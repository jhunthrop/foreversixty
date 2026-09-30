# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.1. Weights run: 1.4s. Verify run: 1.6s. 217 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.185 ± 0.045, crit=0.477 ± 0.019 per rating point (14 rating = 1%, 6.678 per %), hit=0.165 ± 0.007 per rating point (10 rating = 1%, 1.648 per %), melee_haste=3.462 ± 0.795

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.5 ranged_attack_power points (1.04 DPS) | yes | Lucky Fishing Hat (19972, -1.02 DPS, sim-verified) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.9 ranged_attack_power points (0.65 DPS) | yes | Forest Leather Mantle (4709, -0.77 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Cape of the Brotherhood (5193, +0.00 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [world_drop]; Hide of Lupos (3018, -0.26 DPS) [world] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 24.0 ranged_attack_power points (1.44 DPS) | yes | Brawler's Leather Armor (252490, -0.35 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.52 DPS) [crafted]; Prospector's Chestpiece (14562, -0.65 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.9 ranged_attack_power points (0.65 DPS) | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Bravo's Armbands (270015, -0.13 DPS) [quest]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Gloves of the Fang (10413, -0.08 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.26 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.26 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.07 DPS) | yes | Dusty Belt (279897, -0.42 DPS) [quest]; Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Dark Leather Belt (4249, -0.55 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 19.7 ranged_attack_power points (1.17 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.13 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.5 ranged_attack_power points (1.04 DPS) | yes | Footpads of the Fang (10411, -0.26 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.26 DPS) [dungeon]; Bristlebark Boots (14568, -0.39 DPS) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Lavishly Jeweled Ring (1156, -0.52 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.7 ranged_attack_power points (0.52 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.39 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 23.5 ranged_attack_power points (1.40 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.49 DPS) [world]; Pearl-encrusted Spear (1406, -0.62 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (10.65 DPS) | yes | Lil Timmy's Peashooter (13136, -0.69 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.30 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.53 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 30 (dwarf, 0000000000000000-35305500000000000-000000000000000000)

Set DPS (verified): 89.3. Weights run: 1.5s. Verify run: 1.7s. 363 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.180 ± 0.048, crit=0.587 ± 0.023 per rating point (14 rating = 1%, 8.214 per %), hit=0.184 ± 0.007 per rating point (10 rating = 1%, 1.844 per %), melee_haste=7.178 ± 0.865

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.8 ranged_attack_power points (1.29 DPS) | yes | Tribal Worg Helm (6204, +0.00 DPS, sim-verified) [world]; Brawler's Leather Hood (252504, -0.26 DPS) [crafted]; Humbert's Helm (4724, -0.39 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Ghostshard Talisman (7731, -0.20 DPS, sim-verified) [dungeon]; Sentinel's Medallion (20444, -0.26 DPS) [rep]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.0 ranged_attack_power points (1.42 DPS) | yes | Mantle of Thieves (2264, +0.00 DPS, sim-verified) [dungeon]; Insignia Mantle (4721, -0.51 DPS) [world_drop]; Cloudy Gustwoven Spaulders (277043, -0.51 DPS) [crafted] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Hawkeye's Cloak (14593, -0.14 DPS, sim-verified) [world_drop]; Cloak of Night (4447, -0.26 DPS) [world]; Fenrus' Hide (6340, -0.26 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.5 ranged_attack_power points (1.80 DPS) | yes | Tunic of Westfall (2041, -0.39 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.77 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.77 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 13.1 ranged_attack_power points (0.77 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS, sim-verified) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.13 DPS) [world] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Heavy Earthen Gloves (7359, -0.08 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -0.26 DPS) [world_drop]; Ebon Vise (7690, -0.26 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 ranged_attack_power points (1.42 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Skulker's Leather Belt (252520, -0.26 DPS) [crafted]; Stalker's Leather Belt (252521, -0.26 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Ambusher [dungeon] | 30.5 ranged_attack_power points (1.80 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.13 DPS) [crafted]; Ferine Leggings (6690, -0.27 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS, sim-verified) [vendor]; Lancer Boots (6752, -0.13 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 ranged_attack_power points (1.16 DPS) | yes | Ring of Precision (1491, -0.39 DPS) [dungeon]; Protector's Band (19517, -0.39 DPS) [rep]; Signet of the Zhevra (285330, -0.39 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.3 ranged_attack_power points (0.90 DPS) | yes | Ring of Precision (1491, -0.13 DPS) [dungeon]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Protector's Band (19517, -0.31 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.9 ranged_attack_power points (1.00 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.1 ranged_attack_power points (0.77 DPS) | yes | Prison Shank (2941, +0.00 DPS, sim-verified) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (15.44 DPS) | yes | Glass Shooter (9456, -0.62 DPS) [dungeon]; Ironweaver (13137, -1.13 DPS) [world_drop]; Silver Star (3463, -1.26 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Highlander's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 363, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (dwarf, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 110.4. Weights run: 1.5s. Verify run: 1.8s. 592 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.222 ± 0.063, crit=0.834 ± 0.034 per rating point (14 rating = 1%, 11.681 per %), hit=0.204 ± 0.008 per rating point (10 rating = 1%, 2.043 per %), melee_haste=not significant (3.437 ± 0.971)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 28.9 ranged_attack_power points (1.69 DPS) | yes | Guard's Chain Helm (250499, -0.13 DPS) [crafted]; Skullsplitter Helm (1624, -0.26 DPS) [world]; Nightscape Headband (8176, -0.29 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 24.4 ranged_attack_power points (1.43 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.36 DPS, sim-verified) [quest]; Sentinel's Medallion (19541, -0.39 DPS) [rep]; Ghostshard Talisman (7731, -0.61 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.4 ranged_attack_power points (2.14 DPS) | yes | Forest Tracker Epaulets (2278, -0.70 DPS) [world_drop]; Nightscape Shoulders (8192, -0.72 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.83 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 22.2 ranged_attack_power points (1.30 DPS) | yes | Imperial Cloak (6432, -0.26 DPS) [world_drop]; Tigerstrike Mantle (13108, -0.26 DPS) [world_drop]; Parachute Cloak (10518, -0.30 DPS, sim-verified) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 42.2 ranged_attack_power points (2.48 DPS) | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.52 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.52 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Imperial Leather Bracers (4061, -0.07 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.13 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.26 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (1.88 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS, sim-verified) [world_drop]; Dragonscale Gauntlets (8347, -0.41 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.57 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.76 DPS) | yes | Highlander's Leather Girdle (20117, -0.35 DPS) [rep]; Highlander's Chain Girdle (20090, -0.36 DPS, sim-verified) [rep]; Scorpashi Sash (14652, -0.46 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.7 ranged_attack_power points (2.74 DPS) | yes | Triprunner Dungarees (9624, -0.14 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.91 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.91 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.9 ranged_attack_power points (1.69 DPS) | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.26 DPS) [world_drop]; Stalker's Mail Boots (252562, -0.26 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 22.2 ranged_attack_power points (1.30 DPS) | yes | Ironspine's Eye (7686, -0.13 DPS) [dungeon]; Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Mark of Kern (2262, -0.00 DPS) [dungeon]; Assault Band (13095, -0.00 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (110.4 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.33 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS, sim-verified) [crafted]; Satyr's Rod (15962, -1.04 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (20.53 DPS) | yes | Shadowforge Bushmaster (9422, -1.84 DPS) [dungeon]; Swiftwind (13038, -2.17 DPS) [world_drop]; The Silencer (13138, -2.72 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 592, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5500000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 136.1. Weights run: 1.5s. Verify run: 1.9s. 755 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.307 ± 0.077, crit=0.964 ± 0.038 per rating point (14 rating = 1%, 13.494 per %), hit=0.240 ± 0.010 per rating point (10 rating = 1%, 2.404 per %), melee_haste=13.786 ± 1.194

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 50.4 ranged_attack_power points (2.95 DPS) | yes | Ebon Mask (19984, -0.79 DPS) [quest]; Lordrec Helmet (10741, -0.79 DPS) [quest]; Helm of Fire (8348, -2.02 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.0 ranged_attack_power points (1.75 DPS) | yes | Sentinel's Medallion (19539, -0.10 DPS, sim-verified) [rep]; Sentinel's Medallion (19540, -0.27 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.58 DPS) [quest] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 43.5 ranged_attack_power points (2.54 DPS) | yes | Phytoskin Spaulders (17749, -0.38 DPS) [dungeon]; Sunburn Spaulders (274751, -0.77 DPS, sim-verified) [vendor]; Khan's Mantle (14787, -0.92 DPS) [world_drop] |
| back | Blisterbane Wrap (12552) (or Dark Phantom Cape (13122)) | Blackrock Depths: Anvilrage Overseer [dungeon] | 34.6 ranged_attack_power points (2.02 DPS) | yes | Dark Phantom Cape (13122, +0.00 DPS, sim-verified) [world_drop]; Blackveil Cape (11626, -0.13 DPS) [dungeon]; Duskbat Drape (19982, -0.13 DPS) [quest] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 57.7 ranged_attack_power points (3.37 DPS) | yes | Blazewind Breastplate (11193, -0.30 DPS, sim-verified) [quest]; Knight's Chain Armor (220828, -0.56 DPS) [vendor]; Quillward Harness (10583, -0.81 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.6 ranged_attack_power points (2.02 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Bloodlust Bracelets (14807, -0.54 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.54 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 50.8 ranged_attack_power points (2.97 DPS) | yes | Gloves of Holy Might (867, -1.01 DPS) [world_drop]; Sergeant Major's Chain Gauntlets (220829, -1.04 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.08 DPS) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 47.6 ranged_attack_power points (2.78 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Highlander's Chain Girdle (20088, -0.82 DPS) [rep]; Highlander's Leather Girdle (20115, -0.82 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 48.5 ranged_attack_power points (2.83 DPS) | yes | Infernal Trickster Leggings (17754, +0.00 DPS, sim-verified) [dungeon]; Knight's Chain Legplates (220832, -0.16 DPS) [vendor]; Keeper's Woolies (14668, -0.27 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 46.1 ranged_attack_power points (2.70 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.27 DPS) [world_drop]; Whisperwalk Boots (20255, -0.27 DPS) [quest] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.3 ranged_attack_power points (1.89 DPS) | yes | Blackstone Ring (17713, -0.58 DPS) [dungeon]; Falcon's Hook (7552, -0.67 DPS) [world_drop]; Protector's Band (19516, -0.67 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.1 ranged_attack_power points (1.35 DPS) | yes | Blackstone Ring (17713, +0.00 DPS, sim-verified) [dungeon]; Falcon's Hook (7552, -0.13 DPS) [world_drop]; Protector's Band (19516, -0.13 DPS) [rep] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (136.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Fire Ruby (20036, -1.73 DPS, sim-verified) [quest] |
| trinket2 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (136.1 DPS) | yes | Fire Ruby (20036, +0.00 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (136.1 DPS) | yes | Steel Spear (250605, -0.76 DPS) [crafted]; Manslayer (10570, -0.88 DPS) [dungeon]; Hanzo Sword (8190, -3.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (136.1 DPS) | yes | Hurricane (2824, -1.42 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -1.70 DPS) [world_drop]; Dark Iron Rifle (16004, -1.84 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Blisterbane Wrap; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Devilsaur Eye; trinket2: Sanctified Orb; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 755, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 227.8. Weights run: 1.5s. Verify run: 1.9s. 1590 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.339 ± 0.094, crit=1.438 ± 0.055 per rating point (14 rating = 1%, 20.135 per %), hit=0.343 ± 0.014 per rating point (10 rating = 1%, 3.434 per %), melee_haste=10.525 ± 1.657

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Visor (239532) | Leonid Barthalomew the Revered [vendor] | 196.3 ranged_attack_power points (11.14 DPS) | yes | Field Marshal's Chain Helm (231580, -5.48 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, -5.70 DPS) [vendor]; Dawnstalker Headpiece (239540, -6.20 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 44.4 ranged_attack_power points (2.52 DPS) | yes | Medallion of the Dawn (22659, +0.00 DPS, sim-verified) [quest]; Pendant of Celerity (22340, -0.51 DPS) [dungeon]; Maelstrom's Tendril (19620, -0.53 DPS) [quest] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | 143.1 ranged_attack_power points (8.12 DPS) | yes | Dawnstalker Spaulders (239542, -2.72 DPS, sim-verified) [vendor]; Darkspear Pauldrons (272105, -3.39 DPS) [vendor]; Darkspear Epaulets (272106, -3.39 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 55.1 ranged_attack_power points (3.13 DPS) | yes | Cloak of the Honor Guard (20073, -0.83 DPS, sim-verified) [rep]; Shifting Cloak (18511, -0.87 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.87 DPS) [dungeon] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (227.8 DPS) | yes | Dawnstalker Tunic (239543, -1.70 DPS) [vendor]; Dawn Armor (252483, -5.38 DPS) [crafted]; Tunic of Undead Slaying (23089, -16.79 DPS, sim-verified) [world] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | sim-verified (227.8 DPS) | yes | Dawnstalker Wristguards (239544, -0.66 DPS) [vendor]; Bracers of the Eclipse (18375, -2.96 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.47 DPS, sim-verified) [world] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | 145.6 ranged_attack_power points (8.27 DPS) | yes | Dawnstalker Handguards (239539, -1.55 DPS, sim-verified) [vendor]; Marshal's Chain Vices (231578, -4.34 DPS) [vendor]; Marshal's Chain Grips (231560, -4.55 DPS) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 122.3 ranged_attack_power points (6.94 DPS) | yes | Dawnstalker Girdle (239538, -1.37 DPS, sim-verified) [vendor]; Molten Belt (19163, -3.23 DPS) [crafted]; Dense Timbermaw Belt (227807, -3.31 DPS) [vendor] |
| legs | Dawnstalker Leggings (239533) | Leonid Barthalomew the Revered [vendor] | 193.9 ranged_attack_power points (11.01 DPS) | yes | Sentinel's Chain Leggings (237819, -3.88 DPS) [vendor]; Sentinel's Chain Leggings (22748, -5.02 DPS) [rep]; Dawnstalker Legguards (239541, -6.38 DPS, sim-verified) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | 154.3 ranged_attack_power points (8.76 DPS) | yes | Dawnstalker Boots (239537, -3.68 DPS, sim-verified) [vendor]; Marshal's Chain Boots (16462, -5.11 DPS) [vendor]; Marshal's Chain Greaves (231579, -5.11 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (227.8 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.27 DPS) [quest]; Naglering (11669, -6.48 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (227.8 DPS) | yes | Cutthroat's Signet (272408, -0.19 DPS) [vendor]; Tarnished Elven Ring (18500, -0.24 DPS) [dungeon]; Naglering (11669, -5.24 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (227.8 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -1.51 DPS, sim-verified) [quest] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (227.8 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS, sim-verified) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (227.8 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -7.17 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (227.8 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [vendor]; Grand Marshal's Repeater (234586, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -8.86 DPS, sim-verified) [crafted] |

**New at 60:** head: Dawnstalker Visor; neck: Amulet of the Darkmoon; shoulder: Dawnstalker Pauldrons; back: Cape of the Black Baron; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Leggings; feet: Dawnstalker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1590, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.8. Weights run: 1.4s. Verify run: 1.6s. 213 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.185 ± 0.045, crit=0.477 ± 0.019 per rating point (14 rating = 1%, 6.678 per %), hit=0.165 ± 0.007 per rating point (10 rating = 1%, 1.648 per %), melee_haste=3.462 ± 0.795

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.5 ranged_attack_power points (1.04 DPS) | yes | Lucky Fishing Hat (19972, -1.02 DPS, sim-verified) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.9 ranged_attack_power points (0.65 DPS) | yes | Forest Leather Mantle (4709, -0.88 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Sentry Cloak (2059, -0.26 DPS) [world_drop]; Hide of Lupos (3018, -0.26 DPS) [world]; Cape of the Brotherhood (5193, -0.30 DPS, sim-verified) [dungeon] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | sim-verified (75.8 DPS) | yes | Dark Leather Tunic (2317, -0.13 DPS) [crafted]; Prospector's Chestpiece (14562, -0.13 DPS) [world_drop]; Brawler's Leather Armor (252490, -0.76 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.9 ranged_attack_power points (0.65 DPS) | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor]; Spare Part Bindings (279875, -0.26 DPS) [quest] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Gloves of the Fang (10413, -0.08 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.26 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.26 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.07 DPS) | yes | Deviate Scale Belt (6468, -0.45 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.55 DPS) [world_drop]; Dark Leather Belt (4249, -0.55 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 19.7 ranged_attack_power points (1.17 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.13 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.5 ranged_attack_power points (1.04 DPS) | yes | Footpads of the Fang (10411, -0.25 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.26 DPS) [dungeon]; Bristlebark Boots (14568, -0.39 DPS) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Bounty Hunter's Ring (5351, -0.39 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.52 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.7 ranged_attack_power points (0.52 DPS) | yes | Bounty Hunter's Ring (5351, -0.13 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; The 1 Ring (8350, -0.39 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 23.5 ranged_attack_power points (1.40 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.49 DPS) [world]; Crescent Staff (6505, -0.49 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (10.65 DPS) | yes | Lil Timmy's Peashooter (13136, -1.76 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.30 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.53 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 213, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-35305500000000000-000000000000000000)

Set DPS (verified): 91.4. Weights run: 1.5s. Verify run: 1.8s. 360 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.180 ± 0.048, crit=0.587 ± 0.023 per rating point (14 rating = 1%, 8.214 per %), hit=0.184 ± 0.007 per rating point (10 rating = 1%, 1.844 per %), melee_haste=7.178 ± 0.865

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tribal Worg Helm (6204) | Fenros [world] | sim-verified (91.4 DPS) | yes | Brawler's Leather Hood (252504, +0.00 DPS) [crafted]; Humbert's Helm (4724, -0.13 DPS) [world]; Brawler's Leather Helm (252512, -1.37 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Ghostshard Talisman (7731, -0.23 DPS, sim-verified) [dungeon]; Scout's Medallion (20442, -0.26 DPS) [rep]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.0 ranged_attack_power points (1.42 DPS) | yes | Mantle of Thieves (2264, +0.00 DPS, sim-verified) [dungeon]; Insignia Mantle (4721, -0.51 DPS) [world_drop]; Cloudy Gustwoven Spaulders (277043, -0.51 DPS) [crafted] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Hawkeye's Cloak (14593, -0.13 DPS, sim-verified) [world_drop]; Cloak of Night (4447, -0.26 DPS) [world]; Swiftrunner Cape (6745, -0.26 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.5 ranged_attack_power points (1.80 DPS) | yes | Panther Armor (6670, -0.68 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.77 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.77 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 13.1 ranged_attack_power points (0.77 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS, sim-verified) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.13 DPS) [world] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Heavy Earthen Gloves (7359, -0.11 DPS, sim-verified) [crafted]; Braced Handguards (6784, -0.13 DPS) [quest]; Ebon Vise (7690, -0.26 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 ranged_attack_power points (1.42 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Deftkin Belt (16659, -0.19 DPS) [quest]; Skulker's Leather Belt (252520, -0.26 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Ambusher [dungeon] | 30.5 ranged_attack_power points (1.80 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.13 DPS) [crafted]; Ferine Leggings (6690, -0.27 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS, sim-verified) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 ranged_attack_power points (1.16 DPS) | yes | Ring of Precision (1491, -0.39 DPS) [dungeon]; Legionnaire's Band (19513, -0.39 DPS) [rep]; Signet of the Zhevra (285330, -0.39 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.3 ranged_attack_power points (0.90 DPS) | yes | Ring of Precision (1491, -0.13 DPS) [dungeon]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Legionnaire's Band (19513, -0.44 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.9 ranged_attack_power points (1.00 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Outlaw Sabre (16886) | Baron Aquanis [quest] | 15.0 ranged_attack_power points (0.89 DPS) | yes | Vendetta (776, -0.11 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -0.76 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (15.44 DPS) | yes | Glass Shooter (9456, -0.62 DPS) [dungeon]; Ironweaver (13137, -1.13 DPS) [world_drop]; Silver Star (3463, -1.54 DPS, sim-verified) [quest] |

**New at 30:** head: Tribal Worg Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Defiler's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Outlaw Sabre; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 360, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 110.6. Weights run: 1.5s. Verify run: 1.7s. 580 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.222 ± 0.063, crit=0.834 ± 0.034 per rating point (14 rating = 1%, 11.681 per %), hit=0.204 ± 0.008 per rating point (10 rating = 1%, 2.043 per %), melee_haste=not significant (3.437 ± 0.971)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 28.9 ranged_attack_power points (1.69 DPS) | yes | Nightscape Headband (8176, +0.00 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -0.13 DPS) [crafted]; Skullsplitter Helm (1624, -0.26 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 24.4 ranged_attack_power points (1.43 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.28 DPS, sim-verified) [quest]; Scout's Medallion (19537, -0.39 DPS) [rep]; Ghostshard Talisman (7731, -0.61 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.4 ranged_attack_power points (2.14 DPS) | yes | Forest Tracker Epaulets (2278, -0.70 DPS) [world_drop]; Nightscape Shoulders (8192, -0.72 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.83 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 22.2 ranged_attack_power points (1.30 DPS) | yes | Imperial Cloak (6432, -0.26 DPS) [world_drop]; Tigerstrike Mantle (13108, -0.26 DPS) [world_drop]; Parachute Cloak (10518, -0.27 DPS, sim-verified) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 42.2 ranged_attack_power points (2.48 DPS) | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.52 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.52 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Imperial Leather Bracers (4061, -0.09 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.13 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.26 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (1.88 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS, sim-verified) [world_drop]; Dragonscale Gauntlets (8347, -0.41 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.57 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.76 DPS) | yes | Defiler's Leather Girdle (20191, -0.35 DPS) [rep]; Defiler's Chain Girdle (20152, -0.36 DPS, sim-verified) [rep]; Scorpashi Sash (14652, -0.46 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.7 ranged_attack_power points (2.74 DPS) | yes | Triprunner Dungarees (9624, -0.30 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.91 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.91 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.9 ranged_attack_power points (1.69 DPS) | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.26 DPS) [world_drop]; Stalker's Mail Boots (252562, -0.26 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 22.2 ranged_attack_power points (1.30 DPS) | yes | Ironspine's Eye (7686, -0.13 DPS) [dungeon]; Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Mark of Kern (2262, -0.00 DPS) [dungeon]; Assault Band (13095, -0.00 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (110.6 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.32 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS, sim-verified) [crafted]; Satyr's Rod (15962, -1.04 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (20.53 DPS) | yes | Shadowforge Bushmaster (9422, -1.84 DPS) [dungeon]; The Silencer (13138, -2.12 DPS, sim-verified) [world_drop]; Swiftwind (13038, -2.17 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 580, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5500000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 141.0. Weights run: 1.5s. Verify run: 1.8s. 744 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.307 ± 0.077, crit=0.964 ± 0.038 per rating point (14 rating = 1%, 13.494 per %), hit=0.240 ± 0.010 per rating point (10 rating = 1%, 2.404 per %), melee_haste=13.786 ± 1.194

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) | Lady Palanseer [vendor] | 50.4 ranged_attack_power points (2.95 DPS) | yes | Ebon Mask (19984, -0.79 DPS) [quest]; Bloomsprout Headpiece (17767, -0.84 DPS) [dungeon]; Helm of Fire (8348, -2.49 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.0 ranged_attack_power points (1.75 DPS) | yes | Scout's Medallion (19536, -0.27 DPS) [rep]; Scout's Medallion (19535, -0.49 DPS, sim-verified) [rep]; Woven Ivy Necklace (19159, -0.54 DPS) [quest] |
| shoulder | Blood Guard's Chain Epaulets (220824) | Lady Palanseer [vendor] | 43.5 ranged_attack_power points (2.54 DPS) | yes | Phytoskin Spaulders (17749, -0.38 DPS) [dungeon]; Sunburn Spaulders (274751, -0.76 DPS, sim-verified) [vendor]; Khan's Mantle (14787, -0.92 DPS) [world_drop] |
| back | Blisterbane Wrap (12552) (or Dark Phantom Cape (13122)) | Blackrock Depths: Anvilrage Overseer [dungeon] | 34.6 ranged_attack_power points (2.02 DPS) | yes | Dark Phantom Cape (13122, +0.00 DPS, sim-verified) [world_drop]; Blackveil Cape (11626, -0.13 DPS) [dungeon]; Duskbat Drape (19982, -0.13 DPS) [quest] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 57.7 ranged_attack_power points (3.37 DPS) | yes | Blazewind Breastplate (11193, -0.29 DPS, sim-verified) [quest]; Stone Guard's Chain Armor (220827, -0.56 DPS) [vendor]; Quillward Harness (10583, -0.81 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.6 ranged_attack_power points (2.02 DPS) | yes | Bracers of the Stone Princess (17714, -0.11 DPS, sim-verified) [dungeon]; Bloodlust Bracelets (14807, -0.54 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.54 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 50.8 ranged_attack_power points (2.97 DPS) | yes | Gloves of Holy Might (867, -1.01 DPS) [world_drop]; First Sergeant's Chain Gauntlets (220830, -1.04 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.08 DPS) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 47.6 ranged_attack_power points (2.78 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Defiler's Chain Girdle (20151, -0.82 DPS) [rep]; Defiler's Leather Girdle (20193, -0.82 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 48.5 ranged_attack_power points (2.83 DPS) | yes | Infernal Trickster Leggings (17754, +0.00 DPS, sim-verified) [dungeon]; Stone Guard's Chain Legplates (220833, -0.16 DPS) [vendor]; Keeper's Woolies (14668, -0.27 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 46.1 ranged_attack_power points (2.70 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.27 DPS) [world_drop]; Whisperwalk Boots (20255, -0.27 DPS) [quest] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.3 ranged_attack_power points (1.89 DPS) | yes | Ring of the Underwood (2951, -0.54 DPS) [world_drop]; Blackstone Ring (17713, -0.58 DPS) [dungeon]; Legionnaire's Band (19511, -0.67 DPS) [rep] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 ranged_attack_power points (1.40 DPS) | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Blackstone Ring (17713, -0.09 DPS) [dungeon]; Legionnaire's Band (19511, -0.19 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (141.0 DPS) | yes | Frozen Heart of the Mountain (249469, -4.88 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (141.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Fire Ruby (20036, -1.39 DPS, sim-verified) [quest] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (141.0 DPS) | yes | Steel Spear (250605, -0.76 DPS) [crafted]; Manslayer (10570, -0.88 DPS) [dungeon]; Hanzo Sword (8190, -3.48 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (141.0 DPS) | yes | Dark Iron Rifle (16004, -1.23 DPS, sim-verified) [crafted]; Hurricane (2824, -1.42 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -1.70 DPS) [world_drop] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Skibi's Pendant; shoulder: Blood Guard's Chain Epaulets; back: Blisterbane Wrap; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 744, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 232.4. Weights run: 1.5s. Verify run: 1.8s. 1579 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.339 ± 0.094, crit=1.438 ± 0.055 per rating point (14 rating = 1%, 20.135 per %), hit=0.343 ± 0.014 per rating point (10 rating = 1%, 3.434 per %), melee_haste=10.525 ± 1.657

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Visor (239532) | Leonid Barthalomew the Revered [vendor] | 196.3 ranged_attack_power points (11.14 DPS) | yes | Warlord's Chain Helmet (16566, -5.48 DPS) [vendor]; Warlord's Chain Helm (231571, -5.48 DPS) [vendor]; Dawnstalker Headpiece (239540, -6.78 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 44.4 ranged_attack_power points (2.52 DPS) | yes | Medallion of the Dawn (22659, +0.00 DPS, sim-verified) [quest]; Pendant of Celerity (22340, -0.51 DPS) [dungeon]; Maelstrom's Tendril (19620, -0.53 DPS) [quest] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | 143.1 ranged_attack_power points (8.12 DPS) | yes | Dawnstalker Spaulders (239542, -3.35 DPS, sim-verified) [vendor]; Darkspear Pauldrons (272105, -3.39 DPS) [vendor]; Darkspear Epaulets (272106, -3.39 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 55.1 ranged_attack_power points (3.13 DPS) | yes | Deathguard's Cloak (20068, -0.78 DPS, sim-verified) [rep]; Shifting Cloak (18511, -0.87 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.87 DPS) [dungeon] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (232.4 DPS) | yes | Dawnstalker Tunic (239543, -1.70 DPS) [vendor]; Dawn Armor (252483, -5.38 DPS) [crafted]; Tunic of Undead Slaying (23089, -16.79 DPS, sim-verified) [world] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | sim-verified (232.4 DPS) | yes | Dawnstalker Wristguards (239544, -0.66 DPS) [vendor]; Bracers of the Eclipse (18375, -2.96 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.35 DPS, sim-verified) [world] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | 145.6 ranged_attack_power points (8.27 DPS) | yes | Dawnstalker Handguards (239539, -1.45 DPS, sim-verified) [vendor]; General's Chain Gloves (16571, -4.34 DPS) [vendor]; General's Chain Vices (231575, -4.34 DPS) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 122.3 ranged_attack_power points (6.94 DPS) | yes | Dawnstalker Girdle (239538, -1.37 DPS, sim-verified) [vendor]; Molten Belt (19163, -3.23 DPS) [crafted]; Dense Timbermaw Belt (227807, -3.31 DPS) [vendor] |
| legs | Dawnstalker Leggings (239533) | Leonid Barthalomew the Revered [vendor] | 193.9 ranged_attack_power points (11.01 DPS) | yes | Sentinel's Chain Leggings (237819, -3.88 DPS) [vendor]; Outrider's Chain Leggings (22673, -5.02 DPS) [rep]; Dawnstalker Legguards (239541, -6.95 DPS, sim-verified) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | 154.3 ranged_attack_power points (8.76 DPS) | yes | Dawnstalker Boots (239537, -3.81 DPS, sim-verified) [vendor]; General's Chain Greaves (231570, -5.11 DPS) [vendor]; General's Chain Sabatons (231564, -5.25 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (232.4 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.27 DPS) [quest]; Naglering (11669, -6.98 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (232.4 DPS) | yes | Cutthroat's Signet (272408, -0.19 DPS) [vendor]; Tarnished Elven Ring (18500, -0.24 DPS) [dungeon]; Naglering (11669, -5.78 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (232.4 DPS) | yes | Counterattack Lodestone (18537, -3.66 DPS) [dungeon]; Hand of Justice (11815, -3.77 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -4.73 DPS) [crafted] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (232.4 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -1.82 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (232.4 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -7.66 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (232.4 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -8.97 DPS, sim-verified) [crafted] |

**New at 60:** head: Dawnstalker Visor; neck: Amulet of the Darkmoon; shoulder: Dawnstalker Pauldrons; back: Cape of the Black Baron; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Leggings; feet: Dawnstalker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Burst of Knowledge; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1579, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

