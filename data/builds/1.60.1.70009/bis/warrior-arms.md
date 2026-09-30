# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 29.6. Weights run: 1.2s. Verify run: 1.2s. 263 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000), hit=0.529 ± 0.020, melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.18 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 attack_power points (0.26 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Catacomb Cloak (279899, -0.10 DPS) [quest]; Grave Shroud (279865, -0.10 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.18 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.27 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Cryptwalker Bracers (280095, -0.10 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted]; Burnished Bracers (3211, -0.26 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 attack_power points (0.70 DPS) | yes | Gold-flecked Gloves (5195, -0.10 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -0.18 DPS) [quest]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.18 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Support Girdle (1215, -0.32 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.9 attack_power points (0.96 DPS) | yes | Defender's Leather Pants (252445, -0.18 DPS) [crafted]; Totemic Leather Pants (252446, -0.18 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.18 DPS, sim-verified) [crafted] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | sim-verified (29.6 DPS) | yes | Smite's Mighty Hammer (7230, -0.20 DPS) [dungeon]; Living Root (6631, -0.28 DPS) [dungeon]; The Axe of Severing (23171, -13.30 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Dwarven Fishing Pole (3567) (or Cracked Blacksmith Hammer (285279)) | Murloc Poachers [quest] | 4.2 attack_power points (0.18 DPS) | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Fine Longbow (11304, -0.01 DPS) [vendor]; Daryl's Hunting Rifle (2904, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Veteran's Boots; finger1: Demon Band; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Forsaken Greataxe; ranged: Dwarven Fishing Pole

