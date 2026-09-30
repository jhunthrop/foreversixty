# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.6. Weights run: 1.1s. Verify run: 1.4s. 161 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.067 ± 0.020, crit=2.221 ± 0.111, hit=0.606 ± 0.034, melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 | yes | Shadow Goggles (4373, -0.40 DPS) [crafted]; Lucky Fishing Hat (19972, -0.40 DPS) [quest]; Flying Tiger Goggles (4368, -0.66 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.4 | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.30 DPS) [quest]; Tarnished Locket (279870, -0.30 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 | yes | Reinforced Woolen Shoulders (4315, -0.25 DPS) [crafted]; Forest Leather Mantle (4709, -0.25 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.42 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Catacomb Cloak (279899, -0.09 DPS, sim-verified) [quest]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.7 | yes | Brawler's Leather Armor (252490, -0.01 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.20 DPS) [crafted]; Dark Leather Tunic (2317, -0.25 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.3 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.09 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 31.1 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -1.15 DPS) [dungeon]; Forest Leather Gloves (3058, -1.25 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.64 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.6 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.5 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.15 DPS) [crafted]; Blackened Defias Boots (10402, -0.37 DPS, sim-verified) [dungeon] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 6.7 | yes | Protector's Band (20439, -0.11 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; The 1 Ring (8350, -0.26 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 | yes | Protector's Band (20439, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; The 1 Ring (8350, -0.25 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Barrens Basher (274744, -1.20 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop]; Grayson's Torch (1172, -10.71 DPS) [quest]; Pulsating Hydra Heart (5183, -10.71 DPS) [world] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.3 | yes | Fine Longbow (11304, -0.07 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Pyrewood Signet Ring; finger2: Signet of the Zhevra; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 161, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde; 20438 Outrunner's Bow

### Band 30 (night-elf, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 43.4. Weights run: 1.2s. Verify run: 1.7s. 286 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.016, crit=3.306 ± 0.128, hit=0.769 ± 0.040, melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Humbert's Helm (4724, -0.16 DPS) [world]; Tribal Worg Helm (6204, -0.16 DPS, sim-verified) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.22 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.35 DPS) [rep]; Kaleidoscope Chain (13084, -0.45 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 | yes | Dark Leather Shoulders (4252, -0.21 DPS) [crafted]; Insignia Mantle (4721, -0.21 DPS) [world_drop]; Mantle of Thieves (2264, -0.39 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Tigerstrike Mantle (13108, +0.00 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.11 DPS) [world_drop]; Cloak of Night (4447, -0.16 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, +0.00 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.19 DPS) [quest]; Green Leather Armor (4255, -0.34 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.02 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.16 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.16 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 46.3 | yes | Heavy Earthen Gloves (7359, +0.00 DPS, sim-verified) [crafted]; Pilferer's Gloves (7358, -1.77 DPS) [crafted]; Wolfclaw Gloves (1978, -1.87 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.67 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Petrolspill Leggings (9509, -0.46 DPS, sim-verified) [dungeon]; Troll's Bane Leggings (13114, -0.50 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.55 DPS) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.15 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.15 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Pyrewood Signet Ring (277210, -0.11 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.00 DPS, sim-verified) [quest]; Pyrewood Signet Ring (277210, -0.07 DPS) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Zealot Blade (13033, -0.57 DPS) [world_drop]; Electrocutioner Leg (9446, -0.62 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (43.4 DPS) | yes | Ironspine's Fist (7687, -0.72 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -14.96 DPS) [world_drop]; Totem of Infliction (1131, -15.01 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Silver Star (3463, -0.16 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop]; Crystalpine Stinger (13037, -0.22 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14145 Cursed Felblade

