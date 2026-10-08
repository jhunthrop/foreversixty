# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.3. Weights run: 2.5s. Verify run: 1.6s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.133 ± 0.006, crit=0.284 ± 0.011 per rating point (14 rating = 1%, 3.972 per %), hit=0.817 ± 0.020 per rating point (10 rating = 1%, 8.174 per %), melee_haste=5.325 ± 0.302

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 9.1 attack_power points (0.50 DPS) | yes | Defender's Leather Hood (252447, -0.06 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.8 attack_power points (0.37 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.7 attack_power points (0.31 DPS) | yes | Slime-encrusted Pads (6461, -0.33 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.8 attack_power points (0.37 DPS) | yes | Catacomb Cloak (279899, -0.04 DPS) [quest]; Cape of the Brotherhood (5193, -0.06 DPS) [dungeon]; Dark Leather Cloak (2316, -0.08 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.9 attack_power points (0.71 DPS) | yes | Tunic of Westfall (2041, -0.03 DPS) [quest]; Defender's Leather Armor (252434, -0.14 DPS) [crafted]; Prospector's Chestpiece (14562, -0.23 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.5 attack_power points (0.36 DPS) | yes | Forest Leather Bracers (3202, -0.05 DPS) [world_drop]; Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.8 attack_power points (0.59 DPS) | yes | Brawler's Leather Gloves (252494, -0.12 DPS) [crafted]; Bristlebark Gloves (14572, -0.13 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.21 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.99 DPS) | yes | Brawler's Leather Belt (252428, -0.52 DPS) [crafted]; Deviate Scale Belt (6468, -0.57 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.66 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 16.2 attack_power points (0.89 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.15 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 12.1 attack_power points (0.66 DPS) | yes | Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Blackened Defias Boots (10402, -0.29 DPS) [dungeon]; Footpads of the Fang (10411, -0.29 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.5 attack_power points (0.47 DPS) | yes | Signet of the Zhevra (285330, -0.10 DPS) [world]; Demon Band (12054, -0.25 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 7.8 attack_power points (0.43 DPS) | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Demon Band (12054, -0.21 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (12.56 DPS) | yes | Evocator's Blade (2567, -0.78 DPS) [dungeon]; Buzzer Blade (2169, -1.79 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.48 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 225.0 attack_power points (12.36 DPS) | yes | Evocator's Blade (2567, -6.36 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.5 attack_power points (0.25 DPS) | yes | Fine Longbow (11304, -0.03 DPS) [vendor]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted]; Light Bow (4576, -0.12 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 60.7. Weights run: 2.6s. Verify run: 1.8s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.136 ± 0.007, crit=0.292 ± 0.012 per rating point (14 rating = 1%, 4.082 per %), hit=1.071 ± 0.029 per rating point (10 rating = 1%, 10.712 per %), melee_haste=7.844 ± 0.353

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.71 DPS) | yes | Brawler's Leather Helm (252512, -0.04 DPS) [crafted]; Tribal Worg Helm (6204, -0.17 DPS) [world]; Brawler's Leather Hood (252504, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.83 DPS) | yes | Kaleidoscope Chain (13084, -0.32 DPS) [world_drop]; Sentinel's Medallion (19541, -0.35 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 17.5 attack_power points (1.04 DPS) | yes | Mantle of Thieves (2264, -0.40 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.40 DPS) [crafted]; Bristlebark Amice (14573, -0.46 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.9 attack_power points (0.65 DPS) | yes | Wolfmaster Cape (6314, -0.06 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.14 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.95 DPS) | yes | Dusky Leather Armor (7374, -0.01 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.05 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.8 attack_power points (0.64 DPS) | yes | Cultist's Armguards (270032, -0.05 DPS) [quest]; Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.13 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.95 DPS) | yes | Insignia Gloves (6408, -0.13 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.19 DPS) [crafted]; Wolfclaw Gloves (1978, -0.25 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.43 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.36 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.46 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.54 DPS) | yes | Brawler's Leather Pants (252500, -0.58 DPS) [crafted]; Trapper's Leather Pants (252501, -0.58 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.60 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 12.1 attack_power points (0.72 DPS) | yes | Disjointed Shoes (277226, -0.01 DPS) [quest]; Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Insignia Boots (4055, -0.18 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 14.2 attack_power points (0.84 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Insurgent's Band (272067, -0.31 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.32 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.8 attack_power points (0.76 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Insurgent's Band (272067, -0.23 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (19.13 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.31 DPS) [vendor]; Torturing Poker (7682, -1.51 DPS) [dungeon]; Thornspike (6681, -2.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (19.01 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.90 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -18.94 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS) [world_drop]; Silver Star (3463, -0.20 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 113.1. Weights run: 2.8s. Verify run: 1.8s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.376 ± 0.011, crit=0.792 ± 0.019 per rating point (14 rating = 1%, 11.088 per %), hit=1.157 ± 0.040 per rating point (10 rating = 1%, 11.566 per %), melee_haste=8.764 ± 0.453

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.1 attack_power points (1.71 DPS) | yes | Warden's Wizard Hat (14604, -0.54 DPS) [world_drop]; Hawkeye's Helm (14591, -0.55 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.63 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.31 DPS) | yes | Sentinel's Medallion (19540, -0.32 DPS) [rep]; Ghostshard Talisman (7731, -0.39 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.1 attack_power points (1.78 DPS) | yes | Forest Tracker Epaulets (2278, -0.55 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.55 DPS) [dungeon]; Nightscape Shoulders (8192, -0.79 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.8 attack_power points (1.16 DPS) | yes | Sergeant Major's Cape (16336, -0.23 DPS) [pvp]; Hawkeye's Cloak (14593, -0.34 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.43 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 34.1 attack_power points (2.24 DPS) | yes | Wolffear Harness (13110, -0.85 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.89 DPS) [crafted]; Dusky Leather Armor (7374, -0.98 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.31 DPS) | yes | Hawkeye's Bracers (14590, -0.55 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.59 DPS) [dungeon]; Dusky Bracers (7378, -0.59 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 31.1 attack_power points (2.04 DPS) | yes | Prowler's Leather Gloves (252524, -0.57 DPS) [crafted]; Imperial Leather Gloves (4063, -0.64 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.82 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.97 DPS) | yes | Highlander's Chain Girdle (20090, -0.47 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.57 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.76 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 28.9 attack_power points (1.90 DPS) | yes | Triprunner Dungarees (9624, -0.07 DPS) [quest]; Ferine Leggings (6690, -0.19 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.62 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.1 attack_power points (1.45 DPS) | yes | Prowler's Leather Shoes (252465, -0.10 DPS) [crafted]; Imperial Leather Boots (6431, -0.13 DPS) [dungeon]; Excelsior Boots (4109, -0.18 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.31 DPS) | yes | Protector's Band (19515, -0.07 DPS) [rep]; Field Researcher's Loop (281634, -0.22 DPS) [quest]; Ironspine's Eye (7686, -0.24 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.31 DPS) | yes | Protector's Band (19515, -0.07 DPS) [rep]; Field Researcher's Loop (281634, -0.22 DPS) [quest]; Ironspine's Eye (7686, -0.24 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (113.1 DPS) | yes | Black Menace (6831, -3.88 DPS) [quest]; Darkspear Insurgent's Spellblade (272085, -4.55 DPS) [vendor]; Coldrage Dagger (10761, -4.92 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 442.0 attack_power points (28.99 DPS) | yes | Black Menace (6831, -0.86 DPS, sim-verified) [quest]; Stonecloth Branch (15963, -28.80 DPS) [world_drop]; Satyr's Rod (15962, -28.90 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (113.1 DPS) | yes | Monolithic Bow (9426, -0.25 DPS) [dungeon]; Swiftwind (13038, -0.29 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.09 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 166.4. Weights run: 3.2s. Verify run: 2.2s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.434 ± 0.012, crit=0.912 ± 0.020 per rating point (14 rating = 1%, 12.770 per %), hit=1.539 ± 0.059 per rating point (10 rating = 1%, 15.391 per %), melee_haste=9.419 ± 0.850

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 48.8 attack_power points (4.04 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.38 DPS) [vendor]; Embrace of the Lycan (9479, -0.73 DPS) [dungeon]; White Bandit Mask (10008, -1.82 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.6 attack_power points (1.96 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.30 DPS) [quest]; Sentinel's Medallion (19539, -0.53 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 36.2 attack_power points (2.99 DPS) | yes | Skulker's Leather Shoulder (252535, -0.94 DPS) [crafted]; Failed Flying Experiment (9647, -0.98 DPS) [quest]; Sunburn Spaulders (274751, -1.47 DPS, sim-verified) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.1 attack_power points (2.16 DPS) | yes | Blisterbane Wrap (12552, -0.38 DPS) [dungeon]; Dark Phantom Cape (13122, -0.38 DPS) [world_drop]; Duskbat Drape (19982, -0.50 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 46.2 attack_power points (3.82 DPS) | yes | Blazewind Breastplate (11193, -0.84 DPS) [quest]; Fungus Shroud Armor (17742, -0.85 DPS) [dungeon]; Warbear Harness (15064, -1.52 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.5 attack_power points (2.11 DPS) | yes | Skulker's Leather Bracers (252540, -0.46 DPS) [crafted]; Pridelord Bands (14672, -0.51 DPS) [world_drop]; Branded Leather Bracers (19508, -0.60 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 41.6 attack_power points (3.44 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.73 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.03 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.15 DPS) | yes | Highlander's Leather Girdle (20115, -0.43 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.49 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.56 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 46.2 attack_power points (3.82 DPS) | yes | Serpentskin Leggings (8262, -0.85 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.33 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.55 DPS, sim-verified) [quest] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 30.4 attack_power points (2.52 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Prowler's Leather Boots (252468, -0.13 DPS) [crafted]; Albino Crocscale Boots (17728, -0.14 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.4 attack_power points (2.93 DPS) | yes | Masons Fraternity Ring (9533, -1.27 DPS) [quest]; Mark of Kern (2262, -1.27 DPS) [dungeon]; Assault Band (13095, -1.27 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.9 attack_power points (1.90 DPS) | yes | Masons Fraternity Ring (9533, -0.23 DPS) [quest]; Mark of Kern (2262, -0.24 DPS) [dungeon]; Assault Band (13095, -0.24 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (166.4 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (166.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (166.4 DPS) | yes | Barman Shanker (12791, -2.41 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -3.47 DPS) [quest]; Charstone Dirk (17710, -3.47 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (42.37 DPS) | yes | Barman Shanker (12791, -18.91 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -42.01 DPS) [quest]; Stonecloth Branch (15963, -42.12 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (166.4 DPS) | yes | Stinging Bow (10624, -0.26 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.26 DPS) [world_drop]; Dark Iron Rifle (16004, -1.93 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 257.3. Weights run: 2.7s. Verify run: 2.2s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.079 ± 0.025, crit=2.249 ± 0.043 per rating point (14 rating = 1%, 31.482 per %), hit=3.228 ± 0.166 per rating point (10 rating = 1%, 32.284 per %), melee_haste=15.611 ± 2.128

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 96.1 attack_power points (7.83 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 63.5 attack_power points (5.17 DPS) | yes | Mark of Fordring (15411, -0.49 DPS) [quest]; Beads of Ogre Might (22150, -0.59 DPS) [quest]; Medallion of the Dawn (22659, -0.65 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 88.6 attack_power points (7.22 DPS) | yes | Darkspear Pauldrons (272105, -0.08 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.23 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, -1.03 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 attack_power points (4.91 DPS) | yes | Cape of the Black Baron (13340, -0.74 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.30 DPS) [rep]; Windshear Cape (20691, -1.72 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, -1.86 DPS) [pvp]; Darkmantle Tunic (226825, -2.27 DPS) [quest]; Tunic of Undead Slaying (23089, -9.69 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Armsplints (16460, -0.08 DPS) [pvp]; Bracers of the Eclipse (18375, -0.47 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.11 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 67.1 attack_power points (5.47 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.62 DPS) [crafted]; Stormshroud Gloves (21278, -1.83 DPS, sim-verified) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 108.2 attack_power points (8.82 DPS) | yes | Belt of Preserved Heads (20216, -1.47 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -3.48 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.82 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 119.1 attack_power points (9.71 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -1.74 DPS) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 72.5 attack_power points (5.91 DPS) | yes | Shadowcraft Boots (16711, -1.56 DPS) [dungeon]; Darkmantle Boots (22003, -1.84 DPS) [quest]; Fine Dawn Treaders (227815, -1.96 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.33 DPS) [dungeon]; Cutthroat's Signet (272408, -1.50 DPS) [vendor]; Naglering (11669, -5.40 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.51 DPS) [dungeon]; Cutthroat's Signet (272408, -0.68 DPS) [vendor]; Naglering (11669, -4.88 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -5.57 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.76 DPS) [crafted]; Counterattack Lodestone (18537, -3.34 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+9.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (257.3 DPS) | yes | The Lobotomizer (19324, -4.06 DPS, sim-verified) [rep]; Distracting Dagger (18392, -5.50 DPS) [dungeon]; Scepter of Interminable Focus (22329, -46.83 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.26 DPS) [dungeon]; The Purifier (22656, -0.57 DPS) [quest]; Dark Iron Rifle (16004, -2.41 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 623.1. Weights run: 2.9s. Verify run: 2.2s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.132 ± 0.026, crit=2.354 ± 0.044 per rating point (14 rating = 1%, 32.962 per %), hit=3.962 ± 0.225 per rating point (10 rating = 1%, 39.618 per %), melee_haste=12.634 ± 2.984

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 112.2 attack_power points (19.10 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -0.62 DPS) [pvp]; Outlaw's Collar (279253, -3.06 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 71.6 attack_power points (12.19 DPS) | yes | Beads of Ogre Might (22150, -1.36 DPS) [quest]; Mark of Fordring (15411, -2.15 DPS) [quest]; Medallion of the Dawn (22659, -2.49 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 103.2 attack_power points (17.57 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, -1.47 DPS) [pvp]; Darkspear Pauldrons (272105, -2.34 DPS, sim-verified) [vendor]; Field Marshal's Leather Epaulets (231547, -3.21 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 67.6 attack_power points (11.51 DPS) | yes | Cape of the Black Baron (13340, -2.66 DPS) [dungeon]; Cloak of the Honor Guard (20073, -3.91 DPS) [rep]; Windshear Cape (20691, -4.70 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (623.1 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -3.99 DPS) [pvp]; Darkmantle Tunic (226825, -4.42 DPS) [quest]; Tunic of Undead Slaying (23089, -23.12 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (623.1 DPS) | yes | Marshal's Leather Armsplints (16460, -0.17 DPS) [pvp]; Bracers of the Eclipse (18375, -1.05 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -9.42 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 72.6 attack_power points (12.35 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; Devilsaur Gauntlets (15063, -1.98 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 124.0 attack_power points (21.10 DPS) | yes | Belt of Preserved Heads (20216, -4.88 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -9.71 DPS) [rep]; Ferocity of the Timbermaw (227805, -10.51 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 123.5 attack_power points (21.02 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -2.88 DPS) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 78.9 attack_power points (13.43 DPS) | yes | Fine Dawn Treaders (227815, -2.54 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -3.79 DPS) [dungeon]; Darkmantle Boots (22003, -4.72 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (623.1 DPS) | yes | Tarnished Elven Ring (18500, -2.89 DPS) [dungeon]; Cutthroat's Signet (272408, -3.25 DPS) [vendor]; Naglering (11669, -13.29 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (623.1 DPS) | yes | Tarnished Elven Ring (18500, -1.09 DPS) [dungeon]; Cutthroat's Signet (272408, -1.45 DPS) [vendor]; Naglering (11669, -12.18 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (623.1 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -4.92 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (623.1 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Hand of Justice (11815, -2.09 DPS, sim-verified) [dungeon]; Frozen Heart of the Mountain (249469, -5.15 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (623.1 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Felstriker (12590, -21.69 DPS, sim-verified) [dungeon] |
| off_hand | The Lobotomizer (19324) | Stormpike Guard [rep] | 661.1 attack_power points (112.51 DPS) | yes | Felstriker (12590, +0.00 DPS, sim-verified) [dungeon]; Distracting Dagger (18392, -15.37 DPS) [dungeon]; Scepter of Interminable Focus (22329, -100.16 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (623.1 DPS) | yes | Blackcrow (12651, -0.67 DPS) [dungeon]; The Purifier (22656, -2.22 DPS) [quest]; Dark Iron Rifle (16004, -6.82 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.0. Weights run: 2.5s. Verify run: 1.7s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.133 ± 0.006, crit=0.284 ± 0.011 per rating point (14 rating = 1%, 3.972 per %), hit=0.817 ± 0.020 per rating point (10 rating = 1%, 8.174 per %), melee_haste=5.325 ± 0.302

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 9.1 attack_power points (0.50 DPS) | yes | Defender's Leather Hood (252447, -0.06 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.8 attack_power points (0.37 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.7 attack_power points (0.31 DPS) | yes | Slime-encrusted Pads (6461, -0.32 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.8 attack_power points (0.37 DPS) | yes | Catacomb Cloak (279899, -0.04 DPS) [quest]; Cape of the Brotherhood (5193, -0.06 DPS) [dungeon]; Dark Leather Cloak (2316, -0.08 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.9 attack_power points (0.71 DPS) | yes | Defender's Leather Armor (252434, -0.14 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.23 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.25 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.7 attack_power points (0.31 DPS) | yes | Bristlebark Bindings (14569, -0.01 DPS) [world_drop]; Wolf Bracers (4794, -0.06 DPS) [vendor]; Ratchet Wristwraps (274742, -0.12 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.8 attack_power points (0.59 DPS) | yes | Brawler's Leather Gloves (252494, -0.12 DPS) [crafted]; Bristlebark Gloves (14572, -0.12 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.21 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.99 DPS) | yes | Brawler's Leather Belt (252428, -0.52 DPS) [crafted]; Deviate Scale Belt (6468, -0.57 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.66 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 16.2 attack_power points (0.89 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.15 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 12.1 attack_power points (0.66 DPS) | yes | Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Blackened Defias Boots (10402, -0.29 DPS) [dungeon]; Footpads of the Fang (10411, -0.29 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.5 attack_power points (0.47 DPS) | yes | Signet of the Zhevra (285330, -0.10 DPS) [world]; Demon Band (12054, -0.25 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 7.8 attack_power points (0.43 DPS) | yes | Signet of the Zhevra (285330, +0.00 DPS, sim-verified) [world]; Demon Band (12054, -0.21 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (12.56 DPS) | yes | Edward's Knife (251485, -0.30 DPS) [quest]; Scout's Blade (20441, -0.55 DPS) [pvp]; Evocator's Blade (2567, -0.78 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 225.0 attack_power points (12.36 DPS) | yes | Edward's Knife (251485, -0.16 DPS, sim-verified) [quest]; Tork Wrench (11855, -12.25 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.5 attack_power points (0.25 DPS) | yes | Fine Longbow (11304, -0.03 DPS) [vendor]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted]; Light Bow (4576, -0.12 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 60.4. Weights run: 2.6s. Verify run: 1.8s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.136 ± 0.007, crit=0.292 ± 0.012 per rating point (14 rating = 1%, 4.082 per %), hit=1.071 ± 0.029 per rating point (10 rating = 1%, 10.712 per %), melee_haste=7.844 ± 0.353

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.71 DPS) | yes | Brawler's Leather Helm (252512, -0.04 DPS) [crafted]; Tribal Worg Helm (6204, -0.17 DPS) [world]; Brawler's Leather Hood (252504, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.83 DPS) | yes | Kaleidoscope Chain (13084, -0.32 DPS) [world_drop]; Scout's Medallion (19537, -0.35 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 17.5 attack_power points (1.04 DPS) | yes | Mantle of Thieves (2264, -0.40 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.40 DPS) [crafted]; Bristlebark Amice (14573, -0.46 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.9 attack_power points (0.65 DPS) | yes | Wolfmaster Cape (6314, -0.06 DPS) [dungeon]; Wildhunter Cloak (16658, -0.06 DPS) [quest]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 15.9 attack_power points (0.94 DPS) | yes | Brawler's Leather Tunic (252508, -0.05 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted]; Defender's Leather Tunic (252450, -0.18 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.8 attack_power points (0.64 DPS) | yes | Cultist's Armguards (270032, -0.05 DPS) [quest]; Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.13 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.95 DPS) | yes | Insignia Gloves (6408, -0.13 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.19 DPS) [crafted]; Wolfclaw Gloves (1978, -0.25 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.43 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.36 DPS) [dungeon]; Deftkin Belt (16659, -0.44 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.54 DPS) | yes | Brawler's Leather Pants (252500, -0.58 DPS) [crafted]; Trapper's Leather Pants (252501, -0.58 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.60 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 12.1 attack_power points (0.72 DPS) | yes | Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 14.2 attack_power points (0.84 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Insurgent's Band (272067, -0.31 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.32 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.8 attack_power points (0.76 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Insurgent's Band (272067, -0.23 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (19.13 DPS) | yes | Scout's Blade (19545, -1.04 DPS) [pvp]; Darkspear Insurgent's Spellblade (272086, -1.31 DPS) [vendor]; Torturing Poker (7682, -1.51 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (19.01 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.89 DPS, sim-verified) [vendor]; Tork Wrench (11855, -18.89 DPS) [quest]; Satyr's Rod (15962, -18.94 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.53 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS) [world_drop]; Silver Star (3463, -0.20 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 111.9. Weights run: 2.8s. Verify run: 1.8s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.376 ± 0.011, crit=0.792 ± 0.019 per rating point (14 rating = 1%, 11.088 per %), hit=1.157 ± 0.040 per rating point (10 rating = 1%, 11.566 per %), melee_haste=8.764 ± 0.453

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.1 attack_power points (1.71 DPS) | yes | Hawkeye's Helm (14591, -0.54 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.54 DPS) [world_drop]; Nightscape Headband (8176, -0.63 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.31 DPS) | yes | Scout's Medallion (19536, -0.32 DPS) [rep]; Ghostshard Talisman (7731, -0.39 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.1 attack_power points (1.78 DPS) | yes | Forest Tracker Epaulets (2278, -0.54 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.55 DPS) [dungeon]; Nightscape Shoulders (8192, -0.79 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.8 attack_power points (1.16 DPS) | yes | First Sergeant's Cloak (16340, -0.23 DPS) [pvp]; Hawkeye's Cloak (14593, -0.34 DPS) [world_drop]; Parachute Cloak (10518, -0.44 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 34.1 attack_power points (2.24 DPS) | yes | Wolffear Harness (13110, -0.82 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.89 DPS) [crafted]; Dusky Leather Armor (7374, -0.98 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.31 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.55 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 31.1 attack_power points (2.04 DPS) | yes | Prowler's Leather Gloves (252524, -0.57 DPS) [crafted]; Imperial Leather Gloves (4063, -0.64 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.79 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.97 DPS) | yes | Defiler's Chain Girdle (20152, -0.46 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.57 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.76 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 28.9 attack_power points (1.90 DPS) | yes | Triprunner Dungarees (9624, -0.07 DPS) [quest]; Ferine Leggings (6690, -0.19 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.62 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.1 attack_power points (1.45 DPS) | yes | Prowler's Leather Shoes (252465, -0.10 DPS) [crafted]; Imperial Leather Boots (6431, -0.13 DPS) [dungeon]; Excelsior Boots (4109, -0.18 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.31 DPS) | yes | Legionnaire's Band (19512, -0.07 DPS) [rep]; Field Researcher's Loop (281634, -0.22 DPS) [quest]; Ironspine's Eye (7686, -0.24 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.31 DPS) | yes | Legionnaire's Band (19512, -0.07 DPS) [rep]; Field Researcher's Loop (281634, -0.22 DPS) [quest]; Ironspine's Eye (7686, -0.24 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (111.9 DPS) | yes | Scout's Blade (19544, -3.66 DPS) [pvp]; Coldrage Dagger (10761, -4.49 DPS, sim-verified) [dungeon]; Darkspear Insurgent's Spellblade (272085, -4.55 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 442.0 attack_power points (28.99 DPS) | yes | Coldrage Dagger (10761, -0.52 DPS, sim-verified) [dungeon]; Stonecloth Branch (15963, -28.80 DPS) [world_drop]; Tork Wrench (11855, -28.86 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (111.9 DPS) | yes | Monolithic Bow (9426, -0.25 DPS) [dungeon]; Swiftwind (13038, -0.29 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.08 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 168.8. Weights run: 3.2s. Verify run: 2.2s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.434 ± 0.012, crit=0.912 ± 0.020 per rating point (14 rating = 1%, 12.770 per %), hit=1.539 ± 0.059 per rating point (10 rating = 1%, 15.391 per %), melee_haste=9.419 ± 0.850

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 48.8 attack_power points (4.04 DPS) | yes | Blood Guard's Leather Headband (220851, -0.38 DPS) [vendor]; Embrace of the Lycan (9479, -1.20 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.82 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.6 attack_power points (1.96 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.30 DPS) [quest]; Woven Ivy Necklace (19159, -0.39 DPS) [quest]; Scout's Medallion (19535, -0.53 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.8 attack_power points (2.30 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.25 DPS) [crafted]; Failed Flying Experiment (9647, -0.28 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.1 attack_power points (2.16 DPS) | yes | Blisterbane Wrap (12552, -0.38 DPS) [dungeon]; Dark Phantom Cape (13122, -0.38 DPS) [world_drop]; Duskbat Drape (19982, -0.50 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 46.2 attack_power points (3.82 DPS) | yes | Blazewind Breastplate (11193, -0.84 DPS) [quest]; Fungus Shroud Armor (17742, -0.85 DPS) [dungeon]; Warbear Harness (15064, -1.29 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.5 attack_power points (2.11 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.68 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 41.6 attack_power points (3.44 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.73 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.03 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.15 DPS) | yes | Defiler's Leather Girdle (20193, -0.43 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.49 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.56 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 46.2 attack_power points (3.82 DPS) | yes | Basilisk Hide Pants (1718, -1.33 DPS) [world_drop]; Serpentskin Leggings (8262, -1.41 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -1.44 DPS) [quest] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 30.4 attack_power points (2.52 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Prowler's Leather Boots (252468, -0.13 DPS) [crafted]; Albino Crocscale Boots (17728, -0.14 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 35.4 attack_power points (2.93 DPS) | yes | Legionnaire's Band (19511, -1.03 DPS) [rep]; Masons Fraternity Ring (9533, -1.27 DPS) [quest]; Mark of Kern (2262, -1.27 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.99 DPS) | yes | Legionnaire's Band (19511, -0.09 DPS) [rep]; Masons Fraternity Ring (9533, -0.32 DPS) [quest]; Mark of Kern (2262, -0.33 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (168.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (168.8 DPS) | yes | Molten Heart of the Mountain (249470, -1.15 DPS, sim-verified) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (168.8 DPS) | yes | Barman Shanker (12791, -2.50 DPS, sim-verified) [dungeon]; Scout's Blade (19543, -3.25 DPS) [pvp]; Lifeforce Dirk (10750, -3.47 DPS) [quest] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (42.37 DPS) | yes | Barman Shanker (12791, -19.50 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -42.01 DPS) [quest]; Stonecloth Branch (15963, -42.12 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (168.8 DPS) | yes | Stinging Bow (10624, -0.26 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.26 DPS) [world_drop]; Dark Iron Rifle (16004, -1.97 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 253.8. Weights run: 2.7s. Verify run: 2.1s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.079 ± 0.025, crit=2.249 ± 0.043 per rating point (14 rating = 1%, 31.482 per %), hit=3.228 ± 0.166 per rating point (10 rating = 1%, 32.284 per %), melee_haste=15.611 ± 2.128

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 96.1 attack_power points (7.83 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.42 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 63.5 attack_power points (5.17 DPS) | yes | Mark of Fordring (15411, -0.49 DPS) [quest]; Beads of Ogre Might (22150, -0.59 DPS) [quest]; Medallion of the Dawn (22659, -0.65 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 88.6 attack_power points (7.22 DPS) | yes | Darkspear Pauldrons (272105, -0.08 DPS) [vendor]; Champion's Leather Shoulders (227056, -0.23 DPS) [pvp]; Warlord's Leather Spaulders (231551, -1.03 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 attack_power points (4.91 DPS) | yes | Cape of the Black Baron (13340, -0.74 DPS) [dungeon]; Deathguard's Cloak (20068, -1.30 DPS) [rep]; Windshear Cape (20691, -1.72 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (253.8 DPS) | yes | Warlord's Leather Breastplate (231549, -1.86 DPS) [pvp]; Darkmantle Tunic (226825, -2.27 DPS) [quest]; Tunic of Undead Slaying (23089, -9.44 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (253.8 DPS) | yes | General's Leather Armsplints (16559, -0.08 DPS) [pvp]; Bracers of the Eclipse (18375, -0.47 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.22 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 67.1 attack_power points (5.47 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.62 DPS) [crafted]; Stormshroud Gloves (21278, -2.64 DPS, sim-verified) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 108.2 attack_power points (8.82 DPS) | yes | Belt of Preserved Heads (20216, -1.14 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -3.48 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.82 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 119.1 attack_power points (9.71 DPS) | yes | General's Leather Legguards (231554, +0.00 DPS) [pvp]; Legionnaire's Leather Legguards (227059, -1.74 DPS) [pvp]; Plaguehound Leggings (18736, -2.06 DPS, sim-verified) [dungeon] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 72.5 attack_power points (5.91 DPS) | yes | Fine Dawn Treaders (227815, -1.33 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -1.56 DPS) [dungeon]; Darkmantle Boots (22003, -1.84 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (253.8 DPS) | yes | Tarnished Elven Ring (18500, -1.33 DPS) [dungeon]; Cutthroat's Signet (272408, -1.50 DPS) [vendor]; Naglering (11669, -5.00 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (253.8 DPS) | yes | Tarnished Elven Ring (18500, -0.51 DPS) [dungeon]; Cutthroat's Signet (272408, -0.68 DPS) [vendor]; Naglering (11669, -4.52 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (253.8 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (253.8 DPS) | yes | Royal Seal of Eldre'Thalas (18465, -0.00 DPS) [quest]; Blackhand's Breadth (13965, -1.67 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -2.90 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (253.8 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Felstriker (12590, -9.83 DPS, sim-verified) [dungeon] |
| off_hand | The Lobotomizer (19324) | Frostwolf Clan [rep] | 661.1 attack_power points (53.88 DPS) | yes | Felstriker (12590, +0.00 DPS, sim-verified) [dungeon]; Distracting Dagger (18392, -7.36 DPS) [dungeon]; Scepter of Interminable Focus (22329, -48.69 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (253.8 DPS) | yes | Blackcrow (12651, -0.26 DPS) [dungeon]; The Purifier (22656, -0.57 DPS) [quest]; Dark Iron Rifle (16004, -2.05 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 618.6. Weights run: 2.9s. Verify run: 2.2s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.132 ± 0.026, crit=2.354 ± 0.044 per rating point (14 rating = 1%, 32.962 per %), hit=3.962 ± 0.225 per rating point (10 rating = 1%, 39.618 per %), melee_haste=12.634 ± 2.984

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 112.2 attack_power points (19.10 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Champion's Leather Helm (227057, -0.62 DPS) [pvp]; Outlaw's Collar (279253, -3.06 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 71.6 attack_power points (12.19 DPS) | yes | Beads of Ogre Might (22150, -1.36 DPS) [quest]; Mark of Fordring (15411, -2.15 DPS) [quest]; Medallion of the Dawn (22659, -2.49 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 103.2 attack_power points (17.57 DPS) | yes | Champion's Leather Shoulders (227056, -1.47 DPS) [pvp]; Darkspear Pauldrons (272105, -2.16 DPS) [vendor]; Warlord's Leather Spaulders (231551, -3.21 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 67.6 attack_power points (11.51 DPS) | yes | Cape of the Black Baron (13340, -2.66 DPS) [dungeon]; Deathguard's Cloak (20068, -3.91 DPS) [rep]; Windshear Cape (20691, -4.70 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (618.6 DPS) | yes | Warlord's Leather Breastplate (231549, -3.99 DPS) [pvp]; Darkmantle Tunic (226825, -4.42 DPS) [quest]; Tunic of Undead Slaying (23089, -22.93 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (618.6 DPS) | yes | General's Leather Armsplints (16559, -0.17 DPS) [pvp]; Bracers of the Eclipse (18375, -1.05 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -9.45 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 72.6 attack_power points (12.35 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; Devilsaur Gauntlets (15063, -1.98 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 124.0 attack_power points (21.10 DPS) | yes | Belt of Preserved Heads (20216, -3.42 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -9.71 DPS) [rep]; Ferocity of the Timbermaw (227805, -10.51 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 123.5 attack_power points (21.02 DPS) | yes | General's Leather Legguards (231554, +0.00 DPS) [pvp]; Legionnaire's Leather Legguards (227059, -2.88 DPS) [pvp]; Plaguehound Leggings (18736, -4.22 DPS, sim-verified) [dungeon] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 78.9 attack_power points (13.43 DPS) | yes | Shadowcraft Boots (16711, -3.79 DPS) [dungeon]; Darkmantle Boots (22003, -4.72 DPS) [quest]; Fine Dawn Treaders (227815, -4.86 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (618.6 DPS) | yes | Tarnished Elven Ring (18500, -2.89 DPS) [dungeon]; Cutthroat's Signet (272408, -3.25 DPS) [vendor]; Naglering (11669, -12.31 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (618.6 DPS) | yes | Tarnished Elven Ring (18500, -1.09 DPS) [dungeon]; Cutthroat's Signet (272408, -1.45 DPS) [vendor]; Naglering (11669, -10.91 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (618.6 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.80 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (618.6 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.15 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (618.6 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Felstriker (12590, -24.52 DPS, sim-verified) [dungeon] |
| off_hand | The Lobotomizer (19324) | Frostwolf Clan [rep] | 661.1 attack_power points (112.51 DPS) | yes | Felstriker (12590, +0.00 DPS, sim-verified) [dungeon]; Distracting Dagger (18392, -15.37 DPS) [dungeon]; Scepter of Interminable Focus (22329, -100.16 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (618.6 DPS) | yes | Blackcrow (12651, -0.67 DPS) [dungeon]; The Purifier (22656, -2.22 DPS) [quest]; Dark Iron Rifle (16004, -5.17 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Sentinel's Leather Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket2: Blackhand's Breadth; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

