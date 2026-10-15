# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 28.6. Weights run: 0.7s. Verify run: 0.6s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.548 ± 0.009, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.118 per %), hit=0.301 ± 0.009 per rating point (10 rating = 1%, 3.009 per %), spell_haste=not significant (0.255 ± 0.146), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.49 DPS) | yes | Shadow Goggles (4373, -1.12 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.9 spell_power points (0.82 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.49 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Caretaker's Cape (20428, -0.08 DPS) [rep]; Pearl-clasped Cloak (5542, -0.18 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.7 spell_power points (0.64 DPS) | yes | Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.31 DPS) [crafted]; Gray Woolen Robe (2585, -0.51 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 2.7 spell_power points (0.23 DPS) | yes | Bright Bracers (3647, -0.05 DPS) [world_drop]; Mystic's Bracelets (14366, -0.14 DPS) [world_drop]; Repurposed Hair Band (281256, -0.14 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.58 DPS) | yes | Pristine Gloves (253913, -0.11 DPS) [crafted]; Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Tomb Robber's Gloves (280096, -0.31 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.2 spell_power points (0.51 DPS) | yes | Keller's Girdle (2911, -0.15 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.36 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.4 spell_power points (1.10 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.53 DPS) [dungeon]; Rumpled Kilt (274741, -0.69 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.2 spell_power points (0.76 DPS) | yes | Pristine Boots (253889, -0.10 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.43 DPS) [crafted]; Woolen Boots (2583, -0.50 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.1 spell_power points (0.50 DPS) | yes | Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; The 1 Ring (8350, -0.46 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.41 DPS) | yes | Volcanic Rock Ring (12053, -0.28 DPS) [world_drop]; The 1 Ring (8350, -0.37 DPS) [world]; Lavishly Jeweled Ring (1156, -0.43 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 5.5 spell_power points (0.45 DPS) | yes | Lesser Staff of the Spire (1300, -0.18 DPS) [world]; Witching Stave (1484, -0.23 DPS) [dungeon]; Staff of Westfall (2042, -0.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 273.8 spell_power points (22.56 DPS) | yes | Skycaller (12984, -1.00 DPS) [world_drop]; Deepblaze (279896, -4.03 DPS) [quest]; Sizzle Stick (8071, -4.50 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 58.3. Weights run: 0.8s. Verify run: 0.6s. 256 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.563 ± 0.007, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.402 per %), hit=0.223 ± 0.008 per rating point (10 rating = 1%, 2.226 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (58.3 DPS) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Enchanter's Cowl (4322, -0.93 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (1.18 DPS) | yes | Darkspear Warding Pendant (272075, -0.86 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.92 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.04 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (1.60 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Magician's Mantle (12998, -0.45 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.57 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272078, -0.06 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.3 spell_power points (1.85 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Death Speaker Robes (6682, -0.36 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.02 DPS) | yes | Nightsky Wristbands (6407, -0.64 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.65 DPS, sim-verified) [world_drop]; Stonecloth Bindings (14416, -0.70 DPS) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (0.80 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.01 DPS) [quest]; Truefaith Gloves (7049, -0.04 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.7 spell_power points (1.44 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Crimson Silk Belt (7055, -0.31 DPS) [crafted]; Invoker's Cord (215366, -0.33 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.5 spell_power points (1.53 DPS) | yes | Pristine Leggings (253987, -0.29 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.38 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.47 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.24 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Glinteye Slippers (273024, -0.45 DPS) [dungeon]; Acidic Walkers (9454, -1.02 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.00 DPS) | yes | Minor Channeling Ring (1449, -0.31 DPS) [quest]; Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.32 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.80 DPS) | yes | Ring of Power Regulation (273030, -0.11 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.11 DPS) [vendor]; Minor Channeling Ring (1449, -0.47 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.09 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.82 DPS) [quest]; Glimmering Staff (249392, -3.38 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 291.6 spell_power points (33.12 DPS) | yes | Wand of Mana Concentration (273043, -0.81 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.55 DPS) [crafted]; Gravestone Scepter (7001, -4.12 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 256, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 129.6. Weights run: 0.8s. Verify run: 0.7s. 336 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.091 ± 0.012, crit=0.014 ± 0.001 per rating point (14 rating = 1%, 0.200 per %), hit=0.273 ± 0.014 per rating point (10 rating = 1%, 2.735 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.6 spell_power points (5.05 DPS) | yes | Spellpower Goggles Xtreme (10502, -1.06 DPS) [crafted]; Thinking Cap (2624, -1.52 DPS) [world]; Augural Shroud (2620, -1.99 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (2.99 DPS) | yes | Triune Amulet (7722, -1.47 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.54 DPS) [vendor]; Scorn's Icy Choker (23169, -1.59 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.2 spell_power points (4.03 DPS) | yes | Bloodmage Mantle (7684, -0.45 DPS) [dungeon]; Death Speaker Mantle (6685, -0.61 DPS) [dungeon]; Green Silken Shoulders (7057, -1.49 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 12.0 spell_power points (2.29 DPS) | yes | Long Silken Cloak (4326, -0.10 DPS) [crafted]; Guardian Cloak (5965, -0.10 DPS) [crafted]; Blackforge Cape (6424, -0.62 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.5 spell_power points (5.43 DPS) | yes | Dreamweave Vest (10021, -0.14 DPS) [crafted]; Robe of Power (7054, -0.28 DPS) [crafted]; Green Silk Armor (7065, -1.02 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 9.8 spell_power points (1.87 DPS) | yes | Arcane Runed Bracers (4744, +0.00 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Mistscape Bracers (4045, -0.21 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.4 spell_power points (4.26 DPS) | yes | Red Mageweave Gloves (10018, -0.09 DPS) [crafted]; Stormcloth Gloves (10011, -1.00 DPS) [crafted]; Gilded Handwraps (254021, -1.28 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.4 spell_power points (4.45 DPS) | yes | Highlander's Cloth Girdle (20098, -0.95 DPS) [rep]; Gilded Cord (254037, -1.26 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.1 spell_power points (5.16 DPS) | yes | Crimson Silk Pantaloons (7062, -1.21 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.78 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.16 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.57 DPS) | yes | Gilded Slippers (254001, -1.08 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.95 DPS) [dungeon]; Spidersilk Boots (4320, -2.40 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.94 DPS) | yes | Ring of Forlorn Spirits (2043, -1.42 DPS) [quest]; Voodoo Band (1996, -1.49 DPS) [world]; Snake Hoop (6750, -3.94 DPS, sim-verified) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-verified (129.6 DPS) | yes | Ring of Forlorn Spirits (2043, -0.19 DPS) [quest]; Voodoo Band (1996, -0.26 DPS) [world]; Snake Hoop (6750, -1.49 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.42 DPS) [dungeon]; Spine of the Devourer (273040, -3.18 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 16.9 spell_power points (3.22 DPS) | yes | Thrash's Trash (276204, -0.55 DPS) [vendor]; Orb of Mystic Insight (249394, -0.64 DPS) [crafted]; Orb of the Forgotten Seer (7685, -0.99 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 213.4 spell_power points (40.63 DPS) | yes | Earthen Rod (9381, -1.08 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.96 DPS) [dungeon]; Twisted Nether Wand (249144, -5.48 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 336, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 210.1. Weights run: 0.8s. Verify run: 0.9s. 430 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.473 ± 0.019, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.393 per %), hit=0.497 ± 0.026 per rating point (10 rating = 1%, 4.974 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.5 spell_power points (8.58 DPS) | yes | Dreamweave Circlet (10041, -2.15 DPS, sim-verified) [crafted]; Papal Fez (9431, -2.73 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Hat (220889, -2.90 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.6 spell_power points (3.65 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Scorn's Icy Choker (23169, -0.85 DPS) [dungeon]; Mindburst Medallion (11196, -1.02 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.5 spell_power points (6.99 DPS) | yes | Inquisitor's Shawl (19507, -2.37 DPS) [dungeon]; Green Silken Shoulders (7057, -2.71 DPS) [crafted]; Red Mageweave Shoulders (10029, -2.85 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.8 spell_power points (4.04 DPS) | yes | Darkspear Raider's Cloak (272076, -0.39 DPS) [vendor]; Big Voodoo Cloak (8216, -0.81 DPS) [crafted]; Runecloth Cloak (13860, -1.12 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.5 spell_power points (8.58 DPS) | yes | Runecloth Tunic (13857, -2.70 DPS) [crafted]; Robe of Power (7054, -2.97 DPS) [crafted]; Runecloth Robe (13858, -3.42 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 18.3 spell_power points (3.23 DPS) | yes | Forgotten Wraps (9433, -0.10 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.10 DPS) [quest]; Nethergeld Cuffs (254061, -0.17 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.4 spell_power points (6.97 DPS) | yes | Virtuous Hands (226958, -1.05 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.33 DPS) [vendor]; Red Mageweave Gloves (10018, -2.42 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 34.0 spell_power points (6.01 DPS) | yes | Satyrmane Sash (17755, -0.93 DPS) [dungeon]; Runecloth Belt (13856, -1.47 DPS) [crafted]; Deathmage Sash (10771, -2.16 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 31.7 spell_power points (5.61 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.28 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.38 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.80 DPS) [crafted] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Southsea Mojo Boots (20641, -0.01 DPS) [quest]; Earthen Silk Slippers (254013, -0.05 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.61 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.4 spell_power points (3.07 DPS) | yes | Band of the Unicorn (7553, -0.77 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.95 DPS) [rep]; Voodoo Band (1996, -1.25 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 13.4 spell_power points (2.37 DPS) | yes | Band of the Unicorn (7553, -0.06 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.24 DPS) [rep]; Voodoo Band (1996, -0.54 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -3.42 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.81 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -3.39 DPS) [dungeon]; Hammer of the Northern Wind (810, -26.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.46 DPS) [world_drop]; Pyric Caduceus (11748, -3.13 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.72 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Dawnspire Cord; feet: Gilded Sandals; finger2: Snake Hoop; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Noxious Shooter

