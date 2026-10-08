# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 33.7. Weights run: 2.3s. Verify run: 3.9s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.037 ± 0.002, crit=0.080 ± 0.004 per rating point (14 rating = 1%, 1.114 per %), hit=0.768 ± 0.018 per rating point (10 rating = 1%, 7.680 per %), melee_haste=4.663 ± 0.256

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.47 DPS) | yes | Defender's Leather Hood (252447, -0.02 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.2 attack_power points (0.35 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.29 DPS) | yes | Slime-encrusted Pads (6461, -0.31 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.35 DPS) | yes | Catacomb Cloak (279899, -0.01 DPS) [quest]; Cape of the Brotherhood (5193, -0.06 DPS) [dungeon]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 attack_power points (0.69 DPS) | yes | Tunic of Westfall (2041, -0.05 DPS) [quest]; Defender's Leather Armor (252434, -0.12 DPS) [crafted]; Prospector's Chestpiece (14562, -0.23 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.1 attack_power points (0.34 DPS) | yes | Forest Leather Bracers (3202, -0.05 DPS) [world_drop]; Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (33.7 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.08 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.01 DPS) | yes | Brawler's Leather Belt (252428, -0.55 DPS) [crafted]; Ruffian Belt (5975, -0.67 DPS) [world]; Deviate Scale Belt (6468, -3.92 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (33.7 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.15 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (33.7 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.04 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.1 attack_power points (0.46 DPS) | yes | Signet of the Zhevra (285330, -0.11 DPS) [world]; Demon Band (12054, -0.23 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 7.2 attack_power points (0.41 DPS) | yes | Signet of the Zhevra (285330, -0.06 DPS) [world]; Demon Band (12054, -0.18 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (12.83 DPS) | yes | Evocator's Blade (2567, -0.80 DPS) [dungeon]; Buzzer Blade (2169, -1.83 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.56 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.6 attack_power points (12.60 DPS) | yes | Evocator's Blade (2567, -6.37 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.23 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted]; Light Bow (4576, -0.12 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 53.0. Weights run: 2.6s. Verify run: 4.2s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.033 ± 0.003, crit=0.071 ± 0.004 per rating point (14 rating = 1%, 0.998 per %), hit=1.004 ± 0.030 per rating point (10 rating = 1%, 10.043 per %), melee_haste=7.951 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.70 DPS) | yes | Brawler's Leather Helm (252512, -0.10 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.18 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.82 DPS) | yes | Kaleidoscope Chain (13084, -0.34 DPS) [world_drop]; Sentinel's Medallion (19541, -0.36 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.4 attack_power points (0.96 DPS) | yes | Barbaric Shoulders (5964, -0.36 DPS) [crafted]; Mantle of Thieves (2264, -0.37 DPS, sim-verified) [dungeon]; Bristlebark Amice (14573, -0.42 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.2 attack_power points (0.60 DPS) | yes | Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.12 DPS) [pvp] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (53.0 DPS) | yes | Dusky Leather Armor (7374, +0.00 DPS) [crafted]; Brawler's Leather Tunic (252508, +0.00 DPS) [crafted]; Raptorbane Armor (3566, -3.28 DPS, sim-verified) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.2 attack_power points (0.60 DPS) | yes | Cultist's Armguards (270032, -0.01 DPS) [quest]; Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.12 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (53.0 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Heavy Earthen Gloves (7359, -3.21 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (53.0 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.16 DPS) [crafted]; Highlander's Chain Girdle (20090, -3.46 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.52 DPS) | yes | Brawler's Leather Legguards (252516, -0.60 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.62 DPS) [crafted]; Trapper's Leather Pants (252501, -0.62 DPS) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (53.0 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Disjointed Shoes (277226, -3.47 DPS, sim-verified) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.3 attack_power points (0.78 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Insurgent's Band (272067, -0.25 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.30 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.2 attack_power points (0.71 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Insurgent's Band (272067, -0.19 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (18.81 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.28 DPS) [vendor]; Torturing Poker (7682, -1.49 DPS) [dungeon]; Thornspike (6681, -1.97 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (18.69 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.12 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -18.63 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.22 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.28 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Blackened Defias Armor; wrist: Hawkeye's Bracers; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 96.9. Weights run: 2.8s. Verify run: 2.4s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.037 ± 0.002, crit=0.075 ± 0.004 per rating point (14 rating = 1%, 1.056 per %), hit=0.993 ± 0.031 per rating point (10 rating = 1%, 9.931 per %), melee_haste=5.215 ± 0.372

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.4 attack_power points (1.58 DPS) | yes | Hawkeye's Helm (14591, -0.57 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.63 DPS) [world_drop]; Nightscape Headband (8176, -0.70 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.41 DPS) | yes | Ghostshard Talisman (7731, -0.48 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19540, -0.60 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.4 attack_power points (1.65 DPS) | yes | Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.84 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.4 attack_power points (1.01 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.29 DPS) [world_drop]; Wolfmaster Cape (6314, -0.31 DPS) [dungeon] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.7 attack_power points (1.95 DPS) | yes | Raptorbane Armor (3566, -0.82 DPS) [quest]; Nightscape Tunic (8175, -0.85 DPS) [crafted]; Wolffear Harness (13110, -0.88 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Hawkeye's Bracers (14590, -0.62 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.70 DPS) [quest]; Dusky Bracers (7378, -0.82 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 21.1 attack_power points (1.48 DPS) | yes | Prowler's Leather Gloves (252524, -0.12 DPS) [crafted]; Imperial Leather Gloves (4063, -0.19 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.72 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.11 DPS) | yes | Highlander's Chain Girdle (20090, -0.48 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.82 DPS) [world_drop]; Blackened Defias Belt (10403, -0.84 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.83 DPS) | yes | Basilisk Hide Pants (1718, -0.30 DPS) [world_drop]; Triprunner Dungarees (9624, -0.30 DPS) [quest]; Brawler's Leather Legguards (252516, -0.68 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.4 attack_power points (1.29 DPS) | yes | Prowler's Leather Shoes (252465, -0.01 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.15 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Protector's Band (19515, -0.26 DPS) [rep]; Field Researcher's Loop (281634, -0.40 DPS) [quest]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.41 DPS) | yes | Protector's Band (19515, -0.26 DPS) [rep]; Field Researcher's Loop (281634, -0.40 DPS) [quest]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (96.9 DPS) | yes | Black Menace (6831, -4.15 DPS) [quest]; Darkspear Insurgent's Spellblade (272085, -4.88 DPS) [vendor]; Coldrage Dagger (10761, -8.26 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 439.0 attack_power points (30.86 DPS) | yes | Black Menace (6831, -1.22 DPS, sim-verified) [quest]; Stonecloth Branch (15963, -30.65 DPS) [world_drop]; Satyr's Rod (15962, -30.79 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (96.9 DPS) | yes | Monolithic Bow (9426, -0.34 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.35 DPS) [vendor]; Bow of Searing Arrows (2825, -1.13 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 142.2. Weights run: 3.4s. Verify run: 3.3s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.358 ± 0.006, crit=0.753 ± 0.011 per rating point (14 rating = 1%, 10.545 per %), hit=1.564 ± 0.057 per rating point (10 rating = 1%, 15.635 per %), melee_haste=5.321 ± 0.809

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 46.5 attack_power points (3.90 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.37 DPS) [vendor]; Embrace of the Lycan (9479, -0.55 DPS) [dungeon]; White Bandit Mask (10008, -1.73 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 22.7 attack_power points (1.90 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.22 DPS) [quest]; Sentinel's Medallion (19539, -0.53 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 34.2 attack_power points (2.86 DPS) | yes | Sunburn Spaulders (274751, -0.76 DPS, sim-verified) [vendor]; Skulker's Leather Shoulder (252535, -0.86 DPS) [crafted]; Failed Flying Experiment (9647, -0.89 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.0 attack_power points (2.10 DPS) | yes | Blisterbane Wrap (12552, -0.39 DPS) [dungeon]; Dark Phantom Cape (13122, -0.39 DPS) [world_drop]; Duskbat Drape (19982, -0.50 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 44.2 attack_power points (3.70 DPS) | yes | Blazewind Breastplate (11193, -0.83 DPS) [quest]; Fungus Shroud Armor (17742, -0.86 DPS) [dungeon]; Warbear Harness (15064, -0.89 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 24.4 attack_power points (2.04 DPS) | yes | Skulker's Leather Bracers (252540, -0.43 DPS) [crafted]; Pridelord Bands (14672, -0.49 DPS) [world_drop]; Branded Leather Bracers (19508, -0.52 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 39.9 attack_power points (3.34 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.78 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.99 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.18 DPS) | yes | Highlander's Leather Girdle (20115, -0.62 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.65 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.65 DPS, sim-verified) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 44.2 attack_power points (3.70 DPS) | yes | Serpentskin Leggings (8262, -0.79 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -0.90 DPS, sim-verified) [quest]; Basilisk Hide Pants (1718, -1.31 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 29.1 attack_power points (2.44 DPS) | yes | Skulker's Leather Boots (252469, -0.04 DPS) [crafted]; Prowler's Leather Boots (252468, -0.10 DPS) [crafted]; Albino Crocscale Boots (17728, -0.16 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.6 attack_power points (2.99 DPS) | yes | Mark of Kern (2262, -1.31 DPS) [dungeon]; Assault Band (13095, -1.31 DPS) [world_drop]; Masons Fraternity Ring (9533, -1.39 DPS) [quest] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.2 attack_power points (1.86 DPS) | yes | Mark of Kern (2262, -0.19 DPS) [dungeon]; Assault Band (13095, -0.19 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.27 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (142.2 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (142.2 DPS) | yes | Mark of the Chosen (17774, -0.50 DPS, sim-verified) [quest] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (142.2 DPS) | yes | Lifeforce Dirk (10750, -0.66 DPS) [quest]; Charstone Dirk (17710, -0.66 DPS) [dungeon]; Shadowblade (2163, -4.67 DPS, sim-verified) [world_drop] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (42.87 DPS) | yes | Thermotastic Egg Timer (9644, -42.53 DPS) [quest]; Stonecloth Branch (15963, -42.62 DPS) [world_drop]; Windchaser Orb (15965, -42.76 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (142.2 DPS) | yes | Stinging Bow (10624, -0.17 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -2.00 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 263.2. Weights run: 3.3s. Verify run: 9.1s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.875 ± 0.018, crit=1.812 ± 0.032 per rating point (14 rating = 1%, 25.366 per %), hit=2.906 ± 0.132 per rating point (10 rating = 1%, 29.060 per %), melee_haste=10.925 ± 1.906

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 83.5 attack_power points (8.35 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.61 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 57.2 attack_power points (5.72 DPS) | yes | Beads of Ogre Might (22150, -0.41 DPS) [quest]; Mark of Fordring (15411, -0.58 DPS) [quest]; Medallion of the Dawn (22659, -0.78 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (263.2 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -13.66 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 57.1 attack_power points (5.71 DPS) | yes | Cape of the Black Baron (13340, -0.89 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.37 DPS) [rep]; Windshear Cape (20691, -2.09 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (263.2 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -2.06 DPS) [pvp]; Darkmantle Tunic (226825, -2.21 DPS) [quest]; Tunic of Undead Slaying (23089, -10.74 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (263.2 DPS) | yes | Marshal's Leather Armsplints (16460, -0.10 DPS) [pvp]; Bracers of the Eclipse (18375, -0.39 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -5.06 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (263.2 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.11 DPS) [crafted]; Raider Gloves (272099, -15.88 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 97.5 attack_power points (9.75 DPS) | yes | Belt of Preserved Heads (20216, -1.59 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -3.81 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.96 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (263.2 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -14.00 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 65.3 attack_power points (6.54 DPS) | yes | Fine Dawn Treaders (227815, -1.71 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -1.73 DPS) [dungeon]; Darkmantle Boots (22003, -2.03 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (263.2 DPS) | yes | Tarnished Elven Ring (18500, -1.32 DPS) [dungeon]; Cutthroat's Signet (272408, -1.51 DPS) [vendor]; Naglering (11669, -6.02 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (263.2 DPS) | yes | Tarnished Elven Ring (18500, -0.56 DPS) [dungeon]; Cutthroat's Signet (272408, -0.75 DPS) [vendor]; Naglering (11669, -5.65 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (263.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (263.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Heart of Wyrmthalak (22321, -2.69 DPS, sim-verified) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (263.2 DPS) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (263.2 DPS) | yes | The Lobotomizer (19324, -4.70 DPS, sim-verified) [rep]; Distracting Dagger (18392, -6.75 DPS) [dungeon]; Scepter of Interminable Focus (22329, -58.40 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (263.2 DPS) | yes | Blackcrow (12651, -0.29 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.84 DPS) [world_drop]; Dark Iron Rifle (16004, -2.84 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 621.3. Weights run: 3.5s. Verify run: 9.6s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.309 ± 0.024, crit=2.291 ± 0.040 per rating point (14 rating = 1%, 32.072 per %), hit=4.258 ± 0.245 per rating point (10 rating = 1%, 42.584 per %), melee_haste=17.990 ± 3.114

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 117.2 attack_power points (21.46 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -1.21 DPS) [pvp]; Outlaw's Collar (279253, -3.78 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 77.2 attack_power points (14.14 DPS) | yes | Beads of Ogre Might (22150, -1.95 DPS) [quest]; Mark of Fordring (15411, -3.51 DPS) [quest]; Medallion of the Dawn (22659, -3.87 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (621.3 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -30.54 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 70.6 attack_power points (12.92 DPS) | yes | Cape of the Black Baron (13340, -2.92 DPS) [dungeon]; Cloak of the Honor Guard (20073, -4.58 DPS) [rep]; Windshear Cape (20691, -4.97 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (621.3 DPS) | yes | Darkmantle Tunic (226825, -4.51 DPS) [quest]; Field Marshal's Leather Chestpiece (231543, -4.65 DPS) [pvp]; Tunic of Undead Slaying (23089, -24.41 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (621.3 DPS) | yes | Marshal's Leather Armsplints (16460, -0.20 DPS) [pvp]; Blackmist Armguards (12966, -1.44 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -10.81 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 74.7 attack_power points (13.67 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -1.95 DPS) [quest]; Raider Gloves (272099, -41.06 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 133.7 attack_power points (24.47 DPS) | yes | Belt of Preserved Heads (20216, -3.94 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -12.05 DPS) [vendor]; Highlander's Leather Girdle (20045, -12.37 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (621.3 DPS) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -33.44 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 85.2 attack_power points (15.60 DPS) | yes | Fine Dawn Treaders (227815, -3.19 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -4.39 DPS) [dungeon]; Darkmantle Boots (22003, -5.46 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (621.3 DPS) | yes | Tarnished Elven Ring (18500, -2.46 DPS) [dungeon]; Cutthroat's Signet (272408, -2.88 DPS) [vendor]; Naglering (11669, -12.92 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (621.3 DPS) | yes | Tarnished Elven Ring (18500, -1.27 DPS) [dungeon]; Cutthroat's Signet (272408, -1.69 DPS) [vendor]; Naglering (11669, -12.60 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (621.3 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (621.3 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -4.73 DPS) [crafted]; Counterattack Lodestone (18537, -7.71 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (621.3 DPS) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (621.3 DPS) | yes | The Lobotomizer (19324, -9.21 DPS, sim-verified) [rep]; Distracting Dagger (18392, -12.35 DPS) [dungeon]; Scepter of Interminable Focus (22329, -103.17 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (621.3 DPS) | yes | Blackcrow (12651, -0.78 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -3.15 DPS) [world_drop]; Dark Iron Rifle (16004, -6.40 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 33.5. Weights run: 2.3s. Verify run: 3.9s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.037 ± 0.002, crit=0.080 ± 0.004 per rating point (14 rating = 1%, 1.114 per %), hit=0.768 ± 0.018 per rating point (10 rating = 1%, 7.680 per %), melee_haste=4.663 ± 0.256

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.47 DPS) | yes | Defender's Leather Hood (252447, -0.02 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.2 attack_power points (0.35 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.29 DPS) | yes | Slime-encrusted Pads (6461, -0.30 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.35 DPS) | yes | Catacomb Cloak (279899, -0.01 DPS) [quest]; Cape of the Brotherhood (5193, -0.06 DPS) [dungeon]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 attack_power points (0.69 DPS) | yes | Defender's Leather Armor (252434, -0.12 DPS) [crafted]; Prospector's Chestpiece (14562, -0.23 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.23 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.2 attack_power points (0.29 DPS) | yes | Bristlebark Bindings (14569, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.06 DPS) [vendor]; Ratchet Wristwraps (274742, -0.12 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (33.5 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.12 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.01 DPS) | yes | Brawler's Leather Belt (252428, -0.55 DPS) [crafted]; Ruffian Belt (5975, -0.67 DPS) [world]; Deviate Scale Belt (6468, -3.94 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (33.5 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.18 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (33.5 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.06 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.1 attack_power points (0.46 DPS) | yes | Signet of the Zhevra (285330, -0.11 DPS) [world]; Demon Band (12054, -0.23 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 7.2 attack_power points (0.41 DPS) | yes | Signet of the Zhevra (285330, -0.06 DPS) [world]; Demon Band (12054, -0.18 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.23 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (12.83 DPS) | yes | Edward's Knife (251485, -0.33 DPS) [quest]; Scout's Blade (20441, -0.58 DPS) [pvp]; Evocator's Blade (2567, -0.80 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.6 attack_power points (12.60 DPS) | yes | Edward's Knife (251485, +0.00 DPS, sim-verified) [quest]; Tork Wrench (11855, -12.49 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.23 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted]; Light Bow (4576, -0.12 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 52.5. Weights run: 2.6s. Verify run: 4.2s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.033 ± 0.003, crit=0.071 ± 0.004 per rating point (14 rating = 1%, 0.998 per %), hit=1.004 ± 0.030 per rating point (10 rating = 1%, 10.043 per %), melee_haste=7.951 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.70 DPS) | yes | Brawler's Leather Helm (252512, -0.10 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.18 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.82 DPS) | yes | Kaleidoscope Chain (13084, -0.34 DPS) [world_drop]; Scout's Medallion (19537, -0.37 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.4 attack_power points (0.96 DPS) | yes | Barbaric Shoulders (5964, -0.36 DPS) [crafted]; Mantle of Thieves (2264, -0.37 DPS, sim-verified) [dungeon]; Bristlebark Amice (14573, -0.42 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.2 attack_power points (0.60 DPS) | yes | Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Wildhunter Cloak (16658, -0.01 DPS) [quest]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (52.5 DPS) | yes | Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Brawler's Leather Tunic (252508, +0.00 DPS) [crafted]; Dusky Leather Armor (7374, -3.00 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.2 attack_power points (0.60 DPS) | yes | Cultist's Armguards (270032, -0.01 DPS) [quest]; Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.12 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (52.5 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Heavy Earthen Gloves (7359, -2.82 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (52.5 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Deftkin Belt (16659, -0.11 DPS) [quest]; Defiler's Chain Girdle (20152, -3.07 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.52 DPS) | yes | Brawler's Leather Legguards (252516, -0.61 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.62 DPS) [crafted]; Trapper's Leather Pants (252501, -0.62 DPS) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (52.5 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.13 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.3 attack_power points (0.78 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Insurgent's Band (272067, -0.25 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.30 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.2 attack_power points (0.71 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Insurgent's Band (272067, -0.19 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (18.81 DPS) | yes | Scout's Blade (19545, -1.07 DPS) [pvp]; Darkspear Insurgent's Spellblade (272086, -1.28 DPS) [vendor]; Torturing Poker (7682, -1.49 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (18.69 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -4.80 DPS, sim-verified) [vendor]; Tork Wrench (11855, -18.58 DPS) [quest]; Satyr's Rod (15962, -18.63 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.22 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.28 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Blackened Defias Armor; wrist: Hawkeye's Bracers; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 95.9. Weights run: 2.8s. Verify run: 2.5s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.037 ± 0.002, crit=0.075 ± 0.004 per rating point (14 rating = 1%, 1.056 per %), hit=0.993 ± 0.031 per rating point (10 rating = 1%, 9.931 per %), melee_haste=5.215 ± 0.372

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.4 attack_power points (1.58 DPS) | yes | Hawkeye's Helm (14591, -0.56 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.63 DPS) [world_drop]; Nightscape Headband (8176, -0.70 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.41 DPS) | yes | Ghostshard Talisman (7731, -0.48 DPS, sim-verified) [dungeon]; Scout's Medallion (19536, -0.60 DPS) [rep]; Ethereal Talisman (4430, -0.76 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.4 attack_power points (1.65 DPS) | yes | Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Nightscape Shoulders (8192, -0.84 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.4 attack_power points (1.01 DPS) | yes | First Sergeant's Cloak (16340, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.29 DPS) [world_drop]; Wildhunter Cloak (16658, -0.31 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.7 attack_power points (1.95 DPS) | yes | Nightscape Tunic (8175, -0.85 DPS) [crafted]; Wolffear Harness (13110, -0.86 DPS, sim-verified) [world_drop]; Barbaric Harness (5739, -0.87 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.59 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.70 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 21.1 attack_power points (1.48 DPS) | yes | Prowler's Leather Gloves (252524, -0.12 DPS) [crafted]; Imperial Leather Gloves (4063, -0.19 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.69 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.11 DPS) | yes | Defiler's Chain Girdle (20152, -0.48 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.82 DPS) [world_drop]; Blackened Defias Belt (10403, -0.84 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.83 DPS) | yes | Basilisk Hide Pants (1718, -0.30 DPS) [world_drop]; Triprunner Dungarees (9624, -0.30 DPS) [quest]; Brawler's Leather Legguards (252516, -0.68 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.4 attack_power points (1.29 DPS) | yes | Prowler's Leather Shoes (252465, -0.01 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.15 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Legionnaire's Band (19512, -0.26 DPS) [rep]; Field Researcher's Loop (281634, -0.40 DPS) [quest]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.41 DPS) | yes | Legionnaire's Band (19512, -0.26 DPS) [rep]; Field Researcher's Loop (281634, -0.40 DPS) [quest]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (95.9 DPS) | yes | Scout's Blade (19544, -4.11 DPS) [pvp]; Darkspear Insurgent's Spellblade (272085, -4.88 DPS) [vendor]; Coldrage Dagger (10761, -7.98 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 439.0 attack_power points (30.86 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Stonecloth Branch (15963, -30.65 DPS) [world_drop]; Tork Wrench (11855, -30.72 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (95.9 DPS) | yes | Monolithic Bow (9426, -0.34 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.35 DPS) [vendor]; Bow of Searing Arrows (2825, -1.12 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 145.6. Weights run: 3.4s. Verify run: 3.0s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.358 ± 0.006, crit=0.753 ± 0.011 per rating point (14 rating = 1%, 10.545 per %), hit=1.564 ± 0.057 per rating point (10 rating = 1%, 15.635 per %), melee_haste=5.321 ± 0.809

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 46.5 attack_power points (3.90 DPS) | yes | Blood Guard's Leather Headband (220851, -0.37 DPS) [vendor]; Embrace of the Lycan (9479, -0.83 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.73 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 22.7 attack_power points (1.90 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.22 DPS) [quest]; Woven Ivy Necklace (19159, -0.37 DPS) [quest]; Scout's Medallion (19535, -0.53 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.9 attack_power points (2.26 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.25 DPS) [crafted]; Failed Flying Experiment (9647, -0.28 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.0 attack_power points (2.10 DPS) | yes | Blisterbane Wrap (12552, -0.39 DPS) [dungeon]; Dark Phantom Cape (13122, -0.39 DPS) [world_drop]; Duskbat Drape (19982, -0.50 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 44.2 attack_power points (3.70 DPS) | yes | Blazewind Breastplate (11193, -0.83 DPS) [quest]; Fungus Shroud Armor (17742, -0.86 DPS) [dungeon]; Warbear Harness (15064, -1.03 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 24.4 attack_power points (2.04 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.37 DPS) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 39.9 attack_power points (3.34 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.78 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.99 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.18 DPS) | yes | Skulker's Leather Waistguard (252474, -0.61 DPS, sim-verified) [crafted]; Defiler's Leather Girdle (20193, -0.62 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.65 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 44.2 attack_power points (3.70 DPS) | yes | Serpentskin Leggings (8262, -1.13 DPS, sim-verified) [world_drop]; Basilisk Hide Pants (1718, -1.31 DPS) [world_drop]; Triprunner Dungarees (9624, -1.40 DPS) [quest] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 29.1 attack_power points (2.44 DPS) | yes | Skulker's Leather Boots (252469, -0.04 DPS) [crafted]; Prowler's Leather Boots (252468, -0.10 DPS) [crafted]; Albino Crocscale Boots (17728, -0.16 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.6 attack_power points (2.99 DPS) | yes | Legionnaire's Band (19511, -1.12 DPS) [rep]; Mark of Kern (2262, -1.31 DPS) [dungeon]; Assault Band (13095, -1.31 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.01 DPS) | yes | Legionnaire's Band (19511, -0.15 DPS) [rep]; Mark of Kern (2262, -0.34 DPS) [dungeon]; Assault Band (13095, -0.34 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (145.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (145.6 DPS) | yes | Molten Heart of the Mountain (249470, -0.74 DPS, sim-verified) [crafted] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (145.6 DPS) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Scout's Blade (19543, -0.51 DPS) [pvp]; Searing Needle (12531, -5.18 DPS, sim-verified) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (42.87 DPS) | yes | Thermotastic Egg Timer (9644, -42.53 DPS) [quest]; Stonecloth Branch (15963, -42.62 DPS) [world_drop]; Tork Wrench (11855, -42.70 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (145.6 DPS) | yes | Stinging Bow (10624, -0.17 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -1.94 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 262.7. Weights run: 3.3s. Verify run: 9.1s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.875 ± 0.018, crit=1.812 ± 0.032 per rating point (14 rating = 1%, 25.366 per %), hit=2.906 ± 0.132 per rating point (10 rating = 1%, 29.060 per %), melee_haste=10.925 ± 1.906

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 83.5 attack_power points (8.35 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 57.2 attack_power points (5.72 DPS) | yes | Beads of Ogre Might (22150, -0.41 DPS) [quest]; Mark of Fordring (15411, -0.58 DPS) [quest]; Medallion of the Dawn (22659, -0.78 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (262.7 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -12.72 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 57.1 attack_power points (5.71 DPS) | yes | Cape of the Black Baron (13340, -0.89 DPS) [dungeon]; Deathguard's Cloak (20068, -1.37 DPS) [rep]; Windshear Cape (20691, -2.09 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (262.7 DPS) | yes | Warlord's Leather Breastplate (231549, -2.06 DPS) [pvp]; Darkmantle Tunic (226825, -2.21 DPS) [quest]; Tunic of Undead Slaying (23089, -10.07 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (262.7 DPS) | yes | General's Leather Armsplints (16559, -0.10 DPS) [pvp]; Bracers of the Eclipse (18375, -0.39 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.86 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (262.7 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.11 DPS) [crafted]; Raider Gloves (272099, -14.71 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 97.5 attack_power points (9.75 DPS) | yes | Belt of Preserved Heads (20216, -1.00 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -3.81 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.96 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (262.7 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -12.51 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 65.3 attack_power points (6.54 DPS) | yes | Fine Dawn Treaders (227815, -1.04 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -1.73 DPS) [dungeon]; Darkmantle Boots (22003, -2.03 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (262.7 DPS) | yes | Tarnished Elven Ring (18500, -1.32 DPS) [dungeon]; Cutthroat's Signet (272408, -1.51 DPS) [vendor]; Naglering (11669, -5.37 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (262.7 DPS) | yes | Tarnished Elven Ring (18500, -0.56 DPS) [dungeon]; Cutthroat's Signet (272408, -0.75 DPS) [vendor]; Naglering (11669, -5.01 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (262.7 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (262.7 DPS) | yes | Blackhand's Breadth (13965, -1.16 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, -1.75 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -3.62 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (262.7 DPS) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (262.7 DPS) | yes | The Lobotomizer (19324, -4.77 DPS, sim-verified) [rep]; Distracting Dagger (18392, -6.75 DPS) [dungeon]; Scepter of Interminable Focus (22329, -58.40 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (262.7 DPS) | yes | Blackcrow (12651, -0.29 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.84 DPS) [world_drop]; Dark Iron Rifle (16004, -2.22 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 624.2. Weights run: 3.5s. Verify run: 9.8s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.309 ± 0.024, crit=2.291 ± 0.040 per rating point (14 rating = 1%, 32.072 per %), hit=4.258 ± 0.245 per rating point (10 rating = 1%, 42.584 per %), melee_haste=17.990 ± 3.114

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 117.2 attack_power points (21.46 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Champion's Leather Helm (227057, -1.21 DPS) [pvp]; Outlaw's Collar (279253, -3.78 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 77.2 attack_power points (14.14 DPS) | yes | Beads of Ogre Might (22150, -1.95 DPS) [quest]; Mark of Fordring (15411, -3.51 DPS) [quest]; Medallion of the Dawn (22659, -3.87 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -32.39 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 70.6 attack_power points (12.92 DPS) | yes | Cape of the Black Baron (13340, -2.92 DPS) [dungeon]; Deathguard's Cloak (20068, -4.58 DPS) [rep]; Windshear Cape (20691, -4.97 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmantle Tunic (226825, -4.51 DPS) [quest]; Warlord's Leather Breastplate (231549, -4.65 DPS) [pvp]; Tunic of Undead Slaying (23089, -26.42 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Armsplints (16559, -0.20 DPS) [pvp]; Blackmist Armguards (12966, -1.44 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -11.35 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 74.7 attack_power points (13.67 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -1.95 DPS) [quest]; Raider Gloves (272099, -38.51 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 133.7 attack_power points (24.47 DPS) | yes | Belt of Preserved Heads (20216, -4.87 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -12.05 DPS) [vendor]; Defiler's Leather Girdle (20190, -12.37 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -30.07 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 85.2 attack_power points (15.60 DPS) | yes | Fine Dawn Treaders (227815, -3.16 DPS) [vendor]; Shadowcraft Boots (16711, -4.39 DPS) [dungeon]; Darkmantle Boots (22003, -5.46 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.46 DPS) [dungeon]; Cutthroat's Signet (272408, -2.88 DPS) [vendor]; Naglering (11669, -14.50 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.27 DPS) [dungeon]; Cutthroat's Signet (272408, -1.69 DPS) [vendor]; Naglering (11669, -13.75 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+12.4 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (624.2 DPS) | yes | Distracting Dagger (18392, -12.35 DPS) [dungeon]; The Lobotomizer (19324, -12.48 DPS, sim-verified) [rep]; Scepter of Interminable Focus (22329, -103.17 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.78 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -3.15 DPS) [world_drop]; Dark Iron Rifle (16004, -7.02 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

