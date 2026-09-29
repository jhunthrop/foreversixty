# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.5. Weights run: 1.1s. Verify run: 1.2s. 360 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.009 ± 0.003, crit=0.390 ± 0.062, hit=not significant (0.899 ± 0.286), melee_haste=not significant (0.636 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Flying Tiger Goggles (4368, -8.1) [crafted]; Shadow Goggles (4373, -8.1) [crafted]; Lucky Fishing Hat (19972, -8.1) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 | yes | Tarnished Locket (279870, -6.1) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -5.0) [crafted]; Reinforced Woolen Shoulders (4315, -5.0) [crafted]; Forest Leather Mantle (4709, -5.0) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.1) [quest]; Cape of the Brotherhood (5193, -1.0) [dungeon]; Sentry Cloak (2059, -2.0) [dungeon] |
| chest | Tunic of Westfall (2041) | Quests [quest] | 11.1 | yes | Brawler's Leather Armor (252490, -4.0) [crafted]; Trapper's Leather Armor (252491, -4.0) [crafted]; Dark Leather Tunic (2317, -5.0) [crafted] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 5.0 | yes | Wolf Bracers (4794, -1.0) [vendor]; Bravo's Armbands (270015, -1.0) [quest]; Ratchet Wristwraps (274742, -2.0) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.0) [dungeon]; Fletcher's Gloves (7348, -0.6) [crafted]; Forest Leather Gloves (3058, -2.0) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -13.0) [crafted]; Dusty Belt (279897, -13.0) [quest]; Guardsman Belt (3429, -14.0) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.0) [crafted]; Trapper's Leather Pants (252501, +0.0) [crafted]; Bluegill Breeches (3022, -1.0) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.1 | yes | Blackened Defias Boots (10402, -2.0) [dungeon]; Footpads of the Fang (10411, -2.0) [dungeon]; Dark Leather Boots (2315, -3.0) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Lavishly Jeweled Ring (1156, -4.0) [dungeon]; The 1 Ring (8350, -5.0) [world]; Minor Channeling Ring (1449, -6.1) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 4.0 | yes | Lavishly Jeweled Ring (1156, -2.0) [dungeon]; The 1 Ring (8350, -3.0) [world]; Minor Channeling Ring (1449, -4.0) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -20.3) [dungeon]; Diamond Hammer (2194, -22.1) [dungeon]; Barrens Basher (274744, -25.6) [vendor] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -229.6) [quest]; Pulsating Hydra Heart (5183, -229.6) [world]; Tear of Grief (5611, -229.6) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Deadly Blunderbuss (4369, -2.0) [crafted]; Light Bow (4576, -2.0) [dungeon]; Owlsight Rifle (15205, -2.0) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 360, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (night-elf, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.7. Weights run: 1.0s. Verify run: 1.6s. 691 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.033 ± 0.019, crit=0.465 ± 0.062, hit=2.484 ± 0.367, melee_haste=not significant (0.623 ± 0.730)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.3 | yes | Tribal Worg Helm (6204, -2.1) [world]; Brawler's Leather Hood (252504, -2.1) [crafted]; Holy Shroud (2721, -10.3) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -5.7) [rep]; Sentinel's Medallion (20444, -7.8) [rep]; Pendant of Myzrael (4614, -14.0) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 11.4 | yes | Mantle of Thieves (2264, -1.0) [dungeon]; Dark Leather Shoulders (4252, -4.1) [crafted]; Insignia Mantle (4721, -4.1) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Cloak of Night (4447, -3.8) [world]; Fenrus' Hide (6340, -3.8) [dungeon]; Glowing Lizardscale Cloak (6449, -3.8) [dungeon] |
| chest | Raptorbane Armor (3566) | Quests [quest] | 16.0 | yes | Dusky Leather Armor (7374, -1.5) [crafted]; Tunic of Westfall (2041, -4.6) [quest]; Green Leather Armor (4255, -7.7) [crafted] |
| wrist | Cultist's Armguards (270032) | Quests [quest] | 10.0 | yes | Jurassic Wristguards (6198, -3.8) [world]; Insignia Bracers (6410, -3.8) [dungeon]; Madwolf Bracers (897, -4.8) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Pilferer's Gloves (7358, -7.7) [crafted]; Fletcher's Gloves (7348, -9.5) [crafted]; Wolfclaw Gloves (1978, -9.8) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.0) [rep]; Blackened Defias Belt (10403, -6.0) [dungeon]; Skulker's Leather Belt (252520, -14.7) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -12.6) [crafted]; Insignia Leggings (4054, -16.7) [dungeon]; Leggings of the Fang (10410, -16.7) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.3 | yes | Insignia Boots (4055, +0.0) [dungeon]; Highlander's Mail Greaves (20123, +0.0) [vendor]; Lancer Boots (6752, -1.0) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.3 | yes | Monkey Ring (6748, -2.1) [quest]; Ring of Precision (1491, -3.1) [dungeon]; Protector's Band (19517, -3.1) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -1.8) [quest]; Ring of Precision (1491, -2.8) [dungeon]; Protector's Band (19517, -2.8) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, -1.1) [dungeon]; Electrocutioner Leg (9446, -13.2) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -16.0) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -318.3) [quest]; Rod of Molten Fire (2565, -318.3) [dungeon]; Eye of Paleth (2943, -318.3) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Moonsight Rifle (4383, -1.5) [crafted]; Precision Bow (217315, -1.5) [quest]; Silver Star (3463, -3.8) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 88.9. Weights run: 1.1s. Verify run: 1.5s. 958 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.962 ± 0.146, hit=4.384 ± 0.987, melee_haste=not significant (-0.246 ± 1.799)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.1 | yes | White Bandit Mask (10008, -1.0) [crafted]; Hawkeye's Helm (14591, -1.0) [world]; Brawler's Leather Helm (252512, -2.0) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, -2.9) [rep]; Sentinel's Medallion (19541, -5.9) [rep]; Sentinel's Medallion (20444, -7.9) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Mantle of Thieves (2264, -13.0) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, -1.9) [dungeon]; Parachute Cloak (10518, -1.9) [crafted]; Yeti Fur Cloak (2805, -3.9) [quest] |
| chest | Raptorbane Armor (3566) | Quests [quest] | 16.0 | yes | Nightscape Tunic (8175, -0.8) [crafted]; Dusky Leather Armor (7374, -1.9) [crafted]; Hawkeye's Tunic (14592, -3.9) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Cultist's Armguards (270032, -10.0) [quest]; Imperial Leather Bracers (4061, -11.9) [dungeon]; Dusky Bracers (7378, -11.9) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 33.5 | yes | Heavy Earthen Gloves (7359, -17.5) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Chain Girdle (20090, -6.0) [rep]; Highlander's Leather Girdle (20117, -6.0) [rep]; Blackened Defias Belt (10403, -12.0) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, -4.8) [dungeon]; Triprunner Dungarees (9624, -7.8) [quest]; Hawkeye's Breeches (14595, -11.9) [world] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 13.1 | yes | Imperial Leather Boots (6431, -2.0) [dungeon]; Dusky Boots (7390, -2.0) [crafted]; Skulker's Leather Shoes (252531, -2.0) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -2.9) [dungeon]; Insurgent's Band (272067, -3.0) [vendor]; Protector's Band (19515, -3.9) [rep] |
| finger2 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 10.1 | yes | Ironspine's Eye (7686, -1.0) [dungeon]; Protector's Band (19515, -2.0) [rep]; Disengagement Ring (276202, -2.0) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -27.3) [vendor]; Sword of Serenity (6829, -40.0) [quest]; Hand of Righteousness (7721, -42.7) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Quests [quest] | 442.0 | yes | Grayson's Torch (1172, -442.0) [quest]; Rod of Molten Fire (2565, -442.0) [dungeon]; Eye of Paleth (2943, -442.0) [quest] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 13.2 | yes | Precision Bow (217315, +0.0) [quest]; Booty Bay Bruiser's Buckshot (274748, -4.2) [vendor]; Master Hunter's Bow (17686, -7.1) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Moonsight Rifle

