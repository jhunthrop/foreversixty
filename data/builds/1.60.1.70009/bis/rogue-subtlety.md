# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.7. Weights run: 1.1s. Verify run: 1.4s. 194 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.021 ± 0.003 per rating point (14 rating = 1%, 0.290 per %), hit=0.059 ± 0.003 per rating point (10 rating = 1%, 0.591 per %), melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.37 DPS) | yes | Lucky Fishing Hat (19972, -0.45 DPS, sim-verified) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 attack_power points (0.28 DPS) | yes | Erudite's Amulet (277204, -0.11 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.23 DPS) | yes | Forest Leather Mantle (4709, -0.28 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, +0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Hide of Lupos (3018, -0.09 DPS) [world] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.2 attack_power points (0.51 DPS) | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.18 DPS) [crafted]; Prospector's Chestpiece (14562, -0.23 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 attack_power points (0.23 DPS) | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.09 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.09 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.09 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.82 DPS) | yes | Deviate Scale Belt (6468, -0.52 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Dark Leather Belt (4249, -0.63 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 9.1 attack_power points (0.41 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 attack_power points (0.37 DPS) | yes | Blackened Defias Boots (10402, -0.09 DPS) [dungeon]; Bristlebark Boots (14568, -0.14 DPS) [world_drop]; Footpads of the Fang (10411, -0.28 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.28 DPS) | yes | Protector's Band (20439, -0.09 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; The 1 Ring (8350, -0.23 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 4.3 attack_power points (0.19 DPS) | yes | Protector's Band (20439, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The 1 Ring (8350, -0.15 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (10.35 DPS) | yes | Evocator's Blade (2567, -0.65 DPS) [dungeon]; Buzzer Blade (2169, -1.48 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.07 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 221.5 attack_power points (10.03 DPS) | yes | Evocator's Blade (2567, -6.38 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 194, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14145 Cursed Felblade

