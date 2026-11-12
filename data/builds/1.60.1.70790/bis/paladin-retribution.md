# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 49.4. Weights run: 6.9s. Verify run: 17.0s. 312 eligible items had no known source.

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

Set DPS (verified): 90.2. Weights run: 7.6s. Verify run: 13.3s. 537 eligible items had no known source.

Race sweep: human 90.2 ± 0.1, dwarf 88.7 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.030 ± 0.034, crit=1.560 ± 0.034 per rating point (14 rating = 1%, 21.838 per %), hit=2.088 ± 0.087 per rating point (10 rating = 1%, 20.882 per %), melee_haste=8.219 ± 1.074, intellect=1.510 ± 0.038, spell_power=3.364 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 69.1 attack_power points (3.36 DPS) | yes | Totemic Leather Helm (252456, -1.55 DPS, sim-verified) [crafted]; Holy Shroud (2721, -1.56 DPS) [world_drop]; Crusader's Silvered Chain Helm (250532, -1.56 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 22.2 attack_power points (1.08 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Sentinel's Medallion (19541, -0.29 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 43.9 attack_power points (2.14 DPS) | yes | Death Speaker Mantle (6685, -0.39 DPS, sim-verified) [dungeon]; Forest Tracker Epaulets (2278, -0.56 DPS) [world_drop]; Invoker's Mantle (215365, -0.62 DPS) [crafted] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.09 DPS) [world_drop]; Hillman's Cloak (3719, -0.26 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.26 DPS) [quest] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 52.4 attack_power points (2.55 DPS) | yes | Beguiler Robes (7728, -0.19 DPS) [dungeon]; Totemic Leather Tunic (252451, -0.26 DPS) [crafted]; Green Silk Armor (7065, -0.49 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 30.3 attack_power points (1.47 DPS) | yes | Yorgen Bracers (13012, -0.59 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.59 DPS) [world_drop]; Hawkeye's Bracers (14590, -1.08 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 35.7 attack_power points (1.74 DPS) | yes | Naga Battle Gloves (888, -0.53 DPS) [dungeon]; Totemic Leather Gloves (252497, -0.53 DPS) [crafted]; Insignia Gloves (6408, -0.59 DPS, sim-verified) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 52.4 attack_power points (2.55 DPS) | yes | Skycaller's Leather Belt (252522, -0.47 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.53 DPS) [rep]; Skirmisher's Leather Belt (252461, -0.54 DPS, sim-verified) [crafted] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 57.7 attack_power points (2.81 DPS) | yes | Totemic Leather Leggings (252458, -0.49 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.49 DPS) [crafted]; Dark Ritual Leggings (270031, -0.52 DPS) [quest] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 37.5 attack_power points (1.83 DPS) | yes | Totemic Leather Boots (252442, -0.36 DPS) [crafted]; Spidersilk Boots (4320, -0.39 DPS) [crafted]; Gilded Slippers (254001, -0.53 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.8 attack_power points (1.50 DPS) | yes | Thunderbrow Ring (13097, -0.42 DPS) [world_drop]; Ring of Power Regulation (273030, -0.51 DPS) [dungeon]; Protector's Band (19517, -1.60 DPS, sim-verified) [rep] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Ring of Power Regulation (273030, -0.16 DPS) [dungeon]; Protector's Band (19517, -0.41 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.16 DPS) [world]; Corpsemaker (6687, -0.63 DPS) [dungeon]; Viscous Hammer (13045, -6.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Sergeant Major's Cape; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 537, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 144.0. Weights run: 7.5s. Verify run: 21.2s. 808 eligible items had no known source.

Race sweep: human 144.0 ± 0.2, dwarf 141.1 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=3.519 ± 0.069, crit=3.627 ± 0.084 per rating point (14 rating = 1%, 50.780 per %), hit=3.260 ± 0.157 per rating point (10 rating = 1%, 32.597 per %), melee_haste=14.041 ± 1.763, intellect=3.341 ± 0.045, spell_power=3.523 ± 0.006

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 105.7 attack_power points (5.31 DPS) | yes | Miner's Hat of the Deep (9429, -0.73 DPS, sim-verified) [dungeon]; Papal Fez (9431, -1.04 DPS) [dungeon]; Big Voodoo Mask (8201, -1.37 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 51.4 attack_power points (2.58 DPS) | yes | Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.97 DPS) [world]; Sentinel's Medallion (19540, -1.39 DPS, sim-verified) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 68.1 attack_power points (3.42 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Bloodmage Mantle (7684, -0.32 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.45 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.78 DPS) [crafted]; Expeditionary's Cape (271922, -0.84 DPS) [rep]; Long Silken Cloak (4326, -1.32 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 97.6 attack_power points (4.90 DPS) | yes | Robe of Power (7054, -0.41 DPS) [crafted]; Quillward Harness (10583, -0.72 DPS) [dungeon]; Dreamweave Vest (10021, -0.80 DPS, sim-verified) [crafted] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 44.7 attack_power points (2.24 DPS) | yes | Guardian Leather Bracers (4260, -0.18 DPS) [crafted]; Turtle Scale Bracers (8198, -0.34 DPS) [crafted]; Green Leather Bracers (4259, -0.34 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, -1.16 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 74.8 attack_power points (3.76 DPS) | yes | Highlander's Chain Girdle (20089, -0.80 DPS) [rep]; Skirmisher's Leather Belt (252461, -0.92 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.13 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 89.4 attack_power points (4.49 DPS) | yes | Crimson Silk Pantaloons (7062, -0.89 DPS) [crafted]; Triprunner Dungarees (9624, -1.00 DPS) [quest]; Basilisk Hide Pants (1718, -1.60 DPS, sim-verified) [world_drop] |
| feet | Skirmisher's Mail Boots (252564) | Leatherworking [crafted] | 70.9 attack_power points (3.56 DPS) | yes | Skycaller's Mail Boots (252563, -0.08 DPS) [crafted]; Justicar's Boots (250550, -0.19 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.92 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 51.9 attack_power points (2.61 DPS) | yes | Protector's Band (19515, -0.37 DPS) [rep]; Ring of the Underwood (2951, -0.64 DPS) [world_drop]; Falcon's Hook (7552, -0.71 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 44.5 attack_power points (2.24 DPS) | yes | Ring of the Underwood (2951, -0.26 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon]; Protector's Band (19515, -0.77 DPS, sim-verified) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -1.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Sergeant Major's Cape; chest: Robe of the Magi; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Skirmisher's Mail Boots; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 808, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 207.9. Weights run: 7.5s. Verify run: 26.3s. 1047 eligible items had no known source.

Race sweep: human 207.9 ± 0.2, dwarf 205.8 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.931 ± 0.064, crit=3.476 ± 0.085 per rating point (14 rating = 1%, 48.662 per %), hit=2.726 ± 0.122 per rating point (10 rating = 1%, 27.265 per %), melee_haste=15.701 ± 2.055, intellect=2.834 ± 0.094, spell_power=3.249 ± 0.006

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 118.4 attack_power points (6.22 DPS) | yes | Soothsayer's Headdress (17740, -0.06 DPS) [dungeon]; Knight-Lieutenant's Lamellar Helm (220819, -0.26 DPS) [vendor]; Foreman's Head Protector (22223, -0.46 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 49.1 attack_power points (2.58 DPS) | yes | Glowing Eye of Mordresh (10769, -0.19 DPS) [dungeon]; Talvash's Enhancing Necklace (7673, -0.38 DPS) [quest]; Gemshard Heart (17707, -0.39 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 105.4 attack_power points (5.53 DPS) | yes | Knight-Lieutenant's Lamellar Pauldrons (220818, -0.73 DPS) [vendor]; Lead Surveyor's Mantle (11842, -0.89 DPS) [dungeon]; Rotgrip Mantle (17732, -1.63 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -0.56 DPS) [crafted]; Sergeant Major's Cape (16336, -0.77 DPS) [pvp]; Blackveil Cape (11626, -1.10 DPS, sim-verified) [dungeon] |
| chest | Spiderfang Carapace (11633) | Blackrock Depths: Hedrum the Creeper [dungeon] | 109.9 attack_power points (5.77 DPS) | yes | Knight's Lamellar Chestplate (220815, -0.07 DPS) [vendor]; Robes of Insight (940, -0.32 DPS) [world_drop]; Feathered Breastplate (8349, -1.04 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 60.8 attack_power points (3.19 DPS) | yes | Skirmisher's Mail Bracers (252572, -0.05 DPS) [crafted]; Aristocratic Cuffs (12546, -0.28 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.33 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Lamellar Gauntlets (220817, +0.00 DPS) [vendor]; Skirmisher's Mail Gauntlets (252586, +0.00 DPS) [crafted]; Stonerender Gauntlets (17007, -2.54 DPS, sim-verified) [world_drop] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 95.5 attack_power points (5.02 DPS) | yes | Ban'thok Sash (11662, -0.05 DPS) [dungeon]; Highlander's Cloth Girdle (20097, -0.18 DPS) [rep]; Skycaller's Leather Waistguard (252476, -0.50 DPS) [crafted] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 108.6 attack_power points (5.70 DPS) | yes | Knight's Plate Leggings (220797, -0.22 DPS) [vendor]; Spellshock Leggings (9484, -0.29 DPS) [dungeon]; Knight's Lamellar Legplates (220816, -1.27 DPS, sim-verified) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 91.3 attack_power points (4.79 DPS) | yes | Skycaller's Mail Sabatons (252577, -0.60 DPS) [crafted]; Entrenching Boots (22270, -0.84 DPS) [dungeon]; Skycaller's Leather Boots (252471, -1.49 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 64.1 attack_power points (3.36 DPS) | yes | Chivalrous Signet (20505, -0.49 DPS) [quest]; Protector's Band (19516, -0.82 DPS) [rep]; Woodseed Hoop (17768, -0.85 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 57.9 attack_power points (3.04 DPS) | yes | Chivalrous Signet (20505, -0.17 DPS) [quest]; Protector's Band (19516, -0.50 DPS) [rep]; Woodseed Hoop (17768, -0.53 DPS) [quest] |
| trinket1 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (+6.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.27 DPS) [crafted]; Uther's Strength (11302, -1.53 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -3.24 DPS, sim-verified) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -3.44 DPS) [world_drop]; The Rockpounder (9413, -4.08 DPS) [dungeon]; Taran Icebreaker (2915, -4.60 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Red Mageweave Headband; neck: Skibi's Pendant; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Spiderfang Carapace; wrist: Runic Leather Bracers; waist: Highlander's Lamellar Girdle; legs: Infernal Trickster Leggings; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Sanctified Orb; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 1047, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 361.8. Weights run: 7.8s. Verify run: 92.4s. 2436 eligible items had no known source.

Race sweep: human 361.8 ± 0.3, dwarf 358.9 ± 0.3 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=2.009 ± 0.045, crit=3.085 ± 0.064 per rating point (14 rating = 1%, 43.194 per %), hit=4.212 ± 0.287 per rating point (10 rating = 1%, 42.116 per %), melee_haste=not significant (12.556 ± 3.683), intellect=7.304 ± 0.137, spell_power=4.382 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 362.9 attack_power points (21.81 DPS) | yes | Field Marshal's Lamellar Faceguard (16474, -1.08 DPS) [vendor]; Field Marshal's Lamellar Headguard (231648, -4.19 DPS) [vendor]; Blue Dragonscale Helm (252604, -5.93 DPS, sim-verified) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 187.0 attack_power points (11.23 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -1.76 DPS) [world_drop]; Nacreous Shell Necklace (22403, -2.82 DPS) [dungeon]; Archlight Talisman (15856, -3.69 DPS) [quest] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 267.3 attack_power points (16.06 DPS) | yes | Theramore Shoulderpads (271953, -1.05 DPS) [rep]; Theramore Shoulders (271954, -1.05 DPS) [rep]; Mantle of the Timbermaw (19050, -3.29 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.71 DPS) [quest]; Shroud of Arcane Mastery (22330, -2.90 DPS) [dungeon]; Royal Tribunal Cloak (13376, -4.35 DPS, sim-verified) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 257.1 attack_power points (15.45 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Vest of Elements (16666, -0.61 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 154.8 attack_power points (9.30 DPS) | yes | Funeral Cuffs (12626, -1.58 DPS) [dungeon]; Aristocratic Cuffs (12546, -1.67 DPS) [dungeon]; Manacle Cuffs (11962, -2.41 DPS, sim-verified) [quest] |
| hands | Theramore Handguards (271950) | Theramore Expeditionary Force [rep] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Lamellar Gloves (231643, -0.56 DPS) [pvp]; Theramore Handwraps (271945, -0.73 DPS, sim-verified) [rep] |
| waist | Crystallized Girdle (12606) | Blackrock Spire: Solakar Flamewreath [dungeon] | 195.7 attack_power points (11.76 DPS) | yes | Lightforge Belt (16723, -0.75 DPS) [dungeon]; Clutch of Andros (13956, -0.89 DPS) [dungeon]; Magister's Belt (16685, -1.38 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Leggings (231647, +0.00 DPS) [vendor]; Marshal's Lamellar Legplates (16475, -0.69 DPS) [vendor]; Sentinel's Lamellar Legguards (237814, -2.00 DPS, sim-verified) [vendor] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 195.7 attack_power points (11.76 DPS) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor]; Kayser's Boots of Precision (22231, -0.28 DPS) [dungeon] |
| finger1 | Seal of Rivendare (13345) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.76 DPS) [quest]; Maiden's Circle (13001, -1.76 DPS) [world_drop]; Naglering (11669, -6.91 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.86 DPS) [quest]; Maiden's Circle (13001, -0.86 DPS) [world_drop]; Deep Rooted Ring (19109, -1.13 DPS) [world] |
| trinket1 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (+7.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Smolderweb's Eye (13213, -1.42 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.67 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -1.99 DPS, sim-verified) [vendor] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -23.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, -4.69 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Theramore Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Theramore Handguards; waist: Crystallized Girdle; legs: Ironfeather Leggings; feet: Dragonrider Boots; finger1: Seal of Rivendare; finger2: Channeler's Ring; trinket2: Blackhand's Breadth; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2436, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60, raid preset (human, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 867.1. Weights run: 7.9s. Verify run: 110.0s. 2436 eligible items had no known source.

Race sweep: human 867.1 ± 1.1, dwarf 859.6 ± 1.1 (dps); the guide race human is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.006, agility=2.380 ± 0.049, crit=3.179 ± 0.067 per rating point (14 rating = 1%, 44.505 per %), hit=6.249 ± 0.389 per rating point (10 rating = 1%, 62.493 per %), melee_haste=30.205 ± 4.732, intellect=2.054 ± 0.129, spell_power=2.031 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 257.6 attack_power points (39.95 DPS) | yes | Stalwart Helm (250599, -4.09 DPS) [crafted]; Field Marshal's Lamellar Faceguard (16474, -10.92 DPS) [vendor]; Black Dragonscale Helm (252605, -13.19 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Nacreous Shell Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Lightforge Spaulders (16729, +0.00 DPS) [dungeon]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -3.25 DPS) [vendor]; Shroud of Arcane Mastery (22330, -4.09 DPS) [dungeon]; Cloak of Flames (3475, -7.59 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -4.41 DPS) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor] |
| waist | Highlander's Lamellar Girdle (20042) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Sash of the Grand Hunt (22207, +0.00 DPS) [dungeon]; Marksman's Girdle (22232, +0.00 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -2.79 DPS) [vendor]; Sentinel's Lamellar Legguards (22753, -4.07 DPS) [rep] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Kayser's Boots of Precision (22231, +0.00 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -4.61 DPS) [vendor]; Blessed Band of Light (272407, -5.92 DPS) [vendor]; Blackstone Ring (17713, -5.97 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -0.71 DPS) [vendor]; Blessed Band of Light (272407, -2.02 DPS) [vendor]; Blackstone Ring (17713, -2.07 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Cannonball Runner (13382, -0.83 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Lionheart Helm; shoulder: Theramore Pauldrons; back: Arcanoweave Cloak; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Highlander's Lamellar Girdle; legs: Sentinel's Chain Leggings; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Weakness Analyzer; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 2436, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (undead, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 46.1. Weights run: 6.9s. Verify run: 17.1s. 299 eligible items had no known source.

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

Set DPS (verified): 90.8. Weights run: 7.6s. Verify run: 13.0s. 524 eligible items had no known source.

Race sweep: undead 90.8 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.030 ± 0.034, crit=1.560 ± 0.034 per rating point (14 rating = 1%, 21.838 per %), hit=2.088 ± 0.087 per rating point (10 rating = 1%, 20.882 per %), melee_haste=8.219 ± 1.074, intellect=1.510 ± 0.038, spell_power=3.364 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 69.1 attack_power points (3.36 DPS) | yes | Totemic Leather Helm (252456, -1.24 DPS, sim-verified) [crafted]; Holy Shroud (2721, -1.56 DPS) [world_drop]; Crusader's Silvered Chain Helm (250532, -1.56 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 22.2 attack_power points (1.08 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Scout's Medallion (19537, -0.29 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 43.9 attack_power points (2.14 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.56 DPS) [world_drop]; Invoker's Mantle (215365, -0.62 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.16 DPS) [crafted]; Windsong Drape (15468, -0.16 DPS) [quest]; Tigerstrike Mantle (13108, -0.19 DPS) [world_drop] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 52.4 attack_power points (2.55 DPS) | yes | Beguiler Robes (7728, -0.19 DPS) [dungeon]; Totemic Leather Tunic (252451, -0.26 DPS) [crafted]; Green Silk Armor (7065, -0.60 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 30.3 attack_power points (1.47 DPS) | yes | Hawkeye's Bracers (14590, -0.49 DPS) [world_drop]; Yorgen Bracers (13012, -0.59 DPS) [world_drop]; Undead Knight's Bracers (251965, -1.12 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 35.7 attack_power points (1.74 DPS) | yes | Oilrag Handwraps (16741, -0.12 DPS) [quest]; Jutebraid Gloves (10654, -0.39 DPS) [quest]; Insignia Gloves (6408, -0.46 DPS) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 52.4 attack_power points (2.55 DPS) | yes | Skirmisher's Leather Belt (252461, -0.23 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.47 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.53 DPS) [rep] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 57.7 attack_power points (2.81 DPS) | yes | Totemic Leather Leggings (252458, -0.39 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.49 DPS) [crafted]; Dark Ritual Leggings (270031, -0.52 DPS) [quest] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 37.5 attack_power points (1.83 DPS) | yes | Gilded Slippers (254001, -0.17 DPS) [crafted]; Totemic Leather Boots (252442, -0.36 DPS) [crafted]; Spidersilk Boots (4320, -0.39 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.8 attack_power points (1.50 DPS) | yes | Thunderbrow Ring (13097, -0.42 DPS) [world_drop]; Ring of Power Regulation (273030, -0.51 DPS) [dungeon]; Legionnaire's Band (19513, -1.43 DPS, sim-verified) [rep] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Ring of Power Regulation (273030, -0.16 DPS) [dungeon]; Legionnaire's Band (19513, -0.48 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.16 DPS) [world]; Corpsemaker (6687, -0.63 DPS) [dungeon]; Viscous Hammer (13045, -6.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Hawkeye's Cloak; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 524, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 142.4. Weights run: 7.5s. Verify run: 20.6s. 792 eligible items had no known source.

Race sweep: undead 142.4 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=3.519 ± 0.069, crit=3.627 ± 0.084 per rating point (14 rating = 1%, 50.780 per %), hit=3.260 ± 0.157 per rating point (10 rating = 1%, 32.597 per %), melee_haste=14.041 ± 1.763, intellect=3.341 ± 0.045, spell_power=3.523 ± 0.006

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 105.7 attack_power points (5.31 DPS) | yes | Miner's Hat of the Deep (9429, -0.33 DPS) [dungeon]; Papal Fez (9431, -1.04 DPS) [dungeon]; Big Voodoo Mask (8201, -1.37 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 51.4 attack_power points (2.58 DPS) | yes | Scout's Medallion (19536, -0.64 DPS) [rep]; Kaleidoscope Chain (13084, -0.79 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.97 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 68.1 attack_power points (3.42 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Bloodmage Mantle (7684, -0.32 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.45 DPS) [crafted] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.05 DPS) [vendor]; Hawkeye's Cloak (14593, -0.36 DPS) [world_drop] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 97.6 attack_power points (4.90 DPS) | yes | Dreamweave Vest (10021, -0.20 DPS) [crafted]; Robe of Power (7054, -0.41 DPS) [crafted]; Quillward Harness (10583, -0.72 DPS) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 44.7 attack_power points (2.24 DPS) | yes | Guardian Leather Bracers (4260, -0.18 DPS) [crafted]; Radiant Silver Bracers (4545, -0.19 DPS) [quest]; Turtle Scale Bracers (8198, -0.34 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, -1.89 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 74.8 attack_power points (3.76 DPS) | yes | Defiler's Cloth Girdle (20166, -0.61 DPS) [rep]; Defiler's Chain Girdle (20153, -0.80 DPS) [rep]; Skirmisher's Leather Belt (252461, -0.92 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 89.4 attack_power points (4.49 DPS) | yes | Basilisk Hide Pants (1718, -0.78 DPS) [world_drop]; Crimson Silk Pantaloons (7062, -0.89 DPS) [crafted]; Triprunner Dungarees (9624, -1.00 DPS) [quest] |
| feet | Skirmisher's Mail Boots (252564) | Leatherworking [crafted] | 70.9 attack_power points (3.56 DPS) | yes | Skycaller's Leather Shoes (252532, -0.08 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.08 DPS) [crafted]; Justicar's Boots (250550, -0.19 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 51.9 attack_power points (2.61 DPS) | yes | Legionnaire's Band (19512, -0.37 DPS) [rep]; Ring of the Underwood (2951, -0.64 DPS) [world_drop]; Falcon's Hook (7552, -0.71 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 44.5 attack_power points (2.24 DPS) | yes | Legionnaire's Band (19512, -0.00 DPS) [rep]; Ring of the Underwood (2951, -0.26 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -0.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Skirmisher's Mail Boots; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 792, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 202.7. Weights run: 7.5s. Verify run: 29.7s. 1047 eligible items had no known source.

Race sweep: undead 202.7 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.931 ± 0.064, crit=3.476 ± 0.085 per rating point (14 rating = 1%, 48.662 per %), hit=2.726 ± 0.122 per rating point (10 rating = 1%, 27.265 per %), melee_haste=15.701 ± 2.055, intellect=2.834 ± 0.094, spell_power=3.249 ± 0.006

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 118.4 attack_power points (6.22 DPS) | yes | Soothsayer's Headdress (17740, -0.06 DPS) [dungeon]; Foreman's Head Protector (22223, -0.46 DPS) [dungeon]; Blood Guard's Plate Helm (220803, -0.73 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 49.1 attack_power points (2.58 DPS) | yes | Glowing Eye of Mordresh (10769, -0.19 DPS) [dungeon]; Jarkal's Enhancing Necklace (7888, -0.38 DPS) [quest]; Gemshard Heart (17707, -0.39 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 105.4 attack_power points (5.53 DPS) | yes | Lead Surveyor's Mantle (11842, -0.89 DPS) [dungeon]; Earthslag Shoulders (11632, -0.97 DPS) [dungeon]; Rotgrip Mantle (17732, -1.15 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.11 DPS) [dungeon]; Blackveil Cape (11626, -0.54 DPS) [dungeon]; Runecloth Cloak (13860, -0.66 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Spiderfang Carapace (11633, -0.55 DPS, sim-verified) [dungeon]; Feathered Breastplate (8349, -0.72 DPS) [crafted]; Robe of the Magi (1716, -0.80 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 60.8 attack_power points (3.19 DPS) | yes | Skirmisher's Mail Bracers (252572, -0.05 DPS) [crafted]; Aristocratic Cuffs (12546, -0.28 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.33 DPS) [crafted] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skirmisher's Mail Gauntlets (252586, -0.50 DPS) [crafted]; Justicar's Gloves (250555, -0.61 DPS) [crafted]; Gauntlets of the Sea (8346, -0.74 DPS) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | 95.5 attack_power points (5.02 DPS) | yes | Ban'thok Sash (11662, -0.05 DPS) [dungeon]; Defiler's Cloth Girdle (20165, -0.18 DPS) [rep]; Skycaller's Leather Waistguard (252476, -0.50 DPS) [crafted] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 108.6 attack_power points (5.70 DPS) | yes | Spellshock Leggings (9484, -0.29 DPS) [dungeon]; Stormshroud Pants (15057, -0.59 DPS) [crafted]; Stone Guard's Plate Leggings (220798, -1.60 DPS, sim-verified) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 91.3 attack_power points (4.79 DPS) | yes | Skycaller's Mail Sabatons (252577, -0.60 DPS) [crafted]; Entrenching Boots (22270, -0.84 DPS) [dungeon]; Skycaller's Leather Boots (252471, -2.10 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 64.1 attack_power points (3.36 DPS) | yes | Legionnaire's Band (19511, -0.82 DPS) [rep]; Woodseed Hoop (17768, -0.85 DPS) [quest]; Mindseye Circle (10634, -0.86 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 57.9 attack_power points (3.04 DPS) | yes | Legionnaire's Band (19511, -0.50 DPS) [rep]; Woodseed Hoop (17768, -0.53 DPS) [quest]; Mindseye Circle (10634, -0.54 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.71 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.50 DPS, sim-verified) [world_drop] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | The Rockpounder (9413, -0.64 DPS) [dungeon]; Thorium Greatmace (250613, -0.68 DPS, sim-verified) [crafted]; Taran Icebreaker (2915, -1.16 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Red Mageweave Headband; neck: Skibi's Pendant; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Robes of Insight; wrist: Runic Leather Bracers; hands: Stonerender Gauntlets; waist: Defiler's Lamellar Girdle; legs: Infernal Trickster Leggings; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Warmonger

No-known-source sample (15 of 1047, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 364.4. Weights run: 7.8s. Verify run: 91.2s. 2461 eligible items had no known source.

Race sweep: undead 364.4 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=2.009 ± 0.045, crit=3.085 ± 0.064 per rating point (14 rating = 1%, 43.194 per %), hit=4.212 ± 0.287 per rating point (10 rating = 1%, 42.116 per %), melee_haste=not significant (12.556 ± 3.683), intellect=7.304 ± 0.137, spell_power=4.382 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 362.9 attack_power points (21.81 DPS) | yes | Stalwart Helm (250599, -5.15 DPS) [crafted]; Blue Dragonscale Helm (252604, -5.74 DPS, sim-verified) [crafted]; Magister's Crown (16686, -5.74 DPS) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 187.0 attack_power points (11.23 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -1.36 DPS, sim-verified) [world_drop]; Nacreous Shell Necklace (22403, -2.82 DPS) [dungeon]; Archlight Talisman (15856, -3.69 DPS) [quest] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 267.3 attack_power points (16.06 DPS) | yes | Darkspear Shoulderpads (272103, -1.05 DPS) [vendor]; Darkspear Shoulders (272104, -1.05 DPS) [vendor]; Mantle of the Timbermaw (19050, -3.29 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of the Exile (15421, -2.71 DPS) [quest]; Shroud of Arcane Mastery (22330, -2.90 DPS) [dungeon]; Royal Tribunal Cloak (13376, -3.82 DPS, sim-verified) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 257.1 attack_power points (15.45 DPS) | yes | Alanna's Embrace (13314, -0.61 DPS) [dungeon]; Vest of Elements (16666, -0.61 DPS) [dungeon]; Blue Dragonscale Breastplate (15048, -1.05 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 154.8 attack_power points (9.30 DPS) | yes | Funeral Cuffs (12626, -1.58 DPS) [dungeon]; Aristocratic Cuffs (12546, -1.67 DPS) [dungeon]; Manacle Cuffs (11962, -2.09 DPS, sim-verified) [quest] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Raider Handwraps (272098, +0.00 DPS) [vendor]; Raider Handguards (272101, +0.00 DPS) [vendor] |
| waist | Crystallized Girdle (12606) | Blackrock Spire: Solakar Flamewreath [dungeon] | 195.7 attack_power points (11.76 DPS) | yes | Magister's Belt (16685, -0.70 DPS) [dungeon]; Lightforge Belt (16723, -0.75 DPS) [dungeon]; Clutch of Andros (13956, -0.89 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Lizardhide Pants (22741, -3.83 DPS) [rep]; Outrider's Silk Leggings (22747, -4.79 DPS) [rep]; Red Dragonscale Leggings (252603, -4.87 DPS, sim-verified) [crafted] |
| feet | Dragonrider Boots (18102) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 195.7 attack_power points (11.76 DPS) | yes | Kayser's Boots of Precision (22231, -0.28 DPS) [dungeon]; Omnicast Boots (11822, -1.23 DPS) [dungeon]; Greaves of Withering Despair (22240, -1.95 DPS) [dungeon] |
| finger1 | Seal of Rivendare (13345) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.76 DPS) [quest]; Maiden's Circle (13001, -1.76 DPS) [world_drop]; Naglering (11669, -6.78 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.86 DPS) [quest]; Deep Rooted Ring (19109, -1.13 DPS) [world]; Maiden's Circle (13001, -1.47 DPS, sim-verified) [world_drop] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -2.13 DPS) [quest]; Smolderweb's Eye (13213, -2.21 DPS) [dungeon]; Hand of Justice (11815, -4.75 DPS, sim-verified) [dungeon] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -21.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, -4.69 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Darkspear Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Storm Gauntlets; waist: Crystallized Girdle; legs: Ironfeather Leggings; feet: Dragonrider Boots; finger1: Seal of Rivendare; finger2: Channeler's Ring; trinket1: Blackhand's Breadth; trinket2: Weakness Analyzer; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2461, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (undead, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 876.4. Weights run: 7.9s. Verify run: 102.2s. 2461 eligible items had no known source.

Race sweep: undead 876.4 ± 1.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=1.670 ± 0.032, crit=2.253 ± 0.044 per rating point (14 rating = 1%, 31.546 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=not significant (14.247 ± 3.704), intellect=1.442 ± 0.053, spell_power=1.879 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | sim-verified (+22.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Stalwart Helm (250599, -1.25 DPS) [crafted]; Eye of Rend (12587, -4.65 DPS) [dungeon]; Black Dragonscale Helm (252605, -22.13 DPS, sim-verified) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 57.5 attack_power points (11.61 DPS) | yes | Lady Maye's Pendant (14558, -1.91 DPS) [world_drop]; Emberfury Talisman (12929, -2.97 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -4.10 DPS) [world_drop] |
| shoulder | Defiler's Lamellar Spaulders (20184) | The Defilers [rep] | 86.4 attack_power points (17.42 DPS) | yes | Mantle of the Timbermaw (19050, -0.83 DPS) [crafted]; Darkspear Pauldrons (272105, -1.96 DPS) [vendor]; Darkspear Epaulets (272106, -1.96 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -0.34 DPS) [world]; Deathguard's Cloak (20068, -0.54 DPS) [rep]; Arcanoweave Cloak (272411, -0.69 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 141.5 attack_power points (28.53 DPS) | yes | Savage Gladiator Chain (11726, -5.24 DPS, sim-verified) [dungeon]; Breastplate of Bloodthirst (12757, -9.46 DPS) [quest]; Dawn Armor (252483, -9.71 DPS) [crafted] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | 59.3 attack_power points (11.97 DPS) | yes | Forest Stalker's Bracers (19587, -0.20 DPS) [rep]; Dryad's Wrist Bindings (19595, -1.30 DPS) [rep] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Raider Gauntlets (272095, +0.00 DPS) [vendor] |
| waist | Vindicator's Belt (272404) | Pix Xizzix [vendor] | 88.2 attack_power points (17.78 DPS) | yes | Defiler's Lamellar Girdle (20177, -1.19 DPS) [rep]; Might of the Timbermaw (19044, -2.81 DPS) [crafted]; Defiler's Plate Girdle (20204, -3.12 DPS) [rep] |
| legs | Champion's Legplates (250598) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pristine Scorpid Leggings (252606, -1.90 DPS) [crafted]; Ironfeather Leggings (252486, -3.55 DPS) [crafted]; Outrider's Lizardhide Pants (22741, -4.88 DPS) [rep] |
| feet | Defiler's Lamellar Greaves (20181) | The Defilers [rep] | 65.4 attack_power points (13.20 DPS) | yes | Death Knight Sabatons (18692, -0.46 DPS) [dungeon]; Dragonrider Boots (18102, -1.72 DPS) [dungeon]; Skirmisher's Mail Sabatons (252578, -1.93 DPS) [crafted] |
| finger1 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -0.97 DPS) [rep]; Maiden's Circle (13001, -1.67 DPS) [world_drop]; Naglering (11669, -10.78 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (+12.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Orgrimmar (12545, -0.33 DPS) [quest]; Maiden's Circle (13001, -0.33 DPS) [world_drop]; Legionnaire's Band (19510, -12.55 DPS, sim-verified) [rep] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (+10.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -0.17 DPS) [vendor]; Hand of Justice (11815, -0.40 DPS) [dungeon] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Unstoppable Force (19323, +0.00 DPS) [rep]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, +0.00 DPS) [vendor] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Defiler's Lamellar Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Storm Gauntlets; waist: Vindicator's Belt; legs: Champion's Legplates; feet: Defiler's Lamellar Greaves; finger1: Don Mauricio's Band of Domination; finger2: Don Julio's Band; trinket1: Blackhand's Breadth; trinket2: Rune of the Guard Captain; main_hand: Blackblade of Shahram

No-known-source sample (15 of 2461, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

