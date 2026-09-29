# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.5. Weights run: 1.1s. Verify run: 1.2s. 361 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.009 ± 0.003, crit=0.390 ± 0.062, hit=not significant (0.899 ± 0.286), melee_haste=not significant (0.636 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.38 DPS) [crafted]; Lucky Fishing Hat (19972, -0.38 DPS) [quest]; Flying Tiger Goggles (4368, -0.54 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 | yes | Tarnished Locket (279870, -0.41 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Forest Leather Mantle (4709, -0.24 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.34 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.02 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.1 | yes | Brawler's Leather Armor (252490, +0.05 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.19 DPS) [crafted]; Dark Leather Tunic (2317, -0.24 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.0 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.07 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.26 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.03 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.06 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.33 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world]; Minor Channeling Ring (1449, -0.29 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 4.0 | yes | The 1 Ring (8350, -0.14 DPS) [world]; Minor Channeling Ring (1449, -0.19 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.40 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Barrens Basher (274744, -1.21 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.85 DPS) [quest]; Pulsating Hydra Heart (5183, -10.85 DPS) [world]; Tear of Grief (5611, -10.85 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Light Bow (4576, -0.09 DPS) [world_drop]; Owlsight Rifle (15205, -0.09 DPS) [quest]; Deadly Blunderbuss (4369, -0.13 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 361, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (night-elf, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.7. Weights run: 1.0s. Verify run: 1.6s. 694 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.033 ± 0.019, crit=0.465 ± 0.062, hit=2.484 ± 0.367, melee_haste=not significant (0.623 ± 0.730)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.3 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Holy Shroud (2721, -0.50 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.36 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.38 DPS) [rep]; Pendant of Myzrael (4614, -0.68 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.4 | yes | Dark Leather Shoulders (4252, -0.20 DPS) [crafted]; Insignia Mantle (4721, -0.20 DPS) [world_drop]; Mantle of Thieves (2264, -0.39 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Fenrus' Hide (6340, -0.18 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.18 DPS) [dungeon]; Cloak of Night (4447, -0.24 DPS, sim-verified) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.08 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.23 DPS) [quest]; Green Leather Armor (4255, -0.38 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.11 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.18 DPS) [world_drop]; Madwolf Bracers (897, -0.23 DPS) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Fletcher's Gloves (7348, -0.46 DPS) [crafted]; Wolfclaw Gloves (1978, -0.48 DPS) [dungeon]; Pilferer's Gloves (7358, -0.49 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.71 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -0.80 DPS, sim-verified) [crafted]; Insignia Leggings (4054, -0.81 DPS) [world_drop]; Leggings of the Fang (10410, -0.81 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 8.3 | yes | Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.05 DPS) [quest]; Insignia Boots (4055, -0.19 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.3 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.15 DPS) [dungeon]; Protector's Band (19517, -0.15 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.11 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +2.07 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.64 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.78 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -15.45 DPS) [quest]; Rod of Molten Fire (2565, -15.45 DPS) [world_drop]; Eye of Paleth (2943, -15.45 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Precision Bow (217315, -0.08 DPS) [quest]; Silver Star (3463, -0.19 DPS) [quest]; Moonsight Rifle (4383, -0.41 DPS, sim-verified) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 694, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 88.9. Weights run: 1.1s. Verify run: 1.5s. 962 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.962 ± 0.146, hit=4.384 ± 0.987, melee_haste=not significant (-0.246 ± 1.799)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.1 | yes | White Bandit Mask (10008, +0.77 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.05 DPS) [world]; Brawler's Leather Helm (252512, -0.10 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.19 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.30 DPS) [rep]; Sentinel's Medallion (20444, -0.40 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.16 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.10 DPS) [crafted]; Yeti Fur Cloak (2805, -0.20 DPS) [quest] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Nightscape Tunic (8175, +0.50 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.09 DPS) [crafted]; Hawkeye's Tunic (14592, -0.19 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.60 DPS) [world_drop]; Dusky Bracers (7378, -0.60 DPS) [crafted]; Cultist's Armguards (270032, -0.81 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.5 | yes | Fletcher's Gloves (7348, -1.00 DPS) [crafted]; Shadowskin Gloves (18238, -1.00 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.28 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.49 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.41 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.39 DPS) [quest]; Hawkeye's Breeches (14595, -0.60 DPS) [world] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.1 | yes | Imperial Leather Boots (6431, +0.17 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.10 DPS) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Protector's Band (19515, -0.20 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 10.1 | yes | Ironspine's Eye (7686, +0.04 DPS, sim-verified) [dungeon]; Protector's Band (19515, -0.10 DPS) [rep]; Disengagement Ring (276202, -0.10 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -1.37 DPS) [vendor]; Sword of Serenity (6829, -2.01 DPS) [quest]; Hand of Righteousness (7721, -2.14 DPS) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Grayson's Torch (1172, -22.18 DPS) [quest]; Rod of Molten Fire (2565, -22.18 DPS) [world_drop]; Eye of Paleth (2943, -22.18 DPS) [quest] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 13.2 | yes | Precision Bow (217315, +0.00 DPS, sim-verified) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.21 DPS) [vendor]; Master Hunter's Bow (17686, -0.36 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Moonsight Rifle

