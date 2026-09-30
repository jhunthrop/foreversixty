# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.3. Weights run: 1.1s. Verify run: 1.3s. 161 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.404 ± 0.062, hit=0.311 ± 0.030, melee_haste=not significant (0.665 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.38 DPS) [crafted]; Lucky Fishing Hat (19972, -0.38 DPS) [quest]; Flying Tiger Goggles (4368, -0.54 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 | yes | Erudite's Amulet (277204, -0.14 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.29 DPS) [quest]; Tarnished Locket (279870, -0.29 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Forest Leather Mantle (4709, -0.24 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.34 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.01 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.1 | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.19 DPS) [crafted]; Dark Leather Tunic (2317, -0.24 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.07 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.02 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.33 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Protector's Band (20439, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 5.3 | yes | Protector's Band (20439, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon]; The 1 Ring (8350, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Barrens Basher (274744, -1.21 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop]; Grayson's Torch (1172, -10.85 DPS) [quest]; Pulsating Hydra Heart (5183, -10.85 DPS) [world] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 | yes | Fine Longbow (11304, -0.01 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 161, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde; 20438 Outrunner's Bow

### Band 30 (night-elf, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.8. Weights run: 1.1s. Verify run: 1.6s. 286 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.485 ± 0.060, hit=0.428 ± 0.033, melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.14 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.36 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.38 DPS) [rep]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.39 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Tigerstrike Mantle (13108, -0.10 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.19 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.08 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.24 DPS) [quest]; Green Leather Armor (4255, -0.38 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.10 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.19 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Fletcher's Gloves (7348, -0.44 DPS) [crafted]; Wolfclaw Gloves (1978, -0.48 DPS) [dungeon]; Pilferer's Gloves (7358, -0.49 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.72 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.73 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.04 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.10 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Zealot Blade (13033, -0.58 DPS) [world_drop]; Electrocutioner Leg (9446, -0.63 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (46.8 DPS) | yes | Ironspine's Fist (7687, -0.50 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.23 DPS) [world_drop]; Totem of Infliction (1131, -15.28 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Silver Star (3463, -0.24 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop]; Crystalpine Stinger (13037, -0.24 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14145 Cursed Felblade

