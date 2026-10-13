# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 40.9. Weights run: 1.9s. Verify run: 5.9s. 148 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=1.324 ± 0.008, spirit=0.604 ± 0.004, mp5=1.547 ± 0.015, crit=0.068 ± 0.005 per rating point (14 rating = 1%, 0.952 per %), spell_haste=not significant (-0.012 ± 0.020)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.63 DPS) | yes | Shadow Goggles (4373, -1.12 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.41 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 2.4 healing_power points (0.22 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.9 healing_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -0.66 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.94 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 10.2 healing_power points (0.93 DPS) | yes | Pearl-clasped Cloak (5542, -0.57 DPS) [crafted]; Seer's Cape (6378, -0.58 DPS) [dungeon]; Sanguine Cape (14376, -0.83 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.4 healing_power points (1.94 DPS) | yes | Robe of the Moccasin (6465, -0.58 DPS) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon]; Corsair's Overshirt (5202, -0.82 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 healing_power points (1.02 DPS) | yes | Repurposed Hair Band (281256, -0.67 DPS) [quest]; Mystic's Bracelets (14366, -0.78 DPS) [world_drop]; Bright Bracers (3647, -1.14 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Tomb Robber's Gloves (280096, -0.64 DPS) [quest]; Bright Gloves (3066, -0.66 DPS) [world_drop]; Magefist Gloves (12977, -1.04 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Keller's Girdle (2911, -0.52 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.58 DPS, sim-verified) [crafted]; Wise Man's Belt (4786, -0.88 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.57 DPS) | yes | Abomination Skin Leggings (23173, -1.61 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Darkweave Breeches (12987, -1.99 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Spidersilk Boots (4320, -0.70 DPS) [crafted]; Sanguine Sandals (14374, -0.70 DPS) [world_drop]; Kimbra Boots (6191, -0.92 DPS, sim-verified) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.9 healing_power points (0.72 DPS) | yes | Band of Purification (12996, -0.39 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.44 DPS) [rep]; Deep Fathom Ring (6463, -0.45 DPS) [dungeon] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | 4.0 healing_power points (0.36 DPS) | yes | Lorekeeper's Ring (20431, -0.08 DPS) [rep]; Deep Fathom Ring (6463, -0.09 DPS) [dungeon]; Band of Purification (12996, -0.21 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 13.2 healing_power points (1.20 DPS) | yes | Lesser Staff of the Spire (1300, -0.48 DPS) [world]; Staff of Westfall (2042, -0.53 DPS, sim-verified) [quest]; Witching Stave (1484, -0.60 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) (or Moonstone Wand (15204)) | The Escape [quest] | 2.6 healing_power points (0.24 DPS) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.08 DPS) [quest]; Elven Wand (5604, -0.12 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 65.9. Weights run: 2.0s. Verify run: 6.8s. 256 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.026, intellect=1.413 ± 0.010, spirit=0.849 ± 0.006, mp5=1.645 ± 0.022, crit=0.097 ± 0.007 per rating point (14 rating = 1%, 1.352 per %), spell_haste=0.349 ± 0.067

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.1 healing_power points (3.41 DPS) | yes | Nightsky Cowl (4039, -1.51 DPS) [world_drop]; Resilient Cap (14401, -1.71 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -1.72 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.0 healing_power points (0.81 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Scorn's Icy Choker (23169, -0.09 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272075, -0.18 DPS) [vendor] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 15.8 healing_power points (1.42 DPS) | yes | Death Speaker Mantle (6685, -0.03 DPS) [dungeon]; Nightsky Mantle (4718, -0.06 DPS, sim-verified) [world_drop]; Faerie Mantle (5820, -0.20 DPS) [quest] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 17.6 healing_power points (1.57 DPS) | yes | Caretaker's Cape (19533, -0.11 DPS) [rep]; Prelacy Cape (7004, -0.13 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.33 DPS) [vendor] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.5 healing_power points (3.18 DPS) | yes | Pristine Gown (253961, -0.30 DPS, sim-verified) [crafted]; Robes of Arugal (6324, -1.11 DPS) [dungeon]; Filigreed Pristine Gown (253901, -1.16 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.0 healing_power points (1.07 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.26 DPS, sim-verified) [world_drop]; Stonecloth Bindings (14416, -0.44 DPS) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.72 DPS) | yes | Gloves of Old (9395, -0.27 DPS, sim-verified) [world_drop]; Hotshot Pilot's Gloves (9491, -0.33 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.33 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.7 healing_power points (1.49 DPS) | yes | Resilient Cord (14406, -0.43 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.44 DPS, sim-verified) [crafted]; Dreamer's Belt (4829, -0.45 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.1 healing_power points (3.23 DPS) | yes | Filigreed Pristine Leggings (253937, -0.69 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.06 DPS) [crafted]; Necromancer Leggings (2277, -1.84 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.3 healing_power points (2.98 DPS) | yes | Acidic Walkers (9454, -1.66 DPS) [dungeon]; Nimbus Boots (6998, -1.73 DPS, sim-verified) [quest]; Glinteye Slippers (273024, -1.76 DPS) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.61 DPS) | yes | Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; The Queen's Jewel (13094, -0.75 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.85 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 15.8 healing_power points (1.42 DPS) | yes | Electrocutioner Lagnut (9447, -0.56 DPS, sim-verified) [dungeon]; The Queen's Jewel (13094, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.66 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | sim-verified (65.9 DPS) | yes | Glimmering Staff (249392, -2.28 DPS) [crafted]; Lorekeeper's Staff (212580, -2.32 DPS) [vendor]; Death Speaker Scepter (2816, -2.95 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Summoner's Wand (5245) (or Starfaller (13063), Lesser Mystic Wand (11289)) | Dalaran Summoner [world] | 5.7 healing_power points (0.51 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Starfaller (13063, +0.00 DPS) [world_drop]; Consecrated Wand (5244, -0.10 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Repairman's Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; main_hand: Hardened Root Staff; ranged: Summoner's Wand

No-known-source sample (15 of 256, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 83.6. Weights run: 2.0s. Verify run: 15.9s. 331 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.127, intellect=4.455 ± 0.051, spirit=0.708 ± 0.030, mp5=0.656 ± 0.036, crit=0.207 ± 0.024 per rating point (14 rating = 1%, 2.901 per %), spell_haste=not significant (1.312 ± 0.820)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Craftsman's Monocle (4393) | Engineering [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Cowl of Necromancy (2621, +0.00 DPS) [world]; Nightsky Cowl (4039, -0.45 DPS) [world_drop]; Thinking Cap (2624, -1.01 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 48.6 healing_power points (2.25 DPS) | yes | Darkspear Warding Pendant (272074, -0.81 DPS) [vendor]; Scorn's Icy Choker (23169, -1.01 DPS) [dungeon]; Necklace of Calisea (1714, -3.36 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 72.9 healing_power points (3.38 DPS) | yes | Mistscape Mantle (4734, -0.94 DPS) [dungeon]; Silksand Shoulder Pads (14423, -1.01 DPS) [world_drop]; Windchaser Amice (14432, -1.70 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 51.8 healing_power points (2.40 DPS) | yes | Blackforge Cape (6424, -0.62 DPS) [dungeon]; Glyphed Cloak (4732, -0.96 DPS) [world_drop]; Sergeant Major's Cape (16336, -0.97 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silksand Tunic (14417, +0.00 DPS) [world_drop]; Silksand Wraps (14425, +0.00 DPS) [world_drop]; Red Mageweave Vest (10007, -10.88 DPS, sim-verified) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 40.1 healing_power points (1.86 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -0.21 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 64.5 healing_power points (2.98 DPS) | yes | Town Clerk's Mittens (270029, -0.72 DPS) [quest]; Red Mageweave Gloves (10018, -0.92 DPS) [crafted]; Gilded Handwraps (254021, -3.08 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 66.8 healing_power points (3.09 DPS) | yes | Gilded Cord (254037, -0.23 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.62 DPS) [quest]; Teacher's Sash (10747, -0.62 DPS) [quest] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Silk Pantaloons (7062, +0.00 DPS) [crafted]; Pristine Leggings (253987, +0.00 DPS) [crafted]; Stoneweaver Leggings (9407, -2.79 DPS, sim-verified) [dungeon] |
| feet | Furen's Boots (13100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kodo Rustler Boots (15697, -0.15 DPS) [quest]; Acidic Walkers (9454, -0.19 DPS) [dungeon]; Gilded Slippers (254001, -2.45 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 36.1 healing_power points (1.67 DPS) | yes | Ogremind Ring (1993, -0.13 DPS) [world_drop]; Mindbender Loop (5009, -0.16 DPS) [world_drop]; Welken Ring (5011, -0.41 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 33.3 healing_power points (1.54 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.03 DPS) [world_drop]; Welken Ring (5011, -0.28 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+10.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Hypnotic Blade (7714, -0.83 DPS) [dungeon]; Illusionary Rod (7713, -1.11 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 56.6 healing_power points (2.62 DPS) | yes | Arcane Infused Rod (279838, -0.39 DPS, sim-verified) [quest]; Aurora Sphere (7610, -1.04 DPS) [dungeon]; Skull of Impending Doom (4984, -1.17 DPS) [quest] |
| ranged | Goblin Igniter (5253) | Venture Company Mining [quest] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Flash Wand (5248, +0.00 DPS) [quest]; Summoner's Wand (5245, -0.10 DPS) [world]; Jaina's Firestarter (13064, -1.01 DPS, sim-verified) [world_drop] |

**New at 40:** head: Craftsman's Monocle; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Deathmage Sash; legs: Stormcloth Pants; feet: Furen's Boots; finger1: Snake Hoop; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica; ranged: Goblin Igniter

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 104.4. Weights run: 2.1s. Verify run: 18.7s. 422 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.079, intellect=4.782 ± 0.085, spirit=0.224 ± 0.042, mp5=1.345 ± 0.085, crit=0.285 ± 0.029 per rating point (14 rating = 1%, 3.988 per %), spell_haste=-25.421 ± 1.490

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Satin Cover (220896) | Captain Dirgehammer [vendor] | sim-verified (104.4 DPS) | yes | Red Mageweave Headband (10033, -0.11 DPS) [crafted]; Gemburst Circlet (10751, -0.17 DPS) [quest]; Bad Mojo Mask (9470, -6.14 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 67.8 healing_power points (3.20 DPS) | yes | Glowing Eye of Mordresh (10769, -0.78 DPS) [dungeon]; Darkspear Warding Pendant (272073, -1.17 DPS) [vendor]; Darkmoon Necklace (19303, -1.47 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Satin Pads (220894) | Captain Dirgehammer [vendor] | sim-verified (104.4 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Imperial Red Mantle (8250, -0.16 DPS) [world_drop]; Rotgrip Mantle (17732, -5.26 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 68.1 healing_power points (3.21 DPS) | yes | Imperial Red Cloak (8248, -0.79 DPS, sim-verified) [world_drop]; Keeper's Cloak (14665, -0.96 DPS) [world_drop]; Highborne Cloak (14450, -1.14 DPS) [world_drop] |
| chest | Knight's Satin Armor (220892) | Captain Dirgehammer [vendor] | sim-verified (104.4 DPS) | yes | Acumen Robes (17775, +0.00 DPS) [quest]; Red Mageweave Vest (10007, -0.20 DPS) [crafted]; Hibernal Robe (8113, -5.94 DPS, sim-verified) [world_drop] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | 58.1 healing_power points (2.74 DPS) | yes | Forgotten Wraps (9433, -0.03 DPS) [world_drop]; Nethergeld Cuffs (254061, -0.16 DPS) [crafted]; Imperial Red Bracers (8247, -0.26 DPS) [world_drop] |
| hands | Sergeant Major's Satin Gloves (220897) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (104.4 DPS) | yes | Virtuous Mitts (226950, +0.00 DPS) [vendor]; Gilded Gloves (254095, +0.00 DPS) [crafted]; Raider Handwraps (272098, -0.50 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 102.9 healing_power points (4.85 DPS) | yes | Gilded Waistcord (254081, -1.45 DPS) [crafted]; Imperial Red Sash (8253, -1.47 DPS) [world_drop]; Serenity Belt (13144, -2.05 DPS, sim-verified) [world_drop] |
| legs | Knight's Satin Leggings (220893) | Captain Dirgehammer [vendor] | sim-verified (104.4 DPS) | yes | Venomshroud Leggings (14444, -0.75 DPS) [world_drop]; Imperial Red Pants (8251, -0.88 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -4.33 DPS, sim-verified) [dungeon] |
| feet | Sergeant Major's Satin Boots (220895) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 77.4 healing_power points (3.65 DPS) | yes | Highborne Footpads (14447, -0.94 DPS) [world_drop]; Vinerot Sandals (17748, -0.95 DPS) [dungeon]; Gilded Sandals (254107, -3.35 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 35.0 healing_power points (1.65 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Eye of Adaegus (5266, -0.17 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 34.1 healing_power points (1.61 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.01 DPS) [world_drop]; Eye of Adaegus (5266, -0.13 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (104.4 DPS) | yes | Ankh of Life (1713, -0.83 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.85 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.87 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (104.4 DPS) | yes | Evonice's Landin' Pilla (18951, -0.74 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.76 DPS) [quest]; Ankh of Life (1713, -0.87 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (104.4 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Spellshifter Rod (9527, -0.34 DPS) [quest]; Glowing Brightwood Staff (812, -1.44 DPS) [world_drop] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 59.8 healing_power points (2.82 DPS) | yes | Mistscape Stave (7611, -0.79 DPS) [world_drop]; Skullspell Orb (10708, -0.90 DPS, sim-verified) [quest]; Arcane Infused Rod (279838, -1.02 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 29.4 healing_power points (1.39 DPS) | yes | Cairnstone Sliver (9654, -0.20 DPS) [quest]; Flash Wand (5248, -0.45 DPS) [quest]; Goblin Igniter (5253, -0.45 DPS) [quest] |

**New at 50:** head: Knight-Lieutenant's Satin Cover; neck: Horizon Choker; shoulder: Knight-Lieutenant's Satin Pads; back: Darkspear Raider's Cloak; chest: Knight's Satin Armor; wrist: Shizzle's Nozzle Wiper; hands: Sergeant Major's Satin Gloves; waist: Dawnspire Cord; legs: Knight's Satin Leggings; feet: Sergeant Major's Satin Boots; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; ranged: Jaina's Firestarter

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 383.0. Weights run: 2.1s. Verify run: 20.5s. 933 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.161, intellect=1.080 ± 0.069, spirit=0.442 ± 0.099, mp5=0.658 ± 0.068, crit=0.165 ± 0.030 per rating point (14 rating = 1%, 2.316 per %), spell_haste=-7.216 ± 1.412

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 60.8 healing_power points (13.39 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -16.74 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (383.0 DPS) | yes | The Eye of Zuldazar (19593, -0.74 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -0.74 DPS) [quest]; Drake Tooth Necklace (21531, -4.04 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 67.4 healing_power points (14.85 DPS) | yes | Field Marshal's Satin Mantle (231628, -3.48 DPS) [pvp]; Lieutenant Commander's Satin Mantle (227119, -3.80 DPS) [pvp]; Mooncloth Shoulders (14139, -9.43 DPS, sim-verified) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.8 healing_power points (11.64 DPS) | yes | Drape of Recovery (272413, -4.41 DPS, sim-verified) [vendor]; Caretaker's Cape (19530, -5.13 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 80.1 healing_power points (17.65 DPS) | yes | Field Marshal's Satin Tunic (231624, -1.07 DPS) [vendor]; Virtuous Robe (226945, -3.35 DPS, sim-verified) [quest]; Knight-Captain's Satin Tunic (227122, -3.90 DPS) [pvp] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nethergeld Cuffs (254061, -0.62 DPS) [crafted]; Earthenweave Cuffs (254131, -1.36 DPS) [crafted]; Bracers of Hope (22667, -11.91 DPS, sim-verified) [quest] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 58.6 healing_power points (12.92 DPS) | yes | Raider Handwraps (272097, -0.60 DPS, sim-verified) [vendor]; Mooncloth Gloves (18409, -1.84 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Waistcord (254081, -0.43 DPS) [crafted]; Devout Belt (16696, -0.90 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -9.27 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 56.2 healing_power points (12.39 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Devout Skirt (16694, -13.79 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 65.3 healing_power points (14.39 DPS) | yes | Marshal's Satin Walkers (231627, -3.72 DPS) [vendor]; Virtuous Sandals (226952, -4.65 DPS) [quest]; Mooncloth Boots (15802, -6.21 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 48.0 healing_power points (10.58 DPS) | yes | Blessed Band of Light (272407, -2.97 DPS) [vendor]; Channeler's Ring (272406, -6.47 DPS) [vendor]; Sea Giant's Toe Ring (274746, -6.61 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 35.4 healing_power points (7.81 DPS) | yes | Blessed Band of Light (272407, -1.68 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.70 DPS) [vendor]; Sea Giant's Toe Ring (274746, -3.84 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -1.01 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.20 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -5.29 DPS) [world_drop]; Ankh of Life (1713, -8.09 DPS) [world_drop]; Shard of the Splithooves (10659, -9.41 DPS, sim-verified) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -6.98 DPS) [dungeon]; Spellshifter Rod (9527, -11.10 DPS) [quest] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 43.3 healing_power points (9.54 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Twisting Essence Jar (249456, -4.57 DPS) [crafted]; Orb of Lorica (11262, -7.67 DPS, sim-verified) [quest] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 21.9 healing_power points (4.82 DPS) | yes | Cairnstone Sliver (9654, -3.14 DPS) [quest]; Brilliant Wand (249385, -3.15 DPS) [crafted]; Jaina's Firestarter (13064, -5.00 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 933, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 758.4. Weights run: 1.8s. Verify run: 11.2s. 933 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.417, intellect=1.190 ± 0.059, spirit=1.056 ± 0.060, mp5=2.129 ± 0.068, crit=0.443 ± 0.043 per rating point (14 rating = 1%, 6.199 per %), spell_haste=not significant (-0.487 ± 1.301)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 76.3 healing_power points (21.94 DPS) | yes | Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Lieutenant Commander's Satin Hood (227121, -0.04 DPS) [pvp]; Mooncloth Circlet (14140, -16.68 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 48.2 healing_power points (13.85 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -1.99 DPS) [quest]; The Eye of Zuldazar (19593, -2.07 DPS, sim-verified) [quest]; Drake Tooth Necklace (21531, -2.27 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 82.3 healing_power points (23.67 DPS) | yes | Lieutenant Commander's Satin Mantle (227119, -6.40 DPS) [pvp]; Field Marshal's Satin Mantle (231628, -8.20 DPS) [pvp]; Virtuous Mantle (226951, -13.12 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.9 healing_power points (15.49 DPS) | yes | Drape of Recovery (272413, -5.02 DPS, sim-verified) [vendor]; Caretaker's Cape (19530, -5.59 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 89.9 healing_power points (25.84 DPS) | yes | Field Marshal's Satin Tunic (231624, -1.75 DPS) [vendor]; Virtuous Robe (226945, -4.65 DPS, sim-verified) [quest]; Knight-Captain's Satin Tunic (227122, -4.96 DPS) [pvp] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (758.5 DPS) | yes | Nethergeld Cuffs (254061, -1.93 DPS) [crafted]; Marshal's Satin Bracers (17606, -2.06 DPS) [pvp]; Bracers of Hope (22667, -12.42 DPS, sim-verified) [quest] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 62.2 healing_power points (17.89 DPS) | yes | Virtuous Mitts (226950, -1.76 DPS) [vendor]; Mooncloth Gloves (18409, -2.93 DPS) [crafted]; Raider Handwraps (272097, -3.48 DPS, sim-verified) [vendor] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (758.5 DPS) | yes | Penitent's Cinch (272394, -1.17 DPS) [vendor]; Devout Belt (16696, -2.33 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -10.76 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 69.1 healing_power points (19.85 DPS) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Devout Skirt (16694, -15.06 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 75.4 healing_power points (21.69 DPS) | yes | Virtuous Sandals (226952, -5.65 DPS, sim-verified) [quest]; Mooncloth Boots (15802, -6.05 DPS) [crafted]; Marshal's Satin Walkers (231627, -7.19 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 56.4 healing_power points (16.20 DPS) | yes | Blessed Band of Light (272407, -4.86 DPS) [vendor]; Channeler's Ring (272406, -8.98 DPS) [vendor]; Band of the Hierophant (13096, -9.44 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 42.4 healing_power points (12.19 DPS) | yes | Blessed Band of Light (272407, -0.85 DPS) [vendor]; Channeler's Ring (272406, -4.97 DPS) [vendor]; Band of the Hierophant (13096, -5.43 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (758.5 DPS) | yes | Uther's Strength (11302, -6.90 DPS) [world_drop]; Shard of the Splithooves (10659, -8.40 DPS) [quest]; Ankh of Life (1713, -8.43 DPS) [world_drop] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (758.5 DPS) | yes | Uther's Strength (11302, -2.56 DPS, sim-verified) [world_drop]; Shard of the Splithooves (10659, -5.51 DPS) [quest]; Ankh of Life (1713, -5.54 DPS) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (758.5 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -8.88 DPS) [dungeon]; Spellshifter Rod (9527, -14.94 DPS) [quest] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 50.6 healing_power points (14.56 DPS) | yes | Grand Marshal's Tome of Restoration (234590, -1.43 DPS) [pvp]; Orb of Lorica (11262, -7.69 DPS) [quest]; Twisting Essence Jar (249456, -7.83 DPS, sim-verified) [crafted] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 28.6 healing_power points (8.23 DPS) | yes | Jaina's Firestarter (13064, -5.26 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -5.59 DPS) [quest]; Cairnstone Sliver (9654, -5.88 DPS, sim-verified) [quest] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Serenity Field; trinket2: Darkspear Voodoo Seal; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 933, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 38.8. Weights run: 1.9s. Verify run: 5.8s. 138 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=1.324 ± 0.008, spirit=0.604 ± 0.004, mp5=1.547 ± 0.015, crit=0.068 ± 0.005 per rating point (14 rating = 1%, 0.952 per %), spell_haste=not significant (-0.012 ± 0.020)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.63 DPS) | yes | Shadow Goggles (4373, -0.84 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.41 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.8 healing_power points (0.16 DPS) | yes | Scholarly Pendant (277203, -0.07 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.9 healing_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -0.66 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.91 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 10.2 healing_power points (0.93 DPS) | yes | Pearl-clasped Cloak (5542, -0.57 DPS) [crafted]; Seer's Cape (6378, -0.58 DPS) [dungeon]; Sanguine Cape (14376, -0.66 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 21.4 healing_power points (1.94 DPS) | yes | Corsair's Overshirt (5202, -0.36 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.58 DPS) [dungeon]; Bloody Apron (6226, -0.77 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.3 healing_power points (1.02 DPS) | yes | Bright Bracers (3647, -0.54 DPS) [world_drop]; Crystalline Cuffs (14148, -0.62 DPS) [dungeon]; Featherbead Bracers (15452, -1.02 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 15.0 healing_power points (1.36 DPS) | yes | Blight Gloves (279877, -0.11 DPS, sim-verified) [quest]; Magefist Gloves (12977, -0.54 DPS) [world_drop]; Bright Gloves (3066, -0.66 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.3 healing_power points (1.48 DPS) | yes | Novice Ardent's Sash (253887, -0.44 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.52 DPS) [world_drop]; Wise Man's Belt (4786, -0.88 DPS) [vendor] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 28.4 healing_power points (2.57 DPS) | yes | Abomination Skin Leggings (23173, -1.61 DPS) [dungeon]; Filigreed Silky Leggings (253939, -1.63 DPS) [crafted]; Darkweave Breeches (12987, -1.78 DPS, sim-verified) [world_drop] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 13.0 healing_power points (1.18 DPS) | yes | Spidersilk Boots (4320, -0.70 DPS) [crafted]; Bluegill Sandals (1560, -0.72 DPS) [world]; Sanguine Sandals (14374, -1.02 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.9 healing_power points (0.72 DPS) | yes | Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Band of Purification (12996, -0.39 DPS) [world_drop]; Advisor's Ring (20426, -0.44 DPS) [rep] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.6 healing_power points (0.60 DPS) | yes | Band of Purification (12996, -0.27 DPS) [world_drop]; Advisor's Ring (20426, -0.32 DPS) [rep]; Volcanic Rock Ring (12053, -0.45 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 13.2 healing_power points (1.20 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.17 DPS) [quest]; Advisor's Gnarled Staff (20425, -0.30 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 2.6 healing_power points (0.24 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.12 DPS) [world] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 138, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263007 Skyseer's Vest; 263402 Sash of Sorrow

### Band 30 (troll, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 60.6. Weights run: 2.0s. Verify run: 6.3s. 243 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.026, intellect=1.413 ± 0.010, spirit=0.849 ± 0.006, mp5=1.645 ± 0.022, crit=0.097 ± 0.007 per rating point (14 rating = 1%, 1.352 per %), spell_haste=0.349 ± 0.067

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 38.1 healing_power points (3.41 DPS) | yes | Nightsky Cowl (4039, -1.51 DPS) [world_drop]; Filigreed Pristine Circlet (253975, -1.64 DPS, sim-verified) [crafted]; Resilient Cap (14401, -1.71 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 9.0 healing_power points (0.81 DPS) | yes | Scorn's Icy Choker (23169, -0.05 DPS) [dungeon]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.18 DPS) [vendor] |
| shoulder | Mantle of Woe (7750) | Vorrel's Revenge [quest] | 17.0 healing_power points (1.52 DPS) | yes | Nightsky Mantle (4718, -0.13 DPS) [world_drop]; Death Speaker Mantle (6685, -0.13 DPS) [dungeon]; Ghostly Mantle (3324, -0.13 DPS, sim-verified) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 16.4 healing_power points (1.47 DPS) | yes | Glowing Thresher Cape (6901, -0.23 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.28 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.71 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.5 healing_power points (3.18 DPS) | yes | Pristine Gown (253961, -0.17 DPS, sim-verified) [crafted]; Robes of Arugal (6324, -1.11 DPS) [dungeon]; Filigreed Pristine Gown (253901, -1.16 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 12.0 healing_power points (1.07 DPS) | yes | Nightsky Wristbands (6407, -0.09 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.38 DPS, sim-verified) [world_drop]; Featherbead Bracers (15452, -0.44 DPS) [quest] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.72 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.33 DPS) [dungeon]; Pristine Gloves (253913, -0.36 DPS) [crafted]; Gloves of Old (9395, -0.81 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.7 healing_power points (1.49 DPS) | yes | Lilac Sash (6780, -0.18 DPS, sim-verified) [quest]; Novice Ardent's Sash (253887, -0.35 DPS) [crafted]; Resilient Cord (14406, -0.43 DPS) [world_drop] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 36.1 healing_power points (3.23 DPS) | yes | Filigreed Pristine Leggings (253937, -0.66 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -1.06 DPS) [crafted]; Necromancer Leggings (2277, -1.84 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 33.3 healing_power points (2.98 DPS) | yes | Glinteye Slippers (273024, -1.76 DPS) [dungeon]; Frothing Slippers (254003, -1.79 DPS) [crafted]; Acidic Walkers (9454, -2.04 DPS, sim-verified) [dungeon] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.61 DPS) | yes | Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; The Queen's Jewel (13094, -0.75 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.85 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 15.8 healing_power points (1.42 DPS) | yes | Electrocutioner Lagnut (9447, -0.49 DPS, sim-verified) [dungeon]; The Queen's Jewel (13094, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.66 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 68.0 healing_power points (6.09 DPS) | yes | Glimmering Staff (249392, -0.39 DPS, sim-verified) [crafted]; Advisor's Gnarled Staff (19569, -4.61 DPS) [pvp]; Lorekeeper's Staff (212580, -4.74 DPS) [vendor] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 18.1 healing_power points (1.62 DPS) | yes | Orb of Mistmantle (13031, -0.56 DPS, sim-verified) [world_drop]; Defective Samophlange (274743, -0.61 DPS) [vendor]; Witch's Finger (16887, -0.61 DPS) [quest] |
| ranged | Summoner's Wand (5245) (or Starfaller (13063), Lesser Mystic Wand (11289)) | Dalaran Summoner [world] | 5.7 healing_power points (0.51 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Starfaller (13063, +0.00 DPS) [world_drop]; Wand of Decay (5252, -0.25 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Mantle of Woe; back: Battle Healer's Cloak; chest: Death Speaker Robes; hands: Truefaith Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Summoner's Wand

No-known-source sample (15 of 243, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 77.5. Weights run: 2.0s. Verify run: 16.6s. 315 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.127, intellect=4.455 ± 0.051, spirit=0.708 ± 0.030, mp5=0.656 ± 0.036, crit=0.207 ± 0.024 per rating point (14 rating = 1%, 2.901 per %), spell_haste=not significant (1.312 ± 0.820)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Thinking Cap (2624) | Dustbelcher Shaman [world] | 75.7 healing_power points (3.51 DPS) | yes | Cowl of Necromancy (2621, -0.41 DPS) [world]; Craftsman's Monocle (4393, -0.78 DPS, sim-verified) [crafted]; Nightsky Cowl (4039, -0.87 DPS) [world_drop] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Warding Pendant (272074, -0.23 DPS) [vendor]; Scorn's Icy Choker (23169, -0.44 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.85 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 72.9 healing_power points (3.38 DPS) | yes | Mantle of Woe (7750, -0.90 DPS) [quest]; Mistscape Mantle (4734, -0.94 DPS) [dungeon]; Windchaser Amice (14432, -1.59 DPS, sim-verified) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 51.8 healing_power points (2.40 DPS) | yes | Blackforge Cape (6424, -0.34 DPS, sim-verified) [dungeon]; Glyphed Cloak (4732, -0.96 DPS) [world_drop]; Windchaser Cloak (14430, -1.03 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silksand Tunic (14417, +0.00 DPS) [world_drop]; Silksand Wraps (14425, +0.00 DPS) [world_drop]; Red Mageweave Vest (10007, -13.83 DPS, sim-verified) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 40.1 healing_power points (1.86 DPS) | yes | Radiant Silver Bracers (4545, -0.21 DPS) [quest]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest]; Mistscape Bracers (4045, -0.28 DPS, sim-verified) [dungeon] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 64.5 healing_power points (2.98 DPS) | yes | Red Mageweave Gloves (10018, -0.92 DPS) [crafted]; Mistscape Gloves (6428, -1.13 DPS) [dungeon]; Gilded Handwraps (254021, -6.49 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Razzeric's Customized Seatbelt (6726, -0.39 DPS) [quest]; Mistscape Sash (4736, -0.60 DPS) [dungeon]; Deathmage Sash (10771, -4.75 DPS, sim-verified) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Silk Pantaloons (7062, +0.00 DPS) [crafted]; Pristine Leggings (253987, +0.00 DPS) [crafted]; Stoneweaver Leggings (9407, -7.82 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 54.0 healing_power points (2.50 DPS) | yes | Furen's Boots (13100, -0.53 DPS) [world_drop]; Kodo Rustler Boots (15697, -0.69 DPS) [quest]; Boots of the Maharishi (9658, -0.94 DPS, sim-verified) [quest] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 36.1 healing_power points (1.67 DPS) | yes | Ogremind Ring (1993, -0.13 DPS) [world_drop]; Mindbender Loop (5009, -0.16 DPS) [world_drop]; Welken Ring (5011, -0.41 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 33.3 healing_power points (1.54 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.03 DPS) [world_drop]; Welken Ring (5011, -0.28 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+16.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Hypnotic Blade (7714, -0.83 DPS) [dungeon]; Illusionary Rod (7713, -1.11 DPS) [dungeon] |
| off_hand | Aurora Sphere (7610) | Uldaman: Ancient Treasure [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Skull of Impending Doom (4984, -0.13 DPS) [quest]; Nightsky Orb (15929, -0.24 DPS) [world_drop]; Witch's Finger (16887, -2.42 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 28.9 healing_power points (1.34 DPS) | yes | Flash Wand (5248, -0.41 DPS) [quest]; Goblin Igniter (5253, -0.41 DPS) [quest]; Dancing Flame (6806, -2.79 DPS, sim-verified) [quest] |

**New at 40:** head: Thinking Cap; neck: Necklace of Calisea; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; finger1: Snake Hoop; finger2: Voodoo Band; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Aurora Sphere; ranged: Jaina's Firestarter

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (troll, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 93.1. Weights run: 2.1s. Verify run: 17.8s. 403 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.079, intellect=4.782 ± 0.085, spirit=0.224 ± 0.042, mp5=1.345 ± 0.085, crit=0.285 ± 0.029 per rating point (14 rating = 1%, 3.988 per %), spell_haste=-25.421 ± 1.490

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 100.4 healing_power points (4.74 DPS) | yes | Blood Guard's Satin Cover (220899, -0.12 DPS) [vendor]; Red Mageweave Headband (10033, -0.25 DPS, sim-verified) [crafted]; Gemburst Circlet (10751, -0.29 DPS) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 67.8 healing_power points (3.20 DPS) | yes | Glowing Eye of Mordresh (10769, -0.78 DPS) [dungeon]; Darkspear Warding Pendant (272073, -1.17 DPS) [vendor]; Darkmoon Necklace (19303, -1.47 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 86.1 healing_power points (4.06 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Satin Pads (220901, -0.51 DPS) [vendor]; Imperial Red Mantle (8250, -0.68 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 68.1 healing_power points (3.21 DPS) | yes | Imperial Red Cloak (8248, -0.40 DPS, sim-verified) [world_drop]; Keeper's Cloak (14665, -0.96 DPS) [world_drop]; Highborne Cloak (14450, -1.14 DPS) [world_drop] |
| chest | Stone Guard's Satin Armor (220903) | Lady Palanseer [vendor] | sim-verified (93.1 DPS) | yes | Acumen Robes (17775, +0.00 DPS) [quest]; Red Mageweave Vest (10007, -0.20 DPS) [crafted]; Hibernal Robe (8113, -1.13 DPS, sim-verified) [world_drop] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | 58.1 healing_power points (2.74 DPS) | yes | Forgotten Wraps (9433, -0.03 DPS) [world_drop]; Nethergeld Cuffs (254061, -0.16 DPS) [crafted]; Imperial Red Bracers (8247, -0.26 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 127.2 healing_power points (6.00 DPS) | yes | Virtuous Mitts (226950, -1.21 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -2.29 DPS) [crafted]; Greenleaf Handwraps (19116, -2.48 DPS) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 102.9 healing_power points (4.85 DPS) | yes | Serenity Belt (13144, -1.42 DPS, sim-verified) [world_drop]; Gilded Waistcord (254081, -1.45 DPS) [crafted]; Imperial Red Sash (8253, -1.47 DPS) [world_drop] |
| legs | Stone Guard's Satin Leggings (220902) | Lady Palanseer [vendor] | sim-verified (93.1 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.38 DPS, sim-verified) [dungeon]; Venomshroud Leggings (14444, -0.75 DPS) [world_drop]; Imperial Red Pants (8251, -0.88 DPS) [world_drop] |
| feet | First Sergeant's Satin Boots (220900) | PvP rank 9 · First Sergeant · Horde [vendor] | 77.4 healing_power points (3.65 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Highborne Footpads (14447, -0.94 DPS) [world_drop]; Vinerot Sandals (17748, -0.95 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 35.0 healing_power points (1.65 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Mindbender Loop (5009, -0.05 DPS) [world_drop]; Eye of Adaegus (5266, -0.17 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 34.1 healing_power points (1.61 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.01 DPS) [world_drop]; Eye of Adaegus (5266, -0.13 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (93.1 DPS) | yes | Ankh of Life (1713, -0.83 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.85 DPS) [quest]; Alchemist's Stone (13503, -0.91 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (93.1 DPS) | yes | Evonice's Landin' Pilla (18951, -0.74 DPS) [quest]; Alchemist's Stone (13503, -0.81 DPS) [crafted]; Ankh of Life (1713, -1.28 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (93.1 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Spellshifter Rod (9527, -0.34 DPS) [quest]; Glowing Brightwood Staff (812, -1.44 DPS) [world_drop] |
| off_hand | Mistscape Stave (7611) | World drop [world_drop] | 43.0 healing_power points (2.03 DPS) | yes | Witch's Finger (16887, -0.23 DPS) [quest]; Truesilver Conduit (249455, -0.23 DPS) [crafted]; Flask of Forest Mojo (19115, -0.25 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 29.4 healing_power points (1.39 DPS) | yes | Nature's Breath (19118, -0.03 DPS) [quest]; Dancing Flame (6806, -0.26 DPS) [quest]; Goblin Igniter (5253, -0.45 DPS) [quest] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Stone Guard's Satin Armor; wrist: Shizzle's Nozzle Wiper; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Stone Guard's Satin Leggings; feet: First Sergeant's Satin Boots; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; off_hand: Mistscape Stave

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 369.3. Weights run: 2.1s. Verify run: 19.6s. 926 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.161, intellect=1.080 ± 0.069, spirit=0.442 ± 0.099, mp5=0.658 ± 0.068, crit=0.165 ± 0.030 per rating point (14 rating = 1%, 2.316 per %), spell_haste=-7.216 ± 1.412

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 60.8 healing_power points (13.39 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -15.69 DPS, sim-verified) [crafted] |
| neck | Drake Tooth Necklace (21531) | The Nightmare Manifests [quest] | 39.0 healing_power points (8.59 DPS) | yes | Wavefront Necklace (20685, +0.00 DPS, sim-verified) [world]; The Eye of Zuldazar (19593, -1.45 DPS) [quest]; The All-Seeing Eye of Zuldazar (19594, -1.45 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 67.4 healing_power points (14.85 DPS) | yes | Warlord's Satin Mantle (231631, -3.48 DPS) [pvp]; Champion's Satin Mantle (227120, -3.80 DPS) [pvp]; Mooncloth Shoulders (14139, -12.31 DPS, sim-verified) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.8 healing_power points (11.64 DPS) | yes | Drape of Recovery (272413, -2.83 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19526, -5.13 DPS) [rep] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 80.1 healing_power points (17.65 DPS) | yes | Warlord's Satin Tunic (231632, -1.07 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -3.90 DPS) [pvp]; Virtuous Robe (226945, -3.94 DPS, sim-verified) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nethergeld Cuffs (254061, -0.62 DPS) [crafted]; Earthenweave Cuffs (254131, -1.36 DPS) [crafted]; Bracers of Hope (22667, -9.96 DPS, sim-verified) [quest] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 58.6 healing_power points (12.92 DPS) | yes | Raider Handwraps (272097, -1.66 DPS, sim-verified) [vendor]; Mooncloth Gloves (18409, -1.84 DPS) [crafted] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Waistcord (254081, -0.43 DPS) [crafted]; Devout Belt (16696, -0.90 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -10.35 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 56.2 healing_power points (12.39 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Devout Skirt (16694, -12.36 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 65.3 healing_power points (14.39 DPS) | yes | General's Satin Walkers (231630, -3.72 DPS) [vendor]; Virtuous Sandals (226952, -4.65 DPS) [quest]; Mooncloth Boots (15802, -6.33 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 48.0 healing_power points (10.58 DPS) | yes | Blessed Band of Light (272407, -2.97 DPS) [vendor]; Channeler's Ring (272406, -6.47 DPS) [vendor]; Sea Giant's Toe Ring (274746, -6.61 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 35.4 healing_power points (7.81 DPS) | yes | Blessed Band of Light (272407, -1.69 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.70 DPS) [vendor]; Sea Giant's Toe Ring (274746, -3.84 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -1.01 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.20 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (369.2 DPS) | yes | Uther's Strength (11302, -5.29 DPS) [world_drop]; Ankh of Life (1713, -8.09 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -8.89 DPS, sim-verified) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -6.98 DPS) [dungeon]; Spellshifter Rod (9527, -11.10 DPS) [quest] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 43.3 healing_power points (9.54 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Beacon of Hope (9393, -6.07 DPS) [dungeon]; Twisting Essence Jar (249456, -7.52 DPS, sim-verified) [crafted] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 21.9 healing_power points (4.82 DPS) | yes | Brilliant Wand (249385, -3.15 DPS) [crafted]; Captain Rackmore's Tiller (16789, -3.20 DPS) [quest]; Jaina's Firestarter (13064, -5.10 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Drake Tooth Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 926, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 745.5. Weights run: 1.8s. Verify run: 11.6s. 926 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.417, intellect=1.190 ± 0.059, spirit=1.056 ± 0.060, mp5=2.129 ± 0.068, crit=0.443 ± 0.043 per rating point (14 rating = 1%, 6.199 per %), spell_haste=not significant (-0.487 ± 1.301)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Virtuous Crown (226947) | Saving the Best for Last [quest] | 76.3 healing_power points (21.94 DPS) | yes | Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, -0.04 DPS) [pvp]; Mooncloth Circlet (14140, -25.29 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 48.2 healing_power points (13.85 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, -1.99 DPS) [quest]; Drake Tooth Necklace (21531, -2.27 DPS) [quest]; The Eye of Zuldazar (19593, -2.83 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 82.3 healing_power points (23.67 DPS) | yes | Champion's Satin Mantle (227120, -6.40 DPS) [pvp]; Warlord's Satin Mantle (231631, -8.20 DPS) [pvp]; Virtuous Mantle (226951, -16.41 DPS, sim-verified) [quest] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.9 healing_power points (15.49 DPS) | yes | Battle Healer's Cloak (19526, -5.59 DPS) [rep]; Drape of Recovery (272413, -6.15 DPS, sim-verified) [vendor] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 89.9 healing_power points (25.84 DPS) | yes | Warlord's Satin Tunic (231632, -1.75 DPS) [vendor]; Legionnaire's Satin Tunic (227124, -4.96 DPS) [pvp]; Virtuous Robe (226945, -6.66 DPS, sim-verified) [quest] |
| wrist | Virtuous Bracers (226949) | An Earnest Proposition [quest] | sim-verified (745.6 DPS) | yes | Nethergeld Cuffs (254061, -1.93 DPS) [crafted]; General's Satin Bracers (17619, -2.06 DPS) [pvp]; Bracers of Hope (22667, -18.42 DPS, sim-verified) [quest] |
| hands | Desert Bloom Gloves (20717) | Armaments of War [quest] | 62.2 healing_power points (17.89 DPS) | yes | Virtuous Mitts (226950, -1.76 DPS) [vendor]; Mooncloth Gloves (18409, -2.93 DPS) [crafted]; Raider Handwraps (272097, -3.51 DPS, sim-verified) [vendor] |
| waist | Virtuous Belt (226948) | Just Compensation [quest] | sim-verified (745.6 DPS) | yes | Penitent's Cinch (272394, -1.17 DPS) [vendor]; Devout Belt (16696, -2.33 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -16.59 DPS, sim-verified) [crafted] |
| legs | Virtuous Skirt (226946) | Anthion's Parting Words [quest] | 69.1 healing_power points (19.85 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Devout Skirt (16694, -21.17 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 75.4 healing_power points (21.69 DPS) | yes | Mooncloth Boots (15802, -6.05 DPS) [crafted]; Virtuous Sandals (226952, -6.76 DPS, sim-verified) [quest]; General's Satin Walkers (231630, -7.19 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 56.4 healing_power points (16.20 DPS) | yes | Blessed Band of Light (272407, -4.86 DPS) [vendor]; Channeler's Ring (272406, -8.98 DPS) [vendor]; Band of the Hierophant (13096, -9.44 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 42.4 healing_power points (12.19 DPS) | yes | Blessed Band of Light (272407, -0.85 DPS) [vendor]; Channeler's Ring (272406, -4.97 DPS) [vendor]; Band of the Hierophant (13096, -5.43 DPS) [world_drop] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (745.6 DPS) | yes | Uther's Strength (11302, -6.90 DPS) [world_drop]; Shard of the Splithooves (10659, -8.40 DPS) [quest]; Ankh of Life (1713, -8.43 DPS) [world_drop] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (745.6 DPS) | yes | Shard of the Splithooves (10659, -5.51 DPS) [quest]; Ankh of Life (1713, -5.54 DPS) [world_drop]; Uther's Strength (11302, -5.74 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (745.6 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -8.88 DPS) [dungeon]; Spellshifter Rod (9527, -14.94 DPS) [quest] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 50.6 healing_power points (14.56 DPS) | yes | High Warlord's Tome of Mending (234564, -1.43 DPS) [pvp]; Lapidis Tankard of Tidesippe (4696, -7.99 DPS) [world_drop]; Twisting Essence Jar (249456, -10.99 DPS, sim-verified) [crafted] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 28.6 healing_power points (8.23 DPS) | yes | Captain Rackmore's Tiller (16789, -5.59 DPS) [quest]; Brilliant Wand (249385, -5.83 DPS) [crafted]; Jaina's Firestarter (13064, -7.57 DPS, sim-verified) [world_drop] |

**New at 60:** head: Virtuous Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Virtuous Bracers; hands: Desert Bloom Gloves; waist: Virtuous Belt; legs: Virtuous Skirt; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Serenity Field; trinket2: Darkspear Voodoo Seal; off_hand: Lei of the Lifegiver; ranged: Torch of Light

No-known-source sample (15 of 926, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

