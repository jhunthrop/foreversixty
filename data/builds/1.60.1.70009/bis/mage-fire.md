# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 32.9. Weights run: 0.6s. Verify run: 0.6s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.612 ± 0.017, crit=0.136 ± 0.007 per rating point (14 rating = 1%, 1.900 per %), hit=0.395 ± 0.003 per rating point (10 rating = 1%, 3.948 per %), spell_haste=not significant (-0.590 ± 0.325), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.44 DPS) | yes | Shadow Goggles (4373, -0.61 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.5 spell_power points (0.77 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.32 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.29 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Black Whelp Cloak (7283, -0.07 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.1 spell_power points (0.59 DPS) | yes | Green Woolen Robe (6243, -0.24 DPS) [crafted]; Mystic's Wrap (14369, -0.28 DPS) [world_drop]; Gray Woolen Robe (2585, -0.77 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 3.1 spell_power points (0.23 DPS) | yes | Mystic's Bracelets (14366, -0.14 DPS) [world_drop]; Repurposed Hair Band (281256, -0.14 DPS) [quest]; Bright Bracers (3647, -0.52 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.52 DPS) | yes | Pristine Gloves (253913, -0.09 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Tomb Robber's Gloves (280096, -0.25 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.4 spell_power points (0.47 DPS) | yes | Keller's Girdle (2911, -0.11 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.65 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.9 spell_power points (1.02 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.51 DPS) [dungeon]; Rumpled Kilt (274741, -0.65 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.4 spell_power points (0.70 DPS) | yes | Feather Padded Treads (285345, -0.27 DPS, sim-verified) [world]; Pristine Boots (253889, -0.34 DPS) [crafted]; Red Woolen Boots (4313, -0.40 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.46 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Sludge-Stained Band (286535, -0.24 DPS) [world]; Volcanic Rock Ring (12053, -0.32 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.37 DPS) | yes | Sludge-Stained Band (286535, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.84 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 6.1 spell_power points (0.45 DPS) | yes | Channeler's Staff (4437, +0.00 DPS, sim-verified) [world]; Lesser Staff of the Spire (1300, -0.18 DPS) [world]; Staff of Westfall (2042, -0.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 306.0 spell_power points (22.53 DPS) | yes | Skycaller (12984, -0.75 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.24 DPS) [dungeon]; Deepblaze (279896, -4.00 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 58.9. Weights run: 0.6s. Verify run: 0.6s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.956 ± 0.035, crit=0.216 ± 0.017 per rating point (14 rating = 1%, 3.028 per %), hit=0.427 ± 0.005 per rating point (10 rating = 1%, 4.270 per %), spell_haste=not significant (-0.495 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.6 spell_power points (1.48 DPS) | yes | Nightsky Cowl (4039, -0.39 DPS, sim-verified) [world_drop]; Holy Shroud (2721, -0.43 DPS) [world_drop]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 spell_power points (1.21 DPS) | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.14 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.6 spell_power points (1.67 DPS) | yes | Death Speaker Mantle (6685, -0.28 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) | Lord Malathrom [world] | 7.6 spell_power points (0.73 DPS) | yes | Repairman's Cape (9605, -0.08 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.09 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.18 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.4 spell_power points (2.03 DPS) | yes | Death Speaker Robes (6682, -0.51 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.67 DPS) [dungeon]; Pristine Gown (253961, -0.73 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Nightsky Wristbands (6407, -0.31 DPS) [world_drop]; Stonecloth Bindings (14416, -0.40 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.58 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 14.5 spell_power points (1.38 DPS) | yes | Truefaith Gloves (7049, -0.47 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.65 DPS) [dungeon]; Shilly Mitts (9609, -0.71 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.9 spell_power points (1.32 DPS) | yes | Crimson Silk Belt (7055, +0.00 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (58.9 DPS) | yes | Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Abomination Skin Leggings (23173, -0.79 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.7 spell_power points (1.30 DPS) | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.66 DPS) [crafted]; Acidic Walkers (9454, -1.17 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.03 DPS) [world]; Snake Hoop (6750, -0.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, +0.00 DPS, sim-verified) [world]; Snake Hoop (6750, -0.02 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.5 spell_power points (1.00 DPS) | yes | Scorn's Focal Dagger (23168, -0.14 DPS) [dungeon]; Twisted Chanter's Staff (890, -0.22 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.27 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 353.0 spell_power points (33.49 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 105.3. Weights run: 0.7s. Verify run: 0.6s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.846 ± 0.054, crit=0.207 ± 0.021 per rating point (14 rating = 1%, 2.900 per %), hit=0.499 ± 0.008 per rating point (10 rating = 1%, 4.986 per %), spell_haste=6.506 ± 0.958, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.75 DPS) | yes | Augural Shroud (2620, -0.15 DPS, sim-verified) [world]; Corpseshroud (10574, -0.64 DPS) [dungeon]; Enchanter's Cowl (4322, -0.86 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.58 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.76 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.81 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.81 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.0 spell_power points (2.35 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest]; Green Silken Shoulders (7057, -0.65 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.6 spell_power points (2.17 DPS) | yes | Guardian Cloak (5965, -0.84 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.96 DPS) [vendor]; Long Silken Cloak (4326, -1.78 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.1 spell_power points (3.54 DPS) | yes | Robe of Power (7054, -0.38 DPS) [crafted]; Dreamweave Vest (10021, -0.63 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.80 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.18 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.18 DPS) [world_drop]; Condor Bracers (15864, -0.26 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.4 spell_power points (2.80 DPS) | yes | Black Mageweave Gloves (10003, -0.84 DPS) [crafted]; Stormcloth Gloves (10011, -0.95 DPS) [crafted]; Red Mageweave Gloves (10018, -0.97 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.7 spell_power points (2.58 DPS) | yes | Gilded Cord (254037, -0.64 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.81 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.2 spell_power points (3.16 DPS) | yes | Abomination Skin Leggings (23173, -1.10 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.12 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.36 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.14 DPS) | yes | Gilded Slippers (254001, -0.99 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.60 DPS) [dungeon]; Spidersilk Boots (4320, -1.78 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.97 DPS) | yes | Ring of Forlorn Spirits (2043, -0.93 DPS) [quest]; Reedknot Ring (9622, -1.06 DPS) [quest]; Minor Channeling Ring (1449, -1.10 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.18 DPS) | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Minor Channeling Ring (1449, -0.30 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.25 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (105.3 DPS) | yes | Windweaver Staff (7757, -0.96 DPS) [dungeon]; Staff of Jordan (873, -1.40 DPS) [world_drop]; Gut Ripper (2164, -4.29 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 306.1 spell_power points (40.04 DPS) | yes | Nether Force Wand (11263, -1.84 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.52 DPS) [quest]; Ragefire Wand (7513, -2.57 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 152.4. Weights run: 0.6s. Verify run: 0.7s. 422 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.186 ± 0.069, crit=0.346 ± 0.031 per rating point (14 rating = 1%, 4.846 per %), hit=0.739 ± 0.011 per rating point (10 rating = 1%, 7.393 per %), spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 42.7 spell_power points (5.42 DPS) | yes | Dreamweave Circlet (10041, -1.25 DPS) [crafted]; Chief Architect's Monocle (11839, -1.36 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Hat (220889, -2.04 DPS, sim-verified) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 23.1 spell_power points (2.93 DPS) | yes | Scorn's Icy Choker (23169, -1.14 DPS) [dungeon]; Mindburst Medallion (11196, -1.27 DPS) [quest]; Horizon Choker (13085, -2.58 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 34.3 spell_power points (4.36 DPS) | yes | Kentic Amice (11624, -0.23 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.21 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.37 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.1 spell_power points (2.68 DPS) | yes | Runecloth Cloak (13860, -0.33 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.57 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.41 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.7 spell_power points (5.42 DPS) | yes | Runecloth Tunic (13857, -1.61 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.63 DPS) [vendor]; Runecloth Robe (13858, -2.36 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.8 spell_power points (2.26 DPS) | yes | Nethergeld Cuffs (254061, -0.32 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.45 DPS) [quest]; Bloodband Bracers (11469, -0.93 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 36.0 spell_power points (4.57 DPS) | yes | Raider Handwraps (272098, -1.05 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.56 DPS) [vendor]; Red Mageweave Gloves (10018, -1.67 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 28.5 spell_power points (3.62 DPS) | yes | Ban'thok Sash (11662, -0.35 DPS) [dungeon]; Deathmage Sash (10771, -0.47 DPS) [dungeon]; Satyrmane Sash (17755, -0.54 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.9 spell_power points (4.42 DPS) | yes | Red Mageweave Pants (10009, -0.84 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.33 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -2.33 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 26.1 spell_power points (3.31 DPS) | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -0.56 DPS) [crafted]; Southsea Mojo Boots (20641, -0.64 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.8 spell_power points (2.26 DPS) | yes | Philanthropist's Ring (281635, -0.08 DPS) [quest]; Mindseye Circle (10634, -0.45 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.3 spell_power points (2.20 DPS) | yes | Mindseye Circle (10634, -0.39 DPS) [dungeon]; Band of the Unicorn (7553, -0.55 DPS) [world_drop]; Philanthropist's Ring (281635, -1.12 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (152.4 DPS) | yes | Uther's Strength (11302, -0.32 DPS, sim-verified) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (152.4 DPS) | yes | Spellshifter Rod (9527, -0.90 DPS) [quest]; Radiant Staff (249453, -1.50 DPS) [crafted]; Blade of Eternal Darkness (17780, -1.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -0.95 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 457.8. Weights run: 0.7s. Verify run: 0.7s. 1030 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.011, intellect=0.971 ± 0.137, crit=0.545 ± 0.040 per rating point (14 rating = 1%, 7.628 per %), hit=1.387 ± 0.025 per rating point (10 rating = 1%, 13.866 per %), spell_haste=not significant (-3.482 ± 2.444), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 86.1 spell_power points (17.88 DPS) | yes | Fireleaf Hood (240048, -5.47 DPS, sim-verified) [vendor]; Field Marshal's Coronet (231604, -6.01 DPS) [pvp] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (457.8 DPS) | yes | Beads of Ogre Mojo (22149, +0.00 DPS) [quest]; Chains of the Lich (23125, +0.00 DPS) [dungeon]; Amulet of the Dawn (22657, -5.74 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 63.2 spell_power points (13.13 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -2.91 DPS) [vendor]; Field Marshal's Silk Spaulders (16444, -4.91 DPS) [vendor]; Fireleaf Mantle (240046, -5.14 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 37.6 spell_power points (7.82 DPS) | yes | Hide of the Wild (18510, -2.89 DPS) [crafted]; Spritecaster Cape (11623, -3.70 DPS) [dungeon]; Crystalline Threaded Cape (20697, -6.72 DPS, sim-verified) [world] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 96.7 spell_power points (20.09 DPS) | yes | Fireleaf Garb (240051, -1.76 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -7.77 DPS) [crafted]; Field Marshal's Silk Vestments (16443, -8.22 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 58.2 spell_power points (12.08 DPS) | yes | Dryad's Wrist Bindings (19595, -5.90 DPS) [rep]; Fireleaf Wristwraps (240044, -6.72 DPS, sim-verified) [vendor] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 70.9 spell_power points (14.72 DPS) | yes | Fireleaf Mitts (240049, -0.76 DPS, sim-verified) [vendor]; Sorcerer's Gloves (22066, -6.52 DPS) [quest]; Sorcerer's Gauntlets (226930, -6.52 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 72.5 spell_power points (15.06 DPS) | yes | Knowledge of the Timbermaw (228190, -2.55 DPS) [vendor]; Fireleaf Waistguard (240045, -4.36 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -6.09 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 86.1 spell_power points (17.88 DPS) | yes | Fireleaf Pants (240047, -2.59 DPS, sim-verified) [vendor]; Marshal's Silk Leggings (231605, -6.03 DPS) [pvp] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 64.2 spell_power points (13.34 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (231606, -3.27 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (457.8 DPS) | yes | Channeler's Ring (272406, -4.81 DPS) [vendor]; Maiden's Circle (13001, -5.36 DPS) [world_drop]; Naglering (11669, -24.32 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (457.8 DPS) | yes | Channeler's Ring (272406, -1.08 DPS) [vendor]; Maiden's Circle (13001, -1.64 DPS) [world_drop]; Naglering (11669, -20.31 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (457.8 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (457.8 DPS) | yes | Weakness Analyzer (272438, -1.45 DPS) [vendor]; Serenity Field (272439, -3.12 DPS) [vendor]; Blackhand's Breadth (13965, -6.92 DPS, sim-verified) [quest] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (457.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.36 DPS) [world]; Teebu's Blazing Longsword (1728, -18.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 375.6 spell_power points (78.01 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -3.79 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.28 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.82 DPS) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Briarwood Reed; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1030, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 30.1. Weights run: 0.6s. Verify run: 0.6s. 141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.612 ± 0.017, crit=0.136 ± 0.007 per rating point (14 rating = 1%, 1.900 per %), hit=0.395 ± 0.003 per rating point (10 rating = 1%, 3.948 per %), spell_haste=not significant (-0.590 ± 0.325), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.44 DPS) | yes | Shadow Goggles (4373, -0.48 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.5 spell_power points (0.77 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.36 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.29 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Black Whelp Cloak (7283, -0.07 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.1 spell_power points (0.59 DPS) | yes | Green Woolen Robe (6243, -0.24 DPS) [crafted]; Mystic's Wrap (14369, -0.28 DPS) [world_drop]; Gray Woolen Robe (2585, -0.63 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.7 spell_power points (0.27 DPS) | yes | Mindthrust Bracers (1974, -0.05 DPS) [dungeon]; Bright Bracers (3647, -0.09 DPS) [world_drop]; Featherbead Bracers (15452, -0.18 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.52 DPS) | yes | Pristine Gloves (253913, -0.09 DPS) [crafted]; Gnoll Casting Gloves (892, -0.09 DPS, sim-verified) [world]; Blight Gloves (279877, -0.20 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.4 spell_power points (0.47 DPS) | yes | Keller's Girdle (2911, -0.11 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.50 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (30.1 DPS) | yes | Silk-threaded Trousers (1929, -0.20 DPS) [dungeon]; Rumpled Kilt (274741, -0.34 DPS) [vendor]; Abomination Skin Leggings (23173, -0.39 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.4 spell_power points (0.70 DPS) | yes | Pristine Boots (253889, -0.34 DPS) [crafted]; Red Woolen Boots (4313, -0.40 DPS) [crafted]; Feather Padded Treads (285345, -0.44 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.37 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Sludge-Stained Band (286535, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 3.7 spell_power points (0.27 DPS) | yes | Sludge-Stained Band (286535, -0.05 DPS) [world]; Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; Loop of Sacrifice (281673, -0.18 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 6.1 spell_power points (0.45 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.09 DPS) [world]; Lesser Staff of the Spire (1300, -0.18 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 306.0 spell_power points (22.53 DPS) | yes | Skycaller (12984, -0.61 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.24 DPS) [dungeon]; Sizzle Stick (8071, -4.51 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (orc, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 54.1. Weights run: 0.6s. Verify run: 0.6s. 236 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.956 ± 0.035, crit=0.216 ± 0.017 per rating point (14 rating = 1%, 3.028 per %), hit=0.427 ± 0.005 per rating point (10 rating = 1%, 4.270 per %), spell_haste=not significant (-0.495 ± 0.543), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.6 spell_power points (1.48 DPS) | yes | Nightsky Cowl (4039, -0.06 DPS, sim-verified) [world_drop]; Holy Shroud (2721, -0.43 DPS) [world_drop]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 spell_power points (1.21 DPS) | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.99 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.6 spell_power points (1.67 DPS) | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.6 spell_power points (0.73 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.18 DPS) [world_drop]; Hillman's Cloak (3719, -0.25 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.4 spell_power points (2.03 DPS) | yes | Death Speaker Robes (6682, -0.43 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.67 DPS) [dungeon]; Pristine Gown (253961, -0.73 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.85 DPS) | yes | Nightsky Wristbands (6407, -0.31 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.31 DPS) [quest]; Glowing Magical Bracelets (13106, -0.49 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.8 spell_power points (1.02 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.30 DPS) [dungeon]; Serpent Gloves (5970, -0.36 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.9 spell_power points (1.32 DPS) | yes | Crimson Silk Belt (7055, +0.00 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (54.1 DPS) | yes | Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Abomination Skin Leggings (23173, -0.90 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.7 spell_power points (1.30 DPS) | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.66 DPS) [crafted]; Acidic Walkers (9454, -0.75 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.66 DPS) | yes | Snake Hoop (6750, -0.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 6.7 spell_power points (0.64 DPS) | yes | Snake Hoop (6750, +0.00 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.07 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.5 spell_power points (1.00 DPS) | yes | Twisted Chanter's Staff (890, -0.09 DPS) [world_drop]; Scorn's Focal Dagger (23168, -0.14 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.16 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 353.0 spell_power points (33.49 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Black Widow Band; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 236, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 96.3. Weights run: 0.7s. Verify run: 0.6s. 318 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.846 ± 0.054, crit=0.207 ± 0.021 per rating point (14 rating = 1%, 2.900 per %), hit=0.499 ± 0.008 per rating point (10 rating = 1%, 4.986 per %), spell_haste=6.506 ± 0.958, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.75 DPS) | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Corpseshroud (10574, -0.64 DPS) [dungeon]; Enchanter's Cowl (4322, -0.86 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.58 DPS) | yes | Triune Amulet (7722, -0.81 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.81 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -0.84 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.0 spell_power points (2.35 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Green Silken Shoulders (7057, -0.30 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.6 spell_power points (2.17 DPS) | yes | Guardian Cloak (5965, -0.84 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.96 DPS) [vendor]; Long Silken Cloak (4326, -1.51 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.1 spell_power points (3.54 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.38 DPS) [crafted]; Elemental Raiment (9434, -0.80 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.8 spell_power points (1.41 DPS) | yes | Spidertank Oilrag (9448, -0.14 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.41 DPS) [world_drop]; Condor Bracers (15864, -0.49 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.4 spell_power points (2.80 DPS) | yes | Black Mageweave Gloves (10003, -0.84 DPS) [crafted]; Red Mageweave Gloves (10018, -0.89 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.95 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.7 spell_power points (2.58 DPS) | yes | Gilded Cord (254037, -0.64 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.81 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.2 spell_power points (3.16 DPS) | yes | Abomination Skin Leggings (23173, -1.10 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.19 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.36 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.14 DPS) | yes | Gilded Slippers (254001, -0.50 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.60 DPS) [dungeon]; Spidersilk Boots (4320, -1.78 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.97 DPS) | yes | Reedknot Ring (9622, -1.06 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.19 DPS) [vendor]; Black Widow Band (6199, -1.20 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.18 DPS) | yes | Sea Giant's Toe Ring (274746, -0.39 DPS) [vendor]; Black Widow Band (6199, -0.40 DPS) [world]; Reedknot Ring (9622, -0.97 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (96.3 DPS) | yes | Windweaver Staff (7757, -0.96 DPS) [dungeon]; Staff of Jordan (873, -1.40 DPS) [world_drop]; Gut Ripper (2164, -3.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 306.1 spell_power points (40.04 DPS) | yes | Nether Force Wand (11263, -1.50 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.52 DPS) [quest]; Ragefire Wand (7513, -2.57 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 138.9. Weights run: 0.6s. Verify run: 0.7s. 412 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.186 ± 0.069, crit=0.346 ± 0.031 per rating point (14 rating = 1%, 4.846 per %), hit=0.739 ± 0.011 per rating point (10 rating = 1%, 7.393 per %), spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 42.7 spell_power points (5.42 DPS) | yes | Dreamweave Circlet (10041, -1.25 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.30 DPS, sim-verified) [vendor]; Chief Architect's Monocle (11839, -1.36 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 23.1 spell_power points (2.93 DPS) | yes | Scorn's Icy Choker (23169, -1.14 DPS) [dungeon]; Mindburst Medallion (11196, -1.27 DPS) [quest]; Horizon Choker (13085, -1.59 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 34.3 spell_power points (4.36 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -1.21 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.37 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.7 spell_power points (2.88 DPS) | yes | Spritecaster Cape (11623, -0.08 DPS, sim-verified) [dungeon]; Mantle of Lady Falther'ess (23178, -0.38 DPS) [dungeon]; Runecloth Cloak (13860, -0.53 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.7 spell_power points (5.42 DPS) | yes | Runecloth Tunic (13857, -1.61 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.63 DPS) [vendor]; Runecloth Robe (13858, -2.09 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.8 spell_power points (2.26 DPS) | yes | Nethergeld Cuffs (254061, -0.32 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.45 DPS) [quest]; Bloodband Bracers (11469, -0.48 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 36.0 spell_power points (4.57 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.56 DPS) [vendor]; Red Mageweave Gloves (10018, -1.67 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 28.5 spell_power points (3.62 DPS) | yes | Ban'thok Sash (11662, -0.35 DPS) [dungeon]; Satyrmane Sash (17755, -0.45 DPS, sim-verified) [dungeon]; Deathmage Sash (10771, -0.47 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.9 spell_power points (4.42 DPS) | yes | Red Mageweave Pants (10009, -0.84 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -0.91 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.33 DPS) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 26.1 spell_power points (3.31 DPS) | yes | Gilded Sandals (254107, -0.56 DPS) [crafted]; Southsea Mojo Boots (20641, -0.64 DPS) [quest]; Earthen Silk Slippers (254013, -0.68 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.8 spell_power points (2.26 DPS) | yes | Philanthropist's Ring (281635, -0.08 DPS) [quest]; Mindseye Circle (10634, -0.45 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.3 spell_power points (2.20 DPS) | yes | Mindseye Circle (10634, -0.39 DPS) [dungeon]; Band of the Unicorn (7553, -0.55 DPS) [world_drop]; Philanthropist's Ring (281635, -0.90 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (138.9 DPS) | yes | Uther's Strength (11302, -0.08 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (138.9 DPS) | yes | Uther's Strength (11302, -0.06 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (138.9 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.90 DPS) [quest]; Radiant Staff (249453, -1.50 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 412, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 430.4. Weights run: 0.7s. Verify run: 0.7s. 1021 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.011, intellect=0.971 ± 0.137, crit=0.545 ± 0.040 per rating point (14 rating = 1%, 7.628 per %), hit=1.387 ± 0.025 per rating point (10 rating = 1%, 13.866 per %), spell_haste=not significant (-3.482 ± 2.444), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 86.1 spell_power points (17.88 DPS) | yes | Fireleaf Hood (240048, -3.67 DPS, sim-verified) [vendor]; Warlord's Silk Cowl (231601, -6.01 DPS) [pvp] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (430.4 DPS) | yes | Beads of Ogre Mojo (22149, -0.62 DPS) [quest]; Chains of the Lich (23125, -1.17 DPS) [dungeon]; Jewel of Kajaro (19601, -3.37 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 63.2 spell_power points (13.13 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -2.91 DPS) [vendor]; Fireleaf Mantle (240046, -3.36 DPS, sim-verified) [vendor]; Warlord's Silk Amice (16536, -4.91 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 37.6 spell_power points (7.82 DPS) | yes | Hide of the Wild (18510, -2.89 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.51 DPS) [quest]; Crystalline Threaded Cape (20697, -12.34 DPS, sim-verified) [world] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 96.7 spell_power points (20.09 DPS) | yes | Fireleaf Garb (240051, -1.55 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -7.77 DPS) [crafted]; Warlord's Silk Raiment (16535, -8.22 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 58.2 spell_power points (12.08 DPS) | yes | Dryad's Wrist Bindings (19595, -5.90 DPS) [rep]; Fireleaf Wristwraps (240044, -11.61 DPS, sim-verified) [vendor] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 70.9 spell_power points (14.72 DPS) | yes | Fireleaf Mitts (240049, -0.63 DPS, sim-verified) [vendor]; Sorcerer's Gloves (22066, -6.52 DPS) [quest]; Sorcerer's Gauntlets (226930, -6.52 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 72.5 spell_power points (15.06 DPS) | yes | Fireleaf Waistguard (240045, -0.55 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -2.55 DPS) [vendor]; Belt of the Archmage (18405, -6.09 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 86.1 spell_power points (17.88 DPS) | yes | Fireleaf Pants (240047, -4.00 DPS, sim-verified) [vendor]; General's Silk Trousers (231595, -6.03 DPS) [pvp] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 64.2 spell_power points (13.34 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (231597, -3.27 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (430.4 DPS) | yes | Channeler's Ring (272406, -4.81 DPS) [vendor]; Maiden's Circle (13001, -5.36 DPS) [world_drop]; Naglering (11669, -21.02 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (430.4 DPS) | yes | Channeler's Ring (272406, -1.08 DPS) [vendor]; Maiden's Circle (13001, -1.64 DPS) [world_drop]; Naglering (11669, -9.94 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (430.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (430.4 DPS) | yes | Weakness Analyzer (272438, -1.45 DPS) [vendor]; Serenity Field (272439, -3.12 DPS) [vendor]; Blackhand's Breadth (13965, -5.52 DPS, sim-verified) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (430.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.34 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -20.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 375.6 spell_power points (78.01 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -2.13 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -12.28 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.82 DPS) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Amulet of the Dawn; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1021, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

