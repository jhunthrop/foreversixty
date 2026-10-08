# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 26.8. Weights run: 4.1s. Verify run: 13.2s. 239 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=0.900 ± 0.002, spirit=0.222 ± 0.001, mp5=1.695 ± 0.023, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.515 per %), spell_haste=0.104 ± 0.012

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.74 DPS) | yes | Wisdom's Leather Hood (252507, -0.33 DPS) [crafted]; Pristine Circlet (253949, -0.39 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.14 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 0.9 healing_power points (0.07 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.67 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.37 DPS) [crafted]; Forest Leather Mantle (4709, -0.60 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 9.4 healing_power points (0.78 DPS) | yes | Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.60 DPS) [dungeon]; Sanguine Cape (14376, -0.70 DPS, sim-verified) [world_drop] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 20.5 healing_power points (1.70 DPS) | yes | Wisdom's Leather Armor (252493, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.62 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 9.6 healing_power points (0.79 DPS) | yes | Bright Bracers (3647, -0.50 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.50 DPS) [vendor]; Owl Bracers (4796, -0.95 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.5 healing_power points (1.20 DPS) | yes | Wisdom's Leather Gloves (252499, -0.07 DPS) [crafted]; Pristine Gloves (253913, -0.14 DPS, sim-verified) [crafted]; Tomb Robber's Gloves (280096, -0.75 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.21 DPS) | yes | Wisdom's Leather Belt (252433, -0.25 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.26 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.43 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 27.3 healing_power points (2.26 DPS) | yes | Wisdom's Leather Pants (252503, -0.05 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.25 DPS) [crafted]; Dreamer's Leggings (270016, -1.44 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.5 healing_power points (1.29 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.24 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.45 DPS) | yes | Black Pearl Ring (6332, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop]; Minor Channeling Ring (1449, -0.30 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 3.4 healing_power points (0.28 DPS) | yes | Volcanic Rock Ring (12053, -0.06 DPS) [world_drop]; Black Pearl Ring (6332, -0.10 DPS, sim-verified) [world]; Minor Channeling Ring (1449, -0.13 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 8.1 healing_power points (0.67 DPS) | yes | Monastic Hammer (270005, -0.52 DPS) [quest]; Impaling Harpoon (5200, -0.57 DPS, sim-verified) [dungeon]; Trogg Slicer (6186, -0.58 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Lorekeeper's Ring; main_hand: Verigan's Fist

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 51.7. Weights run: 4.0s. Verify run: 13.4s. 404 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.980 ± 0.003, spirit=0.488 ± 0.002, mp5=2.096 ± 0.007, crit=0.092 ± 0.004 per rating point (14 rating = 1%, 1.293 per %), spell_haste=0.037 ± 0.009

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Chain Helm (250501) | Blacksmithing [crafted] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Circlet (253975, -0.28 DPS) [crafted]; Wisdom's Leather Helm (252515, -0.37 DPS) [crafted]; Holy Shroud (2721, -3.70 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 5.9 healing_power points (0.54 DPS) | yes | Kaleidoscope Chain (13084, -0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.09 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 13.2 healing_power points (1.22 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.27 DPS) [world_drop]; Mantle of Honor (3560, -0.27 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 15.0 healing_power points (1.38 DPS) | yes | Glowing Thresher Cape (6901, -0.37 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.52 DPS) [vendor]; Prelacy Cape (7004, -2.01 DPS, sim-verified) [quest] |
| chest | Acolyte's Silvered Chain Shirt (250521) | Blacksmithing [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gown (253961, -0.00 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.19 DPS) [crafted]; Death Speaker Robes (6682, -2.64 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.2 healing_power points (1.04 DPS) | yes | Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Spidertank Oilrag (9448, -0.45 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.64 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.9 healing_power points (1.66 DPS) | yes | Acolyte's Gloves (250511, -0.28 DPS) [crafted]; Silvered Gauntlets (270025, -0.36 DPS) [quest]; Naga Battle Gloves (888, -2.82 DPS, sim-verified) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 28.9 healing_power points (2.67 DPS) | yes | Mender's Leather Belt (252523, -0.19 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.29 DPS) [crafted]; Acolyte's Chain Belt (250516, -1.47 DPS) [crafted] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 32.8 healing_power points (3.04 DPS) | yes | Wisdom's Leather Leggings (252519, -0.14 DPS) [crafted]; Pristine Leggings (253987, -0.33 DPS, sim-verified) [crafted]; Acolyte's Chain Leggings (250496, -0.46 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.8 healing_power points (2.67 DPS) | yes | Acolyte's Boots (250506, -1.19 DPS) [crafted]; Wisdom's Leather Boots (252444, -1.19 DPS) [crafted]; Nimbus Boots (6998, -4.23 DPS, sim-verified) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.67 DPS) | yes | Darkspear Signet (272071, -0.89 DPS) [vendor]; Black Widow Band (6199, -1.03 DPS) [world]; Lavishly Jeweled Ring (1156, -1.12 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 healing_power points (0.95 DPS) | yes | Black Widow Band (6199, -0.32 DPS) [world]; Lavishly Jeweled Ring (1156, -0.41 DPS) [dungeon]; Darkspear Signet (272071, -0.52 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -2.09 DPS, sim-verified) [dungeon]; Haunting Blade (6641, -3.24 DPS) [dungeon]; Verigan's Fist (6953, -5.21 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 15.9 healing_power points (1.47 DPS) | yes | Eye of Paleth (2943, -0.18 DPS, sim-verified) [quest]; Alliance Outrunner Healing Rod (285348, -0.36 DPS) [world]; Orb of Mistmantle (13031, -0.46 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Acolyte's Chain Helm; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Caretaker's Cape; chest: Acolyte's Silvered Chain Shirt; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 404, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 76.5. Weights run: 7.1s. Verify run: 26.2s. 562 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.159 ± 0.006, spirit=0.735 ± 0.005, mp5=2.691 ± 0.008, crit=0.146 ± 0.006 per rating point (14 rating = 1%, 2.039 per %), spell_haste=not significant (0.046 ± 0.016)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 48.3 healing_power points (4.32 DPS) | yes | Earthen Silk Hood (254015, -1.69 DPS) [crafted]; Miner's Hat of the Deep (9429, -1.90 DPS) [dungeon]; Holy Shroud (2721, -7.26 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.9 healing_power points (1.24 DPS) | yes | Necklace of Calisea (1714, -0.06 DPS) [dungeon]; Triune Amulet (7722, -0.10 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.20 DPS) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 23.9 healing_power points (2.13 DPS) | yes | Sheepshear Mantle (13115, +0.00 DPS, sim-verified) [world_drop]; Mistscape Mantle (4734, -0.67 DPS) [dungeon]; Batwing Mantle (6697, -0.67 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | 21.7 healing_power points (1.94 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.27 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272077, -0.53 DPS) [vendor]; Prelacy Cape (7004, -0.56 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.9 healing_power points (5.71 DPS) | yes | Pristine Gown (253961, -2.93 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -3.02 DPS) [crafted]; Death Speaker Robes (6682, -7.77 DPS, sim-verified) [dungeon] |
| wrist | Reflective Wristguards (274750) | Rettrick [vendor] | 18.3 healing_power points (1.63 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.41 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.56 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 31.4 healing_power points (2.81 DPS) | yes | Prefect's Gauntlet (250569, -0.21 DPS) [crafted]; Gloves of the Greatfather (17721, -0.49 DPS) [crafted]; Gilded Handwraps (254021, -0.49 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.7 healing_power points (3.19 DPS) | yes | Prefect's Belt (250559, -0.50 DPS) [crafted]; Mender's Leather Belt (252523, -1.07 DPS, sim-verified) [crafted]; Highlander's Lizardhide Girdle (20104, -1.63 DPS) [rep] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 34.3 healing_power points (3.06 DPS) | yes | Pristine Leggings (253987, -0.04 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.39 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.42 DPS, sim-verified) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 38.4 healing_power points (3.43 DPS) | yes | Mender's Mail Boots (252565, -0.10 DPS) [crafted]; Mender's Leather Shoes (252533, -0.61 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -0.66 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.61 DPS) | yes | Snake Hoop (6750, -0.42 DPS) [quest]; Welken Ring (5011, -0.63 DPS) [world_drop]; Voodoo Band (1996, -0.69 DPS) [world] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 16.1 healing_power points (1.44 DPS) | yes | Welken Ring (5011, -0.46 DPS) [world_drop]; Voodoo Band (1996, -0.52 DPS) [world]; Snake Hoop (6750, -2.20 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (76.5 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (76.5 DPS) | yes | Haunting Blade (6641, -3.13 DPS) [dungeon]; Frost Tiger Blade (3854, -3.16 DPS, sim-verified) [crafted]; Mograine's Might (7723, -4.42 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 24.9 healing_power points (2.23 DPS) | yes | Eye of Paleth (2943, -1.07 DPS) [quest]; Ravager's Shield (14777, -1.11 DPS) [world_drop]; Orb of Souls (249395, -1.12 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Reflective Wristguards; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger2: Darkspear Signet; trinket2: Ankh of Life; off_hand: Beacon of Hope

No-known-source sample (15 of 562, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 98.4. Weights run: 7.5s. Verify run: 33.4s. 722 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=1.839 ± 0.013, spirit=1.415 ± 0.011, mp5=4.679 ± 0.015, crit=0.316 ± 0.015 per rating point (14 rating = 1%, 4.425 per %), spell_haste=not significant (0.112 ± 0.046)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulcatcher Halo (10630, -0.42 DPS) [dungeon]; Braincage (12549, -0.52 DPS) [dungeon]; Knight-Lieutenant's Imbued Helmet (220810, -2.59 DPS, sim-verified) [vendor] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 39.1 healing_power points (2.79 DPS) | yes | Gemshard Heart (17707, -0.87 DPS) [dungeon]; Horizon Choker (13085, -0.89 DPS, sim-verified) [world_drop]; Glowing Eye of Mordresh (10769, -1.02 DPS) [dungeon] |
| shoulder | Lead Surveyor's Mantle (11842) | Blackrock Depths: Fineous Darkvire [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Shoulder (252538, -0.11 DPS) [crafted]; Mender's Mail Shoulder (252569, -0.11 DPS) [crafted]; Knight-Lieutenant's Imbued Pauldrons (220808, -1.98 DPS, sim-verified) [vendor] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 32.8 healing_power points (2.34 DPS) | yes | Caretaker's Cape (19531, +0.00 DPS, sim-verified) [rep]; Featherskin Cape (10843, -0.30 DPS) [world]; Imperial Red Cloak (8248, -0.49 DPS) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 73.7 healing_power points (5.25 DPS) | yes | Stormcloth Vest (10020, -0.31 DPS) [crafted]; Robes of Insight (940, -0.46 DPS) [world_drop]; Knight's Imbued Armor (220813, -1.43 DPS, sim-verified) [vendor] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 40.4 healing_power points (2.88 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Gauntlets (252551, -0.15 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.15 DPS) [crafted]; Soulforge Fists (226982, -3.83 DPS, sim-verified) [vendor] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 52.1 healing_power points (3.71 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.26 DPS) [crafted]; Gilded Waistcord (254081, -0.46 DPS) [crafted] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 71.6 healing_power points (5.10 DPS) | yes | Kilt of the Atal'ai Prophet (10807, +0.00 DPS, sim-verified) [dungeon]; Dalewind Trousers (13008, -1.40 DPS) [world_drop]; Windscale Sarong (10842, -1.78 DPS) [world] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 48.2 healing_power points (3.44 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Mail Boots (252565, -0.10 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.7 healing_power points (2.47 DPS) | yes | Eye of Adaegus (5266, -0.47 DPS) [world_drop]; Chivalrous Signet (20505, -0.48 DPS) [quest]; Cyclopean Band (11824, -0.51 DPS) [dungeon] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 32.8 healing_power points (2.33 DPS) | yes | Eye of Adaegus (5266, -0.34 DPS) [world_drop]; Chivalrous Signet (20505, -0.35 DPS) [quest]; Cyclopean Band (11824, -0.37 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -3.72 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.99 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.20 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.20 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.40 DPS) [quest] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Haunting Blade (6641, -2.49 DPS) [dungeon]; Mograine's Might (7723, -2.75 DPS) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | 37.1 healing_power points (2.64 DPS) | yes | Cloud Stone (17737, -0.33 DPS) [dungeon]; Enthralled Sphere (11625, -0.47 DPS, sim-verified) [dungeon]; Twisting Essence Jar (249456, -0.48 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** neck: Darkmoon Necklace; shoulder: Lead Surveyor's Mantle; back: Darkspear Raider's Cloak; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Raider Handwraps; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 722, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 180.3. Weights run: 8.5s. Verify run: 92.5s. 1674 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.126, intellect=4.925 ± 0.058, spirit=1.708 ± 0.014, mp5=7.069 ± 0.046, crit=0.442 ± 0.030 per rating point (14 rating = 1%, 6.193 per %), spell_haste=not significant (0.104 ± 0.054)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 216.3 healing_power points (13.08 DPS) | yes | Living Crown (252561, -1.86 DPS) [crafted]; Crown of the Penitent (13216, -2.34 DPS) [quest]; Black Dragonscale Helm (252605, -2.68 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Jeweled Amulet of Cainwyn (1443, -0.26 DPS) [world_drop]; Tooth of Gnarr (13141, -1.21 DPS) [dungeon]; Lady Maye's Pendant (14558, -6.67 DPS, sim-verified) [world_drop] |
| shoulder | Darkspear Shoulderpads (272103) (or Darkspear Shoulders (272104), Darkspear Shoulderguards (272958)) | Creeg Bothunk [vendor] | 133.0 healing_power points (8.04 DPS) | yes | Darkspear Shoulders (272104, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Devout Mantle (16695, -0.19 DPS) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 91.2 healing_power points (5.52 DPS) | yes | Shroud of the Exile (15421, -0.33 DPS) [quest]; Faded Hakkari Cloak (20218, -0.57 DPS) [quest]; Darkspear Raider's Cloak (272063, -5.17 DPS, sim-verified) [vendor] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Chestplate (227151, -0.02 DPS) [vendor]; Devout Robe (16690, -0.09 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -14.68 DPS, sim-verified) [world] |
| wrist | Gallant's Wristguards (18459) | Dire Maul: Guard Fengus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Hope (22667, -0.32 DPS) [quest]; Bracers of Mending (23129, -0.81 DPS) [dungeon]; Bracers of Undead Slaying (23090, -8.66 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -2.39 DPS) [dungeon]; Marshal's Lamellar Gloves (231643, -2.73 DPS) [pvp]; Razor Gauntlets (18326, -10.96 DPS, sim-verified) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 133.2 healing_power points (8.06 DPS) | yes | Whipvine Cord (18327, -0.94 DPS) [dungeon]; Elunarian Belt (14465, -1.07 DPS) [world_drop]; Devout Belt (16696, -3.80 DPS, sim-verified) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Dragonscale Leggings (252603, -0.32 DPS) [crafted]; Martyr's Legplates (250600, -0.62 DPS) [crafted]; Cloudkeeper Legplates (14554, -15.75 DPS, sim-verified) [world_drop] |
| feet | Soulforge Treads (226983) | Mokvar [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Lamellar Greaves (227153, -0.54 DPS) [vendor]; Marshal's Lamellar Boots (16472, -0.97 DPS) [vendor]; Incandescent Mooncloth Boots (227862, -2.23 DPS, sim-verified) [vendor] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Demonic Guile (18314, -0.48 DPS) [dungeon]; Emerald Flame Ring (18395, -0.71 DPS) [dungeon]; Naglering (11669, -9.62 DPS, sim-verified) [dungeon] |
| finger2 | Seal of Rivendare (13345) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Demonic Guile (18314, -0.24 DPS) [dungeon]; Emerald Flame Ring (18395, -0.48 DPS) [dungeon]; Naglering (11669, -5.51 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -3.85 DPS) [quest]; Serenity Field (272439, -4.25 DPS, sim-verified) [vendor]; Briarwood Reed (12930, -4.66 DPS) [dungeon] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -2.14 DPS) [quest]; Serenity Field (272439, -2.51 DPS, sim-verified) [vendor]; Briarwood Reed (12930, -2.95 DPS) [dungeon] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Demolisher (234568, -3.67 DPS) [pvp]; Hammer of Divine Might (22333, -3.90 DPS) [dungeon]; Hand of Edward the Odd (2243, -7.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Darkspear Shoulderpads; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Gallant's Wristguards; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Padre's Trousers; feet: Soulforge Treads; finger1: Band of Piety; finger2: Seal of Rivendare; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60, raid preset (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 509.0. Weights run: 5.1s. Verify run: 39.7s. 1674 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.361, intellect=1.096 ± 0.055, spirit=0.446 ± 0.053, mp5=2.718 ± 0.069, crit=0.601 ± 0.062 per rating point (14 rating = 1%, 8.416 per %), spell_haste=not significant (0.056 ± 0.782)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 99.0 healing_power points (25.97 DPS) | yes | Lieutenant Commander's Lamellar Helmet (227149, -4.63 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, -7.71 DPS) [vendor]; Soulforge Crown (226981, -17.18 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.3 healing_power points (13.73 DPS) | yes | Drake Tooth Necklace (21531, -3.66 DPS, sim-verified) [quest]; Animated Chain Necklace (18723, -4.37 DPS) [dungeon]; Amulet of the Redeemed (22327, -6.27 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 79.8 healing_power points (20.93 DPS) | yes | Shimmering Dawnbringer Shoulders (227859, +0.00 DPS, sim-verified) [vendor]; Lieutenant Commander's Lamellar Pauldrons (227148, -5.05 DPS) [vendor]; Dawnbringer Shoulders (12625, -8.21 DPS) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.0 healing_power points (13.90 DPS) | yes | Drape of Recovery (272413, -3.71 DPS, sim-verified) [vendor]; Cloak of the Cosmos (18389, -3.91 DPS) [dungeon]; Caretaker's Cape (19530, -6.14 DPS) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | sim-verified (509.0 DPS) | yes | Knight-Captain's Lamellar Chestplate (227151, -0.04 DPS) [vendor]; Breastplate of Salvation (250601, -2.19 DPS) [crafted]; Breastplate of Undead Slaying (23087, -22.25 DPS, sim-verified) [world] |
| wrist | Gallant's Wristguards (18459) | Dire Maul: Guard Fengus [dungeon] | sim-verified (509.0 DPS) | yes | Loomguard Armbraces (13969, -0.96 DPS) [dungeon]; Bracers of Hope (22667, -2.35 DPS) [quest]; Bracers of Undead Slaying (23090, -14.40 DPS, sim-verified) [world] |
| hands | Harmonious Gauntlets (18527) | Dire Maul: King Gordok [dungeon] | sim-verified (509.0 DPS) | yes | Raider Handwraps (272097, -1.34 DPS) [vendor]; Hands of the Exalted Herald (12554, -1.60 DPS) [dungeon]; Razor Gauntlets (18326, -17.71 DPS, sim-verified) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 57.5 healing_power points (15.07 DPS) | yes | Whipvine Cord (18327, +0.00 DPS, sim-verified) [dungeon]; Eyestalk Cord (18391, -1.58 DPS) [dungeon]; Belt of the Ordained (18702, -1.75 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-verified (509.0 DPS) | yes | Padre's Trousers (18386, -4.23 DPS) [dungeon]; Knight-Captain's Lamellar Legguards (227150, -4.47 DPS) [vendor]; Cloudkeeper Legplates (14554, -24.80 DPS, sim-verified) [world_drop] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 65.6 healing_power points (17.21 DPS) | yes | Knight-Lieutenant's Lamellar Greaves (227153, +0.00 DPS, sim-verified) [vendor]; Soulforge Treads (226983, -4.46 DPS) [vendor]; Mooncloth Boots (15802, -5.36 DPS) [crafted] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (509.0 DPS) | yes | Fordring's Seal (16058, -0.35 DPS) [quest]; Band of Mending (22334, -0.65 DPS) [dungeon]; Naglering (11669, -14.23 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-verified (509.0 DPS) | yes | Fordring's Seal (16058, -0.03 DPS) [quest]; Band of Mending (22334, -0.33 DPS) [dungeon]; Naglering (11669, -12.69 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (509.0 DPS) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -2.85 DPS) [dungeon]; Briarwood Reed (12930, -3.09 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (509.0 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (509.0 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hammer of the Grand Crusader (18717, -3.54 DPS) [dungeon]; Hammer of Divine Might (22333, -6.44 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | sim-verified (509.0 DPS) | yes | Grand Marshal's Tome of Restoration (234590, -2.42 DPS) [pvp]; Tome of Divine Right (22319, -2.95 DPS) [dungeon]; Skullflame Shield (1168, -15.70 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Gallant's Wristguards; hands: Harmonious Gauntlets; waist: Sash of Mercy; legs: Martyr's Legplates; feet: Incandescent Mooncloth Boots; finger1: Band of Piety; finger2: Rosewine Circle; trinket2: Draconic Infused Emblem; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 26.3. Weights run: 4.1s. Verify run: 13.4s. 219 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=0.900 ± 0.002, spirit=0.222 ± 0.001, mp5=1.695 ± 0.023, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.515 per %), spell_haste=0.104 ± 0.012

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.74 DPS) | yes | Wisdom's Leather Hood (252507, -0.33 DPS) [crafted]; Pristine Circlet (253949, -0.38 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.14 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.9 healing_power points (0.07 DPS) | yes | Roadwatcher's Confidence (281265, -0.30 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.67 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.37 DPS) [crafted]; Slime-encrusted Pads (6461, -0.43 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.60 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 9.4 healing_power points (0.78 DPS) | yes | Pearl-clasped Cloak (5542, -0.56 DPS) [crafted]; Seer's Cape (6378, -0.60 DPS) [dungeon]; Sanguine Cape (14376, -0.79 DPS, sim-verified) [world_drop] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 20.5 healing_power points (1.70 DPS) | yes | Wisdom's Leather Armor (252493, -0.13 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.62 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 9.6 healing_power points (0.79 DPS) | yes | Garrison Cuffs (270003, -0.39 DPS) [quest]; Owl Bracers (4796, -0.42 DPS) [vendor]; Tabitha's Cuffs (251486, -0.68 DPS, sim-verified) [quest] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.5 healing_power points (1.20 DPS) | yes | Wisdom's Leather Gloves (252499, -0.07 DPS) [crafted]; Pristine Gloves (253913, -0.23 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.59 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.21 DPS) | yes | Wisdom's Leather Belt (252433, -0.25 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.26 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.43 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 27.3 healing_power points (2.26 DPS) | yes | Wisdom's Leather Pants (252503, -0.22 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.25 DPS) [crafted]; Ghastly Trousers (15449, -1.61 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.5 healing_power points (1.29 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.24 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.45 DPS) | yes | Advisor's Ring (20426, -0.17 DPS) [rep]; Black Pearl Ring (6332, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 4.5 healing_power points (0.37 DPS) | yes | Black Pearl Ring (6332, -0.11 DPS) [world]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop]; Advisor's Ring (20426, -0.33 DPS, sim-verified) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 4.5 healing_power points (0.37 DPS) | yes | Mug of Muddled Memories (277247, -0.28 DPS) [quest]; Rakzur Club (12983, -0.28 DPS) [world_drop]; Samophlange Screwdriver (11854, -0.77 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Impaling Harpoon

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 51.4. Weights run: 4.0s. Verify run: 13.6s. 381 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.980 ± 0.003, spirit=0.488 ± 0.002, mp5=2.096 ± 0.007, crit=0.092 ± 0.004 per rating point (14 rating = 1%, 1.293 per %), spell_haste=0.037 ± 0.009

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Chain Helm (250501) | Blacksmithing [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Circlet (253975, -0.28 DPS) [crafted]; Wisdom's Leather Helm (252515, -0.37 DPS) [crafted]; Holy Shroud (2721, -3.42 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 5.9 healing_power points (0.54 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.09 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Nightsky Mantle (4718, -0.27 DPS) [world_drop]; Ghostly Mantle (3324, -1.74 DPS, sim-verified) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 15.0 healing_power points (1.38 DPS) | yes | Darkspear Raider's Cloak (272078, -0.52 DPS) [vendor]; Cloak of Rot (4462, -0.66 DPS) [world]; Glowing Thresher Cape (6901, -1.54 DPS, sim-verified) [dungeon] |
| chest | Acolyte's Silvered Chain Shirt (250521) | Blacksmithing [crafted] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gown (253961, -0.00 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.19 DPS) [crafted]; Death Speaker Robes (6682, -2.64 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.2 healing_power points (1.04 DPS) | yes | Glowing Magical Bracelets (13106, -0.12 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Spidertank Oilrag (9448, -0.45 DPS) [dungeon] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.9 healing_power points (1.66 DPS) | yes | Acolyte's Gloves (250511, -0.28 DPS) [crafted]; Pristine Gloves (253913, -0.37 DPS) [crafted]; Naga Battle Gloves (888, -2.70 DPS, sim-verified) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 28.9 healing_power points (2.67 DPS) | yes | Mender's Leather Belt (252523, +0.00 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.29 DPS) [crafted]; Acolyte's Chain Belt (250516, -1.47 DPS) [crafted] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 32.8 healing_power points (3.04 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Leggings (252519, -0.14 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.46 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.8 healing_power points (2.67 DPS) | yes | Wisdom's Leather Boots (252444, -1.19 DPS) [crafted]; Glowing Copper Boots (250482, -1.47 DPS) [crafted]; Acolyte's Boots (250506, -1.83 DPS, sim-verified) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.67 DPS) | yes | Darkspear Signet (272071, -0.89 DPS) [vendor]; Black Widow Band (6199, -1.03 DPS) [world]; Lavishly Jeweled Ring (1156, -1.12 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 healing_power points (0.95 DPS) | yes | Darkspear Signet (272071, -0.15 DPS, sim-verified) [vendor]; Black Widow Band (6199, -0.32 DPS) [world]; Lavishly Jeweled Ring (1156, -0.41 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -2.10 DPS, sim-verified) [dungeon]; Haunting Blade (6641, -3.24 DPS) [dungeon]; Wolfsbane (267369, -5.38 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 15.9 healing_power points (1.47 DPS) | yes | Orb of Mistmantle (13031, -0.46 DPS) [world_drop]; Defective Samophlange (274743, -0.79 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -2.19 DPS, sim-verified) [world] |
| ranged | - | - |  |  |  |

**New at 30:** head: Acolyte's Chain Helm; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Battle Healer's Cloak; chest: Acolyte's Silvered Chain Shirt; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 381, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 75.9. Weights run: 7.1s. Verify run: 26.5s. 531 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.007, intellect=1.159 ± 0.006, spirit=0.735 ± 0.005, mp5=2.691 ± 0.008, crit=0.146 ± 0.006 per rating point (14 rating = 1%, 2.039 per %), spell_haste=not significant (0.046 ± 0.016)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 48.3 healing_power points (4.32 DPS) | yes | Holy Shroud (2721, -0.97 DPS) [world_drop]; Earthen Silk Hood (254015, -1.69 DPS) [crafted]; Miner's Hat of the Deep (9429, -1.90 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (75.9 DPS) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.15 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | sim-verified (75.9 DPS) | yes | Earthen Silk Shoulders (254033, +0.00 DPS) [crafted]; Mistscape Mantle (4734, -0.58 DPS) [dungeon]; Batwing Mantle (6697, -0.58 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | sim-verified (75.9 DPS) | yes | Battle Healer's Cloak (19528, +0.00 DPS) [rep]; Darkspear Raider's Cloak (272077, -0.33 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.9 healing_power points (5.71 DPS) | yes | Death Speaker Robes (6682, -2.78 DPS) [dungeon]; Pristine Gown (253961, -2.93 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -3.02 DPS) [crafted] |
| wrist | Reflective Wristguards (274750) | Rettrick [vendor] | 18.3 healing_power points (1.63 DPS) | yes | Mindthrust Bracers (1974, -0.39 DPS) [dungeon]; Earthen Silk Cuffs (254019, -0.41 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.56 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 31.4 healing_power points (2.81 DPS) | yes | Gilded Handwraps (254021, -0.12 DPS) [crafted]; Prefect's Gauntlet (250569, -0.21 DPS) [crafted]; Gloves of the Greatfather (17721, -0.49 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.7 healing_power points (3.19 DPS) | yes | Mender's Leather Belt (252523, -0.48 DPS) [crafted]; Prefect's Belt (250559, -0.50 DPS) [crafted]; Highlander's Mail Girdle (20119, -1.63 DPS) [vendor] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 34.3 healing_power points (3.06 DPS) | yes | Wisdom's Leather Leggings (252519, -0.00 DPS) [crafted]; Pristine Leggings (253987, -0.04 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.39 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 38.4 healing_power points (3.43 DPS) | yes | Mender's Leather Shoes (252533, -0.10 DPS) [crafted]; Mender's Mail Boots (252565, -0.10 DPS) [crafted]; Gilded Slippers (254001, -0.66 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.61 DPS) | yes | Snake Hoop (6750, -0.42 DPS) [quest]; Welken Ring (5011, -0.63 DPS) [world_drop]; Voodoo Band (1996, -0.69 DPS) [world] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 16.1 healing_power points (1.44 DPS) | yes | Snake Hoop (6750, -0.26 DPS) [quest]; Welken Ring (5011, -0.46 DPS) [world_drop]; Voodoo Band (1996, -0.52 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Haunting Blade (6641, -3.13 DPS) [dungeon]; Mograine's Might (7723, -4.42 DPS) [dungeon]; Wolfsbane (267369, -5.04 DPS) [quest] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 24.9 healing_power points (2.23 DPS) | yes | Orb of Souls (249395, -0.67 DPS) [crafted]; Prophetic Cane (6803, -0.99 DPS) [quest]; Ravager's Shield (14777, -1.11 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Sheepshear Mantle; back: Mantle of Lady Falther'ess; chest: Stormcloth Vest; wrist: Reflective Wristguards; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger2: Darkspear Signet; trinket2: Ankh of Life; off_hand: Beacon of Hope

No-known-source sample (15 of 531, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 101.7. Weights run: 7.5s. Verify run: 32.0s. 702 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=1.839 ± 0.013, spirit=1.415 ± 0.011, mp5=4.679 ± 0.015, crit=0.316 ± 0.015 per rating point (14 rating = 1%, 4.425 per %), spell_haste=not significant (0.112 ± 0.046)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 66.0 healing_power points (4.70 DPS) | yes | Braincage (12549, -0.52 DPS) [dungeon]; Helm of Exile (11124, -0.53 DPS) [quest]; Soulcatcher Halo (10630, -1.87 DPS, sim-verified) [dungeon] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 39.1 healing_power points (2.79 DPS) | yes | Gemshard Heart (17707, -0.87 DPS) [dungeon]; Horizon Choker (13085, -0.99 DPS, sim-verified) [world_drop]; Glowing Eye of Mordresh (10769, -1.02 DPS) [dungeon] |
| shoulder | Lead Surveyor's Mantle (11842) | Blackrock Depths: Fineous Darkvire [dungeon] | 52.1 healing_power points (3.71 DPS) | yes | Mender's Leather Shoulder (252538, +0.00 DPS, sim-verified) [crafted]; Mender's Mail Shoulder (252569, -0.11 DPS) [crafted]; Living Shoulders (15061, -0.19 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 32.8 healing_power points (2.34 DPS) | yes | Battle Healer's Cloak (19527, +0.00 DPS, sim-verified) [rep]; Featherskin Cape (10843, -0.30 DPS) [world]; Imperial Red Cloak (8248, -0.49 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Robes of Insight (940, -0.15 DPS) [world_drop]; Ghostweave Vest (14141, -0.26 DPS) [crafted]; Embrace of the Wind Serpent (12462, -3.39 DPS, sim-verified) [world] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 40.4 healing_power points (2.88 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Gauntlets (252551, -0.15 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.15 DPS) [crafted]; Soulforge Fists (226982, -3.79 DPS, sim-verified) [vendor] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 52.1 healing_power points (3.71 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.26 DPS) [crafted]; Gilded Waistcord (254081, -0.46 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 58.6 healing_power points (4.17 DPS) | yes | Dalewind Trousers (13008, -0.67 DPS, sim-verified) [world_drop]; Windscale Sarong (10842, -0.85 DPS) [world]; Jinxed Hoodoo Kilt (9474, -0.95 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 48.2 healing_power points (3.44 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Leather Shoes (252533, -0.10 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.7 healing_power points (2.47 DPS) | yes | Eye of Adaegus (5266, -0.47 DPS) [world_drop]; Cyclopean Band (11824, -0.51 DPS) [dungeon]; Snake Hoop (6750, -0.85 DPS) [quest] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 32.8 healing_power points (2.33 DPS) | yes | Eye of Adaegus (5266, -0.34 DPS) [world_drop]; Cyclopean Band (11824, -0.37 DPS) [dungeon]; Snake Hoop (6750, -0.71 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -3.72 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.99 DPS) [quest]; Alchemists' Stone (13503, -4.60 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.23 DPS, sim-verified) [quest]; Alchemists' Stone (13503, -0.81 DPS) [crafted] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Haunting Blade (6641, -2.49 DPS) [dungeon]; Mograine's Might (7723, -2.75 DPS) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | 37.1 healing_power points (2.64 DPS) | yes | Cloud Stone (17737, -0.33 DPS) [dungeon]; Twisting Essence Jar (249456, -0.48 DPS) [crafted]; Enthralled Sphere (11625, -0.57 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** neck: Darkmoon Necklace; shoulder: Lead Surveyor's Mantle; back: Darkspear Raider's Cloak; wrist: Mender's Leather Bracers; hands: Raider Handwraps; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 702, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 179.3. Weights run: 8.5s. Verify run: 88.4s. 1699 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.126, intellect=4.925 ± 0.058, spirit=1.708 ± 0.014, mp5=7.069 ± 0.046, crit=0.442 ± 0.030 per rating point (14 rating = 1%, 6.193 per %), spell_haste=not significant (0.104 ± 0.054)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 216.3 healing_power points (13.08 DPS) | yes | Living Crown (252561, -1.86 DPS) [crafted]; Crown of the Penitent (13216, -2.34 DPS) [quest]; Black Dragonscale Helm (252605, -2.67 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Jeweled Amulet of Cainwyn (1443, -0.26 DPS) [world_drop]; Tooth of Gnarr (13141, -1.21 DPS) [dungeon]; Lady Maye's Pendant (14558, -6.69 DPS, sim-verified) [world_drop] |
| shoulder | Darkspear Shoulderpads (272103) (or Darkspear Shoulders (272104), Darkspear Shoulderguards (272958)) | Creeg Bothunk [vendor] | 133.0 healing_power points (8.04 DPS) | yes | Darkspear Shoulders (272104, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Devout Mantle (16695, -0.19 DPS) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 91.2 healing_power points (5.52 DPS) | yes | Shroud of the Exile (15421, -0.33 DPS) [quest]; Faded Hakkari Cloak (20218, -0.57 DPS) [quest]; Darkspear Raider's Cloak (272063, -5.25 DPS, sim-verified) [vendor] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devout Robe (16690, -0.09 DPS) [dungeon]; Mooncloth Vest (14138, -0.51 DPS) [crafted]; Breastplate of Undead Slaying (23087, -14.33 DPS, sim-verified) [world] |
| wrist | Gallant's Wristguards (18459) | Dire Maul: Guard Fengus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Hope (22667, -0.32 DPS) [quest]; Bracers of Mending (23129, -0.81 DPS) [dungeon]; Bracers of Undead Slaying (23090, -8.53 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -2.39 DPS) [dungeon]; Mooncloth Gloves (18409, -2.73 DPS) [crafted]; Razor Gauntlets (18326, -11.00 DPS, sim-verified) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 133.2 healing_power points (8.06 DPS) | yes | Whipvine Cord (18327, -0.94 DPS) [dungeon]; Elunarian Belt (14465, -1.07 DPS) [world_drop]; Devout Belt (16696, -4.12 DPS, sim-verified) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Dragonscale Leggings (252603, -0.32 DPS) [crafted]; Martyr's Legplates (250600, -0.62 DPS) [crafted]; Cloudkeeper Legplates (14554, -15.84 DPS, sim-verified) [world_drop] |
| feet | Soulforge Treads (226983) | Mokvar [vendor] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Boots (15802, -1.04 DPS) [crafted]; Faith Healer's Boots (22247, -1.62 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -2.25 DPS, sim-verified) [vendor] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Demonic Guile (18314, -0.48 DPS) [dungeon]; Emerald Flame Ring (18395, -0.71 DPS) [dungeon]; Naglering (11669, -9.85 DPS, sim-verified) [dungeon] |
| finger2 | Seal of Rivendare (13345) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Demonic Guile (18314, -0.24 DPS) [dungeon]; Emerald Flame Ring (18395, -0.48 DPS) [dungeon]; Naglering (11669, -5.38 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -3.85 DPS) [quest]; Serenity Field (272439, -3.87 DPS) [vendor]; Briarwood Reed (12930, -4.66 DPS) [dungeon] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -2.14 DPS) [quest]; Serenity Field (272439, -2.85 DPS, sim-verified) [vendor]; Briarwood Reed (12930, -2.95 DPS) [dungeon] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Destroyer (234546, -3.67 DPS) [pvp]; Hammer of Divine Might (22333, -3.90 DPS) [dungeon]; Hand of Edward the Odd (2243, -7.01 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Darkspear Shoulderpads; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Gallant's Wristguards; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Padre's Trousers; feet: Soulforge Treads; finger1: Band of Piety; finger2: Seal of Rivendare; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 506.7. Weights run: 5.1s. Verify run: 38.5s. 1699 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.361, intellect=1.096 ± 0.055, spirit=0.446 ± 0.053, mp5=2.718 ± 0.069, crit=0.601 ± 0.062 per rating point (14 rating = 1%, 8.416 per %), spell_haste=not significant (0.056 ± 0.782)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 99.0 healing_power points (25.97 DPS) | yes | Sanctified Leather Helm (22689, -8.60 DPS) [quest]; Insightful Hood (18490, -9.73 DPS) [dungeon]; Soulforge Crown (226981, -17.25 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.3 healing_power points (13.73 DPS) | yes | Drake Tooth Necklace (21531, -3.39 DPS, sim-verified) [quest]; Animated Chain Necklace (18723, -4.37 DPS) [dungeon]; Amulet of the Redeemed (22327, -6.27 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 79.8 healing_power points (20.93 DPS) | yes | Shimmering Dawnbringer Shoulders (227859, +0.00 DPS, sim-verified) [vendor]; Dawnbringer Shoulders (12625, -8.21 DPS) [crafted]; Darkspear Mantle (272107, -9.01 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.0 healing_power points (13.90 DPS) | yes | Drape of Recovery (272413, -2.67 DPS, sim-verified) [vendor]; Cloak of the Cosmos (18389, -3.91 DPS) [dungeon]; Battle Healer's Cloak (19526, -6.14 DPS) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | sim-verified (506.7 DPS) | yes | Breastplate of Salvation (250601, -2.19 DPS) [crafted]; Stormcloth Vest (10020, -4.41 DPS) [crafted]; Breastplate of Undead Slaying (23087, -22.01 DPS, sim-verified) [world] |
| wrist | Gallant's Wristguards (18459) | Dire Maul: Guard Fengus [dungeon] | sim-verified (506.7 DPS) | yes | Loomguard Armbraces (13969, -0.96 DPS) [dungeon]; Bracers of Hope (22667, -2.35 DPS) [quest]; Bracers of Undead Slaying (23090, -14.28 DPS, sim-verified) [world] |
| hands | Harmonious Gauntlets (18527) | Dire Maul: King Gordok [dungeon] | sim-verified (506.7 DPS) | yes | Raider Handwraps (272097, -1.34 DPS) [vendor]; Hands of the Exalted Herald (12554, -1.60 DPS) [dungeon]; Razor Gauntlets (18326, -17.39 DPS, sim-verified) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 57.5 healing_power points (15.07 DPS) | yes | Whipvine Cord (18327, +0.00 DPS, sim-verified) [dungeon]; Eyestalk Cord (18391, -1.58 DPS) [dungeon]; Belt of the Ordained (18702, -1.75 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-verified (506.7 DPS) | yes | Padre's Trousers (18386, -4.23 DPS) [dungeon]; Soulforge Leggings (226980, -10.44 DPS) [vendor]; Cloudkeeper Legplates (14554, -24.51 DPS, sim-verified) [world_drop] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 65.6 healing_power points (17.21 DPS) | yes | Soulforge Treads (226983, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -5.36 DPS) [crafted]; Faith Healer's Boots (22247, -5.53 DPS) [dungeon] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (506.7 DPS) | yes | Fordring's Seal (16058, -0.35 DPS) [quest]; Band of Mending (22334, -0.65 DPS) [dungeon]; Naglering (11669, -13.86 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-verified (506.7 DPS) | yes | Fordring's Seal (16058, -0.03 DPS) [quest]; Band of Mending (22334, -0.33 DPS) [dungeon]; Naglering (11669, -12.27 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (506.7 DPS) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -2.85 DPS) [dungeon]; Briarwood Reed (12930, -3.09 DPS) [dungeon] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (506.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (506.7 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hammer of the Grand Crusader (18717, -3.54 DPS) [dungeon]; Hammer of Divine Might (22333, -6.44 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | sim-verified (506.7 DPS) | yes | High Warlord's Tome of Mending (234564, -2.42 DPS) [pvp]; Tome of Divine Right (22319, -2.95 DPS) [dungeon]; Skullflame Shield (1168, -15.44 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Gallant's Wristguards; hands: Harmonious Gauntlets; waist: Sash of Mercy; legs: Martyr's Legplates; feet: Incandescent Mooncloth Boots; finger1: Band of Piety; finger2: Rosewine Circle; trinket2: Draconic Infused Emblem; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

