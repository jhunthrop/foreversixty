# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5510000000000000)

Set DPS (verified): 25.0. Weights run: 4.7s. Verify run: 17.7s. 213 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.093 ± 0.007, spirit=1.372 ± 0.003, mp5=4.674 ± 0.016, crit=0.122 ± 0.006 per rating point (14 rating = 1%, 1.709 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (25.0 DPS) | yes | Wisdom's Leather Hood (252507, -0.04 DPS) [crafted]; Trapper's Leather Hood (252505, -0.05 DPS) [crafted]; Shadow Goggles (4373, -0.32 DPS, sim-verified) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.5 healing_power points (0.21 DPS) | yes | Scholarly Pendant (277203, -0.11 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 18.8 healing_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, -0.06 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.51 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 11.7 healing_power points (0.45 DPS) | yes | Seer's Cape (6378, -0.19 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.21 DPS) [crafted]; Sanguine Cape (14376, -0.34 DPS, sim-verified) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.6 healing_power points (1.14 DPS) | yes | Filigreed Pristine Gown (253901, -0.11 DPS, sim-verified) [crafted]; Robe of the Moccasin (6465, -0.26 DPS) [dungeon]; Corsair's Overshirt (5202, -0.29 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 24.5 healing_power points (0.94 DPS) | yes | Bright Bracers (3647, -0.62 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.62 DPS) [vendor]; Owl Bracers (4796, -0.68 DPS, sim-verified) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 18.4 healing_power points (0.71 DPS) | yes | Pristine Gloves (253913, -0.08 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.09 DPS) [world_drop]; Bright Gloves (3066, -0.17 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 19.4 healing_power points (0.75 DPS) | yes | Novice Ardent's Sash (253887, -0.11 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.19 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.0 healing_power points (1.46 DPS) | yes | Filigreed Pristine Leggings (253937, -0.11 DPS, sim-verified) [crafted]; Dreamer's Leggings (270016, -0.58 DPS) [quest]; Darkweave Breeches (12987, -0.58 DPS) [world_drop] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 21.5 healing_power points (0.83 DPS) | yes | Pristine Boots (253889, -0.24 DPS) [crafted]; Black Whelp Slippers (252424, -0.25 DPS, sim-verified) [crafted]; Kimbra Boots (6191, -0.35 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 12.6 healing_power points (0.48 DPS) | yes | Band of Purification (12996, -0.17 DPS) [world_drop]; Deep Fathom Ring (6463, -0.22 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.3 healing_power points (0.36 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Deep Fathom Ring (6463, -0.10 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 20.9 healing_power points (0.81 DPS) | yes | Staff of Westfall (2042, +0.00 DPS, sim-verified) [quest]; Staff of the Blessed Seer (2271, -0.28 DPS) [dungeon]; Lesser Staff of the Spire (1300, -0.32 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 213, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (dwarf, 0000000000000000-000000000000000000-5532003300000000)

Set DPS (verified): 42.4. Weights run: 4.8s. Verify run: 17.5s. 388 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.993 ± 0.007, spirit=2.468 ± 0.013, mp5=6.283 ± 0.014, crit=0.174 ± 0.008 per rating point (14 rating = 1%, 2.431 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.2 healing_power points (2.03 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.53 DPS) [world_drop]; Dokebi Hat (14584, -0.62 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.8 healing_power points (0.74 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.13 DPS) [dungeon]; Glowing Green Talisman (5002, -0.13 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 31.2 healing_power points (1.29 DPS) | yes | Nightsky Mantle (4718, -0.12 DPS) [world_drop]; Faerie Mantle (5820, -0.18 DPS) [quest]; Talbar Mantle (10657, -0.28 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 26.7 healing_power points (1.11 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Darkspear Raider's Cloak (272078, -0.14 DPS) [vendor]; Caretaker's Cape (19533, -0.16 DPS) [rep] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.8 healing_power points (1.85 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.05 DPS) [dungeon]; Death Speaker Robes (6682, -0.12 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 28.8 healing_power points (1.19 DPS) | yes | Nightsky Wristbands (6407, -0.35 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.53 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 39.6 healing_power points (1.64 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.16 DPS, sim-verified) [dungeon]; Zodiac Gloves (7106, -0.59 DPS) [quest]; Silver-thread Gloves (6393, -0.72 DPS) [world_drop] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.8 healing_power points (1.89 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.90 DPS) [rep]; Highlander's Mail Girdle (20120, -0.90 DPS) [vendor]; Silver-lined Belt (13011, -1.27 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.2 healing_power points (2.16 DPS) | yes | Pristine Leggings (253987, -0.16 DPS) [crafted]; Earthen Leggings (253999, -0.42 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.8 healing_power points (1.81 DPS) | yes | Acidic Walkers (9454, -0.75 DPS) [dungeon]; Frothing Slippers (254003, -0.83 DPS) [crafted]; Soggy Boots (274747, -1.37 DPS, sim-verified) [vendor] |
| finger1 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 25.1 healing_power points (1.04 DPS) | yes | Electrocutioner Lagnut (9447, -0.08 DPS) [dungeon]; Runemender's Seal (273031, -0.08 DPS) [world]; Ring of Power Regulation (273030, -0.26 DPS) [dungeon] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 23.7 healing_power points (0.98 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS, sim-verified) [dungeon]; Runemender's Seal (273031, -0.02 DPS) [world]; Ring of Power Regulation (273030, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (42.4 DPS) | yes | Hardened Root Staff (1317, -1.07 DPS) [quest]; Manual Crowd Pummeler (9449, -1.13 DPS, sim-verified) [dungeon]; Lorekeeper's Staff (212580, -1.28 DPS) [vendor] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 27.8 healing_power points (1.15 DPS) | yes | Defective Samophlange (274743, -0.24 DPS, sim-verified) [vendor]; Orb of Mistmantle (13031, -0.37 DPS) [world_drop]; Insignia Buckler (4066, -0.44 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Darkspear Signet; finger2: The Queen's Jewel; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots

### Band 40 (dwarf, 0000000000000000-000000000000000000-5532003315400000)

Set DPS (verified): 73.8. Weights run: 5.7s. Verify run: 24.4s. 606 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.563 ± 0.017, spirit=2.761 ± 0.012, mp5=7.128 ± 0.021, crit=0.285 ± 0.015 per rating point (14 rating = 1%, 3.996 per %), spell_haste=not significant (0.016 ± 0.016)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 82.4 healing_power points (3.39 DPS) | yes | Whitemane's Chapeau (7720, -0.67 DPS, sim-verified) [dungeon]; Imperial Leather Helm (6433, -0.88 DPS) [dungeon]; Enduring Cap (3020, -0.91 DPS) [world_drop] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 37.3 healing_power points (1.53 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.40 DPS) [quest]; Explorers' League Commendation (7746, -0.51 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 57.6 healing_power points (2.37 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS, sim-verified) [dungeon]; Mistscape Mantle (4734, -0.64 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.72 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 39.2 healing_power points (1.61 DPS) | yes | Sergeant Major's Cape (16336, -0.23 DPS, sim-verified) [pvp]; Caretaker's Cape (19532, -0.31 DPS) [rep]; Blackforge Cape (6424, -0.32 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 80.1 healing_power points (3.29 DPS) | yes | Dreamweave Vest (10021, -0.76 DPS) [crafted]; Doomsayer's Robe (4746, -0.84 DPS) [quest]; Deathchill Armor (10764, -1.48 DPS, sim-verified) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 37.9 healing_power points (1.56 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Silkstream Cuffs (16791, -0.58 DPS) [quest]; Earthen Silk Cuffs (254019, -0.58 DPS) [crafted] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 51.6 healing_power points (2.12 DPS) | yes | Mender's Leather Gloves (252530, -0.31 DPS) [crafted]; Stormcloth Gloves (10011, -0.40 DPS) [crafted]; Gloves of Old (9395, -0.56 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 59.1 healing_power points (2.43 DPS) | yes | Mender's Leather Belt (252523, -0.40 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.56 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 70.0 healing_power points (2.88 DPS) | yes | Stoneweaver Leggings (9407, -0.14 DPS) [dungeon]; Misplaced Pantaloons (276201, -0.36 DPS, sim-verified) [vendor]; Wisdom's Leather Leggings (252519, -0.48 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -0.50 DPS) [crafted]; Furen's Boots (13100, -0.92 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.8 healing_power points (1.76 DPS) | yes | The Queen's Jewel (13094, -0.64 DPS) [world_drop]; Ogremind Ring (1993, -0.68 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 32.1 healing_power points (1.32 DPS) | yes | Ogremind Ring (1993, -0.24 DPS) [world_drop]; Voodoo Band (1996, -0.24 DPS) [world]; The Queen's Jewel (13094, -0.25 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -1.54 DPS) [dungeon]; Hypnotic Blade (7714, -1.78 DPS) [dungeon] |
| off_hand | Battered Viking Shield (9403) | Uldaman: Olaf [dungeon] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ravager's Shield (14777, -0.05 DPS) [world_drop]; Orb of Souls (249395, -0.25 DPS) [crafted]; Orb of Lorica (11262, -7.50 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Battered Viking Shield

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 50 (dwarf, 0000000000000000-000000000000000000-5532003315513131)

Set DPS (verified): 124.7. Weights run: 9.7s. Verify run: 41.7s. 774 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.070, intellect=2.374 ± 0.023, spirit=2.996 ± 0.018, mp5=7.513 ± 0.025, crit=0.449 ± 0.025 per rating point (14 rating = 1%, 6.293 per %), spell_haste=not significant (0.132 ± 0.076)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 103.3 healing_power points (5.29 DPS) | yes | Winged Helm (13112, -0.96 DPS) [world_drop]; Serpentskin Helm (8261, -1.04 DPS) [world_drop]; Engineer's Guild Headpiece (9534, -1.40 DPS, sim-verified) [quest] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 59.3 healing_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.56 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.11 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.40 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 71.4 healing_power points (3.66 DPS) | yes | Living Shoulders (15061, +0.00 DPS, sim-verified) [crafted]; Mender's Leather Shoulders (252538, -0.26 DPS) [crafted]; Mender's Mail Shoulders (252569, -0.26 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 48.2 healing_power points (2.47 DPS) | yes | Caretaker's Cape (19531, -0.42 DPS) [rep]; Imperial Red Cloak (8248, -0.52 DPS) [world_drop]; Wingveil Cloak (10802, -0.53 DPS) [dungeon] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 120.0 healing_power points (6.15 DPS) | yes | Feathered Breastplate (8349, -1.08 DPS, sim-verified) [crafted]; Stormcloth Vest (10020, -1.95 DPS) [crafted]; Deathchill Armor (10764, -2.09 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 53.6 healing_power points (2.75 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.72 DPS) [world_drop] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 104.0 healing_power points (5.33 DPS) | yes | Mender's Leather Gauntlets (252551, -0.72 DPS, sim-verified) [crafted]; Mender's Mail Gauntlets (252587, -1.27 DPS) [crafted]; Gilded Gloves (254095, -1.40 DPS) [crafted] |
| waist | Bloodlust Belt (14803) | World drop [world_drop] | 60.2 healing_power points (3.09 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -0.06 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.09 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 102.6 healing_power points (5.26 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.17 DPS, sim-verified) [dungeon]; Dalewind Trousers (13008, -0.33 DPS) [world_drop]; Senior Designer's Pantaloons (11841, -0.36 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 78.9 healing_power points (4.04 DPS) | yes | Furen's Boots (13100, -0.62 DPS, sim-verified) [world_drop]; Vinerot Sandals (17748, -0.80 DPS) [dungeon]; Mender's Leather Shoes (252533, -0.89 DPS) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 52.6 healing_power points (2.70 DPS) | yes | Choking Band (11868, -0.70 DPS) [quest]; Coldwater Ring (4550, -0.95 DPS) [quest]; Welken Ring (5011, -1.01 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 50.2 healing_power points (2.57 DPS) | yes | Coldwater Ring (4550, -0.83 DPS) [quest]; Choking Band (11868, -0.86 DPS, sim-verified) [quest]; Welken Ring (5011, -0.89 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (124.7 DPS) | yes | Ankh of Life (1713, -3.93 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.24 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.55 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (124.7 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.50 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (124.7 DPS) | yes | Charstone Dirk (17710, -4.54 DPS) [dungeon]; Hand of Righteousness (7721, -6.55 DPS) [dungeon]; Hammer of the Northern Wind (810, -20.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Bloodlust Belt; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 774, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60 (dwarf, 5300000000000000-000000000000000000-5532003315513151)

Set DPS (verified): 235.0. Weights run: 11.4s. Verify run: 131.3s. 1484 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=5.110 ± 0.061, spirit=4.488 ± 0.045, mp5=11.210 ± 0.052, crit=0.755 ± 0.051 per rating point (14 rating = 1%, 10.576 per %), spell_haste=not significant (0.197 ± 0.160)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 259.3 healing_power points (11.61 DPS) | yes | Living Crown (252561, +0.00 DPS, sim-verified) [crafted]; Crown of the Penitent (13216, -1.74 DPS) [quest]; Devout Crown (16693, -2.21 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 144.3 healing_power points (6.46 DPS) | yes | Stormpike Sage's Pendant (19098, -2.17 DPS) [rep]; Horizon Choker (13085, -2.46 DPS) [world_drop]; Darkmoon Necklace (19303, -4.47 DPS, sim-verified) [vendor] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 171.1 healing_power points (7.66 DPS) | yes | Mantle of The Five Thunders (227011, -1.29 DPS) [vendor]; Drakesfire Epaulets (13133, -1.45 DPS) [world_drop]; Devout Mantle (16695, -3.39 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | 108.7 healing_power points (4.87 DPS) | yes | Faded Hakkari Cloak (20218, -0.02 DPS) [quest]; Hide of the Wild (18510, -0.70 DPS) [crafted]; Sergeant Major's Cape (16337, -1.43 DPS) [pvp] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 213.5 healing_power points (9.56 DPS) | yes | Devout Robe (16690, -0.38 DPS, sim-verified) [dungeon]; Tunic of The Five Thunders (227016, -0.75 DPS) [vendor]; Dreamwalker Armor (13123, -0.96 DPS) [world_drop] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 128.1 healing_power points (5.73 DPS) | yes | Bracers of The Five Thunders (227009, -0.77 DPS, sim-verified) [vendor]; Devout Bracers (16697, -1.40 DPS) [dungeon]; Magister's Bindings (16683, -1.55 DPS) [dungeon] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | sim-verified (234.9 DPS) | yes | Devout Gloves (16692, -0.83 DPS) [dungeon]; Raider Handwraps (272098, -1.19 DPS) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 165.3 healing_power points (7.40 DPS) | yes | Devout Belt (16696, -0.67 DPS, sim-verified) [dungeon]; Elunarian Belt (14465, -0.82 DPS) [world_drop]; Sash of The Five Thunders (227010, -0.97 DPS) [vendor] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | sim-verified (234.9 DPS) | yes | Red Dragonscale Leggings (252603, -0.45 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -0.66 DPS) [dungeon]; Devout Skirt (16694, -5.56 DPS, sim-verified) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 198.6 healing_power points (8.89 DPS) | yes | Mooncloth Boots (15802, -2.00 DPS) [crafted]; Devout Sandals (16691, -2.70 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -5.51 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 137.0 healing_power points (6.14 DPS) | yes | Channeler's Ring (272406, -1.32 DPS) [vendor]; Darkspear Signet (272068, -1.62 DPS) [vendor]; Band of the Hierophant (13096, -1.64 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 117.9 healing_power points (5.28 DPS) | yes | Darkspear Signet (272068, -0.76 DPS) [vendor]; Band of the Hierophant (13096, -0.78 DPS) [world_drop]; Channeler's Ring (272406, -2.42 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (234.9 DPS) | yes | Ankh of Life (1713, -5.12 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -5.52 DPS) [quest]; Serenity Field (272439, -5.65 DPS) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (234.9 DPS) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -0.60 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (234.9 DPS) | yes | Staff of Hale Magefire (13000, -5.02 DPS) [world_drop]; Charstone Dirk (17710, -6.44 DPS) [dungeon]; Hand of Edward the Odd (2243, -12.33 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Darkspear Raider's Cloak; chest: Mooncloth Vest; wrist: Bracers of Hope; waist: Belt of Tiny Heads; legs: Leggings of The Five Thunders; feet: Greaves of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1484, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 5300000000000000-000000000000000000-5532003315513151)

Set DPS (verified): 630.4. Weights run: 8.2s. Verify run: 53.7s. 1484 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.144, intellect=1.144 ± 0.028, spirit=0.728 ± 0.022, mp5=2.098 ± 0.027, crit=0.481 ± 0.026 per rating point (14 rating = 1%, 6.740 per %), spell_haste=not significant (0.960 ± 0.372)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 103.8 healing_power points (30.32 DPS) | yes | Crown of The Five Thunders (227013, +0.00 DPS, sim-verified) [vendor]; Sanctified Leather Helm (22689, -9.73 DPS) [quest]; Mooncloth Circlet (14140, -12.23 DPS) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 47.7 healing_power points (13.92 DPS) | yes | Drake Tooth Necklace (21531, -3.05 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -7.45 DPS) [dungeon]; Stormpike Sage's Pendant (19098, -8.12 DPS) [rep] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 78.9 healing_power points (23.04 DPS) | yes | Mantle of The Five Thunders (227011, -0.88 DPS, sim-verified) [vendor]; Mooncloth Shoulders (14139, -9.89 DPS) [crafted]; Eternity Pauldrons (279255, -10.45 DPS) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.4 healing_power points (15.61 DPS) | yes | Drape of Recovery (272413, -3.54 DPS, sim-verified) [vendor]; Caretaker's Cape (19530, -6.31 DPS) [rep] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | 73.0 healing_power points (21.31 DPS) | yes | Mooncloth Vest (14138, -2.99 DPS) [crafted]; Devout Robe (16690, -4.26 DPS) [dungeon]; Stormcloth Vest (10020, -6.34 DPS, sim-verified) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 39.3 healing_power points (11.49 DPS) | yes | Bracers of The Five Thunders (227009, +0.00 DPS, sim-verified) [vendor]; Nethergeld Cuffs (254061, -2.24 DPS) [crafted]; Mender's Leather Bracers (252543, -2.32 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 54.9 healing_power points (16.03 DPS) | yes | Grasp of The Five Thunders (227014, -1.52 DPS) [vendor]; Mooncloth Gloves (18409, -1.79 DPS, sim-verified) [crafted]; Mender's Leather Gauntlets (252551, -2.06 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 53.6 healing_power points (15.65 DPS) | yes | Sash of The Five Thunders (227010, +0.00 DPS, sim-verified) [vendor]; Mender's Leather Waistguard (252477, -2.88 DPS) [crafted]; Mender's Mail Belt (252591, -2.88 DPS) [crafted] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | 62.4 healing_power points (18.23 DPS) | yes | Devout Skirt (16694, -1.32 DPS) [dungeon]; Elderwild Pants (279254, -1.44 DPS) [crafted]; Red Dragonscale Leggings (252603, -1.69 DPS) [crafted] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 70.2 healing_power points (20.51 DPS) | yes | Greaves of The Five Thunders (227015, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -6.05 DPS) [crafted]; Swarmtender's Footpads (275609, -7.95 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 55.8 healing_power points (16.30 DPS) | yes | Blessed Band of Light (272407, -5.54 DPS) [vendor]; Channeler's Ring (272406, -9.91 DPS) [vendor]; Band of the Hierophant (13096, -10.61 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 41.8 healing_power points (12.22 DPS) | yes | Blessed Band of Light (272407, -2.80 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -5.84 DPS) [vendor]; Band of the Hierophant (13096, -6.54 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (630.2 DPS) | yes | Uther's Strength (11302, -3.94 DPS) [world_drop]; Serenity Field (272439, -4.01 DPS, sim-verified) [vendor]; Ankh of Life (1713, -6.64 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (630.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -1.12 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (630.2 DPS) | yes | Charstone Dirk (17710, -16.64 DPS) [dungeon]; Hand of Edward the Odd (2243, -17.00 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -26.26 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Leggings of The Five Thunders; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1484, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

## Horde

### Band 20 (tauren, 0000000000000000-000000000000000000-5510000000000000)

Set DPS (verified): 25.0. Weights run: 4.7s. Verify run: 16.8s. 193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.093 ± 0.007, spirit=1.372 ± 0.003, mp5=4.674 ± 0.016, crit=0.122 ± 0.006 per rating point (14 rating = 1%, 1.709 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (25.0 DPS) | yes | Wisdom's Leather Hood (252507, -0.04 DPS) [crafted]; Trapper's Leather Hood (252505, -0.05 DPS) [crafted]; Shadow Goggles (4373, -0.36 DPS, sim-verified) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.1 healing_power points (0.16 DPS) | yes | Scholarly Pendant (277203, -0.05 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 18.8 healing_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, -0.09 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.51 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 11.7 healing_power points (0.45 DPS) | yes | Seer's Cape (6378, -0.19 DPS) [dungeon]; Traveler's Shawl (277289, -0.19 DPS) [quest]; Sanguine Cape (14376, -0.40 DPS, sim-verified) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.6 healing_power points (1.14 DPS) | yes | Filigreed Pristine Gown (253901, -0.12 DPS, sim-verified) [crafted]; Robe of the Moccasin (6465, -0.26 DPS) [dungeon]; Corsair's Overshirt (5202, -0.29 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 24.5 healing_power points (0.94 DPS) | yes | Featherbead Bracers (15452, -0.54 DPS) [quest]; Bright Bracers (3647, -0.62 DPS) [world_drop]; Owl Bracers (4796, -0.63 DPS, sim-verified) [vendor] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 21.5 healing_power points (0.83 DPS) | yes | Wisdom's Leather Gloves (252499, +0.00 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.16 DPS) [crafted]; Magefist Gloves (12977, -0.21 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 19.4 healing_power points (0.75 DPS) | yes | Novice Ardent's Sash (253887, -0.11 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.36 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.0 healing_power points (1.46 DPS) | yes | Filigreed Pristine Leggings (253937, -0.12 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.58 DPS) [world_drop]; Ghastly Trousers (15449, -0.70 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 21.5 healing_power points (0.83 DPS) | yes | Black Whelp Slippers (252424, -0.13 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.24 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.42 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 12.6 healing_power points (0.48 DPS) | yes | Advisor's Ring (20426, -0.12 DPS) [rep]; Band of Purification (12996, -0.17 DPS) [world_drop]; Deep Fathom Ring (6463, -0.22 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 10.5 healing_power points (0.40 DPS) | yes | Advisor's Ring (20426, +0.00 DPS, sim-verified) [rep]; Band of Purification (12996, -0.09 DPS) [world_drop]; Deep Fathom Ring (6463, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 24.6 healing_power points (0.95 DPS) | yes | Advisor's Gnarled Staff (20425, -0.08 DPS) [pvp]; Gnarled Necromancer's Staff (251534, -0.12 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -0.14 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 206376 Bulwark Icon; 206383 Lithic Icon; 206384 Astral Icon; 206385 Clastic Icon; 206389 Brimstone Icon; 206390 Ichthyan Icon; 206391 Syzygic Icon; 241089 Scarlet Dagger

### Band 30 (tauren, 0000000000000000-000000000000000000-5532003300000000)

Set DPS (verified): 42.5. Weights run: 4.8s. Verify run: 16.5s. 368 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.993 ± 0.007, spirit=2.468 ± 0.013, mp5=6.283 ± 0.014, crit=0.174 ± 0.008 per rating point (14 rating = 1%, 2.431 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.2 healing_power points (2.03 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.53 DPS) [world_drop]; Dokebi Hat (14584, -0.62 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.8 healing_power points (0.74 DPS) | yes | Pendant of Myzrael (4614, -0.13 DPS) [dungeon]; Glowing Green Talisman (5002, -0.13 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.37 DPS, sim-verified) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 31.2 healing_power points (1.29 DPS) | yes | Nightsky Mantle (4718, -0.15 DPS, sim-verified) [world_drop]; Talbar Mantle (10657, -0.28 DPS) [quest]; Desert Shoulders (15457, -0.37 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 26.7 healing_power points (1.11 DPS) | yes | Darkspear Raider's Cloak (272078, -0.14 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19529, -0.16 DPS) [rep]; Amy's Blanket (13005, -0.39 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.8 healing_power points (1.85 DPS) | yes | Beguiler Robes (7728, -0.05 DPS) [dungeon]; Death Speaker Robes (6682, -0.12 DPS) [dungeon]; Pristine Gown (253961, -0.18 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 28.8 healing_power points (1.19 DPS) | yes | Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Nightsky Wristbands (6407, -0.52 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.53 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 39.6 healing_power points (1.64 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.48 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.55 DPS) [quest]; Tattered Mittens (270030, -0.61 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.8 healing_power points (1.89 DPS) | yes | Lilac Sash (6780, -0.85 DPS) [quest]; Highlander's Mail Girdle (20120, -0.90 DPS) [vendor]; Silver-lined Belt (13011, -1.26 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.2 healing_power points (2.16 DPS) | yes | Pristine Leggings (253987, -0.32 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.42 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.8 healing_power points (1.81 DPS) | yes | Acidic Walkers (9454, -0.75 DPS) [dungeon]; Frothing Slippers (254003, -0.83 DPS) [crafted]; Soggy Boots (274747, -1.67 DPS, sim-verified) [vendor] |
| finger1 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 25.1 healing_power points (1.04 DPS) | yes | Electrocutioner Lagnut (9447, -0.08 DPS) [dungeon]; Runemender's Seal (273031, -0.08 DPS) [world]; Solid Gold Band (279845, -0.24 DPS) [quest] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 23.7 healing_power points (0.98 DPS) | yes | Runemender's Seal (273031, -0.02 DPS) [world]; Solid Gold Band (279845, -0.18 DPS) [quest]; Electrocutioner Lagnut (9447, -0.20 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (42.5 DPS) | yes | Advisor's Gnarled Staff (19569, -1.20 DPS) [pvp]; Lorekeeper's Staff (212580, -1.28 DPS) [vendor]; Manual Crowd Pummeler (9449, -1.31 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 27.8 healing_power points (1.15 DPS) | yes | Orb of Mistmantle (13031, -0.37 DPS) [world_drop]; Insignia Buckler (4066, -0.44 DPS) [world_drop]; Defective Samophlange (274743, -0.57 DPS, sim-verified) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Darkspear Signet; finger2: The Queen's Jewel; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 368, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-000000000000000000-5532003315400000)

Set DPS (verified): 73.5. Weights run: 5.7s. Verify run: 23.2s. 567 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.563 ± 0.017, spirit=2.761 ± 0.012, mp5=7.128 ± 0.021, crit=0.285 ± 0.015 per rating point (14 rating = 1%, 3.996 per %), spell_haste=not significant (0.016 ± 0.016)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 82.4 healing_power points (3.39 DPS) | yes | Whitemane's Chapeau (7720, -0.78 DPS, sim-verified) [dungeon]; Imperial Leather Helm (6433, -0.88 DPS) [dungeon]; Enduring Cap (3020, -0.91 DPS) [world_drop] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 37.3 healing_power points (1.53 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.40 DPS) [quest]; Kaleidoscope Chain (13084, -0.66 DPS) [world_drop] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 57.6 healing_power points (2.37 DPS) | yes | Inquisitor's Shawl (19507, -0.38 DPS) [dungeon]; Mistscape Mantle (4734, -0.64 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.72 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 39.2 healing_power points (1.61 DPS) | yes | Battle Healer's Cloak (19528, -0.31 DPS) [rep]; Blackforge Cape (6424, -0.32 DPS) [dungeon]; Cloak of Blight (6832, -0.58 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 80.1 healing_power points (3.29 DPS) | yes | Dreamweave Vest (10021, -0.76 DPS) [crafted]; Doomsayer's Robe (4746, -0.84 DPS) [quest]; Deathchill Armor (10764, -1.31 DPS, sim-verified) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 37.9 healing_power points (1.56 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Windtalker's Wristguards (19584, -0.36 DPS) [pvp]; Dryad's Wrist Bindings (19597, -0.36 DPS) [pvp] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 51.6 healing_power points (2.12 DPS) | yes | Mender's Leather Gloves (252530, -0.31 DPS) [crafted]; Stormcloth Gloves (10011, -0.40 DPS) [crafted]; Gloves of Old (9395, -0.67 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 59.1 healing_power points (2.43 DPS) | yes | Mender's Leather Belt (252523, -0.51 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.56 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 70.0 healing_power points (2.88 DPS) | yes | Stoneweaver Leggings (9407, -0.14 DPS) [dungeon]; Misplaced Pantaloons (276201, -0.24 DPS, sim-verified) [vendor]; Wisdom's Leather Leggings (252519, -0.48 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (73.5 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -0.50 DPS) [crafted]; Furen's Boots (13100, -0.85 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.8 healing_power points (1.76 DPS) | yes | The Queen's Jewel (13094, -0.64 DPS) [world_drop]; Ogremind Ring (1993, -0.68 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 32.1 healing_power points (1.32 DPS) | yes | Ogremind Ring (1993, -0.24 DPS) [world_drop]; Voodoo Band (1996, -0.24 DPS) [world]; The Queen's Jewel (13094, -0.34 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -1.54 DPS) [dungeon]; Hypnotic Blade (7714, -1.78 DPS) [dungeon]; Gut Ripper (2164, -8.02 DPS, sim-verified) [world_drop] |
| off_hand | Battered Viking Shield (9403) | Uldaman: Olaf [dungeon] | 35.6 healing_power points (1.46 DPS) | yes | Ravager's Shield (14777, -0.05 DPS) [world_drop]; Forcestone Buckler (17508, -0.06 DPS) [quest]; Orb of Souls (249395, -0.25 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Battered Viking Shield

No-known-source sample (15 of 567, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (tauren, 0000000000000000-000000000000000000-5532003315513131)

Set DPS (verified): 124.3. Weights run: 9.7s. Verify run: 40.9s. 717 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.070, intellect=2.374 ± 0.023, spirit=2.996 ± 0.018, mp5=7.513 ± 0.025, crit=0.449 ± 0.025 per rating point (14 rating = 1%, 6.293 per %), spell_haste=not significant (0.132 ± 0.076)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 103.3 healing_power points (5.29 DPS) | yes | Winged Helm (13112, -0.96 DPS) [world_drop]; Serpentskin Helm (8261, -1.04 DPS) [world_drop]; Engineer's Guild Headpiece (9534, -1.39 DPS, sim-verified) [quest] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 59.3 healing_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.60 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.11 DPS) [dungeon]; Choker of the High Shaman (4112, -1.20 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 71.4 healing_power points (3.66 DPS) | yes | Living Shoulders (15061, +0.00 DPS, sim-verified) [crafted]; Mender's Leather Shoulders (252538, -0.26 DPS) [crafted]; Mender's Mail Shoulders (252569, -0.26 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 48.2 healing_power points (2.47 DPS) | yes | Battle Healer's Cloak (19527, -0.42 DPS) [rep]; Cloak of Blight (6832, -0.47 DPS) [quest]; Imperial Red Cloak (8248, -0.52 DPS) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 120.0 healing_power points (6.15 DPS) | yes | Feathered Breastplate (8349, -1.07 DPS, sim-verified) [crafted]; Stormcloth Vest (10020, -1.95 DPS) [crafted]; Deathchill Armor (10764, -2.09 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 53.6 healing_power points (2.75 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.72 DPS) [world_drop] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 104.0 healing_power points (5.33 DPS) | yes | Mender's Leather Gauntlets (252551, -0.77 DPS, sim-verified) [crafted]; Mender's Mail Gauntlets (252587, -1.27 DPS) [crafted]; Gilded Gloves (254095, -1.40 DPS) [crafted] |
| waist | Bloodlust Belt (14803) | World drop [world_drop] | 60.2 healing_power points (3.09 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -0.06 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.09 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 102.6 healing_power points (5.26 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.17 DPS, sim-verified) [dungeon]; Dalewind Trousers (13008, -0.33 DPS) [world_drop]; Senior Designer's Pantaloons (11841, -0.36 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 78.9 healing_power points (4.04 DPS) | yes | Furen's Boots (13100, -0.66 DPS, sim-verified) [world_drop]; Vinerot Sandals (17748, -0.80 DPS) [dungeon]; Mender's Leather Shoes (252533, -0.89 DPS) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 52.6 healing_power points (2.70 DPS) | yes | Coldwater Ring (4550, -0.95 DPS) [quest]; Welken Ring (5011, -1.01 DPS) [world_drop]; Ring of Subtlety (19038, -1.16 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 50.2 healing_power points (2.57 DPS) | yes | Welken Ring (5011, -0.89 DPS) [world_drop]; Coldwater Ring (4550, -0.90 DPS, sim-verified) [quest]; Ring of Subtlety (19038, -1.03 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (124.3 DPS) | yes | Ankh of Life (1713, -3.93 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.24 DPS) [quest]; Alchemist's Stone (13503, -5.16 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (124.3 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.31 DPS) [crafted]; Ankh of Life (1713, -0.49 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (124.3 DPS) | yes | Charstone Dirk (17710, -4.54 DPS) [dungeon]; Hand of Righteousness (7721, -6.55 DPS) [dungeon]; Hammer of the Northern Wind (810, -20.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Bloodlust Belt; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 717, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (tauren, 5300000000000000-000000000000000000-5532003315513151)

Set DPS (verified): 234.0. Weights run: 11.4s. Verify run: 128.6s. 1407 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.042, intellect=5.110 ± 0.061, spirit=4.488 ± 0.045, mp5=11.210 ± 0.052, crit=0.755 ± 0.051 per rating point (14 rating = 1%, 10.576 per %), spell_haste=not significant (0.197 ± 0.160)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 259.3 healing_power points (11.61 DPS) | yes | Living Crown (252561, +0.00 DPS, sim-verified) [crafted]; Crown of the Penitent (13216, -1.74 DPS) [quest]; Devout Crown (16693, -2.21 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 144.3 healing_power points (6.46 DPS) | yes | Frostwolf Advisor's Pendant (19096, -2.17 DPS) [rep]; Vision of Voodress (19608, -2.37 DPS) [quest]; Darkmoon Necklace (19303, -4.27 DPS, sim-verified) [vendor] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 171.1 healing_power points (7.66 DPS) | yes | Mantle of The Five Thunders (227011, -1.29 DPS) [vendor]; Drakesfire Epaulets (13133, -1.45 DPS) [world_drop]; Devout Mantle (16695, -3.09 DPS, sim-verified) [dungeon] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | 108.7 healing_power points (4.87 DPS) | yes | Faded Hakkari Cloak (20218, -0.02 DPS) [quest]; Hide of the Wild (18510, -0.70 DPS) [crafted]; Drape of Recovery (272413, -1.47 DPS) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 213.5 healing_power points (9.56 DPS) | yes | Devout Robe (16690, -0.33 DPS, sim-verified) [dungeon]; Tunic of The Five Thunders (227016, -0.75 DPS) [vendor]; Dreamwalker Armor (13123, -0.96 DPS) [world_drop] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 128.1 healing_power points (5.73 DPS) | yes | Bracers of The Five Thunders (227009, -0.75 DPS, sim-verified) [vendor]; Devout Bracers (16697, -1.40 DPS) [dungeon]; General's Mail Bracers (16576, -1.44 DPS) [pvp] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | sim-verified (234.0 DPS) | yes | Devout Gloves (16692, -0.83 DPS) [dungeon]; Raider Handwraps (272098, -1.19 DPS) [vendor] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 165.3 healing_power points (7.40 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Devout Belt (16696, -0.44 DPS, sim-verified) [dungeon]; Elunarian Belt (14465, -0.82 DPS) [world_drop] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | sim-verified (234.0 DPS) | yes | Outrider's Mail Leggings (22676, -0.15 DPS) [rep]; Red Dragonscale Leggings (252603, -0.45 DPS) [crafted]; Devout Skirt (16694, -5.60 DPS, sim-verified) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 198.6 healing_power points (8.89 DPS) | yes | Mooncloth Boots (15802, -2.00 DPS) [crafted]; Boots of The Five Thunders (22096, -2.13 DPS) [quest]; Incandescent Mooncloth Boots (227862, -5.27 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 137.0 healing_power points (6.14 DPS) | yes | Channeler's Ring (272406, -1.32 DPS) [vendor]; Darkspear Signet (272068, -1.62 DPS) [vendor]; Band of the Hierophant (13096, -1.64 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 117.9 healing_power points (5.28 DPS) | yes | Darkspear Signet (272068, -0.76 DPS) [vendor]; Band of the Hierophant (13096, -0.78 DPS) [world_drop]; Channeler's Ring (272406, -2.44 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (234.0 DPS) | yes | Ankh of Life (1713, -5.12 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -5.52 DPS) [quest]; Serenity Field (272439, -5.65 DPS) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (234.0 DPS) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -0.60 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (234.0 DPS) | yes | Staff of Hale Magefire (13000, -5.02 DPS) [world_drop]; Charstone Dirk (17710, -6.44 DPS) [dungeon]; Hand of Edward the Odd (2243, -12.31 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Darkspear Raider's Cloak; chest: Mooncloth Vest; wrist: Bracers of Hope; waist: Belt of Tiny Heads; legs: Leggings of The Five Thunders; feet: Greaves of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1407, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (tauren, 5300000000000000-000000000000000000-5532003315513151)

Set DPS (verified): 629.3. Weights run: 8.2s. Verify run: 53.0s. 1407 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.144, intellect=1.144 ± 0.028, spirit=0.728 ± 0.022, mp5=2.098 ± 0.027, crit=0.481 ± 0.026 per rating point (14 rating = 1%, 6.740 per %), spell_haste=not significant (0.960 ± 0.372)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 103.8 healing_power points (30.32 DPS) | yes | Crown of The Five Thunders (227013, +0.00 DPS, sim-verified) [vendor]; Sanctified Leather Helm (22689, -9.73 DPS) [quest]; Mooncloth Circlet (14140, -12.23 DPS) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 47.7 healing_power points (13.92 DPS) | yes | Drake Tooth Necklace (21531, -2.98 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -7.45 DPS) [dungeon]; Frostwolf Advisor's Pendant (19096, -8.12 DPS) [rep] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 78.9 healing_power points (23.04 DPS) | yes | Mantle of The Five Thunders (227011, -0.80 DPS, sim-verified) [vendor]; Champion's Mail Epaulets (227166, -8.03 DPS) [vendor]; Warlord's Mail Epaulets (231665, -8.71 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 53.4 healing_power points (15.61 DPS) | yes | Drape of Recovery (272413, -3.98 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19526, -6.31 DPS) [rep] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | 73.0 healing_power points (21.31 DPS) | yes | Legionnaire's Mail Chestguard (227171, -2.52 DPS) [vendor]; Mooncloth Vest (14138, -2.99 DPS) [crafted]; Stormcloth Vest (10020, -6.01 DPS, sim-verified) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 39.3 healing_power points (11.49 DPS) | yes | Bracers of The Five Thunders (227009, +0.00 DPS, sim-verified) [vendor]; Nethergeld Cuffs (254061, -2.24 DPS) [crafted]; Mender's Leather Bracers (252543, -2.32 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 54.9 healing_power points (16.03 DPS) | yes | General's Mail Gloves (231666, -1.46 DPS) [vendor]; Grasp of The Five Thunders (227014, -1.52 DPS) [vendor]; Mooncloth Gloves (18409, -1.92 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 53.6 healing_power points (15.65 DPS) | yes | Sash of The Five Thunders (227010, +0.00 DPS, sim-verified) [vendor]; Mender's Leather Waistguard (252477, -2.88 DPS) [crafted]; Mender's Mail Belt (252591, -2.88 DPS) [crafted] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | 62.4 healing_power points (18.23 DPS) | yes | Legionnaire's Mail Pants (227167, +0.00 DPS) [vendor]; General's Mail Pants (231670, -0.48 DPS) [vendor]; Devout Skirt (16694, -0.64 DPS, sim-verified) [dungeon] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 70.2 healing_power points (20.51 DPS) | yes | Greaves of The Five Thunders (227015, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -6.05 DPS) [crafted]; General's Mail Boots (231667, -7.57 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 55.8 healing_power points (16.30 DPS) | yes | Blessed Band of Light (272407, -5.54 DPS) [vendor]; Channeler's Ring (272406, -9.91 DPS) [vendor]; Band of the Hierophant (13096, -10.61 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 41.8 healing_power points (12.22 DPS) | yes | Blessed Band of Light (272407, -2.32 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -5.84 DPS) [vendor]; Band of the Hierophant (13096, -6.54 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (629.1 DPS) | yes | Serenity Field (272439, -3.75 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.94 DPS) [world_drop]; Ankh of Life (1713, -6.64 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (629.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -1.12 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (629.1 DPS) | yes | Hand of Edward the Odd (2243, -16.59 DPS, sim-verified) [world_drop]; Charstone Dirk (17710, -16.64 DPS) [dungeon]; Hand of Righteousness (7721, -26.26 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Leggings of The Five Thunders; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1407, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

