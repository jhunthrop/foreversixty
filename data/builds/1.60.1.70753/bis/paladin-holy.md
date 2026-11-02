# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 50320100000000000-00000000000000000-00000000000000000)

Set DPS (verified): 28.4. Weights run: 4.0s. Verify run: 14.6s. 262 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.003 ± 0.002, spirit=0.518 ± 0.003, mp5=1.938 ± 0.012, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), spell_haste=not significant (0.016 ± 0.005)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.82 DPS) | yes | Wisdom's Leather Hood (252507, -0.35 DPS) [crafted]; Pristine Circlet (253949, -0.36 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.12 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.0 healing_power points (0.09 DPS) | yes | Tarnished Locket (279870, -0.18 DPS, sim-verified) [quest] |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Reinforced Woolen Shoulders (4315, -0.16 DPS) [crafted]; Serpent's Shoulders (5404, -0.19 DPS) [dungeon]; Magician's Mantle (12998, -0.29 DPS, sim-verified) [world_drop] |
| back | Seer's Cape (6378) | Wailing Caverns: Large Battered Chest [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Pearl-clasped Cloak (5542, -0.00 DPS) [crafted]; Spirit Cloak (4792, -0.08 DPS) [vendor]; Sanguine Cape (14376, -0.68 DPS, sim-verified) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 21.6 healing_power points (1.87 DPS) | yes | Acolyte's Chain Shirt (250491, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.17 DPS) [crafted]; Corsair's Overshirt (5202, -0.56 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.8 healing_power points (0.94 DPS) | yes | Bright Bracers (3647, -0.59 DPS) [world_drop]; Owl Bracers (4796, -0.66 DPS, sim-verified) [vendor]; Drakewing Bands (12999, -0.67 DPS) [world_drop] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 15.0 healing_power points (1.30 DPS) | yes | Pristine Gloves (253913, -0.09 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.15 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.69 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.0 healing_power points (1.30 DPS) | yes | Acolyte's Chain Belt (250516, -0.24 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.26 DPS) [crafted]; Novice Ardent's Sash (253887, -0.39 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.17 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.73 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.55 DPS) [world_drop] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 16.0 healing_power points (1.39 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.26 DPS) [crafted]; Black Whelp Slippers (252424, -0.26 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.0 healing_power points (0.52 DPS) | yes | Volcanic Rock Ring (12053, -0.26 DPS) [world_drop]; Minor Channeling Ring (1449, -0.35 DPS) [quest]; Furlfeather Ring (286745, -0.35 DPS) [world] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 3.1 healing_power points (0.27 DPS) | yes | Volcanic Rock Ring (12053, -0.01 DPS) [world_drop]; Minor Channeling Ring (1449, -0.10 DPS) [quest]; Furlfeather Ring (286745, -0.10 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 11.2 healing_power points (0.97 DPS) | yes | Sword of the Fallen (286977, -0.15 DPS, sim-verified) [dungeon]; Impaling Harpoon (5200, -0.54 DPS) [dungeon]; Trogg Slicer (6186, -0.75 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Scholarly Pendant; shoulder: Slime-encrusted Pads; back: Seer's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Verigan's Fist

