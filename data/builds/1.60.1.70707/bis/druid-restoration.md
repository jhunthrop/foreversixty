# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5051000000000000)

Set DPS (verified): 55.1. Weights run: 4.5s. Verify run: 15.8s. 230 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.279 ± 0.009, spirit=1.547 ± 0.011, mp5=3.188 ± 0.015, crit=0.074 ± 0.003 per rating point (14 rating = 1%, 1.042 per %), spell_haste=-3.937 ± 0.154

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.95 DPS) | yes | Wisdom's Leather Hood (252507, -0.14 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.25 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.44 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 3.1 healing_power points (0.34 DPS) | yes | Tarnished Locket (279870, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 healing_power points (1.25 DPS) | yes | Serpent's Shoulders (5404, +0.00 DPS, sim-verified) [dungeon]; Slime-encrusted Pads (6461, -0.21 DPS) [dungeon]; Forest Leather Mantle (4709, -0.58 DPS) [world_drop] |
| back | Spirit Cloak (4792) (or Sylvan Cloak (4793), Heavy Woolen Cloak (4311)) | Jennabink Powerseam [vendor] | 6.2 healing_power points (0.67 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Sylvan Cloak (4793, +0.00 DPS) [vendor]; Seer's Cape (6378, -0.06 DPS) [dungeon] |
| chest | Robe of the Moccasin (6465) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (55.1 DPS) | yes | Filigreed Pristine Gown (253901, -0.05 DPS) [crafted]; Corsair's Overshirt (5202, -0.12 DPS) [dungeon]; Wisdom's Leather Armor (252493, -1.36 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.0 healing_power points (1.73 DPS) | yes | Drakewing Bands (12999, -0.98 DPS, sim-verified) [world_drop]; Owl Bracers (4796, -1.04 DPS) [vendor]; Repurposed Hair Band (281256, -1.12 DPS) [quest] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 15.1 healing_power points (1.64 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.27 DPS) [world_drop]; Bright Gloves (3066, -0.41 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.1 healing_power points (1.74 DPS) | yes | Wisdom's Leather Belt (252433, -0.32 DPS) [crafted]; Keller's Girdle (2911, -0.64 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.85 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.67 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.69 DPS) [world_drop]; Stormrider's Leather Pants (252502, -2.17 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.88 DPS) | yes | Black Whelp Slippers (252424, -0.12 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.49 DPS) [crafted]; Kimbra Boots (6191, -0.83 DPS) [quest] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (1.01 DPS) | yes | Volcanic Rock Ring (12053, -0.59 DPS) [world_drop]; Demon Band (12054, -0.67 DPS) [world_drop]; The 1 Ring (8350, -0.70 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.7 healing_power points (0.83 DPS) | yes | Volcanic Rock Ring (12053, -0.42 DPS) [world_drop]; Demon Band (12054, -0.50 DPS) [world_drop]; The 1 Ring (8350, -0.52 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 15.7 healing_power points (1.70 DPS) | yes | Twisted Chanter's Staff (890, -0.31 DPS) [world_drop]; Evocator's Blade (2567, -0.58 DPS) [dungeon]; Staff of the Blessed Seer (2271, -2.90 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Spirit Cloak; chest: Robe of the Moccasin; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Lavishly Jeweled Ring; main_hand: Staff of Westfall

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5053005110010000)

Set DPS (verified): 116.4. Weights run: 6.0s. Verify run: 21.0s. 414 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.614 ± 0.014, spirit=3.959 ± 0.021, mp5=4.933 ± 0.028, crit=0.126 ± 0.004 per rating point (14 rating = 1%, 1.765 per %), spell_haste=not significant (-0.210 ± 0.293)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 56.8 healing_power points (6.03 DPS) | yes | Enduring Cap (3020, -0.86 DPS, sim-verified) [world_drop]; Stormrider's Leather Helm (252514, -1.82 DPS) [crafted]; Nightsky Cowl (4039, -1.87 DPS) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 23.8 healing_power points (2.52 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.58 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 39.0 healing_power points (4.15 DPS) | yes | Faerie Mantle (5820, -0.48 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.67 DPS) [world_drop]; Insignia Mantle (4721, -1.20 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 38.7 healing_power points (4.11 DPS) | yes | Repairman's Cape (9605, -0.85 DPS) [quest]; Amy's Blanket (13005, -1.16 DPS) [world_drop]; Nightsky Cloak (4719, -1.59 DPS) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Beguiler Robes (7728, -0.04 DPS) [dungeon]; Silver-thread Armor (7110, -0.15 DPS) [world_drop]; Pressed Felt Robe (1997, -1.61 DPS, sim-verified) [world] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Wristbands (6407, -0.14 DPS) [world_drop]; Dokebi Bracers (14580, -0.23 DPS) [world_drop]; Drakewing Bands (12999, -1.34 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 55.6 healing_power points (5.91 DPS) | yes | Hotshot Pilot's Gloves (9491, -2.43 DPS) [dungeon]; Stonecloth Gloves (14411, -2.54 DPS) [world_drop]; Zodiac Gloves (7106, -2.73 DPS, sim-verified) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 52.4 healing_power points (5.57 DPS) | yes | Silver-lined Belt (13011, -2.28 DPS, sim-verified) [world_drop]; Dokebi Cord (14578, -2.36 DPS) [world_drop]; Resilient Cord (14406, -2.86 DPS) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 60.0 healing_power points (6.38 DPS) | yes | Stormrider's Leather Kilt (252518, -0.55 DPS) [crafted]; Pristine Leggings (253987, -0.74 DPS) [crafted]; Earthen Leggings (253999, -1.52 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 47.1 healing_power points (5.01 DPS) | yes | Resilient Boots (14399, -1.64 DPS) [world_drop]; Stonecloth Boots (14408, -1.64 DPS) [world_drop]; Boots of the Enchanter (4325, -3.68 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 36.7 healing_power points (3.90 DPS) | yes | The Queen's Jewel (13094, -0.19 DPS) [world_drop]; Ring of Calm (6790, -0.96 DPS) [quest]; Band of Thorns (5007, -1.38 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 36.7 healing_power points (3.90 DPS) | yes | The Queen's Jewel (13094, -0.19 DPS) [world_drop]; Ring of Calm (6790, -0.96 DPS) [quest]; Band of Thorns (5007, -1.38 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Friar (3415, -1.66 DPS) [dungeon]; Odo's Ley Staff (6318, -2.18 DPS) [dungeon]; Gnarled Ash Staff (791, -3.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5053005153112000)

Set DPS (verified): 171.0. Weights run: 6.7s. Verify run: 20.2s. 549 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.432 ± 0.017, spirit=3.498 ± 0.018, mp5=4.742 ± 0.031, crit=0.187 ± 0.006 per rating point (14 rating = 1%, 2.623 per %), spell_haste=not significant (0.313 ± 0.157)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 77.8 healing_power points (8.35 DPS) | yes | Whitemane's Chapeau (7720, -0.45 DPS, sim-verified) [dungeon]; Earthen Silk Hood (254015, -2.24 DPS) [crafted]; Imperial Leather Helm (6433, -2.38 DPS) [dungeon] |
| neck | Amberglow Talisman (10824) | Bring the Light [quest] | 35.0 healing_power points (3.75 DPS) | yes | Necklace of Calisea (1714, -0.05 DPS) [dungeon]; Explorers' League Commendation (7746, -0.38 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.38 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 46.0 healing_power points (4.93 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Sheepshear Mantle (13115, -0.68 DPS, sim-verified) [world_drop]; Mantle of Doan (7712, -1.01 DPS) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 35.0 healing_power points (3.75 DPS) | yes | Ceremonial Centaur Blanket (6789, -0.07 DPS) [quest]; Energy Cloak (9397, -0.29 DPS, sim-verified) [dungeon]; Sergeant Major's Cape (16336, -0.58 DPS) [pvp] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 86.0 healing_power points (9.22 DPS) | yes | Civinad Robes (9623, -1.93 DPS) [quest]; Dreamweave Vest (10021, -2.59 DPS) [crafted]; Doomsayer's Robe (4746, -3.09 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 40.7 healing_power points (4.36 DPS) | yes | Earthen Silk Cuffs (254019, -1.32 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -2.07 DPS) [dungeon]; Silkstream Cuffs (16791, -2.10 DPS) [quest] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 49.1 healing_power points (5.27 DPS) | yes | Earthen Silk Gloves (254017, +0.00 DPS, sim-verified) [crafted]; Windchaser Handguards (14431, -0.39 DPS) [world_drop]; Warden's Gloves (14606, -0.53 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 54.4 healing_power points (5.84 DPS) | yes | Mender's Leather Belt (252523, -0.84 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.23 DPS) [world_drop]; Earthen Silk Belt (7061, -1.98 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 81.6 healing_power points (8.75 DPS) | yes | Stoneweaver Leggings (9407, -1.34 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -1.51 DPS) [crafted]; Wisdom's Leather Leggings (252519, -2.79 DPS) [crafted] |
| feet | Furen's Boots (13100) | World drop [world_drop] | 66.0 healing_power points (7.08 DPS) | yes | Mender's Leather Shoes (252533, -0.80 DPS) [crafted]; Thoughtcast Boots (10578, -0.81 DPS) [dungeon]; Gilded Slippers (254001, -2.36 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 33.5 healing_power points (3.59 DPS) | yes | Welken Ring (5011, -0.20 DPS) [world_drop]; The Queen's Jewel (13094, -0.28 DPS) [world_drop]; Blush Ember Ring (13093, -0.59 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 33.5 healing_power points (3.59 DPS) | yes | The Queen's Jewel (13094, -0.28 DPS) [world_drop]; Welken Ring (5011, -0.39 DPS, sim-verified) [world_drop]; Blush Ember Ring (13093, -0.59 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (171.0 DPS) | yes | Gnarled Ash Staff (791, -1.20 DPS) [world_drop]; Silver Spade (4128, -1.20 DPS) [quest]; Death Speaker Scepter (2816, -8.76 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Amberglow Talisman; shoulder: Earthen Silk Shoulders; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; waist: Gilded Cord; legs: Warchief Kilt; feet: Furen's Boots; trinket1: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 549, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5053005153113200)

Set DPS (verified): 266.3. Weights run: 6.8s. Verify run: 25.4s. 719 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.710 ± 0.020, spirit=3.574 ± 0.022, mp5=4.583 ± 0.026, crit=0.269 ± 0.008 per rating point (14 rating = 1%, 3.766 per %), spell_haste=not significant (0.388 ± 0.175)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 102.4 healing_power points (10.78 DPS) | yes | Engineer's Guild Headpiece (9534, -0.50 DPS) [quest]; Knight-Lieutenant's Restored Leather Helm (220874, -1.02 DPS) [vendor]; Papal Fez (9431, -2.02 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 38.2 healing_power points (4.03 DPS) | yes | Necklace of Calisea (1714, -0.13 DPS) [dungeon]; Amberglow Talisman (10824, -0.26 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.48 DPS) [dungeon] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 77.5 healing_power points (8.16 DPS) | yes | Knight-Lieutenant's Restored Leather Spaulders (220876, -0.28 DPS) [vendor]; Mender's Leather Shoulders (252538, -1.38 DPS) [crafted]; Ironfeather Shoulders (15067, -1.55 DPS) [crafted] |
| back | Wingveil Cloak (10802) | Sunken Temple: Spawn of Hakkar [dungeon] | 44.3 healing_power points (4.67 DPS) | yes | Arcane Cloak (8286, -0.15 DPS) [world_drop]; Chainlink Towel (9648, -0.36 DPS) [quest]; Traveler's Cloak (8297, -0.53 DPS) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 108.2 healing_power points (11.39 DPS) | yes | Feathered Breastplate (8349, -0.56 DPS) [crafted]; Knight's Restored Leather Jerkin (220878, -2.05 DPS) [vendor]; Grimlok's Tribal Vestments (9415, -2.17 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 52.4 healing_power points (5.52 DPS) | yes | Nethergeld Cuffs (254061, -0.27 DPS) [crafted]; Enchanted Kodo Bracers (13119, -1.04 DPS) [world_drop]; Venomshroud Armguards (14439, -1.79 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | 78.6 healing_power points (8.27 DPS) | yes | Gilded Gloves (254095, -0.41 DPS) [crafted]; Earthenweave Gloves (254075, -1.39 DPS) [crafted]; Keeper's Gloves (14666, -2.53 DPS) [world_drop] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 65.3 healing_power points (6.88 DPS) | yes | Gilded Cord (254037, -0.86 DPS) [crafted]; Mender's Leather Belt (252523, -1.54 DPS) [crafted]; Mender's Leather Waistguard (252477, -1.56 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | sim-verified (266.3 DPS) | yes | Dalewind Trousers (13008, +0.00 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -0.02 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.03 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 85.2 healing_power points (8.97 DPS) | yes | Sergeant Major's Restored Leather Boots (220884, -1.30 DPS) [vendor]; Furen's Boots (13100, -1.69 DPS) [world_drop]; Earthenweave Boots (254093, -2.09 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 53.1 healing_power points (5.60 DPS) | yes | Coldwater Ring (4550, -1.67 DPS) [quest]; Seedtime Hoop (9655, -1.83 DPS) [quest]; Electrocutioner Lagnut (9447, -2.01 DPS) [dungeon] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 46.5 healing_power points (4.89 DPS) | yes | Coldwater Ring (4550, -0.97 DPS) [quest]; Seedtime Hoop (9655, -1.13 DPS) [quest]; Electrocutioner Lagnut (9447, -1.31 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -1.51 DPS) [quest]; Uther's Strength (11302, -2.62 DPS) [world_drop]; Alchemist's Stone (13503, -3.01 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -0.75 DPS) [quest]; Uther's Strength (11302, -1.87 DPS) [world_drop]; Alchemist's Stone (13503, -2.26 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (266.3 DPS) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Living Shoulders; back: Wingveil Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 719, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5053005153113200)

Set DPS (verified): 383.8. Weights run: 7.0s. Verify run: 64.1s. 1597 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=2.376 ± 0.052, spirit=4.523 ± 0.036, mp5=6.292 ± 0.047, crit=0.409 ± 0.013 per rating point (14 rating = 1%, 5.727 per %), spell_haste=not significant (-1.657 ± 0.446)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 179.1 healing_power points (20.19 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -2.94 DPS) [vendor]; Wildheart Cowl (16720, -2.94 DPS) [dungeon]; Devout Crown (16693, -3.86 DPS) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 67.8 healing_power points (7.65 DPS) | yes | Horizon Choker (13085, -1.86 DPS) [world_drop]; Necklace of Calisea (1714, -2.20 DPS) [dungeon]; Amberglow Talisman (10824, -2.55 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 101.6 healing_power points (11.45 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -0.30 DPS) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 63.3 healing_power points (7.14 DPS) | yes | Ethereal Mist Cape (11873, -0.22 DPS) [quest]; Indomitable Cloak (14683, -0.51 DPS) [world_drop]; Sergeant Major's Cape (16337, -0.92 DPS) [pvp] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 159.5 healing_power points (17.97 DPS) | yes | Ironfeather Breastplate (15066, -0.48 DPS) [crafted]; Wildheart Vest (16706, -1.18 DPS) [dungeon]; Devout Robe (16690, -1.64 DPS) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 75.2 healing_power points (8.48 DPS) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, -0.90 DPS) [crafted]; Mender's Leather Bracers (252543, -1.41 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; Swarmtender's Gloves (275607, -1.04 DPS) [crafted]; Marshal's Dragonhide Gauntlets (231706, -1.78 DPS) [pvp] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Waistguard (16447, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, +0.00 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 163.7 healing_power points (18.45 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, -1.33 DPS) [vendor]; Elderwild Pants (279254, -1.70 DPS) [crafted]; Living Leggings (15060, -2.77 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Sandals of the Insurgent (13111, -0.25 DPS) [world_drop]; Wildheart Boots (16715, -0.27 DPS) [dungeon] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 73.5 healing_power points (8.29 DPS) | yes | Choking Band (11868, -1.66 DPS) [quest]; Coldwater Ring (4550, -2.89 DPS) [quest]; Seedtime Hoop (9655, -3.19 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 68.5 healing_power points (7.73 DPS) | yes | Choking Band (11868, -1.10 DPS) [quest]; Coldwater Ring (4550, -2.33 DPS) [quest]; Seedtime Hoop (9655, -2.63 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.86 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.04 DPS) [quest]; Alchemist's Stone (13503, -4.08 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -0.84 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.02 DPS) [quest]; Alchemist's Stone (13503, -3.06 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (383.9 DPS) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5053005153113200)

Set DPS (verified): 742.0. Weights run: 5.0s. Verify run: 43.2s. 1597 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.152, intellect=1.314 ± 0.056, spirit=1.840 ± 0.055, mp5=2.190 ± 0.061, crit=0.428 ± 0.026 per rating point (14 rating = 1%, 5.994 per %), spell_haste=not significant (0.973 ± 0.916)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 122.4 healing_power points (34.68 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -8.46 DPS) [vendor]; Devout Crown (16693, -12.26 DPS) [dungeon]; Mooncloth Circlet (14140, -12.41 DPS) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | sim-verified (741.9 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Horizon Choker (13085, -0.52 DPS) [world_drop]; Necklace of Calisea (1714, -1.57 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Spaulders (16718, +0.00 DPS) [dungeon]; Lieutenant Commander's Dragonhide Pauldrons (227201, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | sim-verified (741.9 DPS) | yes | Ethereal Mist Cape (11873, +0.00 DPS) [quest]; Wingveil Cloak (10802, -0.15 DPS) [dungeon]; Sergeant Major's Cape (16337, -0.15 DPS) [pvp] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-verified (741.9 DPS) | yes | Knight-Captain's Dragonhide Tunic (227196, -0.50 DPS) [pvp]; Field Marshal's Dragonhide Tunic (231702, -1.76 DPS) [vendor]; Devout Robe (16690, -2.30 DPS) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Mender's Leather Bracers (252543, -0.19 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; Mender's Leather Gauntlets (252551, -0.01 DPS) [crafted]; Marshal's Dragonhide Gauntlets (231706, -0.66 DPS) [pvp] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (741.9 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, +0.00 DPS) [crafted]; Wildheart Belt (16716, -0.60 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 86.0 healing_power points (24.37 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, +0.00 DPS) [vendor]; Elderwild Pants (279254, -0.53 DPS) [crafted]; Marshal's Dragonhide Legguards (231703, -2.90 DPS) [pvp] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Wildheart Boots (16715, -0.37 DPS) [dungeon]; Sergeant Major's Restored Leather Boots (220884, -1.10 DPS) [vendor] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 33.4 healing_power points (9.46 DPS) | yes | Choking Band (11868, -2.68 DPS) [quest]; Electrocutioner Lagnut (9447, -3.26 DPS) [dungeon]; Runemender's Seal (273031, -3.26 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 30.0 healing_power points (8.49 DPS) | yes | Choking Band (11868, -1.71 DPS) [quest]; Electrocutioner Lagnut (9447, -2.29 DPS) [dungeon]; Runemender's Seal (273031, -2.29 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -1.04 DPS) [quest]; Uther's Strength (11302, -1.16 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -2.09 DPS) [quest] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (741.9 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (741.9 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Spire of Hakkar (10844, +0.00 DPS) [world]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5051000000000000)

Set DPS (verified): 50.9. Weights run: 4.5s. Verify run: 14.4s. 226 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.279 ± 0.009, spirit=1.547 ± 0.011, mp5=3.188 ± 0.015, crit=0.074 ± 0.003 per rating point (14 rating = 1%, 1.042 per %), spell_haste=-3.937 ± 0.154

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.95 DPS) | yes | Wisdom's Leather Hood (252507, -0.14 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.25 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.44 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.6 healing_power points (0.50 DPS) | yes | Scholarly Pendant (277203, -0.54 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 healing_power points (1.25 DPS) | yes | Slime-encrusted Pads (6461, -0.21 DPS) [dungeon]; Forest Leather Mantle (4709, -0.58 DPS) [world_drop]; Serpent's Shoulders (5404, -1.54 DPS, sim-verified) [dungeon] |
| back | Spirit Cloak (4792) (or Antiquated Cloak (4799), Sylvan Cloak (4793), Heavy Woolen Cloak (4311)) | Jennabink Powerseam [vendor] | 6.2 healing_power points (0.67 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Sylvan Cloak (4793, +0.00 DPS) [vendor]; Antiquated Cloak (4799, +0.00 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.0 healing_power points (2.82 DPS) | yes | Filigreed Pristine Gown (253901, -0.22 DPS) [crafted]; Corsair's Overshirt (5202, -0.29 DPS) [dungeon]; Robe of the Moccasin (6465, -1.14 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 16.0 healing_power points (1.73 DPS) | yes | Drakewing Bands (12999, -1.03 DPS, sim-verified) [world_drop]; Owl Bracers (4796, -1.04 DPS) [vendor]; Bright Bracers (3647, -1.17 DPS) [world_drop] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 15.1 healing_power points (1.64 DPS) | yes | Pristine Gloves (253913, -0.18 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.27 DPS) [world_drop]; Bright Gloves (3066, -0.41 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.1 healing_power points (1.74 DPS) | yes | Wisdom's Leather Belt (252433, -0.32 DPS) [crafted]; Novice Ardent's Sash (253887, -0.58 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.64 DPS) [world_drop] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.67 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.69 DPS) [world_drop]; Stormrider's Leather Pants (252502, -2.17 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.88 DPS) | yes | Pristine Boots (253889, -0.49 DPS) [crafted]; Black Whelp Slippers (252424, -0.58 DPS, sim-verified) [crafted]; Smoldering Boots (3076, -0.88 DPS) [world] |
| finger1 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (1.01 DPS) | yes | Loop of Sacrifice (281673, -0.31 DPS) [quest]; Volcanic Rock Ring (12053, -0.59 DPS) [world_drop]; Horn Ring (17692, -0.67 DPS) [quest] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.7 healing_power points (0.83 DPS) | yes | Loop of Sacrifice (281673, -0.36 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.42 DPS) [world_drop]; Horn Ring (17692, -0.50 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Staff of the Blessed Seer (2271), Hammerbone (270018)) | The Wrath of Rath'mael [quest] | 15.5 healing_power points (1.68 DPS) | yes | Staff of the Blessed Seer (2271, +0.00 DPS, sim-verified) [dungeon]; Hammerbone (270018, +0.00 DPS) [quest]; Twisted Chanter's Staff (890, -0.29 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Spirit Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Band of Purification; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff

No-known-source sample (15 of 226, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (tauren, 0000000000000000-00000000000000000000-5053005110010000)

Set DPS (verified): 112.3. Weights run: 6.0s. Verify run: 21.2s. 414 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.614 ± 0.014, spirit=3.959 ± 0.021, mp5=4.933 ± 0.028, crit=0.126 ± 0.004 per rating point (14 rating = 1%, 1.765 per %), spell_haste=not significant (-0.210 ± 0.293)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 56.8 healing_power points (6.03 DPS) | yes | Enduring Cap (3020, -0.67 DPS, sim-verified) [world_drop]; Stormrider's Leather Helm (252514, -1.82 DPS) [crafted]; Nightsky Cowl (4039, -1.87 DPS) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 23.8 healing_power points (2.52 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.58 DPS) [world_drop] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 32.7 healing_power points (3.48 DPS) | yes | Desert Shoulders (15457, -0.51 DPS) [quest]; Insignia Mantle (4721, -0.53 DPS) [world_drop]; Ghostly Mantle (3324, -1.10 DPS, sim-verified) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 38.7 healing_power points (4.11 DPS) | yes | Amy's Blanket (13005, -1.41 DPS, sim-verified) [world_drop]; Nightsky Cloak (4719, -1.59 DPS) [world_drop]; Insignia Cloak (4722, -1.59 DPS) [world_drop] |
| chest | Pressed Felt Robe (1997) | Bloodscalp Mystic [world] | 52.3 healing_power points (5.56 DPS) | yes | Wisdom's Leather Tunic (252511, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Silver-thread Armor (7110, -0.25 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightsky Wristbands (6407, -0.14 DPS) [world_drop]; Dokebi Bracers (14580, -0.23 DPS) [world_drop]; Drakewing Bands (12999, -1.58 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 55.6 healing_power points (5.91 DPS) | yes | Hotshot Pilot's Gloves (9491, -2.00 DPS, sim-verified) [dungeon]; Stonecloth Gloves (14411, -2.54 DPS) [world_drop]; Cartilage Shapers (273054, -2.54 DPS) [world] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 52.4 healing_power points (5.57 DPS) | yes | Dokebi Cord (14578, -2.36 DPS) [world_drop]; Lilac Sash (6780, -2.77 DPS) [quest]; Silver-lined Belt (13011, -2.83 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 60.0 healing_power points (6.38 DPS) | yes | Stormrider's Leather Kilt (252518, -0.55 DPS) [crafted]; Pristine Leggings (253987, -0.74 DPS) [crafted]; Earthen Leggings (253999, -1.43 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 47.1 healing_power points (5.01 DPS) | yes | Resilient Boots (14399, -1.64 DPS) [world_drop]; Stonecloth Boots (14408, -1.64 DPS) [world_drop]; Boots of the Enchanter (4325, -3.56 DPS, sim-verified) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 36.7 healing_power points (3.90 DPS) | yes | The Queen's Jewel (13094, -0.19 DPS) [world_drop]; Solid Gold Band (279845, -0.63 DPS) [quest]; Band of Thorns (5007, -1.38 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 36.7 healing_power points (3.90 DPS) | yes | The Queen's Jewel (13094, -0.47 DPS, sim-verified) [world_drop]; Solid Gold Band (279845, -0.63 DPS) [quest]; Band of Thorns (5007, -1.38 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Friar (3415, -1.66 DPS) [dungeon]; Odo's Ley Staff (6318, -2.18 DPS) [dungeon]; Gnarled Ash Staff (791, -3.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Nightsky Mantle; back: Glowing Thresher Cape; chest: Pressed Felt Robe; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 0000000000000000-00000000000000000000-5053005153112000)

Set DPS (verified): 170.6. Weights run: 6.7s. Verify run: 20.1s. 547 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.432 ± 0.017, spirit=3.498 ± 0.018, mp5=4.742 ± 0.031, crit=0.187 ± 0.006 per rating point (14 rating = 1%, 2.623 per %), spell_haste=not significant (0.313 ± 0.157)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 77.8 healing_power points (8.35 DPS) | yes | Whitemane's Chapeau (7720, -0.36 DPS, sim-verified) [dungeon]; Earthen Silk Hood (254015, -2.24 DPS) [crafted]; Imperial Leather Helm (6433, -2.38 DPS) [dungeon] |
| neck | Amberglow Talisman (10824) | Bring the End [quest] | 35.0 healing_power points (3.75 DPS) | yes | Necklace of Calisea (1714, -0.05 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.38 DPS) [dungeon]; Triune Amulet (7722, -1.13 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 46.0 healing_power points (4.93 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Mantle of Doan (7712, -1.01 DPS) [dungeon]; Sheepshear Mantle (13115, -1.24 DPS, sim-verified) [world_drop] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 45.5 healing_power points (4.88 DPS) | yes | Glowing Thresher Cape (6901, -0.84 DPS, sim-verified) [dungeon]; Energy Cloak (9397, -1.13 DPS) [dungeon]; Stargazer Cloak (9660, -1.13 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 86.0 healing_power points (9.22 DPS) | yes | Civinad Robes (9623, -1.93 DPS) [quest]; Dreamweave Vest (10021, -2.59 DPS) [crafted]; Doomsayer's Robe (4746, -3.11 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 40.7 healing_power points (4.36 DPS) | yes | Earthen Silk Cuffs (254019, -1.59 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -2.07 DPS) [dungeon]; Silkstream Cuffs (16791, -2.10 DPS) [quest] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 49.1 healing_power points (5.27 DPS) | yes | Earthen Silk Gloves (254017, -0.34 DPS) [crafted]; Windchaser Handguards (14431, -0.39 DPS) [world_drop]; Warden's Gloves (14606, -0.53 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 54.4 healing_power points (5.84 DPS) | yes | Mender's Leather Belt (252523, -1.08 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.23 DPS) [world_drop]; Earthen Silk Belt (7061, -1.98 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 81.6 healing_power points (8.75 DPS) | yes | Stoneweaver Leggings (9407, -1.25 DPS, sim-verified) [dungeon]; Stormcloth Pants (10010, -1.51 DPS) [crafted]; Mistspray Kilt (4976, -2.37 DPS) [quest] |
| feet | Furen's Boots (13100) | World drop [world_drop] | 66.0 healing_power points (7.08 DPS) | yes | Mender's Leather Shoes (252533, -0.80 DPS) [crafted]; Thoughtcast Boots (10578, -0.81 DPS) [dungeon]; Gilded Slippers (254001, -2.36 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 33.5 healing_power points (3.59 DPS) | yes | Welken Ring (5011, -0.20 DPS) [world_drop]; The Queen's Jewel (13094, -0.28 DPS) [world_drop]; Solid Gold Band (279845, -0.54 DPS) [quest] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 33.5 healing_power points (3.59 DPS) | yes | The Queen's Jewel (13094, -0.28 DPS) [world_drop]; Welken Ring (5011, -0.36 DPS, sim-verified) [world_drop]; Solid Gold Band (279845, -0.54 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (170.6 DPS) | yes | Gnarled Ash Staff (791, -1.20 DPS) [world_drop]; Silver Spade (4128, -1.20 DPS) [quest]; Death Speaker Scepter (2816, -9.97 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Amberglow Talisman; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; waist: Gilded Cord; legs: Warchief Kilt; feet: Furen's Boots; trinket1: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 547, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 50 (tauren, 4300000000000000-00000000000000000000-5053005153113200)

Set DPS (verified): 263.1. Weights run: 6.8s. Verify run: 25.4s. 717 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.710 ± 0.020, spirit=3.574 ± 0.022, mp5=4.583 ± 0.026, crit=0.269 ± 0.008 per rating point (14 rating = 1%, 3.766 per %), spell_haste=not significant (0.388 ± 0.175)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 102.4 healing_power points (10.78 DPS) | yes | Engineer's Guild Headpiece (9534, -0.50 DPS) [quest]; Blood Guard's Restored Leather Helm (220875, -1.02 DPS) [vendor]; Papal Fez (9431, -2.02 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 38.2 healing_power points (4.03 DPS) | yes | Choker of the High Shaman (4112, -0.11 DPS) [quest]; Necklace of Calisea (1714, -0.13 DPS) [dungeon]; Amberglow Talisman (10824, -0.26 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 77.5 healing_power points (8.16 DPS) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.28 DPS) [vendor]; Mender's Leather Shoulders (252538, -1.38 DPS) [crafted]; Ironfeather Shoulders (15067, -1.55 DPS) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 46.5 healing_power points (4.89 DPS) | yes | Wingveil Cloak (10802, -0.23 DPS) [dungeon]; Arcane Cloak (8286, -0.38 DPS) [world_drop]; Chainlink Towel (9648, -0.59 DPS) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 108.2 healing_power points (11.39 DPS) | yes | Feathered Breastplate (8349, -0.56 DPS) [crafted]; Stone Guard's Restored Leather Jerkin (220879, -2.05 DPS) [vendor]; Grimlok's Tribal Vestments (9415, -2.17 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 52.4 healing_power points (5.52 DPS) | yes | Nethergeld Cuffs (254061, -0.27 DPS) [crafted]; Enchanted Kodo Bracers (13119, -1.04 DPS) [world_drop]; Venomshroud Armguards (14439, -1.79 DPS) [world_drop] |
| hands | Mender's Leather Gauntlets (252551) | Leatherworking [crafted] | 78.6 healing_power points (8.27 DPS) | yes | Gilded Gloves (254095, -0.41 DPS) [crafted]; Earthenweave Gloves (254075, -1.39 DPS) [crafted]; Keeper's Gloves (14666, -2.53 DPS) [world_drop] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 65.3 healing_power points (6.88 DPS) | yes | Gilded Cord (254037, -0.86 DPS) [crafted]; Mender's Leather Belt (252523, -1.54 DPS) [crafted]; Mender's Leather Waistguard (252477, -1.56 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | sim-verified (263.2 DPS) | yes | Dalewind Trousers (13008, +0.00 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -0.02 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.03 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 85.2 healing_power points (8.97 DPS) | yes | First Sergeant's Restored Leather Boots (220885, -1.30 DPS) [vendor]; Furen's Boots (13100, -1.69 DPS) [world_drop]; Earthenweave Boots (254093, -2.09 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 53.1 healing_power points (5.60 DPS) | yes | Electrocutioner Lagnut (9447, -2.01 DPS) [dungeon]; Runemender's Seal (273031, -2.01 DPS) [world]; Welken Ring (5011, -2.06 DPS) [world_drop] |
| finger2 | Coldwater Ring (4550) | Death From Below [quest] | 37.3 healing_power points (3.93 DPS) | yes | Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; Runemender's Seal (273031, -0.35 DPS) [world]; Welken Ring (5011, -0.39 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -2.62 DPS) [world_drop]; Alchemist's Stone (13503, -3.01 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.87 DPS) [world_drop]; Alchemist's Stone (13503, -2.26 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (263.2 DPS) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Living Shoulders; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Coldwater Ring; trinket2: Evonice's Landin' Pilla; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 717, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60 (tauren, 4532200100000000-00000000000000000000-5053005153113200)

Set DPS (verified): 380.5. Weights run: 7.0s. Verify run: 63.6s. 1597 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.006, intellect=2.376 ± 0.052, spirit=4.523 ± 0.036, mp5=6.292 ± 0.047, crit=0.409 ± 0.013 per rating point (14 rating = 1%, 5.727 per %), spell_haste=not significant (-1.657 ± 0.446)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 179.1 healing_power points (20.19 DPS) | yes | Champion's Dragonhide Headdress (227205, -2.94 DPS) [vendor]; Wildheart Cowl (16720, -2.94 DPS) [dungeon]; Devout Crown (16693, -3.86 DPS) [dungeon] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 67.8 healing_power points (7.65 DPS) | yes | Horizon Choker (13085, -1.86 DPS) [world_drop]; Necklace of Calisea (1714, -2.20 DPS) [dungeon]; Choker of the High Shaman (4112, -2.23 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 101.6 healing_power points (11.45 DPS) | yes | Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Wildheart Spaulders (16718, -0.30 DPS) [dungeon] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 63.3 healing_power points (7.14 DPS) | yes | Cloak of Blight (6832, -0.51 DPS) [quest]; Indomitable Cloak (14683, -0.51 DPS) [world_drop]; Wingveil Cloak (10802, -0.97 DPS) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 159.5 healing_power points (17.97 DPS) | yes | Ironfeather Breastplate (15066, -0.48 DPS) [crafted]; Wildheart Vest (16706, -1.18 DPS) [dungeon]; Devout Robe (16690, -1.64 DPS) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 75.2 healing_power points (8.48 DPS) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, -0.90 DPS) [crafted]; Mender's Leather Bracers (252543, -1.41 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; Swarmtender's Gloves (275607, -1.04 DPS) [crafted]; General's Dragonhide Gauntlets (231676, -1.78 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Belt (16556, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, +0.00 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 163.7 healing_power points (18.45 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, -1.33 DPS) [vendor]; Elderwild Pants (279254, -1.70 DPS) [crafted]; Living Leggings (15060, -2.77 DPS) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Sandals of the Insurgent (13111, -0.25 DPS) [world_drop]; Wildheart Boots (16715, -0.27 DPS) [dungeon] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 73.5 healing_power points (8.29 DPS) | yes | Coldwater Ring (4550, -2.89 DPS) [quest]; Maiden's Circle (13001, -3.35 DPS) [world_drop]; Welken Ring (5011, -3.38 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 68.5 healing_power points (7.73 DPS) | yes | Coldwater Ring (4550, -2.33 DPS) [quest]; Maiden's Circle (13001, -2.79 DPS) [world_drop]; Welken Ring (5011, -2.82 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -1.02 DPS) [quest]; Alchemist's Stone (13503, -4.08 DPS) [crafted]; Uther's Strength (11302, -4.09 DPS) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -2.22 DPS) [crafted]; Uther's Strength (11302, -2.23 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (380.6 DPS) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket2: Shard of the Splithooves

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5053005153113200)

Set DPS (verified): 739.8. Weights run: 5.0s. Verify run: 43.1s. 1597 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.152, intellect=1.314 ± 0.056, spirit=1.840 ± 0.055, mp5=2.190 ± 0.061, crit=0.428 ± 0.026 per rating point (14 rating = 1%, 5.994 per %), spell_haste=not significant (0.973 ± 0.916)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 122.4 healing_power points (34.68 DPS) | yes | Champion's Dragonhide Headdress (227205, -8.46 DPS) [vendor]; Devout Crown (16693, -12.26 DPS) [dungeon]; Mooncloth Circlet (14140, -12.41 DPS) [crafted] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | sim-verified (739.7 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Horizon Choker (13085, -0.52 DPS) [world_drop]; Necklace of Calisea (1714, -1.57 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Spaulders (16718, +0.00 DPS) [dungeon]; Champion's Dragonhide Pauldrons (227207, +0.00 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor] |
| back | Cloak of Earth and Sky (279270) | Tailoring [crafted] | 25.8 healing_power points (7.30 DPS) | yes | Wingveil Cloak (10802, -0.15 DPS) [dungeon]; Cloak of Blight (6832, -0.52 DPS) [quest]; Indomitable Cloak (14683, -0.52 DPS) [world_drop] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-verified (739.7 DPS) | yes | Legionnaire's Dragonhide Tunic (227202, -0.50 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -1.76 DPS) [vendor]; Devout Robe (16690, -2.30 DPS) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted]; Mender's Leather Bracers (252543, -0.19 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wildheart Gloves (16717, +0.00 DPS) [dungeon]; Mender's Leather Gauntlets (252551, -0.01 DPS) [crafted]; General's Dragonhide Gauntlets (231676, -0.66 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (739.7 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Elderwild Waistcord (279252, +0.00 DPS) [crafted]; Wildheart Belt (16716, -0.60 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 86.0 healing_power points (24.37 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; Elderwild Pants (279254, -0.53 DPS) [crafted]; General's Dragonhide Legguards (231673, -2.90 DPS) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Wildheart Boots (16715, -0.37 DPS) [dungeon]; First Sergeant's Restored Leather Boots (220885, -1.10 DPS) [vendor] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 33.4 healing_power points (9.46 DPS) | yes | Electrocutioner Lagnut (9447, -3.26 DPS) [dungeon]; Runemender's Seal (273031, -3.26 DPS) [world]; Coldwater Ring (4550, -3.65 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 30.0 healing_power points (8.49 DPS) | yes | Electrocutioner Lagnut (9447, -2.29 DPS) [dungeon]; Runemender's Seal (273031, -2.29 DPS) [world]; Coldwater Ring (4550, -2.68 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.16 DPS) [world_drop]; Shard of the Splithooves (10659, -2.53 DPS) [quest]; Alchemist's Stone (13503, -4.17 DPS) [crafted] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (739.7 DPS) | yes | Uther's Strength (11302, -0.11 DPS) [world_drop]; Shard of the Splithooves (10659, -1.49 DPS) [quest]; Alchemist's Stone (13503, -3.13 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (739.7 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Spire of Hakkar (10844, +0.00 DPS) [world]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; back: Cloak of Earth and Sky; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

