# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 37.1. Weights run: 0.7s. Verify run: 0.5s. 128 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.654 ± 0.006, crit=0.654 ± 0.041, hit=1.133 ± 0.017, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -1.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (1.44 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.42 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.91 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Pearl-clasped Cloak (5542, -0.13 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.3 spell_power points (1.09 DPS) | yes | Green Woolen Robe (6243, -0.44 DPS) [crafted]; Mystic's Wrap (14369, -0.49 DPS) [world_drop]; Gray Woolen Robe (2585, -1.08 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.9 spell_power points (0.52 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.17 DPS) [world_drop]; Mystic's Bracelets (14366, -0.35 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Pristine Gloves (253913, -0.14 DPS) [crafted]; Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Blight Gloves (279877, -0.32 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.6 spell_power points (0.87 DPS) | yes | Keller's Girdle (2911, -0.18 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.35 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.69 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.88 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.95 DPS) [dungeon]; Colorful Kilt (10048, -1.22 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (1.27 DPS) | yes | Pristine Boots (253889, -0.61 DPS) [crafted]; Red Woolen Boots (4313, -0.74 DPS) [crafted]; Feather Padded Treads (285345, -0.84 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.3 spell_power points (0.83 DPS) | yes | Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Loop of Sacrifice (281673, -0.40 DPS) [quest]; Sludge-Stained Band (286535, -0.44 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.23 DPS) [quest]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.55 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 6.5 spell_power points (0.86 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.17 DPS) [world]; Lesser Staff of the Spire (1300, -0.35 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 171.9 spell_power points (22.71 DPS) | yes | Skycaller (12984, -0.89 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 128, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 65.0. Weights run: 0.7s. Verify run: 0.6s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.638 ± 0.009, crit=1.663 ± 0.106, hit=1.646 ± 0.025, spell_haste=0.154 ± 0.033, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.91 DPS) | yes | Holy Shroud (2721, -0.11 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.52 DPS) [crafted]; Embalmed Shroud (7691, -0.68 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.67 DPS) | yes | Crystal Starfire Medallion (5003, -1.28 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.28 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.53 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (2.28 DPS) | yes | Death Speaker Mantle (6685, -0.39 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Magician's Mantle (12998, -0.62 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.6 spell_power points (0.86 DPS) | yes | Cloak of Rot (4462, +0.00 DPS, sim-verified) [world]; Darkspear Raider's Cloak (272078, -0.07 DPS) [vendor]; Hillman's Cloak (3719, -0.09 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (2.67 DPS) | yes | Death Speaker Robes (6682, -0.42 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -0.66 DPS) [dungeon]; Pristine Gown (253961, -0.90 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Nightsky Wristbands (6407, -0.80 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.80 DPS) [quest]; Glowing Magical Bracelets (13106, -1.31 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 11.0 spell_power points (1.70 DPS) | yes | Shilly Mitts (9609, -0.62 DPS) [quest]; Truefaith Gloves (7049, -0.63 DPS) [crafted]; Serpent Gloves (5970, -0.82 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.9 spell_power points (1.99 DPS) | yes | Belt of Arugal (6392, -0.28 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.38 DPS) [crafted]; Invoker's Cord (215366, -0.42 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (2.18 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.41 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.77 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Nimbus Boots (6998, -0.84 DPS) [quest]; Acidic Walkers (9454, -1.66 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.08 DPS) | yes | Lorekeeper's Ring (20431, -0.31 DPS) [rep]; Black Widow Band (6199, -0.39 DPS) [world]; Minor Channeling Ring (1449, -1.71 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (65.0 DPS) | yes | Black Widow Band (6199, -0.24 DPS) [world]; Snake Hoop (6750, -0.24 DPS) [quest]; Minor Channeling Ring (1449, -1.49 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Talisman of Arathor (21119, -2.10 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Twisted Chanter's Staff (890, -0.40 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.40 DPS) [quest]; Glimmering Staff (249392, -0.61 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.8 spell_power points (1.67 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.97 DPS) [vendor]; Dwarven Tome (279898, -0.98 DPS, sim-verified) [quest]; Eye of Paleth (2943, -1.05 DPS) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 218.2 spell_power points (33.67 DPS) | yes | Starfaller (13063, -0.01 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.90 DPS) [crafted]; Gravestone Scepter (7001, -4.67 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 93.4. Weights run: 0.7s. Verify run: 0.6s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.863 ± 0.018, crit=2.753 ± 0.178, hit=2.986 ± 0.049, spell_haste=0.902 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.14 DPS) | yes | Corpseshroud (10574, -0.69 DPS) [dungeon]; Augural Shroud (2620, -0.93 DPS, sim-verified) [world]; Thinking Cap (2624, -0.95 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.2 spell_power points (1.82 DPS) | yes | Necklace of Calisea (1714, -0.92 DPS) [world_drop]; Triune Amulet (7722, -0.92 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.67 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.2 spell_power points (2.73 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.22 DPS) [dungeon]; Berylline Pads (4197, -0.39 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (93.4 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.12 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.68 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.2 spell_power points (4.07 DPS) | yes | Dreamweave Vest (10021, -0.15 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.42 DPS) [crafted]; Elemental Raiment (9434, -0.92 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.35 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.18 DPS) [world_drop]; Condor Bracers (15864, -0.30 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.5 spell_power points (3.21 DPS) | yes | Red Mageweave Gloves (10018, -0.96 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.97 DPS) [crafted]; Stormcloth Gloves (10011, -1.06 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.9 spell_power points (2.98 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.75 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.95 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.4 spell_power points (3.64 DPS) | yes | Crimson Silk Pantaloons (7062, -1.12 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.26 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.56 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | Gilded Slippers (254001, -1.42 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.81 DPS) [dungeon]; Spidersilk Boots (4320, -2.03 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.2 spell_power points (2.27 DPS) | yes | Ring of Forlorn Spirits (2043, -1.07 DPS) [quest]; Reedknot Ring (9622, -1.22 DPS) [quest]; Minor Channeling Ring (1449, -1.26 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.35 DPS) | yes | Ring of Forlorn Spirits (2043, -0.24 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.30 DPS) [quest]; Lorekeeper's Ring (19525, -0.30 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, -1.54 DPS, sim-verified) [world_drop]; Spellforce Rod (1664, -3.55 DPS) [world_drop]; Windweaver Staff (7757, -4.61 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 268.3 spell_power points (40.15 DPS) | yes | Nether Force Wand (11263, -1.23 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.58 DPS) [quest]; Ragefire Wand (7513, -2.63 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 122.9. Weights run: 0.5s. Verify run: 0.6s. 404 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.911 ± 0.050, crit=5.153 ± 0.298, hit=5.744 ± 0.087, spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 97.1 spell_power points (12.74 DPS) | yes | Eye of Theradras (17715, -0.44 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -7.85 DPS) [crafted]; Dreamweave Circlet (10041, -8.79 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (117.1 DPS) | yes | Scorn's Icy Choker (23169, -0.04 DPS) [dungeon]; Mindburst Medallion (11196, -0.17 DPS) [quest]; Arcane Crystal Pendant (20037, -2.60 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) | Captain Dirgehammer [vendor] | 88.3 spell_power points (11.59 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Kentic Amice (11624, -8.20 DPS) [dungeon]; Red Mageweave Shoulders (10029, -8.88 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.5 spell_power points (2.55 DPS) | yes | Runecloth Cloak (13860, -0.42 DPS) [crafted]; Big Voodoo Cloak (8216, -0.82 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -3.88 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (115.9 DPS) | yes | Robe of the Magi (1716, -1.28 DPS) [world_drop]; Runecloth Tunic (13857, -1.34 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.45 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (118.5 DPS) | yes | Bloodband Bracers (11469, -0.02 DPS) [quest]; Forgotten Wraps (9433, -0.32 DPS) [world_drop]; Aristocratic Cuffs (12546, -4.00 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 82.2 spell_power points (10.78 DPS) | yes | Raider Handwraps (272098, -0.56 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -7.94 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -8.00 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 85.7 spell_power points (11.24 DPS) | yes | Ban'thok Sash (11662, -3.44 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -8.19 DPS) [dungeon]; Satyrmane Sash (17755, -8.21 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 95.1 spell_power points (12.47 DPS) | yes | Spellshock Leggings (9484, -4.83 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -9.20 DPS) [crafted]; Crimson Silk Pantaloons (7062, -9.87 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 73.6 spell_power points (9.66 DPS) | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -7.14 DPS) [crafted]; Southsea Mojo Boots (20641, -7.30 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 spell_power points (7.54 DPS) | yes | Cyclopean Band (11824, -5.52 DPS) [dungeon]; Brainlash (6440, -5.74 DPS) [dungeon]; Band of the Unicorn (7553, -5.83 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.03 DPS) | yes | Brainlash (6440, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Cyclopean Band (11824, -2.85 DPS, sim-verified) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Thunderbrew's Boot Flask (744, -0.56 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.72 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.43 DPS) [dungeon]; Shortsword of Vengeance (754, -2.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -13.77 DPS, sim-verified) [quest] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 404, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 241.0. Weights run: 0.8s. Verify run: 0.7s. 953 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.272 ± 0.036, crit=5.578 ± 0.410, hit=6.478 ± 0.118, spell_haste=not significant (0.090 ± 0.410), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (211.0 DPS) | yes | Field Marshal's Coronet (16441, -3.64 DPS) [vendor]; Field Marshal's Coronet (231604, -3.64 DPS) [vendor]; Bloodvine Goggles (19999, -3.98 DPS, sim-verified) [crafted] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | 0.0 spell_power points (0.00 DPS) | yes | Beads of Ogre Might (22150, +0.00 DPS) [quest]; Amulet of the Dawn (22657, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -6.35 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 138.2 spell_power points (17.02 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -2.74 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -3.27 DPS) [crafted]; Lieutenant Commander's Silk Mantle (23319, -3.83 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 91.0 spell_power points (11.20 DPS) | yes | Fel Cape (279269, -1.58 DPS) [crafted]; Earthweave Cloak (21187, -3.22 DPS) [quest]; Chromatic Cloak (18509, -6.06 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 243.4 spell_power points (29.97 DPS) | yes | Bloodvine Vest (19682, -6.49 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -9.22 DPS) [vendor]; Robe of the Archmage (14152, -13.55 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 133.4 spell_power points (16.44 DPS) | yes | Fireleaf Wristwraps (240044, -1.57 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -5.13 DPS) [quest]; Dryad's Wrist Bindings (19595, -12.47 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 216.0 spell_power points (26.60 DPS) | yes | Gloves of Spell Mastery (14146, -6.65 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -8.85 DPS) [vendor]; Sorcerer's Gloves (22066, -14.95 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 139.9 spell_power points (17.23 DPS) | yes | Fireleaf Belt (240053, +0.00 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -2.64 DPS) [crafted]; Knowledge of the Timbermaw (228190, -2.73 DPS) [vendor] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (212.2 DPS) | yes | Marshal's Silk Leggings (16442, -3.54 DPS) [vendor]; Marshal's Silk Leggings (231605, -3.54 DPS) [vendor]; Sentinel's Silk Leggings (237815, -5.14 DPS, sim-verified) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 138.9 spell_power points (17.11 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Bloodvine Boots (19684, -4.28 DPS) [crafted]; Marshal's Silk Footwraps (16437, -4.35 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 0.0 spell_power points (0.00 DPS) | yes | Wrath of Cenarius (21190, -1.29 DPS, sim-verified) [quest]; Mindtear Band (20632, -4.33 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -4.76 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 0.0 spell_power points (0.00 DPS) | yes | Wrath of Cenarius (21190, -1.29 DPS, sim-verified) [quest]; Mindtear Band (20632, -4.33 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -4.76 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+9.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -1.18 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | 0.0 spell_power points (0.00 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -6.26 DPS, sim-verified) [world_drop]; Grand Marshal's Stave (18873, -8.46 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (229.9 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.29 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.89 DPS) [dungeon]; Cold Snap (19130, -22.80 DPS, sim-verified) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 953, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 34.2. Weights run: 0.7s. Verify run: 0.6s. 124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.654 ± 0.006, crit=0.654 ± 0.041, hit=1.133 ± 0.017, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -1.21 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (1.44 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.34 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.91 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Pearl-clasped Cloak (5542, -0.13 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.3 spell_power points (1.09 DPS) | yes | Green Woolen Robe (6243, -0.44 DPS) [crafted]; Mystic's Wrap (14369, -0.49 DPS) [world_drop]; Gray Woolen Robe (2585, -0.92 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (34.1 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.09 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.54 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Pristine Gloves (253913, -0.14 DPS) [crafted]; Gnoll Casting Gloves (892, -0.15 DPS, sim-verified) [world]; Blight Gloves (279877, -0.32 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.6 spell_power points (0.87 DPS) | yes | Keller's Girdle (2911, -0.18 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.35 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.71 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (34.1 DPS) | yes | Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.47 DPS, sim-verified) [dungeon]; Colorful Kilt (10048, -0.65 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (1.27 DPS) | yes | Pristine Boots (253889, -0.61 DPS) [crafted]; Red Woolen Boots (4313, -0.74 DPS) [crafted]; Feather Padded Treads (285345, -0.98 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.23 DPS) [quest]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 3.9 spell_power points (0.52 DPS) | yes | Sludge-Stained Band (286535, -0.12 DPS) [world]; Loop of Sacrifice (281673, -0.17 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 6.5 spell_power points (0.86 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.17 DPS) [world]; Lesser Staff of the Spire (1300, -0.35 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 171.9 spell_power points (22.71 DPS) | yes | Skycaller (12984, -0.97 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 58.9. Weights run: 0.7s. Verify run: 0.6s. 217 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.638 ± 0.009, crit=1.663 ± 0.106, hit=1.646 ± 0.025, spell_haste=0.154 ± 0.033, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.91 DPS) | yes | Holy Shroud (2721, -0.30 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.52 DPS) [crafted]; Embalmed Shroud (7691, -0.68 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.67 DPS) | yes | Crystal Starfire Medallion (5003, -1.28 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.28 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.56 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (2.28 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Magician's Mantle (12998, -0.62 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 5.1 spell_power points (0.79 DPS) | yes | Darkspear Raider's Cloak (272078, -0.01 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.02 DPS) [crafted]; Windsong Drape (15468, -0.02 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (2.67 DPS) | yes | Tree Bark Jacket (1486, -0.66 DPS) [dungeon]; Pristine Gown (253961, -0.90 DPS) [crafted]; Death Speaker Robes (6682, -0.99 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Nightsky Wristbands (6407, -0.80 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.80 DPS) [quest]; Glowing Magical Bracelets (13106, -1.55 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.2 spell_power points (1.42 DPS) | yes | Truefaith Gloves (7049, -0.35 DPS) [crafted]; Gnoll Casting Gloves (892, -0.49 DPS) [world]; Serpent Gloves (5970, -0.76 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.9 spell_power points (1.99 DPS) | yes | Belt of Arugal (6392, -0.31 DPS) [dungeon]; Crimson Silk Belt (7055, -0.38 DPS) [crafted]; Warsong Sash (16975, -0.73 DPS, sim-verified) [quest] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (2.18 DPS) | yes | Gaze Dreamer Pants (6903, -0.36 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.41 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.77 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Boots of the Enchanter (4325, -1.00 DPS) [crafted]; Acidic Walkers (9454, -1.49 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.08 DPS) | yes | Advisor's Ring (20426, -0.31 DPS) [rep]; Black Widow Band (6199, -0.39 DPS) [world]; Snake Hoop (6750, -0.39 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.93 DPS) | yes | Snake Hoop (6750, -0.24 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Black Widow Band (6199, -1.63 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (58.9 DPS) | yes | Defiler's Talisman (21120, -2.00 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Twisted Chanter's Staff (890, -0.40 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.40 DPS) [quest]; Glimmering Staff (249392, -0.91 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.8 spell_power points (1.67 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.97 DPS) [vendor]; Witch's Finger (16887, -0.98 DPS) [quest]; Dwarven Tome (279898, -1.35 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 218.2 spell_power points (33.67 DPS) | yes | Starfaller (13063, -0.34 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.90 DPS) [crafted]; Gravestone Scepter (7001, -4.67 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 85.5. Weights run: 0.7s. Verify run: 0.6s. 300 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.863 ± 0.018, crit=2.753 ± 0.178, hit=2.986 ± 0.049, spell_haste=0.902 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.14 DPS) | yes | Corpseshroud (10574, -0.69 DPS) [dungeon]; Augural Shroud (2620, -0.75 DPS, sim-verified) [world]; Thinking Cap (2624, -0.95 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.2 spell_power points (1.82 DPS) | yes | Necklace of Calisea (1714, -0.92 DPS) [world_drop]; Triune Amulet (7722, -0.92 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.81 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.2 spell_power points (2.73 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.22 DPS) [dungeon]; Berylline Pads (4197, -0.39 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (85.5 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.12 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.76 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.2 spell_power points (4.07 DPS) | yes | Dreamweave Vest (10021, -0.21 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.42 DPS) [crafted]; Elemental Raiment (9434, -0.92 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.9 spell_power points (1.63 DPS) | yes | Windchaser Cuffs (14429, -0.47 DPS) [world_drop]; Spidertank Oilrag (9448, -0.49 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.58 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.5 spell_power points (3.21 DPS) | yes | Red Mageweave Gloves (10018, -0.59 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.97 DPS) [crafted]; Stormcloth Gloves (10011, -1.06 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.9 spell_power points (2.98 DPS) | yes | Gilded Cord (254037, -0.75 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.95 DPS) [rep]; Defiler's Cloth Girdle (20166, -1.49 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.4 spell_power points (3.64 DPS) | yes | Crimson Silk Pantaloons (7062, -0.69 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.26 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.56 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.59 DPS) | yes | Gilded Slippers (254001, -0.84 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.81 DPS) [dungeon]; Spidersilk Boots (4320, -2.03 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.2 spell_power points (2.27 DPS) | yes | Reedknot Ring (9622, -1.22 DPS) [quest]; Ogremind Ring (1993, -1.37 DPS) [world_drop]; Voodoo Band (1996, -1.37 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.35 DPS) | yes | Advisor's Ring (19521, -0.30 DPS) [rep]; Ogremind Ring (1993, -0.44 DPS) [world_drop]; Reedknot Ring (9622, -1.19 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, -2.47 DPS, sim-verified) [world_drop]; Spellforce Rod (1664, -3.55 DPS) [world_drop]; Windweaver Staff (7757, -4.61 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 268.3 spell_power points (40.15 DPS) | yes | Nether Force Wand (11263, -2.03 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.58 DPS) [quest]; Ragefire Wand (7513, -2.63 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (troll, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 110.7. Weights run: 0.5s. Verify run: 0.6s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.911 ± 0.050, crit=5.153 ± 0.298, hit=5.744 ± 0.087, spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 97.1 spell_power points (12.74 DPS) | yes | Eye of Theradras (17715, -0.31 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -7.85 DPS) [crafted]; Dreamweave Circlet (10041, -8.79 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (105.0 DPS) | yes | Scorn's Icy Choker (23169, -0.04 DPS) [dungeon]; Mindburst Medallion (11196, -0.17 DPS) [quest]; Arcane Crystal Pendant (20037, -2.91 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (103.4 DPS) | yes | Kentic Amice (11624, -0.47 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.15 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.32 DPS, sim-verified) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.2 spell_power points (2.65 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.39 DPS) [dungeon]; Runecloth Cloak (13860, -0.51 DPS) [crafted]; Spritecaster Cape (11623, -0.65 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (103.7 DPS) | yes | Robe of the Magi (1716, -1.28 DPS) [world_drop]; Runecloth Tunic (13857, -1.34 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -1.63 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (105.8 DPS) | yes | Bloodband Bracers (11469, -0.02 DPS) [quest]; Radiant Silver Bracers (4545, -0.27 DPS) [quest]; Aristocratic Cuffs (12546, -3.74 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 82.2 spell_power points (10.78 DPS) | yes | Raider Handwraps (272098, -0.76 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -7.94 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -8.00 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 85.7 spell_power points (11.24 DPS) | yes | Ban'thok Sash (11662, -3.27 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -8.19 DPS) [dungeon]; Satyrmane Sash (17755, -8.21 DPS) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 95.1 spell_power points (12.47 DPS) | yes | Spellshock Leggings (9484, -4.59 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -9.20 DPS) [crafted]; Crimson Silk Pantaloons (7062, -9.87 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 73.6 spell_power points (9.66 DPS) | yes | Earthen Silk Slippers (254013, -0.50 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -7.14 DPS) [crafted]; Southsea Mojo Boots (20641, -7.30 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 spell_power points (7.54 DPS) | yes | Cyclopean Band (11824, -5.52 DPS) [dungeon]; Brainlash (6440, -5.74 DPS) [dungeon]; Band of the Unicorn (7553, -5.83 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.03 DPS) | yes | Brainlash (6440, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.32 DPS) [world_drop]; Cyclopean Band (11824, -3.10 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -4.49 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, -0.34 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.72 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.43 DPS) [dungeon]; Shortsword of Vengeance (754, -2.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.84 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -15.49 DPS, sim-verified) [quest] |

