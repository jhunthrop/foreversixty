# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 37.1. Weights run: 0.7s. Verify run: 0.5s. 128 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.502 ± 0.045, crit=0.654 ± 0.041, hit=1.133 ± 0.017, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -1.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.5 spell_power points (1.26 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.42 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Pearl-clasped Cloak (5542, -0.13 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.5 spell_power points (0.99 DPS) | yes | Green Woolen Robe (6243, -0.40 DPS) [crafted]; Green Woolen Vest (2582, -0.46 DPS) [crafted]; Gray Woolen Robe (2585, -1.08 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.40 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.13 DPS) [world_drop]; Mystic's Bracelets (14366, -0.27 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.20 DPS) [crafted]; Blight Gloves (279877, -0.46 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.33 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.69 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.0 spell_power points (1.72 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.79 DPS) [dungeon]; Colorful Kilt (10048, -1.06 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Pristine Boots (253889, -0.59 DPS) [crafted]; Red Woolen Boots (4313, -0.66 DPS) [crafted]; Feather Padded Treads (285345, -0.84 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.0 spell_power points (0.79 DPS) | yes | Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; Sludge-Stained Band (286535, -0.40 DPS) [world]; Loop of Sacrifice (281673, -0.46 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.66 DPS) | yes | Sludge-Stained Band (286535, -0.26 DPS) [world]; Loop of Sacrifice (281673, -0.33 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.55 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 5.0 spell_power points (0.66 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.13 DPS) [world]; Lesser Staff of the Spire (1300, -0.27 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 171.9 spell_power points (22.71 DPS) | yes | Skycaller (12984, -0.89 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 128, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 64.1. Weights run: 0.7s. Verify run: 0.5s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.649 ± 0.085, crit=1.663 ± 0.106, hit=1.646 ± 0.025, spell_haste=0.154 ± 0.033, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.5 spell_power points (1.93 DPS) | yes | Holy Shroud (2721, -0.06 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.54 DPS) [crafted]; Embalmed Shroud (7691, -0.69 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.9 spell_power points (1.68 DPS) | yes | Crystal Starfire Medallion (5003, -1.28 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.28 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.72 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (2.29 DPS) | yes | Fairywing Mantle (9536, -0.46 DPS) [quest]; Death Speaker Mantle (6685, -0.47 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.62 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.6 spell_power points (0.86 DPS) | yes | Darkspear Raider's Cloak (272078, -0.06 DPS) [vendor]; Hillman's Cloak (3719, -0.09 DPS) [crafted]; Cloak of Rot (4462, -0.22 DPS, sim-verified) [world] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.4 spell_power points (2.69 DPS) | yes | Tree Bark Jacket (1486, -0.69 DPS) [dungeon]; Death Speaker Robes (6682, -0.73 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.91 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Nightsky Wristbands (6407, -0.79 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.79 DPS) [quest]; Glowing Magical Bracelets (13106, -1.45 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 11.1 spell_power points (1.72 DPS) | yes | Shilly Mitts (9609, -0.64 DPS) [quest]; Truefaith Gloves (7049, -0.65 DPS) [crafted]; Serpent Gloves (5970, -0.88 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.9 spell_power points (2.00 DPS) | yes | Crimson Silk Belt (7055, -0.37 DPS) [crafted]; Invoker's Cord (215366, -0.42 DPS) [crafted]; Belt of Arugal (6392, -0.42 DPS, sim-verified) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (2.19 DPS) | yes | Gaze Dreamer Pants (6903, -0.22 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.41 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.78 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Nimbus Boots (6998, -0.86 DPS) [quest]; Acidic Walkers (9454, -1.89 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.08 DPS) | yes | Lorekeeper's Ring (20431, -0.31 DPS) [rep]; Black Widow Band (6199, -0.38 DPS) [world]; Minor Channeling Ring (1449, -2.03 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (64.1 DPS) | yes | Black Widow Band (6199, -0.22 DPS) [world]; Snake Hoop (6750, -0.22 DPS) [quest]; Minor Channeling Ring (1449, -1.21 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (62.9 DPS) | yes | Talisman of Arathor (21119, -2.20 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 7.1 spell_power points (1.10 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.10 DPS) [quest]; Twisted Chanter's Staff (890, -0.21 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.30 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 218.2 spell_power points (33.67 DPS) | yes | Starfaller (13063, -0.10 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.90 DPS) [crafted]; Gravestone Scepter (7001, -4.67 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 93.8. Weights run: 0.7s. Verify run: 0.6s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.160 ± 0.169), crit=2.753 ± 0.178, hit=2.986 ± 0.049, spell_haste=0.902 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (92.6 DPS) | yes | Holy Shroud (2721, -0.24 DPS) [world_drop]; Silk Headband (7050, -0.54 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -0.93 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (1.19 DPS) | yes | Necklace of Calisea (1714, -1.02 DPS) [world_drop]; Triune Amulet (7722, -1.02 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.79 DPS, sim-verified) [quest] |
| shoulder | Green Silken Shoulders (7057) | Tailoring [crafted] | sim-verified (93.4 DPS) | yes | Inquisitor's Shawl (19507, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.17 DPS) [quest]; Bloodmage Mantle (7684, -1.73 DPS, sim-verified) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (93.7 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Caretaker's Cape (19532, -0.12 DPS) [rep]; Icy Cloak (4327, -2.06 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (3.44 DPS) | yes | Dreamweave Vest (10021, -0.53 DPS) [crafted]; Elemental Raiment (9434, -0.72 DPS, sim-verified) [world_drop]; Robe of Power (7054, -1.05 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (94.0 DPS) | yes | Condor Bracers (15864, -0.30 DPS) [quest]; Earthen Silk Cuffs (254019, -0.75 DPS) [crafted]; Arcane Runed Bracers (4744, -2.34 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (2.79 DPS) | yes | Black Mageweave Gloves (10003, -0.74 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.90 DPS) [crafted]; Gilded Handwraps (254021, -1.42 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.6 spell_power points (2.19 DPS) | yes | Star Belt (4329, -0.28 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -0.47 DPS) [rep]; Belt of Arugal (6392, -0.77 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.9 spell_power points (2.38 DPS) | yes | Abomination Skin Leggings (23173, -0.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.87 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.55 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | Gilded Slippers (254001, +0.00 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.45 DPS) [crafted]; Acidic Walkers (9454, -2.65 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (1.64 DPS) | yes | Ring of Forlorn Spirits (2043, -0.44 DPS) [quest]; Reedknot Ring (9622, -0.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.74 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.35 DPS) | yes | Ring of Forlorn Spirits (2043, -0.14 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.30 DPS) [quest]; Lorekeeper's Ring (19525, -0.30 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (92.7 DPS) | yes | Rune of Duty (21567, -1.06 DPS, sim-verified) [rep] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (91.7 DPS) | yes | Gut Ripper (2164, -1.48 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -5.55 DPS) [dungeon]; Staff of Jordan (873, -5.65 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 264.1 spell_power points (39.52 DPS) | yes | Nether Force Wand (11263, -0.38 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.26 DPS) [quest]; Ragefire Wand (7513, -2.31 DPS) [quest] |

**New at 40:** head: Augural Shroud; shoulder: Green Silken Shoulders; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 121.2. Weights run: 0.6s. Verify run: 0.6s. 403 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.275 ± 0.650), crit=5.153 ± 0.298, hit=5.744 ± 0.087, spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 101.4 spell_power points (13.31 DPS) | yes | Eye of Theradras (17715, -0.49 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -7.47 DPS) [crafted]; Chief Architect's Monocle (11839, -8.79 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.8 spell_power points (2.34 DPS) | yes | Scorn's Icy Choker (23169, -0.06 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.55 DPS) [quest]; Gemshard Heart (17707, -0.67 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (114.6 DPS) | yes | Red Mageweave Shoulders (10029, -1.29 DPS) [crafted]; Inquisitor's Shawl (19507, -1.62 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.77 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.6 spell_power points (2.84 DPS) | yes | Darkspear Raider's Cloak (272076, -0.50 DPS) [vendor]; Big Voodoo Cloak (8216, -0.68 DPS) [crafted]; Runecloth Cloak (13860, -0.94 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (115.0 DPS) | yes | Runecloth Robe (13858, -1.55 DPS) [crafted]; Robes of Insight (940, -1.66 DPS) [world_drop]; Knight's Dreadweave Vest (220886, -2.13 DPS, sim-verified) [vendor] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (115.6 DPS) | yes | Nethergeld Cuffs (254061, -0.07 DPS) [crafted]; Forgotten Wraps (9433, -0.15 DPS) [world_drop]; Aristocratic Cuffs (12546, -2.75 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 87.3 spell_power points (11.45 DPS) | yes | Raider Handwraps (272098, -0.63 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -8.24 DPS) [vendor]; Red Mageweave Gloves (10018, -8.34 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 87.5 spell_power points (11.48 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.97 DPS) [dungeon]; Deathmage Sash (10771, -8.05 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 99.4 spell_power points (13.05 DPS) | yes | Red Mageweave Pants (10009, -0.57 DPS, sim-verified) [crafted]; Kilt of the Atal'ai Prophet (10807, -9.64 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -9.82 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 76.9 spell_power points (10.09 DPS) | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -7.14 DPS) [crafted]; Southsea Mojo Boots (20641, -7.20 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 spell_power points (7.54 DPS) | yes | Brainlash (6440, -1.84 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -5.53 DPS) [dungeon]; Band of the Unicorn (7553, -5.83 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (115.8 DPS) | yes | Mindseye Circle (10634, -0.31 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Brainlash (6440, -2.97 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (112.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (112.9 DPS) | yes | Illusionary Rod (7713, -1.00 DPS) [dungeon]; Inventor's Focal Sword (17719, -2.01 DPS) [dungeon]; Shortsword of Vengeance (754, -2.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -12.84 DPS, sim-verified) [quest] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Guardian Talisman; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 235.0. Weights run: 0.7s. Verify run: 0.7s. 952 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.725 ± 0.246, crit=5.578 ± 0.410, hit=6.478 ± 0.118, spell_haste=not significant (0.090 ± 0.410), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (206.1 DPS) | yes | Sorcerer's Crown (226935, -3.53 DPS) [quest]; Bloodvine Goggles (19999, -3.73 DPS, sim-verified) [crafted]; Field Marshal's Coronet (16441, -3.75 DPS) [vendor] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (202.3 DPS) | yes | Beads of Ogre Might (22150, +0.00 DPS) [quest]; Amulet of the Dawn (22657, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -5.19 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 145.0 spell_power points (17.86 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -2.74 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -3.38 DPS) [crafted]; Shroud of the Nathrezim (18720, -3.98 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 94.6 spell_power points (11.65 DPS) | yes | Fel Cape (279269, -2.03 DPS) [crafted]; Earthweave Cloak (21187, -3.67 DPS) [quest]; Chromatic Cloak (18509, -5.63 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 252.0 spell_power points (31.04 DPS) | yes | Bloodvine Vest (19682, -6.56 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -9.16 DPS) [vendor]; Earthpower Vest (21183, -13.66 DPS) [quest] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 140.7 spell_power points (17.33 DPS) | yes | Fireleaf Wristwraps (240044, -2.12 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -6.02 DPS) [quest]; Dryad's Wrist Bindings (19595, -12.92 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 222.3 spell_power points (27.39 DPS) | yes | Gloves of Spell Mastery (14146, -7.20 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -8.79 DPS) [vendor]; Sorcerer's Gloves (22066, -14.95 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 146.2 spell_power points (18.01 DPS) | yes | Fireleaf Belt (240053, +0.00 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -2.28 DPS) [vendor]; Belt of the Archmage (18405, -2.53 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (208.7 DPS) | yes | Marshal's Silk Leggings (16442, -3.48 DPS) [vendor]; Marshal's Silk Leggings (231605, -3.48 DPS) [vendor]; Sentinel's Silk Leggings (237815, -6.41 DPS, sim-verified) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 145.2 spell_power points (17.89 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Bloodvine Boots (19684, -4.17 DPS) [crafted]; Marshal's Silk Footwraps (16437, -4.35 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (202.3 DPS) | yes | Wrath of Cenarius (21190, -1.17 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.99 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -4.26 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (202.3 DPS) | yes | Wrath of Cenarius (21190, -1.17 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.99 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -4.26 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (202.3 DPS) | yes | Uther's Strength (11302, -0.99 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.84 DPS, sim-verified) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (202.3 DPS) | yes | Uther's Strength (11302, -1.97 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.66 DPS, sim-verified) [crafted] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (202.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -5.24 DPS, sim-verified) [world_drop]; Grand Marshal's Stave (18873, -7.73 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (224.7 DPS) | yes | Oblivion's Touch (18761, -12.66 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.26 DPS) [world]; Cold Snap (19130, -22.42 DPS, sim-verified) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 952, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 34.2. Weights run: 0.7s. Verify run: 0.6s. 124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.502 ± 0.045, crit=0.654 ± 0.041, hit=1.133 ± 0.017, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -1.21 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.5 spell_power points (1.26 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.34 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Pearl-clasped Cloak (5542, -0.13 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.5 spell_power points (0.99 DPS) | yes | Green Woolen Robe (6243, -0.40 DPS) [crafted]; Green Woolen Vest (2582, -0.46 DPS) [crafted]; Gray Woolen Robe (2585, -0.92 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (34.1 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.07 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.54 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Gnoll Casting Gloves (892, -0.15 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.20 DPS) [crafted]; Apothecary Gloves (10919, -0.40 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Keller's Girdle (2911, -0.26 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.33 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (34.1 DPS) | yes | Silk-threaded Trousers (1929, -0.27 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.47 DPS, sim-verified) [dungeon]; Colorful Kilt (10048, -0.53 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Pristine Boots (253889, -0.59 DPS) [crafted]; Red Woolen Boots (4313, -0.66 DPS) [crafted]; Feather Padded Treads (285345, -0.98 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.66 DPS) | yes | Sludge-Stained Band (286535, -0.26 DPS) [world]; Loop of Sacrifice (281673, -0.33 DPS) [quest]; Volcanic Rock Ring (12053, -0.46 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 3.0 spell_power points (0.40 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; Volcanic Rock Ring (12053, -0.20 DPS) [world_drop]; Sludge-Stained Band (286535, -0.29 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 5.0 spell_power points (0.66 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.13 DPS) [world]; Lesser Staff of the Spire (1300, -0.27 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 171.9 spell_power points (22.71 DPS) | yes | Skycaller (12984, -0.97 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 58.0. Weights run: 0.7s. Verify run: 0.5s. 217 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.649 ± 0.085, crit=1.663 ± 0.106, hit=1.646 ± 0.025, spell_haste=0.154 ± 0.033, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.5 spell_power points (1.93 DPS) | yes | Holy Shroud (2721, -0.15 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.54 DPS) [crafted]; Embalmed Shroud (7691, -0.69 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.9 spell_power points (1.68 DPS) | yes | Crystal Starfire Medallion (5003, -1.28 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.28 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.41 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (2.29 DPS) | yes | Death Speaker Mantle (6685, -0.35 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Magician's Mantle (12998, -0.62 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 5.2 spell_power points (0.80 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.03 DPS) [crafted]; Windsong Drape (15468, -0.03 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.4 spell_power points (2.69 DPS) | yes | Death Speaker Robes (6682, -0.61 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -0.69 DPS) [dungeon]; Pristine Gown (253961, -0.91 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Nightsky Wristbands (6407, -0.79 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.79 DPS) [quest]; Glowing Magical Bracelets (13106, -1.38 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.2 spell_power points (1.43 DPS) | yes | Serpent Gloves (5970, -0.21 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.35 DPS) [crafted]; Gnoll Casting Gloves (892, -0.50 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.9 spell_power points (2.00 DPS) | yes | Warsong Sash (16975, -0.20 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.31 DPS) [dungeon]; Crimson Silk Belt (7055, -0.37 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (2.19 DPS) | yes | Gaze Dreamer Pants (6903, -0.06 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.41 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.78 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Boots of the Enchanter (4325, -1.01 DPS) [crafted]; Acidic Walkers (9454, -1.45 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.08 DPS) | yes | Advisor's Ring (20426, -0.31 DPS) [rep]; Black Widow Band (6199, -0.38 DPS) [world]; Snake Hoop (6750, -0.38 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.93 DPS) | yes | Snake Hoop (6750, -0.22 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Black Widow Band (6199, -1.45 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (58.0 DPS) | yes | Defiler's Talisman (21120, -2.02 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 7.1 spell_power points (1.10 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.10 DPS) [quest]; Twisted Chanter's Staff (890, -0.10 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.30 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 218.2 spell_power points (33.67 DPS) | yes | Starfaller (13063, -0.15 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.90 DPS) [crafted]; Gravestone Scepter (7001, -4.67 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 83.3. Weights run: 0.7s. Verify run: 0.5s. 300 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.160 ± 0.169), crit=2.753 ± 0.178, hit=2.986 ± 0.049, spell_haste=0.902 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.14 DPS) | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -1.50 DPS) [world_drop]; Silk Headband (7050, -1.80 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (1.19 DPS) | yes | Necklace of Calisea (1714, -1.02 DPS) [world_drop]; Triune Amulet (7722, -1.02 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.35 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (1.56 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.20 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (1.05 DPS) | yes | Long Silken Cloak (4326, +0.00 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.03 DPS) [crafted]; Battle Healer's Cloak (19528, -0.15 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.0 spell_power points (3.44 DPS) | yes | Dreamweave Vest (10021, -0.53 DPS) [crafted]; Elemental Raiment (9434, -0.68 DPS, sim-verified) [world_drop]; Robe of Power (7054, -1.05 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.35 DPS) | yes | Radiant Silver Bracers (4545, -0.56 DPS) [quest]; Earthen Silk Cuffs (254019, -0.75 DPS) [crafted]; Condor Bracers (15864, -0.83 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (2.79 DPS) | yes | Black Mageweave Gloves (10003, -0.74 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.90 DPS) [crafted]; Gilded Handwraps (254021, -1.42 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.6 spell_power points (2.19 DPS) | yes | Star Belt (4329, -0.08 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -0.47 DPS) [rep]; Warsong Sash (16975, -0.54 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.9 spell_power points (2.38 DPS) | yes | Abomination Skin Leggings (23173, -0.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.87 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.22 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | Gilded Slippers (254001, -0.73 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.45 DPS) [crafted]; Acidic Walkers (9454, -2.65 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (1.64 DPS) | yes | Reedknot Ring (9622, -0.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.74 DPS) [vendor]; Electrocutioner Lagnut (9447, -1.19 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.35 DPS) | yes | Advisor's Ring (19521, -0.30 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.45 DPS) [vendor]; Reedknot Ring (9622, -0.83 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (83.3 DPS) | yes | Gut Ripper (2164, -1.07 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -5.55 DPS) [dungeon]; Staff of Jordan (873, -5.65 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 264.1 spell_power points (39.52 DPS) | yes | Nether Force Wand (11263, -0.83 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.26 DPS) [quest]; Ragefire Wand (7513, -2.31 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 103.7. Weights run: 0.6s. Verify run: 0.6s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.275 ± 0.650), crit=5.153 ± 0.298, hit=5.744 ± 0.087, spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 101.4 spell_power points (13.31 DPS) | yes | Eye of Theradras (17715, -2.01 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -7.47 DPS) [crafted]; Chief Architect's Monocle (11839, -8.79 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.8 spell_power points (2.34 DPS) | yes | Mindburst Medallion (11196, -0.55 DPS) [quest]; Gemshard Heart (17707, -0.67 DPS) [dungeon]; Scorn's Icy Choker (23169, -1.17 DPS, sim-verified) [dungeon] |
| shoulder | Blood Guard's Dreadweave Mantle (220905) | Lady Palanseer [vendor] | 91.6 spell_power points (12.02 DPS) | yes | Rotgrip Mantle (17732, -0.23 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -8.59 DPS) [crafted]; Inquisitor's Shawl (19507, -8.93 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 23.5 spell_power points (3.08 DPS) | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.56 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.74 DPS) [vendor] |
| chest | Stone Guard's Dreadweave Vest (220904) | Lady Palanseer [vendor] | 98.2 spell_power points (12.88 DPS) | yes | Acumen Robes (17775, +0.00 DPS, sim-verified) [quest]; Runecloth Robe (13858, -8.59 DPS) [crafted]; Robes of Insight (940, -8.70 DPS) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 19.1 spell_power points (2.51 DPS) | yes | Bloodband Bracers (11469, +0.00 DPS, sim-verified) [quest]; Nethergeld Cuffs (254061, -0.42 DPS) [crafted]; Forgotten Wraps (9433, -0.50 DPS) [world_drop] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 87.3 spell_power points (11.45 DPS) | yes | Raider Handwraps (272098, -1.49 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -8.24 DPS) [vendor]; Red Mageweave Gloves (10018, -8.34 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 87.5 spell_power points (11.48 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.97 DPS) [dungeon]; Deathmage Sash (10771, -8.05 DPS) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 99.4 spell_power points (13.05 DPS) | yes | Red Mageweave Pants (10009, -0.52 DPS, sim-verified) [crafted]; Kilt of the Atal'ai Prophet (10807, -9.64 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -9.82 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 76.9 spell_power points (10.09 DPS) | yes | Earthen Silk Slippers (254013, -0.19 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -7.14 DPS) [crafted]; Southsea Mojo Boots (20641, -7.20 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 spell_power points (7.54 DPS) | yes | Brainlash (6440, -1.39 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -5.53 DPS) [dungeon]; Band of the Unicorn (7553, -5.83 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (103.7 DPS) | yes | Mindseye Circle (10634, -0.31 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Brainlash (6440, -1.69 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (102.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -0.02 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (102.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted]; Uther's Strength (11302, -4.49 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (102.0 DPS) | yes | Illusionary Rod (7713, -1.00 DPS) [dungeon]; Inventor's Focal Sword (17719, -2.01 DPS) [dungeon]; Shortsword of Vengeance (754, -4.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -14.82 DPS, sim-verified) [quest] |

**New at 50:** head: Blood Guard's Dreadweave Hat; neck: Horizon Choker; shoulder: Blood Guard's Dreadweave Mantle; back: Deep Woodlands Cloak; chest: Stone Guard's Dreadweave Vest; wrist: Aristocratic Cuffs; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 212.7. Weights run: 0.7s. Verify run: 0.8s. 946 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.725 ± 0.246, crit=5.578 ± 0.410, hit=6.478 ± 0.118, spell_haste=not significant (0.090 ± 0.410), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (181.3 DPS) | yes | Sorcerer's Crown (226935, -3.53 DPS) [quest]; Warlord's Silk Cowl (16533, -3.75 DPS) [vendor]; Bloodvine Goggles (19999, -5.81 DPS, sim-verified) [crafted] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (175.5 DPS) | yes | Beads of Ogre Might (22150, +0.00 DPS) [quest]; Amulet of the Dawn (22657, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -5.62 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 145.0 spell_power points (17.86 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -2.38 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -3.38 DPS) [crafted]; Shroud of the Nathrezim (18720, -3.98 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 94.6 spell_power points (11.65 DPS) | yes | Fel Cape (279269, -2.03 DPS) [crafted]; Earthweave Cloak (21187, -3.67 DPS) [quest]; Chromatic Cloak (18509, -6.41 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 252.0 spell_power points (31.04 DPS) | yes | Bloodvine Vest (19682, -7.05 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -9.16 DPS) [vendor]; Earthpower Vest (21183, -13.66 DPS) [quest] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 140.7 spell_power points (17.33 DPS) | yes | Fireleaf Wristwraps (240044, -2.34 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -6.02 DPS) [quest]; Dryad's Wrist Bindings (19595, -12.92 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 222.3 spell_power points (27.39 DPS) | yes | Gloves of Spell Mastery (14146, -6.73 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -8.79 DPS) [vendor]; Sorcerer's Gloves (22066, -14.95 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 146.2 spell_power points (18.01 DPS) | yes | Fireleaf Belt (240053, +0.00 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -2.28 DPS) [vendor]; Belt of the Archmage (18405, -2.53 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (181.4 DPS) | yes | General's Silk Trousers (16534, -3.48 DPS) [vendor]; General's Silk Trousers (231595, -3.48 DPS) [vendor]; Sentinel's Silk Leggings (237815, -5.92 DPS, sim-verified) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 145.2 spell_power points (17.89 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Bloodvine Boots (19684, -4.17 DPS) [crafted]; General's Silk Boots (16539, -4.35 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (175.5 DPS) | yes | Wrath of Cenarius (21190, -0.19 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.99 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -4.26 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (175.5 DPS) | yes | Wrath of Cenarius (21190, -0.19 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.99 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -4.26 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (175.5 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.99 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (175.5 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Uther's Strength (11302, -1.97 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.60 DPS, sim-verified) [crafted] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (175.5 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -6.03 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -9.19 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (200.4 DPS) | yes | Oblivion's Touch (18761, -12.66 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.26 DPS) [world]; Cold Snap (19130, -24.91 DPS, sim-verified) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 946, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