No-known-source sample (15 of 262, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers; 14148 Crystalline Cuffs

### Band 30 (human, 50320215012000000-00000000000000000-00000000000000000)

Set DPS (verified): 50.9. Weights run: 4.4s. Verify run: 17.0s. 484 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=1.315 ± 0.009, spirit=0.661 ± 0.005, mp5=2.497 ± 0.009, crit=0.110 ± 0.007 per rating point (14 rating = 1%, 1.537 per %), spell_haste=not significant (0.032 ± 0.016)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.0 healing_power points (2.99 DPS) | yes | Acolyte's Chain Helm (250501, -0.97 DPS) [crafted]; Filigreed Pristine Circlet (253975, -1.21 DPS) [crafted]; Enduring Cap (3020, -1.48 DPS, sim-verified) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.9 healing_power points (0.64 DPS) | yes | Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Pendant of Myzrael (4614, -0.32 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.65 DPS, sim-verified) [world] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 14.5 healing_power points (1.17 DPS) | yes | Nightsky Mantle (4718, -0.05 DPS) [world_drop]; Faerie Mantle (5820, -0.21 DPS) [quest]; Mantle of Honor (3560, -0.51 DPS, sim-verified) [quest] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 13.3 healing_power points (1.07 DPS) | yes | Prelacy Cape (7004, -0.35 DPS) [quest]; Sergeant Major's Cape (16315, -0.43 DPS) [pvp]; Glowing Thresher Cape (6901, -0.53 DPS, sim-verified) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 34.5 healing_power points (2.79 DPS) | yes | Acolyte's Silvered Chain Shirt (250521, -0.26 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.37 DPS) [crafted]; Pristine Gown (253961, -0.84 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.1 healing_power points (1.14 DPS) | yes | Nightsky Wristbands (6407, -0.34 DPS) [world_drop]; Spidertank Oilrag (9448, -0.53 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.40 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.9 healing_power points (1.53 DPS) | yes | Acolyte's Gloves (250511, -0.30 DPS, sim-verified) [crafted]; Wisdom's Leather Gloves (252499, -0.30 DPS) [crafted]; Naga Battle Gloves (888, -0.32 DPS) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 31.2 healing_power points (2.53 DPS) | yes | Mender's Leather Belt (252523, -0.63 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.21 DPS) [crafted]; Silver-lined Belt (13011, -1.35 DPS) [world_drop] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 35.5 healing_power points (2.87 DPS) | yes | Pristine Leggings (253987, -0.08 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.40 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.47 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 31.9 healing_power points (2.58 DPS) | yes | Wisdom's Leather Boots (252444, -1.16 DPS) [crafted]; Glowing Copper Boots (250482, -1.42 DPS) [crafted]; Acolyte's Boots (250506, -2.12 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.6 healing_power points (1.10 DPS) | yes | The Queen's Jewel (13094, -0.46 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Ring of Power Regulation (273030, -0.50 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 13.6 healing_power points (1.10 DPS) | yes | Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Ring of Power Regulation (273030, -0.50 DPS) [dungeon]; The Queen's Jewel (13094, -0.68 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (50.9 DPS) | yes | Pysan's Old Greatsword (1975, -4.19 DPS, sim-verified) [dungeon]; Mechanic's Pipehammer (9604, -4.23 DPS) [quest]; Verigan's Fist (6953, -4.33 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Repairman's Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter

No-known-source sample (15 of 484, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 50320215012101251-00000000000000000-00000000000000000)

Set DPS (verified): 68.8. Weights run: 6.8s. Verify run: 26.3s. 672 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.044, intellect=1.533 ± 0.014, spirit=0.963 ± 0.008, mp5=3.277 ± 0.012, crit=0.227 ± 0.014 per rating point (14 rating = 1%, 3.176 per %), spell_haste=not significant (0.025 ± 0.030)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 56.7 healing_power points (4.13 DPS) | yes | Whitemane's Chapeau (7720, -0.95 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.31 DPS) [world_drop]; Earthen Silk Hood (254015, -1.83 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.3 healing_power points (2.35 DPS) | yes | Necklace of Calisea (1714, -1.38 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -1.62 DPS) [world_drop]; Amberglow Talisman (10824, -1.65 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 41.9 healing_power points (3.05 DPS) | yes | Earthen Silk Shoulders (254033, -1.18 DPS) [crafted]; Sheepshear Mantle (13115, -1.30 DPS, sim-verified) [world_drop]; Mistscape Mantle (4734, -1.47 DPS) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 16.1 healing_power points (1.17 DPS) | yes | Repairman's Cape (9605, -0.05 DPS) [quest]; Sergeant Major's Cape (16336, -0.08 DPS) [pvp]; Glowing Thresher Cape (6901, -0.10 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.7 healing_power points (4.78 DPS) | yes | Deathchill Armor (10764, -2.00 DPS, sim-verified) [dungeon]; Civinad Robes (9623, -2.05 DPS) [quest]; Death Speaker Robes (6682, -2.10 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 17.5 healing_power points (1.27 DPS) | yes | Earthen Silk Cuffs (254019, -0.20 DPS) [crafted]; Windchaser Cuffs (14429, -0.27 DPS) [world_drop]; Enchanted Kodo Bracers (13119, -0.44 DPS, sim-verified) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 34.8 healing_power points (2.53 DPS) | yes | Gilded Handwraps (254021, -0.15 DPS) [crafted]; Prefect's Gauntlet (250569, -0.22 DPS) [crafted]; Stormcloth Gloves (10011, -0.39 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 40.0 healing_power points (2.91 DPS) | yes | Prefect's Belt (250559, -0.53 DPS) [crafted]; Mender's Leather Belt (252523, -0.59 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.24 DPS) [dungeon] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 46.7 healing_power points (3.39 DPS) | yes | Stoneweaver Leggings (9407, -0.36 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -0.60 DPS) [crafted]; Pristine Leggings (253987, -0.66 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 41.8 healing_power points (3.04 DPS) | yes | Mender's Leather Shoes (252533, +0.00 DPS, sim-verified) [crafted]; Mender's Mail Boots (252565, -0.02 DPS) [crafted]; Gilded Slippers (254001, -0.52 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 15.7 healing_power points (1.14 DPS) | yes | Welken Ring (5011, -0.10 DPS) [world_drop]; Ogremind Ring (1993, -0.15 DPS) [world_drop]; Voodoo Band (1996, -0.15 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 15.7 healing_power points (1.14 DPS) | yes | Welken Ring (5011, -0.14 DPS, sim-verified) [world_drop]; Ogremind Ring (1993, -0.15 DPS) [world_drop]; Voodoo Band (1996, -0.15 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (68.9 DPS) | yes | Pysan's Old Greatsword (1975, +0.00 DPS, sim-verified) [dungeon]; Mograine's Might (7723, -3.16 DPS) [dungeon]; Verigan's Fist (6953, -3.58 DPS) [quest] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 27.3 healing_power points (1.99 DPS) | yes | Beacon of Hope (9393, -0.11 DPS) [dungeon]; Battered Viking Shield (9403, -0.34 DPS) [dungeon]; Orb of Souls (249395, -0.62 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Prefect's Boots; trinket1: Ankh of Life; off_hand: Orb of Lorica

No-known-source sample (15 of 672, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (human, 50320215012101251-50320000000000000-00000000000000000)

Set DPS (verified): 95.7. Weights run: 7.0s. Verify run: 33.1s. 869 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.050, intellect=1.958 ± 0.018, spirit=1.458 ± 0.012, mp5=4.543 ± 0.016, crit=0.346 ± 0.019 per rating point (14 rating = 1%, 4.851 per %), spell_haste=not significant (0.114 ± 0.046)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 68.5 healing_power points (4.89 DPS) | yes | Papal Fez (9431, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Imbued Helmet (220810, +0.00 DPS) [vendor]; Winged Helm (13112, -0.64 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 35.7 healing_power points (2.55 DPS) | yes | Horizon Choker (13085, -0.58 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -0.84 DPS) [dungeon]; Mindburst Medallion (11196, -1.19 DPS) [quest] |
| shoulder | Mender's Leather Shoulders (252538) (or Mender's Mail Shoulders (252569)) | Leatherworking [crafted] | 51.8 healing_power points (3.70 DPS) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, +0.00 DPS) [vendor]; Mender's Mail Shoulders (252569, +0.00 DPS) [crafted]; Ironfeather Shoulders (15067, -0.07 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 27.4 healing_power points (1.95 DPS) | yes | Highborne Cloak (14450, -0.31 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -0.42 DPS) [dungeon]; Wingveil Cloak (10802, -0.45 DPS) [dungeon] |
| chest | Knight's Imbued Armor (220813) | Captain Dirgehammer [vendor] | 72.6 healing_power points (5.19 DPS) | yes | Stormcloth Vest (10020, -0.31 DPS, sim-verified) [crafted]; Ghostweave Vest (14141, -0.51 DPS) [crafted]; Deathchill Armor (10764, -0.94 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 41.5 healing_power points (2.96 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.20 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 61.7 healing_power points (4.41 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.11 DPS) [crafted]; Prefect's Gloves (250554, -0.37 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 53.5 healing_power points (3.82 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.28 DPS) [crafted]; Dawnspire Cord (12466, -0.31 DPS) [dungeon] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 72.6 healing_power points (5.19 DPS) | yes | Senior Designer's Pantaloons (11841, -0.55 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.67 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.28 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 49.5 healing_power points (3.54 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.07 DPS) [crafted]; Mender's Leather Shoes (252533, -0.12 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 29.2 healing_power points (2.09 DPS) | yes | Electrocutioner Lagnut (9447, -0.72 DPS) [dungeon]; Runemender's Seal (273031, -0.72 DPS) [world]; Coldwater Ring (4550, -0.73 DPS) [quest] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 20.0 healing_power points (1.43 DPS) | yes | Electrocutioner Lagnut (9447, -0.06 DPS) [dungeon]; Runemender's Seal (273031, -0.06 DPS) [world]; Coldwater Ring (4550, -0.07 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (95.8 DPS) | yes | Evonice's Landin' Pilla (18951, -0.24 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.45 DPS) [quest]; Alchemist's Stone (13503, -0.87 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (95.8 DPS) | yes | Evonice's Landin' Pilla (18951, -0.21 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -0.42 DPS) [quest]; Alchemist's Stone (13503, -0.83 DPS) [crafted] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (95.8 DPS) | yes | Inventor's Focal Sword (17719, -0.22 DPS) [dungeon]; Mograine's Might (7723, -2.62 DPS) [dungeon]; Pysan's Old Greatsword (1975, -2.78 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 31.6 healing_power points (2.26 DPS) | yes | Twisting Essence Jar (249456, -0.20 DPS, sim-verified) [crafted]; Beacon of Hope (9393, -0.27 DPS) [dungeon]; Battered Viking Shield (9403, -0.29 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; shoulder: Mender's Leather Shoulders; back: Imperial Red Cloak; chest: Knight's Imbued Armor; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Mender's Leather Boots; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life

No-known-source sample (15 of 869, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 166.9. Weights run: 6.9s. Verify run: 88.9s. 1792 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.049, intellect=3.366 ± 0.040, spirit=1.418 ± 0.013, mp5=5.083 ± 0.026, crit=0.400 ± 0.022 per rating point (14 rating = 1%, 5.595 per %), spell_haste=not significant (0.063 ± 0.073)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 153.4 healing_power points (12.60 DPS) | yes | Lieutenant Commander's Lamellar Helmet (227149, -1.60 DPS) [vendor]; Black Dragonscale Helm (252605, -1.64 DPS) [crafted]; Crown of the Penitent (13216, -1.75 DPS) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (166.8 DPS) | yes | Glowing Eye of Mordresh (10769, -0.48 DPS) [dungeon]; Opaline Medallion (16623, -0.74 DPS) [quest]; Archlight Talisman (15856, -1.57 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 94.4 healing_power points (7.76 DPS) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, -1.03 DPS) [pvp]; Knight-Lieutenant's Imbued Pauldrons (220808, -1.24 DPS) [vendor] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 42.7 healing_power points (3.51 DPS) | yes | Sergeant Major's Cape (16337, -0.36 DPS) [pvp]; Highborne Cloak (14450, -0.55 DPS) [world_drop]; Keeper's Cloak (14665, -0.74 DPS) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Chestplate (227151, +0.00 DPS) [vendor]; Breastplate of Salvation (250601, +0.00 DPS) [crafted]; Field Marshal's Lamellar Chestplate (231641, -0.04 DPS) [pvp] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; Magister's Bindings (16683, +0.00 DPS) [dungeon]; Mender's Leather Bracers (252543, -0.09 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Mender's Leather Gauntlets (252551, +0.00 DPS) [crafted]; Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 91.1 healing_power points (7.48 DPS) | yes | Wisdom of the Timbermaw (19047, -0.02 DPS) [crafted]; Elunarian Belt (14465, -0.79 DPS) [world_drop]; Marshal's Lamellar Belt (16470, -1.19 DPS) [pvp] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Legguards (227150, -1.12 DPS) [vendor]; Red Dragonscale Leggings (252603, -1.47 DPS) [crafted]; Devout Skirt (16694, -3.17 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 49.3 healing_power points (4.04 DPS) | yes | Glowing Crystal Ring (18402, -0.73 DPS) [quest]; Frigid Ring (18679, -1.28 DPS) [world]; Maiden's Circle (13001, -1.41 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Crystal Ring (18402, +0.00 DPS) [quest]; Frigid Ring (18679, -0.29 DPS) [world]; Maiden's Circle (13001, -0.42 DPS) [world_drop] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (166.8 DPS) | yes | Ankh of Life (1713, -0.08 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.31 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.55 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (166.8 DPS) | yes | Ankh of Life (1713, -1.11 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.34 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.57 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (166.8 DPS) | yes | Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Inventor's Focal Sword (17719, +0.00 DPS) [dungeon]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp] |
| off_hand | Elunarian Sphere (15968) | World drop [world_drop] | 47.1 healing_power points (3.87 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Orb of Lorica (11262, -0.12 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Horizon Choker; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves; main_hand: Hand of Edward the Odd; off_hand: Elunarian Sphere

No-known-source sample (15 of 1792, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 553.1. Weights run: 4.1s. Verify run: 30.6s. 1792 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.513, intellect=1.394 ± 0.068, spirit=0.575 ± 0.067, mp5=3.397 ± 0.085, crit=0.929 ± 0.079 per rating point (14 rating = 1%, 13.004 per %), spell_haste=not significant (0.682 ± 1.061)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 106.2 healing_power points (24.66 DPS) | yes | Lieutenant Commander's Lamellar Helmet (227149, -3.73 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, -5.50 DPS) [vendor]; Crown of the Penitent (13216, -8.62 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 31.2 healing_power points (7.24 DPS) | yes | Horizon Choker (13085, -2.17 DPS) [world_drop]; Opaline Medallion (16623, -3.03 DPS) [quest]; Archlight Talisman (15856, -4.00 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, +0.00 DPS) [vendor]; Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 17.6 healing_power points (4.10 DPS) | yes | Sergeant Major's Cape (16337, -0.44 DPS) [pvp]; Highborne Cloak (14450, -0.65 DPS) [world_drop]; Keeper's Cloak (14665, -0.86 DPS) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (553.2 DPS) | yes | Knight-Captain's Lamellar Chestplate (227151, +0.00 DPS) [vendor]; Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Breastplate of Salvation (250601, +0.00 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prefect's Wristguards (250584, +0.00 DPS) [crafted]; Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Prefect's Gloves (250554, +0.00 DPS) [crafted]; Mender's Leather Gauntlets (252551, +0.00 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Mender's Mail Belt (252591, +0.00 DPS) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Legguards (227150, -3.43 DPS) [vendor]; Marshal's Lamellar Legguards (231639, -6.50 DPS) [vendor]; Red Dragonscale Leggings (252603, -8.88 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | sim-verified (553.2 DPS) | yes | Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Marshal's Lamellar Greaves (231646, -1.55 DPS) [vendor]; Gilded Sandals (254107, -1.73 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 20.3 healing_power points (4.71 DPS) | yes | Eye of Adaegus (5266, -1.16 DPS) [world_drop]; Frigid Ring (18679, -1.47 DPS) [world]; Maiden's Circle (13001, -1.64 DPS) [world_drop] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 16.7 healing_power points (3.89 DPS) | yes | Eye of Adaegus (5266, -0.34 DPS) [world_drop]; Frigid Ring (18679, -0.65 DPS) [world]; Maiden's Circle (13001, -0.82 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -3.13 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.40 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.67 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -2.58 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.85 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.11 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Inventor's Focal Sword (17719, +0.00 DPS) [dungeon]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | 36.0 healing_power points (8.36 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, -0.48 DPS) [vendor]; Orb of Lorica (11262, -2.34 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Glowing Crystal Ring; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Evergreen Shield

No-known-source sample (15 of 1792, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (undead, 50320100000000000-00000000000000000-00000000000000000)

Set DPS (verified): 28.0. Weights run: 4.0s. Verify run: 14.6s. 252 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.003 ± 0.002, spirit=0.518 ± 0.003, mp5=1.938 ± 0.012, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), spell_haste=not significant (0.016 ± 0.005)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.82 DPS) | yes | Wisdom's Leather Hood (252507, -0.35 DPS) [crafted]; Pristine Circlet (253949, -0.36 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.12 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.6 healing_power points (0.14 DPS) | yes | Scholarly Pendant (277203, -0.20 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.0 healing_power points (0.78 DPS) | yes | Slime-encrusted Pads (6461, -0.28 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.44 DPS) [crafted]; Serpent's Shoulders (5404, -0.47 DPS) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 4.0 healing_power points (0.35 DPS) | yes | Seer's Cape (6378, -0.08 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.09 DPS) [crafted]; Spirit Cloak (4792, -0.17 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 21.6 healing_power points (1.87 DPS) | yes | Filigreed Pristine Gown (253901, -0.17 DPS) [crafted]; Acolyte's Chain Shirt (250491, -0.21 DPS, sim-verified) [crafted]; Corsair's Overshirt (5202, -0.56 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.8 healing_power points (0.94 DPS) | yes | Bright Bracers (3647, -0.59 DPS) [world_drop]; Drakewing Bands (12999, -0.67 DPS) [world_drop]; Owl Bracers (4796, -1.85 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 15.0 healing_power points (1.30 DPS) | yes | Pristine Gloves (253913, -0.09 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.19 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.69 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.0 healing_power points (1.30 DPS) | yes | Acolyte's Chain Belt (250516, -0.24 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.26 DPS) [crafted]; Novice Ardent's Sash (253887, -0.39 DPS) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 28.1 healing_power points (2.44 DPS) | yes | Acolyte's Chain Leggings (250496, -0.15 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.17 DPS) [crafted]; Darkweave Breeches (12987, -1.56 DPS) [world_drop] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 16.0 healing_power points (1.39 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.26 DPS) [crafted]; Black Whelp Slippers (252424, -0.26 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.0 healing_power points (0.52 DPS) | yes | Volcanic Rock Ring (12053, -0.26 DPS) [world_drop]; Furlfeather Ring (286745, -0.35 DPS) [world]; Loop of Sacrifice (281673, -0.71 DPS, sim-verified) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | sim-verified (28.0 DPS) | yes | Volcanic Rock Ring (12053, -0.01 DPS) [world_drop]; Furlfeather Ring (286745, -0.10 DPS) [world]; Loop of Sacrifice (281673, -0.28 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 9.0 healing_power points (0.78 DPS) | yes | Hammerbone (270018, +0.00 DPS, sim-verified) [quest]; Impaling Harpoon (5200, -0.35 DPS) [dungeon]; Samophlange Screwdriver (11854, -0.51 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Sword of the Fallen

No-known-source sample (15 of 252, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band

### Band 30 (undead, 50320215012000000-00000000000000000-00000000000000000)

Set DPS (verified): 49.7. Weights run: 4.4s. Verify run: 17.0s. 473 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=1.315 ± 0.009, spirit=0.661 ± 0.005, mp5=2.497 ± 0.009, crit=0.110 ± 0.007 per rating point (14 rating = 1%, 1.537 per %), spell_haste=not significant (0.032 ± 0.016)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 37.0 healing_power points (2.99 DPS) | yes | Acolyte's Chain Helm (250501, -0.97 DPS) [crafted]; Enduring Cap (3020, -1.19 DPS, sim-verified) [world_drop]; Filigreed Pristine Circlet (253975, -1.21 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.9 healing_power points (0.64 DPS) | yes | Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Pendant of Myzrael (4614, -0.32 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -0.61 DPS, sim-verified) [world] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 14.5 healing_power points (1.17 DPS) | yes | Bloodmage Mantle (7684, -0.21 DPS) [dungeon]; Resilient Mantle (14397, -0.21 DPS) [world_drop]; Nightsky Mantle (4718, -0.31 DPS, sim-verified) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.3 healing_power points (0.99 DPS) | yes | Silver-thread Cloak (4713, -0.46 DPS) [world_drop]; Soft Willow Cape (16661, -0.46 DPS) [quest]; Resilient Cape (14400, -1.00 DPS, sim-verified) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 34.5 healing_power points (2.79 DPS) | yes | Acolyte's Silvered Chain Shirt (250521, -0.26 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.37 DPS) [crafted]; Pristine Gown (253961, -0.50 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.1 healing_power points (1.14 DPS) | yes | Nightsky Wristbands (6407, -0.34 DPS) [world_drop]; Spidertank Oilrag (9448, -0.53 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.38 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.9 healing_power points (1.53 DPS) | yes | Acolyte's Gloves (250511, -0.29 DPS, sim-verified) [crafted]; Wisdom's Leather Gloves (252499, -0.30 DPS) [crafted]; Naga Battle Gloves (888, -0.32 DPS) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 31.2 healing_power points (2.53 DPS) | yes | Mender's Leather Belt (252523, -0.43 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.21 DPS) [crafted]; Silver-lined Belt (13011, -1.35 DPS) [world_drop] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 35.5 healing_power points (2.87 DPS) | yes | Pristine Leggings (253987, -0.08 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.40 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.51 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 31.9 healing_power points (2.58 DPS) | yes | Wisdom's Leather Boots (252444, -1.16 DPS) [crafted]; Glowing Copper Boots (250482, -1.42 DPS) [crafted]; Acolyte's Boots (250506, -1.91 DPS, sim-verified) [crafted] |
| finger1 | Solid Gold Band (279845) | Opportunistic Education [quest] | 14.3 healing_power points (1.16 DPS) | yes | Runemender's Seal (273031, -0.05 DPS) [world]; The Queen's Jewel (13094, -0.52 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.52 DPS) [dungeon] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 13.6 healing_power points (1.10 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.46 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (49.7 DPS) | yes | Pysan's Old Greatsword (1975, -4.13 DPS, sim-verified) [dungeon]; Wolfsbane (267369, -4.44 DPS) [quest]; Sword of the Fallen (286977, -4.55 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Death Speaker Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Solid Gold Band; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter

No-known-source sample (15 of 473, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (undead, 50320215012101251-00000000000000000-00000000000000000)

Set DPS (verified): 76.9. Weights run: 6.8s. Verify run: 28.8s. 658 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.044, intellect=1.533 ± 0.014, spirit=0.963 ± 0.008, mp5=3.277 ± 0.012, crit=0.227 ± 0.014 per rating point (14 rating = 1%, 3.176 per %), spell_haste=not significant (0.025 ± 0.030)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 56.7 healing_power points (4.13 DPS) | yes | Whitemane's Chapeau (7720, -0.72 DPS) [dungeon]; Holy Shroud (2721, -1.31 DPS) [world_drop]; Earthen Silk Hood (254015, -1.83 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 32.3 healing_power points (2.35 DPS) | yes | Necklace of Calisea (1714, -1.08 DPS) [dungeon]; Kaleidoscope Chain (13084, -1.62 DPS) [world_drop]; Amberglow Talisman (10824, -1.65 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 41.9 healing_power points (3.05 DPS) | yes | Sheepshear Mantle (13115, -0.85 DPS) [world_drop]; Earthen Silk Shoulders (254033, -1.18 DPS) [crafted]; Mistscape Mantle (4734, -1.47 DPS) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 16.1 healing_power points (1.17 DPS) | yes | Glowing Thresher Cape (6901, -0.10 DPS) [dungeon]; Windchaser Cloak (14430, -0.22 DPS) [world_drop]; Cloak of Blight (6832, -0.26 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.7 healing_power points (4.78 DPS) | yes | Deathchill Armor (10764, -1.27 DPS) [dungeon]; Civinad Robes (9623, -2.05 DPS) [quest]; Death Speaker Robes (6682, -2.10 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 17.5 healing_power points (1.27 DPS) | yes | Enchanted Kodo Bracers (13119, -0.13 DPS) [world_drop]; Earthen Silk Cuffs (254019, -0.20 DPS) [crafted]; Windchaser Cuffs (14429, -0.27 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 34.8 healing_power points (2.53 DPS) | yes | Gilded Handwraps (254021, -0.15 DPS) [crafted]; Prefect's Gauntlet (250569, -0.22 DPS) [crafted]; Stormcloth Gloves (10011, -0.39 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 40.0 healing_power points (2.91 DPS) | yes | Mender's Leather Belt (252523, -0.44 DPS) [crafted]; Prefect's Belt (250559, -0.53 DPS) [crafted]; Deathmage Sash (10771, -1.24 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | sim-verified (76.9 DPS) | yes | Warchief Kilt (7760, +0.00 DPS) [dungeon]; Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Pristine Leggings (253987, -0.07 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 41.8 healing_power points (3.04 DPS) | yes | Mender's Leather Shoes (252533, -0.02 DPS) [crafted]; Mender's Mail Boots (252565, -0.02 DPS) [crafted]; Gilded Slippers (254001, -0.52 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | sim-verified (76.9 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, +0.00 DPS) [quest] |
| finger2 | Voodoo Band (1996) | Bloodscalp Witch Doctor [world] | sim-verified (76.9 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, +0.00 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (76.9 DPS) | yes | Pysan's Old Greatsword (1975, +0.00 DPS) [dungeon]; Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Mograine's Might (7723, +0.00 DPS) [dungeon] |
| off_hand | Forcestone Buckler (17508) | Compendium of the Fallen [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beacon of Hope (9393, +0.00 DPS) [dungeon]; Battered Viking Shield (9403, -0.22 DPS) [dungeon]; Orb of Souls (249395, -0.50 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Wisdom's Leather Leggings; feet: Prefect's Boots; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Forcestone Buckler

No-known-source sample (15 of 658, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (undead, 50320215012101251-50320000000000000-00000000000000000)

Set DPS (verified): 100.7. Weights run: 7.0s. Verify run: 35.1s. 868 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.050, intellect=1.958 ± 0.018, spirit=1.458 ± 0.012, mp5=4.543 ± 0.016, crit=0.346 ± 0.019 per rating point (14 rating = 1%, 4.851 per %), spell_haste=not significant (0.114 ± 0.046)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 68.5 healing_power points (4.89 DPS) | yes | Papal Fez (9431, -0.01 DPS) [dungeon]; Winged Helm (13112, -0.64 DPS) [world_drop]; Stormcloth Headband (10032, -0.74 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (100.8 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Necklace of Calisea (1714, -0.67 DPS) [dungeon]; Choker of the High Shaman (4112, -0.84 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (100.8 DPS) | yes | Mender's Leather Shoulders (252538, +0.00 DPS) [crafted]; Mender's Mail Shoulders (252569, +0.00 DPS) [crafted]; Nethergeld Shoulders (254049, -0.03 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 27.4 healing_power points (1.95 DPS) | yes | Highborne Cloak (14450, -0.28 DPS) [world_drop]; Blackforge Cape (6424, -0.42 DPS) [dungeon]; Wingveil Cloak (10802, -0.45 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 69.7 healing_power points (4.97 DPS) | yes | Ghostweave Vest (14141, -0.30 DPS) [crafted]; Deathchill Armor (10764, -0.73 DPS) [dungeon]; Grimlok's Tribal Vestments (9415, -0.75 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 41.5 healing_power points (2.96 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.20 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 61.7 healing_power points (4.41 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.11 DPS) [crafted]; Prefect's Gloves (250554, -0.37 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 53.5 healing_power points (3.82 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.28 DPS) [crafted]; Dawnspire Cord (12466, -0.31 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 67.5 healing_power points (4.82 DPS) | yes | Senior Designer's Pantaloons (11841, -0.18 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.30 DPS) [dungeon]; Warchief Kilt (7760, -0.81 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 49.5 healing_power points (3.54 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.07 DPS) [crafted]; Mender's Leather Shoes (252533, -0.12 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 29.2 healing_power points (2.09 DPS) | yes | Electrocutioner Lagnut (9447, -0.72 DPS) [dungeon]; Runemender's Seal (273031, -0.72 DPS) [world]; Coldwater Ring (4550, -0.73 DPS) [quest] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 20.0 healing_power points (1.43 DPS) | yes | Electrocutioner Lagnut (9447, -0.06 DPS) [dungeon]; Runemender's Seal (273031, -0.06 DPS) [world]; Coldwater Ring (4550, -0.07 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.24 DPS) [quest]; Alchemist's Stone (13503, -0.87 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.21 DPS) [quest]; Alchemist's Stone (13503, -0.83 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (100.8 DPS) | yes | Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Inventor's Focal Sword (17719, +0.00 DPS) [dungeon] |
| off_hand | Protector Buckler (14790) | World drop [world_drop] | sim-verified (100.8 DPS) | yes | Beacon of Hope (9393, +0.00 DPS) [dungeon]; Forcestone Buckler (17508, +0.00 DPS) [quest]; Twisting Essence Jar (249456, +0.00 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Mender's Leather Boots; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life; off_hand: Protector Buckler

No-known-source sample (15 of 868, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 160.1. Weights run: 6.9s. Verify run: 84.4s. 1811 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.049, intellect=3.366 ± 0.040, spirit=1.418 ± 0.013, mp5=5.083 ± 0.026, crit=0.400 ± 0.022 per rating point (14 rating = 1%, 5.595 per %), spell_haste=not significant (0.063 ± 0.073)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 153.4 healing_power points (12.60 DPS) | yes | Crown of the Penitent (13216, -1.75 DPS) [quest]; Devout Crown (16693, -2.58 DPS) [dungeon]; Black Dragonscale Helm (252605, -6.12 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 52.8 healing_power points (4.34 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Archlight Talisman (15856, -1.57 DPS) [quest]; Necklace of Calisea (1714, -1.59 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 94.4 healing_power points (7.76 DPS) | yes | Magister's Mantle (16689, -1.67 DPS) [dungeon]; Shroud of the Nathrezim (18720, -1.94 DPS) [dungeon]; Ironfeather Shoulders (15067, -9.44 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 42.7 healing_power points (3.51 DPS) | yes | Keeper's Cloak (14665, -0.74 DPS) [world_drop]; Highborne Cloak (14450, -0.80 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -0.83 DPS) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (160.1 DPS) | yes | Mooncloth Vest (14138, -0.15 DPS) [crafted]; Mooncloth Robe (18486, -0.98 DPS) [crafted]; Breastplate of Salvation (250601, -4.71 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (160.1 DPS) | yes | Magister's Bindings (16683, +0.00 DPS) [dungeon]; Mender's Leather Bracers (252543, -0.09 DPS) [crafted]; Magiskull Cuffs (13107, -8.17 DPS, sim-verified) [world_drop] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (160.1 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, +0.00 DPS) [crafted]; Mender's Leather Gauntlets (252551, -5.92 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 91.1 healing_power points (7.48 DPS) | yes | Elunarian Belt (14465, -0.79 DPS) [world_drop]; Dawnspire Cord (12466, -1.24 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -6.04 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-verified (160.1 DPS) | yes | Devout Skirt (16694, -3.17 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -4.40 DPS) [dungeon]; Red Dragonscale Leggings (252603, -7.77 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (160.1 DPS) | yes | Mender's Leather Boots (252472, -0.31 DPS) [crafted]; Mender's Mail Sabatons (252579, -0.31 DPS) [crafted]; Mooncloth Boots (15802, -4.42 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 49.3 healing_power points (4.04 DPS) | yes | Chemist's Ring (15702, -0.82 DPS) [quest]; Eye of Adaegus (5266, -0.99 DPS) [world_drop]; Frigid Ring (18679, -1.28 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 40.4 healing_power points (3.32 DPS) | yes | Chemist's Ring (15702, -0.09 DPS) [quest]; Eye of Adaegus (5266, -0.26 DPS) [world_drop]; Frigid Ring (18679, -0.55 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (160.1 DPS) | yes | Ankh of Life (1713, -1.11 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.34 DPS) [quest]; Alchemist's Stone (13503, -2.04 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (160.1 DPS) | yes | Ankh of Life (1713, -0.08 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.31 DPS) [quest]; Mark of the Chosen (17774, -2.75 DPS, sim-verified) [quest] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (160.1 DPS) | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Inventor's Focal Sword (17719, -0.20 DPS) [dungeon]; Hand of Edward the Odd (2243, -0.83 DPS) [world_drop] |
| off_hand | Elunarian Sphere (15968) | World drop [world_drop] | 47.1 healing_power points (3.87 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Arcane Star (15931, -0.17 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Glowing Crystal Ring; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; main_hand: Death Speaker Scepter; off_hand: Elunarian Sphere

No-known-source sample (15 of 1811, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 543.8. Weights run: 4.1s. Verify run: 31.2s. 1811 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.513, intellect=1.394 ± 0.068, spirit=0.575 ± 0.067, mp5=3.397 ± 0.085, crit=0.929 ± 0.079 per rating point (14 rating = 1%, 13.004 per %), spell_haste=not significant (0.682 ± 1.061)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 106.2 healing_power points (24.66 DPS) | yes | Crown of the Penitent (13216, -8.62 DPS) [quest]; Black Dragonscale Helm (252605, -9.33 DPS) [crafted]; Whitesoul Helm (12633, -9.66 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 31.2 healing_power points (7.24 DPS) | yes | Horizon Choker (13085, -2.17 DPS) [world_drop]; Archlight Talisman (15856, -4.00 DPS) [quest]; Necklace of Calisea (1714, -4.03 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawnbringer Shoulders (12625, +0.00 DPS) [crafted]; Mooncloth Shoulders (14139, -0.50 DPS) [crafted]; Prefect's Pauldrons (250579, -0.58 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 17.6 healing_power points (4.10 DPS) | yes | Highborne Cloak (14450, -0.65 DPS) [world_drop]; Keeper's Cloak (14665, -0.86 DPS) [world_drop]; Blackforge Cape (6424, -0.97 DPS) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (543.9 DPS) | yes | Stormcloth Vest (10020, +0.00 DPS) [crafted]; Mooncloth Vest (14138, +0.00 DPS) [crafted]; Breastplate of Salvation (250601, +0.00 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prefect's Wristguards (250584, +0.00 DPS) [crafted]; Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prefect's Gloves (250554, +0.00 DPS) [crafted]; Mender's Leather Gauntlets (252551, +0.00 DPS) [crafted]; Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Mender's Mail Belt (252591, +0.00 DPS) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Dragonscale Leggings (252603, -8.88 DPS) [crafted]; Devout Skirt (16694, -10.39 DPS) [dungeon]; Elderwild Pants (279254, -11.40 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | sim-verified (543.9 DPS) | yes | Gilded Sandals (254107, -1.73 DPS) [crafted]; Mender's Leather Boots (252472, -1.78 DPS) [crafted]; Mender's Mail Sabatons (252579, -1.78 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 20.3 healing_power points (4.71 DPS) | yes | Chemist's Ring (15702, -0.93 DPS) [quest]; Eye of Adaegus (5266, -1.16 DPS) [world_drop]; Frigid Ring (18679, -1.47 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 16.7 healing_power points (3.89 DPS) | yes | Chemist's Ring (15702, -0.11 DPS) [quest]; Eye of Adaegus (5266, -0.34 DPS) [world_drop]; Frigid Ring (18679, -0.65 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -3.13 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.40 DPS) [quest]; Alchemist's Stone (13503, -4.20 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -2.58 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.85 DPS) [quest]; Alchemist's Stone (13503, -3.65 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Inventor's Focal Sword (17719, +0.00 DPS) [dungeon]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | 36.0 healing_power points (8.36 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, -0.48 DPS) [vendor]; Beacon of Hope (9393, -2.72 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Glowing Crystal Ring; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Evergreen Shield

No-known-source sample (15 of 1811, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

