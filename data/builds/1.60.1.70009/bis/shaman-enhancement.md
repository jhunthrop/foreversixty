# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 36.4. Weights run: 1.8s. Verify run: 0.9s. 225 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.201 ± 0.003, crit=0.283 ± 0.005 per rating point (14 rating = 1%, 3.963 per %), hit=0.215 ± 0.002 per rating point (10 rating = 1%, 2.149 per %), melee_haste=2.477 ± 0.031

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.55 DPS) | yes | Brawler's Leather Hood (252504, -0.42 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 1.2 attack_power points (0.04 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 1.0 attack_power points (0.03 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Defender's Leather Armor (252434) | Leatherworking [crafted] | 14.6 attack_power points (0.50 DPS) | yes | Totemic Leather Armor (252435, -0.02 DPS) [crafted]; Armor of the Fang (6473, -0.09 DPS) [dungeon]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 4.8 attack_power points (0.17 DPS) | yes | Bristlebark Bindings (14569, -0.01 DPS) [world_drop]; Light Leather Bracers (7281, -0.09 DPS) [crafted]; Cloudy Gustwoven Bracers (276999, -0.10 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.0 attack_power points (0.48 DPS) | yes | Blackened Defias Gloves (10401, -0.07 DPS) [dungeon]; Foreman's Gloves (2167, -0.14 DPS) [world]; Gloves of the Fang (10413, -0.17 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Ruffian Belt (5975, -0.21 DPS, sim-verified) [world]; Support Girdle (1215, -0.28 DPS) [world]; Brawler's Leather Belt (252428, -0.32 DPS) [crafted] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 18.8 attack_power points (0.65 DPS) | yes | Totemic Leather Pants (252446, +0.00 DPS, sim-verified) [crafted]; Deepgrave Trousers (279900, -0.14 DPS) [quest]; Brawler's Leather Pants (252500, -0.17 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.0 attack_power points (0.38 DPS) | yes | Defender's Leather Boots (252441, -0.03 DPS) [crafted]; Totemic Leather Boots (252442, -0.03 DPS) [crafted]; Forest Leather Boots (3057, -0.10 DPS) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.8 attack_power points (0.30 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop]; Signet of the Zhevra (285330, -0.26 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.17 DPS, sim-verified) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop]; Signet of the Zhevra (285330, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | sim-verified (36.4 DPS) | yes | Living Root (6631, -0.06 DPS) [dungeon]; Night Reaver (1318, -0.35 DPS) [dungeon]; The Axe of Severing (23171, -3.02 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Defender's Leather Armor; wrist: Bravo's Armbands; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Defender's Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 92.8. Weights run: 1.9s. Verify run: 1.1s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.367 ± 0.006, crit=0.515 ± 0.009 per rating point (14 rating = 1%, 7.205 per %), hit=0.303 ± 0.003 per rating point (10 rating = 1%, 3.026 per %), melee_haste=3.293 ± 0.044

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (0.85 DPS) | yes | Azure Gustwoven Hood (277050, -0.21 DPS) [crafted]; Defender's Leather Hood (252447, -0.28 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.49 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Sentinel's Medallion (19541, -0.39 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.0 attack_power points (0.50 DPS) | yes | Barbaric Shoulders (5964, -0.08 DPS) [crafted]; Bristlebark Amice (14573, -0.21 DPS) [world_drop]; Watchman Pauldrons (7727, -0.28 DPS) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (92.8 DPS) | yes | Hawkeye's Cloak (14593, -0.03 DPS) [world_drop]; Slayer's Cape (14752, -0.05 DPS) [world_drop]; Wolfmaster Cape (6314, -0.94 DPS, sim-verified) [dungeon] |
| chest | Thick Murloc Armor (5782) (or Nightwalker Armor (2234)) | Leatherworking [crafted] | 18.0 attack_power points (0.64 DPS) | yes | Nightwalker Armor (2234, +0.00 DPS) [world]; Raptorbane Armor (3566, -0.07 DPS) [quest]; Totemic Leather Tunic (252451, -0.07 DPS) [crafted] |
| wrist | Bands of Serra'kis (6902) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.0 attack_power points (0.42 DPS) | yes | Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Cultist's Armguards (270032, -0.07 DPS) [quest]; Technician's Bracers (270042, -0.07 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.2 attack_power points (0.57 DPS) | yes | Heavy Earthen Gloves (7359, -0.01 DPS) [crafted]; Brawler Gloves (720, -0.01 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.07 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (0.85 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.13 DPS) [crafted]; Blackened Defias Belt (10403, -0.21 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Totemic Leather Pants (252446, -0.28 DPS) [crafted]; Totemic Leather Leggings (252458, -0.28 DPS) [crafted]; Defender's Leather Pants (252445, -1.75 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.42 DPS) | yes | Brawler's Leather Boots (252439, -0.01 DPS) [crafted]; Draftsman Boots (6668, -0.07 DPS) [quest]; Defender's Leather Boots (252441, -0.07 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.60 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.20 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.25 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.2 attack_power points (0.50 DPS) | yes | Tiger Band (6749, -0.08 DPS) [quest]; Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.16 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.61 DPS) [vendor]; Viscous Hammer (13045, -24.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Thick Murloc Armor; wrist: Bands of Serra'kis; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 115.5. Weights run: 2.2s. Verify run: 1.1s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.425 ± 0.008, crit=0.599 ± 0.011 per rating point (14 rating = 1%, 8.384 per %), hit=0.337 ± 0.003 per rating point (10 rating = 1%, 3.370 per %), melee_haste=3.459 ± 0.123

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 34.4 attack_power points (1.42 DPS) | yes | White Bandit Mask (10008, -0.32 DPS) [crafted]; Tusken Helm (6686, -0.35 DPS) [dungeon]; Hard Gold Coif (250537, -1.61 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.82 DPS) | yes | Ghostshard Talisman (7731, -0.25 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.42 DPS) [world_drop]; River Pride Choker (13087, -0.49 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (0.91 DPS) | yes | Imperial Leather Spaulders (4737, -0.16 DPS) [dungeon]; Wrangling Spaulders (15698, -0.21 DPS) [quest]; Sunburn Spaulders (274751, -0.22 DPS) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.6 attack_power points (0.60 DPS) | yes | Wolfmaster Cape (6314, -0.19 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Dark Hooded Cape (5257, -3.09 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.0 attack_power points (1.28 DPS) | yes | Avenger's Armor (1488, -0.04 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Golden Scale Cuirass (3845, -0.12 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Pugilist Bracers (4438, -0.16 DPS) [dungeon]; Ravager's Armguards (14770, -0.18 DPS) [world_drop]; Yorgen Bracers (13012, -0.28 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (115.5 DPS) | yes | Scarlet Gauntlets (10331, -0.01 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.10 DPS) [world_drop]; Gauntlets of Divinity (7724, -1.28 DPS, sim-verified) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.24 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Highlander's Chain Girdle (20090, -0.25 DPS) [rep]; Scarlet Belt (10329, -0.25 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.73 DPS) | yes | Firemane Leggings (13129, -0.16 DPS) [world_drop]; Orcish War Leggings (7929, -0.33 DPS) [crafted]; Legguards of the Vault (9396, -0.57 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 25.0 attack_power points (1.03 DPS) | yes | Skirmisher's Mail Boots (252564, -0.12 DPS) [crafted]; Ironheel Boots (4653, -0.21 DPS) [quest]; Blackforge Greaves (6423, -1.21 DPS, sim-verified) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Protector's Band (19515, -0.02 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.82 DPS) | yes | Protector's Band (19515, -0.02 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pendulum of Doom (9425) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; The Jackhammer (9423, -0.97 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Pendulum of Doom

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 158.3. Weights run: 2.2s. Verify run: 1.2s. 762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.496 ± 0.010, crit=0.697 ± 0.014 per rating point (14 rating = 1%, 9.752 per %), hit=0.414 ± 0.004 per rating point (10 rating = 1%, 4.142 per %), melee_haste=4.035 ± 0.161

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.02 DPS) | yes | Raging Berserker's Helm (7719, -0.52 DPS) [dungeon]; Knight-Lieutenant's Mail Helmet (223075, -0.68 DPS) [vendor]; Bloomsprout Headpiece (17767, -2.57 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.84 DPS) | yes | Skibi's Pendant (13089, -0.15 DPS) [world_drop]; Ghostshard Talisman (7731, -0.25 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.42 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Mail Epaulets (223073) | Captain Dirgehammer [vendor] | 27.8 attack_power points (1.17 DPS) | yes | Prowler's Leather Shoulder (252534, -0.05 DPS) [crafted]; Failed Flying Experiment (9647, -0.12 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.18 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.9 attack_power points (0.80 DPS) | yes | Sergeant Major's Cape (16336, -0.17 DPS) [pvp]; Dark Hooded Cape (5257, -0.25 DPS) [world]; Bloodlust Cape (14801, -1.86 DPS, sim-verified) [world_drop] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | sim-verified (158.3 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Warbear Harness (15064, -0.12 DPS) [crafted]; Mixologist's Tunic (12793, -3.27 DPS, sim-verified) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.18 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.28 DPS) [crafted]; Branded Leather Bracers (19508, -0.34 DPS) [dungeon] |
| hands | Maddening Gauntlets (11867) | Ogre Head On A Stick = Party [quest] | 32.5 attack_power points (1.37 DPS) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS, sim-verified) [crafted]; Gauntlets of Divinity (7724, -0.02 DPS) [dungeon]; Raider Gloves (272100, -0.07 DPS) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.94 DPS) | yes | Belt of the Gladiator (13134, -0.42 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.51 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.63 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.77 DPS) | yes | Firemane Leggings (13129, -0.17 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -0.19 DPS) [quest]; Orcish War Leggings (7929, -0.34 DPS) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 31.5 attack_power points (1.33 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -0.23 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.5 attack_power points (1.03 DPS) | yes | Mark of Kern (2262, -0.19 DPS) [dungeon]; Assault Band (13095, -0.19 DPS) [world_drop]; Thunderbrow Ring (13097, -0.29 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 24.1 attack_power points (1.02 DPS) | yes | Assault Band (13095, -0.17 DPS) [world_drop]; Thunderbrow Ring (13097, -0.28 DPS) [world_drop]; Mark of Kern (2262, -3.05 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted] |
| main_hand | Ragehammer (10626) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Glowing Brightwood Staff (812, -4.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Mail Epaulets; back: Blackveil Cape; chest: Knight's Mail Armor; wrist: Arena Bands; hands: Maddening Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Ragehammer

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 231.0. Weights run: 2.1s. Verify run: 1.1s. 1751 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.580 ± 0.011, crit=0.814 ± 0.016 per rating point (14 rating = 1%, 11.397 per %), hit=0.563 ± 0.005 per rating point (10 rating = 1%, 5.627 per %), melee_haste=4.709 ± 0.187

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.98 DPS) | yes | Warbear Helm (252485, -0.16 DPS) [crafted]; Face of The Five Thunders (227021, -0.19 DPS) [vendor]; Blue Suede Hat (252482, -0.19 DPS) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 35.4 attack_power points (1.46 DPS) | yes | Amulet of the Darkmoon (19491, -0.18 DPS) [quest]; Will of the Martyr (17044, -0.22 DPS) [quest]; Imperial Jewel (11933, -1.83 DPS, sim-verified) [dungeon] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 40.4 attack_power points (1.67 DPS) | yes | Black Dragonscale Shoulders (15051, -0.02 DPS) [crafted]; Highlander's Lizardhide Shoulders (20060, -0.14 DPS) [rep]; Golden Mantle of the Dawn (19058, -0.33 DPS) [crafted] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 36.9 attack_power points (1.52 DPS) | yes | Shroud of Domination (22337, -0.12 DPS) [dungeon]; Howler's Furs (272414, -0.13 DPS) [vendor]; Cape of the Black Baron (13340, -0.34 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (231.0 DPS) | yes | Obsidian Mail Tunic (22191, -0.02 DPS) [crafted]; Cadaverous Armor (14637, -0.30 DPS) [dungeon]; Tunic of Undead Slaying (23089, -12.16 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | sim-verified (231.0 DPS) | yes | Forest Stalker's Bracers (19587, -0.20 DPS) [rep]; Tranquil Wristguards (279256, -0.25 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -7.66 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 50.1 attack_power points (2.07 DPS) | yes | Studded Timbermaw Brawlers (227809, -0.18 DPS) [vendor]; Cadaverous Gloves (14640, -0.25 DPS) [dungeon]; Skul's Fingerbone Claws (13395, -0.42 DPS) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.64 DPS) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS, sim-verified) [vendor]; Might of the Timbermaw (19044, -0.57 DPS) [crafted]; Windseeker's Belt (272410, -0.66 DPS) [vendor] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 66.4 attack_power points (2.74 DPS) | yes | Devilsaur Leggings (15062, -0.37 DPS) [crafted]; Black Dragonscale Leggings (15052, -0.51 DPS) [crafted]; Cadaverous Leggings (14638, -0.60 DPS) [dungeon] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 42.6 attack_power points (1.76 DPS) | yes | Pads of the Dread Wolf (13210, -0.11 DPS) [dungeon]; Drudge Boots (21532, -0.28 DPS) [quest]; Boots of Ferocity (22472, -0.37 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (231.0 DPS) | yes | Band of the Ogre King (18522, -0.21 DPS) [dungeon]; Blackstone Ring (17713, -0.30 DPS) [dungeon]; Naglering (11669, -9.59 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (231.0 DPS) | yes | Band of the Ogre King (18522, -0.10 DPS) [dungeon]; Blackstone Ring (17713, -0.20 DPS) [dungeon]; Naglering (11669, -4.11 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (231.0 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.62 DPS) [crafted]; Blackhand's Breadth (13965, -0.76 DPS) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (231.0 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Burst of Knowledge (11832, -4.49 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (231.0 DPS) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Seeping Willow (12969, -15.14 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** neck: Medallion of the Dawn; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Timbermaw Brawlers; waist: Dense Timbermaw Belt; legs: Warbear Woolies; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Hand of Justice; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 0000000000000000-253100000000000000-0000000000000000)

Set DPS (verified): 38.0. Weights run: 1.8s. Verify run: 0.8s. 205 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.201 ± 0.003, crit=0.283 ± 0.005 per rating point (14 rating = 1%, 3.963 per %), hit=0.215 ± 0.002 per rating point (10 rating = 1%, 2.149 per %), melee_haste=2.477 ± 0.031

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.55 DPS) | yes | Brawler's Leather Hood (252504, -0.53 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 1.2 attack_power points (0.04 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 1.0 attack_power points (0.03 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Defender's Leather Armor (252434) | Leatherworking [crafted] | 14.6 attack_power points (0.50 DPS) | yes | Totemic Leather Armor (252435, -0.02 DPS) [crafted]; Armor of the Fang (6473, -0.09 DPS) [dungeon]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 4.6 attack_power points (0.16 DPS) | yes | Light Leather Bracers (7281, -0.08 DPS) [crafted]; Cloudy Gustwoven Bracers (276999, -0.09 DPS) [crafted]; Azure Gustwoven Bracers (277023, -0.09 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.0 attack_power points (0.48 DPS) | yes | Blackened Defias Gloves (10401, -0.07 DPS) [dungeon]; Foreman's Gloves (2167, -0.14 DPS) [world]; Gloves of the Fang (10413, -0.17 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.62 DPS) | yes | Ruffian Belt (5975, -0.23 DPS, sim-verified) [world]; Support Girdle (1215, -0.28 DPS) [world]; Brawler's Leather Belt (252428, -0.32 DPS) [crafted] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 18.8 attack_power points (0.65 DPS) | yes | Totemic Leather Pants (252446, +0.00 DPS, sim-verified) [crafted]; Deepgrave Trousers (279900, -0.14 DPS) [quest]; Brawler's Leather Pants (252500, -0.17 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.0 attack_power points (0.38 DPS) | yes | Defender's Leather Boots (252441, -0.03 DPS) [crafted]; Totemic Leather Boots (252442, -0.03 DPS) [crafted]; Forest Leather Boots (3057, -0.10 DPS) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.8 attack_power points (0.30 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.20 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | The Axe of Severing (23171) | Shadowfang Keep: Sever [dungeon] | 322.0 attack_power points (11.11 DPS) | yes | Forsaken Greataxe (251533, -0.66 DPS) [quest]; Smite's Mighty Hammer (7230, -0.80 DPS) [dungeon]; Hammerbone (270018, -4.84 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Defender's Leather Armor; wrist: Bristlebark Bindings; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Defender's Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: The Axe of Severing

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 0000000000000000-253130030004000000-0000000000000000)

Set DPS (verified): 84.9. Weights run: 1.9s. Verify run: 1.0s. 349 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.367 ± 0.006, crit=0.515 ± 0.009 per rating point (14 rating = 1%, 7.205 per %), hit=0.303 ± 0.003 per rating point (10 rating = 1%, 3.026 per %), melee_haste=3.293 ± 0.044

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (0.85 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.21 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.21 DPS) [crafted]; Defender's Leather Hood (252447, -0.28 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.49 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.21 DPS) [world_drop]; Scout's Medallion (19537, -0.39 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 14.0 attack_power points (0.50 DPS) | yes | Barbaric Shoulders (5964, -0.08 DPS) [crafted]; Bristlebark Amice (14573, -0.21 DPS) [world_drop]; Watchman Pauldrons (7727, -0.28 DPS) [dungeon] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.35 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.05 DPS) [world_drop]; Slayer's Cape (14752, -0.07 DPS) [world_drop] |
| chest | Thick Murloc Armor (5782) (or Nightwalker Armor (2234)) | Leatherworking [crafted] | 18.0 attack_power points (0.64 DPS) | yes | Nightwalker Armor (2234, +0.00 DPS) [world]; Totemic Leather Tunic (252451, -0.07 DPS) [crafted]; Defender's Leather Armor (252434, -0.10 DPS) [crafted] |
| wrist | Bands of Serra'kis (6902) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.0 attack_power points (0.42 DPS) | yes | Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Cultist's Armguards (270032, -0.07 DPS) [quest]; Technician's Bracers (270042, -0.07 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.2 attack_power points (0.57 DPS) | yes | Heavy Earthen Gloves (7359, -0.01 DPS) [crafted]; Brawler Gloves (720, -0.01 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.07 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (0.85 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.13 DPS) [crafted]; Blackened Defias Belt (10403, -0.21 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (0.92 DPS) | yes | Totemic Leather Pants (252446, -0.28 DPS) [crafted]; Totemic Leather Leggings (252458, -0.28 DPS) [crafted]; Defender's Leather Pants (252445, -1.61 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 11.8 attack_power points (0.42 DPS) | yes | Draftsman Boots (6668, -0.06 DPS) [quest]; Defender's Leather Boots (252441, -0.06 DPS) [crafted]; Totemic Leather Boots (252442, -0.06 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.60 DPS) | yes | Tiger Band (6749, -0.18 DPS) [quest]; Ironspine's Eye (7686, -0.20 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.25 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.2 attack_power points (0.50 DPS) | yes | Tiger Band (6749, -0.08 DPS) [quest]; Ironspine's Eye (7686, -0.10 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (84.9 DPS) | yes | Corpsemaker (6687, -0.16 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.61 DPS) [vendor]; Viscous Hammer (13045, -22.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wildhunter Cloak; chest: Thick Murloc Armor; wrist: Bands of Serra'kis; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 0000000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 125.0. Weights run: 2.2s. Verify run: 1.0s. 555 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.425 ± 0.008, crit=0.599 ± 0.011 per rating point (14 rating = 1%, 8.384 per %), hit=0.337 ± 0.003 per rating point (10 rating = 1%, 3.370 per %), melee_haste=3.459 ± 0.123

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 34.4 attack_power points (1.42 DPS) | yes | White Bandit Mask (10008, -0.32 DPS) [crafted]; Tusken Helm (6686, -0.35 DPS) [dungeon]; Hard Gold Coif (250537, -1.74 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.82 DPS) | yes | Ghostshard Talisman (7731, -0.25 DPS) [dungeon]; Ethereal Talisman (4430, -0.34 DPS) [quest]; Kaleidoscope Chain (13084, -0.42 DPS) [world_drop] |
| shoulder | Imperial Leather Spaulders (4737) | Uldaman: Ancient Treasure [dungeon] | sim-verified (125.0 DPS) | yes | Wrangling Spaulders (15698, -0.05 DPS) [quest]; Sunburn Spaulders (274751, -0.05 DPS) [vendor]; Hard Gold Pauldrons (250539, -1.27 DPS, sim-verified) [crafted] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 14.6 attack_power points (0.60 DPS) | yes | Dark Hooded Cape (5257, -0.09 DPS) [world]; Wolfmaster Cape (6314, -0.19 DPS) [dungeon]; Wildhunter Cloak (16658, -0.19 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.0 attack_power points (1.28 DPS) | yes | Avenger's Armor (1488, -0.04 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Golden Scale Cuirass (3845, -0.12 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Pugilist Bracers (4438, -0.16 DPS) [dungeon] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.32 DPS) | yes | Gloves of Holy Might (867, -0.15 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.15 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.25 DPS) [world_drop] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.24 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Tharg's Shoelace (9705, -0.16 DPS) [quest]; Defiler's Chain Girdle (20152, -0.25 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.73 DPS) | yes | Firemane Leggings (13129, -0.16 DPS) [world_drop]; Orcish War Leggings (7929, -0.33 DPS) [crafted]; Legguards of the Vault (9396, -0.57 DPS) [dungeon] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 25.0 attack_power points (1.03 DPS) | yes | Skirmisher's Mail Boots (252564, -0.12 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.26 DPS) [crafted]; Blackforge Greaves (6423, -1.11 DPS, sim-verified) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Legionnaire's Band (19512, -0.02 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.82 DPS) | yes | Legionnaire's Band (19512, -0.02 DPS) [rep]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pendulum of Doom (9425) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Fiery War Axe (870, -5.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Imperial Leather Spaulders; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Prowler's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Pendulum of Doom

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 5500000000000000-253130030005102051-0000000000000000)

Set DPS (verified): 159.7. Weights run: 2.2s. Verify run: 1.1s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.496 ± 0.010, crit=0.697 ± 0.014 per rating point (14 rating = 1%, 9.752 per %), hit=0.414 ± 0.004 per rating point (10 rating = 1%, 4.142 per %), melee_haste=4.035 ± 0.161

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (2.02 DPS) | yes | Bloomsprout Headpiece (17767, -0.51 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.52 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -0.68 DPS) [vendor] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.84 DPS) | yes | Woven Ivy Necklace (19159, -0.15 DPS) [quest]; Skibi's Pendant (13089, -0.15 DPS) [world_drop]; Ghostshard Talisman (7731, -0.25 DPS) [dungeon] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 26.5 attack_power points (1.12 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.13 DPS) [crafted]; Failed Flying Experiment (9647, -1.62 DPS, sim-verified) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.9 attack_power points (0.80 DPS) | yes | First Sergeant's Cloak (16340, -0.17 DPS) [pvp]; Dark Hooded Cape (5257, -0.25 DPS) [world]; Bloodlust Cape (14801, -1.58 DPS, sim-verified) [world_drop] |
| chest | Stone Guard's Mail Armor (220826) | Lady Palanseer [vendor] | sim-verified (159.7 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Warbear Harness (15064, -0.12 DPS) [crafted]; Mixologist's Tunic (12793, -2.70 DPS, sim-verified) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.18 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | 32.5 attack_power points (1.37 DPS) | yes | Raider Gloves (272100, -0.07 DPS) [vendor]; Gloves of Holy Might (867, -0.11 DPS) [world_drop]; Gauntlets of Divinity (7724, -2.44 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (1.94 DPS) | yes | Belt of the Gladiator (13134, -0.42 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.51 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.63 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.77 DPS) | yes | Firemane Leggings (13129, -0.17 DPS) [world_drop]; Orcish War Leggings (7929, -0.34 DPS) [crafted]; Serpentskin Leggings (8262, -0.34 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 31.5 attack_power points (1.33 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Skirmisher's Mail Sabatons (252578, -0.23 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.5 attack_power points (1.03 DPS) | yes | White Bone Band (11862, -0.02 DPS) [quest]; Mark of Kern (2262, -0.19 DPS) [dungeon]; Assault Band (13095, -0.19 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 24.1 attack_power points (1.02 DPS) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; White Bone Band (11862, -2.93 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -2.87 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.55 DPS) [crafted]; Darkspear Raider's Reaper (272080, -0.96 DPS) [vendor]; Fiery War Axe (870, -6.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Stone Guard's Mail Armor; wrist: Arena Bands; hands: Prowler's Leather Gauntlets; waist: Girdle of Beastial Fury; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 5533220000000000-253130030005102051-0000000000000000)

Set DPS (verified): 213.6. Weights run: 2.1s. Verify run: 1.0s. 1672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.580 ± 0.011, crit=0.814 ± 0.016 per rating point (14 rating = 1%, 11.397 per %), hit=0.563 ± 0.005 per rating point (10 rating = 1%, 5.627 per %), melee_haste=4.709 ± 0.187

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.0 attack_power points (1.98 DPS) | yes | Champion's Mail Headguard (227155, +0.00 DPS) [pvp]; Warbear Helm (252485, -0.16 DPS) [crafted]; Face of The Five Thunders (227021, -0.19 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 35.4 attack_power points (1.46 DPS) | yes | Amulet of the Darkmoon (19491, -0.18 DPS) [quest]; Will of the Martyr (17044, -0.22 DPS) [quest]; Imperial Jewel (11933, -2.35 DPS, sim-verified) [dungeon] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 40.4 attack_power points (1.67 DPS) | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Black Dragonscale Shoulders (15051, -0.02 DPS) [crafted] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 36.9 attack_power points (1.52 DPS) | yes | Shroud of Domination (22337, -0.12 DPS) [dungeon]; Howler's Furs (272414, -0.13 DPS) [vendor]; Cape of the Black Baron (13340, -0.34 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (213.6 DPS) | yes | Obsidian Mail Tunic (22191, -0.02 DPS) [crafted]; Cadaverous Armor (14637, -0.30 DPS) [dungeon]; Tunic of Undead Slaying (23089, -11.74 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | sim-verified (213.6 DPS) | yes | Forest Stalker's Bracers (19587, -0.20 DPS) [rep]; Tranquil Wristguards (279256, -0.25 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -6.35 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 50.1 attack_power points (2.07 DPS) | yes | General's Mail Vices (231655, -0.05 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.18 DPS) [vendor]; Cadaverous Gloves (14640, -0.25 DPS) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (2.64 DPS) | yes | Ferocity of the Timbermaw (227805, -0.09 DPS) [vendor]; Might of the Timbermaw (19044, -0.57 DPS) [crafted]; Windseeker's Belt (272410, -0.66 DPS) [vendor] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 66.4 attack_power points (2.74 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; Legionnaire's Mail Legguards (227156, -0.22 DPS) [pvp]; Devilsaur Leggings (15062, -0.37 DPS) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 42.6 attack_power points (1.76 DPS) | yes | Pads of the Dread Wolf (13210, -0.11 DPS) [dungeon]; General's Mail Greaves (231656, -0.12 DPS) [vendor]; Drudge Boots (21532, -0.28 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (213.6 DPS) | yes | Band of the Ogre King (18522, -0.21 DPS) [dungeon]; Blackstone Ring (17713, -0.30 DPS) [dungeon]; Naglering (11669, -3.12 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (213.6 DPS) | yes | Band of the Ogre King (18522, -0.10 DPS) [dungeon]; Blackstone Ring (17713, -0.20 DPS) [dungeon]; Naglering (11669, -4.02 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (213.6 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (213.6 DPS) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (213.6 DPS) | yes | Gravestone War Axe (13983, +0.00 DPS) [dungeon]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** neck: Medallion of the Dawn; shoulder: Defiler's Leather Shoulders; back: Deathguard's Cloak; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Timbermaw Brawlers; waist: Dense Timbermaw Belt; legs: Warbear Woolies; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Blackhand's Breadth; trinket2: Burst of Knowledge; main_hand: The Unstoppable Force

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

