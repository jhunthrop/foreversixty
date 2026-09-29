# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (dwarf, 5420000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 75.1. Weights run: 1.5s. Verify run: 1.6s. 384 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.067 ± 0.030, crit=8.180 ± 0.347, hit=4.699 ± 0.333, melee_haste=5.581 ± 0.983

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.5 | yes | Flying Tiger Goggles (4368, -16.5) [crafted]; Shadow Goggles (4373, -16.5) [crafted]; Lucky Fishing Hat (19972, -16.5) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.4 | yes | Tarnished Locket (279870, -12.4) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.3 | yes | Double-Stitched Woolen Shoulders (4314, -10.3) [crafted]; Reinforced Woolen Shoulders (4315, -10.3) [crafted]; Forest Leather Mantle (4709, -10.3) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.4 | yes | Cape of the Brotherhood (5193, -2.1) [dungeon]; Sentry Cloak (2059, -4.1) [dungeon]; Hide of Lupos (3018, -4.1) [world] |
| chest | Tunic of Westfall (2041) | Quests [quest] | 22.7 | yes | Brawler's Leather Armor (252490, -8.3) [crafted]; Trapper's Leather Armor (252491, -8.3) [crafted]; Dark Leather Tunic (2317, -10.3) [crafted] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 10.3 | yes | Wolf Bracers (4794, -2.1) [vendor]; Bravo's Armbands (270015, -2.1) [quest]; Ratchet Wristwraps (274742, -4.1) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 114.5 | yes | Serpent Gloves (5970, -102.1) [dungeon]; Gloves of the Fang (10413, -102.1) [dungeon]; Forest Leather Gloves (3058, -106.3) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -7.7) [crafted]; Dusty Belt (279897, -7.7) [quest]; Guardsman Belt (3429, -9.7) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.6 | yes | Brawler's Leather Pants (252500, +0.0) [crafted]; Trapper's Leather Pants (252501, +0.0) [crafted]; Bluegill Breeches (3022, -2.1) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 16.5 | yes | Blackened Defias Boots (10402, -4.1) [dungeon]; Footpads of the Fang (10411, -4.1) [dungeon]; Dark Leather Boots (2315, -6.2) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.4 | yes | Lavishly Jeweled Ring (1156, -8.3) [dungeon]; The 1 Ring (8350, -10.3) [world]; Minor Channeling Ring (1449, -12.4) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 | yes | Lavishly Jeweled Ring (1156, -4.1) [dungeon]; The 1 Ring (8350, -6.2) [world]; Minor Channeling Ring (1449, -8.3) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 22.5 | yes | Impaling Harpoon (5200, -3.9) [dungeon]; Scythe Axe (5749, -8.1) [world]; Lupine Axe (1220, -10.1) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 178.3 | yes | Cracked Blacksmith Hammer (285279, -38.3) [crafted]; Lovingly Crafted Boomstick (4372, -42.2) [crafted]; Venomstrike (6469, -44.2) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 384, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 30 (dwarf, 5420001504000000-00000000000000000-000000000000000000)

