# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 21.7. Weights run: 0.8s. Verify run: 0.8s. 127 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.740 ± 0.014, crit=0.376 ± 0.034, hit=1.173 ± 0.026, spell_haste=not significant (1.258 ± 0.345), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Shadow Goggles (4373, -1.17 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.7 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.68 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.37 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.11 DPS) [dungeon]; Black Whelp Cloak (7283, -0.11 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.7 spell_power points (0.77 DPS) | yes | Green Woolen Robe (6243, -0.31 DPS) [crafted]; Mystic's Wrap (14369, -0.31 DPS) [world_drop]; Gray Woolen Robe (2585, -0.50 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.4 spell_power points (0.39 DPS) | yes | Mindthrust Bracers (1974, -0.03 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.13 DPS) [world_drop]; Mystic's Bracelets (14366, -0.26 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.62 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.09 DPS) [world]; Blight Gloves (279877, -0.16 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.0 spell_power points (0.62 DPS) | yes | Keller's Girdle (2911, -0.09 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.24 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.35 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (21.7 DPS) | yes | Abomination Skin Leggings (23173, -0.29 DPS, sim-verified) [dungeon]; Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Darkweave Breeches (12987, -0.47 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 spell_power points (0.88 DPS) | yes | Pristine Boots (253889, -0.10 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.44 DPS) [world]; Red Woolen Boots (4313, -0.53 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (0.58 DPS) | yes | Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Loop of Sacrifice (281673, -0.25 DPS) [quest]; Sludge-Stained Band (286535, -0.31 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.44 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Sludge-Stained Band (286535, -0.18 DPS) [world]; Lavishly Jeweled Ring (1156, -0.46 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.4 spell_power points (0.66 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.13 DPS) [world]; Lesser Staff of the Spire (1300, -0.26 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 254.3 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.28 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 127, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 54.3. Weights run: 0.8s. Verify run: 0.6s. 219 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.984 ± 0.010, crit=0.303 ± 0.045, hit=2.022 ± 0.044, spell_haste=not significant (0.291 ± 0.092), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.8 spell_power points (1.42 DPS) | yes | Holy Shroud (2721, -0.43 DPS) [world_drop]; Shadow Hood (4323, -0.45 DPS) [crafted]; Nightsky Cowl (4039, -0.74 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.9 spell_power points (1.16 DPS) | yes | Darkspear Warding Pendant (272075, -0.59 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.80 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.80 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.9 spell_power points (1.60 DPS) | yes | Fairywing Mantle (9536, -0.27 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop]; Death Speaker Mantle (6685, -0.45 DPS, sim-verified) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.9 spell_power points (0.71 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.08 DPS) [quest]; Resilient Cape (14400, -0.18 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.8 spell_power points (1.95 DPS) | yes | Death Speaker Robes (6682, -0.52 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.63 DPS) [dungeon]; Pristine Gown (253961, -0.71 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Nightsky Wristbands (6407, -0.28 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.28 DPS) [quest]; Glowing Magical Bracelets (13106, -0.99 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 14.8 spell_power points (1.33 DPS) | yes | Truefaith Gloves (7049, -0.43 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.62 DPS) [dungeon]; Serpent Gloves (5970, -0.70 DPS) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 14.0 spell_power points (1.25 DPS) | yes | Crimson Silk Belt (7055, +0.00 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Invoker's Cord (215366, -0.18 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (54.3 DPS) | yes | Gaze Dreamer Pants (6903, -0.17 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Abomination Skin Leggings (23173, -0.95 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.9 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.26 DPS) [crafted]; Frothing Slippers (254003, -0.63 DPS) [crafted]; Acidic Walkers (9454, -1.10 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.63 DPS) | yes | Black Widow Band (6199, -0.01 DPS) [world]; Snake Hoop (6750, -0.01 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.0 spell_power points (0.62 DPS) | yes | Snake Hoop (6750, -0.01 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor]; Black Widow Band (6199, -0.25 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Talisman of Arathor (21119, -2.51 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.8 spell_power points (0.97 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Scorn's Focal Dagger (23168, -0.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 373.4 spell_power points (33.48 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; trinket1: Darkspear Voodoo Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 78.5. Weights run: 0.9s. Verify run: 0.6s. 304 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.216 ± 0.020, crit=0.512 ± 0.078, hit=3.296 ± 0.077, spell_haste=not significant (0.544 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 23.2 spell_power points (2.24 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.21 DPS) [crafted]; Thinking Cap (2624, -0.24 DPS) [world_drop]; Corpseshroud (10574, -2.51 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.3 spell_power points (1.38 DPS) | yes | Necklace of Calisea (1714, -0.56 DPS) [world_drop]; Triune Amulet (7722, -0.56 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.82 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.8 spell_power points (2.21 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Death Speaker Mantle (6685, -0.33 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (78.5 DPS) | yes | Long Silken Cloak (4326, -0.13 DPS) [crafted]; Guardian Cloak (5965, -0.13 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.30 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.3 spell_power points (2.83 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.07 DPS) [crafted]; Green Silk Armor (7065, -0.43 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 10.9 spell_power points (1.06 DPS) | yes | Aurora Bracers (4043, +0.00 DPS, sim-verified) [world_drop]; Mistscape Bracers (4045, -0.12 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.12 DPS) [quest] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 23.2 spell_power points (2.24 DPS) | yes | Dreamweave Gloves (10019, -0.20 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.44 DPS) [crafted]; Town Clerk's Mittens (270029, -0.56 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.2 spell_power points (2.44 DPS) | yes | Gilded Cord (254037, -0.73 DPS) [crafted]; Highlander's Cloth Girdle (20099, -1.02 DPS) [rep]; Highlander's Cloth Girdle (20098, -1.07 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.6 spell_power points (2.76 DPS) | yes | Crimson Silk Pantaloons (7062, -0.42 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.95 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.15 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.32 DPS) | yes | Gilded Slippers (254001, -0.29 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.90 DPS) [dungeon]; Spidersilk Boots (4320, -1.17 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.3 spell_power points (1.67 DPS) | yes | Ogremind Ring (1993, -0.85 DPS) [world_drop]; Voodoo Band (1996, -0.85 DPS) [world_drop]; Mindbender Loop (5009, -0.85 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.87 DPS) | yes | Voodoo Band (1996, -0.05 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Ogremind Ring (1993, -0.74 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -0.17 DPS) [dungeon]; Illusionary Rod (7713, -0.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 414.7 spell_power points (40.09 DPS) | yes | Umbral Wand (5216, -1.60 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.50 DPS) [dungeon]; Twisted Nether Wand (249144, -5.51 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Red Mageweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 304, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 124.5. Weights run: 0.6s. Verify run: 0.6s. 402 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.099 ± 0.047, crit=1.607 ± 0.120, hit=6.384 ± 0.118, spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 49.7 spell_power points (5.18 DPS) | yes | Eye of Theradras (17715, -0.10 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -0.91 DPS) [crafted]; Dreamweave Circlet (10041, -1.84 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (115.3 DPS) | yes | Scorn's Icy Choker (23169, -0.19 DPS) [dungeon]; Mindburst Medallion (11196, -0.29 DPS) [quest]; Arcane Crystal Pendant (20037, -2.74 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (114.7 DPS) | yes | Kentic Amice (11624, -0.47 DPS) [dungeon]; Red Mageweave Shoulders (10029, -0.97 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.17 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.6 spell_power points (2.15 DPS) | yes | Runecloth Cloak (13860, -0.29 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.54 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -3.16 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (115.0 DPS) | yes | Runecloth Robe (13858, -1.18 DPS) [crafted]; Runecloth Tunic (13857, -1.24 DPS) [crafted]; Knight's Dreadweave Vest (220886, -2.43 DPS, sim-verified) [vendor] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (114.8 DPS) | yes | Nethergeld Cuffs (254061, -0.02 DPS) [crafted]; Forgotten Wraps (9433, -0.18 DPS) [world_drop]; Aristocratic Cuffs (12546, -2.29 DPS, sim-verified) [dungeon] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 33.3 spell_power points (3.47 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -1.08 DPS) [vendor]; Dreamweave Gloves (10019, -1.14 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 37.0 spell_power points (3.85 DPS) | yes | Dawnspire Cord (12466, -1.05 DPS) [dungeon]; Satyrmane Sash (17755, -1.25 DPS) [dungeon]; Ban'thok Sash (11662, -2.23 DPS, sim-verified) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 47.7 spell_power points (4.97 DPS) | yes | Red Mageweave Pants (10009, -2.14 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -2.59 DPS) [dungeon]; Spellshock Leggings (9484, -3.31 DPS, sim-verified) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 81.7 spell_power points (8.52 DPS) | yes | Earthen Silk Slippers (254013, -0.23 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -6.34 DPS) [crafted]; Southsea Mojo Boots (20641, -6.42 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 63.8 spell_power points (6.65 DPS) | yes | Cyclopean Band (11824, -1.01 DPS, sim-verified) [dungeon]; Brainlash (6440, -4.93 DPS) [dungeon]; Mindseye Circle (10634, -5.28 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (115.1 DPS) | yes | Brainlash (6440, -0.01 DPS) [dungeon]; Mindseye Circle (10634, -0.35 DPS) [dungeon]; Cyclopean Band (11824, -2.60 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Ankh of Life (1713, -0.02 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Might of Hakkar (10838, -0.39 DPS) [world]; Glowing Brightwood Staff (812, -0.40 DPS) [world_drop]; Hammer of the Northern Wind (810, -1.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.98 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -15.48 DPS, sim-verified) [quest] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Virtuous Hands; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 245.3. Weights run: 1.0s. Verify run: 1.0s. 997 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=1.637 ± 0.035, crit=1.001 ± 0.150, hit=7.692 ± 0.171, spell_haste=not significant (0.018 ± 0.293), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Circlet of Revelation (239585) | Leonid Barthalomew the Revered [vendor] | sim-verified (193.6 DPS) | yes | Field Marshal's Headdress (17602, -1.39 DPS) [vendor]; Field Marshal's Satin Crown (231616, -1.39 DPS) [vendor]; Bloodvine Goggles (19999, -8.48 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 0.0 spell_power points (0.00 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Amulet of the Dawn (22657, -4.58 DPS) [quest]; Beads of Ogre Mojo (22149, -4.98 DPS) [quest] |
| shoulder | Virtuous Epaulets (226955) | Mokvar [vendor] | 115.6 spell_power points (13.01 DPS) | yes | Rugged Mantle of the Timbermaw (227808, +0.00 DPS, sim-verified) [vendor]; Shoulderpads of Revelation (239586, -5.81 DPS) [vendor]; Field Marshal's Satin Mantle (17604, -6.51 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 106.0 spell_power points (11.93 DPS) | yes | Howler's Furs (272414, -3.28 DPS) [vendor]; Stalwart Cloak (272415, -3.28 DPS) [vendor]; Earthweave Cloak (21187, -3.67 DPS, sim-verified) [quest] |
| chest | Robe of Revelation (239591) | Leonid Barthalomew the Revered [vendor] | sim-verified (195.6 DPS) | yes | Field Marshal's Satin Vestments (17605, -0.74 DPS) [vendor]; Field Marshal's Satin Robe (231618, -0.74 DPS) [vendor]; Bloodvine Vest (19682, -10.50 DPS, sim-verified) [crafted] |
| wrist | Bindings of Revelation (239588) | Leonid Barthalomew the Revered [vendor] | sim-verified (190.9 DPS) | yes | Wrists of Revelation (239583, -1.77 DPS) [vendor]; Dryad's Wrist Bindings (19595, -2.53 DPS) [rep]; Rockfury Bracers (21186, -5.82 DPS, sim-verified) [quest] |
| hands | Gloves of Revelation (239584) | Leonid Barthalomew the Revered [vendor] | sim-verified (191.0 DPS) | yes | Gloves of Spell Mastery (14146, -2.72 DPS) [crafted]; Gloves of Delusional Power (20618, -2.75 DPS) [world]; Dreadmist Wraps (16705, -5.88 DPS, sim-verified) [dungeon] |
| waist | Belt of Revelation (239590) | Leonid Barthalomew the Revered [vendor] | sim-verified (190.0 DPS) | yes | Belt of the Archmage (18405, -0.75 DPS) [crafted]; Girdle of Revelation (239582, -1.89 DPS) [vendor]; Knowledge of the Timbermaw (228190, -4.96 DPS, sim-verified) [vendor] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (187.4 DPS) | yes | Leggings of Revelation (239587, -0.45 DPS) [vendor]; Bloodvine Leggings (19683, -2.26 DPS, sim-verified) [crafted]; Sentinel's Silk Leggings (237815, -2.75 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 122.1 spell_power points (13.75 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.52 DPS, sim-verified) [vendor]; Argent Elite Boots (227816, -5.09 DPS) [vendor]; Sandals of Revelation (239589, -6.55 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.41 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.52 DPS) [vendor]; Wrath of Cenarius (21190, -2.78 DPS, sim-verified) [quest] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Don Julio's Band (19325, -1.00 DPS) [rep]; Band of Earthen Might (21182, -1.00 DPS) [quest]; Wrath of Cenarius (21190, -1.89 DPS, sim-verified) [quest] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -0.92 DPS, sim-verified) [vendor] |
| main_hand | Persuader (22384) | Blacksmithing [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -0.22 DPS, sim-verified) [world_drop] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | sim-verified (189.6 DPS) | yes | Grand Marshal's Tome of Power (23452, +0.00 DPS) [vendor]; Grand Marshal's Tome of Power (234589, +0.00 DPS) [vendor]; Trance Stone (20582, -4.52 DPS, sim-verified) [world] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (212.4 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.35 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.83 DPS) [dungeon]; Cold Snap (19130, -27.27 DPS, sim-verified) [world] |

