# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 82.1. Weights run: 2.1s. Verify run: 1.9s. 220 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.148 ± 0.007, crit=0.502 ± 0.014 per rating point (14 rating = 1%, 7.027 per %), hit=0.867 ± 0.033 per rating point (10 rating = 1%, 8.669 per %), melee_haste=6.773 ± 0.552

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.2 ranged_attack_power points (1.16 DPS) | yes | Red Winter Hat (21524, -1.13 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.9 ranged_attack_power points (0.87 DPS) | yes | Erudite's Amulet (277204, -0.28 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.7 ranged_attack_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, -0.66 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.9 ranged_attack_power points (0.87 DPS) | yes | Hide of Lupos (3018, -0.29 DPS) [world]; Bristlebark Cape (14571, -0.29 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.36 DPS, sim-verified) [dungeon] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 23.6 ranged_attack_power points (1.59 DPS) | yes | Brawler's Leather Armor (252490, -0.34 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.58 DPS) [crafted]; Dark Leather Tunic (2317, -0.72 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.7 ranged_attack_power points (0.72 DPS) | yes | Wolf Bracers (4794, -0.14 DPS) [vendor]; Bravo's Armbands (270015, -0.14 DPS) [quest]; Ratchet Wristwraps (274742, -0.29 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (82.1 DPS) | yes | Forest Leather Gloves (3058, -0.29 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.29 DPS) [crafted]; Serpent Gloves (5970, -0.84 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.21 DPS) | yes | Deviate Scale Belt (6468, -0.49 DPS) [crafted]; Dusty Belt (279897, -0.59 DPS, sim-verified) [quest]; Dark Leather Belt (4249, -0.63 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.3 ranged_attack_power points (1.30 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.14 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.2 ranged_attack_power points (1.16 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.29 DPS) [dungeon]; Agile Boots (4788, -0.43 DPS) [vendor] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.9 ranged_attack_power points (0.87 DPS) | yes | Lavishly Jeweled Ring (1156, -0.58 DPS) [dungeon]; The 1 Ring (8350, -0.72 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.6 ranged_attack_power points (0.58 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.43 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 23.2 ranged_attack_power points (1.56 DPS) | yes | Impaling Harpoon (5200, -0.26 DPS) [dungeon]; Scythe Axe (5749, -0.55 DPS) [world]; Lupine Axe (1220, -0.69 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (12.02 DPS) | yes | Lil Timmy's Peashooter (13136, -1.11 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.59 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.85 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 108.0. Weights run: 2.2s. Verify run: 2.0s. 366 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.162 ± 0.007, crit=0.604 ± 0.016 per rating point (14 rating = 1%, 8.456 per %), hit=0.968 ± 0.044 per rating point (10 rating = 1%, 9.679 per %), melee_haste=8.765 ± 0.713

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 ranged_attack_power points (1.44 DPS) | yes | Tribal Worg Helm (6204, -0.29 DPS) [world]; Brawler's Leather Hood (252504, -0.29 DPS) [crafted]; Humbert's Helm (4724, -0.43 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 17.3 ranged_attack_power points (1.15 DPS) | yes | Ghostshard Talisman (7731, -0.22 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.57 DPS) [world_drop]; Erudite's Amulet (277204, -0.57 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 ranged_attack_power points (1.58 DPS) | yes | Mantle of Thieves (2264, -0.14 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.57 DPS) [crafted]; Insignia Mantle (4721, -0.57 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.3 ranged_attack_power points (1.15 DPS) | yes | Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.29 DPS) [world]; Fenrus' Hide (6340, -0.29 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.3 ranged_attack_power points (2.01 DPS) | yes | Tunic of Westfall (2041, -0.45 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.86 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.86 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 13.0 ranged_attack_power points (0.86 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.14 DPS) [world_drop] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 17.3 ranged_attack_power points (1.15 DPS) | yes | Heavy Earthen Gloves (7359, -0.09 DPS) [crafted]; Serpent Gloves (5970, -0.29 DPS) [dungeon]; Insignia Gloves (6408, -0.29 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 ranged_attack_power points (1.59 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.30 DPS) [crafted]; Stalker's Leather Belt (252521, -0.30 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 30.3 ranged_attack_power points (2.01 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.14 DPS) [crafted]; Ferine Leggings (6690, -0.28 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.15 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.14 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.5 ranged_attack_power points (1.29 DPS) | yes | Ring of Precision (1491, -0.43 DPS) [dungeon]; Protector's Band (19517, -0.43 DPS) [rep]; Signet of the Zhevra (285330, -0.43 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.1 ranged_attack_power points (1.01 DPS) | yes | Protector's Band (19517, -0.14 DPS) [rep]; Signet of the Zhevra (285330, -0.14 DPS) [world]; Ring of Precision (1491, -0.34 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 ranged_attack_power points (1.12 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.0 ranged_attack_power points (0.86 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.38 DPS) | yes | Glass Shooter (9456, -0.70 DPS) [dungeon]; Silver Star (3463, -1.09 DPS, sim-verified) [quest]; Ironweaver (13137, -1.28 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Highlander's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 145.6. Weights run: 2.4s. Verify run: 2.5s. 597 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.171 ± 0.008, crit=0.645 ± 0.017 per rating point (14 rating = 1%, 9.028 per %), hit=1.026 ± 0.074 per rating point (10 rating = 1%, 10.260 per %), melee_haste=10.582 ± 1.100

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 28.2 ranged_attack_power points (1.83 DPS) | yes | Guard's Chain Helm (250499, -0.14 DPS) [crafted]; Skullsplitter Helm (1624, -0.28 DPS) [world]; Nightscape Headband (8176, -0.46 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 23.9 ranged_attack_power points (1.55 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.25 DPS) [quest]; Ghostshard Talisman (7731, -0.64 DPS) [dungeon]; Erudite's Amulet (277204, -0.99 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 35.9 ranged_attack_power points (2.33 DPS) | yes | Forest Tracker Epaulets (2278, -0.78 DPS) [world_drop]; Nightscape Shoulders (8192, -0.79 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.92 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 21.7 ranged_attack_power points (1.41 DPS) | yes | Imperial Cloak (6432, -0.28 DPS) [dungeon]; Parachute Cloak (10518, -0.28 DPS) [crafted]; Tigerstrike Mantle (13108, -0.28 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 41.3 ranged_attack_power points (2.68 DPS) | yes | Wolffear Harness (13110, -0.28 DPS) [world_drop]; Nightscape Tunic (8175, -0.56 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.56 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.30 DPS) | yes | Imperial Leather Bracers (4061, -0.17 DPS) [dungeon]; Dusky Bracers (7378, -0.17 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.31 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.08 DPS) | yes | Gloves of Holy Might (867, -0.19 DPS) [world_drop]; Dragonscale Gauntlets (8347, -0.65 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.67 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.95 DPS) | yes | Highlander's Chain Girdle (20090, -0.39 DPS) [rep]; Scorpashi Sash (14652, -0.54 DPS) [world_drop]; Blackforge Girdle (6425, -0.68 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.6 ranged_attack_power points (2.96 DPS) | yes | Triprunner Dungarees (9624, -0.42 DPS) [quest]; Petrolspill Leggings (9509, -0.99 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.99 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.2 ranged_attack_power points (1.83 DPS) | yes | Imperial Leather Boots (6431, -0.28 DPS) [dungeon]; Dusky Boots (7390, -0.28 DPS) [crafted]; Worn Running Boots (9398, -0.28 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 21.7 ranged_attack_power points (1.41 DPS) | yes | Assault Band (13095, -0.11 DPS) [world_drop]; Falcon's Hook (7552, -0.14 DPS) [dungeon]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 ranged_attack_power points (1.30 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.03 DPS) [dungeon]; Ironspine's Eye (7686, -0.03 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (145.6 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.44 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 19.5 ranged_attack_power points (1.27 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Satyr's Rod (15962, -1.13 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (22.75 DPS) | yes | Shadowforge Bushmaster (9422, -2.04 DPS) [dungeon]; Swiftwind (13038, -2.42 DPS) [world_drop]; The Silencer (13138, -3.12 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Mark of Kern; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 180.1. Weights run: 2.5s. Verify run: 2.8s. 754 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.192 ± 0.009, crit=0.705 ± 0.018 per rating point (14 rating = 1%, 9.872 per %), hit=1.019 ± 0.073 per rating point (10 rating = 1%, 10.185 per %), melee_haste=8.585 ± 1.178

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 44.9 ranged_attack_power points (3.03 DPS) | yes | Bloomsprout Headpiece (17767, -0.60 DPS) [dungeon]; Lordrec Helmet (10741, -0.67 DPS) [quest]; Helm of Fire (8348, -1.71 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.5 ranged_attack_power points (1.92 DPS) | yes | Sentinel's Medallion (19539, -0.15 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.57 DPS) [quest] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 38.4 ranged_attack_power points (2.59 DPS) | yes | Phytoskin Spaulders (17749, -0.22 DPS) [dungeon]; Sunburn Spaulders (274751, -0.56 DPS, sim-verified) [vendor]; Khan's Mantle (14787, -0.81 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 32.9 ranged_attack_power points (2.22 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.15 DPS) [dungeon]; Blackmetal Cape (9512, -0.44 DPS) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 54.8 ranged_attack_power points (3.69 DPS) | yes | Blazewind Breastplate (11193, -0.30 DPS) [quest]; Knight's Chain Armor (220828, -0.81 DPS) [vendor]; Quillward Harness (10583, -0.89 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 32.9 ranged_attack_power points (2.22 DPS) | yes | Bracers of the Stone Princess (17714, -0.33 DPS) [dungeon]; Arena Bands (18711, -0.33 DPS) [world]; Bloodlust Bracelets (14807, -0.59 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 48.2 ranged_attack_power points (3.25 DPS) | yes | Gauntlets of Divinity (7724, -1.09 DPS) [dungeon]; Sergeant Major's Chain Gauntlets (220829, -1.18 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.18 DPS) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 45.9 ranged_attack_power points (3.09 DPS) | yes | Sagebrush Girdle (17778, -0.88 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.02 DPS) [crafted]; Stalker's Mail Belt (252588, -1.02 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.0 ranged_attack_power points (3.10 DPS) | yes | Infernal Trickster Leggings (17754, +0.00 DPS, sim-verified) [dungeon]; Keeper's Woolies (14668, -0.30 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.37 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 43.8 ranged_attack_power points (2.95 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.30 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.44 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.7 ranged_attack_power points (2.07 DPS) | yes | Ring of the Underwood (2951, -0.59 DPS) [world_drop]; Mark of Kern (2262, -0.72 DPS) [dungeon]; Assault Band (13095, -0.72 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 30.2 ranged_attack_power points (2.03 DPS) | yes | Ring of the Underwood (2951, -0.67 DPS, sim-verified) [world_drop]; Mark of Kern (2262, -0.69 DPS) [dungeon]; Assault Band (13095, -0.69 DPS) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (180.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (180.1 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (180.1 DPS) | yes | Steel Spear (250605, -0.80 DPS) [crafted]; Manslayer (10570, -0.84 DPS) [dungeon]; Hanzo Sword (8190, -3.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (180.1 DPS) | yes | Hurricane (2824, -1.64 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -2.06 DPS) [world_drop]; Dark Iron Rifle (16004, -2.36 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 243.7. Weights run: 2.4s. Verify run: 9.8s. 1670 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.392 ± 0.016, crit=1.489 ± 0.037 per rating point (14 rating = 1%, 20.840 per %), hit=2.860 ± 0.131 per rating point (10 rating = 1%, 28.602 per %), melee_haste=17.222 ± 1.918

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (243.7 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -8.03 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.5 ranged_attack_power points (4.32 DPS) | yes | Beads of Ogre Might (22150, -0.88 DPS, sim-verified) [quest]; Mark of Fordring (15411, -1.18 DPS) [quest]; Amulet of the Darkmoon (19491, -1.27 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 85.4 ranged_attack_power points (5.72 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 56.6 ranged_attack_power points (3.79 DPS) | yes | Cape of the Black Baron (13340, -0.05 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.71 DPS) [rep]; Shadow Prowler's Cloak (22269, -1.07 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (243.7 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.40 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.40 DPS) [vendor]; Tunic of Undead Slaying (23089, -10.56 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (243.7 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Bracers of the Eclipse (18375, +0.00 DPS) [dungeon]; Beaststalker's Bindings (16681, -8.04 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (243.7 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (243.7 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -6.62 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 154.0 ranged_attack_power points (10.31 DPS) | yes | Sentinel's Leather Pants (237818, -3.20 DPS) [vendor]; Marshal's Chain Legguards (231577, -3.47 DPS) [pvp]; Plaguehound Leggings (18736, -3.59 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (243.7 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -6.79 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (243.7 DPS) | yes | Tarnished Elven Ring (18500, -0.48 DPS) [dungeon]; Cutthroat's Signet (272408, -0.64 DPS) [vendor]; Naglering (11669, -5.56 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (243.7 DPS) | yes | Tarnished Elven Ring (18500, -0.06 DPS) [dungeon]; Cutthroat's Signet (272408, -0.22 DPS) [vendor]; Naglering (11669, -4.92 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (243.7 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (243.7 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Smolderweb's Eye (13213, -2.00 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (243.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -7.77 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (243.7 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.98 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 668.8. Weights run: 2.5s. Verify run: 9.7s. 1670 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.451 ± 0.017, crit=1.609 ± 0.037 per rating point (14 rating = 1%, 22.531 per %), hit=3.234 ± 0.184 per rating point (10 rating = 1%, 32.335 per %), melee_haste=15.624 ± 2.490

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (668.8 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -9.76 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 69.1 ranged_attack_power points (9.68 DPS) | yes | Beads of Ogre Might (22150, -1.90 DPS, sim-verified) [quest]; Mark of Fordring (15411, -2.88 DPS) [quest]; Amulet of the Darkmoon (19491, -3.16 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 88.7 ranged_attack_power points (12.43 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 ranged_attack_power points (8.45 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -1.97 DPS) [rep]; Shadow Prowler's Cloak (22269, -2.62 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (668.8 DPS) | yes | Field Marshal's Chain Breastplate (16466, -5.56 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -5.56 DPS) [vendor]; Tunic of Undead Slaying (23089, -18.21 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (668.8 DPS) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -12.00 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 86.7 ranged_attack_power points (12.15 DPS) | yes | Gauntlets of Accuracy (18349, +0.00 DPS, sim-verified) [dungeon]; Marshal's Chain Vices (231578, -1.78 DPS) [vendor]; Marshal's Chain Grips (231560, -2.48 DPS) [pvp] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (668.8 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Marksman's Girdle (22232, -7.78 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 163.2 ranged_attack_power points (22.86 DPS) | yes | Sentinel's Leather Pants (237818, -7.28 DPS) [vendor]; Plaguehound Leggings (18736, -8.03 DPS) [dungeon]; Marshal's Chain Legguards (231577, -8.03 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (668.8 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -10.63 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (668.8 DPS) | yes | Tarnished Elven Ring (18500, -1.03 DPS) [dungeon]; Cutthroat's Signet (272408, -1.37 DPS) [vendor]; Naglering (11669, -7.33 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (668.8 DPS) | yes | Tarnished Elven Ring (18500, -0.25 DPS) [dungeon]; Cutthroat's Signet (272408, -0.59 DPS) [vendor]; Naglering (11669, -6.17 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (668.8 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS, sim-verified) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (668.8 DPS) | yes | Devilsaur Eye (19991, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.24 DPS) [crafted]; Counterattack Lodestone (18537, -3.23 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (668.8 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -11.79 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (668.8 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -24.25 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Second Wind; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 82.9. Weights run: 2.1s. Verify run: 3.3s. 209 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.148 ± 0.007, crit=0.502 ± 0.014 per rating point (14 rating = 1%, 7.027 per %), hit=0.867 ± 0.033 per rating point (10 rating = 1%, 8.669 per %), melee_haste=6.773 ± 0.552

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.2 ranged_attack_power points (1.16 DPS) | yes | Red Winter Hat (21524, -1.18 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.9 ranged_attack_power points (0.87 DPS) | yes | Erudite's Amulet (277204, -0.29 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.7 ranged_attack_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, -0.47 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.9 ranged_attack_power points (0.87 DPS) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Hide of Lupos (3018, -0.29 DPS) [world]; Bristlebark Cape (14571, -0.29 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 15.0 ranged_attack_power points (1.01 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.14 DPS) [crafted]; Prospector's Chestpiece (14562, -0.14 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.7 ranged_attack_power points (0.72 DPS) | yes | Wolf Bracers (4794, -0.14 DPS) [vendor]; Bristlebark Bindings (14569, -0.29 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.29 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (82.9 DPS) | yes | Forest Leather Gloves (3058, -0.29 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.29 DPS) [crafted]; Serpent Gloves (5970, -0.59 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.21 DPS) | yes | Deviate Scale Belt (6468, -0.52 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.63 DPS) [world]; Dark Leather Belt (4249, -0.63 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.3 ranged_attack_power points (1.30 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.14 DPS) [world]; Brawler's Leather Pants (252500, -0.85 DPS, sim-verified) [crafted] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (82.9 DPS) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Agile Boots (4788, -0.14 DPS) [vendor]; Feet of the Lynx (1121, -0.40 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.9 ranged_attack_power points (0.87 DPS) | yes | Bounty Hunter's Ring (5351, -0.43 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.58 DPS) [dungeon]; The 1 Ring (8350, -0.72 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.6 ranged_attack_power points (0.58 DPS) | yes | Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; The 1 Ring (8350, -0.43 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 23.2 ranged_attack_power points (1.56 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.55 DPS) [world]; Crescent Staff (6505, -0.55 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (12.02 DPS) | yes | Outrider's Bow (20437, -0.72 DPS) [pvp]; Lil Timmy's Peashooter (13136, -1.18 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.59 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 109.1. Weights run: 2.2s. Verify run: 2.0s. 352 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.162 ± 0.007, crit=0.604 ± 0.016 per rating point (14 rating = 1%, 8.456 per %), hit=0.968 ± 0.044 per rating point (10 rating = 1%, 9.679 per %), melee_haste=8.765 ± 0.713

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 ranged_attack_power points (1.44 DPS) | yes | Tribal Worg Helm (6204, -0.29 DPS) [world]; Brawler's Leather Hood (252504, -0.29 DPS) [crafted]; Humbert's Helm (4724, -0.43 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 17.3 ranged_attack_power points (1.15 DPS) | yes | Ghostshard Talisman (7731, -0.22 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.57 DPS) [world_drop]; Erudite's Amulet (277204, -0.57 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 ranged_attack_power points (1.58 DPS) | yes | Mantle of Thieves (2264, -0.14 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.57 DPS) [crafted]; Insignia Mantle (4721, -0.57 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.3 ranged_attack_power points (1.15 DPS) | yes | Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.29 DPS) [world]; Swiftrunner Cape (6745, -0.29 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.3 ranged_attack_power points (2.01 DPS) | yes | Panther Armor (6670, -0.73 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.86 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.86 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 13.0 ranged_attack_power points (0.86 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Loamflake Bracers (15462, -0.14 DPS) [quest] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 17.3 ranged_attack_power points (1.15 DPS) | yes | Heavy Earthen Gloves (7359, -0.09 DPS) [crafted]; Braced Handguards (6784, -0.14 DPS) [quest]; Insignia Gloves (6408, -0.29 DPS) [world_drop] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 ranged_attack_power points (1.59 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Deftkin Belt (16659, -0.22 DPS) [quest]; Skulker's Leather Belt (252520, -0.30 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 30.3 ranged_attack_power points (2.01 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.14 DPS) [crafted]; Ferine Leggings (6690, -0.28 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.15 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.5 ranged_attack_power points (1.29 DPS) | yes | Ring of Precision (1491, -0.43 DPS) [dungeon]; Legionnaire's Band (19513, -0.43 DPS) [rep]; Signet of the Zhevra (285330, -0.43 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.1 ranged_attack_power points (1.01 DPS) | yes | Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep]; Signet of the Zhevra (285330, -0.14 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 ranged_attack_power points (1.12 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Outlaw Sabre (16886) | Baron Aquanis [quest] | 15.0 ranged_attack_power points (1.00 DPS) | yes | Vendetta (776, +0.00 DPS) [dungeon]; Satyr's Rod (15962, -0.85 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.38 DPS) | yes | Silver Star (3463, -0.43 DPS, sim-verified) [quest]; Glass Shooter (9456, -0.70 DPS) [dungeon]; Ironweaver (13137, -1.28 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Defiler's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Outlaw Sabre; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 146.7. Weights run: 2.4s. Verify run: 2.4s. 563 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.171 ± 0.008, crit=0.645 ± 0.017 per rating point (14 rating = 1%, 9.028 per %), hit=1.026 ± 0.074 per rating point (10 rating = 1%, 10.260 per %), melee_haste=10.582 ± 1.100

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 28.2 ranged_attack_power points (1.83 DPS) | yes | Nightscape Headband (8176, -0.14 DPS) [crafted]; Guard's Chain Helm (250499, -0.14 DPS) [crafted]; Skullsplitter Helm (1624, -0.28 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 23.9 ranged_attack_power points (1.55 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.25 DPS) [quest]; Ghostshard Talisman (7731, -0.64 DPS) [dungeon]; Erudite's Amulet (277204, -0.99 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 35.9 ranged_attack_power points (2.33 DPS) | yes | Forest Tracker Epaulets (2278, -0.78 DPS) [world_drop]; Nightscape Shoulders (8192, -0.80 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.92 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 21.7 ranged_attack_power points (1.41 DPS) | yes | Imperial Cloak (6432, -0.28 DPS) [dungeon]; Parachute Cloak (10518, -0.28 DPS) [crafted]; Tigerstrike Mantle (13108, -0.28 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 41.3 ranged_attack_power points (2.68 DPS) | yes | Wolffear Harness (13110, -0.28 DPS) [world_drop]; Nightscape Tunic (8175, -0.56 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.56 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.30 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Dusky Bracers (7378, -0.17 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.08 DPS) | yes | Gloves of Holy Might (867, -0.19 DPS) [world_drop]; Dragonscale Gauntlets (8347, -0.65 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.67 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.95 DPS) | yes | Defiler's Chain Girdle (20152, -0.39 DPS) [rep]; Scorpashi Sash (14652, -0.54 DPS) [world_drop]; Deftkin Belt (16659, -0.61 DPS) [quest] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.6 ranged_attack_power points (2.96 DPS) | yes | Triprunner Dungarees (9624, -0.56 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.99 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.99 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.2 ranged_attack_power points (1.83 DPS) | yes | Imperial Leather Boots (6431, -0.28 DPS) [dungeon]; Dusky Boots (7390, -0.28 DPS) [crafted]; Worn Running Boots (9398, -0.28 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 21.7 ranged_attack_power points (1.41 DPS) | yes | Assault Band (13095, -0.11 DPS) [world_drop]; Falcon's Hook (7552, -0.14 DPS) [dungeon]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 ranged_attack_power points (1.30 DPS) | yes | Assault Band (13095, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.03 DPS) [dungeon]; Ironspine's Eye (7686, -0.03 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (146.7 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.46 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 19.5 ranged_attack_power points (1.27 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Satyr's Rod (15962, -1.13 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (22.75 DPS) | yes | Outrider's Bow (19560, -1.66 DPS) [pvp]; Shadowforge Bushmaster (9422, -2.04 DPS) [dungeon]; The Silencer (13138, -2.85 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Mark of Kern; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 184.9. Weights run: 2.5s. Verify run: 2.8s. 713 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.192 ± 0.009, crit=0.705 ± 0.018 per rating point (14 rating = 1%, 9.872 per %), hit=1.019 ± 0.073 per rating point (10 rating = 1%, 10.185 per %), melee_haste=8.585 ± 1.178

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 37.3 ranged_attack_power points (2.51 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.08 DPS) [dungeon]; Sprightring Helm (17776, -0.30 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.5 ranged_attack_power points (1.92 DPS) | yes | Scout's Medallion (19535, -0.15 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.57 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.1 ranged_attack_power points (2.43 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -0.07 DPS) [dungeon]; Khan's Mantle (14787, -0.66 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 32.9 ranged_attack_power points (2.22 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS, sim-verified) [dungeon]; Blackveil Cape (11626, -0.15 DPS) [dungeon]; Blackmetal Cape (9512, -0.44 DPS) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 54.8 ranged_attack_power points (3.69 DPS) | yes | Blazewind Breastplate (11193, -0.30 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.81 DPS) [vendor]; Quillward Harness (10583, -0.89 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 32.9 ranged_attack_power points (2.22 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 48.2 ranged_attack_power points (3.25 DPS) | yes | Gauntlets of Divinity (7724, -1.09 DPS) [dungeon]; First Sergeant's Chain Gauntlets (220830, -1.16 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.18 DPS) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 45.9 ranged_attack_power points (3.09 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.02 DPS) [crafted]; Stalker's Mail Belt (252588, -1.02 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.0 ranged_attack_power points (3.10 DPS) | yes | Infernal Trickster Leggings (17754, +0.00 DPS, sim-verified) [dungeon]; Keeper's Woolies (14668, -0.30 DPS) [world_drop]; Stone Guard's Chain Legplates (220833, -0.37 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 43.8 ranged_attack_power points (2.95 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.30 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.44 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.7 ranged_attack_power points (2.07 DPS) | yes | White Bone Band (11862, -0.45 DPS) [quest]; Ring of the Underwood (2951, -0.59 DPS) [world_drop]; Mark of Kern (2262, -0.72 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 30.2 ranged_attack_power points (2.03 DPS) | yes | Ring of the Underwood (2951, -0.56 DPS) [world_drop]; Mark of Kern (2262, -0.69 DPS) [dungeon]; White Bone Band (11862, -0.89 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (184.9 DPS) | yes | Frozen Heart of the Mountain (249469, -5.52 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (184.9 DPS) | yes | Frozen Heart of the Mountain (249469, -1.97 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (184.9 DPS) | yes | Steel Spear (250605, -0.80 DPS) [crafted]; Manslayer (10570, -0.84 DPS) [dungeon]; Hanzo Sword (8190, -3.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (184.9 DPS) | yes | Hurricane (2824, -1.64 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -2.06 DPS) [world_drop]; Dark Iron Rifle (16004, -2.32 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 250.5. Weights run: 2.4s. Verify run: 9.3s. 1650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.392 ± 0.016, crit=1.489 ± 0.037 per rating point (14 rating = 1%, 20.840 per %), hit=2.860 ± 0.131 per rating point (10 rating = 1%, 28.602 per %), melee_haste=17.222 ± 1.918

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (250.5 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -8.96 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.5 ranged_attack_power points (4.32 DPS) | yes | Beads of Ogre Might (22150, -0.94 DPS, sim-verified) [quest]; Mark of Fordring (15411, -1.18 DPS) [quest]; Amulet of the Darkmoon (19491, -1.27 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 85.4 ranged_attack_power points (5.72 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 56.6 ranged_attack_power points (3.79 DPS) | yes | Cape of the Black Baron (13340, -0.05 DPS) [dungeon]; Deathguard's Cloak (20068, -0.71 DPS) [rep]; Shadow Prowler's Cloak (22269, -1.07 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (250.5 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.40 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.40 DPS) [vendor]; Tunic of Undead Slaying (23089, -11.20 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (250.5 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Bracers of the Eclipse (18375, +0.00 DPS) [dungeon]; Beaststalker's Bindings (16681, -8.07 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (250.5 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Chain Gloves (16571, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (250.5 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -7.12 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 154.0 ranged_attack_power points (10.31 DPS) | yes | Outrider's Chain Leggings (22673, -1.51 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -3.20 DPS) [vendor]; General's Chain Legguards (231574, -3.47 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (250.5 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -7.65 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (250.5 DPS) | yes | Tarnished Elven Ring (18500, -0.48 DPS) [dungeon]; Cutthroat's Signet (272408, -0.64 DPS) [vendor]; Naglering (11669, -6.04 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (250.5 DPS) | yes | Tarnished Elven Ring (18500, -0.06 DPS) [dungeon]; Cutthroat's Signet (272408, -0.22 DPS) [vendor]; Naglering (11669, -5.29 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (250.5 DPS) | yes | Blackhand's Breadth (13965, -4.18 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.24 DPS) [crafted]; Counterattack Lodestone (18537, -5.49 DPS) [dungeon] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (250.5 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Second Wind (11819, -2.16 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (250.5 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -8.25 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (250.5 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -11.60 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Burst of Knowledge; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 683.1. Weights run: 2.5s. Verify run: 9.8s. 1650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.451 ± 0.017, crit=1.609 ± 0.037 per rating point (14 rating = 1%, 22.531 per %), hit=3.234 ± 0.184 per rating point (10 rating = 1%, 32.335 per %), melee_haste=15.624 ± 2.490

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (683.1 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -8.07 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 69.1 ranged_attack_power points (9.68 DPS) | yes | Beads of Ogre Might (22150, -1.99 DPS, sim-verified) [quest]; Mark of Fordring (15411, -2.88 DPS) [quest]; Amulet of the Darkmoon (19491, -3.16 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 88.7 ranged_attack_power points (12.43 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 ranged_attack_power points (8.45 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -1.97 DPS) [rep]; Shadow Prowler's Cloak (22269, -2.62 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (683.1 DPS) | yes | Warlord's Chain Chestpiece (16565, -5.56 DPS) [vendor]; Warlord's Chain Hauberk (231573, -5.56 DPS) [vendor]; Tunic of Undead Slaying (23089, -18.29 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (683.1 DPS) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -8.59 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | 86.7 ranged_attack_power points (12.15 DPS) | yes | Gauntlets of Accuracy (18349, +0.00 DPS, sim-verified) [dungeon]; General's Chain Gloves (16571, -1.78 DPS) [vendor]; General's Chain Vices (231575, -1.78 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (683.1 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Marksman's Girdle (22232, -6.90 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 163.2 ranged_attack_power points (22.86 DPS) | yes | Outrider's Chain Leggings (22673, -3.29 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -7.28 DPS) [vendor]; Plaguehound Leggings (18736, -8.03 DPS) [dungeon] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (683.1 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -9.49 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (683.1 DPS) | yes | Tarnished Elven Ring (18500, -1.03 DPS) [dungeon]; Cutthroat's Signet (272408, -1.37 DPS) [vendor]; Naglering (11669, -7.35 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (683.1 DPS) | yes | Tarnished Elven Ring (18500, -0.25 DPS) [dungeon]; Cutthroat's Signet (272408, -0.59 DPS) [vendor]; Naglering (11669, -6.02 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (683.1 DPS) | yes | Blackhand's Breadth (13965, -8.63 DPS) [quest]; Frozen Heart of the Mountain (249469, -10.86 DPS) [crafted]; Counterattack Lodestone (18537, -11.86 DPS) [dungeon] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (683.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Second Wind (11819, -1.74 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (683.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -11.74 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (683.1 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -23.89 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Burst of Knowledge; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