Set DPS (verified): 100.0. Weights run: 1.6s. Verify run: 2.0s. 731 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.156 ± 0.049, crit=9.416 ± 0.378, hit=5.360 ± 0.397, melee_haste=not significant (1.327 ± 1.223)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 | yes | Tribal Worg Helm (6204, -4.3) [world]; Brawler's Leather Hood (252504, -4.3) [crafted]; Holy Shroud (2721, -21.6) [dungeon] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 17.2 | yes | Ghostshard Talisman (7731, -3.2) [dungeon]; Sentinel's Medallion (20444, -4.3) [rep]; Pendant of Myzrael (4614, -17.2) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 23.7 | yes | Mantle of Thieves (2264, -2.2) [dungeon]; Dark Leather Shoulders (4252, -8.6) [crafted]; Insignia Mantle (4721, -8.6) [dungeon] |
| back | Cloak of Night (4447) (or Fenrus' Hide (6340), Glowing Lizardscale Cloak (6449)) | Rohh the Silent [world] | 12.9 | yes | Fenrus' Hide (6340, +0.0) [dungeon]; Glowing Lizardscale Cloak (6449, +0.0) [dungeon]; Cape of the Brotherhood (5193, -2.2) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.2 | yes | Tunic of Westfall (2041, -6.5) [quest]; Green Leather Armor (4255, -12.9) [crafted]; Brawler's Leather Tunic (252508, -12.9) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410)) | Razormaw Matriarch [world] | 12.9 | yes | Insignia Bracers (6410, +0.0) [dungeon]; Madwolf Bracers (897, -2.2) [world]; Forest Leather Bracers (3202, -2.2) [dungeon] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 131.8 | yes | Pilferer's Gloves (7358, -114.6) [crafted]; Heavy Earthen Gloves (7359, -115.8) [crafted]; Wolfclaw Gloves (1978, -118.9) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.0) [rep]; Skulker's Leather Belt (252520, -4.6) [crafted]; Stalker's Leather Belt (252521, -4.6) [crafted] |
| legs | Dusky Leather Leggings (7373) | Leatherworking [crafted] | 28.0 | yes | Ferine Leggings (6690, -2.0) [dungeon]; Insignia Leggings (4054, -8.6) [dungeon]; Leggings of the Fang (10410, -8.6) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 17.2 | yes | Insignia Boots (4055, +0.0) [dungeon]; Highlander's Mail Greaves (20123, +0.0) [vendor]; Lancer Boots (6752, -2.2) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.4 | yes | Ring of Precision (1491, -6.5) [dungeon]; Protector's Band (19517, -6.5) [rep]; Signet of the Zhevra (285330, -6.5) [world] |
| finger2 | Monkey Ring (6748) | Quests [quest] | 15.1 | yes | Ring of Precision (1491, -2.2) [dungeon]; Protector's Band (19517, -2.2) [rep]; Signet of the Zhevra (285330, -2.2) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 | yes | Armor Piercer (6679, +6.9) [dungeon]; Bronze Dory (250603, +6.5) [crafted]; Kam's Walking Stick (2280, +4.8) [dungeon] |
| off_hand | Prison Shank (2941) (or Talon of Vultros (4454), Sentinel's Blade (212583), Scout's Blade (212587)) | The Stockade: Bruegal Ironknuckle [dungeon] | 12.9 | yes | Talon of Vultros (4454, +0.0) [world]; Sentinel's Blade (212583, +0.0) [vendor]; Scout's Blade (212587, +0.0) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 | yes | Silver Star (3463, -4.2) [quest]; Nightstalker Bow (6696, -28.6) [dungeon]; Satchel of Bronze Bombs (285276, -48.1) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Cloak of Night; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Highlander's Chain Girdle; legs: Dusky Leather Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Alliance Outrunner's Sword; off_hand: Prison Shank; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 731, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword; 7958 Bronze Battle Axe

### Band 40 (dwarf, 5420001505001251-00000000000000000-000000000000000000)

Set DPS (verified): 124.8. Weights run: 1.8s. Verify run: 2.1s. 1253 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.142 ± 0.050, crit=10.937 ± 0.446, hit=6.223 ± 0.768, melee_haste=not significant (4.753 ± 1.936)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 153.1 | yes | Nightscape Headband (8176, -127.4) [crafted]; Guard's Chain Helm (250499, -127.4) [crafted]; White Bandit Mask (10008, -129.6) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 23.6 | yes | Sentinel's Medallion (19541, -6.4) [rep]; Ghostshard Talisman (7731, -9.6) [dungeon]; Sentinel's Medallion (20444, -10.7) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 35.6 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Mantle of Thieves (2264, -14.1) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | Gnomeregan: Caverndeep Burrower [dungeon] | 17.1 | yes | Parachute Cloak (10518, +0.0) [crafted]; Yeti Fur Cloak (2805, -4.3) [quest]; Darktide Cape (4114, -4.3) [quest] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 32.1 | yes | Tough Scorpid Breastplate (8203, +0.0) [crafted]; Dusky Leather Armor (7374, -2.1) [crafted]; Hawkeye's Tunic (14592, -6.4) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -2.9) [dungeon]; Dusky Bracers (7378, -2.9) [crafted]; Tough Scorpid Bracers (8205, -5.0) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 173.1 | yes | Dragonscale Gauntlets (8347, -7.1) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 161.1 | yes | Highlander's Leather Girdle (20116, -131.1) [rep]; Highlander's Chain Girdle (20090, -137.1) [rep]; Highlander's Leather Girdle (20117, -137.1) [rep] |
| legs | Basilisk Hide Pants (1718) | Gnomeregan: Leprous Machinesmith [dungeon] | 45.0 | yes | Triprunner Dungarees (9624, -6.4) [quest]; Hawkeye's Breeches (14595, -15.0) [world]; Dusky Leather Leggings (7373, -17.1) [crafted] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 27.8 | yes | Imperial Leather Boots (6431, -4.3) [dungeon]; Dusky Boots (7390, -4.3) [crafted]; Skulker's Leather Shoes (252531, -4.3) [crafted] |
| finger1 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 21.4 | yes | Protector's Band (19515, -4.3) [rep]; Disengagement Ring (276202, -4.3) [vendor]; Monkey Ring (6748, -6.4) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.3 | yes | Protector's Band (19515, -2.1) [rep]; Disengagement Ring (276202, -2.1) [vendor]; Monkey Ring (6748, -4.3) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Blazing Emblem (2802, +0.0) [dungeon]; Cold Basilisk Eye (5079, +0.0) [world] |
| main_hand | Frost Tiger Blade (3854) (or Illusionary Rod (7713)) | Blacksmithing [crafted] | 153.1 | yes | Illusionary Rod (7713, +0.0) [dungeon]; Steel Spear (250605, -115.3) [crafted]; Loksey's Training Stick (7710, -125.1) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Master Hunter's Bow (17686) | Quests [quest] | 284.2 | yes | Sniper Rifle (3430, -1.8) [dungeon]; Mithril Heavy-bore Rifle (10510, -1.8) [crafted]; Master Hunter's Rifle (17687, -3.7) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Frost Tiger Blade; ranged: Master Hunter's Bow

