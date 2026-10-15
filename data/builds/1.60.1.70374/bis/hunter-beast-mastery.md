# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 110.3. Weights run: 2.0s. Verify run: 1.8s. 212 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.101 ± 0.014, crit=0.593 ± 0.016 per rating point (14 rating = 1%, 8.306 per %), hit=1.060 ± 0.050 per rating point (10 rating = 1%, 10.605 per %), melee_haste=8.276 ± 0.789

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.8 ranged_attack_power points (1.73 DPS) | yes | Red Winter Hat (21524, -1.84 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.93 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -1.13 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Cape of the Brotherhood (5193, -0.22 DPS) [dungeon]; Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 34.1 ranged_attack_power points (2.37 DPS) | yes | Trapper's Leather Armor (252491, -0.86 DPS) [crafted]; Brawler's Leather Armor (252490, -0.91 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -1.08 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bravo's Armbands (270015, -0.22 DPS) [quest]; Ratchet Wristwraps (274742, -0.43 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (110.3 DPS) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -1.28 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Dusty Belt (279897, -0.17 DPS) [quest]; Dark Leather Belt (4249, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.9 ranged_attack_power points (1.94 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.8 ranged_attack_power points (1.73 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.43 DPS) [dungeon]; Agile Boots (4788, -0.65 DPS) [vendor] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.4 ranged_attack_power points (0.86 DPS) | yes | The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.2 ranged_attack_power points (0.43 DPS) | yes | The 1 Ring (8350, -0.92 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.8 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, -0.20 DPS) [dungeon]; Lupine Axe (1220, -0.85 DPS) [world]; Piercing Axe (6094, -0.85 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.48 DPS) | yes | Lil Timmy's Peashooter (13136, -1.21 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.01 DPS) [crafted]; Venomstrike (6469, -3.15 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 212, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 139.5. Weights run: 2.2s. Verify run: 1.9s. 382 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.892 ± 0.013, crit=0.667 ± 0.017 per rating point (14 rating = 1%, 9.344 per %), hit=1.211 ± 0.055 per rating point (10 rating = 1%, 12.106 per %), melee_haste=9.499 ± 0.922

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Brawler's Leather Hood (252504, -0.40 DPS) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world]; Resilient Cap (14401, -1.58 DPS) [world_drop] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.76 DPS) [world]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Erudite's Amulet (277204, -1.19 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.18 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.40 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.40 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Tunic of Westfall (2041, -0.65 DPS, sim-verified) [quest]; Garb of Florid Feathers (273026, -0.99 DPS) [dungeon]; Green Leather Armor (4255, -1.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.20 DPS) [world]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Serpent Gloves (5970, -0.59 DPS) [dungeon]; Insignia Gloves (6408, -0.59 DPS) [world_drop]; Pilferer's Gloves (7358, -0.84 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.14 DPS) [rep]; Highlander's Leather Girdle (20117, -0.14 DPS) [rep] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Footpads of the Fang (10411, -0.40 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) (or Protector's Band (19517)) | Blackfathom Deeps [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Insurgent's Band (272067, -0.57 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Insurgent's Band (272067, -0.57 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) (or Sentinel's Blade (212583)) | Razorfen Kraul [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Sentinel's Blade (212583) | Illiyana Moonblaze [vendor] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Jade Serpentblade (3850, +0.00 DPS) [crafted] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.91 DPS) | yes | Ironweaver (13137, -1.32 DPS) [world_drop]; Golemsight Long Gun (273029, -1.47 DPS, sim-verified) [dungeon]; Crystalpine Stinger (13037, -1.89 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Protector's Band; main_hand: Vendetta; off_hand: Sentinel's Blade; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 382, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 182.6. Weights run: 2.4s. Verify run: 2.3s. 602 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.797 ± 0.013, crit=0.727 ± 0.020 per rating point (14 rating = 1%, 10.182 per %), hit=1.484 ± 0.093 per rating point (10 rating = 1%, 14.837 per %), melee_haste=9.080 ± 1.427

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Nightscape Headband (8176, -0.18 DPS) [crafted]; Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 30.8 ranged_attack_power points (2.03 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.24 DPS) [world]; Kaleidoscope Chain (13084, -1.29 DPS) [world_drop]; Erudite's Amulet (277204, -1.66 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.8 ranged_attack_power points (2.82 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.83 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 47.6 ranged_attack_power points (3.14 DPS) | yes | Tough Scorpid Breastplate (8203, +0.00 DPS, sim-verified) [crafted]; Nightscape Tunic (8175, -0.37 DPS) [crafted]; Dusky Leather Armor (7374, -0.55 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Branded Leather Bracers (19508, -0.16 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.18 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.11 DPS) | yes | Gloves of Holy Might (867, -0.12 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.27 DPS) [crafted]; Scarlet Gauntlets (10331, -0.27 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.98 DPS) | yes | Scorpashi Sash (14652, -0.13 DPS) [world_drop]; Blackforge Girdle (6425, -0.32 DPS) [dungeon]; Ogron's Sash (13117, -0.32 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.7 ranged_attack_power points (3.88 DPS) | yes | Triprunner Dungarees (9624, -0.55 DPS) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.37 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.0 ranged_attack_power points (1.85 DPS) | yes | Protector's Band (19515, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor]; Mark of Kern (2262, -0.53 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.2 ranged_attack_power points (1.66 DPS) | yes | Protector's Band (19515, -0.18 DPS) [rep]; Disengagement Ring (276202, -0.18 DPS) [vendor]; Mark of Kern (2262, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (182.6 DPS) | yes | Hellslayer Battle Axe (13017, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -0.76 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (23.11 DPS) | yes | Swiftwind (13038, -2.17 DPS) [world_drop]; Baelog's Shortbow (9400, -2.21 DPS) [dungeon]; The Silencer (13138, -2.82 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Parachute Cloak; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 602, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 50 (dwarf, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 221.6. Weights run: 2.5s. Verify run: 2.9s. 758 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.771 ± 0.013, crit=0.900 ± 0.024 per rating point (14 rating = 1%, 12.597 per %), hit=1.628 ± 0.098 per rating point (10 rating = 1%, 16.277 per %), melee_haste=9.891 ± 1.505

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 56.9 ranged_attack_power points (3.90 DPS) | yes | Lordrec Helmet (10741, -0.86 DPS) [quest]; Sprightring Helm (17776, -1.05 DPS) [quest]; Helm of Fire (8348, -1.35 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.0 ranged_attack_power points (2.47 DPS) | yes | Sentinel's Medallion (19539, -0.19 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 48.6 ranged_attack_power points (3.33 DPS) | yes | Sunburn Spaulders (274751, -0.69 DPS, sim-verified) [vendor]; Khan's Mantle (14787, -1.05 DPS) [world_drop]; Nightscape Shoulders (8192, -1.24 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 41.6 ranged_attack_power points (2.84 DPS) | yes | Serpentskin Cloak (8259, -0.76 DPS) [world_drop]; Blackflame Cape (13109, -0.86 DPS, sim-verified) [world_drop]; Nightscape Cloak (8195, -0.95 DPS) [crafted] |
| chest | Blazewind Breastplate (11193) | Tremors of the Earth [quest] | 63.7 ranged_attack_power points (4.36 DPS) | yes | Knight's Chain Armor (220828, -0.66 DPS) [vendor]; Warbear Harness (15064, -0.95 DPS) [crafted]; Charred Leather Tunic (19127, -0.95 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.6 ranged_attack_power points (2.84 DPS) | yes | Wicked Leather Bracers (15084, -0.76 DPS) [crafted]; Arena Bands (18711, -0.93 DPS) [world]; Bloodlust Bracelets (14807, -1.03 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 61.0 ranged_attack_power points (4.17 DPS) | yes | Sergeant Major's Chain Gauntlets (220829, -1.31 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.52 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.52 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.6 ranged_attack_power points (3.73 DPS) | yes | Sagebrush Girdle (17778, -0.89 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.08 DPS) [crafted]; Stalker's Mail Belt (252588, -1.08 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.2 ranged_attack_power points (3.98 DPS) | yes | Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.47 DPS) [vendor]; Oilskin Leggings (9414, -0.57 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.4 ranged_attack_power points (3.79 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; Sergeant Major's Chain Sabatons (220837, -0.95 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.8 ranged_attack_power points (2.66 DPS) | yes | Falcon's Hook (7552, -0.95 DPS) [dungeon]; Protector's Band (19516, -0.95 DPS) [rep]; Disengagement Ring (276202, -1.14 DPS) [vendor] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.7 ranged_attack_power points (1.90 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Protector's Band (19516, -0.19 DPS) [rep]; Disengagement Ring (276202, -0.38 DPS) [vendor] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (221.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (221.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-verified (221.6 DPS) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Hanzo Sword (8190, -1.99 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 24.9 ranged_attack_power points (1.71 DPS) | yes | Widowmaker (4091, +0.00 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (221.6 DPS) | yes | Gryphonwing Long Bow (13022, -0.51 DPS) [dungeon]; Dark Iron Rifle (16004, -1.85 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -2.27 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Blazewind Breastplate; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Jhordy's Misplaced Screwdriver; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 758, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60 (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 293.4. Weights run: 2.4s. Verify run: 8.6s. 1388 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.832 ± 0.016, crit=1.317 ± 0.033 per rating point (14 rating = 1%, 18.432 per %), hit=2.045 ± 0.128 per rating point (10 rating = 1%, 20.454 per %), melee_haste=14.665 ± 1.837

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (293.4 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -9.18 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 53.8 ranged_attack_power points (3.69 DPS) | yes | Beads of Ogre Might (22150, -0.64 DPS) [quest]; Sentinel's Medallion (19538, -0.78 DPS) [rep]; Maelstrom's Tendril (19620, -0.78 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.9 ranged_attack_power points (6.50 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Field Marshal's Chain Spaulders (16468, -0.06 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -0.06 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.5 ranged_attack_power points (4.28 DPS) | yes | Howler's Furs (272414, -0.96 DPS) [vendor]; Cloak of the Honor Guard (20073, -0.98 DPS) [rep]; Shifting Cloak (18511, -0.98 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 143.7 ranged_attack_power points (9.84 DPS) | yes | Field Marshal's Chain Breastplate (16466, -1.98 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -1.98 DPS) [vendor]; Obsidian Mail Tunic (22191, -4.69 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (293.4 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -8.01 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (293.4 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (293.4 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -7.70 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 156.4 ranged_attack_power points (10.71 DPS) | yes | Marshal's Chain Legguards (231577, -2.86 DPS) [pvp]; Sentinel's Leather Pants (237818, -2.95 DPS) [vendor]; Marshal's Chain Legplates (231558, -3.69 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (293.4 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -8.07 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 88.4 ranged_attack_power points (6.06 DPS) | yes | Don Julio's Band (19325, -2.30 DPS) [rep]; Masons Fraternity Ring (9533, -3.34 DPS) [quest]; Protector's Band (19514, -3.92 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 60.1 ranged_attack_power points (4.12 DPS) | yes | Don Julio's Band (19325, -0.36 DPS) [rep]; Masons Fraternity Ring (9533, -1.40 DPS) [quest]; Protector's Band (19514, -1.98 DPS) [rep] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (293.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-verified (293.4 DPS) | yes | Darkmoon Card: Blue Dragon (19288, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (293.4 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -8.23 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (293.4 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -11.48 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket2: Darkmoon Card: Heroism; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1388, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 763.3. Weights run: 2.4s. Verify run: 8.5s. 1388 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.881 ± 0.016, crit=1.419 ± 0.031 per rating point (14 rating = 1%, 19.866 per %), hit=3.213 ± 0.182 per rating point (10 rating = 1%, 32.131 per %), melee_haste=17.459 ± 2.433

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 98.4 ranged_attack_power points (14.00 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -16.33 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 56.1 ranged_attack_power points (7.99 DPS) | yes | Medallion of the Dawn (22659, -1.75 DPS) [quest]; Sentinel's Medallion (19538, -1.84 DPS) [rep]; Amulet of the Darkmoon (19491, -2.10 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 97.7 ranged_attack_power points (13.90 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 63.2 ranged_attack_power points (9.00 DPS) | yes | Howler's Furs (272414, -0.44 DPS) [vendor]; Shifting Cloak (18511, -2.03 DPS) [crafted]; Cloak of the Honor Guard (20073, -2.11 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 158.6 ranged_attack_power points (22.57 DPS) | yes | Field Marshal's Chain Breastplate (16466, -5.80 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -5.80 DPS) [vendor]; Shadowcraft Tunic (16721, -10.06 DPS, sim-verified) [dungeon] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (763.3 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -10.53 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 77.8 ranged_attack_power points (11.07 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Marshal's Chain Grips (231560, -1.01 DPS) [pvp]; Mourning Handguards (279257, -2.46 DPS) [crafted] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (763.3 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -8.30 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 172.7 ranged_attack_power points (24.58 DPS) | yes | Marshal's Chain Legguards (231577, -7.81 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.85 DPS) [vendor]; Marshal's Chain Legplates (231558, -9.66 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (763.3 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -13.53 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 101.3 ranged_attack_power points (14.41 DPS) | yes | Don Julio's Band (19325, -4.74 DPS) [rep]; Masons Fraternity Ring (9533, -8.67 DPS) [quest]; Channeler's Ring (272406, -9.84 DPS) [vendor] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 72.5 ranged_attack_power points (10.31 DPS) | yes | Don Julio's Band (19325, -0.64 DPS) [rep]; Masons Fraternity Ring (9533, -4.57 DPS) [quest]; Channeler's Ring (272406, -5.74 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (763.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (763.3 DPS) | yes | Darkmoon Card: Heroism (19287, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (763.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -19.48 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (763.3 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Bow of Searing Arrows (2825, -29.82 DPS, sim-verified) [world_drop] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1388, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 110.3. Weights run: 2.0s. Verify run: 1.7s. 201 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.101 ± 0.014, crit=0.593 ± 0.016 per rating point (14 rating = 1%, 8.306 per %), hit=1.060 ± 0.050 per rating point (10 rating = 1%, 10.605 per %), melee_haste=8.276 ± 0.789

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.8 ranged_attack_power points (1.73 DPS) | yes | Red Winter Hat (21524, -1.85 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.96 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -0.37 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.48 DPS, sim-verified) [dungeon] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 21.7 ranged_attack_power points (1.51 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Dark Leather Tunic (2317, -0.22 DPS) [crafted]; Prospector's Chestpiece (14562, -0.22 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bristlebark Bindings (14569, -0.43 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.43 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (110.3 DPS) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -1.14 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Guardsman Belt (3429, -0.39 DPS) [world]; Dark Leather Belt (4249, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.9 ranged_attack_power points (1.94 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.8 ranged_attack_power points (1.73 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.43 DPS) [dungeon]; Agile Boots (4788, -0.65 DPS) [vendor] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.4 ranged_attack_power points (0.86 DPS) | yes | Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 9.3 ranged_attack_power points (0.65 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.43 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.8 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Crescent Staff (6505, -0.63 DPS) [quest]; Lupine Axe (1220, -0.85 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.48 DPS) | yes | Outrider's Bow (20437, -0.68 DPS) [pvp]; Lil Timmy's Peashooter (13136, -1.47 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.01 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Bounty Hunter's Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 201, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 209611 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 141.2. Weights run: 2.2s. Verify run: 1.9s. 365 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.892 ± 0.013, crit=0.667 ± 0.017 per rating point (14 rating = 1%, 9.344 per %), hit=1.211 ± 0.055 per rating point (10 rating = 1%, 12.106 per %), melee_haste=9.499 ± 0.922

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Brawler's Leather Hood (252504, -0.42 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world]; Resilient Cap (14401, -1.58 DPS) [world_drop] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.76 DPS) [world]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Erudite's Amulet (277204, -1.19 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.18 DPS) | yes | Mantle of Thieves (2264, -0.70 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.40 DPS) [dungeon]; Swiftrunner Cape (6745, -0.40 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Garb of Florid Feathers (273026, +0.00 DPS, sim-verified) [dungeon]; Panther Armor (6670, -0.99 DPS) [quest]; Green Leather Armor (4255, -1.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop]; Loamflake Bracers (15462, -0.20 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Braced Handguards (6784, -0.40 DPS) [quest]; Insignia Gloves (6408, -0.59 DPS) [world_drop]; Pilferer's Gloves (7358, -0.88 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.14 DPS) [rep]; Defiler's Leather Girdle (20191, -0.14 DPS) [rep] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 40.5 ranged_attack_power points (2.77 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ring of Precision (1491) (or Legionnaire's Band (19513)) | Blackfathom Deeps [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Band of the Fist (17694, -0.40 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Band of the Fist (17694, -0.40 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) (or Sentinel's Blade (212583)) | Razorfen Kraul [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Sentinel's Blade (212583) | Illiyana Moonblaze [vendor] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Outlaw Sabre (16886, +0.00 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.91 DPS) | yes | Ironweaver (13137, -1.32 DPS) [world_drop]; Outrider's Bow (19561, -1.53 DPS) [pvp]; Golemsight Long Gun (273029, -1.72 DPS, sim-verified) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Legionnaire's Band; main_hand: Vendetta; off_hand: Sentinel's Blade; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 365, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 184.6. Weights run: 2.4s. Verify run: 2.3s. 566 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.797 ± 0.013, crit=0.727 ± 0.020 per rating point (14 rating = 1%, 10.182 per %), hit=1.484 ± 0.093 per rating point (10 rating = 1%, 14.837 per %), melee_haste=9.080 ± 1.427

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Nightscape Headband (8176, -0.18 DPS) [crafted]; Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 30.8 ranged_attack_power points (2.03 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.24 DPS) [world]; Ethereal Talisman (4430, -1.29 DPS) [quest]; Kaleidoscope Chain (13084, -1.29 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.8 ranged_attack_power points (2.82 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.84 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 47.6 ranged_attack_power points (3.14 DPS) | yes | Nightscape Tunic (8175, -0.37 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.37 DPS) [crafted]; Dusky Leather Armor (7374, -0.55 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.11 DPS) | yes | Gloves of Holy Might (867, -0.12 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.27 DPS) [crafted]; Scarlet Gauntlets (10331, -0.27 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.98 DPS) | yes | Scorpashi Sash (14652, -0.13 DPS) [world_drop]; Blackforge Girdle (6425, -0.32 DPS) [dungeon]; Ogron's Sash (13117, -0.32 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.7 ranged_attack_power points (3.88 DPS) | yes | Triprunner Dungarees (9624, -0.62 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.37 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.0 ranged_attack_power points (1.85 DPS) | yes | Legionnaire's Band (19512, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor]; Mark of Kern (2262, -0.53 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.2 ranged_attack_power points (1.66 DPS) | yes | Disengagement Ring (276202, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Band (19512, -0.18 DPS) [rep]; Mark of Kern (2262, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (184.6 DPS) | yes | Hellslayer Battle Axe (13017, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.25 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (23.11 DPS) | yes | Outrider's Bow (19560, -1.56 DPS) [pvp]; Swiftwind (13038, -2.17 DPS) [world_drop]; The Silencer (13138, -4.14 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Parachute Cloak; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 566, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 228.5. Weights run: 2.5s. Verify run: 2.8s. 715 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.771 ± 0.013, crit=0.900 ± 0.024 per rating point (14 rating = 1%, 12.597 per %), hit=1.628 ± 0.098 per rating point (10 rating = 1%, 16.277 per %), melee_haste=9.891 ± 1.505

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 47.1 ranged_attack_power points (3.22 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.38 DPS) [quest]; Tough Scorpid Helm (8208, -0.57 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.0 ranged_attack_power points (2.47 DPS) | yes | Scout's Medallion (19535, -0.19 DPS) [rep]; Woven Ivy Necklace (19159, -0.76 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.5 ranged_attack_power points (2.91 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Khan's Mantle (14787, -0.63 DPS) [world_drop]; Nightscape Shoulders (8192, -0.82 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 41.6 ranged_attack_power points (2.84 DPS) | yes | Blackflame Cape (13109, -0.75 DPS, sim-verified) [world_drop]; Serpentskin Cloak (8259, -0.76 DPS) [world_drop]; Nightscape Cloak (8195, -0.95 DPS) [crafted] |
| chest | Blazewind Breastplate (11193) | Broken Alliances [quest] | 63.7 ranged_attack_power points (4.36 DPS) | yes | Stone Guard's Chain Armor (220827, +0.00 DPS, sim-verified) [vendor]; Warbear Harness (15064, -0.95 DPS) [crafted]; Charred Leather Tunic (19127, -0.95 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.6 ranged_attack_power points (2.84 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Bloodlust Bracelets (14807, -0.67 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 61.0 ranged_attack_power points (4.17 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -1.34 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.52 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.52 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.6 ranged_attack_power points (3.73 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.08 DPS) [crafted]; Stalker's Mail Belt (252588, -1.08 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.2 ranged_attack_power points (3.98 DPS) | yes | Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Stone Guard's Chain Legplates (220833, -0.47 DPS) [vendor]; Oilskin Leggings (9414, -0.57 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.4 ranged_attack_power points (3.79 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; First Sergeant's Chain Sabatons (220836, -0.95 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.8 ranged_attack_power points (2.66 DPS) | yes | Falcon's Hook (7552, -0.95 DPS) [dungeon]; Legionnaire's Band (19511, -0.95 DPS) [rep]; White Bone Band (11862, -1.01 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.7 ranged_attack_power points (1.90 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Legionnaire's Band (19511, -0.19 DPS) [rep]; White Bone Band (11862, -0.25 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -6.73 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (228.5 DPS) | yes | White Bone Shredder (11863, -0.38 DPS) [quest]; Force of the Hippogryph (9684, -2.80 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -0.51 DPS) [dungeon]; Dark Iron Rifle (16004, -1.67 DPS, sim-verified) [crafted]; Outrider's Bow (19559, -2.25 DPS) [pvp] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Blazewind Breastplate; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Jhordy's Misplaced Screwdriver; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 715, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 307.3. Weights run: 2.4s. Verify run: 8.2s. 1368 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.832 ± 0.016, crit=1.317 ± 0.033 per rating point (14 rating = 1%, 18.432 per %), hit=2.045 ± 0.128 per rating point (10 rating = 1%, 20.454 per %), melee_haste=14.665 ± 1.837

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (307.3 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -9.77 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 53.8 ranged_attack_power points (3.69 DPS) | yes | Scout's Medallion (19534, -0.78 DPS) [rep]; Maelstrom's Tendril (19620, -0.78 DPS) [quest]; Beads of Ogre Might (22150, -0.88 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.9 ranged_attack_power points (6.50 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -0.06 DPS) [pvp]; Warlord's Chain Pauldrons (231565, -0.66 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.5 ranged_attack_power points (4.28 DPS) | yes | Deathguard's Cloak (20068, -0.98 DPS) [rep]; Shifting Cloak (18511, -0.98 DPS) [crafted]; Howler's Furs (272414, -1.22 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 143.7 ranged_attack_power points (9.84 DPS) | yes | Warlord's Chain Chestpiece (16565, -1.98 DPS) [vendor]; Warlord's Chain Hauberk (231573, -1.98 DPS) [vendor]; Obsidian Mail Tunic (22191, -5.01 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (307.3 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -9.33 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (307.3 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (307.3 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -8.79 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 156.4 ranged_attack_power points (10.71 DPS) | yes | Outrider's Chain Leggings (22673, -1.98 DPS, sim-verified) [rep]; General's Chain Legguards (231574, -2.86 DPS) [pvp]; Sentinel's Leather Pants (237818, -2.95 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (307.3 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -8.87 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 88.4 ranged_attack_power points (6.06 DPS) | yes | Don Julio's Band (19325, -2.30 DPS) [rep]; Masons Fraternity Ring (9533, -3.34 DPS) [quest]; Legionnaire's Band (19510, -3.92 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 60.1 ranged_attack_power points (4.12 DPS) | yes | Don Julio's Band (19325, -0.36 DPS) [rep]; Masons Fraternity Ring (9533, -1.40 DPS) [quest]; Legionnaire's Band (19510, -1.98 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (307.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (307.3 DPS) | yes | Molten Heart of the Mountain (249470, -1.41 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (307.3 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -8.34 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (307.3 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -11.44 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1368, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 780.9. Weights run: 2.4s. Verify run: 8.3s. 1368 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.881 ± 0.016, crit=1.419 ± 0.031 per rating point (14 rating = 1%, 19.866 per %), hit=3.213 ± 0.182 per rating point (10 rating = 1%, 32.131 per %), melee_haste=17.459 ± 2.433

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 98.4 ranged_attack_power points (14.00 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -16.47 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 56.1 ranged_attack_power points (7.99 DPS) | yes | Amulet of the Darkmoon (19491, -0.20 DPS) [quest]; Medallion of the Dawn (22659, -1.75 DPS) [quest]; Scout's Medallion (19534, -1.84 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 97.7 ranged_attack_power points (13.90 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -0.01 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 63.2 ranged_attack_power points (9.00 DPS) | yes | Shifting Cloak (18511, -2.03 DPS) [crafted]; Deathguard's Cloak (20068, -2.11 DPS) [rep]; Howler's Furs (272414, -3.22 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 158.6 ranged_attack_power points (22.57 DPS) | yes | Warlord's Chain Chestpiece (16565, -5.80 DPS) [vendor]; Warlord's Chain Hauberk (231573, -5.80 DPS) [vendor]; Shadowcraft Tunic (16721, -10.48 DPS, sim-verified) [dungeon] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (780.9 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -11.31 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 77.8 ranged_attack_power points (11.07 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; General's Chain Grips (231569, -1.01 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (780.9 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -9.03 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 172.7 ranged_attack_power points (24.58 DPS) | yes | Outrider's Chain Leggings (22673, -3.96 DPS, sim-verified) [rep]; General's Chain Legguards (231574, -7.81 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.85 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (780.9 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -13.22 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 101.3 ranged_attack_power points (14.41 DPS) | yes | Don Julio's Band (19325, -4.74 DPS) [rep]; Masons Fraternity Ring (9533, -8.67 DPS) [quest]; Channeler's Ring (272406, -9.84 DPS) [vendor] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 72.5 ranged_attack_power points (10.31 DPS) | yes | Don Julio's Band (19325, -0.64 DPS) [rep]; Masons Fraternity Ring (9533, -4.57 DPS) [quest]; Channeler's Ring (272406, -5.74 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (780.9 DPS) | yes | Frozen Heart of the Mountain (249469, -13.89 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (780.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (780.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -18.37 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (780.9 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -26.82 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1368, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

