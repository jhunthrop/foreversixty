# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 27.3. Weights run: 4.2s. Verify run: 13.7s. 239 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.899 ± 0.002, spirit=0.222 ± 0.001, mp5=1.695 ± 0.023, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.497 per %), spell_haste=0.105 ± 0.012

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.76 DPS) | yes | Wisdom's Leather Hood (252507, -0.34 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.16 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 0.9 healing_power points (0.07 DPS) | yes | Scholarly Pendant (277203, -0.19 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.38 DPS) [crafted]; Slime-encrusted Pads (6461, -0.40 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.60 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 9.4 healing_power points (0.79 DPS) | yes | Pearl-clasped Cloak (5542, -0.57 DPS) [crafted]; Seer's Cape (6378, -0.60 DPS) [dungeon]; Sanguine Cape (14376, -0.83 DPS, sim-verified) [world_drop] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 20.5 healing_power points (1.72 DPS) | yes | Wisdom's Leather Armor (252493, -0.13 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.63 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 9.6 healing_power points (0.80 DPS) | yes | Bright Bracers (3647, -0.50 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.50 DPS) [vendor]; Owl Bracers (4796, -1.05 DPS, sim-verified) [vendor] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.5 healing_power points (1.22 DPS) | yes | Wisdom's Leather Gloves (252499, -0.08 DPS) [crafted]; Pristine Gloves (253913, -0.23 DPS, sim-verified) [crafted]; Tomb Robber's Gloves (280096, -0.76 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.22 DPS) | yes | Wisdom's Leather Belt (252433, -0.25 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.26 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.44 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 27.3 healing_power points (2.29 DPS) | yes | Filigreed Pristine Leggings (253937, -0.25 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.44 DPS, sim-verified) [crafted]; Dreamer's Leggings (270016, -1.46 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.5 healing_power points (1.30 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.24 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.45 DPS) | yes | Black Pearl Ring (6332, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop]; Minor Channeling Ring (1449, -0.30 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 3.4 healing_power points (0.28 DPS) | yes | Volcanic Rock Ring (12053, -0.06 DPS) [world_drop]; Minor Channeling Ring (1449, -0.13 DPS) [quest]; Black Pearl Ring (6332, -0.34 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 8.1 healing_power points (0.68 DPS) | yes | Impaling Harpoon (5200, -0.30 DPS) [dungeon]; Verigan's Fist (6953, -0.33 DPS, sim-verified) [quest]; Monastic Hammer (270005, -0.53 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Lorekeeper's Ring; main_hand: Sword of the Fallen

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 50.1. Weights run: 4.1s. Verify run: 13.9s. 431 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.977 ± 0.003, spirit=0.493 ± 0.003, mp5=2.099 ± 0.007, crit=0.093 ± 0.004 per rating point (14 rating = 1%, 1.295 per %), spell_haste=0.041 ± 0.009

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.0 healing_power points (3.37 DPS) | yes | Filigreed Pristine Circlet (253975, -1.31 DPS) [crafted]; Embalmed Shroud (7691, -1.32 DPS) [dungeon]; Acolyte's Chain Helm (250501, -2.20 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.9 healing_power points (0.55 DPS) | yes | Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.09 DPS) [vendor]; Scorn's Icy Choker (23169, -0.76 DPS, sim-verified) [dungeon] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 13.2 healing_power points (1.24 DPS) | yes | Mantle of Honor (3560, -0.27 DPS) [quest]; Nightsky Mantle (4718, -0.27 DPS) [world_drop]; Death Speaker Mantle (6685, -1.14 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | 15.0 healing_power points (1.40 DPS) | yes | Repairman's Cape (9605, -0.15 DPS) [quest]; Glowing Thresher Cape (6901, -0.38 DPS) [dungeon]; Prelacy Cape (7004, -0.39 DPS, sim-verified) [quest] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 30.7 healing_power points (2.88 DPS) | yes | Pristine Gown (253961, -0.18 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.37 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -0.52 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.2 healing_power points (1.05 DPS) | yes | Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.39 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.9 healing_power points (1.68 DPS) | yes | Acolyte's Gloves (250511, -0.29 DPS) [crafted]; Silvered Gauntlets (270025, -0.37 DPS) [quest]; Naga Battle Gloves (888, -0.44 DPS, sim-verified) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 28.8 healing_power points (2.70 DPS) | yes | Mender's Leather Belt (252523, -0.92 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.31 DPS) [crafted]; Acolyte's Chain Belt (250516, -1.49 DPS) [crafted] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 32.8 healing_power points (3.08 DPS) | yes | Wisdom's Leather Leggings (252519, -0.14 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.47 DPS) [crafted]; Pristine Leggings (253987, -0.98 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.8 healing_power points (2.70 DPS) | yes | Acolyte's Boots (250506, -1.21 DPS) [crafted]; Wisdom's Leather Boots (252444, -1.21 DPS) [crafted]; Nimbus Boots (6998, -2.59 DPS, sim-verified) [quest] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.69 DPS) | yes | Electrocutioner Lagnut (9447, -0.80 DPS) [dungeon]; Darkspear Signet (272071, -0.90 DPS) [vendor]; Black Widow Band (6199, -1.05 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 healing_power points (0.96 DPS) | yes | Darkspear Signet (272071, -0.18 DPS) [vendor]; Black Widow Band (6199, -0.32 DPS) [world]; Electrocutioner Lagnut (9447, -0.54 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | sim-verified (50.1 DPS) | yes | Death Speaker Scepter (2816, -0.98 DPS) [dungeon]; Manual Crowd Pummeler (9449, -2.71 DPS, sim-verified) [dungeon]; Haunting Blade (6641, -4.26 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.0 healing_power points (1.50 DPS) | yes | Alliance Outrunner Healing Rod (285348, -0.37 DPS) [world]; Orb of Mistmantle (13031, -0.47 DPS) [world_drop]; Eye of Paleth (2943, -1.06 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Caretaker's Cape; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; main_hand: Royal Diplomatic Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 431, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 79.3. Weights run: 7.4s. Verify run: 29.4s. 589 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.162 ± 0.006, spirit=0.735 ± 0.005, mp5=2.690 ± 0.008, crit=0.150 ± 0.006 per rating point (14 rating = 1%, 2.097 per %), spell_haste=not significant (0.060 ± 0.017)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 48.4 healing_power points (4.38 DPS) | yes | Earthen Silk Hood (254015, -1.72 DPS) [crafted]; Miner's Hat of the Deep (9429, -1.93 DPS) [dungeon]; Holy Shroud (2721, -2.26 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.3 healing_power points (2.02 DPS) | yes | Necklace of Calisea (1714, -0.82 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.97 DPS) [quest]; Triune Amulet (7722, -1.43 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 30.1 healing_power points (2.73 DPS) | yes | Sheepshear Mantle (13115, -0.64 DPS) [world_drop]; Mistscape Mantle (4734, -1.24 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.38 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | 21.7 healing_power points (1.96 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.20 DPS) [dungeon]; Repairman's Cape (9605, -0.53 DPS) [quest]; Darkspear Raider's Cloak (272077, -0.54 DPS) [vendor] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.9 healing_power points (5.78 DPS) | yes | Death Speaker Robes (6682, -2.82 DPS) [dungeon]; Doomsayer's Robe (4746, -2.85 DPS) [quest]; Deathchill Armor (10764, -4.26 DPS, sim-verified) [dungeon] |
| wrist | Reflective Wristguards (274750) | Rettrick [vendor] | 18.3 healing_power points (1.66 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.42 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.57 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 31.5 healing_power points (2.85 DPS) | yes | Gilded Handwraps (254021, -0.11 DPS, sim-verified) [crafted]; Prefect's Gauntlet (250569, -0.21 DPS) [crafted]; Gloves of the Greatfather (17721, -0.49 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.7 healing_power points (3.23 DPS) | yes | Prefect's Belt (250559, -0.50 DPS) [crafted]; Mender's Leather Belt (252523, -0.73 DPS, sim-verified) [crafted]; Highlander's Lizardhide Girdle (20104, -1.65 DPS) [rep] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 34.3 healing_power points (3.11 DPS) | yes | Wisdom's Leather Leggings (252519, -0.00 DPS) [crafted]; Pristine Leggings (253987, -0.04 DPS) [crafted]; Stoneweaver Leggings (9407, -0.29 DPS) [dungeon] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 38.5 healing_power points (3.48 DPS) | yes | Mender's Mail Boots (252565, -0.11 DPS) [crafted]; Mender's Leather Shoes (252533, -0.23 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -0.67 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.63 DPS) | yes | Snake Hoop (6750, -0.43 DPS) [quest]; Electrocutioner Lagnut (9447, -0.62 DPS) [dungeon]; Welken Ring (5011, -0.64 DPS) [world_drop] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 16.1 healing_power points (1.46 DPS) | yes | Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon]; Welken Ring (5011, -0.47 DPS) [world_drop]; Snake Hoop (6750, -0.98 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (79.3 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (79.3 DPS) | yes | Royal Diplomatic Scepter (9457, -1.79 DPS) [dungeon]; Death Speaker Scepter (2816, -2.89 DPS) [dungeon]; Frost Tiger Blade (3854, -3.23 DPS, sim-verified) [crafted] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 23.6 healing_power points (2.14 DPS) | yes | Beacon of Hope (9393, -0.61 DPS) [dungeon]; Eye of Paleth (2943, -0.96 DPS) [quest]; Orb of Souls (249395, -1.33 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Reflective Wristguards; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Orb of Lorica

No-known-source sample (15 of 589, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 125.9. Weights run: 7.8s. Verify run: 39.3s. 753 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=1.850 ± 0.013, spirit=1.414 ± 0.011, mp5=4.678 ± 0.015, crit=0.311 ± 0.015 per rating point (14 rating = 1%, 4.352 per %), spell_haste=not significant (0.102 ± 0.042)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Imbued Helmet (220810) | Captain Dirgehammer [vendor] | 78.3 healing_power points (5.61 DPS) | yes | Gemburst Circlet (10751, -0.89 DPS) [quest]; Soulcatcher Halo (10630, -1.28 DPS) [dungeon]; Papal Fez (9431, -2.37 DPS, sim-verified) [dungeon] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 39.2 healing_power points (2.81 DPS) | yes | Glowing Eye of Mordresh (10769, -0.81 DPS) [dungeon]; Gemshard Heart (17707, -0.87 DPS) [dungeon]; Horizon Choker (13085, -1.46 DPS, sim-verified) [world_drop] |
| shoulder | Knight-Lieutenant's Imbued Pauldrons (220808) | Captain Dirgehammer [vendor] | 61.1 healing_power points (4.37 DPS) | yes | Mender's Leather Shoulder (252538, -0.75 DPS) [crafted]; Mender's Mail Shoulder (252569, -0.75 DPS) [crafted]; Lead Surveyor's Mantle (11842, -1.84 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 33.0 healing_power points (2.36 DPS) | yes | Caretaker's Cape (19531, +0.00 DPS, sim-verified) [rep]; Featherskin Cape (10843, -0.31 DPS) [world]; Imperial Red Cloak (8248, -0.50 DPS) [world_drop] |
| chest | Knight's Imbued Armor (220813) | Captain Dirgehammer [vendor] | sim-verified (125.9 DPS) | yes | Stormcloth Vest (10020, -0.18 DPS) [crafted]; Robes of Insight (940, -0.31 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -1.98 DPS, sim-verified) [world] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 40.4 healing_power points (2.90 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Soulforge Fists (226982) | Mokvar [vendor] | 69.8 healing_power points (5.00 DPS) | yes | Raider Handwraps (272098, -0.51 DPS) [vendor]; Mender's Leather Gauntlets (252551, -0.68 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.68 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 52.2 healing_power points (3.74 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.27 DPS) [crafted]; Dawnspire Cord (12466, -0.36 DPS) [dungeon] |
| legs | Knight's Imbued Leggings (220809) | Captain Dirgehammer [vendor] | 71.8 healing_power points (5.14 DPS) | yes | Senior Designer's Pantaloons (11841, -0.60 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.75 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -2.43 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 48.4 healing_power points (3.46 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Leather Shoes (252533, -0.11 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.8 healing_power points (2.49 DPS) | yes | Eye of Adaegus (5266, -0.48 DPS) [world_drop]; Chivalrous Signet (20505, -0.49 DPS) [quest]; Cyclopean Band (11824, -0.52 DPS) [dungeon] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 32.7 healing_power points (2.35 DPS) | yes | Eye of Adaegus (5266, -0.34 DPS) [world_drop]; Chivalrous Signet (20505, -0.34 DPS) [quest]; Cyclopean Band (11824, -0.37 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -3.81 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.01 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.22 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.28 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.48 DPS) [quest]; Ankh of Life (1713, -1.21 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Diplomatic Scepter (9457, -1.46 DPS) [dungeon]; Death Speaker Scepter (2816, -2.67 DPS) [dungeon]; Hanzo Sword (8190, -15.45 DPS, sim-verified) [world_drop] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | 37.2 healing_power points (2.67 DPS) | yes | Cloud Stone (17737, -0.33 DPS) [dungeon]; Orb of Lorica (11262, -0.48 DPS) [quest]; Enthralled Sphere (11625, -16.19 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Imbued Helmet; neck: Darkmoon Necklace; shoulder: Knight-Lieutenant's Imbued Pauldrons; back: Darkspear Raider's Cloak; chest: Knight's Imbued Armor; wrist: Mender's Leather Bracers; hands: Soulforge Fists; waist: Mender's Leather Waistguard; legs: Knight's Imbued Leggings; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 753, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 203.5. Weights run: 9.0s. Verify run: 121.5s. 1685 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.126, intellect=4.925 ± 0.058, spirit=1.708 ± 0.014, mp5=7.069 ± 0.046, crit=0.442 ± 0.030 per rating point (14 rating = 1%, 6.193 per %), spell_haste=not significant (0.104 ± 0.054)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 216.3 healing_power points (13.08 DPS) | yes | Crown of the Penitent (13216, -1.80 DPS) [quest]; Living Crown (252561, -1.86 DPS) [crafted]; Black Dragonscale Helm (252605, -3.50 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-verified (203.6 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.26 DPS) [world_drop]; Tooth of Gnarr (13141, -1.21 DPS) [dungeon]; Lady Maye's Pendant (14558, -7.48 DPS, sim-verified) [world_drop] |
| shoulder | Soulforge Epaulets (226979) | Mokvar [vendor] | sim-verified (203.6 DPS) | yes | Darkspear Shoulders (272104, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.27 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 91.2 healing_power points (5.52 DPS) | yes | Shroud of the Exile (15421, -0.33 DPS) [quest]; Faded Hakkari Cloak (20218, -0.57 DPS) [quest]; Darkspear Raider's Cloak (272063, -5.43 DPS, sim-verified) [vendor] |
| chest | Knight-Captain's Lamellar Chestplate (227151) | Captain Dirgehammer [vendor] | sim-verified (203.6 DPS) | yes | Breastplate of Salvation (250601, +0.00 DPS) [crafted]; Devout Robe (16690, -0.08 DPS) [dungeon]; Soulforge Embrace (226984, -5.20 DPS, sim-verified) [vendor] |
| wrist | Soulforge Bindings (226977) | Mokvar [vendor] | sim-verified (203.6 DPS) | yes | Gallant's Wristguards (18459, +0.00 DPS, sim-verified) [dungeon]; Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (203.6 DPS) | yes | Hands of the Exalted Herald (12554, -2.39 DPS) [dungeon]; Marshal's Lamellar Gloves (231643, -2.73 DPS) [pvp]; Razor Gauntlets (18326, -12.17 DPS, sim-verified) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 133.2 healing_power points (8.06 DPS) | yes | Whipvine Cord (18327, -0.94 DPS) [dungeon]; Elunarian Belt (14465, -1.07 DPS) [world_drop]; Devout Belt (16696, -2.56 DPS, sim-verified) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | sim-verified (203.6 DPS) | yes | Red Dragonscale Leggings (252603, -0.32 DPS) [crafted]; Martyr's Legplates (250600, -0.62 DPS) [crafted]; Cloudkeeper Legplates (14554, -17.97 DPS, sim-verified) [world_drop] |
| feet | Knight-Lieutenant's Lamellar Greaves (227153) | Captain Dirgehammer [vendor] | sim-verified (203.6 DPS) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Marshal's Lamellar Boots (16472, -0.43 DPS) [vendor]; Soulforge Treads (226983, -2.03 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-verified (203.6 DPS) | yes | Seal of Rivendare (13345, -1.15 DPS) [dungeon]; Ring of Demonic Guile (18314, -1.39 DPS) [dungeon]; Naglering (11669, -14.02 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (203.6 DPS) | yes | Seal of Rivendare (13345, -0.23 DPS) [dungeon]; Ring of Demonic Guile (18314, -0.48 DPS) [dungeon]; Naglering (11669, -10.82 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (203.6 DPS) | yes | Shard of the Splithooves (10659, -3.85 DPS) [quest]; Briarwood Reed (12930, -4.66 DPS) [dungeon]; Serenity Field (272439, -7.71 DPS, sim-verified) [vendor] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-verified (203.6 DPS) | yes | Shard of the Splithooves (10659, -2.14 DPS) [quest]; Briarwood Reed (12930, -2.95 DPS) [dungeon]; Serenity Field (272439, -3.75 DPS, sim-verified) [vendor] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-verified (203.6 DPS) | yes | Hand of Righteousness (7721, -3.49 DPS) [dungeon]; Grand Marshal's Demolisher (234568, -3.67 DPS) [pvp]; Hand of Edward the Odd (2243, -7.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Soulforge Epaulets; back: Hide of the Wild; chest: Knight-Captain's Lamellar Chestplate; wrist: Soulforge Bindings; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Padre's Trousers; feet: Knight-Lieutenant's Lamellar Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60, raid preset (human, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 575.0. Weights run: 5.2s. Verify run: 49.9s. 1685 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.361, intellect=1.096 ± 0.055, spirit=0.446 ± 0.053, mp5=2.718 ± 0.069, crit=0.601 ± 0.062 per rating point (14 rating = 1%, 8.416 per %), spell_haste=not significant (0.056 ± 0.782)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 99.0 healing_power points (25.97 DPS) | yes | Lieutenant Commander's Lamellar Helmet (227149, -4.63 DPS) [vendor]; Field Marshal's Lamellar Helmet (231640, -7.71 DPS) [vendor]; Soulforge Crown (226981, -8.34 DPS) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.3 healing_power points (13.73 DPS) | yes | Drake Tooth Necklace (21531, -3.45 DPS) [quest]; Animated Chain Necklace (18723, -4.37 DPS) [dungeon]; Amulet of the Redeemed (22327, -6.27 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 79.8 healing_power points (20.93 DPS) | yes | Shimmering Dawnbringer Shoulders (227859, -0.36 DPS) [vendor]; Lieutenant Commander's Lamellar Pauldrons (227148, -5.05 DPS) [vendor]; Dawnbringer Shoulders (12625, -8.21 DPS) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.0 healing_power points (13.90 DPS) | yes | Drape of Recovery (272413, -2.91 DPS) [vendor]; Cloak of the Cosmos (18389, -3.91 DPS) [dungeon]; Caretaker's Cape (19530, -6.14 DPS) [rep] |
| chest | Knight-Captain's Lamellar Chestplate (227151) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robes of the Exalted (13346, +0.00 DPS) [dungeon]; Breastplate of Salvation (250601, -2.14 DPS) [crafted]; Field Marshal's Lamellar Chestplate (231641, -2.27 DPS) [pvp] |
| wrist | Gallant's Wristguards (18459) | Dire Maul: Guard Fengus [dungeon] | sim-verified (575.1 DPS) | yes | Loomguard Armbraces (13969, -0.96 DPS) [dungeon]; Bracers of Hope (22667, -2.35 DPS) [quest]; Bracers of Prosperity (18525, -2.40 DPS) [dungeon] |
| hands | Harmonious Gauntlets (18527) | Dire Maul: King Gordok [dungeon] | sim-verified (575.1 DPS) | yes | Raider Handwraps (272097, -1.34 DPS) [vendor]; Hands of the Exalted Herald (12554, -1.60 DPS) [dungeon]; Mooncloth Gloves (18409, -2.14 DPS) [crafted] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | sim-verified (575.1 DPS) | yes | Whipvine Cord (18327, -0.07 DPS) [dungeon]; Eyestalk Cord (18391, -1.58 DPS) [dungeon]; Belt of the Ordained (18702, -1.75 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Padre's Trousers (18386, -4.23 DPS) [dungeon]; Knight-Captain's Lamellar Legguards (227150, -4.47 DPS) [vendor]; Marshal's Lamellar Legguards (231639, -8.40 DPS) [vendor] |
| feet | Knight-Lieutenant's Lamellar Greaves (227153) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Soulforge Treads (226983, -1.87 DPS) [vendor]; Mooncloth Boots (15802, -2.78 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -4.16 DPS) [dungeon]; Fordring's Seal (16058, -4.19 DPS) [quest]; Band of Mending (22334, -4.49 DPS) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -0.33 DPS) [dungeon]; Fordring's Seal (16058, -0.35 DPS) [quest]; Band of Mending (22334, -0.65 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -2.85 DPS) [dungeon]; Briarwood Reed (12930, -3.09 DPS) [dungeon]; Second Wind (11819, -4.92 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -3.18 DPS) [dungeon]; Briarwood Reed (12930, -3.41 DPS) [dungeon]; Second Wind (11819, -5.25 DPS) [dungeon] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Diplomatic Scepter (9457, -5.10 DPS) [dungeon]; Death Speaker Scepter (2816, -7.76 DPS) [dungeon]; Inventor's Focal Sword (17719, -9.86 DPS) [dungeon] |
| off_hand | Skullflame Shield (1168) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lei of the Lifegiver (19312, +0.00 DPS) [rep]; Tome of Divine Right (22319, +0.00 DPS) [dungeon]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Knight-Captain's Lamellar Chestplate; wrist: Gallant's Wristguards; hands: Harmonious Gauntlets; waist: Sash of Mercy; legs: Martyr's Legplates; feet: Knight-Lieutenant's Lamellar Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Skullflame Shield

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 05320001000000000-0000000000000000-00000000000000000)

Set DPS (verified): 26.9. Weights run: 4.2s. Verify run: 13.7s. 219 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.899 ± 0.002, spirit=0.222 ± 0.001, mp5=1.695 ± 0.023, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.497 per %), spell_haste=0.105 ± 0.012

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Acolyte's Silvered Chain Helm (250531) | Blacksmithing [crafted] | 21.0 healing_power points (1.76 DPS) | yes | Wisdom's Leather Hood (252507, -0.34 DPS) [crafted]; Pristine Circlet (253949, -0.39 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -1.16 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.7 healing_power points (0.06 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 healing_power points (0.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.38 DPS) [crafted]; Slime-encrusted Pads (6461, -0.41 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.60 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 9.4 healing_power points (0.79 DPS) | yes | Sanguine Cape (14376, -0.49 DPS, sim-verified) [world_drop]; Pearl-clasped Cloak (5542, -0.57 DPS) [crafted]; Seer's Cape (6378, -0.60 DPS) [dungeon] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | 20.5 healing_power points (1.72 DPS) | yes | Wisdom's Leather Armor (252493, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.20 DPS) [crafted]; Bloody Apron (6226, -0.63 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 9.6 healing_power points (0.80 DPS) | yes | Owl Bracers (4796, -0.43 DPS) [vendor]; Featherbead Bracers (15452, -0.43 DPS) [quest]; Garrison Cuffs (270003, -0.93 DPS, sim-verified) [quest] |
| hands | Acolyte's Gloves (250511) | Blacksmithing [crafted] | 14.5 healing_power points (1.22 DPS) | yes | Wisdom's Leather Gloves (252499, -0.08 DPS) [crafted]; Pristine Gloves (253913, -0.19 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.59 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 14.6 healing_power points (1.22 DPS) | yes | Wisdom's Leather Belt (252433, -0.25 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.26 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.44 DPS) [crafted] |
| legs | Acolyte's Chain Leggings (250496) | Blacksmithing [crafted] | 27.3 healing_power points (2.29 DPS) | yes | Wisdom's Leather Pants (252503, -0.14 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.25 DPS) [crafted]; Ghastly Trousers (15449, -1.63 DPS) [quest] |
| feet | Acolyte's Boots (250506) (or Wisdom's Leather Boots (252444)) | Blacksmithing [crafted] | 15.5 healing_power points (1.30 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Glowing Copper Boots (250482, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.24 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 healing_power points (0.45 DPS) | yes | Advisor's Ring (20426, -0.17 DPS) [rep]; Black Pearl Ring (6332, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.23 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 4.5 healing_power points (0.38 DPS) | yes | Advisor's Ring (20426, +0.00 DPS, sim-verified) [rep]; Black Pearl Ring (6332, -0.11 DPS) [world]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Sword of the Fallen (286977) | Ruins of Lordaeron: Ragged Ghoul [dungeon] | 8.1 healing_power points (0.68 DPS) | yes | Impaling Harpoon (5200, -0.54 DPS, sim-verified) [dungeon]; Samophlange Screwdriver (11854, -0.57 DPS) [quest]; Mug of Muddled Memories (277247, -0.58 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Acolyte's Silvered Chain Helm; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Acolyte's Gloves; waist: Pristine Sash; legs: Acolyte's Chain Leggings; feet: Acolyte's Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Sword of the Fallen

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 05320003224000000-0000000000000000-00000000000000000)

Set DPS (verified): 49.3. Weights run: 4.1s. Verify run: 14.7s. 409 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=0.977 ± 0.003, spirit=0.493 ± 0.003, mp5=2.099 ± 0.007, crit=0.093 ± 0.004 per rating point (14 rating = 1%, 1.295 per %), spell_haste=0.041 ± 0.009

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 36.0 healing_power points (3.37 DPS) | yes | Acolyte's Chain Helm (250501, -1.07 DPS, sim-verified) [crafted]; Filigreed Pristine Circlet (253975, -1.31 DPS) [crafted]; Embalmed Shroud (7691, -1.32 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystal Starfire Medallion (5003, -0.04 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.09 DPS) [vendor]; Kaleidoscope Chain (13084, -0.50 DPS, sim-verified) [world_drop] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mantle of Woe (7750, -0.14 DPS) [quest]; Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Ghostly Mantle (3324, -1.32 DPS, sim-verified) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | 15.0 healing_power points (1.40 DPS) | yes | Glowing Thresher Cape (6901, -0.46 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272078, -0.53 DPS) [vendor]; Cloak of Rot (4462, -0.67 DPS) [world] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 30.7 healing_power points (2.88 DPS) | yes | Pristine Gown (253961, -0.18 DPS) [crafted]; Wisdom's Leather Tunic (252511, -0.37 DPS) [crafted]; Acolyte's Silvered Chain Shirt (250521, -0.51 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 11.2 healing_power points (1.05 DPS) | yes | Nightsky Wristbands (6407, -0.36 DPS) [world_drop]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.60 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 17.9 healing_power points (1.68 DPS) | yes | Acolyte's Gloves (250511, -0.29 DPS) [crafted]; Pristine Gloves (253913, -0.38 DPS) [crafted]; Naga Battle Gloves (888, -0.65 DPS, sim-verified) [dungeon] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 28.8 healing_power points (2.70 DPS) | yes | Mender's Leather Belt (252523, +0.00 DPS, sim-verified) [crafted]; Pristine Sash (253925, -1.31 DPS) [crafted]; Acolyte's Chain Belt (250516, -1.49 DPS) [crafted] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 32.8 healing_power points (3.08 DPS) | yes | Pristine Leggings (253987, -0.08 DPS, sim-verified) [crafted]; Wisdom's Leather Leggings (252519, -0.14 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.47 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 28.8 healing_power points (2.70 DPS) | yes | Wisdom's Leather Boots (252444, -1.21 DPS) [crafted]; Glowing Copper Boots (250482, -1.49 DPS) [crafted]; Acolyte's Boots (250506, -1.50 DPS, sim-verified) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.69 DPS) | yes | Electrocutioner Lagnut (9447, -0.80 DPS) [dungeon]; Darkspear Signet (272071, -0.90 DPS) [vendor]; Black Widow Band (6199, -1.05 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 healing_power points (0.96 DPS) | yes | Darkspear Signet (272071, -0.18 DPS) [vendor]; Black Widow Band (6199, -0.32 DPS) [world]; Electrocutioner Lagnut (9447, -0.61 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Royal Diplomatic Scepter (9457) | Gnomeregan: Dark Iron Ambassador [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Death Speaker Scepter (2816, -0.98 DPS) [dungeon]; Manual Crowd Pummeler (9449, -1.54 DPS, sim-verified) [dungeon]; Haunting Blade (6641, -4.26 DPS) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 16.0 healing_power points (1.50 DPS) | yes | Alliance Outrunner Healing Rod (285348, -0.11 DPS, sim-verified) [world]; Orb of Mistmantle (13031, -0.47 DPS) [world_drop]; Heart of Agamaggan (6694, -0.56 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Battle Healer's Cloak; chest: Death Speaker Robes; hands: Truefaith Gloves; waist: Prefect's Belt; legs: Acolyte's Silvered Chain Leggings; feet: Gilded Slippers; finger1: Sea Giant's Toe Ring; finger2: Snake Hoop; main_hand: Royal Diplomatic Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 409, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 05320003225111051-0000000000000000-00000000000000000)

Set DPS (verified): 93.6. Weights run: 7.4s. Verify run: 29.4s. 558 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=1.162 ± 0.006, spirit=0.735 ± 0.005, mp5=2.690 ± 0.008, crit=0.150 ± 0.006 per rating point (14 rating = 1%, 2.097 per %), spell_haste=not significant (0.060 ± 0.017)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 48.4 healing_power points (4.38 DPS) | yes | Earthen Silk Hood (254015, -1.72 DPS) [crafted]; Miner's Hat of the Deep (9429, -1.93 DPS) [dungeon]; Holy Shroud (2721, -3.65 DPS, sim-verified) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.3 healing_power points (2.02 DPS) | yes | Necklace of Calisea (1714, -0.82 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.97 DPS) [quest]; Triune Amulet (7722, -1.15 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 30.1 healing_power points (2.73 DPS) | yes | Sheepshear Mantle (13115, -0.64 DPS) [world_drop]; Mistscape Mantle (4734, -1.24 DPS) [dungeon]; Earthen Silk Shoulders (254033, -1.62 DPS, sim-verified) [crafted] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | 21.7 healing_power points (1.96 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.14 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272077, -0.54 DPS) [vendor]; Glowing Thresher Cape (6901, -0.80 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 63.9 healing_power points (5.78 DPS) | yes | Death Speaker Robes (6682, -2.82 DPS) [dungeon]; Doomsayer's Robe (4746, -2.85 DPS) [quest]; Deathchill Armor (10764, -3.76 DPS, sim-verified) [dungeon] |
| wrist | Reflective Wristguards (274750) | Rettrick [vendor] | 18.3 healing_power points (1.66 DPS) | yes | Mindthrust Bracers (1974, -0.41 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.42 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.57 DPS) [world_drop] |
| hands | Mender's Leather Gloves (252530) | Leatherworking [crafted] | 31.5 healing_power points (2.85 DPS) | yes | Prefect's Gauntlet (250569, -0.21 DPS) [crafted]; Gloves of the Greatfather (17721, -0.49 DPS) [crafted]; Gilded Handwraps (254021, -0.54 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 35.7 healing_power points (3.23 DPS) | yes | Prefect's Belt (250559, -0.50 DPS) [crafted]; Mender's Leather Belt (252523, -1.15 DPS, sim-verified) [crafted]; Highlander's Mail Girdle (20119, -1.65 DPS) [vendor] |
| legs | Acolyte's Silvered Chain Leggings (250526) | Blacksmithing [crafted] | 34.3 healing_power points (3.11 DPS) | yes | Wisdom's Leather Leggings (252519, +0.00 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -0.04 DPS) [crafted]; Stoneweaver Leggings (9407, -0.29 DPS) [dungeon] |
| feet | Prefect's Boots (250549) | Blacksmithing [crafted] | 38.5 healing_power points (3.48 DPS) | yes | Mender's Leather Shoes (252533, -0.11 DPS) [crafted]; Mender's Mail Boots (252565, -0.11 DPS) [crafted]; Gilded Slippers (254001, -0.67 DPS) [crafted] |
| finger1 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 18.0 healing_power points (1.63 DPS) | yes | Snake Hoop (6750, -0.43 DPS) [quest]; Electrocutioner Lagnut (9447, -0.62 DPS) [dungeon]; Welken Ring (5011, -0.64 DPS) [world_drop] |
| finger2 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 16.1 healing_power points (1.46 DPS) | yes | Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon]; Welken Ring (5011, -0.47 DPS) [world_drop]; Snake Hoop (6750, -2.05 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (93.6 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (93.6 DPS) | yes | Royal Diplomatic Scepter (9457, -1.79 DPS) [dungeon]; Death Speaker Scepter (2816, -2.89 DPS) [dungeon]; Frost Tiger Blade (3854, -18.20 DPS, sim-verified) [crafted] |
| off_hand | Forcestone Buckler (17508) | Compendium of the Fallen [quest] | 17.9 healing_power points (1.62 DPS) | yes | Beacon of Hope (9393, -0.09 DPS) [dungeon]; Prophetic Cane (6803, -0.36 DPS) [quest]; Orb of Souls (249395, -16.25 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; wrist: Reflective Wristguards; hands: Mender's Leather Gloves; waist: Gilded Cord; feet: Prefect's Boots; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Forcestone Buckler

No-known-source sample (15 of 558, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 05320003225111051-5500000000000000-00000000000000000)

Set DPS (verified): 118.4. Weights run: 7.8s. Verify run: 38.2s. 733 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.004, intellect=1.850 ± 0.013, spirit=1.414 ± 0.011, mp5=4.678 ± 0.015, crit=0.311 ± 0.015 per rating point (14 rating = 1%, 4.352 per %), spell_haste=not significant (0.102 ± 0.042)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (118.4 DPS) | yes | Gemburst Circlet (10751, -0.02 DPS) [quest]; Soulcatcher Halo (10630, -0.41 DPS) [dungeon]; Braincage (12549, -0.52 DPS) [dungeon] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 39.2 healing_power points (2.81 DPS) | yes | Horizon Choker (13085, -0.55 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -0.81 DPS) [dungeon]; Gemshard Heart (17707, -0.87 DPS) [dungeon] |
| shoulder | Mender's Leather Shoulder (252538) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lead Surveyor's Mantle (11842, +0.00 DPS) [dungeon]; Mender's Mail Shoulder (252569, +0.00 DPS) [crafted]; Living Shoulders (15061, -0.08 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 33.0 healing_power points (2.36 DPS) | yes | Battle Healer's Cloak (19527, -0.18 DPS) [rep]; Featherskin Cape (10843, -0.31 DPS) [world]; Imperial Red Cloak (8248, -0.50 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (118.4 DPS) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Robes of Insight (940, -0.13 DPS) [world_drop]; Ghostweave Vest (14141, -0.26 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 40.4 healing_power points (2.90 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Prefect's Wristguards (250584, -0.19 DPS) [crafted] |
| hands | Soulforge Fists (226982) | Mokvar [vendor] | 69.8 healing_power points (5.00 DPS) | yes | Raider Handwraps (272098, -0.51 DPS) [vendor]; Mender's Leather Gauntlets (252551, -0.68 DPS) [crafted]; Mender's Mail Gauntlets (252587, -0.68 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 52.2 healing_power points (3.74 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Prefect's Waistguard (250574, -0.27 DPS) [crafted]; Dawnspire Cord (12466, -0.36 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 64.8 healing_power points (4.64 DPS) | yes | Senior Designer's Pantaloons (11841, -0.10 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.24 DPS) [dungeon]; Dalewind Trousers (13008, -0.91 DPS) [world_drop] |
| feet | Mender's Leather Boots (252472) (or Mender's Mail Sabatons (252579)) | Leatherworking [crafted] | 48.4 healing_power points (3.46 DPS) | yes | Mender's Mail Sabatons (252579, +0.00 DPS) [crafted]; Gilded Sandals (254107, -0.05 DPS) [crafted]; Mender's Mail Boots (252565, -0.11 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.8 healing_power points (2.49 DPS) | yes | Eye of Adaegus (5266, -0.48 DPS) [world_drop]; Cyclopean Band (11824, -0.52 DPS) [dungeon]; Snake Hoop (6750, -0.86 DPS) [quest] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 32.7 healing_power points (2.35 DPS) | yes | Eye of Adaegus (5266, -0.34 DPS) [world_drop]; Cyclopean Band (11824, -0.37 DPS) [dungeon]; Snake Hoop (6750, -0.71 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -3.81 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.01 DPS) [quest]; Alchemist's Stone (13503, -4.62 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.07 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.28 DPS) [quest]; Alchemist's Stone (13503, -0.88 DPS) [crafted] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Diplomatic Scepter (9457, -1.46 DPS) [dungeon]; Death Speaker Scepter (2816, -2.67 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.25 DPS) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | 37.2 healing_power points (2.67 DPS) | yes | Enthralled Sphere (11625, -0.30 DPS) [dungeon]; Cloud Stone (17737, -0.33 DPS) [dungeon]; Twisting Essence Jar (249456, -0.49 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** neck: Darkmoon Necklace; shoulder: Mender's Leather Shoulder; back: Darkspear Raider's Cloak; wrist: Mender's Leather Bracers; hands: Soulforge Fists; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Mender's Leather Boots; finger1: Brainlash; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 733, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 192.8. Weights run: 9.0s. Verify run: 117.5s. 1710 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.126, intellect=4.925 ± 0.058, spirit=1.708 ± 0.014, mp5=7.069 ± 0.046, crit=0.442 ± 0.030 per rating point (14 rating = 1%, 6.193 per %), spell_haste=not significant (0.104 ± 0.054)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 216.3 healing_power points (13.08 DPS) | yes | Crown of the Penitent (13216, -1.80 DPS) [quest]; Living Crown (252561, -1.86 DPS) [crafted]; Black Dragonscale Helm (252605, -3.19 DPS, sim-verified) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, -0.26 DPS) [world_drop]; Tooth of Gnarr (13141, -1.21 DPS) [dungeon]; Lady Maye's Pendant (14558, -7.29 DPS, sim-verified) [world_drop] |
| shoulder | Soulforge Epaulets (226979) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Shoulders (272104, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.38 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 91.2 healing_power points (5.52 DPS) | yes | Shroud of the Exile (15421, -0.33 DPS) [quest]; Faded Hakkari Cloak (20218, -0.57 DPS) [quest]; Darkspear Raider's Cloak (272063, -5.76 DPS, sim-verified) [vendor] |
| chest | Breastplate of Salvation (250601) | Blacksmithing [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Devout Robe (16690, -0.09 DPS) [dungeon]; Mooncloth Vest (14138, -0.51 DPS) [crafted]; Soulforge Embrace (226984, -2.26 DPS, sim-verified) [vendor] |
| wrist | Soulforge Bindings (226977) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gallant's Wristguards (18459, +0.00 DPS, sim-verified) [dungeon]; Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, -2.39 DPS) [dungeon]; Mooncloth Gloves (18409, -2.73 DPS) [crafted]; Razor Gauntlets (18326, -11.80 DPS, sim-verified) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 133.2 healing_power points (8.06 DPS) | yes | Whipvine Cord (18327, -0.94 DPS) [dungeon]; Elunarian Belt (14465, -1.07 DPS) [world_drop]; Devout Belt (16696, -2.52 DPS, sim-verified) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Dragonscale Leggings (252603, -0.32 DPS) [crafted]; Martyr's Legplates (250600, -0.62 DPS) [crafted]; Cloudkeeper Legplates (14554, -17.03 DPS, sim-verified) [world_drop] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Boots (15802, -1.31 DPS) [crafted]; Faith Healer's Boots (22247, -1.89 DPS) [dungeon]; Soulforge Treads (226983, -2.00 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -1.15 DPS) [dungeon]; Ring of Demonic Guile (18314, -1.39 DPS) [dungeon]; Naglering (11669, -13.84 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.23 DPS) [dungeon]; Ring of Demonic Guile (18314, -0.48 DPS) [dungeon]; Naglering (11669, -10.80 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -3.85 DPS) [quest]; Briarwood Reed (12930, -4.66 DPS) [dungeon]; Serenity Field (272439, -6.75 DPS, sim-verified) [vendor] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -2.14 DPS) [quest]; Serenity Field (272439, -2.80 DPS, sim-verified) [vendor]; Briarwood Reed (12930, -2.95 DPS) [dungeon] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Righteousness (7721, -3.49 DPS) [dungeon]; High Warlord's Destroyer (234546, -3.67 DPS) [pvp]; Hand of Edward the Odd (2243, -7.68 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Soulforge Epaulets; back: Hide of the Wild; chest: Breastplate of Salvation; wrist: Soulforge Bindings; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1710, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 05320003225111051-5532500000000000-00000000000000000)

Set DPS (verified): 560.3. Weights run: 5.2s. Verify run: 40.2s. 1710 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.361, intellect=1.096 ± 0.055, spirit=0.446 ± 0.053, mp5=2.718 ± 0.069, crit=0.601 ± 0.062 per rating point (14 rating = 1%, 8.416 per %), spell_haste=not significant (0.056 ± 0.782)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 99.0 healing_power points (25.97 DPS) | yes | Soulforge Crown (226981, -8.34 DPS) [vendor]; Sanctified Leather Helm (22689, -8.60 DPS) [quest]; Insightful Hood (18490, -9.73 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 52.3 healing_power points (13.73 DPS) | yes | Drake Tooth Necklace (21531, -3.45 DPS) [quest]; Animated Chain Necklace (18723, -4.37 DPS) [dungeon]; Amulet of the Redeemed (22327, -6.27 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 79.8 healing_power points (20.93 DPS) | yes | Shimmering Dawnbringer Shoulders (227859, -0.36 DPS) [vendor]; Dawnbringer Shoulders (12625, -8.21 DPS) [crafted]; Darkspear Mantle (272107, -9.01 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.0 healing_power points (13.90 DPS) | yes | Drape of Recovery (272413, -2.91 DPS) [vendor]; Cloak of the Cosmos (18389, -3.91 DPS) [dungeon]; Battle Healer's Cloak (19526, -6.14 DPS) [rep] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | sim-verified (560.2 DPS) | yes | Breastplate of Salvation (250601, -2.19 DPS) [crafted]; Stormcloth Vest (10020, -4.41 DPS) [crafted]; Soulforge Embrace (226984, -4.50 DPS) [vendor] |
| wrist | Gallant's Wristguards (18459) | Dire Maul: Guard Fengus [dungeon] | sim-verified (560.2 DPS) | yes | Loomguard Armbraces (13969, -0.96 DPS) [dungeon]; Bracers of Hope (22667, -2.35 DPS) [quest]; Bracers of Prosperity (18525, -2.40 DPS) [dungeon] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Harmonious Gauntlets (18527, +0.00 DPS) [dungeon]; Hands of the Exalted Herald (12554, -0.26 DPS) [dungeon]; Mooncloth Gloves (18409, -0.80 DPS) [crafted] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | sim-verified (560.2 DPS) | yes | Whipvine Cord (18327, -0.07 DPS) [dungeon]; Eyestalk Cord (18391, -1.58 DPS) [dungeon]; Belt of the Ordained (18702, -1.75 DPS) [dungeon] |
| legs | Martyr's Legplates (250600) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Padre's Trousers (18386, -4.23 DPS) [dungeon]; Soulforge Leggings (226980, -10.44 DPS) [vendor]; Ghoul Skin Leggings (18682, -10.50 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 65.6 healing_power points (17.21 DPS) | yes | Soulforge Treads (226983, -4.46 DPS) [vendor]; Mooncloth Boots (15802, -5.36 DPS) [crafted]; Faith Healer's Boots (22247, -5.53 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -4.16 DPS) [dungeon]; Fordring's Seal (16058, -4.19 DPS) [quest]; Band of Mending (22334, -4.49 DPS) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, -0.33 DPS) [dungeon]; Fordring's Seal (16058, -0.35 DPS) [quest]; Band of Mending (22334, -0.65 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -2.85 DPS) [dungeon]; Briarwood Reed (12930, -3.09 DPS) [dungeon]; Second Wind (11819, -4.92 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -3.18 DPS) [dungeon]; Briarwood Reed (12930, -3.41 DPS) [dungeon]; Second Wind (11819, -5.25 DPS) [dungeon] |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Diplomatic Scepter (9457, -5.10 DPS) [dungeon]; Death Speaker Scepter (2816, -7.76 DPS) [dungeon]; Inventor's Focal Sword (17719, -9.86 DPS) [dungeon] |
| off_hand | Skullflame Shield (1168) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lei of the Lifegiver (19312, +0.00 DPS) [rep]; Tome of Divine Right (22319, +0.00 DPS) [dungeon]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Gallant's Wristguards; hands: Raider Handwraps; waist: Sash of Mercy; legs: Martyr's Legplates; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Skullflame Shield

No-known-source sample (15 of 1710, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

