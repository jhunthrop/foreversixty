# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 33.0. Weights run: 0.8s. Verify run: 0.6s. 128 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.967 ± 0.065, crit=0.863 ± 0.037, hit=1.647 ± 0.023, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.49 DPS) | yes | Shadow Goggles (4373, -0.80 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.7 spell_power points (1.12 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.20 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.79 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.9 spell_power points (0.40 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Sanguine Cape (14376, -0.08 DPS) [world_drop]; Feyscale Cloak (6632, -0.15 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.8 spell_power points (0.80 DPS) | yes | Mystic's Wrap (14369, -0.25 DPS) [world_drop]; Mystic's Robe (14371, -0.25 DPS) [world_drop]; Gray Woolen Robe (2585, -0.59 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (32.1 DPS) | yes | Bright Bracers (3647, -0.08 DPS) [world_drop]; Mystic's Bracelets (14366, -0.24 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.35 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (32.1 DPS) | yes | Blight Gloves (279877, -0.01 DPS) [quest]; Gnoll Casting Gloves (892, -0.07 DPS) [world]; Serpent Gloves (5970, -0.33 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.9 spell_power points (0.64 DPS) | yes | Novice Arcanist's Sash (253885, -0.08 DPS) [crafted]; Novice Ardent's Sash (253887, -0.24 DPS) [crafted]; Keller's Girdle (2911, -0.39 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (32.3 DPS) | yes | Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Darkweave Breeches (12987, -0.41 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.47 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (0.88 DPS) | yes | Pristine Boots (253889, -0.16 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.48 DPS) [world]; Red Woolen Boots (4313, -0.56 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 spell_power points (0.56 DPS) | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Sludge-Stained Band (286535, -0.32 DPS) [world]; Lavishly Jeweled Ring (1156, -0.60 DPS, sim-verified) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (32.2 DPS) | yes | Loop of Sacrifice (281673, -0.01 DPS) [quest]; Sludge-Stained Band (286535, -0.16 DPS) [world]; Lavishly Jeweled Ring (1156, -0.44 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 9.7 spell_power points (0.79 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.16 DPS) [world]; Lesser Staff of the Spire (1300, -0.31 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 277.1 spell_power points (22.55 DPS) | yes | Skycaller (12984, -0.66 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.26 DPS) [dungeon]; Deepblaze (279896, -4.02 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 128, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 57.2. Weights run: 0.8s. Verify run: 0.6s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.639 ± 0.067, crit=1.166 ± 0.044, hit=2.253 ± 0.046, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.43 DPS) | yes | Holy Shroud (2721, -0.23 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.39 DPS) [crafted]; Embalmed Shroud (7691, -0.51 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.25 DPS) | yes | Crystal Starfire Medallion (5003, -0.96 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.96 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.01 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (1.71 DPS) | yes | Fairywing Mantle (9536, -0.35 DPS) [quest]; Magician's Mantle (12998, -0.46 DPS) [world_drop]; Death Speaker Mantle (6685, -0.51 DPS, sim-verified) [dungeon] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.6 spell_power points (0.64 DPS) | yes | Darkspear Raider's Cloak (272078, -0.05 DPS) [vendor]; Hillman's Cloak (3719, -0.06 DPS) [crafted]; Cloak of Rot (4462, -0.67 DPS, sim-verified) [world] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (2.00 DPS) | yes | Death Speaker Robes (6682, -0.20 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -0.50 DPS) [dungeon]; Pristine Gown (253961, -0.68 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Nightsky Wristbands (6407, -0.60 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.60 DPS) [quest]; Glowing Magical Bracelets (13106, -1.36 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 11.0 spell_power points (1.28 DPS) | yes | Shilly Mitts (9609, -0.47 DPS) [quest]; Truefaith Gloves (7049, -0.48 DPS) [crafted]; Serpent Gloves (5970, -0.68 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.9 spell_power points (1.50 DPS) | yes | Belt of Arugal (6392, -0.19 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.28 DPS) [crafted]; Invoker's Cord (215366, -0.31 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (1.63 DPS) | yes | Gaze Dreamer Pants (6903, -0.05 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.31 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.50 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.33 DPS) | yes | Spidersilk Boots (4320, -0.22 DPS) [crafted]; Nimbus Boots (6998, -0.63 DPS) [quest]; Acidic Walkers (9454, -1.25 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.81 DPS) | yes | Lorekeeper's Ring (20431, -0.23 DPS) [rep]; Black Widow Band (6199, -0.29 DPS) [world]; Minor Channeling Ring (1449, -1.29 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (57.2 DPS) | yes | Black Widow Band (6199, -0.18 DPS) [world]; Snake Hoop (6750, -0.18 DPS) [quest]; Minor Channeling Ring (1449, -1.16 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (56.1 DPS) | yes | Talisman of Arathor (21119, -1.63 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 7.0 spell_power points (0.81 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.07 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 289.9 spell_power points (33.56 DPS) | yes | Starfaller (13063, -0.43 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.56 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 179.3. Weights run: 0.7s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.091 ± 0.015, crit=3.159 ± 0.088, hit=3.082 ± 0.038, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.20 DPS) | yes | Augural Shroud (2620, -2.99 DPS, sim-verified) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop]; Silk Headband (7050, -4.11 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 spell_power points (2.59 DPS) | yes | Necklace of Calisea (1714, -2.37 DPS) [world_drop]; Triune Amulet (7722, -2.37 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.71 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (3.37 DPS) | yes | Green Silken Shoulders (7057, -0.24 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Berylline Pads (4197, -0.65 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.40 DPS) | yes | Long Silken Cloak (4326, -0.08 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted]; Caretaker's Cape (19532, -0.34 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 spell_power points (7.73 DPS) | yes | Elemental Raiment (9434, -0.62 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.28 DPS) [crafted]; Robe of Power (7054, -2.56 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.08 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.69 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (6.29 DPS) | yes | Black Mageweave Gloves (10003, -1.22 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.4 spell_power points (4.92 DPS) | yes | Star Belt (4329, -0.53 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -1.06 DPS) [rep]; Belt of Arugal (6392, -1.75 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 spell_power points (5.17 DPS) | yes | Gaze Dreamer Pants (6903, -1.25 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.03 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.23 DPS) | yes | Gilded Slippers (254001, -3.26 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Nimbus Boots (6998, -6.17 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 spell_power points (3.61 DPS) | yes | Ring of Forlorn Spirits (2043, -0.87 DPS) [quest]; Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (3.08 DPS) | yes | Ring of Forlorn Spirits (2043, -0.35 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.69 DPS) [quest]; Lorekeeper's Ring (19525, -0.69 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (179.3 DPS) | yes | Gut Ripper (2164, -1.88 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -14.88 DPS) [dungeon]; Staff of Jordan (873, -15.00 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 115.4 spell_power points (39.57 DPS) | yes | Nether Force Wand (11263, -0.14 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.28 DPS) [quest]; Ragefire Wand (7513, -2.33 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 255.3. Weights run: 0.7s. Verify run: 0.7s. 403 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.020, crit=4.698 ± 0.137, hit=4.690 ± 0.060, spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 81.1 spell_power points (28.40 DPS) | yes | Eye of Theradras (17715, -5.03 DPS, sim-verified) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.94 DPS) [crafted]; Dreamweave Circlet (10041, -20.67 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (2.68 DPS) | yes | Mindburst Medallion (11196, -0.39 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.15 DPS) [world_drop]; Gemshard Heart (17707, -2.30 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) | Captain Dirgehammer [vendor] | 74.7 spell_power points (26.19 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -22.34 DPS) [crafted]; Bloodmage Mantle (7684, -22.69 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.13 DPS) | yes | Runecloth Cloak (13860, -1.81 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.98 DPS) [dungeon]; Icy Cloak (4327, -2.68 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) | Captain Dirgehammer [vendor] | 79.0 spell_power points (27.66 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Acumen Robes (17775, -20.25 DPS) [quest]; Elemental Raiment (9434, -20.31 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.15 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.44 DPS) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 60.4 spell_power points (21.16 DPS) | yes | Dreamweave Gloves (10019, -0.89 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -15.91 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -16.27 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 75.3 spell_power points (26.39 DPS) | yes | Satyrmane Sash (17755, -0.11 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -21.33 DPS) [rep]; Ghostweave Cord (254073, -21.48 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 79.1 spell_power points (27.70 DPS) | yes | Wizardweave Leggings (14132, -0.16 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -22.34 DPS) [crafted]; Senior Designer's Pantaloons (11841, -23.23 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (252.8 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -2.97 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -4.21 DPS) [crafted]; Black Mageweave Boots (10026, -4.29 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.9 spell_power points (16.43 DPS) | yes | Lorekeeper's Ring (19523, -12.23 DPS) [rep]; Philanthropist's Ring (281635, -12.70 DPS) [quest]; Lorekeeper's Ring (19524, -13.28 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.55 DPS) | yes | Lorekeeper's Ring (19523, -0.39 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.82 DPS) [quest]; Lorekeeper's Ring (19524, -1.40 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (252.5 DPS) | yes | Guardian Talisman (1490, -2.64 DPS, sim-verified) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (249.9 DPS) | yes | - |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (249.9 DPS) | yes | Illusionary Rod (7713, -0.23 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.45 DPS) [dungeon]; Shortsword of Vengeance (754, -2.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Woestave (20082, -1.18 DPS) [quest]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 512.0. Weights run: 0.7s. Verify run: 0.8s. 952 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.101 ± 0.023, crit=6.865 ± 0.198, hit=6.837 ± 0.088, spell_haste=4.459 ± 0.153, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (504.1 DPS) | yes | Field Marshal's Coronet (16441, -9.51 DPS) [vendor]; Field Marshal's Coronet (231604, -9.51 DPS) [vendor]; Bloodvine Goggles (19999, -13.22 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (490.9 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -9.70 DPS) [quest]; Orb of the Darkmoon (19426, -25.91 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 138.6 spell_power points (48.47 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -5.87 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -8.46 DPS) [crafted]; Lieutenant Commander's Silk Mantle (23319, -9.23 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 96.1 spell_power points (33.61 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Arcanoweave Cloak (272411, -3.82 DPS) [vendor]; Earthweave Cloak (21187, -9.70 DPS) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 257.1 spell_power points (89.92 DPS) | yes | Bloodvine Vest (19682, -13.84 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -32.87 DPS) [vendor]; Robe of the Archmage (14152, -41.90 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 132.7 spell_power points (46.41 DPS) | yes | Fireleaf Wristwraps (240044, -3.22 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -13.06 DPS) [quest]; Dryad's Wrist Bindings (19595, -38.44 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 235.6 spell_power points (82.40 DPS) | yes | Gloves of Spell Mastery (14146, -14.10 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -31.82 DPS) [vendor]; Sorcerer's Gloves (22066, -53.80 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 141.5 spell_power points (49.49 DPS) | yes | Fireleaf Belt (240053, -9.04 DPS) [vendor]; Belt of the Archmage (18405, -9.93 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -10.77 DPS) [rep] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (499.5 DPS) | yes | Sentinel's Silk Leggings (237815, -8.62 DPS, sim-verified) [vendor]; Marshal's Silk Leggings (16442, -10.46 DPS) [vendor]; Marshal's Silk Leggings (231605, -10.46 DPS) [vendor] |
| feet | Fireleaf Boots (240050) | Leonid Barthalomew the Revered [vendor] | 140.5 spell_power points (49.14 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (16437, -17.39 DPS) [vendor]; Marshal's Silk Footwraps (231606, -17.39 DPS) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (490.9 DPS) | yes | Mindtear Band (20632, +0.00 DPS) [world]; Ritssyn's Ring of Chaos (21836, +0.00 DPS) [world_drop]; Don Julio's Band (19325, -11.15 DPS, sim-verified) [rep] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (490.9 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Naglering (11669, -9.34 DPS, sim-verified) [dungeon]; Ritssyn's Ring of Chaos (21836, -15.17 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (490.9 DPS) | yes | Uther's Strength (11302, -2.80 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -6.55 DPS, sim-verified) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (490.9 DPS) | yes | Frozen Heart of the Mountain (249469, -4.01 DPS, sim-verified) [crafted]; Uther's Strength (11302, -5.60 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (490.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -8.27 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -33.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cold Snap (19130) | Azuregos [world] | 244.6 spell_power points (85.54 DPS) | yes | Torch of Light (279246, +0.00 DPS, sim-verified) [crafted]; Wand of Biting Cold (19108, -21.54 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, -21.69 DPS) [dungeon] |

**New at 60:** head: Fireleaf Circlet; neck: Medallion of the Dawn; shoulder: Fireleaf Shoulderpads; back: Chromatic Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Wrath of Cenarius; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Cold Snap

No-known-source sample (15 of 952, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 30.2. Weights run: 0.8s. Verify run: 0.6s. 124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.967 ± 0.065, crit=0.863 ± 0.037, hit=1.647 ± 0.023, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.49 DPS) | yes | Shadow Goggles (4373, -0.52 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (29.5 DPS) | yes | Scout's Medallion (20442, -0.30 DPS, sim-verified) [rep] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.7 spell_power points (1.12 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.28 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.79 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.9 spell_power points (0.40 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Sanguine Cape (14376, -0.08 DPS) [world_drop]; Feyscale Cloak (6632, -0.15 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.8 spell_power points (0.80 DPS) | yes | Mystic's Wrap (14369, -0.25 DPS) [world_drop]; Mystic's Robe (14371, -0.25 DPS) [world_drop]; Gray Woolen Robe (2585, -0.54 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (29.5 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.08 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.31 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (29.6 DPS) | yes | Blight Gloves (279877, -0.01 DPS) [quest]; Gnoll Casting Gloves (892, -0.07 DPS) [world]; Serpent Gloves (5970, -0.36 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.9 spell_power points (0.64 DPS) | yes | Novice Arcanist's Sash (253885, -0.08 DPS) [crafted]; Novice Ardent's Sash (253887, -0.24 DPS) [crafted]; Keller's Girdle (2911, -0.33 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (29.5 DPS) | yes | Abomination Skin Leggings (23173, -0.33 DPS, sim-verified) [dungeon]; Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Darkweave Breeches (12987, -0.41 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (0.88 DPS) | yes | Pristine Boots (253889, -0.03 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.48 DPS) [world]; Red Woolen Boots (4313, -0.56 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 5.8 spell_power points (0.47 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; Sludge-Stained Band (286535, -0.23 DPS) [world]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.41 DPS) | yes | Sludge-Stained Band (286535, -0.16 DPS) [world]; Volcanic Rock Ring (12053, -0.17 DPS) [world_drop]; Loop of Sacrifice (281673, -0.37 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 9.7 spell_power points (0.79 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.16 DPS) [world]; Lesser Staff of the Spire (1300, -0.31 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 277.1 spell_power points (22.55 DPS) | yes | Skycaller (12984, -0.63 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.26 DPS) [dungeon]; Deepblaze (279896, -4.02 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (orc, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 50.1. Weights run: 0.8s. Verify run: 0.6s. 217 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.639 ± 0.067, crit=1.166 ± 0.044, hit=2.253 ± 0.046, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.43 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Silk Headband (7050, -0.39 DPS) [crafted]; Embalmed Shroud (7691, -0.51 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.25 DPS) | yes | Darkspear Warding Pendant (272075, -0.70 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.96 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.96 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (1.71 DPS) | yes | Fairywing Mantle (9536, -0.35 DPS) [quest]; Death Speaker Mantle (6685, -0.40 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.46 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 5.1 spell_power points (0.59 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.01 DPS) [crafted]; Windsong Drape (15468, -0.01 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (2.00 DPS) | yes | Death Speaker Robes (6682, -0.07 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -0.50 DPS) [dungeon]; Pristine Gown (253961, -0.68 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.04 DPS) | yes | Nightsky Wristbands (6407, -0.60 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.60 DPS) [quest]; Glowing Magical Bracelets (13106, -0.98 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.2 spell_power points (1.06 DPS) | yes | Serpent Gloves (5970, -0.11 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.26 DPS) [crafted]; Gnoll Casting Gloves (892, -0.37 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.9 spell_power points (1.50 DPS) | yes | Warsong Sash (16975, +0.00 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.23 DPS) [dungeon]; Crimson Silk Belt (7055, -0.28 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (1.63 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.31 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.50 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.33 DPS) | yes | Spidersilk Boots (4320, -0.22 DPS) [crafted]; Boots of the Enchanter (4325, -0.75 DPS) [crafted]; Acidic Walkers (9454, -1.20 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.81 DPS) | yes | Advisor's Ring (20426, -0.23 DPS) [rep]; Black Widow Band (6199, -0.29 DPS) [world]; Snake Hoop (6750, -0.29 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.69 DPS) | yes | Snake Hoop (6750, -0.18 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Black Widow Band (6199, -0.91 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (50.1 DPS) | yes | Defiler's Talisman (21120, -1.28 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 7.0 spell_power points (0.81 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.07 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 289.9 spell_power points (33.56 DPS) | yes | Starfaller (13063, -0.29 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.56 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 172.6. Weights run: 0.7s. Verify run: 0.7s. 300 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.091 ± 0.015, crit=3.159 ± 0.088, hit=3.082 ± 0.038, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.20 DPS) | yes | Augural Shroud (2620, -2.89 DPS, sim-verified) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop]; Silk Headband (7050, -4.11 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 spell_power points (2.59 DPS) | yes | Necklace of Calisea (1714, -2.37 DPS) [world_drop]; Triune Amulet (7722, -2.37 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.62 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (3.37 DPS) | yes | Green Silken Shoulders (7057, -0.25 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Chestnut Mantle (17695, -0.62 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.40 DPS) | yes | Long Silken Cloak (4326, -0.13 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted]; Battle Healer's Cloak (19528, -0.34 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 spell_power points (7.73 DPS) | yes | Elemental Raiment (9434, -0.61 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.28 DPS) [crafted]; Robe of Power (7054, -2.56 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.08 DPS) | yes | Condor Bracers (15864, -0.67 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.46 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 spell_power points (6.29 DPS) | yes | Black Mageweave Gloves (10003, -1.17 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.4 spell_power points (4.92 DPS) | yes | Star Belt (4329, -0.50 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -1.06 DPS) [rep]; Warsong Sash (16975, -1.15 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 spell_power points (5.17 DPS) | yes | Gaze Dreamer Pants (6903, -1.25 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.03 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.23 DPS) | yes | Gilded Slippers (254001, -3.17 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Acidic Walkers (9454, -6.26 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 spell_power points (3.61 DPS) | yes | Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor]; Electrocutioner Lagnut (9447, -2.59 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (3.08 DPS) | yes | Reedknot Ring (9622, -0.67 DPS, sim-verified) [quest]; Advisor's Ring (19521, -0.69 DPS) [rep]; Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (172.6 DPS) | yes | Gut Ripper (2164, -1.90 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -14.88 DPS) [dungeon]; Staff of Jordan (873, -15.00 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 115.4 spell_power points (39.57 DPS) | yes | Nether Force Wand (11263, -0.12 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.28 DPS) [quest]; Ragefire Wand (7513, -2.33 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 245.2. Weights run: 0.7s. Verify run: 0.7s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.020, crit=4.698 ± 0.137, hit=4.690 ± 0.060, spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 81.1 spell_power points (28.40 DPS) | yes | Eye of Theradras (17715, -4.96 DPS, sim-verified) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.94 DPS) [crafted]; Dreamweave Circlet (10041, -20.67 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 spell_power points (2.68 DPS) | yes | Mindburst Medallion (11196, -0.38 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.15 DPS) [world_drop]; Gemshard Heart (17707, -2.30 DPS) [dungeon] |
| shoulder | Blood Guard's Dreadweave Mantle (220905) | Lady Palanseer [vendor] | 74.7 spell_power points (26.19 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -22.34 DPS) [crafted]; Bloodmage Mantle (7684, -22.69 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (5.13 DPS) | yes | Deep Woodlands Cloak (19121, -0.63 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.68 DPS) [crafted]; Nightfall Drape (12465, -1.98 DPS) [dungeon] |
| chest | Stone Guard's Dreadweave Vest (220904) | Lady Palanseer [vendor] | 79.0 spell_power points (27.66 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Acumen Robes (17775, -20.25 DPS) [quest]; Elemental Raiment (9434, -20.31 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.15 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest]; Bloodband Bracers (11469, -1.06 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 60.4 spell_power points (21.16 DPS) | yes | Dreamweave Gloves (10019, -0.60 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -15.91 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -16.27 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 75.3 spell_power points (26.39 DPS) | yes | Satyrmane Sash (17755, -0.08 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -21.33 DPS) [rep]; Ghostweave Cord (254073, -21.48 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 79.1 spell_power points (27.70 DPS) | yes | Wizardweave Leggings (14132, +0.00 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -22.34 DPS) [crafted]; Senior Designer's Pantaloons (11841, -23.23 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (245.2 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -3.17 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -4.21 DPS) [crafted]; Black Mageweave Boots (10026, -4.29 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.9 spell_power points (16.43 DPS) | yes | Advisor's Ring (19519, -12.23 DPS) [rep]; Philanthropist's Ring (281635, -12.70 DPS) [quest]; Advisor's Ring (19520, -13.28 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.55 DPS) | yes | Advisor's Ring (19519, -0.38 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.82 DPS) [quest]; Advisor's Ring (19520, -1.40 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (242.0 DPS) | yes | Uther's Strength (11302, -0.02 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (242.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (242.0 DPS) | yes | Illusionary Rod (7713, -0.23 DPS) [dungeon]; Inventor's Focal Sword (17719, -0.45 DPS) [dungeon]; Shortsword of Vengeance (754, -2.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Woestave (20082, -1.18 DPS) [quest]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted] |

**New at 50:** head: Blood Guard's Dreadweave Hat; shoulder: Blood Guard's Dreadweave Mantle; back: Spritecaster Cape; chest: Stone Guard's Dreadweave Vest; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 491.6. Weights run: 0.7s. Verify run: 0.8s. 946 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.101 ± 0.023, crit=6.865 ± 0.198, hit=6.837 ± 0.088, spell_haste=4.459 ± 0.153, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (479.1 DPS) | yes | Warlord's Silk Cowl (16533, -9.51 DPS) [vendor]; Warlord's Silk Cowl (231601, -9.51 DPS) [vendor]; Bloodvine Goggles (19999, -12.10 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (467.0 DPS) | yes | Jewel of Kajaro (19601, +0.00 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -9.70 DPS) [quest]; Orb of the Darkmoon (19426, -25.91 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 138.6 spell_power points (48.47 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -5.68 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -8.46 DPS) [crafted]; Champion's Silk Mantle (23264, -9.23 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 96.1 spell_power points (33.61 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Arcanoweave Cloak (272411, -3.82 DPS) [vendor]; Earthweave Cloak (21187, -9.70 DPS) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 257.1 spell_power points (89.92 DPS) | yes | Bloodvine Vest (19682, -14.28 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -32.87 DPS) [vendor]; Robe of the Archmage (14152, -41.90 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 132.7 spell_power points (46.41 DPS) | yes | Fireleaf Wristwraps (240044, -3.21 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -13.06 DPS) [quest]; Dryad's Wrist Bindings (19595, -38.44 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 235.6 spell_power points (82.40 DPS) | yes | Gloves of Spell Mastery (14146, -13.75 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -31.82 DPS) [vendor]; Sorcerer's Gloves (22066, -53.80 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 141.5 spell_power points (49.49 DPS) | yes | Fireleaf Belt (240053, -9.04 DPS) [vendor]; Belt of the Archmage (18405, -9.64 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -10.77 DPS) [rep] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (475.1 DPS) | yes | Sentinel's Silk Leggings (237815, -8.10 DPS, sim-verified) [vendor]; General's Silk Trousers (16534, -10.46 DPS) [vendor]; General's Silk Trousers (231595, -10.46 DPS) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 140.5 spell_power points (49.14 DPS) | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (16539, -17.39 DPS) [vendor]; General's Silk Boots (231597, -17.39 DPS) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (467.0 DPS) | yes | Mindtear Band (20632, +0.00 DPS) [world]; Ritssyn's Ring of Chaos (21836, +0.00 DPS) [world_drop]; Don Julio's Band (19325, -11.64 DPS, sim-verified) [rep] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (467.0 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Naglering (11669, -8.77 DPS, sim-verified) [dungeon]; Ritssyn's Ring of Chaos (21836, -15.17 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (467.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -2.43 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.80 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (467.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS, sim-verified) [vendor]; Uther's Strength (11302, -14.64 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (467.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -8.09 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -33.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (471.8 DPS) | yes | Cold Snap (19130, -4.73 DPS, sim-verified) [world]; Wand of Biting Cold (19108, -14.86 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, -15.01 DPS) [dungeon] |

**New at 60:** head: Fireleaf Circlet; neck: Medallion of the Dawn; shoulder: Fireleaf Shoulderpads; back: Chromatic Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Wrath of Cenarius; finger2: Band of Earthen Might; trinket1: Serenity Field; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 946, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

