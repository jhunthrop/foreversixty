# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 29.5. Weights run: 1.5s. Verify run: 1.3s. 241 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000), hit=not significant (0.000 ± 0.000), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 | yes | Defender's Leather Hood (252447, -0.21 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.88 DPS) [crafted]; Shadow Goggles (4373, -0.88 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.26 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 | yes | Catacomb Cloak (279899, -0.10 DPS) [quest]; Grave Shroud (279865, -0.11 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.18 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 | yes | Veteran's Chain Shirt (250488, -0.22 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 | yes | Cryptwalker Bracers (280095, -0.11 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted]; Burnished Bracers (3211, -0.26 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 | yes | Gold-flecked Gloves (5195, -0.11 DPS, sim-verified) [dungeon]; Polar Gauntlets (7606, -0.18 DPS) [quest]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.21 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Support Girdle (1215, -0.32 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.9 | yes | Defender's Leather Pants (252445, -0.18 DPS) [crafted]; Totemic Leather Pants (252446, -0.18 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.21 DPS, sim-verified) [crafted] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Protector's Band (20439)) | World drop [world_drop] | 8.3 | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 | yes | Loop of Sacrifice (281673, -0.11 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.7 | yes | Living Root (6631, -0.28 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.31 DPS, sim-verified) [dungeon]; Duskbringer (2205, -0.39 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 299.6 | yes | Lil Timmy's Peashooter (13136, +0.00 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.41 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.87 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Veteran's Boots; finger1: Demon Band; finger2: Protector's Band; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Forsaken Greataxe; ranged: Ranger Bow

