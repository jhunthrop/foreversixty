# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 33.7. Weights run: 1.1s. Verify run: 1.3s. 360 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.009 ± 0.003, crit=0.279 ± 0.037, hit=1.342 ± 0.255, melee_haste=not significant (1.305 ± 0.713)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Flying Tiger Goggles (4368, -8.1) [crafted]; Shadow Goggles (4373, -8.1) [crafted]; Lucky Fishing Hat (19972, -8.1) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 | yes | Tarnished Locket (279870, -6.1) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -5.0) [crafted]; Reinforced Woolen Shoulders (4315, -5.0) [crafted]; Forest Leather Mantle (4709, -5.0) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.1) [quest]; Cape of the Brotherhood (5193, -1.0) [dungeon]; Sentry Cloak (2059, -2.0) [dungeon] |
| chest | Tunic of Westfall (2041) | Quests [quest] | 11.1 | yes | Brawler's Leather Armor (252490, -4.0) [crafted]; Trapper's Leather Armor (252491, -4.0) [crafted]; Dark Leather Tunic (2317, -5.0) [crafted] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 5.0 | yes | Wolf Bracers (4794, -1.0) [vendor]; Bravo's Armbands (270015, -1.0) [quest]; Ratchet Wristwraps (274742, -2.0) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.0) [dungeon]; Forest Leather Gloves (3058, -2.0) [dungeon]; Nimble Leather Gloves (7285, -2.0) [crafted] |
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