No-known-source sample (15 of 1253, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 5420001505001251-35200000000000000-000000000000000000)

Set DPS (verified): 146.8. Weights run: 1.8s. Verify run: 2.3s. 1664 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.214 ± 0.065, crit=10.645 ± 0.461, hit=7.032 ± 0.782, melee_haste=not significant (3.939 ± 2.022)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) (or Eye of Theradras (17715)) | Scarlet Monastery: Herod [dungeon] | 149.0 | yes | Eye of Theradras (17715, +0.0) [dungeon]; Helm of Fire (8348, -111.4) [crafted]; Lordrec Helmet (10741, -113.6) [quest] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 26.6 | yes | Sentinel's Medallion (19540, -2.2) [rep]; Sentinel's Medallion (19541, -8.9) [rep]; Ghostshard Talisman (7731, -12.6) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.4 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Penance Spaulders (11963, -12.0) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 24.4 | yes | Nightscape Cloak (8195, -2.2) [crafted]; Pridelord Cape (14673, -4.4) [dungeon]; Imperial Cloak (6432, -6.6) [dungeon] |
| chest | Blazewind Breastplate (11193) | Quests [quest] | 50.9 | yes | Wildthorn Mail (12624, +19.4) [crafted]; Warbear Harness (15064, -11.1) [crafted]; Charred Leather Tunic (19127, -11.1) [quest] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 33.2 | yes | Bracers of the Stone Princess (17714, -5.2) [dungeon]; Bloodlust Bracelets (14807, -8.9) [dungeon]; Wicked Leather Bracers (15084, -8.9) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 169.0 | yes | Dragonscale Gauntlets (8347, -6.7) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 169.0 | yes | Highlander's Leather Girdle (20115, +0.0) [rep]; Highlander's Chain Girdle (20089, -12.0) [rep]; Highlander's Cloth Girdle (20097, -20.0) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 298.1 | yes | Basilisk Hide Pants (1718, -251.6) [dungeon]; Keeper's Woolies (14668, -256.0) [dungeon]; Oilskin Leggings (9414, -258.2) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 70.3 | yes | Albino Crocscale Boots (17728, -26.0) [dungeon]; Fleetfoot Greaves (11627, -28.2) [dungeon]; Elven Chain Boots (13125, -30.5) [world] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 90.3 | yes | Ring of the Underwood (2951, -68.2) [dungeon]; Ironspine's Eye (7686, -70.4) [dungeon]; Protector's Band (19516, -70.4) [rep] |
| finger2 | Masons Fraternity Ring (9533) | Quests [quest] | 31.0 | yes | Ring of the Underwood (2951, -8.9) [dungeon]; Ironspine's Eye (7686, -11.1) [dungeon]; Protector's Band (19516, -11.1) [rep] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.0) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Dawn's Edge (12774) (or Inventor's Focal Sword (17719)) | Blacksmithing [crafted] | 149.0 | yes | Frost Tiger Blade (3854, +0.0) [crafted]; Illusionary Rod (7713, +0.0) [dungeon]; Kindling Stave (11750, +0.0) [dungeon] |
| off_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 149.0 | yes | Thermotastic Egg Timer (9644, -142.4) [quest]; Grayson's Torch (1172, -149.0) [quest]; Rod of Molten Fire (2565, -149.0) [dungeon] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 448.7 | yes | Precisely Calibrated Boomstick (2100, -30.3) [dungeon]; Dark Iron Rifle (16004, -52.1) [crafted]; Houndmaster's Bow (11628, -55.6) [dungeon] |

**New at 50:** neck: Sentinel's Medallion; back: Serpentskin Cloak; chest: Blazewind Breastplate; wrist: Deepfury Bracers; waist: Highlander's Chain Girdle; legs: Stormshroud Pants; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Dawn's Edge; off_hand: Inventor's Focal Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1664, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 5420001505001251-35510000000000000-510000000000000000)

