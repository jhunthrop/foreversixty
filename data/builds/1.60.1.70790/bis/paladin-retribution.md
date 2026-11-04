# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 42.1. Weights run: 7.2s. Verify run: 18.2s. 265 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.221 ± 0.019, crit=0.724 ± 0.017 per rating point (14 rating = 1%, 10.135 per %), hit=1.112 ± 0.044 per rating point (10 rating = 1%, 11.120 per %), melee_haste=6.441 ± 0.732, intellect=1.157 ± 0.035, spell_power=2.269 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 25.0 attack_power points (1.09 DPS) | yes | Totemic Leather Hood (252448, -0.19 DPS, sim-verified) [crafted]; Veteran's Silvered Chain Helm (250528, -0.22 DPS) [crafted]; Defender's Leather Hood (252447, -0.39 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 7.3 attack_power points (0.32 DPS) | yes | Erudite's Amulet (277204, -0.20 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.58 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.8 attack_power points (0.95 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.28 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.55 DPS) [crafted]; Serpent's Shoulders (5404, -0.68 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.05 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 31.1 attack_power points (1.36 DPS) | yes | Crusader's Chain Shirt (250492, -0.22 DPS, sim-verified) [crafted]; Stormrider's Leather Armor (252492, -0.31 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.48 DPS) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.44 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.01 DPS) [quest]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 23.2 attack_power points (1.02 DPS) | yes | Totemic Leather Gloves (252497, -0.17 DPS) [crafted]; Crusader's Gloves (250512, -0.26 DPS) [crafted]; Thorbia's Gauntlets (12994, -0.32 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Cobrahn's Grasp (6460, -0.01 DPS) [dungeon]; Crusader's Chain Belt (250517, -0.03 DPS) [crafted]; Totemic Leather Belt (252431, -0.04 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 38.5 attack_power points (1.68 DPS) | yes | Stormrider's Leather Pants (252502, -0.39 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted]; Crusader's Chain Leggings (250497, -0.52 DPS, sim-verified) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 23.6 attack_power points (1.03 DPS) | yes | Crusader's Boots (250507, -0.09 DPS) [crafted]; Spidersilk Boots (4320, -0.14 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.18 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 13.7 attack_power points (0.60 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Demon Band (12054, -0.25 DPS) [world_drop]; Protector's Band (20439, -0.63 DPS, sim-verified) [rep] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Demon Band (12054, -0.15 DPS) [world_drop]; Protector's Band (20439, -0.24 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sword of the Fallen (286977, -2.94 DPS) [dungeon]; Duskbringer (2205, -3.12 DPS) [dungeon]; Smite's Mighty Hammer (7230, -4.20 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Totemic Leather Armor; wrist: Patterned Bronze Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Verigan's Fist