### Band 30 (night-elf, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 40.1. Weights run: 1.2s. Verify run: 1.8s. 691 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.042 ± 0.026, crit=0.407 ± 0.046, hit=not significant (0.961 ± 0.247), melee_haste=not significant (1.094 ± 0.685)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.4 | yes | Tribal Worg Helm (6204, -2.1) [world]; Brawler's Leather Hood (252504, -2.1) [crafted]; Holy Shroud (2721, -10.4) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -5.7) [rep]; Sentinel's Medallion (20444, -7.7) [rep]; Pendant of Myzrael (4614, -14.0) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 11.5 | yes | Mantle of Thieves (2264, -1.0) [dungeon]; Dark Leather Shoulders (4252, -4.2) [crafted]; Insignia Mantle (4721, -4.2) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Cloak of Night (4447, -3.7) [world]; Fenrus' Hide (6340, -3.7) [dungeon]; Glowing Lizardscale Cloak (6449, -3.7) [dungeon] |
| chest | Raptorbane Armor (3566) | Quests [quest] | 16.0 | yes | Dusky Leather Armor (7374, -1.4) [crafted]; Tunic of Westfall (2041, -4.5) [quest]; Green Leather Armor (4255, -7.7) [crafted] |
| wrist | Cultist's Armguards (270032) | Quests [quest] | 10.0 | yes | Jurassic Wristguards (6198, -3.7) [world]; Insignia Bracers (6410, -3.7) [dungeon]; Madwolf Bracers (897, -4.8) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Pilferer's Gloves (7358, -7.7) [crafted]; Wolfclaw Gloves (1978, -9.7) [dungeon]; Toughened Leather Gloves (4253, -9.7) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.0) [rep]; Blackened Defias Belt (10403, -6.0) [dungeon]; Skulker's Leather Belt (252520, -14.6) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -12.5) [crafted]; Insignia Leggings (4054, -16.6) [dungeon]; Leggings of the Fang (10410, -16.6) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.3 | yes | Insignia Boots (4055, +0.0) [dungeon]; Highlander's Mail Greaves (20123, +0.0) [vendor]; Lancer Boots (6752, -1.0) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.4 | yes | Monkey Ring (6748, -2.1) [quest]; Ring of Precision (1491, -3.1) [dungeon]; Protector's Band (19517, -3.1) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -1.7) [quest]; Ring of Precision (1491, -2.7) [dungeon]; Protector's Band (19517, -2.7) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, -1.1) [dungeon]; Electrocutioner Leg (9446, -13.2) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -16.0) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -318.3) [quest]; Rod of Molten Fire (2565, -318.3) [dungeon]; Eye of Paleth (2943, -318.3) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Silver Star (3463, -3.8) [quest]; BKP "Sparrow" Smallbore (3042, -4.8) [dungeon]; Alliance Outrunner Bow (285347, -4.8) [world] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 77.4. Weights run: 1.3s. Verify run: 1.7s. 958 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.023 ± 0.008, crit=0.659 ± 0.081, hit=not significant (2.739 ± 0.703), melee_haste=not significant (2.974 ± 1.701)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.3 | yes | White Bandit Mask (10008, -1.0) [crafted]; Hawkeye's Helm (14591, -1.0) [world]; Brawler's Leather Helm (252512, -2.0) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, -2.7) [rep]; Sentinel's Medallion (19541, -5.8) [rep]; Sentinel's Medallion (20444, -7.9) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Mantle of Thieves (2264, -13.0) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, -1.8) [dungeon]; Parachute Cloak (10518, -1.8) [crafted]; Yeti Fur Cloak (2805, -3.9) [quest] |
| chest | Raptorbane Armor (3566) | Quests [quest] | 16.0 | yes | Nightscape Tunic (8175, -0.7) [crafted]; Dusky Leather Armor (7374, -1.7) [crafted]; Hawkeye's Tunic (14592, -3.7) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Cultist's Armguards (270032, -10.0) [quest]; Imperial Leather Bracers (4061, -11.8) [dungeon]; Dusky Bracers (7378, -11.8) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 29.2 | yes | Heavy Earthen Gloves (7359, -13.2) [crafted]; Skulker's Leather Gloves (252525, -19.0) [crafted]; Stalker's Leather Gloves (252526, -19.0) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Chain Girdle (20090, -6.0) [rep]; Highlander's Leather Girdle (20117, -6.0) [rep]; Blackened Defias Belt (10403, -12.0) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, -4.5) [dungeon]; Triprunner Dungarees (9624, -7.6) [quest]; Hawkeye's Breeches (14595, -11.7) [world] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 13.3 | yes | Imperial Leather Boots (6431, -2.0) [dungeon]; Dusky Boots (7390, -2.0) [crafted]; Skulker's Leather Shoes (252531, -2.0) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -2.8) [dungeon]; Insurgent's Band (272067, -3.0) [vendor]; Protector's Band (19515, -3.8) [rep] |
| finger2 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 10.2 | yes | Ironspine's Eye (7686, -1.0) [dungeon]; Protector's Band (19515, -2.0) [rep]; Disengagement Ring (276202, -2.0) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -27.2) [vendor]; Sword of Serenity (6829, -40.0) [quest]; Hand of Righteousness (7721, -42.7) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Quests [quest] | 442.0 | yes | Grayson's Torch (1172, -442.0) [quest]; Rod of Molten Fire (2565, -442.0) [dungeon]; Eye of Paleth (2943, -442.0) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Moonsight Rifle (4383, -0.8) [crafted]; Precision Bow (217315, -0.8) [quest]; Master Hunter's Bow (17686, -2.9) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 958, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 101.0. Weights run: 1.3s. Verify run: 1.5s. 1261 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.395 ± 0.118, crit=6.478 ± 0.350, hit=4.201 ± 0.976, melee_haste=not significant (0.827 ± 2.686)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 90.7 | yes | Helm of Fire (8348, -67.0) [crafted]; Lordrec Helmet (10741, -68.4) [quest]; Sprightring Helm (17776, -69.8) [quest] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 16.7 | yes | Sentinel's Medallion (19540, -1.4) [rep]; Ghostshard Talisman (7731, -2.7) [dungeon]; Sentinel's Medallion (19541, -5.6) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.3 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Penance Spaulders (11963, -12.0) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 15.3 | yes | Nightscape Cloak (8195, -1.4) [crafted]; Pridelord Cape (14673, -2.8) [dungeon]; Imperial Cloak (6432, -4.2) [dungeon] |
| chest | Blazewind Breastplate (11193) | Quests [quest] | 32.1 | yes | Warbear Harness (15064, -7.0) [crafted]; Charred Leather Tunic (19127, -7.0) [quest]; Nightscape Tunic (8175, -11.2) [crafted] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 20.9 | yes | Branded Leather Bracers (19508, -0.9) [dungeon]; Wicked Leather Bracers (15084, -5.6) [crafted]; Pridelord Bands (14672, -7.0) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 110.7 | yes | Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted]; Wicked Leather Gauntlets (15083, -93.9) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 110.7 | yes | Highlander's Cloth Girdle (20097, -20.0) [rep]; Highlander's Lizardhide Girdle (20103, -20.0) [rep]; Highlander's Leather Girdle (20116, -80.7) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 181.4 | yes | Basilisk Hide Pants (1718, -152.1) [dungeon]; Keeper's Woolies (14668, -154.9) [dungeon]; Ferine Leggings (6690, -155.4) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 27.9 | yes | Sandstalker Ankleguards (12470, -4.2) [dungeon]; Swampwalker Boots (2276, -9.8) [dungeon]; Skulker's Leather Boots (252469, -9.8) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 62.0 | yes | Insurgent's Band (272065, -47.0) [vendor]; Ring of the Underwood (2951, -48.1) [dungeon]; Ironspine's Eye (7686, -49.5) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Quests [quest] | 19.5 | yes | Insurgent's Band (272065, -4.5) [vendor]; Ring of the Underwood (2951, -5.6) [dungeon]; Ironspine's Eye (7686, -7.0) [dungeon] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 583.9 | yes | Might of Hakkar (10838, -61.3) [world]; Julie's Dagger (6660, -72.4) [dungeon]; Lifeforce Dirk (10750, -80.8) [quest] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -71.1) [dungeon]; Thermotastic Egg Timer (9644, -549.1) [quest]; Grayson's Torch (1172, -553.3) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | Blackrock Depths: Anvilrage Overseer [dungeon] | 19.5 | yes | Moonsight Rifle (4383, -6.9) [crafted]; Precision Bow (217315, -6.9) [quest]; Houndmaster's Bow (11628, -7.5) [dungeon] |

