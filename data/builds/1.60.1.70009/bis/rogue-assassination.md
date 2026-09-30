# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.7. Weights run: 0.8s. Verify run: 1.0s. 194 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.067 ± 0.020, crit=0.159 ± 0.008 per rating point (14 rating = 1%, 2.221 per %), hit=0.061 ± 0.003 per rating point (10 rating = 1%, 0.606 per %), melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 attack_power points (0.40 DPS) | yes | Lucky Fishing Hat (19972, -0.66 DPS, sim-verified) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.4 attack_power points (0.30 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 attack_power points (0.25 DPS) | yes | Forest Leather Mantle (4709, -0.42 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Hide of Lupos (3018, -0.10 DPS) [world]; Catacomb Cloak (279899, -0.10 DPS, sim-verified) [quest] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.7 attack_power points (0.55 DPS) | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.20 DPS) [crafted]; Prospector's Chestpiece (14562, -0.25 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.3 attack_power points (0.25 DPS) | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.08 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.10 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.84 DPS) | yes | Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Dark Leather Belt (4249, -0.64 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 9.6 attack_power points (0.45 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.5 attack_power points (0.40 DPS) | yes | Blackened Defias Boots (10402, -0.10 DPS) [dungeon]; Bristlebark Boots (14568, -0.15 DPS) [world_drop]; Footpads of the Fang (10411, -0.36 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 attack_power points (0.30 DPS) | yes | Protector's Band (20439, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; The 1 Ring (8350, -0.25 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 4.5 attack_power points (0.21 DPS) | yes | Protector's Band (20439, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; The 1 Ring (8350, -0.16 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.61 DPS) | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Barrens Basher (274744, -1.20 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 attack_power points (10.71 DPS) | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.3 attack_power points (0.20 DPS) | yes | Fine Longbow (11304, -0.06 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 194, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14145 Cursed Felblade