No-known-source sample (15 of 265, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers; 14148 Crystalline Cuffs

### Band 30 (human, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 72.7. Weights run: 7.5s. Verify run: 18.1s. 479 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.385 ± 0.023, crit=1.082 ± 0.023 per rating point (14 rating = 1%, 15.151 per %), hit=1.648 ± 0.065 per rating point (10 rating = 1%, 16.480 per %), melee_haste=8.503 ± 0.797, intellect=1.083 ± 0.034, spell_power=2.699 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 59.8 attack_power points (2.84 DPS) | yes | Barbaric Iron Helm (7915, -1.39 DPS) [crafted]; Holy Shroud (2721, -1.43 DPS) [world_drop]; Totemic Leather Helm (252456, -1.43 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 18.5 attack_power points (0.88 DPS) | yes | Kaleidoscope Chain (13084, -0.34 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.35 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 34.0 attack_power points (1.62 DPS) | yes | Forest Tracker Epaulets (2278, -0.42 DPS) [world_drop]; Invoker's Mantle (215365, -0.46 DPS) [crafted]; Death Speaker Mantle (6685, -0.75 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (72.7 DPS) | yes | Hawkeye's Cloak (14593, -0.10 DPS) [world_drop]; Hillman's Cloak (3719, -0.21 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.21 DPS) [quest] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 44.4 attack_power points (2.11 DPS) | yes | Green Silk Armor (7065, -0.29 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Totemic Leather Tunic (252451, -0.49 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 24.3 attack_power points (1.16 DPS) | yes | Yorgen Bracers (13012, -0.39 DPS) [world_drop]; Pugilist Bracers (4438, -0.39 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.95 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 27.2 attack_power points (1.29 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.25 DPS) [world]; Mail Combat Gauntlets (4075, -0.27 DPS) [world_drop]; Insignia Gloves (6408, -0.34 DPS, sim-verified) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 44.4 attack_power points (2.11 DPS) | yes | Skycaller's Leather Belt (252522, -0.52 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.54 DPS) [rep]; Skirmisher's Leather Belt (252461, -0.58 DPS, sim-verified) [crafted] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 49.1 attack_power points (2.33 DPS) | yes | Totemic Leather Leggings (252458, -0.37 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.37 DPS) [crafted]; Ferine Leggings (6690, -0.48 DPS) [dungeon] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 32.9 attack_power points (1.56 DPS) | yes | Totemic Leather Boots (252442, -0.32 DPS) [crafted]; Disjointed Shoes (277226, -0.35 DPS) [quest]; Gilded Slippers (254001, -0.60 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 21.6 attack_power points (1.03 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Lorekeeper's Ring (19525, -0.13 DPS) [rep]; Ring of Power Regulation (273030, -0.26 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 20.3 attack_power points (0.97 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Lorekeeper's Ring (19525, -0.07 DPS) [rep]; Ring of Power Regulation (273030, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (72.7 DPS) | yes | Corpsemaker (6687, -0.46 DPS) [dungeon]; Felbough Arm (273037, -0.52 DPS) [world]; Viscous Hammer (13045, -4.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Sergeant Major's Cape; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 479, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 123.4. Weights run: 7.6s. Verify run: 18.6s. 657 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=2.542 ± 0.053, crit=2.572 ± 0.062 per rating point (14 rating = 1%, 36.013 per %), hit=2.658 ± 0.119 per rating point (10 rating = 1%, 26.579 per %), melee_haste=12.529 ± 1.361, intellect=3.010 ± 0.044, spell_power=2.942 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 74.7 attack_power points (3.97 DPS) | yes | Big Voodoo Mask (8201, -0.32 DPS) [crafted]; Crusader's Chain Helm (250502, -0.59 DPS) [crafted]; Raging Berserker's Helm (7719, -0.65 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 44.7 attack_power points (2.37 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.78 DPS) [world]; Sentinel's Medallion (19540, -0.89 DPS) [rep]; Kaleidoscope Chain (13084, -0.98 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 59.7 attack_power points (3.17 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Bloodmage Mantle (7684, -0.33 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.37 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.66 DPS) [rep]; Long Silken Cloak (4326, -0.68 DPS) [crafted]; Guardian Cloak (5965, -0.68 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (123.4 DPS) | yes | Dreamweave Vest (10021, -0.15 DPS) [crafted]; Robe of Power (7054, -0.29 DPS) [crafted]; Icemail Jerkin (1981, -1.40 DPS, sim-verified) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 35.9 attack_power points (1.91 DPS) | yes | Guardian Leather Bracers (4260, -0.01 DPS) [crafted]; Turtle Scale Bracers (8198, -0.17 DPS) [crafted]; Windchaser Cuffs (14429, -0.20 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, +0.00 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 65.7 attack_power points (3.49 DPS) | yes | Highlander's Cloth Girdle (20098, -0.67 DPS) [rep]; Skirmisher's Leather Belt (252461, -0.78 DPS) [crafted]; Gilded Cord (254037, -0.96 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 77.3 attack_power points (4.11 DPS) | yes | Crimson Silk Pantaloons (7062, -0.74 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.81 DPS) [dungeon]; Totemic Leather Leggings (252458, -0.92 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 70.6 attack_power points (3.75 DPS) | yes | Skirmisher's Mail Boots (252564, -0.35 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.60 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.60 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 44.5 attack_power points (2.36 DPS) | yes | Snake Hoop (6750, -0.63 DPS) [quest]; Ring of the Underwood (2951, -0.80 DPS) [world_drop]; Falcon's Hook (7552, -0.82 DPS) [dungeon] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 36.7 attack_power points (1.95 DPS) | yes | Snake Hoop (6750, -0.21 DPS) [quest]; Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.41 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, -0.16 DPS) [vendor]; Nightblade (1982, -1.25 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Sergeant Major's Cape; chest: Robe of the Magi; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Protector's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 657, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (human, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 178.2. Weights run: 7.8s. Verify run: 22.6s. 840 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=3.689 ± 0.081, crit=4.458 ± 0.107 per rating point (14 rating = 1%, 62.405 per %), hit=3.296 ± 0.198 per rating point (10 rating = 1%, 32.958 per %), melee_haste=16.070 ± 2.453, intellect=4.789 ± 0.062, spell_power=3.510 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 162.5 attack_power points (8.46 DPS) | yes | Eye of Theradras (17715, -0.47 DPS) [dungeon]; Soothsayer's Headdress (17740, -0.52 DPS) [dungeon]; Knight-Lieutenant's Lamellar Helm (220819, -1.12 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 67.1 attack_power points (3.49 DPS) | yes | Glowing Eye of Mordresh (10769, -0.22 DPS) [dungeon]; Skibi's Pendant (13089, -0.42 DPS) [world_drop]; Mindburst Medallion (11196, -0.90 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 148.4 attack_power points (7.73 DPS) | yes | Rotgrip Mantle (17732, -1.15 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Lamellar Pauldrons (220818, -1.78 DPS) [vendor]; Rockshard Pauldrons (9411, -1.95 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (178.2 DPS) | yes | Runecloth Cloak (13860, -0.42 DPS) [crafted]; Expeditionary's Cape (271921, -0.56 DPS) [rep]; Sergeant Major's Cape (16336, -0.72 DPS) [pvp] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 162.5 attack_power points (8.46 DPS) | yes | Spiderfang Carapace (11633, -1.24 DPS) [dungeon]; Knight's Lamellar Chestplate (220815, -1.49 DPS) [vendor]; Icemail Jerkin (1981, -2.13 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 83.0 attack_power points (4.32 DPS) | yes | Skirmisher's Mail Bracers (252572, -0.58 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.70 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.75 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (178.2 DPS) | yes | Sergeant Major's Lamellar Gauntlets (220817, +0.00 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, +0.00 DPS) [crafted]; Stonerender Gauntlets (17007, -1.97 DPS, sim-verified) [world_drop] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 124.9 attack_power points (6.51 DPS) | yes | Highlander's Cloth Girdle (20097, -0.36 DPS) [rep]; Skycaller's Leather Waistguard (252476, -0.59 DPS) [crafted]; Skycaller's Mail Belt (252589, -0.59 DPS) [crafted] |
| legs | Knight's Lamellar Legplates (220816) | Captain Dirgehammer [vendor] | 133.9 attack_power points (6.97 DPS) | yes | Imperial Red Pants (8251, -0.35 DPS) [world_drop]; Knight's Plate Leggings (220797, -0.43 DPS) [vendor]; Stormshroud Pants (15057, -0.47 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 119.5 attack_power points (6.22 DPS) | yes | Skycaller's Leather Boots (252471, -0.74 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.74 DPS) [crafted]; Fleetfoot Greaves (11627, -1.07 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 59.0 attack_power points (3.07 DPS) | yes | Blackstone Ring (17713, -0.32 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.39 DPS) [quest]; Band of the Unicorn (7553, -0.70 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 55.2 attack_power points (2.87 DPS) | yes | Blackstone Ring (17713, -0.12 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.19 DPS) [quest]; Band of the Unicorn (7553, -0.50 DPS) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (178.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (178.2 DPS) | yes | Uther's Strength (11302, -0.45 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (178.2 DPS) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Warmonger (13052, -2.84 DPS) [world_drop]; The Rockpounder (9413, -2.93 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Highlander's Lamellar Girdle; legs: Knight's Lamellar Legplates; feet: Greaves of Withering Despair; finger2: Protector's Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 840, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (human, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 282.8. Weights run: 7.9s. Verify run: 96.2s. 1646 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=3.666 ± 0.086, crit=5.280 ± 0.122 per rating point (14 rating = 1%, 73.915 per %), hit=3.526 ± 0.217 per rating point (10 rating = 1%, 35.263 per %), melee_haste=19.230 ± 2.507, intellect=4.187 ± 0.149, spell_power=2.836 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 310.0 attack_power points (16.08 DPS) | yes | Field Marshal's Lamellar Faceguard (16474, -1.83 DPS) [vendor]; Field Marshal's Lamellar Headguard (231648, -3.39 DPS) [vendor]; Lionheart Helm (12640, -7.84 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 75.9 attack_power points (3.94 DPS) | yes | Horizon Choker (13085, -0.90 DPS) [world_drop]; Sentinel's Medallion (19538, -1.08 DPS) [rep]; Skibi's Pendant (13089, -2.99 DPS, sim-verified) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (282.8 DPS) | yes | Theramore Pauldrons (271955, +0.00 DPS) [rep]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Mantle of the Timbermaw (19050, -3.80 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16337) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (282.8 DPS) | yes | Fel Cape (279269, -0.34 DPS) [crafted]; Expeditionary's Cape (271907, -0.70 DPS) [rep]; Cape of the Black Baron (13340, -2.76 DPS, sim-verified) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (282.8 DPS) | yes | Savage Gladiator Chain (11726, +0.00 DPS) [dungeon]; Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Dawn Armor (252483, -5.01 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 95.9 attack_power points (4.97 DPS) | yes | Spitfire Bracers (20481, -0.13 DPS) [crafted]; Vambraces of the Sadist (13400, -0.34 DPS) [dungeon]; Forest Stalker's Bracers (19587, -4.16 DPS, sim-verified) [rep] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-verified (282.8 DPS) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Mourning Handguards (279257, -2.99 DPS, sim-verified) [crafted] |
| waist | Highlander's Lamellar Girdle (20042) | The League of Arathor [rep] | 148.8 attack_power points (7.72 DPS) | yes | Highlander's Mail Girdle (20044, -0.19 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, -0.19 DPS) [rep]; Highlander's Cloth Girdle (20047, -0.52 DPS) [rep] |
| legs | Sentinel's Lamellar Legguards (237814) | Illiyana Moonblaze [vendor] | sim-verified (282.8 DPS) | yes | Sentinel's Silk Leggings (237815, -0.81 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -0.81 DPS) [vendor]; Sentinel's Chain Leggings (237819, -4.26 DPS, sim-verified) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 108.3 attack_power points (5.62 DPS) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor]; Highlander's Lamellar Greaves (20049, -1.34 DPS, sim-verified) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 123.2 attack_power points (6.39 DPS) | yes | Deep Rooted Ring (19109, -2.45 DPS) [world]; Band of the Penitent (13217, -2.56 DPS) [quest]; Don Rodrigo's Band (21563, -2.56 DPS) [rep] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 80.4 attack_power points (4.17 DPS) | yes | Deep Rooted Ring (19109, -0.23 DPS) [world]; Band of the Penitent (13217, -0.33 DPS) [quest]; Don Rodrigo's Band (21563, -0.33 DPS) [rep] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (282.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -2.70 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (282.8 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, -0.76 DPS) [world_drop] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (282.8 DPS) | yes | Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Arcanite Champion (12790, -5.63 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-verified (282.8 DPS) | yes | Libram of Fervor (23203, +0.00 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; back: Sergeant Major's Cape; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Theramore Handwraps; waist: Highlander's Lamellar Girdle; legs: Sentinel's Lamellar Legguards; finger1: Don Julio's Band; finger2: Maiden's Circle; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1646, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (human, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 779.9. Weights run: 8.0s. Verify run: 99.7s. 1646 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=1.688 ± 0.032, crit=2.253 ± 0.043 per rating point (14 rating = 1%, 31.538 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=24.905 ± 4.031, intellect=1.725 ± 0.062, spell_power=1.760 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | sim-verified (779.9 DPS) | yes | Field Marshal's Lamellar Faceguard (16474, +0.00 DPS) [vendor]; Field Marshal's Lamellar Headguard (231648, +0.00 DPS) [vendor]; Black Dragonscale Helm (252605, -15.50 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 38.4 attack_power points (7.25 DPS) | yes | Skibi's Pendant (13089, -0.82 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -2.32 DPS) [dungeon]; Sentinel's Medallion (19538, -2.47 DPS) [rep] |
| shoulder | Highlander's Lamellar Spaulders (20058) | The League of Arathor [rep] | 89.5 attack_power points (16.92 DPS) | yes | Field Marshal's Lamellar Shoulders (231651, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -1.06 DPS) [crafted]; Lieutenant Commander's Lamellar Shoulders (23277, -1.25 DPS) [vendor] |
| back | Cloak of Flames (3475) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, +0.00 DPS) [dungeon]; Sergeant Major's Cape (16337, +0.00 DPS) [pvp]; Cloak of the Honor Guard (20073, +0.00 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 141.8 attack_power points (26.80 DPS) | yes | Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Knight-Captain's Lamellar Breastplate (227142, -5.20 DPS) [vendor]; Savage Gladiator Chain (11726, -5.88 DPS, sim-verified) [dungeon] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | 59.5 attack_power points (11.24 DPS) | yes | Forest Stalker's Bracers (19587, -0.15 DPS) [rep]; Tranquil Wristguards (279256, -0.66 DPS) [crafted]; Dryad's Wrist Bindings (19595, -1.32 DPS) [rep] |
| hands | Theramore Gauntlets (271940) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Marshal's Lamellar Gloves (231643, -1.04 DPS) [pvp] |
| waist | Highlander's Lamellar Girdle (20042) | The League of Arathor [rep] | 85.1 attack_power points (16.08 DPS) | yes | Might of the Timbermaw (19044, -2.01 DPS) [crafted]; Highlander's Plate Girdle (20041, -2.35 DPS) [rep]; Highlander's Cloth Girdle (20047, -3.51 DPS) [rep] |
| legs | Champion's Legplates (250598) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legplates (16475, +0.00 DPS) [vendor]; Marshal's Lamellar Leggings (231647, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, -1.01 DPS) [vendor] |
| feet | Depleted Thorium Sabatons (279263) | Blacksmithing [crafted] | 70.8 attack_power points (13.39 DPS) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Highlander's Lamellar Greaves (20049, +0.00 DPS) [rep]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | 47.6 attack_power points (9.00 DPS) | yes | Maiden's Circle (13001, -0.73 DPS) [world_drop]; Deep Rooted Ring (19109, -1.75 DPS) [world]; Band of the Penitent (13217, -3.04 DPS) [quest] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 45.5 attack_power points (8.61 DPS) | yes | Deep Rooted Ring (19109, -1.35 DPS) [world]; Band of the Penitent (13217, -2.65 DPS) [quest]; Maiden's Circle (13001, -9.83 DPS, sim-verified) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -5.48 DPS, sim-verified) [world_drop] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcanite Champion (12790, +0.00 DPS) [crafted]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Archlight Talisman; shoulder: Highlander's Lamellar Spaulders; back: Cloak of Flames; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Theramore Gauntlets; waist: Highlander's Lamellar Girdle; legs: Champion's Legplates; feet: Depleted Thorium Sabatons; finger1: Protector's Band; finger2: Don Julio's Band; main_hand: The Unstoppable Force

No-known-source sample (15 of 1646, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (undead, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 38.8. Weights run: 7.2s. Verify run: 18.5s. 252 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.221 ± 0.019, crit=0.724 ± 0.017 per rating point (14 rating = 1%, 10.135 per %), hit=1.112 ± 0.044 per rating point (10 rating = 1%, 11.120 per %), melee_haste=6.441 ± 0.732, intellect=1.157 ± 0.035, spell_power=2.269 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 25.0 attack_power points (1.09 DPS) | yes | Totemic Leather Hood (252448, -0.20 DPS) [crafted]; Veteran's Silvered Chain Helm (250528, -0.22 DPS) [crafted]; Defender's Leather Hood (252447, -0.39 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 7.3 attack_power points (0.32 DPS) | yes | Erudite's Amulet (277204, -0.18 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.8 attack_power points (0.95 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.55 DPS) [crafted]; Serpent's Shoulders (5404, -0.68 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.05 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop]; Dark Leather Cloak (2316, -0.06 DPS) [crafted] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 31.1 attack_power points (1.36 DPS) | yes | Crusader's Chain Shirt (250492, -0.12 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.31 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.48 DPS) [dungeon] |
| wrist | Death Bindings (286981) | Ruins of Lordaeron: Skeleton [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Patterned Bronze Bracers (2868, -0.12 DPS, sim-verified) [crafted]; Tabitha's Cuffs (251486, -0.14 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 23.2 attack_power points (1.02 DPS) | yes | Totemic Leather Gloves (252497, -0.17 DPS) [crafted]; Crusader's Gloves (250512, -0.26 DPS) [crafted]; Thorbia's Gauntlets (12994, -0.32 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Cobrahn's Grasp (6460, -0.01 DPS) [dungeon]; Crusader's Chain Belt (250517, -0.03 DPS) [crafted]; Totemic Leather Belt (252431, -0.04 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 38.5 attack_power points (1.68 DPS) | yes | Crusader's Chain Leggings (250497, -0.24 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.39 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 23.6 attack_power points (1.03 DPS) | yes | Crusader's Boots (250507, -0.09 DPS) [crafted]; Spidersilk Boots (4320, -0.14 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.18 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.9 attack_power points (0.56 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon]; Demon Band (12054, -0.21 DPS) [world_drop]; Loop of Sacrifice (281673, -0.57 DPS, sim-verified) [quest] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Demon Band (12054, -0.15 DPS) [world_drop]; Loop of Sacrifice (281673, -0.33 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.0 attack_power points (13.26 DPS) | yes | Smite's Mighty Hammer (7230, -0.01 DPS) [dungeon]; Hammerbone (270018, -0.12 DPS) [quest]; Sword of the Fallen (286977, -0.20 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Totemic Leather Armor; wrist: Death Bindings; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Legionnaire's Band; finger2: Advisor's Ring; main_hand: Forsaken Greataxe

No-known-source sample (15 of 252, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 20430 Legionnaire's Sword; 23173 Abomination Skin Leggings; 205218 Libram of Discovery

### Band 30 (undead, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 73.6. Weights run: 7.5s. Verify run: 18.9s. 466 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.385 ± 0.023, crit=1.082 ± 0.023 per rating point (14 rating = 1%, 15.151 per %), hit=1.648 ± 0.065 per rating point (10 rating = 1%, 16.480 per %), melee_haste=8.503 ± 0.797, intellect=1.083 ± 0.034, spell_power=2.699 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 59.8 attack_power points (2.84 DPS) | yes | Totemic Leather Helm (252456, -1.12 DPS, sim-verified) [crafted]; Barbaric Iron Helm (7915, -1.39 DPS) [crafted]; Holy Shroud (2721, -1.43 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Archmagister's Faceted Pendant (273049, -0.13 DPS, sim-verified) [world]; Scout's Medallion (19537, -0.32 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 34.0 attack_power points (1.62 DPS) | yes | Death Speaker Mantle (6685, -0.37 DPS, sim-verified) [dungeon]; Forest Tracker Epaulets (2278, -0.42 DPS) [world_drop]; Invoker's Mantle (215365, -0.46 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.10 DPS) [crafted]; Windsong Drape (15468, -0.10 DPS) [quest]; Tigerstrike Mantle (13108, -0.22 DPS) [world_drop] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 44.4 attack_power points (2.11 DPS) | yes | Green Silk Armor (7065, -0.29 DPS) [crafted]; Hard Gold Cuirass (250533, -0.29 DPS) [crafted]; Totemic Leather Tunic (252451, -0.30 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 24.3 attack_power points (1.16 DPS) | yes | Yorgen Bracers (13012, -0.39 DPS) [world_drop]; Pugilist Bracers (4438, -0.39 DPS) [dungeon]; Hawkeye's Bracers (14590, -1.05 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 27.2 attack_power points (1.29 DPS) | yes | Oilrag Handwraps (16741, -0.04 DPS) [quest]; Insignia Gloves (6408, -0.23 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.25 DPS) [world] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 44.4 attack_power points (2.11 DPS) | yes | Skirmisher's Leather Belt (252461, -0.40 DPS, sim-verified) [crafted]; Skycaller's Leather Belt (252522, -0.52 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.54 DPS) [rep] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 49.1 attack_power points (2.33 DPS) | yes | Totemic Leather Leggings (252458, -0.26 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -0.37 DPS) [crafted]; Ferine Leggings (6690, -0.48 DPS) [dungeon] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 32.9 attack_power points (1.56 DPS) | yes | Gilded Slippers (254001, -0.31 DPS) [crafted]; Totemic Leather Boots (252442, -0.32 DPS) [crafted]; Crusader's Boots (250507, -0.41 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 21.6 attack_power points (1.03 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Advisor's Ring (19521, -0.13 DPS) [rep]; Ring of Power Regulation (273030, -0.26 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 20.3 attack_power points (0.97 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Advisor's Ring (19521, -0.07 DPS) [rep]; Ring of Power Regulation (273030, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.46 DPS) [dungeon]; Felbough Arm (273037, -0.52 DPS) [world]; Viscous Hammer (13045, -4.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Hawkeye's Cloak; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Snake Hoop; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 466, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (undead, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 122.3. Weights run: 7.6s. Verify run: 20.5s. 641 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=2.542 ± 0.053, crit=2.572 ± 0.062 per rating point (14 rating = 1%, 36.013 per %), hit=2.658 ± 0.119 per rating point (10 rating = 1%, 26.579 per %), melee_haste=12.529 ± 1.361, intellect=3.010 ± 0.044, spell_power=2.942 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 74.7 attack_power points (3.97 DPS) | yes | Big Voodoo Mask (8201, -0.33 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.59 DPS) [crafted]; Raging Berserker's Helm (7719, -0.65 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 44.7 attack_power points (2.37 DPS) | yes | Kaleidoscope Chain (13084, -0.76 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.78 DPS) [world]; Scout's Medallion (19536, -0.89 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 59.7 attack_power points (3.17 DPS) | yes | Green Silken Shoulders (7057, -0.16 DPS) [crafted]; Bloodmage Mantle (7684, -0.33 DPS) [dungeon]; Shining Mithril Pauldrons (250541, -0.37 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.02 DPS) [crafted]; Guardian Cloak (5965, -0.02 DPS) [crafted]; Blackforge Cape (6424, -0.48 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.15 DPS) [crafted]; Robe of Power (7054, -0.29 DPS) [crafted]; Icemail Jerkin (1981, -0.59 DPS, sim-verified) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 35.9 attack_power points (1.91 DPS) | yes | Radiant Silver Bracers (4545, -0.00 DPS) [quest]; Guardian Leather Bracers (4260, -0.01 DPS) [crafted]; Turtle Scale Bracers (8198, -0.17 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Gloves of the Greatfather (17721, +0.00 DPS) [crafted]; Dreamweave Gloves (10019, -0.77 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 65.7 attack_power points (3.49 DPS) | yes | Defiler's Cloth Girdle (20166, -0.56 DPS, sim-verified) [rep]; Skirmisher's Leather Belt (252461, -0.78 DPS) [crafted]; Gilded Cord (254037, -0.96 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 77.3 attack_power points (4.11 DPS) | yes | Crimson Silk Pantaloons (7062, -0.67 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.81 DPS) [dungeon]; Totemic Leather Leggings (252458, -0.92 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 70.6 attack_power points (3.75 DPS) | yes | Skirmisher's Mail Boots (252564, -0.35 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.60 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.60 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 44.5 attack_power points (2.36 DPS) | yes | Snake Hoop (6750, -0.63 DPS) [quest]; Ring of the Underwood (2951, -0.80 DPS) [world_drop]; Falcon's Hook (7552, -0.82 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 36.7 attack_power points (1.95 DPS) | yes | Snake Hoop (6750, -0.21 DPS) [quest]; Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.41 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, -0.16 DPS) [vendor]; Nightblade (1982, -1.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Legionnaire's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 641, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (undead, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 176.7. Weights run: 7.8s. Verify run: 28.8s. 837 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.866 ± 0.057, crit=3.362 ± 0.073 per rating point (14 rating = 1%, 47.066 per %), hit=2.687 ± 0.131 per rating point (10 rating = 1%, 26.869 per %), melee_haste=12.895 ± 1.525, intellect=3.744 ± 0.076, spell_power=3.243 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 136.5 attack_power points (6.95 DPS) | yes | Soothsayer's Headdress (17740, -0.29 DPS) [dungeon]; Eye of Theradras (17715, -0.93 DPS) [dungeon]; Dreamweave Circlet (10041, -1.58 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Skibi's Pendant (13089, -0.21 DPS) [world_drop]; Mindburst Medallion (11196, -0.53 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.59 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 123.5 attack_power points (6.29 DPS) | yes | Rotgrip Mantle (17732, -1.14 DPS, sim-verified) [dungeon]; Earthslag Shoulders (11632, -1.49 DPS) [dungeon]; Rockshard Pauldrons (9411, -1.49 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.24 DPS) [dungeon]; Runecloth Cloak (13860, -0.69 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.03 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 136.5 attack_power points (6.95 DPS) | yes | Spiderfang Carapace (11633, -0.76 DPS) [dungeon]; Runecloth Robe (13858, -1.89 DPS) [crafted]; Feathered Breastplate (8349, -1.91 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 69.9 attack_power points (3.56 DPS) | yes | Skirmisher's Mail Bracers (252572, -0.28 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.57 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.57 DPS) [crafted] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Skirmisher's Mail Gauntlets (252586, -0.53 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.70 DPS) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | 101.2 attack_power points (5.15 DPS) | yes | Skycaller's Mail Belt (252589, -0.22 DPS) [crafted]; Defiler's Cloth Girdle (20165, -0.32 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.07 DPS, sim-verified) [crafted] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 102.1 attack_power points (5.20 DPS) | yes | Imperial Red Pants (8251, -0.15 DPS) [world_drop]; Stormshroud Pants (15057, -0.41 DPS) [crafted]; Red Mageweave Pants (10009, -0.60 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 100.0 attack_power points (5.09 DPS) | yes | Skycaller's Leather Boots (252471, -0.52 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.52 DPS) [crafted]; Earthen Silk Slippers (254013, -1.13 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 51.1 attack_power points (2.60 DPS) | yes | Blackstone Ring (17713, -0.22 DPS) [dungeon]; Band of the Unicorn (7553, -0.46 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.56 DPS) [quest] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 47.8 attack_power points (2.43 DPS) | yes | Blackstone Ring (17713, -0.05 DPS) [dungeon]; Band of the Unicorn (7553, -0.29 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.39 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.68 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.24 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Warmonger (13052, -3.39 DPS) [world_drop]; The Rockpounder (9413, -4.12 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Stonerender Gauntlets; waist: Defiler's Lamellar Girdle; legs: Stone Guard's Plate Leggings; feet: Greaves of Withering Despair; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 837, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (undead, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 280.2. Weights run: 7.9s. Verify run: 95.0s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.921 ± 0.047, crit=2.904 ± 0.066 per rating point (14 rating = 1%, 40.653 per %), hit=3.042 ± 0.168 per rating point (10 rating = 1%, 30.416 per %), melee_haste=15.715 ± 2.674, intellect=6.883 ± 0.113, spell_power=3.217 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 342.6 attack_power points (22.64 DPS) | yes | Magister's Crown (16686, -6.37 DPS, sim-verified) [dungeon]; Blue Dragonscale Helm (252604, -7.13 DPS) [crafted]; Coif of Elements (16667, -8.61 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 107.4 attack_power points (7.10 DPS) | yes | Horizon Choker (13085, -0.73 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.97 DPS) [dungeon]; Darkspear Warding Pendant (272072, -2.55 DPS) [vendor] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 237.3 attack_power points (15.69 DPS) | yes | Darkspear Shoulderpads (272103, -0.85 DPS) [vendor]; Darkspear Shoulders (272104, -0.85 DPS) [vendor]; Ironfeather Shoulders (15067, -3.40 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.63 DPS) [quest]; Spritecaster Cape (11623, -1.57 DPS) [dungeon]; Runecloth Cloak (13860, -1.73 DPS) [crafted] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 231.3 attack_power points (15.29 DPS) | yes | Blue Dragonscale Breastplate (15048, -0.85 DPS) [crafted]; Vest of Elements (16666, -1.30 DPS) [dungeon]; Mooncloth Robe (18486, -2.00 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 125.8 attack_power points (8.32 DPS) | yes | Spitfire Bracers (20481, -1.38 DPS) [crafted]; Runecloth Cuffs (254123, -1.43 DPS) [crafted]; Magiskull Cuffs (13107, -2.17 DPS, sim-verified) [world_drop] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handguards (272101, -1.03 DPS) [vendor]; Spitfire Gauntlets (20480, -4.66 DPS) [crafted] |
| waist | Lightforge Belt (16723) | Stratholme: Rockwing Gargoyle [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Lizardhide Girdle (20171, -0.35 DPS) [rep]; Defiler's Mail Girdle (20195, -0.35 DPS) [rep]; Magister's Belt (16685, -1.12 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Lizardhide Pants (22741, -2.14 DPS) [rep]; Outrider's Mail Leggings (22676, -4.22 DPS) [rep]; Red Dragonscale Leggings (252603, -4.35 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 146.9 attack_power points (9.71 DPS) | yes | Greaves of Withering Despair (22240, -0.81 DPS) [dungeon]; Skycaller's Leather Boots (252471, -1.52 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.52 DPS) [crafted] |
| finger1 | Deep Rooted Ring (19109) | Ivus the Forest Lord [world] | 107.4 attack_power points (7.10 DPS) | yes | Glowing Crystal Ring (18402, -1.64 DPS) [quest]; Frigid Ring (18679, -2.55 DPS) [world]; Maiden's Circle (13001, -5.20 DPS, sim-verified) [world_drop] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Crystal Ring (18402, -0.16 DPS) [quest]; Frigid Ring (18679, -1.07 DPS) [world]; Maiden's Circle (13001, -2.40 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.06 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.53 DPS) [world_drop] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Needler (13060, +0.00 DPS) [world_drop]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, +0.00 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; shoulder: Darkspear Shoulderguards; back: Darkspear Raider's Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Lightforge Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Deep Rooted Ring; finger2: Don Julio's Band; main_hand: Forest Defender's Axe

No-known-source sample (15 of 1670, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (undead, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 793.1. Weights run: 8.0s. Verify run: 95.6s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=1.688 ± 0.032, crit=2.253 ± 0.043 per rating point (14 rating = 1%, 31.538 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=24.905 ± 4.031, intellect=1.725 ± 0.062, spell_power=1.760 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | sim-verified (793.1 DPS) | yes | Stalwart Helm (250599, -2.03 DPS) [crafted]; Outlaw's Collar (279253, -4.51 DPS) [crafted]; Black Dragonscale Helm (252605, -13.76 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 38.4 attack_power points (7.25 DPS) | yes | Skibi's Pendant (13089, -0.82 DPS) [world_drop]; Woven Ivy Necklace (19159, -1.64 DPS) [quest]; Glowing Eye of Mordresh (10769, -2.32 DPS) [dungeon] |
| shoulder | Defiler's Lamellar Spaulders (20184) | The Defilers [rep] | 89.5 attack_power points (16.92 DPS) | yes | Mantle of the Timbermaw (19050, -1.06 DPS) [crafted]; Darkspear Pauldrons (272105, -2.34 DPS) [vendor]; Darkspear Epaulets (272106, -2.34 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Flames (3475, +0.00 DPS) [world_drop]; Deathguard's Cloak (20068, -0.55 DPS) [rep]; Deep Woodlands Cloak (19121, -1.64 DPS) [quest] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 141.8 attack_power points (26.80 DPS) | yes | Savage Gladiator Chain (11726, -5.50 DPS, sim-verified) [dungeon]; Dawn Armor (252483, -9.03 DPS) [crafted]; Spitfire Breastplate (20479, -11.49 DPS) [crafted] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | 59.5 attack_power points (11.24 DPS) | yes | Forest Stalker's Bracers (19587, -0.15 DPS) [rep]; Tranquil Wristguards (279256, -0.66 DPS) [crafted]; Dryad's Wrist Bindings (19595, -1.32 DPS) [rep] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, -1.44 DPS) [world_drop]; Timbermaw Brawlers (19049, -1.79 DPS) [crafted]; Gloves of Holy Might (867, -5.67 DPS, sim-verified) [world_drop] |
| waist | Defiler's Lamellar Girdle (20177) | The Defilers [rep] | 85.1 attack_power points (16.08 DPS) | yes | Might of the Timbermaw (19044, -2.01 DPS) [crafted]; Defiler's Plate Girdle (20204, -2.35 DPS) [rep]; Defiler's Cloth Girdle (20163, -3.51 DPS) [rep] |
| legs | Champion's Legplates (250598) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pristine Scorpid Leggings (252606, -1.66 DPS) [crafted]; Ironfeather Leggings (252486, -2.36 DPS) [crafted]; Outrider's Lizardhide Pants (22741, -2.75 DPS) [rep] |
| feet | Depleted Thorium Sabatons (279263) | Blacksmithing [crafted] | 70.8 attack_power points (13.39 DPS) | yes | Defiler's Lamellar Greaves (20181, -0.55 DPS) [rep]; Omnicast Boots (11822, -2.82 DPS) [dungeon]; Skirmisher's Mail Sabatons (252578, -3.12 DPS) [crafted] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | 47.6 attack_power points (9.00 DPS) | yes | Maiden's Circle (13001, -0.73 DPS) [world_drop]; Deep Rooted Ring (19109, -1.75 DPS) [world]; Band of the Penitent (13217, -3.04 DPS) [quest] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 45.5 attack_power points (8.61 DPS) | yes | Deep Rooted Ring (19109, -1.35 DPS) [world]; Band of the Penitent (13217, -2.65 DPS) [quest]; Maiden's Circle (13001, -12.96 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+9.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.65 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Molten Heart of the Mountain (249470, -8.15 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Arcanite Champion (12790, -4.38 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Archlight Talisman; shoulder: Defiler's Lamellar Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Defiler's Lamellar Girdle; legs: Champion's Legplates; feet: Depleted Thorium Sabatons; finger1: Legionnaire's Band; finger2: Don Julio's Band; main_hand: The Unstoppable Force

No-known-source sample (15 of 1670, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

