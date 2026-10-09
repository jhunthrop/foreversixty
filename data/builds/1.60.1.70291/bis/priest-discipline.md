# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.9. Weights run: 10.0s. Verify run: 29.0s. 151 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=2.355 ± 0.009, spirit=1.374 ± 0.005, mp5=3.085 ± 0.033, crit=0.151 ± 0.007 per rating point (14 rating = 1%, 2.114 per %), spell_haste=not significant (0.008 ± 0.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.0 healing_power points (0.94 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.68 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.5 healing_power points (0.26 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.2 healing_power points (0.99 DPS) | yes | Slime-encrusted Pads (6461, -0.56 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.65 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 11.7 healing_power points (0.55 DPS) | yes | Seer's Cape (6378, -0.20 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.22 DPS) [crafted]; Sanguine Cape (14376, -0.35 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 28.9 healing_power points (1.35 DPS) | yes | Corsair's Overshirt (5202, -0.32 DPS) [dungeon]; Seer's Robe (2981, -0.50 DPS) [world_drop]; Robe of the Moccasin (6465, -0.79 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.0 healing_power points (0.98 DPS) | yes | Repurposed Hair Band (281256, -0.63 DPS) [quest]; Seer's Cuffs (3645, -0.74 DPS) [dungeon]; Bright Bracers (3647, -0.76 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 18.1 healing_power points (0.84 DPS) | yes | Bright Gloves (3066, -0.15 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.18 DPS) [quest]; Magefist Gloves (12977, -0.25 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 20.4 healing_power points (0.95 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Tarantula Silk Sash (3229, -0.28 DPS) [world]; Keller's Girdle (2911, -0.28 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 37.6 healing_power points (1.76 DPS) | yes | Filigreed Silky Leggings (253939, -0.84 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.84 DPS) [crafted]; Darkweave Breeches (12987, -1.08 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.1 healing_power points (0.75 DPS) | yes | Bluegill Sandals (1560, -0.27 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop]; Kimbra Boots (6191, -0.42 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 14.1 healing_power points (0.66 DPS) | yes | Band of Purification (12996, -0.27 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; Deep Fathom Ring (6463, -0.34 DPS) [dungeon] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 13.0 healing_power points (0.61 DPS) | yes | Band of Purification (12996, -0.25 DPS, sim-verified) [world_drop]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop]; Deep Fathom Ring (6463, -0.28 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 23.5 healing_power points (1.10 DPS) | yes | Channeler's Staff (4437, -0.22 DPS) [world]; Lesser Staff of the Spire (1300, -0.44 DPS) [world]; Staff of Westfall (2042, -0.48 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 4.7 healing_power points (0.22 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.03 DPS) [quest]; Dwarven Flamestick (5241, -0.09 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 151, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (human, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 80.6. Weights run: 10.2s. Verify run: 26.6s. 262 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.195 ± 0.012, spirit=1.739 ± 0.010, mp5=3.807 ± 0.011, crit=0.208 ± 0.009 per rating point (14 rating = 1%, 2.919 per %), spell_haste=not significant (0.302 ± 0.095)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.4 healing_power points (2.97 DPS) | yes | Nightsky Cowl (4039, -0.57 DPS) [world_drop]; Resilient Cap (14401, -0.84 DPS) [world_drop]; Embalmed Shroud (7691, -0.97 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.7 healing_power points (1.08 DPS) | yes | Scorn's Icy Choker (23169, -0.18 DPS) [dungeon]; Crystal Starfire Medallion (5003, -0.29 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.33 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 32.8 healing_power points (2.24 DPS) | yes | Nightsky Mantle (4718, -0.45 DPS) [world_drop]; Mantle of Honor (3560, -0.58 DPS, sim-verified) [quest]; Death Speaker Mantle (6685, -0.59 DPS) [dungeon] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 27.0 healing_power points (1.84 DPS) | yes | Prelacy Cape (7004, -0.38 DPS) [quest]; Glowing Thresher Cape (6901, -0.41 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.46 DPS, sim-verified) [vendor] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 44.1 healing_power points (3.02 DPS) | yes | Pristine Gown (253961, -0.42 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.87 DPS) [dungeon]; Silver-thread Robe (4035, -0.89 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.4 healing_power points (1.53 DPS) | yes | Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Nightsky Wristbands (6407, -0.35 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.75 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.8 healing_power points (2.18 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.20 DPS, sim-verified) [dungeon]; Town Clerk's Mittens (270029, -0.53 DPS) [quest]; Truefaith Gloves (7049, -0.70 DPS) [crafted] |
| waist | Resilient Cord (14406) | World drop [world_drop] | 20.1 healing_power points (1.38 DPS) | yes | Pristine Sash (253925, +0.00 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.09 DPS) [vendor]; Novice Ardent's Sash (253887, -0.16 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.1 healing_power points (3.15 DPS) | yes | Filigreed Pristine Leggings (253937, -0.78 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.82 DPS) [crafted]; Darkweave Breeches (12987, -1.39 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.3 healing_power points (2.89 DPS) | yes | Glinteye Slippers (273024, -1.28 DPS) [dungeon]; Soggy Boots (274747, -1.34 DPS) [vendor]; Acidic Walkers (9454, -1.63 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.5 healing_power points (1.88 DPS) | yes | Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.65 DPS) [vendor]; Black Widow Band (6199, -0.83 DPS) [world] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 18.3 healing_power points (1.25 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS, sim-verified) [dungeon]; Sea Giant's Toe Ring (274746, -0.02 DPS) [vendor]; Black Widow Band (6199, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 109.8 healing_power points (7.51 DPS) | yes | Royal Diplomatic Scepter (9457, -1.66 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -2.86 DPS) [dungeon]; Hardened Root Staff (1317, -3.91 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Summoner's Wand (5245) (or Starfaller (13063), Lesser Mystic Wand (11289)) | Dalaran Summoner [world] | 8.8 healing_power points (0.60 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Starfaller (13063, +0.00 DPS) [world_drop]; Consecrated Wand (5244, -0.06 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Repairman's Cape; chest: Death Speaker Robes; hands: Gloves of Old; waist: Resilient Cord; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; main_hand: Wind Spirit Staff; ranged: Summoner's Wand

No-known-source sample (15 of 262, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 139.0. Weights run: 11.9s. Verify run: 81.6s. 343 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=2.488 ± 0.018, spirit=not significant (0.119 ± 0.030), mp5=2.616 ± 0.016, crit=0.435 ± 0.013 per rating point (14 rating = 1%, 6.089 per %), spell_haste=-7.906 ± 0.294

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 65.4 healing_power points (4.92 DPS) | yes | Miner's Hat of the Deep (9429, -1.65 DPS) [dungeon]; Thinking Cap (2624, -1.73 DPS) [world]; Corpseshroud (10574, -2.49 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.9 healing_power points (2.47 DPS) | yes | Necklace of Calisea (1714, -1.10 DPS) [dungeon]; Triune Amulet (7722, -1.10 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.39 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 47.3 healing_power points (3.56 DPS) | yes | Mistscape Mantle (4734, -1.46 DPS) [dungeon]; Batwing Mantle (6697, -1.46 DPS) [dungeon]; Windchaser Amice (14432, -1.51 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (139.0 DPS) | yes | Mantle of Lady Falther'ess (23178, +0.00 DPS) [dungeon]; Blackforge Cape (6424, -0.56 DPS) [dungeon]; Cloak of Rot (4462, -0.60 DPS) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 58.9 healing_power points (4.43 DPS) | yes | Red Mageweave Vest (10007, -1.07 DPS) [crafted]; Silksand Tunic (14417, -1.40 DPS) [world_drop]; Death Speaker Robes (6682, -5.85 DPS, sim-verified) [dungeon] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 22.4 healing_power points (1.68 DPS) | yes | Mistscape Bracers (4045, -0.19 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.19 DPS) [quest]; Mindthrust Bracers (1974, -0.21 DPS, sim-verified) [dungeon] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 40.9 healing_power points (3.07 DPS) | yes | Town Clerk's Mittens (270029, -1.01 DPS) [quest]; Red Mageweave Gloves (10018, -1.20 DPS) [crafted]; Gilded Handwraps (254021, -2.21 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 42.6 healing_power points (3.20 DPS) | yes | Razzeric's Customized Seatbelt (6726, -0.96 DPS) [quest]; Teacher's Sash (10747, -0.96 DPS) [quest]; Deathmage Sash (10771, -2.36 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (139.0 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, -2.10 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.9 healing_power points (2.85 DPS) | yes | Acidic Walkers (9454, -1.32 DPS) [dungeon]; Furen's Boots (13100, -1.40 DPS) [world_drop]; Kodo Rustler Boots (15697, -2.10 DPS, sim-verified) [quest] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 18.2 healing_power points (1.37 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Voodoo Band (1996, -0.04 DPS) [world]; Mindbender Loop (5009, -0.04 DPS) [world_drop] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.35 DPS) | yes | Ogremind Ring (1993, -0.02 DPS) [world_drop]; Mindbender Loop (5009, -0.03 DPS) [world_drop]; Voodoo Band (1996, -0.12 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (139.0 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (139.0 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -0.72 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -1.44 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 36.9 healing_power points (2.77 DPS) | yes | Aurora Sphere (7610, -1.43 DPS) [dungeon]; Skull of Impending Doom (4984, -1.46 DPS) [quest]; Arcane Infused Rod (279838, -2.06 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.3 healing_power points (1.15 DPS) | yes | Flash Wand (5248, -0.37 DPS) [quest]; Summoner's Wand (5245, -0.40 DPS) [world]; Goblin Igniter (5253, -0.74 DPS, sim-verified) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 343, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 197.6. Weights run: 11.8s. Verify run: 51.8s. 441 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.502 ± 0.033, spirit=1.680 ± 0.045, mp5=4.037 ± 0.076, crit=0.602 ± 0.017 per rating point (14 rating = 1%, 8.435 per %), spell_haste=-6.644 ± 0.420

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 81.8 healing_power points (6.63 DPS) | yes | Knight-Lieutenant's Satin Cover (220896, -0.01 DPS) [vendor]; Papal Fez (9431, -0.17 DPS) [dungeon]; Soulcatcher Halo (10630, -0.20 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 41.7 healing_power points (3.38 DPS) | yes | Darkmoon Necklace (19303, -0.20 DPS) [vendor]; Gemshard Heart (17707, -0.54 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.71 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Satin Pads (220894) | Captain Dirgehammer [vendor] | 63.2 healing_power points (5.12 DPS) | yes | Kentic Amice (11624, -0.53 DPS) [dungeon]; Nethergeld Shoulders (254049, -0.57 DPS) [crafted]; Inquisitor's Shawl (19507, -1.27 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 43.4 healing_power points (3.52 DPS) | yes | Featherskin Cape (10843, -0.67 DPS) [world]; Imperial Red Cloak (8248, -0.74 DPS) [world_drop]; Caretaker's Cape (19531, -0.92 DPS) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 92.9 healing_power points (7.53 DPS) | yes | Robes of Insight (940, -0.42 DPS) [world_drop]; Vestments of the Atal'ai Prophet (10806, -1.63 DPS) [dungeon]; Stormcloth Vest (10020, -1.74 DPS) [crafted] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | sim-verified (197.6 DPS) | yes | Nethergeld Cuffs (254061, -0.14 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -1.02 DPS) [quest]; Forgotten Wraps (9433, -1.43 DPS) [world_drop] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 81.8 healing_power points (6.63 DPS) | yes | Raider Handwraps (272098, -0.38 DPS) [vendor]; Gilded Gloves (254095, -1.16 DPS) [crafted]; Virtuous Hands (226958, -2.36 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 59.5 healing_power points (4.83 DPS) | yes | Gilded Cord (254037, -0.60 DPS) [crafted]; Gilded Waistcord (254081, -0.65 DPS) [crafted]; Earthenweave Cord (254077, -1.22 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 81.3 healing_power points (6.59 DPS) | yes | Senior Designer's Pantaloons (11841, -0.69 DPS) [dungeon]; Knight's Satin Leggings (220893, -0.82 DPS) [vendor]; Dalewind Trousers (13008, -1.36 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 53.5 healing_power points (4.34 DPS) | yes | Vinerot Sandals (17748, -0.18 DPS) [dungeon]; Coldstone Slippers (18697, -0.19 DPS) [dungeon]; Mistwalker Boots (10629, -0.26 DPS) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 45.9 healing_power points (3.72 DPS) | yes | Cyclopean Band (11824, -1.03 DPS) [dungeon]; Mindseye Circle (10634, -1.29 DPS) [dungeon]; Snake Hoop (6750, -1.35 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 35.2 healing_power points (2.85 DPS) | yes | Cyclopean Band (11824, -0.16 DPS) [dungeon]; Mindseye Circle (10634, -0.42 DPS) [dungeon]; Snake Hoop (6750, -0.48 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -3.27 DPS) [world_drop]; Uther's Strength (11302, -3.45 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -3.82 DPS) [quest] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.27 DPS) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -3.25 DPS) [quest]; Wind Spirit Staff (6689, -3.34 DPS) [dungeon]; Hand of Righteousness (7721, -3.58 DPS) [dungeon] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (197.6 DPS) | yes | Cloud Stone (17737, -0.13 DPS) [dungeon]; Orb of Lorica (11262, -0.52 DPS) [quest]; Twisting Essence Jar (249456, -0.89 DPS) [crafted] |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 20.9 healing_power points (1.70 DPS) | yes | Jaina's Firestarter (13064, -0.07 DPS) [world_drop]; Flash Wand (5248, -0.48 DPS) [quest]; Goblin Igniter (5253, -0.48 DPS) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Knight-Lieutenant's Satin Pads; back: Darkspear Raider's Cloak; chest: Embrace of the Wind Serpent; wrist: Aristocratic Cuffs; hands: Virtuous Mitts; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Brainlash; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Evonice's Landin' Pilla; main_hand: Charstone Dirk; off_hand: Enthralled Sphere; ranged: Cairnstone Sliver

No-known-source sample (15 of 441, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 313.8. Weights run: 11.6s. Verify run: 68.2s. 1121 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.101, intellect=2.853 ± 0.047, spirit=5.865 ± 0.121, mp5=2.109 ± 0.143, crit=0.762 ± 0.031 per rating point (14 rating = 1%, 10.669 per %), spell_haste=-16.366 ± 0.818

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 186.0 healing_power points (15.88 DPS) | yes | Devout Crown (16693, -0.82 DPS) [dungeon]; Crown of Caer Darrow (13986, -0.99 DPS) [quest]; Field Marshal's Satin Hood (231622, -1.05 DPS) [vendor] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 117.1 healing_power points (9.99 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Lady Maye's Pendant (14558, -0.36 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.60 DPS) [world_drop] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Devout Mantle (16695, -0.49 DPS) [dungeon] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 104.6 healing_power points (8.93 DPS) | yes | Featherskin Cape (10843, -0.45 DPS) [world]; Butcher's Apron (12608, -0.92 DPS) [dungeon]; Shroud of the Exile (15421, -1.77 DPS) [quest] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Alanna's Embrace (13314, -0.01 DPS) [dungeon]; Vestments of the Atal'ai Prophet (10806, -0.41 DPS) [dungeon] |
| wrist | Bracers of Mending (23129) | Dire Maul: Revanchion [dungeon] | sim-verified (313.8 DPS) | yes | Bracers of Hope (22667, -0.01 DPS) [quest]; Wyrmthalak's Shackles (13958, -0.54 DPS) [quest]; Marshal's Satin Bracers (17606, -0.60 DPS) [pvp] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 161.0 healing_power points (13.74 DPS) | yes | Devout Gloves (16692, -1.51 DPS) [dungeon]; Hands of the Exalted Herald (12554, -1.75 DPS) [dungeon]; Swarmtender's Gloves (275607, -2.36 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Penitent's Cinch (272394, +0.00 DPS) [vendor]; Wisdom of the Timbermaw (19047, -0.68 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 201.7 healing_power points (17.22 DPS) | yes | Haunting Specter Leggings (11929, -0.28 DPS) [dungeon]; Virtuous Skirt (226946, -1.66 DPS) [quest]; The Postmaster's Trousers (13389, -2.33 DPS) [dungeon] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | 178.6 healing_power points (15.25 DPS) | yes | Incandescent Mooncloth Boots (227862, -1.07 DPS) [vendor]; Mooncloth Boots (15802, -3.28 DPS) [crafted]; Devout Sandals (16691, -3.36 DPS) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of the Hierophant (13096, -1.30 DPS) [world_drop]; Seal of Rivendare (13345, -1.60 DPS) [dungeon]; Eye of Adaegus (5266, -1.77 DPS) [world_drop] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (313.8 DPS) | yes | Band of the Hierophant (13096, -0.27 DPS) [world_drop]; Seal of Rivendare (13345, -0.56 DPS) [dungeon]; Eye of Adaegus (5266, -0.74 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -0.34 DPS) [world_drop]; Soulkeeper (1607, -1.75 DPS) [world_drop]; Wind Spirit Staff (6689, -2.15 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 81.5 healing_power points (6.96 DPS) | yes | Cairnstone Sliver (9654, -3.24 DPS) [quest]; Jaina's Firestarter (13064, -4.00 DPS) [world_drop]; Oblivion's Touch (18761, -4.28 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Virtuous Mantle; back: Frostweaver Cape; chest: Mooncloth Vest; wrist: Bracers of Mending; waist: Virtuous Belt; legs: Devout Skirt; feet: Virtuous Sandals; finger1: The Postmaster's Seal; finger2: Emerald Flame Ring; trinket2: Royal Seal of Eldre'Thalas; main_hand: Dancing Sliver; ranged: Torch of Light

No-known-source sample (15 of 1121, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 581.0. Weights run: 7.6s. Verify run: 37.6s. 1121 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.251, intellect=3.031 ± 0.065, spirit=2.557 ± 0.062, mp5=3.741 ± 0.067, crit=0.952 ± 0.047 per rating point (14 rating = 1%, 13.329 per %), spell_haste=not significant (-1.927 ± 1.308)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 138.8 healing_power points (16.27 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.99 DPS) [pvp]; Devout Crown (16693, -13.03 DPS, sim-verified) [dungeon] |
| neck | The Eye of Zuldazar (19593) | The Eye of Zuldazar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 107.0 healing_power points (12.55 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -0.45 DPS) [pvp]; Virtuous Mantle (226951, -1.75 DPS) [quest]; Devout Mantle (16695, -12.50 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 72.3 healing_power points (8.48 DPS) | yes | Darkspear Raider's Cloak (272063, -0.99 DPS) [vendor]; Shroud of the Exile (15421, -1.05 DPS) [quest]; Frostweaver Cape (12968, -17.51 DPS, sim-verified) [dungeon] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | 137.2 healing_power points (16.08 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Alanna's Embrace (13314, -0.64 DPS) [dungeon]; Mooncloth Vest (14138, -11.11 DPS, sim-verified) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 82.0 healing_power points (9.62 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Bracers of Mending (23129, -0.99 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -1.01 DPS) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 105.8 healing_power points (12.41 DPS) | yes | Hands of the Exalted Herald (12554, -0.32 DPS) [dungeon]; Virtuous Mitts (226950, -0.79 DPS, sim-verified) [vendor]; Devout Gloves (16692, -2.00 DPS) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | 97.6 healing_power points (11.44 DPS) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Devout Belt (16696, -0.34 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -8.91 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-verified (+9.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Padre's Trousers (18386, -0.05 DPS) [dungeon]; Devout Skirt (16694, -9.92 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 122.2 healing_power points (14.33 DPS) | yes | Virtuous Sandals (226952, +0.00 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -2.64 DPS) [crafted]; Faith Healer's Boots (22247, -3.42 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, -1.07 DPS) [dungeon]; Seal of Rivendare (13345, -1.36 DPS) [dungeon]; Naglering (11669, -17.64 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, +0.00 DPS) [dungeon]; Emerald Flame Ring (18395, +0.00 DPS) [dungeon]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -1.75 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18469, -2.29 DPS, sim-verified) [quest]; Ankh of Life (1713, -2.98 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Mindtap Talisman (18371, -0.10 DPS) [dungeon]; Ankh of Life (1713, -1.33 DPS) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, -2.91 DPS) [world_drop]; Wind Spirit Staff (6689, -3.49 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | sim-verified (581.0 DPS) | yes | Thaurissan's Royal Scepter (11928, +0.00 DPS) [dungeon]; Book of the Dead (13353, +0.00 DPS) [dungeon]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 45.1 healing_power points (5.29 DPS) | yes | Sparkling Crystal Wand (20672, -1.51 DPS) [world]; Cairnstone Sliver (9654, -2.02 DPS) [quest]; Oblivion's Touch (18761, -5.99 DPS, sim-verified) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 1121, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 025003010000000000-00000000000000000-000000000000000000)

Set DPS (verified): 35.5. Weights run: 10.0s. Verify run: 28.7s. 141 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=2.355 ± 0.009, spirit=1.374 ± 0.005, mp5=3.085 ± 0.033, crit=0.151 ± 0.007 per rating point (14 rating = 1%, 2.114 per %), spell_haste=not significant (0.008 ± 0.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.0 healing_power points (0.94 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.68 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.1 healing_power points (0.19 DPS) | yes | Scholarly Pendant (277203, -0.08 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 21.2 healing_power points (0.99 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.53 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.56 DPS) [dungeon] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 11.7 healing_power points (0.55 DPS) | yes | Sanguine Cape (14376, -0.17 DPS, sim-verified) [world_drop]; Seer's Cape (6378, -0.20 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.22 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 28.9 healing_power points (1.35 DPS) | yes | Corsair's Overshirt (5202, -0.32 DPS) [dungeon]; Seer's Robe (2981, -0.50 DPS) [world_drop]; Robe of the Moccasin (6465, -0.60 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.0 healing_power points (0.98 DPS) | yes | Bright Bracers (3647, -0.54 DPS) [world_drop]; Crystalline Cuffs (14148, -0.57 DPS) [dungeon]; Featherbead Bracers (15452, -0.69 DPS, sim-verified) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 23.4 healing_power points (1.09 DPS) | yes | Pristine Gloves (253913, -0.25 DPS) [crafted]; Magefist Gloves (12977, -0.28 DPS) [world_drop]; Bright Gloves (3066, -0.39 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 20.4 healing_power points (0.95 DPS) | yes | Novice Ardent's Sash (253887, -0.15 DPS) [crafted]; Keller's Girdle (2911, -0.20 DPS, sim-verified) [world_drop]; Tarantula Silk Sash (3229, -0.28 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 37.6 healing_power points (1.76 DPS) | yes | Filigreed Silky Leggings (253939, -0.84 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.84 DPS) [crafted]; Darkweave Breeches (12987, -1.01 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 16.1 healing_power points (0.75 DPS) | yes | Walking Boots (4660, -0.31 DPS) [world]; Sanguine Sandals (14374, -0.31 DPS) [world_drop]; Bluegill Sandals (1560, -0.70 DPS, sim-verified) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 14.1 healing_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Band of Purification (12996, -0.27 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | 13.0 healing_power points (0.61 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Band of Purification (12996, -0.22 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 23.5 healing_power points (1.10 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.05 DPS) [quest]; Channeler's Staff (4437, -0.22 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 4.7 healing_power points (0.22 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.11 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 141, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (undead, 025003031304000000-00000000000000000-000000000000000000)

Set DPS (verified): 79.7. Weights run: 10.2s. Verify run: 26.7s. 249 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.195 ± 0.012, spirit=1.739 ± 0.010, mp5=3.807 ± 0.011, crit=0.208 ± 0.009 per rating point (14 rating = 1%, 2.919 per %), spell_haste=not significant (0.302 ± 0.095)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 43.4 healing_power points (2.97 DPS) | yes | Nightsky Cowl (4039, -0.57 DPS) [world_drop]; Embalmed Shroud (7691, -0.81 DPS, sim-verified) [dungeon]; Resilient Cap (14401, -0.84 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.7 healing_power points (1.08 DPS) | yes | Crystal Starfire Medallion (5003, -0.10 DPS, sim-verified) [world_drop]; Scorn's Icy Choker (23169, -0.18 DPS) [dungeon]; Darkspear Warding Pendant (272075, -0.33 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 32.8 healing_power points (2.24 DPS) | yes | Mantle of Woe (7750, -0.21 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.45 DPS) [world_drop]; Ghostly Mantle (3324, -0.56 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 22.8 healing_power points (1.56 DPS) | yes | Glowing Thresher Cape (6901, -0.13 DPS) [dungeon]; Battle Healer's Cloak (19529, -0.19 DPS) [rep]; Cloak of Rot (4462, -0.36 DPS) [world] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 44.1 healing_power points (3.02 DPS) | yes | Pristine Gown (253961, -0.12 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.87 DPS) [dungeon]; Silver-thread Robe (4035, -0.89 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.4 healing_power points (1.53 DPS) | yes | Nightsky Wristbands (6407, -0.13 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.33 DPS) [world_drop]; Spidertank Oilrag (9448, -0.75 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 31.8 healing_power points (2.18 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.38 DPS) [dungeon]; Blight Gloves (279877, -0.53 DPS) [quest]; Truefaith Gloves (7049, -0.70 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 25.0 healing_power points (1.71 DPS) | yes | Resilient Cord (14406, -0.27 DPS, sim-verified) [world_drop]; Pristine Sash (253925, -0.35 DPS) [crafted]; Dreamer's Belt (4829, -0.42 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 46.1 healing_power points (3.15 DPS) | yes | Filigreed Pristine Leggings (253937, -0.54 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.82 DPS) [crafted]; Sacred Burial Trousers (6282, -1.33 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.3 healing_power points (2.89 DPS) | yes | Glinteye Slippers (273024, -1.28 DPS) [dungeon]; Soggy Boots (274747, -1.34 DPS) [vendor]; Acidic Walkers (9454, -1.95 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 27.5 healing_power points (1.88 DPS) | yes | Electrocutioner Lagnut (9447, -0.64 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.65 DPS) [vendor]; Black Widow Band (6199, -0.83 DPS) [world] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 18.3 healing_power points (1.25 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS, sim-verified) [dungeon]; Sea Giant's Toe Ring (274746, -0.02 DPS) [vendor]; Black Widow Band (6199, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | 109.8 healing_power points (7.51 DPS) | yes | Royal Diplomatic Scepter (9457, -1.39 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -2.86 DPS) [dungeon]; Advisor's Gnarled Staff (19569, -5.41 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Summoner's Wand (5245) (or Starfaller (13063), Lesser Mystic Wand (11289)) | Dalaran Summoner [world] | 8.8 healing_power points (0.60 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Starfaller (13063, +0.00 DPS) [world_drop]; Necrotic Wand (7708, -0.19 DPS) [dungeon] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Death Speaker Robes; hands: Gloves of Old; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; main_hand: Wind Spirit Staff; ranged: Summoner's Wand

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 025003031305101520-00000000000000000-000000000000000000)

Set DPS (verified): 136.6. Weights run: 11.9s. Verify run: 88.2s. 326 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=2.488 ± 0.018, spirit=not significant (0.119 ± 0.030), mp5=2.616 ± 0.016, crit=0.435 ± 0.013 per rating point (14 rating = 1%, 6.089 per %), spell_haste=-7.906 ± 0.294

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 65.4 healing_power points (4.92 DPS) | yes | Corpseshroud (10574, -1.31 DPS) [dungeon]; Miner's Hat of the Deep (9429, -1.65 DPS) [dungeon]; Thinking Cap (2624, -1.73 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.9 healing_power points (2.47 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.60 DPS) [quest]; Necklace of Calisea (1714, -1.10 DPS) [dungeon]; Triune Amulet (7722, -1.10 DPS) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 47.3 healing_power points (3.56 DPS) | yes | Windchaser Amice (14432, -1.13 DPS) [world_drop]; Mantle of Woe (7750, -1.32 DPS) [quest]; Mistscape Mantle (4734, -1.46 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (136.6 DPS) | yes | Mantle of Lady Falther'ess (23178, +0.00 DPS) [dungeon]; Blackforge Cape (6424, -0.56 DPS) [dungeon]; Cloak of Rot (4462, -0.60 DPS) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 58.9 healing_power points (4.43 DPS) | yes | Death Speaker Robes (6682, -0.87 DPS) [dungeon]; Red Mageweave Vest (10007, -1.07 DPS) [crafted]; Silksand Tunic (14417, -1.40 DPS) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 22.4 healing_power points (1.68 DPS) | yes | Mindthrust Bracers (1974, -0.16 DPS) [dungeon]; Mistscape Bracers (4045, -0.19 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.19 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 40.9 healing_power points (3.07 DPS) | yes | Gilded Handwraps (254021, -0.11 DPS) [crafted]; Red Mageweave Gloves (10018, -1.20 DPS) [crafted]; Truefaith Gloves (7049, -1.38 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 42.6 healing_power points (3.20 DPS) | yes | Deathmage Sash (10771, -0.40 DPS) [dungeon]; Razzeric's Customized Seatbelt (6726, -0.96 DPS) [quest]; Mistscape Sash (4736, -1.15 DPS) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Filigreed Pristine Leggings (253937, +0.00 DPS) [crafted]; Pristine Leggings (253987, +0.00 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 37.9 healing_power points (2.85 DPS) | yes | Boots of the Maharishi (9658, -1.11 DPS) [quest]; Kodo Rustler Boots (15697, -1.31 DPS) [quest]; Acidic Walkers (9454, -1.32 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 18.2 healing_power points (1.37 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Voodoo Band (1996, -0.04 DPS) [world]; Mindbender Loop (5009, -0.04 DPS) [world_drop] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.35 DPS) | yes | Ogremind Ring (1993, -0.02 DPS) [world_drop]; Voodoo Band (1996, -0.02 DPS) [world]; Mindbender Loop (5009, -0.03 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wind Spirit Staff (6689, -0.72 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -1.44 DPS) [dungeon]; Death Speaker Scepter (2816, -2.03 DPS) [dungeon] |
| off_hand | Prophetic Cane (6803) | Into The Scarlet Monastery [quest] | sim-verified (136.6 DPS) | yes | Witch's Finger (16887, -0.75 DPS) [quest]; Aurora Sphere (7610, -0.90 DPS) [dungeon]; Skull of Impending Doom (4984, -0.94 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.3 healing_power points (1.15 DPS) | yes | Dancing Flame (6806, -0.21 DPS) [quest]; Flash Wand (5248, -0.37 DPS) [quest]; Goblin Igniter (5253, -0.37 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Prophetic Cane; ranged: Jaina's Firestarter

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 025003031305101520-03502000000000000-000000000000000000)

Set DPS (verified): 197.3. Weights run: 11.8s. Verify run: 51.4s. 421 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.502 ± 0.033, spirit=1.680 ± 0.045, mp5=4.037 ± 0.076, crit=0.602 ± 0.017 per rating point (14 rating = 1%, 8.435 per %), spell_haste=-6.644 ± 0.420

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 81.8 healing_power points (6.63 DPS) | yes | Blood Guard's Satin Cover (220899, -0.01 DPS) [vendor]; Soulcatcher Halo (10630, -0.20 DPS) [dungeon]; Papal Fez (9431, -0.51 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 41.7 healing_power points (3.38 DPS) | yes | Darkmoon Necklace (19303, +0.00 DPS, sim-verified) [vendor]; Gemshard Heart (17707, -0.54 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.71 DPS) [dungeon] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | sim-verified (197.3 DPS) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Inquisitor's Shawl (19507, -0.69 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 43.4 healing_power points (3.52 DPS) | yes | Featherskin Cape (10843, -0.73 DPS, sim-verified) [world]; Imperial Red Cloak (8248, -0.74 DPS) [world_drop]; Battle Healer's Cloak (19527, -0.92 DPS) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 92.9 healing_power points (7.53 DPS) | yes | Vestments of the Atal'ai Prophet (10806, -1.63 DPS) [dungeon]; Stormcloth Vest (10020, -1.74 DPS) [crafted]; Robes of Insight (940, -2.51 DPS, sim-verified) [world_drop] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 47.6 healing_power points (3.86 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -1.02 DPS) [quest]; Forgotten Wraps (9433, -1.43 DPS) [world_drop] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 81.8 healing_power points (6.63 DPS) | yes | Gilded Gloves (254095, -1.16 DPS) [crafted]; Virtuous Hands (226958, -2.36 DPS) [vendor]; Raider Handwraps (272098, -2.67 DPS, sim-verified) [vendor] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Waistcord (254081, -0.05 DPS) [crafted]; Earthenweave Cord (254077, -0.62 DPS) [crafted]; Dawnspire Cord (12466, -2.45 DPS, sim-verified) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 81.3 healing_power points (6.59 DPS) | yes | Senior Designer's Pantaloons (11841, -0.69 DPS) [dungeon]; Stone Guard's Satin Leggings (220902, -0.82 DPS) [vendor]; Dalewind Trousers (13008, -1.36 DPS) [world_drop] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | 53.5 healing_power points (4.34 DPS) | yes | Coldstone Slippers (18697, -0.19 DPS) [dungeon]; Mistwalker Boots (10629, -0.26 DPS) [dungeon]; Vinerot Sandals (17748, -0.86 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 45.9 healing_power points (3.72 DPS) | yes | Cyclopean Band (11824, -1.03 DPS) [dungeon]; Mindseye Circle (10634, -1.29 DPS) [dungeon]; Snake Hoop (6750, -1.35 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 35.2 healing_power points (2.85 DPS) | yes | Cyclopean Band (11824, -0.30 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -0.42 DPS) [dungeon]; Snake Hoop (6750, -0.48 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -3.27 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.55 DPS) [quest]; Alchemist's Stone (13503, -4.36 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.73 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -0.91 DPS) [crafted] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, +0.00 DPS) [dungeon]; Spellshifter Rod (9527, -3.25 DPS) [quest]; Wind Spirit Staff (6689, -3.34 DPS) [dungeon] |
| off_hand | Enthralled Sphere (11625) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 43.4 healing_power points (3.52 DPS) | yes | Cloud Stone (17737, +0.00 DPS, sim-verified) [dungeon]; Twisting Essence Jar (249456, -0.89 DPS) [crafted]; Desertwalker Cane (12471, -0.94 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 20.1 healing_power points (1.63 DPS) | yes | Flash Wand (5248, -0.41 DPS) [quest]; Nature's Breath (19118, -0.41 DPS) [quest]; Goblin Igniter (5253, -0.76 DPS, sim-verified) [quest] |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Nethergeld Shoulders; back: Darkspear Raider's Cloak; chest: Embrace of the Wind Serpent; wrist: Aristocratic Cuffs; hands: Virtuous Mitts; legs: Kilt of the Atal'ai Prophet; feet: Gilded Sandals; finger1: Brainlash; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; off_hand: Enthralled Sphere

No-known-source sample (15 of 421, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 316.5. Weights run: 11.6s. Verify run: 68.8s. 1112 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.101, intellect=2.853 ± 0.047, spirit=5.865 ± 0.121, mp5=2.109 ± 0.143, crit=0.762 ± 0.031 per rating point (14 rating = 1%, 10.669 per %), spell_haste=-16.366 ± 0.818

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 186.0 healing_power points (15.88 DPS) | yes | Devout Crown (16693, -0.82 DPS) [dungeon]; Crown of Caer Darrow (13986, -0.99 DPS) [quest]; Warlord's Satin Hood (231635, -1.05 DPS) [vendor] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 117.1 healing_power points (9.99 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Lady Maye's Pendant (14558, -0.36 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.60 DPS) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 130.3 healing_power points (11.12 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Virtuous Mantle (226951, -0.07 DPS) [quest]; Devout Mantle (16695, -0.56 DPS) [dungeon] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | sim-verified (316.6 DPS) | yes | Featherskin Cape (10843, -0.45 DPS) [world]; Butcher's Apron (12608, -0.92 DPS) [dungeon]; Shroud of the Exile (15421, -1.77 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 224.4 healing_power points (19.16 DPS) | yes | Mooncloth Vest (14138, -2.56 DPS) [crafted]; Alanna's Embrace (13314, -2.57 DPS) [dungeon]; Vestments of the Atal'ai Prophet (10806, -2.96 DPS) [dungeon] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, +0.00 DPS) [dungeon]; Wyrmthalak's Shackles (13958, -0.52 DPS) [quest]; General's Satin Bracers (17619, -0.59 DPS) [pvp] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 161.0 healing_power points (13.74 DPS) | yes | Devout Gloves (16692, -1.51 DPS) [dungeon]; Hands of the Exalted Herald (12554, -1.75 DPS) [dungeon]; Swarmtender's Gloves (275607, -2.36 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Penitent's Cinch (272394, +0.00 DPS) [vendor]; Wisdom of the Timbermaw (19047, -0.68 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 201.7 healing_power points (17.22 DPS) | yes | Haunting Specter Leggings (11929, -0.28 DPS) [dungeon]; Virtuous Skirt (226946, -1.66 DPS) [quest]; The Postmaster's Trousers (13389, -2.33 DPS) [dungeon] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | 178.6 healing_power points (15.25 DPS) | yes | Incandescent Mooncloth Boots (227862, -1.07 DPS) [vendor]; Mooncloth Boots (15802, -3.28 DPS) [crafted]; Devout Sandals (16691, -3.36 DPS) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of the Hierophant (13096, -1.30 DPS) [world_drop]; Seal of Rivendare (13345, -1.60 DPS) [dungeon]; Eye of Adaegus (5266, -1.77 DPS) [world_drop] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (316.6 DPS) | yes | Band of the Hierophant (13096, -0.27 DPS) [world_drop]; Seal of Rivendare (13345, -0.56 DPS) [dungeon]; Eye of Adaegus (5266, -0.74 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -0.34 DPS) [world_drop]; Soulkeeper (1607, -1.75 DPS) [world_drop]; Wind Spirit Staff (6689, -2.15 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 81.5 healing_power points (6.96 DPS) | yes | Chillnail Splinter (10704, -3.95 DPS) [quest]; Jaina's Firestarter (13064, -4.00 DPS) [world_drop]; Oblivion's Touch (18761, -4.28 DPS) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Frostweaver Cape; wrist: Bracers of Hope; waist: Virtuous Belt; legs: Devout Skirt; feet: Virtuous Sandals; finger1: The Postmaster's Seal; finger2: Emerald Flame Ring; trinket2: Royal Seal of Eldre'Thalas; main_hand: Dancing Sliver; ranged: Torch of Light

No-known-source sample (15 of 1112, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 025003031305101520-03505003030100000-000000000000000000)

Set DPS (verified): 576.5. Weights run: 7.6s. Verify run: 36.4s. 1112 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.251, intellect=3.031 ± 0.065, spirit=2.557 ± 0.062, mp5=3.741 ± 0.067, crit=0.952 ± 0.047 per rating point (14 rating = 1%, 13.329 per %), spell_haste=not significant (-1.927 ± 1.308)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 138.8 healing_power points (16.27 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.99 DPS) [pvp]; Devout Crown (16693, -10.26 DPS, sim-verified) [dungeon] |
| neck | The Eye of Zuldazar (19593) | The Eye of Zuldazar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 107.0 healing_power points (12.55 DPS) | yes | Champion's Satin Mantle (227120, -0.45 DPS) [pvp]; Virtuous Mantle (226951, -1.75 DPS) [quest]; Devout Mantle (16695, -9.90 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 72.3 healing_power points (8.48 DPS) | yes | Darkspear Raider's Cloak (272063, -0.99 DPS) [vendor]; Shroud of the Exile (15421, -1.05 DPS) [quest]; Frostweaver Cape (12968, -12.16 DPS, sim-verified) [dungeon] |
| chest | Virtuous Robe (226945) | Saving the Best for Last [quest] | 137.2 healing_power points (16.08 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Alanna's Embrace (13314, -0.64 DPS) [dungeon]; Mooncloth Vest (14138, -8.43 DPS, sim-verified) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 82.0 healing_power points (9.62 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Bracers of Mending (23129, +0.00 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -1.01 DPS) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 105.8 healing_power points (12.41 DPS) | yes | Virtuous Mitts (226950, +0.00 DPS, sim-verified) [vendor]; Hands of the Exalted Herald (12554, -0.32 DPS) [dungeon]; Devout Gloves (16692, -2.00 DPS) [dungeon] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | 97.6 healing_power points (11.44 DPS) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Devout Belt (16696, -0.34 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -6.25 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | sim-verified (+12.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Padre's Trousers (18386, -0.05 DPS) [dungeon]; Devout Skirt (16694, -12.36 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 122.2 healing_power points (14.33 DPS) | yes | Virtuous Sandals (226952, +0.00 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -2.64 DPS) [crafted]; Faith Healer's Boots (22247, -3.42 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, -1.07 DPS) [dungeon]; Seal of Rivendare (13345, -1.36 DPS) [dungeon]; Naglering (11669, -14.82 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, +0.00 DPS) [dungeon]; Emerald Flame Ring (18395, +0.00 DPS) [dungeon]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -1.75 DPS) [dungeon]; Ankh of Life (1713, -2.98 DPS) [world_drop]; Royal Seal of Eldre'Thalas (18469, -13.99 DPS, sim-verified) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18469, +0.00 DPS) [quest]; Mindtap Talisman (18371, -0.10 DPS) [dungeon]; Ankh of Life (1713, -1.33 DPS) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, -2.91 DPS) [world_drop]; Wind Spirit Staff (6689, -3.49 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | sim-verified (576.4 DPS) | yes | Thaurissan's Royal Scepter (11928, +0.00 DPS) [dungeon]; Book of the Dead (13353, +0.00 DPS) [dungeon]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 45.1 healing_power points (5.29 DPS) | yes | Sparkling Crystal Wand (20672, -1.51 DPS) [world]; Jaina's Firestarter (13064, -2.26 DPS) [world_drop]; Oblivion's Touch (18761, -2.90 DPS, sim-verified) [dungeon] |

**New at 60:** head: Virtuous Crown; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Virtuous Robe; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 1112, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

