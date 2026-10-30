# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 110.9. Weights run: 2.2s. Verify run: 1.9s. 217 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.112 ± 0.015, crit=0.605 ± 0.016 per rating point (14 rating = 1%, 8.470 per %), hit=1.109 ± 0.048 per rating point (10 rating = 1%, 11.089 per %), melee_haste=7.909 ± 0.726

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.9 ranged_attack_power points (1.72 DPS) | yes | Red Winter Hat (21524, -1.82 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.7 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.92 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.6 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -0.77 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.7 ranged_attack_power points (1.29 DPS) | yes | Cape of the Brotherhood (5193, -0.22 DPS) [dungeon]; Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 34.2 ranged_attack_power points (2.37 DPS) | yes | Trapper's Leather Armor (252491, -0.86 DPS) [crafted]; Brawler's Leather Armor (252490, -0.87 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -1.08 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.6 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bravo's Armbands (270015, -0.22 DPS) [quest]; Ratchet Wristwraps (274742, -0.43 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -1.79 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Dark Leather Belt (4249, -0.38 DPS) [crafted]; Dusty Belt (279897, -0.39 DPS, sim-verified) [quest] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 28.0 ranged_attack_power points (1.94 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Gray Leather Boots (2315, -0.22 DPS) [crafted]; Feet of the Lynx (1121, -1.47 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.4 ranged_attack_power points (0.86 DPS) | yes | The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.2 ranged_attack_power points (0.43 DPS) | yes | The 1 Ring (8350, -0.51 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.9 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, -0.20 DPS) [dungeon]; Lupine Axe (1220, -0.85 DPS) [world]; Piercing Axe (6094, -0.85 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.42 DPS) | yes | Lil Timmy's Peashooter (13136, -0.62 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.00 DPS) [crafted]; Venomstrike (6469, -3.13 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Protector's Band; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 138.8. Weights run: 2.3s. Verify run: 2.2s. 397 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.894 ± 0.013, crit=0.674 ± 0.017 per rating point (14 rating = 1%, 9.436 per %), hit=1.207 ± 0.058 per rating point (10 rating = 1%, 12.067 per %), melee_haste=11.026 ± 0.959

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.97 DPS) | yes | Brawler's Leather Hood (252504, -0.42 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world]; Resilient Cap (14401, -1.58 DPS) [world_drop] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.76 DPS) [world]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Erudite's Amulet (277204, -1.18 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.17 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.39 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.39 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.76 DPS) | yes | Tunic of Westfall (2041, -0.63 DPS, sim-verified) [quest]; Garb of Florid Feathers (273026, -0.99 DPS) [dungeon]; Green Leather Armor (4255, -1.18 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.20 DPS) [world]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 26.1 ranged_attack_power points (1.78 DPS) | yes | Pilferer's Gloves (7358, -0.51 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.59 DPS) [dungeon]; Insignia Gloves (6408, -0.59 DPS) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.1 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.14 DPS) [rep]; Highlander's Leather Girdle (20117, -0.14 DPS) [rep] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 40.5 ranged_attack_power points (2.76 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Footpads of the Fang (10411, -0.39 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) (or Protector's Band (19517)) | Blackfathom Deeps [dungeon] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Insurgent's Band (272067, -0.57 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Insurgent's Band (272067, -0.57 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) (or Sentinel's Blade (212583)) | Razorfen Kraul [dungeon] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Sentinel's Blade (212583) | Illiyana Moonblaze [vendor] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Jade Serpentblade (3850, +0.00 DPS) [crafted] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.84 DPS) | yes | Golemsight Long Gun (273029, -0.51 DPS, sim-verified) [dungeon]; Ironweaver (13137, -1.31 DPS) [world_drop]; Crystalpine Stinger (13037, -1.88 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Protector's Band; main_hand: Vendetta; off_hand: Sentinel's Blade; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 183.5. Weights run: 2.5s. Verify run: 2.5s. 626 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.808 ± 0.013, crit=0.752 ± 0.020 per rating point (14 rating = 1%, 10.534 per %), hit=1.307 ± 0.095 per rating point (10 rating = 1%, 13.070 per %), melee_haste=12.121 ± 1.404

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.5 ranged_attack_power points (2.40 DPS) | yes | Nightscape Headband (8176, -0.18 DPS) [crafted]; Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 30.9 ranged_attack_power points (2.03 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.24 DPS) [world]; Kaleidoscope Chain (13084, -1.29 DPS) [world_drop]; Erudite's Amulet (277204, -1.66 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.9 ranged_attack_power points (2.82 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.82 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.97 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 67.4 ranged_attack_power points (4.42 DPS) | yes | Wolffear Harness (13110, -1.02 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.66 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.66 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Branded Leather Bracers (19508, -0.16 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.18 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.10 DPS) | yes | Gloves of Holy Might (867, -0.10 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.26 DPS) [crafted]; Scarlet Gauntlets (10331, -0.26 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.97 DPS) | yes | Scorpashi Sash (14652, -0.13 DPS) [world_drop]; Blackforge Girdle (6425, -0.31 DPS) [dungeon]; Ogron's Sash (13117, -0.31 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 59.0 ranged_attack_power points (3.87 DPS) | yes | Triprunner Dungarees (9624, -0.55 DPS) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 36.5 ranged_attack_power points (2.40 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.1 ranged_attack_power points (1.84 DPS) | yes | Protector's Band (19515, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor]; Mark of Kern (2262, -0.53 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.3 ranged_attack_power points (1.66 DPS) | yes | Protector's Band (19515, -0.18 DPS) [rep]; Disengagement Ring (276202, -0.18 DPS) [vendor]; Mark of Kern (2262, -0.35 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (183.5 DPS) | yes | Hellslayer Battle Axe (13017, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.55 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (22.98 DPS) | yes | Swiftwind (13038, -2.16 DPS) [world_drop]; Baelog's Shortbow (9400, -2.20 DPS) [dungeon]; The Silencer (13138, -2.31 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 626, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 50 (dwarf, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 223.8. Weights run: 2.5s. Verify run: 3.0s. 782 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.784 ± 0.013, crit=0.920 ± 0.024 per rating point (14 rating = 1%, 12.877 per %), hit=1.562 ± 0.105 per rating point (10 rating = 1%, 15.618 per %), melee_haste=10.622 ± 1.555

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 57.4 ranged_attack_power points (3.90 DPS) | yes | Lordrec Helmet (10741, -0.87 DPS) [quest]; Sprightring Helm (17776, -1.06 DPS) [quest]; Helm of Fire (8348, -1.91 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.2 ranged_attack_power points (2.46 DPS) | yes | Sentinel's Medallion (19540, -0.38 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 49.1 ranged_attack_power points (3.33 DPS) | yes | Sunburn Spaulders (274751, -0.73 DPS, sim-verified) [vendor]; Khan's Mantle (14787, -1.06 DPS) [world_drop]; Nightscape Shoulders (8192, -1.25 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Serpentskin Cloak (8259, -0.76 DPS) [world_drop]; Nightscape Cloak (8195, -0.94 DPS) [crafted]; Blackflame Cape (13109, -1.73 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 66.8 ranged_attack_power points (4.54 DPS) | yes | Blazewind Breastplate (11193, -0.19 DPS) [quest]; Knight's Chain Armor (220828, -0.83 DPS) [vendor]; Charred Leather Tunic (19127, -1.13 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Wicked Leather Bracers (15084, -0.76 DPS) [crafted]; Arena Bands (18711, -0.93 DPS) [world]; Bloodlust Bracelets (14807, -1.95 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 61.2 ranged_attack_power points (4.16 DPS) | yes | Sergeant Major's Chain Gauntlets (220829, -1.33 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.51 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.51 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.8 ranged_attack_power points (3.72 DPS) | yes | Sagebrush Girdle (17778, -0.88 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.07 DPS) [crafted]; Stalker's Mail Belt (252588, -1.07 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.5 ranged_attack_power points (3.97 DPS) | yes | Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.45 DPS) [vendor]; Oilskin Leggings (9414, -0.57 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.7 ranged_attack_power points (3.78 DPS) | yes | Elven Chain Boots (13125, -0.38 DPS) [world_drop]; Fleetfoot Greaves (11627, -0.66 DPS, sim-verified) [dungeon]; Sergeant Major's Chain Sabatons (220837, -0.94 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.0 ranged_attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.94 DPS) [dungeon]; Protector's Band (19516, -0.94 DPS) [rep]; Disengagement Ring (276202, -1.13 DPS) [vendor] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.8 ranged_attack_power points (1.89 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Protector's Band (19516, -0.19 DPS) [rep]; Disengagement Ring (276202, -0.38 DPS) [vendor] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (223.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (223.8 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-verified (223.8 DPS) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Hanzo Sword (8190, -2.00 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 25.1 ranged_attack_power points (1.70 DPS) | yes | Widowmaker (4091, -1.04 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (223.8 DPS) | yes | Gryphonwing Long Bow (13022, -0.51 DPS) [dungeon]; Houndmaster's Bow (11628, -2.26 DPS) [dungeon]; Dark Iron Rifle (16004, -5.05 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Jhordy's Misplaced Screwdriver; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 782, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60 (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 296.7. Weights run: 2.5s. Verify run: 9.1s. 1415 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.830 ± 0.016, crit=1.293 ± 0.033 per rating point (14 rating = 1%, 18.095 per %), hit=2.300 ± 0.124 per rating point (10 rating = 1%, 23.004 per %), melee_haste=13.152 ± 1.790

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 88.1 ranged_attack_power points (6.00 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -10.61 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 53.8 ranged_attack_power points (3.66 DPS) | yes | Sentinel's Medallion (19538, -0.77 DPS) [rep]; Maelstrom's Tendril (19620, -0.77 DPS) [quest]; Beads of Ogre Might (22150, -1.15 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.5 ranged_attack_power points (6.44 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.5 ranged_attack_power points (4.25 DPS) | yes | Cloak of the Honor Guard (20073, -0.97 DPS) [rep]; Shifting Cloak (18511, -0.98 DPS) [crafted]; Howler's Furs (272414, -1.51 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 145.8 ranged_attack_power points (9.93 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.15 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.15 DPS) [vendor]; Obsidian Mail Tunic (22191, -5.34 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (296.7 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -7.78 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 76.4 ranged_attack_power points (5.21 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Marshal's Chain Grips (231560, -0.55 DPS) [pvp]; Mourning Handguards (279257, -1.16 DPS) [crafted] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (296.7 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, -3.96 DPS, sim-verified) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 158.3 ranged_attack_power points (10.78 DPS) | yes | Marshal's Chain Legguards (231577, -2.99 DPS) [pvp]; Sentinel's Leather Pants (237818, -3.11 DPS) [vendor]; Marshal's Chain Legplates (231558, -3.82 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (296.7 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beastmaster's Boots (22061, -8.88 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 90.9 ranged_attack_power points (6.19 DPS) | yes | Don Julio's Band (19325, -2.44 DPS) [rep]; Masons Fraternity Ring (9533, -3.50 DPS) [quest]; Protector's Band (19514, -4.07 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 62.6 ranged_attack_power points (4.27 DPS) | yes | Don Julio's Band (19325, -0.51 DPS) [rep]; Masons Fraternity Ring (9533, -1.57 DPS) [quest]; Protector's Band (19514, -2.15 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (296.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (296.7 DPS) | yes | Darkmoon Card: Heroism (19287, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (296.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -8.60 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (296.7 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -12.05 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1415, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 751.3. Weights run: 2.6s. Verify run: 9.3s. 1415 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.903 ± 0.016, crit=1.460 ± 0.031 per rating point (14 rating = 1%, 20.434 per %), hit=2.925 ± 0.181 per rating point (10 rating = 1%, 29.249 per %), melee_haste=19.008 ± 2.444

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 96.0 ranged_attack_power points (13.54 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -11.16 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 55.2 ranged_attack_power points (7.78 DPS) | yes | Beads of Ogre Might (22150, -0.27 DPS) [quest]; Medallion of the Dawn (22659, -1.51 DPS) [quest]; Sentinel's Medallion (19538, -1.64 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 98.8 ranged_attack_power points (13.94 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 63.5 ranged_attack_power points (8.96 DPS) | yes | Howler's Furs (272414, -0.89 DPS) [vendor]; Shifting Cloak (18511, -2.00 DPS) [crafted]; Cloak of the Honor Guard (20073, -2.12 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.1 ranged_attack_power points (22.16 DPS) | yes | Field Marshal's Chain Breastplate (16466, -5.35 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -5.35 DPS) [vendor]; Shadowcraft Tunic (16721, -10.96 DPS, sim-verified) [dungeon] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (751.3 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -4.68 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 78.4 ranged_attack_power points (11.05 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Marshal's Chain Grips (231560, -0.98 DPS) [pvp]; Mourning Handguards (279257, -2.46 DPS) [crafted] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 72.8 ranged_attack_power points (10.27 DPS) | yes | Ranger's Belt (272397, -0.44 DPS) [vendor]; Dense Timbermaw Belt (227807, -1.24 DPS) [vendor]; Substandard Belt Chain (274757, -2.29 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 171.7 ranged_attack_power points (24.22 DPS) | yes | Sentinel's Leather Pants (237818, -7.40 DPS) [vendor]; Marshal's Chain Legguards (231577, -7.42 DPS) [pvp]; Marshal's Chain Legplates (231558, -9.30 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (751.3 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -7.62 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 98.9 ranged_attack_power points (13.95 DPS) | yes | Don Julio's Band (19325, -4.97 DPS) [rep]; Masons Fraternity Ring (9533, -8.22 DPS) [quest]; Protector's Band (19514, -9.45 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 69.9 ranged_attack_power points (9.86 DPS) | yes | Don Julio's Band (19325, -0.88 DPS) [rep]; Masons Fraternity Ring (9533, -4.13 DPS) [quest]; Protector's Band (19514, -5.35 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (751.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (751.3 DPS) | yes | Darkmoon Card: Heroism (19287, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (751.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -20.02 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (751.3 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -23.12 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1415, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 111.5. Weights run: 2.2s. Verify run: 1.8s. 206 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.112 ± 0.015, crit=0.605 ± 0.016 per rating point (14 rating = 1%, 8.470 per %), hit=1.109 ± 0.048 per rating point (10 rating = 1%, 11.089 per %), melee_haste=7.909 ± 0.726

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.9 ranged_attack_power points (1.72 DPS) | yes | Red Winter Hat (21524, -1.79 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.7 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.91 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.6 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.7 ranged_attack_power points (1.29 DPS) | yes | Cape of the Brotherhood (5193, -0.22 DPS) [dungeon]; Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 21.8 ranged_attack_power points (1.51 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.22 DPS) [crafted]; Prospector's Chestpiece (14562, -0.22 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.6 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bristlebark Bindings (14569, -0.43 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.43 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -1.43 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Guardsman Belt (3429, -0.38 DPS) [world]; Dark Leather Belt (4249, -0.38 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 28.0 ranged_attack_power points (1.94 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Gray Leather Boots (2315, -0.22 DPS) [crafted]; Feet of the Lynx (1121, -1.09 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.4 ranged_attack_power points (0.86 DPS) | yes | The 1 Ring (8350, -0.65 DPS) [world]; Bounty Hunter's Ring (5351, -0.91 DPS, sim-verified) [quest] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | The 1 Ring (8350, -0.22 DPS) [world]; Bounty Hunter's Ring (5351, -1.14 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.9 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Crescent Staff (6505, -0.63 DPS) [quest]; Lupine Axe (1220, -0.85 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.42 DPS) | yes | Outrider's Bow (20437, -0.68 DPS) [pvp]; Lil Timmy's Peashooter (13136, -1.09 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.00 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Legionnaire's Band; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 206, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 209611 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 142.8. Weights run: 2.3s. Verify run: 2.0s. 380 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.894 ± 0.013, crit=0.674 ± 0.017 per rating point (14 rating = 1%, 9.436 per %), hit=1.207 ± 0.058 per rating point (10 rating = 1%, 12.067 per %), melee_haste=11.026 ± 0.959

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.97 DPS) | yes | Brawler's Leather Hood (252504, -0.43 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world]; Resilient Cap (14401, -1.58 DPS) [world_drop] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.76 DPS) [world]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Erudite's Amulet (277204, -1.18 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.17 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.39 DPS) [dungeon]; Swiftrunner Cape (6745, -0.39 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.76 DPS) | yes | Garb of Florid Feathers (273026, +0.00 DPS, sim-verified) [dungeon]; Panther Armor (6670, -0.99 DPS) [quest]; Green Leather Armor (4255, -1.18 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop]; Loamflake Bracers (15462, -0.20 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 26.1 ranged_attack_power points (1.78 DPS) | yes | Pilferer's Gloves (7358, -0.20 DPS) [crafted]; Braced Handguards (6784, -0.39 DPS) [quest]; Insignia Gloves (6408, -0.59 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (142.8 DPS) | yes | Defiler's Chain Girdle (20152, -0.14 DPS) [rep]; Defiler's Leather Girdle (20191, -0.14 DPS) [rep]; Skulker's Leather Belt (252520, -1.97 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 40.5 ranged_attack_power points (2.76 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ring of Precision (1491) (or Legionnaire's Band (19513)) | Blackfathom Deeps [dungeon] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Band of the Fist (17694, -0.39 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Band of the Fist (17694, -0.39 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) (or Sentinel's Blade (212583)) | Razorfen Kraul [dungeon] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Sentinel's Blade (212583) | Illiyana Moonblaze [vendor] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Outlaw Sabre (16886, +0.00 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.84 DPS) | yes | Golemsight Long Gun (273029, -1.28 DPS, sim-verified) [dungeon]; Ironweaver (13137, -1.31 DPS) [world_drop]; Outrider's Bow (19561, -1.53 DPS) [pvp] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Legionnaire's Band; main_hand: Vendetta; off_hand: Sentinel's Blade; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 185.7. Weights run: 2.5s. Verify run: 2.4s. 590 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.808 ± 0.013, crit=0.752 ± 0.020 per rating point (14 rating = 1%, 10.534 per %), hit=1.307 ± 0.095 per rating point (10 rating = 1%, 13.070 per %), melee_haste=12.121 ± 1.404

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.5 ranged_attack_power points (2.40 DPS) | yes | Nightscape Headband (8176, -0.18 DPS) [crafted]; Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 30.9 ranged_attack_power points (2.03 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.24 DPS) [world]; Ethereal Talisman (4430, -1.29 DPS) [quest]; Kaleidoscope Chain (13084, -1.29 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.9 ranged_attack_power points (2.82 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.84 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.97 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 67.4 ranged_attack_power points (4.42 DPS) | yes | Wolffear Harness (13110, -1.03 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.66 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.66 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.10 DPS) | yes | Gloves of Holy Might (867, -0.10 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.26 DPS) [crafted]; Scarlet Gauntlets (10331, -0.26 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.97 DPS) | yes | Scorpashi Sash (14652, -0.13 DPS) [world_drop]; Blackforge Girdle (6425, -0.31 DPS) [dungeon]; Ogron's Sash (13117, -0.31 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 59.0 ranged_attack_power points (3.87 DPS) | yes | Triprunner Dungarees (9624, -0.65 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 36.5 ranged_attack_power points (2.40 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.1 ranged_attack_power points (1.84 DPS) | yes | Legionnaire's Band (19512, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor]; Mark of Kern (2262, -0.53 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.3 ranged_attack_power points (1.66 DPS) | yes | Legionnaire's Band (19512, -0.18 DPS) [rep]; Disengagement Ring (276202, -0.18 DPS) [vendor]; Mark of Kern (2262, -0.35 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (185.7 DPS) | yes | Hellslayer Battle Axe (13017, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.84 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (22.98 DPS) | yes | Outrider's Bow (19560, -1.55 DPS) [pvp]; Swiftwind (13038, -2.16 DPS) [world_drop]; The Silencer (13138, -3.17 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 590, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 227.1. Weights run: 2.5s. Verify run: 2.8s. 739 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.784 ± 0.013, crit=0.920 ± 0.024 per rating point (14 rating = 1%, 12.877 per %), hit=1.562 ± 0.105 per rating point (10 rating = 1%, 15.618 per %), melee_haste=10.622 ± 1.555

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 47.3 ranged_attack_power points (3.21 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.38 DPS) [quest]; Tough Scorpid Helm (8208, -0.57 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.2 ranged_attack_power points (2.46 DPS) | yes | Scout's Medallion (19535, -0.19 DPS) [rep]; Woven Ivy Necklace (19159, -0.76 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.6 ranged_attack_power points (2.89 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Khan's Mantle (14787, -0.63 DPS) [world_drop]; Nightscape Shoulders (8192, -0.81 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Blackflame Cape (13109, -0.57 DPS) [world_drop]; Serpentskin Cloak (8259, -0.76 DPS) [world_drop]; Nightscape Cloak (8195, -0.94 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 66.8 ranged_attack_power points (4.54 DPS) | yes | Blazewind Breastplate (11193, -0.19 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.83 DPS) [vendor]; Charred Leather Tunic (19127, -1.13 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Bloodlust Bracelets (14807, -0.83 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 61.2 ranged_attack_power points (4.16 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -1.33 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.51 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.51 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.8 ranged_attack_power points (3.72 DPS) | yes | Sagebrush Girdle (17778, -0.88 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.07 DPS) [crafted]; Stalker's Mail Belt (252588, -1.07 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.5 ranged_attack_power points (3.97 DPS) | yes | Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Stone Guard's Chain Legplates (220833, -0.45 DPS) [vendor]; Oilskin Leggings (9414, -0.57 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.7 ranged_attack_power points (3.78 DPS) | yes | Fleetfoot Greaves (11627, -0.19 DPS) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; First Sergeant's Chain Sabatons (220836, -0.94 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.0 ranged_attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.94 DPS) [dungeon]; Legionnaire's Band (19511, -0.94 DPS) [rep]; White Bone Band (11862, -1.02 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.8 ranged_attack_power points (1.89 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Legionnaire's Band (19511, -0.19 DPS) [rep]; White Bone Band (11862, -0.26 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (227.1 DPS) | yes | Molten Heart of the Mountain (249470, -7.25 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (227.1 DPS) | yes | Molten Heart of the Mountain (249470, -0.85 DPS, sim-verified) [crafted] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-verified (227.1 DPS) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Force of the Hippogryph (9684) (or Jhordy's Misplaced Screwdriver (274753)) | Weapons of Spirit [quest] | 25.1 ranged_attack_power points (1.70 DPS) | yes | Jhordy's Misplaced Screwdriver (274753, +0.00 DPS, sim-verified) [vendor]; White Bone Shredder (11863, -0.38 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (227.1 DPS) | yes | Gryphonwing Long Bow (13022, -0.51 DPS) [dungeon]; Outrider's Bow (19559, -2.24 DPS) [pvp]; Dark Iron Rifle (16004, -3.40 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Force of the Hippogryph; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 739, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 307.0. Weights run: 2.5s. Verify run: 8.8s. 1395 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.830 ± 0.016, crit=1.293 ± 0.033 per rating point (14 rating = 1%, 18.095 per %), hit=2.300 ± 0.124 per rating point (10 rating = 1%, 23.004 per %), melee_haste=13.152 ± 1.790

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 88.1 ranged_attack_power points (6.00 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -10.77 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 53.8 ranged_attack_power points (3.66 DPS) | yes | Beads of Ogre Might (22150, -0.46 DPS) [quest]; Scout's Medallion (19534, -0.77 DPS) [rep]; Maelstrom's Tendril (19620, -0.77 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.5 ranged_attack_power points (6.44 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -0.46 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.5 ranged_attack_power points (4.25 DPS) | yes | Howler's Furs (272414, -0.78 DPS) [vendor]; Deathguard's Cloak (20068, -0.97 DPS) [rep]; Shifting Cloak (18511, -0.98 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 145.8 ranged_attack_power points (9.93 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.15 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.15 DPS) [vendor]; Obsidian Mail Tunic (22191, -4.29 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (307.0 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -8.25 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 76.4 ranged_attack_power points (5.21 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; General's Chain Grips (231569, -0.55 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (307.0 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, -5.66 DPS, sim-verified) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 158.3 ranged_attack_power points (10.78 DPS) | yes | Outrider's Chain Leggings (22673, -1.97 DPS, sim-verified) [rep]; General's Chain Legguards (231574, -2.99 DPS) [pvp]; Sentinel's Leather Pants (237818, -3.11 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (307.0 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beastmaster's Boots (22061, -8.53 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 90.9 ranged_attack_power points (6.19 DPS) | yes | Don Julio's Band (19325, -2.44 DPS) [rep]; Masons Fraternity Ring (9533, -3.50 DPS) [quest]; Legionnaire's Band (19510, -4.07 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 62.6 ranged_attack_power points (4.27 DPS) | yes | Don Julio's Band (19325, -0.51 DPS) [rep]; Masons Fraternity Ring (9533, -1.57 DPS) [quest]; Legionnaire's Band (19510, -2.15 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (307.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (307.0 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (307.0 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -7.51 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (307.0 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.05 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1395, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 781.9. Weights run: 2.6s. Verify run: 9.0s. 1395 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.903 ± 0.016, crit=1.460 ± 0.031 per rating point (14 rating = 1%, 20.434 per %), hit=2.925 ± 0.181 per rating point (10 rating = 1%, 29.249 per %), melee_haste=19.008 ± 2.444

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 96.0 ranged_attack_power points (13.54 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -18.16 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 55.2 ranged_attack_power points (7.78 DPS) | yes | Beads of Ogre Might (22150, -0.27 DPS) [quest]; Medallion of the Dawn (22659, -1.51 DPS) [quest]; Scout's Medallion (19534, -1.64 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 98.8 ranged_attack_power points (13.94 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -0.54 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 63.5 ranged_attack_power points (8.96 DPS) | yes | Howler's Furs (272414, -0.89 DPS) [vendor]; Shifting Cloak (18511, -2.00 DPS) [crafted]; Deathguard's Cloak (20068, -2.12 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.1 ranged_attack_power points (22.16 DPS) | yes | Warlord's Chain Chestpiece (16565, -5.35 DPS) [vendor]; Warlord's Chain Hauberk (231573, -5.35 DPS) [vendor]; Shadowcraft Tunic (16721, -10.16 DPS, sim-verified) [dungeon] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (781.9 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -14.20 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 78.4 ranged_attack_power points (11.05 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; General's Chain Grips (231569, -0.98 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (781.9 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -10.76 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 171.7 ranged_attack_power points (24.22 DPS) | yes | Outrider's Chain Leggings (22673, -3.80 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -7.40 DPS) [vendor]; General's Chain Legguards (231574, -7.42 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (781.9 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -15.51 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 98.9 ranged_attack_power points (13.95 DPS) | yes | Don Julio's Band (19325, -4.97 DPS) [rep]; Masons Fraternity Ring (9533, -8.22 DPS) [quest]; Legionnaire's Band (19510, -9.45 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 69.9 ranged_attack_power points (9.86 DPS) | yes | Don Julio's Band (19325, -0.88 DPS) [rep]; Masons Fraternity Ring (9533, -4.13 DPS) [quest]; Legionnaire's Band (19510, -5.35 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (781.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (781.9 DPS) | yes | Molten Heart of the Mountain (249470, -7.03 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (781.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -19.52 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (781.9 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Bow of Searing Arrows (2825, -28.18 DPS, sim-verified) [world_drop] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1395, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

