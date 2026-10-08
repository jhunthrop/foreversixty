# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.7. Weights run: 8.1s. Verify run: 3.9s. 150 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.010, intellect=2.580 ± 0.011, spirit=1.351 ± 0.004, mp5=2.935 ± 0.023, crit=0.149 ± 0.007 per rating point (14 rating = 1%, 2.093 per %), spell_haste=0.153 ± 0.035

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 21.0 healing_power points (1.02 DPS) | yes | Pristine Circlet (253949, -0.06 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.76 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 5.4 healing_power points (0.26 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 23.2 healing_power points (1.13 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.10 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.70 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (20428, -0.02 DPS, sim-verified) [rep]; Seer's Cape (6378, -0.12 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 30.0 healing_power points (1.46 DPS) | yes | Robe of the Moccasin (6465, -0.10 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.40 DPS) [dungeon]; Seer's Robe (2981, -0.51 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.7 healing_power points (1.06 DPS) | yes | Bright Bracers (3647, -0.06 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.67 DPS) [quest]; Seer's Cuffs (3645, -0.80 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 18.7 healing_power points (0.91 DPS) | yes | Magefist Gloves (12977, -0.02 DPS) [world_drop]; Bright Gloves (3066, -0.15 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.16 DPS) [quest] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Sash (253925, -0.02 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.14 DPS) [crafted]; Tarantula Silk Sash (3229, -0.25 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 38.9 healing_power points (1.89 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -0.88 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.88 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.7 healing_power points (0.81 DPS) | yes | Kimbra Boots (6191, +0.00 DPS, sim-verified) [quest]; Bluegill Sandals (1560, -0.30 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 15.5 healing_power points (0.75 DPS) | yes | Band of Purification (12996, -0.36 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Deep Fathom Ring (6463, -0.42 DPS) [dungeon] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 13.3 healing_power points (0.65 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop]; Deep Fathom Ring (6463, -0.32 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 25.8 healing_power points (1.26 DPS) | yes | Staff of Westfall (2042, -0.08 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.25 DPS) [world]; Lesser Staff of the Spire (1300, -0.50 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 5.2 healing_power points (0.25 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.05 DPS) [quest]; Dwarven Flamestick (5241, -0.12 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Keller's Girdle; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (human, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 73.4. Weights run: 8.1s. Verify run: 3.9s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.014, intellect=2.158 ± 0.012, spirit=1.926 ± 0.011, mp5=3.911 ± 0.011, crit=0.228 ± 0.009 per rating point (14 rating = 1%, 3.186 per %), spell_haste=not significant (-0.128 ± 0.086)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Resilient Cap (14401, -0.27 DPS) [world_drop]; Holy Shroud (2721, -0.51 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.78 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.3 healing_power points (1.09 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS, sim-verified) [world_drop]; Necklace of Harmony (5180, -0.19 DPS) [world]; Scorn's Icy Choker (23169, -0.23 DPS) [dungeon] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 33.4 healing_power points (2.22 DPS) | yes | Mantle of Honor (3560, -0.15 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.43 DPS) [world_drop]; Faerie Mantle (5820, -0.59 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 23.0 healing_power points (1.53 DPS) | yes | Glowing Thresher Cape (6901, -0.04 DPS) [dungeon]; Repairman's Cape (9605, -0.06 DPS) [quest]; Prelacy Cape (7004, -0.26 DPS, sim-verified) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.06 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.10 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.69 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.5 healing_power points (1.50 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.35 DPS) [world_drop]; Spidertank Oilrag (9448, -0.72 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 33.9 healing_power points (2.25 DPS) | yes | Hotshot Pilot's Gloves (9491, +0.00 DPS, sim-verified) [dungeon]; Town Clerk's Mittens (270029, -0.68 DPS) [quest]; Zodiac Gloves (7106, -0.78 DPS) [quest] |
| waist | Resilient Cord (14406) | World drop [world_drop] | 20.7 healing_power points (1.37 DPS) | yes | Pristine Sash (253925, +0.00 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.11 DPS) [vendor]; Novice Ardent's Sash (253887, -0.16 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.7 healing_power points (3.11 DPS) | yes | Filigreed Pristine Leggings (253937, -0.07 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.70 DPS) [crafted]; Blighted Leggings (7709, -1.31 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.8 healing_power points (2.85 DPS) | yes | Acidic Walkers (9454, -0.09 DPS, sim-verified) [dungeon]; Soggy Boots (274747, -1.25 DPS) [vendor]; Frothing Slippers (254003, -1.33 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 28.6 healing_power points (1.90 DPS) | yes | Sea Giant's Toe Ring (274746, -0.70 DPS) [vendor]; Black Pearl Ring (6332, -0.85 DPS) [world]; Darkspear Signet (272071, -0.86 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 19.7 healing_power points (1.31 DPS) | yes | Sea Giant's Toe Ring (274746, -0.07 DPS, sim-verified) [vendor]; Black Pearl Ring (6332, -0.26 DPS) [world]; Darkspear Signet (272071, -0.27 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.22 DPS, sim-verified) [dungeon]; Staff of the Friar (3415, -0.67 DPS) [dungeon]; Gnarled Ash Staff (791, -0.72 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 8.6 healing_power points (0.57 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Consecrated Wand (5244, -0.03 DPS) [quest]; Gravestone Scepter (7001, -0.19 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Pristine Gown; hands: Gloves of Old; waist: Resilient Cord; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 161.2. Weights run: 9.1s. Verify run: 4.5s. 328 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.092, intellect=1.518 ± 0.033, spirit=0.538 ± 0.027, mp5=0.420 ± 0.042, crit=0.473 ± 0.019 per rating point (14 rating = 1%, 6.621 per %), spell_haste=not significant (0.436 ± 0.299)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 52.6 healing_power points (5.37 DPS) | yes | Corpseshroud (10574, -2.10 DPS) [dungeon]; Miner's Hat of the Deep (9429, -2.19 DPS) [dungeon]; Holy Shroud (2721, -2.41 DPS, sim-verified) [world_drop] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 15.2 healing_power points (1.55 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.08 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.17 DPS) [dungeon] |
| shoulder | Windchaser Amice (14432) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.04 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.50 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 22.7 healing_power points (2.31 DPS) | yes | Caretaker's Cape (19532, -0.37 DPS, sim-verified) [rep]; Darkspear Raider's Cloak (272077, -0.39 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.3 healing_power points (6.36 DPS) | yes | Death Speaker Robes (6682, -0.95 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -3.01 DPS) [crafted]; Red Mageweave Vest (10007, -3.57 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 13.7 healing_power points (1.39 DPS) | yes | Mistscape Bracers (4045, -0.15 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.15 DPS) [quest]; Earthen Silk Cuffs (254019, -0.34 DPS, sim-verified) [crafted] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Handwraps (254021, -0.28 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.70 DPS) [crafted]; Truefaith Gloves (7049, -0.99 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 37.4 healing_power points (3.81 DPS) | yes | Deathmage Sash (10771, -0.15 DPS, sim-verified) [dungeon]; Sutarn's Ring (13105, -1.94 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.95 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 35.3 healing_power points (3.60 DPS) | yes | Filigreed Pristine Leggings (253937, -0.22 DPS, sim-verified) [crafted]; Stormcloth Pants (10010, -1.25 DPS) [crafted]; Aurora Pants (4044, -1.35 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.8 healing_power points (3.34 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Nimbus Boots (6998, -1.61 DPS) [quest]; Thoughtcast Boots (10578, -1.70 DPS) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.84 DPS) | yes | Ogremind Ring (1993, -0.59 DPS) [world_drop]; Voodoo Band (1996, -0.59 DPS) [world]; Mindbender Loop (5009, -0.64 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 14.4 healing_power points (1.47 DPS) | yes | Voodoo Band (1996, -0.15 DPS, sim-verified) [world]; Ogremind Ring (1993, -0.22 DPS) [world_drop]; Mindbender Loop (5009, -0.27 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -4.62 DPS) [dungeon]; Staff of Jordan (873, -4.63 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 24.2 healing_power points (2.46 DPS) | yes | Orb of Souls (249395, -0.17 DPS, sim-verified) [crafted]; Orb of Lorica (11262, -0.92 DPS) [quest]; Eye of Paleth (2943, -1.14 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 10.7 healing_power points (1.09 DPS) | yes | Goblin Igniter (5253, -0.16 DPS, sim-verified) [quest]; Flash Wand (5248, -0.31 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.32 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Prodigious Shadowshard Pendant; shoulder: Windchaser Amice; back: Mantle of Lady Falther'ess; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 235.9. Weights run: 8.9s. Verify run: 4.7s. 423 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.143, intellect=1.579 ± 0.053, spirit=not significant (-0.139 ± 0.092), mp5=2.676 ± 0.137, crit=0.726 ± 0.026 per rating point (14 rating = 1%, 10.161 per %), spell_haste=5.244 ± 0.510

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Cassandra's Grace (13102, -0.45 DPS) [world_drop]; Chief Architect's Monocle (11839, -0.57 DPS) [dungeon]; Knight-Lieutenant's Satin Cover (220896, -1.33 DPS, sim-verified) [vendor] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 25.5 healing_power points (2.35 DPS) | yes | Horizon Choker (13085, +0.00 DPS, sim-verified) [world_drop]; Gemshard Heart (17707, -0.90 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.90 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Satin Pads (220894, -0.35 DPS, sim-verified) [vendor]; Kentic Amice (11624, -0.38 DPS) [dungeon]; Rotgrip Mantle (17732, -0.94 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (19531, -0.01 DPS) [rep]; Mantle of Lady Falther'ess (23178, -0.47 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 58.0 healing_power points (5.34 DPS) | yes | Robes of Insight (940, -1.71 DPS) [world_drop]; Knight's Satin Armor (220892, -1.82 DPS, sim-verified) [vendor]; Death Speaker Robes (6682, -1.90 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 31.1 healing_power points (2.86 DPS) | yes | Aristocratic Cuffs (12546, +0.00 DPS, sim-verified) [dungeon]; Forgotten Wraps (9433, -1.12 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -1.12 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 56.7 healing_power points (5.23 DPS) | yes | Gilded Gloves (254095, -1.10 DPS, sim-verified) [crafted]; Virtuous Mitts (226950, -1.42 DPS) [vendor]; Sergeant Major's Satin Gloves (220897, -1.79 DPS) [vendor] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 43.2 healing_power points (3.98 DPS) | yes | Gilded Cord (254037, -0.41 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.22 DPS) [dungeon]; Ban'thok Sash (11662, -1.28 DPS) [dungeon] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | 56.0 healing_power points (5.16 DPS) | yes | Spellshock Leggings (9484, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -2.11 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.48 DPS) [vendor] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 45.2 healing_power points (4.17 DPS) | yes | Coldstone Slippers (18697, -1.14 DPS) [dungeon]; Sergeant Major's Satin Boots (220895, -1.20 DPS, sim-verified) [vendor]; Gilded Slippers (254001, -1.30 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 23.7 healing_power points (2.18 DPS) | yes | Mindseye Circle (10634, -0.44 DPS) [dungeon]; Darkspear Signet (272069, -0.46 DPS) [vendor]; Sea Giant's Toe Ring (274746, -0.52 DPS) [vendor] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 20.1 healing_power points (1.85 DPS) | yes | Darkspear Signet (272069, -0.12 DPS) [vendor]; Sea Giant's Toe Ring (274746, -0.19 DPS) [vendor]; Mindseye Circle (10634, -0.54 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, +0.00 DPS) [quest] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -2.05 DPS) [world_drop]; Spellshifter Rod (9527, -2.92 DPS) [quest] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.1 healing_power points (2.04 DPS) | yes | Twisting Essence Jar (249456, -0.29 DPS) [crafted]; Skullspell Orb (10708, -0.29 DPS) [quest]; Beacon of Hope (9393, -0.97 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 9.5 healing_power points (0.87 DPS) | yes | Cairnstone Sliver (9654, +0.00 DPS, sim-verified) [quest]; Lesser Mystic Wand (11289, -0.29 DPS) [crafted]; Starfaller (13063, -0.29 DPS) [world_drop] |

**New at 50:** neck: Darkmoon Necklace; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Knight's Satin Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; off_hand: Enthralled Sphere

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 384.7. Weights run: 8.3s. Verify run: 3.4s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.197, intellect=1.751 ± 0.049, spirit=0.754 ± 0.056, mp5=1.962 ± 0.079, crit=0.628 ± 0.034 per rating point (14 rating = 1%, 8.791 per %), spell_haste=not significant (1.101 ± 0.661)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Virtuous Crown (226947, -1.77 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 50.2 healing_power points (7.17 DPS) | yes | Drake Tooth Necklace (21531, -0.46 DPS) [quest]; The Eye of Zuldazar (19593, -1.20 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.20 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+9.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -9.70 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 59.5 healing_power points (8.50 DPS) | yes | Drape of Recovery (272413, -2.95 DPS) [vendor]; Darkspear Raider's Cloak (272063, -3.85 DPS) [vendor]; Cloak of the Cosmos (18389, -4.31 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-verified (+12.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.00 DPS) [dungeon]; Virtuous Robe (226945, -12.64 DPS, sim-verified) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 46.3 healing_power points (6.61 DPS) | yes | Virtuous Bracers (226949, -0.61 DPS) [quest]; Marshal's Satin Bracers (17606, -0.79 DPS) [pvp]; Bracers of Mending (23129, -1.20 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 71.3 healing_power points (10.18 DPS) | yes | Desert Bloom Gloves (20717, -1.11 DPS) [quest]; Mooncloth Gloves (18409, -1.47 DPS) [crafted]; Hands of the Exalted Herald (12554, -2.04 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 62.3 healing_power points (8.90 DPS) | yes | Whipvine Cord (18327, -0.82 DPS, sim-verified) [dungeon]; Virtuous Belt (226948, -0.94 DPS) [quest]; Devout Belt (16696, -1.36 DPS) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 88.8 healing_power points (12.68 DPS) | yes | Marshal's Satin Legguards (231626, -0.89 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -1.71 DPS) [pvp]; Virtuous Skirt (226946, -13.82 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Walkers (231627, -0.18 DPS) [vendor]; Mooncloth Boots (15802, -0.49 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -4.42 DPS, sim-verified) [vendor] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.29 DPS) [quest]; Emerald Flame Ring (18395, -0.82 DPS) [dungeon]; Naglering (11669, -13.92 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.23 DPS) [quest]; Emerald Flame Ring (18395, -0.76 DPS) [dungeon]; Naglering (11669, -14.17 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (+10.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -1.69 DPS) [dungeon]; Second Wind (11819, -2.69 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Briarwood Reed (12930, -0.06 DPS) [dungeon]; Second Wind (11819, -1.06 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -0.57 DPS) [dungeon]; Death Speaker Scepter (2816, -1.01 DPS) [dungeon]; Hand of Edward the Odd (2243, -12.37 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 25.3 healing_power points (3.61 DPS) | yes | Oblivion's Touch (18761, -0.86 DPS) [dungeon]; Bonecreeper Stylus (13938, -1.04 DPS) [dungeon]; Sparkling Crystal Wand (20672, -1.18 DPS, sim-verified) [world] |

**New at 60:** head: Devout Crown; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Band of Piety; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 510.2. Weights run: 4.8s. Verify run: 1.8s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.329, intellect=1.254 ± 0.040, spirit=1.190 ± 0.037, mp5=3.009 ± 0.042, crit=0.641 ± 0.038 per rating point (14 rating = 1%, 8.971 per %), spell_haste=not significant (-1.758 ± 0.692)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 82.3 healing_power points (17.73 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.67 DPS) [pvp]; Mooncloth Circlet (14140, -2.90 DPS) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 55.6 healing_power points (11.98 DPS) | yes | The Eye of Zuldazar (19593, -0.98 DPS, sim-verified) [quest]; The All-Seeing Eye of Zuldazar (19594, -2.60 DPS) [quest]; Drake Tooth Necklace (21531, -3.13 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 89.0 healing_power points (19.16 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -5.65 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -7.30 DPS) [pvp]; Virtuous Mantle (226951, -7.84 DPS) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 54.5 healing_power points (11.75 DPS) | yes | Drape of Recovery (272413, -0.71 DPS, sim-verified) [vendor]; Cloak of the Cosmos (18389, -3.18 DPS) [dungeon]; Caretaker's Cape (19530, -4.10 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.0 healing_power points (19.83 DPS) | yes | Robes of the Exalted (13346, -0.75 DPS, sim-verified) [dungeon]; Field Marshal's Satin Tunic (231624, -1.13 DPS) [vendor]; Virtuous Robe (226945, -1.76 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 46.1 healing_power points (9.93 DPS) | yes | Bracers of Mending (23129, -0.28 DPS) [dungeon]; Virtuous Bracers (226949, -0.80 DPS) [quest]; Marshal's Satin Bracers (17606, -2.23 DPS) [pvp] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 63.6 healing_power points (13.70 DPS) | yes | Desert Bloom Gloves (20717, -0.76 DPS, sim-verified) [quest]; Virtuous Mitts (226950, -0.94 DPS) [vendor]; Raider Handwraps (272097, -1.23 DPS) [vendor] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 60.3 healing_power points (13.00 DPS) | yes | Wisdom of the Timbermaw (19047, -0.23 DPS) [crafted]; Virtuous Belt (226948, -0.80 DPS) [quest]; Penitent's Cinch (272394, -1.89 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 85.1 healing_power points (18.34 DPS) | yes | Virtuous Skirt (226946, -1.42 DPS, sim-verified) [quest]; Marshal's Satin Legguards (231626, -1.96 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -2.08 DPS) [pvp] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 78.2 healing_power points (16.85 DPS) | yes | Virtuous Sandals (226952, -0.85 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -4.56 DPS) [crafted]; Faith Healer's Boots (22247, -4.93 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (510.5 DPS) | yes | Rosewine Circle (13178, -0.62 DPS) [dungeon]; Fordring's Seal (16058, -0.79 DPS) [quest]; Naglering (11669, -1.05 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (510.5 DPS) | yes | Rosewine Circle (13178, -0.54 DPS) [dungeon]; Fordring's Seal (16058, -0.71 DPS) [quest]; Naglering (11669, -0.90 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | Alchemists' Stone (13503) | Alchemy [crafted] | sim-verified (510.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (510.5 DPS) | yes | Hand of Edward the Odd (2243, -1.42 DPS, sim-verified) [world_drop]; Staff of Metanoia (22394, -2.37 DPS) [dungeon]; Death Speaker Scepter (2816, -2.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.1 healing_power points (6.48 DPS) | yes | Sparkling Crystal Wand (20672, -2.97 DPS) [world]; Bonecreeper Stylus (13938, -3.03 DPS) [dungeon]; Oblivion's Touch (18761, -3.51 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Hands of the Exalted Herald; waist: Whipvine Cord; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Band of Mending; finger2: Band of Piety; trinket2: Alchemists' Stone; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.4. Weights run: 8.1s. Verify run: 3.8s. 140 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.010, intellect=2.580 ± 0.011, spirit=1.351 ± 0.004, mp5=2.935 ± 0.023, crit=0.149 ± 0.007 per rating point (14 rating = 1%, 2.093 per %), spell_haste=0.153 ± 0.035

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 21.0 healing_power points (1.02 DPS) | yes | Pristine Circlet (253949, -0.06 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.76 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 5.4 healing_power points (0.26 DPS) | yes | Roadwatcher's Confidence (281265, -0.07 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 23.2 healing_power points (1.13 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.06 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.70 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (20427, -0.05 DPS, sim-verified) [rep]; Seer's Cape (6378, -0.12 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.13 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 30.0 healing_power points (1.46 DPS) | yes | Robe of the Moccasin (6465, -0.06 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.40 DPS) [dungeon]; Seer's Robe (2981, -0.51 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.7 healing_power points (1.06 DPS) | yes | Tabitha's Cuffs (251486, +0.00 DPS, sim-verified) [quest]; Featherbead Bracers (15452, -0.43 DPS) [quest]; Bright Bracers (3647, -0.55 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 24.8 healing_power points (1.21 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.32 DPS) [world_drop]; Bright Gloves (3066, -0.44 DPS) [world_drop] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Sash (253925, -0.04 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.14 DPS) [crafted]; Tarantula Silk Sash (3229, -0.25 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 38.9 healing_power points (1.89 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -0.88 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.88 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.7 healing_power points (0.81 DPS) | yes | Bluegill Sandals (1560, +0.00 DPS, sim-verified) [world]; Walking Boots (4660, -0.31 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 15.5 healing_power points (0.75 DPS) | yes | Black Pearl Ring (6332, -0.11 DPS, sim-verified) [world]; Band of Purification (12996, -0.36 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Pearl Ring (6332, -0.04 DPS, sim-verified) [world]; Band of Purification (12996, -0.23 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 25.8 healing_power points (1.26 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.14 DPS) [quest]; Channeler's Staff (4437, -0.25 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 5.2 healing_power points (0.25 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.13 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Keller's Girdle; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.8. Weights run: 8.1s. Verify run: 3.9s. 231 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.014, intellect=2.158 ± 0.012, spirit=1.926 ± 0.011, mp5=3.911 ± 0.011, crit=0.228 ± 0.009 per rating point (14 rating = 1%, 3.186 per %), spell_haste=not significant (-0.128 ± 0.086)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Resilient Cap (14401, -0.27 DPS) [world_drop]; Holy Shroud (2721, -0.52 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.78 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.3 healing_power points (1.09 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Necklace of Harmony (5180, -0.19 DPS) [world]; Scorn's Icy Choker (23169, -0.23 DPS) [dungeon] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 33.4 healing_power points (2.22 DPS) | yes | Nightsky Mantle (4718, -0.11 DPS, sim-verified) [world_drop]; Ghostly Mantle (3324, -0.47 DPS) [quest]; Death Speaker Mantle (6685, -0.64 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 23.0 healing_power points (1.53 DPS) | yes | Battle Healer's Cloak (19529, -0.16 DPS) [rep]; Glowing Thresher Cape (6901, -0.24 DPS, sim-verified) [dungeon]; Cloak of Rot (4462, -0.38 DPS) [world] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.06 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.10 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.69 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.5 healing_power points (1.50 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.35 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.64 DPS) [quest] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan: Caverndeep Burrower [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of Old (9395, -0.04 DPS, sim-verified) [world_drop]; Blight Gloves (279877, -0.14 DPS) [quest]; Tattered Mittens (270030, -0.35 DPS) [quest] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 25.2 healing_power points (1.67 DPS) | yes | Resilient Cord (14406, -0.08 DPS, sim-verified) [world_drop]; Pristine Sash (253925, -0.37 DPS) [crafted]; Dreamer's Belt (4829, -0.41 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.7 healing_power points (3.11 DPS) | yes | Filigreed Pristine Leggings (253937, -0.07 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.70 DPS) [crafted]; Sacred Burial Trousers (6282, -1.24 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.8 healing_power points (2.85 DPS) | yes | Acidic Walkers (9454, -0.10 DPS, sim-verified) [dungeon]; Soggy Boots (274747, -1.25 DPS) [vendor]; Frothing Slippers (254003, -1.33 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 28.6 healing_power points (1.90 DPS) | yes | Sea Giant's Toe Ring (274746, -0.70 DPS) [vendor]; Black Pearl Ring (6332, -0.85 DPS) [world]; Darkspear Signet (272071, -0.86 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 19.7 healing_power points (1.31 DPS) | yes | Sea Giant's Toe Ring (274746, -0.05 DPS, sim-verified) [vendor]; Black Pearl Ring (6332, -0.26 DPS) [world]; Darkspear Signet (272071, -0.27 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.23 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (19569, -0.59 DPS) [pvp]; Staff of the Friar (3415, -0.67 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 8.6 healing_power points (0.57 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Gravestone Scepter (7001, -0.19 DPS) [quest]; Wand of Decay (5252, -0.29 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Pristine Gown; hands: Hotshot Pilot's Gloves; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 158.7. Weights run: 9.1s. Verify run: 4.5s. 311 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.092, intellect=1.518 ± 0.033, spirit=0.538 ± 0.027, mp5=0.420 ± 0.042, crit=0.473 ± 0.019 per rating point (14 rating = 1%, 6.621 per %), spell_haste=not significant (0.436 ± 0.299)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 52.6 healing_power points (5.37 DPS) | yes | Corpseshroud (10574, -2.10 DPS) [dungeon]; Miner's Hat of the Deep (9429, -2.19 DPS) [dungeon]; Holy Shroud (2721, -2.39 DPS, sim-verified) [world_drop] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 15.2 healing_power points (1.55 DPS) | yes | Necklace of Calisea (1714, -0.08 DPS) [dungeon]; Triune Amulet (7722, -0.08 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.17 DPS) [dungeon] |
| shoulder | Windchaser Amice (14432) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.04 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.53 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 22.7 healing_power points (2.31 DPS) | yes | Darkspear Raider's Cloak (272077, -0.39 DPS) [vendor]; Battle Healer's Cloak (19528, -0.41 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.3 healing_power points (6.36 DPS) | yes | Death Speaker Robes (6682, -0.95 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -3.01 DPS) [crafted]; Red Mageweave Vest (10007, -3.57 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 13.7 healing_power points (1.39 DPS) | yes | Mistscape Bracers (4045, -0.15 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.15 DPS) [quest]; Earthen Silk Cuffs (254019, -0.37 DPS, sim-verified) [crafted] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Handwraps (254021, -0.16 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.70 DPS) [crafted]; Truefaith Gloves (7049, -0.99 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 37.4 healing_power points (3.81 DPS) | yes | Deathmage Sash (10771, -0.26 DPS, sim-verified) [dungeon]; Sutarn's Ring (13105, -1.94 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.95 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 35.3 healing_power points (3.60 DPS) | yes | Filigreed Pristine Leggings (253937, -0.26 DPS, sim-verified) [crafted]; Stormcloth Pants (10010, -1.25 DPS) [crafted]; Aurora Pants (4044, -1.35 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.8 healing_power points (3.34 DPS) | yes | Furen's Boots (13100, -1.38 DPS) [world_drop]; Boots of the Maharishi (9658, -1.62 DPS) [quest]; Thoughtcast Boots (10578, -1.70 DPS) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.84 DPS) | yes | Ogremind Ring (1993, -0.59 DPS) [world_drop]; Voodoo Band (1996, -0.59 DPS) [world]; Mindbender Loop (5009, -0.64 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 14.4 healing_power points (1.47 DPS) | yes | Voodoo Band (1996, -0.16 DPS, sim-verified) [world]; Ogremind Ring (1993, -0.22 DPS) [world_drop]; Mindbender Loop (5009, -0.27 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -4.62 DPS) [dungeon]; Staff of Jordan (873, -4.63 DPS) [world_drop] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Souls (249395, -0.20 DPS) [crafted]; Beacon of Hope (9393, -0.47 DPS, sim-verified) [dungeon]; Aurora Sphere (7610, -0.55 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 10.7 healing_power points (1.09 DPS) | yes | Goblin Igniter (5253, -0.21 DPS, sim-verified) [quest]; Flash Wand (5248, -0.31 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.32 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Prodigious Shadowshard Pendant; shoulder: Windchaser Amice; back: Mantle of Lady Falther'ess; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Prophetic Cane; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 240.5. Weights run: 8.9s. Verify run: 4.7s. 403 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.143, intellect=1.579 ± 0.053, spirit=not significant (-0.139 ± 0.092), mp5=2.676 ± 0.137, crit=0.726 ± 0.026 per rating point (14 rating = 1%, 10.161 per %), spell_haste=5.244 ± 0.510

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 48.8 healing_power points (4.50 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Chief Architect's Monocle (11839, -0.57 DPS) [dungeon]; Cassandra's Grace (13102, -3.73 DPS, sim-verified) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Gemshard Heart (17707, -0.58 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.58 DPS) [quest]; Darkmoon Necklace (19303, -1.35 DPS, sim-verified) [vendor] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 38.6 healing_power points (3.56 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Kentic Amice (11624, -0.38 DPS) [dungeon]; Rotgrip Mantle (17732, -0.94 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (19527, -0.01 DPS) [rep]; Mantle of Lady Falther'ess (23178, -0.34 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 58.0 healing_power points (5.34 DPS) | yes | Stone Guard's Satin Armor (220903, -1.47 DPS, sim-verified) [vendor]; Robes of Insight (940, -1.71 DPS) [world_drop]; Death Speaker Robes (6682, -1.90 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 31.1 healing_power points (2.86 DPS) | yes | Aristocratic Cuffs (12546, +0.00 DPS, sim-verified) [dungeon]; Forgotten Wraps (9433, -1.12 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -1.12 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 56.7 healing_power points (5.23 DPS) | yes | Gilded Gloves (254095, -1.07 DPS, sim-verified) [crafted]; Virtuous Mitts (226950, -1.42 DPS) [vendor]; Greenleaf Handwraps (19116, -1.60 DPS) [quest] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 43.2 healing_power points (3.98 DPS) | yes | Gilded Cord (254037, -0.33 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.22 DPS) [dungeon]; Ban'thok Sash (11662, -1.28 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.53 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -0.89 DPS) [vendor]; Stone Guard's Satin Leggings (220902, -0.98 DPS, sim-verified) [vendor] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 45.2 healing_power points (4.17 DPS) | yes | Coldstone Slippers (18697, -1.14 DPS) [dungeon]; First Sergeant's Satin Boots (220900, -1.27 DPS, sim-verified) [vendor]; Gilded Slippers (254001, -1.30 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 23.7 healing_power points (2.18 DPS) | yes | Mindseye Circle (10634, -0.44 DPS) [dungeon]; Darkspear Signet (272069, -0.46 DPS) [vendor]; Sea Giant's Toe Ring (274746, -0.52 DPS) [vendor] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 20.1 healing_power points (1.85 DPS) | yes | Mindseye Circle (10634, +0.00 DPS, sim-verified) [dungeon]; Darkspear Signet (272069, -0.12 DPS) [vendor]; Sea Giant's Toe Ring (274746, -0.19 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Prayer Beads (19990, +0.00 DPS) [quest] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -2.05 DPS) [world_drop]; Spellshifter Rod (9527, -2.92 DPS) [quest] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.1 healing_power points (2.04 DPS) | yes | Twisting Essence Jar (249456, -0.29 DPS) [crafted]; Prophetic Cane (6803, -0.29 DPS) [quest]; Beacon of Hope (9393, -0.74 DPS, sim-verified) [dungeon] |
| ranged | Nature's Breath (19118) (or Jaina's Firestarter (13064)) | Dark Vessels [quest] | 9.5 healing_power points (0.87 DPS) | yes | Jaina's Firestarter (13064, +0.00 DPS, sim-verified) [world_drop]; Lesser Mystic Wand (11289, -0.29 DPS) [crafted]; Starfaller (13063, -0.29 DPS) [world_drop] |

**New at 50:** neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Spellshock Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; off_hand: Enthralled Sphere; ranged: Nature's Breath

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 379.6. Weights run: 8.3s. Verify run: 3.3s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.197, intellect=1.751 ± 0.049, spirit=0.754 ± 0.056, mp5=1.962 ± 0.079, crit=0.628 ± 0.034 per rating point (14 rating = 1%, 8.791 per %), spell_haste=not significant (1.101 ± 0.661)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 83.6 healing_power points (11.94 DPS) | yes | Devout Crown (16693, +0.00 DPS, sim-verified) [dungeon]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.22 DPS) [pvp] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 50.2 healing_power points (7.17 DPS) | yes | The Eye of Zuldazar (19593, -1.20 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.20 DPS) [quest]; Drake Tooth Necklace (21531, -4.29 DPS, sim-verified) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+7.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -7.24 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 59.5 healing_power points (8.50 DPS) | yes | Drape of Recovery (272413, -2.95 DPS) [vendor]; Darkspear Raider's Cloak (272063, -3.85 DPS) [vendor]; Cloak of the Cosmos (18389, -5.34 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-verified (+10.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.00 DPS) [dungeon]; Virtuous Robe (226945, -10.01 DPS, sim-verified) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 46.3 healing_power points (6.61 DPS) | yes | Virtuous Bracers (226949, -0.61 DPS) [quest]; General's Satin Bracers (17619, -0.79 DPS) [pvp]; Bracers of Mending (23129, -1.78 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 71.3 healing_power points (10.18 DPS) | yes | Desert Bloom Gloves (20717, -1.11 DPS) [quest]; Mooncloth Gloves (18409, -1.47 DPS) [crafted]; Hands of the Exalted Herald (12554, -4.01 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 62.3 healing_power points (8.90 DPS) | yes | Virtuous Belt (226948, -0.94 DPS) [quest]; Devout Belt (16696, -1.36 DPS) [dungeon]; Whipvine Cord (18327, -2.67 DPS, sim-verified) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 88.8 healing_power points (12.68 DPS) | yes | General's Satin Legguards (231634, -0.89 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -1.71 DPS) [pvp]; Virtuous Skirt (226946, -14.60 DPS, sim-verified) [quest] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 79.1 healing_power points (11.29 DPS) | yes | Virtuous Sandals (226952, +0.00 DPS, sim-verified) [quest]; General's Satin Walkers (231630, -2.65 DPS) [vendor]; Mooncloth Boots (15802, -2.96 DPS) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.29 DPS) [quest]; Emerald Flame Ring (18395, -0.82 DPS) [dungeon]; Naglering (11669, -14.28 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.23 DPS) [quest]; Emerald Flame Ring (18395, -0.76 DPS) [dungeon]; Naglering (11669, -16.63 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.51 DPS, sim-verified) [vendor]; Briarwood Reed (12930, -1.69 DPS) [dungeon]; Second Wind (11819, -2.69 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, -0.06 DPS) [dungeon]; Second Wind (11819, -1.06 DPS) [dungeon]; Serenity Field (272439, -1.15 DPS, sim-verified) [vendor] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -0.57 DPS) [dungeon]; Death Speaker Scepter (2816, -1.01 DPS) [dungeon]; Hand of Edward the Odd (2243, -13.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 25.3 healing_power points (3.61 DPS) | yes | Oblivion's Touch (18761, -0.86 DPS) [dungeon]; Bonecreeper Stylus (13938, -1.04 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.44 DPS, sim-verified) [world] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Band of Mending; finger2: Band of Piety; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 509.9. Weights run: 4.8s. Verify run: 1.9s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.329, intellect=1.254 ± 0.040, spirit=1.190 ± 0.037, mp5=3.009 ± 0.042, crit=0.641 ± 0.038 per rating point (14 rating = 1%, 8.971 per %), spell_haste=not significant (-1.758 ± 0.692)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 82.3 healing_power points (17.73 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.67 DPS) [pvp]; Mooncloth Circlet (14140, -2.90 DPS) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 55.6 healing_power points (11.98 DPS) | yes | The Eye of Zuldazar (19593, -2.60 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -2.60 DPS) [quest]; Drake Tooth Necklace (21531, -3.13 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 89.0 healing_power points (19.16 DPS) | yes | Virtuous Mantle (226951, +0.00 DPS, sim-verified) [quest]; Champion's Satin Mantle (227120, -5.65 DPS) [pvp]; Warlord's Satin Mantle (231631, -7.30 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 54.5 healing_power points (11.75 DPS) | yes | Drape of Recovery (272413, -2.48 DPS) [vendor]; Cloak of the Cosmos (18389, -3.18 DPS) [dungeon]; Battle Healer's Cloak (19526, -4.10 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.0 healing_power points (19.83 DPS) | yes | Robes of the Exalted (13346, -0.55 DPS, sim-verified) [dungeon]; Warlord's Satin Tunic (231632, -1.13 DPS) [vendor]; Virtuous Robe (226945, -1.76 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 46.1 healing_power points (9.93 DPS) | yes | Bracers of Mending (23129, -0.28 DPS) [dungeon]; Virtuous Bracers (226949, -0.80 DPS) [quest]; General's Satin Bracers (17619, -2.23 DPS) [pvp] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 63.6 healing_power points (13.70 DPS) | yes | Desert Bloom Gloves (20717, -0.08 DPS) [quest]; Virtuous Mitts (226950, -0.94 DPS) [vendor]; Raider Handwraps (272097, -1.23 DPS) [vendor] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 60.3 healing_power points (13.00 DPS) | yes | Wisdom of the Timbermaw (19047, -0.23 DPS) [crafted]; Virtuous Belt (226948, -0.80 DPS) [quest]; Penitent's Cinch (272394, -1.89 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 85.1 healing_power points (18.34 DPS) | yes | General's Satin Legguards (231634, -1.96 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -2.08 DPS) [pvp]; Virtuous Skirt (226946, -2.72 DPS) [quest] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 78.2 healing_power points (16.85 DPS) | yes | Virtuous Sandals (226952, -2.61 DPS) [quest]; Mooncloth Boots (15802, -4.56 DPS) [crafted]; Faith Healer's Boots (22247, -4.93 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (509.9 DPS) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Rosewine Circle (13178, -0.62 DPS) [dungeon]; Fordring's Seal (16058, -0.79 DPS) [quest] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (509.9 DPS) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Rosewine Circle (13178, -0.54 DPS) [dungeon]; Fordring's Seal (16058, -0.71 DPS) [quest] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (509.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-verified (509.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (509.9 DPS) | yes | The Lobotomizer (19324, +0.00 DPS) [rep]; Staff of Metanoia (22394, -2.37 DPS) [dungeon]; Death Speaker Scepter (2816, -2.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.1 healing_power points (6.48 DPS) | yes | Sparkling Crystal Wand (20672, -2.97 DPS) [world]; Bonecreeper Stylus (13938, -3.03 DPS) [dungeon]; Oblivion's Touch (18761, -3.51 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Hands of the Exalted Herald; waist: Whipvine Cord; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Band of Mending; finger2: Band of Piety; trinket1: Burst of Knowledge; trinket2: Darkmoon Card: Heroism; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

