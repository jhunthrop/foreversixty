# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 33.0. Weights run: 0.8s. Verify run: 0.6s. 147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.965 ± 0.012, crit=0.069 ± 0.006 per rating point (14 rating = 1%, 0.972 per %), hit=0.183 ± 0.003 per rating point (10 rating = 1%, 1.826 per %), spell_haste=-0.708 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.45 DPS) | yes | Shadow Goggles (4373, -0.99 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.7 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.39 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.9 spell_power points (0.37 DPS) | yes | Sanguine Cape (14376, -0.08 DPS) [world_drop]; Caretaker's Cape (20428, -0.14 DPS) [rep]; Heavy Woolen Cloak (4311, -0.44 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.8 spell_power points (0.74 DPS) | yes | Mystic's Wrap (14369, -0.23 DPS) [world_drop]; Mystic's Robe (14371, -0.23 DPS) [world_drop]; Gray Woolen Robe (2585, -0.57 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 4.8 spell_power points (0.36 DPS) | yes | Mystic's Bracelets (14366, -0.22 DPS) [world_drop]; Repurposed Hair Band (281256, -0.22 DPS) [quest]; Bright Bracers (3647, -0.65 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.53 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.08 DPS) [world]; Tomb Robber's Gloves (280096, -0.09 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.9 spell_power points (0.59 DPS) | yes | Novice Arcanist's Sash (253885, -0.07 DPS) [crafted]; Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.58 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.7 spell_power points (1.26 DPS) | yes | Filigreed Pristine Leggings (253937, +0.21 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.73 DPS) [dungeon]; Darkweave Breeches (12987, -0.75 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.21 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.44 DPS) [world]; Red Woolen Boots (4313, -0.52 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 spell_power points (0.52 DPS) | yes | Sludge-Stained Band (286535, -0.30 DPS) [world]; Volcanic Rock Ring (12053, -0.30 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.07 DPS, sim-verified) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (33.0 DPS) | yes | Sludge-Stained Band (286535, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.44 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 9.6 spell_power points (0.73 DPS) | yes | Lesser Staff of the Spire (1300, -0.29 DPS) [world_drop]; Staff of Westfall (2042, -0.36 DPS) [quest]; Channeler's Staff (4437, -0.56 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 299.5 spell_power points (22.54 DPS) | yes | Skycaller (12984, -0.93 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 147, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 58.2. Weights run: 1.0s. Verify run: 0.6s. 246 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.146 ± 0.023, crit=0.181 ± 0.016 per rating point (14 rating = 1%, 2.538 per %), hit=0.284 ± 0.005 per rating point (10 rating = 1%, 2.838 per %), spell_haste=not significant (0.015 ± 0.231), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 17.5 spell_power points (1.61 DPS) | yes | Shadow Hood (4323, -0.45 DPS) [crafted]; Resilient Cap (14401, -0.45 DPS) [world_drop]; Nightsky Cowl (4039, -0.50 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.9 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.17 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.3 spell_power points (1.78 DPS) | yes | Death Speaker Mantle (6685, -0.19 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.37 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.2 spell_power points (0.84 DPS) | yes | Darkspear Raider's Cloak (272078, +0.09 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.15 DPS) [quest]; Resilient Cape (14400, -0.21 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.9 spell_power points (2.20 DPS) | yes | Death Speaker Robes (6682, -0.39 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.62 DPS) [dungeon]; Pristine Gown (253961, -0.82 DPS) [crafted] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 9.2 spell_power points (0.84 DPS) | yes | Spidertank Oilrag (9448, +0.57 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -0.21 DPS) [world_drop]; Stonecloth Bindings (14416, -0.32 DPS) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 16.6 spell_power points (1.53 DPS) | yes | Truefaith Gloves (7049, -0.75 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.79 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.84 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 14.4 spell_power points (1.33 DPS) | yes | Invoker's Cord (215366, -0.16 DPS) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Crimson Silk Belt (7055, -0.28 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (58.2 DPS) | yes | Filigreed Pristine Leggings (253937, -0.20 DPS) [crafted]; Necromancer Leggings (2277, -0.22 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.68 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 spell_power points (1.38 DPS) | yes | Spidersilk Boots (4320, -0.32 DPS) [crafted]; Frothing Slippers (254003, -0.65 DPS) [crafted]; Acidic Walkers (9454, -1.12 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) (or Black Widow Band (6199)) | Willix the Importer [quest] | 8.0 spell_power points (0.74 DPS) | yes | Minor Channeling Ring (1449, -0.07 DPS) [quest]; Lorekeeper's Ring (19525, -0.09 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) | Leech Widow [world] | 8.0 spell_power points (0.74 DPS) | yes | Minor Channeling Ring (1449, -0.09 DPS, sim-verified) [quest]; Lorekeeper's Ring (19525, -0.09 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 12.6 spell_power points (1.16 DPS) | yes | Twisted Chanter's Staff (890, -0.20 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.32 DPS) [world]; Scorn's Focal Dagger (23168, -0.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 363.3 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.11 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Greater Mystic Wand (11290, -4.49 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Glowing Magical Bracelets; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Black Widow Band; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 246, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 99.7. Weights run: 1.1s. Verify run: 0.6s. 329 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.993 ± 0.029, crit=0.199 ± 0.020 per rating point (14 rating = 1%, 2.791 per %), hit=0.387 ± 0.007 per rating point (10 rating = 1%, 3.869 per %), spell_haste=6.535 ± 0.548, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.64 DPS) | yes | Augural Shroud (2620, +0.80 DPS, sim-verified) [world]; Corpseshroud (10574, -0.27 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.52 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.0 spell_power points (1.63 DPS) | yes | Necklace of Calisea (1714, -0.76 DPS) [world_drop]; Triune Amulet (7722, -0.76 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.51 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.9 spell_power points (2.51 DPS) | yes | Green Silken Shoulders (7057, +0.06 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.25 DPS) [dungeon]; Berylline Pads (4197, -0.37 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -2.72 DPS, sim-verified) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.13 DPS) [crafted]; Crimson Silk Vest (7058, -0.63 DPS) [crafted]; Robe of the Magi (1716, -1.65 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Windchaser Cuffs (14429, -0.01 DPS) [world_drop]; Mistscape Bracers (4045, -0.13 DPS) [world_drop]; Arcane Runed Bracers (4744, -1.11 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 spell_power points (2.76 DPS) | yes | Red Mageweave Gloves (10018, +0.52 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.76 DPS) [crafted]; Black Mageweave Gloves (10003, -0.88 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.9 spell_power points (2.75 DPS) | yes | Highlander's Cloth Girdle (20098, -0.52 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.75 DPS) [crafted]; Highlander's Cloth Girdle (20099, -1.00 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.9 spell_power points (3.26 DPS) | yes | Crimson Silk Pantaloons (7062, -0.65 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.02 DPS) | yes | Gilded Slippers (254001, -0.38 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.39 DPS) [dungeon]; Spidersilk Boots (4320, -1.64 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.0 spell_power points (2.01 DPS) | yes | Ring of Forlorn Spirits (2043, -1.00 DPS) [quest]; Reedknot Ring (9622, -1.13 DPS) [quest]; Minor Channeling Ring (1449, -1.13 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.13 DPS) | yes | Reedknot Ring (9622, -0.25 DPS) [quest]; Lorekeeper's Ring (19525, -0.25 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.41 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -0.64 DPS) [dungeon]; Staff of Jordan (873, -1.14 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 319.0 spell_power points (40.13 DPS) | yes | Nether Force Wand (11263, -1.25 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.56 DPS) [quest]; Ragefire Wand (7513, -2.61 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Spidertank Oilrag; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 329, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 139.1. Weights run: 0.6s. Verify run: 0.7s. 426 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.186 ± 0.069, crit=0.346 ± 0.031 per rating point (14 rating = 1%, 4.846 per %), hit=0.739 ± 0.011 per rating point (10 rating = 1%, 7.393 per %), spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 42.7 spell_power points (5.42 DPS) | yes | Knight-Lieutenant's Dreadweave Hat (220889, -0.70 DPS, sim-verified) [vendor]; Dreamweave Circlet (10041, -1.25 DPS) [crafted]; Chief Architect's Monocle (11839, -1.36 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.32 DPS) [dungeon]; Mindburst Medallion (11196, -0.44 DPS) [quest]; Arcane Crystal Pendant (20037, -3.05 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 34.3 spell_power points (4.36 DPS) | yes | Red Mageweave Shoulders (10029, -1.21 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.37 DPS) [vendor]; Kentic Amice (11624, -6.96 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.1 spell_power points (2.68 DPS) | yes | Runecloth Cloak (13860, -0.33 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.57 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -3.10 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.7 spell_power points (5.42 DPS) | yes | Runecloth Tunic (13857, -1.61 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.63 DPS) [vendor]; Runecloth Robe (13858, -1.89 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Nethergeld Cuffs (254061, -0.05 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.18 DPS) [quest]; Aristocratic Cuffs (12546, -1.84 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 36.0 spell_power points (4.57 DPS) | yes | Raider Handwraps (272098, +0.22 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.56 DPS) [vendor]; Red Mageweave Gloves (10018, -1.67 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 28.5 spell_power points (3.62 DPS) | yes | Satyrmane Sash (17755, +0.23 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.35 DPS) [dungeon]; Deathmage Sash (10771, -0.47 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.36 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -0.85 DPS) [dungeon]; Spellshock Leggings (9484, -5.57 DPS, sim-verified) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 26.1 spell_power points (3.31 DPS) | yes | Earthen Silk Slippers (254013, +0.84 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -0.56 DPS) [crafted]; Southsea Mojo Boots (20641, -0.64 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.8 spell_power points (2.26 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -0.45 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindseye Circle (10634, -0.37 DPS) [dungeon]; Band of the Unicorn (7553, -0.52 DPS) [world_drop]; Cyclopean Band (11824, -2.47 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Blessed Prayer Beads (19990, -0.99 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Spellshifter Rod (9527, -0.90 DPS) [quest]; Radiant Staff (249453, -1.50 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -23.15 DPS, sim-verified) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 438.9. Weights run: 1.0s. Verify run: 0.7s. 1020 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.011, intellect=1.454 ± 0.084, crit=0.475 ± 0.039 per rating point (14 rating = 1%, 6.656 per %), hit=1.212 ± 0.029 per rating point (10 rating = 1%, 12.117 per %), spell_haste=13.681 ± 1.123, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 94.3 spell_power points (19.34 DPS) | yes | Fireleaf Hood (240048, -2.90 DPS, sim-verified) [vendor]; Field Marshal's Coronet (16441, -6.14 DPS) [vendor]; Field Marshal's Coronet (231604, -6.14 DPS) [pvp] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (438.9 DPS) | yes | Jewel of Kajaro (19601, -0.58 DPS, sim-verified) [quest]; Beads of Ogre Mojo (22149, -0.71 DPS) [quest]; Pebble of Kajaro (19600, -1.23 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 69.5 spell_power points (14.25 DPS) | yes | Fireleaf Mantle (240046, -2.58 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -2.87 DPS) [vendor]; Darkspear Shoulderpads (272103, -3.74 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 39.7 spell_power points (8.16 DPS) | yes | Hide of the Wild (18510, +1.53 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -2.86 DPS) [world_drop]; Darkspear Raider's Cloak (272063, -3.38 DPS) [vendor] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 103.9 spell_power points (21.32 DPS) | yes | Fireleaf Garb (240051, +0.63 DPS, sim-verified) [vendor]; Field Marshal's Silk Vestments (16443, -8.12 DPS) [vendor]; Field Marshal's Silk Vestments (231603, -8.12 DPS) [pvp] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 64.9 spell_power points (13.32 DPS) | yes | Arcanist Bindings (16799, -6.38 DPS) [world_drop]; Dryad's Wrist Bindings (19595, -6.42 DPS) [rep]; Fireleaf Wristwraps (240044, -6.72 DPS, sim-verified) [vendor] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 75.7 spell_power points (15.52 DPS) | yes | Fireleaf Mitts (240049, +1.58 DPS, sim-verified) [vendor]; Raider Handwraps (272097, -5.83 DPS) [vendor]; Sorcerer's Gloves (22066, -6.40 DPS) [quest] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 77.0 spell_power points (15.80 DPS) | yes | Knowledge of the Timbermaw (228190, -1.62 DPS) [vendor]; Magician's Cord (272393, -4.75 DPS) [vendor]; Fireleaf Waistguard (240045, -5.41 DPS, sim-verified) [vendor] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 94.3 spell_power points (19.34 DPS) | yes | Fireleaf Pants (240047, -0.87 DPS, sim-verified) [vendor]; Marshal's Silk Leggings (16442, -5.86 DPS) [vendor]; Marshal's Silk Leggings (231605, -5.86 DPS) [pvp] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 70.0 spell_power points (14.36 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (16437, -3.39 DPS) [vendor]; Marshal's Silk Footwraps (231606, -3.39 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (438.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.71 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.91 DPS) [vendor]; Naglering (11669, -15.69 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (438.9 DPS) | yes | Channeler's Ring (272406, -1.13 DPS) [vendor]; Cauterizing Band (19140, -1.14 DPS) [world_drop]; Naglering (11669, -11.11 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (438.9 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (438.9 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, -0.41 DPS) [dungeon]; Weakness Analyzer (272438, -1.90 DPS, sim-verified) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (438.9 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Teebu's Blazing Longsword (1728, -19.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 380.1 spell_power points (77.99 DPS) | yes | Bonecreeper Stylus (13938, -11.91 DPS) [dungeon]; Sparkling Crystal Wand (20672, -11.95 DPS) [world]; Ritssyn's Wand of Bad Mojo (22408, -26.18 DPS, sim-verified) [dungeon] |