### Band 30 (night-elf, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 39.8. Weights run: 1.3s. Verify run: 1.6s. 328 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.030 ± 0.003 per rating point (14 rating = 1%, 0.421 per %), hit=0.024 ± 0.002 per rating point (10 rating = 1%, 0.238 per %), melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.1 attack_power points (0.47 DPS) | yes | Brawler's Leather Hood (252504, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.11 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.65 DPS) | yes | Sentinel's Medallion (19541, -0.31 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.46 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 attack_power points (0.52 DPS) | yes | Insignia Mantle (4721, -0.19 DPS) [world_drop]; Cloudy Gustwoven Spaulders (277043, -0.19 DPS) [crafted]; Mantle of Thieves (2264, -0.33 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.46 DPS) | yes | Tigerstrike Mantle (13108, -0.08 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.18 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.74 DPS) | yes | Dusky Leather Armor (7374, -0.06 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.23 DPS) [quest]; Green Leather Armor (4255, -0.37 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.46 DPS) | yes | Unearthed Bands (9428, -0.11 DPS, sim-verified) [dungeon]; Jurassic Wristguards (6198, -0.18 DPS) [world]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.74 DPS) | yes | Pilferer's Gloves (7358, -0.42 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -0.46 DPS) [world_drop]; Ebon Vise (7690, -0.46 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.12 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.69 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.21 DPS) | yes | Troll's Bane Leggings (13114, -0.55 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.60 DPS) [crafted]; Petrolspill Leggings (9509, -0.61 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.56 DPS) | yes | Feet of the Lynx (1121, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor]; Insignia Boots (4055, -0.19 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.1 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, -0.09 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, -0.09 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (14.97 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.02 DPS) [vendor]; Torturing Poker (7682, -1.18 DPS) [dungeon]; Thornspike (6681, -1.57 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (14.88 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.43 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -14.83 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Silver Star (3463, -0.21 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.23 DPS) [world_drop]; Crystalpine Stinger (13037, -0.23 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (night-elf, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 81.2. Weights run: 1.3s. Verify run: 1.6s. 456 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.047 ± 0.006 per rating point (14 rating = 1%, 0.663 per %), hit=0.058 ± 0.006 per rating point (10 rating = 1%, 0.583 per %), melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.2 attack_power points (0.65 DPS) | yes | White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop]; Nightscape Headband (8176, -0.11 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.99 DPS) | yes | Ghostshard Talisman (7731, -0.38 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19540, -0.44 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 attack_power points (1.15 DPS) | yes | Forest Tracker Epaulets (2278, -0.59 DPS) [world_drop]; Flintrock Shoulders (7755, -0.64 DPS) [dungeon]; Nightscape Shoulders (8192, -0.76 DPS, sim-verified) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.1 attack_power points (0.50 DPS) | yes | Imperial Cloak (6432, -0.10 DPS) [dungeon]; Parachute Cloak (10518, -0.10 DPS) [crafted]; Wolfmaster Cape (6314, -0.65 DPS, sim-verified) [dungeon] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 19.2 attack_power points (0.95 DPS) | yes | Raptorbane Armor (3566, -0.16 DPS) [quest]; Nightscape Tunic (8175, -0.20 DPS) [crafted]; Wolffear Harness (13110, -0.72 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Imperial Leather Bracers (4061, -0.59 DPS) [dungeon]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.64 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.7 attack_power points (1.02 DPS) | yes | Skulker's Leather Gloves (252525, -0.52 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.52 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.35 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.49 DPS) | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.38 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest]; Hawkeye's Breeches (14595, -0.59 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.2 attack_power points (0.65 DPS) | yes | Imperial Leather Boots (6431, -0.10 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.56 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Insurgent's Band (272066, -0.40 DPS) [vendor]; Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Falcon's Hook (7552, -0.54 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Insurgent's Band (272066, -0.51 DPS, sim-verified) [vendor]; Falcon's Hook (7552, -0.54 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.50 DPS) | yes | Black Menace (6831, -2.93 DPS) [quest]; Coldrage Dagger (10761, -2.93 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272085, -3.44 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 432.8 attack_power points (21.44 DPS) | yes | Black Menace (6831, -0.43 DPS, sim-verified) [quest]; Satyr's Rod (15962, -21.39 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (81.2 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Swiftwind (13038, -0.34 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.89 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 50 (night-elf, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 118.7. Weights run: 1.4s. Verify run: 2.0s. 582 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.218 ± 0.061, crit=0.464 ± 0.025 per rating point (14 rating = 1%, 6.499 per %), hit=0.053 ± 0.009 per rating point (10 rating = 1%, 0.529 per %), melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 36.5 attack_power points (1.97 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.73 DPS) [vendor]; Helm of Fire (8348, -0.85 DPS) [crafted]; Embrace of the Lycan (9479, -1.16 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.08 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19539, -0.29 DPS) [rep]; Ghostshard Talisman (7731, -0.32 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.4 attack_power points (1.37 DPS) | yes | Phytoskin Spaulders (17749, -0.27 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Leather Shoulders (220852, -0.56 DPS) [vendor]; Penance Spaulders (11963, -0.65 DPS) [quest] |
| back | Blisterbane Wrap (12552) (or Dark Phantom Cape (13122)) | Blackrock Depths: Anvilrage Overseer [dungeon] | 18.3 attack_power points (0.99 DPS) | yes | Dark Phantom Cape (13122, +0.00 DPS, sim-verified) [world_drop]; Blackveil Cape (11626, -0.07 DPS) [dungeon]; Duskbat Drape (19982, -0.07 DPS) [quest] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 30.5 attack_power points (1.65 DPS) | yes | Blazewind Breastplate (11193, +0.00 DPS, sim-verified) [quest]; Knight's Leather Armor (220854, -0.29 DPS) [vendor]; Quillward Harness (10583, -0.40 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.08 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.36 DPS) [crafted]; Pridelord Bands (14672, -0.42 DPS) [world_drop] |
| hands | Darkmantle Grips (226828) (or Raider Gloves (272100)) | Mokvar [vendor] | 26.8 attack_power points (1.45 DPS) | yes | Raider Gloves (272100, +0.00 DPS, sim-verified) [vendor]; Gloves of Holy Might (867, -0.02 DPS) [world_drop]; Sergeant Major's Leather Gauntlets (220856, -0.34 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) (or Highlander's Leather Girdle (20116)) | Blackrock Depths: Eviscerator [dungeon] | 30.0 attack_power points (1.62 DPS) | yes | Highlander's Leather Girdle (20115, -0.19 DPS) [rep]; Highlander's Chain Girdle (20090, -0.32 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.41 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Knight's Leather Pants (220858, -0.05 DPS) [vendor]; Keeper's Woolies (14668, -0.15 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.4 attack_power points (1.32 DPS) | yes | Sandstalker Ankleguards (12470, -0.20 DPS) [dungeon]; Whisperwalk Boots (20255, -0.26 DPS, sim-verified) [quest]; Sergeant Major's Leather Boots (220860, -0.34 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.5 attack_power points (1.11 DPS) | yes | Mark of Kern (2262, -0.03 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.19 DPS) [quest]; Insurgent's Band (272065, -0.30 DPS) [vendor] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (1.08 DPS) | yes | Mark of Kern (2262, +0.00 DPS, sim-verified) [dungeon]; Masons Fraternity Ring (9533, -0.16 DPS) [quest]; Insurgent's Band (272065, -0.27 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Lifeforce Dirk (10750, -0.42 DPS) [quest]; Charstone Dirk (17710, -0.42 DPS) [dungeon]; Shadowblade (2163, -2.50 DPS, sim-verified) [world_drop] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (27.69 DPS) | yes | Thermotastic Egg Timer (9644, -27.50 DPS) [quest]; Satyr's Rod (15962, -27.63 DPS) [world_drop]; Windchaser Orb (15965, -27.63 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, -0.17 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -1.78 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; back: Blisterbane Wrap; chest: Fungus Shroud Armor; hands: Darkmantle Grips; waist: Girdle of Beastial Fury; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Assault Band; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 582, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 60 (night-elf, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 226.7. Weights run: 1.4s. Verify run: 2.0s. 1335 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=0.580 ± 0.036 per rating point (14 rating = 1%, 8.118 per %), hit=not significant (0.046 ± 0.013) per rating point (10 rating = 1%, 0.458 per %), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 66.9 attack_power points (3.56 DPS) | yes | Lieutenant Commander's Leather Helm (227055, -1.19 DPS) [pvp]; Field Marshal's Leather Mask (16455, -1.41 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 32.1 attack_power points (1.71 DPS) | yes | Will of the Martyr (17044, -0.11 DPS) [quest]; Mark of Fordring (15411, -0.29 DPS) [quest]; Imperial Jewel (11933, -1.32 DPS, sim-verified) [dungeon] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 51.3 attack_power points (2.72 DPS) | yes | Duskwraith Pauldrons (239559, -0.57 DPS) [vendor]; Darkspear Pauldrons (272105, -0.60 DPS) [vendor]; Highlander's Lizardhide Shoulders (20060, -0.96 DPS, sim-verified) [rep] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 39.9 attack_power points (2.12 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Howler's Furs (272414, -0.61 DPS) [vendor]; Stormpike Soldier's Cloak (19084, -0.85 DPS) [rep] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Duskwraith Breastplate (239562, -0.16 DPS) [vendor]; Nightbrace Tunic (12603, -0.72 DPS) [dungeon]; Tunic of Undead Slaying (23089, -7.38 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Duskwraith Bracers (239555, -0.31 DPS) [vendor]; Marshal's Leather Armsplints (16460, -0.71 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.73 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 49.9 attack_power points (2.65 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.32 DPS) [pvp]; Skul's Fingerbone Claws (13395, -0.53 DPS) [dungeon]; Cadaverous Gloves (14640, -8.16 DPS, sim-verified) [dungeon] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 47.5 attack_power points (2.53 DPS) | yes | Cadaverous Belt (14636, -0.40 DPS) [dungeon]; Girdle of Beastial Fury (11686, -0.93 DPS) [dungeon]; Highlander's Leather Girdle (20045, -6.32 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 68.1 attack_power points (3.62 DPS) | yes | Cadaverous Leggings (14638, -0.85 DPS) [dungeon]; Sentinel's Leather Pants (237818, -1.06 DPS) [vendor]; Devilsaur Leggings (15062, -10.96 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 49.9 attack_power points (2.65 DPS) | yes | Highlander's Leather Boots (20052, -1.05 DPS) [rep]; Darkmantle Footpads (226831, -1.13 DPS) [vendor]; Pads of the Dread Wolf (13210, -8.77 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Naglering (11669, -6.88 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 0.0 attack_power points (0.00 DPS) | yes | Blackstone Ring (17713, -0.22 DPS) [dungeon]; Assault Band (13095, -0.24 DPS) [world_drop]; Naglering (11669, -6.51 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | 0.0 attack_power points (0.00 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18465, -4.72 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS, sim-verified) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -1.00 DPS) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (226.7 DPS) | yes | Distracting Dagger (18392, -3.59 DPS) [dungeon]; The Lobotomizer (19324, -4.66 DPS, sim-verified) [rep]; Tome of Knowledge (13385, -33.42 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.29 DPS) [world_drop]; Dark Iron Rifle (16004, -1.98 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Medallion of the Dawn; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Riphook

No-known-source sample (15 of 1335, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.5. Weights run: 1.1s. Verify run: 1.5s. 193 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.021 ± 0.003 per rating point (14 rating = 1%, 0.290 per %), hit=0.059 ± 0.003 per rating point (10 rating = 1%, 0.591 per %), melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.37 DPS) | yes | Lucky Fishing Hat (19972, -0.44 DPS, sim-verified) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 attack_power points (0.28 DPS) | yes | Erudite's Amulet (277204, -0.11 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.23 DPS) | yes | Forest Leather Mantle (4709, -0.27 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Catacomb Cloak (279899, +0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Hide of Lupos (3018, -0.09 DPS) [world] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 attack_power points (0.32 DPS) | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Prospector's Chestpiece (14562, -0.05 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.27 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 attack_power points (0.23 DPS) | yes | Wolf Bracers (4794, -0.05 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.09 DPS) [vendor]; Spare Part Bindings (279875, -0.09 DPS) [quest] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 attack_power points (0.28 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.09 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.09 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.82 DPS) | yes | Deviate Scale Belt (6468, -0.54 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.63 DPS) [world]; Dark Leather Belt (4249, -0.63 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 9.1 attack_power points (0.41 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 attack_power points (0.37 DPS) | yes | Blackened Defias Boots (10402, -0.09 DPS) [dungeon]; Bristlebark Boots (14568, -0.14 DPS) [world_drop]; Footpads of the Fang (10411, -0.27 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.28 DPS) | yes | Legionnaire's Band (20429, -0.09 DPS) [rep]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 4.3 attack_power points (0.19 DPS) | yes | Legionnaire's Band (20429, +0.00 DPS, sim-verified) [rep]; Bounty Hunter's Ring (5351, -0.06 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (10.35 DPS) | yes | Assassin's Blade (1935, -0.32 DPS) [dungeon]; Evocator's Blade (2567, -0.65 DPS) [dungeon]; Buzzer Blade (2169, -1.48 DPS) [dungeon] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | 222.7 attack_power points (10.09 DPS) | yes | Assassin's Blade (1935, -0.15 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Edward's Knife; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 39.0. Weights run: 1.3s. Verify run: 1.6s. 328 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.030 ± 0.003 per rating point (14 rating = 1%, 0.421 per %), hit=0.024 ± 0.002 per rating point (10 rating = 1%, 0.238 per %), melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.1 attack_power points (0.47 DPS) | yes | Brawler's Leather Hood (252504, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.11 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.65 DPS) | yes | Scout's Medallion (19537, -0.33 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.46 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 attack_power points (0.52 DPS) | yes | Insignia Mantle (4721, -0.19 DPS) [world_drop]; Cloudy Gustwoven Spaulders (277043, -0.19 DPS) [crafted]; Mantle of Thieves (2264, -0.33 DPS, sim-verified) [dungeon] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.46 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 attack_power points (0.66 DPS) | yes | Panther Armor (6670, -0.17 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.28 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.28 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.46 DPS) | yes | Unearthed Bands (9428, -0.11 DPS, sim-verified) [dungeon]; Jurassic Wristguards (6198, -0.18 DPS) [world]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.74 DPS) | yes | Braced Handguards (6784, -0.42 DPS) [quest]; Pilferer's Gloves (7358, -0.44 DPS, sim-verified) [crafted]; Ebon Vise (7690, -0.46 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.12 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.37 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.21 DPS) | yes | Troll's Bane Leggings (13114, -0.55 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.60 DPS) [crafted]; Petrolspill Leggings (9509, -0.65 DPS, sim-verified) [dungeon] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 8.1 attack_power points (0.38 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS, sim-verified) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.1 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, -0.09 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, -0.11 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (14.97 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.02 DPS) [vendor]; Torturing Poker (7682, -1.18 DPS) [dungeon]; Thornspike (6681, -1.57 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (14.88 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.35 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -14.83 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Silver Star (3463, -0.22 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.23 DPS) [world_drop]; Crystalpine Stinger (13037, -0.23 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wildhunter Cloak; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 79.3. Weights run: 1.3s. Verify run: 1.6s. 456 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.047 ± 0.006 per rating point (14 rating = 1%, 0.663 per %), hit=0.058 ± 0.006 per rating point (10 rating = 1%, 0.583 per %), melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.2 attack_power points (0.65 DPS) | yes | Nightscape Headband (8176, -0.10 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.99 DPS) | yes | Ghostshard Talisman (7731, -0.38 DPS, sim-verified) [dungeon]; Scout's Medallion (19536, -0.44 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 attack_power points (1.15 DPS) | yes | Forest Tracker Epaulets (2278, -0.59 DPS) [world_drop]; Flintrock Shoulders (7755, -0.64 DPS) [dungeon]; Nightscape Shoulders (8192, -0.75 DPS, sim-verified) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.1 attack_power points (0.50 DPS) | yes | Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Parachute Cloak (10518, -0.10 DPS) [crafted]; Wildhunter Cloak (16658, -0.64 DPS, sim-verified) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 19.2 attack_power points (0.95 DPS) | yes | Nightscape Tunic (8175, -0.20 DPS) [crafted]; Dusky Leather Armor (7374, -0.25 DPS) [crafted]; Wolffear Harness (13110, -0.70 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Imperial Leather Bracers (4061, -0.59 DPS) [dungeon]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.63 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.7 attack_power points (1.02 DPS) | yes | Skulker's Leather Gloves (252525, -0.52 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.52 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.33 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.49 DPS) | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.38 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest]; Hawkeye's Breeches (14595, -0.59 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.2 attack_power points (0.65 DPS) | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.10 DPS) [dungeon]; Dusky Boots (7390, -0.10 DPS) [crafted] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Insurgent's Band (272066, -0.40 DPS) [vendor]; Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Falcon's Hook (7552, -0.54 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Insurgent's Band (272066, -0.50 DPS, sim-verified) [vendor]; Falcon's Hook (7552, -0.54 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.50 DPS) | yes | Coldrage Dagger (10761, -2.93 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272085, -3.44 DPS) [vendor]; Hypnotic Blade (7714, -4.92 DPS) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 432.8 attack_power points (21.44 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -21.39 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (79.3 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Swiftwind (13038, -0.34 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.88 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 119.2. Weights run: 1.4s. Verify run: 2.0s. 583 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.218 ± 0.061, crit=0.464 ± 0.025 per rating point (14 rating = 1%, 6.499 per %), hit=0.053 ± 0.009 per rating point (10 rating = 1%, 0.529 per %), melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 36.5 attack_power points (1.97 DPS) | yes | Blood Guard's Leather Headband (220851, -0.73 DPS) [vendor]; Helm of Fire (8348, -0.85 DPS) [crafted]; Embrace of the Lycan (9479, -1.19 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.08 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Scout's Medallion (19535, -0.29 DPS) [rep]; Ghostshard Talisman (7731, -0.32 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.4 attack_power points (1.37 DPS) | yes | Phytoskin Spaulders (17749, -0.31 DPS, sim-verified) [dungeon]; Blood Guard's Leather Shoulders (220853, -0.56 DPS) [vendor]; Penance Spaulders (11963, -0.65 DPS) [quest] |
| back | Blisterbane Wrap (12552) (or Dark Phantom Cape (13122)) | Blackrock Depths: Anvilrage Overseer [dungeon] | 18.3 attack_power points (0.99 DPS) | yes | Dark Phantom Cape (13122, +0.00 DPS, sim-verified) [world_drop]; Blackveil Cape (11626, -0.07 DPS) [dungeon]; Duskbat Drape (19982, -0.07 DPS) [quest] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 30.5 attack_power points (1.65 DPS) | yes | Blazewind Breastplate (11193, +0.00 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -0.29 DPS) [vendor]; Quillward Harness (10583, -0.40 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.08 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.36 DPS) [crafted]; Pridelord Bands (14672, -0.42 DPS) [world_drop] |
| hands | Darkmantle Grips (226828) (or Raider Gloves (272100)) | Mokvar [vendor] | 26.8 attack_power points (1.45 DPS) | yes | Raider Gloves (272100, +0.00 DPS, sim-verified) [vendor]; Gloves of Holy Might (867, -0.02 DPS) [world_drop]; First Sergeant's Leather Gauntlets (220857, -0.34 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) (or Defiler's Leather Girdle (20192)) | Blackrock Depths: Eviscerator [dungeon] | 30.0 attack_power points (1.62 DPS) | yes | Defiler's Leather Girdle (20193, -0.19 DPS) [rep]; Defiler's Chain Girdle (20152, -0.32 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.41 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Stone Guard's Leather Pants (220859, -0.05 DPS) [vendor]; Keeper's Woolies (14668, -0.15 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.4 attack_power points (1.32 DPS) | yes | Sandstalker Ankleguards (12470, -0.20 DPS) [dungeon]; Whisperwalk Boots (20255, -0.24 DPS, sim-verified) [quest]; First Sergeant's Leather Boots (220861, -0.34 DPS) [vendor] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.30 DPS) | yes | Mark of Kern (2262, -0.22 DPS) [dungeon]; Assault Band (13095, -0.22 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.38 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.5 attack_power points (1.11 DPS) | yes | Mark of Kern (2262, -0.03 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.19 DPS) [quest]; Assault Band (13095, -0.49 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (119.2 DPS) | yes | Frozen Heart of the Mountain (249469, -2.27 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (119.2 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (119.2 DPS) | yes | Lifeforce Dirk (10750, -0.42 DPS) [quest]; Charstone Dirk (17710, -0.42 DPS) [dungeon]; Shadowblade (2163, -2.15 DPS, sim-verified) [world_drop] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (27.69 DPS) | yes | Thermotastic Egg Timer (9644, -27.50 DPS) [quest]; Satyr's Rod (15962, -27.63 DPS) [world_drop]; Windchaser Orb (15965, -27.63 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (119.2 DPS) | yes | Stinging Bow (10624, -0.17 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -1.73 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; back: Blisterbane Wrap; chest: Fungus Shroud Armor; hands: Darkmantle Grips; waist: Girdle of Beastial Fury; feet: Albino Crocscale Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 583, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 225.1. Weights run: 1.4s. Verify run: 2.1s. 1336 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=0.580 ± 0.036 per rating point (14 rating = 1%, 8.118 per %), hit=not significant (0.046 ± 0.013) per rating point (10 rating = 1%, 0.458 per %), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 66.9 attack_power points (3.56 DPS) | yes | Champion's Leather Helm (227057, -1.19 DPS) [pvp]; Warlord's Leather Helm (16561, -1.41 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 32.1 attack_power points (1.71 DPS) | yes | Will of the Martyr (17044, -0.11 DPS) [quest]; Mark of Fordring (15411, -0.29 DPS) [quest]; Imperial Jewel (11933, -1.53 DPS, sim-verified) [dungeon] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 51.3 attack_power points (2.72 DPS) | yes | Duskwraith Pauldrons (239559, -0.57 DPS) [vendor]; Darkspear Pauldrons (272105, -0.60 DPS) [vendor]; Defiler's Lizardhide Shoulders (20175, -1.03 DPS, sim-verified) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 39.9 attack_power points (2.12 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Howler's Furs (272414, -0.61 DPS) [vendor]; Frostwolf Legionnaire's Cloak (19083, -0.85 DPS) [rep] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Duskwraith Breastplate (239562, -0.16 DPS) [vendor]; Nightbrace Tunic (12603, -0.72 DPS) [dungeon]; Tunic of Undead Slaying (23089, -7.40 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Duskwraith Bracers (239555, -0.31 DPS) [vendor]; General's Leather Armsplints (16559, -0.71 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.80 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 49.9 attack_power points (2.65 DPS) | yes | Blood Guard's Leather Vices (16499, -0.32 DPS) [pvp]; Skul's Fingerbone Claws (13395, -0.53 DPS) [dungeon]; Cadaverous Gloves (14640, -8.51 DPS, sim-verified) [dungeon] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 47.5 attack_power points (2.53 DPS) | yes | Cadaverous Belt (14636, -0.40 DPS) [dungeon]; Girdle of Beastial Fury (11686, -0.93 DPS) [dungeon]; Defiler's Leather Girdle (20190, -6.73 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 68.1 attack_power points (3.62 DPS) | yes | Cadaverous Leggings (14638, -0.85 DPS) [dungeon]; Sentinel's Leather Pants (237818, -1.06 DPS) [vendor]; Devilsaur Leggings (15062, -11.87 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 49.9 attack_power points (2.65 DPS) | yes | Defiler's Leather Boots (20186, -1.05 DPS) [rep]; Darkmantle Footpads (226831, -1.13 DPS) [vendor]; Pads of the Dread Wolf (13210, -9.11 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | 0.0 attack_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Naglering (11669, -7.22 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 0.0 attack_power points (0.00 DPS) | yes | White Bone Band (11862, -0.03 DPS) [quest]; Blackstone Ring (17713, -0.22 DPS) [dungeon]; Naglering (11669, -6.91 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 attack_power points (0.00 DPS) | yes | Counterattack Lodestone (18537, -1.08 DPS) [dungeon]; Hand of Justice (11815, -1.19 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18465, -1.28 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (225.1 DPS) | yes | Distracting Dagger (18392, -3.59 DPS) [dungeon]; The Lobotomizer (19324, -3.70 DPS, sim-verified) [rep]; Tome of Knowledge (13385, -33.42 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | 0.0 attack_power points (0.00 DPS) | yes | Malgen's Long Bow (22318, -0.11 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.29 DPS) [world_drop]; Dark Iron Rifle (16004, -1.96 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Medallion of the Dawn; shoulder: Defiler's Leather Shoulders; back: Deathguard's Cloak; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Riphook

No-known-source sample (15 of 1336, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

