# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 37.5. Weights run: 0.7s. Verify run: 0.5s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.654 ± 0.006, crit=0.047 ± 0.003 per rating point (14 rating = 1%, 0.654 per %), hit=0.113 ± 0.002 per rating point (10 rating = 1%, 1.133 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -1.42 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (1.44 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.40 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.91 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Caretaker's Cape (20428, -0.13 DPS) [rep]; Pearl-clasped Cloak (5542, -0.19 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.3 spell_power points (1.09 DPS) | yes | Green Woolen Robe (6243, -0.44 DPS) [crafted]; Mystic's Wrap (14369, -0.49 DPS) [world_drop]; Gray Woolen Robe (2585, -1.02 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 3.3 spell_power points (0.43 DPS) | yes | Mystic's Bracelets (14366, -0.26 DPS) [world_drop]; Repurposed Hair Band (281256, -0.26 DPS) [quest]; Bright Bracers (3647, -0.59 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Pristine Gloves (253913, -0.14 DPS) [crafted]; Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Tomb Robber's Gloves (280096, -0.41 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.6 spell_power points (0.87 DPS) | yes | Keller's Girdle (2911, -0.18 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.35 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.60 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.88 DPS) | yes | Filigreed Pristine Leggings (253937, +0.28 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.95 DPS) [dungeon]; Rumpled Kilt (274741, -1.22 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (1.27 DPS) | yes | Pristine Boots (253889, -0.61 DPS) [crafted]; Red Woolen Boots (4313, -0.74 DPS) [crafted]; Feather Padded Treads (285345, -0.74 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.3 spell_power points (0.83 DPS) | yes | Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Sludge-Stained Band (286535, -0.44 DPS) [world]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.66 DPS) | yes | Sludge-Stained Band (286535, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.31 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 6.5 spell_power points (0.86 DPS) | yes | Lesser Staff of the Spire (1300, -0.35 DPS) [world_drop]; Channeler's Staff (4437, -0.35 DPS, sim-verified) [world]; Staff of Westfall (2042, -0.43 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 171.9 spell_power points (22.71 DPS) | yes | Skycaller (12984, -1.07 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 62.4. Weights run: 0.7s. Verify run: 0.5s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.638 ± 0.009, crit=0.119 ± 0.008 per rating point (14 rating = 1%, 1.663 per %), hit=0.165 ± 0.002 per rating point (10 rating = 1%, 1.646 per %), spell_haste=0.154 ± 0.033, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.91 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.52 DPS) [crafted]; Embalmed Shroud (7691, -0.68 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.67 DPS) | yes | Crystal Starfire Medallion (5003, -1.28 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.28 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.53 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (2.28 DPS) | yes | Death Speaker Mantle (6685, -0.44 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Magician's Mantle (12998, -0.62 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.6 spell_power points (0.86 DPS) | yes | Darkspear Raider's Cloak (272078, -0.07 DPS) [vendor]; Hillman's Cloak (3719, -0.09 DPS) [crafted]; Cloak of Rot (4462, -0.13 DPS, sim-verified) [world] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (2.67 DPS) | yes | Tree Bark Jacket (1486, -0.66 DPS) [dungeon]; Death Speaker Robes (6682, -0.85 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.90 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Nightsky Wristbands (6407, -0.80 DPS) [world_drop]; Stonecloth Bindings (14416, -0.90 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.25 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 11.0 spell_power points (1.70 DPS) | yes | Shilly Mitts (9609, -0.40 DPS, sim-verified) [quest]; Serpent Gloves (5970, -0.62 DPS) [dungeon]; Truefaith Gloves (7049, -0.63 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.9 spell_power points (1.99 DPS) | yes | Belt of Arugal (6392, -0.07 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.38 DPS) [crafted]; Invoker's Cord (215366, -0.42 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (2.18 DPS) | yes | Gaze Dreamer Pants (6903, -0.10 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.41 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.77 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Nimbus Boots (6998, -0.84 DPS) [quest]; Acidic Walkers (9454, -1.50 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.08 DPS) | yes | Lorekeeper's Ring (20431, -0.31 DPS) [rep]; Snake Hoop (6750, -0.39 DPS) [quest]; Minor Channeling Ring (1449, -1.84 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (62.4 DPS) | yes | Black Widow Band (6199, -0.24 DPS) [world]; Snake Hoop (6750, -0.24 DPS) [quest]; Minor Channeling Ring (1449, -0.96 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Twisted Chanter's Staff (890, -0.40 DPS) [world_drop]; Channeler's Staff (4437, -0.60 DPS) [world]; Glimmering Staff (249392, -0.71 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.8 spell_power points (1.67 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.97 DPS) [vendor]; Eye of Paleth (2943, -1.05 DPS) [quest]; Dwarven Tome (279898, -1.26 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 218.2 spell_power points (33.67 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.90 DPS) [crafted]; Greater Mystic Wand (11290, -4.67 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 92.2. Weights run: 0.7s. Verify run: 0.5s. 329 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.863 ± 0.018, crit=0.197 ± 0.013 per rating point (14 rating = 1%, 2.753 per %), hit=0.299 ± 0.005 per rating point (10 rating = 1%, 2.986 per %), spell_haste=0.902 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.14 DPS) | yes | Corpseshroud (10574, -0.69 DPS) [dungeon]; Augural Shroud (2620, -0.87 DPS, sim-verified) [world]; Miner's Hat of the Deep (9429, -0.95 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.2 spell_power points (1.82 DPS) | yes | Necklace of Calisea (1714, -0.92 DPS) [world_drop]; Triune Amulet (7722, -0.92 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.70 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.2 spell_power points (2.73 DPS) | yes | Bloodmage Mantle (7684, -0.22 DPS) [dungeon]; Green Silken Shoulders (7057, -0.23 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.39 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (92.2 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.12 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.96 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.2 spell_power points (4.07 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.42 DPS) [crafted]; Elemental Raiment (9434, -0.92 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.35 DPS) | yes | Spidertank Oilrag (9448, +0.65 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.18 DPS) [world_drop]; Condor Bracers (15864, -0.30 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.5 spell_power points (3.21 DPS) | yes | Red Mageweave Gloves (10018, -0.60 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.97 DPS) [crafted]; Stormcloth Gloves (10011, -1.06 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.9 spell_power points (2.98 DPS) | yes | Gilded Cord (254037, -0.75 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.95 DPS) [rep]; Highlander's Cloth Girdle (20098, -1.15 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.4 spell_power points (3.64 DPS) | yes | Crimson Silk Pantaloons (7062, -0.89 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.26 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.56 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | Gilded Slippers (254001, -1.14 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.81 DPS) [dungeon]; Spidersilk Boots (4320, -2.03 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.2 spell_power points (2.27 DPS) | yes | Ring of Forlorn Spirits (2043, -1.07 DPS) [quest]; Reedknot Ring (9622, -1.22 DPS) [quest]; Minor Channeling Ring (1449, -1.26 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.35 DPS) | yes | Reedknot Ring (9622, -0.30 DPS) [quest]; Lorekeeper's Ring (19525, -0.30 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.31 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -1.06 DPS) [dungeon]; Staff of Jordan (873, -1.57 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 268.3 spell_power points (40.15 DPS) | yes | Nether Force Wand (11263, -1.79 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.58 DPS) [quest]; Ragefire Wand (7513, -2.63 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 329, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 121.5. Weights run: 0.5s. Verify run: 0.6s. 426 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.911 ± 0.050, crit=0.368 ± 0.021 per rating point (14 rating = 1%, 5.153 per %), hit=0.574 ± 0.009 per rating point (10 rating = 1%, 5.744 per %), spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.2 spell_power points (4.88 DPS) | yes | Dreamweave Circlet (10041, +0.08 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -0.94 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.34 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.04 DPS) [dungeon]; Mindburst Medallion (11196, -0.17 DPS) [quest]; Arcane Crystal Pendant (20037, -2.63 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.4 spell_power points (3.86 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -1.06 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted]; Kentic Amice (11624, -4.33 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.5 spell_power points (2.55 DPS) | yes | Runecloth Cloak (13860, -0.42 DPS) [crafted]; Big Voodoo Cloak (8216, -0.82 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -3.12 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.2 spell_power points (4.88 DPS) | yes | Robe of the Magi (1716, -0.91 DPS, sim-verified) [world_drop]; Knight's Dreadweave Vest (220886, -1.32 DPS) [vendor]; Runecloth Tunic (13857, -1.34 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodband Bracers (11469, -0.02 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.32 DPS) [quest]; Aristocratic Cuffs (12546, -4.10 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.5 spell_power points (4.00 DPS) | yes | Raider Handwraps (272098, +0.51 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.16 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.22 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.3 spell_power points (3.06 DPS) | yes | Satyrmane Sash (17755, +0.44 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.09 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.41 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.08 DPS) [crafted]; Spellshock Leggings (9484, -5.20 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.15 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.11 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -0.63 DPS) [crafted]; Southsea Mojo Boots (20641, -0.78 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.03 DPS) | yes | Brainlash (6440, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.45 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.02 DPS) | yes | Brainlash (6440, +0.00 DPS, sim-verified) [dungeon]; Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.44 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Blessed Prayer Beads (19990, -1.00 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Spellshifter Rod (9527, -0.72 DPS) [quest]; Spellforce Rod (1664, -0.84 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -11.85 DPS, sim-verified) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 256.4. Weights run: 0.8s. Verify run: 0.6s. 1020 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.272 ± 0.036, crit=0.398 ± 0.029 per rating point (14 rating = 1%, 5.578 per %), hit=0.648 ± 0.012 per rating point (10 rating = 1%, 6.478 per %), spell_haste=not significant (0.090 ± 0.410), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 89.7 spell_power points (11.05 DPS) | yes | Fireleaf Hood (240048, -1.13 DPS, sim-verified) [vendor]; Field Marshal's Coronet (16441, -3.64 DPS) [vendor]; Field Marshal's Coronet (231604, -3.64 DPS) [pvp] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (256.4 DPS) | yes | Jewel of Kajaro (19601, -0.18 DPS, sim-verified) [quest]; Beads of Ogre Mojo (22149, -0.40 DPS) [quest]; Pebble of Kajaro (19600, -0.74 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 65.7 spell_power points (8.09 DPS) | yes | Fireleaf Mantle (240046, -0.94 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -1.72 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.38 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.7 spell_power points (4.02 DPS) | yes | Crystalline Threaded Cape (20697, -0.93 DPS) [world_drop]; Spritecaster Cape (11623, -1.36 DPS) [dungeon]; Hide of the Wild (18510, -1.46 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 98.3 spell_power points (12.11 DPS) | yes | Fireleaf Garb (240051, -0.85 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -4.62 DPS) [crafted]; Field Marshal's Silk Vestments (16443, -4.70 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 60.9 spell_power points (7.50 DPS) | yes | Fireleaf Wristwraps (240044, -1.91 DPS, sim-verified) [vendor]; Dryad's Wrist Bindings (19595, -3.54 DPS) [rep]; Arcanist Bindings (16799, -3.68 DPS) [world_drop] |
| hands | Fireleaf Mitts (240049) | Leonid Barthalomew the Revered [vendor] | 71.7 spell_power points (8.83 DPS) | yes | Fireleaf Gloves (240057, +1.40 DPS, sim-verified) [vendor]; Raider Handwraps (272097, -3.61 DPS) [vendor]; Marshal's Silk Gloves (16440, -3.62 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 69.0 spell_power points (8.50 DPS) | yes | Knowledge of the Timbermaw (228190, -1.18 DPS) [vendor]; Magician's Cord (272393, -2.40 DPS) [vendor]; Fireleaf Waistguard (240045, -2.47 DPS, sim-verified) [vendor] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 89.7 spell_power points (11.05 DPS) | yes | Fireleaf Pants (240047, -1.77 DPS, sim-verified) [vendor]; Marshal's Silk Leggings (16442, -3.54 DPS) [vendor]; Marshal's Silk Leggings (231605, -3.54 DPS) [pvp] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 66.4 spell_power points (8.18 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (16437, -2.60 DPS) [vendor]; Marshal's Silk Footwraps (231606, -2.60 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (256.4 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.40 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.53 DPS) [vendor]; Naglering (11669, -11.14 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (256.4 DPS) | yes | Ritssyn's Ring of Chaos (21836, -0.67 DPS) [world_drop]; Cauterizing Band (19140, -0.71 DPS) [world_drop]; Naglering (11669, -6.90 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (256.4 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (256.4 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, -0.25 DPS) [dungeon]; Weakness Analyzer (272438, -1.14 DPS, sim-verified) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (256.4 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Teebu's Blazing Longsword (1728, -13.37 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 629.2 spell_power points (77.50 DPS) | yes | Bonecreeper Stylus (13938, -12.89 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.14 DPS) [world]; Ritssyn's Wand of Bad Mojo (22408, -21.92 DPS, sim-verified) [dungeon] |

**New at 60:** head: Fireleaf Circlet; neck: Amulet of the Dawn; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Mitts; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1020, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 34.2. Weights run: 0.7s. Verify run: 0.5s. 141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.654 ± 0.006, crit=0.047 ± 0.003 per rating point (14 rating = 1%, 0.654 per %), hit=0.113 ± 0.002 per rating point (10 rating = 1%, 1.133 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -1.31 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (1.44 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.45 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.91 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Pearl-clasped Cloak (5542, -0.10 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.13 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.3 spell_power points (1.09 DPS) | yes | Green Woolen Robe (6243, -0.44 DPS) [crafted]; Mystic's Wrap (14369, -0.49 DPS) [world_drop]; Gray Woolen Robe (2585, -0.89 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.9 spell_power points (0.52 DPS) | yes | Mindthrust Bracers (1974, +0.25 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.09 DPS) [quest]; Bright Bracers (3647, -0.17 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Pristine Gloves (253913, -0.14 DPS) [crafted]; Gnoll Casting Gloves (892, -0.15 DPS, sim-verified) [world]; Blight Gloves (279877, -0.32 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.6 spell_power points (0.87 DPS) | yes | Keller's Girdle (2911, -0.18 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.35 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.63 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (34.2 DPS) | yes | Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.43 DPS, sim-verified) [dungeon]; Rumpled Kilt (274741, -0.65 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (1.27 DPS) | yes | Pristine Boots (253889, -0.61 DPS) [crafted]; Red Woolen Boots (4313, -0.74 DPS) [crafted]; Feather Padded Treads (285345, -0.74 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.23 DPS) [quest]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 3.9 spell_power points (0.52 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -0.12 DPS) [world]; Volcanic Rock Ring (12053, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 6.5 spell_power points (0.86 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.17 DPS) [world]; Lesser Staff of the Spire (1300, -0.35 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 171.9 spell_power points (22.71 DPS) | yes | Skycaller (12984, -1.15 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Sizzle Stick (8071, -4.40 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 56.9. Weights run: 0.7s. Verify run: 0.5s. 236 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.638 ± 0.009, crit=0.119 ± 0.008 per rating point (14 rating = 1%, 1.663 per %), hit=0.165 ± 0.002 per rating point (10 rating = 1%, 1.646 per %), spell_haste=0.154 ± 0.033, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.91 DPS) | yes | Silk Headband (7050, -0.52 DPS) [crafted]; Holy Shroud (2721, -0.57 DPS, sim-verified) [world_drop]; Embalmed Shroud (7691, -0.68 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.67 DPS) | yes | Crystal Starfire Medallion (5003, -1.28 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.28 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.50 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (2.28 DPS) | yes | Death Speaker Mantle (6685, -0.44 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Magician's Mantle (12998, -0.62 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 5.1 spell_power points (0.79 DPS) | yes | Darkspear Raider's Cloak (272078, +0.09 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.02 DPS) [crafted]; Windsong Drape (15468, -0.02 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (2.67 DPS) | yes | Tree Bark Jacket (1486, -0.66 DPS) [dungeon]; Death Speaker Robes (6682, -0.81 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.90 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Nightsky Wristbands (6407, -0.80 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.80 DPS) [quest]; Glowing Magical Bracelets (13106, -1.59 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.2 spell_power points (1.42 DPS) | yes | Truefaith Gloves (7049, -0.35 DPS) [crafted]; Gnoll Casting Gloves (892, -0.49 DPS) [world]; Serpent Gloves (5970, -0.98 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.9 spell_power points (1.99 DPS) | yes | Belt of Arugal (6392, -0.31 DPS) [dungeon]; Crimson Silk Belt (7055, -0.38 DPS) [crafted]; Warsong Sash (16975, -0.72 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (2.18 DPS) | yes | Pristine Leggings (253987, -0.41 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.48 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.77 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Boots of the Enchanter (4325, -1.00 DPS) [crafted]; Acidic Walkers (9454, -1.55 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.08 DPS) | yes | Advisor's Ring (20426, -0.31 DPS) [rep]; Black Widow Band (6199, -0.39 DPS) [world]; Snake Hoop (6750, -0.39 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.93 DPS) | yes | Black Widow Band (6199, -0.24 DPS) [world]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Snake Hoop (6750, -1.22 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Twisted Chanter's Staff (890, -0.40 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.40 DPS) [quest]; Glimmering Staff (249392, -0.93 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.8 spell_power points (1.67 DPS) | yes | Witch's Finger (16887, -0.98 DPS) [quest]; Orb of Souls (249395, -1.05 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -1.34 DPS, sim-verified) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 218.2 spell_power points (33.67 DPS) | yes | Starfaller (13063, -0.43 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.90 DPS) [crafted]; Greater Mystic Wand (11290, -4.67 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 236, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 83.8. Weights run: 0.7s. Verify run: 0.5s. 319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.863 ± 0.018, crit=0.197 ± 0.013 per rating point (14 rating = 1%, 2.753 per %), hit=0.299 ± 0.005 per rating point (10 rating = 1%, 2.986 per %), spell_haste=0.902 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.14 DPS) | yes | Augural Shroud (2620, +0.41 DPS, sim-verified) [world]; Corpseshroud (10574, -0.69 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.95 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.2 spell_power points (1.82 DPS) | yes | Necklace of Calisea (1714, -0.92 DPS) [world_drop]; Triune Amulet (7722, -0.92 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.45 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.2 spell_power points (2.73 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.22 DPS) [dungeon]; Berylline Pads (4197, -0.39 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (83.8 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.12 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -2.55 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.2 spell_power points (4.07 DPS) | yes | Dreamweave Vest (10021, +0.64 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.42 DPS) [crafted]; Elemental Raiment (9434, -0.92 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.9 spell_power points (1.63 DPS) | yes | Spidertank Oilrag (9448, -0.09 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.47 DPS) [world_drop]; Condor Bracers (15864, -0.58 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.5 spell_power points (3.21 DPS) | yes | Red Mageweave Gloves (10018, +0.23 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.97 DPS) [crafted]; Stormcloth Gloves (10011, -1.06 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.9 spell_power points (2.98 DPS) | yes | Defiler's Cloth Girdle (20166, -0.73 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.75 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.95 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.4 spell_power points (3.64 DPS) | yes | Crimson Silk Pantaloons (7062, -0.93 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.26 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.56 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | Gilded Slippers (254001, -0.40 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.81 DPS) [dungeon]; Spidersilk Boots (4320, -2.03 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.2 spell_power points (2.27 DPS) | yes | Reedknot Ring (9622, -1.22 DPS) [quest]; Ogremind Ring (1993, -1.37 DPS) [world_drop]; Voodoo Band (1996, -1.37 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.35 DPS) | yes | Advisor's Ring (19521, -0.30 DPS) [rep]; Voodoo Band (1996, -0.44 DPS) [world_drop]; Reedknot Ring (9622, -0.99 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -1.06 DPS) [dungeon]; Staff of Jordan (873, -1.57 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 268.3 spell_power points (40.15 DPS) | yes | Nether Force Wand (11263, -1.44 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.58 DPS) [quest]; Ragefire Wand (7513, -2.63 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 108.7. Weights run: 0.5s. Verify run: 0.6s. 416 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.911 ± 0.050, crit=0.368 ± 0.021 per rating point (14 rating = 1%, 5.153 per %), hit=0.574 ± 0.009 per rating point (10 rating = 1%, 5.744 per %), spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.2 spell_power points (4.88 DPS) | yes | Blood Guard's Dreadweave Hat (220907, -0.94 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.34 DPS) [crafted]; Dreamweave Circlet (10041, -1.97 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.04 DPS) [dungeon]; Mindburst Medallion (11196, -0.17 DPS) [quest]; Arcane Crystal Pendant (20037, -1.90 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.4 spell_power points (3.86 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -1.06 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted]; Kentic Amice (11624, -4.99 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.2 spell_power points (2.65 DPS) | yes | Spritecaster Cape (11623, +0.41 DPS, sim-verified) [dungeon]; Mantle of Lady Falther'ess (23178, -0.39 DPS) [dungeon]; Runecloth Cloak (13860, -0.51 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.2 spell_power points (4.88 DPS) | yes | Stone Guard's Dreadweave Vest (220904, -1.32 DPS) [vendor]; Runecloth Tunic (13857, -1.34 DPS) [crafted]; Robe of the Magi (1716, -2.31 DPS, sim-verified) [world_drop] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodband Bracers (11469, -0.02 DPS) [quest]; Radiant Silver Bracers (4545, -0.27 DPS) [quest]; Aristocratic Cuffs (12546, -2.96 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.5 spell_power points (4.00 DPS) | yes | Raider Handwraps (272098, -0.25 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.16 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.22 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.3 spell_power points (3.06 DPS) | yes | Ban'thok Sash (11662, -0.09 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon]; Satyrmane Sash (17755, -1.47 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.41 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.08 DPS) [crafted]; Spellshock Leggings (9484, -3.73 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.15 DPS) | yes | Gilded Sandals (254107, -0.63 DPS) [crafted]; Southsea Mojo Boots (20641, -0.78 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -0.99 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.03 DPS) | yes | Brainlash (6440, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Advisor's Ring (19519, -0.45 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.02 DPS) | yes | Brainlash (6440, +0.00 DPS, sim-verified) [dungeon]; Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Advisor's Ring (19519, -0.44 DPS) [rep] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.08 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.09 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Spellshifter Rod (9527, -0.72 DPS) [quest]; Spellforce Rod (1664, -0.84 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -14.13 DPS, sim-verified) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; finger2: Cyclopean Band; trinket1: Uther's Strength; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 416, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 216.9. Weights run: 0.8s. Verify run: 0.6s. 1011 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.272 ± 0.036, crit=0.398 ± 0.029 per rating point (14 rating = 1%, 5.578 per %), hit=0.648 ± 0.012 per rating point (10 rating = 1%, 6.478 per %), spell_haste=not significant (0.090 ± 0.410), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 89.7 spell_power points (11.05 DPS) | yes | Fireleaf Hood (240048, -1.88 DPS, sim-verified) [vendor]; Warlord's Silk Cowl (16533, -3.64 DPS) [vendor]; Warlord's Silk Cowl (231601, -3.64 DPS) [pvp] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (216.9 DPS) | yes | Jewel of Kajaro (19601, +0.22 DPS, sim-verified) [quest]; Beads of Ogre Mojo (22149, -0.40 DPS) [quest]; Pebble of Kajaro (19600, -0.74 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 65.7 spell_power points (8.09 DPS) | yes | Fireleaf Mantle (240046, -1.72 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -1.72 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.38 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.7 spell_power points (4.02 DPS) | yes | Hide of the Wild (18510, +0.87 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -0.93 DPS) [world_drop]; Deep Woodlands Cloak (19121, -1.13 DPS) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 98.3 spell_power points (12.11 DPS) | yes | Fireleaf Garb (240051, -0.30 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -4.62 DPS) [crafted]; Warlord's Silk Raiment (16535, -4.70 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 60.9 spell_power points (7.50 DPS) | yes | Fireleaf Wristwraps (240044, -2.01 DPS, sim-verified) [vendor]; Dryad's Wrist Bindings (19595, -3.54 DPS) [rep]; Arcanist Bindings (16799, -3.68 DPS) [world_drop] |
| hands | Fireleaf Mitts (240049) | Leonid Barthalomew the Revered [vendor] | 71.7 spell_power points (8.83 DPS) | yes | Fireleaf Gloves (240057, +0.20 DPS, sim-verified) [vendor]; Raider Handwraps (272097, -3.61 DPS) [vendor]; General's Silk Handguards (16540, -3.62 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 69.0 spell_power points (8.50 DPS) | yes | Fireleaf Waistguard (240045, -0.64 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -1.18 DPS) [vendor]; Magician's Cord (272393, -2.40 DPS) [vendor] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 89.7 spell_power points (11.05 DPS) | yes | Fireleaf Pants (240047, -1.77 DPS, sim-verified) [vendor]; General's Silk Trousers (16534, -3.54 DPS) [vendor]; General's Silk Trousers (231595, -3.54 DPS) [pvp] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 66.4 spell_power points (8.18 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (16539, -2.60 DPS) [vendor]; General's Silk Boots (231597, -2.60 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (216.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.40 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.53 DPS) [vendor]; Naglering (11669, -7.99 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (216.9 DPS) | yes | Ritssyn's Ring of Chaos (21836, -0.67 DPS) [world_drop]; Cauterizing Band (19140, -0.71 DPS) [world_drop]; Naglering (11669, -4.82 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (216.9 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (216.9 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, -0.25 DPS) [dungeon]; Weakness Analyzer (272438, -0.95 DPS, sim-verified) [vendor] |
| main_hand | Trindlehaven Staff (13161) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-verified (216.9 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS, sim-verified) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.04 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 629.2 spell_power points (77.50 DPS) | yes | Bonecreeper Stylus (13938, -12.89 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.14 DPS) [world]; Ritssyn's Wand of Bad Mojo (22408, -24.20 DPS, sim-verified) [dungeon] |

**New at 60:** head: Fireleaf Circlet; neck: Amulet of the Dawn; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Mitts; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Trindlehaven Staff; ranged: Torch of Light

No-known-source sample (15 of 1011, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

