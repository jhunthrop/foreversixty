# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5032100000000000)

Set DPS (verified): 27.5. Weights run: 4.7s. Verify run: 15.4s. 226 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.854 ± 0.007, spirit=1.848 ± 0.004, mp5=4.988 ± 0.007, crit=0.120 ± 0.006 per rating point (14 rating = 1%, 1.676 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.4 healing_power points (0.87 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.14 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.16 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 7.4 healing_power points (0.32 DPS) | yes | Scholarly Pendant (277203, -0.15 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.7 healing_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, -0.20 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 12.7 healing_power points (0.54 DPS) | yes | Regent's Cloak (5969, -0.18 DPS, sim-verified) [world]; Sanguine Cape (14376, -0.23 DPS) [world_drop]; Seer's Cape (6378, -0.23 DPS) [dungeon] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.8 healing_power points (1.28 DPS) | yes | Robe of the Moccasin (6465, -0.10 DPS) [dungeon]; Filigreed Pristine Gown (253901, -0.13 DPS, sim-verified) [crafted]; Corsair's Overshirt (5202, -0.17 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 24.2 healing_power points (1.04 DPS) | yes | Owl Bracers (4796, -0.64 DPS) [vendor]; Drakewing Bands (12999, -0.72 DPS, sim-verified) [world_drop]; Bright Bracers (3647, -0.72 DPS) [world_drop] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 17.4 healing_power points (0.75 DPS) | yes | Pristine Gloves (253913, -0.04 DPS) [crafted]; Bright Gloves (3066, -0.11 DPS) [world_drop]; Magefist Gloves (12977, -0.15 DPS, sim-verified) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.4 healing_power points (0.79 DPS) | yes | Novice Ardent's Sash (253887, -0.06 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.15 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.5 healing_power points (1.65 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.62 DPS) [world_drop]; Dreamer's Leggings (270016, -0.78 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.3 healing_power points (0.87 DPS) | yes | Pristine Boots (253889, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.29 DPS, sim-verified) [crafted]; Kimbra Boots (6191, -0.31 DPS) [quest] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 14.8 healing_power points (0.64 DPS) | yes | Band of Purification (12996, -0.16 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.21 DPS) [rep]; Deep Fathom Ring (6463, -0.24 DPS) [dungeon] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.1 healing_power points (0.48 DPS) | yes | Lorekeeper's Ring (20431, -0.05 DPS) [rep]; Deep Fathom Ring (6463, -0.08 DPS) [dungeon]; Band of Purification (12996, -0.24 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 20.4 healing_power points (0.87 DPS) | yes | Staff of the Blessed Seer (2271, -0.08 DPS) [dungeon]; Twisted Chanter's Staff (890, -0.12 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.24 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 226, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-000000000000000000-5032503300000000)

Set DPS (verified): 49.2. Weights run: 4.8s. Verify run: 15.1s. 395 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.993 ± 0.007, spirit=2.468 ± 0.013, mp5=6.283 ± 0.014, crit=0.174 ± 0.008 per rating point (14 rating = 1%, 2.431 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.2 healing_power points (2.03 DPS) | yes | Holy Shroud (2721, -0.06 DPS) [world_drop]; Embalmed Shroud (7691, -0.15 DPS) [dungeon]; Whisperwind Headdress (6688, -0.43 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.8 healing_power points (0.74 DPS) | yes | Pendant of Myzrael (4614, -0.13 DPS) [dungeon]; Glowing Green Talisman (5002, -0.13 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.48 DPS, sim-verified) [world_drop] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 34.3 healing_power points (1.42 DPS) | yes | Nightsky Mantle (4718, -0.25 DPS) [world_drop]; Faerie Mantle (5820, -0.31 DPS) [quest]; Mantle of Honor (3560, -0.49 DPS, sim-verified) [quest] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 31.3 healing_power points (1.29 DPS) | yes | Prelacy Cape (7004, -0.23 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.33 DPS) [vendor]; Glowing Thresher Cape (6901, -0.46 DPS, sim-verified) [dungeon] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.8 healing_power points (1.85 DPS) | yes | Death Speaker Robes (6682, -0.12 DPS) [dungeon]; Pristine Gown (253961, -0.25 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.38 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 28.8 healing_power points (1.19 DPS) | yes | Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.53 DPS) [world_drop]; Nightsky Wristbands (6407, -0.66 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 39.6 healing_power points (1.64 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.53 DPS, sim-verified) [dungeon]; Zodiac Gloves (7106, -0.59 DPS) [quest]; Silver-thread Gloves (6393, -0.72 DPS) [world_drop] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.8 healing_power points (1.89 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.90 DPS) [rep]; Highlander's Mail Girdle (20120, -0.90 DPS) [vendor]; Silver-lined Belt (13011, -1.16 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.2 healing_power points (2.16 DPS) | yes | Pristine Leggings (253987, -0.35 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.42 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.8 healing_power points (1.81 DPS) | yes | Glinteye Slippers (273024, -0.71 DPS) [dungeon]; Acidic Walkers (9454, -0.75 DPS) [dungeon]; Soggy Boots (274747, -1.44 DPS, sim-verified) [vendor] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 31.2 healing_power points (1.29 DPS) | yes | The Queen's Jewel (13094, -0.31 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.33 DPS) [dungeon]; Black Pearl Ring (6332, -0.51 DPS) [world] |
| finger2 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 25.1 healing_power points (1.04 DPS) | yes | Electrocutioner Lagnut (9447, -0.08 DPS) [dungeon]; Black Pearl Ring (6332, -0.26 DPS) [world]; The Queen's Jewel (13094, -0.39 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (49.2 DPS) | yes | Royal Diplomatic Scepter (9457, -1.17 DPS) [dungeon]; Death Speaker Scepter (2816, -2.17 DPS) [dungeon]; Manual Crowd Pummeler (9449, -6.88 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Repairman's Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Darkspear Signet; main_hand: Wind Spirit Staff

No-known-source sample (15 of 395, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-000000000000000000-5032503315400000)

Set DPS (verified): 72.6. Weights run: 5.8s. Verify run: 23.8s. 622 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.563 ± 0.017, spirit=2.761 ± 0.012, mp5=7.128 ± 0.021, crit=0.285 ± 0.015 per rating point (14 rating = 1%, 3.996 per %), spell_haste=not significant (0.016 ± 0.016)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 90.4 healing_power points (3.72 DPS) | yes | Miner's Hat of the Deep (9429, -0.79 DPS) [dungeon]; Corpseshroud (10574, -1.03 DPS) [dungeon]; Whitemane's Chapeau (7720, -1.76 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) (or Necklace of Calisea (1714)) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 37.3 healing_power points (1.53 DPS) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.15 DPS) [dungeon]; Amberglow Talisman (10824, -0.40 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 57.6 healing_power points (2.37 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS, sim-verified) [dungeon]; Mistscape Mantle (4734, -0.64 DPS) [dungeon]; Batwing Mantle (6697, -0.64 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 39.2 healing_power points (1.61 DPS) | yes | Repairman's Cape (9605, -0.17 DPS, sim-verified) [quest]; Mantle of Lady Falther'ess (23178, -0.29 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.30 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Deathchill Armor (10764, -0.11 DPS) [dungeon]; Icemail Jerkin (1981, -0.57 DPS) [world_drop]; Doomsayer's Robe (4746, -1.77 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 37.9 healing_power points (1.56 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Silkstream Cuffs (16791, -0.58 DPS) [quest]; Earthen Silk Cuffs (254019, -0.58 DPS) [crafted] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 51.6 healing_power points (2.12 DPS) | yes | Mender's Leather Gloves (252530, -0.31 DPS) [crafted]; Stormcloth Gloves (10011, -0.40 DPS) [crafted]; Gloves of Old (9395, -0.61 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 59.1 healing_power points (2.43 DPS) | yes | Mender's Leather Belt (252523, -0.54 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.56 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Misplaced Pantaloons (276201) | Friz Frazzlespark [vendor] | 69.2 healing_power points (2.85 DPS) | yes | Stoneweaver Leggings (9407, -0.11 DPS) [dungeon]; Warchief Kilt (7760, -0.36 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -0.45 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Boots (252565, -0.00 DPS) [crafted]; Gilded Slippers (254001, -0.50 DPS) [crafted]; Furen's Boots (13100, -0.99 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.8 healing_power points (1.76 DPS) | yes | Welken Ring (5011, -0.44 DPS) [world_drop]; The Queen's Jewel (13094, -0.64 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 37.3 healing_power points (1.53 DPS) | yes | Welken Ring (5011, -0.30 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.41 DPS) [world_drop]; Voodoo Band (1996, -0.45 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Righteousness (7721, -0.49 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -1.38 DPS) [dungeon]; Gut Ripper (2164, -8.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Misplaced Pantaloons; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life

No-known-source sample (15 of 622, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 0000000000000000-000000000000000000-5032503315513131)

Set DPS (verified): 122.8. Weights run: 9.8s. Verify run: 40.7s. 796 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.070, intellect=2.374 ± 0.023, spirit=2.996 ± 0.018, mp5=7.513 ± 0.025, crit=0.449 ± 0.025 per rating point (14 rating = 1%, 6.293 per %), spell_haste=not significant (0.132 ± 0.076)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 103.3 healing_power points (5.29 DPS) | yes | Helm of Exile (11124, -0.23 DPS, sim-verified) [quest]; Papal Fez (9431, -0.72 DPS) [dungeon]; Soulcatcher Halo (10630, -0.72 DPS) [dungeon] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 59.3 healing_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.49 DPS, sim-verified) [world_drop]; Lei of Lilies (1315, -0.74 DPS) [world_drop]; Gemshard Heart (17707, -0.90 DPS) [dungeon] |
| shoulder | Lead Surveyor's Mantle (11842) | Blackrock Depths: Fineous Darkvire [dungeon] | 72.2 healing_power points (3.70 DPS) | yes | Living Shoulders (15061, -0.12 DPS) [crafted]; Mender's Leather Shoulder (252538, -0.30 DPS) [crafted]; Ironfeather Shoulders (15067, -0.61 DPS, sim-verified) [crafted] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | 54.4 healing_power points (2.79 DPS) | yes | Darkspear Raider's Cloak (272076, +0.00 DPS, sim-verified) [vendor]; Caretaker's Cape (19531, -0.74 DPS) [rep]; Imperial Red Cloak (8248, -0.84 DPS) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 130.2 healing_power points (6.68 DPS) | yes | Ghostweave Vest (14141, -0.96 DPS, sim-verified) [crafted]; Vestments of the Atal'ai Prophet (10806, -1.19 DPS) [dungeon]; Robes of Insight (940, -1.33 DPS) [world_drop] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 53.6 healing_power points (2.75 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Aristocratic Cuffs (12546, -0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 104.0 healing_power points (5.33 DPS) | yes | Mender's Leather Gauntlets (252551, -1.27 DPS) [crafted]; Mender's Mail Gauntlets (252587, -1.27 DPS) [crafted]; Stonerender Gauntlets (17007, -2.43 DPS, sim-verified) [world_drop] |
| waist | Bloodlust Belt (14803) | World drop [world_drop] | 60.2 healing_power points (3.09 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -0.06 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.09 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 102.6 healing_power points (5.26 DPS) | yes | Dalewind Trousers (13008, -0.33 DPS) [world_drop]; Jinxed Hoodoo Kilt (9474, -0.34 DPS, sim-verified) [dungeon]; Senior Designer's Pantaloons (11841, -0.36 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 78.9 healing_power points (4.04 DPS) | yes | Mistwalker Boots (10629, -0.24 DPS, sim-verified) [dungeon]; Furen's Boots (13100, -0.74 DPS) [world_drop]; Vinerot Sandals (17748, -0.80 DPS) [dungeon] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 52.6 healing_power points (2.70 DPS) | yes | Eye of Adaegus (5266, -0.12 DPS) [world_drop]; Choking Band (11868, -0.70 DPS) [quest]; Cyclopean Band (11824, -0.77 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 50.6 healing_power points (2.59 DPS) | yes | Choking Band (11868, -0.60 DPS) [quest]; Cyclopean Band (11824, -0.67 DPS) [dungeon]; Eye of Adaegus (5266, -0.78 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (122.8 DPS) | yes | Ankh of Life (1713, -3.93 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.24 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.55 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (122.8 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.54 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (122.8 DPS) | yes | Wind Spirit Staff (6689, -1.35 DPS) [dungeon]; Hand of Righteousness (7721, -2.01 DPS) [dungeon]; Barman Shanker (12791, -14.38 DPS, sim-verified) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | 53.8 healing_power points (2.76 DPS) | yes | Desertwalker Cane (12471, -0.27 DPS) [dungeon]; Enthralled Sphere (11625, -0.29 DPS) [dungeon]; Cloud Stone (17737, -14.66 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Lead Surveyor's Mantle; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Bloodlust Belt; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 796, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 229.2. Weights run: 11.2s. Verify run: 140.1s. 1759 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.110, intellect=4.396 ± 0.057, spirit=4.317 ± 0.051, mp5=10.946 ± 0.059, crit=0.807 ± 0.056 per rating point (14 rating = 1%, 11.298 per %), spell_haste=not significant (0.050 ± 0.260)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (229.2 DPS) | yes | Crown of the Penitent (13216, -0.62 DPS) [quest]; Devout Crown (16693, -1.10 DPS) [dungeon]; Gnomish Turban of Psychic Might (21517, -2.83 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 137.9 healing_power points (6.80 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.77 DPS) [world_drop]; Lady Maye's Pendant (14558, -1.79 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -2.15 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 167.9 healing_power points (8.27 DPS) | yes | Soulstealer Mantle (13374, -1.59 DPS) [dungeon]; Mantle of The Five Thunders (227011, -1.82 DPS) [vendor]; Devout Mantle (16695, -3.48 DPS, sim-verified) [dungeon] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 104.6 healing_power points (5.15 DPS) | yes | Faded Hakkari Cloak (20218, -0.18 DPS) [quest]; Gracious Cape (18743, -0.40 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -0.41 DPS) [vendor] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-verified (229.2 DPS) | yes | Mooncloth Vest (14138, -0.41 DPS) [crafted]; Alanna's Embrace (13314, -0.49 DPS) [dungeon]; Tunic of Undead Slaying (23089, -11.89 DPS, sim-verified) [world] |
| wrist | Bracers of The Five Thunders (227009) | Mokvar [vendor] | sim-verified (229.2 DPS) | yes | Bleak Howler Armguards (13208, +0.00 DPS) [dungeon]; Bracers of Mending (23129, +0.00 DPS) [dungeon]; Bracers of Hope (22667, -4.31 DPS, sim-verified) [quest] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 150.1 healing_power points (7.39 DPS) | yes | Hands of the Exalted Herald (12554, -0.40 DPS) [dungeon]; Devout Gloves (16692, -0.87 DPS) [dungeon]; Raider Handwraps (272097, -3.54 DPS, sim-verified) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 151.4 healing_power points (7.46 DPS) | yes | Devout Belt (16696, -0.67 DPS) [dungeon]; Whipvine Cord (18327, -0.74 DPS) [dungeon]; Sash of The Five Thunders (227010, -0.92 DPS) [vendor] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 235.4 healing_power points (11.60 DPS) | yes | Legplates of the Chromatic Defier (12945, -1.95 DPS) [quest]; Padre's Trousers (18386, -1.96 DPS) [dungeon]; Ghostloom Leggings (14545, -2.67 DPS, sim-verified) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 185.4 healing_power points (9.13 DPS) | yes | Mooncloth Boots (15802, -2.15 DPS) [crafted]; Faith Healer's Boots (22247, -2.70 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -4.34 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-verified (229.2 DPS) | yes | Ring of Demonic Guile (18314, -0.97 DPS) [dungeon]; Seal of Rivendare (13345, -1.20 DPS) [dungeon]; Naglering (11669, -10.51 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (229.2 DPS) | yes | Ring of Demonic Guile (18314, -0.01 DPS) [dungeon]; Seal of Rivendare (13345, -0.24 DPS) [dungeon]; Naglering (11669, -8.53 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (229.2 DPS) | yes | Shard of the Splithooves (10659, -4.85 DPS) [quest]; Ankh of Life (1713, -5.54 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -5.96 DPS) [quest] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-verified (229.2 DPS) | yes | Serenity Field (272439, -1.71 DPS, sim-verified) [vendor]; Shard of the Splithooves (10659, -2.70 DPS) [quest]; Ankh of Life (1713, -3.38 DPS) [world_drop] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-verified (229.2 DPS) | yes | Quel'dorai Channeling Rod (18311, -0.18 DPS) [dungeon]; Charstone Dirk (17710, -0.92 DPS) [dungeon]; Hand of Edward the Odd (2243, -6.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Frostweaver Cape; wrist: Bracers of The Five Thunders; waist: Belt of Tiny Heads; legs: Leggings of Arcana; feet: Greaves of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Mindtap Talisman; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 627.4. Weights run: 8.1s. Verify run: 81.1s. 1759 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.172, intellect=1.187 ± 0.027, spirit=0.802 ± 0.022, mp5=2.197 ± 0.028, crit=0.551 ± 0.028 per rating point (14 rating = 1%, 7.714 per %), spell_haste=not significant (0.648 ± 0.369)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 105.6 healing_power points (29.21 DPS) | yes | Crown of The Five Thunders (227013, -4.59 DPS, sim-verified) [vendor]; Sanctified Leather Helm (22689, -9.24 DPS) [quest]; Insightful Hood (18490, -11.03 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 48.7 healing_power points (13.47 DPS) | yes | Animated Chain Necklace (18723, -3.01 DPS) [dungeon]; Drake Tooth Necklace (21531, -3.60 DPS, sim-verified) [quest]; Amulet of the Redeemed (22327, -4.59 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 80.2 healing_power points (22.18 DPS) | yes | Royal Cap Spaulders (14548, -7.51 DPS, sim-verified) [dungeon]; Mantle of The Five Thunders (227011, -9.17 DPS) [vendor]; Mooncloth Shoulders (14139, -9.51 DPS) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.9 healing_power points (14.90 DPS) | yes | Drape of Recovery (272413, -2.61 DPS, sim-verified) [vendor]; Cloak of the Cosmos (18389, -4.10 DPS) [dungeon]; Caretaker's Cape (19530, -5.94 DPS) [rep] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | sim-verified (627.2 DPS) | yes | Mooncloth Vest (14138, -2.98 DPS) [crafted]; Stormcloth Vest (10020, -3.13 DPS) [crafted]; Robes of the Exalted (13346, -9.05 DPS, sim-verified) [dungeon] |
| wrist | Bracers of The Five Thunders (227009) | Mokvar [vendor] | sim-verified (627.2 DPS) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, -7.35 DPS, sim-verified) [dungeon] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | sim-verified (627.2 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Harmonious Gauntlets (18527, -7.04 DPS, sim-verified) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 61.0 healing_power points (16.88 DPS) | yes | Whipvine Cord (18327, +0.00 DPS, sim-verified) [dungeon]; Wisdom of the Timbermaw (19047, -1.71 DPS) [crafted]; Eyestalk Cord (18391, -2.27 DPS) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 78.9 healing_power points (21.83 DPS) | yes | Leggings of Arcana (12756, -2.15 DPS, sim-verified) [quest]; Leggings of The Five Thunders (227012, -4.01 DPS) [vendor]; Devout Skirt (16694, -5.17 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 71.9 healing_power points (19.87 DPS) | yes | Greaves of The Five Thunders (227015, -2.37 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -5.75 DPS) [crafted]; Faith Healer's Boots (22247, -6.08 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-verified (627.2 DPS) | yes | Fordring's Seal (16058, -3.80 DPS) [quest]; Band of Piety (22681, -3.88 DPS) [quest]; Naglering (11669, -13.15 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (627.2 DPS) | yes | Fordring's Seal (16058, -0.33 DPS) [quest]; Band of Piety (22681, -0.40 DPS) [quest]; Naglering (11669, -9.60 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (627.2 DPS) | yes | Briarwood Reed (12930, -1.09 DPS) [dungeon]; Mindtap Talisman (18371, -2.99 DPS, sim-verified) [dungeon]; Second Wind (11819, -3.03 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (627.2 DPS) | yes | Briarwood Reed (12930, -3.60 DPS) [dungeon]; Mindtap Talisman (18371, -4.93 DPS) [dungeon]; Second Wind (11819, -5.53 DPS) [dungeon] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (627.2 DPS) | yes | Hand of Righteousness (7721, -9.14 DPS) [dungeon]; Wind Spirit Staff (6689, -12.08 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.02 DPS, sim-verified) [world_drop] |
| off_hand | Skullflame Shield (1168) | World drop [world_drop] | sim-verified (627.2 DPS) | yes | Tome of Divine Right (22319, +0.00 DPS) [dungeon]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Lei of the Lifegiver (19312, -16.02 DPS, sim-verified) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of The Five Thunders; waist: Sash of Mercy; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Mending; trinket2: Serenity Field; off_hand: Skullflame Shield

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (tauren, 0000000000000000-000000000000000000-5032100000000000)

Set DPS (verified): 27.7. Weights run: 4.7s. Verify run: 14.5s. 206 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.854 ± 0.007, spirit=1.848 ± 0.004, mp5=4.988 ± 0.007, crit=0.120 ± 0.006 per rating point (14 rating = 1%, 1.676 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.4 healing_power points (0.87 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.14 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.16 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 5.5 healing_power points (0.24 DPS) | yes | Scholarly Pendant (277203, -0.26 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.7 healing_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, -0.16 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 12.7 healing_power points (0.54 DPS) | yes | Traveler's Shawl (277289, -0.15 DPS) [quest]; Sanguine Cape (14376, -0.23 DPS) [world_drop]; Regent's Cloak (5969, -0.36 DPS, sim-verified) [world] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.8 healing_power points (1.28 DPS) | yes | Robe of the Moccasin (6465, -0.10 DPS) [dungeon]; Filigreed Pristine Gown (253901, -0.13 DPS, sim-verified) [crafted]; Corsair's Overshirt (5202, -0.17 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 24.2 healing_power points (1.04 DPS) | yes | Owl Bracers (4796, -0.64 DPS) [vendor]; Featherbead Bracers (15452, -0.64 DPS) [quest]; Drakewing Bands (12999, -0.69 DPS, sim-verified) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 22.2 healing_power points (0.95 DPS) | yes | Wisdom's Leather Gloves (252499, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.24 DPS) [world_drop]; Pristine Gloves (253913, -0.24 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.4 healing_power points (0.79 DPS) | yes | Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.15 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.21 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.5 healing_power points (1.65 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.62 DPS) [world_drop]; Ghastly Trousers (15449, -0.83 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.3 healing_power points (0.87 DPS) | yes | Pristine Boots (253889, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.39 DPS, sim-verified) [crafted]; Bluegill Sandals (1560, -0.39 DPS) [world] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 14.8 healing_power points (0.64 DPS) | yes | Band of Purification (12996, -0.16 DPS) [world_drop]; Advisor's Ring (20426, -0.21 DPS) [rep]; Loop of Sacrifice (281673, -0.24 DPS) [quest] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.1 healing_power points (0.48 DPS) | yes | Advisor's Ring (20426, -0.05 DPS) [rep]; Band of Purification (12996, -0.08 DPS, sim-verified) [world_drop]; Loop of Sacrifice (281673, -0.08 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 26.6 healing_power points (1.14 DPS) | yes | Advisor's Gnarled Staff (20425, -0.18 DPS) [pvp]; Twisted Chanter's Staff (890, -0.35 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.36 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 206, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (tauren, 0000000000000000-000000000000000000-5032503300000000)

Set DPS (verified): 48.5. Weights run: 4.8s. Verify run: 14.4s. 377 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.993 ± 0.007, spirit=2.468 ± 0.013, mp5=6.283 ± 0.014, crit=0.174 ± 0.008 per rating point (14 rating = 1%, 2.431 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.2 healing_power points (2.03 DPS) | yes | Whisperwind Headdress (6688, +0.00 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -0.06 DPS) [world_drop]; Embalmed Shroud (7691, -0.15 DPS) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.8 healing_power points (0.74 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.13 DPS) [dungeon]; Glowing Green Talisman (5002, -0.13 DPS) [world_drop] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 34.3 healing_power points (1.42 DPS) | yes | Nightsky Mantle (4718, -0.25 DPS) [world_drop]; Ghostly Mantle (3324, -0.39 DPS, sim-verified) [quest]; Talbar Mantle (10657, -0.40 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 26.7 healing_power points (1.11 DPS) | yes | Darkspear Raider's Cloak (272078, -0.16 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19529, -0.16 DPS) [rep]; Construct Cloak (279848, -0.39 DPS) [quest] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.8 healing_power points (1.85 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.12 DPS) [dungeon]; Beguiler Robes (7728, -0.38 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 28.8 healing_power points (1.19 DPS) | yes | Nightsky Wristbands (6407, -0.36 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.53 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 39.6 healing_power points (1.64 DPS) | yes | Hotshot Pilot's Gloves (9491, +0.00 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.55 DPS) [quest]; Tattered Mittens (270030, -0.61 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.8 healing_power points (1.89 DPS) | yes | Lilac Sash (6780, -0.85 DPS) [quest]; Highlander's Mail Girdle (20120, -0.90 DPS) [vendor]; Silver-lined Belt (13011, -1.15 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.2 healing_power points (2.16 DPS) | yes | Pristine Leggings (253987, -0.08 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.42 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.8 healing_power points (1.81 DPS) | yes | Glinteye Slippers (273024, -0.71 DPS) [dungeon]; Acidic Walkers (9454, -0.75 DPS) [dungeon]; Soggy Boots (274747, -1.44 DPS, sim-verified) [vendor] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 31.2 healing_power points (1.29 DPS) | yes | The Queen's Jewel (13094, -0.31 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.33 DPS) [dungeon]; Black Pearl Ring (6332, -0.51 DPS) [world] |
| finger2 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 25.1 healing_power points (1.04 DPS) | yes | Electrocutioner Lagnut (9447, -0.08 DPS) [dungeon]; The Queen's Jewel (13094, -0.12 DPS, sim-verified) [world_drop]; Black Pearl Ring (6332, -0.26 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (48.5 DPS) | yes | Royal Diplomatic Scepter (9457, -1.17 DPS) [dungeon]; Death Speaker Scepter (2816, -2.17 DPS) [dungeon]; Manual Crowd Pummeler (9449, -6.83 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Darkspear Signet; main_hand: Wind Spirit Staff

No-known-source sample (15 of 377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-000000000000000000-5032503315400000)

Set DPS (verified): 72.3. Weights run: 5.8s. Verify run: 22.7s. 584 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.563 ± 0.017, spirit=2.761 ± 0.012, mp5=7.128 ± 0.021, crit=0.285 ± 0.015 per rating point (14 rating = 1%, 3.996 per %), spell_haste=not significant (0.016 ± 0.016)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 90.4 healing_power points (3.72 DPS) | yes | Miner's Hat of the Deep (9429, -0.79 DPS) [dungeon]; Corpseshroud (10574, -1.03 DPS) [dungeon]; Whitemane's Chapeau (7720, -1.79 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) (or Necklace of Calisea (1714)) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 37.3 healing_power points (1.53 DPS) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.15 DPS) [dungeon]; Amberglow Talisman (10824, -0.40 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 57.6 healing_power points (2.37 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS, sim-verified) [dungeon]; Mistscape Mantle (4734, -0.64 DPS) [dungeon]; Batwing Mantle (6697, -0.64 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 39.2 healing_power points (1.61 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.29 DPS) [dungeon]; First Sergeant's Cloak (16340, -0.30 DPS) [pvp]; Cloak of Blight (6832, -0.70 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Deathchill Armor (10764, -0.11 DPS) [dungeon]; Icemail Jerkin (1981, -0.57 DPS) [world_drop]; Doomsayer's Robe (4746, -1.74 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 37.9 healing_power points (1.56 DPS) | yes | Mindthrust Bracers (1974, -0.15 DPS) [dungeon]; Windtalker's Wristguards (19584, -0.36 DPS) [pvp]; Dryad's Wrist Bindings (19597, -0.36 DPS) [pvp] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 51.6 healing_power points (2.12 DPS) | yes | Mender's Leather Gloves (252530, -0.31 DPS) [crafted]; Stormcloth Gloves (10011, -0.40 DPS) [crafted]; Gloves of Old (9395, -0.73 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 59.1 healing_power points (2.43 DPS) | yes | Mender's Leather Belt (252523, -0.54 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.56 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Misplaced Pantaloons (276201) | Friz Frazzlespark [vendor] | 69.2 healing_power points (2.85 DPS) | yes | Stoneweaver Leggings (9407, -0.11 DPS) [dungeon]; Wisdom's Leather Leggings (252519, -0.45 DPS) [crafted]; Warchief Kilt (7760, -0.47 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -0.50 DPS) [crafted]; Furen's Boots (13100, -0.89 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.8 healing_power points (1.76 DPS) | yes | Welken Ring (5011, -0.44 DPS) [world_drop]; The Queen's Jewel (13094, -0.64 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 37.3 healing_power points (1.53 DPS) | yes | Welken Ring (5011, -0.30 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.41 DPS) [world_drop]; Voodoo Band (1996, -0.45 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Righteousness (7721, -0.49 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -1.38 DPS) [dungeon]; Gut Ripper (2164, -8.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Misplaced Pantaloons; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (tauren, 0000000000000000-000000000000000000-5032503315513131)

Set DPS (verified): 122.4. Weights run: 9.8s. Verify run: 40.1s. 737 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.070, intellect=2.374 ± 0.023, spirit=2.996 ± 0.018, mp5=7.513 ± 0.025, crit=0.449 ± 0.025 per rating point (14 rating = 1%, 6.293 per %), spell_haste=not significant (0.132 ± 0.076)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 103.3 healing_power points (5.29 DPS) | yes | Helm of Exile (11124, -0.34 DPS) [quest]; Papal Fez (9431, -0.72 DPS) [dungeon]; Soulcatcher Halo (10630, -0.72 DPS) [dungeon] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 59.3 healing_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.62 DPS, sim-verified) [world_drop]; Lei of Lilies (1315, -0.74 DPS) [world_drop]; Gemshard Heart (17707, -0.90 DPS) [dungeon] |
| shoulder | Lead Surveyor's Mantle (11842) | Blackrock Depths: Fineous Darkvire [dungeon] | 72.2 healing_power points (3.70 DPS) | yes | Living Shoulders (15061, -0.12 DPS) [crafted]; Mender's Leather Shoulder (252538, -0.30 DPS) [crafted]; Ironfeather Shoulders (15067, -0.62 DPS, sim-verified) [crafted] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | 54.4 healing_power points (2.79 DPS) | yes | Darkspear Raider's Cloak (272076, +0.00 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19527, -0.74 DPS) [rep]; Cloak of Blight (6832, -0.79 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 130.2 healing_power points (6.68 DPS) | yes | Ghostweave Vest (14141, -0.92 DPS, sim-verified) [crafted]; Vestments of the Atal'ai Prophet (10806, -1.19 DPS) [dungeon]; Robes of Insight (940, -1.33 DPS) [world_drop] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 53.6 healing_power points (2.75 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Aristocratic Cuffs (12546, -0.00 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 104.0 healing_power points (5.33 DPS) | yes | Mender's Leather Gauntlets (252551, -1.27 DPS) [crafted]; Mender's Mail Gauntlets (252587, -1.27 DPS) [crafted]; Stonerender Gauntlets (17007, -2.43 DPS, sim-verified) [world_drop] |
| waist | Bloodlust Belt (14803) | World drop [world_drop] | 60.2 healing_power points (3.09 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -0.06 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.09 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 102.6 healing_power points (5.26 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.33 DPS, sim-verified) [dungeon]; Dalewind Trousers (13008, -0.33 DPS) [world_drop]; Senior Designer's Pantaloons (11841, -0.36 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 78.9 healing_power points (4.04 DPS) | yes | Mistwalker Boots (10629, -0.30 DPS, sim-verified) [dungeon]; Furen's Boots (13100, -0.74 DPS) [world_drop]; Vinerot Sandals (17748, -0.80 DPS) [dungeon] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 52.6 healing_power points (2.70 DPS) | yes | Eye of Adaegus (5266, -0.12 DPS) [world_drop]; Cyclopean Band (11824, -0.77 DPS) [dungeon]; Snake Hoop (6750, -0.77 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 50.6 healing_power points (2.59 DPS) | yes | Cyclopean Band (11824, -0.67 DPS) [dungeon]; Snake Hoop (6750, -0.67 DPS) [quest]; Eye of Adaegus (5266, -0.71 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (122.4 DPS) | yes | Ankh of Life (1713, -3.93 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.24 DPS) [quest]; Alchemist's Stone (13503, -5.16 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (122.4 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.31 DPS) [crafted]; Ankh of Life (1713, -0.45 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (122.4 DPS) | yes | Wind Spirit Staff (6689, -1.35 DPS) [dungeon]; Hand of Righteousness (7721, -2.01 DPS) [dungeon]; Barman Shanker (12791, -14.39 DPS, sim-verified) [dungeon] |
| off_hand | Gizlock's Hypertech Buckler (17718) | Maraudon: Tinkerer Gizlock [dungeon] | 53.8 healing_power points (2.76 DPS) | yes | Desertwalker Cane (12471, -0.27 DPS) [dungeon]; Enthralled Sphere (11625, -0.29 DPS) [dungeon]; Cloud Stone (17737, -14.76 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Lead Surveyor's Mantle; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Bloodlust Belt; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; off_hand: Gizlock's Hypertech Buckler

No-known-source sample (15 of 737, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (tauren, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 228.6. Weights run: 11.2s. Verify run: 137.3s. 1679 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.110, intellect=4.396 ± 0.057, spirit=4.317 ± 0.051, mp5=10.946 ± 0.059, crit=0.807 ± 0.056 per rating point (14 rating = 1%, 11.298 per %), spell_haste=not significant (0.050 ± 0.260)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 234.8 healing_power points (11.57 DPS) | yes | Living Crown (252561, +0.00 DPS, sim-verified) [crafted]; Crown of the Penitent (13216, -1.72 DPS) [quest]; Devout Crown (16693, -2.19 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 137.9 healing_power points (6.80 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.77 DPS) [world_drop]; Lady Maye's Pendant (14558, -1.81 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -2.15 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 167.9 healing_power points (8.27 DPS) | yes | Soulstealer Mantle (13374, -1.59 DPS) [dungeon]; Mantle of The Five Thunders (227011, -1.82 DPS) [vendor]; Devout Mantle (16695, -3.84 DPS, sim-verified) [dungeon] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 104.6 healing_power points (5.15 DPS) | yes | Faded Hakkari Cloak (20218, -0.18 DPS) [quest]; Gracious Cape (18743, -0.40 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -0.41 DPS) [vendor] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | sim-verified (228.6 DPS) | yes | Alanna's Embrace (13314, +0.00 DPS) [dungeon]; Mooncloth Vest (14138, +0.00 DPS) [crafted]; Embrace of the Wind Serpent (12462, -3.32 DPS, sim-verified) [world] |
| wrist | Bracers of The Five Thunders (227009) | Mokvar [vendor] | sim-verified (228.6 DPS) | yes | Bleak Howler Armguards (13208, +0.00 DPS) [dungeon]; Bracers of Hope (22667, +0.00 DPS, sim-verified) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 150.1 healing_power points (7.39 DPS) | yes | Raider Handwraps (272097, +0.00 DPS, sim-verified) [vendor]; Hands of the Exalted Herald (12554, -0.40 DPS) [dungeon]; Devout Gloves (16692, -0.87 DPS) [dungeon] |
| waist | Sash of The Five Thunders (227010) | Mokvar [vendor] | sim-verified (228.6 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Devout Belt (16696, +0.00 DPS) [dungeon]; Belt of Tiny Heads (20217, +0.00 DPS) [quest] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 235.4 healing_power points (11.60 DPS) | yes | Legplates of the Chromatic Defier (12945, -1.95 DPS) [quest]; Padre's Trousers (18386, -1.96 DPS) [dungeon]; Ghostloom Leggings (14545, -2.72 DPS, sim-verified) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 185.4 healing_power points (9.13 DPS) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -2.15 DPS) [crafted]; Boots of The Five Thunders (22096, -2.25 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-verified (228.6 DPS) | yes | Ring of Demonic Guile (18314, -0.97 DPS) [dungeon]; Seal of Rivendare (13345, -1.20 DPS) [dungeon]; Naglering (11669, -10.56 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (228.6 DPS) | yes | Ring of Demonic Guile (18314, -0.01 DPS) [dungeon]; Seal of Rivendare (13345, -0.24 DPS) [dungeon]; Naglering (11669, -8.48 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (228.6 DPS) | yes | Shard of the Splithooves (10659, -4.85 DPS) [quest]; Ankh of Life (1713, -5.54 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -5.96 DPS) [quest] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-verified (228.6 DPS) | yes | Serenity Field (272439, -1.82 DPS, sim-verified) [vendor]; Shard of the Splithooves (10659, -2.70 DPS) [quest]; Ankh of Life (1713, -3.38 DPS) [world_drop] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-verified (228.6 DPS) | yes | Quel'dorai Channeling Rod (18311, -0.18 DPS) [dungeon]; Charstone Dirk (17710, -0.92 DPS) [dungeon]; Hand of Edward the Odd (2243, -6.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Frostweaver Cape; chest: Tunic of The Five Thunders; wrist: Bracers of The Five Thunders; waist: Sash of The Five Thunders; legs: Leggings of Arcana; feet: Greaves of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Mindtap Talisman; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (tauren, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 626.8. Weights run: 8.1s. Verify run: 79.2s. 1679 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.172, intellect=1.187 ± 0.027, spirit=0.802 ± 0.022, mp5=2.197 ± 0.028, crit=0.551 ± 0.028 per rating point (14 rating = 1%, 7.714 per %), spell_haste=not significant (0.648 ± 0.369)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 105.6 healing_power points (29.21 DPS) | yes | Crown of The Five Thunders (227013, -5.30 DPS, sim-verified) [vendor]; Sanctified Leather Helm (22689, -9.24 DPS) [quest]; Insightful Hood (18490, -11.03 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 48.7 healing_power points (13.47 DPS) | yes | Animated Chain Necklace (18723, -3.01 DPS) [dungeon]; Drake Tooth Necklace (21531, -3.87 DPS, sim-verified) [quest]; Amulet of the Redeemed (22327, -4.59 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 80.2 healing_power points (22.18 DPS) | yes | Royal Cap Spaulders (14548, -7.45 DPS, sim-verified) [dungeon]; Champion's Mail Epaulets (227166, -7.53 DPS) [vendor]; Warlord's Mail Epaulets (231665, -8.16 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.9 healing_power points (14.90 DPS) | yes | Drape of Recovery (272413, -3.00 DPS, sim-verified) [vendor]; Cloak of the Cosmos (18389, -4.10 DPS) [dungeon]; Battle Healer's Cloak (19526, -5.94 DPS) [rep] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | sim-verified (626.8 DPS) | yes | Legionnaire's Mail Chestguard (227171, -2.67 DPS) [vendor]; Mooncloth Vest (14138, -2.98 DPS) [crafted]; Robes of the Exalted (13346, -8.32 DPS, sim-verified) [dungeon] |
| wrist | Bracers of The Five Thunders (227009) | Mokvar [vendor] | sim-verified (626.8 DPS) | yes | Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon]; Loomguard Armbraces (13969, -6.22 DPS, sim-verified) [dungeon] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | sim-verified (626.8 DPS) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Harmonious Gauntlets (18527, -6.12 DPS, sim-verified) [dungeon] |
| waist | Sash of Mercy (14553) | World drop [world_drop] | 61.0 healing_power points (16.88 DPS) | yes | Whipvine Cord (18327, -1.70 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.71 DPS) [crafted]; Eyestalk Cord (18391, -2.27 DPS) [dungeon] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 78.9 healing_power points (21.83 DPS) | yes | Leggings of Arcana (12756, -2.59 DPS, sim-verified) [quest]; Legionnaire's Mail Pants (227167, -3.05 DPS) [vendor]; Leggings of The Five Thunders (227012, -4.01 DPS) [vendor] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 71.9 healing_power points (19.87 DPS) | yes | Greaves of The Five Thunders (227015, -1.96 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -5.75 DPS) [crafted]; Faith Healer's Boots (22247, -6.08 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-verified (626.8 DPS) | yes | Fordring's Seal (16058, -3.80 DPS) [quest]; Band of Piety (22681, -3.88 DPS) [quest]; Naglering (11669, -13.78 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (626.8 DPS) | yes | Fordring's Seal (16058, -0.33 DPS) [quest]; Band of Piety (22681, -0.40 DPS) [quest]; Naglering (11669, -10.40 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (626.8 DPS) | yes | Briarwood Reed (12930, -1.09 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, -2.75 DPS) [quest]; Serenity Field (272439, -3.34 DPS, sim-verified) [vendor] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-verified (626.8 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor]; Royal Seal of Eldre'Thalas (18471, -0.32 DPS) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (626.8 DPS) | yes | Hand of Righteousness (7721, -9.14 DPS) [dungeon]; Wind Spirit Staff (6689, -12.08 DPS) [dungeon]; Hand of Edward the Odd (2243, -21.57 DPS, sim-verified) [world_drop] |
| off_hand | Skullflame Shield (1168) | World drop [world_drop] | sim-verified (626.8 DPS) | yes | Tome of Divine Right (22319, +0.00 DPS) [dungeon]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Lei of the Lifegiver (19312, -16.78 DPS, sim-verified) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of The Five Thunders; waist: Sash of Mercy; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Mending; trinket2: Mindtap Talisman; off_hand: Skullflame Shield

No-known-source sample (15 of 1679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

