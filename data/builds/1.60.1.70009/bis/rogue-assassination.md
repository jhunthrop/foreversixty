# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 34.5. Weights run: 1.6s. Verify run: 1.1s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.313 ± 0.007, crit=0.643 ± 0.012 per rating point (14 rating = 1%, 9.004 per %), hit=0.873 ± 0.032 per rating point (10 rating = 1%, 8.735 per %), melee_haste=5.171 ± 0.355

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 10.5 attack_power points (0.56 DPS) | yes | Defender's Leather Hood (252447, -0.13 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 7.9 attack_power points (0.42 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.6 attack_power points (0.35 DPS) | yes | Slime-encrusted Pads (6461, -0.42 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 7.9 attack_power points (0.42 DPS) | yes | Cape of the Brotherhood (5193, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 14.4 attack_power points (0.77 DPS) | yes | Brawler's Leather Armor (252490, -0.01 DPS) [crafted]; Defender's Leather Armor (252434, -0.19 DPS) [crafted]; Prospector's Chestpiece (14562, -0.24 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 7.3 attack_power points (0.39 DPS) | yes | Forest Leather Bracers (3202, -0.04 DPS) [world_drop]; Bristlebark Bindings (14569, -0.07 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 11.9 attack_power points (0.64 DPS) | yes | Brawler's Leather Gloves (252494, -0.14 DPS) [crafted]; Fletcher's Gloves (7348, -0.15 DPS) [crafted]; Bristlebark Gloves (14572, -0.17 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.97 DPS) | yes | Brawler's Leather Belt (252428, -0.47 DPS) [crafted]; Deviate Scale Belt (6468, -0.60 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.61 DPS) [quest] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 17.8 attack_power points (0.96 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.19 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 13.5 attack_power points (0.72 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.30 DPS) [dungeon]; Footpads of the Fang (10411, -0.30 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 9.3 attack_power points (0.50 DPS) | yes | Signet of the Zhevra (285330, -0.07 DPS) [world]; Demon Band (12054, -0.28 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 8.7 attack_power points (0.47 DPS) | yes | Signet of the Zhevra (285330, -0.05 DPS) [world]; Demon Band (12054, -0.25 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.35 DPS) | yes | Blackfang (2236, -1.09 DPS) [world_drop]; Assassin's Blade (1935, -1.25 DPS) [dungeon]; Diamond Hammer (2194, -2.66 DPS, sim-verified) [world_drop] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (34.5 DPS) | yes | Diamond Hammer (2194, -0.42 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.3 attack_power points (0.28 DPS) | yes | Fine Longbow (11304, -0.07 DPS) [vendor]; Deadly Blunderbuss (4369, -0.14 DPS) [crafted]; Light Bow (4576, -0.14 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 61.1. Weights run: 1.6s. Verify run: 1.0s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.364 ± 0.009, crit=0.741 ± 0.015 per rating point (14 rating = 1%, 10.371 per %), hit=1.143 ± 0.043 per rating point (10 rating = 1%, 11.426 per %), melee_haste=6.763 ± 0.450

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 13.6 attack_power points (0.77 DPS) | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.15 DPS) [world]; Brawler's Leather Hood (252504, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.79 DPS) | yes | Sentinel's Medallion (19541, -0.18 DPS) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.0 attack_power points (1.13 DPS) | yes | Mantle of Thieves (2264, -0.45 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.46 DPS) [crafted]; Bristlebark Amice (14573, -0.50 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.5 attack_power points (0.71 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.18 DPS) [pvp] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 19.1 attack_power points (1.08 DPS) | yes | Brawler's Leather Tunic (252508, -0.12 DPS) [crafted]; Raptorbane Armor (3566, -0.18 DPS) [quest]; Tunic of Westfall (2041, -0.23 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.2 attack_power points (0.69 DPS) | yes | Jurassic Wristguards (6198, -0.11 DPS) [world]; Cultist's Armguards (270032, -0.12 DPS) [quest]; Barbaric Bracers (18948, -0.15 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.91 DPS) | yes | Insignia Gloves (6408, -0.05 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.16 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.36 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.32 DPS) [crafted]; Blackened Defias Belt (10403, -0.34 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.47 DPS) | yes | Petrolspill Leggings (9509, -0.39 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.39 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.50 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 13.9 attack_power points (0.79 DPS) | yes | Disjointed Shoes (277226, -0.11 DPS) [quest]; Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Insignia Boots (4055, -0.17 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 16.3 attack_power points (0.92 DPS) | yes | Thunderbrow Ring (13097, -0.24 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.35 DPS) [quest]; Monkey Ring (6748, -0.38 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.2 attack_power points (0.80 DPS) | yes | Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.24 DPS) [quest]; Monkey Ring (6748, -0.26 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (18.59 DPS) | yes | Swinetusk Shank (6691, -0.33 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.44 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.54 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (18.29 DPS) | yes | Swinetusk Shank (6691, -12.54 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -18.21 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.51 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Silver Star (3463, -0.12 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.20 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 32502110551501000-00000000000000000-0000000000000000000)

Set DPS (verified): 111.6. Weights run: 1.8s. Verify run: 1.1s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.369 ± 0.011, crit=0.767 ± 0.018 per rating point (14 rating = 1%, 10.741 per %), hit=2.206 ± 0.075 per rating point (10 rating = 1%, 22.058 per %), melee_haste=26.345 ± 0.714

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.1 attack_power points (0.89 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.28 DPS) [world_drop]; Nightscape Headband (8176, -0.33 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.69 DPS) | yes | Sentinel's Medallion (19540, -0.17 DPS) [rep]; Ghostshard Talisman (7731, -0.21 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.1 attack_power points (0.93 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Flintrock Shoulders (7755, -0.29 DPS) [dungeon]; Nightscape Shoulders (8192, -0.41 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.7 attack_power points (0.61 DPS) | yes | Sergeant Major's Cape (16336, -0.12 DPS) [pvp]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.22 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 34.0 attack_power points (1.17 DPS) | yes | Wolffear Harness (13110, -0.40 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.46 DPS) [crafted]; Dusky Leather Armor (7374, -0.51 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Hawkeye's Bracers (14590, -0.27 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.31 DPS) [dungeon]; Dusky Bracers (7378, -0.31 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.7 attack_power points (1.06 DPS) | yes | Skulker's Leather Gloves (252525, -0.28 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.29 DPS) [crafted]; Imperial Leather Gloves (4063, -0.32 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.03 DPS) | yes | Highlander's Chain Girdle (20090, -0.21 DPS) [rep]; Ogron's Sash (13117, -0.30 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.40 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 28.8 attack_power points (0.99 DPS) | yes | Triprunner Dungarees (9624, -0.04 DPS) [quest]; Ferine Leggings (6690, -0.09 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.32 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.1 attack_power points (0.76 DPS) | yes | Prowler's Leather Shoes (252465, -0.05 DPS) [crafted]; Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Excelsior Boots (4109, -0.09 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Protector's Band (19515, -0.04 DPS) [rep]; Field Researcher's Loop (281634, -0.12 DPS) [quest]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.69 DPS) | yes | Protector's Band (19515, -0.04 DPS) [rep]; Field Researcher's Loop (281634, -0.12 DPS) [quest]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.49 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.12 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (111.6 DPS) | yes | Stonecloth Branch (15963, -15.07 DPS) [world_drop]; Satyr's Rod (15962, -15.12 DPS) [world_drop]; Ardent Custodian (868, -64.19 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.13 DPS) [dungeon]; Swiftwind (13038, -0.15 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.51 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 175.0. Weights run: 2.4s. Verify run: 1.5s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.467 ± 0.009, crit=0.987 ± 0.016 per rating point (14 rating = 1%, 13.812 per %), hit=2.105 ± 0.089 per rating point (10 rating = 1%, 21.053 per %), melee_haste=6.379 ± 1.375

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 50.9 attack_power points (4.09 DPS) | yes | Ebon Mask (19984, -0.08 DPS) [quest]; Embrace of the Lycan (9479, -0.87 DPS) [dungeon]; White Bandit Mask (10008, -1.91 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.1 attack_power points (1.94 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.33 DPS) [quest]; Sentinel's Medallion (19539, -0.52 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 42.9 attack_power points (3.45 DPS) | yes | Sunburn Spaulders (274751, -0.74 DPS, sim-verified) [vendor]; Skulker's Leather Shoulder (252535, -1.43 DPS) [crafted]; Failed Flying Experiment (9647, -1.46 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.5 attack_power points (2.14 DPS) | yes | Blisterbane Wrap (12552, -0.36 DPS) [dungeon]; Dark Phantom Cape (13122, -0.36 DPS) [world_drop]; Duskbat Drape (19982, -0.48 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 52.9 attack_power points (4.25 DPS) | yes | Warbear Harness (15064, -0.79 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -1.30 DPS) [quest]; Fungus Shroud Armor (17742, -1.30 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 26.0 attack_power points (2.09 DPS) | yes | Skulker's Leather Bracers (252540, -0.47 DPS) [crafted]; Branded Leather Bracers (19508, -0.48 DPS) [dungeon]; Pridelord Bands (14672, -0.51 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 42.3 attack_power points (3.40 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.68 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.02 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.06 DPS) | yes | Highlander's Leather Girdle (20115, -0.34 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.44 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.51 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 52.9 attack_power points (4.25 DPS) | yes | Gryphon Rider's Leggings (9652, -0.85 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -1.32 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.77 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 30.9 attack_power points (2.49 DPS) | yes | Skulker's Leather Boots (252469, -0.07 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.15 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.1 attack_power points (3.30 DPS) | yes | Masons Fraternity Ring (9533, -1.65 DPS) [quest]; Mark of Kern (2262, -1.69 DPS) [dungeon]; Assault Band (13095, -1.69 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 23.2 attack_power points (1.87 DPS) | yes | Masons Fraternity Ring (9533, -0.21 DPS) [quest]; Mark of Kern (2262, -0.26 DPS) [dungeon]; Assault Band (13095, -0.26 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (175.0 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (175.0 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (175.0 DPS) | yes | Thorium Cestus (250614, -2.16 DPS) [crafted]; Doomforged Straightedge (12535, -2.43 DPS) [dungeon]; Hanzo Sword (8190, -2.50 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (43.86 DPS) | yes | Thorium Cestus (250614, +0.00 DPS) [crafted]; Claw of Celebras (17738, -5.06 DPS) [dungeon]; Thermotastic Egg Timer (9644, -43.51 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (175.0 DPS) | yes | Stinging Bow (10624, -0.28 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.28 DPS) [world_drop]; Dark Iron Rifle (16004, -1.95 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 275.8. Weights run: 2.3s. Verify run: 1.5s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.055 ± 0.022, crit=2.233 ± 0.038 per rating point (14 rating = 1%, 31.256 per %), hit=4.760 ± 0.173 per rating point (10 rating = 1%, 47.597 per %), melee_haste=16.275 ± 2.607

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 126.4 attack_power points (11.31 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -1.04 DPS) [pvp]; Knight-Lieutenant's Leather Headband (220850, -2.83 DPS) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 78.4 attack_power points (7.01 DPS) | yes | Beads of Ogre Might (22150, -0.61 DPS) [quest]; Mark of Fordring (15411, -1.89 DPS) [quest]; Medallion of the Dawn (22659, -2.07 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 119.2 attack_power points (10.66 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, -1.64 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, -2.54 DPS) [pvp]; Wyrmhide Spaulders (12082, -2.65 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 75.6 attack_power points (6.76 DPS) | yes | Cape of the Black Baron (13340, -2.21 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.50 DPS) [vendor]; Stalwart Cloak (272415, -2.50 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (275.8 DPS) | yes | Darkmantle Tunic (226825, -1.77 DPS) [quest]; Field Marshal's Leather Chestpiece (231543, -2.02 DPS) [pvp]; Tunic of Undead Slaying (23089, -10.65 DPS, sim-verified) [world] |
| wrist | Blackmist Armguards (12966) | Blackrock Spire: The Beast [dungeon] | sim-verified (275.8 DPS) | yes | Forest Stalker's Bracers (19587, -0.23 DPS) [rep]; Marshal's Leather Armsplints (16460, -0.32 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -2.61 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 78.9 attack_power points (7.05 DPS) | yes | Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; Marshal's Leather Handgrips (231544, -0.58 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.75 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 138.4 attack_power points (12.37 DPS) | yes | Belt of Preserved Heads (20216, -1.61 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -6.54 DPS) [rep]; Ferocity of the Timbermaw (227805, -6.92 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 118.0 attack_power points (10.55 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -0.46 DPS) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 82.6 attack_power points (7.39 DPS) | yes | Fine Dawn Treaders (227815, -2.18 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -2.25 DPS) [dungeon]; Darkmantle Boots (22003, -2.98 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (275.8 DPS) | yes | Tarnished Elven Ring (18500, -1.47 DPS) [dungeon]; Cutthroat's Signet (272408, -1.65 DPS) [vendor]; Naglering (11669, -5.81 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (275.8 DPS) | yes | Tarnished Elven Ring (18500, -0.55 DPS) [dungeon]; Cutthroat's Signet (272408, -0.74 DPS) [vendor]; Naglering (11669, -5.29 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (275.8 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -8.72 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (275.8 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Heart of Wyrmthalak (22321, -2.01 DPS, sim-verified) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (275.8 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 834.8 attack_power points (74.64 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -7.39 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -22.83 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (275.8 DPS) | yes | Blackcrow (12651, -0.43 DPS) [dungeon]; The Purifier (22656, -2.01 DPS) [quest]; Dark Iron Rifle (16004, -2.59 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Blackmist Armguards; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 556.5. Weights run: 2.3s. Verify run: 1.6s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.124 ± 0.025, crit=2.326 ± 0.042 per rating point (14 rating = 1%, 32.558 per %), hit=4.712 ± 0.160 per rating point (10 rating = 1%, 47.119 per %), melee_haste=18.540 ± 2.831

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 126.8 attack_power points (20.75 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -1.82 DPS) [pvp]; Knight-Lieutenant's Leather Headband (220850, -5.09 DPS) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 79.0 attack_power points (12.92 DPS) | yes | Beads of Ogre Might (22150, -1.29 DPS) [quest]; Mark of Fordring (15411, -3.34 DPS) [quest]; Medallion of the Dawn (22659, -3.67 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 118.2 attack_power points (19.35 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, -2.71 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, -4.34 DPS) [pvp]; Wyrmhide Spaulders (12082, -4.95 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 75.1 attack_power points (12.29 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Arcanoweave Cloak (272411, -4.58 DPS) [vendor]; Stalwart Cloak (272415, -4.58 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmantle Tunic (226825, -3.56 DPS) [quest]; Field Marshal's Leather Chestpiece (231543, -3.82 DPS) [pvp]; Tunic of Undead Slaying (23089, -18.43 DPS, sim-verified) [world] |
| wrist | Blackmist Armguards (12966) | Blackrock Spire: The Beast [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.12 DPS) [rep]; Marshal's Leather Armsplints (16460, -0.29 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -2.72 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (556.5 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.27 DPS) [crafted]; Stormshroud Gloves (21278, -5.77 DPS, sim-verified) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 138.8 attack_power points (22.72 DPS) | yes | Belt of Preserved Heads (20216, -7.51 DPS) [quest]; Highlander's Leather Girdle (20045, -11.83 DPS) [rep]; Ferocity of the Timbermaw (227805, -12.56 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 122.5 attack_power points (20.04 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -1.44 DPS) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 84.0 attack_power points (13.74 DPS) | yes | Fine Dawn Treaders (227815, -3.13 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -4.13 DPS) [dungeon]; Darkmantle Boots (22003, -5.40 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.73 DPS) [dungeon]; Cutthroat's Signet (272408, -3.08 DPS) [vendor]; Naglering (11669, -9.05 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.04 DPS) [dungeon]; Cutthroat's Signet (272408, -1.39 DPS) [vendor]; Naglering (11669, -7.97 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -6.03 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -1.90 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -4.52 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 835.4 attack_power points (136.71 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -11.15 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -41.89 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.77 DPS) [dungeon]; Dark Iron Rifle (16004, -2.83 DPS, sim-verified) [crafted]; The Purifier (22656, -3.43 DPS) [quest] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Blackmist Armguards; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 34.0. Weights run: 1.6s. Verify run: 1.1s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.313 ± 0.007, crit=0.643 ± 0.012 per rating point (14 rating = 1%, 9.004 per %), hit=0.873 ± 0.032 per rating point (10 rating = 1%, 8.735 per %), melee_haste=5.171 ± 0.355

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 10.5 attack_power points (0.56 DPS) | yes | Defender's Leather Hood (252447, -0.13 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 7.9 attack_power points (0.42 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.6 attack_power points (0.35 DPS) | yes | Slime-encrusted Pads (6461, -0.41 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 7.9 attack_power points (0.42 DPS) | yes | Cape of the Brotherhood (5193, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 14.2 attack_power points (0.76 DPS) | yes | Defender's Leather Armor (252434, -0.20 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.23 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.27 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 6.6 attack_power points (0.35 DPS) | yes | Bristlebark Bindings (14569, -0.03 DPS) [world_drop]; Wolf Bracers (4794, -0.07 DPS) [vendor]; Ratchet Wristwraps (274742, -0.14 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 11.9 attack_power points (0.64 DPS) | yes | Brawler's Leather Gloves (252494, -0.14 DPS) [crafted]; Fletcher's Gloves (7348, -0.15 DPS) [crafted]; Bristlebark Gloves (14572, -0.17 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.97 DPS) | yes | Brawler's Leather Belt (252428, -0.47 DPS) [crafted]; Deviate Scale Belt (6468, -0.58 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.64 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 17.8 attack_power points (0.96 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.19 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 13.5 attack_power points (0.72 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.30 DPS) [dungeon]; Footpads of the Fang (10411, -0.30 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 9.3 attack_power points (0.50 DPS) | yes | Signet of the Zhevra (285330, -0.07 DPS) [world]; Demon Band (12054, -0.28 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 8.7 attack_power points (0.47 DPS) | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Demon Band (12054, -0.25 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.26 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.35 DPS) | yes | Blackfang (2236, -1.09 DPS) [world_drop]; Wingblade (6504, -1.22 DPS) [quest]; Diamond Hammer (2194, -2.39 DPS, sim-verified) [world_drop] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (34.0 DPS) | yes | Diamond Hammer (2194, -0.49 DPS, sim-verified) [world_drop]; Tork Wrench (11855, -12.20 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.3 attack_power points (0.28 DPS) | yes | Fine Longbow (11304, -0.07 DPS) [vendor]; Deadly Blunderbuss (4369, -0.14 DPS) [crafted]; Light Bow (4576, -0.14 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 60.4. Weights run: 1.6s. Verify run: 1.1s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.364 ± 0.009, crit=0.741 ± 0.015 per rating point (14 rating = 1%, 10.371 per %), hit=1.143 ± 0.043 per rating point (10 rating = 1%, 11.426 per %), melee_haste=6.763 ± 0.450

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 13.6 attack_power points (0.77 DPS) | yes | Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.15 DPS) [world]; Brawler's Leather Hood (252504, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.79 DPS) | yes | Scout's Medallion (19537, -0.18 DPS) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.0 attack_power points (1.13 DPS) | yes | Mantle of Thieves (2264, -0.45 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.46 DPS) [crafted]; Bristlebark Amice (14573, -0.50 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.5 attack_power points (0.71 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Wildhunter Cloak (16658, -0.14 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 19.1 attack_power points (1.08 DPS) | yes | Brawler's Leather Tunic (252508, -0.12 DPS) [crafted]; Brawler's Leather Armor (252490, -0.26 DPS) [crafted]; Panther Armor (6670, -0.27 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.2 attack_power points (0.69 DPS) | yes | Jurassic Wristguards (6198, -0.11 DPS) [world]; Cultist's Armguards (270032, -0.12 DPS) [quest]; Barbaric Bracers (18948, -0.15 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.91 DPS) | yes | Insignia Gloves (6408, -0.05 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Wolfclaw Gloves (1978, -0.16 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.36 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.32 DPS) [crafted]; Blackened Defias Belt (10403, -0.34 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.47 DPS) | yes | Petrolspill Leggings (9509, -0.39 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.39 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.48 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 13.9 attack_power points (0.79 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Insignia Boots (4055, -0.17 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.17 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 16.3 attack_power points (0.92 DPS) | yes | Thunderbrow Ring (13097, -0.24 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.35 DPS) [quest]; Monkey Ring (6748, -0.38 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.2 attack_power points (0.80 DPS) | yes | Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.24 DPS) [quest]; Monkey Ring (6748, -0.26 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (18.59 DPS) | yes | Swinetusk Shank (6691, -0.33 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.44 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.54 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (18.29 DPS) | yes | Swinetusk Shank (6691, -12.20 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -18.18 DPS) [quest]; Satyr's Rod (15962, -18.21 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.51 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Silver Star (3463, -0.12 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.20 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 32502110551501000-00000000000000000-0000000000000000000)

Set DPS (verified): 110.3. Weights run: 1.8s. Verify run: 1.1s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.369 ± 0.011, crit=0.767 ± 0.018 per rating point (14 rating = 1%, 10.741 per %), hit=2.206 ± 0.075 per rating point (10 rating = 1%, 22.058 per %), melee_haste=26.345 ± 0.714

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.1 attack_power points (0.89 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.28 DPS) [world_drop]; Nightscape Headband (8176, -0.33 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.69 DPS) | yes | Scout's Medallion (19536, -0.17 DPS) [rep]; Ghostshard Talisman (7731, -0.21 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.1 attack_power points (0.93 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Flintrock Shoulders (7755, -0.29 DPS) [dungeon]; Nightscape Shoulders (8192, -0.41 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.7 attack_power points (0.61 DPS) | yes | First Sergeant's Cloak (16340, -0.12 DPS) [pvp]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop]; Parachute Cloak (10518, -0.23 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 34.0 attack_power points (1.17 DPS) | yes | Wolffear Harness (13110, -0.38 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.46 DPS) [crafted]; Dusky Leather Armor (7374, -0.51 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.27 DPS) [world_drop]; Dusky Bracers (7378, -0.31 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.7 attack_power points (1.06 DPS) | yes | Prowler's Leather Gloves (252524, -0.29 DPS) [crafted]; Imperial Leather Gloves (4063, -0.32 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.33 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.03 DPS) | yes | Defiler's Chain Girdle (20152, -0.21 DPS) [rep]; Ogron's Sash (13117, -0.30 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.40 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 28.8 attack_power points (0.99 DPS) | yes | Triprunner Dungarees (9624, -0.04 DPS) [quest]; Ferine Leggings (6690, -0.09 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.32 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.1 attack_power points (0.76 DPS) | yes | Prowler's Leather Shoes (252465, -0.05 DPS) [crafted]; Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Excelsior Boots (4109, -0.09 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Legionnaire's Band (19512, -0.04 DPS) [rep]; Field Researcher's Loop (281634, -0.12 DPS) [quest]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.69 DPS) | yes | Legionnaire's Band (19512, -0.04 DPS) [rep]; Field Researcher's Loop (281634, -0.12 DPS) [quest]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.49 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.12 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (110.3 DPS) | yes | Stonecloth Branch (15963, -15.07 DPS) [world_drop]; Tork Wrench (11855, -15.10 DPS) [quest]; Ardent Custodian (868, -63.40 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.13 DPS) [dungeon]; Swiftwind (13038, -0.15 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.50 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 176.7. Weights run: 2.4s. Verify run: 1.5s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.467 ± 0.009, crit=0.987 ± 0.016 per rating point (14 rating = 1%, 13.812 per %), hit=2.105 ± 0.089 per rating point (10 rating = 1%, 21.053 per %), melee_haste=6.379 ± 1.375

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 49.8 attack_power points (4.01 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -1.31 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.82 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.1 attack_power points (1.94 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.33 DPS) [quest]; Woven Ivy Necklace (19159, -0.39 DPS) [quest]; Scout's Medallion (19535, -0.52 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.1 attack_power points (2.26 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.24 DPS) [crafted]; Failed Flying Experiment (9647, -0.28 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.5 attack_power points (2.14 DPS) | yes | Blisterbane Wrap (12552, -0.36 DPS) [dungeon]; Dark Phantom Cape (13122, -0.36 DPS) [world_drop]; Duskbat Drape (19982, -0.48 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 52.9 attack_power points (4.25 DPS) | yes | Blazewind Breastplate (11193, -1.30 DPS) [quest]; Fungus Shroud Armor (17742, -1.30 DPS) [dungeon]; Warbear Harness (15064, -2.14 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 26.0 attack_power points (2.09 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.47 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 42.3 attack_power points (3.40 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.68 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.02 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.06 DPS) | yes | Defiler's Leather Girdle (20193, -0.34 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.44 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.51 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 52.9 attack_power points (4.25 DPS) | yes | Basilisk Hide Pants (1718, -1.77 DPS) [world_drop]; Triprunner Dungarees (9624, -1.89 DPS) [quest]; Serpentskin Leggings (8262, -2.26 DPS, sim-verified) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 30.9 attack_power points (2.49 DPS) | yes | Skulker's Leather Boots (252469, -0.07 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.15 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.1 attack_power points (3.30 DPS) | yes | Legionnaire's Band (19511, -1.44 DPS) [rep]; Masons Fraternity Ring (9533, -1.65 DPS) [quest]; Mark of Kern (2262, -1.69 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.93 DPS) | yes | Legionnaire's Band (19511, -0.06 DPS) [rep]; Masons Fraternity Ring (9533, -0.28 DPS) [quest]; Mark of Kern (2262, -0.32 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (176.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (176.7 DPS) | yes | Molten Heart of the Mountain (249470, -1.93 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (176.7 DPS) | yes | Thorium Cestus (250614, -2.16 DPS) [crafted]; Bloodrazor (809, -2.28 DPS, sim-verified) [world_drop]; Doomforged Straightedge (12535, -2.43 DPS) [dungeon] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (43.86 DPS) | yes | Thorium Cestus (250614, +0.00 DPS) [crafted]; Claw of Celebras (17738, -5.06 DPS) [dungeon]; White Bone Shredder (11863, -7.71 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (176.7 DPS) | yes | Stinging Bow (10624, -0.28 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.28 DPS) [world_drop]; Dark Iron Rifle (16004, -2.07 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 275.4. Weights run: 2.3s. Verify run: 1.5s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.055 ± 0.022, crit=2.233 ± 0.038 per rating point (14 rating = 1%, 31.256 per %), hit=4.760 ± 0.173 per rating point (10 rating = 1%, 47.597 per %), melee_haste=16.275 ± 2.607

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 126.4 attack_power points (11.31 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted]; Champion's Leather Helm (227057, -1.04 DPS) [pvp] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 78.4 attack_power points (7.01 DPS) | yes | Beads of Ogre Might (22150, -0.61 DPS) [quest]; Mark of Fordring (15411, -1.89 DPS) [quest]; Medallion of the Dawn (22659, -2.07 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 119.2 attack_power points (10.66 DPS) | yes | Champion's Leather Shoulders (227056, -1.64 DPS) [pvp]; Warlord's Leather Spaulders (231551, -2.54 DPS) [pvp]; Wyrmhide Spaulders (12082, -2.62 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 75.6 attack_power points (6.76 DPS) | yes | Cape of the Black Baron (13340, -2.21 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.50 DPS) [vendor]; Stalwart Cloak (272415, -2.50 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (275.4 DPS) | yes | Darkmantle Tunic (226825, -1.77 DPS) [quest]; Warlord's Leather Breastplate (231549, -2.02 DPS) [pvp]; Tunic of Undead Slaying (23089, -10.21 DPS, sim-verified) [world] |
| wrist | Blackmist Armguards (12966) | Blackrock Spire: The Beast [dungeon] | sim-verified (275.4 DPS) | yes | Forest Stalker's Bracers (19587, -0.23 DPS) [rep]; General's Leather Armsplints (16559, -0.32 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -2.11 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 78.9 attack_power points (7.05 DPS) | yes | Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; General's Leather Mitts (231555, -0.58 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.75 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 138.4 attack_power points (12.37 DPS) | yes | Belt of Preserved Heads (20216, -1.13 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -6.54 DPS) [rep]; Ferocity of the Timbermaw (227805, -6.92 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 118.0 attack_power points (10.55 DPS) | yes | General's Leather Legguards (231554, +0.00 DPS) [pvp]; Legionnaire's Leather Legguards (227059, -0.46 DPS) [pvp]; Plaguehound Leggings (18736, -2.40 DPS, sim-verified) [dungeon] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 82.6 attack_power points (7.39 DPS) | yes | Fine Dawn Treaders (227815, -1.26 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -2.25 DPS) [dungeon]; Darkmantle Boots (22003, -2.98 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (275.4 DPS) | yes | Tarnished Elven Ring (18500, -1.47 DPS) [dungeon]; Cutthroat's Signet (272408, -1.65 DPS) [vendor]; Naglering (11669, -5.34 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (275.4 DPS) | yes | Tarnished Elven Ring (18500, -0.55 DPS) [dungeon]; Cutthroat's Signet (272408, -0.74 DPS) [vendor]; Naglering (11669, -4.83 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (275.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (275.4 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -1.46 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -2.90 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (275.4 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 834.8 attack_power points (74.64 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -6.68 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -22.83 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (275.4 DPS) | yes | Blackcrow (12651, -0.43 DPS) [dungeon]; The Purifier (22656, -2.01 DPS) [quest]; Dark Iron Rifle (16004, -2.08 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Blackmist Armguards; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 550.3. Weights run: 2.3s. Verify run: 1.5s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.124 ± 0.025, crit=2.326 ± 0.042 per rating point (14 rating = 1%, 32.558 per %), hit=4.712 ± 0.160 per rating point (10 rating = 1%, 47.119 per %), melee_haste=18.540 ± 2.831

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 126.8 attack_power points (20.75 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted]; Champion's Leather Helm (227057, -1.82 DPS) [pvp] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 79.0 attack_power points (12.92 DPS) | yes | Beads of Ogre Might (22150, -1.29 DPS) [quest]; Mark of Fordring (15411, -3.34 DPS) [quest]; Medallion of the Dawn (22659, -3.67 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 118.2 attack_power points (19.35 DPS) | yes | Champion's Leather Shoulders (227056, -2.71 DPS) [pvp]; Warlord's Leather Spaulders (231551, -4.34 DPS) [pvp]; Wyrmhide Spaulders (12082, -4.90 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 75.1 attack_power points (12.29 DPS) | yes | Cape of the Black Baron (13340, -3.81 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.58 DPS) [vendor]; Stalwart Cloak (272415, -4.58 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (550.3 DPS) | yes | Darkmantle Tunic (226825, -3.56 DPS) [quest]; Warlord's Leather Breastplate (231549, -3.82 DPS) [pvp]; Tunic of Undead Slaying (23089, -21.90 DPS, sim-verified) [world] |
| wrist | Blackmist Armguards (12966) | Blackrock Spire: The Beast [dungeon] | sim-verified (550.3 DPS) | yes | Forest Stalker's Bracers (19587, -0.12 DPS) [rep]; General's Leather Armsplints (16559, -0.29 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -5.36 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 79.7 attack_power points (13.04 DPS) | yes | Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; General's Leather Mitts (231555, -0.76 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.13 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 138.8 attack_power points (22.72 DPS) | yes | Belt of Preserved Heads (20216, -3.52 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -11.83 DPS) [rep]; Ferocity of the Timbermaw (227805, -12.56 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 122.5 attack_power points (20.04 DPS) | yes | General's Leather Legguards (231554, +0.00 DPS) [pvp]; Legionnaire's Leather Legguards (227059, -1.44 DPS) [pvp]; Plaguehound Leggings (18736, -5.53 DPS, sim-verified) [dungeon] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 84.0 attack_power points (13.74 DPS) | yes | Fine Dawn Treaders (227815, -3.07 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -4.13 DPS) [dungeon]; Darkmantle Boots (22003, -5.40 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (550.3 DPS) | yes | Tarnished Elven Ring (18500, -2.73 DPS) [dungeon]; Cutthroat's Signet (272408, -3.08 DPS) [vendor]; Naglering (11669, -12.06 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (550.3 DPS) | yes | Tarnished Elven Ring (18500, -1.04 DPS) [dungeon]; Cutthroat's Signet (272408, -1.39 DPS) [vendor]; Naglering (11669, -10.62 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (550.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Hand of Justice (11815, -3.85 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (550.3 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -3.98 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -5.33 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (550.3 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -5.03 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 835.4 attack_power points (136.71 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -12.77 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -41.89 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (550.3 DPS) | yes | Blackcrow (12651, -0.77 DPS) [dungeon]; The Purifier (22656, -3.43 DPS) [quest]; Dark Iron Rifle (16004, -5.33 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Blackmist Armguards; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

