# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 31.1. Weights run: 1.1s. Verify run: 0.8s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.142 ± 0.005, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.010 per %), hit=0.216 ± 0.001 per rating point (10 rating = 1%, 2.156 per %), spell_haste=0.353 ± 0.085, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.618 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.66 DPS) | yes | Shadow Goggles (4373, -1.26 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.3 spell_power points (0.69 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.35 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.44 DPS) | yes | Feyscale Cloak (6632, -0.11 DPS) [dungeon]; Caretaker's Cape (20428, -0.11 DPS) [rep]; Black Whelp Cloak (7283, -0.22 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.7 spell_power points (0.63 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -0.78 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (31.1 DPS) | yes | Bright Bracers (3647, -0.02 DPS) [world_drop]; Repurposed Hair Band (281256, -0.05 DPS) [quest]; Windsong Bangles (263336, -0.62 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.77 DPS) | yes | Gnoll Casting Gloves (892, -0.11 DPS) [world]; Pristine Gloves (253913, -0.28 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.52 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.6 spell_power points (0.50 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.38 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.53 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.1 spell_power points (1.11 DPS) | yes | Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted]; Rumpled Kilt (274741, -0.56 DPS) [vendor]; Silk-threaded Trousers (1929, -0.93 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (0.83 DPS) | yes | Red Woolen Boots (4313, -0.39 DPS) [crafted]; Pristine Boots (253889, -0.45 DPS) [crafted]; Feather Padded Treads (285345, -0.48 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.3 spell_power points (0.58 DPS) | yes | Sludge-Stained Band (286535, -0.25 DPS) [world]; Lavishly Jeweled Ring (1156, -0.49 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.55 DPS) | yes | Sludge-Stained Band (286535, -0.22 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 1.4 spell_power points (0.16 DPS) | yes | Channeler's Staff (4437, +0.00 DPS, sim-verified) [world]; Lesser Staff of the Spire (1300, -0.06 DPS) [world]; Staff of Westfall (2042, -0.08 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 206.7 spell_power points (22.64 DPS) | yes | Skycaller (12984, -0.41 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.35 DPS) [dungeon]; Deepblaze (279896, -4.11 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 50.5. Weights run: 1.2s. Verify run: 0.8s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.280 ± 0.009, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.036 per %), hit=0.218 ± 0.002 per rating point (10 rating = 1%, 2.184 per %), spell_haste=not significant (0.157 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.692 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.51 DPS) | yes | Enchanter's Cowl (4322, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.41 DPS) [dungeon]; Silk Headband (7050, -0.65 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (1.19 DPS) | yes | Crystal Starfire Medallion (5003, -1.04 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.04 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.22 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.5 spell_power points (1.58 DPS) | yes | Death Speaker Mantle (6685, -0.33 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.41 DPS) [quest]; Invoker's Mantle (215365, -0.43 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (0.69 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.14 DPS) [crafted]; Prelacy Cape (7004, -0.14 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (1.79 DPS) | yes | Green Silk Armor (7065, -0.05 DPS) [crafted]; Death Speaker Robes (6682, -0.40 DPS) [dungeon]; Pristine Gown (253961, -0.56 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.24 DPS) | yes | Nightsky Wristbands (6407, -1.01 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.01 DPS, sim-verified) [world_drop]; Stonecloth Bindings (14416, -1.05 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.1 spell_power points (0.97 DPS) | yes | Serpent Gloves (5970, -0.01 DPS) [dungeon]; Shilly Mitts (9609, -0.01 DPS) [quest]; Gnoll Casting Gloves (892, -0.15 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.8 spell_power points (1.63 DPS) | yes | Belt of Arugal (6392, -0.28 DPS) [dungeon]; Invoker's Cord (215366, -0.47 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.53 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.65 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.42 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.0 spell_power points (1.23 DPS) | yes | Acidic Walkers (9454, -0.24 DPS) [dungeon]; Nimbus Boots (6998, -0.41 DPS) [quest]; Spidersilk Boots (4320, -1.56 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.96 DPS) | yes | Minor Channeling Ring (1449, -0.20 DPS) [quest]; Electrocutioner Lagnut (9447, -0.55 DPS) [dungeon]; Sludge-Stained Band (286535, -0.55 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.83 DPS) | yes | Electrocutioner Lagnut (9447, -0.41 DPS) [dungeon]; Sludge-Stained Band (286535, -0.41 DPS) [world]; Minor Channeling Ring (1449, -0.96 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.24 DPS) | yes | Glimmering Staff (249392, -0.70 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.85 DPS) [world_drop]; Channeler's Staff (4437, -0.93 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.7 spell_power points (1.19 DPS) | yes | Dwarven Tome (279898, -0.29 DPS, sim-verified) [quest]; Eye of Paleth (2943, -0.64 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -0.64 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 244.5 spell_power points (33.62 DPS) | yes | Starfaller (13063, -0.61 DPS) [world_drop]; Greater Mystic Wand (217287, -3.93 DPS) [crafted]; Gravestone Scepter (7001, -4.62 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 253225113100011400-00000000000000000-0000000000000000000)

