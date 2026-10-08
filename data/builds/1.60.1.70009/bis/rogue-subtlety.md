# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 31.2. Weights run: 2.5s. Verify run: 1.6s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.035 ± 0.002, crit=0.078 ± 0.004 per rating point (14 rating = 1%, 1.086 per %), hit=0.778 ± 0.019 per rating point (10 rating = 1%, 7.783 per %), melee_haste=4.795 ± 0.256

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.46 DPS) | yes | Defender's Leather Hood (252447, -0.02 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.2 attack_power points (0.35 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.29 DPS) | yes | Slime-encrusted Pads (6461, -0.31 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.35 DPS) | yes | Catacomb Cloak (279899, -0.01 DPS) [quest]; Cape of the Brotherhood (5193, -0.06 DPS) [dungeon]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.2 attack_power points (0.69 DPS) | yes | Tunic of Westfall (2041, -0.05 DPS) [quest]; Defender's Leather Armor (252434, -0.12 DPS) [crafted]; Prospector's Chestpiece (14562, -0.23 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.1 attack_power points (0.34 DPS) | yes | Forest Leather Bracers (3202, -0.05 DPS) [world_drop]; Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.2 attack_power points (0.57 DPS) | yes | Brawler's Leather Gloves (252494, -0.12 DPS) [crafted]; Bristlebark Gloves (14572, -0.12 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.18 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.01 DPS) | yes | Brawler's Leather Belt (252428, -0.55 DPS) [crafted]; Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.67 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.3 attack_power points (0.86 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.06 DPS) [dungeon]; Defender's Leather Pants (252445, -0.12 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 attack_power points (0.63 DPS) | yes | Brawler's Leather Boots (252439, -0.06 DPS) [crafted]; Blackened Defias Boots (10402, -0.28 DPS) [dungeon]; Footpads of the Fang (10411, -0.28 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.1 attack_power points (0.46 DPS) | yes | Signet of the Zhevra (285330, -0.11 DPS) [world]; Demon Band (12054, -0.23 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 7.3 attack_power points (0.41 DPS) | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Demon Band (12054, -0.18 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (12.82 DPS) | yes | Evocator's Blade (2567, -0.80 DPS) [dungeon]; Buzzer Blade (2169, -1.83 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.55 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.6 attack_power points (12.59 DPS) | yes | Evocator's Blade (2567, -5.75 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.23 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted]; Light Bow (4576, -0.12 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 50.8. Weights run: 2.8s. Verify run: 1.8s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.035 ± 0.003, crit=0.073 ± 0.004 per rating point (14 rating = 1%, 1.025 per %), hit=0.995 ± 0.029 per rating point (10 rating = 1%, 9.952 per %), melee_haste=7.294 ± 0.338

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.70 DPS) | yes | Brawler's Leather Helm (252512, -0.10 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.18 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.82 DPS) | yes | Kaleidoscope Chain (13084, -0.34 DPS) [world_drop]; Sentinel's Medallion (19541, -0.36 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.4 attack_power points (0.96 DPS) | yes | Barbaric Shoulders (5964, -0.36 DPS) [crafted]; Mantle of Thieves (2264, -0.37 DPS, sim-verified) [dungeon]; Bristlebark Amice (14573, -0.42 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.2 attack_power points (0.60 DPS) | yes | Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.12 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.93 DPS) | yes | Dusky Leather Armor (7374, -0.09 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.10 DPS) [crafted]; Brawler's Leather Armor (252490, -0.22 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.2 attack_power points (0.60 DPS) | yes | Cultist's Armguards (270032, -0.01 DPS) [quest]; Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.12 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.93 DPS) | yes | Insignia Gloves (6408, -0.16 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.22 DPS) [crafted]; Wolfclaw Gloves (1978, -0.28 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.40 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.35 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.51 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.52 DPS) | yes | Brawler's Leather Legguards (252516, -0.60 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.62 DPS) [crafted]; Trapper's Leather Pants (252501, -0.62 DPS) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.70 DPS) | yes | Feet of the Lynx (1121, -0.04 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Insignia Boots (4055, -0.22 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.3 attack_power points (0.78 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Insurgent's Band (272067, -0.25 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.30 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.2 attack_power points (0.71 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Insurgent's Band (272067, -0.19 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (18.81 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.28 DPS) [vendor]; Torturing Poker (7682, -1.49 DPS) [dungeon]; Thornspike (6681, -1.97 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (18.69 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -4.02 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -18.63 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.22 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.28 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 98.3. Weights run: 3.0s. Verify run: 1.9s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.038 ± 0.003, crit=0.079 ± 0.004 per rating point (14 rating = 1%, 1.105 per %), hit=0.969 ± 0.037 per rating point (10 rating = 1%, 9.691 per %), melee_haste=6.921 ± 0.499

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.4 attack_power points (1.58 DPS) | yes | Hawkeye's Helm (14591, -0.57 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.63 DPS) [world_drop]; Nightscape Headband (8176, -0.70 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.41 DPS) | yes | Ghostshard Talisman (7731, -0.48 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19540, -0.60 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.4 attack_power points (1.65 DPS) | yes | Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Nightscape Shoulders (8192, -0.84 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.4 attack_power points (1.01 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.29 DPS) [world_drop]; Wolfmaster Cape (6314, -0.31 DPS) [dungeon] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.7 attack_power points (1.95 DPS) | yes | Raptorbane Armor (3566, -0.82 DPS) [quest]; Nightscape Tunic (8175, -0.86 DPS) [crafted]; Wolffear Harness (13110, -0.87 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Hawkeye's Bracers (14590, -0.64 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.70 DPS) [quest]; Dusky Bracers (7378, -0.82 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 21.1 attack_power points (1.49 DPS) | yes | Prowler's Leather Gloves (252524, -0.12 DPS) [crafted]; Imperial Leather Gloves (4063, -0.19 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.72 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.11 DPS) | yes | Highlander's Chain Girdle (20090, -0.48 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.82 DPS) [world_drop]; Blackened Defias Belt (10403, -0.84 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.83 DPS) | yes | Basilisk Hide Pants (1718, -0.30 DPS) [world_drop]; Triprunner Dungarees (9624, -0.30 DPS) [quest]; Brawler's Leather Legguards (252516, -0.68 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.4 attack_power points (1.30 DPS) | yes | Prowler's Leather Shoes (252465, -0.01 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.15 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Protector's Band (19515, -0.26 DPS) [rep]; Field Researcher's Loop (281634, -0.40 DPS) [quest]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.41 DPS) | yes | Protector's Band (19515, -0.26 DPS) [rep]; Field Researcher's Loop (281634, -0.40 DPS) [quest]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (98.3 DPS) | yes | Black Menace (6831, -4.16 DPS) [quest]; Darkspear Insurgent's Spellblade (272085, -4.89 DPS) [vendor]; Coldrage Dagger (10761, -8.36 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 439.0 attack_power points (30.90 DPS) | yes | Black Menace (6831, -1.00 DPS, sim-verified) [quest]; Stonecloth Branch (15963, -30.69 DPS) [world_drop]; Satyr's Rod (15962, -30.83 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (98.3 DPS) | yes | Monolithic Bow (9426, -0.34 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.35 DPS) [vendor]; Bow of Searing Arrows (2825, -1.13 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 141.9. Weights run: 3.5s. Verify run: 2.3s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.367 ± 0.008, crit=0.763 ± 0.013 per rating point (14 rating = 1%, 10.678 per %), hit=1.666 ± 0.067 per rating point (10 rating = 1%, 16.656 per %), melee_haste=7.125 ± 0.900

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 46.7 attack_power points (3.78 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.27 DPS) [vendor]; Embrace of the Lycan (9479, -0.54 DPS) [dungeon]; White Bandit Mask (10008, -1.67 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 22.8 attack_power points (1.84 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.22 DPS) [quest]; Sentinel's Medallion (19539, -0.52 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 35.3 attack_power points (2.86 DPS) | yes | Skulker's Leather Shoulder (252535, -0.91 DPS) [crafted]; Failed Flying Experiment (9647, -0.94 DPS) [quest]; Sunburn Spaulders (274751, -1.04 DPS, sim-verified) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.1 attack_power points (2.03 DPS) | yes | Blisterbane Wrap (12552, -0.37 DPS) [dungeon]; Dark Phantom Cape (13122, -0.37 DPS) [world_drop]; Duskbat Drape (19982, -0.49 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 45.3 attack_power points (3.67 DPS) | yes | Blazewind Breastplate (11193, -0.88 DPS) [quest]; Fungus Shroud Armor (17742, -0.90 DPS) [dungeon]; Warbear Harness (15064, -1.13 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 24.5 attack_power points (1.98 DPS) | yes | Branded Leather Bracers (19508, -0.36 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.42 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 40.1 attack_power points (3.24 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.76 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.97 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.08 DPS) | yes | Highlander's Leather Girdle (20115, -0.59 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.61 DPS, sim-verified) [crafted]; Prowler's Leather Waistguard (252473, -0.61 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 45.3 attack_power points (3.67 DPS) | yes | Serpentskin Leggings (8262, -0.85 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.15 DPS, sim-verified) [quest]; Basilisk Hide Pants (1718, -1.35 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 29.2 attack_power points (2.37 DPS) | yes | Skulker's Leather Boots (252469, -0.04 DPS) [crafted]; Prowler's Leather Boots (252468, -0.10 DPS) [crafted]; Albino Crocscale Boots (17728, -0.15 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.7 attack_power points (2.97 DPS) | yes | Mark of Kern (2262, -1.35 DPS) [dungeon]; Assault Band (13095, -1.35 DPS) [world_drop]; Masons Fraternity Ring (9533, -1.42 DPS) [quest] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.3 attack_power points (1.80 DPS) | yes | Mark of Kern (2262, -0.19 DPS) [dungeon]; Assault Band (13095, -0.19 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.26 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (141.9 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (141.9 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (141.9 DPS) | yes | Lifeforce Dirk (10750, -0.63 DPS) [quest]; Charstone Dirk (17710, -0.63 DPS) [dungeon]; Shadowblade (2163, -5.10 DPS, sim-verified) [world_drop] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (41.40 DPS) | yes | Thermotastic Egg Timer (9644, -41.07 DPS) [quest]; Stonecloth Branch (15963, -41.16 DPS) [world_drop]; Windchaser Orb (15965, -41.29 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (141.9 DPS) | yes | Stinging Bow (10624, -0.17 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -1.92 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 248.6. Weights run: 3.4s. Verify run: 2.4s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.876 ± 0.019, crit=1.804 ± 0.033 per rating point (14 rating = 1%, 25.256 per %), hit=3.012 ± 0.159 per rating point (10 rating = 1%, 30.124 per %), melee_haste=11.265 ± 2.259

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 85.5 attack_power points (8.13 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.78 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 58.3 attack_power points (5.54 DPS) | yes | Beads of Ogre Might (22150, -0.39 DPS) [quest]; Mark of Fordring (15411, -0.67 DPS) [quest]; Medallion of the Dawn (22659, -0.86 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 84.2 attack_power points (8.01 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, -0.65 DPS) [pvp]; Darkspear Pauldrons (272105, -0.79 DPS) [vendor]; Field Marshal's Leather Epaulets (231547, -1.40 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 58.1 attack_power points (5.52 DPS) | yes | Cape of the Black Baron (13340, -0.95 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.40 DPS) [rep]; Windshear Cape (20691, -2.09 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, -1.96 DPS) [pvp]; Darkmantle Tunic (226825, -2.04 DPS) [quest]; Tunic of Undead Slaying (23089, -9.30 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Armsplints (16460, -0.10 DPS) [pvp]; Bracers of the Eclipse (18375, -0.37 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.49 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 61.7 attack_power points (5.86 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -0.80 DPS) [quest]; Stormshroud Gloves (21278, -3.03 DPS, sim-verified) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 99.7 attack_power points (9.47 DPS) | yes | Belt of Preserved Heads (20216, -0.93 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -3.84 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.97 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 101.2 attack_power points (9.62 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -1.12 DPS) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 66.1 attack_power points (6.28 DPS) | yes | Shadowcraft Boots (16711, -1.68 DPS) [dungeon]; Fine Dawn Treaders (227815, -1.97 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -2.00 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.25 DPS) [dungeon]; Cutthroat's Signet (272408, -1.42 DPS) [vendor]; Naglering (11669, -4.92 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.54 DPS) [dungeon]; Cutthroat's Signet (272408, -0.71 DPS) [vendor]; Naglering (11669, -4.50 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+10.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Royal Seal of Eldre'Thalas (18465, -1.10 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+9.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (248.6 DPS) | yes | The Lobotomizer (19324, -2.80 DPS, sim-verified) [rep]; Distracting Dagger (18392, -6.41 DPS) [dungeon]; Scepter of Interminable Focus (22329, -55.40 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.29 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.90 DPS) [world_drop]; Dark Iron Rifle (16004, -2.00 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 528.5. Weights run: 3.6s. Verify run: 2.6s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.941 ± 0.021, crit=1.941 ± 0.035 per rating point (14 rating = 1%, 27.173 per %), hit=3.173 ± 0.254 per rating point (10 rating = 1%, 31.734 per %), melee_haste=13.442 ± 3.231

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 90.6 attack_power points (15.93 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -1.60 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 60.8 attack_power points (10.69 DPS) | yes | Beads of Ogre Might (22150, -0.90 DPS) [quest]; Mark of Fordring (15411, -1.35 DPS) [quest]; Medallion of the Dawn (22659, -1.70 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 87.5 attack_power points (15.37 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, -1.15 DPS) [pvp]; Darkspear Pauldrons (272105, -2.31 DPS, sim-verified) [vendor]; Field Marshal's Leather Epaulets (231547, -2.63 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.7 attack_power points (10.50 DPS) | yes | Cape of the Black Baron (13340, -1.87 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.82 DPS) [rep]; Windshear Cape (20691, -3.98 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, -3.75 DPS) [pvp]; Darkmantle Tunic (226825, -4.03 DPS) [quest]; Tunic of Undead Slaying (23089, -20.72 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Armsplints (16460, -0.18 DPS) [pvp]; Bracers of the Eclipse (18375, -0.79 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.83 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 63.4 attack_power points (11.14 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.45 DPS) [crafted]; Stormshroud Gloves (21278, -6.08 DPS, sim-verified) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 104.2 attack_power points (18.32 DPS) | yes | Belt of Preserved Heads (20216, -3.61 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -7.57 DPS) [rep]; Ferocity of the Timbermaw (227805, -7.95 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 106.7 attack_power points (18.76 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -2.43 DPS) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 68.8 attack_power points (12.09 DPS) | yes | Shadowcraft Boots (16711, -3.25 DPS) [dungeon]; Darkmantle Boots (22003, -3.90 DPS) [quest]; Fine Dawn Treaders (227815, -5.02 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.47 DPS) [dungeon]; Cutthroat's Signet (272408, -2.81 DPS) [vendor]; Naglering (11669, -11.61 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.02 DPS) [dungeon]; Cutthroat's Signet (272408, -1.36 DPS) [vendor]; Naglering (11669, -10.60 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+12.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -1.60 DPS) [quest]; Frozen Heart of the Mountain (249469, -6.14 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+19.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (528.5 DPS) | yes | The Lobotomizer (19324, -6.13 DPS, sim-verified) [rep]; Distracting Dagger (18392, -11.86 DPS) [dungeon]; Scepter of Interminable Focus (22329, -101.83 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.56 DPS) [dungeon]; The Purifier (22656, -1.83 DPS) [quest]; Dark Iron Rifle (16004, -5.41 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Royal Seal of Eldre'Thalas; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.9. Weights run: 2.5s. Verify run: 1.7s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.035 ± 0.002, crit=0.078 ± 0.004 per rating point (14 rating = 1%, 1.086 per %), hit=0.778 ± 0.019 per rating point (10 rating = 1%, 7.783 per %), melee_haste=4.795 ± 0.256

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.46 DPS) | yes | Defender's Leather Hood (252447, -0.02 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.2 attack_power points (0.35 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.29 DPS) | yes | Slime-encrusted Pads (6461, -0.30 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.35 DPS) | yes | Catacomb Cloak (279899, -0.01 DPS) [quest]; Cape of the Brotherhood (5193, -0.06 DPS) [dungeon]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.2 attack_power points (0.69 DPS) | yes | Defender's Leather Armor (252434, -0.12 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.23 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.23 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.2 attack_power points (0.29 DPS) | yes | Bristlebark Bindings (14569, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.06 DPS) [vendor]; Ratchet Wristwraps (274742, -0.12 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.2 attack_power points (0.57 DPS) | yes | Brawler's Leather Gloves (252494, -0.12 DPS) [crafted]; Bristlebark Gloves (14572, -0.12 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.18 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.01 DPS) | yes | Brawler's Leather Belt (252428, -0.55 DPS) [crafted]; Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.67 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.3 attack_power points (0.86 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.06 DPS) [dungeon]; Defender's Leather Pants (252445, -0.12 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 attack_power points (0.63 DPS) | yes | Brawler's Leather Boots (252439, -0.06 DPS) [crafted]; Blackened Defias Boots (10402, -0.28 DPS) [dungeon]; Footpads of the Fang (10411, -0.28 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.1 attack_power points (0.46 DPS) | yes | Signet of the Zhevra (285330, -0.11 DPS) [world]; Demon Band (12054, -0.23 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 7.3 attack_power points (0.41 DPS) | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Demon Band (12054, -0.18 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.23 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (12.82 DPS) | yes | Edward's Knife (251485, -0.33 DPS) [quest]; Scout's Blade (20441, -0.58 DPS) [pvp]; Evocator's Blade (2567, -0.80 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.6 attack_power points (12.59 DPS) | yes | Edward's Knife (251485, +0.00 DPS) [quest]; Tork Wrench (11855, -12.48 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.23 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted]; Light Bow (4576, -0.12 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 50.5. Weights run: 2.8s. Verify run: 1.8s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.035 ± 0.003, crit=0.073 ± 0.004 per rating point (14 rating = 1%, 1.025 per %), hit=0.995 ± 0.029 per rating point (10 rating = 1%, 9.952 per %), melee_haste=7.294 ± 0.338

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.70 DPS) | yes | Brawler's Leather Helm (252512, -0.10 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.18 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.18 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.82 DPS) | yes | Kaleidoscope Chain (13084, -0.34 DPS) [world_drop]; Scout's Medallion (19537, -0.37 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.4 attack_power points (0.96 DPS) | yes | Barbaric Shoulders (5964, -0.36 DPS) [crafted]; Mantle of Thieves (2264, -0.37 DPS, sim-verified) [dungeon]; Bristlebark Amice (14573, -0.42 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.2 attack_power points (0.60 DPS) | yes | Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Wildhunter Cloak (16658, -0.01 DPS) [quest]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.5 attack_power points (0.85 DPS) | yes | Brawler's Leather Tunic (252508, -0.01 DPS) [crafted]; Brawler's Leather Armor (252490, -0.13 DPS) [crafted]; Defender's Leather Tunic (252450, -0.13 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.2 attack_power points (0.60 DPS) | yes | Cultist's Armguards (270032, -0.01 DPS) [quest]; Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.12 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.93 DPS) | yes | Insignia Gloves (6408, -0.16 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.22 DPS) [crafted]; Wolfclaw Gloves (1978, -0.28 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.40 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.35 DPS) [dungeon]; Deftkin Belt (16659, -0.46 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.52 DPS) | yes | Brawler's Leather Legguards (252516, -0.61 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.62 DPS) [crafted]; Trapper's Leather Pants (252501, -0.62 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 attack_power points (0.66 DPS) | yes | Brawler's Leather Boots (252439, -0.06 DPS) [crafted]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.3 attack_power points (0.78 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Insurgent's Band (272067, -0.25 DPS) [vendor]; Band of the Fist (17694, -0.30 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.2 attack_power points (0.71 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Insurgent's Band (272067, -0.19 DPS) [vendor]; Band of the Fist (17694, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (18.81 DPS) | yes | Scout's Blade (19545, -1.07 DPS) [pvp]; Darkspear Insurgent's Spellblade (272086, -1.28 DPS) [vendor]; Torturing Poker (7682, -1.49 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (18.69 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -3.90 DPS, sim-verified) [vendor]; Tork Wrench (11855, -18.58 DPS) [quest]; Satyr's Rod (15962, -18.63 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.22 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.28 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 97.9. Weights run: 3.0s. Verify run: 2.0s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.038 ± 0.003, crit=0.079 ± 0.004 per rating point (14 rating = 1%, 1.105 per %), hit=0.969 ± 0.037 per rating point (10 rating = 1%, 9.691 per %), melee_haste=6.921 ± 0.499

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.4 attack_power points (1.58 DPS) | yes | Hawkeye's Helm (14591, -0.56 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.63 DPS) [world_drop]; Nightscape Headband (8176, -0.70 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.41 DPS) | yes | Ghostshard Talisman (7731, -0.48 DPS, sim-verified) [dungeon]; Scout's Medallion (19536, -0.60 DPS) [rep]; Ethereal Talisman (4430, -0.76 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.4 attack_power points (1.65 DPS) | yes | Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Nightscape Shoulders (8192, -0.84 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.4 attack_power points (1.01 DPS) | yes | First Sergeant's Cloak (16340, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.29 DPS) [world_drop]; Wildhunter Cloak (16658, -0.31 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.7 attack_power points (1.95 DPS) | yes | Nightscape Tunic (8175, -0.86 DPS) [crafted]; Wolffear Harness (13110, -0.86 DPS, sim-verified) [world_drop]; Barbaric Harness (5739, -0.87 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.63 DPS, sim-verified) [world_drop]; Cultist's Armguards (270032, -0.70 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 21.1 attack_power points (1.49 DPS) | yes | Prowler's Leather Gloves (252524, -0.12 DPS) [crafted]; Imperial Leather Gloves (4063, -0.19 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.66 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.11 DPS) | yes | Defiler's Chain Girdle (20152, -0.48 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.82 DPS) [world_drop]; Blackened Defias Belt (10403, -0.84 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.83 DPS) | yes | Basilisk Hide Pants (1718, -0.30 DPS) [world_drop]; Triprunner Dungarees (9624, -0.30 DPS) [quest]; Brawler's Leather Legguards (252516, -0.68 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.4 attack_power points (1.30 DPS) | yes | Prowler's Leather Shoes (252465, -0.01 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.15 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Legionnaire's Band (19512, -0.26 DPS) [rep]; Field Researcher's Loop (281634, -0.40 DPS) [quest]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.41 DPS) | yes | Legionnaire's Band (19512, -0.26 DPS) [rep]; Field Researcher's Loop (281634, -0.40 DPS) [quest]; Ironspine's Eye (7686, -0.47 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (97.9 DPS) | yes | Scout's Blade (19544, -4.12 DPS) [pvp]; Darkspear Insurgent's Spellblade (272085, -4.89 DPS) [vendor]; Coldrage Dagger (10761, -8.04 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 439.0 attack_power points (30.90 DPS) | yes | Coldrage Dagger (10761, -0.76 DPS, sim-verified) [dungeon]; Stonecloth Branch (15963, -30.69 DPS) [world_drop]; Tork Wrench (11855, -30.76 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (97.9 DPS) | yes | Monolithic Bow (9426, -0.34 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.35 DPS) [vendor]; Bow of Searing Arrows (2825, -1.12 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 145.4. Weights run: 3.5s. Verify run: 2.3s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.367 ± 0.008, crit=0.763 ± 0.013 per rating point (14 rating = 1%, 10.678 per %), hit=1.666 ± 0.067 per rating point (10 rating = 1%, 16.656 per %), melee_haste=7.125 ± 0.900

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 46.7 attack_power points (3.78 DPS) | yes | Blood Guard's Leather Headband (220851, -0.27 DPS) [vendor]; Embrace of the Lycan (9479, -0.90 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.67 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 22.8 attack_power points (1.84 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.22 DPS) [quest]; Woven Ivy Necklace (19159, -0.36 DPS) [quest]; Scout's Medallion (19535, -0.52 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.0 attack_power points (2.19 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.24 DPS) [crafted]; Failed Flying Experiment (9647, -0.27 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.1 attack_power points (2.03 DPS) | yes | Blisterbane Wrap (12552, -0.37 DPS) [dungeon]; Dark Phantom Cape (13122, -0.37 DPS) [world_drop]; Duskbat Drape (19982, -0.49 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 45.3 attack_power points (3.67 DPS) | yes | Blazewind Breastplate (11193, -0.88 DPS) [quest]; Fungus Shroud Armor (17742, -0.90 DPS) [dungeon]; Warbear Harness (15064, -0.99 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 24.5 attack_power points (1.98 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.57 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 40.1 attack_power points (3.24 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.76 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.97 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.08 DPS) | yes | Skulker's Leather Waistguard (252474, -0.56 DPS, sim-verified) [crafted]; Defiler's Leather Girdle (20193, -0.59 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.61 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 45.3 attack_power points (3.67 DPS) | yes | Serpentskin Leggings (8262, -1.09 DPS, sim-verified) [world_drop]; Basilisk Hide Pants (1718, -1.35 DPS) [world_drop]; Triprunner Dungarees (9624, -1.43 DPS) [quest] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 29.2 attack_power points (2.37 DPS) | yes | Skulker's Leather Boots (252469, -0.04 DPS) [crafted]; Prowler's Leather Boots (252468, -0.10 DPS) [crafted]; Albino Crocscale Boots (17728, -0.15 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.7 attack_power points (2.97 DPS) | yes | Legionnaire's Band (19511, -1.16 DPS) [rep]; Mark of Kern (2262, -1.35 DPS) [dungeon]; Assault Band (13095, -1.35 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.94 DPS) | yes | Legionnaire's Band (19511, -0.14 DPS) [rep]; Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (145.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (145.4 DPS) | yes | Molten Heart of the Mountain (249470, -1.08 DPS, sim-verified) [crafted] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (145.4 DPS) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Scout's Blade (19543, -0.48 DPS) [pvp]; Searing Needle (12531, -5.36 DPS, sim-verified) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (41.40 DPS) | yes | Thermotastic Egg Timer (9644, -41.07 DPS) [quest]; Stonecloth Branch (15963, -41.16 DPS) [world_drop]; Tork Wrench (11855, -41.24 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (145.4 DPS) | yes | Stinging Bow (10624, -0.17 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -1.97 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 250.1. Weights run: 3.4s. Verify run: 2.5s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.876 ± 0.019, crit=1.804 ± 0.033 per rating point (14 rating = 1%, 25.256 per %), hit=3.012 ± 0.159 per rating point (10 rating = 1%, 30.124 per %), melee_haste=11.265 ± 2.259

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 85.5 attack_power points (8.13 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 58.3 attack_power points (5.54 DPS) | yes | Beads of Ogre Might (22150, -0.39 DPS) [quest]; Mark of Fordring (15411, -0.67 DPS) [quest]; Medallion of the Dawn (22659, -0.86 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 84.2 attack_power points (8.01 DPS) | yes | Champion's Leather Shoulders (227056, -0.65 DPS) [pvp]; Darkspear Pauldrons (272105, -0.79 DPS) [vendor]; Warlord's Leather Spaulders (231551, -1.40 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 58.1 attack_power points (5.52 DPS) | yes | Cape of the Black Baron (13340, -0.95 DPS) [dungeon]; Deathguard's Cloak (20068, -1.40 DPS) [rep]; Windshear Cape (20691, -2.09 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, -1.96 DPS) [pvp]; Darkmantle Tunic (226825, -2.04 DPS) [quest]; Tunic of Undead Slaying (23089, -9.58 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Armsplints (16559, -0.10 DPS) [pvp]; Bracers of the Eclipse (18375, -0.37 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.42 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 61.7 attack_power points (5.86 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -0.80 DPS) [quest]; Stormshroud Gloves (21278, -2.66 DPS, sim-verified) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 99.7 attack_power points (9.47 DPS) | yes | Belt of Preserved Heads (20216, -1.18 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -3.84 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.97 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 101.2 attack_power points (9.62 DPS) | yes | General's Leather Legguards (231554, +0.00 DPS) [pvp]; Legionnaire's Leather Legguards (227059, -1.12 DPS) [pvp]; Plaguehound Leggings (18736, -1.39 DPS, sim-verified) [dungeon] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 66.1 attack_power points (6.28 DPS) | yes | Shadowcraft Boots (16711, -1.68 DPS) [dungeon]; Fine Dawn Treaders (227815, -1.73 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -2.00 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.25 DPS) [dungeon]; Cutthroat's Signet (272408, -1.42 DPS) [vendor]; Naglering (11669, -5.11 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.54 DPS) [dungeon]; Cutthroat's Signet (272408, -0.71 DPS) [vendor]; Naglering (11669, -4.76 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+10.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, -0.27 DPS) [quest]; Blackhand's Breadth (13965, -1.20 DPS) [quest]; Hand of Justice (11815, -1.52 DPS, sim-verified) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (250.1 DPS) | yes | The Lobotomizer (19324, -4.21 DPS, sim-verified) [rep]; Distracting Dagger (18392, -6.41 DPS) [dungeon]; Scepter of Interminable Focus (22329, -55.40 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.29 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.90 DPS) [world_drop]; Dark Iron Rifle (16004, -2.20 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 527.2. Weights run: 3.6s. Verify run: 2.6s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.941 ± 0.021, crit=1.941 ± 0.035 per rating point (14 rating = 1%, 27.173 per %), hit=3.173 ± 0.254 per rating point (10 rating = 1%, 31.734 per %), melee_haste=13.442 ± 3.231

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 90.6 attack_power points (15.93 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -1.60 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 60.8 attack_power points (10.69 DPS) | yes | Beads of Ogre Might (22150, -0.90 DPS) [quest]; Mark of Fordring (15411, -1.35 DPS) [quest]; Medallion of the Dawn (22659, -1.70 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 87.5 attack_power points (15.37 DPS) | yes | Champion's Leather Shoulders (227056, -1.15 DPS) [pvp]; Darkspear Pauldrons (272105, -1.39 DPS) [vendor]; Warlord's Leather Spaulders (231551, -2.63 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.7 attack_power points (10.50 DPS) | yes | Cape of the Black Baron (13340, -2.39 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -2.82 DPS) [rep]; Windshear Cape (20691, -3.98 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, -3.75 DPS) [pvp]; Darkmantle Tunic (226825, -4.03 DPS) [quest]; Tunic of Undead Slaying (23089, -23.06 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Armsplints (16559, -0.18 DPS) [pvp]; Bracers of the Eclipse (18375, -0.79 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -9.07 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 63.4 attack_power points (11.14 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.45 DPS) [crafted]; Stormshroud Gloves (21278, -5.95 DPS, sim-verified) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 104.2 attack_power points (18.32 DPS) | yes | Belt of Preserved Heads (20216, -5.11 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -7.57 DPS) [rep]; Ferocity of the Timbermaw (227805, -7.95 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 106.7 attack_power points (18.76 DPS) | yes | General's Leather Legguards (231554, +0.00 DPS) [pvp]; Legionnaire's Leather Legguards (227059, -2.43 DPS) [pvp]; Plaguehound Leggings (18736, -5.35 DPS, sim-verified) [dungeon] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 68.8 attack_power points (12.09 DPS) | yes | Shadowcraft Boots (16711, -3.25 DPS) [dungeon]; Darkmantle Boots (22003, -3.90 DPS) [quest]; Fine Dawn Treaders (227815, -4.20 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.47 DPS) [dungeon]; Cutthroat's Signet (272408, -2.81 DPS) [vendor]; Naglering (11669, -13.41 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.02 DPS) [dungeon]; Cutthroat's Signet (272408, -1.36 DPS) [vendor]; Naglering (11669, -12.17 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+12.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18465, -0.13 DPS) [quest]; Blackhand's Breadth (13965, -1.73 DPS) [quest]; Frozen Heart of the Mountain (249469, -6.27 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+20.3 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (527.2 DPS) | yes | The Lobotomizer (19324, -5.83 DPS, sim-verified) [rep]; Distracting Dagger (18392, -11.86 DPS) [dungeon]; Scepter of Interminable Focus (22329, -101.83 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.56 DPS) [dungeon]; The Purifier (22656, -1.83 DPS) [quest]; Dark Iron Rifle (16004, -6.91 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