**New at 50:** head: Eye of Theradras; neck: Sentinel's Medallion; back: Serpentskin Cloak; chest: Blazewind Breastplate; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Smoking Heart of the Mountain; trinket2: Talisman of Arathor; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1261, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 304.7. Weights run: 1.3s. Verify run: 2.1s. 1770 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=8.118 ± 0.504, hit=not significant (5.733 ± 1.563), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Quests [quest] | 320.1 | yes | Mask of the Unforgiven (13404, -91.7) [dungeon]; Bloodvine Goggles (19999, -91.7) [crafted]; Ragefury Eyepatch (11735, -92.8) [dungeon] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 253.3 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Onyxia Tooth Pendant (18404, -74.1) [quest]; Fury of the Forgotten Swarm (21809, -82.3) [raid] |
| shoulder | Bonescythe Pauldrons (22479) | Quests [quest] | 197.0 | yes | Lieutenant Commander's Leather Shoulders (23313, -4.0) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -4.0) [pvp]; Champion's Leather Shoulders (227056, -4.0) [pvp] |
| back | Cloak of Veiled Shadows (21406) | Quests [quest] | 78.6 | yes | Chromatic Cloak (18509, +35.1) [crafted]; Earthweave Cloak (21187, -3.5) [quest]; Cloak of the Honor Guard (20073, -38.7) [rep] |
| chest | Bonescythe Breastplate (22476) | Quests [quest] | 364.6 | yes | Zandalar Madcap's Tunic (19834, -93.3) [quest]; Stormshroud Armor (15056, -137.3) [crafted]; Deathdealer's Vest (21364, -150.0) [quest] |
| wrist | Bonescythe Bracers (22483) | Quests [quest] | 144.4 | yes | Qiraji Execution Bracers (21602, -68.1) [raid]; Primal Batskin Bracers (19687, -70.5) [crafted]; Rockfury Bracers (21186, -87.0) [quest] |
| hands | Bonescythe Gauntlets (22481) | Quests [quest] | 237.0 | yes | Stormshroud Gloves (21278, -66.0) [crafted]; Devilsaur Gauntlets (15063, -95.3) [crafted]; Marshal's Leather Handgrips (16454, -99.7) [vendor] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 235.0 | yes | Highlander's Leather Girdle (20045, -87.3) [rep]; Bonescythe Waistguard (22482, -93.0) [quest]; Highlander's Leather Girdle (20115, -101.3) [rep] |
| legs | Marshal's Leather Leggings (16456) (or General's Leather Legguards (16564), Marshal's Leather Leggings (231548), General's Leather Legguards (231554)) | Captain Dirgehammer [vendor] | 260.2 | yes | General's Leather Legguards (16564, +0.0) [vendor]; Marshal's Leather Leggings (231548, +0.0) [pvp]; General's Leather Legguards (231554, +0.0) [pvp] |
| feet | Bonescythe Sabatons (22480) | Quests [quest] | 235.0 | yes | Deathdealer's Boots (21359, -147.0) [quest]; Bloodvine Boots (19684, -177.7) [crafted]; Shadowcraft Boots (16711, -193.0) [dungeon] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 223.0 | yes | Band of Earthen Might (21182, -52.0) [quest]; Seal of the Damned (23025, -52.0) [raid]; Ring of the Qiraji Fury (21677, -69.3) [raid] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 187.0 | yes | Band of Earthen Might (21182, -16.0) [quest]; Seal of the Damned (23025, -16.0) [raid]; Ring of the Qiraji Fury (21677, -33.3) [raid] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +163.3) [raid]; Drake Fang Talisman (19406, +106.7) [raid]; Neltharion's Tear (19379, +50.7) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 171.0 | yes | Eye of Diminution (23001, +56.3) [raid]; Drake Fang Talisman (19406, -0.3) [raid]; Neltharion's Tear (19379, -56.3) [raid] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Quests [quest] | 761.2 | yes | Blessed Qiraji Pugio (21244, +276.0) [quest]; High Warlord's Quickblade (234553, +247.6) [pvp]; Grand Marshal's Swiftblade (234579, +247.6) [pvp] |
| off_hand | The Hungering Cold (23577) | Naxxramas [raid] | 1022.0 | yes | High Warlord's Left Claw (234558, -16.3) [pvp]; Grand Marshal's Left Hand Blade (234584, -16.3) [pvp]; Grand Marshal's Left Hand Blade (18847, -47.6) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 131.7 | yes | The Purifier (22656, -18.0) [quest]; Fahrad's Reloading Repeater (22347, -69.6) [quest]; Core Marksman Rifle (18282, -74.3) [crafted] |

