# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 46.7. Weights run: 5.6s. Verify run: 16.3s. 312 eligible items had no known source.

Race sweep: human 46.7 ± 0.0, dwarf 46.1 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.319 ± 0.020, crit=0.796 ± 0.018 per rating point (14 rating = 1%, 11.139 per %), hit=1.450 ± 0.050 per rating point (10 rating = 1%, 14.503 per %), melee_haste=5.157 ± 0.678, intellect=1.066 ± 0.026, spell_power=2.415 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 26.6 attack_power points (1.14 DPS) | yes | Totemic Leather Hood (252448, -0.21 DPS) [crafted]; Veteran's Silvered Chain Helm (250528, -0.28 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.41 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 7.9 attack_power points (0.34 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.60 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.7 attack_power points (0.93 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.51 DPS) [crafted]; Serpent's Shoulders (5404, -0.64 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.07 DPS) [crafted]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop]; Dark Leather Cloak (2316, -0.07 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 31.7 attack_power points (1.36 DPS) | yes | Crusader's Chain Shirt (250492, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.30 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.50 DPS) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.00 DPS) [quest]; Cryptwalker Bracers (280095, -0.09 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 24.8 attack_power points (1.06 DPS) | yes | Totemic Leather Gloves (252497, -0.20 DPS) [crafted]; Crusader's Gloves (250512, -0.29 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.36 DPS) [crafted] |
| waist | Crusader's Chain Belt (250517) | Blacksmithing [crafted] | 18.1 attack_power points (0.77 DPS) | yes | Blackened Defias Belt (10403, -0.00 DPS) [dungeon]; Cobrahn's Grasp (6460, -0.01 DPS) [dungeon]; Totemic Leather Belt (252431, -0.02 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 39.2 attack_power points (1.68 DPS) | yes | Crusader's Chain Leggings (250497, -0.24 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.37 DPS) [crafted]; Trapper's Leather Pants (252501, -0.47 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 24.5 attack_power points (1.05 DPS) | yes | Crusader's Boots (250507, -0.09 DPS) [crafted]; Spidersilk Boots (4320, -0.14 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.20 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 14.2 attack_power points (0.61 DPS) | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Demon Band (12054, -0.27 DPS) [world_drop]; Protector's Band (20439, -0.51 DPS, sim-verified) [rep] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Protector's Band (20439, -0.10 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Demon Band (12054, -0.17 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sword of the Fallen (286977, -2.89 DPS) [dungeon]; Duskbringer (2205, -3.03 DPS) [dungeon]; Smite's Mighty Hammer (7230, -4.01 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Totemic Leather Armor; wrist: Patterned Bronze Bracers; hands: Serpent Gloves; waist: Crusader's Chain Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Verigan's Fist

No-known-source sample (15 of 312, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 30 (human, 00000000000000000-00000000000000000-05233201001300100)

Set DPS (verified): 85.1. Weights run: 6.2s. Verify run: 13.2s. 537 eligible items had no known source.

Race sweep: human 85.1 ± 0.1, dwarf 83.7 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.802 ± 0.029, crit=1.414 ± 0.031 per rating point (14 rating = 1%, 19.801 per %), hit=1.764 ± 0.073 per rating point (10 rating = 1%, 17.644 per %), melee_haste=7.617 ± 0.971, intellect=1.344 ± 0.035, spell_power=2.604 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 58.5 attack_power points (2.90 DPS) | yes | Totemic Leather Helm (252456, -1.35 DPS) [crafted]; Enchanter's Cowl (4322, -1.46 DPS) [crafted]; Barbaric Iron Helm (7915, -1.65 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 20.6 attack_power points (1.02 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Sentinel's Medallion (19541, -0.31 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 35.5 attack_power points (1.77 DPS) | yes | Forest Tracker Epaulets (2278, -0.28 DPS) [world_drop]; Death Speaker Mantle (6685, -0.33 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.52 DPS) [world_drop] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.10 DPS) [world_drop]; Tigerstrike Mantle (13108, -0.31 DPS) [world_drop]; Vine Pruner's Cloak (279835, -0.38 DPS) [quest] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 43.2 attack_power points (2.15 DPS) | yes | Green Silk Armor (7065, -0.12 DPS) [crafted]; Beguiler Robes (7728, -0.18 DPS) [dungeon]; Totemic Leather Tunic (252451, -0.59 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 23.4 attack_power points (1.16 DPS) | yes | Yorgen Bracers (13012, -0.30 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.89 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 29.0 attack_power points (1.44 DPS) | yes | Mail Combat Gauntlets (4075, -0.29 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.31 DPS) [crafted]; Insignia Gloves (6408, -0.37 DPS, sim-verified) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 43.2 attack_power points (2.15 DPS) | yes | Skirmisher's Leather Belt (252461, -0.40 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -0.45 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.45 DPS) [crafted] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 47.9 attack_power points (2.38 DPS) | yes | Ferine Leggings (6690, -0.28 DPS) [dungeon]; Totemic Leather Pants (252446, -0.31 DPS) [crafted]; Totemic Leather Leggings (252458, -0.32 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 32.2 attack_power points (1.60 DPS) | yes | Totemic Leather Boots (252442, -0.33 DPS) [crafted]; Disjointed Shoes (277226, -0.36 DPS) [quest]; Gilded Slippers (254001, -0.53 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 25.0 attack_power points (1.24 DPS) | yes | Thunderbrow Ring (13097, -0.18 DPS) [world_drop]; Lorekeeper's Ring (19525, -0.34 DPS) [rep]; Seal of Wrynn (2933, -0.41 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 22.8 attack_power points (1.13 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Lorekeeper's Ring (19525, -0.23 DPS) [rep]; Seal of Wrynn (2933, -0.30 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.38 DPS) [world]; Corpsemaker (6687, -0.58 DPS) [dungeon]; Viscous Hammer (13045, -6.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Sergeant Major's Cape; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 537, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 136.2. Weights run: 6.1s. Verify run: 16.5s. 808 eligible items had no known source.

Race sweep: human 136.2 ± 0.2, dwarf 133.8 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=2.566 ± 0.053, crit=2.565 ± 0.061 per rating point (14 rating = 1%, 35.916 per %), hit=2.070 ± 0.089 per rating point (10 rating = 1%, 20.702 per %), melee_haste=12.104 ± 1.356, intellect=2.640 ± 0.048, spell_power=2.519 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 80.6 attack_power points (4.19 DPS) | yes | Miner's Hat of the Deep (9429, -0.29 DPS) [dungeon]; Papal Fez (9431, -0.81 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.95 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 38.8 attack_power points (2.02 DPS) | yes | Sentinel's Medallion (19540, -0.55 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.57 DPS) [world]; Kaleidoscope Chain (13084, -1.06 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 51.9 attack_power points (2.70 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS) [crafted]; Shining Mithril Pauldrons (250541, -0.20 DPS) [crafted]; Bloodmage Mantle (7684, -0.29 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.79 DPS) [crafted]; Guardian Cloak (5965, -0.79 DPS) [crafted]; Expeditionary's Cape (271922, -0.79 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.11 DPS) [crafted]; Robe of Power (7054, -0.22 DPS) [crafted]; Icemail Jerkin (1981, -0.53 DPS, sim-verified) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 33.8 attack_power points (1.76 DPS) | yes | Guardian Leather Bracers (4260, -0.15 DPS) [crafted]; Windchaser Cuffs (14429, -0.26 DPS) [world_drop]; Warden's Wristbands (14600, -0.26 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, -2.22 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 57.2 attack_power points (2.98 DPS) | yes | Highlander's Cloth Girdle (20098, -0.59 DPS) [rep]; Prowler's Leather Belt (252459, -0.67 DPS) [crafted]; Skirmisher's Leather Belt (252461, -2.08 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 66.9 attack_power points (3.48 DPS) | yes | Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Totemic Leather Leggings (252458, -0.66 DPS) [crafted]; Ferine Leggings (6690, -0.71 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 60.5 attack_power points (3.15 DPS) | yes | Skirmisher's Mail Boots (252564, -0.11 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.36 DPS) [crafted]; Justicar's Boots (250550, -0.44 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 38.4 attack_power points (2.00 DPS) | yes | Snake Hoop (6750, -0.25 DPS) [quest]; Ring of the Underwood (2951, -0.45 DPS) [world_drop]; Falcon's Hook (7552, -0.48 DPS) [dungeon] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 36.8 attack_power points (1.92 DPS) | yes | Snake Hoop (6750, -0.17 DPS) [quest]; Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, -0.15 DPS) [vendor]; Nightblade (1982, -1.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Sergeant Major's Cape; chest: Robe of the Magi; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Protector's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 808, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 195.6. Weights run: 6.4s. Verify run: 22.0s. 1046 eligible items had no known source.

Race sweep: human 195.6 ± 0.2, dwarf 193.2 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=3.441 ± 0.077, crit=4.160 ± 0.100 per rating point (14 rating = 1%, 58.244 per %), hit=3.219 ± 0.184 per rating point (10 rating = 1%, 32.192 per %), melee_haste=18.442 ± 2.281, intellect=4.328 ± 0.090, spell_power=2.944 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 142.5 attack_power points (7.68 DPS) | yes | Eye of Theradras (17715, -0.11 DPS) [dungeon]; Chief Architect's Monocle (11839, -0.35 DPS) [dungeon]; Soothsayer's Headdress (17740, -0.53 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 60.6 attack_power points (3.27 DPS) | yes | Skibi's Pendant (13089, -0.26 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -0.29 DPS) [dungeon]; Gemshard Heart (17707, -0.37 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 130.7 attack_power points (7.05 DPS) | yes | Lead Surveyor's Mantle (11842, -1.40 DPS) [dungeon]; Knight-Lieutenant's Lamellar Pauldrons (220818, -1.51 DPS) [vendor]; Rotgrip Mantle (17732, -1.65 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.31 DPS) [dungeon]; Runecloth Cloak (13860, -0.33 DPS) [crafted]; Expeditionary's Cape (271921, -0.36 DPS) [rep] |
| chest | Robes of Insight (940) | World drop [world_drop] | 136.3 attack_power points (7.35 DPS) | yes | Spiderfang Carapace (11633, -0.59 DPS) [dungeon]; Knight's Lamellar Chestplate (220815, -0.86 DPS) [vendor]; Icemail Jerkin (1981, -1.16 DPS) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 76.7 attack_power points (4.14 DPS) | yes | Runic Leather Bracers (15092, -0.21 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.60 DPS) [crafted]; Skirmisher's Mail Bracers (252572, -0.63 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Lamellar Gauntlets (220817, +0.00 DPS) [vendor]; Skirmisher's Mail Gauntlets (252586, +0.00 DPS) [crafted]; Stonerender Gauntlets (17007, -3.53 DPS, sim-verified) [world_drop] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 117.1 attack_power points (6.31 DPS) | yes | Ban'thok Sash (11662, -0.34 DPS) [dungeon]; Highlander's Cloth Girdle (20097, -0.58 DPS) [rep]; Highlander's Lizardhide Girdle (20103, -0.84 DPS) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 135.6 attack_power points (7.31 DPS) | yes | Knight's Plate Leggings (220797, -0.89 DPS) [vendor]; Imperial Red Pants (8251, -1.03 DPS) [world_drop]; Knight's Lamellar Legplates (220816, -2.85 DPS, sim-verified) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 107.9 attack_power points (5.82 DPS) | yes | Skycaller's Mail Sabatons (252577, -0.87 DPS) [crafted]; Fleetfoot Greaves (11627, -0.89 DPS) [dungeon]; Skycaller's Leather Boots (252471, -2.52 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 75.7 attack_power points (4.08 DPS) | yes | Cyclopean Band (11824, -0.54 DPS) [dungeon]; Chivalrous Signet (20505, -0.66 DPS) [quest]; Woodseed Hoop (17768, -0.79 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 68.0 attack_power points (3.67 DPS) | yes | Chivalrous Signet (20505, -0.25 DPS) [quest]; Woodseed Hoop (17768, -0.38 DPS) [quest]; Cyclopean Band (11824, -1.29 DPS, sim-verified) [dungeon] |
| trinket1 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.82 DPS) [crafted]; Uther's Strength (11302, -1.43 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -3.15 DPS, sim-verified) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -2.36 DPS) [world_drop]; The Rockpounder (9413, -2.78 DPS) [dungeon]; Stoneraven (13059, -4.08 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; waist: Highlander's Lamellar Girdle; legs: Infernal Trickster Leggings; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Sanctified Orb; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 1046, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 344.2. Weights run: 6.4s. Verify run: 83.8s. 2409 eligible items had no known source.

Race sweep: human 344.2 ± 0.3, dwarf 341.8 ± 0.3 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.761 ± 0.042, crit=2.676 ± 0.060 per rating point (14 rating = 1%, 37.467 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=18.303 ± 2.481, intellect=5.642 ± 0.120, spell_power=3.184 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 288.8 attack_power points (20.20 DPS) | yes | Field Marshal's Lamellar Faceguard (16474, -0.89 DPS) [vendor]; Field Marshal's Lamellar Headguard (231648, -3.60 DPS) [vendor]; Magister's Crown (16686, -7.66 DPS, sim-verified) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 142.2 attack_power points (9.94 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -1.51 DPS) [world_drop]; Archlight Talisman (15856, -3.33 DPS) [quest]; Tooth of Gnarr (13141, -4.42 DPS) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 203.3 attack_power points (14.21 DPS) | yes | Theramore Shoulderpads (271953, -0.89 DPS) [rep]; Theramore Shoulders (271954, -0.89 DPS) [rep]; Field Marshal's Lamellar Shoulders (231651, -2.39 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Shroud of the Exile (15421, -0.10 DPS) [quest]; Expeditionary's Cape (271907, -0.41 DPS) [rep]; Royal Tribunal Cloak (13376, -3.68 DPS, sim-verified) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 196.2 attack_power points (13.72 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Vest of Elements (16666, -0.71 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 115.2 attack_power points (8.05 DPS) | yes | Funeral Cuffs (12626, -1.20 DPS) [dungeon]; Aristocratic Cuffs (12546, -1.25 DPS) [dungeon]; Manacle Cuffs (11962, -2.18 DPS, sim-verified) [quest] |
| hands | Theramore Handguards (271950) | Theramore Expeditionary Force [rep] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Lamellar Gloves (231643, -0.12 DPS) [pvp]; Theramore Handwraps (271945, -0.68 DPS, sim-verified) [rep] |
| waist | Crystallized Girdle (12606) | Blackrock Spire: Solakar Flamewreath [dungeon] | 148.6 attack_power points (10.39 DPS) | yes | Magister's Belt (16685, -0.55 DPS) [dungeon]; Highlander's Mail Girdle (20044, -1.06 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, -1.06 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legplates (16475, +0.00 DPS) [vendor]; Marshal's Lamellar Leggings (231647, +0.00 DPS) [vendor]; Red Dragonscale Leggings (252603, -5.27 DPS, sim-verified) [crafted] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 147.6 attack_power points (10.32 DPS) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor] |
| finger1 | Seal of Rivendare (13345) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.72 DPS) [quest]; Deep Rooted Ring (19109, -1.87 DPS) [world]; Naglering (11669, -6.47 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Blood of the Martyr (17045, -0.15 DPS) [quest]; Deep Rooted Ring (19109, -0.15 DPS) [world] |
| trinket1 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -1.78 DPS) [vendor]; Uther's Strength (11302, -2.00 DPS) [world_drop] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS) [vendor]; Weakness Analyzer (272438, -1.91 DPS, sim-verified) [vendor] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -22.22 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, -4.61 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Theramore Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Theramore Handguards; waist: Crystallized Girdle; legs: Ironfeather Leggings; feet: Dragonrider Boots; finger1: Seal of Rivendare; finger2: Maiden's Circle; trinket2: Blackhand's Breadth; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2409, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60, raid preset (human, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 850.7. Weights run: 6.1s. Verify run: 96.5s. 2409 eligible items had no known source.

Race sweep: human 850.7 ± 1.1, dwarf 843.3 ± 1.1 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=1.628 ± 0.031, crit=2.200 ± 0.043 per rating point (14 rating = 1%, 30.800 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=not significant (13.719 ± 3.721), intellect=1.302 ± 0.056, spell_power=1.685 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Lamellar Faceguard (16474, +0.00 DPS) [vendor]; Field Marshal's Lamellar Headguard (231648, +0.00 DPS) [vendor]; Black Dragonscale Helm (252605, +0.00 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Mark of Fordring (15411, +0.00 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Lieutenant Commander's Lamellar Shoulders (23277, +0.00 DPS) [vendor]; Field Marshal's Lamellar Shoulders (231651, +0.00 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, +0.00 DPS) [dungeon]; Sergeant Major's Cape (16337, +0.00 DPS) [pvp]; Cloak of the Honor Guard (20073, +0.00 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 140.0 attack_power points (28.21 DPS) | yes | Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Knight-Captain's Lamellar Breastplate (227142, -5.71 DPS) [vendor]; Savage Gladiator Chain (11726, -6.50 DPS) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Windtalker's Wristguards (19582, -1.84 DPS) [rep] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Theramore Gauntlets (271940, +0.00 DPS) [rep] |
| waist | Highlander's Lamellar Girdle (20042) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Vindicator's Belt (272404, +0.00 DPS) [vendor]; Might of the Timbermaw (19044, -1.31 DPS) [crafted]; Highlander's Plate Girdle (20041, -1.65 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legplates (16475, +0.00 DPS) [vendor]; Marshal's Lamellar Leggings (231647, +0.00 DPS) [vendor]; Champion's Legplates (250598, +0.00 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Sabatons (227146, +0.00 DPS) [vendor]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magni's Will (12548, +0.00 DPS) [quest]; Protector's Band (19514, +0.00 DPS) [rep]; Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magni's Will (12548, +0.00 DPS) [quest]; Protector's Band (19514, +0.00 DPS) [rep]; Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanctified Orb (20512, +0.00 DPS) [quest]; Uther's Strength (11302, -0.38 DPS) [world_drop]; Second Wind (11819, -0.72 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanctified Orb (20512, -1.02 DPS) [quest]; Serenity Field (272439, -3.74 DPS) [vendor]; Uther's Strength (11302, -4.08 DPS) [world_drop] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Lionheart Helm; neck: Skibi's Pendant; shoulder: Theramore Pauldrons; back: Arcanoweave Cloak; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Highlander's Lamellar Girdle; legs: Sentinel's Chain Leggings; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Weakness Analyzer; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 2409, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (undead, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 43.6. Weights run: 5.6s. Verify run: 14.7s. 299 eligible items had no known source.

Race sweep: undead 43.6 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.319 ± 0.020, crit=0.796 ± 0.018 per rating point (14 rating = 1%, 11.139 per %), hit=1.450 ± 0.050 per rating point (10 rating = 1%, 14.503 per %), melee_haste=5.157 ± 0.678, intellect=1.066 ± 0.026, spell_power=2.415 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 26.6 attack_power points (1.14 DPS) | yes | Totemic Leather Hood (252448, -0.21 DPS) [crafted]; Veteran's Silvered Chain Helm (250528, -0.28 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.41 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 7.9 attack_power points (0.34 DPS) | yes | Erudite's Amulet (277204, -0.23 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.7 attack_power points (0.93 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.29 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.51 DPS) [crafted]; Serpent's Shoulders (5404, -0.64 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.07 DPS) [crafted]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop]; Dark Leather Cloak (2316, -0.07 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 31.7 attack_power points (1.36 DPS) | yes | Crusader's Chain Shirt (250492, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.30 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.50 DPS) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.12 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 24.8 attack_power points (1.06 DPS) | yes | Totemic Leather Gloves (252497, -0.20 DPS) [crafted]; Crusader's Gloves (250512, -0.29 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.36 DPS) [crafted] |
| waist | Crusader's Chain Belt (250517) | Blacksmithing [crafted] | 18.1 attack_power points (0.77 DPS) | yes | Blackened Defias Belt (10403, -0.00 DPS) [dungeon]; Cobrahn's Grasp (6460, -0.01 DPS) [dungeon]; Totemic Leather Belt (252431, -0.02 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 39.2 attack_power points (1.68 DPS) | yes | Crusader's Chain Leggings (250497, -0.34 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.37 DPS) [crafted]; Trapper's Leather Pants (252501, -0.47 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 24.5 attack_power points (1.05 DPS) | yes | Crusader's Boots (250507, -0.09 DPS) [crafted]; Spidersilk Boots (4320, -0.14 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.20 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.3 attack_power points (0.57 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Demon Band (12054, -0.23 DPS) [world_drop] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 12.1 attack_power points (0.52 DPS) | yes | Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Demon Band (12054, -0.17 DPS) [world_drop]; Loop of Sacrifice (281673, -0.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Hammerbone (270018, -0.12 DPS) [quest]; Sword of the Fallen (286977, -0.23 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.27 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Totemic Leather Armor; wrist: Patterned Bronze Bracers; hands: Serpent Gloves; waist: Crusader's Chain Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Legionnaire's Band; finger2: Advisor's Ring; main_hand: Forsaken Greataxe

No-known-source sample (15 of 299, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (undead, 00000000000000000-00000000000000000-05233201001300100)

Set DPS (verified): 85.7. Weights run: 6.2s. Verify run: 14.3s. 524 eligible items had no known source.

Race sweep: undead 85.7 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.802 ± 0.029, crit=1.414 ± 0.031 per rating point (14 rating = 1%, 19.801 per %), hit=1.764 ± 0.073 per rating point (10 rating = 1%, 17.644 per %), melee_haste=7.617 ± 0.971, intellect=1.344 ± 0.035, spell_power=2.604 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 58.5 attack_power points (2.90 DPS) | yes | Totemic Leather Helm (252456, -1.35 DPS) [crafted]; Enchanter's Cowl (4322, -1.46 DPS) [crafted]; Barbaric Iron Helm (7915, -1.58 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 20.6 attack_power points (1.02 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Scout's Medallion (19537, -0.31 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 35.5 attack_power points (1.77 DPS) | yes | Death Speaker Mantle (6685, -0.25 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.28 DPS) [world_drop]; Magician's Mantle (12998, -0.52 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.21 DPS) [world_drop]; Hillman's Cloak (3719, -0.28 DPS) [crafted]; Windsong Drape (15468, -0.28 DPS) [quest] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 43.2 attack_power points (2.15 DPS) | yes | Totemic Leather Tunic (252451, -0.11 DPS) [crafted]; Green Silk Armor (7065, -0.12 DPS) [crafted]; Beguiler Robes (7728, -0.18 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 23.4 attack_power points (1.16 DPS) | yes | Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Yorgen Bracers (13012, -0.30 DPS) [world_drop]; Undead Knight's Bracers (251965, -0.87 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 29.0 attack_power points (1.44 DPS) | yes | Oilrag Handwraps (16741, -0.14 DPS) [quest]; Insignia Gloves (6408, -0.21 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.29 DPS) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 43.2 attack_power points (2.15 DPS) | yes | Skirmisher's Leather Belt (252461, -0.08 DPS) [crafted]; Prowler's Leather Belt (252459, -0.45 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.45 DPS) [crafted] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 47.9 attack_power points (2.38 DPS) | yes | Totemic Leather Leggings (252458, -0.05 DPS) [crafted]; Ferine Leggings (6690, -0.28 DPS) [dungeon]; Totemic Leather Pants (252446, -0.31 DPS) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 32.2 attack_power points (1.60 DPS) | yes | Gilded Slippers (254001, -0.23 DPS) [crafted]; Totemic Leather Boots (252442, -0.33 DPS) [crafted]; Acidic Walkers (9454, -0.42 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 25.0 attack_power points (1.24 DPS) | yes | Thunderbrow Ring (13097, -0.18 DPS) [world_drop]; Advisor's Ring (19521, -0.34 DPS) [rep]; Runemender's Seal (273031, -0.45 DPS) [world] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 22.8 attack_power points (1.13 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Advisor's Ring (19521, -0.23 DPS) [rep]; Runemender's Seal (273031, -0.34 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.38 DPS) [world]; Corpsemaker (6687, -0.58 DPS) [dungeon]; Viscous Hammer (13045, -6.37 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Hawkeye's Cloak; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 524, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 134.5. Weights run: 6.1s. Verify run: 20.7s. 792 eligible items had no known source.

Race sweep: undead 134.5 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=2.566 ± 0.053, crit=2.565 ± 0.061 per rating point (14 rating = 1%, 35.916 per %), hit=2.070 ± 0.089 per rating point (10 rating = 1%, 20.702 per %), melee_haste=12.104 ± 1.356, intellect=2.640 ± 0.048, spell_power=2.519 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 80.6 attack_power points (4.19 DPS) | yes | Miner's Hat of the Deep (9429, -0.35 DPS, sim-verified) [dungeon]; Papal Fez (9431, -0.81 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.95 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 38.8 attack_power points (2.02 DPS) | yes | Scout's Medallion (19536, -0.55 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.57 DPS) [world]; Kaleidoscope Chain (13084, -0.89 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 51.9 attack_power points (2.70 DPS) | yes | Shining Mithril Pauldrons (250541, -0.20 DPS) [crafted]; Bloodmage Mantle (7684, -0.29 DPS) [dungeon]; Green Silken Shoulders (7057, -0.32 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.04 DPS) [crafted]; Hawkeye's Cloak (14593, -0.26 DPS) [world_drop]; Long Silken Cloak (4326, -0.93 DPS, sim-verified) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 82.3 attack_power points (4.28 DPS) | yes | Robe of the Magi (1716, -0.57 DPS) [world_drop]; Dreamweave Vest (10021, -0.69 DPS) [crafted]; Robe of Power (7054, -0.80 DPS) [crafted] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 33.8 attack_power points (1.76 DPS) | yes | Radiant Silver Bracers (4545, -0.14 DPS) [quest]; Guardian Leather Bracers (4260, -0.15 DPS) [crafted]; Windchaser Cuffs (14429, -0.26 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, -1.62 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 57.2 attack_power points (2.98 DPS) | yes | Defiler's Cloth Girdle (20166, -0.59 DPS) [rep]; Prowler's Leather Belt (252459, -0.67 DPS) [crafted]; Skirmisher's Leather Belt (252461, -1.58 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 66.9 attack_power points (3.48 DPS) | yes | Ferine Leggings (6690, -0.48 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Totemic Leather Leggings (252458, -0.66 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 60.5 attack_power points (3.15 DPS) | yes | Skirmisher's Mail Boots (252564, -0.11 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.36 DPS) [crafted]; Justicar's Boots (250550, -0.44 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 38.4 attack_power points (2.00 DPS) | yes | Snake Hoop (6750, -0.25 DPS) [quest]; Ring of the Underwood (2951, -0.45 DPS) [world_drop]; Falcon's Hook (7552, -0.48 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 36.8 attack_power points (1.92 DPS) | yes | Snake Hoop (6750, -0.17 DPS) [quest]; Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, -0.15 DPS) [vendor]; Nightblade (1982, -1.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Legionnaire's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 792, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 191.3. Weights run: 6.4s. Verify run: 25.3s. 1046 eligible items had no known source.

Race sweep: undead 191.3 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=3.441 ± 0.077, crit=4.160 ± 0.100 per rating point (14 rating = 1%, 58.244 per %), hit=3.219 ± 0.184 per rating point (10 rating = 1%, 32.192 per %), melee_haste=18.442 ± 2.281, intellect=4.328 ± 0.090, spell_power=2.944 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Chief Architect's Monocle (11839, -0.24 DPS) [dungeon]; Red Mageweave Headband (10033, -0.40 DPS, sim-verified) [crafted]; Soothsayer's Headdress (17740, -0.42 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Eye of Mordresh (10769, -0.03 DPS) [dungeon]; Gemshard Heart (17707, -0.11 DPS) [dungeon]; Horizon Choker (13085, -0.53 DPS, sim-verified) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 130.7 attack_power points (7.05 DPS) | yes | Rotgrip Mantle (17732, -0.96 DPS, sim-verified) [dungeon]; Lead Surveyor's Mantle (11842, -1.40 DPS) [dungeon]; Rockshard Pauldrons (9411, -1.75 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.38 DPS) [dungeon]; Blackveil Cape (11626, -0.70 DPS) [dungeon]; Runecloth Cloak (13860, -0.71 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 136.3 attack_power points (7.35 DPS) | yes | Spiderfang Carapace (11633, -0.59 DPS) [dungeon]; Icemail Jerkin (1981, -1.16 DPS) [world_drop]; Acumen Robes (17775, -1.62 DPS) [quest] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 76.7 attack_power points (4.14 DPS) | yes | Runic Leather Bracers (15092, -0.21 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.60 DPS) [crafted]; Skirmisher's Mail Bracers (252572, -0.63 DPS) [crafted] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Gauntlets (252547, -0.70 DPS) [crafted]; Stalker's Mail Gauntlets (252584, -0.74 DPS) [crafted]; Skirmisher's Mail Gauntlets (252586, -1.11 DPS, sim-verified) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | 117.1 attack_power points (6.31 DPS) | yes | Ban'thok Sash (11662, -0.34 DPS) [dungeon]; Defiler's Cloth Girdle (20165, -0.58 DPS) [rep]; Defiler's Lizardhide Girdle (20174, -0.84 DPS) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 135.6 attack_power points (7.31 DPS) | yes | Stone Guard's Plate Leggings (220798, -0.89 DPS) [vendor]; Imperial Red Pants (8251, -1.03 DPS) [world_drop]; Stormshroud Pants (15057, -1.03 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 107.9 attack_power points (5.82 DPS) | yes | Skycaller's Leather Boots (252471, -0.87 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.87 DPS) [crafted]; Fleetfoot Greaves (11627, -0.89 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 75.7 attack_power points (4.08 DPS) | yes | Cyclopean Band (11824, -0.54 DPS) [dungeon]; Woodseed Hoop (17768, -0.79 DPS) [quest]; Legionnaire's Band (19511, -1.22 DPS) [rep] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 68.0 attack_power points (3.67 DPS) | yes | Cyclopean Band (11824, -0.13 DPS) [dungeon]; Woodseed Hoop (17768, -0.38 DPS) [quest]; Legionnaire's Band (19511, -0.81 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.02 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.61 DPS) [world_drop] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | The Rockpounder (9413, -0.42 DPS) [dungeon]; Thorium Greatmace (250613, -0.71 DPS, sim-verified) [crafted]; Stoneraven (13059, -1.72 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Eye of Theradras; neck: Skibi's Pendant; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Stonerender Gauntlets; waist: Defiler's Lamellar Girdle; legs: Infernal Trickster Leggings; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Warmonger

No-known-source sample (15 of 1046, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 343.8. Weights run: 6.4s. Verify run: 83.3s. 2435 eligible items had no known source.

Race sweep: undead 343.8 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.761 ± 0.042, crit=2.676 ± 0.060 per rating point (14 rating = 1%, 37.467 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=18.303 ± 2.481, intellect=5.642 ± 0.120, spell_power=3.184 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 288.8 attack_power points (20.20 DPS) | yes | The Postmaster's Band (13390, -6.99 DPS) [dungeon]; Coif of Elements (16667, -7.37 DPS) [dungeon]; Magister's Crown (16686, -8.14 DPS, sim-verified) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 142.2 attack_power points (9.94 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -1.31 DPS, sim-verified) [world_drop]; Archlight Talisman (15856, -3.33 DPS) [quest]; Tooth of Gnarr (13141, -4.42 DPS) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 203.3 attack_power points (14.21 DPS) | yes | Darkspear Shoulderpads (272103, -0.89 DPS) [vendor]; Darkspear Shoulders (272104, -0.89 DPS) [vendor]; Mantle of the Timbermaw (19050, -2.68 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Shroud of the Exile (15421, -0.10 DPS) [quest]; Darkspear Raider's Cloak (272063, -0.41 DPS) [vendor]; Royal Tribunal Cloak (13376, -4.07 DPS, sim-verified) [dungeon] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Alanna's Embrace (13314, +0.00 DPS) [dungeon]; Blue Dragonscale Breastplate (15048, -0.18 DPS) [crafted]; Magister's Robes (16688, -0.90 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 115.2 attack_power points (8.05 DPS) | yes | Funeral Cuffs (12626, -1.20 DPS) [dungeon]; Aristocratic Cuffs (12546, -1.25 DPS) [dungeon]; Manacle Cuffs (11962, -2.38 DPS, sim-verified) [quest] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Raider Handwraps (272097, -0.50 DPS, sim-verified) [vendor]; Hands of Power (13253, -3.18 DPS) [dungeon] |
| waist | Crystallized Girdle (12606) | Blackrock Spire: Solakar Flamewreath [dungeon] | 148.6 attack_power points (10.39 DPS) | yes | Defiler's Lizardhide Girdle (20171, -1.06 DPS) [rep]; Defiler's Mail Girdle (20195, -1.06 DPS) [rep]; Magister's Belt (16685, -1.29 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Lizardhide Pants (22741, -2.44 DPS) [rep]; Champion's Legplates (250598, -4.33 DPS) [crafted]; Red Dragonscale Leggings (252603, -5.78 DPS, sim-verified) [crafted] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 147.6 attack_power points (10.32 DPS) | yes | Omnicast Boots (11822, -1.13 DPS) [dungeon]; Firemoss Boots (22275, -1.66 DPS) [dungeon]; Kayser's Boots of Precision (22231, -2.33 DPS) [dungeon] |
| finger1 | Seal of Rivendare (13345) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.72 DPS) [quest]; Deep Rooted Ring (19109, -1.87 DPS) [world]; Naglering (11669, -5.66 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Blood of the Martyr (17045, -0.15 DPS) [quest]; Deep Rooted Ring (19109, -0.15 DPS) [world] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, -0.73 DPS, sim-verified) [dungeon]; Serenity Field (272439, -2.45 DPS) [vendor]; Uther's Strength (11302, -2.67 DPS) [world_drop] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -23.85 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, -4.65 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Darkspear Shoulderguards; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Crystallized Girdle; legs: Ironfeather Leggings; feet: Dragonrider Boots; finger1: Seal of Rivendare; finger2: Maiden's Circle; trinket1: Blackhand's Breadth; trinket2: Weakness Analyzer; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2435, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (undead, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 847.0. Weights run: 6.1s. Verify run: 88.3s. 2435 eligible items had no known source.

Race sweep: undead 847.0 ± 1.0 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=1.628 ± 0.031, crit=2.200 ± 0.043 per rating point (14 rating = 1%, 30.800 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=not significant (13.719 ± 3.721), intellect=1.302 ± 0.056, spell_power=1.685 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | sim-verified (+22.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Stalwart Helm (250599, -2.43 DPS) [crafted]; Eye of Rend (12587, -4.57 DPS) [dungeon]; Black Dragonscale Helm (252605, -22.91 DPS, sim-verified) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 56.8 attack_power points (11.45 DPS) | yes | Lady Maye's Pendant (14558, -2.73 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -4.69 DPS) [world_drop]; Skibi's Pendant (13089, -4.74 DPS) [world_drop] |
| shoulder | Defiler's Lamellar Spaulders (20184) | The Defilers [rep] | 84.2 attack_power points (16.98 DPS) | yes | Mantle of the Timbermaw (19050, -1.59 DPS) [crafted]; Darkspear Pauldrons (272105, -1.91 DPS) [vendor]; Darkspear Epaulets (272106, -1.91 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.46 DPS) [rep]; Crystalline Threaded Cape (20697, -1.11 DPS) [world]; Arcanoweave Cloak (272411, -1.42 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 140.0 attack_power points (28.21 DPS) | yes | Savage Gladiator Chain (11726, -5.27 DPS, sim-verified) [dungeon]; Breastplate of Bloodthirst (12757, -9.45 DPS) [quest]; Dawn Armor (252483, -9.86 DPS) [crafted] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | 59.0 attack_power points (11.89 DPS) | yes | Forest Stalker's Bracers (19587, -0.29 DPS) [rep]; Windtalker's Wristguards (19582, -2.13 DPS) [rep] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Raider Gauntlets (272095, +0.00 DPS) [vendor] |
| waist | Vindicator's Belt (272404) | Pix Xizzix [vendor] | 85.0 attack_power points (17.14 DPS) | yes | Defiler's Lamellar Girdle (20177, -0.99 DPS) [rep]; Might of the Timbermaw (19044, -2.31 DPS) [crafted]; Defiler's Plate Girdle (20204, -2.64 DPS) [rep] |
| legs | Champion's Legplates (250598) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pristine Scorpid Leggings (252606, -1.70 DPS) [crafted]; Ironfeather Leggings (252486, -4.06 DPS) [crafted]; Outrider's Lizardhide Pants (22741, -4.38 DPS) [rep] |
| feet | Defiler's Lamellar Greaves (20181) | The Defilers [rep] | 63.8 attack_power points (12.87 DPS) | yes | Battlechaser's Greaves (12555, -1.77 DPS) [dungeon]; Defiler's Plate Greaves (20208, -2.10 DPS) [rep]; Skirmisher's Mail Sabatons (252578, -2.11 DPS) [crafted] |
| finger1 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.51 DPS) [dungeon]; Don Julio's Band (19325, -0.91 DPS) [rep]; Naglering (11669, -10.27 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.03 DPS) [dungeon]; Don Julio's Band (19325, -0.43 DPS) [rep]; Maiden's Circle (13001, -1.51 DPS) [world_drop] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (+10.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Hand of Justice (11815, -0.40 DPS) [dungeon]; Serenity Field (272439, -0.44 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, +0.00 DPS) [vendor] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Defiler's Lamellar Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Storm Gauntlets; waist: Vindicator's Belt; legs: Champion's Legplates; feet: Defiler's Lamellar Greaves; finger1: Don Mauricio's Band of Domination; finger2: Legionnaire's Band; trinket1: Blackhand's Breadth; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force

No-known-source sample (15 of 2435, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

