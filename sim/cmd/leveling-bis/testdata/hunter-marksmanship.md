# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (dwarf, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.1. Weights run: 0.7s. Verify run: 1.0s. 384 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.004, agility=2.000 ± 0.009, crit=6.978 ± 1.340, hit=not significant (4.521 ± 1.506), melee_haste=not significant (6.060 ± 4.886)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.0 | yes | Shadow Goggles (4373, -0.96 DPS) [crafted]; Lucky Fishing Hat (19972, -0.96 DPS) [quest]; Flying Tiger Goggles (4368, -1.38 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.0 | yes | Tarnished Locket (279870, -1.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 | yes | Reinforced Woolen Shoulders (4315, -0.60 DPS) [crafted]; Forest Leather Mantle (4709, -0.60 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.61 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 | yes | Cape of the Brotherhood (5193, +0.02 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.24 DPS) [dungeon]; Hide of Lupos (3018, -0.24 DPS) [world] |
| chest | Tunic of Westfall (2041) | Quests [quest] | 22.0 | yes | Trapper's Leather Armor (252491, -0.48 DPS) [crafted]; Dark Leather Tunic (2317, -0.60 DPS) [crafted]; Brawler's Leather Armor (252490, -0.64 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 10.0 | yes | Bravo's Armbands (270015, -0.12 DPS) [quest]; Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.24 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 97.7 | yes | Serpent Gloves (5970, +0.41 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -5.12 DPS) [dungeon]; Forest Leather Gloves (3058, -5.36 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.46 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.48 DPS) [quest]; Guardsman Belt (3429, -0.60 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.0 | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.12 DPS) [world]; Brawler's Leather Pants (252500, -0.35 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 16.0 | yes | Footpads of the Fang (10411, -0.24 DPS) [dungeon]; Blackened Defias Boots (10402, -0.26 DPS, sim-verified) [dungeon]; Dark Leather Boots (2315, -0.36 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.0 | yes | Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon]; The 1 Ring (8350, -0.60 DPS) [world]; Minor Channeling Ring (1449, -0.72 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.0 | yes | Lavishly Jeweled Ring (1156, -0.08 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.36 DPS) [world]; Minor Channeling Ring (1449, -0.48 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 22.0 | yes | Impaling Harpoon (5200, +0.03 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.48 DPS) [world]; Lupine Axe (1220, -0.60 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 178.3 | yes | Lovingly Crafted Boomstick (4372, -2.52 DPS) [crafted]; Venomstrike (6469, -2.64 DPS) [dungeon]; Cracked Blacksmith Hammer (285279, -3.58 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 384, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 40 (dwarf, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 102.1. Weights run: 0.7s. Verify run: 1.3s. 1253 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.006, agility=2.000 ± 0.012, crit=13.274 ± 2.589, hit=not significant (6.955 ± 2.095), melee_haste=not significant (8.860 ± 5.066)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 185.8 | yes | Nightscape Headband (8176, +0.87 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -9.44 DPS) [crafted]; White Bandit Mask (10008, -9.56 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 22.0 | yes | Sentinel's Medallion (19541, -0.42 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.47 DPS) [dungeon]; Sentinel's Medallion (20444, -0.58 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 34.0 | yes | Nightscape Shoulders (8192, -0.70 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.77 DPS, sim-verified) [dungeon]; Mantle of Thieves (2264, -0.82 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | Gnomeregan: Caverndeep Burrower [dungeon] | 16.0 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Yeti Fur Cloak (2805, -0.23 DPS) [quest]; Darktide Cape (4114, -0.23 DPS) [quest] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 30.0 | yes | Tough Scorpid Breastplate (8203, +0.81 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.12 DPS) [crafted]; Hawkeye's Tunic (14592, -0.35 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.13 DPS, sim-verified) [dungeon]; Dusky Bracers (7378, -0.23 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 205.8 | yes | Dragonscale Gauntlets (8347, -0.43 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 193.8 | yes | Highlander's Leather Girdle (20116, +0.55 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -9.91 DPS) [rep]; Highlander's Leather Girdle (20117, -9.91 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | Gnomeregan: Leprous Machinesmith [dungeon] | 42.0 | yes | Triprunner Dungarees (9624, -0.58 DPS, sim-verified) [quest]; Hawkeye's Breeches (14595, -0.82 DPS) [world]; Ferine Leggings (6690, -0.93 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 26.0 | yes | Dusky Boots (7390, -0.23 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.23 DPS) [crafted]; Imperial Leather Boots (6431, -0.28 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 20.0 | yes | Protector's Band (19515, -0.23 DPS) [rep]; Disengagement Ring (276202, -0.23 DPS) [vendor]; Monkey Ring (6748, -0.35 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.0 | yes | Disengagement Ring (276202, -0.12 DPS) [vendor]; Protector's Band (19515, -0.14 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.23 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Frost Tiger Blade (3854) (or Illusionary Rod (7713)) | Blacksmithing [crafted] | 185.8 | yes | Illusionary Rod (7713, +0.42 DPS, sim-verified) [dungeon]; Steel Spear (250605, -8.74 DPS) [crafted]; Loksey's Training Stick (7710, -9.21 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) | Gnomeregan: Dark Iron Agent [dungeon] | 282.4 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS) [crafted]; Master Hunter's Rifle (17687, -0.15 DPS) [quest]; Master Hunter's Bow (17686, -1.20 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Frost Tiger Blade; ranged: Sniper Rifle

No-known-source sample (15 of 1253, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 60 (dwarf, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 240.7. Weights run: 0.7s. Verify run: 1.3s. 2330 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=18.230 ± 3.490, hit=not significant (0.000 ± 0.000), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Quests [quest] | 572.4 | yes | Lieutenant Commander's Chain Helm (23306, -1.48 DPS) [vendor]; Lieutenant Commander's Chain Helm (227066, -1.48 DPS) [pvp]; Champion's Chain Helm (23251, -3.81 DPS, sim-verified) [vendor] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 536.4 | yes | Gem of Trapped Innocents (23057, -1.48 DPS) [raid]; Barbed Choker (21664, -13.54 DPS) [raid]; Medallion of the Dawn (22659, -14.69 DPS) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Quests [quest] | 313.2 | yes | Lieutenant Commander's Chain Pauldrons (227084, -1.14 DPS) [pvp]; Champion's Chain Shoulders (23252, -1.26 DPS) [vendor]; Champion's Chain Pauldrons (227078, -3.97 DPS, sim-verified) [pvp] |
| back | Cloak of the Fallen God (21710) | Quests [quest] | 52.0 | yes | Cape of the Black Baron (13340, -0.11 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.46 DPS) [rep]; Chromatic Cloak (18509, -3.26 DPS, sim-verified) [crafted] |
| chest | Legionnaire's Chain Hauberk (22874) (or Knight-Captain's Chain Hauberk (23292), Legionnaire's Chain Hauberk (227071)) | Lady Palanseer [vendor] | 542.4 | yes | Legionnaire's Chain Hauberk (227071, +0.00 DPS) [pvp]; Bloodsoul Breastplate (19690, -0.80 DPS) [crafted]; Knight-Captain's Chain Hauberk (23292, -3.97 DPS, sim-verified) [vendor] |
| wrist | Cryptstalker Wristguards (22443) | Quests [quest] | 52.0 | yes | General's Chain Wristguards (16570, -0.69 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.80 DPS) [rep]; Marshal's Chain Bracers (16461, -2.35 DPS, sim-verified) [pvp] |
| hands | Cryptstalker Handguards (22441) | Quests [quest] | 303.2 | yes | Chromatic Gauntlets (19157, -0.23 DPS) [crafted]; General's Chain Gloves (16571, -0.34 DPS) [vendor]; Marshal's Chain Grips (16463, -3.20 DPS, sim-verified) [vendor] |
| waist | Cryptstalker Girdle (22442) | Quests [quest] | 301.2 | yes | Highlander's Chain Girdle (20043, -0.69 DPS) [rep]; Highlander's Leather Girdle (20045, -0.69 DPS) [rep]; Belt of Never-ending Agony (21586, -11.69 DPS, sim-verified) [raid] |
| legs | Legionnaire's Chain Legguards (22875) (or Knight-Captain's Chain Legguards (23293), Knight-Captain's Chain Legguards (227072), Legionnaire's Chain Legguards (227073)) | Lady Palanseer [vendor] | 542.4 | yes | Knight-Captain's Chain Legguards (227072, +0.00 DPS) [pvp]; Legionnaire's Chain Legguards (227073, +0.00 DPS) [pvp]; Knight-Captain's Chain Legguards (23293, -3.97 DPS, sim-verified) [vendor] |
| feet | Cryptstalker Boots (22440) | Quests [quest] | 66.0 | yes | General's Chain Sabatons (231564, -0.69 DPS) [pvp]; Marshal's Chain Boots (16462, -0.80 DPS) [vendor]; Striker's Footguards (21365, -1.74 DPS, sim-verified) [quest] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 307.2 | yes | Quick Strike Ring (18821, -1.26 DPS) [raid]; Don Julio's Band (19325, -2.06 DPS) [rep]; Band of the Penitent (13217, -2.97 DPS) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj [raid] | 295.2 | yes | Quick Strike Ring (18821, -0.82 DPS, sim-verified) [raid]; Don Julio's Band (19325, -1.37 DPS) [rep]; Band of the Penitent (13217, -2.28 DPS) [quest] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Drake Fang Talisman (19406, -0.46 DPS) [raid]; Thunderbrew's Boot Flask (744, -3.65 DPS) [quest]; Eye of Diminution (23001, -5.89 DPS, sim-verified) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 255.2 | yes | Eye of Diminution (23001, -4.65 DPS, sim-verified) [raid]; Drake Fang Talisman (19406, -11.37 DPS) [raid]; Thunderbrew's Boot Flask (744, -14.57 DPS) [quest] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 283.2 | yes | Ironbark Staff (20069, +12.97 DPS) [rep]; Atiesh, Greatstaff of the Guardian (22630, +12.97 DPS) [quest]; High Warlord's War Staff (234549, +12.97 DPS) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 283.2 | yes | Grand Marshal's Handaxe (18827, +0.00 DPS, sim-verified) [vendor]; High Warlord's Cleaver (18828, +0.00 DPS) [vendor]; Grand Marshal's Dirk (18838, +0.00 DPS) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 961.6 | yes | The Purifier (22656, -13.29 DPS) [quest]; Huhuran's Stinger (21616, -15.84 DPS) [raid]; High Warlord's Recurve (234559, -16.39 DPS) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Stormrage's Talisman of Seething; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Hauberk; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legguards; feet: Cryptstalker Boots; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2330, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 74.6. Weights run: 0.7s. Verify run: 1.0s. 379 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.004, agility=2.000 ± 0.009, crit=6.978 ± 1.340, hit=not significant (4.521 ± 1.506), melee_haste=not significant (6.060 ± 4.886)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.0 | yes | Flying Tiger Goggles (4368, -0.91 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.96 DPS) [crafted]; Lucky Fishing Hat (19972, -0.96 DPS) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.0 | yes | Tarnished Locket (279870, -0.65 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.59 DPS, sim-verified) [crafted]; Reinforced Woolen Shoulders (4315, -0.60 DPS) [crafted]; Forest Leather Mantle (4709, -0.60 DPS) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 | yes | Cape of the Brotherhood (5193, -0.23 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.24 DPS) [dungeon]; Hide of Lupos (3018, -0.24 DPS) [world] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 14.0 | yes | Trapper's Leather Armor (252491, +0.40 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.12 DPS) [crafted]; Heckler's Hide (286536, -0.24 DPS) [world] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 10.0 | yes | Bravo's Armbands (270015, -0.12 DPS) [quest]; Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.24 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 97.7 | yes | Serpent Gloves (5970, +0.42 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -5.12 DPS) [dungeon]; Forest Leather Gloves (3058, -5.36 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.48 DPS) [quest]; Guardsman Belt (3429, -0.60 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.0 | yes | Brawler's Leather Pants (252500, +0.12 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.12 DPS) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 16.0 | yes | Footpads of the Fang (10411, -0.24 DPS) [dungeon]; Blackened Defias Boots (10402, -0.25 DPS, sim-verified) [dungeon]; Dark Leather Boots (2315, -0.36 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.0 | yes | Bounty Hunter's Ring (5351, -0.36 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon]; The 1 Ring (8350, -0.60 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 | yes | Bounty Hunter's Ring (5351, -0.13 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; The 1 Ring (8350, -0.36 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 22.0 | yes | Impaling Harpoon (5200, +0.55 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.48 DPS) [world]; Crescent Staff (6505, -0.48 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 178.3 | yes | Lovingly Crafted Boomstick (4372, -2.52 DPS) [crafted]; Venomstrike (6469, -2.64 DPS) [dungeon]; Cracked Blacksmith Hammer (285279, -3.23 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 379, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 40 (troll, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 103.7. Weights run: 0.7s. Verify run: 1.2s. 1248 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.006, agility=2.000 ± 0.012, crit=13.274 ± 2.589, hit=not significant (6.955 ± 2.095), melee_haste=not significant (8.860 ± 5.066)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 185.8 | yes | Nightscape Headband (8176, +0.78 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -9.44 DPS) [crafted]; White Bandit Mask (10008, -9.56 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 22.0 | yes | Scout's Medallion (19537, -0.45 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.47 DPS) [dungeon]; Scout's Medallion (20442, -0.58 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 34.0 | yes | Nightscape Shoulders (8192, -0.70 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.76 DPS, sim-verified) [dungeon]; Mantle of Thieves (2264, -0.82 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | Gnomeregan: Caverndeep Burrower [dungeon] | 16.0 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Darktide Cape (4114, -0.23 DPS) [quest]; Cloak of Night (4447, -0.23 DPS) [world] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 30.0 | yes | Tough Scorpid Breastplate (8203, +0.64 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.12 DPS) [crafted]; Hawkeye's Tunic (14592, -0.35 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.11 DPS, sim-verified) [dungeon]; Dusky Bracers (7378, -0.23 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 205.8 | yes | Dragonscale Gauntlets (8347, -0.40 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 193.8 | yes | Defiler's Leather Girdle (20192, +0.43 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -9.91 DPS) [rep]; Defiler's Leather Girdle (20191, -9.91 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | Gnomeregan: Leprous Machinesmith [dungeon] | 42.0 | yes | Triprunner Dungarees (9624, -0.30 DPS, sim-verified) [quest]; Hawkeye's Breeches (14595, -0.82 DPS) [world]; Ferine Leggings (6690, -0.93 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 26.0 | yes | Dusky Boots (7390, -0.23 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.23 DPS) [crafted]; Imperial Leather Boots (6431, -0.30 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 20.0 | yes | Legionnaire's Band (19512, -0.23 DPS) [rep]; Disengagement Ring (276202, -0.23 DPS) [vendor]; Monkey Ring (6748, -0.35 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.0 | yes | Disengagement Ring (276202, -0.12 DPS) [vendor]; Legionnaire's Band (19512, -0.14 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.23 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 185.8 | yes | Frost Tiger Blade (3854, -1.23 DPS, sim-verified) [crafted]; Steel Spear (250605, -8.74 DPS) [crafted]; Loksey's Training Stick (7710, -9.21 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) | Gnomeregan: Dark Iron Agent [dungeon] | 282.4 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS) [crafted]; Master Hunter's Rifle (17687, -0.15 DPS) [quest]; Master Hunter's Bow (17686, -2.18 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Sniper Rifle

No-known-source sample (15 of 1248, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 60 (troll, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 241.5. Weights run: 0.7s. Verify run: 1.3s. 2324 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=18.230 ± 3.490, hit=not significant (0.000 ± 0.000), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Quests [quest] | 572.4 | yes | Lieutenant Commander's Chain Helm (23306, -1.48 DPS) [vendor]; Lieutenant Commander's Chain Helm (227066, -1.48 DPS) [pvp]; Champion's Chain Helm (23251, -4.43 DPS, sim-verified) [vendor] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 536.4 | yes | Gem of Trapped Innocents (23057, -1.48 DPS) [raid]; Barbed Choker (21664, -13.54 DPS) [raid]; Medallion of the Dawn (22659, -14.69 DPS) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Quests [quest] | 313.2 | yes | Lieutenant Commander's Chain Pauldrons (227084, -1.14 DPS) [pvp]; Champion's Chain Shoulders (23252, -1.26 DPS) [vendor]; Champion's Chain Pauldrons (227078, -3.55 DPS, sim-verified) [pvp] |
| back | Cloak of the Fallen God (21710) | Quests [quest] | 52.0 | yes | Cape of the Black Baron (13340, -0.11 DPS) [dungeon]; Deathguard's Cloak (20068, -0.46 DPS) [rep]; Chromatic Cloak (18509, -3.25 DPS, sim-verified) [crafted] |
| chest | Legionnaire's Chain Hauberk (22874) (or Knight-Captain's Chain Hauberk (23292), Legionnaire's Chain Hauberk (227071)) | Lady Palanseer [vendor] | 542.4 | yes | Legionnaire's Chain Hauberk (227071, +0.00 DPS) [pvp]; Bloodsoul Breastplate (19690, -0.80 DPS) [crafted]; Knight-Captain's Chain Hauberk (23292, -4.02 DPS, sim-verified) [vendor] |
| wrist | Cryptstalker Wristguards (22443) | Quests [quest] | 52.0 | yes | General's Chain Wristguards (16570, -0.69 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.80 DPS) [rep]; Marshal's Chain Bracers (16461, -2.44 DPS, sim-verified) [pvp] |
| hands | Cryptstalker Handguards (22441) | Quests [quest] | 303.2 | yes | Chromatic Gauntlets (19157, -0.23 DPS) [crafted]; General's Chain Gloves (16571, -0.34 DPS) [vendor]; Marshal's Chain Grips (16463, -2.12 DPS, sim-verified) [vendor] |
| waist | Cryptstalker Girdle (22442) | Quests [quest] | 301.2 | yes | Defiler's Chain Girdle (20150, -0.69 DPS) [rep]; Defiler's Leather Girdle (20190, -0.69 DPS) [rep]; Belt of Never-ending Agony (21586, -11.38 DPS, sim-verified) [raid] |
| legs | Legionnaire's Chain Legguards (22875) (or Knight-Captain's Chain Legguards (23293), Knight-Captain's Chain Legguards (227072), Legionnaire's Chain Legguards (227073)) | Lady Palanseer [vendor] | 542.4 | yes | Knight-Captain's Chain Legguards (227072, +0.00 DPS) [pvp]; Legionnaire's Chain Legguards (227073, +0.00 DPS) [pvp]; Knight-Captain's Chain Legguards (23293, -4.02 DPS, sim-verified) [vendor] |
| feet | Cryptstalker Boots (22440) | Quests [quest] | 66.0 | yes | General's Chain Sabatons (231564, -0.69 DPS) [pvp]; Marshal's Chain Boots (16462, -0.80 DPS) [vendor]; Striker's Footguards (21365, -2.42 DPS, sim-verified) [quest] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 307.2 | yes | Quick Strike Ring (18821, -1.26 DPS) [raid]; Don Julio's Band (19325, -2.06 DPS) [rep]; Band of the Penitent (13217, -2.97 DPS) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj [raid] | 295.2 | yes | Quick Strike Ring (18821, -0.82 DPS, sim-verified) [raid]; Don Julio's Band (19325, -1.37 DPS) [rep]; Band of the Penitent (13217, -2.28 DPS) [quest] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Rune of the Guard Captain (19120, +1.14 DPS) [quest]; Drake Fang Talisman (19406, -0.46 DPS) [raid]; Eye of Diminution (23001, -6.43 DPS, sim-verified) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 255.2 | yes | Eye of Diminution (23001, -5.46 DPS, sim-verified) [raid]; Rune of the Guard Captain (19120, -9.78 DPS) [quest]; Drake Fang Talisman (19406, -11.37 DPS) [raid] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 283.2 | yes | Ironbark Staff (20220, +12.97 DPS) [rep]; Atiesh, Greatstaff of the Guardian (22630, +12.97 DPS) [quest]; High Warlord's War Staff (234549, +12.97 DPS) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 283.2 | yes | Grand Marshal's Handaxe (18827, +0.00 DPS, sim-verified) [vendor]; High Warlord's Cleaver (18828, +0.00 DPS) [vendor]; Grand Marshal's Dirk (18838, +0.00 DPS) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 961.6 | yes | The Purifier (22656, -13.29 DPS) [quest]; Huhuran's Stinger (21616, -15.84 DPS) [raid]; High Warlord's Recurve (234559, -16.39 DPS) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Stormrage's Talisman of Seething; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Hauberk; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legguards; feet: Cryptstalker Boots; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2324, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