**New at 60:** head: Circlet of Revelation; neck: Beads of Ogre Might; shoulder: Virtuous Epaulets; back: Arcanoweave Cloak; chest: Robe of Revelation; wrist: Bindings of Revelation; hands: Gloves of Revelation; waist: Belt of Revelation; legs: Magister's Leggings; feet: Bloodvine Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Channeler's Ring; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Persuader; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 997, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 26.1. Weights run: 0.8s. Verify run: 0.8s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.740 ± 0.014, crit=0.376 ± 0.034, hit=1.173 ± 0.026, spell_haste=not significant (1.258 ± 0.345), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Shadow Goggles (4373, -1.23 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.7 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.01 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.68 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.37 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.11 DPS) [dungeon]; Black Whelp Cloak (7283, -0.11 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.7 spell_power points (0.77 DPS) | yes | Green Woolen Robe (6243, -0.31 DPS) [crafted]; Mystic's Wrap (14369, -0.31 DPS) [world_drop]; Gray Woolen Robe (2585, -0.80 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.4 spell_power points (0.39 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest]; Bright Bracers (3647, -0.13 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.62 DPS) | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Blight Gloves (279877, -0.16 DPS) [quest]; Pristine Gloves (253913, -0.17 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.0 spell_power points (0.62 DPS) | yes | Keller's Girdle (2911, -0.09 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.24 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.52 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (26.1 DPS) | yes | Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.33 DPS, sim-verified) [dungeon]; Darkweave Breeches (12987, -0.47 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 spell_power points (0.88 DPS) | yes | Pristine Boots (253889, -0.26 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.44 DPS) [world]; Red Woolen Boots (4313, -0.53 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.44 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Sludge-Stained Band (286535, -0.18 DPS) [world]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 4.4 spell_power points (0.39 DPS) | yes | Sludge-Stained Band (286535, -0.13 DPS) [world]; Loop of Sacrifice (281673, -0.16 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.20 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.4 spell_power points (0.66 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.13 DPS) [world]; Lesser Staff of the Spire (1300, -0.26 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 254.3 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.78 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 209613 Insignia of the Alliance

### Band 30 (undead, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 48.7. Weights run: 0.8s. Verify run: 0.6s. 216 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.984 ± 0.010, crit=0.303 ± 0.045, hit=2.022 ± 0.044, spell_haste=not significant (0.291 ± 0.092), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.8 spell_power points (1.42 DPS) | yes | Nightsky Cowl (4039, -0.36 DPS, sim-verified) [world_drop]; Holy Shroud (2721, -0.43 DPS) [world_drop]; Shadow Hood (4323, -0.45 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.9 spell_power points (1.16 DPS) | yes | Darkspear Warding Pendant (272075, -0.76 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.80 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.80 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.9 spell_power points (1.60 DPS) | yes | Death Speaker Mantle (6685, -0.03 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.27 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.9 spell_power points (0.71 DPS) | yes | Resilient Cape (14400, -0.18 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -0.20 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.26 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.8 spell_power points (1.95 DPS) | yes | Death Speaker Robes (6682, -0.43 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.63 DPS) [dungeon]; Pristine Gown (253961, -0.71 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Nightsky Wristbands (6407, -0.28 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.28 DPS) [quest]; Glowing Magical Bracelets (13106, -0.75 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.9 spell_power points (0.98 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.27 DPS) [dungeon]; Serpent Gloves (5970, -0.35 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 14.0 spell_power points (1.25 DPS) | yes | Crimson Silk Belt (7055, -0.06 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Invoker's Cord (215366, -0.18 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (48.7 DPS) | yes | Gaze Dreamer Pants (6903, -0.17 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Abomination Skin Leggings (23173, -0.68 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.9 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.26 DPS) [crafted]; Frothing Slippers (254003, -0.63 DPS) [crafted]; Acidic Walkers (9454, -0.91 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.63 DPS) | yes | Snake Hoop (6750, -0.01 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 6.9 spell_power points (0.62 DPS) | yes | Snake Hoop (6750, +0.00 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.08 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Defiler's Talisman (21120, -1.74 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.8 spell_power points (0.97 DPS) | yes | Twisted Chanter's Staff (890, -0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Scorn's Focal Dagger (23168, -0.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 373.4 spell_power points (33.48 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Black Widow Band; trinket1: Darkspear Voodoo Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 216, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 70.8. Weights run: 0.9s. Verify run: 0.6s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.216 ± 0.020, crit=0.512 ± 0.078, hit=3.296 ± 0.077, spell_haste=not significant (0.544 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 23.2 spell_power points (2.24 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.21 DPS) [crafted]; Thinking Cap (2624, -0.24 DPS) [world_drop]; Corpseshroud (10574, -2.19 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.3 spell_power points (1.38 DPS) | yes | Necklace of Calisea (1714, -0.56 DPS) [world_drop]; Triune Amulet (7722, -0.56 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.85 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.8 spell_power points (2.21 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Death Speaker Mantle (6685, -0.33 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (70.1 DPS) | yes | Long Silken Cloak (4326, -0.13 DPS) [crafted]; Guardian Cloak (5965, -0.13 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.63 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.3 spell_power points (2.83 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.07 DPS) [crafted]; Green Silk Armor (7065, -0.43 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 13.7 spell_power points (1.33 DPS) | yes | Windchaser Cuffs (14429, +0.00 DPS, sim-verified) [world_drop]; Aurora Bracers (4043, -0.39 DPS) [world_drop]; Mistscape Bracers (4045, -0.39 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (69.2 DPS) | yes | Stormcloth Gloves (10011, -0.41 DPS) [crafted]; Gilded Handwraps (254021, -0.61 DPS) [crafted]; Red Mageweave Gloves (10018, -0.78 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.2 spell_power points (2.44 DPS) | yes | Defiler's Cloth Girdle (20166, -0.58 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.73 DPS) [crafted]; Defiler's Cloth Girdle (20164, -1.02 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.6 spell_power points (2.76 DPS) | yes | Crimson Silk Pantaloons (7062, -0.08 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.95 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.15 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.32 DPS) | yes | Gilded Slippers (254001, +0.00 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.90 DPS) [dungeon]; Spidersilk Boots (4320, -1.17 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.3 spell_power points (1.67 DPS) | yes | Ogremind Ring (1993, -0.85 DPS) [world_drop]; Voodoo Band (1996, -0.85 DPS) [world_drop]; Mindbender Loop (5009, -0.85 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.87 DPS) | yes | Voodoo Band (1996, -0.05 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Ogremind Ring (1993, -0.23 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -0.17 DPS) [dungeon]; Illusionary Rod (7713, -0.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 414.7 spell_power points (40.09 DPS) | yes | Umbral Wand (5216, -1.03 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.50 DPS) [dungeon]; Twisted Nether Wand (249144, -5.51 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 108.6. Weights run: 0.6s. Verify run: 0.6s. 399 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.099 ± 0.047, crit=1.607 ± 0.120, hit=6.384 ± 0.118, spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 49.7 spell_power points (5.18 DPS) | yes | Eye of Theradras (17715, -0.49 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -0.91 DPS) [crafted]; Dreamweave Circlet (10041, -1.84 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (101.3 DPS) | yes | Scorn's Icy Choker (23169, -0.19 DPS) [dungeon]; Mindburst Medallion (11196, -0.29 DPS) [quest]; Arcane Crystal Pendant (20037, -1.99 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (101.1 DPS) | yes | Kentic Amice (11624, -0.47 DPS) [dungeon]; Red Mageweave Shoulders (10029, -0.97 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.84 DPS, sim-verified) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.9 spell_power points (2.28 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.31 DPS) [dungeon]; Runecloth Cloak (13860, -0.43 DPS) [crafted]; Spritecaster Cape (11623, -0.59 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (101.4 DPS) | yes | Runecloth Robe (13858, -1.18 DPS) [crafted]; Runecloth Tunic (13857, -1.24 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -2.07 DPS, sim-verified) [vendor] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (101.5 DPS) | yes | Nethergeld Cuffs (254061, -0.02 DPS) [crafted]; Forgotten Wraps (9433, -0.18 DPS) [world_drop]; Aristocratic Cuffs (12546, -2.17 DPS, sim-verified) [dungeon] |
| hands | Virtuous Hands (226958) | Mokvar [vendor] | 33.3 spell_power points (3.47 DPS) | yes | Raider Handwraps (272098, -0.06 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.08 DPS) [vendor]; Dreamweave Gloves (10019, -1.14 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 37.0 spell_power points (3.85 DPS) | yes | Dawnspire Cord (12466, -1.05 DPS) [dungeon]; Satyrmane Sash (17755, -1.25 DPS) [dungeon]; Ban'thok Sash (11662, -2.23 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 47.7 spell_power points (4.97 DPS) | yes | Red Mageweave Pants (10009, -2.14 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -2.59 DPS) [dungeon]; Spellshock Leggings (9484, -3.33 DPS, sim-verified) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 81.7 spell_power points (8.52 DPS) | yes | Earthen Silk Slippers (254013, -0.20 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -6.34 DPS) [crafted]; Southsea Mojo Boots (20641, -6.42 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 63.8 spell_power points (6.65 DPS) | yes | Cyclopean Band (11824, -1.11 DPS, sim-verified) [dungeon]; Brainlash (6440, -4.93 DPS) [dungeon]; Mindseye Circle (10634, -5.28 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (100.6 DPS) | yes | Brainlash (6440, -0.01 DPS) [dungeon]; Mindseye Circle (10634, -0.35 DPS) [dungeon]; Cyclopean Band (11824, -1.34 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted]; Uther's Strength (11302, -4.03 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -0.08 DPS, sim-verified) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Might of Hakkar (10838, -0.39 DPS) [world]; Glowing Brightwood Staff (812, -0.40 DPS) [world_drop]; Hammer of the Northern Wind (810, -2.06 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.98 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -17.36 DPS, sim-verified) [quest] |

**New at 50:** head: Blood Guard's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Virtuous Hands; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Rune of the Guard Captain; trinket2: Ankh of Life; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 399, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 229.1. Weights run: 1.0s. Verify run: 1.0s. 994 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=1.637 ± 0.035, crit=1.001 ± 0.150, hit=7.692 ± 0.171, spell_haste=not significant (0.018 ± 0.293), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Circlet of Revelation (239585) | Leonid Barthalomew the Revered [vendor] | sim-verified (175.0 DPS) | yes | Warlord's Satin Cowl (17623, -1.39 DPS) [vendor]; Warlord's Satin Crown (231615, -1.39 DPS) [vendor]; Bloodvine Goggles (19999, -8.29 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 0.0 spell_power points (0.00 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Amulet of the Dawn (22657, -4.58 DPS) [quest]; Beads of Ogre Mojo (22149, -4.98 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | sim-verified (168.8 DPS) | yes | Shoulderpads of Revelation (239586, -0.18 DPS) [vendor]; Warlord's Satin Mantle (17622, -0.88 DPS) [vendor]; Virtuous Epaulets (226955, -2.15 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 106.0 spell_power points (11.93 DPS) | yes | Howler's Furs (272414, -3.28 DPS) [vendor]; Stalwart Cloak (272415, -3.28 DPS) [vendor]; Earthweave Cloak (21187, -4.49 DPS, sim-verified) [quest] |
| chest | Robe of Revelation (239591) | Leonid Barthalomew the Revered [vendor] | sim-verified (177.6 DPS) | yes | Warlord's Satin Robes (17624, -0.74 DPS) [vendor]; Warlord's Satin Robes (231612, -0.74 DPS) [vendor]; Bloodvine Vest (19682, -10.93 DPS, sim-verified) [crafted] |
| wrist | Bindings of Revelation (239588) | Leonid Barthalomew the Revered [vendor] | sim-verified (172.8 DPS) | yes | Wrists of Revelation (239583, -1.77 DPS) [vendor]; Dryad's Wrist Bindings (19595, -2.53 DPS) [rep]; Rockfury Bracers (21186, -6.08 DPS, sim-verified) [quest] |
| hands | Gloves of Revelation (239584) | Leonid Barthalomew the Revered [vendor] | sim-verified (172.1 DPS) | yes | Gloves of Spell Mastery (14146, -2.72 DPS) [crafted]; Gloves of Delusional Power (20618, -2.75 DPS) [world]; Dreadmist Wraps (16705, -5.39 DPS, sim-verified) [dungeon] |
| waist | Belt of Revelation (239590) | Leonid Barthalomew the Revered [vendor] | sim-verified (169.4 DPS) | yes | Belt of the Archmage (18405, -0.75 DPS) [crafted]; Girdle of Revelation (239582, -1.89 DPS) [vendor]; Knowledge of the Timbermaw (228190, -2.69 DPS, sim-verified) [vendor] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (170.5 DPS) | yes | Leggings of Revelation (239587, -0.45 DPS) [vendor]; Sentinel's Silk Leggings (237815, -2.75 DPS) [vendor]; Bloodvine Leggings (19683, -3.84 DPS, sim-verified) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 122.1 spell_power points (13.75 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -3.32 DPS, sim-verified) [vendor]; Argent Elite Boots (227816, -5.09 DPS) [vendor]; Sandals of Revelation (239589, -6.55 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.41 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.52 DPS) [vendor]; Wrath of Cenarius (21190, -1.69 DPS, sim-verified) [quest] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Don Julio's Band (19325, -1.00 DPS) [rep]; Band of Earthen Might (21182, -1.00 DPS) [quest]; Wrath of Cenarius (21190, -1.97 DPS, sim-verified) [quest] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -0.81 DPS, sim-verified) [vendor] |
| main_hand | Persuader (22384) | Blacksmithing [crafted] | 0.0 spell_power points (0.00 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -0.12 DPS, sim-verified) [world_drop]; Swiftstrike Cudgel (11964, -1.58 DPS) [quest] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | sim-verified (170.8 DPS) | yes | High Warlord's Tome of Destruction (23468, +0.00 DPS) [vendor]; High Warlord's Tome of Destruction (234563, +0.00 DPS) [vendor]; Trance Stone (20582, -4.07 DPS, sim-verified) [world] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (196.2 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.35 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.83 DPS) [dungeon]; Cold Snap (19130, -29.56 DPS, sim-verified) [world] |

**New at 60:** head: Circlet of Revelation; neck: Beads of Ogre Might; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Revelation; wrist: Bindings of Revelation; hands: Gloves of Revelation; waist: Belt of Revelation; legs: Magister's Leggings; feet: Bloodvine Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Channeler's Ring; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Persuader; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 994, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

