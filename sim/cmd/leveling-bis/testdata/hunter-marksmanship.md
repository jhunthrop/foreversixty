# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (dwarf, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.1. Weights run: 0.7s. Verify run: 1.0s. 384 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.004, agility=2.000 ± 0.009, crit=6.978 ± 1.340, hit=not significant (4.521 ± 1.506), melee_haste=not significant (6.060 ± 4.886)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.0 | yes | Flying Tiger Goggles (4368, -16.0) [crafted]; Shadow Goggles (4373, -16.0) [crafted]; Lucky Fishing Hat (19972, -16.0) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.0 | yes | Tarnished Locket (279870, -12.0) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 | yes | Double-Stitched Woolen Shoulders (4314, -10.0) [crafted]; Reinforced Woolen Shoulders (4315, -10.0) [crafted]; Forest Leather Mantle (4709, -10.0) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 | yes | Cape of the Brotherhood (5193, -2.0) [dungeon]; Sentry Cloak (2059, -4.0) [dungeon]; Hide of Lupos (3018, -4.0) [world] |
| chest | Tunic of Westfall (2041) | Quests [quest] | 22.0 | yes | Brawler's Leather Armor (252490, -8.0) [crafted]; Trapper's Leather Armor (252491, -8.0) [crafted]; Dark Leather Tunic (2317, -10.0) [crafted] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 10.0 | yes | Wolf Bracers (4794, -2.0) [vendor]; Bravo's Armbands (270015, -2.0) [quest]; Ratchet Wristwraps (274742, -4.0) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 97.7 | yes | Serpent Gloves (5970, -85.7) [dungeon]; Gloves of the Fang (10413, -85.7) [dungeon]; Forest Leather Gloves (3058, -89.7) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -8.0) [crafted]; Dusty Belt (279897, -8.0) [quest]; Guardsman Belt (3429, -10.0) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.0 | yes | Brawler's Leather Pants (252500, +0.0) [crafted]; Trapper's Leather Pants (252501, +0.0) [crafted]; Bluegill Breeches (3022, -2.0) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 16.0 | yes | Blackened Defias Boots (10402, -4.0) [dungeon]; Footpads of the Fang (10411, -4.0) [dungeon]; Dark Leather Boots (2315, -6.0) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.0 | yes | Lavishly Jeweled Ring (1156, -8.0) [dungeon]; The 1 Ring (8350, -10.0) [world]; Minor Channeling Ring (1449, -12.0) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.0 | yes | Lavishly Jeweled Ring (1156, -4.0) [dungeon]; The 1 Ring (8350, -6.0) [world]; Minor Channeling Ring (1449, -8.0) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 22.0 | yes | Impaling Harpoon (5200, -4.0) [dungeon]; Scythe Axe (5749, -8.0) [world]; Lupine Axe (1220, -10.0) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 178.3 | yes | Cracked Blacksmith Hammer (285279, -38.3) [crafted]; Lovingly Crafted Boomstick (4372, -42.2) [crafted]; Venomstrike (6469, -44.1) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 384, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 40 (dwarf, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 101.3. Weights run: 0.6s. Verify run: 1.3s. 1266 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.006, agility=2.000 ± 0.012, crit=13.274 ± 2.589, hit=not significant (6.955 ± 2.095), melee_haste=not significant (8.860 ± 5.066)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 185.8 | yes | Nightscape Headband (8176, -161.8) [crafted]; Guard's Chain Helm (250499, -161.8) [crafted]; White Bandit Mask (10008, -163.8) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 22.0 | yes | Sentinel's Medallion (19541, -6.0) [rep]; Ghostshard Talisman (7731, -8.0) [dungeon]; Sentinel's Medallion (20444, -10.0) [rep] |
| shoulder | Forest Tracker Epaulets (2278) (or Nightscape Shoulders (8192)) | Gnomeregan: Mechanized Sentry [dungeon] | 22.0 | yes | Nightscape Shoulders (8192, +0.0) [crafted]; Mantle of Thieves (2264, -2.0) [dungeon]; Hawkeye's Epaulets (14596, -4.0) [world] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | Gnomeregan: Caverndeep Burrower [dungeon] | 16.0 | yes | Parachute Cloak (10518, +0.0) [crafted]; Yeti Fur Cloak (2805, -4.0) [quest]; Darktide Cape (4114, -4.0) [quest] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 30.0 | yes | Tough Scorpid Breastplate (8203, +0.0) [crafted]; Dusky Leather Armor (7374, -2.0) [crafted]; Hawkeye's Tunic (14592, -6.0) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -4.0) [dungeon]; Dusky Bracers (7378, -4.0) [crafted]; Tough Scorpid Bracers (8205, -6.0) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 205.8 | yes | Dragonscale Gauntlets (8347, -8.0) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 193.8 | yes | Highlander's Leather Girdle (20116, -163.8) [rep]; Highlander's Chain Girdle (20090, -169.8) [rep]; Highlander's Leather Girdle (20117, -169.8) [rep] |
| legs | Basilisk Hide Pants (1718) | Gnomeregan: Leprous Machinesmith [dungeon] | 42.0 | yes | Triprunner Dungarees (9624, -6.0) [quest]; Hawkeye's Breeches (14595, -14.0) [world]; Ferine Leggings (6690, -16.0) [dungeon] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 26.0 | yes | Imperial Leather Boots (6431, -4.0) [dungeon]; Dusky Boots (7390, -4.0) [crafted]; Skulker's Leather Shoes (252531, -4.0) [crafted] |
| finger1 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 20.0 | yes | Protector's Band (19515, -4.0) [rep]; Monkey Ring (6748, -6.0) [quest]; Ring of Precision (1491, -8.0) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.0 | yes | Protector's Band (19515, -2.0) [rep]; Monkey Ring (6748, -4.0) [quest]; Ring of Precision (1491, -6.0) [dungeon] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Frost Tiger Blade (3854) (or Illusionary Rod (7713)) | Blacksmithing [crafted] | 185.8 | yes | Illusionary Rod (7713, +0.0) [dungeon]; Steel Spear (250605, -149.8) [crafted]; Loksey's Training Stick (7710, -157.8) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) | Gnomeregan: Dark Iron Agent [dungeon] | 282.4 | yes | Master Hunter's Bow (17686, +0.9) [quest]; Mithril Heavy-bore Rifle (10510, +0.0) [crafted]; Master Hunter's Rifle (17687, -2.5) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Frost Tiger Blade; ranged: Sniper Rifle

