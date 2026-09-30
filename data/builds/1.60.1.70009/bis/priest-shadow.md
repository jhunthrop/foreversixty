# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 21.6. Weights run: 1.0s. Verify run: 1.1s. 127 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.740 ± 0.014, crit=0.027 ± 0.002 per rating point (14 rating = 1%, 0.376 per %), hit=0.117 ± 0.003 per rating point (10 rating = 1%, 1.173 per %), spell_haste=not significant (1.258 ± 0.345), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Shadow Goggles (4373, -1.15 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.7 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.68 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.37 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.11 DPS) [dungeon]; Caretaker's Cape (20428, -0.11 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.7 spell_power points (0.77 DPS) | yes | Green Woolen Robe (6243, -0.31 DPS) [crafted]; Mystic's Wrap (14369, -0.31 DPS) [world_drop]; Gray Woolen Robe (2585, -0.50 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.4 spell_power points (0.39 DPS) | yes | Mindthrust Bracers (1974, -0.01 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.13 DPS) [world_drop]; Repurposed Hair Band (281256, -0.26 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.62 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.09 DPS) [world]; Blight Gloves (279877, -0.16 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.0 spell_power points (0.62 DPS) | yes | Keller's Girdle (2911, -0.09 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.24 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.36 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (21.6 DPS) | yes | Abomination Skin Leggings (23173, -0.29 DPS, sim-verified) [dungeon]; Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Darkweave Breeches (12987, -0.47 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 spell_power points (0.88 DPS) | yes | Pristine Boots (253889, -0.10 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.44 DPS) [world]; Red Woolen Boots (4313, -0.53 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (0.58 DPS) | yes | Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Loop of Sacrifice (281673, -0.25 DPS) [quest]; Sludge-Stained Band (286535, -0.31 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.44 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Sludge-Stained Band (286535, -0.18 DPS) [world]; Lavishly Jeweled Ring (1156, -0.46 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.4 spell_power points (0.66 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.13 DPS) [world]; Lesser Staff of the Spire (1300, -0.26 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 254.3 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.29 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 127, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 54.3. Weights run: 1.1s. Verify run: 0.8s. 219 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.984 ± 0.010, crit=0.022 ± 0.003 per rating point (14 rating = 1%, 0.303 per %), hit=0.202 ± 0.004 per rating point (10 rating = 1%, 2.022 per %), spell_haste=not significant (0.291 ± 0.092), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.8 spell_power points (1.42 DPS) | yes | Holy Shroud (2721, -0.43 DPS) [world_drop]; Shadow Hood (4323, -0.45 DPS) [crafted]; Nightsky Cowl (4039, -0.74 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.9 spell_power points (1.16 DPS) | yes | Darkspear Warding Pendant (272075, -0.59 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.80 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.80 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.9 spell_power points (1.60 DPS) | yes | Fairywing Mantle (9536, -0.27 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop]; Death Speaker Mantle (6685, -0.45 DPS, sim-verified) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.9 spell_power points (0.71 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.08 DPS) [quest]; Resilient Cape (14400, -0.18 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.8 spell_power points (1.95 DPS) | yes | Death Speaker Robes (6682, -0.52 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.63 DPS) [dungeon]; Pristine Gown (253961, -0.71 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Nightsky Wristbands (6407, -0.28 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.28 DPS) [quest]; Glowing Magical Bracelets (13106, -0.99 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 14.8 spell_power points (1.33 DPS) | yes | Truefaith Gloves (7049, -0.43 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.62 DPS) [dungeon]; Shilly Mitts (9609, -0.70 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 14.0 spell_power points (1.25 DPS) | yes | Crimson Silk Belt (7055, +0.00 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Invoker's Cord (215366, -0.18 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (54.3 DPS) | yes | Gaze Dreamer Pants (6903, -0.17 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Abomination Skin Leggings (23173, -0.95 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.9 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.26 DPS) [crafted]; Frothing Slippers (254003, -0.63 DPS) [crafted]; Acidic Walkers (9454, -1.10 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.63 DPS) | yes | Black Widow Band (6199, -0.01 DPS) [world]; Snake Hoop (6750, -0.01 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (0.62 DPS) | yes | Black Widow Band (6199, -0.01 DPS) [world]; Snake Hoop (6750, -0.05 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Talisman of Arathor (21119, -2.51 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.8 spell_power points (0.97 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Scorn's Focal Dagger (23168, -0.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 373.4 spell_power points (33.48 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Greater Mystic Wand (11290, -4.48 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; trinket1: Darkspear Voodoo Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 78.5. Weights run: 1.2s. Verify run: 0.8s. 304 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.216 ± 0.020, crit=0.037 ± 0.006 per rating point (14 rating = 1%, 0.512 per %), hit=0.330 ± 0.008 per rating point (10 rating = 1%, 3.296 per %), spell_haste=not significant (0.544 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 23.2 spell_power points (2.24 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.21 DPS) [crafted]; Miner's Hat of the Deep (9429, -0.24 DPS) [dungeon]; Corpseshroud (10574, -2.51 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.3 spell_power points (1.38 DPS) | yes | Necklace of Calisea (1714, -0.56 DPS) [world_drop]; Triune Amulet (7722, -0.56 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.82 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.8 spell_power points (2.21 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Death Speaker Mantle (6685, -0.33 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (78.5 DPS) | yes | Long Silken Cloak (4326, -0.13 DPS) [crafted]; Guardian Cloak (5965, -0.13 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.30 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.3 spell_power points (2.83 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.07 DPS) [crafted]; Green Silk Armor (7065, -0.43 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 10.9 spell_power points (1.06 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [world_drop]; Aurora Bracers (4043, -0.12 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.12 DPS) [quest] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 23.2 spell_power points (2.24 DPS) | yes | Dreamweave Gloves (10019, -0.20 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.44 DPS) [crafted]; Town Clerk's Mittens (270029, -0.56 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.2 spell_power points (2.44 DPS) | yes | Gilded Cord (254037, -0.73 DPS) [crafted]; Highlander's Cloth Girdle (20099, -1.02 DPS) [rep]; Highlander's Cloth Girdle (20098, -1.07 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.6 spell_power points (2.76 DPS) | yes | Crimson Silk Pantaloons (7062, -0.42 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.95 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.15 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.32 DPS) | yes | Gilded Slippers (254001, -0.29 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.90 DPS) [dungeon]; Spidersilk Boots (4320, -1.17 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.3 spell_power points (1.67 DPS) | yes | Ogremind Ring (1993, -0.85 DPS) [world_drop]; Voodoo Band (1996, -0.85 DPS) [world_drop]; Mindbender Loop (5009, -0.85 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.87 DPS) | yes | Ogremind Ring (1993, -0.05 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Voodoo Band (1996, -0.74 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -0.17 DPS) [dungeon]; Staff of Jordan (873, -0.64 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 414.7 spell_power points (40.09 DPS) | yes | Umbral Wand (5216, -1.60 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.50 DPS) [dungeon]; Twisted Nether Wand (249144, -5.51 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Red Mageweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 304, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 124.3. Weights run: 0.7s. Verify run: 0.8s. 402 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.099 ± 0.047, crit=0.115 ± 0.009 per rating point (14 rating = 1%, 1.607 per %), hit=0.638 ± 0.012 per rating point (10 rating = 1%, 6.384 per %), spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.0 spell_power points (4.27 DPS) | yes | Dreamweave Circlet (10041, -0.17 DPS, sim-verified) [crafted]; Chief Architect's Monocle (11839, -1.18 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Hat (220889, -1.27 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.19 DPS) [dungeon]; Mindburst Medallion (11196, -0.29 DPS) [quest]; Arcane Crystal Pendant (20037, -2.49 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.8 spell_power points (3.42 DPS) | yes | Red Mageweave Shoulders (10029, -0.97 DPS) [crafted]; Inquisitor's Shawl (19507, -1.20 DPS) [dungeon]; Kentic Amice (11624, -4.31 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.6 spell_power points (2.15 DPS) | yes | Runecloth Cloak (13860, -0.29 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.54 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -3.04 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.0 spell_power points (4.27 DPS) | yes | Runecloth Tunic (13857, -1.24 DPS) [crafted]; Robe of the Magi (1716, -1.29 DPS) [world_drop]; Runecloth Robe (13858, -1.41 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Nethergeld Cuffs (254061, -0.02 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.18 DPS) [quest]; Aristocratic Cuffs (12546, -2.42 DPS, sim-verified) [dungeon] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 33.3 spell_power points (3.47 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.08 DPS) [vendor]; Dreamweave Gloves (10019, -1.14 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.9 spell_power points (2.80 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.22 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight's Dreadweave Leggings (220888, -0.04 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.46 DPS) [dungeon]; Spellshock Leggings (9484, -3.69 DPS, sim-verified) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 24.3 spell_power points (2.53 DPS) | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -0.35 DPS) [crafted]; Southsea Mojo Boots (20641, -0.44 DPS) [quest] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.7 spell_power points (1.74 DPS) | yes | Brainlash (6440, -0.02 DPS) [dungeon]; Mindseye Circle (10634, -0.37 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.6 spell_power points (1.73 DPS) | yes | Mindseye Circle (10634, -0.35 DPS) [dungeon]; Band of the Unicorn (7553, -0.37 DPS) [world_drop]; Brainlash (6440, -1.91 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Smoking Heart of the Mountain (11811, -0.79 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.69 DPS) [quest]; Radiant Staff (249453, -1.15 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.98 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -14.64 DPS, sim-verified) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Virtuous Hands; waist: Dawnspire Cord; feet: Sergeant Major's Dreadweave Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 262.1. Weights run: 1.3s. Verify run: 0.8s. 1036 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.637 ± 0.035, crit=0.071 ± 0.011 per rating point (14 rating = 1%, 1.001 per %), hit=0.769 ± 0.017 per rating point (10 rating = 1%, 7.692 per %), spell_haste=not significant (0.018 ± 0.293), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Field Marshal's Satin Crown (231616) | Captain Dirgehammer [vendor] | 78.8 spell_power points (8.87 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Circlet of Revelation (239585, -1.54 DPS) [vendor]; Field Marshal's Satin Hood (231622, -3.70 DPS, sim-verified) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 36.3 spell_power points (4.08 DPS) | yes | Lady Maye's Pendant (14558, -0.58 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.77 DPS) [world_drop]; Beads of Ogre Mojo (22149, -1.60 DPS, sim-verified) [quest] |
| shoulder | Field Marshal's Satin Epaulets (231621) | Captain Dirgehammer [vendor] | 57.7 spell_power points (6.50 DPS) | yes | Field Marshal's Satin Mantle (17604, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -0.58 DPS) [vendor]; Darkspear Shoulderpads (272103, -1.91 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.8 spell_power points (4.14 DPS) | yes | Hide of the Wild (18510, -0.72 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -1.15 DPS) [world_drop]; Darkspear Raider's Cloak (272063, -1.19 DPS) [vendor] |
| chest | Robe of Revelation (239591) | Leonid Barthalomew the Revered [vendor] | 85.4 spell_power points (9.61 DPS) | yes | Field Marshal's Satin Vestments (17605, -0.74 DPS) [vendor]; Field Marshal's Satin Tunic (231624, -2.09 DPS) [vendor]; Field Marshal's Satin Robe (231618, -3.05 DPS, sim-verified) [vendor] |
| wrist | Bindings of Revelation (239588) | Leonid Barthalomew the Revered [vendor] | 44.6 spell_power points (5.02 DPS) | yes | Marshal's Satin Bracers (17606, -1.51 DPS) [pvp]; Dryad's Wrist Bindings (19596, -1.66 DPS) [rep]; Dryad's Wrist Bindings (19595, -4.32 DPS, sim-verified) [rep] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 52.2 spell_power points (5.88 DPS) | yes | Gloves of Revelation (239584, -0.65 DPS, sim-verified) [vendor]; Marshal's Satin Gloves (17608, -0.89 DPS) [vendor]; Marshal's Satin Grips (231617, -0.89 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 68.7 spell_power points (7.73 DPS) | yes | Belt of Revelation (239590, -1.90 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -2.42 DPS) [crafted]; Girdle of Prophecy (16817, -2.67 DPS) [world_drop] |
| legs | Leggings of Revelation (239587) | Leonid Barthalomew the Revered [vendor] | 70.7 spell_power points (7.96 DPS) | yes | Marshal's Satin Pants (17603, -0.49 DPS) [vendor]; Sentinel's Silk Leggings (22752, -1.31 DPS) [rep]; Marshal's Satin Leggings (231619, -3.42 DPS, sim-verified) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 52.9 spell_power points (5.95 DPS) | yes | Marshal's Satin Sandals (17607, -0.05 DPS) [vendor]; Sandals of Revelation (239589, -0.22 DPS) [vendor]; Marshal's Satin Treads (231620, -2.13 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (262.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.41 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.52 DPS) [vendor]; Naglering (11669, -12.13 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (262.1 DPS) | yes | Cauterizing Band (19140, -0.60 DPS) [world_drop]; Channeler's Ring (272406, -1.06 DPS) [vendor]; Naglering (11669, -8.71 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (262.1 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (262.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, -0.23 DPS) [dungeon]; Weakness Analyzer (272438, -1.07 DPS, sim-verified) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (262.1 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -13.07 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 687.9 spell_power points (77.44 DPS) | yes | Bonecreeper Stylus (13938, -12.83 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.88 DPS) [world]; Ritssyn's Wand of Bad Mojo (22408, -23.04 DPS, sim-verified) [dungeon] |

**New at 60:** head: Field Marshal's Satin Crown; neck: Amulet of the Dawn; shoulder: Field Marshal's Satin Epaulets; back: Arcanoweave Cloak; chest: Robe of Revelation; wrist: Bindings of Revelation; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Leggings of Revelation; feet: Bloodvine Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 1036, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 26.4. Weights run: 1.0s. Verify run: 1.0s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.740 ± 0.014, crit=0.027 ± 0.002 per rating point (14 rating = 1%, 0.376 per %), hit=0.117 ± 0.003 per rating point (10 rating = 1%, 1.173 per %), spell_haste=not significant (1.258 ± 0.345), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Shadow Goggles (4373, -1.04 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.7 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.68 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Feyscale Cloak (6632, -0.09 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Pearl-clasped Cloak (5542, -0.33 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.7 spell_power points (0.77 DPS) | yes | Green Woolen Robe (6243, -0.31 DPS) [crafted]; Mystic's Wrap (14369, -0.31 DPS) [world_drop]; Gray Woolen Robe (2585, -0.44 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.4 spell_power points (0.39 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest]; Bright Bracers (3647, -0.13 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.62 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.09 DPS) [world]; Blight Gloves (279877, -0.16 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.0 spell_power points (0.62 DPS) | yes | Keller's Girdle (2911, -0.09 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.20 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.24 DPS) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.35 DPS, sim-verified) [dungeon]; Darkweave Breeches (12987, -0.47 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 spell_power points (0.88 DPS) | yes | Pristine Boots (253889, +0.00 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.44 DPS) [world]; Red Woolen Boots (4313, -0.53 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.44 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Sludge-Stained Band (286535, -0.18 DPS) [world]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 4.4 spell_power points (0.39 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -0.13 DPS) [world]; Volcanic Rock Ring (12053, -0.20 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.4 spell_power points (0.66 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.13 DPS) [world]; Lesser Staff of the Spire (1300, -0.26 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 254.3 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.59 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 209613 Insignia of the Alliance

### Band 30 (undead, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 48.7. Weights run: 1.1s. Verify run: 0.8s. 216 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.984 ± 0.010, crit=0.022 ± 0.003 per rating point (14 rating = 1%, 0.303 per %), hit=0.202 ± 0.004 per rating point (10 rating = 1%, 2.022 per %), spell_haste=not significant (0.291 ± 0.092), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.8 spell_power points (1.42 DPS) | yes | Holy Shroud (2721, -0.43 DPS) [world_drop]; Nightsky Cowl (4039, -0.45 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.45 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.9 spell_power points (1.16 DPS) | yes | Darkspear Warding Pendant (272075, -0.71 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.80 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.80 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.9 spell_power points (1.60 DPS) | yes | Death Speaker Mantle (6685, -0.16 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.27 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.9 spell_power points (0.71 DPS) | yes | Darkspear Raider's Cloak (272078, -0.02 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.18 DPS) [world_drop]; Hillman's Cloak (3719, -0.26 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.8 spell_power points (1.95 DPS) | yes | Death Speaker Robes (6682, -0.35 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.63 DPS) [dungeon]; Pristine Gown (253961, -0.71 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Nightsky Wristbands (6407, -0.28 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.28 DPS) [quest]; Glowing Magical Bracelets (13106, -0.62 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.9 spell_power points (0.98 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.27 DPS) [dungeon]; Serpent Gloves (5970, -0.35 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 14.0 spell_power points (1.25 DPS) | yes | Crimson Silk Belt (7055, -0.17 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Invoker's Cord (215366, -0.18 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (48.7 DPS) | yes | Gaze Dreamer Pants (6903, -0.17 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Abomination Skin Leggings (23173, -0.66 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.9 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.26 DPS) [crafted]; Frothing Slippers (254003, -0.63 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.63 DPS) | yes | Black Widow Band (6199, -0.01 DPS) [world]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) (or Black Widow Band (6199)) | Willix the Importer [quest] | 6.9 spell_power points (0.62 DPS) | yes | Black Widow Band (6199, -0.07 DPS, sim-verified) [world]; Sea Giant's Toe Ring (274746, -0.08 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Defiler's Talisman (21120, -1.60 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.8 spell_power points (0.97 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Scorn's Focal Dagger (23168, -0.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 373.4 spell_power points (33.48 DPS) | yes | Starfaller (13063, -0.03 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Greater Mystic Wand (11290, -4.48 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 216, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 70.8. Weights run: 1.2s. Verify run: 0.8s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.216 ± 0.020, crit=0.037 ± 0.006 per rating point (14 rating = 1%, 0.512 per %), hit=0.330 ± 0.008 per rating point (10 rating = 1%, 3.296 per %), spell_haste=not significant (0.544 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 23.2 spell_power points (2.24 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.21 DPS) [crafted]; Miner's Hat of the Deep (9429, -0.24 DPS) [dungeon]; Corpseshroud (10574, -2.19 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.3 spell_power points (1.38 DPS) | yes | Necklace of Calisea (1714, -0.56 DPS) [world_drop]; Triune Amulet (7722, -0.56 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.85 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.8 spell_power points (2.21 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Death Speaker Mantle (6685, -0.33 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Long Silken Cloak (4326, -0.13 DPS) [crafted]; Guardian Cloak (5965, -0.13 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.63 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.3 spell_power points (2.83 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.07 DPS) [crafted]; Green Silk Armor (7065, -0.43 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 13.7 spell_power points (1.33 DPS) | yes | Windchaser Cuffs (14429, +0.00 DPS, sim-verified) [world_drop]; Mistscape Bracers (4045, -0.39 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.39 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Gloves (10011, -0.41 DPS) [crafted]; Gilded Handwraps (254021, -0.61 DPS) [crafted]; Red Mageweave Gloves (10018, -0.78 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.2 spell_power points (2.44 DPS) | yes | Defiler's Cloth Girdle (20166, -0.58 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.73 DPS) [crafted]; Defiler's Cloth Girdle (20164, -1.02 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.6 spell_power points (2.76 DPS) | yes | Crimson Silk Pantaloons (7062, -0.08 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.95 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.15 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.32 DPS) | yes | Gilded Slippers (254001, +0.00 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.90 DPS) [dungeon]; Spidersilk Boots (4320, -1.17 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.3 spell_power points (1.67 DPS) | yes | Ogremind Ring (1993, -0.85 DPS) [world_drop]; Voodoo Band (1996, -0.85 DPS) [world_drop]; Mindbender Loop (5009, -0.85 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.87 DPS) | yes | Ogremind Ring (1993, -0.05 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Voodoo Band (1996, -0.23 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -0.17 DPS) [dungeon]; Staff of Jordan (873, -0.64 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 414.7 spell_power points (40.09 DPS) | yes | Umbral Wand (5216, -1.03 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.50 DPS) [dungeon]; Twisted Nether Wand (249144, -5.51 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 109.0. Weights run: 0.7s. Verify run: 0.8s. 399 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.099 ± 0.047, crit=0.115 ± 0.009 per rating point (14 rating = 1%, 1.607 per %), hit=0.638 ± 0.012 per rating point (10 rating = 1%, 6.384 per %), spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 41.0 spell_power points (4.27 DPS) | yes | Dreamweave Circlet (10041, -0.00 DPS, sim-verified) [crafted]; Chief Architect's Monocle (11839, -1.18 DPS) [dungeon]; Blood Guard's Dreadweave Hat (220907, -1.27 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.19 DPS) [dungeon]; Mindburst Medallion (11196, -0.29 DPS) [quest]; Arcane Crystal Pendant (20037, -2.86 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.8 spell_power points (3.42 DPS) | yes | Red Mageweave Shoulders (10029, -0.97 DPS) [crafted]; Inquisitor's Shawl (19507, -1.20 DPS) [dungeon]; Kentic Amice (11624, -3.51 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.9 spell_power points (2.28 DPS) | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Mantle of Lady Falther'ess (23178, -0.31 DPS) [dungeon]; Runecloth Cloak (13860, -0.43 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 41.0 spell_power points (4.27 DPS) | yes | Runecloth Robe (13858, -1.12 DPS, sim-verified) [crafted]; Runecloth Tunic (13857, -1.24 DPS) [crafted]; Robe of the Magi (1716, -1.29 DPS) [world_drop] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Nethergeld Cuffs (254061, -0.02 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.18 DPS) [quest]; Aristocratic Cuffs (12546, -2.49 DPS, sim-verified) [dungeon] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 33.3 spell_power points (3.47 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.08 DPS) [vendor]; Dreamweave Gloves (10019, -1.14 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.9 spell_power points (2.80 DPS) | yes | Satyrmane Sash (17755, -0.14 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.22 DPS) [dungeon]; Deathmage Sash (10771, -0.35 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Stone Guard's Dreadweave Leggings (220906, -0.04 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.46 DPS) [dungeon]; Spellshock Leggings (9484, -3.58 DPS, sim-verified) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 24.3 spell_power points (2.53 DPS) | yes | Earthen Silk Slippers (254013, -0.10 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -0.35 DPS) [crafted]; Southsea Mojo Boots (20641, -0.44 DPS) [quest] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.7 spell_power points (1.74 DPS) | yes | Brainlash (6440, -0.02 DPS) [dungeon]; Mindseye Circle (10634, -0.37 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.6 spell_power points (1.73 DPS) | yes | Mindseye Circle (10634, -0.35 DPS) [dungeon]; Band of the Unicorn (7553, -0.37 DPS) [world_drop]; Brainlash (6440, -1.37 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, -0.34 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.69 DPS) [quest]; Radiant Staff (249453, -1.15 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.98 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -15.94 DPS, sim-verified) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Virtuous Hands; waist: Dawnspire Cord; feet: First Sergeant's Dreadweave Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 399, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 230.7. Weights run: 1.3s. Verify run: 0.8s. 1033 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.637 ± 0.035, crit=0.071 ± 0.011 per rating point (14 rating = 1%, 1.001 per %), hit=0.769 ± 0.017 per rating point (10 rating = 1%, 7.692 per %), spell_haste=not significant (0.018 ± 0.293), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warlord's Satin Crown (231615) | Lady Palanseer [vendor] | 78.8 spell_power points (8.87 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Circlet of Revelation (239585, -1.54 DPS) [vendor]; Warlord's Satin Hood (231635, -2.22 DPS, sim-verified) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 36.3 spell_power points (4.08 DPS) | yes | Lady Maye's Pendant (14558, -0.58 DPS) [world_drop]; Beads of Ogre Mojo (22149, -0.68 DPS, sim-verified) [quest]; Jeweled Amulet of Cainwyn (1443, -0.77 DPS) [world_drop] |
| shoulder | Warlord's Satin Epaulets (231611) | Lady Palanseer [vendor] | 57.7 spell_power points (6.50 DPS) | yes | Warlord's Satin Mantle (17622, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -0.58 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.67 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.8 spell_power points (4.14 DPS) | yes | Hide of the Wild (18510, -0.87 DPS, sim-verified) [crafted]; Deep Woodlands Cloak (19121, -1.13 DPS) [quest]; Crystalline Threaded Cape (20697, -1.15 DPS) [world_drop] |
| chest | Robe of Revelation (239591) | Leonid Barthalomew the Revered [vendor] | 85.4 spell_power points (9.61 DPS) | yes | Warlord's Satin Robes (17624, -0.74 DPS) [vendor]; Warlord's Satin Robes (231612, -0.74 DPS) [pvp]; Warlord's Satin Tunic (231632, -4.98 DPS, sim-verified) [vendor] |
| wrist | Bindings of Revelation (239588) | Leonid Barthalomew the Revered [vendor] | 44.6 spell_power points (5.02 DPS) | yes | Dryad's Wrist Bindings (19595, -1.39 DPS, sim-verified) [rep]; General's Satin Bracers (17619, -1.51 DPS) [pvp]; Dryad's Wrist Bindings (19596, -1.66 DPS) [rep] |
| hands | Gloves of Revelation (239584) | Leonid Barthalomew the Revered [vendor] | sim-verified (230.7 DPS) | yes | General's Satin Gloves (17620, -0.82 DPS) [vendor]; General's Satin Grips (231613, -0.82 DPS) [vendor]; Raider Handwraps (272097, -2.34 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 68.7 spell_power points (7.73 DPS) | yes | Belt of Revelation (239590, -0.25 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -2.42 DPS) [crafted]; Girdle of Prophecy (16817, -2.67 DPS) [world_drop] |
| legs | Leggings of Revelation (239587) | Leonid Barthalomew the Revered [vendor] | 70.7 spell_power points (7.96 DPS) | yes | General's Satin Leggings (17625, -0.49 DPS) [vendor]; General's Satin Leggings (231614, -0.49 DPS) [pvp]; Outrider's Silk Leggings (22747, -2.97 DPS, sim-verified) [rep] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 52.9 spell_power points (5.95 DPS) | yes | General's Satin Boots (17618, -0.05 DPS) [vendor]; Sandals of Revelation (239589, -0.22 DPS) [vendor]; General's Satin Treads (231610, -2.09 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.41 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.52 DPS) [vendor]; Naglering (11669, -8.38 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 0.0 spell_power points (0.00 DPS) | yes | Cauterizing Band (19140, -0.60 DPS) [world_drop]; Channeler's Ring (272406, -1.06 DPS) [vendor]; Naglering (11669, -6.18 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, -0.23 DPS) [dungeon]; Weakness Analyzer (272438, -0.93 DPS, sim-verified) [vendor] |
| main_hand | Trindlehaven Staff (13161) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS, sim-verified) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Argent Crusader (13249, -0.65 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 687.9 spell_power points (77.44 DPS) | yes | Bonecreeper Stylus (13938, -12.83 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.88 DPS) [world]; Ritssyn's Wand of Bad Mojo (22408, -22.18 DPS, sim-verified) [dungeon] |

**New at 60:** head: Warlord's Satin Crown; neck: Amulet of the Dawn; shoulder: Warlord's Satin Epaulets; back: Arcanoweave Cloak; chest: Robe of Revelation; wrist: Bindings of Revelation; hands: Gloves of Revelation; waist: Knowledge of the Timbermaw; legs: Leggings of Revelation; feet: Bloodvine Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Trindlehaven Staff; ranged: Torch of Light

No-known-source sample (15 of 1033, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

