# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 32.6. Weights run: 0.7s. Verify run: 0.7s. 123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.395 ± 0.090, crit=1.480 ± 0.048, hit=3.723 ± 0.108, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -0.49 DPS) [crafted]; Lucky Fishing Hat (19972, -0.49 DPS) [quest]; Shadow Goggles (4373, -0.85 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 | yes | Double-Stitched Woolen Shoulders (4314, -0.37 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.52 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.70 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.19 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.23 DPS) [crafted]; Green Woolen Vest (2582, -0.24 DPS) [crafted]; Gray Woolen Robe (2585, -0.79 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.06 DPS) [world_drop]; Windsong Bangles (263336, -0.11 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.15 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.35 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.20 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.59 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.42 DPS) [dungeon]; Colorful Kilt (10048, -0.59 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Pristine Boots (253889, -0.36 DPS) [crafted]; Red Woolen Boots (4313, -0.38 DPS) [crafted]; Feather Padded Treads (285345, -0.46 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Sludge-Stained Band (286535, -0.23 DPS) [world]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Loop of Sacrifice (281673, -0.31 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Loop of Sacrifice (281673, -0.25 DPS) [quest]; Sludge-Stained Band (286535, -0.43 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (32.6 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Staff of Westfall (2042, -0.10 DPS) [quest]; Living Root (6631, -0.49 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 275.0 | yes | Skycaller (12984, -0.75 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.27 DPS) [dungeon]; Deepblaze (279896, -4.03 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 123, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 20434 Lorekeeper's Staff; 20440 Protector's Sword; 20443 Sentinel's Blade; 209618 Insignia of the Alliance; 209623 Insignia of the Horde

