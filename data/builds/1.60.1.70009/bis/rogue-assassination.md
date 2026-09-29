# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (night-elf, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.8. Weights run: 1.1s. Verify run: 1.3s. 360 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.028, crit=2.157 ± 0.106, hit=1.402 ± 0.268, melee_haste=not significant (1.169 ± 0.749)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.8 | yes | Flying Tiger Goggles (4368, -8.8) [crafted]; Shadow Goggles (4373, -8.8) [crafted]; Lucky Fishing Hat (19972, -8.8) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.6 | yes | Tarnished Locket (279870, -6.6) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.5 | yes | Double-Stitched Woolen Shoulders (4314, -5.5) [crafted]; Reinforced Woolen Shoulders (4315, -5.5) [crafted]; Forest Leather Mantle (4709, -5.5) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.6 | yes | Catacomb Cloak (279899, -0.6) [quest]; Cape of the Brotherhood (5193, -1.1) [dungeon]; Sentry Cloak (2059, -2.2) [dungeon] |
| chest | Tunic of Westfall (2041) | Quests [quest] | 12.1 | yes | Brawler's Leather Armor (252490, -4.4) [crafted]; Trapper's Leather Armor (252491, -4.4) [crafted]; Dark Leather Tunic (2317, -5.5) [crafted] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 5.5 | yes | Wolf Bracers (4794, -1.1) [vendor]; Bravo's Armbands (270015, -1.1) [quest]; Ratchet Wristwraps (274742, -2.2) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 30.2 | yes | Serpent Gloves (5970, -23.6) [dungeon]; Gloves of the Fang (10413, -23.6) [dungeon]; Forest Leather Gloves (3058, -25.8) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -12.5) [crafted]; Dusty Belt (279897, -12.5) [quest]; Guardsman Belt (3429, -13.6) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.9 | yes | Brawler's Leather Pants (252500, +0.0) [crafted]; Trapper's Leather Pants (252501, +0.0) [crafted]; Bluegill Breeches (3022, -1.1) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.8 | yes | Blackened Defias Boots (10402, -2.2) [dungeon]; Footpads of the Fang (10411, -2.2) [dungeon]; Dark Leather Boots (2315, -3.3) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.6 | yes | Lavishly Jeweled Ring (1156, -4.4) [dungeon]; The 1 Ring (8350, -5.5) [world]; Minor Channeling Ring (1449, -6.6) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 4.4 | yes | Lavishly Jeweled Ring (1156, -2.2) [dungeon]; The 1 Ring (8350, -3.3) [world]; Minor Channeling Ring (1449, -4.4) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -20.3) [dungeon]; Diamond Hammer (2194, -22.1) [dungeon]; Barrens Basher (274744, -25.6) [vendor] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -229.6) [quest]; Pulsating Hydra Heart (5183, -229.6) [world]; Tear of Grief (5611, -229.6) [quest] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Deadly Blunderbuss (4369, -1.8) [crafted]; Light Bow (4576, -1.8) [dungeon]; Owlsight Rifle (15205, -1.8) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 360, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (night-elf, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 43.5. Weights run: 1.1s. Verify run: 1.7s. 691 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.099 ± 0.022, crit=3.270 ± 0.122, hit=1.653 ± 0.289, melee_haste=not significant (1.560 ± 0.801)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 | yes | Tribal Worg Helm (6204, -2.2) [world]; Brawler's Leather Hood (252504, -2.2) [crafted]; Holy Shroud (2721, -11.0) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -5.2) [rep]; Sentinel's Medallion (20444, -7.4) [rep]; Pendant of Myzrael (4614, -14.0) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 12.1 | yes | Mantle of Thieves (2264, -1.1) [dungeon]; Dark Leather Shoulders (4252, -4.4) [crafted]; Insignia Mantle (4721, -4.4) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Cloak of Night (4447, -3.4) [world]; Fenrus' Hide (6340, -3.4) [dungeon]; Glowing Lizardscale Cloak (6449, -3.4) [dungeon] |
| chest | Raptorbane Armor (3566) | Quests [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.6) [crafted]; Tunic of Westfall (2041, -3.9) [quest]; Green Leather Armor (4255, -7.2) [crafted] |
| wrist | Cultist's Armguards (270032) | Quests [quest] | 10.0 | yes | Jurassic Wristguards (6198, -3.4) [world]; Insignia Bracers (6410, -3.4) [dungeon]; Madwolf Bracers (897, -4.5) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 45.8 | yes | Heavy Earthen Gloves (7359, -29.8) [crafted]; Pilferer's Gloves (7358, -37.0) [crafted]; Wolfclaw Gloves (1978, -39.2) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.0) [rep]; Blackened Defias Belt (10403, -6.0) [dungeon]; Skulker's Leather Belt (252520, -14.1) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -11.7) [crafted]; Insignia Leggings (4054, -16.1) [dungeon]; Leggings of the Fang (10410, -16.1) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.8 | yes | Insignia Boots (4055, +0.0) [dungeon]; Highlander's Mail Greaves (20123, +0.0) [vendor]; Lancer Boots (6752, -1.1) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 | yes | Monkey Ring (6748, -2.2) [quest]; Ring of Precision (1491, -3.3) [dungeon]; Protector's Band (19517, -3.3) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -1.3) [quest]; Ring of Precision (1491, -2.4) [dungeon]; Protector's Band (19517, -2.4) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, -1.1) [dungeon]; Electrocutioner Leg (9446, -13.2) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -16.0) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -318.3) [quest]; Rod of Molten Fire (2565, -318.3) [dungeon]; Eye of Paleth (2943, -318.3) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Silver Star (3463, -3.5) [quest]; Moonsight Rifle (4383, -4.0) [crafted]; Precision Bow (217315, -4.0) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 85.4. Weights run: 1.1s. Verify run: 1.6s. 958 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.100 ± 0.014, crit=2.696 ± 0.072, hit=1.350 ± 0.103, melee_haste=2.063 ± 0.041

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 13.2 | yes | White Bandit Mask (10008, -1.1) [crafted]; Hawkeye's Helm (14591, -1.1) [world]; Brawler's Leather Helm (252512, -2.2) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, -1.9) [rep]; Sentinel's Medallion (19541, -5.2) [rep]; Sentinel's Medallion (20444, -7.4) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.1 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Mantle of Thieves (2264, -13.1) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, -1.2) [dungeon]; Parachute Cloak (10518, -1.2) [crafted]; Yeti Fur Cloak (2805, -3.4) [quest] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 16.5 | yes | Raptorbane Armor (3566, -0.5) [quest]; Dusky Leather Armor (7374, -1.1) [crafted]; Hawkeye's Tunic (14592, -3.3) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Cultist's Armguards (270032, -10.0) [quest]; Imperial Leather Bracers (4061, -11.2) [dungeon]; Dusky Bracers (7378, -11.2) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 57.7 | yes | Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted]; Heavy Earthen Gloves (7359, -41.7) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Chain Girdle (20090, -6.0) [rep]; Highlander's Leather Girdle (20117, -6.0) [rep]; Blackened Defias Belt (10403, -12.0) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, -2.9) [dungeon]; Triprunner Dungarees (9624, -6.2) [quest]; Hawkeye's Breeches (14595, -10.6) [world] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 14.3 | yes | Imperial Leather Boots (6431, -2.2) [dungeon]; Dusky Boots (7390, -2.2) [crafted]; Skulker's Leather Shoes (252531, -2.2) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -2.1) [dungeon]; Insurgent's Band (272067, -3.0) [vendor]; Protector's Band (19515, -3.2) [rep] |
| finger2 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 11.0 | yes | Ironspine's Eye (7686, -1.1) [dungeon]; Protector's Band (19515, -2.2) [rep]; Disengagement Ring (276202, -2.2) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -26.5) [vendor]; Sword of Serenity (6829, -40.0) [quest]; Hand of Righteousness (7721, -42.7) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Quests [quest] | 442.0 | yes | Grayson's Torch (1172, -442.0) [quest]; Rod of Molten Fire (2565, -442.0) [dungeon]; Eye of Paleth (2943, -442.0) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Master Hunter's Bow (17686, -2.4) [quest]; Silver Star (3463, -3.5) [quest]; Mithril Blunderbuss (10508, -3.5) [crafted] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 958, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 117.5. Weights run: 1.2s. Verify run: 1.6s. 1261 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.103 ± 0.015, crit=3.283 ± 0.085, hit=1.682 ± 0.129, melee_haste=2.580 ± 0.051

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 18.8 | yes | Eye of Theradras (17715, +27.2) [dungeon]; Lordrec Helmet (10741, -1.1) [quest]; Sprightring Helm (17776, -2.2) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19539, -0.8) [rep]; Sentinel's Medallion (19540, -1.9) [rep]; Sentinel's Medallion (19541, -5.2) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.1 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Penance Spaulders (11963, -12.0) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 12.1 | yes | Nightscape Cloak (8195, -1.1) [crafted]; Wolfmaster Cape (6314, -2.1) [dungeon]; Pridelord Cape (14673, -2.2) [dungeon] |
| chest | Blazewind Breastplate (11193) | Quests [quest] | 25.4 | yes | Warbear Harness (15064, -5.5) [crafted]; Charred Leather Tunic (19127, -5.5) [quest]; Nightscape Tunic (8175, -8.8) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, -3.5) [world]; Wicked Leather Bracers (15084, -7.9) [crafted]; Pridelord Bands (14672, -9.0) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 66.0 | yes | Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted]; Heavy Earthen Gloves (7359, -50.0) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 66.0 | yes | Highlander's Cloth Girdle (20097, -20.0) [rep]; Highlander's Lizardhide Girdle (20103, -20.0) [rep]; Highlander's Leather Girdle (20116, -36.0) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 91.9 | yes | Ferine Leggings (6690, -65.9) [dungeon]; Basilisk Hide Pants (1718, -68.8) [dungeon]; Keeper's Woolies (14668, -71.0) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 | yes | Sandstalker Ankleguards (12470, -3.3) [dungeon]; Swampwalker Boots (2276, -7.7) [dungeon]; Skulker's Leather Boots (252469, -7.7) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.8 | yes | Insurgent's Band (272065, -21.8) [vendor]; Insurgent's Band (272066, -24.8) [vendor]; Ring of the Underwood (2951, -25.8) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Quests [quest] | 15.4 | yes | Insurgent's Band (272065, -0.4) [vendor]; Insurgent's Band (272066, -3.4) [vendor]; Ring of the Underwood (2951, -4.4) [dungeon] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Inventor's Focal Sword (17719, -14.1) [dungeon]; Might of Hakkar (10838, -43.2) [world]; Lifeforce Dirk (10750, -50.1) [quest] |
| off_hand | Julie's Dagger (6660) | Blackrock Depths: Shadowforge Peasant [dungeon] | 511.6 | yes | Claw of Celebras (17738, -29.4) [dungeon]; Thermotastic Egg Timer (9644, -508.3) [quest]; Grayson's Torch (1172, -511.6) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | Blackrock Depths: Anvilrage Overseer [dungeon] | 15.4 | yes | Houndmaster's Bow (11628, -3.4) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -6.4) [vendor]; Guttbuster (13139, -6.6) [world] |

