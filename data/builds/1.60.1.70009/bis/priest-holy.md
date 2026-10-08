# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 39.3. Weights run: 7.7s. Verify run: 3.7s. 150 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.330 ± 0.006, spirit=0.667 ± 0.005, mp5=1.558 ± 0.013, crit=0.060 ± 0.003 per rating point (14 rating = 1%, 0.836 per %), spell_haste=0.086 ± 0.019

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, -0.03 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.72 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 2.7 healing_power points (0.24 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.09 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.66 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (20428, -0.03 DPS, sim-verified) [rep]; Seer's Cape (6378, -0.12 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.12 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.7 healing_power points (1.96 DPS) | yes | Corsair's Overshirt (5202, -0.14 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.54 DPS) [dungeon]; Bloody Apron (6226, -0.78 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.3 healing_power points (1.03 DPS) | yes | Bright Bracers (3647, +0.00 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.67 DPS) [quest]; Seer's Cuffs (3645, -0.79 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Magefist Gloves (12977, +0.00 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.64 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.64 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.52 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.76 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.6 healing_power points (2.60 DPS) | yes | Darkweave Breeches (12987, -0.07 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.63 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Kimbra Boots (6191, +0.00 DPS, sim-verified) [quest]; Bluegill Sandals (1560, -0.70 DPS) [world]; Sanguine Sandals (14374, -0.70 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.72 DPS) | yes | Band of Purification (12996, -0.36 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Deep Fathom Ring (6463, -0.42 DPS) [dungeon] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 6.7 healing_power points (0.60 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop]; Deep Fathom Ring (6463, -0.30 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 13.3 healing_power points (1.21 DPS) | yes | Staff of Westfall (2042, -0.08 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.24 DPS) [world]; Lesser Staff of the Spire (1300, -0.48 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 2.7 healing_power points (0.24 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.06 DPS) [quest]; Dwarven Flamestick (5241, -0.12 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 66.7. Weights run: 7.6s. Verify run: 3.9s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.014, intellect=1.312 ± 0.008, spirit=0.835 ± 0.004, mp5=1.564 ± 0.014, crit=0.100 ± 0.006 per rating point (14 rating = 1%, 1.403 per %), spell_haste=not significant (-0.055 ± 0.046)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Pristine Circlet (253975) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.12 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.20 DPS) [world_drop]; Pristine Circlet (253949, -0.38 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.04 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.12 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 18.6 healing_power points (1.77 DPS) | yes | Mantle of Honor (3560, -0.08 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.37 DPS) [world_drop]; Death Speaker Mantle (6685, -0.40 DPS) [dungeon] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 16.3 healing_power points (1.56 DPS) | yes | Prelacy Cape (7004, -0.08 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.25 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.32 DPS) [vendor] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.05 DPS, sim-verified) [dungeon]; Robes of Arugal (6324, -0.91 DPS) [dungeon]; Beguiler Robes (7728, -0.96 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.3 healing_power points (1.07 DPS) | yes | Glowing Magical Bracelets (13106, +0.00 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.08 DPS) [world_drop]; Stonecloth Bindings (14416, -0.45 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.9 healing_power points (1.80 DPS) | yes | Gloves of Old (9395, -0.22 DPS) [world_drop]; Pristine Gloves (253913, -0.38 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.41 DPS) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.2 healing_power points (1.55 DPS) | yes | Novice Ardent's Sash (253887, -0.07 DPS, sim-verified) [crafted]; Resilient Cord (14406, -0.48 DPS) [world_drop]; Dreamer's Belt (4829, -0.51 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 35.4 healing_power points (3.37 DPS) | yes | Filigreed Pristine Leggings (253937, -0.07 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.06 DPS) [crafted]; Necromancer Leggings (2277, -1.99 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.5 healing_power points (3.10 DPS) | yes | Nimbus Boots (6998, -0.31 DPS, sim-verified) [quest]; Acidic Walkers (9454, -1.78 DPS) [dungeon]; Pristine Boots (253889, -1.86 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.71 DPS) | yes | The Queen's Jewel (13094, -0.83 DPS) [world_drop]; Black Widow Band (6199, -0.84 DPS) [world]; Lavishly Jeweled Ring (1156, -0.96 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 15.0 healing_power points (1.43 DPS) | yes | The Queen's Jewel (13094, -0.12 DPS, sim-verified) [world_drop]; Black Widow Band (6199, -0.56 DPS) [world]; Lavishly Jeweled Ring (1156, -0.68 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.07 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -0.44 DPS) [crafted]; Lorekeeper's Staff (212580, -0.47 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 5.2 healing_power points (0.50 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Consecrated Wand (5244, -0.09 DPS) [quest]; Flaring Baton (5326, -0.25 DPS) [quest] |

**New at 30:** head: Filigreed Pristine Circlet; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Caretaker's Cape; chest: Pristine Gown; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 93.5. Weights run: 7.9s. Verify run: 4.0s. 328 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.073, intellect=1.728 ± 0.039, spirit=0.581 ± 0.017, mp5=0.854 ± 0.021, crit=0.173 ± 0.014 per rating point (14 rating = 1%, 2.426 per %), spell_haste=2.593 ± 0.224

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 56.6 healing_power points (3.79 DPS) | yes | Holy Shroud (2721, -1.24 DPS, sim-verified) [world_drop]; Corpseshroud (10574, -1.36 DPS) [dungeon]; Miner's Hat of the Deep (9429, -1.43 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.08 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.20 DPS, sim-verified) [quest] |
| shoulder | Windchaser Amice (14432) | World drop [world_drop] | sim-verified (6.4 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.04 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.11 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 24.6 healing_power points (1.64 DPS) | yes | Darkspear Raider's Cloak (272077, +0.00 DPS, sim-verified) [vendor]; Caretaker's Cape (19532, -0.24 DPS) [rep]; Blackforge Cape (6424, -0.56 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.6 healing_power points (4.19 DPS) | yes | Death Speaker Robes (6682, -0.68 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -1.89 DPS) [crafted]; Red Mageweave Vest (10007, -2.11 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 15.6 healing_power points (1.04 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -0.12 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.12 DPS) [quest] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 34.1 healing_power points (2.28 DPS) | yes | Stormcloth Gloves (10011, -0.16 DPS) [crafted]; Earthen Silk Gloves (254017, -0.77 DPS) [crafted]; Truefaith Gloves (7049, -0.93 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 39.3 healing_power points (2.63 DPS) | yes | Deathmage Sash (10771, -0.31 DPS, sim-verified) [dungeon]; Sutarn's Ring (13105, -1.24 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.24 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 37.0 healing_power points (2.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.10 DPS, sim-verified) [crafted]; Aurora Pants (4044, -0.81 DPS) [world_drop]; Stormcloth Pants (10010, -0.89 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 34.4 healing_power points (2.30 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.10 DPS) [dungeon]; Nimbus Boots (6998, -1.17 DPS) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.20 DPS) | yes | Ogremind Ring (1993, -0.28 DPS) [world_drop]; Voodoo Band (1996, -0.28 DPS) [world]; Mindbender Loop (5009, -0.32 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 16.2 healing_power points (1.08 DPS) | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world]; Ogremind Ring (1993, -0.16 DPS) [world_drop]; Mindbender Loop (5009, -0.19 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -2.81 DPS) [dungeon]; Staff of Jordan (873, -2.85 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 24.3 healing_power points (1.63 DPS) | yes | Orb of Lorica (11262, +0.00 DPS, sim-verified) [quest]; Orb of Souls (249395, -0.52 DPS) [crafted]; Aurora Sphere (7610, -0.66 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 12.1 healing_power points (0.81 DPS) | yes | Goblin Igniter (5253, -0.08 DPS, sim-verified) [quest]; Flash Wand (5248, -0.23 DPS) [quest]; Captain Rackmore's Tiller (16789, -0.29 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Windchaser Amice; back: Mantle of Lady Falther'ess; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Gilded Handwraps; waist: Gilded Cord; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 117.6. Weights run: 7.9s. Verify run: 3.9s. 423 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.087, intellect=3.463 ± 0.029, spirit=0.331 ± 0.015, mp5=1.134 ± 0.032, crit=0.223 ± 0.016 per rating point (14 rating = 1%, 3.120 per %), spell_haste=not significant (0.243 ± 0.160)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 94.5 healing_power points (5.76 DPS) | yes | Soulcatcher Halo (10630, -0.28 DPS) [dungeon]; Papal Fez (9431, -0.65 DPS) [dungeon]; Knight-Lieutenant's Satin Cover (220896, -0.90 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 49.8 healing_power points (3.04 DPS) | yes | Gemshard Heart (17707, -0.31 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.93 DPS) [quest]; Darkspear Warding Pendant (272073, -1.14 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 62.3 healing_power points (3.80 DPS) | yes | Kentic Amice (11624, -0.08 DPS) [dungeon]; Knight-Lieutenant's Satin Pads (220894, -0.31 DPS, sim-verified) [vendor]; Nethergeld Shoulders (254049, -0.41 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 50.1 healing_power points (3.06 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.08 DPS, sim-verified) [dungeon]; Imperial Red Cloak (8248, -0.65 DPS) [world_drop]; Keeper's Cloak (14665, -0.95 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 91.5 healing_power points (5.58 DPS) | yes | Knight's Satin Armor (220892, -0.87 DPS, sim-verified) [vendor]; Hibernal Robe (8113, -1.36 DPS) [world_drop]; Acumen Robes (17775, -1.36 DPS) [quest] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 53.9 healing_power points (3.29 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -0.69 DPS) [quest]; Forgotten Wraps (9433, -0.75 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 98.2 healing_power points (5.99 DPS) | yes | Virtuous Mitts (226950, -0.93 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.95 DPS) [crafted]; Stormcloth Gloves (10011, -2.78 DPS) [crafted] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Belt (13144, -0.08 DPS) [world_drop]; Dawnspire Cord (12466, -0.26 DPS, sim-verified) [dungeon]; Imperial Red Sash (8253, -0.50 DPS) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight's Satin Leggings (220893, -0.37 DPS, sim-verified) [vendor]; Spellshock Leggings (9484, -0.65 DPS) [dungeon]; Venomshroud Leggings (14444, -0.75 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 62.2 healing_power points (3.79 DPS) | yes | Coldstone Slippers (18697, -0.56 DPS) [dungeon]; Sergeant Major's Satin Boots (220895, -0.83 DPS, sim-verified) [vendor]; Gilded Slippers (254001, -1.01 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 53.6 healing_power points (3.27 DPS) | yes | Cyclopean Band (11824, -1.16 DPS) [dungeon]; Woodseed Hoop (17768, -1.37 DPS) [quest]; Snake Hoop (6750, -1.65 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 41.6 healing_power points (2.53 DPS) | yes | Cyclopean Band (11824, -0.43 DPS) [dungeon]; Woodseed Hoop (17768, -0.63 DPS) [quest]; Snake Hoop (6750, -0.91 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.84 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.88 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.04 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.08 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -1.45 DPS) [quest]; Barman Shanker (12791, -1.94 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -2.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 21.8 healing_power points (1.33 DPS) | yes | Cairnstone Sliver (9654, -0.09 DPS, sim-verified) [quest]; Flash Wand (5248, -0.42 DPS) [quest]; Goblin Igniter (5253, -0.42 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 330.2. Weights run: 6.9s. Verify run: 3.0s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.360, intellect=1.877 ± 0.062, spirit=0.906 ± 0.071, mp5=1.181 ± 0.066, crit=0.293 ± 0.035 per rating point (14 rating = 1%, 4.107 per %), spell_haste=not significant (1.302 ± 1.075)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Virtuous Crown (226947, -2.21 DPS, sim-verified) [quest] |
| neck | Drake Tooth Necklace (21531) | The Nightmare Manifests [quest] | 48.5 healing_power points (6.74 DPS) | yes | The Eye of Zuldazar (19593, +0.00 DPS, sim-verified) [quest]; The All-Seeing Eye of Zuldazar (19594, -0.52 DPS) [quest]; Lady Maye's Pendant (14558, -0.53 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -5.45 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 60.8 healing_power points (8.45 DPS) | yes | Drape of Recovery (272413, -3.48 DPS) [vendor]; Darkspear Raider's Cloak (272063, -3.52 DPS) [vendor]; Cloak of the Cosmos (18389, -4.02 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 87.5 healing_power points (12.16 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Virtuous Robe (226945, -0.38 DPS) [quest]; Robes of the Exalted (13346, -0.53 DPS, sim-verified) [dungeon] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 49.5 healing_power points (6.88 DPS) | yes | Bracers of Mending (23129, -0.27 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -0.65 DPS) [quest]; Marshal's Satin Bracers (17606, -0.67 DPS) [pvp] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 74.7 healing_power points (10.38 DPS) | yes | Hands of the Exalted Herald (12554, -1.08 DPS, sim-verified) [dungeon]; Desert Bloom Gloves (20717, -1.36 DPS) [quest]; Mooncloth Gloves (18409, -1.62 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 65.4 healing_power points (9.09 DPS) | yes | Marshal's Satin Sash (17609, -1.21 DPS) [pvp]; Devout Belt (16696, -1.21 DPS) [dungeon]; Virtuous Belt (226948, -5.44 DPS, sim-verified) [quest] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 86.6 healing_power points (12.04 DPS) | yes | Marshal's Satin Legguards (231626, -0.01 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.80 DPS) [pvp]; Virtuous Skirt (226946, -9.40 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Walkers (231627, -0.29 DPS) [vendor]; Mooncloth Boots (15802, -0.39 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -3.05 DPS, sim-verified) [vendor] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.57 DPS) [quest]; Emerald Flame Ring (18395, -0.67 DPS) [dungeon]; Naglering (11669, -11.20 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.12 DPS) [quest]; Emerald Flame Ring (18395, -0.22 DPS) [dungeon]; Naglering (11669, -10.94 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -1.21 DPS) [dungeon]; Second Wind (11819, -2.19 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor]; Mindtap Talisman (18371, -2.57 DPS, sim-verified) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -0.27 DPS) [dungeon]; Death Speaker Scepter (2816, -1.23 DPS) [dungeon]; Hand of Edward the Odd (2243, -9.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.0 healing_power points (3.75 DPS) | yes | Oblivion's Touch (18761, -0.88 DPS) [dungeon]; Bonecreeper Stylus (13938, -1.18 DPS) [dungeon]; Sparkling Crystal Wand (20672, -2.33 DPS, sim-verified) [world] |

**New at 60:** head: Devout Crown; neck: Drake Tooth Necklace; shoulder: Devout Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 514.0. Weights run: 4.3s. Verify run: 1.9s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.267, intellect=0.926 ± 0.027, spirit=0.885 ± 0.026, mp5=1.838 ± 0.029, crit=0.282 ± 0.021 per rating point (14 rating = 1%, 3.943 per %), spell_haste=1.769 ± 0.438

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 66.8 healing_power points (27.78 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -1.07 DPS) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 44.3 healing_power points (18.40 DPS) | yes | Animated Chain Necklace (18723, -2.47 DPS) [dungeon]; Drake Tooth Necklace (21531, -2.97 DPS) [quest]; The Eye of Zuldazar (19593, -3.05 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 78.9 healing_power points (32.79 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -10.13 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -12.62 DPS) [pvp]; Mooncloth Shoulders (14139, -13.36 DPS) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 51.3 healing_power points (21.31 DPS) | yes | Drape of Recovery (272413, -0.74 DPS, sim-verified) [vendor]; Cloak of the Cosmos (18389, -6.27 DPS) [dungeon]; Caretaker's Cape (19530, -7.56 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 87.2 healing_power points (36.23 DPS) | yes | Robes of the Exalted (13346, -1.99 DPS) [dungeon]; Field Marshal's Satin Tunic (231624, -5.10 DPS) [vendor]; Virtuous Robe (226945, -7.53 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 38.8 healing_power points (16.13 DPS) | yes | Bracers of Mending (23129, -0.40 DPS) [dungeon]; Virtuous Bracers (226949, -1.14 DPS) [quest]; Nethergeld Cuffs (254061, -3.28 DPS) [crafted] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 60.1 healing_power points (24.96 DPS) | yes | Hands of the Exalted Herald (12554, -1.83 DPS) [dungeon]; Virtuous Mitts (226950, -4.13 DPS) [vendor]; Raider Handwraps (272097, -4.59 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 51.9 healing_power points (21.59 DPS) | yes | Whipvine Cord (18327, -0.69 DPS, sim-verified) [dungeon]; Virtuous Belt (226948, -2.33 DPS) [quest]; Penitent's Cinch (272394, -3.20 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 71.5 healing_power points (29.74 DPS) | yes | Marshal's Satin Legguards (231626, -2.13 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -2.19 DPS) [pvp]; Virtuous Skirt (226946, -4.07 DPS) [quest] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 69.4 healing_power points (28.83 DPS) | yes | Virtuous Sandals (226952, -6.63 DPS) [quest]; Mooncloth Boots (15802, -8.68 DPS) [crafted]; Faith Healer's Boots (22247, -8.99 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (514.0 DPS) | yes | Naglering (11669, -0.61 DPS, sim-verified) [dungeon]; Band of Piety (22681, -1.74 DPS) [quest]; Rosewine Circle (13178, -1.91 DPS) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (514.0 DPS) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Band of Piety (22681, -1.25 DPS) [quest]; Rosewine Circle (13178, -1.43 DPS) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (514.0 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (514.0 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -3.58 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -6.14 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.7 healing_power points (11.11 DPS) | yes | Bonecreeper Stylus (13938, -5.00 DPS) [dungeon]; Sparkling Crystal Wand (20672, -5.57 DPS) [world]; Ritssyn's Wand of Bad Mojo (22408, -6.54 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Desert Bloom Gloves; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Draconic Infused Emblem; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 36.9. Weights run: 7.7s. Verify run: 3.8s. 140 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.330 ± 0.006, spirit=0.667 ± 0.005, mp5=1.558 ± 0.013, crit=0.060 ± 0.003 per rating point (14 rating = 1%, 0.836 per %), spell_haste=0.086 ± 0.019

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, -0.04 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.72 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 2.7 healing_power points (0.24 DPS) | yes | Roadwatcher's Confidence (281265, -0.06 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.07 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.66 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (20427, -0.03 DPS, sim-verified) [rep]; Seer's Cape (6378, -0.12 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.12 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.7 healing_power points (1.96 DPS) | yes | Corsair's Overshirt (5202, -0.10 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.54 DPS) [dungeon]; Bloody Apron (6226, -0.78 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.3 healing_power points (1.03 DPS) | yes | Tabitha's Cuffs (251486, +0.00 DPS, sim-verified) [quest]; Featherbead Bracers (15452, -0.42 DPS) [quest]; Bright Bracers (3647, -0.54 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gloves (253913, -0.05 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.30 DPS) [world_drop]; Bright Gloves (3066, -0.42 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.52 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.76 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.6 healing_power points (2.60 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Filigreed Flame Leggings (253941, -1.63 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Bluegill Sandals (1560, +0.00 DPS, sim-verified) [world]; Walking Boots (4660, -0.70 DPS) [world]; Sanguine Sandals (14374, -0.70 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.72 DPS) | yes | Black Pearl Ring (6332, -0.08 DPS, sim-verified) [world]; Band of Purification (12996, -0.36 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Pearl Ring (6332, -0.04 DPS, sim-verified) [world]; Band of Purification (12996, -0.24 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 13.3 healing_power points (1.21 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.14 DPS) [quest]; Channeler's Staff (4437, -0.24 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.7 healing_power points (0.24 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.12 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (troll, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 63.1. Weights run: 7.6s. Verify run: 3.9s. 231 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.014, intellect=1.312 ± 0.008, spirit=0.835 ± 0.004, mp5=1.564 ± 0.014, crit=0.100 ± 0.006 per rating point (14 rating = 1%, 1.403 per %), spell_haste=not significant (-0.055 ± 0.046)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Pristine Circlet (253975) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.10 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.20 DPS) [world_drop]; Pristine Circlet (253949, -0.38 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.03 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.12 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 18.6 healing_power points (1.77 DPS) | yes | Ghostly Mantle (3324, -0.21 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.37 DPS) [world_drop]; Death Speaker Mantle (6685, -0.40 DPS) [dungeon] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 16.3 healing_power points (1.56 DPS) | yes | Glowing Thresher Cape (6901, -0.06 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -0.32 DPS) [vendor]; Cloak of Rot (4462, -0.56 DPS) [world] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.03 DPS, sim-verified) [dungeon]; Robes of Arugal (6324, -0.91 DPS) [dungeon]; Beguiler Robes (7728, -0.96 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.3 healing_power points (1.07 DPS) | yes | Glowing Magical Bracelets (13106, +0.00 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.08 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.32 DPS) [quest] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.9 healing_power points (1.80 DPS) | yes | Gloves of Old (9395, +0.00 DPS, sim-verified) [world_drop]; Tattered Mittens (270030, -0.37 DPS) [quest]; Pristine Gloves (253913, -0.38 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Sash (253925, -0.07 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Resilient Cord (14406, -0.30 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 35.4 healing_power points (3.37 DPS) | yes | Filigreed Pristine Leggings (253937, -0.05 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.06 DPS) [crafted]; Necromancer Leggings (2277, -1.99 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.5 healing_power points (3.10 DPS) | yes | Acidic Walkers (9454, -0.09 DPS, sim-verified) [dungeon]; Pristine Boots (253889, -1.86 DPS) [crafted]; Frothing Slippers (254003, -1.90 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.71 DPS) | yes | The Queen's Jewel (13094, -0.83 DPS) [world_drop]; Black Widow Band (6199, -0.84 DPS) [world]; Lavishly Jeweled Ring (1156, -0.96 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 15.0 healing_power points (1.43 DPS) | yes | The Queen's Jewel (13094, -0.12 DPS, sim-verified) [world_drop]; Black Widow Band (6199, -0.56 DPS) [world]; Lavishly Jeweled Ring (1156, -0.68 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.06 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (19569, -0.35 DPS) [pvp]; Glimmering Staff (249392, -0.44 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 5.2 healing_power points (0.50 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Wand of Decay (5252, -0.25 DPS) [quest]; Flaring Baton (5326, -0.25 DPS) [quest] |

**New at 30:** head: Filigreed Pristine Circlet; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Battle Healer's Cloak; chest: Pristine Gown; hands: Truefaith Gloves; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 93.7. Weights run: 7.9s. Verify run: 4.1s. 311 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.073, intellect=1.728 ± 0.039, spirit=0.581 ± 0.017, mp5=0.854 ± 0.021, crit=0.173 ± 0.014 per rating point (14 rating = 1%, 2.426 per %), spell_haste=2.593 ± 0.224

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 56.6 healing_power points (3.79 DPS) | yes | Corpseshroud (10574, -1.36 DPS) [dungeon]; Miner's Hat of the Deep (9429, -1.43 DPS) [dungeon]; Holy Shroud (2721, -1.48 DPS, sim-verified) [world_drop] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 17.3 healing_power points (1.16 DPS) | yes | Necklace of Calisea (1714, -0.07 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.15 DPS) [dungeon]; Triune Amulet (7722, -0.16 DPS, sim-verified) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 22.6 healing_power points (1.51 DPS) | yes | Inquisitor's Shawl (19507, -0.01 DPS) [dungeon]; Mistscape Mantle (4734, -0.05 DPS) [dungeon]; Windchaser Amice (14432, -0.08 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (5.7 DPS) | yes | Battle Healer's Cloak (19528, -0.03 DPS) [rep]; Mantle of Lady Falther'ess (23178, -0.14 DPS, sim-verified) [dungeon]; Blackforge Cape (6424, -0.35 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 62.6 healing_power points (4.19 DPS) | yes | Death Speaker Robes (6682, -0.71 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -1.89 DPS) [crafted]; Red Mageweave Vest (10007, -2.11 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 15.6 healing_power points (1.04 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [dungeon]; Radiant Silver Bracers (4545, -0.12 DPS) [quest]; Enchanted Stonecloth Bracers (4979, -0.12 DPS) [quest] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | 34.1 healing_power points (2.28 DPS) | yes | Stormcloth Gloves (10011, +0.00 DPS, sim-verified) [crafted]; Earthen Silk Gloves (254017, -0.77 DPS) [crafted]; Truefaith Gloves (7049, -0.93 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 39.3 healing_power points (2.63 DPS) | yes | Deathmage Sash (10771, -0.36 DPS, sim-verified) [dungeon]; Sutarn's Ring (13105, -1.24 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -1.24 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 37.0 healing_power points (2.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.30 DPS, sim-verified) [crafted]; Aurora Pants (4044, -0.81 DPS) [world_drop]; Stormcloth Pants (10010, -0.89 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 34.4 healing_power points (2.30 DPS) | yes | Furen's Boots (13100, -0.09 DPS, sim-verified) [world_drop]; Boots of the Maharishi (9658, -1.03 DPS) [quest]; Thoughtcast Boots (10578, -1.10 DPS) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.20 DPS) | yes | Ogremind Ring (1993, -0.28 DPS) [world_drop]; Voodoo Band (1996, -0.28 DPS) [world]; Mindbender Loop (5009, -0.32 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 16.2 healing_power points (1.08 DPS) | yes | Ogremind Ring (1993, -0.16 DPS) [world_drop]; Mindbender Loop (5009, -0.19 DPS) [world_drop]; Voodoo Band (1996, -0.22 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Windweaver Staff (7757, -2.81 DPS) [dungeon]; Staff of Jordan (873, -2.85 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 24.3 healing_power points (1.63 DPS) | yes | Prophetic Cane (6803, +0.00 DPS, sim-verified) [quest]; Orb of Souls (249395, -0.52 DPS) [crafted]; Aurora Sphere (7610, -0.66 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 12.1 healing_power points (0.81 DPS) | yes | Flash Wand (5248, -0.23 DPS) [quest]; Goblin Igniter (5253, -0.29 DPS, sim-verified) [quest]; Captain Rackmore's Tiller (16789, -0.29 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Prodigious Shadowshard Pendant; shoulder: Earthen Silk Shoulders; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Gilded Handwraps; waist: Gilded Cord; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (troll, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 112.0. Weights run: 7.9s. Verify run: 3.9s. 403 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.087, intellect=3.463 ± 0.029, spirit=0.331 ± 0.015, mp5=1.134 ± 0.032, crit=0.223 ± 0.016 per rating point (14 rating = 1%, 3.120 per %), spell_haste=not significant (0.243 ± 0.160)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 94.5 healing_power points (5.76 DPS) | yes | Soulcatcher Halo (10630, -0.28 DPS) [dungeon]; Papal Fez (9431, -0.65 DPS) [dungeon]; Blood Guard's Satin Cover (220899, -0.90 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 49.8 healing_power points (3.04 DPS) | yes | Gemshard Heart (17707, -0.21 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.93 DPS) [quest]; Darkspear Warding Pendant (272073, -1.14 DPS) [vendor] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Rotgrip Mantle (17732, -0.23 DPS, sim-verified) [dungeon]; Nethergeld Shoulders (254049, -0.32 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 50.1 healing_power points (3.06 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.61 DPS) [dungeon]; Imperial Red Cloak (8248, -0.65 DPS) [world_drop]; Keeper's Cloak (14665, -0.95 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 91.5 healing_power points (5.58 DPS) | yes | Stone Guard's Satin Armor (220903, -0.73 DPS, sim-verified) [vendor]; Hibernal Robe (8113, -1.36 DPS) [world_drop]; Acumen Robes (17775, -1.36 DPS) [quest] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Aristocratic Cuffs (12546, -0.16 DPS, sim-verified) [dungeon]; Shizzle's Nozzle Wiper (11917, -0.20 DPS) [quest]; Forgotten Wraps (9433, -0.26 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 98.2 healing_power points (5.99 DPS) | yes | Virtuous Mitts (226950, -0.90 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.95 DPS) [crafted]; Greenleaf Handwraps (19116, -2.32 DPS) [quest] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Belt (13144, -0.08 DPS) [world_drop]; Dawnspire Cord (12466, -0.31 DPS, sim-verified) [dungeon]; Imperial Red Sash (8253, -0.50 DPS) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Stone Guard's Satin Leggings (220902, -0.48 DPS, sim-verified) [vendor]; Spellshock Leggings (9484, -0.65 DPS) [dungeon]; Venomshroud Leggings (14444, -0.75 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 62.2 healing_power points (3.79 DPS) | yes | Coldstone Slippers (18697, -0.56 DPS) [dungeon]; First Sergeant's Satin Boots (220900, -0.66 DPS, sim-verified) [vendor]; Gilded Slippers (254001, -1.01 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 53.6 healing_power points (3.27 DPS) | yes | Cyclopean Band (11824, -1.16 DPS) [dungeon]; Woodseed Hoop (17768, -1.37 DPS) [quest]; Snake Hoop (6750, -1.65 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 41.6 healing_power points (2.53 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Woodseed Hoop (17768, -0.63 DPS) [quest]; Snake Hoop (6750, -0.91 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.84 DPS) [quest]; Alchemists' Stone (13503, -0.96 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.04 DPS) [quest]; Alchemists' Stone (13503, -0.16 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -1.45 DPS) [quest]; Barman Shanker (12791, -1.49 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -2.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 21.8 healing_power points (1.33 DPS) | yes | Nature's Breath (19118, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.42 DPS) [quest]; Goblin Igniter (5253, -0.42 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Kentic Amice; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Nethergeld Cuffs; hands: Raider Handwraps; waist: Gilded Waistcord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 308.2. Weights run: 6.9s. Verify run: 3.2s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.360, intellect=1.877 ± 0.062, spirit=0.906 ± 0.071, mp5=1.181 ± 0.066, crit=0.293 ± 0.035 per rating point (14 rating = 1%, 4.107 per %), spell_haste=not significant (1.302 ± 1.075)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Virtuous Crown (226947, -1.42 DPS, sim-verified) [quest] |
| neck | Drake Tooth Necklace (21531) | The Nightmare Manifests [quest] | 48.5 healing_power points (6.74 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -0.52 DPS) [quest]; Lady Maye's Pendant (14558, -0.53 DPS) [world_drop]; The Eye of Zuldazar (19593, -0.68 DPS, sim-verified) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, -4.07 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 60.8 healing_power points (8.45 DPS) | yes | Cloak of the Cosmos (18389, -2.35 DPS, sim-verified) [dungeon]; Drape of Recovery (272413, -3.48 DPS) [vendor]; Darkspear Raider's Cloak (272063, -3.52 DPS) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 87.5 healing_power points (12.16 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -0.31 DPS, sim-verified) [dungeon]; Virtuous Robe (226945, -0.38 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 49.5 healing_power points (6.88 DPS) | yes | Virtuous Bracers (226949, -0.65 DPS) [quest]; General's Satin Bracers (17619, -0.67 DPS) [pvp]; Bracers of Mending (23129, -0.90 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 74.7 healing_power points (10.38 DPS) | yes | Desert Bloom Gloves (20717, -1.36 DPS) [quest]; Mooncloth Gloves (18409, -1.62 DPS) [crafted]; Hands of the Exalted Herald (12554, -2.40 DPS, sim-verified) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 65.4 healing_power points (9.09 DPS) | yes | General's Satin Cinch (17621, -1.21 DPS) [pvp]; Devout Belt (16696, -1.21 DPS) [dungeon]; Virtuous Belt (226948, -4.26 DPS, sim-verified) [quest] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 86.6 healing_power points (12.04 DPS) | yes | General's Satin Legguards (231634, -0.01 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.80 DPS) [pvp]; Virtuous Skirt (226946, -7.62 DPS, sim-verified) [quest] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 83.0 healing_power points (11.53 DPS) | yes | Virtuous Sandals (226952, +0.00 DPS, sim-verified) [quest]; General's Satin Walkers (231630, -2.80 DPS) [vendor]; Mooncloth Boots (15802, -2.91 DPS) [crafted] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.57 DPS) [quest]; Emerald Flame Ring (18395, -0.67 DPS) [dungeon]; Naglering (11669, -10.14 DPS, sim-verified) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Piety (22681, -0.12 DPS) [quest]; Emerald Flame Ring (18395, -0.22 DPS) [dungeon]; Naglering (11669, -7.32 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18469, -0.59 DPS) [quest]; Draconic Infused Emblem (22268, -1.59 DPS, sim-verified) [dungeon]; Briarwood Reed (12930, -1.81 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -0.27 DPS) [dungeon]; Death Speaker Scepter (2816, -1.23 DPS) [dungeon]; Hand of Edward the Odd (2243, -5.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.0 healing_power points (3.75 DPS) | yes | Oblivion's Touch (18761, -0.88 DPS) [dungeon]; Bonecreeper Stylus (13938, -1.18 DPS) [dungeon]; Sparkling Crystal Wand (20672, -1.94 DPS, sim-verified) [world] |

**New at 60:** head: Devout Crown; neck: Drake Tooth Necklace; shoulder: Devout Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket2: Serenity Field; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 514.2. Weights run: 4.3s. Verify run: 2.0s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.267, intellect=0.926 ± 0.027, spirit=0.885 ± 0.026, mp5=1.838 ± 0.029, crit=0.282 ± 0.021 per rating point (14 rating = 1%, 3.943 per %), spell_haste=1.769 ± 0.438

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 66.8 healing_power points (27.78 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -1.07 DPS) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 44.3 healing_power points (18.40 DPS) | yes | Animated Chain Necklace (18723, -2.47 DPS) [dungeon]; Drake Tooth Necklace (21531, -2.97 DPS) [quest]; The Eye of Zuldazar (19593, -3.05 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 78.9 healing_power points (32.79 DPS) | yes | Champion's Satin Mantle (227120, -10.13 DPS) [pvp]; Warlord's Satin Mantle (231631, -12.62 DPS) [pvp]; Mooncloth Shoulders (14139, -13.36 DPS) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 51.3 healing_power points (21.31 DPS) | yes | Drape of Recovery (272413, -1.02 DPS, sim-verified) [vendor]; Cloak of the Cosmos (18389, -6.27 DPS) [dungeon]; Battle Healer's Cloak (19526, -7.56 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 87.2 healing_power points (36.23 DPS) | yes | Robes of the Exalted (13346, -1.99 DPS) [dungeon]; Warlord's Satin Tunic (231632, -5.10 DPS) [vendor]; Virtuous Robe (226945, -7.53 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 38.8 healing_power points (16.13 DPS) | yes | Bracers of Mending (23129, -0.40 DPS) [dungeon]; Virtuous Bracers (226949, -1.14 DPS) [quest]; Nethergeld Cuffs (254061, -3.28 DPS) [crafted] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 60.1 healing_power points (24.96 DPS) | yes | Hands of the Exalted Herald (12554, -1.83 DPS) [dungeon]; Virtuous Mitts (226950, -4.13 DPS) [vendor]; Raider Handwraps (272097, -4.59 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 51.9 healing_power points (21.59 DPS) | yes | Whipvine Cord (18327, -0.65 DPS) [dungeon]; Virtuous Belt (226948, -2.33 DPS) [quest]; Penitent's Cinch (272394, -3.20 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 71.5 healing_power points (29.74 DPS) | yes | General's Satin Legguards (231634, -2.13 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -2.19 DPS) [pvp]; Virtuous Skirt (226946, -4.07 DPS) [quest] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 69.4 healing_power points (28.83 DPS) | yes | Virtuous Sandals (226952, -6.63 DPS) [quest]; Mooncloth Boots (15802, -8.68 DPS) [crafted]; Faith Healer's Boots (22247, -8.99 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (514.1 DPS) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Band of Piety (22681, -1.74 DPS) [quest]; Rosewine Circle (13178, -1.91 DPS) [dungeon] |
| finger2 | Fordring's Seal (16058) | In Dreams [quest] | sim-verified (514.1 DPS) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Band of Piety (22681, -1.25 DPS) [quest]; Rosewine Circle (13178, -1.43 DPS) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (514.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (514.1 DPS) | yes | Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor]; Alchemists' Stone (13503, -0.82 DPS, sim-verified) [crafted] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (514.1 DPS) | yes | The Lobotomizer (19324, -0.76 DPS, sim-verified) [rep]; Death Speaker Scepter (2816, -3.58 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -6.14 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 26.7 healing_power points (11.11 DPS) | yes | Bonecreeper Stylus (13938, -5.00 DPS) [dungeon]; Sparkling Crystal Wand (20672, -5.57 DPS) [world]; Ritssyn's Wand of Bad Mojo (22408, -6.54 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Desert Bloom Gloves; waist: Wisdom of the Timbermaw; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Band of Mending; finger2: Fordring's Seal; trinket1: Draconic Infused Emblem; trinket2: Second Wind; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

