# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 110.7. Weights run: 2.8s. Verify run: 2.4s. 221 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.101 ± 0.014, crit=0.593 ± 0.016 per rating point (14 rating = 1%, 8.306 per %), hit=1.060 ± 0.050 per rating point (10 rating = 1%, 10.605 per %), melee_haste=8.276 ± 0.789

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.8 ranged_attack_power points (1.73 DPS) | yes | Red Winter Hat (21524, -1.81 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.91 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -0.91 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Hide of Lupos (3018, -0.43 DPS) [world]; Bristlebark Cape (14571, -0.43 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.64 DPS, sim-verified) [dungeon] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 34.1 ranged_attack_power points (2.37 DPS) | yes | Trapper's Leather Armor (252491, -0.86 DPS) [crafted]; Dark Leather Tunic (2317, -1.08 DPS) [crafted]; Brawler's Leather Armor (252490, -1.33 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bravo's Armbands (270015, -0.22 DPS) [quest]; Ratchet Wristwraps (274742, -0.43 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (110.7 DPS) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -1.33 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Dark Leather Belt (4249, -0.39 DPS) [crafted]; Dusty Belt (279897, -0.59 DPS, sim-verified) [quest] |
| legs | Leggings of the Fang (10410) (or Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.9 ranged_attack_power points (1.94 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world]; Brawler's Leather Pants (252500, -0.41 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.8 ranged_attack_power points (1.73 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.43 DPS) [dungeon]; Agile Boots (4788, -0.65 DPS) [vendor] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Lavishly Jeweled Ring (1156, -0.86 DPS) [dungeon]; The 1 Ring (8350, -1.08 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.4 ranged_attack_power points (0.86 DPS) | yes | Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.8 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.63 DPS) [world]; Lupine Axe (1220, -0.85 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.48 DPS) | yes | Lil Timmy's Peashooter (13136, -1.35 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.74 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -3.01 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 221, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 140.5. Weights run: 3.0s. Verify run: 2.6s. 395 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.892 ± 0.013, crit=0.667 ± 0.017 per rating point (14 rating = 1%, 9.344 per %), hit=1.211 ± 0.055 per rating point (10 rating = 1%, 12.106 per %), melee_haste=9.499 ± 0.922

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Tribal Worg Helm (6204, -0.40 DPS) [world]; Brawler's Leather Hood (252504, -0.44 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Ghostshard Talisman (7731, -0.63 DPS) [dungeon]; Wolfpack Medallion (5754, -0.79 DPS) [world]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.18 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Cloak of Night (4447, -0.40 DPS) [world]; Fenrus' Hide (6340, -0.40 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Tunic of Westfall (2041, -0.67 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.19 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Serpent Gloves (5970, -0.40 DPS) [dungeon]; Gloves of the Fang (10413, -0.40 DPS) [dungeon]; Insignia Gloves (6408, -0.44 DPS, sim-verified) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.14 DPS) [rep]; Highlander's Leather Girdle (20117, -0.14 DPS) [rep] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.20 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Ring of Precision (1491, -0.59 DPS) [dungeon]; Protector's Band (19517, -0.59 DPS) [rep]; Signet of the Zhevra (285330, -0.59 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 20.2 ranged_attack_power points (1.39 DPS) | yes | Protector's Band (19517, -0.20 DPS) [rep]; Signet of the Zhevra (285330, -0.20 DPS) [world]; Ring of Precision (1491, -0.89 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 20.5 ranged_attack_power points (1.40 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 263.0 ranged_attack_power points (18.01 DPS) | yes | Silver Star (3463, -0.14 DPS) [quest]; Booty Bay Bruiser's Buckshot (274748, -0.57 DPS, sim-verified) [vendor]; Golemsight Long Gun (273029, -0.91 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Glass Shooter

No-known-source sample (15 of 395, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 183.0. Weights run: 3.2s. Verify run: 3.5s. 628 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.797 ± 0.013, crit=0.727 ± 0.020 per rating point (14 rating = 1%, 10.182 per %), hit=1.484 ± 0.093 per rating point (10 rating = 1%, 14.837 per %), melee_haste=9.080 ± 1.427

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world]; Nightscape Headband (8176, -1.25 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 30.8 ranged_attack_power points (2.03 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.71 DPS) [quest]; Ghostshard Talisman (7731, -1.11 DPS) [dungeon]; Wolfpack Medallion (5754, -1.29 DPS) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.8 ranged_attack_power points (2.82 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.83 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 28.0 ranged_attack_power points (1.85 DPS) | yes | Imperial Cloak (6432, -0.37 DPS) [dungeon]; Parachute Cloak (10518, -0.37 DPS) [crafted]; Tigerstrike Mantle (13108, -0.37 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (183.0 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.37 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.37 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Branded Leather Bracers (19508, -0.16 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.18 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.11 DPS) | yes | Gloves of Holy Might (867, -0.12 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.27 DPS) [crafted]; Scarlet Gauntlets (10331, -0.27 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.98 DPS) | yes | Scorpashi Sash (14652, -0.13 DPS) [world_drop]; Blackforge Girdle (6425, -0.32 DPS) [dungeon]; Ogron's Sash (13117, -0.32 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.7 ranged_attack_power points (3.88 DPS) | yes | Triprunner Dungarees (9624, -0.76 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted]; Worn Running Boots (9398, -0.37 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.0 ranged_attack_power points (1.85 DPS) | yes | Ironspine's Eye (7686, -0.18 DPS) [dungeon]; Protector's Band (19515, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 25.2 ranged_attack_power points (1.66 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Protector's Band (19515, -0.18 DPS) [rep]; Disengagement Ring (276202, -0.18 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -2.04 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Ginn-su Sword (9424), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Ginn-su Sword (9424, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (23.11 DPS) | yes | Shadowforge Bushmaster (9422, -2.07 DPS) [dungeon]; Swiftwind (13038, -2.17 DPS) [world_drop]; The Silencer (13138, -4.24 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 628, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 226.8. Weights run: 3.3s. Verify run: 4.5s. 787 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.771 ± 0.013, crit=0.900 ± 0.024 per rating point (14 rating = 1%, 12.597 per %), hit=1.628 ± 0.098 per rating point (10 rating = 1%, 16.277 per %), melee_haste=9.891 ± 1.505

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 56.9 ranged_attack_power points (3.90 DPS) | yes | Lordrec Helmet (10741, -0.86 DPS) [quest]; Sprightring Helm (17776, -1.05 DPS) [quest]; Helm of Fire (8348, -2.94 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.0 ranged_attack_power points (2.47 DPS) | yes | Sentinel's Medallion (19539, -0.19 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 48.6 ranged_attack_power points (3.33 DPS) | yes | Phytoskin Spaulders (17749, -0.29 DPS) [dungeon]; Sunburn Spaulders (274751, -0.42 DPS) [vendor]; Khan's Mantle (14787, -1.05 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 41.6 ranged_attack_power points (2.84 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.19 DPS) [dungeon]; Blackmetal Cape (9512, -0.57 DPS) [dungeon] |
| chest | Blazewind Breastplate (11193) | Tremors of the Earth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fungus Shroud Armor (17742, +0.00 DPS) [dungeon]; Knight's Chain Armor (220828, -0.66 DPS) [vendor]; Quillward Harness (10583, -0.76 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.6 ranged_attack_power points (2.84 DPS) | yes | Bloodlust Bracelets (14807, -0.76 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.76 DPS) [crafted]; Arena Bands (18711, -0.93 DPS) [world] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 61.0 ranged_attack_power points (4.17 DPS) | yes | Sergeant Major's Chain Gauntlets (220829, -1.34 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.52 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.52 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.6 ranged_attack_power points (3.73 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.08 DPS) [crafted]; Stalker's Mail Belt (252588, -1.08 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.2 ranged_attack_power points (3.98 DPS) | yes | Infernal Trickster Leggings (17754, -0.19 DPS) [dungeon]; Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.47 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.4 ranged_attack_power points (3.79 DPS) | yes | Fleetfoot Greaves (11627, -0.19 DPS) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.57 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.8 ranged_attack_power points (2.66 DPS) | yes | Falcon's Hook (7552, -0.95 DPS) [dungeon]; Ironspine's Eye (7686, -0.95 DPS) [dungeon]; Protector's Band (19516, -0.95 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.7 ranged_attack_power points (1.90 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Protector's Band (19516, -0.19 DPS) [rep] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Tooth (19992, -2.98 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (226.8 DPS) | yes | - |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -1.02 DPS) [world_drop]; Steel Spear (250605, -1.21 DPS) [crafted]; Hanzo Sword (8190, -5.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.54 DPS) [world_drop]; Hurricane (2824, -1.67 DPS) [world_drop]; Dark Iron Rifle (16004, -3.13 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Blazewind Breastplate; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 787, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 302.6. Weights run: 3.3s. Verify run: 13.5s. 1685 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.966 ± 0.021, crit=1.719 ± 0.043 per rating point (14 rating = 1%, 24.067 per %), hit=2.762 ± 0.158 per rating point (10 rating = 1%, 27.621 per %), melee_haste=18.958 ± 2.212

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -8.39 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Beads of Ogre Might (22150, -0.32 DPS) [quest]; Mark of Fordring (15411, -0.43 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 104.2 ranged_attack_power points (7.12 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 64.5 ranged_attack_power points (4.41 DPS) | yes | Howler's Furs (272414, -0.61 DPS) [vendor]; Shifting Cloak (18511, -0.96 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.96 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Breastplate (16466, -2.50 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.50 DPS) [vendor]; Tunic of Undead Slaying (23089, -12.65 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -7.91 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS, sim-verified) [quest]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -6.99 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 179.6 ranged_attack_power points (12.27 DPS) | yes | Sentinel's Leather Pants (237818, -3.51 DPS) [vendor]; Marshal's Chain Legguards (231577, -3.73 DPS) [pvp]; Plaguehound Leggings (18736, -4.30 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -7.71 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.82 DPS) [dungeon]; Don Julio's Band (19325, -2.13 DPS) [rep]; Naglering (11669, -7.76 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, +0.00 DPS) [dungeon]; Don Julio's Band (19325, -0.10 DPS) [rep]; Innervating Band (18701, -1.68 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Second Wind (11819, -5.28 DPS, sim-verified) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (302.6 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -0.20 DPS) [dungeon]; Hand of Justice (11815, -0.33 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -7.98 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -9.76 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 762.8. Weights run: 3.4s. Verify run: 13.5s. 1685 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.971 ± 0.019, crit=1.772 ± 0.040 per rating point (14 rating = 1%, 24.802 per %), hit=3.898 ± 0.214 per rating point (10 rating = 1%, 38.981 per %), melee_haste=20.965 ± 2.769

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -10.91 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.93 DPS) [quest]; Mark of Fordring (15411, -1.73 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 105.0 ranged_attack_power points (14.92 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 67.0 ranged_attack_power points (9.51 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Shifting Cloak (18511, -2.34 DPS) [crafted]; Shadow Prowler's Cloak (22269, -2.34 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Breastplate (16466, -6.80 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -6.80 DPS) [vendor]; Tunic of Undead Slaying (23089, -22.51 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -13.50 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Gauntlets of Accuracy (18349, +0.00 DPS) [dungeon]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Marksman's Girdle (22232, -10.77 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 192.6 ranged_attack_power points (27.35 DPS) | yes | Sentinel's Leather Pants (237818, -8.91 DPS) [vendor]; Plaguehound Leggings (18736, -9.16 DPS) [dungeon]; Marshal's Chain Legguards (231577, -9.48 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -14.05 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.80 DPS) [dungeon]; Don Julio's Band (19325, -4.33 DPS) [rep]; Naglering (11669, -12.24 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, +0.00 DPS) [dungeon]; Don Julio's Band (19325, -0.11 DPS) [rep]; Innervating Band (18701, -5.11 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -1.86 DPS) [dungeon]; Hand of Justice (11815, -2.14 DPS) [dungeon] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-verified (762.8 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -12.62 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Bow of Searing Arrows (2825, -32.51 DPS, sim-verified) [world_drop] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Darkmoon Card: Heroism; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 111.4. Weights run: 2.8s. Verify run: 2.3s. 210 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.101 ± 0.014, crit=0.593 ± 0.016 per rating point (14 rating = 1%, 8.306 per %), hit=1.060 ± 0.050 per rating point (10 rating = 1%, 10.605 per %), melee_haste=8.276 ± 0.789

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.8 ranged_attack_power points (1.73 DPS) | yes | Red Winter Hat (21524, -1.84 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.92 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Cape of the Brotherhood (5193, -0.22 DPS) [dungeon]; Hide of Lupos (3018, -0.43 DPS) [world]; Bristlebark Cape (14571, -0.43 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 21.7 ranged_attack_power points (1.51 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.22 DPS) [crafted]; Prospector's Chestpiece (14562, -0.22 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bristlebark Bindings (14569, -0.43 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.43 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -1.57 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Guardsman Belt (3429, -0.39 DPS) [world]; Dark Leather Belt (4249, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.9 ranged_attack_power points (1.94 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Agile Boots (4788, -0.22 DPS) [vendor]; Feet of the Lynx (1121, -1.22 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Bounty Hunter's Ring (5351, -0.65 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.86 DPS) [dungeon]; The 1 Ring (8350, -1.08 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.4 ranged_attack_power points (0.86 DPS) | yes | Bounty Hunter's Ring (5351, -0.22 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.8 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.63 DPS) [world]; Crescent Staff (6505, -0.63 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.48 DPS) | yes | Outrider's Bow (20437, -0.68 DPS) [pvp]; Lil Timmy's Peashooter (13136, -1.74 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.74 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 210, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 142.5. Weights run: 3.0s. Verify run: 2.8s. 380 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.892 ± 0.013, crit=0.667 ± 0.017 per rating point (14 rating = 1%, 9.344 per %), hit=1.211 ± 0.055 per rating point (10 rating = 1%, 12.106 per %), melee_haste=9.499 ± 0.922

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Tribal Worg Helm (6204, -0.40 DPS) [world]; Brawler's Leather Hood (252504, -0.47 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Ghostshard Talisman (7731, -0.63 DPS) [dungeon]; Wolfpack Medallion (5754, -0.79 DPS) [world]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.18 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Cloak of Night (4447, -0.40 DPS) [world]; Swiftrunner Cape (6745, -0.40 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Panther Armor (6670, -1.13 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.19 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Loamflake Bracers (15462, -0.20 DPS) [quest] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Braced Handguards (6784, -0.20 DPS) [quest]; Serpent Gloves (5970, -0.40 DPS) [dungeon]; Insignia Gloves (6408, -0.40 DPS) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.14 DPS) [rep]; Defiler's Leather Girdle (20191, -0.14 DPS) [rep] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Vorrel's Boots (7751) | Vorrel's Revenge [quest] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Warsong Boots (16977, -0.40 DPS) [quest]; Highlander's Mail Greaves (20123, -0.40 DPS) [vendor]; Insignia Boots (4055, -1.13 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Ring of Precision (1491, -0.59 DPS) [dungeon]; Legionnaire's Band (19513, -0.59 DPS) [rep]; Signet of the Zhevra (285330, -0.59 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 20.2 ranged_attack_power points (1.39 DPS) | yes | Ring of Precision (1491, -0.20 DPS) [dungeon]; Legionnaire's Band (19513, -0.20 DPS) [rep]; Signet of the Zhevra (285330, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | 24.5 ranged_attack_power points (1.67 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 20.5 ranged_attack_power points (1.40 DPS) | yes | Vendetta (776, +0.00 DPS) [dungeon]; Satyr's Rod (15962, -1.20 DPS) [world_drop] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 263.0 ranged_attack_power points (18.01 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, +0.00 DPS, sim-verified) [vendor]; Silver Star (3463, -0.14 DPS) [quest]; Golemsight Long Gun (273029, -0.91 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Vorrel's Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Serrated Raptor Claw; off_hand: Alliance Outrunner's Sword; ranged: Glass Shooter

No-known-source sample (15 of 380, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 185.0. Weights run: 3.2s. Verify run: 3.4s. 593 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.797 ± 0.013, crit=0.727 ± 0.020 per rating point (14 rating = 1%, 10.182 per %), hit=1.484 ± 0.093 per rating point (10 rating = 1%, 14.837 per %), melee_haste=9.080 ± 1.427

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Nightscape Headband (8176, -0.18 DPS) [crafted]; Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 30.8 ranged_attack_power points (2.03 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.71 DPS) [quest]; Ghostshard Talisman (7731, -1.11 DPS) [dungeon]; Wolfpack Medallion (5754, -1.29 DPS) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.8 ranged_attack_power points (2.82 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.84 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 28.0 ranged_attack_power points (1.85 DPS) | yes | Imperial Cloak (6432, -0.37 DPS) [dungeon]; Parachute Cloak (10518, -0.37 DPS) [crafted]; Tigerstrike Mantle (13108, -0.37 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (185.0 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.37 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.37 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.11 DPS) | yes | Gloves of Holy Might (867, -0.12 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.27 DPS) [crafted]; Scarlet Gauntlets (10331, -0.27 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.98 DPS) | yes | Scorpashi Sash (14652, -0.13 DPS) [world_drop]; Blackforge Girdle (6425, -0.32 DPS) [dungeon]; Ogron's Sash (13117, -0.32 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.7 ranged_attack_power points (3.88 DPS) | yes | Triprunner Dungarees (9624, -0.69 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted]; Worn Running Boots (9398, -0.37 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.0 ranged_attack_power points (1.85 DPS) | yes | Ironspine's Eye (7686, -0.18 DPS) [dungeon]; Legionnaire's Band (19512, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 25.2 ranged_attack_power points (1.66 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Legionnaire's Band (19512, -0.18 DPS) [rep]; Disengagement Ring (276202, -0.18 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.07 DPS, sim-verified) [world_drop] |
| off_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | 24.0 ranged_attack_power points (1.58 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Satyr's Rod (15962, -1.40 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (23.11 DPS) | yes | Outrider's Bow (19560, -1.56 DPS) [pvp]; Shadowforge Bushmaster (9422, -2.07 DPS) [dungeon]; The Silencer (13138, -3.79 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Serrated Raptor Claw; ranged: Bow of Searing Arrows

No-known-source sample (15 of 593, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 231.9. Weights run: 3.3s. Verify run: 4.2s. 745 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.771 ± 0.013, crit=0.900 ± 0.024 per rating point (14 rating = 1%, 12.597 per %), hit=1.628 ± 0.098 per rating point (10 rating = 1%, 16.277 per %), melee_haste=9.891 ± 1.505

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 47.1 ranged_attack_power points (3.22 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.38 DPS) [quest]; Tough Scorpid Helm (8208, -0.57 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.0 ranged_attack_power points (2.47 DPS) | yes | Scout's Medallion (19535, -0.19 DPS) [rep]; Woven Ivy Necklace (19159, -0.76 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phytoskin Spaulders (17749, +0.00 DPS) [dungeon]; Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Khan's Mantle (14787, -0.63 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 41.6 ranged_attack_power points (2.84 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.19 DPS) [dungeon]; Blackmetal Cape (9512, -0.57 DPS) [dungeon] |
| chest | Blazewind Breastplate (11193) | Broken Alliances [quest] | sim-verified (231.9 DPS) | yes | Fungus Shroud Armor (17742, +0.00 DPS) [dungeon]; Stone Guard's Chain Armor (220827, -0.66 DPS) [vendor]; Quillward Harness (10583, -0.76 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.6 ranged_attack_power points (2.84 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Bloodlust Bracelets (14807, -1.04 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 61.0 ranged_attack_power points (4.17 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -1.31 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.52 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.52 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.6 ranged_attack_power points (3.73 DPS) | yes | Sagebrush Girdle (17778, -0.89 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.08 DPS) [crafted]; Stalker's Mail Belt (252588, -1.08 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.2 ranged_attack_power points (3.98 DPS) | yes | Infernal Trickster Leggings (17754, -0.19 DPS) [dungeon]; Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Stone Guard's Chain Legplates (220833, -0.47 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.4 ranged_attack_power points (3.79 DPS) | yes | Fleetfoot Greaves (11627, -0.19 DPS) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.57 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.8 ranged_attack_power points (2.66 DPS) | yes | Falcon's Hook (7552, -0.95 DPS) [dungeon]; Ironspine's Eye (7686, -0.95 DPS) [dungeon]; Legionnaire's Band (19511, -0.95 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.7 ranged_attack_power points (1.90 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Legionnaire's Band (19511, -0.19 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -5.53 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.50 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -1.02 DPS) [world_drop]; Steel Spear (250605, -1.21 DPS) [crafted]; Hanzo Sword (8190, -4.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.54 DPS) [world_drop]; Hurricane (2824, -1.67 DPS) [world_drop]; Dark Iron Rifle (16004, -2.07 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Blazewind Breastplate; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 745, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 311.0. Weights run: 3.3s. Verify run: 12.8s. 1664 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.966 ± 0.021, crit=1.719 ± 0.043 per rating point (14 rating = 1%, 24.067 per %), hit=2.762 ± 0.158 per rating point (10 rating = 1%, 27.621 per %), melee_haste=18.958 ± 2.212

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -9.77 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Beads of Ogre Might (22150, -0.32 DPS) [quest]; Mark of Fordring (15411, -0.43 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 104.2 ranged_attack_power points (7.12 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -0.68 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 64.5 ranged_attack_power points (4.41 DPS) | yes | Howler's Furs (272414, -0.61 DPS) [vendor]; Shifting Cloak (18511, -0.96 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.96 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Chestpiece (16565, -2.50 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.50 DPS) [vendor]; Tunic of Undead Slaying (23089, -13.83 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -8.38 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -7.43 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 179.6 ranged_attack_power points (12.27 DPS) | yes | Outrider's Chain Leggings (22673, -1.81 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -3.51 DPS) [vendor]; General's Chain Legguards (231574, -3.73 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -8.01 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.82 DPS) [dungeon]; Don Julio's Band (19325, -2.13 DPS) [rep]; Naglering (11669, -8.81 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-verified (311.0 DPS) | yes | Tarnished Elven Ring (18500, +0.00 DPS) [dungeon]; Don Julio's Band (19325, -0.10 DPS) [rep]; Innervating Band (18701, -1.68 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+9.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, -3.77 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.36 DPS) [crafted]; Counterattack Lodestone (18537, -5.56 DPS) [dungeon] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -2.63 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -8.98 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.20 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket2: Burst of Knowledge; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1664, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 776.3. Weights run: 3.4s. Verify run: 13.1s. 1664 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.971 ± 0.019, crit=1.772 ± 0.040 per rating point (14 rating = 1%, 24.802 per %), hit=3.898 ± 0.214 per rating point (10 rating = 1%, 38.981 per %), melee_haste=20.965 ± 2.769

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -16.37 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.93 DPS) [quest]; Mark of Fordring (15411, -1.73 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 105.0 ranged_attack_power points (14.92 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 67.0 ranged_attack_power points (9.51 DPS) | yes | Cape of the Black Baron (13340, -0.34 DPS) [dungeon]; Shifting Cloak (18511, -2.34 DPS) [crafted]; Shadow Prowler's Cloak (22269, -2.34 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Chestpiece (16565, -6.80 DPS) [vendor]; Warlord's Chain Hauberk (231573, -6.80 DPS) [vendor]; Tunic of Undead Slaying (23089, -26.96 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -12.98 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -6.73 DPS, sim-verified) [quest] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -12.32 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 192.6 ranged_attack_power points (27.35 DPS) | yes | Outrider's Chain Leggings (22673, -3.53 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -8.91 DPS) [vendor]; Plaguehound Leggings (18736, -9.16 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -11.89 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.80 DPS) [dungeon]; Don Julio's Band (19325, -4.33 DPS) [rep]; Naglering (11669, -16.74 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, +0.00 DPS) [dungeon]; Don Julio's Band (19325, -0.11 DPS) [rep]; Innervating Band (18701, -5.11 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Counterattack Lodestone (18537, -12.68 DPS) [dungeon]; Hand of Justice (11815, -12.97 DPS) [dungeon]; Blackhand's Breadth (13965, -19.12 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (776.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -1.86 DPS) [dungeon]; Hand of Justice (11815, -2.14 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -17.08 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -24.09 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1664, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