**New at 50:** head: Helm of Fire; back: Serpentskin Cloak; chest: Blazewind Breastplate; waist: Highlander's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Smoking Heart of the Mountain; trinket2: Talisman of Arathor; main_hand: Hammer of the Northern Wind; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1261, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 288.3. Weights run: 1.2s. Verify run: 1.9s. 1770 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=not significant (0.000 ± 0.000), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Quests [quest] | 149.4 | yes | Ragefury Eyepatch (11735, -33.9) [dungeon]; Bloodvine Lens (19998, -33.9) [crafted]; Champion's Leather Helm (23257, -55.7) [vendor] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 141.5 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Barbed Choker (21664, -39.7) [raid]; Medallion of the Dawn (22659, -59.7) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Quests [quest] | 82.6 | yes | Champion's Leather Shoulders (23258, -2.9) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -2.9) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -2.9) [pvp] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 57.7 | yes | Cloak of the Honor Guard (20073, -18.1) [rep]; Cape of the Black Baron (13340, -20.8) [dungeon]; Cloak of the Fallen God (21710, -28.3) [quest] |
| chest | Bonescythe Breastplate (22476) | Quests [quest] | 195.5 | yes | Zandalar Madcap's Tunic (19834, -36.0) [quest]; Stormshroud Armor (15056, -80.0) [crafted]; Deathdealer's Vest (21364, -95.9) [quest] |
| wrist | Bonescythe Bracers (22483) | Quests [quest] | 87.1 | yes | Marshal's Leather Armsplints (16460, -65.7) [pvp]; General's Leather Armsplints (16559, -65.7) [pvp]; Forest Stalker's Bracers (19587, -65.7) [rep] |
| hands | Bonescythe Gauntlets (22481) | Quests [quest] | 123.7 | yes | Devilsaur Gauntlets (15063, -38.0) [crafted]; Marshal's Leather Handgrips (16454, -43.4) [vendor]; General's Leather Mitts (16560, -43.4) [vendor] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 121.7 | yes | Highlander's Leather Girdle (20045, -30.0) [rep]; Bonescythe Waistguard (22482, -36.9) [quest]; Highlander's Leather Girdle (20115, -44.0) [rep] |
| legs | Stormshroud Pants (15057) (or Knight-Captain's Leather Legguards (16419), Legionnaire's Leather Leggings (16508)) | Leatherworking [crafted] | 115.5 | yes | Knight-Captain's Leather Legguards (16419, +0.0) [pvp]; Legionnaire's Leather Leggings (16508, +0.0) [pvp]; Devilsaur Leggings (15062, -11.7) [crafted] |
| feet | Bonescythe Sabatons (22480) | Quests [quest] | 121.7 | yes | Highlander's Leather Boots (20052, -92.2) [rep]; Deathdealer's Boots (21359, -92.3) [quest]; Blood Guard's Leather Walkers (22856, -93.7) [vendor] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 109.7 | yes | Quick Strike Ring (18821, -22.0) [raid]; Don Julio's Band (19325, -36.0) [rep]; Band of the Penitent (13217, -52.0) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj [raid] | 97.7 | yes | Quick Strike Ring (18821, -10.0) [raid]; Don Julio's Band (19325, -24.0) [rep]; Band of the Penitent (13217, -40.0) [quest] |
| trinket1 | Kiss of the Spider (22954) | Naxxramas [raid] | 57.7 | yes | Eye of Diminution (23001, +57.7) [raid]; Drake Fang Talisman (19406, -1.7) [raid]; Thunderbrew's Boot Flask (744, -57.7) [quest] |
| trinket2 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +51.5) [raid]; Drake Fang Talisman (19406, -8.0) [raid]; Thunderbrew's Boot Flask (744, -64.0) [quest] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Quests [quest] | 761.0 | yes | The Hungering Cold (23577, +261.0) [raid]; Death's Sting (21126, +206.5) [raid]; Grand Marshal's Swiftblade (234579, +191.9) [pvp] |
| off_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 952.9 | yes | High Warlord's Left Claw (234558, -3.1) [pvp]; Grand Marshal's Left Hand Blade (234584, -3.1) [pvp]; Grand Marshal's Left Hand Blade (18847, -34.4) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 75.7 | yes | The Purifier (22656, -18.0) [quest]; Huhuran's Stinger (21616, -55.4) [raid]; Precisely Calibrated Boomstick (2100, -59.9) [dungeon] |

