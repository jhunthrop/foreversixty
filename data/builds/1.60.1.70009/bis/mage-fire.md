# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 32.9. Weights run: 1.1s. Verify run: 0.8s. 149 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.617 ± 0.012, crit=0.140 ± 0.005 per rating point (14 rating = 1%, 1.962 per %), hit=0.389 ± 0.002 per rating point (10 rating = 1%, 3.886 per %), spell_haste=not significant (-0.587 ± 0.226), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.44 DPS) | yes | Shadow Goggles (4373, -0.61 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.6 spell_power points (0.78 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.32 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.30 DPS) | yes | Pearl-clasped Cloak (5542, -0.01 DPS) [crafted]; Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Black Whelp Cloak (7283, -0.07 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.1 spell_power points (0.60 DPS) | yes | Green Woolen Robe (6243, -0.24 DPS) [crafted]; Mystic's Wrap (14369, -0.28 DPS) [world_drop]; Gray Woolen Robe (2585, -0.77 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 3.1 spell_power points (0.23 DPS) | yes | Mystic's Bracelets (14366, -0.14 DPS) [world_drop]; Repurposed Hair Band (281256, -0.14 DPS) [quest]; Bright Bracers (3647, -0.52 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.52 DPS) | yes | Gnoll Casting Gloves (892, -0.07 DPS) [world]; Pristine Gloves (253913, -0.09 DPS) [crafted]; Tomb Robber's Gloves (280096, -0.24 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.5 spell_power points (0.48 DPS) | yes | Keller's Girdle (2911, -0.11 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.65 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.9 spell_power points (1.03 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.51 DPS) [dungeon]; Rumpled Kilt (274741, -0.66 DPS) [vendor] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.5 spell_power points (0.70 DPS) | yes | Feather Padded Treads (285345, -0.27 DPS, sim-verified) [world]; Pristine Boots (253889, -0.34 DPS) [crafted]; Red Woolen Boots (4313, -0.40 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.46 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Sludge-Stained Band (286535, -0.24 DPS) [world]; Volcanic Rock Ring (12053, -0.32 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.37 DPS) | yes | Sludge-Stained Band (286535, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.84 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 6.2 spell_power points (0.46 DPS) | yes | Channeler's Staff (4437, -0.09 DPS) [world]; Lesser Staff of the Spire (1300, -0.18 DPS) [world]; Staff of Westfall (2042, -0.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 304.3 spell_power points (22.53 DPS) | yes | Skycaller (12984, -0.75 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.24 DPS) [dungeon]; Deepblaze (279896, -4.00 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 149, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 58.9. Weights run: 1.2s. Verify run: 0.8s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.948 ± 0.026, crit=0.225 ± 0.012 per rating point (14 rating = 1%, 3.154 per %), hit=0.419 ± 0.004 per rating point (10 rating = 1%, 4.189 per %), spell_haste=not significant (-0.185 ± 0.403), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.5 spell_power points (1.48 DPS) | yes | Nightsky Cowl (4039, -0.39 DPS) [world_drop]; Holy Shroud (2721, -0.43 DPS) [world_drop]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 spell_power points (1.21 DPS) | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.14 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.5 spell_power points (1.68 DPS) | yes | Death Speaker Mantle (6685, -0.11 DPS) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.6 spell_power points (0.72 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Repairman's Cape (9605, -0.08 DPS) [quest]; Resilient Cape (14400, -0.18 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.3 spell_power points (2.04 DPS) | yes | Death Speaker Robes (6682, -0.51 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.68 DPS) [dungeon]; Pristine Gown (253961, -0.73 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.86 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Stonecloth Bindings (14416, -0.41 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.58 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 14.4 spell_power points (1.38 DPS) | yes | Truefaith Gloves (7049, -0.63 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.65 DPS) [dungeon]; Shilly Mitts (9609, -0.71 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.8 spell_power points (1.32 DPS) | yes | Crimson Silk Belt (7055, -0.12 DPS) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (58.9 DPS) | yes | Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Abomination Skin Leggings (23173, -0.79 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.6 spell_power points (1.30 DPS) | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -1.17 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.67 DPS) | yes | Black Widow Band (6199, -0.04 DPS) [world]; Snake Hoop (6750, -0.04 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 spell_power points (0.66 DPS) | yes | Black Widow Band (6199, -0.03 DPS) [world]; Snake Hoop (6750, -0.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.4 spell_power points (1.00 DPS) | yes | Twisted Chanter's Staff (890, -0.09 DPS) [world_drop]; Scorn's Focal Dagger (23168, -0.14 DPS) [dungeon]; Channeler's Staff (4437, -0.27 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 350.3 spell_power points (33.50 DPS) | yes | Starfaller (13063, -0.27 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 105.3. Weights run: 1.2s. Verify run: 0.7s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.865 ± 0.039, crit=0.197 ± 0.014 per rating point (14 rating = 1%, 2.758 per %), hit=0.487 ± 0.006 per rating point (10 rating = 1%, 4.875 per %), spell_haste=6.790 ± 0.684, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.76 DPS) | yes | Augural Shroud (2620, -0.18 DPS) [world]; Corpseshroud (10574, -0.60 DPS) [dungeon]; Thinking Cap (2624, -0.83 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.2 spell_power points (1.60 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.47 DPS) [quest]; Triune Amulet (7722, -0.81 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.81 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.2 spell_power points (2.40 DPS) | yes | Green Silken Shoulders (7057, -0.10 DPS) [crafted]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.34 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.8 spell_power points (2.21 DPS) | yes | Guardian Cloak (5965, -0.85 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.96 DPS) [vendor]; Long Silken Cloak (4326, -1.78 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.2 spell_power points (3.57 DPS) | yes | Dreamweave Vest (10021, -0.18 DPS) [crafted]; Robe of Power (7054, -0.37 DPS) [crafted]; Elemental Raiment (9434, -0.81 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.18 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Windchaser Cuffs (14429, -0.16 DPS) [world_drop]; Condor Bracers (15864, -0.26 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.5 spell_power points (2.82 DPS) | yes | Black Mageweave Gloves (10003, -0.85 DPS) [crafted]; Stormcloth Gloves (10011, -0.93 DPS) [crafted]; Red Mageweave Gloves (10018, -0.97 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.0 spell_power points (2.63 DPS) | yes | Gilded Cord (254037, -0.66 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.84 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.4 spell_power points (3.20 DPS) | yes | Abomination Skin Leggings (23173, -1.11 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.12 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.37 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.15 DPS) | yes | Gilded Slippers (254001, -0.99 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.59 DPS) [dungeon]; Spidersilk Boots (4320, -1.78 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.2 spell_power points (2.00 DPS) | yes | Ring of Forlorn Spirits (2043, -0.94 DPS) [quest]; Reedknot Ring (9622, -1.08 DPS) [quest]; Minor Channeling Ring (1449, -1.11 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.18 DPS) | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Minor Channeling Ring (1449, -0.30 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.25 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (105.3 DPS) | yes | Windweaver Staff (7757, -0.92 DPS) [dungeon]; Staff of Jordan (873, -1.38 DPS) [world_drop]; Gut Ripper (2164, -4.29 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 304.8 spell_power points (40.06 DPS) | yes | Nether Force Wand (11263, -1.84 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.53 DPS) [quest]; Ragefire Wand (7513, -2.58 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 152.4. Weights run: 1.1s. Verify run: 0.9s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.153 ± 0.047, crit=0.326 ± 0.021 per rating point (14 rating = 1%, 4.568 per %), hit=0.720 ± 0.008 per rating point (10 rating = 1%, 7.203 per %), spell_haste=8.547 ± 0.996, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 42.1 spell_power points (5.37 DPS) | yes | Dreamweave Circlet (10041, -1.22 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.23 DPS) [vendor]; Chief Architect's Monocle (11839, -1.40 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 22.9 spell_power points (2.93 DPS) | yes | Scorn's Icy Choker (23169, -1.15 DPS) [dungeon]; Mindburst Medallion (11196, -1.28 DPS) [quest]; Horizon Choker (13085, -2.58 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.8 spell_power points (4.31 DPS) | yes | Kentic Amice (11624, -0.61 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.21 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.38 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.9 spell_power points (2.67 DPS) | yes | Runecloth Cloak (13860, -0.34 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.61 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.41 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.1 spell_power points (5.37 DPS) | yes | Runecloth Tunic (13857, -1.58 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.64 DPS) [vendor]; Runecloth Robe (13858, -2.36 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.3 spell_power points (2.21 DPS) | yes | Bloodband Bracers (11469, -0.25 DPS) [quest]; Nethergeld Cuffs (254061, -0.28 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.44 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 35.3 spell_power points (4.51 DPS) | yes | Raider Handwraps (272098, -0.38 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.53 DPS) [vendor]; Dreamweave Gloves (10019, -1.63 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 27.9 spell_power points (3.56 DPS) | yes | Satyrmane Sash (17755, -0.30 DPS) [dungeon]; Ban'thok Sash (11662, -0.32 DPS) [dungeon]; Deathmage Sash (10771, -0.46 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.5 spell_power points (4.41 DPS) | yes | Red Mageweave Pants (10009, -0.85 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.38 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -2.33 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 25.6 spell_power points (3.27 DPS) | yes | Earthen Silk Slippers (254013, -0.20 DPS) [crafted]; Gilded Sandals (254107, -0.54 DPS) [crafted]; Southsea Mojo Boots (20641, -0.63 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.3 spell_power points (2.21 DPS) | yes | Philanthropist's Ring (281635, -0.05 DPS) [quest]; Mindseye Circle (10634, -0.44 DPS) [dungeon]; Band of the Unicorn (7553, -0.55 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.1 spell_power points (2.18 DPS) | yes | Philanthropist's Ring (281635, -0.02 DPS) [quest]; Mindseye Circle (10634, -0.41 DPS) [dungeon]; Band of the Unicorn (7553, -0.52 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (152.4 DPS) | yes | Uther's Strength (11302, -0.06 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | sim-verified (152.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (152.4 DPS) | yes | Spellshifter Rod (9527, -0.88 DPS) [quest]; Radiant Staff (249453, -1.47 DPS) [crafted]; Blade of Eternal Darkness (17780, -1.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 411.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.86 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 390.1. Weights run: 1.3s. Verify run: 1.0s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=1.140 ± 0.094, crit=0.526 ± 0.027 per rating point (14 rating = 1%, 7.368 per %), hit=1.366 ± 0.017 per rating point (10 rating = 1%, 13.656 per %), spell_haste=not significant (-1.453 ± 1.641), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 56.9 spell_power points (11.76 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.65 DPS) [pvp]; Magister's Crown (16686, -2.41 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beads of Ogre Mojo (22149, -0.65 DPS) [quest]; Pebble of Kajaro (19600, -1.24 DPS) [quest]; Jewel of Kajaro (19601, -1.24 DPS) [quest] |
| shoulder | Darkspear Shoulderpads (272103) | Creeg Bothunk [vendor] | sim-verified (390.1 DPS) | yes | Field Marshal's Silk Spaulders (231602, -0.14 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.74 DPS) [crafted]; Rugged Mantle of the Timbermaw (227808, -7.34 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 38.8 spell_power points (8.02 DPS) | yes | Hide of the Wild (18510, -2.77 DPS) [crafted]; Crystalline Threaded Cape (20697, -2.94 DPS) [world]; Spritecaster Cape (11623, -3.71 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 61.0 spell_power points (12.63 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.27 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -2.75 DPS) [pvp]; Sorcerer's Robes (226932, -3.42 DPS) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.1 spell_power points (6.44 DPS) | yes | Sublime Wristguards (18497, -1.60 DPS) [dungeon]; Runecloth Cuffs (254123, -1.80 DPS) [crafted]; Marshal's Silk Bracers (16438, -1.96 DPS) [pvp] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 41.6 spell_power points (8.61 DPS) | yes | Sorcerer's Gloves (22066, -0.00 DPS) [quest]; Marshal's Silk Gloves (16440, -0.19 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -0.19 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 63.7 spell_power points (13.18 DPS) | yes | Belt of the Archmage (18405, -3.75 DPS) [crafted]; Magister's Belt (16685, -6.78 DPS) [dungeon]; Magician's Cord (272393, -6.83 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 58.1 spell_power points (12.02 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Sorcerer's Leggings (226933, -1.50 DPS) [quest]; Knight-Captain's Silk Legguards (227109, -2.15 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 39.2 spell_power points (8.12 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.62 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Signet Ring of the Bronze Dragonflight (21206, -1.15 DPS) [quest]; Maiden's Circle (13001, -1.77 DPS) [world_drop] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.13 DPS) [quest]; Maiden's Circle (13001, -0.75 DPS) [world_drop]; Naglering (11669, -8.61 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+14.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -1.45 DPS) [vendor]; Serenity Field (272439, -3.10 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.91 DPS) [world]; Teebu's Blazing Longsword (1728, -13.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 377.2 spell_power points (78.00 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.88 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.15 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.52 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Darkspear Shoulderpads; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Briarwood Reed; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1074, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 30.1. Weights run: 1.1s. Verify run: 0.8s. 138 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.617 ± 0.012, crit=0.140 ± 0.005 per rating point (14 rating = 1%, 1.962 per %), hit=0.389 ± 0.002 per rating point (10 rating = 1%, 3.886 per %), spell_haste=not significant (-0.587 ± 0.226), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.44 DPS) | yes | Shadow Goggles (4373, -0.48 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.6 spell_power points (0.78 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.36 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.30 DPS) | yes | Pearl-clasped Cloak (5542, -0.01 DPS) [crafted]; Feyscale Cloak (6632, -0.07 DPS) [dungeon]; Black Whelp Cloak (7283, -0.07 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.1 spell_power points (0.60 DPS) | yes | Green Woolen Robe (6243, -0.24 DPS) [crafted]; Mystic's Wrap (14369, -0.28 DPS) [world_drop]; Gray Woolen Robe (2585, -0.63 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.7 spell_power points (0.27 DPS) | yes | Mindthrust Bracers (1974, -0.05 DPS) [dungeon]; Bright Bracers (3647, -0.09 DPS) [world_drop]; Featherbead Bracers (15452, -0.18 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.52 DPS) | yes | Gnoll Casting Gloves (892, -0.07 DPS) [world]; Pristine Gloves (253913, -0.09 DPS) [crafted]; Blight Gloves (279877, -0.20 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.5 spell_power points (0.48 DPS) | yes | Keller's Girdle (2911, -0.11 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.19 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.50 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (30.1 DPS) | yes | Silk-threaded Trousers (1929, -0.20 DPS) [dungeon]; Rumpled Kilt (274741, -0.35 DPS) [vendor]; Abomination Skin Leggings (23173, -0.39 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.5 spell_power points (0.70 DPS) | yes | Pristine Boots (253889, -0.34 DPS) [crafted]; Red Woolen Boots (4313, -0.40 DPS) [crafted]; Feather Padded Treads (285345, -0.44 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.37 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Sludge-Stained Band (286535, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.7 spell_power points (0.27 DPS) | yes | Sludge-Stained Band (286535, -0.05 DPS) [world]; Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; Loop of Sacrifice (281673, -0.18 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 6.2 spell_power points (0.46 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.09 DPS) [world]; Lesser Staff of the Spire (1300, -0.18 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 304.3 spell_power points (22.53 DPS) | yes | Skycaller (12984, -0.61 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.24 DPS) [dungeon]; Sizzle Stick (8071, -4.51 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (orc, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 54.1. Weights run: 1.2s. Verify run: 0.8s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.948 ± 0.026, crit=0.225 ± 0.012 per rating point (14 rating = 1%, 3.154 per %), hit=0.419 ± 0.004 per rating point (10 rating = 1%, 4.189 per %), spell_haste=not significant (-0.185 ± 0.403), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.5 spell_power points (1.48 DPS) | yes | Nightsky Cowl (4039, -0.39 DPS) [world_drop]; Holy Shroud (2721, -0.43 DPS) [world_drop]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 spell_power points (1.21 DPS) | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.99 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.5 spell_power points (1.68 DPS) | yes | Death Speaker Mantle (6685, -0.11 DPS) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.6 spell_power points (0.72 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Resilient Cape (14400, -0.18 DPS) [world_drop]; Hillman's Cloak (3719, -0.25 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.3 spell_power points (2.04 DPS) | yes | Death Speaker Robes (6682, -0.43 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.68 DPS) [dungeon]; Pristine Gown (253961, -0.73 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.86 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.32 DPS) [quest]; Glowing Magical Bracelets (13106, -0.49 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.7 spell_power points (1.03 DPS) | yes | Truefaith Gloves (7049, -0.28 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.30 DPS) [dungeon]; Serpent Gloves (5970, -0.36 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.8 spell_power points (1.32 DPS) | yes | Crimson Silk Belt (7055, -0.12 DPS) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (54.1 DPS) | yes | Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Abomination Skin Leggings (23173, -0.90 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.6 spell_power points (1.30 DPS) | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -0.75 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.67 DPS) | yes | Snake Hoop (6750, -0.04 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 6.6 spell_power points (0.63 DPS) | yes | Snake Hoop (6750, +0.00 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.06 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.4 spell_power points (1.00 DPS) | yes | Twisted Chanter's Staff (890, -0.09 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Scorn's Focal Dagger (23168, -0.14 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 350.3 spell_power points (33.50 DPS) | yes | Starfaller (13063, -0.27 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Black Widow Band; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 96.3. Weights run: 1.2s. Verify run: 0.8s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.865 ± 0.039, crit=0.197 ± 0.014 per rating point (14 rating = 1%, 2.758 per %), hit=0.487 ± 0.006 per rating point (10 rating = 1%, 4.875 per %), spell_haste=6.790 ± 0.684, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.76 DPS) | yes | Augural Shroud (2620, -0.18 DPS) [world]; Corpseshroud (10574, -0.60 DPS) [dungeon]; Thinking Cap (2624, -0.83 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.2 spell_power points (1.60 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.47 DPS) [quest]; Triune Amulet (7722, -0.81 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.81 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.2 spell_power points (2.40 DPS) | yes | Green Silken Shoulders (7057, -0.10 DPS) [crafted]; Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Berylline Pads (4197, -0.34 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 16.8 spell_power points (2.21 DPS) | yes | Guardian Cloak (5965, -0.85 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.96 DPS) [vendor]; Long Silken Cloak (4326, -1.51 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.2 spell_power points (3.57 DPS) | yes | Dreamweave Vest (10021, -0.18 DPS) [crafted]; Robe of Power (7054, -0.37 DPS) [crafted]; Elemental Raiment (9434, -0.81 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.9 spell_power points (1.44 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Spidertank Oilrag (9448, -0.25 DPS) [dungeon]; Windchaser Cuffs (14429, -0.41 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.5 spell_power points (2.82 DPS) | yes | Black Mageweave Gloves (10003, -0.85 DPS) [crafted]; Red Mageweave Gloves (10018, -0.89 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.93 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.0 spell_power points (2.63 DPS) | yes | Gilded Cord (254037, -0.66 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.84 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.4 spell_power points (3.20 DPS) | yes | Abomination Skin Leggings (23173, -1.11 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.19 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.37 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.15 DPS) | yes | Gilded Slippers (254001, -1.44 DPS) [crafted]; Acidic Walkers (9454, -1.59 DPS) [dungeon]; Spidersilk Boots (4320, -1.78 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.2 spell_power points (2.00 DPS) | yes | Reedknot Ring (9622, -1.08 DPS) [quest]; Voodoo Band (1996, -1.20 DPS) [world]; Black Widow Band (6199, -1.20 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.18 DPS) | yes | Voodoo Band (1996, -0.39 DPS) [world]; Black Widow Band (6199, -0.39 DPS) [world]; Reedknot Ring (9622, -0.97 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (96.3 DPS) | yes | Windweaver Staff (7757, -0.92 DPS) [dungeon]; Staff of Jordan (873, -1.38 DPS) [world_drop]; Gut Ripper (2164, -3.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 304.8 spell_power points (40.06 DPS) | yes | Nether Force Wand (11263, -1.50 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.53 DPS) [quest]; Ragefire Wand (7513, -2.58 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 138.9. Weights run: 1.1s. Verify run: 0.9s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.153 ± 0.047, crit=0.326 ± 0.021 per rating point (14 rating = 1%, 4.568 per %), hit=0.720 ± 0.008 per rating point (10 rating = 1%, 7.203 per %), spell_haste=8.547 ± 0.996, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 42.1 spell_power points (5.37 DPS) | yes | Dreamweave Circlet (10041, -1.22 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.23 DPS) [vendor]; Chief Architect's Monocle (11839, -1.40 DPS) [dungeon] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 22.9 spell_power points (2.93 DPS) | yes | Scorn's Icy Choker (23169, -1.15 DPS) [dungeon]; Mindburst Medallion (11196, -1.28 DPS) [quest]; Horizon Choker (13085, -1.59 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 33.8 spell_power points (4.31 DPS) | yes | Kentic Amice (11624, -0.61 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.21 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.38 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.4 spell_power points (2.86 DPS) | yes | Spritecaster Cape (11623, -0.19 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.38 DPS) [dungeon]; Runecloth Cloak (13860, -0.53 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.1 spell_power points (5.37 DPS) | yes | Runecloth Tunic (13857, -1.58 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.64 DPS) [vendor]; Runecloth Robe (13858, -2.09 DPS, sim-verified) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.3 spell_power points (2.21 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Bloodband Bracers (11469, -0.25 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 35.3 spell_power points (4.51 DPS) | yes | Raider Handwraps (272098, -0.38 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.53 DPS) [vendor]; Dreamweave Gloves (10019, -1.63 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 27.9 spell_power points (3.56 DPS) | yes | Satyrmane Sash (17755, -0.30 DPS) [dungeon]; Ban'thok Sash (11662, -0.32 DPS) [dungeon]; Deathmage Sash (10771, -0.46 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.5 spell_power points (4.41 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.53 DPS) [vendor]; Red Mageweave Pants (10009, -0.85 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.38 DPS) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 25.6 spell_power points (3.27 DPS) | yes | Earthen Silk Slippers (254013, -0.20 DPS) [crafted]; Gilded Sandals (254107, -0.54 DPS) [crafted]; Southsea Mojo Boots (20641, -0.63 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.3 spell_power points (2.21 DPS) | yes | Philanthropist's Ring (281635, -0.05 DPS) [quest]; Mindseye Circle (10634, -0.44 DPS) [dungeon]; Band of the Unicorn (7553, -0.55 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.1 spell_power points (2.18 DPS) | yes | Philanthropist's Ring (281635, -0.02 DPS) [quest]; Mindseye Circle (10634, -0.41 DPS) [dungeon]; Band of the Unicorn (7553, -0.52 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (138.9 DPS) | yes | Uther's Strength (11302, -0.06 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (138.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (138.9 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellshifter Rod (9527, -0.88 DPS) [quest]; Radiant Staff (249453, -1.47 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 411.0 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.86 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 370.0. Weights run: 1.3s. Verify run: 0.9s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=1.140 ± 0.094, crit=0.526 ± 0.027 per rating point (14 rating = 1%, 7.368 per %), hit=1.366 ± 0.017 per rating point (10 rating = 1%, 13.656 per %), spell_haste=not significant (-1.453 ± 1.641), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 56.9 spell_power points (11.76 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.65 DPS) [pvp]; Magister's Crown (16686, -10.81 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (370.0 DPS) | yes | Beads of Ogre Mojo (22149, -0.65 DPS) [quest]; Pebble of Kajaro (19600, -1.24 DPS) [quest]; Jewel of Kajaro (19601, -1.24 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 51.5 spell_power points (10.64 DPS) | yes | Darkspear Shoulderpads (272103, -1.80 DPS) [vendor]; Warlord's Silk Amice (231594, -1.94 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.54 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 38.8 spell_power points (8.02 DPS) | yes | Hide of the Wild (18510, -2.77 DPS) [crafted]; Crystalline Threaded Cape (20697, -2.94 DPS) [world]; Deep Woodlands Cloak (19121, -3.42 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 61.0 spell_power points (12.63 DPS) | yes | Warlord's Silk Raiment (231596, -0.27 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -2.75 DPS) [pvp]; Sorcerer's Robes (226932, -3.42 DPS) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.1 spell_power points (6.44 DPS) | yes | Sublime Wristguards (18497, -1.60 DPS) [dungeon]; Runecloth Cuffs (254123, -1.80 DPS) [crafted]; General's Silk Cuffs (16538, -1.96 DPS) [pvp] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 41.6 spell_power points (8.61 DPS) | yes | Sorcerer's Gloves (22066, -0.00 DPS) [quest]; General's Silk Handguards (16540, -0.19 DPS) [vendor]; General's Silk Gauntlets (231599, -0.19 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 63.7 spell_power points (13.18 DPS) | yes | Belt of the Archmage (18405, -3.75 DPS) [crafted]; Magister's Belt (16685, -6.78 DPS) [dungeon]; Magician's Cord (272393, -11.29 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 58.1 spell_power points (12.02 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -1.75 DPS) [rep]; Sorcerer's Leggings (226933, -14.02 DPS, sim-verified) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 39.2 spell_power points (8.12 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.62 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (370.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -1.15 DPS) [quest]; Maiden's Circle (13001, -1.77 DPS) [world_drop]; Naglering (11669, -13.02 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (370.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21206, -0.13 DPS) [quest]; Maiden's Circle (13001, -0.75 DPS) [world_drop]; Naglering (11669, -12.69 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (370.0 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Second Wind (11819, -12.68 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (370.0 DPS) | yes | Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -1.45 DPS) [vendor]; Serenity Field (272439, -3.10 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (370.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.31 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -17.30 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 377.2 spell_power points (78.00 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.88 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.15 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.52 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Briarwood Reed; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

