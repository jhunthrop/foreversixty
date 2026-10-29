# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 25.7. Weights run: 5.5s. Verify run: 19.0s. 219 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.206 ± 0.002, spirit=0.284 ± 0.002, mp5=1.832 ± 0.023, crit=0.040 ± 0.002 per rating point (14 rating = 1%, 0.562 per %), spell_haste=not significant (0.001 ± 0.001)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.65 DPS) | yes | Pristine Circlet (253949, -0.31 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.32 DPS) [crafted]; Trapper's Leather Hood (252505, -0.89 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 1.1 healing_power points (0.09 DPS) | yes | Scholarly Pendant (277203, -0.07 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 healing_power points (0.86 DPS) | yes | Slime-encrusted Pads (6461, -0.42 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.48 DPS) [crafted]; Forest Leather Mantle (4709, -0.77 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 9.6 healing_power points (0.75 DPS) | yes | Pearl-clasped Cloak (5542, -0.47 DPS) [crafted]; Seer's Cape (6378, -0.52 DPS) [dungeon]; Sanguine Cape (14376, -0.62 DPS, sim-verified) [world_drop] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 22.0 healing_power points (1.74 DPS) | yes | Wisdom's Leather Armor (252493, -0.12 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.17 DPS) [crafted]; Corsair's Overshirt (5202, -0.69 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.5 healing_power points (0.91 DPS) | yes | Bright Bracers (3647, -0.53 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.53 DPS) [vendor]; Owl Bracers (4796, -0.76 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 16.0 healing_power points (1.26 DPS) | yes | Pristine Gloves (253913, -0.11 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.25 DPS, sim-verified) [crafted]; Tomb Robber's Gloves (280096, -0.69 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.8 healing_power points (1.25 DPS) | yes | Acolyte's Chain Belt (250516, -0.21 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.24 DPS) [crafted]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 29.4 healing_power points (2.32 DPS) | yes | Wisdom's Leather Pants (252503, -0.14 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.24 DPS) [crafted]; Dreamer's Leggings (270016, -1.27 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 17.0 healing_power points (1.34 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.25 DPS) [crafted]; Black Whelp Slippers (252424, -0.25 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.2 healing_power points (0.57 DPS) | yes | Volcanic Rock Ring (12053, -0.29 DPS) [world_drop]; Minor Channeling Ring (1449, -0.38 DPS) [quest]; Band of Purification (12996, -0.44 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 3.7 healing_power points (0.29 DPS) | yes | Minor Channeling Ring (1449, -0.10 DPS) [quest]; Volcanic Rock Ring (12053, -0.13 DPS, sim-verified) [world_drop]; Band of Purification (12996, -0.15 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 10.9 healing_power points (0.86 DPS) | yes | Verigan's Fist (6953, +0.00 DPS, sim-verified) [quest]; Impaling Harpoon (5200, -0.38 DPS) [dungeon]; Monastic Hammer (270005, -0.67 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Lorekeeper's Ring; main_hand: Sword of the Fallen

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (human, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 43.5. Weights run: 5.5s. Verify run: 20.6s. 418 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.397 ± 0.009, spirit=0.578 ± 0.003, mp5=2.401 ± 0.009, crit=0.117 ± 0.007 per rating point (14 rating = 1%, 1.631 per %), spell_haste=not significant (0.041 ± 0.012)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.5 healing_power points (2.96 DPS) | yes | Enduring Cap (3020, -0.82 DPS) [world_drop]; Acolyte's Chain Helm (250501, -0.93 DPS) [crafted]; Filigreed Pristine Circlet (253975, -1.17 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (43.5 DPS) | yes | Scorn's Icy Choker (23169, +0.00 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, +0.00 DPS) [world]; Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 15.4 healing_power points (1.25 DPS) | yes | Nightsky Mantle (4718, -0.11 DPS) [world_drop]; Mantle of Honor (3560, -0.12 DPS) [quest]; Talbar Mantle (10657, -0.18 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 15.3 healing_power points (1.24 DPS) | yes | Prelacy Cape (7004, -0.07 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.20 DPS) [vendor]; Glowing Thresher Cape (6901, -0.30 DPS) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.4 healing_power points (2.87 DPS) | yes | Pristine Gown (253961, -0.27 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -0.29 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.45 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.2 healing_power points (1.15 DPS) | yes | Glowing Magical Bracelets (13106, -0.24 DPS) [world_drop]; Nightsky Wristbands (6407, -0.33 DPS) [world_drop]; Technician's Bracers (270042, -0.47 DPS) [quest] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.56 DPS) | yes | Acolyte's Gloves (250511, -0.18 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.29 DPS) [crafted]; Pristine Gloves (253913, -0.32 DPS) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 31.8 healing_power points (2.58 DPS) | yes | Mender's Leather Belt (252523, -0.08 DPS) [crafted]; Highlander's Lizardhide Girdle (20105, -1.22 DPS) [rep]; Highlander's Mail Girdle (20120, -1.22 DPS) [vendor] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 36.2 healing_power points (2.93 DPS) | yes | Wisdom's Leather Leggings (252519, -0.11 DPS) [crafted]; Pristine Leggings (253987, -0.12 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.1 healing_power points (2.60 DPS) | yes | Acolyte's Boots (250506, -1.14 DPS) [crafted]; Wisdom's Leather Boots (252444, -1.14 DPS) [crafted]; Nimbus Boots (6998, -1.22 DPS) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.46 DPS) | yes | Runemender's Seal (273031, -0.65 DPS) [world]; Darkspear Signet (272071, -0.68 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.78 DPS) [dungeon] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.0 healing_power points (0.82 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Darkspear Signet (272071, -0.04 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pysan's Old Greatsword (1975, -3.98 DPS) [dungeon]; Mechanic's Pipehammer (9604, -4.16 DPS) [quest]; Verigan's Fist (6953, -4.37 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.5 healing_power points (1.34 DPS) | yes | Eye of Paleth (2943, -0.28 DPS) [quest]; Orb of Mistmantle (13031, -0.42 DPS) [world_drop]; Heart of Agamaggan (6694, -0.52 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 418, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 40 (human, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 73.7. Weights run: 9.0s. Verify run: 38.3s. 580 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.043, intellect=1.723 ± 0.014, spirit=0.891 ± 0.007, mp5=3.195 ± 0.011, crit=0.172 ± 0.011 per rating point (14 rating = 1%, 2.409 per %), spell_haste=not significant (0.090 ± 0.036)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 51.3 healing_power points (3.85 DPS) | yes | Whitemane's Chapeau (7720, -0.78 DPS) [dungeon]; Holy Shroud (2721, -0.97 DPS) [world_drop]; Enduring Cap (3020, -1.32 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.8 healing_power points (2.01 DPS) | yes | Necklace of Calisea (1714, -0.64 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.11 DPS) [vendor]; Kaleidoscope Chain (13084, -1.23 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 37.4 healing_power points (2.81 DPS) | yes | Sheepshear Mantle (13115, -0.34 DPS) [world_drop]; Earthen Silk Shoulders (254033, -0.92 DPS) [crafted]; Mistscape Mantle (4734, -1.05 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-verified (73.7 DPS) | yes | Darkspear Raider's Cloak (272077, +0.00 DPS) [vendor]; Blackforge Cape (6424, -0.38 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.1 healing_power points (4.89 DPS) | yes | Deathchill Armor (10764, -1.75 DPS) [dungeon]; Death Speaker Robes (6682, -1.97 DPS) [dungeon]; Pristine Gown (253961, -2.22 DPS) [crafted] |
| wrist | Reflective Wristguards (274750) | Rettrick [vendor] | 22.8 healing_power points (1.71 DPS) | yes | Mindthrust Bracers (1974, -0.34 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -0.52 DPS) [world_drop]; Windchaser Cuffs (14429, -0.55 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 36.5 healing_power points (2.74 DPS) | yes | Gilded Handwraps (254021, -0.18 DPS) [crafted]; Prefect's Gauntlet (250569, -0.26 DPS) [crafted]; Stormcloth Gloves (10011, -0.36 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.1 healing_power points (3.09 DPS) | yes | Mender's Leather Belt (252523, -0.48 DPS) [crafted]; Prefect's Belt (250559, -0.53 DPS) [crafted]; Highlander's Lizardhide Girdle (20104, -1.15 DPS) [rep] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | sim-verified (73.7 DPS) | yes | Warchief Kilt (7760, +0.00 DPS) [dungeon]; Wisdom's Leather Leggings (252519, +0.00 DPS) [crafted]; Pristine Leggings (253987, -0.02 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 43.5 healing_power points (3.27 DPS) | yes | Mender's Leather Shoes (252533, -0.09 DPS) [crafted]; Mender's Mail Boots (252565, -0.09 DPS) [crafted]; Gilded Slippers (254001, -0.59 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (73.7 DPS) | yes | Welken Ring (5011, -0.24 DPS) [world_drop]; Ogremind Ring (1993, -0.25 DPS) [world_drop]; Voodoo Band (1996, -0.25 DPS) [world] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | sim-verified (73.7 DPS) | yes | Welken Ring (5011, -0.32 DPS) [world_drop]; Ogremind Ring (1993, -0.33 DPS) [world_drop]; Voodoo Band (1996, -0.33 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -2.49 DPS) [dungeon]; Mograine's Might (7723, -5.53 DPS) [dungeon]; Mechanic's Pipehammer (9604, -6.04 DPS) [quest] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 29.2 healing_power points (2.20 DPS) | yes | Orb of Souls (249395, -0.82 DPS) [crafted]; Beacon of Hope (9393, -0.88 DPS) [dungeon]; Ravager's Shield (14777, -0.89 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Reflective Wristguards; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 580, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

### Band 50 (human, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 107.3. Weights run: 9.3s. Verify run: 46.2s. 740 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.029, intellect=2.209 ± 0.021, spirit=1.518 ± 0.012, mp5=5.076 ± 0.017, crit=0.309 ± 0.017 per rating point (14 rating = 1%, 4.332 per %), spell_haste=not significant (0.127 ± 0.047)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Imbued Helmet (220810) | Captain Dirgehammer [vendor] | 85.3 healing_power points (5.71 DPS) | yes | Winged Helm (13112, -1.29 DPS) [world_drop]; Papal Fez (9431, -1.34 DPS) [dungeon]; Gemburst Circlet (10751, -2.52 DPS, sim-verified) [quest] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 43.7 healing_power points (2.93 DPS) | yes | Glowing Eye of Mordresh (10769, -0.87 DPS) [dungeon]; Horizon Choker (13085, -0.91 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.18 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Imbued Pauldrons (220808) | Captain Dirgehammer [vendor] | 66.6 healing_power points (4.46 DPS) | yes | Mender's Leather Shoulders (252538, -0.81 DPS) [crafted]; Mender's Mail Shoulders (252569, -0.81 DPS) [crafted]; Ironfeather Shoulders (15067, -2.16 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 38.5 healing_power points (2.58 DPS) | yes | Caretaker's Cape (19531, +0.00 DPS, sim-verified) [rep]; Imperial Red Cloak (8248, -0.55 DPS) [world_drop]; Highborne Cloak (14450, -0.84 DPS) [world_drop] |
| chest | Knight's Imbued Armor (220813) | Captain Dirgehammer [vendor] | 78.0 healing_power points (5.22 DPS) | yes | Stormcloth Vest (10020, -0.53 DPS) [crafted]; Earth Warder's Vest (21311, -1.26 DPS) [quest]; Ghostweave Vest (14141, -1.54 DPS, sim-verified) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 43.6 healing_power points (2.92 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Soulforge Fists (226982) | Mokvar [vendor] | 75.6 healing_power points (5.06 DPS) | yes | Raider Handwraps (272098, -0.34 DPS) [vendor]; Mender's Leather Gauntlets (252551, -0.74 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.74 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 56.5 healing_power points (3.78 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.17 DPS) [dungeon]; Prefect's Waistguard (250574, -0.30 DPS) [crafted] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 78.0 healing_power points (5.22 DPS) | yes | Senior Designer's Pantaloons (11841, -0.69 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.74 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -2.08 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 52.3 healing_power points (3.50 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.09 DPS) [crafted]; Mender's Leather Shoes (252533, -0.15 DPS) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 35.5 healing_power points (2.38 DPS) | yes | Welken Ring (5011, -0.93 DPS) [world_drop]; Lorekeeper's Ring (19523, -1.02 DPS) [rep]; Coldwater Ring (4550, -1.02 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 31.5 healing_power points (2.11 DPS) | yes | Welken Ring (5011, -0.55 DPS, sim-verified) [world_drop]; Lorekeeper's Ring (19523, -0.75 DPS) [rep]; Coldwater Ring (4550, -0.75 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (107.3 DPS) | yes | Ankh of Life (1713, -3.88 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.08 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.28 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (107.3 DPS) | yes | Evonice's Landin' Pilla (18951, -0.19 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.39 DPS) [quest]; Ankh of Life (1713, -0.88 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (107.3 DPS) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.55 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.09 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 34.1 healing_power points (2.28 DPS) | yes | Twisting Essence Jar (249456, -0.20 DPS) [crafted]; Protector Buckler (14790, -0.30 DPS) [world_drop]; Mountainside Buckler (13082, -0.32 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Imbued Helmet; neck: Darkmoon Necklace; shoulder: Knight-Lieutenant's Imbued Pauldrons; back: Darkspear Raider's Cloak; chest: Knight's Imbued Armor; wrist: Mender's Leather Bracers; hands: Soulforge Fists; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Mender's Leather Boots; finger1: Darkspear Signet; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength

No-known-source sample (15 of 740, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

### Band 60 (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 204.4. Weights run: 9.1s. Verify run: 113.8s. 1396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.102, intellect=4.366 ± 0.043, spirit=1.713 ± 0.015, mp5=6.525 ± 0.031, crit=0.415 ± 0.025 per rating point (14 rating = 1%, 5.814 per %), spell_haste=not significant (0.004 ± 0.075)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (204.5 DPS) | yes | Black Dragonscale Helm (252605, -0.24 DPS) [crafted]; Crown of the Penitent (13216, -0.46 DPS) [quest]; Gnomish Turban of Psychic Might (21517, -2.47 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 102.4 healing_power points (6.33 DPS) | yes | Stormpike Sage's Pendant (19098, -2.02 DPS) [rep]; Horizon Choker (13085, -2.13 DPS) [world_drop]; Drake Tooth Necklace (21531, -6.37 DPS, sim-verified) [quest] |
| shoulder | Soulforge Epaulets (226979) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Darkspear Shoulders (272104, +0.00 DPS) [vendor]; Devout Mantle (16695, -8.77 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 85.7 healing_power points (5.29 DPS) | yes | Faded Hakkari Cloak (20218, -0.72 DPS) [quest]; Darkspear Raider's Cloak (272076, -0.99 DPS) [vendor] |
| chest | Knight-Captain's Lamellar Chestplate (227151) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devout Robe (16690, -0.15 DPS) [dungeon]; Field Marshal's Lamellar Chestplate (231641, -0.36 DPS) [pvp]; Breastplate of Salvation (250601, -10.50 DPS, sim-verified) [crafted] |
| wrist | Soulforge Bindings (226977) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magiskull Cuffs (13107, -0.12 DPS) [world_drop]; Magister's Bindings (16683, -0.56 DPS) [dungeon]; Bracers of Hope (22667, -7.62 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 141.9 healing_power points (8.77 DPS) | yes | Soulforge Fists (226982, -2.41 DPS) [vendor]; Mooncloth Gloves (18409, -2.41 DPS) [crafted]; Marshal's Lamellar Gloves (231643, -2.50 DPS) [pvp] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 119.9 healing_power points (7.41 DPS) | yes | Wisdom of the Timbermaw (19047, -0.76 DPS) [crafted]; Elunarian Belt (14465, -0.95 DPS) [world_drop]; Devout Belt (16696, -2.77 DPS, sim-verified) [dungeon] |
| legs | Knight-Captain's Lamellar Legguards (227150) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Dragonscale Leggings (252603, +0.00 DPS) [crafted]; Soulforge Leggings (226980, -1.58 DPS) [vendor]; Martyr's Legplates (250600, -6.82 DPS, sim-verified) [crafted] |
| feet | Soulforge Treads (226983) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Lamellar Greaves (227153, -0.62 DPS) [vendor]; Mooncloth Boots (15802, -1.03 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -6.58 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 106.9 healing_power points (6.61 DPS) | yes | Channeler's Ring (272406, -1.98 DPS) [vendor]; Band of the Hierophant (13096, -2.75 DPS) [world_drop]; Darkspear Signet (272068, -2.98 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 91.8 healing_power points (5.67 DPS) | yes | Band of the Hierophant (13096, -1.81 DPS) [world_drop]; Darkspear Signet (272068, -2.04 DPS) [vendor]; Channeler's Ring (272406, -6.15 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -3.45 DPS) [vendor]; Ankh of Life (1713, -4.78 DPS) [world_drop]; Uther's Strength (11302, -4.94 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -1.15 DPS) [world_drop]; Uther's Strength (11302, -1.31 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -3.47 DPS) [dungeon]; Grand Marshal's Demolisher (234568, -4.37 DPS) [pvp] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 69.6 healing_power points (4.30 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, -0.37 DPS) [vendor]; Elunarian Sphere (15968, -8.34 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Soulforge Epaulets; back: Hide of the Wild; chest: Knight-Captain's Lamellar Chestplate; wrist: Soulforge Bindings; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Knight-Captain's Lamellar Legguards; feet: Soulforge Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves; main_hand: Hand of Edward the Odd; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

### Band 60, raid preset (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 579.2. Weights run: 5.4s. Verify run: 32.6s. 1396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.354, intellect=1.501 ± 0.060, spirit=0.587 ± 0.057, mp5=3.344 ± 0.075, crit=0.789 ± 0.062 per rating point (14 rating = 1%, 11.051 per %), spell_haste=not significant (0.409 ± 0.889)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 108.2 healing_power points (22.35 DPS) | yes | Lieutenant Commander's Lamellar Helmet (227149, -3.41 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, -5.29 DPS) [vendor]; Soulforge Crown (226981, -9.09 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 59.8 healing_power points (12.34 DPS) | yes | Drake Tooth Necklace (21531, -2.63 DPS, sim-verified) [quest]; Darkmoon Necklace (19303, -6.34 DPS) [vendor]; Stormpike Sage's Pendant (19098, -6.48 DPS) [rep] |
| shoulder | Shimmering Dawnbringer Shoulders (227859) | Argent Quartermaster Hasana [vendor] | 85.1 healing_power points (17.57 DPS) | yes | Argent Elite Shoulders (227888, -2.48 DPS, sim-verified) [vendor]; Lieutenant Commander's Lamellar Pauldrons (227148, -3.51 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, -6.82 DPS) [pvp] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 57.0 healing_power points (11.77 DPS) | yes | Drape of Recovery (272413, -4.22 DPS, sim-verified) [vendor]; Faded Hakkari Cloak (20218, -5.15 DPS) [quest]; Caretaker's Cape (19530, -5.44 DPS) [rep] |
| chest | Knight-Captain's Lamellar Chestplate (227151) | Captain Dirgehammer [vendor] | 88.2 healing_power points (18.22 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, -1.16 DPS) [pvp]; Breastplate of Salvation (250601, -2.84 DPS, sim-verified) [crafted]; Soulforge Embrace (226984, -3.80 DPS) [vendor] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 41.6 healing_power points (8.58 DPS) | yes | Soulforge Bindings (226977, -0.95 DPS, sim-verified) [vendor]; Prefect's Wristguards (250584, -1.35 DPS) [crafted]; Nethergeld Cuffs (254061, -1.68 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 64.5 healing_power points (13.33 DPS) | yes | Mooncloth Gloves (18409, -1.83 DPS, sim-verified) [crafted]; Marshal's Lamellar Gloves (231643, -1.95 DPS) [pvp]; Prefect's Gloves (250554, -2.69 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 57.3 healing_power points (11.83 DPS) | yes | Belt of Tiny Heads (20217, -1.73 DPS) [quest]; Mender's Leather Waistguard (252477, -1.92 DPS) [crafted]; Soulforge Cord (226978, -2.40 DPS, sim-verified) [vendor] |
| legs | Knight-Captain's Lamellar Legguards (227150) | Captain Dirgehammer [vendor] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Lamellar Legguards (231639, -3.03 DPS) [vendor]; Red Dragonscale Leggings (252603, -4.48 DPS) [crafted]; Martyr's Legplates (250600, -8.71 DPS, sim-verified) [crafted] |
| feet | Knight-Lieutenant's Lamellar Greaves (227153) | Captain Dirgehammer [vendor] | sim-verified (+10.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulforge Treads (226983, -1.20 DPS) [vendor]; Mooncloth Boots (15802, -2.29 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -10.72 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 65.2 healing_power points (13.47 DPS) | yes | Blessed Band of Light (272407, -6.10 DPS) [vendor]; Darkspear Signet (272068, -7.26 DPS) [vendor]; Channeler's Ring (272406, -8.16 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 50.4 healing_power points (10.41 DPS) | yes | Blessed Band of Light (272407, -3.23 DPS, sim-verified) [vendor]; Darkspear Signet (272068, -4.19 DPS) [vendor]; Channeler's Ring (272406, -5.10 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+11.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -6.22 DPS) [quest]; Uther's Strength (11302, -6.64 DPS) [world_drop]; Ankh of Life (1713, -8.90 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -4.96 DPS) [world_drop]; Ankh of Life (1713, -7.22 DPS) [world_drop]; Shard of the Splithooves (10659, -7.89 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -3.45 DPS) [dungeon]; Death Speaker Scepter (2816, -9.79 DPS) [dungeon] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | sim-verified (+31.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Lei of the Lifegiver (19312, -31.55 DPS, sim-verified) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Shimmering Dawnbringer Shoulders; back: Hide of the Wild; chest: Knight-Captain's Lamellar Chestplate; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Knight-Captain's Lamellar Legguards; feet: Knight-Lieutenant's Lamellar Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; main_hand: Hand of Edward the Odd; off_hand: Evergreen Shield

No-known-source sample (15 of 1396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

## Horde

### Band 20 (undead, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 25.6. Weights run: 5.5s. Verify run: 19.1s. 199 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.206 ± 0.002, spirit=0.284 ± 0.002, mp5=1.832 ± 0.023, crit=0.040 ± 0.002 per rating point (14 rating = 1%, 0.562 per %), spell_haste=not significant (0.001 ± 0.001)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.65 DPS) | yes | Pristine Circlet (253949, -0.31 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.32 DPS) [crafted]; Trapper's Leather Hood (252505, -0.89 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.9 healing_power points (0.07 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 healing_power points (0.86 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.48 DPS) [crafted]; Slime-encrusted Pads (6461, -0.52 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.77 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 9.6 healing_power points (0.75 DPS) | yes | Pearl-clasped Cloak (5542, -0.47 DPS) [crafted]; Seer's Cape (6378, -0.52 DPS) [dungeon]; Sanguine Cape (14376, -0.56 DPS, sim-verified) [world_drop] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 22.0 healing_power points (1.74 DPS) | yes | Wisdom's Leather Armor (252493, -0.07 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.17 DPS) [crafted]; Corsair's Overshirt (5202, -0.69 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.5 healing_power points (0.91 DPS) | yes | Featherbead Bracers (15452, -0.43 DPS) [quest]; Garrison Cuffs (270003, -0.50 DPS) [quest]; Owl Bracers (4796, -0.87 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 16.0 healing_power points (1.26 DPS) | yes | Pristine Gloves (253913, -0.11 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.21 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.49 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.8 healing_power points (1.25 DPS) | yes | Acolyte's Chain Belt (250516, -0.21 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.24 DPS) [crafted]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 29.4 healing_power points (2.32 DPS) | yes | Wisdom's Leather Pants (252503, -0.06 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.24 DPS) [crafted]; Darkweave Breeches (12987, -1.52 DPS) [world_drop] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 17.0 healing_power points (1.34 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.25 DPS) [crafted]; Black Whelp Slippers (252424, -0.25 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.2 healing_power points (0.57 DPS) | yes | Advisor's Ring (20426, -0.28 DPS) [rep]; Volcanic Rock Ring (12053, -0.29 DPS) [world_drop]; Band of Purification (12996, -0.44 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 healing_power points (0.48 DPS) | yes | Volcanic Rock Ring (12053, -0.19 DPS) [world_drop]; Advisor's Ring (20426, -0.30 DPS, sim-verified) [rep]; Band of Purification (12996, -0.34 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 10.9 healing_power points (0.86 DPS) | yes | Impaling Harpoon (5200, -0.53 DPS, sim-verified) [dungeon]; Samophlange Screwdriver (11854, -0.72 DPS) [quest]; Mug of Muddled Memories (277247, -0.74 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Sword of the Fallen

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 205218 Libram of Discovery; 209614 Insignia of the Alliance; 213513 Libram of Deliverance; 248005 Northshire Hammer

### Band 30 (undead, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 43.7. Weights run: 5.5s. Verify run: 21.5s. 394 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.397 ± 0.009, spirit=0.578 ± 0.003, mp5=2.401 ± 0.009, crit=0.117 ± 0.007 per rating point (14 rating = 1%, 1.631 per %), spell_haste=not significant (0.041 ± 0.012)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.5 healing_power points (2.96 DPS) | yes | Acolyte's Chain Helm (250501, -0.93 DPS) [crafted]; Enduring Cap (3020, -0.99 DPS, sim-verified) [world_drop]; Filigreed Pristine Circlet (253975, -1.17 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) (or Scorn's Icy Choker (23169)) | Shade of the Archmage [world] | 8.4 healing_power points (0.68 DPS) | yes | Scorn's Icy Choker (23169, +0.00 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.04 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.09 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | sim-verified (43.7 DPS) | yes | Nightsky Mantle (4718, -0.01 DPS) [world_drop]; Talbar Mantle (10657, -0.08 DPS) [quest]; Death Speaker Mantle (6685, -0.59 DPS, sim-verified) [dungeon] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 15.3 healing_power points (1.24 DPS) | yes | Darkspear Raider's Cloak (272078, -0.25 DPS, sim-verified) [vendor]; Glowing Thresher Cape (6901, -0.30 DPS) [dungeon]; Resilient Cape (14400, -0.56 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 35.4 healing_power points (2.87 DPS) | yes | Acolyte's Silvered Chain Shirt (250521, -0.29 DPS) [crafted]; Pristine Gown (253961, -0.35 DPS, sim-verified) [crafted]; Wisdom's Leather Tunic (252511, -0.45 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.2 healing_power points (1.15 DPS) | yes | Nightsky Wristbands (6407, -0.33 DPS) [world_drop]; Technician's Bracers (270042, -0.47 DPS) [quest]; Glowing Magical Bracelets (13106, -0.85 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 19.2 healing_power points (1.56 DPS) | yes | Acolyte's Gloves (250511, -0.15 DPS, sim-verified) [crafted]; Wisdom's Leather Gloves (252499, -0.29 DPS) [crafted]; Pristine Gloves (253913, -0.32 DPS) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 31.8 healing_power points (2.58 DPS) | yes | Mender's Leather Belt (252523, +0.00 DPS, sim-verified) [crafted]; Highlander's Mail Girdle (20120, -1.22 DPS) [vendor]; Defiler's Lizardhide Girdle (20172, -1.22 DPS) [rep] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 36.2 healing_power points (2.93 DPS) | yes | Wisdom's Leather Leggings (252519, +0.00 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -0.12 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 32.1 healing_power points (2.60 DPS) | yes | Wisdom's Leather Boots (252444, -1.14 DPS) [crafted]; Glowing Copper Boots (250482, -1.42 DPS) [crafted]; Acolyte's Boots (250506, -1.67 DPS, sim-verified) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.46 DPS) | yes | Runemender's Seal (273031, -0.65 DPS) [world]; Solid Gold Band (279845, -0.66 DPS) [quest]; Darkspear Signet (272071, -0.68 DPS) [vendor] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.0 healing_power points (0.82 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, -0.01 DPS) [quest]; Darkspear Signet (272071, -0.04 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -1.78 DPS, sim-verified) [dungeon]; Pysan's Old Greatsword (1975, -3.98 DPS) [dungeon]; Wolfsbane (267369, -4.38 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.5 healing_power points (1.34 DPS) | yes | Witch's Finger (16887, -0.43 DPS) [quest]; Heart of Agamaggan (6694, -0.52 DPS) [dungeon]; Orb of Mistmantle (13031, -0.68 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Ghostly Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 394, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets

### Band 40 (undead, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 86.8. Weights run: 9.0s. Verify run: 36.3s. 548 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.043, intellect=1.723 ± 0.014, spirit=0.891 ± 0.007, mp5=3.195 ± 0.011, crit=0.172 ± 0.011 per rating point (14 rating = 1%, 2.409 per %), spell_haste=not significant (0.090 ± 0.036)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 51.3 healing_power points (3.85 DPS) | yes | Whitemane's Chapeau (7720, -0.95 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -0.97 DPS) [world_drop]; Enduring Cap (3020, -1.32 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.8 healing_power points (2.01 DPS) | yes | Necklace of Calisea (1714, -0.99 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.11 DPS) [vendor]; Kaleidoscope Chain (13084, -1.23 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 37.4 healing_power points (2.81 DPS) | yes | Sheepshear Mantle (13115, -0.65 DPS, sim-verified) [world_drop]; Earthen Silk Shoulders (254033, -0.92 DPS) [crafted]; Mistscape Mantle (4734, -1.05 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 22.5 healing_power points (1.69 DPS) | yes | Battle Healer's Cloak (19528, +0.00 DPS, sim-verified) [rep]; Blackforge Cape (6424, -0.39 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.1 healing_power points (4.89 DPS) | yes | Death Speaker Robes (6682, -1.97 DPS) [dungeon]; Pristine Gown (253961, -2.22 DPS) [crafted]; Deathchill Armor (10764, -3.22 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (86.8 DPS) | yes | Enchanted Kodo Bracers (13119, -0.18 DPS) [world_drop]; Windchaser Cuffs (14429, -0.20 DPS) [world_drop]; Reflective Wristguards (274750, -1.19 DPS, sim-verified) [vendor] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 36.5 healing_power points (2.74 DPS) | yes | Gilded Handwraps (254021, -0.18 DPS) [crafted]; Prefect's Gauntlet (250569, -0.26 DPS) [crafted]; Stormcloth Gloves (10011, -0.36 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 41.1 healing_power points (3.09 DPS) | yes | Prefect's Belt (250559, -0.53 DPS) [crafted]; Mender's Leather Belt (252523, -0.58 DPS, sim-verified) [crafted]; Highlander's Mail Girdle (20119, -1.15 DPS) [vendor] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 39.3 healing_power points (2.95 DPS) | yes | Warchief Kilt (7760, -0.04 DPS) [dungeon]; Acolyte's Silvered Chain Leggings (250526, -0.04 DPS) [crafted]; Pristine Leggings (253987, -0.06 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 43.5 healing_power points (3.27 DPS) | yes | Mender's Leather Shoes (252533, -0.09 DPS) [crafted]; Mender's Mail Boots (252565, -0.09 DPS) [crafted]; Gilded Slippers (254001, -0.59 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 19.2 healing_power points (1.44 DPS) | yes | Welken Ring (5011, -0.32 DPS) [world_drop]; Ogremind Ring (1993, -0.33 DPS) [world_drop]; Voodoo Band (1996, -0.33 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.35 DPS) | yes | Ogremind Ring (1993, -0.25 DPS) [world_drop]; Voodoo Band (1996, -0.25 DPS) [world]; Welken Ring (5011, -0.47 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+6.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -2.49 DPS) [dungeon]; Skullbreaker (17039, -4.70 DPS) [quest]; Frost Tiger Blade (3854, -15.02 DPS, sim-verified) [crafted] |
| off_hand | Forcestone Buckler (17508) | Compendium of the Fallen [quest] | 19.1 healing_power points (1.44 DPS) | yes | Beacon of Hope (9393, -0.12 DPS) [dungeon]; Ravager's Shield (14777, -0.13 DPS) [world_drop]; Orb of Souls (249395, -13.23 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Wisdom's Leather Leggings; feet: Prefect's Boots; finger1: Darkspear Signet; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Forcestone Buckler

No-known-source sample (15 of 548, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (undead, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 111.8. Weights run: 9.3s. Verify run: 42.2s. 716 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.029, intellect=2.209 ± 0.021, spirit=1.518 ± 0.012, mp5=5.076 ± 0.017, crit=0.309 ± 0.017 per rating point (14 rating = 1%, 4.332 per %), spell_haste=not significant (0.127 ± 0.047)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 73.9 healing_power points (4.94 DPS) | yes | Winged Helm (13112, -0.53 DPS) [world_drop]; Papal Fez (9431, -0.58 DPS) [dungeon]; Stormcloth Headband (10032, -0.84 DPS) [crafted] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 43.7 healing_power points (2.93 DPS) | yes | Horizon Choker (13085, -0.45 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -0.87 DPS) [dungeon]; Necklace of Calisea (1714, -1.18 DPS) [dungeon] |
| shoulder | Mender's Leather Shoulders (252538) | Leatherworking [crafted] | sim-verified (111.8 DPS) | yes | Ironfeather Shoulders (15067, +0.00 DPS) [crafted]; Mender's Mail Shoulders (252569, +0.00 DPS) [crafted]; Nethergeld Shoulders (254049, -0.12 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 38.5 healing_power points (2.58 DPS) | yes | Battle Healer's Cloak (19527, -0.50 DPS) [rep]; Imperial Red Cloak (8248, -0.55 DPS) [world_drop]; Highborne Cloak (14450, -0.84 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ghostweave Vest (14141, +0.00 DPS) [crafted]; Earth Warder's Vest (21311, -0.73 DPS) [quest]; Feathered Breastplate (8349, -0.78 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 43.6 healing_power points (2.92 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Soulforge Fists (226982) | Mokvar [vendor] | 75.6 healing_power points (5.06 DPS) | yes | Raider Handwraps (272098, -0.34 DPS) [vendor]; Mender's Leather Gauntlets (252551, -0.74 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.74 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 56.5 healing_power points (3.78 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.17 DPS) [dungeon]; Prefect's Waistguard (250574, -0.30 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 73.1 healing_power points (4.89 DPS) | yes | Senior Designer's Pantaloons (11841, -0.35 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.41 DPS) [dungeon]; Dalewind Trousers (13008, -1.02 DPS) [world_drop] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 52.3 healing_power points (3.50 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.09 DPS) [crafted]; Mender's Leather Shoes (252533, -0.15 DPS) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 35.5 healing_power points (2.38 DPS) | yes | Welken Ring (5011, -0.93 DPS) [world_drop]; Ring of Subtlety (19038, -1.02 DPS) [quest]; Advisor's Ring (19519, -1.02 DPS) [rep] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 31.5 healing_power points (2.11 DPS) | yes | Welken Ring (5011, -0.66 DPS) [world_drop]; Ring of Subtlety (19038, -0.75 DPS) [quest]; Advisor's Ring (19519, -0.75 DPS) [rep] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -3.88 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.08 DPS) [quest]; Alchemist's Stone (13503, -4.69 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.19 DPS) [quest]; Alchemist's Stone (13503, -0.80 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -2.55 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.09 DPS) [dungeon]; Galgann's Firehammer (9419, -3.49 DPS) [dungeon] |
| off_hand | Mountainside Buckler (13082) | World drop [world_drop] | sim-verified (111.8 DPS) | yes | Protector Buckler (14790, +0.00 DPS) [world_drop]; Twisting Essence Jar (249456, +0.00 DPS) [crafted]; Ravager's Shield (14777, -0.31 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Mender's Leather Shoulders; back: Darkspear Raider's Cloak; wrist: Mender's Leather Bracers; hands: Soulforge Fists; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Mender's Leather Boots; finger1: Darkspear Signet; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; off_hand: Mountainside Buckler

No-known-source sample (15 of 716, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 194.8. Weights run: 9.1s. Verify run: 109.0s. 1418 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.102, intellect=4.366 ± 0.043, spirit=1.713 ± 0.015, mp5=6.525 ± 0.031, crit=0.415 ± 0.025 per rating point (14 rating = 1%, 5.814 per %), spell_haste=not significant (0.004 ± 0.075)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (194.8 DPS) | yes | Black Dragonscale Helm (252605, -0.24 DPS) [crafted]; Crown of the Penitent (13216, -0.46 DPS) [quest]; Gnomish Turban of Psychic Might (21517, -2.70 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 102.4 healing_power points (6.33 DPS) | yes | Frostwolf Advisor's Pendant (19096, -2.02 DPS) [rep]; Horizon Choker (13085, -2.13 DPS) [world_drop]; Drake Tooth Necklace (21531, -5.97 DPS, sim-verified) [quest] |
| shoulder | Soulforge Epaulets (226979) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Darkspear Shoulders (272104, +0.00 DPS) [vendor]; Devout Mantle (16695, -8.17 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 85.7 healing_power points (5.29 DPS) | yes | Faded Hakkari Cloak (20218, -0.72 DPS) [quest]; Darkspear Raider's Cloak (272076, -0.99 DPS) [vendor] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 157.7 healing_power points (9.75 DPS) | yes | Mooncloth Vest (14138, -0.73 DPS) [crafted]; Magister's Robes (16688, -1.34 DPS) [dungeon]; Devout Robe (16690, -3.60 DPS, sim-verified) [dungeon] |
| wrist | Soulforge Bindings (226977) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magiskull Cuffs (13107, -0.12 DPS) [world_drop]; Magister's Bindings (16683, -0.56 DPS) [dungeon]; Bracers of Hope (22667, -7.16 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 141.9 healing_power points (8.77 DPS) | yes | Soulforge Fists (226982, -2.41 DPS) [vendor]; Mooncloth Gloves (18409, -2.41 DPS) [crafted]; Raider Handguards (272101, -3.10 DPS) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 119.9 healing_power points (7.41 DPS) | yes | Wisdom of the Timbermaw (19047, -0.76 DPS) [crafted]; Elunarian Belt (14465, -0.95 DPS) [world_drop]; Devout Belt (16696, -2.64 DPS, sim-verified) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | 165.0 healing_power points (10.20 DPS) | yes | Outrider's Mail Leggings (22676, -1.48 DPS) [rep]; Soulforge Leggings (226980, -2.06 DPS) [vendor]; Red Dragonscale Leggings (252603, -7.91 DPS, sim-verified) [crafted] |
| feet | Soulforge Treads (226983) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, -1.03 DPS) [crafted]; Devout Sandals (16691, -2.49 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -6.23 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 106.9 healing_power points (6.61 DPS) | yes | Channeler's Ring (272406, -1.98 DPS) [vendor]; Band of the Hierophant (13096, -2.75 DPS) [world_drop]; Darkspear Signet (272068, -2.98 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 91.8 healing_power points (5.67 DPS) | yes | Band of the Hierophant (13096, -1.81 DPS) [world_drop]; Darkspear Signet (272068, -2.04 DPS) [vendor]; Channeler's Ring (272406, -5.81 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -3.45 DPS) [vendor]; Ankh of Life (1713, -4.78 DPS) [world_drop]; Uther's Strength (11302, -4.94 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -1.15 DPS) [world_drop]; Uther's Strength (11302, -1.31 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -3.47 DPS) [dungeon]; High Warlord's Destroyer (234546, -4.37 DPS) [pvp] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 69.6 healing_power points (4.30 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, -0.37 DPS) [vendor]; Elunarian Sphere (15968, -7.91 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Soulforge Epaulets; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Soulforge Bindings; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Martyr's Legplates; feet: Soulforge Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves; main_hand: Hand of Edward the Odd; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1418, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 569.4. Weights run: 5.4s. Verify run: 31.6s. 1418 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.354, intellect=1.501 ± 0.060, spirit=0.587 ± 0.057, mp5=3.344 ± 0.075, crit=0.789 ± 0.062 per rating point (14 rating = 1%, 11.051 per %), spell_haste=not significant (0.409 ± 0.889)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 108.2 healing_power points (22.35 DPS) | yes | Gnomish Turban of Psychic Might (21517, -6.53 DPS) [quest]; Sanctified Leather Helm (22689, -6.75 DPS) [quest]; Soulforge Crown (226981, -9.13 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 59.8 healing_power points (12.34 DPS) | yes | Drake Tooth Necklace (21531, -3.07 DPS, sim-verified) [quest]; Darkmoon Necklace (19303, -6.34 DPS) [vendor]; Frostwolf Advisor's Pendant (19096, -6.48 DPS) [rep] |
| shoulder | Shimmering Dawnbringer Shoulders (227859) | Argent Quartermaster Hasana [vendor] | 85.1 healing_power points (17.57 DPS) | yes | Argent Elite Shoulders (227888, -2.61 DPS, sim-verified) [vendor]; Dawnbringer Shoulders (12625, -7.27 DPS) [crafted]; Soulforge Epaulets (226979, -7.54 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 57.0 healing_power points (11.77 DPS) | yes | Drape of Recovery (272413, -3.86 DPS, sim-verified) [vendor]; Faded Hakkari Cloak (20218, -5.15 DPS) [quest]; Battle Healer's Cloak (19526, -5.44 DPS) [rep] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 80.4 healing_power points (16.60 DPS) | yes | Soulforge Embrace (226984, -2.31 DPS, sim-verified) [vendor]; Mooncloth Vest (14138, -2.73 DPS) [crafted]; Devout Robe (16690, -3.21 DPS) [dungeon] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 41.6 healing_power points (8.58 DPS) | yes | Soulforge Bindings (226977, -1.11 DPS, sim-verified) [vendor]; Prefect's Wristguards (250584, -1.35 DPS) [crafted]; Nethergeld Cuffs (254061, -1.68 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 64.5 healing_power points (13.33 DPS) | yes | Mooncloth Gloves (18409, -1.73 DPS, sim-verified) [crafted]; Prefect's Gloves (250554, -2.69 DPS) [crafted]; Soulforge Fists (226982, -2.78 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 57.3 healing_power points (11.83 DPS) | yes | Belt of Tiny Heads (20217, -1.73 DPS) [quest]; Soulforge Cord (226978, -1.81 DPS, sim-verified) [vendor]; Mender's Leather Waistguard (252477, -1.92 DPS) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | 104.6 healing_power points (21.60 DPS) | yes | Soulforge Leggings (226980, -8.05 DPS) [vendor]; Outrider's Mail Leggings (22676, -8.35 DPS) [rep]; Red Dragonscale Leggings (252603, -10.94 DPS, sim-verified) [crafted] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 73.2 healing_power points (15.12 DPS) | yes | Mooncloth Boots (15802, -4.25 DPS) [crafted]; Soulforge Treads (226983, -5.31 DPS, sim-verified) [vendor]; Mender's Leather Boots (252472, -5.93 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 65.2 healing_power points (13.47 DPS) | yes | Blessed Band of Light (272407, -6.10 DPS) [vendor]; Darkspear Signet (272068, -7.26 DPS) [vendor]; Channeler's Ring (272406, -8.16 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 50.4 healing_power points (10.41 DPS) | yes | Blessed Band of Light (272407, -3.32 DPS, sim-verified) [vendor]; Darkspear Signet (272068, -4.19 DPS) [vendor]; Channeler's Ring (272406, -5.10 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+11.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -6.22 DPS) [quest]; Uther's Strength (11302, -6.64 DPS) [world_drop]; Ankh of Life (1713, -8.90 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -4.53 DPS) [quest]; Ankh of Life (1713, -7.22 DPS) [world_drop]; Uther's Strength (11302, -7.70 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -3.45 DPS) [dungeon]; Death Speaker Scepter (2816, -9.79 DPS) [dungeon] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | sim-verified (569.3 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Lei of the Lifegiver (19312, -31.31 DPS, sim-verified) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Shimmering Dawnbringer Shoulders; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Martyr's Legplates; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; main_hand: Hand of Edward the Odd; off_hand: Evergreen Shield

No-known-source sample (15 of 1418, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

