# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.7. Weights run: 10.0s. Verify run: 32.8s. 150 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=2.347 ± 0.010, spirit=1.375 ± 0.005, mp5=3.102 ± 0.033, crit=0.157 ± 0.007 per rating point (14 rating = 1%, 2.200 per %), spell_haste=not significant (0.006 ± 0.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.0 healing_power points (0.93 DPS) | yes | Pristine Circlet (253949, -0.09 DPS) [crafted]; Flying Tiger Goggles (4368, -0.67 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 5.5 healing_power points (0.26 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.1 healing_power points (0.98 DPS) | yes | Slime-encrusted Pads (6461, -0.55 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.66 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 11.7 healing_power points (0.55 DPS) | yes | Seer's Cape (6378, -0.20 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.22 DPS) [crafted]; Sanguine Cape (14376, -0.36 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 28.9 healing_power points (1.34 DPS) | yes | Corsair's Overshirt (5202, -0.32 DPS) [dungeon]; Seer's Robe (2981, -0.49 DPS) [world_drop]; Robe of the Moccasin (6465, -1.41 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.0 healing_power points (0.98 DPS) | yes | Repurposed Hair Band (281256, -0.63 DPS) [quest]; Seer's Cuffs (3645, -0.74 DPS) [dungeon]; Bright Bracers (3647, -0.76 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 18.0 healing_power points (0.84 DPS) | yes | Bright Gloves (3066, -0.15 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.18 DPS) [quest]; Magefist Gloves (12977, -0.25 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 20.4 healing_power points (0.95 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Tarantula Silk Sash (3229, -0.27 DPS) [world]; Keller's Girdle (2911, -0.28 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 37.6 healing_power points (1.74 DPS) | yes | Filigreed Silky Leggings (253939, -0.84 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.84 DPS) [crafted]; Darkweave Breeches (12987, -1.07 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.0 healing_power points (0.74 DPS) | yes | Bluegill Sandals (1560, -0.27 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop]; Kimbra Boots (6191, -0.42 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 14.1 healing_power points (0.65 DPS) | yes | Band of Purification (12996, -0.27 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; Deep Fathom Ring (6463, -0.33 DPS) [dungeon] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 12.9 healing_power points (0.60 DPS) | yes | Band of Purification (12996, -0.26 DPS, sim-verified) [world_drop]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop]; Deep Fathom Ring (6463, -0.28 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 23.5 healing_power points (1.09 DPS) | yes | Channeler's Staff (4437, -0.22 DPS) [world]; Lesser Staff of the Spire (1300, -0.44 DPS) [world]; Staff of Westfall (2042, -0.50 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 4.7 healing_power points (0.22 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.03 DPS) [quest]; Dwarven Flamestick (5241, -0.09 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (human, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.6. Weights run: 10.1s. Verify run: 33.3s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.199 ± 0.011, spirit=1.744 ± 0.010, mp5=3.814 ± 0.010, crit=0.217 ± 0.008 per rating point (14 rating = 1%, 3.038 per %), spell_haste=not significant (0.299 ± 0.097)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Holy Shroud (2721, +0.00 DPS) [world_drop]; Resilient Cap (14401, -0.27 DPS) [world_drop]; Shadow Hood (4323, -0.74 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.8 healing_power points (1.07 DPS) | yes | Crystal Starfire Medallion (5003, -0.12 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.17 DPS) [dungeon]; Necklace of Harmony (5180, -0.24 DPS) [world] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 32.9 healing_power points (2.23 DPS) | yes | Mantle of Honor (3560, -0.36 DPS) [quest]; Nightsky Mantle (4718, -0.45 DPS) [world_drop]; Death Speaker Mantle (6685, -0.59 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.8 healing_power points (1.54 DPS) | yes | Prelacy Cape (7004, -0.09 DPS) [quest]; Repairman's Cape (9605, -0.12 DPS) [quest]; Glowing Thresher Cape (6901, -0.13 DPS) [dungeon] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Robes (6682, +0.00 DPS) [dungeon]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.64 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.4 healing_power points (1.52 DPS) | yes | Nightsky Wristbands (6407, -0.27 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Spidertank Oilrag (9448, -0.74 DPS) [dungeon] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan: Caverndeep Burrower [dungeon] | sim-verified (72.7 DPS) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Town Clerk's Mittens (270029, -0.14 DPS) [quest]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Resilient Cord (14406) | World drop [world_drop] | 20.2 healing_power points (1.37 DPS) | yes | Pristine Sash (253925, -0.03 DPS) [crafted]; Dreamer's Belt (4829, -0.09 DPS) [vendor]; Novice Ardent's Sash (253887, -0.16 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.1 healing_power points (3.12 DPS) | yes | Filigreed Pristine Leggings (253937, -0.54 DPS) [crafted]; Earthen Leggings (253999, -0.81 DPS) [crafted]; Darkweave Breeches (12987, -1.37 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.4 healing_power points (2.87 DPS) | yes | Acidic Walkers (9454, -1.21 DPS) [dungeon]; Soggy Boots (274747, -1.33 DPS) [vendor]; Frothing Slippers (254003, -1.35 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.6 healing_power points (1.87 DPS) | yes | Sea Giant's Toe Ring (274746, -0.65 DPS) [vendor]; Black Widow Band (6199, -0.83 DPS) [world]; Darkspear Signet (272071, -0.84 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 18.3 healing_power points (1.24 DPS) | yes | Sea Giant's Toe Ring (274746, -0.02 DPS) [vendor]; Black Widow Band (6199, -0.20 DPS) [world]; Darkspear Signet (272071, -0.21 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (72.7 DPS) | yes | Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Lorekeeper's Staff (212580, -0.59 DPS) [vendor]; Staff of the Friar (3415, -0.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 8.8 healing_power points (0.60 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Consecrated Wand (5244, -0.06 DPS) [quest]; Gravestone Scepter (7001, -0.24 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Pristine Gown; hands: Hotshot Pilot's Gloves; waist: Resilient Cord; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 163.7. Weights run: 11.7s. Verify run: 79.6s. 328 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.031, intellect=2.090 ± 0.030, spirit=0.746 ± 0.020, mp5=0.672 ± 0.044, crit=0.467 ± 0.014 per rating point (14 rating = 1%, 6.542 per %), spell_haste=-2.645 ± 0.368

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 64.3 healing_power points (5.92 DPS) | yes | Miner's Hat of the Deep (9429, -1.96 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -2.21 DPS) [dungeon]; Corpseshroud (10574, -3.36 DPS, sim-verified) [dungeon] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 20.9 healing_power points (1.93 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.10 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.21 DPS) [dungeon] |
| shoulder | Windchaser Amice (14432) (or Inquisitor's Shawl (19507)) | World drop [world_drop] | 27.2 healing_power points (2.50 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.04 DPS) [dungeon]; Batwing Mantle (6697, -0.04 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 27.8 healing_power points (2.56 DPS) | yes | Darkspear Raider's Cloak (272077, +0.00 DPS, sim-verified) [vendor]; Caretaker's Cape (19532, -0.56 DPS) [rep]; Blackforge Cape (6424, -0.75 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.0 healing_power points (5.90 DPS) | yes | Red Mageweave Vest (10007, -2.43 DPS) [crafted]; Pristine Gown (253961, -2.43 DPS) [crafted]; Death Speaker Robes (6682, -10.02 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 18.8 healing_power points (1.73 DPS) | yes | Aurora Bracers (4043, -0.19 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.19 DPS) [quest]; Mistscape Bracers (4045, -0.44 DPS, sim-verified) [dungeon] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (163.7 DPS) | yes | Earthen Silk Gloves (254017, -1.12 DPS) [crafted]; Town Clerk's Mittens (270029, -1.21 DPS) [quest]; Gilded Handwraps (254021, -2.48 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.2 healing_power points (3.98 DPS) | yes | Sutarn's Ring (13105, -1.64 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.67 DPS) [quest]; Deathmage Sash (10771, -2.91 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (163.7 DPS) | yes | Aurora Pants (4044, +0.00 DPS) [world_drop]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -2.76 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.6 healing_power points (3.47 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.42 DPS) [dungeon]; Kodo Rustler Boots (15697, -1.58 DPS) [quest] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 19.9 healing_power points (1.83 DPS) | yes | Ogremind Ring (1993, -0.28 DPS) [world_drop]; Voodoo Band (1996, -0.28 DPS) [world]; Mindbender Loop (5009, -0.34 DPS) [world_drop] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.66 DPS) | yes | Ogremind Ring (1993, -0.10 DPS) [world_drop]; Voodoo Band (1996, -0.10 DPS) [world]; Mindbender Loop (5009, -0.17 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (163.7 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (163.7 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -3.38 DPS) [dungeon]; Staff of Jordan (873, -3.39 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 25.0 healing_power points (2.30 DPS) | yes | Aurora Sphere (7610, -0.68 DPS) [dungeon]; Orb of Souls (249395, -0.69 DPS) [crafted]; Orb of Lorica (11262, -2.03 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 14.8 healing_power points (1.36 DPS) | yes | Goblin Igniter (5253, -0.37 DPS, sim-verified) [quest]; Flash Wand (5248, -0.39 DPS) [quest]; Starfaller (13063, -0.59 DPS) [world_drop] |

**New at 40:** head: Papal Fez; neck: Prodigious Shadowshard Pendant; shoulder: Windchaser Amice; back: Mantle of Lady Falther'ess; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger2: Sea Giant's Toe Ring; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 234.8. Weights run: 11.9s. Verify run: 48.3s. 423 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.010, intellect=1.734 ± 0.047, spirit=not significant (0.300 ± 0.088), mp5=1.779 ± 0.153, crit=0.662 ± 0.018 per rating point (14 rating = 1%, 9.268 per %), spell_haste=not significant (1.685 ± 0.479)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Satin Cover (220896, +0.00 DPS) [vendor]; Chief Architect's Monocle (11839, -0.68 DPS) [dungeon]; Cassandra's Grace (13102, -0.79 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 25.5 healing_power points (2.68 DPS) | yes | Darkmoon Necklace (19303, -0.46 DPS) [vendor]; Gemshard Heart (17707, -0.67 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.86 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | sim-verified (234.8 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Kentic Amice (11624, -0.35 DPS) [dungeon]; Rotgrip Mantle (17732, -1.10 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 25.8 healing_power points (2.71 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.12 DPS) [dungeon]; Caretaker's Cape (19531, -0.21 DPS) [rep]; Imperial Red Cloak (8248, -0.58 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 60.4 healing_power points (6.36 DPS) | yes | Knight's Satin Armor (220892, -0.55 DPS) [vendor]; Robes of Insight (940, -1.32 DPS) [world_drop]; Death Speaker Robes (6682, -2.25 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 33.6 healing_power points (3.54 DPS) | yes | Aristocratic Cuffs (12546, -0.61 DPS) [dungeon]; Shizzle's Nozzle Wiper (11917, -1.26 DPS) [quest]; Forgotten Wraps (9433, -1.35 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 60.1 healing_power points (6.33 DPS) | yes | Gilded Gloves (254095, -1.20 DPS) [crafted]; Virtuous Mitts (226950, -1.23 DPS) [vendor]; Sergeant Major's Satin Gloves (220897, -2.34 DPS) [vendor] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 44.6 healing_power points (4.69 DPS) | yes | Gilded Cord (254037, -0.73 DPS) [crafted]; Dawnspire Cord (12466, -1.23 DPS) [dungeon]; Ban'thok Sash (11662, -1.42 DPS) [dungeon] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | 55.1 healing_power points (5.80 DPS) | yes | Spellshock Leggings (9484, -1.56 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.95 DPS) [dungeon]; Pristine Leggings (253987, -2.05 DPS) [crafted] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 46.6 healing_power points (4.90 DPS) | yes | Sergeant Major's Satin Boots (220895, -0.61 DPS) [vendor]; Gilded Slippers (254001, -1.40 DPS) [crafted]; Coldstone Slippers (18697, -1.60 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.5 healing_power points (2.89 DPS) | yes | Mindseye Circle (10634, -0.71 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -1.00 DPS) [vendor]; Woodseed Hoop (17768, -1.25 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 22.3 healing_power points (2.35 DPS) | yes | Mindseye Circle (10634, -0.16 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.46 DPS) [vendor]; Woodseed Hoop (17768, -0.71 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (234.8 DPS) | yes | Uther's Strength (11302, -0.91 DPS) [world_drop]; Ankh of Life (1713, -2.43 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.49 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (234.8 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -1.58 DPS) [world_drop]; Spellshifter Rod (9527, -2.96 DPS) [quest]; Radiant Staff (249453, -3.69 DPS) [crafted] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.8 healing_power points (2.71 DPS) | yes | Beacon of Hope (9393, -0.27 DPS) [dungeon]; Twisting Essence Jar (249456, -0.46 DPS) [crafted]; Skullspell Orb (10708, -0.52 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 11.3 healing_power points (1.19 DPS) | yes | Cairnstone Sliver (9654, -0.12 DPS) [quest]; Flash Wand (5248, -0.36 DPS) [quest]; Goblin Igniter (5253, -0.36 DPS) [quest] |

**New at 50:** neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Knight's Satin Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Darkspear Voodoo Seal; trinket2: Thunderbrew's Boot Flask; off_hand: Enthralled Sphere

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 384.3. Weights run: 11.2s. Verify run: 94.8s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.163, intellect=1.782 ± 0.053, spirit=0.888 ± 0.067, mp5=1.464 ± 0.099, crit=0.682 ± 0.034 per rating point (14 rating = 1%, 9.543 per %), spell_haste=-9.207 ± 1.040

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 87.0 healing_power points (11.33 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.33 DPS) [pvp]; Devout Crown (16693, -9.70 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (384.4 DPS) | yes | The Eye of Zuldazar (19593, -0.34 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -0.34 DPS) [quest]; Drake Tooth Necklace (21531, -5.49 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 76.7 healing_power points (9.98 DPS) | yes | Field Marshal's Satin Mantle (231628, -1.44 DPS) [pvp]; Lieutenant Commander's Satin Mantle (227119, -1.54 DPS) [pvp]; Devout Mantle (16695, -2.57 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 59.8 healing_power points (7.79 DPS) | yes | Drape of Recovery (272413, -2.99 DPS) [vendor]; Cloak of the Cosmos (18389, -3.00 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272063, -3.38 DPS) [vendor] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-verified (384.4 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.22 DPS) [dungeon]; Truefaith Vestments (14154, -4.48 DPS, sim-verified) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 48.3 healing_power points (6.28 DPS) | yes | Bracers of Mending (23129, +0.00 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -0.58 DPS) [quest]; Marshal's Satin Bracers (17606, -0.72 DPS) [pvp] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-verified (384.4 DPS) | yes | Desert Bloom Gloves (20717, -0.32 DPS) [quest]; Mooncloth Gloves (18409, -0.69 DPS) [crafted]; Raider Handwraps (272097, -1.75 DPS, sim-verified) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 63.9 healing_power points (8.32 DPS) | yes | Virtuous Belt (226948, -0.93 DPS, sim-verified) [quest]; Whipvine Cord (18327, -1.06 DPS) [dungeon]; Devout Belt (16696, -1.21 DPS) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 86.4 healing_power points (11.25 DPS) | yes | Marshal's Satin Legguards (231626, -0.26 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.94 DPS) [pvp]; Virtuous Skirt (226946, -5.65 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (384.4 DPS) | yes | Marshal's Satin Walkers (231627, -0.36 DPS) [vendor]; Mooncloth Boots (15802, -0.44 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -11.51 DPS, sim-verified) [vendor] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (384.4 DPS) | yes | Band of Piety (22681, -0.41 DPS) [quest]; Emerald Flame Ring (18395, -0.69 DPS) [dungeon]; Naglering (11669, -13.13 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (384.4 DPS) | yes | Band of Piety (22681, -0.03 DPS) [quest]; Emerald Flame Ring (18395, -0.32 DPS) [dungeon]; Naglering (11669, -9.20 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (384.4 DPS) | yes | Mindtap Talisman (18371, +0.00 DPS, sim-verified) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -1.28 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (384.4 DPS) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Serenity Field (272439, -6.62 DPS, sim-verified) [vendor] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (384.4 DPS) | yes | Staff of Metanoia (22394, -0.44 DPS) [dungeon]; Death Speaker Scepter (2816, -1.13 DPS) [dungeon]; Hand of Edward the Odd (2243, -10.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.8 healing_power points (3.48 DPS) | yes | Oblivion's Touch (18761, -0.93 DPS) [dungeon]; Bonecreeper Stylus (13938, -1.13 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.57 DPS, sim-verified) [world] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Bracers of Hope; hands: Hands of the Exalted Herald; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 636.9. Weights run: 7.5s. Verify run: 32.7s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.160, intellect=1.450 ± 0.037, spirit=1.220 ± 0.036, mp5=2.526 ± 0.040, crit=0.634 ± 0.029 per rating point (14 rating = 1%, 8.877 per %), spell_haste=not significant (-1.555 ± 0.684)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 86.1 healing_power points (20.21 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.67 DPS) [pvp]; Devout Crown (16693, -3.04 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.9 healing_power points (12.42 DPS) | yes | The Eye of Zuldazar (19593, -1.75 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.75 DPS) [quest]; Drake Tooth Necklace (21531, -2.23 DPS) [quest] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (637.0 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 56.5 healing_power points (13.27 DPS) | yes | Cloak of the Cosmos (18389, -3.42 DPS) [dungeon]; Drape of Recovery (272413, -3.62 DPS) [vendor]; Caretaker's Cape (19530, -4.87 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.5 healing_power points (21.73 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.90 DPS) [dungeon]; Virtuous Robe (226945, -1.21 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 48.6 healing_power points (11.41 DPS) | yes | Bracers of Mending (23129, -0.39 DPS) [dungeon]; Virtuous Bracers (226949, -0.97 DPS) [quest]; Marshal's Satin Bracers (17606, -2.08 DPS) [pvp] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 66.5 healing_power points (15.61 DPS) | yes | Desert Bloom Gloves (20717, -0.50 DPS) [quest]; Raider Handwraps (272097, -0.79 DPS) [vendor]; Virtuous Mitts (226950, -1.07 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 62.3 healing_power points (14.63 DPS) | yes | Whipvine Cord (18327, -0.72 DPS) [dungeon]; Virtuous Belt (226948, -0.88 DPS) [quest]; Penitent's Cinch (272394, -2.35 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 86.2 healing_power points (20.23 DPS) | yes | Marshal's Satin Legguards (231626, -1.34 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -1.72 DPS) [pvp]; Virtuous Skirt (226946, -2.35 DPS) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (637.0 DPS) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -1.91 DPS) [crafted]; Faith Healer's Boots (22247, -2.41 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.99 DPS) [quest]; Emerald Flame Ring (18395, -1.48 DPS) [dungeon]; Rosewine Circle (13178, -1.61 DPS) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.55 DPS) [quest]; Emerald Flame Ring (18395, -1.04 DPS) [dungeon]; Rosewine Circle (13178, -1.17 DPS) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, -1.22 DPS) [vendor]; Briarwood Reed (12930, -3.31 DPS) [dungeon]; Mindtap Talisman (18371, -3.60 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (637.0 DPS) | yes | Darkspear Voodoo Seal (272061, -0.97 DPS) [vendor]; Briarwood Reed (12930, -3.05 DPS) [dungeon]; Mindtap Talisman (18371, -3.34 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -1.89 DPS) [dungeon]; Death Speaker Scepter (2816, -2.97 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -3.63 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.4 healing_power points (7.14 DPS) | yes | Sparkling Crystal Wand (20672, -2.91 DPS) [world]; Bonecreeper Stylus (13938, -3.20 DPS) [dungeon]; Oblivion's Touch (18761, -3.40 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Hands of the Exalted Herald; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Band of Piety; trinket1: Royal Seal of Eldre'Thalas; trinket2: Serenity Field; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.5. Weights run: 10.0s. Verify run: 32.5s. 140 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=2.347 ± 0.010, spirit=1.375 ± 0.005, mp5=3.102 ± 0.033, crit=0.157 ± 0.007 per rating point (14 rating = 1%, 2.200 per %), spell_haste=not significant (0.006 ± 0.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.0 healing_power points (0.93 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.67 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 5.5 healing_power points (0.26 DPS) | yes | Roadwatcher's Confidence (281265, -0.16 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.1 healing_power points (0.98 DPS) | yes | Slime-encrusted Pads (6461, -0.55 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.70 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 11.7 healing_power points (0.55 DPS) | yes | Seer's Cape (6378, -0.20 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.22 DPS) [crafted]; Sanguine Cape (14376, -0.23 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 28.9 healing_power points (1.34 DPS) | yes | Corsair's Overshirt (5202, -0.32 DPS) [dungeon]; Seer's Robe (2981, -0.49 DPS) [world_drop]; Robe of the Moccasin (6465, -1.32 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.0 healing_power points (0.98 DPS) | yes | Featherbead Bracers (15452, -0.43 DPS) [quest]; Tabitha's Cuffs (251486, -0.50 DPS, sim-verified) [quest]; Bright Bracers (3647, -0.54 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 23.3 healing_power points (1.08 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.28 DPS) [world_drop]; Bright Gloves (3066, -0.39 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 20.4 healing_power points (0.95 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Tarantula Silk Sash (3229, -0.27 DPS) [world]; Keller's Girdle (2911, -0.37 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 37.6 healing_power points (1.74 DPS) | yes | Filigreed Silky Leggings (253939, -0.84 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.84 DPS) [crafted]; Darkweave Breeches (12987, -1.14 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.0 healing_power points (0.74 DPS) | yes | Walking Boots (4660, -0.31 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop]; Bluegill Sandals (1560, -0.75 DPS, sim-verified) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 14.1 healing_power points (0.65 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Band of Purification (12996, -0.27 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 12.9 healing_power points (0.60 DPS) | yes | Loop of Sacrifice (281673, -0.05 DPS, sim-verified) [quest]; Band of Purification (12996, -0.22 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 23.5 healing_power points (1.09 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.05 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 4.7 healing_power points (0.22 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.11 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 72.4. Weights run: 10.1s. Verify run: 33.5s. 231 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.199 ± 0.011, spirit=1.744 ± 0.010, mp5=3.814 ± 0.010, crit=0.217 ± 0.008 per rating point (14 rating = 1%, 3.038 per %), spell_haste=not significant (0.299 ± 0.097)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Holy Shroud (2721, +0.00 DPS) [world_drop]; Resilient Cap (14401, -0.27 DPS) [world_drop]; Shadow Hood (4323, -0.74 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.8 healing_power points (1.07 DPS) | yes | Crystal Starfire Medallion (5003, -0.12 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.17 DPS) [dungeon]; Necklace of Harmony (5180, -0.24 DPS) [world] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 32.9 healing_power points (2.23 DPS) | yes | Nightsky Mantle (4718, -0.45 DPS) [world_drop]; Ghostly Mantle (3324, -0.56 DPS) [quest]; Death Speaker Mantle (6685, -0.59 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.8 healing_power points (1.54 DPS) | yes | Glowing Thresher Cape (6901, -0.13 DPS) [dungeon]; Battle Healer's Cloak (19529, -0.19 DPS) [rep]; Cloak of Rot (4462, -0.35 DPS) [world] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Robes (6682, +0.00 DPS) [dungeon]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.64 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.4 healing_power points (1.52 DPS) | yes | Nightsky Wristbands (6407, -0.27 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.63 DPS) [quest] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan: Caverndeep Burrower [dungeon] | sim-verified (72.4 DPS) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Blight Gloves (279877, -0.15 DPS) [quest]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 25.0 healing_power points (1.69 DPS) | yes | Resilient Cord (14406, -0.33 DPS) [world_drop]; Pristine Sash (253925, -0.35 DPS) [crafted]; Dreamer's Belt (4829, -0.42 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.1 healing_power points (3.12 DPS) | yes | Filigreed Pristine Leggings (253937, -0.54 DPS) [crafted]; Earthen Leggings (253999, -0.81 DPS) [crafted]; Sacred Burial Trousers (6282, -1.31 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.4 healing_power points (2.87 DPS) | yes | Acidic Walkers (9454, -1.21 DPS) [dungeon]; Soggy Boots (274747, -1.33 DPS) [vendor]; Frothing Slippers (254003, -1.35 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.6 healing_power points (1.87 DPS) | yes | Sea Giant's Toe Ring (274746, -0.65 DPS) [vendor]; Black Widow Band (6199, -0.83 DPS) [world]; Darkspear Signet (272071, -0.84 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 18.3 healing_power points (1.24 DPS) | yes | Sea Giant's Toe Ring (274746, -0.02 DPS) [vendor]; Black Widow Band (6199, -0.20 DPS) [world]; Darkspear Signet (272071, -0.21 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (72.4 DPS) | yes | Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Advisor's Gnarled Staff (19569, -0.44 DPS) [pvp]; Lorekeeper's Staff (212580, -0.59 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 8.8 healing_power points (0.60 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Gravestone Scepter (7001, -0.24 DPS) [quest]; Wand of Decay (5252, -0.30 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Pristine Gown; hands: Hotshot Pilot's Gloves; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 162.2. Weights run: 11.7s. Verify run: 80.4s. 311 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.031, intellect=2.090 ± 0.030, spirit=0.746 ± 0.020, mp5=0.672 ± 0.044, crit=0.467 ± 0.014 per rating point (14 rating = 1%, 6.542 per %), spell_haste=-2.645 ± 0.368

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 64.3 healing_power points (5.92 DPS) | yes | Miner's Hat of the Deep (9429, -1.96 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -2.21 DPS) [dungeon]; Corpseshroud (10574, -3.01 DPS, sim-verified) [dungeon] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 20.9 healing_power points (1.93 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.10 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.21 DPS) [dungeon] |
| shoulder | Windchaser Amice (14432) (or Inquisitor's Shawl (19507)) | World drop [world_drop] | 27.2 healing_power points (2.50 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.04 DPS) [dungeon]; Batwing Mantle (6697, -0.04 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 27.8 healing_power points (2.56 DPS) | yes | Darkspear Raider's Cloak (272077, +0.00 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19528, -0.56 DPS) [rep]; Blackforge Cape (6424, -0.75 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.0 healing_power points (5.90 DPS) | yes | Red Mageweave Vest (10007, -2.43 DPS) [crafted]; Pristine Gown (253961, -2.43 DPS) [crafted]; Death Speaker Robes (6682, -8.98 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 18.8 healing_power points (1.73 DPS) | yes | Radiant Silver Bracers (4545, -0.19 DPS) [quest]; Enchanted Stonecloth Bracers (4979, -0.19 DPS) [quest]; Mistscape Bracers (4045, -0.31 DPS, sim-verified) [dungeon] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Silk Gloves (254017, -1.12 DPS) [crafted]; Truefaith Gloves (7049, -1.37 DPS) [crafted]; Gilded Handwraps (254021, -1.88 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 43.2 healing_power points (3.98 DPS) | yes | Sutarn's Ring (13105, -1.64 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.67 DPS) [quest]; Deathmage Sash (10771, -2.67 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aurora Pants (4044, +0.00 DPS) [world_drop]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -1.91 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.6 healing_power points (3.47 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Boots of the Maharishi (9658, -1.32 DPS) [quest]; Thoughtcast Boots (10578, -1.42 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 19.9 healing_power points (1.83 DPS) | yes | Ogremind Ring (1993, -0.28 DPS) [world_drop]; Voodoo Band (1996, -0.28 DPS) [world]; Mindbender Loop (5009, -0.34 DPS) [world_drop] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.66 DPS) | yes | Ogremind Ring (1993, -0.10 DPS) [world_drop]; Voodoo Band (1996, -0.10 DPS) [world]; Mindbender Loop (5009, -0.17 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -3.38 DPS) [dungeon]; Staff of Jordan (873, -3.39 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (162.2 DPS) | yes | Aurora Sphere (7610, -0.68 DPS) [dungeon]; Orb of Souls (249395, -0.69 DPS) [crafted]; Prophetic Cane (6803, -1.69 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 14.8 healing_power points (1.36 DPS) | yes | Goblin Igniter (5253, -0.38 DPS, sim-verified) [quest]; Flash Wand (5248, -0.39 DPS) [quest]; Starfaller (13063, -0.59 DPS) [world_drop] |

**New at 40:** head: Papal Fez; neck: Prodigious Shadowshard Pendant; shoulder: Windchaser Amice; back: Mantle of Lady Falther'ess; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger2: Sea Giant's Toe Ring; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 238.2. Weights run: 11.9s. Verify run: 46.8s. 403 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.010, intellect=1.734 ± 0.047, spirit=not significant (0.300 ± 0.088), mp5=1.779 ± 0.153, crit=0.662 ± 0.018 per rating point (14 rating = 1%, 9.268 per %), spell_haste=not significant (1.685 ± 0.479)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 54.2 healing_power points (5.70 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Chief Architect's Monocle (11839, -0.20 DPS, sim-verified) [dungeon]; Cassandra's Grace (13102, -0.79 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 25.5 healing_power points (2.68 DPS) | yes | Gemshard Heart (17707, -0.67 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.86 DPS) [quest]; Darkmoon Necklace (19303, -5.68 DPS, sim-verified) [vendor] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 41.7 healing_power points (4.39 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Kentic Amice (11624, -0.66 DPS, sim-verified) [dungeon]; Rotgrip Mantle (17732, -1.10 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 25.8 healing_power points (2.71 DPS) | yes | Battle Healer's Cloak (19527, -0.21 DPS) [rep]; Imperial Red Cloak (8248, -0.58 DPS) [world_drop]; Mantle of Lady Falther'ess (23178, -1.21 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 60.4 healing_power points (6.36 DPS) | yes | Robes of Insight (940, -1.32 DPS) [world_drop]; Death Speaker Robes (6682, -2.25 DPS) [dungeon]; Stone Guard's Satin Armor (220903, -3.71 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 33.6 healing_power points (3.54 DPS) | yes | Aristocratic Cuffs (12546, -0.33 DPS, sim-verified) [dungeon]; Shizzle's Nozzle Wiper (11917, -1.26 DPS) [quest]; Forgotten Wraps (9433, -1.35 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 60.1 healing_power points (6.33 DPS) | yes | Virtuous Mitts (226950, -1.23 DPS) [vendor]; Greenleaf Handwraps (19116, -2.01 DPS) [quest]; Gilded Gloves (254095, -7.65 DPS, sim-verified) [crafted] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 44.6 healing_power points (4.69 DPS) | yes | Dawnspire Cord (12466, -1.23 DPS) [dungeon]; Ban'thok Sash (11662, -1.42 DPS) [dungeon]; Gilded Cord (254037, -5.10 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (238.2 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.39 DPS) [dungeon]; Pristine Leggings (253987, -0.49 DPS) [crafted]; Stone Guard's Satin Leggings (220902, -2.60 DPS, sim-verified) [vendor] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 46.6 healing_power points (4.90 DPS) | yes | Gilded Slippers (254001, -1.40 DPS) [crafted]; Coldstone Slippers (18697, -1.60 DPS) [dungeon]; First Sergeant's Satin Boots (220900, -4.08 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.5 healing_power points (2.89 DPS) | yes | Mindseye Circle (10634, -0.71 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -1.00 DPS) [vendor]; Woodseed Hoop (17768, -1.25 DPS) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 22.3 healing_power points (2.35 DPS) | yes | Sea Giant's Toe Ring (274746, -0.46 DPS) [vendor]; Woodseed Hoop (17768, -0.71 DPS) [quest]; Mindseye Circle (10634, -2.39 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.44 DPS, sim-verified) [quest]; Ankh of Life (1713, -1.31 DPS) [world_drop] |
| trinket2 | Alchemists' Stone (13503) | Alchemy [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -2.62 DPS, sim-verified) [world_drop] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -1.58 DPS) [world_drop]; Spellshifter Rod (9527, -2.96 DPS) [quest] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.8 healing_power points (2.71 DPS) | yes | Beacon of Hope (9393, +0.00 DPS, sim-verified) [dungeon]; Twisting Essence Jar (249456, -0.46 DPS) [crafted]; Prophetic Cane (6803, -0.52 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 11.3 healing_power points (1.19 DPS) | yes | Nature's Breath (19118, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.36 DPS) [quest]; Goblin Igniter (5253, -0.36 DPS) [quest] |

**New at 50:** neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; wrist: Nethergeld Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Spellshock Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket2: Alchemists' Stone; off_hand: Enthralled Sphere

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 382.9. Weights run: 11.2s. Verify run: 60.6s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.163, intellect=1.782 ± 0.053, spirit=0.888 ± 0.067, mp5=1.464 ± 0.099, crit=0.682 ± 0.034 per rating point (14 rating = 1%, 9.543 per %), spell_haste=-9.207 ± 1.040

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 87.0 healing_power points (11.33 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.33 DPS) [pvp]; Devout Crown (16693, -4.39 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | The Eye of Zuldazar (19593, -0.34 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -0.34 DPS) [quest]; Drake Tooth Necklace (21531, -5.71 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 76.7 healing_power points (9.98 DPS) | yes | Devout Mantle (16695, +0.00 DPS, sim-verified) [dungeon]; Warlord's Satin Mantle (231631, -1.44 DPS) [pvp]; Champion's Satin Mantle (227120, -1.54 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 59.8 healing_power points (7.79 DPS) | yes | Cloak of the Cosmos (18389, -2.62 DPS, sim-verified) [dungeon]; Drape of Recovery (272413, -2.99 DPS) [vendor]; Darkspear Raider's Cloak (272063, -3.38 DPS) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | sim-verified (+10.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.07 DPS) [dungeon]; Virtuous Robe (226945, -10.74 DPS, sim-verified) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 48.3 healing_power points (6.28 DPS) | yes | Bracers of Mending (23129, +0.00 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -0.58 DPS) [quest]; General's Satin Bracers (17619, -0.72 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 72.1 healing_power points (9.39 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS, sim-verified) [dungeon]; Desert Bloom Gloves (20717, -1.01 DPS) [quest]; Mooncloth Gloves (18409, -1.38 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Whipvine Cord (18327, -0.12 DPS) [dungeon]; Devout Belt (16696, -0.28 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -7.33 DPS, sim-verified) [crafted] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 86.4 healing_power points (11.25 DPS) | yes | Virtuous Skirt (226946, +0.00 DPS, sim-verified) [quest]; General's Satin Legguards (231634, -0.26 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.94 DPS) [pvp] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+11.9 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Walkers (231630, -0.36 DPS) [vendor]; Mooncloth Boots (15802, -0.44 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -11.94 DPS, sim-verified) [vendor] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.41 DPS) [quest]; Emerald Flame Ring (18395, -0.69 DPS) [dungeon]; Naglering (11669, -8.95 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.03 DPS) [quest]; Emerald Flame Ring (18395, -0.32 DPS) [dungeon]; Naglering (11669, -7.77 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, +0.00 DPS, sim-verified) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -1.28 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -0.44 DPS) [dungeon]; Death Speaker Scepter (2816, -1.13 DPS) [dungeon]; Hand of Edward the Odd (2243, -8.07 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.8 healing_power points (3.48 DPS) | yes | Sparkling Crystal Wand (20672, -0.70 DPS, sim-verified) [world]; Oblivion's Touch (18761, -0.93 DPS) [dungeon]; Bonecreeper Stylus (13938, -1.13 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Virtuous Belt; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Darkspear Voodoo Seal; trinket2: Royal Seal of Eldre'Thalas; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 630.5. Weights run: 7.5s. Verify run: 31.2s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.160, intellect=1.450 ± 0.037, spirit=1.220 ± 0.036, mp5=2.526 ± 0.040, crit=0.634 ± 0.029 per rating point (14 rating = 1%, 8.877 per %), spell_haste=not significant (-1.555 ± 0.684)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 86.1 healing_power points (20.21 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.67 DPS) [pvp]; Devout Crown (16693, -3.04 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.9 healing_power points (12.42 DPS) | yes | The Eye of Zuldazar (19593, -1.75 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.75 DPS) [quest]; Drake Tooth Necklace (21531, -2.23 DPS) [quest] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (630.6 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 56.5 healing_power points (13.27 DPS) | yes | Cloak of the Cosmos (18389, -3.42 DPS) [dungeon]; Drape of Recovery (272413, -3.62 DPS) [vendor]; Battle Healer's Cloak (19526, -4.87 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.5 healing_power points (21.73 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.90 DPS) [dungeon]; Virtuous Robe (226945, -1.21 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 48.6 healing_power points (11.41 DPS) | yes | Bracers of Mending (23129, -0.39 DPS) [dungeon]; Virtuous Bracers (226949, -0.97 DPS) [quest]; General's Satin Bracers (17619, -2.08 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (630.6 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Desert Bloom Gloves (20717, +0.00 DPS) [quest]; Virtuous Mitts (226950, -0.29 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 62.3 healing_power points (14.63 DPS) | yes | Whipvine Cord (18327, -0.72 DPS) [dungeon]; Virtuous Belt (226948, -0.88 DPS) [quest]; Penitent's Cinch (272394, -2.35 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 86.2 healing_power points (20.23 DPS) | yes | General's Satin Legguards (231634, -1.34 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -1.72 DPS) [pvp]; Virtuous Skirt (226946, -2.35 DPS) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (630.6 DPS) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -1.91 DPS) [crafted]; Faith Healer's Boots (22247, -2.41 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.99 DPS) [quest]; Emerald Flame Ring (18395, -1.48 DPS) [dungeon]; Rosewine Circle (13178, -1.61 DPS) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, -0.55 DPS) [quest]; Emerald Flame Ring (18395, -1.04 DPS) [dungeon]; Rosewine Circle (13178, -1.17 DPS) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, -1.22 DPS) [vendor]; Briarwood Reed (12930, -3.31 DPS) [dungeon]; Mindtap Talisman (18371, -3.60 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (630.6 DPS) | yes | Darkspear Voodoo Seal (272061, -0.97 DPS) [vendor]; Briarwood Reed (12930, -3.05 DPS) [dungeon]; Mindtap Talisman (18371, -3.34 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -1.89 DPS) [dungeon]; Death Speaker Scepter (2816, -2.97 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -3.63 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.4 healing_power points (7.14 DPS) | yes | Sparkling Crystal Wand (20672, -2.91 DPS) [world]; Bonecreeper Stylus (13938, -3.20 DPS) [dungeon]; Oblivion's Touch (18761, -3.40 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Band of Piety; trinket1: Royal Seal of Eldre'Thalas; trinket2: Serenity Field; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