No-known-source sample (15 of 1266, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 60 (dwarf, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 240.7. Weights run: 0.7s. Verify run: 1.3s. 2418 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=18.230 ± 3.490, hit=not significant (0.000 ± 0.000), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Quests [quest] | 572.4 | yes | Champion's Chain Helm (23251, -26.0) [vendor]; Lieutenant Commander's Chain Helm (23306, -26.0) [vendor]; Lieutenant Commander's Chain Helm (227066, -26.0) [pvp] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 536.4 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Barbed Choker (21664, -237.2) [raid]; Medallion of the Dawn (22659, -257.2) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Quests [quest] | 313.2 | yes | Champion's Chain Pauldrons (227078, -20.0) [pvp]; Lieutenant Commander's Chain Pauldrons (227084, -20.0) [pvp]; Champion's Chain Shoulders (23252, -22.0) [vendor] |
| back | Cloak of the Fallen God (21710) | Quests [quest] | 52.0 | yes | Chromatic Cloak (18509, +203.2) [crafted]; Cape of the Black Baron (13340, -2.0) [dungeon]; Cloak of the Honor Guard (20073, -8.0) [rep] |
| chest | Legionnaire's Chain Hauberk (22874) (or Knight-Captain's Chain Hauberk (23292), Legionnaire's Chain Hauberk (227071)) | Lady Palanseer [vendor] | 542.4 | yes | Knight-Captain's Chain Hauberk (23292, +0.0) [vendor]; Legionnaire's Chain Hauberk (227071, +0.0) [pvp]; Bloodsoul Breastplate (19690, -14.0) [crafted] |
| wrist | Cryptstalker Wristguards (22443) | Quests [quest] | 52.0 | yes | Marshal's Chain Bracers (16461, -12.0) [pvp]; General's Chain Wristguards (16570, -12.0) [pvp]; Forest Stalker's Bracers (19587, -14.0) [rep] |
| hands | Cryptstalker Handguards (22441) | Quests [quest] | 303.2 | yes | Chromatic Gauntlets (19157, -4.0) [crafted]; Marshal's Chain Grips (16463, -6.0) [vendor]; General's Chain Gloves (16571, -6.0) [vendor] |
| waist | Cryptstalker Girdle (22442) | Quests [quest] | 301.2 | yes | Belt of Never-ending Agony (21586, +18.0) [raid]; Highlander's Chain Girdle (20043, -12.0) [rep]; Highlander's Leather Girdle (20045, -12.0) [rep] |
| legs | Legionnaire's Chain Legguards (22875) (or Knight-Captain's Chain Legguards (23293), Knight-Captain's Chain Legguards (227072), Legionnaire's Chain Legguards (227073)) | Lady Palanseer [vendor] | 542.4 | yes | Knight-Captain's Chain Legguards (23293, +0.0) [vendor]; Knight-Captain's Chain Legguards (227072, +0.0) [pvp]; Legionnaire's Chain Legguards (227073, +0.0) [pvp] |
| feet | Cryptstalker Boots (22440) | Quests [quest] | 66.0 | yes | Striker's Footguards (21365, -4.0) [quest]; General's Chain Sabatons (231564, -12.0) [pvp]; Marshal's Chain Boots (16462, -14.0) [vendor] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 307.2 | yes | Quick Strike Ring (18821, -22.0) [raid]; Don Julio's Band (19325, -36.0) [rep]; Band of the Penitent (13217, -52.0) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj [raid] | 295.2 | yes | Quick Strike Ring (18821, -10.0) [raid]; Don Julio's Band (19325, -24.0) [rep]; Band of the Penitent (13217, -40.0) [quest] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +446.4) [raid]; Drake Fang Talisman (19406, -8.0) [raid]; Thunderbrew's Boot Flask (744, -64.0) [quest] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 255.2 | yes | Eye of Diminution (23001, +255.2) [raid]; Drake Fang Talisman (19406, -199.2) [raid]; Thunderbrew's Boot Flask (744, -255.2) [quest] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 283.2 | yes | Ironbark Staff (20069, +227.2) [rep]; Atiesh, Greatstaff of the Guardian (22630, +227.2) [quest]; High Warlord's War Staff (234549, +227.2) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 283.2 | yes | Grand Marshal's Handaxe (18827, +0.0) [vendor]; High Warlord's Cleaver (18828, +0.0) [vendor]; Grand Marshal's Dirk (18838, +0.0) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 961.6 | yes | The Purifier (22656, -232.8) [quest]; Huhuran's Stinger (21616, -277.4) [raid]; High Warlord's Recurve (234559, -287.1) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Stormrage's Talisman of Seething; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Hauberk; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legguards; feet: Cryptstalker Boots; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2418, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 74.6. Weights run: 0.7s. Verify run: 1.0s. 379 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.004, agility=2.000 ± 0.009, crit=6.978 ± 1.340, hit=not significant (4.521 ± 1.506), melee_haste=not significant (6.060 ± 4.886)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.0 | yes | Flying Tiger Goggles (4368, -16.0) [crafted]; Shadow Goggles (4373, -16.0) [crafted]; Lucky Fishing Hat (19972, -16.0) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.0 | yes | Tarnished Locket (279870, -12.0) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.0 | yes | Double-Stitched Woolen Shoulders (4314, -10.0) [crafted]; Reinforced Woolen Shoulders (4315, -10.0) [crafted]; Forest Leather Mantle (4709, -10.0) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.0 | yes | Cape of the Brotherhood (5193, -2.0) [dungeon]; Sentry Cloak (2059, -4.0) [dungeon]; Hide of Lupos (3018, -4.0) [world] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 14.0 | yes | Trapper's Leather Armor (252491, +0.0) [crafted]; Dark Leather Tunic (2317, -2.0) [crafted]; Heckler's Hide (286536, -4.0) [world] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 10.0 | yes | Wolf Bracers (4794, -2.0) [vendor]; Bravo's Armbands (270015, -2.0) [quest]; Ratchet Wristwraps (274742, -4.0) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 97.7 | yes | Serpent Gloves (5970, -85.7) [dungeon]; Gloves of the Fang (10413, -85.7) [dungeon]; Forest Leather Gloves (3058, -89.7) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -8.0) [crafted]; Dusty Belt (279897, -8.0) [quest]; Guardsman Belt (3429, -10.0) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.0 | yes | Brawler's Leather Pants (252500, +0.0) [crafted]; Trapper's Leather Pants (252501, +0.0) [crafted]; Bluegill Breeches (3022, -2.0) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 16.0 | yes | Blackened Defias Boots (10402, -4.0) [dungeon]; Footpads of the Fang (10411, -4.0) [dungeon]; Dark Leather Boots (2315, -6.0) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.0 | yes | Bounty Hunter's Ring (5351, -6.0) [quest]; Lavishly Jeweled Ring (1156, -8.0) [dungeon]; The 1 Ring (8350, -10.0) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 | yes | Bounty Hunter's Ring (5351, -2.0) [quest]; Lavishly Jeweled Ring (1156, -4.0) [dungeon]; The 1 Ring (8350, -6.0) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 22.0 | yes | Impaling Harpoon (5200, -4.0) [dungeon]; Scythe Axe (5749, -8.0) [world]; Crescent Staff (6505, -8.0) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 178.3 | yes | Cracked Blacksmith Hammer (285279, -38.3) [crafted]; Lovingly Crafted Boomstick (4372, -42.2) [crafted]; Venomstrike (6469, -44.1) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 379, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 40 (troll, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 103.0. Weights run: 0.6s. Verify run: 1.2s. 1261 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.006, agility=2.000 ± 0.012, crit=13.274 ± 2.589, hit=not significant (6.955 ± 2.095), melee_haste=not significant (8.860 ± 5.066)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 185.8 | yes | Nightscape Headband (8176, -161.8) [crafted]; Guard's Chain Helm (250499, -161.8) [crafted]; White Bandit Mask (10008, -163.8) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 22.0 | yes | Scout's Medallion (19537, -6.0) [rep]; Ghostshard Talisman (7731, -8.0) [dungeon]; Scout's Medallion (20442, -10.0) [rep] |
| shoulder | Forest Tracker Epaulets (2278) (or Nightscape Shoulders (8192)) | Gnomeregan: Mechanized Sentry [dungeon] | 22.0 | yes | Nightscape Shoulders (8192, +0.0) [crafted]; Mantle of Thieves (2264, -2.0) [dungeon]; Hawkeye's Epaulets (14596, -4.0) [world] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | Gnomeregan: Caverndeep Burrower [dungeon] | 16.0 | yes | Parachute Cloak (10518, +0.0) [crafted]; Darktide Cape (4114, -4.0) [quest]; Cloak of Night (4447, -4.0) [world] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 30.0 | yes | Tough Scorpid Breastplate (8203, +0.0) [crafted]; Dusky Leather Armor (7374, -2.0) [crafted]; Hawkeye's Tunic (14592, -6.0) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -4.0) [dungeon]; Dusky Bracers (7378, -4.0) [crafted]; Tough Scorpid Bracers (8205, -6.0) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 205.8 | yes | Dragonscale Gauntlets (8347, -8.0) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 193.8 | yes | Defiler's Leather Girdle (20192, -163.8) [rep]; Defiler's Chain Girdle (20152, -169.8) [rep]; Defiler's Leather Girdle (20191, -169.8) [rep] |
| legs | Basilisk Hide Pants (1718) | Gnomeregan: Leprous Machinesmith [dungeon] | 42.0 | yes | Triprunner Dungarees (9624, -6.0) [quest]; Hawkeye's Breeches (14595, -14.0) [world]; Ferine Leggings (6690, -16.0) [dungeon] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 26.0 | yes | Imperial Leather Boots (6431, -4.0) [dungeon]; Dusky Boots (7390, -4.0) [crafted]; Skulker's Leather Shoes (252531, -4.0) [crafted] |
| finger1 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 20.0 | yes | Legionnaire's Band (19512, -4.0) [rep]; Monkey Ring (6748, -6.0) [quest]; Ring of Precision (1491, -8.0) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.0 | yes | Legionnaire's Band (19512, -2.0) [rep]; Monkey Ring (6748, -4.0) [quest]; Ring of Precision (1491, -6.0) [dungeon] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 185.8 | yes | Frost Tiger Blade (3854, +0.0) [crafted]; Steel Spear (250605, -149.8) [crafted]; Loksey's Training Stick (7710, -157.8) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) | Gnomeregan: Dark Iron Agent [dungeon] | 282.4 | yes | Master Hunter's Bow (17686, +0.9) [quest]; Mithril Heavy-bore Rifle (10510, +0.0) [crafted]; Master Hunter's Rifle (17687, -2.5) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Sniper Rifle

