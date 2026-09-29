# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (dwarf, 5420000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 75.1. Weights run: 1.5s. Verify run: 1.6s. 385 eligible items had no known source.

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

### Band 30 (dwarf, 5420001504000000-00000000000000000-000000000000000000)

Set DPS (verified): 100.0. Weights run: 1.6s. Verify run: 2.0s. 734 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.156 ± 0.049, crit=9.416 ± 0.378, hit=5.360 ± 0.397, melee_haste=not significant (1.327 ± 1.223)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 | yes | Tribal Worg Helm (6204, +0.44 DPS, sim-verified) [world]; Brawler's Leather Hood (252504, -0.25 DPS) [crafted]; Holy Shroud (2721, -1.27 DPS) [world_drop] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 17.2 | yes | Ghostshard Talisman (7731, -0.19 DPS, sim-verified) [dungeon]; Sentinel's Medallion (20444, -0.25 DPS) [rep]; Pendant of Myzrael (4614, -1.01 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.7 | yes | Mantle of Thieves (2264, -0.09 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.51 DPS) [crafted]; Insignia Mantle (4721, -0.51 DPS) [world_drop] |
| back | Cloak of Night (4447) (or Fenrus' Hide (6340), Glowing Lizardscale Cloak (6449)) | Rohh the Silent [world] | 12.9 | yes | Fenrus' Hide (6340, +0.00 DPS, sim-verified) [dungeon]; Glowing Lizardscale Cloak (6449, +0.00 DPS) [dungeon]; Cape of the Brotherhood (5193, -0.13 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.2 | yes | Tunic of Westfall (2041, -0.40 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.76 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.76 DPS) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410)) | Razormaw Matriarch [world] | 12.9 | yes | Insignia Bracers (6410, +0.00 DPS, sim-verified) [world_drop]; Madwolf Bracers (897, -0.13 DPS) [world]; Forest Leather Bracers (3202, -0.13 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 131.8 | yes | Pilferer's Gloves (7358, +0.53 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -6.81 DPS) [crafted]; Wolfclaw Gloves (1978, -6.99 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Skulker's Leather Belt (252520, -0.27 DPS) [crafted]; Stalker's Leather Belt (252521, -0.27 DPS) [crafted] |
| legs | Dusky Leather Leggings (7373) | Leatherworking [crafted] | 28.0 | yes | Ferine Leggings (6690, +0.68 DPS, sim-verified) [dungeon]; Insignia Leggings (4054, -0.51 DPS) [world_drop]; Leggings of the Fang (10410, -0.51 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 17.2 | yes | Insignia Boots (4055, +0.00 DPS, sim-verified) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.13 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.4 | yes | Ring of Precision (1491, -0.38 DPS) [dungeon]; Protector's Band (19517, -0.38 DPS) [rep]; Signet of the Zhevra (285330, -0.38 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.1 | yes | Protector's Band (19517, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Ring of Precision (1491, -0.29 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 | yes | Armor Piercer (6679, +0.41 DPS) [dungeon]; Bronze Dory (250603, +0.38 DPS) [crafted]; Kam's Walking Stick (2280, +0.28 DPS) [dungeon] |
| off_hand | Prison Shank (2941) (or Talon of Vultros (4454), Sentinel's Blade (212583), Scout's Blade (212587)) | The Stockade: Bruegal Ironknuckle [dungeon] | 12.9 | yes | Talon of Vultros (4454, +0.04 DPS, sim-verified) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor]; Scout's Blade (212587, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 | yes | Silver Star (3463, -1.43 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -1.68 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -2.83 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Cloak of Night; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Highlander's Chain Girdle; legs: Dusky Leather Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Alliance Outrunner's Sword; off_hand: Prison Shank; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 734, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword; 7958 Bronze Battle Axe

### Band 40 (dwarf, 5420001505001251-00000000000000000-000000000000000000)

Set DPS (verified): 124.8. Weights run: 1.8s. Verify run: 2.1s. 1260 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.142 ± 0.050, crit=10.937 ± 0.446, hit=6.223 ± 0.768, melee_haste=not significant (4.753 ± 1.936)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 153.1 | yes | Nightscape Headband (8176, +0.96 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -7.20 DPS) [crafted]; White Bandit Mask (10008, -7.33 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 23.6 | yes | Sentinel's Medallion (19541, -0.40 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.54 DPS) [dungeon]; Sentinel's Medallion (20444, -0.61 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 35.6 | yes | Nightscape Shoulders (8192, -0.68 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.72 DPS, sim-verified) [world_drop]; Mantle of Thieves (2264, -0.80 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | World drop [world_drop] | 17.1 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Yeti Fur Cloak (2805, -0.24 DPS) [quest]; Darktide Cape (4114, -0.24 DPS) [quest] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 32.1 | yes | Tough Scorpid Breastplate (8203, +0.25 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.12 DPS) [crafted]; Hawkeye's Tunic (14592, -0.36 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.14 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.16 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.28 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 173.1 | yes | Dragonscale Gauntlets (8347, -0.42 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.13 DPS) [crafted]; Shadowskin Gloves (18238, -1.13 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 161.1 | yes | Highlander's Leather Girdle (20116, +0.71 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -7.75 DPS) [rep]; Highlander's Leather Girdle (20117, -7.75 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.0 | yes | Hawkeye's Breeches (14595, -0.85 DPS) [world]; Dusky Leather Leggings (7373, -0.97 DPS) [crafted]; Triprunner Dungarees (9624, -1.54 DPS, sim-verified) [quest] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 27.8 | yes | Dusky Boots (7390, -0.24 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.24 DPS) [crafted]; Imperial Leather Boots (6431, -0.27 DPS, sim-verified) [world_drop] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 21.4 | yes | Protector's Band (19515, -0.24 DPS) [rep]; Disengagement Ring (276202, -0.24 DPS) [vendor]; Monkey Ring (6748, -0.36 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.3 | yes | Disengagement Ring (276202, -0.12 DPS) [vendor]; Protector's Band (19515, -0.13 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.24 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Frost Tiger Blade (3854) (or Illusionary Rod (7713)) | Blacksmithing [crafted] | 153.1 | yes | Illusionary Rod (7713, +0.19 DPS, sim-verified) [dungeon]; Steel Spear (250605, -6.52 DPS) [crafted]; Loksey's Training Stick (7710, -7.07 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Master Hunter's Bow (17686) | Big Game Hunter [quest] | 284.2 | yes | Sniper Rifle (3430, +0.87 DPS, sim-verified) [world_drop]; Mithril Heavy-bore Rifle (10510, -0.10 DPS) [crafted]; Master Hunter's Rifle (17687, -0.21 DPS) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Frost Tiger Blade; ranged: Master Hunter's Bow

No-known-source sample (15 of 1260, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 5420001505001251-35200000000000000-000000000000000000)

Set DPS (verified): 150.9. Weights run: 1.8s. Verify run: 2.3s. 1602 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.214 ± 0.065, crit=10.645 ± 0.461, hit=7.032 ± 0.782, melee_haste=not significant (3.939 ± 2.022)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) (or Knight-Lieutenant's Chain Helmet (220822)) | Lady Palanseer [vendor] | 184.5 | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -1.97 DPS) [dungeon]; Eye of Theradras (17715, -1.97 DPS) [dungeon] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 26.6 | yes | Sentinel's Medallion (19540, -0.14 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.49 DPS) [rep]; Ghostshard Talisman (7731, -0.70 DPS) [dungeon] |
| shoulder | Blood Guard's Chain Epaulets (220824) (or Knight-Lieutenant's Chain Epaulets (220825)) | Lady Palanseer [vendor] | 177.8 | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Epaulets (220823, -1.60 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -1.60 DPS) [vendor] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 24.4 | yes | Nightscape Cloak (8195, -0.14 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.25 DPS) [dungeon]; Imperial Cloak (6432, -0.37 DPS) [world_drop] |
| chest | Stone Guard's Chain Armor (220827) (or Knight's Chain Armor (220828)) | Lady Palanseer [vendor] | 182.2 | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Mail Armor (220826, -1.85 DPS) [vendor]; Knight's Mail Armor (223078, -1.85 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 33.2 | yes | Bracers of the Stone Princess (17714, +0.05 DPS, sim-verified) [dungeon]; Bloodlust Bracelets (14807, -0.49 DPS) [dungeon]; Wicked Leather Bracers (15084, -0.49 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 169.0 | yes | Dragonscale Gauntlets (8347, -0.40 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.11 DPS) [crafted]; Shadowskin Gloves (18238, -1.11 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 169.0 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20089, -0.67 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.11 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 180.0 | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS) [vendor]; Stone Guard's Mail Legplates (220834, -1.73 DPS) [vendor]; Stormshroud Pants (15057, -1.90 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 70.3 | yes | Albino Crocscale Boots (17728, +0.94 DPS, sim-verified) [dungeon]; Fleetfoot Greaves (11627, -1.57 DPS) [dungeon]; Elven Chain Boots (13125, -1.70 DPS) [world] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 90.3 | yes | Ring of the Underwood (2951, -3.80 DPS) [world_drop]; Ironspine's Eye (7686, -3.92 DPS) [dungeon]; Protector's Band (19516, -3.92 DPS) [rep] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.0 | yes | Ring of the Underwood (2951, -0.55 DPS, sim-verified) [world_drop]; Ironspine's Eye (7686, -0.62 DPS) [dungeon]; Protector's Band (19516, -0.62 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 63.3 | yes | Thunderbrew's Boot Flask (744, -3.52 DPS) [quest]; Tidal Charm (1404, -3.52 DPS) [vendor]; Guardian Talisman (1490, -3.52 DPS) [quest] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Dawn's Edge (12774) (or Inventor's Focal Sword (17719)) | Blacksmithing [crafted] | 149.0 | yes | Frost Tiger Blade (3854, +0.00 DPS) [crafted]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Kindling Stave (11750, +0.00 DPS) [dungeon] |
| off_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 149.0 | yes | Thermotastic Egg Timer (9644, -7.93 DPS) [quest]; Grayson's Torch (1172, -8.30 DPS) [quest]; Rod of Molten Fire (2565, -8.30 DPS) [world_drop] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 448.7 | yes | Dark Iron Rifle (16004, -1.50 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -1.69 DPS) [world_drop]; Houndmaster's Bow (11628, -3.10 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Sentinel's Medallion; shoulder: Blood Guard's Chain Epaulets; back: Serpentskin Cloak; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; waist: Highlander's Chain Girdle; legs: Knight's Chain Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Inventor's Focal Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1602, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 5420001505001251-35510000000000000-510000000000000000)

Set DPS (verified): 245.5. Weights run: 1.8s. Verify run: 2.4s. 2248 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.233 ± 0.071, crit=15.231 ± 0.576, hit=not significant (0.000 ± 0.000), melee_haste=12.796 ± 2.185

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Cryptstalker Headpiece [quest] | 495.7 | yes | Lieutenant Commander's Chain Helm (23306, -1.56 DPS) [vendor]; Lieutenant Commander's Chain Helm (227066, -1.56 DPS) [pvp]; Champion's Chain Helm (23251, -3.75 DPS, sim-verified) [vendor] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas: Kel'Thuzad [raid] | 452.5 | yes | Gem of Trapped Innocents (23057, -1.39 DPS) [raid]; Barbed Choker (21664, -10.46 DPS) [raid]; Medallion of the Dawn (22659, -11.53 DPS) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Cryptstalker Spaulders [quest] | 278.0 | yes | Lieutenant Commander's Chain Shoulders (23307, -1.32 DPS) [vendor]; Lieutenant Commander's Chain Shoulders (227068, -1.32 DPS) [pvp]; Champion's Chain Shoulders (23252, -3.68 DPS, sim-verified) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 58.1 | yes | Cape of the Black Baron (13340, -0.24 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.69 DPS) [rep]; Chromatic Cloak (18509, -3.19 DPS, sim-verified) [crafted] |
| chest | Legionnaire's Chain Hauberk (22874) (or Knight-Captain's Chain Hauberk (23292), Legionnaire's Chain Hauberk (227071)) | Lady Palanseer [vendor] | 462.2 | yes | Legionnaire's Chain Hauberk (227071, +0.00 DPS) [pvp]; Bloodsoul Breastplate (19690, -0.84 DPS) [crafted]; Knight-Captain's Chain Hauberk (23292, -3.41 DPS, sim-verified) [vendor] |
| wrist | Cryptstalker Wristguards (22443) | Cryptstalker Wristguards [quest] | 58.1 | yes | General's Chain Wristguards (16570, -0.72 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.84 DPS) [rep]; Marshal's Chain Bracers (16461, -2.80 DPS, sim-verified) [pvp] |
| hands | Cryptstalker Handguards (22441) | Cryptstalker Handguards [quest] | 266.8 | yes | General's Chain Gloves (16571, -0.36 DPS) [vendor]; Marshal's Chain Grips (231560, -0.50 DPS) [pvp]; Marshal's Chain Grips (16463, -3.07 DPS, sim-verified) [vendor] |
| waist | Cryptstalker Girdle (22442) | Cryptstalker Girdle [quest] | 264.6 | yes | Highlander's Chain Girdle (20043, -0.93 DPS) [rep]; Highlander's Leather Girdle (20045, -0.93 DPS) [rep]; Belt of Never-ending Agony (21586, -7.05 DPS, sim-verified) [raid] |
| legs | Legionnaire's Chain Legguards (22875) (or Knight-Captain's Chain Legguards (23293), Knight-Captain's Chain Legguards (227072), Legionnaire's Chain Legguards (227073)) | Lady Palanseer [vendor] | 462.2 | yes | Knight-Captain's Chain Legguards (227072, +0.00 DPS) [pvp]; Legionnaire's Chain Legguards (227073, +0.00 DPS) [pvp]; Knight-Captain's Chain Legguards (23293, -3.41 DPS, sim-verified) [vendor] |
| feet | Cryptstalker Boots (22440) | Cryptstalker Boots [quest] | 73.7 | yes | Marshal's Chain Boots (16462, -0.84 DPS) [vendor]; General's Chain Sabatons (16569, -0.84 DPS) [vendor]; Striker's Footguards (21365, -1.86 DPS, sim-verified) [quest] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas: Loatheb [raid] | 265.2 | yes | Quick Strike Ring (18821, -1.18 DPS) [raid]; Don Julio's Band (19325, -1.93 DPS) [rep]; Band of the Penitent (13217, -2.79 DPS) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj: Viscidus [raid] | 253.2 | yes | Quick Strike Ring (18821, -0.74 DPS, sim-verified) [raid]; Don Julio's Band (19325, -1.29 DPS) [rep]; Band of the Penitent (13217, -2.14 DPS) [quest] |
| trinket1 | Slayer's Crest (23041) | Naxxramas: Sapphiron [raid] | 64.0 | yes | Drake Fang Talisman (19406, -0.43 DPS) [raid]; Thunderbrew's Boot Flask (744, -3.43 DPS) [quest]; Eye of Diminution (23001, -6.70 DPS, sim-verified) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas: Maexxna [raid] | 213.2 | yes | Eye of Diminution (23001, -4.34 DPS, sim-verified) [raid]; Drake Fang Talisman (19406, -8.43 DPS) [raid]; Thunderbrew's Boot Flask (744, -11.43 DPS) [quest] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 241.2 | yes | Ironbark Staff (20069, +9.93 DPS) [rep]; Atiesh, Greatstaff of the Guardian (22630, +9.93 DPS) [quest]; High Warlord's War Staff (234549, +9.93 DPS) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 241.2 | yes | Grand Marshal's Handaxe (18827, +0.00 DPS, sim-verified) [vendor]; High Warlord's Cleaver (18828, +0.00 DPS) [vendor]; Grand Marshal's Dirk (18838, +0.00 DPS) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj: Ouro [raid] | 919.6 | yes | Huhuran's Stinger (21616, -12.39 DPS) [raid]; The Purifier (22656, -12.47 DPS) [quest]; High Warlord's Recurve (234559, -13.13 DPS) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Stormrage's Talisman of Seething; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Hauberk; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legguards; feet: Cryptstalker Boots; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2248, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 5420000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 74.6. Weights run: 1.5s. Verify run: 1.6s. 383 eligible items had no known source.

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

### Band 30 (troll, 5420001504000000-00000000000000000-000000000000000000)

Set DPS (verified): 102.5. Weights run: 1.6s. Verify run: 2.1s. 734 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.156 ± 0.049, crit=9.416 ± 0.378, hit=5.360 ± 0.397, melee_haste=not significant (1.327 ± 1.223)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tribal Worg Helm (6204) | Fenros [world] | 17.2 | yes | Brawler's Leather Hood (252504, +0.00 DPS) [crafted]; Holy Shroud (2721, -1.01 DPS) [world_drop]; Brawler's Leather Helm (252512, -1.31 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 17.2 | yes | Ghostshard Talisman (7731, -0.18 DPS, sim-verified) [dungeon]; Scout's Medallion (20442, -0.25 DPS) [rep]; Pendant of Myzrael (4614, -1.01 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.7 | yes | Mantle of Thieves (2264, -0.02 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.51 DPS) [crafted]; Insignia Mantle (4721, -0.51 DPS) [world_drop] |
| back | Cloak of Night (4447) (or Fenrus' Hide (6340), Glowing Lizardscale Cloak (6449), Swiftrunner Cape (6745)) | Rohh the Silent [world] | 12.9 | yes | Fenrus' Hide (6340, +0.00 DPS, sim-verified) [dungeon]; Glowing Lizardscale Cloak (6449, +0.00 DPS) [dungeon]; Swiftrunner Cape (6745, +0.00 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.2 | yes | Panther Armor (6670, -0.65 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.76 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.76 DPS) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410)) | Razormaw Matriarch [world] | 12.9 | yes | Insignia Bracers (6410, +0.00 DPS, sim-verified) [world_drop]; Madwolf Bracers (897, -0.13 DPS) [world]; Forest Leather Bracers (3202, -0.13 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 131.8 | yes | Pilferer's Gloves (7358, +0.56 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -6.81 DPS) [crafted]; Braced Handguards (6784, -6.87 DPS) [quest] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Deftkin Belt (16659, -0.20 DPS) [quest]; Skulker's Leather Belt (252520, -0.27 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Insignia Leggings (4054, -0.39 DPS) [world_drop]; Leggings of the Fang (10410, -0.39 DPS) [dungeon]; Dusky Leather Leggings (7373, -1.37 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 17.2 | yes | Insignia Boots (4055, +0.00 DPS, sim-verified) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.4 | yes | Ring of Precision (1491, -0.38 DPS) [dungeon]; Legionnaire's Band (19513, -0.38 DPS) [rep]; Signet of the Zhevra (285330, -0.38 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.1 | yes | Legionnaire's Band (19513, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Ring of Precision (1491, -0.37 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 | yes | Armor Piercer (6679, +0.41 DPS) [dungeon]; Bronze Dory (250603, +0.38 DPS) [crafted]; Kam's Walking Stick (2280, +0.28 DPS) [dungeon] |
| off_hand | Prison Shank (2941) (or Talon of Vultros (4454), Sentinel's Blade (212583), Scout's Blade (212587)) | The Stockade: Bruegal Ironknuckle [dungeon] | 12.9 | yes | Talon of Vultros (4454, +0.11 DPS, sim-verified) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor]; Scout's Blade (212587, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 | yes | Silver Star (3463, -1.56 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -1.68 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -2.83 DPS) [crafted] |

**New at 30:** head: Tribal Worg Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Cloak of Night; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Alliance Outrunner's Sword; off_hand: Prison Shank; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 734, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword

### Band 40 (troll, 5420001505001251-00000000000000000-000000000000000000)

Set DPS (verified): 127.0. Weights run: 1.8s. Verify run: 2.2s. 1258 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.142 ± 0.050, crit=10.937 ± 0.446, hit=6.223 ± 0.768, melee_haste=not significant (4.753 ± 1.936)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 153.1 | yes | Nightscape Headband (8176, +0.98 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -7.20 DPS) [crafted]; White Bandit Mask (10008, -7.33 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 23.6 | yes | Scout's Medallion (19537, -0.41 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.54 DPS) [dungeon]; Scout's Medallion (20442, -0.61 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 35.6 | yes | Nightscape Shoulders (8192, -0.68 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.72 DPS, sim-verified) [world_drop]; Mantle of Thieves (2264, -0.80 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | World drop [world_drop] | 17.1 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Darktide Cape (4114, -0.24 DPS) [quest]; Cloak of Night (4447, -0.24 DPS) [world] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 32.1 | yes | Tough Scorpid Breastplate (8203, +0.45 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.12 DPS) [crafted]; Hawkeye's Tunic (14592, -0.36 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Dusky Bracers (7378, -0.16 DPS) [crafted]; Imperial Leather Bracers (4061, -0.18 DPS, sim-verified) [world_drop]; Tough Scorpid Bracers (8205, -0.28 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 173.1 | yes | Dragonscale Gauntlets (8347, -0.44 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.13 DPS) [crafted]; Shadowskin Gloves (18238, -1.13 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 161.1 | yes | Defiler's Leather Girdle (20192, +0.69 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -7.75 DPS) [rep]; Defiler's Leather Girdle (20191, -7.75 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.0 | yes | Triprunner Dungarees (9624, -0.47 DPS, sim-verified) [quest]; Hawkeye's Breeches (14595, -0.85 DPS) [world]; Dusky Leather Leggings (7373, -0.97 DPS) [crafted] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 27.8 | yes | Dusky Boots (7390, -0.24 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.24 DPS) [crafted]; Imperial Leather Boots (6431, -0.28 DPS, sim-verified) [world_drop] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 21.4 | yes | Legionnaire's Band (19512, -0.24 DPS) [rep]; Disengagement Ring (276202, -0.24 DPS) [vendor]; Monkey Ring (6748, -0.36 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.3 | yes | Disengagement Ring (276202, -0.12 DPS) [vendor]; Legionnaire's Band (19512, -0.13 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.24 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| trinket2 | Rune of Duty (21567) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 153.1 | yes | Frost Tiger Blade (3854, -1.44 DPS, sim-verified) [crafted]; Steel Spear (250605, -6.52 DPS) [crafted]; Loksey's Training Stick (7710, -7.07 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) | World drop [world_drop] | 282.4 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS) [crafted]; Master Hunter's Rifle (17687, -0.11 DPS) [quest]; Master Hunter's Bow (17686, -2.14 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Illusionary Rod; ranged: Sniper Rifle

No-known-source sample (15 of 1258, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (troll, 5420001505001251-35200000000000000-000000000000000000)

Set DPS (verified): 153.0. Weights run: 1.8s. Verify run: 2.3s. 1601 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.214 ± 0.065, crit=10.645 ± 0.461, hit=7.032 ± 0.782, melee_haste=not significant (3.939 ± 2.022)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) (or Knight-Lieutenant's Chain Helmet (220822)) | Lady Palanseer [vendor] | 184.5 | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -1.97 DPS) [dungeon]; Eye of Theradras (17715, -1.97 DPS) [dungeon] |
| neck | Scout's Medallion (19535) | Warsong Outriders [rep] | 26.6 | yes | Scout's Medallion (19536, -0.13 DPS, sim-verified) [rep]; Woven Ivy Necklace (19159, -0.37 DPS) [quest]; Scout's Medallion (19537, -0.49 DPS) [rep] |
| shoulder | Blood Guard's Chain Epaulets (220824) (or Knight-Lieutenant's Chain Epaulets (220825)) | Lady Palanseer [vendor] | 177.8 | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Epaulets (220823, -1.60 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -1.60 DPS) [vendor] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 24.4 | yes | Nightscape Cloak (8195, -0.13 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.25 DPS) [dungeon]; Imperial Cloak (6432, -0.37 DPS) [world_drop] |
| chest | Stone Guard's Chain Armor (220827) (or Knight's Chain Armor (220828)) | Lady Palanseer [vendor] | 182.2 | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Mail Armor (220826, -1.85 DPS) [vendor]; Knight's Mail Armor (223078, -1.85 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 33.2 | yes | Bracers of the Stone Princess (17714, -0.22 DPS, sim-verified) [dungeon]; Bloodlust Bracelets (14807, -0.49 DPS) [dungeon]; Wicked Leather Bracers (15084, -0.49 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 169.0 | yes | Dragonscale Gauntlets (8347, -0.41 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.11 DPS) [crafted]; Shadowskin Gloves (18238, -1.11 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 169.0 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20153, -0.67 DPS) [rep]; Highlander's Mail Girdle (20118, -1.11 DPS) [vendor] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 180.0 | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS) [vendor]; Stone Guard's Mail Legplates (220834, -1.73 DPS) [vendor]; Stormshroud Pants (15057, -2.31 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 70.3 | yes | Albino Crocscale Boots (17728, +1.00 DPS, sim-verified) [dungeon]; Fleetfoot Greaves (11627, -1.57 DPS) [dungeon]; Elven Chain Boots (13125, -1.70 DPS) [world] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 90.3 | yes | White Bone Band (11862, -3.69 DPS) [quest]; Ring of the Underwood (2951, -3.80 DPS) [world_drop]; Ironspine's Eye (7686, -3.92 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.0 | yes | White Bone Band (11862, -0.46 DPS, sim-verified) [quest]; Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Ironspine's Eye (7686, -0.62 DPS) [dungeon] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, +7.42 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 63.3 | yes | Rune of the Guard Captain (19120, +3.90 DPS) [quest]; Tidal Charm (1404, -3.52 DPS) [vendor]; Guardian Talisman (1490, -3.52 DPS) [quest] |
| main_hand | Dawn's Edge (12774) (or Inventor's Focal Sword (17719)) | Blacksmithing [crafted] | 149.0 | yes | Frost Tiger Blade (3854, +0.00 DPS) [crafted]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Inventor's Focal Sword (17719, +0.00 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 22.0 | yes | White Bone Shredder (11863, -0.36 DPS) [quest]; Thermotastic Egg Timer (9644, -0.86 DPS) [quest]; Grayson's Torch (1172, -1.23 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 448.7 | yes | Dark Iron Rifle (16004, -0.83 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -1.69 DPS) [world_drop]; Houndmaster's Bow (11628, -3.10 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Scout's Medallion; shoulder: Blood Guard's Chain Epaulets; back: Serpentskin Cloak; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; waist: Defiler's Chain Girdle; legs: Knight's Chain Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Smoking Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Vanquisher's Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1601, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (troll, 5420001505001251-35510000000000000-510000000000000000)

Set DPS (verified): 247.8. Weights run: 1.8s. Verify run: 2.3s. 2245 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.233 ± 0.071, crit=15.231 ± 0.576, hit=not significant (0.000 ± 0.000), melee_haste=12.796 ± 2.185

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cryptstalker Headpiece (22438) | Cryptstalker Headpiece [quest] | 495.7 | yes | Lieutenant Commander's Chain Helm (23306, -1.56 DPS) [vendor]; Lieutenant Commander's Chain Helm (227066, -1.56 DPS) [pvp]; Champion's Chain Helm (23251, -2.88 DPS, sim-verified) [vendor] |
| neck | Stormrage's Talisman of Seething (23053) | Naxxramas: Kel'Thuzad [raid] | 452.5 | yes | Gem of Trapped Innocents (23057, -1.39 DPS) [raid]; Barbed Choker (21664, -10.46 DPS) [raid]; Medallion of the Dawn (22659, -11.53 DPS) [quest] |
| shoulder | Cryptstalker Spaulders (22439) | Cryptstalker Spaulders [quest] | 278.0 | yes | Lieutenant Commander's Chain Shoulders (23307, -1.32 DPS) [vendor]; Lieutenant Commander's Chain Shoulders (227068, -1.32 DPS) [pvp]; Champion's Chain Shoulders (23252, -2.80 DPS, sim-verified) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 58.1 | yes | Cape of the Black Baron (13340, -0.24 DPS) [dungeon]; Deathguard's Cloak (20068, -0.69 DPS) [rep]; Chromatic Cloak (18509, -3.11 DPS, sim-verified) [crafted] |
| chest | Legionnaire's Chain Hauberk (22874) (or Knight-Captain's Chain Hauberk (23292), Legionnaire's Chain Hauberk (227071)) | Lady Palanseer [vendor] | 462.2 | yes | Legionnaire's Chain Hauberk (227071, +0.00 DPS) [pvp]; Bloodsoul Breastplate (19690, -0.84 DPS) [crafted]; Knight-Captain's Chain Hauberk (23292, -3.49 DPS, sim-verified) [vendor] |
| wrist | Cryptstalker Wristguards (22443) | Cryptstalker Wristguards [quest] | 58.1 | yes | General's Chain Wristguards (16570, -0.72 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.84 DPS) [rep]; Marshal's Chain Bracers (16461, -2.40 DPS, sim-verified) [pvp] |
| hands | Cryptstalker Handguards (22441) | Cryptstalker Handguards [quest] | 266.8 | yes | General's Chain Gloves (16571, -0.36 DPS) [vendor]; Marshal's Chain Grips (231560, -0.50 DPS) [pvp]; Marshal's Chain Grips (16463, -1.83 DPS, sim-verified) [vendor] |
| waist | Cryptstalker Girdle (22442) | Cryptstalker Girdle [quest] | 264.6 | yes | Defiler's Chain Girdle (20150, -0.93 DPS) [rep]; Defiler's Leather Girdle (20190, -0.93 DPS) [rep]; Belt of Never-ending Agony (21586, -8.89 DPS, sim-verified) [raid] |
| legs | Legionnaire's Chain Legguards (22875) (or Knight-Captain's Chain Legguards (23293), Knight-Captain's Chain Legguards (227072), Legionnaire's Chain Legguards (227073)) | Lady Palanseer [vendor] | 462.2 | yes | Knight-Captain's Chain Legguards (227072, +0.00 DPS) [pvp]; Legionnaire's Chain Legguards (227073, +0.00 DPS) [pvp]; Knight-Captain's Chain Legguards (23293, -3.49 DPS, sim-verified) [vendor] |
| feet | Cryptstalker Boots (22440) | Cryptstalker Boots [quest] | 73.7 | yes | Marshal's Chain Boots (16462, -0.84 DPS) [vendor]; General's Chain Sabatons (16569, -0.84 DPS) [vendor]; Striker's Footguards (21365, -1.94 DPS, sim-verified) [quest] |
| finger1 | Band of Unnatural Forces (23038) | Naxxramas: Loatheb [raid] | 265.2 | yes | Quick Strike Ring (18821, -1.18 DPS) [raid]; Don Julio's Band (19325, -1.93 DPS) [rep]; Band of the Penitent (13217, -2.79 DPS) [quest] |
| finger2 | Ring of the Qiraji Fury (21677) | Ahn'Qiraj: Viscidus [raid] | 253.2 | yes | Quick Strike Ring (18821, -0.74 DPS, sim-verified) [raid]; Don Julio's Band (19325, -1.29 DPS) [rep]; Band of the Penitent (13217, -2.14 DPS) [quest] |
| trinket1 | Slayer's Crest (23041) | Naxxramas: Sapphiron [raid] | 64.0 | yes | Rune of the Guard Captain (19120, +1.07 DPS) [quest]; Drake Fang Talisman (19406, -0.43 DPS) [raid]; Eye of Diminution (23001, -7.35 DPS, sim-verified) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas: Maexxna [raid] | 213.2 | yes | Eye of Diminution (23001, -4.32 DPS, sim-verified) [raid]; Rune of the Guard Captain (19120, -6.93 DPS) [quest]; Drake Fang Talisman (19406, -8.43 DPS) [raid] |
| main_hand | Grand Marshal's Longsword (12584) | Captain O'Neal [vendor] | 241.2 | yes | Ironbark Staff (20220, +9.93 DPS) [rep]; Atiesh, Greatstaff of the Guardian (22630, +9.93 DPS) [quest]; High Warlord's War Staff (234549, +9.93 DPS) [pvp] |
| off_hand | High Warlord's Blade (16345) (or Grand Marshal's Handaxe (18827), High Warlord's Cleaver (18828), Grand Marshal's Dirk (18838), High Warlord's Razor (18840), Grand Marshal's Right Hand Blade (18843), High Warlord's Right Claw (18844), Grand Marshal's Left Hand Blade (18847), High Warlord's Left Claw (18848), Grand Marshal's Swiftblade (23456), High Warlord's Quickblade (23467), High Warlord's Blade (234552), High Warlord's Quickblade (234553), High Warlord's Cleaver (234554), High Warlord's Razor (234556), High Warlord's Right Claw (234557), High Warlord's Left Claw (234558), Grand Marshal's Swiftblade (234579), Grand Marshal's Handaxe (234580), Grand Marshal's Dirk (234582), Grand Marshal's Right Hand Blade (234583), Grand Marshal's Left Hand Blade (234584)) | Sergeant Thunderhorn [vendor] | 241.2 | yes | Grand Marshal's Handaxe (18827, +0.00 DPS, sim-verified) [vendor]; High Warlord's Cleaver (18828, +0.00 DPS) [vendor]; Grand Marshal's Dirk (18838, +0.00 DPS) [vendor] |
| ranged | Larvae of the Great Worm (23557) | Ahn'Qiraj: Ouro [raid] | 919.6 | yes | Huhuran's Stinger (21616, -12.39 DPS) [raid]; The Purifier (22656, -12.47 DPS) [quest]; High Warlord's Recurve (234559, -13.13 DPS) [pvp] |

**New at 60:** head: Cryptstalker Headpiece; neck: Stormrage's Talisman of Seething; shoulder: Cryptstalker Spaulders; back: Cloak of the Fallen God; chest: Legionnaire's Chain Hauberk; wrist: Cryptstalker Wristguards; hands: Cryptstalker Handguards; waist: Cryptstalker Girdle; legs: Legionnaire's Chain Legguards; feet: Cryptstalker Boots; finger1: Band of Unnatural Forces; finger2: Ring of the Qiraji Fury; trinket1: Slayer's Crest; trinket2: Kiss of the Spider; main_hand: Grand Marshal's Longsword; off_hand: High Warlord's Blade; ranged: Larvae of the Great Worm

No-known-source sample (15 of 2245, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

