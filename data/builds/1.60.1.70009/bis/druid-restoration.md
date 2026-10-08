# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 43.5. Weights run: 5.4s. Verify run: 6.3s. 193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.281 ± 0.004, spirit=1.542 ± 0.008, mp5=3.288 ± 0.015, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.529 per %), spell_haste=-5.534 ± 0.104

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.10 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.24 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 6.2 healing_power points (0.63 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Mantle (4709, -0.55 DPS) [world_drop]; Prospector's Pads (14566, -0.55 DPS) [world_drop]; Slime-encrusted Pads (6461, -1.33 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Spirit Cloak (4792, -0.61 DPS) [vendor]; Sylvan Cloak (4793, -0.61 DPS) [vendor]; Regent's Cloak (5969, -1.47 DPS, sim-verified) [world] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wisdom's Leather Armor (252493, +0.00 DPS) [crafted]; Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Robe of the Moccasin (6465, -9.92 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 16.3 healing_power points (1.67 DPS) | yes | Owl Bracers (4796, -1.01 DPS) [vendor]; Bravo's Armbands (270015, -1.04 DPS) [quest]; Drakewing Bands (12999, -9.79 DPS, sim-verified) [world_drop] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Magefist Gloves (12977, -0.26 DPS) [world_drop]; Bright Gloves (3066, -0.39 DPS) [world_drop]; Pristine Gloves (253913, -2.82 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Belt (252433, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.60 DPS) [world_drop]; Novice Ardent's Sash (253887, -1.05 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.19 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Pristine Boots (253889, +0.00 DPS) [crafted]; Black Whelp Slippers (252424, -7.50 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.8 healing_power points (1.21 DPS) | yes | Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.54 DPS) [rep]; Band of Purification (12996, -1.52 DPS, sim-verified) [world_drop] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lavishly Jeweled Ring (1156, -0.00 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.12 DPS) [rep]; Band of Purification (12996, -2.73 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 15.7 healing_power points (1.61 DPS) | yes | Twisted Chanter's Staff (890, -0.29 DPS) [world_drop]; Gnarled Hermit's Staff (1539, -0.50 DPS) [world]; Staff of the Blessed Seer (2271, -3.70 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Armor of the Fang; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Footpads of the Fang; finger1: Black Pearl Ring; finger2: Deep Fathom Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 83.2. Weights run: 7.7s. Verify run: 4.4s. 322 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.448 ± 0.014, spirit=2.905 ± 0.018, mp5=4.429 ± 0.027, crit=0.046 ± 0.004 per rating point (14 rating = 1%, 0.641 per %), spell_haste=-1.040 ± 0.256

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Whisperwind Headdress (6688) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+9.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Enduring Cap (3020, -0.44 DPS) [world_drop]; Embalmed Shroud (7691, -1.32 DPS) [dungeon]; Holy Shroud (2721, -9.14 DPS, sim-verified) [world_drop] |
| neck | Pendant of Myzrael (4614) | Razorfen Downs: Splinterbone Captain [dungeon] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.00 DPS) [world_drop]; Necklace of Harmony (5180, -4.61 DPS, sim-verified) [world] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Faerie Mantle (5820, -0.44 DPS) [quest]; Nightsky Mantle (4718, -0.44 DPS) [world_drop]; Mantle of Honor (3560, -3.89 DPS, sim-verified) [quest] |
| back | Prelacy Cape (7004) | Researching the Corruption [quest] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Repairman's Cape (9605, -0.23 DPS) [quest]; Caretaker's Cape (19533, -0.39 DPS) [rep]; Glowing Thresher Cape (6901, -1.32 DPS, sim-verified) [dungeon] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.12 DPS) [dungeon]; Pressed Felt Robe (1997, -0.26 DPS) [world]; Wisdom's Leather Tunic (252511, -4.73 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.5 healing_power points (2.09 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop]; Drakewing Bands (12999, -1.25 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) | Gyrodrillmatic Excavationators [quest] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Zodiac Gloves (7106, -0.00 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.00 DPS) [dungeon]; Gloves of Old (9395, -1.29 DPS, sim-verified) [world_drop] |
| waist | Silver-lined Belt (13011) | World drop [world_drop] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Dokebi Cord (14578, -0.15 DPS) [world_drop]; Resilient Cord (14406, -0.44 DPS) [world_drop]; Mender's Leather Belt (252523, -4.93 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.5 healing_power points (5.23 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted]; Earthen Leggings (253999, -3.40 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.24 DPS) | yes | Boots of the Enchanter (4325, -1.88 DPS) [crafted]; Stonecloth Boots (14408, -1.88 DPS) [world_drop]; Soggy Boots (274747, -4.75 DPS, sim-verified) [vendor] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.5 healing_power points (3.09 DPS) | yes | Monkey Ring (6748, -1.03 DPS) [quest]; Ring of Calm (6790, -1.03 DPS) [quest]; Electrocutioner Lagnut (9447, -1.03 DPS) [dungeon] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 26.1 healing_power points (2.65 DPS) | yes | Ring of Calm (6790, -0.19 DPS, sim-verified) [quest]; Monkey Ring (6748, -0.59 DPS) [quest]; Electrocutioner Lagnut (9447, -0.59 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wind Spirit Staff (6689, -1.75 DPS) [dungeon]; Manual Crowd Pummeler (9449, -2.32 DPS, sim-verified) [dungeon]; Gnarled Ash Staff (791, -2.48 DPS) [world_drop] |
| off_hand | Defective Samophlange (274743) | Winklespark [vendor] | sim-verified (+7.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Mistmantle (13031, -0.12 DPS) [world_drop]; Strength of Will (4837, -0.44 DPS) [vendor]; Orb of Souls (249395, -7.43 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Whisperwind Headdress; neck: Pendant of Myzrael; shoulder: Batwing Mantle; back: Prelacy Cape; chest: Pristine Gown; hands: Shilly Mitts; waist: Silver-lined Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Defective Samophlange

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 135.4. Weights run: 8.6s. Verify run: 9.2s. 438 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.339 ± 0.016, spirit=2.500 ± 0.011, mp5=4.838 ± 0.034, crit=0.050 ± 0.005 per rating point (14 rating = 1%, 0.694 per %), spell_haste=not significant (0.513 ± 0.312)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 67.3 healing_power points (6.90 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -1.76 DPS) [dungeon]; Holy Shroud (2721, -1.98 DPS) [world_drop]; Whitemane's Chapeau (7720, -5.75 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (74.6 DPS) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Amberglow Talisman (10824, -0.19 DPS) [quest]; Glowing Eye of Mordresh (10769, -3.15 DPS, sim-verified) [dungeon] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | sim-verified (74.6 DPS) | yes | Crimson Silk Shoulders (7059, -0.84 DPS) [crafted]; Mantle of Doan (7712, -0.84 DPS) [dungeon]; Earthen Silk Shoulders (254033, -7.56 DPS, sim-verified) [crafted] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | 30.5 healing_power points (3.13 DPS) | yes | Prelacy Cape (7004, -0.46 DPS) [quest]; Ceremonial Centaur Blanket (6789, -0.55 DPS) [quest]; Glowing Thresher Cape (6901, -1.46 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.0 healing_power points (8.00 DPS) | yes | Dreamweave Vest (10021, -3.18 DPS) [crafted]; Black Mageweave Robe (10001, -3.31 DPS) [crafted]; Doomsayer's Robe (4746, -8.69 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.4 healing_power points (3.12 DPS) | yes | Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Silkstream Cuffs (16791, -1.40 DPS) [quest]; Earthen Silk Cuffs (254019, -6.92 DPS, sim-verified) [crafted] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (74.6 DPS) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Mender's Leather Gloves (252530, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, -7.99 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.7 healing_power points (4.90 DPS) | yes | Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop]; Mender's Leather Belt (252523, -4.29 DPS, sim-verified) [crafted] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (74.6 DPS) | yes | Wisdom's Leather Leggings (252519, -0.47 DPS) [crafted]; Pristine Leggings (253987, -0.89 DPS) [crafted]; Warchief Kilt (7760, -7.32 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 50.9 healing_power points (5.22 DPS) | yes | Thoughtcast Boots (10578, -1.06 DPS) [dungeon]; Gilded Slippers (254001, -1.18 DPS) [crafted]; Furen's Boots (13100, -3.38 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.0 healing_power points (2.98 DPS) | yes | Welken Ring (5011, -0.50 DPS) [world_drop]; The Queen's Jewel (13094, -0.65 DPS) [world_drop]; Blush Ember Ring (13093, -0.93 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 26.9 healing_power points (2.76 DPS) | yes | The Queen's Jewel (13094, -0.43 DPS) [world_drop]; Blush Ember Ring (13093, -0.71 DPS) [world_drop]; Welken Ring (5011, -4.05 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (74.6 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (74.6 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -2.44 DPS) [dungeon]; Staff of Jordan (873, -2.64 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 32.0 healing_power points (3.28 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.46 DPS) [dungeon]; Silksand Star (15964, -0.97 DPS) [world_drop]; Orb of Souls (249395, -2.37 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Sheepshear Mantle; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; off_hand: Beacon of Hope

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 198.9. Weights run: 8.7s. Verify run: 5.4s. 577 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.522 ± 0.017, spirit=2.750 ± 0.016, mp5=4.393 ± 0.016, crit=0.069 ± 0.007 per rating point (14 rating = 1%, 0.967 per %), spell_haste=0.966 ± 0.177

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Papal Fez (9431, -0.27 DPS) [dungeon]; Cassandra's Grace (13102, -0.66 DPS) [world_drop]; Knight-Lieutenant's Restored Leather Helm (220874, -2.96 DPS, sim-verified) [vendor] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 41.2 healing_power points (4.09 DPS) | yes | Darkmoon Necklace (19303, -0.57 DPS) [vendor]; Horizon Choker (13085, -0.89 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.66 DPS, sim-verified) [dungeon] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 66.7 healing_power points (6.62 DPS) | yes | Knight-Lieutenant's Restored Leather Spaulders (220876, -0.76 DPS, sim-verified) [vendor]; Mender's Leather Shoulder (252538, -0.97 DPS) [crafted]; Nethergeld Shoulders (254049, -1.20 DPS) [crafted] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | 47.3 healing_power points (4.69 DPS) | yes | Darkspear Raider's Cloak (272076, -1.22 DPS) [vendor]; Arcane Cloak (8286, -1.42 DPS) [world_drop]; Caretaker's Cape (19531, -4.06 DPS, sim-verified) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 108.4 healing_power points (10.74 DPS) | yes | Ghostweave Vest (14141, -1.81 DPS) [crafted]; Forest's Embrace (22272, -2.17 DPS) [quest]; Vestments of the Atal'ai Prophet (10806, -3.30 DPS, sim-verified) [dungeon] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.57 DPS) | yes | Aristocratic Cuffs (12546, -0.68 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop]; Nethergeld Cuffs (254061, -2.79 DPS, sim-verified) [crafted] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.2 healing_power points (8.94 DPS) | yes | Mender's Leather Gauntlets (252551, +0.00 DPS, sim-verified) [crafted]; Gilded Gloves (254095, -2.38 DPS) [crafted]; Earthenweave Gloves (254075, -3.37 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.2 healing_power points (5.58 DPS) | yes | Gilded Cord (254037, -0.71 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -0.79 DPS) [crafted]; Mender's Leather Belt (252523, -1.15 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Rainstrider Leggings (11123, -0.26 DPS) [quest]; Windscale Sarong (10842, -0.66 DPS) [world]; Dalewind Trousers (13008, -1.48 DPS, sim-verified) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.2 healing_power points (6.66 DPS) | yes | Mistwalker Boots (10629, -0.55 DPS) [dungeon]; Earthenweave Boots (254093, -1.08 DPS) [crafted]; Sergeant Major's Restored Leather Boots (220884, -1.13 DPS, sim-verified) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.1 healing_power points (4.18 DPS) | yes | Choking Band (11868, -0.63 DPS) [quest]; Darkspear Signet (272069, -1.13 DPS) [vendor]; Cyclopean Band (11824, -1.14 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 36.6 healing_power points (3.63 DPS) | yes | Darkspear Signet (272069, -0.58 DPS) [vendor]; Cyclopean Band (11824, -0.59 DPS) [dungeon]; Choking Band (11868, -3.99 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -3.81 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.35 DPS) [quest]; Uther's Strength (11302, -4.75 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.55 DPS) [quest]; Uther's Strength (11302, -1.49 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -1.52 DPS, sim-verified) [quest] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -1.92 DPS) [world_drop]; Death Speaker Scepter (2816, -2.01 DPS) [dungeon]; Barman Shanker (12791, -9.28 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Lei of Lilies; shoulder: Living Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Soulkeeper

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 324.7. Weights run: 9.1s. Verify run: 13.2s. 1449 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 173.1 healing_power points (17.46 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -2.67 DPS) [vendor]; Feralheart Headdress (226786, -3.26 DPS) [vendor]; Wildheart Cowl (16720, -8.37 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 92.3 healing_power points (9.31 DPS) | yes | Lady Maye's Pendant (14558, +0.00 DPS, sim-verified) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.94 DPS) [world_drop]; Heart of the Fiend (13960, -1.72 DPS) [dungeon] |
| shoulder | Feralheart Mantle (226785) | Mokvar [vendor] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Argent Elite Shoulders (227888, -2.14 DPS, sim-verified) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 78.1 healing_power points (7.88 DPS) | yes | Butcher's Apron (12608, -0.98 DPS) [dungeon]; Hide of the Wild (18510, -1.39 DPS) [crafted]; Featherskin Cape (10843, -6.07 DPS, sim-verified) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, -1.45 DPS) [crafted]; Alanna's Embrace (13314, -1.62 DPS) [dungeon]; Tunic of Undead Slaying (23089, -15.77 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, -0.02 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.12 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -13.60 DPS, sim-verified) [world] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 125.5 healing_power points (12.66 DPS) | yes | Hands of the Exalted Herald (12554, -1.23 DPS) [dungeon]; Devout Gloves (16692, -1.56 DPS) [dungeon]; Wildheart Gloves (16717, -3.52 DPS, sim-verified) [dungeon] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Elderwild Waistcord (279252, -0.72 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.34 DPS) [crafted]; Caretaker's Cord (272398, -3.85 DPS, sim-verified) [vendor] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 191.0 healing_power points (19.26 DPS) | yes | Haunting Specter Leggings (11929, -4.48 DPS) [dungeon]; Knight-Captain's Dragonhide Legguards (227200, -4.55 DPS) [vendor]; Devout Skirt (16694, -9.75 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 148.3 healing_power points (14.96 DPS) | yes | Mooncloth Boots (15802, -3.78 DPS) [crafted]; Devout Sandals (16691, -4.26 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -4.97 DPS, sim-verified) [vendor] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.72 DPS) [dungeon]; Band of Piety (22681, -0.77 DPS) [quest]; Naglering (11669, -11.04 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.38 DPS) [dungeon]; Band of Piety (22681, -0.43 DPS) [quest]; Naglering (11669, -10.49 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+11.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindtap Talisman (18371, -2.77 DPS) [dungeon]; Ankh of Life (1713, -5.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.07 DPS) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Serenity Field (272439, -0.50 DPS, sim-verified) [vendor] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -0.09 DPS) [world_drop]; Soulkeeper (1607, -1.52 DPS) [world_drop]; Hand of Edward the Odd (2243, -10.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Feralheart Mantle; back: Frostweaver Cape; wrist: Bracers of Hope; waist: Feralheart Cord; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: The Postmaster's Seal; finger2: Emerald Flame Ring; trinket2: Royal Seal of Eldre'Thalas; main_hand: Dancing Sliver

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 598.5. Weights run: 6.5s. Verify run: 10.2s. 1449 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.236, intellect=2.459 ± 0.082, spirit=2.390 ± 0.078, mp5=3.433 ± 0.102, crit=0.313 ± 0.025 per rating point (14 rating = 1%, 4.384 per %), spell_haste=not significant (-0.554 ± 1.262)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 150.7 healing_power points (24.33 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -4.54 DPS) [vendor]; Sanctified Leather Helm (22689, -5.70 DPS) [quest]; Feralheart Headdress (226786, -8.54 DPS, sim-verified) [vendor] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 70.6 healing_power points (11.40 DPS) | yes | Wavefront Necklace (20685, -0.71 DPS) [world]; Amulet of the Redeemed (22327, -2.32 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -8.37 DPS, sim-verified) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 103.5 healing_power points (16.70 DPS) | yes | Devout Mantle (16695, +0.00 DPS, sim-verified) [dungeon]; Lieutenant Commander's Dragonhide Pauldrons (227201, -1.14 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -1.36 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 66.6 healing_power points (10.75 DPS) | yes | Darkspear Raider's Cloak (272063, -2.08 DPS) [vendor]; Shroud of the Exile (15421, -2.09 DPS) [quest]; Frostweaver Cape (12968, -3.77 DPS, sim-verified) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-verified (526.9 DPS) | yes | Alanna's Embrace (13314, -0.58 DPS) [dungeon]; Devout Robe (16690, -0.92 DPS) [dungeon]; Tunic of Undead Slaying (23089, -29.67 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (526.9 DPS) | yes | Bracers of Mending (23129, -0.41 DPS) [dungeon]; Marshal's Dragonhide Bracers (16445, -0.50 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -16.39 DPS, sim-verified) [world] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 94.1 healing_power points (15.18 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS, sim-verified) [dungeon]; Raider Handwraps (272097, -0.59 DPS) [vendor]; Wildheart Gloves (16717, -1.73 DPS) [dungeon] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | 87.5 healing_power points (14.12 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Marshal's Dragonhide Waistguard (16447, -0.34 DPS) [pvp]; Devout Belt (16696, -0.94 DPS) [dungeon] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 138.9 healing_power points (22.42 DPS) | yes | Devout Skirt (16694, -3.72 DPS) [dungeon]; Feralheart Pants (226787, -4.12 DPS) [vendor]; Knight-Captain's Dragonhide Legguards (227200, -17.21 DPS, sim-verified) [vendor] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 111.9 healing_power points (18.06 DPS) | yes | Feralheart Sandals (226781, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.61 DPS) [crafted]; Faith Healer's Boots (22247, -4.47 DPS) [dungeon] |
| finger1 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (526.9 DPS) | yes | Band of Piety (22681, -0.53 DPS) [quest]; Seal of Rivendare (13345, -0.82 DPS) [dungeon]; Naglering (11669, -22.97 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (526.9 DPS) | yes | Band of Piety (22681, -0.36 DPS) [quest]; Seal of Rivendare (13345, -0.65 DPS) [dungeon]; Naglering (11669, -22.00 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-verified (526.9 DPS) | yes | Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Mindtap Talisman (18371, -1.01 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (526.9 DPS) | yes | Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Mindtap Talisman (18371, -0.68 DPS) [dungeon]; Ankh of Life (1713, -2.05 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-verified (526.9 DPS) | yes | Hammer of the Grand Crusader (18717, -1.07 DPS) [dungeon]; Staff of Metanoia (22394, -1.79 DPS) [dungeon]; Hand of Edward the Odd (2243, -17.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Maye's Pendant; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Bracers of Hope; waist: Feralheart Cord; legs: Leggings of Arcana; feet: Incandescent Mooncloth Boots; finger1: Emerald Flame Ring; finger2: Band of Mending; trinket1: Royal Seal of Eldre'Thalas; trinket2: Serenity Field; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 42.4. Weights run: 5.4s. Verify run: 6.2s. 183 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.281 ± 0.004, spirit=1.542 ± 0.008, mp5=3.288 ± 0.015, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.529 per %), spell_haste=-5.534 ± 0.104

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.09 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.24 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Scholarly Pendant (277203, -3.39 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 healing_power points (1.18 DPS) | yes | Forest Leather Mantle (4709, -0.55 DPS) [world_drop]; Prospector's Pads (14566, -0.55 DPS) [world_drop]; Slime-encrusted Pads (6461, -1.56 DPS, sim-verified) [dungeon] |
| back | Regent's Cloak (5969) | Ravenclaw Regent [world] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Traveler's Shawl (277289, +0.00 DPS) [quest]; Spirit Cloak (4792, -0.16 DPS) [vendor]; Battle Healer's Cloak (20427, -6.75 DPS, sim-verified) [rep] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, -0.04 DPS) [crafted]; Corsair's Overshirt (5202, -0.11 DPS) [dungeon]; Wisdom's Leather Armor (252493, -2.00 DPS, sim-verified) [crafted] |
| wrist | Drakewing Bands (12999) | World drop [world_drop] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Tabitha's Cuffs (251486, -0.16 DPS) [quest]; Crystalline Cuffs (14148, -0.21 DPS) [dungeon]; Mindthrust Bracers (1974, -3.51 DPS, sim-verified) [dungeon] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gloves (253913, -0.03 DPS) [crafted]; Magefist Gloves (12977, -0.26 DPS) [world_drop]; Blight Gloves (279877, -4.27 DPS, sim-verified) [quest] |
| waist | Novice Ardent's Sash (253887) | Tailoring [crafted] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Belt (252433, -0.14 DPS) [crafted]; Keller's Girdle (2911, -0.43 DPS) [world_drop]; Pristine Sash (253925, -5.81 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.48 DPS) | yes | Filigreed Pristine Leggings (253937, -0.17 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Black Whelp Slippers (252424) | Leatherworking [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Boots (253889, -0.13 DPS) [crafted]; Smoldering Boots (3076, -0.50 DPS) [world]; Wisdom's Leather Boots (252444, -1.55 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.8 healing_power points (1.21 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; Advisor's Ring (20426, -0.54 DPS) [rep] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.00 DPS) [dungeon]; Advisor's Ring (20426, -0.12 DPS) [rep]; Band of Purification (12996, -3.39 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Gnarled Staff (20425, -0.04 DPS) [pvp]; Gnarled Necromancer's Staff (251534, -0.27 DPS) [quest]; Staff of Orgrimmar (15444, -5.12 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Regent's Cloak; chest: Robe of the Moccasin; wrist: Drakewing Bands; hands: Wisdom's Leather Gloves; waist: Novice Ardent's Sash; legs: Wisdom's Leather Pants; feet: Black Whelp Slippers; finger1: Black Pearl Ring; finger2: Deep Fathom Ring; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 86.3. Weights run: 7.7s. Verify run: 4.4s. 315 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.448 ± 0.014, spirit=2.905 ± 0.018, mp5=4.429 ± 0.027, crit=0.046 ± 0.004 per rating point (14 rating = 1%, 0.641 per %), spell_haste=-1.040 ± 0.256

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Whisperwind Headdress (6688) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+9.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Enduring Cap (3020, -0.44 DPS) [world_drop]; Embalmed Shroud (7691, -1.32 DPS) [dungeon]; Holy Shroud (2721, -9.87 DPS, sim-verified) [world_drop] |
| neck | Necklace of Harmony (5180) | Singer [world] | 20.3 healing_power points (2.06 DPS) | yes | Pendant of Myzrael (4614, +0.00 DPS, sim-verified) [dungeon]; Glowing Green Talisman (5002, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.30 DPS) [world_drop] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Mantle (4718, -0.44 DPS) [world_drop]; Desert Shoulders (15457, -0.88 DPS) [quest]; Ghostly Mantle (3324, -7.52 DPS, sim-verified) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Amy's Blanket (13005, -0.44 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -0.44 DPS) [vendor]; Glowing Thresher Cape (6901, -0.75 DPS, sim-verified) [dungeon] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.1 healing_power points (4.48 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.36 DPS) [dungeon]; Pressed Felt Robe (1997, -0.50 DPS) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.5 healing_power points (2.09 DPS) | yes | Drakewing Bands (12999, +0.00 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.1 healing_power points (4.28 DPS) | yes | Tattered Mittens (270030, -0.49 DPS, sim-verified) [quest]; Hotshot Pilot's Gloves (9491, -1.62 DPS) [dungeon]; Blight Gloves (279877, -1.77 DPS) [quest] |
| waist | Silver-lined Belt (13011) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dokebi Cord (14578, -0.15 DPS) [world_drop]; Lilac Sash (6780, -0.29 DPS) [quest]; Mender's Leather Belt (252523, -2.25 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.5 healing_power points (5.23 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted]; Earthen Leggings (253999, -2.19 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.24 DPS) | yes | Boots of the Enchanter (4325, -1.88 DPS) [crafted]; Stonecloth Boots (14408, -1.88 DPS) [world_drop]; Soggy Boots (274747, -3.94 DPS, sim-verified) [vendor] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.5 healing_power points (3.09 DPS) | yes | Electrocutioner Lagnut (9447, -1.03 DPS) [dungeon]; Black Pearl Ring (6332, -1.03 DPS) [world]; The Queen's Jewel (13094, -4.03 DPS, sim-verified) [world_drop] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Black Pearl Ring (6332, -0.00 DPS) [world]; The Queen's Jewel (13094, -0.46 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -1.63 DPS, sim-verified) [dungeon]; Wind Spirit Staff (6689, -1.75 DPS) [dungeon]; Gnarled Ash Staff (791, -2.48 DPS) [world_drop] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.4 healing_power points (3.09 DPS) | yes | Defective Samophlange (274743, -0.58 DPS, sim-verified) [vendor]; Orb of Mistmantle (13031, -1.00 DPS) [world_drop]; Strength of Will (4837, -1.32 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Whisperwind Headdress; neck: Necklace of Harmony; shoulder: Batwing Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Tunic; wrist: Mindthrust Bracers; hands: Gloves of Old; waist: Silver-lined Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 131.7. Weights run: 8.6s. Verify run: 5.4s. 426 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.339 ± 0.016, spirit=2.500 ± 0.011, mp5=4.838 ± 0.034, crit=0.050 ± 0.005 per rating point (14 rating = 1%, 0.694 per %), spell_haste=not significant (0.513 ± 0.312)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 67.3 healing_power points (6.90 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -1.76 DPS) [dungeon]; Holy Shroud (2721, -1.98 DPS) [world_drop]; Whitemane's Chapeau (7720, -2.51 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 34.2 healing_power points (3.51 DPS) | yes | Triune Amulet (7722, -0.45 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.75 DPS) [dungeon]; Amberglow Talisman (10824, -0.94 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ghostly Mantle (3324, -0.51 DPS) [quest]; Crimson Silk Shoulders (7059, -0.84 DPS) [crafted]; Earthen Silk Shoulders (254033, -6.39 DPS, sim-verified) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 32.5 healing_power points (3.34 DPS) | yes | Battle Healer's Cloak (19528, +0.00 DPS, sim-verified) [rep]; Glowing Thresher Cape (6901, -0.56 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.75 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.0 healing_power points (8.00 DPS) | yes | Doomsayer's Robe (4746, -1.18 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -3.18 DPS) [crafted]; Black Mageweave Robe (10001, -3.31 DPS) [crafted] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.4 healing_power points (3.12 DPS) | yes | Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Dryad's Wrist Bindings (19597, -1.01 DPS) [pvp]; Earthen Silk Cuffs (254019, -1.48 DPS, sim-verified) [crafted] |
| hands | Gloves of Old (9395) | World drop [world_drop] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Gloves (252530, -0.37 DPS) [crafted]; Warden's Gloves (14606, -0.39 DPS) [world_drop]; Earthen Silk Gloves (254017, -1.76 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.7 healing_power points (4.90 DPS) | yes | Mender's Leather Belt (252523, -1.00 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 55.5 healing_power points (5.70 DPS) | yes | Wisdom's Leather Leggings (252519, -0.79 DPS) [crafted]; Stormcloth Pants (10010, -0.79 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -1.20 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 50.9 healing_power points (5.22 DPS) | yes | Furen's Boots (13100, +0.00 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.06 DPS) [dungeon]; Gilded Slippers (254001, -1.18 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.0 healing_power points (2.98 DPS) | yes | Welken Ring (5011, -0.50 DPS) [world_drop]; The Queen's Jewel (13094, -0.65 DPS) [world_drop]; Blush Ember Ring (13093, -0.93 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 26.9 healing_power points (2.76 DPS) | yes | The Queen's Jewel (13094, -0.43 DPS) [world_drop]; Blush Ember Ring (13093, -0.71 DPS) [world_drop]; Welken Ring (5011, -0.79 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -2.44 DPS) [dungeon]; Staff of Jordan (873, -2.64 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 32.0 healing_power points (3.28 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.46 DPS) [dungeon]; Silksand Star (15964, -0.97 DPS) [world_drop]; Orb of Souls (249395, -1.70 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Sheepshear Mantle; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; off_hand: Beacon of Hope

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 195.2. Weights run: 8.7s. Verify run: 5.1s. 561 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.522 ± 0.017, spirit=2.750 ± 0.016, mp5=4.393 ± 0.016, crit=0.069 ± 0.007 per rating point (14 rating = 1%, 0.967 per %), spell_haste=0.966 ± 0.177

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 75.4 healing_power points (7.47 DPS) | yes | Blood Guard's Restored Leather Helm (220875, +0.00 DPS) [vendor]; Cassandra's Grace (13102, -0.66 DPS) [world_drop]; Papal Fez (9431, -3.61 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkmoon Necklace (19303, -0.23 DPS) [vendor]; Horizon Choker (13085, -0.55 DPS) [world_drop]; Lei of Lilies (1315, -1.65 DPS, sim-verified) [world_drop] |
| shoulder | Mender's Leather Shoulder (252538) | Leatherworking [crafted] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Restored Leather Spaulders (220877, +0.00 DPS) [vendor]; Nethergeld Shoulders (254049, -0.23 DPS) [crafted]; Living Shoulders (15061, -4.68 DPS, sim-verified) [crafted] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | 47.3 healing_power points (4.69 DPS) | yes | Battle Healer's Cloak (19527, -0.94 DPS, sim-verified) [rep]; Cloak of Blight (6832, -1.15 DPS) [quest]; Darkspear Raider's Cloak (272076, -1.22 DPS) [vendor] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 108.4 healing_power points (10.74 DPS) | yes | Ghostweave Vest (14141, -1.81 DPS) [crafted]; Forest's Embrace (22272, -2.17 DPS) [quest]; Vestments of the Atal'ai Prophet (10806, -6.12 DPS, sim-verified) [dungeon] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.57 DPS) | yes | Nethergeld Cuffs (254061, -0.52 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -0.68 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.2 healing_power points (8.94 DPS) | yes | Gilded Gloves (254095, -2.38 DPS) [crafted]; Mender's Leather Gauntlets (252551, -3.06 DPS, sim-verified) [crafted]; Earthenweave Gloves (254075, -3.37 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Waistguard (252477, -0.24 DPS) [crafted]; Mender's Leather Belt (252523, -0.60 DPS) [crafted]; Earthenweave Cord (254077, -3.42 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Rainstrider Leggings (11123, -0.26 DPS) [quest]; Windscale Sarong (10842, -0.66 DPS) [world]; Dalewind Trousers (13008, -3.98 DPS, sim-verified) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.2 healing_power points (6.66 DPS) | yes | Mistwalker Boots (10629, -0.55 DPS) [dungeon]; Earthenweave Boots (254093, -1.08 DPS) [crafted]; First Sergeant's Restored Leather Boots (220885, -3.38 DPS, sim-verified) [vendor] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.1 healing_power points (4.18 DPS) | yes | Darkspear Signet (272069, -1.13 DPS) [vendor]; Cyclopean Band (11824, -1.14 DPS) [dungeon]; Snake Hoop (6750, -1.21 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 36.6 healing_power points (3.63 DPS) | yes | Cyclopean Band (11824, -0.59 DPS) [dungeon]; Snake Hoop (6750, -0.66 DPS) [quest]; Darkspear Signet (272069, -6.21 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -4.75 DPS) [world_drop]; Alchemists' Stone (13503, -5.44 DPS) [crafted]; Evonice's Landin' Pilla (18951, -6.03 DPS, sim-verified) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.55 DPS) [quest]; Alchemists' Stone (13503, -2.18 DPS) [crafted]; Uther's Strength (11302, -2.57 DPS, sim-verified) [world_drop] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -1.92 DPS) [world_drop]; Death Speaker Scepter (2816, -2.01 DPS) [dungeon]; Barman Shanker (12791, -10.84 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; shoulder: Mender's Leather Shoulder; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Soulkeeper

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 297.7. Weights run: 9.1s. Verify run: 16.9s. 1446 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-verified (193.8 DPS) | yes | Wildheart Cowl (16720, +0.00 DPS) [dungeon]; Champion's Dragonhide Headdress (227205, +0.00 DPS) [vendor]; Living Crown (252561, -8.93 DPS, sim-verified) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-verified (193.8 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.23 DPS) [world_drop]; Heart of the Fiend (13960, -1.00 DPS) [dungeon]; Wavefront Necklace (20685, -12.91 DPS, sim-verified) [world] |
| shoulder | Feralheart Mantle (226785) | Mokvar [vendor] | sim-verified (193.8 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Argent Elite Shoulders (227888, -29.18 DPS, sim-verified) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 78.1 healing_power points (7.88 DPS) | yes | Butcher's Apron (12608, -0.98 DPS) [dungeon]; Hide of the Wild (18510, -1.39 DPS) [crafted]; Featherskin Cape (10843, -18.59 DPS, sim-verified) [world] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-verified (193.8 DPS) | yes | Mooncloth Vest (14138, -1.45 DPS) [crafted]; Alanna's Embrace (13314, -1.62 DPS) [dungeon]; Tunic of Undead Slaying (23089, -36.39 DPS, sim-verified) [world] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (193.8 DPS) | yes | Bleak Howler Armguards (13208, +0.00 DPS) [dungeon]; Bracers of Hope (22667, +0.00 DPS, sim-verified) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-verified (193.8 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Wildheart Gloves (16717, +0.00 DPS) [dungeon]; Feralheart Gauntlets (226784, -2.80 DPS, sim-verified) [vendor] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | sim-verified (193.8 DPS) | yes | Elderwild Waistcord (279252, -0.72 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.34 DPS) [crafted]; Caretaker's Cord (272398, -32.72 DPS, sim-verified) [vendor] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 191.0 healing_power points (19.26 DPS) | yes | Haunting Specter Leggings (11929, -4.48 DPS) [dungeon]; Legionnaire's Dragonhide Legguards (227206, -4.55 DPS) [vendor]; Devout Skirt (16694, -25.98 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 148.3 healing_power points (14.96 DPS) | yes | Mooncloth Boots (15802, -3.78 DPS) [crafted]; Devout Sandals (16691, -4.26 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -22.27 DPS, sim-verified) [vendor] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-verified (193.8 DPS) | yes | Band of Mending (22334, -0.72 DPS) [dungeon]; Band of Piety (22681, -0.77 DPS) [quest]; Naglering (11669, -12.96 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (193.8 DPS) | yes | Band of Mending (22334, -0.38 DPS) [dungeon]; Band of Piety (22681, -0.43 DPS) [quest]; Naglering (11669, -34.50 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (193.8 DPS) | yes | Ankh of Life (1713, -5.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.07 DPS) [quest]; Mindtap Talisman (18371, -11.04 DPS, sim-verified) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-verified (193.8 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Serenity Field (272439, -0.37 DPS, sim-verified) [vendor] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-verified (193.8 DPS) | yes | Staff of Hale Magefire (13000, -0.09 DPS) [world_drop]; Soulkeeper (1607, -1.52 DPS) [world_drop]; Hand of Edward the Odd (2243, -13.70 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: The Postmaster's Band; neck: Lady Maye's Pendant; shoulder: Feralheart Mantle; back: Frostweaver Cape; wrist: Magister's Bindings; hands: Magister's Gloves; waist: Feralheart Cord; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: The Postmaster's Seal; finger2: Emerald Flame Ring; trinket2: Royal Seal of Eldre'Thalas; main_hand: Dancing Sliver

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 583.4. Weights run: 6.5s. Verify run: 13.3s. 1446 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.236, intellect=2.459 ± 0.082, spirit=2.390 ± 0.078, mp5=3.433 ± 0.102, crit=0.313 ± 0.025 per rating point (14 rating = 1%, 4.384 per %), spell_haste=not significant (-0.554 ± 1.262)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 150.7 healing_power points (24.33 DPS) | yes | Champion's Dragonhide Headdress (227205, -4.54 DPS) [vendor]; Sanctified Leather Helm (22689, -5.70 DPS) [quest]; Feralheart Headdress (226786, -24.59 DPS, sim-verified) [vendor] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 70.6 healing_power points (11.40 DPS) | yes | Wavefront Necklace (20685, -0.71 DPS) [world]; Amulet of the Redeemed (22327, -2.32 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -4.08 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (535.9 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Argent Elite Shoulders (227888, -31.39 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 66.6 healing_power points (10.75 DPS) | yes | Darkspear Raider's Cloak (272063, -2.08 DPS) [vendor]; Shroud of the Exile (15421, -2.09 DPS) [quest]; Frostweaver Cape (12968, -10.58 DPS, sim-verified) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (535.9 DPS) | yes | Alanna's Embrace (13314, +0.00 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -0.22 DPS) [world]; Mooncloth Vest (14138, -13.49 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (535.9 DPS) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Bracers of Mending (23129, +0.00 DPS) [dungeon]; Bracers of Hope (22667, -9.58 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (535.9 DPS) | yes | Feralheart Gauntlets (226784, +0.00 DPS) [vendor]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Hands of the Exalted Herald (12554, -12.37 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (535.9 DPS) | yes | General's Dragonhide Belt (16556, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Feralheart Cord (226780, -19.20 DPS, sim-verified) [vendor] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 138.9 healing_power points (22.42 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, -2.85 DPS) [vendor]; Feralheart Pants (226787, -4.12 DPS) [vendor]; Devout Skirt (16694, -9.77 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | sim-verified (535.9 DPS) | yes | Mooncloth Boots (15802, -2.77 DPS) [crafted]; Faith Healer's Boots (22247, -3.62 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -12.61 DPS, sim-verified) [vendor] |
| finger1 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (535.9 DPS) | yes | Band of Piety (22681, -0.53 DPS) [quest]; Seal of Rivendare (13345, -0.82 DPS) [dungeon]; Naglering (11669, -27.98 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (535.9 DPS) | yes | Band of Piety (22681, -0.36 DPS) [quest]; Seal of Rivendare (13345, -0.65 DPS) [dungeon]; Naglering (11669, -23.50 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (535.9 DPS) | yes | Royal Seal of Eldre'Thalas (18470, -1.21 DPS) [quest]; Serenity Field (272439, -1.53 DPS) [vendor]; Mindtap Talisman (18371, -2.22 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (535.9 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18470, +0.00 DPS) [quest]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-verified (535.9 DPS) | yes | Hammer of the Grand Crusader (18717, -1.07 DPS) [dungeon]; Staff of Metanoia (22394, -1.79 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Maye's Pendant; shoulder: Devout Mantle; back: Hide of the Wild; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: Emerald Flame Ring; finger2: Band of Mending; trinket2: Blackhand's Breadth; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