### Band 40 (night-elf, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 90.9. Weights run: 1.2s. Verify run: 1.5s. 409 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.925 ± 0.148, hit=0.646 ± 0.078, melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.3 | yes | White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop]; Nightscape Headband (8176, -0.12 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.29 DPS) [rep]; Sentinel's Medallion (20444, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 | yes | Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.00 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 17.4 | yes | Nightscape Tunic (8175, -0.10 DPS) [crafted]; Dusky Leather Armor (7374, -0.15 DPS) [crafted]; Raptorbane Armor (3566, -0.73 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.81 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.0 | yes | Fletcher's Gloves (7348, -0.99 DPS) [crafted]; Shadowskin Gloves (18238, -0.99 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.47 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.49 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Petrolspill Leggings (9509, -0.58 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.57 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Falcon's Hook (7552, -0.53 DPS) [world_drop]; Ironspine's Eye (7686, -0.53 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.14 DPS) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Southsea Lamp (9359, -1.23 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.35 DPS) [vendor]; Heaven's Light (13026, -1.53 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Southsea Lamp (9359, -4.49 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -21.88 DPS) [world_drop]; Totem of Infliction (1131, -21.93 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 14.0 | yes | Swiftwind (13038, -0.34 DPS) [world_drop]; Master Hunter's Bow (17686, -0.39 DPS) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.40 DPS, sim-verified) [vendor] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (night-elf, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 129.9. Weights run: 1.2s. Verify run: 1.5s. 528 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=6.578 ± 0.398, hit=0.803 ± 0.104, melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 116.1 | yes | Eye of Theradras (17715, -2.03 DPS, sim-verified) [dungeon]; Helm of Fire (8348, -5.19 DPS) [crafted]; Lordrec Helmet (10741, -5.25 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 14.6 | yes | Sentinel's Medallion (19539, -0.06 DPS) [rep]; Sentinel's Medallion (19540, -0.12 DPS) [rep]; Ghostshard Talisman (7731, -0.83 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 108.1 | yes | Sunburn Spaulders (274751, -0.25 DPS, sim-verified) [vendor]; Forest Tracker Epaulets (2278, -5.12 DPS) [world_drop]; Nightscape Shoulders (8192, -5.12 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 16.9 | yes | Serpentskin Cloak (8259, -0.24 DPS) [world_drop]; Nightscape Cloak (8195, -0.30 DPS) [crafted]; Blackflame Cape (13109, -0.35 DPS, sim-verified) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 118.1 | yes | Blazewind Breastplate (11193, -0.36 DPS, sim-verified) [quest]; Warbear Harness (15064, -5.23 DPS) [crafted]; Charred Leather Tunic (19127, -5.23 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 112.1 | yes | Sergeant Major's Leather Gauntlets (220856, -0.47 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.07 DPS) [crafted]; Shadowskin Gloves (18238, -1.07 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 112.1 | yes | Highlander's Lizardhide Girdle (20103, -1.07 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.58 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -4.39 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (112.1 DPS) | yes | Stormshroud Pants (15057, -1.67 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -4.92 DPS) [dungeon]; Basilisk Hide Pants (1718, -5.05 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.24 DPS) [vendor]; Swampwalker Boots (2276, -0.42 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Masons Fraternity Ring (9533, -0.66 DPS) [quest]; Insurgent's Band (272065, -0.70 DPS) [vendor]; Insurgent's Band (272066, -0.86 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Masons Fraternity Ring (9533, -0.01 DPS, sim-verified) [quest]; Insurgent's Band (272065, -0.27 DPS) [vendor]; Insurgent's Band (272066, -0.43 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (111.2 DPS) | yes | Thunderbrew's Boot Flask (744, -0.39 DPS) [quest]; Tidal Charm (1404, -0.39 DPS) [vendor]; Smoking Heart of the Mountain (11811, -0.67 DPS, sim-verified) [crafted] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | sim-verified (111.2 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Hammer of the Northern Wind (810, -1.71 DPS) [world_drop]; Julie's Dagger (6660, -3.94 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (128.1 DPS) | yes | Claw of Celebras (17738, -2.37 DPS) [dungeon]; Hammer of the Northern Wind (810, -17.63 DPS, sim-verified) [world_drop]; Thermotastic Egg Timer (9644, -27.96 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (111.2 DPS) | yes | Skull Splitting Crossbow (13039, -0.09 DPS) [world_drop]; The Silencer (13138, -0.09 DPS) [world_drop]; Dark Iron Rifle (16004, -1.64 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Assault Band; trinket1: Guardian Talisman; trinket2: Frozen Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 528, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (night-elf, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 253.2. Weights run: 1.2s. Verify run: 1.7s. 1013 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=0.900 ± 0.162, melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ragefury Eyepatch (11735) (or Bloodvine Lens (19998)) | Blackrock Depths: Guzzler [dungeon] | 230.2 | yes | Bloodvine Lens (19998, -0.69 DPS, sim-verified) [crafted]; Outlaw's Collar (279253, -1.51 DPS) [crafted]; Duskwraith Helmet (239560, -1.54 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (240.0 DPS) | yes | Blazefury Medallion (17111, -0.77 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -5.68 DPS) [quest]; Amulet of the Darkmoon (19491, -6.03 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 152.8 | yes | Lieutenant Commander's Leather Shoulders (23313, -0.36 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.36 DPS) [vendor]; Knight-Lieutenant's Leather Shoulders (220852, -2.05 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 115.1 | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Cloak of the Honor Guard (20073, -3.97 DPS) [rep]; Cape of the Black Baron (13340, -3.97 DPS) [dungeon] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | sim-verified (253.2 DPS) | yes | Dawn Armor (252483, -0.90 DPS) [crafted]; Field Marshal's Leather Chestpiece (16453, -1.72 DPS) [vendor]; Stormshroud Armor (15056, -12.82 DPS, sim-verified) [crafted] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 43.9 | yes | Duskwraith Wristguards (239547, -0.82 DPS) [vendor]; Marshal's Leather Armsplints (16460, -0.93 DPS) [pvp]; Primal Batskin Bracers (19687, -2.84 DPS, sim-verified) [crafted] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 173.0 | yes | Marshal's Leather Handgrips (16454, -1.60 DPS) [vendor]; Marshal's Leather Handgrips (231544, -1.60 DPS) [vendor]; Devilsaur Gauntlets (15063, -6.84 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 170.2 | yes | Highlander's Leather Girdle (20115, -1.88 DPS) [rep]; Belt of the Archmage (18405, -2.95 DPS) [crafted]; Highlander's Leather Girdle (20045, -5.47 DPS, sim-verified) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 267.9 | yes | Knight-Captain's Leather Legguards (16419, -2.02 DPS) [pvp]; Sentinel's Silk Leggings (237815, -2.02 DPS) [vendor]; Stormshroud Pants (15057, -5.02 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 173.0 | yes | Duskwraith Treads (239553, -7.36 DPS) [vendor]; Darkmantle Footpads (226831, -7.43 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -7.47 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (240.0 DPS) | yes | Band of the Penitent (13217, -1.34 DPS) [quest]; Ring of Entropy (18543, -1.34 DPS) [world]; Wrath of Cenarius (21190, -5.77 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (240.0 DPS) | yes | Band of the Penitent (13217, -0.48 DPS) [quest]; Ring of Entropy (18543, -0.48 DPS) [world]; Wrath of Cenarius (21190, -4.64 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (240.0 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -7.86 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (240.0 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Ebon Hand (19170, -5.17 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 815.9 | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Persuader (22384, -14.90 DPS, sim-verified) [crafted] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (240.0 DPS) | yes | Dark Iron Rifle (16004, -2.08 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -5.12 DPS) [world_drop]; Skull Splitting Crossbow (13039, -5.41 DPS) [world_drop] |