### Band 40 (night-elf, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 86.2. Weights run: 1.2s. Verify run: 1.6s. 409 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.105 ± 0.014, crit=2.769 ± 0.073, hit=0.555 ± 0.022, melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 14.4 | yes | White Bandit Mask (10008, -0.07 DPS) [crafted]; Hawkeye's Helm (14591, -0.07 DPS) [world_drop]; Nightscape Headband (8176, -0.10 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.17 DPS) [rep]; Sentinel's Medallion (20444, -0.25 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 | yes | Nightscape Shoulders (8192, -0.40 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.48 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.00 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.04 DPS) [crafted]; Tigerstrike Mantle (13108, -0.04 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 18.8 | yes | Raptorbane Armor (3566, -0.09 DPS) [quest]; Dusky Leather Armor (7374, -0.11 DPS) [crafted]; Nightscape Tunic (8175, -0.21 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.37 DPS) [world_drop]; Dusky Bracers (7378, -0.37 DPS) [crafted]; Cultist's Armguards (270032, -0.69 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 58.8 | yes | Shadowskin Gloves (18238, -0.67 DPS) [crafted]; Fletcher's Gloves (7348, -1.38 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -1.43 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.20 DPS) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon]; Highlander's Chain Girdle (20090, -0.41 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.20 DPS) [quest]; Petrolspill Leggings (9509, -0.35 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.4 | yes | Imperial Leather Boots (6431, +0.00 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.07 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.07 DPS) [crafted] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [world_drop]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.07 DPS) [world_drop]; Ironspine's Eye (7686, -0.07 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Southsea Lamp (9359, -0.81 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -0.89 DPS) [vendor]; Heaven's Light (13026, -1.02 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Southsea Lamp (9359, -0.31 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -14.77 DPS) [world_drop]; Totem of Infliction (1131, -14.80 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 14.0 | yes | Swiftwind (13038, -0.21 DPS) [world_drop]; Master Hunter's Bow (17686, -0.25 DPS) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.35 DPS, sim-verified) [vendor] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (night-elf, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 123.2. Weights run: 1.3s. Verify run: 1.6s. 528 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.107 ± 0.016, crit=3.272 ± 0.082, hit=0.678 ± 0.027, melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 68.6 | yes | Eye of Theradras (17715, -1.14 DPS, sim-verified) [dungeon]; Helm of Fire (8348, -1.68 DPS) [crafted]; Lordrec Helmet (10741, -1.72 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 14.4 | yes | Sentinel's Medallion (19539, -0.04 DPS) [rep]; Sentinel's Medallion (19540, -0.07 DPS) [rep]; Ghostshard Talisman (7731, -0.33 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 60.6 | yes | Sunburn Spaulders (274751, -0.16 DPS, sim-verified) [vendor]; Forest Tracker Epaulets (2278, -1.63 DPS) [world_drop]; Nightscape Shoulders (8192, -1.63 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 16.6 | yes | Serpentskin Cloak (8259, -0.15 DPS) [world_drop]; Blackflame Cape (13109, -0.15 DPS, sim-verified) [world_drop]; Nightscape Cloak (8195, -0.19 DPS) [crafted] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 70.6 | yes | Blazewind Breastplate (11193, -0.27 DPS, sim-verified) [quest]; Warbear Harness (15064, -1.71 DPS) [crafted]; Charred Leather Tunic (19127, -1.71 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.26 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.8 | yes | Sergeant Major's Leather Gauntlets (220856, -0.24 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -0.67 DPS) [crafted]; Shadowskin Gloves (18238, -0.67 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 65.8 | yes | Highlander's Lizardhide Girdle (20103, -0.67 DPS) [rep]; Highlander's Cloth Girdle (20097, -0.80 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -1.21 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (42.7 DPS) | yes | Stormshroud Pants (15057, -0.97 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -1.50 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.60 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.14 DPS) [vendor]; Swampwalker Boots (2276, -0.26 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 26.8 | yes | Masons Fraternity Ring (9533, -0.38 DPS) [quest]; Insurgent's Band (272065, -0.40 DPS) [vendor]; Insurgent's Band (272066, -0.50 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Masons Fraternity Ring (9533, -0.05 DPS, sim-verified) [quest]; Insurgent's Band (272065, -0.17 DPS) [vendor]; Insurgent's Band (272066, -0.27 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (41.6 DPS) | yes | Thunderbrew's Boot Flask (744, -0.21 DPS) [quest]; Tidal Charm (1404, -0.21 DPS) [vendor]; Smoking Heart of the Mountain (11811, -0.42 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (41.6 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Inventor's Focal Sword (17719, -0.48 DPS) [dungeon]; Julie's Dagger (6660, -1.41 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (121.7 DPS) | yes | Claw of Celebras (17738, -1.49 DPS) [dungeon]; Thermotastic Egg Timer (9644, -17.65 DPS) [quest]; Inventor's Focal Sword (17719, -80.05 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (41.6 DPS) | yes | Skull Splitting Crossbow (13039, -0.05 DPS) [world_drop]; The Silencer (13138, -0.05 DPS) [world_drop]; Dark Iron Rifle (16004, -0.74 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Assault Band; trinket1: Frozen Heart of the Mountain; trinket2: Guardian Talisman; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 528, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (night-elf, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 255.0. Weights run: 1.3s. Verify run: 1.8s. 1013 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=0.863 ± 0.035, melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 130.4 | yes | Bloodvine Lens (19998, -0.49 DPS) [crafted]; Outlaw's Collar (279253, -0.66 DPS) [crafted]; Ragefury Eyepatch (11735, -13.81 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (247.1 DPS) | yes | Blazefury Medallion (17111, -1.57 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -1.63 DPS) [quest]; Dragonheart Necklace (20622, -1.91 DPS) [world] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 88.3 | yes | Lieutenant Commander's Leather Shoulders (23313, +0.00 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [vendor]; Knight-Lieutenant's Leather Shoulders (220852, -0.39 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 57.7 | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Cloak of the Honor Guard (20073, -0.60 DPS) [rep]; Cape of the Black Baron (13340, -0.69 DPS) [dungeon] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | 121.8 | yes | Dawn Armor (252483, -0.45 DPS) [crafted]; Knight-Captain's Leather Chestpiece (23298, -0.71 DPS) [vendor]; Stormshroud Armor (15056, -11.62 DPS, sim-verified) [crafted] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 36.9 | yes | Duskwraith Wristguards (239547, -0.41 DPS) [vendor]; Marshal's Leather Armsplints (16460, -0.51 DPS) [pvp]; Primal Batskin Bracers (19687, -2.22 DPS, sim-verified) [crafted] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 105.9 | yes | Marshal's Leather Handgrips (16454, -0.85 DPS) [vendor]; Marshal's Leather Handgrips (231544, -0.85 DPS) [vendor]; Devilsaur Gauntlets (15063, -7.66 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 103.7 | yes | Highlander's Leather Girdle (20115, -0.86 DPS) [rep]; Belt of the Archmage (18405, -1.52 DPS) [crafted]; Highlander's Leather Girdle (20045, -6.62 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | sim-verified (255.0 DPS) | yes | Stormshroud Pants (15057, -0.53 DPS) [crafted]; Knight-Captain's Leather Legguards (16419, -0.53 DPS) [pvp]; Sentinel's Leather Pants (237818, -8.69 DPS, sim-verified) [vendor] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 105.9 | yes | Duskwraith Treads (239553, -2.51 DPS) [vendor]; Highlander's Leather Boots (20052, -2.53 DPS) [rep]; Darkmantle Footpads (226831, -6.39 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (247.1 DPS) | yes | Band of the Penitent (13217, -0.82 DPS) [quest]; Ring of Entropy (18543, -0.82 DPS) [world]; Wrath of Cenarius (21190, -6.74 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (247.1 DPS) | yes | Band of the Penitent (13217, -0.29 DPS) [quest]; Ring of Entropy (18543, -0.29 DPS) [world]; Wrath of Cenarius (21190, -5.91 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (243.3 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (247.1 DPS) | yes | Thunderbrew's Boot Flask (744, -0.26 DPS) [quest]; Tidal Charm (1404, -0.26 DPS) [vendor]; Shard of the Fallen Star (21891, -3.19 DPS, sim-verified) [world_drop] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (247.1 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Masterwork Stormhammer (12794, -9.64 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 813.5 | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Hammer of Bestial Fury (20580, -149.69 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (247.1 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.39 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.45 DPS) [world_drop]; Dark Iron Rifle (16004, -1.84 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1013, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 34.9. Weights run: 1.1s. Verify run: 1.4s. 165 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.067 ± 0.020, crit=2.221 ± 0.111, hit=0.606 ± 0.034, melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 | yes | Shadow Goggles (4373, -0.40 DPS) [crafted]; Lucky Fishing Hat (19972, -0.40 DPS) [quest]; Flying Tiger Goggles (4368, -0.63 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.4 | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.30 DPS) [quest]; Tarnished Locket (279870, -0.30 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 | yes | Reinforced Woolen Shoulders (4315, -0.25 DPS) [crafted]; Forest Leather Mantle (4709, -0.25 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.39 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Catacomb Cloak (279899, -0.08 DPS, sim-verified) [quest]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.5 | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Prospector's Chestpiece (14562, -0.05 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.33 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.3 | yes | Wolf Bracers (4794, -0.08 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 31.1 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -1.15 DPS) [dungeon]; Forest Leather Gloves (3058, -1.25 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.64 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.6 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.5 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.15 DPS) [crafted]; Blackened Defias Boots (10402, -0.35 DPS, sim-verified) [dungeon] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 6.7 | yes | Legionnaire's Band (20429, -0.11 DPS) [rep]; Bounty Hunter's Ring (5351, -0.16 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 | yes | Legionnaire's Band (20429, +0.00 DPS, sim-verified) [rep]; Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Wingblade (6504, -1.12 DPS) [quest] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop]; Grayson's Torch (1172, -10.71 DPS) [quest]; Nightglow Concoction (3451, -10.71 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.3 | yes | Fine Longbow (11304, -0.05 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Pyrewood Signet Ring; finger2: Signet of the Zhevra; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 165, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance; 241089 Scarlet Dagger; 254779 A'sharahm, the Roiling Tempest

