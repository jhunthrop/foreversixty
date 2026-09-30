# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 29.3. Weights run: 1.6s. Verify run: 1.2s. 438 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000), hit=not significant (0.000 ± 0.000), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 | yes | Defender's Leather Hood (252447, -0.09 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.88 DPS) [crafted]; Shadow Goggles (4373, -0.88 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.26 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 | yes | Grave Shroud (279865, -0.02 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.18 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 | yes | Veteran's Chain Shirt (250488, -0.23 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 | yes | Cryptwalker Bracers (280095, -0.02 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted]; Burnished Bracers (3211, -0.26 DPS) [world_drop] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.6 | yes | Polar Gauntlets (7606, -0.02 DPS, sim-verified) [quest]; Blackened Defias Gloves (10401, -0.09 DPS) [dungeon]; Foreman's Gloves (2167, -0.18 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.09 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Support Girdle (1215, -0.32 DPS) [world] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.9 | yes | Veteran's Chain Leggings (250493, -0.09 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.18 DPS) [crafted]; Totemic Leather Pants (252446, -0.18 DPS) [crafted] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.3 | yes | The 1 Ring (8350, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Minor Channeling Ring (1449, -0.35 DPS) [quest] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.3 | yes | The 1 Ring (8350, -0.09 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Minor Channeling Ring (1449, -0.26 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.7 | yes | Smite's Mighty Hammer (7230, -0.21 DPS, sim-verified) [dungeon]; Living Root (6631, -0.28 DPS) [dungeon]; Duskbringer (2205, -0.39 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 299.6 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.87 DPS) [crafted]; Venomstrike (6469, -3.01 DPS) [dungeon] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Veteran's Boots; finger1: Protector's Band; finger2: Loop of Sacrifice; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Forsaken Greataxe; ranged: Ranger Bow

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 30 (human, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 70.3. Weights run: 1.7s. Verify run: 1.5s. 891 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.171 ± 0.023, hit=not significant (0.000 ± 0.000), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 | yes | Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Veteran's Chain Helm (250498, -0.16 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (20444, -0.68 DPS) [rep]; Erudite's Amulet (277204, -0.68 DPS) [quest]; Sentinel's Medallion (19541, -0.94 DPS, sim-verified) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.14 DPS, sim-verified) [pvp]; Catacomb Cloak (279899, -0.20 DPS) [quest] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 27.9 | yes | Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.31 DPS, sim-verified) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.39 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | World drop [world_drop] | 15.9 | yes | Cultist's Armguards (270032, -0.29 DPS) [quest]; Patterned Bronze Bracers (2868, -0.29 DPS) [crafted]; Bands of Serra'kis (6902, -0.31 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 | yes | Heavy Earthen Gloves (7359, -0.29 DPS) [crafted]; Mail Combat Gauntlets (4075, -0.29 DPS) [world_drop]; Bonefist Gauntlets (4465, -0.31 DPS, sim-verified) [world] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Highlander's Plate Girdle (20126, -0.01 DPS) [rep]; Highlander's Lamellar Girdle (20108, -0.10 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Chausses of Westfall (6087, -0.20 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.29 DPS, sim-verified) [crafted] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 | yes | Hard Gold Boots (250534, -0.02 DPS, sim-verified) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.0 | yes | Insurgent's Band (272067, -0.15 DPS) [vendor]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Protector's Band (20439, -0.20 DPS) [rep] |
| finger2 | Silverlaine's Family Seal (6321) | Shadowfang Keep: Baron Silverlaine [dungeon] | 10.0 | yes | Insurgent's Band (272067, -0.01 DPS, sim-verified) [vendor]; Ironspine's Eye (7686, -0.09 DPS) [dungeon]; Seal of Wrynn (2933, -0.19 DPS) [quest] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Corpsemaker (6687, -0.13 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.75 DPS) [vendor]; Morbid Dawn (7689, -17.65 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 378.6 | yes | Silver Star (3463, -0.58 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.30 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.23 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Protector's Band; finger2: Silverlaine's Family Seal; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 891, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer

### Band 40 (human, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 95.2. Weights run: 2.1s. Verify run: 1.8s. 1330 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.304), strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.473 ± 0.059, hit=not significant (0.000 ± 0.000), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 38.8 | yes | Hard Gold Coif (250537, -0.17 DPS) [crafted]; Tusken Helm (6686, -0.27 DPS) [dungeon]; Icemetal Barbute (10763, -1.05 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, -0.07 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.58 DPS) [rep]; Sentinel's Medallion (20444, -0.58 DPS) [rep] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 | yes | Shining Mithril Pauldrons (250541, -0.08 DPS, sim-verified) [crafted]; Imperial Leather Spaulders (4737, -0.20 DPS) [world_drop]; Wrangling Spaulders (15698, -0.31 DPS) [quest] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 14.9 | yes | Sergeant Major's Cape (16315, -0.20 DPS) [pvp]; Lambent Scale Cloak (4706, -0.21 DPS) [world_drop]; Wolfmaster Cape (6314, -0.64 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 34.7 | yes | Golden Scale Cuirass (3845, -0.00 DPS) [crafted]; Shining Mithril Breastplate (250540, -0.00 DPS) [crafted]; Shining Silver Breastplate (2870, -0.56 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.11 DPS) [world]; Golden Scale Bracers (6040, -0.21 DPS) [crafted]; Pugilist Bracers (4438, -0.32 DPS, sim-verified) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 | yes | Gauntlets of Divinity (7724, +0.00 DPS, sim-verified) [dungeon]; Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Seawolf Gloves (4509, -0.51 DPS) [quest] |
| waist | Highlander's Plate Girdle (20125) | The League of Arathor [rep] | 37.1 | yes | Highlander's Leather Girdle (20116, +0.00 DPS, sim-verified) [rep]; Scarlet Belt (10329, -0.31 DPS) [dungeon]; Highlander's Lamellar Girdle (20107, -0.31 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 | yes | Orcish War Leggings (7929, -0.58 DPS, sim-verified) [crafted]; Ornate Mithril Pants (7926, -0.92 DPS) [crafted]; Dual Reinforced Leggings (9625, -0.92 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Ironheel Boots (4653, -0.31 DPS) [quest]; Prowler's Leather Shoes (252465, -0.32 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 19.8 | yes | Protector's Band (19517, -0.20 DPS) [rep]; Silverlaine's Family Seal (6321, -0.31 DPS) [dungeon]; Insurgent's Band (272066, -0.32 DPS) [vendor] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 17.3 | yes | Insurgent's Band (272066, -0.22 DPS) [vendor]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.32 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Bonebiter (6830, +0.00 DPS) [quest]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | World drop [world_drop] | 488.3 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Rifle (17687, -0.34 DPS) [quest]; Mithril Blacksmith Hammer (285280, -0.37 DPS) [crafted] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Suspicious Spare Part; trinket1: Ankh of Life; trinket2: Rune of Perfection; ranged: Sniper Rifle

No-known-source sample (15 of 1330, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (human, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 117.7. Weights run: 2.2s. Verify run: 1.9s. 1762 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.275), strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=6.161 ± 0.515, hit=not significant (0.000 ± 0.000), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | 0.0 | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.37 DPS) [crafted]; Raging Berserker's Helm (7719, -1.23 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19539, -0.38 DPS, sim-verified) [rep]; Talisman of the Naga Lord (5029, -0.84 DPS) [world]; Sentinel's Medallion (19540, -0.84 DPS) [rep] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 103.6 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -6.26 DPS) [crafted]; Wyrmslayer Spaulders (13066, -6.39 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 | yes | Sergeant Major's Cape (16336, -0.21 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.32 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.52 DPS) [pvp] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 109.9 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -1.83 DPS) [crafted]; Warforged Chestplate (11195, -5.59 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Officer's Wristguards (250581, -0.70 DPS) [crafted]; Branded Leather Bracers (19508, -0.70 DPS, sim-verified) [dungeon]; Prowler's Leather Bracers (252539, -0.92 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 106.3 | yes | Dragonscale Gauntlets (8347, -1.17 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.55 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.55 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 106.3 | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Lamellar Girdle (20106, -0.21 DPS) [rep]; Highlander's Plate Girdle (20124, -0.33 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 0.0 | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.94 DPS, sim-verified) [crafted]; Scarlet Leggings (10330, -5.81 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 23.6 | yes | Officer's Boots (250546, -0.09 DPS) [crafted]; Officer's Sabatons (250561, -0.10 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -0.20 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Insurgent's Band (272065, -0.39 DPS) [vendor]; Insurgent's Band (272066, -0.62 DPS) [vendor]; Suspicious Spare Part (274754, -0.70 DPS) [vendor] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 18.3 | yes | Protector's Band (19515, -0.27 DPS) [rep]; Insurgent's Band (272066, -0.49 DPS) [vendor]; Insurgent's Band (272065, -1.38 DPS, sim-verified) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | 0.0 | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 0.0 | yes | Dark Iron Rifle (16004, -0.89 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -4.27 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -4.80 DPS) [world_drop] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Blight; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1762, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (human, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 183.5. Weights run: 2.2s. Verify run: 1.9s. 2570 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.801), strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=20.276 ± 1.619, hit=not significant (0.000 ± 0.000), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 637.7 | yes | Bloodvine Lens (19998, -2.24 DPS) [crafted]; Ragefury Eyepatch (11735, -6.78 DPS, sim-verified) [dungeon]; Field Marshal's Plate Helm (16478, -7.84 DPS) [vendor] |
| neck | Fury of the Forgotten Swarm (21809) | World drop [world_drop] | 0.0 | yes | Medallion of the Dawn (22659, -0.23 DPS) [quest]; Onyxia Tooth Pendant (18404, -0.78 DPS) [quest]; Amulet of the Darkmoon (19491, -8.26 DPS) [quest] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 326.6 | yes | Champion's Plate Shoulders (23243, +0.00 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (23315, +0.00 DPS) [vendor]; Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 0.0 | yes | Drape of Unyielding Strength (21394, -0.02 DPS) [quest]; Tattered Hakkari Cape (20219, -0.80 DPS) [quest]; Chromatic Cloak (18509, -2.47 DPS, sim-verified) [crafted] |
| chest | Bloodsoul Breastplate (19690) | Blacksmithing [crafted] | 576.4 | yes | Stormshroud Armor (15056, -1.06 DPS, sim-verified) [crafted]; Savage Gladiator Chain (11726, -2.40 DPS) [dungeon]; Timbermaw Tunic (252484, -4.71 DPS) [crafted] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | 311.1 | yes | Deeprock Bracers (21184, -0.06 DPS, sim-verified) [quest]; Berserker Bracers (19578, -7.35 DPS) [rep]; Marshal's Plate Bracers (16481, -7.63 DPS) [pvp] |
| hands | Gauntlets of Heroism (21998) (or Gauntlets of Heroism (226861)) | Just Compensation [quest] | 353.8 | yes | Marshal's Plate Gauntlets (16484, +0.00 DPS) [vendor]; General's Plate Gauntlets (16548, +0.00 DPS) [vendor]; Gauntlets of Heroism (226861, +0.00 DPS, sim-verified) [quest] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 388.8 | yes | Zandalar Vindicator's Belt (19823, +0.00 DPS, sim-verified) [quest]; Highlander's Plate Girdle (20041, -1.24 DPS) [rep]; Highlander's Lamellar Girdle (20042, -1.49 DPS) [rep] |
| legs | Knight-Captain's Plate Leggings (227047) | Captain Dirgehammer [vendor] | 614.4 | yes | Marshal's Plate Legguards (16479, +0.00 DPS) [vendor]; General's Plate Leggings (16543, +0.00 DPS) [vendor]; General's Plate Leggings (231533, +0.00 DPS) [vendor] |
| feet | Conqueror's Greaves (21333) | Conqueror's Greaves [quest] | 97.9 | yes | Marshal's Plate Boots (16483, -0.53 DPS) [vendor]; General's Plate Boots (16545, -0.53 DPS) [vendor]; General's Plate Boots (231531, -0.53 DPS) [vendor] |
| finger1 | Signet of Unyielding Strength (21393) | Signet of Unyielding Strength [quest] | 326.6 | yes | Don Julio's Band (19325, -0.86 DPS) [rep]; Band of the Penitent (13217, -1.37 DPS) [quest]; Dragonslayer's Signet (18403, -1.37 DPS) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 307.2 | yes | Don Julio's Band (19325, +0.00 DPS, sim-verified) [rep]; Band of the Penitent (13217, -0.75 DPS) [quest]; Dragonslayer's Signet (18403, -0.75 DPS) [quest] |
| trinket1 | Onyxia Blood Talisman (18406) | Celebrating Good Times [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | 0.0 | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Glaive (234569, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; High Warlord's Street Sweeper (234561, +0.00 DPS) [vendor] |

**New at 60:** head: Lionheart Helm; neck: Fury of the Forgotten Swarm; back: Cloak of the Fallen God; chest: Bloodsoul Breastplate; wrist: Vambraces of the Sadist; hands: Gauntlets of Heroism; waist: Radiant Girdle of the Dawn; legs: Knight-Captain's Plate Leggings; feet: Conqueror's Greaves; finger1: Signet of Unyielding Strength; finger2: Band of Earthen Might; trinket1: Onyxia Blood Talisman; trinket2: Talisman of Arathor; main_hand: Arcanite Champion; ranged: The Purifier

No-known-source sample (15 of 2570, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (orc, 35300000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 32.2. Weights run: 1.6s. Verify run: 1.3s. 440 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.045, strength=2.084 ± 0.056, agility=not significant (0.000 ± 0.000), crit=not significant (0.000 ± 0.000), hit=not significant (0.000 ± 0.000), melee_haste=1.835 ± 0.076

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.8 | yes | Defender's Leather Hood (252447, -0.21 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.88 DPS) [crafted]; Shadow Goggles (4373, -0.88 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.3 | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.26 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.3 | yes | Subterranean Cape (14149, -0.08 DPS, sim-verified) [dungeon]; Grave Shroud (279865, -0.09 DPS) [quest]; Catacomb Cloak (279899, -0.10 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.8 | yes | Defender's Leather Armor (252434, -0.26 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.33 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.4 | yes | Cryptwalker Bracers (280095, -0.08 DPS, sim-verified) [quest]; Raptorcrest Bracers (270010, -0.18 DPS) [quest]; Runed Copper Bracers (2854, -0.26 DPS) [crafted] |
| hands | Gold-flecked Gloves (5195) | The Deadmines: Sneed [dungeon] | 14.6 | yes | Blackened Defias Gloves (10401, -0.08 DPS, sim-verified) [dungeon]; Foreman's Gloves (2167, -0.18 DPS) [world_drop]; Dagmire Gauntlets (6481, -0.18 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Cobrahn's Grasp (6460, -0.13 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.23 DPS) [world]; Support Girdle (1215, -0.32 DPS) [world] |
| legs | Veteran's Chain Leggings (250493) (or Defender's Leather Pants (252445), Totemic Leather Pants (252446)) | Blacksmithing [crafted] | 18.8 | yes | Defender's Leather Pants (252445, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, +0.00 DPS) [crafted]; Deepgrave Trousers (279900, -0.18 DPS) [quest] |
| feet | Veteran's Boots (250503) (or Guard's Boots (250504), Brawler's Leather Boots (252439), Defender's Leather Boots (252441), Totemic Leather Boots (252442)) | Blacksmithing [crafted] | 10.4 | yes | Guard's Boots (250504, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Defender's Leather Boots (252441, +0.00 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.3 | yes | The 1 Ring (8350, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Ring of the Shadow (1462, -0.35 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.3 | yes | The 1 Ring (8350, -0.21 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Ring of the Shadow (1462, -0.26 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 0.0 | yes | Smite's Mighty Hammer (7230, -0.20 DPS) [dungeon]; Living Root (6631, -0.28 DPS) [dungeon]; Hammerbone (270018, -2.23 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 299.6 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.87 DPS) [crafted]; Venomstrike (6469, -3.01 DPS) [dungeon] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Gold-flecked Gloves; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Veteran's Boots; finger1: Legionnaire's Band; finger2: Loop of Sacrifice; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Forsaken Greataxe; ranged: Ranger Bow

No-known-source sample (15 of 440, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore; 7956 Bronze Warhammer; 8178 Training Sword; 9602 Brushwood Blade; 9747 Simple Britches; 9748 Simple Robe

### Band 30 (orc, 35325210000000000-000000000000000000-000000000000000000)

Set DPS (verified): 71.4. Weights run: 1.7s. Verify run: 1.5s. 895 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.089, strength=1.991 ± 0.124, agility=not significant (0.008 ± 0.005), crit=0.171 ± 0.023, hit=not significant (0.000 ± 0.000), melee_haste=2.607 ± 0.360

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 25.9 | yes | Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Veteran's Chain Helm (250498, -0.16 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.19 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (20442, -0.68 DPS) [rep]; Erudite's Amulet (277204, -0.68 DPS) [quest]; Scout's Medallion (19537, -1.11 DPS, sim-verified) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 13.9 | yes | Mail Combat Spaulders (6404, +0.00 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.09 DPS) [crafted]; Elite Shoulders (4835, -0.10 DPS) [vendor] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Sergeant Major's Cape (16315, -0.10 DPS) [pvp]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop] |
| chest | Shining Silver Breastplate (2870) | Blacksmithing [crafted] | 27.9 | yes | Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.39 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.41 DPS, sim-verified) [crafted] |
| wrist | Pugilist Bracers (4438) | World drop [world_drop] | 15.9 | yes | Cultist's Armguards (270032, -0.29 DPS) [quest]; Grimtoll Wristguards (15459, -0.29 DPS) [quest]; Bands of Serra'kis (6902, -0.41 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 | yes | Warsong Gauntlets (16978, -0.05 DPS, sim-verified) [quest]; Bonefist Gauntlets (4465, -0.20 DPS) [world]; Heavy Earthen Gloves (7359, -0.29 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20207, -0.01 DPS) [rep]; Officer's Belt (250556, -0.20 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Golden Scale Leggings (3843, -0.20 DPS) [crafted]; Veteran's Silvered Chain Leggings (250523, -0.27 DPS, sim-verified) [crafted]; Juggernaut Leggings (6671, -0.30 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 14.0 | yes | Hard Gold Boots (250534, -0.02 DPS, sim-verified) [crafted]; Disjointed Shoes (277226, -0.10 DPS) [quest]; Glimmering Mail Greaves (4073, -0.10 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.0 | yes | Insurgent's Band (272067, -0.15 DPS) [vendor]; Ironspine's Eye (7686, -0.19 DPS) [dungeon]; Band of the Fist (17694, -0.20 DPS) [quest] |
| finger2 | Silverlaine's Family Seal (6321) | Shadowfang Keep: Baron Silverlaine [dungeon] | 10.0 | yes | Insurgent's Band (272067, -0.09 DPS, sim-verified) [vendor]; Ironspine's Eye (7686, -0.09 DPS) [dungeon]; Band of the Fist (17694, -0.10 DPS) [quest] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Corpsemaker (6687, -0.13 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.75 DPS) [vendor]; Morbid Dawn (7689, -21.89 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 378.6 | yes | Silver Star (3463, -0.66 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.30 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.23 DPS) [crafted] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Shining Silver Breastplate; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Trouncing Boots; finger1: Legionnaire's Band; finger2: Silverlaine's Family Seal; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Manual Crowd Pummeler; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 895, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7955 Copper Claymore

### Band 40 (orc, 35325213032010001-000000000000000000-000000000000000000)

Set DPS (verified): 97.4. Weights run: 2.1s. Verify run: 1.8s. 1333 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.304), strength=2.474 ± 0.448, agility=not significant (0.006 ± 0.006), crit=0.473 ± 0.059, hit=not significant (0.000 ± 0.000), melee_haste=3.320 ± 0.694

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 38.8 | yes | Hard Gold Coif (250537, -0.17 DPS) [crafted]; Tusken Helm (6686, -0.27 DPS) [dungeon]; Icemetal Barbute (10763, -0.64 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Ethereal Talisman (4430, -0.17 DPS, sim-verified) [quest]; Scout's Medallion (19536, -0.58 DPS) [rep]; Scout's Medallion (19537, -0.58 DPS) [rep] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 27.2 | yes | Imperial Leather Spaulders (4737, -0.20 DPS) [world_drop]; Wrangling Spaulders (15698, -0.31 DPS) [quest]; Shining Mithril Pauldrons (250541, -0.36 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16336) | Rank 9 [pvp] | 14.9 | yes | Wildhunter Cloak (16658, -0.20 DPS) [quest]; Sergeant Major's Cape (16315, -0.20 DPS) [pvp]; Wolfmaster Cape (6314, -0.86 DPS, sim-verified) [dungeon] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 34.7 | yes | Golden Scale Cuirass (3845, -0.00 DPS) [crafted]; Shining Mithril Breastplate (250540, -0.00 DPS) [crafted]; Shining Silver Breastplate (2870, -0.64 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Ravager's Armguards (14770, -0.11 DPS) [world]; Darkspear Armsplints (4132, -0.11 DPS) [quest]; Pugilist Bracers (4438, -0.42 DPS, sim-verified) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 39.6 | yes | Scarlet Gauntlets (10331, -0.41 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.47 DPS, sim-verified) [dungeon]; Seawolf Gloves (4509, -0.51 DPS) [quest] |
| waist | Defiler's Plate Girdle (20206) | The Defilers [rep] | 37.1 | yes | Defiler's Leather Girdle (20192, -0.29 DPS) [rep]; Scarlet Belt (10329, -0.31 DPS) [dungeon]; Tharg's Shoelace (9705, -0.39 DPS, sim-verified) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 52.0 | yes | Orcish War Leggings (7929, -0.58 DPS, sim-verified) [crafted]; Ornate Mithril Pants (7926, -0.92 DPS) [crafted]; Dual Reinforced Leggings (9625, -0.92 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 32.2 | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.39 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.41 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 19.8 | yes | Legionnaire's Band (19513, -0.20 DPS) [rep]; Silverlaine's Family Seal (6321, -0.31 DPS) [dungeon]; Insurgent's Band (272066, -0.32 DPS) [vendor] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 17.3 | yes | Insurgent's Band (272066, -0.22 DPS) [vendor]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Silverlaine's Family Seal (6321, -0.39 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | World drop [world_drop] | 488.3 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Rifle (17687, -0.34 DPS) [quest]; Mithril Blacksmith Hammer (285280, -0.37 DPS) [crafted] |

**New at 40:** head: Raging Berserker's Helm; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Suspicious Spare Part; trinket1: Ankh of Life; trinket2: Rune of Perfection; ranged: Sniper Rifle

No-known-source sample (15 of 1333, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (orc, 35325213032010001-050500000000000000-000000000000000000)

Set DPS (verified): 126.8. Weights run: 2.2s. Verify run: 1.9s. 1765 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.275), strength=1.574 ± 0.365, agility=not significant (0.284 ± 0.128), crit=6.161 ± 0.515, hit=not significant (0.000 ± 0.000), melee_haste=3.257 ± 0.606

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Plate Helm (220803) | Lady Palanseer [vendor] | 0.0 | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.37 DPS) [crafted]; Raging Berserker's Helm (7719, -1.83 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Woven Ivy Necklace (19159, +0.00 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -0.39 DPS) [quest]; Scout's Medallion (19535, -0.82 DPS) [rep] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 103.6 | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor]; Officer's Pauldrons (250576, -6.26 DPS) [crafted]; Wyrmslayer Spaulders (13066, -6.39 DPS) [world_drop] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 14.2 | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.32 DPS) [dungeon]; Battlehard Cape (11858, -0.32 DPS) [quest] |
| chest | Knight's Plate Hauberk (220794) (or Stone Guard's Plate Armor (220801)) | Captain Dirgehammer [vendor] | 109.9 | yes | Stone Guard's Plate Armor (220801, +0.00 DPS, sim-verified) [vendor]; Ornate Mithril Breastplate (7935, -1.83 DPS) [crafted]; Warforged Chestplate (11195, -5.59 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 | yes | Branded Leather Bracers (19508, -0.39 DPS, sim-verified) [dungeon]; Officer's Wristguards (250581, -0.70 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.92 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 106.3 | yes | Dragonscale Gauntlets (8347, -0.93 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.55 DPS) [crafted]; Ornate Mithril Gloves (7927, -1.55 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 106.3 | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Plate Girdle (20205, -0.33 DPS) [rep]; Defiler's Chain Girdle (20153, -0.93 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 0.0 | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.49 DPS, sim-verified) [crafted]; Scarlet Leggings (10330, -5.81 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 23.6 | yes | Officer's Boots (250546, -0.09 DPS) [crafted]; Officer's Sabatons (250561, -0.17 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -0.20 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Legionnaire's Band (19511, -0.44 DPS) [rep]; Insurgent's Band (272065, -0.70 DPS) [vendor]; Legionnaire's Band (19512, -0.71 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Legionnaire's Band (19511, -0.03 DPS, sim-verified) [rep]; Insurgent's Band (272065, -0.39 DPS) [vendor]; Legionnaire's Band (19512, -0.40 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Guardian Talisman (1490, -3.26 DPS) [quest]; Ankh of Life (1713, -3.26 DPS) [world_drop]; Blazing Emblem (2802, -3.26 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | 0.0 | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 0.0 | yes | Dark Iron Rifle (16004, -1.00 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -4.27 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -4.80 DPS) [world_drop] |

**New at 50:** head: Blood Guard's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Knight's Plate Hauberk; wrist: Bracers of the Stone Princess; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Fiery War Axe; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1765, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (orc, 35325213032010001-050500000000000000-500500000000000000)

Set DPS (verified): 198.0. Weights run: 2.2s. Verify run: 1.9s. 2573 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=not significant (1.000 ± 0.801), strength=not significant (3.886 ± 1.021), agility=not significant (0.959 ± 0.373), crit=20.276 ± 1.619, hit=not significant (0.000 ± 0.000), melee_haste=9.244 ± 1.838

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 637.7 | yes | Bloodvine Lens (19998, -2.24 DPS) [crafted]; Ragefury Eyepatch (11735, -6.85 DPS, sim-verified) [dungeon]; Field Marshal's Plate Helm (16478, -7.84 DPS) [vendor] |
| neck | Fury of the Forgotten Swarm (21809) | World drop [world_drop] | 0.0 | yes | Medallion of the Dawn (22659, -0.23 DPS) [quest]; Onyxia Tooth Pendant (18404, -0.78 DPS) [quest]; Amulet of the Darkmoon (19491, -8.26 DPS) [quest] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) (or Blood Guard's Plate Pauldrons (220796)) | Captain Dirgehammer [vendor] | 326.6 | yes | Champion's Plate Shoulders (23243, +0.00 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (23315, +0.00 DPS) [vendor]; Blood Guard's Plate Pauldrons (220796, +0.00 DPS, sim-verified) [vendor] |
| back | Cloak of the Fallen God (21710) | The Savior of Kalimdor [quest] | 0.0 | yes | Drape of Unyielding Strength (21394, -0.02 DPS) [quest]; Tattered Hakkari Cape (20219, -0.80 DPS) [quest]; Chromatic Cloak (18509, -2.20 DPS, sim-verified) [crafted] |
| chest | Bloodsoul Breastplate (19690) | Blacksmithing [crafted] | 576.4 | yes | Stormshroud Armor (15056, -0.92 DPS, sim-verified) [crafted]; Savage Gladiator Chain (11726, -2.40 DPS) [dungeon]; Timbermaw Tunic (252484, -4.71 DPS) [crafted] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | 311.1 | yes | Deeprock Bracers (21184, +0.00 DPS, sim-verified) [quest]; Berserker Bracers (19578, -7.35 DPS) [rep]; Marshal's Plate Bracers (16481, -7.63 DPS) [pvp] |
| hands | Gauntlets of Heroism (21998) (or Gauntlets of Heroism (226861)) | Just Compensation [quest] | 353.8 | yes | Marshal's Plate Gauntlets (16484, +0.00 DPS) [vendor]; General's Plate Gauntlets (16548, +0.00 DPS) [vendor]; Gauntlets of Heroism (226861, +0.00 DPS, sim-verified) [quest] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 388.8 | yes | Zandalar Vindicator's Belt (19823, -0.85 DPS, sim-verified) [quest]; Defiler's Plate Girdle (20204, -1.24 DPS) [rep]; Defiler's Plate Girdle (20205, -2.12 DPS) [rep] |
| legs | Knight-Captain's Plate Leggings (227047) | Captain Dirgehammer [vendor] | 614.4 | yes | Marshal's Plate Legguards (16479, +0.00 DPS) [vendor]; General's Plate Leggings (16543, +0.00 DPS) [vendor]; General's Plate Leggings (231533, +0.00 DPS) [vendor] |
| feet | Conqueror's Greaves (21333) | Conqueror's Greaves [quest] | 97.9 | yes | Marshal's Plate Boots (16483, -0.53 DPS) [vendor]; General's Plate Boots (16545, -0.53 DPS) [vendor]; General's Plate Boots (231531, -0.53 DPS) [vendor] |
| finger1 | Signet of Unyielding Strength (21393) | Signet of Unyielding Strength [quest] | 326.6 | yes | Don Julio's Band (19325, -0.86 DPS) [rep]; Band of the Penitent (13217, -1.37 DPS) [quest]; Dragonslayer's Signet (18403, -1.37 DPS) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 307.2 | yes | Don Julio's Band (19325, -0.49 DPS, sim-verified) [rep]; Band of the Penitent (13217, -0.75 DPS) [quest]; Dragonslayer's Signet (18403, -0.75 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Guardian Talisman (1490, -1.34 DPS) [quest]; Ankh of Life (1713, -1.34 DPS) [world_drop]; Blazing Emblem (2802, -1.34 DPS) [world_drop] |
| trinket2 | Onyxia Blood Talisman (18406) | For All To See [quest] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| main_hand | Nightfall (19169) | Blacksmithing [crafted] | 0.0 | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Glaive (234569, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; High Warlord's Street Sweeper (234561, +0.00 DPS) [vendor] |

**New at 60:** head: Lionheart Helm; neck: Fury of the Forgotten Swarm; back: Cloak of the Fallen God; chest: Bloodsoul Breastplate; wrist: Vambraces of the Sadist; hands: Gauntlets of Heroism; waist: Radiant Girdle of the Dawn; legs: Knight-Captain's Plate Leggings; feet: Conqueror's Greaves; finger1: Signet of Unyielding Strength; finger2: Band of Earthen Might; trinket2: Onyxia Blood Talisman; main_hand: Nightfall; ranged: The Purifier

No-known-source sample (15 of 2573, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

