# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.3. Weights run: 1.1s. Verify run: 1.3s. 159 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.404 ± 0.062, hit=not significant (0.000 ± 0.000), melee_haste=not significant (0.665 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.38 DPS) [crafted]; Lucky Fishing Hat (19972, -0.38 DPS) [quest]; Flying Tiger Goggles (4368, -0.54 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 | yes | Erudite's Amulet (277204, -0.14 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.29 DPS) [quest]; Tarnished Locket (279870, -0.29 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Forest Leather Mantle (4709, -0.24 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.34 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.02 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.1 | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.19 DPS) [crafted]; Dark Leather Tunic (2317, -0.24 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.07 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.02 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.33 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Pyrewood Signet Ring (277210, -0.10 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Protector's Band (20439) (or Pyrewood Signet Ring (277210)) | Silverwing Sentinels [rep] | 4.1 | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The 1 Ring (8350, -0.14 DPS) [world]; Pyrewood Signet Ring (277210, -0.21 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Barrens Basher (274744, -1.21 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop]; Grayson's Torch (1172, -10.85 DPS) [quest]; Pulsating Hydra Heart (5183, -10.85 DPS) [world] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 267.4 | yes | Lil Timmy's Peashooter (13136, +0.00 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.64 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Ranger Bow

No-known-source sample (15 of 159, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 20438 Outrunner's Bow; 20440 Protector's Sword; 20443 Sentinel's Blade; 202256 Privateer's Ornate Pistol

### Band 30 (night-elf, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.7. Weights run: 1.1s. Verify run: 1.5s. 279 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.485 ± 0.060, hit=2.379 ± 0.384, melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.37 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.38 DPS) [rep]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.39 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Tigerstrike Mantle (13108, -0.11 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.19 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.09 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.24 DPS) [quest]; Green Leather Armor (4255, -0.38 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.11 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.19 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Fletcher's Gloves (7348, -0.44 DPS) [crafted]; Wolfclaw Gloves (1978, -0.48 DPS) [dungeon]; Pilferer's Gloves (7358, -0.50 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.72 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.73 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.04 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 13.5 | yes | Insurgent's Band (272067, -0.22 DPS) [vendor]; Monkey Ring (6748, -0.31 DPS) [quest]; Ring of Precision (1491, -0.36 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Insurgent's Band (272067, -0.29 DPS, sim-verified) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Zealot Blade (13033, -0.58 DPS) [world_drop]; Electrocutioner Leg (9446, -0.63 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (46.7 DPS) | yes | Ironspine's Fist (7687, -0.54 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.23 DPS) [world_drop]; Totem of Infliction (1131, -15.28 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 384.9 | yes | Silver Star (3463, -0.25 DPS, sim-verified) [quest]; Ironweaver (13137, -1.16 DPS) [world_drop]; Nightstalker Bow (6696, -2.15 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Pyrewood Signet Ring; finger2: Ironspine's Eye; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 279, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 90.8. Weights run: 1.2s. Verify run: 1.5s. 402 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.925 ± 0.148, hit=not significant (3.384 ± 0.980), melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.3 | yes | White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop]; Nightscape Headband (8176, -0.12 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.29 DPS) [rep]; Sentinel's Medallion (20444, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 | yes | Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.00 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 17.4 | yes | Nightscape Tunic (8175, -0.10 DPS) [crafted]; Dusky Leather Armor (7374, -0.15 DPS) [crafted]; Raptorbane Armor (3566, -0.70 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.82 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.0 | yes | Fletcher's Gloves (7348, -0.99 DPS) [crafted]; Shadowskin Gloves (18238, -0.99 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.42 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.49 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Petrolspill Leggings (9509, -0.58 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.56 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Insurgent's Band (272066, -0.40 DPS) [vendor]; Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Falcon's Hook (7552, -0.53 DPS) [world_drop] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 17.6 | yes | Insurgent's Band (272066, +0.00 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Falcon's Hook (7552, -0.42 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Southsea Lamp (9359, -1.23 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.35 DPS) [vendor]; Heaven's Light (13026, -1.53 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Southsea Lamp (9359, -4.45 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -21.88 DPS) [world_drop]; Totem of Infliction (1131, -21.93 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 463.9 | yes | Swiftwind (13038, -0.27 DPS, sim-verified) [world_drop]; Skystriker Bow (13020, -2.05 DPS) [world_drop]; Crusader Bow (15287, -2.10 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Pyrewood Signet Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 7948 Girdle of Thero-shan

### Band 50 (night-elf, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 110.0. Weights run: 1.2s. Verify run: 1.4s. 513 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=6.578 ± 0.398, hit=not significant (5.130 ± 1.364), melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 159.4 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.60 DPS) [dungeon]; Helm of Fire (8348, -7.50 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 14.6 | yes | Sentinel's Medallion (19539, -0.06 DPS) [rep]; Sentinel's Medallion (19540, -0.12 DPS) [rep]; Ghostshard Talisman (7731, -0.76 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 151.4 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -6.79 DPS) [vendor]; Forest Tracker Epaulets (2278, -7.43 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 16.9 | yes | Serpentskin Cloak (8259, -0.24 DPS) [dungeon]; Nightscape Cloak (8195, -0.30 DPS) [crafted]; Blackflame Cape (13109, -0.34 DPS, sim-verified) [world_drop] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 161.4 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -7.25 DPS) [quest]; Warbear Harness (15064, -7.55 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 112.1 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.47 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.07 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 112.1 | yes | Highlander's Lizardhide Girdle (20103, -1.07 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.58 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -4.39 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (110.0 DPS) | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.43 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -7.24 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.24 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.24 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 71.3 | yes | Assault Band (13095, -2.74 DPS) [world_drop]; Masons Fraternity Ring (9533, -2.97 DPS) [quest]; Insurgent's Band (272065, -3.01 DPS) [vendor] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 25.0 | yes | Assault Band (13095, +0.00 DPS, sim-verified) [world_drop]; Masons Fraternity Ring (9533, -0.50 DPS) [quest]; Insurgent's Band (272065, -0.54 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (109.5 DPS) | yes | Thunderbrew's Boot Flask (744, -2.47 DPS) [quest]; Tidal Charm (1404, -2.47 DPS) [vendor]; Guardian Talisman (1490, -2.47 DPS) [quest] |
| trinket2 | - | - |  |  |  |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | sim-verified (109.5 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Might of Hakkar (10838, -3.10 DPS) [world]; Thorium Cestus (250614, -3.15 DPS) [crafted] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Claw of Celebras (17738, -3.80 DPS) [dungeon]; Might of Hakkar (10838, -5.04 DPS, sim-verified) [world]; Thermotastic Egg Timer (9644, -29.39 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (109.5 DPS) | yes | Dark Iron Rifle (16004, -0.47 DPS, sim-verified) [crafted]; Gryphonwing Long Bow (13022, -3.24 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -3.54 DPS) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 513, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (night-elf, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 296.0. Weights run: 1.2s. Verify run: 1.8s. 957 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 271.5 | yes | Bloodvine Lens (19998, -2.21 DPS) [crafted]; Mask of the Unforgiven (13404, -3.66 DPS) [dungeon]; Ragefury Eyepatch (11735, -15.14 DPS, sim-verified) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (267.2 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -2.15 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 167.1 | yes | Champion's Leather Shoulders (23258, +0.00 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, +0.00 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (269.9 DPS) | yes | Earthweave Cloak (21187, -0.38 DPS) [quest]; Arcanoweave Cloak (272411, -1.50 DPS) [vendor]; Chromatic Cloak (18509, -2.93 DPS, sim-verified) [crafted] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | sim-verified (280.1 DPS) | yes | Dawn Armor (252483, -0.90 DPS) [crafted]; Field Marshal's Leather Chestpiece (16453, -1.72 DPS) [vendor]; Stormshroud Armor (15056, -13.16 DPS, sim-verified) [crafted] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 78.9 | yes | Duskwraith Wristguards (239547, -0.82 DPS) [vendor]; Rockfury Bracers (21186, -1.87 DPS) [quest]; Primal Batskin Bracers (19687, -3.07 DPS, sim-verified) [crafted] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 208.0 | yes | Devilsaur Gauntlets (15063, -3.47 DPS) [crafted]; Marshal's Leather Handgrips (16454, -3.48 DPS) [vendor]; Stormshroud Gloves (21278, -8.46 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 205.2 | yes | Highlander's Leather Girdle (20115, -3.75 DPS) [rep]; Assassin's Waistguard (272395, -4.71 DPS) [vendor]; Highlander's Leather Girdle (20045, -6.73 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 272.9 | yes | Marshal's Leather Leggings (16456, -1.72 DPS) [vendor]; General's Leather Legguards (16564, -1.72 DPS) [vendor]; Sentinel's Leather Pants (237818, -9.89 DPS, sim-verified) [vendor] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 208.0 | yes | Duskwraith Treads (239553, -6.58 DPS, sim-verified) [vendor]; Darkmantle Footpads (226831, -7.69 DPS) [vendor]; Fine Dawn Treaders (227815, -7.96 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (264.8 DPS) | yes | Band of the Penitent (13217, -3.21 DPS) [quest]; Ring of Entropy (18543, -3.21 DPS) [world]; Wrath of Cenarius (21190, -6.97 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (264.8 DPS) | yes | Band of the Penitent (13217, -2.35 DPS) [quest]; Ring of Entropy (18543, -2.35 DPS) [world]; Wrath of Cenarius (21190, -5.80 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (264.3 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (264.8 DPS) | yes | Thunderbrew's Boot Flask (744, -2.12 DPS) [quest]; Tidal Charm (1404, -2.12 DPS) [vendor]; Shard of the Fallen Star (21891, -2.58 DPS, sim-verified) [world_drop] |
| main_hand | High Warlord's Bonecracker (235477) | Sergeant Thunderhorn [vendor] | sim-verified (264.8 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor]; Ebon Hand (19170, -8.13 DPS, sim-verified) [crafted] |
| off_hand | High Warlord's Shiv (235478) | Sergeant Thunderhorn [vendor] | sim-verified (279.0 DPS) | yes | High Warlord's Left Claw (234558, -0.05 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, -0.05 DPS) [vendor]; Grand Marshal's Bonecracker (235480, -12.10 DPS, sim-verified) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (267.2 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -2.42 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Blazefury Medallion; back: Howler's Furs; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: High Warlord's Bonecracker; off_hand: High Warlord's Shiv; ranged: The Purifier

No-known-source sample (15 of 957, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.8. Weights run: 1.1s. Verify run: 1.3s. 166 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.404 ± 0.062, hit=not significant (0.000 ± 0.000), melee_haste=not significant (0.665 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.38 DPS) [crafted]; Lucky Fishing Hat (19972, -0.38 DPS) [quest]; Flying Tiger Goggles (4368, -0.51 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.29 DPS) [quest]; Tarnished Locket (279870, -0.29 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Forest Leather Mantle (4709, -0.24 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.32 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Prospector's Chestpiece (14562, -0.05 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.32 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 | yes | Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.02 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.32 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Pyrewood Signet Ring (277210, -0.10 DPS) [quest]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| finger2 | Legionnaire's Band (20429) (or Pyrewood Signet Ring (277210)) | Warsong Outriders [rep] | 4.1 | yes | Bounty Hunter's Ring (5351, -0.05 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.25 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Wingblade (6504, -1.15 DPS) [quest] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop]; Grayson's Torch (1172, -10.85 DPS) [quest]; Nightglow Concoction (3451, -10.85 DPS) [quest] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 267.4 | yes | Lil Timmy's Peashooter (13136, +0.00 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.64 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Ranger Bow

No-known-source sample (15 of 166, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance; 209622 Insignia of the Horde; 241089 Scarlet Dagger

### Band 30 (troll, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.4. Weights run: 1.1s. Verify run: 1.5s. 291 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.485 ± 0.060, hit=2.379 ± 0.384, melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.38 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.38 DPS) [rep]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.39 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 | yes | Panther Armor (6670, -0.20 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.29 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.29 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.13 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.19 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Braced Handguards (6784, -0.43 DPS) [quest]; Fletcher's Gloves (7348, -0.44 DPS) [crafted]; Pilferer's Gloves (7358, -0.51 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Deftkin Belt (16659, -0.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.76 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.06 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Warsong Boots (16977, -0.19 DPS) [quest] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 13.5 | yes | Insurgent's Band (272067, -0.22 DPS) [vendor]; Monkey Ring (6748, -0.31 DPS) [quest]; Ring of Precision (1491, -0.36 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Insurgent's Band (272067, -0.26 DPS, sim-verified) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Zealot Blade (13033, -0.58 DPS) [world_drop]; Electrocutioner Leg (9446, -0.63 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (46.4 DPS) | yes | Ironspine's Fist (7687, -0.70 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.23 DPS) [world_drop]; Grayson's Torch (1172, -15.28 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 384.9 | yes | Silver Star (3463, -0.25 DPS, sim-verified) [quest]; Ironweaver (13137, -1.16 DPS) [world_drop]; Nightstalker Bow (6696, -2.15 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Pyrewood Signet Ring; finger2: Ironspine's Eye; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 291, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 89.3. Weights run: 1.2s. Verify run: 1.5s. 414 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.925 ± 0.148, hit=not significant (3.384 ± 0.980), melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.3 | yes | White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop]; Nightscape Headband (8176, -0.12 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.29 DPS) [rep]; Scout's Medallion (20442, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 | yes | Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.09 DPS) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 17.4 | yes | Dusky Leather Armor (7374, -0.15 DPS) [crafted]; Nightscape Tunic (8175, -0.24 DPS, sim-verified) [crafted]; Hawkeye's Tunic (14592, -0.25 DPS) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.80 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.0 | yes | Fletcher's Gloves (7348, -0.99 DPS) [crafted]; Shadowskin Gloves (18238, -0.99 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.44 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.48 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Petrolspill Leggings (9509, -0.58 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.56 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Insurgent's Band (272066, -0.40 DPS) [vendor]; Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Falcon's Hook (7552, -0.53 DPS) [world_drop] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 17.6 | yes | Insurgent's Band (272066, -0.04 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Falcon's Hook (7552, -0.42 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Southsea Lamp (9359, -1.23 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.35 DPS) [vendor]; Heaven's Light (13026, -1.53 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Southsea Lamp (9359, -4.35 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -21.88 DPS) [world_drop]; Grayson's Torch (1172, -21.93 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 463.9 | yes | Swiftwind (13038, -0.29 DPS, sim-verified) [world_drop]; Skystriker Bow (13020, -2.05 DPS) [world_drop]; Crusader Bow (15287, -2.10 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Pyrewood Signet Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 111.3. Weights run: 1.2s. Verify run: 1.4s. 525 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=6.578 ± 0.398, hit=not significant (5.130 ± 1.364), melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 159.4 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.60 DPS) [dungeon]; Helm of Fire (8348, -7.50 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 14.6 | yes | Scout's Medallion (19535, -0.06 DPS) [rep]; Scout's Medallion (19536, -0.12 DPS) [rep]; Ghostshard Talisman (7731, -0.75 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 151.4 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -6.79 DPS) [vendor]; Forest Tracker Epaulets (2278, -7.43 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 16.9 | yes | Serpentskin Cloak (8259, -0.24 DPS) [dungeon]; Nightscape Cloak (8195, -0.30 DPS) [crafted]; Blackflame Cape (13109, -0.35 DPS, sim-verified) [world_drop] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 161.4 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -7.25 DPS) [quest]; Warbear Harness (15064, -7.55 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 112.1 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.47 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.07 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 112.1 | yes | Defiler's Lizardhide Girdle (20174, -1.07 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.58 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -4.39 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 184.2 | yes | Knight's Leather Pants (220858, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Leather Pants (220859, -1.22 DPS) [vendor]; Ferine Leggings (6690, -8.45 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.24 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.24 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 71.3 | yes | White Bone Band (11862, -2.53 DPS) [quest]; Assault Band (13095, -2.74 DPS) [world_drop]; Masons Fraternity Ring (9533, -2.97 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 25.0 | yes | White Bone Band (11862, +0.00 DPS, sim-verified) [quest]; Assault Band (13095, -0.27 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.50 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (112.2 DPS) | yes | Tidal Charm (1404, -4.16 DPS) [vendor]; Guardian Talisman (1490, -4.16 DPS) [quest]; Smoking Heart of the Mountain (11811, -4.29 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (112.2 DPS) | yes | Smoking Heart of the Mountain (11811, -1.16 DPS, sim-verified) [crafted]; Tidal Charm (1404, -2.47 DPS) [vendor]; Guardian Talisman (1490, -2.47 DPS) [quest] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | sim-verified (112.2 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Might of Hakkar (10838, -3.10 DPS) [world]; Thorium Cestus (250614, -3.15 DPS) [crafted] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Claw of Celebras (17738, -3.80 DPS) [dungeon]; Might of Hakkar (10838, -5.68 DPS, sim-verified) [world]; White Bone Shredder (11863, -5.90 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (112.2 DPS) | yes | Dark Iron Rifle (16004, -0.47 DPS, sim-verified) [crafted]; Gryphonwing Long Bow (13022, -3.24 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -3.54 DPS) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 297.1. Weights run: 1.2s. Verify run: 1.8s. 970 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 271.5 | yes | Bloodvine Lens (19998, -2.21 DPS) [crafted]; Mask of the Unforgiven (13404, -3.66 DPS) [dungeon]; Ragefury Eyepatch (11735, -13.41 DPS, sim-verified) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (267.7 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -0.23 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 167.1 | yes | Champion's Leather Shoulders (23258, +0.00 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, +0.00 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (270.5 DPS) | yes | Earthweave Cloak (21187, -0.38 DPS) [quest]; Arcanoweave Cloak (272411, -1.50 DPS) [vendor]; Chromatic Cloak (18509, -3.38 DPS, sim-verified) [crafted] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | sim-verified (281.0 DPS) | yes | Dawn Armor (252483, -0.90 DPS) [crafted]; Field Marshal's Leather Chestpiece (16453, -1.72 DPS) [vendor]; Stormshroud Armor (15056, -13.85 DPS, sim-verified) [crafted] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 78.9 | yes | Duskwraith Wristguards (239547, -0.82 DPS) [vendor]; Rockfury Bracers (21186, -1.87 DPS) [quest]; Primal Batskin Bracers (19687, -3.10 DPS, sim-verified) [crafted] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 208.0 | yes | Devilsaur Gauntlets (15063, -3.47 DPS) [crafted]; Marshal's Leather Handgrips (16454, -3.48 DPS) [vendor]; Stormshroud Gloves (21278, -6.66 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 205.2 | yes | Defiler's Leather Girdle (20193, -3.75 DPS) [rep]; Assassin's Waistguard (272395, -4.71 DPS) [vendor]; Defiler's Leather Girdle (20190, -6.53 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 272.9 | yes | Marshal's Leather Leggings (16456, -1.72 DPS) [vendor]; General's Leather Legguards (16564, -1.72 DPS) [vendor]; Sentinel's Leather Pants (237818, -8.13 DPS, sim-verified) [vendor] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 208.0 | yes | Duskwraith Treads (239553, -6.60 DPS, sim-verified) [vendor]; Darkmantle Footpads (226831, -7.69 DPS) [vendor]; Fine Dawn Treaders (227815, -7.96 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (266.8 DPS) | yes | Band of the Penitent (13217, -3.21 DPS) [quest]; Ring of Entropy (18543, -3.21 DPS) [world]; Wrath of Cenarius (21190, -6.80 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (266.8 DPS) | yes | Band of the Penitent (13217, -2.35 DPS) [quest]; Ring of Entropy (18543, -2.35 DPS) [world]; Wrath of Cenarius (21190, -5.64 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (262.4 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (266.8 DPS) | yes | Tidal Charm (1404, -3.90 DPS) [vendor]; Guardian Talisman (1490, -3.90 DPS) [quest]; Frozen Heart of the Mountain (249469, -4.65 DPS, sim-verified) [crafted] |
| main_hand | High Warlord's Bonecracker (235477) | Sergeant Thunderhorn [vendor] | sim-verified (266.8 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor]; Ebon Hand (19170, -6.30 DPS, sim-verified) [crafted] |
| off_hand | High Warlord's Shiv (235478) | Sergeant Thunderhorn [vendor] | sim-verified (279.9 DPS) | yes | High Warlord's Left Claw (234558, -0.05 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, -0.05 DPS) [vendor]; Grand Marshal's Bonecracker (235480, -12.70 DPS, sim-verified) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (267.7 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -2.51 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Blazefury Medallion; back: Howler's Furs; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: High Warlord's Bonecracker; off_hand: High Warlord's Shiv; ranged: The Purifier

No-known-source sample (15 of 970, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