**New at 60:** head: Bonescythe Helmet; neck: Stormrage's Talisman of Seething; shoulder: Bonescythe Pauldrons; back: Cloak of Veiled Shadows; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; legs: Marshal's Leather Leggings; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Don Julio's Band; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: The Hungering Cold; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1770, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 33.1. Weights run: 1.1s. Verify run: 1.3s. 355 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.009 ± 0.003, crit=0.279 ± 0.037, hit=1.342 ± 0.255, melee_haste=not significant (1.305 ± 0.713)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Flying Tiger Goggles (4368, -8.1) [crafted]; Shadow Goggles (4373, -8.1) [crafted]; Lucky Fishing Hat (19972, -8.1) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 | yes | Tarnished Locket (279870, -6.1) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -5.0) [crafted]; Reinforced Woolen Shoulders (4315, -5.0) [crafted]; Forest Leather Mantle (4709, -5.0) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.1) [quest]; Cape of the Brotherhood (5193, -1.0) [dungeon]; Sentry Cloak (2059, -2.0) [dungeon] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 | yes | Trapper's Leather Armor (252491, +0.0) [crafted]; Dark Leather Tunic (2317, -1.0) [crafted]; Heckler's Hide (286536, -2.0) [world] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 5.0 | yes | Wolf Bracers (4794, -1.0) [vendor]; Bravo's Armbands (270015, -1.0) [quest]; Ratchet Wristwraps (274742, -2.0) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.0) [dungeon]; Forest Leather Gloves (3058, -2.0) [dungeon]; Nimble Leather Gloves (7285, -2.0) [crafted] |
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