**New at 60:** head: Bonescythe Helmet; neck: Stormrage's Talisman of Seething; shoulder: Bonescythe Pauldrons; back: Chromatic Cloak; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Kiss of the Spider; trinket2: Slayer's Crest; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: High Warlord's Quickblade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1770, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (troll, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.1. Weights run: 1.1s. Verify run: 1.3s. 355 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.028, crit=2.157 ± 0.106, hit=1.402 ± 0.268, melee_haste=not significant (1.169 ± 0.749)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.8 | yes | Flying Tiger Goggles (4368, -8.8) [crafted]; Shadow Goggles (4373, -8.8) [crafted]; Lucky Fishing Hat (19972, -8.8) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.6 | yes | Tarnished Locket (279870, -6.6) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.5 | yes | Double-Stitched Woolen Shoulders (4314, -5.5) [crafted]; Reinforced Woolen Shoulders (4315, -5.5) [crafted]; Forest Leather Mantle (4709, -5.5) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.6 | yes | Catacomb Cloak (279899, -0.6) [quest]; Cape of the Brotherhood (5193, -1.1) [dungeon]; Sentry Cloak (2059, -2.2) [dungeon] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.7 | yes | Trapper's Leather Armor (252491, +0.0) [crafted]; Dark Leather Tunic (2317, -1.1) [crafted]; Heckler's Hide (286536, -2.2) [world] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 5.5 | yes | Wolf Bracers (4794, -1.1) [vendor]; Bravo's Armbands (270015, -1.1) [quest]; Ratchet Wristwraps (274742, -2.2) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 30.2 | yes | Serpent Gloves (5970, -23.6) [dungeon]; Gloves of the Fang (10413, -23.6) [dungeon]; Forest Leather Gloves (3058, -25.8) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -12.5) [crafted]; Dusty Belt (279897, -12.5) [quest]; Guardsman Belt (3429, -13.6) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.9 | yes | Brawler's Leather Pants (252500, +0.0) [crafted]; Trapper's Leather Pants (252501, +0.0) [crafted]; Bluegill Breeches (3022, -1.1) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.8 | yes | Blackened Defias Boots (10402, -2.2) [dungeon]; Footpads of the Fang (10411, -2.2) [dungeon]; Dark Leather Boots (2315, -3.3) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.6 | yes | Bounty Hunter's Ring (5351, -3.3) [quest]; Lavishly Jeweled Ring (1156, -4.4) [dungeon]; The 1 Ring (8350, -5.5) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 4.4 | yes | Bounty Hunter's Ring (5351, -1.1) [quest]; Lavishly Jeweled Ring (1156, -2.2) [dungeon]; The 1 Ring (8350, -3.3) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -20.3) [dungeon]; Diamond Hammer (2194, -22.1) [dungeon]; Wingblade (6504, -23.9) [quest] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -229.6) [quest]; Nightglow Concoction (3451, -229.6) [quest]; Pulsating Hydra Heart (5183, -229.6) [world] |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 | yes | Deadly Blunderbuss (4369, -1.8) [crafted]; Light Bow (4576, -1.8) [dungeon]; Privateer Musket (5309, -1.8) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Fine Longbow

