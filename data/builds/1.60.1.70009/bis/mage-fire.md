# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 32.1. Weights run: 0.7s. Verify run: 0.8s. 123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.495 ± 0.285, crit=1.741 ± 0.088, hit=4.115 ± 0.204, spell_haste=not significant (0.159 ± 0.399), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (31.8 DPS) | yes | Flying Tiger Goggles (4368, -0.45 DPS) [crafted]; Lucky Fishing Hat (19972, -0.45 DPS) [quest]; Shadow Goggles (4373, -0.77 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 18.5 | yes | Reinforced Woolen Shoulders (4315, -0.45 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.09 DPS) [crafted]; Slime-encrusted Pads (6461, -1.39 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.5 | yes | Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Feyscale Cloak (6632, -0.26 DPS) [dungeon]; Sanguine Cape (14376, -0.33 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 12.5 | yes | Mystic's Wrap (14369, -0.15 DPS) [world_drop]; Mystic's Robe (14371, -0.15 DPS) [world_drop]; Gray Woolen Robe (2585, -0.50 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 9.0 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.23 DPS) [world_drop]; Mystic's Bracelets (14366, -0.45 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 10.5 | yes | Pristine Gloves (253913, -0.15 DPS) [crafted]; Tomb Robber's Gloves (280096, -0.17 DPS, sim-verified) [quest]; Adept's Gloves (4768, -0.23 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (31.5 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Tarantula Silk Sash (3229, -0.19 DPS) [world]; Keller's Girdle (2911, -0.53 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 21.0 | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.79 DPS) [world_drop]; Filigreed Silky Leggings (253939, -0.90 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 13.0 | yes | Pristine Boots (253889, -0.22 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.53 DPS) [world]; Kimbra Boots (6191, -0.53 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 9.0 | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Lorekeeper's Ring (20431, -0.30 DPS) [rep]; Volcanic Rock Ring (12053, -0.34 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 8.0 | yes | Lorekeeper's Ring (20431, -0.23 DPS) [rep]; Volcanic Rock Ring (12053, -0.26 DPS) [world_drop]; Loop of Sacrifice (281673, -0.42 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 285.9 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Living Root (6631, -0.36 DPS) [dungeon]; Staff of Westfall (2042, -0.50 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 298.7 | yes | Skycaller (12984, -0.66 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Blight Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 123, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 20434 Lorekeeper's Staff; 20440 Protector's Sword; 20443 Sentinel's Blade; 209618 Insignia of the Alliance; 209623 Insignia of the Horde

### Band 30 (gnome, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 60.9. Weights run: 0.8s. Verify run: 0.8s. 209 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.956 ± 0.566), crit=3.069 ± 0.225, hit=4.419 ± 0.395, spell_haste=not significant (0.233 ± 0.614), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.6 | yes | Nightsky Cowl (4039, +0.00 DPS, sim-verified) [world_drop]; Holy Shroud (2721, -0.44 DPS) [world_drop]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.90 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.6 | yes | Death Speaker Mantle (6685, -0.06 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.7 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.08 DPS) [quest]; Resilient Cape (14400, -0.18 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.4 | yes | Death Speaker Robes (6682, -0.35 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.74 DPS) [crafted]; Tree Bark Jacket (1486, -0.81 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.31 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.31 DPS) [quest]; Glowing Magical Bracelets (13106, -0.91 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 14.5 | yes | Truefaith Gloves (7049, -0.40 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.66 DPS) [dungeon]; Serpent Gloves (5970, -0.72 DPS) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.9 | yes | Crimson Silk Belt (7055, +0.00 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (60.7 DPS) | yes | Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Abomination Skin Leggings (23173, -1.25 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.7 | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Black Widow Band (6199, -0.03 DPS) [world]; Snake Hoop (6750, -0.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 | yes | Black Widow Band (6199, +0.00 DPS, sim-verified) [world]; Snake Hoop (6750, -0.02 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (58.6 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Insignia of the Horde (18850, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | sim-verified (60.3 DPS) | yes | Lorekeeper's Staff (212580, -0.35 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.35 DPS) [vendor]; Gnarled Ash Staff (791, -0.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 349.6 | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 19549 Sentinel's Blade

