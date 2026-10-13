# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 40.9. Weights run: 5.9s. Verify run: 17.0s. 151 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.334 ± 0.006, spirit=0.605 ± 0.003, mp5=1.556 ± 0.011, crit=0.070 ± 0.003 per rating point (14 rating = 1%, 0.982 per %), spell_haste=not significant (-0.005 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.63 DPS) | yes | Shadow Goggles (4373, -1.03 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.41 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 2.4 healing_power points (0.22 DPS) | yes | Scholarly Pendant (277203, -0.11 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.69 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 10.2 healing_power points (0.93 DPS) | yes | Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.57 DPS) [dungeon]; Sanguine Cape (14376, -0.67 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.5 healing_power points (1.95 DPS) | yes | Robe of the Moccasin (6465, -0.58 DPS) [dungeon]; Corsair's Overshirt (5202, -0.69 DPS, sim-verified) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.3 healing_power points (1.03 DPS) | yes | Repurposed Hair Band (281256, -0.68 DPS) [quest]; Mystic's Bracelets (14366, -0.79 DPS) [world_drop]; Bright Bracers (3647, -0.91 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Tomb Robber's Gloves (280096, -0.63 DPS) [quest]; Bright Gloves (3066, -0.66 DPS) [world_drop]; Magefist Gloves (12977, -0.94 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Keller's Girdle (2911, -0.51 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.58 DPS, sim-verified) [crafted]; Tarantula Silk Sash (3229, -0.77 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.58 DPS) | yes | Abomination Skin Leggings (23173, -1.61 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Darkweave Breeches (12987, -1.90 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Walking Boots (4660, -0.69 DPS) [world]; Sanguine Sandals (14374, -0.69 DPS) [world_drop]; Kimbra Boots (6191, -0.85 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.73 DPS) | yes | Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Band of Purification (12996, -0.40 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.44 DPS) [rep] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 6.3 healing_power points (0.57 DPS) | yes | Volcanic Rock Ring (12053, +0.00 DPS, sim-verified) [world_drop]; Band of Purification (12996, -0.24 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.29 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 13.3 healing_power points (1.21 DPS) | yes | Channeler's Staff (4437, -0.24 DPS, sim-verified) [world]; Staff of Westfall (2042, -0.28 DPS) [quest]; Lesser Staff of the Spire (1300, -0.48 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 2.7 healing_power points (0.24 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.08 DPS) [quest]; Elven Wand (5604, -0.12 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 151, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 70.6. Weights run: 6.0s. Verify run: 16.5s. 262 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.413 ± 0.007, spirit=0.841 ± 0.004, mp5=1.628 ± 0.015, crit=0.098 ± 0.005 per rating point (14 rating = 1%, 1.379 per %), spell_haste=0.351 ± 0.047

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.0 healing_power points (3.46 DPS) | yes | Embalmed Shroud (7691, -1.03 DPS, sim-verified) [dungeon]; Filigreed Pristine Circlet (253975, -1.46 DPS) [crafted]; Nightsky Cowl (4039, -1.54 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.0 healing_power points (0.82 DPS) | yes | Archmagister's Faceted Pendant (273049, +0.00 DPS, sim-verified) [world]; Scorn's Icy Choker (23169, -0.05 DPS) [dungeon]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 15.8 healing_power points (1.43 DPS) | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Batwing Mantle (6697, -0.02 DPS) [dungeon]; Nightsky Mantle (4718, -0.02 DPS) [world_drop] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 16.4 healing_power points (1.49 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Darkspear Raider's Cloak (272078, -0.23 DPS) [vendor]; Glowing Thresher Cape (6901, -0.24 DPS) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.5 healing_power points (3.23 DPS) | yes | Pristine Gown (253961, -0.15 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.08 DPS) [dungeon]; Robes of Arugal (6324, -1.14 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.9 healing_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.30 DPS, sim-verified) [world_drop]; Stonecloth Bindings (14416, -0.44 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.75 DPS) | yes | Gloves of Old (9395, -0.11 DPS, sim-verified) [world_drop]; Hotshot Pilot's Gloves (9491, -0.34 DPS) [dungeon]; Pristine Gloves (253913, -0.36 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.7 healing_power points (1.51 DPS) | yes | Novice Ardent's Sash (253887, -0.12 DPS, sim-verified) [crafted]; Resilient Cord (14406, -0.44 DPS) [world_drop]; Dreamer's Belt (4829, -0.46 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.1 healing_power points (3.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.35 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.08 DPS) [crafted]; Necromancer Leggings (2277, -1.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.3 healing_power points (3.02 DPS) | yes | Acidic Walkers (9454, -1.69 DPS) [dungeon]; Frothing Slippers (254003, -1.82 DPS) [crafted]; Nimbus Boots (6998, -1.92 DPS, sim-verified) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Runemender's Seal (273031, -0.56 DPS) [world]; Black Widow Band (6199, -0.74 DPS) [world]; The Queen's Jewel (13094, -0.77 DPS) [world_drop] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.9 healing_power points (1.08 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Black Widow Band (6199, -0.18 DPS) [world]; The Queen's Jewel (13094, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 84.4 healing_power points (7.68 DPS) | yes | Royal Diplomatic Scepter (9457, -0.33 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -1.50 DPS) [dungeon]; Spine of the Devourer (273040, -4.41 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.9 healing_power points (0.81 DPS) | yes | Summoner's Wand (5245, -0.30 DPS) [world]; Starfaller (13063, -0.30 DPS) [world_drop]; Necrotic Wand (7708, -0.70 DPS, sim-verified) [dungeon] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Caretaker's Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Electrocutioner Lagnut; main_hand: Wind Spirit Staff; ranged: Wand of Mana Concentration

No-known-source sample (15 of 262, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 85.8. Weights run: 6.2s. Verify run: 44.0s. 348 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.066, intellect=4.383 ± 0.038, spirit=0.671 ± 0.023, mp5=0.754 ± 0.022, crit=0.207 ± 0.015 per rating point (14 rating = 1%, 2.895 per %), spell_haste=7.521 ± 0.580

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 102.6 healing_power points (4.91 DPS) | yes | Miner's Hat of the Deep (9429, -1.02 DPS) [dungeon]; Thinking Cap (2624, -1.34 DPS) [world]; Corpseshroud (10574, -1.44 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 48.1 healing_power points (2.30 DPS) | yes | Necklace of Calisea (1714, -0.61 DPS) [dungeon]; Triune Amulet (7722, -0.61 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.94 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 72.0 healing_power points (3.45 DPS) | yes | Mistscape Mantle (4734, -0.98 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.04 DPS) [world_drop]; Windchaser Amice (14432, -1.65 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 50.9 healing_power points (2.44 DPS) | yes | Mantle of Lady Falther'ess (23178, +0.00 DPS, sim-verified) [dungeon]; Blackforge Cape (6424, -0.63 DPS) [dungeon]; Cloak of Rot (4462, -0.76 DPS) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Mageweave Vest (10007, +0.00 DPS) [crafted]; Silksand Wraps (14425, +0.00 DPS) [world_drop]; Silksand Tunic (14417, -1.61 DPS, sim-verified) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 39.4 healing_power points (1.89 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -0.21 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 63.6 healing_power points (3.05 DPS) | yes | Gilded Handwraps (254021, -0.52 DPS) [crafted]; Red Mageweave Gloves (10018, -0.95 DPS) [crafted]; Mistscape Gloves (6428, -1.16 DPS) [dungeon] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razzeric's Customized Seatbelt (6726, -0.41 DPS) [quest]; Teacher's Sash (10747, -0.41 DPS) [quest]; Deathmage Sash (10771, -0.79 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Crimson Silk Pantaloons (7062, +0.00 DPS) [crafted]; Pristine Leggings (253987, -0.00 DPS) [crafted]; Stormcloth Pants (10010, -3.10 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Kodo Rustler Boots (15697, -0.72 DPS) [quest]; Acidic Walkers (9454, -0.75 DPS) [dungeon]; Furen's Boots (13100, -4.62 DPS, sim-verified) [world_drop] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 32.7 healing_power points (1.57 DPS) | yes | Mindbender Loop (5009, -0.03 DPS) [world_drop]; Black Widow Band (6199, -0.10 DPS) [world]; Welken Ring (5011, -0.29 DPS) [world_drop] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 32.7 healing_power points (1.57 DPS) | yes | Mindbender Loop (5009, +0.00 DPS, sim-verified) [world_drop]; Black Widow Band (6199, -0.10 DPS) [world]; Welken Ring (5011, -0.29 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -0.84 DPS) [dungeon]; Hypnotic Blade (7714, -0.88 DPS) [dungeon] |
| off_hand | Aurora Sphere (7610) | Uldaman: Ancient Treasure [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Skull of Impending Doom (4984, -0.13 DPS) [quest]; Nightsky Orb (15929, -0.24 DPS) [world_drop]; Orb of Lorica (11262, -0.97 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 28.3 healing_power points (1.36 DPS) | yes | Goblin Igniter (5253, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.42 DPS) [quest]; Wand of Mana Concentration (273043, -0.44 DPS) [world] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Aurora Sphere; ranged: Jaina's Firestarter

No-known-source sample (15 of 348, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 105.8. Weights run: 6.4s. Verify run: 52.8s. 449 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=4.816 ± 0.060, spirit=0.208 ± 0.028, mp5=1.346 ± 0.060, crit=0.289 ± 0.021 per rating point (14 rating = 1%, 4.052 per %), spell_haste=-25.020 ± 1.049

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 130.6 healing_power points (6.08 DPS) | yes | Soulcatcher Halo (10630, +0.00 DPS, sim-verified) [dungeon]; Papal Fez (9431, -1.16 DPS) [dungeon]; Bad Mojo Mask (9470, -1.37 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 68.3 healing_power points (3.18 DPS) | yes | Glowing Eye of Mordresh (10769, -0.16 DPS, sim-verified) [dungeon]; Gemshard Heart (17707, -0.88 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.94 DPS) [quest] |
| shoulder | Knight-Lieutenant's Satin Pads (220894) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Rotgrip Mantle (17732, -0.93 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 68.5 healing_power points (3.19 DPS) | yes | Imperial Red Cloak (8248, -0.69 DPS, sim-verified) [world_drop]; Mantle of Lady Falther'ess (23178, -0.75 DPS) [dungeon]; Keeper's Cloak (14665, -0.95 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 123.5 healing_power points (5.75 DPS) | yes | Acumen Robes (17775, -1.27 DPS) [quest]; Knight's Satin Armor (220892, -1.52 DPS) [vendor]; Hibernal Robe (8113, -1.76 DPS, sim-verified) [world_drop] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Forgotten Wraps (9433, -0.03 DPS) [world_drop]; Nethergeld Cuffs (254061, -0.17 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 127.9 healing_power points (5.96 DPS) | yes | Virtuous Mitts (226950, -0.19 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -2.29 DPS) [crafted]; Stormcloth Gloves (10011, -2.76 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 103.5 healing_power points (4.82 DPS) | yes | Gilded Waistcord (254081, -1.45 DPS) [crafted]; Imperial Red Sash (8253, -1.46 DPS) [world_drop]; Serenity Belt (13144, -1.78 DPS, sim-verified) [world_drop] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kilt of the Atal'ai Prophet (10807, -0.39 DPS, sim-verified) [dungeon]; Venomshroud Leggings (14444, -0.75 DPS) [world_drop]; Imperial Red Pants (8251, -0.87 DPS) [world_drop] |
| feet | Sergeant Major's Satin Boots (220895) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 77.8 healing_power points (3.62 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Coldstone Slippers (18697, -0.23 DPS) [dungeon]; Highborne Footpads (14447, -0.93 DPS) [world_drop] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 73.3 healing_power points (3.41 DPS) | yes | Cyclopean Band (11824, -1.38 DPS) [dungeon]; Woodseed Hoop (17768, -1.39 DPS) [quest]; Mindseye Circle (10634, -3.95 DPS, sim-verified) [dungeon] |
| finger2 | Voodoo Band (1996) | Bloodscalp Witch Doctor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindseye Circle (10634, +0.00 DPS) [dungeon]; Cyclopean Band (11824, +0.00 DPS) [dungeon]; Woodseed Hoop (17768, +0.00 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.84 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.86 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.72 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.76 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.92 DPS, sim-verified) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Spellshifter Rod (9527, -0.32 DPS) [quest]; Glowing Brightwood Staff (812, -1.40 DPS) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (105.8 DPS) | yes | Enthralled Sphere (11625, +0.00 DPS) [dungeon]; Skullspell Orb (10708, -0.11 DPS) [quest]; Cloud Stone (17737, -0.46 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 29.5 healing_power points (1.37 DPS) | yes | Cairnstone Sliver (9654, -0.20 DPS) [quest]; Wand of Mana Concentration (273043, -0.35 DPS) [world]; Goblin Igniter (5253, -0.45 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Knight-Lieutenant's Satin Pads; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Shizzle's Nozzle Wiper; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Knight's Satin Leggings; feet: Sergeant Major's Satin Boots; finger1: Brainlash; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; off_hand: Orb of Lorica

No-known-source sample (15 of 449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 378.1. Weights run: 5.9s. Verify run: 50.1s. 1129 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.173, intellect=1.298 ± 0.050, spirit=0.941 ± 0.058, mp5=0.645 ± 0.064, crit=0.186 ± 0.025 per rating point (14 rating = 1%, 2.604 per %), spell_haste=-10.031 ± 0.944

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 72.7 healing_power points (12.46 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Devout Crown (16693, -16.83 DPS, sim-verified) [dungeon] |
| neck | The Eye of Zuldazar (19593) | The Eye of Zuldazar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Animated Chain Necklace (18723, -0.34 DPS) [dungeon]; Drake Tooth Necklace (21531, -2.35 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 72.3 healing_power points (12.39 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -2.15 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -2.80 DPS) [pvp]; Virtuous Mantle (226951, -12.10 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.0 healing_power points (9.42 DPS) | yes | Drape of Recovery (272413, -3.67 DPS) [vendor]; Caretaker's Cape (19530, -3.68 DPS) [rep]; Cloak of the Cosmos (18389, -4.61 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 88.1 healing_power points (15.09 DPS) | yes | Field Marshal's Satin Tunic (231624, -0.39 DPS) [vendor]; Robes of the Exalted (13346, -1.18 DPS, sim-verified) [dungeon]; Virtuous Robe (226945, -2.33 DPS) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, +0.00 DPS) [dungeon]; Marshal's Satin Bracers (17606, -1.02 DPS) [pvp]; Bracers of Hope (22667, -12.53 DPS, sim-verified) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Desert Bloom Gloves (20717, -13.22 DPS, sim-verified) [quest] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 57.6 healing_power points (9.88 DPS) | yes | Whipvine Cord (18327, -1.90 DPS) [dungeon]; Penitent's Cinch (272394, -2.07 DPS) [vendor]; Virtuous Belt (226948, -2.88 DPS, sim-verified) [quest] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [vendor]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Padre's Trousers (18386, -11.41 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 75.3 healing_power points (12.91 DPS) | yes | Mooncloth Boots (15802, -3.59 DPS) [crafted]; Faith Healer's Boots (22247, -3.85 DPS) [dungeon]; Virtuous Sandals (226952, -7.44 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.72 DPS) [dungeon]; Fordring's Seal (16058, -1.10 DPS) [quest]; Naglering (11669, -19.21 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon]; Emerald Flame Ring (18395, -0.07 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Mindtap Talisman (18371, -4.29 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (378.1 DPS) | yes | Royal Seal of Eldre'Thalas (18469, -1.10 DPS) [quest]; Briarwood Reed (12930, -2.23 DPS) [dungeon]; Second Wind (11819, -3.43 DPS) [dungeon] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -5.15 DPS) [dungeon]; Wind Spirit Staff (6689, -7.80 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 43.2 healing_power points (7.41 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Brightly Glowing Stone (18523, -1.07 DPS) [dungeon]; Thaurissan's Royal Scepter (11928, -3.41 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.3 healing_power points (4.69 DPS) | yes | Bonecreeper Stylus (13938, -1.91 DPS) [dungeon]; Oblivion's Touch (18761, -2.24 DPS) [dungeon]; Sparkling Crystal Wand (20672, -4.78 DPS, sim-verified) [world] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Virtuous Mitts; waist: Wisdom of the Timbermaw; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket2: Serenity Field; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 751.2. Weights run: 3.9s. Verify run: 18.2s. 1129 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.322, intellect=1.599 ± 0.050, spirit=1.187 ± 0.049, mp5=2.324 ± 0.056, crit=0.477 ± 0.039 per rating point (14 rating = 1%, 6.680 per %), spell_haste=not significant (-0.271 ± 1.106)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 85.9 healing_power points (22.45 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.12 DPS) [pvp]; Devout Crown (16693, -19.59 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.2 healing_power points (13.64 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -1.56 DPS) [quest]; Drake Tooth Necklace (21531, -1.83 DPS) [quest]; The Eye of Zuldazar (19593, -2.64 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 84.8 healing_power points (22.18 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -4.71 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -5.97 DPS) [pvp]; Virtuous Mantle (226951, -13.25 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 58.0 healing_power points (15.16 DPS) | yes | Drape of Recovery (272413, -4.63 DPS) [vendor]; Cloak of the Cosmos (18389, -5.73 DPS, sim-verified) [dungeon]; Caretaker's Cape (19530, -5.88 DPS) [rep] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Truefaith Vestments (14154, -7.44 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, +0.00 DPS) [dungeon]; Marshal's Satin Bracers (17606, -0.84 DPS) [pvp]; Bracers of Hope (22667, -12.99 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Desert Bloom Gloves (20717, -0.58 DPS) [quest]; Virtuous Mitts (226950, -1.10 DPS) [vendor] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Whipvine Cord (18327, -0.09 DPS) [dungeon]; Marshal's Satin Sash (17609, -1.33 DPS) [pvp]; Wisdom of the Timbermaw (19047, -11.49 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Padre's Trousers (18386, +0.00 DPS) [dungeon]; Knight-Captain's Satin Legguards (227125, +0.00 DPS) [vendor]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 83.0 healing_power points (21.70 DPS) | yes | Mooncloth Boots (15802, -5.54 DPS) [crafted]; Faith Healer's Boots (22247, -6.16 DPS) [dungeon]; Virtuous Sandals (226952, -6.62 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -3.06 DPS) [dungeon]; Fordring's Seal (16058, -4.19 DPS) [quest]; Naglering (11669, -22.67 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Fordring's Seal (16058, -0.60 DPS) [quest]; Emerald Flame Ring (18395, -0.94 DPS) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (751.2 DPS) | yes | Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, -1.87 DPS) [vendor]; Briarwood Reed (12930, -3.40 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (-2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -9.08 DPS) [dungeon]; Spellshifter Rod (9527, -12.41 DPS) [quest] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 51.6 healing_power points (13.50 DPS) | yes | Grand Marshal's Tome of Restoration (234590, -0.06 DPS) [pvp]; Thaurissan's Royal Scepter (11928, -2.55 DPS) [dungeon]; Tome of Divine Right (22319, -4.08 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.1 healing_power points (7.86 DPS) | yes | Oblivion's Touch (18761, -3.26 DPS) [dungeon]; Bonecreeper Stylus (13938, -3.31 DPS) [dungeon]; Sparkling Crystal Wand (20672, -5.04 DPS, sim-verified) [world] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Raider Handwraps; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Serenity Field; trinket2: Draconic Infused Emblem; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 38.8. Weights run: 5.9s. Verify run: 16.8s. 141 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.334 ± 0.006, spirit=0.605 ± 0.003, mp5=1.556 ± 0.011, crit=0.070 ± 0.003 per rating point (14 rating = 1%, 0.982 per %), spell_haste=not significant (-0.005 ± 0.014)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.63 DPS) | yes | Shadow Goggles (4373, -0.84 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.41 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.8 healing_power points (0.16 DPS) | yes | Scholarly Pendant (277203, -0.07 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 healing_power points (1.09 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.91 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 10.2 healing_power points (0.93 DPS) | yes | Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.57 DPS) [dungeon]; Sanguine Cape (14376, -0.66 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.5 healing_power points (1.95 DPS) | yes | Corsair's Overshirt (5202, -0.36 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.58 DPS) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.3 healing_power points (1.03 DPS) | yes | Bright Bracers (3647, -0.54 DPS) [world_drop]; Crystalline Cuffs (14148, -0.62 DPS) [dungeon]; Featherbead Bracers (15452, -1.02 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Blight Gloves (279877, -0.11 DPS, sim-verified) [quest]; Magefist Gloves (12977, -0.54 DPS) [world_drop]; Bright Gloves (3066, -0.66 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Novice Ardent's Sash (253887, -0.44 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.51 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.77 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.58 DPS) | yes | Abomination Skin Leggings (23173, -1.61 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Darkweave Breeches (12987, -1.78 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Spidersilk Boots (4320, -0.69 DPS) [crafted]; Walking Boots (4660, -0.69 DPS) [world]; Sanguine Sandals (14374, -1.02 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 8.0 healing_power points (0.73 DPS) | yes | Black Pearl Ring (6332, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Band of Purification (12996, -0.40 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.7 healing_power points (0.60 DPS) | yes | Black Pearl Ring (6332, +0.00 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop]; Band of Purification (12996, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 13.3 healing_power points (1.21 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.17 DPS) [quest]; Channeler's Staff (4437, -0.24 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.7 healing_power points (0.24 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.12 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (troll, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 68.4. Weights run: 6.0s. Verify run: 17.8s. 248 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.413 ± 0.007, spirit=0.841 ± 0.004, mp5=1.628 ± 0.015, crit=0.098 ± 0.005 per rating point (14 rating = 1%, 1.379 per %), spell_haste=0.351 ± 0.047

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.0 healing_power points (3.46 DPS) | yes | Embalmed Shroud (7691, -1.19 DPS, sim-verified) [dungeon]; Filigreed Pristine Circlet (253975, -1.46 DPS) [crafted]; Nightsky Cowl (4039, -1.54 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.0 healing_power points (0.82 DPS) | yes | Archmagister's Faceted Pendant (273049, +0.00 DPS, sim-verified) [world]; Scorn's Icy Choker (23169, -0.05 DPS) [dungeon]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 16.6 healing_power points (1.51 DPS) | yes | Death Speaker Mantle (6685, -0.09 DPS) [dungeon]; Batwing Mantle (6697, -0.09 DPS) [dungeon]; Nightsky Mantle (4718, -0.10 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 16.4 healing_power points (1.49 DPS) | yes | Darkspear Raider's Cloak (272078, -0.23 DPS) [vendor]; Glowing Thresher Cape (6901, -0.24 DPS) [dungeon]; Cloak of Rot (4462, -0.46 DPS) [world] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.5 healing_power points (3.23 DPS) | yes | Pristine Gown (253961, -0.09 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -1.08 DPS) [dungeon]; Robes of Arugal (6324, -1.14 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.9 healing_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.29 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.44 DPS) [quest] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.75 DPS) | yes | Gloves of Old (9395, +0.00 DPS, sim-verified) [world_drop]; Hotshot Pilot's Gloves (9491, -0.34 DPS) [dungeon]; Pristine Gloves (253913, -0.36 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | sim-verified (68.4 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Resilient Cord (14406, -0.31 DPS) [world_drop]; Pristine Sash (253925, -0.73 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.1 healing_power points (3.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.53 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.08 DPS) [crafted]; Necromancer Leggings (2277, -1.87 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.3 healing_power points (3.02 DPS) | yes | Frothing Slippers (254003, -1.82 DPS) [crafted]; Black Slippers (254007, -1.82 DPS) [crafted]; Acidic Walkers (9454, -1.94 DPS, sim-verified) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Runemender's Seal (273031, -0.56 DPS) [world]; Solid Gold Band (279845, -0.62 DPS) [quest]; Black Widow Band (6199, -0.74 DPS) [world] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.9 healing_power points (1.08 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, -0.06 DPS) [quest]; Black Widow Band (6199, -0.18 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 84.4 healing_power points (7.68 DPS) | yes | Royal Diplomatic Scepter (9457, -0.06 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -1.50 DPS) [dungeon]; Spine of the Devourer (273040, -4.41 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 8.9 healing_power points (0.81 DPS) | yes | Summoner's Wand (5245, -0.30 DPS) [world]; Starfaller (13063, -0.30 DPS) [world_drop]; Necrotic Wand (7708, -0.61 DPS, sim-verified) [dungeon] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Electrocutioner Lagnut; main_hand: Wind Spirit Staff; ranged: Wand of Mana Concentration

No-known-source sample (15 of 248, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 86.5. Weights run: 6.2s. Verify run: 45.3s. 330 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.066, intellect=4.383 ± 0.038, spirit=0.671 ± 0.023, mp5=0.754 ± 0.022, crit=0.207 ± 0.015 per rating point (14 rating = 1%, 2.895 per %), spell_haste=7.521 ± 0.580

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 102.6 healing_power points (4.91 DPS) | yes | Miner's Hat of the Deep (9429, -1.02 DPS) [dungeon]; Thinking Cap (2624, -1.34 DPS) [world]; Corpseshroud (10574, -4.13 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 48.1 healing_power points (2.30 DPS) | yes | Necklace of Calisea (1714, -0.61 DPS) [dungeon]; Triune Amulet (7722, -0.61 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.11 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 72.0 healing_power points (3.45 DPS) | yes | Mistscape Mantle (4734, -0.98 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.04 DPS) [world_drop]; Windchaser Amice (14432, -1.66 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 50.9 healing_power points (2.44 DPS) | yes | Blackforge Cape (6424, -0.63 DPS) [dungeon]; Cloak of Rot (4462, -0.76 DPS) [world]; Mantle of Lady Falther'ess (23178, -2.60 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silksand Tunic (14417, +0.00 DPS) [world_drop]; Silksand Wraps (14425, +0.00 DPS) [world_drop]; Red Mageweave Vest (10007, -15.51 DPS, sim-verified) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 39.4 healing_power points (1.89 DPS) | yes | Mistscape Bracers (4045, -0.20 DPS, sim-verified) [dungeon]; Radiant Silver Bracers (4545, -0.21 DPS) [quest]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 63.6 healing_power points (3.05 DPS) | yes | Red Mageweave Gloves (10018, -0.95 DPS) [crafted]; Mistscape Gloves (6428, -1.16 DPS) [dungeon]; Gilded Handwraps (254021, -7.41 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Razzeric's Customized Seatbelt (6726, -0.41 DPS) [quest]; Mistscape Sash (4736, -0.62 DPS) [dungeon]; Deathmage Sash (10771, -1.70 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Silk Pantaloons (7062, +0.00 DPS) [crafted]; Pristine Leggings (253987, +0.00 DPS) [crafted]; Stoneweaver Leggings (9407, -8.51 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 53.4 healing_power points (2.56 DPS) | yes | Furen's Boots (13100, -0.57 DPS) [world_drop]; Kodo Rustler Boots (15697, -0.72 DPS) [quest]; Boots of the Maharishi (9658, -3.01 DPS, sim-verified) [quest] |
| finger1 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 32.7 healing_power points (1.57 DPS) | yes | Mindbender Loop (5009, -0.03 DPS) [world_drop]; Black Widow Band (6199, -0.10 DPS) [world]; Welken Ring (5011, -0.29 DPS) [world_drop] |
| finger2 | Ogremind Ring (1993) | World drop [world_drop] | 32.7 healing_power points (1.57 DPS) | yes | Black Widow Band (6199, -0.10 DPS) [world]; Welken Ring (5011, -0.29 DPS) [world_drop]; Mindbender Loop (5009, -0.42 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+17.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -0.84 DPS) [dungeon]; Hypnotic Blade (7714, -0.88 DPS) [dungeon] |
| off_hand | Witch's Finger (16887) | Baron Aquanis [quest] | sim-verified (86.5 DPS) | yes | Prophetic Cane (6803, +0.00 DPS) [quest]; Aurora Sphere (7610, -0.08 DPS) [dungeon]; Skull of Impending Doom (4984, -0.21 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 28.3 healing_power points (1.36 DPS) | yes | Flash Wand (5248, -0.42 DPS) [quest]; Wand of Mana Concentration (273043, -0.44 DPS) [world]; Goblin Igniter (5253, -0.78 DPS, sim-verified) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Voodoo Band; finger2: Ogremind Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Witch's Finger; ranged: Jaina's Firestarter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (troll, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 98.4. Weights run: 6.4s. Verify run: 32.2s. 428 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=4.816 ± 0.060, spirit=0.208 ± 0.028, mp5=1.346 ± 0.060, crit=0.289 ± 0.021 per rating point (14 rating = 1%, 4.052 per %), spell_haste=-25.020 ± 1.049

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 130.6 healing_power points (6.08 DPS) | yes | Soulcatcher Halo (10630, -0.18 DPS, sim-verified) [dungeon]; Papal Fez (9431, -1.16 DPS) [dungeon]; Bad Mojo Mask (9470, -1.37 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 68.3 healing_power points (3.18 DPS) | yes | Glowing Eye of Mordresh (10769, -0.36 DPS, sim-verified) [dungeon]; Gemshard Heart (17707, -0.88 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.94 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 86.7 healing_power points (4.04 DPS) | yes | Kentic Amice (11624, +0.00 DPS, sim-verified) [dungeon]; Inquisitor's Shawl (19507, -0.42 DPS) [dungeon]; Blood Guard's Satin Pads (220901, -0.52 DPS) [vendor] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 68.5 healing_power points (3.19 DPS) | yes | Imperial Red Cloak (8248, -0.57 DPS, sim-verified) [world_drop]; Mantle of Lady Falther'ess (23178, -0.75 DPS) [dungeon]; Keeper's Cloak (14665, -0.95 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 123.5 healing_power points (5.75 DPS) | yes | Acumen Robes (17775, -1.27 DPS) [quest]; Stone Guard's Satin Armor (220903, -1.52 DPS) [vendor]; Hibernal Robe (8113, -1.77 DPS, sim-verified) [world_drop] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | sim-verified (98.4 DPS) | yes | Aristocratic Cuffs (12546, +0.00 DPS) [dungeon]; Forgotten Wraps (9433, -0.03 DPS) [world_drop]; Nethergeld Cuffs (254061, -0.17 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 127.9 healing_power points (5.96 DPS) | yes | Virtuous Mitts (226950, -1.42 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -2.29 DPS) [crafted]; Greenleaf Handwraps (19116, -2.47 DPS) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 103.5 healing_power points (4.82 DPS) | yes | Gilded Waistcord (254081, -1.45 DPS) [crafted]; Imperial Red Sash (8253, -1.46 DPS) [world_drop]; Serenity Belt (13144, -1.56 DPS, sim-verified) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 96.4 healing_power points (4.49 DPS) | yes | Stone Guard's Satin Leggings (220902, +0.00 DPS, sim-verified) [vendor]; Venomshroud Leggings (14444, -1.01 DPS) [world_drop]; Imperial Red Pants (8251, -1.13 DPS) [world_drop] |
| feet | First Sergeant's Satin Boots (220900) | PvP rank 9 · First Sergeant · Horde [vendor] | 77.8 healing_power points (3.62 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Coldstone Slippers (18697, -0.23 DPS) [dungeon]; Highborne Footpads (14447, -0.93 DPS) [world_drop] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 73.3 healing_power points (3.41 DPS) | yes | Cyclopean Band (11824, -1.38 DPS) [dungeon]; Woodseed Hoop (17768, -1.39 DPS) [quest]; Voodoo Band (1996, -1.81 DPS) [world] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 57.8 healing_power points (2.69 DPS) | yes | Cyclopean Band (11824, -0.66 DPS) [dungeon]; Woodseed Hoop (17768, -0.67 DPS) [quest]; Voodoo Band (1996, -1.09 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.84 DPS) [quest]; Alchemist's Stone (13503, -0.90 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.74 DPS) [quest]; Alchemist's Stone (13503, -0.80 DPS) [crafted]; Ankh of Life (1713, -1.36 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Spellshifter Rod (9527, -0.32 DPS) [quest]; Glowing Brightwood Staff (812, -1.40 DPS) [world_drop] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 68.5 healing_power points (3.19 DPS) | yes | Prophetic Cane (6803, -0.48 DPS, sim-verified) [quest]; Cloud Stone (17737, -0.85 DPS) [dungeon]; Mistscape Stave (7611, -1.17 DPS) [world_drop] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 29.5 healing_power points (1.37 DPS) | yes | Nature's Breath (19118, -0.24 DPS, sim-verified) [quest]; Wand of Mana Concentration (273043, -0.35 DPS) [world]; Goblin Igniter (5253, -0.45 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Shizzle's Nozzle Wiper; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: First Sergeant's Satin Boots; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; off_hand: Enthralled Sphere

No-known-source sample (15 of 428, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 369.4. Weights run: 5.9s. Verify run: 50.8s. 1119 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.173, intellect=1.298 ± 0.050, spirit=0.941 ± 0.058, mp5=0.645 ± 0.064, crit=0.186 ± 0.025 per rating point (14 rating = 1%, 2.604 per %), spell_haste=-10.031 ± 0.944

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 72.7 healing_power points (12.46 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Devout Crown (16693, -16.54 DPS, sim-verified) [dungeon] |
| neck | The Eye of Zuldazar (19593) | The Eye of Zuldazar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Drake Tooth Necklace (21531, +0.00 DPS) [quest]; Animated Chain Necklace (18723, -0.34 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 72.3 healing_power points (12.39 DPS) | yes | Champion's Satin Mantle (227120, -2.15 DPS) [pvp]; Warlord's Satin Mantle (231631, -2.80 DPS) [pvp]; Virtuous Mantle (226951, -10.12 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 55.0 healing_power points (9.42 DPS) | yes | Drape of Recovery (272413, -3.67 DPS) [vendor]; Battle Healer's Cloak (19526, -3.68 DPS) [rep]; Cloak of the Cosmos (18389, -4.60 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 88.1 healing_power points (15.09 DPS) | yes | Robes of the Exalted (13346, +0.00 DPS, sim-verified) [dungeon]; Warlord's Satin Tunic (231632, -0.39 DPS) [vendor]; Virtuous Robe (226945, -2.33 DPS) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, +0.00 DPS) [dungeon]; General's Satin Bracers (17619, -1.02 DPS) [pvp]; Bracers of Hope (22667, -11.35 DPS, sim-verified) [quest] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -0.18 DPS) [dungeon]; Raider Handwraps (272097, -0.54 DPS) [vendor]; Virtuous Mitts (226950, -1.20 DPS) [vendor] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Whipvine Cord (18327, -0.50 DPS) [dungeon]; Penitent's Cinch (272394, -0.66 DPS) [vendor]; Wisdom of the Timbermaw (19047, -9.50 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Padre's Trousers (18386, -8.17 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 75.3 healing_power points (12.91 DPS) | yes | Mooncloth Boots (15802, -3.59 DPS) [crafted]; Faith Healer's Boots (22247, -3.85 DPS) [dungeon]; Virtuous Sandals (226952, -6.80 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.72 DPS) [dungeon]; Fordring's Seal (16058, -1.10 DPS) [quest]; Naglering (11669, -17.82 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon]; Emerald Flame Ring (18395, -0.07 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Mindtap Talisman (18371, -2.93 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (369.4 DPS) | yes | Royal Seal of Eldre'Thalas (18469, -1.10 DPS) [quest]; Briarwood Reed (12930, -2.23 DPS) [dungeon]; Second Wind (11819, -3.43 DPS) [dungeon] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -5.15 DPS) [dungeon]; Wind Spirit Staff (6689, -7.80 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 43.2 healing_power points (7.41 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Brightly Glowing Stone (18523, -1.07 DPS) [dungeon]; Thaurissan's Royal Scepter (11928, -3.56 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.3 healing_power points (4.69 DPS) | yes | Bonecreeper Stylus (13938, -1.91 DPS) [dungeon]; Oblivion's Touch (18761, -2.24 DPS) [dungeon]; Sparkling Crystal Wand (20672, -3.98 DPS, sim-verified) [world] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Blessed Band of Light; trinket2: Serenity Field; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 1119, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 733.6. Weights run: 3.9s. Verify run: 19.5s. 1119 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.322, intellect=1.599 ± 0.050, spirit=1.187 ± 0.049, mp5=2.324 ± 0.056, crit=0.477 ± 0.039 per rating point (14 rating = 1%, 6.680 per %), spell_haste=not significant (-0.271 ± 1.106)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 85.9 healing_power points (22.45 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.12 DPS) [pvp]; Devout Crown (16693, -28.55 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.2 healing_power points (13.64 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -1.56 DPS) [quest]; Drake Tooth Necklace (21531, -1.83 DPS) [quest]; The Eye of Zuldazar (19593, -3.94 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 84.8 healing_power points (22.18 DPS) | yes | Champion's Satin Mantle (227120, -4.71 DPS) [pvp]; Warlord's Satin Mantle (231631, -5.97 DPS) [pvp]; Virtuous Mantle (226951, -19.67 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 58.0 healing_power points (15.16 DPS) | yes | Drape of Recovery (272413, -4.63 DPS) [vendor]; Battle Healer's Cloak (19526, -5.88 DPS) [rep]; Cloak of the Cosmos (18389, -8.77 DPS, sim-verified) [dungeon] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Truefaith Vestments (14154, -11.95 DPS, sim-verified) [crafted] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, +0.00 DPS) [dungeon]; General's Satin Bracers (17619, -0.84 DPS) [pvp]; Bracers of Hope (22667, -18.73 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Desert Bloom Gloves (20717, -0.58 DPS) [quest]; Virtuous Mitts (226950, -1.10 DPS) [vendor] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Whipvine Cord (18327, -0.09 DPS) [dungeon]; General's Satin Cinch (17621, -1.33 DPS) [pvp]; Wisdom of the Timbermaw (19047, -15.85 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Padre's Trousers (18386, +0.00 DPS) [dungeon]; Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 83.0 healing_power points (21.70 DPS) | yes | Mooncloth Boots (15802, -5.54 DPS) [crafted]; Faith Healer's Boots (22247, -6.16 DPS) [dungeon]; Virtuous Sandals (226952, -9.53 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -3.06 DPS) [dungeon]; Fordring's Seal (16058, -4.19 DPS) [quest]; Naglering (11669, -31.10 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; Fordring's Seal (16058, -0.60 DPS) [quest]; Emerald Flame Ring (18395, -0.94 DPS) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (733.6 DPS) | yes | Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Briarwood Reed (12930, -3.40 DPS) [dungeon]; Mindtap Talisman (18371, -4.30 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (-6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Briarwood Reed (12930, -1.53 DPS) [dungeon]; Mindtap Talisman (18371, -2.43 DPS) [dungeon] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -9.08 DPS) [dungeon]; Spellshifter Rod (9527, -12.41 DPS) [quest] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 51.6 healing_power points (13.50 DPS) | yes | High Warlord's Tome of Mending (234564, -0.06 DPS) [pvp]; Thaurissan's Royal Scepter (11928, -2.55 DPS) [dungeon]; Tome of Divine Right (22319, -6.74 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 30.1 healing_power points (7.86 DPS) | yes | Oblivion's Touch (18761, -3.26 DPS) [dungeon]; Bonecreeper Stylus (13938, -3.31 DPS) [dungeon]; Sparkling Crystal Wand (20672, -8.04 DPS, sim-verified) [world] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Virtuous Bracers; hands: Raider Handwraps; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Serenity Field; trinket2: Darkspear Voodoo Seal; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 1119, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

