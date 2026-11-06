# Leveling BiS: Enhancement

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-055010000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 3.6s. Verify run: 9.0s. 260 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.860 ± 0.020, crit=0.665 ± 0.020 per rating point (14 rating = 1%, 9.307 per %), hit=1.281 ± 0.052 per rating point (10 rating = 1%, 12.808 per %), melee_haste=6.212 ± 0.202, intellect=1.590 ± 0.024, spell_power=1.218 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.58 DPS) | yes | Trapper's Leather Hood (252505, -0.12 DPS) [crafted]; Totemic Leather Hood (252448, -0.18 DPS) [crafted]; Shadow Goggles (4373, -0.29 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 5.2 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS) [quest]; Tarnished Locket (279870, -0.40 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 20.4 attack_power points (0.73 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.25 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.56 DPS) [crafted]; Serpent's Shoulders (5404, -0.58 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pearl-clasped Cloak (5542, -0.03 DPS) [crafted]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Sanguine Cape (14376, -0.06 DPS) [world_drop] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 26.1 attack_power points (0.94 DPS) | yes | Stormrider's Leather Armor (252492, -0.32 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.34 DPS) [crafted]; Brawler's Leather Armor (252490, -0.36 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 8.6 attack_power points (0.31 DPS) | yes | Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Owl Bracers (4796, -0.02 DPS) [vendor]; Bristlebark Bindings (14569, -0.07 DPS) [world_drop] |
| hands | Totemic Leather Gloves (252497) | Leatherworking [crafted] | 14.1 attack_power points (0.51 DPS) | yes | Gold-flecked Gloves (5195, -0.00 DPS) [dungeon]; Serpent Gloves (5970, -0.01 DPS) [dungeon]; Gloves of the Fang (10413, -0.03 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Totemic Leather Belt (252431, -0.18 DPS) [crafted]; Keller's Girdle (2911, -0.19 DPS) [world_drop]; Brawler's Leather Belt (252428, -0.24 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 32.9 attack_power points (1.18 DPS) | yes | Trapper's Leather Pants (252501, -0.22 DPS, sim-verified) [crafted]; Stormrider's Leather Pants (252502, -0.40 DPS) [crafted]; Defender's Leather Pants (252445, -0.41 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 17.3 attack_power points (0.62 DPS) | yes | Stormrider's Leather Boots (252443, -0.07 DPS) [crafted]; Spidersilk Boots (4320, -0.09 DPS) [crafted]; Brawler's Leather Boots (252439, -0.11 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.4 attack_power points (0.41 DPS) | yes | Minor Channeling Ring (1449, -0.08 DPS) [quest]; Demon Band (12054, -0.12 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.19 DPS) [rep] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.3 attack_power points (0.41 DPS) | yes | Minor Channeling Ring (1449, -0.07 DPS) [quest]; Demon Band (12054, -0.12 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.19 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 302.1 attack_power points (10.88 DPS) | yes | Smite's Mighty Hammer (7230, -0.02 DPS) [dungeon]; Rockslicer (872, -0.44 DPS) [dungeon]; Night Reaver (1318, -0.95 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Lambent Scale Cloak; chest: Totemic Leather Armor; wrist: Bravo's Armbands; hands: Totemic Leather Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Protector's Band; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 260, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (dwarf, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 103.6. Weights run: 3.8s. Verify run: 7.6s. 448 eligible items had no known source.

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

### Band 40 (dwarf, 0000000000000000-055030310051202031-0000000000000000)

Set DPS (verified): 156.8. Weights run: 3.9s. Verify run: 13.8s. 696 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.392 ± 0.045, crit=1.528 ± 0.057 per rating point (14 rating = 1%, 21.388 per %), hit=2.291 ± 0.100 per rating point (10 rating = 1%, 22.907 per %), melee_haste=7.430 ± 0.822, intellect=3.548 ± 0.027, spell_power=0.375 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Miner's Hat of the Deep (9429, +0.00 DPS) [dungeon]; Corpseshroud (10574, +0.00 DPS) [dungeon]; Thinking Cap (2624, -0.28 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prodigious Shadowshard Pendant (17773, +0.00 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.21 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.52 DPS) [world_drop] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 56.8 attack_power points (5.32 DPS) | yes | Windchaser Amice (14432, -0.61 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.75 DPS) [dungeon]; Green Silken Shoulders (7057, -1.38 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.25 DPS) [rep]; Blackforge Cape (6424, -1.24 DPS) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 65.8 attack_power points (6.17 DPS) | yes | Silksand Tunic (14417, -0.84 DPS) [world_drop]; Silksand Wraps (14425, -0.84 DPS) [world_drop]; Red Mageweave Vest (10007, -1.66 DPS, sim-verified) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 34.7 attack_power points (3.26 DPS) | yes | Glowing Magical Bracelets (13106, -0.03 DPS) [world_drop]; Warden's Wristbands (14600, -0.20 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Stormcloth Gloves (10011, +0.00 DPS) [crafted]; Engineering Gloves (1659, -2.54 DPS, sim-verified) [world] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 55.8 attack_power points (5.24 DPS) | yes | Highlander's Lizardhide Girdle (20104, -0.25 DPS) [rep]; Highlander's Mail Girdle (20119, -0.25 DPS) [vendor]; Saltscale Girdle (273023, -0.55 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 68.6 attack_power points (6.43 DPS) | yes | Crimson Silk Pantaloons (7062, -1.44 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -1.95 DPS) [crafted]; Misplaced Pantaloons (276201, -2.11 DPS) [vendor] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 49.5 attack_power points (4.64 DPS) | yes | Skirmisher's Mail Boots (252564, -0.60 DPS) [crafted]; Stalker's Mail Boots (252562, -0.88 DPS) [crafted]; Briar Tredders (10582, -0.99 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 27.1 attack_power points (2.55 DPS) | yes | Ogremind Ring (1993, -0.22 DPS) [world_drop]; Mindbender Loop (5009, -0.22 DPS) [world_drop]; Seal of Wrynn (2933, -0.26 DPS) [quest] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009)) | Bloodscalp Witch Doctor [world] | 24.8 attack_power points (2.33 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Seal of Wrynn (2933, -0.04 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Bonebiter (6830, -1.15 DPS) [quest]; Pendulum of Doom (9425, -6.04 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; shoulder: Sheepshear Mantle; back: Sergeant Major's Cape; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 696, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (dwarf, 3230000000000000-055030310051202051-0000000000000000)

Set DPS (verified): 241.3. Weights run: 3.9s. Verify run: 12.5s. 893 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.498 ± 0.080, crit=2.277 ± 0.113 per rating point (14 rating = 1%, 31.872 per %), hit=4.290 ± 0.235 per rating point (10 rating = 1%, 42.905 per %), melee_haste=13.063 ± 1.631, intellect=4.221 ± 0.082, spell_power=1.087 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Chief Architect's Monocle (11839, +0.00 DPS) [dungeon]; Soulcatcher Halo (10630, -0.63 DPS) [dungeon]; Braincage (12549, -0.63 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 59.1 attack_power points (5.65 DPS) | yes | Gemshard Heart (17707, -1.61 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.69 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.49 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 100.7 attack_power points (9.63 DPS) | yes | Rotgrip Mantle (17732, -2.01 DPS, sim-verified) [dungeon]; Lead Surveyor's Mantle (11842, -2.42 DPS) [dungeon]; Dregmetal Spaulders (11722, -2.43 DPS) [dungeon] |
| back | Expeditionary's Cape (271921) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Red Cloak (8248, -1.21 DPS) [world_drop]; Sergeant Major's Cape (16336, -1.22 DPS) [pvp]; Runecloth Cloak (13860, -1.49 DPS) [crafted] |
| chest | Knight's Mail Armor (223078) | Captain Dirgehammer [vendor] | 106.5 attack_power points (10.18 DPS) | yes | Robes of Insight (940, -0.10 DPS) [world_drop]; Acumen Robes (17775, -0.14 DPS) [quest]; Relentless Chain (17777, -1.86 DPS) [quest] |
| wrist | Forgotten Wraps (9433) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Runic Leather Bracers (15092, -0.20 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.43 DPS) [quest] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Mail Gauntlets (223076, -0.99 DPS) [vendor]; Prowler's Leather Gauntlets (252547, -1.20 DPS) [crafted]; Rockgrip Gauntlets (17736, -2.23 DPS, sim-verified) [dungeon] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Highlander's Lizardhide Girdle (20103, -1.21 DPS) [rep]; Highlander's Mail Girdle (20118, -1.21 DPS) [vendor] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | 106.5 attack_power points (10.18 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -2.61 DPS) [dungeon]; Ferine Leggings (6690, -2.86 DPS) [dungeon]; Imperial Red Pants (8251, -3.28 DPS, sim-verified) [world_drop] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 97.1 attack_power points (9.28 DPS) | yes | Skycaller's Leather Boots (252471, -3.28 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -3.28 DPS) [crafted]; Rancor Boots (11865, -3.58 DPS, sim-verified) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 63.3 attack_power points (6.05 DPS) | yes | Mindseye Circle (10634, -1.21 DPS) [dungeon]; Cyclopean Band (11824, -1.53 DPS) [dungeon]; Woodseed Hoop (17768, -2.42 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 62.9 attack_power points (6.01 DPS) | yes | Mindseye Circle (10634, -1.17 DPS) [dungeon]; Cyclopean Band (11824, -1.49 DPS) [dungeon]; Woodseed Hoop (17768, -2.38 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -1.73 DPS, sim-verified) [quest]; Uther's Strength (11302, -3.07 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -2.26 DPS, sim-verified) [world_drop]; Soulkeeper (1607, -3.01 DPS) [world_drop]; Taran Icebreaker (2915, -3.46 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Expeditionary's Cape; chest: Knight's Mail Armor; wrist: Forgotten Wraps; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Knight's Mail Legplates; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 893, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (dwarf, 3230030000000000-055030310051202051-0520000000000000)

Set DPS (verified): 378.5. Weights run: 4.0s. Verify run: 48.1s. 2086 eligible items had no known source.

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

Set DPS (verified): 922.3. Weights run: 3.5s. Verify run: 45.2s. 2086 eligible items had no known source.

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

### Band 20 (orc, 0000000000000000-055010000000000000-0000000000000000)

Set DPS (verified): 36.4. Weights run: 3.6s. Verify run: 9.0s. 249 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.860 ± 0.020, crit=0.665 ± 0.020 per rating point (14 rating = 1%, 9.307 per %), hit=1.281 ± 0.052 per rating point (10 rating = 1%, 12.808 per %), melee_haste=6.212 ± 0.202, intellect=1.590 ± 0.024, spell_power=1.218 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (0.58 DPS) | yes | Trapper's Leather Hood (252505, -0.12 DPS) [crafted]; Totemic Leather Hood (252448, -0.18 DPS) [crafted]; Shadow Goggles (4373, -0.29 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 5.2 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 20.4 attack_power points (0.73 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.29 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.56 DPS) [crafted]; Serpent's Shoulders (5404, -0.58 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Sanguine Cape (14376, -0.03 DPS) [world_drop]; Lambent Scale Cloak (4706, -0.07 DPS, sim-verified) [world_drop] |
| chest | Totemic Leather Armor (252435) | Leatherworking [crafted] | 26.1 attack_power points (0.94 DPS) | yes | Stormrider's Leather Armor (252492, -0.30 DPS) [crafted]; Defender's Leather Armor (252434, -0.34 DPS) [crafted]; Brawler's Leather Armor (252490, -0.36 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bristlebark Bindings (14569, -0.05 DPS) [world_drop]; Bright Bracers (3647, -0.06 DPS) [world_drop]; Owl Bracers (4796, -0.41 DPS, sim-verified) [vendor] |
| hands | Totemic Leather Gloves (252497) | Leatherworking [crafted] | 14.1 attack_power points (0.51 DPS) | yes | Gold-flecked Gloves (5195, -0.00 DPS) [dungeon]; Serpent Gloves (5970, -0.01 DPS) [dungeon]; Gloves of the Fang (10413, -0.03 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Totemic Leather Belt (252431, -0.18 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.19 DPS) [world_drop]; Brawler's Leather Belt (252428, -0.24 DPS) [crafted] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 32.9 attack_power points (1.18 DPS) | yes | Trapper's Leather Pants (252501, -0.24 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.40 DPS) [crafted]; Defender's Leather Pants (252445, -0.41 DPS) [crafted] |
| feet | Totemic Leather Boots (252442) | Leatherworking [crafted] | 17.3 attack_power points (0.62 DPS) | yes | Stormrider's Leather Boots (252443, -0.07 DPS) [crafted]; Spidersilk Boots (4320, -0.09 DPS) [crafted]; Brawler's Leather Boots (252439, -0.11 DPS) [crafted] |
| finger1 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 14.0 attack_power points (0.50 DPS) | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; Demon Band (12054, -0.21 DPS) [world_drop]; Advisor's Ring (20426, -0.28 DPS) [rep] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.4 attack_power points (0.41 DPS) | yes | Lavishly Jeweled Ring (1156, -0.01 DPS) [dungeon]; Demon Band (12054, -0.12 DPS) [world_drop]; Advisor's Ring (20426, -0.19 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.0 attack_power points (10.91 DPS) | yes | Smite's Mighty Hammer (7230, -0.06 DPS) [dungeon]; Hammerbone (270018, -0.10 DPS) [quest]; Twisted Chanter's Staff (890, -0.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Totemic Leather Armor; wrist: Mindthrust Bracers; hands: Totemic Leather Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Totemic Leather Boots; finger1: Loop of Sacrifice; finger2: Legionnaire's Band; main_hand: Forsaken Greataxe

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads

### Band 30 (orc, 0000000000000000-055030310031000000-0000000000000000)

Set DPS (verified): 103.9. Weights run: 3.8s. Verify run: 7.5s. 442 eligible items had no known source.

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

### Band 40 (orc, 0000000000000000-055030310051202031-0000000000000000)

Set DPS (verified): 157.4. Weights run: 3.9s. Verify run: 12.3s. 679 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.392 ± 0.045, crit=1.528 ± 0.057 per rating point (14 rating = 1%, 21.388 per %), hit=2.291 ± 0.100 per rating point (10 rating = 1%, 22.907 per %), melee_haste=7.430 ± 0.822, intellect=3.548 ± 0.027, spell_power=0.375 ± 0.001

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Miner's Hat of the Deep (9429, +0.00 DPS) [dungeon]; Corpseshroud (10574, +0.00 DPS) [dungeon]; Thinking Cap (2624, -0.28 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prodigious Shadowshard Pendant (17773, +0.00 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.21 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.52 DPS) [world_drop] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 56.8 attack_power points (5.32 DPS) | yes | Windchaser Amice (14432, -0.61 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.75 DPS) [dungeon]; Green Silken Shoulders (7057, -1.38 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackforge Cape (6424, -1.00 DPS) [dungeon]; Glyphed Cloak (4732, -1.33 DPS) [world_drop]; Resilient Cape (14400, -1.66 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 65.8 attack_power points (6.17 DPS) | yes | Red Mageweave Vest (10007, -0.18 DPS) [crafted]; Silksand Tunic (14417, -0.84 DPS) [world_drop]; Silksand Wraps (14425, -0.84 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 34.7 attack_power points (3.26 DPS) | yes | Glowing Magical Bracelets (13106, -0.03 DPS) [world_drop]; Warden's Wristbands (14600, -0.20 DPS) [world_drop]; Tough Scorpid Bracers (8205, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Engineering Gloves (1659, +0.00 DPS) [world]; Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Rustler Gloves (9704, +0.00 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 55.8 attack_power points (5.24 DPS) | yes | Defiler's Lizardhide Girdle (20173, -0.25 DPS) [rep]; Defiler's Mail Girdle (20198, -0.25 DPS) [rep]; Saltscale Girdle (273023, -0.55 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 68.6 attack_power points (6.43 DPS) | yes | Crimson Silk Pantaloons (7062, -1.82 DPS) [crafted]; Red Mageweave Pants (10009, -1.95 DPS) [crafted]; Misplaced Pantaloons (276201, -2.11 DPS) [vendor] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 49.5 attack_power points (4.64 DPS) | yes | Skirmisher's Mail Boots (252564, -0.60 DPS) [crafted]; Stalker's Mail Boots (252562, -0.88 DPS) [crafted]; Briar Tredders (10582, -0.99 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 27.1 attack_power points (2.55 DPS) | yes | Ogremind Ring (1993, -0.22 DPS) [world_drop]; Mindbender Loop (5009, -0.22 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993), Mindbender Loop (5009)) | Bloodscalp Witch Doctor [world] | 24.8 attack_power points (2.33 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.07 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Thornstone Sledgehammer (1722, -3.30 DPS) [world_drop]; The Jackhammer (9423, -3.89 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; wrist: Windchaser Cuffs; hands: Gloves of Holy Might; waist: Deathmage Sash; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Fiery War Axe

No-known-source sample (15 of 679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 50 (orc, 3230000000000000-055030310051202051-0000000000000000)

Set DPS (verified): 242.3. Weights run: 3.9s. Verify run: 12.3s. 862 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.498 ± 0.080, crit=2.277 ± 0.113 per rating point (14 rating = 1%, 31.872 per %), hit=4.290 ± 0.235 per rating point (10 rating = 1%, 42.905 per %), melee_haste=13.063 ± 1.631, intellect=4.221 ± 0.082, spell_power=1.087 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Chief Architect's Monocle (11839, +0.00 DPS) [dungeon]; Soulcatcher Halo (10630, -0.63 DPS) [dungeon]; Braincage (12549, -0.63 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 59.1 attack_power points (5.65 DPS) | yes | Gemshard Heart (17707, -1.61 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.69 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.60 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 100.7 attack_power points (9.63 DPS) | yes | Rotgrip Mantle (17732, -1.01 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -2.42 DPS) [dungeon]; Dregmetal Spaulders (11722, -2.43 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Red Cloak (8248, -1.21 DPS) [world_drop]; Runecloth Cloak (13860, -1.49 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.27 DPS, sim-verified) [quest] |
| chest | Stone Guard's Pulsing Breastplate (220844) | Lady Palanseer [vendor] | 113.1 attack_power points (10.81 DPS) | yes | Stone Guard's Mail Armor (220826, -0.63 DPS) [vendor]; Robes of Insight (940, -0.73 DPS) [world_drop]; Acumen Robes (17775, -0.77 DPS) [quest] |
| wrist | Forgotten Wraps (9433) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Runic Leather Bracers (15092, -0.20 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.43 DPS) [quest] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Azurite Fists (20369, -0.85 DPS) [quest]; Rockgrip Gauntlets (17736, -0.87 DPS) [dungeon]; First Sergeant's Mail Gauntlets (220831, -0.99 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Defiler's Lizardhide Girdle (20174, -1.21 DPS) [rep]; Defiler's Mail Girdle (20196, -1.21 DPS) [rep] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 113.1 attack_power points (10.81 DPS) | yes | Stone Guard's Mail Legplates (220834, -0.63 DPS) [vendor]; Imperial Red Pants (8251, -2.61 DPS) [world_drop]; Stone Guard's Inscribed Legplates (220839, -2.93 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 97.1 attack_power points (9.28 DPS) | yes | Skycaller's Leather Boots (252471, -3.28 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -3.28 DPS) [crafted]; Rancor Boots (11865, -3.34 DPS, sim-verified) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 63.3 attack_power points (6.05 DPS) | yes | Mindseye Circle (10634, -1.21 DPS) [dungeon]; Cyclopean Band (11824, -1.53 DPS) [dungeon]; Woodseed Hoop (17768, -2.42 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 62.9 attack_power points (6.01 DPS) | yes | Mindseye Circle (10634, -1.17 DPS) [dungeon]; Cyclopean Band (11824, -1.49 DPS) [dungeon]; Woodseed Hoop (17768, -2.38 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -3.59 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -2.52 DPS, sim-verified) [crafted]; Uther's Strength (11302, -3.07 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildstaff (20556, +0.00 DPS) [quest]; Glowing Brightwood Staff (812, -1.95 DPS) [world_drop]; Soulkeeper (1607, -3.01 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Forgotten Wraps; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Stone Guard's Pulsing Legplates; feet: Greaves of Withering Despair; finger1: Brainlash; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 862, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle

### Band 60 (orc, 3230030000000000-055030310051202051-0520000000000000)

Set DPS (verified): 378.7. Weights run: 4.0s. Verify run: 48.2s. 2043 eligible items had no known source.

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

Set DPS (verified): 906.7. Weights run: 3.5s. Verify run: 44.0s. 2043 eligible items had no known source.

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

