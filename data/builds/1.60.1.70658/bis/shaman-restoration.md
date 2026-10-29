# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5510000000000000)

Set DPS (verified): 26.8. Weights run: 6.1s. Verify run: 22.7s. 213 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=1.938 ± 0.009, spirit=1.454 ± 0.008, mp5=4.313 ± 0.013, crit=0.120 ± 0.006 per rating point (14 rating = 1%, 1.677 per %), spell_haste=0.228 ± 0.043

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 18.4 healing_power points (0.78 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.06 DPS) [crafted]; Trapper's Leather Hood (252505, -0.12 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.8 healing_power points (0.25 DPS) | yes | Scholarly Pendant (277203, -0.16 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.4 healing_power points (0.74 DPS) | yes | Slime-encrusted Pads (6461, -0.20 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.41 DPS) [crafted]; Forest Leather Mantle (4709, -0.49 DPS) [world_drop] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 11.9 healing_power points (0.50 DPS) | yes | Seer's Cape (6378, -0.22 DPS) [dungeon]; Spirit Cloak (4792, -0.26 DPS) [vendor]; Sanguine Cape (14376, -0.28 DPS, sim-verified) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.1 healing_power points (1.23 DPS) | yes | Filigreed Pristine Gown (253901, -0.13 DPS, sim-verified) [crafted]; Robe of the Moccasin (6465, -0.23 DPS) [dungeon]; Corsair's Overshirt (5202, -0.27 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.6 healing_power points (0.95 DPS) | yes | Drakewing Bands (12999, -0.59 DPS) [world_drop]; Bright Bracers (3647, -0.63 DPS) [world_drop]; Owl Bracers (4796, -0.66 DPS, sim-verified) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 17.8 healing_power points (0.75 DPS) | yes | Pristine Gloves (253913, -0.04 DPS) [crafted]; Magefist Gloves (12977, -0.09 DPS) [world_drop]; Bright Gloves (3066, -0.18 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.8 healing_power points (0.79 DPS) | yes | Novice Ardent's Sash (253887, -0.06 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.14 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 37.4 healing_power points (1.58 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.64 DPS) [world_drop]; Dreamer's Leggings (270016, -0.68 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.7 healing_power points (0.87 DPS) | yes | Black Whelp Slippers (252424, -0.21 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.25 DPS) [crafted]; Kimbra Boots (6191, -0.36 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.6 healing_power points (0.49 DPS) | yes | Deep Fathom Ring (6463, -0.18 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; Band of Purification (12996, -0.53 DPS, sim-verified) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (26.8 DPS) | yes | Deep Fathom Ring (6463, -0.06 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; Band of Purification (12996, -0.29 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 19.4 healing_power points (0.82 DPS) | yes | Staff of the Blessed Seer (2271, -0.20 DPS) [dungeon]; Staff of Westfall (2042, -0.23 DPS, sim-verified) [quest]; Taskmaster Axe (5194, -0.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 213, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (dwarf, 0000000000000000-000000000000000000-5532003300000000)

Set DPS (verified): 45.5. Weights run: 6.1s. Verify run: 21.9s. 388 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=2.113 ± 0.011, spirit=2.159 ± 0.011, mp5=5.739 ± 0.012, crit=0.170 ± 0.009 per rating point (14 rating = 1%, 2.378 per %), spell_haste=not significant (0.046 ± 0.017)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 48.9 healing_power points (2.25 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.59 DPS) [world_drop]; Dokebi Hat (14584, -0.69 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.1 healing_power points (0.79 DPS) | yes | Crystal Starfire Medallion (5003, -0.06 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.19 DPS) [dungeon]; Glowing Green Talisman (5002, -0.19 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 29.9 healing_power points (1.38 DPS) | yes | Nightsky Mantle (4718, +0.00 DPS, sim-verified) [world_drop]; Faerie Mantle (5820, -0.20 DPS) [quest]; Talbar Mantle (10657, -0.26 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 24.3 healing_power points (1.12 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Darkspear Raider's Cloak (272078, -0.04 DPS) [vendor]; Caretaker's Cape (19533, -0.12 DPS) [rep] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 43.6 healing_power points (2.01 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.02 DPS) [dungeon]; Beguiler Robes (7728, -0.05 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 27.8 healing_power points (1.28 DPS) | yes | Spidertank Oilrag (9448, -0.49 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.50 DPS) [world_drop]; Nightsky Wristbands (6407, -0.61 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 36.5 healing_power points (1.68 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.24 DPS, sim-verified) [dungeon]; Zodiac Gloves (7106, -0.59 DPS) [quest]; Silver-thread Gloves (6393, -0.70 DPS) [world_drop] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 44.6 healing_power points (2.06 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.89 DPS) [rep]; Highlander's Mail Girdle (20120, -0.89 DPS) [vendor]; Silver-lined Belt (13011, -1.26 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 50.9 healing_power points (2.34 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.44 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.53 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.4 healing_power points (2.00 DPS) | yes | Acidic Walkers (9454, -0.82 DPS) [dungeon]; Frothing Slippers (254003, -0.92 DPS) [crafted]; Soggy Boots (274747, -1.61 DPS, sim-verified) [vendor] |
| finger1 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 23.0 healing_power points (1.06 DPS) | yes | Electrocutioner Lagnut (9447, -0.09 DPS) [dungeon]; Runemender's Seal (273031, -0.09 DPS) [world]; Sea Giant's Toe Ring (274746, -0.23 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 21.5 healing_power points (0.99 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS, sim-verified) [dungeon]; Runemender's Seal (273031, -0.02 DPS) [world]; Sea Giant's Toe Ring (274746, -0.16 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (45.5 DPS) | yes | Hardened Root Staff (1317, -1.32 DPS) [quest]; Spine of the Devourer (273040, -1.47 DPS) [world]; Manual Crowd Pummeler (9449, -1.54 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 26.0 healing_power points (1.20 DPS) | yes | Orb of Mistmantle (13031, -0.38 DPS) [world_drop]; Defective Samophlange (274743, -0.40 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -0.41 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Darkspear Signet; finger2: The Queen's Jewel; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots

### Band 40 (dwarf, 0000000000000000-000000000000000000-5532003315400000)

Set DPS (verified): 78.7. Weights run: 6.9s. Verify run: 29.5s. 609 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.018, intellect=2.852 ± 0.017, spirit=2.774 ± 0.013, mp5=7.149 ± 0.022, crit=0.310 ± 0.018 per rating point (14 rating = 1%, 4.346 per %), spell_haste=not significant (0.075 ± 0.036)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 87.5 healing_power points (3.64 DPS) | yes | Whitemane's Chapeau (7720, -0.53 DPS) [dungeon]; Enduring Cap (3020, -0.93 DPS) [world_drop]; Imperial Leather Helm (6433, -0.95 DPS) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 39.4 healing_power points (1.64 DPS) | yes | Glowing Eye of Mordresh (10769, -0.15 DPS) [dungeon]; Amberglow Talisman (10824, -0.48 DPS) [quest]; Explorers' League Commendation (7746, -0.60 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 62.3 healing_power points (2.59 DPS) | yes | Inquisitor's Shawl (19507, -0.42 DPS) [dungeon]; Mistscape Mantle (4734, -0.71 DPS) [dungeon]; Crimson Silk Shoulders (7059, -0.83 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 42.5 healing_power points (1.77 DPS) | yes | Blackforge Cape (6424, -0.36 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.36 DPS) [pvp]; Caretaker's Cape (19532, -0.44 DPS) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathchill Armor (10764, +0.00 DPS) [dungeon]; Dreamweave Vest (10021, -0.65 DPS) [crafted]; Black Mageweave Robe (10001, -0.77 DPS) [crafted] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 39.1 healing_power points (1.63 DPS) | yes | Mindthrust Bracers (1974, -0.14 DPS) [dungeon]; Windchaser Cuffs (14429, -0.56 DPS) [world_drop]; Nightsky Wristbands (6407, -0.57 DPS) [world_drop] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 54.3 healing_power points (2.26 DPS) | yes | Gloves of Old (9395, -0.28 DPS) [world_drop]; Mender's Leather Gloves (252530, -0.32 DPS) [crafted]; Stormcloth Gloves (10011, -0.38 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 61.5 healing_power points (2.56 DPS) | yes | Mender's Leather Belt (252523, -0.36 DPS) [crafted]; Windchaser Cinch (14435, -0.57 DPS) [world_drop]; Sutarn's Ring (13105, -0.68 DPS) [world_drop] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | sim-verified (78.7 DPS) | yes | Misplaced Pantaloons (276201, +0.00 DPS) [vendor]; Stoneweaver Leggings (9407, -0.12 DPS) [dungeon]; Wisdom's Leather Leggings (252519, -0.48 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -0.51 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.9 healing_power points (1.78 DPS) | yes | Ogremind Ring (1993, -0.61 DPS) [world_drop]; Voodoo Band (1996, -0.61 DPS) [world]; The Queen's Jewel (13094, -0.62 DPS) [world_drop] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 33.7 healing_power points (1.40 DPS) | yes | Ogremind Ring (1993, -0.22 DPS) [world_drop]; Voodoo Band (1996, -0.22 DPS) [world]; The Queen's Jewel (13094, -0.24 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -1.48 DPS) [dungeon]; Hypnotic Blade (7714, -1.71 DPS) [dungeon]; Death Speaker Scepter (2816, -2.00 DPS) [dungeon] |
| off_hand | Battered Viking Shield (9403) | Uldaman: Olaf [dungeon] | sim-verified (78.7 DPS) | yes | Orb of Lorica (11262, +0.00 DPS) [quest]; Ravager's Shield (14777, +0.00 DPS) [world_drop]; Aurora Sphere (7610, -0.19 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Battered Viking Shield

No-known-source sample (15 of 609, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 50 (dwarf, 0000000000000000-000000000000000000-5532003315513131)

Set DPS (verified): 151.6. Weights run: 10.7s. Verify run: 46.2s. 777 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.090, intellect=2.458 ± 0.029, spirit=1.957 ± 0.019, mp5=5.200 ± 0.026, crit=0.350 ± 0.023 per rating point (14 rating = 1%, 4.893 per %), spell_haste=not significant (0.183 ± 0.076)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 86.0 healing_power points (6.72 DPS) | yes | Winged Helm (13112, -0.64 DPS, sim-verified) [world_drop]; Papal Fez (9431, -0.99 DPS) [dungeon]; Serpentskin Helm (8261, -1.35 DPS) [world_drop] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 45.9 healing_power points (3.59 DPS) | yes | Horizon Choker (13085, -0.29 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.04 DPS) [dungeon]; Necklace of Calisea (1714, -1.18 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 64.8 healing_power points (5.06 DPS) | yes | Mender's Leather Shoulders (252538, -0.39 DPS) [crafted]; Mender's Mail Shoulders (252569, -0.39 DPS) [crafted]; Nethergeld Shoulders (254049, -0.58 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 44.2 healing_power points (3.45 DPS) | yes | Caretaker's Cape (19531, -0.82 DPS) [rep]; Imperial Red Cloak (8248, -0.88 DPS, sim-verified) [world_drop]; Highborne Cloak (14450, -1.11 DPS) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 80.7 healing_power points (6.31 DPS) | yes | Stormcloth Vest (10020, +0.00 DPS, sim-verified) [crafted]; Feathered Breastplate (8349, -0.72 DPS) [crafted]; Deathchill Armor (10764, -1.27 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 47.9 healing_power points (3.75 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.07 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.98 DPS) [quest] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 86.3 healing_power points (6.74 DPS) | yes | Raider Handwraps (272098, -0.80 DPS) [vendor]; Mender's Leather Gauntlets (252551, -1.21 DPS) [crafted]; Mender's Mail Gauntlets (252587, -1.21 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 59.5 healing_power points (4.65 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.06 DPS) [dungeon]; Gilded Cord (254037, -0.48 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 85.5 healing_power points (6.68 DPS) | yes | Senior Designer's Pantaloons (11841, -0.63 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.73 DPS, sim-verified) [dungeon]; Dalewind Trousers (13008, -1.13 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 58.8 healing_power points (4.59 DPS) | yes | Mender's Leather Boots (252472, +0.00 DPS, sim-verified) [crafted]; Mender's Mail Sabatons (252579, -0.29 DPS) [crafted]; Mender's Leather Shoes (252533, -0.30 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 38.2 healing_power points (2.99 DPS) | yes | Welken Ring (5011, -0.96 DPS) [world_drop]; Choking Band (11868, -1.00 DPS) [quest]; Coldwater Ring (4550, -1.03 DPS) [quest] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 36.4 healing_power points (2.84 DPS) | yes | Welken Ring (5011, -0.81 DPS) [world_drop]; Choking Band (11868, -0.86 DPS) [quest]; Coldwater Ring (4550, -0.89 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (151.6 DPS) | yes | Ankh of Life (1713, -4.26 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.57 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.87 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (151.6 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.18 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (151.6 DPS) | yes | Charstone Dirk (17710, -6.68 DPS) [dungeon]; Spellshifter Rod (9527, -10.03 DPS) [quest]; Hammer of the Northern Wind (810, -26.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 777, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60 (dwarf, 5300000000000000-000000000000000000-5532003315513151)

Set DPS (verified): 231.8. Weights run: 11.0s. Verify run: 92.1s. 1487 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.092, intellect=4.037 ± 0.058, spirit=4.036 ± 0.043, mp5=9.490 ± 0.050, crit=0.657 ± 0.045 per rating point (14 rating = 1%, 9.195 per %), spell_haste=not significant (-0.162 ± 0.117)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 210.5 healing_power points (11.07 DPS) | yes | Living Crown (252561, -0.44 DPS) [crafted]; Crown of the Penitent (13216, -1.59 DPS) [quest]; Devout Crown (16693, -1.74 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 124.1 healing_power points (6.53 DPS) | yes | Darkmoon Necklace (19303, -2.26 DPS) [vendor]; Stormpike Sage's Pendant (19098, -2.41 DPS) [rep]; Drake Tooth Necklace (21531, -2.61 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 156.3 healing_power points (8.22 DPS) | yes | Devout Mantle (16695, -1.27 DPS) [dungeon]; Mantle of The Five Thunders (227011, -1.76 DPS) [vendor]; Drakesfire Epaulets (13133, -2.06 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | sim-verified (231.7 DPS) | yes | Faded Hakkari Cloak (20218, +0.00 DPS) [quest]; Hide of the Wild (18510, -0.34 DPS) [crafted]; Drape of Recovery (272413, -1.04 DPS) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 183.4 healing_power points (9.64 DPS) | yes | Devout Robe (16690, -0.32 DPS) [dungeon]; Tunic of The Five Thunders (227016, -0.64 DPS) [vendor]; Dreamwalker Armor (13123, -1.15 DPS) [world_drop] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (231.7 DPS) | yes | Bracers of The Five Thunders (227009, -0.64 DPS) [vendor]; Devout Bracers (16697, -1.43 DPS) [dungeon]; Magister's Bindings (16683, -1.80 DPS) [dungeon] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 141.1 healing_power points (7.42 DPS) | yes | Raider Handwraps (272097, -0.42 DPS) [vendor]; Devout Gloves (16692, -0.90 DPS) [dungeon]; Magister's Gloves (16684, -1.47 DPS) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 135.1 healing_power points (7.10 DPS) | yes | Devout Belt (16696, -0.37 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -0.59 DPS) [crafted]; Sash of The Five Thunders (227010, -0.61 DPS) [vendor] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devout Skirt (16694, +0.00 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.89 DPS) [dungeon]; Red Dragonscale Leggings (252603, -1.00 DPS) [crafted] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 172.2 healing_power points (9.05 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.80 DPS) [vendor]; Mooncloth Boots (15802, -2.06 DPS) [crafted]; Devout Sandals (16691, -2.74 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 118.8 healing_power points (6.25 DPS) | yes | Channeler's Ring (272406, -1.58 DPS) [vendor]; Darkspear Signet (272068, -1.75 DPS) [vendor]; Band of the Hierophant (13096, -1.79 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 100.3 healing_power points (5.28 DPS) | yes | Channeler's Ring (272406, -0.61 DPS) [vendor]; Darkspear Signet (272068, -0.78 DPS) [vendor]; Band of the Hierophant (13096, -0.82 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -4.94 DPS) [world_drop]; Serenity Field (272439, -5.28 DPS) [vendor]; Evonice's Landin' Pilla (18951, -5.36 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.45 DPS) [world_drop]; Serenity Field (272439, -0.79 DPS) [vendor]; Evonice's Landin' Pilla (18951, -0.87 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -5.98 DPS) [world_drop]; Hand of Edward the Odd (2243, -6.34 DPS) [world_drop]; Charstone Dirk (17710, -6.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Darkspear Raider's Cloak; chest: Mooncloth Vest; wrist: Bracers of Hope; waist: Belt of Tiny Heads; legs: Leggings of The Five Thunders; feet: Greaves of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1487, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 5300000000000000-000000000000000000-5532003315513151)

Set DPS (verified): 631.3. Weights run: 8.1s. Verify run: 52.2s. 1487 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.224, intellect=1.823 ± 0.038, spirit=1.117 ± 0.032, mp5=3.283 ± 0.040, crit=0.800 ± 0.041 per rating point (14 rating = 1%, 11.193 per %), spell_haste=not significant (1.668 ± 0.540)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 121.5 healing_power points (22.06 DPS) | yes | Crown of The Five Thunders (227013, +0.00 DPS, sim-verified) [vendor]; Sanctified Leather Helm (22689, -6.07 DPS) [quest]; Gnomish Turban of Psychic Might (21517, -6.43 DPS) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 61.2 healing_power points (11.12 DPS) | yes | Drake Tooth Necklace (21531, -4.33 DPS, sim-verified) [quest]; Stormpike Sage's Pendant (19098, -5.42 DPS) [rep]; Darkmoon Necklace (19303, -5.55 DPS) [vendor] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 89.9 healing_power points (16.32 DPS) | yes | Mantle of The Five Thunders (227011, -5.59 DPS) [vendor]; Devout Mantle (16695, -6.68 DPS, sim-verified) [dungeon]; Eternity Pauldrons (279255, -7.22 DPS) [crafted] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 60.2 healing_power points (10.94 DPS) | yes | Drape of Recovery (272413, -2.76 DPS, sim-verified) [vendor]; Darkspear Raider's Cloak (272063, -4.42 DPS) [vendor]; Caretaker's Cape (19530, -4.59 DPS) [rep] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | 94.1 healing_power points (17.09 DPS) | yes | Devout Robe (16690, -2.47 DPS) [dungeon]; Mooncloth Vest (14138, -2.97 DPS, sim-verified) [crafted]; Mooncloth Robe (18486, -4.10 DPS) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 51.5 healing_power points (9.35 DPS) | yes | Bracers of The Five Thunders (227009, +0.00 DPS, sim-verified) [vendor]; Mender's Leather Bracers (252543, -2.36 DPS) [crafted]; Mender's Mail Bracers (252573, -2.36 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 73.2 healing_power points (13.30 DPS) | yes | Grasp of The Five Thunders (227014, +0.00 DPS, sim-verified) [vendor]; Mooncloth Gloves (18409, -2.01 DPS) [crafted]; Mender's Leather Gauntlets (252551, -2.86 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 66.6 healing_power points (12.09 DPS) | yes | Sash of The Five Thunders (227010, +0.00 DPS, sim-verified) [vendor]; Devout Belt (16696, -1.65 DPS) [dungeon]; Belt of Tiny Heads (20217, -2.29 DPS) [quest] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | 80.3 healing_power points (14.58 DPS) | yes | Red Dragonscale Leggings (252603, -0.35 DPS) [crafted]; Devout Skirt (16694, -0.59 DPS) [dungeon]; Elderwild Pants (279254, -2.44 DPS) [crafted] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 85.2 healing_power points (15.47 DPS) | yes | Greaves of The Five Thunders (227015, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.84 DPS) [crafted]; Devout Sandals (16691, -6.71 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 67.8 healing_power points (12.32 DPS) | yes | Blessed Band of Light (272407, -5.06 DPS) [vendor]; Channeler's Ring (272406, -6.06 DPS) [vendor]; Band of the Hierophant (13096, -6.78 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 53.4 healing_power points (9.69 DPS) | yes | Blessed Band of Light (272407, -2.93 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.43 DPS) [vendor]; Band of the Hierophant (13096, -4.15 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (631.3 DPS) | yes | Serenity Field (272439, -4.29 DPS, sim-verified) [vendor]; Uther's Strength (11302, -5.68 DPS) [world_drop]; Ankh of Life (1713, -6.51 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (631.3 DPS) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Uther's Strength (11302, -0.31 DPS) [world_drop]; Ankh of Life (1713, -1.14 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (631.3 DPS) | yes | Charstone Dirk (17710, -12.77 DPS) [dungeon]; Hand of Edward the Odd (2243, -17.85 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -19.97 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Leggings of The Five Thunders; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1487, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

## Horde

### Band 20 (tauren, 0000000000000000-000000000000000000-5510000000000000)

Set DPS (verified): 26.8. Weights run: 6.1s. Verify run: 20.0s. 193 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=1.938 ± 0.009, spirit=1.454 ± 0.008, mp5=4.313 ± 0.013, crit=0.120 ± 0.006 per rating point (14 rating = 1%, 1.677 per %), spell_haste=0.228 ± 0.043

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (26.8 DPS) | yes | Shadow Goggles (4373, +0.00 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.04 DPS) [crafted]; Trapper's Leather Hood (252505, -0.11 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.4 healing_power points (0.18 DPS) | yes | Scholarly Pendant (277203, -0.06 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.4 healing_power points (0.74 DPS) | yes | Slime-encrusted Pads (6461, -0.19 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.41 DPS) [crafted]; Forest Leather Mantle (4709, -0.49 DPS) [world_drop] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 11.9 healing_power points (0.50 DPS) | yes | Sanguine Cape (14376, -0.18 DPS) [world_drop]; Seer's Cape (6378, -0.22 DPS) [dungeon]; Spirit Cloak (4792, -0.26 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.1 healing_power points (1.23 DPS) | yes | Filigreed Pristine Gown (253901, -0.08 DPS) [crafted]; Robe of the Moccasin (6465, -0.23 DPS) [dungeon]; Corsair's Overshirt (5202, -0.27 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.6 healing_power points (0.95 DPS) | yes | Owl Bracers (4796, -0.55 DPS) [vendor]; Featherbead Bracers (15452, -0.55 DPS) [quest]; Drakewing Bands (12999, -0.59 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 20.8 healing_power points (0.88 DPS) | yes | Wisdom's Leather Gloves (252499, -0.13 DPS) [crafted]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Magefist Gloves (12977, -0.22 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.8 healing_power points (0.79 DPS) | yes | Novice Ardent's Sash (253887, -0.11 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.14 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 37.4 healing_power points (1.58 DPS) | yes | Filigreed Pristine Leggings (253937, -0.08 DPS) [crafted]; Darkweave Breeches (12987, -0.64 DPS) [world_drop]; Ghastly Trousers (15449, -0.81 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.7 healing_power points (0.87 DPS) | yes | Black Whelp Slippers (252424, -0.17 DPS) [crafted]; Pristine Boots (253889, -0.25 DPS) [crafted]; Bluegill Sandals (1560, -0.46 DPS) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.6 healing_power points (0.49 DPS) | yes | Band of Purification (12996, -0.12 DPS) [world_drop]; Advisor's Ring (20426, -0.13 DPS) [rep]; Deep Fathom Ring (6463, -0.18 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 9.7 healing_power points (0.41 DPS) | yes | Band of Purification (12996, -0.04 DPS) [world_drop]; Advisor's Ring (20426, -0.04 DPS) [rep]; Deep Fathom Ring (6463, -0.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 23.6 healing_power points (1.00 DPS) | yes | Advisor's Gnarled Staff (20425, -0.12 DPS) [pvp]; Twisted Chanter's Staff (890, -0.18 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.18 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 206376 Bulwark Icon; 206383 Lithic Icon; 206384 Astral Icon; 206385 Clastic Icon; 206389 Brimstone Icon; 206390 Ichthyan Icon; 206391 Syzygic Icon; 241089 Scarlet Dagger

### Band 30 (tauren, 0000000000000000-000000000000000000-5532003300000000)

Set DPS (verified): 45.4. Weights run: 6.1s. Verify run: 20.7s. 368 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=2.113 ± 0.011, spirit=2.159 ± 0.011, mp5=5.739 ± 0.012, crit=0.170 ± 0.009 per rating point (14 rating = 1%, 2.378 per %), spell_haste=not significant (0.046 ± 0.017)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 48.9 healing_power points (2.25 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.59 DPS) [world_drop]; Dokebi Hat (14584, -0.69 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.1 healing_power points (0.79 DPS) | yes | Crystal Starfire Medallion (5003, -0.19 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.19 DPS) [dungeon]; Glowing Green Talisman (5002, -0.19 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 28.4 healing_power points (1.31 DPS) | yes | Nightsky Mantle (4718, -0.19 DPS, sim-verified) [world_drop]; Talbar Mantle (10657, -0.20 DPS) [quest]; Death Speaker Mantle (6685, -0.24 DPS) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 24.3 healing_power points (1.12 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Battle Healer's Cloak (19529, -0.12 DPS) [rep]; Amy's Blanket (13005, -0.42 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 43.6 healing_power points (2.01 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.02 DPS) [dungeon]; Beguiler Robes (7728, -0.05 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 27.8 healing_power points (1.28 DPS) | yes | Spidertank Oilrag (9448, -0.49 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.50 DPS) [world_drop]; Nightsky Wristbands (6407, -0.64 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 36.5 healing_power points (1.68 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.30 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.50 DPS) [quest]; Tattered Mittens (270030, -0.62 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 44.6 healing_power points (2.06 DPS) | yes | Lilac Sash (6780, -0.88 DPS) [quest]; Highlander's Mail Girdle (20120, -0.89 DPS) [vendor]; Silver-lined Belt (13011, -1.36 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 50.9 healing_power points (2.34 DPS) | yes | Pristine Leggings (253987, -0.35 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.44 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.53 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.4 healing_power points (2.00 DPS) | yes | Acidic Walkers (9454, -0.82 DPS) [dungeon]; Frothing Slippers (254003, -0.92 DPS) [crafted]; Soggy Boots (274747, -1.74 DPS, sim-verified) [vendor] |
| finger1 | Darkspear Signet (272071) | Creeg Bothunk [vendor] | 23.0 healing_power points (1.06 DPS) | yes | Electrocutioner Lagnut (9447, -0.09 DPS) [dungeon]; Runemender's Seal (273031, -0.09 DPS) [world]; Sea Giant's Toe Ring (274746, -0.23 DPS) [vendor] |
| finger2 | The Queen's Jewel (13094) | World drop [world_drop] | 21.5 healing_power points (0.99 DPS) | yes | Electrocutioner Lagnut (9447, +0.00 DPS, sim-verified) [dungeon]; Runemender's Seal (273031, -0.02 DPS) [world]; Sea Giant's Toe Ring (274746, -0.16 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (45.4 DPS) | yes | Advisor's Gnarled Staff (19569, -1.39 DPS) [pvp]; Spine of the Devourer (273040, -1.47 DPS) [world]; Manual Crowd Pummeler (9449, -1.55 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 26.0 healing_power points (1.20 DPS) | yes | Orb of Mistmantle (13031, -0.38 DPS) [world_drop]; Satyr's Rod (15962, -0.41 DPS) [world_drop]; Defective Samophlange (274743, -0.59 DPS, sim-verified) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Darkspear Signet; finger2: The Queen's Jewel; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 368, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-000000000000000000-5532003315400000)

Set DPS (verified): 78.2. Weights run: 6.9s. Verify run: 27.7s. 570 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.018, intellect=2.852 ± 0.017, spirit=2.774 ± 0.013, mp5=7.149 ± 0.022, crit=0.310 ± 0.018 per rating point (14 rating = 1%, 4.346 per %), spell_haste=not significant (0.075 ± 0.036)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 87.5 healing_power points (3.64 DPS) | yes | Whitemane's Chapeau (7720, -0.82 DPS, sim-verified) [dungeon]; Enduring Cap (3020, -0.93 DPS) [world_drop]; Imperial Leather Helm (6433, -0.95 DPS) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 39.4 healing_power points (1.64 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.48 DPS) [quest]; Kaleidoscope Chain (13084, -0.70 DPS) [world_drop] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 62.3 healing_power points (2.59 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS, sim-verified) [dungeon]; Mistscape Mantle (4734, -0.71 DPS) [dungeon]; Crimson Silk Shoulders (7059, -0.83 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 42.5 healing_power points (1.77 DPS) | yes | Blackforge Cape (6424, -0.36 DPS) [dungeon]; Battle Healer's Cloak (19528, -0.44 DPS) [rep]; Cloak of Blight (6832, -0.81 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.65 DPS) [crafted]; Black Mageweave Robe (10001, -0.77 DPS) [crafted]; Deathchill Armor (10764, -1.66 DPS, sim-verified) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 39.1 healing_power points (1.63 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Windtalker's Wristguards (19584, -0.34 DPS) [pvp]; Dryad's Wrist Bindings (19597, -0.34 DPS) [pvp] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 54.3 healing_power points (2.26 DPS) | yes | Mender's Leather Gloves (252530, -0.32 DPS) [crafted]; Stormcloth Gloves (10011, -0.38 DPS) [crafted]; Gloves of Old (9395, -0.71 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 61.5 healing_power points (2.56 DPS) | yes | Mender's Leather Belt (252523, -0.43 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.57 DPS) [world_drop]; Sutarn's Ring (13105, -0.68 DPS) [world_drop] |
| legs | Misplaced Pantaloons (276201) | Friz Frazzlespark [vendor] | 73.1 healing_power points (3.04 DPS) | yes | Warchief Kilt (7760, +0.00 DPS, sim-verified) [dungeon]; Stoneweaver Leggings (9407, -0.17 DPS) [dungeon]; Wisdom's Leather Leggings (252519, -0.53 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Gilded Slippers (254001, -0.51 DPS) [crafted]; Furen's Boots (13100, -1.21 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 42.9 healing_power points (1.78 DPS) | yes | Ogremind Ring (1993, -0.61 DPS) [world_drop]; Voodoo Band (1996, -0.61 DPS) [world]; The Queen's Jewel (13094, -0.62 DPS) [world_drop] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 33.7 healing_power points (1.40 DPS) | yes | Ogremind Ring (1993, -0.22 DPS) [world_drop]; Voodoo Band (1996, -0.22 DPS) [world]; The Queen's Jewel (13094, -0.24 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -1.48 DPS) [dungeon]; Hypnotic Blade (7714, -1.71 DPS) [dungeon]; Gut Ripper (2164, -8.78 DPS, sim-verified) [world_drop] |
| off_hand | Ravager's Shield (14777) | World drop [world_drop] | 36.6 healing_power points (1.52 DPS) | yes | Battered Viking Shield (9403, -0.04 DPS) [dungeon]; Forcestone Buckler (17508, -0.10 DPS) [quest]; Aurora Sphere (7610, -0.23 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Sheepshear Mantle; back: Darkspear Raider's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Misplaced Pantaloons; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Welken Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Ravager's Shield

No-known-source sample (15 of 570, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (tauren, 0000000000000000-000000000000000000-5532003315513131)

Set DPS (verified): 151.1. Weights run: 10.7s. Verify run: 44.8s. 720 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.090, intellect=2.458 ± 0.029, spirit=1.957 ± 0.019, mp5=5.200 ± 0.026, crit=0.350 ± 0.023 per rating point (14 rating = 1%, 4.893 per %), spell_haste=not significant (0.183 ± 0.076)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 86.0 healing_power points (6.72 DPS) | yes | Winged Helm (13112, -0.78 DPS, sim-verified) [world_drop]; Papal Fez (9431, -0.99 DPS) [dungeon]; Serpentskin Helm (8261, -1.35 DPS) [world_drop] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 45.9 healing_power points (3.59 DPS) | yes | Horizon Choker (13085, -0.29 DPS, sim-verified) [world_drop]; Glowing Eye of Mordresh (10769, -1.04 DPS) [dungeon]; Necklace of Calisea (1714, -1.18 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 64.8 healing_power points (5.06 DPS) | yes | Mender's Leather Shoulders (252538, -0.20 DPS, sim-verified) [crafted]; Mender's Mail Shoulders (252569, -0.39 DPS) [crafted]; Nethergeld Shoulders (254049, -0.58 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 44.2 healing_power points (3.45 DPS) | yes | Battle Healer's Cloak (19527, -0.82 DPS) [rep]; Imperial Red Cloak (8248, -0.91 DPS, sim-verified) [world_drop]; Highborne Cloak (14450, -1.11 DPS) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 80.7 healing_power points (6.31 DPS) | yes | Stormcloth Vest (10020, +0.00 DPS, sim-verified) [crafted]; Feathered Breastplate (8349, -0.72 DPS) [crafted]; Deathchill Armor (10764, -1.27 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 47.9 healing_power points (3.75 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.07 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.98 DPS) [quest] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 86.3 healing_power points (6.74 DPS) | yes | Raider Handwraps (272098, -0.80 DPS) [vendor]; Mender's Leather Gauntlets (252551, -1.21 DPS) [crafted]; Mender's Mail Gauntlets (252587, -1.21 DPS) [crafted] |
| waist | Mender's Leather Waistguard (252477) (or Mender's Mail Belt (252591)) | Leatherworking [crafted] | 59.5 healing_power points (4.65 DPS) | yes | Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.06 DPS) [dungeon]; Gilded Cord (254037, -0.48 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 85.5 healing_power points (6.68 DPS) | yes | Senior Designer's Pantaloons (11841, -0.63 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.90 DPS, sim-verified) [dungeon]; Dalewind Trousers (13008, -1.13 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 58.8 healing_power points (4.59 DPS) | yes | Mender's Leather Boots (252472, +0.00 DPS, sim-verified) [crafted]; Mender's Mail Sabatons (252579, -0.29 DPS) [crafted]; Mender's Leather Shoes (252533, -0.30 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 38.2 healing_power points (2.99 DPS) | yes | Welken Ring (5011, -0.96 DPS) [world_drop]; Coldwater Ring (4550, -1.03 DPS) [quest]; Voodoo Band (1996, -1.18 DPS) [world] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 36.4 healing_power points (2.84 DPS) | yes | Welken Ring (5011, -0.81 DPS) [world_drop]; Coldwater Ring (4550, -0.89 DPS) [quest]; Voodoo Band (1996, -1.04 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (151.1 DPS) | yes | Ankh of Life (1713, -4.26 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.57 DPS) [quest]; Alchemist's Stone (13503, -5.48 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (151.1 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.36 DPS, sim-verified) [world_drop]; Alchemist's Stone (13503, -0.80 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (151.1 DPS) | yes | Charstone Dirk (17710, -6.68 DPS) [dungeon]; Spellshifter Rod (9527, -10.03 DPS) [quest]; Hammer of the Northern Wind (810, -26.88 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Grasp of The Five Thunders; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 720, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (tauren, 5300000000000000-000000000000000000-5532003315513151)

Set DPS (verified): 231.0. Weights run: 11.0s. Verify run: 89.8s. 1410 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.092, intellect=4.037 ± 0.058, spirit=4.036 ± 0.043, mp5=9.490 ± 0.050, crit=0.657 ± 0.045 per rating point (14 rating = 1%, 9.195 per %), spell_haste=not significant (-0.162 ± 0.117)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gnomish Turban of Psychic Might (21517) | The Only Prescription [quest] | 210.5 healing_power points (11.07 DPS) | yes | Living Crown (252561, -0.44 DPS) [crafted]; Crown of the Penitent (13216, -1.59 DPS) [quest]; Devout Crown (16693, -1.74 DPS) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 124.1 healing_power points (6.53 DPS) | yes | Darkmoon Necklace (19303, -2.26 DPS) [vendor]; Frostwolf Advisor's Pendant (19096, -2.41 DPS) [rep]; Vision of Voodress (19608, -2.49 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 156.3 healing_power points (8.22 DPS) | yes | Devout Mantle (16695, -1.27 DPS) [dungeon]; Mantle of The Five Thunders (227011, -1.76 DPS) [vendor]; Drakesfire Epaulets (13133, -2.06 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | sim-verified (231.0 DPS) | yes | Faded Hakkari Cloak (20218, +0.00 DPS) [quest]; Hide of the Wild (18510, -0.34 DPS) [crafted]; Drape of Recovery (272413, -1.04 DPS) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 183.4 healing_power points (9.64 DPS) | yes | Devout Robe (16690, -0.32 DPS) [dungeon]; Tunic of The Five Thunders (227016, -0.64 DPS) [vendor]; Dreamwalker Armor (13123, -1.15 DPS) [world_drop] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (231.0 DPS) | yes | Bracers of The Five Thunders (227009, -0.64 DPS) [vendor]; Devout Bracers (16697, -1.43 DPS) [dungeon]; General's Mail Bracers (16576, -1.58 DPS) [pvp] |
| hands | Grasp of The Five Thunders (227014) | Mokvar [vendor] | 141.1 healing_power points (7.42 DPS) | yes | Raider Handwraps (272097, -0.42 DPS) [vendor]; Devout Gloves (16692, -0.90 DPS) [dungeon]; Magister's Gloves (16684, -1.47 DPS) [dungeon] |
| waist | Belt of Tiny Heads (20217) | A Collection of Heads [quest] | 135.1 healing_power points (7.10 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Devout Belt (16696, -0.37 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -0.59 DPS) [crafted] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devout Skirt (16694, +0.00 DPS) [dungeon]; Outrider's Mail Leggings (22676, -0.70 DPS) [rep]; Kilt of the Atal'ai Prophet (10807, -0.89 DPS) [dungeon] |
| feet | Greaves of The Five Thunders (227015) | Mokvar [vendor] | 172.2 healing_power points (9.05 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.80 DPS) [vendor]; Mooncloth Boots (15802, -2.06 DPS) [crafted]; Boots of The Five Thunders (22096, -2.39 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 118.8 healing_power points (6.25 DPS) | yes | Channeler's Ring (272406, -1.58 DPS) [vendor]; Darkspear Signet (272068, -1.75 DPS) [vendor]; Band of the Hierophant (13096, -1.79 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 100.3 healing_power points (5.28 DPS) | yes | Channeler's Ring (272406, -0.61 DPS) [vendor]; Darkspear Signet (272068, -0.78 DPS) [vendor]; Band of the Hierophant (13096, -0.82 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -4.94 DPS) [world_drop]; Serenity Field (272439, -5.28 DPS) [vendor]; Evonice's Landin' Pilla (18951, -5.36 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.45 DPS) [world_drop]; Serenity Field (272439, -0.79 DPS) [vendor]; Evonice's Landin' Pilla (18951, -0.87 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -5.98 DPS) [world_drop]; Hand of Edward the Odd (2243, -6.34 DPS) [world_drop]; Charstone Dirk (17710, -6.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Gnomish Turban of Psychic Might; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Darkspear Raider's Cloak; chest: Mooncloth Vest; wrist: Bracers of Hope; waist: Belt of Tiny Heads; legs: Leggings of The Five Thunders; feet: Greaves of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (tauren, 5300000000000000-000000000000000000-5532003315513151)

Set DPS (verified): 630.3. Weights run: 8.1s. Verify run: 51.0s. 1410 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.224, intellect=1.823 ± 0.038, spirit=1.117 ± 0.032, mp5=3.283 ± 0.040, crit=0.800 ± 0.041 per rating point (14 rating = 1%, 11.193 per %), spell_haste=not significant (1.668 ± 0.540)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 121.5 healing_power points (22.06 DPS) | yes | Crown of The Five Thunders (227013, +0.00 DPS, sim-verified) [vendor]; Sanctified Leather Helm (22689, -6.07 DPS) [quest]; Gnomish Turban of Psychic Might (21517, -6.43 DPS) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 61.2 healing_power points (11.12 DPS) | yes | Drake Tooth Necklace (21531, -4.38 DPS, sim-verified) [quest]; Frostwolf Advisor's Pendant (19096, -5.42 DPS) [rep]; Darkmoon Necklace (19303, -5.55 DPS) [vendor] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 89.9 healing_power points (16.32 DPS) | yes | Champion's Mail Epaulets (227166, -4.52 DPS) [vendor]; Warlord's Mail Epaulets (231665, -4.63 DPS) [vendor]; Devout Mantle (16695, -5.96 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 60.2 healing_power points (10.94 DPS) | yes | Drape of Recovery (272413, -2.99 DPS, sim-verified) [vendor]; Darkspear Raider's Cloak (272063, -4.42 DPS) [vendor]; Battle Healer's Cloak (19526, -4.59 DPS) [rep] |
| chest | Tunic of The Five Thunders (227016) | Mokvar [vendor] | 94.1 healing_power points (17.09 DPS) | yes | Legionnaire's Mail Chestguard (227171, -2.38 DPS) [vendor]; Devout Robe (16690, -2.47 DPS) [dungeon]; Mooncloth Vest (14138, -3.15 DPS, sim-verified) [crafted] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 51.5 healing_power points (9.35 DPS) | yes | Bracers of The Five Thunders (227009, +0.00 DPS, sim-verified) [vendor]; Mender's Leather Bracers (252543, -2.36 DPS) [crafted]; Mender's Mail Bracers (252573, -2.36 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 73.2 healing_power points (13.30 DPS) | yes | Grasp of The Five Thunders (227014, +0.00 DPS, sim-verified) [vendor]; General's Mail Gloves (231666, -1.58 DPS) [vendor]; Mooncloth Gloves (18409, -2.01 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 66.6 healing_power points (12.09 DPS) | yes | Sash of The Five Thunders (227010, +0.00 DPS, sim-verified) [vendor]; Devout Belt (16696, -1.65 DPS) [dungeon]; General's Mail Waistband (16575, -1.79 DPS) [pvp] |
| legs | Leggings of The Five Thunders (227012) | Mokvar [vendor] | 80.3 healing_power points (14.58 DPS) | yes | Legionnaire's Mail Pants (227167, +0.00 DPS) [vendor]; General's Mail Pants (231670, -0.15 DPS) [vendor]; Red Dragonscale Leggings (252603, -0.35 DPS) [crafted] |
| feet | Incandescent Mooncloth Boots (227862) | Meilosh [vendor] | 85.2 healing_power points (15.47 DPS) | yes | Greaves of The Five Thunders (227015, +0.00 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.84 DPS) [crafted]; General's Mail Boots (231667, -5.45 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 67.8 healing_power points (12.32 DPS) | yes | Blessed Band of Light (272407, -5.06 DPS) [vendor]; Channeler's Ring (272406, -6.06 DPS) [vendor]; Band of the Hierophant (13096, -6.78 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 53.4 healing_power points (9.69 DPS) | yes | Blessed Band of Light (272407, -3.38 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.43 DPS) [vendor]; Band of the Hierophant (13096, -4.15 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (630.3 DPS) | yes | Serenity Field (272439, -1.32 DPS) [vendor]; Uther's Strength (11302, -5.68 DPS) [world_drop]; Ankh of Life (1713, -6.51 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (630.3 DPS) | yes | Serenity Field (272439, +0.00 DPS, sim-verified) [vendor]; Uther's Strength (11302, -0.31 DPS) [world_drop]; Ankh of Life (1713, -1.14 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (630.3 DPS) | yes | Charstone Dirk (17710, -12.77 DPS) [dungeon]; Hand of Edward the Odd (2243, -18.18 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -19.97 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Tunic of The Five Thunders; wrist: Bracers of Hope; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Leggings of The Five Thunders; feet: Incandescent Mooncloth Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