No-known-source sample (15 of 958, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 105.9. Weights run: 1.1s. Verify run: 1.3s. 1261 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.239 ± 0.059, crit=6.333 ± 0.367, hit=not significant (5.215 ± 1.384), melee_haste=not significant (-1.529 ± 2.421)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 88.7 | yes | Helm of Fire (8348, -67.6) [crafted]; Lordrec Helmet (10741, -68.8) [quest]; Sprightring Helm (17776, -70.1) [quest] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 14.9 | yes | Ghostshard Talisman (7731, -0.9) [dungeon]; Sentinel's Medallion (19540, -1.2) [rep]; Sentinel's Medallion (19541, -5.0) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.6 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Penance Spaulders (11963, -12.0) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 13.6 | yes | Nightscape Cloak (8195, -1.2) [crafted]; Pridelord Cape (14673, -2.5) [dungeon]; Wolfmaster Cape (6314, -3.6) [dungeon] |
| chest | Blazewind Breastplate (11193) | Quests [quest] | 28.5 | yes | Warbear Harness (15064, -6.2) [crafted]; Charred Leather Tunic (19127, -6.2) [quest]; Nightscape Tunic (8175, -9.9) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, -1.4) [world]; Wicked Leather Bracers (15084, -6.4) [crafted]; Pridelord Bands (14672, -7.6) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 108.7 | yes | Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted]; Heavy Earthen Gloves (7359, -92.7) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 108.7 | yes | Highlander's Cloth Girdle (20097, -20.0) [rep]; Highlander's Lizardhide Girdle (20103, -20.0) [rep]; Highlander's Leather Girdle (20116, -78.7) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 177.3 | yes | Basilisk Hide Pants (1718, -151.3) [dungeon]; Ferine Leggings (6690, -151.3) [dungeon]; Keeper's Woolies (14668, -153.8) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.8 | yes | Sandstalker Ankleguards (12470, -3.7) [dungeon]; Swampwalker Boots (2276, -8.7) [dungeon]; Skulker's Leather Boots (252469, -8.7) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 72.1 | yes | Insurgent's Band (272065, -57.1) [vendor]; Ring of the Underwood (2951, -59.8) [dungeon]; Insurgent's Band (272066, -60.1) [vendor] |
| finger2 | Masons Fraternity Ring (9533) | Quests [quest] | 17.3 | yes | Insurgent's Band (272065, -2.3) [vendor]; Ring of the Underwood (2951, -5.0) [dungeon]; Insurgent's Band (272066, -5.3) [vendor] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 581.9 | yes | Might of Hakkar (10838, -54.2) [world]; Julie's Dagger (6660, -70.3) [dungeon]; Lifeforce Dirk (10750, -78.7) [quest] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -71.1) [dungeon]; Thermotastic Egg Timer (9644, -549.6) [quest]; Grayson's Torch (1172, -553.3) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.3 | yes | Moonsight Rifle (4383, -1.7) [crafted]; Precision Bow (217315, -1.7) [quest]; Houndmaster's Bow (11628, -5.3) [dungeon] |

