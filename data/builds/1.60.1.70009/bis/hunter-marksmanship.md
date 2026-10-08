# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 98.1. Weights run: 2.2s. Verify run: 2.0s. 220 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.620 ± 0.008, crit=0.619 ± 0.016 per rating point (14 rating = 1%, 8.660 per %), hit=0.998 ± 0.037 per rating point (10 rating = 1%, 9.979 per %), melee_haste=not significant (0.108 ± 0.741)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 21.0 ranged_attack_power points (1.36 DPS) | yes | Red Winter Hat (21524, -1.38 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 15.7 ranged_attack_power points (1.02 DPS) | yes | Erudite's Amulet (277204, -0.34 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 13.1 ranged_attack_power points (0.85 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 15.7 ranged_attack_power points (1.02 DPS) | yes | Hide of Lupos (3018, -0.34 DPS) [world]; Bristlebark Cape (14571, -0.34 DPS) [world_drop]; Cape of the Brotherhood (5193, -1.15 DPS, sim-verified) [dungeon] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 28.8 ranged_attack_power points (1.87 DPS) | yes | Trapper's Leather Armor (252491, -0.68 DPS) [crafted]; Brawler's Leather Armor (252490, -0.80 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.85 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 13.1 ranged_attack_power points (0.85 DPS) | yes | Wolf Bracers (4794, -0.17 DPS) [vendor]; Bravo's Armbands (270015, -0.17 DPS) [quest]; Ratchet Wristwraps (274742, -0.34 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (98.1 DPS) | yes | Forest Leather Gloves (3058, -0.34 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.34 DPS) [crafted]; Serpent Gloves (5970, -1.12 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.17 DPS) | yes | Deviate Scale Belt (6468, -0.32 DPS) [crafted]; Dark Leather Belt (4249, -0.49 DPS) [crafted]; Dusty Belt (279897, -0.60 DPS, sim-verified) [quest] |
| legs | Leggings of the Fang (10410) (or Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 23.6 ranged_attack_power points (1.53 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.17 DPS) [world]; Brawler's Leather Pants (252500, -0.30 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.0 ranged_attack_power points (1.36 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.34 DPS) [dungeon]; Agile Boots (4788, -0.51 DPS) [vendor] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 15.7 ranged_attack_power points (1.02 DPS) | yes | Lavishly Jeweled Ring (1156, -0.68 DPS) [dungeon]; The 1 Ring (8350, -0.85 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.5 ranged_attack_power points (0.68 DPS) | yes | Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; The 1 Ring (8350, -0.51 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 27.0 ranged_attack_power points (1.75 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.56 DPS) [world]; Lupine Axe (1220, -0.73 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.9 ranged_attack_power points (11.64 DPS) | yes | Lil Timmy's Peashooter (13136, -1.83 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.53 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.78 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 123.7. Weights run: 2.4s. Verify run: 2.2s. 366 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.715 ± 0.012, crit=0.888 ± 0.023 per rating point (14 rating = 1%, 12.437 per %), hit=1.205 ± 0.052 per rating point (10 rating = 1%, 12.054 per %), melee_haste=10.286 ± 0.925

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 27.1 ranged_attack_power points (1.78 DPS) | yes | Tribal Worg Helm (6204, -0.36 DPS) [world]; Brawler's Leather Hood (252504, -0.36 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.53 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 21.7 ranged_attack_power points (1.42 DPS) | yes | Ghostshard Talisman (7731, -0.51 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.71 DPS) [world_drop]; Erudite's Amulet (277204, -0.71 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.9 ranged_attack_power points (1.96 DPS) | yes | Mantle of Thieves (2264, -0.18 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.71 DPS) [crafted]; Insignia Mantle (4721, -0.71 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 21.7 ranged_attack_power points (1.42 DPS) | yes | Hawkeye's Cloak (14593, -0.18 DPS) [world_drop]; Cloak of Night (4447, -0.36 DPS) [world]; Fenrus' Hide (6340, -0.36 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 38.0 ranged_attack_power points (2.49 DPS) | yes | Tunic of Westfall (2041, -0.54 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.07 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.07 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 16.3 ranged_attack_power points (1.07 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.18 DPS) [world_drop] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 21.7 ranged_attack_power points (1.42 DPS) | yes | Serpent Gloves (5970, -0.36 DPS) [dungeon]; Gloves of the Fang (10413, -0.36 DPS) [dungeon]; Insignia Gloves (6408, -0.36 DPS, sim-verified) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 24.4 ranged_attack_power points (1.60 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.03 DPS) [rep]; Highlander's Leather Girdle (20117, -0.03 DPS) [rep] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 38.0 ranged_attack_power points (2.49 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.18 DPS) [crafted]; Ferine Leggings (6690, -0.79 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 21.7 ranged_attack_power points (1.42 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.18 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 24.4 ranged_attack_power points (1.60 DPS) | yes | Ring of Precision (1491, -0.53 DPS) [dungeon]; Protector's Band (19517, -0.53 DPS) [rep]; Signet of the Zhevra (285330, -0.53 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 19.0 ranged_attack_power points (1.25 DPS) | yes | Protector's Band (19517, -0.18 DPS) [rep]; Signet of the Zhevra (285330, -0.18 DPS) [world]; Ring of Precision (1491, -0.69 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 19.6 ranged_attack_power points (1.28 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.3 ranged_attack_power points (1.07 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.14 DPS) | yes | Glass Shooter (9456, -0.69 DPS) [dungeon]; Ironweaver (13137, -1.26 DPS) [world_drop]; Silver Star (3463, -2.13 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 161.0. Weights run: 2.5s. Verify run: 2.6s. 597 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.753 ± 0.014, crit=1.080 ± 0.029 per rating point (14 rating = 1%, 15.122 per %), hit=1.304 ± 0.063 per rating point (10 rating = 1%, 13.042 per %), melee_haste=6.902 ± 0.972

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 35.8 ranged_attack_power points (2.51 DPS) | yes | Nightscape Headband (8176, -0.19 DPS) [crafted]; Guard's Chain Helm (250499, -0.19 DPS) [crafted]; Skullsplitter Helm (1624, -0.39 DPS) [world] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 30.3 ranged_attack_power points (2.13 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.72 DPS) [quest]; Ghostshard Talisman (7731, -1.14 DPS) [dungeon]; Erudite's Amulet (277204, -1.35 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.3 ranged_attack_power points (2.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.84 DPS) [world_drop]; Nightscape Shoulders (8192, -0.88 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -1.04 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 27.5 ranged_attack_power points (1.93 DPS) | yes | Imperial Cloak (6432, -0.39 DPS) [dungeon]; Parachute Cloak (10518, -0.39 DPS) [crafted]; Tigerstrike Mantle (13108, -0.39 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 52.3 ranged_attack_power points (3.67 DPS) | yes | Wolffear Harness (13110, -0.39 DPS) [world_drop]; Nightscape Tunic (8175, -0.77 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.77 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.0 ranged_attack_power points (1.55 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Branded Leather Bracers (19508, -0.14 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.19 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 35.1 ranged_attack_power points (2.46 DPS) | yes | Gauntlets of Divinity (7724, -0.22 DPS) [dungeon]; Dragonscale Gauntlets (8347, -0.24 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.53 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (2.11 DPS) | yes | Scorpashi Sash (14652, -0.17 DPS) [world_drop]; Blackforge Girdle (6425, -0.37 DPS) [dungeon]; Ogron's Sash (13117, -0.37 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 57.8 ranged_attack_power points (4.06 DPS) | yes | Triprunner Dungarees (9624, -0.85 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.35 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.35 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 35.8 ranged_attack_power points (2.51 DPS) | yes | Imperial Leather Boots (6431, -0.39 DPS) [dungeon]; Dusky Boots (7390, -0.39 DPS) [crafted]; Worn Running Boots (9398, -0.39 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 27.5 ranged_attack_power points (1.93 DPS) | yes | Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Protector's Band (19515, -0.39 DPS) [rep]; Disengagement Ring (276202, -0.39 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 24.8 ranged_attack_power points (1.74 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Protector's Band (19515, -0.19 DPS) [rep]; Disengagement Ring (276202, -0.19 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (161.0 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -2.22 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Ginn-su Sword (9424), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.0 ranged_attack_power points (1.55 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Ginn-su Sword (9424, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (24.56 DPS) | yes | Shadowforge Bushmaster (9422, -2.20 DPS) [dungeon]; Swiftwind (13038, -2.33 DPS) [world_drop]; The Silencer (13138, -2.60 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 222.7. Weights run: 2.6s. Verify run: 3.3s. 754 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.809 ± 0.015, crit=1.241 ± 0.031 per rating point (14 rating = 1%, 17.373 per %), hit=1.586 ± 0.082 per rating point (10 rating = 1%, 15.855 per %), melee_haste=10.926 ± 1.165

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 62.3 ranged_attack_power points (4.76 DPS) | yes | Lordrec Helmet (10741, -1.33 DPS) [quest]; Sprightring Helm (17776, -1.54 DPS) [quest]; Helm of Fire (8348, -4.02 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.5 ranged_attack_power points (2.79 DPS) | yes | Sentinel's Medallion (19539, -0.21 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 53.9 ranged_attack_power points (4.12 DPS) | yes | Sunburn Spaulders (274751, -0.84 DPS) [vendor]; Phytoskin Spaulders (17749, -1.09 DPS, sim-verified) [dungeon]; Khan's Mantle (14787, -1.54 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 42.1 ranged_attack_power points (3.22 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.21 DPS) [dungeon]; Blackmetal Cape (9512, -0.64 DPS) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 70.2 ranged_attack_power points (5.36 DPS) | yes | Blazewind Breastplate (11193, -0.43 DPS) [quest]; Knight's Chain Armor (220828, -0.82 DPS) [vendor]; Quillward Harness (10583, -1.29 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 42.1 ranged_attack_power points (3.22 DPS) | yes | Bloodlust Bracelets (14807, -0.86 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.86 DPS) [crafted]; Pridelord Bands (14672, -1.07 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 61.8 ranged_attack_power points (4.72 DPS) | yes | Beastmaster's Gauntlets (226883, -1.72 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.72 DPS) [crafted]; Sergeant Major's Chain Gauntlets (220829, -1.72 DPS, sim-verified) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 55.1 ranged_attack_power points (4.21 DPS) | yes | Sagebrush Girdle (17778, -0.99 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.21 DPS) [crafted]; Stalker's Mail Belt (252588, -1.21 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 59.0 ranged_attack_power points (4.50 DPS) | yes | Knight's Chain Legplates (220832, +0.00 DPS, sim-verified) [vendor]; Infernal Trickster Leggings (17754, -0.21 DPS) [dungeon]; Keeper's Woolies (14668, -0.43 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 56.2 ranged_attack_power points (4.29 DPS) | yes | Fleetfoot Greaves (11627, -0.21 DPS) [dungeon]; Elven Chain Boots (13125, -0.43 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.64 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.3 ranged_attack_power points (3.00 DPS) | yes | Ring of the Underwood (2951, -0.86 DPS) [world_drop]; Falcon's Hook (7552, -1.07 DPS) [dungeon]; Ironspine's Eye (7686, -1.07 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.9 ranged_attack_power points (2.74 DPS) | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.81 DPS) [dungeon]; Ironspine's Eye (7686, -0.81 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (222.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Devilsaur Tooth (19992) | The Green Drake [quest] | sim-verified (222.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (222.7 DPS) | yes | Warmonger (13052, -1.30 DPS) [world_drop]; Steel Spear (250605, -1.38 DPS) [crafted]; Hanzo Sword (8190, -5.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (222.7 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.68 DPS) [world_drop]; Hurricane (2824, -1.86 DPS) [world_drop]; Dark Iron Rifle (16004, -5.91 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Devilsaur Tooth; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 306.8. Weights run: 2.5s. Verify run: 10.5s. 1670 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.110 ± 0.026, crit=2.158 ± 0.054 per rating point (14 rating = 1%, 30.209 per %), hit=3.513 ± 0.145 per rating point (10 rating = 1%, 35.129 per %), melee_haste=19.852 ± 2.143

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (306.8 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -11.28 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 81.8 ranged_attack_power points (6.20 DPS) | yes | Amulet of the Darkmoon (19491, -1.72 DPS) [quest]; Mark of Fordring (15411, -1.94 DPS) [quest]; Beads of Ogre Might (22150, -2.05 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 114.2 ranged_attack_power points (8.65 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.7 ranged_attack_power points (5.05 DPS) | yes | Shifting Cloak (18511, -1.04 DPS) [crafted]; Shadow Prowler's Cloak (22269, -1.04 DPS) [dungeon]; Howler's Furs (272414, -1.82 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (306.8 DPS) | yes | Field Marshal's Chain Breastplate (16466, -3.37 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -3.37 DPS) [vendor]; Tunic of Undead Slaying (23089, -16.05 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (306.8 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -10.28 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (306.8 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (306.8 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -10.45 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 204.4 ranged_attack_power points (15.49 DPS) | yes | Sentinel's Leather Pants (237818, -4.55 DPS) [vendor]; Marshal's Chain Legguards (231577, -5.19 DPS) [pvp]; Plaguehound Leggings (18736, -5.76 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (306.8 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -10.73 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (306.8 DPS) | yes | Don Julio's Band (19325, -0.74 DPS) [rep]; Cutthroat's Signet (272408, -0.94 DPS) [vendor]; Naglering (11669, -7.86 DPS, sim-verified) [dungeon] |
| finger2 | Tarnished Elven Ring (18500) | Dire Maul: Tribute [dungeon] | sim-verified (306.8 DPS) | yes | Don Julio's Band (19325, -0.03 DPS) [rep]; Cutthroat's Signet (272408, -0.24 DPS) [vendor]; Naglering (11669, -6.97 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (306.8 DPS) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (306.8 DPS) | yes | Devilsaur Eye (19991, -1.43 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -2.18 DPS) [crafted]; Counterattack Lodestone (18537, -2.91 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (306.8 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -9.61 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (306.8 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -12.78 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Tarnished Elven Ring; trinket1: Burst of Knowledge; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 762.2. Weights run: 2.6s. Verify run: 11.2s. 1670 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.129 ± 0.025, crit=2.277 ± 0.051 per rating point (14 rating = 1%, 31.876 per %), hit=3.998 ± 0.199 per rating point (10 rating = 1%, 39.984 per %), melee_haste=21.245 ± 2.668

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (762.2 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -2.80 DPS, sim-verified) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 86.9 ranged_attack_power points (14.36 DPS) | yes | Amulet of the Darkmoon (19491, -4.54 DPS) [quest]; Beads of Ogre Might (22150, -4.56 DPS, sim-verified) [quest]; Mark of Fordring (15411, -4.80 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 116.4 ranged_attack_power points (19.23 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (762.2 DPS) | yes | Shifting Cloak (18511, -2.27 DPS) [crafted]; Shadow Prowler's Cloak (22269, -2.27 DPS) [dungeon]; Howler's Furs (272414, -7.89 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (762.2 DPS) | yes | Field Marshal's Chain Breastplate (16466, -8.16 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -8.16 DPS) [vendor]; Tunic of Undead Slaying (23089, -28.46 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (762.2 DPS) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -7.29 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 108.1 ranged_attack_power points (17.87 DPS) | yes | Gauntlets of Accuracy (18349, +0.00 DPS, sim-verified) [dungeon]; Marshal's Chain Vices (231578, -1.74 DPS) [vendor]; Marshal's Chain Grips (231560, -3.79 DPS) [pvp] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 105.7 ranged_attack_power points (17.47 DPS) | yes | Belt of Preserved Heads (20216, -3.77 DPS, sim-verified) [quest]; Warpwood Binding (18393, -4.96 DPS) [dungeon]; Ranger's Belt (272397, -5.06 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 213.3 ranged_attack_power points (35.24 DPS) | yes | Sentinel's Leather Pants (237818, -10.74 DPS) [vendor]; Marshal's Chain Legguards (231577, -12.39 DPS) [pvp]; Plaguehound Leggings (18736, -13.12 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (762.2 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -8.26 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (762.2 DPS) | yes | Tarnished Elven Ring (18500, -1.55 DPS) [dungeon]; Cutthroat's Signet (272408, -2.07 DPS) [vendor]; Naglering (11669, -11.41 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (762.2 DPS) | yes | Tarnished Elven Ring (18500, -0.15 DPS) [dungeon]; Cutthroat's Signet (272408, -0.67 DPS) [vendor]; Naglering (11669, -8.99 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (762.2 DPS) | yes | Frozen Heart of the Mountain (249469, -4.59 DPS) [crafted]; Counterattack Lodestone (18537, -6.90 DPS) [dungeon]; Hand of Justice (11815, -7.23 DPS) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (762.2 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Heart of Wyrmthalak (22321, -7.54 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (762.2 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -15.03 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (762.2 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -30.76 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Marksman's Girdle; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Blackhand's Breadth; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 99.0. Weights run: 2.2s. Verify run: 1.9s. 209 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.620 ± 0.008, crit=0.619 ± 0.016 per rating point (14 rating = 1%, 8.660 per %), hit=0.998 ± 0.037 per rating point (10 rating = 1%, 9.979 per %), melee_haste=not significant (0.108 ± 0.741)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 21.0 ranged_attack_power points (1.36 DPS) | yes | Red Winter Hat (21524, -1.40 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 15.7 ranged_attack_power points (1.02 DPS) | yes | Erudite's Amulet (277204, -0.35 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 13.1 ranged_attack_power points (0.85 DPS) | yes | Slime-encrusted Pads (6461, -0.75 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 15.7 ranged_attack_power points (1.02 DPS) | yes | Cape of the Brotherhood (5193, -0.31 DPS, sim-verified) [dungeon]; Hide of Lupos (3018, -0.34 DPS) [world]; Bristlebark Cape (14571, -0.34 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 18.3 ranged_attack_power points (1.19 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Dark Leather Tunic (2317, -0.17 DPS) [crafted]; Prospector's Chestpiece (14562, -0.17 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 13.1 ranged_attack_power points (0.85 DPS) | yes | Wolf Bracers (4794, -0.17 DPS) [vendor]; Bristlebark Bindings (14569, -0.34 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.34 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.34 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.34 DPS) [crafted]; Serpent Gloves (5970, -1.27 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.17 DPS) | yes | Deviate Scale Belt (6468, -0.31 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.49 DPS) [world]; Dark Leather Belt (4249, -0.49 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 23.6 ranged_attack_power points (1.53 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.17 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Agile Boots (4788, -0.17 DPS) [vendor]; Feet of the Lynx (1121, -0.98 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 15.7 ranged_attack_power points (1.02 DPS) | yes | Bounty Hunter's Ring (5351, -0.51 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.68 DPS) [dungeon]; The 1 Ring (8350, -0.85 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.5 ranged_attack_power points (0.68 DPS) | yes | Bounty Hunter's Ring (5351, -0.17 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; The 1 Ring (8350, -0.51 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 27.0 ranged_attack_power points (1.75 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.56 DPS) [world]; Crescent Staff (6505, -0.56 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.9 ranged_attack_power points (11.64 DPS) | yes | Outrider's Bow (20437, -0.67 DPS) [pvp]; Lil Timmy's Peashooter (13136, -0.97 DPS) [world_drop]; Cracked Blacksmith Hammer (285279, -2.53 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 126.0. Weights run: 2.4s. Verify run: 2.2s. 352 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.715 ± 0.012, crit=0.888 ± 0.023 per rating point (14 rating = 1%, 12.437 per %), hit=1.205 ± 0.052 per rating point (10 rating = 1%, 12.054 per %), melee_haste=10.286 ± 0.925

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 27.1 ranged_attack_power points (1.78 DPS) | yes | Tribal Worg Helm (6204, -0.36 DPS) [world]; Brawler's Leather Hood (252504, -0.37 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.53 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 21.7 ranged_attack_power points (1.42 DPS) | yes | Ghostshard Talisman (7731, -0.51 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.71 DPS) [world_drop]; Erudite's Amulet (277204, -0.71 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.9 ranged_attack_power points (1.96 DPS) | yes | Mantle of Thieves (2264, +0.00 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.71 DPS) [crafted]; Insignia Mantle (4721, -0.71 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 21.7 ranged_attack_power points (1.42 DPS) | yes | Hawkeye's Cloak (14593, -0.18 DPS) [world_drop]; Cloak of Night (4447, -0.36 DPS) [world]; Swiftrunner Cape (6745, -0.36 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 38.0 ranged_attack_power points (2.49 DPS) | yes | Panther Armor (6670, -0.93 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.07 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.07 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 16.3 ranged_attack_power points (1.07 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Loamflake Bracers (15462, -0.18 DPS) [quest] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 21.7 ranged_attack_power points (1.42 DPS) | yes | Braced Handguards (6784, -0.18 DPS) [quest]; Serpent Gloves (5970, -0.36 DPS) [dungeon]; Insignia Gloves (6408, -0.36 DPS) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 24.4 ranged_attack_power points (1.60 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.03 DPS) [rep]; Defiler's Leather Girdle (20191, -0.03 DPS) [rep] |
| legs | Troll's Bane Leggings (13114) | World drop [world_drop] | sim-verified (126.0 DPS) | yes | Dusky Leather Leggings (7373, -0.18 DPS) [crafted]; Ferine Leggings (6690, -0.79 DPS) [dungeon]; Petrolspill Leggings (9509, -1.39 DPS, sim-verified) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 21.7 ranged_attack_power points (1.42 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 24.4 ranged_attack_power points (1.60 DPS) | yes | Ring of Precision (1491, -0.53 DPS) [dungeon]; Legionnaire's Band (19513, -0.53 DPS) [rep]; Signet of the Zhevra (285330, -0.53 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 19.0 ranged_attack_power points (1.25 DPS) | yes | Ring of Precision (1491, -0.18 DPS) [dungeon]; Legionnaire's Band (19513, -0.18 DPS) [rep]; Signet of the Zhevra (285330, -0.18 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 19.6 ranged_attack_power points (1.28 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.3 ranged_attack_power points (1.07 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.14 DPS) | yes | Glass Shooter (9456, -0.69 DPS) [dungeon]; Ironweaver (13137, -1.26 DPS) [world_drop]; Silver Star (3463, -1.49 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Skulker's Leather Belt; legs: Troll's Bane Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 163.8. Weights run: 2.5s. Verify run: 2.5s. 563 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.753 ± 0.014, crit=1.080 ± 0.029 per rating point (14 rating = 1%, 15.122 per %), hit=1.304 ± 0.063 per rating point (10 rating = 1%, 13.042 per %), melee_haste=6.902 ± 0.972

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 35.8 ranged_attack_power points (2.51 DPS) | yes | Nightscape Headband (8176, -0.19 DPS) [crafted]; Guard's Chain Helm (250499, -0.19 DPS) [crafted]; Skullsplitter Helm (1624, -0.39 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 30.3 ranged_attack_power points (2.13 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.72 DPS) [quest]; Ghostshard Talisman (7731, -1.14 DPS) [dungeon]; Erudite's Amulet (277204, -1.35 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.3 ranged_attack_power points (2.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.84 DPS) [world_drop]; Nightscape Shoulders (8192, -0.89 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -1.04 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 27.5 ranged_attack_power points (1.93 DPS) | yes | Imperial Cloak (6432, -0.39 DPS) [dungeon]; Parachute Cloak (10518, -0.39 DPS) [crafted]; Tigerstrike Mantle (13108, -0.39 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 52.3 ranged_attack_power points (3.67 DPS) | yes | Wolffear Harness (13110, -0.39 DPS) [world_drop]; Nightscape Tunic (8175, -0.77 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.77 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.0 ranged_attack_power points (1.55 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 35.1 ranged_attack_power points (2.46 DPS) | yes | Gauntlets of Divinity (7724, -0.22 DPS) [dungeon]; Dragonscale Gauntlets (8347, -0.24 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.53 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (2.11 DPS) | yes | Scorpashi Sash (14652, -0.17 DPS) [world_drop]; Blackforge Girdle (6425, -0.37 DPS) [dungeon]; Ogron's Sash (13117, -0.37 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 57.8 ranged_attack_power points (4.06 DPS) | yes | Triprunner Dungarees (9624, -1.08 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.35 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.35 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 35.8 ranged_attack_power points (2.51 DPS) | yes | Imperial Leather Boots (6431, -0.39 DPS) [dungeon]; Dusky Boots (7390, -0.39 DPS) [crafted]; Worn Running Boots (9398, -0.39 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 27.5 ranged_attack_power points (1.93 DPS) | yes | Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Legionnaire's Band (19512, -0.39 DPS) [rep]; Disengagement Ring (276202, -0.39 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 24.8 ranged_attack_power points (1.74 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Legionnaire's Band (19512, -0.19 DPS) [rep]; Disengagement Ring (276202, -0.19 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (163.8 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -2.21 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Ginn-su Sword (9424), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.0 ranged_attack_power points (1.55 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Ginn-su Sword (9424, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (24.56 DPS) | yes | Outrider's Bow (19560, -1.67 DPS) [pvp]; Shadowforge Bushmaster (9422, -2.20 DPS) [dungeon]; The Silencer (13138, -3.87 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 228.0. Weights run: 2.6s. Verify run: 3.1s. 713 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.809 ± 0.015, crit=1.241 ± 0.031 per rating point (14 rating = 1%, 17.373 per %), hit=1.586 ± 0.082 per rating point (10 rating = 1%, 15.855 per %), melee_haste=10.926 ± 1.165

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 47.7 ranged_attack_power points (3.65 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.43 DPS) [quest]; Tough Scorpid Helm (8208, -0.64 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.5 ranged_attack_power points (2.79 DPS) | yes | Scout's Medallion (19535, -0.21 DPS) [rep]; Woven Ivy Necklace (19159, -0.86 DPS) [quest] |
| shoulder | Phytoskin Spaulders (17749) | Maraudon: Razorlash [dungeon] | 44.9 ranged_attack_power points (3.43 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Sunburn Spaulders (274751, -0.16 DPS) [vendor]; Khan's Mantle (14787, -0.86 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 42.1 ranged_attack_power points (3.22 DPS) | yes | Blackveil Cape (11626, -0.21 DPS) [dungeon]; Blackmetal Cape (9512, -0.64 DPS) [dungeon]; Blisterbane Wrap (12552, -1.24 DPS, sim-verified) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 70.2 ranged_attack_power points (5.36 DPS) | yes | Blazewind Breastplate (11193, -0.43 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.82 DPS) [vendor]; Quillward Harness (10583, -1.29 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 42.1 ranged_attack_power points (3.22 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Bloodlust Bracelets (14807, -1.09 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 61.8 ranged_attack_power points (4.72 DPS) | yes | Beastmaster's Gauntlets (226883, -1.72 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.72 DPS) [crafted]; First Sergeant's Chain Gauntlets (220830, -1.74 DPS, sim-verified) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 55.1 ranged_attack_power points (4.21 DPS) | yes | Sagebrush Girdle (17778, -0.99 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.21 DPS) [crafted]; Stalker's Mail Belt (252588, -1.21 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 59.0 ranged_attack_power points (4.50 DPS) | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS, sim-verified) [vendor]; Infernal Trickster Leggings (17754, -0.21 DPS) [dungeon]; Keeper's Woolies (14668, -0.43 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 56.2 ranged_attack_power points (4.29 DPS) | yes | Fleetfoot Greaves (11627, -0.21 DPS) [dungeon]; Elven Chain Boots (13125, -0.43 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.64 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.3 ranged_attack_power points (3.00 DPS) | yes | Ring of the Underwood (2951, -0.86 DPS) [world_drop]; Falcon's Hook (7552, -1.07 DPS) [dungeon]; Ironspine's Eye (7686, -1.07 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.9 ranged_attack_power points (2.74 DPS) | yes | Falcon's Hook (7552, -0.81 DPS) [dungeon]; Ironspine's Eye (7686, -0.81 DPS) [dungeon]; Ring of the Underwood (2951, -1.45 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (228.0 DPS) | yes | Frozen Heart of the Mountain (249469, -6.17 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (228.0 DPS) | yes | Frozen Heart of the Mountain (249469, -2.79 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (228.0 DPS) | yes | Warmonger (13052, -1.30 DPS) [world_drop]; Steel Spear (250605, -1.38 DPS) [crafted]; Hanzo Sword (8190, -5.81 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (228.0 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.68 DPS) [world_drop]; Hurricane (2824, -1.86 DPS) [world_drop]; Dark Iron Rifle (16004, -3.63 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Phytoskin Spaulders; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 314.5. Weights run: 2.5s. Verify run: 10.2s. 1650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.110 ± 0.026, crit=2.158 ± 0.054 per rating point (14 rating = 1%, 30.209 per %), hit=3.513 ± 0.145 per rating point (10 rating = 1%, 35.129 per %), melee_haste=19.852 ± 2.143

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (314.5 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -9.16 DPS, sim-verified) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 81.8 ranged_attack_power points (6.20 DPS) | yes | Amulet of the Darkmoon (19491, -1.72 DPS) [quest]; Mark of Fordring (15411, -1.94 DPS) [quest]; Beads of Ogre Might (22150, -2.09 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 114.2 ranged_attack_power points (8.65 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -0.80 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.7 ranged_attack_power points (5.05 DPS) | yes | Shifting Cloak (18511, -1.04 DPS) [crafted]; Shadow Prowler's Cloak (22269, -1.04 DPS) [dungeon]; Howler's Furs (272414, -2.55 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (314.5 DPS) | yes | Warlord's Chain Chestpiece (16565, -3.37 DPS) [vendor]; Warlord's Chain Hauberk (231573, -3.37 DPS) [vendor]; Tunic of Undead Slaying (23089, -15.80 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (314.5 DPS) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -8.56 DPS, sim-verified) [rep] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-verified (314.5 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Chain Gloves (16571, +0.00 DPS) [vendor]; Raider Gloves (272099, -5.17 DPS, sim-verified) [vendor] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 100.4 ranged_attack_power points (7.61 DPS) | yes | Belt of Preserved Heads (20216, -1.81 DPS, sim-verified) [quest]; Ranger's Belt (272397, -1.95 DPS) [vendor]; Warpwood Binding (18393, -2.02 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 204.4 ranged_attack_power points (15.49 DPS) | yes | Outrider's Chain Leggings (22673, -2.61 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -4.55 DPS) [vendor]; General's Chain Legguards (231574, -5.19 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (314.5 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Windreaver Greaves (13967, -9.11 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (314.5 DPS) | yes | Don Julio's Band (19325, -0.74 DPS) [rep]; Cutthroat's Signet (272408, -0.94 DPS) [vendor]; Naglering (11669, -7.71 DPS, sim-verified) [dungeon] |
| finger2 | Tarnished Elven Ring (18500) | Dire Maul: Tribute [dungeon] | sim-verified (314.5 DPS) | yes | Don Julio's Band (19325, -0.03 DPS) [rep]; Cutthroat's Signet (272408, -0.24 DPS) [vendor]; Naglering (11669, -6.84 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (314.5 DPS) | yes | Blackhand's Breadth (13965, -3.65 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.83 DPS) [crafted]; Counterattack Lodestone (18537, -6.56 DPS) [dungeon] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (314.5 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (314.5 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -9.55 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (314.5 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -11.65 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Marksman's Girdle; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Tarnished Elven Ring; trinket2: Burst of Knowledge; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 785.4. Weights run: 2.6s. Verify run: 10.7s. 1650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.129 ± 0.025, crit=2.277 ± 0.051 per rating point (14 rating = 1%, 31.876 per %), hit=3.998 ± 0.199 per rating point (10 rating = 1%, 39.984 per %), melee_haste=21.245 ± 2.668

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (785.4 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -13.33 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 86.9 ranged_attack_power points (14.36 DPS) | yes | Amulet of the Darkmoon (19491, -4.54 DPS) [quest]; Beads of Ogre Might (22150, -4.61 DPS, sim-verified) [quest]; Mark of Fordring (15411, -4.80 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 116.4 ranged_attack_power points (19.23 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.27 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (785.4 DPS) | yes | Shifting Cloak (18511, -2.27 DPS) [crafted]; Shadow Prowler's Cloak (22269, -2.27 DPS) [dungeon]; Howler's Furs (272414, -7.98 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (785.4 DPS) | yes | Warlord's Chain Chestpiece (16565, -8.16 DPS) [vendor]; Warlord's Chain Hauberk (231573, -8.16 DPS) [vendor]; Tunic of Undead Slaying (23089, -30.39 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (785.4 DPS) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -12.99 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | 108.1 ranged_attack_power points (17.87 DPS) | yes | General's Chain Gloves (16571, -1.74 DPS) [vendor]; General's Chain Vices (231575, -1.74 DPS) [vendor]; Gauntlets of Accuracy (18349, -3.50 DPS) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (785.4 DPS) | yes | Warpwood Binding (18393, +0.00 DPS) [dungeon]; Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, -9.74 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 213.3 ranged_attack_power points (35.24 DPS) | yes | Outrider's Chain Leggings (22673, -5.71 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -10.74 DPS) [vendor]; General's Chain Legguards (231574, -12.39 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (785.4 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -12.77 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (785.4 DPS) | yes | Tarnished Elven Ring (18500, -1.55 DPS) [dungeon]; Cutthroat's Signet (272408, -2.07 DPS) [vendor]; Naglering (11669, -13.17 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (785.4 DPS) | yes | Tarnished Elven Ring (18500, -0.15 DPS) [dungeon]; Cutthroat's Signet (272408, -0.67 DPS) [vendor]; Naglering (11669, -10.57 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (785.4 DPS) | yes | Frozen Heart of the Mountain (249469, -12.56 DPS) [crafted]; Counterattack Lodestone (18537, -14.87 DPS) [dungeon]; Hand of Justice (11815, -15.20 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (785.4 DPS) | yes | Frozen Heart of the Mountain (249469, -4.59 DPS) [crafted]; Burst of Knowledge (11832, -6.35 DPS, sim-verified) [dungeon]; Counterattack Lodestone (18537, -6.90 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (785.4 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -16.81 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (785.4 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -28.59 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

