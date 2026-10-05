# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 43.3. Weights run: 0.9s. Verify run: 0.8s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.126, intellect=0.577 ± 0.022, crit=0.051 ± 0.002 per rating point (14 rating = 1%, 0.715 per %), hit=0.170 ± 0.002 per rating point (10 rating = 1%, 1.703 per %), spell_haste=0.818 ± 0.113, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.654 ± 0.126, fire_power=0.351 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.76 DPS) | yes | Shadow Goggles (4373, -3.28 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.2 spell_power points (1.29 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.78 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.51 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.50 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.9 spell_power points (1.00 DPS) | yes | Green Woolen Robe (6243, -0.40 DPS) [crafted]; Mystic's Wrap (14369, -0.49 DPS) [world_drop]; Gray Woolen Robe (2585, -1.63 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 2.9 spell_power points (0.37 DPS) | yes | Mystic's Bracelets (14366, -0.22 DPS) [world_drop]; Repurposed Hair Band (281256, -0.22 DPS) [quest]; Bright Bracers (3647, -0.22 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.89 DPS) | yes | Pristine Gloves (253913, -0.16 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS, sim-verified) [world]; Tomb Robber's Gloves (280096, -0.45 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.3 spell_power points (0.80 DPS) | yes | Keller's Girdle (2911, -0.21 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.33 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.04 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (43.3 DPS) | yes | Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Rumpled Kilt (274741, -0.56 DPS) [vendor]; Abomination Skin Leggings (23173, -0.76 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.3 spell_power points (1.18 DPS) | yes | Pristine Boots (253889, -0.58 DPS) [crafted]; Red Woolen Boots (4313, -0.67 DPS) [crafted]; Feather Padded Treads (285345, -0.71 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.78 DPS) | yes | Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Sludge-Stained Band (286535, -0.40 DPS) [world]; Volcanic Rock Ring (12053, -0.56 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.63 DPS) | yes | Sludge-Stained Band (286535, -0.25 DPS) [world]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.29 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 5.8 spell_power points (0.73 DPS) | yes | Channeler's Staff (4437, -0.15 DPS) [world]; Lesser Staff of the Spire (1300, -0.29 DPS) [world]; Staff of Westfall (2042, -0.37 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 179.1 spell_power points (22.69 DPS) | yes | Skycaller (12984, -2.39 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.40 DPS) [dungeon]; Deepblaze (279896, -4.16 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 66.2. Weights run: 1.0s. Verify run: 0.8s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.113, intellect=0.241 ± 0.007, crit=0.049 ± 0.003 per rating point (14 rating = 1%, 0.682 per %), hit=0.103 ± 0.001 per rating point (10 rating = 1%, 1.033 per %), spell_haste=0.638 ± 0.111, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.833 ± 0.114, fire_power=0.166 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.80 DPS) | yes | Silk Headband (7050, -0.57 DPS, sim-verified) [crafted]; Enchanter's Cowl (4322, -0.66 DPS) [crafted]; Embalmed Shroud (7691, -0.76 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 spell_power points (2.15 DPS) | yes | Darkspear Warding Pendant (272075, -1.83 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.91 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.91 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.2 spell_power points (2.85 DPS) | yes | Invoker's Mantle (215365, -0.76 DPS) [crafted]; Fairywing Mantle (9536, -0.76 DPS) [quest]; Death Speaker Mantle (6685, -0.79 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.27 DPS) | yes | Prelacy Cape (7004, -0.25 DPS) [quest]; Caretaker's Cape (19533, -0.25 DPS) [rep]; Heavy Woolen Cloak (4311, -0.65 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.31 DPS) | yes | Green Silk Armor (7065, -0.22 DPS) [crafted]; Death Speaker Robes (6682, -0.85 DPS) [dungeon]; Pristine Gown (253961, -1.10 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.29 DPS) | yes | Glowing Magical Bracelets (13106, -1.77 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.93 DPS) [world_drop]; Stonecloth Bindings (14416, -1.99 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.78 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Town Clerk's Mittens (270029, -0.09 DPS) [quest]; Gnoll Casting Gloves (892, -0.25 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.7 spell_power points (2.99 DPS) | yes | Belt of Arugal (6392, -0.53 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.90 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.95 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.06 DPS) | yes | Abomination Skin Leggings (23173, -0.27 DPS) [dungeon]; Pristine Leggings (253987, -0.85 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.16 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.7 spell_power points (2.21 DPS) | yes | Acidic Walkers (9454, -0.45 DPS) [dungeon]; Nimbus Boots (6998, -0.68 DPS) [quest]; Spidersilk Boots (4320, -1.95 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.78 DPS) | yes | Minor Channeling Ring (1449, -0.39 DPS) [quest]; Electrocutioner Lagnut (9447, -1.02 DPS) [dungeon]; Sludge-Stained Band (286535, -1.02 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.53 DPS) | yes | Electrocutioner Lagnut (9447, -0.76 DPS) [dungeon]; Sludge-Stained Band (286535, -0.76 DPS) [world]; Minor Channeling Ring (1449, -1.65 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.29 DPS) | yes | Glimmering Staff (249392, -1.15 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.68 DPS) [world_drop]; Channeler's Staff (4437, -1.80 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.4 spell_power points (2.15 DPS) | yes | Eye of Paleth (2943, -1.13 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -1.13 DPS) [world]; Dwarven Tome (279898, -1.21 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 133.3 spell_power points (33.97 DPS) | yes | Starfaller (13063, -0.83 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.70 DPS) [crafted]; Gravestone Scepter (7001, -4.97 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 117.3. Weights run: 1.0s. Verify run: 0.7s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.273, intellect=0.462 ± 0.017, crit=0.089 ± 0.004 per rating point (14 rating = 1%, 1.241 per %), hit=0.198 ± 0.002 per rating point (10 rating = 1%, 1.975 per %), spell_haste=not significant (0.432 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.422 ± 0.273), fire_power=0.574 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.14 DPS) | yes | Augural Shroud (2620, -1.23 DPS, sim-verified) [world]; Living Cowl (5608, -1.96 DPS) [world]; Holy Shroud (2721, -2.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (2.39 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.68 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.60 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.60 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (3.22 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Inquisitor's Shawl (19507, -0.04 DPS) [dungeon]; Berylline Pads (4197, -0.38 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.2 spell_power points (3.22 DPS) | yes | Long Silken Cloak (4326, -0.52 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -1.19 DPS) [crafted]; Icy Cloak (4327, -1.51 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.8 spell_power points (6.06 DPS) | yes | Dreamweave Vest (10021, -0.64 DPS) [crafted]; Elemental Raiment (9434, -0.92 DPS) [world_drop]; Robe of Power (7054, -1.28 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.20 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.49 DPS) [quest]; Windchaser Cuffs (14429, -1.18 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (4.86 DPS) | yes | Black Mageweave Gloves (10003, -1.19 DPS) [crafted]; Red Mageweave Gloves (10018, -1.28 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.11 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.8 spell_power points (3.88 DPS) | yes | Deathmage Sash (10771, -0.60 DPS, sim-verified) [dungeon]; Star Belt (4329, -0.70 DPS) [crafted]; Gilded Cord (254037, -1.02 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.5 spell_power points (4.78 DPS) | yes | Crimson Silk Pantaloons (7062, -1.04 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.68 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.85 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.87 DPS) | yes | Gilded Slippers (254001, -1.67 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.71 DPS) [crafted]; Acidic Walkers (9454, -3.74 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (3.13 DPS) | yes | Ring of Forlorn Spirits (2043, -1.17 DPS) [quest]; Reedknot Ring (9622, -1.41 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.66 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.20 DPS) | yes | Ring of Forlorn Spirits (2043, -0.24 DPS) [quest]; Reedknot Ring (9622, -0.49 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.73 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (117.3 DPS) | yes | Scorn's Focal Dagger (23168, -2.69 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.17 DPS) [quest]; Gut Ripper (2164, -5.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 163.7 spell_power points (40.06 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.47 DPS) [dungeon]; Twisted Nether Wand (249144, -4.59 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 152.9. Weights run: 1.0s. Verify run: 0.8s. 417 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.605, intellect=1.132 ± 0.038, crit=0.187 ± 0.010 per rating point (14 rating = 1%, 2.617 per %), hit=0.418 ± 0.005 per rating point (10 rating = 1%, 4.179 per %), spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.6 spell_power points (6.45 DPS) | yes | Dreamweave Circlet (10041, -1.04 DPS, sim-verified) [crafted]; Chief Architect's Monocle (11839, -1.72 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Hat (220889, -1.77 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 15.8 spell_power points (2.45 DPS) | yes | Scorn's Icy Choker (23169, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.47 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.70 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.4 spell_power points (5.17 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.46 DPS) [crafted]; Inquisitor's Shawl (19507, -1.81 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.8 spell_power points (3.22 DPS) | yes | Runecloth Cloak (13860, -0.42 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.77 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.04 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.6 spell_power points (6.45 DPS) | yes | Runecloth Tunic (13857, -1.89 DPS) [crafted]; Robe of the Magi (1716, -1.99 DPS) [world_drop]; Runecloth Robe (13858, -2.21 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Bloodband Bracers (11469, -0.28 DPS) [quest]; Nethergeld Cuffs (254061, -0.32 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.53 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 31.9 spell_power points (4.94 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.25 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.45 DPS) [crafted]; Red Mageweave Gloves (10018, -1.48 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 27.5 spell_power points (4.26 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.41 DPS) [dungeon]; Deathmage Sash (10771, -0.55 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.3 spell_power points (5.32 DPS) | yes | Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.70 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -4.24 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.72 DPS) | yes | Gilded Sandals (254107, -0.44 DPS) [crafted]; Southsea Mojo Boots (20641, -0.55 DPS) [quest]; Sergeant Major's Dreadweave Boots (220891, -2.45 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Philanthropist's Ring (281635, -0.03 DPS) [quest]; Mindseye Circle (10634, -0.53 DPS) [dungeon]; Band of the Unicorn (7553, -0.62 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.9 spell_power points (2.62 DPS) | yes | Mindseye Circle (10634, -0.52 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Philanthropist's Ring (281635, -1.19 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.93 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Soul Harvester (20536, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -2.99 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (152.9 DPS) | yes | Wand of Allistarj (13065, -2.35 DPS) [world_drop]; Flaming Incinerator (9483, -3.55 DPS) [dungeon]; Pyric Caduceus (11748, -4.94 DPS, sim-verified) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; ranged: Noxious Shooter

