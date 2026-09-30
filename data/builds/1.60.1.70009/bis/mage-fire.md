# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 32.0. Weights run: 0.8s. Verify run: 0.6s. 128 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.392 ± 0.085, crit=0.972 ± 0.082, hit=1.826 ± 0.028, spell_haste=-0.708 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (31.7 DPS) | yes | Shadow Goggles (4373, -0.93 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.5 spell_power points (1.32 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.42 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.02 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.2 spell_power points (0.46 DPS) | yes | Heavy Woolen Cloak (4311, -0.16 DPS) [crafted]; Feyscale Cloak (6632, -0.24 DPS) [dungeon]; Sanguine Cape (14376, -0.30 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 12.0 spell_power points (0.90 DPS) | yes | Mystic's Wrap (14369, -0.17 DPS) [world_drop]; Mystic's Robe (14371, -0.17 DPS) [world_drop]; Gray Woolen Robe (2585, -0.48 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 8.4 spell_power points (0.63 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.21 DPS) [world_drop]; Mystic's Bracelets (14366, -0.42 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 9.7 spell_power points (0.73 DPS) | yes | Tomb Robber's Gloves (280096, -0.06 DPS, sim-verified) [quest]; Pristine Gloves (253913, -0.12 DPS) [crafted]; Serpent Gloves (5970, -0.21 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (31.3 DPS) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Tarantula Silk Sash (3229, -0.20 DPS) [world]; Keller's Girdle (2911, -0.52 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 20.1 spell_power points (1.52 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.78 DPS) [world_drop]; Filigreed Silky Leggings (253939, -0.89 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 12.6 spell_power points (0.95 DPS) | yes | Pristine Boots (253889, -0.10 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.53 DPS) [world]; Kimbra Boots (6191, -0.53 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 8.4 spell_power points (0.63 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; Lorekeeper's Ring (20431, -0.25 DPS) [rep]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.8 spell_power points (0.59 DPS) | yes | Lorekeeper's Ring (20431, -0.21 DPS) [rep]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop]; Loop of Sacrifice (281673, -0.32 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 13.9 spell_power points (1.05 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.21 DPS) [world]; Lesser Staff of the Spire (1300, -0.42 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 299.5 spell_power points (22.54 DPS) | yes | Skycaller (12984, -0.63 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Blight Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 128, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 61.3. Weights run: 1.0s. Verify run: 0.6s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=1.030 ± 0.168, crit=2.538 ± 0.231, hit=2.838 ± 0.050, spell_haste=not significant (0.015 ± 0.231), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 16.3 spell_power points (1.50 DPS) | yes | Nightsky Cowl (4039, -0.43 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.46 DPS) [crafted]; Resilient Cap (14401, -0.46 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.2 spell_power points (1.21 DPS) | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.84 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.03 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.3 spell_power points (1.68 DPS) | yes | Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.37 DPS) [world_drop]; Death Speaker Mantle (6685, -0.65 DPS, sim-verified) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 8.2 spell_power points (0.76 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.10 DPS) [quest]; Resilient Cape (14400, -0.19 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 22.4 spell_power points (2.06 DPS) | yes | Mechbuilder's Overalls (9508, -0.64 DPS) [dungeon]; Death Speaker Robes (6682, -0.68 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.75 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.83 DPS) | yes | Nightsky Wristbands (6407, -0.26 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.26 DPS) [quest]; Glowing Magical Bracelets (13106, -0.95 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 15.3 spell_power points (1.41 DPS) | yes | Truefaith Gloves (7049, -0.67 DPS) [crafted]; Blight Gloves (279877, -0.75 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.97 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 14.1 spell_power points (1.30 DPS) | yes | Invoker's Cord (215366, -0.18 DPS) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Crimson Silk Belt (7055, -0.22 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (61.3 DPS) | yes | Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.20 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.74 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 14.2 spell_power points (1.31 DPS) | yes | Spidersilk Boots (4320, -0.28 DPS) [crafted]; Frothing Slippers (254003, -0.65 DPS) [crafted]; Acidic Walkers (9454, -1.20 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 7.2 spell_power points (0.66 DPS) | yes | Minor Channeling Ring (1449, -0.01 DPS) [quest]; Lorekeeper's Ring (19525, -0.02 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.2 spell_power points (0.66 DPS) | yes | Lorekeeper's Ring (19525, -0.02 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Minor Channeling Ring (1449, -0.56 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (60.6 DPS) | yes | Talisman of Arathor (21119, -2.52 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 11.3 spell_power points (1.04 DPS) | yes | Twisted Chanter's Staff (890, -0.02 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Channeler's Staff (4437, -0.28 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 363.3 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.04 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 99.9. Weights run: 1.1s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (-0.083 ± 0.167), crit=2.791 ± 0.286, hit=3.869 ± 0.072, spell_haste=6.535 ± 0.548, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.64 DPS) | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -1.26 DPS) [world_drop]; Silk Headband (7050, -1.51 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 spell_power points (0.88 DPS) | yes | Necklace of Calisea (1714, -0.39 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (1.13 DPS) | yes | Green Silken Shoulders (7057, -0.11 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.25 DPS) [quest]; Moonlit Amice (11884, -0.25 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (0.88 DPS) | yes | Long Silken Cloak (4326, +0.00 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.13 DPS) [crafted]; Caretaker's Cape (19532, -0.13 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.0 spell_power points (2.77 DPS) | yes | Dreamweave Vest (10021, -0.50 DPS) [crafted]; Robe of Power (7054, -1.01 DPS) [crafted]; Elemental Raiment (9434, -1.95 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (99.9 DPS) | yes | Condor Bracers (15864, -0.25 DPS) [quest]; Earthen Silk Cuffs (254019, -0.63 DPS) [crafted]; Arcane Runed Bracers (4744, -1.29 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 spell_power points (2.26 DPS) | yes | Red Mageweave Gloves (10018, -0.88 DPS) [crafted]; Gilded Handwraps (254021, -1.26 DPS) [crafted]; Black Mageweave Gloves (10003, -1.99 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.0 spell_power points (1.76 DPS) | yes | Star Belt (4329, -0.27 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -0.38 DPS) [rep]; Belt of Arugal (6392, -0.63 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.0 spell_power points (1.76 DPS) | yes | Abomination Skin Leggings (23173, -0.63 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.75 DPS) [crafted]; Gaze Dreamer Pants (6903, -3.24 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.02 DPS) | yes | Gilded Slippers (254001, -2.14 DPS) [crafted]; Spidersilk Boots (4320, -2.24 DPS, sim-verified) [crafted]; Nimbus Boots (6998, -2.26 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.0 spell_power points (1.26 DPS) | yes | Ring of Forlorn Spirits (2043, -0.25 DPS) [quest]; Reedknot Ring (9622, -0.38 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.50 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.13 DPS) | yes | Reedknot Ring (9622, -0.25 DPS) [quest]; Lorekeeper's Ring (19525, -0.25 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.66 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (98.6 DPS) | yes | Gut Ripper (2164, -2.33 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 313.0 spell_power points (39.38 DPS) | yes | Nether Force Wand (11263, -1.01 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.19 DPS) [quest]; Plaguerot Sprig (10766, -2.19 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 140.8. Weights run: 0.6s. Verify run: 0.7s. 403 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.053 ± 1.050), crit=4.846 ± 0.439, hit=7.393 ± 0.114, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 94.5 spell_power points (11.99 DPS) | yes | Eye of Theradras (17715, -0.39 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -6.91 DPS) [crafted]; Dreamweave Circlet (10041, -7.99 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.7 spell_power points (1.87 DPS) | yes | Scorn's Icy Choker (23169, -0.11 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.31 DPS) [quest]; Gemshard Heart (17707, -0.53 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (135.8 DPS) | yes | Red Mageweave Shoulders (10029, -1.16 DPS) [crafted]; Inquisitor's Shawl (19507, -1.43 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.41 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.3 spell_power points (2.58 DPS) | yes | Darkspear Raider's Cloak (272076, -0.71 DPS) [vendor]; Big Voodoo Cloak (8216, -0.74 DPS) [crafted]; Runecloth Cloak (13860, -0.92 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (136.1 DPS) | yes | Runecloth Robe (13858, -1.42 DPS) [crafted]; Runecloth Tunic (13857, -1.46 DPS) [crafted]; Knight's Dreadweave Vest (220886, -2.76 DPS, sim-verified) [vendor] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (136.6 DPS) | yes | Nethergeld Cuffs (254061, -0.01 DPS) [crafted]; Forgotten Wraps (9433, -0.23 DPS) [world_drop]; Aristocratic Cuffs (12546, -3.24 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 100.7 spell_power points (12.78 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -9.92 DPS) [vendor]; Dreamweave Gloves (10019, -9.96 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (135.8 DPS) | yes | Satyrmane Sash (17755, -0.19 DPS) [dungeon]; Deathmage Sash (10771, -0.41 DPS) [dungeon]; Highlander's Cloth Girdle (20097, -2.46 DPS, sim-verified) [rep] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 92.5 spell_power points (11.74 DPS) | yes | Red Mageweave Pants (10009, -0.22 DPS, sim-verified) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.95 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -8.98 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 91.4 spell_power points (11.60 DPS) | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -9.00 DPS) [crafted]; Southsea Mojo Boots (20641, -9.11 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 73.9 spell_power points (9.38 DPS) | yes | Brainlash (6440, -7.38 DPS) [dungeon]; Band of the Unicorn (7553, -7.73 DPS) [world_drop]; Mindseye Circle (10634, -7.78 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.3 spell_power points (2.07 DPS) | yes | Band of the Unicorn (7553, -0.42 DPS) [world_drop]; Mindseye Circle (10634, -0.47 DPS) [dungeon]; Brainlash (6440, -3.02 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (133.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (133.4 DPS) | yes | Illusionary Rod (7713, -0.80 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.60 DPS) [dungeon]; Shortsword of Vengeance (754, -1.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -24.56 DPS, sim-verified) [quest] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Guardian Talisman; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 390.9. Weights run: 1.0s. Verify run: 0.9s. 952 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.696 ± 0.593), crit=6.656 ± 0.541, hit=12.117 ± 0.288, spell_haste=13.681 ± 1.123, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (359.2 DPS) | yes | Bloodvine Goggles (19999, -4.29 DPS, sim-verified) [crafted]; Sorcerer's Crown (226935, -5.91 DPS) [quest]; Field Marshal's Coronet (16441, -6.24 DPS) [vendor] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (354.9 DPS) | yes | Amulet of the Dawn (22657, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Beads of Ogre Might (22150, -10.40 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 159.6 spell_power points (32.75 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -4.72 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.62 DPS) [crafted]; Shroud of the Nathrezim (18720, -6.63 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 150.7 spell_power points (30.93 DPS) | yes | Howler's Furs (272414, -6.07 DPS) [vendor]; Stalwart Cloak (272415, -6.07 DPS) [vendor]; Earthweave Cloak (21187, -10.63 DPS, sim-verified) [quest] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 291.4 spell_power points (59.78 DPS) | yes | Fireleaf Robe (240059, +0.00 DPS, sim-verified) [vendor]; Fireleaf Garb (240051, -20.37 DPS) [vendor]; Earthpower Vest (21183, -27.88 DPS) [quest] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 155.3 spell_power points (31.87 DPS) | yes | Fireleaf Wristwraps (240044, -2.83 DPS) [vendor]; Rockfury Bracers (21186, -16.13 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -24.57 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 252.1 spell_power points (51.73 DPS) | yes | Gloves of Spell Mastery (14146, -15.80 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -17.74 DPS) [vendor]; Sorcerer's Gloves (22066, -19.53 DPS) [quest] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 189.2 spell_power points (38.82 DPS) | yes | Fireleaf Waistguard (240045, -5.81 DPS) [vendor]; Knowledge of the Timbermaw (228190, -5.98 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -10.03 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 235.3 spell_power points (48.28 DPS) | yes | Fireleaf Leggings (240055, +0.00 DPS, sim-verified) [vendor]; Bloodvine Leggings (19683, -13.74 DPS) [crafted]; Sorcerer's Leggings (226933, -13.96 DPS) [quest] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 167.3 spell_power points (34.33 DPS) | yes | Fireleaf Boots (240050, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (16437, -0.29 DPS) [vendor]; Marshal's Silk Footwraps (231606, -0.29 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (354.9 DPS) | yes | Wrath of Cenarius (21190, +0.00 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -10.24 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -11.00 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (354.9 DPS) | yes | Wrath of Cenarius (21190, +0.00 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -10.24 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -11.00 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (354.9 DPS) | yes | Weakness Analyzer (272438, -17.86 DPS) [vendor]; Uther's Strength (11302, -21.14 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (354.9 DPS) | yes | Uther's Strength (11302, -1.64 DPS) [world_drop]; Weakness Analyzer (272438, -2.02 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (354.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Sageblade (22383, -14.77 DPS) [crafted]; Shortsword of Vengeance (754, -18.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (381.6 DPS) | yes | Oblivion's Touch (18761, -11.66 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.53 DPS) [world]; Cold Snap (19130, -26.64 DPS, sim-verified) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Bloodvine Vest; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Sentinel's Silk Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Frozen Heart of the Mountain; trinket2: Serenity Field; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 952, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 30.5. Weights run: 0.8s. Verify run: 0.6s. 124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.392 ± 0.085, crit=0.972 ± 0.082, hit=1.826 ± 0.028, spell_haste=-0.708 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (28.9 DPS) | yes | Shadow Goggles (4373, -0.75 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.5 spell_power points (1.32 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.02 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.2 spell_power points (0.46 DPS) | yes | Sanguine Cape (14376, -0.16 DPS, sim-verified) [world_drop]; Heavy Woolen Cloak (4311, -0.16 DPS) [crafted]; Feyscale Cloak (6632, -0.24 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 12.0 spell_power points (0.90 DPS) | yes | Mystic's Wrap (14369, -0.17 DPS) [world_drop]; Mystic's Robe (14371, -0.17 DPS) [world_drop]; Gray Woolen Robe (2585, -0.72 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (28.5 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.10 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.35 DPS, sim-verified) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 9.7 spell_power points (0.73 DPS) | yes | Tomb Robber's Gloves (280096, -0.08 DPS, sim-verified) [quest]; Pristine Gloves (253913, -0.12 DPS) [crafted]; Serpent Gloves (5970, -0.21 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (28.8 DPS) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Tarantula Silk Sash (3229, -0.20 DPS) [world]; Keller's Girdle (2911, -0.63 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (28.5 DPS) | yes | Abomination Skin Leggings (23173, -0.32 DPS, sim-verified) [dungeon]; Darkweave Breeches (12987, -0.35 DPS) [world_drop]; Filigreed Silky Leggings (253939, -0.45 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 12.6 spell_power points (0.95 DPS) | yes | Pristine Boots (253889, -0.09 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.53 DPS) [world]; Sanguine Sandals (14374, -0.53 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 8.4 spell_power points (0.63 DPS) | yes | Loop of Sacrifice (281673, -0.26 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; Sludge-Stained Band (286535, -0.40 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (28.7 DPS) | yes | Volcanic Rock Ring (12053, -0.06 DPS) [world_drop]; Sludge-Stained Band (286535, -0.15 DPS) [world]; Loop of Sacrifice (281673, -0.58 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 13.9 spell_power points (1.05 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.21 DPS) [world]; Lesser Staff of the Spire (1300, -0.42 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 299.5 spell_power points (22.54 DPS) | yes | Skycaller (12984, -0.68 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (orc, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 56.2. Weights run: 1.0s. Verify run: 0.6s. 217 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=1.030 ± 0.168, crit=2.538 ± 0.231, hit=2.838 ± 0.050, spell_haste=not significant (0.015 ± 0.231), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 16.3 spell_power points (1.50 DPS) | yes | Nightsky Cowl (4039, -0.33 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.46 DPS) [crafted]; Resilient Cap (14401, -0.46 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.2 spell_power points (1.21 DPS) | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.84 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.97 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.3 spell_power points (1.68 DPS) | yes | Death Speaker Mantle (6685, -0.27 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.37 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 8.2 spell_power points (0.76 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.19 DPS) [world_drop]; Soft Willow Cape (16661, -0.28 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 22.4 spell_power points (2.06 DPS) | yes | Death Speaker Robes (6682, -0.63 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.64 DPS) [dungeon]; Pristine Gown (253961, -0.75 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.83 DPS) | yes | Nightsky Wristbands (6407, -0.26 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.26 DPS) [quest]; Glowing Magical Bracelets (13106, -0.66 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 11.2 spell_power points (1.03 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.16 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.28 DPS) [crafted]; Blight Gloves (279877, -0.36 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 14.1 spell_power points (1.30 DPS) | yes | Invoker's Cord (215366, -0.18 DPS) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Crimson Silk Belt (7055, -0.37 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (56.2 DPS) | yes | Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.20 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.82 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 14.2 spell_power points (1.31 DPS) | yes | Spidersilk Boots (4320, -0.28 DPS) [crafted]; Frothing Slippers (254003, -0.65 DPS) [crafted]; Acidic Walkers (9454, -0.93 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 7.2 spell_power points (0.66 DPS) | yes | Advisor's Ring (19521, -0.02 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.11 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.2 spell_power points (0.66 DPS) | yes | Advisor's Ring (19521, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.11 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (55.4 DPS) | yes | Defiler's Talisman (21120, -2.14 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 11.3 spell_power points (1.04 DPS) | yes | Twisted Chanter's Staff (890, -0.08 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Channeler's Staff (4437, -0.28 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 363.3 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.23 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 91.0. Weights run: 1.1s. Verify run: 0.6s. 300 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (-0.083 ± 0.167), crit=2.791 ± 0.286, hit=3.869 ± 0.072, spell_haste=6.535 ± 0.548, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.64 DPS) | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -1.26 DPS) [world_drop]; Silk Headband (7050, -1.51 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 spell_power points (0.88 DPS) | yes | Necklace of Calisea (1714, +0.00 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (1.13 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Chestnut Mantle (17695, -0.13 DPS) [quest]; Berylline Pads (4197, -0.25 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (0.88 DPS) | yes | Long Silken Cloak (4326, +0.00 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.13 DPS) [crafted]; Battle Healer's Cloak (19528, -0.13 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.0 spell_power points (2.77 DPS) | yes | Dreamweave Vest (10021, -0.50 DPS) [crafted]; Elemental Raiment (9434, -0.79 DPS, sim-verified) [world_drop]; Robe of Power (7054, -1.01 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.13 DPS) | yes | Radiant Silver Bracers (4545, -0.63 DPS) [quest]; Earthen Silk Cuffs (254019, -0.63 DPS) [crafted]; Condor Bracers (15864, -1.22 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 spell_power points (2.26 DPS) | yes | Red Mageweave Gloves (10018, -0.88 DPS) [crafted]; Black Mageweave Gloves (10003, -1.07 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.26 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.0 spell_power points (1.76 DPS) | yes | Star Belt (4329, +0.00 DPS, sim-verified) [crafted]; Warsong Sash (16975, -0.38 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.38 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.0 spell_power points (1.76 DPS) | yes | Abomination Skin Leggings (23173, -0.63 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.75 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.46 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.02 DPS) | yes | Spidersilk Boots (4320, -1.29 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -2.14 DPS) [crafted]; Boots of the Enchanter (4325, -2.39 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.0 spell_power points (1.26 DPS) | yes | Reedknot Ring (9622, -0.38 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.50 DPS) [vendor]; Electrocutioner Lagnut (9447, -0.88 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.13 DPS) | yes | Advisor's Ring (19521, -0.25 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor]; Reedknot Ring (9622, -1.22 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (91.0 DPS) | yes | Gut Ripper (2164, -0.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 313.0 spell_power points (39.38 DPS) | yes | Nether Force Wand (11263, -1.17 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.19 DPS) [quest]; Plaguerot Sprig (10766, -2.19 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 128.2. Weights run: 0.6s. Verify run: 0.7s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.053 ± 1.050), crit=4.846 ± 0.439, hit=7.393 ± 0.114, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 94.5 spell_power points (11.99 DPS) | yes | Eye of Theradras (17715, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -6.91 DPS) [crafted]; Dreamweave Circlet (10041, -7.99 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.7 spell_power points (1.87 DPS) | yes | Scorn's Icy Choker (23169, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.31 DPS) [quest]; Gemshard Heart (17707, -0.53 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (121.8 DPS) | yes | Red Mageweave Shoulders (10029, -1.16 DPS) [crafted]; Inquisitor's Shawl (19507, -1.43 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.80 DPS, sim-verified) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.5 spell_power points (2.73 DPS) | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.51 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.85 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (122.1 DPS) | yes | Runecloth Robe (13858, -1.42 DPS) [crafted]; Runecloth Tunic (13857, -1.46 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -2.11 DPS, sim-verified) [vendor] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (123.0 DPS) | yes | Nethergeld Cuffs (254061, -0.01 DPS) [crafted]; Forgotten Wraps (9433, -0.23 DPS) [world_drop]; Aristocratic Cuffs (12546, -3.01 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 100.7 spell_power points (12.78 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -9.92 DPS) [vendor]; Dreamweave Gloves (10019, -9.96 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (121.7 DPS) | yes | Satyrmane Sash (17755, -0.19 DPS) [dungeon]; Deathmage Sash (10771, -0.41 DPS) [dungeon]; Defiler's Cloth Girdle (20165, -1.73 DPS, sim-verified) [rep] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 92.5 spell_power points (11.74 DPS) | yes | Red Mageweave Pants (10009, -0.10 DPS, sim-verified) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.95 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -8.98 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (121.2 DPS) | yes | Gilded Sandals (254107, -0.45 DPS) [crafted]; Southsea Mojo Boots (20641, -0.56 DPS) [quest]; First Sergeant's Dreadweave Boots (220909, -1.25 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 73.9 spell_power points (9.38 DPS) | yes | Brainlash (6440, -7.38 DPS) [dungeon]; Band of the Unicorn (7553, -7.73 DPS) [world_drop]; Mindseye Circle (10634, -7.78 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.3 spell_power points (2.07 DPS) | yes | Band of the Unicorn (7553, -0.42 DPS) [world_drop]; Mindseye Circle (10634, -0.47 DPS) [dungeon]; Brainlash (6440, -2.01 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (120.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted]; Uther's Strength (11302, -5.81 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (120.0 DPS) | yes | Illusionary Rod (7713, -0.80 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.60 DPS) [dungeon]; Shortsword of Vengeance (754, -1.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -24.75 DPS, sim-verified) [quest] |