Set DPS (verified): 245.5. Weights run: 1.8s. Verify run: 2.4s. 2330 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.233 ± 0.071, crit=15.231 ± 0.576, hit=not significant (0.000 ± 0.000), melee_haste=12.796 ± 2.185

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Quests [quest] | 495.7 | yes | Champion's Chain Helm (23251, -29.0) [vendor]; Lieutenant Commander's Chain Helm (23306, -29.0) [vendor]; Lieutenant Commander's Chain Helm (227066, -29.0) [pvp] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 452.5 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Barbed Choker (21664, -195.2) [raid]; Medallion of the Dawn (22659, -215.2) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Quests [quest] | 278.0 | yes | Champion's Chain Shoulders (23252, -24.6) [vendor]; Lieutenant Commander's Chain Shoulders (23307, -24.6) [vendor]; Lieutenant Commander's Chain Shoulders (227068, -24.6) [pvp] |
| back | Cloak of the Fallen God (21710) | Quests [quest] | 58.1 | yes | Chromatic Cloak (18509, +155.2) [crafted]; Cape of the Black Baron (13340, -4.6) [dungeon]; Cloak of the Honor Guard (20073, -12.9) [rep] |
| chest | Legionnaire's Chain Hauberk (22874) (or Knight-Captain's Chain Hauberk (23292), Legionnaire's Chain Hauberk (227071)) | Lady Palanseer [vendor] | 462.2 | yes | Knight-Captain's Chain Hauberk (23292, +0.0) [vendor]; Legionnaire's Chain Hauberk (227071, +0.0) [pvp]; Bloodsoul Breastplate (19690, -15.6) [crafted] |
| wrist | Cryptstalker Wristguards (22443) | Quests [quest] | 58.1 | yes | Marshal's Chain Bracers (16461, -13.4) [pvp]; General's Chain Wristguards (16570, -13.4) [pvp]; Forest Stalker's Bracers (19587, -15.6) [rep] |
| hands | Cryptstalker Handguards (22441) | Quests [quest] | 266.8 | yes | Marshal's Chain Grips (16463, -6.7) [vendor]; General's Chain Gloves (16571, -6.7) [vendor]; Marshal's Chain Grips (231560, -9.3) [pvp] |
| waist | Cryptstalker Girdle (22442) | Quests [quest] | 264.6 | yes | Belt of Never-ending Agony (21586, +12.6) [raid]; Highlander's Chain Girdle (20043, -17.4) [rep]; Highlander's Leather Girdle (20045, -17.4) [rep] |
| legs | Legionnaire's Chain Legguards (22875) (or Knight-Captain's Chain Legguards (23293), Knight-Captain's Chain Legguards (227072), Legionnaire's Chain Legguards (227073)) | Lady Palanseer [vendor] | 462.2 | yes | Knight-Captain's Chain Legguards (23293, +0.0) [vendor]; Knight-Captain's Chain Legguards (227072, +0.0) [pvp]; Legionnaire's Chain Legguards (227073, +0.0) [pvp] |
| feet | Cryptstalker Boots (22440) | Quests [quest] | 73.7 | yes | Striker's Footguards (21365, -4.5) [quest]; Marshal's Chain Boots (16462, -15.6) [vendor]; General's Chain Sabatons (16569, -15.6) [vendor] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 265.2 | yes | Quick Strike Ring (18821, -22.0) [raid]; Don Julio's Band (19325, -36.0) [rep]; Band of the Penitent (13217, -52.0) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj [raid] | 253.2 | yes | Quick Strike Ring (18821, -10.0) [raid]; Don Julio's Band (19325, -24.0) [rep]; Band of the Penitent (13217, -40.0) [quest] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +362.5) [raid]; Drake Fang Talisman (19406, -8.0) [raid]; Thunderbrew's Boot Flask (744, -64.0) [quest] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 213.2 | yes | Eye of Diminution (23001, +213.2) [raid]; Drake Fang Talisman (19406, -157.2) [raid]; Thunderbrew's Boot Flask (744, -213.2) [quest] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 241.2 | yes | Ironbark Staff (20069, +185.2) [rep]; Atiesh, Greatstaff of the Guardian (22630, +185.2) [quest]; High Warlord's War Staff (234549, +185.2) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 241.2 | yes | Grand Marshal's Handaxe (18827, +0.0) [vendor]; High Warlord's Cleaver (18828, +0.0) [vendor]; Grand Marshal's Dirk (18838, +0.0) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 919.6 | yes | Huhuran's Stinger (21616, -231.2) [raid]; The Purifier (22656, -232.8) [quest]; High Warlord's Recurve (234559, -245.1) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Stormrage's Talisman of Seething; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Hauberk; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legguards; feet: Cryptstalker Boots; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2330, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 5420000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 74.6. Weights run: 1.5s. Verify run: 1.6s. 379 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.067 ± 0.030, crit=8.180 ± 0.347, hit=4.699 ± 0.333, melee_haste=5.581 ± 0.983

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.5 | yes | Flying Tiger Goggles (4368, -16.5) [crafted]; Shadow Goggles (4373, -16.5) [crafted]; Lucky Fishing Hat (19972, -16.5) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.4 | yes | Tarnished Locket (279870, -12.4) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.3 | yes | Double-Stitched Woolen Shoulders (4314, -10.3) [crafted]; Reinforced Woolen Shoulders (4315, -10.3) [crafted]; Forest Leather Mantle (4709, -10.3) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.4 | yes | Cape of the Brotherhood (5193, -2.1) [dungeon]; Sentry Cloak (2059, -4.1) [dungeon]; Hide of Lupos (3018, -4.1) [world] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 14.5 | yes | Trapper's Leather Armor (252491, +0.0) [crafted]; Dark Leather Tunic (2317, -2.1) [crafted]; Heckler's Hide (286536, -4.1) [world] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 10.3 | yes | Wolf Bracers (4794, -2.1) [vendor]; Bravo's Armbands (270015, -2.1) [quest]; Ratchet Wristwraps (274742, -4.1) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 114.5 | yes | Serpent Gloves (5970, -102.1) [dungeon]; Gloves of the Fang (10413, -102.1) [dungeon]; Forest Leather Gloves (3058, -106.3) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -7.7) [crafted]; Dusty Belt (279897, -7.7) [quest]; Guardsman Belt (3429, -9.7) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.6 | yes | Brawler's Leather Pants (252500, +0.0) [crafted]; Trapper's Leather Pants (252501, +0.0) [crafted]; Bluegill Breeches (3022, -2.1) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 16.5 | yes | Blackened Defias Boots (10402, -4.1) [dungeon]; Footpads of the Fang (10411, -4.1) [dungeon]; Dark Leather Boots (2315, -6.2) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.4 | yes | Bounty Hunter's Ring (5351, -6.2) [quest]; Lavishly Jeweled Ring (1156, -8.3) [dungeon]; The 1 Ring (8350, -10.3) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 | yes | Bounty Hunter's Ring (5351, -2.1) [quest]; Lavishly Jeweled Ring (1156, -4.1) [dungeon]; The 1 Ring (8350, -6.2) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 22.5 | yes | Impaling Harpoon (5200, -3.9) [dungeon]; Scythe Axe (5749, -8.1) [world]; Crescent Staff (6505, -8.1) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 178.3 | yes | Cracked Blacksmith Hammer (285279, -38.3) [crafted]; Lovingly Crafted Boomstick (4372, -42.2) [crafted]; Venomstrike (6469, -44.2) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 379, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 30 (troll, 5420001504000000-00000000000000000-000000000000000000)

