# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 32.9. Weights run: 1.1s. Verify run: 0.7s. 149 eligible items had no known source.

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

### Band 30 (gnome, 000000000000000000-23552100030000000-0000000000000000000)

Set DPS (verified): 59.8. Weights run: 1.2s. Verify run: 0.8s. 248 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.076 ± 0.029, crit=0.232 ± 0.012 per rating point (14 rating = 1%, 3.248 per %), hit=0.436 ± 0.004 per rating point (10 rating = 1%, 4.361 per %), spell_haste=not significant (0.788 ± 0.519), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 16.8 spell_power points (1.59 DPS) | yes | Nightsky Cowl (4039, -0.37 DPS) [world_drop]; Shadow Hood (4323, -0.47 DPS) [crafted]; Resilient Cap (14401, -0.47 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.5 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -0.87 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.87 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.00 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.7 spell_power points (1.78 DPS) | yes | Death Speaker Mantle (6685, -0.08 DPS) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 8.6 spell_power points (0.82 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Repairman's Cape (9605, -0.12 DPS) [quest]; Resilient Cape (14400, -0.20 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.0 spell_power points (2.19 DPS) | yes | Mechbuilder's Overalls (9508, -0.65 DPS) [dungeon]; Death Speaker Robes (6682, -0.68 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.80 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.86 DPS) | yes | Nightsky Wristbands (6407, -0.24 DPS) [world_drop]; Stonecloth Bindings (14416, -0.34 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.03 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 15.8 spell_power points (1.51 DPS) | yes | Truefaith Gloves (7049, -0.72 DPS) [crafted]; Pristine Gloves (253913, -0.82 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -1.01 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 14.2 spell_power points (1.35 DPS) | yes | Crimson Silk Belt (7055, -0.07 DPS) [crafted]; Invoker's Cord (215366, -0.18 DPS) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (59.8 DPS) | yes | Filigreed Pristine Leggings (253937, -0.20 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.24 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.88 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 14.5 spell_power points (1.38 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -0.81 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 7.5 spell_power points (0.72 DPS) | yes | Minor Channeling Ring (1449, -0.04 DPS) [quest]; Lorekeeper's Ring (19525, -0.05 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.5 spell_power points (0.72 DPS) | yes | Minor Channeling Ring (1449, -0.04 DPS) [quest]; Lorekeeper's Ring (19525, -0.05 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 11.8 spell_power points (1.13 DPS) | yes | Twisted Chanter's Staff (890, -0.10 DPS) [world_drop]; Scorn's Focal Dagger (23168, -0.27 DPS) [dungeon]; Channeler's Staff (4437, -0.31 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 352.4 spell_power points (33.50 DPS) | yes | Starfaller (13063, -0.23 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape

### Band 40 (gnome, 000000000000000000-23552100030023050-0000000000000000000)

Set DPS (verified): 99.1. Weights run: 1.2s. Verify run: 0.8s. 330 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.014 ± 0.045, crit=0.334 ± 0.026 per rating point (14 rating = 1%, 4.682 per %), hit=0.548 ± 0.006 per rating point (10 rating = 1%, 5.484 per %), spell_haste=not significant (0.133 ± 0.753), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 21.1 spell_power points (2.48 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.02 DPS) [crafted]; Corpseshroud (10574, -0.22 DPS) [dungeon]; Thinking Cap (2624, -0.46 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.1 spell_power points (1.54 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.35 DPS) [quest]; Triune Amulet (7722, -0.70 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.70 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.2 spell_power points (2.37 DPS) | yes | Green Silken Shoulders (7057, -0.12 DPS) [crafted]; Bloodmage Mantle (7684, -0.24 DPS) [dungeon]; Death Speaker Mantle (6685, -0.36 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 18.1 spell_power points (2.13 DPS) | yes | Long Silken Cloak (4326, -0.83 DPS) [crafted]; Guardian Cloak (5965, -0.83 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.58 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.1 spell_power points (3.30 DPS) | yes | Dreamweave Vest (10021, -0.11 DPS) [crafted]; Robe of Power (7054, -0.22 DPS) [crafted]; Green Silk Armor (7065, -0.69 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 9.1 spell_power points (1.07 DPS) | yes | Arcane Runed Bracers (4744, -0.02 DPS) [quest]; Spidertank Oilrag (9448, -0.02 DPS) [dungeon]; Mistscape Bracers (4045, -0.12 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.1 spell_power points (2.59 DPS) | yes | Red Mageweave Gloves (10018, -0.11 DPS) [crafted]; Stormcloth Gloves (10011, -0.69 DPS) [crafted]; Town Clerk's Mittens (270029, -0.81 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.2 spell_power points (2.61 DPS) | yes | Gilded Cord (254037, -0.72 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.96 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.2 spell_power points (3.07 DPS) | yes | Crimson Silk Pantaloons (7062, -0.59 DPS) [crafted]; Abomination Skin Leggings (23173, -1.06 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.30 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.82 DPS) | yes | Gilded Slippers (254001, -1.16 DPS) [crafted]; Acidic Walkers (9454, -1.28 DPS) [dungeon]; Spidersilk Boots (4320, -1.52 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (1.89 DPS) | yes | Ring of Forlorn Spirits (2043, -0.95 DPS) [quest]; Voodoo Band (1996, -1.06 DPS) [world]; Black Widow Band (6199, -1.06 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.06 DPS) | yes | Voodoo Band (1996, -0.22 DPS) [world]; Black Widow Band (6199, -0.22 DPS) [world]; Ring of Forlorn Spirits (2043, -1.03 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (99.1 DPS) | yes | Windweaver Staff (7757, -0.56 DPS) [dungeon]; Staff of Jordan (873, -1.04 DPS) [world_drop]; Gut Ripper (2164, -4.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 341.2 spell_power points (40.09 DPS) | yes | Nether Force Wand (11263, -2.41 DPS) [quest]; Icefury Wand (7514, -2.55 DPS) [quest]; Ragefire Wand (7513, -2.60 DPS) [quest] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 50 (gnome, 205011000000000000-23552100030023051-0000000000000000000)

Set DPS (verified): 170.0. Weights run: 1.2s. Verify run: 0.9s. 424 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.907 ± 0.053, crit=0.321 ± 0.030 per rating point (14 rating = 1%, 4.488 per %), hit=0.749 ± 0.010 per rating point (10 rating = 1%, 7.495 per %), spell_haste=6.546 ± 1.047, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.1 spell_power points (5.43 DPS) | yes | Dreamweave Circlet (10041, -1.03 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.14 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.48 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.4 spell_power points (3.14 DPS) | yes | Horizon Choker (13085, -1.28 DPS) [world_drop]; Scorn's Icy Choker (23169, -1.32 DPS) [dungeon]; Mindburst Medallion (11196, -1.46 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.3 spell_power points (4.29 DPS) | yes | Kentic Amice (11624, -0.52 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.27 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.28 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.4 spell_power points (2.84 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.33 DPS) [dungeon]; Runecloth Cloak (13860, -0.47 DPS) [crafted]; Big Voodoo Cloak (8216, -0.92 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.1 spell_power points (5.43 DPS) | yes | Robe of the Magi (1716, -1.42 DPS) [world_drop]; Runecloth Tunic (13857, -1.49 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.56 DPS) [vendor] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.6 spell_power points (1.99 DPS) | yes | Nethergeld Cuffs (254061, -0.04 DPS) [crafted]; Bloodband Bracers (11469, -0.06 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -0.40 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 32.2 spell_power points (4.71 DPS) | yes | Raider Handwraps (272098, -0.77 DPS) [vendor]; Dreamweave Gloves (10019, -1.55 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.61 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.2 spell_power points (3.40 DPS) | yes | Satyrmane Sash (17755, -0.02 DPS) [dungeon]; Ban'thok Sash (11662, -0.07 DPS) [dungeon]; Deathmage Sash (10771, -0.38 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.1 spell_power points (4.69 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.69 DPS) [vendor]; Red Mageweave Pants (10009, -1.05 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.80 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.51 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.05 DPS) [vendor]; Gilded Sandals (254107, -0.71 DPS) [crafted]; Southsea Mojo Boots (20641, -0.88 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (2.26 DPS) | yes | Brainlash (6440, -0.27 DPS) [dungeon]; Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.50 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.3 spell_power points (2.25 DPS) | yes | Brainlash (6440, -0.26 DPS) [dungeon]; Band of the Unicorn (7553, -0.34 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.49 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (170.0 DPS) | yes | Uther's Strength (11302, -0.11 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | sim-verified (170.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (170.0 DPS) | yes | Spellshifter Rod (9527, -0.80 DPS) [quest]; Spellforce Rod (1664, -0.92 DPS) [world]; Blade of Eternal Darkness (17780, -2.89 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 358.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands

### Band 60 (gnome, 205015100000000000-23552100030023051-0050000000000000000)

Set DPS (verified): 416.0. Weights run: 1.2s. Verify run: 0.9s. 1074 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=1.190 ± 0.077, crit=0.576 ± 0.043 per rating point (14 rating = 1%, 8.061 per %), hit=1.110 ± 0.016 per rating point (10 rating = 1%, 11.098 per %), spell_haste=not significant (1.489 ± 1.547), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 58.8 spell_power points (13.81 DPS) | yes | Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Cowl (227103, -1.96 DPS) [pvp]; Magister's Crown (16686, -2.84 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (416.0 DPS) | yes | Beads of Ogre Mojo (22149, -0.75 DPS) [quest]; Pebble of Kajaro (19600, -1.41 DPS) [quest]; Jewel of Kajaro (19601, -1.41 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 52.9 spell_power points (12.42 DPS) | yes | Darkspear Shoulderpads (272103, -2.06 DPS) [vendor]; Field Marshal's Silk Spaulders (231602, -2.36 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.91 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.6 spell_power points (8.60 DPS) | yes | Hide of the Wild (18510, -2.52 DPS) [crafted]; Crystalline Threaded Cape (20697, -2.78 DPS) [world]; Spritecaster Cape (11623, -3.63 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 62.3 spell_power points (14.63 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.25 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.06 DPS) [pvp]; Sorcerer's Robes (226932, -3.89 DPS) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 31.5 spell_power points (7.40 DPS) | yes | Sublime Wristguards (18497, -1.79 DPS) [dungeon]; Runecloth Cuffs (254123, -2.02 DPS) [crafted]; Marshal's Silk Bracers (16438, -2.09 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 40.1 spell_power points (9.42 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 62.3 spell_power points (14.62 DPS) | yes | Belt of the Archmage (18405, -3.56 DPS) [crafted]; Magister's Belt (16685, -7.11 DPS) [dungeon]; Magician's Cord (272393, -12.14 DPS, sim-verified) [vendor] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 60.0 spell_power points (14.09 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -2.52 DPS) [pvp]; Sorcerer's Leggings (226933, -2.55 DPS) [quest] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 40.0 spell_power points (9.40 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.70 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (416.0 DPS) | yes | Channeler's Ring (272406, -1.72 DPS) [vendor]; Maiden's Circle (13001, -2.06 DPS) [world_drop]; Naglering (11669, -18.11 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (416.0 DPS) | yes | Channeler's Ring (272406, -0.37 DPS) [vendor]; Maiden's Circle (13001, -0.70 DPS) [world_drop]; Naglering (11669, -17.35 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (416.0 DPS) | yes | Darkmoon Card: Blue Dragon (19288, +0.00 DPS) [quest]; Weakness Analyzer (272438, -1.64 DPS) [vendor]; Serenity Field (272439, -3.52 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (416.0 DPS) | yes | Darkmoon Card: Blue Dragon (19288, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (416.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -2.36 DPS) [world]; Teebu's Blazing Longsword (1728, -21.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 333.0 spell_power points (78.17 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.74 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.84 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.15 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Torch of Light

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

### Band 30 (orc, 000000000000000000-23552100030000000-0000000000000000000)

Set DPS (verified): 54.9. Weights run: 1.2s. Verify run: 0.8s. 230 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.076 ± 0.029, crit=0.232 ± 0.012 per rating point (14 rating = 1%, 3.248 per %), hit=0.436 ± 0.004 per rating point (10 rating = 1%, 4.361 per %), spell_haste=not significant (0.788 ± 0.519), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 16.8 spell_power points (1.59 DPS) | yes | Nightsky Cowl (4039, -0.37 DPS) [world_drop]; Shadow Hood (4323, -0.47 DPS) [crafted]; Resilient Cap (14401, -0.47 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.5 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -0.87 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.87 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.89 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.7 spell_power points (1.78 DPS) | yes | Death Speaker Mantle (6685, -0.08 DPS) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 8.6 spell_power points (0.82 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Resilient Cape (14400, -0.20 DPS) [world_drop]; Soft Willow Cape (16661, -0.31 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.0 spell_power points (2.19 DPS) | yes | Death Speaker Robes (6682, -0.39 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.65 DPS) [dungeon]; Pristine Gown (253961, -0.80 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.86 DPS) | yes | Nightsky Wristbands (6407, -0.24 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.24 DPS) [quest]; Glowing Magical Bracelets (13106, -1.00 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 11.4 spell_power points (1.08 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.26 DPS) [dungeon]; Truefaith Gloves (7049, -0.30 DPS) [crafted]; Blight Gloves (279877, -0.37 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 14.2 spell_power points (1.35 DPS) | yes | Crimson Silk Belt (7055, -0.07 DPS) [crafted]; Invoker's Cord (215366, -0.18 DPS) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (54.9 DPS) | yes | Filigreed Pristine Leggings (253937, -0.20 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.24 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.82 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 14.5 spell_power points (1.38 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -0.81 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 7.5 spell_power points (0.72 DPS) | yes | Advisor's Ring (19521, -0.05 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.15 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.5 spell_power points (0.72 DPS) | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.15 DPS) [vendor]; Advisor's Ring (19521, -0.59 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 11.8 spell_power points (1.13 DPS) | yes | Twisted Chanter's Staff (890, -0.10 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.10 DPS) [quest]; Scorn's Focal Dagger (23168, -0.27 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 352.4 spell_power points (33.50 DPS) | yes | Starfaller (13063, -0.23 DPS) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552100030023050-0000000000000000000)

Set DPS (verified): 90.6. Weights run: 1.2s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.014 ± 0.045, crit=0.334 ± 0.026 per rating point (14 rating = 1%, 4.682 per %), hit=0.548 ± 0.006 per rating point (10 rating = 1%, 5.484 per %), spell_haste=not significant (0.133 ± 0.753), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 21.1 spell_power points (2.48 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.02 DPS) [crafted]; Corpseshroud (10574, -0.22 DPS) [dungeon]; Thinking Cap (2624, -0.46 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.1 spell_power points (1.54 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.35 DPS) [quest]; Triune Amulet (7722, -0.70 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.70 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.2 spell_power points (2.37 DPS) | yes | Green Silken Shoulders (7057, -0.12 DPS) [crafted]; Bloodmage Mantle (7684, -0.24 DPS) [dungeon]; Death Speaker Mantle (6685, -0.36 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 18.1 spell_power points (2.13 DPS) | yes | Darkspear Raider's Cloak (272077, -0.82 DPS) [vendor]; Long Silken Cloak (4326, -0.83 DPS) [crafted]; Guardian Cloak (5965, -0.83 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.1 spell_power points (3.30 DPS) | yes | Dreamweave Vest (10021, -0.11 DPS) [crafted]; Robe of Power (7054, -0.22 DPS) [crafted]; Green Silk Armor (7065, -0.69 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 12.1 spell_power points (1.42 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Windchaser Cuffs (14429, -0.35 DPS) [world_drop]; Spidertank Oilrag (9448, -0.37 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.1 spell_power points (2.59 DPS) | yes | Red Mageweave Gloves (10018, -0.11 DPS) [crafted]; Stormcloth Gloves (10011, -0.69 DPS) [crafted]; Gilded Handwraps (254021, -0.82 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.2 spell_power points (2.61 DPS) | yes | Defiler's Cloth Girdle (20166, -0.49 DPS) [rep]; Gilded Cord (254037, -0.72 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.2 spell_power points (3.07 DPS) | yes | Crimson Silk Pantaloons (7062, -0.59 DPS) [crafted]; Abomination Skin Leggings (23173, -1.06 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.30 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.82 DPS) | yes | Gilded Slippers (254001, -1.16 DPS) [crafted]; Acidic Walkers (9454, -1.28 DPS) [dungeon]; Spidersilk Boots (4320, -1.52 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (1.89 DPS) | yes | Ogremind Ring (1993, -1.06 DPS) [world_drop]; Voodoo Band (1996, -1.06 DPS) [world]; Black Widow Band (6199, -1.06 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.06 DPS) | yes | Ogremind Ring (1993, -0.22 DPS) [world_drop]; Voodoo Band (1996, -0.22 DPS) [world]; Black Widow Band (6199, -1.53 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-verified (90.6 DPS) | yes | Windweaver Staff (7757, -0.56 DPS) [dungeon]; Staff of Jordan (873, -1.04 DPS) [world_drop]; Gut Ripper (2164, -3.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 341.2 spell_power points (40.09 DPS) | yes | Nether Force Wand (11263, -1.12 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.55 DPS) [quest]; Ragefire Wand (7513, -2.60 DPS) [quest] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 205011000000000000-23552100030023051-0000000000000000000)

Set DPS (verified): 156.5. Weights run: 1.2s. Verify run: 0.9s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.907 ± 0.053, crit=0.321 ± 0.030 per rating point (14 rating = 1%, 4.488 per %), hit=0.749 ± 0.010 per rating point (10 rating = 1%, 7.495 per %), spell_haste=6.546 ± 1.047, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 37.1 spell_power points (5.43 DPS) | yes | Dreamweave Circlet (10041, -1.03 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.14 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.48 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 21.4 spell_power points (3.14 DPS) | yes | Scorn's Icy Choker (23169, -1.32 DPS) [dungeon]; Mindburst Medallion (11196, -1.46 DPS) [quest]; Horizon Choker (13085, -2.03 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 29.3 spell_power points (4.29 DPS) | yes | Kentic Amice (11624, -0.52 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.27 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.28 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.2 spell_power points (2.95 DPS) | yes | Spritecaster Cape (11623, -0.11 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.44 DPS) [dungeon]; Runecloth Cloak (13860, -0.57 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 37.1 spell_power points (5.43 DPS) | yes | Runecloth Tunic (13857, -1.49 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.56 DPS) [vendor]; Robe of the Magi (1716, -2.36 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 13.6 spell_power points (1.99 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Nethergeld Cuffs (254061, -0.04 DPS) [crafted] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 32.2 spell_power points (4.71 DPS) | yes | Dreamweave Gloves (10019, -1.55 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.61 DPS) [vendor]; Raider Handwraps (272098, -2.10 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 23.2 spell_power points (3.40 DPS) | yes | Satyrmane Sash (17755, -0.02 DPS) [dungeon]; Ban'thok Sash (11662, -0.07 DPS) [dungeon]; Deathmage Sash (10771, -0.38 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.1 spell_power points (4.69 DPS) | yes | Red Mageweave Pants (10009, -1.05 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.80 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.22 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.51 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.05 DPS) [vendor]; Gilded Sandals (254107, -0.71 DPS) [crafted]; Southsea Mojo Boots (20641, -0.88 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (2.26 DPS) | yes | Brainlash (6440, -0.27 DPS) [dungeon]; Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Advisor's Ring (19519, -0.50 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.3 spell_power points (2.25 DPS) | yes | Brainlash (6440, -0.26 DPS) [dungeon]; Band of the Unicorn (7553, -0.34 DPS) [world_drop]; Advisor's Ring (19519, -0.49 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (156.5 DPS) | yes | Uther's Strength (11302, -0.11 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (156.5 DPS) | yes | Spellshifter Rod (9527, -0.80 DPS) [quest]; Spellforce Rod (1664, -0.92 DPS) [world]; Blade of Eternal Darkness (17780, -2.76 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 358.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.77 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Flaming Incinerator (9483, -5.28 DPS) [dungeon] |

**New at 50:** head: Red Mageweave Headband; neck: Arcane Crystal Pendant; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 205015100000000000-23552100030023051-0050000000000000000)

Set DPS (verified): 387.4. Weights run: 1.2s. Verify run: 0.9s. 1062 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=1.190 ± 0.077, crit=0.576 ± 0.043 per rating point (14 rating = 1%, 8.061 per %), hit=1.110 ± 0.016 per rating point (10 rating = 1%, 11.098 per %), spell_haste=not significant (1.489 ± 1.547), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sorcerer's Crown (226935) | Saving the Best for Last [quest] | 58.8 spell_power points (13.81 DPS) | yes | Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Champion's Silk Cowl (227105, -1.96 DPS) [pvp]; Magister's Crown (16686, -2.84 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (387.4 DPS) | yes | Beads of Ogre Mojo (22149, -0.75 DPS) [quest]; Pebble of Kajaro (19600, -1.41 DPS) [quest]; Jewel of Kajaro (19601, -1.41 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 52.9 spell_power points (12.42 DPS) | yes | Darkspear Shoulderpads (272103, -2.06 DPS) [vendor]; Warlord's Silk Amice (231594, -2.36 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.91 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.6 spell_power points (8.60 DPS) | yes | Hide of the Wild (18510, -2.52 DPS) [crafted]; Crystalline Threaded Cape (20697, -2.78 DPS) [world]; Deep Woodlands Cloak (19121, -3.26 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 62.3 spell_power points (14.63 DPS) | yes | Warlord's Silk Raiment (231596, -0.25 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.06 DPS) [pvp]; Sorcerer's Robes (226932, -3.89 DPS) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 31.5 spell_power points (7.40 DPS) | yes | Sublime Wristguards (18497, -1.79 DPS) [dungeon]; Runecloth Cuffs (254123, -2.02 DPS) [crafted]; General's Silk Cuffs (16538, -2.09 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 40.1 spell_power points (9.42 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Sorcerer's Gloves (22066, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 62.3 spell_power points (14.62 DPS) | yes | Magician's Cord (272393, -3.45 DPS) [vendor]; Belt of the Archmage (18405, -3.56 DPS) [crafted]; Magister's Belt (16685, -7.11 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 60.0 spell_power points (14.09 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.21 DPS) [rep]; Legionnaire's Silk Legguards (227107, -2.52 DPS) [pvp] |
| feet | Sorcerer's Boots (22064) (or Sorcerer's Sandals (226931)) | Anthion's Parting Words [quest] | 40.0 spell_power points (9.40 DPS) | yes | Sorcerer's Sandals (226931, +0.00 DPS) [vendor]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.70 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (387.4 DPS) | yes | Channeler's Ring (272406, -1.72 DPS) [vendor]; Maiden's Circle (13001, -2.06 DPS) [world_drop]; Naglering (11669, -11.84 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (387.4 DPS) | yes | Channeler's Ring (272406, -0.37 DPS) [vendor]; Maiden's Circle (13001, -0.70 DPS) [world_drop]; Naglering (11669, -12.54 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (387.4 DPS) | yes | Weakness Analyzer (272438, -1.64 DPS) [vendor]; Serenity Field (272439, -3.52 DPS) [vendor]; Burst of Knowledge (11832, -3.99 DPS) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (387.4 DPS) | yes | Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (387.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Trindlehaven Staff (13161, -0.22 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -17.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 333.0 spell_power points (78.17 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.74 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.84 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.15 DPS) [world] |

**New at 60:** head: Sorcerer's Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Sorcerer's Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Amethyst War Staff; ranged: Torch of Light

No-known-source sample (15 of 1062, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

