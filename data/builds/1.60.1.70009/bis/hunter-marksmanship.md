# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (dwarf, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.1. Weights run: 1.1s. Verify run: 1.2s. 385 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.067 ± 0.030, crit=8.180 ± 0.347, hit=4.699 ± 0.333, melee_haste=5.581 ± 0.983

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.5 | yes | Shadow Goggles (4373, -1.00 DPS) [crafted]; Lucky Fishing Hat (19972, -1.00 DPS) [quest]; Flying Tiger Goggles (4368, -1.38 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.4 | yes | Tarnished Locket (279870, -1.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.3 | yes | Double-Stitched Woolen Shoulders (4314, -0.61 DPS, sim-verified) [crafted]; Reinforced Woolen Shoulders (4315, -0.62 DPS) [crafted]; Forest Leather Mantle (4709, -0.62 DPS) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.4 | yes | Cape of the Brotherhood (5193, +0.02 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.25 DPS) [world_drop]; Hide of Lupos (3018, -0.25 DPS) [world] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 22.7 | yes | Trapper's Leather Armor (252491, -0.50 DPS) [crafted]; Dark Leather Tunic (2317, -0.62 DPS) [crafted]; Brawler's Leather Armor (252490, -0.64 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.3 | yes | Bravo's Armbands (270015, -0.12 DPS) [quest]; Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.25 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 114.5 | yes | Serpent Gloves (5970, +0.41 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -6.16 DPS) [dungeon]; Forest Leather Gloves (3058, -6.41 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.46 DPS) [quest]; Deviate Scale Belt (6468, -0.46 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.59 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.6 | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.12 DPS) [world]; Brawler's Leather Pants (252500, -0.35 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.5 | yes | Footpads of the Fang (10411, -0.25 DPS) [dungeon]; Blackened Defias Boots (10402, -0.26 DPS, sim-verified) [dungeon]; Dark Leather Boots (2315, -0.37 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.4 | yes | Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon]; The 1 Ring (8350, -0.62 DPS) [world]; Minor Channeling Ring (1449, -0.75 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 | yes | Lavishly Jeweled Ring (1156, -0.08 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.37 DPS) [world]; Minor Channeling Ring (1449, -0.50 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 22.5 | yes | Impaling Harpoon (5200, +0.03 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.49 DPS) [world]; Lupine Axe (1220, -0.61 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.3 | yes | Lovingly Crafted Boomstick (4372, -2.55 DPS) [crafted]; Venomstrike (6469, -2.67 DPS) [dungeon]; Cracked Blacksmith Hammer (285279, -3.58 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 385, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 30 (dwarf, 0000000000000000-35305500000000000-000000000000000000)

Set DPS (verified): 90.1. Weights run: 1.2s. Verify run: 1.5s. 734 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.151 ± 0.048, crit=9.471 ± 0.389, hit=5.605 ± 0.389, melee_haste=not significant (1.703 ± 1.175)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.5 | yes | Tribal Worg Helm (6204, +0.48 DPS, sim-verified) [world]; Brawler's Leather Hood (252504, -0.26 DPS) [crafted]; Holy Shroud (2721, -1.28 DPS) [world_drop] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 17.2 | yes | Ghostshard Talisman (7731, -0.23 DPS, sim-verified) [dungeon]; Sentinel's Medallion (20444, -0.26 DPS) [rep]; Pendant of Myzrael (4614, -1.02 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.7 | yes | Dark Leather Shoulders (4252, -0.51 DPS) [crafted]; Insignia Mantle (4721, -0.51 DPS) [world_drop]; Mantle of Thieves (2264, -1.05 DPS, sim-verified) [dungeon] |
| back | Cloak of Night (4447) (or Fenrus' Hide (6340), Glowing Lizardscale Cloak (6449)) | Rohh the Silent [world] | 12.9 | yes | Fenrus' Hide (6340, +0.00 DPS, sim-verified) [dungeon]; Glowing Lizardscale Cloak (6449, +0.00 DPS) [dungeon]; Cape of the Brotherhood (5193, -0.13 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.1 | yes | Tunic of Westfall (2041, -0.41 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.77 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.77 DPS) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410)) | Razormaw Matriarch [world] | 12.9 | yes | Insignia Bracers (6410, +0.00 DPS, sim-verified) [world_drop]; Madwolf Bracers (897, -0.13 DPS) [world]; Forest Leather Bracers (3202, -0.13 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 132.6 | yes | Pilferer's Gloves (7358, +0.53 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -6.94 DPS) [crafted]; Wolfclaw Gloves (1978, -7.12 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted]; Stalker's Leather Belt (252521, -0.28 DPS) [crafted] |
| legs | Dusky Leather Leggings (7373) | Leatherworking [crafted] | 28.0 | yes | Ferine Leggings (6690, +0.74 DPS, sim-verified) [dungeon]; Insignia Leggings (4054, -0.51 DPS) [world_drop]; Leggings of the Fang (10410, -0.51 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 17.2 | yes | Insignia Boots (4055, +0.00 DPS, sim-verified) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.13 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.4 | yes | Ring of Precision (1491, -0.38 DPS) [dungeon]; Protector's Band (19517, -0.38 DPS) [rep]; Signet of the Zhevra (285330, -0.38 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.1 | yes | Protector's Band (19517, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Ring of Precision (1491, -0.32 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Prison Shank (2941) (or Talon of Vultros (4454), Sentinel's Blade (212583), Scout's Blade (212587)) | The Stockade: Bruegal Ironknuckle [dungeon] | 12.9 | yes | Sentinel's Blade (212583, +0.00 DPS) [vendor]; Scout's Blade (212587, +0.00 DPS) [vendor]; Talon of Vultros (4454, -0.91 DPS, sim-verified) [world] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 | yes | Silver Star (3463, -1.34 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -1.70 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -2.86 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Cloak of Night; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Highlander's Chain Girdle; legs: Dusky Leather Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Alliance Outrunner's Sword; off_hand: Prison Shank; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 734, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword; 7958 Bronze Battle Axe

### Band 40 (dwarf, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 102.1. Weights run: 1.2s. Verify run: 1.5s. 1260 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.199 ± 0.068, crit=13.685 ± 0.578, hit=6.174 ± 0.454, melee_haste=7.173 ± 1.232

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 191.6 | yes | Nightscape Headband (8176, +0.87 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -9.77 DPS) [crafted]; White Bandit Mask (10008, -9.90 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 24.2 | yes | Sentinel's Medallion (19541, -0.42 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.60 DPS) [dungeon]; Sentinel's Medallion (20444, -0.65 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.2 | yes | Nightscape Shoulders (8192, -0.71 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.77 DPS, sim-verified) [world_drop]; Mantle of Thieves (2264, -0.84 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | World drop [world_drop] | 17.6 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Yeti Fur Cloak (2805, -0.26 DPS) [quest]; Darktide Cape (4114, -0.26 DPS) [quest] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 33.0 | yes | Tough Scorpid Breastplate (8203, +0.81 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.13 DPS) [crafted]; Hawkeye's Tunic (14592, -0.39 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.13 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.14 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.27 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 211.6 | yes | Dragonscale Gauntlets (8347, -0.43 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.18 DPS) [crafted]; Shadowskin Gloves (18238, -1.18 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 199.6 | yes | Highlander's Leather Girdle (20116, +0.55 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -10.38 DPS) [rep]; Highlander's Leather Girdle (20117, -10.38 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.2 | yes | Triprunner Dungarees (9624, -0.58 DPS, sim-verified) [quest]; Hawkeye's Breeches (14595, -0.91 DPS) [world]; Dusky Leather Leggings (7373, -1.04 DPS) [crafted] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.6 | yes | Dusky Boots (7390, -0.26 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.26 DPS) [crafted]; Imperial Leather Boots (6431, -0.28 DPS, sim-verified) [world_drop] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 22.0 | yes | Protector's Band (19515, -0.26 DPS) [rep]; Disengagement Ring (276202, -0.26 DPS) [vendor]; Monkey Ring (6748, -0.39 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.8 | yes | Disengagement Ring (276202, -0.13 DPS) [vendor]; Protector's Band (19515, -0.14 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.26 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Frost Tiger Blade (3854) (or Illusionary Rod (7713)) | Blacksmithing [crafted] | 191.6 | yes | Illusionary Rod (7713, +0.42 DPS, sim-verified) [dungeon]; Steel Spear (250605, -9.05 DPS) [crafted]; Loksey's Training Stick (7710, -9.67 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) | World drop [world_drop] | 282.4 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS) [crafted]; Master Hunter's Rifle (17687, -0.10 DPS) [quest]; Master Hunter's Bow (17686, -1.20 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Frost Tiger Blade; ranged: Sniper Rifle

No-known-source sample (15 of 1260, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 5500000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 129.2. Weights run: 1.2s. Verify run: 1.6s. 1602 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.235 ± 0.084, crit=14.111 ± 0.595, hit=7.061 ± 0.559, melee_haste=6.806 ± 1.475

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) (or Knight-Lieutenant's Chain Helmet (220822)) | Lady Palanseer [vendor] | 233.3 | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -2.07 DPS) [dungeon]; Eye of Theradras (17715, -2.07 DPS) [dungeon] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 26.8 | yes | Sentinel's Medallion (19540, -0.16 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.52 DPS) [rep]; Ghostshard Talisman (7731, -0.74 DPS) [dungeon] |
| shoulder | Blood Guard's Chain Epaulets (220824) (or Knight-Lieutenant's Chain Epaulets (220825)) | Lady Palanseer [vendor] | 226.6 | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Epaulets (220823, -1.69 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -1.69 DPS) [vendor] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 24.6 | yes | Nightscape Cloak (8195, -0.16 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.26 DPS) [dungeon]; Imperial Cloak (6432, -0.39 DPS) [world_drop] |
| chest | Stone Guard's Chain Armor (220827) (or Knight's Chain Armor (220828)) | Lady Palanseer [vendor] | 231.1 | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Mail Armor (220826, -1.94 DPS) [vendor]; Knight's Mail Armor (223078, -1.94 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 33.5 | yes | Bracers of the Stone Princess (17714, -0.40 DPS, sim-verified) [dungeon]; Bloodlust Bracelets (14807, -0.52 DPS) [dungeon]; Wicked Leather Bracers (15084, -0.52 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 217.5 | yes | Dragonscale Gauntlets (8347, -0.45 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.16 DPS) [crafted]; Shadowskin Gloves (18238, -1.16 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 217.5 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20089, -0.70 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.16 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 228.8 | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.77 DPS, sim-verified) [crafted]; Stone Guard's Mail Legplates (220834, -1.82 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 70.6 | yes | Albino Crocscale Boots (17728, +0.94 DPS, sim-verified) [dungeon]; Fleetfoot Greaves (11627, -1.63 DPS) [dungeon]; Elven Chain Boots (13125, -1.76 DPS) [world] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 90.6 | yes | Ring of the Underwood (2951, -3.96 DPS) [world_drop]; Ironspine's Eye (7686, -4.09 DPS) [dungeon]; Protector's Band (19516, -4.09 DPS) [rep] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.3 | yes | Ring of the Underwood (2951, -0.64 DPS, sim-verified) [world_drop]; Ironspine's Eye (7686, -0.65 DPS) [dungeon]; Protector's Band (19516, -0.65 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 63.6 | yes | Tidal Charm (1404, -3.69 DPS) [vendor]; Guardian Talisman (1490, -3.69 DPS) [quest]; Ankh of Life (1713, -3.69 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 | yes | Ankh of Life (1713, +0.12 DPS, sim-verified) [world_drop]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Dawn's Edge (12774) (or Inventor's Focal Sword (17719)) | Blacksmithing [crafted] | 197.5 | yes | Frost Tiger Blade (3854, +0.00 DPS) [crafted]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Kindling Stave (11750, +0.00 DPS) [dungeon] |
| off_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 197.5 | yes | Thermotastic Egg Timer (9644, -11.07 DPS) [quest]; Grayson's Torch (1172, -11.46 DPS) [quest]; Rod of Molten Fire (2565, -11.46 DPS) [world_drop] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 448.7 | yes | Precisely Calibrated Boomstick (2100, -1.74 DPS) [world_drop]; Dark Iron Rifle (16004, -2.31 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -3.23 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Sentinel's Medallion; shoulder: Blood Guard's Chain Epaulets; back: Serpentskin Cloak; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; waist: Highlander's Chain Girdle; legs: Knight's Chain Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Dawn's Edge; off_hand: Inventor's Focal Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1602, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 188.3. Weights run: 1.2s. Verify run: 1.6s. 2288 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.339 ± 0.094, crit=20.135 ± 0.769, hit=not significant (0.000 ± 0.000), melee_haste=10.525 ± 1.657

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Cryptstalker Headpiece [quest] | 636.3 | yes | Lieutenant Commander's Chain Greathelm (227086, -1.29 DPS) [vendor]; Champion's Chain Helm (23251, -1.73 DPS) [vendor]; Champion's Chain Greathelm (227080, -10.81 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 305.9 | yes | Onyxia Tooth Pendant (18404, -0.43 DPS) [quest]; Amulet of the Darkmoon (19491, -14.84 DPS) [quest]; Choker of the Shifting Sands (21505, -14.98 DPS) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Cryptstalker Spaulders [quest] | 349.7 | yes | Lieutenant Commander's Chain Shoulders (23307, -1.46 DPS) [vendor]; Lieutenant Commander's Chain Shoulders (227068, -1.46 DPS) [pvp]; Champion's Chain Shoulders (23252, -10.65 DPS, sim-verified) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 60.8 | yes | Cape of the Black Baron (13340, -0.32 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.86 DPS) [rep]; Chromatic Cloak (18509, -2.90 DPS, sim-verified) [crafted] |
| chest | Legionnaire's Chain Armor (227083) (or Knight-Captain's Chain Armor (227089)) | Lady Palanseer [vendor] | 607.2 | yes | Knight-Captain's Chain Armor (227089, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Chain Hauberk (22874, -0.34 DPS) [vendor]; Knight-Captain's Chain Hauberk (23292, -0.34 DPS) [vendor] |
| wrist | Cryptstalker Wristguards (22443) | Cryptstalker Wristguards [quest] | 60.8 | yes | General's Chain Wristguards (16570, -0.80 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.93 DPS) [rep]; Marshal's Chain Bracers (16461, -9.27 DPS, sim-verified) [pvp] |
| hands | Cryptstalker Handguards (22441) | Cryptstalker Handguards [quest] | 338.0 | yes | General's Chain Gloves (16571, -0.40 DPS) [vendor]; Marshal's Chain Grips (231560, -0.61 DPS) [pvp]; Marshal's Chain Grips (16463, -9.05 DPS, sim-verified) [vendor] |
| waist | Cryptstalker Girdle (22442) | Cryptstalker Girdle [quest] | 335.7 | yes | Highlander's Leather Girdle (20045, -1.12 DPS) [rep]; Light Obsidian Belt (22195, -1.24 DPS) [crafted]; Highlander's Chain Girdle (20043, -12.24 DPS, sim-verified) [rep] |
| legs | Legionnaire's Chain Legplates (227079) (or Knight-Captain's Chain Legplates (227085)) | Lady Palanseer [vendor] | 607.2 | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Chain Legguards (22875, -0.34 DPS) [vendor]; Knight-Captain's Chain Legguards (23293, -0.34 DPS) [vendor] |
| feet | Cryptstalker Boots (22440) | Cryptstalker Boots [quest] | 77.2 | yes | Marshal's Chain Boots (16462, -0.93 DPS) [vendor]; General's Chain Sabatons (16569, -0.93 DPS) [vendor]; Striker's Footguards (21365, -9.29 DPS, sim-verified) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 297.9 | yes | Dragonslayer's Signet (18403, -0.91 DPS) [quest]; Ring of Entropy (18543, -0.91 DPS) [world]; Mindtear Band (20632, -0.91 DPS) [world] |
| finger2 | Band of the Penitent (13217) (or Dragonslayer's Signet (18403), Ring of Entropy (18543), Mindtear Band (20632), Band of Earthen Wrath (21179), Band of Earthen Might (21182), Don Rodrigo's Band (21563), Ritssyn's Ring of Chaos (21836), Ring of the Eternal Flame (23237)) | Houses of the Holy [quest] | 281.9 | yes | Dragonslayer's Signet (18403, +1.38 DPS, sim-verified) [quest]; Ring of Entropy (18543, +0.00 DPS) [world]; Mindtear Band (20632, +0.00 DPS) [world] |
| trinket1 | Onyxia Blood Talisman (18406) | Celebrating Good Times [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 309.9 | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 309.9 | yes | Grand Marshal's Handaxe (18827, +0.00 DPS, sim-verified) [vendor]; High Warlord's Cleaver (18828, +0.00 DPS) [vendor]; Grand Marshal's Dirk (18838, +0.00 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 755.5 | yes | High Warlord's Recurve (234559, -4.60 DPS) [pvp]; High Warlord's Crossbow (234560, -4.60 DPS) [pvp]; High Warlord's Street Sweeper (234561, -4.60 DPS) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Medallion of the Dawn; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Armor; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legplates; feet: Cryptstalker Boots; finger1: Don Julio's Band; finger2: Band of the Penitent; trinket1: Onyxia Blood Talisman; trinket2: Talisman of Arathor; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: The Purifier

No-known-source sample (15 of 2288, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 74.6. Weights run: 1.1s. Verify run: 1.2s. 383 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.067 ± 0.030, crit=8.180 ± 0.347, hit=4.699 ± 0.333, melee_haste=5.581 ± 0.983

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.5 | yes | Flying Tiger Goggles (4368, -0.91 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.00 DPS) [crafted]; Lucky Fishing Hat (19972, -1.00 DPS) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.4 | yes | Tarnished Locket (279870, -0.65 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.3 | yes | Double-Stitched Woolen Shoulders (4314, -0.59 DPS, sim-verified) [crafted]; Reinforced Woolen Shoulders (4315, -0.62 DPS) [crafted]; Forest Leather Mantle (4709, -0.62 DPS) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.4 | yes | Cape of the Brotherhood (5193, -0.23 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.25 DPS) [world_drop]; Hide of Lupos (3018, -0.25 DPS) [world] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 14.5 | yes | Trapper's Leather Armor (252491, +0.40 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.12 DPS) [crafted]; Heckler's Hide (286536, -0.25 DPS) [world] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.3 | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.25 DPS) [vendor]; Spare Part Bindings (279875, -0.25 DPS) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 114.5 | yes | Serpent Gloves (5970, +0.42 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -6.16 DPS) [dungeon]; Forest Leather Gloves (3058, -6.41 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.46 DPS) [quest]; Guardsman Belt (3429, -0.59 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 18.6 | yes | Brawler's Leather Pants (252500, +0.12 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.12 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.5 | yes | Footpads of the Fang (10411, -0.25 DPS) [dungeon]; Blackened Defias Boots (10402, -0.25 DPS, sim-verified) [dungeon]; Dark Leather Boots (2315, -0.37 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.4 | yes | Bounty Hunter's Ring (5351, -0.37 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon]; The 1 Ring (8350, -0.62 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 | yes | Bounty Hunter's Ring (5351, -0.13 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; The 1 Ring (8350, -0.37 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 22.5 | yes | Impaling Harpoon (5200, +0.55 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.49 DPS) [world]; Crescent Staff (6505, -0.49 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.3 | yes | Lovingly Crafted Boomstick (4372, -2.55 DPS) [crafted]; Venomstrike (6469, -2.67 DPS) [dungeon]; Cracked Blacksmith Hammer (285279, -3.23 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 383, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers

### Band 30 (troll, 0000000000000000-35305500000000000-000000000000000000)

Set DPS (verified): 91.4. Weights run: 1.2s. Verify run: 1.4s. 734 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.151 ± 0.048, crit=9.471 ± 0.389, hit=5.605 ± 0.389, melee_haste=not significant (1.703 ± 1.175)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.5 | yes | Tribal Worg Helm (6204, +0.27 DPS, sim-verified) [world]; Brawler's Leather Hood (252504, -0.26 DPS) [crafted]; Holy Shroud (2721, -1.28 DPS) [world_drop] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 17.2 | yes | Ghostshard Talisman (7731, -0.21 DPS, sim-verified) [dungeon]; Scout's Medallion (20442, -0.26 DPS) [rep]; Pendant of Myzrael (4614, -1.02 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.7 | yes | Mantle of Thieves (2264, -0.40 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.51 DPS) [crafted]; Insignia Mantle (4721, -0.51 DPS) [world_drop] |
| back | Cloak of Night (4447) (or Fenrus' Hide (6340), Glowing Lizardscale Cloak (6449), Swiftrunner Cape (6745)) | Rohh the Silent [world] | 12.9 | yes | Fenrus' Hide (6340, +0.00 DPS, sim-verified) [dungeon]; Glowing Lizardscale Cloak (6449, +0.00 DPS) [dungeon]; Swiftrunner Cape (6745, +0.00 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.1 | yes | Panther Armor (6670, -0.69 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.77 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.77 DPS) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410)) | Razormaw Matriarch [world] | 12.9 | yes | Insignia Bracers (6410, +0.00 DPS, sim-verified) [world_drop]; Madwolf Bracers (897, -0.13 DPS) [world]; Forest Leather Bracers (3202, -0.13 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 132.6 | yes | Pilferer's Gloves (7358, +0.51 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -6.94 DPS) [crafted]; Braced Handguards (6784, -6.99 DPS) [quest] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Deftkin Belt (16659, -0.20 DPS) [quest]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted] |
| legs | Dusky Leather Leggings (7373) | Leatherworking [crafted] | 28.0 | yes | Ferine Leggings (6690, +0.60 DPS, sim-verified) [dungeon]; Insignia Leggings (4054, -0.51 DPS) [world_drop]; Leggings of the Fang (10410, -0.51 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 17.2 | yes | Insignia Boots (4055, +0.00 DPS, sim-verified) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.4 | yes | Ring of Precision (1491, -0.38 DPS) [dungeon]; Legionnaire's Band (19513, -0.38 DPS) [rep]; Signet of the Zhevra (285330, -0.38 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.1 | yes | Legionnaire's Band (19513, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Ring of Precision (1491, -0.21 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Prison Shank (2941) (or Talon of Vultros (4454), Sentinel's Blade (212583), Scout's Blade (212587)) | The Stockade: Bruegal Ironknuckle [dungeon] | 12.9 | yes | Sentinel's Blade (212583, +0.00 DPS) [vendor]; Scout's Blade (212587, +0.00 DPS) [vendor]; Talon of Vultros (4454, -0.26 DPS, sim-verified) [world] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 | yes | Nightstalker Bow (6696, -1.70 DPS) [dungeon]; Silver Star (3463, -1.92 DPS, sim-verified) [quest]; Satchel of Bronze Bombs (285276, -2.86 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Cloak of Night; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Defiler's Chain Girdle; legs: Dusky Leather Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Alliance Outrunner's Sword; off_hand: Prison Shank; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 734, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword

### Band 40 (troll, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 103.7. Weights run: 1.2s. Verify run: 1.5s. 1258 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.199 ± 0.068, crit=13.685 ± 0.578, hit=6.174 ± 0.454, melee_haste=7.173 ± 1.232

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 191.6 | yes | Nightscape Headband (8176, +0.78 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -9.77 DPS) [crafted]; White Bandit Mask (10008, -9.90 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 24.2 | yes | Scout's Medallion (19537, -0.45 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.60 DPS) [dungeon]; Scout's Medallion (20442, -0.65 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.2 | yes | Nightscape Shoulders (8192, -0.71 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.76 DPS, sim-verified) [world_drop]; Mantle of Thieves (2264, -0.84 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | World drop [world_drop] | 17.6 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Darktide Cape (4114, -0.26 DPS) [quest]; Cloak of Night (4447, -0.26 DPS) [world] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 33.0 | yes | Tough Scorpid Breastplate (8203, +0.64 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.13 DPS) [crafted]; Hawkeye's Tunic (14592, -0.39 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.11 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.14 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.27 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 211.6 | yes | Dragonscale Gauntlets (8347, -0.40 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.18 DPS) [crafted]; Shadowskin Gloves (18238, -1.18 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 199.6 | yes | Defiler's Leather Girdle (20192, +0.43 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -10.38 DPS) [rep]; Defiler's Leather Girdle (20191, -10.38 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.2 | yes | Triprunner Dungarees (9624, -0.30 DPS, sim-verified) [quest]; Hawkeye's Breeches (14595, -0.91 DPS) [world]; Dusky Leather Leggings (7373, -1.04 DPS) [crafted] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.6 | yes | Dusky Boots (7390, -0.26 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.26 DPS) [crafted]; Imperial Leather Boots (6431, -0.30 DPS, sim-verified) [world_drop] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 22.0 | yes | Legionnaire's Band (19512, -0.26 DPS) [rep]; Disengagement Ring (276202, -0.26 DPS) [vendor]; Monkey Ring (6748, -0.39 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.8 | yes | Disengagement Ring (276202, -0.13 DPS) [vendor]; Legionnaire's Band (19512, -0.14 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.26 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 191.6 | yes | Frost Tiger Blade (3854, -1.23 DPS, sim-verified) [crafted]; Steel Spear (250605, -9.05 DPS) [crafted]; Loksey's Training Stick (7710, -9.67 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) | World drop [world_drop] | 282.4 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS) [crafted]; Master Hunter's Rifle (17687, -0.10 DPS) [quest]; Master Hunter's Bow (17686, -2.18 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Sniper Rifle

No-known-source sample (15 of 1258, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (troll, 5500000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 137.8. Weights run: 1.2s. Verify run: 1.6s. 1601 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.235 ± 0.084, crit=14.111 ± 0.595, hit=7.061 ± 0.559, melee_haste=6.806 ± 1.475

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) (or Knight-Lieutenant's Chain Helmet (220822)) | Lady Palanseer [vendor] | 233.3 | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -2.07 DPS) [dungeon]; Eye of Theradras (17715, -2.07 DPS) [dungeon] |
| neck | Scout's Medallion (19535) | Warsong Outriders [rep] | 26.8 | yes | Scout's Medallion (19536, -0.16 DPS, sim-verified) [rep]; Woven Ivy Necklace (19159, -0.39 DPS) [quest]; Scout's Medallion (19537, -0.52 DPS) [rep] |
| shoulder | Blood Guard's Chain Epaulets (220824) (or Knight-Lieutenant's Chain Epaulets (220825)) | Lady Palanseer [vendor] | 226.6 | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Epaulets (220823, -1.69 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -1.69 DPS) [vendor] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 24.6 | yes | Nightscape Cloak (8195, -0.16 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.26 DPS) [dungeon]; Imperial Cloak (6432, -0.39 DPS) [world_drop] |
| chest | Stone Guard's Chain Armor (220827) (or Knight's Chain Armor (220828)) | Lady Palanseer [vendor] | 231.1 | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Mail Armor (220826, -1.94 DPS) [vendor]; Knight's Mail Armor (223078, -1.94 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 33.5 | yes | Bracers of the Stone Princess (17714, +0.54 DPS, sim-verified) [dungeon]; Bloodlust Bracelets (14807, -0.52 DPS) [dungeon]; Wicked Leather Bracers (15084, -0.52 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 217.5 | yes | Dragonscale Gauntlets (8347, -0.41 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.16 DPS) [crafted]; Shadowskin Gloves (18238, -1.16 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 217.5 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20153, -0.70 DPS) [rep]; Highlander's Mail Girdle (20118, -1.16 DPS) [vendor] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 228.8 | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS) [vendor]; Stone Guard's Mail Legplates (220834, -1.82 DPS) [vendor]; Stormshroud Pants (15057, -2.39 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 70.6 | yes | Albino Crocscale Boots (17728, +0.69 DPS, sim-verified) [dungeon]; Fleetfoot Greaves (11627, -1.63 DPS) [dungeon]; Elven Chain Boots (13125, -1.76 DPS) [world] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 90.6 | yes | White Bone Band (11862, -3.86 DPS) [quest]; Ring of the Underwood (2951, -3.96 DPS) [world_drop]; Ironspine's Eye (7686, -4.09 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.3 | yes | Ring of the Underwood (2951, -0.52 DPS) [world_drop]; White Bone Band (11862, -0.58 DPS, sim-verified) [quest]; Ironspine's Eye (7686, -0.65 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 133.4 | yes | Ankh of Life (1713, -5.33 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -7.74 DPS) [vendor]; Guardian Talisman (1490, -7.74 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 63.6 | yes | Ankh of Life (1713, +0.29 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -3.69 DPS) [vendor]; Guardian Talisman (1490, -3.69 DPS) [quest] |
| main_hand | Dawn's Edge (12774) (or Inventor's Focal Sword (17719)) | Blacksmithing [crafted] | 197.5 | yes | Frost Tiger Blade (3854, +0.00 DPS) [crafted]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Inventor's Focal Sword (17719, +0.00 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 22.0 | yes | White Bone Shredder (11863, -0.37 DPS) [quest]; Thermotastic Egg Timer (9644, -0.89 DPS) [quest]; Grayson's Torch (1172, -1.28 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 448.7 | yes | Dark Iron Rifle (16004, -1.41 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -1.74 DPS) [world_drop]; Houndmaster's Bow (11628, -3.23 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Scout's Medallion; shoulder: Blood Guard's Chain Epaulets; back: Serpentskin Cloak; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; waist: Defiler's Chain Girdle; legs: Knight's Chain Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Vanquisher's Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1601, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (troll, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 200.0. Weights run: 1.2s. Verify run: 1.6s. 2285 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.339 ± 0.094, crit=20.135 ± 0.769, hit=not significant (0.000 ± 0.000), melee_haste=10.525 ± 1.657

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Cryptstalker Headpiece [quest] | 636.3 | yes | Lieutenant Commander's Chain Greathelm (227086, -1.29 DPS) [vendor]; Champion's Chain Helm (23251, -1.73 DPS) [vendor]; Champion's Chain Greathelm (227080, -10.73 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 305.9 | yes | Onyxia Tooth Pendant (18404, -0.43 DPS) [quest]; Amulet of the Darkmoon (19491, -14.84 DPS) [quest]; Choker of the Shifting Sands (21505, -14.98 DPS) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Cryptstalker Spaulders [quest] | 349.7 | yes | Lieutenant Commander's Chain Shoulders (23307, -1.46 DPS) [vendor]; Lieutenant Commander's Chain Shoulders (227068, -1.46 DPS) [pvp]; Champion's Chain Shoulders (23252, -10.33 DPS, sim-verified) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 60.8 | yes | Cape of the Black Baron (13340, -0.32 DPS) [dungeon]; Deathguard's Cloak (20068, -0.86 DPS) [rep]; Chromatic Cloak (18509, -3.42 DPS, sim-verified) [crafted] |
| chest | Legionnaire's Chain Armor (227083) (or Knight-Captain's Chain Armor (227089)) | Lady Palanseer [vendor] | 607.2 | yes | Knight-Captain's Chain Armor (227089, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Chain Hauberk (22874, -0.34 DPS) [vendor]; Knight-Captain's Chain Hauberk (23292, -0.34 DPS) [vendor] |
| wrist | Cryptstalker Wristguards (22443) | Cryptstalker Wristguards [quest] | 60.8 | yes | General's Chain Wristguards (16570, -0.80 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.93 DPS) [rep]; Marshal's Chain Bracers (16461, -9.18 DPS, sim-verified) [pvp] |
| hands | Cryptstalker Handguards (22441) | Cryptstalker Handguards [quest] | 338.0 | yes | General's Chain Gloves (16571, -0.40 DPS) [vendor]; Marshal's Chain Grips (231560, -0.61 DPS) [pvp]; Marshal's Chain Grips (16463, -8.56 DPS, sim-verified) [vendor] |
| waist | Cryptstalker Girdle (22442) | Cryptstalker Girdle [quest] | 335.7 | yes | Defiler's Leather Girdle (20190, -1.12 DPS) [rep]; Light Obsidian Belt (22195, -1.24 DPS) [crafted]; Defiler's Chain Girdle (20150, -11.92 DPS, sim-verified) [rep] |
| legs | Legionnaire's Chain Legplates (227079) (or Knight-Captain's Chain Legplates (227085)) | Lady Palanseer [vendor] | 607.2 | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Chain Legguards (22875, -0.34 DPS) [vendor]; Knight-Captain's Chain Legguards (23293, -0.34 DPS) [vendor] |
| feet | Cryptstalker Boots (22440) | Cryptstalker Boots [quest] | 77.2 | yes | Marshal's Chain Boots (16462, -0.93 DPS) [vendor]; General's Chain Sabatons (16569, -0.93 DPS) [vendor]; Striker's Footguards (21365, -7.48 DPS, sim-verified) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 297.9 | yes | Ring of Entropy (18543, -0.91 DPS) [world]; Mindtear Band (20632, -0.91 DPS) [world]; Band of the Penitent (13217, -4.94 DPS, sim-verified) [quest] |
| finger2 | Dragonslayer's Signet (18403) | For All To See [quest] | 281.9 | yes | Ring of Entropy (18543, +0.00 DPS) [world]; Mindtear Band (20632, +0.00 DPS) [world]; Band of the Penitent (13217, -1.98 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 84.0 | yes | Tidal Charm (1404, -4.77 DPS) [vendor]; Guardian Talisman (1490, -4.77 DPS) [quest]; Blazing Emblem (2802, -4.77 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 309.9 | yes | Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 309.9 | yes | Grand Marshal's Handaxe (18827, +0.00 DPS, sim-verified) [vendor]; High Warlord's Cleaver (18828, +0.00 DPS) [vendor]; Grand Marshal's Dirk (18838, +0.00 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 755.5 | yes | High Warlord's Recurve (234559, -4.60 DPS) [pvp]; High Warlord's Crossbow (234560, -4.60 DPS) [pvp]; High Warlord's Street Sweeper (234561, -4.60 DPS) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Medallion of the Dawn; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Armor; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legplates; feet: Cryptstalker Boots; finger1: Don Julio's Band; finger2: Dragonslayer's Signet; trinket2: Ankh of Life; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: The Purifier

No-known-source sample (15 of 2285, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

