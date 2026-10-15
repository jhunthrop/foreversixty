# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5032100000000000)

Set DPS (verified): 27.3. Weights run: 1.4s. Verify run: 4.0s. 217 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.853 ± 0.010, spirit=1.848 ± 0.005, mp5=4.991 ± 0.010, crit=0.119 ± 0.008 per rating point (14 rating = 1%, 1.664 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 20.4 healing_power points (0.87 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.14 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.16 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 7.4 healing_power points (0.32 DPS) | yes | Scholarly Pendant (277203, -0.11 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.7 healing_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 12.7 healing_power points (0.55 DPS) | yes | Seer's Cape (6378, -0.23 DPS) [dungeon]; Spirit Cloak (4792, -0.23 DPS) [vendor]; Sanguine Cape (14376, -0.42 DPS, sim-verified) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.8 healing_power points (1.28 DPS) | yes | Robe of the Moccasin (6465, -0.10 DPS) [dungeon]; Filigreed Pristine Gown (253901, -0.13 DPS, sim-verified) [crafted]; Corsair's Overshirt (5202, -0.17 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 24.2 healing_power points (1.04 DPS) | yes | Owl Bracers (4796, -0.64 DPS) [vendor]; Drakewing Bands (12999, -0.68 DPS, sim-verified) [world_drop]; Bright Bracers (3647, -0.72 DPS) [world_drop] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 17.4 healing_power points (0.75 DPS) | yes | Pristine Gloves (253913, -0.04 DPS) [crafted]; Magefist Gloves (12977, -0.08 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.11 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.4 healing_power points (0.79 DPS) | yes | Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.15 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.23 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.5 healing_power points (1.65 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.62 DPS) [world_drop]; Dreamer's Leggings (270016, -0.78 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.3 healing_power points (0.87 DPS) | yes | Black Whelp Slippers (252424, -0.14 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.24 DPS) [crafted]; Kimbra Boots (6191, -0.31 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.1 healing_power points (0.48 DPS) | yes | Lorekeeper's Ring (20431, -0.05 DPS) [rep]; Deep Fathom Ring (6463, -0.08 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 11.1 healing_power points (0.48 DPS) | yes | Lorekeeper's Ring (20431, -0.05 DPS) [rep]; Deep Fathom Ring (6463, -0.08 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 20.4 healing_power points (0.87 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Staff of the Blessed Seer (2271, -0.08 DPS) [dungeon]; Taskmaster Axe (5194, -0.24 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Staff of Westfall

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (dwarf, 0000000000000000-000000000000000000-5032503300000000)

Set DPS (verified): 42.4. Weights run: 1.3s. Verify run: 4.1s. 383 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.998 ± 0.009, spirit=2.478 ± 0.018, mp5=6.276 ± 0.020, crit=0.174 ± 0.011 per rating point (14 rating = 1%, 2.441 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.3 healing_power points (2.04 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.54 DPS) [world_drop]; Dokebi Hat (14584, -0.62 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.9 healing_power points (0.74 DPS) | yes | Crystal Starfire Medallion (5003, -0.08 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.13 DPS) [dungeon]; Glowing Green Talisman (5002, -0.13 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 31.3 healing_power points (1.30 DPS) | yes | Nightsky Mantle (4718, -0.12 DPS) [world_drop]; Faerie Mantle (5820, -0.19 DPS) [quest]; Talbar Mantle (10657, -0.28 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 26.8 healing_power points (1.11 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Darkspear Raider's Cloak (272078, -0.14 DPS) [vendor]; Caretaker's Cape (19533, -0.16 DPS) [rep] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.9 healing_power points (1.86 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.04 DPS) [dungeon]; Death Speaker Robes (6682, -0.12 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 28.8 healing_power points (1.19 DPS) | yes | Nightsky Wristbands (6407, -0.35 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.53 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 39.7 healing_power points (1.64 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.16 DPS, sim-verified) [dungeon]; Zodiac Gloves (7106, -0.60 DPS) [quest]; Silver-thread Gloves (6393, -0.72 DPS) [world_drop] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.9 healing_power points (1.90 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.91 DPS) [rep]; Highlander's Mail Girdle (20120, -0.91 DPS) [vendor]; Silver-lined Belt (13011, -1.27 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.3 healing_power points (2.17 DPS) | yes | Pristine Leggings (253987, -0.16 DPS) [crafted]; Earthen Leggings (253999, -0.42 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.9 healing_power points (1.82 DPS) | yes | Acidic Walkers (9454, -0.75 DPS) [dungeon]; Frothing Slippers (254003, -0.83 DPS) [crafted]; Soggy Boots (274747, -1.37 DPS, sim-verified) [vendor] |
| finger1 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 25.1 healing_power points (1.04 DPS) | yes | Electrocutioner Lagnut (9447, -0.07 DPS) [dungeon]; Runemender's Seal (273031, -0.07 DPS) [world]; Ring of Power Regulation (273030, -0.26 DPS) [dungeon] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 23.8 healing_power points (0.99 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS, sim-verified) [dungeon]; Runemender's Seal (273031, -0.02 DPS) [world]; Ring of Power Regulation (273030, -0.21 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (42.4 DPS) | yes | Hardened Root Staff (1317, -1.06 DPS) [quest]; Manual Crowd Pummeler (9449, -1.13 DPS, sim-verified) [dungeon]; Gnarled Ash Staff (791, -1.28 DPS) [world_drop] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 27.9 healing_power points (1.15 DPS) | yes | Defective Samophlange (274743, -0.24 DPS, sim-verified) [vendor]; Orb of Mistmantle (13031, -0.37 DPS) [world_drop]; Insignia Buckler (4066, -0.44 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Darkspear Signet; finger2: The Queen's Jewel; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 383, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots

### Band 40 (dwarf, 0000000000000000-000000000000000000-5032503315400000)

Set DPS (verified): 73.8. Weights run: 1.4s. Verify run: 6.1s. 597 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.557 ± 0.024, spirit=2.749 ± 0.016, mp5=7.117 ± 0.029, crit=0.276 ± 0.021 per rating point (14 rating = 1%, 3.865 per %), spell_haste=not significant (0.032 ± 0.032)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 82.2 healing_power points (3.37 DPS) | yes | Enduring Cap (3020, -0.91 DPS) [world_drop]; Silksand Circlet (14421, -1.30 DPS) [world_drop]; Imperial Leather Helm (6433, -1.49 DPS, sim-verified) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 37.1 healing_power points (1.52 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.40 DPS) [quest]; Explorers' League Commendation (7746, -0.51 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 57.4 healing_power points (2.36 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS, sim-verified) [dungeon]; Mistscape Mantle (4734, -0.64 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.71 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 39.1 healing_power points (1.61 DPS) | yes | Sergeant Major's Cape (16336, -0.23 DPS, sim-verified) [pvp]; Caretaker's Cape (19532, -0.30 DPS) [rep]; Blackforge Cape (6424, -0.31 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 80.0 healing_power points (3.28 DPS) | yes | Icemail Jerkin (1981, -0.58 DPS) [world_drop]; Dreamweave Vest (10021, -0.76 DPS) [crafted]; Deathchill Armor (10764, -1.48 DPS, sim-verified) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 37.7 healing_power points (1.55 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Silkstream Cuffs (16791, -0.57 DPS) [quest]; Earthen Silk Cuffs (254019, -0.57 DPS) [crafted] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 51.5 healing_power points (2.11 DPS) | yes | Mender's Leather Gloves (252530, -0.31 DPS) [crafted]; Stormcloth Gloves (10011, -0.40 DPS) [crafted]; Gloves of Old (9395, -0.56 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 58.9 healing_power points (2.42 DPS) | yes | Mender's Leather Belt (252523, -0.40 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.56 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 69.8 healing_power points (2.87 DPS) | yes | Stoneweaver Leggings (9407, -0.14 DPS) [dungeon]; Misplaced Pantaloons (276201, -0.36 DPS, sim-verified) [vendor]; Wisdom's Leather Leggings (252519, -0.48 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -0.50 DPS) [crafted]; Furen's Boots (13100, -0.92 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.7 healing_power points (1.75 DPS) | yes | The Queen's Jewel (13094, -0.64 DPS) [world_drop]; Ogremind Ring (1993, -0.68 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 32.0 healing_power points (1.31 DPS) | yes | Ogremind Ring (1993, -0.24 DPS) [world_drop]; Voodoo Band (1996, -0.24 DPS) [world]; The Queen's Jewel (13094, -0.25 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -1.54 DPS) [dungeon]; Hypnotic Blade (7714, -1.78 DPS) [dungeon] |
| off_hand | Battered Viking Shield (9403) | Uldaman: Olaf [dungeon] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ravager's Shield (14777, -0.05 DPS) [world_drop]; Orb of Souls (249395, -0.25 DPS) [crafted]; Orb of Lorica (11262, -7.50 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Battered Viking Shield

No-known-source sample (15 of 597, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 50 (dwarf, 0000000000000000-000000000000000000-5032503315513131)

Set DPS (verified): 120.2. Weights run: 1.7s. Verify run: 9.1s. 765 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.080, intellect=2.455 ± 0.035, spirit=3.099 ± 0.027, mp5=7.743 ± 0.036, crit=0.433 ± 0.033 per rating point (14 rating = 1%, 6.058 per %), spell_haste=not significant (0.090 ± 0.106)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 106.5 healing_power points (5.29 DPS) | yes | Engineer's Guild Headpiece (9534, -0.96 DPS) [quest]; Winged Helm (13112, -1.03 DPS, sim-verified) [world_drop]; Serpentskin Helm (8261, -1.03 DPS) [world_drop] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 61.2 healing_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.65 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.11 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.42 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 73.9 healing_power points (3.67 DPS) | yes | Living Shoulders (15061, +0.00 DPS, sim-verified) [crafted]; Mender's Leather Shoulder (252538, -0.30 DPS) [crafted]; Mender's Mail Shoulder (252569, -0.30 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 49.9 healing_power points (2.48 DPS) | yes | Caretaker's Cape (19531, -0.21 DPS, sim-verified) [rep]; Imperial Red Cloak (8248, -0.52 DPS) [world_drop]; Wingveil Cloak (10802, -0.53 DPS) [dungeon] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 123.9 healing_power points (6.15 DPS) | yes | Feathered Breastplate (8349, -1.13 DPS, sim-verified) [crafted]; Stormcloth Vest (10020, -2.04 DPS) [crafted]; Deathchill Armor (10764, -2.10 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 54.8 healing_power points (2.72 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.69 DPS) [world_drop] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 106.8 healing_power points (5.30 DPS) | yes | Mender's Leather Gauntlets (252551, -0.85 DPS, sim-verified) [crafted]; Mender's Mail Gauntlets (252587, -1.28 DPS) [crafted]; Gilded Gloves (254095, -1.41 DPS) [crafted] |
| waist | Bloodlust Belt (14803) | World drop [world_drop] | 62.3 healing_power points (3.09 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -0.11 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.14 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 106.0 healing_power points (5.26 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.32 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.40 DPS) [dungeon]; Dalewind Trousers (13008, -1.39 DPS, sim-verified) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 81.6 healing_power points (4.05 DPS) | yes | Furen's Boots (13100, -0.62 DPS, sim-verified) [world_drop]; Vinerot Sandals (17748, -0.82 DPS) [dungeon]; Mender's Leather Shoes (252533, -0.93 DPS) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 54.2 healing_power points (2.69 DPS) | yes | Choking Band (11868, -0.69 DPS) [quest]; Coldwater Ring (4550, -0.94 DPS) [quest]; Welken Ring (5011, -1.00 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 51.9 healing_power points (2.58 DPS) | yes | Coldwater Ring (4550, -0.83 DPS) [quest]; Welken Ring (5011, -0.89 DPS) [world_drop]; Choking Band (11868, -1.03 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (120.2 DPS) | yes | Ankh of Life (1713, -3.92 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.23 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.54 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (120.2 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.57 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (120.2 DPS) | yes | Hand of Righteousness (7721, -1.97 DPS) [dungeon]; Spire of Hakkar (10844, -2.33 DPS) [world]; Hammer of the Northern Wind (810, -14.36 DPS, sim-verified) [world_drop] |
| off_hand | Mountainside Buckler (13082) | World drop [world_drop] | 49.5 healing_power points (2.46 DPS) | yes | Protector Buckler (14790, -0.47 DPS) [world_drop]; Battered Viking Shield (9403, -0.52 DPS) [dungeon]; Twisting Essence Jar (249456, -13.76 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Bloodlust Belt; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; off_hand: Mountainside Buckler

No-known-source sample (15 of 765, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60 (dwarf, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 223.1. Weights run: 1.8s. Verify run: 22.4s. 1473 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.129, intellect=5.471 ± 0.104, spirit=5.362 ± 0.090, mp5=13.450 ± 0.105, crit=0.857 ± 0.086 per rating point (14 rating = 1%, 11.996 per %), spell_haste=not significant (-0.260 ± 0.452)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 290.6 healing_power points (11.69 DPS) | yes | Living Crown (252561, +0.00 DPS, sim-verified) [crafted]; Crown of the Penitent (13216, -1.82 DPS) [quest]; Devout Crown (16693, -2.37 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 164.4 healing_power points (6.61 DPS) | yes | Stormpike Sage's Pendant (19098, -2.25 DPS) [rep]; Horizon Choker (13085, -2.67 DPS) [world_drop]; Darkmoon Necklace (19303, -3.98 DPS, sim-verified) [vendor] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 193.3 healing_power points (7.77 DPS) | yes | Mantle of The Five Thunders (227011, -1.44 DPS) [vendor]; Drakesfire Epaulets (13133, -1.46 DPS) [world_drop]; Devout Mantle (16695, -3.38 DPS, sim-verified) [dungeon] |
| back | Faded Hakkari Cloak (20218) | Confront Yeh'kinya [quest] | 124.5 healing_power points (5.00 DPS) | yes | Darkspear Raider's Cloak (272063, -0.19 DPS) [vendor]; Hide of the Wild (18510, -1.12 DPS) [crafted] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 237.3 healing_power points (9.54 DPS) | yes | Devout Robe (16690, -0.41 DPS, sim-verified) [dungeon]; Tunic of The Five Thunders (227016, -0.82 DPS) [vendor]; Dreamwalker Armor (13123, -0.83 DPS) [world_drop] |
| wrist | Bracers of The Five Thunders (227009) | Mokvar [vendor] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Devout Bracers (16697, -0.72 DPS) [dungeon]; Magister's Bindings (16683, -0.92 DPS) [dungeon]; Bracers of Hope (22667, -4.11 DPS, sim-verified) [quest] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 180.7 healing_power points (7.27 DPS) | yes | Raider Handwraps (272097, +0.00 DPS, sim-verified) [vendor]; Devout Gloves (16692, -0.80 DPS) [dungeon]; Magister's Gloves (16684, -1.17 DPS) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 187.2 healing_power points (7.53 DPS) | yes | Devout Belt (16696, -0.74 DPS) [dungeon]; Elunarian Belt (14465, -0.97 DPS) [world_drop]; Sash of The Five Thunders (227010, -1.11 DPS) [vendor] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Kilt of the Atal'ai Prophet (10807, -0.57 DPS) [dungeon]; Red Dragonscale Leggings (252603, -0.80 DPS) [crafted]; Devout Skirt (16694, -5.21 DPS, sim-verified) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 225.3 healing_power points (9.06 DPS) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -2.21 DPS) [crafted]; Devout Sandals (16691, -2.75 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 151.5 healing_power points (6.09 DPS) | yes | Darkspear Signet (272068, -1.22 DPS) [vendor]; Channeler's Ring (272406, -1.29 DPS) [vendor]; Band of the Hierophant (13096, -1.52 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 130.5 healing_power points (5.25 DPS) | yes | Channeler's Ring (272406, -0.44 DPS) [vendor]; Band of the Hierophant (13096, -0.68 DPS) [world_drop]; Darkspear Signet (272068, -2.51 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -5.52 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -5.96 DPS) [quest]; Thunderbrew's Boot Flask (744, -6.39 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -0.66 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.09 DPS) [quest] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -1.92 DPS) [world_drop]; Charstone Dirk (17710, -2.08 DPS) [dungeon]; Hand of Edward the Odd (2243, -6.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Faded Hakkari Cloak; chest: Mooncloth Vest; wrist: Bracers of The Five Thunders; waist: Belt of Tiny Heads; legs: Leggings of The Five Thunders; feet: Greaves of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1473, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 623.8. Weights run: 1.6s. Verify run: 16.7s. 1473 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.213, intellect=1.026 ± 0.033, spirit=0.657 ± 0.025, mp5=1.860 ± 0.033, crit=0.528 ± 0.035 per rating point (14 rating = 1%, 7.395 per %), spell_haste=not significant (0.296 ± 0.449)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 100.7 healing_power points (32.70 DPS) | yes | Crown of The Five Thunders (227013, +0.00 DPS, sim-verified) [vendor]; Sanctified Leather Helm (22689, -10.81 DPS) [quest]; Mooncloth Circlet (14140, -12.93 DPS) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 45.0 healing_power points (14.63 DPS) | yes | Drake Tooth Necklace (21531, -3.11 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -7.74 DPS) [dungeon]; Stormpike Sage's Pendant (19098, -8.88 DPS) [rep] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 76.7 healing_power points (24.92 DPS) | yes | Mooncloth Shoulders (14139, -7.36 DPS, sim-verified) [crafted]; Mantle of The Five Thunders (227011, -10.88 DPS) [vendor]; Eternity Pauldrons (279255, -11.34 DPS) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.3 healing_power points (16.98 DPS) | yes | Drape of Recovery (272413, -2.52 DPS, sim-verified) [vendor]; Caretaker's Cape (19530, -6.82 DPS) [rep] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | 70.7 healing_power points (22.95 DPS) | yes | Mooncloth Vest (14138, -3.79 DPS) [crafted]; Red Dragonscale Breastplate (15047, -4.44 DPS) [crafted]; Stormcloth Vest (10020, -4.93 DPS, sim-verified) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 37.2 healing_power points (12.07 DPS) | yes | Bracers of The Five Thunders (227009, +0.00 DPS, sim-verified) [vendor]; Nethergeld Cuffs (254061, -2.18 DPS) [crafted]; Mender's Leather Bracers (252543, -2.29 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 51.7 healing_power points (16.80 DPS) | yes | Grasp of The Five Thunders (227014, -1.49 DPS) [vendor]; Mender's Leather Gauntlets (252551, -1.81 DPS) [crafted]; Mooncloth Gloves (18409, -2.17 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 51.3 healing_power points (16.66 DPS) | yes | Sash of The Five Thunders (227010, +0.00 DPS, sim-verified) [vendor]; Mender's Leather Waistguard (252477, -2.91 DPS) [crafted]; Mender's Mail Belt (252591, -2.91 DPS) [crafted] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | 59.2 healing_power points (19.24 DPS) | yes | Devout Skirt (16694, -1.54 DPS) [dungeon]; Red Dragonscale Leggings (252603, -2.08 DPS) [crafted]; Elderwild Pants (279254, -2.28 DPS, sim-verified) [crafted] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 67.6 healing_power points (21.95 DPS) | yes | Greaves of The Five Thunders (227015, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -6.71 DPS) [crafted]; Swarmtender's Footpads (275609, -7.98 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 53.5 healing_power points (17.39 DPS) | yes | Blessed Band of Light (272407, -5.61 DPS) [vendor]; Channeler's Ring (272406, -11.02 DPS) [vendor]; Sea Giant's Toe Ring (274746, -11.54 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 39.7 healing_power points (12.90 DPS) | yes | Blessed Band of Light (272407, -2.46 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -6.52 DPS) [vendor]; Sea Giant's Toe Ring (274746, -7.05 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -2.03 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.21 DPS) [world_drop]; Ankh of Life (1713, -6.50 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -1.06 DPS) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -10.32 DPS) [dungeon]; Spellshifter Rod (9527, -17.35 DPS) [quest] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | sim-verified (623.9 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, -0.04 DPS) [vendor]; Lei of the Lifegiver (19312, -22.19 DPS, sim-verified) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Leggings of The Five Thunders; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves; off_hand: Evergreen Shield

No-known-source sample (15 of 1473, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

## Horde

### Band 20 (tauren, 0000000000000000-000000000000000000-5032100000000000)

Set DPS (verified): 27.7. Weights run: 1.4s. Verify run: 3.9s. 197 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.853 ± 0.010, spirit=1.848 ± 0.005, mp5=4.991 ± 0.010, crit=0.119 ± 0.008 per rating point (14 rating = 1%, 1.664 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (27.7 DPS) | yes | Wisdom's Leather Hood (252507, -0.04 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.06 DPS) [crafted]; Shadow Goggles (4373, -0.30 DPS, sim-verified) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 5.5 healing_power points (0.24 DPS) | yes | Scholarly Pendant (277203, -0.11 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.7 healing_power points (0.72 DPS) | yes | Slime-encrusted Pads (6461, -0.07 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.40 DPS) [crafted]; Forest Leather Mantle (4709, -0.40 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 12.7 healing_power points (0.55 DPS) | yes | Traveler's Shawl (277289, -0.15 DPS, sim-verified) [quest]; Sanguine Cape (14376, -0.23 DPS) [world_drop]; Seer's Cape (6378, -0.23 DPS) [dungeon] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.8 healing_power points (1.28 DPS) | yes | Robe of the Moccasin (6465, -0.10 DPS) [dungeon]; Filigreed Pristine Gown (253901, -0.13 DPS, sim-verified) [crafted]; Corsair's Overshirt (5202, -0.17 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 24.2 healing_power points (1.04 DPS) | yes | Owl Bracers (4796, -0.64 DPS) [vendor]; Featherbead Bracers (15452, -0.64 DPS) [quest]; Drakewing Bands (12999, -0.73 DPS, sim-verified) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 22.2 healing_power points (0.95 DPS) | yes | Wisdom's Leather Gloves (252499, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.24 DPS) [world_drop]; Pristine Gloves (253913, -0.24 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.4 healing_power points (0.79 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.15 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 38.5 healing_power points (1.65 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.62 DPS) [world_drop]; Ghastly Trousers (15449, -0.83 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.3 healing_power points (0.87 DPS) | yes | Black Whelp Slippers (252424, -0.24 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.24 DPS) [crafted]; Bluegill Sandals (1560, -0.39 DPS) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.1 healing_power points (0.48 DPS) | yes | Advisor's Ring (20426, -0.05 DPS) [rep]; Loop of Sacrifice (281673, -0.08 DPS) [quest]; Deep Fathom Ring (6463, -0.08 DPS) [dungeon] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 11.1 healing_power points (0.48 DPS) | yes | Advisor's Ring (20426, -0.05 DPS) [rep]; Loop of Sacrifice (281673, -0.08 DPS) [quest]; Deep Fathom Ring (6463, -0.08 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 26.6 healing_power points (1.14 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.14 DPS, sim-verified) [quest]; Advisor's Gnarled Staff (20425, -0.18 DPS) [pvp]; Twisted Chanter's Staff (890, -0.35 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 206376 Bulwark Icon; 206383 Lithic Icon; 206384 Astral Icon; 206385 Clastic Icon; 206389 Brimstone Icon; 206390 Ichthyan Icon; 206391 Syzygic Icon; 241089 Scarlet Dagger

### Band 30 (tauren, 0000000000000000-000000000000000000-5032503300000000)

Set DPS (verified): 42.5. Weights run: 1.3s. Verify run: 3.9s. 363 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.998 ± 0.009, spirit=2.478 ± 0.018, mp5=6.276 ± 0.020, crit=0.174 ± 0.011 per rating point (14 rating = 1%, 2.441 per %), spell_haste=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.3 healing_power points (2.04 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.54 DPS) [world_drop]; Dokebi Hat (14584, -0.62 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.9 healing_power points (0.74 DPS) | yes | Pendant of Myzrael (4614, -0.13 DPS) [dungeon]; Glowing Green Talisman (5002, -0.13 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.37 DPS, sim-verified) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 31.3 healing_power points (1.30 DPS) | yes | Nightsky Mantle (4718, -0.15 DPS, sim-verified) [world_drop]; Talbar Mantle (10657, -0.28 DPS) [quest]; Desert Shoulders (15457, -0.37 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 26.8 healing_power points (1.11 DPS) | yes | Darkspear Raider's Cloak (272078, -0.14 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19529, -0.16 DPS) [rep]; Amy's Blanket (13005, -0.39 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.9 healing_power points (1.86 DPS) | yes | Beguiler Robes (7728, -0.04 DPS) [dungeon]; Death Speaker Robes (6682, -0.12 DPS) [dungeon]; Pristine Gown (253961, -0.18 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 28.8 healing_power points (1.19 DPS) | yes | Spidertank Oilrag (9448, -0.41 DPS) [dungeon]; Nightsky Wristbands (6407, -0.52 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.53 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 39.7 healing_power points (1.64 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.48 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.55 DPS) [quest]; Tattered Mittens (270030, -0.61 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.9 healing_power points (1.90 DPS) | yes | Lilac Sash (6780, -0.85 DPS) [quest]; Highlander's Mail Girdle (20120, -0.91 DPS) [vendor]; Silver-lined Belt (13011, -1.26 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 52.3 healing_power points (2.17 DPS) | yes | Pristine Leggings (253987, -0.32 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.42 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.9 healing_power points (1.82 DPS) | yes | Acidic Walkers (9454, -0.75 DPS) [dungeon]; Frothing Slippers (254003, -0.83 DPS) [crafted]; Soggy Boots (274747, -1.67 DPS, sim-verified) [vendor] |
| finger1 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 25.1 healing_power points (1.04 DPS) | yes | Electrocutioner Lagnut (9447, -0.07 DPS) [dungeon]; Runemender's Seal (273031, -0.07 DPS) [world]; Solid Gold Band (279845, -0.24 DPS) [quest] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 23.8 healing_power points (0.99 DPS) | yes | Runemender's Seal (273031, -0.02 DPS) [world]; Solid Gold Band (279845, -0.18 DPS) [quest]; Electrocutioner Lagnut (9447, -0.20 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (42.5 DPS) | yes | Advisor's Gnarled Staff (19569, -1.20 DPS) [pvp]; Gnarled Ash Staff (791, -1.28 DPS) [world_drop]; Manual Crowd Pummeler (9449, -1.31 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 27.9 healing_power points (1.15 DPS) | yes | Orb of Mistmantle (13031, -0.37 DPS) [world_drop]; Insignia Buckler (4066, -0.44 DPS) [world_drop]; Defective Samophlange (274743, -0.57 DPS, sim-verified) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Darkspear Signet; finger2: The Queen's Jewel; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 363, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-000000000000000000-5032503315400000)

Set DPS (verified): 73.5. Weights run: 1.4s. Verify run: 5.7s. 558 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=2.557 ± 0.024, spirit=2.749 ± 0.016, mp5=7.117 ± 0.029, crit=0.276 ± 0.021 per rating point (14 rating = 1%, 3.865 per %), spell_haste=not significant (0.032 ± 0.032)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 82.2 healing_power points (3.37 DPS) | yes | Enduring Cap (3020, -0.91 DPS) [world_drop]; Silksand Circlet (14421, -1.30 DPS) [world_drop]; Imperial Leather Helm (6433, -1.54 DPS, sim-verified) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 37.1 healing_power points (1.52 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.40 DPS) [quest]; Kaleidoscope Chain (13084, -0.65 DPS) [world_drop] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 57.4 healing_power points (2.36 DPS) | yes | Inquisitor's Shawl (19507, -0.38 DPS) [dungeon]; Mistscape Mantle (4734, -0.64 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.71 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 39.1 healing_power points (1.61 DPS) | yes | Battle Healer's Cloak (19528, -0.30 DPS) [rep]; Blackforge Cape (6424, -0.31 DPS) [dungeon]; Cloak of Blight (6832, -0.58 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 80.0 healing_power points (3.28 DPS) | yes | Icemail Jerkin (1981, -0.58 DPS) [world_drop]; Dreamweave Vest (10021, -0.76 DPS) [crafted]; Deathchill Armor (10764, -1.31 DPS, sim-verified) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 37.7 healing_power points (1.55 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Windtalker's Wristguards (19584, -0.35 DPS) [pvp]; Dryad's Wrist Bindings (19597, -0.35 DPS) [pvp] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 51.5 healing_power points (2.11 DPS) | yes | Mender's Leather Gloves (252530, -0.31 DPS) [crafted]; Stormcloth Gloves (10011, -0.40 DPS) [crafted]; Gloves of Old (9395, -0.67 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 58.9 healing_power points (2.42 DPS) | yes | Mender's Leather Belt (252523, -0.51 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.56 DPS) [world_drop]; Sutarn's Ring (13105, -0.69 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 69.8 healing_power points (2.87 DPS) | yes | Stoneweaver Leggings (9407, -0.14 DPS) [dungeon]; Misplaced Pantaloons (276201, -0.24 DPS, sim-verified) [vendor]; Wisdom's Leather Leggings (252519, -0.48 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (73.5 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -0.50 DPS) [crafted]; Furen's Boots (13100, -0.85 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.7 healing_power points (1.75 DPS) | yes | The Queen's Jewel (13094, -0.64 DPS) [world_drop]; Ogremind Ring (1993, -0.68 DPS) [world_drop]; Voodoo Band (1996, -0.68 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 32.0 healing_power points (1.31 DPS) | yes | Ogremind Ring (1993, -0.24 DPS) [world_drop]; Voodoo Band (1996, -0.24 DPS) [world]; The Queen's Jewel (13094, -0.34 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -1.54 DPS) [dungeon]; Hypnotic Blade (7714, -1.78 DPS) [dungeon]; Gut Ripper (2164, -8.02 DPS, sim-verified) [world_drop] |
| off_hand | Battered Viking Shield (9403) | Uldaman: Olaf [dungeon] | 35.5 healing_power points (1.46 DPS) | yes | Ravager's Shield (14777, -0.05 DPS) [world_drop]; Forcestone Buckler (17508, -0.06 DPS) [quest]; Orb of Souls (249395, -0.25 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Battered Viking Shield

No-known-source sample (15 of 558, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (tauren, 0000000000000000-000000000000000000-5032503315513131)

Set DPS (verified): 119.8. Weights run: 1.7s. Verify run: 9.3s. 708 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.080, intellect=2.455 ± 0.035, spirit=3.099 ± 0.027, mp5=7.743 ± 0.036, crit=0.433 ± 0.033 per rating point (14 rating = 1%, 6.058 per %), spell_haste=not significant (0.090 ± 0.106)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 106.5 healing_power points (5.29 DPS) | yes | Engineer's Guild Headpiece (9534, -0.96 DPS) [quest]; Serpentskin Helm (8261, -1.03 DPS) [world_drop]; Winged Helm (13112, -1.03 DPS, sim-verified) [world_drop] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 61.2 healing_power points (3.04 DPS) | yes | Horizon Choker (13085, -0.64 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -1.11 DPS) [dungeon]; Choker of the High Shaman (4112, -1.20 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 73.9 healing_power points (3.67 DPS) | yes | Living Shoulders (15061, +0.00 DPS, sim-verified) [crafted]; Mender's Leather Shoulder (252538, -0.30 DPS) [crafted]; Mender's Mail Shoulder (252569, -0.30 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 49.9 healing_power points (2.48 DPS) | yes | Battle Healer's Cloak (19527, -0.26 DPS, sim-verified) [rep]; Cloak of Blight (6832, -0.48 DPS) [quest]; Imperial Red Cloak (8248, -0.52 DPS) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 123.9 healing_power points (6.15 DPS) | yes | Feathered Breastplate (8349, -1.08 DPS, sim-verified) [crafted]; Stormcloth Vest (10020, -2.04 DPS) [crafted]; Deathchill Armor (10764, -2.10 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 54.8 healing_power points (2.72 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted]; Enchanted Kodo Bracers (13119, -0.69 DPS) [world_drop] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 106.8 healing_power points (5.30 DPS) | yes | Mender's Leather Gauntlets (252551, -0.81 DPS, sim-verified) [crafted]; Mender's Mail Gauntlets (252587, -1.28 DPS) [crafted]; Gilded Gloves (254095, -1.41 DPS) [crafted] |
| waist | Bloodlust Belt (14803) | World drop [world_drop] | 62.3 healing_power points (3.09 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Earthenweave Cord (254077, -0.11 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.14 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 106.0 healing_power points (5.26 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.32 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.40 DPS) [dungeon]; Dalewind Trousers (13008, -1.42 DPS, sim-verified) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 81.6 healing_power points (4.05 DPS) | yes | Furen's Boots (13100, -0.70 DPS, sim-verified) [world_drop]; Vinerot Sandals (17748, -0.82 DPS) [dungeon]; Mender's Leather Shoes (252533, -0.93 DPS) [crafted] |
| finger1 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 54.2 healing_power points (2.69 DPS) | yes | Coldwater Ring (4550, -0.94 DPS) [quest]; Welken Ring (5011, -1.00 DPS) [world_drop]; Ring of Subtlety (19038, -1.15 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 51.9 healing_power points (2.58 DPS) | yes | Welken Ring (5011, -0.89 DPS) [world_drop]; Coldwater Ring (4550, -0.97 DPS, sim-verified) [quest]; Ring of Subtlety (19038, -1.04 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (119.8 DPS) | yes | Ankh of Life (1713, -3.92 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.23 DPS) [quest]; Alchemist's Stone (13503, -5.15 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (119.8 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.28 DPS) [crafted]; Ankh of Life (1713, -0.61 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (119.8 DPS) | yes | Hand of Righteousness (7721, -1.97 DPS) [dungeon]; Spire of Hakkar (10844, -2.33 DPS) [world]; Hammer of the Northern Wind (810, -14.32 DPS, sim-verified) [world_drop] |
| off_hand | Mountainside Buckler (13082) | World drop [world_drop] | 49.5 healing_power points (2.46 DPS) | yes | Protector Buckler (14790, -0.47 DPS) [world_drop]; Battered Viking Shield (9403, -0.52 DPS) [dungeon]; Twisting Essence Jar (249456, -13.67 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Bloodlust Belt; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Darkspear Signet; finger2: Eye of Adaegus; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Charstone Dirk; off_hand: Mountainside Buckler

No-known-source sample (15 of 708, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (tauren, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 222.6. Weights run: 1.8s. Verify run: 21.7s. 1396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.129, intellect=5.471 ± 0.104, spirit=5.362 ± 0.090, mp5=13.450 ± 0.105, crit=0.857 ± 0.086 per rating point (14 rating = 1%, 11.996 per %), spell_haste=not significant (-0.260 ± 0.452)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 290.6 healing_power points (11.69 DPS) | yes | Living Crown (252561, +0.00 DPS, sim-verified) [crafted]; Crown of the Penitent (13216, -1.82 DPS) [quest]; Devout Crown (16693, -2.37 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 164.4 healing_power points (6.61 DPS) | yes | Frostwolf Advisor's Pendant (19096, -2.25 DPS) [rep]; Vision of Voodress (19608, -2.47 DPS) [quest]; Darkmoon Necklace (19303, -4.16 DPS, sim-verified) [vendor] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 193.3 healing_power points (7.77 DPS) | yes | Mantle of The Five Thunders (227011, -1.44 DPS) [vendor]; Drakesfire Epaulets (13133, -1.46 DPS) [world_drop]; Devout Mantle (16695, -3.25 DPS, sim-verified) [dungeon] |
| back | Faded Hakkari Cloak (20218) | Confront Yeh'kinya [quest] | 124.5 healing_power points (5.00 DPS) | yes | Darkspear Raider's Cloak (272063, -0.19 DPS) [vendor]; Hide of the Wild (18510, -1.12 DPS) [crafted] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 237.3 healing_power points (9.54 DPS) | yes | Devout Robe (16690, -0.30 DPS, sim-verified) [dungeon]; Tunic of The Five Thunders (227016, -0.82 DPS) [vendor]; Dreamwalker Armor (13123, -0.83 DPS) [world_drop] |
| wrist | Bracers of The Five Thunders (227009) | Mokvar [vendor] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Devout Bracers (16697, -0.72 DPS) [dungeon]; General's Mail Bracers (16576, -0.72 DPS) [pvp]; Bracers of Hope (22667, -4.32 DPS, sim-verified) [quest] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 180.7 healing_power points (7.27 DPS) | yes | Raider Handwraps (272097, +0.00 DPS, sim-verified) [vendor]; Devout Gloves (16692, -0.80 DPS) [dungeon]; Magister's Gloves (16684, -1.17 DPS) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 187.2 healing_power points (7.53 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Devout Belt (16696, -0.74 DPS) [dungeon]; Elunarian Belt (14465, -0.97 DPS) [world_drop] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Outrider's Mail Leggings (22676, -0.08 DPS) [rep]; Kilt of the Atal'ai Prophet (10807, -0.57 DPS) [dungeon]; Devout Skirt (16694, -5.27 DPS, sim-verified) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 225.3 healing_power points (9.06 DPS) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS, sim-verified) [vendor]; Boots of The Five Thunders (22096, -2.10 DPS) [quest]; Mooncloth Boots (15802, -2.21 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 151.5 healing_power points (6.09 DPS) | yes | Darkspear Signet (272068, -1.22 DPS) [vendor]; Channeler's Ring (272406, -1.29 DPS) [vendor]; Band of the Hierophant (13096, -1.52 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 130.5 healing_power points (5.25 DPS) | yes | Channeler's Ring (272406, -0.44 DPS) [vendor]; Band of the Hierophant (13096, -0.68 DPS) [world_drop]; Darkspear Signet (272068, -2.49 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -5.52 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -5.96 DPS) [quest]; Serenity Field (272439, -6.42 DPS) [vendor] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -0.66 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.09 DPS) [quest] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -1.92 DPS) [world_drop]; Charstone Dirk (17710, -2.08 DPS) [dungeon]; Hand of Edward the Odd (2243, -5.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Faded Hakkari Cloak; chest: Mooncloth Vest; wrist: Bracers of The Five Thunders; waist: Belt of Tiny Heads; legs: Leggings of The Five Thunders; feet: Greaves of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves; main_hand: Staff of Hale Magefire

No-known-source sample (15 of 1396, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (tauren, 5300000000000000-000000000000000000-5032503315513151)

Set DPS (verified): 622.9. Weights run: 1.6s. Verify run: 17.2s. 1396 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.213, intellect=1.026 ± 0.033, spirit=0.657 ± 0.025, mp5=1.860 ± 0.033, crit=0.528 ± 0.035 per rating point (14 rating = 1%, 7.395 per %), spell_haste=not significant (0.296 ± 0.449)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 100.7 healing_power points (32.70 DPS) | yes | Crown of The Five Thunders (227013, -10.50 DPS) [vendor]; Sanctified Leather Helm (22689, -10.81 DPS) [quest]; Mooncloth Circlet (14140, -12.93 DPS) [crafted] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 45.0 healing_power points (14.63 DPS) | yes | Drake Tooth Necklace (21531, -3.22 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -7.74 DPS) [dungeon]; Frostwolf Advisor's Pendant (19096, -8.88 DPS) [rep] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 76.7 healing_power points (24.92 DPS) | yes | Mooncloth Shoulders (14139, -7.77 DPS, sim-verified) [crafted]; Champion's Mail Epaulets (227166, -8.56 DPS) [vendor]; Warlord's Mail Epaulets (231665, -9.39 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 52.3 healing_power points (16.98 DPS) | yes | Drape of Recovery (272413, -2.73 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19526, -6.82 DPS) [rep] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | 70.7 healing_power points (22.95 DPS) | yes | Legionnaire's Mail Chestguard (227171, -2.53 DPS) [vendor]; Mooncloth Vest (14138, -3.79 DPS) [crafted]; Stormcloth Vest (10020, -5.14 DPS, sim-verified) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 37.2 healing_power points (12.07 DPS) | yes | Bracers of The Five Thunders (227009, +0.00 DPS, sim-verified) [vendor]; Nethergeld Cuffs (254061, -2.18 DPS) [crafted]; Mender's Leather Bracers (252543, -2.29 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 51.7 healing_power points (16.80 DPS) | yes | General's Mail Gloves (231666, -0.95 DPS) [vendor]; Grasp of The Five Thunders (227014, -1.49 DPS) [vendor]; Mooncloth Gloves (18409, -1.96 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 51.3 healing_power points (16.66 DPS) | yes | Sash of The Five Thunders (227010, +0.00 DPS, sim-verified) [vendor]; Mender's Leather Waistguard (252477, -2.91 DPS) [crafted]; Mender's Mail Belt (252591, -2.91 DPS) [crafted] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | 59.2 healing_power points (19.24 DPS) | yes | Legionnaire's Mail Pants (227167, +0.00 DPS) [vendor]; General's Mail Pants (231670, -0.10 DPS) [vendor]; Elderwild Pants (279254, -2.17 DPS, sim-verified) [crafted] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 67.6 healing_power points (21.95 DPS) | yes | Greaves of The Five Thunders (227015, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -6.71 DPS) [crafted]; Swarmtender's Footpads (275609, -7.98 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 53.5 healing_power points (17.39 DPS) | yes | Blessed Band of Light (272407, -5.61 DPS) [vendor]; Channeler's Ring (272406, -11.02 DPS) [vendor]; Sea Giant's Toe Ring (274746, -11.54 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 39.7 healing_power points (12.90 DPS) | yes | Blessed Band of Light (272407, -2.31 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -6.52 DPS) [vendor]; Sea Giant's Toe Ring (274746, -7.05 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -2.23 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.21 DPS) [world_drop]; Ankh of Life (1713, -6.50 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Ankh of Life (1713, -1.06 DPS) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Hand of Righteousness (7721, -10.32 DPS) [dungeon]; Spellshifter Rod (9527, -17.35 DPS) [quest] |
| off_hand | Evergreen Shield (279262) | Blacksmithing [crafted] | sim-verified (622.9 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, -0.04 DPS) [vendor]; Lei of the Lifegiver (19312, -21.78 DPS, sim-verified) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Leggings of The Five Thunders; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves; off_hand: Evergreen Shield

No-known-source sample (15 of 1396, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

