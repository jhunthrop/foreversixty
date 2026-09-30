# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.3. Weights run: 0.8s. Verify run: 1.1s. 169 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.404 ± 0.062, hit=0.311 ± 0.030, melee_haste=not significant (0.665 ± 0.825)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.38 DPS) | yes | Flying Tiger Goggles (4368, -0.54 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.14 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.24 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.34 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Catacomb Cloak (279899, -0.01 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.1 attack_power points (0.53 DPS) | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.19 DPS) [crafted]; Dark Leather Tunic (2317, -0.24 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 attack_power points (0.24 DPS) | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.07 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.02 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.85 DPS) | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 attack_power points (0.43 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 attack_power points (0.38 DPS) | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.33 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.29 DPS) | yes | Protector's Band (20439, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 5.3 attack_power points (0.25 DPS) | yes | Protector's Band (20439, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon]; The 1 Ring (8350, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.76 DPS) | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Barrens Basher (274744, -1.21 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 attack_power points (10.85 DPS) | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.19 DPS) | yes | Fine Longbow (11304, -0.01 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 169, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6255 Fishing Pole (JEFFTEST); 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 48.3. Weights run: 0.8s. Verify run: 1.1s. 300 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.485 ± 0.060, hit=0.428 ± 0.033, melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 attack_power points (0.48 DPS) | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.11 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Sentinel's Medallion (19541, -0.32 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.38 DPS) [rep]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 attack_power points (0.53 DPS) | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.33 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.48 DPS) | yes | Tigerstrike Mantle (13108, -0.10 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.19 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.77 DPS) | yes | Dusky Leather Armor (7374, -0.08 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.24 DPS) [quest]; Green Leather Armor (4255, -0.38 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.48 DPS) | yes | Jurassic Wristguards (6198, -0.10 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.19 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.77 DPS) | yes | Pilferer's Gloves (7358, -0.43 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -0.44 DPS) [crafted]; Wolfclaw Gloves (1978, -0.48 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.15 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.72 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.25 DPS) | yes | Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.63 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.58 DPS) | yes | Feet of the Lynx (1121, -0.04 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 attack_power points (0.43 DPS) | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Monkey Ring (6748, -0.10 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.48 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -0.12 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.20 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | sim-verified (48.3 DPS) | yes | Swinetusk Shank (6691, -10.42 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.35 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Silver Star (3463, -0.21 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop]; Crystalpine Stinger (13037, -0.24 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Royal Diplomatic Scepter; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (night-elf, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 78.2. Weights run: 0.9s. Verify run: 1.1s. 429 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.925 ± 0.148, hit=0.646 ± 0.078, melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.3 attack_power points (0.66 DPS) | yes | Nightscape Headband (8176, -0.10 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.69 DPS) | yes | Sentinel's Medallion (19540, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.29 DPS) [rep]; Sentinel's Medallion (20444, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 attack_power points (1.15 DPS) | yes | Forest Tracker Epaulets (2278, -0.45 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.3 attack_power points (0.51 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Imperial Cloak (6432, -0.10 DPS) [world_drop]; Parachute Cloak (10518, -0.10 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (73.0 DPS) | yes | Raptorbane Armor (3566, -0.07 DPS) [quest]; Nightscape Tunic (8175, -0.10 DPS) [crafted]; Quillward Harness (10583, -1.55 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.65 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.0 attack_power points (1.63 DPS) | yes | Heavy Earthen Gloves (7359, -0.72 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -0.99 DPS) [crafted]; Shadowskin Gloves (18238, -0.99 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.49 DPS) | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.39 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Petrolspill Leggings (9509, -0.58 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 attack_power points (0.66 DPS) | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.40 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Falcon's Hook (7552, -0.53 DPS) [world_drop]; Ironspine's Eye (7686, -0.53 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 attack_power points (0.60 DPS) | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.14 DPS) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.54 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Dazzling Longsword (869, -1.68 DPS) [world_drop]; Southsea Lamp (9359, -1.95 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (76.6 DPS) | yes | Ardent Custodian (868, -5.20 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -21.88 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (71.4 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Swiftwind (13038, -0.34 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.90 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 429, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 50 (night-elf, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 112.1. Weights run: 0.9s. Verify run: 1.1s. 556 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=6.578 ± 0.398, hit=0.803 ± 0.104, melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 116.1 attack_power points (6.21 DPS) | yes | Eye of Theradras (17715, -1.95 DPS, sim-verified) [dungeon]; Helm of Fire (8348, -5.19 DPS) [crafted]; Lordrec Helmet (10741, -5.25 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 14.6 attack_power points (0.78 DPS) | yes | Sentinel's Medallion (19539, -0.06 DPS) [rep]; Sentinel's Medallion (19540, -0.12 DPS) [rep]; Ghostshard Talisman (7731, -0.75 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 108.1 attack_power points (5.78 DPS) | yes | Sunburn Spaulders (274751, -0.16 DPS, sim-verified) [vendor]; Phytoskin Spaulders (17749, -4.82 DPS) [dungeon]; Forest Tracker Epaulets (2278, -5.12 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (110.4 DPS) | yes | Blackveil Cape (11626, -0.06 DPS) [dungeon]; Duskbat Drape (19982, -0.06 DPS) [quest]; Blisterbane Wrap (12552, -1.75 DPS, sim-verified) [dungeon] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 118.1 attack_power points (6.31 DPS) | yes | Fungus Shroud Armor (17742, -3.07 DPS, sim-verified) [dungeon]; Blazewind Breastplate (11193, -4.93 DPS) [quest]; Quillward Harness (10583, -5.17 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.07 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 112.1 attack_power points (5.99 DPS) | yes | Sergeant Major's Leather Gauntlets (220856, -0.47 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.07 DPS) [crafted]; Shadowskin Gloves (18238, -1.07 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 112.1 attack_power points (5.99 DPS) | yes | Highlander's Lizardhide Girdle (20103, -1.07 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.57 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -4.39 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (110.4 DPS) | yes | Stormshroud Pants (15057, -1.71 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -4.92 DPS) [dungeon]; Basilisk Hide Pants (1718, -5.05 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 attack_power points (1.20 DPS) | yes | Sandstalker Ankleguards (12470, -0.18 DPS) [dungeon]; Sergeant Major's Leather Boots (220860, -0.24 DPS) [vendor]; Whisperwalk Boots (20255, -2.21 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.50 DPS) | yes | Masons Fraternity Ring (9533, -0.66 DPS) [quest]; Insurgent's Band (272065, -0.70 DPS) [vendor]; Insurgent's Band (272066, -0.86 DPS) [vendor] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.07 DPS) | yes | Masons Fraternity Ring (9533, +0.00 DPS, sim-verified) [quest]; Insurgent's Band (272065, -0.27 DPS) [vendor]; Insurgent's Band (272066, -0.43 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (108.7 DPS) | yes | Smoking Heart of the Mountain (11811, -0.60 DPS, sim-verified) [crafted] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | sim-verified (108.7 DPS) | yes | Shadowblade (2163, +0.00 DPS, sim-verified) [world_drop]; Thorium Cestus (250614, -3.15 DPS) [crafted]; Bloodrazor (809, -3.57 DPS) [world_drop] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 attack_power points (29.57 DPS) | yes | Claw of Celebras (17738, -3.80 DPS) [dungeon]; Shadowblade (2163, -13.65 DPS, sim-verified) [world_drop]; Thermotastic Egg Timer (9644, -29.39 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (108.7 DPS) | yes | Skull Splitting Crossbow (13039, -0.09 DPS) [world_drop]; The Silencer (13138, -0.09 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.56 DPS, sim-verified) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Assault Band; trinket1: Guardian Talisman; trinket2: Frozen Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 556, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 60 (night-elf, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 239.6. Weights run: 0.9s. Verify run: 1.2s. 1222 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=0.900 ± 0.162, melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ragefury Eyepatch (11735) (or Bloodvine Lens (19998)) | Blackrock Depths: Guzzler [dungeon] | 230.2 attack_power points (12.33 DPS) | yes | Bloodvine Lens (19998, -0.71 DPS, sim-verified) [crafted]; Outlaw's Collar (279253, -1.51 DPS) [crafted]; Duskwraith Helmet (239560, -1.54 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (239.6 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -5.68 DPS) [quest]; Amulet of the Darkmoon (19491, -6.03 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 152.8 attack_power points (8.18 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, -0.36 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -0.36 DPS) [vendor]; Knight-Lieutenant's Leather Shoulders (220852, -2.46 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 115.1 attack_power points (6.16 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Cloak of the Honor Guard (20073, -3.97 DPS) [rep]; Cape of the Black Baron (13340, -3.97 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-verified (239.6 DPS) | yes | Duskwraith Breastplate (239562, -2.02 DPS) [vendor]; Dawn Armor (252483, -2.91 DPS) [crafted]; Tunic of Undead Slaying (23089, -3.85 DPS, sim-verified) [world] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | sim-verified (239.6 DPS) | yes | Primal Batskin Bracers (19687, -0.82 DPS) [crafted]; Duskwraith Wristguards (239547, -0.82 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -9.84 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 173.0 attack_power points (9.26 DPS) | yes | Marshal's Leather Handgrips (16454, -1.60 DPS) [vendor]; Marshal's Leather Handgrips (231544, -1.60 DPS) [vendor]; Devilsaur Gauntlets (15063, -9.38 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 170.2 attack_power points (9.11 DPS) | yes | Highlander's Leather Girdle (20115, -1.88 DPS) [rep]; Belt of the Archmage (18405, -2.95 DPS) [crafted]; Highlander's Leather Girdle (20045, -7.98 DPS, sim-verified) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 267.9 attack_power points (14.35 DPS) | yes | Knight-Captain's Leather Legguards (16419, -2.02 DPS) [pvp]; Sentinel's Silk Leggings (237815, -2.02 DPS) [vendor]; Stormshroud Pants (15057, -5.01 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 173.0 attack_power points (9.26 DPS) | yes | Duskwraith Treads (239553, -7.36 DPS) [vendor]; Darkmantle Boots (22003, -7.47 DPS) [quest]; Darkmantle Footpads (226831, -7.73 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (239.6 DPS) | yes | Band of the Penitent (13217, -1.34 DPS) [quest]; Ring of Entropy (18543, -1.34 DPS) [world]; Naglering (11669, -8.30 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (239.6 DPS) | yes | Band of the Penitent (13217, -0.48 DPS) [quest]; Ring of Entropy (18543, -0.48 DPS) [world]; Naglering (11669, -7.15 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (239.6 DPS) | yes | Frozen Heart of the Mountain (249469, -8.31 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | sim-verified (239.6 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Misplaced Servo Arm (23221, -57.64 DPS, sim-verified) [world_drop] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (239.6 DPS) | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -58.30 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (239.6 DPS) | yes | Bow of Searing Arrows (2825, -2.06 DPS, sim-verified) [world_drop]; Precisely Calibrated Boomstick (2100, -5.12 DPS) [world_drop]; Ancient Bone Bow (18680, -5.34 DPS) [dungeon] |

