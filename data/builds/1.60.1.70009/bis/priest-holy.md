# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 41.1. Weights run: 6.2s. Verify run: 17.7s. 150 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.344 ± 0.006, spirit=0.601 ± 0.003, mp5=1.547 ± 0.011, crit=0.071 ± 0.003 per rating point (14 rating = 1%, 0.990 per %), spell_haste=not significant (-0.006 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.64 DPS) | yes | Shadow Goggles (4373, -0.99 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.42 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 2.4 healing_power points (0.22 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.1 healing_power points (1.10 DPS) | yes | Slime-encrusted Pads (6461, -0.68 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.69 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 10.2 healing_power points (0.93 DPS) | yes | Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.58 DPS) [dungeon]; Sanguine Cape (14376, -0.64 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.5 healing_power points (1.96 DPS) | yes | Robe of the Moccasin (6465, -0.59 DPS) [dungeon]; Bloody Apron (6226, -0.78 DPS) [dungeon]; Corsair's Overshirt (5202, -1.91 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.4 healing_power points (1.04 DPS) | yes | Repurposed Hair Band (281256, -0.68 DPS) [quest]; Mystic's Bracelets (14366, -0.79 DPS) [world_drop]; Bright Bracers (3647, -0.90 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.37 DPS) | yes | Tomb Robber's Gloves (280096, -0.64 DPS) [quest]; Bright Gloves (3066, -0.66 DPS) [world_drop]; Magefist Gloves (12977, -0.91 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.4 healing_power points (1.49 DPS) | yes | Keller's Girdle (2911, -0.51 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.56 DPS, sim-verified) [crafted]; Tarantula Silk Sash (3229, -0.77 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.5 healing_power points (2.60 DPS) | yes | Abomination Skin Leggings (23173, -1.62 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.64 DPS) [crafted]; Darkweave Breeches (12987, -1.88 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.19 DPS) | yes | Walking Boots (4660, -0.70 DPS) [world]; Sanguine Sandals (14374, -0.70 DPS) [world_drop]; Kimbra Boots (6191, -0.82 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.1 healing_power points (0.74 DPS) | yes | Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; Band of Purification (12996, -0.41 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.45 DPS) [rep] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 6.3 healing_power points (0.57 DPS) | yes | Volcanic Rock Ring (12053, -0.21 DPS) [world_drop]; Band of Purification (12996, -0.25 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.29 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 13.4 healing_power points (1.23 DPS) | yes | Channeler's Staff (4437, -0.23 DPS, sim-verified) [world]; Staff of Westfall (2042, -0.28 DPS) [quest]; Lesser Staff of the Spire (1300, -0.49 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 2.7 healing_power points (0.25 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.08 DPS) [quest]; Elven Wand (5604, -0.12 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 64.2. Weights run: 6.2s. Verify run: 19.1s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.413 ± 0.007, spirit=0.841 ± 0.004, mp5=1.628 ± 0.015, crit=0.098 ± 0.005 per rating point (14 rating = 1%, 1.379 per %), spell_haste=0.351 ± 0.047

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Pristine Circlet (253975) | Tailoring [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Cowl (4039, -0.08 DPS) [world_drop]; Resilient Cap (14401, -0.28 DPS) [world_drop]; Holy Shroud (2721, -2.17 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.0 healing_power points (0.82 DPS) | yes | Scorn's Icy Choker (23169, -0.05 DPS) [dungeon]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.18 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 19.7 healing_power points (1.80 DPS) | yes | Death Speaker Mantle (6685, -0.38 DPS) [dungeon]; Nightsky Mantle (4718, -0.39 DPS) [world_drop]; Mantle of Honor (3560, -0.78 DPS, sim-verified) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 16.4 healing_power points (1.49 DPS) | yes | Darkspear Raider's Cloak (272078, -0.23 DPS) [vendor]; Glowing Thresher Cape (6901, -0.24 DPS) [dungeon]; Prelacy Cape (7004, -1.57 DPS, sim-verified) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.87 DPS) [dungeon]; Robes of Arugal (6324, -0.93 DPS) [dungeon]; Death Speaker Robes (6682, -1.78 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.9 healing_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.38 DPS, sim-verified) [world_drop]; Stonecloth Bindings (14416, -0.44 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.75 DPS) | yes | Town Clerk's Mittens (270029, -0.34 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.34 DPS) [dungeon]; Gloves of Old (9395, -1.24 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.7 healing_power points (1.51 DPS) | yes | Resilient Cord (14406, -0.44 DPS) [world_drop]; Dreamer's Belt (4829, -0.46 DPS) [vendor]; Novice Ardent's Sash (253887, -0.87 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.1 healing_power points (3.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.89 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.08 DPS) [crafted]; Necromancer Leggings (2277, -1.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.3 healing_power points (3.02 DPS) | yes | Acidic Walkers (9454, -1.69 DPS) [dungeon]; Frothing Slippers (254003, -1.82 DPS) [crafted]; Nimbus Boots (6998, -3.57 DPS, sim-verified) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Black Widow Band (6199, -0.74 DPS) [world]; The Queen's Jewel (13094, -0.77 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.87 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 15.8 healing_power points (1.43 DPS) | yes | Black Widow Band (6199, -0.51 DPS, sim-verified) [world]; The Queen's Jewel (13094, -0.57 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.66 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 68.0 healing_power points (6.19 DPS) | yes | Wind Spirit Staff (6689, -0.61 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -4.77 DPS) [crafted]; Lorekeeper's Staff (212580, -4.82 DPS) [vendor] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 18.0 healing_power points (1.64 DPS) | yes | Eye of Paleth (2943, -0.48 DPS, sim-verified) [quest]; Orb of Mistmantle (13031, -0.52 DPS) [world_drop]; Alliance Outrunner Healing Rod (285348, -0.55 DPS) [world] |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 5.7 healing_power points (0.51 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Consecrated Wand (5244, -0.10 DPS) [quest]; Flaring Baton (5326, -0.26 DPS) [quest] |

**New at 30:** head: Filigreed Pristine Circlet; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Caretaker's Cape; chest: Pristine Gown; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Starfaller

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 85.7. Weights run: 6.5s. Verify run: 41.8s. 328 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.066, intellect=4.383 ± 0.038, spirit=0.671 ± 0.023, mp5=0.754 ± 0.022, crit=0.207 ± 0.015 per rating point (14 rating = 1%, 2.895 per %), spell_haste=7.521 ± 0.580

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 102.6 healing_power points (4.91 DPS) | yes | Miner's Hat of the Deep (9429, -1.02 DPS) [dungeon]; Thinking Cap (2624, -1.34 DPS) [world]; Corpseshroud (10574, -4.63 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (85.7 DPS) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.22 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.64 DPS, sim-verified) [quest] |
| shoulder | Windchaser Amice (14432) (or Inquisitor's Shawl (19507)) | World drop [world_drop] | 57.0 healing_power points (2.73 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.26 DPS) [dungeon]; Batwing Mantle (6697, -0.26 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 50.9 healing_power points (2.44 DPS) | yes | Blackforge Cape (6424, -0.63 DPS) [dungeon]; Cloak of Rot (4462, -0.76 DPS) [world]; Mantle of Lady Falther'ess (23178, -1.59 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (85.7 DPS) | yes | Red Mageweave Vest (10007, +0.00 DPS) [crafted]; Silksand Wraps (14425, +0.00 DPS) [world_drop]; Silksand Tunic (14417, -5.95 DPS, sim-verified) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 39.4 healing_power points (1.89 DPS) | yes | Aurora Bracers (4043, -0.21 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest]; Mistscape Bracers (4045, -0.21 DPS, sim-verified) [dungeon] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (85.7 DPS) | yes | Town Clerk's Mittens (270029, -0.74 DPS) [quest]; Red Mageweave Gloves (10018, -0.95 DPS) [crafted]; Gilded Handwraps (254021, -1.10 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (85.7 DPS) | yes | Razzeric's Customized Seatbelt (6726, -0.41 DPS) [quest]; Teacher's Sash (10747, -0.41 DPS) [quest]; Deathmage Sash (10771, -2.79 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (85.7 DPS) | yes | Aurora Pants (4044, +0.00 DPS) [world_drop]; Crimson Silk Pantaloons (7062, +0.00 DPS) [crafted]; Pristine Leggings (253987, -1.78 DPS, sim-verified) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | sim-verified (85.7 DPS) | yes | Gilded Slippers (254001, +0.00 DPS, sim-verified) [crafted]; Kodo Rustler Boots (15697, -0.14 DPS) [quest]; Acidic Walkers (9454, -0.18 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 35.4 healing_power points (1.70 DPS) | yes | Ogremind Ring (1993, -0.13 DPS) [world_drop]; Mindbender Loop (5009, -0.16 DPS) [world_drop]; Black Widow Band (6199, -0.22 DPS) [world] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 32.7 healing_power points (1.57 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.03 DPS) [world_drop]; Black Widow Band (6199, -0.10 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (85.7 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (85.7 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -0.11 DPS) [dungeon]; Staff of Jordan (873, -0.59 DPS) [world_drop] |
| off_hand | Aurora Sphere (7610) | Uldaman: Ancient Treasure [dungeon] | sim-verified (85.7 DPS) | yes | Skull of Impending Doom (4984, -0.13 DPS) [quest]; Nightsky Orb (15929, -0.24 DPS) [world_drop]; Orb of Lorica (11262, -1.97 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 28.3 healing_power points (1.36 DPS) | yes | Goblin Igniter (5253, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.42 DPS) [quest]; Starfaller (13063, -0.52 DPS) [world_drop] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Windchaser Amice; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; feet: Furen's Boots; finger1: Snake Hoop; finger2: Voodoo Band; trinket2: Ankh of Life; off_hand: Aurora Sphere; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 108.4. Weights run: 6.5s. Verify run: 22.7s. 423 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=4.816 ± 0.060, spirit=0.208 ± 0.028, mp5=1.346 ± 0.060, crit=0.289 ± 0.021 per rating point (14 rating = 1%, 4.052 per %), spell_haste=-25.020 ± 1.049

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 130.6 healing_power points (6.08 DPS) | yes | Soulcatcher Halo (10630, +0.00 DPS, sim-verified) [dungeon]; Papal Fez (9431, -1.16 DPS) [dungeon]; Bad Mojo Mask (9470, -1.37 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 68.3 healing_power points (3.18 DPS) | yes | Gemshard Heart (17707, -0.88 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.94 DPS) [quest]; Darkspear Warding Pendant (272073, -1.16 DPS) [vendor] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Satin Pads (220894, -0.11 DPS) [vendor]; Imperial Red Mantle (8250, -0.26 DPS) [world_drop]; Rotgrip Mantle (17732, -1.19 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 68.5 healing_power points (3.19 DPS) | yes | Imperial Red Cloak (8248, -0.27 DPS, sim-verified) [world_drop]; Mantle of Lady Falther'ess (23178, -0.75 DPS) [dungeon]; Keeper's Cloak (14665, -0.95 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 123.5 healing_power points (5.75 DPS) | yes | Acumen Robes (17775, -1.27 DPS) [quest]; Knight's Satin Armor (220892, -1.52 DPS) [vendor]; Hibernal Robe (8113, -1.63 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 73.5 healing_power points (3.42 DPS) | yes | Shizzle's Nozzle Wiper (11917, -0.40 DPS, sim-verified) [quest]; Forgotten Wraps (9433, -0.73 DPS) [world_drop]; Nethergeld Cuffs (254061, -0.87 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 127.9 healing_power points (5.96 DPS) | yes | Virtuous Mitts (226950, -1.48 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -2.29 DPS) [crafted]; Stormcloth Gloves (10011, -2.76 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 91.5 healing_power points (4.26 DPS) | yes | Serenity Belt (13144, -0.16 DPS, sim-verified) [world_drop]; Gilded Waistcord (254081, -0.89 DPS) [crafted]; Imperial Red Sash (8253, -0.90 DPS) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Venomshroud Leggings (14444, -0.73 DPS) [world_drop]; Imperial Red Pants (8251, -0.85 DPS) [world_drop]; Knight's Satin Leggings (220893, -3.88 DPS, sim-verified) [vendor] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Coldstone Slippers (18697, -0.07 DPS) [dungeon]; Highborne Footpads (14447, -0.77 DPS) [world_drop]; Sergeant Major's Satin Boots (220895, -2.80 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 73.3 healing_power points (3.41 DPS) | yes | Cyclopean Band (11824, -1.38 DPS) [dungeon]; Woodseed Hoop (17768, -1.39 DPS) [quest]; Snake Hoop (6750, -1.77 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 57.8 healing_power points (2.69 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Woodseed Hoop (17768, -0.67 DPS) [quest]; Snake Hoop (6750, -1.05 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.10 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.84 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.86 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.02 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.04 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -1.43 DPS) [quest]; Radiant Staff (249453, -2.33 DPS) [crafted]; Barman Shanker (12791, -6.25 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 29.5 healing_power points (1.37 DPS) | yes | Cairnstone Sliver (9654, -0.15 DPS, sim-verified) [quest]; Flash Wand (5248, -0.45 DPS) [quest]; Goblin Igniter (5253, -0.45 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Kentic Amice; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 343.6. Weights run: 5.9s. Verify run: 31.9s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.173, intellect=1.298 ± 0.050, spirit=0.941 ± 0.058, mp5=0.645 ± 0.064, crit=0.186 ± 0.025 per rating point (14 rating = 1%, 2.604 per %), spell_haste=-10.031 ± 0.944

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 72.7 healing_power points (12.46 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Devout Crown (16693, -1.28 DPS) [dungeon] |
| neck | Drake Tooth Necklace (21531) | The Nightmare Manifests [quest] | 41.6 healing_power points (7.13 DPS) | yes | The Eye of Zuldazar (19593, -0.16 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -0.16 DPS) [quest]; Animated Chain Necklace (18723, -0.50 DPS) [dungeon] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (343.5 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.0 healing_power points (9.42 DPS) | yes | Cloak of the Cosmos (18389, -2.52 DPS) [dungeon]; Drape of Recovery (272413, -3.67 DPS) [vendor]; Caretaker's Cape (19530, -3.68 DPS) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | sim-verified (343.5 DPS) | yes | Truefaith Vestments (14154, +0.00 DPS) [crafted]; Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Virtuous Robe (226945, -1.78 DPS) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (343.5 DPS) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon]; Marshal's Satin Bracers (17606, -1.02 DPS) [pvp] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-verified (343.5 DPS) | yes | Desert Bloom Gloves (20717, +0.00 DPS) [quest]; Raider Handwraps (272097, -0.36 DPS) [vendor]; Virtuous Mitts (226950, -1.02 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 57.6 healing_power points (9.88 DPS) | yes | Virtuous Belt (226948, -1.40 DPS) [quest]; Whipvine Cord (18327, -1.90 DPS) [dungeon]; Penitent's Cinch (272394, -2.07 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 71.8 healing_power points (12.31 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Virtuous Skirt (226946, -0.54 DPS) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (343.5 DPS) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -0.44 DPS) [crafted]; Faith Healer's Boots (22247, -0.70 DPS) [dungeon] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (343.5 DPS) | yes | Blessed Band of Light (272407, -0.85 DPS) [vendor]; Emerald Flame Ring (18395, -0.92 DPS) [dungeon]; Band of Piety (22681, -1.02 DPS) [quest] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (343.5 DPS) | yes | Blessed Band of Light (272407, -1.23 DPS) [vendor]; Emerald Flame Ring (18395, -1.31 DPS) [dungeon]; Band of Piety (22681, -1.40 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (343.5 DPS) | yes | Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -1.13 DPS) [dungeon]; Second Wind (11819, -2.33 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -1.59 DPS) [dungeon]; Staff of Metanoia (22394, -1.82 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -2.55 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.3 healing_power points (4.69 DPS) | yes | Sparkling Crystal Wand (20672, -1.83 DPS) [world]; Bonecreeper Stylus (13938, -1.91 DPS) [dungeon]; Oblivion's Touch (18761, -2.24 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Drake Tooth Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Virtuous Bracers; hands: Hands of the Exalted Herald; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Fordring's Seal; finger2: Band of Mending; trinket2: Royal Seal of Eldre'Thalas; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 713.9. Weights run: 4.0s. Verify run: 14.3s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.322, intellect=1.599 ± 0.050, spirit=1.187 ± 0.049, mp5=2.324 ± 0.056, crit=0.477 ± 0.039 per rating point (14 rating = 1%, 6.680 per %), spell_haste=not significant (-0.271 ± 1.106)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 85.9 healing_power points (22.45 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.12 DPS) [pvp]; Devout Crown (16693, -2.53 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.2 healing_power points (13.64 DPS) | yes | The Eye of Zuldazar (19593, -1.56 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.56 DPS) [quest]; Drake Tooth Necklace (21531, -1.83 DPS) [quest] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (714.0 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 58.0 healing_power points (15.16 DPS) | yes | Cloak of the Cosmos (18389, -3.77 DPS) [dungeon]; Drape of Recovery (272413, -4.63 DPS) [vendor]; Caretaker's Cape (19530, -5.88 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.0 healing_power points (24.05 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.77 DPS) [dungeon]; Virtuous Robe (226945, -1.29 DPS) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (714.0 DPS) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon]; Marshal's Satin Bracers (17606, -0.84 DPS) [pvp] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | sim-verified (714.0 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Virtuous Mitts (226950, -0.52 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 64.1 healing_power points (16.75 DPS) | yes | Virtuous Belt (226948, -1.15 DPS) [quest]; Whipvine Cord (18327, -1.24 DPS) [dungeon]; Marshal's Satin Sash (17609, -2.48 DPS) [pvp] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 87.9 healing_power points (22.99 DPS) | yes | Marshal's Satin Legguards (231626, -1.22 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -1.91 DPS) [pvp]; Virtuous Skirt (226946, -2.61 DPS) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (714.0 DPS) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -1.92 DPS) [crafted]; Faith Healer's Boots (22247, -2.55 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.53 DPS) [quest]; Emerald Flame Ring (18395, -1.47 DPS) [dungeon]; Rosewine Circle (13178, -2.27 DPS) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (714.0 DPS) | yes | Band of Piety (22681, +0.00 DPS) [quest]; Emerald Flame Ring (18395, -0.34 DPS) [dungeon]; Rosewine Circle (13178, -1.14 DPS) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (714.0 DPS) | yes | Darkspear Voodoo Seal (272061, -1.87 DPS) [vendor]; Briarwood Reed (12930, -3.40 DPS) [dungeon]; Mindtap Talisman (18371, -4.30 DPS) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (714.0 DPS) | yes | Darkspear Voodoo Seal (272061, -1.95 DPS) [vendor]; Briarwood Reed (12930, -3.48 DPS) [dungeon]; Mindtap Talisman (18371, -4.38 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -1.52 DPS) [dungeon]; Death Speaker Scepter (2816, -3.20 DPS) [dungeon]; Staff of Hale Magefire (13000, -3.70 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.1 healing_power points (7.86 DPS) | yes | Sparkling Crystal Wand (20672, -2.79 DPS) [world]; Oblivion's Touch (18761, -3.26 DPS) [dungeon]; Bonecreeper Stylus (13938, -3.31 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Serenity Field; trinket2: Royal Seal of Eldre'Thalas; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 39.1. Weights run: 6.2s. Verify run: 17.8s. 140 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.344 ± 0.006, spirit=0.601 ± 0.003, mp5=1.547 ± 0.011, crit=0.071 ± 0.003 per rating point (14 rating = 1%, 0.990 per %), spell_haste=not significant (-0.006 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.64 DPS) | yes | Shadow Goggles (4373, -0.83 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.42 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 2.4 healing_power points (0.22 DPS) | yes | Roadwatcher's Confidence (281265, -0.08 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.1 healing_power points (1.10 DPS) | yes | Slime-encrusted Pads (6461, -0.68 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.88 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 10.2 healing_power points (0.93 DPS) | yes | Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.58 DPS) [dungeon]; Sanguine Cape (14376, -0.63 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.5 healing_power points (1.96 DPS) | yes | Robe of the Moccasin (6465, -0.59 DPS) [dungeon]; Bloody Apron (6226, -0.78 DPS) [dungeon]; Corsair's Overshirt (5202, -1.59 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.4 healing_power points (1.04 DPS) | yes | Featherbead Bracers (15452, -0.42 DPS) [quest]; Bright Bracers (3647, -0.55 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.87 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.37 DPS) | yes | Blight Gloves (279877, -0.16 DPS, sim-verified) [quest]; Magefist Gloves (12977, -0.54 DPS) [world_drop]; Bright Gloves (3066, -0.66 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.4 healing_power points (1.49 DPS) | yes | Novice Ardent's Sash (253887, -0.47 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.51 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.77 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.5 healing_power points (2.60 DPS) | yes | Abomination Skin Leggings (23173, -1.62 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.64 DPS) [crafted]; Darkweave Breeches (12987, -1.81 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.19 DPS) | yes | Spidersilk Boots (4320, -0.70 DPS) [crafted]; Walking Boots (4660, -0.70 DPS) [world]; Sanguine Sandals (14374, -1.02 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.1 healing_power points (0.74 DPS) | yes | Black Pearl Ring (6332, -0.16 DPS) [world]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; Band of Purification (12996, -0.41 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.7 healing_power points (0.61 DPS) | yes | Black Pearl Ring (6332, +0.00 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; Band of Purification (12996, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 13.4 healing_power points (1.23 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.18 DPS) [quest]; Channeler's Staff (4437, -0.25 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.7 healing_power points (0.25 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.12 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (troll, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 61.5. Weights run: 6.2s. Verify run: 19.4s. 231 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.413 ± 0.007, spirit=0.841 ± 0.004, mp5=1.628 ± 0.015, crit=0.098 ± 0.005 per rating point (14 rating = 1%, 1.379 per %), spell_haste=0.351 ± 0.047

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Pristine Circlet (253975) | Tailoring [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Cowl (4039, -0.08 DPS) [world_drop]; Resilient Cap (14401, -0.28 DPS) [world_drop]; Holy Shroud (2721, -1.92 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.0 healing_power points (0.82 DPS) | yes | Scorn's Icy Choker (23169, +0.00 DPS, sim-verified) [dungeon]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.18 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 19.7 healing_power points (1.80 DPS) | yes | Death Speaker Mantle (6685, -0.38 DPS) [dungeon]; Nightsky Mantle (4718, -0.39 DPS) [world_drop]; Ghostly Mantle (3324, -1.55 DPS, sim-verified) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 16.4 healing_power points (1.49 DPS) | yes | Glowing Thresher Cape (6901, -0.24 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.43 DPS, sim-verified) [vendor]; Cloak of Rot (4462, -0.46 DPS) [world] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.87 DPS) [dungeon]; Robes of Arugal (6324, -0.93 DPS) [dungeon]; Death Speaker Robes (6682, -1.71 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.9 healing_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.32 DPS) [quest]; Glowing Magical Bracelets (13106, -0.84 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.75 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.34 DPS) [dungeon]; Pristine Gloves (253913, -0.36 DPS) [crafted]; Gloves of Old (9395, -1.27 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.7 healing_power points (1.51 DPS) | yes | Novice Ardent's Sash (253887, -0.35 DPS) [crafted]; Resilient Cord (14406, -0.44 DPS) [world_drop]; Lilac Sash (6780, -0.54 DPS, sim-verified) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.1 healing_power points (3.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.62 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.08 DPS) [crafted]; Necromancer Leggings (2277, -1.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.3 healing_power points (3.02 DPS) | yes | Frothing Slippers (254003, -1.82 DPS) [crafted]; Fiery Slippers (254005, -1.82 DPS) [crafted]; Acidic Walkers (9454, -2.29 DPS, sim-verified) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Black Widow Band (6199, -0.74 DPS) [world]; The Queen's Jewel (13094, -0.77 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.87 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 15.8 healing_power points (1.43 DPS) | yes | The Queen's Jewel (13094, -0.57 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.66 DPS) [dungeon]; Black Widow Band (6199, -0.83 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 68.0 healing_power points (6.19 DPS) | yes | Wind Spirit Staff (6689, -0.64 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (19569, -4.69 DPS) [pvp]; Glimmering Staff (249392, -4.77 DPS) [crafted] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 18.0 healing_power points (1.64 DPS) | yes | Alliance Outrunner Healing Rod (285348, -0.55 DPS) [world]; Defective Samophlange (274743, -0.62 DPS) [vendor]; Orb of Mistmantle (13031, -1.62 DPS, sim-verified) [world_drop] |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 5.7 healing_power points (0.51 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Wand of Decay (5252, -0.26 DPS) [quest]; Flaring Baton (5326, -0.26 DPS) [quest] |

**New at 30:** head: Filigreed Pristine Circlet; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Battle Healer's Cloak; chest: Pristine Gown; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Starfaller

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 83.7. Weights run: 6.5s. Verify run: 43.7s. 311 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.066, intellect=4.383 ± 0.038, spirit=0.671 ± 0.023, mp5=0.754 ± 0.022, crit=0.207 ± 0.015 per rating point (14 rating = 1%, 2.895 per %), spell_haste=7.521 ± 0.580

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 102.6 healing_power points (4.91 DPS) | yes | Miner's Hat of the Deep (9429, -1.02 DPS) [dungeon]; Thinking Cap (2624, -1.34 DPS) [world]; Corpseshroud (10574, -1.65 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.22 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -1.30 DPS, sim-verified) [quest] |
| shoulder | Windchaser Amice (14432) (or Inquisitor's Shawl (19507)) | World drop [world_drop] | 57.0 healing_power points (2.73 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.26 DPS) [dungeon]; Batwing Mantle (6697, -0.26 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 50.9 healing_power points (2.44 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.12 DPS, sim-verified) [dungeon]; Blackforge Cape (6424, -0.63 DPS) [dungeon]; Cloak of Rot (4462, -0.76 DPS) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Vest (10007, +0.00 DPS) [crafted]; Silksand Wraps (14425, +0.00 DPS) [world_drop]; Silksand Tunic (14417, -6.05 DPS, sim-verified) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 39.4 healing_power points (1.89 DPS) | yes | Radiant Silver Bracers (4545, -0.21 DPS) [quest]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest]; Mistscape Bracers (4045, -0.46 DPS, sim-verified) [dungeon] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 63.6 healing_power points (3.05 DPS) | yes | Red Mageweave Gloves (10018, -0.95 DPS) [crafted]; Mistscape Gloves (6428, -1.16 DPS) [dungeon]; Gilded Handwraps (254021, -5.84 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razzeric's Customized Seatbelt (6726, -0.41 DPS) [quest]; Mistscape Sash (4736, -0.62 DPS) [dungeon]; Deathmage Sash (10771, -3.14 DPS, sim-verified) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (83.7 DPS) | yes | Crimson Silk Pantaloons (7062, +0.00 DPS) [crafted]; Aurora Pants (4044, -0.05 DPS) [world_drop]; Stormcloth Pants (10010, -1.89 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 53.4 healing_power points (2.56 DPS) | yes | Furen's Boots (13100, -0.57 DPS) [world_drop]; Kodo Rustler Boots (15697, -0.72 DPS) [quest]; Boots of the Maharishi (9658, -2.40 DPS, sim-verified) [quest] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 35.4 healing_power points (1.70 DPS) | yes | Ogremind Ring (1993, -0.13 DPS) [world_drop]; Mindbender Loop (5009, -0.16 DPS) [world_drop]; Black Widow Band (6199, -0.22 DPS) [world] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 32.7 healing_power points (1.57 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.03 DPS) [world_drop]; Black Widow Band (6199, -0.10 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -0.11 DPS) [dungeon]; Staff of Jordan (873, -0.59 DPS) [world_drop] |
| off_hand | Aurora Sphere (7610) | Uldaman: Ancient Treasure [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull of Impending Doom (4984, -0.13 DPS) [quest]; Witch's Finger (16887, -0.13 DPS) [quest]; Prophetic Cane (6803, -0.91 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 28.3 healing_power points (1.36 DPS) | yes | Goblin Igniter (5253, -0.25 DPS, sim-verified) [quest]; Flash Wand (5248, -0.42 DPS) [quest]; Starfaller (13063, -0.52 DPS) [world_drop] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Windchaser Amice; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; finger1: Snake Hoop; finger2: Voodoo Band; trinket2: Ankh of Life; off_hand: Aurora Sphere; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (troll, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 101.0. Weights run: 6.5s. Verify run: 25.0s. 403 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=4.816 ± 0.060, spirit=0.208 ± 0.028, mp5=1.346 ± 0.060, crit=0.289 ± 0.021 per rating point (14 rating = 1%, 4.052 per %), spell_haste=-25.020 ± 1.049

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 130.6 healing_power points (6.08 DPS) | yes | Soulcatcher Halo (10630, -0.38 DPS) [dungeon]; Papal Fez (9431, -1.16 DPS) [dungeon]; Bad Mojo Mask (9470, -1.37 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 68.3 healing_power points (3.18 DPS) | yes | Gemshard Heart (17707, -0.88 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.94 DPS) [quest]; Darkspear Warding Pendant (272073, -1.16 DPS) [vendor] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (101.0 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS) [dungeon]; Blood Guard's Satin Pads (220901, -0.11 DPS) [vendor]; Imperial Red Mantle (8250, -0.26 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 68.5 healing_power points (3.19 DPS) | yes | Imperial Red Cloak (8248, -0.68 DPS) [world_drop]; Mantle of Lady Falther'ess (23178, -0.75 DPS) [dungeon]; Keeper's Cloak (14665, -0.95 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 123.5 healing_power points (5.75 DPS) | yes | Hibernal Robe (8113, -1.27 DPS) [world_drop]; Acumen Robes (17775, -1.27 DPS) [quest]; Stone Guard's Satin Armor (220903, -1.52 DPS) [vendor] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 73.5 healing_power points (3.42 DPS) | yes | Shizzle's Nozzle Wiper (11917, -0.70 DPS) [quest]; Forgotten Wraps (9433, -0.73 DPS) [world_drop]; Nethergeld Cuffs (254061, -0.87 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 127.9 healing_power points (5.96 DPS) | yes | Virtuous Mitts (226950, -2.20 DPS) [vendor]; Gilded Gloves (254095, -2.29 DPS) [crafted]; Greenleaf Handwraps (19116, -2.47 DPS) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 91.5 healing_power points (4.26 DPS) | yes | Serenity Belt (13144, -0.45 DPS) [world_drop]; Gilded Waistcord (254081, -0.89 DPS) [crafted]; Imperial Red Sash (8253, -0.90 DPS) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stone Guard's Satin Leggings (220902, +0.00 DPS) [vendor]; Venomshroud Leggings (14444, -0.73 DPS) [world_drop]; Imperial Red Pants (8251, -0.85 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | First Sergeant's Satin Boots (220900, +0.00 DPS) [vendor]; Coldstone Slippers (18697, -0.07 DPS) [dungeon]; Highborne Footpads (14447, -0.77 DPS) [world_drop] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 73.3 healing_power points (3.41 DPS) | yes | Cyclopean Band (11824, -1.38 DPS) [dungeon]; Woodseed Hoop (17768, -1.39 DPS) [quest]; Snake Hoop (6750, -1.77 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 57.8 healing_power points (2.69 DPS) | yes | Cyclopean Band (11824, -0.66 DPS) [dungeon]; Woodseed Hoop (17768, -0.67 DPS) [quest]; Snake Hoop (6750, -1.05 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.10 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.84 DPS) [quest]; Alchemists' Stone (13503, -0.90 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.02 DPS) [quest]; Alchemists' Stone (13503, -0.08 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -1.43 DPS) [quest]; Radiant Staff (249453, -2.33 DPS) [crafted]; Strength of the Treant (9683, -2.55 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 29.5 healing_power points (1.37 DPS) | yes | Nature's Breath (19118, -0.03 DPS) [quest]; Flash Wand (5248, -0.45 DPS) [quest]; Goblin Igniter (5253, -0.45 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Kentic Amice; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 335.3. Weights run: 5.9s. Verify run: 32.7s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.173, intellect=1.298 ± 0.050, spirit=0.941 ± 0.058, mp5=0.645 ± 0.064, crit=0.186 ± 0.025 per rating point (14 rating = 1%, 2.604 per %), spell_haste=-10.031 ± 0.944

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 72.7 healing_power points (12.46 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Devout Crown (16693, -1.28 DPS) [dungeon] |
| neck | Drake Tooth Necklace (21531) | The Nightmare Manifests [quest] | 41.6 healing_power points (7.13 DPS) | yes | The Eye of Zuldazar (19593, -0.16 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -0.16 DPS) [quest]; Animated Chain Necklace (18723, -0.50 DPS) [dungeon] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (335.2 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.0 healing_power points (9.42 DPS) | yes | Cloak of the Cosmos (18389, -2.52 DPS) [dungeon]; Drape of Recovery (272413, -3.67 DPS) [vendor]; Battle Healer's Cloak (19526, -3.68 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 88.1 healing_power points (15.09 DPS) | yes | Warlord's Satin Tunic (231632, -0.39 DPS) [vendor]; Robes of the Exalted (13346, -0.55 DPS) [dungeon]; Virtuous Robe (226945, -2.33 DPS) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (335.2 DPS) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon]; General's Satin Bracers (17619, -1.02 DPS) [pvp] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-verified (335.2 DPS) | yes | Desert Bloom Gloves (20717, +0.00 DPS) [quest]; Raider Handwraps (272097, -0.36 DPS) [vendor]; Virtuous Mitts (226950, -1.02 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 57.6 healing_power points (9.88 DPS) | yes | Virtuous Belt (226948, -1.40 DPS) [quest]; Whipvine Cord (18327, -1.90 DPS) [dungeon]; Penitent's Cinch (272394, -2.07 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 71.8 healing_power points (12.31 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Virtuous Skirt (226946, -0.54 DPS) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -0.44 DPS) [crafted]; Faith Healer's Boots (22247, -0.70 DPS) [dungeon] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (335.2 DPS) | yes | Blessed Band of Light (272407, -0.85 DPS) [vendor]; Emerald Flame Ring (18395, -0.92 DPS) [dungeon]; Band of Piety (22681, -1.02 DPS) [quest] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (335.2 DPS) | yes | Blessed Band of Light (272407, -1.23 DPS) [vendor]; Emerald Flame Ring (18395, -1.31 DPS) [dungeon]; Band of Piety (22681, -1.40 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (335.2 DPS) | yes | Royal Seal of Eldre'Thalas (18469, -1.10 DPS) [quest]; Briarwood Reed (12930, -2.23 DPS) [dungeon]; Second Wind (11819, -3.43 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -1.59 DPS) [dungeon]; Staff of Metanoia (22394, -1.82 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -2.55 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.3 healing_power points (4.69 DPS) | yes | Sparkling Crystal Wand (20672, -1.83 DPS) [world]; Bonecreeper Stylus (13938, -1.91 DPS) [dungeon]; Oblivion's Touch (18761, -2.24 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Drake Tooth Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Hands of the Exalted Herald; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Fordring's Seal; finger2: Band of Mending; trinket2: Serenity Field; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 693.4. Weights run: 4.0s. Verify run: 14.8s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.322, intellect=1.599 ± 0.050, spirit=1.187 ± 0.049, mp5=2.324 ± 0.056, crit=0.477 ± 0.039 per rating point (14 rating = 1%, 6.680 per %), spell_haste=not significant (-0.271 ± 1.106)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 85.9 healing_power points (22.45 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.12 DPS) [pvp]; Devout Crown (16693, -2.53 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.2 healing_power points (13.64 DPS) | yes | The Eye of Zuldazar (19593, -1.56 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.56 DPS) [quest]; Drake Tooth Necklace (21531, -1.83 DPS) [quest] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (693.5 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 58.0 healing_power points (15.16 DPS) | yes | Cloak of the Cosmos (18389, -3.77 DPS) [dungeon]; Drape of Recovery (272413, -4.63 DPS) [vendor]; Battle Healer's Cloak (19526, -5.88 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 92.0 healing_power points (24.05 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.77 DPS) [dungeon]; Virtuous Robe (226945, -1.29 DPS) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (693.5 DPS) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon]; General's Satin Bracers (17619, -0.84 DPS) [pvp] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | sim-verified (693.5 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Virtuous Mitts (226950, -0.52 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 64.1 healing_power points (16.75 DPS) | yes | Virtuous Belt (226948, -1.15 DPS) [quest]; Whipvine Cord (18327, -1.24 DPS) [dungeon]; General's Satin Cinch (17621, -2.48 DPS) [pvp] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 87.9 healing_power points (22.99 DPS) | yes | General's Satin Legguards (231634, -1.22 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -1.91 DPS) [pvp]; Virtuous Skirt (226946, -2.61 DPS) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (693.5 DPS) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -1.92 DPS) [crafted]; Faith Healer's Boots (22247, -2.55 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.53 DPS) [quest]; Emerald Flame Ring (18395, -1.47 DPS) [dungeon]; Rosewine Circle (13178, -2.27 DPS) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (693.5 DPS) | yes | Band of Piety (22681, +0.00 DPS) [quest]; Emerald Flame Ring (18395, -0.34 DPS) [dungeon]; Rosewine Circle (13178, -1.14 DPS) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, -1.95 DPS) [vendor]; Briarwood Reed (12930, -3.48 DPS) [dungeon]; Mindtap Talisman (18371, -4.38 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (693.5 DPS) | yes | Darkspear Voodoo Seal (272061, -1.87 DPS) [vendor]; Briarwood Reed (12930, -3.40 DPS) [dungeon]; Mindtap Talisman (18371, -4.30 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -1.52 DPS) [dungeon]; Death Speaker Scepter (2816, -3.20 DPS) [dungeon]; Staff of Hale Magefire (13000, -3.70 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.1 healing_power points (7.86 DPS) | yes | Sparkling Crystal Wand (20672, -2.79 DPS) [world]; Oblivion's Touch (18761, -3.26 DPS) [dungeon]; Bonecreeper Stylus (13938, -3.31 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Royal Seal of Eldre'Thalas; trinket2: Serenity Field; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