**New at 50:** head: Eye of Theradras; neck: Sentinel's Medallion; back: Serpentskin Cloak; chest: Blazewind Breastplate; waist: Highlander's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Smoking Heart of the Mountain; trinket2: Talisman of Arathor; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1261, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 370.7. Weights run: 1.1s. Verify run: 1.8s. 1770 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Quests [quest] | 316.1 | yes | Ragefury Eyepatch (11735, -85.9) [dungeon]; Bloodvine Lens (19998, -85.9) [crafted]; Mask of the Unforgiven (13404, -113.0) [dungeon] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 256.2 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Onyxia Tooth Pendant (18404, -87.4) [quest]; Barbed Choker (21664, -97.1) [raid] |
| shoulder | Bonescythe Pauldrons (22479) | Quests [quest] | 189.8 | yes | Lieutenant Commander's Leather Shoulders (227054, -8.7) [pvp]; Champion's Leather Shoulders (227056, -8.7) [pvp]; Champion's Leather Shoulders (23258, -8.7) [vendor] |
| back | Cloak of Veiled Shadows (21406) | Quests [quest] | 69.1 | yes | Chromatic Cloak (18509, +46.0) [crafted]; Earthweave Cloak (21187, -4.2) [quest]; Cloak of the Honor Guard (20073, -28.1) [rep] |
| chest | Bonescythe Breastplate (22476) | Quests [quest] | 354.2 | yes | Zandalar Madcap's Tunic (19834, -80.0) [quest]; Stormshroud Armor (15056, -124.0) [crafted]; Deathdealer's Vest (21364, -143.4) [quest] |
| wrist | Bonescythe Bracers (22483) | Quests [quest] | 151.4 | yes | Qiraji Execution Bracers (21602, -85.1) [raid]; Primal Batskin Bracers (19687, -87.9) [crafted]; Rockfury Bracers (21186, -107.5) [quest] |
| hands | Bonescythe Gauntlets (22481) | Quests [quest] | 225.1 | yes | Stormshroud Gloves (21278, -66.0) [crafted]; Devilsaur Gauntlets (15063, -82.0) [crafted]; Marshal's Leather Handgrips (16454, -82.0) [vendor] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 223.1 | yes | Highlander's Leather Girdle (20045, -74.0) [rep]; Bonescythe Waistguard (22482, -74.5) [quest]; Highlander's Leather Girdle (20115, -88.0) [rep] |
| legs | Marshal's Leather Leggings (231548) (or General's Leather Legguards (231554)) | Rank 16 [pvp] | 240.8 | yes | General's Leather Legguards (231554, +0.0) [pvp]; Marshal's Leather Leggings (16456, -0.0) [vendor]; General's Leather Legguards (16564, -0.0) [vendor] |
| feet | Bonescythe Sabatons (22480) | Quests [quest] | 223.1 | yes | Deathdealer's Boots (21359, -142.8) [quest]; Bloodvine Boots (19684, -179.1) [crafted]; Shadowcraft Boots (16711, -180.6) [dungeon] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 211.1 | yes | Band of Earthen Might (21182, -52.0) [quest]; Seal of the Damned (23025, -52.0) [raid]; Ring of the Qiraji Fury (21677, -56.0) [raid] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 175.1 | yes | Band of Earthen Might (21182, -16.0) [quest]; Seal of the Damned (23025, -16.0) [raid]; Ring of the Qiraji Fury (21677, -20.0) [raid] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +166.2) [raid]; Drake Fang Talisman (19406, +79.9) [raid]; Neltharion's Tear (19379, +23.9) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 159.1 | yes | Eye of Diminution (23001, +71.1) [raid]; Drake Fang Talisman (19406, -15.1) [raid]; Neltharion's Tear (19379, -71.1) [raid] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Quests [quest] | 762.3 | yes | Blessed Qiraji Pugio (21244, +263.1) [quest]; High Warlord's Quickblade (234553, +248.0) [pvp]; Grand Marshal's Swiftblade (234579, +248.0) [pvp] |
| off_hand | The Hungering Cold (23577) | Naxxramas [raid] | 1022.0 | yes | High Warlord's Left Claw (234558, -14.8) [pvp]; Grand Marshal's Left Hand Blade (234584, -14.8) [pvp]; Grand Marshal's Left Hand Blade (18847, -46.2) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 133.1 | yes | The Purifier (22656, -18.0) [quest]; Fahrad's Reloading Repeater (22347, -83.6) [quest]; Core Marksman Rifle (18282, -89.1) [crafted] |

