# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 25.0. Weights run: 4.1s. Verify run: 14.1s. 225 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.206 ± 0.002, spirit=0.284 ± 0.002, mp5=1.832 ± 0.023, crit=0.040 ± 0.002 per rating point (14 rating = 1%, 0.562 per %), spell_haste=not significant (0.001 ± 0.001)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.65 DPS) | yes | Pristine Circlet (253949, -0.31 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.32 DPS) [crafted]; Trapper's Leather Hood (252505, -0.89 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 1.1 healing_power points (0.09 DPS) | yes | Scholarly Pendant (277203, -0.04 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 healing_power points (0.86 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.48 DPS) [crafted]; Slime-encrusted Pads (6461, -0.49 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.77 DPS) [world_drop] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 4.8 healing_power points (0.38 DPS) | yes | Pearl-clasped Cloak (5542, -0.06 DPS, sim-verified) [crafted]; Seer's Cape (6378, -0.15 DPS) [dungeon]; Forest Cloak (4710, -0.19 DPS) [world_drop] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 22.0 healing_power points (1.74 DPS) | yes | Wisdom's Leather Armor (252493, -0.01 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.17 DPS) [crafted]; Corsair's Overshirt (5202, -0.69 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.5 healing_power points (0.91 DPS) | yes | Bright Bracers (3647, -0.53 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.62 DPS) [dungeon]; Owl Bracers (4796, -0.65 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 16.0 healing_power points (1.26 DPS) | yes | Wisdom's Leather Gloves (252499, -0.06 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.11 DPS) [crafted]; Tomb Robber's Gloves (280096, -0.69 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.8 healing_power points (1.25 DPS) | yes | Acolyte's Chain Belt (250516, -0.20 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.24 DPS) [crafted]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 29.4 healing_power points (2.32 DPS) | yes | Wisdom's Leather Pants (252503, -0.10 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.24 DPS) [crafted]; Dreamer's Leggings (270016, -1.27 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 17.0 healing_power points (1.34 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.25 DPS) [crafted]; Black Whelp Slippers (252424, -0.25 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.2 healing_power points (0.57 DPS) | yes | Minor Channeling Ring (1449, -0.38 DPS) [quest]; Band of Purification (12996, -0.44 DPS) [world_drop]; The 1 Ring (8350, -0.45 DPS) [world] |
| finger2 | Volcanic Rock Ring (12053) | World drop [world_drop] | 3.6 healing_power points (0.29 DPS) | yes | Minor Channeling Ring (1449, -0.06 DPS, sim-verified) [quest]; Band of Purification (12996, -0.15 DPS) [world_drop]; The 1 Ring (8350, -0.17 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 10.9 healing_power points (0.86 DPS) | yes | Verigan's Fist (6953, +0.00 DPS, sim-verified) [quest]; Impaling Harpoon (5200, -0.38 DPS) [dungeon]; Monastic Hammer (270005, -0.67 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Volcanic Rock Ring; main_hand: Sword of the Fallen

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (human, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 43.4. Weights run: 4.1s. Verify run: 15.4s. 447 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.397 ± 0.009, spirit=0.578 ± 0.003, mp5=2.401 ± 0.009, crit=0.117 ± 0.007 per rating point (14 rating = 1%, 1.631 per %), spell_haste=not significant (0.041 ± 0.012)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.5 healing_power points (2.96 DPS) | yes | Acolyte's Chain Helm (250501, -0.93 DPS) [crafted]; Filigreed Pristine Circlet (253975, -1.17 DPS) [crafted]; Enduring Cap (3020, -1.30 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 8.4 healing_power points (0.68 DPS) | yes | Crystal Starfire Medallion (5003, -0.09 DPS) [world_drop]; Pendant of Myzrael (4614, -0.40 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.61 DPS, sim-verified) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 15.4 healing_power points (1.25 DPS) | yes | Mantle of Honor (3560, -0.12 DPS) [quest]; Talbar Mantle (10657, -0.18 DPS) [quest]; Nightsky Mantle (4718, -1.02 DPS, sim-verified) [world_drop] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 14.5 healing_power points (1.17 DPS) | yes | Repairman's Cape (9605, -0.27 DPS) [quest]; Resilient Cape (14400, -0.49 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.53 DPS, sim-verified) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.4 healing_power points (2.87 DPS) | yes | Acolyte's Silvered Chain Shirt (250521, -0.29 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.45 DPS) [crafted]; Pristine Gown (253961, -1.00 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.2 healing_power points (1.15 DPS) | yes | Nightsky Wristbands (6407, -0.33 DPS) [world_drop]; Technician's Bracers (270042, -0.47 DPS) [quest]; Glowing Magical Bracelets (13106, -1.50 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.56 DPS) | yes | Acolyte's Gloves (250511, -0.14 DPS, sim-verified) [crafted]; Wisdom's Leather Gloves (252499, -0.29 DPS) [crafted]; Pristine Gloves (253913, -0.32 DPS) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 31.8 healing_power points (2.58 DPS) | yes | Mender's Leather Belt (252523, -1.01 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.23 DPS) [crafted]; Silver-lined Belt (13011, -1.37 DPS) [world_drop] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 36.2 healing_power points (2.93 DPS) | yes | Pristine Leggings (253987, -0.12 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.44 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.97 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.1 healing_power points (2.60 DPS) | yes | Wisdom's Leather Boots (252444, -1.14 DPS) [crafted]; Nimbus Boots (6998, -1.22 DPS) [quest]; Acolyte's Boots (250506, -2.37 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.0 healing_power points (0.82 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Defias Mage Ring (1077, -0.22 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 10.0 healing_power points (0.82 DPS) | yes | The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Defias Mage Ring (1077, -0.22 DPS) [world]; Lavishly Jeweled Ring (1156, -0.98 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (43.4 DPS) | yes | Manual Crowd Pummeler (9449, -2.31 DPS, sim-verified) [dungeon]; Pysan's Old Greatsword (1975, -3.98 DPS) [dungeon]; Mechanic's Pipehammer (9604, -4.16 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.5 healing_power points (1.34 DPS) | yes | Orb of Mistmantle (13031, -0.42 DPS) [world_drop]; Heart of Agamaggan (6694, -0.52 DPS) [dungeon]; Eye of Paleth (2943, -1.17 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Prelacy Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 447, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (human, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 66.9. Weights run: 6.8s. Verify run: 26.0s. 636 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.043, intellect=1.723 ± 0.014, spirit=0.891 ± 0.007, mp5=3.195 ± 0.011, crit=0.172 ± 0.011 per rating point (14 rating = 1%, 2.409 per %), spell_haste=not significant (0.090 ± 0.036)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 51.3 healing_power points (3.85 DPS) | yes | Holy Shroud (2721, -0.97 DPS) [world_drop]; Whitemane's Chapeau (7720, -1.01 DPS, sim-verified) [dungeon]; Enduring Cap (3020, -1.32 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.8 healing_power points (2.01 DPS) | yes | Necklace of Calisea (1714, -1.06 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -1.23 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.24 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 37.4 healing_power points (2.81 DPS) | yes | Sheepshear Mantle (13115, -0.66 DPS, sim-verified) [world_drop]; Earthen Silk Shoulders (254033, -0.92 DPS) [crafted]; Mistscape Mantle (4734, -1.05 DPS) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 17.3 healing_power points (1.30 DPS) | yes | Prelacy Cape (7004, -0.07 DPS) [quest]; Sergeant Major's Cape (16336, -0.12 DPS) [pvp]; Glowing Thresher Cape (6901, -0.24 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.1 healing_power points (4.89 DPS) | yes | Death Speaker Robes (6682, -1.97 DPS) [dungeon]; Pristine Gown (253961, -2.22 DPS) [crafted]; Deathchill Armor (10764, -3.15 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 18.2 healing_power points (1.37 DPS) | yes | Windchaser Cuffs (14429, -0.20 DPS) [world_drop]; Earthen Silk Cuffs (254019, -0.28 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.44 DPS, sim-verified) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 36.5 healing_power points (2.74 DPS) | yes | Gilded Handwraps (254021, -0.22 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.26 DPS) [crafted]; Stormcloth Gloves (10011, -0.36 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.1 healing_power points (3.09 DPS) | yes | Prefect's Belt (250559, -0.53 DPS) [crafted]; Mender's Leather Belt (252523, -0.68 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.15 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 39.3 healing_power points (2.95 DPS) | yes | Acolyte's Silvered Chain Leggings (250526, -0.04 DPS) [crafted]; Pristine Leggings (253987, -0.06 DPS) [crafted]; Warchief Kilt (7760, -0.14 DPS, sim-verified) [dungeon] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 43.5 healing_power points (3.27 DPS) | yes | Mender's Mail Boots (252565, -0.09 DPS) [crafted]; Mender's Leather Shoes (252533, -0.16 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -0.59 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 14.9 healing_power points (1.12 DPS) | yes | Ogremind Ring (1993, -0.01 DPS) [world_drop]; Mindbender Loop (5009, -0.08 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.20 DPS) [dungeon] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 14.7 healing_power points (1.11 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.07 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.19 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (66.9 DPS) | yes | Death Speaker Scepter (2816, -2.49 DPS) [dungeon]; Frost Tiger Blade (3854, -2.58 DPS, sim-verified) [crafted]; Mograine's Might (7723, -5.53 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 29.2 healing_power points (2.20 DPS) | yes | Beacon of Hope (9393, -0.88 DPS) [dungeon]; Ravager's Shield (14777, -0.89 DPS) [world_drop]; Orb of Souls (249395, -1.04 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Wisdom's Leather Leggings; feet: Prefect's Boots; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 636, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (human, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 93.2. Weights run: 7.0s. Verify run: 32.8s. 833 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.029, intellect=2.209 ± 0.021, spirit=1.518 ± 0.012, mp5=5.076 ± 0.017, crit=0.309 ± 0.017 per rating point (14 rating = 1%, 4.332 per %), spell_haste=not significant (0.127 ± 0.047)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 73.9 healing_power points (4.94 DPS) | yes | Knight-Lieutenant's Imbued Helmet (220810, +0.00 DPS) [vendor]; Papal Fez (9431, -0.58 DPS) [dungeon]; Winged Helm (13112, -1.02 DPS, sim-verified) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 37.0 healing_power points (2.48 DPS) | yes | Glowing Eye of Mordresh (10769, -0.42 DPS) [dungeon]; Necklace of Calisea (1714, -0.73 DPS) [dungeon]; Mindburst Medallion (11196, -1.08 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 56.3 healing_power points (3.77 DPS) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, +0.00 DPS) [vendor]; Mender's Leather Shoulders (252538, +0.00 DPS, sim-verified) [crafted]; Mender's Mail Shoulders (252569, -0.12 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 30.4 healing_power points (2.03 DPS) | yes | Highborne Cloak (14450, -0.30 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -0.44 DPS) [dungeon]; Spritecaster Cape (11623, -0.54 DPS) [dungeon] |
| chest | Knight's Imbued Armor (220813) | Captain Dirgehammer [vendor] | 78.0 healing_power points (5.22 DPS) | yes | Stormcloth Vest (10020, -0.53 DPS) [crafted]; Earth Warder's Vest (21311, -1.26 DPS) [quest]; Ghostweave Vest (14141, -1.87 DPS, sim-verified) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 43.6 healing_power points (2.92 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 64.5 healing_power points (4.32 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.09 DPS) [crafted]; Prefect's Gloves (250554, -0.35 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 56.5 healing_power points (3.78 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.17 DPS) [dungeon]; Prefect's Waistguard (250574, -0.30 DPS) [crafted] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 78.0 healing_power points (5.22 DPS) | yes | Senior Designer's Pantaloons (11841, -0.69 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.74 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.51 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 52.3 healing_power points (3.50 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.09 DPS) [crafted]; Mender's Leather Shoes (252533, -0.15 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 31.5 healing_power points (2.11 DPS) | yes | Coldwater Ring (4550, -0.75 DPS) [quest]; Ogremind Ring (1993, -0.77 DPS) [world_drop]; Voodoo Band (1996, -0.77 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 21.7 healing_power points (1.45 DPS) | yes | Coldwater Ring (4550, -0.09 DPS) [quest]; Ogremind Ring (1993, -0.11 DPS) [world_drop]; Voodoo Band (1996, -0.11 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (93.3 DPS) | yes | Evonice's Landin' Pilla (18951, -0.19 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.39 DPS) [quest]; Alchemist's Stone (13503, -0.80 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (93.3 DPS) | yes | Evonice's Landin' Pilla (18951, -0.29 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -0.41 DPS) [quest]; Alchemist's Stone (13503, -0.81 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (93.3 DPS) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.55 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.09 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 34.1 healing_power points (2.28 DPS) | yes | Twisting Essence Jar (249456, -0.20 DPS) [crafted]; Protector Buckler (14790, -0.30 DPS) [world_drop]; Mountainside Buckler (13082, -0.32 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; chest: Knight's Imbued Armor; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Mender's Leather Boots; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life

No-known-source sample (15 of 833, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 154.6. Weights run: 6.9s. Verify run: 89.0s. 1755 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.102, intellect=4.366 ± 0.043, spirit=1.713 ± 0.015, mp5=6.525 ± 0.031, crit=0.415 ± 0.025 per rating point (14 rating = 1%, 5.814 per %), spell_haste=not significant (0.004 ± 0.075)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 175.6 healing_power points (10.85 DPS) | yes | Crown of the Penitent (13216, -0.46 DPS) [quest]; Lieutenant Commander's Lamellar Helmet (227149, -1.01 DPS) [vendor]; Black Dragonscale Helm (252605, -7.17 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 68.0 healing_power points (4.20 DPS) | yes | Opaline Medallion (16623, -1.14 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -1.24 DPS) [dungeon]; Archlight Talisman (15856, -1.50 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 118.1 healing_power points (7.30 DPS) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, -0.03 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, -1.25 DPS) [pvp]; Ironfeather Shoulders (15067, -11.03 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 54.9 healing_power points (3.39 DPS) | yes | Sergeant Major's Cape (16337, +0.00 DPS, sim-verified) [pvp]; Highborne Cloak (14450, -0.54 DPS) [world_drop]; Keeper's Cloak (14665, -0.69 DPS) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (154.6 DPS) | yes | Knight-Captain's Lamellar Chestplate (227151, +0.00 DPS) [vendor]; Field Marshal's Lamellar Chestplate (231641, -0.21 DPS) [pvp]; Breastplate of Salvation (250601, -4.10 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (154.6 DPS) | yes | Magister's Bindings (16683, +0.00 DPS) [dungeon]; Indomitable Armguards (14682, -0.13 DPS) [world_drop]; Magiskull Cuffs (13107, -10.33 DPS, sim-verified) [world_drop] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (154.6 DPS) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Mender's Leather Gauntlets (252551, -0.13 DPS) [crafted]; Gilded Gloves (254095, -8.82 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 113.7 healing_power points (7.03 DPS) | yes | Elunarian Belt (14465, -0.57 DPS) [world_drop]; Marshal's Lamellar Belt (16470, -1.02 DPS) [pvp]; Wisdom of the Timbermaw (19047, -7.70 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | 165.0 healing_power points (10.20 DPS) | yes | Knight-Captain's Lamellar Legguards (227150, -0.48 DPS) [vendor]; Devout Skirt (16694, -2.23 DPS) [dungeon]; Red Dragonscale Leggings (252603, -7.75 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (154.6 DPS) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -6.84 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 62.5 healing_power points (3.86 DPS) | yes | Frigid Ring (18679, -1.16 DPS) [world]; Maiden's Circle (13001, -1.34 DPS) [world_drop]; Glowing Crystal Ring (18402, -3.59 DPS, sim-verified) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | sim-verified (154.6 DPS) | yes | Frigid Ring (18679, -0.19 DPS) [world]; Maiden's Circle (13001, -0.37 DPS) [world_drop]; Glowing Crystal Ring (18402, -2.37 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (154.6 DPS) | yes | Evonice's Landin' Pilla (18951, -0.05 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.27 DPS) [quest]; Ankh of Life (1713, -0.90 DPS, sim-verified) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (154.6 DPS) | yes | Evonice's Landin' Pilla (18951, -1.36 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.57 DPS) [quest]; Ankh of Life (1713, -3.68 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (154.6 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -3.47 DPS) [dungeon]; Grand Marshal's Demolisher (234568, -4.37 DPS) [pvp] |
| off_hand | Elunarian Sphere (15968) | World drop [world_drop] | 61.1 healing_power points (3.78 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Arcane Star (15931, -0.23 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves; main_hand: Hand of Edward the Odd; off_hand: Elunarian Sphere

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 501.2. Weights run: 4.1s. Verify run: 31.8s. 1755 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.354, intellect=1.501 ± 0.060, spirit=0.587 ± 0.057, mp5=3.344 ± 0.075, crit=0.789 ± 0.062 per rating point (14 rating = 1%, 11.051 per %), spell_haste=not significant (0.409 ± 0.889)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 108.2 healing_power points (22.35 DPS) | yes | Lieutenant Commander's Lamellar Helmet (227149, -3.41 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, -5.29 DPS) [vendor]; Crown of the Penitent (13216, -13.17 DPS, sim-verified) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.0 healing_power points (5.16 DPS) | yes | Horizon Choker (13085, -0.87 DPS, sim-verified) [world_drop]; Opaline Medallion (16623, -1.13 DPS) [quest]; Archlight Talisman (15856, -2.06 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (501.3 DPS) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Dawnbringer Shoulders (12625, -21.34 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 18.9 healing_power points (3.89 DPS) | yes | Sergeant Major's Cape (16337, -0.45 DPS) [pvp]; Highborne Cloak (14450, -0.62 DPS) [world_drop]; Keeper's Cloak (14665, -0.79 DPS) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (501.3 DPS) | yes | Knight-Captain's Lamellar Chestplate (227151, +0.00 DPS) [vendor]; Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Breastplate of Salvation (250601, -17.49 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (501.3 DPS) | yes | Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Prefect's Wristguards (250584, -19.07 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (501.3 DPS) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Mender's Leather Gauntlets (252551, +0.00 DPS) [crafted]; Prefect's Gloves (250554, -17.59 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (501.3 DPS) | yes | Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -18.48 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | 104.6 healing_power points (21.60 DPS) | yes | Knight-Captain's Lamellar Legguards (227150, -3.06 DPS) [vendor]; Marshal's Lamellar Legguards (231639, -6.09 DPS) [vendor]; Red Dragonscale Leggings (252603, -10.96 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 52.6 healing_power points (10.87 DPS) | yes | Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Marshal's Lamellar Greaves (231646, -1.37 DPS) [vendor]; Mender's Leather Boots (252472, -2.28 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 21.5 healing_power points (4.43 DPS) | yes | Eye of Adaegus (5266, -1.12 DPS) [world_drop]; Frigid Ring (18679, -1.33 DPS) [world]; Maiden's Circle (13001, -1.54 DPS) [world_drop] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 18.0 healing_power points (3.72 DPS) | yes | Frigid Ring (18679, -0.62 DPS) [world]; Maiden's Circle (13001, -0.82 DPS) [world_drop]; Eye of Adaegus (5266, -1.15 DPS, sim-verified) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (501.3 DPS) | yes | Ankh of Life (1713, -2.69 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.93 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.17 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (501.3 DPS) | yes | Evonice's Landin' Pilla (18951, -2.51 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.75 DPS) [quest]; Ankh of Life (1713, -4.07 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (501.3 DPS) | yes | Hand of Righteousness (7721, -3.45 DPS) [dungeon]; Death Speaker Scepter (2816, -9.79 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -34.28 DPS, sim-verified) [world_drop] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | 34.1 healing_power points (7.03 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Orb of Lorica (11262, -37.03 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Glowing Crystal Ring; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Evergreen Shield

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (undead, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 25.1. Weights run: 4.1s. Verify run: 14.0s. 215 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.206 ± 0.002, spirit=0.284 ± 0.002, mp5=1.832 ± 0.023, crit=0.040 ± 0.002 per rating point (14 rating = 1%, 0.562 per %), spell_haste=not significant (0.001 ± 0.001)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.65 DPS) | yes | Pristine Circlet (253949, -0.31 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.32 DPS) [crafted]; Trapper's Leather Hood (252505, -0.89 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.9 healing_power points (0.07 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 healing_power points (0.86 DPS) | yes | Slime-encrusted Pads (6461, -0.38 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.48 DPS) [crafted]; Forest Leather Mantle (4709, -0.77 DPS) [world_drop] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 4.8 healing_power points (0.38 DPS) | yes | Pearl-clasped Cloak (5542, -0.06 DPS, sim-verified) [crafted]; Seer's Cape (6378, -0.15 DPS) [dungeon]; Forest Cloak (4710, -0.19 DPS) [world_drop] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 22.0 healing_power points (1.74 DPS) | yes | Wisdom's Leather Armor (252493, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.17 DPS) [crafted]; Corsair's Overshirt (5202, -0.69 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.5 healing_power points (0.91 DPS) | yes | Featherbead Bracers (15452, -0.43 DPS) [quest]; Garrison Cuffs (270003, -0.50 DPS) [quest]; Owl Bracers (4796, -0.71 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 16.0 healing_power points (1.26 DPS) | yes | Wisdom's Leather Gloves (252499, -0.06 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.11 DPS) [crafted]; Blight Gloves (279877, -0.49 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.8 healing_power points (1.25 DPS) | yes | Acolyte's Chain Belt (250516, -0.21 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.24 DPS) [crafted]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 29.4 healing_power points (2.32 DPS) | yes | Wisdom's Leather Pants (252503, -0.16 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.24 DPS) [crafted]; Darkweave Breeches (12987, -1.52 DPS) [world_drop] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 17.0 healing_power points (1.34 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.25 DPS) [crafted]; Black Whelp Slippers (252424, -0.25 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.2 healing_power points (0.57 DPS) | yes | Volcanic Rock Ring (12053, -0.29 DPS) [world_drop]; Band of Purification (12996, -0.44 DPS) [world_drop]; The 1 Ring (8350, -0.45 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 healing_power points (0.48 DPS) | yes | Volcanic Rock Ring (12053, -0.23 DPS, sim-verified) [world_drop]; Band of Purification (12996, -0.34 DPS) [world_drop]; The 1 Ring (8350, -0.36 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 10.9 healing_power points (0.86 DPS) | yes | Impaling Harpoon (5200, -0.34 DPS, sim-verified) [dungeon]; Samophlange Screwdriver (11854, -0.72 DPS) [quest]; Mug of Muddled Memories (277247, -0.74 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Sword of the Fallen

No-known-source sample (15 of 215, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (undead, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 42.7. Weights run: 4.1s. Verify run: 16.0s. 436 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.397 ± 0.009, spirit=0.578 ± 0.003, mp5=2.401 ± 0.009, crit=0.117 ± 0.007 per rating point (14 rating = 1%, 1.631 per %), spell_haste=not significant (0.041 ± 0.012)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.5 healing_power points (2.96 DPS) | yes | Acolyte's Chain Helm (250501, -0.93 DPS) [crafted]; Filigreed Pristine Circlet (253975, -1.17 DPS) [crafted]; Enduring Cap (3020, -1.25 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 8.4 healing_power points (0.68 DPS) | yes | Crystal Starfire Medallion (5003, -0.09 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.34 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.40 DPS) [dungeon] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 15.4 healing_power points (1.25 DPS) | yes | Ghostly Mantle (3324, +0.00 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.11 DPS) [world_drop]; Talbar Mantle (10657, -0.18 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 11.6 healing_power points (0.94 DPS) | yes | Soft Willow Cape (16661, -0.38 DPS) [quest]; Silver-thread Cloak (4713, -0.40 DPS) [world_drop]; Resilient Cape (14400, -0.65 DPS, sim-verified) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.4 healing_power points (2.87 DPS) | yes | Acolyte's Silvered Chain Shirt (250521, -0.29 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.45 DPS) [crafted]; Pristine Gown (253961, -0.80 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.2 healing_power points (1.15 DPS) | yes | Nightsky Wristbands (6407, -0.33 DPS) [world_drop]; Technician's Bracers (270042, -0.47 DPS) [quest]; Glowing Magical Bracelets (13106, -0.87 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.56 DPS) | yes | Acolyte's Gloves (250511, -0.12 DPS, sim-verified) [crafted]; Wisdom's Leather Gloves (252499, -0.29 DPS) [crafted]; Pristine Gloves (253913, -0.32 DPS) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 31.8 healing_power points (2.58 DPS) | yes | Mender's Leather Belt (252523, -0.26 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.23 DPS) [crafted]; Silver-lined Belt (13011, -1.37 DPS) [world_drop] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 36.2 healing_power points (2.93 DPS) | yes | Pristine Leggings (253987, -0.12 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.27 DPS, sim-verified) [crafted]; Acolyte's Chain Leggings (250496, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.1 healing_power points (2.60 DPS) | yes | Wisdom's Leather Boots (252444, -1.14 DPS) [crafted]; Acolyte's Boots (250506, -1.31 DPS, sim-verified) [crafted]; Glowing Copper Boots (250482, -1.42 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.0 healing_power points (0.82 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; The Queen's Jewel (13094, -0.21 DPS) [world_drop]; Runemender's Seal (273031, -1.29 DPS, sim-verified) [world] |
| finger2 | Solid Gold Band (279845) | Opportunistic Education [quest] | sim-verified (42.7 DPS) | yes | Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon]; The Queen's Jewel (13094, -0.20 DPS) [world_drop]; Runemender's Seal (273031, -0.44 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -1.56 DPS, sim-verified) [dungeon]; Pysan's Old Greatsword (1975, -3.98 DPS) [dungeon]; Wolfsbane (267369, -4.38 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.5 healing_power points (1.34 DPS) | yes | Orb of Mistmantle (13031, -0.08 DPS, sim-verified) [world_drop]; Witch's Finger (16887, -0.43 DPS) [quest]; Heart of Agamaggan (6694, -0.52 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Solid Gold Band; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets

### Band 40 (undead, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 76.8. Weights run: 6.8s. Verify run: 26.1s. 622 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.043, intellect=1.723 ± 0.014, spirit=0.891 ± 0.007, mp5=3.195 ± 0.011, crit=0.172 ± 0.011 per rating point (14 rating = 1%, 2.409 per %), spell_haste=not significant (0.090 ± 0.036)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 51.3 healing_power points (3.85 DPS) | yes | Holy Shroud (2721, -0.97 DPS) [world_drop]; Whitemane's Chapeau (7720, -1.04 DPS, sim-verified) [dungeon]; Enduring Cap (3020, -1.32 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.8 healing_power points (2.01 DPS) | yes | Necklace of Calisea (1714, -0.92 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -1.23 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.24 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 37.4 healing_power points (2.81 DPS) | yes | Sheepshear Mantle (13115, -0.64 DPS, sim-verified) [world_drop]; Earthen Silk Shoulders (254033, -0.92 DPS) [crafted]; Mistscape Mantle (4734, -1.05 DPS) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 17.3 healing_power points (1.30 DPS) | yes | Glowing Thresher Cape (6901, -0.25 DPS, sim-verified) [dungeon]; Windchaser Cloak (14430, -0.26 DPS) [world_drop]; Glyphed Cloak (4732, -0.40 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.1 healing_power points (4.89 DPS) | yes | Death Speaker Robes (6682, -1.97 DPS) [dungeon]; Pristine Gown (253961, -2.22 DPS) [crafted]; Deathchill Armor (10764, -2.90 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 18.2 healing_power points (1.37 DPS) | yes | Windchaser Cuffs (14429, -0.20 DPS) [world_drop]; Earthen Silk Cuffs (254019, -0.28 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.66 DPS, sim-verified) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 36.5 healing_power points (2.74 DPS) | yes | Prefect's Gauntlet (250569, -0.26 DPS) [crafted]; Gilded Handwraps (254021, -0.34 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.36 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.1 healing_power points (3.09 DPS) | yes | Prefect's Belt (250559, -0.53 DPS) [crafted]; Mender's Leather Belt (252523, -0.81 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.15 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 39.3 healing_power points (2.95 DPS) | yes | Warchief Kilt (7760, -0.04 DPS) [dungeon]; Acolyte's Silvered Chain Leggings (250526, -0.04 DPS) [crafted]; Pristine Leggings (253987, -0.06 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 43.5 healing_power points (3.27 DPS) | yes | Mender's Leather Shoes (252533, -0.09 DPS) [crafted]; Mender's Mail Boots (252565, -0.09 DPS) [crafted]; Gilded Slippers (254001, -0.59 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 14.9 healing_power points (1.12 DPS) | yes | Ogremind Ring (1993, -0.01 DPS) [world_drop]; Mindbender Loop (5009, -0.08 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.20 DPS) [dungeon] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 14.7 healing_power points (1.11 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.07 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.19 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (76.8 DPS) | yes | Death Speaker Scepter (2816, -2.49 DPS) [dungeon]; Skullbreaker (17039, -4.70 DPS) [quest]; Frost Tiger Blade (3854, -12.92 DPS, sim-verified) [crafted] |
| off_hand | Forcestone Buckler (17508) | Compendium of the Fallen [quest] | 19.1 healing_power points (1.44 DPS) | yes | Beacon of Hope (9393, -0.12 DPS) [dungeon]; Ravager's Shield (14777, -0.13 DPS) [world_drop]; Orb of Souls (249395, -11.47 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Wisdom's Leather Leggings; feet: Prefect's Boots; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Forcestone Buckler

No-known-source sample (15 of 622, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (undead, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 101.6. Weights run: 7.0s. Verify run: 33.1s. 832 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.029, intellect=2.209 ± 0.021, spirit=1.518 ± 0.012, mp5=5.076 ± 0.017, crit=0.309 ± 0.017 per rating point (14 rating = 1%, 4.332 per %), spell_haste=not significant (0.127 ± 0.047)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 73.9 healing_power points (4.94 DPS) | yes | Winged Helm (13112, -0.53 DPS) [world_drop]; Papal Fez (9431, -0.58 DPS) [dungeon]; Stormcloth Headband (10032, -0.84 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 37.0 healing_power points (2.48 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.73 DPS) [dungeon]; Choker of the High Shaman (4112, -0.92 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 56.3 healing_power points (3.77 DPS) | yes | Mender's Leather Shoulders (252538, +0.00 DPS, sim-verified) [crafted]; Mender's Mail Shoulders (252569, -0.12 DPS) [crafted]; Nethergeld Shoulders (254049, -0.24 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 30.4 healing_power points (2.03 DPS) | yes | Highborne Cloak (14450, -0.36 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -0.44 DPS) [dungeon]; Spritecaster Cape (11623, -0.54 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Earth Warder's Vest (21311, -0.73 DPS) [quest]; Feathered Breastplate (8349, -0.78 DPS) [crafted]; Ghostweave Vest (14141, -2.51 DPS, sim-verified) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 43.6 healing_power points (2.92 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 64.5 healing_power points (4.32 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.09 DPS) [crafted]; Prefect's Gloves (250554, -0.35 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 56.5 healing_power points (3.78 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.17 DPS) [dungeon]; Prefect's Waistguard (250574, -0.30 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 73.1 healing_power points (4.89 DPS) | yes | Senior Designer's Pantaloons (11841, -0.35 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.41 DPS) [dungeon]; Dalewind Trousers (13008, -1.02 DPS) [world_drop] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 52.3 healing_power points (3.50 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.09 DPS) [crafted]; Mender's Leather Shoes (252533, -0.15 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 31.5 healing_power points (2.11 DPS) | yes | Ring of Subtlety (19038, -0.75 DPS) [quest]; Coldwater Ring (4550, -0.75 DPS) [quest]; Voodoo Band (1996, -0.77 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 21.7 healing_power points (1.45 DPS) | yes | Ring of Subtlety (19038, +0.00 DPS, sim-verified) [quest]; Coldwater Ring (4550, -0.09 DPS) [quest]; Voodoo Band (1996, -0.11 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.19 DPS) [quest]; Alchemist's Stone (13503, -1.47 DPS, sim-verified) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.20 DPS) [quest]; Alchemist's Stone (13503, -0.81 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.55 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.09 DPS) [dungeon] |
| off_hand | Protector Buckler (14790) | World drop [world_drop] | sim-verified (+11.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mountainside Buckler (13082, -0.03 DPS) [world_drop]; Ravager's Shield (14777, -0.34 DPS) [world_drop]; Twisting Essence Jar (249456, -11.27 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Mender's Leather Boots; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life; off_hand: Protector Buckler

No-known-source sample (15 of 832, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 151.6. Weights run: 6.9s. Verify run: 89.0s. 1774 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.102, intellect=4.366 ± 0.043, spirit=1.713 ± 0.015, mp5=6.525 ± 0.031, crit=0.415 ± 0.025 per rating point (14 rating = 1%, 5.814 per %), spell_haste=not significant (0.004 ± 0.075)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 175.6 healing_power points (10.85 DPS) | yes | Crown of the Penitent (13216, -0.46 DPS) [quest]; Devout Crown (16693, -1.55 DPS) [dungeon]; Black Dragonscale Helm (252605, -6.26 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 68.0 healing_power points (4.20 DPS) | yes | Glowing Eye of Mordresh (10769, -0.61 DPS, sim-verified) [dungeon]; Archlight Talisman (15856, -1.50 DPS) [quest]; Necklace of Calisea (1714, -1.57 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 118.1 healing_power points (7.30 DPS) | yes | Magister's Mantle (16689, -1.36 DPS) [dungeon]; Blue Dragonscale Shoulders (15049, -1.63 DPS) [crafted]; Ironfeather Shoulders (15067, -9.04 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 54.9 healing_power points (3.39 DPS) | yes | Keeper's Cloak (14665, -0.69 DPS) [world_drop]; Blackforge Cape (6424, -0.81 DPS) [dungeon]; Highborne Cloak (14450, -0.87 DPS, sim-verified) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (151.5 DPS) | yes | Mooncloth Vest (14138, -0.29 DPS) [crafted]; Magister's Robes (16688, -0.90 DPS) [dungeon]; Breastplate of Salvation (250601, -4.70 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (151.5 DPS) | yes | Magister's Bindings (16683, +0.00 DPS) [dungeon]; Indomitable Armguards (14682, -0.13 DPS) [world_drop]; Magiskull Cuffs (13107, -7.92 DPS, sim-verified) [world_drop] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (151.5 DPS) | yes | Mender's Leather Gauntlets (252551, -0.13 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.13 DPS) [crafted]; Gilded Gloves (254095, -6.25 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (151.5 DPS) | yes | Elunarian Belt (14465, -0.57 DPS) [world_drop]; Dawnspire Cord (12466, -1.16 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -6.10 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | 165.0 healing_power points (10.20 DPS) | yes | Devout Skirt (16694, -2.23 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -3.06 DPS) [dungeon]; Red Dragonscale Leggings (252603, -7.58 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (151.5 DPS) | yes | Mender's Leather Boots (252472, -0.48 DPS) [crafted]; Mender's Mail Sabatons (252579, -0.48 DPS) [crafted]; Mooncloth Boots (15802, -4.26 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 62.5 healing_power points (3.86 DPS) | yes | Chemist's Ring (15702, -0.74 DPS) [quest]; Eye of Adaegus (5266, -0.97 DPS) [world_drop]; Frigid Ring (18679, -1.16 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 52.4 healing_power points (3.24 DPS) | yes | Chemist's Ring (15702, -0.12 DPS) [quest]; Eye of Adaegus (5266, -0.35 DPS) [world_drop]; Frigid Ring (18679, -0.54 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (151.5 DPS) | yes | Ankh of Life (1713, -1.15 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.36 DPS) [quest]; Alchemist's Stone (13503, -2.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (151.5 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.05 DPS) [quest]; Alchemist's Stone (13503, -0.69 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (151.5 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -3.47 DPS) [dungeon]; High Warlord's Destroyer (234546, -4.37 DPS) [pvp] |
| off_hand | Elunarian Sphere (15968) | World drop [world_drop] | 61.1 healing_power points (3.78 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Arcane Star (15931, -0.23 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Glowing Crystal Ring; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Elunarian Sphere

No-known-source sample (15 of 1774, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 498.7. Weights run: 4.1s. Verify run: 31.8s. 1774 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.354, intellect=1.501 ± 0.060, spirit=0.587 ± 0.057, mp5=3.344 ± 0.075, crit=0.789 ± 0.062 per rating point (14 rating = 1%, 11.051 per %), spell_haste=not significant (0.409 ± 0.889)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 108.2 healing_power points (22.35 DPS) | yes | Black Dragonscale Helm (252605, -8.29 DPS) [crafted]; Whitesoul Helm (12633, -8.66 DPS) [crafted]; Crown of the Penitent (13216, -12.15 DPS, sim-verified) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.0 healing_power points (5.16 DPS) | yes | Horizon Choker (13085, -0.68 DPS, sim-verified) [world_drop]; Archlight Talisman (15856, -2.06 DPS) [quest]; Necklace of Calisea (1714, -2.15 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (498.7 DPS) | yes | Prefect's Pauldrons (250579, -0.78 DPS) [crafted]; Mooncloth Shoulders (14139, -0.90 DPS) [crafted]; Dawnbringer Shoulders (12625, -21.37 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 18.9 healing_power points (3.89 DPS) | yes | Keeper's Cloak (14665, -0.79 DPS) [world_drop]; Blackforge Cape (6424, -0.93 DPS) [dungeon]; Highborne Cloak (14450, -0.96 DPS, sim-verified) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (498.7 DPS) | yes | Mooncloth Vest (14138, +0.00 DPS) [crafted]; Mooncloth Robe (18486, -0.27 DPS) [crafted]; Breastplate of Salvation (250601, -17.76 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (498.7 DPS) | yes | Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Prefect's Wristguards (250584, -18.69 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (498.7 DPS) | yes | Mender's Leather Gauntlets (252551, +0.00 DPS) [crafted]; Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Prefect's Gloves (250554, -17.47 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (498.7 DPS) | yes | Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -18.13 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | 104.6 healing_power points (21.60 DPS) | yes | Devout Skirt (16694, -9.20 DPS) [dungeon]; Elderwild Pants (279254, -10.43 DPS) [crafted]; Red Dragonscale Leggings (252603, -10.44 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 52.6 healing_power points (10.87 DPS) | yes | Mender's Leather Boots (252472, -1.60 DPS, sim-verified) [crafted]; Mender's Mail Sabatons (252579, -1.68 DPS) [crafted]; Gilded Sandals (254107, -1.68 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 21.5 healing_power points (4.43 DPS) | yes | Chemist's Ring (15702, -0.85 DPS) [quest]; Eye of Adaegus (5266, -1.12 DPS) [world_drop]; Frigid Ring (18679, -1.33 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 18.0 healing_power points (3.72 DPS) | yes | Chemist's Ring (15702, -0.14 DPS) [quest]; Eye of Adaegus (5266, -0.41 DPS) [world_drop]; Frigid Ring (18679, -0.62 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (498.7 DPS) | yes | Ankh of Life (1713, -2.69 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.93 DPS) [quest]; Alchemist's Stone (13503, -3.66 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (498.7 DPS) | yes | Evonice's Landin' Pilla (18951, -2.51 DPS) [quest]; Alchemist's Stone (13503, -3.23 DPS) [crafted]; Ankh of Life (1713, -3.78 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (498.7 DPS) | yes | Hand of Righteousness (7721, -3.45 DPS) [dungeon]; Death Speaker Scepter (2816, -9.79 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -33.75 DPS, sim-verified) [world_drop] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | 34.1 healing_power points (7.03 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Twisting Essence Jar (249456, -37.92 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Glowing Crystal Ring; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Evergreen Shield

No-known-source sample (15 of 1774, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

