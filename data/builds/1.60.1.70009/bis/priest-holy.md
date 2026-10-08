# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 39.0. Weights run: 7.5s. Verify run: 3.8s. 150 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.330 ± 0.006, spirit=0.600 ± 0.003, mp5=1.543 ± 0.010, crit=0.061 ± 0.003 per rating point (14 rating = 1%, 0.856 per %), spell_haste=not significant (-0.002 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, -0.04 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.71 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 2.4 healing_power points (0.22 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.08 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.67 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (20428, -0.02 DPS, sim-verified) [rep]; Pearl-clasped Cloak (5542, -0.12 DPS) [crafted]; Seer's Cape (6378, -0.13 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.4 healing_power points (1.95 DPS) | yes | Corsair's Overshirt (5202, -0.12 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.59 DPS) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.3 healing_power points (1.02 DPS) | yes | Bright Bracers (3647, +0.00 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.67 DPS) [quest]; Mystic's Bracelets (14366, -0.78 DPS) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Magefist Gloves (12977, +0.00 DPS, sim-verified) [world_drop]; Tomb Robber's Gloves (280096, -0.64 DPS) [quest]; Bright Gloves (3066, -0.66 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.52 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.77 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.58 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Abomination Skin Leggings (23173, -1.61 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Kimbra Boots (6191, +0.00 DPS, sim-verified) [quest]; Walking Boots (4660, -0.70 DPS) [world]; Sanguine Sandals (14374, -0.70 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.72 DPS) | yes | Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Band of Purification (12996, -0.40 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.44 DPS) [rep] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 6.3 healing_power points (0.57 DPS) | yes | Volcanic Rock Ring (12053, +0.00 DPS, sim-verified) [world_drop]; Band of Purification (12996, -0.24 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.29 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 13.3 healing_power points (1.21 DPS) | yes | Channeler's Staff (4437, +0.00 DPS, sim-verified) [world]; Staff of Westfall (2042, -0.28 DPS) [quest]; Lesser Staff of the Spire (1300, -0.48 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 2.7 healing_power points (0.24 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.08 DPS) [quest]; Elven Wand (5604, -0.12 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 63.2. Weights run: 7.5s. Verify run: 3.9s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.390 ± 0.007, spirit=0.834 ± 0.004, mp5=1.615 ± 0.015, crit=0.092 ± 0.005 per rating point (14 rating = 1%, 1.293 per %), spell_haste=0.334 ± 0.046

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Pristine Circlet (253975) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.08 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.11 DPS) [world_drop]; Resilient Cap (14401, -0.31 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystal Starfire Medallion (5003, -0.03 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.04 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.13 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 19.5 healing_power points (1.77 DPS) | yes | Mantle of Honor (3560, -0.08 DPS, sim-verified) [quest]; Death Speaker Mantle (6685, -0.38 DPS) [dungeon]; Nightsky Mantle (4718, -0.38 DPS) [world_drop] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 16.3 healing_power points (1.49 DPS) | yes | Prelacy Cape (7004, -0.06 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.24 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.25 DPS) [vendor] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.02 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.89 DPS) [dungeon]; Robes of Arugal (6324, -0.92 DPS) [dungeon] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindthrust Bracers (1974, -0.02 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -0.03 DPS) [world_drop]; Stonecloth Bindings (14416, -0.38 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.75 DPS) | yes | Gloves of Old (9395, +0.00 DPS, sim-verified) [world_drop]; Town Clerk's Mittens (270029, -0.35 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.35 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.6 healing_power points (1.51 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Resilient Cord (14406, -0.45 DPS) [world_drop]; Dreamer's Belt (4829, -0.47 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 35.9 healing_power points (3.27 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.07 DPS) [crafted]; Necromancer Leggings (2277, -1.88 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.1 healing_power points (3.01 DPS) | yes | Nimbus Boots (6998, -0.24 DPS, sim-verified) [quest]; Acidic Walkers (9454, -1.70 DPS) [dungeon]; Pristine Boots (253889, -1.81 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Black Widow Band (6199, -0.75 DPS) [world]; The Queen's Jewel (13094, -0.78 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.88 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 15.6 healing_power points (1.42 DPS) | yes | Black Widow Band (6199, +0.00 DPS, sim-verified) [world]; The Queen's Jewel (13094, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.66 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.07 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -0.38 DPS) [crafted]; Lorekeeper's Staff (212580, -0.42 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 5.6 healing_power points (0.51 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Consecrated Wand (5244, -0.10 DPS) [quest]; Flaring Baton (5326, -0.25 DPS) [quest] |

**New at 30:** head: Filigreed Pristine Circlet; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Caretaker's Cape; chest: Pristine Gown; wrist: Glowing Magical Bracelets; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 70.4. Weights run: 7.7s. Verify run: 4.2s. 328 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.057, intellect=3.559 ± 0.033, spirit=0.357 ± 0.013, mp5=0.747 ± 0.022, crit=0.201 ± 0.014 per rating point (14 rating = 1%, 2.811 per %), spell_haste=not significant (0.020 ± 0.062)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 85.7 healing_power points (3.65 DPS) | yes | Corpseshroud (10574, -0.22 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -0.92 DPS) [dungeon]; Thinking Cap (2624, -1.08 DPS) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.06 DPS, sim-verified) [quest]; Darkspear Warding Pendant (272074, -0.11 DPS) [vendor] |
| shoulder | Windchaser Amice (14432) (or Inquisitor's Shawl (19507)) | World drop [world_drop] | 46.3 healing_power points (1.97 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.23 DPS) [dungeon]; Batwing Mantle (6697, -0.23 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mantle of Lady Falther'ess (23178, -0.06 DPS, sim-verified) [dungeon]; Blackforge Cape (6424, -0.46 DPS) [dungeon]; Cloak of Rot (4462, -0.52 DPS) [world] |
| chest | Red Mageweave Vest (10007) | Tailoring [crafted] | 64.1 healing_power points (2.73 DPS) | yes | Stormcloth Vest (10020, -0.14 DPS) [crafted]; Death Speaker Robes (6682, -0.21 DPS) [dungeon]; Silksand Tunic (14417, -0.23 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 32.0 healing_power points (1.37 DPS) | yes | Mistscape Bracers (4045, -0.11 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -0.15 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.15 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 53.7 healing_power points (2.29 DPS) | yes | Gilded Handwraps (254021, -0.13 DPS, sim-verified) [crafted]; Town Clerk's Mittens (270029, -0.62 DPS) [quest]; Red Mageweave Gloves (10018, -0.77 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 53.4 healing_power points (2.28 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Razzeric's Customized Seatbelt (6726, -0.46 DPS) [quest]; Teacher's Sash (10747, -0.46 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 48.7 healing_power points (2.08 DPS) | yes | Crimson Silk Pantaloons (7062, -0.14 DPS, sim-verified) [crafted]; Aurora Pants (4044, -0.25 DPS) [world_drop]; Red Mageweave Pants (10009, -0.26 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Kodo Rustler Boots (15697, -0.02 DPS) [quest]; Acidic Walkers (9454, -0.03 DPS) [dungeon]; Gilded Slippers (254001, -0.18 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.4 healing_power points (1.17 DPS) | yes | Ogremind Ring (1993, -0.06 DPS) [world_drop]; Mindbender Loop (5009, -0.08 DPS) [world_drop]; Black Widow Band (6199, -0.11 DPS) [world] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 26.0 healing_power points (1.11 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.02 DPS) [world_drop]; Black Widow Band (6199, -0.05 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -0.62 DPS) [dungeon]; Staff of Jordan (873, -1.06 DPS) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 35.6 healing_power points (1.52 DPS) | yes | Aurora Sphere (7610, +0.00 DPS, sim-verified) [dungeon]; Skull of Impending Doom (4984, -0.46 DPS) [quest]; Beacon of Hope (9393, -0.52 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 22.4 healing_power points (0.96 DPS) | yes | Goblin Igniter (5253, -0.11 DPS, sim-verified) [quest]; Flash Wand (5248, -0.30 DPS) [quest]; Starfaller (13063, -0.35 DPS) [world_drop] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Windchaser Amice; back: Darkspear Raider's Cloak; chest: Red Mageweave Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Deathmage Sash; feet: Furen's Boots; finger1: Snake Hoop; finger2: Voodoo Band; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 100.2. Weights run: 8.0s. Verify run: 4.7s. 423 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.055, intellect=3.195 ± 0.048, spirit=0.554 ± 0.037, mp5=0.522 ± 0.033, crit=0.270 ± 0.018 per rating point (14 rating = 1%, 3.779 per %), spell_haste=5.156 ± 0.445

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 87.9 healing_power points (4.31 DPS) | yes | Soulcatcher Halo (10630, +0.00 DPS, sim-verified) [dungeon]; Papal Fez (9431, -0.32 DPS) [dungeon]; Knight-Lieutenant's Satin Cover (220896, -0.72 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 46.9 healing_power points (2.30 DPS) | yes | Gemshard Heart (17707, -0.09 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.73 DPS) [quest]; Darkspear Warding Pendant (272073, -0.89 DPS) [vendor] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 58.9 healing_power points (2.88 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Satin Pads (220894, -0.14 DPS) [vendor]; Nethergeld Shoulders (254049, -0.19 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 47.5 healing_power points (2.33 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.09 DPS, sim-verified) [dungeon]; Imperial Red Cloak (8248, -0.50 DPS) [world_drop]; Keeper's Cloak (14665, -0.76 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 88.2 healing_power points (4.32 DPS) | yes | Embrace of the Wind Serpent (12462, -0.06 DPS, sim-verified) [world]; Knight's Satin Armor (220892, -0.94 DPS) [vendor]; Hibernal Robe (8113, -1.19 DPS) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 51.2 healing_power points (2.51 DPS) | yes | Nethergeld Cuffs (254061, -0.06 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -0.55 DPS) [quest]; Forgotten Wraps (9433, -0.63 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 92.3 healing_power points (4.52 DPS) | yes | Virtuous Mitts (226950, -0.20 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.32 DPS) [crafted]; Stormcloth Gloves (10011, -2.10 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 60.7 healing_power points (2.97 DPS) | yes | Gilded Waistcord (254081, +0.00 DPS, sim-verified) [crafted]; Serenity Belt (13144, -0.31 DPS) [world_drop]; Gilded Cord (254037, -0.48 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (5.8 DPS) | yes | Knight's Satin Leggings (220893, -0.26 DPS, sim-verified) [vendor]; Spellshock Leggings (9484, -0.61 DPS) [dungeon]; Venomshroud Leggings (14444, -0.63 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 59.8 healing_power points (2.93 DPS) | yes | Sergeant Major's Satin Boots (220895, -0.14 DPS, sim-verified) [vendor]; Coldstone Slippers (18697, -0.63 DPS) [dungeon]; Gilded Slippers (254001, -0.74 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 50.7 healing_power points (2.48 DPS) | yes | Cyclopean Band (11824, -0.84 DPS) [dungeon]; Woodseed Hoop (17768, -1.07 DPS) [quest]; Snake Hoop (6750, -1.20 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 38.3 healing_power points (1.88 DPS) | yes | Cyclopean Band (11824, -0.05 DPS, sim-verified) [dungeon]; Woodseed Hoop (17768, -0.47 DPS) [quest]; Snake Hoop (6750, -0.59 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.11 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.17 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -0.11 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -0.80 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -1.18 DPS) [quest]; Death Speaker Scepter (2816, -1.45 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 20.8 healing_power points (1.02 DPS) | yes | Cairnstone Sliver (9654, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.31 DPS) [quest]; Goblin Igniter (5253, -0.31 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Kentic Amice; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 331.2. Weights run: 6.9s. Verify run: 10.8s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.179, intellect=1.043 ± 0.061, spirit=0.550 ± 0.061, mp5=0.486 ± 0.068, crit=0.275 ± 0.026 per rating point (14 rating = 1%, 3.849 per %), spell_haste=not significant (-0.629 ± 0.862)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -2.67 DPS, sim-verified) [crafted] |
| neck | Drake Tooth Necklace (21531) | The Nightmare Manifests [quest] | 38.5 healing_power points (5.91 DPS) | yes | Wavefront Necklace (20685, -0.67 DPS) [world]; The Eye of Zuldazar (19593, -0.77 DPS) [quest]; Animated Chain Necklace (18723, -1.21 DPS, sim-verified) [dungeon] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Mooncloth Shoulders (14139, -2.87 DPS, sim-verified) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.4 healing_power points (8.05 DPS) | yes | Cloak of the Cosmos (18389, -0.87 DPS, sim-verified) [dungeon]; Drape of Recovery (272413, -2.99 DPS) [vendor]; Caretaker's Cape (19530, -3.38 DPS) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | sim-verified (34.5 DPS) | yes | Truefaith Vestments (14154, -0.60 DPS, sim-verified) [crafted]; Field Marshal's Satin Tunic (231624, -0.63 DPS) [vendor]; Virtuous Robe (226945, -2.07 DPS) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.52 DPS) [crafted]; Bracers of Hope (22667, -1.91 DPS, sim-verified) [quest] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handwraps (272097, -0.15 DPS) [vendor]; Desert Bloom Gloves (20717, -0.46 DPS, sim-verified) [quest]; Mooncloth Gloves (18409, -0.53 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 50.6 healing_power points (7.76 DPS) | yes | Whipvine Cord (18327, -0.43 DPS, sim-verified) [dungeon]; Virtuous Belt (226948, -1.45 DPS) [quest]; Gilded Waistcord (254081, -1.87 DPS) [crafted] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 65.8 healing_power points (10.10 DPS) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.37 DPS) [pvp]; Virtuous Skirt (226946, -3.32 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -2.59 DPS, sim-verified) [crafted] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.05 DPS) [vendor]; Band of Piety (22681, -1.20 DPS) [quest]; Naglering (11669, -3.63 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.95 DPS) [vendor]; Band of Piety (22681, -1.11 DPS) [quest]; Naglering (11669, -4.27 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -0.91 DPS) [dungeon]; Mindtap Talisman (18371, -2.19 DPS, sim-verified) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -0.71 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -2.16 DPS) [dungeon]; Hand of Edward the Odd (2243, -3.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 23.0 healing_power points (3.54 DPS) | yes | Bonecreeper Stylus (13938, -0.64 DPS, sim-verified) [dungeon]; Sparkling Crystal Wand (20672, -1.33 DPS) [world]; Oblivion's Touch (18761, -1.78 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Drake Tooth Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Virtuous Bracers; hands: Hands of the Exalted Herald; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Fordring's Seal; finger2: Band of Mending; trinket2: Royal Seal of Eldre'Thalas; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 659.2. Weights run: 4.8s. Verify run: 4.5s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.210, intellect=1.089 ± 0.035, spirit=1.061 ± 0.033, mp5=2.035 ± 0.039, crit=0.346 ± 0.024 per rating point (14 rating = 1%, 4.839 per %), spell_haste=not significant (1.068 ± 0.630)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mooncloth Circlet (14140) | Tailoring [crafted] | sim-verified (+18.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Virtuous Crown (226947, -17.95 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 46.8 healing_power points (14.94 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -2.01 DPS) [quest]; Animated Chain Necklace (18723, -2.38 DPS) [dungeon]; The Eye of Zuldazar (19593, -11.77 DPS, sim-verified) [quest] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -3.77 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.9 healing_power points (16.87 DPS) | yes | Cloak of the Cosmos (18389, -4.76 DPS) [dungeon]; Caretaker's Cape (19530, -5.87 DPS) [rep]; Drape of Recovery (272413, -17.62 DPS, sim-verified) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 90.0 healing_power points (28.71 DPS) | yes | Field Marshal's Satin Tunic (231624, -2.85 DPS) [vendor]; Virtuous Robe (226945, -4.72 DPS) [quest]; Robes of the Exalted (13346, -4.80 DPS, sim-verified) [dungeon] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 42.7 healing_power points (13.63 DPS) | yes | Bracers of Mending (23129, -0.36 DPS) [dungeon]; Virtuous Bracers (226949, -1.03 DPS) [quest]; Nethergeld Cuffs (254061, -3.12 DPS) [crafted] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 61.7 healing_power points (19.70 DPS) | yes | Hands of the Exalted Herald (12554, -0.60 DPS) [dungeon]; Virtuous Mitts (226950, -2.13 DPS) [vendor]; Raider Handwraps (272097, -2.66 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 55.8 healing_power points (17.80 DPS) | yes | Virtuous Belt (226948, -1.54 DPS) [quest]; Penitent's Cinch (272394, -2.34 DPS) [vendor]; Whipvine Cord (18327, -3.53 DPS, sim-verified) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 76.0 healing_power points (24.24 DPS) | yes | Marshal's Satin Legguards (231626, -1.46 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -1.53 DPS) [pvp]; Virtuous Skirt (226946, -42.23 DPS, sim-verified) [quest] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 74.1 healing_power points (23.64 DPS) | yes | Virtuous Sandals (226952, +0.00 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -6.72 DPS) [crafted]; Faith Healer's Boots (22247, -7.11 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -1.26 DPS) [quest]; Blessed Band of Light (272407, -1.75 DPS) [vendor]; Naglering (11669, -46.92 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.45 DPS) [quest]; Blessed Band of Light (272407, -0.94 DPS) [vendor]; Naglering (11669, -39.52 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (+42.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Voodoo Seal (272061, -3.39 DPS) [vendor]; Briarwood Reed (12930, -3.87 DPS) [dungeon]; Mindtap Talisman (18371, -5.99 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Voodoo Seal (272061, -3.66 DPS) [vendor]; Briarwood Reed (12930, -4.15 DPS) [dungeon]; Draconic Infused Emblem (22268, -7.97 DPS, sim-verified) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -3.42 DPS) [dungeon]; Staff of Metanoia (22394, -4.34 DPS) [dungeon]; Hand of Edward the Odd (2243, -46.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 28.7 healing_power points (9.15 DPS) | yes | Sparkling Crystal Wand (20672, -4.43 DPS) [world]; Oblivion's Touch (18761, -5.33 DPS) [dungeon]; Bonecreeper Stylus (13938, -12.37 DPS, sim-verified) [dungeon] |

**New at 60:** head: Mooncloth Circlet; neck: Wavefront Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Desert Bloom Gloves; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Royal Seal of Eldre'Thalas; trinket2: Serenity Field; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 36.7. Weights run: 7.5s. Verify run: 3.8s. 140 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.330 ± 0.006, spirit=0.600 ± 0.003, mp5=1.543 ± 0.010, crit=0.061 ± 0.003 per rating point (14 rating = 1%, 0.856 per %), spell_haste=not significant (-0.002 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, -0.05 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.71 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 2.4 healing_power points (0.22 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.07 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.67 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (20427, -0.04 DPS, sim-verified) [rep]; Pearl-clasped Cloak (5542, -0.12 DPS) [crafted]; Seer's Cape (6378, -0.13 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.4 healing_power points (1.95 DPS) | yes | Corsair's Overshirt (5202, -0.08 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.59 DPS) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.3 healing_power points (1.02 DPS) | yes | Tabitha's Cuffs (251486, +0.00 DPS, sim-verified) [quest]; Featherbead Bracers (15452, -0.42 DPS) [quest]; Bright Bracers (3647, -0.54 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gloves (253913, -0.05 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.30 DPS) [world_drop]; Bright Gloves (3066, -0.42 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.52 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.77 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.58 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Abomination Skin Leggings (23173, -1.61 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Spidersilk Boots (4320, -0.70 DPS) [crafted]; Walking Boots (4660, -0.70 DPS) [world]; Sanguine Sandals (14374, -0.70 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.72 DPS) | yes | Black Pearl Ring (6332, -0.16 DPS) [world]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Band of Purification (12996, -0.40 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.6 healing_power points (0.60 DPS) | yes | Black Pearl Ring (6332, +0.00 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop]; Band of Purification (12996, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 13.3 healing_power points (1.21 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.17 DPS) [quest]; Channeler's Staff (4437, -0.24 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.7 healing_power points (0.24 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.12 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (troll, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 58.3. Weights run: 7.5s. Verify run: 4.0s. 231 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.390 ± 0.007, spirit=0.834 ± 0.004, mp5=1.615 ± 0.015, crit=0.092 ± 0.005 per rating point (14 rating = 1%, 1.293 per %), spell_haste=0.334 ± 0.046

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Pristine Circlet (253975) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.06 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.11 DPS) [world_drop]; Resilient Cap (14401, -0.31 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystal Starfire Medallion (5003, -0.03 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.03 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.13 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 19.5 healing_power points (1.77 DPS) | yes | Ghostly Mantle (3324, -0.19 DPS, sim-verified) [quest]; Death Speaker Mantle (6685, -0.38 DPS) [dungeon]; Nightsky Mantle (4718, -0.38 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 16.3 healing_power points (1.49 DPS) | yes | Glowing Thresher Cape (6901, +0.00 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -0.25 DPS) [vendor]; Cloak of Rot (4462, -0.48 DPS) [world] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.3 healing_power points (3.21 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.09 DPS) [dungeon]; Robes of Arugal (6324, -1.12 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.8 healing_power points (1.07 DPS) | yes | Glowing Magical Bracelets (13106, -0.06 DPS) [world_drop]; Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.31 DPS) [quest] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.75 DPS) | yes | Gloves of Old (9395, +0.00 DPS, sim-verified) [world_drop]; Hotshot Pilot's Gloves (9491, -0.35 DPS) [dungeon]; Pristine Gloves (253913, -0.36 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Sash (253925, -0.05 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Resilient Cord (14406, -0.30 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 35.9 healing_power points (3.27 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.07 DPS) [crafted]; Necromancer Leggings (2277, -1.88 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.1 healing_power points (3.01 DPS) | yes | Acidic Walkers (9454, -0.05 DPS, sim-verified) [dungeon]; Pristine Boots (253889, -1.81 DPS) [crafted]; Frothing Slippers (254003, -1.82 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Black Widow Band (6199, -0.75 DPS) [world]; The Queen's Jewel (13094, -0.78 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.88 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 15.6 healing_power points (1.42 DPS) | yes | Black Widow Band (6199, +0.00 DPS, sim-verified) [world]; The Queen's Jewel (13094, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.66 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.05 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (19569, -0.30 DPS) [pvp]; Glimmering Staff (249392, -0.38 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 5.6 healing_power points (0.51 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Wand of Decay (5252, -0.25 DPS) [quest]; Flaring Baton (5326, -0.25 DPS) [quest] |

**New at 30:** head: Filigreed Pristine Circlet; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 74.0. Weights run: 7.7s. Verify run: 4.2s. 311 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.057, intellect=3.559 ± 0.033, spirit=0.357 ± 0.013, mp5=0.747 ± 0.022, crit=0.201 ± 0.014 per rating point (14 rating = 1%, 2.811 per %), spell_haste=not significant (0.020 ± 0.062)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 85.7 healing_power points (3.65 DPS) | yes | Corpseshroud (10574, +0.00 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -0.92 DPS) [dungeon]; Thinking Cap (2624, -1.08 DPS) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.11 DPS) [vendor]; Prodigious Shadowshard Pendant (17773, -0.14 DPS, sim-verified) [quest] |
| shoulder | Windchaser Amice (14432) (or Inquisitor's Shawl (19507)) | World drop [world_drop] | 46.3 healing_power points (1.97 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.23 DPS) [dungeon]; Batwing Mantle (6697, -0.23 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mantle of Lady Falther'ess (23178, -0.22 DPS, sim-verified) [dungeon]; Blackforge Cape (6424, -0.46 DPS) [dungeon]; Cloak of Rot (4462, -0.52 DPS) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.07 DPS) [dungeon]; Red Mageweave Vest (10007, -0.08 DPS, sim-verified) [crafted]; Silksand Tunic (14417, -0.09 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 32.0 healing_power points (1.37 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [dungeon]; Radiant Silver Bracers (4545, -0.15 DPS) [quest]; Enchanted Stonecloth Bracers (4979, -0.15 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 53.7 healing_power points (2.29 DPS) | yes | Gilded Handwraps (254021, +0.00 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.77 DPS) [crafted]; Mistscape Gloves (6428, -0.92 DPS) [dungeon] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Deathmage Sash (10771, -0.13 DPS, sim-verified) [dungeon]; Razzeric's Customized Seatbelt (6726, -0.42 DPS) [quest]; Mistscape Sash (4736, -0.57 DPS) [dungeon] |
| legs | Crimson Silk Pantaloons (7062) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.06 DPS, sim-verified) [crafted]; Aurora Pants (4044, -0.15 DPS) [world_drop]; Red Mageweave Pants (10009, -0.15 DPS) [crafted] |
| feet | Boots of the Maharishi (9658) | Gordunni Cobalt [quest] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Slippers (254001, -0.10 DPS, sim-verified) [crafted]; Furen's Boots (13100, -0.15 DPS) [world_drop]; Kodo Rustler Boots (15697, -0.17 DPS) [quest] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.4 healing_power points (1.17 DPS) | yes | Ogremind Ring (1993, -0.06 DPS) [world_drop]; Mindbender Loop (5009, -0.08 DPS) [world_drop]; Black Widow Band (6199, -0.11 DPS) [world] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 26.0 healing_power points (1.11 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.02 DPS) [world_drop]; Black Widow Band (6199, -0.05 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -0.62 DPS) [dungeon]; Staff of Jordan (873, -1.06 DPS) [world_drop] |
| off_hand | Aurora Sphere (7610) | Uldaman: Ancient Treasure [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Prophetic Cane (6803, -0.04 DPS, sim-verified) [quest]; Skull of Impending Doom (4984, -0.06 DPS) [quest]; Witch's Finger (16887, -0.06 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 22.4 healing_power points (0.96 DPS) | yes | Goblin Igniter (5253, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.30 DPS) [quest]; Starfaller (13063, -0.35 DPS) [world_drop] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Windchaser Amice; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Crimson Silk Pantaloons; feet: Boots of the Maharishi; finger1: Snake Hoop; finger2: Voodoo Band; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Aurora Sphere; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (troll, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 93.9. Weights run: 8.0s. Verify run: 4.5s. 403 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.055, intellect=3.195 ± 0.048, spirit=0.554 ± 0.037, mp5=0.522 ± 0.033, crit=0.270 ± 0.018 per rating point (14 rating = 1%, 3.779 per %), spell_haste=5.156 ± 0.445

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 87.9 healing_power points (4.31 DPS) | yes | Soulcatcher Halo (10630, -0.12 DPS) [dungeon]; Papal Fez (9431, -0.32 DPS) [dungeon]; Blood Guard's Satin Cover (220899, -0.72 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 46.9 healing_power points (2.30 DPS) | yes | Gemshard Heart (17707, -0.12 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.73 DPS) [quest]; Darkspear Warding Pendant (272073, -0.89 DPS) [vendor] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 58.9 healing_power points (2.88 DPS) | yes | Rotgrip Mantle (17732, -0.07 DPS) [dungeon]; Blood Guard's Satin Pads (220901, -0.14 DPS) [vendor]; Nethergeld Shoulders (254049, -0.19 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 47.5 healing_power points (2.33 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.12 DPS, sim-verified) [dungeon]; Imperial Red Cloak (8248, -0.50 DPS) [world_drop]; Keeper's Cloak (14665, -0.76 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 88.2 healing_power points (4.32 DPS) | yes | Embrace of the Wind Serpent (12462, -0.16 DPS, sim-verified) [world]; Stone Guard's Satin Armor (220903, -0.94 DPS) [vendor]; Hibernal Robe (8113, -1.19 DPS) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 51.2 healing_power points (2.51 DPS) | yes | Nethergeld Cuffs (254061, -0.06 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -0.55 DPS) [quest]; Forgotten Wraps (9433, -0.63 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 92.3 healing_power points (4.52 DPS) | yes | Virtuous Mitts (226950, -0.36 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.32 DPS) [crafted]; Greenleaf Handwraps (19116, -1.72 DPS) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 60.7 healing_power points (2.97 DPS) | yes | Gilded Waistcord (254081, +0.00 DPS, sim-verified) [crafted]; Serenity Belt (13144, -0.31 DPS) [world_drop]; Gilded Cord (254037, -0.48 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (5.0 DPS) | yes | Stone Guard's Satin Leggings (220902, -0.20 DPS, sim-verified) [vendor]; Spellshock Leggings (9484, -0.61 DPS) [dungeon]; Venomshroud Leggings (14444, -0.63 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 59.8 healing_power points (2.93 DPS) | yes | First Sergeant's Satin Boots (220900, -0.15 DPS, sim-verified) [vendor]; Coldstone Slippers (18697, -0.63 DPS) [dungeon]; Gilded Slippers (254001, -0.74 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 50.7 healing_power points (2.48 DPS) | yes | Cyclopean Band (11824, -0.84 DPS) [dungeon]; Woodseed Hoop (17768, -1.07 DPS) [quest]; Snake Hoop (6750, -1.20 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 38.3 healing_power points (1.88 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Woodseed Hoop (17768, -0.47 DPS) [quest]; Snake Hoop (6750, -0.59 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.11 DPS) [quest]; Alchemists' Stone (13503, -0.27 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.05 DPS) [quest]; Alchemists' Stone (13503, -0.22 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -0.78 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -1.18 DPS) [quest]; Death Speaker Scepter (2816, -1.45 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 20.8 healing_power points (1.02 DPS) | yes | Nature's Breath (19118, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.31 DPS) [quest]; Goblin Igniter (5253, -0.31 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Kentic Amice; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 323.2. Weights run: 6.9s. Verify run: 10.8s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.179, intellect=1.043 ± 0.061, spirit=0.550 ± 0.061, mp5=0.486 ± 0.068, crit=0.275 ± 0.026 per rating point (14 rating = 1%, 3.849 per %), spell_haste=not significant (-0.629 ± 0.862)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Devout Crown (16693, -2.47 DPS, sim-verified) [dungeon] |
| neck | Drake Tooth Necklace (21531) | The Nightmare Manifests [quest] | 38.5 healing_power points (5.91 DPS) | yes | Wavefront Necklace (20685, -0.67 DPS) [world]; The Eye of Zuldazar (19593, -0.77 DPS) [quest]; Animated Chain Necklace (18723, -0.96 DPS, sim-verified) [dungeon] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Devout Mantle (16695, -2.13 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.4 healing_power points (8.05 DPS) | yes | Cloak of the Cosmos (18389, -1.18 DPS, sim-verified) [dungeon]; Drape of Recovery (272413, -2.99 DPS) [vendor]; Battle Healer's Cloak (19526, -3.38 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 81.8 healing_power points (12.56 DPS) | yes | Robes of the Exalted (13346, -0.25 DPS, sim-verified) [dungeon]; Warlord's Satin Tunic (231632, -1.02 DPS) [vendor]; Virtuous Robe (226945, -2.46 DPS) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.52 DPS) [crafted]; Bracers of Hope (22667, -1.85 DPS, sim-verified) [quest] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handwraps (272097, -0.15 DPS) [vendor]; Mooncloth Gloves (18409, -0.53 DPS) [crafted]; Desert Bloom Gloves (20717, -0.70 DPS, sim-verified) [quest] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 50.6 healing_power points (7.76 DPS) | yes | Whipvine Cord (18327, -0.69 DPS, sim-verified) [dungeon]; Virtuous Belt (226948, -1.45 DPS) [quest]; Gilded Waistcord (254081, -1.87 DPS) [crafted] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 65.8 healing_power points (10.10 DPS) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.37 DPS) [pvp]; Virtuous Skirt (226946, -3.31 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; General's Satin Walkers (231630, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -1.85 DPS, sim-verified) [crafted] |
| finger1 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.05 DPS) [vendor]; Band of Piety (22681, -1.20 DPS) [quest]; Naglering (11669, -3.48 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.95 DPS) [vendor]; Band of Piety (22681, -1.11 DPS) [quest]; Naglering (11669, -3.68 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Mindtap Talisman (18371, -3.02 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (29.9 DPS) | yes | Mindtap Talisman (18371, -0.69 DPS, sim-verified) [dungeon]; Royal Seal of Eldre'Thalas (18469, -1.08 DPS) [quest]; Briarwood Reed (12930, -2.00 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -0.71 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -2.16 DPS) [dungeon]; Hand of Edward the Odd (2243, -2.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 23.0 healing_power points (3.54 DPS) | yes | Bonecreeper Stylus (13938, -0.87 DPS, sim-verified) [dungeon]; Sparkling Crystal Wand (20672, -1.33 DPS) [world]; Oblivion's Touch (18761, -1.78 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Drake Tooth Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Hands of the Exalted Herald; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Fordring's Seal; finger2: Band of Mending; trinket2: Serenity Field; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 650.7. Weights run: 4.8s. Verify run: 6.4s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.210, intellect=1.089 ± 0.035, spirit=1.061 ± 0.033, mp5=2.035 ± 0.039, crit=0.346 ± 0.024 per rating point (14 rating = 1%, 4.839 per %), spell_haste=not significant (1.068 ± 0.630)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | sim-verified (369.2 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -24.00 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 46.8 healing_power points (14.94 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -2.01 DPS) [quest]; Animated Chain Necklace (18723, -2.38 DPS) [dungeon]; The Eye of Zuldazar (19593, -9.59 DPS, sim-verified) [quest] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (369.2 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -47.63 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.9 healing_power points (16.87 DPS) | yes | Cloak of the Cosmos (18389, -4.76 DPS) [dungeon]; Battle Healer's Cloak (19526, -5.87 DPS) [rep]; Drape of Recovery (272413, -17.43 DPS, sim-verified) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 90.0 healing_power points (28.71 DPS) | yes | Warlord's Satin Tunic (231632, -2.85 DPS) [vendor]; Robes of the Exalted (13346, -4.31 DPS, sim-verified) [dungeon]; Virtuous Robe (226945, -4.72 DPS) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (369.2 DPS) | yes | Bracers of Mending (23129, +0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.09 DPS) [crafted]; Bracers of Hope (22667, -14.50 DPS, sim-verified) [quest] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 61.7 healing_power points (19.70 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS, sim-verified) [dungeon]; Virtuous Mitts (226950, -2.13 DPS) [vendor]; Raider Handwraps (272097, -2.66 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 55.8 healing_power points (17.80 DPS) | yes | Whipvine Cord (18327, -1.33 DPS, sim-verified) [dungeon]; Virtuous Belt (226948, -1.54 DPS) [quest]; Penitent's Cinch (272394, -2.34 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 76.0 healing_power points (24.24 DPS) | yes | General's Satin Legguards (231634, -1.46 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -1.53 DPS) [pvp]; Virtuous Skirt (226946, -41.37 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (369.2 DPS) | yes | Mooncloth Boots (15802, -2.04 DPS) [crafted]; Faith Healer's Boots (22247, -2.44 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -47.01 DPS, sim-verified) [vendor] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (369.2 DPS) | yes | Band of Piety (22681, -1.26 DPS) [quest]; Blessed Band of Light (272407, -1.75 DPS) [vendor]; Naglering (11669, -42.14 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (369.2 DPS) | yes | Band of Piety (22681, -0.45 DPS) [quest]; Blessed Band of Light (272407, -0.94 DPS) [vendor]; Naglering (11669, -42.21 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (369.2 DPS) | yes | Darkspear Voodoo Seal (272061, -3.39 DPS) [vendor]; Briarwood Reed (12930, -3.87 DPS) [dungeon]; Mindtap Talisman (18371, -5.99 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (369.2 DPS) | yes | Darkspear Voodoo Seal (272061, -3.66 DPS) [vendor]; Briarwood Reed (12930, -4.15 DPS) [dungeon]; Draconic Infused Emblem (22268, -13.66 DPS, sim-verified) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (369.2 DPS) | yes | Death Speaker Scepter (2816, -3.42 DPS) [dungeon]; Staff of Metanoia (22394, -4.34 DPS) [dungeon]; Hand of Edward the Odd (2243, -49.14 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 28.7 healing_power points (9.15 DPS) | yes | Sparkling Crystal Wand (20672, -4.43 DPS) [world]; Oblivion's Touch (18761, -5.33 DPS) [dungeon]; Bonecreeper Stylus (13938, -10.03 DPS, sim-verified) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Royal Seal of Eldre'Thalas; trinket2: Serenity Field; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

