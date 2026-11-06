# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 46.1. Weights run: 8.1s. Verify run: 21.2s. 281 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.306 ± 0.019, crit=0.777 ± 0.019 per rating point (14 rating = 1%, 10.880 per %), hit=1.284 ± 0.049 per rating point (10 rating = 1%, 12.838 per %), melee_haste=6.394 ± 0.679, intellect=1.508 ± 0.038, spell_power=2.333 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 25.7 attack_power points (1.12 DPS) | yes | Totemic Leather Hood (252448, -0.20 DPS) [crafted]; Veteran's Silvered Chain Helm (250528, -0.25 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.41 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 7.8 attack_power points (0.34 DPS) | yes | Erudite's Amulet (277204, -0.23 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.60 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 25.2 attack_power points (1.10 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.56 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.69 DPS) [crafted]; Serpent's Shoulders (5404, -0.81 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.01 DPS) [crafted]; Lambent Scale Cloak (4706, -0.06 DPS) [world_drop]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 32.5 attack_power points (1.42 DPS) | yes | Crusader's Chain Shirt (250492, -0.15 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.28 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.55 DPS) [dungeon] |
| wrist | Death Bindings (286981) | Ruins of Lordaeron: Skeleton [dungeon] | sim-verified (46.1 DPS) | yes | Bravo's Armbands (270015, -0.00 DPS) [quest]; Patterned Bronze Bracers (2868, -0.09 DPS, sim-verified) [crafted]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 24.2 attack_power points (1.05 DPS) | yes | Totemic Leather Gloves (252497, -0.20 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.28 DPS) [crafted]; Crusader's Gloves (250512, -0.28 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.78 DPS) | yes | Cobrahn's Grasp (6460, -0.00 DPS) [dungeon]; Crusader's Chain Belt (250517, -0.01 DPS) [crafted]; Totemic Leather Belt (252431, -0.03 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 40.4 attack_power points (1.76 DPS) | yes | Crusader's Chain Leggings (250497, -0.38 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.42 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, -0.46 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 24.0 attack_power points (1.05 DPS) | yes | Spidersilk Boots (4320, -0.07 DPS) [crafted]; Crusader's Boots (250507, -0.09 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.11 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 14.7 attack_power points (0.64 DPS) | yes | Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Demon Band (12054, -0.29 DPS) [world_drop]; Protector's Band (20439, -0.74 DPS, sim-verified) [rep] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.00 DPS) [dungeon]; Demon Band (12054, -0.16 DPS) [world_drop]; Protector's Band (20439, -0.36 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sword of the Fallen (286977, -2.89 DPS) [dungeon]; Duskbringer (2205, -3.20 DPS) [dungeon]; Smite's Mighty Hammer (7230, -4.14 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Totemic Leather Armor; wrist: Death Bindings; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Verigan's Fist

No-known-source sample (15 of 281, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 30 (human, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 81.8. Weights run: 8.7s. Verify run: 16.1s. 504 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.529 ± 0.025, crit=1.191 ± 0.026 per rating point (14 rating = 1%, 16.680 per %), hit=1.811 ± 0.072 per rating point (10 rating = 1%, 18.112 per %), melee_haste=9.597 ± 0.838, intellect=1.427 ± 0.042, spell_power=2.688 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 59.6 attack_power points (2.82 DPS) | yes | Barbaric Iron Helm (7915, -1.32 DPS) [crafted]; Enchanter's Cowl (4322, -1.38 DPS) [crafted]; Totemic Leather Helm (252456, -1.43 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 20.6 attack_power points (0.97 DPS) | yes | Kaleidoscope Chain (13084, -0.03 DPS) [world_drop]; Sentinel's Medallion (19541, -0.39 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 37.0 attack_power points (1.75 DPS) | yes | Death Speaker Mantle (6685, -0.46 DPS, sim-verified) [dungeon]; Forest Tracker Epaulets (2278, -0.48 DPS) [world_drop]; Magician's Mantle (12998, -0.51 DPS) [world_drop] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (81.8 DPS) | yes | Hawkeye's Cloak (14593, -0.15 DPS) [world_drop]; Hillman's Cloak (3719, -0.30 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.30 DPS) [quest] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 44.3 attack_power points (2.09 DPS) | yes | Totemic Leather Tunic (252451, -0.11 DPS) [crafted]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Green Silk Armor (7065, -0.37 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 24.2 attack_power points (1.14 DPS) | yes | Hawkeye's Bracers (14590, -0.33 DPS) [world_drop]; Yorgen Bracers (13012, -0.36 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.16 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 28.0 attack_power points (1.32 DPS) | yes | Mail Combat Gauntlets (4075, -0.28 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.28 DPS) [world]; Insignia Gloves (6408, -0.35 DPS, sim-verified) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 44.3 attack_power points (2.09 DPS) | yes | Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Skirmisher's Leather Belt (252461, -0.49 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -0.49 DPS) [rep] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 48.9 attack_power points (2.31 DPS) | yes | Totemic Leather Leggings (252458, -0.04 DPS) [crafted]; Ferine Leggings (6690, -0.28 DPS) [dungeon]; Totemic Leather Pants (252446, -0.30 DPS) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 32.8 attack_power points (1.55 DPS) | yes | Totemic Leather Boots (252442, -0.32 DPS) [crafted]; Disjointed Shoes (277226, -0.35 DPS) [quest]; Gilded Slippers (254001, -0.48 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 23.3 attack_power points (1.10 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Lorekeeper's Ring (19525, -0.21 DPS) [rep]; Seal of Wrynn (2933, -0.33 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.2 attack_power points (1.00 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Lorekeeper's Ring (19525, -0.11 DPS) [rep]; Seal of Wrynn (2933, -0.23 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (81.8 DPS) | yes | Felbough Arm (273037, -0.41 DPS) [world]; Corpsemaker (6687, -0.49 DPS) [dungeon]; Viscous Hammer (13045, -5.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Sergeant Major's Cape; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 504, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 134.1. Weights run: 8.8s. Verify run: 15.9s. 691 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=2.642 ± 0.054, crit=2.671 ± 0.064 per rating point (14 rating = 1%, 37.399 per %), hit=3.116 ± 0.127 per rating point (10 rating = 1%, 31.160 per %), melee_haste=14.857 ± 1.555, intellect=3.093 ± 0.055, spell_power=2.797 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 75.0 attack_power points (3.96 DPS) | yes | Corpseshroud (10574, -0.16 DPS) [dungeon]; Big Voodoo Mask (8201, -0.34 DPS) [crafted]; Raging Berserker's Helm (7719, -0.58 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 44.3 attack_power points (2.34 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.71 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.73 DPS) [world]; Kaleidoscope Chain (13084, -1.38 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 59.8 attack_power points (3.16 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.46 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.76 DPS) [crafted]; Guardian Cloak (5965, -0.76 DPS) [crafted]; Expeditionary's Cape (271922, -0.92 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.10 DPS) [crafted]; Robe of Power (7054, -0.20 DPS) [crafted]; Icemail Jerkin (1981, -0.40 DPS, sim-verified) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 37.1 attack_power points (1.96 DPS) | yes | Guardian Leather Bracers (4260, -0.09 DPS) [crafted]; Windchaser Cuffs (14429, -0.21 DPS) [world_drop]; Warden's Wristbands (14600, -0.23 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, -1.83 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 66.0 attack_power points (3.48 DPS) | yes | Skirmisher's Leather Belt (252461, -0.83 DPS) [crafted]; Gilded Cord (254037, -1.00 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.95 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 76.3 attack_power points (4.03 DPS) | yes | Ferine Leggings (6690, -0.71 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -0.72 DPS) [crafted]; Totemic Leather Leggings (252458, -0.90 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 67.1 attack_power points (3.54 DPS) | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.48 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.48 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 43.4 attack_power points (2.29 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Ring of the Underwood (2951, -0.68 DPS) [world_drop]; Falcon's Hook (7552, -0.71 DPS) [dungeon] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 37.5 attack_power points (1.98 DPS) | yes | Snake Hoop (6750, -0.27 DPS) [quest]; Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, -0.11 DPS) [vendor]; Nightblade (1982, -1.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Sergeant Major's Cape; chest: Robe of the Magi; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Protector's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 691, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 188.8. Weights run: 8.9s. Verify run: 41.0s. 883 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.969 ± 0.061, crit=3.495 ± 0.078 per rating point (14 rating = 1%, 48.926 per %), hit=2.856 ± 0.150 per rating point (10 rating = 1%, 28.564 per %), melee_haste=16.580 ± 1.739, intellect=4.153 ± 0.096, spell_power=2.676 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Soothsayer's Headdress (17740, -0.20 DPS) [dungeon]; Knight-Lieutenant's Lamellar Helm (220819, -0.36 DPS) [vendor]; Red Mageweave Headband (10033, -0.74 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 58.1 attack_power points (2.93 DPS) | yes | Glowing Eye of Mordresh (10769, -0.31 DPS) [dungeon]; Skibi's Pendant (13089, -0.43 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.84 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 123.2 attack_power points (6.20 DPS) | yes | Rotgrip Mantle (17732, -1.12 DPS, sim-verified) [dungeon]; Lead Surveyor's Mantle (11842, -1.25 DPS) [dungeon]; Knight-Lieutenant's Lamellar Pauldrons (220818, -1.28 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -0.21 DPS) [rep]; Runecloth Cloak (13860, -0.26 DPS) [crafted]; Sergeant Major's Cape (16336, -0.32 DPS) [pvp] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 133.9 attack_power points (6.74 DPS) | yes | Spiderfang Carapace (11633, -0.72 DPS) [dungeon]; Knight's Lamellar Chestplate (220815, -0.98 DPS) [vendor]; Robes of Insight (940, -1.51 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 68.3 attack_power points (3.44 DPS) | yes | Aristocratic Cuffs (12546, -0.30 DPS) [dungeon]; Skirmisher's Mail Bracers (252572, -0.32 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Sergeant Major's Lamellar Gauntlets (220817, +0.00 DPS) [vendor]; Skirmisher's Mail Gauntlets (252586, +0.00 DPS) [crafted]; Stonerender Gauntlets (17007, -2.32 DPS, sim-verified) [world_drop] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ban'thok Sash (11662, -0.54 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -0.57 DPS) [dungeon]; Sagebrush Girdle (17778, -0.60 DPS) [quest] |
| legs | Knight's Lamellar Legplates (220816) | Captain Dirgehammer [vendor] | 114.4 attack_power points (5.76 DPS) | yes | Imperial Red Pants (8251, -0.38 DPS) [world_drop]; Knight's Plate Leggings (220797, -0.47 DPS) [vendor]; Spellshock Leggings (9484, -0.57 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 99.5 attack_power points (5.01 DPS) | yes | Skycaller's Leather Boots (252471, -0.69 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.69 DPS) [crafted]; Fleetfoot Greaves (11627, -0.92 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (188.8 DPS) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon] |
| trinket1 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.17 DPS, sim-verified) [world_drop] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Thorium Greatmace (250613, -0.48 DPS, sim-verified) [crafted]; The Rockpounder (9413, -0.79 DPS) [dungeon]; Stoneraven (13059, -1.37 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Highlander's Lamellar Girdle; legs: Knight's Lamellar Legplates; feet: Greaves of Withering Despair; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Sanctified Orb; trinket2: Thunderbrew's Boot Flask; main_hand: Warmonger

No-known-source sample (15 of 883, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 336.1. Weights run: 8.9s. Verify run: 131.1s. 2115 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=3.712 ± 0.085, crit=5.376 ± 0.121 per rating point (14 rating = 1%, 75.264 per %), hit=4.311 ± 0.225 per rating point (10 rating = 1%, 43.108 per %), melee_haste=22.913 ± 2.589, intellect=4.624 ± 0.181, spell_power=2.825 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 328.9 attack_power points (16.92 DPS) | yes | Field Marshal's Lamellar Faceguard (16474, -2.28 DPS) [vendor]; Field Marshal's Lamellar Headguard (231648, -3.94 DPS) [vendor]; Lionheart Helm (12640, -8.96 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Pauldrons (271955, +0.00 DPS) [rep]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Mantle of the Timbermaw (19050, -4.73 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16337, -2.11 DPS) [pvp]; Shroud of Arcane Mastery (22330, -2.29 DPS, sim-verified) [dungeon]; Cape of the Black Baron (13340, -2.55 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -8.63 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spitfire Bracers (20481, -0.08 DPS) [crafted]; Forest Stalker's Bracers (19587, -0.23 DPS) [rep]; Battleborn Armbraces (12936, -3.50 DPS, sim-verified) [dungeon] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Mourning Handguards (279257, -2.83 DPS, sim-verified) [crafted] |
| waist | Highlander's Lamellar Girdle (20042) | The League of Arathor [rep] | 154.5 attack_power points (7.95 DPS) | yes | Highlander's Mail Girdle (20044, -0.03 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, -0.03 DPS) [rep]; Highlander's Cloth Girdle (20047, -0.61 DPS) [rep] |
| legs | Sentinel's Lamellar Legguards (237814) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Leggings (231647, -1.57 DPS) [vendor]; Sentinel's Silk Leggings (237815, -1.61 DPS) [vendor]; Sentinel's Chain Leggings (237819, -6.55 DPS, sim-verified) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (336.1 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.34 DPS) [dungeon]; Cutthroat's Signet (272408, -1.92 DPS) [vendor]; Naglering (11669, -7.39 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.66 DPS) [vendor]; Painweaver Band (13098, -0.85 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -2.56 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -0.71 DPS) [dungeon]; Second Wind (11819, -2.83 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.16 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.20 DPS) [crafted]; Thunderbrew's Boot Flask (744, -2.33 DPS, sim-verified) [quest] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -12.50 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, -4.97 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Highlander's Lamellar Girdle; legs: Sentinel's Lamellar Legguards; finger1: Don Julio's Band; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2115, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60, raid preset (human, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 841.8. Weights run: 9.6s. Verify run: 126.9s. 2115 eligible items had no known source.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.006, agility=2.523 ± 0.052, crit=3.366 ± 0.071 per rating point (14 rating = 1%, 47.123 per %), hit=6.021 ± 0.370 per rating point (10 rating = 1%, 60.213 per %), melee_haste=32.736 ± 4.634, intellect=1.776 ± 0.112, spell_power=1.828 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 258.2 attack_power points (39.25 DPS) | yes | Stalwart Helm (250599, -5.97 DPS) [crafted]; Field Marshal's Lamellar Faceguard (16474, -12.02 DPS) [vendor]; Black Dragonscale Helm (252605, -13.78 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -2.35 DPS) [vendor]; Shroud of Arcane Mastery (22330, -3.63 DPS) [dungeon]; Cloak of Flames (3475, -6.60 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -0.62 DPS) [dungeon]; Field Marshal's Lamellar Breastplate (231649, -0.78 DPS) [vendor]; Dawn Armor (252483, -9.53 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, -0.34 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor] |
| waist | Highlander's Lamellar Girdle (20042) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Brigam Girdle (13142, -1.08 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -3.49 DPS) [vendor]; Sentinel's Lamellar Legguards (22753, -5.91 DPS) [rep] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (841.8 DPS) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Knight-Lieutenant's Lamellar Sabatons (227146, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -5.51 DPS) [vendor]; Blackstone Ring (17713, -6.25 DPS) [dungeon]; Naglering (11669, -9.85 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -1.59 DPS) [vendor]; Blackstone Ring (17713, -2.33 DPS) [dungeon]; Blessed Band of Light (272407, -2.59 DPS) [vendor] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.37 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Serenity Field (272439, -2.22 DPS) [vendor]; Burst of Knowledge (11832, -5.32 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackblade of Shahram (12592, +0.00 DPS) [dungeon]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, +0.00 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Skibi's Pendant; shoulder: Theramore Pauldrons; back: Arcanoweave Cloak; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Highlander's Lamellar Girdle; legs: Sentinel's Chain Leggings; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Weakness Analyzer; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 2115, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (undead, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 42.9. Weights run: 8.1s. Verify run: 21.3s. 268 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.306 ± 0.019, crit=0.777 ± 0.019 per rating point (14 rating = 1%, 10.880 per %), hit=1.284 ± 0.049 per rating point (10 rating = 1%, 12.838 per %), melee_haste=6.394 ± 0.679, intellect=1.508 ± 0.038, spell_power=2.333 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 25.7 attack_power points (1.12 DPS) | yes | Totemic Leather Hood (252448, -0.20 DPS) [crafted]; Veteran's Silvered Chain Helm (250528, -0.25 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.41 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 7.8 attack_power points (0.34 DPS) | yes | Erudite's Amulet (277204, -0.23 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 25.2 attack_power points (1.10 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.33 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.69 DPS) [crafted]; Serpent's Shoulders (5404, -0.81 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.01 DPS) [crafted]; Lambent Scale Cloak (4706, -0.06 DPS) [world_drop]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 32.5 attack_power points (1.42 DPS) | yes | Crusader's Chain Shirt (250492, -0.15 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.28 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.55 DPS) [dungeon] |
| wrist | Death Bindings (286981) | Ruins of Lordaeron: Skeleton [dungeon] | sim-verified (42.9 DPS) | yes | Patterned Bronze Bracers (2868, +0.00 DPS) [crafted]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop]; Owl Bracers (4796, -0.11 DPS) [vendor] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 24.2 attack_power points (1.05 DPS) | yes | Totemic Leather Gloves (252497, -0.20 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.28 DPS) [crafted]; Crusader's Gloves (250512, -0.28 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.78 DPS) | yes | Cobrahn's Grasp (6460, -0.00 DPS) [dungeon]; Crusader's Chain Belt (250517, -0.01 DPS) [crafted]; Totemic Leather Belt (252431, -0.03 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 40.4 attack_power points (1.76 DPS) | yes | Stormrider's Leather Pants (252502, -0.35 DPS) [crafted]; Crusader's Chain Leggings (250497, -0.38 DPS) [crafted]; Trapper's Leather Pants (252501, -0.46 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 24.0 attack_power points (1.05 DPS) | yes | Spidersilk Boots (4320, -0.07 DPS) [crafted]; Crusader's Boots (250507, -0.09 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.11 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | sim-verified (42.9 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.07 DPS) [dungeon]; Demon Band (12054, -0.23 DPS) [world_drop] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Loop of Sacrifice (281673, +0.00 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.00 DPS) [dungeon]; Demon Band (12054, -0.16 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Sword of the Fallen (286977, -0.06 DPS) [dungeon]; Hammerbone (270018, -0.12 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Totemic Leather Armor; wrist: Death Bindings; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Legionnaire's Band; finger2: Advisor's Ring; main_hand: Forsaken Greataxe

No-known-source sample (15 of 268, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (undead, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 83.0. Weights run: 8.7s. Verify run: 16.2s. 491 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.529 ± 0.025, crit=1.191 ± 0.026 per rating point (14 rating = 1%, 16.680 per %), hit=1.811 ± 0.072 per rating point (10 rating = 1%, 18.112 per %), melee_haste=9.597 ± 0.838, intellect=1.427 ± 0.042, spell_power=2.688 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 59.6 attack_power points (2.82 DPS) | yes | Totemic Leather Helm (252456, -1.13 DPS, sim-verified) [crafted]; Barbaric Iron Helm (7915, -1.32 DPS) [crafted]; Enchanter's Cowl (4322, -1.38 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 20.6 attack_power points (0.97 DPS) | yes | Kaleidoscope Chain (13084, -0.33 DPS, sim-verified) [world_drop]; Scout's Medallion (19537, -0.39 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 37.0 attack_power points (1.75 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS, sim-verified) [dungeon]; Forest Tracker Epaulets (2278, -0.48 DPS) [world_drop]; Magician's Mantle (12998, -0.51 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (83.0 DPS) | yes | Hillman's Cloak (3719, -0.15 DPS) [crafted]; Windsong Drape (15468, -0.15 DPS) [quest]; Tigerstrike Mantle (13108, -0.21 DPS) [world_drop] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 44.3 attack_power points (2.09 DPS) | yes | Totemic Leather Tunic (252451, -0.11 DPS) [crafted]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Green Silk Armor (7065, -0.78 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 24.2 attack_power points (1.14 DPS) | yes | Hawkeye's Bracers (14590, -0.33 DPS) [world_drop]; Yorgen Bracers (13012, -0.36 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.22 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 28.0 attack_power points (1.32 DPS) | yes | Oilrag Handwraps (16741, -0.04 DPS) [quest]; Jutebraid Gloves (10654, -0.22 DPS) [quest]; Insignia Gloves (6408, -0.23 DPS) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 44.3 attack_power points (2.09 DPS) | yes | Skirmisher's Leather Belt (252461, -0.08 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.49 DPS) [rep] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 48.9 attack_power points (2.31 DPS) | yes | Ferine Leggings (6690, -0.28 DPS) [dungeon]; Totemic Leather Pants (252446, -0.30 DPS) [crafted]; Totemic Leather Leggings (252458, -0.57 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 32.8 attack_power points (1.55 DPS) | yes | Totemic Leather Boots (252442, -0.32 DPS) [crafted]; Acidic Walkers (9454, -0.38 DPS) [dungeon]; Gilded Slippers (254001, -0.47 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 23.3 attack_power points (1.10 DPS) | yes | Thunderbrow Ring (13097, -0.13 DPS) [world_drop]; Advisor's Ring (19521, -0.21 DPS) [rep]; Ring of Power Regulation (273030, -0.34 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.2 attack_power points (1.00 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Advisor's Ring (19521, -0.11 DPS) [rep]; Ring of Power Regulation (273030, -0.24 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (83.0 DPS) | yes | Felbough Arm (273037, -0.41 DPS) [world]; Corpsemaker (6687, -0.49 DPS) [dungeon]; Viscous Hammer (13045, -5.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hawkeye's Cloak; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 491, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 132.7. Weights run: 8.8s. Verify run: 19.5s. 675 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=2.642 ± 0.054, crit=2.671 ± 0.064 per rating point (14 rating = 1%, 37.399 per %), hit=3.116 ± 0.127 per rating point (10 rating = 1%, 31.160 per %), melee_haste=14.857 ± 1.555, intellect=3.093 ± 0.055, spell_power=2.797 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 75.0 attack_power points (3.96 DPS) | yes | Big Voodoo Mask (8201, -0.34 DPS) [crafted]; Corpseshroud (10574, -0.41 DPS, sim-verified) [dungeon]; Raging Berserker's Helm (7719, -0.58 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 44.3 attack_power points (2.34 DPS) | yes | Kaleidoscope Chain (13084, -0.70 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.71 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.73 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 59.8 attack_power points (3.16 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.46 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (132.7 DPS) | yes | Long Silken Cloak (4326, -0.09 DPS) [crafted]; Guardian Cloak (5965, -0.09 DPS) [crafted]; Blackforge Cape (6424, -0.49 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 94.6 attack_power points (4.99 DPS) | yes | Robe of the Magi (1716, -0.76 DPS) [world_drop]; Dreamweave Vest (10021, -0.86 DPS) [crafted]; Robe of Power (7054, -0.97 DPS) [crafted] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 37.1 attack_power points (1.96 DPS) | yes | Radiant Silver Bracers (4545, -0.06 DPS) [quest]; Guardian Leather Bracers (4260, -0.09 DPS) [crafted]; Windchaser Cuffs (14429, -0.21 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (132.7 DPS) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, -1.25 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 66.0 attack_power points (3.48 DPS) | yes | Defiler's Cloth Girdle (20166, -0.76 DPS) [rep]; Skirmisher's Leather Belt (252461, -0.83 DPS) [crafted]; Gilded Cord (254037, -1.00 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 76.3 attack_power points (4.03 DPS) | yes | Ferine Leggings (6690, -0.69 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.72 DPS) [crafted]; Totemic Leather Leggings (252458, -0.90 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 67.1 attack_power points (3.54 DPS) | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.48 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.48 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 43.4 attack_power points (2.29 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Ring of the Underwood (2951, -0.68 DPS) [world_drop]; Falcon's Hook (7552, -0.71 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 37.5 attack_power points (1.98 DPS) | yes | Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon]; Snake Hoop (6750, -0.72 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (132.7 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, -0.11 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Legionnaire's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 675, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 187.8. Weights run: 8.9s. Verify run: 38.6s. 883 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.969 ± 0.061, crit=3.495 ± 0.078 per rating point (14 rating = 1%, 48.926 per %), hit=2.856 ± 0.150 per rating point (10 rating = 1%, 28.564 per %), melee_haste=16.580 ± 1.739, intellect=4.153 ± 0.096, spell_power=2.676 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Soothsayer's Headdress (17740, -0.20 DPS) [dungeon]; Chief Architect's Monocle (11839, -0.79 DPS) [dungeon]; Red Mageweave Headband (10033, -1.04 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 58.1 attack_power points (2.93 DPS) | yes | Glowing Eye of Mordresh (10769, -0.31 DPS) [dungeon]; Skibi's Pendant (13089, -0.43 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.84 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 123.2 attack_power points (6.20 DPS) | yes | Rotgrip Mantle (17732, -0.86 DPS, sim-verified) [dungeon]; Lead Surveyor's Mantle (11842, -1.25 DPS) [dungeon]; Rockshard Pauldrons (9411, -1.54 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.36 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -0.57 DPS) [vendor]; Runecloth Cloak (13860, -0.61 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 133.9 attack_power points (6.74 DPS) | yes | Spiderfang Carapace (11633, -0.72 DPS) [dungeon]; Robes of Insight (940, -1.51 DPS) [world_drop]; Runecloth Robe (13858, -1.70 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 68.3 attack_power points (3.44 DPS) | yes | Aristocratic Cuffs (12546, -0.30 DPS) [dungeon]; Skirmisher's Mail Bracers (252572, -0.32 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.35 DPS) [crafted] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skirmisher's Mail Gauntlets (252586, -0.60 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -0.74 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.85 DPS) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Dawnspire Cord (12466, -0.57 DPS) [dungeon]; Sagebrush Girdle (17778, -0.60 DPS) [quest]; Ban'thok Sash (11662, -0.61 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Spellshock Leggings (9484, -0.10 DPS) [dungeon]; Stormshroud Pants (15057, -0.36 DPS) [crafted]; Imperial Red Pants (8251, -1.00 DPS, sim-verified) [world_drop] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 99.5 attack_power points (5.01 DPS) | yes | Skycaller's Mail Sabatons (252577, -0.69 DPS) [crafted]; Fleetfoot Greaves (11627, -0.92 DPS) [dungeon]; Skycaller's Leather Boots (252471, -1.23 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 62.3 attack_power points (3.14 DPS) | yes | Mindseye Circle (10634, -0.63 DPS) [dungeon]; Blackstone Ring (17713, -0.69 DPS) [dungeon]; Cyclopean Band (11824, -3.15 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | sim-verified (187.8 DPS) | yes | Mindseye Circle (10634, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon]; Blackstone Ring (17713, -0.01 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.49 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.95 DPS, sim-verified) [world_drop] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | The Rockpounder (9413, -0.79 DPS) [dungeon]; Thorium Greatmace (250613, -0.82 DPS, sim-verified) [crafted]; Stoneraven (13059, -1.37 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Stonerender Gauntlets; waist: Defiler's Lamellar Girdle; legs: Stone Guard's Plate Leggings; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Legionnaire's Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Warmonger

No-known-source sample (15 of 883, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 336.9. Weights run: 8.9s. Verify run: 129.1s. 2141 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=3.712 ± 0.085, crit=5.376 ± 0.121 per rating point (14 rating = 1%, 75.264 per %), hit=4.311 ± 0.225 per rating point (10 rating = 1%, 43.108 per %), melee_haste=22.913 ± 2.589, intellect=4.624 ± 0.181, spell_power=2.825 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 328.9 attack_power points (16.92 DPS) | yes | Stalwart Helm (250599, -5.61 DPS) [crafted]; Blue Dragonscale Helm (252604, -6.79 DPS) [crafted]; Lionheart Helm (12640, -8.25 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -3.04 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -2.52 DPS, sim-verified) [dungeon]; Cape of the Black Baron (13340, -2.55 DPS) [dungeon]; Deep Woodlands Cloak (19121, -2.56 DPS) [quest] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, +0.00 DPS) [dungeon]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -6.84 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spitfire Bracers (20481, -0.08 DPS) [crafted]; Forest Stalker's Bracers (19587, -0.23 DPS) [rep]; Battleborn Armbraces (12936, -3.23 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handguards (272101, -0.26 DPS) [vendor]; Bloodmail Gauntlets (14615, -0.32 DPS) [dungeon]; Mourning Handguards (279257, -3.59 DPS, sim-verified) [crafted] |
| waist | Defiler's Lizardhide Girdle (20171) | The Defilers [rep] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Mail Girdle (20195, +0.00 DPS) [rep]; Defiler's Cloth Girdle (20163, -0.58 DPS) [rep]; Defiler's Lamellar Girdle (20177, -0.96 DPS, sim-verified) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Legplates (250598, -2.24 DPS) [crafted]; Outrider's Lizardhide Pants (22741, -2.44 DPS) [rep]; Outrider's Chain Leggings (22673, -9.48 DPS, sim-verified) [rep] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (336.9 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Windreaver Greaves (13967, -0.16 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.34 DPS) [dungeon]; Cutthroat's Signet (272408, -1.92 DPS) [vendor]; Naglering (11669, -8.24 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.66 DPS) [vendor]; Painweaver Band (13098, -0.85 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -3.93 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Second Wind (11819, -2.83 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.92 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.20 DPS) [crafted]; Serenity Field (272439, -1.97 DPS, sim-verified) [vendor] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -14.19 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, -4.83 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Defiler's Lizardhide Girdle; legs: Ironfeather Leggings; finger1: Don Julio's Band; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2141, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (undead, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 832.5. Weights run: 9.6s. Verify run: 136.6s. 2141 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=1.780 ± 0.034, crit=2.404 ± 0.045 per rating point (14 rating = 1%, 33.649 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=28.356 ± 3.927, intellect=1.462 ± 0.080, spell_power=1.774 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Black Dragonscale Helm (252605, +0.00 DPS) [crafted]; Eye of Rend (12587, -2.28 DPS) [dungeon]; Stalwart Helm (250599, -2.71 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -0.13 DPS) [world_drop]; Emberfury Talisman (12929, -0.43 DPS) [dungeon] |
| shoulder | Defiler's Lamellar Spaulders (20184) | The Defilers [rep] | 88.4 attack_power points (16.67 DPS) | yes | Mantle of the Timbermaw (19050, -1.06 DPS) [crafted]; Darkspear Pauldrons (272105, -1.27 DPS) [vendor]; Darkspear Epaulets (272106, -1.27 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.72 DPS) [rep]; Crystalline Threaded Cape (20697, -1.01 DPS) [world]; Arcanoweave Cloak (272411, -1.25 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 145.5 attack_power points (27.44 DPS) | yes | Savage Gladiator Chain (11726, -2.75 DPS) [dungeon]; Dawn Armor (252483, -8.68 DPS) [crafted]; Breastplate of Bloodthirst (12757, -8.82 DPS) [quest] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (832.5 DPS) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Tranquil Wristguards (279256, -1.30 DPS) [crafted] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (832.5 DPS) | yes | Timbermaw Brawlers (19049, -1.73 DPS) [crafted]; Raider Gloves (272099, -1.93 DPS) [vendor]; Stonerender Gauntlets (17007, -1.94 DPS) [world_drop] |
| waist | Defiler's Lamellar Girdle (20177) | The Defilers [rep] | sim-verified (832.5 DPS) | yes | Vindicator's Belt (272404, +0.00 DPS) [vendor]; Might of the Timbermaw (19044, -1.66 DPS) [crafted]; Defiler's Plate Girdle (20204, -1.84 DPS) [rep] |
| legs | Champion's Legplates (250598) | Blacksmithing [crafted] | sim-verified (832.5 DPS) | yes | Pristine Scorpid Leggings (252606, -1.67 DPS) [crafted]; Ironfeather Leggings (252486, -3.25 DPS) [crafted]; Outrider's Lizardhide Pants (22741, -4.13 DPS) [rep] |
| feet | Depleted Thorium Sabatons (279263) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Lamellar Greaves (20181, -0.79 DPS) [rep]; Battlechaser's Greaves (12555, -2.66 DPS) [dungeon]; Dragonrider Boots (18102, -2.98 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Maiden's Circle (13001, -1.22 DPS) [world_drop] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (832.5 DPS) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Maiden's Circle (13001, -1.03 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (832.5 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (832.5 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Archlight Talisman; shoulder: Defiler's Lamellar Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Defiler's Lamellar Girdle; legs: Champion's Legplates; feet: Depleted Thorium Sabatons; finger1: Legionnaire's Band; finger2: Don Julio's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: The Unstoppable Force

No-known-source sample (15 of 2141, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

