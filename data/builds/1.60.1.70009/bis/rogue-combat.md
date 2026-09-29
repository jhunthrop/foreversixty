# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.5. Weights run: 0.8s. Verify run: 1.0s. 360 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.009 ± 0.003, crit=0.390 ± 0.062, hit=not significant (0.899 ± 0.286), melee_haste=not significant (0.636 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.38 DPS) [crafted]; Lucky Fishing Hat (19972, -0.38 DPS) [quest]; Flying Tiger Goggles (4368, -0.54 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 | yes | Tarnished Locket (279870, -0.41 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Forest Leather Mantle (4709, -0.24 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.34 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.02 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [dungeon] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.1 | yes | Brawler's Leather Armor (252490, +0.05 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.19 DPS) [crafted]; Dark Leather Tunic (2317, -0.24 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 5.0 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.07 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.26 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.03 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.06 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.1 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.33 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world]; Minor Channeling Ring (1449, -0.29 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 4.0 | yes | The 1 Ring (8350, -0.14 DPS) [world]; Minor Channeling Ring (1449, -0.19 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.40 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [dungeon]; Diamond Hammer (2194, -1.05 DPS) [dungeon]; Barrens Basher (274744, -1.21 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.85 DPS) [quest]; Pulsating Hydra Heart (5183, -10.85 DPS) [world]; Tear of Grief (5611, -10.85 DPS) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Light Bow (4576, -0.09 DPS) [dungeon]; Owlsight Rifle (15205, -0.09 DPS) [quest]; Deadly Blunderbuss (4369, -0.13 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 360, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (night-elf, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.7. Weights run: 0.7s. Verify run: 1.2s. 691 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.033 ± 0.019, crit=0.465 ± 0.062, hit=2.484 ± 0.367, melee_haste=not significant (0.623 ± 0.730)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.3 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Holy Shroud (2721, -0.50 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.36 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.38 DPS) [rep]; Pendant of Myzrael (4614, -0.68 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 11.4 | yes | Dark Leather Shoulders (4252, -0.20 DPS) [crafted]; Insignia Mantle (4721, -0.20 DPS) [dungeon]; Mantle of Thieves (2264, -0.39 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Fenrus' Hide (6340, -0.18 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.18 DPS) [dungeon]; Cloak of Night (4447, -0.24 DPS, sim-verified) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.08 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.23 DPS) [quest]; Green Leather Armor (4255, -0.38 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.11 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.18 DPS) [dungeon]; Madwolf Bracers (897, -0.23 DPS) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Fletcher's Gloves (7348, -0.46 DPS) [crafted]; Wolfclaw Gloves (1978, -0.48 DPS) [dungeon]; Pilferer's Gloves (7358, -0.49 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.71 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -0.80 DPS, sim-verified) [crafted]; Insignia Leggings (4054, -0.81 DPS) [dungeon]; Leggings of the Fang (10410, -0.81 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.3 | yes | Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.05 DPS) [quest]; Insignia Boots (4055, -0.19 DPS, sim-verified) [dungeon] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.3 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.15 DPS) [dungeon]; Protector's Band (19517, -0.15 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.11 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +2.07 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.64 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.78 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -15.45 DPS) [quest]; Rod of Molten Fire (2565, -15.45 DPS) [dungeon]; Eye of Paleth (2943, -15.45 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Precision Bow (217315, -0.08 DPS) [quest]; Silver Star (3463, -0.19 DPS) [quest]; Moonsight Rifle (4383, -0.41 DPS, sim-verified) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 88.9. Weights run: 0.8s. Verify run: 1.2s. 958 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.962 ± 0.146, hit=4.384 ± 0.987, melee_haste=not significant (-0.246 ± 1.799)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.1 | yes | White Bandit Mask (10008, +0.77 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.05 DPS) [world]; Brawler's Leather Helm (252512, -0.10 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.19 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.30 DPS) [rep]; Sentinel's Medallion (20444, -0.40 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [dungeon]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.16 DPS, sim-verified) [dungeon]; Parachute Cloak (10518, -0.10 DPS) [crafted]; Yeti Fur Cloak (2805, -0.20 DPS) [quest] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Nightscape Tunic (8175, +0.50 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.09 DPS) [crafted]; Hawkeye's Tunic (14592, -0.19 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.60 DPS) [dungeon]; Dusky Bracers (7378, -0.60 DPS) [crafted]; Cultist's Armguards (270032, -0.81 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 33.5 | yes | Fletcher's Gloves (7348, -1.00 DPS) [crafted]; Shadowskin Gloves (18238, -1.00 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.28 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.49 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.41 DPS, sim-verified) [dungeon]; Triprunner Dungarees (9624, -0.39 DPS) [quest]; Hawkeye's Breeches (14595, -0.60 DPS) [world] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 13.1 | yes | Imperial Leather Boots (6431, +0.17 DPS, sim-verified) [dungeon]; Dusky Boots (7390, -0.10 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.10 DPS) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Protector's Band (19515, -0.20 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 10.1 | yes | Ironspine's Eye (7686, +0.04 DPS, sim-verified) [dungeon]; Protector's Band (19515, -0.10 DPS) [rep]; Disengagement Ring (276202, -0.10 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -1.37 DPS) [vendor]; Sword of Serenity (6829, -2.01 DPS) [quest]; Hand of Righteousness (7721, -2.14 DPS) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Grayson's Torch (1172, -22.18 DPS) [quest]; Rod of Molten Fire (2565, -22.18 DPS) [dungeon]; Eye of Paleth (2943, -22.18 DPS) [quest] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 13.2 | yes | Precision Bow (217315, +0.00 DPS, sim-verified) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.21 DPS) [vendor]; Master Hunter's Bow (17686, -0.36 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Moonsight Rifle

