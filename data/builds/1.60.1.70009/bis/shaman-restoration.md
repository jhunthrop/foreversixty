# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5032100000000000)

Set DPS (verified): 27.0. Weights run: 4.6s. Verify run: 3.1s. 225 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.856 ± 0.007, spirit=1.848 ± 0.003, mp5=4.989 ± 0.007, crit=0.124 ± 0.006 per rating point (14 rating = 1%, 1.731 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.4 healing_power points (0.87 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.14 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.16 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 7.4 healing_power points (0.32 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.7 healing_power points (0.71 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 12.7 healing_power points (0.54 DPS) | yes | Regent's Cloak (5969, +0.00 DPS, sim-verified) [world]; Sanguine Cape (14376, -0.22 DPS) [world_drop]; Seer's Cape (6378, -0.23 DPS) [dungeon] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.8 healing_power points (1.27 DPS) | yes | Filigreed Pristine Gown (253901, +0.00 DPS, sim-verified) [crafted]; Robe of the Moccasin (6465, -0.10 DPS) [dungeon]; Corsair's Overshirt (5202, -0.17 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 24.2 healing_power points (1.03 DPS) | yes | Drakewing Bands (12999, +0.00 DPS, sim-verified) [world_drop]; Owl Bracers (4796, -0.64 DPS) [vendor]; Bright Bracers (3647, -0.72 DPS) [world_drop] |
| hands | Magefist Gloves (12977) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gloves (253913, -0.00 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.02 DPS, sim-verified) [crafted]; Bright Gloves (3066, -0.08 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.4 healing_power points (0.79 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.15 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.5 healing_power points (1.64 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.62 DPS) [world_drop]; Dreamer's Leggings (270016, -0.77 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.3 healing_power points (0.86 DPS) | yes | Black Whelp Slippers (252424, +0.00 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.24 DPS) [crafted]; Kimbra Boots (6191, -0.31 DPS) [quest] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 14.8 healing_power points (0.63 DPS) | yes | Band of Purification (12996, -0.16 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.21 DPS) [rep]; Deep Fathom Ring (6463, -0.24 DPS) [dungeon] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.1 healing_power points (0.47 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Lorekeeper's Ring (20431, -0.05 DPS) [rep]; Deep Fathom Ring (6463, -0.08 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (0.4 DPS) | yes | Staff of the Blessed Seer (2271, -0.00 DPS) [dungeon]; Staff of Westfall (2042, -0.03 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.16 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Magefist Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 0000000000000000-000000000000000000-5032503300000000)

Set DPS (verified): 43.0. Weights run: 4.7s. Verify run: 3.2s. 366 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.995 ± 0.006, spirit=2.468 ± 0.013, mp5=6.283 ± 0.014, crit=0.185 ± 0.009 per rating point (14 rating = 1%, 2.592 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.2 healing_power points (2.03 DPS) | yes | Whisperwind Headdress (6688, +0.00 DPS, sim-verified) [dungeon]; Holy Shroud (2721, -0.06 DPS) [world_drop]; Nightsky Cowl (4039, -0.53 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.9 healing_power points (0.73 DPS) | yes | Necklace of Harmony (5180, +0.00 DPS, sim-verified) [world]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Pendant of Myzrael (4614, -0.13 DPS) [dungeon] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 34.3 healing_power points (1.41 DPS) | yes | Mantle of Honor (3560, +0.00 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.25 DPS) [world_drop]; Faerie Mantle (5820, -0.31 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 26.7 healing_power points (1.10 DPS) | yes | Prelacy Cape (7004, -0.04 DPS) [quest]; Repairman's Cape (9605, -0.06 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.14 DPS) [vendor] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.8 healing_power points (1.84 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.05 DPS) [dungeon]; Death Speaker Robes (6682, -0.12 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 28.8 healing_power points (1.19 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.53 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan: Caverndeep Burrower [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of Old (9395, -0.04 DPS, sim-verified) [world_drop]; Zodiac Gloves (7106, -0.13 DPS) [quest]; Silver-thread Gloves (6393, -0.25 DPS) [world_drop] |
| waist | Silver-lined Belt (13011) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Leather Belt (252523, -0.02 DPS, sim-verified) [crafted]; Highlander's Lizardhide Girdle (20105, -0.16 DPS) [rep]; Highlander's Mail Girdle (20120, -0.16 DPS) [vendor] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.2 healing_power points (2.15 DPS) | yes | Pristine Leggings (253987, -0.16 DPS) [crafted]; Earthen Leggings (253999, -0.42 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.8 healing_power points (1.80 DPS) | yes | Soggy Boots (274747, +0.00 DPS, sim-verified) [vendor]; Acidic Walkers (9454, -0.74 DPS) [dungeon]; Frothing Slippers (254003, -0.82 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 31.2 healing_power points (1.29 DPS) | yes | The Queen's Jewel (13094, -0.31 DPS) [world_drop]; Black Pearl Ring (6332, -0.51 DPS) [world]; Sea Giant's Toe Ring (274746, -0.55 DPS) [vendor] |
| finger2 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 25.1 healing_power points (1.03 DPS) | yes | The Queen's Jewel (13094, +0.00 DPS, sim-verified) [world_drop]; Black Pearl Ring (6332, -0.26 DPS) [world]; Sea Giant's Toe Ring (274746, -0.29 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, +0.00 DPS, sim-verified) [dungeon]; Wind Spirit Staff (6689, -0.86 DPS) [dungeon]; Lorekeeper's Staff (212580, -1.27 DPS) [vendor] |
| off_hand | Defective Samophlange (274743) | Winklespark [vendor] | sim-verified (0.9 DPS) | yes | Orb of Souls (249395, -0.04 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -0.14 DPS) [world_drop]; Insignia Buckler (4066, -0.21 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Hotshot Pilot's Gloves; waist: Silver-lined Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Defective Samophlange

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 0000000000000000-000000000000000000-5032503315400000)

Set DPS (verified): 65.5. Weights run: 5.5s. Verify run: 4.2s. 592 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.522 ± 0.017, spirit=2.749 ± 0.011, mp5=7.094 ± 0.020, crit=0.303 ± 0.015 per rating point (14 rating = 1%, 4.235 per %), spell_haste=not significant (0.033 ± 0.023)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 89.6 healing_power points (3.67 DPS) | yes | Whitemane's Chapeau (7720, -0.05 DPS, sim-verified) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.77 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.79 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.02 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.38 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 56.8 healing_power points (2.33 DPS) | yes | Mistscape Mantle (4734, -0.08 DPS, sim-verified) [dungeon]; Batwing Mantle (6697, -0.63 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.69 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 38.7 healing_power points (1.58 DPS) | yes | Caretaker's Cape (19532, -0.12 DPS, sim-verified) [rep]; Mantle of Lady Falther'ess (23178, -0.29 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.29 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 80.0 healing_power points (3.27 DPS) | yes | Doomsayer's Robe (4746, +0.00 DPS, sim-verified) [quest]; Icemail Jerkin (1981, -0.57 DPS) [world_drop]; Deathchill Armor (10764, -0.66 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Enchanted Kodo Bracers (13119, -0.02 DPS, sim-verified) [world_drop]; Earthen Silk Cuffs (254019, -0.41 DPS) [crafted]; Silkstream Cuffs (16791, -0.42 DPS) [quest] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 51.1 healing_power points (2.09 DPS) | yes | Gloves of Old (9395, -0.06 DPS, sim-verified) [world_drop]; Mender's Leather Gloves (252530, -0.30 DPS) [crafted]; Stormcloth Gloves (10011, -0.40 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 58.7 healing_power points (2.40 DPS) | yes | Mender's Leather Belt (252523, +0.00 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.56 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Misplaced Pantaloons (276201) | Friz Frazzlespark [vendor] | 68.5 healing_power points (2.80 DPS) | yes | Warchief Kilt (7760, -0.08 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -0.43 DPS) [crafted]; Stormcloth Pants (10010, -0.50 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Furen's Boots (13100, -0.02 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -0.41 DPS) [dungeon] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.6 healing_power points (1.74 DPS) | yes | Welken Ring (5011, -0.44 DPS) [world_drop]; The Queen's Jewel (13094, -0.64 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 36.9 healing_power points (1.51 DPS) | yes | Welken Ring (5011, +0.00 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.40 DPS) [world_drop]; Voodoo Band (1996, -0.45 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Staff of Jordan (873, -0.41 DPS) [world_drop]; Wind Spirit Staff (6689, -0.58 DPS) [dungeon] |
| off_hand | Ravager's Shield (14777) | World drop [world_drop] | 34.1 healing_power points (1.40 DPS) | yes | Beacon of Hope (9393, -0.06 DPS, sim-verified) [dungeon]; Mordresh's Lifeless Skull (10770, -0.16 DPS) [dungeon]; Orb of Souls (249395, -0.19 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; hands: Bonefingers; waist: Gilded Cord; legs: Misplaced Pantaloons; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; off_hand: Ravager's Shield

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 0000000000000000-000000000000000000-5032503315513131)

Set DPS (verified): 104.0. Weights run: 9.6s. Verify run: 6.9s. 762 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.062, intellect=2.339 ± 0.023, spirit=3.097 ± 0.019, mp5=7.723 ± 0.026, crit=0.443 ± 0.024 per rating point (14 rating = 1%, 6.203 per %), spell_haste=not significant (0.206 ± 0.075)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | 97.8 healing_power points (4.86 DPS) | yes | Gemburst Circlet (10751, -0.09 DPS, sim-verified) [quest]; Papal Fez (9431, -0.41 DPS) [dungeon]; Soulcatcher Halo (10630, -0.42 DPS) [dungeon] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 60.4 healing_power points (3.00 DPS) | yes | Lei of Lilies (1315, -0.20 DPS, sim-verified) [world_drop]; Glowing Eye of Mordresh (10769, -0.72 DPS) [dungeon]; Horizon Choker (13085, -0.76 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Living Shoulders (15061, -0.01 DPS) [crafted]; Lead Surveyor's Mantle (11842, -0.10 DPS, sim-verified) [dungeon]; Mender's Leather Shoulder (252538, -0.24 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Featherskin Cape (10843, -0.16 DPS, sim-verified) [world]; Caretaker's Cape (19531, -0.38 DPS) [rep]; Imperial Red Cloak (8248, -0.50 DPS) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 132.7 healing_power points (6.58 DPS) | yes | Ghostweave Vest (14141, -0.36 DPS, sim-verified) [crafted]; Vestments of the Atal'ai Prophet (10806, -1.16 DPS) [dungeon]; Robes of Insight (940, -1.38 DPS) [world_drop] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 54.0 healing_power points (2.68 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Aristocratic Cuffs (12546, -0.01 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 105.5 healing_power points (5.23 DPS) | yes | Stonerender Gauntlets (17007, -0.07 DPS, sim-verified) [world_drop]; Mender's Leather Gauntlets (252551, -1.27 DPS) [crafted]; Mender's Mail Gauntlets (252587, -1.27 DPS) [crafted] |
| waist | Bloodlust Belt (14803) | World drop [world_drop] | 60.6 healing_power points (3.01 DPS) | yes | Gilded Cord (254037, -0.07 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.13 DPS) [crafted]; Earthenweave Cord (254077, -0.29 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Dalewind Trousers (13008, -0.17 DPS, sim-verified) [world_drop]; Windscale Sarong (10842, -0.62 DPS) [world]; Rainstrider Leggings (11123, -0.71 DPS) [quest] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 80.7 healing_power points (4.00 DPS) | yes | Mistwalker Boots (10629, -0.31 DPS) [dungeon]; Furen's Boots (13100, -0.73 DPS) [world_drop]; Coldstone Slippers (18697, -0.84 DPS) [dungeon] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 54.1 healing_power points (2.68 DPS) | yes | Eye of Adaegus (5266, -0.16 DPS, sim-verified) [world_drop]; Choking Band (11868, -0.68 DPS) [quest]; Snake Hoop (6750, -0.79 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Adaegus (5266, -0.17 DPS, sim-verified) [world_drop]; Choking Band (11868, -0.51 DPS) [quest]; Snake Hoop (6750, -0.62 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -4.21 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.52 DPS) [quest]; Uther's Strength (11302, -4.86 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.31 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.61 DPS) [quest]; Uther's Strength (11302, -0.95 DPS) [world_drop] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -0.44 DPS, sim-verified) [dungeon]; Glowing Brightwood Staff (812, -0.52 DPS) [world_drop]; Resurgence Rod (17743, -1.06 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Helm of Exile; neck: Darkmoon Necklace; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Bloodlust Belt; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Soulkeeper

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 211.7. Weights run: 10.8s. Verify run: 18.3s. 1751 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.084, intellect=4.714 ± 0.062, spirit=4.716 ± 0.056, mp5=11.904 ± 0.064, crit=0.868 ± 0.061 per rating point (14 rating = 1%, 12.149 per %), spell_haste=not significant (0.114 ± 0.283)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Gnomish Turban of Psychic Might (21517, -0.61 DPS, sim-verified) [quest]; Crown of the Penitent (13216, -0.81 DPS) [quest]; Devout Crown (16693, -0.90 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 147.5 healing_power points (6.67 DPS) | yes | Lady Maye's Pendant (14558, -0.08 DPS, sim-verified) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.70 DPS) [world_drop]; Tooth of Gnarr (13141, -2.07 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulstealer Mantle (13374, -0.28 DPS) [dungeon]; Sunderseer Mantle (13185, -0.56 DPS) [dungeon]; Argent Elite Shoulders (227888, -1.09 DPS, sim-verified) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 113.2 healing_power points (5.12 DPS) | yes | Faded Hakkari Cloak (20218, -0.08 DPS, sim-verified) [quest]; Gracious Cape (18743, -0.40 DPS) [dungeon]; Shroud of the Exile (15421, -0.43 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, -0.53 DPS) [crafted]; Alanna's Embrace (13314, -0.59 DPS) [dungeon]; Tunic of Undead Slaying (23089, -1.70 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, -0.21 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.35 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -1.48 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Hands of the Exalted Herald (12554, -0.02 DPS) [dungeon]; Devout Gloves (16692, -0.41 DPS) [dungeon]; Grasp of The Five Thunders (227014, -1.48 DPS, sim-verified) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 163.5 healing_power points (7.39 DPS) | yes | Devout Belt (16696, -0.09 DPS, sim-verified) [dungeon]; Whipvine Cord (18327, -0.84 DPS) [dungeon]; Feralsurge Girdle (18104, -0.95 DPS) [dungeon] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 253.8 healing_power points (11.48 DPS) | yes | Ghostloom Leggings (14545, -0.65 DPS, sim-verified) [dungeon]; Legplates of the Chromatic Defier (12945, -1.88 DPS) [quest]; Padre's Trousers (18386, -2.08 DPS) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 199.4 healing_power points (9.02 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.19 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -2.18 DPS) [crafted]; Faith Healer's Boots (22247, -2.73 DPS) [dungeon] |
| finger1 | Ring of Demonic Guile (18314) | Dire Maul: Alzzin the Wildshaper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.24 DPS) [dungeon]; Emerald Flame Ring (18395, -0.42 DPS) [dungeon]; Naglering (11669, -1.00 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.16 DPS) [dungeon]; Emerald Flame Ring (18395, -0.34 DPS) [dungeon]; Naglering (11669, -1.51 DPS, sim-verified) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Mindtap Talisman (18371, +0.00 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -0.22 DPS, sim-verified) [dungeon]; Shard of the Splithooves (10659, -4.85 DPS) [quest]; Ankh of Life (1713, -5.52 DPS) [world_drop] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quel'dorai Channeling Rod (18311, -0.17 DPS) [dungeon]; Hand of Edward the Odd (2243, -0.96 DPS, sim-verified) [world_drop]; Dancing Sliver (15854, -1.28 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Frostweaver Cape; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Leggings of Arcana; feet: Greaves of The Five Thunders; finger1: Ring of Demonic Guile; finger2: Band of Piety; trinket1: Serenity Field; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 570.5. Weights run: 8.0s. Verify run: 12.6s. 1751 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.250, intellect=1.881 ± 0.048, spirit=1.359 ± 0.041, mp5=3.953 ± 0.051, crit=0.786 ± 0.042 per rating point (14 rating = 1%, 10.999 per %), spell_haste=not significant (1.903 ± 0.637)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 125.9 healing_power points (19.70 DPS) | yes | Gnomish Turban of Psychic Might (21517, -5.01 DPS) [quest]; Sanctified Leather Helm (22689, -5.30 DPS) [quest]; Crown of The Five Thunders (227013, -45.92 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 66.9 healing_power points (10.47 DPS) | yes | Drake Tooth Necklace (21531, -2.87 DPS) [quest]; Jeweled Amulet of Cainwyn (1443, -3.04 DPS) [world_drop]; Lady Maye's Pendant (14558, -3.09 DPS, sim-verified) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 96.3 healing_power points (15.07 DPS) | yes | Mantle of The Five Thunders (227011, -5.25 DPS) [vendor]; Devout Mantle (16695, -5.25 DPS) [dungeon]; Royal Cap Spaulders (14548, -5.59 DPS) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 60.8 healing_power points (9.52 DPS) | yes | Cloak of the Cosmos (18389, -2.21 DPS) [dungeon]; Frostweaver Cape (12968, -3.43 DPS) [dungeon]; Drape of Recovery (272413, -3.59 DPS, sim-verified) [vendor] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | sim-verified (570.2 DPS) | yes | Robes of the Exalted (13346, -0.96 DPS) [dungeon]; Mooncloth Vest (14138, -1.41 DPS) [crafted]; Tunic of Undead Slaying (23089, -38.83 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (570.2 DPS) | yes | Bracers of Mending (23129, -0.38 DPS) [dungeon]; Bracers of The Five Thunders (227009, -0.80 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -70.09 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 74.8 healing_power points (11.71 DPS) | yes | Hands of the Exalted Herald (12554, -0.16 DPS) [dungeon]; Grasp of The Five Thunders (227014, -0.88 DPS) [vendor]; Harmonious Gauntlets (18527, -1.19 DPS) [dungeon] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 71.7 healing_power points (11.21 DPS) | yes | Wisdom of the Timbermaw (19047, -0.33 DPS) [crafted]; Sash of Mercy (14553, -0.79 DPS) [world_drop]; Sash of The Five Thunders (227010, -1.16 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 103.3 healing_power points (16.17 DPS) | yes | Leggings of Arcana (12756, -1.66 DPS, sim-verified) [quest]; Leggings of The Five Thunders (227012, -2.78 DPS) [vendor]; Devout Skirt (16694, -3.11 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 89.4 healing_power points (13.99 DPS) | yes | Greaves of The Five Thunders (227015, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.34 DPS) [crafted]; Faith Healer's Boots (22247, -3.83 DPS) [dungeon] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (570.2 DPS) | yes | Rosewine Circle (13178, -1.23 DPS) [dungeon]; Emerald Flame Ring (18395, -1.28 DPS) [dungeon]; Naglering (11669, -53.61 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (570.2 DPS) | yes | Rosewine Circle (13178, -0.56 DPS) [dungeon]; Emerald Flame Ring (18395, -0.61 DPS) [dungeon]; Naglering (11669, -61.60 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (570.2 DPS) | yes | Serenity Field (272439, -2.71 DPS) [vendor]; Briarwood Reed (12930, -4.54 DPS, sim-verified) [dungeon]; Shard of the Splithooves (10659, -5.57 DPS) [quest] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-verified (570.2 DPS) | yes | Serenity Field (272439, -0.23 DPS) [vendor]; Briarwood Reed (12930, -2.27 DPS) [dungeon]; Shard of the Splithooves (10659, -3.09 DPS) [quest] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-verified (570.2 DPS) | yes | Redemption (22406, -0.13 DPS) [dungeon]; Staff of Metanoia (22394, -0.35 DPS) [dungeon]; Hand of Edward the Odd (2243, -43.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Whipvine Cord; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Band of Piety; finger2: Band of Mending; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (tauren, 0000000000000000-000000000000000000-5032100000000000)

Set DPS (verified): 27.1. Weights run: 4.6s. Verify run: 3.0s. 205 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.856 ± 0.007, spirit=1.848 ± 0.003, mp5=4.989 ± 0.007, crit=0.124 ± 0.006 per rating point (14 rating = 1%, 1.731 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.4 healing_power points (0.87 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.14 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.16 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 7.4 healing_power points (0.32 DPS) | yes | Roadwatcher's Confidence (281265, -0.08 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.7 healing_power points (0.71 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 12.7 healing_power points (0.54 DPS) | yes | Regent's Cloak (5969, +0.00 DPS, sim-verified) [world]; Traveler's Shawl (277289, -0.15 DPS) [quest]; Sanguine Cape (14376, -0.22 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.8 healing_power points (1.27 DPS) | yes | Filigreed Pristine Gown (253901, +0.00 DPS, sim-verified) [crafted]; Robe of the Moccasin (6465, -0.10 DPS) [dungeon]; Corsair's Overshirt (5202, -0.17 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 24.2 healing_power points (1.03 DPS) | yes | Tabitha's Cuffs (251486, +0.00 DPS, sim-verified) [quest]; Drakewing Bands (12999, -0.56 DPS) [world_drop]; Owl Bracers (4796, -0.64 DPS) [vendor] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 22.2 healing_power points (0.95 DPS) | yes | Wisdom's Leather Gloves (252499, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.24 DPS) [world_drop]; Pristine Gloves (253913, -0.24 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.4 healing_power points (0.79 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.15 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.5 healing_power points (1.64 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.62 DPS) [world_drop]; Ghastly Trousers (15449, -0.82 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.3 healing_power points (0.86 DPS) | yes | Black Whelp Slippers (252424, +0.00 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.24 DPS) [crafted]; Bluegill Sandals (1560, -0.39 DPS) [world] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 14.8 healing_power points (0.63 DPS) | yes | Band of Purification (12996, -0.16 DPS) [world_drop]; Advisor's Ring (20426, -0.21 DPS) [rep]; Loop of Sacrifice (281673, -0.24 DPS) [quest] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.1 healing_power points (0.47 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Advisor's Ring (20426, -0.05 DPS) [rep]; Loop of Sacrifice (281673, -0.08 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) | The Wrath of Rath'mael [quest] | sim-verified (0.4 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Advisor's Gnarled Staff (20425, +0.00 DPS) [pvp]; Staff of Orgrimmar (15444, -0.03 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (tauren, 0000000000000000-000000000000000000-5032503300000000)

Set DPS (verified): 45.2. Weights run: 4.7s. Verify run: 3.1s. 349 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.995 ± 0.006, spirit=2.468 ± 0.013, mp5=6.283 ± 0.014, crit=0.185 ± 0.009 per rating point (14 rating = 1%, 2.592 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.2 healing_power points (2.03 DPS) | yes | Holy Shroud (2721, -0.06 DPS) [world_drop]; Whisperwind Headdress (6688, -0.06 DPS, sim-verified) [dungeon]; Nightsky Cowl (4039, -0.53 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.9 healing_power points (0.73 DPS) | yes | Necklace of Harmony (5180, +0.00 DPS, sim-verified) [world]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop]; Pendant of Myzrael (4614, -0.13 DPS) [dungeon] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 34.3 healing_power points (1.41 DPS) | yes | Ghostly Mantle (3324, -0.10 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.25 DPS) [world_drop]; Talbar Mantle (10657, -0.40 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-verified (0.9 DPS) | yes | Battle Healer's Cloak (19529, -0.02 DPS) [rep]; Glowing Thresher Cape (6901, -0.07 DPS, sim-verified) [dungeon]; Amy's Blanket (13005, -0.25 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.8 healing_power points (1.84 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.05 DPS) [dungeon]; Death Speaker Robes (6682, -0.12 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 28.8 healing_power points (1.19 DPS) | yes | Nightsky Wristbands (6407, +0.00 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.53 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 39.6 healing_power points (1.63 DPS) | yes | Hotshot Pilot's Gloves (9491, +0.00 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.55 DPS) [quest]; Tattered Mittens (270030, -0.61 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.8 healing_power points (1.88 DPS) | yes | Silver-lined Belt (13011, +0.00 DPS, sim-verified) [world_drop]; Lilac Sash (6780, -0.84 DPS) [quest]; Highlander's Mail Girdle (20120, -0.90 DPS) [vendor] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.2 healing_power points (2.15 DPS) | yes | Pristine Leggings (253987, -0.16 DPS) [crafted]; Earthen Leggings (253999, -0.42 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.8 healing_power points (1.80 DPS) | yes | Soggy Boots (274747, -0.07 DPS, sim-verified) [vendor]; Acidic Walkers (9454, -0.74 DPS) [dungeon]; Frothing Slippers (254003, -0.82 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 31.2 healing_power points (1.29 DPS) | yes | The Queen's Jewel (13094, -0.31 DPS) [world_drop]; Black Pearl Ring (6332, -0.51 DPS) [world]; Sea Giant's Toe Ring (274746, -0.55 DPS) [vendor] |
| finger2 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 25.1 healing_power points (1.03 DPS) | yes | The Queen's Jewel (13094, -0.06 DPS) [world_drop]; Black Pearl Ring (6332, -0.26 DPS) [world]; Sea Giant's Toe Ring (274746, -0.29 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Manual Crowd Pummeler (9449, +0.00 DPS, sim-verified) [dungeon]; Wind Spirit Staff (6689, -0.86 DPS) [dungeon]; Advisor's Gnarled Staff (19569, -1.19 DPS) [pvp] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 27.8 healing_power points (1.14 DPS) | yes | Defective Samophlange (274743, -0.23 DPS) [vendor]; Orb of Mistmantle (13031, -0.37 DPS) [world_drop]; Insignia Buckler (4066, -0.43 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-000000000000000000-5032503315400000)

Set DPS (verified): 65.1. Weights run: 5.5s. Verify run: 4.0s. 555 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.522 ± 0.017, spirit=2.749 ± 0.011, mp5=7.094 ± 0.020, crit=0.303 ± 0.015 per rating point (14 rating = 1%, 4.235 per %), spell_haste=not significant (0.033 ± 0.023)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 89.6 healing_power points (3.67 DPS) | yes | Whitemane's Chapeau (7720, -0.05 DPS, sim-verified) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.77 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.79 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.02 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.38 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 56.8 healing_power points (2.33 DPS) | yes | Mistscape Mantle (4734, -0.06 DPS, sim-verified) [dungeon]; Batwing Mantle (6697, -0.63 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.69 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 38.7 healing_power points (1.58 DPS) | yes | Cloak of Blight (6832, -0.10 DPS, sim-verified) [quest]; Battle Healer's Cloak (19528, -0.29 DPS) [rep]; Mantle of Lady Falther'ess (23178, -0.29 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 80.0 healing_power points (3.27 DPS) | yes | Doomsayer's Robe (4746, +0.00 DPS, sim-verified) [quest]; Icemail Jerkin (1981, -0.57 DPS) [world_drop]; Deathchill Armor (10764, -0.66 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Enchanted Kodo Bracers (13119, -0.03 DPS, sim-verified) [world_drop]; Windtalker's Wristguards (19584, -0.21 DPS) [pvp]; Dryad's Wrist Bindings (19597, -0.21 DPS) [pvp] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 51.1 healing_power points (2.09 DPS) | yes | Gloves of Old (9395, +0.00 DPS, sim-verified) [world_drop]; Mender's Leather Gloves (252530, -0.30 DPS) [crafted]; Stormcloth Gloves (10011, -0.40 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 58.7 healing_power points (2.40 DPS) | yes | Mender's Leather Belt (252523, +0.00 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.56 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Misplaced Pantaloons (276201) | Friz Frazzlespark [vendor] | 68.5 healing_power points (2.80 DPS) | yes | Warchief Kilt (7760, +0.00 DPS, sim-verified) [dungeon]; Wisdom's Leather Leggings (252519, -0.43 DPS) [crafted]; Stormcloth Pants (10010, -0.50 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Furen's Boots (13100, -0.02 DPS, sim-verified) [world_drop]; Thoughtcast Boots (10578, -0.41 DPS) [dungeon] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.6 healing_power points (1.74 DPS) | yes | Welken Ring (5011, -0.44 DPS) [world_drop]; The Queen's Jewel (13094, -0.64 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 36.9 healing_power points (1.51 DPS) | yes | Welken Ring (5011, +0.00 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.40 DPS) [world_drop]; Voodoo Band (1996, -0.45 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Advisor's Gnarled Staff (19568, -0.11 DPS) [pvp]; Staff of Jordan (873, -0.41 DPS) [world_drop] |
| off_hand | Ravager's Shield (14777) | World drop [world_drop] | 34.1 healing_power points (1.40 DPS) | yes | Beacon of Hope (9393, -0.06 DPS, sim-verified) [dungeon]; Prophetic Cane (6803, -0.16 DPS) [quest]; Mordresh's Lifeless Skull (10770, -0.16 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; hands: Bonefingers; waist: Gilded Cord; legs: Misplaced Pantaloons; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket2: Ankh of Life; off_hand: Ravager's Shield

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (tauren, 0000000000000000-000000000000000000-5032503315513131)

Set DPS (verified): 103.7. Weights run: 9.6s. Verify run: 6.7s. 704 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.062, intellect=2.339 ± 0.023, spirit=3.097 ± 0.019, mp5=7.723 ± 0.026, crit=0.443 ± 0.024 per rating point (14 rating = 1%, 6.203 per %), spell_haste=not significant (0.206 ± 0.075)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | 97.8 healing_power points (4.86 DPS) | yes | Gemburst Circlet (10751, +0.00 DPS, sim-verified) [quest]; Papal Fez (9431, -0.41 DPS) [dungeon]; Soulcatcher Halo (10630, -0.42 DPS) [dungeon] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 60.4 healing_power points (3.00 DPS) | yes | Lei of Lilies (1315, -0.19 DPS, sim-verified) [world_drop]; Glowing Eye of Mordresh (10769, -0.72 DPS) [dungeon]; Horizon Choker (13085, -0.76 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Living Shoulders (15061, -0.01 DPS) [crafted]; Lead Surveyor's Mantle (11842, -0.11 DPS, sim-verified) [dungeon]; Mender's Leather Shoulder (252538, -0.24 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Featherskin Cape (10843, -0.21 DPS, sim-verified) [world]; Battle Healer's Cloak (19527, -0.38 DPS) [rep]; Cloak of Blight (6832, -0.40 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 132.7 healing_power points (6.58 DPS) | yes | Ghostweave Vest (14141, -0.29 DPS, sim-verified) [crafted]; Vestments of the Atal'ai Prophet (10806, -1.16 DPS) [dungeon]; Robes of Insight (940, -1.38 DPS) [world_drop] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 54.0 healing_power points (2.68 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Aristocratic Cuffs (12546, -0.01 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 105.5 healing_power points (5.23 DPS) | yes | Stonerender Gauntlets (17007, -1.00 DPS) [world_drop]; Mender's Leather Gauntlets (252551, -1.27 DPS) [crafted]; Mender's Mail Gauntlets (252587, -1.27 DPS) [crafted] |
| waist | Bloodlust Belt (14803) | World drop [world_drop] | 60.6 healing_power points (3.01 DPS) | yes | Gilded Cord (254037, -0.07 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.13 DPS) [crafted]; Earthenweave Cord (254077, -0.27 DPS, sim-verified) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Dalewind Trousers (13008, -0.17 DPS, sim-verified) [world_drop]; Windscale Sarong (10842, -0.62 DPS) [world]; Rainstrider Leggings (11123, -0.71 DPS) [quest] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 80.7 healing_power points (4.00 DPS) | yes | Mistwalker Boots (10629, +0.00 DPS, sim-verified) [dungeon]; Furen's Boots (13100, -0.73 DPS) [world_drop]; Coldstone Slippers (18697, -0.84 DPS) [dungeon] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 54.1 healing_power points (2.68 DPS) | yes | Eye of Adaegus (5266, -0.11 DPS, sim-verified) [world_drop]; Snake Hoop (6750, -0.79 DPS) [quest]; Cyclopean Band (11824, -0.81 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Adaegus (5266, -0.17 DPS, sim-verified) [world_drop]; Snake Hoop (6750, -0.62 DPS) [quest]; Cyclopean Band (11824, -0.64 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -4.21 DPS) [quest]; Uther's Strength (11302, -4.86 DPS) [world_drop]; Alchemists' Stone (13503, -5.13 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS, sim-verified) [quest]; Uther's Strength (11302, -0.95 DPS) [world_drop]; Alchemists' Stone (13503, -1.23 DPS) [crafted] |
| main_hand | Soulkeeper (1607) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -0.42 DPS, sim-verified) [dungeon]; Glowing Brightwood Staff (812, -0.52 DPS) [world_drop]; Resurgence Rod (17743, -1.06 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Helm of Exile; neck: Darkmoon Necklace; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Bloodlust Belt; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Soulkeeper

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (tauren, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 211.2. Weights run: 10.8s. Verify run: 17.9s. 1672 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.084, intellect=4.714 ± 0.062, spirit=4.716 ± 0.056, mp5=11.904 ± 0.064, crit=0.868 ± 0.061 per rating point (14 rating = 1%, 12.149 per %), spell_haste=not significant (0.114 ± 0.283)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Gnomish Turban of Psychic Might (21517, -0.56 DPS, sim-verified) [quest]; Crown of the Penitent (13216, -0.81 DPS) [quest]; Devout Crown (16693, -0.90 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 147.5 healing_power points (6.67 DPS) | yes | Lady Maye's Pendant (14558, -0.18 DPS, sim-verified) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.70 DPS) [world_drop]; Tooth of Gnarr (13141, -2.07 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Soulstealer Mantle (13374, -0.28 DPS) [dungeon]; Sunderseer Mantle (13185, -0.56 DPS) [dungeon]; Argent Elite Shoulders (227888, -1.13 DPS, sim-verified) [vendor] |
| back | Frostweaver Cape (12968) | Blackrock Spire: The Beast [dungeon] | 113.2 healing_power points (5.12 DPS) | yes | Faded Hakkari Cloak (20218, -0.14 DPS, sim-verified) [quest]; Gracious Cape (18743, -0.40 DPS) [dungeon]; Shroud of the Exile (15421, -0.43 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, -0.53 DPS) [crafted]; Alanna's Embrace (13314, -0.59 DPS) [dungeon]; Tunic of Undead Slaying (23089, -1.73 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Mending (23129, -0.21 DPS) [dungeon]; Bleak Howler Armguards (13208, -0.35 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -1.51 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Hands of the Exalted Herald (12554, -0.02 DPS) [dungeon]; Devout Gloves (16692, -0.41 DPS) [dungeon]; Grasp of The Five Thunders (227014, -1.37 DPS, sim-verified) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 163.5 healing_power points (7.39 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Devout Belt (16696, -0.10 DPS, sim-verified) [dungeon]; Whipvine Cord (18327, -0.84 DPS) [dungeon] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 253.8 healing_power points (11.48 DPS) | yes | Ghostloom Leggings (14545, -0.60 DPS, sim-verified) [dungeon]; Legplates of the Chromatic Defier (12945, -1.88 DPS) [quest]; Padre's Trousers (18386, -2.08 DPS) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 199.4 healing_power points (9.02 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.27 DPS, sim-verified) [vendor]; Boots of The Five Thunders (22096, -2.17 DPS) [quest]; Mooncloth Boots (15802, -2.18 DPS) [crafted] |
| finger1 | Ring of Demonic Guile (18314) | Dire Maul: Alzzin the Wildshaper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.24 DPS) [dungeon]; Emerald Flame Ring (18395, -0.42 DPS) [dungeon]; Naglering (11669, -1.17 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Rivendare (13345, -0.16 DPS) [dungeon]; Emerald Flame Ring (18395, -0.34 DPS) [dungeon]; Naglering (11669, -1.55 DPS, sim-verified) [dungeon] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Mindtap Talisman (18371, +0.00 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -0.34 DPS, sim-verified) [dungeon]; Shard of the Splithooves (10659, -4.85 DPS) [quest]; Ankh of Life (1713, -5.52 DPS) [world_drop] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quel'dorai Channeling Rod (18311, -0.17 DPS) [dungeon]; Hand of Edward the Odd (2243, -1.06 DPS, sim-verified) [world_drop]; Dancing Sliver (15854, -1.28 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Devout Mantle; back: Frostweaver Cape; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Belt of Tiny Heads; legs: Leggings of Arcana; feet: Greaves of The Five Thunders; finger1: Ring of Demonic Guile; finger2: Band of Piety; trinket1: Serenity Field; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (tauren, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 569.9. Weights run: 8.0s. Verify run: 12.4s. 1672 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.250, intellect=1.881 ± 0.048, spirit=1.359 ± 0.041, mp5=3.953 ± 0.051, crit=0.786 ± 0.042 per rating point (14 rating = 1%, 10.999 per %), spell_haste=not significant (1.903 ± 0.637)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 125.9 healing_power points (19.70 DPS) | yes | Gnomish Turban of Psychic Might (21517, -5.01 DPS) [quest]; Sanctified Leather Helm (22689, -5.30 DPS) [quest]; Crown of The Five Thunders (227013, -61.60 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 66.9 healing_power points (10.47 DPS) | yes | Drake Tooth Necklace (21531, -2.87 DPS) [quest]; Jeweled Amulet of Cainwyn (1443, -3.04 DPS) [world_drop]; Lady Maye's Pendant (14558, -3.83 DPS, sim-verified) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 96.3 healing_power points (15.07 DPS) | yes | Champion's Mail Epaulets (227166, -4.63 DPS) [vendor]; Warlord's Mail Epaulets (231665, -4.88 DPS) [vendor]; Mantle of The Five Thunders (227011, -5.25 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 60.8 healing_power points (9.52 DPS) | yes | Cloak of the Cosmos (18389, -2.21 DPS) [dungeon]; Drape of Recovery (272413, -3.13 DPS, sim-verified) [vendor]; Frostweaver Cape (12968, -3.43 DPS) [dungeon] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | sim-verified (569.9 DPS) | yes | Robes of the Exalted (13346, -0.96 DPS) [dungeon]; Mooncloth Vest (14138, -1.41 DPS) [crafted]; Tunic of Undead Slaying (23089, -48.21 DPS, sim-verified) [world] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (569.9 DPS) | yes | Bracers of Mending (23129, -0.38 DPS) [dungeon]; Bracers of The Five Thunders (227009, -0.80 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -67.14 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 74.8 healing_power points (11.71 DPS) | yes | Grasp of The Five Thunders (227014, -0.88 DPS) [vendor]; Harmonious Gauntlets (18527, -1.19 DPS) [dungeon]; Hands of the Exalted Herald (12554, -1.31 DPS, sim-verified) [dungeon] |
| waist | Whipvine Cord (18327) | Dire Maul: Alzzin the Wildshaper [dungeon] | 71.7 healing_power points (11.21 DPS) | yes | Wisdom of the Timbermaw (19047, -0.33 DPS) [crafted]; Sash of Mercy (14553, -0.79 DPS) [world_drop]; Sash of The Five Thunders (227010, -1.16 DPS) [vendor] |
| legs | Padre's Trousers (18386) | Dire Maul: Illyanna Ravenoak [dungeon] | 103.3 healing_power points (16.17 DPS) | yes | Leggings of Arcana (12756, -2.50 DPS, sim-verified) [quest]; Leggings of The Five Thunders (227012, -2.78 DPS) [vendor]; Devout Skirt (16694, -3.11 DPS) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 89.4 healing_power points (13.99 DPS) | yes | Greaves of The Five Thunders (227015, -1.54 DPS) [vendor]; Mooncloth Boots (15802, -3.34 DPS) [crafted]; Faith Healer's Boots (22247, -3.83 DPS) [dungeon] |
| finger1 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (569.9 DPS) | yes | Rosewine Circle (13178, -1.23 DPS) [dungeon]; Emerald Flame Ring (18395, -1.28 DPS) [dungeon]; Naglering (11669, -54.56 DPS, sim-verified) [dungeon] |
| finger2 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-verified (569.9 DPS) | yes | Rosewine Circle (13178, -0.56 DPS) [dungeon]; Emerald Flame Ring (18395, -0.61 DPS) [dungeon]; Naglering (11669, -59.55 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (569.9 DPS) | yes | Serenity Field (272439, -2.71 DPS) [vendor]; Shard of the Splithooves (10659, -5.57 DPS) [quest]; Briarwood Reed (12930, -7.72 DPS, sim-verified) [dungeon] |
| trinket2 | Mindtap Talisman (18371) | Dire Maul: Magister Kalendris [dungeon] | sim-verified (569.9 DPS) | yes | Serenity Field (272439, -0.23 DPS) [vendor]; Briarwood Reed (12930, -2.27 DPS) [dungeon]; Shard of the Splithooves (10659, -3.09 DPS) [quest] |
| main_hand | Hammer of the Grand Crusader (18717) | Stratholme: Balnazzar [dungeon] | sim-verified (569.9 DPS) | yes | Redemption (22406, -0.13 DPS) [dungeon]; Staff of Metanoia (22394, -0.35 DPS) [dungeon]; Hand of Edward the Odd (2243, -65.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Whipvine Cord; legs: Padre's Trousers; feet: Incandescent Mooncloth Boots; finger1: Band of Piety; finger2: Band of Mending; trinket2: Mindtap Talisman; main_hand: Hammer of the Grand Crusader

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