### Band 40 (gnome, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 101.8. Weights run: 0.8s. Verify run: 0.8s. 286 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (0.249 ± 0.763), crit=2.900 ± 0.288, hit=5.720 ± 0.629, spell_haste=6.506 ± 0.958, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -1.31 DPS) [world_drop]; Silk Headband (7050, -1.57 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 | yes | Necklace of Calisea (1714, +0.00 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.88 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.95 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.2 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.03 DPS) [crafted]; Caretaker's Cape (19532, -0.16 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.5 | yes | Dreamweave Vest (10021, -0.43 DPS) [crafted]; Robe of Power (7054, -0.85 DPS) [crafted]; Elemental Raiment (9434, -1.15 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (101.8 DPS) | yes | Condor Bracers (15864, -0.26 DPS) [quest]; Earthen Silk Cuffs (254019, -0.65 DPS) [crafted]; Arcane Runed Bracers (4744, -1.50 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | Red Mageweave Gloves (10018, -0.72 DPS) [crafted]; Black Mageweave Gloves (10003, -1.14 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.21 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.0 | yes | Star Belt (4329, -0.33 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -0.43 DPS) [rep]; Deathmage Sash (10771, -0.56 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.0 | yes | Crimson Silk Pantaloons (7062, -0.75 DPS) [crafted]; Abomination Skin Leggings (23173, -0.78 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.86 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, +0.00 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.09 DPS) [crafted]; Acidic Walkers (9454, -2.23 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 | yes | Ring of Forlorn Spirits (2043, -0.46 DPS) [quest]; Reedknot Ring (9622, -0.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.72 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (99.0 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Rune of Duty (21567, -0.32 DPS, sim-verified) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 311.6 | yes | Illusionary Rod (7713, -0.02 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -4.84 DPS) [quest]; Black Duskwood Staff (937, -7.20 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 302.5 | yes | Nether Force Wand (11263, -1.42 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.29 DPS) [quest]; Ragefire Wand (7513, -2.34 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs

