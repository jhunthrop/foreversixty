# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (night-elf, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.8. Weights run: 1.1s. Verify run: 1.3s. 361 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.028, crit=2.157 ± 0.106, hit=1.402 ± 0.268, melee_haste=not significant (1.169 ± 0.749)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.8 | yes | Shadow Goggles (4373, -0.41 DPS) [crafted]; Lucky Fishing Hat (19972, -0.41 DPS) [quest]; Flying Tiger Goggles (4368, -0.66 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.6 | yes | Tarnished Locket (279870, -0.49 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.5 | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS) [crafted]; Forest Leather Mantle (4709, -0.26 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.41 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.6 | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Catacomb Cloak (279899, -0.09 DPS, sim-verified) [quest]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 12.1 | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.21 DPS) [crafted]; Dark Leather Tunic (2317, -0.26 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.5 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.08 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 30.2 | yes | Serpent Gloves (5970, +0.02 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -1.10 DPS) [dungeon]; Forest Leather Gloves (3058, -1.21 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.58 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.58 DPS) [quest]; Guardsman Belt (3429, -0.64 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.9 | yes | Brawler's Leather Pants (252500, +0.07 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.8 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.15 DPS) [crafted]; Blackened Defias Boots (10402, -0.36 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.6 | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; The 1 Ring (8350, -0.26 DPS) [world]; Minor Channeling Ring (1449, -0.31 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 4.4 | yes | The 1 Ring (8350, -0.15 DPS) [world]; Minor Channeling Ring (1449, -0.21 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.42 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Barrens Basher (274744, -1.20 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.73 DPS) [quest]; Pulsating Hydra Heart (5183, -10.73 DPS) [world]; Tear of Grief (5611, -10.73 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Light Bow (4576, -0.08 DPS) [world_drop]; Owlsight Rifle (15205, -0.08 DPS) [quest]; Deadly Blunderbuss (4369, -0.10 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 361, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (night-elf, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 43.5. Weights run: 1.1s. Verify run: 1.6s. 694 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.099 ± 0.022, crit=3.270 ± 0.122, hit=1.653 ± 0.289, melee_haste=not significant (1.560 ± 0.801)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.16 DPS, sim-verified) [world]; Holy Shroud (2721, -0.52 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.20 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.35 DPS) [rep]; Pendant of Myzrael (4614, -0.66 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 | yes | Dark Leather Shoulders (4252, -0.21 DPS) [crafted]; Insignia Mantle (4721, -0.21 DPS) [world_drop]; Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Cloak of Night (4447, -0.12 DPS, sim-verified) [world]; Fenrus' Hide (6340, -0.16 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.16 DPS) [dungeon] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, +0.15 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.18 DPS) [quest]; Green Leather Armor (4255, -0.34 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, +0.00 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.16 DPS) [world_drop]; Madwolf Bracers (897, -0.21 DPS) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 45.8 | yes | Heavy Earthen Gloves (7359, +0.42 DPS, sim-verified) [crafted]; Pilferer's Gloves (7358, -1.75 DPS) [crafted]; Wolfclaw Gloves (1978, -1.85 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.67 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -0.52 DPS, sim-verified) [crafted]; Insignia Leggings (4054, -0.76 DPS) [world_drop]; Leggings of the Fang (10410, -0.76 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 8.8 | yes | Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.05 DPS) [quest]; Insignia Boots (4055, -0.18 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon]; Protector's Band (19517, -0.16 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, +0.02 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon]; Protector's Band (19517, -0.11 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +4.50 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.62 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.75 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -15.04 DPS) [quest]; Rod of Molten Fire (2565, -15.04 DPS) [world_drop]; Eye of Paleth (2943, -15.04 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Silver Star (3463, -0.14 DPS, sim-verified) [quest]; Moonsight Rifle (4383, -0.19 DPS) [crafted]; Precision Bow (217315, -0.19 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 694, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 85.4. Weights run: 1.2s. Verify run: 1.6s. 962 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.100 ± 0.014, crit=2.696 ± 0.072, hit=1.350 ± 0.103, melee_haste=2.063 ± 0.041

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 13.2 | yes | White Bandit Mask (10008, +0.67 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.04 DPS) [world]; Brawler's Leather Helm (252512, -0.07 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.20 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.17 DPS) [rep]; Sentinel's Medallion (20444, -0.25 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.1 | yes | Nightscape Shoulders (8192, -0.40 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.49 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.16 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.04 DPS) [crafted]; Yeti Fur Cloak (2805, -0.11 DPS) [quest] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 16.5 | yes | Dusky Leather Armor (7374, -0.04 DPS) [crafted]; Hawkeye's Tunic (14592, -0.11 DPS) [world]; Raptorbane Armor (3566, -0.44 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.38 DPS) [world_drop]; Dusky Bracers (7378, -0.38 DPS) [crafted]; Cultist's Armguards (270032, -0.69 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 57.7 | yes | Shadowskin Gloves (18238, -0.67 DPS) [crafted]; Fletcher's Gloves (7348, -1.39 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -1.40 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.20 DPS) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon]; Highlander's Chain Girdle (20090, -0.42 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.43 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.21 DPS) [quest]; Hawkeye's Breeches (14595, -0.36 DPS) [world] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.3 | yes | Imperial Leather Boots (6431, +0.15 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.07 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.07 DPS) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.07 DPS) [dungeon]; Insurgent's Band (272067, -0.10 DPS) [vendor]; Protector's Band (19515, -0.11 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 11.0 | yes | Ironspine's Eye (7686, +0.04 DPS, sim-verified) [dungeon]; Protector's Band (19515, -0.07 DPS) [rep]; Disengagement Ring (276202, -0.07 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -0.89 DPS) [vendor]; Sword of Serenity (6829, -1.34 DPS) [quest]; Hand of Righteousness (7721, -1.43 DPS) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Grayson's Torch (1172, -14.83 DPS) [quest]; Rod of Molten Fire (2565, -14.83 DPS) [world_drop]; Eye of Paleth (2943, -14.83 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Master Hunter's Bow (17686, +0.02 DPS, sim-verified) [quest]; Silver Star (3463, -0.12 DPS) [quest]; Mithril Blunderbuss (10508, -0.12 DPS) [crafted] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 962, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 122.0. Weights run: 1.2s. Verify run: 1.5s. 1220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.103 ± 0.015, crit=3.283 ± 0.085, hit=1.682 ± 0.129, melee_haste=2.580 ± 0.051

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 78.8 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -1.11 DPS) [dungeon]; Helm of Fire (8348, -2.03 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19539, +0.08 DPS, sim-verified) [rep]; Sentinel's Medallion (19540, -0.06 DPS) [rep]; Sentinel's Medallion (19541, -0.17 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 70.8 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -1.58 DPS) [vendor]; Forest Tracker Epaulets (2278, -1.98 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 12.1 | yes | Nightscape Cloak (8195, -0.06 DPS, sim-verified) [crafted]; Wolfmaster Cape (6314, -0.07 DPS) [dungeon]; Pridelord Cape (14673, -0.07 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 80.8 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -1.87 DPS) [quest]; Warbear Harness (15064, -2.06 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.16 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.27 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 66.0 | yes | First Sergeant's Leather Gauntlets (220857, -0.20 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.24 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -0.68 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 66.0 | yes | Highlander's Lizardhide Girdle (20103, -0.68 DPS) [rep]; Highlander's Cloth Girdle (20097, -0.80 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -1.22 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 80.8 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -0.97 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -1.85 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 | yes | Sandstalker Ankleguards (12470, +0.08 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.14 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.14 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.8 | yes | Insurgent's Band (272065, -0.74 DPS) [vendor]; Insurgent's Band (272066, -0.84 DPS) [vendor]; Ring of the Underwood (2951, -0.87 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.4 | yes | Insurgent's Band (272066, -0.12 DPS) [vendor]; Insurgent's Band (272065, -0.15 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.15 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 15.1 | yes | Thunderbrew's Boot Flask (744, -0.51 DPS) [quest]; Tidal Charm (1404, -0.51 DPS) [vendor]; Guardian Talisman (1490, -0.51 DPS) [quest] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Inventor's Focal Sword (17719, -0.48 DPS) [dungeon]; Julie's Dagger (6660, -1.41 DPS) [world_drop]; Might of Hakkar (10838, -1.46 DPS) [world] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 | yes | Claw of Celebras (17738, -1.50 DPS) [dungeon]; Thermotastic Egg Timer (9644, -17.68 DPS) [quest]; Grayson's Torch (1172, -17.79 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 15.4 | yes | Houndmaster's Bow (11628, -0.12 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.22 DPS) [vendor]; Guttbuster (13139, -0.22 DPS) [world] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1220, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 288.3. Weights run: 1.2s. Verify run: 1.9s. 1717 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=not significant (0.000 ± 0.000), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 149.4 | yes | Bloodvine Lens (19998, -1.12 DPS) [crafted]; Champion's Leather Helm (23257, -1.85 DPS) [vendor]; Ragefury Eyepatch (11735, -10.27 DPS, sim-verified) [dungeon] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas: Kel'Thuzad [raid] | 141.5 | yes | Gem of Trapped Innocents (23057, -0.86 DPS) [raid]; Barbed Choker (21664, -1.32 DPS) [raid]; Medallion of the Dawn (22659, -1.98 DPS) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 82.6 | yes | Lieutenant Commander's Leather Shoulders (23313, -0.10 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.10 DPS) [pvp]; Champion's Leather Shoulders (23258, -3.43 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 57.7 | yes | Cloak of the Honor Guard (20073, +1.41 DPS, sim-verified) [rep]; Cape of the Black Baron (13340, -0.69 DPS) [dungeon]; Cloak of the Fallen God (21710, -0.94 DPS) [quest] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 195.5 | yes | Stormshroud Armor (15056, -2.65 DPS) [crafted]; Deathdealer's Vest (21364, -3.18 DPS) [quest]; Zandalar Madcap's Tunic (19834, -7.54 DPS, sim-verified) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 87.1 | yes | General's Leather Armsplints (16559, -2.18 DPS) [pvp]; Forest Stalker's Bracers (19587, -2.18 DPS) [rep]; Marshal's Leather Armsplints (16460, -2.45 DPS, sim-verified) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 123.7 | yes | Marshal's Leather Handgrips (16454, -1.44 DPS) [vendor]; General's Leather Mitts (16560, -1.44 DPS) [vendor]; Devilsaur Gauntlets (15063, -7.70 DPS, sim-verified) [crafted] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj: C'Thun [raid] | 121.7 | yes | Bonescythe Waistguard (22482, -1.22 DPS) [quest]; Highlander's Leather Girdle (20115, -1.46 DPS) [rep]; Highlander's Leather Girdle (20045, -7.06 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) (or Knight-Captain's Leather Legguards (16419), Legionnaire's Leather Leggings (16508)) | Leatherworking [crafted] | 115.5 | yes | Knight-Captain's Leather Legguards (16419, +0.57 DPS, sim-verified) [pvp]; Legionnaire's Leather Leggings (16508, +0.00 DPS) [pvp]; Devilsaur Leggings (15062, -0.39 DPS) [crafted] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 121.7 | yes | Deathdealer's Boots (21359, -3.06 DPS) [quest]; Blood Guard's Leather Walkers (22856, -3.11 DPS) [vendor]; Highlander's Leather Boots (20052, -8.82 DPS, sim-verified) [rep] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas: Loatheb [raid] | 109.7 | yes | Quick Strike Ring (18821, -0.73 DPS) [raid]; Don Julio's Band (19325, -1.19 DPS) [rep]; Band of the Penitent (13217, -1.72 DPS) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj: Viscidus [raid] | 97.7 | yes | Quick Strike Ring (18821, -0.41 DPS, sim-verified) [raid]; Don Julio's Band (19325, -0.80 DPS) [rep]; Band of the Penitent (13217, -1.33 DPS) [quest] |
| trinket1 | Kiss of the Spider (22954) | Naxxramas: Maexxna [raid] | 57.7 | yes | Drake Fang Talisman (19406, -0.06 DPS) [raid]; Thunderbrew's Boot Flask (744, -1.91 DPS) [quest]; Eye of Diminution (23001, -6.86 DPS, sim-verified) [raid] |
| trinket2 | Slayer's Crest (23041) | Naxxramas: Sapphiron [raid] | 64.0 | yes | Drake Fang Talisman (19406, -0.27 DPS) [raid]; Thunderbrew's Boot Flask (744, -2.12 DPS) [quest]; Eye of Diminution (23001, -5.45 DPS, sim-verified) [raid] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Rise, Thunderfury! [quest] | 761.0 | yes | Death's Sting (21126, +6.85 DPS) [raid]; Grand Marshal's Swiftblade (234579, +6.36 DPS) [pvp]; The Hungering Cold (23577, -17.02 DPS, sim-verified) [raid] |
| off_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 952.9 | yes | High Warlord's Left Claw (234558, -0.10 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.10 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.14 DPS) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj: Ouro [raid] | 75.7 | yes | The Purifier (22656, -0.60 DPS) [quest]; Huhuran's Stinger (21616, -1.84 DPS) [raid]; Precisely Calibrated Boomstick (2100, -1.99 DPS) [world_drop] |

**New at 60:** head: Bonescythe Helmet; neck: Stormrage's Talisman of Seething; shoulder: Bonescythe Pauldrons; back: Chromatic Cloak; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; legs: Stormshroud Pants; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Kiss of the Spider; trinket2: Slayer's Crest; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: High Warlord's Quickblade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1717, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (troll, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.1. Weights run: 1.1s. Verify run: 1.3s. 360 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.028, crit=2.157 ± 0.106, hit=1.402 ± 0.268, melee_haste=not significant (1.169 ± 0.749)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.8 | yes | Shadow Goggles (4373, -0.41 DPS) [crafted]; Lucky Fishing Hat (19972, -0.41 DPS) [quest]; Flying Tiger Goggles (4368, -0.65 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.6 | yes | Tarnished Locket (279870, -0.48 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.5 | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS) [crafted]; Forest Leather Mantle (4709, -0.26 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.6 | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Catacomb Cloak (279899, -0.09 DPS, sim-verified) [quest]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.7 | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Heckler's Hide (286536, -0.10 DPS) [world]; Trapper's Leather Armor (252491, -0.33 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.5 | yes | Wolf Bracers (4794, -0.08 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor]; Spare Part Bindings (279875, -0.10 DPS) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 30.2 | yes | Serpent Gloves (5970, +0.02 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -1.10 DPS) [dungeon]; Forest Leather Gloves (3058, -1.21 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.58 DPS) [quest]; Deviate Scale Belt (6468, -0.60 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.64 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.9 | yes | Brawler's Leather Pants (252500, +0.07 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.8 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.15 DPS) [crafted]; Blackened Defias Boots (10402, -0.35 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.6 | yes | Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; The 1 Ring (8350, -0.26 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 4.4 | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The 1 Ring (8350, -0.15 DPS) [world]; Bounty Hunter's Ring (5351, -0.34 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Wingblade (6504, -1.12 DPS) [quest] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.73 DPS) [quest]; Nightglow Concoction (3451, -10.73 DPS) [quest]; Pulsating Hydra Heart (5183, -10.73 DPS) [world] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Light Bow (4576, -0.08 DPS) [world_drop]; Privateer Musket (5309, -0.08 DPS) [quest]; Deadly Blunderbuss (4369, -0.11 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 360, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers

### Band 30 (troll, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 43.1. Weights run: 1.1s. Verify run: 1.7s. 695 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.099 ± 0.022, crit=3.270 ± 0.122, hit=1.653 ± 0.289, melee_haste=not significant (1.560 ± 0.801)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.15 DPS, sim-verified) [world]; Holy Shroud (2721, -0.52 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.18 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.35 DPS) [rep]; Pendant of Myzrael (4614, -0.66 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 | yes | Dark Leather Shoulders (4252, -0.21 DPS) [crafted]; Insignia Mantle (4721, -0.21 DPS) [world_drop]; Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Cloak of Night (4447, -0.16 DPS) [world]; Fenrus' Hide (6340, -0.16 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 15.4 | yes | Panther Armor (6670, -0.27 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.31 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.31 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, +0.01 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.16 DPS) [world_drop]; Madwolf Bracers (897, -0.21 DPS) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 45.8 | yes | Heavy Earthen Gloves (7359, +0.41 DPS, sim-verified) [crafted]; Pilferer's Gloves (7358, -1.75 DPS) [crafted]; Braced Handguards (6784, -1.80 DPS) [quest] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.36 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -0.51 DPS, sim-verified) [crafted]; Insignia Leggings (4054, -0.76 DPS) [world_drop]; Leggings of the Fang (10410, -0.76 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 8.8 | yes | Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Insignia Boots (4055, -0.18 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon]; Legionnaire's Band (19513, -0.16 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, +0.03 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon]; Legionnaire's Band (19513, -0.11 DPS) [rep] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +4.71 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.62 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.75 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -15.04 DPS) [quest]; Rod of Molten Fire (2565, -15.04 DPS) [world_drop]; Nightglow Concoction (3451, -15.04 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Silver Star (3463, -0.14 DPS, sim-verified) [quest]; Moonsight Rifle (4383, -0.19 DPS) [crafted]; Precision Bow (217315, -0.19 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 695, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow

### Band 40 (troll, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 83.1. Weights run: 1.2s. Verify run: 1.7s. 962 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.100 ± 0.014, crit=2.696 ± 0.072, hit=1.350 ± 0.103, melee_haste=2.063 ± 0.041

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 13.2 | yes | White Bandit Mask (10008, +0.64 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.04 DPS) [world]; Spirit Hunter Headdress (6720, -0.07 DPS) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.20 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.17 DPS) [rep]; Scout's Medallion (20442, -0.25 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.1 | yes | Nightscape Shoulders (8192, -0.40 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.48 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.04 DPS) [world_drop]; Parachute Cloak (10518, -0.04 DPS) [crafted] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 16.5 | yes | Dusky Leather Armor (7374, -0.11 DPS, sim-verified) [crafted]; Hawkeye's Tunic (14592, -0.11 DPS) [world]; Panther Armor (6670, -0.22 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.38 DPS) [world_drop]; Dusky Bracers (7378, -0.38 DPS) [crafted]; Cultist's Armguards (270032, -0.68 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 57.7 | yes | Shadowskin Gloves (18238, -0.67 DPS) [crafted]; Fletcher's Gloves (7348, -1.37 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -1.40 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.20 DPS) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon]; Defiler's Chain Girdle (20152, -0.41 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.44 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.21 DPS) [quest]; Hawkeye's Breeches (14595, -0.36 DPS) [world] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.3 | yes | Imperial Leather Boots (6431, +0.12 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.07 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.07 DPS) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.07 DPS) [dungeon]; Insurgent's Band (272067, -0.10 DPS) [vendor]; Legionnaire's Band (19512, -0.11 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 11.0 | yes | Ironspine's Eye (7686, +0.03 DPS, sim-verified) [dungeon]; Legionnaire's Band (19512, -0.07 DPS) [rep]; Disengagement Ring (276202, -0.07 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -0.89 DPS) [vendor]; Hand of Righteousness (7721, -1.43 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272093, -1.44 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Grayson's Torch (1172, -14.83 DPS) [quest]; Rod of Molten Fire (2565, -14.83 DPS) [world_drop]; Nightglow Concoction (3451, -14.83 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Master Hunter's Bow (17686, -0.01 DPS, sim-verified) [quest]; Silver Star (3463, -0.12 DPS) [quest]; Mithril Blunderbuss (10508, -0.12 DPS) [crafted] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 962, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (troll, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 120.5. Weights run: 1.2s. Verify run: 1.6s. 1220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.103 ± 0.015, crit=3.283 ± 0.085, hit=1.682 ± 0.129, melee_haste=2.580 ± 0.051

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 78.8 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -1.11 DPS) [dungeon]; Helm of Fire (8348, -2.03 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19535, +0.09 DPS, sim-verified) [rep]; Scout's Medallion (19536, -0.06 DPS) [rep]; Woven Ivy Necklace (19159, -0.14 DPS) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 70.8 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -1.58 DPS) [vendor]; Forest Tracker Epaulets (2278, -1.98 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 12.1 | yes | Nightscape Cloak (8195, -0.05 DPS, sim-verified) [crafted]; Wolfmaster Cape (6314, -0.07 DPS) [dungeon]; Battlehard Cape (11858, -0.07 DPS) [quest] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 80.8 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -1.87 DPS) [quest]; Warbear Harness (15064, -2.06 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.17 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.27 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 66.0 | yes | First Sergeant's Leather Gauntlets (220857, -0.20 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.24 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -0.68 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 66.0 | yes | Defiler's Lizardhide Girdle (20174, -0.68 DPS) [rep]; Defiler's Cloth Girdle (20165, -0.79 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -1.22 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 80.8 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -0.99 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -1.85 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 | yes | Sandstalker Ankleguards (12470, +0.08 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.14 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.14 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.8 | yes | Masons Fraternity Ring (9533, -0.72 DPS) [quest]; Insurgent's Band (272065, -0.74 DPS) [vendor]; Insurgent's Band (272066, -0.84 DPS) [vendor] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Masons Fraternity Ring (9533, -0.20 DPS, sim-verified) [quest]; Insurgent's Band (272065, -0.30 DPS) [vendor]; Insurgent's Band (272066, -0.41 DPS) [vendor] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, +1.82 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 15.1 | yes | Rune of the Guard Captain (19120, +1.31 DPS) [quest]; Tidal Charm (1404, -0.51 DPS) [vendor]; Guardian Talisman (1490, -0.51 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Inventor's Focal Sword (17719, +0.00 DPS, sim-verified) [dungeon]; Julie's Dagger (6660, -1.41 DPS) [world_drop]; Might of Hakkar (10838, -1.46 DPS) [world] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 | yes | Claw of Celebras (17738, -1.50 DPS) [dungeon]; White Bone Shredder (11863, -2.83 DPS) [quest]; Thermotastic Egg Timer (9644, -17.68 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 15.4 | yes | Houndmaster's Bow (11628, -0.12 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.22 DPS) [vendor]; Guttbuster (13139, -0.22 DPS) [world] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Defiler's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Smoking Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1220, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 288.0. Weights run: 1.2s. Verify run: 1.9s. 1716 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=not significant (0.000 ± 0.000), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 149.4 | yes | Bloodvine Lens (19998, -1.12 DPS) [crafted]; Champion's Leather Helm (23257, -1.85 DPS) [vendor]; Ragefury Eyepatch (11735, -11.56 DPS, sim-verified) [dungeon] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas: Kel'Thuzad [raid] | 141.5 | yes | Gem of Trapped Innocents (23057, -0.86 DPS) [raid]; Barbed Choker (21664, -1.32 DPS) [raid]; Medallion of the Dawn (22659, -1.98 DPS) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 82.6 | yes | Lieutenant Commander's Leather Shoulders (23313, -0.10 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.10 DPS) [pvp]; Champion's Leather Shoulders (23258, -3.48 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 57.7 | yes | Deathguard's Cloak (20068, +1.28 DPS, sim-verified) [rep]; Cape of the Black Baron (13340, -0.69 DPS) [dungeon]; Cloak of the Fallen God (21710, -0.94 DPS) [quest] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 195.5 | yes | Stormshroud Armor (15056, -2.65 DPS) [crafted]; Deathdealer's Vest (21364, -3.18 DPS) [quest]; Zandalar Madcap's Tunic (19834, -8.59 DPS, sim-verified) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 87.1 | yes | General's Leather Armsplints (16559, -2.18 DPS) [pvp]; Forest Stalker's Bracers (19587, -2.18 DPS) [rep]; Marshal's Leather Armsplints (16460, -2.61 DPS, sim-verified) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 123.7 | yes | Marshal's Leather Handgrips (16454, -1.44 DPS) [vendor]; General's Leather Mitts (16560, -1.44 DPS) [vendor]; Devilsaur Gauntlets (15063, -8.75 DPS, sim-verified) [crafted] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj: C'Thun [raid] | 121.7 | yes | Bonescythe Waistguard (22482, -1.22 DPS) [quest]; Defiler's Leather Girdle (20193, -1.46 DPS) [rep]; Defiler's Leather Girdle (20190, -8.11 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) (or Knight-Captain's Leather Legguards (16419), Legionnaire's Leather Leggings (16508)) | Leatherworking [crafted] | 115.5 | yes | Knight-Captain's Leather Legguards (16419, +0.57 DPS, sim-verified) [pvp]; Legionnaire's Leather Leggings (16508, +0.00 DPS) [pvp]; Devilsaur Leggings (15062, -0.39 DPS) [crafted] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 121.7 | yes | Deathdealer's Boots (21359, -3.06 DPS) [quest]; Blood Guard's Leather Walkers (22856, -3.11 DPS) [vendor]; Defiler's Leather Boots (20186, -9.98 DPS, sim-verified) [rep] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas: Loatheb [raid] | 109.7 | yes | Quick Strike Ring (18821, -0.73 DPS) [raid]; Don Julio's Band (19325, -1.19 DPS) [rep]; Band of the Penitent (13217, -1.72 DPS) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj: Viscidus [raid] | 97.7 | yes | Quick Strike Ring (18821, -0.41 DPS, sim-verified) [raid]; Don Julio's Band (19325, -0.80 DPS) [rep]; Band of the Penitent (13217, -1.33 DPS) [quest] |
| trinket1 | Kiss of the Spider (22954) | Naxxramas: Maexxna [raid] | 57.7 | yes | Drake Fang Talisman (19406, -0.06 DPS) [raid]; Rune of the Guard Captain (19120, -0.52 DPS) [quest]; Eye of Diminution (23001, -8.78 DPS, sim-verified) [raid] |
| trinket2 | Slayer's Crest (23041) | Naxxramas: Sapphiron [raid] | 64.0 | yes | Drake Fang Talisman (19406, -0.27 DPS) [raid]; Rune of the Guard Captain (19120, -0.73 DPS) [quest]; Eye of Diminution (23001, -5.32 DPS, sim-verified) [raid] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Rise, Thunderfury! [quest] | 761.0 | yes | Death's Sting (21126, +6.85 DPS) [raid]; Grand Marshal's Swiftblade (234579, +6.36 DPS) [pvp]; The Hungering Cold (23577, -18.10 DPS, sim-verified) [raid] |
| off_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 952.9 | yes | High Warlord's Left Claw (234558, -0.10 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.10 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.14 DPS) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj: Ouro [raid] | 75.7 | yes | The Purifier (22656, -0.60 DPS) [quest]; Huhuran's Stinger (21616, -1.84 DPS) [raid]; Precisely Calibrated Boomstick (2100, -1.99 DPS) [world_drop] |

**New at 60:** head: Bonescythe Helmet; neck: Stormrage's Talisman of Seething; shoulder: Bonescythe Pauldrons; back: Chromatic Cloak; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; legs: Stormshroud Pants; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Kiss of the Spider; trinket2: Slayer's Crest; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: High Warlord's Quickblade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1716, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