Set DPS (verified): 102.5. Weights run: 1.6s. Verify run: 2.1s. 727 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.156 ± 0.049, crit=9.416 ± 0.378, hit=5.360 ± 0.397, melee_haste=not significant (1.327 ± 1.223)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tribal Worg Helm (6204) | Fenros [world] | 17.2 | yes | Brawler's Leather Helm (252512, +4.3) [crafted]; Brawler's Leather Hood (252504, +0.0) [crafted]; Holy Shroud (2721, -17.2) [dungeon] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 17.2 | yes | Ghostshard Talisman (7731, -3.2) [dungeon]; Scout's Medallion (20442, -4.3) [rep]; Pendant of Myzrael (4614, -17.2) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 23.7 | yes | Mantle of Thieves (2264, -2.2) [dungeon]; Dark Leather Shoulders (4252, -8.6) [crafted]; Insignia Mantle (4721, -8.6) [dungeon] |
| back | Cloak of Night (4447) (or Fenrus' Hide (6340), Glowing Lizardscale Cloak (6449), Swiftrunner Cape (6745)) | Rohh the Silent [world] | 12.9 | yes | Fenrus' Hide (6340, +0.0) [dungeon]; Glowing Lizardscale Cloak (6449, +0.0) [dungeon]; Swiftrunner Cape (6745, +0.0) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.2 | yes | Panther Armor (6670, -10.8) [quest]; Green Leather Armor (4255, -12.9) [crafted]; Brawler's Leather Tunic (252508, -12.9) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410)) | Razormaw Matriarch [world] | 12.9 | yes | Insignia Bracers (6410, +0.0) [dungeon]; Madwolf Bracers (897, -2.2) [world]; Forest Leather Bracers (3202, -2.2) [dungeon] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 131.8 | yes | Pilferer's Gloves (7358, -114.6) [crafted]; Heavy Earthen Gloves (7359, -115.8) [crafted]; Braced Handguards (6784, -116.7) [quest] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.0) [rep]; Deftkin Belt (16659, -3.4) [quest]; Skulker's Leather Belt (252520, -4.6) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, +2.0) [crafted]; Insignia Leggings (4054, -6.6) [dungeon]; Leggings of the Fang (10410, -6.6) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 17.2 | yes | Insignia Boots (4055, +0.0) [dungeon]; Warsong Boots (16977, +0.0) [quest]; Highlander's Mail Greaves (20123, +0.0) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.4 | yes | Ring of Precision (1491, -6.5) [dungeon]; Legionnaire's Band (19513, -6.5) [rep]; Signet of the Zhevra (285330, -6.5) [world] |
| finger2 | Monkey Ring (6748) | Quests [quest] | 15.1 | yes | Ring of Precision (1491, -2.2) [dungeon]; Legionnaire's Band (19513, -2.2) [rep]; Signet of the Zhevra (285330, -2.2) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.0) [rep]; Rune of Duty (21568, +0.0) [rep]; Relentless Raider's Seal (272062, +0.0) [vendor] |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 | yes | Armor Piercer (6679, +6.9) [dungeon]; Bronze Dory (250603, +6.5) [crafted]; Kam's Walking Stick (2280, +4.8) [dungeon] |
| off_hand | Prison Shank (2941) (or Talon of Vultros (4454), Sentinel's Blade (212583), Scout's Blade (212587)) | The Stockade: Bruegal Ironknuckle [dungeon] | 12.9 | yes | Talon of Vultros (4454, +0.0) [world]; Sentinel's Blade (212583, +0.0) [vendor]; Scout's Blade (212587, +0.0) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 | yes | Silver Star (3463, -4.2) [quest]; Nightstalker Bow (6696, -28.6) [dungeon]; Satchel of Bronze Bombs (285276, -48.1) [crafted] |

