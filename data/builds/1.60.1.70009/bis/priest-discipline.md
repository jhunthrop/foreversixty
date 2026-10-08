# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 28.5. Weights run: 9.1s. Verify run: 4.3s. 150 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=1.947 ± 0.013, spirit=1.465 ± 0.008, mp5=3.514 ± 0.025, crit=0.143 ± 0.007 per rating point (14 rating = 1%, 2.001 per %), spell_haste=0.448 ± 0.083

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 18.5 healing_power points (0.76 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.52 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 5.9 healing_power points (0.24 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.5 healing_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (0.4 DPS) | yes | Caretaker's Cape (20428, -0.01 DPS, sim-verified) [rep]; Regent's Cloak (5969, -0.02 DPS) [world]; Seer's Cape (6378, -0.04 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 27.1 healing_power points (1.11 DPS) | yes | Robe of the Moccasin (6465, +0.00 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.18 DPS) [dungeon]; Seer's Robe (2981, -0.45 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.3 healing_power points (0.83 DPS) | yes | Bright Bracers (3647, +0.00 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.55 DPS) [quest]; Seer's Cuffs (3645, -0.63 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 16.8 healing_power points (0.69 DPS) | yes | Magefist Gloves (12977, +0.00 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.13 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.21 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.8 healing_power points (0.77 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.13 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.25 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 35.5 healing_power points (1.45 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -0.73 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.73 DPS) [crafted] |
| feet | Kimbra Boots (6191) | WANTED: Chok'sul [quest] | sim-verified (0.4 DPS) | yes | Pristine Boots (253889, -0.00 DPS, sim-verified) [crafted]; Bluegill Sandals (1560, -0.10 DPS) [world]; Smoldering Boots (3076, -0.14 DPS) [world] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 12.7 healing_power points (0.52 DPS) | yes | Band of Purification (12996, -0.16 DPS) [world_drop]; Deep Fathom Ring (6463, -0.22 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.23 DPS) [rep] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.7 healing_power points (0.48 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Deep Fathom Ring (6463, -0.18 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.19 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 19.5 healing_power points (0.79 DPS) | yes | Staff of Westfall (2042, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.16 DPS) [world]; Staff of the Blessed Seer (2271, -0.20 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) | The Escape [quest] | sim-verified (0.4 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.01 DPS, sim-verified) [quest]; Dwarven Flamestick (5241, -0.04 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Kimbra Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (human, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 56.5. Weights run: 9.3s. Verify run: 4.4s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=1.999 ± 0.011, spirit=2.351 ± 0.009, mp5=4.895 ± 0.011, crit=0.214 ± 0.010 per rating point (14 rating = 1%, 2.992 per %), spell_haste=not significant (-0.085 ± 0.047)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.14 DPS, sim-verified) [world_drop]; Resilient Cap (14401, -0.22 DPS) [world_drop]; Embalmed Shroud (7691, -0.39 DPS) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.4 healing_power points (0.89 DPS) | yes | Necklace of Harmony (5180, +0.00 DPS, sim-verified) [world]; Crystal Starfire Medallion (5003, -0.12 DPS) [world_drop]; Pendant of Myzrael (4614, -0.17 DPS) [dungeon] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 33.7 healing_power points (1.73 DPS) | yes | Mantle of Honor (3560, +0.00 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.31 DPS) [world_drop]; Faerie Mantle (5820, -0.39 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 25.8 healing_power points (1.32 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Repairman's Cape (9605, -0.07 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.14 DPS) [vendor] |
| chest | Beguiler Robes (7728) | Scarlet Monastery: Houndmaster Loksey [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gown (253961, -0.02 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.04 DPS) [dungeon]; Pressed Felt Robe (1997, -0.44 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 24.7 healing_power points (1.27 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.45 DPS) [world_drop]; Spidertank Oilrag (9448, -0.51 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 38.2 healing_power points (1.96 DPS) | yes | Hotshot Pilot's Gloves (9491, +0.00 DPS, sim-verified) [dungeon]; Zodiac Gloves (7106, -0.71 DPS) [quest]; Town Clerk's Mittens (270029, -0.83 DPS) [quest] |
| waist | Resilient Cord (14406) | World drop [world_drop] | 21.4 healing_power points (1.10 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Pristine Sash (253925, -0.12 DPS) [crafted]; Dreamer's Belt (4829, -0.14 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 47.7 healing_power points (2.45 DPS) | yes | Earthen Leggings (253999, -0.08 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.43 DPS) [crafted]; Blighted Leggings (7709, -0.76 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.4 healing_power points (2.23 DPS) | yes | Soggy Boots (274747, -0.07 DPS, sim-verified) [vendor]; Acidic Walkers (9454, -0.92 DPS) [dungeon]; Frothing Slippers (254003, -1.03 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.4 healing_power points (1.56 DPS) | yes | Darkspear Signet (272071, -0.56 DPS) [vendor]; Black Pearl Ring (6332, -0.63 DPS) [world]; Sea Giant's Toe Ring (274746, -0.64 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 22.8 healing_power points (1.17 DPS) | yes | Darkspear Signet (272071, +0.00 DPS, sim-verified) [vendor]; Black Pearl Ring (6332, -0.24 DPS) [world]; Sea Giant's Toe Ring (274746, -0.25 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.05 DPS, sim-verified) [dungeon]; Gnarled Ash Staff (791, -0.51 DPS) [world_drop]; Staff of the Friar (3415, -0.57 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Consecrated Wand (5244, -0.02 DPS, sim-verified) [quest]; Gravestone Scepter (7001, -0.05 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Glowing Thresher Cape; chest: Beguiler Robes; hands: Gloves of Old; waist: Resilient Cord; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 136.5. Weights run: 11.3s. Verify run: 5.5s. 328 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.108, intellect=1.534 ± 0.038, spirit=0.766 ± 0.041, mp5=not significant (0.194 ± 0.066), crit=0.426 ± 0.020 per rating point (14 rating = 1%, 5.959 per %), spell_haste=not significant (-0.288 ± 0.304)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 55.0 healing_power points (4.83 DPS) | yes | Holy Shroud (2721, -1.02 DPS, sim-verified) [world_drop]; Corpseshroud (10574, -1.87 DPS) [dungeon]; Miner's Hat of the Deep (9429, -1.87 DPS) [dungeon] |
| neck | Triune Amulet (7722) (or Necklace of Calisea (1714)) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 16.1 healing_power points (1.42 DPS) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.00 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.07 DPS) [quest] |
| shoulder | Mistscape Mantle (4734) | Uldaman: Ancient Treasure [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Batwing Mantle (6697, +0.00 DPS) [dungeon]; Windchaser Amice (14432, -0.07 DPS) [world_drop]; Earthen Silk Shoulders (254033, -0.20 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 22.8 healing_power points (2.01 DPS) | yes | Darkspear Raider's Cloak (272077, -0.25 DPS) [vendor]; Caretaker's Cape (19532, -0.26 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.1 healing_power points (5.64 DPS) | yes | Death Speaker Robes (6682, -0.32 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -2.67 DPS) [crafted]; Silksand Tunic (14417, -3.14 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mistscape Bracers (4045, -0.13 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.13 DPS) [quest]; Earthen Silk Cuffs (254019, -0.20 DPS, sim-verified) [crafted] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Handwraps (254021, -0.09 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.46 DPS) [crafted]; Truefaith Gloves (7049, -0.86 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 38.9 healing_power points (3.42 DPS) | yes | Deathmage Sash (10771, -0.08 DPS, sim-verified) [dungeon]; Sutarn's Ring (13105, -1.66 DPS) [world_drop]; Windchaser Cinch (14435, -1.73 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.6 healing_power points (3.22 DPS) | yes | Filigreed Pristine Leggings (253937, -0.15 DPS, sim-verified) [crafted]; Stormcloth Pants (10010, -0.89 DPS) [crafted]; Aurora Pants (4044, -1.06 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.8 healing_power points (2.97 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.29 DPS) [dungeon]; Nimbus Boots (6998, -1.48 DPS) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.58 DPS) | yes | Ogremind Ring (1993, -0.44 DPS) [world_drop]; Voodoo Band (1996, -0.44 DPS) [world]; Welken Ring (5011, -0.44 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 16.1 healing_power points (1.42 DPS) | yes | Voodoo Band (1996, -0.07 DPS, sim-verified) [world]; Ogremind Ring (1993, -0.27 DPS) [world_drop]; Welken Ring (5011, -0.27 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Staff of Jordan (873, -3.75 DPS) [world_drop]; Windweaver Staff (7757, -3.96 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 25.1 healing_power points (2.20 DPS) | yes | Orb of Souls (249395, -0.09 DPS, sim-verified) [crafted]; Orb of Lorica (11262, -0.86 DPS) [quest]; Aurora Sphere (7610, -0.99 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 11.5 healing_power points (1.01 DPS) | yes | Goblin Igniter (5253, -0.10 DPS, sim-verified) [quest]; Flash Wand (5248, -0.27 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.28 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Mistscape Mantle; back: Mantle of Lady Falther'ess; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 190.1. Weights run: 11.0s. Verify run: 5.8s. 423 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=1.387 ± 0.020, spirit=not significant (0.158 ± 0.092), mp5=2.113 ± 0.148, crit=0.632 ± 0.019 per rating point (14 rating = 1%, 8.850 per %), spell_haste=not significant (-0.025 ± 0.272)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Cassandra's Grace (13102, -0.15 DPS) [world_drop]; Knight-Lieutenant's Satin Cover (220896, -0.53 DPS, sim-verified) [vendor]; Stormcloth Headband (10032, -0.71 DPS) [crafted] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 21.0 healing_power points (1.98 DPS) | yes | Horizon Choker (13085, +0.00 DPS, sim-verified) [world_drop]; Gemshard Heart (17707, -0.58 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.67 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Satin Pads (220894, -0.36 DPS, sim-verified) [vendor]; Kentic Amice (11624, -0.48 DPS) [dungeon]; Rotgrip Mantle (17732, -1.23 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Raider's Cloak (272076, -0.12 DPS) [vendor]; Caretaker's Cape (19531, -0.15 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 59.3 healing_power points (5.59 DPS) | yes | Knight's Satin Armor (220892, -0.47 DPS, sim-verified) [vendor]; Robes of Insight (940, -2.09 DPS) [world_drop]; Death Speaker Robes (6682, -2.26 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 30.5 healing_power points (2.87 DPS) | yes | Aristocratic Cuffs (12546, +0.00 DPS, sim-verified) [dungeon]; Shizzle's Nozzle Wiper (11917, -1.26 DPS) [quest]; Forgotten Wraps (9433, -1.31 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 52.5 healing_power points (4.95 DPS) | yes | Gilded Gloves (254095, -0.20 DPS, sim-verified) [crafted]; Virtuous Mitts (226950, -0.98 DPS) [vendor]; Sergeant Major's Satin Gloves (220897, -1.51 DPS) [vendor] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 41.5 healing_power points (3.91 DPS) | yes | Gilded Cord (254037, -0.70 DPS) [crafted]; Earthenweave Cord (254077, -1.30 DPS) [crafted]; Ban'thok Sash (11662, -1.34 DPS) [dungeon] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | 52.0 healing_power points (4.90 DPS) | yes | Spellshock Leggings (9484, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.84 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -2.28 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 43.5 healing_power points (4.10 DPS) | yes | Sergeant Major's Satin Boots (220895, -0.49 DPS, sim-verified) [vendor]; Gilded Slippers (254001, -1.24 DPS) [crafted]; Coldstone Slippers (18697, -1.47 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 21.6 healing_power points (2.04 DPS) | yes | Sea Giant's Toe Ring (274746, -0.34 DPS) [vendor]; Mindseye Circle (10634, -0.47 DPS) [dungeon]; Darkspear Signet (272069, -0.64 DPS) [vendor] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.3 healing_power points (1.82 DPS) | yes | Mindseye Circle (10634, -0.25 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.39 DPS, sim-verified) [vendor]; Darkspear Signet (272069, -0.43 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.99 DPS, sim-verified) [quest]; Uther's Strength (11302, -1.29 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -2.87 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS, sim-verified) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.03 DPS) [quest] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -2.48 DPS) [world_drop]; Spellshifter Rod (9527, -3.40 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 22.6 healing_power points (2.13 DPS) | yes | Twisting Essence Jar (249456, +0.00 DPS, sim-verified) [crafted]; Enthralled Sphere (11625, -0.23 DPS) [dungeon]; Skullspell Orb (10708, -0.56 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 8.8 healing_power points (0.83 DPS) | yes | Cairnstone Sliver (9654, +0.00 DPS, sim-verified) [quest]; Captain Rackmore's Tiller (16789, -0.22 DPS) [quest]; Goblin Igniter (5253, -0.26 DPS) [quest] |

**New at 50:** neck: Darkmoon Necklace; shoulder: Nethergeld Shoulders; wrist: Nethergeld Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Knight's Satin Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Darkspear Voodoo Seal

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 313.6. Weights run: 11.2s. Verify run: 5.3s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.193, intellect=1.076 ± 0.053, spirit=-0.682 ± 0.050, mp5=1.039 ± 0.133, crit=0.706 ± 0.035 per rating point (14 rating = 1%, 9.883 per %), spell_haste=not significant (-0.289 ± 0.428)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 61.2 healing_power points (6.64 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.20 DPS) [pvp]; Mooncloth Circlet (14140, -0.96 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Animated Chain Necklace (18723, -0.63 DPS) [dungeon]; Drake Tooth Necklace (21531, -1.23 DPS, sim-verified) [quest]; The Eye of Zuldazar (19593, -1.32 DPS) [quest] |
| shoulder | Burial Shawl (18681) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -1.60 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.8 healing_power points (5.73 DPS) | yes | Cloak of the Cosmos (18389, -0.76 DPS, sim-verified) [dungeon]; Drape of Recovery (272413, -1.91 DPS) [vendor]; Caretaker's Cape (19530, -2.90 DPS) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | 73.4 healing_power points (7.96 DPS) | yes | Truefaith Vestments (14154, -0.09 DPS, sim-verified) [crafted]; Field Marshal's Satin Tunic (231624, -0.25 DPS) [vendor]; Virtuous Robe (226945, -1.01 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.8 healing_power points (3.24 DPS) | yes | Bracers of Mending (23129, -0.13 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -0.23 DPS) [quest]; Nethergeld Cuffs (254061, -0.25 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Gloves (18409, -0.31 DPS) [crafted]; Hands of the Exalted Herald (12554, -0.66 DPS) [dungeon]; Desert Bloom Gloves (20717, -2.02 DPS, sim-verified) [quest] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 46.9 healing_power points (5.09 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Gilded Waistcord (254081, -0.89 DPS) [crafted]; Virtuous Belt (226948, -1.18 DPS) [quest] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 69.7 healing_power points (7.57 DPS) | yes | Marshal's Satin Legguards (231626, -0.89 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -1.48 DPS) [pvp]; Virtuous Skirt (226946, -3.15 DPS, sim-verified) [quest] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 59.1 healing_power points (6.41 DPS) | yes | Marshal's Satin Walkers (231627, -1.16 DPS) [vendor]; Knight-Lieutenant's Satin Walkers (227129, -1.84 DPS) [pvp]; Swarmtender's Footpads (275609, -2.34 DPS, sim-verified) [crafted] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.59 DPS) [quest]; Rosewine Circle (13178, -0.89 DPS) [dungeon]; Naglering (11669, -3.28 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.18 DPS) [quest]; Rosewine Circle (13178, -0.47 DPS) [dungeon]; Naglering (11669, -4.14 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, +0.00 DPS, sim-verified) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -0.89 DPS) [dungeon] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Redemption (22406, -0.22 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -1.63 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 45.2 healing_power points (4.90 DPS) | yes | Grand Marshal's Tome of Restoration (234590, -0.12 DPS) [pvp]; Tome of Divine Right (22319, -1.05 DPS) [dungeon]; Brightly Glowing Stone (18523, -1.47 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 17.0 healing_power points (1.85 DPS) | yes | Sparkling Crystal Wand (20672, -0.25 DPS) [world]; Oblivion's Touch (18761, -0.56 DPS) [dungeon]; Bonecreeper Stylus (13938, -1.34 DPS, sim-verified) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Burial Shawl; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Whipvine Cord; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Fordring's Seal; finger2: Band of Mending; trinket2: Royal Seal of Eldre'Thalas; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 426.0. Weights run: 9.7s. Verify run: 3.9s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.441, intellect=0.811 ± 0.113, spirit=1.147 ± 0.133, mp5=3.132 ± 0.139, crit=not significant (0.278 ± 0.075) per rating point (14 rating = 1%, 3.888 per %), spell_haste=not significant (0.810 ± 1.378)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Virtuous Crown (226947, -2.69 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 53.9 healing_power points (14.15 DPS) | yes | The Eye of Zuldazar (19593, -3.81 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -3.81 DPS) [quest]; Animated Chain Necklace (18723, -11.21 DPS, sim-verified) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Virtuous Mantle (226951, -0.87 DPS) [quest]; Argent Elite Shoulders (227888, -5.44 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 50.1 healing_power points (13.15 DPS) | yes | Caretaker's Cape (19530, -3.92 DPS) [rep]; Cloak of the Cosmos (18389, -3.99 DPS) [dungeon]; Drape of Recovery (272413, -4.46 DPS, sim-verified) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 91.4 healing_power points (23.98 DPS) | yes | Robes of the Exalted (13346, -2.23 DPS, sim-verified) [dungeon]; Field Marshal's Satin Tunic (231624, -4.55 DPS) [vendor]; Virtuous Robe (226945, -5.32 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 40.7 healing_power points (10.68 DPS) | yes | Virtuous Bracers (226949, -0.73 DPS) [quest]; Earthenweave Cuffs (254131, -1.89 DPS) [crafted]; Bracers of Mending (23129, -2.48 DPS, sim-verified) [dungeon] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 60.8 healing_power points (15.95 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS, sim-verified) [dungeon]; Virtuous Mitts (226950, -1.89 DPS) [vendor]; Gilded Gloves (254095, -3.81 DPS) [crafted] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 57.1 healing_power points (14.98 DPS) | yes | Virtuous Belt (226948, -1.71 DPS) [quest]; Penitent's Cinch (272394, -1.72 DPS) [vendor]; Wisdom of the Timbermaw (19047, -2.03 DPS, sim-verified) [crafted] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 77.0 healing_power points (20.21 DPS) | yes | Knight-Captain's Satin Legguards (227125, -2.30 DPS) [pvp]; Marshal's Satin Legguards (231626, -2.80 DPS) [vendor]; Virtuous Skirt (226946, -18.42 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+7.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Boots (15802, -2.49 DPS) [crafted]; Faith Healer's Boots (22247, -2.69 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -7.84 DPS, sim-verified) [vendor] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.53 DPS) [quest]; Fordring's Seal (16058, -0.95 DPS) [quest]; Naglering (11669, -21.83 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.24 DPS) [quest]; Fordring's Seal (16058, -0.66 DPS) [quest]; Naglering (11669, -23.10 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+26.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -1.31 DPS) [vendor]; Mindtap Talisman (18371, -3.29 DPS) [dungeon]; Briarwood Reed (12930, -4.72 DPS) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -2.77 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -2.91 DPS) [dungeon]; Briarwood Reed (12930, -4.34 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -3.09 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -4.01 DPS) [dungeon]; Hand of Edward the Odd (2243, -19.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 29.6 healing_power points (7.77 DPS) | yes | Mana Channeling Wand (18483, -4.49 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.55 DPS) [world]; Bonecreeper Stylus (13938, -5.87 DPS, sim-verified) [dungeon] |

**New at 60:** head: Mooncloth Circlet; neck: Wavefront Necklace; shoulder: Mooncloth Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Desert Bloom Gloves; waist: Whipvine Cord; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Rosewine Circle; finger2: Band of Mending; trinket2: Royal Seal of Eldre'Thalas; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 28.9. Weights run: 9.1s. Verify run: 4.3s. 140 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=1.947 ± 0.013, spirit=1.465 ± 0.008, mp5=3.514 ± 0.025, crit=0.143 ± 0.007 per rating point (14 rating = 1%, 2.001 per %), spell_haste=0.448 ± 0.083

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 18.5 healing_power points (0.76 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.52 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 5.9 healing_power points (0.24 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.5 healing_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 11.9 healing_power points (0.49 DPS) | yes | Sanguine Cape (14376, -0.17 DPS) [world_drop]; Regent's Cloak (5969, -0.19 DPS) [world]; Traveler's Shawl (277289, -0.19 DPS) [quest] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 27.1 healing_power points (1.11 DPS) | yes | Robe of the Moccasin (6465, +0.00 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.18 DPS) [dungeon]; Seer's Robe (2981, -0.45 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.3 healing_power points (0.83 DPS) | yes | Tabitha's Cuffs (251486, +0.00 DPS, sim-verified) [quest]; Featherbead Bracers (15452, -0.43 DPS) [quest]; Crystalline Cuffs (14148, -0.49 DPS) [dungeon] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 21.0 healing_power points (0.86 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.22 DPS) [world_drop]; Bright Gloves (3066, -0.30 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.8 healing_power points (0.77 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.13 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.25 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 35.5 healing_power points (1.45 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -0.73 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.73 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 14.8 healing_power points (0.61 DPS) | yes | Bluegill Sandals (1560, +0.00 DPS, sim-verified) [world]; Smoldering Boots (3076, -0.25 DPS) [world]; Sanguine Sandals (14374, -0.29 DPS) [world_drop] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 12.7 healing_power points (0.52 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Band of Purification (12996, -0.16 DPS) [world_drop]; Deep Fathom Ring (6463, -0.22 DPS) [dungeon] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.7 healing_power points (0.48 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Band of Purification (12996, -0.12 DPS) [world_drop]; Deep Fathom Ring (6463, -0.18 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) | The Wrath of Rath'mael [quest] | sim-verified (0.4 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.01 DPS, sim-verified) [quest]; Advisor's Gnarled Staff (20425, -0.05 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 3.9 healing_power points (0.16 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.08 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 57.3. Weights run: 9.3s. Verify run: 4.4s. 231 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=1.999 ± 0.011, spirit=2.351 ± 0.009, mp5=4.895 ± 0.011, crit=0.214 ± 0.010 per rating point (14 rating = 1%, 2.992 per %), spell_haste=not significant (-0.085 ± 0.047)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.14 DPS, sim-verified) [world_drop]; Resilient Cap (14401, -0.22 DPS) [world_drop]; Embalmed Shroud (7691, -0.39 DPS) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.4 healing_power points (0.89 DPS) | yes | Necklace of Harmony (5180, +0.00 DPS, sim-verified) [world]; Crystal Starfire Medallion (5003, -0.12 DPS) [world_drop]; Pendant of Myzrael (4614, -0.17 DPS) [dungeon] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 33.7 healing_power points (1.73 DPS) | yes | Ghostly Mantle (3324, -0.12 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.31 DPS) [world_drop]; Mantle of Woe (7750, -0.53 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (19529, -0.03 DPS) [rep]; Glowing Thresher Cape (6901, -0.07 DPS, sim-verified) [dungeon]; Amy's Blanket (13005, -0.34 DPS) [world_drop] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 43.4 healing_power points (2.23 DPS) | yes | Beguiler Robes (7728, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.07 DPS) [dungeon]; Pressed Felt Robe (1997, -0.47 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 24.7 healing_power points (1.27 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.45 DPS) [world_drop]; Spidertank Oilrag (9448, -0.51 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 38.2 healing_power points (1.96 DPS) | yes | Hotshot Pilot's Gloves (9491, +0.00 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.64 DPS) [quest]; Tattered Mittens (270030, -0.72 DPS) [quest] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 25.0 healing_power points (1.28 DPS) | yes | Resilient Cord (14406, +0.00 DPS, sim-verified) [world_drop]; Novice Ardent's Sash (253887, -0.31 DPS) [crafted]; Pristine Sash (253925, -0.31 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 47.7 healing_power points (2.45 DPS) | yes | Earthen Leggings (253999, -0.08 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.43 DPS) [crafted]; Blighted Leggings (7709, -0.76 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.4 healing_power points (2.23 DPS) | yes | Soggy Boots (274747, -0.07 DPS, sim-verified) [vendor]; Acidic Walkers (9454, -0.92 DPS) [dungeon]; Frothing Slippers (254003, -1.03 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.4 healing_power points (1.56 DPS) | yes | Darkspear Signet (272071, -0.56 DPS) [vendor]; Black Pearl Ring (6332, -0.63 DPS) [world]; Sea Giant's Toe Ring (274746, -0.64 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 22.8 healing_power points (1.17 DPS) | yes | Darkspear Signet (272071, +0.00 DPS, sim-verified) [vendor]; Black Pearl Ring (6332, -0.24 DPS) [world]; Sea Giant's Toe Ring (274746, -0.25 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.06 DPS, sim-verified) [dungeon]; Gnarled Ash Staff (791, -0.51 DPS) [world_drop]; Staff of the Friar (3415, -0.57 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 8.0 healing_power points (0.41 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Gravestone Scepter (7001, -0.05 DPS) [quest]; Wand of Decay (5252, -0.21 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Pristine Gown; hands: Gloves of Old; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 133.7. Weights run: 11.3s. Verify run: 5.6s. 311 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.108, intellect=1.534 ± 0.038, spirit=0.766 ± 0.041, mp5=not significant (0.194 ± 0.066), crit=0.426 ± 0.020 per rating point (14 rating = 1%, 5.959 per %), spell_haste=not significant (-0.288 ± 0.304)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 55.0 healing_power points (4.83 DPS) | yes | Holy Shroud (2721, -0.98 DPS, sim-verified) [world_drop]; Corpseshroud (10574, -1.87 DPS) [dungeon]; Miner's Hat of the Deep (9429, -1.87 DPS) [dungeon] |
| neck | Triune Amulet (7722) (or Necklace of Calisea (1714)) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 16.1 healing_power points (1.42 DPS) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.00 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.07 DPS) [quest] |
| shoulder | Mistscape Mantle (4734) | Uldaman: Ancient Treasure [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Batwing Mantle (6697, +0.00 DPS) [dungeon]; Windchaser Amice (14432, -0.07 DPS) [world_drop]; Earthen Silk Shoulders (254033, -0.29 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 22.8 healing_power points (2.01 DPS) | yes | Battle Healer's Cloak (19528, -0.19 DPS, sim-verified) [rep]; Darkspear Raider's Cloak (272077, -0.25 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 64.1 healing_power points (5.64 DPS) | yes | Death Speaker Robes (6682, -0.26 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -2.67 DPS) [crafted]; Silksand Tunic (14417, -3.14 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dryad's Wrist Bindings (19597, -0.07 DPS) [pvp]; Mistscape Bracers (4045, -0.13 DPS) [dungeon]; Earthen Silk Cuffs (254019, -0.26 DPS, sim-verified) [crafted] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Handwraps (254021, -0.18 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.46 DPS) [crafted]; Truefaith Gloves (7049, -0.86 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 38.9 healing_power points (3.42 DPS) | yes | Deathmage Sash (10771, +0.00 DPS, sim-verified) [dungeon]; Sutarn's Ring (13105, -1.66 DPS) [world_drop]; Windchaser Cinch (14435, -1.73 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.6 healing_power points (3.22 DPS) | yes | Filigreed Pristine Leggings (253937, -0.12 DPS, sim-verified) [crafted]; Stormcloth Pants (10010, -0.89 DPS) [crafted]; Aurora Pants (4044, -1.06 DPS) [world_drop] |
| feet | Furen's Boots (13100) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Slippers (254001, -0.08 DPS, sim-verified) [crafted]; Thoughtcast Boots (10578, -0.34 DPS) [dungeon]; Boots of the Maharishi (9658, -0.40 DPS) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.58 DPS) | yes | Ogremind Ring (1993, -0.44 DPS) [world_drop]; Voodoo Band (1996, -0.44 DPS) [world]; Welken Ring (5011, -0.44 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 16.1 healing_power points (1.42 DPS) | yes | Voodoo Band (1996, -0.08 DPS, sim-verified) [world]; Ogremind Ring (1993, -0.27 DPS) [world_drop]; Welken Ring (5011, -0.27 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Staff of Jordan (873, -3.75 DPS) [world_drop]; Windweaver Staff (7757, -3.96 DPS) [dungeon] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Souls (249395, -0.07 DPS) [crafted]; Beacon of Hope (9393, -0.27 DPS, sim-verified) [dungeon]; Aurora Sphere (7610, -0.40 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 11.5 healing_power points (1.01 DPS) | yes | Goblin Igniter (5253, -0.11 DPS, sim-verified) [quest]; Flash Wand (5248, -0.27 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.28 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Mistscape Mantle; back: Mantle of Lady Falther'ess; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; feet: Furen's Boots; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Prophetic Cane; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 192.5. Weights run: 11.0s. Verify run: 5.6s. 403 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=1.387 ± 0.020, spirit=not significant (0.158 ± 0.092), mp5=2.113 ± 0.148, crit=0.632 ± 0.019 per rating point (14 rating = 1%, 8.850 per %), spell_haste=not significant (-0.025 ± 0.272)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 47.0 healing_power points (4.43 DPS) | yes | Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor]; Stormcloth Headband (10032, -0.71 DPS) [crafted]; Cassandra's Grace (13102, -1.59 DPS, sim-verified) [world_drop] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 21.0 healing_power points (1.98 DPS) | yes | Horizon Choker (13085, -0.22 DPS, sim-verified) [world_drop]; Gemshard Heart (17707, -0.58 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.67 DPS) [quest] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 38.0 healing_power points (3.59 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -1.23 DPS) [dungeon] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | 22.9 healing_power points (2.16 DPS) | yes | Darkspear Raider's Cloak (272076, -0.26 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -0.45 DPS, sim-verified) [dungeon]; Imperial Red Cloak (8248, -0.67 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 59.3 healing_power points (5.59 DPS) | yes | Stone Guard's Satin Armor (220903, -0.68 DPS, sim-verified) [vendor]; Robes of Insight (940, -2.09 DPS) [world_drop]; Death Speaker Robes (6682, -2.26 DPS) [dungeon] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 30.5 healing_power points (2.87 DPS) | yes | Aristocratic Cuffs (12546, +0.00 DPS, sim-verified) [dungeon]; Shizzle's Nozzle Wiper (11917, -1.26 DPS) [quest]; Forgotten Wraps (9433, -1.31 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 52.5 healing_power points (4.95 DPS) | yes | Gilded Gloves (254095, -0.85 DPS, sim-verified) [crafted]; Virtuous Mitts (226950, -0.98 DPS) [vendor]; Greenleaf Handwraps (19116, -1.44 DPS) [quest] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | 41.5 healing_power points (3.91 DPS) | yes | Gilded Cord (254037, -0.41 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -1.30 DPS) [crafted]; Ban'thok Sash (11662, -1.34 DPS) [dungeon] |
| legs | Stone Guard's Satin Leggings (220902) | Lady Palanseer [vendor] | 52.0 healing_power points (4.90 DPS) | yes | Spellshock Leggings (9484, -0.42 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -1.84 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -2.28 DPS) [dungeon] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 43.5 healing_power points (4.10 DPS) | yes | First Sergeant's Satin Boots (220900, -0.59 DPS, sim-verified) [vendor]; Gilded Slippers (254001, -1.24 DPS) [crafted]; Coldstone Slippers (18697, -1.47 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 21.6 healing_power points (2.04 DPS) | yes | Sea Giant's Toe Ring (274746, -0.34 DPS) [vendor]; Mindseye Circle (10634, -0.47 DPS) [dungeon]; Darkspear Signet (272069, -0.64 DPS) [vendor] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.3 healing_power points (1.82 DPS) | yes | Mindseye Circle (10634, -0.25 DPS) [dungeon]; Darkspear Signet (272069, -0.43 DPS) [vendor]; Sea Giant's Toe Ring (274746, -0.92 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (14.4 DPS) | yes | Ankh of Life (1713, -0.69 DPS, sim-verified) [world_drop]; Uther's Strength (11302, -1.29 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.84 DPS) [quest] |
| trinket2 | Alchemists' Stone (13503) | Alchemy [crafted] | sim-verified (14.4 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.14 DPS, sim-verified) [world_drop] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (14.4 DPS) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Glowing Brightwood Staff (812, -2.48 DPS) [world_drop]; Spellshifter Rod (9527, -3.40 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 22.6 healing_power points (2.13 DPS) | yes | Twisting Essence Jar (249456, -0.17 DPS, sim-verified) [crafted]; Enthralled Sphere (11625, -0.23 DPS) [dungeon]; Prophetic Cane (6803, -0.56 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 8.8 healing_power points (0.83 DPS) | yes | Captain Rackmore's Tiller (16789, -0.22 DPS) [quest]; Goblin Igniter (5253, -0.26 DPS) [quest]; Nature's Breath (19118, -0.45 DPS, sim-verified) [quest] |

**New at 50:** neck: Darkmoon Necklace; shoulder: Nethergeld Shoulders; back: Battle Healer's Cloak; wrist: Nethergeld Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Stone Guard's Satin Leggings; feet: Gilded Sandals; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Darkspear Voodoo Seal; trinket2: Alchemists' Stone; off_hand: Beacon of Hope

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 307.9. Weights run: 11.2s. Verify run: 5.2s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.193, intellect=1.076 ± 0.053, spirit=-0.682 ± 0.050, mp5=1.039 ± 0.133, crit=0.706 ± 0.035 per rating point (14 rating = 1%, 9.883 per %), spell_haste=not significant (-0.289 ± 0.428)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 61.2 healing_power points (6.64 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.20 DPS) [pvp]; Mooncloth Circlet (14140, -1.44 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Animated Chain Necklace (18723, -0.63 DPS) [dungeon]; Drake Tooth Necklace (21531, -1.09 DPS, sim-verified) [quest]; The Eye of Zuldazar (19593, -1.32 DPS) [quest] |
| shoulder | Burial Shawl (18681) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -1.05 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.8 healing_power points (5.73 DPS) | yes | Cloak of the Cosmos (18389, -0.91 DPS, sim-verified) [dungeon]; Drape of Recovery (272413, -1.91 DPS) [vendor]; Battle Healer's Cloak (19526, -2.90 DPS) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | 73.4 healing_power points (7.96 DPS) | yes | Warlord's Satin Tunic (231632, -0.25 DPS) [vendor]; Truefaith Vestments (14154, -0.69 DPS, sim-verified) [crafted]; Virtuous Robe (226945, -1.01 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.8 healing_power points (3.24 DPS) | yes | Virtuous Bracers (226949, -0.23 DPS) [quest]; Nethergeld Cuffs (254061, -0.25 DPS) [crafted]; Bracers of Mending (23129, -0.31 DPS, sim-verified) [dungeon] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 56.4 healing_power points (6.12 DPS) | yes | Raider Handwraps (272097, +0.00 DPS, sim-verified) [vendor]; Mooncloth Gloves (18409, -0.67 DPS) [crafted]; Hands of the Exalted Herald (12554, -1.02 DPS) [dungeon] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 46.9 healing_power points (5.09 DPS) | yes | Wisdom of the Timbermaw (19047, -0.09 DPS) [crafted]; Gilded Waistcord (254081, -0.89 DPS) [crafted]; Virtuous Belt (226948, -1.18 DPS) [quest] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 69.7 healing_power points (7.57 DPS) | yes | General's Satin Legguards (231634, -0.89 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -1.48 DPS) [pvp]; Virtuous Skirt (226946, -3.38 DPS, sim-verified) [quest] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 59.1 healing_power points (6.41 DPS) | yes | General's Satin Walkers (231630, -1.16 DPS) [vendor]; Swarmtender's Footpads (275609, -1.67 DPS, sim-verified) [crafted]; Blood Guard's Satin Walkers (227127, -1.84 DPS) [pvp] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.59 DPS) [quest]; Rosewine Circle (13178, -0.89 DPS) [dungeon]; Naglering (11669, -3.43 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.18 DPS) [quest]; Rosewine Circle (13178, -0.47 DPS) [dungeon]; Naglering (11669, -4.12 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor]; Mindtap Talisman (18371, -1.33 DPS, sim-verified) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS) [vendor]; Mindtap Talisman (18371, -0.42 DPS, sim-verified) [dungeon]; Briarwood Reed (12930, -0.89 DPS) [dungeon] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Redemption (22406, -0.22 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -1.63 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 45.2 healing_power points (4.90 DPS) | yes | High Warlord's Tome of Mending (234564, -0.12 DPS) [pvp]; Tome of Divine Right (22319, -1.05 DPS) [dungeon]; Brightly Glowing Stone (18523, -1.93 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 17.0 healing_power points (1.85 DPS) | yes | Sparkling Crystal Wand (20672, -0.25 DPS) [world]; Oblivion's Touch (18761, -0.56 DPS) [dungeon]; Bonecreeper Stylus (13938, -1.14 DPS, sim-verified) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Burial Shawl; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Bracers of Hope; hands: Desert Bloom Gloves; waist: Whipvine Cord; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Fordring's Seal; finger2: Band of Mending; trinket2: Royal Seal of Eldre'Thalas; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 423.3. Weights run: 9.7s. Verify run: 4.0s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.441, intellect=0.811 ± 0.113, spirit=1.147 ± 0.133, mp5=3.132 ± 0.139, crit=not significant (0.278 ± 0.075) per rating point (14 rating = 1%, 3.888 per %), spell_haste=not significant (0.810 ± 1.378)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Virtuous Crown (226947, -2.62 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 53.9 healing_power points (14.15 DPS) | yes | The Eye of Zuldazar (19593, -3.81 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -3.81 DPS) [quest]; Animated Chain Necklace (18723, -10.08 DPS, sim-verified) [dungeon] |
| shoulder | Mooncloth Shoulders (14139) | Tailoring [crafted] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Virtuous Mantle (226951, -0.87 DPS) [quest]; Argent Elite Shoulders (227888, -7.59 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 50.1 healing_power points (13.15 DPS) | yes | Battle Healer's Cloak (19526, -3.92 DPS) [rep]; Cloak of the Cosmos (18389, -3.99 DPS) [dungeon]; Drape of Recovery (272413, -5.18 DPS, sim-verified) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 91.4 healing_power points (23.98 DPS) | yes | Robes of the Exalted (13346, -0.52 DPS, sim-verified) [dungeon]; Warlord's Satin Tunic (231632, -4.55 DPS) [vendor]; Virtuous Robe (226945, -5.32 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 40.7 healing_power points (10.68 DPS) | yes | Bracers of Mending (23129, -0.67 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -0.73 DPS) [quest]; Earthenweave Cuffs (254131, -1.89 DPS) [crafted] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Virtuous Mitts (226950, -0.98 DPS) [vendor]; Desert Bloom Gloves (20717, -2.80 DPS, sim-verified) [quest]; Gilded Gloves (254095, -2.89 DPS) [crafted] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 57.1 healing_power points (14.98 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Virtuous Belt (226948, -1.71 DPS) [quest]; Penitent's Cinch (272394, -1.72 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 77.0 healing_power points (20.21 DPS) | yes | Legionnaire's Satin Legguards (227123, -2.30 DPS) [pvp]; General's Satin Legguards (231634, -2.80 DPS) [vendor]; Virtuous Skirt (226946, -18.21 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Boots (15802, -2.49 DPS) [crafted]; Faith Healer's Boots (22247, -2.69 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -6.69 DPS, sim-verified) [vendor] |
| finger1 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.53 DPS) [quest]; Fordring's Seal (16058, -0.95 DPS) [quest]; Naglering (11669, -20.51 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.24 DPS) [quest]; Fordring's Seal (16058, -0.66 DPS) [quest]; Naglering (11669, -21.09 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+20.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -1.31 DPS) [vendor]; Mindtap Talisman (18371, -3.29 DPS) [dungeon]; Briarwood Reed (12930, -4.72 DPS) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -2.39 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -2.91 DPS) [dungeon]; Briarwood Reed (12930, -4.34 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -3.09 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -4.01 DPS) [dungeon]; Hand of Edward the Odd (2243, -18.43 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 29.6 healing_power points (7.77 DPS) | yes | Bonecreeper Stylus (13938, -3.36 DPS, sim-verified) [dungeon]; Mana Channeling Wand (18483, -4.49 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.55 DPS) [world] |

**New at 60:** head: Mooncloth Circlet; neck: Wavefront Necklace; shoulder: Mooncloth Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Hands of the Exalted Herald; waist: Whipvine Cord; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Rosewine Circle; finger2: Band of Mending; trinket2: Royal Seal of Eldre'Thalas; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

