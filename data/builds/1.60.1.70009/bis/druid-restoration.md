# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 47.9. Weights run: 5.3s. Verify run: 18.2s. 193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.281 ± 0.004, spirit=1.542 ± 0.008, mp5=3.288 ± 0.015, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.529 per %), spell_haste=-5.534 ± 0.104

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.24 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 6.2 healing_power points (0.63 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 healing_power points (1.18 DPS) | yes | Forest Leather Mantle (4709, -0.55 DPS) [world_drop]; Prospector's Pads (14566, -0.55 DPS) [world_drop]; Slime-encrusted Pads (6461, -0.80 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 12.1 healing_power points (1.24 DPS) | yes | Spirit Cloak (4792, -0.61 DPS) [vendor]; Sylvan Cloak (4793, -0.61 DPS) [vendor]; Regent's Cloak (5969, -2.19 DPS, sim-verified) [world] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.0 healing_power points (2.67 DPS) | yes | Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon]; Robe of the Moccasin (6465, -2.21 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 16.3 healing_power points (1.67 DPS) | yes | Owl Bracers (4796, -1.01 DPS) [vendor]; Bravo's Armbands (270015, -1.04 DPS) [quest]; Drakewing Bands (12999, -2.53 DPS, sim-verified) [world_drop] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 15.1 healing_power points (1.55 DPS) | yes | Magefist Gloves (12977, -0.26 DPS) [world_drop]; Bright Gloves (3066, -0.39 DPS) [world_drop]; Pristine Gloves (253913, -1.51 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.1 healing_power points (1.66 DPS) | yes | Wisdom's Leather Belt (252433, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.60 DPS) [world_drop]; Novice Ardent's Sash (253887, -1.72 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.26 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.79 DPS) | yes | Pristine Boots (253889, -0.47 DPS) [crafted]; Kimbra Boots (6191, -0.79 DPS) [quest]; Black Whelp Slippers (252424, -1.89 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.8 healing_power points (1.21 DPS) | yes | Deep Fathom Ring (6463, -0.42 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.54 DPS) [rep] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.95 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.27 DPS) [rep]; Deep Fathom Ring (6463, -2.09 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 15.7 healing_power points (1.61 DPS) | yes | Twisted Chanter's Staff (890, -0.29 DPS) [world_drop]; Gnarled Hermit's Staff (1539, -0.50 DPS) [world]; Staff of the Blessed Seer (2271, -1.57 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Band of Purification; main_hand: Staff of Westfall

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 89.2. Weights run: 7.8s. Verify run: 26.9s. 322 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.448 ± 0.014, spirit=2.905 ± 0.018, mp5=4.429 ± 0.027, crit=0.046 ± 0.004 per rating point (14 rating = 1%, 0.641 per %), spell_haste=-1.040 ± 0.256

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Whisperwind Headdress (6688) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (89.2 DPS) | yes | Enduring Cap (3020, -0.44 DPS) [world_drop]; Embalmed Shroud (7691, -1.32 DPS) [dungeon]; Holy Shroud (2721, -2.53 DPS, sim-verified) [world_drop] |
| neck | Necklace of Harmony (5180) | Singer [world] | 20.3 healing_power points (2.06 DPS) | yes | Glowing Green Talisman (5002, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.30 DPS) [world_drop]; Pendant of Myzrael (4614, -0.40 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 30.5 healing_power points (3.09 DPS) | yes | Batwing Mantle (6697, +0.00 DPS, sim-verified) [dungeon]; Faerie Mantle (5820, -0.44 DPS) [quest]; Nightsky Mantle (4718, -0.44 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.2 healing_power points (3.07 DPS) | yes | Repairman's Cape (9605, -0.42 DPS) [quest]; Caretaker's Cape (19533, -0.57 DPS) [rep]; Prelacy Cape (7004, -0.66 DPS, sim-verified) [quest] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.1 healing_power points (4.48 DPS) | yes | Pristine Gown (253961, -0.23 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.36 DPS) [dungeon]; Pressed Felt Robe (1997, -0.50 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.5 healing_power points (2.09 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop]; Drakewing Bands (12999, -0.75 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.1 healing_power points (4.28 DPS) | yes | Zodiac Gloves (7106, -1.62 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.62 DPS) [dungeon]; Shilly Mitts (9609, -1.69 DPS, sim-verified) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.1 healing_power points (4.58 DPS) | yes | Dokebi Cord (14578, -2.22 DPS) [world_drop]; Resilient Cord (14406, -2.52 DPS) [world_drop]; Silver-lined Belt (13011, -3.06 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.5 healing_power points (5.23 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted]; Earthen Leggings (253999, -1.06 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.24 DPS) | yes | Boots of the Enchanter (4325, -1.88 DPS) [crafted]; Stonecloth Boots (14408, -1.88 DPS) [world_drop]; Soggy Boots (274747, -2.66 DPS, sim-verified) [vendor] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.5 healing_power points (3.09 DPS) | yes | Monkey Ring (6748, -1.03 DPS) [quest]; Ring of Calm (6790, -1.03 DPS) [quest]; Electrocutioner Lagnut (9447, -1.03 DPS) [dungeon] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 26.1 healing_power points (2.65 DPS) | yes | Ring of Calm (6790, -0.56 DPS, sim-verified) [quest]; Monkey Ring (6748, -0.59 DPS) [quest]; Electrocutioner Lagnut (9447, -0.59 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wind Spirit Staff (6689, -1.75 DPS) [dungeon]; Gnarled Ash Staff (791, -2.48 DPS) [world_drop]; Manual Crowd Pummeler (9449, -3.52 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.4 healing_power points (3.09 DPS) | yes | Orb of Mistmantle (13031, -1.00 DPS) [world_drop]; Strength of Will (4837, -1.32 DPS) [vendor]; Defective Samophlange (274743, -1.41 DPS, sim-verified) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Whisperwind Headdress; neck: Necklace of Harmony; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 136.9. Weights run: 8.7s. Verify run: 54.2s. 438 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.339 ± 0.016, spirit=2.500 ± 0.011, mp5=4.838 ± 0.034, crit=0.050 ± 0.005 per rating point (14 rating = 1%, 0.694 per %), spell_haste=not significant (0.513 ± 0.312)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 67.3 healing_power points (6.90 DPS) | yes | Whitemane's Chapeau (7720, -1.25 DPS, sim-verified) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.76 DPS) [dungeon]; Holy Shroud (2721, -1.98 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 34.2 healing_power points (3.51 DPS) | yes | Triune Amulet (7722, -0.72 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.75 DPS) [dungeon]; Amberglow Talisman (10824, -0.94 DPS) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.0 healing_power points (3.90 DPS) | yes | Crimson Silk Shoulders (7059, -1.00 DPS) [crafted]; Mantle of Doan (7712, -1.00 DPS) [dungeon]; Sheepshear Mantle (13115, -1.25 DPS, sim-verified) [world_drop] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | 30.5 healing_power points (3.13 DPS) | yes | Prelacy Cape (7004, -0.46 DPS) [quest]; Ceremonial Centaur Blanket (6789, -0.55 DPS) [quest]; Glowing Thresher Cape (6901, -0.89 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.0 healing_power points (8.00 DPS) | yes | Dreamweave Vest (10021, -3.18 DPS) [crafted]; Black Mageweave Robe (10001, -3.31 DPS) [crafted]; Doomsayer's Robe (4746, -6.85 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.4 healing_power points (3.12 DPS) | yes | Earthen Silk Cuffs (254019, -0.50 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Silkstream Cuffs (16791, -1.40 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (136.9 DPS) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Mender's Leather Gloves (252530, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, -0.75 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.7 healing_power points (4.90 DPS) | yes | Mender's Leather Belt (252523, -0.62 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (136.9 DPS) | yes | Wisdom's Leather Leggings (252519, -0.47 DPS) [crafted]; Pristine Leggings (253987, -0.89 DPS) [crafted]; Warchief Kilt (7760, -3.18 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 50.9 healing_power points (5.22 DPS) | yes | Furen's Boots (13100, -0.92 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.06 DPS) [dungeon]; Gilded Slippers (254001, -1.18 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.0 healing_power points (2.98 DPS) | yes | Welken Ring (5011, -0.50 DPS) [world_drop]; The Queen's Jewel (13094, -0.65 DPS) [world_drop]; Blush Ember Ring (13093, -0.93 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 26.9 healing_power points (2.76 DPS) | yes | Welken Ring (5011, -0.15 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.43 DPS) [world_drop]; Blush Ember Ring (13093, -0.71 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (136.9 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (136.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -2.44 DPS) [dungeon]; Staff of Jordan (873, -2.64 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 32.0 healing_power points (3.28 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.46 DPS) [dungeon]; Orb of Souls (249395, -0.89 DPS, sim-verified) [crafted]; Silksand Star (15964, -0.97 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; off_hand: Beacon of Hope

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 199.7. Weights run: 8.6s. Verify run: 35.0s. 577 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.522 ± 0.017, spirit=2.750 ± 0.016, mp5=4.393 ± 0.016, crit=0.069 ± 0.007 per rating point (14 rating = 1%, 0.967 per %), spell_haste=0.966 ± 0.177

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | sim-verified (199.7 DPS) | yes | Papal Fez (9431, -0.27 DPS) [dungeon]; Cassandra's Grace (13102, -0.66 DPS) [world_drop]; Knight-Lieutenant's Restored Leather Helm (220874, -2.38 DPS, sim-verified) [vendor] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 41.2 healing_power points (4.09 DPS) | yes | Glowing Eye of Mordresh (10769, -0.37 DPS, sim-verified) [dungeon]; Darkmoon Necklace (19303, -0.57 DPS) [vendor]; Horizon Choker (13085, -0.89 DPS) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 66.7 healing_power points (6.62 DPS) | yes | Mender's Leather Shoulder (252538, -0.97 DPS) [crafted]; Nethergeld Shoulders (254049, -1.20 DPS) [crafted]; Knight-Lieutenant's Restored Leather Spaulders (220876, -4.42 DPS, sim-verified) [vendor] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | 47.3 healing_power points (4.69 DPS) | yes | Caretaker's Cape (19531, +0.00 DPS, sim-verified) [rep]; Darkspear Raider's Cloak (272076, -1.22 DPS) [vendor]; Arcane Cloak (8286, -1.42 DPS) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 108.4 healing_power points (10.74 DPS) | yes | Ghostweave Vest (14141, -1.81 DPS) [crafted]; Vestments of the Atal'ai Prophet (10806, -2.10 DPS, sim-verified) [dungeon]; Forest's Embrace (22272, -2.17 DPS) [quest] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.57 DPS) | yes | Nethergeld Cuffs (254061, -0.49 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -0.68 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.2 healing_power points (8.94 DPS) | yes | Mender's Leather Gauntlets (252551, +0.00 DPS, sim-verified) [crafted]; Gilded Gloves (254095, -2.38 DPS) [crafted]; Earthenweave Gloves (254075, -3.37 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.2 healing_power points (5.58 DPS) | yes | Mender's Leather Waistguard (252477, -0.79 DPS) [crafted]; Gilded Cord (254037, -0.95 DPS, sim-verified) [crafted]; Mender's Leather Belt (252523, -1.15 DPS) [crafted] |
| legs | Dalewind Trousers (13008) | World drop [world_drop] | 82.4 healing_power points (8.17 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.58 DPS, sim-verified) [dungeon]; Rainstrider Leggings (11123, -0.81 DPS) [quest]; Windscale Sarong (10842, -1.21 DPS) [world] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.2 healing_power points (6.66 DPS) | yes | Mistwalker Boots (10629, -0.55 DPS) [dungeon]; Earthenweave Boots (254093, -1.08 DPS) [crafted]; Sergeant Major's Restored Leather Boots (220884, -2.70 DPS, sim-verified) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.1 healing_power points (4.18 DPS) | yes | Choking Band (11868, -0.63 DPS) [quest]; Darkspear Signet (272069, -1.13 DPS) [vendor]; Cyclopean Band (11824, -1.14 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 36.6 healing_power points (3.63 DPS) | yes | Choking Band (11868, +0.00 DPS, sim-verified) [quest]; Darkspear Signet (272069, -0.58 DPS) [vendor]; Cyclopean Band (11824, -0.59 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -3.81 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.35 DPS) [quest]; Uther's Strength (11302, -4.75 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.92 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -1.09 DPS) [quest]; Uther's Strength (11302, -1.49 DPS) [world_drop] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -1.92 DPS) [world_drop]; Death Speaker Scepter (2816, -2.01 DPS) [dungeon]; Barman Shanker (12791, -9.27 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Lei of Lilies; shoulder: Living Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Dalewind Trousers; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Soulkeeper

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 324.6. Weights run: 9.0s. Verify run: 72.4s. 1449 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 173.1 healing_power points (17.46 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -2.67 DPS) [vendor]; Wildheart Cowl (16720, -2.81 DPS) [dungeon]; Feralheart Headdress (226786, -3.26 DPS) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 92.3 healing_power points (9.31 DPS) | yes | Lady Maye's Pendant (14558, -0.72 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.94 DPS) [world_drop]; Heart of the Fiend (13960, -1.72 DPS) [dungeon] |
| shoulder | Feralheart Mantle (226785) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Argent Elite Shoulders (227888, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 78.1 healing_power points (7.88 DPS) | yes | Featherskin Cape (10843, -0.51 DPS) [world]; Butcher's Apron (12608, -0.98 DPS) [dungeon]; Hide of the Wild (18510, -1.39 DPS) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, -1.45 DPS) [crafted]; Alanna's Embrace (13314, -1.62 DPS) [dungeon]; Ironfeather Breastplate (15066, -1.99 DPS) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, -0.02 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.12 DPS) [dungeon]; Marshal's Dragonhide Bracers (16445, -0.88 DPS) [pvp] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 125.5 healing_power points (12.66 DPS) | yes | Wildheart Gloves (16717, -0.47 DPS) [dungeon]; Hands of the Exalted Herald (12554, -1.23 DPS) [dungeon]; Devout Gloves (16692, -1.56 DPS) [dungeon] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | sim-verified (324.6 DPS) | yes | Caretaker's Cord (272398, +0.00 DPS) [vendor]; Elderwild Waistcord (279252, -0.72 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.34 DPS) [crafted] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 191.0 healing_power points (19.26 DPS) | yes | Devout Skirt (16694, -3.54 DPS) [dungeon]; Haunting Specter Leggings (11929, -4.48 DPS) [dungeon]; Knight-Captain's Dragonhide Legguards (227200, -4.55 DPS) [vendor] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 148.3 healing_power points (14.96 DPS) | yes | Incandescent Mooncloth Boots (227862, -1.33 DPS) [vendor]; Mooncloth Boots (15802, -3.78 DPS) [crafted]; Devout Sandals (16691, -4.26 DPS) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.72 DPS) [dungeon]; Band of Piety (22681, -0.77 DPS) [quest]; Band of the Hierophant (13096, -1.01 DPS) [world_drop] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.38 DPS) [dungeon]; Band of Piety (22681, -0.43 DPS) [quest]; Band of the Hierophant (13096, -0.67 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -2.77 DPS) [dungeon]; Ankh of Life (1713, -5.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.07 DPS) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, -0.12 DPS) [quest] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -0.09 DPS) [world_drop]; Soulkeeper (1607, -1.52 DPS) [world_drop]; Quel'dorai Channeling Rod (18311, -2.17 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Feralheart Mantle; back: Frostweaver Cape; wrist: Bracers of Hope; waist: Feralheart Cord; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: The Postmaster's Seal; finger2: Emerald Flame Ring; trinket2: Royal Seal of Eldre'Thalas; main_hand: Dancing Sliver

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 599.1. Weights run: 6.5s. Verify run: 48.7s. 1449 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.204, intellect=2.451 ± 0.080, spirit=2.281 ± 0.075, mp5=3.336 ± 0.099, crit=0.327 ± 0.030 per rating point (14 rating = 1%, 4.573 per %), spell_haste=not significant (-0.743 ± 1.216)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 149.0 healing_power points (24.64 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -4.67 DPS) [vendor]; Sanctified Leather Helm (22689, -5.80 DPS) [quest]; Feralheart Headdress (226786, -16.56 DPS, sim-verified) [vendor] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 69.4 healing_power points (11.47 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.41 DPS) [world_drop]; Wavefront Necklace (20685, -0.66 DPS) [world]; Drake Tooth Necklace (21531, -2.31 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 101.8 healing_power points (16.83 DPS) | yes | Devout Mantle (16695, +0.00 DPS, sim-verified) [dungeon]; Lieutenant Commander's Dragonhide Pauldrons (227201, -1.16 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -1.32 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 66.5 healing_power points (10.99 DPS) | yes | Cloak of the Cosmos (18389, -2.24 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -2.25 DPS) [vendor]; Frostweaver Cape (12968, -4.83 DPS, sim-verified) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-verified (599.2 DPS) | yes | Alanna's Embrace (13314, -0.61 DPS) [dungeon]; Devout Robe (16690, -0.88 DPS) [dungeon]; Tunic of Undead Slaying (23089, -19.17 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (599.2 DPS) | yes | Bracers of Mending (23129, -0.43 DPS) [dungeon]; Marshal's Dragonhide Bracers (16445, -0.49 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -11.38 DPS, sim-verified) [world] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 92.2 healing_power points (15.25 DPS) | yes | Raider Handwraps (272097, -0.34 DPS) [vendor]; Wildheart Gloves (16717, -1.86 DPS) [dungeon]; Feralheart Gauntlets (226784, -5.97 DPS, sim-verified) [vendor] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | 85.9 healing_power points (14.20 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Marshal's Dragonhide Waistguard (16447, -0.30 DPS) [pvp]; Devout Belt (16696, -0.89 DPS) [dungeon] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 135.4 healing_power points (22.39 DPS) | yes | Devout Skirt (16694, -3.67 DPS) [dungeon]; Feralheart Pants (226787, -4.00 DPS) [vendor]; Knight-Captain's Dragonhide Legguards (227200, -11.52 DPS, sim-verified) [vendor] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 110.2 healing_power points (18.23 DPS) | yes | Feralheart Sandals (226781, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.68 DPS) [crafted]; Faith Healer's Boots (22247, -4.54 DPS) [dungeon] |
| finger1 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (599.2 DPS) | yes | Band of Piety (22681, -0.46 DPS) [quest]; Seal of Rivendare (13345, -0.83 DPS) [dungeon]; Naglering (11669, -9.75 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (599.2 DPS) | yes | Band of Piety (22681, -0.33 DPS) [quest]; Seal of Rivendare (13345, -0.70 DPS) [dungeon]; Naglering (11669, -11.64 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-verified (599.2 DPS) | yes | Serenity Field (272439, -0.52 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -1.21 DPS) [dungeon]; Briarwood Reed (12930, -2.48 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (599.2 DPS) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -2.21 DPS) [dungeon]; Briarwood Reed (12930, -3.48 DPS) [dungeon] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-verified (599.2 DPS) | yes | Hammer of the Grand Crusader (18717, -0.86 DPS) [dungeon]; Staff of Metanoia (22394, -1.66 DPS) [dungeon]; Hand of Edward the Odd (2243, -7.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Maye's Pendant; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Bracers of Hope; hands: Hands of the Exalted Herald; waist: Feralheart Cord; legs: Leggings of Arcana; feet: Incandescent Mooncloth Boots; finger1: Emerald Flame Ring; finger2: Band of Mending; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 47.5. Weights run: 5.3s. Verify run: 18.7s. 183 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.281 ± 0.004, spirit=1.542 ± 0.008, mp5=3.288 ± 0.015, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.529 per %), spell_haste=-5.534 ± 0.104

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.24 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 6.2 healing_power points (0.63 DPS) | yes | Roadwatcher's Confidence (281265, -2.07 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 healing_power points (1.18 DPS) | yes | Slime-encrusted Pads (6461, -0.38 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.55 DPS) [world_drop]; Prospector's Pads (14566, -0.55 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 12.1 healing_power points (1.24 DPS) | yes | Traveler's Shawl (277289, -0.45 DPS) [quest]; Spirit Cloak (4792, -0.61 DPS) [vendor]; Regent's Cloak (5969, -0.85 DPS, sim-verified) [world] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.0 healing_power points (2.67 DPS) | yes | Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon]; Robe of the Moccasin (6465, -1.41 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 16.3 healing_power points (1.67 DPS) | yes | Drakewing Bands (12999, -0.29 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.88 DPS) [quest]; Crystalline Cuffs (14148, -0.93 DPS) [dungeon] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 16.7 healing_power points (1.71 DPS) | yes | Wisdom's Leather Gloves (252499, +0.00 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Magefist Gloves (12977, -0.42 DPS) [world_drop] |
| waist | Novice Ardent's Sash (253887) | Tailoring [crafted] | sim-verified (47.6 DPS) | yes | Wisdom's Leather Belt (252433, -0.14 DPS) [crafted]; Keller's Girdle (2911, -0.43 DPS) [world_drop]; Pristine Sash (253925, -0.56 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.79 DPS) | yes | Pristine Boots (253889, -0.47 DPS) [crafted]; Smoldering Boots (3076, -0.84 DPS) [world]; Black Whelp Slippers (252424, -2.43 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.8 healing_power points (1.21 DPS) | yes | Deep Fathom Ring (6463, -0.42 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; Advisor's Ring (20426, -0.54 DPS) [rep] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.95 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Advisor's Ring (20426, -0.27 DPS) [rep]; Deep Fathom Ring (6463, -2.07 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 19.4 healing_power points (1.99 DPS) | yes | Staff of the Blessed Seer (2271, -0.37 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (20425, -0.45 DPS) [pvp]; Gnarled Necromancer's Staff (251534, -0.68 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Novice Ardent's Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Band of Purification; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 89.0. Weights run: 7.8s. Verify run: 26.8s. 315 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.448 ± 0.014, spirit=2.905 ± 0.018, mp5=4.429 ± 0.027, crit=0.046 ± 0.004 per rating point (14 rating = 1%, 0.641 per %), spell_haste=-1.040 ± 0.256

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Whisperwind Headdress (6688) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Enduring Cap (3020, -0.44 DPS) [world_drop]; Embalmed Shroud (7691, -1.32 DPS) [dungeon]; Holy Shroud (2721, -2.84 DPS, sim-verified) [world_drop] |
| neck | Necklace of Harmony (5180) | Singer [world] | 20.3 healing_power points (2.06 DPS) | yes | Pendant of Myzrael (4614, -0.29 DPS) [dungeon]; Glowing Green Talisman (5002, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.30 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 35.1 healing_power points (3.57 DPS) | yes | Batwing Mantle (6697, +0.00 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.92 DPS) [world_drop]; Desert Shoulders (15457, -1.36 DPS) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Amy's Blanket (13005, -0.44 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -0.44 DPS) [vendor]; Glowing Thresher Cape (6901, -1.33 DPS, sim-verified) [dungeon] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.1 healing_power points (4.48 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.36 DPS) [dungeon]; Pressed Felt Robe (1997, -0.50 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.5 healing_power points (2.09 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Drakewing Bands (12999, -0.40 DPS, sim-verified) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.1 healing_power points (4.28 DPS) | yes | Hotshot Pilot's Gloves (9491, -1.62 DPS) [dungeon]; Blight Gloves (279877, -1.77 DPS) [quest]; Tattered Mittens (270030, -1.81 DPS, sim-verified) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.1 healing_power points (4.58 DPS) | yes | Dokebi Cord (14578, -2.22 DPS) [world_drop]; Lilac Sash (6780, -2.37 DPS) [quest]; Silver-lined Belt (13011, -2.86 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.5 healing_power points (5.23 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Earthen Leggings (253999, -0.66 DPS, sim-verified) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.24 DPS) | yes | Boots of the Enchanter (4325, -1.88 DPS) [crafted]; Stonecloth Boots (14408, -1.88 DPS) [world_drop]; Soggy Boots (274747, -2.21 DPS, sim-verified) [vendor] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.5 healing_power points (3.09 DPS) | yes | Monkey Ring (6748, -1.03 DPS) [quest]; Electrocutioner Lagnut (9447, -1.03 DPS) [dungeon]; Black Pearl Ring (6332, -1.03 DPS) [world] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 26.1 healing_power points (2.65 DPS) | yes | Monkey Ring (6748, -0.18 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.59 DPS) [dungeon]; Black Pearl Ring (6332, -0.59 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wind Spirit Staff (6689, -1.75 DPS) [dungeon]; Gnarled Ash Staff (791, -2.48 DPS) [world_drop]; Manual Crowd Pummeler (9449, -3.19 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.4 healing_power points (3.09 DPS) | yes | Orb of Mistmantle (13031, -1.00 DPS) [world_drop]; Strength of Will (4837, -1.32 DPS) [vendor]; Defective Samophlange (274743, -1.45 DPS, sim-verified) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Whisperwind Headdress; neck: Necklace of Harmony; shoulder: Ghostly Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 135.8. Weights run: 8.7s. Verify run: 53.9s. 426 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.339 ± 0.016, spirit=2.500 ± 0.011, mp5=4.838 ± 0.034, crit=0.050 ± 0.005 per rating point (14 rating = 1%, 0.694 per %), spell_haste=not significant (0.513 ± 0.312)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 67.3 healing_power points (6.90 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -1.76 DPS) [dungeon]; Holy Shroud (2721, -1.98 DPS) [world_drop]; Whitemane's Chapeau (7720, -2.43 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 34.2 healing_power points (3.51 DPS) | yes | Triune Amulet (7722, -0.10 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.75 DPS) [dungeon]; Amberglow Talisman (10824, -0.94 DPS) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.0 healing_power points (3.90 DPS) | yes | Ghostly Mantle (3324, -0.67 DPS) [quest]; Sheepshear Mantle (13115, -0.85 DPS, sim-verified) [world_drop]; Crimson Silk Shoulders (7059, -1.00 DPS) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 32.5 healing_power points (3.34 DPS) | yes | Battle Healer's Cloak (19528, +0.00 DPS, sim-verified) [rep]; Glowing Thresher Cape (6901, -0.56 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.75 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.0 healing_power points (8.00 DPS) | yes | Dreamweave Vest (10021, -3.18 DPS) [crafted]; Black Mageweave Robe (10001, -3.31 DPS) [crafted]; Doomsayer's Robe (4746, -7.10 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.4 healing_power points (3.12 DPS) | yes | Earthen Silk Cuffs (254019, -0.10 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Dryad's Wrist Bindings (19597, -1.01 DPS) [pvp] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (135.8 DPS) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Mender's Leather Gloves (252530, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, -1.40 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.7 healing_power points (4.90 DPS) | yes | Mender's Leather Belt (252523, -0.37 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (135.8 DPS) | yes | Wisdom's Leather Leggings (252519, -0.47 DPS) [crafted]; Pristine Leggings (253987, -0.89 DPS) [crafted]; Warchief Kilt (7760, -3.30 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 50.9 healing_power points (5.22 DPS) | yes | Thoughtcast Boots (10578, -1.06 DPS) [dungeon]; Furen's Boots (13100, -1.14 DPS, sim-verified) [world_drop]; Gilded Slippers (254001, -1.18 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.0 healing_power points (2.98 DPS) | yes | Welken Ring (5011, -0.50 DPS) [world_drop]; The Queen's Jewel (13094, -0.65 DPS) [world_drop]; Blush Ember Ring (13093, -0.93 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 26.9 healing_power points (2.76 DPS) | yes | Welken Ring (5011, +0.00 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.43 DPS) [world_drop]; Blush Ember Ring (13093, -0.71 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (135.8 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (135.8 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -2.44 DPS) [dungeon]; Staff of Jordan (873, -2.64 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 32.0 healing_power points (3.28 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.46 DPS) [dungeon]; Orb of Souls (249395, -0.62 DPS, sim-verified) [crafted]; Silksand Star (15964, -0.97 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; off_hand: Beacon of Hope

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 198.9. Weights run: 8.6s. Verify run: 29.7s. 561 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.522 ± 0.017, spirit=2.750 ± 0.016, mp5=4.393 ± 0.016, crit=0.069 ± 0.007 per rating point (14 rating = 1%, 0.967 per %), spell_haste=0.966 ± 0.177

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 75.4 healing_power points (7.47 DPS) | yes | Papal Fez (9431, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Restored Leather Helm (220875, +0.00 DPS) [vendor]; Cassandra's Grace (13102, -0.66 DPS) [world_drop] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 41.2 healing_power points (4.09 DPS) | yes | Darkmoon Necklace (19303, -0.57 DPS) [vendor]; Glowing Eye of Mordresh (10769, -0.78 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -0.89 DPS) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 66.7 healing_power points (6.62 DPS) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.20 DPS) [vendor]; Nethergeld Shoulders (254049, -1.20 DPS) [crafted]; Mender's Leather Shoulder (252538, -1.81 DPS, sim-verified) [crafted] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | 47.3 healing_power points (4.69 DPS) | yes | Battle Healer's Cloak (19527, -0.31 DPS, sim-verified) [rep]; Cloak of Blight (6832, -1.15 DPS) [quest]; Darkspear Raider's Cloak (272076, -1.22 DPS) [vendor] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 108.4 healing_power points (10.74 DPS) | yes | Ghostweave Vest (14141, -1.81 DPS) [crafted]; Forest's Embrace (22272, -2.17 DPS) [quest]; Vestments of the Atal'ai Prophet (10806, -2.29 DPS, sim-verified) [dungeon] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.57 DPS) | yes | Aristocratic Cuffs (12546, -0.68 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop]; Nethergeld Cuffs (254061, -1.41 DPS, sim-verified) [crafted] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.2 healing_power points (8.94 DPS) | yes | Mender's Leather Gauntlets (252551, +0.00 DPS, sim-verified) [crafted]; Gilded Gloves (254095, -2.38 DPS) [crafted]; Earthenweave Gloves (254075, -3.37 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.2 healing_power points (5.58 DPS) | yes | Mender's Leather Waistguard (252477, -0.79 DPS) [crafted]; Mender's Leather Belt (252523, -1.15 DPS) [crafted]; Gilded Cord (254037, -1.29 DPS, sim-verified) [crafted] |
| legs | Dalewind Trousers (13008) | World drop [world_drop] | 82.4 healing_power points (8.17 DPS) | yes | Rainstrider Leggings (11123, -0.81 DPS) [quest]; Kilt of the Atal'ai Prophet (10807, -1.02 DPS, sim-verified) [dungeon]; Windscale Sarong (10842, -1.21 DPS) [world] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.2 healing_power points (6.66 DPS) | yes | Mistwalker Boots (10629, -0.55 DPS) [dungeon]; Earthenweave Boots (254093, -1.08 DPS) [crafted]; First Sergeant's Restored Leather Boots (220885, -3.16 DPS, sim-verified) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.1 healing_power points (4.18 DPS) | yes | Darkspear Signet (272069, -1.13 DPS) [vendor]; Cyclopean Band (11824, -1.14 DPS) [dungeon]; Snake Hoop (6750, -1.21 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 36.6 healing_power points (3.63 DPS) | yes | Cyclopean Band (11824, -0.59 DPS) [dungeon]; Snake Hoop (6750, -0.66 DPS) [quest]; Darkspear Signet (272069, -1.11 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (198.9 DPS) | yes | Evonice's Landin' Pilla (18951, -3.81 DPS) [quest]; Uther's Strength (11302, -4.75 DPS) [world_drop]; Alchemists' Stone (13503, -5.44 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (198.9 DPS) | yes | Evonice's Landin' Pilla (18951, -1.06 DPS, sim-verified) [quest]; Uther's Strength (11302, -1.49 DPS) [world_drop]; Alchemists' Stone (13503, -2.18 DPS) [crafted] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-verified (198.9 DPS) | yes | Glowing Brightwood Staff (812, -1.92 DPS) [world_drop]; Death Speaker Scepter (2816, -2.01 DPS) [dungeon]; Barman Shanker (12791, -9.70 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Lei of Lilies; shoulder: Living Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Dalewind Trousers; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Soulkeeper

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 321.8. Weights run: 9.0s. Verify run: 69.5s. 1446 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 173.1 healing_power points (17.46 DPS) | yes | Champion's Dragonhide Headdress (227205, -2.67 DPS) [vendor]; Feralheart Headdress (226786, -3.26 DPS) [vendor]; Wildheart Cowl (16720, -10.34 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 92.3 healing_power points (9.31 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.94 DPS) [world_drop]; Heart of the Fiend (13960, -1.72 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.46 DPS, sim-verified) [world_drop] |
| shoulder | Feralheart Mantle (226785) | Mokvar [vendor] | sim-verified (321.8 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Argent Elite Shoulders (227888, -4.89 DPS, sim-verified) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 78.1 healing_power points (7.88 DPS) | yes | Featherskin Cape (10843, -0.80 DPS, sim-verified) [world]; Butcher's Apron (12608, -0.98 DPS) [dungeon]; Hide of the Wild (18510, -1.39 DPS) [crafted] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, -1.45 DPS) [crafted]; Alanna's Embrace (13314, -1.62 DPS) [dungeon]; Tunic of Undead Slaying (23089, -13.76 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, -0.02 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.12 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -10.06 DPS, sim-verified) [world] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 125.5 healing_power points (12.66 DPS) | yes | Wildheart Gloves (16717, +0.00 DPS, sim-verified) [dungeon]; Hands of the Exalted Herald (12554, -1.23 DPS) [dungeon]; Devout Gloves (16692, -1.56 DPS) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | 125.6 healing_power points (12.67 DPS) | yes | Feralheart Cord (226780, +0.00 DPS, sim-verified) [vendor]; Elderwild Waistcord (279252, -1.88 DPS) [crafted]; Wisdom of the Timbermaw (19047, -2.51 DPS) [crafted] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 191.0 healing_power points (19.26 DPS) | yes | Haunting Specter Leggings (11929, -4.48 DPS) [dungeon]; Legionnaire's Dragonhide Legguards (227206, -4.55 DPS) [vendor]; Devout Skirt (16694, -7.70 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 148.3 healing_power points (14.96 DPS) | yes | Incandescent Mooncloth Boots (227862, -3.09 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.78 DPS) [crafted]; Devout Sandals (16691, -4.26 DPS) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.72 DPS) [dungeon]; Band of Piety (22681, -0.77 DPS) [quest]; Naglering (11669, -6.51 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.38 DPS) [dungeon]; Band of Piety (22681, -0.43 DPS) [quest]; Naglering (11669, -8.09 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+8.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindtap Talisman (18371, -2.77 DPS) [dungeon]; Ankh of Life (1713, -5.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.07 DPS) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Serenity Field (272439, -0.36 DPS, sim-verified) [vendor] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -0.09 DPS) [world_drop]; Soulkeeper (1607, -1.52 DPS) [world_drop]; Hand of Edward the Odd (2243, -6.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Feralheart Mantle; back: Frostweaver Cape; wrist: Bracers of Hope; waist: Caretaker's Cord; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: The Postmaster's Seal; finger2: Emerald Flame Ring; trinket2: Royal Seal of Eldre'Thalas; main_hand: Dancing Sliver

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 597.6. Weights run: 6.5s. Verify run: 47.3s. 1446 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.204, intellect=2.451 ± 0.080, spirit=2.281 ± 0.075, mp5=3.336 ± 0.099, crit=0.327 ± 0.030 per rating point (14 rating = 1%, 4.573 per %), spell_haste=not significant (-0.743 ± 1.216)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 149.0 healing_power points (24.64 DPS) | yes | Champion's Dragonhide Headdress (227205, -4.67 DPS) [vendor]; Sanctified Leather Helm (22689, -5.80 DPS) [quest]; Feralheart Headdress (226786, -16.93 DPS, sim-verified) [vendor] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 69.4 healing_power points (11.47 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.41 DPS) [world_drop]; Wavefront Necklace (20685, -0.66 DPS) [world]; Drake Tooth Necklace (21531, -2.31 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 101.8 healing_power points (16.83 DPS) | yes | Devout Mantle (16695, +0.00 DPS, sim-verified) [dungeon]; Champion's Dragonhide Pauldrons (227207, -1.16 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, -1.32 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 66.5 healing_power points (10.99 DPS) | yes | Cloak of the Cosmos (18389, -2.24 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -2.25 DPS) [vendor]; Frostweaver Cape (12968, -4.70 DPS, sim-verified) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-verified (597.6 DPS) | yes | Alanna's Embrace (13314, -0.61 DPS) [dungeon]; Devout Robe (16690, -0.88 DPS) [dungeon]; Tunic of Undead Slaying (23089, -19.71 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (597.6 DPS) | yes | Bracers of Mending (23129, -0.43 DPS) [dungeon]; General's Dragonhide Bracers (16553, -0.49 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -11.83 DPS, sim-verified) [world] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 92.2 healing_power points (15.25 DPS) | yes | Raider Handwraps (272097, -0.34 DPS) [vendor]; Wildheart Gloves (16717, -1.86 DPS) [dungeon]; Feralheart Gauntlets (226784, -6.27 DPS, sim-verified) [vendor] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | 85.9 healing_power points (14.20 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; General's Dragonhide Belt (16556, -0.30 DPS) [pvp]; Devout Belt (16696, -0.89 DPS) [dungeon] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 135.4 healing_power points (22.39 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, -2.65 DPS) [vendor]; Feralheart Pants (226787, -4.00 DPS) [vendor]; Devout Skirt (16694, -8.61 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 110.2 healing_power points (18.23 DPS) | yes | Feralheart Sandals (226781, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.68 DPS) [crafted]; Faith Healer's Boots (22247, -4.54 DPS) [dungeon] |
| finger1 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (597.6 DPS) | yes | Band of Piety (22681, -0.46 DPS) [quest]; Seal of Rivendare (13345, -0.83 DPS) [dungeon]; Naglering (11669, -10.31 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (597.6 DPS) | yes | Band of Piety (22681, -0.33 DPS) [quest]; Seal of Rivendare (13345, -0.70 DPS) [dungeon]; Naglering (11669, -12.34 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-verified (597.6 DPS) | yes | Serenity Field (272439, -0.56 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -1.21 DPS) [dungeon]; Briarwood Reed (12930, -2.48 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (597.6 DPS) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -2.21 DPS) [dungeon]; Briarwood Reed (12930, -3.48 DPS) [dungeon] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-verified (597.6 DPS) | yes | Hammer of the Grand Crusader (18717, -0.86 DPS) [dungeon]; Staff of Metanoia (22394, -1.66 DPS) [dungeon]; Hand of Edward the Odd (2243, -7.39 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Maye's Pendant; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Bracers of Hope; hands: Hands of the Exalted Herald; waist: Feralheart Cord; legs: Leggings of Arcana; feet: Incandescent Mooncloth Boots; finger1: Emerald Flame Ring; finger2: Band of Mending; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