### Band 30 (troll, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 43.0. Weights run: 1.2s. Verify run: 1.7s. 292 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.016, crit=3.306 ± 0.128, hit=0.769 ± 0.040, melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Humbert's Helm (4724, -0.16 DPS) [world]; Tribal Worg Helm (6204, -0.17 DPS, sim-verified) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.21 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.35 DPS) [rep]; Kaleidoscope Chain (13084, -0.45 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 | yes | Dark Leather Shoulders (4252, -0.21 DPS) [crafted]; Insignia Mantle (4721, -0.21 DPS) [world_drop]; Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Tigerstrike Mantle (13108, -0.06 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.11 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 15.4 | yes | Panther Armor (6670, -0.28 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.31 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.31 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.00 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.16 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.16 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 46.3 | yes | Heavy Earthen Gloves (7359, +0.00 DPS, sim-verified) [crafted]; Pilferer's Gloves (7358, -1.77 DPS) [crafted]; Braced Handguards (6784, -1.82 DPS) [quest] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.36 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Petrolspill Leggings (9509, -0.44 DPS, sim-verified) [dungeon]; Troll's Bane Leggings (13114, -0.50 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.55 DPS) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.15 DPS) [world_drop]; Warsong Boots (16977, -0.15 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Pyrewood Signet Ring (277210, -0.11 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, +0.00 DPS, sim-verified) [quest]; Pyrewood Signet Ring (277210, -0.07 DPS) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Zealot Blade (13033, -0.57 DPS) [world_drop]; Electrocutioner Leg (9446, -0.62 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (43.0 DPS) | yes | Ironspine's Fist (7687, -0.79 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -14.96 DPS) [world_drop]; Grayson's Torch (1172, -15.01 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 | yes | Silver Star (3463, -0.14 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop]; Crystalpine Stinger (13037, -0.22 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 83.9. Weights run: 1.2s. Verify run: 1.6s. 415 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.105 ± 0.014, crit=2.769 ± 0.073, hit=0.555 ± 0.022, melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 14.4 | yes | White Bandit Mask (10008, -0.07 DPS) [crafted]; Hawkeye's Helm (14591, -0.07 DPS) [world_drop]; Nightscape Headband (8176, -0.11 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.17 DPS) [rep]; Scout's Medallion (20442, -0.25 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 | yes | Nightscape Shoulders (8192, -0.40 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.48 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.04 DPS) [world_drop]; Parachute Cloak (10518, -0.04 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 18.8 | yes | Dusky Leather Armor (7374, -0.11 DPS) [crafted]; Hawkeye's Tunic (14592, -0.18 DPS) [world_drop]; Nightscape Tunic (8175, -0.22 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.37 DPS) [world_drop]; Dusky Bracers (7378, -0.37 DPS) [crafted]; Cultist's Armguards (270032, -0.68 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 58.8 | yes | Shadowskin Gloves (18238, -0.67 DPS) [crafted]; Fletcher's Gloves (7348, -1.36 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -1.43 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.20 DPS) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon]; Defiler's Chain Girdle (20152, -0.41 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.20 DPS) [quest]; Petrolspill Leggings (9509, -0.35 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.4 | yes | Imperial Leather Boots (6431, +0.00 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.07 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.07 DPS) [crafted] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [world_drop]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.07 DPS) [world_drop]; Ironspine's Eye (7686, -0.07 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Southsea Lamp (9359, -0.81 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -0.89 DPS) [vendor]; Heaven's Light (13026, -1.02 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Southsea Lamp (9359, -0.03 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -14.77 DPS) [world_drop]; Grayson's Torch (1172, -14.80 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 14.0 | yes | Swiftwind (13038, -0.21 DPS) [world_drop]; Master Hunter's Bow (17686, -0.25 DPS) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.34 DPS, sim-verified) [vendor] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 415, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 126.3. Weights run: 1.3s. Verify run: 1.6s. 534 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.107 ± 0.016, crit=3.272 ± 0.082, hit=0.678 ± 0.027, melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Leather Headband (220851) | Lady Palanseer [vendor] | 68.6 | yes | Eye of Theradras (17715, -1.38 DPS, sim-verified) [dungeon]; Helm of Fire (8348, -1.68 DPS) [crafted]; Sprightring Helm (17776, -1.75 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 14.4 | yes | Scout's Medallion (19535, -0.04 DPS) [rep]; Scout's Medallion (19536, -0.07 DPS) [rep]; Ghostshard Talisman (7731, -0.35 DPS, sim-verified) [dungeon] |
| shoulder | Blood Guard's Leather Shoulders (220853) | Lady Palanseer [vendor] | 60.6 | yes | Sunburn Spaulders (274751, -0.42 DPS, sim-verified) [vendor]; Forest Tracker Epaulets (2278, -1.63 DPS) [world_drop]; Nightscape Shoulders (8192, -1.63 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 16.6 | yes | Serpentskin Cloak (8259, -0.15 DPS) [world_drop]; Blackflame Cape (13109, -0.17 DPS, sim-verified) [world_drop]; Nightscape Cloak (8195, -0.19 DPS) [crafted] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 70.6 | yes | Blazewind Breastplate (11193, -0.53 DPS, sim-verified) [quest]; Warbear Harness (15064, -1.71 DPS) [crafted]; Charred Leather Tunic (19127, -1.71 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.26 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.8 | yes | First Sergeant's Leather Gauntlets (220857, -0.24 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -0.67 DPS) [crafted]; Shadowskin Gloves (18238, -0.67 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 65.8 | yes | Defiler's Lizardhide Girdle (20174, -0.67 DPS) [rep]; Defiler's Cloth Girdle (20165, -0.80 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -1.21 DPS) [rep] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | sim-verified (44.9 DPS) | yes | Stormshroud Pants (15057, -0.94 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -1.50 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.60 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; First Sergeant's Leather Boots (220861, -0.14 DPS) [vendor]; Swampwalker Boots (2276, -0.26 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 26.8 | yes | Assault Band (13095, -0.23 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.38 DPS) [quest]; Insurgent's Band (272065, -0.40 DPS) [vendor] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Assault Band (13095, -0.16 DPS, sim-verified) [world_drop]; Masons Fraternity Ring (9533, -0.29 DPS) [quest]; Insurgent's Band (272065, -0.30 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (43.8 DPS) | yes | Tidal Charm (1404, -1.58 DPS) [vendor]; Guardian Talisman (1490, -1.58 DPS) [quest]; Smoking Heart of the Mountain (11811, -2.22 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (43.8 DPS) | yes | Tidal Charm (1404, -0.21 DPS) [vendor]; Guardian Talisman (1490, -0.21 DPS) [quest]; Smoking Heart of the Mountain (11811, -0.67 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (43.8 DPS) | yes | Hanzo Sword (8190, +0.00 DPS, sim-verified) [world_drop]; Inventor's Focal Sword (17719, -0.48 DPS) [dungeon]; Julie's Dagger (6660, -1.41 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (124.0 DPS) | yes | Claw of Celebras (17738, -1.49 DPS) [dungeon]; White Bone Shredder (11863, -2.82 DPS) [quest]; Inventor's Focal Sword (17719, -80.02 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (43.8 DPS) | yes | Skull Splitting Crossbow (13039, -0.05 DPS) [world_drop]; The Silencer (13138, -0.05 DPS) [world_drop]; Dark Iron Rifle (16004, -0.76 DPS, sim-verified) [crafted] |

**New at 50:** head: Blood Guard's Leather Headband; neck: Skibi's Pendant; shoulder: Blood Guard's Leather Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 534, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 254.9. Weights run: 1.3s. Verify run: 1.8s. 1020 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=0.863 ± 0.035, melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 130.4 | yes | Bloodvine Lens (19998, -0.49 DPS) [crafted]; Outlaw's Collar (279253, -0.66 DPS) [crafted]; Ragefury Eyepatch (11735, -13.67 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (248.0 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -1.63 DPS) [quest]; Dragonheart Necklace (20622, -1.91 DPS) [world] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 88.3 | yes | Champion's Leather Shoulders (23258, +0.00 DPS) [vendor]; Champion's Leather Shoulders (227056, +0.00 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, -1.02 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 57.7 | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Deathguard's Cloak (20068, -0.60 DPS) [rep]; Cape of the Black Baron (13340, -0.69 DPS) [dungeon] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | 121.8 | yes | Dawn Armor (252483, -0.45 DPS) [crafted]; Legionnaire's Leather Chestpiece (22879, -0.71 DPS) [vendor]; Stormshroud Armor (15056, -10.99 DPS, sim-verified) [crafted] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 36.9 | yes | Duskwraith Wristguards (239547, -0.41 DPS) [vendor]; General's Leather Armsplints (16559, -0.51 DPS) [pvp]; Primal Batskin Bracers (19687, -2.35 DPS, sim-verified) [crafted] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 105.9 | yes | General's Leather Mitts (16560, -0.85 DPS) [vendor]; General's Leather Mitts (231555, -0.85 DPS) [vendor]; Devilsaur Gauntlets (15063, -7.22 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 103.7 | yes | Defiler's Leather Girdle (20193, -0.86 DPS) [rep]; Belt of the Archmage (18405, -1.52 DPS) [crafted]; Defiler's Leather Girdle (20190, -6.14 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | sim-verified (254.9 DPS) | yes | Stormshroud Pants (15057, -0.53 DPS) [crafted]; Legionnaire's Leather Leggings (16508, -0.53 DPS) [pvp]; Sentinel's Leather Pants (237818, -8.78 DPS, sim-verified) [vendor] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 105.9 | yes | Duskwraith Treads (239553, -2.51 DPS) [vendor]; Defiler's Leather Boots (20186, -2.53 DPS) [rep]; Darkmantle Footpads (226831, -6.87 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (248.0 DPS) | yes | Band of the Penitent (13217, -0.82 DPS) [quest]; Ring of Entropy (18543, -0.82 DPS) [world]; Wrath of Cenarius (21190, -6.28 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (248.0 DPS) | yes | Band of the Penitent (13217, -0.29 DPS) [quest]; Ring of Entropy (18543, -0.29 DPS) [world]; Wrath of Cenarius (21190, -5.47 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (245.4 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (248.0 DPS) | yes | Tidal Charm (1404, -1.59 DPS) [vendor]; Guardian Talisman (1490, -1.59 DPS) [quest]; Frozen Heart of the Mountain (249469, -3.02 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (248.0 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Ebon Hand (19170, -7.39 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 813.5 | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Hammer of Bestial Fury (20580, -147.74 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (248.0 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.39 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.45 DPS) [world_drop]; Dark Iron Rifle (16004, -2.08 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1020, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

