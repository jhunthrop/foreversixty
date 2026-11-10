# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-053030000000000000-0000000000000000)

Set DPS (verified): 39.8. Weights run: 3.9s. Verify run: 13.7s. 286 eligible items had no known source.

Race sweep: dwarf 39.8 ± 0.0 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.394 ± 0.028, crit=1.076 ± 0.028 per rating point (14 rating = 1%, 15.063 per %), hit=1.844 ± 0.072 per rating point (10 rating = 1%, 18.438 per %), melee_haste=9.856 ± 0.228, intellect=2.589 ± 0.025, spell_power=1.263 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Trapper's Leather Hood (252505) | Leatherworking [crafted] | 20.7 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.27 DPS) [crafted]; Totemic Leather Hood (252448, -0.32 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.19 DPS) [quest]; Tarnished Locket (279870, -0.49 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 29.6 attack_power points (1.01 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.52 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.77 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.84 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (39.8 DPS) | yes | Pearl-clasped Cloak (5542, -0.00 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon]; Dark Leather Cloak (2316, -0.07 DPS) [crafted] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mystic's Wrap (14369, +0.00 DPS) [world_drop]; Stormrider's Leather Armor (252492, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.65 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bright Bracers (3647, -0.09 DPS) [world_drop]; Bravo's Armbands (270015, -0.09 DPS) [quest]; Owl Bracers (4796, -0.41 DPS, sim-verified) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormrider's Leather Gloves (252498, +0.00 DPS) [crafted]; Fletcher's Gloves (7348, -0.04 DPS) [crafted]; Serpent Gloves (5970, -0.94 DPS, sim-verified) [dungeon] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 20.7 attack_power points (0.71 DPS) | yes | Blackened Defias Belt (10403, -0.09 DPS) [dungeon]; Stormrider's Leather Belt (252432, -0.18 DPS) [crafted]; Pristine Sash (253925, -0.18 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 37.2 attack_power points (1.27 DPS) | yes | Stormrider's Leather Pants (252502, -0.31 DPS) [crafted]; Trapper's Leather Pants (252501, -0.38 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.43 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 20.5 attack_power points (0.70 DPS) | yes | Spidersilk Boots (4320, -0.05 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.09 DPS) [crafted]; Totemic Leather Boots (252442, -0.10 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 18.3 attack_power points (0.63 DPS) | yes | Minor Channeling Ring (1449, -0.23 DPS) [quest]; Demon Band (12054, -0.35 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.6 attack_power points (0.46 DPS) | yes | Minor Channeling Ring (1449, -0.07 DPS) [quest]; Demon Band (12054, -0.19 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.20 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 312.1 attack_power points (10.68 DPS) | yes | Smite's Mighty Hammer (7230, -0.24 DPS, sim-verified) [dungeon]; Rockslicer (872, -0.76 DPS) [dungeon]; Night Reaver (1318, -1.24 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Trapper's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Armor of the Fang; wrist: Mindthrust Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Totemic Leather Pants; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish

### Band 30 (dwarf, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 103.6. Weights run: 4.0s. Verify run: 7.7s. 475 eligible items had no known source.

Race sweep: dwarf 103.6 ± 0.1 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.170 ± 0.032, crit=1.054 ± 0.036 per rating point (14 rating = 1%, 14.752 per %), hit=1.804 ± 0.072 per rating point (10 rating = 1%, 18.037 per %), melee_haste=4.966 ± 0.533, intellect=2.697 ± 0.018, spell_power=0.449 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 43.2 attack_power points (3.82 DPS) | yes | Dokebi Hat (14584, -0.49 DPS) [world_drop]; Nightsky Cowl (4039, -0.96 DPS) [world_drop]; Resilient Cap (14401, -0.99 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 28.2 attack_power points (2.50 DPS) | yes | Kaleidoscope Chain (13084, -0.42 DPS) [world_drop]; Seal of the Expedition (271920, -1.30 DPS) [rep]; Crystal Starfire Medallion (5003, -1.54 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 32.4 attack_power points (2.87 DPS) | yes | Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Magician's Mantle (12998, -0.52 DPS) [world_drop]; Feline Mantle (3748, -0.59 DPS) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resilient Cape (14400, -0.65 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.82 DPS) [world_drop]; Expeditionary's Cape (271923, -0.97 DPS, sim-verified) [rep] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 51.0 attack_power points (4.52 DPS) | yes | Spirewind Fetter (9406, -0.41 DPS) [dungeon]; Green Silk Armor (7065, -1.05 DPS) [crafted]; Beguiler Robes (7728, -1.29 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 27.6 attack_power points (2.44 DPS) | yes | Technician's Bracers (270042, -0.78 DPS) [quest]; Nightsky Wristbands (6407, -1.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -1.11 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 26.3 attack_power points (2.33 DPS) | yes | Insignia Gloves (6408, -0.46 DPS) [world_drop]; Town Clerk's Mittens (270029, -0.53 DPS) [quest]; Resilient Handgrips (14403, -0.58 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 35.8 attack_power points (3.17 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.30 DPS) [rep]; Highlander's Mail Girdle (20120, -0.30 DPS) [vendor]; Skirmisher's Leather Belt (252461, -0.34 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 58.4 attack_power points (5.17 DPS) | yes | Totemic Leather Leggings (252458, -1.61 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, -2.22 DPS) [crafted]; Stormrider's Leather Kilt (252518, -2.26 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 23.8 attack_power points (2.11 DPS) | yes | Gilded Slippers (254001, -0.16 DPS) [crafted]; Frothing Slippers (254003, -0.44 DPS) [crafted]; Fiery Slippers (254005, -0.44 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 21.6 attack_power points (1.91 DPS) | yes | Thunderbrow Ring (13097, -0.18 DPS) [world_drop]; Protector's Band (19517, -0.23 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 20.3 attack_power points (1.80 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Protector's Band (19517, -0.11 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -2.40 DPS) [dungeon]; Manual Crowd Pummeler (9449, -2.92 DPS, sim-verified) [dungeon]; Cobalt Crusher (7730, -5.11 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Sergeant Major's Cape; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Snake Hoop; finger2: Seal of Wrynn; main_hand: Viscous Hammer

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 0000000000000000-055030310051002051-0000000000000000)

Set DPS (verified): 175.5. Weights run: 4.0s. Verify run: 10.8s. 804 eligible items had no known source.

Race sweep: dwarf 175.5 ± 0.3 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.362 ± 0.046, crit=1.481 ± 0.057 per rating point (14 rating = 1%, 20.732 per %), hit=2.538 ± 0.115 per rating point (10 rating = 1%, 25.380 per %), melee_haste=9.804 ± 0.840, intellect=4.042 ± 0.057, spell_power=1.202 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 92.0 attack_power points (7.29 DPS) | yes | Miner's Hat of the Deep (9429, -1.08 DPS, sim-verified) [dungeon]; Papal Fez (9431, -1.09 DPS) [dungeon]; Thinking Cap (2624, -1.85 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 40.7 attack_power points (3.23 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.36 DPS) [world]; Kaleidoscope Chain (13084, -0.88 DPS) [world_drop]; Seal of the Expedition (271919, -0.99 DPS) [rep] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 64.7 attack_power points (5.13 DPS) | yes | Inquisitor's Shawl (19507, -0.29 DPS) [dungeon]; Windchaser Amice (14432, -0.64 DPS) [world_drop]; Green Silken Shoulders (7057, -0.84 DPS) [crafted] |
| back | Expeditionary's Cape (271922) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16336, -0.00 DPS) [pvp]; Blackforge Cape (6424, -0.96 DPS) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 72.9 attack_power points (5.78 DPS) | yes | Red Mageweave Vest (10007, -0.01 DPS) [crafted]; Big Voodoo Robe (8200, -0.43 DPS) [crafted]; Robe of Power (7054, -0.60 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 39.1 attack_power points (3.10 DPS) | yes | Glowing Magical Bracelets (13106, -0.06 DPS) [world_drop]; Warden's Wristbands (14600, -0.21 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.42 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Stormcloth Gloves (10011, +0.00 DPS) [crafted]; Red Mageweave Gloves (10018, -2.35 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 69.0 attack_power points (5.47 DPS) | yes | Highlander's Lizardhide Girdle (20104, -0.67 DPS) [rep]; Highlander's Mail Girdle (20119, -0.67 DPS) [vendor]; Saltscale Girdle (273023, -1.16 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 74.5 attack_power points (5.91 DPS) | yes | Red Mageweave Pants (10009, -0.73 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.98 DPS) [crafted]; Misplaced Pantaloons (276201, -1.74 DPS) [vendor] |
| feet | Skirmisher's Mail Boots (252564) | Leatherworking [crafted] | 53.0 attack_power points (4.20 DPS) | yes | Prowler's Leather Shoes (252465, -0.10 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.72 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.72 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 35.5 attack_power points (2.81 DPS) | yes | Ogremind Ring (1993, -0.57 DPS) [world_drop]; Voodoo Band (1996, -0.57 DPS) [world]; Mindbender Loop (5009, -0.57 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 32.2 attack_power points (2.55 DPS) | yes | Ogremind Ring (1993, -0.31 DPS) [world_drop]; Mindbender Loop (5009, -0.31 DPS) [world_drop]; Voodoo Band (1996, -1.74 DPS, sim-verified) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Bonebiter (6830, -0.98 DPS) [quest]; Mograine's Might (7723, -6.53 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Sheepshear Mantle; back: Expeditionary's Cape; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Skirmisher's Mail Boots; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 3000000000000000-055030310051202051-0500000000000000)

Set DPS (verified): 258.6. Weights run: 4.3s. Verify run: 15.0s. 1050 eligible items had no known source.

Race sweep: dwarf 258.6 ± 0.4 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.579 ± 0.073, crit=2.213 ± 0.097 per rating point (14 rating = 1%, 30.982 per %), hit=3.978 ± 0.151 per rating point (10 rating = 1%, 39.781 per %), melee_haste=11.708 ± 1.380, intellect=4.480 ± 0.070, spell_power=1.120 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 147.3 attack_power points (14.91 DPS) | yes | Braincage (12549, -1.81 DPS) [dungeon]; Soulcatcher Halo (10630, -2.89 DPS) [dungeon]; Eye of Theradras (17715, -3.16 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 62.7 attack_power points (6.35 DPS) | yes | Gemshard Heart (17707, -1.70 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.93 DPS) [dungeon]; Seal of the Expedition (271918, -2.27 DPS) [rep] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 106.4 attack_power points (10.77 DPS) | yes | Rotgrip Mantle (17732, -1.13 DPS) [dungeon]; Dregmetal Spaulders (11722, -1.44 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -2.72 DPS) [dungeon] |
| back | Expeditionary's Cape (271921) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Red Cloak (8248, -1.36 DPS) [world_drop]; Sergeant Major's Cape (16336, -1.45 DPS) [pvp]; Big Voodoo Cloak (8216, -1.70 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 119.8 attack_power points (12.13 DPS) | yes | Knight's Mail Armor (223078, -1.12 DPS) [vendor]; Hibernal Robe (8113, -3.06 DPS) [world_drop]; Runecloth Robe (13858, -3.18 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 71.7 attack_power points (7.26 DPS) | yes | Forgotten Wraps (9433, -1.33 DPS) [world_drop]; Runic Leather Bracers (15092, -1.59 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -1.81 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Sergeant Major's Mail Gauntlets (223076, +0.00 DPS) [vendor]; Rockgrip Gauntlets (17736, -2.04 DPS, sim-verified) [dungeon] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Belt (13144, -1.59 DPS) [world_drop]; Highlander's Lizardhide Girdle (20103, -1.63 DPS) [rep]; Ban'thok Sash (11662, -2.81 DPS, sim-verified) [dungeon] |
| legs | Knight's Mail Legplates (223074) | PvP rank 6 · Knight · Alliance [vendor] | 108.7 attack_power points (11.01 DPS) | yes | Imperial Red Pants (8251, -1.81 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -2.51 DPS) [dungeon]; Infernal Trickster Leggings (17754, -2.64 DPS, sim-verified) [dungeon] |
| feet | Bloodshot Greaves (10846) | Avatar of Hakkar [world] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Leather Boots (252471, -1.18 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.18 DPS) [crafted]; Greaves of Withering Despair (22240, -1.50 DPS, sim-verified) [dungeon] |
| finger1 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 73.6 attack_power points (7.46 DPS) | yes | Woodseed Hoop (17768, -2.24 DPS) [quest]; Cyclopean Band (11824, -2.45 DPS) [dungeon]; Blackstone Ring (17713, -5.95 DPS, sim-verified) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Woodseed Hoop (17768, -0.57 DPS) [quest]; Cyclopean Band (11824, -0.78 DPS) [dungeon]; Blackstone Ring (17713, -3.59 DPS, sim-verified) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Molten Heart of the Mountain (249470, -2.81 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.95 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -1.22 DPS) [world_drop]; Soulkeeper (1607, -2.98 DPS) [world_drop]; Fiery War Axe (870, -5.29 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Expeditionary's Cape; chest: Robes of Insight; wrist: Aristocratic Cuffs; waist: Dawnspire Cord; legs: Knight's Mail Legplates; feet: Bloodshot Greaves; finger1: Mindseye Circle; finger2: Brainlash; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 1050, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 3230030000000000-055030310051202051-0520000000000000)

Set DPS (verified): 378.2. Weights run: 4.4s. Verify run: 49.7s. 2350 eligible items had no known source.

Race sweep: dwarf 378.2 ± 0.5 (dps); the guide race dwarf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.807 ± 0.082, crit=2.997 ± 0.122 per rating point (14 rating = 1%, 41.962 per %), hit=4.194 ± 0.247 per rating point (10 rating = 1%, 41.939 per %), melee_haste=15.818 ± 1.714, intellect=5.241 ± 0.124, spell_power=0.970 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 279.1 attack_power points (24.80 DPS) | yes | Cap of the Scarlet Savant (12752, -5.58 DPS, sim-verified) [quest]; Magister's Crown (16686, -9.88 DPS) [dungeon]; Chief Architect's Monocle (11839, -9.90 DPS) [dungeon] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 110.3 attack_power points (9.80 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.90 DPS) [world_drop]; Nacreous Shell Necklace (22403, -1.57 DPS) [dungeon]; Star of Mystaria (12103, -1.69 DPS) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Theramore Shoulderpads (271953, -0.34 DPS) [rep]; Theramore Shoulders (271954, -0.34 DPS) [rep]; Soulstealer Mantle (13374, -5.99 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, +0.00 DPS) [dungeon]; Royal Tribunal Cloak (13376, -0.95 DPS) [dungeon]; Expeditionary's Cape (271907, -1.38 DPS) [rep] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 158.4 attack_power points (14.08 DPS) | yes | Blue Dragonscale Breastplate (15048, -0.34 DPS) [crafted]; Dawn Armor (252483, -0.68 DPS) [crafted]; Timbermaw Tunic (252484, -1.59 DPS) [crafted] |
| wrist | Magiskull Cuffs (13107) | World drop [world_drop] | 89.1 attack_power points (7.92 DPS) | yes | Manacle Cuffs (11962, -0.33 DPS) [quest]; Aristocratic Cuffs (12546, -0.59 DPS) [dungeon]; Windtalker's Wristguards (19582, -0.81 DPS) [rep] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, -2.08 DPS) [dungeon]; Theramore Handguards (271950, -2.10 DPS) [rep]; Gloves of Holy Might (867, -2.36 DPS, sim-verified) [world_drop] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Mail Girdle (20044, -0.91 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, -0.91 DPS) [rep]; Clutch of Andros (13956, -5.99 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 189.1 attack_power points (16.81 DPS) | yes | Sentinel's Lizardhide Pants (22750, -0.09 DPS) [rep]; Red Dragonscale Leggings (252603, -0.18 DPS) [crafted]; Ironfeather Leggings (252486, -1.42 DPS) [crafted] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 140.2 attack_power points (12.46 DPS) | yes | Bloodmail Boots (14616, -1.19 DPS) [dungeon]; Greaves of Withering Despair (22240, -3.12 DPS) [dungeon]; Dragonrider Boots (18102, -3.45 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -1.64 DPS) [dungeon]; Mindseye Circle (10634, -2.80 DPS) [dungeon]; Naglering (11669, -6.74 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.09 DPS) [dungeon]; Mindseye Circle (10634, -1.25 DPS) [dungeon]; Cutthroat's Signet (272408, -2.73 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+10.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -4.04 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -2.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (272432) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Rage (22395, -2.22 DPS, sim-verified) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Theramore Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Magiskull Cuffs; hands: Theramore Handwraps; waist: Stormseeker's Girdle; legs: Sentinel's Chain Leggings; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: The Unstoppable Force; ranged: Totem of the Storm

No-known-source sample (15 of 2350, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 0533001000000000-053030310051202051-0530000000000000)

Set DPS (verified): 940.6. Weights run: 3.7s. Verify run: 46.6s. 2350 eligible items had no known source.

Race sweep: dwarf 940.6 ± 0.9 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.826 ± 0.067, crit=3.013 ± 0.097 per rating point (14 rating = 1%, 42.181 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=19.431 ± 2.186, intellect=2.315 ± 0.020, spell_power=1.062 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 168.5 attack_power points (31.66 DPS) | yes | Cap of the Scarlet Savant (12752, -8.72 DPS, sim-verified) [quest]; Ragefury Eyepatch (11735, -13.33 DPS) [dungeon]; Eye of Rend (12587, -13.57 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 68.2 attack_power points (12.81 DPS) | yes | Lady Maye's Pendant (14558, -2.35 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -3.78 DPS) [world_drop]; Emberfury Talisman (12929, -4.88 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 91.5 attack_power points (17.19 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Mantle of the Timbermaw (19050, -0.22 DPS) [crafted]; Shroud of the Nathrezim (18720, -0.91 DPS) [dungeon] |
| back | Sergeant Major's Cape (16337) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.63 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.43 DPS) [rep]; Royal Tribunal Cloak (13376, -1.57 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 147.7 attack_power points (27.74 DPS) | yes | Savage Gladiator Chain (11726, -4.56 DPS, sim-verified) [dungeon]; Breastplate of Bloodthirst (12757, -6.52 DPS) [quest]; Tombstone Breastplate (13944, -6.93 DPS) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 58.9 attack_power points (11.06 DPS) | yes | Windtalker's Wristguards (19582, -0.45 DPS) [rep]; Spitfire Bracers (20481, -2.47 DPS) [crafted] |
| hands | Bloodmail Gauntlets (14615) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Theramore Gloves (271948, -3.01 DPS) [rep]; Timbermaw Brawlers (19049, -3.34 DPS) [crafted]; Gloves of Holy Might (867, -14.23 DPS, sim-verified) [world_drop] |
| waist | Highlander's Mail Girdle (20044) (or Highlander's Lizardhide Girdle (20046)) | Samuel Hawke [vendor] | 81.5 attack_power points (15.32 DPS) | yes | Highlander's Lizardhide Girdle (20046, +0.00 DPS) [rep]; Highlander's Chain Girdle (20043, -1.01 DPS) [rep]; Highlander's Leather Girdle (20045, -1.01 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 148.3 attack_power points (27.86 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.27 DPS) [vendor]; Sentinel's Leather Pants (237818, -2.74 DPS) [vendor]; Sentinel's Silk Leggings (237815, -17.00 DPS, sim-verified) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | 57.6 attack_power points (10.81 DPS) | yes | Timmy's Galoshes (13402, -0.29 DPS) [dungeon]; Scalegut Treaders (275618, -0.93 DPS) [crafted]; Dragonrider Boots (18102, -15.03 DPS, sim-verified) [dungeon] |
| finger1 | Painweaver Band (13098) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magni's Will (12548, -0.11 DPS) [quest]; Don Mauricio's Band of Domination (22433, -0.81 DPS) [dungeon]; Naglering (11669, -12.40 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (+10.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Don Mauricio's Band of Domination (22433, -0.43 DPS) [dungeon]; Seal of Rivendare (13345, -1.56 DPS) [dungeon]; Magni's Will (12548, -10.92 DPS, sim-verified) [quest] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, -0.86 DPS) [vendor]; Heart of Wyrmthalak (22321, -8.48 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Rebels' Rugged Reaper (271924, -27.47 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (272432) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Rage (22395, -5.94 DPS, sim-verified) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Sergeant Major's Cape; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Bloodmail Gauntlets; waist: Highlander's Mail Girdle; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Painweaver Band; finger2: Don Julio's Band; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Totem of the Storm

No-known-source sample (15 of 2350, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 0000000000000000-053030000000000000-0000000000000000)

Set DPS (verified): 40.1. Weights run: 3.9s. Verify run: 14.8s. 275 eligible items had no known source.

Race sweep: windshaper-skyborne 40.5 ± 0.0, orc 40.1 ± 0.0, troll 40.1 ± 0.0, tauren 39.7 ± 0.0 (dps); windshaper-skyborne beats the guide race orc by 0.4 (1.00%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.973 ± 0.023, crit=0.783 ± 0.022 per rating point (14 rating = 1%, 10.957 per %), hit=1.448 ± 0.061 per rating point (10 rating = 1%, 14.480 per %), melee_haste=6.807 ± 0.217, intellect=2.659 ± 0.024, spell_power=1.385 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Trapper's Leather Hood (252505) | Leatherworking [crafted] | 21.3 attack_power points (0.75 DPS) | yes | Defender's Leather Hood (252447, -0.19 DPS) [crafted]; Shadow Goggles (4373, -0.28 DPS) [crafted]; Totemic Leather Hood (252448, -0.31 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 5.8 attack_power points (0.21 DPS) | yes | Erudite's Amulet (277204, -0.14 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 30.9 attack_power points (1.09 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.39 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.89 DPS) [crafted]; Serpent's Shoulders (5404, -0.91 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-verified (40.1 DPS) | yes | Sanguine Cape (14376, -0.00 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Dark Leather Cloak (2316, -0.13 DPS) [crafted] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mystic's Wrap (14369, +0.00 DPS) [world_drop]; Stormrider's Leather Armor (252492, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.47 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bright Bracers (3647, -0.09 DPS) [world_drop]; Featherbead Bracers (15452, -0.19 DPS) [quest]; Owl Bracers (4796, -0.22 DPS, sim-verified) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Totemic Leather Gloves (252497, +0.00 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.78 DPS, sim-verified) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 21.3 attack_power points (0.75 DPS) | yes | Blackened Defias Belt (10403, -0.12 DPS) [dungeon]; Stormrider's Leather Belt (252432, -0.18 DPS) [crafted]; Pristine Sash (253925, -0.18 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 38.3 attack_power points (1.35 DPS) | yes | Trapper's Leather Pants (252501, -0.25 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.30 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.45 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 21.6 attack_power points (0.76 DPS) | yes | Spidersilk Boots (4320, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.10 DPS) [crafted]; Totemic Leather Boots (252442, -0.12 DPS) [crafted] |
| finger1 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 19.3 attack_power points (0.68 DPS) | yes | Legionnaire's Band (20429, -0.26 DPS) [rep]; Demon Band (12054, -0.40 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 17.9 attack_power points (0.63 DPS) | yes | Legionnaire's Band (20429, -0.21 DPS) [rep]; Demon Band (12054, -0.35 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 312.8 attack_power points (11.01 DPS) | yes | Smite's Mighty Hammer (7230, -0.38 DPS) [dungeon]; Forsaken Greataxe (251533, -0.41 DPS, sim-verified) [quest]; Hammerbone (270018, -0.44 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Trapper's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Armor of the Fang; wrist: Mindthrust Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Totemic Leather Pants; feet: Stormrider's Leather Boots; finger1: Loop of Sacrifice; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 275, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield

### Band 30 (orc, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 104.2. Weights run: 4.0s. Verify run: 8.5s. 469 eligible items had no known source.

Race sweep: windshaper-skyborne 105.0 ± 0.1, orc 104.2 ± 0.1, troll 103.8 ± 0.1, tauren 103.0 ± 0.1 (dps); windshaper-skyborne beats the guide race orc by 0.8 (0.73%) ± 0.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.170 ± 0.032, crit=1.054 ± 0.036 per rating point (14 rating = 1%, 14.752 per %), hit=1.804 ± 0.072 per rating point (10 rating = 1%, 18.037 per %), melee_haste=4.966 ± 0.533, intellect=2.697 ± 0.018, spell_power=0.449 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 43.2 attack_power points (3.82 DPS) | yes | Dokebi Hat (14584, -0.49 DPS) [world_drop]; Nightsky Cowl (4039, -0.96 DPS) [world_drop]; Resilient Cap (14401, -0.99 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 28.2 attack_power points (2.50 DPS) | yes | Kaleidoscope Chain (13084, -0.42 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.30 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.54 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 32.4 attack_power points (2.87 DPS) | yes | Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Magician's Mantle (12998, -0.52 DPS) [world_drop]; Feline Mantle (3748, -0.59 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.65 DPS) [world_drop]; Resilient Cape (14400, -0.69 DPS, sim-verified) [world_drop]; Soft Willow Cape (16661, -0.72 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 51.0 attack_power points (4.52 DPS) | yes | Spirewind Fetter (9406, -0.41 DPS) [dungeon]; Green Silk Armor (7065, -1.05 DPS) [crafted]; Beguiler Robes (7728, -1.29 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 27.6 attack_power points (2.44 DPS) | yes | Technician's Bracers (270042, -0.82 DPS, sim-verified) [quest]; Nightsky Wristbands (6407, -1.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -1.11 DPS) [world_drop] |
| hands | Gloves of Kapelan (6744) | Alliance Relations [quest] | 30.2 attack_power points (2.67 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.35 DPS) [dungeon]; Insignia Gloves (6408, -0.81 DPS) [world_drop]; Resilient Handgrips (14403, -0.93 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 35.8 attack_power points (3.17 DPS) | yes | Defiler's Mail Girdle (20197, -0.30 DPS) [rep]; Skirmisher's Leather Belt (252461, -0.34 DPS) [crafted]; Defiler's Lizardhide Girdle (20172, -1.63 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 58.4 attack_power points (5.17 DPS) | yes | Totemic Leather Leggings (252458, -1.64 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, -2.22 DPS) [crafted]; Stormrider's Leather Kilt (252518, -2.26 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 23.8 attack_power points (2.11 DPS) | yes | Gilded Slippers (254001, -0.16 DPS) [crafted]; Frothing Slippers (254003, -0.44 DPS) [crafted]; Fiery Slippers (254005, -0.44 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 21.6 attack_power points (1.91 DPS) | yes | Loop of Sacrifice (281673, -0.19 DPS) [quest]; Legionnaire's Band (19513, -0.23 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.5 attack_power points (1.73 DPS) | yes | Loop of Sacrifice (281673, -0.00 DPS) [quest]; Legionnaire's Band (19513, -0.04 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -2.40 DPS) [dungeon]; Manual Crowd Pummeler (9449, -2.83 DPS, sim-verified) [dungeon]; Cobalt Crusher (7730, -5.11 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Darkspear Raider's Cloak; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Gloves of Kapelan; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Snake Hoop; finger2: Thunderbrow Ring; main_hand: Viscous Hammer

No-known-source sample (15 of 469, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 0000000000000000-055030310051002051-0000000000000000)

Set DPS (verified): 175.4. Weights run: 4.0s. Verify run: 11.6s. 787 eligible items had no known source.

Race sweep: windshaper-skyborne 176.6 ± 0.3, orc 175.4 ± 0.3, troll 175.0 ± 0.3, tauren 172.8 ± 0.3 (dps); windshaper-skyborne beats the guide race orc by 1.1 (0.63%) ± 0.4, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.362 ± 0.046, crit=1.481 ± 0.057 per rating point (14 rating = 1%, 20.732 per %), hit=2.538 ± 0.115 per rating point (10 rating = 1%, 25.380 per %), melee_haste=9.804 ± 0.840, intellect=4.042 ± 0.057, spell_power=1.202 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 92.0 attack_power points (7.29 DPS) | yes | Papal Fez (9431, -1.09 DPS) [dungeon]; Miner's Hat of the Deep (9429, -1.64 DPS, sim-verified) [dungeon]; Thinking Cap (2624, -1.85 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 40.7 attack_power points (3.23 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.36 DPS) [world]; Kaleidoscope Chain (13084, -0.88 DPS) [world_drop]; Darkspear Warding Pendant (272074, -0.99 DPS) [vendor] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 64.7 attack_power points (5.13 DPS) | yes | Inquisitor's Shawl (19507, -0.29 DPS) [dungeon]; Windchaser Amice (14432, -0.64 DPS) [world_drop]; Green Silken Shoulders (7057, -0.84 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackforge Cape (6424, -1.23 DPS, sim-verified) [dungeon]; Glyphed Cloak (4732, -1.28 DPS) [world_drop]; Long Silken Cloak (4326, -1.35 DPS) [crafted] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 72.9 attack_power points (5.78 DPS) | yes | Red Mageweave Vest (10007, -0.01 DPS) [crafted]; Big Voodoo Robe (8200, -0.43 DPS) [crafted]; Robe of Power (7054, -0.60 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 39.1 attack_power points (3.10 DPS) | yes | Glowing Magical Bracelets (13106, -0.06 DPS) [world_drop]; Radiant Silver Bracers (4545, -0.16 DPS) [quest]; Warden's Wristbands (14600, -0.21 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Gloves (10011, +0.00 DPS) [crafted]; Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Rustler Gloves (9704, -2.75 DPS, sim-verified) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 69.0 attack_power points (5.47 DPS) | yes | Defiler's Lizardhide Girdle (20173, -0.67 DPS) [rep]; Defiler's Mail Girdle (20198, -0.67 DPS) [rep]; Saltscale Girdle (273023, -1.16 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 74.5 attack_power points (5.91 DPS) | yes | Red Mageweave Pants (10009, -0.73 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.98 DPS) [crafted]; Misplaced Pantaloons (276201, -1.74 DPS) [vendor] |
| feet | Skirmisher's Mail Boots (252564) | Leatherworking [crafted] | 53.0 attack_power points (4.20 DPS) | yes | Prowler's Leather Shoes (252465, -0.10 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.72 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.72 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 35.5 attack_power points (2.81 DPS) | yes | Ogremind Ring (1993, -0.57 DPS) [world_drop]; Voodoo Band (1996, -0.57 DPS) [world]; Mindbender Loop (5009, -0.57 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 32.2 attack_power points (2.55 DPS) | yes | Ogremind Ring (1993, -0.31 DPS) [world_drop]; Voodoo Band (1996, -0.31 DPS) [world]; Mindbender Loop (5009, -0.31 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Pendulum of Doom (9425, -6.28 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Corpseshroud; neck: Glowing Eye of Mordresh; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Skirmisher's Mail Boots; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 787, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 3000000000000000-055030310051202051-0500000000000000)

Set DPS (verified): 257.8. Weights run: 4.3s. Verify run: 17.1s. 1018 eligible items had no known source.

Race sweep: windshaper-skyborne 259.9 ± 0.4, orc 257.8 ± 0.4, troll 257.7 ± 0.4, tauren 255.3 ± 0.4 (dps); windshaper-skyborne beats the guide race orc by 2.1 (0.80%) ± 0.5, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.579 ± 0.073, crit=2.213 ± 0.097 per rating point (14 rating = 1%, 30.982 per %), hit=3.978 ± 0.151 per rating point (10 rating = 1%, 39.781 per %), melee_haste=11.708 ± 1.380, intellect=4.480 ± 0.070, spell_power=1.120 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 147.3 attack_power points (14.91 DPS) | yes | Braincage (12549, -1.81 DPS) [dungeon]; Soulcatcher Halo (10630, -2.89 DPS) [dungeon]; Eye of Theradras (17715, -3.16 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 62.7 attack_power points (6.35 DPS) | yes | Gemshard Heart (17707, -1.70 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.93 DPS) [dungeon]; Darkspear Warding Pendant (272073, -2.27 DPS) [vendor] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 106.4 attack_power points (10.77 DPS) | yes | Rotgrip Mantle (17732, -1.13 DPS) [dungeon]; Dregmetal Spaulders (11722, -1.44 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -2.72 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.91 DPS) [quest]; Imperial Red Cloak (8248, -1.36 DPS) [world_drop]; Big Voodoo Cloak (8216, -1.70 DPS) [crafted] |
| chest | Robes of Insight (940) | World drop [world_drop] | 119.8 attack_power points (12.13 DPS) | yes | Stone Guard's Mail Armor (220826, -1.12 DPS) [vendor]; Stone Guard's Pulsing Breastplate (220844, -1.64 DPS, sim-verified) [vendor]; Hibernal Robe (8113, -3.06 DPS) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths [dungeon] | 71.7 attack_power points (7.26 DPS) | yes | Forgotten Wraps (9433, -1.33 DPS) [world_drop]; Runic Leather Bracers (15092, -1.59 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -1.81 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Rockgrip Gauntlets (17736, +0.00 DPS) [dungeon]; Azurite Fists (20369, +0.00 DPS) [quest]; Stonerender Gauntlets (17007, -0.92 DPS, sim-verified) [world_drop] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Belt (13144, -1.59 DPS) [world_drop]; Defiler's Lizardhide Girdle (20174, -1.63 DPS) [rep]; Ban'thok Sash (11662, -2.74 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Mail Legplates (220834) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Infernal Trickster Leggings (17754, -0.16 DPS) [dungeon]; Imperial Red Pants (8251, -1.81 DPS) [world_drop]; Stone Guard's Pulsing Legplates (220847, -2.36 DPS, sim-verified) [vendor] |
| feet | Bloodshot Greaves (10846) | Avatar of Hakkar [world] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Leather Boots (252471, -1.18 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.18 DPS) [crafted]; Greaves of Withering Despair (22240, -1.88 DPS, sim-verified) [dungeon] |
| finger1 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 73.6 attack_power points (7.46 DPS) | yes | Woodseed Hoop (17768, -2.24 DPS) [quest]; Cyclopean Band (11824, -2.45 DPS) [dungeon]; Blackstone Ring (17713, -3.68 DPS, sim-verified) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Woodseed Hoop (17768, -0.57 DPS) [quest]; Cyclopean Band (11824, -0.78 DPS) [dungeon]; Blackstone Ring (17713, -2.65 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -3.56 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.82 DPS, sim-verified) [crafted]; Uther's Strength (11302, -2.95 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -1.74 DPS, sim-verified) [world_drop]; Soulkeeper (1607, -2.98 DPS) [world_drop]; Taran Icebreaker (2915, -3.49 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; waist: Dawnspire Cord; legs: Stone Guard's Mail Legplates; feet: Bloodshot Greaves; finger1: Mindseye Circle; finger2: Brainlash; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 1018, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 3230030000000000-055030310051202051-0520000000000000)

Set DPS (verified): 376.1. Weights run: 4.4s. Verify run: 49.5s. 2307 eligible items had no known source.

Race sweep: windshaper-skyborne 377.7 ± 0.5, orc 376.1 ± 0.5, troll 375.9 ± 0.5, tauren 373.1 ± 0.5 (dps); windshaper-skyborne beats the guide race orc by 1.6 (0.43%) ± 0.7, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.807 ± 0.082, crit=2.997 ± 0.122 per rating point (14 rating = 1%, 41.962 per %), hit=4.194 ± 0.247 per rating point (10 rating = 1%, 41.939 per %), melee_haste=15.818 ± 1.714, intellect=5.241 ± 0.124, spell_power=0.970 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 279.1 attack_power points (24.80 DPS) | yes | Cap of the Scarlet Savant (12752, -4.04 DPS, sim-verified) [quest]; Warlord's Mail Headguard (231657, -6.47 DPS) [vendor]; Warlord's Mail Helm (231663, -9.77 DPS) [pvp] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 110.3 attack_power points (9.80 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.90 DPS) [world_drop]; Nacreous Shell Necklace (22403, -1.57 DPS) [dungeon]; Star of Mystaria (12103, -1.69 DPS) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.34 DPS) [vendor]; Soulstealer Mantle (13374, -1.92 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Tribunal Cloak (13376, -0.95 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -1.38 DPS) [vendor]; Shroud of Arcane Mastery (22330, -1.60 DPS, sim-verified) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 158.4 attack_power points (14.08 DPS) | yes | Blue Dragonscale Breastplate (15048, +0.00 DPS) [crafted]; Warlord's Mail Armor (16577, +0.00 DPS) [vendor]; Warlord's Mail Hauberk (231653, +0.00 DPS) [vendor] |
| wrist | Magiskull Cuffs (13107) | World drop [world_drop] | 89.1 attack_power points (7.92 DPS) | yes | Manacle Cuffs (11962, -0.33 DPS) [quest]; Aristocratic Cuffs (12546, -0.59 DPS) [dungeon]; Windtalker's Wristguards (19582, -0.81 DPS) [rep] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; General's Mail Vices (231655, -0.17 DPS) [vendor]; General's Mail Gauntlets (231660, -0.74 DPS) [pvp] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Waistband (16575, -0.03 DPS) [pvp]; Defiler's Lizardhide Girdle (20171, -0.91 DPS) [rep]; Clutch of Andros (13956, -1.92 DPS, sim-verified) [dungeon] |
| legs | Outrider's Lizardhide Pants (22741) | Warsong Outriders [rep] | 188.0 attack_power points (16.71 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; General's Mail Leggings (231664, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, -2.45 DPS, sim-verified) [crafted] |
| feet | Kayser's Boots of Precision (22231) | Blackrock Spire: War Master Voone [dungeon] | 140.2 attack_power points (12.46 DPS) | yes | General's Mail Greaves (231656, -0.12 DPS) [vendor]; Bloodmail Boots (14616, -1.19 DPS) [dungeon]; General's Mail Sabatons (231661, -2.81 DPS) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Julio's Band (19325, -1.55 DPS) [rep]; Mindseye Circle (10634, -2.80 DPS) [dungeon]; Naglering (11669, -5.02 DPS, sim-verified) [dungeon] |
| finger2 | Seal of Rivendare (13345) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindseye Circle (10634, -1.16 DPS) [dungeon]; Don Julio's Band (19325, -1.94 DPS, sim-verified) [rep]; Cutthroat's Signet (272408, -2.63 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Hand of Justice (11815, -3.78 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -2.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (272432) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Rage (22395, -2.96 DPS, sim-verified) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Darkspear Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Magiskull Cuffs; hands: Raider Handwraps; waist: Stormseeker's Girdle; legs: Outrider's Lizardhide Pants; feet: Kayser's Boots of Precision; finger1: Channeler's Ring; finger2: Seal of Rivendare; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: The Unstoppable Force; ranged: Totem of the Storm

No-known-source sample (15 of 2307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 0533001000000000-053030310051202051-0530000000000000)

Set DPS (verified): 926.9. Weights run: 3.7s. Verify run: 47.9s. 2307 eligible items had no known source.

Race sweep: orc 926.9 ± 1.0, windshaper-skyborne 924.7 ± 0.9, troll 923.9 ± 1.0, tauren 915.8 ± 0.9 (dps); the guide race orc is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.826 ± 0.067, crit=3.013 ± 0.097 per rating point (14 rating = 1%, 42.181 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=19.431 ± 2.186, intellect=2.315 ± 0.020, spell_power=1.062 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 168.5 attack_power points (31.66 DPS) | yes | Warlord's Mail Headguard (231657, -6.64 DPS) [vendor]; Cap of the Scarlet Savant (12752, -6.68 DPS, sim-verified) [quest]; Skyfury Helm (20134, -7.11 DPS) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 68.2 attack_power points (12.81 DPS) | yes | Lady Maye's Pendant (14558, -2.35 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -3.78 DPS) [world_drop]; Emberfury Talisman (12929, -4.88 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 91.5 attack_power points (17.19 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Flames (3475, +0.00 DPS) [world_drop]; Deathguard's Cloak (20068, -0.80 DPS) [rep]; Royal Tribunal Cloak (13376, -0.95 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 147.7 attack_power points (27.74 DPS) | yes | Savage Gladiator Chain (11726, -4.53 DPS, sim-verified) [dungeon]; Warlord's Mail Hauberk (231653, -5.07 DPS) [vendor]; Warlord's Mail Armor (16577, -5.60 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 58.9 attack_power points (11.06 DPS) | yes | Windtalker's Wristguards (19582, -0.45 DPS) [rep]; Spitfire Bracers (20481, -2.47 DPS) [crafted] |
| hands | Bloodmail Gauntlets (14615) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, +0.00 DPS) [vendor]; Raider Gloves (272099, -5.17 DPS, sim-verified) [vendor] |
| waist | Defiler's Lizardhide Girdle (20171) (or Defiler's Mail Girdle (20195)) | The Defilers [rep] | 81.5 attack_power points (15.32 DPS) | yes | Defiler's Mail Girdle (20195, +0.00 DPS) [rep]; Defiler's Chain Girdle (20150, -1.01 DPS) [rep]; Defiler's Leather Girdle (20190, -1.01 DPS) [rep] |
| legs | Outrider's Lizardhide Pants (22741) | Warsong Outriders [rep] | 129.3 attack_power points (24.29 DPS) | yes | General's Mail Leggings (231664, -0.34 DPS) [pvp]; General's Mail Legguards (231658, -0.44 DPS) [vendor]; Ironfeather Leggings (252486, -0.95 DPS) [crafted] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | 57.6 attack_power points (10.81 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -17.42 DPS, sim-verified) [dungeon] |
| finger1 | Painweaver Band (13098) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.81 DPS) [dungeon]; Seal of Rivendare (13345, -1.94 DPS) [dungeon]; Naglering (11669, -10.18 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -1.56 DPS) [dungeon]; Legionnaire's Band (19510, -1.82 DPS) [rep]; Don Mauricio's Band of Domination (22433, -14.89 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Hand of Justice (11815, -13.54 DPS, sim-verified) [dungeon] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Hand of Justice (11815, -1.88 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Darkspear Raider's Reaper (272079, -23.32 DPS, sim-verified) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Rage (22395) | Blackrock Depths: Magmus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, +0.00 DPS) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Bloodmail Gauntlets; waist: Defiler's Lizardhide Girdle; legs: Outrider's Lizardhide Pants; feet: Bloodmail Boots; finger1: Painweaver Band; finger2: Don Julio's Band; trinket1: Blackhand's Breadth; trinket2: Smolderweb's Eye; main_hand: The Unstoppable Force; ranged: Totem of Rage

No-known-source sample (15 of 2307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

