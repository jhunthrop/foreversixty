# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (night-elf, 0000000000000000-5420000000000000000-0000000000000000)

Set DPS (verified): 61.9. Weights run: 1.2s. Verify run: 0.9s. 358 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.436 ± 0.049, crit=8.959 ± 0.224, hit=not significant (0.000 ± 0.000), melee_haste=5.281 ± 0.416

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 18.6 | yes | Brawler's Leather Hood (252504, -0.36 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.03 DPS) [crafted]; Shadow Goggles (4373, -1.03 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.6 | yes | Tarnished Locket (279870, -0.53 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.2 | yes | Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.45 DPS, sim-verified) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 9.8 | yes | Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS, sim-verified) [dungeon]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 21.7 | yes | Defender's Leather Armor (252434, -0.10 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.30 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.32 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.4 | yes | Forest Leather Bracers (3202, -0.18 DPS, sim-verified) [dungeon]; Wolf Bracers (4794, -0.26 DPS) [vendor]; Ratchet Wristwraps (274742, -0.34 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 125.4 | yes | Gloves of the Fang (10413, +0.45 DPS, sim-verified) [dungeon]; Gold-flecked Gloves (5195, -6.05 DPS) [dungeon]; Brawler's Leather Gloves (252494, -6.12 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Brawler's Leather Belt (252428, -0.15 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Ruffian Belt (5975, -0.23 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 26.8 | yes | Trapper's Leather Pants (252501, +0.17 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 18.8 | yes | Feet of the Lynx (1121, -0.01 DPS, sim-verified) [dungeon]; Defender's Leather Boots (252441, -0.40 DPS) [crafted]; Totemic Leather Boots (252442, -0.40 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 15.0 | yes | The 1 Ring (8350, -0.62 DPS) [world]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Minor Channeling Ring (1449, -0.83 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 8.6 | yes | The 1 Ring (8350, -0.28 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Minor Channeling Ring (1449, -0.48 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 309.4 | yes | Smite's Mighty Hammer (7230, +0.49 DPS, sim-verified) [dungeon]; Living Root (6631, -0.69 DPS) [dungeon]; Staff of Westfall (2042, -1.24 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Signet of the Zhevra; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Hammerbone

No-known-source sample (15 of 358, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers

### Band 30 (night-elf, 0000000000000000-5423222100000000000-0000000000000000)

Set DPS (verified): 101.2. Weights run: 1.1s. Verify run: 1.1s. 679 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.494 ± 0.064, crit=10.211 ± 0.292, hit=not significant (0.000 ± 0.000), melee_haste=6.340 ± 0.693

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 27.8 | yes | Azure Gustwoven Hood (277050, -0.36 DPS) [crafted]; Defender's Leather Hood (252447, -0.48 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.05 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.26 DPS) [rep]; Pendant of Myzrael (4614, -0.73 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 28.0 | yes | Mantle of Thieves (2264, -0.68 DPS) [dungeon]; Barbaric Shoulders (5964, -0.74 DPS, sim-verified) [crafted]; Dark Leather Shoulders (4252, -0.92 DPS) [crafted] |
| back | Sergeant Major's Cape (16315) | Rank 9 [pvp] | 15.3 | yes | Grave Shroud (279865, -0.28 DPS) [quest]; Lambent Scale Cloak (4706, -0.31 DPS) [dungeon]; Wolfmaster Cape (6314, -0.60 DPS, sim-verified) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.9 | yes | Brawler's Leather Armor (252490, -0.20 DPS) [crafted]; Defender's Leather Tunic (252450, -0.25 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 15.3 | yes | Jurassic Wristguards (6198, -0.09 DPS) [world]; Bands of Serra'kis (6902, -0.17 DPS, sim-verified) [dungeon]; Technician's Bracers (270042, -0.19 DPS) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 143.0 | yes | Insignia Gloves (6408, +0.88 DPS, sim-verified) [dungeon]; Toughened Leather Gloves (4253, -6.27 DPS) [crafted]; Wolfclaw Gloves (1978, -6.40 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.8 | yes | Highlander's Chain Girdle (20090, -0.31 DPS) [rep]; Highlander's Leather Girdle (20117, -0.31 DPS) [rep]; Skulker's Leather Belt (252520, -0.31 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.7 | yes | Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.15 DPS) [crafted]; Brawler's Leather Pants (252500, -0.17 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.1 | yes | Feet of the Lynx (1121, +0.02 DPS, sim-verified) [dungeon]; Insignia Boots (4055, -0.37 DPS) [dungeon]; Highlander's Mail Greaves (20123, -0.37 DPS) [vendor] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 22.9 | yes | Protector's Band (20439, -0.40 DPS) [rep]; Silverlaine's Family Seal (6321, -0.59 DPS) [dungeon]; Seal of Wrynn (2933, -0.60 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.7 | yes | Seal of Wrynn (2933, -0.59 DPS) [quest]; Monkey Ring (6748, -0.64 DPS) [quest]; Silverlaine's Family Seal (6321, -0.90 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan [dungeon] | 450.6 | yes | Gnarled Ash Staff (791, -4.80 DPS) [dungeon]; Slaghammer (1976, -4.84 DPS) [dungeon]; Cobalt Crusher (7730, -24.11 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Brawler's Leather Tunic; wrist: Barbaric Bracers; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Protector's Band; finger2: Ironspine's Eye; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 679, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (night-elf, 0000000000000000-5423222121032010001-0000000000000000)

Set DPS (verified): 103.5. Weights run: 1.1s. Verify run: 1.1s. 931 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.534 ± 0.071, crit=11.632 ± 0.333, hit=not significant (0.000 ± 0.000), melee_haste=7.650 ± 0.975

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.4 | yes | Hawkeye's Helm (14591, -0.90 DPS) [world]; Defender's Leather Helm (252455, -0.93 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -1.19 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 16.9 | yes | Ghostshard Talisman (7731, -0.24 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19541, -0.25 DPS) [rep]; Sentinel's Medallion (20444, -0.42 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 | yes | Forest Tracker Epaulets (2278, -0.02 DPS, sim-verified) [dungeon]; Imperial Leather Spaulders (4737, -0.44 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.45 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 23.1 | yes | Sergeant Major's Cape (16315, -0.43 DPS) [pvp]; Yeti Fur Cloak (2805, -0.56 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.60 DPS) [dungeon] |
| chest | Barbaric Harness (5739) | Leatherworking [crafted] | 28.5 | yes | Brawler's Leather Tunic (252508, -0.14 DPS, sim-verified) [crafted]; Defender's Leather Tunic (252450, -0.30 DPS) [crafted]; Nightscape Tunic (8175, -0.30 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Barbaric Bracers (18948, -0.24 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.34 DPS) [dungeon]; Jurassic Wristguards (6198, -0.34 DPS) [world] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 182.9 | yes | Shadowskin Gloves (18238, -1.11 DPS) [crafted]; Fletcher's Gloves (7348, -1.19 DPS, sim-verified) [crafted]; Prowler's Leather Gloves (252524, -8.08 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 30.1 | yes | Skulker's Leather Belt (252520, -0.13 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.14 DPS, sim-verified) [rep]; Barbaric Belt (4264, -0.25 DPS) [crafted] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 34.6 | yes | Basilisk Hide Pants (1718, -0.12 DPS, sim-verified) [dungeon]; Brawler's Leather Legguards (252516, -0.25 DPS) [crafted]; Brawler's Leather Pants (252500, -0.38 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 36.3 | yes | Skulker's Leather Shoes (252531, -0.27 DPS, sim-verified) [crafted]; Excelsior Boots (4109, -0.34 DPS) [quest]; Imperial Leather Boots (6431, -0.43 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 30.8 | yes | Protector's Band (19517, -0.47 DPS, sim-verified) [rep]; Ring of the Underwood (2951, -0.60 DPS) [dungeon]; Suspicious Spare Part (274754, -0.81 DPS) [vendor] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.1 | yes | Ring of the Underwood (2951, -0.17 DPS, sim-verified) [dungeon]; Suspicious Spare Part (274754, -0.38 DPS) [vendor]; Disengagement Ring (276202, -0.60 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Primitive Fishing Pole (276203) | Friz Frazzlespark [vendor] | 585.3 | yes | Thornstone Sledgehammer (1722, -1.01 DPS) [dungeon]; Illusionary Rod (7713, -1.04 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -1.08 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Barbaric Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Protector's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Primitive Fishing Pole

No-known-source sample (15 of 931, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (night-elf, 0000000000000000-5423222121032010001-5500000000000000)

Set DPS (verified): 129.6. Weights run: 1.1s. Verify run: 1.2s. 1260 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.597 ± 0.077, crit=12.908 ± 0.369, hit=not significant (0.000 ± 0.000), melee_haste=7.662 ± 1.230

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 43.1 | yes | Defender's Leather Helm (252455, -0.84 DPS) [crafted]; Helm of Fire (8348, -0.88 DPS) [crafted]; Eye of Theradras (17715, -1.50 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 19.2 | yes | Sentinel's Medallion (19540, -0.10 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.29 DPS) [dungeon]; Sentinel's Medallion (19541, -0.35 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 39.9 | yes | Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted]; Failed Flying Experiment (9647, -0.16 DPS, sim-verified) [quest]; Warden's Leather Shoulder (252536, -0.43 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 23.5 | yes | Bloodlust Cape (14801, -0.15 DPS) [dungeon]; Pridelord Cape (14673, -0.23 DPS, sim-verified) [dungeon]; Serpentskin Cloak (8259, -0.33 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 54.3 | yes | Blazewind Breastplate (11193, -0.61 DPS, sim-verified) [quest]; Barbaric Harness (5739, -1.40 DPS) [crafted]; Charred Leather Tunic (19127, -1.41 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 33.2 | yes | Prowler's Leather Bracers (252539, +0.00 DPS, sim-verified) [crafted]; Skulker's Leather Bracers (252540, -0.15 DPS) [crafted]; Pridelord Bands (14672, -0.31 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 200.7 | yes | Shadowskin Gloves (18238, -1.11 DPS) [crafted]; Fletcher's Gloves (7348, -1.21 DPS, sim-verified) [crafted]; Serpentskin Gloves (8260, -8.74 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 200.7 | yes | Highlander's Cloth Girdle (20097, -1.10 DPS, sim-verified) [rep]; Highlander's Lizardhide Girdle (20103, -1.11 DPS) [rep]; Vinehedge Cinch (9657, -9.32 DPS) [quest] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 58.8 | yes | Serpentskin Leggings (8262, -0.17 DPS) [dungeon]; Triprunner Dungarees (9624, -1.28 DPS) [quest]; Stormshroud Pants (15057, -2.00 DPS, sim-verified) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.7 | yes | Skulker's Leather Boots (252469, -0.08 DPS, sim-verified) [crafted]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.61 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 37.6 | yes | Protector's Band (19515, -0.38 DPS, sim-verified) [rep]; Protector's Band (19517, -0.78 DPS) [rep]; Masons Fraternity Ring (9533, -0.84 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.7 | yes | Masons Fraternity Ring (9533, -0.05 DPS, sim-verified) [quest]; Ring of the Underwood (2951, -0.17 DPS) [dungeon]; Blackstone Ring (17713, -0.20 DPS) [dungeon] |
| trinket1 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | 606.6 | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Kindling Stave (11750, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Sentinel's Medallion; shoulder: Prowler's Leather Shoulder; chest: Warbear Harness; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; trinket1: Smoking Heart of the Mountain; trinket2: Talisman of Arathor; main_hand: Blight

No-known-source sample (15 of 1260, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 60 (night-elf, 0000000000000000-5423222121032010001-5553200000000000)

Set DPS (verified): 174.4. Weights run: 1.2s. Verify run: 1.2s. 2038 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.760 ± 0.098, crit=13.720 ± 0.409, hit=not significant (0.000 ± 0.000), melee_haste=not significant (7.156 ± 1.916)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ragefury Eyepatch (11735) | Blackrock Depths: Guzzler [dungeon] | 398.1 | yes | Bloodvine Lens (19998, -0.88 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Helmet (16451, -6.84 DPS) [vendor]; Warlord's Dragonhide Helmet (16550, -6.84 DPS) [vendor] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 22.9 | yes | Onyxia Tooth Pendant (18404, +0.00 DPS) [quest]; Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -1.98 DPS, sim-verified) [quest] |
| shoulder | Lieutenant Commander's Dragonhide Shoulders (227172) | Rank 14 [pvp] | 230.5 | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS, sim-verified) [pvp]; Stormshroud Shoulders (15058, -2.08 DPS) [crafted]; Shroud of the Nathrezim (18720, -2.08 DPS) [dungeon] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 71.3 | yes | Cape of the Black Baron (13340, -1.35 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.55 DPS) [rep]; Chromatic Cloak (18509, -3.58 DPS, sim-verified) [crafted] |
| chest | Field Marshal's Dragonhide Breastplate (16452) | Captain Dirgehammer [vendor] | 274.7 | yes | Warlord's Dragonhide Hauberk (16549, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Chestpiece (227176, -1.38 DPS) [pvp]; Stormshroud Armor (15056, -2.69 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 59.0 | yes | Bracers of Subterfuge (22668, -0.51 DPS) [quest]; Forest Stalker's Bracers (19589, -0.51 DPS, sim-verified) [rep]; Forest Stalker's Bracers (19590, -0.85 DPS) [rep] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 220.1 | yes | Gloves of Holy Might (867, -0.51 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -1.52 DPS) [crafted]; Shadowskin Gloves (18238, -1.52 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20045) | The League of Arathor [rep] | 226.1 | yes | Highlander's Leather Girdle (20115, -0.89 DPS, sim-verified) [rep]; Belt of the Archmage (18405, -1.84 DPS) [crafted]; Highlander's Lizardhide Girdle (20046, -1.84 DPS) [rep] |
| legs | Marshal's Dragonhide Legguards (16450) | Captain Dirgehammer [vendor] | 279.4 | yes | General's Dragonhide Leggings (16552, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, -0.79 DPS) [pvp]; Stormshroud Pants (15057, -3.02 DPS, sim-verified) [crafted] |
| feet | Blood Guard's Dragonhide Treads (227181) (or Knight-Lieutenant's Dragonhide Treads (227182)) | Rank 11 [pvp] | 233.4 | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS, sim-verified) [pvp]; Marshal's Dragonhide Boots (16459, -9.35 DPS) [vendor]; General's Dragonhide Boots (16554, -9.35 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 208.1 | yes | Band of the Penitent (13217, -0.87 DPS) [quest]; Dragonslayer's Signet (18403, -0.87 DPS) [quest]; Ring of Entropy (18543, -0.87 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 206.0 | yes | Dragonslayer's Signet (18403, -0.76 DPS) [quest]; Ring of Entropy (18543, -0.76 DPS) [world]; Band of the Penitent (13217, -2.14 DPS, sim-verified) [quest] |
| trinket1 | Onyxia Blood Talisman (18406) | For All To See [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | High Warlord's Pulverizer (234545) | Rank 18 [pvp] | 1378.4 | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Ragefury Eyepatch; neck: Blazefury Medallion; shoulder: Lieutenant Commander's Dragonhide Shoulders; back: Cloak of the Fallen God; chest: Field Marshal's Dragonhide Breastplate; wrist: Forest Stalker's Bracers; hands: Devilsaur Gauntlets; waist: Highlander's Leather Girdle; legs: Marshal's Dragonhide Legguards; feet: Blood Guard's Dragonhide Treads; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Onyxia Blood Talisman; trinket2: Talisman of Arathor; main_hand: High Warlord's Pulverizer

No-known-source sample (15 of 2038, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

## Horde

### Band 20 (tauren, 0000000000000000-5420000000000000000-0000000000000000)

Set DPS (verified): 61.2. Weights run: 1.2s. Verify run: 0.9s. 354 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.436 ± 0.049, crit=8.959 ± 0.224, hit=not significant (0.000 ± 0.000), melee_haste=5.281 ± 0.416

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 18.6 | yes | Brawler's Leather Hood (252504, -0.37 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.03 DPS) [crafted]; Shadow Goggles (4373, -1.03 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.6 | yes | Tarnished Locket (279870, -0.53 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.2 | yes | Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.44 DPS, sim-verified) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 9.8 | yes | Lambent Scale Cloak (4706, -0.04 DPS, sim-verified) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 21.7 | yes | Defender's Leather Armor (252434, -0.09 DPS, sim-verified) [crafted]; Totemic Leather Armor (252435, -0.30 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.32 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.4 | yes | Forest Leather Bracers (3202, -0.17 DPS, sim-verified) [dungeon]; Wolf Bracers (4794, -0.26 DPS) [vendor]; Ratchet Wristwraps (274742, -0.34 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 125.4 | yes | Gloves of the Fang (10413, +0.40 DPS, sim-verified) [dungeon]; Gold-flecked Gloves (5195, -6.05 DPS) [dungeon]; Brawler's Leather Gloves (252494, -6.12 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 | yes | Brawler's Leather Belt (252428, -0.15 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Ruffian Belt (5975, -0.23 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 26.8 | yes | Trapper's Leather Pants (252501, +0.14 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 18.8 | yes | Feet of the Lynx (1121, -0.01 DPS, sim-verified) [dungeon]; Defender's Leather Boots (252441, -0.40 DPS) [crafted]; Totemic Leather Boots (252442, -0.40 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 15.0 | yes | Bounty Hunter's Ring (5351, -0.59 DPS) [quest]; The 1 Ring (8350, -0.62 DPS) [world]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 8.6 | yes | Bounty Hunter's Ring (5351, -0.26 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.27 DPS) [world]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 309.4 | yes | Smite's Mighty Hammer (7230, +0.47 DPS, sim-verified) [dungeon]; Living Root (6631, -0.69 DPS) [dungeon]; Crescent Staff (6505, -0.81 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Hammerbone

No-known-source sample (15 of 354, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers

### Band 30 (tauren, 0000000000000000-5423222100000000000-0000000000000000)

Set DPS (verified): 99.8. Weights run: 1.1s. Verify run: 1.1s. 676 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.494 ± 0.064, crit=10.211 ± 0.292, hit=not significant (0.000 ± 0.000), melee_haste=6.340 ± 0.693

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 27.8 | yes | Azure Gustwoven Hood (277050, -0.36 DPS) [crafted]; Defender's Leather Hood (252447, -0.48 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.49 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.03 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.26 DPS) [rep]; Pendant of Myzrael (4614, -0.73 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | Gnomeregan: Mechanized Sentry [dungeon] | 28.0 | yes | Mantle of Thieves (2264, -0.68 DPS) [dungeon]; Barbaric Shoulders (5964, -0.71 DPS, sim-verified) [crafted]; Dark Leather Shoulders (4252, -0.92 DPS) [crafted] |
| back | Sergeant Major's Cape (16315) | Rank 9 [pvp] | 15.3 | yes | Wildhunter Cloak (16658, -0.27 DPS) [quest]; Grave Shroud (279865, -0.28 DPS) [quest]; Wolfmaster Cape (6314, -0.58 DPS, sim-verified) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.9 | yes | Brawler's Leather Armor (252490, -0.20 DPS) [crafted]; Defender's Leather Tunic (252450, -0.24 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 15.3 | yes | Jurassic Wristguards (6198, -0.09 DPS) [world]; Bands of Serra'kis (6902, -0.15 DPS, sim-verified) [dungeon]; Technician's Bracers (270042, -0.19 DPS) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 143.0 | yes | Insignia Gloves (6408, +0.86 DPS, sim-verified) [dungeon]; Toughened Leather Gloves (4253, -6.27 DPS) [crafted]; Wolfclaw Gloves (1978, -6.40 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.8 | yes | Skulker's Leather Belt (252520, -0.29 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.31 DPS) [rep]; Defiler's Leather Girdle (20191, -0.31 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.7 | yes | Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.15 DPS) [crafted]; Brawler's Leather Pants (252500, -0.16 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.1 | yes | Feet of the Lynx (1121, +0.06 DPS, sim-verified) [dungeon]; Stomping Boots (3741, -0.20 DPS) [quest]; Insignia Boots (4055, -0.37 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 22.9 | yes | Band of the Fist (17694, -0.40 DPS) [quest]; Legionnaire's Band (20429, -0.40 DPS) [rep]; Silverlaine's Family Seal (6321, -0.59 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.7 | yes | Band of the Fist (17694, -0.58 DPS, sim-verified) [quest]; Silverlaine's Family Seal (6321, -0.58 DPS) [dungeon]; Monkey Ring (6748, -0.64 DPS) [quest] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan [dungeon] | 450.6 | yes | Gnarled Ash Staff (791, -4.80 DPS) [dungeon]; Slaghammer (1976, -4.84 DPS) [dungeon]; Cobalt Crusher (7730, -23.59 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Brawler's Leather Tunic; wrist: Barbaric Bracers; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Legionnaire's Band; finger2: Ironspine's Eye; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 676, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (tauren, 0000000000000000-5423222121032010001-0000000000000000)

Set DPS (verified): 101.4. Weights run: 1.1s. Verify run: 1.1s. 928 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.534 ± 0.071, crit=11.632 ± 0.333, hit=not significant (0.000 ± 0.000), melee_haste=7.650 ± 0.975

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.4 | yes | Hawkeye's Helm (14591, -0.90 DPS) [world]; Defender's Leather Helm (252455, -0.97 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -1.19 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 17.7 | yes | Scout's Medallion (19536, +0.06 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.21 DPS) [dungeon]; Scout's Medallion (19537, -0.30 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 | yes | Forest Tracker Epaulets (2278, -0.02 DPS, sim-verified) [dungeon]; Imperial Leather Spaulders (4737, -0.44 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.45 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 23.1 | yes | Sergeant Major's Cape (16315, -0.51 DPS, sim-verified) [pvp]; Imperial Cloak (6432, -0.60 DPS) [dungeon]; Parachute Cloak (10518, -0.60 DPS) [crafted] |
| chest | Barbaric Harness (5739) | Leatherworking [crafted] | 28.5 | yes | Brawler's Leather Tunic (252508, -0.13 DPS, sim-verified) [crafted]; Defender's Leather Tunic (252450, -0.30 DPS) [crafted]; Nightscape Tunic (8175, -0.30 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Barbaric Bracers (18948, -0.19 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.34 DPS) [dungeon]; Jurassic Wristguards (6198, -0.34 DPS) [world] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 182.9 | yes | Shadowskin Gloves (18238, -1.11 DPS) [crafted]; Fletcher's Gloves (7348, -1.16 DPS, sim-verified) [crafted]; Prowler's Leather Gloves (252524, -8.08 DPS) [crafted] |
| waist | Tharg's Shoelace (9705) | Army of the Black Dragon [quest] | 30.2 | yes | Prowler's Leather Belt (252459, +0.17 DPS, sim-verified) [crafted]; Defiler's Leather Girdle (20192, -0.01 DPS) [rep]; Skulker's Leather Belt (252520, -0.13 DPS) [crafted] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 34.6 | yes | Basilisk Hide Pants (1718, -0.09 DPS, sim-verified) [dungeon]; Brawler's Leather Legguards (252516, -0.25 DPS) [crafted]; Trapper's Leather Pants (252501, -0.38 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 36.3 | yes | Skulker's Leather Shoes (252531, -0.27 DPS, sim-verified) [crafted]; Excelsior Boots (4109, -0.34 DPS) [quest]; Imperial Leather Boots (6431, -0.43 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 30.8 | yes | Legionnaire's Band (19513, -0.46 DPS, sim-verified) [rep]; Ring of the Underwood (2951, -0.60 DPS) [dungeon]; Suspicious Spare Part (274754, -0.81 DPS) [vendor] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.1 | yes | Ring of the Underwood (2951, -0.17 DPS, sim-verified) [dungeon]; Suspicious Spare Part (274754, -0.38 DPS) [vendor]; Band of the Fist (17694, -0.42 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 648.8 | yes | Primitive Fishing Pole (276203, +0.94 DPS, sim-verified) [vendor]; Thornstone Sledgehammer (1722, -4.53 DPS) [dungeon]; Staff of Jordan (873, -4.60 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Barbaric Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Tharg's Shoelace; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod

No-known-source sample (15 of 928, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (tauren, 0000000000000000-5423222121032010001-5500000000000000)

Set DPS (verified): 132.1. Weights run: 1.1s. Verify run: 1.2s. 1257 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.597 ± 0.077, crit=12.908 ± 0.369, hit=not significant (0.000 ± 0.000), melee_haste=7.662 ± 1.230

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 43.1 | yes | Undercity Reservist's Cap (20643, -0.13 DPS) [quest]; Defender's Leather Helm (252455, -0.84 DPS) [crafted]; Eye of Theradras (17715, -1.43 DPS, sim-verified) [dungeon] |
| neck | Woven Ivy Necklace (19159) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 28.3 | yes | Scout's Medallion (19535, -0.52 DPS, sim-verified) [rep]; Ethereal Talisman (4430, -0.57 DPS) [quest]; Scout's Medallion (19536, -0.59 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 39.9 | yes | Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted]; Failed Flying Experiment (9647, -0.15 DPS, sim-verified) [quest]; Warden's Leather Shoulder (252536, -0.43 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 23.5 | yes | Bloodlust Cape (14801, -0.15 DPS) [dungeon]; Pridelord Cape (14673, -0.23 DPS, sim-verified) [dungeon]; Serpentskin Cloak (8259, -0.33 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 54.3 | yes | Blazewind Breastplate (11193, -0.57 DPS, sim-verified) [quest]; Barbaric Harness (5739, -1.40 DPS) [crafted]; Charred Leather Tunic (19127, -1.41 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | Azuregos [world] | 33.2 | yes | Prowler's Leather Bracers (252539, -0.01 DPS, sim-verified) [crafted]; Skulker's Leather Bracers (252540, -0.15 DPS) [crafted]; Pridelord Bands (14672, -0.31 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | Gnomeregan: Dark Iron Ambassador [dungeon] | 200.7 | yes | Shadowskin Gloves (18238, -1.11 DPS) [crafted]; Fletcher's Gloves (7348, -1.20 DPS, sim-verified) [crafted]; Serpentskin Gloves (8260, -8.74 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 200.7 | yes | Defiler's Cloth Girdle (20165, -1.09 DPS, sim-verified) [rep]; Defiler's Lizardhide Girdle (20174, -1.11 DPS) [rep]; Prowler's Leather Belt (252459, -9.43 DPS) [crafted] |
| legs | Serpentskin Leggings (8262) | Maraudon: Princess Theradras [dungeon] | 55.7 | yes | Triprunner Dungarees (9624, -1.11 DPS) [quest]; Dragonflight Leggings (10742, -1.17 DPS) [quest]; Stormshroud Pants (15057, -1.74 DPS, sim-verified) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.7 | yes | Skulker's Leather Boots (252469, -0.05 DPS, sim-verified) [crafted]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.61 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 37.6 | yes | Legionnaire's Band (19512, -0.37 DPS, sim-verified) [rep]; Ironspine's Eye (7686, -0.77 DPS) [dungeon]; Legionnaire's Band (19513, -0.78 DPS) [rep] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Ironspine's Eye (7686, +0.03 DPS, sim-verified) [dungeon]; Masons Fraternity Ring (9533, -0.09 DPS) [quest]; Band of Allegiance (18585, -0.17 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 42.0 | yes | Tidal Charm (1404, -2.33 DPS) [vendor]; Guardian Talisman (1490, -2.33 DPS) [quest]; Ankh of Life (1713, -2.33 DPS) [dungeon] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | 606.6 | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Kindling Stave (11750, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Woven Ivy Necklace; shoulder: Prowler's Leather Shoulder; chest: Warbear Harness; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Smoking Heart of the Mountain; main_hand: Blight

No-known-source sample (15 of 1257, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 60 (tauren, 0000000000000000-5423222121032010001-5553200000000000)

Set DPS (verified): 178.0. Weights run: 1.2s. Verify run: 1.2s. 2034 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.760 ± 0.098, crit=13.720 ± 0.409, hit=not significant (0.000 ± 0.000), melee_haste=not significant (7.156 ± 1.916)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ragefury Eyepatch (11735) | Blackrock Depths: Guzzler [dungeon] | 398.1 | yes | Bloodvine Lens (19998, -0.89 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Helmet (16451, -6.84 DPS) [vendor]; Warlord's Dragonhide Helmet (16550, -6.84 DPS) [vendor] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 22.9 | yes | Onyxia Tooth Pendant (18404, +0.00 DPS) [quest]; Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -2.46 DPS, sim-verified) [quest] |
| shoulder | Lieutenant Commander's Dragonhide Shoulders (227172) (or Champion's Dragonhide Shoulders (227175)) | Rank 14 [pvp] | 230.5 | yes | Champion's Dragonhide Shoulders (227175, +0.00 DPS, sim-verified) [pvp]; Stormshroud Shoulders (15058, -2.08 DPS) [crafted]; Shroud of the Nathrezim (18720, -2.08 DPS) [dungeon] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 71.3 | yes | Cape of the Black Baron (13340, -1.35 DPS) [dungeon]; Deathguard's Cloak (20068, -1.55 DPS) [rep]; Chromatic Cloak (18509, -3.55 DPS, sim-verified) [crafted] |
| chest | Field Marshal's Dragonhide Breastplate (16452) | Captain Dirgehammer [vendor] | 274.7 | yes | Warlord's Dragonhide Hauberk (16549, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Chestpiece (227176, -1.38 DPS) [pvp]; Stormshroud Armor (15056, -2.73 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 59.0 | yes | Forest Stalker's Bracers (19589, -0.50 DPS, sim-verified) [rep]; Bracers of Subterfuge (22668, -0.51 DPS) [quest]; Forest Stalker's Bracers (19590, -0.85 DPS) [rep] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 220.1 | yes | Gloves of Holy Might (867, -0.51 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -1.52 DPS) [crafted]; Shadowskin Gloves (18238, -1.52 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20190) | The Defilers [rep] | 226.1 | yes | Defiler's Leather Girdle (20193, -0.89 DPS, sim-verified) [rep]; Belt of the Archmage (18405, -1.84 DPS) [crafted]; Defiler's Cloth Girdle (20163, -1.84 DPS) [rep] |
| legs | Marshal's Dragonhide Legguards (16450) | Captain Dirgehammer [vendor] | 279.4 | yes | General's Dragonhide Leggings (16552, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, -0.79 DPS) [pvp]; Stormshroud Pants (15057, -3.08 DPS, sim-verified) [crafted] |
| feet | Blood Guard's Dragonhide Treads (227181) (or Knight-Lieutenant's Dragonhide Treads (227182)) | Rank 11 [pvp] | 233.4 | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS, sim-verified) [pvp]; Marshal's Dragonhide Boots (16459, -9.35 DPS) [vendor]; General's Dragonhide Boots (16554, -9.35 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 208.1 | yes | Band of the Penitent (13217, -0.87 DPS) [quest]; Dragonslayer's Signet (18403, -0.87 DPS) [quest]; Ring of Entropy (18543, -0.87 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 206.0 | yes | Dragonslayer's Signet (18403, -0.76 DPS) [quest]; Ring of Entropy (18543, -0.76 DPS) [world]; Band of the Penitent (13217, -2.19 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 42.0 | yes | Tidal Charm (1404, -2.28 DPS) [vendor]; Guardian Talisman (1490, -2.28 DPS) [quest]; Ankh of Life (1713, -2.28 DPS) [dungeon] |
| trinket2 | Onyxia Blood Talisman (18406) | For All To See [quest] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon] |
| main_hand | High Warlord's Pulverizer (234545) | Rank 18 [pvp] | 1378.4 | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Ragefury Eyepatch; neck: Blazefury Medallion; shoulder: Lieutenant Commander's Dragonhide Shoulders; back: Cloak of the Fallen God; chest: Field Marshal's Dragonhide Breastplate; wrist: Forest Stalker's Bracers; hands: Devilsaur Gauntlets; waist: Defiler's Leather Girdle; legs: Marshal's Dragonhide Legguards; feet: Blood Guard's Dragonhide Treads; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket2: Onyxia Blood Talisman; main_hand: High Warlord's Pulverizer

No-known-source sample (15 of 2034, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

