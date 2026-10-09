# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 40.9. Weights run: 2.2s. Verify run: 3.5s. 198 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.032 ± 0.013, crit=0.603 ± 0.012 per rating point (14 rating = 1%, 8.438 per %), hit=0.880 ± 0.030 per rating point (10 rating = 1%, 8.802 per %), melee_haste=4.857 ± 0.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.3 attack_power points (0.94 DPS) | yes | Defender's Leather Hood (252447, -0.46 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.2 attack_power points (0.70 DPS) | yes | Erudite's Amulet (277204, -0.54 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.2 attack_power points (0.58 DPS) | yes | Slime-encrusted Pads (6461, -0.68 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.2 attack_power points (0.70 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Hide of Lupos (3018, -0.23 DPS) [world]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 22.4 attack_power points (1.29 DPS) | yes | Brawler's Leather Armor (252490, -0.18 DPS) [crafted]; Trapper's Leather Armor (252491, -0.47 DPS) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.2 attack_power points (0.58 DPS) | yes | Bravo's Armbands (270015, -0.00 DPS) [quest]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.9 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.97 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.04 DPS) | yes | Brawler's Leather Belt (252428, -0.34 DPS) [crafted]; Dusty Belt (279897, -0.45 DPS) [quest]; Deviate Scale Belt (6468, -3.14 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.9 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.22 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (40.9 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.13 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.2 attack_power points (0.70 DPS) | yes | Pyrewood Signet Ring (277210, -0.03 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon]; Demon Band (12054, -0.47 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.1 attack_power points (0.70 DPS) | yes | Pyrewood Signet Ring (277210, -0.33 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Demon Band (12054, -0.47 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (14.32 DPS) | yes | Diamond Hammer (2194, +0.00 DPS, sim-verified) [world_drop]; Blackfang (2236, -1.17 DPS) [world_drop]; Assassin's Blade (1935, -1.17 DPS) [dungeon] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (40.9 DPS) | yes | Diamond Hammer (2194, -0.40 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.1 attack_power points (0.47 DPS) | yes | Light Bow (4576, -0.23 DPS) [world_drop]; Owlsight Rifle (15205, -0.23 DPS) [quest]; Deadly Blunderbuss (4369, -0.27 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 57.6. Weights run: 2.3s. Verify run: 3.5s. 358 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.737 ± 0.012, crit=0.680 ± 0.013 per rating point (14 rating = 1%, 9.516 per %), hit=1.066 ± 0.036 per rating point (10 rating = 1%, 10.656 per %), melee_haste=5.299 ± 0.419

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.4 attack_power points (1.02 DPS) | yes | Tribal Worg Helm (6204, -0.20 DPS) [world]; Brawler's Leather Hood (252504, -0.20 DPS) [crafted]; Humbert's Helm (4724, -0.31 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.83 DPS) | yes | Sentinel's Medallion (19541, -0.01 DPS) [rep]; Kaleidoscope Chain (13084, -0.18 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (1.42 DPS) | yes | Mantle of Thieves (2264, -0.41 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.61 DPS) [crafted]; Bristlebark Amice (14573, -0.63 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.2 attack_power points (0.89 DPS) | yes | Tigerstrike Mantle (13108, -0.07 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.25 DPS) [pvp]; Cloak of Night (4447, -0.28 DPS) [world] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.3 attack_power points (1.43 DPS) | yes | Brawler's Leather Tunic (252508, -0.21 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.31 DPS) [quest]; Brawler's Leather Armor (252490, -0.42 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.4 attack_power points (0.85 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.20 DPS) [crafted]; Insignia Bracers (6410, -0.24 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (57.6 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.53 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (57.6 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -2.81 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (57.6 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.52 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (57.6 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Feet of the Lynx (1121, -2.80 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 attack_power points (1.16 DPS) | yes | Thunderbrow Ring (13097, -0.38 DPS) [world_drop]; Monkey Ring (6748, -0.44 DPS) [quest]; Pyrewood Signet Ring (277210, -0.50 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.4 attack_power points (0.97 DPS) | yes | Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Monkey Ring (6748, -0.25 DPS) [quest]; Pyrewood Signet Ring (277210, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (19.35 DPS) | yes | Royal Diplomatic Scepter (9457, -0.31 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.46 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.56 DPS) [dungeon] |
| off_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (57.6 DPS) | yes | Royal Diplomatic Scepter (9457, -5.39 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -18.90 DPS) [world_drop] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.71 DPS) | yes | Golemsight Long Gun (273029, -0.09 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.18 DPS) [vendor]; Silver Star (3463, -0.20 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Ironspine's Fist; off_hand: Swinetusk Shank; ranged: Glass Shooter

No-known-source sample (15 of 358, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 32502110551401001-00000000000000000-0000000000000000000)

Set DPS (verified): 109.0. Weights run: 2.3s. Verify run: 2.4s. 483 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.566 ± 0.013, crit=0.729 ± 0.017 per rating point (14 rating = 1%, 10.208 per %), hit=2.147 ± 0.072 per rating point (10 rating = 1%, 21.474 per %), melee_haste=26.807 ± 0.730

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.2 attack_power points (0.97 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.27 DPS) [world_drop]; Nightscape Headband (8176, -0.32 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Ghostshard Talisman (7731, -0.11 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.24 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.2 attack_power points (1.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Flintrock Shoulders (7755, -0.29 DPS) [dungeon]; Nightscape Shoulders (8192, -0.41 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.7 attack_power points (0.68 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Parachute Cloak (10518, -0.25 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (109.0 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.11 DPS) [crafted]; Dusky Leather Armor (7374, -0.16 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.26 DPS) [dungeon]; Dusky Bracers (7378, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.2 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.19 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.21 DPS) [crafted]; Imperial Leather Gloves (4063, -0.24 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.03 DPS) | yes | Highlander's Chain Girdle (20090, -0.21 DPS) [rep]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.34 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 32.9 attack_power points (1.13 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.24 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.38 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.2 attack_power points (0.83 DPS) | yes | Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.11 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 20.5 attack_power points (0.71 DPS) | yes | Assault Band (13095, -0.02 DPS) [world_drop]; Ironspine's Eye (7686, -0.08 DPS) [dungeon]; Ring of the Underwood (2951, -0.10 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ironspine's Eye (7686, -0.07 DPS) [dungeon]; Ring of the Underwood (2951, -0.08 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.50 DPS) [world_drop]; Vanquisher's Sword (10823, -1.12 DPS) [quest] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (+63.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Stonecloth Branch (15963, -15.15 DPS) [world_drop]; Satyr's Rod (15962, -15.20 DPS) [world_drop]; Ardent Custodian (868, -63.29 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glass Shooter (9456, -0.07 DPS) [dungeon]; Swiftwind (13038, -0.10 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.52 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 483, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 172.5. Weights run: 3.2s. Verify run: 3.3s. 610 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.603 ± 0.012, crit=0.995 ± 0.018 per rating point (14 rating = 1%, 13.924 per %), hit=2.074 ± 0.067 per rating point (10 rating = 1%, 20.742 per %), melee_haste=7.447 ± 1.180

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 50.7 attack_power points (4.15 DPS) | yes | Ebon Mask (19984, -0.06 DPS) [quest]; Embrace of the Lycan (9479, -0.87 DPS) [dungeon]; White Bandit Mask (10008, -1.80 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.8 attack_power points (2.12 DPS) | yes | Sentinel's Medallion (19539, -0.54 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.71 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 42.7 attack_power points (3.49 DPS) | yes | Skulker's Leather Shoulder (252535, -1.31 DPS) [crafted]; Sunburn Spaulders (274751, -1.33 DPS, sim-verified) [vendor]; Failed Flying Experiment (9647, -1.36 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Duskbat Drape (19982, -0.13 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 52.7 attack_power points (4.31 DPS) | yes | Blazewind Breastplate (11193, -1.05 DPS) [quest]; Warbear Harness (15064, -1.05 DPS) [crafted]; Fungus Shroud Armor (17742, -1.20 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 28.0 attack_power points (2.30 DPS) | yes | Pridelord Bands (14672, -0.57 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.64 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.69 DPS, sim-verified) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 45.3 attack_power points (3.71 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.93 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.13 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20115, -0.04 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 52.7 attack_power points (4.31 DPS) | yes | Serpentskin Leggings (8262, -1.15 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.37 DPS, sim-verified) [quest]; Basilisk Hide Pants (1718, -1.56 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 33.3 attack_power points (2.72 DPS) | yes | Albino Crocscale Boots (17728, -0.10 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.12 DPS) [crafted]; Prowler's Leather Boots (252468, -0.21 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.4 attack_power points (2.00 DPS) | yes | Mark of Kern (2262, -0.36 DPS) [dungeon]; Assault Band (13095, -0.36 DPS) [world_drop]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.4 attack_power points (1.84 DPS) | yes | Mark of Kern (2262, -0.20 DPS) [dungeon]; Assault Band (13095, -0.20 DPS) [world_drop]; Blackstone Ring (17713, -0.20 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadowblade (2163, -0.68 DPS) [world_drop]; Hanzo Sword (8190, -1.10 DPS, sim-verified) [world_drop]; Doomforged Straightedge (12535, -2.48 DPS) [dungeon] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (172.5 DPS) | yes | Claw of Celebras (17738, -3.63 DPS) [dungeon]; Thermotastic Egg Timer (9644, -42.72 DPS) [quest]; Stonecloth Branch (15963, -42.87 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.45 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.45 DPS) [world_drop]; Dark Iron Rifle (16004, -2.25 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 610, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 291.5. Weights run: 3.0s. Verify run: 8.7s. 1378 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.056 ± 0.022, crit=2.203 ± 0.038 per rating point (14 rating = 1%, 30.841 per %), hit=4.472 ± 0.146 per rating point (10 rating = 1%, 44.722 per %), melee_haste=not significant (9.491 ± 2.425)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 91.6 attack_power points (8.17 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.19 DPS) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (291.5 DPS) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -1.06 DPS) [quest]; Medallion of the Dawn (22659, -1.24 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -17.72 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.7 attack_power points (6.49 DPS) | yes | Cape of the Black Baron (13340, -1.95 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.50 DPS) [vendor]; Stalwart Cloak (272415, -2.50 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmantle Tunic (226825, -1.86 DPS) [quest]; Field Marshal's Leather Chestpiece (231543, -2.02 DPS) [pvp]; Tunic of Undead Slaying (23089, -11.86 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackmist Armguards (12966, -0.03 DPS) [dungeon]; Marshal's Leather Armsplints (16460, -0.09 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -4.88 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 75.6 attack_power points (6.75 DPS) | yes | Marshal's Leather Handgrips (231544, -0.32 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.49 DPS) [crafted]; Raider Gloves (272099, -20.89 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 132.6 attack_power points (11.84 DPS) | yes | Belt of Preserved Heads (20216, -2.06 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -6.05 DPS) [rep]; Ferocity of the Timbermaw (227805, -6.40 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -17.09 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 80.7 attack_power points (7.20 DPS) | yes | Fine Dawn Treaders (227815, -1.31 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -2.15 DPS) [dungeon]; Darkmantle Boots (22003, -2.79 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.65 DPS) [dungeon]; Cutthroat's Signet (272408, -1.84 DPS) [vendor]; Naglering (11669, -7.24 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.43 DPS) [dungeon]; Cutthroat's Signet (272408, -1.61 DPS) [vendor]; Naglering (11669, -6.67 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -8.23 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -1.98 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -6.03 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 834.8 attack_power points (74.52 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -7.42 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -22.80 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Purifier (22656, -1.79 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.97 DPS) [world_drop]; Dark Iron Rifle (16004, -3.07 DPS, sim-verified) [crafted] |

**New at 60:** neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1378, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 01532310421501000-31530300001500000-0020000000000000000)

Set DPS (verified): 746.6. Weights run: 3.2s. Verify run: 9.2s. 1378 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.456 ± 0.027, crit=2.549 ± 0.043 per rating point (14 rating = 1%, 35.691 per %), hit=4.492 ± 0.230 per rating point (10 rating = 1%, 44.917 per %), melee_haste=15.529 ± 3.211

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 105.1 attack_power points (19.60 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -1.50 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -1.35 DPS) [quest]; Medallion of the Dawn (22659, -1.72 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -51.19 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.9 attack_power points (13.59 DPS) | yes | Cape of the Black Baron (13340, -3.00 DPS) [dungeon]; Cloak of the Honor Guard (20073, -4.97 DPS) [rep]; Windshear Cape (20691, -5.08 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, -5.04 DPS) [pvp]; Darkmantle Tunic (226825, -5.21 DPS) [quest]; Tunic of Undead Slaying (23089, -29.14 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Armsplints (16460, -0.21 DPS) [pvp]; Blackmist Armguards (12966, -1.56 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -11.93 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 80.6 attack_power points (15.03 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -2.49 DPS) [quest]; Raider Gloves (272099, -58.02 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 141.4 attack_power points (26.36 DPS) | yes | Belt of Preserved Heads (20216, -5.71 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -13.25 DPS) [vendor]; Highlander's Leather Girdle (20045, -13.37 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -52.17 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 90.4 attack_power points (16.85 DPS) | yes | Shadowcraft Boots (16711, -4.72 DPS) [dungeon]; Fine Dawn Treaders (227815, -5.61 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -5.86 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -4.12 DPS) [dungeon]; Cutthroat's Signet (272408, -4.58 DPS) [vendor]; Naglering (11669, -17.79 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.77 DPS) [dungeon]; Cutthroat's Signet (272408, -3.23 DPS) [vendor]; Naglering (11669, -15.76 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (746.6 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.77 DPS) [crafted]; Counterattack Lodestone (18537, -9.21 DPS) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -8.00 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 839.7 attack_power points (156.55 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -11.07 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -48.52 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Purifier (22656, -3.09 DPS) [quest]; Precisely Calibrated Boomstick (2100, -3.34 DPS) [world_drop]; Dark Iron Rifle (16004, -7.97 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1378, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 39.9. Weights run: 2.2s. Verify run: 3.5s. 191 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.032 ± 0.013, crit=0.603 ± 0.012 per rating point (14 rating = 1%, 8.438 per %), hit=0.880 ± 0.030 per rating point (10 rating = 1%, 8.802 per %), melee_haste=4.857 ± 0.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.3 attack_power points (0.94 DPS) | yes | Defender's Leather Hood (252447, -0.44 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.2 attack_power points (0.70 DPS) | yes | Erudite's Amulet (277204, -0.53 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.2 attack_power points (0.58 DPS) | yes | Slime-encrusted Pads (6461, -0.66 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.2 attack_power points (0.70 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Hide of Lupos (3018, -0.23 DPS) [world]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (39.9 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.07 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.2 attack_power points (0.58 DPS) | yes | Wolf Bracers (4794, -0.12 DPS) [vendor]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.23 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 16.2 attack_power points (0.93 DPS) | yes | Bristlebark Gloves (14572, -0.23 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.23 DPS) [crafted]; Serpent Gloves (5970, -0.31 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.04 DPS) | yes | Brawler's Leather Belt (252428, -0.34 DPS) [crafted]; Dark Leather Belt (4249, -0.57 DPS) [crafted]; Deviate Scale Belt (6468, -3.17 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.9 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.28 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (39.9 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.18 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.2 attack_power points (0.70 DPS) | yes | Pyrewood Signet Ring (277210, -0.03 DPS) [quest]; Bounty Hunter's Ring (5351, -0.35 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.1 attack_power points (0.70 DPS) | yes | Bounty Hunter's Ring (5351, -0.35 DPS) [quest]; Pyrewood Signet Ring (277210, -0.35 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (14.32 DPS) | yes | Wingblade (6504, -1.11 DPS) [quest]; Cruel Barb (5191, -1.11 DPS) [dungeon]; Blackfang (2236, -1.17 DPS) [world_drop] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 229.8 attack_power points (13.22 DPS) | yes | Wingblade (6504, -4.96 DPS, sim-verified) [quest]; Tork Wrench (11855, -13.10 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.1 attack_power points (0.47 DPS) | yes | Light Bow (4576, -0.23 DPS) [world_drop]; Privateer Musket (5309, -0.23 DPS) [quest]; Deadly Blunderbuss (4369, -0.26 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Shadowfang; off_hand: Diamond Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 191, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 57.2. Weights run: 2.3s. Verify run: 3.6s. 348 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.737 ± 0.012, crit=0.680 ± 0.013 per rating point (14 rating = 1%, 9.516 per %), hit=1.066 ± 0.036 per rating point (10 rating = 1%, 10.656 per %), melee_haste=5.299 ± 0.419

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.4 attack_power points (1.02 DPS) | yes | Tribal Worg Helm (6204, -0.20 DPS) [world]; Brawler's Leather Hood (252504, -0.20 DPS) [crafted]; Humbert's Helm (4724, -0.31 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.83 DPS) | yes | Scout's Medallion (19537, -0.01 DPS) [rep]; Kaleidoscope Chain (13084, -0.18 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (1.42 DPS) | yes | Mantle of Thieves (2264, -0.41 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.61 DPS) [crafted]; Bristlebark Amice (14573, -0.63 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.2 attack_power points (0.89 DPS) | yes | Tigerstrike Mantle (13108, -0.07 DPS) [world_drop]; Cloak of Night (4447, -0.28 DPS) [world]; Swiftrunner Cape (6745, -0.28 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.3 attack_power points (1.43 DPS) | yes | Brawler's Leather Tunic (252508, -0.23 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.39 DPS) [quest]; Brawler's Leather Armor (252490, -0.42 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.4 attack_power points (0.85 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.20 DPS) [crafted]; Insignia Bracers (6410, -0.24 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (57.2 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.52 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (57.2 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -2.81 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (57.2 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.53 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (57.2 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Insignia Boots (4055, +0.00 DPS) [world_drop]; Vorrel's Boots (7751, -2.77 DPS, sim-verified) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 attack_power points (1.16 DPS) | yes | Thunderbrow Ring (13097, -0.38 DPS) [world_drop]; Monkey Ring (6748, -0.44 DPS) [quest]; Pyrewood Signet Ring (277210, -0.50 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.4 attack_power points (0.97 DPS) | yes | Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Monkey Ring (6748, -0.25 DPS) [quest]; Pyrewood Signet Ring (277210, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (19.35 DPS) | yes | Royal Diplomatic Scepter (9457, -0.31 DPS) [dungeon]; Serrated Raptor Claw (280805, -0.43 DPS) [quest]; Scorn's Focal Dagger (23168, -0.46 DPS) [dungeon] |
| off_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (57.2 DPS) | yes | Royal Diplomatic Scepter (9457, -5.05 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -18.88 DPS) [quest]; Satyr's Rod (15962, -18.90 DPS) [world_drop] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.71 DPS) | yes | Golemsight Long Gun (273029, -0.09 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.18 DPS) [vendor]; Silver Star (3463, -0.20 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Ironspine's Fist; off_hand: Swinetusk Shank; ranged: Glass Shooter

No-known-source sample (15 of 348, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 32502110551401001-00000000000000000-0000000000000000000)

Set DPS (verified): 108.3. Weights run: 2.3s. Verify run: 2.4s. 465 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.566 ± 0.013, crit=0.729 ± 0.017 per rating point (14 rating = 1%, 10.208 per %), hit=2.147 ± 0.072 per rating point (10 rating = 1%, 21.474 per %), melee_haste=26.807 ± 0.730

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.2 attack_power points (0.97 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.27 DPS) [world_drop]; Nightscape Headband (8176, -0.32 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Ghostshard Talisman (7731, -0.11 DPS) [dungeon]; Ethereal Talisman (4430, -0.20 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.2 attack_power points (1.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Flintrock Shoulders (7755, -0.29 DPS) [dungeon]; Nightscape Shoulders (8192, -0.41 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.7 attack_power points (0.68 DPS) | yes | First Sergeant's Cloak (16340, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Parachute Cloak (10518, -0.25 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (108.3 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.11 DPS) [crafted]; Dusky Leather Armor (7374, -0.16 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Dusky Bracers (7378, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.2 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.19 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.21 DPS) [crafted]; Imperial Leather Gloves (4063, -0.24 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.03 DPS) | yes | Defiler's Chain Girdle (20152, -0.21 DPS) [rep]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.34 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 32.9 attack_power points (1.13 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.24 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.38 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.2 attack_power points (0.83 DPS) | yes | Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.11 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 20.5 attack_power points (0.71 DPS) | yes | Assault Band (13095, -0.02 DPS) [world_drop]; Ironspine's Eye (7686, -0.08 DPS) [dungeon]; Ring of the Underwood (2951, -0.10 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ironspine's Eye (7686, -0.07 DPS) [dungeon]; Ring of the Underwood (2951, -0.08 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.50 DPS) [world_drop]; Vanquisher's Sword (10823, -1.12 DPS) [quest] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (+62.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Stonecloth Branch (15963, -15.15 DPS) [world_drop]; Tork Wrench (11855, -15.19 DPS) [quest]; Ardent Custodian (868, -62.12 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glass Shooter (9456, -0.07 DPS) [dungeon]; Swiftwind (13038, -0.10 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.52 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 465, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 174.9. Weights run: 3.2s. Verify run: 3.4s. 587 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.603 ± 0.012, crit=0.995 ± 0.018 per rating point (14 rating = 1%, 13.924 per %), hit=2.074 ± 0.067 per rating point (10 rating = 1%, 20.742 per %), melee_haste=7.447 ± 1.180

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 49.9 attack_power points (4.09 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -1.22 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.74 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.8 attack_power points (2.12 DPS) | yes | Woven Ivy Necklace (19159, -0.44 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.48 DPS) [quest]; Scout's Medallion (19535, -0.54 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.6 attack_power points (2.43 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.25 DPS) [crafted]; Failed Flying Experiment (9647, -0.30 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Duskbat Drape (19982, -0.13 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 52.7 attack_power points (4.31 DPS) | yes | Blazewind Breastplate (11193, -1.05 DPS) [quest]; Warbear Harness (15064, -1.05 DPS) [crafted]; Fungus Shroud Armor (17742, -1.48 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 28.0 attack_power points (2.30 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.54 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 45.3 attack_power points (3.71 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.93 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.13 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20193, -0.04 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 52.7 attack_power points (4.31 DPS) | yes | Basilisk Hide Pants (1718, -1.56 DPS) [world_drop]; Serpentskin Leggings (8262, -1.67 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -1.70 DPS) [quest] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 33.3 attack_power points (2.72 DPS) | yes | Albino Crocscale Boots (17728, -0.10 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.12 DPS) [crafted]; Prowler's Leather Boots (252468, -0.21 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.4 attack_power points (2.00 DPS) | yes | Masons Fraternity Ring (9533, -0.16 DPS) [quest]; Mark of Kern (2262, -0.36 DPS) [dungeon]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.97 DPS) | yes | Masons Fraternity Ring (9533, -0.13 DPS) [quest]; Mark of Kern (2262, -0.33 DPS) [dungeon]; Blackstone Ring (17713, -0.33 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.53 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Shadowblade (2163, -0.68 DPS) [world_drop]; Doomforged Straightedge (12535, -2.48 DPS) [dungeon] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (174.9 DPS) | yes | Claw of Celebras (17738, -3.63 DPS) [dungeon]; White Bone Shredder (11863, -6.24 DPS) [quest]; Thermotastic Egg Timer (9644, -42.72 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.45 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.45 DPS) [world_drop]; Dark Iron Rifle (16004, -2.24 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 587, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 291.1. Weights run: 3.0s. Verify run: 8.6s. 1372 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.056 ± 0.022, crit=2.203 ± 0.038 per rating point (14 rating = 1%, 30.841 per %), hit=4.472 ± 0.146 per rating point (10 rating = 1%, 44.722 per %), melee_haste=not significant (9.491 ± 2.425)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 89.4 attack_power points (7.98 DPS) | yes | Darkmantle Cap (226829, +0.00 DPS) [quest]; Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (291.1 DPS) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -1.06 DPS) [quest]; Medallion of the Dawn (22659, -1.24 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -15.60 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.7 attack_power points (6.49 DPS) | yes | Cape of the Black Baron (13340, -1.95 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.50 DPS) [vendor]; Stalwart Cloak (272415, -2.50 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmantle Tunic (226825, -1.86 DPS) [quest]; Warlord's Leather Breastplate (231549, -2.02 DPS) [pvp]; Tunic of Undead Slaying (23089, -11.67 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackmist Armguards (12966, -0.03 DPS) [dungeon]; General's Leather Armsplints (16559, -0.09 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -4.87 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Mitts (231555, -0.32 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.49 DPS) [crafted]; Raider Gloves (272099, -18.53 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 132.6 attack_power points (11.84 DPS) | yes | Belt of Preserved Heads (20216, -2.11 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -6.05 DPS) [rep]; Ferocity of the Timbermaw (227805, -6.40 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -16.28 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 80.7 attack_power points (7.20 DPS) | yes | Fine Dawn Treaders (227815, -1.19 DPS) [vendor]; Shadowcraft Boots (16711, -2.15 DPS) [dungeon]; Darkmantle Boots (22003, -2.79 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.65 DPS) [dungeon]; Cutthroat's Signet (272408, -1.84 DPS) [vendor]; Naglering (11669, -7.02 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.43 DPS) [dungeon]; Cutthroat's Signet (272408, -1.61 DPS) [vendor]; Naglering (11669, -6.47 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+10.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -1.04 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.95 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -3.98 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 834.8 attack_power points (74.52 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -5.89 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -22.80 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Purifier (22656, -1.79 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.97 DPS) [world_drop]; Dark Iron Rifle (16004, -3.17 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1372, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 01532310421501000-31530300001500000-0020000000000000000)

Set DPS (verified): 742.0. Weights run: 3.2s. Verify run: 9.0s. 1372 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.456 ± 0.027, crit=2.549 ± 0.043 per rating point (14 rating = 1%, 35.691 per %), hit=4.492 ± 0.230 per rating point (10 rating = 1%, 44.917 per %), melee_haste=15.529 ± 3.211

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 105.1 attack_power points (19.60 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -1.50 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (742.0 DPS) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -1.35 DPS) [quest]; Medallion of the Dawn (22659, -1.72 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -51.78 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.9 attack_power points (13.59 DPS) | yes | Cape of the Black Baron (13340, -4.22 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -4.97 DPS) [rep]; Windshear Cape (20691, -5.08 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, -5.04 DPS) [pvp]; Darkmantle Tunic (226825, -5.21 DPS) [quest]; Tunic of Undead Slaying (23089, -33.22 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Armsplints (16559, -0.21 DPS) [pvp]; Blackmist Armguards (12966, -1.56 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -12.13 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 80.6 attack_power points (15.03 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -2.49 DPS) [quest]; Raider Gloves (272099, -57.61 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 141.4 attack_power points (26.36 DPS) | yes | Belt of Preserved Heads (20216, -8.62 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -13.25 DPS) [vendor]; Defiler's Leather Girdle (20190, -13.37 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -49.09 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 90.4 attack_power points (16.85 DPS) | yes | Shadowcraft Boots (16711, -4.72 DPS) [dungeon]; Fine Dawn Treaders (227815, -5.00 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -5.86 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -4.12 DPS) [dungeon]; Cutthroat's Signet (272408, -4.58 DPS) [vendor]; Naglering (11669, -21.10 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.77 DPS) [dungeon]; Cutthroat's Signet (272408, -3.23 DPS) [vendor]; Naglering (11669, -19.13 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+16.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -0.38 DPS) [quest]; Hand of Justice (11815, -4.86 DPS, sim-verified) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 839.7 attack_power points (156.55 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -13.49 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -48.52 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Purifier (22656, -3.09 DPS) [quest]; Precisely Calibrated Boomstick (2100, -3.34 DPS) [world_drop]; Dark Iron Rifle (16004, -10.81 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1372, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

