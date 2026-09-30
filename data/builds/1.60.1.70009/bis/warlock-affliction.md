# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.6. Weights run: 1.0s. Verify run: 1.0s. 129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.222, intellect=0.726 ± 0.012, crit=0.290 ± 0.024, hit=0.887 ± 0.015, spell_haste=-1.704 ± 0.181, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -2.81 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 spell_power points (1.13 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.74 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (38.1 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.47 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.6 spell_power points (0.85 DPS) | yes | Green Woolen Robe (6243, -0.34 DPS) [crafted]; Mystic's Wrap (14369, -0.35 DPS) [world_drop]; Gray Woolen Robe (2585, -1.28 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.4 spell_power points (0.43 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.14 DPS) [world_drop]; Mystic's Bracelets (14366, -0.29 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Blight Gloves (279877, -0.19 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.9 spell_power points (0.68 DPS) | yes | Keller's Girdle (2911, -0.11 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.96 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (38.2 DPS) | yes | Silk-threaded Trousers (1929, -0.33 DPS) [dungeon]; Darkweave Breeches (12987, -0.52 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.57 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (0.97 DPS) | yes | Pristine Boots (253889, -0.36 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.48 DPS) [world]; Red Woolen Boots (4313, -0.58 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (0.63 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Loop of Sacrifice (281673, -0.28 DPS) [quest]; Sludge-Stained Band (286535, -0.34 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.49 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Sludge-Stained Band (286535, -0.20 DPS) [world]; Lavishly Jeweled Ring (1156, -0.99 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.3 spell_power points (0.71 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.14 DPS) [world]; Lesser Staff of the Spire (1300, -0.29 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 229.9 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.51 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Deepblaze (279896, -4.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 129, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.2. Weights run: 0.9s. Verify run: 1.0s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.064, intellect=0.165 ± 0.004, crit=0.129 ± 0.009, hit=0.366 ± 0.006, spell_haste=0.807 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.064

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (4.20 DPS) | yes | Silk Headband (7050, -0.75 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -1.15 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -1.15 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (3.05 DPS) | yes | Darkspear Warding Pendant (272075, -2.32 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.80 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.80 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (4.00 DPS) | yes | Death Speaker Mantle (6685, -1.02 DPS) [dungeon]; Fairywing Mantle (9536, -1.15 DPS) [quest]; Invoker's Mantle (215365, -1.50 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.91 DPS) | yes | Heavy Woolen Cloak (4311, -0.30 DPS, sim-verified) [crafted]; Prelacy Cape (7004, -0.38 DPS) [quest]; Caretaker's Cape (19533, -0.38 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (4.96 DPS) | yes | Green Silk Armor (7065, -0.56 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.60 DPS) [dungeon]; Pristine Gown (253961, -1.85 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Glowing Magical Bracelets (13106, -2.16 DPS, sim-verified) [world_drop]; Windsong Bangles (263336, -3.05 DPS) [quest]; Nightsky Wristbands (6407, -3.06 DPS) [world_drop] |
| hands | Serpent Gloves (5970) (or Shilly Mitts (9609)) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.67 DPS) | yes | Shilly Mitts (9609, -0.09 DPS, sim-verified) [quest]; Gnoll Casting Gloves (892, -0.38 DPS) [world]; Town Clerk's Mittens (270029, -0.45 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (4.39 DPS) | yes | Belt of Arugal (6392, -0.68 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -1.33 DPS) [dungeon]; Invoker's Cord (215366, -1.40 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (4.58 DPS) | yes | Abomination Skin Leggings (23173, -0.81 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.47 DPS) [crafted]; Silk-threaded Trousers (1929, -1.91 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.2 spell_power points (3.11 DPS) | yes | Acidic Walkers (9454, -0.70 DPS) [dungeon]; Nimbus Boots (6998, -0.82 DPS) [quest]; Spidersilk Boots (4320, -2.43 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (2.67 DPS) | yes | Minor Channeling Ring (1449, -0.64 DPS) [quest]; Lorekeeper's Ring (20431, -0.76 DPS) [rep]; Electrocutioner Lagnut (9447, -1.53 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (2.29 DPS) | yes | Electrocutioner Lagnut (9447, -1.15 DPS) [dungeon]; Sludge-Stained Band (286535, -1.15 DPS) [world]; Minor Channeling Ring (1449, -1.49 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Glimmering Staff (249392, -2.07 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -2.81 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -2.81 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (3.05 DPS) | yes | Eye of Paleth (2943, -1.52 DPS) [quest]; Orb of Souls (249395, -1.52 DPS) [crafted]; Dwarven Tome (279898, -1.66 DPS, sim-verified) [quest] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 90.0 spell_power points (34.36 DPS) | yes | Starfaller (13063, -1.14 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.45 DPS) [crafted]; Thunderwood (13062, -4.95 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 107.7. Weights run: 0.8s. Verify run: 0.9s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.122, intellect=0.301 ± 0.010, crit=0.287 ± 0.020, hit=0.869 ± 0.013, spell_haste=not significant (-0.202 ± 0.176), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.122

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.90 DPS) | yes | Living Cowl (5608, -2.25 DPS) [world]; Augural Shroud (2620, -2.63 DPS, sim-verified) [world]; Holy Shroud (2721, -2.81 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 spell_power points (2.47 DPS) | yes | Necklace of Calisea (1714, -1.88 DPS) [world_drop]; Triune Amulet (7722, -1.88 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.35 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (3.29 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.22 DPS) [dungeon]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (107.7 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.14 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.80 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (6.69 DPS) | yes | Elemental Raiment (9434, -0.57 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.87 DPS) [crafted]; Robe of Power (7054, -1.74 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.53 DPS) | yes | Spidertank Oilrag (9448, -0.47 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest]; Earthen Silk Cuffs (254019, -1.41 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (5.40 DPS) | yes | Black Mageweave Gloves (10003, -1.30 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.46 DPS) [crafted]; Gilded Handwraps (254021, -2.56 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.2 spell_power points (4.27 DPS) | yes | Highlander's Cloth Girdle (20099, -0.93 DPS) [rep]; Star Belt (4329, -0.94 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.04 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.6 spell_power points (4.95 DPS) | yes | Crimson Silk Pantaloons (7062, -1.60 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.72 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.74 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.74 DPS) | yes | Gilded Slippers (254001, -2.89 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.44 DPS) [crafted]; Acidic Walkers (9454, -4.66 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.8 spell_power points (3.32 DPS) | yes | Ring of Forlorn Spirits (2043, -1.07 DPS) [quest]; Reedknot Ring (9622, -1.35 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.63 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.53 DPS) | yes | Ring of Forlorn Spirits (2043, -0.48 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.56 DPS) [quest]; Lorekeeper's Ring (19525, -0.56 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Staff of Dar'Orahil (15106, -2.25 DPS) [quest]; Scorn's Focal Dagger (23168, -3.09 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 141.9 spell_power points (39.89 DPS) | yes | Twisted Nether Wand (249144, +0.00 DPS, sim-verified) [crafted]; Umbral Wand (5216, -4.22 DPS) [world_drop]; Earthen Rod (9381, -4.30 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Spellforce Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 184.9. Weights run: 0.9s. Verify run: 1.0s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.227, intellect=0.211 ± 0.015, crit=0.457 ± 0.026, hit=2.027 ± 0.030, spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.14 DPS) [vendor]; Red Mageweave Headband (10033, -2.12 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (171.2 DPS) | yes | Mindburst Medallion (11196, -0.28 DPS) [quest]; Horizon Choker (13085, -1.49 DPS) [world_drop]; Arcane Crystal Pendant (20037, -2.90 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.8 spell_power points (4.71 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -0.14 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.37 DPS) [crafted]; Kentic Amice (11624, -5.03 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.28 DPS) | yes | Runecloth Cloak (13860, -1.28 DPS) [crafted]; Nightfall Drape (12465, -1.76 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -6.08 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (6.52 DPS) | yes | Acumen Robes (17775, -0.32 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.64 DPS) [world_drop]; Knight's Dreadweave Vest (220886, -0.72 DPS) [vendor] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.15 DPS) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.28 DPS) | yes | Black Mageweave Gloves (10003, -0.99 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.11 DPS) [vendor]; Runecloth Gloves (13863, -1.39 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 16.4 spell_power points (4.61 DPS) | yes | Satyrmane Sash (17755, -0.09 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.45 DPS) [rep]; Ban'thok Sash (11662, -4.27 DPS, sim-verified) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | sim-verified (173.3 DPS) | yes | Wizardweave Leggings (14132, -0.54 DPS) [crafted]; Red Mageweave Pants (10009, -1.23 DPS) [crafted]; Spellshock Leggings (9484, -5.06 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (172.8 DPS) | yes | Gilded Sandals (254107, -3.11 DPS) [crafted]; Black Mageweave Boots (10026, -3.23 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.49 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.3 spell_power points (5.68 DPS) | yes | Lorekeeper's Ring (19523, -2.32 DPS) [rep]; Philanthropist's Ring (281635, -2.52 DPS) [quest]; Cyclopean Band (11824, -2.74 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Lorekeeper's Ring (19523, -0.13 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.49 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -1.01 DPS, sim-verified) [crafted] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 0.0 spell_power points (0.00 DPS) | yes | Staff of Dar'Orahil (15106, -0.30 DPS) [quest]; Spellforce Rod (1664, -1.02 DPS) [world_drop]; Shortsword of Vengeance (754, -2.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (173.5 DPS) | yes | Woestave (20082, -0.08 DPS) [quest]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Pyric Caduceus (11748, -5.17 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 396.0. Weights run: 0.9s. Verify run: 1.0s. 945 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.250, intellect=0.107 ± 0.021, crit=0.421 ± 0.033, hit=2.006 ± 0.034, spell_haste=-3.778 ± 0.375, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.250

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 109.4 spell_power points (38.59 DPS) | yes | Deathmist Mask (226909, -24.26 DPS) [quest]; Bloodvine Goggles (19999, -24.79 DPS, sim-verified) [crafted]; Deathmist Mask (22074, -24.96 DPS) [quest] |
| neck | Orb of the Darkmoon (19426) | 1200 Tickets - Orb of the Darkmoon [quest] | 0.0 spell_power points (0.00 DPS) | yes | Chains of the Lich (23125, +0.00 DPS) [dungeon]; Beads of Ogre Might (22150, -0.68 DPS) [quest]; Blazefury Medallion (17111, -8.69 DPS, sim-verified) [world] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 67.3 spell_power points (23.75 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -7.21 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -13.61 DPS) [vendor]; Field Marshal's Dreadweave Shoulders (17580, -14.29 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.9 spell_power points (13.02 DPS) | yes | Earthweave Cloak (21187, -5.95 DPS) [quest]; Howler's Furs (272414, -5.95 DPS) [vendor]; Crystalline Threaded Cape (20697, -7.83 DPS, sim-verified) [world_drop] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 78.6 spell_power points (27.72 DPS) | yes | Robe of the Void (14153, -11.16 DPS) [crafted]; Heretic Garb (240146, -11.41 DPS) [vendor]; Bloodvine Vest (19682, -13.42 DPS, sim-verified) [crafted] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 47.5 spell_power points (16.74 DPS) | yes | Rockfury Bracers (21186, -0.07 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -4.72 DPS) [vendor]; Black Bark Wristbands (20626, -7.77 DPS) [world] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 58.5 spell_power points (20.62 DPS) | yes | Deathmist Wraps (226911, -8.46 DPS) [quest]; Deathmist Wraps (22077, -9.35 DPS, sim-verified) [quest]; Marshal's Dreadweave Gloves (17584, -9.81 DPS) [vendor] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 73.2 spell_power points (25.83 DPS) | yes | Knowledge of the Timbermaw (228190, -9.10 DPS) [vendor]; Heretic Waistguard (240151, -9.19 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -16.09 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 88.7 spell_power points (31.27 DPS) | yes | Bloodvine Leggings (19683, -9.18 DPS, sim-verified) [crafted]; Heretic Pants (240149, -15.32 DPS) [vendor]; Sentinel's Silk Leggings (237815, -15.45 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 57.6 spell_power points (20.30 DPS) | yes | Bloodvine Boots (19684, -7.68 DPS, sim-verified) [crafted]; Snowblind Shoes (19131, -8.64 DPS) [world]; Sergeant Major's Dreadweave Boots (220891, -10.06 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.74 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -1.10 DPS) [vendor]; Wrath of Cenarius (21190, -3.84 DPS, sim-verified) [quest] |
| finger2 | Ritssyn's Ring of Chaos (21836) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Blessed Band of Light (272407, -0.29 DPS) [vendor]; Mindtear Band (20632, -0.83 DPS) [world]; Wrath of Cenarius (21190, -2.77 DPS, sim-verified) [quest] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+13.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, -3.05 DPS, sim-verified) [vendor] |
| main_hand | Amberseal Keeper (17113) | Lord Kazzak [world] | 0.0 spell_power points (0.00 DPS) | yes | Shortsword of Vengeance (754, +0.00 DPS, sim-verified) [world_drop]; Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (396.0 DPS) | yes | Cold Snap (19130, -6.99 DPS, sim-verified) [world]; Ritssyn's Wand of Bad Mojo (22408, -11.15 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.22 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Ritssyn's Ring of Chaos; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Amberseal Keeper; ranged: Torch of Light

No-known-source sample (15 of 945, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.8. Weights run: 1.0s. Verify run: 1.0s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.222, intellect=0.726 ± 0.012, crit=0.290 ± 0.024, hit=0.887 ± 0.015, spell_haste=-1.704 ± 0.181, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -2.74 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 spell_power points (1.13 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.33 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.74 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.41 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.12 DPS) [dungeon]; Black Whelp Cloak (7283, -0.12 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.6 spell_power points (0.85 DPS) | yes | Green Woolen Robe (6243, -0.34 DPS) [crafted]; Mystic's Wrap (14369, -0.35 DPS) [world_drop]; Gray Woolen Robe (2585, -1.41 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.4 spell_power points (0.43 DPS) | yes | Featherbead Bracers (15452, -0.07 DPS) [quest]; Mindthrust Bracers (1974, -0.09 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.14 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Blight Gloves (279877, -0.19 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.9 spell_power points (0.68 DPS) | yes | Keller's Girdle (2911, -0.11 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.12 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (36.8 DPS) | yes | Silk-threaded Trousers (1929, -0.33 DPS) [dungeon]; Darkweave Breeches (12987, -0.52 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.58 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (0.97 DPS) | yes | Pristine Boots (253889, -0.43 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.48 DPS) [world]; Red Woolen Boots (4313, -0.58 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.49 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Sludge-Stained Band (286535, -0.20 DPS) [world]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 4.4 spell_power points (0.43 DPS) | yes | Sludge-Stained Band (286535, -0.13 DPS) [world]; Loop of Sacrifice (281673, -0.16 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.3 spell_power points (0.71 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.14 DPS) [world]; Lesser Staff of the Spire (1300, -0.29 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 229.9 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.90 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Deepblaze (279896, -4.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 58.6. Weights run: 0.9s. Verify run: 1.0s. 218 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.064, intellect=0.165 ± 0.004, crit=0.129 ± 0.009, hit=0.366 ± 0.006, spell_haste=0.807 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.064

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (4.20 DPS) | yes | Silk Headband (7050, -0.76 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -1.15 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -1.15 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.0 spell_power points (3.05 DPS) | yes | Darkspear Warding Pendant (272075, -1.68 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.80 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.80 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (4.00 DPS) | yes | Invoker's Mantle (215365, -1.02 DPS) [crafted]; Death Speaker Mantle (6685, -1.02 DPS) [dungeon]; Chestnut Mantle (17695, -1.38 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.91 DPS) | yes | Windsong Drape (15468, -0.01 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.38 DPS) [crafted]; Battle Healer's Cloak (19529, -0.38 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 spell_power points (4.96 DPS) | yes | Green Silk Armor (7065, -0.69 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.60 DPS) [dungeon]; High Robe of the Adjudicator (3461, -1.78 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Owlbeard Bracers (16981, -1.65 DPS, sim-verified) [quest]; Glowing Magical Bracelets (13106, -2.93 DPS) [world_drop]; Windsong Bangles (263336, -3.05 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.67 DPS) | yes | Jutebraid Gloves (10654, -0.06 DPS, sim-verified) [quest]; Gnoll Casting Gloves (892, -0.38 DPS) [world]; Truefaith Gloves (7049, -0.57 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (4.39 DPS) | yes | Warsong Sash (16975, -0.04 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.76 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -1.33 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (4.58 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.47 DPS) [crafted]; Silk-threaded Trousers (1929, -1.91 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.2 spell_power points (3.11 DPS) | yes | Acidic Walkers (9454, -0.70 DPS) [dungeon]; Boots of the Enchanter (4325, -1.20 DPS) [crafted]; Spidersilk Boots (4320, -1.95 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (2.67 DPS) | yes | Advisor's Ring (20426, -0.76 DPS) [rep]; Electrocutioner Lagnut (9447, -1.53 DPS) [dungeon]; Sludge-Stained Band (286535, -1.53 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (2.29 DPS) | yes | Sludge-Stained Band (286535, -1.15 DPS) [world]; Sacred Band (6669, -1.53 DPS) [quest]; Electrocutioner Lagnut (9447, -1.93 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Glimmering Staff (249392, -1.62 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -2.81 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -2.81 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 8.0 spell_power points (3.05 DPS) | yes | Dwarven Tome (279898, -1.19 DPS, sim-verified) [quest]; Orb of Souls (249395, -1.52 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -1.52 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 90.0 spell_power points (34.36 DPS) | yes | Starfaller (13063, -1.06 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.45 DPS) [crafted]; Thunderwood (13062, -4.95 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 107.4. Weights run: 0.8s. Verify run: 0.9s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.122, intellect=0.301 ± 0.010, crit=0.287 ± 0.020, hit=0.869 ± 0.013, spell_haste=not significant (-0.202 ± 0.176), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.122

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.90 DPS) | yes | Augural Shroud (2620, -2.21 DPS, sim-verified) [world]; Living Cowl (5608, -2.25 DPS) [world]; Holy Shroud (2721, -2.81 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 spell_power points (2.47 DPS) | yes | Necklace of Calisea (1714, -1.88 DPS) [world_drop]; Triune Amulet (7722, -1.88 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.14 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 spell_power points (3.29 DPS) | yes | Inquisitor's Shawl (19507, -0.22 DPS) [dungeon]; Green Silken Shoulders (7057, -0.26 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.48 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (105.4 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.14 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -1.80 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (6.69 DPS) | yes | Elemental Raiment (9434, -0.29 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.87 DPS) [crafted]; Robe of Power (7054, -1.74 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.53 DPS) | yes | Condor Bracers (15864, -0.64 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -0.73 DPS) [quest]; Earthen Silk Cuffs (254019, -1.41 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.2 spell_power points (5.40 DPS) | yes | Black Mageweave Gloves (10003, -0.96 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.46 DPS) [crafted]; Gilded Handwraps (254021, -2.56 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.2 spell_power points (4.27 DPS) | yes | Star Belt (4329, -0.46 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -0.93 DPS) [rep]; Deathmage Sash (10771, -1.04 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.6 spell_power points (4.95 DPS) | yes | Gaze Dreamer Pants (6903, -0.69 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -1.60 DPS) [crafted]; Abomination Skin Leggings (23173, -1.74 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.74 DPS) | yes | Gilded Slippers (254001, -2.25 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -4.44 DPS) [crafted]; Acidic Walkers (9454, -4.66 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.8 spell_power points (3.32 DPS) | yes | Reedknot Ring (9622, -1.35 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.63 DPS) [vendor]; Electrocutioner Lagnut (9447, -2.47 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.53 DPS) | yes | Advisor's Ring (19521, -0.56 DPS) [rep]; Reedknot Ring (9622, -0.64 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Staff of Dar'Orahil (15106, -2.25 DPS) [quest]; Scorn's Focal Dagger (23168, -3.09 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (105.2 DPS) | yes | Umbral Wand (5216, -0.02 DPS) [world_drop]; Earthen Rod (9381, -0.10 DPS) [dungeon]; Jaina's Firestarter (13064, -1.56 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 184.7. Weights run: 0.9s. Verify run: 1.0s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.227, intellect=0.211 ± 0.015, crit=0.457 ± 0.026, hit=2.027 ± 0.030, spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Dreamweave Circlet (10041, -1.09 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.14 DPS) [vendor]; Red Mageweave Headband (10033, -2.32 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (168.8 DPS) | yes | Mindburst Medallion (11196, -0.28 DPS) [quest]; Horizon Choker (13085, -1.49 DPS) [world_drop]; Arcane Crystal Pendant (20037, -3.11 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.8 spell_power points (4.71 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -0.14 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.37 DPS) [crafted]; Kentic Amice (11624, -5.35 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.3 spell_power points (4.28 DPS) | yes | Mantle of Lady Falther'ess (23178, -1.22 DPS) [dungeon]; Runecloth Cloak (13860, -1.28 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.67 DPS, sim-verified) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (6.52 DPS) | yes | Acumen Robes (17775, -0.33 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.64 DPS) [world_drop]; Stone Guard's Dreadweave Vest (220904, -0.72 DPS) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (167.9 DPS) | yes | Condor Bracers (15864, -0.41 DPS) [quest]; Bloodband Bracers (11469, -0.44 DPS) [quest]; Spidertank Oilrag (9448, -2.16 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.28 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.11 DPS) [vendor]; Runecloth Gloves (13863, -1.39 DPS) [crafted]; Black Mageweave Gloves (10003, -1.87 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 16.4 spell_power points (4.61 DPS) | yes | Satyrmane Sash (17755, -0.09 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.45 DPS) [rep]; Ban'thok Sash (11662, -4.73 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | sim-verified (171.1 DPS) | yes | Wizardweave Leggings (14132, -0.54 DPS) [crafted]; Red Mageweave Pants (10009, -1.23 DPS) [crafted]; Spellshock Leggings (9484, -5.37 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (169.6 DPS) | yes | Gilded Sandals (254107, -3.11 DPS) [crafted]; Black Mageweave Boots (10026, -3.23 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.86 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.3 spell_power points (5.68 DPS) | yes | Advisor's Ring (19519, -2.32 DPS) [rep]; Philanthropist's Ring (281635, -2.52 DPS) [quest]; Cyclopean Band (11824, -2.74 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Philanthropist's Ring (281635, -0.49 DPS) [quest]; Advisor's Ring (19519, -0.61 DPS, sim-verified) [rep]; Cyclopean Band (11824, -0.71 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.94 DPS, sim-verified) [crafted] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 0.0 spell_power points (0.00 DPS) | yes | Staff of Dar'Orahil (15106, -0.30 DPS) [quest]; Spellforce Rod (1664, -1.02 DPS) [world_drop]; Shortsword of Vengeance (754, -2.53 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (170.3 DPS) | yes | Woestave (20082, -0.08 DPS) [quest]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Pyric Caduceus (11748, -4.56 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Nethergeld Cuffs; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 386.8. Weights run: 0.9s. Verify run: 1.0s. 939 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.250, intellect=0.107 ± 0.021, crit=0.421 ± 0.033, hit=2.006 ± 0.034, spell_haste=-3.778 ± 0.375, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.250

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 109.4 spell_power points (38.59 DPS) | yes | Bloodvine Goggles (19999, -23.47 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -24.26 DPS) [quest]; Deathmist Mask (22074, -24.96 DPS) [quest] |
| neck | Orb of the Darkmoon (19426) | 1200 Tickets - Orb of the Darkmoon [quest] | 0.0 spell_power points (0.00 DPS) | yes | Chains of the Lich (23125, +0.00 DPS) [dungeon]; Beads of Ogre Might (22150, -0.68 DPS) [quest]; Blazefury Medallion (17111, -8.84 DPS, sim-verified) [world] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 67.3 spell_power points (23.75 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -6.70 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -13.61 DPS) [vendor]; Warlord's Dreadweave Mantle (17590, -14.29 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.9 spell_power points (13.02 DPS) | yes | Earthweave Cloak (21187, -5.95 DPS) [quest]; Howler's Furs (272414, -5.95 DPS) [vendor]; Crystalline Threaded Cape (20697, -5.96 DPS, sim-verified) [world_drop] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 78.6 spell_power points (27.72 DPS) | yes | Robe of the Void (14153, -11.16 DPS) [crafted]; Heretic Garb (240146, -11.41 DPS) [vendor]; Bloodvine Vest (19682, -12.30 DPS, sim-verified) [crafted] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 47.5 spell_power points (16.74 DPS) | yes | Rockfury Bracers (21186, +0.00 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -4.72 DPS) [vendor]; Black Bark Wristbands (20626, -7.77 DPS) [world] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 58.5 spell_power points (20.62 DPS) | yes | Deathmist Wraps (226911, -8.46 DPS) [quest]; General's Dreadweave Gloves (17588, -9.81 DPS) [vendor]; Deathmist Wraps (22077, -10.07 DPS, sim-verified) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 73.2 spell_power points (25.83 DPS) | yes | Heretic Waistguard (240151, -7.96 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -9.10 DPS) [vendor]; Belt of the Archmage (18405, -16.09 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 88.7 spell_power points (31.27 DPS) | yes | Bloodvine Leggings (19683, -7.09 DPS, sim-verified) [crafted]; Heretic Pants (240149, -15.32 DPS) [vendor]; Sentinel's Silk Leggings (237815, -15.45 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 57.6 spell_power points (20.30 DPS) | yes | Bloodvine Boots (19684, -6.13 DPS, sim-verified) [crafted]; Snowblind Shoes (19131, -8.64 DPS) [world]; First Sergeant's Dreadweave Boots (220909, -10.06 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.74 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -1.10 DPS) [vendor]; Wrath of Cenarius (21190, -4.37 DPS, sim-verified) [quest] |
| finger2 | Ritssyn's Ring of Chaos (21836) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Blessed Band of Light (272407, -0.29 DPS) [vendor]; Mindtear Band (20632, -0.83 DPS) [world]; Wrath of Cenarius (21190, -3.40 DPS, sim-verified) [quest] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+15.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -7.80 DPS, sim-verified) [crafted] |
| main_hand | Amberseal Keeper (17113) | Lord Kazzak [world] | 0.0 spell_power points (0.00 DPS) | yes | Shortsword of Vengeance (754, +0.00 DPS, sim-verified) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Fang of the Mystics (17070, -1.64 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (386.8 DPS) | yes | Cold Snap (19130, -6.61 DPS, sim-verified) [world]; Ritssyn's Wand of Bad Mojo (22408, -11.15 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.22 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Ritssyn's Ring of Chaos; trinket1: Talisman of Ascendance; trinket2: Weakness Analyzer; main_hand: Amberseal Keeper; ranged: Torch of Light

No-known-source sample (15 of 939, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