Set DPS (verified): 170.3. Weights run: 1.5s. Verify run: 0.9s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.510 ± 0.022, crit=0.165 ± 0.004 per rating point (14 rating = 1%, 2.307 per %), hit=0.299 ± 0.005 per rating point (10 rating = 1%, 2.989 per %), spell_haste=not significant (1.153 ± 0.349), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.909 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.20 DPS) | yes | Living Cowl (5608, -2.36 DPS) [world]; Enchanter's Cowl (4322, -2.93 DPS) [crafted]; Augural Shroud (2620, -3.61 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.1 spell_power points (2.97 DPS) | yes | Triune Amulet (7722, -1.92 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.92 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.66 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.6 spell_power points (4.03 DPS) | yes | Green Silken Shoulders (7057, -0.01 DPS) [crafted]; Bloodmage Mantle (7684, -0.01 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.6 spell_power points (4.01 DPS) | yes | Guardian Cloak (5965, -1.49 DPS) [crafted]; Icy Cloak (4327, -1.95 DPS) [crafted]; Long Silken Cloak (4326, -2.54 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.1 spell_power points (7.40 DPS) | yes | Elemental Raiment (9434, -1.20 DPS) [world_drop]; Dreamweave Vest (10021, -1.25 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.46 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.66 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.59 DPS) [quest]; Windchaser Cuffs (14429, -1.30 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.0 spell_power points (5.92 DPS) | yes | Black Mageweave Gloves (10003, -1.49 DPS) [crafted]; Red Mageweave Gloves (10018, -2.44 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.50 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.0 spell_power points (4.74 DPS) | yes | Star Belt (4329, -0.90 DPS) [crafted]; Gilded Cord (254037, -1.17 DPS) [crafted]; Deathmage Sash (10771, -2.30 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.1 spell_power points (5.94 DPS) | yes | Abomination Skin Leggings (23173, -2.08 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.16 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -2.40 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.09 DPS) | yes | Gilded Slippers (254001, -3.47 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.41 DPS) [dungeon]; Spidersilk Boots (4320, -4.42 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.86 DPS) | yes | Ring of Forlorn Spirits (2043, -1.49 DPS) [quest]; Reedknot Ring (9622, -1.79 DPS) [quest]; Minor Channeling Ring (1449, -2.08 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.66 DPS) | yes | Ring of Forlorn Spirits (2043, -0.30 DPS) [quest]; Reedknot Ring (9622, -0.59 DPS) [quest]; Minor Channeling Ring (1449, -0.88 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (170.3 DPS) | yes | Scorn's Focal Dagger (23168, -3.25 DPS) [dungeon]; Windweaver Staff (7757, -3.65 DPS) [dungeon]; Gut Ripper (2164, -9.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 136.3 spell_power points (40.28 DPS) | yes | Nether Force Wand (11263, -2.50 DPS) [quest]; Icefury Wand (7514, -2.64 DPS) [quest]; Ragefire Wand (7513, -2.69 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 253225113100011531-03200000000000000-0000000000000000000)

Set DPS (verified): 279.4. Weights run: 1.5s. Verify run: 1.1s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.647 ± 0.035, crit=0.275 ± 0.006 per rating point (14 rating = 1%, 3.854 per %), hit=0.459 ± 0.007 per rating point (10 rating = 1%, 4.594 per %), spell_haste=not significant (0.973 ± 0.536), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.916 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 31.9 spell_power points (10.32 DPS) | yes | Dreamweave Circlet (10041, -1.45 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.60 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.04 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.9 spell_power points (6.43 DPS) | yes | Mindburst Medallion (11196, -3.23 DPS) [quest]; Horizon Choker (13085, -3.50 DPS) [world_drop]; Scorn's Icy Choker (23169, -6.02 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 24.7 spell_power points (7.97 DPS) | yes | Kentic Amice (11624, -0.72 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.25 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.57 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.9 spell_power points (5.78 DPS) | yes | Runecloth Cloak (13860, -1.20 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.86 DPS, sim-verified) [dungeon]; Big Voodoo Cloak (8216, -2.28 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 31.9 spell_power points (10.32 DPS) | yes | Robe of the Magi (1716, -1.96 DPS) [world_drop]; Runecloth Tunic (13857, -2.53 DPS) [crafted]; Dreamweave Vest (10021, -2.62 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.5 spell_power points (3.73 DPS) | yes | Aristocratic Cuffs (12546, -0.59 DPS) [dungeon]; Arcane Runed Bracers (4744, -0.82 DPS) [quest]; Bloodband Bracers (11469, -4.49 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 25.7 spell_power points (8.29 DPS) | yes | Dreamweave Gloves (10019, -1.64 DPS) [crafted]; Raider Handwraps (272098, -2.19 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.21 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (279.4 DPS) | yes | Dawnspire Cord (12466, -0.41 DPS) [dungeon]; Deathmage Sash (10771, -0.93 DPS) [dungeon]; Satyrmane Sash (17755, -3.66 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.5 spell_power points (9.53 DPS) | yes | Red Mageweave Pants (10009, -2.49 DPS) [crafted]; Wizardweave Leggings (14132, -3.38 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -6.51 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.76 DPS) | yes | Gilded Sandals (254107, -2.32 DPS) [crafted]; Black Mageweave Boots (10026, -2.74 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.88 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (4.49 DPS) | yes | Band of the Unicorn (7553, -0.29 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.61 DPS) [rep]; Brainlash (6440, -1.35 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.5 spell_power points (4.37 DPS) | yes | Band of the Unicorn (7553, -0.17 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.49 DPS) [rep]; Brainlash (6440, -1.23 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -0.40 DPS) [world_drop]; Spellshifter Rod (9527, -1.65 DPS) [quest]; Blade of Eternal Darkness (17780, -2.59 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 162.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -0.88 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.16 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 253225113100011531-03202300000000000-0050000000000000000)

Set DPS (verified): 466.2. Weights run: 1.5s. Verify run: 1.0s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.805 ± 0.046, crit=0.448 ± 0.010 per rating point (14 rating = 1%, 6.270 per %), hit=0.720 ± 0.010 per rating point (10 rating = 1%, 7.203 per %), spell_haste=not significant (1.956 ± 0.763), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.912 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.4 spell_power points (15.54 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.85 DPS) [pvp]; Crimson Felt Hat (18727, -3.59 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (466.2 DPS) | yes | Beads of Ogre Mojo (22149, -0.92 DPS) [quest]; Chains of the Lich (23125, -1.14 DPS) [dungeon]; Jewel of Kajaro (19601, -3.22 DPS, sim-verified) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.3 spell_power points (14.87 DPS) | yes | Field Marshal's Silk Spaulders (231602, -2.71 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.81 DPS) [crafted]; Darkspear Shoulderpads (272103, -9.16 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.6 spell_power points (9.72 DPS) | yes | Crystalline Threaded Cape (20697, -2.11 DPS) [world]; Hide of the Wild (18510, -2.49 DPS) [crafted]; Spritecaster Cape (11623, -3.55 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.9 spell_power points (18.34 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.98 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.91 DPS) [pvp]; Robe of Everlasting Night (18385, -5.98 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.4 spell_power points (9.33 DPS) | yes | Sublime Wristguards (18497, -2.75 DPS) [dungeon]; Runecloth Cuffs (254123, -3.08 DPS) [crafted]; Marshal's Silk Bracers (16438, -4.31 DPS) [pvp] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 31.0 spell_power points (10.17 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 49.9 spell_power points (16.37 DPS) | yes | Belt of the Archmage (18405, -3.76 DPS, sim-verified) [crafted]; Magician's Cord (272393, -3.80 DPS) [vendor]; Stormpike Cloth Girdle (19094, -7.82 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.6 spell_power points (17.24 DPS) | yes | Marshal's Silk Leggings (231605, -0.07 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -3.81 DPS) [pvp]; Skyshroud Leggings (13170, -3.98 DPS) [dungeon] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 33.9 spell_power points (11.11 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.98 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (466.2 DPS) | yes | Songstone of Ironforge (12543, -2.37 DPS) [quest]; Maiden's Circle (13001, -2.37 DPS) [world_drop]; Naglering (11669, -11.99 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (466.2 DPS) | yes | Songstone of Ironforge (12543, -0.98 DPS) [quest]; Maiden's Circle (13001, -0.98 DPS) [world_drop]; Naglering (11669, -11.39 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (466.2 DPS) | yes | Weakness Analyzer (272438, -2.30 DPS) [vendor]; Serenity Field (272439, -4.92 DPS) [vendor]; Burst of Knowledge (11832, -5.57 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (466.2 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -3.02 DPS, sim-verified) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (466.2 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.27 DPS) [world]; Teebu's Blazing Longsword (1728, -14.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 240.1 spell_power points (78.73 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.27 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.43 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.38 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 27.8. Weights run: 1.1s. Verify run: 0.8s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.142 ± 0.005, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.010 per %), hit=0.216 ± 0.001 per rating point (10 rating = 1%, 2.156 per %), spell_haste=0.353 ± 0.085, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.618 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.66 DPS) | yes | Shadow Goggles (4373, -1.01 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.3 spell_power points (0.69 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.25 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.48 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.44 DPS) | yes | Feyscale Cloak (6632, -0.11 DPS) [dungeon]; Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Battle Healer's Cloak (20427, -0.11 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.7 spell_power points (0.63 DPS) | yes | Green Woolen Vest (2582, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.19 DPS) [dungeon]; Gray Woolen Robe (2585, -0.52 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) | Earthen Arise [quest] | 1.3 spell_power points (0.14 DPS) | yes | Tabitha's Cuffs (251486, -0.05 DPS) [quest]; Featherbead Bracers (15452, -0.06 DPS) [quest]; Windsong Bangles (263336, -0.16 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.77 DPS) | yes | Gnoll Casting Gloves (892, -0.11 DPS) [world]; Pristine Gloves (253913, -0.28 DPS) [crafted]; Apothecary Gloves (10919, -0.33 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.6 spell_power points (0.50 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.27 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.38 DPS) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.1 spell_power points (1.11 DPS) | yes | Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted]; Rumpled Kilt (274741, -0.56 DPS) [vendor]; Silk-threaded Trousers (1929, -0.94 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (0.83 DPS) | yes | Red Woolen Boots (4313, -0.39 DPS) [crafted]; Feather Padded Treads (285345, -0.43 DPS, sim-verified) [world]; Pristine Boots (253889, -0.45 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.55 DPS) | yes | Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Loop of Sacrifice (281673, -0.47 DPS) [quest]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 spell_power points (0.33 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.25 DPS) [quest]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 1.4 spell_power points (0.16 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.06 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 206.7 spell_power points (22.64 DPS) | yes | Skycaller (12984, -0.54 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.35 DPS) [dungeon]; Sizzle Stick (8071, -4.44 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 45.6. Weights run: 1.2s. Verify run: 0.9s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.280 ± 0.009, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.036 per %), hit=0.218 ± 0.002 per rating point (10 rating = 1%, 2.184 per %), spell_haste=not significant (0.157 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.692 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.51 DPS) | yes | Enchanter's Cowl (4322, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.41 DPS) [dungeon]; Silk Headband (7050, -0.50 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.7 spell_power points (1.19 DPS) | yes | Crystal Starfire Medallion (5003, -1.04 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.04 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.30 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.5 spell_power points (1.58 DPS) | yes | Fairywing Mantle (9536, -0.41 DPS) [quest]; Invoker's Mantle (215365, -0.43 DPS) [crafted]; Death Speaker Mantle (6685, -0.45 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.69 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.14 DPS) [crafted]; Battle Healer's Cloak (19529, -0.14 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (45.6 DPS) | yes | Death Speaker Robes (6682, -0.35 DPS) [dungeon]; Pristine Gown (253961, -0.51 DPS) [crafted]; Tree Bark Jacket (1486, -0.54 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.24 DPS) | yes | Nightsky Wristbands (6407, -1.01 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.01 DPS) [quest]; Glowing Magical Bracelets (13106, -1.02 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.4 spell_power points (1.02 DPS) | yes | Gnoll Casting Gloves (892, -0.19 DPS) [world]; Truefaith Gloves (7049, -0.21 DPS) [crafted]; Serpent Gloves (5970, -0.57 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.8 spell_power points (1.63 DPS) | yes | Warsong Sash (16975, -0.12 DPS) [quest]; Belt of Arugal (6392, -0.28 DPS) [dungeon]; Invoker's Cord (215366, -0.47 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (1.65 DPS) | yes | Abomination Skin Leggings (23173, -0.10 DPS) [dungeon]; Pristine Leggings (253987, -0.42 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.59 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.0 spell_power points (1.23 DPS) | yes | Acidic Walkers (9454, -0.24 DPS) [dungeon]; Boots of the Enchanter (4325, -0.54 DPS) [crafted]; Spidersilk Boots (4320, -1.24 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.96 DPS) | yes | Electrocutioner Lagnut (9447, -0.55 DPS) [dungeon]; Sludge-Stained Band (286535, -0.55 DPS) [world]; Sacred Band (6669, -0.69 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.83 DPS) | yes | Electrocutioner Lagnut (9447, -0.41 DPS) [dungeon]; Sacred Band (6669, -0.55 DPS) [quest]; Sludge-Stained Band (286535, -1.00 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.24 DPS) | yes | Glimmering Staff (249392, -0.62 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -0.85 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.85 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.7 spell_power points (1.19 DPS) | yes | Orb of Souls (249395, -0.64 DPS) [crafted]; Tome of the Darkspear Prophecy (272090, -0.76 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -1.04 DPS, sim-verified) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 244.5 spell_power points (33.62 DPS) | yes | Starfaller (13063, -0.61 DPS) [world_drop]; Greater Mystic Wand (217287, -3.93 DPS) [crafted]; Gravestone Scepter (7001, -4.62 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 253225113100011400-00000000000000000-0000000000000000000)

Set DPS (verified): 162.8. Weights run: 1.5s. Verify run: 0.9s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.510 ± 0.022, crit=0.165 ± 0.004 per rating point (14 rating = 1%, 2.307 per %), hit=0.299 ± 0.005 per rating point (10 rating = 1%, 2.989 per %), spell_haste=not significant (1.153 ± 0.349), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.909 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.20 DPS) | yes | Augural Shroud (2620, -1.48 DPS, sim-verified) [world]; Living Cowl (5608, -2.36 DPS) [world]; Enchanter's Cowl (4322, -2.93 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.1 spell_power points (2.97 DPS) | yes | Triune Amulet (7722, -1.92 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.92 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -2.52 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.6 spell_power points (4.03 DPS) | yes | Green Silken Shoulders (7057, -0.01 DPS) [crafted]; Bloodmage Mantle (7684, -0.01 DPS) [dungeon]; Berylline Pads (4197, -0.45 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 13.6 spell_power points (4.01 DPS) | yes | Guardian Cloak (5965, -1.49 DPS) [crafted]; Icy Cloak (4327, -1.95 DPS) [crafted]; Long Silken Cloak (4326, -2.15 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.1 spell_power points (7.40 DPS) | yes | Dreamweave Vest (10021, -0.73 DPS) [crafted]; Elemental Raiment (9434, -1.20 DPS) [world_drop]; Robe of Power (7054, -1.46 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.66 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.27 DPS) [quest]; Condor Bracers (15864, -0.59 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.0 spell_power points (5.92 DPS) | yes | Black Mageweave Gloves (10003, -1.49 DPS) [crafted]; Red Mageweave Gloves (10018, -1.92 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.50 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.0 spell_power points (4.74 DPS) | yes | Deathmage Sash (10771, -0.41 DPS) [dungeon]; Star Belt (4329, -0.90 DPS) [crafted]; Gilded Cord (254037, -1.17 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.1 spell_power points (5.94 DPS) | yes | Crimson Silk Pantaloons (7062, -1.85 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.08 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.40 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.09 DPS) | yes | Gilded Slippers (254001, -2.73 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -4.41 DPS) [dungeon]; Spidersilk Boots (4320, -4.42 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.1 spell_power points (3.86 DPS) | yes | Reedknot Ring (9622, -1.79 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.09 DPS) [vendor]; Black Widow Band (6199, -2.80 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.66 DPS) | yes | Reedknot Ring (9622, -0.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.89 DPS) [vendor]; Black Widow Band (6199, -1.60 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (162.8 DPS) | yes | Scorn's Focal Dagger (23168, -3.25 DPS) [dungeon]; Windweaver Staff (7757, -3.65 DPS) [dungeon]; Gut Ripper (2164, -9.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 136.3 spell_power points (40.28 DPS) | yes | Nether Force Wand (11263, -2.50 DPS) [quest]; Icefury Wand (7514, -2.64 DPS) [quest]; Ragefire Wand (7513, -2.69 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253225113100011531-03200000000000000-0000000000000000000)

Set DPS (verified): 271.1. Weights run: 1.5s. Verify run: 1.1s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.647 ± 0.035, crit=0.275 ± 0.006 per rating point (14 rating = 1%, 3.854 per %), hit=0.459 ± 0.007 per rating point (10 rating = 1%, 4.594 per %), spell_haste=not significant (0.973 ± 0.536), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.916 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 31.9 spell_power points (10.32 DPS) | yes | Dreamweave Circlet (10041, -1.45 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.60 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.04 DPS) [vendor] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 19.9 spell_power points (6.43 DPS) | yes | Mindburst Medallion (11196, -3.23 DPS) [quest]; Horizon Choker (13085, -3.50 DPS) [world_drop]; Scorn's Icy Choker (23169, -5.80 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 24.7 spell_power points (7.97 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.25 DPS) [vendor]; Red Mageweave Shoulders (10029, -2.57 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.9 spell_power points (5.78 DPS) | yes | Deep Woodlands Cloak (19121, -0.02 DPS) [quest]; Mantle of Lady Falther'ess (23178, -0.99 DPS) [dungeon]; Runecloth Cloak (13860, -1.20 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 31.9 spell_power points (10.32 DPS) | yes | Robe of the Magi (1716, -1.96 DPS) [world_drop]; Runecloth Tunic (13857, -2.53 DPS) [crafted]; Dreamweave Vest (10021, -2.62 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 11.5 spell_power points (3.73 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -3.23 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 25.7 spell_power points (8.29 DPS) | yes | Raider Handwraps (272098, -1.43 DPS) [vendor]; Dreamweave Gloves (10019, -1.64 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -2.21 DPS) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (271.1 DPS) | yes | Dawnspire Cord (12466, -0.41 DPS) [dungeon]; Deathmage Sash (10771, -0.93 DPS) [dungeon]; Satyrmane Sash (17755, -4.78 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.5 spell_power points (9.53 DPS) | yes | Red Mageweave Pants (10009, -2.49 DPS) [crafted]; Wizardweave Leggings (14132, -3.38 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -5.43 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.76 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -1.80 DPS) [vendor]; Gilded Sandals (254107, -2.32 DPS) [crafted]; Black Mageweave Boots (10026, -2.74 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (4.49 DPS) | yes | Band of the Unicorn (7553, -0.29 DPS) [world_drop]; Advisor's Ring (19519, -0.61 DPS) [rep]; Brainlash (6440, -1.35 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.5 spell_power points (4.37 DPS) | yes | Band of the Unicorn (7553, -0.17 DPS) [world_drop]; Advisor's Ring (19519, -0.49 DPS) [rep]; Brainlash (6440, -1.23 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -0.40 DPS) [world_drop]; Spellshifter Rod (9527, -1.65 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 162.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.16 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Ban'thok Sash; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 253225113100011531-03202300000000000-0050000000000000000)

Set DPS (verified): 455.8. Weights run: 1.5s. Verify run: 1.0s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.805 ± 0.046, crit=0.448 ± 0.010 per rating point (14 rating = 1%, 6.270 per %), hit=0.720 ± 0.010 per rating point (10 rating = 1%, 7.203 per %), spell_haste=not significant (1.956 ± 0.763), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.912 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 47.4 spell_power points (15.54 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.85 DPS) [pvp]; Crimson Felt Hat (18727, -3.59 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (455.8 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS) [quest]; Beads of Ogre Mojo (22149, -0.92 DPS) [quest]; Chains of the Lich (23125, -1.14 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 45.3 spell_power points (14.87 DPS) | yes | Warlord's Silk Amice (231594, -2.71 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.81 DPS) [crafted]; Darkspear Shoulderpads (272103, -8.05 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 29.6 spell_power points (9.72 DPS) | yes | Crystalline Threaded Cape (20697, -2.11 DPS) [world]; Hide of the Wild (18510, -2.49 DPS) [crafted]; Deep Woodlands Cloak (19121, -3.41 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.9 spell_power points (18.34 DPS) | yes | Warlord's Silk Raiment (231596, -0.98 DPS) [pvp]; Robe of Everlasting Night (18385, -4.91 DPS, sim-verified) [dungeon]; Legionnaire's Silk Tunic (227106, -4.91 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.4 spell_power points (9.33 DPS) | yes | Sublime Wristguards (18497, -2.75 DPS) [dungeon]; Runecloth Cuffs (254123, -3.08 DPS) [crafted]; General's Silk Cuffs (16538, -4.31 DPS) [pvp] |
| hands | Sandworm Skin Gloves (20716) | Armaments of War [quest] | 31.0 spell_power points (10.17 DPS) | yes | Hands of Power (13253, +0.00 DPS, sim-verified) [dungeon]; General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 49.9 spell_power points (16.37 DPS) | yes | Belt of the Archmage (18405, -3.48 DPS, sim-verified) [crafted]; Magician's Cord (272393, -3.80 DPS) [vendor]; Frostwolf Cloth Belt (19090, -7.82 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.6 spell_power points (17.24 DPS) | yes | General's Silk Trousers (231595, -0.07 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -3.81 DPS) [pvp]; Outrider's Silk Leggings (22747, -6.63 DPS, sim-verified) [rep] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 33.9 spell_power points (11.11 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.98 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (455.8 DPS) | yes | Eye of Orgrimmar (12545, -2.37 DPS) [quest]; Maiden's Circle (13001, -2.37 DPS) [world_drop]; Naglering (11669, -10.68 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (455.8 DPS) | yes | Eye of Orgrimmar (12545, -0.98 DPS) [quest]; Maiden's Circle (13001, -0.98 DPS) [world_drop]; Naglering (11669, -11.54 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (455.8 DPS) | yes | Weakness Analyzer (272438, -2.30 DPS) [vendor]; Serenity Field (272439, -4.92 DPS) [vendor]; Burst of Knowledge (11832, -5.57 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (455.8 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (455.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.81 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -20.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 240.1 spell_power points (78.73 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.27 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.43 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.38 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Sandworm Skin Gloves; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Draconic Infused Emblem; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

