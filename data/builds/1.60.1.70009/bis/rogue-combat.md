# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.3. Weights run: 1.1s. Verify run: 1.3s. 194 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.029 ± 0.004 per rating point (14 rating = 1%, 0.404 per %), hit=0.031 ± 0.003 per rating point (10 rating = 1%, 0.311 per %), melee_haste=not significant (0.665 ± 0.825)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.38 DPS) | yes | Lucky Fishing Hat (19972, -0.54 DPS, sim-verified) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.14 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.24 DPS) | yes | Forest Leather Mantle (4709, -0.34 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Catacomb Cloak (279899, +0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Hide of Lupos (3018, -0.10 DPS) [world] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.1 attack_power points (0.53 DPS) | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.19 DPS) [crafted]; Prospector's Chestpiece (14562, -0.24 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 attack_power points (0.24 DPS) | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.07 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.10 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.85 DPS) | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Dark Leather Belt (4249, -0.66 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 9.1 attack_power points (0.43 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 attack_power points (0.38 DPS) | yes | Blackened Defias Boots (10402, -0.10 DPS) [dungeon]; Bristlebark Boots (14568, -0.14 DPS) [world_drop]; Footpads of the Fang (10411, -0.33 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.29 DPS) | yes | Protector's Band (20439, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 4.2 attack_power points (0.20 DPS) | yes | Protector's Band (20439, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The 1 Ring (8350, -0.15 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.76 DPS) | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Barrens Basher (274744, -1.21 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 attack_power points (10.85 DPS) | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.19 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 194, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14145 Cursed Felblade