**New at 50:** head: Blood Guard's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (troll, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 216.4. Weights run: 0.8s. Verify run: 0.8s. 947 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.272 ± 0.036, crit=5.578 ± 0.410, hit=6.478 ± 0.118, spell_haste=not significant (0.090 ± 0.410), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (185.6 DPS) | yes | Warlord's Silk Cowl (16533, -3.64 DPS) [vendor]; Warlord's Silk Cowl (231601, -3.64 DPS) [vendor]; Bloodvine Goggles (19999, -5.99 DPS, sim-verified) [crafted] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | 0.0 spell_power points (0.00 DPS) | yes | Beads of Ogre Might (22150, +0.00 DPS) [quest]; Amulet of the Dawn (22657, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -4.81 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 138.2 spell_power points (17.02 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -2.38 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -3.27 DPS) [crafted]; Champion's Silk Mantle (23264, -3.83 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 91.0 spell_power points (11.20 DPS) | yes | Fel Cape (279269, -1.58 DPS) [crafted]; Earthweave Cloak (21187, -3.22 DPS) [quest]; Chromatic Cloak (18509, -6.59 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 243.4 spell_power points (29.97 DPS) | yes | Bloodvine Vest (19682, -7.38 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -9.22 DPS) [vendor]; Robe of the Archmage (14152, -13.55 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 133.4 spell_power points (16.44 DPS) | yes | Fireleaf Wristwraps (240044, -2.63 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -5.13 DPS) [quest]; Dryad's Wrist Bindings (19595, -12.47 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 216.0 spell_power points (26.60 DPS) | yes | Gloves of Spell Mastery (14146, -7.03 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -8.85 DPS) [vendor]; Sorcerer's Gloves (22066, -14.95 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 139.9 spell_power points (17.23 DPS) | yes | Fireleaf Belt (240053, -0.11 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -2.64 DPS) [crafted]; Knowledge of the Timbermaw (228190, -2.73 DPS) [vendor] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (186.0 DPS) | yes | General's Silk Trousers (16534, -3.54 DPS) [vendor]; General's Silk Trousers (231595, -3.54 DPS) [vendor]; Sentinel's Silk Leggings (237815, -6.31 DPS, sim-verified) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 138.9 spell_power points (17.11 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Bloodvine Boots (19684, -4.28 DPS) [crafted]; General's Silk Boots (16539, -4.35 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 0.0 spell_power points (0.00 DPS) | yes | Wrath of Cenarius (21190, -0.15 DPS, sim-verified) [quest]; Mindtear Band (20632, -4.33 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -4.76 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 0.0 spell_power points (0.00 DPS) | yes | Wrath of Cenarius (21190, -0.15 DPS, sim-verified) [quest]; Mindtear Band (20632, -4.33 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -4.76 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+7.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -1.02 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | 0.0 spell_power points (0.00 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -6.04 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -9.31 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (204.8 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -12.29 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.89 DPS) [dungeon]; Cold Snap (19130, -25.15 DPS, sim-verified) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 947, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