**New at 60:** head: Ragefury Eyepatch; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Sentinel's Leather Pants; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Blue Dragon; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1013, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.8. Weights run: 1.1s. Verify run: 1.4s. 165 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.404 ± 0.062, hit=0.311 ± 0.030, melee_haste=not significant (0.665 ± 0.825)

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
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Legionnaire's Band (20429, -0.10 DPS) [rep]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 5.3 | yes | Legionnaire's Band (20429, +0.00 DPS, sim-verified) [rep]; Bounty Hunter's Ring (5351, -0.11 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Wingblade (6504, -1.15 DPS) [quest] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop]; Grayson's Torch (1172, -10.85 DPS) [quest]; Nightglow Concoction (3451, -10.85 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 | yes | Fine Longbow (11304, -0.00 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 165, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance; 241089 Scarlet Dagger; 254779 A'sharahm, the Roiling Tempest

### Band 30 (troll, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.3. Weights run: 1.1s. Verify run: 1.6s. 292 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.485 ± 0.060, hit=0.428 ± 0.033, melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.38 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.38 DPS) [rep]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 | yes | Panther Armor (6670, -0.20 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.29 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.29 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.12 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.19 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Braced Handguards (6784, -0.43 DPS) [quest]; Fletcher's Gloves (7348, -0.44 DPS) [crafted]; Pilferer's Gloves (7358, -0.51 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Deftkin Belt (16659, -0.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.76 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.06 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Warsong Boots (16977, -0.19 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.12 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Zealot Blade (13033, -0.58 DPS) [world_drop]; Electrocutioner Leg (9446, -0.63 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (46.3 DPS) | yes | Ironspine's Fist (7687, -0.64 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.23 DPS) [world_drop]; Grayson's Torch (1172, -15.28 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop]; Crystalpine Stinger (13037, -0.24 DPS) [world_drop]; Silver Star (3463, -0.25 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 89.3. Weights run: 1.2s. Verify run: 1.5s. 415 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.925 ± 0.148, hit=0.646 ± 0.078, melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.3 | yes | White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop]; Nightscape Headband (8176, -0.12 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.29 DPS) [rep]; Scout's Medallion (20442, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 | yes | Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.09 DPS) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 17.4 | yes | Dusky Leather Armor (7374, -0.15 DPS) [crafted]; Nightscape Tunic (8175, -0.23 DPS, sim-verified) [crafted]; Hawkeye's Tunic (14592, -0.25 DPS) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.80 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.0 | yes | Fletcher's Gloves (7348, -0.99 DPS) [crafted]; Shadowskin Gloves (18238, -0.99 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.50 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.48 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Petrolspill Leggings (9509, -0.58 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.61 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Falcon's Hook (7552, -0.53 DPS) [world_drop]; Ironspine's Eye (7686, -0.53 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.14 DPS) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Southsea Lamp (9359, -1.23 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.35 DPS) [vendor]; Heaven's Light (13026, -1.53 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Southsea Lamp (9359, -4.42 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -21.88 DPS) [world_drop]; Grayson's Torch (1172, -21.93 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 14.0 | yes | Swiftwind (13038, -0.34 DPS) [world_drop]; Master Hunter's Bow (17686, -0.39 DPS) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.40 DPS, sim-verified) [vendor] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 415, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 132.5. Weights run: 1.2s. Verify run: 1.5s. 534 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=6.578 ± 0.398, hit=0.803 ± 0.104, melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Leather Headband (220851) | Lady Palanseer [vendor] | 116.1 | yes | Eye of Theradras (17715, -2.42 DPS, sim-verified) [dungeon]; Helm of Fire (8348, -5.19 DPS) [crafted]; Sprightring Helm (17776, -5.31 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 14.6 | yes | Scout's Medallion (19535, -0.06 DPS) [rep]; Scout's Medallion (19536, -0.12 DPS) [rep]; Ghostshard Talisman (7731, -0.79 DPS, sim-verified) [dungeon] |
| shoulder | Blood Guard's Leather Shoulders (220853) | Lady Palanseer [vendor] | 108.1 | yes | Sunburn Spaulders (274751, -0.67 DPS, sim-verified) [vendor]; Forest Tracker Epaulets (2278, -5.12 DPS) [world_drop]; Nightscape Shoulders (8192, -5.12 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 16.9 | yes | Serpentskin Cloak (8259, -0.24 DPS) [world_drop]; Nightscape Cloak (8195, -0.30 DPS) [crafted]; Blackflame Cape (13109, -0.34 DPS, sim-verified) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 118.1 | yes | Blazewind Breastplate (11193, -0.76 DPS, sim-verified) [quest]; Warbear Harness (15064, -5.23 DPS) [crafted]; Charred Leather Tunic (19127, -5.23 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 112.1 | yes | First Sergeant's Leather Gauntlets (220857, -0.47 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.07 DPS) [crafted]; Shadowskin Gloves (18238, -1.07 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 112.1 | yes | Defiler's Lizardhide Girdle (20174, -1.07 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.58 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -4.39 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 184.2 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS, sim-verified) [vendor]; Ferine Leggings (6690, -8.45 DPS) [dungeon]; Basilisk Hide Pants (1718, -8.58 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; First Sergeant's Leather Boots (220861, -0.24 DPS) [vendor]; Swampwalker Boots (2276, -0.42 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Assault Band (13095, -0.43 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.66 DPS) [quest]; Insurgent's Band (272065, -0.70 DPS) [vendor] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Assault Band (13095, -0.32 DPS, sim-verified) [world_drop]; Masons Fraternity Ring (9533, -0.44 DPS) [quest]; Insurgent's Band (272065, -0.48 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (114.8 DPS) | yes | Tidal Charm (1404, -2.55 DPS) [vendor]; Guardian Talisman (1490, -2.55 DPS) [quest]; Smoking Heart of the Mountain (11811, -4.22 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (114.8 DPS) | yes | Tidal Charm (1404, -0.39 DPS) [vendor]; Guardian Talisman (1490, -0.39 DPS) [quest]; Smoking Heart of the Mountain (11811, -1.06 DPS, sim-verified) [crafted] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | sim-verified (114.8 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Hammer of the Northern Wind (810, -1.71 DPS) [world_drop]; Julie's Dagger (6660, -3.94 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (132.5 DPS) | yes | Claw of Celebras (17738, -2.37 DPS) [dungeon]; White Bone Shredder (11863, -4.47 DPS) [quest]; Hammer of the Northern Wind (810, -19.14 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (114.8 DPS) | yes | Skull Splitting Crossbow (13039, -0.09 DPS) [world_drop]; The Silencer (13138, -0.09 DPS) [world_drop]; Dark Iron Rifle (16004, -1.62 DPS, sim-verified) [crafted] |

**New at 50:** head: Blood Guard's Leather Headband; neck: Skibi's Pendant; shoulder: Blood Guard's Leather Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 534, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 257.7. Weights run: 1.2s. Verify run: 1.8s. 1020 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=0.900 ± 0.162, melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ragefury Eyepatch (11735) (or Bloodvine Lens (19998)) | Blackrock Depths: Guzzler [dungeon] | 230.2 | yes | Bloodvine Lens (19998, -0.69 DPS, sim-verified) [crafted]; Outlaw's Collar (279253, -1.51 DPS) [crafted]; Duskwraith Helmet (239560, -1.54 DPS) [vendor] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (245.0 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -0.60 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 152.8 | yes | Champion's Leather Shoulders (23258, -0.36 DPS) [vendor]; Champion's Leather Shoulders (227056, -0.36 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, -2.47 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 115.1 | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Deathguard's Cloak (20068, -3.97 DPS) [rep]; Cape of the Black Baron (13340, -3.97 DPS) [dungeon] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | sim-verified (257.7 DPS) | yes | Dawn Armor (252483, -0.90 DPS) [crafted]; Warlord's Leather Breastplate (16563, -1.72 DPS) [vendor]; Stormshroud Armor (15056, -12.64 DPS, sim-verified) [crafted] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 43.9 | yes | Duskwraith Wristguards (239547, -0.82 DPS) [vendor]; General's Leather Armsplints (16559, -0.93 DPS) [pvp]; Primal Batskin Bracers (19687, -2.83 DPS, sim-verified) [crafted] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 173.0 | yes | General's Leather Mitts (16560, -1.60 DPS) [vendor]; General's Leather Mitts (231555, -1.60 DPS) [vendor]; Devilsaur Gauntlets (15063, -8.46 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 170.2 | yes | Defiler's Leather Girdle (20193, -1.88 DPS) [rep]; Belt of the Archmage (18405, -2.95 DPS) [crafted]; Defiler's Leather Girdle (20190, -7.09 DPS, sim-verified) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 267.9 | yes | Legionnaire's Leather Leggings (16508, -2.02 DPS) [pvp]; Sentinel's Silk Leggings (237815, -2.02 DPS) [vendor]; Stormshroud Pants (15057, -4.92 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 173.0 | yes | Darkmantle Footpads (226831, -7.35 DPS, sim-verified) [vendor]; Duskwraith Treads (239553, -7.36 DPS) [vendor]; Darkmantle Boots (22003, -7.47 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (244.7 DPS) | yes | Band of the Penitent (13217, -1.34 DPS) [quest]; Ring of Entropy (18543, -1.34 DPS) [world]; Wrath of Cenarius (21190, -7.41 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (244.7 DPS) | yes | Band of the Penitent (13217, -0.48 DPS) [quest]; Ring of Entropy (18543, -0.48 DPS) [world]; Wrath of Cenarius (21190, -6.29 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (239.6 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (244.7 DPS) | yes | Tidal Charm (1404, -2.59 DPS) [vendor]; Guardian Talisman (1490, -2.59 DPS) [quest]; Frozen Heart of the Mountain (249469, -4.46 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (244.7 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Ebon Hand (19170, -5.05 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 815.9 | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Persuader (22384, -15.28 DPS, sim-verified) [crafted] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (245.0 DPS) | yes | Dark Iron Rifle (16004, -2.12 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -5.12 DPS) [world_drop]; Skull Splitting Crossbow (13039, -5.41 DPS) [world_drop] |

**New at 60:** head: Ragefury Eyepatch; neck: Blazefury Medallion; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Sentinel's Leather Pants; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1020, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