No-known-source sample (15 of 355, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (troll, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 43.1. Weights run: 1.1s. Verify run: 1.7s. 687 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.099 ± 0.022, crit=3.270 ± 0.122, hit=1.653 ± 0.289, melee_haste=not significant (1.560 ± 0.801)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 | yes | Tribal Worg Helm (6204, -2.2) [world]; Brawler's Leather Hood (252504, -2.2) [crafted]; Holy Shroud (2721, -11.0) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -5.2) [rep]; Scout's Medallion (20442, -7.4) [rep]; Pendant of Myzrael (4614, -14.0) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 12.1 | yes | Mantle of Thieves (2264, -1.1) [dungeon]; Dark Leather Shoulders (4252, -4.4) [crafted]; Insignia Mantle (4721, -4.4) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.0) [quest]; Cloak of Night (4447, -3.4) [world]; Fenrus' Hide (6340, -3.4) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 15.4 | yes | Panther Armor (6670, -5.5) [quest]; Green Leather Armor (4255, -6.6) [crafted]; Brawler's Leather Tunic (252508, -6.6) [crafted] |
| wrist | Cultist's Armguards (270032) | Quests [quest] | 10.0 | yes | Jurassic Wristguards (6198, -3.4) [world]; Insignia Bracers (6410, -3.4) [dungeon]; Madwolf Bracers (897, -4.5) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 45.8 | yes | Heavy Earthen Gloves (7359, -29.8) [crafted]; Pilferer's Gloves (7358, -37.0) [crafted]; Braced Handguards (6784, -38.1) [quest] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.0) [rep]; Blackened Defias Belt (10403, -6.0) [dungeon]; Deftkin Belt (16659, -7.6) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -11.7) [crafted]; Insignia Leggings (4054, -16.1) [dungeon]; Leggings of the Fang (10410, -16.1) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.8 | yes | Insignia Boots (4055, +0.0) [dungeon]; Warsong Boots (16977, +0.0) [quest]; Highlander's Mail Greaves (20123, +0.0) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 | yes | Monkey Ring (6748, -2.2) [quest]; Ring of Precision (1491, -3.3) [dungeon]; Legionnaire's Band (19513, -3.3) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -1.3) [quest]; Ring of Precision (1491, -2.4) [dungeon]; Legionnaire's Band (19513, -2.4) [rep] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, -1.1) [dungeon]; Electrocutioner Leg (9446, -13.2) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -16.0) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -318.3) [quest]; Rod of Molten Fire (2565, -318.3) [dungeon]; Nightglow Concoction (3451, -318.3) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Silver Star (3463, -3.5) [quest]; Moonsight Rifle (4383, -4.0) [crafted]; Precision Bow (217315, -4.0) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 687, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (troll, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 83.1. Weights run: 1.1s. Verify run: 1.7s. 954 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.100 ± 0.014, crit=2.696 ± 0.072, hit=1.350 ± 0.103, melee_haste=2.063 ± 0.041

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 13.2 | yes | White Bandit Mask (10008, -1.1) [crafted]; Hawkeye's Helm (14591, -1.1) [world]; Spirit Hunter Headdress (6720, -2.2) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, -1.9) [rep]; Scout's Medallion (19537, -5.2) [rep]; Scout's Medallion (20442, -7.4) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.1 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Mantle of Thieves (2264, -13.1) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.0) [quest]; Imperial Cloak (6432, -1.2) [dungeon]; Parachute Cloak (10518, -1.2) [crafted] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 16.5 | yes | Dusky Leather Armor (7374, -1.1) [crafted]; Hawkeye's Tunic (14592, -3.3) [world]; Panther Armor (6670, -6.6) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Cultist's Armguards (270032, -10.0) [quest]; Imperial Leather Bracers (4061, -11.2) [dungeon]; Dusky Bracers (7378, -11.2) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 57.7 | yes | Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted]; Heavy Earthen Gloves (7359, -41.7) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Chain Girdle (20152, -6.0) [rep]; Defiler's Leather Girdle (20191, -6.0) [rep]; Blackened Defias Belt (10403, -12.0) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, -2.9) [dungeon]; Triprunner Dungarees (9624, -6.2) [quest]; Hawkeye's Breeches (14595, -10.6) [world] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 14.3 | yes | Imperial Leather Boots (6431, -2.2) [dungeon]; Dusky Boots (7390, -2.2) [crafted]; Skulker's Leather Shoes (252531, -2.2) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -2.1) [dungeon]; Insurgent's Band (272067, -3.0) [vendor]; Legionnaire's Band (19512, -3.2) [rep] |
| finger2 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 11.0 | yes | Ironspine's Eye (7686, -1.1) [dungeon]; Legionnaire's Band (19512, -2.2) [rep]; Disengagement Ring (276202, -2.2) [vendor] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -26.5) [vendor]; Hand of Righteousness (7721, -42.7) [dungeon]; Darkspear Skirmisher's Bludgeon (272093, -43.0) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Quests [quest] | 442.0 | yes | Grayson's Torch (1172, -442.0) [quest]; Rod of Molten Fire (2565, -442.0) [dungeon]; Nightglow Concoction (3451, -442.0) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Master Hunter's Bow (17686, -2.4) [quest]; Silver Star (3463, -3.5) [quest]; Mithril Blunderbuss (10508, -3.5) [crafted] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 954, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 115.4. Weights run: 1.2s. Verify run: 1.6s. 1257 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.103 ± 0.015, crit=3.283 ± 0.085, hit=1.682 ± 0.129, melee_haste=2.580 ± 0.051

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 18.8 | yes | Eye of Theradras (17715, +27.2) [dungeon]; Sprightring Helm (17776, -2.2) [quest]; Nightscape Headband (8176, -5.5) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19535, -0.8) [rep]; Scout's Medallion (19536, -1.9) [rep]; Woven Ivy Necklace (19159, -4.1) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.1 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Penance Spaulders (11963, -12.0) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 12.1 | yes | Nightscape Cloak (8195, -1.1) [crafted]; Wolfmaster Cape (6314, -2.1) [dungeon]; Battlehard Cape (11858, -2.1) [quest] |
| chest | Blazewind Breastplate (11193) | Quests [quest] | 25.4 | yes | Warbear Harness (15064, -5.5) [crafted]; Charred Leather Tunic (19127, -5.5) [quest]; Nightscape Tunic (8175, -8.8) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, -3.5) [world]; Wicked Leather Bracers (15084, -7.9) [crafted]; Pridelord Bands (14672, -9.0) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 66.0 | yes | Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted]; Heavy Earthen Gloves (7359, -50.0) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 66.0 | yes | Defiler's Cloth Girdle (20165, -20.0) [rep]; Defiler's Lizardhide Girdle (20174, -20.0) [rep]; Defiler's Leather Girdle (20192, -36.0) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 91.9 | yes | Ferine Leggings (6690, -65.9) [dungeon]; Basilisk Hide Pants (1718, -68.8) [dungeon]; Keeper's Woolies (14668, -71.0) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 | yes | Sandstalker Ankleguards (12470, -3.3) [dungeon]; Swampwalker Boots (2276, -7.7) [dungeon]; Skulker's Leather Boots (252469, -7.7) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.8 | yes | Masons Fraternity Ring (9533, -21.4) [quest]; Insurgent's Band (272065, -21.8) [vendor]; Insurgent's Band (272066, -24.8) [vendor] |
| finger2 | White Bone Band (11862) | Quests [quest] | 24.0 | yes | Masons Fraternity Ring (9533, -8.6) [quest]; Insurgent's Band (272065, -9.0) [vendor]; Insurgent's Band (272066, -12.0) [vendor] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, +53.8) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Defiler's Talisman (21115) | The Defilers [rep] | 0.0 | yes | Rune of the Guard Captain (19120, +53.8) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Inventor's Focal Sword (17719, -14.1) [dungeon]; Might of Hakkar (10838, -43.2) [world]; Lifeforce Dirk (10750, -50.1) [quest] |
| off_hand | Julie's Dagger (6660) | Blackrock Depths: Shadowforge Peasant [dungeon] | 511.6 | yes | Claw of Celebras (17738, -29.4) [dungeon]; White Bone Shredder (11863, -68.9) [quest]; Thermotastic Egg Timer (9644, -508.3) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | Blackrock Depths: Anvilrage Overseer [dungeon] | 15.4 | yes | Houndmaster's Bow (11628, -3.4) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -6.4) [vendor]; Guttbuster (13139, -6.6) [world] |

