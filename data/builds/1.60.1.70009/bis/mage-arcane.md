# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 32.7. Weights run: 0.8s. Verify run: 0.6s. 128 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.794 ± 0.008, crit=0.863 ± 0.037, hit=1.647 ± 0.023, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.49 DPS) | yes | Shadow Goggles (4373, -0.66 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.1 spell_power points (0.99 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.39 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.66 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.4 spell_power points (0.36 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Sanguine Cape (14376, -0.10 DPS) [world_drop]; Feyscale Cloak (6632, -0.11 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.0 spell_power points (0.73 DPS) | yes | Mystic's Wrap (14369, -0.28 DPS) [world_drop]; Mystic's Robe (14371, -0.28 DPS) [world_drop]; Gray Woolen Robe (2585, -0.67 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (32.6 DPS) | yes | Bright Bracers (3647, -0.06 DPS) [world_drop]; Mystic's Bracelets (14366, -0.19 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.37 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.57 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.08 DPS) [world]; Blight Gloves (279877, -0.12 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.2 spell_power points (0.58 DPS) | yes | Keller's Girdle (2911, -0.07 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.36 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (32.6 DPS) | yes | Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.39 DPS, sim-verified) [dungeon]; Darkweave Breeches (12987, -0.42 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.2 spell_power points (0.83 DPS) | yes | Pristine Boots (253889, -0.09 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.42 DPS) [world]; Red Woolen Boots (4313, -0.50 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.6 spell_power points (0.54 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon]; Loop of Sacrifice (281673, -0.21 DPS) [quest]; Sludge-Stained Band (286535, -0.29 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.41 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; Sludge-Stained Band (286535, -0.16 DPS) [world]; Lavishly Jeweled Ring (1156, -0.44 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.9 spell_power points (0.65 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.13 DPS) [world]; Lesser Staff of the Spire (1300, -0.26 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 277.1 spell_power points (22.55 DPS) | yes | Skycaller (12984, -0.43 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.26 DPS) [dungeon]; Deepblaze (279896, -4.02 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 128, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 58.1. Weights run: 0.8s. Verify run: 0.7s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.727 ± 0.010, crit=1.166 ± 0.044, hit=2.253 ± 0.046, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.3 spell_power points (1.54 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.49 DPS) [crafted]; Nightsky Cowl (4039, -0.53 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 spell_power points (1.32 DPS) | yes | Crystal Starfire Medallion (5003, -0.98 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.98 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.15 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.5 spell_power points (1.80 DPS) | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.35 DPS) [quest]; Magician's Mantle (12998, -0.46 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.9 spell_power points (0.68 DPS) | yes | Darkspear Raider's Cloak (272078, -0.01 DPS) [vendor]; Hillman's Cloak (3719, -0.11 DPS) [crafted]; Cloak of Rot (4462, -0.20 DPS, sim-verified) [world] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.5 spell_power points (2.14 DPS) | yes | Tree Bark Jacket (1486, -0.63 DPS) [dungeon]; Death Speaker Robes (6682, -0.63 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.74 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Nightsky Wristbands (6407, -0.54 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.54 DPS) [quest]; Glowing Magical Bracelets (13106, -0.93 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 12.0 spell_power points (1.39 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.58 DPS) [dungeon]; Shilly Mitts (9609, -0.58 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.2 spell_power points (1.53 DPS) | yes | Belt of Arugal (6392, +0.00 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.24 DPS) [crafted]; Invoker's Cord (215366, -0.29 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (57.5 DPS) | yes | Gaze Dreamer Pants (6903, -0.01 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.20 DPS) [crafted]; Abomination Skin Leggings (23173, -1.07 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.1 spell_power points (1.40 DPS) | yes | Spidersilk Boots (4320, -0.25 DPS) [crafted]; Nimbus Boots (6998, -0.71 DPS) [quest]; Acidic Walkers (9454, -1.04 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.81 DPS) | yes | Black Widow Band (6199, -0.22 DPS) [world]; Snake Hoop (6750, -0.22 DPS) [quest]; Minor Channeling Ring (1449, -1.55 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (57.1 DPS) | yes | Black Widow Band (6199, -0.11 DPS) [world]; Snake Hoop (6750, -0.11 DPS) [quest]; Minor Channeling Ring (1449, -0.69 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Talisman of Arathor (21119, -1.11 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Twisted Chanter's Staff (890, -0.20 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.20 DPS) [quest]; Glimmering Staff (249392, -0.36 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.4 spell_power points (1.32 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.75 DPS) [vendor]; Dwarven Tome (279898, -0.80 DPS, sim-verified) [quest]; Satyr's Rod (15962, -0.81 DPS) [world_drop] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 289.9 spell_power points (33.56 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.56 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 179.3. Weights run: 0.7s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.094 ± 0.003, crit=3.159 ± 0.088, hit=3.082 ± 0.038, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.20 DPS) | yes | Living Cowl (5608, -2.80 DPS, sim-verified) [world]; Augural Shroud (2620, -3.11 DPS) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (2.59 DPS) | yes | Necklace of Calisea (1714, -2.37 DPS) [world_drop]; Triune Amulet (7722, -2.37 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.71 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (3.37 DPS) | yes | Green Silken Shoulders (7057, -0.24 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Berylline Pads (4197, -0.65 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (179.3 DPS) | yes | Long Silken Cloak (4326, -0.18 DPS) [crafted]; Guardian Cloak (5965, -0.18 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -2.45 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.6 spell_power points (7.73 DPS) | yes | Elemental Raiment (9434, -0.62 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.27 DPS) [crafted]; Robe of Power (7054, -2.55 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.08 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.69 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (6.30 DPS) | yes | Black Mageweave Gloves (10003, -1.22 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.4 spell_power points (4.93 DPS) | yes | Star Belt (4329, -0.52 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -1.06 DPS) [rep]; Belt of Arugal (6392, -1.75 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 spell_power points (5.19 DPS) | yes | Gaze Dreamer Pants (6903, -1.24 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.02 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.23 DPS) | yes | Gilded Slippers (254001, -3.27 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Nimbus Boots (6998, -6.17 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.6 spell_power points (3.62 DPS) | yes | Ring of Forlorn Spirits (2043, -0.88 DPS) [quest]; Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (3.08 DPS) | yes | Ring of Forlorn Spirits (2043, -0.35 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.69 DPS) [quest]; Lorekeeper's Ring (19525, -0.69 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, -1.86 DPS, sim-verified) [world_drop]; Spellforce Rod (1664, -8.50 DPS) [world_drop]; Scorn's Focal Dagger (23168, -12.27 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 115.4 spell_power points (39.57 DPS) | yes | Nether Force Wand (11263, -0.14 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.29 DPS) [quest]; Ragefire Wand (7513, -2.34 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 255.3. Weights run: 0.7s. Verify run: 0.7s. 404 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.005, crit=4.698 ± 0.137, hit=4.690 ± 0.060, spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 81.1 spell_power points (28.40 DPS) | yes | Eye of Theradras (17715, -5.13 DPS, sim-verified) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.94 DPS) [crafted]; Dreamweave Circlet (10041, -20.67 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (252.5 DPS) | yes | Mindburst Medallion (11196, -0.35 DPS) [quest]; Horizon Choker (13085, -2.15 DPS) [world_drop]; Arcane Crystal Pendant (20037, -3.02 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) | Captain Dirgehammer [vendor] | 74.7 spell_power points (26.19 DPS) | yes | Kentic Amice (11624, -5.63 DPS, sim-verified) [dungeon]; Rotgrip Mantle (17732, -20.95 DPS) [dungeon]; Black Mageweave Shoulders (10027, -22.34 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.13 DPS) | yes | Runecloth Cloak (13860, -1.68 DPS) [crafted]; Nightfall Drape (12465, -1.98 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -5.75 DPS, sim-verified) [dungeon] |
| chest | Knight's Dreadweave Vest (220886) | Captain Dirgehammer [vendor] | 79.0 spell_power points (27.66 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Acumen Robes (17775, -20.25 DPS) [quest]; Elemental Raiment (9434, -20.31 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.15 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.44 DPS) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 60.4 spell_power points (21.16 DPS) | yes | Dreamweave Gloves (10019, -1.18 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -15.91 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -16.27 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 75.3 spell_power points (26.39 DPS) | yes | Ban'thok Sash (11662, -5.86 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -21.10 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -21.33 DPS) [rep] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 79.1 spell_power points (27.70 DPS) | yes | Spellshock Leggings (9484, -7.30 DPS, sim-verified) [dungeon]; Wizardweave Leggings (14132, -21.04 DPS) [crafted]; Red Mageweave Pants (10009, -22.34 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (252.2 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -2.70 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -4.21 DPS) [crafted]; Black Mageweave Boots (10026, -4.29 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.9 spell_power points (16.43 DPS) | yes | Lorekeeper's Ring (19523, -12.23 DPS) [rep]; Philanthropist's Ring (281635, -12.70 DPS) [quest]; Cyclopean Band (11824, -13.01 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.55 DPS) | yes | Lorekeeper's Ring (19523, -0.39 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.82 DPS) [quest]; Cyclopean Band (11824, -1.14 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Smoking Heart of the Mountain (11811, -2.63 DPS, sim-verified) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.23 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.45 DPS) [dungeon]; Shortsword of Vengeance (754, -2.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Woestave (20082, -1.18 DPS) [quest]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 404, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 514.3. Weights run: 0.7s. Verify run: 0.7s. 953 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.123 ± 0.006, crit=6.865 ± 0.198, hit=6.837 ± 0.088, spell_haste=4.459 ± 0.153, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (501.6 DPS) | yes | Field Marshal's Coronet (16441, -9.53 DPS) [vendor]; Field Marshal's Coronet (231604, -9.53 DPS) [vendor]; Bloodvine Goggles (19999, -13.28 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points (0.00 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -9.70 DPS) [quest]; Orb of the Darkmoon (19426, -25.91 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 139.0 spell_power points (48.59 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -5.87 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -8.48 DPS) [crafted]; Lieutenant Commander's Silk Mantle (23319, -9.26 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 96.1 spell_power points (33.61 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Arcanoweave Cloak (272411, -3.76 DPS) [vendor]; Earthweave Cloak (21187, -9.70 DPS) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 257.6 spell_power points (90.07 DPS) | yes | Bloodvine Vest (19682, -13.78 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -32.87 DPS) [vendor]; Robe of the Archmage (14152, -41.95 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 133.1 spell_power points (46.54 DPS) | yes | Fireleaf Wristwraps (240044, -3.21 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -13.19 DPS) [quest]; Black Bark Wristbands (20626, -37.62 DPS) [world] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 235.9 spell_power points (82.51 DPS) | yes | Gloves of Spell Mastery (14146, -14.10 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -31.82 DPS) [vendor]; Sorcerer's Gloves (22066, -53.80 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 141.8 spell_power points (49.60 DPS) | yes | Fireleaf Belt (240053, -9.04 DPS) [vendor]; Belt of the Archmage (18405, -9.93 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -10.84 DPS) [rep] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (496.9 DPS) | yes | Sentinel's Silk Leggings (237815, -8.63 DPS, sim-verified) [vendor]; Marshal's Silk Leggings (16442, -10.45 DPS) [vendor]; Marshal's Silk Leggings (231605, -10.45 DPS) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 140.8 spell_power points (49.25 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (16437, -17.39 DPS) [vendor]; Marshal's Silk Footwraps (231606, -17.39 DPS) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | 0.0 spell_power points (0.00 DPS) | yes | Mindtear Band (20632, +0.00 DPS) [world]; Ritssyn's Ring of Chaos (21836, +0.00 DPS) [world_drop]; Don Julio's Band (19325, -11.27 DPS, sim-verified) [rep] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 0.0 spell_power points (0.00 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Naglering (11669, -9.31 DPS, sim-verified) [dungeon]; Ritssyn's Ring of Chaos (21836, -15.17 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -2.51 DPS, sim-verified) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | 0.0 spell_power points (0.00 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -8.23 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -33.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (493.2 DPS) | yes | Cold Snap (19130, -4.89 DPS, sim-verified) [world]; Ritssyn's Wand of Bad Mojo (22408, -11.16 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.21 DPS) [dungeon] |

**New at 60:** head: Fireleaf Circlet; neck: Medallion of the Dawn; shoulder: Fireleaf Shoulderpads; back: Chromatic Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Wrath of Cenarius; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Talisman of Ascendance; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 953, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 30.2. Weights run: 0.8s. Verify run: 0.6s. 124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.794 ± 0.008, crit=0.863 ± 0.037, hit=1.647 ± 0.023, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.49 DPS) | yes | Shadow Goggles (4373, -0.52 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (29.5 DPS) | yes | Scout's Medallion (20442, -0.30 DPS, sim-verified) [rep] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.1 spell_power points (0.99 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.28 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.66 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.4 spell_power points (0.36 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Sanguine Cape (14376, -0.10 DPS) [world_drop]; Feyscale Cloak (6632, -0.11 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.0 spell_power points (0.73 DPS) | yes | Mystic's Wrap (14369, -0.28 DPS) [world_drop]; Mystic's Robe (14371, -0.28 DPS) [world_drop]; Gray Woolen Robe (2585, -0.54 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (29.5 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.06 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.31 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (29.6 DPS) | yes | Gnoll Casting Gloves (892, -0.03 DPS) [world]; Blight Gloves (279877, -0.07 DPS) [quest]; Serpent Gloves (5970, -0.36 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.2 spell_power points (0.58 DPS) | yes | Keller's Girdle (2911, -0.07 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.26 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (29.5 DPS) | yes | Silk-threaded Trousers (1929, -0.31 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.33 DPS, sim-verified) [dungeon]; Darkweave Breeches (12987, -0.42 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.2 spell_power points (0.83 DPS) | yes | Pristine Boots (253889, -0.03 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.42 DPS) [world]; Red Woolen Boots (4313, -0.50 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.41 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; Sludge-Stained Band (286535, -0.16 DPS) [world]; Volcanic Rock Ring (12053, -0.21 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 4.8 spell_power points (0.39 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -0.14 DPS) [world]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.9 spell_power points (0.65 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.13 DPS) [world]; Lesser Staff of the Spire (1300, -0.26 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 277.1 spell_power points (22.55 DPS) | yes | Skycaller (12984, -0.63 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.26 DPS) [dungeon]; Deepblaze (279896, -4.02 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (orc, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 51.5. Weights run: 0.8s. Verify run: 0.7s. 217 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.727 ± 0.010, crit=1.166 ± 0.044, hit=2.253 ± 0.046, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.3 spell_power points (1.54 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.49 DPS) [crafted]; Nightsky Cowl (4039, -0.53 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 spell_power points (1.32 DPS) | yes | Darkspear Warding Pendant (272075, -0.89 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.98 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.98 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.5 spell_power points (1.80 DPS) | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.35 DPS) [quest]; Magician's Mantle (12998, -0.46 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 5.8 spell_power points (0.67 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.09 DPS) [crafted]; Windsong Drape (15468, -0.09 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.5 spell_power points (2.14 DPS) | yes | Death Speaker Robes (6682, -0.35 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -0.63 DPS) [dungeon]; Pristine Gown (253961, -0.74 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Nightsky Wristbands (6407, -0.54 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.54 DPS) [quest]; Glowing Magical Bracelets (13106, -0.96 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.6 spell_power points (1.12 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.31 DPS) [dungeon]; Pristine Gloves (253913, -0.40 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.2 spell_power points (1.53 DPS) | yes | Belt of Arugal (6392, -0.15 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.24 DPS) [crafted]; Warsong Sash (16975, -0.25 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (51.5 DPS) | yes | Gaze Dreamer Pants (6903, -0.01 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.20 DPS) [crafted]; Abomination Skin Leggings (23173, -0.75 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.1 spell_power points (1.40 DPS) | yes | Spidersilk Boots (4320, -0.25 DPS) [crafted]; Pristine Boots (253889, -0.80 DPS) [crafted]; Acidic Walkers (9454, -0.83 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.81 DPS) | yes | Black Widow Band (6199, -0.22 DPS) [world]; Snake Hoop (6750, -0.22 DPS) [quest]; Advisor's Ring (20426, -0.23 DPS) [rep] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Snake Hoop (6750, -0.11 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Black Widow Band (6199, -1.27 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Defiler's Talisman (21120, -1.44 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Twisted Chanter's Staff (890, -0.20 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.20 DPS) [quest]; Glimmering Staff (249392, -0.62 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.4 spell_power points (1.32 DPS) | yes | Dwarven Tome (279898, -0.74 DPS) [quest]; Tome of the Darkspear Prophecy (272090, -0.75 DPS) [vendor]; Witch's Finger (16887, -1.10 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 289.9 spell_power points (33.56 DPS) | yes | Starfaller (13063, -0.05 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.56 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 172.6. Weights run: 0.7s. Verify run: 0.7s. 300 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.094 ± 0.003, crit=3.159 ± 0.088, hit=3.082 ± 0.038, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.20 DPS) | yes | Living Cowl (5608, -2.68 DPS, sim-verified) [world]; Augural Shroud (2620, -3.11 DPS) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (2.59 DPS) | yes | Necklace of Calisea (1714, -2.37 DPS) [world_drop]; Triune Amulet (7722, -2.37 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.62 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (3.37 DPS) | yes | Green Silken Shoulders (7057, -0.25 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Chestnut Mantle (17695, -0.63 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (172.6 DPS) | yes | Long Silken Cloak (4326, -0.18 DPS) [crafted]; Guardian Cloak (5965, -0.18 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -2.35 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.6 spell_power points (7.73 DPS) | yes | Elemental Raiment (9434, -0.61 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.27 DPS) [crafted]; Robe of Power (7054, -2.55 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.08 DPS) | yes | Condor Bracers (15864, -0.67 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.46 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (6.30 DPS) | yes | Black Mageweave Gloves (10003, -1.17 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.4 spell_power points (4.93 DPS) | yes | Star Belt (4329, -0.50 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -1.06 DPS) [rep]; Warsong Sash (16975, -1.16 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 spell_power points (5.19 DPS) | yes | Gaze Dreamer Pants (6903, -1.25 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.02 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.23 DPS) | yes | Gilded Slippers (254001, -3.17 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Acidic Walkers (9454, -6.26 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.6 spell_power points (3.62 DPS) | yes | Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor]; Electrocutioner Lagnut (9447, -2.59 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (3.08 DPS) | yes | Reedknot Ring (9622, -0.67 DPS, sim-verified) [quest]; Advisor's Ring (19521, -0.69 DPS) [rep]; Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, -1.87 DPS, sim-verified) [world_drop]; Spellforce Rod (1664, -8.50 DPS) [world_drop]; Scorn's Focal Dagger (23168, -12.27 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 115.4 spell_power points (39.57 DPS) | yes | Nether Force Wand (11263, -0.12 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.29 DPS) [quest]; Ragefire Wand (7513, -2.34 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 244.5. Weights run: 0.7s. Verify run: 0.7s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.005, crit=4.698 ± 0.137, hit=4.690 ± 0.060, spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 81.1 spell_power points (28.40 DPS) | yes | Eye of Theradras (17715, -4.91 DPS, sim-verified) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.94 DPS) [crafted]; Dreamweave Circlet (10041, -20.67 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (242.0 DPS) | yes | Mindburst Medallion (11196, -0.35 DPS) [quest]; Horizon Choker (13085, -2.15 DPS) [world_drop]; Arcane Crystal Pendant (20037, -2.90 DPS, sim-verified) [quest] |
| shoulder | Blood Guard's Dreadweave Mantle (220905) | Lady Palanseer [vendor] | 74.7 spell_power points (26.19 DPS) | yes | Kentic Amice (11624, -5.53 DPS, sim-verified) [dungeon]; Rotgrip Mantle (17732, -20.95 DPS) [dungeon]; Black Mageweave Shoulders (10027, -22.34 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.13 DPS) | yes | Deep Woodlands Cloak (19121, -0.58 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -1.64 DPS) [dungeon]; Runecloth Cloak (13860, -1.68 DPS) [crafted] |
| chest | Stone Guard's Dreadweave Vest (220904) | Lady Palanseer [vendor] | 79.0 spell_power points (27.66 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Acumen Robes (17775, -20.25 DPS) [quest]; Elemental Raiment (9434, -20.31 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.15 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest]; Bloodband Bracers (11469, -1.06 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 60.4 spell_power points (21.16 DPS) | yes | Dreamweave Gloves (10019, -1.20 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -15.91 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -16.27 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 75.3 spell_power points (26.39 DPS) | yes | Ban'thok Sash (11662, -5.68 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -21.10 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -21.33 DPS) [rep] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 79.1 spell_power points (27.70 DPS) | yes | Spellshock Leggings (9484, -7.18 DPS, sim-verified) [dungeon]; Wizardweave Leggings (14132, -21.04 DPS) [crafted]; Red Mageweave Pants (10009, -22.34 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (241.6 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.54 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -4.21 DPS) [crafted]; Black Mageweave Boots (10026, -4.29 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.9 spell_power points (16.43 DPS) | yes | Advisor's Ring (19519, -12.23 DPS) [rep]; Philanthropist's Ring (281635, -12.70 DPS) [quest]; Cyclopean Band (11824, -13.01 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.55 DPS) | yes | Advisor's Ring (19519, -0.37 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.82 DPS) [quest]; Cyclopean Band (11824, -1.14 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -0.04 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.23 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.45 DPS) [dungeon]; Shortsword of Vengeance (754, -2.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Woestave (20082, -1.18 DPS) [quest]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted] |

**New at 50:** head: Blood Guard's Dreadweave Hat; shoulder: Blood Guard's Dreadweave Mantle; back: Spritecaster Cape; chest: Stone Guard's Dreadweave Vest; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 494.3. Weights run: 0.7s. Verify run: 0.8s. 947 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.123 ± 0.006, crit=6.865 ± 0.198, hit=6.837 ± 0.088, spell_haste=4.459 ± 0.153, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (481.9 DPS) | yes | Warlord's Silk Cowl (16533, -9.53 DPS) [vendor]; Warlord's Silk Cowl (231601, -9.53 DPS) [vendor]; Bloodvine Goggles (19999, -12.99 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points (0.00 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -9.70 DPS) [quest]; Orb of the Darkmoon (19426, -25.91 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 139.0 spell_power points (48.59 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -5.64 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -8.48 DPS) [crafted]; Champion's Silk Mantle (23264, -9.26 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 96.1 spell_power points (33.61 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Arcanoweave Cloak (272411, -3.76 DPS) [vendor]; Earthweave Cloak (21187, -9.70 DPS) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 257.6 spell_power points (90.07 DPS) | yes | Bloodvine Vest (19682, -13.33 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -32.87 DPS) [vendor]; Robe of the Archmage (14152, -41.95 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 133.1 spell_power points (46.54 DPS) | yes | Fireleaf Wristwraps (240044, -3.19 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -13.19 DPS) [quest]; Black Bark Wristbands (20626, -37.62 DPS) [world] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 235.9 spell_power points (82.51 DPS) | yes | Gloves of Spell Mastery (14146, -13.65 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -31.82 DPS) [vendor]; Sorcerer's Gloves (22066, -53.80 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 141.8 spell_power points (49.60 DPS) | yes | Fireleaf Belt (240053, -9.04 DPS) [vendor]; Belt of the Archmage (18405, -9.56 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -10.84 DPS) [rep] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (477.2 DPS) | yes | Sentinel's Silk Leggings (237815, -8.25 DPS, sim-verified) [vendor]; General's Silk Trousers (16534, -10.45 DPS) [vendor]; General's Silk Trousers (231595, -10.45 DPS) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 140.8 spell_power points (49.25 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (16539, -17.39 DPS) [vendor]; General's Silk Boots (231597, -17.39 DPS) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | 0.0 spell_power points (0.00 DPS) | yes | Mindtear Band (20632, +0.00 DPS) [world]; Ritssyn's Ring of Chaos (21836, +0.00 DPS) [world_drop]; Don Julio's Band (19325, -10.94 DPS, sim-verified) [rep] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 0.0 spell_power points (0.00 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Naglering (11669, -9.64 DPS, sim-verified) [dungeon]; Ritssyn's Ring of Chaos (21836, -15.17 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -2.42 DPS, sim-verified) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points (0.00 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | 0.0 spell_power points (0.00 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -7.96 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -33.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (473.6 DPS) | yes | Cold Snap (19130, -4.70 DPS, sim-verified) [world]; Ritssyn's Wand of Bad Mojo (22408, -11.16 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.21 DPS) [dungeon] |

**New at 60:** head: Fireleaf Circlet; neck: Medallion of the Dawn; shoulder: Fireleaf Shoulderpads; back: Chromatic Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Wrath of Cenarius; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Talisman of Ascendance; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 947, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