**New at 60:** head: Bonescythe Helmet; neck: Stormrage's Talisman of Seething; shoulder: Bonescythe Pauldrons; back: Cloak of Veiled Shadows; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; legs: Marshal's Leather Leggings; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Don Julio's Band; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: The Hungering Cold; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1770, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.0. Weights run: 1.1s. Verify run: 1.2s. 355 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.009 ± 0.003, crit=0.390 ± 0.062, hit=not significant (0.899 ± 0.286), melee_haste=not significant (0.636 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Flying Tiger Goggles (4368, -8.1) [crafted]; Shadow Goggles (4373, -8.1) [crafted]; Lucky Fishing Hat (19972, -8.1) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 | yes | Tarnished Locket (279870, -6.1) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -5.0) [crafted]; Reinforced Woolen Shoulders (4315, -5.0) [crafted]; Forest Leather Mantle (4709, -5.0) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.1) [quest]; Cape of the Brotherhood (5193, -1.0) [dungeon]; Sentry Cloak (2059, -2.0) [dungeon] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 | yes | Trapper's Leather Armor (252491, +0.0) [crafted]; Dark Leather Tunic (2317, -1.0) [crafted]; Heckler's Hide (286536, -2.0) [world] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 5.0 | yes | Wolf Bracers (4794, -1.0) [vendor]; Bravo's Armbands (270015, -1.0) [quest]; Ratchet Wristwraps (274742, -2.0) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.0) [dungeon]; Fletcher's Gloves (7348, -0.6) [crafted]; Forest Leather Gloves (3058, -2.0) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -13.0) [crafted]; Dusty Belt (279897, -13.0) [quest]; Guardsman Belt (3429, -14.0) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.0) [crafted]; Trapper's Leather Pants (252501, +0.0) [crafted]; Bluegill Breeches (3022, -1.0) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.1 | yes | Blackened Defias Boots (10402, -2.0) [dungeon]; Footpads of the Fang (10411, -2.0) [dungeon]; Dark Leather Boots (2315, -3.0) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Bounty Hunter's Ring (5351, -3.0) [quest]; Lavishly Jeweled Ring (1156, -4.0) [dungeon]; The 1 Ring (8350, -5.0) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 4.0 | yes | Bounty Hunter's Ring (5351, -1.0) [quest]; Lavishly Jeweled Ring (1156, -2.0) [dungeon]; The 1 Ring (8350, -3.0) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -20.3) [dungeon]; Diamond Hammer (2194, -22.1) [dungeon]; Wingblade (6504, -24.4) [quest] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -229.6) [quest]; Nightglow Concoction (3451, -229.6) [quest]; Pulsating Hydra Heart (5183, -229.6) [world] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Deadly Blunderbuss (4369, -2.0) [crafted]; Light Bow (4576, -2.0) [dungeon]; Privateer Musket (5309, -2.0) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 355, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (troll, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.3. Weights run: 1.0s. Verify run: 1.7s. 687 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.033 ± 0.019, crit=0.465 ± 0.062, hit=2.484 ± 0.367, melee_haste=not significant (0.623 ± 0.730)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.3 | yes | Tribal Worg Helm (6204, -2.1) [world]; Brawler's Leather Hood (252504, -2.1) [crafted]; Holy Shroud (2721, -10.3) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -5.7) [rep]; Scout's Medallion (20442, -7.8) [rep]; Pendant of Myzrael (4614, -14.0) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 11.4 | yes | Mantle of Thieves (2264, -1.0) [dungeon]; Dark Leather Shoulders (4252, -4.1) [crafted]; Insignia Mantle (4721, -4.1) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.0) [quest]; Cloak of Night (4447, -3.8) [world]; Fenrus' Hide (6340, -3.8) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.5 | yes | Panther Armor (6670, -5.2) [quest]; Green Leather Armor (4255, -6.2) [crafted]; Brawler's Leather Tunic (252508, -6.2) [crafted] |
| wrist | Cultist's Armguards (270032) | Quests [quest] | 10.0 | yes | Jurassic Wristguards (6198, -3.8) [world]; Insignia Bracers (6410, -3.8) [dungeon]; Madwolf Bracers (897, -4.8) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Pilferer's Gloves (7358, -7.7) [crafted]; Braced Handguards (6784, -8.8) [quest]; Fletcher's Gloves (7348, -9.5) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.0) [rep]; Blackened Defias Belt (10403, -6.0) [dungeon]; Deftkin Belt (16659, -7.9) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -12.6) [crafted]; Insignia Leggings (4054, -16.7) [dungeon]; Leggings of the Fang (10410, -16.7) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.3 | yes | Insignia Boots (4055, +0.0) [dungeon]; Warsong Boots (16977, +0.0) [quest]; Highlander's Mail Greaves (20123, +0.0) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.3 | yes | Monkey Ring (6748, -2.1) [quest]; Ring of Precision (1491, -3.1) [dungeon]; Legionnaire's Band (19513, -3.1) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -1.8) [quest]; Ring of Precision (1491, -2.8) [dungeon]; Legionnaire's Band (19513, -2.8) [rep] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, -1.1) [dungeon]; Electrocutioner Leg (9446, -13.2) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -16.0) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -318.3) [quest]; Rod of Molten Fire (2565, -318.3) [dungeon]; Nightglow Concoction (3451, -318.3) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Moonsight Rifle (4383, -1.5) [crafted]; Precision Bow (217315, -1.5) [quest]; Silver Star (3463, -3.8) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 687, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 87.9. Weights run: 1.1s. Verify run: 1.6s. 954 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.962 ± 0.146, hit=4.384 ± 0.987, melee_haste=not significant (-0.246 ± 1.799)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.1 | yes | White Bandit Mask (10008, -1.0) [crafted]; Hawkeye's Helm (14591, -1.0) [world]; Spirit Hunter Headdress (6720, -2.0) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, -2.9) [rep]; Scout's Medallion (19537, -5.9) [rep]; Scout's Medallion (20442, -7.9) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Mantle of Thieves (2264, -13.0) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.0) [quest]; Imperial Cloak (6432, -1.9) [dungeon]; Parachute Cloak (10518, -1.9) [crafted] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 15.2 | yes | Dusky Leather Armor (7374, -1.0) [crafted]; Hawkeye's Tunic (14592, -3.0) [world]; Panther Armor (6670, -6.1) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Cultist's Armguards (270032, -10.0) [quest]; Imperial Leather Bracers (4061, -11.9) [dungeon]; Dusky Bracers (7378, -11.9) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 33.5 | yes | Heavy Earthen Gloves (7359, -17.5) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Chain Girdle (20152, -6.0) [rep]; Defiler's Leather Girdle (20191, -6.0) [rep]; Blackened Defias Belt (10403, -12.0) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, -4.8) [dungeon]; Triprunner Dungarees (9624, -7.8) [quest]; Hawkeye's Breeches (14595, -11.9) [world] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 13.1 | yes | Imperial Leather Boots (6431, -2.0) [dungeon]; Dusky Boots (7390, -2.0) [crafted]; Skulker's Leather Shoes (252531, -2.0) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -2.9) [dungeon]; Insurgent's Band (272067, -3.0) [vendor]; Legionnaire's Band (19512, -3.9) [rep] |
| finger2 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 10.1 | yes | Ironspine's Eye (7686, -1.0) [dungeon]; Legionnaire's Band (19512, -2.0) [rep]; Disengagement Ring (276202, -2.0) [vendor] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -27.3) [vendor]; Hand of Righteousness (7721, -42.7) [dungeon]; Darkspear Skirmisher's Bludgeon (272093, -43.0) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Quests [quest] | 442.0 | yes | Grayson's Torch (1172, -442.0) [quest]; Rod of Molten Fire (2565, -442.0) [dungeon]; Nightglow Concoction (3451, -442.0) [quest] |
| ranged | Moonsight Rifle (4383) (or Precision Bow (217315)) | Engineering [crafted] | 13.2 | yes | Precision Bow (217315, +0.0) [quest]; Booty Bay Bruiser's Buckshot (274748, -4.2) [vendor]; Master Hunter's Bow (17686, -7.1) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Moonsight Rifle

