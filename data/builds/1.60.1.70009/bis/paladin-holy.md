# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 26.2. Weights run: 4.7s. Verify run: 2.2s. 239 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.002, intellect=0.897 ± 0.002, spirit=0.227 ± 0.001, mp5=1.697 ± 0.023, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.515 per %), spell_haste=0.067 ± 0.009

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.74 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.33 DPS) [crafted]; Trapper's Leather Hood (252505, -1.14 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 0.9 healing_power points (0.08 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.67 DPS) | yes | Slime-encrusted Pads (6461, -0.19 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.37 DPS) [crafted]; Forest Leather Mantle (4709, -0.59 DPS) [world_drop] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pearl-clasped Cloak (5542, -0.07 DPS) [crafted]; Caretaker's Cape (20428, -0.10 DPS, sim-verified) [rep]; Seer's Cape (6378, -0.11 DPS) [dungeon] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Acolyte's Chain Shirt (250491, -0.03 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.17 DPS) [crafted]; Bloody Apron (6226, -0.59 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 9.6 healing_power points (0.79 DPS) | yes | Owl Bracers (4796, -0.12 DPS, sim-verified) [vendor]; Bright Bracers (3647, -0.49 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.49 DPS) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.5 healing_power points (1.20 DPS) | yes | Pristine Gloves (253913, -0.07 DPS, sim-verified) [crafted]; Wisdom's Leather Gloves (252499, -0.07 DPS) [crafted]; Magefist Gloves (12977, -0.75 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.21 DPS) | yes | Acolyte's Chain Belt (250516, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.25 DPS) [crafted]; Novice Ardent's Sash (253887, -0.43 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 27.3 healing_power points (2.25 DPS) | yes | Wisdom's Leather Pants (252503, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.25 DPS) [crafted]; Dreamer's Leggings (270016, -1.44 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.5 healing_power points (1.28 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.24 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.44 DPS) | yes | Volcanic Rock Ring (12053, -0.22 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.23 DPS, sim-verified) [rep]; Minor Channeling Ring (1449, -0.30 DPS) [quest] |
| finger2 | Black Pearl Ring (6332) | Lady Vespira [world] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Volcanic Rock Ring (12053, -0.04 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.05 DPS, sim-verified) [rep]; Minor Channeling Ring (1449, -0.11 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 8.1 healing_power points (0.67 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Monastic Hammer (270005, -0.52 DPS) [quest]; Trogg Slicer (6186, -0.58 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Black Pearl Ring; main_hand: Verigan's Fist

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 49.1. Weights run: 4.6s. Verify run: 2.2s. 404 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.979 ± 0.003, spirit=0.489 ± 0.003, mp5=2.096 ± 0.007, crit=0.092 ± 0.004 per rating point (14 rating = 1%, 1.285 per %), spell_haste=not significant (0.034 ± 0.008)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Chain Helm (250501) | Blacksmithing [crafted] | sim-verified (4.1 DPS) | yes | Holy Shroud (2721, -0.25 DPS, sim-verified) [world_drop]; Filigreed Pristine Circlet (253975, -0.28 DPS) [crafted]; Wisdom's Leather Helm (252515, -0.37 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 5.9 healing_power points (0.54 DPS) | yes | Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.08 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.09 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 13.2 healing_power points (1.22 DPS) | yes | Death Speaker Mantle (6685, -0.06 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.27 DPS) [world_drop]; Mantle of Honor (3560, -0.27 DPS) [quest] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 15.0 healing_power points (1.38 DPS) | yes | Prelacy Cape (7004, -0.17 DPS, sim-verified) [quest]; Glowing Thresher Cape (6901, -0.37 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.52 DPS) [vendor] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 30.8 healing_power points (2.84 DPS) | yes | Acolyte's Silvered Chain Shirt (250521, +0.00 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.18 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.37 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.2 healing_power points (1.03 DPS) | yes | Glowing Magical Bracelets (13106, -0.05 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Spidertank Oilrag (9448, -0.45 DPS) [dungeon] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.9 healing_power points (1.66 DPS) | yes | Acolyte's Gloves (250511, -0.28 DPS) [crafted]; Silvered Gauntlets (270025, -0.36 DPS) [quest]; Naga Battle Gloves (888, -0.42 DPS, sim-verified) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 28.9 healing_power points (2.67 DPS) | yes | Mender's Leather Belt (252523, +0.00 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.29 DPS) [crafted]; Acolyte's Chain Belt (250516, -1.47 DPS) [crafted] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 32.8 healing_power points (3.03 DPS) | yes | Pristine Leggings (253987, -0.05 DPS, sim-verified) [crafted]; Wisdom's Leather Leggings (252519, -0.14 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.46 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.8 healing_power points (2.66 DPS) | yes | Nimbus Boots (6998, -0.77 DPS, sim-verified) [quest]; Acolyte's Boots (250506, -1.19 DPS) [crafted]; Wisdom's Leather Boots (252444, -1.19 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.66 DPS) | yes | Darkspear Signet (272071, -0.89 DPS) [vendor]; Black Widow Band (6199, -1.03 DPS) [world]; Lavishly Jeweled Ring (1156, -1.12 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 healing_power points (0.95 DPS) | yes | Darkspear Signet (272071, -0.31 DPS, sim-verified) [vendor]; Black Widow Band (6199, -0.32 DPS) [world]; Lavishly Jeweled Ring (1156, -0.41 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -0.23 DPS, sim-verified) [dungeon]; Haunting Blade (6641, -3.23 DPS) [dungeon]; Verigan's Fist (6953, -5.20 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 15.9 healing_power points (1.47 DPS) | yes | Eye of Paleth (2943, -0.08 DPS, sim-verified) [quest]; Alliance Outrunner Healing Rod (285348, -0.36 DPS) [world]; Orb of Mistmantle (13031, -0.46 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Acolyte's Chain Helm; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 404, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 76.8. Weights run: 7.8s. Verify run: 3.5s. 562 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.009, intellect=1.167 ± 0.006, spirit=0.727 ± 0.005, mp5=2.764 ± 0.008, crit=0.140 ± 0.006 per rating point (14 rating = 1%, 1.960 per %), spell_haste=not significant (0.045 ± 0.018)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 48.4 healing_power points (4.29 DPS) | yes | Holy Shroud (2721, -1.46 DPS, sim-verified) [world_drop]; Earthen Silk Hood (254015, -1.69 DPS) [crafted]; Miner's Hat of the Deep (9429, -1.88 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.06 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.14 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mistscape Mantle (4734, -0.58 DPS) [dungeon]; Batwing Mantle (6697, -0.58 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.05 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (19533, -0.32 DPS) [rep]; Darkspear Raider's Cloak (272077, -0.33 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.8 healing_power points (5.65 DPS) | yes | Death Speaker Robes (6682, +0.00 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -2.90 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -2.98 DPS) [crafted] |
| wrist | Reflective Wristguards (274750) | Rettrick [vendor] | 18.3 healing_power points (1.62 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.42 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.57 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 31.5 healing_power points (2.79 DPS) | yes | Gilded Handwraps (254021, -0.16 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.21 DPS) [crafted]; Gloves of the Greatfather (17721, -0.49 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.7 healing_power points (3.16 DPS) | yes | Mender's Leather Belt (252523, -0.20 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -0.49 DPS) [crafted]; Highlander's Lizardhide Girdle (20104, -1.61 DPS) [rep] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 34.3 healing_power points (3.04 DPS) | yes | Wisdom's Leather Leggings (252519, +0.00 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -0.05 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.39 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 38.5 healing_power points (3.41 DPS) | yes | Mender's Mail Boots (252565, -0.11 DPS) [crafted]; Mender's Leather Shoes (252533, -0.13 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -0.66 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.59 DPS) | yes | Snake Hoop (6750, -0.42 DPS) [quest]; Welken Ring (5011, -0.63 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 16.6 healing_power points (1.47 DPS) | yes | Snake Hoop (6750, -0.29 DPS) [quest]; Welken Ring (5011, -0.50 DPS) [world_drop]; Voodoo Band (1996, -0.55 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frost Tiger Blade (3854, -0.22 DPS, sim-verified) [crafted]; Haunting Blade (6641, -3.10 DPS) [dungeon]; Mograine's Might (7723, -4.37 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 24.9 healing_power points (2.21 DPS) | yes | Orb of Souls (249395, -0.09 DPS, sim-verified) [crafted]; Eye of Paleth (2943, -1.05 DPS) [quest]; Ravager's Shield (14777, -1.10 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Sheepshear Mantle; back: Mantle of Lady Falther'ess; chest: Stormcloth Vest; wrist: Reflective Wristguards; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger2: Darkspear Signet; trinket2: Ankh of Life; off_hand: Beacon of Hope

No-known-source sample (15 of 562, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 98.1. Weights run: 8.3s. Verify run: 3.7s. 722 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.027, intellect=1.926 ± 0.016, spirit=1.456 ± 0.013, mp5=4.768 ± 0.017, crit=0.317 ± 0.015 per rating point (14 rating = 1%, 4.437 per %), spell_haste=not significant (0.069 ± 0.033)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Imbued Helmet (220810, -0.18 DPS, sim-verified) [vendor]; Soulcatcher Halo (10630, -0.36 DPS) [dungeon]; Braincage (12549, -0.46 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkmoon Necklace (19303, -0.13 DPS, sim-verified) [vendor]; Gemshard Heart (17707, -0.33 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.50 DPS) [dungeon] |
| shoulder | Lead Surveyor's Mantle (11842) | Blackrock Depths: Fineous Darkvire [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Imbued Pauldrons (220808, -0.14 DPS, sim-verified) [vendor]; Mender's Leather Shoulder (252538, -0.15 DPS) [crafted]; Mender's Mail Shoulder (252569, -0.15 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 34.2 healing_power points (2.39 DPS) | yes | Featherskin Cape (10843, -0.33 DPS) [world]; Imperial Red Cloak (8248, -0.51 DPS) [world_drop]; Caretaker's Cape (19531, -0.60 DPS, sim-verified) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 76.4 healing_power points (5.34 DPS) | yes | Knight's Imbued Armor (220813, -0.34 DPS, sim-verified) [vendor]; Robes of Insight (940, -0.45 DPS) [world_drop]; Stormcloth Vest (10020, -0.47 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 41.2 healing_power points (2.88 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.20 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Gauntlets (252551, -0.20 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.20 DPS) [crafted]; Soulforge Fists (226982, -0.75 DPS, sim-verified) [vendor] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 53.1 healing_power points (3.71 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.27 DPS) [crafted]; Gilded Waistcord (254081, -0.47 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight's Imbued Leggings (220809, -0.24 DPS, sim-verified) [vendor]; Dalewind Trousers (13008, -0.50 DPS) [world_drop]; Windscale Sarong (10842, -0.87 DPS) [world] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 49.2 healing_power points (3.44 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.06 DPS) [crafted]; Mender's Leather Shoes (252533, -0.11 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 36.2 healing_power points (2.53 DPS) | yes | Eye of Adaegus (5266, -0.50 DPS) [world_drop]; Chivalrous Signet (20505, -0.54 DPS) [quest]; Cyclopean Band (11824, -0.55 DPS) [dungeon] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 33.4 healing_power points (2.33 DPS) | yes | Eye of Adaegus (5266, +0.00 DPS, sim-verified) [world_drop]; Chivalrous Signet (20505, -0.34 DPS) [quest]; Cyclopean Band (11824, -0.35 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -3.74 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.98 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.18 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.20 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.41 DPS) [quest] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Haunting Blade (6641, -2.44 DPS) [dungeon]; Mograine's Might (7723, -2.60 DPS) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | 38.3 healing_power points (2.68 DPS) | yes | Enthralled Sphere (11625, -0.29 DPS) [dungeon]; Cloud Stone (17737, -0.32 DPS) [dungeon]; Twisting Essence Jar (249456, -0.54 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** neck: Horizon Choker; shoulder: Lead Surveyor's Mantle; back: Darkspear Raider's Cloak; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Raider Handwraps; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 722, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 177.5. Weights run: 9.6s. Verify run: 4.1s. 1674 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.024, intellect=2.838 ± 0.026, spirit=2.289 ± 0.021, mp5=7.125 ± 0.028, crit=0.538 ± 0.026 per rating point (14 rating = 1%, 7.532 per %), spell_haste=not significant (0.203 ± 0.066)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Lamellar Helmet (227149, -1.10 DPS) [vendor]; Living Crown (252561, -1.18 DPS, sim-verified) [crafted]; Soulforge Crown (226981, -1.69 DPS) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 98.0 healing_power points (6.19 DPS) | yes | Lady Maye's Pendant (14558, +0.00 DPS, sim-verified) [world_drop]; Jeweled Amulet of Cainwyn (1443, -1.52 DPS) [world_drop]; Tooth of Gnarr (13141, -2.33 DPS) [dungeon] |
| shoulder | Shimmering Dawnbringer Shoulders (227859) | Argent Quartermaster Hasana [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Lamellar Pauldrons (227148, +0.00 DPS) [vendor]; Devout Mantle (16695, -0.23 DPS) [dungeon]; Argent Elite Shoulders (227888, -0.24 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 70.4 healing_power points (4.44 DPS) | yes | Faded Hakkari Cloak (20218, -0.10 DPS, sim-verified) [quest]; Gracious Cape (18743, -0.49 DPS) [dungeon]; Frostweaver Cape (12968, -0.56 DPS) [dungeon] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Lamellar Chestplate (227151, -1.60 DPS) [vendor]; Mooncloth Vest (14138, -1.83 DPS) [crafted]; Breastplate of Undead Slaying (23087, -2.71 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, -0.21 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.44 DPS) [dungeon]; Bracers of Undead Slaying (23090, -1.72 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -0.21 DPS) [dungeon]; Soulforge Fists (226982, -0.27 DPS) [vendor]; Razor Gauntlets (18326, -2.96 DPS, sim-verified) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom of the Timbermaw (19047, -0.43 DPS) [crafted]; Whipvine Cord (18327, -0.58 DPS, sim-verified) [dungeon]; Devout Belt (16696, -0.62 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, -0.66 DPS) [quest]; Padre's Trousers (18386, -0.78 DPS) [dungeon]; Cloudkeeper Legplates (14554, -2.62 DPS, sim-verified) [world_drop] |
| feet | Soulforge Treads (226983) | Mokvar [vendor] | 116.9 healing_power points (7.38 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.44 DPS, sim-verified) [vendor]; Knight-Lieutenant's Lamellar Greaves (227153, -1.15 DPS) [vendor]; Mooncloth Boots (15802, -1.48 DPS) [crafted] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, -0.73 DPS) [dungeon]; Rosewine Circle (13178, -0.90 DPS) [dungeon]; Naglering (11669, -2.03 DPS, sim-verified) [dungeon] |
| finger2 | Ring of Demonic Guile (18314) | Dire Maul: Alzzin the Wildshaper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, -0.24 DPS) [dungeon]; Rosewine Circle (13178, -0.41 DPS) [dungeon]; Naglering (11669, -1.93 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -4.05 DPS) [quest]; Serenity Field (272439, -4.09 DPS) [vendor]; Briarwood Reed (12930, -4.91 DPS) [dungeon] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.73 DPS, sim-verified) [quest]; Serenity Field (272439, -2.30 DPS) [vendor]; Briarwood Reed (12930, -3.12 DPS) [dungeon] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, -1.62 DPS, sim-verified) [world_drop]; Hammer of Divine Might (22333, -2.96 DPS) [dungeon]; Death Speaker Scepter (2816, -3.05 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Shimmering Dawnbringer Shoulders; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Martyr's Legplates; feet: Soulforge Treads; finger1: Band of Piety; finger2: Ring of Demonic Guile; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60, raid preset (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 314.2. Weights run: 7.8s. Verify run: 2.8s. 1674 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.092, intellect=1.913 ± 0.024, spirit=0.923 ± 0.016, mp5=4.768 ± 0.027, crit=0.633 ± 0.024 per rating point (14 rating = 1%, 8.857 per %), spell_haste=-0.677 ± 0.161

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Lamellar Helmet (227149, +0.00 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, -1.11 DPS) [vendor]; Living Crown (252561, -7.26 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 73.6 healing_power points (9.20 DPS) | yes | Lady Maye's Pendant (14558, -3.51 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -3.74 DPS) [world_drop]; Drake Tooth Necklace (21531, -4.88 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 96.8 healing_power points (12.10 DPS) | yes | Shimmering Dawnbringer Shoulders (227859, -1.23 DPS) [vendor]; Lieutenant Commander's Lamellar Pauldrons (227148, -2.34 DPS) [vendor]; Knight-Lieutenant's Imbued Pauldrons (220808, -4.34 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 61.1 healing_power points (7.64 DPS) | yes | Cloak of the Cosmos (18389, -1.76 DPS) [dungeon]; Faded Hakkari Cloak (20218, -2.15 DPS) [quest]; Drape of Recovery (272413, -2.29 DPS, sim-verified) [vendor] |
| chest | Knight-Captain's Lamellar Chestplate (227151) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Breastplate of Salvation (250601, -0.64 DPS) [crafted]; Field Marshal's Lamellar Chestplate (231641, -1.13 DPS) [pvp]; Breastplate of Undead Slaying (23087, -13.11 DPS, sim-verified) [world] |
| wrist | Gallant's Wristguards (18459) | Dire Maul: Guard Fengus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Hope (22667, -0.71 DPS) [quest]; Bracers of Mending (23129, -1.07 DPS) [dungeon]; Bracers of Undead Slaying (23090, -9.65 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -0.84 DPS) [dungeon]; Harmonious Gauntlets (18527, -1.31 DPS) [dungeon]; Razor Gauntlets (18326, -15.13 DPS, sim-verified) [dungeon] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 76.8 healing_power points (9.60 DPS) | yes | Belt of Tiny Heads (20217, -1.37 DPS) [quest]; Eyestalk Cord (18391, -1.64 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.35 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Padre's Trousers (18386, -0.94 DPS) [dungeon]; Knight-Captain's Lamellar Legguards (227150, -1.51 DPS) [vendor]; Cloudkeeper Legplates (14554, -13.28 DPS, sim-verified) [world_drop] |
| feet | Knight-Lieutenant's Lamellar Greaves (227153) | Captain Dirgehammer [vendor] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulforge Treads (226983, -0.37 DPS) [vendor]; Mooncloth Boots (15802, -1.63 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -2.12 DPS, sim-verified) [vendor] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -1.28 DPS) [dungeon]; Ring of Demonic Guile (18314, -1.56 DPS) [dungeon]; Naglering (11669, -10.80 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.36 DPS) [dungeon]; Ring of Demonic Guile (18314, -0.64 DPS) [dungeon]; Naglering (11669, -7.59 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+14.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -3.69 DPS) [vendor]; Briarwood Reed (12930, -5.31 DPS) [dungeon]; Shard of the Splithooves (10659, -5.36 DPS) [quest] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.31 DPS) [vendor]; Briarwood Reed (12930, -2.93 DPS) [dungeon]; Shard of the Splithooves (10659, -4.91 DPS, sim-verified) [quest] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -1.51 DPS) [dungeon]; Hammer of Divine Might (22333, -3.04 DPS) [dungeon]; Hand of Edward the Odd (2243, -8.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Knight-Captain's Lamellar Chestplate; wrist: Gallant's Wristguards; hands: Raider Handwraps; waist: Whipvine Cord; legs: Martyr's Legplates; feet: Knight-Lieutenant's Lamellar Greaves; finger1: Band of Piety; finger2: Rosewine Circle; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 25.5. Weights run: 4.7s. Verify run: 2.3s. 219 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.002, intellect=0.897 ± 0.002, spirit=0.227 ± 0.001, mp5=1.697 ± 0.023, crit=0.037 ± 0.001 per rating point (14 rating = 1%, 0.515 per %), spell_haste=0.067 ± 0.009

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.74 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.33 DPS) [crafted]; Trapper's Leather Hood (252505, -1.14 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.9 healing_power points (0.08 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.67 DPS) | yes | Slime-encrusted Pads (6461, -0.21 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.37 DPS) [crafted]; Forest Leather Mantle (4709, -0.59 DPS) [world_drop] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (1.6 DPS) | yes | Pearl-clasped Cloak (5542, -0.07 DPS) [crafted]; Battle Healer's Cloak (20427, -0.08 DPS, sim-verified) [rep]; Seer's Cape (6378, -0.11 DPS) [dungeon] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 20.5 healing_power points (1.69 DPS) | yes | Wisdom's Leather Armor (252493, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.19 DPS) [crafted]; Bloody Apron (6226, -0.62 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 9.6 healing_power points (0.79 DPS) | yes | Tabitha's Cuffs (251486, +0.00 DPS, sim-verified) [quest]; Garrison Cuffs (270003, -0.39 DPS) [quest]; Owl Bracers (4796, -0.42 DPS) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.5 healing_power points (1.20 DPS) | yes | Wisdom's Leather Gloves (252499, -0.07 DPS) [crafted]; Pristine Gloves (253913, -0.08 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.58 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.21 DPS) | yes | Acolyte's Chain Belt (250516, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.25 DPS) [crafted]; Novice Ardent's Sash (253887, -0.43 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 27.3 healing_power points (2.25 DPS) | yes | Wisdom's Leather Pants (252503, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.25 DPS) [crafted]; Ghastly Trousers (15449, -1.60 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.5 healing_power points (1.28 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.24 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.44 DPS) | yes | Advisor's Ring (20426, -0.16 DPS) [rep]; Black Pearl Ring (6332, -0.18 DPS) [world]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 4.5 healing_power points (0.37 DPS) | yes | Advisor's Ring (20426, -0.08 DPS, sim-verified) [rep]; Black Pearl Ring (6332, -0.11 DPS) [world]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 4.5 healing_power points (0.37 DPS) | yes | Samophlange Screwdriver (11854, -0.17 DPS, sim-verified) [quest]; Rakzur Club (12983, -0.28 DPS) [world_drop]; Mug of Muddled Memories (277247, -0.28 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Impaling Harpoon

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 48.7. Weights run: 4.6s. Verify run: 2.2s. 381 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.979 ± 0.003, spirit=0.489 ± 0.003, mp5=2.096 ± 0.007, crit=0.092 ± 0.004 per rating point (14 rating = 1%, 1.285 per %), spell_haste=not significant (0.034 ± 0.008)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Chain Helm (250501) | Blacksmithing [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.20 DPS, sim-verified) [world_drop]; Filigreed Pristine Circlet (253975, -0.28 DPS) [crafted]; Wisdom's Leather Helm (252515, -0.37 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 5.9 healing_power points (0.54 DPS) | yes | Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.06 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.09 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Nightsky Mantle (4718, -0.27 DPS) [world_drop]; Ghostly Mantle (3324, -0.69 DPS, sim-verified) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 15.0 healing_power points (1.38 DPS) | yes | Glowing Thresher Cape (6901, -0.08 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -0.52 DPS) [vendor]; Cloak of Rot (4462, -0.66 DPS) [world] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 30.8 healing_power points (2.84 DPS) | yes | Acolyte's Silvered Chain Shirt (250521, +0.00 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.18 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.37 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.2 healing_power points (1.03 DPS) | yes | Glowing Magical Bracelets (13106, +0.00 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Spidertank Oilrag (9448, -0.45 DPS) [dungeon] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.9 healing_power points (1.66 DPS) | yes | Acolyte's Gloves (250511, -0.28 DPS) [crafted]; Naga Battle Gloves (888, -0.33 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.37 DPS) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 28.9 healing_power points (2.67 DPS) | yes | Mender's Leather Belt (252523, +0.00 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.29 DPS) [crafted]; Acolyte's Chain Belt (250516, -1.47 DPS) [crafted] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 32.8 healing_power points (3.03 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Leggings (252519, -0.14 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.46 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.8 healing_power points (2.66 DPS) | yes | Acolyte's Boots (250506, -0.25 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -1.19 DPS) [crafted]; Glowing Copper Boots (250482, -1.47 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.66 DPS) | yes | Darkspear Signet (272071, -0.89 DPS) [vendor]; Black Widow Band (6199, -1.03 DPS) [world]; Lavishly Jeweled Ring (1156, -1.12 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 healing_power points (0.95 DPS) | yes | Darkspear Signet (272071, -0.22 DPS, sim-verified) [vendor]; Black Widow Band (6199, -0.32 DPS) [world]; Lavishly Jeweled Ring (1156, -0.41 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -0.17 DPS, sim-verified) [dungeon]; Haunting Blade (6641, -3.23 DPS) [dungeon]; Wolfsbane (267369, -5.38 DPS) [quest] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 15.9 healing_power points (1.47 DPS) | yes | Alliance Outrunner Healing Rod (285348, -0.19 DPS, sim-verified) [world]; Orb of Mistmantle (13031, -0.46 DPS) [world_drop]; Defective Samophlange (274743, -0.79 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Acolyte's Chain Helm; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 381, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 68.1. Weights run: 7.8s. Verify run: 3.5s. 531 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.009, intellect=1.167 ± 0.006, spirit=0.727 ± 0.005, mp5=2.764 ± 0.008, crit=0.140 ± 0.006 per rating point (14 rating = 1%, 1.960 per %), spell_haste=not significant (0.045 ± 0.018)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 48.4 healing_power points (4.29 DPS) | yes | Holy Shroud (2721, -1.41 DPS, sim-verified) [world_drop]; Earthen Silk Hood (254015, -1.69 DPS) [crafted]; Miner's Hat of the Deep (9429, -1.88 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.06 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.14 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mistscape Mantle (4734, -0.58 DPS) [dungeon]; Batwing Mantle (6697, -0.58 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.03 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Battle Healer's Cloak (19529, -0.32 DPS) [rep]; Darkspear Raider's Cloak (272077, -0.33 DPS) [vendor] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Vest (10020, -0.05 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.16 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -0.24 DPS) [crafted] |
| wrist | Reflective Wristguards (274750) | Rettrick [vendor] | 18.3 healing_power points (1.62 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.42 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.57 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 31.5 healing_power points (2.79 DPS) | yes | Gilded Handwraps (254021, -0.16 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.21 DPS) [crafted]; Gloves of the Greatfather (17721, -0.49 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.7 healing_power points (3.16 DPS) | yes | Mender's Leather Belt (252523, -0.20 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -0.49 DPS) [crafted]; Highlander's Mail Girdle (20119, -1.61 DPS) [vendor] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 34.3 healing_power points (3.04 DPS) | yes | Wisdom's Leather Leggings (252519, +0.00 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -0.05 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.39 DPS) [crafted] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 38.5 healing_power points (3.41 DPS) | yes | Mender's Leather Shoes (252533, -0.10 DPS, sim-verified) [crafted]; Mender's Mail Boots (252565, -0.11 DPS) [crafted]; Gilded Slippers (254001, -0.66 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.59 DPS) | yes | Snake Hoop (6750, -0.42 DPS) [quest]; Welken Ring (5011, -0.63 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 16.6 healing_power points (1.47 DPS) | yes | Snake Hoop (6750, -0.29 DPS) [quest]; Welken Ring (5011, -0.50 DPS) [world_drop]; Voodoo Band (1996, -0.55 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frost Tiger Blade (3854, -0.22 DPS, sim-verified) [crafted]; Haunting Blade (6641, -3.10 DPS) [dungeon]; Mograine's Might (7723, -4.37 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 24.9 healing_power points (2.21 DPS) | yes | Orb of Souls (249395, -0.08 DPS, sim-verified) [crafted]; Prophetic Cane (6803, -0.97 DPS) [quest]; Ravager's Shield (14777, -1.10 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Sheepshear Mantle; back: Mantle of Lady Falther'ess; wrist: Reflective Wristguards; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger2: Darkspear Signet; trinket2: Ankh of Life; off_hand: Beacon of Hope

No-known-source sample (15 of 531, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 96.8. Weights run: 8.3s. Verify run: 3.7s. 702 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.027, intellect=1.926 ± 0.016, spirit=1.456 ± 0.013, mp5=4.768 ± 0.017, crit=0.317 ± 0.015 per rating point (14 rating = 1%, 4.437 per %), spell_haste=not significant (0.069 ± 0.033)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcatcher Halo (10630) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Braincage (12549, -0.10 DPS) [dungeon]; Helm of Exile (11124, -0.13 DPS) [quest]; Papal Fez (9431, -0.22 DPS, sim-verified) [dungeon] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 40.2 healing_power points (2.81 DPS) | yes | Horizon Choker (13085, +0.00 DPS, sim-verified) [world_drop]; Gemshard Heart (17707, -0.85 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.01 DPS) [dungeon] |
| shoulder | Lead Surveyor's Mantle (11842) | Blackrock Depths: Fineous Darkvire [dungeon] | 53.6 healing_power points (3.74 DPS) | yes | Mender's Mail Shoulder (252569, -0.15 DPS) [crafted]; Ironfeather Shoulders (15067, -0.24 DPS) [crafted]; Mender's Leather Shoulder (252538, -0.26 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 34.2 healing_power points (2.39 DPS) | yes | Featherskin Cape (10843, -0.33 DPS) [world]; Imperial Red Cloak (8248, -0.51 DPS) [world_drop]; Battle Healer's Cloak (19527, -0.67 DPS, sim-verified) [rep] |
| chest | Robes of Insight (940) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormcloth Vest (10020, -0.02 DPS) [crafted]; Ghostweave Vest (14141, -0.19 DPS) [crafted]; Embrace of the Wind Serpent (12462, -0.27 DPS, sim-verified) [world] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 41.2 healing_power points (2.88 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.20 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Gauntlets (252551, -0.20 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.20 DPS) [crafted]; Soulforge Fists (226982, -0.72 DPS, sim-verified) [vendor] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 53.1 healing_power points (3.71 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.27 DPS) [crafted]; Gilded Waistcord (254081, -0.47 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 60.9 healing_power points (4.25 DPS) | yes | Dalewind Trousers (13008, -0.49 DPS, sim-verified) [world_drop]; Windscale Sarong (10842, -0.87 DPS) [world]; Jinxed Hoodoo Kilt (9474, -0.97 DPS) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 49.2 healing_power points (3.44 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.06 DPS) [crafted]; Mender's Leather Shoes (252533, -0.11 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 36.2 healing_power points (2.53 DPS) | yes | Eye of Adaegus (5266, -0.50 DPS) [world_drop]; Cyclopean Band (11824, -0.55 DPS) [dungeon]; Snake Hoop (6750, -0.87 DPS) [quest] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 33.4 healing_power points (2.33 DPS) | yes | Eye of Adaegus (5266, -0.06 DPS, sim-verified) [world_drop]; Cyclopean Band (11824, -0.35 DPS) [dungeon]; Snake Hoop (6750, -0.68 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -3.74 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.98 DPS) [quest]; Alchemists' Stone (13503, -4.59 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS, sim-verified) [quest]; Alchemists' Stone (13503, -0.81 DPS) [crafted] |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Haunting Blade (6641, -2.44 DPS) [dungeon]; Mograine's Might (7723, -2.60 DPS) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | 38.3 healing_power points (2.68 DPS) | yes | Enthralled Sphere (11625, +0.00 DPS, sim-verified) [dungeon]; Cloud Stone (17737, -0.32 DPS) [dungeon]; Twisting Essence Jar (249456, -0.54 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Soulcatcher Halo; neck: Darkmoon Necklace; shoulder: Lead Surveyor's Mantle; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Mender's Leather Bracers; hands: Raider Handwraps; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 702, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 171.0. Weights run: 9.6s. Verify run: 4.1s. 1699 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.024, intellect=2.838 ± 0.026, spirit=2.289 ± 0.021, mp5=7.125 ± 0.028, crit=0.538 ± 0.026 per rating point (14 rating = 1%, 7.532 per %), spell_haste=not significant (0.203 ± 0.066)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Living Crown (252561, -1.23 DPS, sim-verified) [crafted]; Soulforge Crown (226981, -1.69 DPS) [vendor]; Crown of the Penitent (13216, -1.80 DPS) [quest] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Jeweled Amulet of Cainwyn (1443, -0.18 DPS) [world_drop]; Wavefront Necklace (20685, -0.21 DPS, sim-verified) [world]; Tooth of Gnarr (13141, -0.99 DPS) [dungeon] |
| shoulder | Shimmering Dawnbringer Shoulders (227859) | Argent Quartermaster Hasana [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Devout Mantle (16695, -0.23 DPS) [dungeon]; Argent Elite Shoulders (227888, -0.31 DPS, sim-verified) [vendor]; Soulforge Epaulets (226979, -0.54 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 70.4 healing_power points (4.44 DPS) | yes | Faded Hakkari Cloak (20218, -0.11 DPS, sim-verified) [quest]; Gracious Cape (18743, -0.49 DPS) [dungeon]; Frostweaver Cape (12968, -0.56 DPS) [dungeon] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, -1.83 DPS) [crafted]; Alanna's Embrace (13314, -2.07 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -2.59 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, -0.21 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.44 DPS) [dungeon]; Bracers of Undead Slaying (23090, -1.68 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -0.21 DPS) [dungeon]; Soulforge Fists (226982, -0.27 DPS) [vendor]; Razor Gauntlets (18326, -2.87 DPS, sim-verified) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom of the Timbermaw (19047, -0.43 DPS) [crafted]; Whipvine Cord (18327, -0.61 DPS, sim-verified) [dungeon]; Devout Belt (16696, -0.62 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, -0.66 DPS) [quest]; Padre's Trousers (18386, -0.78 DPS) [dungeon]; Cloudkeeper Legplates (14554, -2.47 DPS, sim-verified) [world_drop] |
| feet | Soulforge Treads (226983) | Mokvar [vendor] | 116.9 healing_power points (7.38 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.36 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -1.48 DPS) [crafted]; Faith Healer's Boots (22247, -1.86 DPS) [dungeon] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, -0.73 DPS) [dungeon]; Rosewine Circle (13178, -0.90 DPS) [dungeon]; Naglering (11669, -1.85 DPS, sim-verified) [dungeon] |
| finger2 | Ring of Demonic Guile (18314) | Dire Maul: Alzzin the Wildshaper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, -0.24 DPS) [dungeon]; Rosewine Circle (13178, -0.41 DPS) [dungeon]; Naglering (11669, -1.82 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -4.05 DPS) [quest]; Serenity Field (272439, -4.09 DPS) [vendor]; Briarwood Reed (12930, -4.91 DPS) [dungeon] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, -0.77 DPS, sim-verified) [quest]; Shard of the Splithooves (10659, -2.25 DPS) [quest]; Serenity Field (272439, -2.30 DPS) [vendor] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, -1.58 DPS, sim-verified) [world_drop]; Hammer of Divine Might (22333, -2.96 DPS) [dungeon]; Death Speaker Scepter (2816, -3.05 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Lady Maye's Pendant; shoulder: Shimmering Dawnbringer Shoulders; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Martyr's Legplates; feet: Soulforge Treads; finger1: Band of Piety; finger2: Ring of Demonic Guile; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 323.0. Weights run: 7.8s. Verify run: 2.8s. 1699 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.092, intellect=1.913 ± 0.024, spirit=0.923 ± 0.016, mp5=4.768 ± 0.027, crit=0.633 ± 0.024 per rating point (14 rating = 1%, 8.857 per %), spell_haste=-0.677 ± 0.161

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulforge Crown (226981, -1.63 DPS) [vendor]; Sanctified Leather Helm (22689, -1.85 DPS) [quest]; Living Crown (252561, -6.67 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 73.6 healing_power points (9.20 DPS) | yes | Lady Maye's Pendant (14558, -3.51 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -3.74 DPS) [world_drop]; Drake Tooth Necklace (21531, -5.30 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 96.8 healing_power points (12.10 DPS) | yes | Shimmering Dawnbringer Shoulders (227859, -1.23 DPS) [vendor]; Devout Mantle (16695, -4.67 DPS) [dungeon]; Soulforge Epaulets (226979, -4.86 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 61.1 healing_power points (7.64 DPS) | yes | Cloak of the Cosmos (18389, -1.76 DPS) [dungeon]; Faded Hakkari Cloak (20218, -2.15 DPS) [quest]; Drape of Recovery (272413, -2.67 DPS, sim-verified) [vendor] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of the Exalted (13346, -1.19 DPS) [dungeon]; Mooncloth Vest (14138, -1.93 DPS) [crafted]; Breastplate of Undead Slaying (23087, -16.13 DPS, sim-verified) [world] |
| wrist | Gallant's Wristguards (18459) | Dire Maul: Guard Fengus [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Hope (22667, -0.71 DPS) [quest]; Bracers of Mending (23129, -1.07 DPS) [dungeon]; Bracers of Undead Slaying (23090, -10.46 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -0.84 DPS) [dungeon]; Harmonious Gauntlets (18527, -1.31 DPS) [dungeon]; Razor Gauntlets (18326, -15.82 DPS, sim-verified) [dungeon] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 76.8 healing_power points (9.60 DPS) | yes | Belt of Tiny Heads (20217, -1.37 DPS) [quest]; Eyestalk Cord (18391, -1.64 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.37 DPS, sim-verified) [crafted] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Padre's Trousers (18386, -0.94 DPS) [dungeon]; Leggings of Arcana (12756, -4.06 DPS) [quest]; Cloudkeeper Legplates (14554, -14.38 DPS, sim-verified) [world_drop] |
| feet | Soulforge Treads (226983) | Mokvar [vendor] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Boots (15802, -1.26 DPS) [crafted]; Faith Healer's Boots (22247, -1.60 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -2.49 DPS, sim-verified) [vendor] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -1.28 DPS) [dungeon]; Ring of Demonic Guile (18314, -1.56 DPS) [dungeon]; Naglering (11669, -11.26 DPS, sim-verified) [dungeon] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.36 DPS) [dungeon]; Ring of Demonic Guile (18314, -0.64 DPS) [dungeon]; Naglering (11669, -8.42 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+14.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -3.69 DPS) [vendor]; Briarwood Reed (12930, -5.31 DPS) [dungeon]; Shard of the Splithooves (10659, -5.36 DPS) [quest] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.31 DPS) [vendor]; Briarwood Reed (12930, -2.93 DPS) [dungeon]; Shard of the Splithooves (10659, -5.27 DPS, sim-verified) [quest] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -1.51 DPS) [dungeon]; Hammer of Divine Might (22333, -3.04 DPS) [dungeon]; Hand of Edward the Odd (2243, -9.05 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Gallant's Wristguards; hands: Raider Handwraps; waist: Whipvine Cord; legs: Martyr's Legplates; feet: Soulforge Treads; finger1: Band of Piety; finger2: Rosewine Circle; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

