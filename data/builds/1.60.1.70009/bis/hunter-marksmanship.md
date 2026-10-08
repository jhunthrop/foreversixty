# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 79.3. Weights run: 2.2s. Verify run: 2.0s. 220 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.143 ± 0.007, crit=0.545 ± 0.014 per rating point (14 rating = 1%, 7.633 per %), hit=0.840 ± 0.032 per rating point (10 rating = 1%, 8.403 per %), melee_haste=not significant (-1.079 ± 0.660)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.1 ranged_attack_power points (1.10 DPS) | yes | Red Winter Hat (21524, -1.05 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.9 ranged_attack_power points (0.82 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS, sim-verified) [quest] |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent's Shoulders (5404, -0.95 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.9 ranged_attack_power points (0.82 DPS) | yes | Cape of the Brotherhood (5193, +0.00 DPS, sim-verified) [dungeon]; Hide of Lupos (3018, -0.27 DPS) [world]; Bristlebark Cape (14571, -0.27 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 23.6 ranged_attack_power points (1.51 DPS) | yes | Brawler's Leather Armor (252490, -0.23 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.55 DPS) [crafted]; Dark Leather Tunic (2317, -0.69 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.7 ranged_attack_power points (0.69 DPS) | yes | Wolf Bracers (4794, -0.14 DPS) [vendor]; Bravo's Armbands (270015, -0.14 DPS) [quest]; Ratchet Wristwraps (274742, -0.27 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.27 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.27 DPS) [crafted]; Serpent Gloves (5970, -1.60 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.15 DPS) | yes | Deviate Scale Belt (6468, -0.47 DPS) [crafted]; Dusty Belt (279897, -0.47 DPS) [quest]; Dark Leather Belt (4249, -0.60 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.3 ranged_attack_power points (1.23 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.14 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Agile Boots (4788, -0.14 DPS) [vendor]; Feet of the Lynx (1121, -1.37 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.9 ranged_attack_power points (0.82 DPS) | yes | Lavishly Jeweled Ring (1156, -0.55 DPS) [dungeon]; The 1 Ring (8350, -0.69 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.6 ranged_attack_power points (0.55 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.41 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Scythe Axe (5749, -0.27 DPS) [world]; Lupine Axe (1220, -0.41 DPS) [world]; Bronze Dory (250603, -0.82 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (11.41 DPS) | yes | Lil Timmy's Peashooter (13136, -1.48 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.46 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.71 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Slime-encrusted Pads; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 97.3. Weights run: 2.5s. Verify run: 2.5s. 366 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.192 ± 0.009, crit=0.734 ± 0.019 per rating point (14 rating = 1%, 10.273 per %), hit=0.941 ± 0.041 per rating point (10 rating = 1%, 9.405 per %), melee_haste=4.953 ± 0.730

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.9 ranged_attack_power points (1.42 DPS) | yes | Tribal Worg Helm (6204, -0.28 DPS) [world]; Brawler's Leather Hood (252504, -0.29 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.43 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 17.5 ranged_attack_power points (1.14 DPS) | yes | Ghostshard Talisman (7731, -0.23 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.57 DPS) [world_drop]; Erudite's Amulet (277204, -0.57 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 ranged_attack_power points (1.56 DPS) | yes | Dark Leather Shoulders (4252, -0.57 DPS) [crafted]; Insignia Mantle (4721, -0.57 DPS) [world_drop]; Mantle of Thieves (2264, -0.66 DPS, sim-verified) [dungeon] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.5 ranged_attack_power points (1.14 DPS) | yes | Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.28 DPS) [world]; Fenrus' Hide (6340, -0.28 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.7 ranged_attack_power points (1.99 DPS) | yes | Tunic of Westfall (2041, -0.43 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.85 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.85 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 13.2 ranged_attack_power points (0.85 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.14 DPS) [world_drop] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 17.5 ranged_attack_power points (1.14 DPS) | yes | Heavy Earthen Gloves (7359, -0.10 DPS) [crafted]; Serpent Gloves (5970, -0.28 DPS) [dungeon]; Insignia Gloves (6408, -0.28 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 ranged_attack_power points (1.55 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted]; Stalker's Leather Belt (252521, -0.28 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 30.7 ranged_attack_power points (1.99 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.14 DPS) [crafted]; Ferine Leggings (6690, -0.30 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 17.5 ranged_attack_power points (1.14 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.14 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.7 ranged_attack_power points (1.28 DPS) | yes | Ring of Precision (1491, -0.43 DPS) [dungeon]; Protector's Band (19517, -0.43 DPS) [rep]; Signet of the Zhevra (285330, -0.43 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.3 ranged_attack_power points (0.99 DPS) | yes | Protector's Band (19517, -0.14 DPS) [rep]; Signet of the Zhevra (285330, -0.14 DPS) [world]; Ring of Precision (1491, -0.50 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 17.0 ranged_attack_power points (1.10 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.2 ranged_attack_power points (0.85 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (16.94 DPS) | yes | Glass Shooter (9456, -0.68 DPS) [dungeon]; Ironweaver (13137, -1.24 DPS) [world_drop]; Silver Star (3463, -1.25 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Highlander's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 128.0. Weights run: 2.5s. Verify run: 2.6s. 597 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.235 ± 0.010, crit=0.858 ± 0.023 per rating point (14 rating = 1%, 12.009 per %), hit=1.191 ± 0.053 per rating point (10 rating = 1%, 11.913 per %), melee_haste=7.263 ± 0.872

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 29.1 ranged_attack_power points (2.00 DPS) | yes | Guard's Chain Helm (250499, -0.15 DPS) [crafted]; Skullsplitter Helm (1624, -0.31 DPS) [world]; Nightscape Headband (8176, -0.64 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 24.6 ranged_attack_power points (1.69 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.32 DPS) [quest]; Ghostshard Talisman (7731, -0.73 DPS) [dungeon]; Erudite's Amulet (277204, -1.08 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.6 ranged_attack_power points (2.52 DPS) | yes | Forest Tracker Epaulets (2278, -0.83 DPS) [world_drop]; Nightscape Shoulders (8192, -0.86 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 22.3 ranged_attack_power points (1.54 DPS) | yes | Imperial Cloak (6432, -0.31 DPS) [dungeon]; Parachute Cloak (10518, -0.31 DPS) [crafted]; Tigerstrike Mantle (13108, -0.31 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 42.5 ranged_attack_power points (2.93 DPS) | yes | Wolffear Harness (13110, -0.31 DPS) [world_drop]; Nightscape Tunic (8175, -0.62 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.62 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.38 DPS) | yes | Imperial Leather Bracers (4061, -0.15 DPS) [dungeon]; Dusky Bracers (7378, -0.15 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.30 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 32.0 ranged_attack_power points (2.21 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Dragonscale Gauntlets (8347, -0.45 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.67 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (2.07 DPS) | yes | Highlander's Chain Girdle (20090, -0.43 DPS, sim-verified) [rep]; Scorpashi Sash (14652, -0.53 DPS) [world_drop]; Blackforge Girdle (6425, -0.68 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.9 ranged_attack_power points (3.24 DPS) | yes | Triprunner Dungarees (9624, -0.71 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.08 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.08 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 29.1 ranged_attack_power points (2.00 DPS) | yes | Imperial Leather Boots (6431, -0.31 DPS) [dungeon]; Dusky Boots (7390, -0.31 DPS) [crafted]; Worn Running Boots (9398, -0.31 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 22.3 ranged_attack_power points (1.54 DPS) | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Mark of Kern (2262, -0.16 DPS) [dungeon]; Assault Band (13095, -0.16 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 20.1 ranged_attack_power points (1.39 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Mark of Kern (2262, -0.01 DPS) [dungeon]; Assault Band (13095, -0.01 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (128.0 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.57 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 20.1 ranged_attack_power points (1.39 DPS) | yes | Blue Glittering Axe (7942, -0.93 DPS, sim-verified) [crafted]; Satyr's Rod (15962, -1.23 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (24.13 DPS) | yes | Shadowforge Bushmaster (9422, -2.16 DPS) [dungeon]; Swiftwind (13038, -2.54 DPS) [world_drop]; The Silencer (13138, -3.14 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 177.9. Weights run: 2.6s. Verify run: 5.1s. 754 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.267 ± 0.012, crit=0.997 ± 0.025 per rating point (14 rating = 1%, 13.954 per %), hit=1.307 ± 0.068 per rating point (10 rating = 1%, 13.074 per %), melee_haste=8.625 ± 0.942

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 50.2 ranged_attack_power points (3.76 DPS) | yes | Lordrec Helmet (10741, -1.05 DPS) [quest]; Bloomsprout Headpiece (17767, -1.07 DPS) [dungeon]; Helm of Fire (8348, -5.29 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.5 ranged_attack_power points (2.21 DPS) | yes | Sentinel's Medallion (19539, -0.17 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.71 DPS) [quest] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 43.4 ranged_attack_power points (3.25 DPS) | yes | Phytoskin Spaulders (17749, -0.54 DPS) [dungeon]; Khan's Mantle (14787, -1.22 DPS) [world_drop]; Sunburn Spaulders (274751, -3.51 DPS, sim-verified) [vendor] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 34.0 ranged_attack_power points (2.55 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.17 DPS) [dungeon]; Blackmetal Cape (9512, -0.51 DPS) [dungeon] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | sim-verified (177.9 DPS) | yes | Blazewind Breastplate (11193, +0.00 DPS) [quest]; Quillward Harness (10583, -0.37 DPS) [dungeon]; Fungus Shroud Armor (17742, -3.51 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.0 ranged_attack_power points (2.55 DPS) | yes | Bracers of the Stone Princess (17714, -0.45 DPS) [dungeon]; Arena Bands (18711, -0.61 DPS, sim-verified) [world]; Bloodlust Bracelets (14807, -0.68 DPS) [world_drop] |
| hands | Sergeant Major's Chain Gauntlets (220829) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (177.9 DPS) | yes | Gloves of Holy Might (867, -0.17 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.32 DPS) [dungeon]; Raider Gloves (272100, -1.14 DPS, sim-verified) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 47.0 ranged_attack_power points (3.52 DPS) | yes | Sagebrush Girdle (17778, -0.97 DPS) [quest]; Highlander's Chain Girdle (20088, -0.98 DPS) [rep]; Highlander's Leather Girdle (20115, -0.98 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | sim-verified (177.9 DPS) | yes | Infernal Trickster Leggings (17754, -0.03 DPS) [dungeon]; Keeper's Woolies (14668, -0.20 DPS) [world_drop]; Basilisk Hide Pants (1718, -4.27 DPS, sim-verified) [world_drop] |
| feet | Sergeant Major's Chain Sabatons (220837) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (177.9 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS) [dungeon]; Elven Chain Boots (13125, +0.00 DPS) [world_drop]; Albino Crocscale Boots (17728, -1.55 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 33.1 ranged_attack_power points (2.48 DPS) | yes | Ring of the Underwood (2951, -0.78 DPS) [world_drop]; Falcon's Hook (7552, -0.95 DPS) [dungeon]; Ironspine's Eye (7686, -0.95 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.7 ranged_attack_power points (2.38 DPS) | yes | Ring of the Underwood (2951, -0.80 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.85 DPS) [dungeon]; Ironspine's Eye (7686, -0.85 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (177.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (177.9 DPS) | yes | Molten Heart of the Mountain (249470, -1.17 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (177.9 DPS) | yes | Steel Spear (250605, -0.95 DPS) [crafted]; Warmonger (13052, -0.97 DPS) [world_drop]; Hanzo Sword (8190, -4.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (177.9 DPS) | yes | Hurricane (2824, -1.82 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -2.22 DPS) [world_drop]; Dark Iron Rifle (16004, -4.09 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; hands: Sergeant Major's Chain Gauntlets; waist: Substandard Belt Chain; legs: Knight's Chain Legplates; feet: Sergeant Major's Chain Sabatons; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 238.0. Weights run: 2.6s. Verify run: 10.5s. 1670 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.523 ± 0.021, crit=1.912 ± 0.048 per rating point (14 rating = 1%, 26.765 per %), hit=2.727 ± 0.122 per rating point (10 rating = 1%, 27.266 per %), melee_haste=18.676 ± 1.862

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (238.0 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -7.55 DPS, sim-verified) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 65.1 ranged_attack_power points (4.84 DPS) | yes | Mark of Fordring (15411, -0.92 DPS) [quest]; Beads of Ogre Might (22150, -1.03 DPS) [quest]; Medallion of the Dawn (22659, -1.07 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.9 ranged_attack_power points (7.05 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Field Marshal's Chain Spaulders (16468, -0.15 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -0.15 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 57.8 ranged_attack_power points (4.30 DPS) | yes | Cloak of the Honor Guard (20073, -0.83 DPS) [rep]; Shadow Prowler's Cloak (22269, -1.11 DPS) [dungeon]; Howler's Furs (272414, -2.74 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (238.0 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.59 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.59 DPS) [vendor]; Tunic of Undead Slaying (23089, -11.61 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (238.0 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Bracers of the Eclipse (18375, -6.59 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 77.2 ranged_attack_power points (5.74 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; Marshal's Chain Grips (231560, -0.24 DPS) [pvp] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (238.0 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Marksman's Girdle (22232, -6.23 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 169.1 ranged_attack_power points (12.57 DPS) | yes | Sentinel's Leather Pants (237818, -3.53 DPS) [vendor]; Marshal's Chain Legguards (231577, -4.20 DPS) [pvp]; Marshal's Chain Legplates (231558, -4.72 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (238.0 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Windreaver Greaves (13967, -7.88 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (238.0 DPS) | yes | Tarnished Elven Ring (18500, -0.56 DPS) [dungeon]; Cutthroat's Signet (272408, -0.75 DPS) [vendor]; Naglering (11669, -5.18 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (238.0 DPS) | yes | Tarnished Elven Ring (18500, -0.37 DPS) [dungeon]; Cutthroat's Signet (272408, -0.55 DPS) [vendor]; Naglering (11669, -4.65 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (238.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (238.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (238.0 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -7.62 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (238.0 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.84 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 656.9. Weights run: 2.6s. Verify run: 11.3s. 1670 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.553 ± 0.021, crit=2.069 ± 0.046 per rating point (14 rating = 1%, 28.968 per %), hit=3.691 ± 0.188 per rating point (10 rating = 1%, 36.912 per %), melee_haste=20.493 ± 2.386

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (656.9 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -12.69 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 75.2 ranged_attack_power points (12.33 DPS) | yes | Beads of Ogre Might (22150, -2.45 DPS, sim-verified) [quest]; Mark of Fordring (15411, -3.32 DPS) [quest]; Medallion of the Dawn (22659, -3.65 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 97.9 ranged_attack_power points (16.05 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 64.9 ranged_attack_power points (10.64 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -2.97 DPS) [rep]; Shadow Prowler's Cloak (22269, -3.52 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (656.9 DPS) | yes | Field Marshal's Chain Breastplate (16466, -7.30 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -7.30 DPS) [vendor]; Tunic of Undead Slaying (23089, -23.20 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (656.9 DPS) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -12.23 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 96.8 ranged_attack_power points (15.86 DPS) | yes | Gauntlets of Accuracy (18349, +0.00 DPS, sim-verified) [dungeon]; Marshal's Chain Vices (231578, -2.33 DPS) [vendor]; Marshal's Chain Grips (231560, -3.33 DPS) [pvp] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (656.9 DPS) | yes | Warpwood Binding (18393, +0.00 DPS) [dungeon]; Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, -11.11 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 184.2 ranged_attack_power points (30.19 DPS) | yes | Sentinel's Leather Pants (237818, -9.40 DPS) [vendor]; Marshal's Chain Legguards (231577, -11.22 DPS) [pvp]; Plaguehound Leggings (18736, -11.59 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (656.9 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -13.15 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (656.9 DPS) | yes | Tarnished Elven Ring (18500, -1.26 DPS) [dungeon]; Cutthroat's Signet (272408, -1.67 DPS) [vendor]; Naglering (11669, -8.92 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (656.9 DPS) | yes | Tarnished Elven Ring (18500, -1.09 DPS) [dungeon]; Cutthroat's Signet (272408, -1.51 DPS) [vendor]; Naglering (11669, -7.98 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (656.9 DPS) | yes | Frozen Heart of the Mountain (249469, -4.05 DPS) [crafted]; Counterattack Lodestone (18537, -5.89 DPS) [dungeon]; Hand of Justice (11815, -6.22 DPS) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (656.9 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Burst of Knowledge (11832, -2.84 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (656.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -14.12 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (656.9 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -31.67 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Blackhand's Breadth; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 79.9. Weights run: 2.2s. Verify run: 3.9s. 209 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.143 ± 0.007, crit=0.545 ± 0.014 per rating point (14 rating = 1%, 7.633 per %), hit=0.840 ± 0.032 per rating point (10 rating = 1%, 8.403 per %), melee_haste=not significant (-1.079 ± 0.660)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.1 ranged_attack_power points (1.10 DPS) | yes | Red Winter Hat (21524, -1.10 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.9 ranged_attack_power points (0.82 DPS) | yes | Erudite's Amulet (277204, -0.27 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.7 ranged_attack_power points (0.69 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.9 ranged_attack_power points (0.82 DPS) | yes | Hide of Lupos (3018, -0.27 DPS) [world]; Bristlebark Cape (14571, -0.27 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.44 DPS, sim-verified) [dungeon] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Dark Leather Tunic (2317, -0.14 DPS) [crafted]; Prospector's Chestpiece (14562, -0.14 DPS) [world_drop]; Brawler's Leather Armor (252490, -1.21 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.7 ranged_attack_power points (0.69 DPS) | yes | Wolf Bracers (4794, -0.14 DPS) [vendor]; Bristlebark Bindings (14569, -0.27 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.27 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Leather Gloves (3058, -0.27 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.27 DPS) [crafted]; Serpent Gloves (5970, -0.59 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.15 DPS) | yes | Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.60 DPS) [world]; Dark Leather Belt (4249, -0.60 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.3 ranged_attack_power points (1.23 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.14 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Agile Boots (4788, -0.14 DPS) [vendor]; Feet of the Lynx (1121, -0.35 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.9 ranged_attack_power points (0.82 DPS) | yes | Bounty Hunter's Ring (5351, -0.41 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.55 DPS) [dungeon]; The 1 Ring (8350, -0.69 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.6 ranged_attack_power points (0.55 DPS) | yes | Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon]; The 1 Ring (8350, -0.41 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Scythe Axe (5749, -0.27 DPS) [world]; Crescent Staff (6505, -0.27 DPS) [quest]; Bronze Dory (250603, -0.83 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (11.41 DPS) | yes | Outrider's Bow (20437, -0.69 DPS) [pvp]; Lil Timmy's Peashooter (13136, -1.02 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.46 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 97.9. Weights run: 2.5s. Verify run: 2.4s. 352 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.192 ± 0.009, crit=0.734 ± 0.019 per rating point (14 rating = 1%, 10.273 per %), hit=0.941 ± 0.041 per rating point (10 rating = 1%, 9.405 per %), melee_haste=4.953 ± 0.730

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.9 ranged_attack_power points (1.42 DPS) | yes | Brawler's Leather Hood (252504, -0.28 DPS, sim-verified) [crafted]; Tribal Worg Helm (6204, -0.28 DPS) [world]; Humbert's Helm (4724, -0.43 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 17.5 ranged_attack_power points (1.14 DPS) | yes | Ghostshard Talisman (7731, -0.23 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.57 DPS) [world_drop]; Erudite's Amulet (277204, -0.57 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 ranged_attack_power points (1.56 DPS) | yes | Mantle of Thieves (2264, +0.00 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.57 DPS) [crafted]; Insignia Mantle (4721, -0.57 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.5 ranged_attack_power points (1.14 DPS) | yes | Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.28 DPS) [world]; Swiftrunner Cape (6745, -0.28 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.7 ranged_attack_power points (1.99 DPS) | yes | Panther Armor (6670, -0.72 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.85 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.85 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 13.2 ranged_attack_power points (0.85 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Loamflake Bracers (15462, -0.14 DPS) [quest] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 17.5 ranged_attack_power points (1.14 DPS) | yes | Heavy Earthen Gloves (7359, -0.10 DPS) [crafted]; Braced Handguards (6784, -0.14 DPS) [quest]; Insignia Gloves (6408, -0.28 DPS) [world_drop] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 ranged_attack_power points (1.55 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Deftkin Belt (16659, -0.21 DPS) [quest]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 30.7 ranged_attack_power points (1.99 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.14 DPS) [crafted]; Ferine Leggings (6690, -0.30 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 17.5 ranged_attack_power points (1.14 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.7 ranged_attack_power points (1.28 DPS) | yes | Ring of Precision (1491, -0.43 DPS) [dungeon]; Legionnaire's Band (19513, -0.43 DPS) [rep]; Signet of the Zhevra (285330, -0.43 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.3 ranged_attack_power points (0.99 DPS) | yes | Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep]; Signet of the Zhevra (285330, -0.14 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 17.0 ranged_attack_power points (1.10 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Outlaw Sabre (16886) | Baron Aquanis [quest] | 15.0 ranged_attack_power points (0.97 DPS) | yes | Vendetta (776, +0.00 DPS) [dungeon]; Satyr's Rod (15962, -0.83 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (16.94 DPS) | yes | Glass Shooter (9456, -0.68 DPS) [dungeon]; Ironweaver (13137, -1.24 DPS) [world_drop]; Silver Star (3463, -1.34 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Defiler's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Outlaw Sabre; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 129.9. Weights run: 2.5s. Verify run: 2.5s. 563 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.235 ± 0.010, crit=0.858 ± 0.023 per rating point (14 rating = 1%, 12.009 per %), hit=1.191 ± 0.053 per rating point (10 rating = 1%, 11.913 per %), melee_haste=7.263 ± 0.872

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 29.1 ranged_attack_power points (2.00 DPS) | yes | Guard's Chain Helm (250499, -0.15 DPS) [crafted]; Skullsplitter Helm (1624, -0.31 DPS) [world]; Nightscape Headband (8176, -1.04 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 24.6 ranged_attack_power points (1.69 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.32 DPS) [quest]; Ghostshard Talisman (7731, -0.73 DPS) [dungeon]; Erudite's Amulet (277204, -1.08 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.6 ranged_attack_power points (2.52 DPS) | yes | Forest Tracker Epaulets (2278, -0.83 DPS) [world_drop]; Nightscape Shoulders (8192, -0.87 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 22.3 ranged_attack_power points (1.54 DPS) | yes | Imperial Cloak (6432, -0.31 DPS) [dungeon]; Parachute Cloak (10518, -0.31 DPS) [crafted]; Tigerstrike Mantle (13108, -0.31 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 42.5 ranged_attack_power points (2.93 DPS) | yes | Wolffear Harness (13110, -0.47 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.62 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.62 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.38 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Dusky Bracers (7378, -0.15 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 32.0 ranged_attack_power points (2.21 DPS) | yes | Dragonscale Gauntlets (8347, -0.45 DPS) [crafted]; Gauntlets of Divinity (7724, -0.49 DPS, sim-verified) [dungeon]; Tough Scorpid Gloves (8204, -0.67 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (2.07 DPS) | yes | Defiler's Chain Girdle (20152, -0.41 DPS) [rep]; Scorpashi Sash (14652, -0.53 DPS) [world_drop]; Deftkin Belt (16659, -0.62 DPS) [quest] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.9 ranged_attack_power points (3.24 DPS) | yes | Petrolspill Leggings (9509, -1.08 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.08 DPS) [world_drop]; Triprunner Dungarees (9624, -1.10 DPS, sim-verified) [quest] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 29.1 ranged_attack_power points (2.00 DPS) | yes | Imperial Leather Boots (6431, -0.31 DPS) [dungeon]; Dusky Boots (7390, -0.31 DPS) [crafted]; Worn Running Boots (9398, -0.31 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 22.3 ranged_attack_power points (1.54 DPS) | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Mark of Kern (2262, -0.16 DPS) [dungeon]; Assault Band (13095, -0.16 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 20.1 ranged_attack_power points (1.39 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Mark of Kern (2262, -0.01 DPS) [dungeon]; Assault Band (13095, -0.01 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (129.9 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.59 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 20.1 ranged_attack_power points (1.39 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Satyr's Rod (15962, -1.23 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (24.13 DPS) | yes | Outrider's Bow (19560, -1.75 DPS) [pvp]; Shadowforge Bushmaster (9422, -2.16 DPS) [dungeon]; The Silencer (13138, -3.75 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 180.2. Weights run: 2.6s. Verify run: 3.1s. 713 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.267 ± 0.012, crit=0.997 ± 0.025 per rating point (14 rating = 1%, 13.954 per %), hit=1.307 ± 0.068 per rating point (10 rating = 1%, 13.074 per %), melee_haste=8.625 ± 0.942

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 38.5 ranged_attack_power points (2.89 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.19 DPS) [dungeon]; Sprightring Helm (17776, -0.34 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.5 ranged_attack_power points (2.21 DPS) | yes | Scout's Medallion (19535, -0.17 DPS) [rep]; Woven Ivy Necklace (19159, -0.68 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.9 ranged_attack_power points (2.77 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -0.05 DPS) [dungeon]; Khan's Mantle (14787, -0.73 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 34.0 ranged_attack_power points (2.55 DPS) | yes | Blackveil Cape (11626, -0.17 DPS) [dungeon]; Blackmetal Cape (9512, -0.51 DPS) [dungeon]; Blisterbane Wrap (12552, -0.77 DPS, sim-verified) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 56.7 ranged_attack_power points (4.25 DPS) | yes | Blazewind Breastplate (11193, -0.34 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.65 DPS) [vendor]; Quillward Harness (10583, -1.02 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.0 ranged_attack_power points (2.55 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, -0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 49.9 ranged_attack_power points (3.74 DPS) | yes | Gloves of Holy Might (867, -1.19 DPS) [world_drop]; Gauntlets of Divinity (7724, -1.34 DPS) [dungeon]; First Sergeant's Chain Gauntlets (220830, -1.34 DPS, sim-verified) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 47.0 ranged_attack_power points (3.52 DPS) | yes | Sagebrush Girdle (17778, -0.97 DPS) [quest]; Defiler's Chain Girdle (20151, -0.98 DPS) [rep]; Defiler's Leather Girdle (20193, -0.98 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 47.6 ranged_attack_power points (3.57 DPS) | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS, sim-verified) [vendor]; Infernal Trickster Leggings (17754, -0.17 DPS) [dungeon]; Keeper's Woolies (14668, -0.34 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 45.3 ranged_attack_power points (3.40 DPS) | yes | Fleetfoot Greaves (11627, -0.17 DPS) [dungeon]; Elven Chain Boots (13125, -0.34 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.51 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 33.1 ranged_attack_power points (2.48 DPS) | yes | White Bone Band (11862, -0.68 DPS) [quest]; Ring of the Underwood (2951, -0.78 DPS) [world_drop]; Falcon's Hook (7552, -0.95 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.7 ranged_attack_power points (2.38 DPS) | yes | Ring of the Underwood (2951, -0.68 DPS) [world_drop]; White Bone Band (11862, -0.71 DPS, sim-verified) [quest]; Falcon's Hook (7552, -0.85 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (180.2 DPS) | yes | Frozen Heart of the Mountain (249469, -6.10 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (180.2 DPS) | yes | Frozen Heart of the Mountain (249469, -2.63 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (180.2 DPS) | yes | Steel Spear (250605, -0.95 DPS) [crafted]; Warmonger (13052, -0.97 DPS) [world_drop]; Hanzo Sword (8190, -4.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (180.2 DPS) | yes | Hurricane (2824, -1.82 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -2.22 DPS) [world_drop]; Dark Iron Rifle (16004, -3.05 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 242.9. Weights run: 2.6s. Verify run: 10.3s. 1650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.523 ± 0.021, crit=1.912 ± 0.048 per rating point (14 rating = 1%, 26.765 per %), hit=2.727 ± 0.122 per rating point (10 rating = 1%, 27.266 per %), melee_haste=18.676 ± 1.862

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (242.9 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -8.44 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 65.1 ranged_attack_power points (4.84 DPS) | yes | Beads of Ogre Might (22150, -1.03 DPS) [quest]; Medallion of the Dawn (22659, -1.07 DPS) [quest]; Mark of Fordring (15411, -1.61 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.9 ranged_attack_power points (7.05 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -0.15 DPS) [pvp]; Warlord's Chain Pauldrons (231565, -0.51 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 57.8 ranged_attack_power points (4.30 DPS) | yes | Howler's Furs (272414, -0.19 DPS) [vendor]; Deathguard's Cloak (20068, -0.83 DPS) [rep]; Shadow Prowler's Cloak (22269, -1.11 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (242.9 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.59 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.59 DPS) [vendor]; Tunic of Undead Slaying (23089, -13.21 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (242.9 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Bracers of the Eclipse (18375, +0.00 DPS) [dungeon]; Beaststalker's Bindings (16681, -7.14 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (242.9 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (242.9 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -6.63 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 169.1 ranged_attack_power points (12.57 DPS) | yes | Outrider's Chain Leggings (22673, -2.15 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -3.53 DPS) [vendor]; General's Chain Legguards (231574, -4.20 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (242.9 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -6.12 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (242.9 DPS) | yes | Tarnished Elven Ring (18500, -0.56 DPS) [dungeon]; Cutthroat's Signet (272408, -0.75 DPS) [vendor]; Naglering (11669, -7.00 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (242.9 DPS) | yes | Tarnished Elven Ring (18500, -0.37 DPS) [dungeon]; Cutthroat's Signet (272408, -0.55 DPS) [vendor]; Naglering (11669, -6.44 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (242.9 DPS) | yes | Frozen Heart of the Mountain (249469, -5.84 DPS) [crafted]; Counterattack Lodestone (18537, -6.03 DPS) [dungeon]; Hand of Justice (11815, -6.18 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (242.9 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS, sim-verified) [dungeon]; Frozen Heart of the Mountain (249469, -2.16 DPS) [crafted]; Counterattack Lodestone (18537, -2.34 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (242.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -9.33 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (242.9 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.18 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 675.7. Weights run: 2.6s. Verify run: 10.8s. 1650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.553 ± 0.021, crit=2.069 ± 0.046 per rating point (14 rating = 1%, 28.968 per %), hit=3.691 ± 0.188 per rating point (10 rating = 1%, 36.912 per %), melee_haste=20.493 ± 2.386

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (675.7 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -19.73 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 75.2 ranged_attack_power points (12.33 DPS) | yes | Beads of Ogre Might (22150, -2.61 DPS, sim-verified) [quest]; Mark of Fordring (15411, -3.32 DPS) [quest]; Medallion of the Dawn (22659, -3.65 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 97.9 ranged_attack_power points (16.05 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 64.9 ranged_attack_power points (10.64 DPS) | yes | Cape of the Black Baron (13340, -1.08 DPS) [dungeon]; Deathguard's Cloak (20068, -2.97 DPS) [rep]; Shadow Prowler's Cloak (22269, -3.52 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (675.7 DPS) | yes | Warlord's Chain Chestpiece (16565, -7.30 DPS) [vendor]; Warlord's Chain Hauberk (231573, -7.30 DPS) [vendor]; Tunic of Undead Slaying (23089, -29.99 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (675.7 DPS) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -14.50 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (675.7 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -4.77 DPS, sim-verified) [quest] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (675.7 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -13.69 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 184.2 ranged_attack_power points (30.19 DPS) | yes | Outrider's Chain Leggings (22673, -4.69 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -9.40 DPS) [vendor]; General's Chain Legguards (231574, -11.22 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (675.7 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -15.39 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (675.7 DPS) | yes | Tarnished Elven Ring (18500, -1.26 DPS) [dungeon]; Cutthroat's Signet (272408, -1.67 DPS) [vendor]; Naglering (11669, -15.92 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (675.7 DPS) | yes | Tarnished Elven Ring (18500, -1.09 DPS) [dungeon]; Cutthroat's Signet (272408, -1.51 DPS) [vendor]; Naglering (11669, -14.72 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (675.7 DPS) | yes | Frozen Heart of the Mountain (249469, -12.56 DPS) [crafted]; Counterattack Lodestone (18537, -14.40 DPS) [dungeon]; Hand of Justice (11815, -14.72 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (675.7 DPS) | yes | Devilsaur Eye (19991, -2.22 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -4.05 DPS) [crafted]; Counterattack Lodestone (18537, -5.89 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (675.7 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -20.81 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (675.7 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -29.59 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