### Band 30 (night-elf, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 51.5. Weights run: 0.9s. Verify run: 1.1s. 328 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.097 ± 0.016, crit=0.236 ± 0.009 per rating point (14 rating = 1%, 3.306 per %), hit=0.077 ± 0.004 per rating point (10 rating = 1%, 0.769 per %), melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 attack_power points (0.52 DPS) | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.14 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.16 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.66 DPS) | yes | Sentinel's Medallion (19541, -0.23 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.35 DPS) [rep]; Kaleidoscope Chain (13084, -0.45 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 attack_power points (0.57 DPS) | yes | Insignia Mantle (4721, -0.21 DPS) [world_drop]; Cloudy Gustwoven Spaulders (277043, -0.21 DPS) [crafted]; Mantle of Thieves (2264, -0.34 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.47 DPS) | yes | Tigerstrike Mantle (13108, +0.00 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.11 DPS) [world_drop]; Cloak of Night (4447, -0.16 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.75 DPS) | yes | Dusky Leather Armor (7374, +0.00 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.19 DPS) [quest]; Green Leather Armor (4255, -0.34 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.47 DPS) | yes | Unearthed Bands (9428, -0.11 DPS, sim-verified) [dungeon]; Jurassic Wristguards (6198, -0.16 DPS) [world]; Hawkeye's Bracers (14590, -0.16 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.75 DPS) | yes | Pilferer's Gloves (7358, -0.34 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -0.44 DPS) [world_drop]; Ebon Vise (7690, -0.44 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.13 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.67 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.23 DPS) | yes | Petrolspill Leggings (9509, -0.46 DPS, sim-verified) [dungeon]; Troll's Bane Leggings (13114, -0.50 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.55 DPS) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.57 DPS) | yes | Insignia Boots (4055, -0.12 DPS, sim-verified) [world_drop]; Feet of the Lynx (1121, -0.15 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.15 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 attack_power points (0.47 DPS) | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon]; Protector's Band (19517, -0.16 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, +0.00 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon]; Protector's Band (19517, -0.11 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.21 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -0.12 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.20 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | sim-verified (51.5 DPS) | yes | Swinetusk Shank (6691, -13.46 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.08 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Silver Star (3463, -0.16 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop]; Crystalpine Stinger (13037, -0.22 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Royal Diplomatic Scepter; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (night-elf, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 88.9. Weights run: 0.9s. Verify run: 1.1s. 456 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.105 ± 0.014, crit=0.198 ± 0.005 per rating point (14 rating = 1%, 2.769 per %), hit=0.055 ± 0.002 per rating point (10 rating = 1%, 0.555 per %), melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 14.4 attack_power points (0.48 DPS) | yes | Nightscape Headband (8176, -0.05 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.07 DPS) [crafted]; Hawkeye's Helm (14591, -0.07 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Ghostshard Talisman (7731, -0.21 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19540, -0.26 DPS) [rep]; Sentinel's Medallion (19541, -0.37 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.81 DPS) | yes | Forest Tracker Epaulets (2278, -0.40 DPS) [world_drop]; Nightscape Shoulders (8192, -0.42 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.44 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 11.0 attack_power points (0.37 DPS) | yes | Imperial Cloak (6432, -0.07 DPS) [dungeon]; Parachute Cloak (10518, -0.07 DPS) [crafted]; Wolfmaster Cape (6314, -0.28 DPS, sim-verified) [dungeon] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 21.0 attack_power points (0.70 DPS) | yes | Nightscape Tunic (8175, -0.15 DPS) [crafted]; Raptorbane Armor (3566, -0.17 DPS) [quest]; Wolffear Harness (13110, -0.39 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Cultist's Armguards (270032, -0.35 DPS, sim-verified) [quest]; Imperial Leather Bracers (4061, -0.37 DPS) [dungeon]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 22.8 attack_power points (0.76 DPS) | yes | Skulker's Leather Gloves (252525, -0.39 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.39 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.56 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.00 DPS) | yes | Highlander's Leather Girdle (20117, -0.20 DPS) [rep]; Highlander's Chain Girdle (20090, -0.21 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.87 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.20 DPS) [quest]; Hawkeye's Breeches (14595, -0.35 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.4 attack_power points (0.48 DPS) | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Dusky Boots (7390, -0.07 DPS) [crafted] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.67 DPS) | yes | Insurgent's Band (272066, -0.27 DPS) [vendor]; Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Insurgent's Band (272066, -0.28 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (15.89 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Dazzling Longsword (869, -1.13 DPS) [world_drop]; Southsea Lamp (9359, -1.29 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (88.9 DPS) | yes | Satyr's Rod (15962, -14.77 DPS) [world_drop]; Ardent Custodian (868, -50.68 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.17 DPS) [vendor]; Swiftwind (13038, -0.21 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.49 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 50 (night-elf, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 123.8. Weights run: 1.0s. Verify run: 1.5s. 582 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.107 ± 0.016, crit=0.234 ± 0.006 per rating point (14 rating = 1%, 3.272 per %), hit=0.068 ± 0.003 per rating point (10 rating = 1%, 0.678 per %), melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 36.2 attack_power points (1.22 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.55 DPS) [vendor]; Helm of Fire (8348, -0.59 DPS) [crafted]; Embrace of the Lycan (9479, -1.21 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.20 DPS) [dungeon]; Sentinel's Medallion (19539, -0.23 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.82 DPS) | yes | Phytoskin Spaulders (17749, -0.18 DPS, sim-verified) [dungeon]; Penance Spaulders (11963, -0.40 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.40 DPS) [crafted] |
| back | Blisterbane Wrap (12552) (or Dark Phantom Cape (13122)) | Blackrock Depths: Anvilrage Overseer [dungeon] | 16.6 attack_power points (0.56 DPS) | yes | Dark Phantom Cape (13122, +0.00 DPS, sim-verified) [world_drop]; Blackveil Cape (11626, -0.04 DPS) [dungeon]; Duskbat Drape (19982, -0.04 DPS) [quest] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 27.7 attack_power points (0.93 DPS) | yes | Blazewind Breastplate (11193, +0.00 DPS, sim-verified) [quest]; Knight's Leather Armor (220854, -0.19 DPS) [vendor]; Quillward Harness (10583, -0.22 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.26 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [world_drop] |
| hands | Darkmantle Grips (226828) (or Raider Gloves (272100)) | Mokvar [vendor] | 24.4 attack_power points (0.82 DPS) | yes | Raider Gloves (272100, +0.00 DPS, sim-verified) [vendor]; Gloves of Holy Might (867, -0.04 DPS) [world_drop]; Sergeant Major's Leather Gauntlets (220856, -0.24 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) (or Highlander's Leather Girdle (20116)) | Blackrock Depths: Eviscerator [dungeon] | 30.0 attack_power points (1.01 DPS) | yes | Highlander's Chain Girdle (20090, -0.20 DPS) [rep]; Highlander's Leather Girdle (20117, -0.20 DPS) [rep]; Highlander's Leather Girdle (20116, -0.50 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.88 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Knight's Leather Pants (220858, -0.14 DPS) [vendor]; Keeper's Woolies (14668, -0.17 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 attack_power points (0.75 DPS) | yes | Sandstalker Ankleguards (12470, -0.11 DPS) [dungeon]; Sergeant Major's Leather Boots (220860, -0.14 DPS) [vendor]; Whisperwalk Boots (20255, -0.23 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.7 attack_power points (0.70 DPS) | yes | Mark of Kern (2262, -0.02 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.17 DPS) [quest]; Insurgent's Band (272065, -0.19 DPS) [vendor] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.67 DPS) | yes | Mark of Kern (2262, +0.00 DPS, sim-verified) [dungeon]; Masons Fraternity Ring (9533, -0.15 DPS) [quest]; Insurgent's Band (272065, -0.17 DPS) [vendor] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | 0.0 attack_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| trinket2 | Sanctified Orb (20512) | Forging the Mightstone [quest] | 0.0 attack_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | sim-verified (123.8 DPS) | yes | Bloodrazor (809, +0.00 DPS) [world_drop]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hammer of the Northern Wind (810, -1.47 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (18.39 DPS) | yes | Claw of Celebras (17738, -2.12 DPS) [dungeon]; Thorium Cestus (250614, -4.34 DPS, sim-verified) [crafted]; Thermotastic Egg Timer (9644, -18.28 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, -0.05 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.05 DPS) [world_drop]; Dark Iron Rifle (16004, -1.56 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; back: Blisterbane Wrap; chest: Fungus Shroud Armor; hands: Darkmantle Grips; waist: Girdle of Beastial Fury; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Assault Band; trinket1: Diamond Flask; trinket2: Sanctified Orb; main_hand: Hanzo Sword; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 582, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 60 (night-elf, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 246.6. Weights run: 1.0s. Verify run: 1.4s. 1335 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=0.295 ± 0.007 per rating point (14 rating = 1%, 4.124 per %), hit=0.086 ± 0.004 per rating point (10 rating = 1%, 0.863 per %), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 61.3 attack_power points (2.03 DPS) | yes | Lieutenant Commander's Leather Helm (227055, -0.67 DPS) [pvp]; Duskwraith Mask (239550, -0.80 DPS) [vendor]; Lieutenant Commander's Leather Helm (23312, -9.91 DPS, sim-verified) [vendor] |
| neck | Imperial Jewel (11933) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 32.0 attack_power points (1.06 DPS) | yes | Medallion of the Dawn (22659, -0.13 DPS) [quest]; Will of the Martyr (17044, -0.16 DPS, sim-verified) [quest]; Mark of Fordring (15411, -0.19 DPS) [quest] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 50.4 attack_power points (1.67 DPS) | yes | Duskwraith Pauldrons (239559, -0.37 DPS) [vendor]; Darkspear Pauldrons (272105, -0.52 DPS) [vendor]; Highlander's Lizardhide Shoulders (20060, -0.91 DPS, sim-verified) [rep] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 39.7 attack_power points (1.31 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Howler's Furs (272414, -0.36 DPS) [vendor]; Stormpike Soldier's Cloak (19084, -0.52 DPS) [rep] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (246.6 DPS) | yes | Duskwraith Breastplate (239562, -0.29 DPS) [vendor]; Nightbrace Tunic (12603, -0.44 DPS) [dungeon]; Tunic of Undead Slaying (23089, -6.55 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (246.6 DPS) | yes | Duskwraith Bracers (239555, -0.20 DPS) [vendor]; Marshal's Leather Armsplints (16460, -0.46 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.39 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 44.6 attack_power points (1.48 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.04 DPS) [pvp]; Skul's Fingerbone Claws (13395, -0.15 DPS) [dungeon]; Cadaverous Gloves (14640, -8.46 DPS, sim-verified) [dungeon] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 42.3 attack_power points (1.40 DPS) | yes | Highlander's Leather Girdle (20045, -0.14 DPS) [rep]; Girdle of Beastial Fury (11686, -0.41 DPS) [dungeon]; Cadaverous Belt (14636, -8.28 DPS, sim-verified) [dungeon] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 62.4 attack_power points (2.07 DPS) | yes | Devilsaur Leggings (15062, -0.41 DPS) [crafted]; Knight-Captain's Leather Legguards (23299, -0.78 DPS) [vendor]; Cadaverous Leggings (14638, -13.22 DPS, sim-verified) [dungeon] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 44.6 attack_power points (1.48 DPS) | yes | Highlander's Leather Boots (20052, -0.50 DPS) [rep]; Knight-Lieutenant's Leather Walkers (23285, -0.55 DPS) [vendor]; Pads of the Dread Wolf (13210, -8.99 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (246.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.07 DPS) [quest]; Naglering (11669, -6.91 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (246.6 DPS) | yes | Blackstone Ring (17713, -0.00 DPS) [dungeon]; Assault Band (13095, -0.03 DPS) [world_drop]; Naglering (11669, -6.66 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (246.6 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -5.92 DPS, sim-verified) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-verified (246.6 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -2.05 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (246.6 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Alcor's Sunrazor (14555, -2.29 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 813.5 attack_power points (26.97 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -5.64 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -7.76 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (246.6 DPS) | yes | Malgen's Long Bow (22318, -0.07 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.20 DPS) [world_drop]; Dark Iron Rifle (16004, -1.73 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Imperial Jewel; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Royal Seal of Eldre'Thalas; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1335, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

## Horde

### Band 20 (troll, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.0. Weights run: 0.8s. Verify run: 1.0s. 193 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.067 ± 0.020, crit=0.159 ± 0.008 per rating point (14 rating = 1%, 2.221 per %), hit=0.061 ± 0.003 per rating point (10 rating = 1%, 0.606 per %), melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 attack_power points (0.40 DPS) | yes | Lucky Fishing Hat (19972, -0.64 DPS, sim-verified) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.4 attack_power points (0.30 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 attack_power points (0.25 DPS) | yes | Forest Leather Mantle (4709, -0.40 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Catacomb Cloak (279899, -0.09 DPS, sim-verified) [quest]; Hide of Lupos (3018, -0.10 DPS) [world] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.5 attack_power points (0.35 DPS) | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Prospector's Chestpiece (14562, -0.05 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.33 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.3 attack_power points (0.25 DPS) | yes | Wolf Bracers (4794, -0.09 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor]; Spare Part Bindings (279875, -0.10 DPS) [quest] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.10 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.84 DPS) | yes | Deviate Scale Belt (6468, -0.57 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.64 DPS) [world]; Dark Leather Belt (4249, -0.64 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 9.6 attack_power points (0.45 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.5 attack_power points (0.40 DPS) | yes | Blackened Defias Boots (10402, -0.10 DPS) [dungeon]; Bristlebark Boots (14568, -0.15 DPS) [world_drop]; Footpads of the Fang (10411, -0.36 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 attack_power points (0.30 DPS) | yes | Legionnaire's Band (20429, -0.10 DPS) [rep]; Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 4.5 attack_power points (0.21 DPS) | yes | Legionnaire's Band (20429, +0.00 DPS, sim-verified) [rep]; Bounty Hunter's Ring (5351, -0.06 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.61 DPS) | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Wingblade (6504, -1.12 DPS) [quest] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 attack_power points (10.71 DPS) | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.3 attack_power points (0.20 DPS) | yes | Fine Longbow (11304, -0.06 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 50.8. Weights run: 0.9s. Verify run: 1.1s. 328 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.097 ± 0.016, crit=0.236 ± 0.009 per rating point (14 rating = 1%, 3.306 per %), hit=0.077 ± 0.004 per rating point (10 rating = 1%, 0.769 per %), melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 attack_power points (0.52 DPS) | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.16 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.66 DPS) | yes | Scout's Medallion (19537, -0.21 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.35 DPS) [rep]; Kaleidoscope Chain (13084, -0.45 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 attack_power points (0.57 DPS) | yes | Insignia Mantle (4721, -0.21 DPS) [world_drop]; Cloudy Gustwoven Spaulders (277043, -0.21 DPS) [crafted]; Mantle of Thieves (2264, -0.34 DPS, sim-verified) [dungeon] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.47 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Tigerstrike Mantle (13108, -0.06 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.11 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 15.4 attack_power points (0.72 DPS) | yes | Panther Armor (6670, -0.23 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.31 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.31 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.47 DPS) | yes | Unearthed Bands (9428, -0.11 DPS, sim-verified) [dungeon]; Jurassic Wristguards (6198, -0.16 DPS) [world]; Hawkeye's Bracers (14590, -0.16 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.75 DPS) | yes | Pilferer's Gloves (7358, -0.32 DPS, sim-verified) [crafted]; Braced Handguards (6784, -0.39 DPS) [quest]; Ebon Vise (7690, -0.44 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.13 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.36 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.23 DPS) | yes | Petrolspill Leggings (9509, -0.45 DPS, sim-verified) [dungeon]; Troll's Bane Leggings (13114, -0.50 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.55 DPS) [crafted] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Vorrel's Boots (7751), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 8.8 attack_power points (0.41 DPS) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS, sim-verified) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 attack_power points (0.47 DPS) | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon]; Legionnaire's Band (19513, -0.16 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Monkey Ring (6748, +0.00 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon]; Legionnaire's Band (19513, -0.11 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.21 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -0.12 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.20 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | sim-verified (50.8 DPS) | yes | Swinetusk Shank (6691, -13.32 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.08 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.42 DPS) | yes | Silver Star (3463, -0.14 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop]; Crystalpine Stinger (13037, -0.22 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wildhunter Cloak; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Insignia Boots; finger1: Ironspine's Eye; finger2: Insurgent's Band; main_hand: Royal Diplomatic Scepter; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 86.9. Weights run: 0.9s. Verify run: 1.1s. 456 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.105 ± 0.014, crit=0.198 ± 0.005 per rating point (14 rating = 1%, 2.769 per %), hit=0.055 ± 0.002 per rating point (10 rating = 1%, 0.555 per %), melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 14.4 attack_power points (0.48 DPS) | yes | Nightscape Headband (8176, -0.05 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.07 DPS) [crafted]; Hawkeye's Helm (14591, -0.07 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Ghostshard Talisman (7731, -0.21 DPS, sim-verified) [dungeon]; Scout's Medallion (19536, -0.26 DPS) [rep]; Scout's Medallion (19537, -0.37 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.81 DPS) | yes | Forest Tracker Epaulets (2278, -0.40 DPS) [world_drop]; Nightscape Shoulders (8192, -0.42 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.44 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 11.0 attack_power points (0.37 DPS) | yes | Wolfmaster Cape (6314, -0.04 DPS) [dungeon]; Parachute Cloak (10518, -0.07 DPS) [crafted]; Wildhunter Cloak (16658, -0.28 DPS, sim-verified) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 21.0 attack_power points (0.70 DPS) | yes | Nightscape Tunic (8175, -0.15 DPS) [crafted]; Dusky Leather Armor (7374, -0.18 DPS) [crafted]; Wolffear Harness (13110, -0.38 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Cultist's Armguards (270032, -0.35 DPS, sim-verified) [quest]; Imperial Leather Bracers (4061, -0.37 DPS) [dungeon]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 22.8 attack_power points (0.76 DPS) | yes | Skulker's Leather Gloves (252525, -0.39 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.39 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.54 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.00 DPS) | yes | Defiler's Leather Girdle (20191, -0.20 DPS) [rep]; Defiler's Chain Girdle (20152, -0.21 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.87 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.20 DPS) [quest]; Hawkeye's Breeches (14595, -0.35 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.4 attack_power points (0.48 DPS) | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Dusky Boots (7390, -0.07 DPS) [crafted] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (0.67 DPS) | yes | Insurgent's Band (272066, -0.27 DPS) [vendor]; Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Insurgent's Band (272066, -0.28 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (15.89 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Dazzling Longsword (869, -1.13 DPS) [world_drop]; Southsea Lamp (9359, -1.29 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (86.9 DPS) | yes | Satyr's Rod (15962, -14.77 DPS) [world_drop]; Ardent Custodian (868, -48.95 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.17 DPS) [vendor]; Swiftwind (13038, -0.21 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.49 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 456, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 126.4. Weights run: 1.0s. Verify run: 1.5s. 583 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.107 ± 0.016, crit=0.234 ± 0.006 per rating point (14 rating = 1%, 3.272 per %), hit=0.068 ± 0.003 per rating point (10 rating = 1%, 0.678 per %), melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 36.2 attack_power points (1.22 DPS) | yes | Blood Guard's Leather Headband (220851, -0.55 DPS) [vendor]; Helm of Fire (8348, -0.59 DPS) [crafted]; Embrace of the Lycan (9479, -1.06 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.67 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.20 DPS) [dungeon]; Scout's Medallion (19535, -0.23 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (0.82 DPS) | yes | Phytoskin Spaulders (17749, -0.19 DPS, sim-verified) [dungeon]; Penance Spaulders (11963, -0.40 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.40 DPS) [crafted] |
| back | Blisterbane Wrap (12552) (or Dark Phantom Cape (13122)) | Blackrock Depths: Anvilrage Overseer [dungeon] | 16.6 attack_power points (0.56 DPS) | yes | Dark Phantom Cape (13122, +0.00 DPS, sim-verified) [world_drop]; Blackveil Cape (11626, -0.04 DPS) [dungeon]; Duskbat Drape (19982, -0.04 DPS) [quest] |
| chest | Fungus Shroud Armor (17742) | Maraudon: Meshlok the Harvester [dungeon] | 27.7 attack_power points (0.93 DPS) | yes | Blazewind Breastplate (11193, +0.00 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -0.19 DPS) [vendor]; Quillward Harness (10583, -0.22 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.67 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.26 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [world_drop] |
| hands | Darkmantle Grips (226828) (or Raider Gloves (272100)) | Mokvar [vendor] | 24.4 attack_power points (0.82 DPS) | yes | Raider Gloves (272100, +0.00 DPS, sim-verified) [vendor]; Gloves of Holy Might (867, -0.04 DPS) [world_drop]; First Sergeant's Leather Gauntlets (220857, -0.24 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) (or Defiler's Leather Girdle (20192)) | Blackrock Depths: Eviscerator [dungeon] | 30.0 attack_power points (1.01 DPS) | yes | Defiler's Chain Girdle (20152, -0.20 DPS) [rep]; Defiler's Leather Girdle (20191, -0.20 DPS) [rep]; Defiler's Leather Girdle (20192, -0.49 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.88 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Stone Guard's Leather Pants (220859, -0.14 DPS) [vendor]; Keeper's Woolies (14668, -0.17 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 attack_power points (0.75 DPS) | yes | Sandstalker Ankleguards (12470, -0.11 DPS) [dungeon]; First Sergeant's Leather Boots (220861, -0.14 DPS) [vendor]; Whisperwalk Boots (20255, -0.20 DPS, sim-verified) [quest] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (0.81 DPS) | yes | Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.29 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.7 attack_power points (0.70 DPS) | yes | Mark of Kern (2262, -0.02 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.17 DPS) [quest]; Assault Band (13095, -2.76 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (126.4 DPS) | yes | Diamond Flask (20130, -2.84 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (126.4 DPS) | yes | Diamond Flask (20130, -1.03 DPS, sim-verified) [quest] |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | sim-verified (126.4 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Thorium Cestus (250614, +0.00 DPS) [crafted]; Hookfang Shanker (11635, -1.99 DPS, sim-verified) [dungeon] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (18.39 DPS) | yes | Claw of Celebras (17738, -2.12 DPS) [dungeon]; White Bone Shredder (11863, -3.45 DPS) [quest]; Thorium Cestus (250614, -5.05 DPS, sim-verified) [crafted] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (126.4 DPS) | yes | Stinging Bow (10624, -0.05 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.05 DPS) [world_drop]; Dark Iron Rifle (16004, -1.48 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; back: Blisterbane Wrap; chest: Fungus Shroud Armor; hands: Darkmantle Grips; waist: Girdle of Beastial Fury; feet: Albino Crocscale Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hanzo Sword; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 583, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 237.6. Weights run: 1.0s. Verify run: 1.4s. 1336 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=0.295 ± 0.007 per rating point (14 rating = 1%, 4.124 per %), hit=0.086 ± 0.004 per rating point (10 rating = 1%, 0.863 per %), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 61.3 attack_power points (2.03 DPS) | yes | Champion's Leather Helm (227057, -0.67 DPS) [pvp]; Duskwraith Mask (239550, -0.80 DPS) [vendor]; Champion's Leather Helm (23257, -9.81 DPS, sim-verified) [vendor] |
| neck | Imperial Jewel (11933) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 32.0 attack_power points (1.06 DPS) | yes | Medallion of the Dawn (22659, -0.13 DPS) [quest]; Will of the Martyr (17044, -0.15 DPS, sim-verified) [quest]; Mark of Fordring (15411, -0.19 DPS) [quest] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 50.4 attack_power points (1.67 DPS) | yes | Duskwraith Pauldrons (239559, -0.37 DPS) [vendor]; Darkspear Pauldrons (272105, -0.52 DPS) [vendor]; Defiler's Lizardhide Shoulders (20175, -0.96 DPS, sim-verified) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 39.7 attack_power points (1.31 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Howler's Furs (272414, -0.36 DPS) [vendor]; Frostwolf Legionnaire's Cloak (19083, -0.52 DPS) [rep] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (237.6 DPS) | yes | Duskwraith Breastplate (239562, -0.29 DPS) [vendor]; Nightbrace Tunic (12603, -0.44 DPS) [dungeon]; Tunic of Undead Slaying (23089, -6.35 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (237.6 DPS) | yes | Duskwraith Bracers (239555, -0.20 DPS) [vendor]; General's Leather Armsplints (16559, -0.46 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.35 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 44.6 attack_power points (1.48 DPS) | yes | Blood Guard's Leather Vices (16499, -0.04 DPS) [pvp]; Skul's Fingerbone Claws (13395, -0.15 DPS) [dungeon]; Cadaverous Gloves (14640, -8.48 DPS, sim-verified) [dungeon] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 42.3 attack_power points (1.40 DPS) | yes | Defiler's Leather Girdle (20190, -0.14 DPS) [rep]; Girdle of Beastial Fury (11686, -0.41 DPS) [dungeon]; Cadaverous Belt (14636, -8.27 DPS, sim-verified) [dungeon] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 62.4 attack_power points (2.07 DPS) | yes | Devilsaur Leggings (15062, -0.41 DPS) [crafted]; Legionnaire's Leather Legguards (22880, -0.78 DPS) [vendor]; Cadaverous Leggings (14638, -13.72 DPS, sim-verified) [dungeon] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 44.6 attack_power points (1.48 DPS) | yes | Defiler's Leather Boots (20186, -0.50 DPS) [rep]; Blood Guard's Leather Walkers (22856, -0.55 DPS) [vendor]; Pads of the Dread Wolf (13210, -8.98 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (237.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.07 DPS) [quest]; Naglering (11669, -6.96 DPS, sim-verified) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | sim-verified (237.6 DPS) | yes | Don Julio's Band (19325, -0.10 DPS) [rep]; Blackstone Ring (17713, -0.10 DPS) [dungeon]; Naglering (11669, -1.80 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (237.6 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (237.6 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS, sim-verified) [quest]; Counterattack Lodestone (18537, -0.68 DPS) [dungeon]; Hand of Justice (11815, -0.75 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (237.6 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Alcor's Sunrazor (14555, -0.16 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 813.5 attack_power points (26.97 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -3.27 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -7.76 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (237.6 DPS) | yes | Malgen's Long Bow (22318, -0.07 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.20 DPS) [world_drop]; Dark Iron Rifle (16004, -1.65 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Imperial Jewel; shoulder: Defiler's Leather Shoulders; back: Deathguard's Cloak; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Signet Ring of the Bronze Dragonflight; finger2: White Bone Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1336, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

