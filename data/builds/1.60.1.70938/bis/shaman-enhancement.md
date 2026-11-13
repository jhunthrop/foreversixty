# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-053030000000000000-0000000000000000)

Set DPS (verified): 38.3. Weights run: 4.5s. Verify run: 12.0s. 286 eligible items had no known source.

Race sweep: dwarf 38.3 ± 0.0 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.902 ± 0.021, crit=0.755 ± 0.022 per rating point (14 rating = 1%, 10.566 per %), hit=1.375 ± 0.059 per rating point (10 rating = 1%, 13.747 per %), melee_haste=6.350 ± 0.202, intellect=2.340 ± 0.023, spell_power=1.470 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Trapper's Leather Hood (252505) | Leatherworking [crafted] | 18.7 attack_power points (0.65 DPS) | yes | Defender's Leather Hood (252447, -0.09 DPS) [crafted]; Totemic Leather Hood (252448, -0.19 DPS) [crafted]; Shadow Goggles (4373, -0.25 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 5.4 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS) [quest]; Tarnished Locket (279870, -0.40 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 28.4 attack_power points (0.99 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.41 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.79 DPS) [crafted]; Serpent's Shoulders (5404, -0.83 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, +0.00 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop]; Dark Leather Cloak (2316, -0.09 DPS) [crafted] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mystic's Wrap (14369, +0.00 DPS) [world_drop]; Totemic Leather Armor (252435, +0.00 DPS) [crafted]; Stormrider's Leather Armor (252492, +0.00 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.08 DPS) [world_drop]; Bravo's Armbands (270015, -0.10 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Totemic Leather Gloves (252497, +0.00 DPS) [crafted]; Stormrider's Leather Gloves (252498, +0.00 DPS) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 18.7 attack_power points (0.65 DPS) | yes | Blackened Defias Belt (10403, -0.03 DPS) [dungeon]; Stormrider's Leather Belt (252432, -0.12 DPS) [crafted]; Pristine Sash (253925, -0.12 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 37.7 attack_power points (1.31 DPS) | yes | Trapper's Leather Pants (252501, -0.29 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.31 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.47 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 20.5 attack_power points (0.72 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Totemic Leather Boots (252442, -0.06 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.10 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 15.8 attack_power points (0.55 DPS) | yes | Minor Channeling Ring (1449, -0.13 DPS) [quest]; Demon Band (12054, -0.27 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.30 DPS) [rep] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Minor Channeling Ring (1449, +0.00 DPS) [quest]; Demon Band (12054, -0.13 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.15 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 309.6 attack_power points (10.81 DPS) | yes | Smite's Mighty Hammer (7230, -0.28 DPS) [dungeon]; Rockslicer (872, -0.69 DPS) [dungeon]; Night Reaver (1318, -1.18 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Trapper's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Armor of the Fang; wrist: Mindthrust Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Totemic Leather Pants; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish

### Band 30 (dwarf, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 85.5. Weights run: 4.5s. Verify run: 10.1s. 475 eligible items had no known source.

Race sweep: dwarf 85.5 ± 0.1 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.151 ± 0.031, crit=1.101 ± 0.035 per rating point (14 rating = 1%, 15.407 per %), hit=1.605 ± 0.069 per rating point (10 rating = 1%, 16.053 per %), melee_haste=6.329 ± 0.473, intellect=2.455 ± 0.020, spell_power=0.719 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 39.3 attack_power points (2.77 DPS) | yes | Dokebi Hat (14584, -0.30 DPS) [world_drop]; Nightsky Cowl (4039, -0.69 DPS) [world_drop]; Resilient Cap (14401, -0.70 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 26.7 attack_power points (1.89 DPS) | yes | Kaleidoscope Chain (13084, -0.21 DPS, sim-verified) [world_drop]; Seal of the Expedition (271920, -1.02 DPS) [rep]; Crystal Starfire Medallion (5003, -1.19 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.3 attack_power points (2.21 DPS) | yes | Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Magician's Mantle (12998, -0.40 DPS) [world_drop]; Feline Mantle (3748, -0.46 DPS) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271923, -0.20 DPS) [rep]; Resilient Cape (14400, -0.54 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.59 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 47.2 attack_power points (3.33 DPS) | yes | Spirewind Fetter (9406, -0.26 DPS) [dungeon]; Green Silk Armor (7065, -0.62 DPS) [crafted]; Beguiler Robes (7728, -0.79 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 25.6 attack_power points (1.81 DPS) | yes | Technician's Bracers (270042, -0.55 DPS) [quest]; Hawkeye's Bracers (14590, -0.76 DPS) [world_drop]; Nightsky Wristbands (6407, -0.77 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 24.2 attack_power points (1.71 DPS) | yes | Insignia Gloves (6408, -0.24 DPS) [world_drop]; Town Clerk's Mittens (270029, -0.35 DPS) [quest]; Toughened Leather Gloves (4253, -0.38 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 34.7 attack_power points (2.45 DPS) | yes | Skirmisher's Leather Belt (252461, -0.13 DPS) [crafted]; Highlander's Lizardhide Girdle (20105, -0.37 DPS) [rep]; Highlander's Mail Girdle (20120, -0.37 DPS) [vendor] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 55.5 attack_power points (3.91 DPS) | yes | Totemic Leather Leggings (252458, -1.25 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -1.60 DPS) [crafted]; Trapper's Leather Pants (252501, -1.64 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 23.2 attack_power points (1.64 DPS) | yes | Gilded Slippers (254001, -0.07 DPS) [crafted]; Frothing Slippers (254003, -0.43 DPS) [crafted]; Fiery Slippers (254005, -0.43 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 21.5 attack_power points (1.52 DPS) | yes | Protector's Band (19517, -0.18 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Thunderbrow Ring (13097, -1.45 DPS, sim-verified) [world_drop] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Protector's Band (19517, -0.03 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Thunderbrow Ring (13097, -0.23 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -0.89 DPS, sim-verified) [dungeon]; Corpsemaker (6687, -1.92 DPS) [dungeon]; Cobalt Crusher (7730, -4.07 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Sergeant Major's Cape; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Snake Hoop; finger2: Seal of Wrynn; main_hand: Viscous Hammer

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 0000000000000000-055030310051002051-0000000000000000)

Set DPS (verified): 144.5. Weights run: 4.8s. Verify run: 13.6s. 804 eligible items had no known source.

Race sweep: dwarf 144.5 ± 0.2 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.374 ± 0.046, crit=1.475 ± 0.056 per rating point (14 rating = 1%, 20.647 per %), hit=2.233 ± 0.119 per rating point (10 rating = 1%, 22.333 per %), melee_haste=7.729 ± 0.747, intellect=3.618 ± 0.056, spell_power=0.741 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 80.8 attack_power points (5.66 DPS) | yes | Miner's Hat of the Deep (9429, -0.73 DPS) [dungeon]; Papal Fez (9431, -0.94 DPS) [dungeon]; Thinking Cap (2624, -1.35 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 34.1 attack_power points (2.39 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Kaleidoscope Chain (13084, -0.43 DPS) [world_drop]; Seal of the Expedition (271919, -0.62 DPS) [rep] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 57.9 attack_power points (4.05 DPS) | yes | Inquisitor's Shawl (19507, -0.40 DPS) [dungeon]; Windchaser Amice (14432, -0.47 DPS) [world_drop]; Green Silken Shoulders (7057, -0.85 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackforge Cape (6424, -0.91 DPS) [dungeon]; Expeditionary's Cape (271923, -0.91 DPS) [rep] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 66.6 attack_power points (4.67 DPS) | yes | Red Mageweave Vest (10007, -0.49 DPS, sim-verified) [crafted]; Silksand Tunic (14417, -0.61 DPS) [world_drop]; Silksand Wraps (14425, -0.61 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 35.3 attack_power points (2.47 DPS) | yes | Glowing Magical Bracelets (13106, -0.03 DPS) [world_drop]; Warden's Wristbands (14600, -0.16 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.28 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Engineering Gloves (1659, +0.00 DPS) [world]; Stormcloth Gloves (10011, +0.00 DPS) [crafted]; Tough Scorpid Gloves (8204, -2.67 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 59.5 attack_power points (4.16 DPS) | yes | Highlander's Lizardhide Girdle (20104, -0.36 DPS) [rep]; Highlander's Mail Girdle (20119, -0.36 DPS) [vendor]; Saltscale Girdle (273023, -0.62 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 69.4 attack_power points (4.86 DPS) | yes | Red Mageweave Pants (10009, -1.09 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.15 DPS) [crafted]; Misplaced Pantaloons (276201, -1.57 DPS) [vendor] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 49.7 attack_power points (3.48 DPS) | yes | Skirmisher's Mail Boots (252564, -0.21 DPS) [crafted]; Stalker's Mail Boots (252562, -0.65 DPS) [crafted]; Briar Tredders (10582, -0.72 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 29.8 attack_power points (2.08 DPS) | yes | Philanthropist's Ring (281635, -0.30 DPS) [quest]; Ogremind Ring (1993, -0.31 DPS) [world_drop]; Voodoo Band (1996, -0.31 DPS) [world] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | 27.0 attack_power points (1.89 DPS) | yes | Philanthropist's Ring (281635, -0.11 DPS) [quest]; Ogremind Ring (1993, -0.12 DPS) [world_drop]; Voodoo Band (1996, -0.12 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonebiter (6830, -0.86 DPS) [quest]; Staff of Jordan (873, -1.92 DPS) [world_drop]; Mograine's Might (7723, -7.56 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Sheepshear Mantle; back: Sergeant Major's Cape; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger2: Protector's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 1000000000000000-055030310051202251-0500000000000000)

Set DPS (verified): 222.6. Weights run: 4.9s. Verify run: 14.5s. 1050 eligible items had no known source.

Race sweep: dwarf 222.6 ± 0.3 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.216 ± 0.052, crit=1.609 ± 0.069 per rating point (14 rating = 1%, 22.525 per %), hit=2.642 ± 0.143 per rating point (10 rating = 1%, 26.423 per %), melee_haste=9.607 ± 0.998, intellect=4.234 ± 0.071, spell_power=0.762 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 128.0 attack_power points (10.44 DPS) | yes | Braincage (12549, -1.33 DPS, sim-verified) [dungeon]; Soulcatcher Halo (10630, -1.44 DPS) [dungeon]; Eye of Theradras (17715, -2.04 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 59.3 attack_power points (4.83 DPS) | yes | Glowing Eye of Mordresh (10769, -1.64 DPS) [dungeon]; Seal of the Expedition (271918, -1.73 DPS) [rep]; Gemshard Heart (17707, -1.99 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 96.1 attack_power points (7.84 DPS) | yes | Rotgrip Mantle (17732, -0.81 DPS) [dungeon]; Dregmetal Spaulders (11722, -1.69 DPS, sim-verified) [dungeon]; Splinthide Shoulders (11685, -2.03 DPS) [dungeon] |
| back | Expeditionary's Cape (271921) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16336, -1.19 DPS) [pvp]; Keeper's Cloak (14665, -1.38 DPS) [world_drop]; Imperial Red Cloak (8248, -2.06 DPS, sim-verified) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 110.0 attack_power points (8.97 DPS) | yes | Knight's Mail Armor (223078, -1.03 DPS) [vendor]; Hibernal Robe (8113, -2.06 DPS) [world_drop]; Runecloth Robe (13858, -2.42 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 66.6 attack_power points (5.43 DPS) | yes | Shizzle's Nozzle Wiper (11917, -1.28 DPS) [quest]; Runic Leather Bracers (15092, -1.35 DPS) [crafted]; Forgotten Wraps (9433, -1.83 DPS, sim-verified) [world_drop] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Rockgrip Gauntlets (17736, -0.16 DPS) [dungeon]; Prowler's Leather Gauntlets (252547, -0.99 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 85.0 attack_power points (6.93 DPS) | yes | Serenity Belt (13144, -1.06 DPS) [world_drop]; Imperial Red Sash (8253, -1.10 DPS) [world_drop]; Ban'thok Sash (11662, -3.05 DPS, sim-verified) [dungeon] |
| legs | Knight's Mail Legplates (223074) | PvP rank 6 · Knight · Alliance [vendor] | 97.3 attack_power points (7.94 DPS) | yes | Imperial Red Pants (8251, -1.27 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -1.54 DPS) [dungeon]; Infernal Trickster Leggings (17754, -2.39 DPS, sim-verified) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 77.1 attack_power points (6.29 DPS) | yes | Bloodshot Greaves (10846, -1.41 DPS, sim-verified) [world]; Coldstone Slippers (18697, -1.46 DPS) [dungeon]; Skycaller's Leather Boots (252471, -1.56 DPS) [crafted] |
| finger1 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 64.0 attack_power points (5.22 DPS) | yes | Woodseed Hoop (17768, -1.40 DPS) [quest]; Blackstone Ring (17713, -1.44 DPS) [dungeon]; Cyclopean Band (11824, -1.59 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 50.7 attack_power points (4.14 DPS) | yes | Blackstone Ring (17713, -0.35 DPS) [dungeon]; Cyclopean Band (11824, -0.51 DPS) [dungeon]; Woodseed Hoop (17768, -1.68 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -1.39 DPS, sim-verified) [world_drop]; Uther's Strength (11302, -1.57 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Soulkeeper (1607, -1.95 DPS) [world_drop]; Taran Icebreaker (2915, -2.14 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Expeditionary's Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Knight's Mail Legplates; feet: Greaves of Withering Despair; finger1: Mindseye Circle; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 1050, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 5030000000000000-055030310051202251-0530000000000000)

Set DPS (verified): 336.8. Weights run: 4.6s. Verify run: 49.4s. 2374 eligible items had no known source.

Race sweep: dwarf 336.8 ± 0.4 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.695 ± 0.081, crit=2.830 ± 0.122 per rating point (14 rating = 1%, 39.620 per %), hit=4.492 ± 0.250 per rating point (10 rating = 1%, 44.918 per %), melee_haste=14.989 ± 1.788, intellect=3.963 ± 0.104, spell_power=0.868 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 225.8 attack_power points (18.21 DPS) | yes | Cap of the Scarlet Savant (12752, -5.37 DPS, sim-verified) [quest]; Chief Architect's Monocle (11839, -6.92 DPS) [dungeon]; Blue Dragonscale Helm (252604, -7.09 DPS) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nacreous Shell Necklace (22403, -0.03 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.67 DPS) [world_drop]; Star of Mystaria (12103, -1.86 DPS, sim-verified) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Theramore Shoulderpads (271953, -0.28 DPS) [rep]; Theramore Shoulders (271954, -0.28 DPS) [rep]; Soulstealer Mantle (13374, -3.36 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -0.16 DPS) [dungeon]; Howler's Furs (272414, -1.42 DPS) [vendor]; Royal Tribunal Cloak (13376, -1.84 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Magister's Robes (16688, -1.19 DPS) [dungeon]; Blue Dragonscale Breastplate (15048, -1.47 DPS) [crafted]; Dawn Armor (252483, -4.18 DPS, sim-verified) [crafted] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | 69.7 attack_power points (5.62 DPS) | yes | Manacle Cuffs (11962, -0.34 DPS) [quest]; Aristocratic Cuffs (12546, -0.55 DPS) [dungeon]; Magiskull Cuffs (13107, -2.29 DPS, sim-verified) [world_drop] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Bloodmail Gauntlets (14615, -1.02 DPS) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Mail Girdle (20044, -0.16 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, -0.16 DPS) [rep]; Clutch of Andros (13956, -3.36 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Lizardhide Pants (22750) | Silverwing Sentinels [rep] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Sentinel's Silk Leggings (237815, -0.89 DPS) [vendor]; Red Dragonscale Leggings (252603, -1.09 DPS) [crafted]; Sentinel's Chain Leggings (237819, -4.72 DPS, sim-verified) [vendor] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 119.7 attack_power points (9.66 DPS) | yes | Bloodmail Boots (14616, -0.29 DPS) [dungeon]; Windreaver Greaves (13967, -2.02 DPS) [dungeon]; Greaves of Withering Despair (22240, -2.07 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -2.10 DPS) [dungeon]; Mindseye Circle (10634, -2.45 DPS) [dungeon]; Naglering (11669, -7.02 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -1.95 DPS) [dungeon]; Mindseye Circle (10634, -2.30 DPS) [dungeon]; Cutthroat's Signet (272408, -2.41 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -5.75 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Rage (22395) | Blackrock Depths: Magmus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, -1.91 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Theramore Shoulderguards; back: Arcanoweave Cloak; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Theramore Handwraps; waist: Stormseeker's Girdle; legs: Sentinel's Lizardhide Pants; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Totem of Rage

No-known-source sample (15 of 2374, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 0533001000000000-053030310051202051-0530000000000000)

Set DPS (verified): 874.4. Weights run: 4.8s. Verify run: 49.3s. 2374 eligible items had no known source.

Race sweep: dwarf 874.4 ± 1.2 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=1.339 ± 0.052, crit=2.236 ± 0.073 per rating point (14 rating = 1%, 31.306 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=19.365 ± 2.160, intellect=1.583 ± 0.012, spell_power=0.806 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 119.6 attack_power points (26.24 DPS) | yes | Cap of the Scarlet Savant (12752, -5.87 DPS, sim-verified) [quest]; Eye of Rend (12587, -8.70 DPS) [dungeon]; Ragefury Eyepatch (11735, -9.61 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 57.3 attack_power points (12.57 DPS) | yes | Lady Maye's Pendant (14558, -4.03 DPS) [world_drop]; Emberfury Talisman (12929, -4.64 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -5.26 DPS) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 67.5 attack_power points (14.80 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Mantle of the Timbermaw (19050, -0.41 DPS) [crafted]; Highlander's Lizardhide Shoulders (20060, -0.53 DPS) [rep] |
| back | Sergeant Major's Cape (16337) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.06 DPS) [rep]; Cape of the Black Baron (13340, -0.20 DPS) [dungeon]; Blackveil Cape (11626, -1.98 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 128.0 attack_power points (28.08 DPS) | yes | Savage Gladiator Chain (11726, -5.37 DPS, sim-verified) [dungeon]; Breastplate of Bloodthirst (12757, -8.07 DPS) [quest]; Tombstone Breastplate (13944, -8.56 DPS) [dungeon] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | 50.7 attack_power points (11.12 DPS) | yes | Forest Stalker's Bracers (19587, -0.22 DPS) [rep]; Prowler's Leather Bracers (252539, -2.98 DPS) [crafted] |
| hands | Bloodmail Gauntlets (14615) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Brawlers (19049, -1.40 DPS) [crafted]; Stonerender Gauntlets (17007, -1.98 DPS) [world_drop]; Theramore Gloves (271948, -2.41 DPS) [rep] |
| waist | Windseeker's Belt (272410) | Pix Xizzix [vendor] | 65.7 attack_power points (14.41 DPS) | yes | Highlander's Chain Girdle (20043, -0.09 DPS) [rep]; Highlander's Leather Girdle (20045, -0.09 DPS) [rep]; Might of the Timbermaw (19044, -0.16 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 109.5 attack_power points (24.02 DPS) | yes | Sentinel's Silk Leggings (237815, -1.15 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.15 DPS) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | sim-verified (+12.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Boots (252468, -0.66 DPS) [crafted]; Skulker's Leather Boots (252469, -1.04 DPS) [crafted]; Scalegut Treaders (275618, -12.43 DPS, sim-verified) [crafted] |
| finger1 | Magni's Will (12548) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.75 DPS) [dungeon]; Protector's Band (19514, -1.22 DPS) [rep]; Naglering (11669, -11.38 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -0.92 DPS) [rep]; Don Mauricio's Band of Domination (22433, -1.13 DPS) [dungeon]; Painweaver Band (13098, -10.30 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -1.40 DPS) [vendor] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS) [vendor]; Weakness Analyzer (272438, -9.68 DPS, sim-verified) [vendor] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Seeping Willow (12969, -19.52 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Rage (22395) | Blackrock Depths: Magmus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, +0.00 DPS) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Sergeant Major's Cape; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Bloodmail Gauntlets; waist: Windseeker's Belt; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Magni's Will; finger2: Don Julio's Band; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Totem of Rage

No-known-source sample (15 of 2374, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 0000000000000000-053030000000000000-0000000000000000)

Set DPS (verified): 38.5. Weights run: 4.5s. Verify run: 13.3s. 275 eligible items had no known source.

Race sweep: windshaper-skyborne 38.8 ± 0.0, orc 38.5 ± 0.0, troll 38.5 ± 0.0, tauren 38.2 ± 0.0 (dps); windshaper-skyborne beats the guide race orc by 0.3 (0.84%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.902 ± 0.021, crit=0.755 ± 0.022 per rating point (14 rating = 1%, 10.566 per %), hit=1.375 ± 0.059 per rating point (10 rating = 1%, 13.747 per %), melee_haste=6.350 ± 0.202, intellect=2.340 ± 0.023, spell_power=1.470 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Trapper's Leather Hood (252505) | Leatherworking [crafted] | 18.7 attack_power points (0.65 DPS) | yes | Defender's Leather Hood (252447, -0.09 DPS) [crafted]; Totemic Leather Hood (252448, -0.19 DPS) [crafted]; Shadow Goggles (4373, -0.25 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 5.4 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 28.4 attack_power points (0.99 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.41 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.79 DPS) [crafted]; Serpent's Shoulders (5404, -0.83 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanguine Cape (14376, -0.02 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop]; Dark Leather Cloak (2316, -0.11 DPS) [crafted] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mystic's Wrap (14369, +0.00 DPS) [world_drop]; Totemic Leather Armor (252435, +0.00 DPS) [crafted]; Stormrider's Leather Armor (252492, +0.00 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.08 DPS) [world_drop]; Featherbead Bracers (15452, -0.16 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Totemic Leather Gloves (252497, +0.00 DPS) [crafted]; Stormrider's Leather Gloves (252498, +0.00 DPS) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackened Defias Belt (10403, -0.03 DPS) [dungeon]; Stormrider's Leather Belt (252432, -0.12 DPS) [crafted]; Pristine Sash (253925, -0.12 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 37.7 attack_power points (1.31 DPS) | yes | Trapper's Leather Pants (252501, -0.29 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.31 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.47 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 20.5 attack_power points (0.72 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Totemic Leather Boots (252442, -0.06 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.10 DPS) [crafted] |
| finger1 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 17.7 attack_power points (0.62 DPS) | yes | Legionnaire's Band (20429, -0.21 DPS) [rep]; Demon Band (12054, -0.34 DPS) [world_drop]; Advisor's Ring (20426, -0.36 DPS) [rep] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 15.8 attack_power points (0.55 DPS) | yes | Legionnaire's Band (20429, -0.15 DPS) [rep]; Demon Band (12054, -0.27 DPS) [world_drop]; Advisor's Ring (20426, -0.30 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 309.6 attack_power points (10.81 DPS) | yes | Forsaken Greataxe (251533, -0.23 DPS) [quest]; Smite's Mighty Hammer (7230, -0.28 DPS) [dungeon]; Hammerbone (270018, -0.33 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Trapper's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Armor of the Fang; wrist: Mindthrust Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Totemic Leather Pants; feet: Stormrider's Leather Boots; finger1: Loop of Sacrifice; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 275, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield

### Band 30 (orc, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 85.6. Weights run: 4.5s. Verify run: 10.7s. 469 eligible items had no known source.

Race sweep: windshaper-skyborne 86.5 ± 0.1, orc 85.6 ± 0.1, troll 85.4 ± 0.1, tauren 84.9 ± 0.1 (dps); windshaper-skyborne beats the guide race orc by 0.9 (1.09%) ± 0.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.151 ± 0.031, crit=1.101 ± 0.035 per rating point (14 rating = 1%, 15.407 per %), hit=1.605 ± 0.069 per rating point (10 rating = 1%, 16.053 per %), melee_haste=6.329 ± 0.473, intellect=2.455 ± 0.020, spell_power=0.719 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 39.3 attack_power points (2.77 DPS) | yes | Nightsky Cowl (4039, -0.69 DPS) [world_drop]; Resilient Cap (14401, -0.70 DPS) [world_drop]; Dokebi Hat (14584, -0.74 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 26.7 attack_power points (1.89 DPS) | yes | Kaleidoscope Chain (13084, -0.30 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.02 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.19 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.3 attack_power points (2.21 DPS) | yes | Bloodmage Mantle (7684, -0.19 DPS) [dungeon]; Magician's Mantle (12998, -0.40 DPS) [world_drop]; Feline Mantle (3748, -0.46 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resilient Cape (14400, -0.35 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.39 DPS) [world_drop]; Soft Willow Cape (16661, -0.52 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 47.2 attack_power points (3.33 DPS) | yes | Spirewind Fetter (9406, -0.26 DPS) [dungeon]; Green Silk Armor (7065, -0.62 DPS) [crafted]; Beguiler Robes (7728, -0.79 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 25.6 attack_power points (1.81 DPS) | yes | Hawkeye's Bracers (14590, -0.76 DPS) [world_drop]; Nightsky Wristbands (6407, -0.77 DPS) [world_drop]; Technician's Bracers (270042, -1.00 DPS, sim-verified) [quest] |
| hands | Gloves of Kapelan (6744) | Alliance Relations [quest] | 28.7 attack_power points (2.03 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.32 DPS) [dungeon]; Insignia Gloves (6408, -0.55 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.69 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 34.7 attack_power points (2.45 DPS) | yes | Skirmisher's Leather Belt (252461, -0.13 DPS) [crafted]; Defiler's Lizardhide Girdle (20172, -0.37 DPS) [rep]; Defiler's Mail Girdle (20197, -0.37 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 55.5 attack_power points (3.91 DPS) | yes | Totemic Leather Leggings (252458, -1.32 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -1.60 DPS) [crafted]; Trapper's Leather Pants (252501, -1.64 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 23.2 attack_power points (1.64 DPS) | yes | Gilded Slippers (254001, -0.07 DPS) [crafted]; Frothing Slippers (254003, -0.43 DPS) [crafted]; Fiery Slippers (254005, -0.43 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 21.5 attack_power points (1.52 DPS) | yes | Legionnaire's Band (19513, -0.18 DPS) [rep]; Loop of Sacrifice (281673, -0.23 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.5 attack_power points (1.37 DPS) | yes | Legionnaire's Band (19513, -0.04 DPS) [rep]; Loop of Sacrifice (281673, -0.08 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -0.72 DPS, sim-verified) [dungeon]; Corpsemaker (6687, -1.92 DPS) [dungeon]; Cobalt Crusher (7730, -4.07 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Darkspear Raider's Cloak; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Gloves of Kapelan; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Snake Hoop; finger2: Thunderbrow Ring; main_hand: Viscous Hammer

No-known-source sample (15 of 469, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 0000000000000000-055030310051002051-0000000000000000)

Set DPS (verified): 144.0. Weights run: 4.8s. Verify run: 15.0s. 787 eligible items had no known source.

Race sweep: windshaper-skyborne 145.3 ± 0.2, troll 144.4 ± 0.2, orc 144.0 ± 0.2, tauren 143.0 ± 0.2 (dps); windshaper-skyborne beats the guide race orc by 1.3 (0.90%) ± 0.3, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.350 ± 0.055, crit=1.593 ± 0.068 per rating point (14 rating = 1%, 22.301 per %), hit=2.865 ± 0.140 per rating point (10 rating = 1%, 28.647 per %), melee_haste=9.530 ± 1.049, intellect=3.205 ± 0.042, spell_power=1.026 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 74.8 attack_power points (5.59 DPS) | yes | Miner's Hat of the Deep (9429, -0.60 DPS) [dungeon]; Papal Fez (9431, -0.91 DPS) [dungeon]; Thinking Cap (2624, -1.52 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.8 attack_power points (2.45 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.12 DPS) [world]; Kaleidoscope Chain (13084, -0.49 DPS) [world_drop]; Darkspear Warding Pendant (272074, -0.78 DPS) [vendor] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 51.3 attack_power points (3.83 DPS) | yes | Inquisitor's Shawl (19507, -0.18 DPS) [dungeon]; Windchaser Amice (14432, -0.42 DPS) [world_drop]; Green Silken Shoulders (7057, -0.58 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glyphed Cloak (4732, -0.96 DPS) [world_drop]; Long Silken Cloak (4326, -0.98 DPS) [crafted]; Blackforge Cape (6424, -1.21 DPS, sim-verified) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 60.2 attack_power points (4.50 DPS) | yes | Red Mageweave Vest (10007, -0.19 DPS) [crafted]; Big Voodoo Robe (8200, -0.46 DPS) [crafted]; Robe of Power (7054, -0.55 DPS) [crafted] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 31.6 attack_power points (2.37 DPS) | yes | Windchaser Cuffs (14429, -0.01 DPS) [world_drop]; Radiant Silver Bracers (4545, -0.14 DPS) [quest]; Warden's Wristbands (14600, -0.15 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormcloth Gloves (10011, +0.00 DPS) [crafted]; Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Rustler Gloves (9704, -3.00 DPS, sim-verified) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 55.3 attack_power points (4.13 DPS) | yes | Defiler's Lizardhide Girdle (20173, -0.54 DPS) [rep]; Defiler's Mail Girdle (20198, -0.54 DPS) [rep]; Saltscale Girdle (273023, -0.63 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 64.5 attack_power points (4.82 DPS) | yes | Red Mageweave Pants (10009, -0.87 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.09 DPS) [crafted]; Totemic Leather Leggings (252458, -1.42 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 47.5 attack_power points (3.55 DPS) | yes | Skirmisher's Mail Boots (252564, -0.02 DPS) [crafted]; Stalker's Mail Boots (252562, -0.76 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.87 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 28.6 attack_power points (2.14 DPS) | yes | Philanthropist's Ring (281635, -0.17 DPS) [quest]; Ogremind Ring (1993, -0.46 DPS) [world_drop]; Voodoo Band (1996, -0.46 DPS) [world] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 26.8 attack_power points (2.00 DPS) | yes | Philanthropist's Ring (281635, -0.04 DPS) [quest]; Ogremind Ring (1993, -0.33 DPS) [world_drop]; Voodoo Band (1996, -0.33 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, -0.68 DPS) [world_drop]; Thornstone Sledgehammer (1722, -2.63 DPS) [world_drop]; Mograine's Might (7723, -8.34 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger2: Legionnaire's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 787, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 1000000000000000-055030310051202251-0500000000000000)

Set DPS (verified): 222.8. Weights run: 4.9s. Verify run: 15.5s. 1018 eligible items had no known source.

Race sweep: windshaper-skyborne 223.9 ± 0.3, orc 222.8 ± 0.3, troll 222.1 ± 0.3, tauren 220.9 ± 0.3 (dps); windshaper-skyborne beats the guide race orc by 1.1 (0.50%) ± 0.4, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.216 ± 0.052, crit=1.609 ± 0.069 per rating point (14 rating = 1%, 22.525 per %), hit=2.642 ± 0.143 per rating point (10 rating = 1%, 26.423 per %), melee_haste=9.607 ± 0.998, intellect=4.234 ± 0.071, spell_power=0.762 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 128.0 attack_power points (10.44 DPS) | yes | Braincage (12549, -1.38 DPS) [dungeon]; Soulcatcher Halo (10630, -1.44 DPS) [dungeon]; Eye of Theradras (17715, -2.04 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 59.3 attack_power points (4.83 DPS) | yes | Gemshard Heart (17707, -1.41 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -1.64 DPS) [dungeon]; Darkspear Warding Pendant (272073, -1.73 DPS) [vendor] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 96.1 attack_power points (7.84 DPS) | yes | Dregmetal Spaulders (11722, -0.66 DPS) [dungeon]; Rotgrip Mantle (17732, -0.81 DPS) [dungeon]; Splinthide Shoulders (11685, -2.03 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Red Cloak (8248, -1.04 DPS) [world_drop]; Deep Woodlands Cloak (19121, -1.37 DPS, sim-verified) [quest]; Keeper's Cloak (14665, -1.38 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 110.0 attack_power points (8.97 DPS) | yes | Stone Guard's Mail Armor (220826, -1.03 DPS) [vendor]; Stone Guard's Pulsing Breastplate (220844, -1.55 DPS) [vendor]; Hibernal Robe (8113, -2.06 DPS) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 66.6 attack_power points (5.43 DPS) | yes | Shizzle's Nozzle Wiper (11917, -1.28 DPS) [quest]; Runic Leather Bracers (15092, -1.35 DPS) [crafted]; Forgotten Wraps (9433, -1.47 DPS, sim-verified) [world_drop] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rockgrip Gauntlets (17736, -0.16 DPS) [dungeon]; Prowler's Leather Gauntlets (252547, -0.99 DPS) [crafted]; Skirmisher's Mail Gauntlets (252586, -1.20 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 85.0 attack_power points (6.93 DPS) | yes | Serenity Belt (13144, -1.06 DPS) [world_drop]; Imperial Red Sash (8253, -1.10 DPS) [world_drop]; Ban'thok Sash (11662, -1.52 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Mail Legplates (220834) | PvP rank 6 · Stone Guard · Horde [vendor] | 97.3 attack_power points (7.94 DPS) | yes | Infernal Trickster Leggings (17754, -1.04 DPS) [dungeon]; Imperial Red Pants (8251, -1.27 DPS) [world_drop]; Stone Guard's Pulsing Legplates (220847, -2.39 DPS, sim-verified) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 77.1 attack_power points (6.29 DPS) | yes | Coldstone Slippers (18697, -1.46 DPS) [dungeon]; Skycaller's Leather Boots (252471, -1.56 DPS) [crafted]; Bloodshot Greaves (10846, -2.70 DPS, sim-verified) [world] |
| finger1 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 64.0 attack_power points (5.22 DPS) | yes | Woodseed Hoop (17768, -1.40 DPS) [quest]; Blackstone Ring (17713, -1.44 DPS) [dungeon]; Cyclopean Band (11824, -1.59 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 50.7 attack_power points (4.14 DPS) | yes | Woodseed Hoop (17768, -0.31 DPS) [quest]; Blackstone Ring (17713, -0.35 DPS) [dungeon]; Cyclopean Band (11824, -0.51 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.57 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.28 DPS) [world_drop]; Molten Heart of the Mountain (249470, -2.37 DPS, sim-verified) [crafted] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Taran Icebreaker (2915, -0.19 DPS) [world_drop]; Glowing Brightwood Staff (812, -0.40 DPS) [world_drop]; Thorium Greatmace (250613, -1.08 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Mail Legplates; feet: Greaves of Withering Despair; finger1: Mindseye Circle; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Soulkeeper

No-known-source sample (15 of 1018, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 5030000000000000-055030310051202251-0530000000000000)

Set DPS (verified): 334.1. Weights run: 4.6s. Verify run: 52.3s. 2331 eligible items had no known source.

Race sweep: windshaper-skyborne 336.0 ± 0.4, troll 334.3 ± 0.4, orc 334.1 ± 0.4, tauren 331.9 ± 0.4 (dps); windshaper-skyborne beats the guide race orc by 1.8 (0.55%) ± 0.6, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.695 ± 0.081, crit=2.830 ± 0.122 per rating point (14 rating = 1%, 39.620 per %), hit=4.492 ± 0.250 per rating point (10 rating = 1%, 44.918 per %), melee_haste=14.989 ± 1.788, intellect=3.963 ± 0.104, spell_power=0.868 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 225.8 attack_power points (18.21 DPS) | yes | Cap of the Scarlet Savant (12752, -4.23 DPS, sim-verified) [quest]; Warlord's Mail Headguard (231657, -4.45 DPS) [vendor]; Skyfury Helm (20134, -6.68 DPS) [quest] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Nacreous Shell Necklace (22403, -0.03 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.67 DPS) [world_drop]; Star of Mystaria (12103, -1.57 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Mail Pauldrons (227154, +0.00 DPS) [pvp]; Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Soulstealer Mantle (13374, -1.95 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -0.16 DPS) [dungeon]; Howler's Furs (272414, -1.42 DPS) [vendor]; Royal Tribunal Cloak (13376, -1.84 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Mail Armor (16577, +0.00 DPS) [vendor]; Warlord's Mail Hauberk (231653, +0.00 DPS) [vendor]; Dawn Armor (252483, -2.71 DPS, sim-verified) [crafted] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | 69.7 attack_power points (5.62 DPS) | yes | Manacle Cuffs (11962, -0.34 DPS) [quest]; Aristocratic Cuffs (12546, -0.55 DPS) [dungeon]; Magiskull Cuffs (13107, -2.21 DPS, sim-verified) [world_drop] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; General's Mail Vices (231655, +0.00 DPS) [vendor]; Storm Gauntlets (12632, -4.09 DPS, sim-verified) [crafted] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Defiler's Lizardhide Girdle (20171, -0.16 DPS) [rep]; Clutch of Andros (13956, -1.95 DPS, sim-verified) [dungeon] |
| legs | Outrider's Lizardhide Pants (22741) | Warsong Outriders [rep] | 157.7 attack_power points (12.72 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; General's Mail Leggings (231664, +0.00 DPS) [pvp]; Outrider's Mail Leggings (22676, -0.23 DPS) [rep] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; Kayser's Boots of Precision (22231, -0.69 DPS, sim-verified) [dungeon]; Windreaver Greaves (13967, -1.73 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -2.10 DPS) [dungeon]; Mindseye Circle (10634, -2.45 DPS) [dungeon]; Naglering (11669, -6.26 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -1.95 DPS) [dungeon]; Mindseye Circle (10634, -2.30 DPS) [dungeon]; Cutthroat's Signet (272408, -2.41 DPS) [vendor] |
| trinket1 | Heart of Wyrmthalak (22321) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -8.30 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Rage (22395) | Blackrock Depths: Magmus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, +0.00 DPS) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Darkspear Shoulderguards; back: Arcanoweave Cloak; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Raider Handwraps; waist: Stormseeker's Girdle; legs: Outrider's Lizardhide Pants; feet: Bloodmail Boots; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Heart of Wyrmthalak; trinket2: Burst of Knowledge; main_hand: The Unstoppable Force; ranged: Totem of Rage

No-known-source sample (15 of 2331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 0533001000000000-053030310051202051-0530000000000000)

Set DPS (verified): 856.8. Weights run: 4.8s. Verify run: 52.9s. 2331 eligible items had no known source.

Race sweep: windshaper-skyborne 862.8 ± 1.2, orc 856.8 ± 1.2, troll 855.7 ± 1.2, tauren 848.0 ± 1.2 (dps); windshaper-skyborne beats the guide race orc by 6.0 (0.70%) ± 1.7, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=1.339 ± 0.052, crit=2.236 ± 0.073 per rating point (14 rating = 1%, 31.306 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=19.365 ± 2.160, intellect=1.583 ± 0.012, spell_power=0.806 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 119.6 attack_power points (26.24 DPS) | yes | Skyfury Helm (20134, -3.08 DPS) [quest]; Warlord's Mail Headguard (231657, -3.58 DPS) [vendor]; Cap of the Scarlet Savant (12752, -5.55 DPS) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 57.3 attack_power points (12.57 DPS) | yes | Lady Maye's Pendant (14558, -4.03 DPS) [world_drop]; Emberfury Talisman (12929, -4.64 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -5.26 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 67.5 attack_power points (14.80 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, +0.00 DPS) [rep]; Blackveil Cape (11626, -1.79 DPS) [dungeon]; Royal Tribunal Cloak (13376, -2.35 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 128.0 attack_power points (28.08 DPS) | yes | Savage Gladiator Chain (11726, -5.95 DPS) [dungeon]; Warlord's Mail Hauberk (231653, -6.08 DPS) [vendor]; Breastplate of Bloodthirst (12757, -8.07 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windtalker's Wristguards (19582, +0.00 DPS) [rep] |
| hands | Bloodmail Gauntlets (14615) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Vices (231655, +0.00 DPS) [vendor]; General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, -1.10 DPS) [vendor] |
| waist | Defiler's Lizardhide Girdle (20171) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Chain Girdle (20150, +0.00 DPS) [rep]; Defiler's Leather Girdle (20190, +0.00 DPS) [rep]; Windseeker's Belt (272410, +0.00 DPS) [vendor] |
| legs | Outrider's Lizardhide Pants (22741) | Warsong Outriders [rep] | 105.5 attack_power points (23.14 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; Pristine Scorpid Leggings (252606, -1.75 DPS) [crafted]; Outrider's Mail Leggings (22676, -1.88 DPS) [rep] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -0.66 DPS) [crafted] |
| finger1 | Painweaver Band (13098) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -0.48 DPS) [rep]; Don Mauricio's Band of Domination (22433, -0.69 DPS) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -0.92 DPS) [rep]; Don Mauricio's Band of Domination (22433, -1.13 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Rage (22395) | Blackrock Depths: Magmus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Bloodmail Gauntlets; waist: Defiler's Lizardhide Girdle; legs: Outrider's Lizardhide Pants; feet: Bloodmail Boots; finger1: Painweaver Band; finger2: Don Julio's Band; trinket1: Blackhand's Breadth; trinket2: Smolderweb's Eye; main_hand: The Unstoppable Force; ranged: Totem of Rage

No-known-source sample (15 of 2331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

