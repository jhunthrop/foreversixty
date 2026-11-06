# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 46.7. Weights run: 5.4s. Verify run: 14.6s. 281 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.319 ± 0.020, crit=0.796 ± 0.018 per rating point (14 rating = 1%, 11.139 per %), hit=1.450 ± 0.050 per rating point (10 rating = 1%, 14.503 per %), melee_haste=5.157 ± 0.678, intellect=1.066 ± 0.026, spell_power=2.415 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 26.6 attack_power points (1.14 DPS) | yes | Totemic Leather Hood (252448, -0.21 DPS) [crafted]; Veteran's Silvered Chain Helm (250528, -0.28 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.41 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 7.9 attack_power points (0.34 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.60 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.7 attack_power points (0.93 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.51 DPS) [crafted]; Serpent's Shoulders (5404, -0.64 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.07 DPS) [crafted]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop]; Dark Leather Cloak (2316, -0.07 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 31.7 attack_power points (1.36 DPS) | yes | Crusader's Chain Shirt (250492, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.30 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.50 DPS) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.00 DPS) [quest]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop] |
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

No-known-source sample (15 of 281, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 30 (human, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 82.1. Weights run: 5.9s. Verify run: 14.9s. 504 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.561 ± 0.026, crit=1.183 ± 0.027 per rating point (14 rating = 1%, 16.559 per %), hit=1.966 ± 0.076 per rating point (10 rating = 1%, 19.662 per %), melee_haste=11.026 ± 0.885, intellect=0.912 ± 0.034, spell_power=2.522 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 57.3 attack_power points (2.65 DPS) | yes | Totemic Leather Helm (252456, -1.25 DPS) [crafted]; Holy Shroud (2721, -1.37 DPS) [world_drop]; Barbaric Iron Helm (7915, -1.70 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.9 attack_power points (0.83 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.02 DPS) [world]; Sentinel's Medallion (19541, -0.25 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 30.9 attack_power points (1.43 DPS) | yes | Death Speaker Mantle (6685, -0.27 DPS) [dungeon]; Invoker's Mantle (215365, -0.40 DPS) [crafted]; Forest Tracker Epaulets (2278, -0.86 DPS, sim-verified) [world_drop] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.24 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.24 DPS) [quest]; Hawkeye's Cloak (14593, -0.29 DPS, sim-verified) [world_drop] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 42.3 attack_power points (1.95 DPS) | yes | Totemic Leather Tunic (252451, -0.19 DPS) [crafted]; Hard Gold Cuirass (250533, -0.24 DPS) [crafted]; Green Silk Armor (7065, -0.36 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.7 attack_power points (1.05 DPS) | yes | Yorgen Bracers (13012, -0.28 DPS) [world_drop]; Pugilist Bracers (4438, -0.31 DPS) [dungeon]; Hawkeye's Bracers (14590, -1.31 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 27.0 attack_power points (1.25 DPS) | yes | Mail Combat Gauntlets (4075, -0.22 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.23 DPS) [world]; Insignia Gloves (6408, -0.34 DPS, sim-verified) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 42.3 attack_power points (1.95 DPS) | yes | Skirmisher's Leather Belt (252461, -0.14 DPS) [crafted]; Prowler's Leather Belt (252459, -0.52 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.53 DPS) [crafted] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 46.8 attack_power points (2.16 DPS) | yes | Totemic Leather Leggings (252458, -0.14 DPS) [crafted]; Totemic Leather Pants (252446, -0.35 DPS) [crafted]; Ferine Leggings (6690, -0.45 DPS) [dungeon] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 31.7 attack_power points (1.46 DPS) | yes | Disjointed Shoes (277226, -0.33 DPS) [quest]; Gilded Slippers (254001, -0.35 DPS) [crafted]; Totemic Leather Boots (252442, -0.36 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.4 attack_power points (0.99 DPS) | yes | Lorekeeper's Ring (19525, -0.17 DPS) [rep]; Ring of Power Regulation (273030, -0.29 DPS) [dungeon]; Thunderbrow Ring (13097, -1.03 DPS, sim-verified) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (82.1 DPS) | yes | Lorekeeper's Ring (19525, -0.09 DPS) [rep]; Ring of Power Regulation (273030, -0.21 DPS) [dungeon]; Thunderbrow Ring (13097, -0.26 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.32 DPS) [world]; Corpsemaker (6687, -0.49 DPS) [dungeon]; Viscous Hammer (13045, -5.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Sergeant Major's Cape; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Protector's Band; finger2: Snake Hoop; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 504, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (human, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 135.3. Weights run: 6.0s. Verify run: 16.2s. 691 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=2.655 ± 0.055, crit=2.676 ± 0.065 per rating point (14 rating = 1%, 37.471 per %), hit=2.983 ± 0.129 per rating point (10 rating = 1%, 29.826 per %), melee_haste=12.380 ± 1.556, intellect=3.126 ± 0.046, spell_power=2.813 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 75.6 attack_power points (4.00 DPS) | yes | Corpseshroud (10574, -0.16 DPS) [dungeon]; Big Voodoo Mask (8201, -0.35 DPS) [crafted]; Raging Berserker's Helm (7719, -0.62 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 44.7 attack_power points (2.36 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.71 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.74 DPS) [world]; Kaleidoscope Chain (13084, -0.94 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 60.3 attack_power points (3.19 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.47 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.76 DPS) [crafted]; Guardian Cloak (5965, -0.76 DPS) [crafted]; Expeditionary's Cape (271922, -1.37 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.10 DPS) [crafted]; Robe of Power (7054, -0.20 DPS) [crafted]; Icemail Jerkin (1981, -0.64 DPS, sim-verified) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 37.3 attack_power points (1.97 DPS) | yes | Guardian Leather Bracers (4260, -0.09 DPS) [crafted]; Windchaser Cuffs (14429, -0.21 DPS) [world_drop]; Warden's Wristbands (14600, -0.23 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, -1.34 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 66.6 attack_power points (3.52 DPS) | yes | Highlander's Cloth Girdle (20098, -0.79 DPS, sim-verified) [rep]; Skirmisher's Leather Belt (252461, -0.85 DPS) [crafted]; Gilded Cord (254037, -1.01 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 76.9 attack_power points (4.06 DPS) | yes | Ferine Leggings (6690, -0.71 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.73 DPS) [crafted]; Totemic Leather Leggings (252458, -0.91 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 67.5 attack_power points (3.57 DPS) | yes | Skycaller's Leather Shoes (252532, -0.48 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.48 DPS) [crafted]; Skirmisher's Mail Boots (252564, -1.31 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 43.8 attack_power points (2.31 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Ring of the Underwood (2951, -0.69 DPS) [world_drop]; Falcon's Hook (7552, -0.73 DPS) [dungeon] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 37.6 attack_power points (1.98 DPS) | yes | Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon]; Snake Hoop (6750, -1.06 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, -0.10 DPS) [vendor]; Nightblade (1982, -2.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Sergeant Major's Cape; chest: Robe of the Magi; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Protector's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 691, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (human, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 191.3. Weights run: 5.7s. Verify run: 30.7s. 883 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=2.503 ± 0.057, crit=2.993 ± 0.072 per rating point (14 rating = 1%, 41.903 per %), hit=2.970 ± 0.151 per rating point (10 rating = 1%, 29.698 per %), melee_haste=20.847 ± 2.163, intellect=4.089 ± 0.099, spell_power=2.756 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 134.1 attack_power points (7.80 DPS) | yes | Soothsayer's Headdress (17740, -0.55 DPS) [dungeon]; Knight-Lieutenant's Lamellar Helm (220819, -0.75 DPS) [vendor]; Eye of Theradras (17715, -0.85 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 57.2 attack_power points (3.33 DPS) | yes | Glowing Eye of Mordresh (10769, -0.30 DPS) [dungeon]; Skibi's Pendant (13089, -0.80 DPS) [world_drop]; Mindburst Medallion (11196, -0.94 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 123.1 attack_power points (7.15 DPS) | yes | Rotgrip Mantle (17732, -0.80 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -1.43 DPS) [dungeon]; Knight-Lieutenant's Lamellar Pauldrons (220818, -1.45 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -0.33 DPS) [crafted]; Expeditionary's Cape (271921, -0.34 DPS) [rep]; Sergeant Major's Cape (16336, -0.60 DPS) [pvp] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 134.1 attack_power points (7.80 DPS) | yes | Spiderfang Carapace (11633, -0.83 DPS) [dungeon]; Knight's Lamellar Chestplate (220815, -1.12 DPS) [vendor]; Robes of Insight (940, -1.85 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 68.4 attack_power points (3.98 DPS) | yes | Skirmisher's Mail Bracers (252572, -0.36 DPS) [crafted]; Aristocratic Cuffs (12546, -0.41 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.62 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller's Leather Gauntlets (252550, +0.00 DPS) [crafted]; Skirmisher's Mail Gauntlets (252586, +0.00 DPS) [crafted]; Stonerender Gauntlets (17007, -2.33 DPS, sim-verified) [world_drop] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Dawnspire Cord (12466, -0.27 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, -0.33 DPS) [crafted] |
| legs | Knight's Lamellar Legplates (220816) | Captain Dirgehammer [vendor] | 114.9 attack_power points (6.68 DPS) | yes | Spellshock Leggings (9484, -0.62 DPS) [dungeon]; Imperial Red Pants (8251, -0.93 DPS) [world_drop]; Knight's Plate Leggings (220797, -1.25 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 100.9 attack_power points (5.86 DPS) | yes | Skycaller's Leather Boots (252471, -0.85 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.85 DPS) [crafted]; Sergeant Major's Lamellar Boots (220814, -1.35 DPS) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Cyclopean Band (11824, +0.00 DPS) [dungeon]; Mindseye Circle (10634, -0.04 DPS) [dungeon]; Brainlash (6440, -0.45 DPS, sim-verified) [dungeon] |
| trinket1 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.71 DPS, sim-verified) [world_drop] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Thorium Greatmace (250613, -0.85 DPS, sim-verified) [crafted]; Taran Icebreaker (2915, -1.85 DPS) [world_drop]; The Rockpounder (9413, -1.92 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-verified (191.3 DPS) | yes | - |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Highlander's Lamellar Girdle; legs: Knight's Lamellar Legplates; feet: Greaves of Withering Despair; finger2: Blackstone Ring; trinket1: Sanctified Orb; trinket2: Thunderbrew's Boot Flask; main_hand: Warmonger

No-known-source sample (15 of 883, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (human, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 337.7. Weights run: 5.7s. Verify run: 88.1s. 2115 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=3.962 ± 0.092, crit=5.698 ± 0.132 per rating point (14 rating = 1%, 79.777 per %), hit=4.354 ± 0.230 per rating point (10 rating = 1%, 43.542 per %), melee_haste=21.813 ± 2.969, intellect=4.933 ± 0.144, spell_power=2.888 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 350.4 attack_power points (18.05 DPS) | yes | Field Marshal's Lamellar Faceguard (16474, -2.76 DPS) [vendor]; Lionheart Helm (12640, -3.31 DPS) [crafted]; Field Marshal's Lamellar Headguard (231648, -4.51 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Theramore Pauldrons (271955, +0.00 DPS) [rep]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -1.62 DPS) [dungeon]; Sergeant Major's Cape (16337, -2.08 DPS) [pvp]; Fel Cape (279269, -2.55 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, +0.00 DPS) [dungeon]; Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Dawn Armor (252483, +0.00 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-verified (337.7 DPS) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Spitfire Bracers (20481, +0.00 DPS) [crafted]; Forest Stalker's Bracers (19587, -0.18 DPS) [rep] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Mourning Handguards (279257, -0.21 DPS) [crafted] |
| waist | Highlander's Lamellar Girdle (20042) | The League of Arathor [rep] | sim-verified (337.7 DPS) | yes | Highlander's Mail Girdle (20044, +0.00 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, +0.00 DPS) [rep]; Highlander's Cloth Girdle (20047, -0.63 DPS) [rep] |
| legs | Sentinel's Lamellar Legguards (237814) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, -1.80 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.80 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.33 DPS) [dungeon]; Cutthroat's Signet (272408, -1.97 DPS) [vendor]; Painweaver Band (13098, -2.14 DPS) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.05 DPS) [dungeon]; Cutthroat's Signet (272408, -0.70 DPS) [vendor]; Painweaver Band (13098, -0.87 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -0.76 DPS) [dungeon]; Uther's Strength (11302, -0.89 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.19 DPS) [vendor]; Frozen Heart of the Mountain (249469, -1.25 DPS) [crafted]; Hand of Justice (11815, -2.24 DPS) [dungeon] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Highlander's Lamellar Girdle; legs: Sentinel's Lamellar Legguards; finger1: Don Julio's Band; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2115, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60, raid preset (human, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 853.7. Weights run: 6.1s. Verify run: 88.3s. 2115 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.006, agility=2.342 ± 0.049, crit=3.130 ± 0.066 per rating point (14 rating = 1%, 43.820 per %), hit=6.156 ± 0.387 per rating point (10 rating = 1%, 61.557 per %), melee_haste=30.233 ± 4.698, intellect=1.908 ± 0.128, spell_power=1.823 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 254.3 attack_power points (39.45 DPS) | yes | Stalwart Helm (250599, -5.10 DPS) [crafted]; Field Marshal's Lamellar Faceguard (16474, -11.78 DPS) [vendor]; Mask of the Unforgiven (13404, -13.56 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Lightforge Spaulders (16729, +0.00 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -2.55 DPS) [vendor]; Shroud of Arcane Mastery (22330, -3.64 DPS) [dungeon]; Cloak of Flames (3475, -6.89 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Field Marshal's Lamellar Breastplate (231649, -0.02 DPS) [vendor]; Savage Gladiator Chain (11726, -1.04 DPS) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor] |
| waist | Highlander's Lamellar Girdle (20042) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Clutch of Andros (13956, -0.22 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -2.58 DPS) [vendor]; Sentinel's Lamellar Legguards (22753, -4.56 DPS) [rep] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Marshal's Lamellar Boots (16472, -0.46 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -4.82 DPS) [vendor]; Blackstone Ring (17713, -5.87 DPS) [dungeon]; Blessed Band of Light (272407, -6.14 DPS) [vendor] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -0.94 DPS) [vendor]; Blackstone Ring (17713, -1.98 DPS) [dungeon]; Blessed Band of Light (272407, -2.26 DPS) [vendor] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.41 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Serenity Field (272439, -2.26 DPS) [vendor]; Burst of Knowledge (11832, -2.83 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Lionheart Helm; neck: Skibi's Pendant; shoulder: Theramore Pauldrons; back: Arcanoweave Cloak; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Highlander's Lamellar Girdle; legs: Sentinel's Chain Leggings; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Weakness Analyzer; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 2115, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (undead, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 43.6. Weights run: 5.4s. Verify run: 14.1s. 268 eligible items had no known source.

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

No-known-source sample (15 of 268, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (undead, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 83.0. Weights run: 5.9s. Verify run: 16.2s. 491 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.548 ± 0.026, crit=1.198 ± 0.026 per rating point (14 rating = 1%, 16.778 per %), hit=1.911 ± 0.072 per rating point (10 rating = 1%, 19.107 per %), melee_haste=8.565 ± 0.931, intellect=1.530 ± 0.031, spell_power=2.697 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 59.8 attack_power points (2.82 DPS) | yes | Totemic Leather Helm (252456, -1.12 DPS, sim-verified) [crafted]; Barbaric Iron Helm (7915, -1.31 DPS) [crafted]; Enchanter's Cowl (4322, -1.34 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Archmagister's Faceted Pendant (273049, -0.15 DPS, sim-verified) [world]; Scout's Medallion (19537, -0.37 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 38.0 attack_power points (1.80 DPS) | yes | Death Speaker Mantle (6685, -0.24 DPS) [dungeon]; Magician's Mantle (12998, -0.51 DPS) [world_drop]; Forest Tracker Epaulets (2278, -0.52 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.16 DPS) [crafted]; Windsong Drape (15468, -0.16 DPS) [quest]; Tigerstrike Mantle (13108, -0.21 DPS) [world_drop] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 44.4 attack_power points (2.10 DPS) | yes | Green Silk Armor (7065, -0.01 DPS) [crafted]; Beguiler Robes (7728, -0.08 DPS) [dungeon]; Totemic Leather Tunic (252451, -0.09 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 24.3 attack_power points (1.15 DPS) | yes | Hawkeye's Bracers (14590, -0.33 DPS) [world_drop]; Yorgen Bracers (13012, -0.36 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.85 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 28.2 attack_power points (1.33 DPS) | yes | Oilrag Handwraps (16741, -0.04 DPS) [quest]; Jutebraid Gloves (10654, -0.20 DPS) [quest]; Insignia Gloves (6408, -0.23 DPS) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 44.4 attack_power points (2.10 DPS) | yes | Skirmisher's Leather Belt (252461, -0.10 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -0.39 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.48 DPS) [rep] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 49.1 attack_power points (2.32 DPS) | yes | Totemic Leather Leggings (252458, -0.01 DPS) [crafted]; Ferine Leggings (6690, -0.22 DPS) [dungeon]; Totemic Leather Pants (252446, -0.29 DPS) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 32.9 attack_power points (1.55 DPS) | yes | Gilded Slippers (254001, -0.16 DPS) [crafted]; Totemic Leather Boots (252442, -0.32 DPS) [crafted]; Acidic Walkers (9454, -0.34 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 23.8 attack_power points (1.13 DPS) | yes | Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Advisor's Ring (19521, -0.23 DPS) [rep]; Ring of Power Regulation (273030, -0.36 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.3 attack_power points (1.01 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Advisor's Ring (19521, -0.11 DPS) [rep]; Ring of Power Regulation (273030, -0.24 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.35 DPS) [world]; Corpsemaker (6687, -0.50 DPS) [dungeon]; Viscous Hammer (13045, -5.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Hawkeye's Cloak; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 491, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (undead, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 133.8. Weights run: 6.0s. Verify run: 16.1s. 675 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=2.655 ± 0.055, crit=2.676 ± 0.065 per rating point (14 rating = 1%, 37.471 per %), hit=2.983 ± 0.129 per rating point (10 rating = 1%, 29.826 per %), melee_haste=12.380 ± 1.556, intellect=3.126 ± 0.046, spell_power=2.813 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 75.6 attack_power points (4.00 DPS) | yes | Corpseshroud (10574, -0.16 DPS) [dungeon]; Big Voodoo Mask (8201, -0.35 DPS) [crafted]; Raging Berserker's Helm (7719, -0.62 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 44.7 attack_power points (2.36 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.71 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.74 DPS) [world]; Kaleidoscope Chain (13084, -0.90 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 60.3 attack_power points (3.19 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.47 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.10 DPS) [crafted]; Guardian Cloak (5965, -0.10 DPS) [crafted]; Blackforge Cape (6424, -0.50 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 93.5 attack_power points (4.94 DPS) | yes | Robe of the Magi (1716, -0.68 DPS) [world_drop]; Dreamweave Vest (10021, -0.78 DPS) [crafted]; Robe of Power (7054, -0.88 DPS) [crafted] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 37.3 attack_power points (1.97 DPS) | yes | Radiant Silver Bracers (4545, -0.06 DPS) [quest]; Guardian Leather Bracers (4260, -0.09 DPS) [crafted]; Windchaser Cuffs (14429, -0.21 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, -2.12 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 66.6 attack_power points (3.52 DPS) | yes | Skirmisher's Leather Belt (252461, -0.85 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.88 DPS, sim-verified) [rep]; Gilded Cord (254037, -1.01 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 76.9 attack_power points (4.06 DPS) | yes | Ferine Leggings (6690, -0.71 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.73 DPS) [crafted]; Totemic Leather Leggings (252458, -0.91 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 67.5 attack_power points (3.57 DPS) | yes | Skirmisher's Mail Boots (252564, -0.22 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.48 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.48 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 43.8 attack_power points (2.31 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Ring of the Underwood (2951, -0.69 DPS) [world_drop]; Falcon's Hook (7552, -0.73 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 37.6 attack_power points (1.98 DPS) | yes | Snake Hoop (6750, -0.27 DPS) [quest]; Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, -0.10 DPS) [vendor]; Nightblade (1982, -1.99 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Legionnaire's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Libram of Invocation

No-known-source sample (15 of 675, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (undead, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 188.5. Weights run: 5.7s. Verify run: 26.0s. 883 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=3.035 ± 0.063, crit=3.584 ± 0.081 per rating point (14 rating = 1%, 50.171 per %), hit=3.100 ± 0.155 per rating point (10 rating = 1%, 30.995 per %), melee_haste=14.698 ± 1.742, intellect=2.514 ± 0.089, spell_power=2.743 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Eye of Theradras (17715, -0.14 DPS) [dungeon]; Red Mageweave Headband (10033, -0.49 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 50.5 attack_power points (2.54 DPS) | yes | Woven Ivy Necklace (19159, -0.50 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.56 DPS) [dungeon]; Scout's Medallion (19535, -0.71 DPS) [rep] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 91.4 attack_power points (4.61 DPS) | yes | Rotgrip Mantle (17732, -0.53 DPS) [dungeon]; Earthslag Shoulders (11632, -0.63 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -0.76 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Dark Phantom Cape (13122, -0.51 DPS) [world_drop]; Blackveil Cape (11626, -1.20 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Spiderfang Carapace (11633, -0.16 DPS) [dungeon]; Stone Guard's Plate Armor (220801, -0.97 DPS) [vendor]; Icemail Jerkin (1981, -1.97 DPS, sim-verified) [world_drop] |
| wrist | Skirmisher's Mail Bracers (252572) | Leatherworking [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Bracers (252539, -0.04 DPS) [crafted]; Runic Leather Bracers (15092, -0.09 DPS) [crafted]; Deepfury Bracers (13120, -0.99 DPS, sim-verified) [world_drop] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skirmisher's Mail Gauntlets (252586, -0.52 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -0.66 DPS) [crafted]; Justicar's Gloves (250555, -0.72 DPS) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, -0.15 DPS) [dungeon]; Defiler's Cloth Girdle (20165, -0.36 DPS) [rep]; Sagebrush Girdle (17778, -0.95 DPS) [quest] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 106.9 attack_power points (5.39 DPS) | yes | Stormshroud Pants (15057, -0.33 DPS) [crafted]; Spellshock Leggings (9484, -0.94 DPS) [dungeon]; Imperial Red Pants (8251, -1.19 DPS) [world_drop] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 86.3 attack_power points (4.35 DPS) | yes | Fleetfoot Greaves (11627, -0.68 DPS) [dungeon]; Battlechaser's Greaves (12555, -0.81 DPS) [dungeon]; Skycaller's Leather Boots (252471, -0.88 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 51.1 attack_power points (2.58 DPS) | yes | Legionnaire's Band (19511, -0.09 DPS) [rep]; Masons Fraternity Ring (9533, -0.43 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 51.0 attack_power points (2.57 DPS) | yes | Legionnaire's Band (19511, -0.08 DPS) [rep]; Masons Fraternity Ring (9533, -0.43 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.97 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, -0.58 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -1.28 DPS, sim-verified) [world_drop]; The Rockpounder (9413, -3.19 DPS) [dungeon]; Stoneraven (13059, -3.85 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; neck: Skibi's Pendant; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Skirmisher's Mail Bracers; hands: Stonerender Gauntlets; waist: Defiler's Lamellar Girdle; legs: Stone Guard's Plate Leggings; feet: Greaves of Withering Despair; finger1: Cyclopean Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 883, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (undead, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 339.0. Weights run: 5.7s. Verify run: 81.7s. 2141 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.755 ± 0.042, crit=2.663 ± 0.060 per rating point (14 rating = 1%, 37.288 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=28.014 ± 2.574, intellect=4.883 ± 0.126, spell_power=3.143 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 259.7 attack_power points (18.17 DPS) | yes | Magister's Crown (16686, -5.50 DPS) [dungeon]; Living Crown (252561, -6.52 DPS) [crafted]; Blue Dragonscale Helm (252604, -6.52 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Tooth of Gnarr (13141, -1.27 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Darkspear Shoulders (272104, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -0.49 DPS) [world]; Deep Woodlands Cloak (19121, -0.54 DPS) [quest]; Royal Tribunal Cloak (13376, -0.79 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blue Dragonscale Breastplate (15048, +0.00 DPS) [crafted]; Vest of Elements (16666, +0.00 DPS) [dungeon]; Magister's Robes (16688, +0.00 DPS) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 108.2 attack_power points (7.57 DPS) | yes | Spitfire Bracers (20481, -1.63 DPS) [crafted]; Tranquil Wristguards (279256, -1.69 DPS) [crafted]; Runecloth Cuffs (254123, -1.74 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handguards (272101, -0.29 DPS) [vendor]; Spitfire Gauntlets (20480, -2.99 DPS) [crafted] |
| waist | Defiler's Lizardhide Girdle (20171) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystallized Girdle (12606, +0.00 DPS) [dungeon]; Magister's Belt (16685, +0.00 DPS) [dungeon]; Defiler's Mail Girdle (20195, +0.00 DPS) [rep] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Dragonscale Leggings (252603, -1.78 DPS) [crafted]; Outrider's Lizardhide Pants (22741, -2.56 DPS) [rep]; Champion's Legplates (250598, -3.38 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Skycaller's Leather Boots (252471, +0.00 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Maiden's Circle (13001, +0.00 DPS) [world_drop]; Deep Rooted Ring (19109, +0.00 DPS) [world] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Maiden's Circle (13001, +0.00 DPS) [world_drop]; Deep Rooted Ring (19109, +0.00 DPS) [world] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -1.24 DPS) [dungeon]; Uther's Strength (11302, -1.32 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.66 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.76 DPS) [vendor]; Hand of Justice (11815, -3.44 DPS) [dungeon]; Uther's Strength (11302, -3.52 DPS) [world_drop] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Defiler's Lizardhide Girdle; legs: Ironfeather Leggings; finger1: Don Julio's Band; finger2: Channeler's Ring; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 2141, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (undead, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 847.2. Weights run: 6.1s. Verify run: 90.6s. 2141 eligible items had no known source.

8 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.006, agility=2.342 ± 0.049, crit=3.130 ± 0.066 per rating point (14 rating = 1%, 43.820 per %), hit=6.156 ± 0.387 per rating point (10 rating = 1%, 61.557 per %), melee_haste=30.233 ± 4.698, intellect=1.908 ± 0.128, spell_power=1.823 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 254.3 attack_power points (39.45 DPS) | yes | Stalwart Helm (250599, -5.10 DPS) [crafted]; Mask of the Unforgiven (13404, -13.56 DPS) [dungeon]; Black Dragonscale Helm (252605, -13.77 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Defiler's Lamellar Spaulders (20184) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Lightforge Spaulders (16729, +0.00 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, +0.00 DPS) [dungeon]; Arcanoweave Cloak (272411, +0.00 DPS) [vendor]; Howler's Furs (272414, +0.00 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -1.04 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -7.25 DPS) [quest] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Storm Gauntlets (12632, +0.00 DPS) [crafted]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; Raider Gloves (272099, -0.19 DPS) [vendor] |
| waist | Defiler's Lamellar Girdle (20177) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Clutch of Andros (13956, -0.22 DPS) [dungeon] |
| legs | Champion's Legplates (250598) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Plate Legguards (22651, +0.00 DPS) [rep]; Outrider's Chain Leggings (22673, +0.00 DPS) [rep]; Pristine Scorpid Leggings (252606, -1.41 DPS) [crafted] |
| feet | Depleted Thorium Sabatons (279263) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Greaves of Withering Despair (22240, +0.00 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Channeler's Ring (272406, +0.00 DPS) [vendor]; Cutthroat's Signet (272408, +0.00 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -3.88 DPS) [vendor]; Channeler's Ring (272406, -4.82 DPS) [vendor]; Blackstone Ring (17713, -5.87 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.63 DPS) [vendor]; Serenity Field (272439, -4.90 DPS) [vendor]; Burst of Knowledge (11832, -5.46 DPS) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.37 DPS) [vendor]; Serenity Field (272439, -4.63 DPS) [vendor]; Burst of Knowledge (11832, -5.20 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Archlight Talisman; shoulder: Defiler's Lamellar Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Defiler's Lamellar Girdle; legs: Champion's Legplates; feet: Depleted Thorium Sabatons; finger1: Legionnaire's Band; finger2: Don Julio's Band; main_hand: The Unstoppable Force

No-known-source sample (15 of 2141, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh

