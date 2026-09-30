# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 31.9. Weights run: 0.7s. Verify run: 0.7s. 203 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.495 ± 0.285, crit=1.741 ± 0.088, hit=4.115 ± 0.204, spell_haste=not significant (0.159 ± 0.399), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (31.2 DPS) | yes | Flying Tiger Goggles (4368, -0.45 DPS) [crafted]; Lucky Fishing Hat (19972, -0.45 DPS) [quest]; Shadow Goggles (4373, -0.65 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 11.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.51 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.83 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.5 | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.26 DPS) [dungeon]; Black Whelp Cloak (7283, -0.26 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 12.5 | yes | Seer's Robe (2981, -0.26 DPS) [dungeon]; Manaweave Robe (7509, -0.26 DPS) [quest]; Gray Woolen Robe (2585, -0.52 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 9.0 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.23 DPS) [dungeon]; Repurposed Hair Band (281256, -0.45 DPS) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 10.5 | yes | Pristine Gloves (253913, -0.15 DPS) [crafted]; Adept's Gloves (4768, -0.23 DPS) [world]; Tomb Robber's Gloves (280096, -0.23 DPS, sim-verified) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (31.0 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Tarantula Silk Sash (3229, -0.19 DPS) [world]; Keller's Girdle (2911, -0.43 DPS, sim-verified) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 21.0 | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Filigreed Silky Leggings (253939, -0.90 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.90 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 13.0 | yes | Pristine Boots (253889, -0.12 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.53 DPS) [world]; Kimbra Boots (6191, -0.53 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 9.0 | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Lorekeeper's Ring (20431, -0.30 DPS) [rep]; Sludge-Stained Band (286535, -0.45 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 8.0 | yes | Lorekeeper's Ring (20431, -0.23 DPS) [rep]; Loop of Sacrifice (281673, -0.33 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -0.38 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 285.9 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Living Root (6631, -0.36 DPS) [dungeon]; Staff of Westfall (2042, -0.50 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 298.7 | yes | Firebelcher (5243, -0.40 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest]; Sizzle Stick (8071, -4.51 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Blight Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak; 9792 Ivycloth Boots

### Band 30 (gnome, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 60.9. Weights run: 0.8s. Verify run: 0.8s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.956 ± 0.566), crit=3.069 ± 0.225, hit=4.419 ± 0.395, spell_haste=not significant (0.233 ± 0.614), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.6 | yes | Nightsky Cowl (4039, +0.00 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -0.44 DPS) [dungeon]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [dungeon]; Darkspear Warding Pendant (272075, -0.90 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.22 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.6 | yes | Death Speaker Mantle (6685, -0.06 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Invoker's Mantle (215365, -0.56 DPS) [crafted] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.7 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.08 DPS) [quest]; Hillman's Cloak (3719, -0.25 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.4 | yes | Death Speaker Robes (6682, -0.35 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.74 DPS) [crafted]; Tree Bark Jacket (1486, -0.81 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.31 DPS) [quest]; Mindthrust Bracers (1974, -0.40 DPS) [dungeon]; Nightsky Wristbands (6407, -1.00 DPS, sim-verified) [dungeon] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 14.5 | yes | Truefaith Gloves (7049, -0.40 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.66 DPS) [dungeon]; Serpent Gloves (5970, -0.72 DPS) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.9 | yes | Crimson Silk Belt (7055, +0.00 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (60.7 DPS) | yes | Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Abomination Skin Leggings (23173, -1.25 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.7 | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Black Widow Band (6199, -0.03 DPS) [world]; Snake Hoop (6750, -0.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 | yes | Black Widow Band (6199, +0.00 DPS, sim-verified) [world]; Snake Hoop (6750, -0.02 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (58.6 DPS) | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | sim-verified (60.3 DPS) | yes | Lorekeeper's Staff (212580, -0.35 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.35 DPS) [vendor]; Gnarled Ash Staff (791, -0.86 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 349.6 | yes | Greater Mystic Wand (217287, -0.84 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest]; Scorching Wand (5213, -4.65 DPS) [dungeon] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 40 (gnome, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 98.7. Weights run: 0.8s. Verify run: 0.8s. 568 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (0.249 ± 0.763), crit=2.900 ± 0.288, hit=5.720 ± 0.629, spell_haste=6.506 ± 0.958, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -1.31 DPS) [dungeon]; Silk Headband (7050, -1.57 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 | yes | Darkspear Warding Pendant (272074, -0.88 DPS) [vendor]; Necklace of Calisea (1714, -0.89 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272075, -0.95 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.2 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.03 DPS) [crafted]; Caretaker's Cape (19532, -0.16 DPS) [rep] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 23.5 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.85 DPS) [crafted]; Crimson Silk Vest (7058, -1.18 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.26 DPS) [quest]; Earthen Silk Cuffs (254019, -0.65 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | Red Mageweave Gloves (10018, -0.72 DPS) [crafted]; Gilded Handwraps (254021, -1.21 DPS) [crafted]; Black Mageweave Gloves (10003, -1.34 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.0 | yes | Highlander's Cloth Girdle (20099, -0.43 DPS) [rep]; Star Belt (4329, -0.51 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -0.56 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.0 | yes | Crimson Silk Pantaloons (7062, -0.75 DPS) [crafted]; Abomination Skin Leggings (23173, -0.78 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.02 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.31 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.09 DPS) [crafted]; Acidic Walkers (9454, -2.23 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 | yes | Ring of Forlorn Spirits (2043, -0.46 DPS) [quest]; Reedknot Ring (9622, -0.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.72 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.91 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | Gnomeregan: Leprous Assistant [dungeon] | 311.6 | yes | Illusionary Rod (7713, -0.10 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -4.84 DPS) [quest]; Black Duskwood Staff (937, -7.20 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | 286.1 | yes | Ragefire Wand (7513, -0.19 DPS) [quest]; Icefury Wand (7514, -0.65 DPS, sim-verified) [quest]; Umbral Wand (5216, -1.76 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle; 7475 Regal Cuffs

