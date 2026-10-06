# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 75.5. Weights run: 2.6s. Verify run: 1.6s. 220 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.139 ± 0.027, crit=0.489 ± 0.014 per rating point (14 rating = 1%, 6.849 per %), hit=0.165 ± 0.005 per rating point (10 rating = 1%, 1.653 per %), melee_haste=3.588 ± 0.565

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.1 ranged_attack_power points (1.02 DPS) | yes | Red Winter Hat (21524, -1.05 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.7 ranged_attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.81 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Hide of Lupos (3018, -0.26 DPS) [world]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 23.5 ranged_attack_power points (1.40 DPS) | yes | Trapper's Leather Armor (252491, -0.51 DPS) [crafted]; Brawler's Leather Armor (252490, -0.61 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.64 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.7 ranged_attack_power points (0.64 DPS) | yes | Wolf Bracers (4794, -0.13 DPS) [vendor]; Bravo's Armbands (270015, -0.13 DPS) [quest]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Forest Leather Gloves (3058, -0.26 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.26 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.07 DPS) | yes | Deviate Scale Belt (6468, -0.44 DPS) [crafted]; Dark Leather Belt (4249, -0.56 DPS) [crafted]; Dusty Belt (279897, -0.79 DPS, sim-verified) [quest] |
| legs | Leggings of the Fang (10410) (or Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.3 ranged_attack_power points (1.15 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.13 DPS) [world]; Brawler's Leather Pants (252500, -0.35 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.1 ranged_attack_power points (1.02 DPS) | yes | Blackened Defias Boots (10402, -0.26 DPS) [dungeon]; Footpads of the Fang (10411, -0.26 DPS, sim-verified) [dungeon]; Agile Boots (4788, -0.38 DPS) [vendor] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; The 1 Ring (8350, -0.64 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.6 ranged_attack_power points (0.51 DPS) | yes | Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; The 1 Ring (8350, -0.38 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 23.1 ranged_attack_power points (1.38 DPS) | yes | Impaling Harpoon (5200, -0.23 DPS) [dungeon]; Scythe Axe (5749, -0.49 DPS) [world]; Lupine Axe (1220, -0.61 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (10.65 DPS) | yes | Lil Timmy's Peashooter (13136, -1.62 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.29 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.53 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 5420001504000000-00000000000000000-000000000000000000)

Set DPS (verified): 99.6. Weights run: 2.8s. Verify run: 1.9s. 366 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.159 ± 0.032, crit=0.595 ± 0.017 per rating point (14 rating = 1%, 8.325 per %), hit=0.186 ± 0.005 per rating point (10 rating = 1%, 1.862 per %), melee_haste=6.489 ± 0.857

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 ranged_attack_power points (1.25 DPS) | yes | Tribal Worg Helm (6204, -0.25 DPS) [world]; Brawler's Leather Hood (252504, -0.26 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.38 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 17.3 ranged_attack_power points (1.00 DPS) | yes | Ghostshard Talisman (7731, -0.19 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.50 DPS) [world_drop]; Erudite's Amulet (277204, -0.50 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 ranged_attack_power points (1.38 DPS) | yes | Mantle of Thieves (2264, -0.13 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.50 DPS) [crafted]; Insignia Mantle (4721, -0.50 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.3 ranged_attack_power points (1.00 DPS) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Cloak of Night (4447, -0.25 DPS) [world]; Fenrus' Hide (6340, -0.25 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.2 ranged_attack_power points (1.76 DPS) | yes | Tunic of Westfall (2041, -0.39 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.75 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.75 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 13.0 ranged_attack_power points (0.75 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.13 DPS) [world_drop] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 17.3 ranged_attack_power points (1.00 DPS) | yes | Heavy Earthen Gloves (7359, -0.07 DPS) [crafted]; Serpent Gloves (5970, -0.25 DPS) [dungeon]; Insignia Gloves (6408, -0.25 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 ranged_attack_power points (1.39 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.27 DPS) [crafted]; Stalker's Leather Belt (252521, -0.27 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 30.2 ranged_attack_power points (1.76 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.13 DPS) [crafted]; Ferine Leggings (6690, -0.25 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.00 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.13 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.4 ranged_attack_power points (1.13 DPS) | yes | Ring of Precision (1491, -0.38 DPS) [dungeon]; Protector's Band (19517, -0.38 DPS) [rep]; Signet of the Zhevra (285330, -0.38 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.1 ranged_attack_power points (0.88 DPS) | yes | Protector's Band (19517, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Ring of Precision (1491, -0.27 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 ranged_attack_power points (0.98 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.0 ranged_attack_power points (0.75 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (15.20 DPS) | yes | Glass Shooter (9456, -0.61 DPS) [dungeon]; Ironweaver (13137, -1.12 DPS) [world_drop]; Silver Star (3463, -1.31 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Highlander's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 5420001505001251-00000000000000000-000000000000000000)

Set DPS (verified): 134.1. Weights run: 3.1s. Verify run: 2.0s. 597 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.166 ± 0.034, crit=0.652 ± 0.018 per rating point (14 rating = 1%, 9.133 per %), hit=0.201 ± 0.007 per rating point (10 rating = 1%, 2.012 per %), melee_haste=not significant (4.475 ± 1.250)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 28.2 ranged_attack_power points (1.57 DPS) | yes | Nightscape Headband (8176, -0.12 DPS) [crafted]; Guard's Chain Helm (250499, -0.12 DPS) [crafted]; Skullsplitter Helm (1624, -0.24 DPS) [world] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 23.8 ranged_attack_power points (1.33 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.21 DPS) [quest]; Ghostshard Talisman (7731, -0.55 DPS) [dungeon]; Erudite's Amulet (277204, -0.85 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 35.8 ranged_attack_power points (2.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.67 DPS) [world_drop]; Nightscape Shoulders (8192, -0.69 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.79 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 21.7 ranged_attack_power points (1.21 DPS) | yes | Imperial Cloak (6432, -0.24 DPS) [dungeon]; Parachute Cloak (10518, -0.24 DPS) [crafted]; Tigerstrike Mantle (13108, -0.24 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 41.1 ranged_attack_power points (2.29 DPS) | yes | Wolffear Harness (13110, -0.24 DPS) [world_drop]; Nightscape Tunic (8175, -0.48 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.48 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.12 DPS) | yes | Imperial Leather Bracers (4061, -0.15 DPS) [dungeon]; Dusky Bracers (7378, -0.15 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.27 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (1.78 DPS) | yes | Gloves of Holy Might (867, -0.16 DPS) [world_drop]; Dragonscale Gauntlets (8347, -0.55 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.58 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.67 DPS) | yes | Highlander's Chain Girdle (20090, -0.33 DPS) [rep]; Scorpashi Sash (14652, -0.47 DPS) [world_drop]; Blackforge Girdle (6425, -0.59 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.5 ranged_attack_power points (2.54 DPS) | yes | Triprunner Dungarees (9624, -0.36 DPS) [quest]; Petrolspill Leggings (9509, -0.85 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.85 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.2 ranged_attack_power points (1.57 DPS) | yes | Imperial Leather Boots (6431, -0.24 DPS) [dungeon]; Dusky Boots (7390, -0.24 DPS) [crafted]; Worn Running Boots (9398, -0.24 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 21.7 ranged_attack_power points (1.21 DPS) | yes | Assault Band (13095, -0.09 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 ranged_attack_power points (1.12 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.03 DPS) [dungeon]; Ironspine's Eye (7686, -0.03 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (134.1 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.26 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 19.5 ranged_attack_power points (1.09 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS, sim-verified) [crafted]; Satyr's Rod (15962, -0.97 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (19.52 DPS) | yes | Shadowforge Bushmaster (9422, -1.75 DPS) [dungeon]; Swiftwind (13038, -2.08 DPS) [world_drop]; The Silencer (13138, -2.86 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Mark of Kern; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5420001505001251-35200000000000000-000000000000000000)

Set DPS (verified): 159.0. Weights run: 3.1s. Verify run: 2.2s. 754 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.259 ± 0.044, crit=0.734 ± 0.020 per rating point (14 rating = 1%, 10.273 per %), hit=0.236 ± 0.007 per rating point (10 rating = 1%, 2.360 per %), melee_haste=not significant (4.207 ± 1.348)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 46.4 ranged_attack_power points (2.56 DPS) | yes | Lordrec Helmet (10741, -0.57 DPS) [quest]; Bloomsprout Headpiece (17767, -0.57 DPS) [dungeon]; Helm of Fire (8348, -2.47 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.4 ranged_attack_power points (1.62 DPS) | yes | Sentinel's Medallion (19540, -0.25 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.52 DPS) [quest] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 39.6 ranged_attack_power points (2.19 DPS) | yes | Phytoskin Spaulders (17749, -0.19 DPS) [dungeon]; Sunburn Spaulders (274751, -0.49 DPS, sim-verified) [vendor]; Khan's Mantle (14787, -0.69 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 33.9 ranged_attack_power points (1.87 DPS) | yes | Blackveil Cape (11626, -0.12 DPS) [dungeon]; Blackmetal Cape (9512, -0.37 DPS) [dungeon]; Blisterbane Wrap (12552, -0.51 DPS, sim-verified) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 56.5 ranged_attack_power points (3.12 DPS) | yes | Blazewind Breastplate (11193, -0.25 DPS) [quest]; Knight's Chain Armor (220828, -0.68 DPS) [vendor]; Quillward Harness (10583, -0.75 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.9 ranged_attack_power points (1.87 DPS) | yes | Bracers of the Stone Princess (17714, -0.32 DPS) [dungeon]; Arena Bands (18711, -0.32 DPS) [world]; Bloodlust Bracelets (14807, -0.50 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 49.7 ranged_attack_power points (2.74 DPS) | yes | Sergeant Major's Chain Gauntlets (220829, -0.96 DPS, sim-verified) [vendor]; Gauntlets of Divinity (7724, -0.98 DPS) [dungeon]; Beastmaster's Gauntlets (226883, -1.00 DPS) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 46.9 ranged_attack_power points (2.59 DPS) | yes | Sagebrush Girdle (17778, -0.46 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -0.84 DPS) [crafted]; Stalker's Mail Belt (252588, -0.84 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 47.4 ranged_attack_power points (2.62 DPS) | yes | Infernal Trickster Leggings (17754, +0.00 DPS, sim-verified) [dungeon]; Keeper's Woolies (14668, -0.25 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.31 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 45.2 ranged_attack_power points (2.49 DPS) | yes | Fleetfoot Greaves (11627, -0.12 DPS) [dungeon]; Elven Chain Boots (13125, -0.25 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.6 ranged_attack_power points (1.75 DPS) | yes | Blackstone Ring (17713, -0.51 DPS) [dungeon]; Falcon's Hook (7552, -0.62 DPS) [dungeon]; Ironspine's Eye (7686, -0.62 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 22.6 ranged_attack_power points (1.25 DPS) | yes | Blackstone Ring (17713, -0.01 DPS) [dungeon]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (159.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (159.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (159.0 DPS) | yes | Steel Spear (250605, -0.69 DPS) [crafted]; Manslayer (10570, -0.77 DPS) [dungeon]; Hanzo Sword (8190, -3.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (159.0 DPS) | yes | Hurricane (2824, -1.34 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -1.64 DPS) [world_drop]; Dark Iron Rifle (16004, -2.39 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Devilsaur Eye; trinket2: Molten Heart of the Mountain; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5420001505001251-35510000000000000-510000000000000000)

Set DPS (verified): 200.1. Weights run: 3.1s. Verify run: 2.2s. 1670 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.167 ± 0.042, crit=1.067 ± 0.029 per rating point (14 rating = 1%, 14.943 per %), hit=0.352 ± 0.010 per rating point (10 rating = 1%, 3.525 per %), melee_haste=12.115 ± 1.560

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (200.1 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -3.81 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 41.2 ranged_attack_power points (2.21 DPS) | yes | Medallion of the Dawn (22659, -0.12 DPS) [quest]; Pendant of Celerity (22340, -0.45 DPS) [dungeon]; Sentinel's Medallion (19538, -0.46 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 73.4 ranged_attack_power points (3.94 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Highlander's Leather Shoulders (20059, -0.24 DPS) [rep]; Field Marshal's Chain Spaulders (16468, -0.73 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 52.5 ranged_attack_power points (2.81 DPS) | yes | Cloak of the Honor Guard (20073, -0.63 DPS, sim-verified) [rep]; Shifting Cloak (18511, -0.84 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.84 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Obsidian Mail Tunic (22191, -0.41 DPS) [crafted]; Field Marshal's Chain Breastplate (16466, -0.54 DPS) [vendor]; Tunic of Undead Slaying (23089, -8.23 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Bracers (16461, -0.13 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.24 DPS) [rep]; Wristwraps of Undead Slaying (23093, -3.08 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 58.5 ranged_attack_power points (3.14 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Gauntlets of Deftness (22410, -0.58 DPS, sim-verified) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 ranged_attack_power points (3.43 DPS) | yes | Ranger's Belt (272397, +0.00 DPS, sim-verified) [vendor]; Highlander's Chain Girdle (20043, -0.81 DPS) [rep]; Highlander's Leather Girdle (20045, -0.81 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 109.2 ranged_attack_power points (5.85 DPS) | yes | Marshal's Chain Legguards (231577, -1.11 DPS) [pvp]; Sentinel's Leather Pants (237818, -1.12 DPS) [vendor]; Marshal's Chain Legplates (231558, -1.15 DPS) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 55.3 ranged_attack_power points (2.97 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Beastmaster's Boots (22061, +0.00 DPS) [quest]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.46 DPS) [vendor]; Tarnished Elven Ring (18500, -0.52 DPS) [dungeon]; Naglering (11669, -4.19 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.03 DPS) [vendor]; Tarnished Elven Ring (18500, -0.09 DPS) [dungeon]; Naglering (11669, -3.68 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Devilsaur Eye (19991, -1.53 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -5.85 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -7.05 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Bracers of the Eclipse; hands: Raider Gloves; waist: Dense Timbermaw Belt; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 5420000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 75.0. Weights run: 2.6s. Verify run: 1.6s. 209 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.139 ± 0.027, crit=0.489 ± 0.014 per rating point (14 rating = 1%, 6.849 per %), hit=0.165 ± 0.005 per rating point (10 rating = 1%, 1.653 per %), melee_haste=3.588 ± 0.565

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.1 ranged_attack_power points (1.02 DPS) | yes | Red Winter Hat (21524, -1.00 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.7 ranged_attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Cape of the Brotherhood (5193, -0.23 DPS, sim-verified) [dungeon]; Hide of Lupos (3018, -0.26 DPS) [world]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 15.0 ranged_attack_power points (0.89 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.13 DPS) [crafted]; Prospector's Chestpiece (14562, -0.13 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.7 ranged_attack_power points (0.64 DPS) | yes | Wolf Bracers (4794, -0.13 DPS) [vendor]; Bristlebark Bindings (14569, -0.26 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Forest Leather Gloves (3058, -0.26 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.26 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.07 DPS) | yes | Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.56 DPS) [world]; Dark Leather Belt (4249, -0.56 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.3 ranged_attack_power points (1.15 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.13 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.1 ranged_attack_power points (1.02 DPS) | yes | Footpads of the Fang (10411, -0.25 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.26 DPS) [dungeon]; Agile Boots (4788, -0.38 DPS) [vendor] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Bounty Hunter's Ring (5351, -0.38 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; The 1 Ring (8350, -0.64 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.6 ranged_attack_power points (0.51 DPS) | yes | Bounty Hunter's Ring (5351, -0.13 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; The 1 Ring (8350, -0.38 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 23.1 ranged_attack_power points (1.38 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.49 DPS) [world]; Crescent Staff (6505, -0.49 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (10.65 DPS) | yes | Outrider's Bow (20437, -0.64 DPS) [pvp]; Lil Timmy's Peashooter (13136, -0.74 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.29 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 5420001504000000-00000000000000000-000000000000000000)

Set DPS (verified): 100.1. Weights run: 2.8s. Verify run: 1.8s. 352 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.159 ± 0.032, crit=0.595 ± 0.017 per rating point (14 rating = 1%, 8.325 per %), hit=0.186 ± 0.005 per rating point (10 rating = 1%, 1.862 per %), melee_haste=6.489 ± 0.857

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 ranged_attack_power points (1.25 DPS) | yes | Tribal Worg Helm (6204, -0.25 DPS) [world]; Brawler's Leather Hood (252504, -0.26 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.38 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 17.3 ranged_attack_power points (1.00 DPS) | yes | Ghostshard Talisman (7731, -0.19 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.50 DPS) [world_drop]; Erudite's Amulet (277204, -0.50 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 ranged_attack_power points (1.38 DPS) | yes | Mantle of Thieves (2264, +0.00 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.50 DPS) [crafted]; Insignia Mantle (4721, -0.50 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.3 ranged_attack_power points (1.00 DPS) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Cloak of Night (4447, -0.25 DPS) [world]; Swiftrunner Cape (6745, -0.25 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.2 ranged_attack_power points (1.76 DPS) | yes | Panther Armor (6670, -0.67 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.75 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.75 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 13.0 ranged_attack_power points (0.75 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Loamflake Bracers (15462, -0.13 DPS) [quest] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 17.3 ranged_attack_power points (1.00 DPS) | yes | Heavy Earthen Gloves (7359, -0.07 DPS) [crafted]; Braced Handguards (6784, -0.13 DPS) [quest]; Insignia Gloves (6408, -0.25 DPS) [world_drop] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 ranged_attack_power points (1.39 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Deftkin Belt (16659, -0.20 DPS) [quest]; Skulker's Leather Belt (252520, -0.27 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 30.2 ranged_attack_power points (1.76 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.13 DPS) [crafted]; Ferine Leggings (6690, -0.25 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.00 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.4 ranged_attack_power points (1.13 DPS) | yes | Ring of Precision (1491, -0.38 DPS) [dungeon]; Legionnaire's Band (19513, -0.38 DPS) [rep]; Signet of the Zhevra (285330, -0.38 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.1 ranged_attack_power points (0.88 DPS) | yes | Legionnaire's Band (19513, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Ring of Precision (1491, -0.29 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.8 ranged_attack_power points (0.98 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Outlaw Sabre (16886) | Baron Aquanis [quest] | 15.0 ranged_attack_power points (0.87 DPS) | yes | Vendetta (776, +0.00 DPS) [dungeon]; Satyr's Rod (15962, -0.75 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (15.20 DPS) | yes | Glass Shooter (9456, -0.61 DPS) [dungeon]; Ironweaver (13137, -1.12 DPS) [world_drop]; Silver Star (3463, -1.29 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Defiler's Chain Girdle; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Outlaw Sabre; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 5420001505001251-00000000000000000-000000000000000000)

Set DPS (verified): 134.8. Weights run: 3.1s. Verify run: 2.0s. 563 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.166 ± 0.034, crit=0.652 ± 0.018 per rating point (14 rating = 1%, 9.133 per %), hit=0.201 ± 0.007 per rating point (10 rating = 1%, 2.012 per %), melee_haste=not significant (4.475 ± 1.250)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 28.2 ranged_attack_power points (1.57 DPS) | yes | Guard's Chain Helm (250499, -0.12 DPS) [crafted]; Skullsplitter Helm (1624, -0.24 DPS) [world]; Nightscape Headband (8176, -0.63 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 23.8 ranged_attack_power points (1.33 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.21 DPS) [quest]; Ghostshard Talisman (7731, -0.55 DPS) [dungeon]; Erudite's Amulet (277204, -0.85 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 35.8 ranged_attack_power points (2.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.67 DPS) [world_drop]; Nightscape Shoulders (8192, -0.69 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.79 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 21.7 ranged_attack_power points (1.21 DPS) | yes | Imperial Cloak (6432, -0.24 DPS) [dungeon]; Parachute Cloak (10518, -0.24 DPS) [crafted]; Tigerstrike Mantle (13108, -0.24 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 41.1 ranged_attack_power points (2.29 DPS) | yes | Wolffear Harness (13110, -0.24 DPS) [world_drop]; Nightscape Tunic (8175, -0.48 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.48 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.12 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Dusky Bracers (7378, -0.15 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (1.78 DPS) | yes | Gloves of Holy Might (867, -0.16 DPS) [world_drop]; Dragonscale Gauntlets (8347, -0.55 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.58 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.67 DPS) | yes | Defiler's Chain Girdle (20152, -0.33 DPS) [rep]; Scorpashi Sash (14652, -0.47 DPS) [world_drop]; Deftkin Belt (16659, -0.52 DPS) [quest] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.5 ranged_attack_power points (2.54 DPS) | yes | Triprunner Dungarees (9624, -0.60 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.85 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.85 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.2 ranged_attack_power points (1.57 DPS) | yes | Imperial Leather Boots (6431, -0.24 DPS) [dungeon]; Dusky Boots (7390, -0.24 DPS) [crafted]; Worn Running Boots (9398, -0.24 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 21.7 ranged_attack_power points (1.21 DPS) | yes | Assault Band (13095, -0.09 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 ranged_attack_power points (1.12 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.03 DPS) [dungeon]; Ironspine's Eye (7686, -0.03 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (134.8 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.27 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 19.5 ranged_attack_power points (1.09 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Satyr's Rod (15962, -0.97 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (19.52 DPS) | yes | Outrider's Bow (19560, -1.43 DPS) [pvp]; Shadowforge Bushmaster (9422, -1.75 DPS) [dungeon]; The Silencer (13138, -3.46 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Mark of Kern; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5420001505001251-35200000000000000-000000000000000000)

Set DPS (verified): 161.0. Weights run: 3.1s. Verify run: 2.2s. 713 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.259 ± 0.044, crit=0.734 ± 0.020 per rating point (14 rating = 1%, 10.273 per %), hit=0.236 ± 0.007 per rating point (10 rating = 1%, 2.360 per %), melee_haste=not significant (4.207 ± 1.348)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 38.4 ranged_attack_power points (2.12 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.13 DPS) [dungeon]; Sprightring Helm (17776, -0.25 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.4 ranged_attack_power points (1.62 DPS) | yes | Scout's Medallion (19535, -0.12 DPS) [rep]; Woven Ivy Necklace (19159, -0.50 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.8 ranged_attack_power points (2.03 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -0.04 DPS) [dungeon]; Khan's Mantle (14787, -0.54 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 33.9 ranged_attack_power points (1.87 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.12 DPS) [dungeon]; Blackmetal Cape (9512, -0.37 DPS) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 56.5 ranged_attack_power points (3.12 DPS) | yes | Blazewind Breastplate (11193, -0.25 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.68 DPS) [vendor]; Quillward Harness (10583, -0.75 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.9 ranged_attack_power points (1.87 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 49.7 ranged_attack_power points (2.74 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -0.97 DPS, sim-verified) [vendor]; Gauntlets of Divinity (7724, -0.98 DPS) [dungeon]; Beastmaster's Gauntlets (226883, -1.00 DPS) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 46.9 ranged_attack_power points (2.59 DPS) | yes | Sagebrush Girdle (17778, -0.72 DPS) [quest]; Skulker's Leather Waistguard (252474, -0.84 DPS) [crafted]; Stalker's Mail Belt (252588, -0.84 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 47.4 ranged_attack_power points (2.62 DPS) | yes | Infernal Trickster Leggings (17754, -0.12 DPS) [dungeon]; Keeper's Woolies (14668, -0.25 DPS) [world_drop]; Stone Guard's Chain Legplates (220833, -0.31 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 45.2 ranged_attack_power points (2.49 DPS) | yes | Fleetfoot Greaves (11627, -0.12 DPS) [dungeon]; Elven Chain Boots (13125, -0.25 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.6 ranged_attack_power points (1.75 DPS) | yes | Ring of the Underwood (2951, -0.50 DPS) [world_drop]; Blackstone Ring (17713, -0.51 DPS) [dungeon]; Falcon's Hook (7552, -0.62 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 ranged_attack_power points (1.32 DPS) | yes | Ring of the Underwood (2951, -0.08 DPS) [world_drop]; Blackstone Ring (17713, -0.09 DPS) [dungeon]; Falcon's Hook (7552, -0.20 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (161.0 DPS) | yes | Frozen Heart of the Mountain (249469, -4.61 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (161.0 DPS) | yes | Frozen Heart of the Mountain (249469, -1.47 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (161.0 DPS) | yes | Steel Spear (250605, -0.69 DPS) [crafted]; Manslayer (10570, -0.77 DPS) [dungeon]; Hanzo Sword (8190, -3.09 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (161.0 DPS) | yes | Dark Iron Rifle (16004, -1.10 DPS, sim-verified) [crafted]; Hurricane (2824, -1.34 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -1.64 DPS) [world_drop] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5420001505001251-35510000000000000-510000000000000000)

Set DPS (verified): 202.5. Weights run: 3.1s. Verify run: 2.2s. 1650 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.167 ± 0.042, crit=1.067 ± 0.029 per rating point (14 rating = 1%, 14.943 per %), hit=0.352 ± 0.010 per rating point (10 rating = 1%, 3.525 per %), melee_haste=12.115 ± 1.560

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (202.5 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -4.24 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 41.2 ranged_attack_power points (2.21 DPS) | yes | Medallion of the Dawn (22659, -0.12 DPS) [quest]; Pendant of Celerity (22340, -0.45 DPS) [dungeon]; Scout's Medallion (19534, -0.46 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 73.4 ranged_attack_power points (3.94 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Defiler's Leather Shoulders (20194, -0.24 DPS) [rep]; Warlord's Chain Shoulders (231572, -0.73 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 52.5 ranged_attack_power points (2.81 DPS) | yes | Deathguard's Cloak (20068, -0.60 DPS, sim-verified) [rep]; Shifting Cloak (18511, -0.84 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.84 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Obsidian Mail Tunic (22191, -0.41 DPS) [crafted]; Warlord's Chain Chestpiece (16565, -0.54 DPS) [vendor]; Tunic of Undead Slaying (23089, -8.25 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Wristguards (16570, -0.13 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.24 DPS) [rep]; Wristwraps of Undead Slaying (23093, -3.03 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 58.5 ranged_attack_power points (3.14 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Gauntlets of Deftness (22410, -0.60 DPS, sim-verified) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 ranged_attack_power points (3.43 DPS) | yes | Ranger's Belt (272397, +0.00 DPS, sim-verified) [vendor]; Defiler's Chain Girdle (20150, -0.81 DPS) [rep]; Defiler's Leather Girdle (20190, -0.81 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 109.2 ranged_attack_power points (5.85 DPS) | yes | General's Chain Legguards (231574, -1.11 DPS) [pvp]; Sentinel's Leather Pants (237818, -1.12 DPS) [vendor]; Outrider's Chain Leggings (22673, -1.20 DPS, sim-verified) [rep] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 55.3 ranged_attack_power points (2.97 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beastmaster's Boots (22061, -0.18 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.46 DPS) [vendor]; Tarnished Elven Ring (18500, -0.52 DPS) [dungeon]; Naglering (11669, -4.22 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.03 DPS) [vendor]; Tarnished Elven Ring (18500, -0.09 DPS) [dungeon]; Naglering (11669, -3.71 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Counterattack Lodestone (18537, -3.46 DPS) [dungeon]; Hand of Justice (11815, -3.56 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -4.46 DPS) [crafted] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Second Wind (11819, -2.85 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -5.85 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -6.31 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Bracers of the Eclipse; hands: Raider Gloves; waist: Dense Timbermaw Belt; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Burst of Knowledge; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