**New at 60:** head: Fireleaf Circlet; neck: Amulet of the Dawn; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1020, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 30.1. Weights run: 0.8s. Verify run: 0.6s. 141 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.965 ± 0.012, crit=0.069 ± 0.006 per rating point (14 rating = 1%, 0.972 per %), hit=0.183 ± 0.003 per rating point (10 rating = 1%, 1.826 per %), spell_haste=-0.708 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.45 DPS) | yes | Shadow Goggles (4373, -0.54 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.7 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.44 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.9 spell_power points (0.37 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Sanguine Cape (14376, -0.08 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.14 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.8 spell_power points (0.74 DPS) | yes | Mystic's Wrap (14369, -0.23 DPS) [world_drop]; Mystic's Robe (14371, -0.23 DPS) [world_drop]; Gray Woolen Robe (2585, -0.53 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (30.1 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.07 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.37 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.53 DPS) | yes | Pristine Gloves (253913, +0.10 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.02 DPS) [quest]; Gnoll Casting Gloves (892, -0.08 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.9 spell_power points (0.59 DPS) | yes | Novice Arcanist's Sash (253885, -0.07 DPS) [crafted]; Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.48 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.7 spell_power points (1.26 DPS) | yes | Filigreed Pristine Leggings (253937, +0.11 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.73 DPS) [dungeon]; Darkweave Breeches (12987, -0.75 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.38 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.44 DPS) [world]; Red Woolen Boots (4313, -0.52 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 5.8 spell_power points (0.44 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; Sludge-Stained Band (286535, -0.21 DPS) [world]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.38 DPS) | yes | Sludge-Stained Band (286535, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; Loop of Sacrifice (281673, -0.62 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 9.6 spell_power points (0.73 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.15 DPS) [world]; Lesser Staff of the Spire (1300, -0.29 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 299.5 spell_power points (22.54 DPS) | yes | Skycaller (12984, -0.64 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Sizzle Stick (8071, -4.51 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (orc, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 54.5. Weights run: 1.0s. Verify run: 0.6s. 236 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.146 ± 0.023, crit=0.181 ± 0.016 per rating point (14 rating = 1%, 2.538 per %), hit=0.284 ± 0.005 per rating point (10 rating = 1%, 2.838 per %), spell_haste=not significant (0.015 ± 0.231), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 17.5 spell_power points (1.61 DPS) | yes | Nightsky Cowl (4039, -0.27 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.45 DPS) [crafted]; Resilient Cap (14401, -0.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.9 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.96 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.3 spell_power points (1.78 DPS) | yes | Death Speaker Mantle (6685, +0.12 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.37 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.2 spell_power points (0.84 DPS) | yes | Darkspear Raider's Cloak (272078, +0.26 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.21 DPS) [world_drop]; Soft Willow Cape (16661, -0.32 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.9 spell_power points (2.20 DPS) | yes | Death Speaker Robes (6682, -0.34 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.62 DPS) [dungeon]; Pristine Gown (253961, -0.82 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Wristbands (6407, -0.20 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.20 DPS) [quest]; Glowing Magical Bracelets (13106, -0.91 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 11.7 spell_power points (1.08 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.23 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.30 DPS) [crafted]; Blight Gloves (279877, -0.34 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 14.4 spell_power points (1.33 DPS) | yes | Crimson Silk Belt (7055, +0.00 DPS, sim-verified) [crafted]; Invoker's Cord (215366, -0.16 DPS) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.20 DPS) [crafted]; Necromancer Leggings (2277, -0.22 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.75 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 spell_power points (1.38 DPS) | yes | Spidersilk Boots (4320, -0.32 DPS) [crafted]; Frothing Slippers (254003, -0.65 DPS) [crafted]; Acidic Walkers (9454, -0.77 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) (or Black Widow Band (6199)) | Willix the Importer [quest] | 8.0 spell_power points (0.74 DPS) | yes | Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.19 DPS) [vendor]; Black Widow Band (6199, -0.70 DPS, sim-verified) [world] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.01 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor]; Black Widow Band (6199, -0.62 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 12.6 spell_power points (1.16 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.11 DPS) [quest]; Twisted Chanter's Staff (890, -0.11 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.32 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 363.3 spell_power points (33.49 DPS) | yes | Starfaller (13063, +0.16 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Greater Mystic Wand (11290, -4.49 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 236, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 90.9. Weights run: 1.1s. Verify run: 0.6s. 319 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.993 ± 0.029, crit=0.199 ± 0.020 per rating point (14 rating = 1%, 2.791 per %), hit=0.387 ± 0.007 per rating point (10 rating = 1%, 3.869 per %), spell_haste=6.535 ± 0.548, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.64 DPS) | yes | Corpseshroud (10574, -0.27 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.52 DPS) [dungeon]; Augural Shroud (2620, -0.59 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.0 spell_power points (1.63 DPS) | yes | Necklace of Calisea (1714, -0.76 DPS) [world_drop]; Triune Amulet (7722, -0.76 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.18 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.9 spell_power points (2.51 DPS) | yes | Bloodmage Mantle (7684, -0.25 DPS) [dungeon]; Berylline Pads (4197, -0.37 DPS) [quest]; Green Silken Shoulders (7057, -0.40 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (90.9 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.52 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.0 spell_power points (3.52 DPS) | yes | Dreamweave Vest (10021, -0.18 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.26 DPS) [crafted]; Crimson Silk Vest (7058, -0.76 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 11.9 spell_power points (1.50 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.38 DPS) [world_drop]; Mistscape Bracers (4045, -0.50 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 spell_power points (2.76 DPS) | yes | Red Mageweave Gloves (10018, -0.70 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.76 DPS) [crafted]; Black Mageweave Gloves (10003, -0.88 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.9 spell_power points (2.75 DPS) | yes | Gilded Cord (254037, -0.75 DPS) [crafted]; Defiler's Cloth Girdle (20164, -1.00 DPS) [rep]; Defiler's Cloth Girdle (20166, -1.22 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.9 spell_power points (3.26 DPS) | yes | Crimson Silk Pantaloons (7062, -0.26 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.02 DPS) | yes | Gilded Slippers (254001, -0.71 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.39 DPS) [dungeon]; Spidersilk Boots (4320, -1.64 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.0 spell_power points (2.01 DPS) | yes | Reedknot Ring (9622, -1.13 DPS) [quest]; Ogremind Ring (1993, -1.13 DPS) [world_drop]; Voodoo Band (1996, -1.13 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.13 DPS) | yes | Advisor's Ring (19521, -0.25 DPS) [rep]; Voodoo Band (1996, -0.26 DPS) [world_drop]; Reedknot Ring (9622, -1.64 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -0.64 DPS) [dungeon]; Staff of Jordan (873, -1.14 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 319.0 spell_power points (40.13 DPS) | yes | Nether Force Wand (11263, -1.35 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.56 DPS) [quest]; Ragefire Wand (7513, -2.61 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 124.1. Weights run: 0.6s. Verify run: 0.7s. 416 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=1.186 ± 0.069, crit=0.346 ± 0.031 per rating point (14 rating = 1%, 4.846 per %), hit=0.739 ± 0.011 per rating point (10 rating = 1%, 7.393 per %), spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 42.7 spell_power points (5.42 DPS) | yes | Blood Guard's Dreadweave Hat (220907, +0.08 DPS, sim-verified) [vendor]; Dreamweave Circlet (10041, -1.25 DPS) [crafted]; Chief Architect's Monocle (11839, -1.36 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.32 DPS) [dungeon]; Mindburst Medallion (11196, -0.44 DPS) [quest]; Arcane Crystal Pendant (20037, -2.91 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 34.3 spell_power points (4.36 DPS) | yes | Red Mageweave Shoulders (10029, -1.21 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.37 DPS) [vendor]; Kentic Amice (11624, -3.13 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.7 spell_power points (2.88 DPS) | yes | Spritecaster Cape (11623, +1.01 DPS, sim-verified) [dungeon]; Mantle of Lady Falther'ess (23178, -0.38 DPS) [dungeon]; Runecloth Cloak (13860, -0.53 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.7 spell_power points (5.42 DPS) | yes | Runecloth Robe (13858, -0.61 DPS, sim-verified) [crafted]; Runecloth Tunic (13857, -1.61 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.63 DPS) [vendor] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Nethergeld Cuffs (254061, -0.05 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.18 DPS) [quest]; Aristocratic Cuffs (12546, -3.72 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Sorcerer's Gauntlets (226930, -1.16 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.19 DPS) [vendor]; Red Mageweave Gloves (10018, -1.30 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ban'thok Sash (11662, -0.01 DPS) [dungeon]; Deathmage Sash (10771, -0.14 DPS) [dungeon]; Dawnspire Cord (12466, -1.26 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.36 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -0.85 DPS) [dungeon]; Spellshock Leggings (9484, -5.21 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Sandals (254107, -0.30 DPS) [crafted]; Southsea Mojo Boots (20641, -0.38 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -1.32 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.8 spell_power points (2.26 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -0.45 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindseye Circle (10634, -0.37 DPS) [dungeon]; Band of the Unicorn (7553, -0.52 DPS) [world_drop]; Cyclopean Band (11824, -3.85 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, -0.38 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Spellshifter Rod (9527, -0.90 DPS) [quest]; Radiant Staff (249453, -1.50 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -23.84 DPS, sim-verified) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Satyrmane Sash; legs: Stone Guard's Dreadweave Leggings; finger1: Brainlash; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 416, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 398.6. Weights run: 1.0s. Verify run: 0.7s. 1011 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.011, intellect=1.454 ± 0.084, crit=0.475 ± 0.039 per rating point (14 rating = 1%, 6.656 per %), hit=1.212 ± 0.029 per rating point (10 rating = 1%, 12.117 per %), spell_haste=13.681 ± 1.123, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | 94.3 spell_power points (19.34 DPS) | yes | Warlord's Silk Cowl (16533, -6.14 DPS) [vendor]; Warlord's Silk Cowl (231601, -6.14 DPS) [pvp]; Fireleaf Hood (240048, -7.74 DPS, sim-verified) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (398.6 DPS) | yes | Beads of Ogre Mojo (22149, -0.71 DPS) [quest]; Pebble of Kajaro (19600, -1.23 DPS) [quest]; Jewel of Kajaro (19601, -4.85 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 69.5 spell_power points (14.25 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -2.87 DPS) [vendor]; Darkspear Shoulderpads (272103, -3.74 DPS) [vendor]; Fireleaf Mantle (240046, -7.45 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 39.7 spell_power points (8.16 DPS) | yes | Crystalline Threaded Cape (20697, -2.86 DPS) [world_drop]; Deep Woodlands Cloak (19121, -3.01 DPS) [quest]; Hide of the Wild (18510, -4.68 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 103.9 spell_power points (21.32 DPS) | yes | Fireleaf Garb (240051, -5.07 DPS, sim-verified) [vendor]; Warlord's Silk Raiment (16535, -8.12 DPS) [vendor]; Warlord's Silk Raiment (231596, -8.12 DPS) [pvp] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 64.9 spell_power points (13.32 DPS) | yes | Fireleaf Wristwraps (240044, +1.18 DPS, sim-verified) [vendor]; Arcanist Bindings (16799, -6.38 DPS) [world_drop]; Dryad's Wrist Bindings (19595, -6.42 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 75.7 spell_power points (15.52 DPS) | yes | Fireleaf Mitts (240049, -4.19 DPS, sim-verified) [vendor]; Raider Handwraps (272097, -5.83 DPS) [vendor]; Sorcerer's Gloves (22066, -6.40 DPS) [quest] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 77.0 spell_power points (15.80 DPS) | yes | Knowledge of the Timbermaw (228190, -1.62 DPS) [vendor]; Magician's Cord (272393, -4.75 DPS) [vendor]; Fireleaf Waistguard (240045, -7.80 DPS, sim-verified) [vendor] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | 94.3 spell_power points (19.34 DPS) | yes | Fireleaf Pants (240047, -1.98 DPS, sim-verified) [vendor]; General's Silk Trousers (16534, -5.86 DPS) [vendor]; General's Silk Trousers (231595, -5.86 DPS) [pvp] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 70.0 spell_power points (14.36 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (16539, -3.39 DPS) [vendor]; General's Silk Boots (231597, -3.39 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (398.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.71 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.91 DPS) [vendor]; Naglering (11669, -25.28 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (398.6 DPS) | yes | Channeler's Ring (272406, -1.13 DPS) [vendor]; Cauterizing Band (19140, -1.14 DPS) [world_drop]; Naglering (11669, -13.43 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (398.6 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (398.6 DPS) | yes | Serenity Field (272439, +0.19 DPS, sim-verified) [vendor]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Trindlehaven Staff (13161) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-verified (398.6 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS, sim-verified) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.67 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 380.1 spell_power points (77.99 DPS) | yes | Bonecreeper Stylus (13938, -11.91 DPS) [dungeon]; Sparkling Crystal Wand (20672, -11.95 DPS) [world]; Ritssyn's Wand of Bad Mojo (22408, -30.11 DPS, sim-verified) [dungeon] |

**New at 60:** head: Fireleaf Circlet; neck: Amulet of the Dawn; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Frozen Heart of the Mountain; main_hand: Trindlehaven Staff; ranged: Torch of Light

No-known-source sample (15 of 1011, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