### Band 50 (gnome, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 139.1. Weights run: 0.8s. Verify run: 0.8s. 359 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.053 ± 1.050), crit=4.846 ± 0.439, hit=8.500 ± 0.859, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 94.5 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.84 DPS) [dungeon]; Red Mageweave Headband (10033, -6.91 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.7 | yes | Mindburst Medallion (11196, -0.31 DPS) [quest]; Darkspear Warding Pendant (272073, -0.67 DPS) [vendor]; Scorn's Icy Choker (23169, -1.13 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 85.3 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -6.77 DPS) [dungeon]; Red Mageweave Shoulders (10029, -7.93 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.3 | yes | Darkspear Raider's Cloak (272076, -0.71 DPS) [vendor]; Big Voodoo Cloak (8216, -0.74 DPS) [crafted]; Runecloth Cloak (13860, -1.21 DPS, sim-verified) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 91.4 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Acumen Robes (17775, -6.52 DPS) [quest]; Runecloth Robe (13858, -7.93 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (139.1 DPS) | yes | Forgotten Wraps (9433, -0.22 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.22 DPS) [quest]; Bloodband Bracers (11469, -1.41 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 111.7 | yes | Raider Handwraps (272098, -1.27 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -11.33 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -11.33 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 82.1 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.31 DPS) [dungeon]; Deathmage Sash (10771, -7.53 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 92.5 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Red Mageweave Pants (10009, -8.36 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.95 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 102.5 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -9.96 DPS) [crafted]; Gilded Sandals (254107, -10.41 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 85.0 | yes | Band of the Unicorn (7553, -9.14 DPS) [world_drop]; Lorekeeper's Ring (19523, -9.26 DPS) [rep]; Lorekeeper's Ring (19524, -9.65 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.3 | yes | Lorekeeper's Ring (19523, -0.55 DPS) [rep]; Lorekeeper's Ring (19524, -0.93 DPS) [rep]; Band of the Unicorn (7553, -1.67 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (136.8 DPS) | yes | Thunderbrew's Boot Flask (744, -0.78 DPS, sim-verified) [quest]; Tidal Charm (1404, -9.71 DPS) [vendor]; Guardian Talisman (1490, -9.71 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (136.8 DPS) | yes | Thunderbrew's Boot Flask (744, -0.65 DPS, sim-verified) [quest]; Tidal Charm (1404, -0.76 DPS) [vendor]; Guardian Talisman (1490, -0.76 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (136.8 DPS) | yes | Hanzo Sword (8190, -2.93 DPS, sim-verified) [world_drop]; Soulkeeper (1607, -8.02 DPS) [world_drop]; Spire of Hakkar (10844, -9.60 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 | yes | Noxious Shooter (17745, -0.55 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.73 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 359, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (gnome, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 378.0. Weights run: 0.9s. Verify run: 0.9s. 722 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.134 ± 1.662), crit=7.715 ± 0.552, hit=17.053 ± 1.684, spell_haste=not significant (-2.222 ± 2.347), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bloodvine Goggles (19999) | Engineering [crafted] | 449.1 | yes | Fireleaf Circlet (240056, +0.00 DPS, sim-verified) [vendor]; Field Marshal's Coronet (16441, -62.37 DPS) [vendor]; Warlord's Silk Cowl (16533, -62.37 DPS) [vendor] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (369.2 DPS) | yes | Amulet of the Dawn (22657, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Beads of Ogre Might (22150, -3.58 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 166.0 | yes | Rugged Mantle of the Timbermaw (227808, -4.66 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.67 DPS) [crafted]; Champion's Silk Mantle (23264, -6.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 195.6 | yes | Howler's Furs (272414, -5.41 DPS) [vendor]; Stalwart Cloak (272415, -5.41 DPS) [vendor]; Earthweave Cloak (21187, -16.80 DPS, sim-verified) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | sim-verified (378.0 DPS) | yes | Bloodvine Vest (19682, -11.87 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -22.65 DPS) [vendor]; Robe of the Archmage (14152, -30.01 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 197.5 | yes | Fireleaf Bindings (240052, -2.43 DPS, sim-verified) [vendor]; Fireleaf Wristwraps (240044, -10.35 DPS) [vendor]; Dryad's Wrist Bindings (19595, -35.95 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 273.9 | yes | Sorcerer's Gloves (22066, -16.30 DPS) [quest]; Sorcerer's Gauntlets (226930, -16.30 DPS) [vendor]; Gloves of Spell Mastery (14146, -17.50 DPS, sim-verified) [crafted] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 231.3 | yes | Knowledge of the Timbermaw (228190, -4.67 DPS, sim-verified) [vendor]; Fireleaf Waistguard (240045, -13.69 DPS) [vendor]; Belt of the Archmage (18405, -18.38 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 259.4 | yes | Bloodvine Leggings (19683, -6.74 DPS, sim-verified) [crafted]; Sorcerer's Leggings (226933, -11.16 DPS) [quest]; Fireleaf Leggings (240055, -15.08 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 207.7 | yes | Marshal's Silk Footwraps (16437, -0.06 DPS) [vendor]; General's Silk Boots (16539, -0.06 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -13.99 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (367.2 DPS) | yes | Wrath of Cenarius (21190, -8.05 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -15.08 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -15.75 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (367.2 DPS) | yes | Wrath of Cenarius (21190, -8.05 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -15.08 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -15.75 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (365.2 DPS) | yes | Weakness Analyzer (272438, -28.39 DPS) [vendor]; Uther's Strength (11302, -31.85 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -33.15 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (367.2 DPS) | yes | Uther's Strength (11302, -1.73 DPS) [world_drop]; Weakness Analyzer (272438, -2.00 DPS, sim-verified) [vendor]; Thunderbrew's Boot Flask (744, -3.02 DPS) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (367.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; The Lobotomizer (19324, -17.66 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | 296.3 | yes | Brilliant Wand (249385, -0.76 DPS) [crafted]; Stormrager (16997, -3.17 DPS, sim-verified) [quest]; Torch of Austen (13004, -7.21 DPS) [world_drop] |

