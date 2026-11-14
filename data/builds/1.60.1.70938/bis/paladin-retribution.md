# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 49.4. Weights run: 9.2s. Verify run: 22.9s. 312 eligible items had no known source.

Race sweep: human 49.4 ± 0.0, dwarf 48.8 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.380 ± 0.020, crit=0.820 ± 0.019 per rating point (14 rating = 1%, 11.484 per %), hit=1.548 ± 0.054 per rating point (10 rating = 1%, 15.481 per %), melee_haste=5.700 ± 0.667, intellect=1.098 ± 0.027, spell_power=3.033 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 33.4 attack_power points (1.42 DPS) | yes | Totemic Leather Hood (252448, -0.25 DPS, sim-verified) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.52 DPS) [crafted]; Veteran's Silvered Chain Helm (250528, -0.57 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.3 attack_power points (0.35 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.61 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 25.0 attack_power points (1.07 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.55 DPS) [crafted]; Serpent's Shoulders (5404, -0.77 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.12 DPS) [crafted]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Caretaker's Cape (20428, -0.13 DPS) [rep] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 35.5 attack_power points (1.51 DPS) | yes | Crusader's Chain Shirt (250492, -0.01 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.24 DPS) [crafted]; Acolyte's Chain Shirt (250491, -0.63 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.1 attack_power points (0.43 DPS) | yes | Patterned Bronze Bracers (2868, -0.01 DPS) [crafted]; Death Bindings (286981, -0.01 DPS) [dungeon]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 29.5 attack_power points (1.26 DPS) | yes | Totemic Leather Gloves (252497, -0.28 DPS, sim-verified) [crafted]; Crusader's Gloves (250512, -0.36 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.42 DPS) [crafted] |
| waist | Crusader's Chain Belt (250517) | Blacksmithing [crafted] | 21.2 attack_power points (0.90 DPS) | yes | Totemic Leather Belt (252431, -0.04 DPS) [crafted]; Cobrahn's Grasp (6460, -0.13 DPS) [dungeon]; Blackened Defias Belt (10403, -0.13 DPS) [dungeon] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 43.6 attack_power points (1.86 DPS) | yes | Crusader's Chain Leggings (250497, -0.18 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.63 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 28.2 attack_power points (1.20 DPS) | yes | Crusader's Boots (250507, -0.09 DPS) [crafted]; Spidersilk Boots (4320, -0.11 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.19 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 17.4 attack_power points (0.74 DPS) | yes | Protector's Band (20439, -0.16 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Demon Band (12054, -0.40 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 15.2 attack_power points (0.65 DPS) | yes | Protector's Band (20439, -0.24 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Demon Band (12054, -0.31 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sword of the Fallen (286977, -2.87 DPS) [dungeon]; Duskbringer (2205, -3.02 DPS) [dungeon]; Smite's Mighty Hammer (7230, -3.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Totemic Leather Armor; wrist: Bravo's Armbands; hands: Serpent Gloves; waist: Crusader's Chain Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Verigan's Fist

No-known-source sample (15 of 312, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 30 (human, 00000000000000000-00000000000000000-05233201001300100)

Set DPS (verified): 89.9. Weights run: 10.7s. Verify run: 25.2s. 537 eligible items had no known source.

Race sweep: human 89.9 ± 0.1, dwarf 88.4 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.031 ± 0.034, crit=1.560 ± 0.034 per rating point (14 rating = 1%, 21.840 per %), hit=2.123 ± 0.086 per rating point (10 rating = 1%, 21.234 per %), melee_haste=8.699 ± 1.076, intellect=1.518 ± 0.037, spell_power=3.350 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 68.9 attack_power points (3.35 DPS) | yes | Totemic Leather Helm (252456, -1.55 DPS, sim-verified) [crafted]; Holy Shroud (2721, -1.56 DPS) [world_drop]; Crusader's Silvered Chain Helm (250532, -1.56 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 22.2 attack_power points (1.08 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Sentinel's Medallion (19541, -0.29 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 43.8 attack_power points (2.13 DPS) | yes | Forest Tracker Epaulets (2278, -0.56 DPS) [world_drop]; Death Speaker Mantle (6685, -0.57 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.62 DPS) [crafted] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.15 DPS, sim-verified) [world_drop]; Hillman's Cloak (3719, -0.27 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.27 DPS) [quest] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 52.2 attack_power points (2.54 DPS) | yes | Beguiler Robes (7728, -0.19 DPS) [dungeon]; Totemic Leather Tunic (252451, -0.25 DPS) [crafted]; Green Silk Armor (7065, -0.52 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 30.2 attack_power points (1.47 DPS) | yes | Glowing Magical Bracelets (13106, -0.58 DPS) [world_drop]; Yorgen Bracers (13012, -0.59 DPS) [world_drop]; Hawkeye's Bracers (14590, -1.18 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 35.6 attack_power points (1.74 DPS) | yes | Naga Battle Gloves (888, -0.53 DPS) [dungeon]; Totemic Leather Gloves (252497, -0.53 DPS) [crafted]; Insignia Gloves (6408, -0.59 DPS, sim-verified) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 52.2 attack_power points (2.54 DPS) | yes | Skirmisher's Leather Belt (252461, -0.46 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -0.47 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.53 DPS) [rep] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 57.6 attack_power points (2.80 DPS) | yes | Totemic Leather Leggings (252458, -0.48 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.49 DPS) [crafted]; Dark Ritual Leggings (270031, -0.52 DPS) [quest] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 37.5 attack_power points (1.82 DPS) | yes | Totemic Leather Boots (252442, -0.36 DPS) [crafted]; Gilded Slippers (254001, -0.38 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -0.39 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.7 attack_power points (1.35 DPS) | yes | Thunderbrow Ring (13097, -0.27 DPS) [world_drop]; Ring of Power Regulation (273030, -0.37 DPS) [dungeon]; Protector's Band (19517, -1.39 DPS, sim-verified) [rep] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | sim-verified (89.9 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Ring of Power Regulation (273030, -0.16 DPS) [dungeon]; Protector's Band (19517, -0.37 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.14 DPS) [world]; Corpsemaker (6687, -0.63 DPS) [dungeon]; Viscous Hammer (13045, -6.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Sergeant Major's Cape; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 537, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 144.3. Weights run: 10.1s. Verify run: 22.6s. 808 eligible items had no known source.

Race sweep: human 144.3 ± 0.2, dwarf 141.5 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=2.657 ± 0.055, crit=2.654 ± 0.064 per rating point (14 rating = 1%, 37.161 per %), hit=2.133 ± 0.092 per rating point (10 rating = 1%, 21.332 per %), melee_haste=11.817 ± 1.405, intellect=2.990 ± 0.051, spell_power=3.130 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 92.0 attack_power points (4.79 DPS) | yes | Papal Fez (9431, -0.84 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.85 DPS, sim-verified) [dungeon]; Big Voodoo Mask (8201, -1.14 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 45.8 attack_power points (2.38 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.83 DPS) [world]; Sentinel's Medallion (19540, -0.86 DPS) [rep]; Kaleidoscope Chain (13084, -0.90 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 60.8 attack_power points (3.16 DPS) | yes | Green Silken Shoulders (7057, -0.15 DPS) [crafted]; Bloodmage Mantle (7684, -0.30 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.31 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.64 DPS) [crafted]; Expeditionary's Cape (271922, -0.69 DPS) [rep]; Long Silken Cloak (4326, -1.38 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 86.8 attack_power points (4.52 DPS) | yes | Dreamweave Vest (10021, -0.18 DPS) [crafted]; Robe of Power (7054, -0.37 DPS) [crafted]; Icemail Jerkin (1981, -1.58 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 36.7 attack_power points (1.91 DPS) | yes | Tough Scorpid Bracers (8205, -0.01 DPS) [crafted]; Turtle Scale Bracers (8198, -0.16 DPS) [crafted]; Windchaser Cuffs (14429, -0.23 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, -1.45 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 66.8 attack_power points (3.47 DPS) | yes | Skirmisher's Leather Belt (252461, -0.76 DPS) [crafted]; Justicar's Belt (250560, -0.88 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.96 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 79.7 attack_power points (4.15 DPS) | yes | Ferine Leggings (6690, -0.93 DPS) [dungeon]; Totemic Leather Leggings (252458, -0.96 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.03 DPS, sim-verified) [crafted] |
| feet | Skirmisher's Mail Boots (252564) | Leatherworking [crafted] | 65.6 attack_power points (3.41 DPS) | yes | Skycaller's Leather Shoes (252532, -0.21 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.21 DPS) [crafted]; Justicar's Boots (250550, -0.23 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 46.2 attack_power points (2.41 DPS) | yes | Snake Hoop (6750, -0.65 DPS) [quest]; Ring of the Underwood (2951, -0.81 DPS) [world_drop]; Falcon's Hook (7552, -0.84 DPS) [dungeon] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 37.6 attack_power points (1.96 DPS) | yes | Ring of the Underwood (2951, -0.36 DPS) [world_drop]; Falcon's Hook (7552, -0.39 DPS) [dungeon]; Snake Hoop (6750, -0.80 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, -0.10 DPS) [vendor]; Nightblade (1982, -1.37 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Sergeant Major's Cape; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Skirmisher's Mail Boots; finger1: Philanthropist's Ring; finger2: Protector's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 808, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 209.2. Weights run: 10.1s. Verify run: 38.7s. 1050 eligible items had no known source.

Race sweep: human 209.2 ± 0.2, dwarf 206.5 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=3.901 ± 0.086, crit=4.662 ± 0.111 per rating point (14 rating = 1%, 65.262 per %), hit=4.045 ± 0.201 per rating point (10 rating = 1%, 40.451 per %), melee_haste=11.775 ± 2.473, intellect=4.378 ± 0.099, spell_power=3.489 ± 0.006

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 153.8 attack_power points (8.22 DPS) | yes | Eye of Theradras (17715, -0.29 DPS) [dungeon]; Soothsayer's Headdress (17740, -0.42 DPS) [dungeon]; Chief Architect's Monocle (11839, -0.44 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 61.7 attack_power points (3.30 DPS) | yes | Horizon Choker (13085, -0.02 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -0.12 DPS) [dungeon]; Gemshard Heart (17707, -0.26 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 139.9 attack_power points (7.47 DPS) | yes | Rotgrip Mantle (17732, -1.08 DPS, sim-verified) [dungeon]; Lead Surveyor's Mantle (11842, -1.40 DPS) [dungeon]; Knight-Lieutenant's Lamellar Pauldrons (220818, -1.60 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.39 DPS) [dungeon]; Runecloth Cloak (13860, -0.46 DPS) [crafted]; Sergeant Major's Cape (16336, -0.65 DPS) [pvp] |
| chest | Robes of Insight (940) | World drop [world_drop] | 143.4 attack_power points (7.66 DPS) | yes | Icemail Jerkin (1981, -0.50 DPS) [world_drop]; Spiderfang Carapace (11633, -0.55 DPS) [dungeon]; Knight's Lamellar Chestplate (220815, -0.77 DPS) [vendor] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 79.6 attack_power points (4.25 DPS) | yes | Runic Leather Bracers (15092, -0.05 DPS) [crafted]; Skirmisher's Mail Bracers (252572, -0.54 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.57 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Sergeant Major's Imbued Gauntlets (220812, +0.00 DPS) [vendor]; Sergeant Major's Lamellar Gauntlets (220817, -3.48 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 126.1 attack_power points (6.74 DPS) | yes | Highlander's Lamellar Girdle (20106, -0.09 DPS) [rep]; Highlander's Cloth Girdle (20097, -0.40 DPS) [rep]; Highlander's Lizardhide Girdle (20103, -0.91 DPS) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 153.5 attack_power points (8.20 DPS) | yes | Knight's Plate Leggings (220797, -1.21 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -1.23 DPS) [crafted]; Knight's Lamellar Legplates (220816, -1.30 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 122.6 attack_power points (6.55 DPS) | yes | Skycaller's Leather Boots (252471, -1.18 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.18 DPS) [crafted]; Fleetfoot Greaves (11627, -1.19 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 82.2 attack_power points (4.39 DPS) | yes | Cyclopean Band (11824, -0.60 DPS) [dungeon]; Chivalrous Signet (20505, -0.81 DPS) [quest]; Woodseed Hoop (17768, -0.93 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 72.8 attack_power points (3.89 DPS) | yes | Cyclopean Band (11824, -0.10 DPS) [dungeon]; Chivalrous Signet (20505, -0.31 DPS) [quest]; Woodseed Hoop (17768, -0.43 DPS) [quest] |
| trinket1 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.68 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.26 DPS, sim-verified) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | The Rockpounder (9413, -0.99 DPS) [dungeon]; Thorium Greatmace (250613, -1.56 DPS, sim-verified) [crafted]; Stoneraven (13059, -2.71 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Red Mageweave Headband; neck: Skibi's Pendant; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; waist: Ban'thok Sash; legs: Infernal Trickster Leggings; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Sanctified Orb; trinket2: Thunderbrew's Boot Flask; main_hand: Warmonger

No-known-source sample (15 of 1050, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 357.0. Weights run: 10.0s. Verify run: 117.9s. 2437 eligible items had no known source.

Race sweep: human 357.0 ± 0.3, dwarf 355.6 ± 0.3 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.916 ± 0.046, crit=2.910 ± 0.066 per rating point (14 rating = 1%, 40.735 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=11.489 ± 2.600, intellect=6.296 ± 0.134, spell_power=3.671 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 320.2 attack_power points (22.26 DPS) | yes | Field Marshal's Lamellar Faceguard (16474, -1.08 DPS) [vendor]; Field Marshal's Lamellar Headguard (231648, -4.14 DPS) [vendor]; Magister's Crown (16686, -7.09 DPS, sim-verified) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 160.0 attack_power points (11.12 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -1.65 DPS, sim-verified) [world_drop]; Archlight Talisman (15856, -3.68 DPS) [quest]; Horizon Choker (13085, -5.00 DPS) [world_drop] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 228.7 attack_power points (15.90 DPS) | yes | Theramore Shoulderpads (271953, -1.02 DPS) [rep]; Theramore Shoulders (271954, -1.02 DPS) [rep]; Mantle of the Timbermaw (19050, -3.04 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Shroud of the Exile (15421, -0.15 DPS) [quest]; Expeditionary's Cape (271907, -0.58 DPS) [rep]; Royal Tribunal Cloak (13376, -3.56 DPS, sim-verified) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 220.3 attack_power points (15.32 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Vest of Elements (16666, -0.69 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 131.1 attack_power points (9.12 DPS) | yes | Funeral Cuffs (12626, -1.46 DPS) [dungeon]; Aristocratic Cuffs (12546, -1.53 DPS) [dungeon]; Manacle Cuffs (11962, -2.69 DPS, sim-verified) [quest] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Theramore Handwraps (271945, +0.00 DPS) [rep]; Theramore Handguards (271950, -3.30 DPS, sim-verified) [rep] |
| waist | Crystallized Girdle (12606) | Blackrock Spire: Solakar Flamewreath [dungeon] | 167.3 attack_power points (11.63 DPS) | yes | Highlander's Mail Girdle (20044, -1.36 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, -1.36 DPS) [rep]; Magister's Belt (16685, -1.39 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Leggings (231647, +0.00 DPS) [vendor]; Marshal's Lamellar Legplates (16475, -0.19 DPS) [vendor]; Red Dragonscale Leggings (252603, -4.25 DPS, sim-verified) [crafted] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 166.8 attack_power points (11.60 DPS) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor] |
| finger1 | Seal of Rivendare (13345) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.82 DPS) [quest]; Deep Rooted Ring (19109, -2.04 DPS) [world]; Naglering (11669, -8.35 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Blood of the Martyr (17045, -0.22 DPS) [quest]; Deep Rooted Ring (19109, -0.22 DPS) [world] |
| trinket1 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (+8.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -2.04 DPS) [vendor]; Uther's Strength (11302, -2.30 DPS) [world_drop] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; The Unstoppable Force (19323, -16.49 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, -4.61 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Theramore Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Storm Gauntlets; waist: Crystallized Girdle; legs: Ironfeather Leggings; feet: Dragonrider Boots; finger1: Seal of Rivendare; finger2: Maiden's Circle; trinket2: Blackhand's Breadth; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2437, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60, raid preset (human, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 834.3. Weights run: 10.7s. Verify run: 145.2s. 2437 eligible items had no known source.

Race sweep: human 834.3 ± 0.9, dwarf 827.5 ± 0.9 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=2.367 ± 0.050, crit=3.174 ± 0.068 per rating point (14 rating = 1%, 44.440 per %), hit=6.568 ± 0.349 per rating point (10 rating = 1%, 65.678 per %), melee_haste=23.921 ± 4.190, intellect=1.711 ± 0.063, spell_power=2.009 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 263.8 attack_power points (39.58 DPS) | yes | Stalwart Helm (250599, -4.06 DPS) [crafted]; Field Marshal's Lamellar Faceguard (16474, -12.62 DPS) [vendor]; Mask of the Unforgiven (13404, -13.20 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Nacreous Shell Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Lightforge Spaulders (16729, +0.00 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -2.67 DPS) [vendor]; Cloak of Flames (3475, -6.88 DPS) [world_drop]; Stalwart Cloak (272415, -6.88 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -4.26 DPS) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor] |
| waist | Highlander's Lamellar Girdle (20042) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Sash of the Grand Hunt (22207, +0.00 DPS) [dungeon]; Marksman's Girdle (22232, +0.00 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -2.63 DPS) [vendor]; Sentinel's Lamellar Legguards (22753, -3.94 DPS) [rep] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Kayser's Boots of Precision (22231, -0.06 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -5.17 DPS) [vendor]; Blessed Band of Light (272407, -5.75 DPS) [vendor]; Blackstone Ring (17713, -5.77 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.38 DPS) [vendor]; Blessed Band of Light (272407, -1.96 DPS) [vendor]; Blackstone Ring (17713, -1.97 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Cannonball Runner (13382, -0.50 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Lionheart Helm; shoulder: Theramore Pauldrons; back: Arcanoweave Cloak; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Highlander's Lamellar Girdle; legs: Sentinel's Chain Leggings; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Weakness Analyzer; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 2437, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (undead, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 46.1. Weights run: 9.2s. Verify run: 23.0s. 299 eligible items had no known source.

Race sweep: undead 46.1 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.380 ± 0.020, crit=0.820 ± 0.019 per rating point (14 rating = 1%, 11.484 per %), hit=1.548 ± 0.054 per rating point (10 rating = 1%, 15.481 per %), melee_haste=5.700 ± 0.667, intellect=1.098 ± 0.027, spell_power=3.033 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 33.4 attack_power points (1.42 DPS) | yes | Totemic Leather Hood (252448, -0.22 DPS, sim-verified) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.52 DPS) [crafted]; Veteran's Silvered Chain Helm (250528, -0.57 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.3 attack_power points (0.35 DPS) | yes | Erudite's Amulet (277204, -0.24 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 25.0 attack_power points (1.07 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.30 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.55 DPS) [crafted]; Serpent's Shoulders (5404, -0.77 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.12 DPS) [crafted]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Battle Healer's Cloak (20427, -0.13 DPS) [rep] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 35.5 attack_power points (1.51 DPS) | yes | Crusader's Chain Shirt (250492, -0.01 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.24 DPS) [crafted]; Acolyte's Chain Shirt (250491, -0.63 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Tabitha's Cuffs (251486, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.08 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 29.5 attack_power points (1.26 DPS) | yes | Totemic Leather Gloves (252497, -0.27 DPS) [crafted]; Crusader's Gloves (250512, -0.36 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.42 DPS) [crafted] |
| waist | Crusader's Chain Belt (250517) | Blacksmithing [crafted] | 21.2 attack_power points (0.90 DPS) | yes | Totemic Leather Belt (252431, -0.04 DPS) [crafted]; Cobrahn's Grasp (6460, -0.13 DPS) [dungeon]; Blackened Defias Belt (10403, -0.13 DPS) [dungeon] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 43.6 attack_power points (1.86 DPS) | yes | Crusader's Chain Leggings (250497, -0.24 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.63 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 28.2 attack_power points (1.20 DPS) | yes | Crusader's Boots (250507, -0.09 DPS) [crafted]; Spidersilk Boots (4320, -0.11 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.19 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 15.2 attack_power points (0.65 DPS) | yes | Loop of Sacrifice (281673, -0.16 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Demon Band (12054, -0.31 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.5 attack_power points (0.58 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Demon Band (12054, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Hammerbone (270018, -0.12 DPS) [quest]; Sword of the Fallen (286977, -0.22 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Totemic Leather Armor; wrist: Patterned Bronze Bracers; hands: Serpent Gloves; waist: Crusader's Chain Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Advisor's Ring; finger2: Legionnaire's Band; main_hand: Forsaken Greataxe

No-known-source sample (15 of 299, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (undead, 00000000000000000-00000000000000000-05233201001300100)

Set DPS (verified): 90.5. Weights run: 10.7s. Verify run: 23.2s. 524 eligible items had no known source.

Race sweep: undead 90.5 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.031 ± 0.034, crit=1.560 ± 0.034 per rating point (14 rating = 1%, 21.840 per %), hit=2.123 ± 0.086 per rating point (10 rating = 1%, 21.234 per %), melee_haste=8.699 ± 1.076, intellect=1.518 ± 0.037, spell_power=3.350 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 68.9 attack_power points (3.35 DPS) | yes | Totemic Leather Helm (252456, -1.24 DPS, sim-verified) [crafted]; Holy Shroud (2721, -1.56 DPS) [world_drop]; Crusader's Silvered Chain Helm (250532, -1.56 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 22.2 attack_power points (1.08 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Scout's Medallion (19537, -0.29 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 43.8 attack_power points (2.13 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.56 DPS) [world_drop]; Invoker's Mantle (215365, -0.62 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.17 DPS) [crafted]; Windsong Drape (15468, -0.17 DPS) [quest]; Tigerstrike Mantle (13108, -0.19 DPS) [world_drop] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 52.2 attack_power points (2.54 DPS) | yes | Beguiler Robes (7728, -0.19 DPS) [dungeon]; Totemic Leather Tunic (252451, -0.25 DPS) [crafted]; Green Silk Armor (7065, -0.52 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 30.2 attack_power points (1.47 DPS) | yes | Glowing Magical Bracelets (13106, -0.58 DPS) [world_drop]; Yorgen Bracers (13012, -0.59 DPS) [world_drop]; Hawkeye's Bracers (14590, -1.02 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 35.6 attack_power points (1.74 DPS) | yes | Oilrag Handwraps (16741, -0.12 DPS) [quest]; Jutebraid Gloves (10654, -0.39 DPS) [quest]; Insignia Gloves (6408, -0.46 DPS) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 52.2 attack_power points (2.54 DPS) | yes | Skirmisher's Leather Belt (252461, -0.36 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -0.47 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.53 DPS) [rep] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 57.6 attack_power points (2.80 DPS) | yes | Totemic Leather Leggings (252458, -0.18 DPS) [crafted]; Totemic Leather Pants (252446, -0.49 DPS) [crafted]; Dark Ritual Leggings (270031, -0.52 DPS) [quest] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 37.5 attack_power points (1.82 DPS) | yes | Gilded Slippers (254001, -0.16 DPS) [crafted]; Totemic Leather Boots (252442, -0.36 DPS) [crafted]; Spidersilk Boots (4320, -0.39 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.7 attack_power points (1.35 DPS) | yes | Thunderbrow Ring (13097, -0.27 DPS) [world_drop]; Ring of Power Regulation (273030, -0.37 DPS) [dungeon]; Legionnaire's Band (19513, -1.22 DPS, sim-verified) [rep] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | sim-verified (90.5 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Ring of Power Regulation (273030, -0.16 DPS) [dungeon]; Legionnaire's Band (19513, -0.36 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.14 DPS) [world]; Corpsemaker (6687, -0.63 DPS) [dungeon]; Viscous Hammer (13045, -6.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Hawkeye's Cloak; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 524, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 143.0. Weights run: 10.1s. Verify run: 33.0s. 792 eligible items had no known source.

Race sweep: undead 143.0 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=3.337 ± 0.068, crit=3.453 ± 0.082 per rating point (14 rating = 1%, 48.341 per %), hit=3.199 ± 0.154 per rating point (10 rating = 1%, 31.987 per %), melee_haste=13.329 ± 1.762, intellect=3.457 ± 0.043, spell_power=3.493 ± 0.006

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 106.8 attack_power points (5.43 DPS) | yes | Miner's Hat of the Deep (9429, -0.31 DPS) [dungeon]; Papal Fez (9431, -1.02 DPS) [dungeon]; Big Voodoo Mask (8201, -1.37 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 52.1 attack_power points (2.65 DPS) | yes | Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.98 DPS) [world]; Scout's Medallion (19536, -1.96 DPS, sim-verified) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 69.4 attack_power points (3.53 DPS) | yes | Green Silken Shoulders (7057, -0.17 DPS) [crafted]; Bloodmage Mantle (7684, -0.35 DPS) [dungeon]; Berylline Pads (4197, -0.53 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Hawkeye's Cloak (14593, -0.45 DPS) [world_drop] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.18 DPS) [crafted]; Robe of Power (7054, -0.37 DPS) [crafted]; Icemail Jerkin (1981, -0.88 DPS, sim-verified) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 44.1 attack_power points (2.24 DPS) | yes | Guardian Leather Bracers (4260, -0.12 DPS) [crafted]; Radiant Silver Bracers (4545, -0.13 DPS) [quest]; Turtle Scale Bracers (8198, -0.30 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, -2.79 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 76.3 attack_power points (3.88 DPS) | yes | Defiler's Cloth Girdle (20166, -0.69 DPS) [rep]; Skirmisher's Leather Belt (252461, -1.00 DPS) [crafted]; Defiler's Chain Girdle (20153, -1.01 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 90.4 attack_power points (4.59 DPS) | yes | Basilisk Hide Pants (1718, -1.03 DPS) [world_drop]; Ferine Leggings (6690, -1.16 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.40 DPS, sim-verified) [crafted] |
| feet | Skirmisher's Mail Boots (252564) | Leatherworking [crafted] | 71.2 attack_power points (3.62 DPS) | yes | Skycaller's Leather Shoes (252532, -0.08 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.08 DPS) [crafted]; Justicar's Boots (250550, -0.23 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 52.2 attack_power points (2.65 DPS) | yes | Snake Hoop (6750, -0.71 DPS) [quest]; Ring of the Underwood (2951, -0.75 DPS) [world_drop]; Falcon's Hook (7552, -0.82 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 43.0 attack_power points (2.19 DPS) | yes | Snake Hoop (6750, -0.24 DPS) [quest]; Ring of the Underwood (2951, -0.28 DPS) [world_drop]; Falcon's Hook (7552, -0.35 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -2.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Skirmisher's Mail Boots; finger1: Philanthropist's Ring; finger2: Legionnaire's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 792, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 204.5. Weights run: 10.1s. Verify run: 32.7s. 1050 eligible items had no known source.

Race sweep: undead 204.5 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.922 ± 0.064, crit=3.469 ± 0.085 per rating point (14 rating = 1%, 48.564 per %), hit=2.858 ± 0.123 per rating point (10 rating = 1%, 28.581 per %), melee_haste=15.523 ± 2.030, intellect=2.796 ± 0.094, spell_power=3.237 ± 0.006

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Headband (10033, -0.37 DPS, sim-verified) [crafted]; Foreman's Head Protector (22223, -0.38 DPS) [dungeon]; Blood Guard's Plate Helm (220803, -0.56 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 49.0 attack_power points (2.57 DPS) | yes | Glowing Eye of Mordresh (10769, -0.21 DPS) [dungeon]; Jarkal's Enhancing Necklace (7888, -0.39 DPS) [quest]; Gemshard Heart (17707, -0.40 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 104.5 attack_power points (5.48 DPS) | yes | Lead Surveyor's Mantle (11842, -0.88 DPS) [dungeon]; Earthslag Shoulders (11632, -0.94 DPS) [dungeon]; Rotgrip Mantle (17732, -1.33 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Blackveil Cape (11626, -0.52 DPS) [dungeon]; Runecloth Cloak (13860, -0.66 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Icemail Jerkin (1981, -0.21 DPS) [world_drop]; Feathered Breastplate (8349, -0.70 DPS) [crafted]; Spiderfang Carapace (11633, -0.87 DPS, sim-verified) [dungeon] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 60.3 attack_power points (3.17 DPS) | yes | Skirmisher's Mail Bracers (252572, -0.03 DPS) [crafted]; Aristocratic Cuffs (12546, -0.29 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.32 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Justicar's Gloves (250555, +0.00 DPS) [crafted]; Skirmisher's Mail Gauntlets (252586, +0.00 DPS) [crafted]; Stonerender Gauntlets (17007, -2.16 DPS, sim-verified) [world_drop] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 95.4 attack_power points (5.01 DPS) | yes | Defiler's Lamellar Girdle (20179, -0.01 DPS) [rep]; Defiler's Cloth Girdle (20165, -0.19 DPS) [rep]; Skycaller's Leather Waistguard (252476, -0.53 DPS) [crafted] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 109.4 attack_power points (5.74 DPS) | yes | Stone Guard's Plate Leggings (220798, -0.27 DPS) [vendor]; Spellshock Leggings (9484, -0.37 DPS) [dungeon]; Stormshroud Pants (15057, -0.64 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 92.2 attack_power points (4.84 DPS) | yes | Skycaller's Mail Sabatons (252577, -0.67 DPS) [crafted]; Entrenching Boots (22270, -0.91 DPS) [dungeon]; Skycaller's Leather Boots (252471, -1.48 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 63.6 attack_power points (3.34 DPS) | yes | Blackstone Ring (17713, -0.79 DPS) [dungeon]; Legionnaire's Band (19511, -0.80 DPS) [rep]; Mindseye Circle (10634, -0.83 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 57.5 attack_power points (3.02 DPS) | yes | Blackstone Ring (17713, -0.47 DPS) [dungeon]; Legionnaire's Band (19511, -0.48 DPS) [rep]; Mindseye Circle (10634, -0.51 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.33 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.06 DPS, sim-verified) [world_drop] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | The Rockpounder (9413, -0.86 DPS) [dungeon]; Taran Icebreaker (2915, -1.37 DPS) [world_drop]; Thorium Greatmace (250613, -2.07 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Skibi's Pendant; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Robes of Insight; wrist: Runic Leather Bracers; waist: Ban'thok Sash; legs: Infernal Trickster Leggings; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Warmonger

No-known-source sample (15 of 1050, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 357.9. Weights run: 10.0s. Verify run: 127.9s. 2462 eligible items had no known source.

Race sweep: undead 357.9 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=4.079 ± 0.095, crit=5.863 ± 0.136 per rating point (14 rating = 1%, 82.085 per %), hit=4.795 ± 0.242 per rating point (10 rating = 1%, 47.946 per %), melee_haste=20.733 ± 2.966, intellect=5.227 ± 0.147, spell_power=3.362 ± 0.006

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 366.3 attack_power points (18.87 DPS) | yes | Lionheart Helm (12640, -3.43 DPS) [crafted]; Stalwart Helm (250599, -5.99 DPS) [crafted]; Blue Dragonscale Helm (252604, -7.22 DPS) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 136.3 attack_power points (7.02 DPS) | yes | Nacreous Shell Necklace (22403, -0.84 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -1.13 DPS) [world_drop]; Mark of Fordring (15411, -1.45 DPS) [quest] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Darkspear Pauldrons (272105, -0.14 DPS) [vendor]; Darkspear Epaulets (272106, -0.14 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Tribunal Cloak (13376, -2.22 DPS) [dungeon]; Shroud of the Exile (15421, -2.62 DPS) [quest]; Crystalline Threaded Cape (20697, -2.85 DPS) [world] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, +0.00 DPS) [dungeon]; Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Spitfire Bracers (20481, -0.26 DPS) [crafted]; Manacle Cuffs (11962, -0.71 DPS) [quest] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodmail Gauntlets (14615, +0.00 DPS) [dungeon]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Raider Handguards (272101, +0.00 DPS) [vendor] |
| waist | Crystallized Girdle (12606) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Lizardhide Girdle (20171, +0.00 DPS) [rep]; Defiler's Lamellar Girdle (20177, +0.00 DPS) [rep]; Defiler's Mail Girdle (20195, +0.00 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Chain Leggings (22673, -1.60 DPS) [rep]; Champion's Legplates (250598, -2.81 DPS) [crafted]; Pristine Scorpid Leggings (252606, -2.81 DPS) [crafted] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Kayser's Boots of Precision (22231, +0.00 DPS) [dungeon] |
| finger1 | Seal of Rivendare (13345) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, -0.55 DPS) [vendor] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Don Mauricio's Band of Domination (22433, -0.11 DPS) [dungeon]; Cutthroat's Signet (272408, -0.83 DPS) [vendor] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -0.65 DPS) [dungeon]; Rune of the Guard Captain (19120, -0.67 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.89 DPS) [crafted] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Darkspear Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Storm Gauntlets; waist: Crystallized Girdle; legs: Ironfeather Leggings; feet: Dragonrider Boots; finger1: Seal of Rivendare; finger2: Channeler's Ring; trinket1: Blackhand's Breadth; trinket2: Weakness Analyzer; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (undead, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 840.4. Weights run: 10.7s. Verify run: 145.3s. 2462 eligible items had no known source.

Race sweep: undead 840.4 ± 1.0 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=1.644 ± 0.032, crit=2.215 ± 0.044 per rating point (14 rating = 1%, 31.013 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=22.958 ± 3.210, intellect=1.418 ± 0.021, spell_power=1.866 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalwart Helm (250599, -1.10 DPS) [crafted]; Eye of Rend (12587, -4.43 DPS) [dungeon]; Black Dragonscale Helm (252605, -23.34 DPS, sim-verified) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 57.0 attack_power points (11.08 DPS) | yes | Lady Maye's Pendant (14558, -1.85 DPS) [world_drop]; Emberfury Talisman (12929, -2.88 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -3.94 DPS) [world_drop] |
| shoulder | Defiler's Lamellar Spaulders (20184) | The Defilers [rep] | 85.7 attack_power points (16.66 DPS) | yes | Mantle of the Timbermaw (19050, -0.88 DPS) [crafted]; Darkspear Pauldrons (272105, -2.00 DPS) [vendor]; Darkspear Epaulets (272106, -2.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -0.32 DPS) [world]; Deathguard's Cloak (20068, -0.47 DPS) [rep]; Arcanoweave Cloak (272411, -0.67 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 140.5 attack_power points (27.31 DPS) | yes | Savage Gladiator Chain (11726, -5.23 DPS, sim-verified) [dungeon]; Breastplate of Bloodthirst (12757, -9.13 DPS) [quest]; Dawn Armor (252483, -9.45 DPS) [crafted] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | 59.1 attack_power points (11.50 DPS) | yes | Forest Stalker's Bracers (19587, -0.25 DPS) [rep]; Dryad's Wrist Bindings (19595, -1.31 DPS) [rep] |
| hands | Bloodmail Gauntlets (14615) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Raider Gauntlets (272095, +0.00 DPS) [vendor] |
| waist | Vindicator's Belt (272404) | Pix Xizzix [vendor] | 87.9 attack_power points (17.09 DPS) | yes | Defiler's Lamellar Girdle (20177, -1.25 DPS) [rep]; Might of the Timbermaw (19044, -2.74 DPS) [crafted]; Defiler's Plate Girdle (20204, -3.07 DPS) [rep] |
| legs | Champion's Legplates (250598) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pristine Scorpid Leggings (252606, +0.00 DPS) [crafted]; Ironfeather Leggings (252486, -3.50 DPS) [crafted]; Outrider's Lizardhide Pants (22741, -4.68 DPS) [rep] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Death Knight Sabatons (18692, +0.00 DPS) [dungeon]; Defiler's Lamellar Greaves (20181, -13.56 DPS, sim-verified) [rep] |
| finger1 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -0.86 DPS) [rep]; Maiden's Circle (13001, -1.56 DPS) [world_drop]; Naglering (11669, -10.22 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.29 DPS) [quest]; Maiden's Circle (13001, -0.29 DPS) [world_drop]; Legionnaire's Band (19510, -13.99 DPS, sim-verified) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+12.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackblade of Shahram (12592, +0.00 DPS) [dungeon]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, +0.00 DPS) [vendor] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Defiler's Lamellar Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Bloodmail Gauntlets; waist: Vindicator's Belt; legs: Champion's Legplates; feet: Bloodmail Boots; finger1: Don Mauricio's Band of Domination; finger2: Don Julio's Band; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force

No-known-source sample (15 of 2462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

