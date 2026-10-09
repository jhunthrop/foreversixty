# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 107.5. Weights run: 1.6s. Verify run: 1.4s. 221 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.073 ± 0.014, crit=0.589 ± 0.016 per rating point (14 rating = 1%, 8.252 per %), hit=0.952 ± 0.035 per rating point (10 rating = 1%, 9.522 per %), melee_haste=5.285 ± 0.724

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.6 ranged_attack_power points (1.68 DPS) | yes | Red Winter Hat (21524, -1.73 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.4 ranged_attack_power points (1.26 DPS) | yes | Erudite's Amulet (277204, -0.85 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.4 ranged_attack_power points (1.05 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.4 ranged_attack_power points (1.26 DPS) | yes | Hide of Lupos (3018, -0.42 DPS) [world]; Bristlebark Cape (14571, -0.42 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.61 DPS, sim-verified) [dungeon] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 33.8 ranged_attack_power points (2.31 DPS) | yes | Brawler's Leather Armor (252490, -0.84 DPS) [crafted]; Trapper's Leather Armor (252491, -0.84 DPS) [crafted]; Dark Leather Tunic (2317, -1.05 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.4 ranged_attack_power points (1.05 DPS) | yes | Wolf Bracers (4794, -0.21 DPS) [vendor]; Bravo's Armbands (270015, -0.21 DPS) [quest]; Ratchet Wristwraps (274742, -0.42 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.42 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.42 DPS) [crafted]; Serpent Gloves (5970, -1.44 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.23 DPS) | yes | Deviate Scale Belt (6468, -0.18 DPS) [crafted]; Dusty Belt (279897, -0.18 DPS) [quest]; Dark Leather Belt (4249, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.7 ranged_attack_power points (1.89 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.21 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Agile Boots (4788, -0.21 DPS) [vendor]; Feet of the Lynx (1121, -1.10 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 18.4 ranged_attack_power points (1.26 DPS) | yes | Lavishly Jeweled Ring (1156, -0.84 DPS) [dungeon]; The 1 Ring (8350, -1.05 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.3 ranged_attack_power points (0.84 DPS) | yes | Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; The 1 Ring (8350, -0.63 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.6 ranged_attack_power points (2.09 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.62 DPS) [world]; Lupine Axe (1220, -0.83 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.3 ranged_attack_power points (12.24 DPS) | yes | Lil Timmy's Peashooter (13136, -0.91 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.69 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.95 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 221, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 132.6. Weights run: 1.6s. Verify run: 1.9s. 395 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.027 ± 0.016, crit=0.864 ± 0.023 per rating point (14 rating = 1%, 12.090 per %), hit=1.133 ± 0.053 per rating point (10 rating = 1%, 11.328 per %), melee_haste=8.507 ± 1.048

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 30.3 ranged_attack_power points (2.03 DPS) | yes | Tribal Worg Helm (6204, -0.41 DPS) [world]; Brawler's Leather Hood (252504, -0.44 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.61 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Ghostshard Talisman (7731, -0.69 DPS) [dungeon]; Wolfpack Medallion (5754, -0.81 DPS) [world]; Kaleidoscope Chain (13084, -0.81 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.3 ranged_attack_power points (2.24 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.81 DPS) [crafted]; Insignia Mantle (4721, -0.81 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Cloak of Night (4447, -0.41 DPS) [world]; Fenrus' Hide (6340, -0.41 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 42.4 ranged_attack_power points (2.85 DPS) | yes | Tunic of Westfall (2041, -0.66 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.22 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.22 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Serpent Gloves (5970, -0.41 DPS) [dungeon]; Gloves of the Fang (10413, -0.41 DPS) [dungeon]; Insignia Gloves (6408, -0.44 DPS, sim-verified) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (132.6 DPS) | yes | Highlander's Chain Girdle (20090, -0.22 DPS) [rep]; Highlander's Leather Girdle (20117, -0.22 DPS) [rep]; Skulker's Leather Belt (252520, -1.44 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 42.4 ranged_attack_power points (2.85 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -1.02 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.20 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 27.2 ranged_attack_power points (1.83 DPS) | yes | Ring of Precision (1491, -0.61 DPS) [dungeon]; Protector's Band (19517, -0.61 DPS) [rep]; Signet of the Zhevra (285330, -0.61 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 21.2 ranged_attack_power points (1.42 DPS) | yes | Protector's Band (19517, -0.20 DPS) [rep]; Signet of the Zhevra (285330, -0.20 DPS) [world]; Ring of Precision (1491, -0.84 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 21.1 ranged_attack_power points (1.42 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 263.0 ranged_attack_power points (17.68 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.10 DPS) [vendor]; Golemsight Long Gun (273029, -0.84 DPS) [dungeon]; Silver Star (3463, -1.82 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Glass Shooter

No-known-source sample (15 of 395, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 169.6. Weights run: 1.8s. Verify run: 2.0s. 628 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.970 ± 0.017, crit=1.067 ± 0.028 per rating point (14 rating = 1%, 14.943 per %), hit=1.335 ± 0.069 per rating point (10 rating = 1%, 13.355 per %), melee_haste=9.897 ± 1.033

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 38.6 ranged_attack_power points (2.77 DPS) | yes | Guard's Chain Helm (250499, -0.21 DPS) [crafted]; Skullsplitter Helm (1624, -0.43 DPS) [world]; Nightscape Headband (8176, -1.10 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 32.7 ranged_attack_power points (2.34 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.91 DPS) [quest]; Ghostshard Talisman (7731, -1.34 DPS) [dungeon]; Wolfpack Medallion (5754, -1.49 DPS) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 44.7 ranged_attack_power points (3.20 DPS) | yes | Forest Tracker Epaulets (2278, -0.86 DPS) [world_drop]; Nightscape Shoulders (8192, -0.92 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -1.07 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 29.7 ranged_attack_power points (2.13 DPS) | yes | Imperial Cloak (6432, -0.43 DPS) [dungeon]; Parachute Cloak (10518, -0.43 DPS) [crafted]; Tigerstrike Mantle (13108, -0.43 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (169.6 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.43 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.43 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 23.8 ranged_attack_power points (1.70 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.21 DPS) [crafted]; Duracin Bracers (10358, -0.21 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 34.9 ranged_attack_power points (2.50 DPS) | yes | Dragonscale Gauntlets (8347, -0.16 DPS) [crafted]; Gauntlets of Divinity (7724, -0.21 DPS) [dungeon]; Tough Scorpid Gloves (8204, -0.38 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (2.15 DPS) | yes | Scorpashi Sash (14652, -0.02 DPS) [world_drop]; Blackforge Girdle (6425, -0.23 DPS) [dungeon]; Ogron's Sash (13117, -0.23 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 62.4 ranged_attack_power points (4.47 DPS) | yes | Petrolspill Leggings (9509, -1.49 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.49 DPS) [world_drop]; Triprunner Dungarees (9624, -1.49 DPS, sim-verified) [quest] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 38.6 ranged_attack_power points (2.77 DPS) | yes | Imperial Leather Boots (6431, -0.43 DPS) [dungeon]; Dusky Boots (7390, -0.43 DPS) [crafted]; Worn Running Boots (9398, -0.43 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 29.7 ranged_attack_power points (2.13 DPS) | yes | Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Protector's Band (19515, -0.43 DPS) [rep]; Disengagement Ring (276202, -0.43 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 26.7 ranged_attack_power points (1.92 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Protector's Band (19515, -0.21 DPS) [rep]; Disengagement Ring (276202, -0.21 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.72 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Ginn-su Sword (9424), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 23.8 ranged_attack_power points (1.70 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Ginn-su Sword (9424, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (25.08 DPS) | yes | Shadowforge Bushmaster (9422, -2.25 DPS) [dungeon]; Swiftwind (13038, -2.27 DPS) [world_drop]; The Silencer (13138, -3.42 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 628, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 231.2. Weights run: 1.8s. Verify run: 4.4s. 787 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.925 ± 0.018, crit=1.248 ± 0.032 per rating point (14 rating = 1%, 17.474 per %), hit=1.503 ± 0.086 per rating point (10 rating = 1%, 15.032 per %), melee_haste=12.277 ± 1.297

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 64.3 ranged_attack_power points (4.96 DPS) | yes | Lordrec Helmet (10741, -1.35 DPS) [quest]; Sprightring Helm (17776, -1.58 DPS) [quest]; Helm of Fire (8348, -5.97 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 38.0 ranged_attack_power points (2.94 DPS) | yes | Sentinel's Medallion (19539, -0.23 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 55.5 ranged_attack_power points (4.29 DPS) | yes | Sunburn Spaulders (274751, -0.87 DPS) [vendor]; Khan's Mantle (14787, -1.58 DPS) [world_drop]; Phytoskin Spaulders (17749, -3.64 DPS, sim-verified) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 43.9 ranged_attack_power points (3.39 DPS) | yes | Blackveil Cape (11626, -0.23 DPS) [dungeon]; Blackmetal Cape (9512, -0.68 DPS) [dungeon]; Blisterbane Wrap (12552, -0.87 DPS, sim-verified) [dungeon] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blazewind Breastplate (11193, +0.00 DPS) [quest]; Quillward Harness (10583, -0.45 DPS) [dungeon]; Fungus Shroud Armor (17742, -3.59 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 43.9 ranged_attack_power points (3.39 DPS) | yes | Wicked Leather Bracers (15084, -0.90 DPS) [crafted]; Bloodlust Bracelets (14807, -1.00 DPS, sim-verified) [world_drop]; Pridelord Bands (14672, -1.13 DPS) [world_drop] |
| hands | Sergeant Major's Chain Gauntlets (220829) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beastmaster's Gauntlets (226883, -0.45 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -0.45 DPS) [crafted]; Raider Gloves (272100, -0.99 DPS, sim-verified) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 56.9 ranged_attack_power points (4.39 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.23 DPS) [crafted]; Stalker's Mail Belt (252588, -1.23 DPS) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Infernal Trickster Leggings (17754, +0.00 DPS) [dungeon]; Keeper's Woolies (14668, -0.22 DPS) [world_drop]; Basilisk Hide Pants (1718, -3.73 DPS, sim-verified) [world_drop] |
| feet | Sergeant Major's Chain Sabatons (220837) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fleetfoot Greaves (11627, +0.00 DPS) [dungeon]; Elven Chain Boots (13125, +0.00 DPS) [world_drop]; Albino Crocscale Boots (17728, -1.25 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 40.9 ranged_attack_power points (3.16 DPS) | yes | Falcon's Hook (7552, -1.13 DPS) [dungeon]; Ironspine's Eye (7686, -1.13 DPS) [dungeon]; Protector's Band (19516, -1.13 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.2 ranged_attack_power points (2.26 DPS) | yes | Falcon's Hook (7552, -0.23 DPS) [dungeon]; Ironspine's Eye (7686, -0.23 DPS) [dungeon]; Protector's Band (19516, -0.23 DPS) [rep] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -3.84 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (231.2 DPS) | yes | - |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Steel Spear (250605, -1.49 DPS) [crafted]; Warmonger (13052, -1.71 DPS) [world_drop]; Hanzo Sword (8190, -6.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.57 DPS) [world_drop]; Hurricane (2824, -1.88 DPS) [world_drop]; Dark Iron Rifle (16004, -3.90 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; hands: Sergeant Major's Chain Gauntlets; waist: Substandard Belt Chain; legs: Knight's Chain Legplates; feet: Sergeant Major's Chain Sabatons; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 787, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 303.9. Weights run: 1.8s. Verify run: 8.4s. 1685 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.106 ± 0.026, crit=2.176 ± 0.054 per rating point (14 rating = 1%, 30.468 per %), hit=2.977 ± 0.141 per rating point (10 rating = 1%, 29.769 per %), melee_haste=18.361 ± 2.067

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -11.74 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -0.19 DPS) [quest]; Medallion of the Dawn (22659, -0.35 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 114.3 ranged_attack_power points (8.72 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Field Marshal's Chain Spaulders (16468, -0.29 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -0.29 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.6 ranged_attack_power points (5.08 DPS) | yes | Howler's Furs (272414, -0.67 DPS) [vendor]; Shifting Cloak (18511, -1.05 DPS) [crafted]; Shadow Prowler's Cloak (22269, -1.05 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Breastplate (16466, -2.98 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.98 DPS) [vendor]; Tunic of Undead Slaying (23089, -16.13 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -9.37 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS, sim-verified) [quest]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -9.28 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 199.4 ranged_attack_power points (15.22 DPS) | yes | Sentinel's Leather Pants (237818, -4.17 DPS) [vendor]; Marshal's Chain Legguards (231577, -4.83 DPS) [pvp]; Plaguehound Leggings (18736, -5.84 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -10.11 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.13 DPS) [dungeon]; Cutthroat's Signet (272408, -2.37 DPS) [vendor]; Naglering (11669, -9.72 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, -0.23 DPS) [vendor]; Innervating Band (18701, -2.26 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (303.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -0.37 DPS) [dungeon]; Hand of Justice (11815, -0.52 DPS) [dungeon] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -9.78 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.91 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Burst of Knowledge; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 766.2. Weights run: 1.8s. Verify run: 8.6s. 1685 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.125 ± 0.024, crit=2.277 ± 0.051 per rating point (14 rating = 1%, 31.876 per %), hit=4.195 ± 0.204 per rating point (10 rating = 1%, 41.945 per %), melee_haste=17.263 ± 2.660

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -13.71 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Amulet of the Darkmoon (19491, -1.10 DPS) [quest]; Mark of Fordring (15411, -1.35 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 116.3 ranged_attack_power points (19.51 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shifting Cloak (18511, -2.31 DPS) [crafted]; Shadow Prowler's Cloak (22269, -2.31 DPS) [dungeon]; Howler's Furs (272414, -8.13 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Breastplate (16466, -8.61 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -8.61 DPS) [vendor]; Tunic of Undead Slaying (23089, -33.61 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -9.97 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Gauntlets of Accuracy (18349, +0.00 DPS) [dungeon]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warpwood Binding (18393, +0.00 DPS) [dungeon]; Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, -5.67 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 215.1 ranged_attack_power points (36.09 DPS) | yes | Sentinel's Leather Pants (237818, -11.23 DPS) [vendor]; Marshal's Chain Legguards (231577, -12.91 DPS) [pvp]; Plaguehound Leggings (18736, -13.32 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -10.50 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -4.72 DPS) [dungeon]; Cutthroat's Signet (272408, -5.24 DPS) [vendor]; Naglering (11669, -20.01 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.17 DPS) [dungeon]; Cutthroat's Signet (272408, -0.69 DPS) [vendor]; Naglering (11669, -13.50 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (766.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -2.64 DPS) [dungeon]; Hand of Justice (11815, -2.98 DPS) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -2.84 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -19.96 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -24.08 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 108.1. Weights run: 1.6s. Verify run: 1.4s. 210 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.073 ± 0.014, crit=0.589 ± 0.016 per rating point (14 rating = 1%, 8.252 per %), hit=0.952 ± 0.035 per rating point (10 rating = 1%, 9.522 per %), melee_haste=5.285 ± 0.724

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.6 ranged_attack_power points (1.68 DPS) | yes | Red Winter Hat (21524, -1.78 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.4 ranged_attack_power points (1.26 DPS) | yes | Erudite's Amulet (277204, -0.90 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.4 ranged_attack_power points (1.05 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.4 ranged_attack_power points (1.26 DPS) | yes | Cape of the Brotherhood (5193, -0.21 DPS) [dungeon]; Hide of Lupos (3018, -0.42 DPS) [world]; Bristlebark Cape (14571, -0.42 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Dark Leather Tunic (2317, -0.21 DPS) [crafted]; Prospector's Chestpiece (14562, -0.21 DPS) [world_drop]; Brawler's Leather Armor (252490, -1.08 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.4 ranged_attack_power points (1.05 DPS) | yes | Wolf Bracers (4794, -0.21 DPS) [vendor]; Bristlebark Bindings (14569, -0.42 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.42 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.42 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.42 DPS) [crafted]; Serpent Gloves (5970, -1.59 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.23 DPS) | yes | Deviate Scale Belt (6468, -0.18 DPS) [crafted]; Guardsman Belt (3429, -0.39 DPS) [world]; Dark Leather Belt (4249, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.7 ranged_attack_power points (1.89 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.21 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Agile Boots (4788, -0.21 DPS) [vendor]; Feet of the Lynx (1121, -1.22 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 18.4 ranged_attack_power points (1.26 DPS) | yes | Bounty Hunter's Ring (5351, -0.63 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.84 DPS) [dungeon]; The 1 Ring (8350, -1.05 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.3 ranged_attack_power points (0.84 DPS) | yes | Bounty Hunter's Ring (5351, -0.21 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; The 1 Ring (8350, -0.63 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.6 ranged_attack_power points (2.09 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.62 DPS) [world]; Crescent Staff (6505, -0.62 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.3 ranged_attack_power points (12.24 DPS) | yes | Outrider's Bow (20437, -0.67 DPS) [pvp]; Lil Timmy's Peashooter (13136, -0.72 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.69 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 210, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 136.3. Weights run: 1.6s. Verify run: 1.8s. 380 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.027 ± 0.016, crit=0.864 ± 0.023 per rating point (14 rating = 1%, 12.090 per %), hit=1.133 ± 0.053 per rating point (10 rating = 1%, 11.328 per %), melee_haste=8.507 ± 1.048

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 30.3 ranged_attack_power points (2.03 DPS) | yes | Tribal Worg Helm (6204, -0.41 DPS) [world]; Brawler's Leather Hood (252504, -0.48 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.61 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Ghostshard Talisman (7731, -0.69 DPS) [dungeon]; Wolfpack Medallion (5754, -0.81 DPS) [world]; Kaleidoscope Chain (13084, -0.81 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.3 ranged_attack_power points (2.24 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.81 DPS) [crafted]; Insignia Mantle (4721, -0.81 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Cloak of Night (4447, -0.41 DPS) [world]; Swiftrunner Cape (6745, -0.41 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 42.4 ranged_attack_power points (2.85 DPS) | yes | Panther Armor (6670, -1.21 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.22 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.22 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Loamflake Bracers (15462, -0.20 DPS) [quest] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Braced Handguards (6784, -0.20 DPS) [quest]; Serpent Gloves (5970, -0.41 DPS) [dungeon]; Insignia Gloves (6408, -0.41 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (136.3 DPS) | yes | Deftkin Belt (16659, -0.21 DPS) [quest]; Defiler's Chain Girdle (20152, -0.22 DPS) [rep]; Skulker's Leather Belt (252520, -2.15 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 42.4 ranged_attack_power points (2.85 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -1.02 DPS) [world_drop] |
| feet | Vorrel's Boots (7751) | Vorrel's Revenge [quest] | 30.3 ranged_attack_power points (2.03 DPS) | yes | Warsong Boots (16977, -0.41 DPS) [quest]; Highlander's Mail Greaves (20123, -0.41 DPS) [vendor]; Insignia Boots (4055, -1.01 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 27.2 ranged_attack_power points (1.83 DPS) | yes | Ring of Precision (1491, -0.61 DPS) [dungeon]; Legionnaire's Band (19513, -0.61 DPS) [rep]; Signet of the Zhevra (285330, -0.61 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 21.2 ranged_attack_power points (1.42 DPS) | yes | Ring of Precision (1491, -0.20 DPS) [dungeon]; Legionnaire's Band (19513, -0.20 DPS) [rep]; Signet of the Zhevra (285330, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | 25.1 ranged_attack_power points (1.69 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 21.1 ranged_attack_power points (1.42 DPS) | yes | Vendetta (776, +0.00 DPS) [dungeon]; Satyr's Rod (15962, -1.22 DPS) [world_drop] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 263.0 ranged_attack_power points (17.68 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.10 DPS) [vendor]; Golemsight Long Gun (273029, -0.84 DPS) [dungeon]; Silver Star (3463, -1.64 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Vorrel's Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Serrated Raptor Claw; off_hand: Alliance Outrunner's Sword; ranged: Glass Shooter

No-known-source sample (15 of 380, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 171.5. Weights run: 1.8s. Verify run: 2.0s. 593 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.970 ± 0.017, crit=1.067 ± 0.028 per rating point (14 rating = 1%, 14.943 per %), hit=1.335 ± 0.069 per rating point (10 rating = 1%, 13.355 per %), melee_haste=9.897 ± 1.033

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 38.6 ranged_attack_power points (2.77 DPS) | yes | Guard's Chain Helm (250499, -0.21 DPS) [crafted]; Skullsplitter Helm (1624, -0.43 DPS) [world]; Nightscape Headband (8176, -1.12 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 32.7 ranged_attack_power points (2.34 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.91 DPS) [quest]; Ghostshard Talisman (7731, -1.34 DPS) [dungeon]; Wolfpack Medallion (5754, -1.49 DPS) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 44.7 ranged_attack_power points (3.20 DPS) | yes | Forest Tracker Epaulets (2278, -0.86 DPS) [world_drop]; Nightscape Shoulders (8192, -0.94 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -1.07 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 29.7 ranged_attack_power points (2.13 DPS) | yes | Imperial Cloak (6432, -0.43 DPS) [dungeon]; Parachute Cloak (10518, -0.43 DPS) [crafted]; Tigerstrike Mantle (13108, -0.43 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (171.5 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.43 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.43 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 23.8 ranged_attack_power points (1.70 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 34.9 ranged_attack_power points (2.50 DPS) | yes | Dragonscale Gauntlets (8347, -0.16 DPS) [crafted]; Gauntlets of Divinity (7724, -0.21 DPS) [dungeon]; Tough Scorpid Gloves (8204, -0.38 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (2.15 DPS) | yes | Scorpashi Sash (14652, -0.02 DPS) [world_drop]; Blackforge Girdle (6425, -0.23 DPS) [dungeon]; Ogron's Sash (13117, -0.23 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 62.4 ranged_attack_power points (4.47 DPS) | yes | Triprunner Dungarees (9624, -1.14 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.49 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.49 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 38.6 ranged_attack_power points (2.77 DPS) | yes | Imperial Leather Boots (6431, -0.43 DPS) [dungeon]; Dusky Boots (7390, -0.43 DPS) [crafted]; Worn Running Boots (9398, -0.43 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 29.7 ranged_attack_power points (2.13 DPS) | yes | Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Legionnaire's Band (19512, -0.43 DPS) [rep]; Disengagement Ring (276202, -0.43 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 26.7 ranged_attack_power points (1.92 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Legionnaire's Band (19512, -0.21 DPS) [rep]; Disengagement Ring (276202, -0.21 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -2.23 DPS, sim-verified) [world_drop] |
| off_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | 24.8 ranged_attack_power points (1.78 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Satyr's Rod (15962, -1.57 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (25.08 DPS) | yes | Outrider's Bow (19560, -1.66 DPS) [pvp]; Shadowforge Bushmaster (9422, -2.25 DPS) [dungeon]; The Silencer (13138, -3.80 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Serrated Raptor Claw; ranged: Bow of Searing Arrows

No-known-source sample (15 of 593, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 231.8. Weights run: 1.8s. Verify run: 2.7s. 745 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.925 ± 0.018, crit=1.248 ± 0.032 per rating point (14 rating = 1%, 17.474 per %), hit=1.503 ± 0.086 per rating point (10 rating = 1%, 15.032 per %), melee_haste=12.277 ± 1.297

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 49.7 ranged_attack_power points (3.84 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.45 DPS) [quest]; Tough Scorpid Helm (8208, -0.68 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 38.0 ranged_attack_power points (2.94 DPS) | yes | Scout's Medallion (19535, -0.23 DPS) [rep]; Woven Ivy Necklace (19159, -0.90 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phytoskin Spaulders (17749, +0.00 DPS) [dungeon]; Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Khan's Mantle (14787, -0.70 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 43.9 ranged_attack_power points (3.39 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.23 DPS) [dungeon]; Blackmetal Cape (9512, -0.68 DPS) [dungeon] |
| chest | Blazewind Breastplate (11193) | Broken Alliances [quest] | sim-verified (231.8 DPS) | yes | Fungus Shroud Armor (17742, +0.00 DPS) [dungeon]; Stone Guard's Chain Armor (220827, -0.46 DPS) [vendor]; Quillward Harness (10583, -0.90 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 43.9 ranged_attack_power points (3.39 DPS) | yes | Bloodlust Bracelets (14807, +0.00 DPS) [world_drop]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 64.3 ranged_attack_power points (4.97 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -1.59 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.81 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.81 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 56.9 ranged_attack_power points (4.39 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.23 DPS) [crafted]; Stalker's Mail Belt (252588, -1.23 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 61.4 ranged_attack_power points (4.74 DPS) | yes | Infernal Trickster Leggings (17754, -0.23 DPS) [dungeon]; Stone Guard's Chain Legplates (220833, -0.23 DPS) [vendor]; Keeper's Woolies (14668, -0.45 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 58.5 ranged_attack_power points (4.52 DPS) | yes | Fleetfoot Greaves (11627, -0.23 DPS) [dungeon]; Elven Chain Boots (13125, -0.45 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.68 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 40.9 ranged_attack_power points (3.16 DPS) | yes | Falcon's Hook (7552, -1.13 DPS) [dungeon]; Ironspine's Eye (7686, -1.13 DPS) [dungeon]; Legionnaire's Band (19511, -1.13 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.2 ranged_attack_power points (2.26 DPS) | yes | Falcon's Hook (7552, -0.23 DPS) [dungeon]; Ironspine's Eye (7686, -0.23 DPS) [dungeon]; Legionnaire's Band (19511, -0.23 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -6.26 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -3.49 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Steel Spear (250605, -1.49 DPS) [crafted]; Warmonger (13052, -1.71 DPS) [world_drop]; Hanzo Sword (8190, -6.05 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.57 DPS) [world_drop]; Hurricane (2824, -1.88 DPS) [world_drop]; Dark Iron Rifle (16004, -3.42 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Blazewind Breastplate; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 745, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 309.7. Weights run: 1.8s. Verify run: 8.0s. 1664 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.106 ± 0.026, crit=2.176 ± 0.054 per rating point (14 rating = 1%, 30.468 per %), hit=2.977 ± 0.141 per rating point (10 rating = 1%, 29.769 per %), melee_haste=18.361 ± 2.067

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -10.38 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -0.19 DPS) [quest]; Medallion of the Dawn (22659, -0.35 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 114.3 ranged_attack_power points (8.72 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -0.29 DPS) [pvp]; Warlord's Chain Pauldrons (231565, -1.23 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.6 ranged_attack_power points (5.08 DPS) | yes | Howler's Furs (272414, -0.67 DPS) [vendor]; Shifting Cloak (18511, -1.05 DPS) [crafted]; Shadow Prowler's Cloak (22269, -1.05 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Chestpiece (16565, -2.98 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.98 DPS) [vendor]; Tunic of Undead Slaying (23089, -16.25 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -8.91 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -8.56 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 199.4 ranged_attack_power points (15.22 DPS) | yes | Outrider's Chain Leggings (22673, -2.63 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -4.17 DPS) [vendor]; General's Chain Legguards (231574, -4.83 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -9.47 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.13 DPS) [dungeon]; Cutthroat's Signet (272408, -2.37 DPS) [vendor]; Naglering (11669, -9.86 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, -0.23 DPS) [vendor]; Innervating Band (18701, -2.26 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, -3.35 DPS) [quest]; Counterattack Lodestone (18537, -6.32 DPS) [dungeon]; Burst of Knowledge (11832, -10.18 DPS, sim-verified) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (309.7 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -0.37 DPS) [dungeon]; Hand of Justice (11815, -0.52 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -9.95 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -12.76 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1664, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 784.5. Weights run: 1.8s. Verify run: 8.7s. 1664 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.125 ± 0.024, crit=2.277 ± 0.051 per rating point (14 rating = 1%, 31.876 per %), hit=4.195 ± 0.204 per rating point (10 rating = 1%, 41.945 per %), melee_haste=17.263 ± 2.660

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -9.49 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Amulet of the Darkmoon (19491, -1.10 DPS) [quest]; Mark of Fordring (15411, -1.35 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 116.3 ranged_attack_power points (19.51 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -0.95 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shifting Cloak (18511, -2.31 DPS) [crafted]; Shadow Prowler's Cloak (22269, -2.31 DPS) [dungeon]; Howler's Furs (272414, -8.12 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Chestpiece (16565, -8.61 DPS) [vendor]; Warlord's Chain Hauberk (231573, -8.61 DPS) [vendor]; Tunic of Undead Slaying (23089, -31.35 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -8.01 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warpwood Binding (18393, +0.00 DPS) [dungeon]; Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, -5.51 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 215.1 ranged_attack_power points (36.09 DPS) | yes | Outrider's Chain Leggings (22673, -5.99 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -11.23 DPS) [vendor]; General's Chain Legguards (231574, -12.91 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -8.01 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -4.72 DPS) [dungeon]; Cutthroat's Signet (272408, -5.24 DPS) [vendor]; Naglering (11669, -17.22 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.17 DPS) [dungeon]; Cutthroat's Signet (272408, -0.69 DPS) [vendor]; Naglering (11669, -11.19 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Counterattack Lodestone (18537, -15.33 DPS) [dungeon]; Hand of Justice (11815, -15.67 DPS) [dungeon]; Blackhand's Breadth (13965, -17.91 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (784.5 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -2.64 DPS) [dungeon]; Hand of Justice (11815, -2.98 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -17.06 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -25.19 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1664, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