**New at 60:** head: Bloodvine Goggles; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Rockfury Bracers; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Sentinel's Silk Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket2: Serenity Field; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 722, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 30.5. Weights run: 0.7s. Verify run: 0.8s. 122 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.495 ± 0.285, crit=1.741 ± 0.088, hit=4.115 ± 0.204, spell_haste=not significant (0.159 ± 0.399), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (28.9 DPS) | yes | Flying Tiger Goggles (4368, -0.45 DPS) [crafted]; Lucky Fishing Hat (19972, -0.45 DPS) [quest]; Shadow Goggles (4373, -0.74 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 18.5 | yes | Reinforced Woolen Shoulders (4315, -0.18 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.09 DPS) [crafted]; Slime-encrusted Pads (6461, -1.39 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.5 | yes | Sanguine Cape (14376, -0.06 DPS, sim-verified) [world_drop]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Feyscale Cloak (6632, -0.26 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 12.5 | yes | Mystic's Wrap (14369, -0.15 DPS) [world_drop]; Mystic's Robe (14371, -0.15 DPS) [world_drop]; Gray Woolen Robe (2585, -0.65 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (28.7 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.11 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.48 DPS, sim-verified) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 10.5 | yes | Tomb Robber's Gloves (280096, +0.00 DPS, sim-verified) [quest]; Pristine Gloves (253913, -0.15 DPS) [crafted]; Adept's Gloves (4768, -0.23 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (28.9 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Tarantula Silk Sash (3229, -0.19 DPS) [world]; Keller's Girdle (2911, -0.68 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (28.6 DPS) | yes | Darkweave Breeches (12987, -0.34 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.41 DPS, sim-verified) [dungeon]; Filigreed Silky Leggings (253939, -0.45 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 13.0 | yes | Pristine Boots (253889, +0.00 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.53 DPS) [world]; Sanguine Sandals (14374, -0.53 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 9.0 | yes | Loop of Sacrifice (281673, -0.19 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.34 DPS) [world_drop]; Sludge-Stained Band (286535, -0.45 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (28.8 DPS) | yes | Volcanic Rock Ring (12053, -0.04 DPS) [world_drop]; Sludge-Stained Band (286535, -0.15 DPS) [world]; Loop of Sacrifice (281673, -0.62 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 285.9 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Living Root (6631, -0.36 DPS) [dungeon]; Crescent Staff (6505, -1.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 298.7 | yes | Skycaller (12984, -0.59 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 122, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 209623 Insignia of the Horde; 241089 Scarlet Dagger; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest

### Band 30 (orc, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 55.6. Weights run: 0.8s. Verify run: 0.8s. 210 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.956 ± 0.566), crit=3.069 ± 0.225, hit=4.419 ± 0.395, spell_haste=not significant (0.233 ± 0.614), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.6 | yes | Holy Shroud (2721, -0.44 DPS) [world_drop]; Nightsky Cowl (4039, -0.46 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.86 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.6 | yes | Fairywing Mantle (9536, -0.29 DPS) [quest]; Death Speaker Mantle (6685, -0.35 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.7 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.18 DPS) [world_drop]; Hillman's Cloak (3719, -0.25 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.4 | yes | Death Speaker Robes (6682, -0.52 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.74 DPS) [crafted]; Tree Bark Jacket (1486, -0.81 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.31 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.31 DPS) [quest]; Glowing Magical Bracelets (13106, -0.93 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.8 | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.30 DPS) [dungeon]; Serpent Gloves (5970, -0.36 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.9 | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted]; Crimson Silk Belt (7055, -0.36 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (55.6 DPS) | yes | Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Abomination Skin Leggings (23173, -0.73 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.7 | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -0.83 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Snake Hoop (6750, -0.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 6.7 | yes | Snake Hoop (6750, +0.00 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.07 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (54.6 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Insignia of the Horde (18850, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Ash Staff (791) | World drop [world_drop] | 267.4 | yes | Glimmering Staff (249392, +0.00 DPS, sim-verified) [crafted]; Lorekeeper's Staff (212580, -0.41 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.41 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 349.6 | yes | Starfaller (13063, -0.22 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Black Widow Band; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Gnarled Ash Staff; ranged: Necrotic Wand

No-known-source sample (15 of 210, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 19545 Scout's Blade; 19553 Legionnaire's Sword; 19569 Advisor's Gnarled Staff

### Band 40 (orc, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 93.7. Weights run: 0.8s. Verify run: 0.7s. 287 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (0.249 ± 0.763), crit=2.900 ± 0.288, hit=5.720 ± 0.629, spell_haste=6.506 ± 0.958, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -1.31 DPS) [world_drop]; Silk Headband (7050, -1.57 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 | yes | Necklace of Calisea (1714, -0.09 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.88 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.95 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.2 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.03 DPS) [crafted]; Battle Healer's Cloak (19528, -0.16 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.5 | yes | Dreamweave Vest (10021, -0.43 DPS) [crafted]; Robe of Power (7054, -0.85 DPS) [crafted]; Elemental Raiment (9434, -1.44 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Radiant Silver Bracers (4545, -0.39 DPS) [quest]; Earthen Silk Cuffs (254019, -0.65 DPS) [crafted]; Condor Bracers (15864, -1.39 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | Red Mageweave Gloves (10018, -0.72 DPS) [crafted]; Gilded Handwraps (254021, -1.21 DPS) [crafted]; Black Mageweave Gloves (10003, -1.84 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.0 | yes | Defiler's Cloth Girdle (20164, -0.43 DPS) [rep]; Warsong Sash (16975, -0.52 DPS) [quest]; Star Belt (4329, -0.77 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.0 | yes | Crimson Silk Pantaloons (7062, -0.75 DPS) [crafted]; Abomination Skin Leggings (23173, -0.78 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.44 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.74 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.09 DPS) [crafted]; Acidic Walkers (9454, -2.23 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 | yes | Reedknot Ring (9622, -0.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.72 DPS) [vendor]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.39 DPS) [vendor]; Reedknot Ring (9622, -1.39 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (92.4 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Rune of Duty (21567, -0.97 DPS, sim-verified) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 311.6 | yes | Illusionary Rod (7713, -0.46 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -4.84 DPS) [quest]; Black Duskwood Staff (937, -7.20 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 302.5 | yes | Nether Force Wand (11263, -1.77 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.29 DPS) [quest]; Ragefire Wand (7513, -2.34 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 287, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (orc, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 124.1. Weights run: 0.8s. Verify run: 0.7s. 360 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.053 ± 1.050), crit=4.846 ± 0.439, hit=8.500 ± 0.859, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 94.5 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.84 DPS) [dungeon]; Red Mageweave Headband (10033, -6.91 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.7 | yes | Scorn's Icy Choker (23169, -0.27 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.31 DPS) [quest]; Darkspear Warding Pendant (272073, -0.67 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 85.3 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -6.77 DPS) [dungeon]; Red Mageweave Shoulders (10029, -7.93 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.5 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.51 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.85 DPS) [vendor] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 91.4 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Acumen Robes (17775, -6.52 DPS) [quest]; Runecloth Robe (13858, -7.93 DPS) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 14.5 | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Forgotten Wraps (9433, -0.23 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.23 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 111.7 | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -11.33 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -11.33 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 82.1 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.31 DPS) [dungeon]; Deathmage Sash (10771, -7.53 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 92.5 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Red Mageweave Pants (10009, -8.36 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.95 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 102.5 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -9.96 DPS) [crafted]; Gilded Sandals (254107, -10.41 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 85.0 | yes | Band of the Unicorn (7553, -9.14 DPS) [world_drop]; Advisor's Ring (19519, -9.26 DPS) [rep]; Advisor's Ring (19520, -9.65 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.3 | yes | Band of the Unicorn (7553, +0.00 DPS, sim-verified) [world_drop]; Advisor's Ring (19519, -0.55 DPS) [rep]; Advisor's Ring (19520, -0.93 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (123.5 DPS) | yes | Uther's Strength (11302, -0.43 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -9.71 DPS) [vendor]; Guardian Talisman (1490, -9.71 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (123.5 DPS) | yes | Uther's Strength (11302, -0.56 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -7.55 DPS) [vendor]; Guardian Talisman (1490, -7.55 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (123.5 DPS) | yes | Hanzo Sword (8190, -3.45 DPS, sim-verified) [world_drop]; Soulkeeper (1607, -8.02 DPS) [world_drop]; Spire of Hakkar (10844, -9.60 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 | yes | Noxious Shooter (17745, -0.04 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.73 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Deep Woodlands Cloak; chest: Knight's Dreadweave Vest; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 360, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 380.7. Weights run: 0.9s. Verify run: 0.9s. 725 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.134 ± 1.662), crit=7.715 ± 0.552, hit=17.053 ± 1.684, spell_haste=not significant (-2.222 ± 2.347), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (345.9 DPS) | yes | Field Marshal's Coronet (16441, -6.32 DPS) [vendor]; Warlord's Silk Cowl (16533, -6.32 DPS) [vendor]; Bloodvine Goggles (19999, -9.54 DPS, sim-verified) [crafted] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (344.5 DPS) | yes | Amulet of the Dawn (22657, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Beads of Ogre Might (22150, -3.61 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 166.0 | yes | Rugged Mantle of the Timbermaw (227808, -4.06 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.67 DPS) [crafted]; Champion's Silk Mantle (23264, -6.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 195.6 | yes | Howler's Furs (272414, -5.41 DPS) [vendor]; Stalwart Cloak (272415, -5.41 DPS) [vendor]; Earthweave Cloak (21187, -7.72 DPS, sim-verified) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | sim-verified (346.0 DPS) | yes | Bloodvine Vest (19682, -9.64 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -22.65 DPS) [vendor]; Robe of the Archmage (14152, -30.01 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | sim-verified (342.5 DPS) | yes | Fireleaf Wristwraps (240044, -2.49 DPS) [vendor]; Rockfury Bracers (21186, -6.17 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -28.09 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 273.9 | yes | Sorcerer's Gloves (22066, -16.30 DPS) [quest]; Sorcerer's Gauntlets (226930, -16.30 DPS) [vendor]; Gloves of Spell Mastery (14146, -18.22 DPS, sim-verified) [crafted] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 231.3 | yes | Knowledge of the Timbermaw (228190, +0.00 DPS, sim-verified) [vendor]; Fireleaf Waistguard (240045, -13.69 DPS) [vendor]; Belt of the Archmage (18405, -18.38 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 259.4 | yes | Bloodvine Leggings (19683, -5.31 DPS, sim-verified) [crafted]; Sorcerer's Leggings (226933, -11.16 DPS) [quest]; Fireleaf Leggings (240055, -15.08 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 207.7 | yes | Marshal's Silk Footwraps (16437, -0.06 DPS) [vendor]; General's Silk Boots (16539, -0.06 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -5.98 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (343.4 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Wrath of Cenarius (21190, -7.00 DPS, sim-verified) [quest]; Channeler's Ring (272406, -19.90 DPS) [vendor] |
| finger2 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (346.7 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.68 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.89 DPS) [vendor]; Band of Earthen Might (21182, -10.28 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (333.0 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.73 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (333.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Rune of the Guard Captain (19120, -1.28 DPS, sim-verified) [quest]; Uther's Strength (11302, -3.46 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (334.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; The Lobotomizer (19324, -9.32 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Stormrager (16997) | The Scarlet Oracle, Demetria [quest] | sim-verified (340.7 DPS) | yes | Brilliant Wand (249385, -0.68 DPS) [crafted]; Wand of Biting Cold (19108, -4.28 DPS, sim-verified) [quest]; Torch of Austen (13004, -7.12 DPS) [world_drop] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Sentinel's Silk Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Stormrager

No-known-source sample (15 of 725, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

