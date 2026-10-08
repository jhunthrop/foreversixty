# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 44.8. Weights run: 4.7s. Verify run: 2.3s. 193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.220 ± 0.005, spirit=1.493 ± 0.008, mp5=3.165 ± 0.014, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.529 per %), spell_haste=-5.534 ± 0.104

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.09 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.30 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.47 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 6.0 healing_power points (0.61 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Mantle (4709, -0.36 DPS) [world_drop]; Prospector's Pads (14566, -0.36 DPS) [world_drop]; Magician's Mantle (12998, -1.03 DPS, sim-verified) [world_drop] |
| back | Regent's Cloak (5969) | Ravenclaw Regent [world] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Spirit Cloak (4792, -0.15 DPS) [vendor]; Sylvan Cloak (4793, -0.15 DPS) [vendor]; Caretaker's Cape (20428, -2.33 DPS, sim-verified) [rep] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, -0.04 DPS) [crafted]; Corsair's Overshirt (5202, -0.10 DPS) [dungeon]; Wisdom's Leather Armor (252493, -1.92 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 15.6 healing_power points (1.60 DPS) | yes | Owl Bracers (4796, -0.98 DPS) [vendor]; Bravo's Armbands (270015, -0.99 DPS) [quest]; Drakewing Bands (12999, -2.65 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Magefist Gloves (12977, -0.27 DPS) [world_drop]; Bright Gloves (3066, -0.39 DPS) [world_drop]; Wisdom's Leather Gloves (252499, -2.41 DPS, sim-verified) [crafted] |
| waist | Novice Ardent's Sash (253887) | Tailoring [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.45 DPS) [world_drop]; Pristine Sash (253925, -2.20 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.3 healing_power points (3.42 DPS) | yes | Filigreed Pristine Leggings (253937, -0.18 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.62 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Black Whelp Slippers (252424) | Leatherworking [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Boots (253889, -0.13 DPS) [crafted]; Kimbra Boots (6191, -0.46 DPS) [quest]; Wisdom's Leather Boots (252444, -2.11 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.4 healing_power points (1.17 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.52 DPS) [rep] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.01 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.12 DPS) [rep]; Band of Purification (12996, -2.50 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 15.1 healing_power points (1.55 DPS) | yes | Staff of the Blessed Seer (2271, -0.18 DPS, sim-verified) [dungeon]; Twisted Chanter's Staff (890, -0.29 DPS) [world_drop]; Gnarled Hermit's Staff (1539, -0.47 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Slime-encrusted Pads; back: Regent's Cloak; chest: Robe of the Moccasin; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Novice Ardent's Sash; legs: Wisdom's Leather Pants; feet: Black Whelp Slippers; finger1: Black Pearl Ring; finger2: Deep Fathom Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 85.1. Weights run: 6.5s. Verify run: 2.9s. 322 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.009, intellect=1.453 ± 0.012, spirit=2.927 ± 0.024, mp5=4.581 ± 0.031, crit=0.042 ± 0.004 per rating point (14 rating = 1%, 0.595 per %), spell_haste=not significant (0.262 ± 0.224)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Whisperwind Headdress (6688) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (48.7 DPS) | yes | Enduring Cap (3020, -0.45 DPS) [world_drop]; Embalmed Shroud (7691, -1.31 DPS) [dungeon]; Holy Shroud (2721, -2.97 DPS, sim-verified) [world_drop] |
| neck | Necklace of Harmony (5180) | Singer [world] | 20.5 healing_power points (2.06 DPS) | yes | Glowing Green Talisman (5002, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.30 DPS) [world_drop]; Pendant of Myzrael (4614, -1.71 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 30.7 healing_power points (3.08 DPS) | yes | Faerie Mantle (5820, -0.44 DPS) [quest]; Nightsky Mantle (4718, -0.44 DPS) [world_drop]; Batwing Mantle (6697, -1.09 DPS, sim-verified) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.4 healing_power points (3.05 DPS) | yes | Repairman's Cape (9605, -0.41 DPS) [quest]; Caretaker's Cape (19533, -0.57 DPS) [rep]; Prelacy Cape (7004, -1.86 DPS, sim-verified) [quest] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.3 healing_power points (4.44 DPS) | yes | Beguiler Robes (7728, -0.34 DPS) [dungeon]; Pressed Felt Robe (1997, -0.48 DPS) [world]; Pristine Gown (253961, -2.22 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.0 healing_power points (2.11 DPS) | yes | Nightsky Wristbands (6407, -0.35 DPS) [world_drop]; Dokebi Bracers (14580, -0.50 DPS) [world_drop]; Drakewing Bands (12999, -2.80 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.4 healing_power points (4.25 DPS) | yes | Zodiac Gloves (7106, -1.61 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.62 DPS) [dungeon]; Shilly Mitts (9609, -4.38 DPS, sim-verified) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.3 healing_power points (4.54 DPS) | yes | Dokebi Cord (14578, -2.20 DPS) [world_drop]; Resilient Cord (14406, -2.49 DPS) [world_drop]; Silver-lined Belt (13011, -4.16 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.7 healing_power points (5.18 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.93 DPS) [crafted]; Earthen Leggings (253999, -1.76 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.9 healing_power points (4.20 DPS) | yes | Soggy Boots (274747, -0.14 DPS, sim-verified) [vendor]; Boots of the Enchanter (4325, -1.85 DPS) [crafted]; Stonecloth Boots (14408, -1.85 DPS) [world_drop] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.7 healing_power points (3.08 DPS) | yes | Monkey Ring (6748, -1.02 DPS) [quest]; Ring of Calm (6790, -1.02 DPS) [quest]; Electrocutioner Lagnut (9447, -1.02 DPS) [dungeon] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 26.3 healing_power points (2.64 DPS) | yes | Monkey Ring (6748, -0.59 DPS) [quest]; Electrocutioner Lagnut (9447, -0.59 DPS) [dungeon]; Ring of Calm (6790, -2.93 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gnarled Ash Staff (791, +0.00 DPS) [world_drop]; Death Speaker Scepter (2816, +0.00 DPS) [dungeon]; Wind Spirit Staff (6689, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Whisperwind Headdress; neck: Necklace of Harmony; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 134.5. Weights run: 7.3s. Verify run: 3.2s. 438 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.335 ± 0.008, spirit=2.460 ± 0.012, mp5=4.827 ± 0.037, crit=0.046 ± 0.005 per rating point (14 rating = 1%, 0.644 per %), spell_haste=not significant (0.357 ± 0.200)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 66.8 healing_power points (6.84 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -1.77 DPS) [dungeon]; Whitemane's Chapeau (7720, -1.86 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -1.95 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.7 healing_power points (3.45 DPS) | yes | Triune Amulet (7722, -0.70 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.73 DPS) [dungeon]; Amberglow Talisman (10824, -0.94 DPS) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 37.7 healing_power points (3.86 DPS) | yes | Sheepshear Mantle (13115, +0.00 DPS, sim-verified) [world_drop]; Crimson Silk Shoulders (7059, -1.00 DPS) [crafted]; Mantle of Doan (7712, -1.00 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | 30.3 healing_power points (3.10 DPS) | yes | Prelacy Cape (7004, -0.46 DPS) [quest]; Ceremonial Centaur Blanket (6789, -0.56 DPS) [quest]; Glowing Thresher Cape (6901, -1.68 DPS, sim-verified) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 77.7 healing_power points (7.96 DPS) | yes | Doomsayer's Robe (4746, -2.61 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -3.20 DPS) [crafted]; Black Mageweave Robe (10001, -3.33 DPS) [crafted] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 29.9 healing_power points (3.07 DPS) | yes | Earthen Silk Cuffs (254019, -0.53 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.90 DPS) [dungeon]; Silkstream Cuffs (16791, -1.37 DPS) [quest] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 37.7 healing_power points (3.86 DPS) | yes | Mender's Leather Gloves (252530, -0.48 DPS) [crafted]; Warden's Gloves (14606, -0.54 DPS) [world_drop]; Gloves of Old (9395, -0.97 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.4 healing_power points (4.86 DPS) | yes | Mender's Leather Belt (252523, -0.46 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.50 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 54.8 healing_power points (5.61 DPS) | yes | Stormcloth Pants (10010, -0.29 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.74 DPS) [crafted]; Pristine Leggings (253987, -1.14 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 50.6 healing_power points (5.18 DPS) | yes | Thoughtcast Boots (10578, -1.08 DPS) [dungeon]; Furen's Boots (13100, -1.12 DPS, sim-verified) [world_drop]; Gilded Slippers (254001, -1.17 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.0 healing_power points (2.97 DPS) | yes | Welken Ring (5011, -0.52 DPS) [world_drop]; The Queen's Jewel (13094, -0.68 DPS) [world_drop]; Blush Ember Ring (13093, -0.95 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 26.6 healing_power points (2.72 DPS) | yes | Welken Ring (5011, -0.18 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.43 DPS) [world_drop]; Blush Ember Ring (13093, -0.70 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (68.1 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (68.1 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -2.50 DPS) [dungeon]; Staff of Jordan (873, -2.69 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 31.8 healing_power points (3.26 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.49 DPS) [dungeon]; Orb of Souls (249395, -0.93 DPS, sim-verified) [crafted]; Silksand Star (15964, -0.99 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; main_hand: Death Speaker Scepter; off_hand: Beacon of Hope

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 196.0. Weights run: 7.4s. Verify run: 3.0s. 577 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.458 ± 0.018, spirit=3.178 ± 0.016, mp5=4.502 ± 0.016, crit=0.061 ± 0.007 per rating point (14 rating = 1%, 0.858 per %), spell_haste=not significant (0.153 ± 0.162)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Restored Leather Helm (220874) | Captain Dirgehammer [vendor] | 84.3 healing_power points (8.36 DPS) | yes | Gemburst Circlet (10751, -0.62 DPS, sim-verified) [quest]; Papal Fez (9431, -0.88 DPS) [dungeon]; Cassandra's Grace (13102, -1.16 DPS) [world_drop] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 47.7 healing_power points (4.73 DPS) | yes | Glowing Eye of Mordresh (10769, -0.54 DPS) [dungeon]; Darkmoon Necklace (19303, -1.18 DPS) [vendor]; Gemshard Heart (17707, -1.39 DPS) [dungeon] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 72.3 healing_power points (7.18 DPS) | yes | Mender's Leather Shoulder (252538, -1.29 DPS) [crafted]; Nethergeld Shoulders (254049, -1.55 DPS) [crafted]; Knight-Lieutenant's Restored Leather Spaulders (220876, -4.37 DPS, sim-verified) [vendor] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | 53.5 healing_power points (5.31 DPS) | yes | Arcane Cloak (8286, -1.53 DPS) [world_drop]; Caretaker's Cape (19531, -1.58 DPS, sim-verified) [rep]; Darkspear Raider's Cloak (272076, -1.71 DPS) [vendor] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 120.1 healing_power points (11.92 DPS) | yes | Ghostweave Vest (14141, -2.04 DPS) [crafted]; Feathered Breastplate (8349, -2.91 DPS) [crafted]; Vestments of the Atal'ai Prophet (10806, -3.06 DPS, sim-verified) [dungeon] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 48.3 healing_power points (4.79 DPS) | yes | Nethergeld Cuffs (254061, -0.22 DPS) [crafted]; Aristocratic Cuffs (12546, -0.73 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.06 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 97.2 healing_power points (9.65 DPS) | yes | Mender's Leather Gauntlets (252551, +0.00 DPS, sim-verified) [crafted]; Gilded Gloves (254095, -2.80 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -3.34 DPS) [dungeon] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 61.0 healing_power points (6.05 DPS) | yes | Gilded Cord (254037, -0.78 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -1.34 DPS) [crafted]; Mender's Leather Belt (252523, -1.40 DPS) [crafted] |
| legs | Dalewind Trousers (13008) | World drop [world_drop] | 92.6 healing_power points (9.19 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.91 DPS) [dungeon]; Windscale Sarong (10842, -1.43 DPS) [world]; Rainstrider Leggings (11123, -2.52 DPS, sim-verified) [quest] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 75.2 healing_power points (7.47 DPS) | yes | Sergeant Major's Restored Leather Boots (220884, -0.88 DPS) [vendor]; Mistwalker Boots (10629, -1.17 DPS, sim-verified) [dungeon]; Furen's Boots (13100, -1.41 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 46.9 healing_power points (4.65 DPS) | yes | Brainlash (6440, -0.91 DPS) [dungeon]; Coldwater Ring (4550, -1.38 DPS) [quest]; Snake Hoop (6750, -1.43 DPS) [quest] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 41.3 healing_power points (4.10 DPS) | yes | Brainlash (6440, -0.35 DPS) [dungeon]; Coldwater Ring (4550, -0.83 DPS) [quest]; Snake Hoop (6750, -0.88 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (105.4 DPS) | yes | Ankh of Life (1713, -2.92 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -4.18 DPS) [quest]; Uther's Strength (11302, -4.92 DPS) [world_drop] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (105.4 DPS) | yes | Ankh of Life (1713, -0.57 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -0.63 DPS) [quest]; Uther's Strength (11302, -1.37 DPS) [world_drop] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-verified (105.4 DPS) | yes | Glowing Brightwood Staff (812, -2.76 DPS) [world_drop]; Death Speaker Scepter (2816, -3.04 DPS) [dungeon]; Barman Shanker (12791, -10.66 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Restored Leather Helm; neck: Lei of Lilies; shoulder: Living Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Dalewind Trousers; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket1: Darkspear Voodoo Seal; trinket2: Evonice's Landin' Pilla; main_hand: Soulkeeper

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 317.8. Weights run: 7.5s. Verify run: 3.2s. 1449 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.251 ± 0.040, spirit=4.529 ± 0.030, mp5=6.786 ± 0.032, crit=0.094 ± 0.009 per rating point (14 rating = 1%, 1.309 per %), spell_haste=not significant (0.911 ± 0.553)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 176.9 healing_power points (18.21 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -2.70 DPS) [vendor]; Feralheart Headdress (226786, -3.28 DPS) [vendor]; Wildheart Cowl (16720, -8.42 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 91.8 healing_power points (9.45 DPS) | yes | Lady Maye's Pendant (14558, +0.00 DPS, sim-verified) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.62 DPS) [world_drop]; Heart of the Fiend (13960, -1.30 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 145.0 healing_power points (14.93 DPS) | yes | Feralheart Mantle (226785, +0.00 DPS, sim-verified) [vendor]; Lieutenant Commander's Dragonhide Pauldrons (227201, -2.20 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -3.37 DPS) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 81.4 healing_power points (8.37 DPS) | yes | Butcher's Apron (12608, -0.91 DPS) [dungeon]; Featherskin Cape (10843, -1.43 DPS, sim-verified) [world]; Shroud of the Exile (15421, -1.64 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-verified (187.7 DPS) | yes | Mooncloth Vest (14138, -1.76 DPS) [crafted]; Alanna's Embrace (13314, -1.91 DPS) [dungeon]; Tunic of Undead Slaying (23089, -16.07 DPS, sim-verified) [world] |
| wrist | Bracers of Mending (23129) | Dire Maul: Revanchion [dungeon] | sim-verified (187.7 DPS) | yes | Bracers of Hope (22667, -0.00 DPS) [quest]; Bleak Howler Armguards (13208, -0.07 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.15 DPS, sim-verified) [world] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 130.3 healing_power points (13.41 DPS) | yes | Wildheart Gloves (16717, -0.62 DPS, sim-verified) [dungeon]; Hands of the Exalted Herald (12554, -1.41 DPS) [dungeon]; Devout Gloves (16692, -1.62 DPS) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | 131.7 healing_power points (13.56 DPS) | yes | Feralheart Cord (226780, -1.33 DPS, sim-verified) [vendor]; Elderwild Waistcord (279252, -2.08 DPS) [crafted]; Wisdom of the Timbermaw (19047, -2.93 DPS) [crafted] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 198.9 healing_power points (20.47 DPS) | yes | Haunting Specter Leggings (11929, -4.64 DPS) [dungeon]; Knight-Captain's Dragonhide Legguards (227200, -5.04 DPS) [vendor]; Devout Skirt (16694, -7.46 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 153.4 healing_power points (15.79 DPS) | yes | Mooncloth Boots (15802, -4.02 DPS) [crafted]; Devout Sandals (16691, -4.42 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -4.49 DPS, sim-verified) [vendor] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-verified (187.7 DPS) | yes | Band of Mending (22334, -1.01 DPS) [dungeon]; Band of the Hierophant (13096, -1.18 DPS) [world_drop]; Naglering (11669, -5.34 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (187.7 DPS) | yes | Band of Mending (22334, -0.44 DPS) [dungeon]; Band of the Hierophant (13096, -0.61 DPS) [world_drop]; Naglering (11669, -7.63 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (187.7 DPS) | yes | Mindtap Talisman (18371, -2.79 DPS) [dungeon]; Ankh of Life (1713, -4.88 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -5.81 DPS) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-verified (187.7 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Serenity Field (272439, -0.65 DPS, sim-verified) [vendor] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-verified (187.7 DPS) | yes | Staff of Hale Magefire (13000, -0.25 DPS) [world_drop]; Soulkeeper (1607, -1.63 DPS) [world_drop]; Hand of Edward the Odd (2243, -5.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Frostweaver Cape; wrist: Bracers of Mending; waist: Caretaker's Cord; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: The Postmaster's Seal; finger2: Emerald Flame Ring; trinket2: Royal Seal of Eldre'Thalas; main_hand: Dancing Sliver

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 509.9. Weights run: 5.4s. Verify run: 2.6s. 1449 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.266, intellect=1.074 ± 0.045, spirit=2.070 ± 0.045, mp5=3.507 ± 0.053, crit=0.196 ± 0.023 per rating point (14 rating = 1%, 2.749 per %), spell_haste=not significant (2.112 ± 0.836)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 121.3 healing_power points (23.43 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -5.76 DPS) [vendor]; Feralheart Headdress (226786, -6.60 DPS) [vendor]; Sanctified Leather Helm (22689, -6.96 DPS) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 58.5 healing_power points (11.30 DPS) | yes | Animated Chain Necklace (18723, -2.53 DPS) [dungeon]; Amulet of the Redeemed (22327, -3.34 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.36 DPS) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 100.7 healing_power points (19.45 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, -5.17 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -7.06 DPS) [vendor]; Feralheart Mantle (226785, -7.71 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.7 healing_power points (10.18 DPS) | yes | Drape of Recovery (272413, -1.49 DPS) [vendor]; Caretaker's Cape (19530, -1.97 DPS) [rep]; Cloak of the Cosmos (18389, -2.88 DPS) [dungeon] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | sim-verified (510.0 DPS) | yes | Tunic of Undead Slaying (23089, -0.85 DPS, sim-verified) [world]; Mooncloth Vest (14138, -1.80 DPS) [crafted]; Feralheart Embrace (226783, -2.11 DPS) [vendor] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (510.0 DPS) | yes | Wristwraps of Undead Slaying (23093, +0.00 DPS) [world]; Bracers of Mending (23129, -0.02 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.40 DPS) [dungeon] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 73.1 healing_power points (14.11 DPS) | yes | Hands of the Exalted Herald (12554, -0.25 DPS) [dungeon]; Wildheart Gloves (16717, -1.73 DPS) [dungeon]; Devout Gloves (16692, -2.34 DPS) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 73.7 healing_power points (14.23 DPS) | yes | Caretaker's Cord (272398, -0.20 DPS) [vendor]; Elderwild Waistcord (279252, -1.24 DPS) [crafted]; Feralheart Cord (226780, -1.52 DPS) [vendor] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 101.6 healing_power points (19.61 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, -1.42 DPS) [vendor]; Elderwild Pants (279254, -2.30 DPS) [crafted]; Devout Skirt (16694, -2.68 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 88.0 healing_power points (17.00 DPS) | yes | Feralheart Sandals (226781, -1.09 DPS) [vendor]; Mooncloth Boots (15802, -4.26 DPS) [crafted]; Faith Healer's Boots (22247, -4.69 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (510.0 DPS) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Band of Piety (22681, -0.80 DPS) [quest]; Rosewine Circle (13178, -0.85 DPS) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (510.0 DPS) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Band of Piety (22681, -0.15 DPS) [quest]; Rosewine Circle (13178, -0.20 DPS) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (510.0 DPS) | yes | Royal Seal of Eldre'Thalas (18470, +0.00 DPS) [quest]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Ankh of Life (1713, -1.02 DPS, sim-verified) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (510.0 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Staff of Metanoia (22394, -2.47 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -3.31 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Bracers of Hope; waist: Sash of Mercy; legs: Leggings of Arcana; feet: Incandescent Mooncloth Boots; finger1: Band of Mending; finger2: Blessed Band of Light; trinket1: Draconic Infused Emblem; main_hand: Redemption

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 43.9. Weights run: 4.7s. Verify run: 2.3s. 183 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.220 ± 0.005, spirit=1.493 ± 0.008, mp5=3.165 ± 0.014, crit=0.038 ± 0.003 per rating point (14 rating = 1%, 0.529 per %), spell_haste=-5.534 ± 0.104

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.10 DPS) [crafted]; Shadow Goggles (4373, -0.30 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.47 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Scholarly Pendant (277203, -2.26 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.0 healing_power points (1.13 DPS) | yes | Forest Leather Mantle (4709, -0.51 DPS) [world_drop]; Prospector's Pads (14566, -0.51 DPS) [world_drop]; Slime-encrusted Pads (6461, -0.93 DPS, sim-verified) [dungeon] |
| back | Regent's Cloak (5969) | Ravenclaw Regent [world] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Traveler's Shawl (277289, +0.00 DPS) [quest]; Spirit Cloak (4792, -0.15 DPS) [vendor]; Battle Healer's Cloak (20427, -4.68 DPS, sim-verified) [rep] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 25.6 healing_power points (2.63 DPS) | yes | Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.27 DPS) [dungeon]; Robe of the Moccasin (6465, -1.33 DPS, sim-verified) [dungeon] |
| wrist | Drakewing Bands (12999) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Tabitha's Cuffs (251486, -0.17 DPS) [quest]; Crystalline Cuffs (14148, -0.21 DPS) [dungeon]; Mindthrust Bracers (1974, -0.53 DPS, sim-verified) [dungeon] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gloves (253913, -0.02 DPS) [crafted]; Magefist Gloves (12977, -0.29 DPS) [world_drop]; Blight Gloves (279877, -1.90 DPS, sim-verified) [quest] |
| waist | Novice Ardent's Sash (253887) | Tailoring [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.45 DPS) [world_drop]; Pristine Sash (253925, -2.65 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.3 healing_power points (3.42 DPS) | yes | Filigreed Pristine Leggings (253937, -0.12 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.62 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Black Whelp Slippers (252424) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Boots (253889, -0.13 DPS) [crafted]; Smoldering Boots (3076, -0.51 DPS) [world]; Wisdom's Leather Boots (252444, -0.62 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.4 healing_power points (1.17 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; Advisor's Ring (20426, -0.52 DPS) [rep] |
| finger2 | Deep Fathom Ring (6463) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.01 DPS) [dungeon]; Advisor's Ring (20426, -0.12 DPS) [rep]; Band of Purification (12996, -2.26 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Gnarled Staff (20425, -0.06 DPS) [pvp]; Gnarled Necromancer's Staff (251534, -0.28 DPS) [quest]; Staff of Orgrimmar (15444, -2.77 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Regent's Cloak; chest: Wisdom's Leather Armor; wrist: Drakewing Bands; hands: Wisdom's Leather Gloves; waist: Novice Ardent's Sash; legs: Wisdom's Leather Pants; feet: Black Whelp Slippers; finger1: Black Pearl Ring; finger2: Deep Fathom Ring; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 85.9. Weights run: 6.5s. Verify run: 3.1s. 315 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.009, intellect=1.453 ± 0.012, spirit=2.927 ± 0.024, mp5=4.581 ± 0.031, crit=0.042 ± 0.004 per rating point (14 rating = 1%, 0.595 per %), spell_haste=not significant (0.262 ± 0.224)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Whisperwind Headdress (6688) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Enduring Cap (3020, -0.45 DPS) [world_drop]; Embalmed Shroud (7691, -1.31 DPS) [dungeon]; Holy Shroud (2721, -4.14 DPS, sim-verified) [world_drop] |
| neck | Necklace of Harmony (5180) | Singer [world] | 20.5 healing_power points (2.06 DPS) | yes | Pendant of Myzrael (4614, -0.15 DPS, sim-verified) [dungeon]; Glowing Green Talisman (5002, -0.29 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.30 DPS) [world_drop] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Mantle (4718, -0.44 DPS) [world_drop]; Desert Shoulders (15457, -0.87 DPS) [quest]; Ghostly Mantle (3324, -2.20 DPS, sim-verified) [quest] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Amy's Blanket (13005, -0.42 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -0.43 DPS) [vendor]; Glowing Thresher Cape (6901, -0.84 DPS, sim-verified) [dungeon] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.3 healing_power points (4.44 DPS) | yes | Beguiler Robes (7728, -0.34 DPS) [dungeon]; Pressed Felt Robe (1997, -0.48 DPS) [world]; Pristine Gown (253961, -0.62 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 21.0 healing_power points (2.11 DPS) | yes | Nightsky Wristbands (6407, -0.35 DPS) [world_drop]; Dokebi Bracers (14580, -0.50 DPS) [world_drop]; Drakewing Bands (12999, -1.48 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.4 healing_power points (4.25 DPS) | yes | Hotshot Pilot's Gloves (9491, -1.62 DPS) [dungeon]; Blight Gloves (279877, -1.76 DPS) [quest]; Tattered Mittens (270030, -1.90 DPS, sim-verified) [quest] |
| waist | Silver-lined Belt (13011) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dokebi Cord (14578, -0.14 DPS) [world_drop]; Lilac Sash (6780, -0.29 DPS) [quest]; Mender's Leather Belt (252523, -1.28 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.7 healing_power points (5.18 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.93 DPS) [crafted]; Earthen Leggings (253999, -2.63 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.9 healing_power points (4.20 DPS) | yes | Boots of the Enchanter (4325, -1.85 DPS) [crafted]; Stonecloth Boots (14408, -1.85 DPS) [world_drop]; Soggy Boots (274747, -2.48 DPS, sim-verified) [vendor] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.7 healing_power points (3.08 DPS) | yes | Monkey Ring (6748, -1.02 DPS) [quest]; Electrocutioner Lagnut (9447, -1.02 DPS) [dungeon]; Black Pearl Ring (6332, -1.02 DPS) [world] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 26.3 healing_power points (2.64 DPS) | yes | Electrocutioner Lagnut (9447, -0.59 DPS) [dungeon]; Black Pearl Ring (6332, -0.59 DPS) [world]; Monkey Ring (6748, -0.93 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, -1.49 DPS, sim-verified) [dungeon]; Wind Spirit Staff (6689, -1.69 DPS) [dungeon]; Gnarled Ash Staff (791, -2.42 DPS) [world_drop] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 30.6 healing_power points (3.07 DPS) | yes | Defective Samophlange (274743, -0.26 DPS, sim-verified) [vendor]; Orb of Mistmantle (13031, -0.99 DPS) [world_drop]; Strength of Will (4837, -1.30 DPS) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Whisperwind Headdress; neck: Necklace of Harmony; shoulder: Batwing Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Tunic; wrist: Mindthrust Bracers; hands: Gloves of Old; waist: Silver-lined Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 133.4. Weights run: 7.3s. Verify run: 3.3s. 426 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.335 ± 0.008, spirit=2.460 ± 0.012, mp5=4.827 ± 0.037, crit=0.046 ± 0.005 per rating point (14 rating = 1%, 0.644 per %), spell_haste=not significant (0.357 ± 0.200)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 66.8 healing_power points (6.84 DPS) | yes | Whitemane's Chapeau (7720, -1.63 DPS, sim-verified) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.77 DPS) [dungeon]; Holy Shroud (2721, -1.95 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.7 healing_power points (3.45 DPS) | yes | Triune Amulet (7722, -0.62 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.73 DPS) [dungeon]; Amberglow Talisman (10824, -0.94 DPS) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 37.7 healing_power points (3.86 DPS) | yes | Sheepshear Mantle (13115, +0.00 DPS, sim-verified) [world_drop]; Ghostly Mantle (3324, -0.67 DPS) [quest]; Crimson Silk Shoulders (7059, -1.00 DPS) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 32.0 healing_power points (3.28 DPS) | yes | Battle Healer's Cloak (19528, +0.00 DPS, sim-verified) [rep]; Glowing Thresher Cape (6901, -0.54 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.73 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 77.7 healing_power points (7.96 DPS) | yes | Dreamweave Vest (10021, -3.20 DPS) [crafted]; Black Mageweave Robe (10001, -3.33 DPS) [crafted]; Doomsayer's Robe (4746, -3.54 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 29.9 healing_power points (3.07 DPS) | yes | Earthen Silk Cuffs (254019, -0.43 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.90 DPS) [dungeon]; Dryad's Wrist Bindings (19597, -0.99 DPS) [pvp] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 37.7 healing_power points (3.86 DPS) | yes | Mender's Leather Gloves (252530, -0.48 DPS) [crafted]; Warden's Gloves (14606, -0.54 DPS) [world_drop]; Gloves of Old (9395, -0.64 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.4 healing_power points (4.86 DPS) | yes | Mender's Leather Belt (252523, -0.12 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.50 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 54.8 healing_power points (5.61 DPS) | yes | Stormcloth Pants (10010, -0.45 DPS, sim-verified) [crafted]; Wisdom's Leather Leggings (252519, -0.74 DPS) [crafted]; Pristine Leggings (253987, -1.14 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 50.6 healing_power points (5.18 DPS) | yes | Furen's Boots (13100, -0.42 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -1.08 DPS) [dungeon]; Gilded Slippers (254001, -1.17 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.0 healing_power points (2.97 DPS) | yes | Welken Ring (5011, -0.52 DPS) [world_drop]; The Queen's Jewel (13094, -0.68 DPS) [world_drop]; Blush Ember Ring (13093, -0.95 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 26.6 healing_power points (2.72 DPS) | yes | Welken Ring (5011, +0.00 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.43 DPS) [world_drop]; Blush Ember Ring (13093, -0.70 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (67.0 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (67.0 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -2.50 DPS) [dungeon]; Staff of Jordan (873, -2.69 DPS) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman: Shadowforge Relic Hunter [dungeon] | 31.8 healing_power points (3.26 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.49 DPS) [dungeon]; Orb of Souls (249395, -0.85 DPS, sim-verified) [crafted]; Silksand Star (15964, -0.99 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; off_hand: Beacon of Hope

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 196.7. Weights run: 7.4s. Verify run: 3.0s. 561 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.458 ± 0.018, spirit=3.178 ± 0.016, mp5=4.502 ± 0.016, crit=0.061 ± 0.007 per rating point (14 rating = 1%, 0.858 per %), spell_haste=not significant (0.153 ± 0.162)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 82.0 healing_power points (8.14 DPS) | yes | Papal Fez (9431, +0.00 DPS, sim-verified) [dungeon]; Blood Guard's Restored Leather Helm (220875, +0.00 DPS) [vendor]; Cassandra's Grace (13102, -0.93 DPS) [world_drop] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | 47.7 healing_power points (4.73 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Darkmoon Necklace (19303, -1.18 DPS) [vendor]; Gemshard Heart (17707, -1.39 DPS) [dungeon] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 72.3 healing_power points (7.18 DPS) | yes | Mender's Leather Shoulder (252538, +0.00 DPS, sim-verified) [crafted]; Blood Guard's Restored Leather Spaulders (220877, -0.39 DPS) [vendor]; Nethergeld Shoulders (254049, -1.55 DPS) [crafted] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | 53.5 healing_power points (5.31 DPS) | yes | Battle Healer's Cloak (19527, -1.23 DPS) [rep]; Arcane Cloak (8286, -1.53 DPS) [world_drop]; Cloak of Blight (6832, -2.30 DPS, sim-verified) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 120.1 healing_power points (11.92 DPS) | yes | Vestments of the Atal'ai Prophet (10806, -1.50 DPS, sim-verified) [dungeon]; Ghostweave Vest (14141, -2.04 DPS) [crafted]; Feathered Breastplate (8349, -2.91 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 48.3 healing_power points (4.79 DPS) | yes | Nethergeld Cuffs (254061, -0.40 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -0.73 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.06 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 97.2 healing_power points (9.65 DPS) | yes | Mender's Leather Gauntlets (252551, +0.00 DPS, sim-verified) [crafted]; Gilded Gloves (254095, -2.80 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -3.34 DPS) [dungeon] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 61.0 healing_power points (6.05 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -1.34 DPS) [crafted]; Mender's Leather Belt (252523, -1.40 DPS) [crafted] |
| legs | Dalewind Trousers (13008) | World drop [world_drop] | 92.6 healing_power points (9.19 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.91 DPS) [dungeon]; Windscale Sarong (10842, -1.43 DPS) [world]; Rainstrider Leggings (11123, -1.69 DPS, sim-verified) [quest] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 75.2 healing_power points (7.47 DPS) | yes | Mistwalker Boots (10629, -0.63 DPS) [dungeon]; First Sergeant's Restored Leather Boots (220885, -0.88 DPS) [vendor]; Furen's Boots (13100, -1.41 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 46.9 healing_power points (4.65 DPS) | yes | Coldwater Ring (4550, -1.38 DPS) [quest]; Snake Hoop (6750, -1.43 DPS) [quest]; Cyclopean Band (11824, -1.49 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.8 healing_power points (3.75 DPS) | yes | Snake Hoop (6750, -0.53 DPS) [quest]; Cyclopean Band (11824, -0.58 DPS) [dungeon]; Coldwater Ring (4550, -2.13 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (106.3 DPS) | yes | Ankh of Life (1713, -2.92 DPS) [world_drop]; Uther's Strength (11302, -4.92 DPS) [world_drop]; Alchemists' Stone (13503, -5.44 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (106.3 DPS) | yes | Ankh of Life (1713, +0.00 DPS, sim-verified) [world_drop]; Uther's Strength (11302, -1.37 DPS) [world_drop]; Alchemists' Stone (13503, -1.89 DPS) [crafted] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-verified (106.3 DPS) | yes | Glowing Brightwood Staff (812, -2.76 DPS) [world_drop]; Death Speaker Scepter (2816, -3.04 DPS) [dungeon]; Barman Shanker (12791, -5.56 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Lei of Lilies; shoulder: Living Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Dalewind Trousers; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; trinket2: Evonice's Landin' Pilla; main_hand: Soulkeeper

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 316.6. Weights run: 7.5s. Verify run: 3.5s. 1446 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.011, intellect=2.251 ± 0.040, spirit=4.529 ± 0.030, mp5=6.786 ± 0.032, crit=0.094 ± 0.009 per rating point (14 rating = 1%, 1.309 per %), spell_haste=not significant (0.911 ± 0.553)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 176.9 healing_power points (18.21 DPS) | yes | Champion's Dragonhide Headdress (227205, -2.70 DPS) [vendor]; Feralheart Headdress (226786, -3.28 DPS) [vendor]; Wildheart Cowl (16720, -6.74 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 91.8 healing_power points (9.45 DPS) | yes | Lady Maye's Pendant (14558, -0.38 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.62 DPS) [world_drop]; Heart of the Fiend (13960, -1.30 DPS) [dungeon] |
| shoulder | Feralheart Mantle (226785) | Mokvar [vendor] | sim-verified (187.5 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Argent Elite Shoulders (227888, -2.36 DPS, sim-verified) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 81.4 healing_power points (8.37 DPS) | yes | Butcher's Apron (12608, -0.91 DPS) [dungeon]; Featherskin Cape (10843, -1.47 DPS, sim-verified) [world]; Shroud of the Exile (15421, -1.64 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, -1.76 DPS) [crafted]; Alanna's Embrace (13314, -1.91 DPS) [dungeon]; Tunic of Undead Slaying (23089, -15.95 DPS, sim-verified) [world] |
| wrist | Bracers of Mending (23129) | Dire Maul: Revanchion [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Hope (22667, -0.00 DPS) [quest]; Bleak Howler Armguards (13208, -0.07 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.40 DPS, sim-verified) [world] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 130.3 healing_power points (13.41 DPS) | yes | Wildheart Gloves (16717, -0.40 DPS) [dungeon]; Hands of the Exalted Herald (12554, -1.41 DPS) [dungeon]; Devout Gloves (16692, -1.62 DPS) [dungeon] |
| waist | Caretaker's Cord (272398) | Pix Xizzix [vendor] | 131.7 healing_power points (13.56 DPS) | yes | Feralheart Cord (226780, -1.17 DPS, sim-verified) [vendor]; Elderwild Waistcord (279252, -2.08 DPS) [crafted]; Wisdom of the Timbermaw (19047, -2.93 DPS) [crafted] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 198.9 healing_power points (20.47 DPS) | yes | Haunting Specter Leggings (11929, -4.64 DPS) [dungeon]; Legionnaire's Dragonhide Legguards (227206, -5.04 DPS) [vendor]; Devout Skirt (16694, -6.83 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 153.4 healing_power points (15.79 DPS) | yes | Incandescent Mooncloth Boots (227862, -3.37 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -4.02 DPS) [crafted]; Devout Sandals (16691, -4.42 DPS) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -1.01 DPS) [dungeon]; Band of the Hierophant (13096, -1.18 DPS) [world_drop]; Naglering (11669, -4.03 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, -0.44 DPS) [dungeon]; Band of the Hierophant (13096, -0.61 DPS) [world_drop]; Naglering (11669, -7.45 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindtap Talisman (18371, -2.79 DPS) [dungeon]; Ankh of Life (1713, -4.88 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -5.81 DPS) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Mindtap Talisman (18371, +0.00 DPS, sim-verified) [dungeon]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -0.25 DPS) [world_drop]; Soulkeeper (1607, -1.63 DPS) [world_drop]; Hand of Edward the Odd (2243, -3.70 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Feralheart Mantle; back: Frostweaver Cape; wrist: Bracers of Mending; waist: Caretaker's Cord; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: The Postmaster's Seal; finger2: Emerald Flame Ring; trinket2: Royal Seal of Eldre'Thalas; main_hand: Dancing Sliver

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 509.4. Weights run: 5.4s. Verify run: 2.6s. 1446 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.266, intellect=1.074 ± 0.045, spirit=2.070 ± 0.045, mp5=3.507 ± 0.053, crit=0.196 ± 0.023 per rating point (14 rating = 1%, 2.749 per %), spell_haste=not significant (2.112 ± 0.836)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 121.3 healing_power points (23.43 DPS) | yes | Champion's Dragonhide Headdress (227205, -5.76 DPS) [vendor]; Feralheart Headdress (226786, -6.60 DPS) [vendor]; Sanctified Leather Helm (22689, -6.96 DPS) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 58.5 healing_power points (11.30 DPS) | yes | Animated Chain Necklace (18723, -2.53 DPS) [dungeon]; Amulet of the Redeemed (22327, -3.34 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.36 DPS) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 100.7 healing_power points (19.45 DPS) | yes | Champion's Dragonhide Pauldrons (227207, -5.17 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, -7.06 DPS) [vendor]; Feralheart Mantle (226785, -7.71 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.7 healing_power points (10.18 DPS) | yes | Drape of Recovery (272413, -1.49 DPS) [vendor]; Battle Healer's Cloak (19526, -1.97 DPS) [rep]; Raincaster Drape (12110, -2.34 DPS) [quest] |
| chest | Robes of the Exalted (13346) | Stratholme: Baron Rivendare [dungeon] | sim-verified (509.3 DPS) | yes | Tunic of Undead Slaying (23089, -0.62 DPS, sim-verified) [world]; Mooncloth Vest (14138, -1.80 DPS) [crafted]; Feralheart Embrace (226783, -2.11 DPS) [vendor] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (509.3 DPS) | yes | Bracers of Mending (23129, -0.02 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.40 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -0.64 DPS, sim-verified) [world] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 73.1 healing_power points (14.11 DPS) | yes | Hands of the Exalted Herald (12554, -0.25 DPS) [dungeon]; Wildheart Gloves (16717, -1.73 DPS) [dungeon]; Devout Gloves (16692, -2.34 DPS) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 73.7 healing_power points (14.23 DPS) | yes | Caretaker's Cord (272398, -0.20 DPS) [vendor]; Elderwild Waistcord (279252, -1.24 DPS) [crafted]; Feralheart Cord (226780, -1.52 DPS) [vendor] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 101.6 healing_power points (19.61 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, -1.42 DPS) [vendor]; Elderwild Pants (279254, -2.30 DPS) [crafted]; Devout Skirt (16694, -2.68 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 88.0 healing_power points (17.00 DPS) | yes | Feralheart Sandals (226781, -1.09 DPS) [vendor]; Mooncloth Boots (15802, -4.26 DPS) [crafted]; Faith Healer's Boots (22247, -4.69 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (509.3 DPS) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Band of Piety (22681, -0.80 DPS) [quest]; Rosewine Circle (13178, -0.85 DPS) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (509.3 DPS) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Band of Piety (22681, -0.15 DPS) [quest]; Rosewine Circle (13178, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (509.3 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Staff of Metanoia (22394, -2.47 DPS) [dungeon]; Guiding Stave of Wisdom (11932, -3.31 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Robes of the Exalted; wrist: Bracers of Hope; waist: Sash of Mercy; legs: Leggings of Arcana; feet: Incandescent Mooncloth Boots; finger1: Band of Mending; finger2: Blessed Band of Light; main_hand: Redemption

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

