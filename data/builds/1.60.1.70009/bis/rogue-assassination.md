# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 36.2. Weights run: 1.4s. Verify run: 1.4s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.067 ± 0.020, crit=0.159 ± 0.008 per rating point (14 rating = 1%, 2.221 per %), hit=0.061 ± 0.003 per rating point (10 rating = 1%, 0.606 per %), melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 attack_power points (0.40 DPS) | yes | Defender's Leather Hood (252447, -0.02 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.4 attack_power points (0.30 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 attack_power points (0.25 DPS) | yes | Slime-encrusted Pads (6461, -0.40 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Catacomb Cloak (279899, -0.02 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.5 attack_power points (0.58 DPS) | yes | Tunic of Westfall (2041, -0.03 DPS) [quest]; Defender's Leather Armor (252434, -0.11 DPS) [crafted]; Prospector's Chestpiece (14562, -0.19 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.3 attack_power points (0.29 DPS) | yes | Forest Leather Bracers (3202, -0.04 DPS) [world_drop]; Bristlebark Bindings (14569, -0.05 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.4 attack_power points (0.49 DPS) | yes | Brawler's Leather Gloves (252494, -0.10 DPS) [crafted]; Gold-flecked Gloves (5195, -0.16 DPS) [dungeon]; Bristlebark Gloves (14572, -0.17 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.84 DPS) | yes | Brawler's Leather Belt (252428, -0.45 DPS) [crafted]; Ruffian Belt (5975, -0.56 DPS) [world]; Deviate Scale Belt (6468, -0.58 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.6 attack_power points (0.73 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.11 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.5 attack_power points (0.54 DPS) | yes | Brawler's Leather Boots (252439, -0.06 DPS) [crafted]; Blackened Defias Boots (10402, -0.24 DPS) [dungeon]; Footpads of the Fang (10411, -0.24 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 attack_power points (0.39 DPS) | yes | Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Demon Band (12054, -0.20 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 attack_power points (0.30 DPS) | yes | Pyrewood Signet Ring (277210, -0.09 DPS) [quest]; Demon Band (12054, -0.11 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.61 DPS) | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Barrens Basher (274744, -1.10 DPS) [vendor]; Diamond Hammer (2194, -2.82 DPS, sim-verified) [world_drop] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (36.2 DPS) | yes | Diamond Hammer (2194, -0.40 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.3 attack_power points (0.20 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 52.2. Weights run: 1.4s. Verify run: 1.4s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.097 ± 0.016, crit=0.236 ± 0.009 per rating point (14 rating = 1%, 3.306 per %), hit=0.077 ± 0.004 per rating point (10 rating = 1%, 0.769 per %), melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Helm (252512, -0.05 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.66 DPS) | yes | Sentinel's Medallion (19541, -0.20 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 17.1 attack_power points (0.80 DPS) | yes | Barbaric Shoulders (5964, -0.31 DPS) [crafted]; Bristlebark Amice (14573, -0.35 DPS) [world_drop]; Mantle of Thieves (2264, -0.41 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.7 attack_power points (0.50 DPS) | yes | Wolfmaster Cape (6314, -0.03 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.11 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.75 DPS) | yes | Dusky Leather Armor (7374, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Tunic (252508, -0.06 DPS) [crafted]; Brawler's Leather Armor (252490, -0.16 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.6 attack_power points (0.50 DPS) | yes | Cultist's Armguards (270032, -0.03 DPS) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.75 DPS) | yes | Insignia Gloves (6408, -0.11 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.16 DPS) [crafted]; Wolfclaw Gloves (1978, -0.21 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.13 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.38 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.23 DPS) | yes | Brawler's Leather Legguards (252516, -0.44 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.07 DPS) [crafted]; Insignia Boots (4055, -0.15 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.9 attack_power points (0.65 DPS) | yes | Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Insurgent's Band (272067, -0.23 DPS) [vendor]; Monkey Ring (6748, -0.29 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.6 attack_power points (0.59 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Insurgent's Band (272067, -0.17 DPS) [vendor]; Monkey Ring (6748, -0.23 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (15.46 DPS) | yes | Swinetusk Shank (6691, -0.28 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.37 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.45 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.21 DPS) | yes | Swinetusk Shank (6691, -13.77 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.16 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Double-barreled Shotgun (2098, -0.13 DPS) [world_drop]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 89.7. Weights run: 1.4s. Verify run: 1.5s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.105 ± 0.014, crit=0.198 ± 0.005 per rating point (14 rating = 1%, 2.769 per %), hit=0.055 ± 0.002 per rating point (10 rating = 1%, 0.555 per %), melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 23.2 attack_power points (0.78 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.29 DPS) [world_drop]; Nightscape Headband (8176, -0.33 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Ghostshard Talisman (7731, -0.20 DPS) [dungeon]; Sentinel's Medallion (19540, -0.26 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.81 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.27 DPS) [dungeon]; Nightscape Shoulders (8192, -0.40 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 15.0 attack_power points (0.50 DPS) | yes | Sergeant Major's Cape (16336, -0.08 DPS) [pvp]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Wolfmaster Cape (6314, -0.17 DPS) [dungeon] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 29.0 attack_power points (0.97 DPS) | yes | Wolffear Harness (13110, -0.38 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.42 DPS) [crafted]; Raptorbane Armor (3566, -0.43 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Hawkeye's Bracers (14590, -0.25 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.33 DPS) [quest]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 22.8 attack_power points (0.76 DPS) | yes | Prowler's Leather Gloves (252524, -0.09 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.30 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.00 DPS) | yes | Highlander's Chain Girdle (20090, -0.20 DPS) [rep]; Ogron's Sash (13117, -0.37 DPS) [world_drop]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.87 DPS) | yes | Basilisk Hide Pants (1718, -0.09 DPS) [world_drop]; Triprunner Dungarees (9624, -0.10 DPS) [quest]; Brawler's Leather Legguards (252516, -0.30 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 19.2 attack_power points (0.64 DPS) | yes | Prowler's Leather Shoes (252465, -0.01 DPS) [crafted]; Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Excelsior Boots (4109, -0.07 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Protector's Band (19515, -0.11 DPS) [rep]; Field Researcher's Loop (281634, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.20 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.67 DPS) | yes | Protector's Band (19515, -0.11 DPS) [rep]; Field Researcher's Loop (281634, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (15.89 DPS) | yes | Ardent Custodian (868, -0.48 DPS) [world_drop]; Dazzling Longsword (869, -1.13 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.17 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (89.7 DPS) | yes | Stonecloth Branch (15963, -14.70 DPS) [world_drop]; Satyr's Rod (15962, -14.77 DPS) [world_drop]; Ardent Custodian (868, -51.02 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.16 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.17 DPS) [vendor]; Bow of Searing Arrows (2825, -0.49 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 124.1. Weights run: 1.5s. Verify run: 1.9s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.107 ± 0.016, crit=0.234 ± 0.006 per rating point (14 rating = 1%, 3.272 per %), hit=0.068 ± 0.003 per rating point (10 rating = 1%, 0.678 per %), melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (1.35 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -0.57 DPS) [crafted]; Knight-Lieutenant's Leather Headband (220850, -0.68 DPS) [vendor] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Skibi's Pendant (13089, -0.02 DPS) [world_drop]; Ghostshard Talisman (7731, -0.20 DPS) [dungeon]; Sentinel's Medallion (19539, -0.23 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.82 DPS) | yes | Skulker's Leather Shoulder (252535, -0.10 DPS) [crafted]; Failed Flying Experiment (9647, -0.10 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.11 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.5 attack_power points (0.73 DPS) | yes | Blisterbane Wrap (12552, -0.17 DPS) [dungeon]; Dark Phantom Cape (13122, -0.17 DPS) [world_drop]; Duskbat Drape (19982, -0.20 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 30.9 attack_power points (1.04 DPS) | yes | Mixologist's Tunic (12793, -0.03 DPS) [dungeon]; Quillward Harness (10583, -0.06 DPS) [dungeon]; Blazewind Breastplate (11193, -0.08 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 20.6 attack_power points (0.70 DPS) | yes | Skulker's Leather Bracers (252540, -0.12 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.13 DPS) [crafted]; Branded Leather Bracers (19508, -0.57 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 34.4 attack_power points (1.16 DPS) | yes | Darkmantle Grips (226828, -0.03 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -0.33 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -0.35 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (1.28 DPS) | yes | Highlander's Leather Girdle (20116, -0.27 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.35 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.36 DPS) [crafted] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 31.6 attack_power points (1.07 DPS) | yes | Serpentskin Leggings (8262, -0.03 DPS) [world_drop]; Ferine Leggings (6690, -0.19 DPS) [dungeon]; Basilisk Hide Pants (1718, -0.28 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 25.4 attack_power points (0.86 DPS) | yes | Prowler's Leather Boots (252468, -0.01 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.02 DPS) [dungeon]; Shadefiend Boots (11675, -0.11 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.7 attack_power points (0.70 DPS) | yes | Assault Band (13095, -0.02 DPS) [world_drop]; Protector's Band (19516, -0.02 DPS) [rep] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Protector's Band (19516, -0.00 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (124.1 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (124.1 DPS) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted] |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | sim-verified (124.1 DPS) | yes | Doomforged Straightedge (12535, +0.00 DPS) [dungeon]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hammer of the Northern Wind (810, -1.70 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (18.39 DPS) | yes | Claw of Celebras (17738, -2.12 DPS) [dungeon]; Thorium Cestus (250614, -4.49 DPS, sim-verified) [crafted]; Thermotastic Egg Timer (9644, -18.28 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (124.1 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.05 DPS) [world_drop]; Dark Iron Rifle (16004, -1.04 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Mark of Kern; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hanzo Sword; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 188.8. Weights run: 1.5s. Verify run: 1.8s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.131 ± 0.020, crit=0.295 ± 0.007 per rating point (14 rating = 1%, 4.124 per %), hit=0.086 ± 0.004 per rating point (10 rating = 1%, 0.863 per %), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (1.33 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, -0.15 DPS) [pvp] |
| neck | Imperial Jewel (11933) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 32.0 attack_power points (1.06 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS, sim-verified) [quest]; Will of the Martyr (17044, -0.07 DPS) [quest]; Medallion of the Dawn (22659, -0.13 DPS) [quest] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 50.4 attack_power points (1.67 DPS) | yes | Highlander's Lizardhide Shoulders (20060, -0.22 DPS) [rep]; Darkspear Pauldrons (272105, -0.52 DPS) [vendor]; Dark Warder's Pauldrons (22241, -0.67 DPS) [dungeon] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 39.7 attack_power points (1.31 DPS) | yes | Cape of the Black Baron (13340, -0.09 DPS) [dungeon]; Howler's Furs (272414, -0.36 DPS) [vendor]; Windshear Cape (20691, -0.49 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (188.8 DPS) | yes | Timbermaw Tunic (252484, -0.65 DPS) [crafted]; Nightbrace Tunic (12603, -0.71 DPS) [dungeon]; Tunic of Undead Slaying (23089, -5.74 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (188.8 DPS) | yes | Forest Stalker's Bracers (19587, -0.09 DPS) [rep]; Marshal's Leather Armsplints (16460, -0.13 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -2.72 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | 44.0 attack_power points (1.46 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.02 DPS) [pvp]; Raider Gloves (272099, -0.08 DPS) [vendor]; Skul's Fingerbone Claws (13395, -0.13 DPS) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 45.2 attack_power points (1.50 DPS) | yes | Marshal's Leather Cinch (16458, -0.07 DPS) [pvp]; Highlander's Leather Girdle (20045, -0.24 DPS) [rep]; Cadaverous Belt (14636, -1.34 DPS, sim-verified) [dungeon] |
| legs | Cadaverous Leggings (14638) | Scholomance: Lady Illucia Barov [dungeon] | 52.0 attack_power points (1.72 DPS) | yes | Devilsaur Leggings (15062, +0.00 DPS, sim-verified) [crafted]; Warbear Woolies (15065, -0.12 DPS) [crafted]; Knight-Captain's Leather Legguards (23299, -0.43 DPS) [vendor] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (1.33 DPS) | yes | Drudge Boots (21532, -0.26 DPS) [quest]; Dunestalker's Boots (20715, -0.32 DPS) [quest]; Highlander's Leather Boots (20052, -0.35 DPS) [rep] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (188.8 DPS) | yes | Don Julio's Band (19325, -0.11 DPS) [rep]; Blackstone Ring (17713, -0.12 DPS) [dungeon]; Naglering (11669, -2.13 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (188.8 DPS) | yes | Don Julio's Band (19325, -0.01 DPS) [rep]; Blackstone Ring (17713, -0.01 DPS) [dungeon]; Naglering (11669, -4.65 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (188.8 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-verified (188.8 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, -3.62 DPS, sim-verified) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (188.8 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Alcor's Sunrazor (14555, -1.14 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 826.5 attack_power points (27.40 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -1.88 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -8.19 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (188.8 DPS) | yes | Malgen's Long Bow (22318, -0.07 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -1.34 DPS, sim-verified) [crafted] |

**New at 60:** neck: Imperial Jewel; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Cadaverous Gloves; waist: Ferocity of the Timbermaw; legs: Cadaverous Leggings; feet: Pads of the Dread Wolf; finger1: Protector's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Royal Seal of Eldre'Thalas; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.3. Weights run: 1.4s. Verify run: 1.3s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.067 ± 0.020, crit=0.159 ± 0.008 per rating point (14 rating = 1%, 2.221 per %), hit=0.061 ± 0.003 per rating point (10 rating = 1%, 0.606 per %), melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 attack_power points (0.40 DPS) | yes | Defender's Leather Hood (252447, -0.02 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.4 attack_power points (0.30 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 attack_power points (0.25 DPS) | yes | Slime-encrusted Pads (6461, -0.40 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Catacomb Cloak (279899, -0.02 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.5 attack_power points (0.58 DPS) | yes | Defender's Leather Armor (252434, -0.19 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.19 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.20 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.3 attack_power points (0.25 DPS) | yes | Bristlebark Bindings (14569, -0.01 DPS) [world_drop]; Wolf Bracers (4794, -0.05 DPS) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.4 attack_power points (0.49 DPS) | yes | Brawler's Leather Gloves (252494, -0.10 DPS) [crafted]; Bristlebark Gloves (14572, -0.16 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.16 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.84 DPS) | yes | Brawler's Leather Belt (252428, -0.45 DPS) [crafted]; Ruffian Belt (5975, -0.56 DPS) [world]; Deviate Scale Belt (6468, -0.57 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.6 attack_power points (0.73 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.11 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.5 attack_power points (0.54 DPS) | yes | Brawler's Leather Boots (252439, -0.06 DPS) [crafted]; Blackened Defias Boots (10402, -0.24 DPS) [dungeon]; Footpads of the Fang (10411, -0.24 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 attack_power points (0.39 DPS) | yes | Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Demon Band (12054, -0.20 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.24 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 attack_power points (0.30 DPS) | yes | Pyrewood Signet Ring (277210, -0.09 DPS) [quest]; Demon Band (12054, -0.11 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.61 DPS) | yes | Cruel Barb (5191, -0.90 DPS) [dungeon]; Blackfang (2236, -0.95 DPS) [world_drop]; Barrens Basher (274744, -1.10 DPS) [vendor] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 229.8 attack_power points (10.72 DPS) | yes | Cruel Barb (5191, +0.00 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -10.63 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.3 attack_power points (0.20 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Diamond Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 51.9. Weights run: 1.4s. Verify run: 1.4s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.097 ± 0.016, crit=0.236 ± 0.009 per rating point (14 rating = 1%, 3.306 per %), hit=0.077 ± 0.004 per rating point (10 rating = 1%, 0.769 per %), melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Helm (252512, -0.05 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.66 DPS) | yes | Scout's Medallion (19537, -0.20 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 17.1 attack_power points (0.80 DPS) | yes | Barbaric Shoulders (5964, -0.31 DPS) [crafted]; Bristlebark Amice (14573, -0.35 DPS) [world_drop]; Mantle of Thieves (2264, -0.42 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.7 attack_power points (0.50 DPS) | yes | Wolfmaster Cape (6314, -0.03 DPS) [dungeon]; Wildhunter Cloak (16658, -0.03 DPS) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 15.4 attack_power points (0.72 DPS) | yes | Brawler's Leather Tunic (252508, -0.03 DPS) [crafted]; Brawler's Leather Armor (252490, -0.13 DPS) [crafted]; Defender's Leather Tunic (252450, -0.13 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.6 attack_power points (0.50 DPS) | yes | Cultist's Armguards (270032, -0.03 DPS) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.75 DPS) | yes | Insignia Gloves (6408, -0.11 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.16 DPS) [crafted]; Wolfclaw Gloves (1978, -0.21 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.13 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.36 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.23 DPS) | yes | Brawler's Leather Legguards (252516, -0.44 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.8 attack_power points (0.56 DPS) | yes | Brawler's Leather Boots (252439, -0.06 DPS) [crafted]; Insignia Boots (4055, -0.14 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.14 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.9 attack_power points (0.65 DPS) | yes | Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Insurgent's Band (272067, -0.23 DPS) [vendor]; Band of the Fist (17694, -0.26 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.6 attack_power points (0.59 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Insurgent's Band (272067, -0.17 DPS) [vendor]; Band of the Fist (17694, -0.20 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (15.46 DPS) | yes | Swinetusk Shank (6691, -0.28 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.37 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.45 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.21 DPS) | yes | Swinetusk Shank (6691, -13.86 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -15.12 DPS) [quest]; Satyr's Rod (15962, -15.16 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Double-barreled Shotgun (2098, -0.13 DPS) [world_drop]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 87.7. Weights run: 1.4s. Verify run: 1.4s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.105 ± 0.014, crit=0.198 ± 0.005 per rating point (14 rating = 1%, 2.769 per %), hit=0.055 ± 0.002 per rating point (10 rating = 1%, 0.555 per %), melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 23.2 attack_power points (0.78 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.29 DPS) [world_drop]; Nightscape Headband (8176, -0.33 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Ghostshard Talisman (7731, -0.21 DPS, sim-verified) [dungeon]; Scout's Medallion (19536, -0.26 DPS) [rep]; Ethereal Talisman (4430, -0.35 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.81 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.27 DPS) [dungeon]; Nightscape Shoulders (8192, -0.40 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 15.0 attack_power points (0.50 DPS) | yes | First Sergeant's Cloak (16340, -0.08 DPS) [pvp]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Wildhunter Cloak (16658, -0.17 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 29.0 attack_power points (0.97 DPS) | yes | Wolffear Harness (13110, -0.36 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.42 DPS) [crafted]; Barbaric Harness (5739, -0.44 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.26 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.33 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 22.8 attack_power points (0.76 DPS) | yes | Prowler's Leather Gloves (252524, -0.09 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.31 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.00 DPS) | yes | Defiler's Chain Girdle (20152, -0.21 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.37 DPS) [world_drop]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.87 DPS) | yes | Basilisk Hide Pants (1718, -0.09 DPS) [world_drop]; Triprunner Dungarees (9624, -0.10 DPS) [quest]; Brawler's Leather Legguards (252516, -0.30 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 19.2 attack_power points (0.64 DPS) | yes | Prowler's Leather Shoes (252465, -0.01 DPS) [crafted]; Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Excelsior Boots (4109, -0.07 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Legionnaire's Band (19512, -0.11 DPS) [rep]; Field Researcher's Loop (281634, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.20 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.67 DPS) | yes | Legionnaire's Band (19512, -0.11 DPS) [rep]; Field Researcher's Loop (281634, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (15.89 DPS) | yes | Ardent Custodian (868, -0.48 DPS) [world_drop]; Dazzling Longsword (869, -1.13 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.17 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (87.7 DPS) | yes | Stonecloth Branch (15963, -14.70 DPS) [world_drop]; Tork Wrench (11855, -14.74 DPS) [quest]; Ardent Custodian (868, -49.28 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.16 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.17 DPS) [vendor]; Bow of Searing Arrows (2825, -0.48 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 126.7. Weights run: 1.5s. Verify run: 1.9s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.107 ± 0.016, crit=0.234 ± 0.006 per rating point (14 rating = 1%, 3.272 per %), hit=0.068 ± 0.003 per rating point (10 rating = 1%, 0.678 per %), melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (1.35 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -0.57 DPS) [crafted]; Undercity Reservist's Cap (20643, -0.60 DPS) [quest] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Skibi's Pendant (13089, -0.02 DPS) [world_drop]; Woven Ivy Necklace (19159, -0.14 DPS) [quest]; Ghostshard Talisman (7731, -0.20 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.82 DPS) | yes | Skulker's Leather Shoulder (252535, -0.10 DPS) [crafted]; Failed Flying Experiment (9647, -0.10 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.11 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.5 attack_power points (0.73 DPS) | yes | Blisterbane Wrap (12552, -0.17 DPS) [dungeon]; Dark Phantom Cape (13122, -0.17 DPS) [world_drop]; Duskbat Drape (19982, -0.20 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 30.9 attack_power points (1.04 DPS) | yes | Mixologist's Tunic (12793, -0.03 DPS) [dungeon]; Quillward Harness (10583, -0.06 DPS) [dungeon]; Blazewind Breastplate (11193, -0.08 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 20.6 attack_power points (0.70 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.60 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 34.4 attack_power points (1.16 DPS) | yes | Darkmantle Grips (226828, -0.03 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -0.33 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -0.35 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (1.28 DPS) | yes | Defiler's Leather Girdle (20192, -0.27 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.35 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.36 DPS) [crafted] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 30.7 attack_power points (1.04 DPS) | yes | Basilisk Hide Pants (1718, -0.25 DPS) [world_drop]; Triprunner Dungarees (9624, -0.26 DPS) [quest]; Ferine Leggings (6690, -0.89 DPS, sim-verified) [dungeon] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 25.4 attack_power points (0.86 DPS) | yes | Prowler's Leather Boots (252468, -0.01 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.02 DPS) [dungeon]; Shadefiend Boots (11675, -0.11 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (0.81 DPS) | yes | Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop]; Legionnaire's Band (19511, -0.14 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.7 attack_power points (0.70 DPS) | yes | Assault Band (13095, -0.02 DPS) [world_drop]; Legionnaire's Band (19511, -0.02 DPS) [rep]; Mark of Kern (2262, -2.86 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (126.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (126.7 DPS) | yes | Molten Heart of the Mountain (249470, -2.35 DPS, sim-verified) [crafted] |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | sim-verified (126.7 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hookfang Shanker (11635, -2.02 DPS, sim-verified) [dungeon] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (18.39 DPS) | yes | Claw of Celebras (17738, -2.12 DPS) [dungeon]; White Bone Shredder (11863, -3.32 DPS) [quest]; Thorium Cestus (250614, -5.00 DPS, sim-verified) [crafted] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (126.7 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.05 DPS) [world_drop]; Dark Iron Rifle (16004, -1.04 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hanzo Sword; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 182.4. Weights run: 1.5s. Verify run: 1.9s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.131 ± 0.020, crit=0.295 ± 0.007 per rating point (14 rating = 1%, 4.124 per %), hit=0.086 ± 0.004 per rating point (10 rating = 1%, 0.863 per %), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (1.33 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, -0.15 DPS) [pvp] |
| neck | Imperial Jewel (11933) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 32.0 attack_power points (1.06 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS, sim-verified) [quest]; Will of the Martyr (17044, -0.07 DPS) [quest]; Medallion of the Dawn (22659, -0.13 DPS) [quest] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 50.4 attack_power points (1.67 DPS) | yes | Defiler's Lizardhide Shoulders (20175, -0.22 DPS) [rep]; Darkspear Pauldrons (272105, -0.52 DPS) [vendor]; Dark Warder's Pauldrons (22241, -0.67 DPS) [dungeon] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 39.7 attack_power points (1.31 DPS) | yes | Cape of the Black Baron (13340, -0.09 DPS) [dungeon]; Howler's Furs (272414, -0.36 DPS) [vendor]; Windshear Cape (20691, -0.49 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (182.4 DPS) | yes | Timbermaw Tunic (252484, -0.65 DPS) [crafted]; Nightbrace Tunic (12603, -0.71 DPS) [dungeon]; Tunic of Undead Slaying (23089, -5.52 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (182.4 DPS) | yes | Forest Stalker's Bracers (19587, -0.09 DPS) [rep]; General's Leather Armsplints (16559, -0.13 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -2.61 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | 44.0 attack_power points (1.46 DPS) | yes | Blood Guard's Leather Vices (16499, -0.02 DPS) [pvp]; Raider Gloves (272099, -0.08 DPS) [vendor]; Skul's Fingerbone Claws (13395, -0.13 DPS) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 45.2 attack_power points (1.50 DPS) | yes | General's Leather Girdle (16557, -0.07 DPS) [pvp]; Defiler's Leather Girdle (20190, -0.24 DPS) [rep]; Cadaverous Belt (14636, -1.26 DPS, sim-verified) [dungeon] |
| legs | Cadaverous Leggings (14638) | Scholomance: Lady Illucia Barov [dungeon] | 52.0 attack_power points (1.72 DPS) | yes | Devilsaur Leggings (15062, -0.06 DPS) [crafted]; Warbear Woolies (15065, -0.12 DPS) [crafted]; Legionnaire's Leather Legguards (227059, -0.43 DPS) [pvp] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (1.33 DPS) | yes | Drudge Boots (21532, -0.26 DPS) [quest]; Dunestalker's Boots (20715, -0.32 DPS) [quest]; Defiler's Leather Boots (20186, -0.35 DPS) [rep] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (182.4 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.11 DPS) [quest]; Don Julio's Band (19325, -0.11 DPS) [rep]; Naglering (11669, -2.04 DPS, sim-verified) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | sim-verified (182.4 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.09 DPS) [quest]; Don Julio's Band (19325, -0.10 DPS) [rep]; Naglering (11669, -1.41 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (182.4 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -3.02 DPS, sim-verified) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-verified (182.4 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Smolderweb's Eye (13213, -1.02 DPS, sim-verified) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (182.4 DPS) | yes | Alcor's Sunrazor (14555, +0.00 DPS) [world_drop]; High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 826.5 attack_power points (27.40 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -2.39 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -8.19 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (182.4 DPS) | yes | Malgen's Long Bow (22318, -0.07 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -1.29 DPS, sim-verified) [crafted] |

**New at 60:** neck: Imperial Jewel; shoulder: Defiler's Leather Shoulders; back: Deathguard's Cloak; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Cadaverous Gloves; waist: Ferocity of the Timbermaw; legs: Cadaverous Leggings; feet: Pads of the Dread Wolf; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Royal Seal of Eldre'Thalas; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

