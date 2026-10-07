# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.4. Weights run: 2.5s. Verify run: 1.7s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.007 ± 0.001, crit=0.016 ± 0.001 per rating point (14 rating = 1%, 0.218 per %), hit=0.030 ± 0.001 per rating point (10 rating = 1%, 0.299 per %), melee_haste=not significant (0.899 ± 0.364)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.38 DPS) | yes | Defender's Leather Hood (252447, -0.00 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.0 attack_power points (0.28 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 attack_power points (0.24 DPS) | yes | Slime-encrusted Pads (6461, -0.32 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.0 attack_power points (0.57 DPS) | yes | Tunic of Westfall (2041, -0.05 DPS) [quest]; Defender's Leather Armor (252434, -0.10 DPS) [crafted]; Prospector's Chestpiece (14562, -0.19 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.0 attack_power points (0.28 DPS) | yes | Forest Leather Bracers (3202, -0.05 DPS) [world_drop]; Bristlebark Bindings (14569, -0.05 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.0 attack_power points (0.47 DPS) | yes | Brawler's Leather Gloves (252494, -0.09 DPS) [crafted]; Bristlebark Gloves (14572, -0.13 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.85 DPS) | yes | Brawler's Leather Belt (252428, -0.47 DPS) [crafted]; Ruffian Belt (5975, -0.57 DPS) [world]; Deviate Scale Belt (6468, -0.57 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.1 attack_power points (0.71 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.10 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.52 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.24 DPS) [dungeon]; Footpads of the Fang (10411, -0.24 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.0 attack_power points (0.38 DPS) | yes | Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Demon Band (12054, -0.19 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.0 attack_power points (0.28 DPS) | yes | Demon Band (12054, -0.10 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.12 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (10.77 DPS) | yes | Evocator's Blade (2567, -0.67 DPS) [dungeon]; Buzzer Blade (2169, -1.54 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.15 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.4 attack_power points (10.58 DPS) | yes | Evocator's Blade (2567, -6.40 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.0 attack_power points (0.19 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 54.0. Weights run: 2.5s. Verify run: 1.8s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.012 ± 0.001, crit=0.023 ± 0.002 per rating point (14 rating = 1%, 0.319 per %), hit=0.048 ± 0.001 per rating point (10 rating = 1%, 0.482 per %), melee_haste=1.221 ± 0.169

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Helm (252512, -0.09 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Kaleidoscope Chain (13084, -0.28 DPS) [world_drop]; Sentinel's Medallion (19541, -0.33 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.1 attack_power points (0.77 DPS) | yes | Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Bristlebark Amice (14573, -0.34 DPS) [world_drop]; Mantle of Thieves (2264, -0.40 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.1 attack_power points (0.48 DPS) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.10 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.76 DPS) | yes | Dusky Leather Armor (7374, -0.09 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.09 DPS) [crafted]; Brawler's Leather Armor (252490, -0.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.1 attack_power points (0.48 DPS) | yes | Cultist's Armguards (270032, -0.00 DPS) [quest]; Jurassic Wristguards (6198, -0.10 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.76 DPS) | yes | Insignia Gloves (6408, -0.14 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.19 DPS) [crafted]; Wolfclaw Gloves (1978, -0.24 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.15 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.42 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.24 DPS) | yes | Brawler's Leather Pants (252500, -0.52 DPS) [crafted]; Trapper's Leather Pants (252501, -0.52 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.59 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, -0.04 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.09 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.1 attack_power points (0.63 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Monkey Ring (6748, -0.29 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.1 attack_power points (0.58 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Monkey Ring (6748, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (15.37 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.05 DPS) [vendor]; Torturing Poker (7682, -1.22 DPS) [dungeon]; Thornspike (6681, -1.61 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (15.27 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -7.24 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -15.22 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Double-barreled Shotgun (2098, -0.14 DPS) [world_drop]; Silver Star (3463, -0.19 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 104.1. Weights run: 2.7s. Verify run: 1.8s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.022 ± 0.002, crit=0.047 ± 0.004 per rating point (14 rating = 1%, 0.665 per %), hit=0.066 ± 0.002 per rating point (10 rating = 1%, 0.659 per %), melee_haste=not significant (1.117 ± 0.308)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.2 attack_power points (1.12 DPS) | yes | Warden's Wizard Hat (14604, -0.45 DPS) [world_drop]; Nightscape Headband (8176, -0.50 DPS) [crafted]; Hawkeye's Helm (14591, -0.54 DPS, sim-verified) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.01 DPS) | yes | Sentinel's Medallion (19540, -0.44 DPS) [rep]; Ghostshard Talisman (7731, -0.47 DPS, sim-verified) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.2 attack_power points (1.17 DPS) | yes | Flintrock Shoulders (7755, -0.40 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.54 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.2 attack_power points (0.71 DPS) | yes | Sergeant Major's Cape (16336, -0.10 DPS) [pvp]; Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Wolfmaster Cape (6314, -0.21 DPS) [dungeon] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.4 attack_power points (1.38 DPS) | yes | Raptorbane Armor (3566, -0.57 DPS) [quest]; Nightscape Tunic (8175, -0.61 DPS) [crafted]; Wolffear Harness (13110, -0.84 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.01 DPS) | yes | Cultist's Armguards (270032, -0.50 DPS) [quest]; Hawkeye's Bracers (14590, -0.56 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.7 attack_power points (1.04 DPS) | yes | Prowler's Leather Gloves (252524, -0.07 DPS) [crafted]; Imperial Leather Gloves (4063, -0.12 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.84 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.51 DPS) | yes | Highlander's Chain Girdle (20090, -0.47 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.59 DPS) [world_drop]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.31 DPS) | yes | Basilisk Hide Pants (1718, -0.23 DPS) [world_drop]; Triprunner Dungarees (9624, -0.23 DPS) [quest]; Brawler's Leather Legguards (252516, -0.49 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.2 attack_power points (0.92 DPS) | yes | Prowler's Leather Shoes (252465, -0.00 DPS) [crafted]; Imperial Leather Boots (6431, -0.10 DPS) [dungeon]; Excelsior Boots (4109, -0.10 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.01 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.01 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (104.1 DPS) | yes | Black Menace (6831, -2.97 DPS) [quest]; Darkspear Insurgent's Spellblade (272085, -3.49 DPS) [vendor]; Coldrage Dagger (10761, -6.31 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 438.8 attack_power points (22.06 DPS) | yes | Black Menace (6831, -1.89 DPS, sim-verified) [quest]; Stonecloth Branch (15963, -21.90 DPS) [world_drop]; Satyr's Rod (15962, -22.00 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (104.1 DPS) | yes | Monolithic Bow (9426, -0.25 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Bow of Searing Arrows (2825, -1.09 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 145.6. Weights run: 2.7s. Verify run: 2.2s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.274 ± 0.015, crit=0.564 ± 0.026 per rating point (14 rating = 1%, 7.894 per %), hit=0.085 ± 0.003 per rating point (10 rating = 1%, 0.852 per %), melee_haste=not significant (1.725 ± 0.486)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.12 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -0.80 DPS) [crafted]; Knight-Lieutenant's Leather Headband (220850, -0.81 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 21.6 attack_power points (1.14 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.08 DPS) [quest]; Sentinel's Medallion (19539, -0.33 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.0 attack_power points (1.38 DPS) | yes | Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted]; Failed Flying Experiment (9647, -0.17 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.19 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 23.8 attack_power points (1.26 DPS) | yes | Blisterbane Wrap (12552, -0.25 DPS) [dungeon]; Dark Phantom Cape (13122, -0.25 DPS) [world_drop]; Duskbat Drape (19982, -0.32 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 33.9 attack_power points (1.80 DPS) | yes | Blazewind Breastplate (11193, -0.09 DPS) [quest]; Quillward Harness (10583, -0.09 DPS) [dungeon]; Mixologist's Tunic (12793, -0.10 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 23.1 attack_power points (1.23 DPS) | yes | Skulker's Leather Bracers (252540, -0.25 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.28 DPS) [crafted]; Branded Leather Bracers (19508, -0.58 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 38.0 attack_power points (2.02 DPS) | yes | Darkmantle Grips (226828, -0.05 DPS) [vendor]; Gloves of Holy Might (867, -0.54 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.59 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.02 DPS) | yes | Skulker's Leather Waistguard (252474, -0.43 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.46 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.69 DPS, sim-verified) [rep] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 34.1 attack_power points (1.81 DPS) | yes | Serpentskin Leggings (8262, -0.04 DPS) [world_drop]; Basilisk Hide Pants (1718, -0.39 DPS) [world_drop]; Knight's Leather Pants (220858, -0.39 DPS) [vendor] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 27.7 attack_power points (1.47 DPS) | yes | Skulker's Leather Boots (252469, -0.01 DPS) [crafted]; Prowler's Leather Boots (252468, -0.03 DPS) [crafted]; Albino Crocscale Boots (17728, -0.12 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 21.5 attack_power points (1.14 DPS) | yes | Mark of Kern (2262, -0.08 DPS) [dungeon]; Assault Band (13095, -0.08 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.19 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.9 attack_power points (1.11 DPS) | yes | Assault Band (13095, -0.05 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.16 DPS) [quest]; Mark of Kern (2262, -1.28 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (145.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (145.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (145.6 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -2.22 DPS) [quest]; Charstone Dirk (17710, -2.22 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (27.14 DPS) | yes | Barman Shanker (12791, -17.52 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -26.94 DPS) [quest]; Stonecloth Branch (15963, -26.98 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (145.6 DPS) | yes | Stinging Bow (10624, -0.04 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.04 DPS) [world_drop]; Dark Iron Rifle (16004, -1.81 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Gryphon Rider's Leggings; feet: Sandstalker Ankleguards; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 230.8. Weights run: 2.7s. Verify run: 2.2s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.389 ± 0.022, crit=0.793 ± 0.038 per rating point (14 rating = 1%, 11.103 per %), hit=0.126 ± 0.004 per rating point (10 rating = 1%, 1.256 per %), melee_haste=not significant (2.840 ± 0.725)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Darkmantle Cap (226829) | Saving the Best for Last [quest] | 45.8 attack_power points (2.43 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.00 DPS) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 36.4 attack_power points (1.93 DPS) | yes | Medallion of the Dawn (22659, -0.07 DPS) [quest]; Imperial Jewel (11933, -0.23 DPS) [dungeon]; Will of the Martyr (17044, -0.34 DPS) [quest] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 55.0 attack_power points (2.92 DPS) | yes | Darkspear Pauldrons (272105, -0.34 DPS) [vendor]; Highlander's Lizardhide Shoulders (20060, -0.44 DPS) [rep]; Dark Warder's Pauldrons (22241, -1.08 DPS) [dungeon] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 40.9 attack_power points (2.18 DPS) | yes | Cape of the Black Baron (13340, -0.01 DPS) [dungeon]; Howler's Furs (272414, -0.62 DPS) [vendor]; Windshear Cape (20691, -0.64 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.53 DPS) [crafted]; Dawn Armor (252483, -0.82 DPS) [crafted]; Tunic of Undead Slaying (23089, -7.40 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.03 DPS) [rep]; Marshal's Leather Armsplints (16460, -0.08 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.76 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 48.5 attack_power points (2.58 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.14 DPS) [pvp]; Darkmantle Gloves (22006, -0.32 DPS) [quest]; Cadaverous Gloves (14640, -1.34 DPS, sim-verified) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 49.6 attack_power points (2.64 DPS) | yes | Marshal's Leather Cinch (16458, -0.15 DPS) [pvp]; Highlander's Leather Girdle (20045, -0.24 DPS) [rep]; Might of the Timbermaw (19044, -0.49 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 59.7 attack_power points (3.17 DPS) | yes | Warbear Woolies (15065, -0.36 DPS) [crafted]; Cadaverous Leggings (14638, -0.41 DPS) [dungeon]; Devilsaur Leggings (15062, -2.28 DPS, sim-verified) [crafted] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (2.12 DPS) | yes | Drudge Boots (21532, -0.18 DPS) [quest]; Dunestalker's Boots (20715, -0.24 DPS) [quest]; Darkmantle Footpads (226831, -0.31 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.11 DPS) [quest]; Blackstone Ring (17713, -0.38 DPS) [dungeon]; Naglering (11669, -6.78 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.05 DPS) [quest]; Blackstone Ring (17713, -0.32 DPS) [dungeon]; Naglering (11669, -2.83 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18465, -4.77 DPS, sim-verified) [quest] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (230.8 DPS) | yes | The Lobotomizer (19324, -2.67 DPS, sim-verified) [rep]; Distracting Dagger (18392, -3.58 DPS) [dungeon]; Tome of Knowledge (13385, -32.89 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.14 DPS) [world_drop]; Dark Iron Rifle (16004, -2.00 DPS, sim-verified) [crafted] |

**New at 60:** head: Darkmantle Cap; neck: Amulet of the Darkmoon; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Sentinel's Leather Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Smolderweb's Eye; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Riphook

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 489.5. Weights run: 2.8s. Verify run: 2.2s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.537 ± 0.021, crit=1.138 ± 0.037 per rating point (14 rating = 1%, 15.936 per %), hit=0.215 ± 0.005 per rating point (10 rating = 1%, 2.151 per %), melee_haste=6.819 ± 0.835

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 56.2 attack_power points (5.37 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.17 DPS) [quest]; Lieutenant Commander's Leather Helm (227055, -0.20 DPS) [pvp] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 39.9 attack_power points (3.82 DPS) | yes | Amulet of the Darkmoon (19491, -0.07 DPS) [quest]; Imperial Jewel (11933, -0.76 DPS) [dungeon]; Will of the Martyr (17044, -0.95 DPS) [quest] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 57.7 attack_power points (5.52 DPS) | yes | Darkspear Pauldrons (272105, -0.02 DPS) [vendor]; Highlander's Lizardhide Shoulders (20060, -0.88 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, -1.68 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 43.1 attack_power points (4.12 DPS) | yes | Cloak of the Honor Guard (20073, -0.13 DPS) [rep]; Windshear Cape (20691, -1.15 DPS) [world]; Howler's Furs (272414, -1.23 DPS) [vendor] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (489.5 DPS) | yes | Timbermaw Tunic (252484, -0.35 DPS) [crafted]; Dawn Armor (252483, -0.51 DPS) [crafted]; Tunic of Undead Slaying (23089, -13.79 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (489.5 DPS) | yes | Bracers of the Eclipse (18375, -0.08 DPS) [dungeon]; Marshal's Leather Armsplints (16460, -0.10 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -7.67 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 52.5 attack_power points (5.02 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.49 DPS) [pvp]; Marshal's Leather Handgrips (231544, -0.56 DPS) [pvp]; Darkmantle Gloves (22006, -0.64 DPS) [quest] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 52.1 attack_power points (4.99 DPS) | yes | Highlander's Leather Girdle (20045, -0.21 DPS) [rep]; Marshal's Leather Cinch (16458, -0.29 DPS) [pvp]; Might of the Timbermaw (19044, -0.92 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 73.4 attack_power points (7.02 DPS) | yes | Marshal's Leather Leggings (231548, -1.11 DPS) [pvp]; Warbear Woolies (15065, -1.69 DPS) [crafted]; Devilsaur Leggings (15062, -4.49 DPS, sim-verified) [crafted] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (3.83 DPS) | yes | Drudge Boots (21532, -0.08 DPS) [quest]; Darkmantle Footpads (226831, -0.15 DPS) [vendor]; Dunestalker's Boots (20715, -0.17 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (489.5 DPS) | yes | Protector's Band (19514, -0.50 DPS) [rep]; Naglering (11669, -12.49 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (489.5 DPS) | yes | Protector's Band (19514, -0.09 DPS) [rep]; Naglering (11669, -11.32 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (489.5 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-verified (489.5 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -6.80 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (489.5 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Felstriker (12590, -15.09 DPS, sim-verified) [dungeon] |
| off_hand | The Lobotomizer (19324) | Stormpike Guard [rep] | 661.1 attack_power points (63.23 DPS) | yes | Felstriker (12590, +0.00 DPS, sim-verified) [dungeon]; Distracting Dagger (18392, -8.64 DPS) [dungeon]; Tome of Knowledge (13385, -61.29 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (489.5 DPS) | yes | Precisely Calibrated Boomstick (2100, -0.05 DPS) [world_drop]; Malgen's Long Bow (22318, -0.19 DPS) [dungeon]; Dark Iron Rifle (16004, -3.67 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Medallion of the Dawn; shoulder: Highlander's Leather Shoulders; back: Cape of the Black Baron; chest: Cadaverous Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Sentinel's Leather Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Royal Seal of Eldre'Thalas; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer; ranged: Riphook

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.1. Weights run: 2.5s. Verify run: 1.7s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.007 ± 0.001, crit=0.016 ± 0.001 per rating point (14 rating = 1%, 0.218 per %), hit=0.030 ± 0.001 per rating point (10 rating = 1%, 0.299 per %), melee_haste=not significant (0.899 ± 0.364)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.38 DPS) | yes | Defender's Leather Hood (252447, -0.00 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.0 attack_power points (0.28 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 attack_power points (0.24 DPS) | yes | Slime-encrusted Pads (6461, -0.31 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.0 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.0 attack_power points (0.57 DPS) | yes | Defender's Leather Armor (252434, -0.13 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.19 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.19 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.0 attack_power points (0.24 DPS) | yes | Bristlebark Bindings (14569, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.05 DPS) [vendor]; Ratchet Wristwraps (274742, -0.09 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.0 attack_power points (0.47 DPS) | yes | Brawler's Leather Gloves (252494, -0.09 DPS) [crafted]; Bristlebark Gloves (14572, -0.13 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.85 DPS) | yes | Brawler's Leather Belt (252428, -0.47 DPS) [crafted]; Deviate Scale Belt (6468, -0.56 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.57 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.1 attack_power points (0.71 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.10 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.52 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.24 DPS) [dungeon]; Footpads of the Fang (10411, -0.24 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 attack_power points (0.38 DPS) | yes | Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Demon Band (12054, -0.19 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.24 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.0 attack_power points (0.28 DPS) | yes | Demon Band (12054, -0.10 DPS) [world_drop]; Pyrewood Signet Ring (277210, -0.12 DPS, sim-verified) [quest]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (10.77 DPS) | yes | Edward's Knife (251485, -0.28 DPS) [quest]; Scout's Blade (20441, -0.50 DPS) [pvp]; Evocator's Blade (2567, -0.67 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.4 attack_power points (10.58 DPS) | yes | Edward's Knife (251485, -0.17 DPS, sim-verified) [quest]; Tork Wrench (11855, -10.48 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.0 attack_power points (0.19 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 53.7. Weights run: 2.5s. Verify run: 1.8s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.012 ± 0.001, crit=0.023 ± 0.002 per rating point (14 rating = 1%, 0.319 per %), hit=0.048 ± 0.001 per rating point (10 rating = 1%, 0.482 per %), melee_haste=1.221 ± 0.169

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Helm (252512, -0.09 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Kaleidoscope Chain (13084, -0.28 DPS) [world_drop]; Scout's Medallion (19537, -0.35 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.1 attack_power points (0.77 DPS) | yes | Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Bristlebark Amice (14573, -0.34 DPS) [world_drop]; Mantle of Thieves (2264, -0.40 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.1 attack_power points (0.48 DPS) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.2 attack_power points (0.68 DPS) | yes | Brawler's Leather Tunic (252508, -0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.10 DPS) [crafted]; Defender's Leather Tunic (252450, -0.10 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.1 attack_power points (0.48 DPS) | yes | Cultist's Armguards (270032, -0.00 DPS) [quest]; Jurassic Wristguards (6198, -0.10 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.76 DPS) | yes | Insignia Gloves (6408, -0.14 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.19 DPS) [crafted]; Wolfclaw Gloves (1978, -0.24 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.15 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Deftkin Belt (16659, -0.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.24 DPS) | yes | Brawler's Leather Pants (252500, -0.52 DPS) [crafted]; Trapper's Leather Pants (252501, -0.52 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.61 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.53 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Insignia Boots (4055, -0.14 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.14 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.1 attack_power points (0.63 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Band of the Fist (17694, -0.24 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.1 attack_power points (0.58 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Band of the Fist (17694, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (15.37 DPS) | yes | Scout's Blade (19545, -0.88 DPS) [pvp]; Darkspear Insurgent's Spellblade (272086, -1.05 DPS) [vendor]; Torturing Poker (7682, -1.22 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (15.27 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -7.38 DPS, sim-verified) [vendor]; Tork Wrench (11855, -15.18 DPS) [quest]; Satyr's Rod (15962, -15.22 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Double-barreled Shotgun (2098, -0.18 DPS, sim-verified) [world_drop]; Silver Star (3463, -0.19 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 102.8. Weights run: 2.7s. Verify run: 1.8s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.022 ± 0.002, crit=0.047 ± 0.004 per rating point (14 rating = 1%, 0.665 per %), hit=0.066 ± 0.002 per rating point (10 rating = 1%, 0.659 per %), melee_haste=not significant (1.117 ± 0.308)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.2 attack_power points (1.12 DPS) | yes | Warden's Wizard Hat (14604, -0.45 DPS) [world_drop]; Nightscape Headband (8176, -0.50 DPS) [crafted]; Hawkeye's Helm (14591, -0.54 DPS, sim-verified) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.01 DPS) | yes | Scout's Medallion (19536, -0.44 DPS) [rep]; Ghostshard Talisman (7731, -0.46 DPS, sim-verified) [dungeon]; Ethereal Talisman (4430, -0.55 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.2 attack_power points (1.17 DPS) | yes | Flintrock Shoulders (7755, -0.40 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.54 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.2 attack_power points (0.71 DPS) | yes | First Sergeant's Cloak (16340, -0.10 DPS) [pvp]; Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Wildhunter Cloak (16658, -0.21 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.4 attack_power points (1.38 DPS) | yes | Nightscape Tunic (8175, -0.61 DPS) [crafted]; Barbaric Harness (5739, -0.62 DPS) [crafted]; Wolffear Harness (13110, -0.84 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.01 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Cultist's Armguards (270032, -0.50 DPS) [quest]; Hawkeye's Bracers (14590, -0.58 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.7 attack_power points (1.04 DPS) | yes | Prowler's Leather Gloves (252524, -0.07 DPS) [crafted]; Imperial Leather Gloves (4063, -0.12 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.81 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.51 DPS) | yes | Defiler's Chain Girdle (20152, -0.46 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.59 DPS) [world_drop]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.31 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.23 DPS) [quest]; Brawler's Leather Legguards (252516, -0.49 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.2 attack_power points (0.92 DPS) | yes | Prowler's Leather Shoes (252465, -0.00 DPS) [crafted]; Imperial Leather Boots (6431, -0.10 DPS) [dungeon]; Excelsior Boots (4109, -0.10 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.01 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.01 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (102.8 DPS) | yes | Scout's Blade (19544, -2.95 DPS) [pvp]; Darkspear Insurgent's Spellblade (272085, -3.49 DPS) [vendor]; Coldrage Dagger (10761, -5.87 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 438.8 attack_power points (22.06 DPS) | yes | Coldrage Dagger (10761, -0.84 DPS, sim-verified) [dungeon]; Stonecloth Branch (15963, -21.90 DPS) [world_drop]; Tork Wrench (11855, -21.96 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (102.8 DPS) | yes | Monolithic Bow (9426, -0.25 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Bow of Searing Arrows (2825, -1.07 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 146.3. Weights run: 2.7s. Verify run: 2.2s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.274 ± 0.015, crit=0.564 ± 0.026 per rating point (14 rating = 1%, 7.894 per %), hit=0.085 ± 0.003 per rating point (10 rating = 1%, 0.852 per %), melee_haste=not significant (1.725 ± 0.486)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.12 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -0.80 DPS) [crafted]; Blood Guard's Leather Headband (220851, -0.81 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 21.6 attack_power points (1.14 DPS) | yes | Woven Ivy Necklace (19159, -0.22 DPS) [quest]; Scout's Medallion (19535, -0.33 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.52 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.0 attack_power points (1.38 DPS) | yes | Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted]; Failed Flying Experiment (9647, -0.17 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.19 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 23.8 attack_power points (1.26 DPS) | yes | Blisterbane Wrap (12552, -0.25 DPS) [dungeon]; Dark Phantom Cape (13122, -0.25 DPS) [world_drop]; Duskbat Drape (19982, -0.32 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 33.9 attack_power points (1.80 DPS) | yes | Blazewind Breastplate (11193, -0.09 DPS) [quest]; Quillward Harness (10583, -0.09 DPS) [dungeon]; Mixologist's Tunic (12793, -0.10 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 23.1 attack_power points (1.23 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.69 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 38.0 attack_power points (2.02 DPS) | yes | Darkmantle Grips (226828, -0.05 DPS) [vendor]; Gloves of Holy Might (867, -0.54 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.59 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.02 DPS) | yes | Skulker's Leather Waistguard (252474, -0.43 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.46 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.68 DPS, sim-verified) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 33.4 attack_power points (1.77 DPS) | yes | Basilisk Hide Pants (1718, -0.35 DPS) [world_drop]; Stone Guard's Leather Pants (220859, -0.35 DPS) [vendor]; Ferine Leggings (6690, -0.39 DPS) [dungeon] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 27.7 attack_power points (1.47 DPS) | yes | Skulker's Leather Boots (252469, -0.01 DPS) [crafted]; Prowler's Leather Boots (252468, -0.03 DPS) [crafted]; Albino Crocscale Boots (17728, -0.12 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.27 DPS) | yes | Blackstone Ring (17713, -0.17 DPS) [dungeon]; Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 21.5 attack_power points (1.14 DPS) | yes | Blackstone Ring (17713, +0.00 DPS, sim-verified) [dungeon]; Mark of Kern (2262, -0.08 DPS) [dungeon]; Assault Band (13095, -0.08 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (146.3 DPS) | yes | Frozen Heart of the Mountain (249469, -3.57 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (146.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (146.3 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Scout's Blade (19543, -2.18 DPS) [pvp]; Lifeforce Dirk (10750, -2.22 DPS) [quest] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (27.14 DPS) | yes | Barman Shanker (12791, -16.49 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -26.94 DPS) [quest]; Stonecloth Branch (15963, -26.98 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (146.3 DPS) | yes | Stinging Bow (10624, -0.04 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.04 DPS) [world_drop]; Dark Iron Rifle (16004, -1.91 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Serpentskin Leggings; feet: Sandstalker Ankleguards; finger1: White Bone Band; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 231.6. Weights run: 2.7s. Verify run: 2.2s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.389 ± 0.022, crit=0.793 ± 0.038 per rating point (14 rating = 1%, 11.103 per %), hit=0.126 ± 0.004 per rating point (10 rating = 1%, 1.256 per %), melee_haste=not significant (2.840 ± 0.725)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Darkmantle Cap (226829) | Saving the Best for Last [quest] | 45.8 attack_power points (2.43 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.00 DPS) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 36.4 attack_power points (1.93 DPS) | yes | Medallion of the Dawn (22659, -0.07 DPS) [quest]; Imperial Jewel (11933, -0.23 DPS) [dungeon]; Will of the Martyr (17044, -0.34 DPS) [quest] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 55.0 attack_power points (2.92 DPS) | yes | Darkspear Pauldrons (272105, -0.34 DPS) [vendor]; Defiler's Lizardhide Shoulders (20175, -0.44 DPS) [rep]; Dark Warder's Pauldrons (22241, -1.08 DPS) [dungeon] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 40.9 attack_power points (2.18 DPS) | yes | Cape of the Black Baron (13340, -0.01 DPS) [dungeon]; Howler's Furs (272414, -0.62 DPS) [vendor]; Windshear Cape (20691, -0.64 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.53 DPS) [crafted]; Dawn Armor (252483, -0.82 DPS) [crafted]; Tunic of Undead Slaying (23089, -7.46 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.03 DPS) [rep]; General's Leather Armsplints (16559, -0.08 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.80 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 48.5 attack_power points (2.58 DPS) | yes | Blood Guard's Leather Vices (16499, -0.14 DPS) [pvp]; Darkmantle Gloves (22006, -0.32 DPS) [quest]; Cadaverous Gloves (14640, -1.47 DPS, sim-verified) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 49.6 attack_power points (2.64 DPS) | yes | General's Leather Girdle (16557, -0.15 DPS) [pvp]; Defiler's Leather Girdle (20190, -0.24 DPS) [rep]; Might of the Timbermaw (19044, -0.49 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 59.7 attack_power points (3.17 DPS) | yes | Warbear Woolies (15065, -0.36 DPS) [crafted]; Cadaverous Leggings (14638, -0.41 DPS) [dungeon]; Devilsaur Leggings (15062, -2.43 DPS, sim-verified) [crafted] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (2.12 DPS) | yes | Drudge Boots (21532, -0.18 DPS) [quest]; Dunestalker's Boots (20715, -0.24 DPS) [quest]; Darkmantle Footpads (226831, -0.31 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.11 DPS) [quest]; White Bone Band (11862, -0.23 DPS) [quest]; Naglering (11669, -6.77 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.05 DPS) [quest]; White Bone Band (11862, -0.17 DPS) [quest]; Naglering (11669, -2.89 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Counterattack Lodestone (18537, -1.11 DPS) [dungeon]; Hand of Justice (11815, -1.22 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18465, -1.39 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (231.6 DPS) | yes | The Lobotomizer (19324, -2.68 DPS, sim-verified) [rep]; Distracting Dagger (18392, -3.58 DPS) [dungeon]; Tome of Knowledge (13385, -32.89 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.14 DPS) [world_drop]; Dark Iron Rifle (16004, -1.99 DPS, sim-verified) [crafted] |

**New at 60:** head: Darkmantle Cap; neck: Amulet of the Darkmoon; shoulder: Defiler's Leather Shoulders; back: Deathguard's Cloak; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Sentinel's Leather Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Riphook

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 495.2. Weights run: 2.8s. Verify run: 2.3s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.537 ± 0.021, crit=1.138 ± 0.037 per rating point (14 rating = 1%, 15.936 per %), hit=0.215 ± 0.005 per rating point (10 rating = 1%, 2.151 per %), melee_haste=6.819 ± 0.835

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 56.2 attack_power points (5.37 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.17 DPS) [quest]; Champion's Leather Helm (227057, -0.20 DPS) [pvp] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 39.9 attack_power points (3.82 DPS) | yes | Amulet of the Darkmoon (19491, -0.07 DPS) [quest]; Imperial Jewel (11933, -0.76 DPS) [dungeon]; Will of the Martyr (17044, -0.95 DPS) [quest] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 57.7 attack_power points (5.52 DPS) | yes | Darkspear Pauldrons (272105, -0.02 DPS) [vendor]; Defiler's Lizardhide Shoulders (20175, -0.88 DPS) [rep]; Champion's Leather Shoulders (227056, -1.68 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 43.1 attack_power points (4.12 DPS) | yes | Deathguard's Cloak (20068, -0.13 DPS) [rep]; Windshear Cape (20691, -1.15 DPS) [world]; Howler's Furs (272414, -1.23 DPS) [vendor] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.35 DPS) [crafted]; Dawn Armor (252483, -0.51 DPS) [crafted]; Tunic of Undead Slaying (23089, -13.75 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of the Eclipse (18375, -0.08 DPS) [dungeon]; General's Leather Armsplints (16559, -0.10 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -7.64 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 52.5 attack_power points (5.02 DPS) | yes | Blood Guard's Leather Vices (16499, -0.49 DPS) [pvp]; General's Leather Mitts (231555, -0.56 DPS) [pvp]; Darkmantle Gloves (22006, -0.64 DPS) [quest] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 52.1 attack_power points (4.99 DPS) | yes | Defiler's Leather Girdle (20190, -0.21 DPS) [rep]; General's Leather Girdle (16557, -0.29 DPS) [pvp]; Might of the Timbermaw (19044, -0.92 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 73.4 attack_power points (7.02 DPS) | yes | General's Leather Legguards (231554, -1.11 DPS) [pvp]; Outrider's Leather Pants (22740, -1.38 DPS) [rep]; Devilsaur Leggings (15062, -4.90 DPS, sim-verified) [crafted] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (3.83 DPS) | yes | Drudge Boots (21532, -0.08 DPS) [quest]; Darkmantle Footpads (226831, -0.15 DPS) [vendor]; Dunestalker's Boots (20715, -0.17 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -0.50 DPS) [rep]; White Bone Band (11862, -0.96 DPS) [quest]; Naglering (11669, -9.19 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -0.09 DPS) [rep]; White Bone Band (11862, -0.56 DPS) [quest]; Naglering (11669, -8.06 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+14.0 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (495.2 DPS) | yes | Distracting Dagger (18392, -6.45 DPS) [dungeon]; The Lobotomizer (19324, -7.56 DPS, sim-verified) [rep]; Tome of Knowledge (13385, -59.11 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.05 DPS) [world_drop]; Malgen's Long Bow (22318, -0.19 DPS) [dungeon]; Dark Iron Rifle (16004, -3.66 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Medallion of the Dawn; shoulder: Defiler's Leather Shoulders; back: Cape of the Black Baron; chest: Cadaverous Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Sentinel's Leather Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkmoon Card: Maelstrom; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Riphook

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