### Band 30 (night-elf, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 55.3. Weights run: 1.1s. Verify run: 1.3s. 328 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.035 ± 0.004 per rating point (14 rating = 1%, 0.485 per %), hit=0.043 ± 0.003 per rating point (10 rating = 1%, 0.428 per %), melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 attack_power points (0.48 DPS) | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.11 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Sentinel's Medallion (19541, -0.32 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.38 DPS) [rep]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 attack_power points (0.53 DPS) | yes | Insignia Mantle (4721, -0.19 DPS) [world_drop]; Cloudy Gustwoven Spaulders (277043, -0.19 DPS) [crafted]; Mantle of Thieves (2264, -0.33 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.48 DPS) | yes | Tigerstrike Mantle (13108, -0.10 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.19 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.77 DPS) | yes | Dusky Leather Armor (7374, -0.08 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.24 DPS) [quest]; Green Leather Armor (4255, -0.38 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.48 DPS) | yes | Unearthed Bands (9428, -0.11 DPS, sim-verified) [dungeon]; Jurassic Wristguards (6198, -0.19 DPS) [world]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.77 DPS) | yes | Pilferer's Gloves (7358, -0.43 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -0.48 DPS) [world_drop]; Ebon Vise (7690, -0.48 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.15 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.72 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.25 DPS) | yes | Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.63 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.58 DPS) | yes | Feet of the Lynx (1121, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor]; Insignia Boots (4055, -0.21 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 attack_power points (0.43 DPS) | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Monkey Ring (6748, -0.10 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.48 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -0.12 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.20 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | sim-verified (55.3 DPS) | yes | Satyr's Rod (15962, -15.35 DPS) [world_drop]; Swinetusk Shank (6691, -17.36 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Silver Star (3463, -0.21 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop]; Crystalpine Stinger (13037, -0.24 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Royal Diplomatic Scepter; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (night-elf, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 97.6. Weights run: 1.2s. Verify run: 1.4s. 457 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.066 ± 0.011 per rating point (14 rating = 1%, 0.925 per %), hit=0.065 ± 0.008 per rating point (10 rating = 1%, 0.646 per %), melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.3 attack_power points (0.66 DPS) | yes | Nightscape Headband (8176, -0.09 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.99 DPS) | yes | Ghostshard Talisman (7731, -0.39 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19540, -0.43 DPS) [rep]; Sentinel's Medallion (19541, -0.59 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 attack_power points (1.15 DPS) | yes | Forest Tracker Epaulets (2278, -0.60 DPS) [world_drop]; Flintrock Shoulders (7755, -0.65 DPS) [dungeon]; Nightscape Shoulders (8192, -0.79 DPS, sim-verified) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.3 attack_power points (0.51 DPS) | yes | Imperial Cloak (6432, -0.10 DPS) [world_drop]; Parachute Cloak (10518, -0.10 DPS) [crafted]; Wolfmaster Cape (6314, -0.56 DPS, sim-verified) [dungeon] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 19.5 attack_power points (0.97 DPS) | yes | Raptorbane Armor (3566, -0.17 DPS) [quest]; Nightscape Tunic (8175, -0.20 DPS) [crafted]; Wolffear Harness (13110, -0.71 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.65 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.9 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.53 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.53 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.08 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.49 DPS) | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.39 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Hawkeye's Breeches (14595, -0.58 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 attack_power points (0.66 DPS) | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Skulker's Leather Shoes (252531, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.45 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Insurgent's Band (272066, -0.40 DPS) [vendor]; Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Falcon's Hook (7552, -0.53 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Insurgent's Band (272066, -0.52 DPS, sim-verified) [vendor]; Falcon's Hook (7552, -0.53 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.54 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Dazzling Longsword (869, -1.68 DPS) [world_drop]; Southsea Lamp (9359, -1.95 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (97.6 DPS) | yes | Ardent Custodian (868, -21.79 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -21.88 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Swiftwind (13038, -0.34 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.92 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 457, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 50 (night-elf, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 151.7. Weights run: 1.2s. Verify run: 1.8s. 585 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=0.470 ± 0.028 per rating point (14 rating = 1%, 6.578 per %), hit=0.080 ± 0.010 per rating point (10 rating = 1%, 0.803 per %), melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 36.5 attack_power points (1.95 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.70 DPS) [vendor]; Helm of Fire (8348, -0.93 DPS) [crafted]; Embrace of the Lycan (9479, -1.29 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.07 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.32 DPS) [dungeon]; Sentinel's Medallion (19539, -0.35 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.4 attack_power points (1.30 DPS) | yes | Phytoskin Spaulders (17749, -0.39 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Leather Shoulders (220852, -0.48 DPS) [vendor]; Penance Spaulders (11963, -0.64 DPS) [quest] |
| back | Blisterbane Wrap (12552) (or Dark Phantom Cape (13122)) | Blackrock Depths: Anvilrage Overseer [dungeon] | 16.9 attack_power points (0.90 DPS) | yes | Dark Phantom Cape (13122, +0.00 DPS, sim-verified) [world_drop]; Blackveil Cape (11626, -0.06 DPS) [dungeon]; Duskbat Drape (19982, -0.06 DPS) [quest] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 28.1 attack_power points (1.50 DPS) | yes | Blazewind Breastplate (11193, +0.00 DPS, sim-verified) [quest]; Knight's Leather Armor (220854, -0.14 DPS) [vendor]; Quillward Harness (10583, -0.36 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.07 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 26.6 attack_power points (1.42 DPS) | yes | Darkmantle Grips (226828, +0.00 DPS, sim-verified) [vendor]; Raider Gloves (272100, -0.10 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.32 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) (or Highlander's Leather Girdle (20116)) | Blackrock Depths: Eviscerator [dungeon] | 30.0 attack_power points (1.60 DPS) | yes | Highlander's Leather Girdle (20115, -0.18 DPS) [rep]; Highlander's Chain Girdle (20090, -0.32 DPS) [rep]; Highlander's Leather Girdle (20116, -0.79 DPS, sim-verified) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (151.7 DPS) | yes | Basilisk Hide Pants (1718, -0.10 DPS) [world_drop]; Keeper's Woolies (14668, -0.22 DPS) [world_drop]; Ferine Leggings (6690, -2.23 DPS, sim-verified) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 attack_power points (1.20 DPS) | yes | Sandstalker Ankleguards (12470, -0.18 DPS) [dungeon]; Sergeant Major's Leather Boots (220860, -0.24 DPS) [vendor]; Whisperwalk Boots (20255, -0.31 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.8 attack_power points (1.11 DPS) | yes | Mark of Kern (2262, -0.04 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.27 DPS) [quest]; Insurgent's Band (272065, -0.31 DPS) [vendor] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (1.07 DPS) | yes | Mark of Kern (2262, +0.00 DPS, sim-verified) [dungeon]; Masons Fraternity Ring (9533, -0.23 DPS) [quest]; Insurgent's Band (272065, -0.27 DPS) [vendor] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | 0.0 attack_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -0.98 DPS, sim-verified) [crafted] |
| trinket2 | Sanctified Orb (20512) | Forging the Mightstone [quest] | 0.0 attack_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hanzo Sword (8190, -5.32 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (29.13 DPS) | yes | Claw of Celebras (17738, -3.36 DPS) [dungeon]; Thorium Cestus (250614, -3.79 DPS, sim-verified) [crafted]; Thermotastic Egg Timer (9644, -28.95 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, -0.09 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.09 DPS) [world_drop]; Dark Iron Rifle (16004, -2.20 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; back: Blisterbane Wrap; chest: Fungus Shroud Armor; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Assault Band; trinket1: Diamond Flask; trinket2: Sanctified Orb; main_hand: Bloodrazor; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 585, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 60 (night-elf, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 290.5. Weights run: 1.2s. Verify run: 1.7s. 1310 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=0.587 ± 0.039 per rating point (14 rating = 1%, 8.223 per %), hit=0.090 ± 0.016 per rating point (10 rating = 1%, 0.900 per %), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 78.5 attack_power points (4.20 DPS) | yes | Field Marshal's Leather Mask (231545, -1.69 DPS) [vendor]; Lieutenant Commander's Leather Helm (227055, -1.78 DPS) [vendor]; Duskwraith Mask (239550, -9.80 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 32.2 attack_power points (1.73 DPS) | yes | Will of the Martyr (17044, -0.12 DPS) [quest]; Mark of Fordring (15411, -0.30 DPS) [quest]; Imperial Jewel (11933, -1.41 DPS, sim-verified) [dungeon] |
| shoulder | Duskwraith Pauldrons (239559) | Leonid Barthalomew the Revered [vendor] | sim-verified (290.5 DPS) | yes | Highlander's Lizardhide Shoulders (20060, -0.09 DPS) [rep]; Darkspear Pauldrons (272105, -0.13 DPS) [vendor]; Highlander's Leather Shoulders (20059, -4.21 DPS, sim-verified) [rep] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 41.0 attack_power points (2.19 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Howler's Furs (272414, -0.65 DPS) [vendor]; Stormpike Soldier's Cloak (19084, -0.91 DPS) [rep] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Cadaverous Armor (14637, -0.34 DPS) [dungeon]; Dawn Armor (252483, -0.90 DPS) [crafted]; Tunic of Undead Slaying (23089, -18.64 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Duskwraith Bracers (239555, -0.12 DPS) [vendor]; Marshal's Leather Armsplints (16460, -0.61 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -5.12 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 58.0 attack_power points (3.11 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.64 DPS) [pvp]; Skul's Fingerbone Claws (13395, -0.96 DPS) [dungeon]; Cadaverous Gloves (14640, -8.36 DPS, sim-verified) [dungeon] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 55.2 attack_power points (2.96 DPS) | yes | Cadaverous Belt (14636, -0.81 DPS) [dungeon]; Assassin's Waistguard (272395, -1.29 DPS) [vendor]; Highlander's Leather Girdle (20045, -6.36 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 79.9 attack_power points (4.28 DPS) | yes | Sentinel's Leather Pants (237818, -1.38 DPS) [vendor]; Cadaverous Leggings (14638, -1.49 DPS) [dungeon]; Devilsaur Leggings (15062, -10.42 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 58.0 attack_power points (3.11 DPS) | yes | Darkmantle Footpads (226831, -1.28 DPS) [vendor]; Darkmantle Boots (22003, -1.31 DPS) [quest]; Pads of the Dread Wolf (13210, -9.22 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.15 DPS) [quest]; Naglering (11669, -7.22 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 0.0 attack_power points (0.00 DPS) | yes | Tarnished Elven Ring (18500, -0.22 DPS) [dungeon]; Innervating Band (18701, -0.22 DPS) [dungeon]; Naglering (11669, -6.69 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+12.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -0.59 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -1.03 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -4.71 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 815.9 attack_power points (43.68 DPS) | yes | Core Hound Tooth (18805, +0.00 DPS, sim-verified) [world_drop]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Dal'Rend's Tribal Guardian (12939, -12.66 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.13 DPS) [world_drop]; Dark Iron Rifle (16004, -2.76 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Medallion of the Dawn; shoulder: Duskwraith Pauldrons; back: Cloak of the Honor Guard; chest: Duskwraith Breastplate; wrist: Bracers of the Eclipse; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1310, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.8. Weights run: 1.1s. Verify run: 1.2s. 193 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.029 ± 0.004 per rating point (14 rating = 1%, 0.404 per %), hit=0.031 ± 0.003 per rating point (10 rating = 1%, 0.311 per %), melee_haste=not significant (0.665 ± 0.825)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.38 DPS) | yes | Lucky Fishing Hat (19972, -0.51 DPS, sim-verified) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.24 DPS) | yes | Forest Leather Mantle (4709, -0.32 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Catacomb Cloak (279899, +0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Hide of Lupos (3018, -0.10 DPS) [world] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 attack_power points (0.34 DPS) | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Prospector's Chestpiece (14562, -0.05 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.32 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 attack_power points (0.24 DPS) | yes | Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor]; Spare Part Bindings (279875, -0.10 DPS) [quest] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.10 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.85 DPS) | yes | Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world_drop]; Dark Leather Belt (4249, -0.66 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 9.1 attack_power points (0.43 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 attack_power points (0.38 DPS) | yes | Blackened Defias Boots (10402, -0.10 DPS) [dungeon]; Bristlebark Boots (14568, -0.14 DPS) [world_drop]; Footpads of the Fang (10411, -0.32 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.29 DPS) | yes | Legionnaire's Band (20429, -0.10 DPS) [rep]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 4.2 attack_power points (0.20 DPS) | yes | Legionnaire's Band (20429, +0.00 DPS, sim-verified) [rep]; Bounty Hunter's Ring (5351, -0.05 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.76 DPS) | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Wingblade (6504, -1.15 DPS) [quest] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 attack_power points (10.85 DPS) | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.19 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 54.1. Weights run: 1.1s. Verify run: 1.3s. 328 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.035 ± 0.004 per rating point (14 rating = 1%, 0.485 per %), hit=0.043 ± 0.003 per rating point (10 rating = 1%, 0.428 per %), melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 attack_power points (0.48 DPS) | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.11 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Scout's Medallion (19537, -0.32 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.38 DPS) [rep]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 attack_power points (0.53 DPS) | yes | Insignia Mantle (4721, -0.19 DPS) [world_drop]; Cloudy Gustwoven Spaulders (277043, -0.19 DPS) [crafted]; Mantle of Thieves (2264, -0.33 DPS, sim-verified) [dungeon] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.48 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 attack_power points (0.68 DPS) | yes | Panther Armor (6670, -0.17 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.29 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.29 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.48 DPS) | yes | Unearthed Bands (9428, -0.11 DPS, sim-verified) [dungeon]; Jurassic Wristguards (6198, -0.19 DPS) [world]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.77 DPS) | yes | Pilferer's Gloves (7358, -0.43 DPS, sim-verified) [crafted]; Braced Handguards (6784, -0.43 DPS) [quest]; Ebon Vise (7690, -0.48 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.15 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Deftkin Belt (16659, -0.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.25 DPS) | yes | Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.64 DPS, sim-verified) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 8.0 attack_power points (0.39 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS, sim-verified) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 attack_power points (0.43 DPS) | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Monkey Ring (6748, -0.10 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.48 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -0.12 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.20 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | sim-verified (54.1 DPS) | yes | Satyr's Rod (15962, -15.35 DPS) [world_drop]; Swinetusk Shank (6691, -16.86 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Silver Star (3463, -0.21 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop]; Crystalpine Stinger (13037, -0.24 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wildhunter Cloak; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Royal Diplomatic Scepter; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 96.2. Weights run: 1.2s. Verify run: 1.4s. 457 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.066 ± 0.011 per rating point (14 rating = 1%, 0.925 per %), hit=0.065 ± 0.008 per rating point (10 rating = 1%, 0.646 per %), melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.3 attack_power points (0.66 DPS) | yes | White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop]; Nightscape Headband (8176, -0.11 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.99 DPS) | yes | Ghostshard Talisman (7731, -0.39 DPS, sim-verified) [dungeon]; Scout's Medallion (19536, -0.43 DPS) [rep]; Scout's Medallion (19537, -0.59 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 attack_power points (1.15 DPS) | yes | Forest Tracker Epaulets (2278, -0.60 DPS) [world_drop]; Flintrock Shoulders (7755, -0.65 DPS) [dungeon]; Nightscape Shoulders (8192, -0.78 DPS, sim-verified) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.3 attack_power points (0.51 DPS) | yes | Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Parachute Cloak (10518, -0.10 DPS) [crafted]; Wildhunter Cloak (16658, -0.59 DPS, sim-verified) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 19.5 attack_power points (0.97 DPS) | yes | Nightscape Tunic (8175, -0.20 DPS) [crafted]; Dusky Leather Armor (7374, -0.25 DPS) [crafted]; Wolffear Harness (13110, -0.72 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.65 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.9 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.53 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.53 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.17 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.49 DPS) | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.39 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Hawkeye's Breeches (14595, -0.58 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 attack_power points (0.66 DPS) | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Insurgent's Band (272066, -0.40 DPS) [vendor]; Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Falcon's Hook (7552, -0.53 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Insurgent's Band (272066, -0.52 DPS, sim-verified) [vendor]; Falcon's Hook (7552, -0.53 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.54 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Dazzling Longsword (869, -1.68 DPS) [world_drop]; Southsea Lamp (9359, -1.95 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (96.2 DPS) | yes | Ardent Custodian (868, -21.48 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -21.88 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Swiftwind (13038, -0.34 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.90 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 457, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 155.0. Weights run: 1.2s. Verify run: 1.8s. 586 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=0.470 ± 0.028 per rating point (14 rating = 1%, 6.578 per %), hit=0.080 ± 0.010 per rating point (10 rating = 1%, 0.803 per %), melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 36.5 attack_power points (1.95 DPS) | yes | Blood Guard's Leather Headband (220851, -0.70 DPS) [vendor]; Helm of Fire (8348, -0.93 DPS) [crafted]; Embrace of the Lycan (9479, -1.19 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.07 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.32 DPS) [dungeon]; Scout's Medallion (19535, -0.35 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.4 attack_power points (1.30 DPS) | yes | Phytoskin Spaulders (17749, -0.42 DPS, sim-verified) [dungeon]; Blood Guard's Leather Shoulders (220853, -0.48 DPS) [vendor]; Penance Spaulders (11963, -0.64 DPS) [quest] |
| back | Blisterbane Wrap (12552) (or Dark Phantom Cape (13122)) | Blackrock Depths: Anvilrage Overseer [dungeon] | 16.9 attack_power points (0.90 DPS) | yes | Dark Phantom Cape (13122, +0.00 DPS, sim-verified) [world_drop]; Blackveil Cape (11626, -0.06 DPS) [dungeon]; Duskbat Drape (19982, -0.06 DPS) [quest] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 28.1 attack_power points (1.50 DPS) | yes | Blazewind Breastplate (11193, +0.00 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -0.14 DPS) [vendor]; Quillward Harness (10583, -0.36 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.07 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 26.6 attack_power points (1.42 DPS) | yes | Darkmantle Grips (226828, +0.00 DPS, sim-verified) [vendor]; Raider Gloves (272100, -0.10 DPS) [vendor]; First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) (or Defiler's Leather Girdle (20192)) | Blackrock Depths: Eviscerator [dungeon] | 30.0 attack_power points (1.60 DPS) | yes | Defiler's Leather Girdle (20193, -0.18 DPS) [rep]; Defiler's Chain Girdle (20152, -0.32 DPS) [rep]; Defiler's Leather Girdle (20192, -0.79 DPS, sim-verified) [rep] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | sim-verified (155.0 DPS) | yes | Basilisk Hide Pants (1718, -0.10 DPS) [world_drop]; Keeper's Woolies (14668, -0.22 DPS) [world_drop]; Ferine Leggings (6690, -2.18 DPS, sim-verified) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 attack_power points (1.20 DPS) | yes | Sandstalker Ankleguards (12470, -0.18 DPS) [dungeon]; First Sergeant's Leather Boots (220861, -0.24 DPS) [vendor]; Whisperwalk Boots (20255, -0.29 DPS, sim-verified) [quest] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.28 DPS) | yes | Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.44 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.8 attack_power points (1.11 DPS) | yes | Mark of Kern (2262, -0.04 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.27 DPS) [quest]; Assault Band (13095, -1.66 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.24 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | 0.0 attack_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Fire Ruby (20036, -3.18 DPS, sim-verified) [quest] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hanzo Sword (8190, -4.73 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (29.13 DPS) | yes | Claw of Celebras (17738, -3.36 DPS) [dungeon]; Thorium Cestus (250614, -4.34 DPS, sim-verified) [crafted]; White Bone Shredder (11863, -5.46 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, -0.09 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.09 DPS) [world_drop]; Dark Iron Rifle (16004, -2.12 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; back: Blisterbane Wrap; chest: Fungus Shroud Armor; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Bloodrazor; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 586, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 290.4. Weights run: 1.2s. Verify run: 1.6s. 1312 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=0.587 ± 0.039 per rating point (14 rating = 1%, 8.223 per %), hit=0.090 ± 0.016 per rating point (10 rating = 1%, 0.900 per %), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 78.5 attack_power points (4.20 DPS) | yes | Warlord's Leather Helm (231553, -1.69 DPS) [vendor]; Champion's Leather Helm (227057, -1.78 DPS) [vendor]; Duskwraith Mask (239550, -9.65 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 32.2 attack_power points (1.73 DPS) | yes | Will of the Martyr (17044, -0.12 DPS) [quest]; Mark of Fordring (15411, -0.30 DPS) [quest]; Imperial Jewel (11933, -1.44 DPS, sim-verified) [dungeon] |
| shoulder | Duskwraith Pauldrons (239559) | Leonid Barthalomew the Revered [vendor] | sim-verified (290.4 DPS) | yes | Defiler's Lizardhide Shoulders (20175, -0.09 DPS) [rep]; Darkspear Pauldrons (272105, -0.13 DPS) [vendor]; Defiler's Leather Shoulders (20194, -4.97 DPS, sim-verified) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 41.0 attack_power points (2.19 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Howler's Furs (272414, -0.65 DPS) [vendor]; Frostwolf Legionnaire's Cloak (19083, -0.91 DPS) [rep] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | 0.0 attack_power points (0.00 DPS) | yes | Cadaverous Armor (14637, -0.34 DPS) [dungeon]; Dawn Armor (252483, -0.90 DPS) [crafted]; Tunic of Undead Slaying (23089, -18.35 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Duskwraith Bracers (239555, -0.12 DPS) [vendor]; General's Leather Armsplints (16559, -0.61 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -5.05 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 58.0 attack_power points (3.11 DPS) | yes | Blood Guard's Leather Vices (16499, -0.64 DPS) [pvp]; Skul's Fingerbone Claws (13395, -0.96 DPS) [dungeon]; Cadaverous Gloves (14640, -8.14 DPS, sim-verified) [dungeon] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 55.2 attack_power points (2.96 DPS) | yes | Cadaverous Belt (14636, -0.81 DPS) [dungeon]; Assassin's Waistguard (272395, -1.29 DPS) [vendor]; Defiler's Leather Girdle (20190, -6.23 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 79.9 attack_power points (4.28 DPS) | yes | Sentinel's Leather Pants (237818, -1.38 DPS) [vendor]; Cadaverous Leggings (14638, -1.49 DPS) [dungeon]; Devilsaur Leggings (15062, -11.97 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 58.0 attack_power points (3.11 DPS) | yes | Darkmantle Footpads (226831, -1.28 DPS) [vendor]; Darkmantle Boots (22003, -1.31 DPS) [quest]; Pads of the Dread Wolf (13210, -8.98 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.15 DPS) [quest]; Naglering (11669, -7.08 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 0.0 attack_power points (0.00 DPS) | yes | White Bone Band (11862, -0.06 DPS) [quest]; Tarnished Elven Ring (18500, -0.22 DPS) [dungeon]; Naglering (11669, -6.51 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+13.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 attack_power points (0.00 DPS) | yes | Counterattack Lodestone (18537, -1.10 DPS) [dungeon]; Hand of Justice (11815, -2.02 DPS, sim-verified) [dungeon]; Blackhand's Breadth (13965, -2.22 DPS) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -3.74 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 815.9 attack_power points (43.68 DPS) | yes | Core Hound Tooth (18805, +0.00 DPS, sim-verified) [world_drop]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Dal'Rend's Tribal Guardian (12939, -12.66 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.13 DPS) [world_drop]; Dark Iron Rifle (16004, -2.73 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Medallion of the Dawn; shoulder: Duskwraith Pauldrons; back: Deathguard's Cloak; chest: Duskwraith Breastplate; wrist: Bracers of the Eclipse; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1312, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

