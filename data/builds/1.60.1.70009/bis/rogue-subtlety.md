# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 31.3. Weights run: 1.3s. Verify run: 1.4s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.021 ± 0.003 per rating point (14 rating = 1%, 0.290 per %), hit=0.059 ± 0.003 per rating point (10 rating = 1%, 0.591 per %), melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.37 DPS) | yes | Defender's Leather Hood (252447, -0.01 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 attack_power points (0.28 DPS) | yes | Erudite's Amulet (277204, -0.09 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.23 DPS) | yes | Slime-encrusted Pads (6461, -0.28 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.1 attack_power points (0.55 DPS) | yes | Tunic of Westfall (2041, -0.04 DPS) [quest]; Defender's Leather Armor (252434, -0.09 DPS) [crafted]; Prospector's Chestpiece (14562, -0.18 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.1 attack_power points (0.27 DPS) | yes | Forest Leather Bracers (3202, -0.04 DPS) [world_drop]; Bristlebark Bindings (14569, -0.05 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.1 attack_power points (0.46 DPS) | yes | Bristlebark Gloves (14572, -0.09 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.09 DPS) [crafted]; Gold-flecked Gloves (5195, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.82 DPS) | yes | Brawler's Leather Belt (252428, -0.45 DPS) [crafted]; Deviate Scale Belt (6468, -0.54 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.54 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.1 attack_power points (0.69 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.50 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.23 DPS) [dungeon]; Footpads of the Fang (10411, -0.23 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.1 attack_power points (0.36 DPS) | yes | Pyrewood Signet Ring (277210, -0.17 DPS) [quest]; Demon Band (12054, -0.18 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.28 DPS) | yes | Pyrewood Signet Ring (277210, -0.08 DPS) [quest]; Demon Band (12054, -0.09 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (10.35 DPS) | yes | Evocator's Blade (2567, -0.65 DPS) [dungeon]; Buzzer Blade (2169, -1.48 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.07 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.5 attack_power points (10.17 DPS) | yes | Evocator's Blade (2567, -6.42 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 40.1. Weights run: 1.5s. Verify run: 1.6s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.030 ± 0.003 per rating point (14 rating = 1%, 0.421 per %), hit=0.024 ± 0.002 per rating point (10 rating = 1%, 0.238 per %), melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Helm (252512, -0.09 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.65 DPS) | yes | Kaleidoscope Chain (13084, -0.28 DPS) [world_drop]; Sentinel's Medallion (19541, -0.30 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.1 attack_power points (0.75 DPS) | yes | Barbaric Shoulders (5964, -0.28 DPS) [crafted]; Bristlebark Amice (14573, -0.33 DPS) [world_drop]; Mantle of Thieves (2264, -0.34 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.1 attack_power points (0.47 DPS) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.09 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.74 DPS) | yes | Dusky Leather Armor (7374, -0.09 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.09 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.1 attack_power points (0.47 DPS) | yes | Cultist's Armguards (270032, -0.00 DPS) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.09 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.74 DPS) | yes | Insignia Gloves (6408, -0.14 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -0.18 DPS) [crafted]; Wolfclaw Gloves (1978, -0.23 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.12 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.41 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.21 DPS) | yes | Brawler's Leather Pants (252500, -0.51 DPS) [crafted]; Trapper's Leather Pants (252501, -0.51 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.52 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.56 DPS) | yes | Feet of the Lynx (1121, -0.04 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.09 DPS) [crafted]; Insignia Boots (4055, -0.18 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.1 attack_power points (0.61 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.19 DPS) [vendor]; Monkey Ring (6748, -0.28 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.1 attack_power points (0.56 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.14 DPS) [vendor]; Monkey Ring (6748, -0.23 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (14.97 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.02 DPS) [vendor]; Torturing Poker (7682, -1.18 DPS) [dungeon]; Thornspike (6681, -1.57 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (14.88 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.46 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -14.83 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS, sim-verified) [world_drop]; Silver Star (3463, -0.18 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.23 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 82.0. Weights run: 1.6s. Verify run: 1.6s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.047 ± 0.006 per rating point (14 rating = 1%, 0.663 per %), hit=0.058 ± 0.006 per rating point (10 rating = 1%, 0.583 per %), melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.1 attack_power points (1.10 DPS) | yes | Hawkeye's Helm (14591, -0.44 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.44 DPS) [world_drop]; Nightscape Headband (8176, -0.49 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.99 DPS) | yes | Ghostshard Talisman (7731, -0.38 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19540, -0.44 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 attack_power points (1.15 DPS) | yes | Flintrock Shoulders (7755, -0.40 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.44 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.1 attack_power points (0.70 DPS) | yes | Sergeant Major's Cape (16336, -0.10 DPS) [pvp]; Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Wolfmaster Cape (6314, -0.20 DPS) [dungeon] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.2 attack_power points (1.35 DPS) | yes | Raptorbane Armor (3566, -0.56 DPS) [quest]; Nightscape Tunic (8175, -0.60 DPS) [crafted]; Wolffear Harness (13110, -0.71 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Hawkeye's Bracers (14590, -0.39 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.50 DPS) [quest]; Dusky Bracers (7378, -0.59 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.7 attack_power points (1.02 DPS) | yes | Prowler's Leather Gloves (252524, -0.08 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.78 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.49 DPS) | yes | Highlander's Chain Girdle (20090, -0.38 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.59 DPS) [world_drop]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.24 DPS) [quest]; Brawler's Leather Legguards (252516, -0.49 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.1 attack_power points (0.90 DPS) | yes | Prowler's Leather Shoes (252465, -0.00 DPS) [crafted]; Imperial Leather Boots (6431, -0.10 DPS) [dungeon]; Excelsior Boots (4109, -0.10 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Protector's Band (19515, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.50 DPS) | yes | Black Menace (6831, -2.93 DPS) [quest]; Coldrage Dagger (10761, -2.93 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272085, -3.44 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 438.8 attack_power points (21.74 DPS) | yes | Black Menace (6831, -0.42 DPS, sim-verified) [quest]; Stonecloth Branch (15963, -21.59 DPS) [world_drop]; Satyr's Rod (15962, -21.68 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (82.0 DPS) | yes | Monolithic Bow (9426, -0.25 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Bow of Searing Arrows (2825, -0.89 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 118.9. Weights run: 1.6s. Verify run: 1.9s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.218 ± 0.061, crit=0.464 ± 0.025 per rating point (14 rating = 1%, 6.499 per %), hit=0.053 ± 0.009 per rating point (10 rating = 1%, 0.529 per %), melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.17 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -0.84 DPS) [crafted]; Knight-Lieutenant's Leather Headband (220850, -0.92 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.8 attack_power points (1.13 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.05 DPS) [quest]; Sentinel's Medallion (19539, -0.34 DPS) [rep]; Ghostshard Talisman (7731, -0.37 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.4 attack_power points (1.37 DPS) | yes | Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted]; Failed Flying Experiment (9647, -0.17 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.19 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 23.1 attack_power points (1.25 DPS) | yes | Blisterbane Wrap (12552, -0.26 DPS) [dungeon]; Dark Phantom Cape (13122, -0.26 DPS) [world_drop]; Duskbat Drape (19982, -0.32 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 32.9 attack_power points (1.78 DPS) | yes | Mixologist's Tunic (12793, -0.08 DPS) [dungeon]; Quillward Harness (10583, -0.10 DPS) [dungeon]; Blazewind Breastplate (11193, -0.10 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 22.3 attack_power points (1.21 DPS) | yes | Skulker's Leather Bracers (252540, -0.23 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.26 DPS) [crafted]; Branded Leather Bracers (19508, -0.65 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 36.8 attack_power points (1.99 DPS) | yes | Darkmantle Grips (226828, -0.05 DPS) [vendor]; Gloves of Holy Might (867, -0.56 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.58 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.06 DPS) | yes | Skulker's Leather Waistguard (252474, -0.48 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.51 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.61 DPS, sim-verified) [rep] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 33.3 attack_power points (1.80 DPS) | yes | Serpentskin Leggings (8262, -0.04 DPS) [world_drop]; Ferine Leggings (6690, -0.39 DPS) [dungeon]; Basilisk Hide Pants (1718, -0.42 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.8 attack_power points (1.45 DPS) | yes | Sandstalker Ankleguards (12470, -0.01 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 21.0 attack_power points (1.13 DPS) | yes | Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.21 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.5 attack_power points (1.11 DPS) | yes | Assault Band (13095, -0.03 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.19 DPS) [quest]; Mark of Kern (2262, -1.86 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (118.9 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (118.9 DPS) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (118.9 DPS) | yes | Lifeforce Dirk (10750, -0.42 DPS) [quest]; Charstone Dirk (17710, -0.42 DPS) [dungeon]; Shadowblade (2163, -2.95 DPS, sim-verified) [world_drop] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (27.69 DPS) | yes | Thermotastic Egg Timer (9644, -27.50 DPS) [quest]; Stonecloth Branch (15963, -27.53 DPS) [world_drop]; Windchaser Orb (15965, -27.63 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (118.9 DPS) | yes | Stinging Bow (10624, -0.00 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.00 DPS) [world_drop]; Dark Iron Rifle (16004, -1.72 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 183.0. Weights run: 1.6s. Verify run: 2.0s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.181 ± 0.023, crit=0.580 ± 0.036 per rating point (14 rating = 1%, 8.118 per %), hit=not significant (0.046 ± 0.013) per rating point (10 rating = 1%, 0.458 per %), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Darkmantle Cap (226829) | Saving the Best for Last [quest] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Embrace of the Lycan (9479, -2.60 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 32.4 attack_power points (1.72 DPS) | yes | Medallion of the Dawn (22659, -0.02 DPS) [quest]; Imperial Jewel (11933, -0.02 DPS) [dungeon]; Will of the Martyr (17044, -0.13 DPS) [quest] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 51.3 attack_power points (2.72 DPS) | yes | Darkspear Pauldrons (272105, -0.60 DPS) [vendor]; Highlander's Lizardhide Shoulders (20060, -0.80 DPS, sim-verified) [rep]; Dark Warder's Pauldrons (22241, -1.07 DPS) [dungeon] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 39.9 attack_power points (2.12 DPS) | yes | Cape of the Black Baron (13340, -0.12 DPS) [dungeon]; Howler's Furs (272414, -0.61 DPS) [vendor]; Windshear Cape (20691, -0.75 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.80 DPS) [crafted]; Nightbrace Tunic (12603, -1.15 DPS) [dungeon]; Tunic of Undead Slaying (23089, -6.00 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.13 DPS) [rep]; Marshal's Leather Armsplints (16460, -0.18 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.10 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | 44.0 attack_power points (2.34 DPS) | yes | Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Leather Gauntlets (16396, -0.01 DPS) [pvp]; Skul's Fingerbone Claws (13395, -0.21 DPS) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 46.1 attack_power points (2.45 DPS) | yes | Marshal's Leather Cinch (16458, -0.13 DPS) [pvp]; Highlander's Leather Girdle (20045, -0.21 DPS) [rep]; Cadaverous Belt (14636, -0.32 DPS) [dungeon] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 54.1 attack_power points (2.88 DPS) | yes | Cadaverous Leggings (14638, -0.11 DPS) [dungeon]; Warbear Woolies (15065, -0.26 DPS) [crafted]; Sentinel's Leather Pants (237818, -0.32 DPS) [vendor] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (2.13 DPS) | yes | Drudge Boots (21532, -0.37 DPS) [quest]; Dunestalker's Boots (20715, -0.46 DPS) [quest]; Highlander's Leather Boots (20052, -0.52 DPS) [rep] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.17 DPS) [quest]; Blackstone Ring (17713, -0.24 DPS) [dungeon]; Naglering (11669, -2.38 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.15 DPS) [quest]; Blackstone Ring (17713, -0.22 DPS) [dungeon]; Naglering (11669, -5.19 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Hand of Justice (11815, -4.83 DPS, sim-verified) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -0.94 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Distracting Dagger (18392, -3.59 DPS) [dungeon]; The Lobotomizer (19324, -4.91 DPS, sim-verified) [rep]; Tome of Knowledge (13385, -33.00 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.27 DPS) [world_drop]; Dark Iron Rifle (16004, -1.61 DPS, sim-verified) [crafted] |