**New at 50:** head: Helm of Fire; back: Serpentskin Cloak; chest: Blazewind Breastplate; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Smoking Heart of the Mountain; trinket2: Defiler's Talisman; main_hand: Hammer of the Northern Wind; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1257, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 288.0. Weights run: 1.2s. Verify run: 1.9s. 1765 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=not significant (0.000 ± 0.000), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Quests [quest] | 149.4 | yes | Ragefury Eyepatch (11735, -33.9) [dungeon]; Bloodvine Lens (19998, -33.9) [crafted]; Champion's Leather Helm (23257, -55.7) [vendor] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 141.5 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Barbed Choker (21664, -39.7) [raid]; Medallion of the Dawn (22659, -59.7) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Quests [quest] | 82.6 | yes | Champion's Leather Shoulders (23258, -2.9) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -2.9) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -2.9) [pvp] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 57.7 | yes | Deathguard's Cloak (20068, -18.1) [rep]; Cape of the Black Baron (13340, -20.8) [dungeon]; Cloak of the Fallen God (21710, -28.3) [quest] |
| chest | Bonescythe Breastplate (22476) | Quests [quest] | 195.5 | yes | Zandalar Madcap's Tunic (19834, -36.0) [quest]; Stormshroud Armor (15056, -80.0) [crafted]; Deathdealer's Vest (21364, -95.9) [quest] |
| wrist | Bonescythe Bracers (22483) | Quests [quest] | 87.1 | yes | Marshal's Leather Armsplints (16460, -65.7) [pvp]; General's Leather Armsplints (16559, -65.7) [pvp]; Forest Stalker's Bracers (19587, -65.7) [rep] |
| hands | Bonescythe Gauntlets (22481) | Quests [quest] | 123.7 | yes | Devilsaur Gauntlets (15063, -38.0) [crafted]; Marshal's Leather Handgrips (16454, -43.4) [vendor]; General's Leather Mitts (16560, -43.4) [vendor] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 121.7 | yes | Defiler's Leather Girdle (20190, -30.0) [rep]; Bonescythe Waistguard (22482, -36.9) [quest]; Defiler's Leather Girdle (20193, -44.0) [rep] |
| legs | Stormshroud Pants (15057) (or Knight-Captain's Leather Legguards (16419), Legionnaire's Leather Leggings (16508)) | Leatherworking [crafted] | 115.5 | yes | Knight-Captain's Leather Legguards (16419, +0.0) [pvp]; Legionnaire's Leather Leggings (16508, +0.0) [pvp]; Devilsaur Leggings (15062, -11.7) [crafted] |
| feet | Bonescythe Sabatons (22480) | Quests [quest] | 121.7 | yes | Defiler's Leather Boots (20186, -92.2) [rep]; Deathdealer's Boots (21359, -92.3) [quest]; Blood Guard's Leather Walkers (22856, -93.7) [vendor] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 109.7 | yes | Quick Strike Ring (18821, -22.0) [raid]; Don Julio's Band (19325, -36.0) [rep]; Band of the Penitent (13217, -52.0) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj [raid] | 97.7 | yes | Quick Strike Ring (18821, -10.0) [raid]; Don Julio's Band (19325, -24.0) [rep]; Band of the Penitent (13217, -40.0) [quest] |
| trinket1 | Kiss of the Spider (22954) | Naxxramas [raid] | 57.7 | yes | Eye of Diminution (23001, +57.7) [raid]; Drake Fang Talisman (19406, -1.7) [raid]; Rune of the Guard Captain (19120, -15.7) [quest] |
| trinket2 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +51.5) [raid]; Drake Fang Talisman (19406, -8.0) [raid]; Rune of the Guard Captain (19120, -22.0) [quest] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Quests [quest] | 761.0 | yes | The Hungering Cold (23577, +261.0) [raid]; Death's Sting (21126, +206.5) [raid]; Grand Marshal's Swiftblade (234579, +191.9) [pvp] |
| off_hand | High Warlord's Quickblade (234553) | Rank 18 [pvp] | 952.9 | yes | High Warlord's Left Claw (234558, -3.1) [pvp]; Grand Marshal's Left Hand Blade (234584, -3.1) [pvp]; Grand Marshal's Left Hand Blade (18847, -34.4) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 75.7 | yes | The Purifier (22656, -18.0) [quest]; Huhuran's Stinger (21616, -55.4) [raid]; Precisely Calibrated Boomstick (2100, -59.9) [dungeon] |

**New at 60:** head: Bonescythe Helmet; neck: Stormrage's Talisman of Seething; shoulder: Bonescythe Pauldrons; back: Chromatic Cloak; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Kiss of the Spider; trinket2: Slayer's Crest; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: High Warlord's Quickblade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1765, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