**New at 30:** head: Tribal Worg Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Cloak of Night; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Alliance Outrunner's Sword; off_hand: Prison Shank; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 727, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword; 7958 Bronze Battle Axe

### Band 40 (troll, 5420001505001251-00000000000000000-000000000000000000)

Set DPS (verified): 127.0. Weights run: 1.8s. Verify run: 2.2s. 1248 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.142 ± 0.050, crit=10.937 ± 0.446, hit=6.223 ± 0.768, melee_haste=not significant (4.753 ± 1.936)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 153.1 | yes | Nightscape Headband (8176, -127.4) [crafted]; Guard's Chain Helm (250499, -127.4) [crafted]; White Bandit Mask (10008, -129.6) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 23.6 | yes | Scout's Medallion (19537, -6.4) [rep]; Ghostshard Talisman (7731, -9.6) [dungeon]; Scout's Medallion (20442, -10.7) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 35.6 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Mantle of Thieves (2264, -14.1) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | Gnomeregan: Caverndeep Burrower [dungeon] | 17.1 | yes | Parachute Cloak (10518, +0.0) [crafted]; Darktide Cape (4114, -4.3) [quest]; Cloak of Night (4447, -4.3) [world] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 32.1 | yes | Tough Scorpid Breastplate (8203, +0.0) [crafted]; Dusky Leather Armor (7374, -2.1) [crafted]; Hawkeye's Tunic (14592, -6.4) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -2.9) [dungeon]; Dusky Bracers (7378, -2.9) [crafted]; Tough Scorpid Bracers (8205, -5.0) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 173.1 | yes | Dragonscale Gauntlets (8347, -7.1) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 161.1 | yes | Defiler's Leather Girdle (20192, -131.1) [rep]; Defiler's Chain Girdle (20152, -137.1) [rep]; Defiler's Leather Girdle (20191, -137.1) [rep] |
| legs | Basilisk Hide Pants (1718) | Gnomeregan: Leprous Machinesmith [dungeon] | 45.0 | yes | Triprunner Dungarees (9624, -6.4) [quest]; Hawkeye's Breeches (14595, -15.0) [world]; Dusky Leather Leggings (7373, -17.1) [crafted] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 27.8 | yes | Imperial Leather Boots (6431, -4.3) [dungeon]; Dusky Boots (7390, -4.3) [crafted]; Skulker's Leather Shoes (252531, -4.3) [crafted] |
| finger1 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 21.4 | yes | Legionnaire's Band (19512, -4.3) [rep]; Disengagement Ring (276202, -4.3) [vendor]; Monkey Ring (6748, -6.4) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.3 | yes | Legionnaire's Band (19512, -2.1) [rep]; Disengagement Ring (276202, -2.1) [vendor]; Monkey Ring (6748, -4.3) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Ankh of Life (1713, +0.0) [dungeon]; Blazing Emblem (2802, +0.0) [dungeon] |
| trinket2 | Rune of Duty (21567) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.0) [vendor]; Ankh of Life (1713, +0.0) [dungeon]; Blazing Emblem (2802, +0.0) [dungeon] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 153.1 | yes | Frost Tiger Blade (3854, +0.0) [crafted]; Steel Spear (250605, -115.3) [crafted]; Loksey's Training Stick (7710, -125.1) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) | Gnomeregan: Dark Iron Agent [dungeon] | 282.4 | yes | Master Hunter's Bow (17686, +1.8) [quest]; Mithril Heavy-bore Rifle (10510, +0.0) [crafted]; Master Hunter's Rifle (17687, -1.9) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Illusionary Rod; ranged: Sniper Rifle

