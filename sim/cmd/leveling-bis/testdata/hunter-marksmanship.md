# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.1. Weights run: 0.7s. Verify run: 1.2s. 189 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.004, agility=2.567 ± 0.314, crit=0.505 ± 0.072 per rating point (14 rating = 1%, 7.077 per %), hit=0.151 ± 0.027 per rating point (10 rating = 1%, 1.512 per %), melee_haste=not significant (7.284 ± 3.386)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 20.5 ranged_attack_power points (1.23 DPS) | yes | Lucky Fishing Hat (19972, -1.02 DPS, sim-verified) [quest] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Forest Leather Mantle (4709, -0.77 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Cape of the Brotherhood (5193, +0.00 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.31 DPS) [world_drop]; Hide of Lupos (3018, -0.31 DPS) [world] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 28.2 ranged_attack_power points (1.70 DPS) | yes | Brawler's Leather Armor (252490, -0.35 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.62 DPS) [crafted]; Prospector's Chestpiece (14562, -0.77 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Bravo's Armbands (270015, -0.15 DPS) [quest]; Ratchet Wristwraps (274742, -0.31 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Gloves of the Fang (10413, -0.08 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.31 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.31 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.08 DPS) | yes | Dusty Belt (279897, -0.31 DPS) [quest]; Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Dark Leather Belt (4249, -0.46 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 23.1 ranged_attack_power points (1.39 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.15 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.5 ranged_attack_power points (1.23 DPS) | yes | Footpads of the Fang (10411, -0.26 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.31 DPS) [dungeon]; Bristlebark Boots (14568, -0.46 DPS) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; The 1 Ring (8350, -0.77 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.3 ranged_attack_power points (0.62 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.46 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 26.5 ranged_attack_power points (1.59 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.51 DPS) [world]; Pearl-encrusted Spear (1406, -0.67 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.8 ranged_attack_power points (10.75 DPS) | yes | Lil Timmy's Peashooter (13136, -0.69 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.33 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.57 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 189, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (dwarf, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 104.0. Weights run: 0.6s. Verify run: 1.3s. 555 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.008, agility=2.000 ± 0.016, crit=0.832 ± 0.162 per rating point (14 rating = 1%, 11.643 per %), hit=0.218 ± 0.040 per rating point (10 rating = 1%, 2.175 per %), melee_haste=not significant (2.790 ± 4.937)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 26.0 ranged_attack_power points (1.52 DPS) | yes | Nightscape Headband (8176, +0.00 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -0.12 DPS) [crafted]; Skullsplitter Helm (1624, -0.23 DPS) [world] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 22.0 ranged_attack_power points (1.28 DPS) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS, sim-verified) [quest]; Sentinel's Medallion (19541, -0.35 DPS) [rep]; Ghostshard Talisman (7731, -0.47 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 34.0 ranged_attack_power points (1.98 DPS) | yes | Nightscape Shoulders (8192, +0.00 DPS, sim-verified) [crafted]; Forest Tracker Epaulets (2278, -0.70 DPS) [world_drop]; Flintrock Shoulders (7755, -0.82 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Imperial Cloak (6432, -0.23 DPS) [world_drop]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 38.0 ranged_attack_power points (2.22 DPS) | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.47 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.47 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.23 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (1.87 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS, sim-verified) [world_drop]; Dragonscale Gauntlets (8347, -0.49 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.70 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.75 DPS) | yes | Highlander's Chain Girdle (20090, +0.00 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20117, -0.35 DPS) [rep]; Scorpashi Sash (14652, -0.58 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 42.0 ranged_attack_power points (2.45 DPS) | yes | Triprunner Dungarees (9624, +0.00 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.82 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.82 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 26.0 ranged_attack_power points (1.52 DPS) | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.23 DPS) [world_drop]; Stalker's Mail Boots (252562, -0.23 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Mark of Kern (2262, -0.00 DPS) [dungeon]; Falcon's Hook (7552, -0.12 DPS) [world_drop]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Mark of Kern (2262, +0.00 DPS, sim-verified) [dungeon]; Falcon's Hook (7552, -0.12 DPS) [world_drop]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 0.0 ranged_attack_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 18.0 ranged_attack_power points (1.05 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS, sim-verified) [crafted]; Satyr's Rod (15962, -0.93 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (104.0 DPS) | yes | Shadowforge Bushmaster (9422, -0.46 DPS) [dungeon]; Monolithic Bow (9426, -0.85 DPS) [dungeon]; Bow of Searing Arrows (2825, -72.92 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Assault Band; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 555, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 226.5. Weights run: 0.7s. Verify run: 1.2s. 1539 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=1.302 ± 0.249 per rating point (14 rating = 1%, 18.230 per %), hit=0.315 ± 0.055 per rating point (10 rating = 1%, 3.147 per %), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Visor (239532) | Leonid Barthalomew the Revered [vendor] | 176.5 ranged_attack_power points (10.08 DPS) | yes | Field Marshal's Chain Greathelm (231562, -5.04 DPS) [vendor]; Field Marshal's Chain Helm (16465, -5.16 DPS) [vendor]; Dawnstalker Headpiece (239540, -5.62 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 42.2 ranged_attack_power points (2.41 DPS) | yes | Amulet of the Darkmoon (19491, -0.45 DPS, sim-verified) [quest]; Imperial Jewel (11933, -0.58 DPS) [dungeon]; Pendant of Celerity (22340, -0.68 DPS) [dungeon] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | 129.4 ranged_attack_power points (7.39 DPS) | yes | Dawnstalker Spaulders (239542, -2.43 DPS, sim-verified) [vendor]; Darkspear Pauldrons (272105, -3.26 DPS) [vendor]; Darkspear Epaulets (272106, -3.26 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.0 ranged_attack_power points (2.85 DPS) | yes | Cloak of the Honor Guard (20073, -0.86 DPS, sim-verified) [rep]; Shifting Cloak (18511, -0.91 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.91 DPS) [dungeon] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (226.5 DPS) | yes | Dawnstalker Tunic (239543, -1.81 DPS) [vendor]; Dawn Armor (252483, -5.09 DPS) [crafted]; Tunic of Undead Slaying (23089, -16.75 DPS, sim-verified) [world] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | sim-verified (226.5 DPS) | yes | Dawnstalker Wristguards (239544, -0.69 DPS) [vendor]; Bracers of the Eclipse (18375, -2.53 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.54 DPS, sim-verified) [world] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | 132.2 ranged_attack_power points (7.55 DPS) | yes | Dawnstalker Handguards (239539, -1.51 DPS, sim-verified) [vendor]; Chromatic Gauntlets (19157, -4.00 DPS) [crafted]; Marshal's Chain Grips (16463, -4.11 DPS) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 108.2 ranged_attack_power points (6.18 DPS) | yes | Dawnstalker Girdle (239538, -1.36 DPS, sim-verified) [vendor]; Dense Timbermaw Belt (227807, -2.53 DPS) [vendor]; Molten Belt (19163, -2.98 DPS) [crafted] |
| legs | Dawnstalker Leggings (239533) | Leonid Barthalomew the Revered [vendor] | 174.5 ranged_attack_power points (9.96 DPS) | yes | Sentinel's Chain Leggings (237819, -3.71 DPS) [vendor]; Sentinel's Chain Leggings (22748, -4.75 DPS) [rep]; Dawnstalker Legguards (239541, -5.83 DPS, sim-verified) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | 140.2 ranged_attack_power points (8.01 DPS) | yes | Dawnstalker Boots (239537, -3.64 DPS, sim-verified) [vendor]; Marshal's Chain Sabatons (231561, -4.74 DPS) [vendor]; Marshal's Chain Boots (16462, -4.86 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (226.5 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.23 DPS) [quest]; Wrath of Cenarius (21190, -6.22 DPS, sim-verified) [quest] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (226.5 DPS) | yes | Cutthroat's Signet (272408, -0.36 DPS) [vendor]; Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Wrath of Cenarius (21190, -5.07 DPS, sim-verified) [quest] |
| trinket1 | Earthstrike (21180) | Champion's Battlegear [quest] | sim-verified (226.5 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Shard of the Fallen Star (21891, -5.35 DPS, sim-verified) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (226.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -6.79 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Core Marksman Rifle (18282) | Engineering [crafted] | sim-verified (226.5 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [vendor]; Grand Marshal's Repeater (234586, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -10.09 DPS, sim-verified) [crafted] |

**New at 60:** head: Dawnstalker Visor; neck: Medallion of the Dawn; shoulder: Dawnstalker Pauldrons; back: Cape of the Black Baron; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Leggings; feet: Dawnstalker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Earthstrike; main_hand: Legionite Glaive; ranged: Core Marksman Rifle

No-known-source sample (15 of 1539, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.8. Weights run: 0.7s. Verify run: 1.2s. 192 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.004, agility=2.567 ± 0.314, crit=0.505 ± 0.072 per rating point (14 rating = 1%, 7.077 per %), hit=0.151 ± 0.027 per rating point (10 rating = 1%, 1.512 per %), melee_haste=not significant (7.284 ± 3.386)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 20.5 ranged_attack_power points (1.23 DPS) | yes | Lucky Fishing Hat (19972, -1.02 DPS, sim-verified) [quest] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Forest Leather Mantle (4709, -0.88 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Cape of the Brotherhood (5193, -0.30 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.31 DPS) [world_drop]; Hide of Lupos (3018, -0.31 DPS) [world] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | sim-verified (75.8 DPS) | yes | Dark Leather Tunic (2317, -0.15 DPS) [crafted]; Prospector's Chestpiece (14562, -0.15 DPS) [world_drop]; Brawler's Leather Armor (252490, -0.76 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.31 DPS) [vendor]; Spare Part Bindings (279875, -0.31 DPS) [quest] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Gloves of the Fang (10413, -0.08 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.31 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.31 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.08 DPS) | yes | Dusty Belt (279897, -0.31 DPS) [quest]; Deviate Scale Belt (6468, -0.45 DPS, sim-verified) [crafted]; Dark Leather Belt (4249, -0.46 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501), Leggings of the Fang (10410)) | Leatherworking [crafted] | 23.1 ranged_attack_power points (1.39 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Bluegill Breeches (3022, -0.15 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.5 ranged_attack_power points (1.23 DPS) | yes | Footpads of the Fang (10411, -0.25 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.31 DPS) [dungeon]; Bristlebark Boots (14568, -0.46 DPS) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Bounty Hunter's Ring (5351, -0.46 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; The 1 Ring (8350, -0.77 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.3 ranged_attack_power points (0.62 DPS) | yes | Bounty Hunter's Ring (5351, -0.13 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; The 1 Ring (8350, -0.46 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 26.5 ranged_attack_power points (1.59 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.51 DPS) [world]; Crescent Staff (6505, -0.51 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.8 ranged_attack_power points (10.75 DPS) | yes | Lil Timmy's Peashooter (13136, -1.76 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.33 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.57 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6255 Fishing Pole (JEFFTEST); 6478 Rat Stompers

### Band 40 (troll, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 105.6. Weights run: 0.6s. Verify run: 1.3s. 550 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.008, agility=2.000 ± 0.016, crit=0.832 ± 0.162 per rating point (14 rating = 1%, 11.643 per %), hit=0.218 ± 0.040 per rating point (10 rating = 1%, 2.175 per %), melee_haste=not significant (2.790 ± 4.937)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 26.0 ranged_attack_power points (1.52 DPS) | yes | Nightscape Headband (8176, +0.00 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -0.12 DPS) [crafted]; Skullsplitter Helm (1624, -0.23 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 22.0 ranged_attack_power points (1.28 DPS) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS, sim-verified) [quest]; Scout's Medallion (19537, -0.35 DPS) [rep]; Ghostshard Talisman (7731, -0.47 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 34.0 ranged_attack_power points (1.98 DPS) | yes | Nightscape Shoulders (8192, +0.00 DPS, sim-verified) [crafted]; Forest Tracker Epaulets (2278, -0.70 DPS) [world_drop]; Flintrock Shoulders (7755, -0.82 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Parachute Cloak (10518, +0.00 DPS, sim-verified) [crafted]; Imperial Cloak (6432, -0.23 DPS) [world_drop]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 38.0 ranged_attack_power points (2.22 DPS) | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.47 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.47 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.23 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (1.87 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS, sim-verified) [world_drop]; Dragonscale Gauntlets (8347, -0.49 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.70 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.75 DPS) | yes | Defiler's Chain Girdle (20152, +0.00 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20191, -0.35 DPS) [rep]; Scorpashi Sash (14652, -0.58 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 42.0 ranged_attack_power points (2.45 DPS) | yes | Triprunner Dungarees (9624, +0.00 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.82 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.82 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 26.0 ranged_attack_power points (1.52 DPS) | yes | Skulker's Leather Shoes (252531, +0.00 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.23 DPS) [world_drop]; Stalker's Mail Boots (252562, -0.23 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Mark of Kern (2262, -0.00 DPS) [dungeon]; Falcon's Hook (7552, -0.12 DPS) [world_drop]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| finger2 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Mark of Kern (2262, +0.00 DPS, sim-verified) [dungeon]; Falcon's Hook (7552, -0.12 DPS) [world_drop]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 0.0 ranged_attack_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 18.0 ranged_attack_power points (1.05 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS, sim-verified) [crafted]; Satyr's Rod (15962, -0.93 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (105.6 DPS) | yes | Shadowforge Bushmaster (9422, -0.46 DPS) [dungeon]; Monolithic Bow (9426, -0.85 DPS) [dungeon]; Bow of Searing Arrows (2825, -74.53 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Assault Band; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 550, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 232.6. Weights run: 0.7s. Verify run: 1.2s. 1534 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=1.302 ± 0.249 per rating point (14 rating = 1%, 18.230 per %), hit=0.315 ± 0.055 per rating point (10 rating = 1%, 3.147 per %), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Visor (239532) | Leonid Barthalomew the Revered [vendor] | 176.5 ranged_attack_power points (10.08 DPS) | yes | Warlord's Chain Greathelm (231568, -5.04 DPS) [vendor]; Warlord's Chain Helmet (16566, -5.16 DPS) [vendor]; Dawnstalker Headpiece (239540, -5.28 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 42.2 ranged_attack_power points (2.41 DPS) | yes | Amulet of the Darkmoon (19491, -0.32 DPS, sim-verified) [quest]; Imperial Jewel (11933, -0.58 DPS) [dungeon]; Pendant of Celerity (22340, -0.68 DPS) [dungeon] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | 129.4 ranged_attack_power points (7.39 DPS) | yes | Dawnstalker Spaulders (239542, -1.24 DPS, sim-verified) [vendor]; Darkspear Pauldrons (272105, -3.26 DPS) [vendor]; Darkspear Epaulets (272106, -3.26 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.0 ranged_attack_power points (2.85 DPS) | yes | Deathguard's Cloak (20068, -0.91 DPS, sim-verified) [rep]; Shifting Cloak (18511, -0.91 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.91 DPS) [dungeon] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (232.6 DPS) | yes | Dawnstalker Tunic (239543, -1.81 DPS) [vendor]; Dawn Armor (252483, -5.09 DPS) [crafted]; Tunic of Undead Slaying (23089, -16.85 DPS, sim-verified) [world] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | sim-verified (232.6 DPS) | yes | Dawnstalker Wristguards (239544, -0.69 DPS) [vendor]; Bracers of the Eclipse (18375, -2.53 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.43 DPS, sim-verified) [world] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | 132.2 ranged_attack_power points (7.55 DPS) | yes | Dawnstalker Handguards (239539, -1.45 DPS, sim-verified) [vendor]; Chromatic Gauntlets (19157, -4.00 DPS) [crafted]; General's Chain Gloves (16571, -4.11 DPS) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 108.2 ranged_attack_power points (6.18 DPS) | yes | Dawnstalker Girdle (239538, -1.35 DPS, sim-verified) [vendor]; Dense Timbermaw Belt (227807, -2.53 DPS) [vendor]; Molten Belt (19163, -2.98 DPS) [crafted] |
| legs | Dawnstalker Leggings (239533) | Leonid Barthalomew the Revered [vendor] | 174.5 ranged_attack_power points (9.96 DPS) | yes | Sentinel's Chain Leggings (237819, -3.71 DPS) [vendor]; Outrider's Chain Leggings (22673, -4.75 DPS) [rep]; Dawnstalker Legguards (239541, -5.45 DPS, sim-verified) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | 140.2 ranged_attack_power points (8.01 DPS) | yes | Dawnstalker Boots (239537, -3.67 DPS, sim-verified) [vendor]; General's Chain Sabatons (231564, -4.74 DPS) [vendor]; General's Chain Sabatons (16569, -4.86 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (232.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.23 DPS) [quest]; Wrath of Cenarius (21190, -5.08 DPS, sim-verified) [quest] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (232.6 DPS) | yes | Cutthroat's Signet (272408, -0.36 DPS) [vendor]; Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Wrath of Cenarius (21190, -3.78 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (232.6 DPS) | yes | Counterattack Lodestone (18537, -3.67 DPS) [dungeon]; Hand of Justice (11815, -3.78 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -4.76 DPS) [crafted] |
| trinket2 | Earthstrike (21180) | Champion's Battlegear [quest] | sim-verified (232.6 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Shard of the Fallen Star (21891, -5.10 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (232.6 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -5.72 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Core Marksman Rifle (18282) | Engineering [crafted] | sim-verified (232.6 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -8.86 DPS, sim-verified) [crafted] |

**New at 60:** head: Dawnstalker Visor; neck: Medallion of the Dawn; shoulder: Dawnstalker Pauldrons; back: Cape of the Black Baron; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Leggings; feet: Dawnstalker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Rune of the Guard Captain; trinket2: Earthstrike; main_hand: Legionite Glaive; ranged: Core Marksman Rifle

No-known-source sample (15 of 1534, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

