# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 43.2. Weights run: 1.1s. Verify run: 1.1s. 129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.152, intellect=0.841 ± 0.010, crit=0.348 ± 0.028, hit=0.848 ± 0.013, spell_haste=0.663 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.665 ± 0.152, fire_power=0.340 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -3.28 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.6 spell_power points (1.66 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.22 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.13 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (41.1 DPS) | yes | Sanguine Cape (14376, -0.08 DPS) [world_drop]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.44 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.2 spell_power points (1.21 DPS) | yes | Mystic's Wrap (14369, -0.44 DPS) [world_drop]; Mystic's Robe (14371, -0.44 DPS) [world_drop]; Gray Woolen Robe (2585, -1.55 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 5.0 spell_power points (0.67 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.22 DPS) [world_drop]; Mystic's Bracelets (14366, -0.44 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.13 DPS) [world]; Blight Gloves (279877, -0.15 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.4 spell_power points (0.97 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Novice Ardent's Sash (253887, -0.37 DPS) [crafted]; Keller's Girdle (2911, -2.07 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (41.5 DPS) | yes | Silk-threaded Trousers (1929, -0.53 DPS) [dungeon]; Darkweave Breeches (12987, -0.68 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.93 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.4 spell_power points (1.37 DPS) | yes | Pristine Boots (253889, -0.29 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.71 DPS) [world]; Red Woolen Boots (4313, -0.84 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.7 spell_power points (0.88 DPS) | yes | Loop of Sacrifice (281673, -0.33 DPS) [quest]; Sludge-Stained Band (286535, -0.49 DPS) [world]; Lavishly Jeweled Ring (1156, -1.45 DPS, sim-verified) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (41.8 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -1.22 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 8.4 spell_power points (1.11 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.22 DPS) [world]; Lesser Staff of the Spire (1300, -0.44 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 172.2 spell_power points (22.71 DPS) | yes | Skycaller (12984, -2.18 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 129, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 66.1. Weights run: 1.1s. Verify run: 1.0s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.118, intellect=0.426 ± 0.008, crit=0.549 ± 0.036, hit=0.668 ± 0.011, spell_haste=not significant (0.090 ± 0.086), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.832 ± 0.118, fire_power=0.168 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.85 DPS) | yes | Silk Headband (7050, -0.52 DPS) [crafted]; Embalmed Shroud (7691, -0.78 DPS) [dungeon]; Enchanter's Cowl (4322, -1.05 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (2.47 DPS) | yes | Darkspear Warding Pendant (272075, -1.81 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.03 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.03 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.8 spell_power points (3.32 DPS) | yes | Death Speaker Mantle (6685, -0.76 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.78 DPS) [quest]; Invoker's Mantle (215365, -0.96 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 spell_power points (1.29 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.26 DPS) [crafted]; Prelacy Cape (7004, -0.26 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.5 spell_power points (3.76 DPS) | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.74 DPS) [dungeon]; Pristine Gown (253961, -1.18 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.33 DPS) | yes | Nightsky Wristbands (6407, -1.67 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.67 DPS) [quest]; Glowing Magical Bracelets (13106, -2.21 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.7 spell_power points (2.25 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Shilly Mitts (9609, -0.44 DPS) [quest]; Truefaith Gloves (7049, -0.62 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.3 spell_power points (3.18 DPS) | yes | Belt of Arugal (6392, -0.38 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.82 DPS) [crafted]; Crimson Silk Belt (7055, -0.85 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.4 spell_power points (3.21 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.63 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.00 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 spell_power points (2.58 DPS) | yes | Acidic Walkers (9454, -0.41 DPS) [dungeon]; Nimbus Boots (6998, -1.03 DPS) [quest]; Spidersilk Boots (4320, -1.66 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.81 DPS) | yes | Minor Channeling Ring (1449, -0.30 DPS) [quest]; Lorekeeper's Ring (20431, -0.52 DPS) [rep]; Electrocutioner Lagnut (9447, -1.04 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.55 DPS) | yes | Electrocutioner Lagnut (9447, -0.78 DPS) [dungeon]; Sludge-Stained Band (286535, -0.78 DPS) [world]; Minor Channeling Ring (1449, -1.82 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (66.1 DPS) | yes | Talisman of Arathor (21119, +0.00 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.33 DPS) | yes | Twisted Chanter's Staff (890, -1.23 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.23 DPS) [quest]; Glimmering Staff (249392, -1.72 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.6 spell_power points (2.47 DPS) | yes | Dwarven Tome (279898, -0.39 DPS, sim-verified) [quest]; Eye of Paleth (2943, -1.44 DPS) [quest]; Orb of Souls (249395, -1.44 DPS) [crafted] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 131.3 spell_power points (33.99 DPS) | yes | Starfaller (13063, -0.60 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.69 DPS) [crafted]; Gravestone Scepter (7001, -4.99 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 110.3. Weights run: 1.0s. Verify run: 1.0s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.396, intellect=0.842 ± 0.030, crit=1.986 ± 0.117, hit=2.861 ± 0.043, spell_haste=3.882 ± 0.364, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.107 ± 0.396), fire_power=1.100 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.49 DPS) | yes | Corpseshroud (10574, -0.59 DPS) [dungeon]; Enchanter's Cowl (4322, -0.78 DPS) [crafted]; Augural Shroud (2620, -1.73 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.43 DPS) | yes | Necklace of Calisea (1714, -0.73 DPS) [world_drop]; Triune Amulet (7722, -0.73 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.03 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.0 spell_power points (2.13 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (110.3 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.11 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.65 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.1 spell_power points (3.20 DPS) | yes | Robe of Power (7054, -0.35 DPS) [crafted]; Dreamweave Vest (10021, -0.48 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.72 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.07 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.17 DPS) [world_drop]; Condor Bracers (15864, -0.24 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.4 spell_power points (2.53 DPS) | yes | Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Red Mageweave Gloves (10018, -0.86 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.86 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.6 spell_power points (2.33 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.58 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.72 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.1 spell_power points (2.86 DPS) | yes | Abomination Skin Leggings (23173, -0.99 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.23 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.36 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.84 DPS) | yes | Acidic Walkers (9454, -1.45 DPS) [dungeon]; Spidersilk Boots (4320, -1.61 DPS) [crafted]; Gilded Slippers (254001, -1.91 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.78 DPS) | yes | Ring of Forlorn Spirits (2043, -0.84 DPS) [quest]; Reedknot Ring (9622, -0.95 DPS) [quest]; Minor Channeling Ring (1449, -0.99 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.07 DPS) | yes | Ring of Forlorn Spirits (2043, -0.10 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.24 DPS) [quest]; Lorekeeper's Ring (19525, -0.24 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.59 DPS) [dungeon]; Staff of Noh'Orahil (15105, -0.79 DPS) [quest]; Gut Ripper (2164, -1.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 337.6 spell_power points (39.98 DPS) | yes | Umbral Wand (5216, -0.00 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.39 DPS) [dungeon]; Twisted Nether Wand (249144, -5.27 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 138.0. Weights run: 1.0s. Verify run: 1.0s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.605, intellect=1.132 ± 0.038, crit=2.617 ± 0.140, hit=4.179 ± 0.053, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 64.2 spell_power points (9.95 DPS) | yes | Eye of Theradras (17715, -3.27 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -3.50 DPS) [crafted]; Dreamweave Circlet (10041, -4.94 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 22.8 spell_power points (3.53 DPS) | yes | Horizon Choker (13085, +0.00 DPS, sim-verified) [world_drop]; Scorn's Icy Choker (23169, -1.39 DPS) [dungeon]; Mindburst Medallion (11196, -1.55 DPS) [quest] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) | Captain Dirgehammer [vendor] | 54.8 spell_power points (8.49 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Kentic Amice (11624, -4.05 DPS) [dungeon]; Red Mageweave Shoulders (10029, -4.78 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.8 spell_power points (3.22 DPS) | yes | Runecloth Cloak (13860, -0.42 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.77 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -4.18 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (134.2 DPS) | yes | Knight's Dreadweave Vest (220886, -1.75 DPS, sim-verified) [vendor]; Runecloth Robe (13858, -1.77 DPS) [crafted]; Runecloth Tunic (13857, -1.89 DPS) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (134.2 DPS) | yes | Nethergeld Cuffs (254061, -0.04 DPS) [crafted]; Forgotten Wraps (9433, -0.25 DPS) [world_drop]; Aristocratic Cuffs (12546, -1.75 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 31.9 spell_power points (4.94 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.16 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.45 DPS) [crafted]; Red Mageweave Gloves (10018, -1.48 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 51.3 spell_power points (7.95 DPS) | yes | Ban'thok Sash (11662, -3.47 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -3.69 DPS) [dungeon]; Satyrmane Sash (17755, -4.03 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 62.2 spell_power points (9.64 DPS) | yes | Spellshock Leggings (9484, -4.23 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -5.37 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -6.02 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (135.4 DPS) | yes | Gilded Sandals (254107, -0.44 DPS) [crafted]; Southsea Mojo Boots (20641, -0.55 DPS) [quest]; Sergeant Major's Dreadweave Boots (220891, -3.00 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.8 spell_power points (6.48 DPS) | yes | Cyclopean Band (11824, -3.85 DPS) [dungeon]; Philanthropist's Ring (281635, -3.87 DPS) [quest]; Mindseye Circle (10634, -4.37 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Philanthropist's Ring (281635, -0.03 DPS) [quest]; Mindseye Circle (10634, -0.53 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -1.91 DPS, sim-verified) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -0.40 DPS, sim-verified) [crafted] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 0.0 spell_power points (0.00 DPS) | yes | Staff of Dar'Orahil (15106, -0.88 DPS) [quest]; Kindling Stave (11750, -1.50 DPS) [dungeon]; Shortsword of Vengeance (754, -1.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 338.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.73 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -6.47 DPS, sim-verified) [quest] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Arcane Crystal Pendant; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Raider Handwraps; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Brainlash; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Soul Harvester; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 335.5. Weights run: 1.0s. Verify run: 0.9s. 945 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.338, intellect=0.654 ± 0.019, crit=1.580 ± 0.087, hit=2.312 ± 0.031, spell_haste=not significant (0.018 ± 0.286), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.566 ± 0.338), fire_power=0.427 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 138.9 spell_power points (51.50 DPS) | yes | Bloodvine Goggles (19999, -21.62 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -30.43 DPS) [quest]; Lieutenant Commander's Dreadweave Cowl (23310, -31.15 DPS) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points (0.00 DPS) | yes | Beads of Ogre Might (22150, -0.14 DPS) [quest]; Medallion of the Dawn (22659, -0.52 DPS) [quest]; Blazefury Medallion (17111, -4.35 DPS, sim-verified) [world] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 93.7 spell_power points (34.77 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -8.29 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -16.47 DPS) [vendor]; Mantle of the Timbermaw (19050, -17.11 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 44.4 spell_power points (16.45 DPS) | yes | Earthweave Cloak (21187, -4.33 DPS, sim-verified) [quest]; Howler's Furs (272414, -7.88 DPS) [vendor]; Stalwart Cloak (272415, -7.88 DPS) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 120.4 spell_power points (44.64 DPS) | yes | Bloodvine Vest (19682, -12.07 DPS, sim-verified) [crafted]; Heretic Garb (240146, -17.90 DPS) [vendor]; Earthpower Vest (21183, -21.82 DPS) [quest] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 57.6 spell_power points (21.37 DPS) | yes | Rockfury Bracers (21186, -0.75 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -7.61 DPS) [vendor]; Black Bark Wristbands (20626, -11.13 DPS) [world] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 68.6 spell_power points (25.45 DPS) | yes | Deathmist Wraps (22077, -8.90 DPS) [quest]; Deathmist Wraps (226911, -8.90 DPS) [quest]; Gloves of Spell Mastery (14146, -8.94 DPS, sim-verified) [crafted] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 115.9 spell_power points (42.97 DPS) | yes | Heretic Waistguard (240151, -11.18 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -19.78 DPS) [vendor]; Belt of the Archmage (18405, -23.47 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 133.5 spell_power points (49.50 DPS) | yes | Sentinel's Silk Leggings (237815, -10.99 DPS, sim-verified) [vendor]; Heretic Pants (240149, -23.14 DPS) [vendor]; Bloodvine Leggings (19683, -25.75 DPS) [crafted] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 68.3 spell_power points (25.33 DPS) | yes | Bloodvine Boots (19684, -6.57 DPS, sim-verified) [crafted]; Heretic Boots (240153, -8.51 DPS) [vendor]; Snowblind Shoes (19131, -11.03 DPS) [world] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.98 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -1.36 DPS) [vendor]; Wrath of Cenarius (21190, -6.39 DPS, sim-verified) [quest] |
| finger2 | Mindtear Band (20632) | Taerar [world] | 0.0 spell_power points (0.00 DPS) | yes | Ritssyn's Ring of Chaos (21836, -0.34 DPS) [world_drop]; Don Julio's Band (19325, -1.04 DPS) [rep]; Wrath of Cenarius (21190, -4.17 DPS, sim-verified) [quest] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+12.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Weakness Analyzer (272438, -1.79 DPS, sim-verified) [vendor] |
| main_hand | Amberseal Keeper (17113) | Lord Kazzak [world] | 0.0 spell_power points (0.00 DPS) | yes | Shortsword of Vengeance (754, +0.00 DPS, sim-verified) [world_drop]; Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (335.5 DPS) | yes | Cold Snap (19130, -7.01 DPS, sim-verified) [world]; Ritssyn's Wand of Bad Mojo (22408, -11.06 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.30 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Amulet of the Dawn; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Mindtear Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Amberseal Keeper; ranged: Torch of Light