### Band 30 (troll, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 39.4. Weights run: 1.2s. Verify run: 1.8s. 687 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.042 ± 0.026, crit=0.407 ± 0.046, hit=not significant (0.961 ± 0.247), melee_haste=not significant (1.094 ± 0.685)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.4 | yes | Tribal Worg Helm (6204, -2.1) [world]; Brawler's Leather Hood (252504, -2.1) [crafted]; Holy Shroud (2721, -10.4) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -5.7) [rep]; Scout's Medallion (20442, -7.7) [rep]; Pendant of Myzrael (4614, -14.0) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 11.5 | yes | Mantle of Thieves (2264, -1.0) [dungeon]; Dark Leather Shoulders (4252, -4.2) [crafted]; Insignia Mantle (4721, -4.2) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.0) [quest]; Cloak of Night (4447, -3.7) [world]; Fenrus' Hide (6340, -3.7) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.6 | yes | Panther Armor (6670, -5.2) [quest]; Green Leather Armor (4255, -6.3) [crafted]; Brawler's Leather Tunic (252508, -6.3) [crafted] |
| wrist | Cultist's Armguards (270032) | Quests [quest] | 10.0 | yes | Jurassic Wristguards (6198, -3.7) [world]; Insignia Bracers (6410, -3.7) [dungeon]; Madwolf Bracers (897, -4.8) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Pilferer's Gloves (7358, -7.7) [crafted]; Braced Handguards (6784, -8.7) [quest]; Wolfclaw Gloves (1978, -9.7) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.0) [rep]; Blackened Defias Belt (10403, -6.0) [dungeon]; Deftkin Belt (16659, -7.8) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -12.5) [crafted]; Insignia Leggings (4054, -16.6) [dungeon]; Leggings of the Fang (10410, -16.6) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 8.3 | yes | Insignia Boots (4055, +0.0) [dungeon]; Warsong Boots (16977, +0.0) [quest]; Highlander's Mail Greaves (20123, +0.0) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.4 | yes | Monkey Ring (6748, -2.1) [quest]; Ring of Precision (1491, -3.1) [dungeon]; Legionnaire's Band (19513, -3.1) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -1.7) [quest]; Ring of Precision (1491, -2.7) [dungeon]; Legionnaire's Band (19513, -2.7) [rep] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, -1.1) [dungeon]; Electrocutioner Leg (9446, -13.2) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -16.0) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 | yes | Grayson's Torch (1172, -318.3) [quest]; Rod of Molten Fire (2565, -318.3) [dungeon]; Nightglow Concoction (3451, -318.3) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Silver Star (3463, -3.8) [quest]; BKP "Sparrow" Smallbore (3042, -4.8) [dungeon]; Alliance Outrunner Bow (285347, -4.8) [world] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 687, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 76.9. Weights run: 1.3s. Verify run: 1.8s. 954 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.023 ± 0.008, crit=0.659 ± 0.081, hit=not significant (2.739 ± 0.703), melee_haste=not significant (2.974 ± 1.701)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.3 | yes | White Bandit Mask (10008, -1.0) [crafted]; Hawkeye's Helm (14591, -1.0) [world]; Spirit Hunter Headdress (6720, -2.0) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, -2.7) [rep]; Scout's Medallion (19537, -5.8) [rep]; Scout's Medallion (20442, -7.9) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Mantle of Thieves (2264, -13.0) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.0) [quest]; Imperial Cloak (6432, -1.8) [dungeon]; Parachute Cloak (10518, -1.8) [crafted] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 15.3 | yes | Dusky Leather Armor (7374, -1.0) [crafted]; Hawkeye's Tunic (14592, -3.1) [world]; Panther Armor (6670, -6.1) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Cultist's Armguards (270032, -10.0) [quest]; Imperial Leather Bracers (4061, -11.8) [dungeon]; Dusky Bracers (7378, -11.8) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 29.2 | yes | Heavy Earthen Gloves (7359, -13.2) [crafted]; Skulker's Leather Gloves (252525, -19.0) [crafted]; Stalker's Leather Gloves (252526, -19.0) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Chain Girdle (20152, -6.0) [rep]; Defiler's Leather Girdle (20191, -6.0) [rep]; Blackened Defias Belt (10403, -12.0) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, -4.5) [dungeon]; Triprunner Dungarees (9624, -7.6) [quest]; Hawkeye's Breeches (14595, -11.7) [world] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 13.3 | yes | Imperial Leather Boots (6431, -2.0) [dungeon]; Dusky Boots (7390, -2.0) [crafted]; Skulker's Leather Shoes (252531, -2.0) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -2.8) [dungeon]; Insurgent's Band (272067, -3.0) [vendor]; Legionnaire's Band (19512, -3.8) [rep] |
| finger2 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 10.2 | yes | Ironspine's Eye (7686, -1.0) [dungeon]; Legionnaire's Band (19512, -2.0) [rep]; Disengagement Ring (276202, -2.0) [vendor] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Ardent Custodian (868) | Gnomeregan: Dark Iron Ambassador [dungeon] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -27.2) [vendor]; Hand of Righteousness (7721, -42.7) [dungeon]; Darkspear Skirmisher's Bludgeon (272093, -43.0) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Quests [quest] | 442.0 | yes | Grayson's Torch (1172, -442.0) [quest]; Rod of Molten Fire (2565, -442.0) [dungeon]; Nightglow Concoction (3451, -442.0) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Moonsight Rifle (4383, -0.8) [crafted]; Precision Bow (217315, -0.8) [quest]; Master Hunter's Bow (17686, -2.9) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword

No-known-source sample (15 of 954, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 100.2. Weights run: 1.3s. Verify run: 1.5s. 1257 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.395 ± 0.118, crit=6.478 ± 0.350, hit=4.201 ± 0.976, melee_haste=not significant (0.827 ± 2.686)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 90.7 | yes | Helm of Fire (8348, -67.0) [crafted]; Sprightring Helm (17776, -69.8) [quest]; Nightscape Headband (8176, -73.9) [crafted] |
| neck | Scout's Medallion (19535) | Warsong Outriders [rep] | 16.7 | yes | Scout's Medallion (19536, -1.4) [rep]; Ghostshard Talisman (7731, -2.7) [dungeon]; Woven Ivy Necklace (19159, -4.2) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.3 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Penance Spaulders (11963, -12.0) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 15.3 | yes | Nightscape Cloak (8195, -1.4) [crafted]; Pridelord Cape (14673, -2.8) [dungeon]; Imperial Cloak (6432, -4.2) [dungeon] |
| chest | Blazewind Breastplate (11193) | Quests [quest] | 32.1 | yes | Warbear Harness (15064, -7.0) [crafted]; Charred Leather Tunic (19127, -7.0) [quest]; Nightscape Tunic (8175, -11.2) [crafted] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 20.9 | yes | Branded Leather Bracers (19508, -0.9) [dungeon]; Wicked Leather Bracers (15084, -5.6) [crafted]; Pridelord Bands (14672, -7.0) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 110.7 | yes | Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted]; Wicked Leather Gauntlets (15083, -93.9) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 110.7 | yes | Defiler's Cloth Girdle (20165, -20.0) [rep]; Defiler's Lizardhide Girdle (20174, -20.0) [rep]; Defiler's Leather Girdle (20192, -80.7) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 181.4 | yes | Basilisk Hide Pants (1718, -152.1) [dungeon]; Keeper's Woolies (14668, -154.9) [dungeon]; Ferine Leggings (6690, -155.4) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 27.9 | yes | Sandstalker Ankleguards (12470, -4.2) [dungeon]; Swampwalker Boots (2276, -9.8) [dungeon]; Skulker's Leather Boots (252469, -9.8) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 62.0 | yes | Masons Fraternity Ring (9533, -42.5) [quest]; Insurgent's Band (272065, -47.0) [vendor]; Ring of the Underwood (2951, -48.1) [dungeon] |
| finger2 | White Bone Band (11862) | Quests [quest] | 24.0 | yes | Masons Fraternity Ring (9533, -4.5) [quest]; Insurgent's Band (272065, -9.0) [vendor]; Ring of the Underwood (2951, -10.0) [dungeon] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, +71.4) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Defiler's Talisman (21115) | The Defilers [rep] | 0.0 | yes | Rune of the Guard Captain (19120, +71.4) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 583.9 | yes | Might of Hakkar (10838, -61.3) [world]; Julie's Dagger (6660, -72.4) [dungeon]; Lifeforce Dirk (10750, -80.8) [quest] |
| off_hand | Hammer of the Northern Wind (810) | Blackrock Depths: Anvilrage Medic [dungeon] | 553.3 | yes | Claw of Celebras (17738, -71.1) [dungeon]; White Bone Shredder (11863, -108.5) [quest]; Thermotastic Egg Timer (9644, -549.1) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | Blackrock Depths: Anvilrage Overseer [dungeon] | 19.5 | yes | Moonsight Rifle (4383, -6.9) [crafted]; Precision Bow (217315, -6.9) [quest]; Houndmaster's Bow (11628, -7.5) [dungeon] |

