# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 101.9. Weights run: 2.5s. Verify run: 2.3s. 220 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.622 ± 0.008, crit=0.609 ± 0.016 per rating point (14 rating = 1%, 8.519 per %), hit=1.049 ± 0.042 per rating point (10 rating = 1%, 10.487 per %), melee_haste=8.371 ± 0.680

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 21.0 ranged_attack_power points (1.43 DPS) | yes | Red Winter Hat (21524, -1.45 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 15.7 ranged_attack_power points (1.07 DPS) | yes | Erudite's Amulet (277204, -0.36 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 13.1 ranged_attack_power points (0.89 DPS) | yes | Slime-encrusted Pads (6461, -0.68 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 15.7 ranged_attack_power points (1.07 DPS) | yes | Hide of Lupos (3018, -0.36 DPS) [world]; Bristlebark Cape (14571, -0.36 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.64 DPS, sim-verified) [dungeon] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 28.8 ranged_attack_power points (1.96 DPS) | yes | Trapper's Leather Armor (252491, -0.71 DPS) [crafted]; Dark Leather Tunic (2317, -0.89 DPS) [crafted]; Brawler's Leather Armor (252490, -1.06 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 13.1 ranged_attack_power points (0.89 DPS) | yes | Wolf Bracers (4794, -0.18 DPS) [vendor]; Bravo's Armbands (270015, -0.18 DPS) [quest]; Ratchet Wristwraps (274742, -0.36 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (101.9 DPS) | yes | Forest Leather Gloves (3058, -0.36 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.36 DPS) [crafted]; Serpent Gloves (5970, -1.21 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.22 DPS) | yes | Deviate Scale Belt (6468, -0.33 DPS) [crafted]; Dark Leather Belt (4249, -0.51 DPS) [crafted]; Dusty Belt (279897, -0.60 DPS, sim-verified) [quest] |
| legs | Leggings of the Fang (10410) (or Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 23.6 ranged_attack_power points (1.60 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.18 DPS) [world]; Brawler's Leather Pants (252500, -0.27 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.0 ranged_attack_power points (1.43 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.36 DPS) [dungeon]; Agile Boots (4788, -0.53 DPS) [vendor] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 15.7 ranged_attack_power points (1.07 DPS) | yes | Lavishly Jeweled Ring (1156, -0.71 DPS) [dungeon]; The 1 Ring (8350, -0.89 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.5 ranged_attack_power points (0.71 DPS) | yes | Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon]; The 1 Ring (8350, -0.53 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 27.0 ranged_attack_power points (1.83 DPS) | yes | Impaling Harpoon (5200, -0.23 DPS) [dungeon]; Scythe Axe (5749, -0.59 DPS) [world]; Lupine Axe (1220, -0.76 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.9 ranged_attack_power points (12.16 DPS) | yes | Lil Timmy's Peashooter (13136, -1.33 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.64 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.91 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 137.0. Weights run: 2.8s. Verify run: 2.8s. 366 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.649 ± 0.010, crit=0.710 ± 0.019 per rating point (14 rating = 1%, 9.944 per %), hit=1.227 ± 0.054 per rating point (10 rating = 1%, 12.269 per %), melee_haste=9.628 ± 0.925

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 26.5 ranged_attack_power points (1.79 DPS) | yes | Tribal Worg Helm (6204, -0.36 DPS) [world]; Brawler's Leather Hood (252504, -0.36 DPS) [crafted]; Humbert's Helm (4724, -0.54 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 21.2 ranged_attack_power points (1.43 DPS) | yes | Ghostshard Talisman (7731, -0.48 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.71 DPS) [world_drop]; Erudite's Amulet (277204, -0.71 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.1 ranged_attack_power points (1.96 DPS) | yes | Mantle of Thieves (2264, -0.18 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.71 DPS) [crafted]; Insignia Mantle (4721, -0.71 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 21.2 ranged_attack_power points (1.43 DPS) | yes | Hawkeye's Cloak (14593, -0.18 DPS) [world_drop]; Cloak of Night (4447, -0.36 DPS) [world]; Fenrus' Hide (6340, -0.36 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 37.1 ranged_attack_power points (2.50 DPS) | yes | Tunic of Westfall (2041, -0.54 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.07 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.07 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 15.9 ranged_attack_power points (1.07 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.18 DPS) [world_drop] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 21.2 ranged_attack_power points (1.43 DPS) | yes | Heavy Earthen Gloves (7359, -0.35 DPS) [crafted]; Serpent Gloves (5970, -0.36 DPS) [dungeon]; Insignia Gloves (6408, -0.36 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 ranged_attack_power points (1.62 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.01 DPS) [crafted]; Stalker's Leather Belt (252521, -0.01 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 37.1 ranged_attack_power points (2.50 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.18 DPS) [crafted]; Ferine Leggings (6690, -0.75 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 21.2 ranged_attack_power points (1.43 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.18 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.8 ranged_attack_power points (1.61 DPS) | yes | Ring of Precision (1491, -0.54 DPS) [dungeon]; Protector's Band (19517, -0.54 DPS) [rep]; Signet of the Zhevra (285330, -0.54 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 18.5 ranged_attack_power points (1.25 DPS) | yes | Ring of Precision (1491, -0.18 DPS) [dungeon]; Protector's Band (19517, -0.18 DPS) [rep]; Signet of the Zhevra (285330, -0.18 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 19.2 ranged_attack_power points (1.30 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 ranged_attack_power points (1.07 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.63 DPS) | yes | Silver Star (3463, -0.47 DPS, sim-verified) [quest]; Glass Shooter (9456, -0.71 DPS) [dungeon]; Ironweaver (13137, -1.30 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Highlander's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 186.7. Weights run: 3.0s. Verify run: 2.9s. 597 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.656 ± 0.010, crit=0.770 ± 0.020 per rating point (14 rating = 1%, 10.776 per %), hit=1.263 ± 0.100 per rating point (10 rating = 1%, 12.633 per %), melee_haste=9.750 ± 1.505

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 34.5 ranged_attack_power points (2.26 DPS) | yes | Guard's Chain Helm (250499, -0.17 DPS) [crafted]; Skullsplitter Helm (1624, -0.35 DPS) [world]; Nightscape Headband (8176, -0.65 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 29.2 ranged_attack_power points (1.91 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.60 DPS) [quest]; Ghostshard Talisman (7731, -1.00 DPS) [dungeon]; Erudite's Amulet (277204, -1.22 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 41.2 ranged_attack_power points (2.70 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.81 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.96 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 26.6 ranged_attack_power points (1.74 DPS) | yes | Imperial Cloak (6432, -0.35 DPS) [dungeon]; Parachute Cloak (10518, -0.35 DPS) [crafted]; Tigerstrike Mantle (13108, -0.35 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 50.5 ranged_attack_power points (3.31 DPS) | yes | Wolffear Harness (13110, -0.35 DPS) [world_drop]; Nightscape Tunic (8175, -0.70 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.70 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 21.2 ranged_attack_power points (1.39 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Branded Leather Bracers (19508, -0.08 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.17 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.10 DPS) | yes | Gloves of Holy Might (867, -0.08 DPS) [world_drop]; Dragonscale Gauntlets (8347, -0.35 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.36 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.97 DPS) | yes | Scorpashi Sash (14652, -0.23 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.39 DPS) [rep]; Blackforge Girdle (6425, -0.40 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 55.8 ranged_attack_power points (3.65 DPS) | yes | Triprunner Dungarees (9624, -0.69 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.22 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.22 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 34.5 ranged_attack_power points (2.26 DPS) | yes | Imperial Leather Boots (6431, -0.35 DPS) [dungeon]; Dusky Boots (7390, -0.35 DPS) [crafted]; Worn Running Boots (9398, -0.35 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 26.6 ranged_attack_power points (1.74 DPS) | yes | Ironspine's Eye (7686, -0.17 DPS) [dungeon]; Protector's Band (19515, -0.35 DPS) [rep]; Disengagement Ring (276202, -0.35 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 23.9 ranged_attack_power points (1.57 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Protector's Band (19515, -0.17 DPS) [rep]; Disengagement Ring (276202, -0.17 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (186.7 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.66 DPS, sim-verified) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 22.0 ranged_attack_power points (1.44 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Satyr's Rod (15962, -1.27 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (22.93 DPS) | yes | Shadowforge Bushmaster (9422, -2.05 DPS) [dungeon]; Swiftwind (13038, -2.22 DPS) [world_drop]; The Silencer (13138, -3.43 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Vanquisher's Sword; ranged: Bow of Searing Arrows

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 234.9. Weights run: 3.1s. Verify run: 3.5s. 754 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.712 ± 0.012, crit=0.921 ± 0.023 per rating point (14 rating = 1%, 12.900 per %), hit=1.697 ± 0.099 per rating point (10 rating = 1%, 16.969 per %), melee_haste=13.372 ± 1.600

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 56.3 ranged_attack_power points (3.83 DPS) | yes | Lordrec Helmet (10741, -0.88 DPS) [quest]; Sprightring Helm (17776, -1.06 DPS) [quest]; Helm of Fire (8348, -3.54 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.3 ranged_attack_power points (2.40 DPS) | yes | Sentinel's Medallion (19540, -0.37 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 48.2 ranged_attack_power points (3.28 DPS) | yes | Phytoskin Spaulders (17749, -0.32 DPS) [dungeon]; Sunburn Spaulders (274751, -0.43 DPS) [vendor]; Khan's Mantle (14787, -1.06 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 40.7 ranged_attack_power points (2.77 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.18 DPS) [dungeon]; Blackmetal Cape (9512, -0.55 DPS) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 67.8 ranged_attack_power points (4.62 DPS) | yes | Blazewind Breastplate (11193, -0.37 DPS) [quest]; Knight's Chain Armor (220828, -0.97 DPS) [vendor]; Quillward Harness (10583, -1.11 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 40.7 ranged_attack_power points (2.77 DPS) | yes | Wicked Leather Bracers (15084, -0.74 DPS) [crafted]; Arena Bands (18711, -0.86 DPS) [world]; Bloodlust Bracelets (14807, -0.90 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 59.7 ranged_attack_power points (4.06 DPS) | yes | Beastmaster's Gauntlets (226883, -1.48 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.48 DPS) [crafted]; Sergeant Major's Chain Gauntlets (220829, -1.51 DPS, sim-verified) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 53.7 ranged_attack_power points (3.66 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.07 DPS) [crafted]; Stalker's Mail Belt (252588, -1.07 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 56.9 ranged_attack_power points (3.88 DPS) | yes | Infernal Trickster Leggings (17754, -0.18 DPS) [dungeon]; Keeper's Woolies (14668, -0.37 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.41 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 54.2 ranged_attack_power points (3.70 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.37 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.55 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.0 ranged_attack_power points (2.59 DPS) | yes | Ring of the Underwood (2951, -0.74 DPS) [world_drop]; Falcon's Hook (7552, -0.92 DPS) [dungeon]; Ironspine's Eye (7686, -0.92 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 37.0 ranged_attack_power points (2.52 DPS) | yes | Falcon's Hook (7552, -0.86 DPS) [dungeon]; Ironspine's Eye (7686, -0.86 DPS) [dungeon]; Ring of the Underwood (2951, -0.91 DPS, sim-verified) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (234.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (234.9 DPS) | yes | Molten Heart of the Mountain (249470, -1.04 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (234.9 DPS) | yes | Warmonger (13052, -0.78 DPS) [world_drop]; Steel Spear (250605, -1.17 DPS) [crafted]; Hanzo Sword (8190, -4.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (234.9 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.59 DPS) [world_drop]; Hurricane (2824, -1.66 DPS) [world_drop]; Dark Iron Rifle (16004, -3.06 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 314.3. Weights run: 3.0s. Verify run: 11.7s. 1670 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.950 ± 0.021, crit=1.691 ± 0.043 per rating point (14 rating = 1%, 23.668 per %), hit=2.934 ± 0.157 per rating point (10 rating = 1%, 29.343 per %), melee_haste=17.485 ± 2.316

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (314.3 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -9.64 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 73.6 ranged_attack_power points (5.01 DPS) | yes | Beads of Ogre Might (22150, -1.38 DPS) [quest]; Amulet of the Darkmoon (19491, -1.42 DPS, sim-verified) [quest]; Mark of Fordring (15411, -1.63 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 103.3 ranged_attack_power points (7.03 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 64.2 ranged_attack_power points (4.37 DPS) | yes | Shifting Cloak (18511, -0.96 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.96 DPS) [dungeon]; Howler's Furs (272414, -1.56 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (314.3 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.60 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.60 DPS) [vendor]; Tunic of Undead Slaying (23089, -13.02 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (314.3 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -8.46 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (314.3 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (314.3 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -8.12 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 179.9 ranged_attack_power points (12.24 DPS) | yes | Sentinel's Leather Pants (237818, -3.60 DPS) [vendor]; Marshal's Chain Legguards (231577, -3.81 DPS) [pvp]; Plaguehound Leggings (18736, -4.22 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (314.3 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -8.80 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (314.3 DPS) | yes | Cutthroat's Signet (272408, -0.80 DPS) [vendor]; Don Julio's Band (19325, -0.91 DPS) [rep]; Naglering (11669, -6.83 DPS, sim-verified) [dungeon] |
| finger2 | Tarnished Elven Ring (18500) | Dire Maul: Tribute [dungeon] | sim-verified (314.3 DPS) | yes | Cutthroat's Signet (272408, -0.20 DPS) [vendor]; Don Julio's Band (19325, -0.31 DPS) [rep]; Naglering (11669, -6.05 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (314.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Devilsaur Eye (19991, -2.60 DPS, sim-verified) [quest] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (314.3 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -1.05 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (314.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -8.42 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (314.3 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Tarnished Elven Ring; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 786.9. Weights run: 3.1s. Verify run: 12.1s. 1670 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.995 ± 0.020, crit=1.815 ± 0.041 per rating point (14 rating = 1%, 25.414 per %), hit=3.766 ± 0.217 per rating point (10 rating = 1%, 37.656 per %), melee_haste=20.998 ± 2.757

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (786.9 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -10.45 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 82.6 ranged_attack_power points (11.61 DPS) | yes | Beads of Ogre Might (22150, -3.40 DPS, sim-verified) [quest]; Amulet of the Darkmoon (19491, -3.61 DPS) [quest]; Mark of Fordring (15411, -4.38 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 106.3 ranged_attack_power points (14.95 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 65.7 ranged_attack_power points (9.23 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Shifting Cloak (18511, -2.07 DPS) [crafted]; Shadow Prowler's Cloak (22269, -2.07 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (786.9 DPS) | yes | Field Marshal's Chain Breastplate (16466, -6.56 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -6.56 DPS) [vendor]; Tunic of Undead Slaying (23089, -22.33 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (786.9 DPS) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -13.95 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 102.3 ranged_attack_power points (14.38 DPS) | yes | Gauntlets of Accuracy (18349, +0.00 DPS, sim-verified) [dungeon]; Marshal's Chain Vices (231578, -1.96 DPS) [vendor]; Raider Gloves (272099, -3.01 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (786.9 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Marksman's Girdle (22232, -10.07 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 193.3 ranged_attack_power points (27.18 DPS) | yes | Sentinel's Leather Pants (237818, -8.66 DPS) [vendor]; Plaguehound Leggings (18736, -9.25 DPS) [dungeon]; Marshal's Chain Legguards (231577, -9.29 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (786.9 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -12.89 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (786.9 DPS) | yes | Cutthroat's Signet (272408, -1.68 DPS) [vendor]; Don Julio's Band (19325, -1.76 DPS) [rep]; Naglering (11669, -9.10 DPS, sim-verified) [dungeon] |
| finger2 | Tarnished Elven Ring (18500) | Dire Maul: Tribute [dungeon] | sim-verified (786.9 DPS) | yes | Cutthroat's Signet (272408, -0.42 DPS) [vendor]; Don Julio's Band (19325, -0.49 DPS) [rep]; Naglering (11669, -7.60 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (786.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (786.9 DPS) | yes | Devilsaur Eye (19991, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.38 DPS) [crafted]; Counterattack Lodestone (18537, -4.05 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (786.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -12.55 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (786.9 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -23.92 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Tarnished Elven Ring; trinket1: Burst of Knowledge; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 102.5. Weights run: 2.5s. Verify run: 2.1s. 209 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.622 ± 0.008, crit=0.609 ± 0.016 per rating point (14 rating = 1%, 8.519 per %), hit=1.049 ± 0.042 per rating point (10 rating = 1%, 10.487 per %), melee_haste=8.371 ± 0.680

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 21.0 ranged_attack_power points (1.43 DPS) | yes | Red Winter Hat (21524, -1.46 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 15.7 ranged_attack_power points (1.07 DPS) | yes | Erudite's Amulet (277204, -0.37 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 13.1 ranged_attack_power points (0.89 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 15.7 ranged_attack_power points (1.07 DPS) | yes | Cape of the Brotherhood (5193, -0.18 DPS) [dungeon]; Hide of Lupos (3018, -0.36 DPS) [world]; Bristlebark Cape (14571, -0.36 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 18.4 ranged_attack_power points (1.25 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.18 DPS) [crafted]; Prospector's Chestpiece (14562, -0.18 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 13.1 ranged_attack_power points (0.89 DPS) | yes | Wolf Bracers (4794, -0.18 DPS) [vendor]; Bristlebark Bindings (14569, -0.36 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.36 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.36 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.36 DPS) [crafted]; Serpent Gloves (5970, -1.35 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.22 DPS) | yes | Deviate Scale Belt (6468, -0.35 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.51 DPS) [world]; Dark Leather Belt (4249, -0.51 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 23.6 ranged_attack_power points (1.60 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.18 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Agile Boots (4788, -0.18 DPS) [vendor]; Feet of the Lynx (1121, -1.07 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 15.7 ranged_attack_power points (1.07 DPS) | yes | Bounty Hunter's Ring (5351, -0.53 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.71 DPS) [dungeon]; The 1 Ring (8350, -0.89 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.5 ranged_attack_power points (0.71 DPS) | yes | Bounty Hunter's Ring (5351, -0.18 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon]; The 1 Ring (8350, -0.53 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 27.0 ranged_attack_power points (1.83 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.59 DPS) [world]; Crescent Staff (6505, -0.59 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.9 ranged_attack_power points (12.16 DPS) | yes | Outrider's Bow (20437, -0.70 DPS) [pvp]; Lil Timmy's Peashooter (13136, -1.66 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.64 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 138.6. Weights run: 2.8s. Verify run: 2.7s. 352 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.649 ± 0.010, crit=0.710 ± 0.019 per rating point (14 rating = 1%, 9.944 per %), hit=1.227 ± 0.054 per rating point (10 rating = 1%, 12.269 per %), melee_haste=9.628 ± 0.925

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 26.5 ranged_attack_power points (1.79 DPS) | yes | Tribal Worg Helm (6204, -0.36 DPS) [world]; Brawler's Leather Hood (252504, -0.36 DPS) [crafted]; Humbert's Helm (4724, -0.54 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 21.2 ranged_attack_power points (1.43 DPS) | yes | Ghostshard Talisman (7731, -0.48 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.71 DPS) [world_drop]; Erudite's Amulet (277204, -0.71 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.1 ranged_attack_power points (1.96 DPS) | yes | Mantle of Thieves (2264, -0.18 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.71 DPS) [crafted]; Insignia Mantle (4721, -0.71 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 21.2 ranged_attack_power points (1.43 DPS) | yes | Hawkeye's Cloak (14593, -0.18 DPS) [world_drop]; Cloak of Night (4447, -0.36 DPS) [world]; Swiftrunner Cape (6745, -0.36 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 37.1 ranged_attack_power points (2.50 DPS) | yes | Panther Armor (6670, -0.92 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.07 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.07 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 15.9 ranged_attack_power points (1.07 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Loamflake Bracers (15462, -0.18 DPS) [quest] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 21.2 ranged_attack_power points (1.43 DPS) | yes | Braced Handguards (6784, -0.18 DPS) [quest]; Heavy Earthen Gloves (7359, -0.35 DPS) [crafted]; Insignia Gloves (6408, -0.36 DPS) [world_drop] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 ranged_attack_power points (1.62 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.01 DPS) [crafted]; Stalker's Leather Belt (252521, -0.01 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 37.1 ranged_attack_power points (2.50 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.18 DPS) [crafted]; Ferine Leggings (6690, -0.75 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 21.2 ranged_attack_power points (1.43 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.8 ranged_attack_power points (1.61 DPS) | yes | Ring of Precision (1491, -0.54 DPS) [dungeon]; Legionnaire's Band (19513, -0.54 DPS) [rep]; Signet of the Zhevra (285330, -0.54 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 18.5 ranged_attack_power points (1.25 DPS) | yes | Legionnaire's Band (19513, -0.18 DPS) [rep]; Signet of the Zhevra (285330, -0.18 DPS) [world]; Ring of Precision (1491, -0.87 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 19.2 ranged_attack_power points (1.30 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 ranged_attack_power points (1.07 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.63 DPS) | yes | Silver Star (3463, -0.58 DPS, sim-verified) [quest]; Glass Shooter (9456, -0.71 DPS) [dungeon]; Ironweaver (13137, -1.30 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Defiler's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 188.6. Weights run: 3.0s. Verify run: 2.9s. 563 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.656 ± 0.010, crit=0.770 ± 0.020 per rating point (14 rating = 1%, 10.776 per %), hit=1.263 ± 0.100 per rating point (10 rating = 1%, 12.633 per %), melee_haste=9.750 ± 1.505

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 34.5 ranged_attack_power points (2.26 DPS) | yes | Guard's Chain Helm (250499, -0.17 DPS) [crafted]; Skullsplitter Helm (1624, -0.35 DPS) [world]; Nightscape Headband (8176, -0.93 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 29.2 ranged_attack_power points (1.91 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.60 DPS) [quest]; Ghostshard Talisman (7731, -1.00 DPS) [dungeon]; Erudite's Amulet (277204, -1.22 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 41.2 ranged_attack_power points (2.70 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.81 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.96 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 26.6 ranged_attack_power points (1.74 DPS) | yes | Imperial Cloak (6432, -0.35 DPS) [dungeon]; Parachute Cloak (10518, -0.35 DPS) [crafted]; Tigerstrike Mantle (13108, -0.35 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 50.5 ranged_attack_power points (3.31 DPS) | yes | Wolffear Harness (13110, -0.69 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.70 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.70 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 21.2 ranged_attack_power points (1.39 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.10 DPS) | yes | Gloves of Holy Might (867, -0.08 DPS) [world_drop]; Dragonscale Gauntlets (8347, -0.35 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.36 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.97 DPS) | yes | Scorpashi Sash (14652, -0.23 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.39 DPS) [rep]; Blackforge Girdle (6425, -0.40 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 55.8 ranged_attack_power points (3.65 DPS) | yes | Triprunner Dungarees (9624, -0.72 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.22 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.22 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 34.5 ranged_attack_power points (2.26 DPS) | yes | Imperial Leather Boots (6431, -0.35 DPS) [dungeon]; Dusky Boots (7390, -0.35 DPS) [crafted]; Worn Running Boots (9398, -0.35 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 26.6 ranged_attack_power points (1.74 DPS) | yes | Ironspine's Eye (7686, -0.17 DPS) [dungeon]; Legionnaire's Band (19512, -0.35 DPS) [rep]; Disengagement Ring (276202, -0.35 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 23.9 ranged_attack_power points (1.57 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Legionnaire's Band (19512, -0.17 DPS) [rep]; Disengagement Ring (276202, -0.17 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (188.6 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.89 DPS, sim-verified) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 22.0 ranged_attack_power points (1.44 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Satyr's Rod (15962, -1.27 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (22.93 DPS) | yes | Outrider's Bow (19560, -1.58 DPS) [pvp]; Shadowforge Bushmaster (9422, -2.05 DPS) [dungeon]; The Silencer (13138, -3.58 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Vanquisher's Sword; ranged: Bow of Searing Arrows

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 239.8. Weights run: 3.1s. Verify run: 3.5s. 713 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.712 ± 0.012, crit=0.921 ± 0.023 per rating point (14 rating = 1%, 12.900 per %), hit=1.697 ± 0.099 per rating point (10 rating = 1%, 16.969 per %), melee_haste=13.372 ± 1.600

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 46.1 ranged_attack_power points (3.14 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.37 DPS) [quest]; Tough Scorpid Helm (8208, -0.55 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.3 ranged_attack_power points (2.40 DPS) | yes | Scout's Medallion (19536, -0.37 DPS) [rep]; Woven Ivy Necklace (19159, -0.74 DPS) [quest] |
| shoulder | Phytoskin Spaulders (17749) | Maraudon: Razorlash [dungeon] | 43.4 ranged_attack_power points (2.96 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Sunburn Spaulders (274751, -0.11 DPS) [vendor]; Khan's Mantle (14787, -0.74 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 40.7 ranged_attack_power points (2.77 DPS) | yes | Blackveil Cape (11626, -0.18 DPS) [dungeon]; Blackmetal Cape (9512, -0.55 DPS) [dungeon]; Blisterbane Wrap (12552, -1.12 DPS, sim-verified) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 67.8 ranged_attack_power points (4.62 DPS) | yes | Blazewind Breastplate (11193, -0.37 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.97 DPS) [vendor]; Quillward Harness (10583, -1.11 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 40.7 ranged_attack_power points (2.77 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Bloodlust Bracelets (14807, -0.91 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 59.7 ranged_attack_power points (4.06 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -1.46 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.48 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.48 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 53.7 ranged_attack_power points (3.66 DPS) | yes | Sagebrush Girdle (17778, -0.89 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.07 DPS) [crafted]; Stalker's Mail Belt (252588, -1.07 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 56.9 ranged_attack_power points (3.88 DPS) | yes | Infernal Trickster Leggings (17754, -0.18 DPS) [dungeon]; Keeper's Woolies (14668, -0.37 DPS) [world_drop]; Stone Guard's Chain Legplates (220833, -0.41 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 54.2 ranged_attack_power points (3.70 DPS) | yes | Fleetfoot Greaves (11627, -0.18 DPS) [dungeon]; Elven Chain Boots (13125, -0.37 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.55 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.0 ranged_attack_power points (2.59 DPS) | yes | Ring of the Underwood (2951, -0.74 DPS) [world_drop]; Falcon's Hook (7552, -0.92 DPS) [dungeon]; Ironspine's Eye (7686, -0.92 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 37.0 ranged_attack_power points (2.52 DPS) | yes | Ring of the Underwood (2951, -0.67 DPS) [world_drop]; Falcon's Hook (7552, -0.86 DPS) [dungeon]; Ironspine's Eye (7686, -0.86 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (239.8 DPS) | yes | Frozen Heart of the Mountain (249469, -5.49 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (239.8 DPS) | yes | Frozen Heart of the Mountain (249469, -1.48 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (239.8 DPS) | yes | Warmonger (13052, -0.78 DPS) [world_drop]; Steel Spear (250605, -1.17 DPS) [crafted]; Hanzo Sword (8190, -4.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (239.8 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.59 DPS) [world_drop]; Hurricane (2824, -1.66 DPS) [world_drop]; Dark Iron Rifle (16004, -3.55 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Phytoskin Spaulders; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 319.9. Weights run: 3.0s. Verify run: 11.4s. 1650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.950 ± 0.021, crit=1.691 ± 0.043 per rating point (14 rating = 1%, 23.668 per %), hit=2.934 ± 0.157 per rating point (10 rating = 1%, 29.343 per %), melee_haste=17.485 ± 2.316

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (319.9 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -9.87 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 73.6 ranged_attack_power points (5.01 DPS) | yes | Beads of Ogre Might (22150, -1.38 DPS) [quest]; Mark of Fordring (15411, -1.63 DPS) [quest]; Amulet of the Darkmoon (19491, -1.90 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 103.3 ranged_attack_power points (7.03 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -0.52 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 64.2 ranged_attack_power points (4.37 DPS) | yes | Shifting Cloak (18511, -0.96 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.96 DPS) [dungeon]; Howler's Furs (272414, -1.15 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (319.9 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.60 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.60 DPS) [vendor]; Tunic of Undead Slaying (23089, -13.89 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (319.9 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -10.29 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (319.9 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (319.9 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -8.02 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 179.9 ranged_attack_power points (12.24 DPS) | yes | Outrider's Chain Leggings (22673, -1.91 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -3.60 DPS) [vendor]; General's Chain Legguards (231574, -3.81 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (319.9 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -9.20 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (319.9 DPS) | yes | Cutthroat's Signet (272408, -0.80 DPS) [vendor]; Don Julio's Band (19325, -0.91 DPS) [rep]; Naglering (11669, -7.32 DPS, sim-verified) [dungeon] |
| finger2 | Tarnished Elven Ring (18500) | Dire Maul: Tribute [dungeon] | sim-verified (319.9 DPS) | yes | Cutthroat's Signet (272408, -0.20 DPS) [vendor]; Don Julio's Band (19325, -0.31 DPS) [rep]; Naglering (11669, -6.60 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (319.9 DPS) | yes | Blackhand's Breadth (13965, -3.89 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.32 DPS) [crafted]; Counterattack Lodestone (18537, -5.62 DPS) [dungeon] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (319.9 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS, sim-verified) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (319.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -8.99 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (319.9 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -12.85 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Tarnished Elven Ring; trinket2: Second Wind; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 799.2. Weights run: 3.1s. Verify run: 12.1s. 1650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.995 ± 0.020, crit=1.815 ± 0.041 per rating point (14 rating = 1%, 25.414 per %), hit=3.766 ± 0.217 per rating point (10 rating = 1%, 37.656 per %), melee_haste=20.998 ± 2.757

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (799.2 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -12.17 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 82.6 ranged_attack_power points (11.61 DPS) | yes | Beads of Ogre Might (22150, -3.35 DPS, sim-verified) [quest]; Amulet of the Darkmoon (19491, -3.61 DPS) [quest]; Mark of Fordring (15411, -4.38 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 106.3 ranged_attack_power points (14.95 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -0.24 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 65.7 ranged_attack_power points (9.23 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Shifting Cloak (18511, -2.07 DPS) [crafted]; Shadow Prowler's Cloak (22269, -2.07 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (799.2 DPS) | yes | Warlord's Chain Chestpiece (16565, -6.56 DPS) [vendor]; Warlord's Chain Hauberk (231573, -6.56 DPS) [vendor]; Tunic of Undead Slaying (23089, -22.37 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (799.2 DPS) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -15.91 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | 102.3 ranged_attack_power points (14.38 DPS) | yes | Gauntlets of Accuracy (18349, +0.00 DPS, sim-verified) [dungeon]; General's Chain Gloves (16571, -1.96 DPS) [vendor]; General's Chain Vices (231575, -1.96 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (799.2 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Marksman's Girdle (22232, -11.33 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 193.3 ranged_attack_power points (27.18 DPS) | yes | Outrider's Chain Leggings (22673, -3.63 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -8.66 DPS) [vendor]; Plaguehound Leggings (18736, -9.25 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (799.2 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -12.48 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (799.2 DPS) | yes | Cutthroat's Signet (272408, -1.68 DPS) [vendor]; Don Julio's Band (19325, -1.76 DPS) [rep]; Naglering (11669, -9.04 DPS, sim-verified) [dungeon] |
| finger2 | Tarnished Elven Ring (18500) | Dire Maul: Tribute [dungeon] | sim-verified (799.2 DPS) | yes | Cutthroat's Signet (272408, -0.42 DPS) [vendor]; Don Julio's Band (19325, -0.49 DPS) [rep]; Naglering (11669, -7.55 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (799.2 DPS) | yes | Frozen Heart of the Mountain (249469, -10.75 DPS) [crafted]; Counterattack Lodestone (18537, -12.43 DPS) [dungeon]; Hand of Justice (11815, -12.71 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (799.2 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.38 DPS) [crafted]; Counterattack Lodestone (18537, -4.05 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (799.2 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -12.46 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (799.2 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -25.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Tarnished Elven Ring; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

