# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-053030000000000000-0000000000000000)

Set DPS (verified): 39.0. Weights run: 6.8s. Verify run: 14.0s. 286 eligible items had no known source.

Race sweep: dwarf 39.0 ± 0.0 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.902 ± 0.021, crit=0.760 ± 0.021 per rating point (14 rating = 1%, 10.638 per %), hit=1.396 ± 0.060 per rating point (10 rating = 1%, 13.956 per %), melee_haste=6.350 ± 0.202, intellect=2.524 ± 0.024, spell_power=1.655 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Trapper's Leather Hood (252505) | Leatherworking [crafted] | 20.2 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS) [crafted]; Totemic Leather Hood (252448, -0.18 DPS) [crafted]; Shadow Goggles (4373, -0.26 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 5.4 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS) [quest]; Tarnished Locket (279870, -0.40 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 31.0 attack_power points (1.08 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.44 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.85 DPS) [crafted]; Serpent's Shoulders (5404, -0.92 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (39.0 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS) [crafted]; Lambent Scale Cloak (4706, -0.07 DPS) [world_drop]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-verified (39.0 DPS) | yes | Totemic Leather Armor (252435, +0.00 DPS) [crafted]; Stormrider's Leather Armor (252492, +0.00 DPS) [crafted]; Wisdom's Leather Armor (252493, +0.00 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.09 DPS) [world_drop]; Bravo's Armbands (270015, -0.14 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (39.0 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Totemic Leather Gloves (252497, +0.00 DPS) [crafted]; Stormrider's Leather Gloves (252498, +0.00 DPS) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 20.2 attack_power points (0.71 DPS) | yes | Blackened Defias Belt (10403, -0.08 DPS) [dungeon]; Stormrider's Leather Belt (252432, -0.12 DPS) [crafted]; Pristine Sash (253925, -0.12 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 39.7 attack_power points (1.39 DPS) | yes | Stormrider's Leather Pants (252502, -0.28 DPS) [crafted]; Trapper's Leather Pants (252501, -0.33 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.45 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 22.6 attack_power points (0.79 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Totemic Leather Boots (252442, -0.09 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.12 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.9 attack_power points (0.59 DPS) | yes | Minor Channeling Ring (1449, -0.13 DPS) [quest]; Lorekeeper's Ring (20431, -0.30 DPS) [rep]; Demon Band (12054, -0.31 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | sim-verified (39.0 DPS) | yes | Minor Channeling Ring (1449, +0.00 DPS) [quest]; Lorekeeper's Ring (20431, -0.12 DPS) [rep]; Demon Band (12054, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 311.4 attack_power points (10.87 DPS) | yes | Smite's Mighty Hammer (7230, -0.34 DPS) [dungeon]; Rockslicer (872, -0.76 DPS) [dungeon]; Night Reaver (1318, -1.25 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Trapper's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Armor of the Fang; wrist: Mindthrust Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Totemic Leather Pants; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish

### Band 30 (dwarf, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 85.8. Weights run: 6.7s. Verify run: 18.8s. 475 eligible items had no known source.

Race sweep: dwarf 85.8 ± 0.1 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.148 ± 0.029, crit=1.065 ± 0.034 per rating point (14 rating = 1%, 14.907 per %), hit=1.769 ± 0.072 per rating point (10 rating = 1%, 17.686 per %), melee_haste=6.100 ± 0.481, intellect=2.450 ± 0.019, spell_power=0.777 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 39.2 attack_power points (2.79 DPS) | yes | Dokebi Hat (14584, -0.30 DPS) [world_drop]; Nightsky Cowl (4039, -0.70 DPS) [world_drop]; Resilient Cap (14401, -0.71 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 26.7 attack_power points (1.90 DPS) | yes | Kaleidoscope Chain (13084, -0.31 DPS) [world_drop]; Seal of the Expedition (271920, -1.03 DPS) [rep]; Crystal Starfire Medallion (5003, -1.20 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.6 attack_power points (2.25 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Magician's Mantle (12998, -0.40 DPS) [world_drop]; Feline Mantle (3748, -0.47 DPS) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (85.8 DPS) | yes | Expeditionary's Cape (271923, -0.20 DPS) [rep]; Resilient Cape (14400, -0.55 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.59 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 47.1 attack_power points (3.35 DPS) | yes | Spirewind Fetter (9406, -0.38 DPS, sim-verified) [dungeon]; Green Silk Armor (7065, -0.59 DPS) [crafted]; Beguiler Robes (7728, -0.76 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 25.6 attack_power points (1.82 DPS) | yes | Technician's Bracers (270042, -0.55 DPS) [quest]; Hawkeye's Bracers (14590, -0.76 DPS) [world_drop]; Nightsky Wristbands (6407, -0.77 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 24.2 attack_power points (1.72 DPS) | yes | Insignia Gloves (6408, -0.23 DPS) [world_drop]; Town Clerk's Mittens (270029, -0.33 DPS) [quest]; Toughened Leather Gloves (4253, -0.38 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 34.7 attack_power points (2.47 DPS) | yes | Skirmisher's Leather Belt (252461, -0.10 DPS) [crafted]; Highlander's Lizardhide Girdle (20105, -0.38 DPS) [rep]; Highlander's Mail Girdle (20120, -0.38 DPS) [vendor] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 55.4 attack_power points (3.94 DPS) | yes | Totemic Leather Leggings (252458, -1.13 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -1.58 DPS) [crafted]; Trapper's Leather Pants (252501, -1.65 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 23.5 attack_power points (1.67 DPS) | yes | Gilded Slippers (254001, -0.06 DPS) [crafted]; Frothing Slippers (254003, -0.45 DPS) [crafted]; Fiery Slippers (254005, -0.45 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.4 attack_power points (1.38 DPS) | yes | Protector's Band (19517, -0.04 DPS) [rep]; Runemender's Seal (273031, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 19.2 attack_power points (1.37 DPS) | yes | Protector's Band (19517, -0.03 DPS) [rep]; Runemender's Seal (273031, -0.16 DPS) [world]; Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -1.33 DPS) [dungeon]; Corpsemaker (6687, -1.93 DPS) [dungeon]; Cobalt Crusher (7730, -4.10 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Sergeant Major's Cape; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Thunderbrow Ring; finger2: Seal of Wrynn; main_hand: Viscous Hammer

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 0000000000000000-055030310051002051-0000000000000000)

Set DPS (verified): 146.8. Weights run: 7.1s. Verify run: 20.9s. 804 eligible items had no known source.

Race sweep: dwarf 146.8 ± 0.2 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.453 ± 0.058, crit=1.677 ± 0.073 per rating point (14 rating = 1%, 23.476 per %), hit=3.110 ± 0.150 per rating point (10 rating = 1%, 31.097 per %), melee_haste=7.480 ± 1.044, intellect=3.256 ± 0.044, spell_power=1.110 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 76.9 attack_power points (5.74 DPS) | yes | Miner's Hat of the Deep (9429, -0.61 DPS) [dungeon]; Papal Fez (9431, -0.94 DPS) [dungeon]; Big Voodoo Mask (8201, -1.59 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.8 attack_power points (2.52 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.17 DPS) [world]; Kaleidoscope Chain (13084, -0.52 DPS) [world_drop]; Seal of the Expedition (271919, -0.82 DPS) [rep] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 52.1 attack_power points (3.89 DPS) | yes | Inquisitor's Shawl (19507, -0.15 DPS) [dungeon]; Windchaser Amice (14432, -0.40 DPS) [world_drop]; Green Silken Shoulders (7057, -0.55 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.33 DPS) [rep]; Blackforge Cape (6424, -1.06 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 66.0 attack_power points (4.92 DPS) | yes | Garb of Florid Feathers (273026, -0.30 DPS) [dungeon]; Red Mageweave Vest (10007, -0.55 DPS) [crafted]; Big Voodoo Robe (8200, -0.78 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 32.2 attack_power points (2.40 DPS) | yes | Glowing Magical Bracelets (13106, -0.01 DPS) [world_drop]; Warden's Wristbands (14600, -0.13 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.19 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Stormcloth Gloves (10011, +0.00 DPS) [crafted]; Red Mageweave Gloves (10018, -2.35 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 56.6 attack_power points (4.22 DPS) | yes | Highlander's Lizardhide Girdle (20104, -0.58 DPS) [rep]; Highlander's Mail Girdle (20119, -0.58 DPS) [vendor]; Saltscale Girdle (273023, -0.69 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 65.1 attack_power points (4.85 DPS) | yes | Red Mageweave Pants (10009, -0.78 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.03 DPS) [crafted]; Totemic Leather Leggings (252458, -1.39 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 48.5 attack_power points (3.61 DPS) | yes | Skirmisher's Mail Boots (252564, -0.01 DPS) [crafted]; Stalker's Mail Boots (252562, -0.72 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.84 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 27.6 attack_power points (2.06 DPS) | yes | Seal of Wrynn (2933, -0.32 DPS) [quest]; Snake Hoop (6750, -0.35 DPS) [quest]; Runemender's Seal (273031, -0.35 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 27.4 attack_power points (2.04 DPS) | yes | Seal of Wrynn (2933, -0.30 DPS) [quest]; Snake Hoop (6750, -0.33 DPS) [quest]; Runemender's Seal (273031, -0.34 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Pendulum of Doom (9425, -6.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Sheepshear Mantle; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 1000000000000000-055030310051202251-0500000000000000)

Set DPS (verified): 226.2. Weights run: 7.3s. Verify run: 17.5s. 1053 eligible items had no known source.

Race sweep: dwarf 226.2 ± 0.3 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.514 ± 0.071, crit=2.289 ± 0.099 per rating point (14 rating = 1%, 32.052 per %), hit=3.710 ± 0.151 per rating point (10 rating = 1%, 37.103 per %), melee_haste=10.962 ± 1.400, intellect=3.945 ± 0.050, spell_power=1.311 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 131.8 attack_power points (11.16 DPS) | yes | Braincage (12549, -1.34 DPS) [dungeon]; Eye of Theradras (17715, -2.10 DPS) [dungeon]; Soulcatcher Halo (10630, -2.14 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 55.2 attack_power points (4.68 DPS) | yes | Gemshard Heart (17707, -1.12 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.23 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.66 DPS) [world] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 98.6 attack_power points (8.35 DPS) | yes | Rotgrip Mantle (17732, -0.89 DPS) [dungeon]; Dregmetal Spaulders (11722, -1.31 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -2.00 DPS) [dungeon] |
| back | Expeditionary's Cape (271921) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16336, -0.89 DPS) [pvp]; Imperial Red Cloak (8248, -1.00 DPS) [world_drop]; Runecloth Cloak (13860, -1.01 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 109.1 attack_power points (9.24 DPS) | yes | Knight's Mail Armor (223078, -0.48 DPS) [vendor]; Runecloth Robe (13858, -2.34 DPS) [crafted]; Hibernal Robe (8113, -2.56 DPS) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 64.4 attack_power points (5.45 DPS) | yes | Runic Leather Bracers (15092, -1.00 DPS) [crafted]; Forgotten Wraps (9433, -1.06 DPS) [world_drop]; Prowler's Leather Bracers (252539, -1.36 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Rockgrip Gauntlets (17736, +0.00 DPS) [dungeon]; Sergeant Major's Mail Gauntlets (223076, +0.00 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Highlander's Lizardhide Girdle (20103, -0.96 DPS) [rep]; Highlander's Mail Girdle (20118, -0.96 DPS) [vendor] |
| legs | Knight's Mail Legplates (223074) | PvP rank 6 · Knight · Alliance [vendor] | 103.4 attack_power points (8.75 DPS) | yes | Infernal Trickster Leggings (17754, -0.38 DPS) [dungeon]; Imperial Red Pants (8251, -1.82 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -2.41 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 91.0 attack_power points (7.70 DPS) | yes | Bloodshot Greaves (10846, -1.74 DPS) [world]; Skycaller's Leather Boots (252471, -2.36 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -2.36 DPS) [crafted] |
| finger1 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 65.9 attack_power points (5.58 DPS) | yes | Blackstone Ring (17713, -0.74 DPS) [dungeon]; Cyclopean Band (11824, -1.56 DPS) [dungeon]; Woodseed Hoop (17768, -1.57 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Cyclopean Band (11824, -0.55 DPS) [dungeon]; Woodseed Hoop (17768, -0.56 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.16 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Expeditionary's Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; waist: Dawnspire Cord; legs: Knight's Mail Legplates; feet: Greaves of Withering Despair; finger1: Mindseye Circle; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Soulkeeper

No-known-source sample (15 of 1053, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 5030000000000000-055030310051202251-0530000000000000)

Set DPS (verified): 340.8. Weights run: 6.8s. Verify run: 77.8s. 2377 eligible items had no known source.

Race sweep: dwarf 340.8 ± 0.4 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.782 ± 0.082, crit=3.019 ± 0.123 per rating point (14 rating = 1%, 42.269 per %), hit=4.528 ± 0.253 per rating point (10 rating = 1%, 45.282 per %), melee_haste=14.169 ± 1.792, intellect=4.498 ± 0.107, spell_power=0.928 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 250.6 attack_power points (20.22 DPS) | yes | Cap of the Scarlet Savant (12752, -4.18 DPS, sim-verified) [quest]; Chief Architect's Monocle (11839, -7.86 DPS) [dungeon]; Blue Dragonscale Helm (252604, -8.16 DPS) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 95.7 attack_power points (7.72 DPS) | yes | Nacreous Shell Necklace (22403, -0.49 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.74 DPS) [world_drop]; Star of Mystaria (12103, -3.86 DPS, sim-verified) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Shoulderpads (271953, -0.30 DPS) [rep]; Theramore Shoulders (271954, -0.30 DPS) [rep]; Soulstealer Mantle (13374, -4.65 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Tribunal Cloak (13376, -1.57 DPS) [dungeon]; Howler's Furs (272414, -1.84 DPS) [vendor]; Expeditionary's Cape (271907, -1.95 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Robes (16688, -0.26 DPS) [dungeon]; Blue Dragonscale Breastplate (15048, -0.56 DPS) [crafted]; Dawn Armor (252483, -4.08 DPS, sim-verified) [crafted] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manacle Cuffs (11962, -0.00 DPS) [quest]; Aristocratic Cuffs (12546, -0.23 DPS) [dungeon]; Magiskull Cuffs (13107, -2.43 DPS, sim-verified) [world_drop] |
| hands | Bloodmail Gauntlets (14615) | Scholomance [dungeon] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Theramore Handwraps (271946, +0.00 DPS) [rep] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Mail Girdle (20044, -0.33 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, -0.33 DPS) [rep]; Clutch of Andros (13956, -4.65 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 192.2 attack_power points (15.51 DPS) | yes | Sentinel's Lizardhide Pants (22750, -1.71 DPS) [rep]; Red Dragonscale Leggings (252603, -2.40 DPS) [crafted]; Sentinel's Silk Leggings (237815, -2.66 DPS) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greaves of Withering Despair (22240, -1.78 DPS) [dungeon]; Windreaver Greaves (13967, -1.90 DPS) [dungeon]; Kayser's Boots of Precision (22231, -2.21 DPS, sim-verified) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -1.97 DPS) [dungeon]; Mindseye Circle (10634, -2.55 DPS) [dungeon]; Naglering (11669, -7.42 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -1.43 DPS) [dungeon]; Mindseye Circle (10634, -2.01 DPS) [dungeon]; Cutthroat's Signet (272408, -2.53 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Ankh of Life (1713, -2.06 DPS, sim-verified) [world_drop] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -9.78 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Rage (22395) | Blackrock Depths: Magmus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -2.03 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Theramore Shoulderguards; back: Arcanoweave Cloak; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Bloodmail Gauntlets; waist: Stormseeker's Girdle; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: The Unstoppable Force; ranged: Totem of Rage

No-known-source sample (15 of 2377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 0535001000000000-054030310051202051-0500000000000000)

Set DPS (verified): 890.2. Weights run: 6.9s. Verify run: 78.9s. 2377 eligible items had no known source.

Race sweep: dwarf 890.2 ± 1.3 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=1.440 ± 0.057, crit=2.439 ± 0.086 per rating point (14 rating = 1%, 34.143 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=24.720 ± 2.152, intellect=1.691 ± 0.022, spell_power=0.915 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 128.6 attack_power points (27.63 DPS) | yes | Cap of the Scarlet Savant (12752, -5.70 DPS) [quest]; Eye of Rend (12587, -9.54 DPS) [dungeon]; Ragefury Eyepatch (11735, -10.13 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 60.1 attack_power points (12.92 DPS) | yes | Lady Maye's Pendant (14558, -3.86 DPS) [world_drop]; Emberfury Talisman (12929, -4.41 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -5.20 DPS) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 73.0 attack_power points (15.69 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Mantle of the Timbermaw (19050, -0.29 DPS) [crafted]; Shroud of the Nathrezim (18720, -1.16 DPS) [dungeon] |
| back | Sergeant Major's Cape (16337) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.23 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.31 DPS) [rep]; Blackveil Cape (11626, -2.00 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 132.7 attack_power points (28.50 DPS) | yes | Savage Gladiator Chain (11726, -5.60 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -7.68 DPS) [quest]; Tombstone Breastplate (13944, -8.16 DPS) [dungeon] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Prowler's Leather Bracers (252539, -2.83 DPS) [crafted] |
| hands | Bloodmail Gauntlets (14615) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Brawlers (19049, -1.91 DPS) [crafted]; Stonerender Gauntlets (17007, -2.40 DPS) [world_drop]; Theramore Gloves (271948, -2.62 DPS) [rep] |
| waist | Windseeker's Belt (272410) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Chain Girdle (20043, +0.00 DPS) [rep]; Highlander's Leather Girdle (20045, +0.00 DPS) [rep]; Might of the Timbermaw (19044, -0.23 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 118.7 attack_power points (25.50 DPS) | yes | Sentinel's Silk Leggings (237815, -0.90 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -0.90 DPS) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scalegut Treaders (275618, +0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -0.81 DPS) [crafted]; Dragonrider Boots (18102, -1.01 DPS) [dungeon] |
| finger1 | Magni's Will (12548) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.73 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -1.15 DPS) [dungeon]; Protector's Band (19514, -1.57 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.43 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -0.85 DPS) [dungeon]; Protector's Band (19514, -1.27 DPS) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -1.20 DPS) [vendor]; Uther's Strength (11302, -1.40 DPS) [world_drop] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Rage (22395) | Blackrock Depths: Magmus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Sergeant Major's Cape; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Bloodmail Gauntlets; waist: Windseeker's Belt; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Magni's Will; finger2: Don Julio's Band; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Totem of Rage

No-known-source sample (15 of 2377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 0000000000000000-053030000000000000-0000000000000000)

Set DPS (verified): 39.4. Weights run: 6.8s. Verify run: 14.4s. 275 eligible items had no known source.

Race sweep: windshaper-skyborne 39.7 ± 0.0, troll 39.4 ± 0.0, orc 39.4 ± 0.0, tauren 39.0 ± 0.0 (dps); windshaper-skyborne beats the guide race orc by 0.3 (0.89%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.902 ± 0.021, crit=0.760 ± 0.021 per rating point (14 rating = 1%, 10.638 per %), hit=1.396 ± 0.060 per rating point (10 rating = 1%, 13.956 per %), melee_haste=6.350 ± 0.202, intellect=2.524 ± 0.024, spell_power=1.655 ± 0.005

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Trapper's Leather Hood (252505) | Leatherworking [crafted] | 20.2 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS) [crafted]; Totemic Leather Hood (252448, -0.18 DPS) [crafted]; Shadow Goggles (4373, -0.26 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 5.4 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 31.0 attack_power points (1.08 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.44 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.85 DPS) [crafted]; Serpent's Shoulders (5404, -0.92 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanguine Cape (14376, -0.03 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Dark Leather Cloak (2316, -0.15 DPS) [crafted] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totemic Leather Armor (252435, +0.00 DPS) [crafted]; Stormrider's Leather Armor (252492, +0.00 DPS) [crafted]; Wisdom's Leather Armor (252493, +0.00 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.09 DPS) [world_drop]; Featherbead Bracers (15452, -0.18 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Totemic Leather Gloves (252497, +0.00 DPS) [crafted]; Stormrider's Leather Gloves (252498, +0.00 DPS) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 20.2 attack_power points (0.71 DPS) | yes | Blackened Defias Belt (10403, -0.08 DPS) [dungeon]; Stormrider's Leather Belt (252432, -0.12 DPS) [crafted]; Pristine Sash (253925, -0.12 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 39.7 attack_power points (1.39 DPS) | yes | Stormrider's Leather Pants (252502, -0.28 DPS) [crafted]; Trapper's Leather Pants (252501, -0.33 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.45 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 22.6 attack_power points (0.79 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Totemic Leather Boots (252442, -0.09 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.12 DPS) [crafted] |
| finger1 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 18.6 attack_power points (0.65 DPS) | yes | Legionnaire's Band (20429, -0.24 DPS) [rep]; Advisor's Ring (20426, -0.36 DPS) [rep]; Demon Band (12054, -0.37 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.9 attack_power points (0.59 DPS) | yes | Legionnaire's Band (20429, -0.19 DPS) [rep]; Advisor's Ring (20426, -0.30 DPS) [rep]; Demon Band (12054, -0.31 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 311.4 attack_power points (10.87 DPS) | yes | Forsaken Greataxe (251533, -0.29 DPS) [quest]; Smite's Mighty Hammer (7230, -0.34 DPS) [dungeon]; Hammerbone (270018, -0.39 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Trapper's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Armor of the Fang; wrist: Mindthrust Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Totemic Leather Pants; feet: Stormrider's Leather Boots; finger1: Loop of Sacrifice; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 275, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield

### Band 30 (orc, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 85.9. Weights run: 6.7s. Verify run: 21.3s. 469 eligible items had no known source.

Race sweep: windshaper-skyborne 86.9 ± 0.1, troll 86.0 ± 0.1, orc 85.9 ± 0.1, tauren 85.2 ± 0.1 (dps); windshaper-skyborne beats the guide race orc by 1.0 (1.12%) ± 0.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.148 ± 0.029, crit=1.065 ± 0.034 per rating point (14 rating = 1%, 14.907 per %), hit=1.769 ± 0.072 per rating point (10 rating = 1%, 17.686 per %), melee_haste=6.100 ± 0.481, intellect=2.450 ± 0.019, spell_power=0.777 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 39.2 attack_power points (2.79 DPS) | yes | Dokebi Hat (14584, -0.30 DPS) [world_drop]; Nightsky Cowl (4039, -0.70 DPS) [world_drop]; Resilient Cap (14401, -0.71 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 26.7 attack_power points (1.90 DPS) | yes | Kaleidoscope Chain (13084, -0.22 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -1.03 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.20 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 31.6 attack_power points (2.25 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Magician's Mantle (12998, -0.40 DPS) [world_drop]; Feline Mantle (3748, -0.47 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resilient Cape (14400, -0.35 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.40 DPS) [world_drop]; Soft Willow Cape (16661, -0.52 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 47.1 attack_power points (3.35 DPS) | yes | Spirewind Fetter (9406, -0.26 DPS) [dungeon]; Green Silk Armor (7065, -0.59 DPS) [crafted]; Beguiler Robes (7728, -0.76 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 25.6 attack_power points (1.82 DPS) | yes | Technician's Bracers (270042, -0.55 DPS) [quest]; Hawkeye's Bracers (14590, -0.76 DPS) [world_drop]; Nightsky Wristbands (6407, -0.77 DPS) [world_drop] |
| hands | Gloves of Kapelan (6744) | Alliance Relations [quest] | 28.7 attack_power points (2.04 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.32 DPS) [dungeon]; Insignia Gloves (6408, -0.56 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.70 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 34.7 attack_power points (2.47 DPS) | yes | Skirmisher's Leather Belt (252461, -0.22 DPS, sim-verified) [crafted]; Defiler's Lizardhide Girdle (20172, -0.38 DPS) [rep]; Defiler's Mail Girdle (20197, -0.38 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 55.4 attack_power points (3.94 DPS) | yes | Totemic Leather Leggings (252458, -0.84 DPS, sim-verified) [crafted]; Totemic Leather Pants (252446, -1.58 DPS) [crafted]; Trapper's Leather Pants (252501, -1.65 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 23.5 attack_power points (1.67 DPS) | yes | Gilded Slippers (254001, -0.06 DPS) [crafted]; Frothing Slippers (254003, -0.45 DPS) [crafted]; Fiery Slippers (254005, -0.45 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.4 attack_power points (1.38 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; Runemender's Seal (273031, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.9 attack_power points (1.34 DPS) | yes | Loop of Sacrifice (281673, -0.05 DPS) [quest]; Runemender's Seal (273031, -0.13 DPS) [world]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -1.33 DPS) [dungeon]; Corpsemaker (6687, -1.93 DPS) [dungeon]; Cobalt Crusher (7730, -4.10 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Darkspear Raider's Cloak; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Gloves of Kapelan; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Viscous Hammer

No-known-source sample (15 of 469, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 0000000000000000-055030310051002051-0000000000000000)

Set DPS (verified): 145.9. Weights run: 7.1s. Verify run: 22.7s. 787 eligible items had no known source.

Race sweep: windshaper-skyborne 147.8 ± 0.2, troll 146.4 ± 0.2, orc 145.9 ± 0.2, tauren 144.7 ± 0.2 (dps); windshaper-skyborne beats the guide race orc by 1.8 (1.25%) ± 0.3, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.421 ± 0.049, crit=1.563 ± 0.060 per rating point (14 rating = 1%, 21.880 per %), hit=2.454 ± 0.115 per rating point (10 rating = 1%, 24.539 per %), melee_haste=8.052 ± 0.790, intellect=3.824 ± 0.058, spell_power=0.788 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 85.3 attack_power points (5.95 DPS) | yes | Miner's Hat of the Deep (9429, -0.75 DPS) [dungeon]; Papal Fez (9431, -0.97 DPS) [dungeon]; Thinking Cap (2624, -1.41 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 36.1 attack_power points (2.52 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Kaleidoscope Chain (13084, -0.50 DPS) [world_drop]; Darkspear Warding Pendant (272074, -0.65 DPS) [vendor] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 61.2 attack_power points (4.27 DPS) | yes | Inquisitor's Shawl (19507, -0.42 DPS) [dungeon]; Windchaser Amice (14432, -0.50 DPS) [world_drop]; Green Silken Shoulders (7057, -0.89 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackforge Cape (6424, -0.80 DPS) [dungeon]; Glyphed Cloak (4732, -1.07 DPS) [world_drop]; Long Silken Cloak (4326, -1.27 DPS) [crafted] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 70.1 attack_power points (4.89 DPS) | yes | Red Mageweave Vest (10007, -0.09 DPS) [crafted]; Silksand Tunic (14417, -0.63 DPS) [world_drop]; Silksand Wraps (14425, -0.63 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 37.3 attack_power points (2.60 DPS) | yes | Glowing Magical Bracelets (13106, -0.05 DPS) [world_drop]; Warden's Wristbands (14600, -0.17 DPS) [world_drop]; Radiant Silver Bracers (4545, -0.25 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Stormcloth Gloves (10011, +0.00 DPS) [crafted]; Rustler Gloves (9704, -1.37 DPS, sim-verified) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 62.9 attack_power points (4.39 DPS) | yes | Defiler's Lizardhide Girdle (20173, -0.38 DPS) [rep]; Defiler's Mail Girdle (20198, -0.38 DPS) [rep]; Saltscale Girdle (273023, -0.73 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 71.9 attack_power points (5.02 DPS) | yes | Red Mageweave Pants (10009, -1.04 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.11 DPS) [crafted]; Misplaced Pantaloons (276201, -1.55 DPS) [vendor] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 51.1 attack_power points (3.56 DPS) | yes | Skirmisher's Mail Boots (252564, -0.20 DPS) [crafted]; Stalker's Mail Boots (252562, -0.60 DPS) [crafted]; Briar Tredders (10582, -0.67 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 27.4 attack_power points (1.91 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Voodoo Band (1996, -0.04 DPS) [world]; Mindbender Loop (5009, -0.04 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 27.0 attack_power points (1.88 DPS) | yes | Ogremind Ring (1993, -0.02 DPS) [world_drop]; Voodoo Band (1996, -0.02 DPS) [world]; Mindbender Loop (5009, -0.02 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Staff of Jordan (873, -1.20 DPS) [world_drop]; Pendulum of Doom (9425, -5.44 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 787, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 1000000000000000-055030310051202251-0500000000000000)

Set DPS (verified): 225.5. Weights run: 7.3s. Verify run: 22.4s. 1021 eligible items had no known source.

Race sweep: windshaper-skyborne 226.7 ± 0.3, orc 225.5 ± 0.3, troll 225.4 ± 0.3, tauren 223.7 ± 0.3 (dps); windshaper-skyborne beats the guide race orc by 1.1 (0.50%) ± 0.4, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.514 ± 0.071, crit=2.289 ± 0.099 per rating point (14 rating = 1%, 32.052 per %), hit=3.710 ± 0.151 per rating point (10 rating = 1%, 37.103 per %), melee_haste=10.962 ± 1.400, intellect=3.945 ± 0.050, spell_power=1.311 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 131.8 attack_power points (11.16 DPS) | yes | Braincage (12549, -1.34 DPS) [dungeon]; Eye of Theradras (17715, -2.10 DPS) [dungeon]; Soulcatcher Halo (10630, -2.14 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 55.2 attack_power points (4.68 DPS) | yes | Gemshard Heart (17707, -1.12 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.23 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.66 DPS) [world] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 98.6 attack_power points (8.35 DPS) | yes | Rotgrip Mantle (17732, -0.89 DPS) [dungeon]; Dregmetal Spaulders (11722, -1.31 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -2.00 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.34 DPS) [quest]; Imperial Red Cloak (8248, -1.00 DPS) [world_drop]; Runecloth Cloak (13860, -1.01 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 109.1 attack_power points (9.24 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.09 DPS) [vendor]; Stone Guard's Mail Armor (220826, -0.48 DPS) [vendor]; Runecloth Robe (13858, -2.34 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 64.4 attack_power points (5.45 DPS) | yes | Runic Leather Bracers (15092, -1.00 DPS) [crafted]; Forgotten Wraps (9433, -1.06 DPS) [world_drop]; Prowler's Leather Bracers (252539, -1.36 DPS) [crafted] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-verified (225.5 DPS) | yes | Rockgrip Gauntlets (17736, -0.68 DPS) [dungeon]; Azurite Fists (20369, -0.90 DPS) [quest]; First Sergeant's Mail Gauntlets (220831, -0.99 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Defiler's Lizardhide Girdle (20174, -0.96 DPS) [rep]; Defiler's Mail Girdle (20196, -0.96 DPS) [rep] |
| legs | Stone Guard's Mail Legplates (220834) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stone Guard's Pulsing Legplates (220847, +0.00 DPS) [vendor]; Infernal Trickster Leggings (17754, -0.38 DPS) [dungeon]; Imperial Red Pants (8251, -1.82 DPS) [world_drop] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 91.0 attack_power points (7.70 DPS) | yes | Bloodshot Greaves (10846, -1.74 DPS) [world]; Skycaller's Leather Boots (252471, -2.36 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -2.36 DPS) [crafted] |
| finger1 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 65.9 attack_power points (5.58 DPS) | yes | Blackstone Ring (17713, -0.74 DPS) [dungeon]; Cyclopean Band (11824, -1.56 DPS) [dungeon]; Woodseed Hoop (17768, -1.57 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Cyclopean Band (11824, -0.55 DPS) [dungeon]; Woodseed Hoop (17768, -0.56 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (225.5 DPS) | yes | Uther's Strength (11302, -2.16 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (225.5 DPS) | yes | Uther's Strength (11302, -2.72 DPS) [world_drop] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-verified (225.5 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Mail Legplates; feet: Greaves of Withering Despair; finger1: Mindseye Circle; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Soulkeeper

No-known-source sample (15 of 1021, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 5030000000000000-055030310051202251-0530000000000000)

Set DPS (verified): 340.0. Weights run: 6.8s. Verify run: 75.1s. 2334 eligible items had no known source.

Race sweep: windshaper-skyborne 341.4 ± 0.4, troll 340.4 ± 0.4, orc 340.0 ± 0.4, tauren 337.6 ± 0.4 (dps); windshaper-skyborne beats the guide race orc by 1.5 (0.43%) ± 0.6, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.782 ± 0.082, crit=3.019 ± 0.123 per rating point (14 rating = 1%, 42.269 per %), hit=4.528 ± 0.253 per rating point (10 rating = 1%, 45.282 per %), melee_haste=14.169 ± 1.792, intellect=4.498 ± 0.107, spell_power=0.928 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 250.6 attack_power points (20.22 DPS) | yes | Cap of the Scarlet Savant (12752, -4.91 DPS, sim-verified) [quest]; Warlord's Mail Headguard (231657, -5.12 DPS) [vendor]; Warlord's Mail Helm (231663, -7.83 DPS) [pvp] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 95.7 attack_power points (7.72 DPS) | yes | Nacreous Shell Necklace (22403, -0.49 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.74 DPS) [world_drop]; Star of Mystaria (12103, -2.63 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Champion's Mail Pauldrons (227154, -0.02 DPS) [pvp]; Soulstealer Mantle (13374, -1.49 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -1.84 DPS) [vendor]; Darkspear Raider's Cloak (272063, -1.95 DPS) [vendor]; Royal Tribunal Cloak (13376, -2.84 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Mail Armor (16577, +0.00 DPS) [vendor]; Warlord's Mail Hauberk (231653, +0.00 DPS) [vendor]; Dawn Armor (252483, -2.20 DPS, sim-verified) [crafted] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Manacle Cuffs (11962, -0.00 DPS) [quest]; Aristocratic Cuffs (12546, -0.23 DPS) [dungeon]; Magiskull Cuffs (13107, -2.15 DPS, sim-verified) [world_drop] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Vices (231655, +0.00 DPS) [vendor]; General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Trueaim Gauntlets (13255, -3.37 DPS, sim-verified) [dungeon] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Waistband (16575, -0.03 DPS) [pvp]; Defiler's Lizardhide Girdle (20171, -0.33 DPS) [rep]; Clutch of Andros (13956, -1.49 DPS, sim-verified) [dungeon] |
| legs | Outrider's Lizardhide Pants (22741) | Warsong Outriders [rep] | 171.0 attack_power points (13.80 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; General's Mail Leggings (231664, +0.00 DPS) [pvp]; Outrider's Mail Leggings (22676, -0.14 DPS) [rep] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 130.0 attack_power points (10.49 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; Bloodmail Boots (14616, -0.60 DPS) [dungeon]; Greaves of Withering Despair (22240, -2.38 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -1.97 DPS) [dungeon]; Mindseye Circle (10634, -2.55 DPS) [dungeon]; Naglering (11669, -8.89 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindseye Circle (10634, -2.01 DPS) [dungeon]; Cutthroat's Signet (272408, -2.53 DPS) [vendor]; Seal of Rivendare (13345, -2.65 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, -0.03 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.40 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Gravestone War Axe (13983, -8.66 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Rage (22395) | Blackrock Depths: Magmus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Sustaining (23200, -2.45 DPS, sim-verified) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Darkspear Shoulderguards; back: Arcanoweave Cloak; chest: Timbermaw Tunic; wrist: Windtalker's Wristguards; hands: Raider Handwraps; waist: Stormseeker's Girdle; legs: Outrider's Lizardhide Pants; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Burst of Knowledge; main_hand: The Unstoppable Force; ranged: Totem of Rage

No-known-source sample (15 of 2334, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 0535001000000000-054030310051202051-0500000000000000)

Set DPS (verified): 876.0. Weights run: 6.9s. Verify run: 77.0s. 2334 eligible items had no known source.

Race sweep: windshaper-skyborne 880.6 ± 1.2, troll 876.8 ± 1.2, orc 876.0 ± 1.2, tauren 866.8 ± 1.2 (dps); windshaper-skyborne beats the guide race orc by 4.6 (0.52%) ± 1.8, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=1.440 ± 0.057, crit=2.439 ± 0.086 per rating point (14 rating = 1%, 34.143 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=24.720 ± 2.152, intellect=1.691 ± 0.022, spell_power=0.915 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 128.6 attack_power points (27.63 DPS) | yes | Warlord's Mail Headguard (231657, -4.24 DPS) [vendor]; Cap of the Scarlet Savant (12752, -5.70 DPS) [quest]; Skyfury Helm (20134, -8.65 DPS, sim-verified) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 60.1 attack_power points (12.92 DPS) | yes | Emberfury Talisman (12929, -4.41 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -5.20 DPS) [world_drop]; Lady Maye's Pendant (14558, -5.99 DPS, sim-verified) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 73.0 attack_power points (15.69 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.09 DPS) [rep]; Blackveil Cape (11626, -1.77 DPS) [dungeon]; Royal Tribunal Cloak (13376, -2.14 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 132.7 attack_power points (28.50 DPS) | yes | Warlord's Mail Hauberk (231653, -6.04 DPS) [vendor]; Savage Gladiator Chain (11726, -6.79 DPS, sim-verified) [dungeon]; Breastplate of Bloodthirst (12757, -7.68 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 51.6 attack_power points (11.08 DPS) | yes | Windtalker's Wristguards (19582, -0.01 DPS) [rep]; Prowler's Leather Bracers (252539, -2.84 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Vices (231655, +0.00 DPS) [vendor]; General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Bloodmail Gauntlets (14615, -24.14 DPS, sim-verified) [dungeon] |
| waist | Defiler's Chain Girdle (20150) (or Defiler's Leather Girdle (20190)) | The Defilers [rep] | 68.1 attack_power points (14.64 DPS) | yes | Defiler's Leather Girdle (20190, +0.00 DPS) [rep]; Windseeker's Belt (272410, -0.15 DPS) [vendor]; Might of the Timbermaw (19044, -0.38 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; Pristine Scorpid Leggings (252606, +0.00 DPS) [crafted]; Outrider's Lizardhide Pants (22741, -22.15 DPS, sim-verified) [rep] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; General's Mail Sabatons (231661, -0.25 DPS) [vendor]; Scalegut Treaders (275618, -8.43 DPS, sim-verified) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.85 DPS) [dungeon]; Legionnaire's Band (19510, -1.27 DPS) [rep]; Naglering (11669, -19.48 DPS, sim-verified) [dungeon] |
| finger2 | Painweaver Band (13098) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.42 DPS) [dungeon]; Legionnaire's Band (19510, -0.84 DPS) [rep]; Seal of Rivendare (13345, -2.17 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Hand of Justice (11815, -0.43 DPS) [dungeon]; Serenity Field (272439, -1.63 DPS) [vendor] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Heart of Wyrmthalak (22321, -15.21 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Darkspear Raider's Reaper (272079, -32.26 DPS, sim-verified) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Rage (22395) | Blackrock Depths: Magmus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, +0.00 DPS) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Devilsaur Gauntlets; waist: Defiler's Chain Girdle; legs: Devilsaur Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Painweaver Band; trinket1: Rune of the Guard Captain; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Totem of Rage

No-known-source sample (15 of 2334, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