No-known-source sample (15 of 417, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 278.6. Weights run: 3.1s. Verify run: 0.8s. 1075 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± 0.618), intellect=not significant (-2.689 ± 0.038), crit=not significant (-0.544 ± 0.013) per rating point (14 rating = 1%, -7.618 per %), hit=not significant (-1.130 ± 0.006) per rating point (10 rating = 1%, -11.297 per %), spell_haste=not significant (-2.762 ± 0.468), spell_penetration=not significant (-0.000 ± 0.000), shadow_power=not significant (2.596 ± 0.618), fire_power=not significant (-1.624 ± 0.003)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | 49.3 spell_power points | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -0.83 DPS) [pvp]; Magister's Crown (16686, -4.45 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.7 spell_power points | yes | Beads of Ogre Mojo (22149, -0.49 DPS) [quest]; Archlight Talisman (15856, -0.99 DPS) [quest]; Kezan's Taint (19604, -1.03 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.6 spell_power points | yes | Field Marshal's Dreadweave Shoulders (231583, -0.36 DPS) [pvp]; Burial Shawl (18681, -1.31 DPS) [dungeon]; Darkspear Shoulderpads (272103, -7.07 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.2 spell_power points | yes | Hide of the Wild (18510, -0.61 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.73 DPS) [world]; Spritecaster Cape (11623, -1.31 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 56.2 spell_power points | yes | Field Marshal's Dreadweave Robe (231582, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -1.32 DPS) [pvp]; Magister's Robes (16688, -9.36 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.1 spell_power points | yes | Sublime Wristguards (18497, -1.20 DPS) [dungeon]; Runecloth Cuffs (254123, -1.35 DPS) [crafted]; Marshal's Dreadweave Cuffs (17582, -1.48 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 38.6 spell_power points | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dreadweave Gloves (231586, -0.27 DPS) [pvp]; Sandworm Skin Gloves (20716, -0.91 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 54.1 spell_power points | yes | Magister's Belt (16685, -3.61 DPS) [dungeon]; Belt of the Archmage (18405, -3.62 DPS, sim-verified) [crafted]; Dustfeather Sash (12589, -3.83 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.85 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -0.90 DPS) [pvp] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 36.1 spell_power points | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.02 DPS) [pvp]; Omnicast Boots (11822, -3.61 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.32 DPS) [quest]; Maiden's Circle (13001, -1.32 DPS) [world_drop]; Naglering (11669, -8.15 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.46 DPS) [quest]; Maiden's Circle (13001, -0.46 DPS) [world_drop]; Naglering (11669, -7.86 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.08 DPS) [vendor]; Serenity Field (272439, -2.32 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -2.63 DPS, sim-verified) [quest] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Royal Seal of Eldre'Thalas (18467, -3.18 DPS, sim-verified) [quest] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.41 DPS) [world]; Teebu's Blazing Longsword (1728, -11.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.52 DPS) [dungeon]; Sparkling Crystal Wand (20672, -0.87 DPS) [world]; Torch of Light (279246, -5.93 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Dragonrider Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1075, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 41.8. Weights run: 0.9s. Verify run: 0.8s. 139 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.126, intellect=0.577 ± 0.022, crit=0.051 ± 0.002 per rating point (14 rating = 1%, 0.715 per %), hit=0.170 ± 0.002 per rating point (10 rating = 1%, 1.703 per %), spell_haste=0.818 ± 0.113, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.654 ± 0.126, fire_power=0.351 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.76 DPS) | yes | Shadow Goggles (4373, -3.00 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.2 spell_power points (1.29 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.37 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.78 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.51 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.44 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.9 spell_power points (1.00 DPS) | yes | Green Woolen Robe (6243, -0.40 DPS) [crafted]; Mystic's Wrap (14369, -0.49 DPS) [world_drop]; Gray Woolen Robe (2585, -1.44 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.5 spell_power points (0.44 DPS) | yes | Mindthrust Bracers (1974, -0.07 DPS) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest]; Bright Bracers (3647, -0.15 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.89 DPS) | yes | Pristine Gloves (253913, -0.16 DPS) [crafted]; Gnoll Casting Gloves (892, -0.26 DPS, sim-verified) [world]; Blight Gloves (279877, -0.38 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.3 spell_power points (0.80 DPS) | yes | Keller's Girdle (2911, -0.21 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.33 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.94 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Rumpled Kilt (274741, -0.56 DPS) [vendor]; Abomination Skin Leggings (23173, -0.82 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.3 spell_power points (1.18 DPS) | yes | Pristine Boots (253889, -0.58 DPS) [crafted]; Red Woolen Boots (4313, -0.67 DPS) [crafted]; Feather Padded Treads (285345, -0.72 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.63 DPS) | yes | Loop of Sacrifice (281673, -0.27 DPS) [quest]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.44 DPS, sim-verified) [dungeon] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Loop of Sacrifice (281673, -0.01 DPS) [quest]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.69 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 5.8 spell_power points (0.73 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.15 DPS) [world]; Lesser Staff of the Spire (1300, -0.29 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 179.1 spell_power points (22.69 DPS) | yes | Skycaller (12984, -2.23 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.40 DPS) [dungeon]; Sizzle Stick (8071, -4.41 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 139, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 64.7. Weights run: 1.0s. Verify run: 0.8s. 231 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.113, intellect=0.241 ± 0.007, crit=0.049 ± 0.003 per rating point (14 rating = 1%, 0.682 per %), hit=0.103 ± 0.001 per rating point (10 rating = 1%, 1.033 per %), spell_haste=0.638 ± 0.111, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.833 ± 0.114, fire_power=0.166 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.80 DPS) | yes | Silk Headband (7050, -0.62 DPS, sim-verified) [crafted]; Enchanter's Cowl (4322, -0.66 DPS) [crafted]; Embalmed Shroud (7691, -0.76 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 spell_power points (2.15 DPS) | yes | Darkspear Warding Pendant (272075, -1.90 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.91 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.91 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.2 spell_power points (2.85 DPS) | yes | Death Speaker Mantle (6685, -0.60 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.76 DPS) [crafted]; Fairywing Mantle (9536, -0.76 DPS) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.27 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.25 DPS) [crafted]; Battle Healer's Cloak (19529, -0.25 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (3.31 DPS) | yes | Green Silk Armor (7065, -0.22 DPS) [crafted]; Death Speaker Robes (6682, -0.85 DPS) [dungeon]; Pristine Gown (253961, -1.10 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.29 DPS) | yes | Glowing Magical Bracelets (13106, -1.76 DPS, sim-verified) [world_drop]; Owlbeard Bracers (16981, -1.92 DPS) [quest]; Nightsky Wristbands (6407, -1.93 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.2 spell_power points (1.84 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.31 DPS) [world]; Truefaith Gloves (7049, -0.38 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.7 spell_power points (2.99 DPS) | yes | Warsong Sash (16975, -0.18 DPS) [quest]; Belt of Arugal (6392, -0.51 DPS) [dungeon]; Invoker's Cord (215366, -0.90 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.06 DPS) | yes | Abomination Skin Leggings (23173, -0.27 DPS) [dungeon]; Pristine Leggings (253987, -0.85 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.16 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.7 spell_power points (2.21 DPS) | yes | Acidic Walkers (9454, -0.45 DPS) [dungeon]; Boots of the Enchanter (4325, -0.94 DPS) [crafted]; Spidersilk Boots (4320, -1.85 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.78 DPS) | yes | Electrocutioner Lagnut (9447, -1.02 DPS) [dungeon]; Sludge-Stained Band (286535, -1.02 DPS) [world]; Sacred Band (6669, -1.27 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.53 DPS) | yes | Electrocutioner Lagnut (9447, -0.76 DPS) [dungeon]; Sacred Band (6669, -1.02 DPS) [quest]; Sludge-Stained Band (286535, -2.34 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.29 DPS) | yes | Twisted Chanter's Staff (890, -1.68 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.68 DPS) [quest]; Glimmering Staff (249392, -1.93 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.4 spell_power points (2.15 DPS) | yes | Orb of Souls (249395, -1.13 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -1.27 DPS, sim-verified) [world]; Tome of the Darkspear Prophecy (272090, -1.40 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 133.3 spell_power points (33.97 DPS) | yes | Starfaller (13063, -0.58 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.70 DPS) [crafted]; Gravestone Scepter (7001, -4.97 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 115.9. Weights run: 1.0s. Verify run: 0.7s. 308 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.273, intellect=0.462 ± 0.017, crit=0.089 ± 0.004 per rating point (14 rating = 1%, 1.241 per %), hit=0.198 ± 0.002 per rating point (10 rating = 1%, 1.975 per %), spell_haste=not significant (0.432 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.422 ± 0.273), fire_power=0.574 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.14 DPS) | yes | Augural Shroud (2620, -1.58 DPS, sim-verified) [world]; Living Cowl (5608, -1.96 DPS) [world]; Holy Shroud (2721, -2.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (2.39 DPS) | yes | Triune Amulet (7722, -1.60 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.60 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.95 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.2 spell_power points (3.22 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Inquisitor's Shawl (19507, -0.04 DPS) [dungeon]; Berylline Pads (4197, -0.38 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.2 spell_power points (3.22 DPS) | yes | Guardian Cloak (5965, -1.19 DPS) [crafted]; Icy Cloak (4327, -1.51 DPS) [crafted]; Long Silken Cloak (4326, -1.60 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.8 spell_power points (6.06 DPS) | yes | Dreamweave Vest (10021, -0.73 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.92 DPS) [world_drop]; Robe of Power (7054, -1.28 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.20 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Condor Bracers (15864, -0.49 DPS) [quest]; Radiant Silver Bracers (4545, -1.22 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 spell_power points (4.86 DPS) | yes | Black Mageweave Gloves (10003, -1.19 DPS) [crafted]; Red Mageweave Gloves (10018, -1.28 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.11 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.8 spell_power points (3.88 DPS) | yes | Star Belt (4329, -0.70 DPS) [crafted]; Gilded Cord (254037, -1.02 DPS) [crafted]; Deathmage Sash (10771, -1.52 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.5 spell_power points (4.78 DPS) | yes | Crimson Silk Pantaloons (7062, -1.64 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.68 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.85 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.87 DPS) | yes | Gilded Slippers (254001, -1.77 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -3.71 DPS) [crafted]; Acidic Walkers (9454, -3.74 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (3.13 DPS) | yes | Reedknot Ring (9622, -1.41 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.66 DPS) [vendor]; Black Widow Band (6199, -2.33 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.20 DPS) | yes | Sea Giant's Toe Ring (274746, -0.73 DPS) [vendor]; Reedknot Ring (9622, -1.39 DPS, sim-verified) [quest]; Black Widow Band (6199, -1.41 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (115.9 DPS) | yes | Scorn's Focal Dagger (23168, -2.69 DPS) [dungeon]; Staff of Dar'Orahil (15106, -3.17 DPS) [quest]; Gut Ripper (2164, -6.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 163.7 spell_power points (40.06 DPS) | yes | Umbral Wand (5216, -0.61 DPS, sim-verified) [dungeon]; Earthen Rod (9381, -4.47 DPS) [dungeon]; Twisted Nether Wand (249144, -4.59 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 149.7. Weights run: 1.0s. Verify run: 0.8s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.605, intellect=1.132 ± 0.038, crit=0.187 ± 0.010 per rating point (14 rating = 1%, 2.617 per %), hit=0.418 ± 0.005 per rating point (10 rating = 1%, 4.179 per %), spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.6 spell_power points (6.45 DPS) | yes | Dreamweave Circlet (10041, -0.73 DPS, sim-verified) [crafted]; Chief Architect's Monocle (11839, -1.72 DPS) [dungeon]; Blood Guard's Dreadweave Hat (220907, -1.77 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 15.8 spell_power points (2.45 DPS) | yes | Scorn's Icy Choker (23169, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.47 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.70 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.4 spell_power points (5.17 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.46 DPS) [crafted]; Inquisitor's Shawl (19507, -1.81 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.2 spell_power points (3.44 DPS) | yes | Spritecaster Cape (11623, -0.22 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.46 DPS) [dungeon]; Runecloth Cloak (13860, -0.64 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.6 spell_power points (6.45 DPS) | yes | Runecloth Tunic (13857, -1.89 DPS) [crafted]; Robe of the Magi (1716, -1.99 DPS) [world_drop]; Runecloth Robe (13858, -2.58 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.28 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 31.9 spell_power points (4.94 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.25 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.45 DPS) [crafted]; Red Mageweave Gloves (10018, -1.48 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 27.5 spell_power points (4.26 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.41 DPS) [dungeon]; Deathmage Sash (10771, -0.55 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.3 spell_power points (5.32 DPS) | yes | Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.70 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -4.26 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.72 DPS) | yes | Gilded Sandals (254107, -0.44 DPS) [crafted]; Southsea Mojo Boots (20641, -0.55 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -2.88 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Philanthropist's Ring (281635, -0.03 DPS) [quest]; Mindseye Circle (10634, -0.53 DPS) [dungeon]; Band of the Unicorn (7553, -0.62 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.9 spell_power points (2.62 DPS) | yes | Mindseye Circle (10634, -0.52 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Philanthropist's Ring (281635, -1.09 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.93 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.41 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.58 DPS, sim-verified) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Soul Harvester (20536, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -1.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.35 DPS) [world_drop]; Pyric Caduceus (11748, -3.05 DPS, sim-verified) [dungeon]; Flaming Incinerator (9483, -3.55 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Spellshock Leggings; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Eternal Darkness; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 281.9. Weights run: 3.1s. Verify run: 0.8s. 1063 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=not significant (1.000 ± 0.618), intellect=not significant (-2.689 ± 0.038), crit=not significant (-0.544 ± 0.013) per rating point (14 rating = 1%, -7.618 per %), hit=not significant (-1.130 ± 0.006) per rating point (10 rating = 1%, -11.297 per %), spell_haste=not significant (-2.762 ± 0.468), spell_penetration=not significant (-0.000 ± 0.000), shadow_power=not significant (2.596 ± 0.618), fire_power=not significant (-1.624 ± 0.003)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Deathmist Mask (226909) | Saving the Best for Last [quest] | 49.3 spell_power points | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -0.83 DPS) [pvp]; Magister's Crown (16686, -5.01 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.7 spell_power points | yes | Beads of Ogre Mojo (22149, -0.49 DPS) [quest]; Archlight Talisman (15856, -0.99 DPS) [quest]; Kezan's Taint (19604, -1.03 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.6 spell_power points | yes | Warlord's Dreadweave Mantle (231592, -0.36 DPS) [pvp]; Burial Shawl (18681, -1.31 DPS) [dungeon]; Darkspear Shoulderpads (272103, -6.04 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.2 spell_power points | yes | Hide of the Wild (18510, -0.61 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.73 DPS) [world]; Deep Woodlands Cloak (19121, -1.09 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 56.2 spell_power points | yes | Warlord's Dreadweave Robe (231591, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -1.32 DPS) [pvp]; Magister's Robes (16688, -10.00 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.1 spell_power points | yes | Sublime Wristguards (18497, -1.20 DPS) [dungeon]; Runecloth Cuffs (254123, -1.35 DPS) [crafted]; General's Dreadweave Bracers (17587, -1.48 DPS) [pvp] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Sandworm Skin Gloves (20716, -0.02 DPS) [quest]; Raider Handwraps (272097, -3.21 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 54.1 spell_power points | yes | Magister's Belt (16685, -3.61 DPS) [dungeon]; Dustfeather Sash (12589, -3.83 DPS) [dungeon]; Belt of the Archmage (18405, -4.15 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.85 DPS) [dungeon]; Outrider's Silk Leggings (22747, -3.49 DPS, sim-verified) [rep] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 36.1 spell_power points | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Walkers (227098, -0.53 DPS) [pvp]; Omnicast Boots (11822, -1.37 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.32 DPS) [quest]; Maiden's Circle (13001, -1.32 DPS) [world_drop]; Naglering (11669, -7.83 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.46 DPS) [quest]; Maiden's Circle (13001, -0.46 DPS) [world_drop]; Naglering (11669, -7.03 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -1.40 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18467, -0.93 DPS) [quest]; Weakness Analyzer (272438, -1.08 DPS) [vendor]; Draconic Infused Emblem (22268, -4.13 DPS, sim-verified) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Trindlehaven Staff (13161, -0.29 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -14.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+6.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Bonecreeper Stylus (13938, -0.52 DPS) [dungeon]; Sparkling Crystal Wand (20672, -0.87 DPS) [world]; Torch of Light (279246, -6.47 DPS, sim-verified) [crafted] |

**New at 60:** head: Deathmist Mask; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Dragonrider Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1063, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

