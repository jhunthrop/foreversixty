# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 110.3. Weights run: 1.5s. Verify run: 2.6s. 282 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=3.098 ± 0.021, crit=0.580 ± 0.022 per rating point (14 rating = 1%, 8.118 per %), hit=0.950 ± 0.069 per rating point (10 rating = 1%, 9.503 per %), melee_haste=8.075 ± 1.136

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.8 ranged_attack_power points (1.72 DPS) | yes | Red Winter Hat (21524, -1.84 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.93 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -1.13 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Cape of the Brotherhood (5193, -0.22 DPS) [dungeon]; Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 34.1 ranged_attack_power points (2.37 DPS) | yes | Trapper's Leather Armor (252491, -0.86 DPS) [crafted]; Brawler's Leather Armor (252490, -0.91 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -1.08 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bravo's Armbands (270015, -0.22 DPS) [quest]; Ratchet Wristwraps (274742, -0.43 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (110.3 DPS) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -1.28 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Dusty Belt (279897, -0.17 DPS) [quest]; Dark Leather Belt (4249, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.9 ranged_attack_power points (1.94 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.8 ranged_attack_power points (1.72 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.43 DPS) [dungeon]; Agile Boots (4788, -0.65 DPS) [vendor] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.4 ranged_attack_power points (0.86 DPS) | yes | The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.2 ranged_attack_power points (0.43 DPS) | yes | The 1 Ring (8350, -0.92 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.8 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, -0.20 DPS) [dungeon]; Lupine Axe (1220, -0.85 DPS) [world]; Piercing Axe (6094, -0.85 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.46 DPS) | yes | Lil Timmy's Peashooter (13136, -1.21 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.74 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -3.01 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 282, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw; 2241 Desperado Cape; 3000 Brood Mother Carapace; 3018 Hide of Lupos

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 140.6. Weights run: 1.6s. Verify run: 3.1s. 473 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.885 ± 0.018, crit=0.658 ± 0.024 per rating point (14 rating = 1%, 9.208 per %), hit=1.171 ± 0.078 per rating point (10 rating = 1%, 11.705 per %), melee_haste=10.464 ± 1.273

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Brawler's Leather Hood (252504, -0.40 DPS) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world]; Resilient Cap (14401, -1.58 DPS) [world_drop] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Erudite's Amulet (277204, -1.19 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.7 ranged_attack_power points (2.18 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.40 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.40 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.4 ranged_attack_power points (2.77 DPS) | yes | Tunic of Westfall (2041, -0.65 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.19 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.20 DPS) [world]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Serpent Gloves (5970, -0.40 DPS) [dungeon]; Insignia Gloves (6408, -0.40 DPS) [world_drop]; Gloves of the Fang (10413, -0.40 DPS) [dungeon] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (140.6 DPS) | yes | Highlander's Chain Girdle (20090, -0.13 DPS) [rep]; Highlander's Leather Girdle (20117, -0.13 DPS) [rep]; Skulker's Leather Belt (252520, -1.54 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 40.4 ranged_attack_power points (2.77 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Ferine Leggings (6690, -0.99 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.20 DPS) [quest] |
| finger1 | Monkey Ring (6748) | Willix the Importer [quest] | 20.2 ranged_attack_power points (1.39 DPS) | yes | Protector's Band (19517, -0.20 DPS) [rep]; Defias Renegade Ring (1076, -0.40 DPS) [dungeon] |
| finger2 | Ring of Precision (1491) (or Protector's Band (19517)) | Blackfathom Deeps [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Protector's Band (19517, +0.00 DPS) [rep]; Defias Renegade Ring (1076, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) (or Sentinel's Blade (212583)) | Razorfen Kraul [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Sentinel's Blade (212583) | Illiyana Moonblaze [vendor] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Jade Serpentblade (3850, +0.00 DPS) [crafted] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.94 DPS) | yes | Golemsight Long Gun (273029, -0.81 DPS) [dungeon]; Silver Star (3463, -1.17 DPS, sim-verified) [quest]; Ironweaver (13137, -1.32 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Monkey Ring; finger2: Ring of Precision; main_hand: Vendetta; off_hand: Sentinel's Blade; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 473, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw; 2241 Desperado Cape; 2944 Cursed Eye of Paleth

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 182.6. Weights run: 1.5s. Verify run: 3.4s. 712 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.803 ± 0.018, crit=0.723 ± 0.029 per rating point (14 rating = 1%, 10.121 per %), hit=1.458 ± 0.135 per rating point (10 rating = 1%, 14.580 per %), melee_haste=not significant (7.159 ± 2.067)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Nightscape Headband (8176, -0.18 DPS) [crafted]; Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 30.8 ranged_attack_power points (2.03 DPS) | yes | Kaleidoscope Chain (13084, -1.29 DPS) [world_drop]; Erudite's Amulet (277204, -1.66 DPS) [quest]; Amberglow Talisman (10824, -1.85 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.8 ranged_attack_power points (2.83 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.83 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 47.7 ranged_attack_power points (3.14 DPS) | yes | Tough Scorpid Breastplate (8203, +0.00 DPS, sim-verified) [crafted]; Nightscape Tunic (8175, -0.37 DPS) [crafted]; Dusky Leather Armor (7374, -0.55 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Branded Leather Bracers (19508, -0.16 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.18 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.11 DPS) | yes | Gloves of Holy Might (867, -0.12 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.26 DPS) [crafted]; Scarlet Gauntlets (10331, -0.26 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.98 DPS) | yes | Scorpashi Sash (14652, -0.13 DPS) [world_drop]; Blackforge Girdle (6425, -0.31 DPS) [dungeon]; Ogron's Sash (13117, -0.31 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.9 ranged_attack_power points (3.88 DPS) | yes | Triprunner Dungarees (9624, -0.55 DPS) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.37 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.0 ranged_attack_power points (1.85 DPS) | yes | Protector's Band (19515, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor]; Mark of Kern (2262, -0.53 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.2 ranged_attack_power points (1.66 DPS) | yes | Protector's Band (19515, -0.18 DPS) [rep]; Disengagement Ring (276202, -0.18 DPS) [vendor]; Mark of Kern (2262, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (182.6 DPS) | yes | Hellslayer Battle Axe (13017, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -0.76 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (23.09 DPS) | yes | Swiftwind (13038, -2.17 DPS) [world_drop]; Monolithic Bow (9426, -2.35 DPS) [dungeon]; The Silencer (13138, -2.82 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Parachute Cloak; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 712, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1678 Black Ogre Kickers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw; 2241 Desperado Cape

### Band 50 (dwarf, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 222.5. Weights run: 1.5s. Verify run: 4.0s. 867 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.767 ± 0.018, crit=0.899 ± 0.034 per rating point (14 rating = 1%, 12.581 per %), hit=1.475 ± 0.128 per rating point (10 rating = 1%, 14.746 per %), melee_haste=11.870 ± 2.100

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 56.9 ranged_attack_power points (3.89 DPS) | yes | Lordrec Helmet (10741, -0.86 DPS) [quest]; Sprightring Helm (17776, -1.05 DPS) [quest]; Helm of Fire (8348, -2.82 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.0 ranged_attack_power points (2.46 DPS) | yes | Sentinel's Medallion (19539, -0.19 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 48.6 ranged_attack_power points (3.32 DPS) | yes | Sunburn Spaulders (274751, -0.77 DPS, sim-verified) [vendor]; Khan's Mantle (14787, -1.05 DPS) [world_drop]; Nightscape Shoulders (8192, -1.24 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 41.5 ranged_attack_power points (2.84 DPS) | yes | Blackflame Cape (13109, -0.57 DPS) [world_drop]; Serpentskin Cloak (8259, -0.76 DPS) [world_drop]; Nightscape Cloak (8195, -0.95 DPS) [crafted] |
| chest | Blazewind Breastplate (11193) | Tremors of the Earth [quest] | 63.6 ranged_attack_power points (4.35 DPS) | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Warbear Harness (15064, -0.95 DPS) [crafted]; Charred Leather Tunic (19127, -0.95 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.5 ranged_attack_power points (2.84 DPS) | yes | Bloodlust Bracelets (14807, -0.76 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.76 DPS) [crafted]; Arena Bands (18711, -0.92 DPS) [world] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 60.9 ranged_attack_power points (4.16 DPS) | yes | Sergeant Major's Chain Gauntlets (220829, -1.30 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.51 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.51 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.5 ranged_attack_power points (3.73 DPS) | yes | Sagebrush Girdle (17778, -0.89 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.08 DPS) [crafted]; Stalker's Mail Belt (252588, -1.08 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.1 ranged_attack_power points (3.97 DPS) | yes | Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.46 DPS) [vendor]; Oilskin Leggings (9414, -0.57 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.3 ranged_attack_power points (3.78 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; Sergeant Major's Chain Sabatons (220837, -0.95 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.7 ranged_attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.95 DPS) [dungeon]; Protector's Band (19516, -0.95 DPS) [rep]; Disengagement Ring (276202, -1.14 DPS) [vendor] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.7 ranged_attack_power points (1.89 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Protector's Band (19516, -0.19 DPS) [rep]; Disengagement Ring (276202, -0.38 DPS) [vendor] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (222.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (222.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-verified (222.5 DPS) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Hanzo Sword (8190, -1.91 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 24.9 ranged_attack_power points (1.70 DPS) | yes | Widowmaker (4091, +0.00 DPS) [world_drop] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (222.5 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.54 DPS) [world_drop]; Gryphonwing Long Bow (13022, -2.05 DPS) [dungeon]; Dark Iron Rifle (16004, -2.79 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Blazewind Breastplate; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Jhordy's Misplaced Screwdriver; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 867, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1678 Black Ogre Kickers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw; 2241 Desperado Cape

### Band 60 (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 293.4. Weights run: 1.2s. Verify run: 12.0s. 1495 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.832 ± 0.023, crit=1.308 ± 0.048 per rating point (14 rating = 1%, 18.314 per %), hit=2.044 ± 0.181 per rating point (10 rating = 1%, 20.439 per %), melee_haste=15.922 ± 2.619

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (293.4 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -9.18 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 53.8 ranged_attack_power points (3.68 DPS) | yes | Beads of Ogre Might (22150, -0.64 DPS) [quest]; Sentinel's Medallion (19538, -0.77 DPS) [rep]; Maelstrom's Tendril (19620, -0.77 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.8 ranged_attack_power points (6.48 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Field Marshal's Chain Spaulders (16468, -0.05 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -0.05 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.5 ranged_attack_power points (4.27 DPS) | yes | Howler's Furs (272414, -0.96 DPS) [vendor]; Cloak of the Honor Guard (20073, -0.98 DPS) [rep]; Shifting Cloak (18511, -0.98 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 143.5 ranged_attack_power points (9.81 DPS) | yes | Field Marshal's Chain Breastplate (16466, -1.98 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -1.98 DPS) [vendor]; Obsidian Mail Tunic (22191, -4.69 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (293.4 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -8.01 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (293.4 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (293.4 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -7.70 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 156.2 ranged_attack_power points (10.68 DPS) | yes | Marshal's Chain Legguards (231577, -2.84 DPS) [pvp]; Sentinel's Leather Pants (237818, -2.95 DPS) [vendor]; Marshal's Chain Legplates (231558, -3.67 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (293.4 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -8.07 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 88.4 ranged_attack_power points (6.04 DPS) | yes | Don Julio's Band (19325, -2.30 DPS) [rep]; Masons Fraternity Ring (9533, -3.33 DPS) [quest]; Protector's Band (19514, -3.91 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 60.1 ranged_attack_power points (4.11 DPS) | yes | Don Julio's Band (19325, -0.36 DPS) [rep]; Masons Fraternity Ring (9533, -1.40 DPS) [quest]; Protector's Band (19514, -1.98 DPS) [rep] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (293.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-verified (293.4 DPS) | yes | Darkmoon Card: Blue Dragon (19288, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (293.4 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -8.23 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (293.4 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -11.48 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket2: Darkmoon Card: Heroism; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1495, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1678 Black Ogre Kickers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw; 2241 Desperado Cape

### Band 60, raid preset (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 763.3. Weights run: 1.4s. Verify run: 12.0s. 1495 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.875 ± 0.022, crit=1.352 ± 0.044 per rating point (14 rating = 1%, 18.933 per %), hit=3.244 ± 0.247 per rating point (10 rating = 1%, 32.442 per %), melee_haste=16.861 ± 3.307

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 98.6 ranged_attack_power points (14.03 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -16.33 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 56.4 ranged_attack_power points (8.03 DPS) | yes | Sentinel's Medallion (19538, -1.89 DPS) [rep]; Maelstrom's Tendril (19620, -1.89 DPS) [quest]; Amulet of the Darkmoon (19491, -2.10 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 96.6 ranged_attack_power points (13.74 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 63.1 ranged_attack_power points (8.98 DPS) | yes | Howler's Furs (272414, -0.38 DPS) [vendor]; Shifting Cloak (18511, -2.03 DPS) [crafted]; Cloak of the Honor Guard (20073, -2.10 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.8 ranged_attack_power points (22.45 DPS) | yes | Field Marshal's Chain Breastplate (16466, -5.84 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -5.84 DPS) [vendor]; Shadowcraft Tunic (16721, -10.06 DPS, sim-verified) [dungeon] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (763.3 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -10.53 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 77.6 ranged_attack_power points (11.05 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Marshal's Chain Grips (231560, -1.13 DPS) [pvp]; Mourning Handguards (279257, -2.45 DPS) [crafted] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (763.3 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -8.30 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 170.9 ranged_attack_power points (24.32 DPS) | yes | Marshal's Chain Legguards (231577, -7.72 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.89 DPS) [vendor]; Marshal's Chain Legplates (231558, -9.55 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (763.3 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -13.53 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 101.4 ranged_attack_power points (14.43 DPS) | yes | Don Julio's Band (19325, -4.85 DPS) [rep]; Masons Fraternity Ring (9533, -8.71 DPS) [quest]; Channeler's Ring (272406, -9.82 DPS) [vendor] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 72.7 ranged_attack_power points (10.34 DPS) | yes | Don Julio's Band (19325, -0.76 DPS) [rep]; Masons Fraternity Ring (9533, -4.62 DPS) [quest]; Channeler's Ring (272406, -5.73 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (763.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (763.3 DPS) | yes | Darkmoon Card: Heroism (19287, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (763.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -19.48 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (763.3 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Bow of Searing Arrows (2825, -29.82 DPS, sim-verified) [world_drop] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1495, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1678 Black Ogre Kickers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw; 2241 Desperado Cape

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 110.3. Weights run: 1.5s. Verify run: 2.7s. 271 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=3.098 ± 0.021, crit=0.580 ± 0.022 per rating point (14 rating = 1%, 8.118 per %), hit=0.950 ± 0.069 per rating point (10 rating = 1%, 9.503 per %), melee_haste=8.075 ± 1.136

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.8 ranged_attack_power points (1.72 DPS) | yes | Red Winter Hat (21524, -1.85 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Erudite's Amulet (277204, -0.96 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -0.37 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.6 ranged_attack_power points (1.29 DPS) | yes | Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.48 DPS, sim-verified) [dungeon] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 21.7 ranged_attack_power points (1.51 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Dark Leather Tunic (2317, -0.22 DPS) [crafted]; Prospector's Chestpiece (14562, -0.22 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.5 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bristlebark Bindings (14569, -0.43 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.43 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (110.3 DPS) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -1.14 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Guardsman Belt (3429, -0.39 DPS) [world]; Dark Leather Belt (4249, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.9 ranged_attack_power points (1.94 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.8 ranged_attack_power points (1.72 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.43 DPS) [dungeon]; Agile Boots (4788, -0.65 DPS) [vendor] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.4 ranged_attack_power points (0.86 DPS) | yes | Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 9.3 ranged_attack_power points (0.65 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.43 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.8 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Crescent Staff (6505, -0.63 DPS) [quest]; Lupine Axe (1220, -0.85 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.46 DPS) | yes | Outrider's Bow (20437, -0.68 DPS) [pvp]; Lil Timmy's Peashooter (13136, -1.47 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.74 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Bounty Hunter's Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 271, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw; 2241 Desperado Cape; 3000 Brood Mother Carapace

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 142.2. Weights run: 1.6s. Verify run: 2.9s. 458 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.885 ± 0.018, crit=0.658 ± 0.024 per rating point (14 rating = 1%, 9.208 per %), hit=1.171 ± 0.078 per rating point (10 rating = 1%, 11.705 per %), melee_haste=10.464 ± 1.273

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Brawler's Leather Hood (252504, -0.45 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world]; Resilient Cap (14401, -1.58 DPS) [world_drop] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Erudite's Amulet (277204, -1.19 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.7 ranged_attack_power points (2.18 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.40 DPS) [dungeon]; Swiftrunner Cape (6745, -0.40 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.4 ranged_attack_power points (2.77 DPS) | yes | Panther Armor (6670, -1.09 DPS, sim-verified) [quest]; Green Leather Armor (4255, -1.19 DPS) [crafted]; Brawler's Leather Tunic (252508, -1.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop]; Loamflake Bracers (15462, -0.20 DPS) [quest] |
| hands | Pilferer's Gloves (7358) | Leatherworking [crafted] | 23.1 ranged_attack_power points (1.58 DPS) | yes | Braced Handguards (6784, -0.20 DPS) [quest]; Serpent Gloves (5970, -0.40 DPS) [dungeon]; Insignia Gloves (6408, -0.40 DPS) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.0 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.13 DPS) [rep]; Defiler's Leather Girdle (20191, -0.13 DPS) [rep] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 40.4 ranged_attack_power points (2.77 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Ferine Leggings (6690, -0.99 DPS) [dungeon] |
| feet | Vorrel's Boots (7751) | Vorrel's Revenge [quest] | 28.9 ranged_attack_power points (1.98 DPS) | yes | Warsong Boots (16977, -0.40 DPS) [quest]; Highlander's Mail Greaves (20123, -0.40 DPS) [vendor]; Insignia Boots (4055, -0.69 DPS, sim-verified) [world_drop] |
| finger1 | Monkey Ring (6748) | Willix the Importer [quest] | 20.2 ranged_attack_power points (1.39 DPS) | yes | Legionnaire's Band (19513, -0.20 DPS) [rep]; Defias Renegade Ring (1076, -0.40 DPS) [dungeon]; Band of the Fist (17694, -0.59 DPS) [quest] |
| finger2 | Ring of Precision (1491) (or Legionnaire's Band (19513)) | Blackfathom Deeps [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Legionnaire's Band (19513, +0.00 DPS) [rep]; Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Band of the Fist (17694, -0.40 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | 24.4 ranged_attack_power points (1.68 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Sentinel's Blade (212583)) | Razorfen Kraul [dungeon] | 17.3 ranged_attack_power points (1.19 DPS) | yes | Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.94 DPS) | yes | Silver Star (3463, -0.74 DPS, sim-verified) [quest]; Golemsight Long Gun (273029, -0.81 DPS) [dungeon]; Ironweaver (13137, -1.32 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Pilferer's Gloves; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Vorrel's Boots; finger1: Monkey Ring; finger2: Ring of Precision; main_hand: Serrated Raptor Claw; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 458, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw; 2241 Desperado Cape

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 184.6. Weights run: 1.5s. Verify run: 3.5s. 678 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.803 ± 0.018, crit=0.723 ± 0.029 per rating point (14 rating = 1%, 10.121 per %), hit=1.458 ± 0.135 per rating point (10 rating = 1%, 14.580 per %), melee_haste=not significant (7.159 ± 2.067)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Nightscape Headband (8176, -0.18 DPS) [crafted]; Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 30.8 ranged_attack_power points (2.03 DPS) | yes | Ethereal Talisman (4430, -1.29 DPS) [quest]; Kaleidoscope Chain (13084, -1.29 DPS) [world_drop]; Erudite's Amulet (277204, -1.66 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.8 ranged_attack_power points (2.83 DPS) | yes | Forest Tracker Epaulets (2278, -0.79 DPS) [world_drop]; Nightscape Shoulders (8192, -0.84 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.98 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 47.7 ranged_attack_power points (3.14 DPS) | yes | Nightscape Tunic (8175, -0.37 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.37 DPS) [crafted]; Dusky Leather Armor (7374, -0.55 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.4 ranged_attack_power points (1.48 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.11 DPS) | yes | Gloves of Holy Might (867, -0.12 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.26 DPS) [crafted]; Scarlet Gauntlets (10331, -0.26 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.98 DPS) | yes | Scorpashi Sash (14652, -0.13 DPS) [world_drop]; Blackforge Girdle (6425, -0.31 DPS) [dungeon]; Ogron's Sash (13117, -0.31 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.9 ranged_attack_power points (3.88 DPS) | yes | Triprunner Dungarees (9624, -0.60 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 36.4 ranged_attack_power points (2.40 DPS) | yes | Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.37 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.0 ranged_attack_power points (1.85 DPS) | yes | Legionnaire's Band (19512, -0.37 DPS) [rep]; Disengagement Ring (276202, -0.37 DPS) [vendor]; Mark of Kern (2262, -0.53 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.2 ranged_attack_power points (1.66 DPS) | yes | Disengagement Ring (276202, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Band (19512, -0.18 DPS) [rep]; Mark of Kern (2262, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (184.6 DPS) | yes | Hellslayer Battle Axe (13017, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.23 DPS, sim-verified) [world_drop] |
| off_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | 24.0 ranged_attack_power points (1.58 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (23.09 DPS) | yes | Outrider's Bow (19560, -1.56 DPS) [pvp]; Swiftwind (13038, -2.17 DPS) [world_drop]; The Silencer (13138, -4.11 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Parachute Cloak; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Serrated Raptor Claw; ranged: Bow of Searing Arrows

No-known-source sample (15 of 678, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1678 Black Ogre Kickers; 1832 Lucky Trousers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw

### Band 50 (troll, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 229.5. Weights run: 1.5s. Verify run: 3.6s. 826 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.767 ± 0.018, crit=0.899 ± 0.034 per rating point (14 rating = 1%, 12.581 per %), hit=1.475 ± 0.128 per rating point (10 rating = 1%, 14.746 per %), melee_haste=11.870 ± 2.100

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 47.0 ranged_attack_power points (3.22 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.38 DPS) [quest]; Tough Scorpid Helm (8208, -0.57 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.0 ranged_attack_power points (2.46 DPS) | yes | Scout's Medallion (19535, -0.19 DPS) [rep]; Woven Ivy Necklace (19159, -0.76 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 42.4 ranged_attack_power points (2.90 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Khan's Mantle (14787, -0.63 DPS) [world_drop]; Nightscape Shoulders (8192, -0.82 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 41.5 ranged_attack_power points (2.84 DPS) | yes | Serpentskin Cloak (8259, -0.76 DPS) [world_drop]; Nightscape Cloak (8195, -0.95 DPS) [crafted]; Blackflame Cape (13109, -1.26 DPS, sim-verified) [world_drop] |
| chest | Blazewind Breastplate (11193) | Broken Alliances [quest] | 63.6 ranged_attack_power points (4.35 DPS) | yes | Stone Guard's Chain Armor (220827, +0.00 DPS, sim-verified) [vendor]; Warbear Harness (15064, -0.95 DPS) [crafted]; Charred Leather Tunic (19127, -0.95 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.5 ranged_attack_power points (2.84 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Bloodlust Bracelets (14807, -0.88 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 60.9 ranged_attack_power points (4.16 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -1.28 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.51 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.51 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 54.5 ranged_attack_power points (3.73 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.08 DPS) [crafted]; Stalker's Mail Belt (252588, -1.08 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.1 ranged_attack_power points (3.97 DPS) | yes | Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Stone Guard's Chain Legplates (220833, -0.46 DPS) [vendor]; Oilskin Leggings (9414, -0.57 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.3 ranged_attack_power points (3.78 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; First Sergeant's Chain Sabatons (220836, -0.95 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 38.7 ranged_attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.95 DPS) [dungeon]; Legionnaire's Band (19511, -0.95 DPS) [rep]; White Bone Band (11862, -1.01 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.7 ranged_attack_power points (1.89 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Legionnaire's Band (19511, -0.19 DPS) [rep]; White Bone Band (11862, -0.25 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -6.53 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (229.5 DPS) | yes | White Bone Shredder (11863, -0.38 DPS) [quest]; Force of the Hippogryph (9684, -3.13 DPS, sim-verified) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.54 DPS) [world_drop]; Gryphonwing Long Bow (13022, -2.05 DPS) [dungeon]; Dark Iron Rifle (16004, -2.37 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Blazewind Breastplate; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Jhordy's Misplaced Screwdriver; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 826, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1678 Black Ogre Kickers; 1832 Lucky Trousers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw

### Band 60 (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 307.3. Weights run: 1.2s. Verify run: 12.0s. 1477 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.832 ± 0.023, crit=1.308 ± 0.048 per rating point (14 rating = 1%, 18.314 per %), hit=2.044 ± 0.181 per rating point (10 rating = 1%, 20.439 per %), melee_haste=15.922 ± 2.619

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (307.3 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -9.77 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 53.8 ranged_attack_power points (3.68 DPS) | yes | Scout's Medallion (19534, -0.77 DPS) [rep]; Maelstrom's Tendril (19620, -0.77 DPS) [quest]; Beads of Ogre Might (22150, -0.88 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 94.8 ranged_attack_power points (6.48 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -0.05 DPS) [pvp]; Warlord's Chain Pauldrons (231565, -0.65 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.5 ranged_attack_power points (4.27 DPS) | yes | Deathguard's Cloak (20068, -0.98 DPS) [rep]; Shifting Cloak (18511, -0.98 DPS) [crafted]; Howler's Furs (272414, -1.22 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 143.5 ranged_attack_power points (9.81 DPS) | yes | Warlord's Chain Chestpiece (16565, -1.98 DPS) [vendor]; Warlord's Chain Hauberk (231573, -1.98 DPS) [vendor]; Obsidian Mail Tunic (22191, -5.01 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (307.3 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -9.33 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (307.3 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (307.3 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -8.79 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 156.2 ranged_attack_power points (10.68 DPS) | yes | Outrider's Chain Leggings (22673, -1.98 DPS, sim-verified) [rep]; General's Chain Legguards (231574, -2.84 DPS) [pvp]; Sentinel's Leather Pants (237818, -2.95 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (307.3 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -8.87 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 88.4 ranged_attack_power points (6.04 DPS) | yes | Don Julio's Band (19325, -2.30 DPS) [rep]; Masons Fraternity Ring (9533, -3.33 DPS) [quest]; Legionnaire's Band (19510, -3.91 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 60.1 ranged_attack_power points (4.11 DPS) | yes | Don Julio's Band (19325, -0.36 DPS) [rep]; Masons Fraternity Ring (9533, -1.40 DPS) [quest]; Legionnaire's Band (19510, -1.98 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (307.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (307.3 DPS) | yes | Molten Heart of the Mountain (249470, -1.41 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (307.3 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -8.34 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (307.3 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -11.44 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1477, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1678 Black Ogre Kickers; 1832 Lucky Trousers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw

### Band 60, raid preset (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 780.9. Weights run: 1.4s. Verify run: 11.6s. 1477 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.875 ± 0.022, crit=1.352 ± 0.044 per rating point (14 rating = 1%, 18.933 per %), hit=3.244 ± 0.247 per rating point (10 rating = 1%, 32.442 per %), melee_haste=16.861 ± 3.307

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 98.6 ranged_attack_power points (14.03 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -16.47 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 56.4 ranged_attack_power points (8.03 DPS) | yes | Amulet of the Darkmoon (19491, -0.26 DPS) [quest]; Scout's Medallion (19534, -1.89 DPS) [rep]; Maelstrom's Tendril (19620, -1.89 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 96.6 ranged_attack_power points (13.74 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 63.1 ranged_attack_power points (8.98 DPS) | yes | Shifting Cloak (18511, -2.03 DPS) [crafted]; Deathguard's Cloak (20068, -2.10 DPS) [rep]; Howler's Furs (272414, -3.22 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.8 ranged_attack_power points (22.45 DPS) | yes | Warlord's Chain Chestpiece (16565, -5.84 DPS) [vendor]; Warlord's Chain Hauberk (231573, -5.84 DPS) [vendor]; Shadowcraft Tunic (16721, -10.48 DPS, sim-verified) [dungeon] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (780.9 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -11.31 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 77.6 ranged_attack_power points (11.05 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; General's Chain Grips (231569, -1.13 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (780.9 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -9.03 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 170.9 ranged_attack_power points (24.32 DPS) | yes | Outrider's Chain Leggings (22673, -3.96 DPS, sim-verified) [rep]; General's Chain Legguards (231574, -7.72 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.89 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (780.9 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -13.22 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 101.4 ranged_attack_power points (14.43 DPS) | yes | Don Julio's Band (19325, -4.85 DPS) [rep]; Masons Fraternity Ring (9533, -8.71 DPS) [quest]; Channeler's Ring (272406, -9.82 DPS) [vendor] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 72.7 ranged_attack_power points (10.34 DPS) | yes | Don Julio's Band (19325, -0.76 DPS) [rep]; Masons Fraternity Ring (9533, -4.62 DPS) [quest]; Channeler's Ring (272406, -5.73 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (780.9 DPS) | yes | Frozen Heart of the Mountain (249469, -13.89 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (780.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (780.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -18.37 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (780.9 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -26.82 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1477, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1387 Ghoulfang; 1539 Gnarled Hermit's Staff; 1678 Black Ogre Kickers; 1832 Lucky Trousers; 1917 Jeweled Dagger; 1965 White Wolf Gloves; 2035 Sword of the Night Sky; 2069 Black Bear Hide Vest; 2203 Brashclaw's Chopper; 2204 Brashclaw's Skewer; 2235 Brackclaw

