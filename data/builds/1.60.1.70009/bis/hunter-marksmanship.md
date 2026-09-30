# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 73.7. Weights run: 1.6s. Verify run: 1.7s. 309 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.185 ± 0.045, crit=6.678 ± 0.268, hit=not significant (0.000 ± 0.000), melee_haste=3.462 ± 0.795

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.5 | yes | Shadow Goggles (4373, -1.04 DPS) [crafted]; Lucky Fishing Hat (19972, -1.04 DPS) [quest]; Flying Tiger Goggles (4368, -1.22 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 13.1 | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.78 DPS) [quest]; Tarnished Locket (279870, -0.78 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.9 | yes | Double-Stitched Woolen Shoulders (4314, -0.60 DPS, sim-verified) [crafted]; Reinforced Woolen Shoulders (4315, -0.65 DPS) [crafted]; Forest Leather Mantle (4709, -0.65 DPS) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.1 | yes | Cape of the Brotherhood (5193, -0.05 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [dungeon]; Hide of Lupos (3018, -0.26 DPS) [world] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 24.0 | yes | Brawler's Leather Armor (252490, -0.51 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.52 DPS) [crafted]; Dark Leather Tunic (2317, -0.65 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 10.9 | yes | Wolf Bracers (4794, -0.12 DPS, sim-verified) [vendor]; Bravo's Armbands (270015, -0.13 DPS) [quest]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 93.5 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -4.80 DPS) [dungeon]; Forest Leather Gloves (3058, -5.06 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.42 DPS) [quest]; Deviate Scale Belt (6468, -0.45 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.55 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.7 | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.13 DPS) [world]; Brawler's Leather Pants (252500, -0.17 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 17.5 | yes | Blackened Defias Boots (10402, -0.25 DPS, sim-verified) [dungeon]; Footpads of the Fang (10411, -0.26 DPS) [dungeon]; Dark Leather Boots (2315, -0.39 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 13.1 | yes | Lavishly Jeweled Ring (1156, -0.52 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world]; Minor Channeling Ring (1449, -0.78 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.7 | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.39 DPS) [world]; Minor Channeling Ring (1449, -0.52 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 297.8 | yes | Night Reaver (1318, +0.00 DPS) [dungeon]; Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 275.4 | yes | Blackfang (2236, -0.18 DPS, sim-verified) [dungeon]; Grayson's Torch (1172, -16.44 DPS) [quest]; Pulsating Hydra Heart (5183, -16.44 DPS) [world] |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 178.4 | yes | Lovingly Crafted Boomstick (4372, -2.53 DPS) [crafted]; Venomstrike (6469, -2.65 DPS) [dungeon]; Cracked Blacksmith Hammer (285279, -3.25 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Ranger Bow

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash

### Band 30 (dwarf, 0000000000000000-35305500000000000-000000000000000000)

Set DPS (verified): 88.8. Weights run: 1.6s. Verify run: 2.0s. 639 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.180 ± 0.048, crit=8.214 ± 0.316, hit=not significant (0.000 ± 0.000), melee_haste=7.178 ± 0.865

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.8 | yes | Tribal Worg Helm (6204, -0.08 DPS, sim-verified) [world]; Brawler's Leather Hood (252504, -0.26 DPS) [crafted]; Humbert's Helm (4724, -0.39 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 17.4 | yes | Ghostshard Talisman (7731, -0.20 DPS, sim-verified) [dungeon]; Sentinel's Medallion (20444, -0.26 DPS) [rep]; Erudite's Amulet (277204, -0.51 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 24.0 | yes | Dark Leather Shoulders (4252, -0.51 DPS) [crafted]; Insignia Mantle (4721, -0.51 DPS) [dungeon]; Mantle of Thieves (2264, -1.34 DPS, sim-verified) [dungeon] |
| back | Cloak of Night (4447) (or Fenrus' Hide (6340), Glowing Lizardscale Cloak (6449)) | Rohh the Silent [world] | 13.1 | yes | Fenrus' Hide (6340, +0.00 DPS, sim-verified) [dungeon]; Glowing Lizardscale Cloak (6449, +0.00 DPS) [dungeon]; Cape of the Brotherhood (5193, -0.13 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.5 | yes | Tunic of Westfall (2041, -0.40 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.77 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.77 DPS) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410)) | Razormaw Matriarch [world] | 13.1 | yes | Insignia Bracers (6410, +0.00 DPS, sim-verified) [dungeon]; Madwolf Bracers (897, -0.13 DPS) [world]; Forest Leather Bracers (3202, -0.13 DPS) [dungeon] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 115.0 | yes | Pilferer's Gloves (7358, +0.00 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -5.84 DPS) [crafted]; Wolfclaw Gloves (1978, -6.02 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Skulker's Leather Belt (252520, -0.26 DPS) [crafted]; Stalker's Leather Belt (252521, -0.26 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) | Gnomeregan: Caverndeep Ambusher [dungeon] | 30.5 | yes | Dusky Leather Leggings (7373, -0.14 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.27 DPS) [dungeon]; Insignia Leggings (4054, -0.64 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 17.4 | yes | Insignia Boots (4055, +0.00 DPS, sim-verified) [dungeon]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.13 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 | yes | Ring of Precision (1491, -0.39 DPS) [dungeon]; Protector's Band (19517, -0.39 DPS) [rep]; Signet of the Zhevra (285330, -0.39 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.3 | yes | Protector's Band (19517, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Ring of Precision (1491, -0.44 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (88.9 DPS) | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 405.9 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272082, +0.00 DPS) [vendor] |
| off_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 389.6 | yes | Bloody Brass Knuckles (7683, +0.00 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -6.95 DPS) [quest]; Grayson's Torch (1172, -23.00 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 | yes | Nightstalker Bow (6696, -1.68 DPS) [dungeon]; Silver Star (3463, -2.71 DPS, sim-verified) [quest]; Satchel of Bronze Bombs (285276, -2.84 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Cloak of Night; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Highlander's Chain Girdle; legs: Petrolspill Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Pronged Reaver; off_hand: Swinetusk Shank; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 639, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword; 7958 Bronze Battle Axe; 8178 Training Sword

### Band 40 (dwarf, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 104.3. Weights run: 1.6s. Verify run: 2.2s. 1144 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.222 ± 0.063, crit=11.681 ± 0.473, hit=not significant (0.000 ± 0.000), melee_haste=not significant (3.437 ± 0.971)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 163.5 | yes | Nightscape Headband (8176, +0.00 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -8.03 DPS) [crafted]; White Bandit Mask (10008, -8.16 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 24.4 | yes | Sentinel's Medallion (19541, -0.42 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.61 DPS) [dungeon]; Sentinel's Medallion (20444, -0.65 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.4 | yes | Nightscape Shoulders (8192, -0.70 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.76 DPS, sim-verified) [dungeon]; Mantle of Thieves (2264, -0.83 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | Gnomeregan: Caverndeep Burrower [dungeon] | 17.8 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Yeti Fur Cloak (2805, -0.26 DPS) [quest]; Darktide Cape (4114, -0.26 DPS) [quest] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 33.3 | yes | Tough Scorpid Breastplate (8203, +0.00 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.13 DPS) [crafted]; Hawkeye's Tunic (14592, -0.39 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.11 DPS, sim-verified) [dungeon]; Dusky Bracers (7378, -0.13 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 183.5 | yes | Dragonscale Gauntlets (8347, -0.39 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 171.5 | yes | Highlander's Leather Girdle (20116, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -8.66 DPS) [rep]; Highlander's Leather Girdle (20117, -8.66 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | Gnomeregan: Leprous Machinesmith [dungeon] | 46.7 | yes | Petrolspill Leggings (9509, -0.91 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.91 DPS) [world]; Triprunner Dungarees (9624, -1.20 DPS, sim-verified) [quest] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 28.9 | yes | Dusky Boots (7390, -0.26 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.26 DPS) [crafted]; Imperial Leather Boots (6431, -0.28 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 22.2 | yes | Protector's Band (19515, -0.26 DPS) [rep]; Disengagement Ring (276202, -0.26 DPS) [vendor]; Monkey Ring (6748, -0.39 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 20.0 | yes | Disengagement Ring (276202, -0.13 DPS) [vendor]; Protector's Band (19515, -0.15 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.26 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | sim-verified (102.6 DPS) | yes | Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 535.8 | yes | Staff of Jordan (873, +0.00 DPS) [dungeon]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 533.4 | yes | Curve-bladed Ripper (2815, -2.56 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -15.24 DPS) [quest]; Grayson's Torch (1172, -31.29 DPS) [quest] |
| ranged | Sniper Rifle (3430) | Gnomeregan: Dark Iron Agent [dungeon] | sim-verified (104.3 DPS) | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS) [crafted]; Master Hunter's Rifle (17687, -0.09 DPS) [quest]; Master Hunter's Bow (17686, -1.64 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 1144, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (dwarf, 5500000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 130.0. Weights run: 1.6s. Verify run: 2.2s. 1478 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.307 ± 0.077, crit=13.494 ± 0.534, hit=not significant (0.000 ± 0.000), melee_haste=13.786 ± 1.194

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) (or Knight-Lieutenant's Chain Helmet (220822)) | Lady Palanseer [vendor] | 225.8 | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -2.16 DPS) [dungeon]; Eye of Theradras (17715, -2.16 DPS) [dungeon] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 27.7 | yes | Sentinel's Medallion (19540, -0.17 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.54 DPS) [rep]; Ghostshard Talisman (7731, -0.80 DPS) [dungeon] |
| shoulder | Blood Guard's Chain Epaulets (220824) (or Knight-Lieutenant's Chain Epaulets (220825)) | Lady Palanseer [vendor] | 218.9 | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Epaulets (220823, -1.75 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -1.75 DPS) [vendor] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 25.4 | yes | Nightscape Cloak (8195, -0.17 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.27 DPS) [dungeon]; Imperial Cloak (6432, -0.40 DPS) [dungeon] |
| chest | Stone Guard's Chain Armor (220827) (or Knight's Chain Armor (220828)) | Lady Palanseer [vendor] | 223.5 | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Mail Armor (220826, -2.02 DPS) [vendor]; Knight's Mail Armor (223078, -2.02 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.6 | yes | Bracers of the Stone Princess (17714, -0.24 DPS, sim-verified) [dungeon]; Bloodlust Bracelets (14807, -0.54 DPS) [dungeon]; Wicked Leather Bracers (15084, -0.54 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 208.9 | yes | Dragonscale Gauntlets (8347, -0.42 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 208.9 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20089, -0.70 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.17 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | sim-verified (130.0 DPS) | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS) [vendor]; Stone Guard's Mail Legplates (220834, -1.89 DPS) [vendor]; Stormshroud Pants (15057, -1.98 DPS, sim-verified) [crafted] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 46.1 | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.27 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.40 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.3 | yes | Ironspine's Eye (7686, -0.67 DPS) [dungeon]; Protector's Band (19516, -0.67 DPS) [rep]; Blackstone Ring (17713, -0.72 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 23.1 | yes | Protector's Band (19516, -0.13 DPS) [rep]; Ironspine's Eye (7686, -0.17 DPS, sim-verified) [dungeon]; Blackstone Ring (17713, -0.18 DPS) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (128.2 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Smotts' Compass (4130, +0.00 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | sim-verified (128.7 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Smoking Heart of the Mountain (11811, -0.34 DPS, sim-verified) [crafted] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (128.7 DPS) | yes | Flurry Axe (871, +0.00 DPS, sim-verified) [dungeon]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Soulkeeper (1607, -0.89 DPS) [dungeon] |
| off_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 791.8 | yes | Thorium Cestus (250614, +0.00 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -11.83 DPS) [dungeon]; Shoni's Disarming Tool (9608, -30.22 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (128.7 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.70 DPS) [dungeon]; Dark Iron Rifle (16004, -2.25 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -3.25 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Sentinel's Medallion; shoulder: Blood Guard's Chain Epaulets; back: Serpentskin Cloak; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; waist: Highlander's Chain Girdle; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Thunderbrew's Boot Flask; trinket2: Ankh of Life; main_hand: Dawn's Edge; off_hand: Inventor's Focal Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1478, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 207.3. Weights run: 1.7s. Verify run: 2.2s. 2069 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.339 ± 0.094, crit=20.135 ± 0.769, hit=not significant (0.000 ± 0.000), melee_haste=10.525 ± 1.657

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Headpiece (239540) | Leonid Barthalomew the Revered [vendor] | 678.4 | yes | Cryptstalker Headpiece (22438, -0.65 DPS, sim-verified) [quest]; Champion's Chain Greathelm (227080, -3.68 DPS) [vendor]; Lieutenant Commander's Chain Greathelm (227086, -3.68 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (189.9 DPS) | yes | Onyxia Tooth Pendant (18404, -0.43 DPS) [quest]; Blazefury Medallion (17111, -0.83 DPS, sim-verified) [world]; Amulet of the Darkmoon (19491, -14.84 DPS) [quest] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | sim-verified (192.1 DPS) | yes | Cryptstalker Spaulders (22439, -2.93 DPS) [quest]; Dawnstalker Spaulders (239542, -2.99 DPS, sim-verified) [vendor]; Darkspear Pauldrons (272105, -3.20 DPS) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | sim-verified (192.6 DPS) | yes | Cape of the Black Baron (13340, -0.32 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.86 DPS) [rep]; Chromatic Cloak (18509, -3.45 DPS, sim-verified) [crafted] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (191.7 DPS) | yes | Dawnstalker Tunic (239543, -2.55 DPS, sim-verified) [vendor]; Legionnaire's Chain Armor (227083, -6.88 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -6.88 DPS) [vendor] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | 361.4 | yes | Dawnstalker Wristguards (239544, -0.93 DPS, sim-verified) [vendor]; Cryptstalker Wristguards (22443, -17.06 DPS) [quest]; Marshal's Chain Bracers (16461, -17.86 DPS) [pvp] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | sim-verified (191.8 DPS) | yes | Dawnstalker Handguards (239539, -2.63 DPS, sim-verified) [vendor]; Cryptstalker Handguards (22441, -3.94 DPS) [quest]; Marshal's Chain Grips (16463, -4.34 DPS) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 384.1 | yes | Dawnstalker Girdle (239538, -1.43 DPS, sim-verified) [vendor]; Cryptstalker Girdle (22442, -2.75 DPS) [quest]; Highlander's Chain Girdle (20043, -3.87 DPS) [rep] |
| legs | Dawnstalker Legguards (239541) | Leonid Barthalomew the Revered [vendor] | 673.7 | yes | Sentinel's Chain Leggings (237819, -1.29 DPS, sim-verified) [vendor]; Sentinel's Leather Pants (237818, -2.65 DPS) [vendor]; Legionnaire's Chain Legplates (227079, -3.78 DPS) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | sim-verified (193.6 DPS) | yes | Dawnstalker Boots (239537, -4.43 DPS, sim-verified) [vendor]; Cryptstalker Boots (22440, -19.24 DPS) [quest]; Striker's Footguards (21365, -19.50 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 297.9 | yes | Ring of Entropy (18543, -0.91 DPS) [world]; Mindtear Band (20632, -0.91 DPS) [world]; Band of the Penitent (13217, -2.65 DPS, sim-verified) [quest] |
| finger2 | Dragonslayer's Signet (18403) | Celebrating Good Times [quest] | sim-verified (191.9 DPS) | yes | Ring of Entropy (18543, +0.00 DPS) [world]; Mindtear Band (20632, +0.00 DPS) [world]; Band of the Penitent (13217, -2.69 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | sim-verified (188.9 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Smotts' Compass (4130, +0.00 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (189.9 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Serenity Field (272439, -0.69 DPS, sim-verified) [vendor] |
| main_hand | High Warlord's Hacker (235476) | Sergeant Thunderhorn [vendor] | sim-verified (189.9 DPS) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -1.72 DPS, sim-verified) [rep] |
| off_hand | Grand Marshal's Hacker (235481) | Captain O'Neal [vendor] | 1401.2 | yes | High Warlord's Shiv (235478, +0.00 DPS, sim-verified) [vendor]; High Warlord's Left Claw (234558, -0.22 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, -0.22 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (189.9 DPS) | yes | High Warlord's Recurve (234559, -4.60 DPS) [vendor]; High Warlord's Crossbow (234560, -4.60 DPS) [vendor]; Dark Iron Rifle (16004, -7.46 DPS, sim-verified) [crafted] |

**New at 60:** head: Dawnstalker Headpiece; neck: Medallion of the Dawn; shoulder: Dawnstalker Pauldrons; back: Cloak of the Fallen God; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Legguards; feet: Dawnstalker Greaves; finger1: Don Julio's Band; finger2: Dragonslayer's Signet; trinket1: Ankh of Life; trinket2: Thunderbrew's Boot Flask; main_hand: High Warlord's Hacker; off_hand: Grand Marshal's Hacker; ranged: The Purifier

No-known-source sample (15 of 2069, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 73.5. Weights run: 1.6s. Verify run: 1.7s. 315 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.185 ± 0.045, crit=6.678 ± 0.268, hit=not significant (0.000 ± 0.000), melee_haste=3.462 ± 0.795

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.5 | yes | Flying Tiger Goggles (4368, -0.96 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.04 DPS) [crafted]; Lucky Fishing Hat (19972, -1.04 DPS) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 13.1 | yes | Erudite's Amulet (277204, -0.27 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.78 DPS) [quest]; Tarnished Locket (279870, -0.78 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.9 | yes | Double-Stitched Woolen Shoulders (4314, -0.61 DPS, sim-verified) [crafted]; Reinforced Woolen Shoulders (4315, -0.65 DPS) [crafted]; Forest Leather Mantle (4709, -0.65 DPS) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.1 | yes | Cape of the Brotherhood (5193, -0.25 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [dungeon]; Hide of Lupos (3018, -0.26 DPS) [world] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 15.3 | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.13 DPS) [crafted]; Heckler's Hide (286536, -0.26 DPS) [world] |
| wrist | Forest Leather Bracers (3202) | Gnomeregan: Caverndeep Burrower [dungeon] | 10.9 | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor]; Spare Part Bindings (279875, -0.26 DPS) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 93.5 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -4.80 DPS) [dungeon]; Forest Leather Gloves (3058, -5.06 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.42 DPS) [quest]; Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.55 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.7 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.13 DPS) [world] |
| feet | Feet of the Lynx (1121) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 17.5 | yes | Footpads of the Fang (10411, -0.26 DPS) [dungeon]; Blackened Defias Boots (10402, -0.27 DPS, sim-verified) [dungeon]; Dark Leather Boots (2315, -0.39 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 13.1 | yes | Bounty Hunter's Ring (5351, -0.39 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.52 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.7 | yes | Bounty Hunter's Ring (5351, -0.13 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; The 1 Ring (8350, -0.39 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 297.8 | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Crescent Staff (6505, +0.00 DPS) [quest]; Living Root (6631, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 275.4 | yes | Wingblade (6504, +0.00 DPS, sim-verified) [quest]; Grayson's Torch (1172, -16.44 DPS) [quest]; Nightglow Concoction (3451, -16.44 DPS) [quest] |
| ranged | Ranger Bow (3021) | Blackfathom Deeps: Murkshallow Snapclaw [dungeon] | 178.4 | yes | Lovingly Crafted Boomstick (4372, -2.53 DPS) [crafted]; Venomstrike (6469, -2.65 DPS) [dungeon]; Cracked Blacksmith Hammer (285279, -3.42 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Ranger Bow

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 30 (troll, 0000000000000000-35305500000000000-000000000000000000)

Set DPS (verified): 90.4. Weights run: 1.6s. Verify run: 2.1s. 650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.180 ± 0.048, crit=8.214 ± 0.316, hit=not significant (0.000 ± 0.000), melee_haste=7.178 ± 0.865

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.8 | yes | Tribal Worg Helm (6204, +0.00 DPS, sim-verified) [world]; Brawler's Leather Hood (252504, -0.26 DPS) [crafted]; Humbert's Helm (4724, -0.39 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 17.4 | yes | Ghostshard Talisman (7731, -0.19 DPS, sim-verified) [dungeon]; Scout's Medallion (20442, -0.26 DPS) [rep]; Erudite's Amulet (277204, -0.51 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 24.0 | yes | Mantle of Thieves (2264, -0.22 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.51 DPS) [crafted]; Insignia Mantle (4721, -0.51 DPS) [dungeon] |
| back | Cloak of Night (4447) (or Fenrus' Hide (6340), Glowing Lizardscale Cloak (6449), Swiftrunner Cape (6745)) | Rohh the Silent [world] | 13.1 | yes | Fenrus' Hide (6340, +0.00 DPS, sim-verified) [dungeon]; Glowing Lizardscale Cloak (6449, +0.00 DPS) [dungeon]; Swiftrunner Cape (6745, +0.00 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.5 | yes | Panther Armor (6670, -0.66 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.77 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.77 DPS) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410)) | Razormaw Matriarch [world] | 13.1 | yes | Insignia Bracers (6410, +0.00 DPS, sim-verified) [dungeon]; Madwolf Bracers (897, -0.13 DPS) [world]; Forest Leather Bracers (3202, -0.13 DPS) [dungeon] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 115.0 | yes | Pilferer's Gloves (7358, +0.00 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -5.84 DPS) [crafted]; Braced Handguards (6784, -5.89 DPS) [quest] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Deftkin Belt (16659, -0.19 DPS) [quest]; Skulker's Leather Belt (252520, -0.26 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) | Gnomeregan: Caverndeep Ambusher [dungeon] | 30.5 | yes | Dusky Leather Leggings (7373, -0.13 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.27 DPS) [dungeon]; Insignia Leggings (4054, -0.64 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 17.4 | yes | Insignia Boots (4055, +0.00 DPS, sim-verified) [dungeon]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 | yes | Ring of Precision (1491, -0.39 DPS) [dungeon]; Legionnaire's Band (19513, -0.39 DPS) [rep]; Signet of the Zhevra (285330, -0.39 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.3 | yes | Ring of Precision (1491, +0.00 DPS, sim-verified) [dungeon]; Legionnaire's Band (19513, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (89.6 DPS) | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 405.9 | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (90.4 DPS) | yes | Swinetusk Shank (6691, -1.01 DPS, sim-verified) [dungeon]; Grayson's Torch (1172, -22.59 DPS) [quest]; Rod of Molten Fire (2565, -22.59 DPS) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 | yes | Silver Star (3463, -1.58 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -1.68 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -2.84 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Cloak of Night; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Defiler's Chain Girdle; legs: Petrolspill Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 650, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7957 Bronze Greatsword; 7958 Bronze Battle Axe

### Band 40 (troll, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 104.8. Weights run: 1.6s. Verify run: 2.1s. 1150 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.222 ± 0.063, crit=11.681 ± 0.473, hit=not significant (0.000 ± 0.000), melee_haste=not significant (3.437 ± 0.971)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 163.5 | yes | Nightscape Headband (8176, +0.00 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -8.03 DPS) [crafted]; White Bandit Mask (10008, -8.16 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 24.4 | yes | Scout's Medallion (19537, -0.46 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.61 DPS) [dungeon]; Scout's Medallion (20442, -0.65 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.4 | yes | Nightscape Shoulders (8192, -0.70 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.76 DPS, sim-verified) [dungeon]; Mantle of Thieves (2264, -0.83 DPS) [dungeon] |
| back | Imperial Cloak (6432) (or Parachute Cloak (10518)) | Gnomeregan: Caverndeep Burrower [dungeon] | 17.8 | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Darktide Cape (4114, -0.26 DPS) [quest]; Cloak of Night (4447, -0.26 DPS) [world] |
| chest | Nightscape Tunic (8175) (or Tough Scorpid Breastplate (8203)) | Leatherworking [crafted] | 33.3 | yes | Tough Scorpid Breastplate (8203, +0.00 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.13 DPS) [crafted]; Hawkeye's Tunic (14592, -0.39 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.11 DPS, sim-verified) [dungeon]; Dusky Bracers (7378, -0.13 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 183.5 | yes | Dragonscale Gauntlets (8347, -0.40 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 171.5 | yes | Defiler's Leather Girdle (20192, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -8.66 DPS) [rep]; Defiler's Leather Girdle (20191, -8.66 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | Gnomeregan: Leprous Machinesmith [dungeon] | 46.7 | yes | Triprunner Dungarees (9624, -0.69 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.91 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.91 DPS) [world] |
| feet | Swampwalker Boots (2276) | Razorfen Downs: Withered Warrior [dungeon] | 28.9 | yes | Dusky Boots (7390, -0.26 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.26 DPS) [crafted]; Imperial Leather Boots (6431, -0.32 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | Razorfen Downs: Splinterbone Warrior [dungeon] | 22.2 | yes | Legionnaire's Band (19512, -0.26 DPS) [rep]; Disengagement Ring (276202, -0.26 DPS) [vendor]; Monkey Ring (6748, -0.39 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 20.0 | yes | Disengagement Ring (276202, -0.13 DPS) [vendor]; Legionnaire's Band (19512, -0.17 DPS, sim-verified) [rep]; Monkey Ring (6748, -0.26 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | sim-verified (103.5 DPS) | yes | Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 535.8 | yes | Staff of Jordan (873, +0.00 DPS) [dungeon]; Illusionary Rod (7713, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 533.4 | yes | Curve-bladed Ripper (2815, -3.10 DPS, sim-verified) [dungeon]; Grayson's Torch (1172, -31.29 DPS) [quest]; Rod of Molten Fire (2565, -31.29 DPS) [dungeon] |
| ranged | Sniper Rifle (3430) | Gnomeregan: Dark Iron Agent [dungeon] | sim-verified (104.8 DPS) | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS) [crafted]; Master Hunter's Rifle (17687, -0.09 DPS) [quest]; Master Hunter's Bow (17686, -1.37 DPS, sim-verified) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Imperial Cloak; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Ironspine's Eye; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 1150, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 5500000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 135.7. Weights run: 1.6s. Verify run: 2.1s. 1485 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.307 ± 0.077, crit=13.494 ± 0.534, hit=not significant (0.000 ± 0.000), melee_haste=13.786 ± 1.194

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) (or Knight-Lieutenant's Chain Helmet (220822)) | Lady Palanseer [vendor] | 225.8 | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -2.16 DPS) [dungeon]; Eye of Theradras (17715, -2.16 DPS) [dungeon] |
| neck | Scout's Medallion (19535) | Warsong Outriders [rep] | 27.7 | yes | Scout's Medallion (19536, -0.17 DPS, sim-verified) [rep]; Woven Ivy Necklace (19159, -0.40 DPS) [quest]; Scout's Medallion (19537, -0.54 DPS) [rep] |
| shoulder | Blood Guard's Chain Epaulets (220824) (or Knight-Lieutenant's Chain Epaulets (220825)) | Lady Palanseer [vendor] | 218.9 | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Mail Epaulets (220823, -1.75 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -1.75 DPS) [vendor] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 25.4 | yes | Nightscape Cloak (8195, -0.17 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.27 DPS) [dungeon]; Imperial Cloak (6432, -0.40 DPS) [dungeon] |
| chest | Stone Guard's Chain Armor (220827) (or Knight's Chain Armor (220828)) | Lady Palanseer [vendor] | 223.5 | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Mail Armor (220826, -2.02 DPS) [vendor]; Knight's Mail Armor (223078, -2.02 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.6 | yes | Bloodlust Bracelets (14807, -0.54 DPS) [dungeon]; Wicked Leather Bracers (15084, -0.54 DPS) [crafted]; Bracers of the Stone Princess (17714, -0.66 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 208.9 | yes | Dragonscale Gauntlets (8347, -0.39 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 208.9 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20153, -0.70 DPS) [rep]; Highlander's Mail Girdle (20118, -1.17 DPS) [vendor] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | sim-verified (135.7 DPS) | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.63 DPS, sim-verified) [crafted]; Stone Guard's Mail Legplates (220834, -1.89 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 46.1 | yes | Fleetfoot Greaves (11627, -0.24 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.27 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.40 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.3 | yes | Ring of the Underwood (2951, -0.54 DPS) [dungeon]; Ironspine's Eye (7686, -0.67 DPS) [dungeon]; Legionnaire's Band (19511, -0.67 DPS) [rep] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Ring of the Underwood (2951, -0.03 DPS, sim-verified) [dungeon]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Legionnaire's Band (19511, -0.19 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (132.9 DPS) | yes | Guardian Talisman (1490, -4.91 DPS) [quest]; Blazing Emblem (2802, -4.91 DPS) [dungeon]; Smotts' Compass (4130, -4.91 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | sim-verified (134.2 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.85 DPS, sim-verified) [crafted] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (134.2 DPS) | yes | Flurry Axe (871, +0.00 DPS, sim-verified) [dungeon]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Soulkeeper (1607, -0.89 DPS) [dungeon] |
| off_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 791.8 | yes | Thorium Cestus (250614, +0.00 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -11.83 DPS) [dungeon]; White Bone Shredder (11863, -14.26 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (134.2 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.70 DPS) [dungeon]; Dark Iron Rifle (16004, -2.63 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -3.25 DPS) [dungeon] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Scout's Medallion; shoulder: Blood Guard's Chain Epaulets; back: Serpentskin Cloak; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; waist: Defiler's Chain Girdle; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Ankh of Life; main_hand: Dawn's Edge; off_hand: Inventor's Focal Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1485, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (troll, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 211.4. Weights run: 1.7s. Verify run: 2.2s. 2076 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.339 ± 0.094, crit=20.135 ± 0.769, hit=not significant (0.000 ± 0.000), melee_haste=10.525 ± 1.657

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Headpiece (239540) | Leonid Barthalomew the Revered [vendor] | 678.4 | yes | Cryptstalker Headpiece (22438, -1.12 DPS, sim-verified) [quest]; Champion's Chain Greathelm (227080, -3.68 DPS) [vendor]; Lieutenant Commander's Chain Greathelm (227086, -3.68 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (196.8 DPS) | yes | Onyxia Tooth Pendant (18404, -0.43 DPS) [quest]; Blazefury Medallion (17111, -0.97 DPS, sim-verified) [world]; Amulet of the Darkmoon (19491, -14.84 DPS) [quest] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | sim-verified (200.0 DPS) | yes | Cryptstalker Spaulders (22439, -2.93 DPS) [quest]; Darkspear Pauldrons (272105, -3.20 DPS) [vendor]; Dawnstalker Spaulders (239542, -3.21 DPS, sim-verified) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | sim-verified (200.0 DPS) | yes | Cape of the Black Baron (13340, -0.32 DPS) [dungeon]; Deathguard's Cloak (20068, -0.86 DPS) [rep]; Chromatic Cloak (18509, -3.25 DPS, sim-verified) [crafted] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (199.3 DPS) | yes | Dawnstalker Tunic (239543, -2.50 DPS, sim-verified) [vendor]; Legionnaire's Chain Armor (227083, -6.88 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -6.88 DPS) [vendor] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | 361.4 | yes | Dawnstalker Wristguards (239544, -0.90 DPS, sim-verified) [vendor]; Cryptstalker Wristguards (22443, -17.06 DPS) [quest]; General's Chain Wristguards (16570, -17.86 DPS) [pvp] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | sim-verified (199.2 DPS) | yes | Dawnstalker Handguards (239539, -2.39 DPS, sim-verified) [vendor]; Cryptstalker Handguards (22441, -3.94 DPS) [quest]; Marshal's Chain Grips (16463, -4.34 DPS) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 384.1 | yes | Dawnstalker Girdle (239538, -1.43 DPS, sim-verified) [vendor]; Cryptstalker Girdle (22442, -2.75 DPS) [quest]; Defiler's Chain Girdle (20150, -3.87 DPS) [rep] |
| legs | Dawnstalker Legguards (239541) | Leonid Barthalomew the Revered [vendor] | 673.7 | yes | Sentinel's Chain Leggings (237819, -0.92 DPS, sim-verified) [vendor]; Sentinel's Leather Pants (237818, -2.65 DPS) [vendor]; Legionnaire's Chain Legplates (227079, -3.78 DPS) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | sim-verified (201.1 DPS) | yes | Dawnstalker Boots (239537, -4.35 DPS, sim-verified) [vendor]; Cryptstalker Boots (22440, -19.24 DPS) [quest]; Striker's Footguards (21365, -19.50 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 297.9 | yes | Dragonslayer's Signet (18403, -0.91 DPS) [quest]; Ring of Entropy (18543, -0.91 DPS) [world]; Mindtear Band (20632, -0.91 DPS) [world] |
| finger2 | Band of the Penitent (13217) (or Dragonslayer's Signet (18403), Ring of Entropy (18543), Mindtear Band (20632), Band of Earthen Wrath (21179), Band of Earthen Might (21182), Don Rodrigo's Band (21563), Ritssyn's Ring of Chaos (21836), Ring of the Eternal Flame (23237)) | Houses of the Holy [quest] | 281.9 | yes | Dragonslayer's Signet (18403, +0.00 DPS, sim-verified) [quest]; Ring of Entropy (18543, +0.00 DPS) [world]; Mindtear Band (20632, +0.00 DPS) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (196.5 DPS) | yes | Guardian Talisman (1490, -4.77 DPS) [quest]; Ankh of Life (1713, -4.77 DPS) [dungeon]; Blazing Emblem (2802, -4.77 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (196.8 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -0.07 DPS, sim-verified) [vendor] |
| main_hand | High Warlord's Hacker (235476) | Sergeant Thunderhorn [vendor] | sim-verified (196.8 DPS) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Glacial Blade (19099, -1.79 DPS, sim-verified) [rep] |
| off_hand | Grand Marshal's Hacker (235481) | Captain O'Neal [vendor] | 1401.2 | yes | High Warlord's Shiv (235478, +0.00 DPS, sim-verified) [vendor]; High Warlord's Left Claw (234558, -0.22 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, -0.22 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (196.8 DPS) | yes | High Warlord's Recurve (234559, -4.60 DPS) [vendor]; High Warlord's Crossbow (234560, -4.60 DPS) [vendor]; Dark Iron Rifle (16004, -9.41 DPS, sim-verified) [crafted] |

**New at 60:** head: Dawnstalker Headpiece; neck: Medallion of the Dawn; shoulder: Dawnstalker Pauldrons; back: Cloak of the Fallen God; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Legguards; feet: Dawnstalker Greaves; finger1: Don Julio's Band; finger2: Band of the Penitent; trinket2: Serenity Field; main_hand: High Warlord's Hacker; off_hand: Grand Marshal's Hacker; ranged: The Purifier

No-known-source sample (15 of 2076, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

