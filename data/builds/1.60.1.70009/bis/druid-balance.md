# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (night-elf, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 32.2. Weights run: 1.2s. Verify run: 0.6s. 361 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.463 ± 0.112, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | 6.0 | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.38 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.5 | yes | Double-Stitched Woolen Shoulders (4314, -0.65 DPS, sim-verified) [crafted]; Forest Leather Mantle (4709, -0.67 DPS) [world_drop]; Rugged Spaulders (5254, -0.67 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Pearl-clasped Cloak (5542, -0.10 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 | yes | Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted]; Totemic Leather Armor (252435, -0.64 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.9 | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.04 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.50 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.87 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.52 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 | yes | Abomination Skin Leggings (23173, +0.02 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 | yes | Stormrider's Leather Boots (252443, +0.06 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Sludge-Stained Band (286535, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Black Pearl Ring (6332, -0.51 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Sludge-Stained Band (286535, -0.41 DPS, sim-verified) [world]; Black Pearl Ring (6332, -0.43 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 | yes | Twisted Chanter's Staff (890, +0.29 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.42 DPS) [quest]; Channeler's Staff (4437, -0.50 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 361, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals

### Band 30 (night-elf, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.9. Weights run: 1.3s. Verify run: 0.8s. 684 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.971 ± 0.152, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 | yes | Holy Shroud (2721, +0.55 DPS, sim-verified) [world_drop]; Enchanter's Cowl (4322, -0.23 DPS) [crafted]; Silk Headband (7050, -0.40 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.13 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.28 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 | yes | Fairywing Mantle (9536, -0.40 DPS) [quest]; Death Speaker Mantle (6685, -0.43 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 4.7 | yes | Heavy Woolen Cloak (4311, -0.10 DPS) [crafted]; Prelacy Cape (7004, -0.10 DPS) [quest]; Hillman's Cloak (3719, -0.79 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 | yes | Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Guardian Armor (4256, -0.22 DPS, sim-verified) [crafted]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.86 DPS) [quest]; Technician's Bracers (270042, -0.86 DPS) [quest]; Nightsky Wristbands (6407, -1.01 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.7 | yes | Serpent Gloves (5970, -0.23 DPS) [dungeon]; Shilly Mitts (9609, -0.23 DPS) [quest]; Fletcher's Gloves (7348, -0.81 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Belt of Arugal (6392, -0.31 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.46 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon]; Dark Ritual Leggings (270031, -1.17 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -1.76 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.15 DPS) [quest]; Lorekeeper's Ring (20431, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.40 DPS) [dungeon]; Sludge-Stained Band (286535, -0.40 DPS) [world]; Minor Channeling Ring (1449, -1.12 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 21.1 | yes | Glimmering Staff (249392, -2.20 DPS) [crafted]; Twisted Chanter's Staff (890, -2.26 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -4.87 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 9395 Gloves of Old

### Band 40 (night-elf, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 80.3. Weights run: 1.5s. Verify run: 0.9s. 937 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.985 ± 0.232, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Big Voodoo Mask (8201, +0.05 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -0.48 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.98 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -0.38 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 24.3 | yes | Robe of Power (7054, -0.24 DPS) [crafted]; Crimson Silk Vest (7058, -0.68 DPS) [crafted]; Robe of the Magi (1716, -0.83 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 | yes | Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Arcane Runed Bracers (4744, -0.28 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.41 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes | Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted]; Dreamweave Gloves (10019, -1.00 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 | yes | Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.54 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Kodohide Legguards (285338, -0.51 DPS, sim-verified) [world]; Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -1.13 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Minor Channeling Ring (1449, -0.33 DPS) [quest] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 | yes | Reedknot Ring (9622, -0.16 DPS, sim-verified) [quest]; Minor Channeling Ring (1449, -0.20 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.26 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 22.5 | yes | Mograine's Might (7723, -1.44 DPS) [dungeon]; Windweaver Staff (7757, -1.53 DPS) [dungeon]; Illusionary Rod (7713, -1.90 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Forlorn Spirits; trinket1: Rune of Perfection; trinket2: Ankh of Life