No-known-source sample (15 of 954, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 104.8. Weights run: 1.1s. Verify run: 1.3s. 1257 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.239 ± 0.059, crit=6.333 ± 0.367, hit=not significant (5.215 ± 1.384), melee_haste=not significant (-1.529 ± 2.421)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 88.7 | yes | Helm of Fire (8348, -67.6) [crafted]; Sprightring Helm (17776, -70.1) [quest]; Nightscape Headband (8176, -73.8) [crafted] |
| neck | Scout's Medallion (19535) | Warsong Outriders [rep] | 14.9 | yes | Ghostshard Talisman (7731, -0.9) [dungeon]; Scout's Medallion (19536, -1.2) [rep]; Woven Ivy Necklace (19159, -3.7) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.6 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Penance Spaulders (11963, -12.0) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 13.6 | yes | Nightscape Cloak (8195, -1.2) [crafted]; Pridelord Cape (14673, -2.5) [dungeon]; Wolfmaster Cape (6314, -3.6) [dungeon] |
| chest | Blazewind Breastplate (11193) | Quests [quest] | 28.5 | yes | Warbear Harness (15064, -6.2) [crafted]; Charred Leather Tunic (19127, -6.2) [quest]; Nightscape Tunic (8175, -9.9) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, -1.4) [world]; Wicked Leather Bracers (15084, -6.4) [crafted]; Pridelord Bands (14672, -7.6) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 108.7 | yes | Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted]; Heavy Earthen Gloves (7359, -92.7) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 108.7 | yes | Defiler's Cloth Girdle (20165, -20.0) [rep]; Defiler's Lizardhide Girdle (20174, -20.0) [rep]; Defiler's Leather Girdle (20192, -78.7) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 177.3 | yes | Basilisk Hide Pants (1718, -151.3) [dungeon]; Ferine Leggings (6690, -151.3) [dungeon]; Keeper's Woolies (14668, -153.8) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.8 | yes | Sandstalker Ankleguards (12470, -3.7) [dungeon]; Swampwalker Boots (2276, -8.7) [dungeon]; Skulker's Leather Boots (252469, -8.7) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 72.1 | yes | Masons Fraternity Ring (9533, -54.8) [quest]; Insurgent's Band (272065, -57.1) [vendor]; Ring of the Underwood (2951, -59.8) [dungeon] |
| finger2 | White Bone Band (11862) | Quests [quest] | 24.0 | yes | Masons Fraternity Ring (9533, -6.7) [quest]; Insurgent's Band (272065, -9.0) [vendor]; Ring of the Underwood (2951, -11.6) [dungeon] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, +78.5) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Defiler's Talisman (21115) | The Defilers [rep] | 0.0 | yes | Rune of the Guard Captain (19120, +78.5) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 581.9 | yes | Might of Hakkar (10838, -54.2) [world]; Julie's Dagger (6660, -70.3) [dungeon]; Lifeforce Dirk (10750, -78.7) [quest] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -71.1) [dungeon]; White Bone Shredder (11863, -109.6) [quest]; Thermotastic Egg Timer (9644, -549.6) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | Blackrock Depths: Anvilrage Overseer [dungeon] | 17.3 | yes | Moonsight Rifle (4383, -1.7) [crafted]; Precision Bow (217315, -1.7) [quest]; Houndmaster's Bow (11628, -5.3) [dungeon] |