No-known-source sample (15 of 241, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 9602 Brushwood Blade; 10047 Simple Kilt; 10421 Rough Copper Vest; 14145 Cursed Felblade; 14147 Cavedweller Bracers; 14148 Crystalline Cuffs

### Band 30 (human, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 70.5. Weights run: 1.6s. Verify run: 1.5s. 410 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.171 ± 0.023, hit=not significant (0.000 ± 0.000), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 | yes | Veteran's Chain Helm (250498, -0.05 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, -0.21 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -0.29 DPS) [world_drop]; Sentinel's Medallion (19541, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Sergeant Major's Cape (16315, -0.04 DPS, sim-verified) [pvp]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 27.9 | yes | Barbaric Iron Breastplate (7914, -0.17 DPS, sim-verified) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 | yes | Yorgen Bracers (13012, -0.16 DPS, sim-verified) [world_drop]; Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Cultist's Armguards (270032, -0.29 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 | yes | Bonefist Gauntlets (4465, -0.17 DPS, sim-verified) [world]; Heavy Earthen Gloves (7359, -0.29 DPS) [crafted]; Mail Combat Gauntlets (4075, -0.29 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Veteran's Silvered Chain Leggings (250523, -0.15 DPS, sim-verified) [crafted]; Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Chausses of Westfall (6087, -0.20 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 | yes | Hard Gold Boots (250534, -0.02 DPS, sim-verified) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 | yes | Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor]; Ironspine's Eye (7686, -0.39 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.0 | yes | Silverlaine's Family Seal (6321, -0.07 DPS, sim-verified) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Ironspine's Eye (7686, -0.19 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (70.4 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -20.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 378.6 | yes | Silver Star (3463, -0.49 DPS, sim-verified) [quest]; Ironweaver (13137, -1.17 DPS) [world_drop]; Nightstalker Bow (6696, -2.30 DPS) [dungeon] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 7957 Bronze Greatsword; 7958 Bronze Battle Axe

### Band 40 (human, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 97.3. Weights run: 2.1s. Verify run: 1.8s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.304), strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.473 ± 0.059, hit=not significant (0.000 ± 0.000), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 38.8 | yes | Hard Gold Coif (250537, -0.17 DPS) [crafted]; Tusken Helm (6686, -0.27 DPS) [dungeon]; Icemetal Barbute (10763, -0.63 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | River Pride Choker (13087, -0.17 DPS) [world_drop]; Gazlowe's Charm (13088, -0.17 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.45 DPS, sim-verified) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 | yes | Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [world_drop]; Chromite Pauldrons (8144, -0.31 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16336) | Rank 9 (Alliance) [pvp] | 14.9 | yes | Sergeant Major's Cape (16315, -0.20 DPS) [pvp]; Lambent Scale Cloak (4706, -0.21 DPS) [world_drop]; Wolfmaster Cape (6314, -0.77 DPS, sim-verified) [dungeon] |
| chest | Jouster's Chestplate (8157) | World drop [world_drop] | 37.1 | yes | Kolkar Marauder Chain (6773, +0.00 DPS, sim-verified) [quest]; Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Golden Scale Cuirass (3845, -0.10 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Yorgen Bracers (13012, -0.21 DPS) [world_drop]; Pugilist Bracers (4438, -0.47 DPS, sim-verified) [dungeon] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 | yes | Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.89 DPS, sim-verified) [world_drop] |
| waist | Highlander's Plate Girdle (20125) | The League of Arathor [rep] | 37.1 | yes | Highlander's Leather Girdle (20116, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.31 DPS) [world_drop]; Scarlet Belt (10329, -0.31 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 | yes | Orcish War Leggings (7929, -0.41 DPS) [crafted]; Firemane Leggings (13129, -0.47 DPS, sim-verified) [world_drop]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.47 DPS, sim-verified) [crafted] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor]; Silverlaine's Family Seal (6321, -0.32 DPS) [dungeon] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 19.8 | yes | Suspicious Spare Part (274754, -0.10 DPS) [vendor]; Protector's Band (19517, -0.20 DPS) [rep]; Thunderbrow Ring (13097, -0.58 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (97.6 DPS) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Fiery War Axe (870, -18.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | 554.4 | yes | Swiftwind (13038, -1.38 DPS, sim-verified) [world_drop]; Crusader Bow (15287, -1.99 DPS) [world_drop]; Skystriker Bow (13020, -2.19 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; ranged: The Silencer

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (human, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 116.7. Weights run: 2.2s. Verify run: 1.9s. 754 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.275), strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=6.161 ± 0.515, hit=not significant (0.000 ± 0.000), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | sim-verified (115.7 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.37 DPS) [crafted]; Raging Berserker's Helm (7719, -1.51 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop]; River Pride Choker (13087, -0.60 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 103.6 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -6.26 DPS) [crafted]; Wyrmslayer Spaulders (13066, -6.39 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 14.2 | yes | Sergeant Major's Cape (16336, -0.18 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.32 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.52 DPS) [pvp] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 109.9 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -1.83 DPS) [crafted]; Warforged Chestplate (11195, -5.59 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.63 DPS, sim-verified) [dungeon]; Officer's Wristguards (250581, -0.70 DPS) [crafted]; Giantslayer Bracers (13076, -0.84 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 106.3 | yes | Dragonscale Gauntlets (8347, -1.20 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.55 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.55 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 106.3 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20106, -0.21 DPS) [rep]; Highlander's Plate Girdle (20124, -0.33 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | sim-verified (116.0 DPS) | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.80 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -5.69 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 23.6 | yes | Officer's Boots (250546, -0.09 DPS) [crafted]; Officer's Sabatons (250561, -0.13 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -0.20 DPS) [crafted] |
| finger1 | Assault Band (13095) (or Blackstone Ring (17713)) | World drop [world_drop] | 20.0 | yes | Protector's Band (19516, -0.13 DPS) [rep]; Insurgent's Band (272065, -0.39 DPS) [vendor]; Protector's Band (19515, -0.40 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Protector's Band (19516, -0.31 DPS, sim-verified) [rep]; Insurgent's Band (272065, -0.39 DPS) [vendor]; Protector's Band (19515, -0.40 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (98.0 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | - | - |  |  |  |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (112.7 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Fiery War Axe (870, -7.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (112.7 DPS) | yes | Dark Iron Rifle (16004, -0.79 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -4.27 DPS) [dungeon]; Gryphonwing Long Bow (13022, -4.28 DPS) [world_drop] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Blight; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 754, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (human, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 251.7. Weights run: 2.2s. Verify run: 1.8s. 1359 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.801), strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=20.276 ± 1.619, hit=not significant (0.000 ± 0.000), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 637.7 | yes | Bloodvine Lens (19998, -2.24 DPS) [crafted]; Lightbreaker Greathelm (239517, -4.52 DPS) [vendor]; Ragefury Eyepatch (11735, -6.31 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (247.2 DPS) | yes | Blazefury Medallion (17111, -3.59 DPS, sim-verified) [world]; Amulet of the Darkmoon (19491, -8.03 DPS) [quest]; Strength of Mugamba (19576, -8.62 DPS) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 430.4 | yes | Champion's Plate Shoulders (23243, -2.58 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (23315, -2.58 DPS) [vendor]; Darkspear Spaulders (272108, -8.81 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 283.9 | yes | Tattered Hakkari Cape (20219, -1.51 DPS, sim-verified) [quest]; Cloak of the Honor Guard (20073, -7.85 DPS) [rep]; Sergeant Major's Cape (16337, -7.85 DPS) [pvp] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | 769.7 | yes | Stormshroud Armor (15056, -6.47 DPS) [crafted]; Savage Gladiator Chain (11726, -8.59 DPS) [dungeon]; Bloodsoul Breastplate (19690, -10.48 DPS, sim-verified) [crafted] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (251.7 DPS) | yes | Deeprock Bracers (21184, -1.24 DPS) [quest]; Berserker Bracers (19578, -1.31 DPS) [rep]; Vambraces of the Sadist (13400, -5.71 DPS, sim-verified) [dungeon] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | 427.6 | yes | Marshal's Plate Gauntlets (16484, -2.11 DPS) [vendor]; General's Plate Gauntlets (16548, -2.11 DPS) [vendor]; Gauntlets of Heroism (21998, -4.96 DPS, sim-verified) [quest] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 433.4 | yes | Highlander's Plate Girdle (20041, -2.67 DPS) [rep]; Highlander's Lamellar Girdle (20042, -2.92 DPS) [rep]; Radiant Girdle of the Dawn (227814, -4.05 DPS, sim-verified) [vendor] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | 771.6 | yes | Sentinel's Plate Legguards (237825, -3.68 DPS, sim-verified) [vendor]; Marshal's Plate Legguards (16479, -4.04 DPS) [vendor]; General's Plate Leggings (16543, -4.04 DPS) [vendor] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 429.5 | yes | Boots of Heroism (21995, -6.48 DPS, sim-verified) [quest]; Marshal's Plate Boots (16483, -11.14 DPS) [vendor]; General's Plate Boots (16545, -11.14 DPS) [vendor] |
| finger1 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (247.2 DPS) | yes | Band of the Penitent (13217, -0.75 DPS) [quest]; Ring of Entropy (18543, -0.75 DPS) [world]; Wrath of Cenarius (21190, -6.70 DPS, sim-verified) [quest] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (247.2 DPS) | yes | Band of the Penitent (13217, -0.51 DPS) [quest]; Ring of Entropy (18543, -0.51 DPS) [world]; Wrath of Cenarius (21190, -7.33 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (247.2 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Darkmoon Card: Heroism (19287, -6.03 DPS, sim-verified) [quest] |
| trinket2 | - | - |  |  |  |
| main_hand | High Warlord's Pig Poker (234548) | Sergeant Thunderhorn [vendor] | sim-verified (247.2 DPS) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; High Warlord's Pig Sticker (234547, -6.57 DPS) [vendor]; Arcanite Champion (12790, -16.13 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (247.2 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -3.06 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Chromatic Cloak; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Band of Earthen Might; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Shard of the Fallen Star; main_hand: High Warlord's Pig Poker; ranged: The Purifier

No-known-source sample (15 of 1359, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (orc, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 32.4. Weights run: 1.5s. Verify run: 1.4s. 243 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000), hit=not significant (0.000 ± 0.000), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 | yes | Defender's Leather Hood (252447, -0.09 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.88 DPS) [crafted]; Shadow Goggles (4373, -0.88 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.26 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 | yes | Subterranean Cape (14149, -0.07 DPS, sim-verified) [dungeon]; Grave Shroud (279865, -0.09 DPS) [quest]; Catacomb Cloak (279899, -0.10 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 | yes | Veteran's Chain Shirt (250488, -0.17 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 | yes | Cryptwalker Bracers (280095, -0.07 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.18 DPS) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.7 | yes | Gold-flecked Gloves (5195, -0.07 DPS, sim-verified) [dungeon]; Blackened Defias Gloves (10401, -0.18 DPS) [dungeon]; Foreman's Gloves (2167, -0.26 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.09 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Support Girdle (1215, -0.32 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.8 | yes | Defender's Leather Pants (252445, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Hulking Leggings (14748, -0.09 DPS) [world_drop] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Demon Band (12054) (or Legionnaire's Band (20429)) | World drop [world_drop] | 8.3 | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 | yes | Loop of Sacrifice (281673, -0.07 DPS, sim-verified) [quest]; The 1 Ring (8350, -0.26 DPS) [world]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | sim-verified (32.4 DPS) | yes | Smite's Mighty Hammer (7230, -0.20 DPS) [dungeon]; Living Root (6631, -0.28 DPS) [dungeon]; Hammerbone (270018, -2.38 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 299.6 | yes | Lil Timmy's Peashooter (13136, +0.00 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.41 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.87 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Veteran's Boots; finger1: Demon Band; finger2: Legionnaire's Band; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Forsaken Greataxe; ranged: Ranger Bow

No-known-source sample (15 of 243, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 9602 Brushwood Blade; 10047 Simple Kilt; 10421 Rough Copper Vest; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 30 (orc, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 72.0. Weights run: 1.6s. Verify run: 1.5s. 415 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.171 ± 0.023, hit=not significant (0.000 ± 0.000), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 | yes | Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Veteran's Chain Helm (250498, -0.13 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | River Pride Choker (13087, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.39 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.68 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Slayer's Cape (14752, -0.10 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 27.9 | yes | Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.39 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.42 DPS, sim-verified) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 15.9 | yes | Bands of Serra'kis (6902, -0.19 DPS) [dungeon]; Cultist's Armguards (270032, -0.29 DPS) [quest]; Yorgen Bracers (13012, -0.41 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 | yes | Warsong Gauntlets (16978, -0.06 DPS, sim-verified) [quest]; Bonefist Gauntlets (4465, -0.20 DPS) [world]; Heavy Earthen Gloves (7359, -0.29 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Girdle of Golem Strength (9405, -0.01 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.01 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Slayer's Pants (14757, -0.20 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -0.20 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 | yes | Hard Gold Boots (250534, -0.02 DPS, sim-verified) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.0 | yes | Silverlaine's Family Seal (6321, -0.29 DPS) [dungeon]; Insurgent's Band (272067, -0.34 DPS) [vendor]; Ironspine's Eye (7686, -0.39 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.0 | yes | Insurgent's Band (272067, -0.15 DPS) [vendor]; Silverlaine's Family Seal (6321, -0.15 DPS, sim-verified) [dungeon]; Ironspine's Eye (7686, -0.19 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.0 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Corpsemaker (6687, -0.13 DPS) [dungeon]; Viscous Hammer (13045, -21.30 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 378.6 | yes | Silver Star (3463, -0.70 DPS, sim-verified) [quest]; Ironweaver (13137, -1.17 DPS) [world_drop]; Nightstalker Bow (6696, -2.30 DPS) [dungeon] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 415, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7955 Copper Claymore; 7956 Bronze Warhammer; 7957 Bronze Greatsword

### Band 40 (orc, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 98.2. Weights run: 2.1s. Verify run: 1.8s. 589 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.304), strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.473 ± 0.059, hit=not significant (0.000 ± 0.000), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 38.8 | yes | Icemetal Barbute (10763, +0.00 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -0.17 DPS) [crafted]; Tusken Helm (6686, -0.27 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; River Pride Choker (13087, -0.17 DPS) [world_drop]; Ethereal Talisman (4430, -0.46 DPS, sim-verified) [quest] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 | yes | Shining Mithril Pauldrons (250541, -0.10 DPS) [crafted]; Chromite Pauldrons (8144, -0.14 DPS, sim-verified) [world_drop]; Imperial Leather Spaulders (4737, -0.20 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Lambent Scale Cloak (4706, -0.00 DPS) [world_drop]; Warden's Cloak (14602, -0.00 DPS) [world_drop] |
| chest | Jouster's Chestplate (8157) | World drop [world_drop] | 37.1 | yes | Kolkar Marauder Chain (6773, +0.00 DPS, sim-verified) [quest]; Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Golden Scale Cuirass (3845, -0.10 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Darkspear Armsplints (4132, -0.11 DPS) [quest]; Pugilist Bracers (4438, -0.44 DPS, sim-verified) [dungeon] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 | yes | Gauntlets of Divinity (7724, -0.31 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Reticulated Bone Gauntlets (9435, -0.45 DPS, sim-verified) [world_drop] |
| waist | Defiler's Plate Girdle (20206) | The Defilers [rep] | 37.1 | yes | Defiler's Leather Girdle (20192, -0.29 DPS) [rep]; Girdle of Golem Strength (9405, -0.31 DPS) [world_drop]; Tharg's Shoelace (9705, -0.50 DPS, sim-verified) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 | yes | Orcish War Leggings (7929, -0.41 DPS) [crafted]; Firemane Leggings (13129, -0.50 DPS, sim-verified) [world_drop]; Symbolic Legplates (14829, -0.61 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Obsidian Greaves (13068, -0.31 DPS) [world_drop]; Prowler's Leather Shoes (252465, -0.50 DPS, sim-verified) [crafted] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Suspicious Spare Part (274754, -0.11 DPS) [vendor]; Silverlaine's Family Seal (6321, -0.32 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 19.8 | yes | Thunderbrow Ring (13097, -0.03 DPS, sim-verified) [world_drop]; Suspicious Spare Part (274754, -0.10 DPS) [vendor]; Legionnaire's Band (19513, -0.20 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (98.6 DPS) | yes | Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Fiery War Axe (870, -9.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | 554.4 | yes | Swiftwind (13038, -0.76 DPS, sim-verified) [world_drop]; Crusader Bow (15287, -1.99 DPS) [world_drop]; Skystriker Bow (13020, -2.19 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; chest: Jouster's Chestplate; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Assault Band; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; ranged: The Silencer

No-known-source sample (15 of 589, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (orc, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 125.0. Weights run: 2.2s. Verify run: 1.9s. 759 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.275), strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=6.161 ± 0.515, hit=not significant (0.000 ± 0.000), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) (or Blood Guard's Plate Helm (220803), Knight-Lieutenant's Plate Helm (220804)) | Scarlet Monastery: Herod [dungeon] | 106.7 | yes | Blood Guard's Plate Helm (220803, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.37 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Woven Ivy Necklace (19159, +0.00 DPS, sim-verified) [quest]; Skibi's Pendant (13089, -0.19 DPS) [world_drop]; Ethereal Talisman (4430, -0.39 DPS) [quest] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 103.6 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -6.26 DPS) [crafted]; Wyrmslayer Spaulders (13066, -6.39 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | Maraudon: Princess Theradras [dungeon] | 14.2 | yes | Battlehard Cape (11858, -0.32 DPS) [quest]; Wildhunter Cloak (16658, -0.32 DPS) [quest]; Wolfmaster Cape (6314, -0.63 DPS, sim-verified) [dungeon] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 109.9 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -1.83 DPS) [crafted]; Warforged Chestplate (11195, -5.59 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.41 DPS, sim-verified) [dungeon]; Officer's Wristguards (250581, -0.70 DPS) [crafted]; Giantslayer Bracers (13076, -0.84 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 106.3 | yes | Dragonscale Gauntlets (8347, -1.31 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.55 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.55 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 106.3 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.33 DPS) [rep]; Defiler's Chain Girdle (20153, -0.93 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | sim-verified (125.0 DPS) | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.69 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -5.69 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 23.6 | yes | Officer's Sabatons (250561, -0.08 DPS, sim-verified) [crafted]; Officer's Boots (250546, -0.09 DPS) [crafted]; Skulker's Leather Boots (252469, -0.20 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Blackstone Ring (17713, -0.31 DPS) [dungeon]; Legionnaire's Band (19511, -0.44 DPS) [rep]; Insurgent's Band (272065, -0.70 DPS) [vendor] |
| finger2 | Assault Band (13095) (or Blackstone Ring (17713)) | World drop [world_drop] | 20.0 | yes | Blackstone Ring (17713, +0.00 DPS, sim-verified) [dungeon]; Legionnaire's Band (19511, -0.13 DPS) [rep]; Insurgent's Band (272065, -0.39 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (102.3 DPS) | yes | Tidal Charm (1404, -3.26 DPS) [vendor]; Guardian Talisman (1490, -3.26 DPS) [quest]; Smoking Heart of the Mountain (11811, -4.96 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (102.4 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Smoking Heart of the Mountain (11811, -1.63 DPS, sim-verified) [crafted] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (124.4 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Blight (7959, -4.30 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (124.4 DPS) | yes | Dark Iron Rifle (16004, -0.85 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -4.27 DPS) [dungeon]; Gryphonwing Long Bow (13022, -4.28 DPS) [world_drop] |

**New at 50:** shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Assault Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Fiery War Axe; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 759, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (orc, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 261.2. Weights run: 2.2s. Verify run: 1.8s. 1365 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.801), strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=20.276 ± 1.619, hit=not significant (0.000 ± 0.000), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 637.7 | yes | Bloodvine Lens (19998, -2.24 DPS) [crafted]; Lightbreaker Greathelm (239517, -4.52 DPS) [vendor]; Ragefury Eyepatch (11735, -7.03 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (257.0 DPS) | yes | Rage of Mugamba (19577, -2.54 DPS, sim-verified) [quest]; Amulet of the Darkmoon (19491, -8.03 DPS) [quest]; Strength of Mugamba (19576, -8.62 DPS) [quest] |
| shoulder | Lightbreaker Shoulders (239516) | Leonid Barthalomew the Revered [vendor] | 430.4 | yes | Champion's Plate Shoulders (23243, -2.58 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (23315, -2.58 DPS) [vendor]; Darkspear Spaulders (272108, -8.40 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 283.9 | yes | Tattered Hakkari Cape (20219, -0.31 DPS, sim-verified) [quest]; Deathguard's Cloak (20068, -7.85 DPS) [rep]; Bloodlust Cape (14801, -7.97 DPS) [dungeon] |
| chest | Lightbreaker Cuirass (239519) | Leonid Barthalomew the Revered [vendor] | 769.7 | yes | Stormshroud Armor (15056, -6.47 DPS) [crafted]; Savage Gladiator Chain (11726, -8.59 DPS) [dungeon]; Bloodsoul Breastplate (19690, -10.61 DPS, sim-verified) [crafted] |
| wrist | Lightbreaker Wrists (239512) | Leonid Barthalomew the Revered [vendor] | sim-verified (261.2 DPS) | yes | Deeprock Bracers (21184, -1.24 DPS) [quest]; Berserker Bracers (19578, -1.31 DPS) [rep]; Vambraces of the Sadist (13400, -6.76 DPS, sim-verified) [dungeon] |
| hands | Lightbreaker Grips (239514) | Leonid Barthalomew the Revered [vendor] | 427.6 | yes | Marshal's Plate Gauntlets (16484, -2.11 DPS) [vendor]; General's Plate Gauntlets (16548, -2.11 DPS) [vendor]; Gauntlets of Heroism (21998, -4.71 DPS, sim-verified) [quest] |
| waist | Lightbreaker Belt (239513) | Leonid Barthalomew the Revered [vendor] | 433.4 | yes | Defiler's Plate Girdle (20204, -2.67 DPS) [rep]; Radiant Girdle of the Dawn (227814, -3.25 DPS, sim-verified) [vendor]; Defiler's Plate Girdle (20205, -3.54 DPS) [rep] |
| legs | Lightbreaker Tassets (239518) | Leonid Barthalomew the Revered [vendor] | 771.6 | yes | Sentinel's Plate Legguards (237825, -3.10 DPS, sim-verified) [vendor]; Marshal's Plate Legguards (16479, -4.04 DPS) [vendor]; General's Plate Leggings (16543, -4.04 DPS) [vendor] |
| feet | Lightbreaker Greaves (239515) | Leonid Barthalomew the Revered [vendor] | 429.5 | yes | Boots of Heroism (21995, -6.15 DPS, sim-verified) [quest]; Marshal's Plate Boots (16483, -11.14 DPS) [vendor]; General's Plate Boots (16545, -11.14 DPS) [vendor] |
| finger1 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (257.0 DPS) | yes | Band of the Penitent (13217, -0.75 DPS) [quest]; Ring of Entropy (18543, -0.75 DPS) [world]; Wrath of Cenarius (21190, -5.72 DPS, sim-verified) [quest] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (257.0 DPS) | yes | Band of the Penitent (13217, -0.51 DPS) [quest]; Ring of Entropy (18543, -0.51 DPS) [world]; Wrath of Cenarius (21190, -6.95 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (251.0 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (257.0 DPS) | yes | Tidal Charm (1404, -1.34 DPS) [vendor]; Guardian Talisman (1490, -1.34 DPS) [quest]; Shard of the Fallen Star (21891, -5.34 DPS, sim-verified) [world_drop] |
| main_hand | High Warlord's Pig Poker (234548) | Sergeant Thunderhorn [vendor] | sim-verified (257.0 DPS) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; High Warlord's Pig Sticker (234547, -6.57 DPS) [vendor]; Nightfall (19169, -8.41 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (257.0 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -2.13 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Lightbreaker Shoulders; back: Chromatic Cloak; chest: Lightbreaker Cuirass; wrist: Lightbreaker Wrists; hands: Lightbreaker Grips; waist: Lightbreaker Belt; legs: Lightbreaker Tassets; feet: Lightbreaker Greaves; finger1: Band of Earthen Might; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: High Warlord's Pig Poker; ranged: The Purifier

No-known-source sample (15 of 1365, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