No-known-source sample (15 of 962, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 111.1. Weights run: 1.1s. Verify run: 1.4s. 1220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.239 ± 0.059, crit=6.333 ± 0.367, hit=not significant (5.215 ± 1.384), melee_haste=not significant (-1.529 ± 2.421)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 156.8 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.69 DPS) [dungeon]; Helm of Fire (8348, -7.35 DPS) [crafted] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 14.9 | yes | Sentinel's Medallion (19540, -0.07 DPS) [rep]; Ghostshard Talisman (7731, -0.25 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19541, -0.27 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 148.8 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -6.67 DPS) [vendor]; Forest Tracker Epaulets (2278, -7.32 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 13.6 | yes | Nightscape Cloak (8195, -0.10 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.13 DPS) [dungeon]; Wolfmaster Cape (6314, -0.20 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 158.8 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -7.05 DPS) [quest]; Warbear Harness (15064, -7.39 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.43 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.35 DPS) [crafted]; Pridelord Bands (14672, -0.41 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 108.7 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.48 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.08 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 108.7 | yes | Highlander's Lizardhide Girdle (20103, -1.08 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.58 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -4.26 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 158.8 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.66 DPS, sim-verified) [crafted]; Basilisk Hide Pants (1718, -7.19 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.8 | yes | Sandstalker Ankleguards (12470, +0.14 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.37 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.37 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 72.1 | yes | Insurgent's Band (272065, -3.09 DPS) [vendor]; Ring of the Underwood (2951, -3.23 DPS) [world_drop]; Insurgent's Band (272066, -3.26 DPS) [vendor] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 17.3 | yes | Ring of the Underwood (2951, -0.27 DPS) [world_drop]; Insurgent's Band (272066, -0.29 DPS) [vendor]; Insurgent's Band (272065, -0.43 DPS, sim-verified) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 46.9 | yes | Thunderbrew's Boot Flask (744, -2.54 DPS) [quest]; Tidal Charm (1404, -2.54 DPS) [vendor]; Guardian Talisman (1490, -2.54 DPS) [quest] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 581.9 | yes | Might of Hakkar (10838, -2.93 DPS) [world]; Thorium Cestus (250614, -3.00 DPS) [crafted]; Julie's Dagger (6660, -3.81 DPS) [world_drop] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Claw of Celebras (17738, -3.85 DPS) [dungeon]; Thermotastic Egg Timer (9644, -29.74 DPS) [quest]; Grayson's Torch (1172, -29.94 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 17.3 | yes | Moonsight Rifle (4383, -0.09 DPS) [crafted]; Precision Bow (217315, -0.09 DPS) [quest]; Houndmaster's Bow (11628, -0.29 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Sentinel's Medallion; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1220, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 370.7. Weights run: 1.1s. Verify run: 1.8s. 1717 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 316.1 | yes | Bloodvine Lens (19998, -4.60 DPS) [crafted]; Mask of the Unforgiven (13404, -6.05 DPS) [dungeon]; Ragefury Eyepatch (11735, -12.05 DPS, sim-verified) [dungeon] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas: Kel'Thuzad [raid] | 256.2 | yes | Gem of Trapped Innocents (23057, -1.39 DPS) [raid]; Onyxia Tooth Pendant (18404, -4.68 DPS) [quest]; Barbed Choker (21664, -5.20 DPS) [raid] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 189.8 | yes | Lieutenant Commander's Leather Shoulders (23313, -0.47 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.47 DPS) [pvp]; Champion's Leather Shoulders (23258, -5.15 DPS, sim-verified) [vendor] |
| back | Cloak of Veiled Shadows (21406) | Cloak of Veiled Shadows [quest] | 69.1 | yes | Earthweave Cloak (21187, -0.22 DPS) [quest]; Cloak of the Honor Guard (20073, -1.51 DPS) [rep]; Chromatic Cloak (18509, -6.67 DPS, sim-verified) [crafted] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 354.2 | yes | Stormshroud Armor (15056, -6.64 DPS) [crafted]; Deathdealer's Vest (21364, -7.68 DPS) [quest]; Zandalar Madcap's Tunic (19834, -8.06 DPS, sim-verified) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 151.4 | yes | Qiraji Execution Bracers (21602, +2.41 DPS, sim-verified) [raid]; Primal Batskin Bracers (19687, -4.71 DPS) [crafted]; Rockfury Bracers (21186, -5.75 DPS) [quest] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 225.1 | yes | Devilsaur Gauntlets (15063, -4.39 DPS) [crafted]; Marshal's Leather Handgrips (16454, -4.39 DPS) [vendor]; Stormshroud Gloves (21278, -8.68 DPS, sim-verified) [crafted] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj: C'Thun [raid] | 223.1 | yes | Bonescythe Waistguard (22482, -3.99 DPS) [quest]; Highlander's Leather Girdle (20115, -4.71 DPS) [rep]; Highlander's Leather Girdle (20045, -7.27 DPS, sim-verified) [rep] |
| legs | Marshal's Leather Leggings (231548) (or General's Leather Legguards (231554)) | Rank 16 [pvp] | 240.8 | yes | General's Leather Legguards (231554, +0.00 DPS, sim-verified) [pvp]; Marshal's Leather Leggings (16456, -0.00 DPS) [vendor]; General's Leather Legguards (16564, -0.00 DPS) [vendor] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 223.1 | yes | Deathdealer's Boots (21359, -3.33 DPS, sim-verified) [quest]; Bloodvine Boots (19684, -9.59 DPS) [crafted]; Shadowcraft Boots (16711, -9.67 DPS) [dungeon] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas: Loatheb [raid] | 211.1 | yes | Band of Earthen Might (21182, -2.78 DPS) [quest]; Ring of the Qiraji Fury (21677, -3.00 DPS) [raid]; Quick Strike Ring (18821, -3.53 DPS) [raid] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 175.1 | yes | Ring of the Qiraji Fury (21677, -1.07 DPS) [raid]; Band of Earthen Might (21182, -1.32 DPS, sim-verified) [quest]; Quick Strike Ring (18821, -1.60 DPS) [raid] |
| trinket1 | Slayer's Crest (23041) | Naxxramas: Sapphiron [raid] | 64.0 | yes | Drake Fang Talisman (19406, +4.28 DPS) [raid]; Neltharion's Tear (19379, +1.28 DPS) [raid]; Eye of Diminution (23001, -9.44 DPS, sim-verified) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas: Maexxna [raid] | 159.1 | yes | Drake Fang Talisman (19406, -0.81 DPS) [raid]; Neltharion's Tear (19379, -3.81 DPS) [raid]; Eye of Diminution (23001, -6.87 DPS, sim-verified) [raid] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Rise, Thunderfury! [quest] | 762.3 | yes | High Warlord's Quickblade (234553, +13.28 DPS) [pvp]; Grand Marshal's Swiftblade (234579, +13.28 DPS) [pvp]; Blessed Qiraji Pugio (21244, -53.83 DPS, sim-verified) [quest] |
| off_hand | The Hungering Cold (23577) | Naxxramas: Kel'Thuzad [raid] | 1022.0 | yes | High Warlord's Left Claw (234558, -0.79 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.79 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -2.47 DPS) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj: Ouro [raid] | 133.1 | yes | The Purifier (22656, -0.96 DPS) [quest]; Fahrad's Reloading Repeater (22347, -4.47 DPS) [quest]; Core Marksman Rifle (18282, -4.77 DPS) [crafted] |

**New at 60:** head: Bonescythe Helmet; neck: Stormrage's Talisman of Seething; shoulder: Bonescythe Pauldrons; back: Cloak of Veiled Shadows; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; legs: Marshal's Leather Leggings; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Don Julio's Band; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: The Hungering Cold; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1717, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.0. Weights run: 1.1s. Verify run: 1.3s. 360 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.009 ± 0.003, crit=0.390 ± 0.062, hit=not significant (0.899 ± 0.286), melee_haste=not significant (0.636 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.38 DPS) [crafted]; Lucky Fishing Hat (19972, -0.38 DPS) [quest]; Flying Tiger Goggles (4368, -0.51 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 | yes | Tarnished Locket (279870, -0.39 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Forest Leather Mantle (4709, -0.24 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.32 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Heckler's Hide (286536, -0.10 DPS) [world]; Trapper's Leather Armor (252491, -0.32 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.0 | yes | Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor]; Spare Part Bindings (279875, -0.10 DPS) [quest] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.26 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.03 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.06 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.32 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 4.0 | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The 1 Ring (8350, -0.14 DPS) [world]; Bounty Hunter's Ring (5351, -0.32 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Wingblade (6504, -1.15 DPS) [quest] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.85 DPS) [quest]; Nightglow Concoction (3451, -10.85 DPS) [quest]; Pulsating Hydra Heart (5183, -10.85 DPS) [world] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Light Bow (4576, -0.09 DPS) [world_drop]; Privateer Musket (5309, -0.09 DPS) [quest]; Deadly Blunderbuss (4369, -0.13 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 360, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers

### Band 30 (troll, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.3. Weights run: 1.0s. Verify run: 1.6s. 695 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.033 ± 0.019, crit=0.465 ± 0.062, hit=2.484 ± 0.367, melee_haste=not significant (0.623 ± 0.730)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.3 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Holy Shroud (2721, -0.50 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.38 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.38 DPS) [rep]; Pendant of Myzrael (4614, -0.68 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.4 | yes | Dark Leather Shoulders (4252, -0.20 DPS) [crafted]; Insignia Mantle (4721, -0.20 DPS) [world_drop]; Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Cloak of Night (4447, -0.18 DPS) [world]; Fenrus' Hide (6340, -0.18 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.5 | yes | Panther Armor (6670, -0.19 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.30 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.30 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.12 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.18 DPS) [world_drop]; Madwolf Bracers (897, -0.23 DPS) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Braced Handguards (6784, -0.43 DPS) [quest]; Fletcher's Gloves (7348, -0.46 DPS) [crafted]; Pilferer's Gloves (7358, -0.50 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Deftkin Belt (16659, -0.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Insignia Leggings (4054, -0.81 DPS) [world_drop]; Leggings of the Fang (10410, -0.81 DPS) [dungeon]; Dusky Leather Leggings (7373, -0.82 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 8.3 | yes | Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Insignia Boots (4055, -0.19 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.3 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.15 DPS) [dungeon]; Legionnaire's Band (19513, -0.15 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.12 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +1.87 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.64 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.78 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -15.45 DPS) [quest]; Rod of Molten Fire (2565, -15.45 DPS) [world_drop]; Nightglow Concoction (3451, -15.45 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Precision Bow (217315, -0.08 DPS) [quest]; Silver Star (3463, -0.19 DPS) [quest]; Moonsight Rifle (4383, -0.35 DPS, sim-verified) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 695, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow

### Band 40 (troll, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 87.9. Weights run: 1.1s. Verify run: 1.5s. 962 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.962 ± 0.146, hit=4.384 ± 0.987, melee_haste=not significant (-0.246 ± 1.799)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.1 | yes | White Bandit Mask (10008, +0.78 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.05 DPS) [world]; Spirit Hunter Headdress (6720, -0.10 DPS) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.17 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.30 DPS) [rep]; Scout's Medallion (20442, -0.40 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.10 DPS) [world_drop]; Parachute Cloak (10518, -0.10 DPS) [crafted] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 15.2 | yes | Dusky Leather Armor (7374, -0.10 DPS, sim-verified) [crafted]; Hawkeye's Tunic (14592, -0.15 DPS) [world]; Panther Armor (6670, -0.30 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.60 DPS) [world_drop]; Dusky Bracers (7378, -0.60 DPS) [crafted]; Cultist's Armguards (270032, -0.80 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.5 | yes | Fletcher's Gloves (7348, -1.00 DPS) [crafted]; Shadowskin Gloves (18238, -1.00 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.43 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.48 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.38 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.39 DPS) [quest]; Hawkeye's Breeches (14595, -0.60 DPS) [world] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.1 | yes | Imperial Leather Boots (6431, +0.17 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.10 DPS) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Legionnaire's Band (19512, -0.20 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 10.1 | yes | Ironspine's Eye (7686, +0.06 DPS, sim-verified) [dungeon]; Legionnaire's Band (19512, -0.10 DPS) [rep]; Disengagement Ring (276202, -0.10 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -1.37 DPS) [vendor]; Hand of Righteousness (7721, -2.14 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272093, -2.16 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Grayson's Torch (1172, -22.18 DPS) [quest]; Rod of Molten Fire (2565, -22.18 DPS) [world_drop]; Nightglow Concoction (3451, -22.18 DPS) [quest] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 13.2 | yes | Precision Bow (217315, +0.00 DPS, sim-verified) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.21 DPS) [vendor]; Master Hunter's Bow (17686, -0.36 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Moonsight Rifle

No-known-source sample (15 of 962, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (troll, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 109.9. Weights run: 1.1s. Verify run: 1.4s. 1220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.239 ± 0.059, crit=6.333 ± 0.367, hit=not significant (5.215 ± 1.384), melee_haste=not significant (-1.529 ± 2.421)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 156.8 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.69 DPS) [dungeon]; Helm of Fire (8348, -7.35 DPS) [crafted] |
| neck | Scout's Medallion (19535) | Warsong Outriders [rep] | 14.9 | yes | Scout's Medallion (19536, -0.07 DPS) [rep]; Woven Ivy Necklace (19159, -0.20 DPS) [quest]; Ghostshard Talisman (7731, -0.29 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 148.8 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -6.67 DPS) [vendor]; Forest Tracker Epaulets (2278, -7.32 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 13.6 | yes | Nightscape Cloak (8195, -0.13 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.13 DPS) [dungeon]; Wolfmaster Cape (6314, -0.20 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 158.8 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -7.05 DPS) [quest]; Warbear Harness (15064, -7.39 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.40 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.35 DPS) [crafted]; Pridelord Bands (14672, -0.41 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 108.7 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.47 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.08 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 108.7 | yes | Defiler's Lizardhide Girdle (20174, -1.08 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.56 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -4.26 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 158.8 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.68 DPS, sim-verified) [crafted]; Basilisk Hide Pants (1718, -7.19 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.8 | yes | Sandstalker Ankleguards (12470, +0.12 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.37 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.37 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 72.1 | yes | Masons Fraternity Ring (9533, -2.97 DPS) [quest]; Insurgent's Band (272065, -3.09 DPS) [vendor]; Ring of the Underwood (2951, -3.23 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Masons Fraternity Ring (9533, -0.34 DPS, sim-verified) [quest]; Insurgent's Band (272065, -0.49 DPS) [vendor]; Ring of the Underwood (2951, -0.63 DPS) [world_drop] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, +4.25 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 46.9 | yes | Rune of the Guard Captain (19120, +1.71 DPS) [quest]; Tidal Charm (1404, -2.54 DPS) [vendor]; Guardian Talisman (1490, -2.54 DPS) [quest] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 581.9 | yes | Might of Hakkar (10838, -2.93 DPS) [world]; Thorium Cestus (250614, -3.00 DPS) [crafted]; Julie's Dagger (6660, -3.81 DPS) [world_drop] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Claw of Celebras (17738, -3.85 DPS) [dungeon]; White Bone Shredder (11863, -5.93 DPS) [quest]; Thermotastic Egg Timer (9644, -29.74 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 17.3 | yes | Moonsight Rifle (4383, -0.09 DPS) [crafted]; Precision Bow (217315, -0.09 DPS) [quest]; Houndmaster's Bow (11628, -0.29 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Scout's Medallion; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Defiler's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Smoking Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1220, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 360.1. Weights run: 1.1s. Verify run: 1.8s. 1716 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 316.1 | yes | Bloodvine Lens (19998, -4.60 DPS) [crafted]; Mask of the Unforgiven (13404, -6.05 DPS) [dungeon]; Ragefury Eyepatch (11735, -14.20 DPS, sim-verified) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 18.2 | yes | Gem of Trapped Innocents (23057, +11.35 DPS) [raid]; Onyxia Tooth Pendant (18404, +8.07 DPS) [quest]; Stormrage's Talisman of Seething (23053, +0.20 DPS, sim-verified) [raid] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 189.8 | yes | Lieutenant Commander's Leather Shoulders (23313, -0.47 DPS) [vendor]; Champion's Leather Shoulders (227056, -0.47 DPS) [pvp]; Champion's Leather Shoulders (23258, -5.83 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 115.1 | yes | Cloak of Veiled Shadows (21406, +3.55 DPS, sim-verified) [quest]; Earthweave Cloak (21187, -2.69 DPS) [quest]; Deathguard's Cloak (20068, -3.97 DPS) [rep] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 354.2 | yes | Stormshroud Armor (15056, -6.64 DPS) [crafted]; Deathdealer's Vest (21364, -7.68 DPS) [quest]; Zandalar Madcap's Tunic (19834, -9.43 DPS, sim-verified) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 151.4 | yes | Qiraji Execution Bracers (21602, -1.48 DPS, sim-verified) [raid]; Primal Batskin Bracers (19687, -4.71 DPS) [crafted]; Rockfury Bracers (21186, -5.75 DPS) [quest] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 225.1 | yes | Devilsaur Gauntlets (15063, -4.39 DPS) [crafted]; Marshal's Leather Handgrips (16454, -4.39 DPS) [vendor]; Stormshroud Gloves (21278, -8.49 DPS, sim-verified) [crafted] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj: C'Thun [raid] | 223.1 | yes | Bonescythe Waistguard (22482, -3.99 DPS) [quest]; Defiler's Leather Girdle (20193, -4.71 DPS) [rep]; Defiler's Leather Girdle (20190, -8.67 DPS, sim-verified) [rep] |
| legs | Marshal's Leather Leggings (231548) (or General's Leather Legguards (231554)) | Rank 16 [pvp] | 240.8 | yes | General's Leather Legguards (231554, +0.00 DPS, sim-verified) [pvp]; Marshal's Leather Leggings (16456, -0.00 DPS) [vendor]; General's Leather Legguards (16564, -0.00 DPS) [vendor] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 223.1 | yes | Deathdealer's Boots (21359, -3.59 DPS, sim-verified) [quest]; Bloodvine Boots (19684, -9.59 DPS) [crafted]; Shadowcraft Boots (16711, -9.67 DPS) [dungeon] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas: Loatheb [raid] | 211.1 | yes | Band of Earthen Might (21182, -2.78 DPS) [quest]; Ring of the Qiraji Fury (21677, -3.00 DPS) [raid]; Quick Strike Ring (18821, -3.53 DPS) [raid] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 175.1 | yes | Ring of the Qiraji Fury (21677, -1.07 DPS) [raid]; Band of Earthen Might (21182, -1.29 DPS, sim-verified) [quest]; Quick Strike Ring (18821, -1.60 DPS) [raid] |
| trinket1 | Slayer's Crest (23041) | Naxxramas: Sapphiron [raid] | 64.0 | yes | Drake Fang Talisman (19406, +4.28 DPS) [raid]; Neltharion's Tear (19379, +1.28 DPS) [raid]; Eye of Diminution (23001, -11.52 DPS, sim-verified) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas: Maexxna [raid] | 159.1 | yes | Drake Fang Talisman (19406, -0.81 DPS) [raid]; Neltharion's Tear (19379, -3.81 DPS) [raid]; Eye of Diminution (23001, -8.37 DPS, sim-verified) [raid] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Rise, Thunderfury! [quest] | 762.3 | yes | High Warlord's Quickblade (234553, +13.28 DPS) [pvp]; Grand Marshal's Swiftblade (234579, +13.28 DPS) [pvp]; Blessed Qiraji Pugio (21244, -54.17 DPS, sim-verified) [quest] |
| off_hand | The Hungering Cold (23577) | Naxxramas: Kel'Thuzad [raid] | 1022.0 | yes | High Warlord's Left Claw (234558, -0.79 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.79 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -2.47 DPS) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj: Ouro [raid] | 133.1 | yes | The Purifier (22656, -0.96 DPS) [quest]; Fahrad's Reloading Repeater (22347, -4.47 DPS) [quest]; Core Marksman Rifle (18282, -4.77 DPS) [crafted] |

**New at 60:** head: Bonescythe Helmet; neck: Blazefury Medallion; shoulder: Bonescythe Pauldrons; back: Chromatic Cloak; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; legs: Marshal's Leather Leggings; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Don Julio's Band; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: The Hungering Cold; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1716, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

