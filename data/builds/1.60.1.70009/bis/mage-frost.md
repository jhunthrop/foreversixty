# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 37.5. Weights run: 0.7s. Verify run: 0.6s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.218 ± 0.007, crit=0.087 ± 0.004 per rating point (14 rating = 1%, 1.211 per %), hit=0.250 ± 0.002 per rating point (10 rating = 1%, 2.500 per %), spell_haste=0.732 ± 0.122, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.76 DPS) | yes | Shadow Goggles (4373, -1.42 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.0 spell_power points (0.88 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.37 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.40 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.50 DPS) | yes | Caretaker's Cape (20428, +0.00 DPS, sim-verified) [rep]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.1 spell_power points (0.77 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -1.02 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.1 spell_power points (0.14 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Repurposed Hair Band (281256, -0.08 DPS) [quest]; Windsong Bangles (263336, -0.69 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.88 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.30 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.58 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.61 DPS) | yes | Novice Ardent's Sash (253887, -0.28 DPS) [crafted]; Keller's Girdle (2911, -0.39 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.60 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.7 spell_power points (1.35 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.47 DPS) [dungeon]; Rumpled Kilt (274741, -0.72 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.99 DPS) | yes | Red Woolen Boots (4313, -0.49 DPS) [crafted]; Pristine Boots (253889, -0.53 DPS) [crafted]; Feather Padded Treads (285345, -0.74 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.4 spell_power points (0.68 DPS) | yes | Sludge-Stained Band (286535, -0.31 DPS) [world]; Lavishly Jeweled Ring (1156, -0.52 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.60 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.63 DPS) | yes | Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.55 DPS) [world_drop]; Sludge-Stained Band (286535, -0.83 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 2.2 spell_power points (0.27 DPS) | yes | Lesser Staff of the Spire (1300, -0.11 DPS) [world]; Staff of Westfall (2042, -0.14 DPS) [quest]; Channeler's Staff (4437, -0.35 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 180.1 spell_power points (22.69 DPS) | yes | Skycaller (12984, -1.07 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.40 DPS) [dungeon]; Deepblaze (279896, -4.16 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 61.8. Weights run: 0.7s. Verify run: 0.7s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.309 ± 0.013, crit=0.151 ± 0.008 per rating point (14 rating = 1%, 2.110 per %), hit=0.262 ± 0.003 per rating point (10 rating = 1%, 2.617 per %), spell_haste=not significant (-0.035 ± 0.220), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.67 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.46 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (1.35 DPS) | yes | Crystal Starfire Medallion (5003, -1.16 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.16 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.40 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.8 spell_power points (1.79 DPS) | yes | Death Speaker Mantle (6685, -0.24 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Invoker's Mantle (215365, -0.49 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.76 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Prelacy Cape (7004, -0.15 DPS) [quest]; Caretaker's Cape (19533, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.0 spell_power points (1.98 DPS) | yes | Death Speaker Robes (6682, -0.40 DPS) [dungeon]; Pristine Gown (253961, -0.59 DPS) [crafted]; Tree Bark Jacket (1486, -1.28 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.37 DPS) | yes | Nightsky Wristbands (6407, -1.09 DPS) [world_drop]; Stonecloth Bindings (14416, -1.13 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.32 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.4 spell_power points (1.13 DPS) | yes | Serpent Gloves (5970, -0.06 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.21 DPS) [world]; Shilly Mitts (9609, -0.70 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 spell_power points (1.81 DPS) | yes | Belt of Arugal (6392, -0.17 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.51 DPS) [crafted]; Crimson Silk Belt (7055, -0.57 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.83 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.43 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.63 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.2 spell_power points (1.39 DPS) | yes | Acidic Walkers (9454, -0.26 DPS) [dungeon]; Nimbus Boots (6998, -0.48 DPS) [quest]; Spidersilk Boots (4320, -1.97 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.06 DPS) | yes | Minor Channeling Ring (1449, -0.21 DPS) [quest]; Lorekeeper's Ring (20431, -0.30 DPS) [rep]; Electrocutioner Lagnut (9447, -0.61 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.91 DPS) | yes | Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.46 DPS) [world]; Minor Channeling Ring (1449, -1.00 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.37 DPS) | yes | Glimmering Staff (249392, -0.60 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.90 DPS) [world_drop]; Channeler's Staff (4437, -0.99 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.9 spell_power points (1.35 DPS) | yes | Eye of Paleth (2943, -0.74 DPS) [quest]; Orb of Souls (249395, -0.74 DPS) [crafted]; Dwarven Tome (279898, -1.19 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 221.3 spell_power points (33.67 DPS) | yes | Starfaller (13063, -0.07 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.67 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 98.3. Weights run: 0.7s. Verify run: 0.6s. 328 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.687 ± 0.033, crit=0.230 ± 0.013 per rating point (14 rating = 1%, 3.219 per %), hit=0.400 ± 0.005 per rating point (10 rating = 1%, 3.997 per %), spell_haste=4.878 ± 0.478, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.10 DPS) | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Corpseshroud (10574, -1.17 DPS) [dungeon]; Living Cowl (5608, -1.18 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.1 spell_power points (1.64 DPS) | yes | Necklace of Calisea (1714, -0.93 DPS) [dungeon]; Triune Amulet (7722, -0.93 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.15 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.9 spell_power points (2.35 DPS) | yes | Green Silken Shoulders (7057, -0.09 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.2 spell_power points (2.24 DPS) | yes | Guardian Cloak (5965, -0.85 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.13 DPS) [vendor]; Long Silken Cloak (4326, -1.20 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.1 spell_power points (3.86 DPS) | yes | Dreamweave Vest (10021, -0.27 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.57 DPS) [crafted]; Elemental Raiment (9434, -0.76 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.33 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.30 DPS) [quest]; Windchaser Cuffs (14429, -0.42 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 spell_power points (3.07 DPS) | yes | Black Mageweave Gloves (10003, -0.85 DPS) [crafted]; Gilded Handwraps (254021, -1.17 DPS) [crafted]; Red Mageweave Gloves (10018, -1.30 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.3 spell_power points (2.56 DPS) | yes | Highlander's Cloth Girdle (20098, -0.22 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.56 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.2 spell_power points (3.29 DPS) | yes | Crimson Silk Pantaloons (7062, -1.08 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.14 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.44 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.55 DPS) | yes | Gilded Slippers (254001, -1.62 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.99 DPS) [dungeon]; Spidersilk Boots (4320, -2.11 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.09 DPS) | yes | Ring of Forlorn Spirits (2043, -0.90 DPS) [quest]; Reedknot Ring (9622, -1.05 DPS) [quest]; Minor Channeling Ring (1449, -1.14 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.33 DPS) | yes | Reedknot Ring (9622, -0.30 DPS) [quest]; Lorekeeper's Ring (19525, -0.30 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.46 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (98.3 DPS) | yes | Windweaver Staff (7757, -1.43 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.62 DPS) [dungeon]; Gut Ripper (2164, -4.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 270.7 spell_power points (39.99 DPS) | yes | Nether Force Wand (11263, -0.91 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.49 DPS) [quest]; Ragefire Wand (7513, -2.54 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 134.7. Weights run: 0.7s. Verify run: 0.8s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.911 ± 0.050, crit=0.368 ± 0.021 per rating point (14 rating = 1%, 5.153 per %), hit=0.574 ± 0.009 per rating point (10 rating = 1%, 5.744 per %), spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.2 spell_power points (4.88 DPS) | yes | Dreamweave Circlet (10041, -0.79 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -0.94 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.34 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (2.82 DPS) | yes | Scorn's Icy Choker (23169, -1.18 DPS) [dungeon]; Mindburst Medallion (11196, -1.31 DPS) [quest]; Horizon Choker (13085, -2.85 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.4 spell_power points (3.86 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.06 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.5 spell_power points (2.55 DPS) | yes | Runecloth Cloak (13860, -0.42 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -0.56 DPS, sim-verified) [dungeon]; Big Voodoo Cloak (8216, -0.82 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.2 spell_power points (4.88 DPS) | yes | Robe of the Magi (1716, -1.31 DPS, sim-verified) [world_drop]; Knight's Dreadweave Vest (220886, -1.32 DPS) [vendor]; Runecloth Tunic (13857, -1.34 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.7 spell_power points (1.79 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -0.06 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.36 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.5 spell_power points (4.00 DPS) | yes | Dreamweave Gloves (10019, -1.16 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.22 DPS) [vendor]; Raider Handwraps (272098, -1.48 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.3 spell_power points (3.06 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.09 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.1 spell_power points (4.21 DPS) | yes | Red Mageweave Pants (10009, -0.94 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.61 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.86 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.15 DPS) | yes | Gilded Sandals (254107, -0.63 DPS) [crafted]; Southsea Mojo Boots (20641, -0.78 DPS) [quest]; Sergeant Major's Dreadweave Boots (220891, -1.31 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.03 DPS) | yes | Brainlash (6440, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.45 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.02 DPS) | yes | Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.44 DPS) [rep]; Brainlash (6440, -1.54 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (134.7 DPS) | yes | Spellshifter Rod (9527, -0.72 DPS) [quest]; Spellforce Rod (1664, -0.84 DPS) [world]; Blade of Eternal Darkness (17780, -0.94 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.84 DPS) [dungeon]; Woestave (20082, -2.19 DPS, sim-verified) [quest]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; finger2: Cyclopean Band; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 257.8. Weights run: 0.8s. Verify run: 0.8s. 1035 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.019 ± 0.072, crit=0.597 ± 0.035 per rating point (14 rating = 1%, 8.359 per %), hit=1.081 ± 0.014 per rating point (10 rating = 1%, 10.815 per %), spell_haste=not significant (4.959 ± 1.418), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 87.7 spell_power points (10.76 DPS) | yes | Fireleaf Hood (240048, -1.54 DPS, sim-verified) [vendor]; Field Marshal's Coronet (231604, -3.56 DPS) [pvp] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (257.8 DPS) | yes | Beads of Ogre Mojo (22149, -0.37 DPS) [quest]; Pebble of Kajaro (19600, -0.74 DPS) [quest]; Jewel of Kajaro (19601, -1.51 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 64.6 spell_power points (7.93 DPS) | yes | Fireleaf Mantle (240046, -1.35 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -1.72 DPS) [vendor]; Field Marshal's Silk Spaulders (16444, -2.99 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.0 spell_power points (4.29 DPS) | yes | Crystalline Threaded Cape (20697, -1.34 DPS) [world]; Spritecaster Cape (11623, -1.82 DPS) [dungeon]; Hide of the Wild (18510, -2.11 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 99.1 spell_power points (12.15 DPS) | yes | Fireleaf Garb (240051, -1.98 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -4.72 DPS) [crafted]; Field Marshal's Silk Vestments (16443, -4.95 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 59.7 spell_power points (7.32 DPS) | yes | Fireleaf Wristwraps (240044, -2.07 DPS, sim-verified) [vendor]; Dryad's Wrist Bindings (19595, -3.62 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 73.0 spell_power points (8.95 DPS) | yes | Fireleaf Mitts (240049, -1.41 DPS, sim-verified) [vendor]; Marshal's Silk Gloves (16440, -4.14 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -4.14 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 70.1 spell_power points (8.59 DPS) | yes | Knowledge of the Timbermaw (228190, -1.45 DPS) [vendor]; Fireleaf Waistguard (240045, -2.69 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -3.12 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 87.7 spell_power points (10.76 DPS) | yes | Fireleaf Pants (240047, -1.70 DPS, sim-verified) [vendor]; Sentinel's Silk Leggings (237815, -3.53 DPS) [vendor]; Marshal's Silk Leggings (16442, -3.55 DPS) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 65.6 spell_power points (8.05 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (231606, -2.40 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (257.8 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.37 DPS) [vendor]; Naglering (11669, -13.49 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (257.8 DPS) | yes | Songstone of Ironforge (12543, -0.99 DPS) [quest]; Maiden's Circle (13001, -0.99 DPS) [world_drop]; Naglering (11669, -10.32 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (257.8 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -1.55 DPS, sim-verified) [dungeon] |
| trinket2 | - | - |  |  |  |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (257.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.90 DPS) [world]; Teebu's Blazing Longsword (1728, -13.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 631.8 spell_power points (77.50 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -4.20 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -13.02 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.43 DPS) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Amulet of the Dawn; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1035, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 33.6. Weights run: 0.7s. Verify run: 0.7s. 141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.218 ± 0.007, crit=0.087 ± 0.004 per rating point (14 rating = 1%, 1.211 per %), hit=0.250 ± 0.002 per rating point (10 rating = 1%, 2.500 per %), spell_haste=0.732 ± 0.122, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.76 DPS) | yes | Shadow Goggles (4373, -1.01 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.0 spell_power points (0.88 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.37 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.59 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.50 DPS) | yes | Battle Healer's Cloak (20427, +0.00 DPS, sim-verified) [rep]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.1 spell_power points (0.77 DPS) | yes | Green Woolen Vest (2582, -0.26 DPS) [crafted]; Bloody Apron (6226, -0.26 DPS) [dungeon]; Gray Woolen Robe (2585, -0.71 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | sim-verified (33.6 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest]; Owlbeard Bracers (16981, -0.47 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.88 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.30 DPS) [crafted]; Apothecary Gloves (10919, -0.38 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.9 spell_power points (0.61 DPS) | yes | Novice Ardent's Sash (253887, -0.28 DPS) [crafted]; Keller's Girdle (2911, -0.39 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.41 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.7 spell_power points (1.35 DPS) | yes | Filigreed Pristine Leggings (253937, -0.08 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.47 DPS) [dungeon]; Rumpled Kilt (274741, -0.72 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.9 spell_power points (0.99 DPS) | yes | Red Woolen Boots (4313, -0.49 DPS) [crafted]; Pristine Boots (253889, -0.53 DPS) [crafted]; Feather Padded Treads (285345, -0.88 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.63 DPS) | yes | Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon]; Loop of Sacrifice (281673, -0.49 DPS) [quest]; Volcanic Rock Ring (12053, -0.55 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.38 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.24 DPS) [quest]; Volcanic Rock Ring (12053, -0.30 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 2.2 spell_power points (0.27 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.05 DPS) [world]; Lesser Staff of the Spire (1300, -0.11 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 180.1 spell_power points (22.69 DPS) | yes | Skycaller (12984, -1.47 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.40 DPS) [dungeon]; Sizzle Stick (8071, -4.41 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 55.8. Weights run: 0.7s. Verify run: 0.7s. 236 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.309 ± 0.013, crit=0.151 ± 0.008 per rating point (14 rating = 1%, 2.110 per %), hit=0.262 ± 0.003 per rating point (10 rating = 1%, 2.617 per %), spell_haste=not significant (-0.035 ± 0.220), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.67 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.46 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 spell_power points (1.35 DPS) | yes | Crystal Starfire Medallion (5003, -1.16 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.16 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.53 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.8 spell_power points (1.79 DPS) | yes | Death Speaker Mantle (6685, -0.06 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Invoker's Mantle (215365, -0.49 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.76 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.0 spell_power points (1.98 DPS) | yes | Death Speaker Robes (6682, -0.40 DPS) [dungeon]; Pristine Gown (253961, -0.59 DPS) [crafted]; Tree Bark Jacket (1486, -1.12 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.37 DPS) | yes | Nightsky Wristbands (6407, -1.09 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.09 DPS) [quest]; Glowing Magical Bracelets (13106, -1.15 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.5 spell_power points (1.15 DPS) | yes | Gnoll Casting Gloves (892, -0.24 DPS) [world]; Truefaith Gloves (7049, -0.25 DPS) [crafted]; Serpent Gloves (5970, -0.58 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 spell_power points (1.81 DPS) | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Invoker's Cord (215366, -0.51 DPS) [crafted]; Warsong Sash (16975, -0.69 DPS, sim-verified) [quest] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.83 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.43 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.63 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.2 spell_power points (1.39 DPS) | yes | Acidic Walkers (9454, -0.26 DPS) [dungeon]; Boots of the Enchanter (4325, -0.63 DPS) [crafted]; Spidersilk Boots (4320, -1.78 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.06 DPS) | yes | Advisor's Ring (20426, -0.30 DPS) [rep]; Electrocutioner Lagnut (9447, -0.61 DPS) [dungeon]; Sludge-Stained Band (286535, -0.61 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.91 DPS) | yes | Sludge-Stained Band (286535, -0.46 DPS) [world]; Snake Hoop (6750, -0.58 DPS) [quest]; Electrocutioner Lagnut (9447, -1.36 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.37 DPS) | yes | Glimmering Staff (249392, -0.56 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.90 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.90 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.9 spell_power points (1.35 DPS) | yes | Orb of Souls (249395, -0.67 DPS, sim-verified) [crafted]; Alliance Outrunner Healing Rod (285348, -0.74 DPS) [world]; Tome of the Darkspear Prophecy (272090, -0.85 DPS) [vendor] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 221.3 spell_power points (33.67 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.67 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 236, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 89.5. Weights run: 0.7s. Verify run: 0.7s. 318 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.687 ± 0.033, crit=0.230 ± 0.013 per rating point (14 rating = 1%, 3.219 per %), hit=0.400 ± 0.005 per rating point (10 rating = 1%, 3.997 per %), spell_haste=4.878 ± 0.478, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.10 DPS) | yes | Augural Shroud (2620, -0.49 DPS, sim-verified) [world]; Corpseshroud (10574, -1.17 DPS) [dungeon]; Living Cowl (5608, -1.18 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.1 spell_power points (1.64 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.50 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -0.93 DPS) [dungeon]; Triune Amulet (7722, -0.93 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.9 spell_power points (2.35 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.11 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.2 spell_power points (2.24 DPS) | yes | Guardian Cloak (5965, -0.85 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.13 DPS) [vendor]; Long Silken Cloak (4326, -1.32 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.1 spell_power points (3.86 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.57 DPS) [crafted]; Elemental Raiment (9434, -0.76 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.5 spell_power points (1.40 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.37 DPS) [quest]; Windchaser Cuffs (14429, -0.49 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 spell_power points (3.07 DPS) | yes | Black Mageweave Gloves (10003, -0.85 DPS) [crafted]; Red Mageweave Gloves (10018, -1.07 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.17 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.3 spell_power points (2.56 DPS) | yes | Gilded Cord (254037, -0.56 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.63 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.2 spell_power points (3.29 DPS) | yes | Crimson Silk Pantaloons (7062, -0.76 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.14 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.44 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.55 DPS) | yes | Gilded Slippers (254001, -1.42 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.99 DPS) [dungeon]; Spidersilk Boots (4320, -2.11 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (2.09 DPS) | yes | Reedknot Ring (9622, -1.05 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.20 DPS) [vendor]; Voodoo Band (1996, -1.38 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.33 DPS) | yes | Advisor's Ring (19521, -0.30 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.44 DPS) [vendor]; Reedknot Ring (9622, -1.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (89.5 DPS) | yes | Windweaver Staff (7757, -1.43 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.62 DPS) [dungeon]; Gut Ripper (2164, -4.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 270.7 spell_power points (39.99 DPS) | yes | Nether Force Wand (11263, -2.36 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.49 DPS) [quest]; Ragefire Wand (7513, -2.54 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 123.3. Weights run: 0.7s. Verify run: 0.8s. 414 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.911 ± 0.050, crit=0.368 ± 0.021 per rating point (14 rating = 1%, 5.153 per %), hit=0.574 ± 0.009 per rating point (10 rating = 1%, 5.744 per %), spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.2 spell_power points (4.88 DPS) | yes | Dreamweave Circlet (10041, -0.05 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Hat (220907, -0.94 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.34 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.5 spell_power points (2.82 DPS) | yes | Scorn's Icy Choker (23169, -1.18 DPS) [dungeon]; Mindburst Medallion (11196, -1.31 DPS) [quest]; Horizon Choker (13085, -1.70 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.4 spell_power points (3.86 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.06 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.2 spell_power points (2.65 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.39 DPS) [dungeon]; Spritecaster Cape (11623, -0.41 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.51 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.2 spell_power points (4.88 DPS) | yes | Robe of the Magi (1716, -0.81 DPS, sim-verified) [world_drop]; Stone Guard's Dreadweave Vest (220904, -1.32 DPS) [vendor]; Runecloth Tunic (13857, -1.34 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (123.3 DPS) | yes | Bloodband Bracers (11469, -0.02 DPS) [quest]; Radiant Silver Bracers (4545, -0.27 DPS) [quest]; Aristocratic Cuffs (12546, -1.24 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 30.5 spell_power points (4.00 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.16 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.22 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.3 spell_power points (3.06 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.09 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.1 spell_power points (4.21 DPS) | yes | Red Mageweave Pants (10009, -0.94 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.61 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -1.70 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.15 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.07 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -0.63 DPS) [crafted]; Southsea Mojo Boots (20641, -0.78 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.03 DPS) | yes | Brainlash (6440, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Advisor's Ring (19519, -0.45 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.4 spell_power points (2.02 DPS) | yes | Brainlash (6440, -0.22 DPS, sim-verified) [dungeon]; Band of the Unicorn (7553, -0.31 DPS) [world_drop]; Advisor's Ring (19519, -0.44 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Blade of Eternal Darkness (17780, -0.64 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.72 DPS) [quest]; Spellforce Rod (1664, -0.84 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Woestave (20082, +0.00 DPS, sim-verified) [quest]; Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; finger2: Cyclopean Band; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 231.4. Weights run: 0.8s. Verify run: 0.9s. 1026 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.019 ± 0.072, crit=0.597 ± 0.035 per rating point (14 rating = 1%, 8.359 per %), hit=1.081 ± 0.014 per rating point (10 rating = 1%, 10.815 per %), spell_haste=not significant (4.959 ± 1.418), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 87.7 spell_power points (10.76 DPS) | yes | Fireleaf Hood (240048, -1.83 DPS, sim-verified) [vendor]; Warlord's Silk Cowl (231601, -3.56 DPS) [pvp] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points (0.00 DPS) | yes | Beads of Ogre Mojo (22149, -0.37 DPS) [quest]; Pebble of Kajaro (19600, -0.74 DPS) [quest]; Jewel of Kajaro (19601, -0.86 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 64.6 spell_power points (7.93 DPS) | yes | Fireleaf Mantle (240046, -1.67 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -1.72 DPS) [vendor]; Warlord's Silk Amice (16536, -2.99 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 35.0 spell_power points (4.29 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -1.34 DPS) [world]; Deep Woodlands Cloak (19121, -1.69 DPS) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 99.1 spell_power points (12.15 DPS) | yes | Fireleaf Garb (240051, -1.68 DPS, sim-verified) [vendor]; Robe of the Archmage (14152, -4.72 DPS) [crafted]; Warlord's Silk Raiment (16535, -4.95 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 59.7 spell_power points (7.32 DPS) | yes | Fireleaf Wristwraps (240044, -2.27 DPS, sim-verified) [vendor]; Dryad's Wrist Bindings (19595, -3.62 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 73.0 spell_power points (8.95 DPS) | yes | Fireleaf Mitts (240049, -1.20 DPS, sim-verified) [vendor]; General's Silk Handguards (16540, -4.14 DPS) [vendor]; General's Silk Gauntlets (231599, -4.14 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 70.1 spell_power points (8.59 DPS) | yes | Fireleaf Waistguard (240045, -1.16 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -1.45 DPS) [vendor]; Belt of the Archmage (18405, -3.12 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 87.7 spell_power points (10.76 DPS) | yes | Fireleaf Pants (240047, -1.88 DPS, sim-verified) [vendor]; Sentinel's Silk Leggings (237815, -3.53 DPS) [vendor]; General's Silk Trousers (16534, -3.55 DPS) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 65.6 spell_power points (8.05 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (231597, -2.40 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.37 DPS) [vendor]; Naglering (11669, -11.36 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 0.0 spell_power points (0.00 DPS) | yes | Eye of Orgrimmar (12545, -0.99 DPS) [quest]; Maiden's Circle (13001, -0.99 DPS) [world_drop]; Naglering (11669, -8.01 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | 0.0 spell_power points (0.00 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.23 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -11.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 631.8 spell_power points (77.50 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -5.46 DPS, sim-verified) [dungeon]; Bonecreeper Stylus (13938, -13.02 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.43 DPS) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Amulet of the Dawn; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1026, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