### Band 50 (gnome, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 139.1. Weights run: 0.7s. Verify run: 0.8s. 722 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.053 ± 1.050), crit=4.846 ± 0.439, hit=8.500 ± 0.859, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 94.5 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.84 DPS) [dungeon]; Red Mageweave Headband (10033, -6.91 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.7 | yes | Mindburst Medallion (11196, -0.31 DPS) [quest]; Darkspear Warding Pendant (272073, -0.67 DPS) [vendor]; Scorn's Icy Choker (23169, -1.13 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 85.3 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -6.77 DPS) [dungeon]; Red Mageweave Shoulders (10029, -7.93 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.3 | yes | Darkspear Raider's Cloak (272076, -0.71 DPS) [vendor]; Big Voodoo Cloak (8216, -0.74 DPS) [crafted]; Runecloth Cloak (13860, -1.21 DPS, sim-verified) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 91.4 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Acumen Robes (17775, -6.52 DPS) [quest]; Runecloth Robe (13858, -7.93 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (139.1 DPS) | yes | Shizzle's Nozzle Wiper (11917, -0.22 DPS) [quest]; Imperial Red Bracers (8247, -0.35 DPS) [dungeon]; Bloodband Bracers (11469, -1.41 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 111.7 | yes | Raider Handwraps (272098, -1.27 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -11.33 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -11.33 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 82.1 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.31 DPS) [dungeon]; Deathmage Sash (10771, -7.53 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 92.5 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Red Mageweave Pants (10009, -8.36 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.95 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 102.5 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -9.96 DPS) [crafted]; Gilded Sandals (254107, -10.41 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 85.0 | yes | Lorekeeper's Ring (19523, -9.26 DPS) [rep]; Lorekeeper's Ring (19524, -9.65 DPS) [rep]; Ring of Forlorn Spirits (2043, -9.77 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.3 | yes | Lorekeeper's Ring (19523, +0.00 DPS, sim-verified) [rep]; Lorekeeper's Ring (19524, -0.93 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.06 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (136.8 DPS) | yes | Thunderbrew's Boot Flask (744, -0.78 DPS, sim-verified) [quest]; Guardian Talisman (1490, -9.71 DPS) [quest]; Ankh of Life (1713, -9.71 DPS) [dungeon] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (136.8 DPS) | yes | Thunderbrew's Boot Flask (744, -0.65 DPS, sim-verified) [quest]; Guardian Talisman (1490, -0.76 DPS) [quest]; Ankh of Life (1713, -0.76 DPS) [dungeon] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (136.8 DPS) | yes | Thrash Blade (17705, -2.93 DPS, sim-verified) [quest]; Soulkeeper (1607, -8.02 DPS) [dungeon]; Spire of Hakkar (10844, -9.60 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 | yes | Noxious Shooter (17745, -0.55 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.73 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 722, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 433.1. Weights run: 0.8s. Verify run: 0.9s. 1084 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.134 ± 1.662), crit=7.715 ± 0.552, hit=17.053 ± 1.684, spell_haste=not significant (-2.222 ± 2.347), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | sim-verified (383.6 DPS) | yes | Bloodvine Goggles (19999, -11.82 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -23.51 DPS) [quest]; Fireleaf Circlet (240056, -55.74 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | sim-verified (380.3 DPS) | yes | Jewel of Kajaro (19601, -1.58 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -23.33 DPS) [quest]; Medallion of the Dawn (22659, -36.83 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 166.0 | yes | Rugged Mantle of the Timbermaw (227808, -4.64 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.67 DPS) [crafted]; Champion's Silk Mantle (23264, -6.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 195.6 | yes | Howler's Furs (272414, -5.41 DPS) [vendor]; Stalwart Cloak (272415, -5.41 DPS) [vendor]; Earthweave Cloak (21187, -12.04 DPS, sim-verified) [quest] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | sim-verified (379.4 DPS) | yes | Bloodvine Vest (19682, -7.56 DPS, sim-verified) [crafted]; Fireleaf Robe (240059, -12.01 DPS) [vendor]; Zandalar Illusionist's Robe (20034, -28.38 DPS) [quest] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | sim-verified (379.0 DPS) | yes | Fireleaf Wristwraps (240044, -2.49 DPS) [vendor]; Rockfury Bracers (21186, -7.20 DPS, sim-verified) [quest]; Frostfire Bindings (22503, -25.30 DPS) [quest] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 273.9 | yes | Gloves of Spell Mastery (14146, -15.53 DPS, sim-verified) [crafted]; Sorcerer's Gloves (22066, -16.30 DPS) [quest]; Sorcerer's Gauntlets (226930, -16.30 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 231.3 | yes | Frostfire Belt (22502, -0.74 DPS, sim-verified) [quest]; Knowledge of the Timbermaw (228190, -2.33 DPS) [vendor]; Fireleaf Waistguard (240045, -13.69 DPS) [vendor] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | sim-verified (394.2 DPS) | yes | Bloodvine Leggings (19683, -6.84 DPS) [crafted]; Sorcerer's Leggings (226933, -8.28 DPS) [quest]; Sentinel's Silk Leggings (237815, -22.42 DPS, sim-verified) [vendor] |
| feet | Enigma Boots (21344) | Enigma Boots [quest] | 215.5 | yes | Marshal's Silk Footwraps (16437, -1.76 DPS) [vendor]; General's Silk Boots (16539, -1.76 DPS) [vendor]; Bloodvine Boots (19684, -11.68 DPS, sim-verified) [crafted] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 278.5 | yes | Band of Earthen Might (21182, -12.36 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -15.08 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -15.75 DPS) [vendor] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | sim-verified (386.7 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -1.21 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -1.89 DPS) [vendor]; Band of Earthen Might (21182, -14.90 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (380.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Thunderbrew's Boot Flask (744, -3.02 DPS) [quest]; Uther's Strength (11302, -7.28 DPS, sim-verified) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (380.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Thunderbrew's Boot Flask (744, -4.75 DPS) [quest]; Uther's Strength (11302, -5.30 DPS, sim-verified) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (380.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; The Lobotomizer (19324, -10.02 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | 296.3 | yes | Stormrager (16997, +0.00 DPS, sim-verified) [quest]; Brilliant Wand (249385, -0.76 DPS) [crafted]; Torch of Austen (13004, -7.21 DPS) [world_drop] |

**New at 60:** head: Frostfire Circlet; neck: Onyxia Tooth Pendant; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Frostfire Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1084, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 30.1. Weights run: 0.7s. Verify run: 0.7s. 202 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.495 ± 0.285, crit=1.741 ± 0.088, hit=4.115 ± 0.204, spell_haste=not significant (0.159 ± 0.399), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (28.4 DPS) | yes | Shadow Goggles (4373, -0.41 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.45 DPS) [crafted]; Lucky Fishing Hat (19972, -0.45 DPS) [quest] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 11.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.23 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.83 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (28.4 DPS) | yes | Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.38 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 12.5 | yes | Seer's Robe (2981, -0.26 DPS) [dungeon]; Lesser Spellfire Robes (7510, -0.26 DPS) [quest]; Gray Woolen Robe (2585, -0.43 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (28.4 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.11 DPS) [dungeon]; Tabitha's Cuffs (251486, -0.39 DPS, sim-verified) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 10.5 | yes | Tomb Robber's Gloves (280096, +0.00 DPS, sim-verified) [quest]; Pristine Gloves (253913, -0.15 DPS) [crafted]; Adept's Gloves (4768, -0.23 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (28.6 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Tarantula Silk Sash (3229, -0.19 DPS) [world]; Keller's Girdle (2911, -0.61 DPS, sim-verified) [dungeon] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (28.4 DPS) | yes | Abomination Skin Leggings (23173, -0.44 DPS, sim-verified) [dungeon]; Filigreed Silky Leggings (253939, -0.45 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.45 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 13.0 | yes | Pristine Boots (253889, -0.11 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.53 DPS) [world]; Feather Padded Treads (285345, -0.60 DPS) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 9.0 | yes | Sludge-Stained Band (286535, -0.45 DPS) [world]; Black Pearl Ring (6332, -0.45 DPS) [world]; Loop of Sacrifice (281673, -0.51 DPS, sim-verified) [quest] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (28.6 DPS) | yes | Sludge-Stained Band (286535, -0.15 DPS) [world]; Black Pearl Ring (6332, -0.15 DPS) [world]; Loop of Sacrifice (281673, -0.64 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 285.9 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Living Root (6631, -0.36 DPS) [dungeon]; Crescent Staff (6505, -1.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 298.7 | yes | Firebelcher (5243, -0.61 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest]; Sizzle Stick (8071, -4.51 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 202, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (orc, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 55.6. Weights run: 0.8s. Verify run: 0.8s. 410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.956 ± 0.566), crit=3.069 ± 0.225, hit=4.419 ± 0.395, spell_haste=not significant (0.233 ± 0.614), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.6 | yes | Holy Shroud (2721, -0.44 DPS) [dungeon]; Nightsky Cowl (4039, -0.46 DPS, sim-verified) [dungeon]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [dungeon]; Darkspear Warding Pendant (272075, -0.86 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.22 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.6 | yes | Fairywing Mantle (9536, -0.29 DPS) [quest]; Death Speaker Mantle (6685, -0.35 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.56 DPS) [crafted] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.7 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.25 DPS) [crafted]; Windsong Drape (15468, -0.25 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.4 | yes | Death Speaker Robes (6682, -0.52 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.74 DPS) [crafted]; Tree Bark Jacket (1486, -0.81 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.31 DPS) [quest]; Mindthrust Bracers (1974, -0.40 DPS) [dungeon]; Nightsky Wristbands (6407, -0.93 DPS, sim-verified) [dungeon] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.8 | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.30 DPS) [dungeon]; Serpent Gloves (5970, -0.36 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.9 | yes | Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted]; Crimson Silk Belt (7055, -0.36 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (55.6 DPS) | yes | Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Abomination Skin Leggings (23173, -0.73 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.7 | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -0.83 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Snake Hoop (6750, -0.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 6.7 | yes | Snake Hoop (6750, +0.00 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.07 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (54.6 DPS) | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Ash Staff (791) | Gnomeregan: Leprous Technician [dungeon] | 267.4 | yes | Glimmering Staff (249392, +0.00 DPS, sim-verified) [crafted]; Lorekeeper's Staff (212580, -0.41 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.41 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 349.6 | yes | Greater Mystic Wand (217287, -0.87 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest]; Scorching Wand (5213, -4.65 DPS) [dungeon] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Black Widow Band; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Gnarled Ash Staff; ranged: Necrotic Wand

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (orc, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 91.9. Weights run: 0.8s. Verify run: 0.8s. 569 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (0.249 ± 0.763), crit=2.900 ± 0.288, hit=5.720 ± 0.629, spell_haste=6.506 ± 0.958, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -1.31 DPS) [dungeon]; Silk Headband (7050, -1.57 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 | yes | Darkspear Warding Pendant (272074, -0.88 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.95 DPS) [vendor]; Necklace of Calisea (1714, -0.98 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.2 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.03 DPS) [crafted]; Battle Healer's Cloak (19528, -0.16 DPS) [rep] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 23.5 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.85 DPS) [crafted]; Crimson Silk Vest (7058, -1.18 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Radiant Silver Bracers (4545, -0.39 DPS) [quest]; Condor Bracers (15864, -0.63 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -0.65 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | Red Mageweave Gloves (10018, -0.72 DPS) [crafted]; Black Mageweave Gloves (10003, -1.10 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.21 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.0 | yes | Star Belt (4329, -0.12 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -0.43 DPS) [rep]; Warsong Sash (16975, -0.52 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.0 | yes | Crimson Silk Pantaloons (7062, -0.75 DPS) [crafted]; Abomination Skin Leggings (23173, -0.78 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.68 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.29 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.09 DPS) [crafted]; Acidic Walkers (9454, -2.23 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 | yes | Reedknot Ring (9622, -0.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.72 DPS) [vendor]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.39 DPS) [vendor]; Reedknot Ring (9622, -0.63 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | sim-verified (90.5 DPS) | yes | Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | Gnomeregan: Leprous Assistant [dungeon] | 311.6 | yes | Illusionary Rod (7713, +0.00 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -4.84 DPS) [quest]; Black Duskwood Staff (937, -7.20 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | 286.1 | yes | Icefury Wand (7514, -0.03 DPS, sim-verified) [quest]; Ragefire Wand (7513, -0.19 DPS) [quest]; Umbral Wand (5216, -1.76 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 569, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (orc, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 124.1. Weights run: 0.7s. Verify run: 0.8s. 723 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.053 ± 1.050), crit=4.846 ± 0.439, hit=8.500 ± 0.859, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 94.5 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.84 DPS) [dungeon]; Red Mageweave Headband (10033, -6.91 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.7 | yes | Scorn's Icy Choker (23169, -0.27 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.31 DPS) [quest]; Darkspear Warding Pendant (272073, -0.67 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 85.3 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -6.77 DPS) [dungeon]; Red Mageweave Shoulders (10029, -7.93 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.5 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.51 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.85 DPS) [vendor] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 91.4 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Acumen Robes (17775, -6.52 DPS) [quest]; Runecloth Robe (13858, -7.93 DPS) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 14.5 | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -0.23 DPS) [quest]; Radiant Silver Bracers (4545, -0.26 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 111.7 | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -11.33 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -11.33 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 82.1 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.31 DPS) [dungeon]; Deathmage Sash (10771, -7.53 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 92.5 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Red Mageweave Pants (10009, -8.36 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.95 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 102.5 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -9.96 DPS) [crafted]; Gilded Sandals (254107, -10.41 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 85.0 | yes | Advisor's Ring (19519, -9.26 DPS) [rep]; Advisor's Ring (19520, -9.65 DPS) [rep]; Ogremind Ring (1993, -9.85 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.3 | yes | Advisor's Ring (19519, +0.00 DPS, sim-verified) [rep]; Advisor's Ring (19520, -0.93 DPS) [rep]; Ogremind Ring (1993, -1.14 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (123.5 DPS) | yes | Uther's Strength (11302, -0.43 DPS, sim-verified) [world_drop]; Guardian Talisman (1490, -9.71 DPS) [quest]; Ankh of Life (1713, -9.71 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (123.5 DPS) | yes | Uther's Strength (11302, -0.56 DPS, sim-verified) [world_drop]; Guardian Talisman (1490, -7.55 DPS) [quest]; Ankh of Life (1713, -7.55 DPS) [dungeon] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (123.5 DPS) | yes | Thrash Blade (17705, -3.45 DPS, sim-verified) [quest]; Soulkeeper (1607, -8.02 DPS) [dungeon]; Spire of Hakkar (10844, -9.60 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 | yes | Noxious Shooter (17745, -0.04 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.73 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Deep Woodlands Cloak; chest: Knight's Dreadweave Vest; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 723, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 60 (orc, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 399.8. Weights run: 0.8s. Verify run: 0.9s. 1087 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.134 ± 1.662), crit=7.715 ± 0.552, hit=17.053 ± 1.684, spell_haste=not significant (-2.222 ± 2.347), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | sim-verified (353.0 DPS) | yes | Bloodvine Goggles (19999, -6.99 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -23.51 DPS) [quest]; Fireleaf Circlet (240056, -55.74 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | sim-verified (344.1 DPS) | yes | Jewel of Kajaro (19601, -0.80 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -23.33 DPS) [quest]; Medallion of the Dawn (22659, -36.83 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 166.0 | yes | Rugged Mantle of the Timbermaw (227808, -4.37 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.67 DPS) [crafted]; Champion's Silk Mantle (23264, -6.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 195.6 | yes | Howler's Furs (272414, -5.41 DPS) [vendor]; Stalwart Cloak (272415, -5.41 DPS) [vendor]; Earthweave Cloak (21187, -11.34 DPS, sim-verified) [quest] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | sim-verified (354.0 DPS) | yes | Bloodvine Vest (19682, -8.00 DPS, sim-verified) [crafted]; Fireleaf Robe (240059, -12.01 DPS) [vendor]; Zandalar Illusionist's Robe (20034, -28.38 DPS) [quest] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | sim-verified (349.7 DPS) | yes | Fireleaf Wristwraps (240044, -2.49 DPS) [vendor]; Rockfury Bracers (21186, -3.68 DPS, sim-verified) [quest]; Frostfire Bindings (22503, -25.30 DPS) [quest] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 273.9 | yes | Gloves of Spell Mastery (14146, -15.63 DPS, sim-verified) [crafted]; Sorcerer's Gloves (22066, -16.30 DPS) [quest]; Sorcerer's Gauntlets (226930, -16.30 DPS) [vendor] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 231.3 | yes | Frostfire Belt (22502, -1.52 DPS, sim-verified) [quest]; Knowledge of the Timbermaw (228190, -2.33 DPS) [vendor]; Fireleaf Waistguard (240045, -13.69 DPS) [vendor] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | sim-verified (365.5 DPS) | yes | Bloodvine Leggings (19683, -6.84 DPS) [crafted]; Sorcerer's Leggings (226933, -8.28 DPS) [quest]; Sentinel's Silk Leggings (237815, -19.49 DPS, sim-verified) [vendor] |
| feet | Enigma Boots (21344) | Enigma Boots [quest] | 215.5 | yes | Marshal's Silk Footwraps (16437, -1.76 DPS) [vendor]; General's Silk Boots (16539, -1.76 DPS) [vendor]; Bloodvine Boots (19684, -7.40 DPS, sim-verified) [crafted] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 278.5 | yes | Band of Earthen Might (21182, -6.54 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -15.08 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -15.75 DPS) [vendor] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | sim-verified (352.2 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -1.21 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -1.89 DPS) [vendor]; Band of Earthen Might (21182, -6.21 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (341.7 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Uther's Strength (11302, -1.73 DPS) [world_drop]; Weakness Analyzer (272438, -1.87 DPS, sim-verified) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (344.1 DPS) | yes | Rune of the Guard Captain (19120, -1.14 DPS, sim-verified) [quest]; Weakness Analyzer (272438, -28.39 DPS) [vendor]; Uther's Strength (11302, -31.85 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (344.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; The Lobotomizer (19324, -9.91 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Stormrager (16997) | The Scarlet Oracle, Demetria [quest] | sim-verified (352.9 DPS) | yes | Brilliant Wand (249385, -0.68 DPS) [crafted]; Wand of Biting Cold (19108, -6.95 DPS, sim-verified) [quest]; Torch of Austen (13004, -7.12 DPS) [world_drop] |

**New at 60:** head: Frostfire Circlet; neck: Onyxia Tooth Pendant; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Frostfire Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; trinket1: Serenity Field; trinket2: Frozen Heart of the Mountain; main_hand: Ironbark Staff; ranged: Stormrager

No-known-source sample (15 of 1087, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