No-known-source sample (15 of 937, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 50 (night-elf, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 93.9. Weights run: 1.4s. Verify run: 0.8s. 1196 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.782 ± 0.380, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.87 DPS) [dungeon]; Soothsayer's Headdress (17740, -4.69 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 | yes | Horizon Choker (13085, -0.13 DPS) [world]; Mindburst Medallion (11196, -0.38 DPS, sim-verified) [quest]; Darkspear Warding Pendant (272073, -0.55 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.90 DPS) [vendor]; Blood Guard's Crackling Leather Spaulders (220871, -0.90 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.4 | yes | Big Voodoo Cloak (8216, -0.78 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.93 DPS) [vendor]; Runecloth Cloak (13860, -1.33 DPS, sim-verified) [crafted] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 75.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Knight's Crackling Leather Tunic (220868, -3.12 DPS) [vendor]; Stone Guard's Crackling Leather Tunic (220869, -3.12 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 | yes | Skycaller's Leather Bracers (252542, -0.39 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Shadowskin Gloves (18238), Sergeant Major's Leather Gauntlets (220856), First Sergeant's Leather Gauntlets (220857)) | World drop [world_drop] | 27.2 | yes | Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Sergeant Major's Leather Gauntlets (220856, +0.00 DPS) [vendor]; Fletcher's Gloves (7348, -0.59 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 39.9 | yes | Highlander's Lizardhide Girdle (20103, -0.73 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20115, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 94.4 | yes | Stone Guard's Crackling Leather Leggings (220865, +0.00 DPS, sim-verified) [vendor]; Knight's Leather Pants (220858, -2.22 DPS) [vendor]; Stone Guard's Leather Pants (220859, -2.22 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) (or First Sergeant's Crackling Leather Boots (220863)) | Captain Dirgehammer [vendor] | 66.2 | yes | First Sergeant's Crackling Leather Boots (220863, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.83 DPS) [crafted]; Skycaller's Leather Boots (252471, -4.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.8 | yes | Ring of Forlorn Spirits (2043, -4.56 DPS) [quest]; Reedknot Ring (9622, -4.67 DPS) [quest]; Minor Channeling Ring (1449, -4.73 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Ring of Forlorn Spirits (2043, -0.46 DPS) [quest]; Reedknot Ring (9622, -0.57 DPS) [quest]; Lorekeeper's Ring (19524, -0.79 DPS, sim-verified) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, +0.00 DPS) [world] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 43.0 | yes | Uther's Strength (11302, -0.20 DPS, sim-verified) [world]; Thunderbrew's Boot Flask (744, -4.93 DPS) [quest]; Tidal Charm (1404, -4.93 DPS) [vendor] |
| main_hand | - | - |  |  |  |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Spritecaster Cape; chest: Knight's Leather Armor; wrist: Runic Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Cloth Girdle; legs: Knight's Crackling Leather Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Frozen Heart of the Mountain; main_hand: Blight

No-known-source sample (15 of 1196, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (night-elf, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 149.8. Weights run: 1.4s. Verify run: 0.9s. 1913 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.891 ± 0.496, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bloodvine Goggles (19999) | Engineering [crafted] | 153.0 | yes | Mask of the Unforgiven (13404, -3.13 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Leather Headband (220850, -6.51 DPS) [vendor]; Blood Guard's Leather Headband (220851, -6.51 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 94.1 | yes | Beads of Ogre Might (22150, -3.89 DPS) [quest]; Medallion of the Dawn (22659, -6.51 DPS) [quest]; Charm of the Shifting Sands (21504, -6.82 DPS) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 94.1 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Feralheart Spaulders (226778, -0.97 DPS) [quest]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.88 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 58.9 | yes | Chromatic Cloak (18509, -0.78 DPS, sim-verified) [crafted]; Hide of the Wild (18510, -4.28 DPS) [crafted]; Spritecaster Cape (11623, -4.55 DPS) [dungeon] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 152.9 | yes | Knight-Captain's Dragonhide Chestpiece (227176, -3.28 DPS, sim-verified) [pvp]; Legionnaire's Dragonhide Chestpiece (227179, -5.61 DPS) [pvp]; Knight's Leather Armor (220854, -6.50 DPS) [vendor] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 85.9 | yes | Primal Batskin Bracers (19687, -3.77 DPS, sim-verified) [crafted]; Dryad's Wrist Bindings (19595, -6.52 DPS) [rep]; Dryad's Wrist Bindings (19596, -6.88 DPS) [rep] |
| hands | Primal Batskin Gloves (19686) | Leatherworking [crafted] | 117.8 | yes | Stormshroud Gloves (21278, -0.60 DPS, sim-verified) [crafted]; Dreadmist Wraps (16705, -5.90 DPS) [dungeon]; Blood Guard's Dragonhide Grips (227180, -6.51 DPS) [pvp] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 65.1 | yes | Highlander's Cloth Girdle (20047, -1.35 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.97 DPS) [rep]; Belt of Preserved Heads (20216, -4.32 DPS, sim-verified) [quest] |
| legs | Knight's Crackling Leather Leggings (220864) (or Stone Guard's Crackling Leather Leggings (220865)) | Captain Dirgehammer [vendor] | 112.3 | yes | Stone Guard's Crackling Leather Leggings (220865, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Dragonhide Leggings (227177, -1.19 DPS) [pvp]; Knight-Captain's Dragonhide Leggings (227178, -1.19 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 87.8 | yes | First Sergeant's Crackling Leather Boots (220863, -1.30 DPS) [vendor]; Sergeant Major's Crackling Leather Boots (220862, -1.90 DPS, sim-verified) [vendor]; Blood Guard's Dragonhide Treads (227181, -5.55 DPS) [pvp] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 99.6 | yes | Band of Earthen Might (21182, -0.61 DPS) [quest]; Mindtear Band (20632, -4.28 DPS) [world]; Ritssyn's Ring of Chaos (21836, -4.36 DPS) [world] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 94.1 | yes | Band of Earthen Might (21182, -0.13 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.67 DPS) [world]; Ritssyn's Ring of Chaos (21836, -3.75 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.53 DPS, sim-verified) [world] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 165.6 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Atiesh, Greatstaff of the Guardian (22631, -3.13 DPS) [quest]; Enchanted Battlehammer (12776, -5.29 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Bloodvine Goggles; neck: Onyxia Tooth Pendant; back: Earthweave Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Primal Batskin Gloves; waist: Belt of the Archmage; feet: Bloodvine Boots; finger1: Ring of the Fallen God; finger2: Don Julio's Band; trinket2: Thunderbrew's Boot Flask; main_hand: High Warlord's War Staff

No-known-source sample (15 of 1913, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (tauren, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 31.4. Weights run: 1.2s. Verify run: 0.6s. 359 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.463 ± 0.112, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | 6.0 | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.36 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.5 | yes | Double-Stitched Woolen Shoulders (4314, -0.53 DPS, sim-verified) [crafted]; Forest Leather Mantle (4709, -0.67 DPS) [world_drop]; Rugged Spaulders (5254, -0.67 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Pearl-clasped Cloak (5542, -0.04 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 | yes | Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted]; Totemic Leather Armor (252435, -0.47 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 1.9 | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.37 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.84 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.50 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 | yes | Abomination Skin Leggings (23173, +0.13 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 | yes | Stormrider's Leather Boots (252443, +0.16 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Black Pearl Ring (6332, -0.43 DPS) [world]; The 1 Ring (8350, -0.47 DPS) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, +0.26 DPS, sim-verified) [dungeon]; Black Pearl Ring (6332, -0.23 DPS) [world]; The 1 Ring (8350, -0.27 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 | yes | Gnarled Necromancer's Staff (251534, -0.42 DPS) [quest]; Channeler's Staff (4437, -0.50 DPS) [world]; Twisted Chanter's Staff (890, -0.64 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 359, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals

### Band 30 (tauren, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 58.0. Weights run: 1.3s. Verify run: 0.8s. 685 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.971 ± 0.152, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 | yes | Holy Shroud (2721, +0.39 DPS, sim-verified) [world_drop]; Enchanter's Cowl (4322, -0.23 DPS) [crafted]; Silk Headband (7050, -0.40 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.18 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.28 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 | yes | Death Speaker Mantle (6685, -0.04 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.42 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Battle Healer's Cloak (19529, -0.13 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 | yes | Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Guardian Armor (4256, -0.85 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.83 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.86 DPS) [quest]; Technician's Bracers (270042, -0.86 DPS) [quest] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 9.9 | yes | Jutebraid Gloves (10654, -0.23 DPS) [quest]; Serpent Gloves (5970, -0.38 DPS) [dungeon]; Fletcher's Gloves (7348, -1.74 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Warsong Sash (16975, -0.21 DPS) [quest]; Defiler's Cloth Girdle (20164, -1.14 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.73 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -2.40 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; Sludge-Stained Band (286535, -0.54 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.40 DPS) [world]; Black Widow Band (6199, -0.40 DPS) [world]; Electrocutioner Lagnut (9447, -0.91 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 4.7 | yes | Twisted Chanter's Staff (890, -0.06 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.06 DPS) [quest]; Rhahk'Zor's Hammer (5187, -2.47 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff

No-known-source sample (15 of 685, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 9395 Gloves of Old

### Band 40 (tauren, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 82.4. Weights run: 1.5s. Verify run: 0.9s. 937 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.985 ± 0.232, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -0.38 DPS) [world]; Big Voodoo Mask (8201, -0.61 DPS, sim-verified) [crafted]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -0.48 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.98 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Green Silken Shoulders (7057, -0.20 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.27 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 | yes | Dreamweave Vest (10021, +0.41 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.48 DPS) [crafted]; Crimson Silk Vest (7058, -0.92 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 | yes | Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.27 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.41 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes | Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted]; Dreamweave Gloves (10019, -1.30 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 | yes | Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.93 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Kodohide Legguards (285338, -0.76 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -1.28 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Advisor's Ring (20426, -0.51 DPS) [rep]; Reedknot Ring (9622, -2.06 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Ogremind Ring (1993, -0.14 DPS) [world_drop]; Voodoo Band (1996, -0.14 DPS) [world]; Reedknot Ring (9622, -0.81 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Mograine's Might (7723) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 11.2 | yes | Windweaver Staff (7757, -0.09 DPS) [dungeon]; Rhahk'Zor's Hammer (5187, -0.41 DPS) [dungeon]; Illusionary Rod (7713, -4.12 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Mograine's Might

No-known-source sample (15 of 937, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 50 (tauren, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 94.0. Weights run: 1.4s. Verify run: 0.9s. 1196 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.782 ± 0.380, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.87 DPS) [dungeon]; Soothsayer's Headdress (17740, -4.69 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 | yes | Mindburst Medallion (11196, +0.63 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.13 DPS) [world]; Darkspear Warding Pendant (272073, -0.55 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.90 DPS) [vendor]; Blood Guard's Crackling Leather Spaulders (220871, -0.90 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.6 | yes | Spritecaster Cape (11623, +0.35 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.43 DPS) [crafted]; Big Voodoo Cloak (8216, -0.80 DPS) [crafted] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 75.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Knight's Crackling Leather Tunic (220868, -3.12 DPS) [vendor]; Stone Guard's Crackling Leather Tunic (220869, -3.12 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 | yes | Skycaller's Leather Bracers (252542, +0.03 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Shadowskin Gloves (18238), Sergeant Major's Leather Gauntlets (220856), First Sergeant's Leather Gauntlets (220857)) | World drop [world_drop] | 27.2 | yes | Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Sergeant Major's Leather Gauntlets (220856, +0.00 DPS) [vendor]; Fletcher's Gloves (7348, -0.59 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 39.9 | yes | Defiler's Lizardhide Girdle (20174, -0.65 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20193, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | Lady Palanseer [vendor] | 94.4 | yes | Knight's Crackling Leather Leggings (220864, +0.00 DPS, sim-verified) [vendor]; Knight's Leather Pants (220858, -2.22 DPS) [vendor]; Stone Guard's Leather Pants (220859, -2.22 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) (or First Sergeant's Crackling Leather Boots (220863)) | Captain Dirgehammer [vendor] | 66.2 | yes | First Sergeant's Crackling Leather Boots (220863, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.83 DPS) [crafted]; Skycaller's Leather Boots (252471, -4.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.8 | yes | Reedknot Ring (9622, -4.67 DPS) [quest]; Sea Giant's Toe Ring (274746, -4.79 DPS) [vendor]; Ogremind Ring (1993, -4.88 DPS) [world_drop] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Advisor's Ring (19520, -0.50 DPS, sim-verified) [rep]; Reedknot Ring (9622, -0.57 DPS) [quest]; Advisor's Ring (19521, -0.57 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, -0.13 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 33.5 | yes | Frozen Heart of the Mountain (249469, -1.03 DPS, sim-verified) [crafted]; Uther's Strength (11302, -3.15 DPS) [world]; Tidal Charm (1404, -3.83 DPS) [vendor] |
| main_hand | - | - |  |  |  |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Deep Woodlands Cloak; chest: Knight's Leather Armor; wrist: Runic Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Cloth Girdle; legs: Stone Guard's Crackling Leather Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Rune of the Guard Captain; main_hand: Blight

No-known-source sample (15 of 1196, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (tauren, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 147.0. Weights run: 1.4s. Verify run: 0.8s. 1912 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.891 ± 0.496, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) (or Bloodvine Goggles (19999)) | Stratholme: The Unforgiven [dungeon] | 153.0 | yes | Bloodvine Goggles (19999, +1.20 DPS, sim-verified) [crafted]; Knight-Lieutenant's Leather Headband (220850, -6.51 DPS) [vendor]; Blood Guard's Leather Headband (220851, -6.51 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 94.1 | yes | Beads of Ogre Might (22150, -3.89 DPS) [quest]; Medallion of the Dawn (22659, -6.51 DPS) [quest]; Charm of the Shifting Sands (21504, -6.82 DPS) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 94.1 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Feralheart Spaulders (226778, -0.97 DPS) [quest]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.88 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 58.9 | yes | Chromatic Cloak (18509, -1.30 DPS, sim-verified) [crafted]; Hide of the Wild (18510, -4.28 DPS) [crafted]; Spritecaster Cape (11623, -4.55 DPS) [dungeon] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 152.9 | yes | Knight-Captain's Dragonhide Chestpiece (227176, -3.48 DPS, sim-verified) [pvp]; Legionnaire's Dragonhide Chestpiece (227179, -5.61 DPS) [pvp]; Knight's Leather Armor (220854, -6.50 DPS) [vendor] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 85.9 | yes | Primal Batskin Bracers (19687, -3.75 DPS, sim-verified) [crafted]; Dryad's Wrist Bindings (19595, -6.52 DPS) [rep]; Dryad's Wrist Bindings (19596, -6.88 DPS) [rep] |
| hands | Primal Batskin Gloves (19686) | Leatherworking [crafted] | 117.8 | yes | Stormshroud Gloves (21278, -1.21 DPS, sim-verified) [crafted]; Dreadmist Wraps (16705, -5.90 DPS) [dungeon]; Blood Guard's Dragonhide Grips (227180, -6.51 DPS) [pvp] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 65.1 | yes | Defiler's Cloth Girdle (20163, -1.35 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.97 DPS) [rep]; Belt of Preserved Heads (20216, -5.20 DPS, sim-verified) [quest] |
| legs | Knight's Crackling Leather Leggings (220864) (or Stone Guard's Crackling Leather Leggings (220865)) | Captain Dirgehammer [vendor] | 112.3 | yes | Stone Guard's Crackling Leather Leggings (220865, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Dragonhide Leggings (227177, -1.19 DPS) [pvp]; Knight-Captain's Dragonhide Leggings (227178, -1.19 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 87.8 | yes | First Sergeant's Crackling Leather Boots (220863, -1.30 DPS) [vendor]; Sergeant Major's Crackling Leather Boots (220862, -3.03 DPS, sim-verified) [vendor]; Blood Guard's Dragonhide Treads (227181, -5.55 DPS) [pvp] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 99.6 | yes | Band of Earthen Might (21182, -0.61 DPS) [quest]; Mindtear Band (20632, -4.28 DPS) [world]; Ritssyn's Ring of Chaos (21836, -4.36 DPS) [world] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 94.1 | yes | Band of Earthen Might (21182, -0.13 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.67 DPS) [world]; Ritssyn's Ring of Chaos (21836, -3.75 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.98 DPS, sim-verified) [world] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 41.2 | yes | Frozen Heart of the Mountain (249469, -1.16 DPS, sim-verified) [crafted]; Uther's Strength (11302, -3.90 DPS) [world]; Tidal Charm (1404, -4.56 DPS) [vendor] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 165.6 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Atiesh, Greatstaff of the Guardian (22631, -3.13 DPS) [quest]; Enchanted Battlehammer (12776, -5.29 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Mask of the Unforgiven; neck: Onyxia Tooth Pendant; back: Earthweave Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Primal Batskin Gloves; waist: Belt of the Archmage; legs: Knight's Crackling Leather Leggings; feet: Bloodvine Boots; finger1: Ring of the Fallen God; finger2: Don Julio's Band; main_hand: High Warlord's War Staff

No-known-source sample (15 of 1912, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