No-known-source sample (15 of 1261, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 60 (troll, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 241.5. Weights run: 0.7s. Verify run: 1.3s. 2412 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=18.230 ± 3.490, hit=not significant (0.000 ± 0.000), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Quests [quest] | 572.4 | yes | Champion's Chain Helm (23251, -26.0) [vendor]; Lieutenant Commander's Chain Helm (23306, -26.0) [vendor]; Lieutenant Commander's Chain Helm (227066, -26.0) [pvp] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 536.4 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Barbed Choker (21664, -237.2) [raid]; Medallion of the Dawn (22659, -257.2) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Quests [quest] | 313.2 | yes | Champion's Chain Pauldrons (227078, -20.0) [pvp]; Lieutenant Commander's Chain Pauldrons (227084, -20.0) [pvp]; Champion's Chain Shoulders (23252, -22.0) [vendor] |
| back | Cloak of the Fallen God (21710) | Quests [quest] | 52.0 | yes | Chromatic Cloak (18509, +203.2) [crafted]; Cape of the Black Baron (13340, -2.0) [dungeon]; Deathguard's Cloak (20068, -8.0) [rep] |
| chest | Legionnaire's Chain Hauberk (22874) (or Knight-Captain's Chain Hauberk (23292), Legionnaire's Chain Hauberk (227071)) | Lady Palanseer [vendor] | 542.4 | yes | Knight-Captain's Chain Hauberk (23292, +0.0) [vendor]; Legionnaire's Chain Hauberk (227071, +0.0) [pvp]; Bloodsoul Breastplate (19690, -14.0) [crafted] |
| wrist | Cryptstalker Wristguards (22443) | Quests [quest] | 52.0 | yes | Marshal's Chain Bracers (16461, -12.0) [pvp]; General's Chain Wristguards (16570, -12.0) [pvp]; Forest Stalker's Bracers (19587, -14.0) [rep] |
| hands | Cryptstalker Handguards (22441) | Quests [quest] | 303.2 | yes | Chromatic Gauntlets (19157, -4.0) [crafted]; Marshal's Chain Grips (16463, -6.0) [vendor]; General's Chain Gloves (16571, -6.0) [vendor] |
| waist | Cryptstalker Girdle (22442) | Quests [quest] | 301.2 | yes | Belt of Never-ending Agony (21586, +18.0) [raid]; Defiler's Chain Girdle (20150, -12.0) [rep]; Defiler's Leather Girdle (20190, -12.0) [rep] |
| legs | Legionnaire's Chain Legguards (22875) (or Knight-Captain's Chain Legguards (23293), Knight-Captain's Chain Legguards (227072), Legionnaire's Chain Legguards (227073)) | Lady Palanseer [vendor] | 542.4 | yes | Knight-Captain's Chain Legguards (23293, +0.0) [vendor]; Knight-Captain's Chain Legguards (227072, +0.0) [pvp]; Legionnaire's Chain Legguards (227073, +0.0) [pvp] |
| feet | Cryptstalker Boots (22440) | Quests [quest] | 66.0 | yes | Striker's Footguards (21365, -4.0) [quest]; General's Chain Sabatons (231564, -12.0) [pvp]; Marshal's Chain Boots (16462, -14.0) [vendor] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 307.2 | yes | Quick Strike Ring (18821, -22.0) [raid]; Don Julio's Band (19325, -36.0) [rep]; Band of the Penitent (13217, -52.0) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj [raid] | 295.2 | yes | Quick Strike Ring (18821, -10.0) [raid]; Don Julio's Band (19325, -24.0) [rep]; Band of the Penitent (13217, -40.0) [quest] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +446.4) [raid]; Rune of the Guard Captain (19120, +20.0) [quest]; Drake Fang Talisman (19406, -8.0) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 255.2 | yes | Eye of Diminution (23001, +255.2) [raid]; Rune of the Guard Captain (19120, -171.2) [quest]; Drake Fang Talisman (19406, -199.2) [raid] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 283.2 | yes | Ironbark Staff (20220, +227.2) [rep]; Atiesh, Greatstaff of the Guardian (22630, +227.2) [quest]; High Warlord's War Staff (234549, +227.2) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 283.2 | yes | Grand Marshal's Handaxe (18827, +0.0) [vendor]; High Warlord's Cleaver (18828, +0.0) [vendor]; Grand Marshal's Dirk (18838, +0.0) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 961.6 | yes | The Purifier (22656, -232.8) [quest]; Huhuran's Stinger (21616, -277.4) [raid]; High Warlord's Recurve (234559, -287.1) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Stormrage's Talisman of Seething; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Hauberk; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legguards; feet: Cryptstalker Boots; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2412, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

