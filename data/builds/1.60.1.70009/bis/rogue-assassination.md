# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 40.9. Weights run: 1.5s. Verify run: 2.6s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.032 ± 0.013, crit=0.603 ± 0.012 per rating point (14 rating = 1%, 8.438 per %), hit=0.880 ± 0.030 per rating point (10 rating = 1%, 8.802 per %), melee_haste=4.857 ± 0.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.3 attack_power points (0.94 DPS) | yes | Defender's Leather Hood (252447, -0.46 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.2 attack_power points (0.70 DPS) | yes | Erudite's Amulet (277204, -0.27 DPS, sim-verified) [quest] |
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

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 64.7. Weights run: 1.6s. Verify run: 1.4s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.737 ± 0.012, crit=0.680 ± 0.013 per rating point (14 rating = 1%, 9.516 per %), hit=1.066 ± 0.036 per rating point (10 rating = 1%, 10.656 per %), melee_haste=5.299 ± 0.419

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.4 attack_power points (1.02 DPS) | yes | Tribal Worg Helm (6204, -0.20 DPS) [world]; Brawler's Leather Hood (252504, -0.20 DPS) [crafted]; Humbert's Helm (4724, -0.31 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.83 DPS) | yes | Sentinel's Medallion (19541, -0.01 DPS) [rep]; Kaleidoscope Chain (13084, -0.18 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (1.42 DPS) | yes | Mantle of Thieves (2264, -0.52 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.61 DPS) [crafted]; Bristlebark Amice (14573, -0.63 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.2 attack_power points (0.89 DPS) | yes | Tigerstrike Mantle (13108, -0.07 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.25 DPS) [pvp]; Cloak of Night (4447, -0.28 DPS) [world] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.3 attack_power points (1.43 DPS) | yes | Tunic of Westfall (2041, -0.31 DPS) [quest]; Brawler's Leather Tunic (252508, -0.36 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.42 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.4 attack_power points (0.85 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.20 DPS) [crafted]; Insignia Bracers (6410, -0.24 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 17.4 attack_power points (1.03 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.08 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.42 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.14 DPS) [crafted]; Prowler's Leather Belt (252459, -0.27 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.53 DPS) | yes | Petrolspill Leggings (9509, -0.10 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.10 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.20 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.9 attack_power points (1.00 DPS) | yes | Insignia Boots (4055, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor]; Brawler's Leather Boots (252439, -0.19 DPS) [crafted] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 attack_power points (1.16 DPS) | yes | Thunderbrow Ring (13097, -0.38 DPS) [world_drop]; Monkey Ring (6748, -0.44 DPS) [quest]; Pyrewood Signet Ring (277210, -0.50 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.4 attack_power points (0.97 DPS) | yes | Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Monkey Ring (6748, -0.25 DPS) [quest]; Pyrewood Signet Ring (277210, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (19.35 DPS) | yes | Swinetusk Shank (6691, -0.35 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.46 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.56 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (19.03 DPS) | yes | Swinetusk Shank (6691, -14.93 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -18.93 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Silver Star (3463, -0.02 DPS) [quest]; Double-barreled Shotgun (2098, -0.05 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.12 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 32502110551401001-00000000000000000-0000000000000000000)

Set DPS (verified): 109.9. Weights run: 1.5s. Verify run: 1.6s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.566 ± 0.013, crit=0.729 ± 0.017 per rating point (14 rating = 1%, 10.208 per %), hit=2.147 ± 0.072 per rating point (10 rating = 1%, 21.474 per %), melee_haste=26.807 ± 0.730

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.2 attack_power points (0.97 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.27 DPS) [world_drop]; Nightscape Headband (8176, -0.32 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.69 DPS) | yes | Sentinel's Medallion (19540, -0.10 DPS) [rep]; Ghostshard Talisman (7731, -0.21 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.2 attack_power points (1.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Flintrock Shoulders (7755, -0.29 DPS) [dungeon]; Nightscape Shoulders (8192, -0.41 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.7 attack_power points (0.68 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Parachute Cloak (10518, -0.25 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.8 attack_power points (1.30 DPS) | yes | Wolffear Harness (13110, -0.39 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.49 DPS) [crafted]; Dusky Leather Armor (7374, -0.54 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.26 DPS) [dungeon]; Dusky Bracers (7378, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.2 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.19 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.21 DPS) [crafted]; Imperial Leather Gloves (4063, -0.24 DPS) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.03 DPS) | yes | Highlander's Chain Girdle (20090, -0.21 DPS) [rep]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.34 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 32.9 attack_power points (1.13 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.24 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.38 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.2 attack_power points (0.83 DPS) | yes | Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.11 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 20.5 attack_power points (0.71 DPS) | yes | Assault Band (13095, -0.02 DPS) [world_drop]; Ironspine's Eye (7686, -0.08 DPS) [dungeon]; Field Researcher's Loop (281634, -0.09 DPS) [quest] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ironspine's Eye (7686, -0.07 DPS) [dungeon]; Field Researcher's Loop (281634, -0.07 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.50 DPS) [world_drop]; Vanquisher's Sword (10823, -1.12 DPS) [quest] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (109.9 DPS) | yes | Stonecloth Branch (15963, -15.15 DPS) [world_drop]; Satyr's Rod (15962, -15.20 DPS) [world_drop]; Ardent Custodian (868, -63.29 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.10 DPS) [world_drop]; Monolithic Bow (9426, -0.11 DPS) [dungeon]; Bow of Searing Arrows (2825, -0.52 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 176.2. Weights run: 2.1s. Verify run: 2.0s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.603 ± 0.012, crit=0.995 ± 0.018 per rating point (14 rating = 1%, 13.924 per %), hit=2.074 ± 0.067 per rating point (10 rating = 1%, 20.742 per %), melee_haste=7.447 ± 1.180

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 50.7 attack_power points (4.15 DPS) | yes | Ebon Mask (19984, -0.06 DPS) [quest]; Embrace of the Lycan (9479, -0.87 DPS) [dungeon]; White Bandit Mask (10008, -1.80 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.8 attack_power points (2.12 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.48 DPS) [quest]; Sentinel's Medallion (19539, -0.54 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 42.7 attack_power points (3.49 DPS) | yes | Skulker's Leather Shoulder (252535, -1.31 DPS) [crafted]; Failed Flying Experiment (9647, -1.36 DPS) [quest]; Sunburn Spaulders (274751, -1.55 DPS, sim-verified) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.4 attack_power points (2.33 DPS) | yes | Blisterbane Wrap (12552, -0.36 DPS) [dungeon]; Dark Phantom Cape (13122, -0.36 DPS) [world_drop]; Duskbat Drape (19982, -0.49 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 52.7 attack_power points (4.31 DPS) | yes | Blazewind Breastplate (11193, -1.05 DPS) [quest]; Warbear Harness (15064, -1.05 DPS) [crafted]; Fungus Shroud Armor (17742, -1.49 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 28.0 attack_power points (2.30 DPS) | yes | Pridelord Bands (14672, -0.57 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.64 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.65 DPS, sim-verified) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 45.3 attack_power points (3.71 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.93 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.13 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.11 DPS) | yes | Skulker's Leather Waistguard (252474, -0.29 DPS) [crafted]; Highlander's Leather Girdle (20115, -0.33 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.39 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 52.7 attack_power points (4.31 DPS) | yes | Serpentskin Leggings (8262, -1.15 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.56 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.59 DPS, sim-verified) [quest] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 33.3 attack_power points (2.72 DPS) | yes | Albino Crocscale Boots (17728, -0.10 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.12 DPS) [crafted]; Prowler's Leather Boots (252468, -0.21 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.7 attack_power points (3.34 DPS) | yes | Masons Fraternity Ring (9533, -1.50 DPS) [quest]; Mark of Kern (2262, -1.70 DPS) [dungeon]; Assault Band (13095, -1.70 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.4 attack_power points (2.00 DPS) | yes | Masons Fraternity Ring (9533, -0.16 DPS) [quest]; Mark of Kern (2262, -0.36 DPS) [dungeon]; Assault Band (13095, -0.36 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadowblade (2163, -0.68 DPS) [world_drop]; Hanzo Sword (8190, -1.40 DPS, sim-verified) [world_drop]; Doomforged Straightedge (12535, -2.48 DPS) [dungeon] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (176.2 DPS) | yes | Shadowblade (2163, -1.98 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.63 DPS) [dungeon]; Thermotastic Egg Timer (9644, -42.72 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.45 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.45 DPS) [world_drop]; Dark Iron Rifle (16004, -2.17 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 289.9. Weights run: 2.0s. Verify run: 6.1s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.056 ± 0.022, crit=2.203 ± 0.038 per rating point (14 rating = 1%, 30.841 per %), hit=4.472 ± 0.146 per rating point (10 rating = 1%, 44.722 per %), melee_haste=not significant (9.491 ± 2.425)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 120.3 attack_power points (10.74 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -0.78 DPS) [pvp]; Knight-Lieutenant's Leather Headband (220850, -1.51 DPS, sim-verified) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 75.6 attack_power points (6.75 DPS) | yes | Beads of Ogre Might (22150, -0.61 DPS) [quest]; Mark of Fordring (15411, -1.67 DPS) [quest]; Medallion of the Dawn (22659, -1.85 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (289.9 DPS) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -17.44 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.7 attack_power points (6.49 DPS) | yes | Cape of the Black Baron (13340, -1.39 DPS, sim-verified) [dungeon]; Arcanoweave Cloak (272411, -2.50 DPS) [vendor]; Stalwart Cloak (272415, -2.50 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (289.9 DPS) | yes | Darkmantle Tunic (226825, -1.86 DPS) [quest]; Field Marshal's Leather Chestpiece (231543, -2.02 DPS) [pvp]; Tunic of Undead Slaying (23089, -12.36 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (289.9 DPS) | yes | Blackmist Armguards (12966, -0.03 DPS) [dungeon]; Marshal's Leather Armsplints (16460, -0.09 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -4.98 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 75.6 attack_power points (6.75 DPS) | yes | Marshal's Leather Handgrips (231544, -0.32 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.49 DPS) [crafted]; Raider Gloves (272099, -19.84 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 132.6 attack_power points (11.84 DPS) | yes | Belt of Preserved Heads (20216, -2.93 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -6.05 DPS) [rep]; Ferocity of the Timbermaw (227805, -6.40 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (289.9 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -16.92 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 80.7 attack_power points (7.20 DPS) | yes | Shadowcraft Boots (16711, -2.15 DPS) [dungeon]; Fine Dawn Treaders (227815, -2.46 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -2.79 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (289.9 DPS) | yes | Tarnished Elven Ring (18500, -1.43 DPS) [dungeon]; Cutthroat's Signet (272408, -1.61 DPS) [vendor]; Naglering (11669, -7.47 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (289.9 DPS) | yes | Tarnished Elven Ring (18500, -0.55 DPS) [dungeon]; Cutthroat's Signet (272408, -0.73 DPS) [vendor]; Naglering (11669, -6.84 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (289.9 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (289.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.91 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (289.9 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -6.25 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 834.8 attack_power points (74.52 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -5.69 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -22.80 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (289.9 DPS) | yes | Blackcrow (12651, -0.40 DPS) [dungeon]; The Purifier (22656, -1.79 DPS) [quest]; Dark Iron Rifle (16004, -3.94 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 01532310421501000-31530300001500000-0020000000000000000)

Set DPS (verified): 740.0. Weights run: 2.2s. Verify run: 6.2s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.456 ± 0.027, crit=2.549 ± 0.043 per rating point (14 rating = 1%, 35.691 per %), hit=4.492 ± 0.230 per rating point (10 rating = 1%, 44.917 per %), melee_haste=15.529 ± 3.211

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 125.5 attack_power points (23.40 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted]; Lieutenant Commander's Leather Helm (227055, -1.66 DPS) [pvp] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 81.8 attack_power points (15.24 DPS) | yes | Beads of Ogre Might (22150, -2.39 DPS) [quest]; Mark of Fordring (15411, -3.74 DPS) [quest]; Medallion of the Dawn (22659, -4.11 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (740.0 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -43.07 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.9 attack_power points (13.59 DPS) | yes | Cape of the Black Baron (13340, -3.00 DPS) [dungeon]; Cloak of the Honor Guard (20073, -4.97 DPS) [rep]; Windshear Cape (20691, -5.08 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (740.0 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -5.04 DPS) [pvp]; Darkmantle Tunic (226825, -5.21 DPS) [quest]; Tunic of Undead Slaying (23089, -28.19 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (740.0 DPS) | yes | Marshal's Leather Armsplints (16460, -0.21 DPS) [pvp]; Blackmist Armguards (12966, -1.56 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -12.76 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 80.6 attack_power points (15.03 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -2.49 DPS) [quest]; Raider Gloves (272099, -49.04 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 141.4 attack_power points (26.36 DPS) | yes | Belt of Preserved Heads (20216, -8.25 DPS) [quest]; Ferocity of the Timbermaw (227805, -13.25 DPS) [vendor]; Highlander's Leather Girdle (20045, -13.37 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (740.0 DPS) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -40.30 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 90.4 attack_power points (16.85 DPS) | yes | Shadowcraft Boots (16711, -4.72 DPS) [dungeon]; Fine Dawn Treaders (227815, -4.77 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -5.86 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (740.0 DPS) | yes | Tarnished Elven Ring (18500, -2.77 DPS) [dungeon]; Cutthroat's Signet (272408, -3.23 DPS) [vendor]; Naglering (11669, -13.10 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (740.0 DPS) | yes | Tarnished Elven Ring (18500, -1.37 DPS) [dungeon]; Cutthroat's Signet (272408, -1.83 DPS) [vendor]; Naglering (11669, -12.08 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (740.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (740.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (740.0 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -7.20 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 839.7 attack_power points (156.55 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -11.55 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -48.52 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (740.0 DPS) | yes | Blackcrow (12651, -0.84 DPS) [dungeon]; The Purifier (22656, -3.09 DPS) [quest]; Dark Iron Rifle (16004, -4.40 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 39.9. Weights run: 1.5s. Verify run: 2.6s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.032 ± 0.013, crit=0.603 ± 0.012 per rating point (14 rating = 1%, 8.438 per %), hit=0.880 ± 0.030 per rating point (10 rating = 1%, 8.802 per %), melee_haste=4.857 ± 0.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.3 attack_power points (0.94 DPS) | yes | Defender's Leather Hood (252447, -0.44 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.2 attack_power points (0.70 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS, sim-verified) [quest] |
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

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 64.1. Weights run: 1.6s. Verify run: 1.4s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.737 ± 0.012, crit=0.680 ± 0.013 per rating point (14 rating = 1%, 9.516 per %), hit=1.066 ± 0.036 per rating point (10 rating = 1%, 10.656 per %), melee_haste=5.299 ± 0.419

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.4 attack_power points (1.02 DPS) | yes | Tribal Worg Helm (6204, -0.20 DPS) [world]; Brawler's Leather Hood (252504, -0.20 DPS) [crafted]; Humbert's Helm (4724, -0.31 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.83 DPS) | yes | Scout's Medallion (19537, -0.01 DPS) [rep]; Kaleidoscope Chain (13084, -0.18 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (1.42 DPS) | yes | Mantle of Thieves (2264, -0.53 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.61 DPS) [crafted]; Bristlebark Amice (14573, -0.63 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.2 attack_power points (0.89 DPS) | yes | Tigerstrike Mantle (13108, -0.07 DPS) [world_drop]; Cloak of Night (4447, -0.28 DPS) [world]; Swiftrunner Cape (6745, -0.28 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.3 attack_power points (1.43 DPS) | yes | Brawler's Leather Tunic (252508, -0.38 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.39 DPS) [quest]; Brawler's Leather Armor (252490, -0.42 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.4 attack_power points (0.85 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.20 DPS) [crafted]; Insignia Bracers (6410, -0.24 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 17.4 attack_power points (1.03 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.08 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.42 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.14 DPS) [crafted]; Prowler's Leather Belt (252459, -0.27 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.53 DPS) | yes | Petrolspill Leggings (9509, -0.10 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.10 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.20 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.9 attack_power points (1.00 DPS) | yes | Insignia Boots (4055, -0.18 DPS) [world_drop]; Vorrel's Boots (7751, -0.18 DPS) [quest]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 attack_power points (1.16 DPS) | yes | Thunderbrow Ring (13097, -0.38 DPS) [world_drop]; Monkey Ring (6748, -0.44 DPS) [quest]; Pyrewood Signet Ring (277210, -0.50 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.4 attack_power points (0.97 DPS) | yes | Thunderbrow Ring (13097, -0.19 DPS) [world_drop]; Monkey Ring (6748, -0.25 DPS) [quest]; Pyrewood Signet Ring (277210, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (19.35 DPS) | yes | Swinetusk Shank (6691, -0.35 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.46 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.56 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (19.03 DPS) | yes | Swinetusk Shank (6691, -14.66 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -18.91 DPS) [quest]; Satyr's Rod (15962, -18.93 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Silver Star (3463, -0.02 DPS) [quest]; Double-barreled Shotgun (2098, -0.05 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.12 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 32502110551401001-00000000000000000-0000000000000000000)

Set DPS (verified): 109.1. Weights run: 1.5s. Verify run: 1.5s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.566 ± 0.013, crit=0.729 ± 0.017 per rating point (14 rating = 1%, 10.208 per %), hit=2.147 ± 0.072 per rating point (10 rating = 1%, 21.474 per %), melee_haste=26.807 ± 0.730

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.2 attack_power points (0.97 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.27 DPS) [world_drop]; Nightscape Headband (8176, -0.32 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.69 DPS) | yes | Scout's Medallion (19536, -0.10 DPS) [rep]; Ghostshard Talisman (7731, -0.21 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.2 attack_power points (1.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Flintrock Shoulders (7755, -0.29 DPS) [dungeon]; Nightscape Shoulders (8192, -0.41 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.7 attack_power points (0.68 DPS) | yes | First Sergeant's Cloak (16340, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Parachute Cloak (10518, -0.25 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 37.8 attack_power points (1.30 DPS) | yes | Wolffear Harness (13110, -0.41 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.49 DPS) [crafted]; Dusky Leather Armor (7374, -0.54 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Dusky Bracers (7378, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.2 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.19 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.21 DPS) [crafted]; Imperial Leather Gloves (4063, -0.24 DPS) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.03 DPS) | yes | Defiler's Chain Girdle (20152, -0.21 DPS) [rep]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.34 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 32.9 attack_power points (1.13 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.24 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.38 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.2 attack_power points (0.83 DPS) | yes | Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.11 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 20.5 attack_power points (0.71 DPS) | yes | Assault Band (13095, -0.02 DPS) [world_drop]; Ironspine's Eye (7686, -0.08 DPS) [dungeon]; Field Researcher's Loop (281634, -0.09 DPS) [quest] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ironspine's Eye (7686, -0.07 DPS) [dungeon]; Field Researcher's Loop (281634, -0.07 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.50 DPS) [world_drop]; Vanquisher's Sword (10823, -1.12 DPS) [quest] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (109.1 DPS) | yes | Stonecloth Branch (15963, -15.15 DPS) [world_drop]; Tork Wrench (11855, -15.19 DPS) [quest]; Ardent Custodian (868, -62.12 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.10 DPS) [world_drop]; Monolithic Bow (9426, -0.11 DPS) [dungeon]; Bow of Searing Arrows (2825, -0.52 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 177.5. Weights run: 2.1s. Verify run: 2.1s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.603 ± 0.012, crit=0.995 ± 0.018 per rating point (14 rating = 1%, 13.924 per %), hit=2.074 ± 0.067 per rating point (10 rating = 1%, 20.742 per %), melee_haste=7.447 ± 1.180

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 49.9 attack_power points (4.09 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -1.17 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.74 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.8 attack_power points (2.12 DPS) | yes | Woven Ivy Necklace (19159, -0.44 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.48 DPS) [quest]; Scout's Medallion (19535, -0.54 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.6 attack_power points (2.43 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.25 DPS) [crafted]; Failed Flying Experiment (9647, -0.30 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 28.4 attack_power points (2.33 DPS) | yes | Blisterbane Wrap (12552, -0.36 DPS) [dungeon]; Dark Phantom Cape (13122, -0.36 DPS) [world_drop]; Duskbat Drape (19982, -0.49 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 52.7 attack_power points (4.31 DPS) | yes | Blazewind Breastplate (11193, -1.05 DPS) [quest]; Warbear Harness (15064, -1.05 DPS) [crafted]; Fungus Shroud Armor (17742, -2.87 DPS, sim-verified) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 28.0 attack_power points (2.30 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.54 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 45.3 attack_power points (3.71 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.93 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.13 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.11 DPS) | yes | Skulker's Leather Waistguard (252474, -0.29 DPS) [crafted]; Defiler's Leather Girdle (20193, -0.33 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.39 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 52.7 attack_power points (4.31 DPS) | yes | Basilisk Hide Pants (1718, -1.56 DPS) [world_drop]; Triprunner Dungarees (9624, -1.70 DPS) [quest]; Serpentskin Leggings (8262, -3.09 DPS, sim-verified) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 33.3 attack_power points (2.72 DPS) | yes | Albino Crocscale Boots (17728, -0.10 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.12 DPS) [crafted]; Prowler's Leather Boots (252468, -0.21 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.7 attack_power points (3.34 DPS) | yes | White Bone Band (11862, -1.37 DPS) [quest]; Masons Fraternity Ring (9533, -1.50 DPS) [quest]; Mark of Kern (2262, -1.70 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.4 attack_power points (2.00 DPS) | yes | White Bone Band (11862, -0.04 DPS) [quest]; Masons Fraternity Ring (9533, -0.16 DPS) [quest]; Mark of Kern (2262, -0.36 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (177.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (177.5 DPS) | yes | Molten Heart of the Mountain (249470, -2.91 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (177.5 DPS) | yes | Thorium Cestus (250614, -2.20 DPS) [crafted]; Hanzo Sword (8190, -2.22 DPS, sim-verified) [world_drop]; Doomforged Straightedge (12535, -2.48 DPS) [dungeon] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (44.64 DPS) | yes | Thorium Cestus (250614, +0.00 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -5.15 DPS) [dungeon]; White Bone Shredder (11863, -7.77 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (177.5 DPS) | yes | Stinging Bow (10624, -0.45 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.45 DPS) [world_drop]; Dark Iron Rifle (16004, -2.20 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 290.0. Weights run: 2.0s. Verify run: 5.9s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.056 ± 0.022, crit=2.203 ± 0.038 per rating point (14 rating = 1%, 30.841 per %), hit=4.472 ± 0.146 per rating point (10 rating = 1%, 44.722 per %), melee_haste=not significant (9.491 ± 2.425)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 120.3 attack_power points (10.74 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS) [crafted]; Champion's Leather Helm (227057, -0.78 DPS) [pvp] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 75.6 attack_power points (6.75 DPS) | yes | Beads of Ogre Might (22150, -0.61 DPS) [quest]; Mark of Fordring (15411, -1.67 DPS) [quest]; Medallion of the Dawn (22659, -1.85 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (290.0 DPS) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -17.14 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.7 attack_power points (6.49 DPS) | yes | Cape of the Black Baron (13340, -1.95 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.50 DPS) [vendor]; Stalwart Cloak (272415, -2.50 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (290.0 DPS) | yes | Darkmantle Tunic (226825, -1.86 DPS) [quest]; Warlord's Leather Breastplate (231549, -2.02 DPS) [pvp]; Tunic of Undead Slaying (23089, -11.52 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (290.0 DPS) | yes | Blackmist Armguards (12966, -0.03 DPS) [dungeon]; General's Leather Armsplints (16559, -0.09 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -4.88 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 75.6 attack_power points (6.75 DPS) | yes | General's Leather Mitts (231555, -0.32 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.49 DPS) [crafted]; Raider Gloves (272099, -19.12 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 132.6 attack_power points (11.84 DPS) | yes | Belt of Preserved Heads (20216, -1.95 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -6.05 DPS) [rep]; Ferocity of the Timbermaw (227805, -6.40 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (290.0 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -15.91 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 80.7 attack_power points (7.20 DPS) | yes | Fine Dawn Treaders (227815, -1.52 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -2.15 DPS) [dungeon]; Darkmantle Boots (22003, -2.79 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (290.0 DPS) | yes | Tarnished Elven Ring (18500, -1.43 DPS) [dungeon]; Cutthroat's Signet (272408, -1.61 DPS) [vendor]; Naglering (11669, -6.36 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (290.0 DPS) | yes | Tarnished Elven Ring (18500, -0.55 DPS) [dungeon]; Cutthroat's Signet (272408, -0.73 DPS) [vendor]; Naglering (11669, -5.76 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (290.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (290.0 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -1.64 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -2.95 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (290.0 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -4.64 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 834.8 attack_power points (74.52 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -6.59 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -22.80 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (290.0 DPS) | yes | Blackcrow (12651, -0.40 DPS) [dungeon]; The Purifier (22656, -1.79 DPS) [quest]; Dark Iron Rifle (16004, -2.98 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 01532310421501000-31530300001500000-0020000000000000000)

Set DPS (verified): 737.0. Weights run: 2.2s. Verify run: 6.2s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.456 ± 0.027, crit=2.549 ± 0.043 per rating point (14 rating = 1%, 35.691 per %), hit=4.492 ± 0.230 per rating point (10 rating = 1%, 44.917 per %), melee_haste=15.529 ± 3.211

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 125.5 attack_power points (23.40 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted]; Champion's Leather Helm (227057, -1.66 DPS) [pvp] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 81.8 attack_power points (15.24 DPS) | yes | Beads of Ogre Might (22150, -2.39 DPS) [quest]; Mark of Fordring (15411, -3.74 DPS) [quest]; Medallion of the Dawn (22659, -4.11 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (737.0 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -45.09 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 72.9 attack_power points (13.59 DPS) | yes | Cape of the Black Baron (13340, -3.34 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -4.97 DPS) [rep]; Windshear Cape (20691, -5.08 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (737.0 DPS) | yes | Warlord's Leather Breastplate (231549, -5.04 DPS) [pvp]; Darkmantle Tunic (226825, -5.21 DPS) [quest]; Tunic of Undead Slaying (23089, -33.19 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (737.0 DPS) | yes | General's Leather Armsplints (16559, -0.21 DPS) [pvp]; Blackmist Armguards (12966, -1.56 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -12.71 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (737.0 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -2.49 DPS) [quest]; Raider Gloves (272099, -49.30 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 141.4 attack_power points (26.36 DPS) | yes | Belt of Preserved Heads (20216, -8.35 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -13.25 DPS) [vendor]; Defiler's Leather Girdle (20190, -13.37 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (737.0 DPS) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -42.77 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 90.4 attack_power points (16.85 DPS) | yes | Shadowcraft Boots (16711, -4.72 DPS) [dungeon]; Darkmantle Boots (22003, -5.86 DPS) [quest]; Fine Dawn Treaders (227815, -6.09 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (737.0 DPS) | yes | Tarnished Elven Ring (18500, -2.77 DPS) [dungeon]; Cutthroat's Signet (272408, -3.23 DPS) [vendor]; Naglering (11669, -18.81 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (737.0 DPS) | yes | Tarnished Elven Ring (18500, -1.37 DPS) [dungeon]; Cutthroat's Signet (272408, -1.83 DPS) [vendor]; Naglering (11669, -17.89 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (737.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (737.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (737.0 DPS) | yes | Shadowsong's Sorrow (21522, +0.00 DPS) [quest]; High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 839.7 attack_power points (156.55 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -10.60 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -48.52 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (737.0 DPS) | yes | Blackcrow (12651, -0.84 DPS) [dungeon]; The Purifier (22656, -3.09 DPS) [quest]; Dark Iron Rifle (16004, -10.58 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

