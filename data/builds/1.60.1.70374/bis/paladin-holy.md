# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 27.3. Weights run: 1.4s. Verify run: 4.3s. 225 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.901 ± 0.003, spirit=0.222 ± 0.001, mp5=1.699 ± 0.033, crit=0.037 ± 0.002 per rating point (14 rating = 1%, 0.512 per %), spell_haste=0.104 ± 0.016

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.76 DPS) | yes | Wisdom's Leather Hood (252507, -0.34 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.16 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 0.9 healing_power points (0.07 DPS) | yes | Scholarly Pendant (277203, -0.19 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.38 DPS) [crafted]; Slime-encrusted Pads (6461, -0.40 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.61 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 9.4 healing_power points (0.79 DPS) | yes | Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.60 DPS) [dungeon]; Sanguine Cape (14376, -0.83 DPS, sim-verified) [world_drop] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 20.5 healing_power points (1.72 DPS) | yes | Wisdom's Leather Armor (252493, -0.13 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.63 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.6 healing_power points (0.80 DPS) | yes | Bright Bracers (3647, -0.50 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.50 DPS) [vendor]; Owl Bracers (4796, -1.05 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.5 healing_power points (1.22 DPS) | yes | Wisdom's Leather Gloves (252499, -0.08 DPS) [crafted]; Pristine Gloves (253913, -0.23 DPS, sim-verified) [crafted]; Tomb Robber's Gloves (280096, -0.76 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.22 DPS) | yes | Wisdom's Leather Belt (252433, -0.25 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.26 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.44 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 27.3 healing_power points (2.29 DPS) | yes | Filigreed Pristine Leggings (253937, -0.25 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.44 DPS, sim-verified) [crafted]; Dreamer's Leggings (270016, -1.46 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.5 healing_power points (1.30 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.24 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.45 DPS) | yes | Volcanic Rock Ring (12053, -0.23 DPS) [world_drop]; Minor Channeling Ring (1449, -0.30 DPS) [quest]; Band of Purification (12996, -0.34 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 3.4 healing_power points (0.28 DPS) | yes | Minor Channeling Ring (1449, -0.13 DPS) [quest]; Band of Purification (12996, -0.17 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.32 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 8.1 healing_power points (0.68 DPS) | yes | Impaling Harpoon (5200, -0.30 DPS) [dungeon]; Verigan's Fist (6953, -0.33 DPS, sim-verified) [quest]; Monastic Hammer (270005, -0.53 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Lorekeeper's Ring; main_hand: Sword of the Fallen

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (human, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 48.4. Weights run: 1.5s. Verify run: 4.0s. 413 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.985 ± 0.005, spirit=0.496 ± 0.004, mp5=2.104 ± 0.011, crit=0.092 ± 0.005 per rating point (14 rating = 1%, 1.288 per %), spell_haste=not significant (0.037 ± 0.013)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.0 healing_power points (3.37 DPS) | yes | Filigreed Pristine Circlet (253975, -1.31 DPS) [crafted]; Wisdom's Leather Helm (252515, -1.40 DPS) [crafted]; Acolyte's Chain Helm (250501, -1.73 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.9 healing_power points (0.56 DPS) | yes | Scorn's Icy Choker (23169, -0.00 DPS) [dungeon]; Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.16 DPS, sim-verified) [world] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 10.8 healing_power points (1.02 DPS) | yes | Mantle of Honor (3560, +0.00 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.04 DPS) [world_drop]; Talbar Mantle (10657, -0.07 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 15.0 healing_power points (1.41 DPS) | yes | Prelacy Cape (7004, -0.09 DPS) [quest]; Glowing Thresher Cape (6901, -0.38 DPS) [dungeon]; Repairman's Cape (9605, -0.52 DPS) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 30.8 healing_power points (2.89 DPS) | yes | Pristine Gown (253961, -0.18 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.37 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -0.46 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.2 healing_power points (1.05 DPS) | yes | Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.90 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.0 healing_power points (1.68 DPS) | yes | Acolyte's Gloves (250511, -0.28 DPS) [crafted]; Silvered Gauntlets (270025, -0.37 DPS) [quest]; Naga Battle Gloves (888, -0.53 DPS, sim-verified) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 28.9 healing_power points (2.71 DPS) | yes | Mender's Leather Belt (252523, +0.00 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.31 DPS) [crafted]; Acolyte's Chain Belt (250516, -1.50 DPS) [crafted] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 32.9 healing_power points (3.08 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Leggings (252519, -0.14 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.47 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.9 healing_power points (2.71 DPS) | yes | Acolyte's Boots (250506, -1.21 DPS) [crafted]; Wisdom's Leather Boots (252444, -1.21 DPS) [crafted]; Nimbus Boots (6998, -1.96 DPS, sim-verified) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.69 DPS) | yes | Runemender's Seal (273031, -0.80 DPS) [world]; Darkspear Signet (272071, -0.90 DPS) [vendor]; Ring of Power Regulation (273030, -1.10 DPS) [dungeon] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.5 healing_power points (0.89 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Darkspear Signet (272071, -0.10 DPS) [vendor]; Ring of Power Regulation (273030, -0.30 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (48.4 DPS) | yes | Manual Crowd Pummeler (9449, -2.24 DPS, sim-verified) [dungeon]; Pysan's Old Greatsword (1975, -4.59 DPS) [dungeon]; Mechanic's Pipehammer (9604, -5.27 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.0 healing_power points (1.50 DPS) | yes | Orb of Mistmantle (13031, -0.47 DPS) [world_drop]; Heart of Agamaggan (6694, -0.56 DPS) [dungeon]; Eye of Paleth (2943, -0.64 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 413, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 40 (human, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 78.1. Weights run: 1.6s. Verify run: 7.2s. 568 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.164 ± 0.009, spirit=0.728 ± 0.008, mp5=2.684 ± 0.012, crit=0.162 ± 0.009 per rating point (14 rating = 1%, 2.271 per %), spell_haste=not significant (0.080 ± 0.025)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 40.3 healing_power points (3.67 DPS) | yes | Earthen Silk Hood (254015, -1.01 DPS) [crafted]; Holy Shroud (2721, -1.03 DPS, sim-verified) [world_drop]; Acolyte's Chain Helm (250501, -1.40 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.3 healing_power points (2.03 DPS) | yes | Darkspear Warding Pendant (272074, -1.29 DPS) [vendor]; Kaleidoscope Chain (13084, -1.34 DPS) [world_drop]; Necklace of Calisea (1714, -1.43 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 30.1 healing_power points (2.74 DPS) | yes | Sheepshear Mantle (13115, -0.65 DPS) [world_drop]; Mistscape Mantle (4734, -1.25 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.38 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | 21.6 healing_power points (1.97 DPS) | yes | Darkspear Raider's Cloak (272077, -0.54 DPS) [vendor]; Prelacy Cape (7004, -0.57 DPS) [quest]; Glowing Thresher Cape (6901, -0.80 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.8 healing_power points (5.81 DPS) | yes | Death Speaker Robes (6682, -2.83 DPS) [dungeon]; Pristine Gown (253961, -2.99 DPS) [crafted]; Deathchill Armor (10764, -4.32 DPS, sim-verified) [dungeon] |
| wrist | Reflective Wristguards (274750) | Rettrick [vendor] | 18.3 healing_power points (1.67 DPS) | yes | Mindthrust Bracers (1974, -0.40 DPS) [dungeon]; Earthen Silk Cuffs (254019, -0.43 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.58 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 31.5 healing_power points (2.87 DPS) | yes | Gilded Handwraps (254021, -0.12 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.21 DPS) [crafted]; Gloves of the Greatfather (17721, -0.50 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.7 healing_power points (3.25 DPS) | yes | Prefect's Belt (250559, -0.50 DPS) [crafted]; Mender's Leather Belt (252523, -0.73 DPS, sim-verified) [crafted]; Highlander's Lizardhide Girdle (20104, -1.66 DPS) [rep] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 34.3 healing_power points (3.13 DPS) | yes | Wisdom's Leather Leggings (252519, -0.01 DPS) [crafted]; Pristine Leggings (253987, -0.05 DPS) [crafted]; Warchief Kilt (7760, -0.10 DPS) [dungeon] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 38.5 healing_power points (3.50 DPS) | yes | Mender's Mail Boots (252565, -0.11 DPS) [crafted]; Mender's Leather Shoes (252533, -0.25 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -0.68 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Electrocutioner Lagnut (9447, -0.63 DPS) [dungeon]; Runemender's Seal (273031, -0.63 DPS) [world]; Welken Ring (5011, -0.65 DPS) [world_drop] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 16.1 healing_power points (1.47 DPS) | yes | Runemender's Seal (273031, -0.46 DPS) [world]; Welken Ring (5011, -0.47 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.71 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (78.1 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (78.1 DPS) | yes | Death Speaker Scepter (2816, -2.90 DPS) [dungeon]; Frost Tiger Blade (3854, -3.21 DPS, sim-verified) [crafted]; Pysan's Old Greatsword (1975, -7.36 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 23.6 healing_power points (2.15 DPS) | yes | Beacon of Hope (9393, -0.61 DPS) [dungeon]; Battered Viking Shield (9403, -0.76 DPS) [dungeon]; Orb of Souls (249395, -1.32 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Reflective Wristguards; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

### Band 50 (human, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 109.5. Weights run: 1.7s. Verify run: 9.9s. 728 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.002, intellect=1.851 ± 0.018, spirit=1.397 ± 0.016, mp5=4.657 ± 0.021, crit=0.325 ± 0.021 per rating point (14 rating = 1%, 4.550 per %), spell_haste=not significant (0.062 ± 0.067)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Imbued Helmet (220810) | Captain Dirgehammer [vendor] | 78.2 healing_power points (5.62 DPS) | yes | Papal Fez (9431, -1.45 DPS) [dungeon]; Cassandra's Grace (13102, -1.55 DPS) [world_drop]; Gemburst Circlet (10751, -3.07 DPS, sim-verified) [quest] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 39.1 healing_power points (2.81 DPS) | yes | Glowing Eye of Mordresh (10769, -0.81 DPS) [dungeon]; Horizon Choker (13085, -1.16 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.17 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Imbued Pauldrons (220808) | Captain Dirgehammer [vendor] | 61.0 healing_power points (4.38 DPS) | yes | Mender's Mail Shoulder (252569, -0.76 DPS) [crafted]; Nethergeld Shoulders (254049, -0.85 DPS) [crafted]; Mender's Leather Shoulder (252538, -1.71 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 32.9 healing_power points (2.36 DPS) | yes | Caretaker's Cape (19531, -0.18 DPS) [rep]; Imperial Red Cloak (8248, -0.50 DPS) [world_drop] |
| chest | Knight's Imbued Armor (220813) | Captain Dirgehammer [vendor] | 71.7 healing_power points (5.15 DPS) | yes | Stormcloth Vest (10020, -0.28 DPS, sim-verified) [crafted]; Ghostweave Vest (14141, -0.47 DPS) [crafted]; Feathered Breastplate (8349, -1.41 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 40.3 healing_power points (2.90 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.18 DPS) [crafted] |
| hands | Soulforge Fists (226982) | Mokvar [vendor] | 69.5 healing_power points (5.00 DPS) | yes | Raider Handwraps (272098, -0.49 DPS) [vendor]; Mender's Leather Gauntlets (252551, -0.67 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.67 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 52.2 healing_power points (3.75 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.27 DPS) [crafted]; Dawnspire Cord (12466, -0.36 DPS) [dungeon] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 71.7 healing_power points (5.15 DPS) | yes | Senior Designer's Pantaloons (11841, -0.62 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.76 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -2.70 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 48.4 healing_power points (3.48 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Leather Shoes (252533, -0.12 DPS) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 32.6 healing_power points (2.34 DPS) | yes | Welken Ring (5011, -0.98 DPS) [world_drop]; Lorekeeper's Ring (19523, -1.00 DPS) [rep]; Choking Band (11868, -1.04 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 27.9 healing_power points (2.00 DPS) | yes | Lorekeeper's Ring (19523, -0.66 DPS) [rep]; Choking Band (11868, -0.70 DPS) [quest]; Welken Ring (5011, -0.82 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (109.5 DPS) | yes | Ankh of Life (1713, -3.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.02 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.22 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (109.5 DPS) | yes | Ankh of Life (1713, -0.09 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.49 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.58 DPS, sim-verified) [quest] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (109.5 DPS) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.67 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.25 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 30.5 healing_power points (2.19 DPS) | yes | Mountainside Buckler (13082, -0.32 DPS) [world_drop]; Protector Buckler (14790, -0.36 DPS) [world_drop]; Twisting Essence Jar (249456, -0.45 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Imbued Helmet; neck: Darkmoon Necklace; shoulder: Knight-Lieutenant's Imbued Pauldrons; back: Darkspear Raider's Cloak; chest: Knight's Imbued Armor; wrist: Mender's Leather Bracers; hands: Soulforge Fists; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Mender's Leather Boots; finger1: Darkspear Signet; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength

No-known-source sample (15 of 728, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

### Band 60 (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 204.2. Weights run: 1.9s. Verify run: 34.0s. 1382 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.136, intellect=4.683 ± 0.078, spirit=1.589 ± 0.017, mp5=6.636 ± 0.062, crit=0.464 ± 0.044 per rating point (14 rating = 1%, 6.491 per %), spell_haste=not significant (0.057 ± 0.055)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 204.9 healing_power points (13.20 DPS) | yes | Living Crown (252561, -1.63 DPS) [crafted]; Crown of the Penitent (13216, -1.78 DPS) [quest]; Black Dragonscale Helm (252605, -3.44 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 105.2 healing_power points (6.78 DPS) | yes | Stormpike Sage's Pendant (19098, -2.05 DPS) [rep]; Horizon Choker (13085, -2.14 DPS) [world_drop]; Drake Tooth Necklace (21531, -6.06 DPS, sim-verified) [quest] |
| shoulder | Soulforge Epaulets (226979) | Mokvar [vendor] | sim-verified (204.2 DPS) | yes | Darkspear Shoulders (272104, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -10.15 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 88.8 healing_power points (5.72 DPS) | yes | Faded Hakkari Cloak (20218, -0.74 DPS) [quest]; Darkspear Raider's Cloak (272076, -0.99 DPS) [vendor] |
| chest | Knight-Captain's Lamellar Chestplate (227151) | Captain Dirgehammer [vendor] | sim-verified (204.2 DPS) | yes | Breastplate of Salvation (250601, -0.05 DPS) [crafted]; Devout Robe (16690, -0.16 DPS) [dungeon]; Soulforge Embrace (226984, -12.83 DPS, sim-verified) [vendor] |
| wrist | Soulforge Bindings (226977) | Mokvar [vendor] | sim-verified (204.2 DPS) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; Magister's Bindings (16683, -0.47 DPS) [dungeon]; Bracers of Hope (22667, -7.43 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 150.4 healing_power points (9.70 DPS) | yes | Mooncloth Gloves (18409, -2.74 DPS) [crafted]; Marshal's Lamellar Gloves (231643, -2.75 DPS) [pvp]; Soulforge Fists (226982, -2.98 DPS) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 126.1 healing_power points (8.12 DPS) | yes | Wisdom of the Timbermaw (19047, -0.98 DPS) [crafted]; Elunarian Belt (14465, -1.06 DPS) [world_drop]; Devout Belt (16696, -2.68 DPS, sim-verified) [dungeon] |
| legs | Knight-Captain's Lamellar Legguards (227150) | Captain Dirgehammer [vendor] | sim-verified (204.2 DPS) | yes | Red Dragonscale Leggings (252603, +0.00 DPS) [crafted]; Sentinel's Lamellar Legguards (237814, -1.55 DPS) [vendor]; Martyr's Legplates (250600, -6.57 DPS, sim-verified) [crafted] |
| feet | Soulforge Treads (226983) | Mokvar [vendor] | sim-verified (204.2 DPS) | yes | Knight-Lieutenant's Lamellar Greaves (227153, -0.49 DPS) [vendor]; Mooncloth Boots (15802, -1.01 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -6.57 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 110.3 healing_power points (7.11 DPS) | yes | Channeler's Ring (272406, -2.07 DPS) [vendor]; Band of the Hierophant (13096, -2.97 DPS) [world_drop]; Darkspear Signet (272068, -3.26 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 95.4 healing_power points (6.15 DPS) | yes | Band of the Hierophant (13096, -2.00 DPS) [world_drop]; Darkspear Signet (272068, -2.30 DPS) [vendor]; Channeler's Ring (272406, -5.79 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (204.2 DPS) | yes | Ankh of Life (1713, -5.19 DPS) [world_drop]; Uther's Strength (11302, -5.25 DPS) [world_drop]; Serenity Field (272439, -8.16 DPS, sim-verified) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (204.2 DPS) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -1.34 DPS) [world_drop]; Uther's Strength (11302, -1.41 DPS) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (204.2 DPS) | yes | Grand Marshal's Demolisher (234568, -0.42 DPS) [pvp]; Grand Marshal's Warhammer (23454, -1.90 DPS) [vendor]; Hand of Edward the Odd (2243, -1.93 DPS) [world_drop] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 73.2 healing_power points (4.72 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, -0.41 DPS) [vendor]; Elunarian Sphere (15968, -8.58 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Soulforge Epaulets; back: Hide of the Wild; chest: Knight-Captain's Lamellar Chestplate; wrist: Soulforge Bindings; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Knight-Captain's Lamellar Legguards; feet: Soulforge Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1382, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

### Band 60, raid preset (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 575.2. Weights run: 1.5s. Verify run: 18.5s. 1382 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.852, intellect=1.784 ± 0.122, spirit=0.691 ± 0.118, mp5=4.105 ± 0.149, crit=0.893 ± 0.140 per rating point (14 rating = 1%, 12.506 per %), spell_haste=not significant (-2.578 ± 1.753)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 114.8 healing_power points (19.29 DPS) | yes | Lieutenant Commander's Lamellar Helmet (227149, -2.38 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, -3.78 DPS) [vendor]; Gnomish Turban of Psychic Might (21517, -12.36 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 67.5 healing_power points (11.35 DPS) | yes | Drake Tooth Necklace (21531, -2.86 DPS, sim-verified) [quest]; Darkmoon Necklace (19303, -5.41 DPS) [vendor]; Stormpike Sage's Pendant (19098, -5.59 DPS) [rep] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 90.5 healing_power points (15.21 DPS) | yes | Shimmering Dawnbringer Shoulders (227859, +0.00 DPS, sim-verified) [vendor]; Lieutenant Commander's Lamellar Pauldrons (227148, -2.73 DPS) [vendor]; Knight-Lieutenant's Imbued Pauldrons (220808, -5.46 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 59.8 healing_power points (10.06 DPS) | yes | Faded Hakkari Cloak (20218, -3.52 DPS) [quest]; Darkspear Raider's Cloak (272063, -4.56 DPS) [vendor]; Drape of Recovery (272413, -5.66 DPS, sim-verified) [vendor] |
| chest | Knight-Captain's Lamellar Chestplate (227151) | Captain Dirgehammer [vendor] | 96.9 healing_power points (16.27 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, -0.77 DPS) [pvp]; Breastplate of Salvation (250601, -2.68 DPS, sim-verified) [crafted]; Soulforge Embrace (226984, -3.52 DPS) [vendor] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 45.9 healing_power points (7.71 DPS) | yes | Soulforge Bindings (226977, -0.80 DPS, sim-verified) [vendor]; Prefect's Wristguards (250584, -1.45 DPS) [crafted]; Nethergeld Cuffs (254061, -1.68 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 72.2 healing_power points (12.13 DPS) | yes | Mooncloth Gloves (18409, -1.47 DPS, sim-verified) [crafted]; Marshal's Lamellar Gloves (231643, -1.77 DPS) [pvp]; Soulforge Fists (226982, -2.71 DPS) [vendor] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 62.2 healing_power points (10.45 DPS) | yes | Soulforge Cord (226978, -1.09 DPS) [vendor]; Devout Belt (16696, -1.56 DPS) [dungeon]; Belt of Tiny Heads (20217, -4.66 DPS, sim-verified) [quest] |
| legs | Knight-Captain's Lamellar Legguards (227150) | Captain Dirgehammer [vendor] | sim-verified (+7.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Lamellar Legguards (231639, -2.76 DPS) [vendor]; Red Dragonscale Leggings (252603, -3.62 DPS) [crafted]; Martyr's Legplates (250600, -7.93 DPS, sim-verified) [crafted] |
| feet | Knight-Lieutenant's Lamellar Greaves (227153) | Captain Dirgehammer [vendor] | sim-verified (+9.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulforge Treads (226983, -0.96 DPS) [vendor]; Mooncloth Boots (15802, -2.18 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -9.70 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 71.6 healing_power points (12.03 DPS) | yes | Darkspear Signet (272068, -5.82 DPS) [vendor]; Blessed Band of Light (272407, -5.89 DPS) [vendor]; Channeler's Ring (272406, -6.90 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 56.3 healing_power points (9.45 DPS) | yes | Blessed Band of Light (272407, -3.32 DPS) [vendor]; Channeler's Ring (272406, -4.33 DPS) [vendor]; Darkspear Signet (272068, -7.97 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+10.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -6.21 DPS) [quest]; Uther's Strength (11302, -7.32 DPS) [world_drop]; Ankh of Life (1713, -8.95 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -2.92 DPS) [quest]; Ankh of Life (1713, -5.66 DPS) [world_drop]; Uther's Strength (11302, -7.51 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -5.30 DPS) [dungeon]; Inventor's Focal Sword (17719, -6.64 DPS) [dungeon] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | sim-verified (+25.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Lei of the Lifegiver (19312, -25.94 DPS, sim-verified) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Knight-Captain's Lamellar Chestplate; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Knight-Captain's Lamellar Legguards; feet: Knight-Lieutenant's Lamellar Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Evergreen Shield

No-known-source sample (15 of 1382, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

## Horde

### Band 20 (undead, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 26.9. Weights run: 1.4s. Verify run: 4.2s. 205 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.901 ± 0.003, spirit=0.222 ± 0.001, mp5=1.699 ± 0.033, crit=0.037 ± 0.002 per rating point (14 rating = 1%, 0.512 per %), spell_haste=0.104 ± 0.016

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.76 DPS) | yes | Wisdom's Leather Hood (252507, -0.34 DPS) [crafted]; Pristine Circlet (253949, -0.39 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.16 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.7 healing_power points (0.06 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.38 DPS) [crafted]; Slime-encrusted Pads (6461, -0.41 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.61 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 9.4 healing_power points (0.79 DPS) | yes | Sanguine Cape (14376, -0.49 DPS, sim-verified) [world_drop]; Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.60 DPS) [dungeon] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 20.5 healing_power points (1.72 DPS) | yes | Wisdom's Leather Armor (252493, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.63 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.6 healing_power points (0.80 DPS) | yes | Owl Bracers (4796, -0.43 DPS) [vendor]; Featherbead Bracers (15452, -0.43 DPS) [quest]; Garrison Cuffs (270003, -0.93 DPS, sim-verified) [quest] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.5 healing_power points (1.22 DPS) | yes | Wisdom's Leather Gloves (252499, -0.08 DPS) [crafted]; Pristine Gloves (253913, -0.19 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.59 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.22 DPS) | yes | Wisdom's Leather Belt (252433, -0.25 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.26 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.44 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 27.3 healing_power points (2.29 DPS) | yes | Wisdom's Leather Pants (252503, -0.14 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.25 DPS) [crafted]; Ghastly Trousers (15449, -1.63 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.5 healing_power points (1.30 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.24 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.45 DPS) | yes | Advisor's Ring (20426, -0.17 DPS) [rep]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop]; Band of Purification (12996, -0.34 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 4.5 healing_power points (0.38 DPS) | yes | Advisor's Ring (20426, +0.00 DPS, sim-verified) [rep]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop]; Band of Purification (12996, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 8.1 healing_power points (0.68 DPS) | yes | Impaling Harpoon (5200, -0.54 DPS, sim-verified) [dungeon]; Samophlange Screwdriver (11854, -0.57 DPS) [quest]; Mug of Muddled Memories (277247, -0.59 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Sword of the Fallen

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 205218 Libram of Discovery; 209614 Insignia of the Alliance; 213513 Libram of Deliverance; 248005 Northshire Hammer

### Band 30 (undead, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 47.9. Weights run: 1.5s. Verify run: 4.1s. 389 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.985 ± 0.005, spirit=0.496 ± 0.004, mp5=2.104 ± 0.011, crit=0.092 ± 0.005 per rating point (14 rating = 1%, 1.288 per %), spell_haste=not significant (0.037 ± 0.013)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.0 healing_power points (3.37 DPS) | yes | Filigreed Pristine Circlet (253975, -1.31 DPS) [crafted]; Wisdom's Leather Helm (252515, -1.40 DPS) [crafted]; Acolyte's Chain Helm (250501, -1.53 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.9 healing_power points (0.56 DPS) | yes | Scorn's Icy Choker (23169, -0.00 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.00 DPS) [world]; Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 13.5 healing_power points (1.26 DPS) | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.29 DPS) [world_drop]; Talbar Mantle (10657, -0.31 DPS) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 15.0 healing_power points (1.41 DPS) | yes | Darkspear Raider's Cloak (272078, -0.53 DPS) [vendor]; Resilient Cape (14400, -0.85 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.89 DPS, sim-verified) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 30.8 healing_power points (2.89 DPS) | yes | Pristine Gown (253961, -0.18 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.37 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -0.63 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 11.2 healing_power points (1.05 DPS) | yes | Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.01 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.0 healing_power points (1.68 DPS) | yes | Acolyte's Gloves (250511, -0.28 DPS) [crafted]; Pristine Gloves (253913, -0.38 DPS) [crafted]; Naga Battle Gloves (888, -0.61 DPS, sim-verified) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 28.9 healing_power points (2.71 DPS) | yes | Mender's Leather Belt (252523, -0.33 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.31 DPS) [crafted]; Acolyte's Chain Belt (250516, -1.50 DPS) [crafted] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 32.9 healing_power points (3.08 DPS) | yes | Wisdom's Leather Leggings (252519, -0.14 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.47 DPS) [crafted]; Pristine Leggings (253987, -0.51 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.9 healing_power points (2.71 DPS) | yes | Wisdom's Leather Boots (252444, -1.21 DPS) [crafted]; Glowing Copper Boots (250482, -1.49 DPS) [crafted]; Acolyte's Boots (250506, -1.79 DPS, sim-verified) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.69 DPS) | yes | Electrocutioner Lagnut (9447, -0.80 DPS) [dungeon]; Runemender's Seal (273031, -0.80 DPS) [world]; Darkspear Signet (272071, -0.90 DPS) [vendor] |
| finger2 | Solid Gold Band (279845) | Opportunistic Education [quest] | 9.5 healing_power points (0.89 DPS) | yes | Runemender's Seal (273031, -0.00 DPS) [world]; Darkspear Signet (272071, -0.10 DPS) [vendor]; Electrocutioner Lagnut (9447, -0.19 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (47.9 DPS) | yes | Manual Crowd Pummeler (9449, -2.02 DPS, sim-verified) [dungeon]; Pysan's Old Greatsword (1975, -4.59 DPS) [dungeon]; Wolfsbane (267369, -5.45 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.0 healing_power points (1.50 DPS) | yes | Orb of Mistmantle (13031, -0.34 DPS, sim-verified) [world_drop]; Heart of Agamaggan (6694, -0.56 DPS) [dungeon]; Witch's Finger (16887, -0.76 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Solid Gold Band; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 389, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets

### Band 40 (undead, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 92.4. Weights run: 1.6s. Verify run: 7.3s. 536 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.164 ± 0.009, spirit=0.728 ± 0.008, mp5=2.684 ± 0.012, crit=0.162 ± 0.009 per rating point (14 rating = 1%, 2.271 per %), spell_haste=not significant (0.080 ± 0.025)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 40.3 healing_power points (3.67 DPS) | yes | Earthen Silk Hood (254015, -1.01 DPS) [crafted]; Acolyte's Chain Helm (250501, -1.40 DPS) [crafted]; Holy Shroud (2721, -2.42 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.3 healing_power points (2.03 DPS) | yes | Necklace of Calisea (1714, -1.18 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -1.29 DPS) [vendor]; Kaleidoscope Chain (13084, -1.34 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 30.1 healing_power points (2.74 DPS) | yes | Sheepshear Mantle (13115, -0.65 DPS) [world_drop]; Mistscape Mantle (4734, -1.25 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.60 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | 21.6 healing_power points (1.97 DPS) | yes | Darkspear Raider's Cloak (272077, -0.54 DPS) [vendor]; Glowing Thresher Cape (6901, -0.80 DPS) [dungeon]; Blackforge Cape (6424, -0.86 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.8 healing_power points (5.81 DPS) | yes | Death Speaker Robes (6682, -2.83 DPS) [dungeon]; Pristine Gown (253961, -2.99 DPS) [crafted]; Deathchill Armor (10764, -3.84 DPS, sim-verified) [dungeon] |
| wrist | Reflective Wristguards (274750) | Rettrick [vendor] | 18.3 healing_power points (1.67 DPS) | yes | Earthen Silk Cuffs (254019, -0.43 DPS) [crafted]; Mindthrust Bracers (1974, -0.44 DPS, sim-verified) [dungeon]; Enchanted Kodo Bracers (13119, -0.58 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 31.5 healing_power points (2.87 DPS) | yes | Prefect's Gauntlet (250569, -0.21 DPS) [crafted]; Gloves of the Greatfather (17721, -0.50 DPS) [crafted]; Gilded Handwraps (254021, -0.54 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.7 healing_power points (3.25 DPS) | yes | Prefect's Belt (250559, -0.50 DPS) [crafted]; Mender's Leather Belt (252523, -1.14 DPS, sim-verified) [crafted]; Highlander's Mail Girdle (20119, -1.66 DPS) [vendor] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 34.3 healing_power points (3.13 DPS) | yes | Wisdom's Leather Leggings (252519, +0.00 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -0.05 DPS) [crafted]; Warchief Kilt (7760, -0.10 DPS) [dungeon] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 38.5 healing_power points (3.50 DPS) | yes | Mender's Leather Shoes (252533, -0.11 DPS) [crafted]; Mender's Mail Boots (252565, -0.11 DPS) [crafted]; Gilded Slippers (254001, -0.68 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.64 DPS) | yes | Electrocutioner Lagnut (9447, -0.63 DPS) [dungeon]; Runemender's Seal (273031, -0.63 DPS) [world]; Welken Ring (5011, -0.65 DPS) [world_drop] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 16.1 healing_power points (1.47 DPS) | yes | Runemender's Seal (273031, -0.46 DPS) [world]; Welken Ring (5011, -0.47 DPS) [world_drop]; Electrocutioner Lagnut (9447, -2.93 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (92.4 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (92.4 DPS) | yes | Death Speaker Scepter (2816, -2.90 DPS) [dungeon]; Skullbreaker (17039, -5.83 DPS) [quest]; Frost Tiger Blade (3854, -18.18 DPS, sim-verified) [crafted] |
| off_hand | Forcestone Buckler (17508) | Compendium of the Fallen [quest] | 17.8 healing_power points (1.62 DPS) | yes | Beacon of Hope (9393, -0.08 DPS) [dungeon]; Battered Viking Shield (9403, -0.23 DPS) [dungeon]; Orb of Souls (249395, -16.24 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; wrist: Reflective Wristguards; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Forcestone Buckler

No-known-source sample (15 of 536, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (undead, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 113.8. Weights run: 1.7s. Verify run: 9.4s. 704 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.002, intellect=1.851 ± 0.018, spirit=1.397 ± 0.016, mp5=4.657 ± 0.021, crit=0.325 ± 0.021 per rating point (14 rating = 1%, 4.550 per %), spell_haste=not significant (0.062 ± 0.067)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 65.6 healing_power points (4.72 DPS) | yes | Papal Fez (9431, -0.33 DPS, sim-verified) [dungeon]; Cassandra's Grace (13102, -0.65 DPS) [world_drop]; Winged Helm (13112, -0.65 DPS) [world_drop] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 39.1 healing_power points (2.81 DPS) | yes | Glowing Eye of Mordresh (10769, -0.81 DPS) [dungeon]; Horizon Choker (13085, -0.93 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.17 DPS) [dungeon] |
| shoulder | Mender's Leather Shoulder (252538) (or Mender's Mail Shoulder (252569)) | Leatherworking [crafted] | 50.4 healing_power points (3.63 DPS) | yes | Mender's Mail Shoulder (252569, +0.00 DPS) [crafted]; Nethergeld Shoulders (254049, -0.09 DPS) [crafted]; Living Shoulders (15061, -0.09 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 32.9 healing_power points (2.36 DPS) | yes | Battle Healer's Cloak (19527, -0.18 DPS) [rep]; Imperial Red Cloak (8248, -0.50 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 69.2 healing_power points (4.97 DPS) | yes | Feathered Breastplate (8349, -1.23 DPS) [crafted]; Earth Warder's Vest (21311, -1.23 DPS) [quest]; Ghostweave Vest (14141, -2.27 DPS, sim-verified) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 40.3 healing_power points (2.90 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.18 DPS) [crafted] |
| hands | Soulforge Fists (226982) | Mokvar [vendor] | 69.5 healing_power points (5.00 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Mender's Leather Gauntlets (252551, -0.67 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.67 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 52.2 healing_power points (3.75 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.27 DPS) [crafted]; Dawnspire Cord (12466, -0.36 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 64.5 healing_power points (4.63 DPS) | yes | Senior Designer's Pantaloons (11841, +0.00 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.25 DPS) [dungeon]; Dalewind Trousers (13008, -0.93 DPS) [world_drop] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 48.4 healing_power points (3.48 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Leather Shoes (252533, -0.12 DPS) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 32.6 healing_power points (2.34 DPS) | yes | Welken Ring (5011, -0.98 DPS) [world_drop]; Ring of Subtlety (19038, -1.00 DPS) [quest]; Advisor's Ring (19519, -1.00 DPS) [rep] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 27.9 healing_power points (2.00 DPS) | yes | Welken Ring (5011, -0.64 DPS, sim-verified) [world_drop]; Ring of Subtlety (19038, -0.66 DPS) [quest]; Advisor's Ring (19519, -0.66 DPS) [rep] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -3.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.02 DPS) [quest]; Alchemist's Stone (13503, -4.62 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.29 DPS) [quest]; Alchemist's Stone (13503, -0.89 DPS) [crafted]; Ankh of Life (1713, -1.23 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.67 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.25 DPS) [dungeon] |
| off_hand | Mountainside Buckler (13082) | World drop [world_drop] | sim-verified (113.8 DPS) | yes | Protector Buckler (14790, -0.04 DPS) [world_drop]; Forcestone Buckler (17508, -0.20 DPS) [quest]; Twisting Essence Jar (249456, -12.33 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Mender's Leather Shoulder; back: Darkspear Raider's Cloak; wrist: Mender's Leather Bracers; hands: Soulforge Fists; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Mender's Leather Boots; finger1: Darkspear Signet; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; off_hand: Mountainside Buckler

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 194.7. Weights run: 1.9s. Verify run: 33.0s. 1404 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.136, intellect=4.683 ± 0.078, spirit=1.589 ± 0.017, mp5=6.636 ± 0.062, crit=0.464 ± 0.044 per rating point (14 rating = 1%, 6.491 per %), spell_haste=not significant (0.057 ± 0.055)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 204.9 healing_power points (13.20 DPS) | yes | Living Crown (252561, -1.63 DPS) [crafted]; Crown of the Penitent (13216, -1.78 DPS) [quest]; Black Dragonscale Helm (252605, -3.47 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 105.2 healing_power points (6.78 DPS) | yes | Frostwolf Advisor's Pendant (19096, -2.05 DPS) [rep]; Horizon Choker (13085, -2.14 DPS) [world_drop]; Drake Tooth Necklace (21531, -5.97 DPS, sim-verified) [quest] |
| shoulder | Soulforge Epaulets (226979) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Shoulders (272104, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.01 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 88.8 healing_power points (5.72 DPS) | yes | Faded Hakkari Cloak (20218, -0.74 DPS) [quest]; Darkspear Raider's Cloak (272076, -0.99 DPS) [vendor] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | sim-verified (194.7 DPS) | yes | Devout Robe (16690, -0.11 DPS) [dungeon]; Mooncloth Vest (14138, -0.53 DPS) [crafted]; Soulforge Embrace (226984, -2.70 DPS, sim-verified) [vendor] |
| wrist | Soulforge Bindings (226977) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; Bracers of Hope (22667, +0.00 DPS, sim-verified) [quest]; Magister's Bindings (16683, -0.47 DPS) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 150.4 healing_power points (9.70 DPS) | yes | Mooncloth Gloves (18409, -2.74 DPS) [crafted]; Soulforge Fists (226982, -2.98 DPS) [vendor]; Raider Handguards (272101, -3.36 DPS) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 126.1 healing_power points (8.12 DPS) | yes | Wisdom of the Timbermaw (19047, -0.98 DPS) [crafted]; Elunarian Belt (14465, -1.06 DPS) [world_drop]; Devout Belt (16696, -2.83 DPS, sim-verified) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Mail Leggings (22676, -1.15 DPS) [rep]; Sentinel's Lamellar Legguards (237814, -1.79 DPS) [vendor]; Red Dragonscale Leggings (252603, -7.61 DPS, sim-verified) [crafted] |
| feet | Soulforge Treads (226983) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -1.01 DPS) [crafted]; Devout Sandals (16691, -2.65 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 110.3 healing_power points (7.11 DPS) | yes | Channeler's Ring (272406, -2.07 DPS) [vendor]; Band of the Hierophant (13096, -2.97 DPS) [world_drop]; Darkspear Signet (272068, -3.26 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 95.4 healing_power points (6.15 DPS) | yes | Band of the Hierophant (13096, -2.00 DPS) [world_drop]; Darkspear Signet (272068, -2.30 DPS) [vendor]; Channeler's Ring (272406, -5.78 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -5.19 DPS) [world_drop]; Uther's Strength (11302, -5.25 DPS) [world_drop]; Serenity Field (272439, -8.06 DPS, sim-verified) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -1.34 DPS) [world_drop]; Uther's Strength (11302, -1.41 DPS) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Destroyer (234546, -0.42 DPS) [pvp]; High Warlord's Battle Mace (23464, -1.90 DPS) [vendor]; Hand of Edward the Odd (2243, -1.93 DPS) [world_drop] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 73.2 healing_power points (4.72 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, -0.41 DPS) [vendor]; Elunarian Sphere (15968, -8.11 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Soulforge Epaulets; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Soulforge Bindings; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Martyr's Legplates; feet: Soulforge Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1404, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 565.2. Weights run: 1.5s. Verify run: 17.2s. 1404 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.852, intellect=1.784 ± 0.122, spirit=0.691 ± 0.118, mp5=4.105 ± 0.149, crit=0.893 ± 0.140 per rating point (14 rating = 1%, 12.506 per %), spell_haste=not significant (-2.578 ± 1.753)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 114.8 healing_power points (19.29 DPS) | yes | Soulforge Crown (226981, -4.69 DPS) [vendor]; Crown of the Penitent (13216, -5.15 DPS) [quest]; Gnomish Turban of Psychic Might (21517, -12.42 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 67.5 healing_power points (11.35 DPS) | yes | Drake Tooth Necklace (21531, -3.29 DPS, sim-verified) [quest]; Darkmoon Necklace (19303, -5.41 DPS) [vendor]; Frostwolf Advisor's Pendant (19096, -5.59 DPS) [rep] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 90.5 healing_power points (15.21 DPS) | yes | Shimmering Dawnbringer Shoulders (227859, +0.00 DPS, sim-verified) [vendor]; Devout Mantle (16695, -6.02 DPS) [dungeon]; Soulforge Epaulets (226979, -6.23 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 59.8 healing_power points (10.06 DPS) | yes | Faded Hakkari Cloak (20218, -3.52 DPS) [quest]; Darkspear Raider's Cloak (272063, -4.56 DPS) [vendor]; Drape of Recovery (272413, -6.07 DPS, sim-verified) [vendor] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | 87.8 healing_power points (14.75 DPS) | yes | Soulforge Embrace (226984, -2.12 DPS, sim-verified) [vendor]; Mooncloth Vest (14138, -2.18 DPS) [crafted]; Devout Robe (16690, -2.45 DPS) [dungeon] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 45.9 healing_power points (7.71 DPS) | yes | Soulforge Bindings (226977, -0.88 DPS, sim-verified) [vendor]; Prefect's Wristguards (250584, -1.45 DPS) [crafted]; Nethergeld Cuffs (254061, -1.68 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 72.2 healing_power points (12.13 DPS) | yes | Mooncloth Gloves (18409, -1.94 DPS, sim-verified) [crafted]; Soulforge Fists (226982, -2.71 DPS) [vendor]; Prefect's Gloves (250554, -2.95 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 62.2 healing_power points (10.45 DPS) | yes | Soulforge Cord (226978, -1.09 DPS) [vendor]; Devout Belt (16696, -1.56 DPS) [dungeon]; Belt of Tiny Heads (20217, -4.93 DPS, sim-verified) [quest] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | 110.4 healing_power points (18.55 DPS) | yes | Outrider's Mail Leggings (22676, -5.72 DPS) [rep]; Sentinel's Lamellar Legguards (237814, -6.25 DPS) [vendor]; Red Dragonscale Leggings (252603, -10.01 DPS, sim-verified) [crafted] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 78.6 healing_power points (13.21 DPS) | yes | Mooncloth Boots (15802, -3.48 DPS) [crafted]; Mender's Leather Boots (252472, -5.21 DPS) [crafted]; Soulforge Treads (226983, -6.54 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 71.6 healing_power points (12.03 DPS) | yes | Darkspear Signet (272068, -5.82 DPS) [vendor]; Blessed Band of Light (272407, -5.89 DPS) [vendor]; Channeler's Ring (272406, -6.90 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 56.3 healing_power points (9.45 DPS) | yes | Blessed Band of Light (272407, -3.32 DPS) [vendor]; Channeler's Ring (272406, -4.33 DPS) [vendor]; Darkspear Signet (272068, -8.29 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+11.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -6.21 DPS) [quest]; Uther's Strength (11302, -7.32 DPS) [world_drop]; Ankh of Life (1713, -8.95 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -2.92 DPS) [quest]; Ankh of Life (1713, -5.66 DPS) [world_drop]; Uther's Strength (11302, -6.91 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -5.30 DPS) [dungeon]; Inventor's Focal Sword (17719, -6.64 DPS) [dungeon] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | sim-verified (565.1 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Lei of the Lifegiver (19312, -24.76 DPS, sim-verified) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Martyr's Legplates; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Evergreen Shield

No-known-source sample (15 of 1404, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

