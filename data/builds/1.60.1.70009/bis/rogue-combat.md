# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.0. Weights run: 1.4s. Verify run: 1.0s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.010 ± 0.003, crit=0.021 ± 0.002 per rating point (14 rating = 1%, 0.288 per %), hit=0.029 ± 0.001 per rating point (10 rating = 1%, 0.292 per %), melee_haste=not significant (0.486 ± 0.357)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.38 DPS) | yes | Defender's Leather Hood (252447, -0.00 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 attack_power points (0.24 DPS) | yes | Slime-encrusted Pads (6461, -0.34 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.1 attack_power points (0.57 DPS) | yes | Tunic of Westfall (2041, -0.05 DPS) [quest]; Defender's Leather Armor (252434, -0.10 DPS) [crafted]; Prospector's Chestpiece (14562, -0.19 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.0 attack_power points (0.28 DPS) | yes | Forest Leather Bracers (3202, -0.05 DPS) [world_drop]; Bristlebark Bindings (14569, -0.05 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.1 attack_power points (0.47 DPS) | yes | Brawler's Leather Gloves (252494, -0.10 DPS) [crafted]; Bristlebark Gloves (14572, -0.13 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.85 DPS) | yes | Brawler's Leather Belt (252428, -0.47 DPS) [crafted]; Ruffian Belt (5975, -0.57 DPS) [world]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.1 attack_power points (0.71 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.10 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.52 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.24 DPS) [dungeon]; Footpads of the Fang (10411, -0.24 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.0 attack_power points (0.38 DPS) | yes | Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Demon Band (12054, -0.19 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.29 DPS) | yes | Pyrewood Signet Ring (277210, -0.09 DPS) [quest]; Demon Band (12054, -0.10 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.75 DPS) | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Barrens Basher (274744, -1.11 DPS) [vendor]; Diamond Hammer (2194, -2.13 DPS, sim-verified) [world_drop] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (35.0 DPS) | yes | Diamond Hammer (2194, -0.41 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.0 attack_power points (0.19 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 55.8. Weights run: 1.5s. Verify run: 0.9s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.008 ± 0.002, crit=0.033 ± 0.003 per rating point (14 rating = 1%, 0.466 per %), hit=0.042 ± 0.002 per rating point (10 rating = 1%, 0.418 per %), melee_haste=not significant (0.788 ± 0.553)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.58 DPS) | yes | Brawler's Leather Helm (252512, -0.09 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Kaleidoscope Chain (13084, -0.29 DPS) [world_drop]; Sentinel's Medallion (20444, -0.38 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.1 attack_power points (0.77 DPS) | yes | Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Bristlebark Amice (14573, -0.34 DPS) [world_drop]; Mantle of Thieves (2264, -0.44 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.1 attack_power points (0.48 DPS) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.10 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.10 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.77 DPS) | yes | Dusky Leather Armor (7374, -0.09 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.09 DPS) [crafted]; Brawler's Leather Armor (252490, -0.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.0 attack_power points (0.48 DPS) | yes | Cultist's Armguards (270032, -0.00 DPS) [quest]; Jurassic Wristguards (6198, -0.10 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.77 DPS) | yes | Toughened Leather Gloves (4253, -0.19 DPS) [crafted]; Insignia Gloves (6408, -0.19 DPS, sim-verified) [world_drop]; Wolfclaw Gloves (1978, -0.24 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.15 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.43 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.25 DPS) | yes | Brawler's Leather Pants (252500, -0.52 DPS) [crafted]; Trapper's Leather Pants (252501, -0.52 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.68 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.58 DPS) | yes | Feet of the Lynx (1121, -0.04 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.09 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.1 attack_power points (0.63 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Monkey Ring (6748, -0.29 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.0 attack_power points (0.58 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Monkey Ring (6748, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (15.72 DPS) | yes | Swinetusk Shank (6691, -0.28 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.38 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.46 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.46 DPS) | yes | Satyr's Rod (15962, -15.42 DPS) [world_drop]; Swinetusk Shank (6691, -17.55 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Silver Star (3463, -0.19 DPS) [quest]; Double-barreled Shotgun (2098, -0.20 DPS, sim-verified) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 102.6. Weights run: 1.6s. Verify run: 0.9s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.025 ± 0.006, crit=0.061 ± 0.007 per rating point (14 rating = 1%, 0.861 per %), hit=0.057 ± 0.005 per rating point (10 rating = 1%, 0.574 per %), melee_haste=not significant (0.414 ± 1.195)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.3 attack_power points (1.10 DPS) | yes | Warden's Wizard Hat (14604, -0.44 DPS) [world_drop]; Hawkeye's Helm (14591, -0.47 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.49 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.99 DPS) | yes | Ghostshard Talisman (7731, -0.41 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19540, -0.43 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 attack_power points (1.15 DPS) | yes | Flintrock Shoulders (7755, -0.40 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.47 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.2 attack_power points (0.71 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Wolfmaster Cape (6314, -0.21 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.28 DPS, sim-verified) [pvp] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.5 attack_power points (1.36 DPS) | yes | Raptorbane Armor (3566, -0.57 DPS) [quest]; Nightscape Tunic (8175, -0.60 DPS) [crafted]; Wolffear Harness (13110, -0.74 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Cultist's Armguards (270032, -0.50 DPS) [quest]; Hawkeye's Bracers (14590, -0.52 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.58 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.9 attack_power points (1.03 DPS) | yes | Prowler's Leather Gloves (252524, -0.08 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.65 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.49 DPS) | yes | Highlander's Chain Girdle (20090, -0.41 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.58 DPS) [world_drop]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, -0.22 DPS) [world_drop]; Triprunner Dungarees (9624, -0.23 DPS) [quest]; Brawler's Leather Legguards (252516, -0.48 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.3 attack_power points (0.91 DPS) | yes | Prowler's Leather Shoes (252465, -0.00 DPS) [crafted]; Imperial Leather Boots (6431, -0.10 DPS) [dungeon]; Excelsior Boots (4109, -0.10 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.52 DPS) | yes | Ardent Custodian (868, -0.71 DPS) [world_drop]; Dazzling Longsword (869, -1.68 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.76 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (102.6 DPS) | yes | Stonecloth Branch (15963, -21.76 DPS) [world_drop]; Satyr's Rod (15962, -21.86 DPS) [world_drop]; Ardent Custodian (868, -24.03 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.24 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Bow of Searing Arrows (2825, -0.95 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 155.3. Weights run: 1.5s. Verify run: 1.2s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.178 ± 0.040, crit=0.459 ± 0.019 per rating point (14 rating = 1%, 6.430 per %), hit=0.074 ± 0.007 per rating point (10 rating = 1%, 0.744 per %), melee_haste=not significant (0.037 ± 1.743)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.14 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -0.86 DPS) [crafted]; Knight-Lieutenant's Leather Headband (220850, -0.90 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.3 attack_power points (1.09 DPS) | yes | Sentinel's Medallion (19539, -0.33 DPS) [rep]; Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Zealous Shadowshard Pendant (17772, -0.59 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.0 attack_power points (1.33 DPS) | yes | Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted]; Failed Flying Experiment (9647, -0.17 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.18 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.5 attack_power points (1.20 DPS) | yes | Blisterbane Wrap (12552, -0.26 DPS) [dungeon]; Dark Phantom Cape (13122, -0.26 DPS) [world_drop]; Duskbat Drape (19982, -0.32 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 32.2 attack_power points (1.72 DPS) | yes | Mixologist's Tunic (12793, -0.07 DPS) [dungeon]; Quillward Harness (10583, -0.10 DPS) [dungeon]; Blazewind Breastplate (11193, -0.11 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 21.7 attack_power points (1.16 DPS) | yes | Skulker's Leather Bracers (252540, -0.22 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.24 DPS) [crafted]; Branded Leather Bracers (19508, -0.79 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 35.9 attack_power points (1.92 DPS) | yes | Darkmantle Grips (226828, -0.05 DPS) [vendor]; Gloves of Holy Might (867, -0.51 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.56 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.03 DPS) | yes | Skulker's Leather Waistguard (252474, -0.51 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.53 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.83 DPS, sim-verified) [rep] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 32.7 attack_power points (1.75 DPS) | yes | Serpentskin Leggings (8262, -0.04 DPS) [world_drop]; Ferine Leggings (6690, -0.36 DPS) [dungeon]; Knight's Leather Pants (220858, -0.40 DPS) [vendor] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.3 attack_power points (1.41 DPS) | yes | Sandstalker Ankleguards (12470, -0.02 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.15 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.7 attack_power points (1.11 DPS) | yes | Mark of Kern (2262, -0.04 DPS) [dungeon]; Assault Band (13095, -0.04 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.23 DPS) [quest] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 20.6 attack_power points (1.10 DPS) | yes | Mark of Kern (2262, -0.03 DPS) [dungeon]; Assault Band (13095, -0.03 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.22 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (155.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (155.3 DPS) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-verified (155.3 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hanzo Sword (8190, -4.75 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (29.15 DPS) | yes | Claw of Celebras (17738, -3.36 DPS) [dungeon]; Thorium Cestus (250614, -4.92 DPS, sim-verified) [crafted]; Thermotastic Egg Timer (9644, -28.96 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (155.3 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.03 DPS) [world_drop]; Dark Iron Rifle (16004, -1.76 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Bloodrazor; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 236.6. Weights run: 1.5s. Verify run: 1.2s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.379 ± 0.097, crit=0.589 ± 0.028 per rating point (14 rating = 1%, 8.246 per %), hit=0.079 ± 0.012 per rating point (10 rating = 1%, 0.789 per %), melee_haste=not significant (3.357 ± 2.841)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Darkmantle Cap (226829) | Saving the Best for Last [quest] | 42.7 attack_power points (2.29 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.07 DPS) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 36.2 attack_power points (1.94 DPS) | yes | Medallion of the Dawn (22659, -0.21 DPS) [quest]; Imperial Jewel (11933, -0.23 DPS) [dungeon]; Will of the Martyr (17044, -0.33 DPS) [quest] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 54.8 attack_power points (2.94 DPS) | yes | Darkspear Pauldrons (272105, -0.50 DPS) [vendor]; Highlander's Lizardhide Shoulders (20060, -1.07 DPS, sim-verified) [rep]; Dark Warder's Pauldrons (22241, -1.09 DPS) [dungeon] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 40.9 attack_power points (2.19 DPS) | yes | Cape of the Black Baron (13340, -0.01 DPS) [dungeon]; Howler's Furs (272414, -0.65 DPS) [vendor]; Windshear Cape (20691, -0.65 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.69 DPS) [crafted]; Dawn Armor (252483, -1.02 DPS) [crafted]; Tunic of Undead Slaying (23089, -8.24 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.03 DPS) [rep]; Marshal's Leather Armsplints (16460, -0.08 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -4.16 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 48.2 attack_power points (2.59 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.13 DPS) [pvp]; Darkmantle Gloves (22006, -0.32 DPS) [quest]; Cadaverous Gloves (14640, -1.49 DPS, sim-verified) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 49.4 attack_power points (2.65 DPS) | yes | Marshal's Leather Cinch (16458, -0.15 DPS) [pvp]; Highlander's Leather Girdle (20045, -0.39 DPS) [rep]; Might of the Timbermaw (19044, -0.49 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | sim-verified (236.6 DPS) | yes | Warbear Woolies (15065, -0.05 DPS) [crafted]; Cadaverous Leggings (14638, -0.09 DPS) [dungeon]; Devilsaur Leggings (15062, -2.37 DPS, sim-verified) [crafted] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (2.14 DPS) | yes | Drudge Boots (21532, -0.19 DPS) [quest]; Dunestalker's Boots (20715, -0.26 DPS) [quest]; Darkmantle Footpads (226831, -0.34 DPS) [vendor] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -0.11 DPS) [rep]; Blackstone Ring (17713, -0.34 DPS) [dungeon]; Naglering (11669, -3.14 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -0.03 DPS) [rep]; Blackstone Ring (17713, -0.26 DPS) [dungeon]; Naglering (11669, -6.32 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18465, -3.66 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -1.01 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.03 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 828.7 attack_power points (44.41 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -9.50 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -13.36 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.14 DPS) [world_drop]; Dark Iron Rifle (16004, -2.21 DPS, sim-verified) [crafted] |