**New at 60:** head: Darkmantle Cap; neck: Amulet of the Darkmoon; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Cadaverous Gloves; waist: Ferocity of the Timbermaw; legs: Devilsaur Leggings; feet: Pads of the Dread Wolf; finger1: Protector's Band; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Royal Seal of Eldre'Thalas; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Riphook

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.7. Weights run: 1.3s. Verify run: 1.4s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.021 ± 0.003 per rating point (14 rating = 1%, 0.290 per %), hit=0.059 ± 0.003 per rating point (10 rating = 1%, 0.591 per %), melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.37 DPS) | yes | Defender's Leather Hood (252447, -0.01 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 attack_power points (0.28 DPS) | yes | Erudite's Amulet (277204, -0.11 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.23 DPS) | yes | Slime-encrusted Pads (6461, -0.27 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.1 attack_power points (0.55 DPS) | yes | Defender's Leather Armor (252434, -0.11 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.18 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.18 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 attack_power points (0.23 DPS) | yes | Bristlebark Bindings (14569, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.05 DPS) [vendor]; Ratchet Wristwraps (274742, -0.09 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.1 attack_power points (0.46 DPS) | yes | Brawler's Leather Gloves (252494, -0.09 DPS) [crafted]; Bristlebark Gloves (14572, -0.11 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.82 DPS) | yes | Brawler's Leather Belt (252428, -0.45 DPS) [crafted]; Deviate Scale Belt (6468, -0.54 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.54 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.1 attack_power points (0.69 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.50 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.23 DPS) [dungeon]; Footpads of the Fang (10411, -0.23 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.1 attack_power points (0.36 DPS) | yes | Pyrewood Signet Ring (277210, -0.17 DPS) [quest]; Demon Band (12054, -0.18 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.23 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.28 DPS) | yes | Pyrewood Signet Ring (277210, -0.08 DPS) [quest]; Demon Band (12054, -0.09 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (10.35 DPS) | yes | Edward's Knife (251485, -0.27 DPS) [quest]; Scout's Blade (20441, -0.48 DPS) [pvp]; Evocator's Blade (2567, -0.65 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.5 attack_power points (10.17 DPS) | yes | Edward's Knife (251485, +0.00 DPS, sim-verified) [quest]; Tork Wrench (11855, -10.08 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 39.4. Weights run: 1.5s. Verify run: 1.6s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.030 ± 0.003 per rating point (14 rating = 1%, 0.421 per %), hit=0.024 ± 0.002 per rating point (10 rating = 1%, 0.238 per %), melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Helm (252512, -0.09 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.65 DPS) | yes | Kaleidoscope Chain (13084, -0.28 DPS) [world_drop]; Scout's Medallion (19537, -0.33 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.1 attack_power points (0.75 DPS) | yes | Barbaric Shoulders (5964, -0.28 DPS) [crafted]; Bristlebark Amice (14573, -0.33 DPS) [world_drop]; Mantle of Thieves (2264, -0.33 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.1 attack_power points (0.47 DPS) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 attack_power points (0.66 DPS) | yes | Brawler's Leather Tunic (252508, -0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.10 DPS) [crafted]; Defender's Leather Tunic (252450, -0.10 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.1 attack_power points (0.47 DPS) | yes | Cultist's Armguards (270032, -0.00 DPS) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.09 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.74 DPS) | yes | Insignia Gloves (6408, -0.16 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -0.18 DPS) [crafted]; Wolfclaw Gloves (1978, -0.23 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.12 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.37 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.21 DPS) | yes | Brawler's Leather Pants (252500, -0.51 DPS) [crafted]; Trapper's Leather Pants (252501, -0.51 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.55 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.51 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Insignia Boots (4055, -0.14 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.14 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.1 attack_power points (0.61 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.19 DPS) [vendor]; Band of the Fist (17694, -0.23 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.1 attack_power points (0.56 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.14 DPS) [vendor]; Band of the Fist (17694, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (14.97 DPS) | yes | Scout's Blade (19545, -0.86 DPS) [pvp]; Darkspear Insurgent's Spellblade (272086, -1.02 DPS) [vendor]; Torturing Poker (7682, -1.18 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (14.88 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.38 DPS, sim-verified) [vendor]; Tork Wrench (11855, -14.78 DPS) [quest]; Satyr's Rod (15962, -14.83 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Double-barreled Shotgun (2098, -0.16 DPS, sim-verified) [world_drop]; Silver Star (3463, -0.18 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.23 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 80.0. Weights run: 1.6s. Verify run: 1.6s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.047 ± 0.006 per rating point (14 rating = 1%, 0.663 per %), hit=0.058 ± 0.006 per rating point (10 rating = 1%, 0.583 per %), melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.1 attack_power points (1.10 DPS) | yes | Hawkeye's Helm (14591, -0.44 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.44 DPS) [world_drop]; Nightscape Headband (8176, -0.49 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.99 DPS) | yes | Ghostshard Talisman (7731, -0.37 DPS, sim-verified) [dungeon]; Scout's Medallion (19536, -0.44 DPS) [rep]; Ethereal Talisman (4430, -0.54 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 attack_power points (1.15 DPS) | yes | Flintrock Shoulders (7755, -0.40 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.44 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.1 attack_power points (0.70 DPS) | yes | First Sergeant's Cloak (16340, -0.10 DPS) [pvp]; Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Wildhunter Cloak (16658, -0.20 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.2 attack_power points (1.35 DPS) | yes | Nightscape Tunic (8175, -0.60 DPS) [crafted]; Barbaric Harness (5739, -0.60 DPS) [crafted]; Wolffear Harness (13110, -0.70 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.39 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.50 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.7 attack_power points (1.02 DPS) | yes | Prowler's Leather Gloves (252524, -0.08 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.75 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.49 DPS) | yes | Defiler's Chain Girdle (20152, -0.37 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.59 DPS) [world_drop]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.24 DPS) [quest]; Brawler's Leather Legguards (252516, -0.49 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.1 attack_power points (0.90 DPS) | yes | Prowler's Leather Shoes (252465, -0.00 DPS) [crafted]; Imperial Leather Boots (6431, -0.10 DPS) [dungeon]; Excelsior Boots (4109, -0.10 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Legionnaire's Band (19512, -0.19 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest]; Ironspine's Eye (7686, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.50 DPS) | yes | Scout's Blade (19544, -2.91 DPS) [pvp]; Coldrage Dagger (10761, -2.93 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272085, -3.44 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 438.8 attack_power points (21.74 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS, sim-verified) [dungeon]; Stonecloth Branch (15963, -21.59 DPS) [world_drop]; Tork Wrench (11855, -21.64 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (80.0 DPS) | yes | Monolithic Bow (9426, -0.25 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Bow of Searing Arrows (2825, -0.87 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 120.7. Weights run: 1.6s. Verify run: 2.1s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.218 ± 0.061, crit=0.464 ± 0.025 per rating point (14 rating = 1%, 6.499 per %), hit=0.053 ± 0.009 per rating point (10 rating = 1%, 0.529 per %), melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.17 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -0.84 DPS) [crafted]; Undercity Reservist's Cap (20643, -0.90 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.8 attack_power points (1.13 DPS) | yes | Woven Ivy Necklace (19159, -0.21 DPS) [quest]; Scout's Medallion (19535, -0.34 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.48 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.4 attack_power points (1.37 DPS) | yes | Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted]; Failed Flying Experiment (9647, -0.17 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.19 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 23.1 attack_power points (1.25 DPS) | yes | Blisterbane Wrap (12552, -0.26 DPS) [dungeon]; Dark Phantom Cape (13122, -0.26 DPS) [world_drop]; Duskbat Drape (19982, -0.32 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 32.9 attack_power points (1.78 DPS) | yes | Mixologist's Tunic (12793, -0.08 DPS) [dungeon]; Quillward Harness (10583, -0.10 DPS) [dungeon]; Blazewind Breastplate (11193, -0.10 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 22.3 attack_power points (1.21 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.65 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 36.8 attack_power points (1.99 DPS) | yes | Darkmantle Grips (226828, -0.05 DPS) [vendor]; Gloves of Holy Might (867, -0.56 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.58 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.06 DPS) | yes | Skulker's Leather Waistguard (252474, -0.48 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.51 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.60 DPS, sim-verified) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 32.5 attack_power points (1.76 DPS) | yes | Basilisk Hide Pants (1718, -0.37 DPS) [world_drop]; Stone Guard's Leather Pants (220859, -0.40 DPS) [vendor]; Ferine Leggings (6690, -1.01 DPS, sim-verified) [dungeon] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.8 attack_power points (1.45 DPS) | yes | Sandstalker Ankleguards (12470, -0.01 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.30 DPS) | yes | Mark of Kern (2262, -0.22 DPS) [dungeon]; Assault Band (13095, -0.22 DPS) [world_drop]; Legionnaire's Band (19511, -1.80 DPS, sim-verified) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (120.7 DPS) | yes | Mark of Kern (2262, -0.03 DPS) [dungeon]; Assault Band (13095, -0.03 DPS) [world_drop]; Legionnaire's Band (19511, -1.55 DPS, sim-verified) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -3.34 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scout's Blade (19543, -0.41 DPS) [pvp]; Lifeforce Dirk (10750, -0.42 DPS) [quest]; Shadowblade (2163, -2.61 DPS, sim-verified) [world_drop] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (27.69 DPS) | yes | Thermotastic Egg Timer (9644, -27.50 DPS) [quest]; Stonecloth Branch (15963, -27.53 DPS) [world_drop]; Tork Wrench (11855, -27.58 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.00 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.00 DPS) [world_drop]; Dark Iron Rifle (16004, -1.71 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 180.6. Weights run: 1.6s. Verify run: 2.0s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.181 ± 0.023, crit=0.580 ± 0.036 per rating point (14 rating = 1%, 8.118 per %), hit=not significant (0.046 ± 0.013) per rating point (10 rating = 1%, 0.458 per %), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Darkmantle Cap (226829) | Saving the Best for Last [quest] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Embrace of the Lycan (9479, -2.39 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 32.4 attack_power points (1.72 DPS) | yes | Medallion of the Dawn (22659, -0.02 DPS) [quest]; Imperial Jewel (11933, -0.02 DPS) [dungeon]; Will of the Martyr (17044, -0.13 DPS) [quest] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 51.3 attack_power points (2.72 DPS) | yes | Darkspear Pauldrons (272105, -0.60 DPS) [vendor]; Defiler's Lizardhide Shoulders (20175, -0.80 DPS, sim-verified) [rep]; Dark Warder's Pauldrons (22241, -1.07 DPS) [dungeon] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 39.9 attack_power points (2.12 DPS) | yes | Cape of the Black Baron (13340, -0.12 DPS) [dungeon]; Howler's Furs (272414, -0.61 DPS) [vendor]; Windshear Cape (20691, -0.75 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -0.80 DPS) [crafted]; Nightbrace Tunic (12603, -1.15 DPS) [dungeon]; Tunic of Undead Slaying (23089, -5.99 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.13 DPS) [rep]; General's Leather Armsplints (16559, -0.18 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.10 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | 44.0 attack_power points (2.34 DPS) | yes | Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; Blood Guard's Leather Vices (16499, -0.01 DPS) [pvp]; Skul's Fingerbone Claws (13395, -0.21 DPS) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 46.1 attack_power points (2.45 DPS) | yes | General's Leather Girdle (16557, -0.13 DPS) [pvp]; Defiler's Leather Girdle (20190, -0.21 DPS) [rep]; Cadaverous Belt (14636, -0.32 DPS) [dungeon] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 54.1 attack_power points (2.88 DPS) | yes | Warbear Woolies (15065, -0.26 DPS) [crafted]; Sentinel's Leather Pants (237818, -0.32 DPS) [vendor]; Cadaverous Leggings (14638, -0.81 DPS, sim-verified) [dungeon] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (2.13 DPS) | yes | Drudge Boots (21532, -0.37 DPS) [quest]; Dunestalker's Boots (20715, -0.46 DPS) [quest]; Defiler's Leather Boots (20186, -0.52 DPS) [rep] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | White Bone Band (11862, -0.05 DPS) [quest]; Signet Ring of the Bronze Dragonflight (21201, -0.17 DPS) [quest]; Naglering (11669, -2.36 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | White Bone Band (11862, -0.03 DPS) [quest]; Signet Ring of the Bronze Dragonflight (21201, -0.15 DPS) [quest]; Naglering (11669, -5.80 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, -3.85 DPS, sim-verified) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Distracting Dagger (18392, -3.59 DPS) [dungeon]; The Lobotomizer (19324, -5.44 DPS, sim-verified) [rep]; Tome of Knowledge (13385, -33.00 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.27 DPS) [world_drop]; Dark Iron Rifle (16004, -1.59 DPS, sim-verified) [crafted] |

**New at 60:** head: Darkmantle Cap; neck: Amulet of the Darkmoon; shoulder: Defiler's Leather Shoulders; back: Deathguard's Cloak; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Cadaverous Gloves; waist: Ferocity of the Timbermaw; legs: Devilsaur Leggings; feet: Pads of the Dread Wolf; finger1: Legionnaire's Band; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Royal Seal of Eldre'Thalas; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Riphook

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