No-known-source sample (15 of 263, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings

### Band 30 (human, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 70.5. Weights run: 1.2s. Verify run: 1.2s. 446 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.171 ± 0.023, hit=0.743 ± 0.031, melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.26 DPS) | yes | Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Veteran's Chain Helm (250498, -0.17 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.68 DPS) | yes | River Pride Choker (13087, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.46 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.68 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.49 DPS) | yes | Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.16 DPS, sim-verified) [pvp] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (70.5 DPS) | yes | Barbaric Iron Breastplate (7914, -0.19 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Avenger's Armor (1488, -1.83 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 attack_power points (0.78 DPS) | yes | Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Cultist's Armguards (270032, -0.29 DPS) [quest]; Yorgen Bracers (13012, -0.44 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.07 DPS) | yes | Heavy Earthen Gloves (7359, -0.29 DPS) [crafted]; Mail Combat Gauntlets (4075, -0.29 DPS) [world_drop]; Bonefist Gauntlets (4465, -0.44 DPS, sim-verified) [world] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.17 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.27 DPS) | yes | Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Chausses of Westfall (6087, -0.20 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.43 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 attack_power points (0.68 DPS) | yes | Hard Gold Boots (250534, -0.02 DPS, sim-verified) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 attack_power points (0.78 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.0 attack_power points (0.59 DPS) | yes | Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Tiger Band (6749, -0.93 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (68.7 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -20.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.44 DPS) | yes | Long Battle Bow (15284, -0.15 DPS) [world_drop]; Fine Longbow (11304, -0.24 DPS) [vendor]; Double-barreled Shotgun (2098, -0.25 DPS, sim-verified) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 446, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 77.8. Weights run: 1.5s. Verify run: 1.4s. 636 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.304), strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.473 ± 0.059, hit=1.140 ± 0.055, melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 38.8 attack_power points (1.60 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -0.17 DPS) [crafted]; Tusken Helm (6686, -0.27 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.58 DPS) | yes | Kaleidoscope Chain (13084, -0.01 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.17 DPS) [world_drop]; Gazlowe's Charm (13088, -0.17 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 attack_power points (1.12 DPS) | yes | Chromite Pauldrons (8144, +0.00 DPS, sim-verified) [world_drop]; Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.9 attack_power points (0.61 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS, sim-verified) [dungeon]; Dark Hooded Cape (5257, -0.20 DPS) [world]; Sergeant Major's Cape (16315, -0.20 DPS) [pvp] |
| chest | Jouster's Chestplate (8157) | World drop [world_drop] | sim-verified (76.0 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Avenger's Armor (1488, -2.16 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS, sim-verified) [dungeon]; Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Yorgen Bracers (13012, -0.21 DPS) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 attack_power points (1.63 DPS) | yes | Reticulated Bone Gauntlets (9435, -0.15 DPS, sim-verified) [world_drop]; Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon] |
| waist | Highlander's Plate Girdle (20125) | The League of Arathor [rep] | sim-verified (76.0 DPS) | yes | Highlander's Leather Girdle (20116, -0.29 DPS) [rep]; Girdle of Golem Strength (9405, -0.31 DPS) [world_drop]; Boar Champion's Belt (10768, -2.16 DPS, sim-verified) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 attack_power points (2.15 DPS) | yes | Firemane Leggings (13129, +0.00 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.41 DPS) [crafted]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 attack_power points (1.33 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.83 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor]; Tiger Band (6749, -0.21 DPS) [quest] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 19.8 attack_power points (0.82 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS, sim-verified) [world_drop]; Suspicious Spare Part (274754, -0.10 DPS) [vendor]; Protector's Band (19517, -0.20 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (73.9 DPS) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Nightblade (1982, -41.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (73.9 DPS) | yes | Bow of Searing Arrows (2825, +0.00 DPS, sim-verified) [world_drop]; The Silencer (13138, -0.04 DPS) [world_drop]; Explosive Shotgun (8188, -0.10 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 636, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 115.5. Weights run: 1.7s. Verify run: 1.5s. 827 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.275), strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=6.161 ± 0.515, hit=0.978 ± 0.047, melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | 116.5 attack_power points (9.03 DPS) | yes | Ornate Mithril Helm (7937, -1.12 DPS) [crafted]; Raging Berserker's Helm (7719, -1.29 DPS, sim-verified) [dungeon]; Eye of Theradras (17715, -2.34 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (1.09 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop]; River Pride Choker (13087, -0.60 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 103.6 attack_power points (8.03 DPS) | yes | Officer's Pauldrons (250576, -0.54 DPS, sim-verified) [crafted]; Wyrmslayer Spaulders (13066, -6.39 DPS) [world_drop]; Earthslag Shoulders (11632, -6.44 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 attack_power points (1.10 DPS) | yes | Sergeant Major's Cape (16336, -0.23 DPS) [pvp]; Wolfmaster Cape (6314, -0.32 DPS) [dungeon]; Blackveil Cape (11626, -1.66 DPS, sim-verified) [dungeon] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 109.9 attack_power points (8.51 DPS) | yes | Ornate Mithril Breastplate (7935, -2.45 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -5.59 DPS) [quest]; Valorous Chestguard (8274, -6.08 DPS) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.17 DPS) | yes | Branded Leather Bracers (19508, -0.62 DPS) [dungeon]; Officer's Wristguards (250581, -0.70 DPS) [crafted]; Runed Golem Shackles (12550, -2.37 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 106.3 attack_power points (8.24 DPS) | yes | Dragonscale Gauntlets (8347, -1.24 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.55 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.55 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 106.3 attack_power points (8.24 DPS) | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20106, -0.21 DPS) [rep]; Highlander's Plate Girdle (20124, -0.33 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 172.5 attack_power points (13.37 DPS) | yes | Knight's Plate Leggings (220797, +0.00 DPS, sim-verified) [vendor]; Golem Shard Leggings (13074, -10.69 DPS) [world_drop]; Scarlet Leggings (10330, -10.81 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (115.5 DPS) | yes | Officer's Sabatons (250561, -0.04 DPS) [crafted]; Officer's Boots (250546, -0.09 DPS) [crafted]; Battlechaser's Greaves (12555, -2.30 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 29.8 attack_power points (2.31 DPS) | yes | Protector's Band (19516, -0.89 DPS) [rep]; Insurgent's Band (272065, -1.15 DPS) [vendor]; Protector's Band (19515, -1.16 DPS) [rep] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.55 DPS) | yes | Protector's Band (19516, +0.00 DPS, sim-verified) [rep]; Insurgent's Band (272065, -0.39 DPS) [vendor]; Protector's Band (19515, -0.40 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (113.2 DPS) | yes | Smoking Heart of the Mountain (11811, -1.29 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (113.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Fiery War Axe (870, -6.66 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (113.2 DPS) | yes | Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; The Silencer (13138, -0.37 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.64 DPS, sim-verified) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Stormshroud Pants; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Assault Band; trinket1: Frozen Heart of the Mountain; trinket2: Guardian Talisman; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 827, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 219.2. Weights run: 1.6s. Verify run: 1.3s. 1744 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.801), strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=20.276 ± 1.619, hit=3.037 ± 0.150, melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 698.4 attack_power points (22.36 DPS) | yes | Bloodvine Lens (19998, -4.18 DPS) [crafted]; Ragefury Eyepatch (11735, -5.45 DPS, sim-verified) [dungeon]; Lightbreaker Greathelm (239517, -5.49 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (219.2 DPS) | yes | Rage of Mugamba (19577, -1.72 DPS, sim-verified) [quest]; Amulet of the Darkmoon (19491, -8.03 DPS) [quest]; Beads of Ogre Might (22150, -8.12 DPS) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 460.8 attack_power points (14.75 DPS) | yes | Lieutenant Commander's Plate Shoulders (23315, -3.55 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -3.55 DPS) [vendor]; Darkspear Spaulders (272108, -8.28 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 283.9 attack_power points (9.09 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Shroud of Domination (22337, -6.97 DPS) [dungeon]; Howler's Furs (272414, -7.22 DPS) [vendor] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | sim-verified (219.2 DPS) | yes | Bloodsoul Breastplate (19690, -6.19 DPS) [crafted]; Stormshroud Armor (15056, -6.47 DPS) [crafted]; Breastplate of Undead Slaying (23087, -15.71 DPS, sim-verified) [world] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | sim-verified (219.2 DPS) | yes | Bracers of Undead Slaying (23090, -3.12 DPS, sim-verified) [world]; Lightbreaker Wrists (239512, -5.07 DPS) [vendor]; Deeprock Bracers (21184, -7.29 DPS) [quest] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | sim-verified (219.2 DPS) | yes | Marshal's Plate Gauntlets (16484, -2.11 DPS) [vendor]; Marshal's Plate Gauntlets (231541, -2.11 DPS) [vendor]; Razor Gauntlets (18326, -9.88 DPS, sim-verified) [dungeon] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 433.4 attack_power points (13.87 DPS) | yes | Highlander's Plate Girdle (20041, -2.67 DPS) [rep]; Highlander's Lamellar Girdle (20042, -2.92 DPS) [rep]; Radiant Girdle of the Dawn (227814, -3.01 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | sim-verified (219.2 DPS) | yes | Sentinel's Plate Legguards (237825, -2.20 DPS) [vendor]; Marshal's Plate Legguards (16479, -3.07 DPS) [vendor]; Cloudkeeper Legplates (14554, -15.47 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 459.8 attack_power points (14.72 DPS) | yes | Boots of Heroism (21995, -6.82 DPS, sim-verified) [quest]; Marshal's Plate Boots (16483, -11.14 DPS) [vendor]; Marshal's Plate Boots (231539, -11.14 DPS) [vendor] |
| finger1 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (219.2 DPS) | yes | Band of the Penitent (13217, -1.72 DPS) [quest]; Ring of Entropy (18543, -1.72 DPS) [world]; Naglering (11669, -5.02 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (219.2 DPS) | yes | Band of the Penitent (13217, -1.48 DPS) [quest]; Ring of Entropy (18543, -1.48 DPS) [world]; Naglering (11669, -5.54 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (219.2 DPS) | yes | Frozen Heart of the Mountain (249469, -4.09 DPS, sim-verified) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (219.2 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Nightfall (19169, -1.97 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (219.2 DPS) | yes | Bow of Searing Arrows (2825, -2.53 DPS, sim-verified) [world_drop]; Bloodseeker (19107, -7.88 DPS) [quest]; Unsophisticated Hand Cannon (18460, -8.09 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Chromatic Cloak; chest: Lightbreaker Cuirass; wrist: Vambraces of the Sadist; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Band of Earthen Might; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Blue Dragon; main_hand: The Unstoppable Force; ranged: The Purifier

No-known-source sample (15 of 1744, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (orc, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 30.2. Weights run: 1.2s. Verify run: 1.3s. 260 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000), hit=0.529 ± 0.020, melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.22 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 attack_power points (0.26 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Grave Shroud (279865, -0.09 DPS) [quest]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Subterranean Cape (14149, -0.10 DPS, sim-verified) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 attack_power points (0.88 DPS) | yes | Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.29 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Cryptwalker Bracers (280095, -0.10 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.18 DPS) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 attack_power points (0.70 DPS) | yes | Gold-flecked Gloves (5195, -0.10 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon]; Foreman's Gloves (2167, -0.26 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.17 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Support Girdle (1215, -0.32 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.8 attack_power points (0.79 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.09 DPS) [world_drop] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 attack_power points (0.44 DPS) | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | sim-verified (30.2 DPS) | yes | Forsaken Greataxe (251533, -0.14 DPS) [quest]; Smite's Mighty Hammer (7230, -0.34 DPS) [dungeon]; The Axe of Severing (23171, -13.32 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | 4.2 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, -0.01 DPS, sim-verified) [vendor]; Heavy Shortbow (3036, -0.09 DPS) [world_drop]; Orcish Battle Bow (5346, -0.09 DPS) [quest] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Veteran's Boots; finger1: Demon Band; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Hammerbone; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5031 ZZZZZZZZ; 5036 ZZZZZ; 5255 Quilboar Tomahawk; 5748 Centaur Longbow

### Band 30 (orc, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 72.0. Weights run: 1.2s. Verify run: 1.2s. 440 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.171 ± 0.023, hit=0.743 ± 0.031, melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 attack_power points (1.26 DPS) | yes | Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted]; Veteran's Chain Helm (250498, -0.22 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.68 DPS) | yes | Kaleidoscope Chain (13084, -0.28 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.29 DPS) [world_drop]; Scout's Medallion (19537, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 attack_power points (0.68 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.49 DPS) | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | sim-verified (72.0 DPS) | yes | Barbaric Iron Breastplate (7914, -0.19 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Avenger's Armor (1488, -2.19 DPS, sim-verified) [dungeon] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 attack_power points (0.78 DPS) | yes | Yorgen Bracers (13012, -0.14 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Cultist's Armguards (270032, -0.29 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.07 DPS) | yes | Warsong Gauntlets (16978, -0.12 DPS, sim-verified) [quest]; Bonefist Gauntlets (4465, -0.20 DPS) [world]; Heavy Earthen Gloves (7359, -0.29 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.17 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.27 DPS) | yes | Veteran's Silvered Chain Leggings (250523, -0.13 DPS, sim-verified) [crafted]; Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Slayer's Pants (14757, -0.20 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 attack_power points (0.68 DPS) | yes | Hard Gold Boots (250534, -0.03 DPS, sim-verified) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 attack_power points (0.78 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.0 attack_power points (0.59 DPS) | yes | Silverlaine's Family Seal (6321, -0.10 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Tiger Band (6749, -0.92 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (69.8 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -20.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.44 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS, sim-verified) [world_drop]; Long Battle Bow (15284, -0.15 DPS) [world_drop]; Fine Longbow (11304, -0.24 DPS) [vendor] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 440, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (orc, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 87.4. Weights run: 1.5s. Verify run: 1.4s. 628 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.304), strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.473 ± 0.059, hit=1.140 ± 0.055, melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 38.8 attack_power points (1.60 DPS) | yes | Icemetal Barbute (10763, +0.00 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -0.17 DPS) [crafted]; Tusken Helm (6686, -0.27 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.58 DPS) | yes | Ethereal Talisman (4430, +0.00 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; River Pride Choker (13087, -0.17 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 attack_power points (1.12 DPS) | yes | Chromite Pauldrons (8144, -0.05 DPS, sim-verified) [world_drop]; Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.41 DPS) | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Dark Hooded Cape (5257, -0.00 DPS) [world]; Lambent Scale Cloak (4706, -0.00 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) | World drop [world_drop] | sim-verified (85.5 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Avenger's Armor (1488, -2.39 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.83 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS, sim-verified) [dungeon]; Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Darkspear Armsplints (4132, -0.11 DPS) [quest] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 attack_power points (1.63 DPS) | yes | Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.40 DPS, sim-verified) [world_drop]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon] |
| waist | Defiler's Plate Girdle (20206) | The Defilers [rep] | sim-verified (85.5 DPS) | yes | Tharg's Shoelace (9705, -0.20 DPS) [quest]; Defiler's Leather Girdle (20192, -0.29 DPS) [rep]; Boar Champion's Belt (10768, -2.39 DPS, sim-verified) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 attack_power points (2.15 DPS) | yes | Firemane Leggings (13129, -0.08 DPS, sim-verified) [world_drop]; Orcish War Leggings (7929, -0.41 DPS) [crafted]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 attack_power points (1.33 DPS) | yes | Prowler's Leather Shoes (252465, -0.08 DPS, sim-verified) [crafted]; Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.83 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor]; Tiger Band (6749, -0.21 DPS) [quest] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 19.8 attack_power points (0.82 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS, sim-verified) [world_drop]; Suspicious Spare Part (274754, -0.10 DPS) [vendor]; Legionnaire's Band (19513, -0.20 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (83.1 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Nightblade (1982, -49.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (83.1 DPS) | yes | Bow of Searing Arrows (2825, +0.00 DPS, sim-verified) [world_drop]; The Silencer (13138, -0.04 DPS) [world_drop]; Explosive Shotgun (8188, -0.10 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 628, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (orc, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 127.4. Weights run: 1.7s. Verify run: 1.5s. 819 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.275), strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=6.161 ± 0.515, hit=0.978 ± 0.047, melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | 116.5 attack_power points (9.03 DPS) | yes | Ornate Mithril Helm (7937, -1.12 DPS) [crafted]; Raging Berserker's Helm (7719, -1.15 DPS, sim-verified) [dungeon]; Eye of Theradras (17715, -2.34 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (1.09 DPS) | yes | Woven Ivy Necklace (19159, +0.00 DPS, sim-verified) [quest]; Skibi's Pendant (13089, -0.19 DPS) [world_drop]; Ethereal Talisman (4430, -0.39 DPS) [quest] |
| shoulder | Blood Guard's Plate Pauldrons (220796) | Lady Palanseer [vendor] | 103.6 attack_power points (8.03 DPS) | yes | Officer's Pauldrons (250576, -0.98 DPS, sim-verified) [crafted]; Wyrmslayer Spaulders (13066, -6.39 DPS) [world_drop]; Earthslag Shoulders (11632, -6.44 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 attack_power points (1.10 DPS) | yes | Wolfmaster Cape (6314, -0.32 DPS) [dungeon]; Battlehard Cape (11858, -0.32 DPS) [quest]; Blackveil Cape (11626, -1.57 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 109.9 attack_power points (8.51 DPS) | yes | Ornate Mithril Breastplate (7935, -3.13 DPS, sim-verified) [crafted]; Warforged Chestplate (11195, -5.59 DPS) [quest]; Valorous Chestguard (8274, -6.08 DPS) [world_drop] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (2.17 DPS) | yes | Branded Leather Bracers (19508, -0.62 DPS) [dungeon]; Officer's Wristguards (250581, -0.70 DPS) [crafted]; Runed Golem Shackles (12550, -2.71 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 106.3 attack_power points (8.24 DPS) | yes | Dragonscale Gauntlets (8347, -1.05 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.55 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.55 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 106.3 attack_power points (8.24 DPS) | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.33 DPS) [rep]; Defiler's Chain Girdle (20153, -0.93 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | sim-verified (124.5 DPS) | yes | Stormshroud Pants (15057, -1.33 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -5.69 DPS) [world_drop]; Scarlet Leggings (10330, -5.81 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-verified (126.3 DPS) | yes | Officer's Sabatons (250561, -0.04 DPS) [crafted]; Officer's Boots (250546, -0.09 DPS) [crafted]; Battlechaser's Greaves (12555, -3.12 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 29.8 attack_power points (2.31 DPS) | yes | Assault Band (13095, -0.76 DPS) [world_drop]; Legionnaire's Band (19511, -0.89 DPS) [rep]; Insurgent's Band (272065, -1.15 DPS) [vendor] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.86 DPS) | yes | Assault Band (13095, +0.00 DPS, sim-verified) [world_drop]; Legionnaire's Band (19511, -0.44 DPS) [rep]; Insurgent's Band (272065, -0.70 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (123.2 DPS) | yes | Smoking Heart of the Mountain (11811, -4.01 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (123.2 DPS) | yes | Smoking Heart of the Mountain (11811, -1.24 DPS, sim-verified) [crafted] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (123.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Blight (7959, -3.47 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (123.2 DPS) | yes | Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; The Silencer (13138, -0.37 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.73 DPS, sim-verified) [world_drop] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Blood Guard's Plate Pauldrons; back: Bloodlust Cape; chest: Stone Guard's Plate Armor; wrist: Bracers of the Stone Princess; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 819, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (orc, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 246.0. Weights run: 1.6s. Verify run: 1.3s. 1736 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.801), strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=20.276 ± 1.619, hit=3.037 ± 0.150, melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 698.4 attack_power points (22.36 DPS) | yes | Bloodvine Lens (19998, -4.18 DPS) [crafted]; Lightbreaker Greathelm (239517, -5.49 DPS) [vendor]; Ragefury Eyepatch (11735, -7.76 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (246.0 DPS) | yes | Blazefury Medallion (17111, -3.92 DPS, sim-verified) [world]; Amulet of the Darkmoon (19491, -8.03 DPS) [quest]; Beads of Ogre Might (22150, -8.12 DPS) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 460.8 attack_power points (14.75 DPS) | yes | Champion's Plate Shoulders (23243, -3.55 DPS) [vendor]; Champion's Plate Shoulders (227042, -3.55 DPS) [vendor]; Darkspear Spaulders (272108, -8.93 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 283.9 attack_power points (9.09 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Shroud of Domination (22337, -6.97 DPS) [dungeon]; Howler's Furs (272414, -7.22 DPS) [vendor] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | sim-verified (246.0 DPS) | yes | Bloodsoul Breastplate (19690, -6.19 DPS) [crafted]; Stormshroud Armor (15056, -6.47 DPS) [crafted]; Breastplate of Undead Slaying (23087, -19.12 DPS, sim-verified) [world] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | sim-verified (246.0 DPS) | yes | Bracers of Undead Slaying (23090, -4.31 DPS, sim-verified) [world]; Lightbreaker Wrists (239512, -5.07 DPS) [vendor]; Deeprock Bracers (21184, -7.29 DPS) [quest] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | sim-verified (246.0 DPS) | yes | General's Plate Gauntlets (16548, -2.11 DPS) [vendor]; General's Plate Gauntlets (231532, -2.11 DPS) [vendor]; Razor Gauntlets (18326, -13.65 DPS, sim-verified) [dungeon] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 433.4 attack_power points (13.87 DPS) | yes | Defiler's Plate Girdle (20204, -2.67 DPS) [rep]; Radiant Girdle of the Dawn (227814, -3.37 DPS, sim-verified) [vendor]; Defiler's Plate Girdle (20205, -3.54 DPS) [rep] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | sim-verified (246.0 DPS) | yes | Sentinel's Plate Legguards (237825, -2.20 DPS) [vendor]; General's Plate Leggings (16543, -3.07 DPS) [vendor]; Cloudkeeper Legplates (14554, -19.77 DPS, sim-verified) [world_drop] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 459.8 attack_power points (14.72 DPS) | yes | Boots of Heroism (21995, -8.28 DPS, sim-verified) [quest]; General's Plate Boots (16545, -11.14 DPS) [vendor]; General's Plate Boots (231531, -11.14 DPS) [vendor] |
| finger1 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (246.0 DPS) | yes | Band of the Penitent (13217, -1.72 DPS) [quest]; Ring of Entropy (18543, -1.72 DPS) [world]; Naglering (11669, -6.33 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (246.0 DPS) | yes | Band of the Penitent (13217, -1.48 DPS) [quest]; Ring of Entropy (18543, -1.48 DPS) [world]; Naglering (11669, -6.19 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (246.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (246.0 DPS) | yes | Frozen Heart of the Mountain (249469, -5.99 DPS, sim-verified) [crafted] |
| main_hand | Nightfall (19169) | Blacksmithing [crafted] | sim-verified (246.0 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; The Unstoppable Force (19323, -19.31 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (246.0 DPS) | yes | Bow of Searing Arrows (2825, -3.45 DPS, sim-verified) [world_drop]; Bloodseeker (19107, -7.88 DPS) [quest]; Unsophisticated Hand Cannon (18460, -8.09 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Chromatic Cloak; chest: Lightbreaker Cuirass; wrist: Vambraces of the Sadist; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Band of Earthen Might; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Nightfall; ranged: The Purifier

No-known-source sample (15 of 1736, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