**New at 60:** head: Ragefury Eyepatch; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Stormshroud Armor; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Sentinel's Leather Pants; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Blue Dragon; main_hand: Ebon Hand; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1222, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.8. Weights run: 0.8s. Verify run: 1.0s. 173 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.404 ± 0.062, hit=0.311 ± 0.030, melee_haste=not significant (0.665 ± 0.825)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 attack_power points (0.38 DPS) | yes | Flying Tiger Goggles (4368, -0.51 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 attack_power points (0.24 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.32 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 attack_power points (0.34 DPS) | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Prospector's Chestpiece (14562, -0.05 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.32 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 attack_power points (0.24 DPS) | yes | Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 attack_power points (0.29 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.02 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.85 DPS) | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 attack_power points (0.43 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 attack_power points (0.38 DPS) | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.32 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 attack_power points (0.29 DPS) | yes | Legionnaire's Band (20429, -0.10 DPS) [rep]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 5.3 attack_power points (0.25 DPS) | yes | Legionnaire's Band (20429, +0.00 DPS, sim-verified) [rep]; Bounty Hunter's Ring (5351, -0.11 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (11.76 DPS) | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Wingblade (6504, -1.15 DPS) [quest] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 attack_power points (10.85 DPS) | yes | Blackfang (2236, +0.00 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.19 DPS) | yes | Fine Longbow (11304, -0.00 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Pyrewood Signet Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 173, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6255 Fishing Pole (JEFFTEST); 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 30 (troll, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 47.5. Weights run: 0.8s. Verify run: 1.1s. 304 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.485 ± 0.060, hit=0.428 ± 0.033, melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 attack_power points (0.48 DPS) | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.11 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Scout's Medallion (19537, -0.32 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.38 DPS) [rep]; Kaleidoscope Chain (13084, -0.48 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 attack_power points (0.53 DPS) | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.33 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.48 DPS) | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 attack_power points (0.68 DPS) | yes | Panther Armor (6670, -0.17 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.29 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.29 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 attack_power points (0.48 DPS) | yes | Jurassic Wristguards (6198, -0.10 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.19 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.77 DPS) | yes | Pilferer's Gloves (7358, -0.43 DPS, sim-verified) [crafted]; Braced Handguards (6784, -0.43 DPS) [quest]; Fletcher's Gloves (7348, -0.44 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.15 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Deftkin Belt (16659, -0.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.25 DPS) | yes | Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.64 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.58 DPS) | yes | Feet of the Lynx (1121, -0.05 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Vorrel's Boots (7751, -0.19 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 attack_power points (0.43 DPS) | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Monkey Ring (6748, -0.10 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | 322.6 attack_power points (15.48 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Scorn's Focal Dagger (23168, -0.12 DPS) [dungeon]; Bloody Brass Knuckles (7683, -0.20 DPS) [dungeon] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | sim-verified (47.5 DPS) | yes | Swinetusk Shank (6691, -10.06 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -15.35 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Silver Star (3463, -0.21 DPS, sim-verified) [quest]; BKP "Sparrow" Smallbore (3042, -0.24 DPS) [world_drop]; Crystalpine Stinger (13037, -0.24 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Royal Diplomatic Scepter; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 304, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 76.6. Weights run: 0.9s. Verify run: 1.1s. 433 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.925 ± 0.148, hit=0.646 ± 0.078, melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.3 attack_power points (0.66 DPS) | yes | Nightscape Headband (8176, -0.08 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.69 DPS) | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.29 DPS) [rep]; Scout's Medallion (20442, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 attack_power points (1.15 DPS) | yes | Forest Tracker Epaulets (2278, -0.45 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 10.3 attack_power points (0.51 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Wildhunter Cloak (16658, -0.01 DPS) [quest]; Imperial Cloak (6432, -0.10 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (71.9 DPS) | yes | Nightscape Tunic (8175, -0.10 DPS) [crafted]; Dusky Leather Armor (7374, -0.15 DPS) [crafted]; Quillward Harness (10583, -1.62 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.99 DPS) | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.64 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.0 attack_power points (1.63 DPS) | yes | Heavy Earthen Gloves (7359, -0.63 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -0.99 DPS) [crafted]; Shadowskin Gloves (18238, -0.99 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.49 DPS) | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.38 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.29 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Petrolspill Leggings (9509, -0.58 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 attack_power points (0.66 DPS) | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.42 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.99 DPS) | yes | Ring of the Underwood (2951, -0.48 DPS) [world_drop]; Falcon's Hook (7552, -0.53 DPS) [world_drop]; Ironspine's Eye (7686, -0.53 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 attack_power points (0.60 DPS) | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.14 DPS) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (23.54 DPS) | yes | Ardent Custodian (868, +0.00 DPS, sim-verified) [world_drop]; Dazzling Longsword (869, -1.68 DPS) [world_drop]; Southsea Lamp (9359, -1.95 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (74.9 DPS) | yes | Ardent Custodian (868, -4.59 DPS, sim-verified) [world_drop]; Satyr's Rod (15962, -21.88 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (70.3 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.25 DPS) [vendor]; Swiftwind (13038, -0.34 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.89 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 433, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 114.5. Weights run: 0.9s. Verify run: 1.1s. 561 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=6.578 ± 0.398, hit=0.803 ± 0.104, melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Leather Headband (220851) | Lady Palanseer [vendor] | 116.1 attack_power points (6.21 DPS) | yes | Eye of Theradras (17715, -2.40 DPS, sim-verified) [dungeon]; Helm of Fire (8348, -5.19 DPS) [crafted]; Sprightring Helm (17776, -5.31 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 14.6 attack_power points (0.78 DPS) | yes | Scout's Medallion (19535, -0.06 DPS) [rep]; Scout's Medallion (19536, -0.12 DPS) [rep]; Ghostshard Talisman (7731, -0.76 DPS, sim-verified) [dungeon] |
| shoulder | Blood Guard's Leather Shoulders (220853) | Lady Palanseer [vendor] | 108.1 attack_power points (5.78 DPS) | yes | Sunburn Spaulders (274751, -0.58 DPS, sim-verified) [vendor]; Phytoskin Spaulders (17749, -4.82 DPS) [dungeon]; Forest Tracker Epaulets (2278, -5.12 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (113.4 DPS) | yes | Blackveil Cape (11626, -0.06 DPS) [dungeon]; Duskbat Drape (19982, -0.06 DPS) [quest]; Blisterbane Wrap (12552, -1.72 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 118.1 attack_power points (6.31 DPS) | yes | Fungus Shroud Armor (17742, -3.51 DPS, sim-verified) [dungeon]; Blazewind Breastplate (11193, -4.93 DPS) [quest]; Quillward Harness (10583, -5.17 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.07 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 112.1 attack_power points (5.99 DPS) | yes | First Sergeant's Leather Gauntlets (220857, -0.47 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.07 DPS) [crafted]; Shadowskin Gloves (18238, -1.07 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 112.1 attack_power points (5.99 DPS) | yes | Defiler's Lizardhide Girdle (20174, -1.07 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.57 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -4.39 DPS) [rep] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | sim-verified (112.8 DPS) | yes | Stormshroud Pants (15057, -1.12 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -4.92 DPS) [dungeon]; Basilisk Hide Pants (1718, -5.05 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 attack_power points (1.20 DPS) | yes | Sandstalker Ankleguards (12470, -0.18 DPS) [dungeon]; First Sergeant's Leather Boots (220861, -0.24 DPS) [vendor]; Whisperwalk Boots (20255, -2.23 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.50 DPS) | yes | Assault Band (13095, -0.43 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.66 DPS) [quest]; Insurgent's Band (272065, -0.70 DPS) [vendor] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.28 DPS) | yes | Assault Band (13095, -0.31 DPS, sim-verified) [world_drop]; Masons Fraternity Ring (9533, -0.44 DPS) [quest]; Insurgent's Band (272065, -0.48 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (111.7 DPS) | yes | Smoking Heart of the Mountain (11811, -4.21 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (111.7 DPS) | yes | Smoking Heart of the Mountain (11811, -1.08 DPS, sim-verified) [crafted] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | sim-verified (111.7 DPS) | yes | Shadowblade (2163, +0.00 DPS, sim-verified) [world_drop]; Thorium Cestus (250614, -3.15 DPS) [crafted]; Bloodrazor (809, -3.57 DPS) [world_drop] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 attack_power points (29.57 DPS) | yes | Claw of Celebras (17738, -3.80 DPS) [dungeon]; White Bone Shredder (11863, -5.90 DPS) [quest]; Shadowblade (2163, -14.14 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (111.7 DPS) | yes | Skull Splitting Crossbow (13039, -0.09 DPS) [world_drop]; The Silencer (13138, -0.09 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.58 DPS, sim-verified) [world_drop] |

**New at 50:** head: Blood Guard's Leather Headband; neck: Skibi's Pendant; shoulder: Blood Guard's Leather Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 243.2. Weights run: 0.9s. Verify run: 1.2s. 1227 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=0.900 ± 0.162, melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ragefury Eyepatch (11735) (or Bloodvine Lens (19998)) | Blackrock Depths: Guzzler [dungeon] | 230.2 attack_power points (12.33 DPS) | yes | Bloodvine Lens (19998, -0.70 DPS, sim-verified) [crafted]; Outlaw's Collar (279253, -1.51 DPS) [crafted]; Duskwraith Helmet (239560, -1.54 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (243.2 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -5.68 DPS) [quest]; Amulet of the Darkmoon (19491, -6.03 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 152.8 attack_power points (8.18 DPS) | yes | Champion's Leather Shoulders (23258, -0.36 DPS) [vendor]; Champion's Leather Shoulders (227056, -0.36 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, -2.88 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 115.1 attack_power points (6.16 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Deathguard's Cloak (20068, -3.97 DPS) [rep]; Cape of the Black Baron (13340, -3.97 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-verified (243.2 DPS) | yes | Duskwraith Breastplate (239562, -2.02 DPS) [vendor]; Dawn Armor (252483, -2.91 DPS) [crafted]; Tunic of Undead Slaying (23089, -4.04 DPS, sim-verified) [world] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | sim-verified (243.2 DPS) | yes | Primal Batskin Bracers (19687, -0.82 DPS) [crafted]; Duskwraith Wristguards (239547, -0.82 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -8.07 DPS, sim-verified) [world] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 173.0 attack_power points (9.26 DPS) | yes | General's Leather Mitts (16560, -1.60 DPS) [vendor]; General's Leather Mitts (231555, -1.60 DPS) [vendor]; Devilsaur Gauntlets (15063, -7.70 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 170.2 attack_power points (9.11 DPS) | yes | Defiler's Leather Girdle (20193, -1.88 DPS) [rep]; Belt of the Archmage (18405, -2.95 DPS) [crafted]; Defiler's Leather Girdle (20190, -6.31 DPS, sim-verified) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 267.9 attack_power points (14.35 DPS) | yes | Legionnaire's Leather Leggings (16508, -2.02 DPS) [pvp]; Sentinel's Silk Leggings (237815, -2.02 DPS) [vendor]; Stormshroud Pants (15057, -3.41 DPS, sim-verified) [crafted] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 173.0 attack_power points (9.26 DPS) | yes | Darkmantle Footpads (226831, -6.99 DPS, sim-verified) [vendor]; Duskwraith Treads (239553, -7.36 DPS) [vendor]; Darkmantle Boots (22003, -7.47 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (243.2 DPS) | yes | Band of the Penitent (13217, -1.34 DPS) [quest]; Ring of Entropy (18543, -1.34 DPS) [world]; Naglering (11669, -6.53 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (243.2 DPS) | yes | Band of the Penitent (13217, -0.48 DPS) [quest]; Ring of Entropy (18543, -0.48 DPS) [world]; Naglering (11669, -5.38 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (243.2 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (243.2 DPS) | yes | Frozen Heart of the Mountain (249469, -4.51 DPS, sim-verified) [crafted] |
| main_hand | Ebon Hand (19170) | Blacksmithing [crafted] | sim-verified (243.2 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Misplaced Servo Arm (23221, -60.30 DPS, sim-verified) [world_drop] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (243.2 DPS) | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -60.67 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (243.2 DPS) | yes | Bow of Searing Arrows (2825, -2.06 DPS, sim-verified) [world_drop]; Precisely Calibrated Boomstick (2100, -5.12 DPS) [world_drop]; Ancient Bone Bow (18680, -5.34 DPS) [dungeon] |

**New at 60:** head: Ragefury Eyepatch; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Chromatic Cloak; chest: Stormshroud Armor; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Sentinel's Leather Pants; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Ebon Hand; off_hand: Shadowsong's Sorrow; ranged: The Purifier

No-known-source sample (15 of 1227, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