**New at 60:** head: Darkmantle Cap; neck: Amulet of the Darkmoon; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Sentinel's Leather Pants; feet: Pads of the Dread Wolf; finger1: Protector's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.3. Weights run: 1.4s. Verify run: 0.9s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.010 ± 0.003, crit=0.021 ± 0.002 per rating point (14 rating = 1%, 0.288 per %), hit=0.029 ± 0.001 per rating point (10 rating = 1%, 0.292 per %), melee_haste=not significant (0.486 ± 0.357)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.38 DPS) | yes | Defender's Leather Hood (252447, -0.00 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 attack_power points (0.24 DPS) | yes | Slime-encrusted Pads (6461, -0.32 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.1 attack_power points (0.57 DPS) | yes | Defender's Leather Armor (252434, -0.13 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.19 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.19 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.0 attack_power points (0.24 DPS) | yes | Bristlebark Bindings (14569, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.05 DPS) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.1 attack_power points (0.47 DPS) | yes | Brawler's Leather Gloves (252494, -0.10 DPS) [crafted]; Bristlebark Gloves (14572, -0.13 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.85 DPS) | yes | Brawler's Leather Belt (252428, -0.47 DPS) [crafted]; Ruffian Belt (5975, -0.57 DPS) [world]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.1 attack_power points (0.71 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.10 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.52 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.24 DPS) [dungeon]; Footpads of the Fang (10411, -0.24 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 attack_power points (0.38 DPS) | yes | Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Demon Band (12054, -0.19 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.24 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.29 DPS) | yes | Demon Band (12054, -0.10 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.13 DPS, sim-verified) [quest]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.75 DPS) | yes | Cruel Barb (5191, -0.91 DPS) [dungeon]; Blackfang (2236, -0.96 DPS) [world_drop]; Barrens Basher (274744, -1.11 DPS) [vendor] |
| off_hand | Diamond Hammer (2194) | World drop [world_drop] | 229.8 attack_power points (10.85 DPS) | yes | Cruel Barb (5191, +0.00 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -10.75 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.0 attack_power points (0.19 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Diamond Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 55.0. Weights run: 1.5s. Verify run: 1.0s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.008 ± 0.002, crit=0.033 ± 0.003 per rating point (14 rating = 1%, 0.466 per %), hit=0.042 ± 0.002 per rating point (10 rating = 1%, 0.418 per %), melee_haste=not significant (0.788 ± 0.553)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.58 DPS) | yes | Brawler's Leather Helm (252512, -0.09 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Kaleidoscope Chain (13084, -0.29 DPS) [world_drop]; Scout's Medallion (20442, -0.38 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.1 attack_power points (0.77 DPS) | yes | Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Bristlebark Amice (14573, -0.34 DPS) [world_drop]; Mantle of Thieves (2264, -0.42 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.1 attack_power points (0.48 DPS) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest]; Tigerstrike Mantle (13108, -0.10 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 attack_power points (0.68 DPS) | yes | Brawler's Leather Tunic (252508, -0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.10 DPS) [crafted]; Defender's Leather Tunic (252450, -0.10 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.0 attack_power points (0.48 DPS) | yes | Cultist's Armguards (270032, -0.00 DPS) [quest]; Jurassic Wristguards (6198, -0.10 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.77 DPS) | yes | Toughened Leather Gloves (4253, -0.19 DPS) [crafted]; Insignia Gloves (6408, -0.21 DPS, sim-verified) [world_drop]; Wolfclaw Gloves (1978, -0.24 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.15 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Deftkin Belt (16659, -0.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.25 DPS) | yes | Brawler's Leather Pants (252500, -0.52 DPS) [crafted]; Trapper's Leather Pants (252501, -0.52 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.70 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.53 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Insignia Boots (4055, -0.14 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.14 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.1 attack_power points (0.63 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Band of the Fist (17694, -0.24 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.0 attack_power points (0.58 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Band of the Fist (17694, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 327.9 attack_power points (15.72 DPS) | yes | Swinetusk Shank (6691, -0.28 DPS) [dungeon]; Scorn's Focal Dagger (23168, -0.38 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.46 DPS) [dungeon] |
| off_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.46 DPS) | yes | Tork Wrench (11855, -15.37 DPS) [quest]; Satyr's Rod (15962, -15.42 DPS) [world_drop]; Swinetusk Shank (6691, -17.27 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Silver Star (3463, -0.19 DPS) [quest]; Double-barreled Shotgun (2098, -0.21 DPS, sim-verified) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Ironspine's Fist; off_hand: Royal Diplomatic Scepter; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 101.6. Weights run: 1.6s. Verify run: 1.0s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.025 ± 0.006, crit=0.061 ± 0.007 per rating point (14 rating = 1%, 0.861 per %), hit=0.057 ± 0.005 per rating point (10 rating = 1%, 0.574 per %), melee_haste=not significant (0.414 ± 1.195)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.3 attack_power points (1.10 DPS) | yes | Warden's Wizard Hat (14604, -0.44 DPS) [world_drop]; Hawkeye's Helm (14591, -0.47 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.49 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.99 DPS) | yes | Ghostshard Talisman (7731, -0.40 DPS, sim-verified) [dungeon]; Scout's Medallion (19536, -0.43 DPS) [rep]; Ethereal Talisman (4430, -0.54 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 attack_power points (1.15 DPS) | yes | Flintrock Shoulders (7755, -0.40 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.47 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.2 attack_power points (0.71 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Wildhunter Cloak (16658, -0.21 DPS) [quest]; First Sergeant's Cloak (16340, -0.27 DPS, sim-verified) [pvp] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.5 attack_power points (1.36 DPS) | yes | Nightscape Tunic (8175, -0.60 DPS) [crafted]; Barbaric Harness (5739, -0.61 DPS) [crafted]; Wolffear Harness (13110, -0.74 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.45 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.50 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.9 attack_power points (1.03 DPS) | yes | Prowler's Leather Gloves (252524, -0.08 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.66 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.49 DPS) | yes | Defiler's Chain Girdle (20152, -0.40 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.58 DPS) [world_drop]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.23 DPS) [quest]; Brawler's Leather Legguards (252516, -0.48 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.3 attack_power points (0.91 DPS) | yes | Prowler's Leather Shoes (252465, -0.00 DPS) [crafted]; Imperial Leather Boots (6431, -0.10 DPS) [dungeon]; Excelsior Boots (4109, -0.10 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.52 DPS) | yes | Ardent Custodian (868, -0.71 DPS) [world_drop]; Dazzling Longsword (869, -1.68 DPS) [world_drop]; Jhordy's Misplaced Screwdriver (274753, -1.76 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (101.6 DPS) | yes | Stonecloth Branch (15963, -21.76 DPS) [world_drop]; Tork Wrench (11855, -21.81 DPS) [quest]; Ardent Custodian (868, -24.17 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Monolithic Bow (9426, -0.24 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Bow of Searing Arrows (2825, -0.93 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 157.5. Weights run: 1.5s. Verify run: 1.1s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.178 ± 0.040, crit=0.459 ± 0.019 per rating point (14 rating = 1%, 6.430 per %), hit=0.074 ± 0.007 per rating point (10 rating = 1%, 0.744 per %), melee_haste=not significant (0.037 ± 1.743)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.14 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -0.86 DPS) [crafted]; Blood Guard's Leather Headband (220851, -0.90 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.3 attack_power points (1.09 DPS) | yes | Woven Ivy Necklace (19159, -0.20 DPS) [quest]; Scout's Medallion (19535, -0.33 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.49 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.0 attack_power points (1.33 DPS) | yes | Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted]; Failed Flying Experiment (9647, -0.17 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.18 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.5 attack_power points (1.20 DPS) | yes | Blisterbane Wrap (12552, -0.26 DPS) [dungeon]; Dark Phantom Cape (13122, -0.26 DPS) [world_drop]; Duskbat Drape (19982, -0.32 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 32.2 attack_power points (1.72 DPS) | yes | Mixologist's Tunic (12793, -0.07 DPS) [dungeon]; Quillward Harness (10583, -0.10 DPS) [dungeon]; Blazewind Breastplate (11193, -0.11 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 21.7 attack_power points (1.16 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.70 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 35.9 attack_power points (1.92 DPS) | yes | Darkmantle Grips (226828, -0.05 DPS) [vendor]; Gloves of Holy Might (867, -0.51 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.56 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.03 DPS) | yes | Skulker's Leather Waistguard (252474, -0.51 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.53 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.82 DPS, sim-verified) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 31.8 attack_power points (1.70 DPS) | yes | Stone Guard's Leather Pants (220859, -0.36 DPS) [vendor]; Basilisk Hide Pants (1718, -0.38 DPS) [world_drop]; Ferine Leggings (6690, -1.14 DPS, sim-verified) [dungeon] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.3 attack_power points (1.41 DPS) | yes | Sandstalker Ankleguards (12470, -0.02 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.15 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.28 DPS) | yes | Legionnaire's Band (19511, -0.18 DPS) [rep]; Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.7 attack_power points (1.11 DPS) | yes | Mark of Kern (2262, -0.04 DPS) [dungeon]; Assault Band (13095, -0.04 DPS) [world_drop]; Legionnaire's Band (19511, -1.14 DPS, sim-verified) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (157.5 DPS) | yes | Frozen Heart of the Mountain (249469, -4.47 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (157.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-verified (157.5 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hanzo Sword (8190, -4.55 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (29.15 DPS) | yes | Claw of Celebras (17738, -3.36 DPS) [dungeon]; Thorium Cestus (250614, -3.95 DPS, sim-verified) [crafted]; White Bone Shredder (11863, -5.23 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (157.5 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.03 DPS) [world_drop]; Dark Iron Rifle (16004, -1.74 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Bloodrazor; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 233.7. Weights run: 1.5s. Verify run: 1.2s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.379 ± 0.097, crit=0.589 ± 0.028 per rating point (14 rating = 1%, 8.246 per %), hit=0.079 ± 0.012 per rating point (10 rating = 1%, 0.789 per %), melee_haste=not significant (3.357 ± 2.841)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Darkmantle Cap (226829) | Saving the Best for Last [quest] | 42.7 attack_power points (2.29 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.07 DPS) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 36.2 attack_power points (1.94 DPS) | yes | Medallion of the Dawn (22659, -0.21 DPS) [quest]; Imperial Jewel (11933, -0.23 DPS) [dungeon]; Will of the Martyr (17044, -0.33 DPS) [quest] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 54.8 attack_power points (2.94 DPS) | yes | Darkspear Pauldrons (272105, -0.50 DPS) [vendor]; Defiler's Lizardhide Shoulders (20175, -1.07 DPS, sim-verified) [rep]; Dark Warder's Pauldrons (22241, -1.09 DPS) [dungeon] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 40.9 attack_power points (2.19 DPS) | yes | Cape of the Black Baron (13340, -0.01 DPS) [dungeon]; Howler's Furs (272414, -0.65 DPS) [vendor]; Windshear Cape (20691, -0.65 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.69 DPS) [crafted]; Dawn Armor (252483, -1.02 DPS) [crafted]; Tunic of Undead Slaying (23089, -8.20 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.03 DPS) [rep]; General's Leather Armsplints (16559, -0.08 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -4.22 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 48.2 attack_power points (2.59 DPS) | yes | Blood Guard's Leather Vices (16499, -0.13 DPS) [pvp]; Darkmantle Gloves (22006, -0.32 DPS) [quest]; Cadaverous Gloves (14640, -1.60 DPS, sim-verified) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 49.4 attack_power points (2.65 DPS) | yes | General's Leather Girdle (16557, -0.15 DPS) [pvp]; Defiler's Leather Girdle (20190, -0.39 DPS) [rep]; Might of the Timbermaw (19044, -0.49 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | sim-verified (233.7 DPS) | yes | Warbear Woolies (15065, -0.05 DPS) [crafted]; Cadaverous Leggings (14638, -0.09 DPS) [dungeon]; Devilsaur Leggings (15062, -2.41 DPS, sim-verified) [crafted] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (2.14 DPS) | yes | Drudge Boots (21532, -0.19 DPS) [quest]; Dunestalker's Boots (20715, -0.26 DPS) [quest]; Darkmantle Footpads (226831, -0.34 DPS) [vendor] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -0.11 DPS) [rep]; White Bone Band (11862, -0.17 DPS) [quest]; Naglering (11669, -3.20 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -0.03 DPS) [rep]; White Bone Band (11862, -0.09 DPS) [quest]; Naglering (11669, -5.92 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, -3.22 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -1.01 DPS) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 828.7 attack_power points (44.41 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -10.03 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -13.36 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.14 DPS) [world_drop]; Dark Iron Rifle (16004, -2.20 DPS, sim-verified) [crafted] |

**New at 60:** head: Darkmantle Cap; neck: Amulet of the Darkmoon; shoulder: Defiler's Leather Shoulders; back: Deathguard's Cloak; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Sentinel's Leather Pants; feet: Pads of the Dread Wolf; finger1: Legionnaire's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