No-known-source sample (15 of 1248, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 5420001505001251-35200000000000000-000000000000000000)

Set DPS (verified): 148.0. Weights run: 1.8s. Verify run: 2.3s. 1660 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.214 ± 0.065, crit=10.645 ± 0.461, hit=7.032 ± 0.782, melee_haste=not significant (3.939 ± 2.022)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 149.0 | yes | Raging Berserker's Helm (7719, +0.0) [dungeon]; Helm of Fire (8348, -111.4) [crafted]; Sprightring Helm (17776, -115.8) [quest] |
| neck | Scout's Medallion (19535) | Warsong Outriders [rep] | 26.6 | yes | Scout's Medallion (19536, -2.2) [rep]; Woven Ivy Necklace (19159, -6.6) [quest]; Scout's Medallion (19537, -8.9) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.4 | yes | Forest Tracker Epaulets (2278, -12.0) [dungeon]; Nightscape Shoulders (8192, -12.0) [crafted]; Penance Spaulders (11963, -12.0) [quest] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 24.4 | yes | Nightscape Cloak (8195, -2.2) [crafted]; Pridelord Cape (14673, -4.4) [dungeon]; Imperial Cloak (6432, -6.6) [dungeon] |
| chest | Blazewind Breastplate (11193) | Quests [quest] | 50.9 | yes | Wildthorn Mail (12624, +19.4) [crafted]; Warbear Harness (15064, -11.1) [crafted]; Charred Leather Tunic (19127, -11.1) [quest] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 33.2 | yes | Bracers of the Stone Princess (17714, -5.2) [dungeon]; Bloodlust Bracelets (14807, -8.9) [dungeon]; Wicked Leather Bracers (15084, -8.9) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 169.0 | yes | Dragonscale Gauntlets (8347, -6.7) [crafted]; Fletcher's Gloves (7348, -20.0) [crafted]; Shadowskin Gloves (18238, -20.0) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 169.0 | yes | Defiler's Leather Girdle (20193, +0.0) [rep]; Defiler's Chain Girdle (20153, -12.0) [rep]; Highlander's Mail Girdle (20118, -20.0) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 298.1 | yes | Basilisk Hide Pants (1718, -251.6) [dungeon]; Keeper's Woolies (14668, -256.0) [dungeon]; Oilskin Leggings (9414, -258.2) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 70.3 | yes | Albino Crocscale Boots (17728, -26.0) [dungeon]; Fleetfoot Greaves (11627, -28.2) [dungeon]; Elven Chain Boots (13125, -30.5) [world] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 90.3 | yes | White Bone Band (11862, -66.3) [quest]; Ring of the Underwood (2951, -68.2) [dungeon]; Ironspine's Eye (7686, -70.4) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Quests [quest] | 31.0 | yes | White Bone Band (11862, -7.0) [quest]; Ring of the Underwood (2951, -8.9) [dungeon]; Ironspine's Eye (7686, -11.1) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of the Guard Captain (19120, +133.2) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Rune of the Guard Captain (19120, +133.2) [quest]; Tidal Charm (1404, +0.0) [vendor]; Guardian Talisman (1490, +0.0) [quest] |
| main_hand | Dawn's Edge (12774) (or Inventor's Focal Sword (17719)) | Blacksmithing [crafted] | 149.0 | yes | Frost Tiger Blade (3854, +0.0) [crafted]; Illusionary Rod (7713, +0.0) [dungeon]; Kindling Stave (11750, +0.0) [dungeon] |
| off_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 149.0 | yes | White Bone Shredder (11863, -133.5) [quest]; Thermotastic Egg Timer (9644, -142.4) [quest]; Grayson's Torch (1172, -149.0) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 448.7 | yes | Precisely Calibrated Boomstick (2100, -30.3) [dungeon]; Dark Iron Rifle (16004, -52.1) [crafted]; Houndmaster's Bow (11628, -55.6) [dungeon] |

**New at 50:** head: Eye of Theradras; neck: Scout's Medallion; back: Serpentskin Cloak; chest: Blazewind Breastplate; wrist: Deepfury Bracers; waist: Defiler's Chain Girdle; legs: Stormshroud Pants; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Dawn's Edge; off_hand: Inventor's Focal Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1660, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 5420001505001251-35510000000000000-510000000000000000)

