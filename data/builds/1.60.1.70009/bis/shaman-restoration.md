# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5032100000000000)

Set DPS (verified): 26.4. Weights run: 4.4s. Verify run: 2.3s. 225 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.751 ± 0.009, spirit=1.906 ± 0.006, mp5=5.250 ± 0.009, crit=0.123 ± 0.006 per rating point (14 rating = 1%, 1.728 per %), spell_haste=not significant (0.054 ± 0.020)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.2 healing_power points (0.83 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Stormrider's Leather Hood (252506, -0.13 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.13 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 7.6 healing_power points (0.31 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.8 healing_power points (0.65 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.34 DPS) [world_drop]; Prospector's Pads (14566, -0.34 DPS) [world_drop] |
| back | Regent's Cloak (5969) | Ravenclaw Regent [world] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (20428, -0.02 DPS, sim-verified) [rep]; Spirit Cloak (4792, -0.08 DPS) [vendor]; Sylvan Cloak (4793, -0.08 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.5 healing_power points (1.22 DPS) | yes | Robe of the Moccasin (6465, +0.00 DPS, sim-verified) [dungeon]; Filigreed Pristine Gown (253901, -0.08 DPS) [crafted]; Corsair's Overshirt (5202, -0.13 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 24.5 healing_power points (1.01 DPS) | yes | Drakewing Bands (12999, +0.00 DPS, sim-verified) [world_drop]; Owl Bracers (4796, -0.65 DPS) [vendor]; Bravo's Armbands (270015, -0.70 DPS) [quest] |
| hands | Magefist Gloves (12977) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gloves (253913, -0.01 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.03 DPS, sim-verified) [crafted]; Bright Gloves (3066, -0.07 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.0 healing_power points (0.74 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.17 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.1 healing_power points (1.57 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.60 DPS) [world_drop]; Dreamer's Leggings (270016, -0.78 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 19.8 healing_power points (0.82 DPS) | yes | Black Whelp Slippers (252424, +0.00 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.23 DPS) [crafted]; Kimbra Boots (6191, -0.29 DPS) [quest] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 14.9 healing_power points (0.62 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Lorekeeper's Ring (20431, -0.18 DPS) [rep]; Deep Fathom Ring (6463, -0.22 DPS) [dungeon] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Lorekeeper's Ring (20431, -0.00 DPS) [rep]; Band of Purification (12996, -0.04 DPS, sim-verified) [world_drop]; Deep Fathom Ring (6463, -0.04 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 20.2 healing_power points (0.83 DPS) | yes | Staff of the Blessed Seer (2271, +0.00 DPS, sim-verified) [dungeon]; Twisted Chanter's Staff (890, -0.11 DPS) [world_drop]; Taskmaster Axe (5194, -0.20 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Regent's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Magefist Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-000000000000000000-5032503300000000)

Set DPS (verified): 43.5. Weights run: 4.5s. Verify run: 2.4s. 366 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=2.418 ± 0.014, spirit=2.475 ± 0.013, mp5=6.455 ± 0.015, crit=0.185 ± 0.009 per rating point (14 rating = 1%, 2.592 per %), spell_haste=not significant (0.199 ± 0.050)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 56.0 healing_power points (2.21 DPS) | yes | Whisperwind Headdress (6688, -0.05 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -0.32 DPS) [world_drop]; Nightsky Cowl (4039, -0.58 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 19.6 healing_power points (0.77 DPS) | yes | Necklace of Harmony (5180, -0.05 DPS, sim-verified) [world]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Pendant of Myzrael (4614, -0.19 DPS) [dungeon] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 39.0 healing_power points (1.54 DPS) | yes | Mantle of Honor (3560, -0.05 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.29 DPS) [world_drop]; Faerie Mantle (5820, -0.38 DPS) [quest] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 27.0 healing_power points (1.06 DPS) | yes | Darkspear Raider's Cloak (272078, -0.01 DPS) [vendor]; Prelacy Cape (7004, -0.05 DPS) [quest]; Glowing Thresher Cape (6901, -0.05 DPS, sim-verified) [dungeon] |
| chest | Beguiler Robes (7728) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 48.8 healing_power points (1.93 DPS) | yes | Wisdom's Leather Tunic (252511, +0.00 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.08 DPS) [crafted]; Death Speaker Robes (6682, -0.09 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 31.5 healing_power points (1.24 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.48 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.48 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan: Caverndeep Burrower [dungeon] | sim-verified (1.2 DPS) | yes | Gloves of Old (9395, -0.02 DPS, sim-verified) [world_drop]; Zodiac Gloves (7106, -0.19 DPS) [quest]; Town Clerk's Mittens (270029, -0.20 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 48.4 healing_power points (1.91 DPS) | yes | Silver-lined Belt (13011, +0.00 DPS, sim-verified) [world_drop]; Highlander's Lizardhide Girdle (20105, -0.76 DPS) [rep]; Highlander's Mail Girdle (20120, -0.76 DPS) [vendor] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 55.3 healing_power points (2.18 DPS) | yes | Pristine Leggings (253987, -0.16 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 46.8 healing_power points (1.85 DPS) | yes | Soggy Boots (274747, -0.08 DPS, sim-verified) [vendor]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Frothing Slippers (254003, -0.79 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 34.3 healing_power points (1.35 DPS) | yes | Darkspear Signet (272071, -0.09 DPS, sim-verified) [vendor]; Black Pearl Ring (6332, -0.57 DPS) [world]; Sea Giant's Toe Ring (274746, -0.64 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Signet (272071, -0.02 DPS, sim-verified) [vendor]; Black Pearl Ring (6332, -0.20 DPS) [world]; Sea Giant's Toe Ring (274746, -0.26 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (1.2 DPS) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, +0.00 DPS, sim-verified) [dungeon]; Wind Spirit Staff (6689, -0.74 DPS) [dungeon]; Lorekeeper's Staff (212580, -1.09 DPS) [vendor] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 27.9 healing_power points (1.10 DPS) | yes | Defective Samophlange (274743, +0.00 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -0.33 DPS) [world_drop]; Orb of Mistmantle (13031, -0.35 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Repairman's Cape; chest: Beguiler Robes; hands: Hotshot Pilot's Gloves; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: The Queen's Jewel; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-000000000000000000-5032503315400000)

Set DPS (verified): 64.3. Weights run: 5.4s. Verify run: 2.9s. 592 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=2.362 ± 0.020, spirit=2.863 ± 0.016, mp5=7.383 ± 0.023, crit=0.375 ± 0.020 per rating point (14 rating = 1%, 5.256 per %), spell_haste=not significant (0.232 ± 0.065)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 87.9 healing_power points (3.49 DPS) | yes | Whitemane's Chapeau (7720, -0.06 DPS, sim-verified) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.72 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.76 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.03 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.32 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 55.0 healing_power points (2.18 DPS) | yes | Earthen Silk Shoulders (254033, -0.18 DPS, sim-verified) [crafted]; Mistscape Mantle (4734, -0.58 DPS) [dungeon]; Batwing Mantle (6697, -0.58 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 37.4 healing_power points (1.49 DPS) | yes | Caretaker's Cape (19532, -0.10 DPS, sim-verified) [rep]; Sergeant Major's Cape (16336, -0.24 DPS) [pvp]; Ceremonial Centaur Blanket (6789, -0.28 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 80.9 healing_power points (3.21 DPS) | yes | Doomsayer's Robe (4746, -0.05 DPS, sim-verified) [quest]; Icemail Jerkin (1981, -0.48 DPS) [world_drop]; Deathchill Armor (10764, -0.66 DPS) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 38.1 healing_power points (1.51 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.55 DPS) [crafted]; Silkstream Cuffs (16791, -0.59 DPS) [quest] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 50.8 healing_power points (2.02 DPS) | yes | Gloves of Old (9395, -0.06 DPS, sim-verified) [world_drop]; Mender's Leather Gloves (252530, -0.34 DPS) [crafted]; Warden's Gloves (14606, -0.39 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 58.1 healing_power points (2.31 DPS) | yes | Mender's Leather Belt (252523, +0.00 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.53 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Misplaced Pantaloons (276201) | Friz Frazzlespark [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Warchief Kilt (7760, -0.09 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -0.40 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.41 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Furen's Boots (13100, -0.03 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -0.36 DPS) [dungeon] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 44.3 healing_power points (1.76 DPS) | yes | Welken Ring (5011, -0.49 DPS) [world_drop]; The Queen's Jewel (13094, -0.66 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 36.6 healing_power points (1.45 DPS) | yes | Welken Ring (5011, +0.00 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.36 DPS) [world_drop]; Voodoo Band (1996, -0.45 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Staff of Jordan (873, -0.42 DPS) [world_drop]; Wind Spirit Staff (6689, -0.53 DPS) [dungeon] |
| off_hand | Ravager's Shield (14777) | World drop [world_drop] | 33.7 healing_power points (1.34 DPS) | yes | Beacon of Hope (9393, +0.00 DPS, sim-verified) [dungeon]; Mordresh's Lifeless Skull (10770, -0.09 DPS) [dungeon]; Orb of Souls (249395, -0.14 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Misplaced Pantaloons; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; off_hand: Ravager's Shield

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 0000000000000000-000000000000000000-5032503315513140)

Set DPS (verified): 99.3. Weights run: 6.2s. Verify run: 3.2s. 762 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.023, intellect=2.751 ± 0.026, spirit=3.030 ± 0.025, mp5=8.085 ± 0.029, crit=0.480 ± 0.026 per rating point (14 rating = 1%, 6.718 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | 104.1 healing_power points (4.62 DPS) | yes | Gemburst Circlet (10751, -0.06 DPS, sim-verified) [quest]; Soulcatcher Halo (10630, -0.22 DPS) [dungeon]; Braincage (12549, -0.36 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkmoon Necklace (19303, -0.14 DPS, sim-verified) [vendor]; Glowing Eye of Mordresh (10769, -0.16 DPS) [dungeon]; Gemshard Heart (17707, -0.22 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 79.3 healing_power points (3.52 DPS) | yes | Lead Surveyor's Mantle (11842, -0.06 DPS, sim-verified) [dungeon]; Dregmetal Spaulders (11722, -0.34 DPS) [dungeon]; Living Shoulders (15061, -0.39 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Featherskin Cape (10843, -0.20 DPS, sim-verified) [world]; Imperial Red Cloak (8248, -0.50 DPS) [world_drop]; Caretaker's Cape (19531, -0.60 DPS) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 137.7 healing_power points (6.11 DPS) | yes | Ghostweave Vest (14141, -0.40 DPS, sim-verified) [crafted]; Robes of Insight (940, -1.04 DPS) [world_drop]; Vestments of the Atal'ai Prophet (10806, -1.14 DPS) [dungeon] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 59.4 healing_power points (2.64 DPS) | yes | Mender's Leather Bracers (252543, -0.09 DPS, sim-verified) [crafted]; Mender's Mail Bracers (252573, -0.13 DPS) [crafted]; Nethergeld Cuffs (254061, -0.22 DPS) [crafted] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 108.8 healing_power points (4.83 DPS) | yes | Stonerender Gauntlets (17007, +0.00 DPS, sim-verified) [world_drop]; Mender's Leather Gauntlets (252551, -1.14 DPS) [crafted]; Mender's Mail Gauntlets (252587, -1.14 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Gilded Cord (254037, -0.04 DPS) [crafted]; Bloodlust Belt (14803, -0.08 DPS, sim-verified) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 104.1 healing_power points (4.62 DPS) | yes | Dalewind Trousers (13008, -0.18 DPS, sim-verified) [world_drop]; Windscale Sarong (10842, -0.71 DPS) [world]; Jinxed Hoodoo Kilt (9474, -0.84 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 82.6 healing_power points (3.66 DPS) | yes | Mistwalker Boots (10629, -0.27 DPS) [dungeon]; Coldstone Slippers (18697, -0.52 DPS) [dungeon]; Furen's Boots (13100, -0.66 DPS) [world_drop] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 56.6 healing_power points (2.51 DPS) | yes | Eye of Adaegus (5266, -0.17 DPS) [world_drop]; Snake Hoop (6750, -0.72 DPS) [quest]; Cyclopean Band (11824, -0.72 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 56.4 healing_power points (2.50 DPS) | yes | Eye of Adaegus (5266, -0.18 DPS, sim-verified) [world_drop]; Snake Hoop (6750, -0.71 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -4.04 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.30 DPS) [quest]; Uther's Strength (11302, -4.58 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.27 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.54 DPS) [quest]; Uther's Strength (11302, -0.81 DPS) [world_drop] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -0.09 DPS) [world_drop]; Barman Shanker (12791, -0.56 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -0.89 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Helm of Exile; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Embrace of the Wind Serpent; wrist: Aristocratic Cuffs; hands: Grasp of The Five Thunders; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Soulkeeper

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 231.7. Weights run: 10.3s. Verify run: 4.8s. 1751 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.125, intellect=4.812 ± 0.061, spirit=4.336 ± 0.053, mp5=11.044 ± 0.061, crit=0.769 ± 0.054 per rating point (14 rating = 1%, 10.766 per %), spell_haste=not significant (0.236 ± 0.232)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Crown of the Penitent (13216, -0.92 DPS) [quest]; Devout Crown (16693, -1.02 DPS) [dungeon]; Gnomish Turban of Psychic Might (21517, -2.20 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 141.2 healing_power points (7.30 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.58 DPS) [world_drop]; Lady Maye's Pendant (14558, -1.50 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -2.10 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulstealer Mantle (13374, -0.32 DPS) [dungeon]; Sunderseer Mantle (13185, -0.70 DPS) [dungeon]; Argent Elite Shoulders (227888, -2.71 DPS, sim-verified) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 109.8 healing_power points (5.67 DPS) | yes | Darkspear Raider's Cloak (272063, -0.35 DPS) [vendor]; Shroud of the Exile (15421, -0.37 DPS) [quest]; Faded Hakkari Cloak (20218, -0.73 DPS, sim-verified) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, -0.37 DPS) [crafted]; Alanna's Embrace (13314, -0.46 DPS) [dungeon]; Tunic of Undead Slaying (23089, -9.07 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, -0.27 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.45 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -7.55 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Hands of the Exalted Herald (12554, -0.33 DPS) [dungeon]; Devout Gloves (16692, -0.88 DPS) [dungeon]; Grasp of The Five Thunders (227014, -4.22 DPS, sim-verified) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 159.1 healing_power points (8.22 DPS) | yes | Whipvine Cord (18327, -0.96 DPS) [dungeon]; Elunarian Belt (14465, -1.01 DPS) [world_drop]; Devout Belt (16696, -1.52 DPS, sim-verified) [dungeon] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 244.3 healing_power points (12.62 DPS) | yes | Padre's Trousers (18386, -2.06 DPS) [dungeon]; Legplates of the Chromatic Defier (12945, -2.07 DPS) [quest]; Ghostloom Leggings (14545, -3.27 DPS, sim-verified) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 191.4 healing_power points (9.89 DPS) | yes | Incandescent Mooncloth Boots (227862, -2.02 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -2.25 DPS) [crafted]; Faith Healer's Boots (22247, -2.87 DPS) [dungeon] |
| finger1 | Ring of Demonic Guile (18314) | Dire Maul: Alzzin the Wildshaper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.11 DPS) [dungeon]; Emerald Flame Ring (18395, -0.36 DPS) [dungeon]; Naglering (11669, -6.06 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.11 DPS) [dungeon]; Emerald Flame Ring (18395, -0.35 DPS) [dungeon]; Naglering (11669, -7.71 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindtap Talisman (18371, -2.28 DPS) [dungeon]; Shard of the Splithooves (10659, -5.13 DPS) [quest]; Ankh of Life (1713, -5.87 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Mindtap Talisman (18371, +0.00 DPS) [dungeon] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quel'dorai Channeling Rod (18311, -0.31 DPS) [dungeon]; Argent Crusader (13249, -1.44 DPS) [quest]; Hand of Edward the Odd (2243, -5.65 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Frostweaver Cape; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Leggings of Arcana; feet: Greaves of The Five Thunders; finger1: Ring of Demonic Guile; finger2: Band of Piety; trinket2: Serenity Field; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 370.5. Weights run: 9.0s. Verify run: 3.6s. 1751 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.272, intellect=2.675 ± 0.051, spirit=2.238 ± 0.036, mp5=6.265 ± 0.052, crit=0.784 ± 0.045 per rating point (14 rating = 1%, 10.977 per %), spell_haste=not significant (-0.366 ± 0.372)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 152.5 healing_power points (15.68 DPS) | yes | Crown of The Five Thunders (227013, -2.80 DPS) [vendor]; Gnomish Turban of Psychic Might (21517, -3.18 DPS, sim-verified) [quest]; Devout Crown (16693, -3.57 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 90.2 healing_power points (9.27 DPS) | yes | Lady Maye's Pendant (14558, -1.65 DPS, sim-verified) [world_drop]; Jeweled Amulet of Cainwyn (1443, -2.02 DPS) [world_drop]; Drake Tooth Necklace (21531, -3.30 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 119.0 healing_power points (12.23 DPS) | yes | Devout Mantle (16695, +0.00 DPS, sim-verified) [dungeon]; Mantle of The Five Thunders (227011, -3.63 DPS) [vendor]; Royal Cap Spaulders (14548, -4.09 DPS) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 68.7 healing_power points (7.07 DPS) | yes | Frostweaver Cape (12968, -1.01 DPS) [dungeon]; Gracious Cape (18743, -1.28 DPS) [dungeon]; Faded Hakkari Cloak (20218, -2.01 DPS, sim-verified) [quest] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | sim-verified (370.8 DPS) | yes | Mooncloth Vest (14138, -0.24 DPS) [crafted]; Alanna's Embrace (13314, -0.63 DPS) [dungeon]; Tunic of Undead Slaying (23089, -2.68 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (370.8 DPS) | yes | Bracers of Mending (23129, -0.32 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.67 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -3.10 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 96.2 healing_power points (9.89 DPS) | yes | Hands of the Exalted Herald (12554, -0.16 DPS) [dungeon]; Grasp of The Five Thunders (227014, -0.26 DPS) [vendor]; Devout Gloves (16692, -1.69 DPS) [dungeon] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 92.7 healing_power points (9.53 DPS) | yes | Wisdom of the Timbermaw (19047, -0.42 DPS) [crafted]; Sash of The Five Thunders (227010, -0.78 DPS) [vendor]; Belt of Tiny Heads (20217, -1.56 DPS, sim-verified) [quest] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 138.6 healing_power points (14.25 DPS) | yes | Padre's Trousers (18386, -0.57 DPS) [dungeon]; Devout Skirt (16694, -2.37 DPS) [dungeon]; Ghostloom Leggings (14545, -2.45 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 112.8 healing_power points (11.59 DPS) | yes | Greaves of The Five Thunders (227015, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -2.29 DPS) [crafted]; Faith Healer's Boots (22247, -2.86 DPS) [dungeon] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (370.8 DPS) | yes | Ring of Demonic Guile (18314, -0.97 DPS) [dungeon]; Band of Mending (22334, -1.10 DPS) [dungeon]; Naglering (11669, -3.26 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (370.8 DPS) | yes | Ring of Demonic Guile (18314, -0.07 DPS) [dungeon]; Band of Mending (22334, -0.19 DPS) [dungeon]; Naglering (11669, -2.80 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (370.8 DPS) | yes | Mindtap Talisman (18371, -0.58 DPS, sim-verified) [dungeon]; Serenity Field (272439, -5.34 DPS) [vendor]; Briarwood Reed (12930, -6.68 DPS) [dungeon] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (370.8 DPS) | yes | Mindtap Talisman (18371, +0.00 DPS, sim-verified) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -0.88 DPS) [dungeon] |
| main_hand | Quel'dorai Channeling Rod (18311) | Dire Maul: Lethtendris [dungeon] | sim-verified (370.8 DPS) | yes | Staff of Hale Magefire (13000, -0.01 DPS) [world_drop]; Hammer of the Grand Crusader (18717, -0.46 DPS) [dungeon]; Hand of Edward the Odd (2243, -1.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Whipvine Cord; legs: Leggings of Arcana; feet: Incandescent Mooncloth Boots; finger1: Band of Piety; finger2: Emerald Flame Ring; trinket2: Shard of the Splithooves; main_hand: Quel'dorai Channeling Rod

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (tauren, 0000000000000000-000000000000000000-5032100000000000)

Set DPS (verified): 27.0. Weights run: 4.4s. Verify run: 2.3s. 205 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.751 ± 0.009, spirit=1.906 ± 0.006, mp5=5.250 ± 0.009, crit=0.123 ± 0.006 per rating point (14 rating = 1%, 1.728 per %), spell_haste=not significant (0.054 ± 0.020)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.2 healing_power points (0.83 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Stormrider's Leather Hood (252506, -0.13 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.13 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 7.6 healing_power points (0.31 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.8 healing_power points (0.65 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Mantle (4709, -0.34 DPS) [world_drop]; Prospector's Pads (14566, -0.34 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 12.8 healing_power points (0.53 DPS) | yes | Regent's Cloak (5969, +0.00 DPS, sim-verified) [world]; Traveler's Shawl (277289, -0.14 DPS) [quest]; Spirit Cloak (4792, -0.21 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.5 healing_power points (1.22 DPS) | yes | Robe of the Moccasin (6465, +0.00 DPS, sim-verified) [dungeon]; Filigreed Pristine Gown (253901, -0.08 DPS) [crafted]; Corsair's Overshirt (5202, -0.13 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 24.5 healing_power points (1.01 DPS) | yes | Drakewing Bands (12999, +0.00 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.58 DPS) [quest]; Crystalline Cuffs (14148, -0.63 DPS) [dungeon] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 21.8 healing_power points (0.90 DPS) | yes | Wisdom's Leather Gloves (252499, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.22 DPS) [world_drop]; Pristine Gloves (253913, -0.23 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.0 healing_power points (0.74 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.17 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.1 healing_power points (1.57 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.60 DPS) [world_drop]; Ghastly Trousers (15449, -0.78 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 19.8 healing_power points (0.82 DPS) | yes | Black Whelp Slippers (252424, +0.00 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.23 DPS) [crafted]; Smoldering Boots (3076, -0.34 DPS) [world] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 14.9 healing_power points (0.62 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Advisor's Ring (20426, -0.18 DPS) [rep]; Deep Fathom Ring (6463, -0.22 DPS) [dungeon] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (0.5 DPS) | yes | Advisor's Ring (20426, -0.00 DPS) [rep]; Band of Purification (12996, -0.03 DPS, sim-verified) [world_drop]; Deep Fathom Ring (6463, -0.04 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 27.0 healing_power points (1.12 DPS) | yes | Staff of the Blessed Seer (2271, +0.00 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (20425, -0.18 DPS) [pvp]; Gnarled Necromancer's Staff (251534, -0.39 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (tauren, 0000000000000000-000000000000000000-5032503300000000)

Set DPS (verified): 44.1. Weights run: 4.5s. Verify run: 2.4s. 349 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=2.418 ± 0.014, spirit=2.475 ± 0.013, mp5=6.455 ± 0.015, crit=0.185 ± 0.009 per rating point (14 rating = 1%, 2.592 per %), spell_haste=not significant (0.199 ± 0.050)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 56.0 healing_power points (2.21 DPS) | yes | Whisperwind Headdress (6688, -0.06 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -0.32 DPS) [world_drop]; Nightsky Cowl (4039, -0.58 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 19.6 healing_power points (0.77 DPS) | yes | Necklace of Harmony (5180, -0.06 DPS, sim-verified) [world]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Pendant of Myzrael (4614, -0.19 DPS) [dungeon] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 39.0 healing_power points (1.54 DPS) | yes | Nightsky Mantle (4718, -0.06 DPS, sim-verified) [world_drop]; Ghostly Mantle (3324, -0.30 DPS) [quest]; Talbar Mantle (10657, -0.46 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-verified (1.2 DPS) | yes | Glowing Thresher Cape (6901, -0.07 DPS, sim-verified) [dungeon]; Battle Healer's Cloak (19529, -0.15 DPS) [rep]; Cloak of Rot (4462, -0.29 DPS) [world] |
| chest | Beguiler Robes (7728) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 48.8 healing_power points (1.93 DPS) | yes | Wisdom's Leather Tunic (252511, +0.00 DPS, sim-verified) [crafted]; Pristine Gown (253961, -0.08 DPS) [crafted]; Death Speaker Robes (6682, -0.09 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 31.5 healing_power points (1.24 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.48 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.48 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 41.8 healing_power points (1.65 DPS) | yes | Hotshot Pilot's Gloves (9491, +0.00 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.49 DPS) [quest]; Tattered Mittens (270030, -0.67 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 48.4 healing_power points (1.91 DPS) | yes | Silver-lined Belt (13011, +0.00 DPS, sim-verified) [world_drop]; Lilac Sash (6780, -0.76 DPS) [quest]; Highlander's Mail Girdle (20120, -0.76 DPS) [vendor] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 55.3 healing_power points (2.18 DPS) | yes | Pristine Leggings (253987, -0.16 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 46.8 healing_power points (1.85 DPS) | yes | Soggy Boots (274747, -0.08 DPS, sim-verified) [vendor]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Frothing Slippers (254003, -0.79 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 34.3 healing_power points (1.35 DPS) | yes | The Queen's Jewel (13094, -0.38 DPS) [world_drop]; Black Pearl Ring (6332, -0.57 DPS) [world]; Sea Giant's Toe Ring (274746, -0.64 DPS) [vendor] |
| finger2 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 25.8 healing_power points (1.02 DPS) | yes | The Queen's Jewel (13094, +0.00 DPS, sim-verified) [world_drop]; Black Pearl Ring (6332, -0.24 DPS) [world]; Sea Giant's Toe Ring (274746, -0.31 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, +0.00 DPS, sim-verified) [dungeon]; Wind Spirit Staff (6689, -0.74 DPS) [dungeon]; Advisor's Gnarled Staff (19569, -1.00 DPS) [pvp] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 27.9 healing_power points (1.10 DPS) | yes | Defective Samophlange (274743, +0.00 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -0.33 DPS) [world_drop]; Orb of Mistmantle (13031, -0.35 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-000000000000000000-5032503315400000)

Set DPS (verified): 63.2. Weights run: 5.4s. Verify run: 2.9s. 555 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=2.362 ± 0.020, spirit=2.863 ± 0.016, mp5=7.383 ± 0.023, crit=0.375 ± 0.020 per rating point (14 rating = 1%, 5.256 per %), spell_haste=not significant (0.232 ± 0.065)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 87.9 healing_power points (3.49 DPS) | yes | Whitemane's Chapeau (7720, -0.09 DPS, sim-verified) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.72 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.76 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 43.3 healing_power points (1.72 DPS) | yes | Necklace of Calisea (1714, -0.27 DPS) [dungeon]; Triune Amulet (7722, -0.27 DPS) [dungeon]; Amberglow Talisman (10824, -0.58 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 55.0 healing_power points (2.18 DPS) | yes | Earthen Silk Shoulders (254033, -0.20 DPS, sim-verified) [crafted]; Mistscape Mantle (4734, -0.58 DPS) [dungeon]; Batwing Mantle (6697, -0.58 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 37.4 healing_power points (1.49 DPS) | yes | Cloak of Blight (6832, -0.15 DPS, sim-verified) [quest]; Battle Healer's Cloak (19528, -0.20 DPS) [rep]; First Sergeant's Cloak (16340, -0.24 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 80.9 healing_power points (3.21 DPS) | yes | Doomsayer's Robe (4746, -0.08 DPS, sim-verified) [quest]; Icemail Jerkin (1981, -0.48 DPS) [world_drop]; Deathchill Armor (10764, -0.66 DPS) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 38.1 healing_power points (1.51 DPS) | yes | Mindthrust Bracers (1974, -0.16 DPS) [dungeon]; Windtalker's Wristguards (19584, -0.38 DPS) [pvp]; Dryad's Wrist Bindings (19597, -0.38 DPS) [pvp] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 50.8 healing_power points (2.02 DPS) | yes | Gloves of Old (9395, -0.09 DPS, sim-verified) [world_drop]; Mender's Leather Gloves (252530, -0.34 DPS) [crafted]; Warden's Gloves (14606, -0.39 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 58.1 healing_power points (2.31 DPS) | yes | Mender's Leather Belt (252523, -0.05 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.53 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Misplaced Pantaloons (276201) | Friz Frazzlespark [vendor] | sim-verified (1.8 DPS) | yes | Warchief Kilt (7760, -0.06 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -0.40 DPS) [crafted]; Wisdom's Leather Leggings (252519, -0.41 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | 62.3 healing_power points (2.48 DPS) | yes | Mender's Leather Shoes (252533, -0.07 DPS) [crafted]; Mender's Mail Boots (252565, -0.07 DPS) [crafted]; Thoughtcast Boots (10578, -0.43 DPS) [dungeon] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 44.3 healing_power points (1.76 DPS) | yes | Welken Ring (5011, -0.49 DPS) [world_drop]; The Queen's Jewel (13094, -0.66 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 36.6 healing_power points (1.45 DPS) | yes | Welken Ring (5011, +0.00 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.36 DPS) [world_drop]; Voodoo Band (1996, -0.45 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Advisor's Gnarled Staff (19568, -0.10 DPS) [pvp]; Staff of Jordan (873, -0.42 DPS) [world_drop] |
| off_hand | Ravager's Shield (14777) | World drop [world_drop] | 33.7 healing_power points (1.34 DPS) | yes | Beacon of Hope (9393, -0.07 DPS, sim-verified) [dungeon]; Mordresh's Lifeless Skull (10770, -0.09 DPS) [dungeon]; Orb of Souls (249395, -0.14 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Misplaced Pantaloons; feet: Furen's Boots; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; off_hand: Ravager's Shield

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (tauren, 0000000000000000-000000000000000000-5032503315513140)

Set DPS (verified): 99.2. Weights run: 6.2s. Verify run: 3.2s. 704 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.023, intellect=2.751 ± 0.026, spirit=3.030 ± 0.025, mp5=8.085 ± 0.029, crit=0.480 ± 0.026 per rating point (14 rating = 1%, 6.718 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | 104.1 healing_power points (4.62 DPS) | yes | Gemburst Circlet (10751, +0.00 DPS, sim-verified) [quest]; Soulcatcher Halo (10630, -0.22 DPS) [dungeon]; Braincage (12549, -0.36 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkmoon Necklace (19303, -0.10 DPS, sim-verified) [vendor]; Glowing Eye of Mordresh (10769, -0.16 DPS) [dungeon]; Gemshard Heart (17707, -0.22 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 79.3 healing_power points (3.52 DPS) | yes | Lead Surveyor's Mantle (11842, -0.07 DPS, sim-verified) [dungeon]; Dregmetal Spaulders (11722, -0.34 DPS) [dungeon]; Living Shoulders (15061, -0.39 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Featherskin Cape (10843, -0.24 DPS, sim-verified) [world]; Imperial Red Cloak (8248, -0.50 DPS) [world_drop]; Battle Healer's Cloak (19527, -0.60 DPS) [rep] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 137.7 healing_power points (6.11 DPS) | yes | Ghostweave Vest (14141, -0.32 DPS, sim-verified) [crafted]; Robes of Insight (940, -1.04 DPS) [world_drop]; Vestments of the Atal'ai Prophet (10806, -1.14 DPS) [dungeon] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 59.4 healing_power points (2.64 DPS) | yes | Mender's Leather Bracers (252543, -0.09 DPS, sim-verified) [crafted]; Mender's Mail Bracers (252573, -0.13 DPS) [crafted]; Nethergeld Cuffs (254061, -0.22 DPS) [crafted] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 108.8 healing_power points (4.83 DPS) | yes | Stonerender Gauntlets (17007, +0.00 DPS, sim-verified) [world_drop]; Mender's Leather Gauntlets (252551, -1.14 DPS) [crafted]; Mender's Mail Gauntlets (252587, -1.14 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Gilded Cord (254037, -0.04 DPS) [crafted]; Bloodlust Belt (14803, -0.06 DPS, sim-verified) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 104.1 healing_power points (4.62 DPS) | yes | Dalewind Trousers (13008, -0.19 DPS, sim-verified) [world_drop]; Windscale Sarong (10842, -0.71 DPS) [world]; Jinxed Hoodoo Kilt (9474, -0.84 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 82.6 healing_power points (3.66 DPS) | yes | Mistwalker Boots (10629, +0.00 DPS, sim-verified) [dungeon]; Coldstone Slippers (18697, -0.52 DPS) [dungeon]; Furen's Boots (13100, -0.66 DPS) [world_drop] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 56.6 healing_power points (2.51 DPS) | yes | Eye of Adaegus (5266, -0.17 DPS) [world_drop]; Snake Hoop (6750, -0.72 DPS) [quest]; Cyclopean Band (11824, -0.72 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 56.4 healing_power points (2.50 DPS) | yes | Eye of Adaegus (5266, -0.19 DPS, sim-verified) [world_drop]; Snake Hoop (6750, -0.71 DPS) [quest]; Cyclopean Band (11824, -0.71 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -4.04 DPS) [quest]; Uther's Strength (11302, -4.58 DPS) [world_drop]; Alchemists' Stone (13503, -4.84 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS, sim-verified) [quest]; Uther's Strength (11302, -0.81 DPS) [world_drop]; Alchemists' Stone (13503, -1.08 DPS) [crafted] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -0.09 DPS) [world_drop]; Barman Shanker (12791, -0.49 DPS, sim-verified) [dungeon]; Resurgence Rod (17743, -0.89 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Helm of Exile; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Embrace of the Wind Serpent; wrist: Aristocratic Cuffs; hands: Grasp of The Five Thunders; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Soulkeeper

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (tauren, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 231.1. Weights run: 10.3s. Verify run: 4.8s. 1672 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.125, intellect=4.812 ± 0.061, spirit=4.336 ± 0.053, mp5=11.044 ± 0.061, crit=0.769 ± 0.054 per rating point (14 rating = 1%, 10.766 per %), spell_haste=not significant (0.236 ± 0.232)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Crown of the Penitent (13216, -0.92 DPS) [quest]; Devout Crown (16693, -1.02 DPS) [dungeon]; Gnomish Turban of Psychic Might (21517, -2.66 DPS, sim-verified) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 141.2 healing_power points (7.30 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.58 DPS) [world_drop]; Lady Maye's Pendant (14558, -1.05 DPS, sim-verified) [world_drop]; Tooth of Gnarr (13141, -2.10 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulstealer Mantle (13374, -0.32 DPS) [dungeon]; Sunderseer Mantle (13185, -0.70 DPS) [dungeon]; Argent Elite Shoulders (227888, -3.76 DPS, sim-verified) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 109.8 healing_power points (5.67 DPS) | yes | Faded Hakkari Cloak (20218, -0.35 DPS, sim-verified) [quest]; Darkspear Raider's Cloak (272063, -0.35 DPS) [vendor]; Shroud of the Exile (15421, -0.37 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, -0.37 DPS) [crafted]; Alanna's Embrace (13314, -0.46 DPS) [dungeon]; Tunic of Undead Slaying (23089, -8.68 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, -0.27 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.45 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -6.72 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Hands of the Exalted Herald (12554, -0.33 DPS) [dungeon]; Devout Gloves (16692, -0.88 DPS) [dungeon]; Grasp of The Five Thunders (227014, -4.74 DPS, sim-verified) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 159.1 healing_power points (8.22 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Whipvine Cord (18327, -0.96 DPS) [dungeon]; Devout Belt (16696, -1.20 DPS, sim-verified) [dungeon] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 244.3 healing_power points (12.62 DPS) | yes | Padre's Trousers (18386, -2.06 DPS) [dungeon]; Legplates of the Chromatic Defier (12945, -2.07 DPS) [quest]; Ghostloom Leggings (14545, -2.71 DPS, sim-verified) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 191.4 healing_power points (9.89 DPS) | yes | Incandescent Mooncloth Boots (227862, -2.09 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -2.25 DPS) [crafted]; Boots of The Five Thunders (22096, -2.38 DPS) [quest] |
| finger1 | Ring of Demonic Guile (18314) | Dire Maul: Alzzin the Wildshaper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.11 DPS) [dungeon]; Emerald Flame Ring (18395, -0.36 DPS) [dungeon]; Naglering (11669, -5.32 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.11 DPS) [dungeon]; Emerald Flame Ring (18395, -0.35 DPS) [dungeon]; Naglering (11669, -7.19 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -2.26 DPS, sim-verified) [dungeon]; Shard of the Splithooves (10659, -5.13 DPS) [quest]; Ankh of Life (1713, -5.87 DPS) [world_drop] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Mindtap Talisman (18371, +0.00 DPS, sim-verified) [dungeon] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quel'dorai Channeling Rod (18311, -0.31 DPS) [dungeon]; Argent Crusader (13249, -1.44 DPS) [quest]; Hand of Edward the Odd (2243, -4.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Frostweaver Cape; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Leggings of Arcana; feet: Greaves of The Five Thunders; finger1: Ring of Demonic Guile; finger2: Band of Piety; trinket2: Serenity Field; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (tauren, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 370.2. Weights run: 9.0s. Verify run: 3.5s. 1672 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.272, intellect=2.675 ± 0.051, spirit=2.238 ± 0.036, mp5=6.265 ± 0.052, crit=0.784 ± 0.045 per rating point (14 rating = 1%, 10.977 per %), spell_haste=not significant (-0.366 ± 0.372)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 152.5 healing_power points (15.68 DPS) | yes | Crown of The Five Thunders (227013, -2.80 DPS) [vendor]; Gnomish Turban of Psychic Might (21517, -3.15 DPS, sim-verified) [quest]; Devout Crown (16693, -3.57 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 90.2 healing_power points (9.27 DPS) | yes | Lady Maye's Pendant (14558, -1.51 DPS, sim-verified) [world_drop]; Jeweled Amulet of Cainwyn (1443, -2.02 DPS) [world_drop]; Drake Tooth Necklace (21531, -3.30 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 119.0 healing_power points (12.23 DPS) | yes | Devout Mantle (16695, -3.26 DPS) [dungeon]; Mantle of The Five Thunders (227011, -3.63 DPS) [vendor]; Champion's Mail Epaulets (227166, -4.08 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 68.7 healing_power points (7.07 DPS) | yes | Frostweaver Cape (12968, -1.01 DPS) [dungeon]; Gracious Cape (18743, -1.28 DPS) [dungeon]; Faded Hakkari Cloak (20218, -2.04 DPS, sim-verified) [quest] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | sim-verified (370.7 DPS) | yes | Mooncloth Vest (14138, -0.24 DPS) [crafted]; Alanna's Embrace (13314, -0.63 DPS) [dungeon]; Tunic of Undead Slaying (23089, -3.06 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (370.7 DPS) | yes | Bracers of Mending (23129, -0.32 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.67 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -3.04 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 96.2 healing_power points (9.89 DPS) | yes | Hands of the Exalted Herald (12554, -0.16 DPS) [dungeon]; Grasp of The Five Thunders (227014, -0.26 DPS) [vendor]; Devout Gloves (16692, -1.69 DPS) [dungeon] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 92.7 healing_power points (9.53 DPS) | yes | General's Mail Waistband (16575, -0.35 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.42 DPS) [crafted]; Belt of Tiny Heads (20217, -1.55 DPS, sim-verified) [quest] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 138.6 healing_power points (14.25 DPS) | yes | Padre's Trousers (18386, -0.57 DPS) [dungeon]; Devout Skirt (16694, -2.37 DPS) [dungeon]; Ghostloom Leggings (14545, -2.45 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 112.8 healing_power points (11.59 DPS) | yes | Greaves of The Five Thunders (227015, -0.07 DPS) [vendor]; Mooncloth Boots (15802, -2.29 DPS) [crafted]; Faith Healer's Boots (22247, -2.86 DPS) [dungeon] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (370.7 DPS) | yes | Ring of Demonic Guile (18314, -0.97 DPS) [dungeon]; Band of Mending (22334, -1.10 DPS) [dungeon]; Naglering (11669, -3.23 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-verified (370.7 DPS) | yes | Ring of Demonic Guile (18314, -0.07 DPS) [dungeon]; Band of Mending (22334, -0.19 DPS) [dungeon]; Naglering (11669, -2.88 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (370.7 DPS) | yes | Mindtap Talisman (18371, -0.70 DPS, sim-verified) [dungeon]; Serenity Field (272439, -5.34 DPS) [vendor]; Briarwood Reed (12930, -6.68 DPS) [dungeon] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (370.7 DPS) | yes | Mindtap Talisman (18371, +0.00 DPS, sim-verified) [dungeon]; Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -0.88 DPS) [dungeon] |
| main_hand | Quel'dorai Channeling Rod (18311) | Dire Maul: Lethtendris [dungeon] | sim-verified (370.7 DPS) | yes | Staff of Hale Magefire (13000, -0.01 DPS) [world_drop]; Hammer of the Grand Crusader (18717, -0.46 DPS) [dungeon]; Hand of Edward the Odd (2243, -1.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Whipvine Cord; legs: Leggings of Arcana; feet: Incandescent Mooncloth Boots; finger1: Band of Piety; finger2: Emerald Flame Ring; trinket2: Shard of the Splithooves; main_hand: Quel'dorai Channeling Rod

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