No-known-source sample (15 of 430, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 415.9. Weights run: 0.9s. Verify run: 2.7s. 943 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.292 ± 0.018, crit=0.027 ± 0.002 per rating point (14 rating = 1%, 0.377 per %), hit=0.579 ± 0.029 per rating point (10 rating = 1%, 5.787 per %), spell_haste=not significant (-0.005 ± 0.012), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | sim-verified (415.9 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Magister's Crown (16686, -13.13 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 31.8 spell_power points (8.86 DPS) | yes | Beads of Ogre Mojo (22149, -0.68 DPS, sim-verified) [quest]; Archlight Talisman (15856, -1.92 DPS) [quest]; Chains of the Lich (23125, -2.73 DPS) [dungeon] |
| shoulder | Virtuous Epaulets (226955) | Mokvar [vendor] | sim-verified (415.9 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -6.89 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.1 spell_power points (8.95 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Spritecaster Cape (11623, -2.89 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -3.19 DPS) [vendor] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 49.3 spell_power points (13.72 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Magister's Robes (16688, -12.53 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 32.3 spell_power points (9.01 DPS) | yes | Marshal's Satin Bracers (17606, -2.17 DPS) [pvp]; Runecloth Cuffs (254123, -2.35 DPS) [crafted]; Virtuous Wraps (226953, -2.82 DPS) [vendor] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | sim-verified (415.9 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Raider Handwraps (272097, -8.78 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 59.2 spell_power points (16.50 DPS) | yes | Belt of the Archmage (18405, -3.51 DPS, sim-verified) [crafted]; Magister's Belt (16685, -6.99 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -7.88 DPS) [rep] |
| legs | Virtuous Leggings (226956) | Mokvar [vendor] | sim-verified (415.9 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (22752, -6.78 DPS, sim-verified) [rep] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 42.7 spell_power points (11.89 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Omnicast Boots (11822, -8.51 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.4 spell_power points (12.65 DPS) | yes | Maiden's Circle (13001, -5.12 DPS) [world_drop]; Deep Rooted Ring (19109, -5.71 DPS) [world]; Channeler's Ring (272406, -6.00 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 36.2 spell_power points (10.09 DPS) | yes | Maiden's Circle (13001, -1.37 DPS, sim-verified) [world_drop]; Deep Rooted Ring (19109, -3.15 DPS) [world]; Channeler's Ring (272406, -3.44 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (415.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (415.9 DPS) | yes | Weakness Analyzer (272438, -2.04 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.23 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.45 DPS) [crafted] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (415.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -2.29 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (415.9 DPS) | yes | Stormrager (16997, -0.12 DPS) [quest]; Wand of Biting Cold (19108, -0.61 DPS) [quest]; Torch of Light (279246, -8.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Virtuous Epaulets; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Virtuous Hands; waist: Knowledge of the Timbermaw; legs: Virtuous Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 943, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 325001031300000000-00000000000000000-523120501201300251)

Set DPS (verified): 573.1. Weights run: 1.0s. Verify run: 2.0s. 943 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.014 ± 0.002, crit=0.056 ± 0.004 per rating point (14 rating = 1%, 0.784 per %), hit=0.414 ± 0.034 per rating point (10 rating = 1%, 4.136 per %), spell_haste=-7.363 ± 0.349, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 27.0 spell_power points (10.01 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -0.01 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.15 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.53 DPS) [quest]; Beads of Ogre Mojo (22149, -3.27 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.0 spell_power points (10.37 DPS) | yes | Field Marshal's Satin Epaulets (231621, -1.01 DPS) [vendor]; Argent Shoulders (19059, -1.17 DPS, sim-verified) [crafted]; Virtuous Epaulets (226955, -1.74 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.2 spell_power points (7.50 DPS) | yes | Hide of the Wild (18510, -2.26 DPS) [crafted]; Spritecaster Cape (11623, -2.28 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -2.31 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-verified (573.1 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Virtuous Gown (226960, -8.49 DPS, sim-verified) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.1 spell_power points (8.19 DPS) | yes | Runecloth Cuffs (254123, -4.06 DPS) [crafted]; Arcane Runed Bracers (4744, -4.86 DPS) [quest]; Spidertank Oilrag (9448, -4.86 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.15 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Knight-Lieutenant's Satin Grips (227139, -0.34 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 29.4 spell_power points (10.90 DPS) | yes | Stormpike Cloth Girdle (19094, -4.19 DPS) [rep]; Belt of the Archmage (18405, -5.27 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -5.40 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 33.7 spell_power points (12.48 DPS) | yes | Marshal's Satin Pants (17603, -0.52 DPS) [vendor]; Marshal's Satin Leggings (231619, -0.52 DPS) [vendor]; Ghostweave Pants (14144, -2.85 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.89 DPS) | yes | Marshal's Satin Sandals (17607, -0.28 DPS) [vendor]; Marshal's Satin Treads (231620, -0.28 DPS) [vendor]; Virtuous Slippers (226959, -0.66 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 32.3 spell_power points (11.95 DPS) | yes | Maiden's Circle (13001, -5.25 DPS) [world_drop]; Blessed Band of Light (272407, -6.71 DPS) [vendor]; Lorekeeper's Ring (19522, -6.76 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.2 spell_power points (8.20 DPS) | yes | Maiden's Circle (13001, -1.52 DPS, sim-verified) [world_drop]; Blessed Band of Light (272407, -2.97 DPS) [vendor]; Lorekeeper's Ring (19522, -3.02 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+17.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.21 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.96 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.81 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.68 DPS) [dungeon]; The Lobotomizer (19324, -30.60 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 213.2 spell_power points (78.98 DPS) | yes | Wand of Biting Cold (19108, +0.00 DPS, sim-verified) [quest]; Brilliant Wand (249385, -16.03 DPS) [crafted]; Stormrager (16997, -16.27 DPS) [quest] |

**New at 60:** head: Virtuous Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Truefaith Vestments; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Earthen Silk Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 943, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-542000000000000000)

Set DPS (verified): 33.5. Weights run: 0.7s. Verify run: 0.6s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.548 ± 0.009, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.118 per %), hit=0.301 ± 0.009 per rating point (10 rating = 1%, 3.009 per %), spell_haste=not significant (0.255 ± 0.146), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.49 DPS) | yes | Shadow Goggles (4373, -1.04 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.9 spell_power points (0.82 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.49 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Pearl-clasped Cloak (5542, -0.03 DPS) [crafted]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Battle Healer's Cloak (20427, -0.08 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.7 spell_power points (0.64 DPS) | yes | Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.31 DPS) [crafted]; Gray Woolen Robe (2585, -0.70 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.25 DPS) | yes | Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest]; Bright Bracers (3647, -0.07 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.58 DPS) | yes | Gnoll Casting Gloves (892, -0.08 DPS) [world]; Pristine Gloves (253913, -0.11 DPS) [crafted]; Apothecary Gloves (10919, -0.25 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.2 spell_power points (0.51 DPS) | yes | Keller's Girdle (2911, -0.15 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.49 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.4 spell_power points (1.10 DPS) | yes | Filigreed Pristine Leggings (253937, -0.34 DPS) [crafted]; Silk-threaded Trousers (1929, -0.53 DPS) [dungeon]; Rumpled Kilt (274741, -0.69 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.2 spell_power points (0.76 DPS) | yes | Pristine Boots (253889, -0.16 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.43 DPS) [crafted]; Woolen Boots (2583, -0.50 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.41 DPS) | yes | Loop of Sacrifice (281673, -0.19 DPS) [quest]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop]; The 1 Ring (8350, -0.37 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.3 spell_power points (0.27 DPS) | yes | Loop of Sacrifice (281673, -0.05 DPS) [quest]; Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; The 1 Ring (8350, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 5.5 spell_power points (0.45 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Lesser Staff of the Spire (1300, -0.18 DPS) [world]; Clear Crystal Rod (16894, -0.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 273.8 spell_power points (22.56 DPS) | yes | Skycaller (12984, -2.11 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.50 DPS) [quest]; Greater Magic Wand (11288, -4.89 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (undead, 000000000000000000-00000000000000000-543120301200000000)

Set DPS (verified): 58.3. Weights run: 0.8s. Verify run: 0.6s. 242 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.563 ± 0.007, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.402 per %), hit=0.223 ± 0.008 per rating point (10 rating = 1%, 2.226 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (58.3 DPS) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Enchanter's Cowl (4322, -0.92 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (1.18 DPS) | yes | Darkspear Warding Pendant (272075, -0.86 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.92 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.05 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (1.60 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Magician's Mantle (12998, -0.45 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.57 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.06 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.3 spell_power points (1.85 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Death Speaker Robes (6682, -0.36 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.02 DPS) | yes | Nightsky Wristbands (6407, -0.64 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.68 DPS) [quest]; Glowing Magical Bracelets (13106, -0.92 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.8 spell_power points (1.00 DPS) | yes | Truefaith Gloves (7049, -0.24 DPS) [crafted]; Gnoll Casting Gloves (892, -0.32 DPS) [world]; Serpent Gloves (5970, -0.37 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.7 spell_power points (1.44 DPS) | yes | Belt of Arugal (6392, -0.23 DPS) [dungeon]; Crimson Silk Belt (7055, -0.31 DPS) [crafted]; Warsong Sash (16975, -0.35 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.5 spell_power points (1.53 DPS) | yes | Pristine Leggings (253987, -0.29 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.47 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.54 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.24 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Glinteye Slippers (273024, -0.45 DPS) [dungeon]; Acidic Walkers (9454, -1.11 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.00 DPS) | yes | Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.32 DPS) [vendor]; Solid Gold Band (279845, -0.55 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.80 DPS) | yes | Sea Giant's Toe Ring (274746, +0.00 DPS, sim-verified) [vendor]; Ring of Power Regulation (273030, -0.11 DPS) [dungeon]; Solid Gold Band (279845, -0.34 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.09 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -3.38 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -3.45 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 291.6 spell_power points (33.12 DPS) | yes | Wand of Mana Concentration (273043, -1.69 DPS, sim-verified) [world]; Unstable Power Core (279847, -3.41 DPS) [quest]; Greater Mystic Wand (217287, -3.55 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 000000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 120.3. Weights run: 0.8s. Verify run: 0.7s. 319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.091 ± 0.012, crit=0.014 ± 0.001 per rating point (14 rating = 1%, 0.200 per %), hit=0.273 ± 0.014 per rating point (10 rating = 1%, 2.735 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.6 spell_power points (5.05 DPS) | yes | Spellpower Goggles Xtreme (10502, -1.06 DPS) [crafted]; Thinking Cap (2624, -1.52 DPS) [world]; Augural Shroud (2620, -2.82 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (2.99 DPS) | yes | Triune Amulet (7722, -1.47 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.54 DPS) [vendor]; Scorn's Icy Choker (23169, -1.79 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.2 spell_power points (4.03 DPS) | yes | Bloodmage Mantle (7684, -0.45 DPS) [dungeon]; Death Speaker Mantle (6685, -0.61 DPS) [dungeon]; Green Silken Shoulders (7057, -1.81 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 12.0 spell_power points (2.29 DPS) | yes | Long Silken Cloak (4326, -0.10 DPS) [crafted]; Guardian Cloak (5965, -0.10 DPS) [crafted]; Blackforge Cape (6424, -0.62 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.5 spell_power points (5.43 DPS) | yes | Dreamweave Vest (10021, -0.14 DPS) [crafted]; Robe of Power (7054, -0.28 DPS) [crafted]; Green Silk Armor (7065, -1.02 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 12.7 spell_power points (2.42 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.71 DPS) [dungeon]; Windchaser Cuffs (14429, -0.74 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.4 spell_power points (4.26 DPS) | yes | Red Mageweave Gloves (10018, -0.09 DPS) [crafted]; Stormcloth Gloves (10011, -1.00 DPS) [crafted]; Gilded Handwraps (254021, -1.28 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.4 spell_power points (4.45 DPS) | yes | Gilded Cord (254037, -1.26 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.69 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.1 spell_power points (5.16 DPS) | yes | Crimson Silk Pantaloons (7062, -1.24 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.78 DPS) [dungeon]; Stoneweaver Leggings (9407, -2.16 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.57 DPS) | yes | Gilded Slippers (254001, -0.74 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.95 DPS) [dungeon]; Spidersilk Boots (4320, -2.40 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.94 DPS) | yes | Advisor's Ring (19520, -1.23 DPS) [rep]; Ogremind Ring (1993, -1.49 DPS) [world_drop]; Voodoo Band (1996, -1.49 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 11.5 spell_power points (2.18 DPS) | yes | Advisor's Ring (19520, +0.00 DPS, sim-verified) [rep]; Ogremind Ring (1993, -0.73 DPS) [world_drop]; Voodoo Band (1996, -0.73 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (120.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.42 DPS) [dungeon]; Spine of the Devourer (273040, -3.18 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.66 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Mystic Insight (249394, -0.09 DPS) [crafted]; Tome of the Darkspear Prophecy (272089, -0.86 DPS) [vendor] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 213.4 spell_power points (40.63 DPS) | yes | Earthen Rod (9381, -1.90 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.96 DPS) [dungeon]; Twisted Nether Wand (249144, -5.48 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 523000000000000000-00000000000000000-543120301201300051)

Set DPS (verified): 197.0. Weights run: 0.8s. Verify run: 0.9s. 410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.473 ± 0.019, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.393 per %), hit=0.497 ± 0.026 per rating point (10 rating = 1%, 4.974 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.5 spell_power points (8.58 DPS) | yes | Dreamweave Circlet (10041, -1.18 DPS, sim-verified) [crafted]; Papal Fez (9431, -2.73 DPS) [dungeon]; Blood Guard's Dreadweave Hat (220907, -2.90 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.6 spell_power points (3.65 DPS) | yes | Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon]; Scorn's Icy Choker (23169, -0.85 DPS) [dungeon]; Mindburst Medallion (11196, -1.02 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.5 spell_power points (6.99 DPS) | yes | Red Mageweave Shoulders (10029, -2.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -2.37 DPS) [dungeon]; Green Silken Shoulders (7057, -2.71 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 25.3 spell_power points (4.47 DPS) | yes | Spritecaster Cape (11623, -0.43 DPS) [dungeon]; Runecloth Cloak (13860, -0.79 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.82 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.5 spell_power points (8.58 DPS) | yes | Runecloth Robe (13858, -2.52 DPS, sim-verified) [crafted]; Runecloth Tunic (13857, -2.70 DPS) [crafted]; Robe of Power (7054, -2.97 DPS) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 18.3 spell_power points (3.23 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Shizzle's Nozzle Wiper (11917, -0.10 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 39.4 spell_power points (6.97 DPS) | yes | Virtuous Hands (226958, -0.22 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -2.33 DPS) [vendor]; Red Mageweave Gloves (10018, -2.42 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 34.0 spell_power points (6.01 DPS) | yes | Satyrmane Sash (17755, -0.93 DPS) [dungeon]; Deathmage Sash (10771, -1.28 DPS, sim-verified) [dungeon]; Runecloth Belt (13856, -1.47 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 31.7 spell_power points (5.61 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.28 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.38 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.80 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.2 spell_power points (4.64 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Southsea Mojo Boots (20641, -0.36 DPS) [quest]; Earthen Silk Slippers (254013, -0.40 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.4 spell_power points (3.07 DPS) | yes | Band of the Unicorn (7553, -0.77 DPS) [world_drop]; Advisor's Ring (19519, -0.95 DPS) [rep]; Voodoo Band (1996, -1.25 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 13.4 spell_power points (2.37 DPS) | yes | Band of the Unicorn (7553, -0.06 DPS) [world_drop]; Advisor's Ring (19519, -0.24 DPS) [rep]; Voodoo Band (1996, -0.54 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.18 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.78 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.81 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -3.39 DPS) [dungeon]; Hammer of the Northern Wind (810, -24.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (197.0 DPS) | yes | Wand of Allistarj (13065, -2.46 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.72 DPS) [crafted]; Pyric Caduceus (11748, -3.75 DPS, sim-verified) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Dawnspire Cord; feet: First Sergeant's Dreadweave Boots; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spellshifter Rod; ranged: Noxious Shooter

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 524011031300000000-00000000000000000-543120301201300051)

Set DPS (verified): 427.6. Weights run: 0.9s. Verify run: 2.5s. 935 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.292 ± 0.018, crit=0.027 ± 0.002 per rating point (14 rating = 1%, 0.377 per %), hit=0.579 ± 0.029 per rating point (10 rating = 1%, 5.787 per %), spell_haste=not significant (-0.005 ± 0.012), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | sim-verified (427.6 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -13.37 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 31.8 spell_power points (8.86 DPS) | yes | Beads of Ogre Mojo (22149, -0.81 DPS, sim-verified) [quest]; Archlight Talisman (15856, -1.92 DPS) [quest]; Chains of the Lich (23125, -2.73 DPS) [dungeon] |
| shoulder | Virtuous Epaulets (226955) | Mokvar [vendor] | sim-verified (427.6 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -7.81 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.1 spell_power points (8.95 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Deep Woodlands Cloak (19121, -2.37 DPS) [quest]; Spritecaster Cape (11623, -2.89 DPS) [dungeon] |
| chest | Virtuous Gown (226960) | Mokvar [vendor] | 49.3 spell_power points (13.72 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Magister's Robes (16688, -12.71 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 32.3 spell_power points (9.01 DPS) | yes | General's Satin Bracers (17619, -2.17 DPS) [pvp]; Runecloth Cuffs (254123, -2.35 DPS) [crafted]; Virtuous Wraps (226953, -2.82 DPS) [vendor] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | sim-verified (427.6 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Raider Handwraps (272097, -8.68 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 59.2 spell_power points (16.50 DPS) | yes | Belt of the Archmage (18405, -4.03 DPS, sim-verified) [crafted]; Magister's Belt (16685, -6.99 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -7.88 DPS) [rep] |
| legs | Virtuous Leggings (226956) | Mokvar [vendor] | sim-verified (427.6 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Outrider's Silk Leggings (22747, -6.56 DPS, sim-verified) [rep] |
| feet | Virtuous Slippers (226959) | Mokvar [vendor] | 42.7 spell_power points (11.89 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Omnicast Boots (11822, -8.67 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 45.4 spell_power points (12.65 DPS) | yes | Maiden's Circle (13001, -5.12 DPS) [world_drop]; Deep Rooted Ring (19109, -5.71 DPS) [world]; Channeler's Ring (272406, -6.00 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 36.2 spell_power points (10.09 DPS) | yes | Maiden's Circle (13001, -1.55 DPS, sim-verified) [world_drop]; Deep Rooted Ring (19109, -3.15 DPS) [world]; Channeler's Ring (272406, -3.44 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (427.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (427.6 DPS) | yes | Weakness Analyzer (272438, -2.05 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.23 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.45 DPS) [crafted] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (427.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -2.29 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (427.6 DPS) | yes | Stormrager (16997, -0.12 DPS) [quest]; Wand of Biting Cold (19108, -0.61 DPS) [quest]; Torch of Light (279246, -10.38 DPS, sim-verified) [crafted] |

**New at 60:** head: Virtuous Cowl; neck: Amulet of the Dawn; shoulder: Virtuous Epaulets; back: Arcanoweave Cloak; chest: Virtuous Gown; wrist: Dryad's Wrist Bindings; hands: Virtuous Hands; waist: Knowledge of the Timbermaw; legs: Virtuous Leggings; feet: Virtuous Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; ranged: Brilliant Wand

No-known-source sample (15 of 935, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 325001031300000000-00000000000000000-523120501201300251)

Set DPS (verified): 583.4. Weights run: 1.0s. Verify run: 1.9s. 935 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.014 ± 0.002, crit=0.056 ± 0.004 per rating point (14 rating = 1%, 0.784 per %), hit=0.414 ± 0.034 per rating point (10 rating = 1%, 4.136 per %), spell_haste=-7.363 ± 0.349, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Cowl (226957) | Mokvar [vendor] | 27.0 spell_power points (10.01 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -0.01 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.15 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -2.53 DPS) [quest]; Beads of Ogre Mojo (22149, -3.27 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 28.0 spell_power points (10.37 DPS) | yes | Warlord's Satin Epaulets (231611, -1.01 DPS) [vendor]; Argent Shoulders (19059, -1.11 DPS) [crafted]; Virtuous Epaulets (226955, -1.74 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 20.2 spell_power points (7.50 DPS) | yes | Hide of the Wild (18510, -2.26 DPS) [crafted]; Spritecaster Cape (11623, -2.28 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -2.31 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-verified (583.4 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Robe of the Magi (1716, -1.08 DPS) [world_drop]; Virtuous Gown (226960, -8.35 DPS, sim-verified) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.1 spell_power points (8.19 DPS) | yes | Runecloth Cuffs (254123, -4.06 DPS) [crafted]; Spidertank Oilrag (9448, -4.86 DPS) [dungeon]; Virtuous Wraps (226953, -5.17 DPS) [vendor] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.15 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; Virtuous Hands (226958, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 29.4 spell_power points (10.90 DPS) | yes | Frostwolf Cloth Belt (19090, -4.19 DPS) [rep]; Belt of the Archmage (18405, -5.20 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -5.40 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 33.7 spell_power points (12.48 DPS) | yes | General's Satin Leggings (231614, -0.52 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.61 DPS, sim-verified) [rep]; Ghostweave Pants (14144, -2.85 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.89 DPS) | yes | General's Satin Boots (17618, -0.28 DPS) [vendor]; General's Satin Treads (231610, -0.28 DPS) [vendor]; Virtuous Slippers (226959, -0.66 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 32.3 spell_power points (11.95 DPS) | yes | Maiden's Circle (13001, -5.25 DPS) [world_drop]; Blessed Band of Light (272407, -6.71 DPS) [vendor]; Advisor's Ring (19518, -6.76 DPS) [rep] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 22.2 spell_power points (8.20 DPS) | yes | Maiden's Circle (13001, -1.58 DPS, sim-verified) [world_drop]; Blessed Band of Light (272407, -2.97 DPS) [vendor]; Advisor's Ring (19518, -3.02 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+17.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.21 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.96 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.81 DPS) [crafted] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.68 DPS) [dungeon]; The Lobotomizer (19324, -30.70 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 213.2 spell_power points (78.98 DPS) | yes | Wand of Biting Cold (19108, +0.00 DPS, sim-verified) [quest]; Brilliant Wand (249385, -16.03 DPS) [crafted]; Stormrager (16997, -16.27 DPS) [quest] |

**New at 60:** head: Virtuous Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Truefaith Vestments; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Earthen Silk Slippers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Spire of Hakkar; ranged: Torch of Light

No-known-source sample (15 of 935, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