No-known-source sample (15 of 945, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 41.2. Weights run: 1.1s. Verify run: 1.1s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.152, intellect=0.841 ± 0.010, crit=0.348 ± 0.028, hit=0.848 ± 0.013, spell_haste=0.663 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.665 ± 0.152, fire_power=0.340 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -3.11 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.6 spell_power points (1.66 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.13 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (40.4 DPS) | yes | Sanguine Cape (14376, -0.08 DPS) [world_drop]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.47 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.2 spell_power points (1.21 DPS) | yes | Mystic's Wrap (14369, -0.44 DPS) [world_drop]; Mystic's Robe (14371, -0.44 DPS) [world_drop]; Gray Woolen Robe (2585, -1.48 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 5.0 spell_power points (0.67 DPS) | yes | Mindthrust Bracers (1974, -0.05 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.11 DPS) [quest]; Bright Bracers (3647, -0.22 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.13 DPS) [world]; Blight Gloves (279877, -0.15 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.4 spell_power points (0.97 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Novice Ardent's Sash (253887, -0.37 DPS) [crafted]; Keller's Girdle (2911, -2.18 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (40.6 DPS) | yes | Silk-threaded Trousers (1929, -0.53 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.66 DPS, sim-verified) [dungeon]; Darkweave Breeches (12987, -0.68 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.4 spell_power points (1.37 DPS) | yes | Pristine Boots (253889, -0.18 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.71 DPS) [world]; Red Woolen Boots (4313, -0.84 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 5.0 spell_power points (0.67 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Sludge-Stained Band (286535, -0.27 DPS) [world]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.66 DPS) | yes | Sludge-Stained Band (286535, -0.26 DPS) [world]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; Loop of Sacrifice (281673, -1.36 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 8.4 spell_power points (1.11 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.22 DPS) [world]; Lesser Staff of the Spire (1300, -0.44 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 172.2 spell_power points (22.71 DPS) | yes | Skycaller (12984, -2.07 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 65.0. Weights run: 1.1s. Verify run: 1.0s. 218 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.118, intellect=0.426 ± 0.008, crit=0.549 ± 0.036, hit=0.668 ± 0.011, spell_haste=not significant (0.090 ± 0.086), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.832 ± 0.118, fire_power=0.168 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (2.85 DPS) | yes | Silk Headband (7050, -0.52 DPS) [crafted]; Enchanter's Cowl (4322, -0.53 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.78 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (2.47 DPS) | yes | Darkspear Warding Pendant (272075, -1.95 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.03 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.03 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.8 spell_power points (3.32 DPS) | yes | Death Speaker Mantle (6685, -0.19 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.78 DPS) [quest]; Invoker's Mantle (215365, -0.96 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.29 DPS) | yes | Windsong Drape (15468, -0.01 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.26 DPS) [crafted]; Battle Healer's Cloak (19529, -0.26 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.5 spell_power points (3.76 DPS) | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.74 DPS) [dungeon]; Pristine Gown (253961, -1.18 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.33 DPS) | yes | Nightsky Wristbands (6407, -1.67 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.67 DPS) [quest]; Glowing Magical Bracelets (13106, -1.91 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.1 spell_power points (2.10 DPS) | yes | Serpent Gloves (5970, -0.20 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.48 DPS) [crafted]; Gnoll Casting Gloves (892, -0.55 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.3 spell_power points (3.18 DPS) | yes | Warsong Sash (16975, -0.51 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.52 DPS) [dungeon]; Invoker's Cord (215366, -0.82 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.4 spell_power points (3.21 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.63 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.00 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 spell_power points (2.58 DPS) | yes | Acidic Walkers (9454, -0.41 DPS) [dungeon]; Boots of the Enchanter (4325, -1.29 DPS) [crafted]; Spidersilk Boots (4320, -2.26 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.81 DPS) | yes | Advisor's Ring (20426, -0.52 DPS) [rep]; Electrocutioner Lagnut (9447, -1.04 DPS) [dungeon]; Sludge-Stained Band (286535, -1.04 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.55 DPS) | yes | Sludge-Stained Band (286535, -0.78 DPS) [world]; Black Widow Band (6199, -0.78 DPS) [world]; Electrocutioner Lagnut (9447, -2.24 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (65.0 DPS) | yes | Defiler's Talisman (21120, -0.78 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 9.0 spell_power points (2.33 DPS) | yes | Twisted Chanter's Staff (890, -1.23 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.23 DPS) [quest]; Glimmering Staff (249392, -1.47 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.6 spell_power points (2.47 DPS) | yes | Dwarven Tome (279898, -0.44 DPS, sim-verified) [quest]; Orb of Souls (249395, -1.44 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -1.44 DPS) [world] |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 131.3 spell_power points (33.99 DPS) | yes | Starfaller (13063, -1.01 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.69 DPS) [crafted]; Gravestone Scepter (7001, -4.99 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight; ranged: Necrotic Wand

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 107.6. Weights run: 1.0s. Verify run: 1.0s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.396, intellect=0.842 ± 0.030, crit=1.986 ± 0.117, hit=2.861 ± 0.043, spell_haste=3.882 ± 0.364, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.107 ± 0.396), fire_power=1.100 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.49 DPS) | yes | Corpseshroud (10574, -0.59 DPS) [dungeon]; Enchanter's Cowl (4322, -0.78 DPS) [crafted]; Augural Shroud (2620, -1.57 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.43 DPS) | yes | Necklace of Calisea (1714, -0.73 DPS) [world_drop]; Triune Amulet (7722, -0.73 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.20 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.0 spell_power points (2.13 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (107.6 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.11 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.42 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.1 spell_power points (3.20 DPS) | yes | Robe of Power (7054, -0.35 DPS) [crafted]; Dreamweave Vest (10021, -0.39 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.72 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.7 spell_power points (1.27 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.37 DPS) [world_drop]; Condor Bracers (15864, -0.44 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.4 spell_power points (2.53 DPS) | yes | Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Stormcloth Gloves (10011, -0.86 DPS) [crafted]; Red Mageweave Gloves (10018, -1.80 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.6 spell_power points (2.33 DPS) | yes | Defiler's Cloth Girdle (20166, +0.00 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.58 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.72 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.1 spell_power points (2.86 DPS) | yes | Abomination Skin Leggings (23173, -0.99 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.23 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.73 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.84 DPS) | yes | Acidic Walkers (9454, -1.45 DPS) [dungeon]; Spidersilk Boots (4320, -1.61 DPS) [crafted]; Gilded Slippers (254001, -1.82 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.78 DPS) | yes | Reedknot Ring (9622, -0.95 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.07 DPS) [vendor]; Ogremind Ring (1993, -1.08 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.07 DPS) | yes | Advisor's Ring (19521, -0.24 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.36 DPS) [vendor]; Reedknot Ring (9622, -0.59 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.59 DPS) [dungeon]; Staff of Noh'Orahil (15105, -0.79 DPS) [quest]; Gut Ripper (2164, -1.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 337.6 spell_power points (39.98 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.39 DPS) [dungeon]; Twisted Nether Wand (249144, -5.27 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 132.5. Weights run: 1.0s. Verify run: 0.9s. 391 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.605, intellect=1.132 ± 0.038, crit=2.617 ± 0.140, hit=4.179 ± 0.053, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 64.2 spell_power points (9.95 DPS) | yes | Red Mageweave Headband (10033, -3.50 DPS) [crafted]; Eye of Theradras (17715, -3.80 DPS, sim-verified) [dungeon]; Dreamweave Circlet (10041, -4.94 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 22.8 spell_power points (3.53 DPS) | yes | Horizon Choker (13085, +0.00 DPS, sim-verified) [world_drop]; Scorn's Icy Choker (23169, -1.39 DPS) [dungeon]; Mindburst Medallion (11196, -1.55 DPS) [quest] |
| shoulder | Blood Guard's Dreadweave Mantle (220905) | Lady Palanseer [vendor] | 54.8 spell_power points (8.49 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Kentic Amice (11624, -4.05 DPS) [dungeon]; Red Mageweave Shoulders (10029, -4.78 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.2 spell_power points (3.44 DPS) | yes | Spritecaster Cape (11623, -0.18 DPS, sim-verified) [dungeon]; Mantle of Lady Falther'ess (23178, -0.46 DPS) [dungeon]; Runecloth Cloak (13860, -0.64 DPS) [crafted] |
| chest | Stone Guard's Dreadweave Vest (220904) | Lady Palanseer [vendor] | 61.1 spell_power points (9.47 DPS) | yes | Acumen Robes (17775, +0.00 DPS, sim-verified) [quest]; Runecloth Robe (13858, -4.78 DPS) [crafted]; Runecloth Tunic (13857, -4.90 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Bloodband Bracers (11469, +0.00 DPS, sim-verified) [quest]; Nethergeld Cuffs (254061, -0.32 DPS) [crafted]; Forgotten Wraps (9433, -0.53 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 31.9 spell_power points (4.94 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.14 DPS, sim-verified) [vendor]; Dreamweave Gloves (10019, -1.45 DPS) [crafted]; Red Mageweave Gloves (10018, -1.48 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 51.3 spell_power points (7.95 DPS) | yes | Dawnspire Cord (12466, -3.69 DPS) [dungeon]; Ban'thok Sash (11662, -3.76 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -4.03 DPS) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 62.2 spell_power points (9.64 DPS) | yes | Spellshock Leggings (9484, -4.71 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -5.37 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -6.02 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (132.5 DPS) | yes | Gilded Sandals (254107, -0.44 DPS) [crafted]; Southsea Mojo Boots (20641, -0.55 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -1.70 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.8 spell_power points (6.48 DPS) | yes | Cyclopean Band (11824, -3.85 DPS) [dungeon]; Philanthropist's Ring (281635, -3.87 DPS) [quest]; Mindseye Circle (10634, -4.37 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.0 spell_power points (2.63 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Philanthropist's Ring (281635, -0.03 DPS) [quest]; Mindseye Circle (10634, -0.53 DPS) [dungeon] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.84 DPS, sim-verified) [crafted] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 0.0 spell_power points (0.00 DPS) | yes | Staff of Dar'Orahil (15106, -0.88 DPS) [quest]; Kindling Stave (11750, -1.50 DPS) [dungeon]; Shortsword of Vengeance (754, -1.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 338.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.73 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -7.54 DPS, sim-verified) [quest] |

**New at 50:** head: Blood Guard's Dreadweave Hat; neck: Arcane Crystal Pendant; shoulder: Blood Guard's Dreadweave Mantle; back: Deep Woodlands Cloak; chest: Stone Guard's Dreadweave Vest; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Brainlash; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Soul Harvester; ranged: Pyric Caduceus

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 330.1. Weights run: 1.0s. Verify run: 0.9s. 939 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.338, intellect=0.654 ± 0.019, crit=1.580 ± 0.087, hit=2.312 ± 0.031, spell_haste=not significant (0.018 ± 0.286), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.566 ± 0.338), fire_power=0.427 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 138.9 spell_power points (51.50 DPS) | yes | Bloodvine Goggles (19999, -21.61 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -30.43 DPS) [quest]; Champion's Dreadweave Cowl (227090, -31.15 DPS) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points (0.00 DPS) | yes | Beads of Ogre Might (22150, -0.14 DPS) [quest]; Medallion of the Dawn (22659, -0.52 DPS) [quest]; Blazefury Medallion (17111, -4.94 DPS, sim-verified) [world] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 93.7 spell_power points (34.77 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -6.66 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -16.47 DPS) [vendor]; Mantle of the Timbermaw (19050, -17.11 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 44.4 spell_power points (16.45 DPS) | yes | Earthweave Cloak (21187, -4.73 DPS, sim-verified) [quest]; Howler's Furs (272414, -7.88 DPS) [vendor]; Stalwart Cloak (272415, -7.88 DPS) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 120.4 spell_power points (44.64 DPS) | yes | Bloodvine Vest (19682, -12.61 DPS, sim-verified) [crafted]; Heretic Garb (240146, -17.90 DPS) [vendor]; Earthpower Vest (21183, -21.82 DPS) [quest] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 57.6 spell_power points (21.37 DPS) | yes | Rockfury Bracers (21186, -0.26 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -7.61 DPS) [vendor]; Black Bark Wristbands (20626, -11.13 DPS) [world] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 68.6 spell_power points (25.45 DPS) | yes | Gloves of Spell Mastery (14146, -8.45 DPS, sim-verified) [crafted]; Deathmist Wraps (22077, -8.90 DPS) [quest]; Deathmist Wraps (226911, -8.90 DPS) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 115.9 spell_power points (42.97 DPS) | yes | Heretic Waistguard (240151, -8.14 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -19.78 DPS) [vendor]; Belt of the Archmage (18405, -23.47 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 133.5 spell_power points (49.50 DPS) | yes | Sentinel's Silk Leggings (237815, -11.36 DPS, sim-verified) [vendor]; Heretic Pants (240149, -23.14 DPS) [vendor]; Bloodvine Leggings (19683, -25.75 DPS) [crafted] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 68.3 spell_power points (25.33 DPS) | yes | Bloodvine Boots (19684, -5.96 DPS, sim-verified) [crafted]; Heretic Boots (240153, -8.51 DPS) [vendor]; Snowblind Shoes (19131, -11.03 DPS) [world] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.98 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -1.36 DPS) [vendor]; Wrath of Cenarius (21190, -6.08 DPS, sim-verified) [quest] |
| finger2 | Mindtear Band (20632) | Taerar [world] | 0.0 spell_power points (0.00 DPS) | yes | Ritssyn's Ring of Chaos (21836, -0.34 DPS) [world_drop]; Don Julio's Band (19325, -1.04 DPS) [rep]; Wrath of Cenarius (21190, -1.73 DPS, sim-verified) [quest] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+12.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18467, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.26 DPS, sim-verified) [crafted] |
| main_hand | Amberseal Keeper (17113) | Lord Kazzak [world] | 0.0 spell_power points (0.00 DPS) | yes | Shortsword of Vengeance (754, +0.00 DPS, sim-verified) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Ironbark Staff (20220, -2.34 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (330.1 DPS) | yes | Cold Snap (19130, -6.60 DPS, sim-verified) [world]; Ritssyn's Wand of Bad Mojo (22408, -11.06 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.30 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Amulet of the Dawn; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Mindtear Band; trinket1: Talisman of Ascendance; trinket2: Weakness Analyzer; main_hand: Amberseal Keeper; ranged: Torch of Light

No-known-source sample (15 of 939, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