No-known-source sample (15 of 958, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 105.9. Weights run: 0.7s. Verify run: 1.2s. 1261 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.239 ± 0.059, crit=6.333 ± 0.367, hit=not significant (5.215 ± 1.384), melee_haste=not significant (-1.529 ± 2.421)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 88.7 | yes | Helm of Fire (8348, +0.85 DPS, sim-verified) [crafted]; Lordrec Helmet (10741, -3.73 DPS) [quest]; Sprightring Helm (17776, -3.79 DPS) [quest] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 14.9 | yes | Sentinel's Medallion (19540, -0.07 DPS) [rep]; Sentinel's Medallion (19541, -0.27 DPS) [rep]; Ghostshard Talisman (7731, -0.28 DPS, sim-verified) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.6 | yes | Forest Tracker Epaulets (2278, -0.53 DPS, sim-verified) [dungeon]; Nightscape Shoulders (8192, -0.65 DPS) [crafted]; Penance Spaulders (11963, -0.65 DPS) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 13.6 | yes | Nightscape Cloak (8195, -0.13 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.13 DPS) [dungeon]; Wolfmaster Cape (6314, -0.20 DPS) [dungeon] |
| chest | Blazewind Breastplate (11193) | Tremors of the Earth [quest] | 28.5 | yes | Warbear Harness (15064, +0.05 DPS, sim-verified) [crafted]; Charred Leather Tunic (19127, -0.34 DPS) [quest]; Nightscape Tunic (8175, -0.54 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.40 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.35 DPS) [crafted]; Pridelord Bands (14672, -0.41 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 108.7 | yes | Shadowskin Gloves (18238, -1.08 DPS) [crafted]; Fletcher's Gloves (7348, -1.52 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -5.01 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 108.7 | yes | Highlander's Lizardhide Girdle (20103, -1.08 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.52 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -4.26 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 177.3 | yes | Basilisk Hide Pants (1718, +0.30 DPS, sim-verified) [dungeon]; Ferine Leggings (6690, -8.19 DPS) [dungeon]; Keeper's Woolies (14668, -8.32 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.8 | yes | Sandstalker Ankleguards (12470, +0.12 DPS, sim-verified) [dungeon]; Swampwalker Boots (2276, -0.47 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.47 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 72.1 | yes | Insurgent's Band (272065, -3.09 DPS) [vendor]; Ring of the Underwood (2951, -3.23 DPS) [dungeon]; Insurgent's Band (272066, -3.26 DPS) [vendor] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 17.3 | yes | Ring of the Underwood (2951, -0.27 DPS) [dungeon]; Insurgent's Band (272066, -0.29 DPS) [vendor]; Insurgent's Band (272065, -0.43 DPS, sim-verified) [vendor] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 581.9 | yes | Might of Hakkar (10838, -2.93 DPS) [world]; Julie's Dagger (6660, -3.81 DPS) [dungeon]; Lifeforce Dirk (10750, -4.26 DPS) [quest] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -3.85 DPS) [dungeon]; Thermotastic Egg Timer (9644, -29.74 DPS) [quest]; Grayson's Torch (1172, -29.94 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.3 | yes | Moonsight Rifle (4383, -0.09 DPS) [crafted]; Precision Bow (217315, -0.09 DPS) [quest]; Houndmaster's Bow (11628, -0.29 DPS) [dungeon] |

**New at 50:** head: Eye of Theradras; neck: Sentinel's Medallion; back: Serpentskin Cloak; chest: Blazewind Breastplate; waist: Highlander's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Smoking Heart of the Mountain; trinket2: Talisman of Arathor; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1261, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 279.5. Weights run: 0.8s. Verify run: 1.3s. 1866 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 316.1 | yes | Bloodvine Lens (19998, -4.60 DPS) [crafted]; Mask of the Unforgiven (13404, -6.05 DPS) [dungeon]; Ragefury Eyepatch (11735, -10.25 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 168.9 | yes | Medallion of the Dawn (22659, -1.59 DPS) [quest]; Beads of Ogre Might (22150, -5.40 DPS) [quest]; Choker of the Shifting Sands (21505, -6.79 DPS) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 189.8 | yes | Lieutenant Commander's Leather Shoulders (23313, -0.47 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.47 DPS) [pvp]; Champion's Leather Shoulders (23258, -4.83 DPS, sim-verified) [vendor] |
| back | Cloak of Veiled Shadows (21406) | Cloak of Veiled Shadows [quest] | 69.1 | yes | Earthweave Cloak (21187, -0.22 DPS) [quest]; Cloak of the Honor Guard (20073, -1.51 DPS) [rep]; Chromatic Cloak (18509, -3.41 DPS, sim-verified) [crafted] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 354.2 | yes | Stormshroud Armor (15056, -6.64 DPS) [crafted]; Zandalar Madcap's Tunic (19834, -6.73 DPS, sim-verified) [quest]; Deathdealer's Vest (21364, -7.68 DPS) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 151.4 | yes | Primal Batskin Bracers (19687, -2.78 DPS, sim-verified) [crafted]; Rockfury Bracers (21186, -5.75 DPS) [quest]; Marshal's Leather Armsplints (16460, -6.69 DPS) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 225.1 | yes | Devilsaur Gauntlets (15063, -4.39 DPS) [crafted]; Marshal's Leather Handgrips (16454, -4.39 DPS) [vendor]; Stormshroud Gloves (21278, -8.45 DPS, sim-verified) [crafted] |
| waist | Bonescythe Waistguard (22482) | Bonescythe Waistguard [quest] | 148.6 | yes | Highlander's Leather Girdle (20115, -0.72 DPS) [rep]; Belt of the Archmage (18405, -1.79 DPS) [crafted]; Highlander's Leather Girdle (20045, -3.80 DPS, sim-verified) [rep] |
| legs | Marshal's Leather Leggings (16456) (or General's Leather Legguards (16564), Marshal's Leather Leggings (231548), General's Leather Legguards (231554)) | Captain Dirgehammer [vendor] | 240.8 | yes | General's Leather Legguards (16564, +0.00 DPS, sim-verified) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 223.1 | yes | Deathdealer's Boots (21359, -3.34 DPS, sim-verified) [quest]; Bloodvine Boots (19684, -9.59 DPS) [crafted]; Shadowcraft Boots (16711, -9.67 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 175.1 | yes | Band of the Penitent (13217, -3.21 DPS) [quest]; Dragonslayer's Signet (18403, -3.21 DPS) [quest]; Ring of Entropy (18543, -3.21 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 159.1 | yes | Dragonslayer's Signet (18403, -2.35 DPS) [quest]; Ring of Entropy (18543, -2.35 DPS) [world]; Band of the Penitent (13217, -2.91 DPS, sim-verified) [quest] |
| trinket1 | Onyxia Blood Talisman (18406) | For All To See [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 1010.3 | yes | High Warlord's Blade (234552, -0.16 DPS) [pvp]; High Warlord's Bludgeon (234555, -0.16 DPS) [pvp]; High Warlord's Right Claw (234557, -0.16 DPS) [pvp] |
| off_hand | Grand Marshal's Swiftblade (234579) | Rank 18 [pvp] | 1010.3 | yes | High Warlord's Left Claw (234558, -0.16 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.16 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.84 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 115.1 | yes | Fahrad's Reloading Repeater (22347, -3.51 DPS) [quest]; Core Marksman Rifle (18282, -3.81 DPS) [crafted]; Blackcrow (12651, -3.82 DPS) [dungeon] |

**New at 60:** head: Bonescythe Helmet; neck: Onyxia Tooth Pendant; shoulder: Bonescythe Pauldrons; back: Cloak of Veiled Shadows; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Bonescythe Waistguard; legs: Marshal's Leather Leggings; feet: Bonescythe Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Onyxia Blood Talisman; trinket2: Talisman of Arathor; main_hand: High Warlord's Quickblade; off_hand: Grand Marshal's Swiftblade; ranged: The Purifier

No-known-source sample (15 of 1866, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.0. Weights run: 0.8s. Verify run: 1.0s. 355 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.009 ± 0.003, crit=0.390 ± 0.062, hit=not significant (0.899 ± 0.286), melee_haste=not significant (0.636 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.38 DPS) [crafted]; Lucky Fishing Hat (19972, -0.38 DPS) [quest]; Flying Tiger Goggles (4368, -0.51 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 | yes | Tarnished Locket (279870, -0.39 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Forest Leather Mantle (4709, -0.24 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.32 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Heckler's Hide (286536, -0.10 DPS) [world]; Trapper's Leather Armor (252491, -0.32 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 5.0 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.26 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.03 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.06 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.1 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.32 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 4.0 | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The 1 Ring (8350, -0.14 DPS) [world]; Bounty Hunter's Ring (5351, -0.32 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [dungeon]; Diamond Hammer (2194, -1.05 DPS) [dungeon]; Wingblade (6504, -1.15 DPS) [quest] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.85 DPS) [quest]; Nightglow Concoction (3451, -10.85 DPS) [quest]; Pulsating Hydra Heart (5183, -10.85 DPS) [world] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Light Bow (4576, -0.09 DPS) [dungeon]; Privateer Musket (5309, -0.09 DPS) [quest]; Deadly Blunderbuss (4369, -0.13 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 355, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (troll, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.3. Weights run: 0.7s. Verify run: 1.3s. 687 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.033 ± 0.019, crit=0.465 ± 0.062, hit=2.484 ± 0.367, melee_haste=not significant (0.623 ± 0.730)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.3 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Holy Shroud (2721, -0.50 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.38 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.38 DPS) [rep]; Pendant of Myzrael (4614, -0.68 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 11.4 | yes | Dark Leather Shoulders (4252, -0.20 DPS) [crafted]; Insignia Mantle (4721, -0.20 DPS) [dungeon]; Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Cloak of Night (4447, -0.18 DPS) [world]; Fenrus' Hide (6340, -0.18 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.5 | yes | Panther Armor (6670, -0.19 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.30 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.30 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.12 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.18 DPS) [dungeon]; Madwolf Bracers (897, -0.23 DPS) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Braced Handguards (6784, -0.43 DPS) [quest]; Fletcher's Gloves (7348, -0.46 DPS) [crafted]; Pilferer's Gloves (7358, -0.50 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Deftkin Belt (16659, -0.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Insignia Leggings (4054, -0.81 DPS) [dungeon]; Leggings of the Fang (10410, -0.81 DPS) [dungeon]; Dusky Leather Leggings (7373, -0.82 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.3 | yes | Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Insignia Boots (4055, -0.19 DPS, sim-verified) [dungeon] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.3 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.15 DPS) [dungeon]; Legionnaire's Band (19513, -0.15 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.12 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +1.87 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.64 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.78 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -15.45 DPS) [quest]; Rod of Molten Fire (2565, -15.45 DPS) [dungeon]; Nightglow Concoction (3451, -15.45 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Precision Bow (217315, -0.08 DPS) [quest]; Silver Star (3463, -0.19 DPS) [quest]; Moonsight Rifle (4383, -0.35 DPS, sim-verified) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 687, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 87.9. Weights run: 0.8s. Verify run: 1.2s. 954 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.962 ± 0.146, hit=4.384 ± 0.987, melee_haste=not significant (-0.246 ± 1.799)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.1 | yes | White Bandit Mask (10008, +0.78 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.05 DPS) [world]; Spirit Hunter Headdress (6720, -0.10 DPS) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.17 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.30 DPS) [rep]; Scout's Medallion (20442, -0.40 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [dungeon]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.10 DPS) [dungeon]; Parachute Cloak (10518, -0.10 DPS) [crafted] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 15.2 | yes | Dusky Leather Armor (7374, -0.10 DPS, sim-verified) [crafted]; Hawkeye's Tunic (14592, -0.15 DPS) [world]; Panther Armor (6670, -0.30 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.60 DPS) [dungeon]; Dusky Bracers (7378, -0.60 DPS) [crafted]; Cultist's Armguards (270032, -0.80 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 33.5 | yes | Fletcher's Gloves (7348, -1.00 DPS) [crafted]; Shadowskin Gloves (18238, -1.00 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.43 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.48 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.38 DPS, sim-verified) [dungeon]; Triprunner Dungarees (9624, -0.39 DPS) [quest]; Hawkeye's Breeches (14595, -0.60 DPS) [world] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 13.1 | yes | Imperial Leather Boots (6431, +0.17 DPS, sim-verified) [dungeon]; Dusky Boots (7390, -0.10 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.10 DPS) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Legionnaire's Band (19512, -0.20 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 10.1 | yes | Ironspine's Eye (7686, +0.06 DPS, sim-verified) [dungeon]; Legionnaire's Band (19512, -0.10 DPS) [rep]; Disengagement Ring (276202, -0.10 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -1.37 DPS) [vendor]; Hand of Righteousness (7721, -2.14 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272093, -2.16 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Grayson's Torch (1172, -22.18 DPS) [quest]; Rod of Molten Fire (2565, -22.18 DPS) [dungeon]; Nightglow Concoction (3451, -22.18 DPS) [quest] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 13.2 | yes | Precision Bow (217315, +0.00 DPS, sim-verified) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.21 DPS) [vendor]; Master Hunter's Bow (17686, -0.36 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Moonsight Rifle

No-known-source sample (15 of 954, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 108.2. Weights run: 0.7s. Verify run: 1.2s. 1257 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.239 ± 0.059, crit=6.333 ± 0.367, hit=not significant (5.215 ± 1.384), melee_haste=not significant (-1.529 ± 2.421)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 88.7 | yes | Helm of Fire (8348, +0.82 DPS, sim-verified) [crafted]; Sprightring Helm (17776, -3.79 DPS) [quest]; Nightscape Headband (8176, -3.99 DPS) [crafted] |
| neck | Scout's Medallion (19535) | Warsong Outriders [rep] | 14.9 | yes | Scout's Medallion (19536, -0.07 DPS) [rep]; Woven Ivy Necklace (19159, -0.20 DPS) [quest]; Ghostshard Talisman (7731, -0.30 DPS, sim-verified) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.6 | yes | Forest Tracker Epaulets (2278, -0.53 DPS, sim-verified) [dungeon]; Nightscape Shoulders (8192, -0.65 DPS) [crafted]; Penance Spaulders (11963, -0.65 DPS) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 13.6 | yes | Nightscape Cloak (8195, -0.11 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.13 DPS) [dungeon]; Wolfmaster Cape (6314, -0.20 DPS) [dungeon] |
| chest | Blazewind Breastplate (11193) | Tremors of the Earth [quest] | 28.5 | yes | Warbear Harness (15064, +0.05 DPS, sim-verified) [crafted]; Charred Leather Tunic (19127, -0.34 DPS) [quest]; Nightscape Tunic (8175, -0.54 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.38 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.35 DPS) [crafted]; Pridelord Bands (14672, -0.41 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 108.7 | yes | Shadowskin Gloves (18238, -1.08 DPS) [crafted]; Fletcher's Gloves (7348, -1.50 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -5.01 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 108.7 | yes | Defiler's Lizardhide Girdle (20174, -1.08 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.50 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -4.26 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 177.3 | yes | Basilisk Hide Pants (1718, +0.20 DPS, sim-verified) [dungeon]; Ferine Leggings (6690, -8.19 DPS) [dungeon]; Keeper's Woolies (14668, -8.32 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.8 | yes | Sandstalker Ankleguards (12470, +0.13 DPS, sim-verified) [dungeon]; Swampwalker Boots (2276, -0.47 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.47 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 72.1 | yes | Masons Fraternity Ring (9533, -2.97 DPS) [quest]; Insurgent's Band (272065, -3.09 DPS) [vendor]; Ring of the Underwood (2951, -3.23 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Masons Fraternity Ring (9533, -0.33 DPS, sim-verified) [quest]; Insurgent's Band (272065, -0.49 DPS) [vendor]; Ring of the Underwood (2951, -0.63 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 78.5 | yes | Tidal Charm (1404, -4.25 DPS) [vendor]; Guardian Talisman (1490, -4.25 DPS) [quest]; Ankh of Life (1713, -4.25 DPS) [dungeon] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 581.9 | yes | Might of Hakkar (10838, -2.93 DPS) [world]; Julie's Dagger (6660, -3.81 DPS) [dungeon]; Lifeforce Dirk (10750, -4.26 DPS) [quest] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -3.85 DPS) [dungeon]; White Bone Shredder (11863, -5.93 DPS) [quest]; Thermotastic Egg Timer (9644, -29.74 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.3 | yes | Moonsight Rifle (4383, -0.09 DPS) [crafted]; Precision Bow (217315, -0.09 DPS) [quest]; Houndmaster's Bow (11628, -0.29 DPS) [dungeon] |

**New at 50:** head: Eye of Theradras; neck: Scout's Medallion; back: Serpentskin Cloak; chest: Blazewind Breastplate; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Smoking Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1257, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 284.2. Weights run: 0.8s. Verify run: 1.3s. 1861 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 316.1 | yes | Bloodvine Lens (19998, -4.60 DPS) [crafted]; Mask of the Unforgiven (13404, -6.05 DPS) [dungeon]; Ragefury Eyepatch (11735, -11.09 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 168.9 | yes | Medallion of the Dawn (22659, -1.59 DPS) [quest]; Beads of Ogre Might (22150, -5.40 DPS) [quest]; Choker of the Shifting Sands (21505, -6.79 DPS) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 189.8 | yes | Lieutenant Commander's Leather Shoulders (23313, -0.47 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.47 DPS) [pvp]; Champion's Leather Shoulders (23258, -4.78 DPS, sim-verified) [vendor] |
| back | Cloak of Veiled Shadows (21406) | Cloak of Veiled Shadows [quest] | 69.1 | yes | Earthweave Cloak (21187, -0.22 DPS) [quest]; Deathguard's Cloak (20068, -1.51 DPS) [rep]; Chromatic Cloak (18509, -3.17 DPS, sim-verified) [crafted] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 354.2 | yes | Stormshroud Armor (15056, -6.64 DPS) [crafted]; Zandalar Madcap's Tunic (19834, -7.58 DPS, sim-verified) [quest]; Deathdealer's Vest (21364, -7.68 DPS) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 151.4 | yes | Primal Batskin Bracers (19687, -3.44 DPS, sim-verified) [crafted]; Rockfury Bracers (21186, -5.75 DPS) [quest]; Marshal's Leather Armsplints (16460, -6.69 DPS) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 225.1 | yes | Devilsaur Gauntlets (15063, -4.39 DPS) [crafted]; Marshal's Leather Handgrips (16454, -4.39 DPS) [vendor]; Stormshroud Gloves (21278, -8.40 DPS, sim-verified) [crafted] |
| waist | Bonescythe Waistguard (22482) | Bonescythe Waistguard [quest] | 148.6 | yes | Defiler's Leather Girdle (20193, -0.72 DPS) [rep]; Belt of the Archmage (18405, -1.79 DPS) [crafted]; Defiler's Leather Girdle (20190, -3.60 DPS, sim-verified) [rep] |
| legs | Marshal's Leather Leggings (16456) (or General's Leather Legguards (16564), Marshal's Leather Leggings (231548), General's Leather Legguards (231554)) | Captain Dirgehammer [vendor] | 240.8 | yes | General's Leather Legguards (16564, +0.00 DPS, sim-verified) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 223.1 | yes | Deathdealer's Boots (21359, -3.36 DPS, sim-verified) [quest]; Bloodvine Boots (19684, -9.59 DPS) [crafted]; Shadowcraft Boots (16711, -9.67 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 175.1 | yes | Band of the Penitent (13217, -3.21 DPS) [quest]; Dragonslayer's Signet (18403, -3.21 DPS) [quest]; Ring of Entropy (18543, -3.21 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 159.1 | yes | Dragonslayer's Signet (18403, -2.35 DPS) [quest]; Ring of Entropy (18543, -2.35 DPS) [world]; Band of the Penitent (13217, -3.80 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 72.8 | yes | Tidal Charm (1404, -3.90 DPS) [vendor]; Guardian Talisman (1490, -3.90 DPS) [quest]; Ankh of Life (1713, -3.90 DPS) [dungeon] |
| trinket2 | Onyxia Blood Talisman (18406) | For All To See [quest] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon] |
| main_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 1010.3 | yes | High Warlord's Blade (234552, -0.16 DPS) [pvp]; High Warlord's Bludgeon (234555, -0.16 DPS) [pvp]; High Warlord's Right Claw (234557, -0.16 DPS) [pvp] |
| off_hand | Grand Marshal's Swiftblade (234579) | Rank 18 [pvp] | 1010.3 | yes | High Warlord's Left Claw (234558, -0.16 DPS) [pvp]; Grand Marshal's Left Hand Blade (234584, -0.16 DPS) [pvp]; Grand Marshal's Left Hand Blade (18847, -1.84 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 115.1 | yes | Fahrad's Reloading Repeater (22347, -3.51 DPS) [quest]; Core Marksman Rifle (18282, -3.81 DPS) [crafted]; Blackcrow (12651, -3.82 DPS) [dungeon] |

**New at 60:** head: Bonescythe Helmet; neck: Onyxia Tooth Pendant; shoulder: Bonescythe Pauldrons; back: Cloak of Veiled Shadows; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Bonescythe Waistguard; legs: Marshal's Leather Leggings; feet: Bonescythe Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket2: Onyxia Blood Talisman; main_hand: High Warlord's Quickblade; off_hand: Grand Marshal's Swiftblade; ranged: The Purifier

No-known-source sample (15 of 1861, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