**New at 50:** head: Eye of Theradras; neck: Scout's Medallion; back: Serpentskin Cloak; chest: Blazewind Breastplate; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Smoking Heart of the Mountain; trinket2: Defiler's Talisman; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1257, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 360.1. Weights run: 1.1s. Verify run: 1.9s. 1765 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Quests [quest] | 316.1 | yes | Ragefury Eyepatch (11735, -85.9) [dungeon]; Bloodvine Lens (19998, -85.9) [crafted]; Mask of the Unforgiven (13404, -113.0) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 18.2 | yes | Stormrage's Talisman of Seething (23053, +238.1) [raid]; Gem of Trapped Innocents (23057, +212.1) [raid]; Onyxia Tooth Pendant (18404, +150.7) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Quests [quest] | 189.8 | yes | Champion's Leather Shoulders (23258, -8.7) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -8.7) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -8.7) [pvp] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 115.1 | yes | Cloak of Veiled Shadows (21406, -46.0) [quest]; Earthweave Cloak (21187, -50.2) [quest]; Deathguard's Cloak (20068, -74.1) [rep] |
| chest | Bonescythe Breastplate (22476) | Quests [quest] | 354.2 | yes | Zandalar Madcap's Tunic (19834, -80.0) [quest]; Stormshroud Armor (15056, -124.0) [crafted]; Deathdealer's Vest (21364, -143.4) [quest] |
| wrist | Bonescythe Bracers (22483) | Quests [quest] | 151.4 | yes | Qiraji Execution Bracers (21602, -85.1) [raid]; Primal Batskin Bracers (19687, -87.9) [crafted]; Rockfury Bracers (21186, -107.5) [quest] |
| hands | Bonescythe Gauntlets (22481) | Quests [quest] | 225.1 | yes | Stormshroud Gloves (21278, -66.0) [crafted]; Devilsaur Gauntlets (15063, -82.0) [crafted]; Marshal's Leather Handgrips (16454, -82.0) [vendor] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 223.1 | yes | Defiler's Leather Girdle (20190, -74.0) [rep]; Bonescythe Waistguard (22482, -74.5) [quest]; Defiler's Leather Girdle (20193, -88.0) [rep] |
| legs | Marshal's Leather Leggings (231548) (or General's Leather Legguards (231554)) | Rank 16 [pvp] | 240.8 | yes | General's Leather Legguards (231554, +0.0) [pvp]; Marshal's Leather Leggings (16456, -0.0) [vendor]; General's Leather Legguards (16564, -0.0) [vendor] |
| feet | Bonescythe Sabatons (22480) | Quests [quest] | 223.1 | yes | Deathdealer's Boots (21359, -142.8) [quest]; Bloodvine Boots (19684, -179.1) [crafted]; Shadowcraft Boots (16711, -180.6) [dungeon] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 211.1 | yes | Band of Earthen Might (21182, -52.0) [quest]; Seal of the Damned (23025, -52.0) [raid]; Ring of the Qiraji Fury (21677, -56.0) [raid] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 175.1 | yes | Band of Earthen Might (21182, -16.0) [quest]; Seal of the Damned (23025, -16.0) [raid]; Ring of the Qiraji Fury (21677, -20.0) [raid] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +166.2) [raid]; Drake Fang Talisman (19406, +79.9) [raid]; Neltharion's Tear (19379, +23.9) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 159.1 | yes | Eye of Diminution (23001, +71.1) [raid]; Drake Fang Talisman (19406, -15.1) [raid]; Neltharion's Tear (19379, -71.1) [raid] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Quests [quest] | 762.3 | yes | Blessed Qiraji Pugio (21244, +263.1) [quest]; High Warlord's Quickblade (234553, +248.0) [pvp]; Grand Marshal's Swiftblade (234579, +248.0) [pvp] |
| off_hand | The Hungering Cold (23577) | Naxxramas [raid] | 1022.0 | yes | High Warlord's Left Claw (234558, -14.8) [pvp]; Grand Marshal's Left Hand Blade (234584, -14.8) [pvp]; Grand Marshal's Left Hand Blade (18847, -46.2) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 133.1 | yes | The Purifier (22656, -18.0) [quest]; Fahrad's Reloading Repeater (22347, -83.6) [quest]; Core Marksman Rifle (18282, -89.1) [crafted] |

**New at 60:** head: Bonescythe Helmet; neck: Blazefury Medallion; shoulder: Bonescythe Pauldrons; back: Chromatic Cloak; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; legs: Marshal's Leather Leggings; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Don Julio's Band; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: The Hungering Cold; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1765, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

