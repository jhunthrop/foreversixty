# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-053030000000000000-0000000000000000)

Set DPS (verified): 39.8. Weights run: 4.5s. Verify run: 12.6s. 260 eligible items had no known source.

Race sweep: dwarf 39.8 ± 0.0 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.429 ± 0.032, crit=1.079 ± 0.032 per rating point (14 rating = 1%, 15.112 per %), hit=1.858 ± 0.082 per rating point (10 rating = 1%, 18.584 per %), melee_haste=10.287 ± 0.300, intellect=2.789 ± 0.024, spell_power=1.106 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Trapper's Leather Hood (252505) | Leatherworking [crafted] | 22.3 attack_power points (0.77 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.29 DPS) [crafted]; Brawler's Leather Hood (252504, -0.37 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.6 attack_power points (0.29 DPS) | yes | Erudite's Amulet (277204, -0.20 DPS) [quest]; Tarnished Locket (279870, -0.50 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 30.6 attack_power points (1.05 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.52 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.81 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.90 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (39.8 DPS) | yes | Pearl-clasped Cloak (5542, -0.02 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.10 DPS) [crafted] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mystic's Wrap (14369, +0.00 DPS) [world_drop]; Mystic's Robe (14371, +0.00 DPS) [world_drop]; Totemic Leather Armor (252435, -0.65 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bright Bracers (3647, -0.10 DPS) [world_drop]; Bravo's Armbands (270015, -0.13 DPS) [quest]; Owl Bracers (4796, -0.41 DPS, sim-verified) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, -0.01 DPS) [dungeon]; Fletcher's Gloves (7348, -0.05 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.62 DPS, sim-verified) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 22.3 attack_power points (0.77 DPS) | yes | Blackened Defias Belt (10403, -0.15 DPS) [dungeon]; Stormrider's Leather Belt (252432, -0.23 DPS) [crafted]; Pristine Sash (253925, -0.23 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 36.9 attack_power points (1.27 DPS) | yes | Stormrider's Leather Pants (252502, -0.31 DPS) [crafted]; Trapper's Leather Pants (252501, -0.38 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.41 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 20.6 attack_power points (0.71 DPS) | yes | Spidersilk Boots (4320, -0.06 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.08 DPS) [crafted]; Trapper's Leather Boots (252440, -0.08 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 19.6 attack_power points (0.67 DPS) | yes | Minor Channeling Ring (1449, -0.29 DPS) [quest]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Demon Band (12054, -0.40 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.7 attack_power points (0.47 DPS) | yes | Minor Channeling Ring (1449, -0.09 DPS) [quest]; Volcanic Rock Ring (12053, -0.18 DPS) [world_drop]; Demon Band (12054, -0.20 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 314.1 attack_power points (10.77 DPS) | yes | Smite's Mighty Hammer (7230, -0.24 DPS, sim-verified) [dungeon]; Rockslicer (872, -0.83 DPS) [dungeon]; Night Reaver (1318, -1.31 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Trapper's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Armor of the Fang; wrist: Mindthrust Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Totemic Leather Pants; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (dwarf, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 103.6. Weights run: 4.5s. Verify run: 9.0s. 448 eligible items had no known source.

Race sweep: dwarf 103.6 ± 0.1 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.182 ± 0.032, crit=1.076 ± 0.037 per rating point (14 rating = 1%, 15.059 per %), hit=1.812 ± 0.071 per rating point (10 rating = 1%, 18.120 per %), melee_haste=5.337 ± 0.534, intellect=2.721 ± 0.017, spell_power=0.448 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 43.5 attack_power points (3.87 DPS) | yes | Dokebi Hat (14584, -0.50 DPS) [world_drop]; Nightsky Cowl (4039, -0.97 DPS) [world_drop]; Resilient Cap (14401, -1.00 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 28.3 attack_power points (2.52 DPS) | yes | Kaleidoscope Chain (13084, -0.65 DPS, sim-verified) [world_drop]; Seal of the Expedition (271920, -1.31 DPS) [rep]; Crystal Starfire Medallion (5003, -1.55 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 32.6 attack_power points (2.90 DPS) | yes | Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Magician's Mantle (12998, -0.52 DPS) [world_drop]; Feline Mantle (3748, -0.59 DPS) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271923, -0.25 DPS, sim-verified) [rep]; Resilient Cape (14400, -0.65 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.83 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 51.5 attack_power points (4.57 DPS) | yes | Spirewind Fetter (9406, -0.61 DPS, sim-verified) [dungeon]; Green Silk Armor (7065, -1.07 DPS) [crafted]; Beguiler Robes (7728, -1.31 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 27.8 attack_power points (2.47 DPS) | yes | Technician's Bracers (270042, -0.73 DPS, sim-verified) [quest]; Nightsky Wristbands (6407, -1.02 DPS) [world_drop]; Hawkeye's Bracers (14590, -1.13 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 26.5 attack_power points (2.35 DPS) | yes | Insignia Gloves (6408, -0.48 DPS) [world_drop]; Town Clerk's Mittens (270029, -0.54 DPS) [quest]; Resilient Handgrips (14403, -0.59 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 36.0 attack_power points (3.20 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.30 DPS) [rep]; Highlander's Mail Girdle (20120, -0.30 DPS) [vendor]; Skirmisher's Leather Belt (252461, -0.35 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 58.7 attack_power points (5.21 DPS) | yes | Totemic Leather Leggings (252458, -1.73 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, -2.23 DPS) [crafted]; Stormrider's Leather Kilt (252518, -2.27 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 24.0 attack_power points (2.13 DPS) | yes | Gilded Slippers (254001, -0.16 DPS) [crafted]; Frothing Slippers (254003, -0.44 DPS) [crafted]; Fiery Slippers (254005, -0.44 DPS) [crafted] |
| finger1 | Seal of Wrynn (2933) | An Audience with the King [quest] | 20.4 attack_power points (1.81 DPS) | yes | Protector's Band (19517, -0.12 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon]; Snake Hoop (6750, -0.37 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.5 attack_power points (1.74 DPS) | yes | Protector's Band (19517, -0.04 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.08 DPS) [dungeon]; Snake Hoop (6750, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -2.41 DPS) [dungeon]; Manual Crowd Pummeler (9449, -2.99 DPS, sim-verified) [dungeon]; Cobalt Crusher (7730, -5.13 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Sergeant Major's Cape; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger1: Seal of Wrynn; finger2: Thunderbrow Ring; main_hand: Viscous Hammer

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (dwarf, 0000000000000000-055030310051002051-0000000000000000)

Set DPS (verified): 174.8. Weights run: 4.5s. Verify run: 14.1s. 696 eligible items had no known source.

Race sweep: dwarf 174.8 ± 0.3 (dps); the guide race dwarf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.599 ± 0.065, crit=1.766 ± 0.078 per rating point (14 rating = 1%, 24.726 per %), hit=2.896 ± 0.147 per rating point (10 rating = 1%, 28.965 per %), melee_haste=6.840 ± 1.127, intellect=4.144 ± 0.054, spell_power=1.162 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 86.7 attack_power points (7.60 DPS) | yes | Papal Fez (9431, -0.61 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.87 DPS) [dungeon]; Thinking Cap (2624, -1.43 DPS) [world] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 41.4 attack_power points (3.63 DPS) | yes | Glowing Eye of Mordresh (10769, -0.01 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.40 DPS) [world]; Kaleidoscope Chain (13084, -0.92 DPS) [world_drop] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 66.3 attack_power points (5.81 DPS) | yes | Inquisitor's Shawl (19507, -0.38 DPS) [dungeon]; Windchaser Amice (14432, -0.67 DPS) [world_drop]; Green Silken Shoulders (7057, -1.00 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.08 DPS) [rep]; Blackforge Cape (6424, -1.17 DPS) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 76.6 attack_power points (6.71 DPS) | yes | Big Voodoo Robe (8200, -0.71 DPS) [crafted]; Icemail Jerkin (1981, -0.81 DPS) [world_drop]; Red Mageweave Vest (10007, -2.26 DPS, sim-verified) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 40.5 attack_power points (3.55 DPS) | yes | Glowing Magical Bracelets (13106, -0.12 DPS) [world_drop]; Warden's Wristbands (14600, -0.22 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.39 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Stormcloth Gloves (10011, -1.81 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 70.3 attack_power points (6.16 DPS) | yes | Highlander's Lizardhide Girdle (20104, -0.71 DPS) [rep]; Highlander's Mail Girdle (20119, -0.71 DPS) [vendor]; Saltscale Girdle (273023, -1.31 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 75.7 attack_power points (6.63 DPS) | yes | Red Mageweave Pants (10009, -0.85 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.10 DPS) [crafted]; Misplaced Pantaloons (276201, -1.91 DPS) [vendor] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 53.9 attack_power points (4.72 DPS) | yes | Skirmisher's Mail Boots (252564, -0.06 DPS) [crafted]; Stalker's Mail Boots (252562, -0.64 DPS) [crafted]; Briar Tredders (10582, -0.76 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 32.3 attack_power points (2.83 DPS) | yes | Ogremind Ring (1993, -0.29 DPS) [world_drop]; Mindbender Loop (5009, -0.29 DPS) [world_drop]; Protector's Band (19515, -0.31 DPS) [rep] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009)) | Bloodscalp Witch Doctor [world] | 29.0 attack_power points (2.54 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Protector's Band (19515, -0.02 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Bonebiter (6830, -1.08 DPS) [quest]; Mograine's Might (7723, -6.89 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Corpseshroud; neck: Prodigious Shadowshard Pendant; shoulder: Sheepshear Mantle; back: Sergeant Major's Cape; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Philanthropist's Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 696, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 3000000000000000-055030310051202051-0500000000000000)

Set DPS (verified): 254.0. Weights run: 4.8s. Verify run: 16.2s. 893 eligible items had no known source.

Race sweep: dwarf 254.0 ± 0.4 (dps); the guide race dwarf is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.617 ± 0.070, crit=2.282 ± 0.095 per rating point (14 rating = 1%, 31.944 per %), hit=3.742 ± 0.197 per rating point (10 rating = 1%, 37.423 per %), melee_haste=11.723 ± 1.294, intellect=4.048 ± 0.064, spell_power=1.120 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Chief Architect's Monocle (11839, +0.00 DPS) [dungeon]; Red Mageweave Headband (10033, -0.66 DPS) [crafted]; Soulcatcher Halo (10630, -0.77 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 56.7 attack_power points (5.69 DPS) | yes | Gemshard Heart (17707, -1.62 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.62 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.65 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 97.8 attack_power points (9.81 DPS) | yes | Rotgrip Mantle (17732, -1.04 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -2.44 DPS) [dungeon]; Dregmetal Spaulders (11722, -2.51 DPS) [dungeon] |
| back | Expeditionary's Cape (271921) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16336, -1.07 DPS) [pvp]; Imperial Red Cloak (8248, -1.22 DPS) [world_drop]; Runecloth Cloak (13860, -1.43 DPS) [crafted] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | 104.5 attack_power points (10.49 DPS) | yes | Acumen Robes (17775, -0.23 DPS) [quest]; Robes of Insight (940, -0.33 DPS) [world_drop]; Relentless Chain (17777, -1.99 DPS) [quest] |
| wrist | Forgotten Wraps (9433) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Runic Leather Bracers (15092, -0.17 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.39 DPS) [crafted] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rockgrip Gauntlets (17736, -1.00 DPS) [dungeon]; Sergeant Major's Mail Gauntlets (223076, -1.03 DPS) [vendor]; Prowler's Leather Gauntlets (252547, -1.16 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Sagebrush Girdle (17778, -1.08 DPS) [quest]; Highlander's Lizardhide Girdle (20103, -1.12 DPS) [rep] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | 104.5 attack_power points (10.49 DPS) | yes | Imperial Red Pants (8251, -1.96 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -2.84 DPS) [dungeon]; Ferine Leggings (6690, -3.00 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 90.2 attack_power points (9.05 DPS) | yes | Rancor Boots (11865, -2.64 DPS) [quest]; Skycaller's Leather Boots (252471, -2.90 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -2.90 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 60.7 attack_power points (6.09 DPS) | yes | Mindseye Circle (10634, -1.22 DPS) [dungeon]; Cyclopean Band (11824, -1.43 DPS) [dungeon]; Woodseed Hoop (17768, -2.44 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 attack_power points (5.76 DPS) | yes | Mindseye Circle (10634, -0.89 DPS) [dungeon]; Cyclopean Band (11824, -1.10 DPS) [dungeon]; Woodseed Hoop (17768, -2.11 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.70 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.46 DPS) [world_drop]; Soulkeeper (1607, -3.43 DPS) [world_drop]; Taran Icebreaker (2915, -3.70 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Expeditionary's Cape; chest: Knight's Mail Armor; wrist: Forgotten Wraps; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Knight's Mail Legplates; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 893, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 3230030000000000-055030310051202051-0520000000000000)

Set DPS (verified): 378.5. Weights run: 4.8s. Verify run: 55.8s. 2086 eligible items had no known source.

Race sweep: dwarf 378.5 ± 0.5 (dps); the guide race dwarf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.807 ± 0.082, crit=2.997 ± 0.122 per rating point (14 rating = 1%, 41.962 per %), hit=4.194 ± 0.247 per rating point (10 rating = 1%, 41.939 per %), melee_haste=15.818 ± 1.714, intellect=5.241 ± 0.124, spell_power=0.970 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 279.1 attack_power points (24.80 DPS) | yes | Cap of the Scarlet Savant (12752, -8.12 DPS, sim-verified) [quest]; Magister's Crown (16686, -9.88 DPS) [dungeon]; Blue Dragonscale Helm (252604, -10.53 DPS) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 99.6 attack_power points (8.85 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.47 DPS) [world_drop]; Diana's Pearl Necklace (22403, -0.62 DPS) [dungeon]; Star of Mystaria (12103, -0.93 DPS) [dungeon] |
| shoulder | Theramore Shoulderguards (272957) | Theramore Expeditionary Force [rep] | 157.0 attack_power points (13.96 DPS) | yes | Theramore Shoulderpads (271953, -0.34 DPS) [rep]; Theramore Shoulders (271954, -0.34 DPS) [rep]; Shroud of the Nathrezim (18720, -2.17 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Shroud of Arcane Mastery (22330, -1.08 DPS, sim-verified) [dungeon]; Royal Tribunal Cloak (13376, -1.38 DPS) [dungeon]; Expeditionary's Cape (271907, -1.38 DPS) [rep] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 158.4 attack_power points (14.08 DPS) | yes | Blue Dragonscale Breastplate (15048, -0.34 DPS) [crafted]; Dawn Armor (252483, -0.68 DPS) [crafted]; Savage Gladiator Chain (11726, -1.53 DPS) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Windtalker's Wristguards (19582, -0.40 DPS) [rep]; Manacle Cuffs (11962, -0.52 DPS) [quest]; Magiskull Cuffs (13107, -1.01 DPS, sim-verified) [world_drop] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Theramore Handguards (271950, -2.10 DPS) [rep]; Bloodmail Gauntlets (14615, -3.28 DPS) [dungeon]; Mourning Handguards (279257, -3.37 DPS) [crafted] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Mail Girdle (20044, -0.91 DPS) [vendor]; Highlander's Lizardhide Girdle (20046, -0.91 DPS) [rep]; Clutch of Andros (13956, -1.60 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Lizardhide Pants (22750) | Silverwing Sentinels [rep] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Dragonscale Leggings (252603, -0.08 DPS) [crafted]; Ironfeather Leggings (252486, -1.32 DPS) [crafted]; Sentinel's Chain Leggings (237819, -2.28 DPS, sim-verified) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Dragonrider Boots (18102, -0.33 DPS) [dungeon]; The Postmaster's Treads (13391, -1.74 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -2.33 DPS) [dungeon]; Brainlash (6440, -3.26 DPS) [dungeon]; Naglering (11669, -9.56 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.78 DPS) [dungeon]; Brainlash (6440, -1.71 DPS) [dungeon]; Cutthroat's Signet (272408, -2.73 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+11.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -7.16 DPS, sim-verified) [vendor] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -2.04 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (272432) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -3.40 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Theramore Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Tranquil Wristguards; hands: Theramore Handwraps; waist: Stormseeker's Girdle; legs: Sentinel's Lizardhide Pants; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Forest Defender's Axe; ranged: Totem of the Storm

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (dwarf, 0533001000000000-053030310051202051-0530000000000000)

Set DPS (verified): 922.3. Weights run: 4.3s. Verify run: 54.3s. 2086 eligible items had no known source.

Race sweep: dwarf 922.3 ± 0.9 (dps); the guide race dwarf is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.799 ± 0.071, crit=3.079 ± 0.109 per rating point (14 rating = 1%, 43.100 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=20.283 ± 2.270, intellect=2.342 ± 0.024, spell_power=1.079 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 169.9 attack_power points (31.32 DPS) | yes | Cap of the Scarlet Savant (12752, -7.40 DPS, sim-verified) [quest]; Eye of Rend (12587, -10.15 DPS) [dungeon]; Outlaw's Collar (279253, -12.43 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emberfury Talisman (12929, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Mark of Fordring (15411, +0.00 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 91.9 attack_power points (16.94 DPS) | yes | Theramore Pauldrons (271955, -0.04 DPS) [rep]; Theramore Epaulets (271956, -0.04 DPS) [rep]; Shroud of the Nathrezim (18720, -0.69 DPS) [dungeon] |
| back | Sergeant Major's Cape (16337) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.69 DPS) [dungeon]; Fel Cape (279269, -1.41 DPS) [crafted]; Cloak of the Honor Guard (20073, -1.43 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 148.1 attack_power points (27.30 DPS) | yes | Savage Gladiator Chain (11726, -0.28 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -6.14 DPS) [quest]; Dawn Armor (252483, -7.09 DPS) [crafted] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 58.6 attack_power points (10.81 DPS) | yes | Forest Stalker's Bracers (19587, -0.04 DPS) [rep]; Windtalker's Wristguards (19582, -0.35 DPS) [rep] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gargoyle Slashers (13957, +0.00 DPS) [dungeon]; Bloodmail Gauntlets (14615, +0.00 DPS) [dungeon]; Sergeant Major's Mail Gauntlets (223076, -0.07 DPS) [vendor] |
| waist | Highlander's Mail Girdle (20044) (or Highlander's Lizardhide Girdle (20046)) | Samuel Hawke [vendor] | 82.9 attack_power points (15.29 DPS) | yes | Highlander's Lizardhide Girdle (20046, +0.00 DPS) [rep]; Highlander's Chain Girdle (20043, -1.07 DPS) [rep]; Highlander's Leather Girdle (20045, -1.07 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 149.2 attack_power points (27.50 DPS) | yes | Sentinel's Lizardhide Pants (237817, -0.93 DPS) [vendor]; Sentinel's Leather Pants (237818, -2.65 DPS) [vendor]; Sentinel's Silk Leggings (237815, -12.31 DPS, sim-verified) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | 59.4 attack_power points (10.95 DPS) | yes | Timmy's Galoshes (13402, -1.00 DPS) [dungeon]; Scalegut Treaders (275618, -1.29 DPS) [crafted]; Dragonrider Boots (18102, -9.34 DPS, sim-verified) [dungeon] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magni's Will (12548, +0.00 DPS) [quest]; Painweaver Band (13098, +0.00 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -0.39 DPS) [dungeon]; Magni's Will (12548, -8.54 DPS, sim-verified) [quest] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+9.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, -1.30 DPS) [dungeon]; Uther's Strength (11302, -2.49 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.59 DPS) [vendor]; Burst of Knowledge (11832, -1.99 DPS) [dungeon]; Uther's Strength (11302, -3.18 DPS) [world_drop] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (+11.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Forest Defender's Axe (279260, -11.54 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (272432) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -10.04 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Sergeant Major's Cape; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; hands: Theramore Gloves; waist: Highlander's Mail Girdle; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Protector's Band; finger2: Don Julio's Band; trinket1: Hand of Justice; trinket2: Weakness Analyzer; main_hand: The Unstoppable Force; ranged: Totem of the Storm

No-known-source sample (15 of 2086, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

## Horde

### Band 20 (orc, 0000000000000000-053030000000000000-0000000000000000)

Set DPS (verified): 40.1. Weights run: 4.5s. Verify run: 14.9s. 249 eligible items had no known source.

Race sweep: windshaper-skyborne 40.5 ± 0.0, orc 40.1 ± 0.0, troll 40.1 ± 0.0, tauren 39.7 ± 0.0 (dps); windshaper-skyborne beats the guide race orc by 0.4 (1.00%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.977 ± 0.023, crit=0.800 ± 0.023 per rating point (14 rating = 1%, 11.197 per %), hit=1.401 ± 0.063 per rating point (10 rating = 1%, 14.006 per %), melee_haste=6.879 ± 0.218, intellect=2.689 ± 0.023, spell_power=1.414 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Trapper's Leather Hood (252505) | Leatherworking [crafted] | 21.5 attack_power points (0.75 DPS) | yes | Defender's Leather Hood (252447, -0.19 DPS) [crafted]; Shadow Goggles (4373, -0.28 DPS) [crafted]; Totemic Leather Hood (252448, -0.31 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 5.9 attack_power points (0.20 DPS) | yes | Erudite's Amulet (277204, -0.14 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 31.3 attack_power points (1.09 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.39 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.89 DPS) [crafted]; Serpent's Shoulders (5404, -0.92 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-verified (40.1 DPS) | yes | Sanguine Cape (14376, -0.00 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.10 DPS) [world_drop]; Dark Leather Cloak (2316, -0.14 DPS) [crafted] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mystic's Wrap (14369, +0.00 DPS) [world_drop]; Stormrider's Leather Armor (252492, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.47 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bright Bracers (3647, -0.09 DPS) [world_drop]; Featherbead Bracers (15452, -0.19 DPS) [quest]; Owl Bracers (4796, -0.22 DPS, sim-verified) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Totemic Leather Gloves (252497, +0.00 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.78 DPS, sim-verified) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 21.5 attack_power points (0.75 DPS) | yes | Blackened Defias Belt (10403, -0.12 DPS) [dungeon]; Stormrider's Leather Belt (252432, -0.18 DPS) [crafted]; Pristine Sash (253925, -0.18 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 38.7 attack_power points (1.35 DPS) | yes | Trapper's Leather Pants (252501, -0.25 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.44 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 21.9 attack_power points (0.77 DPS) | yes | Spidersilk Boots (4320, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.10 DPS) [crafted]; Totemic Leather Boots (252442, -0.12 DPS) [crafted] |
| finger1 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 19.4 attack_power points (0.68 DPS) | yes | Legionnaire's Band (20429, -0.26 DPS) [rep]; Volcanic Rock Ring (12053, -0.40 DPS) [world_drop]; Demon Band (12054, -0.40 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 18.1 attack_power points (0.63 DPS) | yes | Legionnaire's Band (20429, -0.22 DPS) [rep]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop]; Demon Band (12054, -0.35 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 313.1 attack_power points (10.93 DPS) | yes | Smite's Mighty Hammer (7230, -0.39 DPS) [dungeon]; Forsaken Greataxe (251533, -0.41 DPS, sim-verified) [quest]; Hammerbone (270018, -0.45 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Trapper's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Armor of the Fang; wrist: Mindthrust Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Totemic Leather Pants; feet: Stormrider's Leather Boots; finger1: Loop of Sacrifice; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads

### Band 30 (orc, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 103.9. Weights run: 4.5s. Verify run: 10.0s. 442 eligible items had no known source.

Race sweep: windshaper-skyborne 104.8 ± 0.1, orc 103.9 ± 0.1, troll 103.7 ± 0.1, tauren 102.8 ± 0.1 (dps); windshaper-skyborne beats the guide race orc by 0.9 (0.85%) ± 0.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.182 ± 0.032, crit=1.076 ± 0.037 per rating point (14 rating = 1%, 15.059 per %), hit=1.812 ± 0.071 per rating point (10 rating = 1%, 18.120 per %), melee_haste=5.337 ± 0.534, intellect=2.721 ± 0.017, spell_power=0.448 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 43.5 attack_power points (3.87 DPS) | yes | Dokebi Hat (14584, -0.50 DPS) [world_drop]; Nightsky Cowl (4039, -0.97 DPS) [world_drop]; Resilient Cap (14401, -1.00 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 28.3 attack_power points (2.52 DPS) | yes | Kaleidoscope Chain (13084, -0.42 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.31 DPS) [vendor]; Crystal Starfire Medallion (5003, -1.55 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 32.6 attack_power points (2.90 DPS) | yes | Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Magician's Mantle (12998, -0.52 DPS) [world_drop]; Feline Mantle (3748, -0.59 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resilient Cape (14400, -0.48 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.67 DPS) [world_drop]; Soft Willow Cape (16661, -0.73 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 51.5 attack_power points (4.57 DPS) | yes | Spirewind Fetter (9406, -0.43 DPS) [dungeon]; Green Silk Armor (7065, -1.07 DPS) [crafted]; Beguiler Robes (7728, -1.31 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 27.8 attack_power points (2.47 DPS) | yes | Technician's Bracers (270042, -0.79 DPS) [quest]; Nightsky Wristbands (6407, -1.02 DPS) [world_drop]; Hawkeye's Bracers (14590, -1.13 DPS) [world_drop] |
| hands | Gloves of Kapelan (6744) | Alliance Relations [quest] | 30.3 attack_power points (2.69 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.34 DPS) [dungeon]; Insignia Gloves (6408, -0.82 DPS) [world_drop]; Resilient Handgrips (14403, -0.93 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 36.0 attack_power points (3.20 DPS) | yes | Defiler's Lizardhide Girdle (20172, -0.30 DPS) [rep]; Defiler's Mail Girdle (20197, -0.30 DPS) [rep]; Skirmisher's Leather Belt (252461, -0.35 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 58.7 attack_power points (5.21 DPS) | yes | Totemic Leather Leggings (252458, -1.66 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, -2.23 DPS) [crafted]; Stormrider's Leather Kilt (252518, -2.27 DPS) [crafted] |
| feet | Acidic Walkers (9454) | Gnomeregan: Viscous Fallout [dungeon] | 24.0 attack_power points (2.13 DPS) | yes | Gilded Slippers (254001, -0.16 DPS) [crafted]; Frothing Slippers (254003, -0.44 DPS) [crafted]; Fiery Slippers (254005, -0.44 DPS) [crafted] |
| finger1 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 19.6 attack_power points (1.74 DPS) | yes | Legionnaire's Band (19513, -0.05 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.08 DPS) [dungeon]; Snake Hoop (6750, -0.29 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.5 attack_power points (1.74 DPS) | yes | Legionnaire's Band (19513, -0.04 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.08 DPS) [dungeon]; Snake Hoop (6750, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -2.39 DPS, sim-verified) [dungeon]; Corpsemaker (6687, -2.41 DPS) [dungeon]; Cobalt Crusher (7730, -5.13 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Darkspear Raider's Cloak; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Gloves of Kapelan; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Acidic Walkers; finger2: Thunderbrow Ring; main_hand: Viscous Hammer

No-known-source sample (15 of 442, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (orc, 0000000000000000-055030310051002051-0000000000000000)

Set DPS (verified): 174.2. Weights run: 4.5s. Verify run: 13.9s. 679 eligible items had no known source.

Race sweep: windshaper-skyborne 174.9 ± 0.3, orc 174.2 ± 0.3, troll 173.5 ± 0.3, tauren 171.5 ± 0.3 (dps); windshaper-skyborne beats the guide race orc by 0.7 (0.40%) ± 0.4, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.599 ± 0.065, crit=1.766 ± 0.078 per rating point (14 rating = 1%, 24.726 per %), hit=2.896 ± 0.147 per rating point (10 rating = 1%, 28.965 per %), melee_haste=6.840 ± 1.127, intellect=4.144 ± 0.054, spell_power=1.162 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpseshroud (10574, +0.00 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.25 DPS) [dungeon]; Thinking Cap (2624, -0.81 DPS) [world] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 41.4 attack_power points (3.63 DPS) | yes | Glowing Eye of Mordresh (10769, -0.01 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.40 DPS) [world]; Kaleidoscope Chain (13084, -0.92 DPS) [world_drop] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 66.3 attack_power points (5.81 DPS) | yes | Inquisitor's Shawl (19507, -0.38 DPS) [dungeon]; Windchaser Amice (14432, -0.67 DPS) [world_drop]; Green Silken Shoulders (7057, -1.00 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackforge Cape (6424, -1.09 DPS) [dungeon]; Glyphed Cloak (4732, -1.45 DPS) [world_drop]; Long Silken Cloak (4326, -1.57 DPS) [crafted] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 76.6 attack_power points (6.71 DPS) | yes | Big Voodoo Robe (8200, -0.71 DPS) [crafted]; Icemail Jerkin (1981, -0.81 DPS) [world_drop]; Red Mageweave Vest (10007, -1.20 DPS, sim-verified) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 40.5 attack_power points (3.55 DPS) | yes | Glowing Magical Bracelets (13106, -0.12 DPS) [world_drop]; Warden's Wristbands (14600, -0.22 DPS) [world_drop]; Radiant Silver Bracers (4545, -0.24 DPS) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Gloves (10011, +0.00 DPS) [crafted]; Red Mageweave Gloves (10018, +0.00 DPS) [crafted]; Rustler Gloves (9704, -2.79 DPS, sim-verified) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 70.3 attack_power points (6.16 DPS) | yes | Defiler's Lizardhide Girdle (20173, -0.71 DPS) [rep]; Defiler's Mail Girdle (20198, -0.71 DPS) [rep]; Saltscale Girdle (273023, -1.31 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 75.7 attack_power points (6.63 DPS) | yes | Red Mageweave Pants (10009, -0.85 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.10 DPS) [crafted]; Misplaced Pantaloons (276201, -1.91 DPS) [vendor] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 53.9 attack_power points (4.72 DPS) | yes | Skirmisher's Mail Boots (252564, -0.06 DPS) [crafted]; Stalker's Mail Boots (252562, -0.64 DPS) [crafted]; Briar Tredders (10582, -0.76 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 32.3 attack_power points (2.83 DPS) | yes | Ogremind Ring (1993, -0.29 DPS) [world_drop]; Mindbender Loop (5009, -0.29 DPS) [world_drop]; Legionnaire's Band (19512, -0.31 DPS) [rep] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009)) | Bloodscalp Witch Doctor [world] | 29.0 attack_power points (2.54 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Legionnaire's Band (19512, -0.02 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Pendulum of Doom (9425, -6.04 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Prodigious Shadowshard Pendant; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Philanthropist's Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 3000000000000000-055030310051202051-0500000000000000)

Set DPS (verified): 251.9. Weights run: 4.8s. Verify run: 21.0s. 862 eligible items had no known source.

Race sweep: windshaper-skyborne 253.4 ± 0.4, orc 251.9 ± 0.4, troll 251.7 ± 0.3, tauren 249.3 ± 0.4 (dps); windshaper-skyborne beats the guide race orc by 1.5 (0.58%) ± 0.5, beyond error.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.274 ± 0.050, crit=1.758 ± 0.068 per rating point (14 rating = 1%, 24.616 per %), hit=2.785 ± 0.146 per rating point (10 rating = 1%, 27.849 per %), melee_haste=9.977 ± 0.980, intellect=4.558 ± 0.075, spell_power=1.001 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulcatcher Halo (10630, +0.00 DPS) [dungeon]; Chief Architect's Monocle (11839, +0.00 DPS) [dungeon]; Braincage (12549, +0.00 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 63.8 attack_power points (5.55 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.49 DPS, sim-verified) [quest]; Gemshard Heart (17707, -1.59 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.77 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 106.2 attack_power points (9.24 DPS) | yes | Rotgrip Mantle (17732, -0.97 DPS) [dungeon]; Dregmetal Spaulders (11722, -2.25 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -2.38 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.94 DPS) [quest]; Imperial Red Cloak (8248, -1.19 DPS) [world_drop]; Big Voodoo Cloak (8216, -1.55 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Robes of Insight (940, -0.53 DPS, sim-verified) [world_drop]; Stone Guard's Mail Armor (220826, -0.60 DPS) [vendor]; Stone Guard's Pulsing Breastplate (220844, -0.84 DPS) [vendor] |
| wrist | Forgotten Wraps (9433) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Runic Leather Bracers (15092, -0.25 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.33 DPS) [quest] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Rockgrip Gauntlets (17736, -0.64 DPS) [dungeon]; Prowler's Leather Gauntlets (252547, -1.23 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 92.6 attack_power points (8.06 DPS) | yes | Ban'thok Sash (11662, -0.23 DPS) [dungeon]; Serenity Belt (13144, -1.32 DPS) [world_drop]; Imperial Red Sash (8253, -1.41 DPS) [world_drop] |
| legs | Stone Guard's Mail Legplates (220834) | Lady Palanseer [vendor] | 103.3 attack_power points (8.99 DPS) | yes | Stone Guard's Pulsing Legplates (220847, -0.24 DPS) [vendor]; Imperial Red Pants (8251, -1.38 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -1.59 DPS) [dungeon] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 84.4 attack_power points (7.35 DPS) | yes | Skycaller's Leather Boots (252471, -1.68 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.68 DPS) [crafted]; Rancor Boots (11865, -5.26 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, +0.00 DPS) [dungeon]; Mindseye Circle (10634, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.39 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.66 DPS) [world_drop]; Molten Heart of the Mountain (249470, -3.04 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, -1.13 DPS) [world_drop]; Wildstaff (20556, -1.61 DPS, sim-verified) [quest]; Soulkeeper (1607, -2.25 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Acumen Robes; wrist: Forgotten Wraps; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 862, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 3230030000000000-055030310051202051-0520000000000000)

Set DPS (verified): 378.7. Weights run: 4.8s. Verify run: 56.1s. 2043 eligible items had no known source.

Race sweep: windshaper-skyborne 379.7 ± 0.5, orc 378.7 ± 0.5, troll 378.2 ± 0.5, tauren 375.7 ± 0.5 (dps); windshaper-skyborne beats the guide race orc by 1.0 (0.26%) ± 0.6, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.807 ± 0.082, crit=2.997 ± 0.122 per rating point (14 rating = 1%, 41.962 per %), hit=4.194 ± 0.247 per rating point (10 rating = 1%, 41.939 per %), melee_haste=15.818 ± 1.714, intellect=5.241 ± 0.124, spell_power=0.970 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 279.1 attack_power points (24.80 DPS) | yes | Warlord's Mail Headguard (231657, -6.47 DPS) [vendor]; Cap of the Scarlet Savant (12752, -7.17 DPS, sim-verified) [quest]; Warlord's Mail Helm (231663, -9.77 DPS) [pvp] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 99.6 attack_power points (8.85 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.47 DPS) [world_drop]; Diana's Pearl Necklace (22403, -0.62 DPS) [dungeon]; Star of Mystaria (12103, -0.93 DPS) [dungeon] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 157.0 attack_power points (13.96 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.34 DPS) [vendor]; Darkspear Shoulders (272104, -0.34 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Tribunal Cloak (13376, -1.38 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -1.38 DPS) [vendor]; Shroud of Arcane Mastery (22330, -1.45 DPS, sim-verified) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 158.4 attack_power points (14.08 DPS) | yes | Blue Dragonscale Breastplate (15048, +0.00 DPS) [crafted]; Warlord's Mail Armor (16577, +0.00 DPS) [vendor]; Warlord's Mail Hauberk (231653, +0.00 DPS) [vendor] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Windtalker's Wristguards (19582, -0.40 DPS) [rep]; Manacle Cuffs (11962, -0.52 DPS) [quest]; Magiskull Cuffs (13107, -1.48 DPS, sim-verified) [world_drop] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Vices (231655, -0.17 DPS) [vendor]; General's Mail Gauntlets (231660, -0.74 DPS) [pvp]; General's Mail Gloves (231666, -1.34 DPS) [vendor] |
| waist | Stormseeker's Girdle (272399) | Pix Xizzix [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Waistband (16575, -0.03 DPS) [pvp]; Defiler's Lizardhide Girdle (20171, -0.91 DPS) [rep]; Clutch of Andros (13956, -2.18 DPS, sim-verified) [dungeon] |
| legs | Outrider's Lizardhide Pants (22741) | Warsong Outriders [rep] | 188.0 attack_power points (16.71 DPS) | yes | General's Mail Legguards (231658, +0.00 DPS) [vendor]; General's Mail Leggings (231664, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, -0.08 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; General's Mail Greaves (231656, +0.00 DPS) [vendor]; General's Mail Sabatons (231661, +0.00 DPS) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -2.33 DPS) [dungeon]; Brainlash (6440, -3.26 DPS) [dungeon]; Naglering (11669, -8.99 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brainlash (6440, -1.71 DPS) [dungeon]; Seal of Rivendare (13345, -2.32 DPS, sim-verified) [dungeon]; Cutthroat's Signet (272408, -2.73 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+8.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -3.01 DPS, sim-verified) [dungeon] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -1.39 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (272432) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of Thunder (228176, -3.56 DPS, sim-verified) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Lady Maye's Pendant; shoulder: Darkspear Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Tranquil Wristguards; hands: Raider Handwraps; waist: Stormseeker's Girdle; legs: Outrider's Lizardhide Pants; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Forest Defender's Axe; ranged: Totem of the Storm

No-known-source sample (15 of 2043, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60, raid preset (orc, 0533001000000000-053030310051202051-0530000000000000)

Set DPS (verified): 906.7. Weights run: 4.3s. Verify run: 54.0s. 2043 eligible items had no known source.

Race sweep: orc 906.7 ± 0.9, troll 905.0 ± 0.9, windshaper-skyborne 903.3 ± 0.9, tauren 895.4 ± 0.9 (dps); the guide race orc is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.799 ± 0.071, crit=3.079 ± 0.109 per rating point (14 rating = 1%, 43.100 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=20.283 ± 2.270, intellect=2.342 ± 0.024, spell_power=1.079 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 169.9 attack_power points (31.32 DPS) | yes | Skyfury Helm (20134, -6.36 DPS, sim-verified) [quest]; Warlord's Mail Headguard (231657, -6.47 DPS) [vendor]; Cap of the Scarlet Savant (12752, -6.79 DPS) [quest] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emberfury Talisman (12929, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Mark of Fordring (15411, +0.00 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 91.9 attack_power points (16.94 DPS) | yes | Warlord's Mail Pauldrons (231654, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Flames (3475, +0.00 DPS) [world_drop]; Fel Cape (279269, -0.72 DPS) [crafted]; Deathguard's Cloak (20068, -0.74 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 148.1 attack_power points (27.30 DPS) | yes | Warlord's Mail Hauberk (231653, -4.82 DPS) [vendor]; Warlord's Mail Armor (16577, -5.28 DPS) [vendor]; Savage Gladiator Chain (11726, -6.86 DPS, sim-verified) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 58.6 attack_power points (10.81 DPS) | yes | Forest Stalker's Bracers (19587, -0.04 DPS) [rep]; Windtalker's Wristguards (19582, -0.35 DPS) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Vices (231655, +0.00 DPS) [vendor]; General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, +0.00 DPS) [vendor] |
| waist | Defiler's Lizardhide Girdle (20171) (or Defiler's Mail Girdle (20195)) | The Defilers [rep] | 82.9 attack_power points (15.29 DPS) | yes | Defiler's Mail Girdle (20195, +0.00 DPS) [rep]; Defiler's Chain Girdle (20150, -1.07 DPS) [rep]; Defiler's Leather Girdle (20190, -1.07 DPS) [rep] |
| legs | Outrider's Lizardhide Pants (22741) | Warsong Outriders [rep] | 129.8 attack_power points (23.93 DPS) | yes | General's Mail Leggings (231664, -0.05 DPS) [pvp]; General's Mail Legguards (231658, -0.30 DPS) [vendor]; Ironfeather Leggings (252486, -0.65 DPS) [crafted] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | 59.4 attack_power points (10.95 DPS) | yes | General's Mail Greaves (231656, +0.00 DPS) [vendor]; General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -17.62 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Band of Rumination (18103, -0.57 DPS) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Band of Rumination (18103, -2.58 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -14.84 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (+8.2 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Forest Defender's Axe (279260, -8.24 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Totem (272433, +0.00 DPS) [vendor] |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; hands: Raider Gloves; waist: Defiler's Lizardhide Girdle; legs: Outrider's Lizardhide Pants; feet: Bloodmail Boots; finger1: Legionnaire's Band; finger2: Don Julio's Band; main_hand: The Unstoppable Force; ranged: Totem of Thunder

No-known-source sample (15 of 2043, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