Set DPS (verified): 247.8. Weights run: 1.8s. Verify run: 2.3s. 2324 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.233 ± 0.071, crit=15.231 ± 0.576, hit=not significant (0.000 ± 0.000), melee_haste=12.796 ± 2.185

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Quests [quest] | 495.7 | yes | Champion's Chain Helm (23251, -29.0) [vendor]; Lieutenant Commander's Chain Helm (23306, -29.0) [vendor]; Lieutenant Commander's Chain Helm (227066, -29.0) [pvp] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas [raid] | 452.5 | yes | Gem of Trapped Innocents (23057, -26.0) [raid]; Barbed Choker (21664, -195.2) [raid]; Medallion of the Dawn (22659, -215.2) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Quests [quest] | 278.0 | yes | Champion's Chain Shoulders (23252, -24.6) [vendor]; Lieutenant Commander's Chain Shoulders (23307, -24.6) [vendor]; Lieutenant Commander's Chain Shoulders (227068, -24.6) [pvp] |
| back | Cloak of the Fallen God (21710) | Quests [quest] | 58.1 | yes | Chromatic Cloak (18509, +155.2) [crafted]; Cape of the Black Baron (13340, -4.6) [dungeon]; Deathguard's Cloak (20068, -12.9) [rep] |
| chest | Legionnaire's Chain Hauberk (22874) (or Knight-Captain's Chain Hauberk (23292), Legionnaire's Chain Hauberk (227071)) | Lady Palanseer [vendor] | 462.2 | yes | Knight-Captain's Chain Hauberk (23292, +0.0) [vendor]; Legionnaire's Chain Hauberk (227071, +0.0) [pvp]; Bloodsoul Breastplate (19690, -15.6) [crafted] |
| wrist | Cryptstalker Wristguards (22443) | Quests [quest] | 58.1 | yes | Marshal's Chain Bracers (16461, -13.4) [pvp]; General's Chain Wristguards (16570, -13.4) [pvp]; Forest Stalker's Bracers (19587, -15.6) [rep] |
| hands | Cryptstalker Handguards (22441) | Quests [quest] | 266.8 | yes | Marshal's Chain Grips (16463, -6.7) [vendor]; General's Chain Gloves (16571, -6.7) [vendor]; Marshal's Chain Grips (231560, -9.3) [pvp] |
| waist | Cryptstalker Girdle (22442) | Quests [quest] | 264.6 | yes | Belt of Never-ending Agony (21586, +12.6) [raid]; Defiler's Chain Girdle (20150, -17.4) [rep]; Defiler's Leather Girdle (20190, -17.4) [rep] |
| legs | Legionnaire's Chain Legguards (22875) (or Knight-Captain's Chain Legguards (23293), Knight-Captain's Chain Legguards (227072), Legionnaire's Chain Legguards (227073)) | Lady Palanseer [vendor] | 462.2 | yes | Knight-Captain's Chain Legguards (23293, +0.0) [vendor]; Knight-Captain's Chain Legguards (227072, +0.0) [pvp]; Legionnaire's Chain Legguards (227073, +0.0) [pvp] |
| feet | Cryptstalker Boots (22440) | Quests [quest] | 73.7 | yes | Striker's Footguards (21365, -4.5) [quest]; Marshal's Chain Boots (16462, -15.6) [vendor]; General's Chain Sabatons (16569, -15.6) [vendor] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas [raid] | 265.2 | yes | Quick Strike Ring (18821, -22.0) [raid]; Don Julio's Band (19325, -36.0) [rep]; Band of the Penitent (13217, -52.0) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj [raid] | 253.2 | yes | Quick Strike Ring (18821, -10.0) [raid]; Don Julio's Band (19325, -24.0) [rep]; Band of the Penitent (13217, -40.0) [quest] |
| trinket1 | Slayer's Crest (23041) | Naxxramas [raid] | 64.0 | yes | Eye of Diminution (23001, +362.5) [raid]; Rune of the Guard Captain (19120, +20.0) [quest]; Drake Fang Talisman (19406, -8.0) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas [raid] | 213.2 | yes | Eye of Diminution (23001, +213.2) [raid]; Rune of the Guard Captain (19120, -129.2) [quest]; Drake Fang Talisman (19406, -157.2) [raid] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 241.2 | yes | Ironbark Staff (20220, +185.2) [rep]; Atiesh, Greatstaff of the Guardian (22630, +185.2) [quest]; High Warlord's War Staff (234549, +185.2) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 241.2 | yes | Grand Marshal's Handaxe (18827, +0.0) [vendor]; High Warlord's Cleaver (18828, +0.0) [vendor]; Grand Marshal's Dirk (18838, +0.0) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj [raid] | 919.6 | yes | Huhuran's Stinger (21616, -231.2) [raid]; The Purifier (22656, -232.8) [quest]; High Warlord's Recurve (234559, -245.1) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Stormrage's Talisman of Seething; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Hauberk; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legguards; feet: Cryptstalker Boots; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2324, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

