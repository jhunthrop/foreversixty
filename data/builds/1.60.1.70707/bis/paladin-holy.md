# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 50320100000000000-00000000000000000-00000000000000000)

Set DPS (verified): 29.0. Weights run: 5.4s. Verify run: 18.8s. 232 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.002, intellect=1.018 ± 0.002, spirit=0.462 ± 0.002, mp5=1.888 ± 0.011, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.583 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.84 DPS) | yes | Wisdom's Leather Hood (252507, -0.35 DPS) [crafted]; Pristine Circlet (253949, -0.37 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.12 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 1.8 healing_power points (0.16 DPS) | yes | Scholarly Pendant (277203, -0.31 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.2 healing_power points (0.80 DPS) | yes | Slime-encrusted Pads (6461, -0.25 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.44 DPS) [crafted]; Forest Leather Mantle (4709, -0.64 DPS) [world_drop] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 4.1 healing_power points (0.36 DPS) | yes | Seer's Cape (6378, -0.10 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.17 DPS, sim-verified) [crafted]; Forest Cloak (4710, -0.18 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 21.5 healing_power points (1.88 DPS) | yes | Filigreed Pristine Gown (253901, -0.17 DPS) [crafted]; Acolyte's Chain Shirt (250491, -0.23 DPS, sim-verified) [crafted]; Corsair's Overshirt (5202, -0.59 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.8 healing_power points (0.94 DPS) | yes | Bright Bracers (3647, -0.58 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.67 DPS) [dungeon]; Owl Bracers (4796, -1.96 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 15.1 healing_power points (1.32 DPS) | yes | Pristine Gloves (253913, -0.09 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.17 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.71 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.1 healing_power points (1.32 DPS) | yes | Acolyte's Chain Belt (250516, -0.25 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.26 DPS) [crafted]; Novice Ardent's Sash (253887, -0.40 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 28.1 healing_power points (2.46 DPS) | yes | Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.37 DPS, sim-verified) [crafted]; Dreamer's Leggings (270016, -1.48 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 16.1 healing_power points (1.41 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.26 DPS) [crafted]; Black Whelp Slippers (252424, -0.26 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.1 healing_power points (0.53 DPS) | yes | Deep Fathom Ring (6463, -0.33 DPS) [dungeon]; Minor Channeling Ring (1449, -0.36 DPS) [quest]; Volcanic Rock Ring (12053, -0.80 DPS, sim-verified) [world_drop] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | sim-verified (29.0 DPS) | yes | Deep Fathom Ring (6463, -0.04 DPS) [dungeon]; Minor Channeling Ring (1449, -0.06 DPS) [quest]; Volcanic Rock Ring (12053, -0.62 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 10.7 healing_power points (0.94 DPS) | yes | Impaling Harpoon (5200, -0.49 DPS) [dungeon]; Trogg Slicer (6186, -0.74 DPS) [quest]; Sword of the Fallen (286977, -0.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Verigan's Fist

No-known-source sample (15 of 232, see the JSON for more): 1189 Overseer's Ring; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (human, 50320215012000000-00000000000000000-00000000000000000)

Set DPS (verified): 44.7. Weights run: 6.0s. Verify run: 21.8s. 454 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=1.305 ± 0.009, spirit=0.626 ± 0.005, mp5=2.511 ± 0.010, crit=0.116 ± 0.007 per rating point (14 rating = 1%, 1.628 per %), spell_haste=not significant (0.067 ± 0.018)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.8 healing_power points (3.00 DPS) | yes | Acolyte's Chain Helm (250501, -0.96 DPS) [crafted]; Filigreed Pristine Circlet (253975, -1.21 DPS) [crafted]; Enduring Cap (3020, -1.60 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 7.8 healing_power points (0.64 DPS) | yes | Crystal Starfire Medallion (5003, -0.06 DPS) [world_drop]; Pendant of Myzrael (4614, -0.33 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.42 DPS, sim-verified) [world_drop] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 14.4 healing_power points (1.17 DPS) | yes | Mantle of Honor (3560, -0.07 DPS) [quest]; Talbar Mantle (10657, -0.12 DPS) [quest]; Nightsky Mantle (4718, -0.73 DPS, sim-verified) [world_drop] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 14.8 healing_power points (1.21 DPS) | yes | Repairman's Cape (9605, -0.30 DPS) [quest]; Resilient Cape (14400, -0.57 DPS) [world_drop]; Glowing Thresher Cape (6901, -0.63 DPS, sim-verified) [dungeon] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 34.4 healing_power points (2.81 DPS) | yes | Acolyte's Silvered Chain Shirt (250521, -0.26 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.39 DPS) [crafted]; Pristine Gown (253961, -0.73 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.1 healing_power points (1.15 DPS) | yes | Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Technician's Bracers (270042, -0.51 DPS) [quest]; Glowing Magical Bracelets (13106, -1.44 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.9 healing_power points (1.54 DPS) | yes | Acolyte's Gloves (250511, -0.25 DPS, sim-verified) [crafted]; Wisdom's Leather Gloves (252499, -0.30 DPS) [crafted]; Naga Battle Gloves (888, -0.32 DPS) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 31.1 healing_power points (2.54 DPS) | yes | Mender's Leather Belt (252523, -0.74 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.22 DPS) [crafted]; Silver-lined Belt (13011, -1.38 DPS) [world_drop] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 35.4 healing_power points (2.89 DPS) | yes | Pristine Leggings (253987, -0.10 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.42 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.71 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 31.6 healing_power points (2.58 DPS) | yes | Wisdom's Leather Boots (252444, -1.15 DPS) [crafted]; Nimbus Boots (6998, -1.20 DPS) [quest]; Acolyte's Boots (250506, -2.14 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 healing_power points (0.85 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; The Queen's Jewel (13094, -0.23 DPS) [world_drop]; Ring of Power Regulation (273030, -0.23 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 10.4 healing_power points (0.85 DPS) | yes | The Queen's Jewel (13094, -0.23 DPS) [world_drop]; Ring of Power Regulation (273030, -0.23 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.92 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (44.7 DPS) | yes | Manual Crowd Pummeler (9449, -2.26 DPS, sim-verified) [dungeon]; Pysan's Old Greatsword (1975, -4.00 DPS) [dungeon]; Mechanic's Pipehammer (9604, -4.27 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.8 healing_power points (1.37 DPS) | yes | Orb of Mistmantle (13031, -0.43 DPS) [world_drop]; Heart of Agamaggan (6694, -0.55 DPS) [dungeon]; Eye of Paleth (2943, -1.01 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Prelacy Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 454, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (human, 50320215012101251-00000000000000000-00000000000000000)

Set DPS (verified): 65.2. Weights run: 9.1s. Verify run: 35.5s. 643 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.026, intellect=1.437 ± 0.014, spirit=0.872 ± 0.008, mp5=3.076 ± 0.013, crit=0.214 ± 0.012 per rating point (14 rating = 1%, 2.994 per %), spell_haste=not significant (0.037 ± 0.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 46.3 healing_power points (3.62 DPS) | yes | Whitemane's Chapeau (7720, -0.66 DPS) [dungeon]; Holy Shroud (2721, -0.77 DPS, sim-verified) [world_drop]; Earthen Silk Hood (254015, -1.22 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.5 healing_power points (1.91 DPS) | yes | Necklace of Calisea (1714, -1.04 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -1.19 DPS) [world_drop]; Amberglow Talisman (10824, -1.23 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 33.7 healing_power points (2.63 DPS) | yes | Sheepshear Mantle (13115, -0.66 DPS, sim-verified) [world_drop]; Earthen Silk Shoulders (254033, -0.68 DPS) [crafted]; Mistscape Mantle (4734, -1.06 DPS) [dungeon] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | 16.2 healing_power points (1.27 DPS) | yes | Blackforge Cape (6424, -0.10 DPS) [dungeon]; Glowing Thresher Cape (6901, -0.18 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.19 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.0 healing_power points (5.08 DPS) | yes | Death Speaker Robes (6682, -2.28 DPS) [dungeon]; Pristine Gown (253961, -2.46 DPS) [crafted]; Deathchill Armor (10764, -3.18 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.4 healing_power points (1.28 DPS) | yes | Earthen Silk Cuffs (254019, -0.16 DPS) [crafted]; Windchaser Cuffs (14429, -0.27 DPS) [world_drop]; Enchanted Kodo Bracers (13119, -0.53 DPS, sim-verified) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 33.9 healing_power points (2.65 DPS) | yes | Gilded Handwraps (254021, -0.21 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.22 DPS) [crafted]; Stormcloth Gloves (10011, -0.44 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 38.7 healing_power points (3.03 DPS) | yes | Prefect's Belt (250559, -0.52 DPS) [crafted]; Mender's Leather Belt (252523, -0.66 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.34 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 37.2 healing_power points (2.91 DPS) | yes | Acolyte's Silvered Chain Leggings (250526, -0.05 DPS) [crafted]; Pristine Leggings (253987, -0.06 DPS) [crafted]; Warchief Kilt (7760, -0.13 DPS, sim-verified) [dungeon] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 40.9 healing_power points (3.20 DPS) | yes | Mender's Mail Boots (252565, -0.06 DPS) [crafted]; Mender's Leather Shoes (252533, -0.26 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -0.58 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 13.3 healing_power points (1.04 DPS) | yes | Ogremind Ring (1993, -0.05 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.09 DPS) [dungeon]; Runemender's Seal (273031, -0.09 DPS) [world] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 12.7 healing_power points (0.99 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.04 DPS) [dungeon]; Runemender's Seal (273031, -0.04 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (65.2 DPS) | yes | Frost Tiger Blade (3854, -2.22 DPS, sim-verified) [crafted]; Death Speaker Scepter (2816, -2.58 DPS) [dungeon]; Mograine's Might (7723, -6.10 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 26.4 healing_power points (2.06 DPS) | yes | Beacon of Hope (9393, -0.69 DPS) [dungeon]; Battered Viking Shield (9403, -0.75 DPS) [dungeon]; Orb of Souls (249395, -0.87 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; legs: Wisdom's Leather Leggings; feet: Prefect's Boots; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 643, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (human, 50320215012101251-50320000000000000-00000000000000000)

Set DPS (verified): 94.1. Weights run: 9.4s. Verify run: 44.4s. 840 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.023, intellect=2.277 ± 0.024, spirit=1.533 ± 0.013, mp5=5.131 ± 0.019, crit=0.365 ± 0.018 per rating point (14 rating = 1%, 5.113 per %), spell_haste=not significant (0.137 ± 0.041)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 75.3 healing_power points (4.98 DPS) | yes | Knight-Lieutenant's Imbued Helmet (220810, +0.00 DPS) [vendor]; Papal Fez (9431, -0.58 DPS) [dungeon]; Winged Helm (13112, -0.95 DPS, sim-verified) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 38.0 healing_power points (2.51 DPS) | yes | Glowing Eye of Mordresh (10769, -0.27 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.75 DPS) [dungeon]; Mindburst Medallion (11196, -1.10 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 57.8 healing_power points (3.82 DPS) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, +0.00 DPS) [vendor]; Mender's Leather Shoulders (252538, +0.00 DPS, sim-verified) [crafted]; Mender's Mail Shoulders (252569, -0.17 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 31.2 healing_power points (2.06 DPS) | yes | Highborne Cloak (14450, -0.36 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -0.45 DPS) [dungeon]; Spritecaster Cape (11623, -0.55 DPS) [dungeon] |
| chest | Knight's Imbued Armor (220813) | Captain Dirgehammer [vendor] | 79.1 healing_power points (5.23 DPS) | yes | Stormcloth Vest (10020, -0.59 DPS) [crafted]; Earth Warder's Vest (21311, -1.24 DPS) [quest]; Ghostweave Vest (14141, -1.63 DPS, sim-verified) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 44.1 healing_power points (2.92 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.04 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 65.3 healing_power points (4.32 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.08 DPS) [crafted]; Prefect's Gloves (250554, -0.35 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 57.3 healing_power points (3.79 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.14 DPS) [dungeon]; Prefect's Waistguard (250574, -0.30 DPS) [crafted] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 79.1 healing_power points (5.23 DPS) | yes | Senior Designer's Pantaloons (11841, -0.70 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.74 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.45 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 53.1 healing_power points (3.51 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.10 DPS) [crafted]; Mender's Leather Shoes (252533, -0.16 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 32.1 healing_power points (2.12 DPS) | yes | Coldwater Ring (4550, -0.76 DPS) [quest]; Ogremind Ring (1993, -0.76 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 22.1 healing_power points (1.46 DPS) | yes | Ogremind Ring (1993, -0.10 DPS) [world_drop]; Voodoo Band (1996, -0.10 DPS) [world]; Coldwater Ring (4550, -0.22 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (94.1 DPS) | yes | Evonice's Landin' Pilla (18951, -0.18 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.38 DPS) [quest]; Alchemist's Stone (13503, -0.78 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (94.1 DPS) | yes | Evonice's Landin' Pilla (18951, -0.28 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -0.41 DPS) [quest]; Alchemist's Stone (13503, -0.81 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (94.1 DPS) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.53 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.06 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 34.8 healing_power points (2.30 DPS) | yes | Twisting Essence Jar (249456, -0.24 DPS, sim-verified) [crafted]; Protector Buckler (14790, -0.29 DPS) [world_drop]; Mountainside Buckler (13082, -0.33 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; chest: Knight's Imbued Armor; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Mender's Leather Boots; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life

No-known-source sample (15 of 840, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 166.4. Weights run: 9.1s. Verify run: 121.3s. 1762 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.198, intellect=5.732 ± 0.063, spirit=2.168 ± 0.024, mp5=8.318 ± 0.046, crit=0.613 ± 0.038 per rating point (14 rating = 1%, 8.586 per %), spell_haste=not significant (-0.021 ± 0.085)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (166.4 DPS) | yes | Crown of the Penitent (13216, +0.00 DPS) [quest]; Black Dragonscale Helm (252605, +0.00 DPS) [crafted]; Lieutenant Commander's Lamellar Helmet (227149, -0.71 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 88.9 healing_power points (4.60 DPS) | yes | Opaline Medallion (16623, -0.74 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.55 DPS) [dungeon]; Archlight Talisman (15856, -1.63 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 150.9 healing_power points (7.80 DPS) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, -0.45 DPS) [vendor]; Ironfeather Shoulders (15067, -0.98 DPS) [crafted]; Magister's Mantle (16689, -1.28 DPS) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 71.7 healing_power points (3.71 DPS) | yes | Sergeant Major's Cape (16337, -0.44 DPS) [pvp]; Highborne Cloak (14450, -0.59 DPS) [world_drop]; Keeper's Cloak (14665, -0.74 DPS) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Breastplate of Salvation (250601, +0.00 DPS) [crafted]; Knight-Captain's Lamellar Chestplate (227151, -0.26 DPS) [vendor]; Field Marshal's Lamellar Chestplate (231641, -0.32 DPS) [pvp] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; Indomitable Armguards (14682, +0.00 DPS) [world_drop]; Magister's Bindings (16683, +0.00 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Gloves (16684, +0.00 DPS) [dungeon]; Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Gilded Gloves (254095, -0.29 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 145.2 healing_power points (7.51 DPS) | yes | Elunarian Belt (14465, -0.46 DPS) [world_drop]; Wisdom of the Timbermaw (19047, -0.74 DPS) [crafted]; Marshal's Lamellar Belt (16470, -0.97 DPS) [pvp] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Dragonscale Leggings (252603, +0.00 DPS) [crafted]; Knight-Captain's Lamellar Legguards (227150, -0.08 DPS) [vendor]; Devout Skirt (16694, -1.68 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 81.2 healing_power points (4.20 DPS) | yes | Glowing Crystal Ring (18402, -0.64 DPS) [quest]; Frigid Ring (18679, -1.23 DPS) [world]; Maiden's Circle (13001, -1.45 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | sim-verified (166.4 DPS) | yes | Glowing Crystal Ring (18402, +0.00 DPS) [quest]; Frigid Ring (18679, -0.16 DPS) [world]; Maiden's Circle (13001, -0.38 DPS) [world_drop] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (166.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.03 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (166.4 DPS) | yes | Ankh of Life (1713, -1.24 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.46 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.68 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Demolisher (234568, -3.26 DPS) [pvp]; Hand of Righteousness (7721, -3.87 DPS) [dungeon]; Grand Marshal's Warhammer (23454, -4.68 DPS) [vendor] |
| off_hand | Elunarian Sphere (15968) | World drop [world_drop] | 80.3 healing_power points (4.15 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, -0.15 DPS) [vendor]; Arcane Star (15931, -0.29 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves; main_hand: Hand of Edward the Odd; off_hand: Elunarian Sphere

No-known-source sample (15 of 1762, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (human, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 538.5. Weights run: 5.3s. Verify run: 38.8s. 1762 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.561, intellect=1.727 ± 0.079, spirit=0.574 ± 0.077, mp5=3.637 ± 0.105, crit=0.984 ± 0.089 per rating point (14 rating = 1%, 13.772 per %), spell_haste=not significant (2.709 ± 1.191)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 112.1 healing_power points (21.40 DPS) | yes | Lieutenant Commander's Lamellar Helmet (227149, -2.83 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, -4.04 DPS) [vendor]; Black Dragonscale Helm (252605, -6.43 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.8 healing_power points (5.12 DPS) | yes | Opaline Medallion (16623, -0.83 DPS) [quest]; Archlight Talisman (15856, -1.82 DPS) [quest]; Horizon Choker (13085, -1.88 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 52.4 healing_power points (10.01 DPS) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Field Marshal's Lamellar Pauldrons (231645, +0.00 DPS) [pvp]; Dawnbringer Shoulders (12625, -24.00 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 21.3 healing_power points (4.06 DPS) | yes | Sergeant Major's Cape (16337, -0.55 DPS) [pvp]; Highborne Cloak (14450, -0.66 DPS) [world_drop]; Keeper's Cloak (14665, -0.77 DPS) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (538.6 DPS) | yes | Knight-Captain's Lamellar Chestplate (227151, +0.00 DPS) [vendor]; Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Breastplate of Salvation (250601, -20.37 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (538.6 DPS) | yes | Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Prefect's Wristguards (250584, -22.02 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (538.6 DPS) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Mender's Leather Gauntlets (252551, +0.00 DPS) [crafted]; Prefect's Gloves (250554, -20.00 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (538.6 DPS) | yes | Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -20.79 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | 107.5 healing_power points (20.52 DPS) | yes | Knight-Captain's Lamellar Legguards (227150, -2.33 DPS) [vendor]; Marshal's Lamellar Legguards (231639, -4.81 DPS) [vendor]; Red Dragonscale Leggings (252603, -11.10 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 55.6 healing_power points (10.62 DPS) | yes | Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS) [vendor]; Marshal's Lamellar Greaves (231646, -1.15 DPS) [vendor]; Mender's Leather Boots (252472, -2.22 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 23.6 healing_power points (4.50 DPS) | yes | Frigid Ring (18679, -1.20 DPS) [world]; Eye of Adaegus (5266, -1.21 DPS) [world_drop]; Maiden's Circle (13001, -1.54 DPS) [world_drop] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 20.7 healing_power points (3.96 DPS) | yes | Eye of Adaegus (5266, -0.66 DPS) [world_drop]; Maiden's Circle (13001, -0.99 DPS) [world_drop]; Frigid Ring (18679, -1.26 DPS, sim-verified) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (538.6 DPS) | yes | Ankh of Life (1713, -2.85 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.07 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.29 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (538.6 DPS) | yes | Evonice's Landin' Pilla (18951, -2.34 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.56 DPS) [quest]; Ankh of Life (1713, -4.32 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (538.6 DPS) | yes | Hand of Righteousness (7721, -3.74 DPS) [dungeon]; Death Speaker Scepter (2816, -9.58 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -34.68 DPS, sim-verified) [world_drop] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | 36.8 healing_power points (7.02 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Orb of Lorica (11262, -37.63 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Glowing Crystal Ring; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Evergreen Shield

No-known-source sample (15 of 1762, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (undead, 50320100000000000-00000000000000000-00000000000000000)

Set DPS (verified): 27.6. Weights run: 5.4s. Verify run: 18.2s. 222 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.002, intellect=1.018 ± 0.002, spirit=0.462 ± 0.002, mp5=1.888 ± 0.011, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.583 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.84 DPS) | yes | Wisdom's Leather Hood (252507, -0.35 DPS) [crafted]; Pristine Circlet (253949, -0.36 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.12 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.4 healing_power points (0.12 DPS) | yes | Scholarly Pendant (277203, -0.26 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.2 healing_power points (0.80 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.44 DPS) [crafted]; Forest Leather Mantle (4709, -0.64 DPS) [world_drop] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 4.1 healing_power points (0.36 DPS) | yes | Seer's Cape (6378, -0.10 DPS) [dungeon]; Forest Cloak (4710, -0.18 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.22 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 21.5 healing_power points (1.88 DPS) | yes | Filigreed Pristine Gown (253901, -0.17 DPS) [crafted]; Acolyte's Chain Shirt (250491, -0.48 DPS, sim-verified) [crafted]; Corsair's Overshirt (5202, -0.59 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 10.8 healing_power points (0.94 DPS) | yes | Owl Bracers (4796, -0.50 DPS) [vendor]; Featherbead Bracers (15452, -0.50 DPS) [quest]; Garrison Cuffs (270003, -1.49 DPS, sim-verified) [quest] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 15.1 healing_power points (1.32 DPS) | yes | Pristine Gloves (253913, -0.09 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.22 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.49 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 15.1 healing_power points (1.32 DPS) | yes | Acolyte's Chain Belt (250516, -0.24 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.26 DPS) [crafted]; Novice Ardent's Sash (253887, -0.40 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 28.1 healing_power points (2.46 DPS) | yes | Wisdom's Leather Pants (252503, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Darkweave Breeches (12987, -1.59 DPS) [world_drop] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 16.1 healing_power points (1.41 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.26 DPS) [crafted]; Black Whelp Slippers (252424, -0.26 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.1 healing_power points (0.53 DPS) | yes | Volcanic Rock Ring (12053, -0.27 DPS) [world_drop]; Band of Purification (12996, -0.29 DPS) [world_drop]; Deep Fathom Ring (6463, -0.33 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 5.1 healing_power points (0.44 DPS) | yes | Band of Purification (12996, -0.20 DPS) [world_drop]; Deep Fathom Ring (6463, -0.24 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.34 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 9.2 healing_power points (0.80 DPS) | yes | Impaling Harpoon (5200, -0.51 DPS, sim-verified) [dungeon]; Samophlange Screwdriver (11854, -0.56 DPS) [quest]; Rakzur Club (12983, -0.60 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Sword of the Fallen

No-known-source sample (15 of 222, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (undead, 50320215012000000-00000000000000000-00000000000000000)

Set DPS (verified): 44.9. Weights run: 6.0s. Verify run: 22.7s. 443 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.016, intellect=1.305 ± 0.009, spirit=0.626 ± 0.005, mp5=2.511 ± 0.010, crit=0.116 ± 0.007 per rating point (14 rating = 1%, 1.628 per %), spell_haste=not significant (0.067 ± 0.018)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.8 healing_power points (3.00 DPS) | yes | Acolyte's Chain Helm (250501, -0.96 DPS) [crafted]; Filigreed Pristine Circlet (253975, -1.21 DPS) [crafted]; Enduring Cap (3020, -1.33 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 7.8 healing_power points (0.64 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Crystal Starfire Medallion (5003, -0.06 DPS) [world_drop]; Pendant of Myzrael (4614, -0.33 DPS) [dungeon] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 14.6 healing_power points (1.20 DPS) | yes | Nightsky Mantle (4718, -0.09 DPS) [world_drop]; Talbar Mantle (10657, -0.15 DPS) [quest]; Death Speaker Mantle (6685, -0.60 DPS, sim-verified) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 12.0 healing_power points (0.98 DPS) | yes | Soft Willow Cape (16661, -0.45 DPS) [quest]; Silver-thread Cloak (4713, -0.45 DPS) [world_drop]; Resilient Cape (14400, -0.82 DPS, sim-verified) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 34.4 healing_power points (2.81 DPS) | yes | Pristine Gown (253961, -0.22 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -0.26 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.39 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 14.1 healing_power points (1.15 DPS) | yes | Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Technician's Bracers (270042, -0.51 DPS) [quest]; Glowing Magical Bracelets (13106, -0.59 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 18.9 healing_power points (1.54 DPS) | yes | Acolyte's Gloves (250511, -0.29 DPS, sim-verified) [crafted]; Wisdom's Leather Gloves (252499, -0.30 DPS) [crafted]; Naga Battle Gloves (888, -0.32 DPS) [dungeon] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Prefect's Belt (250559, -0.48 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.17 DPS) [crafted]; Silver-lined Belt (13011, -1.33 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Leggings (253987, -0.02 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.34 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.50 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 31.6 healing_power points (2.58 DPS) | yes | Wisdom's Leather Boots (252444, -1.15 DPS) [crafted]; Glowing Copper Boots (250482, -1.42 DPS) [crafted]; Acolyte's Boots (250506, -1.95 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 healing_power points (0.85 DPS) | yes | Solid Gold Band (279845, -0.02 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; The Queen's Jewel (13094, -0.23 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 10.4 healing_power points (0.85 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; The Queen's Jewel (13094, -0.23 DPS) [world_drop]; Solid Gold Band (279845, -0.29 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -2.19 DPS, sim-verified) [dungeon]; Pysan's Old Greatsword (1975, -4.00 DPS) [dungeon]; Wolfsbane (267369, -4.49 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.8 healing_power points (1.37 DPS) | yes | Witch's Finger (16887, -0.52 DPS) [quest]; Heart of Agamaggan (6694, -0.55 DPS) [dungeon]; Orb of Mistmantle (13031, -0.81 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 443, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets

### Band 40 (undead, 50320215012101251-00000000000000000-00000000000000000)

Set DPS (verified): 74.3. Weights run: 9.1s. Verify run: 35.1s. 629 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.026, intellect=1.437 ± 0.014, spirit=0.872 ± 0.008, mp5=3.076 ± 0.013, crit=0.214 ± 0.012 per rating point (14 rating = 1%, 2.994 per %), spell_haste=not significant (0.037 ± 0.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 46.3 healing_power points (3.62 DPS) | yes | Whitemane's Chapeau (7720, -0.66 DPS) [dungeon]; Holy Shroud (2721, -1.01 DPS, sim-verified) [world_drop]; Earthen Silk Hood (254015, -1.22 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.5 healing_power points (1.91 DPS) | yes | Necklace of Calisea (1714, -0.74 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -1.19 DPS) [world_drop]; Amberglow Talisman (10824, -1.23 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 33.7 healing_power points (2.63 DPS) | yes | Sheepshear Mantle (13115, -0.62 DPS, sim-verified) [world_drop]; Earthen Silk Shoulders (254033, -0.68 DPS) [crafted]; Mistscape Mantle (4734, -1.06 DPS) [dungeon] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 15.0 healing_power points (1.17 DPS) | yes | Glowing Thresher Cape (6901, -0.19 DPS, sim-verified) [dungeon]; Windchaser Cloak (14430, -0.22 DPS) [world_drop]; Cloak of Blight (6832, -0.28 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 65.0 healing_power points (5.08 DPS) | yes | Death Speaker Robes (6682, -2.28 DPS) [dungeon]; Pristine Gown (253961, -2.46 DPS) [crafted]; Deathchill Armor (10764, -2.96 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.4 healing_power points (1.28 DPS) | yes | Earthen Silk Cuffs (254019, -0.16 DPS) [crafted]; Windchaser Cuffs (14429, -0.27 DPS) [world_drop]; Enchanted Kodo Bracers (13119, -0.67 DPS, sim-verified) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 33.9 healing_power points (2.65 DPS) | yes | Prefect's Gauntlet (250569, -0.22 DPS) [crafted]; Gilded Handwraps (254021, -0.28 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.44 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 38.7 healing_power points (3.03 DPS) | yes | Prefect's Belt (250559, -0.52 DPS) [crafted]; Mender's Leather Belt (252523, -0.75 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -1.34 DPS) [dungeon] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 37.2 healing_power points (2.91 DPS) | yes | Warchief Kilt (7760, +0.00 DPS, sim-verified) [dungeon]; Acolyte's Silvered Chain Leggings (250526, -0.05 DPS) [crafted]; Pristine Leggings (253987, -0.06 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 40.9 healing_power points (3.20 DPS) | yes | Mender's Leather Shoes (252533, -0.06 DPS) [crafted]; Mender's Mail Boots (252565, -0.06 DPS) [crafted]; Gilded Slippers (254001, -0.58 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 13.3 healing_power points (1.04 DPS) | yes | Ogremind Ring (1993, -0.05 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.09 DPS) [dungeon]; Runemender's Seal (273031, -0.09 DPS) [world] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 12.7 healing_power points (0.99 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.04 DPS) [dungeon]; Runemender's Seal (273031, -0.04 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (74.4 DPS) | yes | Death Speaker Scepter (2816, -2.58 DPS) [dungeon]; Skullbreaker (17039, -4.99 DPS) [quest]; Frost Tiger Blade (3854, -12.57 DPS, sim-verified) [crafted] |
| off_hand | Forcestone Buckler (17508) | Compendium of the Fallen [quest] | 19.0 healing_power points (1.48 DPS) | yes | Beacon of Hope (9393, -0.12 DPS) [dungeon]; Battered Viking Shield (9403, -0.18 DPS) [dungeon]; Orb of Souls (249395, -11.38 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Blackforge Cape; chest: Stormcloth Vest; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Forcestone Buckler

No-known-source sample (15 of 629, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (undead, 50320215012101251-50320000000000000-00000000000000000)

Set DPS (verified): 101.5. Weights run: 9.4s. Verify run: 44.3s. 839 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.023, intellect=2.277 ± 0.024, spirit=1.533 ± 0.013, mp5=5.131 ± 0.019, crit=0.365 ± 0.018 per rating point (14 rating = 1%, 5.113 per %), spell_haste=not significant (0.137 ± 0.041)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 75.3 healing_power points (4.98 DPS) | yes | Winged Helm (13112, -0.18 DPS, sim-verified) [world_drop]; Papal Fez (9431, -0.58 DPS) [dungeon]; Stormcloth Headband (10032, -0.87 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 38.0 healing_power points (2.51 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.75 DPS) [dungeon]; Choker of the High Shaman (4112, -0.95 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 57.8 healing_power points (3.82 DPS) | yes | Mender's Leather Shoulders (252538, +0.00 DPS, sim-verified) [crafted]; Mender's Mail Shoulders (252569, -0.17 DPS) [crafted]; Nethergeld Shoulders (254049, -0.29 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 31.2 healing_power points (2.06 DPS) | yes | Highborne Cloak (14450, -0.33 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -0.45 DPS) [dungeon]; Spritecaster Cape (11623, -0.55 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Earth Warder's Vest (21311, -0.65 DPS) [quest]; Feathered Breastplate (8349, -0.71 DPS) [crafted]; Ghostweave Vest (14141, -2.18 DPS, sim-verified) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 44.1 healing_power points (2.92 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.04 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 65.3 healing_power points (4.32 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.08 DPS) [crafted]; Prefect's Gloves (250554, -0.35 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 57.3 healing_power points (3.79 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.14 DPS) [dungeon]; Prefect's Waistguard (250574, -0.30 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 74.6 healing_power points (4.93 DPS) | yes | Senior Designer's Pantaloons (11841, -0.40 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.44 DPS) [dungeon]; Dalewind Trousers (13008, -1.04 DPS) [world_drop] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 53.1 healing_power points (3.51 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.10 DPS) [crafted]; Mender's Leather Shoes (252533, -0.16 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 32.1 healing_power points (2.12 DPS) | yes | Coldwater Ring (4550, -0.76 DPS) [quest]; Ogremind Ring (1993, -0.76 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 22.1 healing_power points (1.46 DPS) | yes | Ogremind Ring (1993, -0.10 DPS) [world_drop]; Voodoo Band (1996, -0.10 DPS) [world]; Coldwater Ring (4550, -0.18 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -0.18 DPS) [quest]; Alchemist's Stone (13503, -0.78 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.20 DPS) [quest]; Alchemist's Stone (13503, -0.81 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, -2.53 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.06 DPS) [dungeon] |
| off_hand | Protector Buckler (14790) | World drop [world_drop] | sim-verified (+10.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mountainside Buckler (13082, -0.04 DPS) [world_drop]; Ravager's Shield (14777, -0.35 DPS) [world_drop]; Twisting Essence Jar (249456, -10.90 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Mender's Leather Boots; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life; off_hand: Protector Buckler

No-known-source sample (15 of 839, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 161.2. Weights run: 9.1s. Verify run: 118.4s. 1781 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.198, intellect=5.732 ± 0.063, spirit=2.168 ± 0.024, mp5=8.318 ± 0.046, crit=0.613 ± 0.038 per rating point (14 rating = 1%, 8.586 per %), spell_haste=not significant (-0.021 ± 0.085)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (161.2 DPS) | yes | Crown of the Penitent (13216, +0.00 DPS) [quest]; Black Dragonscale Helm (252605, +0.00 DPS) [crafted]; Devout Crown (16693, -0.85 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 88.9 healing_power points (4.60 DPS) | yes | Glowing Eye of Mordresh (10769, -1.55 DPS) [dungeon]; Archlight Talisman (15856, -1.63 DPS) [quest]; Necklace of Calisea (1714, -1.74 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 150.9 healing_power points (7.80 DPS) | yes | Ironfeather Shoulders (15067, -0.98 DPS) [crafted]; Magister's Mantle (16689, -1.28 DPS) [dungeon]; Blue Dragonscale Shoulders (15049, -1.58 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 71.7 healing_power points (3.71 DPS) | yes | Highborne Cloak (14450, -0.59 DPS) [world_drop]; Keeper's Cloak (14665, -0.74 DPS) [world_drop]; Blackforge Cape (6424, -0.89 DPS) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Breastplate of Salvation (250601, +0.00 DPS) [crafted]; Mooncloth Vest (14138, -0.43 DPS) [crafted]; Magister's Robes (16688, -0.63 DPS) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; Indomitable Armguards (14682, +0.00 DPS) [world_drop]; Magister's Bindings (16683, +0.00 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magister's Gloves (16684, +0.00 DPS) [dungeon]; Gilded Gloves (254095, -0.29 DPS) [crafted]; Mender's Leather Gauntlets (252551, -0.37 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 145.2 healing_power points (7.51 DPS) | yes | Elunarian Belt (14465, -0.46 DPS) [world_drop]; Wisdom of the Timbermaw (19047, -0.74 DPS) [crafted]; Dawnspire Cord (12466, -1.25 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Dragonscale Leggings (252603, +0.00 DPS) [crafted]; Devout Skirt (16694, -1.68 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -2.28 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Mender's Leather Boots (252472, -0.73 DPS) [crafted]; Mender's Mail Sabatons (252579, -0.73 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 81.2 healing_power points (4.20 DPS) | yes | Chemist's Ring (15702, -0.78 DPS) [quest]; Eye of Adaegus (5266, -1.07 DPS) [world_drop]; Frigid Ring (18679, -1.23 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 68.8 healing_power points (3.56 DPS) | yes | Chemist's Ring (15702, -0.14 DPS) [quest]; Eye of Adaegus (5266, -0.43 DPS) [world_drop]; Frigid Ring (18679, -0.59 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -1.24 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.46 DPS) [quest]; Alchemist's Stone (13503, -2.13 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.48 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Destroyer (234546, -3.26 DPS) [pvp]; Hand of Righteousness (7721, -3.87 DPS) [dungeon]; High Warlord's Battle Mace (23464, -4.68 DPS) [vendor] |
| off_hand | Elunarian Sphere (15968) | World drop [world_drop] | 80.3 healing_power points (4.15 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, -0.15 DPS) [vendor]; Arcane Star (15931, -0.29 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Glowing Crystal Ring; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Elunarian Sphere

No-known-source sample (15 of 1781, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (undead, 50320215012101251-50322000000000000-05030000000000000)

Set DPS (verified): 531.1. Weights run: 5.3s. Verify run: 39.2s. 1781 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.561, intellect=1.727 ± 0.079, spirit=0.574 ± 0.077, mp5=3.637 ± 0.105, crit=0.984 ± 0.089 per rating point (14 rating = 1%, 13.772 per %), spell_haste=not significant (2.709 ± 1.191)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 112.1 healing_power points (21.40 DPS) | yes | Black Dragonscale Helm (252605, -6.00 DPS, sim-verified) [crafted]; Crown of the Penitent (13216, -6.25 DPS) [quest]; Devout Crown (16693, -8.03 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 26.8 healing_power points (5.12 DPS) | yes | Horizon Choker (13085, -0.91 DPS, sim-verified) [world_drop]; Archlight Talisman (15856, -1.82 DPS) [quest]; Necklace of Calisea (1714, -2.04 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 52.4 healing_power points (10.01 DPS) | yes | Prefect's Pauldrons (250579, -1.18 DPS) [crafted]; Mender's Leather Shoulders (252538, -1.69 DPS) [crafted]; Dawnbringer Shoulders (12625, -22.98 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 21.3 healing_power points (4.06 DPS) | yes | Keeper's Cloak (14665, -0.77 DPS) [world_drop]; Highborne Cloak (14450, -0.85 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -0.99 DPS) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (531.1 DPS) | yes | Mooncloth Vest (14138, +0.00 DPS) [crafted]; Mooncloth Robe (18486, -0.17 DPS) [crafted]; Breastplate of Salvation (250601, -20.29 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (531.1 DPS) | yes | Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted]; Prefect's Wristguards (250584, -21.46 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (531.1 DPS) | yes | Mender's Leather Gauntlets (252551, +0.00 DPS) [crafted]; Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Prefect's Gloves (250554, -19.69 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (531.1 DPS) | yes | Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -20.73 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | 107.5 healing_power points (20.52 DPS) | yes | Devout Skirt (16694, -8.48 DPS) [dungeon]; Elderwild Pants (279254, -10.26 DPS) [crafted]; Red Dragonscale Leggings (252603, -10.83 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 55.6 healing_power points (10.62 DPS) | yes | Mender's Mail Sabatons (252579, -1.65 DPS) [crafted]; Gilded Sandals (254107, -1.74 DPS) [crafted]; Mender's Leather Boots (252472, -2.11 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 23.6 healing_power points (4.50 DPS) | yes | Chemist's Ring (15702, -0.77 DPS) [quest]; Frigid Ring (18679, -1.20 DPS) [world]; Eye of Adaegus (5266, -1.21 DPS) [world_drop] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 20.7 healing_power points (3.96 DPS) | yes | Chemist's Ring (15702, -0.22 DPS) [quest]; Frigid Ring (18679, -0.66 DPS) [world]; Eye of Adaegus (5266, -0.66 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (531.1 DPS) | yes | Ankh of Life (1713, -2.85 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.07 DPS) [quest]; Alchemist's Stone (13503, -3.73 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (531.1 DPS) | yes | Ankh of Life (1713, -2.12 DPS) [world_drop]; Alchemist's Stone (13503, -3.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, -4.26 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (531.1 DPS) | yes | Hand of Righteousness (7721, -3.74 DPS) [dungeon]; Death Speaker Scepter (2816, -9.58 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -34.48 DPS, sim-verified) [world_drop] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | 36.8 healing_power points (7.02 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Elunarian Sphere (15968, -39.30 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Martyr's Legplates; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Glowing Crystal Ring; trinket1: Shard of the Splithooves; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Evergreen Shield

No-known-source sample (15 of 1781, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

