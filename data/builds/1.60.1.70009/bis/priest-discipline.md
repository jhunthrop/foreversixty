# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.3. Weights run: 7.8s. Verify run: 3.9s. 150 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=2.347 ± 0.010, spirit=1.375 ± 0.005, mp5=3.102 ± 0.033, crit=0.157 ± 0.007 per rating point (14 rating = 1%, 2.200 per %), spell_haste=not significant (0.006 ± 0.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.0 healing_power points (0.93 DPS) | yes | Pristine Circlet (253949, -0.06 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.67 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 5.5 healing_power points (0.26 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.1 healing_power points (0.98 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.07 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.55 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (20428, -0.03 DPS, sim-verified) [rep]; Seer's Cape (6378, -0.09 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.11 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 28.9 healing_power points (1.34 DPS) | yes | Robe of the Moccasin (6465, -0.09 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.32 DPS) [dungeon]; Seer's Robe (2981, -0.49 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.0 healing_power points (0.98 DPS) | yes | Bright Bracers (3647, +0.00 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.63 DPS) [quest]; Seer's Cuffs (3645, -0.74 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 18.0 healing_power points (0.84 DPS) | yes | Magefist Gloves (12977, +0.00 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.15 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.18 DPS) [quest] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Sash (253925, -0.03 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.07 DPS) [crafted]; Tarantula Silk Sash (3229, -0.20 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 37.6 healing_power points (1.74 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -0.84 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.84 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.0 healing_power points (0.74 DPS) | yes | Kimbra Boots (6191, +0.00 DPS, sim-verified) [quest]; Bluegill Sandals (1560, -0.27 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 14.1 healing_power points (0.65 DPS) | yes | Band of Purification (12996, -0.27 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; Deep Fathom Ring (6463, -0.33 DPS) [dungeon] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 12.9 healing_power points (0.60 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop]; Deep Fathom Ring (6463, -0.28 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 23.5 healing_power points (1.09 DPS) | yes | Staff of Westfall (2042, -0.07 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.22 DPS) [world]; Lesser Staff of the Spire (1300, -0.44 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 4.7 healing_power points (0.22 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.03 DPS) [quest]; Dwarven Flamestick (5241, -0.09 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Keller's Girdle; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (human, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.7. Weights run: 7.9s. Verify run: 3.9s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.199 ± 0.011, spirit=1.744 ± 0.010, mp5=3.814 ± 0.010, crit=0.217 ± 0.008 per rating point (14 rating = 1%, 3.038 per %), spell_haste=not significant (0.299 ± 0.097)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Resilient Cap (14401, -0.27 DPS) [world_drop]; Holy Shroud (2721, -0.52 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.74 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.8 healing_power points (1.07 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS, sim-verified) [world_drop]; Scorn's Icy Choker (23169, -0.17 DPS) [dungeon]; Necklace of Harmony (5180, -0.24 DPS) [world] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 32.9 healing_power points (2.23 DPS) | yes | Mantle of Honor (3560, -0.08 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.45 DPS) [world_drop]; Death Speaker Mantle (6685, -0.59 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.8 healing_power points (1.54 DPS) | yes | Repairman's Cape (9605, -0.12 DPS) [quest]; Glowing Thresher Cape (6901, -0.13 DPS) [dungeon]; Prelacy Cape (7004, -0.23 DPS, sim-verified) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.05 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.64 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.4 healing_power points (1.52 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Spidertank Oilrag (9448, -0.74 DPS) [dungeon] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan: Caverndeep Burrower [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of Old (9395, -0.07 DPS, sim-verified) [world_drop]; Town Clerk's Mittens (270029, -0.14 DPS) [quest]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Resilient Cord (14406) | World drop [world_drop] | 20.2 healing_power points (1.37 DPS) | yes | Pristine Sash (253925, +0.00 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.09 DPS) [vendor]; Novice Ardent's Sash (253887, -0.16 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.1 healing_power points (3.12 DPS) | yes | Filigreed Pristine Leggings (253937, -0.05 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.81 DPS) [crafted]; Darkweave Breeches (12987, -1.37 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.4 healing_power points (2.87 DPS) | yes | Acidic Walkers (9454, +0.00 DPS, sim-verified) [dungeon]; Soggy Boots (274747, -1.33 DPS) [vendor]; Frothing Slippers (254003, -1.35 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.6 healing_power points (1.87 DPS) | yes | Sea Giant's Toe Ring (274746, -0.65 DPS) [vendor]; Black Widow Band (6199, -0.83 DPS) [world]; Darkspear Signet (272071, -0.84 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 18.3 healing_power points (1.24 DPS) | yes | Sea Giant's Toe Ring (274746, +0.00 DPS, sim-verified) [vendor]; Black Widow Band (6199, -0.20 DPS) [world]; Darkspear Signet (272071, -0.21 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.28 DPS, sim-verified) [dungeon]; Lorekeeper's Staff (212580, -0.59 DPS) [vendor]; Staff of the Friar (3415, -0.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 8.8 healing_power points (0.60 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Consecrated Wand (5244, -0.06 DPS) [quest]; Gravestone Scepter (7001, -0.24 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Pristine Gown; hands: Hotshot Pilot's Gloves; waist: Resilient Cord; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 155.3. Weights run: 9.1s. Verify run: 4.8s. 328 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.035, intellect=1.741 ± 0.027, spirit=0.595 ± 0.019, mp5=0.429 ± 0.044, crit=0.454 ± 0.014 per rating point (14 rating = 1%, 6.349 per %), spell_haste=not significant (0.671 ± 0.177)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 57.0 healing_power points (5.48 DPS) | yes | Corpseshroud (10574, -0.19 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.96 DPS) [world_drop]; Miner's Hat of the Deep (9429, -2.06 DPS) [dungeon] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 17.4 healing_power points (1.67 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.10 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.21 DPS) [dungeon] |
| shoulder | Windchaser Amice (14432) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.05 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.46 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (19532, -0.05 DPS) [rep]; Mantle of Lady Falther'ess (23178, -0.10 DPS, sim-verified) [dungeon]; Blackforge Cape (6424, -0.50 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.8 healing_power points (6.03 DPS) | yes | Death Speaker Robes (6682, -0.15 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -2.71 DPS) [crafted]; Red Mageweave Vest (10007, -3.02 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 15.7 healing_power points (1.51 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -0.17 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.17 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Handwraps (254021, -0.21 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.88 DPS) [crafted]; Truefaith Gloves (7049, -1.12 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 39.5 healing_power points (3.80 DPS) | yes | Deathmage Sash (10771, +0.00 DPS, sim-verified) [dungeon]; Sutarn's Ring (13105, -1.78 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.79 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 37.2 healing_power points (3.57 DPS) | yes | Filigreed Pristine Leggings (253937, -0.12 DPS, sim-verified) [crafted]; Aurora Pants (4044, -1.16 DPS) [world_drop]; Stormcloth Pants (10010, -1.27 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Slippers (254001, -0.10 DPS, sim-verified) [crafted]; Thoughtcast Boots (10578, -0.34 DPS) [dungeon]; Nimbus Boots (6998, -0.45 DPS) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.73 DPS) | yes | Ogremind Ring (1993, -0.39 DPS) [world_drop]; Voodoo Band (1996, -0.39 DPS) [world]; Mindbender Loop (5009, -0.44 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 16.4 healing_power points (1.57 DPS) | yes | Voodoo Band (1996, -0.07 DPS, sim-verified) [world]; Ogremind Ring (1993, -0.23 DPS) [world_drop]; Mindbender Loop (5009, -0.29 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -4.03 DPS) [dungeon]; Staff of Jordan (873, -4.07 DPS) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Souls (249395, -0.08 DPS) [crafted]; Aurora Sphere (7610, -0.27 DPS) [dungeon]; Beacon of Hope (9393, -0.31 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 12.2 healing_power points (1.18 DPS) | yes | Goblin Igniter (5253, -0.12 DPS, sim-verified) [quest]; Flash Wand (5248, -0.33 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.43 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Prodigious Shadowshard Pendant; shoulder: Windchaser Amice; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; feet: Furen's Boots; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 228.3. Weights run: 9.1s. Verify run: 5.7s. 423 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=1.665 ± 0.034, spirit=not significant (0.273 ± 0.092), mp5=1.735 ± 0.139, crit=0.662 ± 0.018 per rating point (14 rating = 1%, 9.267 per %), spell_haste=0.833 ± 0.190

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Cassandra's Grace (13102, -0.67 DPS) [world_drop]; Knight-Lieutenant's Satin Cover (220896, -0.68 DPS, sim-verified) [vendor]; Chief Architect's Monocle (11839, -0.74 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 24.4 healing_power points (2.59 DPS) | yes | Darkmoon Necklace (19303, -0.41 DPS, sim-verified) [vendor]; Gemshard Heart (17707, -0.65 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.82 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Satin Pads (220894, -0.32 DPS, sim-verified) [vendor]; Kentic Amice (11624, -0.39 DPS) [dungeon]; Rotgrip Mantle (17732, -1.17 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 24.7 healing_power points (2.62 DPS) | yes | Caretaker's Cape (19531, -0.11 DPS) [rep]; Mantle of Lady Falther'ess (23178, -0.45 DPS, sim-verified) [dungeon]; Imperial Red Cloak (8248, -0.56 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 60.2 healing_power points (6.40 DPS) | yes | Knight's Satin Armor (220892, -0.07 DPS, sim-verified) [vendor]; Robes of Insight (940, -1.54 DPS) [world_drop]; Death Speaker Robes (6682, -2.33 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 33.0 healing_power points (3.51 DPS) | yes | Aristocratic Cuffs (12546, +0.00 DPS, sim-verified) [dungeon]; Shizzle's Nozzle Wiper (11917, -1.30 DPS) [quest]; Forgotten Wraps (9433, -1.39 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 58.6 healing_power points (6.23 DPS) | yes | Gilded Gloves (254095, -0.60 DPS, sim-verified) [crafted]; Virtuous Mitts (226950, -1.21 DPS) [vendor]; Sergeant Major's Satin Gloves (220897, -2.23 DPS) [vendor] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 44.0 healing_power points (4.68 DPS) | yes | Gilded Cord (254037, -0.15 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.31 DPS) [dungeon]; Ban'thok Sash (11662, -1.45 DPS) [dungeon] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | 54.2 healing_power points (5.76 DPS) | yes | Spellshock Leggings (9484, -0.18 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -2.04 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -2.05 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 46.0 healing_power points (4.89 DPS) | yes | Sergeant Major's Satin Boots (220895, -0.46 DPS, sim-verified) [vendor]; Gilded Slippers (254001, -1.41 DPS) [crafted]; Coldstone Slippers (18697, -1.67 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.3 healing_power points (2.80 DPS) | yes | Mindseye Circle (10634, -0.68 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.89 DPS) [vendor]; Woodseed Hoop (17768, -1.21 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 21.7 healing_power points (2.31 DPS) | yes | Mindseye Circle (10634, +0.00 DPS, sim-verified) [dungeon]; Sea Giant's Toe Ring (274746, -0.40 DPS) [vendor]; Woodseed Hoop (17768, -0.72 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.64 DPS, sim-verified) [world_drop]; Uther's Strength (11302, -0.85 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.48 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.06 DPS, sim-verified) [world_drop] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -1.84 DPS) [world_drop]; Spellshifter Rod (9527, -3.16 DPS) [quest] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.7 healing_power points (2.62 DPS) | yes | Twisting Essence Jar (249456, -0.37 DPS) [crafted]; Skullspell Orb (10708, -0.50 DPS) [quest]; Beacon of Hope (9393, -0.59 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 10.8 healing_power points (1.15 DPS) | yes | Cairnstone Sliver (9654, -0.12 DPS, sim-verified) [quest]; Flash Wand (5248, -0.35 DPS) [quest]; Goblin Igniter (5253, -0.35 DPS) [quest] |

**New at 50:** neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Knight's Satin Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Darkspear Voodoo Seal; trinket2: Thunderbrew's Boot Flask; off_hand: Enthralled Sphere

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 373.4. Weights run: 8.7s. Verify run: 8.6s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.113, intellect=2.151 ± 0.058, spirit=0.337 ± 0.074, mp5=1.595 ± 0.105, crit=0.791 ± 0.032 per rating point (14 rating = 1%, 11.071 per %), spell_haste=not significant (0.866 ± 0.772)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Virtuous Crown (226947, -1.41 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Drake Tooth Necklace (21531, -0.43 DPS, sim-verified) [quest]; Lady Maye's Pendant (14558, -0.60 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.84 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -1.29 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 63.5 healing_power points (7.08 DPS) | yes | Cloak of the Cosmos (18389, -0.66 DPS, sim-verified) [dungeon]; Drape of Recovery (272413, -2.91 DPS) [vendor]; Darkspear Raider's Cloak (272063, -3.02 DPS) [vendor] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Mooncloth Robe (18486, -0.30 DPS) [crafted]; Virtuous Robe (226945, -0.75 DPS, sim-verified) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 45.7 healing_power points (5.10 DPS) | yes | Marshal's Satin Bracers (17606, -0.16 DPS) [pvp]; Bracers of Mending (23129, -0.42 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -0.52 DPS) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 82.1 healing_power points (9.15 DPS) | yes | Mooncloth Gloves (18409, -1.63 DPS) [crafted]; Hands of the Exalted Herald (12554, -1.90 DPS) [dungeon]; Desert Bloom Gloves (20717, -2.08 DPS) [quest] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 64.1 healing_power points (7.15 DPS) | yes | Whipvine Cord (18327, -0.50 DPS, sim-verified) [dungeon]; Devout Belt (16696, -0.79 DPS) [dungeon]; Virtuous Belt (226948, -1.00 DPS) [quest] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 94.6 healing_power points (10.55 DPS) | yes | Marshal's Satin Legguards (231626, -0.82 DPS) [vendor]; Virtuous Skirt (226946, -1.36 DPS, sim-verified) [quest]; Knight-Captain's Satin Legguards (227125, -1.96 DPS) [pvp] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Knight-Lieutenant's Satin Walkers (227129, -0.03 DPS) [pvp]; Incandescent Mooncloth Boots (227862, -1.02 DPS, sim-verified) [vendor] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.24 DPS) [quest]; Emerald Flame Ring (18395, -0.71 DPS) [dungeon]; Naglering (11669, -3.38 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.04 DPS) [quest]; Emerald Flame Ring (18395, -0.51 DPS) [dungeon]; Naglering (11669, -2.91 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor]; Mindtap Talisman (18371, -0.84 DPS, sim-verified) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -0.38 DPS, sim-verified) [vendor]; Briarwood Reed (12930, -1.16 DPS) [dungeon]; Blackhand's Breadth (13965, -1.92 DPS) [quest] |
| main_hand | Staff of Metanoia (22394) | Scholomance: Jandice Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -0.00 DPS) [pvp]; Redemption (22406, -0.18 DPS) [dungeon]; Hand of Edward the Odd (2243, -2.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Sparkling Crystal Wand (20672) | The Duke of Shards [world] | 24.4 healing_power points (2.72 DPS) | yes | Oblivion's Touch (18761, -0.09 DPS, sim-verified) [dungeon]; Torch of Light (279246, -0.41 DPS) [crafted]; Bonecreeper Stylus (13938, -0.53 DPS) [dungeon] |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Mooncloth Boots; finger1: Band of Piety; finger2: Band of Mending; trinket2: Royal Seal of Eldre'Thalas; main_hand: Staff of Metanoia; ranged: Sparkling Crystal Wand

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 583.9. Weights run: 5.9s. Verify run: 8.4s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.331, intellect=2.552 ± 0.068, spirit=2.112 ± 0.065, mp5=5.114 ± 0.075, crit=1.103 ± 0.047 per rating point (14 rating = 1%, 15.437 per %), spell_haste=not significant (1.751 ± 0.974)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 125.6 healing_power points (16.66 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -1.34 DPS) [pvp]; Gnomish Turban of Psychic Might (21517, -19.64 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 80.2 healing_power points (10.64 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -1.75 DPS) [world_drop]; The Eye of Zuldazar (19593, -1.90 DPS) [quest]; Lady Maye's Pendant (14558, -2.16 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -0.34 DPS) [pvp]; Argent Elite Shoulders (227888, -14.46 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 67.5 healing_power points (8.95 DPS) | yes | Cloak of the Cosmos (18389, -1.78 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -1.86 DPS) [vendor]; Frostweaver Cape (12968, -14.13 DPS, sim-verified) [dungeon] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | 124.9 healing_power points (16.57 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Alanna's Embrace (13314, -1.55 DPS) [dungeon]; Mooncloth Vest (14138, -19.02 DPS, sim-verified) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 71.4 healing_power points (9.47 DPS) | yes | Marshal's Satin Bracers (17606, -0.24 DPS) [pvp]; Virtuous Bracers (226949, -0.96 DPS) [quest]; Bracers of Mending (23129, -1.91 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 92.9 healing_power points (12.32 DPS) | yes | Virtuous Mitts (226950, -0.37 DPS) [vendor]; Devout Gloves (16692, -2.19 DPS) [dungeon]; Hands of the Exalted Herald (12554, -5.73 DPS, sim-verified) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Whipvine Cord (18327, -0.54 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -16.46 DPS, sim-verified) [crafted] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 123.7 healing_power points (16.41 DPS) | yes | Marshal's Satin Legguards (231626, -1.61 DPS) [vendor]; Virtuous Skirt (226946, -1.84 DPS) [quest]; Devout Skirt (16694, -23.16 DPS, sim-verified) [dungeon] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, -2.45 DPS) [crafted]; Faith Healer's Boots (22247, -3.14 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -32.27 DPS, sim-verified) [vendor] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.85 DPS) [dungeon]; Seal of Rivendare (13345, -1.30 DPS) [dungeon]; Naglering (11669, -25.36 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.13 DPS) [dungeon]; Seal of Rivendare (13345, -0.58 DPS) [dungeon]; Naglering (11669, -19.83 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (-1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (229.5 DPS) | yes | Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Darkspear Voodoo Seal (272059, +0.00 DPS) [vendor] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quel'dorai Channeling Rod (18311, -0.86 DPS) [dungeon]; Staff of Metanoia (22394, -1.26 DPS) [dungeon]; Hand of Edward the Odd (2243, -18.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 40.2 healing_power points (5.33 DPS) | yes | Oblivion's Touch (18761, -1.18 DPS, sim-verified) [dungeon]; Sparkling Crystal Wand (20672, -1.63 DPS) [world]; Cairnstone Sliver (9654, -2.24 DPS) [quest] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Hide of the Wild; chest: Virtuous Robe; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Virtuous Belt; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Piety; finger2: Emerald Flame Ring; trinket1: Royal Seal of Eldre'Thalas; trinket2: Serenity Field; main_hand: Staff of Hale Magefire; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.1. Weights run: 7.8s. Verify run: 3.9s. 140 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=2.347 ± 0.010, spirit=1.375 ± 0.005, mp5=3.102 ± 0.033, crit=0.157 ± 0.007 per rating point (14 rating = 1%, 2.200 per %), spell_haste=not significant (0.006 ± 0.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.0 healing_power points (0.93 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.67 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 5.5 healing_power points (0.26 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.1 healing_power points (0.98 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.06 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.55 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (20427, -0.03 DPS, sim-verified) [rep]; Seer's Cape (6378, -0.09 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.11 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 28.9 healing_power points (1.34 DPS) | yes | Robe of the Moccasin (6465, -0.08 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.32 DPS) [dungeon]; Seer's Robe (2981, -0.49 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.0 healing_power points (0.98 DPS) | yes | Tabitha's Cuffs (251486, +0.00 DPS, sim-verified) [quest]; Featherbead Bracers (15452, -0.43 DPS) [quest]; Bright Bracers (3647, -0.54 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 23.3 healing_power points (1.08 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.28 DPS) [world_drop]; Bright Gloves (3066, -0.39 DPS) [world_drop] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Sash (253925, -0.03 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.07 DPS) [crafted]; Tarantula Silk Sash (3229, -0.20 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 37.6 healing_power points (1.74 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -0.84 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.84 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.0 healing_power points (0.74 DPS) | yes | Bluegill Sandals (1560, +0.00 DPS, sim-verified) [world]; Walking Boots (4660, -0.31 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 14.1 healing_power points (0.65 DPS) | yes | Black Pearl Ring (6332, -0.07 DPS, sim-verified) [world]; Band of Purification (12996, -0.27 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Pearl Ring (6332, -0.02 DPS, sim-verified) [world]; Band of Purification (12996, -0.16 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 23.5 healing_power points (1.09 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.05 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 4.7 healing_power points (0.22 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.11 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Keller's Girdle; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.5. Weights run: 7.9s. Verify run: 3.9s. 231 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.199 ± 0.011, spirit=1.744 ± 0.010, mp5=3.814 ± 0.010, crit=0.217 ± 0.008 per rating point (14 rating = 1%, 3.038 per %), spell_haste=not significant (0.299 ± 0.097)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Resilient Cap (14401, -0.27 DPS) [world_drop]; Holy Shroud (2721, -0.51 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.74 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.8 healing_power points (1.07 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS, sim-verified) [world_drop]; Scorn's Icy Choker (23169, -0.17 DPS) [dungeon]; Necklace of Harmony (5180, -0.24 DPS) [world] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 32.9 healing_power points (2.23 DPS) | yes | Nightsky Mantle (4718, -0.08 DPS, sim-verified) [world_drop]; Ghostly Mantle (3324, -0.56 DPS) [quest]; Death Speaker Mantle (6685, -0.59 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.8 healing_power points (1.54 DPS) | yes | Battle Healer's Cloak (19529, -0.19 DPS) [rep]; Glowing Thresher Cape (6901, -0.22 DPS, sim-verified) [dungeon]; Cloak of Rot (4462, -0.35 DPS) [world] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.05 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.64 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.4 healing_power points (1.52 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.63 DPS) [quest] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan: Caverndeep Burrower [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of Old (9395, -0.08 DPS, sim-verified) [world_drop]; Blight Gloves (279877, -0.15 DPS) [quest]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 25.0 healing_power points (1.69 DPS) | yes | Resilient Cord (14406, -0.08 DPS, sim-verified) [world_drop]; Pristine Sash (253925, -0.35 DPS) [crafted]; Dreamer's Belt (4829, -0.42 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.1 healing_power points (3.12 DPS) | yes | Filigreed Pristine Leggings (253937, -0.07 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.81 DPS) [crafted]; Sacred Burial Trousers (6282, -1.31 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.4 healing_power points (2.87 DPS) | yes | Acidic Walkers (9454, +0.00 DPS, sim-verified) [dungeon]; Soggy Boots (274747, -1.33 DPS) [vendor]; Frothing Slippers (254003, -1.35 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.6 healing_power points (1.87 DPS) | yes | Sea Giant's Toe Ring (274746, -0.65 DPS) [vendor]; Black Widow Band (6199, -0.83 DPS) [world]; Darkspear Signet (272071, -0.84 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 18.3 healing_power points (1.24 DPS) | yes | Sea Giant's Toe Ring (274746, +0.00 DPS, sim-verified) [vendor]; Black Widow Band (6199, -0.20 DPS) [world]; Darkspear Signet (272071, -0.21 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.27 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (19569, -0.44 DPS) [pvp]; Lorekeeper's Staff (212580, -0.59 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 8.8 healing_power points (0.60 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Gravestone Scepter (7001, -0.24 DPS) [quest]; Wand of Decay (5252, -0.30 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Pristine Gown; hands: Hotshot Pilot's Gloves; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 153.8. Weights run: 9.1s. Verify run: 4.9s. 311 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.035, intellect=1.741 ± 0.027, spirit=0.595 ± 0.019, mp5=0.429 ± 0.044, crit=0.454 ± 0.014 per rating point (14 rating = 1%, 6.349 per %), spell_haste=not significant (0.671 ± 0.177)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 57.0 healing_power points (5.48 DPS) | yes | Corpseshroud (10574, -0.16 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.96 DPS) [world_drop]; Miner's Hat of the Deep (9429, -2.06 DPS) [dungeon] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 17.4 healing_power points (1.67 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.10 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.21 DPS) [dungeon] |
| shoulder | Windchaser Amice (14432) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.05 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.50 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (19528, -0.05 DPS) [rep]; Mantle of Lady Falther'ess (23178, -0.13 DPS, sim-verified) [dungeon]; Blackforge Cape (6424, -0.50 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.8 healing_power points (6.03 DPS) | yes | Death Speaker Robes (6682, -2.27 DPS) [dungeon]; Pristine Gown (253961, -2.71 DPS) [crafted]; Red Mageweave Vest (10007, -3.02 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 15.7 healing_power points (1.51 DPS) | yes | Mistscape Bracers (4045, -0.08 DPS, sim-verified) [dungeon]; Radiant Silver Bracers (4545, -0.17 DPS) [quest]; Enchanted Stonecloth Bracers (4979, -0.17 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Handwraps (254021, -0.16 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.88 DPS) [crafted]; Truefaith Gloves (7049, -1.12 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 39.5 healing_power points (3.80 DPS) | yes | Deathmage Sash (10771, +0.00 DPS, sim-verified) [dungeon]; Sutarn's Ring (13105, -1.78 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.79 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 37.2 healing_power points (3.57 DPS) | yes | Filigreed Pristine Leggings (253937, -0.14 DPS, sim-verified) [crafted]; Aurora Pants (4044, -1.16 DPS) [world_drop]; Stormcloth Pants (10010, -1.27 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Slippers (254001, -0.10 DPS, sim-verified) [crafted]; Boots of the Maharishi (9658, -0.24 DPS) [quest]; Thoughtcast Boots (10578, -0.34 DPS) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.73 DPS) | yes | Ogremind Ring (1993, -0.39 DPS) [world_drop]; Voodoo Band (1996, -0.39 DPS) [world]; Mindbender Loop (5009, -0.44 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 16.4 healing_power points (1.57 DPS) | yes | Voodoo Band (1996, -0.11 DPS, sim-verified) [world]; Ogremind Ring (1993, -0.23 DPS) [world_drop]; Mindbender Loop (5009, -0.29 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -4.03 DPS) [dungeon]; Staff of Jordan (873, -4.07 DPS) [world_drop] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Souls (249395, -0.42 DPS) [crafted]; Beacon of Hope (9393, -0.46 DPS, sim-verified) [dungeon]; Aurora Sphere (7610, -0.61 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 12.2 healing_power points (1.18 DPS) | yes | Goblin Igniter (5253, -0.20 DPS, sim-verified) [quest]; Flash Wand (5248, -0.33 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.43 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Prodigious Shadowshard Pendant; shoulder: Windchaser Amice; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; feet: Furen's Boots; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Prophetic Cane; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 228.4. Weights run: 9.1s. Verify run: 5.5s. 403 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=1.665 ± 0.034, spirit=not significant (0.273 ± 0.092), mp5=1.735 ± 0.139, crit=0.662 ± 0.018 per rating point (14 rating = 1%, 9.267 per %), spell_haste=0.833 ± 0.190

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 52.8 healing_power points (5.61 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Chief Architect's Monocle (11839, -0.74 DPS) [dungeon]; Cassandra's Grace (13102, -1.61 DPS, sim-verified) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 24.4 healing_power points (2.59 DPS) | yes | Gemshard Heart (17707, -0.65 DPS) [dungeon]; Darkmoon Necklace (19303, -0.79 DPS, sim-verified) [vendor]; Prodigious Shadowshard Pendant (17773, -0.82 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 41.0 healing_power points (4.36 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -1.17 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 24.7 healing_power points (2.62 DPS) | yes | Battle Healer's Cloak (19527, -0.11 DPS) [rep]; Mantle of Lady Falther'ess (23178, -0.40 DPS, sim-verified) [dungeon]; Imperial Red Cloak (8248, -0.56 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 60.2 healing_power points (6.40 DPS) | yes | Stone Guard's Satin Armor (220903, -0.06 DPS, sim-verified) [vendor]; Robes of Insight (940, -1.54 DPS) [world_drop]; Death Speaker Robes (6682, -2.33 DPS) [dungeon] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | sim-verified (15.9 DPS) | yes | Nethergeld Cuffs (254061, -0.21 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -0.62 DPS) [quest]; Forgotten Wraps (9433, -0.71 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 58.6 healing_power points (6.23 DPS) | yes | Gilded Gloves (254095, -1.10 DPS, sim-verified) [crafted]; Virtuous Mitts (226950, -1.21 DPS) [vendor]; Greenleaf Handwraps (19116, -1.95 DPS) [quest] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 44.0 healing_power points (4.68 DPS) | yes | Gilded Cord (254037, -0.42 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.31 DPS) [dungeon]; Ban'thok Sash (11662, -1.45 DPS) [dungeon] |
| legs | Stone Guard's Satin Leggings (220902) | Lady Palanseer [vendor] | 54.2 healing_power points (5.76 DPS) | yes | Spellshock Leggings (9484, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -2.04 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -2.05 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 46.0 healing_power points (4.89 DPS) | yes | First Sergeant's Satin Boots (220900, -0.44 DPS, sim-verified) [vendor]; Gilded Slippers (254001, -1.41 DPS) [crafted]; Coldstone Slippers (18697, -1.67 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 26.3 healing_power points (2.80 DPS) | yes | Mindseye Circle (10634, -0.68 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.89 DPS) [vendor]; Woodseed Hoop (17768, -1.21 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 21.7 healing_power points (2.31 DPS) | yes | Mindseye Circle (10634, -0.19 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.40 DPS) [vendor]; Woodseed Hoop (17768, -0.72 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -1.31 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.37 DPS) [quest] |
| trinket2 | Alchemists' Stone (13503) | Alchemy [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.27 DPS, sim-verified) [world_drop] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -1.84 DPS) [world_drop]; Spellshifter Rod (9527, -3.16 DPS) [quest] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.7 healing_power points (2.62 DPS) | yes | Twisting Essence Jar (249456, -0.37 DPS) [crafted]; Prophetic Cane (6803, -0.50 DPS) [quest]; Beacon of Hope (9393, -0.64 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 10.8 healing_power points (1.15 DPS) | yes | Nature's Breath (19118, -0.06 DPS, sim-verified) [quest]; Flash Wand (5248, -0.35 DPS) [quest]; Goblin Igniter (5253, -0.35 DPS) [quest] |

**New at 50:** neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Stone Guard's Satin Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket2: Alchemists' Stone; off_hand: Enthralled Sphere

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 367.7. Weights run: 8.7s. Verify run: 8.4s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.113, intellect=2.151 ± 0.058, spirit=0.337 ± 0.074, mp5=1.595 ± 0.105, crit=0.791 ± 0.032 per rating point (14 rating = 1%, 11.071 per %), spell_haste=not significant (0.866 ± 0.772)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Virtuous Crown (226947, -1.67 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Drake Tooth Necklace (21531, -0.40 DPS, sim-verified) [quest]; Lady Maye's Pendant (14558, -0.60 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.84 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -1.87 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 63.5 healing_power points (7.08 DPS) | yes | Cloak of the Cosmos (18389, -0.49 DPS, sim-verified) [dungeon]; Drape of Recovery (272413, -2.91 DPS) [vendor]; Darkspear Raider's Cloak (272063, -3.02 DPS) [vendor] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Mooncloth Robe (18486, -0.30 DPS) [crafted]; Virtuous Robe (226945, -0.70 DPS, sim-verified) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 45.7 healing_power points (5.10 DPS) | yes | General's Satin Bracers (17619, -0.16 DPS) [pvp]; Bracers of Mending (23129, -0.30 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -0.52 DPS) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 82.1 healing_power points (9.15 DPS) | yes | Mooncloth Gloves (18409, -1.63 DPS) [crafted]; Hands of the Exalted Herald (12554, -1.90 DPS) [dungeon]; Desert Bloom Gloves (20717, -2.08 DPS) [quest] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 64.1 healing_power points (7.15 DPS) | yes | Whipvine Cord (18327, -0.63 DPS, sim-verified) [dungeon]; Devout Belt (16696, -0.79 DPS) [dungeon]; Virtuous Belt (226948, -1.00 DPS) [quest] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 94.6 healing_power points (10.55 DPS) | yes | General's Satin Legguards (231634, -0.82 DPS) [vendor]; Virtuous Skirt (226946, -1.18 DPS, sim-verified) [quest]; Legionnaire's Satin Legguards (227123, -1.96 DPS) [pvp] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Walkers (231630, +0.00 DPS) [vendor]; Blood Guard's Satin Walkers (227127, -0.03 DPS) [pvp]; Incandescent Mooncloth Boots (227862, -1.14 DPS, sim-verified) [vendor] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.24 DPS) [quest]; Emerald Flame Ring (18395, -0.71 DPS) [dungeon]; Naglering (11669, -3.32 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.04 DPS) [quest]; Emerald Flame Ring (18395, -0.51 DPS) [dungeon]; Naglering (11669, -3.00 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -0.20 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -0.23 DPS, sim-verified) [dungeon]; Royal Seal of Eldre'Thalas (18469, -0.29 DPS) [quest]; Briarwood Reed (12930, -1.45 DPS) [dungeon] |
| main_hand | Staff of Metanoia (22394) | Scholomance: Jandice Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -0.00 DPS) [pvp]; Redemption (22406, -0.18 DPS) [dungeon]; Hand of Edward the Odd (2243, -2.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Sparkling Crystal Wand (20672) | The Duke of Shards [world] | 24.4 healing_power points (2.72 DPS) | yes | Oblivion's Touch (18761, +0.00 DPS, sim-verified) [dungeon]; Torch of Light (279246, -0.41 DPS) [crafted]; Bonecreeper Stylus (13938, -0.53 DPS) [dungeon] |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Mooncloth Boots; finger1: Band of Piety; finger2: Band of Mending; trinket1: Darkspear Voodoo Seal; trinket2: Serenity Field; main_hand: Staff of Metanoia; ranged: Sparkling Crystal Wand

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 580.3. Weights run: 5.9s. Verify run: 5.7s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.331, intellect=2.552 ± 0.068, spirit=2.112 ± 0.065, mp5=5.114 ± 0.075, crit=1.103 ± 0.047 per rating point (14 rating = 1%, 15.437 per %), spell_haste=not significant (1.751 ± 0.974)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 125.6 healing_power points (16.66 DPS) | yes | Gnomish Turban of Psychic Might (21517, +0.00 DPS, sim-verified) [quest]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -1.34 DPS) [pvp] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 80.2 healing_power points (10.64 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -1.75 DPS) [world_drop]; The Eye of Zuldazar (19593, -1.90 DPS) [quest]; Lady Maye's Pendant (14558, -5.09 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+9.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, -0.34 DPS) [pvp]; Argent Elite Shoulders (227888, -9.73 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 67.5 healing_power points (8.95 DPS) | yes | Cloak of the Cosmos (18389, -1.78 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -1.86 DPS) [vendor]; Frostweaver Cape (12968, -11.08 DPS, sim-verified) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Alanna's Embrace (13314, -0.52 DPS) [dungeon]; Virtuous Robe (226945, -6.03 DPS, sim-verified) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 71.4 healing_power points (9.47 DPS) | yes | General's Satin Bracers (17619, -0.24 DPS) [pvp]; Virtuous Bracers (226949, -0.96 DPS) [quest]; Bracers of Mending (23129, -2.09 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 92.9 healing_power points (12.32 DPS) | yes | Virtuous Mitts (226950, -0.37 DPS) [vendor]; Devout Gloves (16692, -2.19 DPS) [dungeon]; Hands of the Exalted Herald (12554, -4.99 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (+7.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Whipvine Cord (18327, -0.14 DPS) [dungeon]; Virtuous Belt (226948, -7.80 DPS, sim-verified) [quest] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 123.7 healing_power points (16.41 DPS) | yes | General's Satin Legguards (231634, -1.61 DPS) [vendor]; Virtuous Skirt (226946, -1.84 DPS) [quest]; Devout Skirt (16694, -19.64 DPS, sim-verified) [dungeon] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+25.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Boots (15802, -2.45 DPS) [crafted]; Faith Healer's Boots (22247, -3.14 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -25.37 DPS, sim-verified) [vendor] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.85 DPS) [dungeon]; Seal of Rivendare (13345, -1.30 DPS) [dungeon]; Naglering (11669, -20.47 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.13 DPS) [dungeon]; Seal of Rivendare (13345, -0.58 DPS) [dungeon]; Naglering (11669, -17.63 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Darkspear Voodoo Seal (272059, -0.99 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Darkspear Voodoo Seal (272059, +0.00 DPS) [vendor] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quel'dorai Channeling Rod (18311, -0.86 DPS) [dungeon]; Staff of Metanoia (22394, -1.26 DPS) [dungeon]; Hand of Edward the Odd (2243, -15.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 40.2 healing_power points (5.33 DPS) | yes | Sparkling Crystal Wand (20672, -1.63 DPS) [world]; Jaina's Firestarter (13064, -2.46 DPS) [world_drop]; Oblivion's Touch (18761, -2.72 DPS, sim-verified) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Piety; finger2: Emerald Flame Ring; trinket1: Royal Seal of Eldre'Thalas; trinket2: Serenity Field; main_hand: Staff of Hale Magefire; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

