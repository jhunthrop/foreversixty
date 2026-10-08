# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 110.7. Weights run: 1.7s. Verify run: 1.5s. 220 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.101 ± 0.014, crit=0.593 ± 0.016 per rating point (14 rating = 1%, 8.306 per %), hit=1.060 ± 0.050 per rating point (10 rating = 1%, 10.605 per %), melee_haste=8.276 ± 0.789

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.8 ranged_attack_power points (1.73 DPS) | yes | Red Winter Hat (21524, -1.81 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.47 DPS, sim-verified) [quest] |
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

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 144.4. Weights run: 1.8s. Verify run: 1.7s. 366 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.892 ± 0.013, crit=0.667 ± 0.017 per rating point (14 rating = 1%, 9.344 per %), hit=1.185 ± 0.057 per rating point (10 rating = 1%, 11.850 per %), melee_haste=9.439 ± 0.975

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Tribal Worg Helm (6204, -0.40 DPS) [world]; Brawler's Leather Hood (252504, -0.40 DPS) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Ghostshard Talisman (7731, -0.63 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Erudite's Amulet (277204, -0.79 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.18 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Cloak of Night (4447, -0.40 DPS) [world]; Fenrus' Hide (6340, -0.40 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Tunic of Westfall (2041, -0.67 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.19 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Serpent Gloves (5970, -0.40 DPS) [dungeon]; Insignia Gloves (6408, -0.40 DPS) [world_drop]; Gloves of the Fang (10413, -0.40 DPS) [dungeon] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.14 DPS) [rep]; Highlander's Leather Girdle (20117, -0.14 DPS) [rep] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.20 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Ring of Precision (1491, -0.59 DPS) [dungeon]; Protector's Band (19517, -0.59 DPS) [rep]; Signet of the Zhevra (285330, -0.59 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 20.2 ranged_attack_power points (1.39 DPS) | yes | Ring of Precision (1491, -0.20 DPS) [dungeon]; Protector's Band (19517, -0.20 DPS) [rep]; Signet of the Zhevra (285330, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 20.5 ranged_attack_power points (1.40 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.91 DPS) | yes | Silver Star (3463, -0.66 DPS, sim-verified) [quest]; Glass Shooter (9456, -0.72 DPS) [dungeon]; Ironweaver (13137, -1.32 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 193.3. Weights run: 1.9s. Verify run: 2.0s. 597 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.816 ± 0.013, crit=0.765 ± 0.020 per rating point (14 rating = 1%, 10.713 per %), hit=1.374 ± 0.095 per rating point (10 rating = 1%, 13.744 per %), melee_haste=9.757 ± 1.481

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.6 ranged_attack_power points (2.42 DPS) | yes | Guard's Chain Helm (250499, -0.19 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world]; Nightscape Headband (8176, -0.85 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 31.0 ranged_attack_power points (2.05 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.73 DPS) [quest]; Ghostshard Talisman (7731, -1.12 DPS) [dungeon]; Erudite's Amulet (277204, -1.30 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 43.0 ranged_attack_power points (2.84 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.83 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 28.2 ranged_attack_power points (1.86 DPS) | yes | Imperial Cloak (6432, -0.37 DPS) [dungeon]; Parachute Cloak (10518, -0.37 DPS) [crafted]; Tigerstrike Mantle (13108, -0.37 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 53.5 ranged_attack_power points (3.54 DPS) | yes | Wolffear Harness (13110, -0.37 DPS) [world_drop]; Nightscape Tunic (8175, -0.75 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.75 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.5 ranged_attack_power points (1.49 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Branded Leather Bracers (19508, -0.17 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.19 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.12 DPS) | yes | Gloves of Holy Might (867, -0.09 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.25 DPS) [crafted]; Scarlet Gauntlets (10331, -0.25 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.99 DPS) | yes | Scorpashi Sash (14652, -0.12 DPS) [world_drop]; Blackforge Girdle (6425, -0.31 DPS) [dungeon]; Ogron's Sash (13117, -0.31 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 59.1 ranged_attack_power points (3.91 DPS) | yes | Triprunner Dungarees (9624, -1.09 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.30 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.30 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 36.6 ranged_attack_power points (2.42 DPS) | yes | Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted]; Worn Running Boots (9398, -0.37 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.2 ranged_attack_power points (1.86 DPS) | yes | Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Protector's Band (19515, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 25.3 ranged_attack_power points (1.68 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Protector's Band (19515, -0.19 DPS) [rep]; Disengagement Ring (276202, -0.19 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (193.3 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.50 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Ginn-su Sword (9424), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.5 ranged_attack_power points (1.49 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Ginn-su Sword (9424, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (23.17 DPS) | yes | Shadowforge Bushmaster (9422, -2.08 DPS) [dungeon]; Swiftwind (13038, -2.17 DPS) [world_drop]; The Silencer (13138, -3.42 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 237.5. Weights run: 2.0s. Verify run: 2.5s. 754 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.765 ± 0.013, crit=0.875 ± 0.023 per rating point (14 rating = 1%, 12.249 per %), hit=1.577 ± 0.100 per rating point (10 rating = 1%, 15.767 per %), melee_haste=12.639 ± 1.609

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 56.5 ranged_attack_power points (3.86 DPS) | yes | Lordrec Helmet (10741, -0.84 DPS) [quest]; Sprightring Helm (17776, -1.03 DPS) [quest]; Helm of Fire (8348, -2.76 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.0 ranged_attack_power points (2.46 DPS) | yes | Sentinel's Medallion (19539, -0.19 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 48.2 ranged_attack_power points (3.30 DPS) | yes | Phytoskin Spaulders (17749, -0.27 DPS) [dungeon]; Sunburn Spaulders (274751, -0.40 DPS) [vendor]; Khan's Mantle (14787, -1.03 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 41.5 ranged_attack_power points (2.84 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.19 DPS) [dungeon]; Blackmetal Cape (9512, -0.57 DPS) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 69.1 ranged_attack_power points (4.73 DPS) | yes | Blazewind Breastplate (11193, -0.38 DPS) [quest]; Knight's Chain Armor (220828, -1.05 DPS) [vendor]; Quillward Harness (10583, -1.13 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.5 ranged_attack_power points (2.84 DPS) | yes | Bloodlust Bracelets (14807, -0.76 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.76 DPS) [crafted]; Arena Bands (18711, -0.92 DPS) [world] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 60.8 ranged_attack_power points (4.16 DPS) | yes | Sergeant Major's Chain Gauntlets (220829, -1.50 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.51 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.51 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.5 ranged_attack_power points (3.73 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.08 DPS) [crafted]; Stalker's Mail Belt (252588, -1.08 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.1 ranged_attack_power points (3.97 DPS) | yes | Infernal Trickster Leggings (17754, -0.19 DPS) [dungeon]; Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.49 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.3 ranged_attack_power points (3.78 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.57 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.7 ranged_attack_power points (2.65 DPS) | yes | Ring of the Underwood (2951, -0.76 DPS) [world_drop]; Falcon's Hook (7552, -0.95 DPS) [dungeon]; Ironspine's Eye (7686, -0.95 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.8 ranged_attack_power points (2.45 DPS) | yes | Ring of the Underwood (2951, -0.55 DPS) [world_drop]; Falcon's Hook (7552, -0.74 DPS) [dungeon]; Ironspine's Eye (7686, -0.74 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (237.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (237.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (237.5 DPS) | yes | Warmonger (13052, -1.11 DPS) [world_drop]; Steel Spear (250605, -1.21 DPS) [crafted]; Hanzo Sword (8190, -5.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (237.5 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.55 DPS) [world_drop]; Hurricane (2824, -1.67 DPS) [world_drop]; Dark Iron Rifle (16004, -3.07 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Molten Heart of the Mountain; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 314.3. Weights run: 2.0s. Verify run: 8.4s. 1670 eligible items had no known source.

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

Set DPS (verified): 786.9. Weights run: 2.1s. Verify run: 8.6s. 1670 eligible items had no known source.

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

Set DPS (verified): 111.4. Weights run: 1.7s. Verify run: 1.5s. 209 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.101 ± 0.014, crit=0.593 ± 0.016 per rating point (14 rating = 1%, 8.306 per %), hit=1.060 ± 0.050 per rating point (10 rating = 1%, 10.605 per %), melee_haste=8.276 ± 0.789

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.8 ranged_attack_power points (1.73 DPS) | yes | Red Winter Hat (21524, -1.84 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.44 DPS, sim-verified) [quest] |
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

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 145.9. Weights run: 1.8s. Verify run: 1.6s. 352 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.892 ± 0.013, crit=0.667 ± 0.017 per rating point (14 rating = 1%, 9.344 per %), hit=1.185 ± 0.057 per rating point (10 rating = 1%, 11.850 per %), melee_haste=9.439 ± 0.975

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Tribal Worg Helm (6204, -0.40 DPS) [world]; Brawler's Leather Hood (252504, -0.44 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Ghostshard Talisman (7731, -0.63 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Erudite's Amulet (277204, -0.79 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.18 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Cloak of Night (4447, -0.40 DPS) [world]; Swiftrunner Cape (6745, -0.40 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Panther Armor (6670, -1.07 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.19 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Jurassic Wristguards (6198), Insignia Bracers (6410)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Jurassic Wristguards (6198, +0.00 DPS) [world]; Insignia Bracers (6410, +0.00 DPS) [world_drop]; Loamflake Bracers (15462, -0.20 DPS) [quest] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Braced Handguards (6784, -0.20 DPS) [quest]; Serpent Gloves (5970, -0.40 DPS) [dungeon]; Insignia Gloves (6408, -0.40 DPS) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.14 DPS) [rep]; Defiler's Leather Girdle (20191, -0.14 DPS) [rep] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Ring of Precision (1491, -0.59 DPS) [dungeon]; Legionnaire's Band (19513, -0.59 DPS) [rep]; Signet of the Zhevra (285330, -0.59 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 20.2 ranged_attack_power points (1.39 DPS) | yes | Ring of Precision (1491, -0.20 DPS) [dungeon]; Legionnaire's Band (19513, -0.20 DPS) [rep]; Signet of the Zhevra (285330, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 20.5 ranged_attack_power points (1.40 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Prison Shank (2941, +0.00 DPS) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.91 DPS) | yes | Silver Star (3463, -0.04 DPS) [quest]; Glass Shooter (9456, -0.72 DPS) [dungeon]; Ironweaver (13137, -1.32 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Monkey Ring; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 195.3. Weights run: 1.9s. Verify run: 2.0s. 563 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.816 ± 0.013, crit=0.765 ± 0.020 per rating point (14 rating = 1%, 10.713 per %), hit=1.374 ± 0.095 per rating point (10 rating = 1%, 13.744 per %), melee_haste=9.757 ± 1.481

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.6 ranged_attack_power points (2.42 DPS) | yes | Nightscape Headband (8176, -0.19 DPS) [crafted]; Guard's Chain Helm (250499, -0.19 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 31.0 ranged_attack_power points (2.05 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.73 DPS) [quest]; Ghostshard Talisman (7731, -1.12 DPS) [dungeon]; Erudite's Amulet (277204, -1.30 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 43.0 ranged_attack_power points (2.84 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.84 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 28.2 ranged_attack_power points (1.86 DPS) | yes | Imperial Cloak (6432, -0.37 DPS) [dungeon]; Parachute Cloak (10518, -0.37 DPS) [crafted]; Tigerstrike Mantle (13108, -0.37 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 53.5 ranged_attack_power points (3.54 DPS) | yes | Wolffear Harness (13110, -0.37 DPS) [world_drop]; Nightscape Tunic (8175, -0.75 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.75 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.5 ranged_attack_power points (1.49 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.12 DPS) | yes | Gloves of Holy Might (867, -0.09 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.25 DPS) [crafted]; Scarlet Gauntlets (10331, -0.25 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.99 DPS) | yes | Scorpashi Sash (14652, -0.12 DPS) [world_drop]; Blackforge Girdle (6425, -0.31 DPS) [dungeon]; Ogron's Sash (13117, -0.31 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 59.1 ranged_attack_power points (3.91 DPS) | yes | Triprunner Dungarees (9624, -0.75 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.30 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.30 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 36.6 ranged_attack_power points (2.42 DPS) | yes | Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted]; Worn Running Boots (9398, -0.37 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.2 ranged_attack_power points (1.86 DPS) | yes | Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Legionnaire's Band (19512, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | Uldaman: Ancient Treasure [dungeon] | 25.3 ranged_attack_power points (1.68 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS) [dungeon]; Legionnaire's Band (19512, -0.19 DPS) [rep]; Disengagement Ring (276202, -0.19 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (195.3 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.98 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Ginn-su Sword (9424), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.5 ranged_attack_power points (1.49 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Ginn-su Sword (9424, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (23.17 DPS) | yes | Outrider's Bow (19560, -1.56 DPS) [pvp]; Shadowforge Bushmaster (9422, -2.08 DPS) [dungeon]; The Silencer (13138, -3.39 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 243.2. Weights run: 2.0s. Verify run: 2.5s. 713 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.765 ± 0.013, crit=0.875 ± 0.023 per rating point (14 rating = 1%, 12.249 per %), hit=1.577 ± 0.100 per rating point (10 rating = 1%, 15.767 per %), melee_haste=12.639 ± 1.609

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 47.0 ranged_attack_power points (3.22 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.38 DPS) [quest]; Tough Scorpid Helm (8208, -0.57 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.0 ranged_attack_power points (2.46 DPS) | yes | Scout's Medallion (19536, -0.38 DPS) [rep]; Woven Ivy Necklace (19159, -0.76 DPS) [quest] |
| shoulder | Phytoskin Spaulders (17749) | Maraudon: Razorlash [dungeon] | 44.2 ranged_attack_power points (3.03 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Sunburn Spaulders (274751, -0.12 DPS) [vendor]; Khan's Mantle (14787, -0.76 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) (or Blisterbane Wrap (12552)) | World drop [world_drop] | 41.5 ranged_attack_power points (2.84 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.19 DPS) [dungeon]; Blackmetal Cape (9512, -0.57 DPS) [dungeon] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 69.1 ranged_attack_power points (4.73 DPS) | yes | Blazewind Breastplate (11193, -0.38 DPS) [quest]; Stone Guard's Chain Armor (220827, -1.05 DPS) [vendor]; Quillward Harness (10583, -1.13 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.5 ranged_attack_power points (2.84 DPS) | yes | Bloodlust Bracelets (14807, +0.00 DPS) [world_drop]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 60.8 ranged_attack_power points (4.16 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -1.51 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.51 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.51 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.5 ranged_attack_power points (3.73 DPS) | yes | Sagebrush Girdle (17778, -0.89 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.08 DPS) [crafted]; Stalker's Mail Belt (252588, -1.08 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.1 ranged_attack_power points (3.97 DPS) | yes | Infernal Trickster Leggings (17754, +0.00 DPS, sim-verified) [dungeon]; Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Stone Guard's Chain Legplates (220833, -0.49 DPS) [vendor] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.3 ranged_attack_power points (3.78 DPS) | yes | Fleetfoot Greaves (11627, -0.19 DPS) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; Sandstalker Ankleguards (12470, -0.57 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.7 ranged_attack_power points (2.65 DPS) | yes | Ring of the Underwood (2951, -0.76 DPS) [world_drop]; Falcon's Hook (7552, -0.95 DPS) [dungeon]; Ironspine's Eye (7686, -0.95 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.8 ranged_attack_power points (2.45 DPS) | yes | Ring of the Underwood (2951, -0.55 DPS) [world_drop]; Falcon's Hook (7552, -0.74 DPS) [dungeon]; Ironspine's Eye (7686, -0.74 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (243.2 DPS) | yes | Frozen Heart of the Mountain (249469, -5.53 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (243.2 DPS) | yes | Frozen Heart of the Mountain (249469, -1.75 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (243.2 DPS) | yes | Warmonger (13052, -1.11 DPS) [world_drop]; Steel Spear (250605, -1.21 DPS) [crafted]; Hanzo Sword (8190, -5.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (243.2 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.55 DPS) [world_drop]; Hurricane (2824, -1.67 DPS) [world_drop]; Dark Iron Rifle (16004, -3.70 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Phytoskin Spaulders; back: Dark Phantom Cape; chest: Fungus Shroud Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Eyegouger; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 319.9. Weights run: 2.0s. Verify run: 8.3s. 1650 eligible items had no known source.

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

Set DPS (verified): 799.2. Weights run: 2.1s. Verify run: 8.8s. 1650 eligible items had no known source.

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