### Band 30 (gnome, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 57.8. Weights run: 0.8s. Verify run: 0.8s. 209 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.236 ± 0.191), crit=1.585 ± 0.050, hit=3.984 ± 0.329, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.34 DPS) [dungeon]; Silk Headband (7050, -0.36 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 | yes | Darkspear Warding Pendant (272075, -0.86 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 | yes | Death Speaker Mantle (6685, -0.09 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.34 DPS) [crafted]; Fairywing Mantle (9536, -0.34 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Prelacy Cape (7004, -0.11 DPS) [quest]; Caretaker's Cape (19533, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (57.8 DPS) | yes | Death Speaker Robes (6682, -0.28 DPS) [dungeon]; Pristine Gown (253961, -0.39 DPS) [crafted]; Tree Bark Jacket (1486, -0.87 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Glowing Magical Bracelets (13106, -0.80 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.87 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.87 DPS) [quest] |
| hands | Serpent Gloves (5970) (or Shilly Mitts (9609)) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Shilly Mitts (9609, +0.00 DPS, sim-verified) [quest]; Town Clerk's Mittens (270029, -0.05 DPS) [quest]; Gnoll Casting Gloves (892, -0.11 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.7 | yes | Belt of Arugal (6392, +0.00 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.41 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.43 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.38 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.53 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.7 | yes | Acidic Walkers (9454, -0.20 DPS) [dungeon]; Nimbus Boots (6998, -0.31 DPS) [quest]; Spidersilk Boots (4320, -1.09 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.18 DPS) [quest]; Lorekeeper's Ring (20431, -0.23 DPS) [rep]; Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.34 DPS) [dungeon]; Sludge-Stained Band (286535, -0.34 DPS) [world]; Minor Channeling Ring (1449, -0.66 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (56.9 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Insignia of the Horde (18850, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Ash Staff (791) | World drop [world_drop] | 222.8 | yes | Glimmering Staff (249392, +0.00 DPS, sim-verified) [crafted]; Lorekeeper's Staff (212580, -0.80 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.80 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 291.8 | yes | Starfaller (13063, -0.12 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Gnarled Ash Staff; ranged: Necrotic Wand

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 19549 Sentinel's Blade

### Band 40 (gnome, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 179.8. Weights run: 0.9s. Verify run: 0.8s. 286 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.091 ± 0.015, crit=3.159 ± 0.088, hit=3.864 ± 0.105, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -3.04 DPS, sim-verified) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop]; Silk Headband (7050, -4.11 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 | yes | Darkspear Warding Pendant (272074, -2.37 DPS) [vendor]; Necklace of Calisea (1714, -2.42 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.43 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 | yes | Green Silken Shoulders (7057, -0.27 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Berylline Pads (4197, -0.65 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Long Silken Cloak (4326, -0.13 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted]; Caretaker's Cape (19532, -0.34 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 | yes | Elemental Raiment (9434, -0.64 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.28 DPS) [crafted]; Robe of Power (7054, -2.56 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.69 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.26 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.4 | yes | Star Belt (4329, -0.56 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -1.06 DPS) [rep]; Belt of Arugal (6392, -1.75 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 | yes | Gaze Dreamer Pants (6903, -1.25 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.03 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -3.31 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Nimbus Boots (6998, -6.17 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 | yes | Ring of Forlorn Spirits (2043, -0.87 DPS) [quest]; Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Ring of Forlorn Spirits (2043, -0.35 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.69 DPS) [quest]; Lorekeeper's Ring (19525, -0.69 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 146.0 | yes | Staff of Jordan (873, -1.31 DPS, sim-verified) [world_drop]; Celestial Stave (9517, -14.14 DPS) [quest]; Black Duskwood Staff (937, -16.49 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 115.4 | yes | Nether Force Wand (11263, -0.17 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.28 DPS) [quest]; Ragefire Wand (7513, -2.33 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs

### Band 50 (gnome, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 254.1. Weights run: 0.9s. Verify run: 0.8s. 359 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.020, crit=4.698 ± 0.137, hit=5.742 ± 0.243, spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 81.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -4.64 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.94 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 | yes | Mindburst Medallion (11196, -0.39 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.15 DPS) [world_drop]; Darkspear Warding Pendant (272073, -2.34 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 74.7 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -20.95 DPS) [dungeon]; Black Mageweave Shoulders (10027, -22.34 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 | yes | Runecloth Cloak (13860, -1.84 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.98 DPS) [dungeon]; Icy Cloak (4327, -2.68 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 79.0 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -19.73 DPS) [world_drop]; Acumen Robes (17775, -20.25 DPS) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.44 DPS) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 70.9 | yes | Dreamweave Gloves (10019, -1.61 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -19.60 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -19.96 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 75.3 | yes | Satyrmane Sash (17755, -0.10 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -21.33 DPS) [rep]; Ghostweave Cord (254073, -21.48 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 79.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -21.05 DPS) [crafted]; Red Mageweave Pants (10009, -22.34 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 66.4 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -14.85 DPS) [crafted]; Gilded Sandals (254107, -19.07 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 | yes | Lorekeeper's Ring (19523, -15.91 DPS) [rep]; Philanthropist's Ring (281635, -16.39 DPS) [quest]; Lorekeeper's Ring (19524, -16.96 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Lorekeeper's Ring (19523, -0.39 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.82 DPS) [quest]; Lorekeeper's Ring (19524, -1.40 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (253.6 DPS) | yes | Smoking Heart of the Mountain (11811, -3.33 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -18.11 DPS) [quest]; Tidal Charm (1404, -18.11 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (253.6 DPS) | yes | Thunderbrew's Boot Flask (744, -2.10 DPS) [quest]; Tidal Charm (1404, -2.10 DPS) [vendor]; Smoking Heart of the Mountain (11811, -2.35 DPS, sim-verified) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (253.6 DPS) | yes | Hanzo Sword (8190, -2.85 DPS, sim-verified) [world_drop]; Inventor's Focal Sword (17719, -11.08 DPS) [dungeon]; Illusionary Rod (7713, -11.38 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 359, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (gnome, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 528.5. Weights run: 0.9s. Verify run: 0.9s. 722 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.137 ± 0.028, crit=6.853 ± 0.198, hit=8.396 ± 0.345, spell_haste=4.452 ± 0.152, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (520.5 DPS) | yes | Field Marshal's Coronet (16441, -9.56 DPS) [vendor]; Warlord's Silk Cowl (16533, -9.56 DPS) [vendor]; Bloodvine Goggles (19999, -15.67 DPS, sim-verified) [crafted] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (510.2 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -0.69 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 139.0 | yes | Rugged Mantle of the Timbermaw (227808, -5.91 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -8.50 DPS) [crafted]; Champion's Silk Mantle (23264, -9.30 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 101.1 | yes | Earthweave Cloak (21187, -5.99 DPS) [quest]; Howler's Furs (272414, -5.99 DPS) [vendor]; Chromatic Cloak (18509, -9.46 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 257.5 | yes | Bloodvine Vest (19682, -12.98 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -32.87 DPS) [vendor]; Robe of the Archmage (14152, -42.01 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 133.1 | yes | Fireleaf Wristwraps (240044, -3.19 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -7.77 DPS) [quest]; Dryad's Wrist Bindings (19595, -38.55 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 235.8 | yes | Gloves of Spell Mastery (14146, -14.15 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -31.81 DPS) [vendor]; Sorcerer's Gloves (22066, -48.32 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 141.9 | yes | Fireleaf Belt (240053, +0.00 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -8.31 DPS) [crafted]; Knowledge of the Timbermaw (228190, -10.47 DPS) [vendor] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (512.9 DPS) | yes | Sentinel's Silk Leggings (237815, -8.10 DPS, sim-verified) [vendor]; Marshal's Silk Leggings (16442, -10.46 DPS) [vendor]; General's Silk Trousers (16534, -10.46 DPS) [vendor] |
| feet | Fireleaf Boots (240050) | Leonid Barthalomew the Revered [vendor] | 140.9 | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (16437, -11.91 DPS) [vendor]; General's Silk Boots (16539, -11.91 DPS) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (510.0 DPS) | yes | Mindtear Band (20632, +0.00 DPS) [world]; Ritssyn's Ring of Chaos (21836, +0.00 DPS) [world_drop]; Don Julio's Band (19325, -11.10 DPS, sim-verified) [rep] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 179.9 | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Ritssyn's Ring of Chaos (21836, +0.00 DPS, sim-verified) [world_drop]; Mindtear Band (20632, -21.42 DPS) [world] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (496.7 DPS) | yes | Uther's Strength (11302, -2.80 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -4.91 DPS) [quest]; Frozen Heart of the Mountain (249469, -7.37 DPS, sim-verified) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (496.7 DPS) | yes | Frozen Heart of the Mountain (249469, -4.81 DPS, sim-verified) [crafted]; Uther's Strength (11302, -5.61 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -7.71 DPS) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (510.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; The Lobotomizer (19324, -8.87 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | 182.7 | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Stormrager (16997, -1.07 DPS) [quest]; Torch of Austen (13004, -7.21 DPS) [world_drop] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Wrath of Cenarius; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 722, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

## Horde

### Band 20 (orc, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 29.6. Weights run: 0.7s. Verify run: 0.7s. 122 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.395 ± 0.090, crit=1.480 ± 0.048, hit=3.723 ± 0.108, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.31 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.49 DPS) [crafted]; Lucky Fishing Hat (19972, -0.49 DPS) [quest] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 | yes | Reinforced Woolen Shoulders (4315, -0.33 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.37 DPS) [crafted]; Slime-encrusted Pads (6461, -0.70 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.19 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.23 DPS) [crafted]; Green Woolen Vest (2582, -0.24 DPS) [crafted]; Gray Woolen Robe (2585, -0.47 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.03 DPS) [quest]; Owlbeard Bracers (16981, -0.05 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.08 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.15 DPS) [crafted]; Apothecary Gloves (10919, -0.25 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.20 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.32 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.42 DPS) [dungeon]; Colorful Kilt (10048, -0.59 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Pristine Boots (253889, -0.36 DPS) [crafted]; Red Woolen Boots (4313, -0.38 DPS) [crafted]; Feather Padded Treads (285345, -0.39 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Loop of Sacrifice (281673, -0.25 DPS) [quest]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, -0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.08 DPS) [quest]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (29.6 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Crescent Staff (6505, -0.42 DPS) [quest]; Living Root (6631, -0.51 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 275.0 | yes | Skycaller (12984, -0.53 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.27 DPS) [dungeon]; Deepblaze (279896, -4.03 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 122, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 209623 Insignia of the Horde; 241089 Scarlet Dagger; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest

### Band 30 (orc, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 50.8. Weights run: 0.8s. Verify run: 0.7s. 210 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.236 ± 0.191), crit=1.585 ± 0.050, hit=3.984 ± 0.329, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.34 DPS) [dungeon]; Silk Headband (7050, -0.47 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 | yes | Darkspear Warding Pendant (272075, -0.78 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 | yes | Death Speaker Mantle (6685, -0.03 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.34 DPS) [crafted]; Fairywing Mantle (9536, -0.34 DPS) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Battle Healer's Cloak (19529, -0.11 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.39 DPS) [dungeon]; Pristine Gown (253961, -0.50 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Glowing Magical Bracelets (13106, -0.80 DPS, sim-verified) [world_drop]; Owlbeard Bracers (16981, -0.87 DPS) [quest]; Nightsky Wristbands (6407, -0.87 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.2 | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Serpent Gloves (5970, -0.16 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.17 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.7 | yes | Warsong Sash (16975, -0.16 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.23 DPS) [dungeon]; Invoker's Cord (215366, -0.41 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.38 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.53 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.7 | yes | Acidic Walkers (9454, -0.20 DPS) [dungeon]; Boots of the Enchanter (4325, -0.42 DPS) [crafted]; Spidersilk Boots (4320, -1.06 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.23 DPS) [rep]; Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.46 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.34 DPS) [world]; Sacred Band (6669, -0.46 DPS) [quest]; Electrocutioner Lagnut (9447, -0.63 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (50.4 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Insignia of the Horde (18850, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Ash Staff (791) | World drop [world_drop] | 222.8 | yes | Glimmering Staff (249392, +0.00 DPS, sim-verified) [crafted]; Lorekeeper's Staff (212580, -0.80 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.80 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 291.8 | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.98 DPS) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Gnarled Ash Staff; ranged: Necrotic Wand

No-known-source sample (15 of 210, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 19545 Scout's Blade; 19553 Legionnaire's Sword; 19569 Advisor's Gnarled Staff

### Band 40 (orc, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 172.9. Weights run: 0.9s. Verify run: 0.8s. 287 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.091 ± 0.015, crit=3.159 ± 0.088, hit=3.864 ± 0.105, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -2.93 DPS, sim-verified) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop]; Silk Headband (7050, -4.11 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 | yes | Necklace of Calisea (1714, -2.32 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -2.37 DPS) [vendor]; Darkspear Warding Pendant (272075, -2.43 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 | yes | Green Silken Shoulders (7057, -0.26 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Chestnut Mantle (17695, -0.62 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Long Silken Cloak (4326, -0.12 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted]; Battle Healer's Cloak (19528, -0.34 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 | yes | Elemental Raiment (9434, -0.59 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.28 DPS) [crafted]; Robe of Power (7054, -2.56 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Condor Bracers (15864, -0.67 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.46 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.14 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.4 | yes | Star Belt (4329, -0.47 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -1.06 DPS) [rep]; Warsong Sash (16975, -1.15 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 | yes | Gaze Dreamer Pants (6903, -1.22 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.03 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -3.17 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Acidic Walkers (9454, -6.26 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 | yes | Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor]; Electrocutioner Lagnut (9447, -2.59 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Reedknot Ring (9622, -0.67 DPS, sim-verified) [quest]; Advisor's Ring (19521, -0.69 DPS) [rep]; Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 146.0 | yes | Staff of Jordan (873, -1.23 DPS, sim-verified) [world_drop]; Celestial Stave (9517, -14.14 DPS) [quest]; Black Duskwood Staff (937, -16.49 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 115.4 | yes | Nether Force Wand (11263, -0.09 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.28 DPS) [quest]; Ragefire Wand (7513, -2.33 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 287, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (orc, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 243.9. Weights run: 0.9s. Verify run: 0.8s. 360 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.020, crit=4.698 ± 0.137, hit=5.742 ± 0.243, spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 81.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -4.64 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.94 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 | yes | Mindburst Medallion (11196, -0.38 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.15 DPS) [world_drop]; Darkspear Warding Pendant (272073, -2.34 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 74.7 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -20.95 DPS) [dungeon]; Black Mageweave Shoulders (10027, -22.34 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 | yes | Deep Woodlands Cloak (19121, -0.64 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.68 DPS) [crafted]; Nightfall Drape (12465, -1.98 DPS) [dungeon] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 79.0 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -19.73 DPS) [world_drop]; Acumen Robes (17775, -20.25 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest]; Bloodband Bracers (11469, -1.06 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 70.9 | yes | Dreamweave Gloves (10019, -1.21 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -19.60 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -19.96 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 75.3 | yes | Satyrmane Sash (17755, -0.11 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -21.33 DPS) [rep]; Ghostweave Cord (254073, -21.48 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 79.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -21.05 DPS) [crafted]; Red Mageweave Pants (10009, -22.34 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 66.4 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -14.85 DPS) [crafted]; Gilded Sandals (254107, -19.07 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 | yes | Advisor's Ring (19519, -15.91 DPS) [rep]; Philanthropist's Ring (281635, -16.39 DPS) [quest]; Advisor's Ring (19520, -16.96 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Advisor's Ring (19519, -0.38 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.82 DPS) [quest]; Advisor's Ring (19520, -1.40 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (244.9 DPS) | yes | Uther's Strength (11302, -0.42 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -18.11 DPS) [vendor]; Guardian Talisman (1490, -18.11 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (244.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -14.08 DPS) [vendor]; Guardian Talisman (1490, -14.08 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (244.9 DPS) | yes | Hanzo Sword (8190, -2.69 DPS, sim-verified) [world_drop]; Inventor's Focal Sword (17719, -11.08 DPS) [dungeon]; Illusionary Rod (7713, -11.38 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 360, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 508.2. Weights run: 0.9s. Verify run: 0.9s. 725 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.137 ± 0.028, crit=6.853 ± 0.198, hit=8.396 ± 0.345, spell_haste=4.452 ± 0.152, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (500.6 DPS) | yes | Field Marshal's Coronet (16441, -9.56 DPS) [vendor]; Warlord's Silk Cowl (16533, -9.56 DPS) [vendor]; Bloodvine Goggles (19999, -14.69 DPS, sim-verified) [crafted] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (491.4 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -0.76 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 139.0 | yes | Rugged Mantle of the Timbermaw (227808, -5.68 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -8.50 DPS) [crafted]; Champion's Silk Mantle (23264, -9.30 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 101.1 | yes | Earthweave Cloak (21187, -5.99 DPS) [quest]; Howler's Furs (272414, -5.99 DPS) [vendor]; Chromatic Cloak (18509, -9.59 DPS, sim-verified) [crafted] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 257.5 | yes | Bloodvine Vest (19682, -14.49 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -32.87 DPS) [vendor]; Robe of the Archmage (14152, -42.01 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 133.1 | yes | Fireleaf Wristwraps (240044, -3.14 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -7.77 DPS) [quest]; Dryad's Wrist Bindings (19595, -38.55 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 235.8 | yes | Gloves of Spell Mastery (14146, -13.69 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -31.81 DPS) [vendor]; Sorcerer's Gloves (22066, -48.32 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 141.9 | yes | Fireleaf Belt (240053, +0.00 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -8.31 DPS) [crafted]; Knowledge of the Timbermaw (228190, -10.47 DPS) [vendor] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (493.7 DPS) | yes | Sentinel's Silk Leggings (237815, -7.84 DPS, sim-verified) [vendor]; Marshal's Silk Leggings (16442, -10.46 DPS) [vendor]; General's Silk Trousers (16534, -10.46 DPS) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 140.9 | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (16437, -11.91 DPS) [vendor]; General's Silk Boots (16539, -11.91 DPS) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (491.1 DPS) | yes | Mindtear Band (20632, +0.00 DPS) [world]; Ritssyn's Ring of Chaos (21836, +0.00 DPS) [world_drop]; Don Julio's Band (19325, -11.52 DPS, sim-verified) [rep] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 179.9 | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Ritssyn's Ring of Chaos (21836, +0.00 DPS, sim-verified) [world_drop]; Mindtear Band (20632, -21.42 DPS) [world] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (469.4 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.80 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (475.7 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -4.92 DPS, sim-verified) [crafted]; Uther's Strength (11302, -5.61 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (491.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; The Lobotomizer (19324, -8.61 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | 182.7 | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Stormrager (16997, -1.07 DPS) [quest]; Torch of Austen (13004, -7.21 DPS) [world_drop] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Wrath of Cenarius; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 725, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

