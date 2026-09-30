# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.5. Weights run: 0.9s. Verify run: 1.3s. 217 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.004, agility=2.567 ± 0.314, crit=0.505 ± 0.072 per rating point (14 rating = 1%, 7.077 per %), hit=0.151 ± 0.027 per rating point (10 rating = 1%, 1.512 per %), melee_haste=not significant (7.284 ± 3.386)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 20.5 ranged_attack_power points (1.23 DPS) | yes | Resilient Cloth Headband (211500, -1.05 DPS, sim-verified) [vendor] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Slime-encrusted Pads (6461, -0.81 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Cape of the Brotherhood (5193, +0.00 DPS, sim-verified) [dungeon]; Hide of Lupos (3018, -0.31 DPS) [world]; Bristlebark Cape (14571, -0.31 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 28.2 ranged_attack_power points (1.70 DPS) | yes | Brawler's Leather Armor (252490, -0.61 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.62 DPS) [crafted]; Dark Leather Tunic (2317, -0.77 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Bravo's Armbands (270015, -0.15 DPS) [quest]; Ratchet Wristwraps (274742, -0.31 DPS) [vendor] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Gloves of the Fang (10413, -0.08 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.31 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.31 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.08 DPS) | yes | Deviate Scale Belt (6468, -0.31 DPS) [crafted]; Dark Leather Belt (4249, -0.46 DPS) [crafted]; Dusty Belt (279897, -0.79 DPS, sim-verified) [quest] |
| legs | Leggings of the Fang (10410) (or Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 23.1 ranged_attack_power points (1.39 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.15 DPS) [world]; Brawler's Leather Pants (252500, -0.35 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.5 ranged_attack_power points (1.23 DPS) | yes | Footpads of the Fang (10411, -0.26 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.31 DPS) [dungeon]; Agile Boots (4788, -0.46 DPS) [vendor] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; The 1 Ring (8350, -0.77 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.3 ranged_attack_power points (0.62 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.46 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 26.5 ranged_attack_power points (1.59 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.51 DPS) [world]; Lupine Axe (1220, -0.67 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.8 ranged_attack_power points (10.75 DPS) | yes | Lil Timmy's Peashooter (13136, -1.62 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.33 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.57 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (dwarf, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 110.3. Weights run: 0.8s. Verify run: 1.3s. 591 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.008, agility=2.000 ± 0.016, crit=0.832 ± 0.162 per rating point (14 rating = 1%, 11.643 per %), hit=0.218 ± 0.040 per rating point (10 rating = 1%, 2.175 per %), melee_haste=not significant (2.790 ± 4.937)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 26.0 ranged_attack_power points (1.52 DPS) | yes | Guard's Chain Helm (250499, -0.12 DPS) [crafted]; Skullsplitter Helm (1624, -0.23 DPS) [world]; Nightscape Headband (8176, -0.24 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 22.0 ranged_attack_power points (1.28 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.32 DPS, sim-verified) [quest]; Ghostshard Talisman (7731, -0.47 DPS) [dungeon]; Erudite's Amulet (277204, -0.82 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 34.0 ranged_attack_power points (1.98 DPS) | yes | Forest Tracker Epaulets (2278, -0.70 DPS) [world_drop]; Nightscape Shoulders (8192, -0.72 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.82 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Imperial Cloak (6432, -0.23 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Parachute Cloak (10518, -0.28 DPS, sim-verified) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 38.0 ranged_attack_power points (2.22 DPS) | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.47 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.47 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Dusky Bracers (7378, -0.07 DPS, sim-verified) [crafted]; Imperial Leather Bracers (4061, -0.23 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (1.87 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS, sim-verified) [world_drop]; Dragonscale Gauntlets (8347, -0.49 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.70 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (1.75 DPS) | yes | Highlander's Chain Girdle (20090, -0.36 DPS, sim-verified) [rep]; Scorpashi Sash (14652, -0.58 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 42.0 ranged_attack_power points (2.45 DPS) | yes | Triprunner Dungarees (9624, -0.11 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.82 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.82 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 26.0 ranged_attack_power points (1.52 DPS) | yes | Dusky Boots (7390, -0.23 DPS) [crafted]; Worn Running Boots (9398, -0.23 DPS) [dungeon]; Imperial Leather Boots (6431, -0.28 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Assault Band (13095, -0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Assault Band (13095, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (110.3 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.32 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 18.0 ranged_attack_power points (1.05 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS, sim-verified) [crafted]; Satyr's Rod (15962, -0.93 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (20.40 DPS) | yes | Shadowforge Bushmaster (9422, -1.83 DPS) [dungeon]; Monolithic Bow (9426, -2.22 DPS) [dungeon]; The Silencer (13138, -2.68 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Mark of Kern; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 591, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 185.9. Weights run: 0.8s. Verify run: 1.9s. 1627 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=1.302 ± 0.249 per rating point (14 rating = 1%, 18.230 per %), hit=0.315 ± 0.055 per rating point (10 rating = 1%, 3.147 per %), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Field Marshal's Chain Greathelm (231562) | Captain Dirgehammer [vendor] | 88.2 ranged_attack_power points (5.04 DPS) | yes | Field Marshal's Chain Helm (231580, -0.11 DPS) [pvp]; Lieutenant Commander's Chain Helm (227066, -0.90 DPS) [pvp]; Lieutenant Commander's Chain Greathelm (227086, -2.07 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 42.2 ranged_attack_power points (2.41 DPS) | yes | Amulet of the Darkmoon (19491, -0.12 DPS, sim-verified) [quest]; Imperial Jewel (11933, -0.58 DPS) [dungeon]; Pendant of Celerity (22340, -0.68 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 72.2 ranged_attack_power points (4.12 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS, sim-verified) [vendor]; Highlander's Leather Shoulders (20059, -0.36 DPS) [rep]; Field Marshal's Chain Pauldrons (231557, -0.86 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.0 ranged_attack_power points (2.85 DPS) | yes | Cloak of the Honor Guard (20073, -0.81 DPS, sim-verified) [rep]; Shifting Cloak (18511, -0.91 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.91 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (185.9 DPS) | yes | Obsidian Mail Tunic (22191, -0.07 DPS) [crafted]; Field Marshal's Chain Armor (231563, -0.41 DPS) [vendor]; Tunic of Undead Slaying (23089, -9.80 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (185.9 DPS) | yes | Marshal's Chain Bracers (16461, -0.23 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.34 DPS) [rep]; Wristwraps of Undead Slaying (23093, -3.54 DPS, sim-verified) [world] |
| hands | Marshal's Chain Vices (231578) | Captain Dirgehammer [vendor] | 60.2 ranged_attack_power points (3.44 DPS) | yes | Marshal's Chain Grips (231560, -0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.80 DPS) [crafted]; Raider Gloves (272099, -1.06 DPS, sim-verified) [vendor] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 ranged_attack_power points (3.65 DPS) | yes | Highlander's Chain Girdle (20043, -0.56 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20045, -0.67 DPS) [rep]; Light Obsidian Belt (22195, -0.79 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 109.6 ranged_attack_power points (6.26 DPS) | yes | Sentinel's Leather Pants (237818, -1.09 DPS) [vendor]; Marshal's Chain Legplates (231558, -1.22 DPS) [vendor]; Marshal's Chain Legguards (231577, -1.33 DPS) [pvp] |
| feet | Marshal's Chain Sabatons (231561) | Captain Dirgehammer [vendor] | 57.1 ranged_attack_power points (3.26 DPS) | yes | Marshal's Chain Greaves (231579, +0.00 DPS, sim-verified) [vendor]; Marshal's Chain Boots (16462, -0.11 DPS) [vendor]; Scalegut Treaders (275618, -0.18 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (185.9 DPS) | yes | Cutthroat's Signet (272408, -1.14 DPS) [vendor]; Tarnished Elven Ring (18500, -1.19 DPS) [dungeon]; Naglering (11669, -6.02 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (185.9 DPS) | yes | Cutthroat's Signet (272408, -0.36 DPS) [vendor]; Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Naglering (11669, -4.61 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (185.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (185.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Devilsaur Eye (19991, -1.01 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (185.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -6.77 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (185.9 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -8.07 DPS, sim-verified) [crafted] |

**New at 60:** head: Field Marshal's Chain Greathelm; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Bracers of the Eclipse; hands: Marshal's Chain Vices; waist: Dense Timbermaw Belt; legs: Sentinel's Chain Leggings; feet: Marshal's Chain Sabatons; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1627, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.0. Weights run: 0.9s. Verify run: 1.3s. 213 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.004, agility=2.567 ± 0.314, crit=0.505 ± 0.072 per rating point (14 rating = 1%, 7.077 per %), hit=0.151 ± 0.027 per rating point (10 rating = 1%, 1.512 per %), melee_haste=not significant (7.284 ± 3.386)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 20.5 ranged_attack_power points (1.23 DPS) | yes | Resilient Cloth Headband (211500, -1.00 DPS, sim-verified) [vendor] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Cape of the Brotherhood (5193, -0.23 DPS, sim-verified) [dungeon]; Hide of Lupos (3018, -0.31 DPS) [world]; Bristlebark Cape (14571, -0.31 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 18.0 ranged_attack_power points (1.08 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.15 DPS) [crafted]; Prospector's Chestpiece (14562, -0.15 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 12.8 ranged_attack_power points (0.77 DPS) | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.31 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.31 DPS) [vendor] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Gloves of the Fang (10413, -0.08 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.31 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.31 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.08 DPS) | yes | Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.46 DPS) [world]; Dark Leather Belt (4249, -0.46 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 23.1 ranged_attack_power points (1.39 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.15 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.5 ranged_attack_power points (1.23 DPS) | yes | Footpads of the Fang (10411, -0.25 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.31 DPS) [dungeon]; Agile Boots (4788, -0.46 DPS) [vendor] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 15.4 ranged_attack_power points (0.93 DPS) | yes | Bounty Hunter's Ring (5351, -0.46 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; The 1 Ring (8350, -0.77 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.3 ranged_attack_power points (0.62 DPS) | yes | Bounty Hunter's Ring (5351, -0.13 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; The 1 Ring (8350, -0.46 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 26.5 ranged_attack_power points (1.59 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.51 DPS) [world]; Crescent Staff (6505, -0.51 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.8 ranged_attack_power points (10.75 DPS) | yes | Lil Timmy's Peashooter (13136, -0.74 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.33 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.57 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 213, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 40 (troll, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 110.5. Weights run: 0.8s. Verify run: 1.4s. 579 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.008, agility=2.000 ± 0.016, crit=0.832 ± 0.162 per rating point (14 rating = 1%, 11.643 per %), hit=0.218 ± 0.040 per rating point (10 rating = 1%, 2.175 per %), melee_haste=not significant (2.790 ± 4.937)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 26.0 ranged_attack_power points (1.52 DPS) | yes | Nightscape Headband (8176, +0.00 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -0.12 DPS) [crafted]; Skullsplitter Helm (1624, -0.23 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 22.0 ranged_attack_power points (1.28 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.32 DPS, sim-verified) [quest]; Ghostshard Talisman (7731, -0.47 DPS) [dungeon]; Erudite's Amulet (277204, -0.82 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 34.0 ranged_attack_power points (1.98 DPS) | yes | Forest Tracker Epaulets (2278, -0.70 DPS) [world_drop]; Nightscape Shoulders (8192, -0.72 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -0.82 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Imperial Cloak (6432, -0.23 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Parachute Cloak (10518, -0.27 DPS, sim-verified) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 38.0 ranged_attack_power points (2.22 DPS) | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.47 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.47 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Dusky Bracers (7378, -0.11 DPS, sim-verified) [crafted]; Imperial Leather Bracers (4061, -0.23 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (1.87 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS, sim-verified) [world_drop]; Dragonscale Gauntlets (8347, -0.49 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.70 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (1.75 DPS) | yes | Defiler's Chain Girdle (20152, -0.36 DPS, sim-verified) [rep]; Scorpashi Sash (14652, -0.58 DPS) [world_drop]; Deftkin Belt (16659, -0.58 DPS) [quest] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 42.0 ranged_attack_power points (2.45 DPS) | yes | Triprunner Dungarees (9624, -0.28 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.82 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.82 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 26.0 ranged_attack_power points (1.52 DPS) | yes | Dusky Boots (7390, -0.23 DPS) [crafted]; Worn Running Boots (9398, -0.23 DPS) [dungeon]; Imperial Leather Boots (6431, -0.27 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Assault Band (13095, -0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Assault Band (13095, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Ironspine's Eye (7686, -0.12 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (110.5 DPS) | yes | Manslayer (10570, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -1.32 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 18.0 ranged_attack_power points (1.05 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS, sim-verified) [crafted]; Satyr's Rod (15962, -0.93 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (20.40 DPS) | yes | Shadowforge Bushmaster (9422, -1.83 DPS) [dungeon]; The Silencer (13138, -2.12 DPS, sim-verified) [world_drop]; Monolithic Bow (9426, -2.22 DPS) [dungeon] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Mark of Kern; main_hand: Vanquisher's Sword; off_hand: Jhordy's Misplaced Screwdriver; ranged: Bow of Searing Arrows

No-known-source sample (15 of 579, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 187.8. Weights run: 0.8s. Verify run: 1.5s. 1616 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.005, agility=2.000 ± 0.009, crit=1.302 ± 0.249 per rating point (14 rating = 1%, 18.230 per %), hit=0.315 ± 0.055 per rating point (10 rating = 1%, 3.147 per %), melee_haste=not significant (4.590 ± 6.727)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warlord's Chain Greathelm (231568) | Lady Palanseer [vendor] | 88.2 ranged_attack_power points (5.04 DPS) | yes | Warlord's Chain Helm (231571, +0.00 DPS, sim-verified) [vendor]; Warlord's Chain Helmet (16566, -0.11 DPS) [vendor]; Champion's Chain Greathelm (227080, -0.33 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 42.2 ranged_attack_power points (2.41 DPS) | yes | Amulet of the Darkmoon (19491, -0.16 DPS, sim-verified) [quest]; Imperial Jewel (11933, -0.58 DPS) [dungeon]; Pendant of Celerity (22340, -0.68 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 72.2 ranged_attack_power points (4.12 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS, sim-verified) [vendor]; Defiler's Leather Shoulders (20194, -0.36 DPS) [rep]; Warlord's Chain Pauldrons (231565, -0.86 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.0 ranged_attack_power points (2.85 DPS) | yes | Deathguard's Cloak (20068, -0.73 DPS, sim-verified) [rep]; Shifting Cloak (18511, -0.91 DPS) [crafted]; Shadow Prowler's Cloak (22269, -0.91 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (187.8 DPS) | yes | Obsidian Mail Tunic (22191, -0.07 DPS) [crafted]; Warlord's Chain Armor (231566, -0.41 DPS) [vendor]; Tunic of Undead Slaying (23089, -9.27 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (187.8 DPS) | yes | General's Chain Wristguards (16570, -0.23 DPS) [pvp]; Forest Stalker's Bracers (19587, -0.34 DPS) [rep]; Wristwraps of Undead Slaying (23093, -3.45 DPS, sim-verified) [world] |
| hands | General's Chain Vices (231575) | Lady Palanseer [vendor] | 60.2 ranged_attack_power points (3.44 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, -0.31 DPS, sim-verified) [vendor]; Raider Gloves (272099, -0.36 DPS) [vendor] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 ranged_attack_power points (3.65 DPS) | yes | Defiler's Chain Girdle (20150, -0.59 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20190, -0.67 DPS) [rep]; Light Obsidian Belt (22195, -0.79 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 109.6 ranged_attack_power points (6.26 DPS) | yes | Sentinel's Leather Pants (237818, -1.09 DPS) [vendor]; General's Chain Legplates (231567, -1.22 DPS) [vendor]; Outrider's Chain Leggings (22673, -1.72 DPS, sim-verified) [rep] |
| feet | General's Chain Greaves (231570) | Lady Palanseer [vendor] | 55.1 ranged_attack_power points (3.15 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; Beastmaster's Boots (22061, -0.41 DPS) [quest]; Scalegut Treaders (275618, -2.47 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (187.8 DPS) | yes | Cutthroat's Signet (272408, -1.14 DPS) [vendor]; Tarnished Elven Ring (18500, -1.19 DPS) [dungeon]; Naglering (11669, -5.32 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (187.8 DPS) | yes | Cutthroat's Signet (272408, -0.36 DPS) [vendor]; Tarnished Elven Ring (18500, -0.40 DPS) [dungeon]; Naglering (11669, -4.04 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (187.8 DPS) | yes | Counterattack Lodestone (18537, -3.67 DPS) [dungeon]; Hand of Justice (11815, -3.78 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -4.76 DPS) [crafted] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (187.8 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Second Wind (11819, -0.87 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (187.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -6.02 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (187.8 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -6.82 DPS, sim-verified) [crafted] |

**New at 60:** head: Warlord's Chain Greathelm; neck: Medallion of the Dawn; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Bracers of the Eclipse; hands: General's Chain Vices; waist: Dense Timbermaw Belt; legs: Sentinel's Chain Leggings; feet: General's Chain Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Rune of the Guard Captain; trinket2: Burst of Knowledge; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1616, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