**New at 50:** head: Eye of Theradras; neck: Scout's Medallion; back: Serpentskin Cloak; chest: Blazewind Breastplate; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Smoking Heart of the Mountain; trinket2: Defiler's Talisman; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1257, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 301.6. Weights run: 1.3s. Verify run: 2.1s. 1765 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=8.118 ± 0.504, hit=not significant (5.733 ± 1.563), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Quests [quest] | 320.1 | yes | Mask of the Unforgiven (13404, -91.7) [dungeon]; Bloodvine Goggles (19999, -91.7) [crafted]; Ragefury Eyepatch (11735, -92.8) [dungeon] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 253.3 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Onyxia Tooth Pendant (18404, -74.1) [quest]; Fury of the Forgotten Swarm (21809, -82.3) [raid] |
| shoulder | Bonescythe Pauldrons (22479) | Quests [quest] | 197.0 | yes | Champion's Leather Shoulders (23258, -4.0) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -4.0) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -4.0) [pvp] |
| back | Cloak of Veiled Shadows (21406) | Quests [quest] | 78.6 | yes | Chromatic Cloak (18509, +35.1) [crafted]; Earthweave Cloak (21187, -3.5) [quest]; Deathguard's Cloak (20068, -38.7) [rep] |
| chest | Bonescythe Breastplate (22476) | Quests [quest] | 364.6 | yes | Zandalar Madcap's Tunic (19834, -93.3) [quest]; Stormshroud Armor (15056, -137.3) [crafted]; Deathdealer's Vest (21364, -150.0) [quest] |
| wrist | Bonescythe Bracers (22483) | Quests [quest] | 144.4 | yes | Qiraji Execution Bracers (21602, -68.1) [raid]; Primal Batskin Bracers (19687, -70.5) [crafted]; Rockfury Bracers (21186, -87.0) [quest] |
| hands | Bonescythe Gauntlets (22481) | Quests [quest] | 237.0 | yes | Stormshroud Gloves (21278, -66.0) [crafted]; Devilsaur Gauntlets (15063, -95.3) [crafted]; Marshal's Leather Handgrips (16454, -99.7) [vendor] |
| waist | Belt of Never-ending Agony (21586) | Ahn'Qiraj [raid] | 235.0 | yes | Defiler's Leather Girdle (20190, -87.3) [rep]; Bonescythe Waistguard (22482, -93.0) [quest]; Defiler's Leather Girdle (20193, -101.3) [rep] |
| legs | Marshal's Leather Leggings (16456) (or General's Leather Legguards (16564), Marshal's Leather Leggings (231548), General's Leather Legguards (231554)) | Captain Dirgehammer [vendor] | 260.2 | yes | General's Leather Legguards (16564, +0.0) [vendor]; Marshal's Leather Leggings (231548, +0.0) [pvp]; General's Leather Legguards (231554, +0.0) [pvp] |
| feet | Bonescythe Sabatons (22480) | Quests [quest] | 235.0 | yes | Deathdealer's Boots (21359, -147.0) [quest]; Bloodvine Boots (19684, -177.7) [crafted]; Shadowcraft Boots (16711, -193.0) [dungeon] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 223.0 | yes | Band of Earthen Might (21182, -52.0) [quest]; Seal of the Damned (23025, -52.0) [raid]; Ring of the Qiraji Fury (21677, -69.3) [raid] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 187.0 | yes | Band of Earthen Might (21182, -16.0) [quest]; Seal of the Damned (23025, -16.0) [raid]; Ring of the Qiraji Fury (21677, -33.3) [raid] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +163.3) [raid]; Drake Fang Talisman (19406, +106.7) [raid]; Neltharion's Tear (19379, +50.7) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 171.0 | yes | Eye of Diminution (23001, +56.3) [raid]; Drake Fang Talisman (19406, -0.3) [raid]; Neltharion's Tear (19379, -56.3) [raid] |
| main_hand | Thunderfury, Blessed Blade of the Windseeker (19019) | Quests [quest] | 761.2 | yes | Blessed Qiraji Pugio (21244, +276.0) [quest]; High Warlord's Quickblade (234553, +247.6) [pvp]; Grand Marshal's Swiftblade (234579, +247.6) [pvp] |
| off_hand | The Hungering Cold (23577) | Naxxramas [raid] | 1022.0 | yes | High Warlord's Left Claw (234558, -16.3) [pvp]; Grand Marshal's Left Hand Blade (234584, -16.3) [pvp]; Grand Marshal's Left Hand Blade (18847, -47.6) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 131.7 | yes | The Purifier (22656, -18.0) [quest]; Fahrad's Reloading Repeater (22347, -69.6) [quest]; Core Marksman Rifle (18282, -74.3) [crafted] |

**New at 60:** head: Bonescythe Helmet; neck: Stormrage's Talisman of Seething; shoulder: Bonescythe Pauldrons; back: Cloak of Veiled Shadows; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Belt of Never-ending Agony; legs: Marshal's Leather Leggings; feet: Bonescythe Sabatons; finger1: Band of Unnatural Forces; finger2: Don Julio's Band; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Thunderfury, Blessed Blade of the Windseeker; off_hand: The Hungering Cold; ranged: Larvae of the Great Worm

No-known-source sample (15 of 1765, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