**New at 50:** head: Blood Guard's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Guardian Talisman; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 363.8. Weights run: 1.0s. Verify run: 0.9s. 946 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.696 ± 0.593), crit=6.656 ± 0.541, hit=12.117 ± 0.288, spell_haste=13.681 ± 1.123, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (324.2 DPS) | yes | Sorcerer's Crown (226935, -5.91 DPS) [quest]; Bloodvine Goggles (19999, -6.00 DPS, sim-verified) [crafted]; Warlord's Silk Cowl (16533, -6.24 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (318.2 DPS) | yes | Jewel of Kajaro (19601, -1.07 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -5.74 DPS) [quest]; Amulet of the Dawn (22657, -17.26 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 159.6 spell_power points (32.75 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -4.26 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.62 DPS) [crafted]; Shroud of the Nathrezim (18720, -6.63 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 150.7 spell_power points (30.93 DPS) | yes | Earthweave Cloak (21187, -5.76 DPS, sim-verified) [quest]; Howler's Furs (272414, -6.07 DPS) [vendor]; Stalwart Cloak (272415, -6.07 DPS) [vendor] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 291.4 spell_power points (59.78 DPS) | yes | Fireleaf Robe (240059, +0.00 DPS, sim-verified) [vendor]; Fireleaf Garb (240051, -20.37 DPS) [vendor]; Earthpower Vest (21183, -27.88 DPS) [quest] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 155.3 spell_power points (31.87 DPS) | yes | Fireleaf Wristwraps (240044, -2.83 DPS) [vendor]; Rockfury Bracers (21186, -6.33 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -24.57 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 252.1 spell_power points (51.73 DPS) | yes | Gloves of Spell Mastery (14146, -8.55 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -17.74 DPS) [vendor]; Sorcerer's Gloves (22066, -19.53 DPS) [quest] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 189.2 spell_power points (38.82 DPS) | yes | Knowledge of the Timbermaw (228190, -3.98 DPS, sim-verified) [vendor]; Fireleaf Waistguard (240045, -5.81 DPS) [vendor]; Belt of the Archmage (18405, -10.03 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (326.9 DPS) | yes | Bloodvine Leggings (19683, -3.50 DPS) [crafted]; Sorcerer's Leggings (226933, -3.71 DPS) [quest]; Sentinel's Silk Leggings (237815, -8.75 DPS, sim-verified) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 167.3 spell_power points (34.33 DPS) | yes | General's Silk Boots (16539, -0.29 DPS) [vendor]; General's Silk Boots (231597, -0.29 DPS) [vendor]; Fireleaf Boots (240050, -0.93 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (318.2 DPS) | yes | Wrath of Cenarius (21190, -1.89 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -10.24 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -11.00 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (318.2 DPS) | yes | Wrath of Cenarius (21190, -1.89 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -10.24 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -11.00 DPS) [vendor] |
| trinket1 | Darkmoon Card: Blue Dragon (19288) | Darkmoon Beast Deck [quest] | sim-verified (318.2 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (318.2 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -1.83 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (318.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -3.21 DPS, sim-verified) [world_drop]; Sageblade (22383, -14.77 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (350.9 DPS) | yes | Oblivion's Touch (18761, -11.66 DPS) [dungeon]; Sparkling Crystal Wand (20672, -12.53 DPS) [world]; Cold Snap (19130, -32.71 DPS, sim-verified) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Beads of Ogre Might; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Bloodvine Vest; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Blue Dragon; trinket2: Serenity Field; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 946, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

