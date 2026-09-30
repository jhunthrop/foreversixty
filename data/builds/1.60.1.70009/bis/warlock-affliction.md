# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 37.3. Weights run: 1.0s. Verify run: 1.0s. 129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.222, intellect=1.135 ± 0.222, crit=0.290 ± 0.024, hit=0.887 ± 0.015, spell_haste=-1.704 ± 0.181, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -3.15 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.2 spell_power points (1.50 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.28 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.10 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.4 spell_power points (0.53 DPS) | yes | Heavy Woolen Cloak (4311, -0.14 DPS) [crafted]; Feyscale Cloak (6632, -0.24 DPS) [dungeon]; Sanguine Cape (14376, -0.53 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 10.7 spell_power points (1.05 DPS) | yes | Mystic's Wrap (14369, -0.27 DPS) [world_drop]; Mystic's Robe (14371, -0.27 DPS) [world_drop]; Gray Woolen Robe (2585, -1.46 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 6.8 spell_power points (0.67 DPS) | yes | Mindthrust Bracers (1974, -0.05 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.22 DPS) [world_drop]; Mystic's Bracelets (14366, -0.45 DPS) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (35.0 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Tomb Robber's Gloves (280096, -0.06 DPS) [quest]; Blight Gloves (279877, -1.60 DPS, sim-verified) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (35.0 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Tarantula Silk Sash (3229, -0.28 DPS) [world]; Keller's Girdle (2911, -1.60 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (33.8 DPS) | yes | Abomination Skin Leggings (23173, -0.47 DPS, sim-verified) [dungeon]; Darkweave Breeches (12987, -0.48 DPS) [world_drop]; Silk-threaded Trousers (1929, -0.57 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.5 spell_power points (1.13 DPS) | yes | Pristine Boots (253889, -0.32 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.64 DPS) [world]; Walking Boots (4660, -0.69 DPS) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.3 spell_power points (0.71 DPS) | yes | Loop of Sacrifice (281673, -0.16 DPS) [quest]; Lorekeeper's Ring (20431, -0.22 DPS) [rep]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 6.8 spell_power points (0.67 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS, sim-verified) [quest]; Lorekeeper's Ring (20431, -0.18 DPS) [rep]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 11.4 spell_power points (1.12 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.22 DPS) [world]; Lesser Staff of the Spire (1300, -0.45 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 229.9 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.70 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Deepblaze (279896, -4.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 129, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 59.3. Weights run: 0.9s. Verify run: 1.0s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.064, intellect=0.690 ± 0.070, crit=0.129 ± 0.009, hit=0.366 ± 0.006, spell_haste=0.807 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.064

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (56.3 DPS) | yes | Silk Headband (7050, -0.76 DPS) [crafted]; Nightsky Cowl (4039, -1.04 DPS) [world_drop]; Enchanter's Cowl (4322, -1.10 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.1 spell_power points (4.25 DPS) | yes | Darkspear Warding Pendant (272075, -1.90 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -3.20 DPS) [world_drop]; Kaleidoscope Chain (13084, -3.20 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.2 spell_power points (5.81 DPS) | yes | Death Speaker Mantle (6685, -0.88 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -1.15 DPS) [quest]; Magician's Mantle (12998, -1.53 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.8 spell_power points (2.20 DPS) | yes | Darkspear Raider's Cloak (272078, -0.09 DPS) [vendor]; Hillman's Cloak (3719, -0.29 DPS) [crafted]; Cloak of Rot (4462, -0.78 DPS, sim-verified) [world] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.0 spell_power points (6.86 DPS) | yes | Death Speaker Robes (6682, -0.67 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -1.90 DPS) [dungeon]; Pristine Gown (253961, -2.34 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Nightsky Wristbands (6407, -1.86 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.86 DPS) [quest]; Glowing Magical Bracelets (13106, -2.14 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | sim-verified (56.6 DPS) | yes | Serpent Gloves (5970, -0.03 DPS) [dungeon]; Shilly Mitts (9609, -0.03 DPS) [quest]; Town Clerk's Mittens (270029, -1.36 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.1 spell_power points (4.99 DPS) | yes | Belt of Arugal (6392, -0.55 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.86 DPS) [crafted]; Invoker's Cord (215366, -1.00 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | sim-verified (55.9 DPS) | yes | Pristine Leggings (253987, -0.06 DPS) [crafted]; Abomination Skin Leggings (23173, -0.70 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.71 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.8 spell_power points (4.52 DPS) | yes | Spidersilk Boots (4320, -0.79 DPS) [crafted]; Nimbus Boots (6998, -2.23 DPS) [quest]; Acidic Walkers (9454, -2.32 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (2.67 DPS) | yes | Lorekeeper's Ring (20431, -0.76 DPS) [rep]; Black Widow Band (6199, -0.83 DPS) [world]; Minor Channeling Ring (1449, -1.90 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (57.2 DPS) | yes | Black Widow Band (6199, -0.45 DPS) [world]; Snake Hoop (6750, -0.45 DPS) [quest]; Minor Channeling Ring (1449, -1.97 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (55.2 DPS) | yes | Talisman of Arathor (21119, -0.32 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 7.6 spell_power points (2.90 DPS) | yes | Twisted Chanter's Staff (890, -0.01 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.26 DPS) [quest]; Channeler's Staff (4437, -0.79 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 90.0 spell_power points (34.36 DPS) | yes | Starfaller (13063, -0.80 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.45 DPS) [crafted]; Thunderwood (13062, -4.95 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 110.0. Weights run: 0.8s. Verify run: 0.8s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.122, intellect=not significant (0.613 ± 0.165), crit=0.287 ± 0.020, hit=0.869 ± 0.013, spell_haste=not significant (-0.202 ± 0.176), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.122

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.90 DPS) | yes | Enchanter's Cowl (4322, -2.49 DPS) [crafted]; Augural Shroud (2620, -2.62 DPS, sim-verified) [world]; Corpseshroud (10574, -2.63 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.7 spell_power points (3.00 DPS) | yes | Necklace of Calisea (1714, -1.79 DPS) [world_drop]; Triune Amulet (7722, -1.79 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.46 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.0 spell_power points (4.21 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.52 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.55 DPS) | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.58 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.65 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (7.22 DPS) | yes | Robe of Power (7054, -1.21 DPS) [crafted]; Dreamweave Vest (10021, -1.26 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.31 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.53 DPS) | yes | Spidertank Oilrag (9448, -0.06 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest]; Windchaser Cuffs (14429, -0.98 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (5.75 DPS) | yes | Black Mageweave Gloves (10003, -1.53 DPS) [crafted]; Red Mageweave Gloves (10018, -1.92 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.29 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.5 spell_power points (4.62 DPS) | yes | Star Belt (4329, -0.97 DPS) [crafted]; Gilded Cord (254037, -1.00 DPS) [crafted]; Deathmage Sash (10771, -1.75 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.4 spell_power points (6.00 DPS) | yes | Crimson Silk Pantaloons (7062, -2.05 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.09 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.63 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.74 DPS) | yes | Gilded Slippers (254001, -2.73 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.96 DPS) [dungeon]; Spidersilk Boots (4320, -4.09 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (3.84 DPS) | yes | Ring of Forlorn Spirits (2043, -1.60 DPS) [quest]; Reedknot Ring (9622, -1.88 DPS) [quest]; Minor Channeling Ring (1449, -2.09 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.53 DPS) | yes | Reedknot Ring (9622, -0.56 DPS) [quest]; Lorekeeper's Ring (19525, -0.56 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.74 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (108.4 DPS) | yes | Gut Ripper (2164, -1.03 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -1.75 DPS) [dungeon]; Illusionary Rod (7713, -2.17 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (110.0 DPS) | yes | Umbral Wand (5216, -0.02 DPS) [world_drop]; Earthen Rod (9381, -0.10 DPS) [dungeon]; Jaina's Firestarter (13064, -1.59 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Twisted Nether Wand

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 187.3. Weights run: 0.9s. Verify run: 1.0s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.227, intellect=not significant (0.199 ± 0.264), crit=0.457 ± 0.026, hit=2.027 ± 0.030, spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Red Mageweave Headband (10033, -1.13 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.18 DPS) [vendor]; Dreamweave Circlet (10041, -2.09 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.2 spell_power points (2.30 DPS) | yes | Mindburst Medallion (11196, -0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.52 DPS) [world_drop]; Gemshard Heart (17707, -1.74 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.6 spell_power points (4.65 DPS) | yes | Black Mageweave Shoulders (10027, -1.34 DPS) [crafted]; Bloodmage Mantle (7684, -1.62 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.71 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 spell_power points (4.26 DPS) | yes | Runecloth Cloak (13860, -1.32 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.74 DPS) [dungeon]; Icy Cloak (4327, -2.30 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 spell_power points (6.50 DPS) | yes | Acumen Robes (17775, -0.45 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.61 DPS) [world_drop]; Knight's Dreadweave Vest (220886, -0.73 DPS) [vendor] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.52 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.27 DPS) | yes | Sergeant Major's Dreadweave Gloves (220890, -1.12 DPS) [vendor]; Runecloth Gloves (13863, -1.40 DPS) [crafted]; Black Mageweave Gloves (10003, -1.49 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 16.4 spell_power points (4.59 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.45 DPS) [rep]; Ghostweave Cord (254073, -0.67 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | sim-verified (178.1 DPS) | yes | Red Mageweave Pants (10009, -0.73 DPS) [crafted]; Senior Designer's Pantaloons (11841, -1.57 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -2.31 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (180.4 DPS) | yes | Gilded Sandals (254107, -3.14 DPS) [crafted]; Black Mageweave Boots (10026, -3.25 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.65 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.3 spell_power points (5.68 DPS) | yes | Lorekeeper's Ring (19523, -2.32 DPS) [rep]; Philanthropist's Ring (281635, -2.54 DPS) [quest]; Lorekeeper's Ring (19524, -3.16 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Lorekeeper's Ring (19523, +0.00 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.51 DPS) [quest]; Lorekeeper's Ring (19524, -1.12 DPS) [rep] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (175.8 DPS) | yes | Frozen Heart of the Mountain (249469, -1.10 DPS, sim-verified) [crafted] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (175.8 DPS) | yes | Frozen Heart of the Mountain (249469, -3.10 DPS, sim-verified) [crafted] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | sim-verified (175.8 DPS) | yes | Staff of Dar'Orahil (15106, -0.28 DPS) [quest]; Shortsword of Vengeance (754, -1.42 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -4.11 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (180.5 DPS) | yes | Woestave (20082, -0.08 DPS) [quest]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Pyric Caduceus (11748, -4.74 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Highlander's Cloth Girdle; legs: Wizardweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Abyss Shard; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 393.9. Weights run: 0.9s. Verify run: 0.9s. 944 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.250, intellect=not significant (-0.863 ± 0.331), crit=0.421 ± 0.033, hit=2.006 ± 0.034, spell_haste=-3.778 ± 0.375, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.250

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 108.0 spell_power points (38.10 DPS) | yes | Bloodvine Goggles (19999, -23.55 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -24.67 DPS) [quest]; Deathmist Mask (22074, -25.38 DPS) [quest] |
| neck | Orb of the Darkmoon (19426) | 1200 Tickets - Orb of the Darkmoon [quest] | sim-verified (388.3 DPS) | yes | Chains of the Lich (23125, +0.00 DPS) [dungeon]; Beads of Ogre Might (22150, -0.68 DPS) [quest]; Blazefury Medallion (17111, -8.80 DPS, sim-verified) [world] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 66.0 spell_power points (23.26 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -6.95 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -13.42 DPS) [vendor]; Field Marshal's Dreadweave Shoulders (17580, -14.44 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.1 spell_power points (12.72 DPS) | yes | Earthweave Cloak (21187, -4.32 DPS, sim-verified) [quest]; Howler's Furs (272414, -5.64 DPS) [vendor]; Stalwart Cloak (272415, -5.64 DPS) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 76.8 spell_power points (27.08 DPS) | yes | Robe of the Void (14153, -10.86 DPS) [crafted]; Heretic Garb (240146, -11.23 DPS) [vendor]; Bloodvine Vest (19682, -11.94 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 47.1 spell_power points (16.60 DPS) | yes | Heretic Bindings (240145, -1.27 DPS, sim-verified) [vendor]; Heretic Wristguards (240152, -4.59 DPS) [vendor]; Dryad's Wrist Bindings (19595, -8.84 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 57.1 spell_power points (20.13 DPS) | yes | Deathmist Wraps (226911, -8.46 DPS) [quest]; Marshal's Dreadweave Gloves (17584, -9.54 DPS) [vendor]; Deathmist Wraps (22077, -9.97 DPS, sim-verified) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 71.8 spell_power points (25.34 DPS) | yes | Heretic Waistguard (240151, -7.71 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -9.44 DPS) [vendor]; Belt of the Archmage (18405, -16.21 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 86.8 spell_power points (30.63 DPS) | yes | Bloodvine Leggings (19683, -7.22 DPS, sim-verified) [crafted]; Heretic Pants (240149, -15.13 DPS) [vendor]; Sentinel's Silk Leggings (237815, -15.19 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 56.1 spell_power points (19.77 DPS) | yes | Bloodvine Boots (19684, -6.42 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -9.88 DPS) [vendor]; Marshal's Dreadweave Boots (17583, -10.60 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (388.3 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.71 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -1.06 DPS) [vendor]; Wrath of Cenarius (21190, -2.27 DPS, sim-verified) [quest] |
| finger2 | Ritssyn's Ring of Chaos (21836) | World drop [world_drop] | sim-verified (388.3 DPS) | yes | Blessed Band of Light (272407, -0.29 DPS) [vendor]; Mindtear Band (20632, -1.06 DPS) [world]; Wrath of Cenarius (21190, -1.50 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (388.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -2.00 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.82 DPS) [world_drop] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (388.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS, sim-verified) [vendor]; Uther's Strength (11302, -2.12 DPS) [world_drop] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (388.3 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -9.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (393.9 DPS) | yes | Cold Snap (19130, -5.61 DPS, sim-verified) [world]; Wand of Biting Cold (19108, -14.88 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, -15.03 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Rockfury Bracers; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Ritssyn's Ring of Chaos; trinket1: Serenity Field; main_hand: Crackling Staff; ranged: Torch of Light

No-known-source sample (15 of 944, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.9. Weights run: 1.0s. Verify run: 1.1s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.222, intellect=1.135 ± 0.222, crit=0.290 ± 0.024, hit=0.887 ± 0.015, spell_haste=-1.704 ± 0.181, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -2.69 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.2 spell_power points (1.50 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.12 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.10 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.4 spell_power points (0.53 DPS) | yes | Heavy Woolen Cloak (4311, -0.14 DPS) [crafted]; Feyscale Cloak (6632, -0.24 DPS) [dungeon]; Sanguine Cape (14376, -0.35 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 10.7 spell_power points (1.05 DPS) | yes | Mystic's Wrap (14369, -0.27 DPS) [world_drop]; Mystic's Robe (14371, -0.27 DPS) [world_drop]; Gray Woolen Robe (2585, -1.34 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 6.8 spell_power points (0.67 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.11 DPS) [quest]; Bright Bracers (3647, -0.22 DPS) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (33.1 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Tomb Robber's Gloves (280096, -0.06 DPS) [quest]; Blight Gloves (279877, -1.79 DPS, sim-verified) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (33.1 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Tarantula Silk Sash (3229, -0.28 DPS) [world]; Keller's Girdle (2911, -1.79 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (32.0 DPS) | yes | Darkweave Breeches (12987, -0.48 DPS) [world_drop]; Silk-threaded Trousers (1929, -0.57 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.67 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 11.5 spell_power points (1.13 DPS) | yes | Pristine Boots (253889, -0.17 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.64 DPS) [world]; Walking Boots (4660, -0.69 DPS) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 6.8 spell_power points (0.67 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; Sludge-Stained Band (286535, -0.37 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (32.5 DPS) | yes | Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; Sludge-Stained Band (286535, -0.20 DPS) [world]; Loop of Sacrifice (281673, -1.20 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 11.4 spell_power points (1.12 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.22 DPS) [world]; Lesser Staff of the Spire (1300, -0.45 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 229.9 spell_power points (22.60 DPS) | yes | Skycaller (12984, -1.40 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.31 DPS) [dungeon]; Deepblaze (279896, -4.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 56.6. Weights run: 0.9s. Verify run: 1.0s. 218 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.064, intellect=0.690 ± 0.070, crit=0.129 ± 0.009, hit=0.366 ± 0.006, spell_haste=0.807 ± 0.067, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.064

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.9 spell_power points (4.93 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Silk Headband (7050, -1.49 DPS) [crafted]; Nightsky Cowl (4039, -1.76 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.1 spell_power points (4.25 DPS) | yes | Darkspear Warding Pendant (272075, -1.87 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -3.20 DPS) [world_drop]; Kaleidoscope Chain (13084, -3.20 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.2 spell_power points (5.81 DPS) | yes | Death Speaker Mantle (6685, -0.84 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -1.15 DPS) [quest]; Magician's Mantle (12998, -1.53 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 5.5 spell_power points (2.11 DPS) | yes | Darkspear Raider's Cloak (272078, -0.08 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.20 DPS) [crafted]; Windsong Drape (15468, -0.20 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.0 spell_power points (6.86 DPS) | yes | Death Speaker Robes (6682, -0.51 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -1.90 DPS) [dungeon]; Pristine Gown (253961, -2.34 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.44 DPS) | yes | Nightsky Wristbands (6407, -1.86 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.86 DPS) [quest]; Glowing Magical Bracelets (13106, -2.03 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | sim-verified (56.6 DPS) | yes | Serpent Gloves (5970, -0.03 DPS) [dungeon]; Pristine Gloves (253913, -0.38 DPS) [crafted]; Jutebraid Gloves (10654, -0.99 DPS, sim-verified) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.1 spell_power points (4.99 DPS) | yes | Belt of Arugal (6392, -0.68 DPS, sim-verified) [dungeon]; Warsong Sash (16975, -0.79 DPS) [quest]; Crimson Silk Belt (7055, -0.86 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.5 spell_power points (5.54 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.03 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.67 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.8 spell_power points (4.52 DPS) | yes | Spidersilk Boots (4320, -0.79 DPS) [crafted]; Acidic Walkers (9454, -2.12 DPS, sim-verified) [dungeon]; Pristine Boots (253889, -2.58 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (2.67 DPS) | yes | Advisor's Ring (20426, -0.76 DPS) [rep]; Black Widow Band (6199, -0.83 DPS) [world]; Snake Hoop (6750, -0.83 DPS) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (2.29 DPS) | yes | Snake Hoop (6750, -0.45 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.71 DPS) [dungeon]; Black Widow Band (6199, -2.76 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 7.6 spell_power points (2.90 DPS) | yes | Twisted Chanter's Staff (890, -0.18 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.26 DPS) [quest]; Channeler's Staff (4437, -0.79 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 90.0 spell_power points (34.36 DPS) | yes | Starfaller (13063, -1.01 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.45 DPS) [crafted]; Thunderwood (13062, -4.95 DPS) [world_drop] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Defiler's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Relentless Raider's Seal; trinket2: Rune of Perfection; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 107.8. Weights run: 0.8s. Verify run: 0.9s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.122, intellect=not significant (0.613 ± 0.165), crit=0.287 ± 0.020, hit=0.869 ± 0.013, spell_haste=not significant (-0.202 ± 0.176), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.122

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.90 DPS) | yes | Augural Shroud (2620, -2.28 DPS, sim-verified) [world]; Enchanter's Cowl (4322, -2.49 DPS) [crafted]; Corpseshroud (10574, -2.63 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.7 spell_power points (3.00 DPS) | yes | Necklace of Calisea (1714, -1.79 DPS) [world_drop]; Triune Amulet (7722, -1.79 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.14 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.0 spell_power points (4.21 DPS) | yes | Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Green Silken Shoulders (7057, -0.13 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.52 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.1 spell_power points (2.55 DPS) | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.58 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.65 DPS) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.7 spell_power points (7.22 DPS) | yes | Dreamweave Vest (10021, -0.77 DPS, sim-verified) [crafted]; Robe of Power (7054, -1.21 DPS) [crafted]; Elemental Raiment (9434, -1.31 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.53 DPS) | yes | Condor Bracers (15864, -0.56 DPS) [quest]; Windchaser Cuffs (14429, -0.98 DPS) [world_drop]; Radiant Silver Bracers (4545, -1.23 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (5.75 DPS) | yes | Black Mageweave Gloves (10003, -1.53 DPS) [crafted]; Red Mageweave Gloves (10018, -1.58 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -2.29 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.5 spell_power points (4.62 DPS) | yes | Star Belt (4329, -0.97 DPS) [crafted]; Gilded Cord (254037, -1.00 DPS) [crafted]; Deathmage Sash (10771, -1.30 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 21.4 spell_power points (6.00 DPS) | yes | Crimson Silk Pantaloons (7062, -1.49 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.09 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.63 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.74 DPS) | yes | Gilded Slippers (254001, -2.46 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.96 DPS) [dungeon]; Spidersilk Boots (4320, -4.09 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (3.84 DPS) | yes | Reedknot Ring (9622, -1.88 DPS) [quest]; Sea Giant's Toe Ring (274746, -2.16 DPS) [vendor]; Ogremind Ring (1993, -2.64 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (2.53 DPS) | yes | Advisor's Ring (19521, -0.56 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.84 DPS) [vendor]; Reedknot Ring (9622, -0.87 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (106.3 DPS) | yes | Gut Ripper (2164, -1.12 DPS, sim-verified) [world_drop]; Windweaver Staff (7757, -1.75 DPS) [dungeon]; Illusionary Rod (7713, -2.17 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (107.8 DPS) | yes | Umbral Wand (5216, -0.02 DPS) [world_drop]; Earthen Rod (9381, -0.10 DPS) [dungeon]; Jaina's Firestarter (13064, -1.56 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Twisted Nether Wand

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 185.1. Weights run: 0.9s. Verify run: 1.0s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.227, intellect=not significant (0.199 ± 0.264), crit=0.457 ± 0.026, hit=2.027 ± 0.030, spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (7.57 DPS) | yes | Red Mageweave Headband (10033, -1.13 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.18 DPS) [vendor]; Dreamweave Circlet (10041, -2.37 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.2 spell_power points (2.30 DPS) | yes | Mindburst Medallion (11196, -0.62 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.52 DPS) [world_drop]; Gemshard Heart (17707, -1.74 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.6 spell_power points (4.65 DPS) | yes | Black Mageweave Shoulders (10027, -1.34 DPS) [crafted]; Bloodmage Mantle (7684, -1.62 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.35 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 spell_power points (4.26 DPS) | yes | Deep Woodlands Cloak (19121, -0.84 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.29 DPS) [crafted]; Nightfall Drape (12465, -1.74 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 spell_power points (6.50 DPS) | yes | Elemental Raiment (9434, -0.61 DPS) [world_drop]; Stone Guard's Dreadweave Vest (220904, -0.73 DPS) [vendor]; Acumen Robes (17775, -1.40 DPS, sim-verified) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest]; Bloodband Bracers (11469, -0.62 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.27 DPS) | yes | Black Mageweave Gloves (10003, -0.92 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.12 DPS) [vendor]; Runecloth Gloves (13863, -1.40 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | sim-verified (174.5 DPS) | yes | Defiler's Cloth Girdle (20166, -0.33 DPS) [rep]; Ghostweave Cord (254073, -0.56 DPS) [crafted]; Defiler's Cloth Girdle (20165, -2.09 DPS, sim-verified) [rep] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | sim-verified (174.6 DPS) | yes | Red Mageweave Pants (10009, -0.73 DPS) [crafted]; Senior Designer's Pantaloons (11841, -1.57 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -2.18 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (176.7 DPS) | yes | Gilded Sandals (254107, -3.14 DPS) [crafted]; Black Mageweave Boots (10026, -3.25 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.25 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.3 spell_power points (5.68 DPS) | yes | Advisor's Ring (19519, -2.32 DPS) [rep]; Philanthropist's Ring (281635, -2.54 DPS) [quest]; Advisor's Ring (19520, -3.16 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.64 DPS) | yes | Advisor's Ring (19519, -0.18 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.51 DPS) [quest]; Advisor's Ring (19520, -1.12 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (172.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.72 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (172.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | sim-verified (172.4 DPS) | yes | Staff of Dar'Orahil (15106, -0.28 DPS) [quest]; Shortsword of Vengeance (754, -2.04 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -4.11 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (177.4 DPS) | yes | Woestave (20082, -0.08 DPS) [quest]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Pyric Caduceus (11748, -4.99 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Abyss Shard; trinket2: Rune of the Guard Captain; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 384.4. Weights run: 0.9s. Verify run: 0.9s. 938 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.250, intellect=not significant (-0.863 ± 0.331), crit=0.421 ± 0.033, hit=2.006 ± 0.034, spell_haste=-3.778 ± 0.375, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.250

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 108.0 spell_power points (38.10 DPS) | yes | Deathmist Mask (226909, -24.67 DPS) [quest]; Bloodvine Goggles (19999, -25.29 DPS, sim-verified) [crafted]; Deathmist Mask (22074, -25.38 DPS) [quest] |
| neck | Orb of the Darkmoon (19426) | 1200 Tickets - Orb of the Darkmoon [quest] | sim-verified (378.8 DPS) | yes | Chains of the Lich (23125, +0.00 DPS) [dungeon]; Beads of Ogre Might (22150, -0.68 DPS) [quest]; Blazefury Medallion (17111, -8.64 DPS, sim-verified) [world] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 66.0 spell_power points (23.26 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -7.60 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -13.42 DPS) [vendor]; Warlord's Dreadweave Mantle (17590, -14.44 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 36.1 spell_power points (12.72 DPS) | yes | Howler's Furs (272414, -5.64 DPS) [vendor]; Stalwart Cloak (272415, -5.64 DPS) [vendor]; Earthweave Cloak (21187, -6.35 DPS, sim-verified) [quest] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 76.8 spell_power points (27.08 DPS) | yes | Robe of the Void (14153, -10.86 DPS) [crafted]; Heretic Garb (240146, -11.23 DPS) [vendor]; Bloodvine Vest (19682, -15.36 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 47.1 spell_power points (16.60 DPS) | yes | Heretic Bindings (240145, +0.00 DPS, sim-verified) [vendor]; Heretic Wristguards (240152, -4.59 DPS) [vendor]; Dryad's Wrist Bindings (19595, -8.84 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 57.1 spell_power points (20.13 DPS) | yes | Deathmist Wraps (226911, -8.46 DPS) [quest]; General's Dreadweave Gloves (17588, -9.54 DPS) [vendor]; Deathmist Wraps (22077, -9.80 DPS, sim-verified) [quest] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 71.8 spell_power points (25.34 DPS) | yes | Knowledge of the Timbermaw (228190, -9.44 DPS) [vendor]; Heretic Waistguard (240151, -9.71 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -16.21 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 86.8 spell_power points (30.63 DPS) | yes | Bloodvine Leggings (19683, -6.92 DPS, sim-verified) [crafted]; Heretic Pants (240149, -15.13 DPS) [vendor]; Sentinel's Silk Leggings (237815, -15.19 DPS) [vendor] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 56.1 spell_power points (19.77 DPS) | yes | Bloodvine Boots (19684, -6.79 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Boots (220909, -9.88 DPS) [vendor]; General's Dreadweave Boots (17586, -10.60 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (378.8 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.71 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -1.06 DPS) [vendor]; Wrath of Cenarius (21190, -4.46 DPS, sim-verified) [quest] |
| finger2 | Ritssyn's Ring of Chaos (21836) | World drop [world_drop] | sim-verified (378.8 DPS) | yes | Blessed Band of Light (272407, -0.29 DPS) [vendor]; Mindtear Band (20632, -1.06 DPS) [world]; Wrath of Cenarius (21190, -3.50 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (378.8 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Abyss Shard (20534, -0.71 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (378.8 DPS) | yes | Frozen Heart of the Mountain (249469, -1.39 DPS) [crafted]; Rune of the Guard Captain (19120, -2.81 DPS) [quest]; Abyss Shard (20534, -4.31 DPS, sim-verified) [quest] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (378.8 DPS) | yes | Shortsword of Vengeance (754, +0.00 DPS, sim-verified) [world_drop]; Soul Harvester (20536, +0.00 DPS) [quest]; High Warlord's War Staff (234549, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (384.4 DPS) | yes | Cold Snap (19130, -5.61 DPS, sim-verified) [world]; Wand of Biting Cold (19108, -14.88 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, -15.03 DPS) [dungeon] |

**New at 60:** head: Heretic Cowl; neck: Orb of the Darkmoon; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Rockfury Bracers; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Ritssyn's Ring of Chaos; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Staff of Dar'Orahil; ranged: Torch of Light

No-known-source sample (15 of 938, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

