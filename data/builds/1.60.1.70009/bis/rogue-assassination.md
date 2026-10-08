# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 34.1. Weights run: 2.1s. Verify run: 1.4s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.307 ± 0.007, crit=0.631 ± 0.012 per rating point (14 rating = 1%, 8.831 per %), hit=0.874 ± 0.032 per rating point (10 rating = 1%, 8.740 per %), melee_haste=5.172 ± 0.354

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 10.5 attack_power points (0.55 DPS) | yes | Defender's Leather Hood (252447, -0.13 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 7.8 attack_power points (0.42 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.5 attack_power points (0.35 DPS) | yes | Slime-encrusted Pads (6461, -0.41 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 7.8 attack_power points (0.42 DPS) | yes | Cape of the Brotherhood (5193, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 14.4 attack_power points (0.76 DPS) | yes | Brawler's Leather Armor (252490, -0.01 DPS) [crafted]; Defender's Leather Armor (252434, -0.18 DPS) [crafted]; Prospector's Chestpiece (14562, -0.24 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 7.2 attack_power points (0.38 DPS) | yes | Forest Leather Bracers (3202, -0.04 DPS) [world_drop]; Bristlebark Bindings (14569, -0.07 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 11.8 attack_power points (0.63 DPS) | yes | Brawler's Leather Gloves (252494, -0.14 DPS) [crafted]; Fletcher's Gloves (7348, -0.16 DPS) [crafted]; Bristlebark Gloves (14572, -0.16 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.96 DPS) | yes | Brawler's Leather Belt (252428, -0.47 DPS) [crafted]; Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.61 DPS) [quest] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 17.8 attack_power points (0.94 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.19 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 13.5 attack_power points (0.71 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.30 DPS) [dungeon]; Footpads of the Fang (10411, -0.30 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 9.2 attack_power points (0.49 DPS) | yes | Signet of the Zhevra (285330, -0.07 DPS) [world]; Demon Band (12054, -0.28 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 8.7 attack_power points (0.46 DPS) | yes | Signet of the Zhevra (285330, -0.05 DPS) [world]; Demon Band (12054, -0.25 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.21 DPS) | yes | Blackfang (2236, -1.08 DPS) [world_drop]; Assassin's Blade (1935, -1.23 DPS) [dungeon]; Diamond Hammer (2194, -2.62 DPS, sim-verified) [world_drop] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (34.1 DPS) | yes | Diamond Hammer (2194, -0.41 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.2 attack_power points (0.28 DPS) | yes | Fine Longbow (11304, -0.07 DPS) [vendor]; Deadly Blunderbuss (4369, -0.14 DPS) [crafted]; Light Bow (4576, -0.14 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 59.7. Weights run: 2.2s. Verify run: 1.4s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.350 ± 0.009, crit=0.713 ± 0.014 per rating point (14 rating = 1%, 9.989 per %), hit=1.143 ± 0.043 per rating point (10 rating = 1%, 11.430 per %), melee_haste=6.766 ± 0.447

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 13.5 attack_power points (0.75 DPS) | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Tribal Worg Helm (6204, -0.15 DPS) [world]; Brawler's Leather Hood (252504, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.78 DPS) | yes | Sentinel's Medallion (19541, -0.18 DPS) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 19.9 attack_power points (1.10 DPS) | yes | Mantle of Thieves (2264, -0.44 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.45 DPS) [crafted]; Bristlebark Amice (14573, -0.49 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.5 attack_power points (0.69 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.17 DPS) [pvp] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 18.9 attack_power points (1.05 DPS) | yes | Brawler's Leather Tunic (252508, -0.12 DPS) [crafted]; Raptorbane Armor (3566, -0.16 DPS) [quest]; Tunic of Westfall (2041, -0.22 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.1 attack_power points (0.67 DPS) | yes | Jurassic Wristguards (6198, -0.11 DPS) [world]; Cultist's Armguards (270032, -0.12 DPS) [quest]; Barbaric Bracers (18948, -0.15 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.89 DPS) | yes | Insignia Gloves (6408, -0.05 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.11 DPS) [crafted]; Wolfclaw Gloves (1978, -0.16 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.33 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.32 DPS) [crafted]; Blackened Defias Belt (10403, -0.33 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.44 DPS) | yes | Petrolspill Leggings (9509, -0.39 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.39 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.50 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 13.8 attack_power points (0.77 DPS) | yes | Disjointed Shoes (277226, -0.10 DPS) [quest]; Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Insignia Boots (4055, -0.17 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 16.2 attack_power points (0.90 DPS) | yes | Thunderbrow Ring (13097, -0.23 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.34 DPS) [quest]; Monkey Ring (6748, -0.37 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.1 attack_power points (0.78 DPS) | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.23 DPS) [quest]; Monkey Ring (6748, -0.26 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (18.19 DPS) | yes | Swinetusk Shank (6691, -0.33 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.44 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.53 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (17.90 DPS) | yes | Swinetusk Shank (6691, -12.21 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -17.82 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.50 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Silver Star (3463, -0.12 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.20 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 32502110551501000-00000000000000000-0000000000000000000)

Set DPS (verified): 109.0. Weights run: 2.3s. Verify run: 1.5s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.356 ± 0.010, crit=0.739 ± 0.017 per rating point (14 rating = 1%, 10.348 per %), hit=2.210 ± 0.075 per rating point (10 rating = 1%, 22.096 per %), melee_haste=26.369 ± 0.714

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.9 attack_power points (0.87 DPS) | yes | Hawkeye's Helm (14591, -0.23 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.28 DPS) [world_drop]; Nightscape Headband (8176, -0.32 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Sentinel's Medallion (19540, -0.17 DPS) [rep]; Ghostshard Talisman (7731, -0.20 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.9 attack_power points (0.90 DPS) | yes | Forest Tracker Epaulets (2278, -0.23 DPS) [world_drop]; Flintrock Shoulders (7755, -0.28 DPS) [dungeon]; Nightscape Shoulders (8192, -0.40 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.6 attack_power points (0.59 DPS) | yes | Sergeant Major's Cape (16336, -0.11 DPS) [pvp]; Hawkeye's Cloak (14593, -0.17 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.22 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 33.8 attack_power points (1.13 DPS) | yes | Wolffear Harness (13110, -0.39 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.45 DPS) [crafted]; Dusky Leather Armor (7374, -0.50 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Hawkeye's Bracers (14590, -0.26 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.31 DPS) [dungeon]; Dusky Bracers (7378, -0.31 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.3 attack_power points (1.02 DPS) | yes | Skulker's Leather Gloves (252525, -0.26 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.27 DPS) [crafted]; Imperial Leather Gloves (4063, -0.31 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.01 DPS) | yes | Highlander's Chain Girdle (20090, -0.20 DPS) [rep]; Ogron's Sash (13117, -0.30 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.40 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 28.5 attack_power points (0.96 DPS) | yes | Triprunner Dungarees (9624, -0.04 DPS) [quest]; Ferine Leggings (6690, -0.08 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.31 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.9 attack_power points (0.74 DPS) | yes | Prowler's Leather Shoes (252465, -0.05 DPS) [crafted]; Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Excelsior Boots (4109, -0.09 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Protector's Band (19515, -0.04 DPS) [rep]; Field Researcher's Loop (281634, -0.12 DPS) [quest]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.67 DPS) | yes | Protector's Band (19515, -0.04 DPS) [rep]; Field Researcher's Loop (281634, -0.12 DPS) [quest]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.48 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.10 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (109.0 DPS) | yes | Stonecloth Branch (15963, -14.74 DPS) [world_drop]; Satyr's Rod (15962, -14.79 DPS) [world_drop]; Ardent Custodian (868, -62.62 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.13 DPS) [dungeon]; Swiftwind (13038, -0.15 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.49 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 170.4. Weights run: 3.2s. Verify run: 2.0s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.451 ± 0.009, crit=0.953 ± 0.016 per rating point (14 rating = 1%, 13.336 per %), hit=2.111 ± 0.089 per rating point (10 rating = 1%, 21.114 per %), melee_haste=6.382 ± 1.375

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 50.4 attack_power points (3.96 DPS) | yes | Ebon Mask (19984, -0.09 DPS) [quest]; Embrace of the Lycan (9479, -0.82 DPS) [dungeon]; White Bandit Mask (10008, -1.84 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.9 attack_power points (1.87 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.30 DPS) [quest]; Sentinel's Medallion (19539, -0.51 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 42.4 attack_power points (3.33 DPS) | yes | Sunburn Spaulders (274751, -0.71 DPS, sim-verified) [vendor]; Skulker's Leather Shoulder (252535, -1.37 DPS) [crafted]; Failed Flying Experiment (9647, -1.41 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.3 attack_power points (2.07 DPS) | yes | Blisterbane Wrap (12552, -0.36 DPS) [dungeon]; Dark Phantom Cape (13122, -0.36 DPS) [world_drop]; Duskbat Drape (19982, -0.47 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 52.4 attack_power points (4.12 DPS) | yes | Warbear Harness (15064, -0.77 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -1.26 DPS) [quest]; Fungus Shroud Armor (17742, -1.27 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.8 attack_power points (2.02 DPS) | yes | Skulker's Leather Bracers (252540, -0.45 DPS) [crafted]; Branded Leather Bracers (19508, -0.45 DPS) [dungeon]; Pridelord Bands (14672, -0.49 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 41.9 attack_power points (3.29 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.67 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.99 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.98 DPS) | yes | Highlander's Leather Girdle (20115, -0.37 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.45 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.52 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 52.4 attack_power points (4.12 DPS) | yes | Gryphon Rider's Leggings (9652, -0.82 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -1.27 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.73 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 30.7 attack_power points (2.41 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.13 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.1 attack_power points (3.23 DPS) | yes | Masons Fraternity Ring (9533, -1.63 DPS) [quest]; Mark of Kern (2262, -1.66 DPS) [dungeon]; Assault Band (13095, -1.66 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 23.1 attack_power points (1.81 DPS) | yes | Masons Fraternity Ring (9533, -0.22 DPS) [quest]; Mark of Kern (2262, -0.24 DPS) [dungeon]; Assault Band (13095, -0.24 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (170.4 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (170.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (170.4 DPS) | yes | Thorium Cestus (250614, -2.11 DPS) [crafted]; Doomforged Straightedge (12535, -2.37 DPS) [dungeon]; Hanzo Sword (8190, -2.42 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (42.77 DPS) | yes | Thorium Cestus (250614, +0.00 DPS) [crafted]; Claw of Celebras (17738, -4.93 DPS) [dungeon]; Thermotastic Egg Timer (9644, -42.43 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (170.4 DPS) | yes | Stinging Bow (10624, -0.26 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.26 DPS) [world_drop]; Dark Iron Rifle (16004, -1.88 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 268.4. Weights run: 3.0s. Verify run: 2.0s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.020 ± 0.021, crit=2.159 ± 0.037 per rating point (14 rating = 1%, 30.219 per %), hit=4.777 ± 0.172 per rating point (10 rating = 1%, 47.766 per %), melee_haste=16.296 ± 2.604

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 125.8 attack_power points (10.95 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -1.02 DPS) [pvp]; Knight-Lieutenant's Leather Headband (220850, -2.77 DPS) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 78.1 attack_power points (6.80 DPS) | yes | Beads of Ogre Might (22150, -0.55 DPS) [quest]; Mark of Fordring (15411, -1.90 DPS) [quest]; Medallion of the Dawn (22659, -2.08 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 119.5 attack_power points (10.41 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, -1.70 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, -2.56 DPS) [pvp]; Wyrmhide Spaulders (12082, -2.58 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 75.8 attack_power points (6.60 DPS) | yes | Cape of the Black Baron (13340, -2.22 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.44 DPS) [vendor]; Stalwart Cloak (272415, -2.44 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (268.4 DPS) | yes | Darkmantle Tunic (226825, -1.61 DPS) [quest]; Field Marshal's Leather Chestpiece (231543, -1.94 DPS) [pvp]; Tunic of Undead Slaying (23089, -10.23 DPS, sim-verified) [world] |
| wrist | Blackmist Armguards (12966) | Blackrock Spire: The Beast [dungeon] | sim-verified (268.4 DPS) | yes | Forest Stalker's Bracers (19587, -0.29 DPS) [rep]; Marshal's Leather Armsplints (16460, -0.38 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -2.55 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 78.0 attack_power points (6.79 DPS) | yes | Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; Marshal's Leather Handgrips (231544, -0.64 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.72 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 138.0 attack_power points (12.02 DPS) | yes | Belt of Preserved Heads (20216, -1.57 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -6.42 DPS) [rep]; Ferocity of the Timbermaw (227805, -6.76 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 115.0 attack_power points (10.02 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -0.26 DPS) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 81.9 attack_power points (7.14 DPS) | yes | Fine Dawn Treaders (227815, -2.08 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -2.19 DPS) [dungeon]; Darkmantle Boots (22003, -2.91 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (268.4 DPS) | yes | Tarnished Elven Ring (18500, -1.39 DPS) [dungeon]; Cutthroat's Signet (272408, -1.56 DPS) [vendor]; Naglering (11669, -5.60 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (268.4 DPS) | yes | Tarnished Elven Ring (18500, -0.53 DPS) [dungeon]; Cutthroat's Signet (272408, -0.70 DPS) [vendor]; Naglering (11669, -5.12 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (268.4 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -8.61 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (268.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Heart of Wyrmthalak (22321, -1.93 DPS, sim-verified) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (268.4 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 834.5 attack_power points (72.68 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -7.17 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -22.21 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (268.4 DPS) | yes | Blackcrow (12651, -0.42 DPS) [dungeon]; The Purifier (22656, -2.06 DPS) [quest]; Dark Iron Rifle (16004, -2.52 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Blackmist Armguards; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 574.9. Weights run: 3.2s. Verify run: 2.0s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.094 ± 0.023, crit=2.267 ± 0.039 per rating point (14 rating = 1%, 31.743 per %), hit=4.563 ± 0.233 per rating point (10 rating = 1%, 45.627 per %), melee_haste=15.582 ± 3.275

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 123.0 attack_power points (20.82 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -1.63 DPS) [pvp]; Knight-Lieutenant's Leather Headband (220850, -5.02 DPS) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 77.0 attack_power points (13.04 DPS) | yes | Beads of Ogre Might (22150, -1.26 DPS) [quest]; Mark of Fordring (15411, -3.27 DPS) [quest]; Medallion of the Dawn (22659, -3.61 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 115.3 attack_power points (19.51 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, -2.69 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, -4.34 DPS) [pvp]; Wyrmhide Spaulders (12082, -5.19 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 73.6 attack_power points (12.47 DPS) | yes | Cape of the Black Baron (13340, -3.76 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.74 DPS) [vendor]; Stalwart Cloak (272415, -4.74 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (574.9 DPS) | yes | Darkmantle Tunic (226825, -3.64 DPS) [quest]; Field Marshal's Leather Chestpiece (231543, -3.90 DPS) [pvp]; Tunic of Undead Slaying (23089, -20.03 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (574.9 DPS) | yes | Blackmist Armguards (12966, -0.03 DPS) [dungeon]; Marshal's Leather Armsplints (16460, -0.17 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -9.45 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 77.4 attack_power points (13.10 DPS) | yes | Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; Marshal's Leather Handgrips (231544, -0.63 DPS) [pvp]; Devilsaur Gauntlets (15063, -2.98 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 135.2 attack_power points (22.90 DPS) | yes | Belt of Preserved Heads (20216, -7.48 DPS) [quest]; Highlander's Leather Girdle (20045, -11.77 DPS) [rep]; Ferocity of the Timbermaw (227805, -12.47 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 120.0 attack_power points (20.32 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -1.47 DPS) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 82.2 attack_power points (13.92 DPS) | yes | Fine Dawn Treaders (227815, -3.45 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -4.15 DPS) [dungeon]; Darkmantle Boots (22003, -5.41 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (574.9 DPS) | yes | Tarnished Elven Ring (18500, -2.76 DPS) [dungeon]; Cutthroat's Signet (272408, -3.12 DPS) [vendor]; Naglering (11669, -10.27 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (574.9 DPS) | yes | Tarnished Elven Ring (18500, -1.06 DPS) [dungeon]; Cutthroat's Signet (272408, -1.42 DPS) [vendor]; Naglering (11669, -9.07 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (574.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (574.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS, sim-verified) [dungeon]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -3.80 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (574.9 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -8.77 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 835.2 attack_power points (141.40 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -9.56 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -43.29 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (574.9 DPS) | yes | Blackcrow (12651, -0.77 DPS) [dungeon]; The Purifier (22656, -3.41 DPS) [quest]; Dark Iron Rifle (16004, -3.73 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 33.7. Weights run: 2.1s. Verify run: 1.4s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.307 ± 0.007, crit=0.631 ± 0.012 per rating point (14 rating = 1%, 8.831 per %), hit=0.874 ± 0.032 per rating point (10 rating = 1%, 8.740 per %), melee_haste=5.172 ± 0.354

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 10.5 attack_power points (0.55 DPS) | yes | Defender's Leather Hood (252447, -0.13 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 7.8 attack_power points (0.42 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.5 attack_power points (0.35 DPS) | yes | Slime-encrusted Pads (6461, -0.41 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 7.8 attack_power points (0.42 DPS) | yes | Cape of the Brotherhood (5193, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 14.1 attack_power points (0.75 DPS) | yes | Defender's Leather Armor (252434, -0.20 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.23 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.27 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 6.5 attack_power points (0.35 DPS) | yes | Bristlebark Bindings (14569, -0.03 DPS) [world_drop]; Wolf Bracers (4794, -0.07 DPS) [vendor]; Ratchet Wristwraps (274742, -0.14 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 11.8 attack_power points (0.63 DPS) | yes | Brawler's Leather Gloves (252494, -0.14 DPS) [crafted]; Fletcher's Gloves (7348, -0.16 DPS) [crafted]; Bristlebark Gloves (14572, -0.17 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.96 DPS) | yes | Brawler's Leather Belt (252428, -0.47 DPS) [crafted]; Deviate Scale Belt (6468, -0.57 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.64 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 17.8 attack_power points (0.94 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.19 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 13.5 attack_power points (0.71 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.30 DPS) [dungeon]; Footpads of the Fang (10411, -0.30 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 9.2 attack_power points (0.49 DPS) | yes | Signet of the Zhevra (285330, -0.07 DPS) [world]; Demon Band (12054, -0.28 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 8.7 attack_power points (0.46 DPS) | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Demon Band (12054, -0.25 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.25 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.21 DPS) | yes | Blackfang (2236, -1.08 DPS) [world_drop]; Wingblade (6504, -1.21 DPS) [quest]; Diamond Hammer (2194, -2.35 DPS, sim-verified) [world_drop] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (33.7 DPS) | yes | Diamond Hammer (2194, -0.49 DPS, sim-verified) [world_drop]; Tork Wrench (11855, -12.07 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.2 attack_power points (0.28 DPS) | yes | Fine Longbow (11304, -0.07 DPS) [vendor]; Deadly Blunderbuss (4369, -0.14 DPS) [crafted]; Light Bow (4576, -0.14 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 59.0. Weights run: 2.2s. Verify run: 1.4s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.350 ± 0.009, crit=0.713 ± 0.014 per rating point (14 rating = 1%, 9.989 per %), hit=1.143 ± 0.043 per rating point (10 rating = 1%, 11.430 per %), melee_haste=6.766 ± 0.447

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 13.5 attack_power points (0.75 DPS) | yes | Defender's Leather Helm (252455, -0.08 DPS) [crafted]; Tribal Worg Helm (6204, -0.15 DPS) [world]; Brawler's Leather Hood (252504, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.78 DPS) | yes | Scout's Medallion (19537, -0.18 DPS) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 19.9 attack_power points (1.10 DPS) | yes | Mantle of Thieves (2264, -0.44 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.45 DPS) [crafted]; Bristlebark Amice (14573, -0.49 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.5 attack_power points (0.69 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Wildhunter Cloak (16658, -0.14 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 18.9 attack_power points (1.05 DPS) | yes | Brawler's Leather Tunic (252508, -0.12 DPS) [crafted]; Brawler's Leather Armor (252490, -0.25 DPS) [crafted]; Panther Armor (6670, -0.26 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.1 attack_power points (0.67 DPS) | yes | Jurassic Wristguards (6198, -0.11 DPS) [world]; Cultist's Armguards (270032, -0.12 DPS) [quest]; Barbaric Bracers (18948, -0.15 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.89 DPS) | yes | Insignia Gloves (6408, -0.05 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.11 DPS) [crafted]; Wolfclaw Gloves (1978, -0.16 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.33 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.32 DPS) [crafted]; Blackened Defias Belt (10403, -0.33 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.44 DPS) | yes | Petrolspill Leggings (9509, -0.39 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.39 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.48 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 13.8 attack_power points (0.77 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Insignia Boots (4055, -0.17 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.17 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 16.2 attack_power points (0.90 DPS) | yes | Thunderbrow Ring (13097, -0.23 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.34 DPS) [quest]; Monkey Ring (6748, -0.37 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.1 attack_power points (0.78 DPS) | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.23 DPS) [quest]; Monkey Ring (6748, -0.26 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (18.19 DPS) | yes | Swinetusk Shank (6691, -0.33 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.44 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.53 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (17.90 DPS) | yes | Swinetusk Shank (6691, -11.90 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -17.79 DPS) [quest]; Satyr's Rod (15962, -17.82 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.50 DPS) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Silver Star (3463, -0.12 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.20 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 32502110551501000-00000000000000000-0000000000000000000)

Set DPS (verified): 107.8. Weights run: 2.3s. Verify run: 1.4s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.356 ± 0.010, crit=0.739 ± 0.017 per rating point (14 rating = 1%, 10.348 per %), hit=2.210 ± 0.075 per rating point (10 rating = 1%, 22.096 per %), melee_haste=26.369 ± 0.714

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.9 attack_power points (0.87 DPS) | yes | Hawkeye's Helm (14591, -0.23 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.28 DPS) [world_drop]; Nightscape Headband (8176, -0.32 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Scout's Medallion (19536, -0.17 DPS) [rep]; Ghostshard Talisman (7731, -0.20 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.9 attack_power points (0.90 DPS) | yes | Forest Tracker Epaulets (2278, -0.23 DPS) [world_drop]; Flintrock Shoulders (7755, -0.28 DPS) [dungeon]; Nightscape Shoulders (8192, -0.40 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.6 attack_power points (0.59 DPS) | yes | First Sergeant's Cloak (16340, -0.11 DPS) [pvp]; Hawkeye's Cloak (14593, -0.17 DPS) [world_drop]; Parachute Cloak (10518, -0.23 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 33.8 attack_power points (1.13 DPS) | yes | Wolffear Harness (13110, -0.37 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.45 DPS) [crafted]; Dusky Leather Armor (7374, -0.50 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.26 DPS) [world_drop]; Dusky Bracers (7378, -0.31 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.3 attack_power points (1.02 DPS) | yes | Prowler's Leather Gloves (252524, -0.27 DPS) [crafted]; Imperial Leather Gloves (4063, -0.31 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.32 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.01 DPS) | yes | Defiler's Chain Girdle (20152, -0.20 DPS) [rep]; Ogron's Sash (13117, -0.30 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.40 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 28.5 attack_power points (0.96 DPS) | yes | Triprunner Dungarees (9624, -0.04 DPS) [quest]; Ferine Leggings (6690, -0.08 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.31 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.9 attack_power points (0.74 DPS) | yes | Prowler's Leather Shoes (252465, -0.05 DPS) [crafted]; Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Excelsior Boots (4109, -0.09 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Legionnaire's Band (19512, -0.04 DPS) [rep]; Field Researcher's Loop (281634, -0.12 DPS) [quest]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.67 DPS) | yes | Legionnaire's Band (19512, -0.04 DPS) [rep]; Field Researcher's Loop (281634, -0.12 DPS) [quest]; Ironspine's Eye (7686, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.48 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.10 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (107.8 DPS) | yes | Stonecloth Branch (15963, -14.74 DPS) [world_drop]; Tork Wrench (11855, -14.77 DPS) [quest]; Ardent Custodian (868, -61.88 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.13 DPS) [dungeon]; Swiftwind (13038, -0.15 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.49 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 172.2. Weights run: 3.2s. Verify run: 2.0s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.451 ± 0.009, crit=0.953 ± 0.016 per rating point (14 rating = 1%, 13.336 per %), hit=2.111 ± 0.089 per rating point (10 rating = 1%, 21.114 per %), melee_haste=6.382 ± 1.375

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 49.3 attack_power points (3.87 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -1.22 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.76 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.9 attack_power points (1.87 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.30 DPS) [quest]; Woven Ivy Necklace (19159, -0.38 DPS) [quest]; Scout's Medallion (19535, -0.51 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.0 attack_power points (2.19 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.24 DPS) [crafted]; Failed Flying Experiment (9647, -0.27 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.3 attack_power points (2.07 DPS) | yes | Blisterbane Wrap (12552, -0.36 DPS) [dungeon]; Dark Phantom Cape (13122, -0.36 DPS) [world_drop]; Duskbat Drape (19982, -0.47 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 52.4 attack_power points (4.12 DPS) | yes | Blazewind Breastplate (11193, -1.26 DPS) [quest]; Fungus Shroud Armor (17742, -1.27 DPS) [dungeon]; Warbear Harness (15064, -2.07 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.8 attack_power points (2.02 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.45 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 41.9 attack_power points (3.29 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.67 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.99 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.98 DPS) | yes | Defiler's Leather Girdle (20193, -0.37 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.45 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.52 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 52.4 attack_power points (4.12 DPS) | yes | Basilisk Hide Pants (1718, -1.73 DPS) [world_drop]; Triprunner Dungarees (9624, -1.83 DPS) [quest]; Serpentskin Leggings (8262, -2.18 DPS, sim-verified) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 30.7 attack_power points (2.41 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.13 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.1 attack_power points (3.23 DPS) | yes | Legionnaire's Band (19511, -1.42 DPS) [rep]; Masons Fraternity Ring (9533, -1.63 DPS) [quest]; Mark of Kern (2262, -1.66 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.88 DPS) | yes | Legionnaire's Band (19511, -0.07 DPS) [rep]; Masons Fraternity Ring (9533, -0.29 DPS) [quest]; Mark of Kern (2262, -0.31 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (172.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (172.2 DPS) | yes | Molten Heart of the Mountain (249470, -1.88 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (172.2 DPS) | yes | Thorium Cestus (250614, -2.11 DPS) [crafted]; Bloodrazor (809, -2.24 DPS, sim-verified) [world_drop]; Doomforged Straightedge (12535, -2.37 DPS) [dungeon] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (42.77 DPS) | yes | Thorium Cestus (250614, +0.00 DPS) [crafted]; Claw of Celebras (17738, -4.93 DPS) [dungeon]; White Bone Shredder (11863, -7.53 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (172.2 DPS) | yes | Stinging Bow (10624, -0.26 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.26 DPS) [world_drop]; Dark Iron Rifle (16004, -1.99 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 268.1. Weights run: 3.0s. Verify run: 1.9s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.020 ± 0.021, crit=2.159 ± 0.037 per rating point (14 rating = 1%, 30.219 per %), hit=4.777 ± 0.172 per rating point (10 rating = 1%, 47.766 per %), melee_haste=16.296 ± 2.604

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 125.8 attack_power points (10.95 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted]; Champion's Leather Helm (227057, -1.02 DPS) [pvp] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 78.1 attack_power points (6.80 DPS) | yes | Beads of Ogre Might (22150, -0.55 DPS) [quest]; Mark of Fordring (15411, -1.90 DPS) [quest]; Medallion of the Dawn (22659, -2.08 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 119.5 attack_power points (10.41 DPS) | yes | Champion's Leather Shoulders (227056, -1.70 DPS) [pvp]; Wyrmhide Spaulders (12082, -2.55 DPS, sim-verified) [quest]; Warlord's Leather Spaulders (231551, -2.56 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 75.8 attack_power points (6.60 DPS) | yes | Cape of the Black Baron (13340, -2.22 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.44 DPS) [vendor]; Stalwart Cloak (272415, -2.44 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (268.1 DPS) | yes | Darkmantle Tunic (226825, -1.61 DPS) [quest]; Warlord's Leather Breastplate (231549, -1.94 DPS) [pvp]; Tunic of Undead Slaying (23089, -9.78 DPS, sim-verified) [world] |
| wrist | Blackmist Armguards (12966) | Blackrock Spire: The Beast [dungeon] | sim-verified (268.1 DPS) | yes | Forest Stalker's Bracers (19587, -0.29 DPS) [rep]; General's Leather Armsplints (16559, -0.38 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -2.05 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 78.0 attack_power points (6.79 DPS) | yes | Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; General's Leather Mitts (231555, -0.64 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.72 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 138.0 attack_power points (12.02 DPS) | yes | Belt of Preserved Heads (20216, -1.09 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -6.42 DPS) [rep]; Ferocity of the Timbermaw (227805, -6.76 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 115.0 attack_power points (10.02 DPS) | yes | General's Leather Legguards (231554, +0.00 DPS) [pvp]; Legionnaire's Leather Legguards (227059, -0.26 DPS) [pvp]; Plaguehound Leggings (18736, -2.19 DPS, sim-verified) [dungeon] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 81.9 attack_power points (7.14 DPS) | yes | Fine Dawn Treaders (227815, -1.18 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -2.19 DPS) [dungeon]; Darkmantle Boots (22003, -2.91 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (268.1 DPS) | yes | Tarnished Elven Ring (18500, -1.39 DPS) [dungeon]; Cutthroat's Signet (272408, -1.56 DPS) [vendor]; Naglering (11669, -5.12 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (268.1 DPS) | yes | Tarnished Elven Ring (18500, -0.53 DPS) [dungeon]; Cutthroat's Signet (272408, -0.70 DPS) [vendor]; Naglering (11669, -4.65 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (268.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (268.1 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -1.54 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -2.83 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (268.1 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 834.5 attack_power points (72.68 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -6.47 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -22.21 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (268.1 DPS) | yes | Blackcrow (12651, -0.42 DPS) [dungeon]; Dark Iron Rifle (16004, -2.02 DPS, sim-verified) [crafted]; The Purifier (22656, -2.06 DPS) [quest] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Blackmist Armguards; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 569.1. Weights run: 3.2s. Verify run: 2.1s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.094 ± 0.023, crit=2.267 ± 0.039 per rating point (14 rating = 1%, 31.743 per %), hit=4.563 ± 0.233 per rating point (10 rating = 1%, 45.627 per %), melee_haste=15.582 ± 3.275

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 123.0 attack_power points (20.82 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS) [crafted]; Champion's Leather Helm (227057, -1.63 DPS) [pvp] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 77.0 attack_power points (13.04 DPS) | yes | Beads of Ogre Might (22150, -1.26 DPS) [quest]; Mark of Fordring (15411, -3.27 DPS) [quest]; Medallion of the Dawn (22659, -3.61 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 115.3 attack_power points (19.51 DPS) | yes | Champion's Leather Shoulders (227056, -2.69 DPS) [pvp]; Warlord's Leather Spaulders (231551, -4.34 DPS) [pvp]; Wyrmhide Spaulders (12082, -4.87 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 73.6 attack_power points (12.47 DPS) | yes | Cape of the Black Baron (13340, -3.76 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.74 DPS) [vendor]; Stalwart Cloak (272415, -4.74 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (569.1 DPS) | yes | Darkmantle Tunic (226825, -3.64 DPS) [quest]; Warlord's Leather Breastplate (231549, -3.90 DPS) [pvp]; Tunic of Undead Slaying (23089, -21.25 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (569.1 DPS) | yes | Blackmist Armguards (12966, -0.03 DPS) [dungeon]; General's Leather Armsplints (16559, -0.17 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -8.88 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 77.4 attack_power points (13.10 DPS) | yes | Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; General's Leather Mitts (231555, -0.63 DPS) [pvp]; Devilsaur Gauntlets (15063, -2.98 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 135.2 attack_power points (22.90 DPS) | yes | Belt of Preserved Heads (20216, -3.49 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -11.77 DPS) [rep]; Ferocity of the Timbermaw (227805, -12.47 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 120.0 attack_power points (20.32 DPS) | yes | General's Leather Legguards (231554, +0.00 DPS) [pvp]; Legionnaire's Leather Legguards (227059, -1.47 DPS) [pvp]; Plaguehound Leggings (18736, -5.41 DPS, sim-verified) [dungeon] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 82.2 attack_power points (13.92 DPS) | yes | Fine Dawn Treaders (227815, -3.32 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -4.15 DPS) [dungeon]; Darkmantle Boots (22003, -5.41 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (569.1 DPS) | yes | Tarnished Elven Ring (18500, -2.76 DPS) [dungeon]; Cutthroat's Signet (272408, -3.12 DPS) [vendor]; Naglering (11669, -11.64 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (569.1 DPS) | yes | Tarnished Elven Ring (18500, -1.06 DPS) [dungeon]; Cutthroat's Signet (272408, -1.42 DPS) [vendor]; Naglering (11669, -10.48 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (569.1 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -1.77 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.57 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (569.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Hand of Justice (11815, -5.99 DPS, sim-verified) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (569.1 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS, sim-verified) [world_drop]; High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 835.2 attack_power points (141.40 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -10.07 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -43.29 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (569.1 DPS) | yes | Blackcrow (12651, -0.77 DPS) [dungeon]; The Purifier (22656, -3.41 DPS) [quest]; Dark Iron Rifle (16004, -5.33 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket2: Darkmoon Card: Maelstrom; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

