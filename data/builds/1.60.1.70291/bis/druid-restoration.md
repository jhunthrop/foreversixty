# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 48.2. Weights run: 1.9s. Verify run: 6.7s. 188 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.284 ± 0.006, spirit=1.552 ± 0.012, mp5=3.307 ± 0.021, crit=0.044 ± 0.004 per rating point (14 rating = 1%, 0.622 per %), spell_haste=-5.580 ± 0.149

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.86 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.23 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 6.2 healing_power points (0.64 DPS) | yes | Scholarly Pendant (277203, -1.34 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.6 healing_power points (1.19 DPS) | yes | Forest Leather Mantle (4709, -0.55 DPS) [world_drop]; Prospector's Pads (14566, -0.55 DPS) [world_drop]; Slime-encrusted Pads (6461, -0.82 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 12.1 healing_power points (1.25 DPS) | yes | Spirit Cloak (4792, -0.61 DPS) [vendor]; Sylvan Cloak (4793, -0.61 DPS) [vendor]; Regent's Cloak (5969, -2.21 DPS, sim-verified) [world] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.1 healing_power points (2.69 DPS) | yes | Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.27 DPS) [dungeon]; Robe of the Moccasin (6465, -1.02 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 16.3 healing_power points (1.68 DPS) | yes | Owl Bracers (4796, -1.02 DPS) [vendor]; Bravo's Armbands (270015, -1.04 DPS) [quest]; Drakewing Bands (12999, -2.58 DPS, sim-verified) [world_drop] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 15.1 healing_power points (1.56 DPS) | yes | Magefist Gloves (12977, -0.26 DPS) [world_drop]; Bright Gloves (3066, -0.39 DPS) [world_drop]; Pristine Gloves (253913, -1.53 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.1 healing_power points (1.66 DPS) | yes | Wisdom's Leather Belt (252433, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.60 DPS) [world_drop]; Novice Ardent's Sash (253887, -1.74 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.50 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.80 DPS) | yes | Pristine Boots (253889, -0.47 DPS) [crafted]; Kimbra Boots (6191, -0.79 DPS) [quest]; Black Whelp Slippers (252424, -1.92 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.9 healing_power points (1.22 DPS) | yes | Deep Fathom Ring (6463, -0.42 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.54 DPS) [rep] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.28 DPS) [rep]; Deep Fathom Ring (6463, -2.12 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 15.7 healing_power points (1.62 DPS) | yes | Twisted Chanter's Staff (890, -0.30 DPS) [world_drop]; Gnarled Hermit's Staff (1539, -0.50 DPS) [world]; Staff of the Blessed Seer (2271, -1.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Band of Purification; main_hand: Staff of Westfall

No-known-source sample (15 of 188, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 100.7. Weights run: 1.9s. Verify run: 7.6s. 340 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.002, intellect=1.454 ± 0.020, spirit=2.901 ± 0.026, mp5=4.443 ± 0.039, crit=0.058 ± 0.007 per rating point (14 rating = 1%, 0.818 per %), spell_haste=not significant (-0.802 ± 0.370)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embalmed Shroud (7691) | Scarlet Monastery: Fallen Champion [dungeon] | 50.8 healing_power points (5.18 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Whisperwind Headdress (6688, -0.30 DPS) [dungeon]; Enduring Cap (3020, -0.74 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.4 healing_power points (1.78 DPS) | yes | Glowing Green Talisman (5002, -0.00 DPS) [world_drop]; Pendant of Myzrael (4614, -0.24 DPS, sim-verified) [dungeon]; Crystal Starfire Medallion (5003, -0.30 DPS) [world_drop] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 30.5 healing_power points (3.11 DPS) | yes | Mantle of Honor (3560, -0.23 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.44 DPS) [world_drop]; Faerie Mantle (5820, -0.45 DPS) [quest] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 32.1 healing_power points (3.27 DPS) | yes | Glowing Thresher Cape (6901, -0.36 DPS, sim-verified) [dungeon]; Prelacy Cape (7004, -0.38 DPS) [quest]; Caretaker's Cape (19533, -0.77 DPS) [rep] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.1 healing_power points (4.50 DPS) | yes | Pristine Gown (253961, -0.22 DPS, sim-verified) [crafted]; Pressed Felt Robe (1997, -0.50 DPS) [world]; Robe of the Moccasin (6465, -0.62 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.6 healing_power points (2.10 DPS) | yes | Drakewing Bands (12999, -0.33 DPS) [world_drop]; Dokebi Bracers (14580, -0.47 DPS) [world_drop]; Nightsky Wristbands (6407, -0.69 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.1 healing_power points (4.29 DPS) | yes | Zodiac Gloves (7106, -1.63 DPS) [quest]; Shilly Mitts (9609, -1.63 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.73 DPS, sim-verified) [dungeon] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.1 healing_power points (4.60 DPS) | yes | Dokebi Cord (14578, -2.23 DPS) [world_drop]; Resilient Cord (14406, -2.53 DPS) [world_drop]; Silver-lined Belt (13011, -2.66 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.5 healing_power points (5.25 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.96 DPS) [crafted]; Earthen Leggings (253999, -1.00 DPS, sim-verified) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.26 DPS) | yes | Glinteye Slippers (273024, -1.60 DPS) [dungeon]; Acidic Walkers (9454, -1.89 DPS) [dungeon]; Soggy Boots (274747, -2.56 DPS, sim-verified) [vendor] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.5 healing_power points (3.11 DPS) | yes | The Queen's Jewel (13094, -0.45 DPS) [world_drop]; Black Pearl Ring (6332, -1.04 DPS) [world]; Ring of Calm (6790, -1.04 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.3 healing_power points (2.68 DPS) | yes | The Queen's Jewel (13094, -0.07 DPS, sim-verified) [world_drop]; Black Pearl Ring (6332, -0.61 DPS) [world]; Ring of Calm (6790, -0.61 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (100.7 DPS) | yes | Royal Diplomatic Scepter (9457, -2.99 DPS) [dungeon]; Death Speaker Scepter (2816, -5.78 DPS) [dungeon]; Manual Crowd Pummeler (9449, -15.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Embalmed Shroud; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Repairman's Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Electrocutioner Lagnut; main_hand: Wind Spirit Staff

No-known-source sample (15 of 340, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 145.6. Weights run: 2.0s. Verify run: 9.5s. 445 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.352 ± 0.022, spirit=2.509 ± 0.016, mp5=4.836 ± 0.048, crit=0.058 ± 0.008 per rating point (14 rating = 1%, 0.814 per %), spell_haste=not significant (0.504 ± 0.436)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 48.1 healing_power points (4.94 DPS) | yes | Embalmed Shroud (7691, -0.20 DPS) [dungeon]; Imperial Leather Helm (6433, -0.44 DPS) [dungeon]; Earthen Silk Hood (254015, -0.51 DPS, sim-verified) [crafted] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 27.0 healing_power points (2.78 DPS) | yes | Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon]; Explorers' League Commendation (7746, -0.46 DPS) [quest]; Amberglow Talisman (10824, -0.55 DPS, sim-verified) [quest] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Crimson Silk Shoulders (7059, -1.00 DPS) [crafted]; Sheepshear Mantle (13115, -1.27 DPS, sim-verified) [world_drop] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | 30.5 healing_power points (3.14 DPS) | yes | Glowing Thresher Cape (6901, -0.36 DPS) [dungeon]; Prelacy Cape (7004, -0.46 DPS) [quest]; Repairman's Cape (9605, -1.18 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.1 healing_power points (8.03 DPS) | yes | Doomsayer's Robe (4746, -2.25 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -3.16 DPS) [crafted]; Civinad Robes (9623, -3.18 DPS) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.5 healing_power points (3.14 DPS) | yes | Mindthrust Bracers (1974, -0.95 DPS) [dungeon]; Earthen Silk Cuffs (254019, -1.11 DPS, sim-verified) [crafted]; Silkstream Cuffs (16791, -1.41 DPS) [quest] |
| hands | Earthen Silk Gloves (254017) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Restorer's Fine Gloves (270059, -0.15 DPS) [quest]; Mender's Leather Gloves (252530, -0.50 DPS) [crafted]; Gloves of Old (9395, -0.97 DPS, sim-verified) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.9 healing_power points (4.92 DPS) | yes | Mender's Leather Belt (252523, -0.59 DPS) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 53.9 healing_power points (5.54 DPS) | yes | Wisdom's Leather Leggings (252519, -0.61 DPS) [crafted]; Stormcloth Pants (10010, -0.83 DPS, sim-verified) [crafted]; Pristine Leggings (253987, -1.02 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 51.0 healing_power points (5.25 DPS) | yes | Everlast Boots (10359, -0.86 DPS) [quest]; Thoughtcast Boots (10578, -1.07 DPS) [dungeon]; Furen's Boots (13100, -1.31 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.0 healing_power points (2.98 DPS) | yes | Welken Ring (5011, -0.48 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.56 DPS) [dungeon]; The Queen's Jewel (13094, -0.64 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 27.0 healing_power points (2.78 DPS) | yes | Welken Ring (5011, +0.00 DPS, sim-verified) [world_drop]; Electrocutioner Lagnut (9447, -0.36 DPS) [dungeon]; The Queen's Jewel (13094, -0.44 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (145.6 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (145.6 DPS) | yes | Hand of Righteousness (7721, -0.39 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -2.61 DPS) [dungeon]; Manual Crowd Pummeler (9449, -17.27 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Holy Shroud; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Earthen Silk Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life

No-known-source sample (15 of 445, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 200.5. Weights run: 2.1s. Verify run: 10.1s. 580 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.534 ± 0.024, spirit=2.780 ± 0.024, mp5=4.433 ± 0.022, crit=0.074 ± 0.011 per rating point (14 rating = 1%, 1.036 per %), spell_haste=1.008 ± 0.239

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 85.1 healing_power points (8.41 DPS) | yes | Knight-Lieutenant's Restored Leather Helm (220874, -0.50 DPS) [vendor]; Cassandra's Grace (13102, -1.59 DPS) [world_drop]; Engineer's Guild Headpiece (9534, -1.60 DPS, sim-verified) [quest] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 35.8 healing_power points (3.54 DPS) | yes | Horizon Choker (13085, -0.44 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -0.55 DPS) [dungeon]; Amberglow Talisman (10824, -0.79 DPS) [quest] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 67.1 healing_power points (6.64 DPS) | yes | Knight-Lieutenant's Restored Leather Spaulders (220876, -0.37 DPS, sim-verified) [vendor]; Mender's Leather Shoulder (252538, -0.98 DPS) [crafted]; Nethergeld Shoulders (254049, -1.21 DPS) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | 38.7 healing_power points (3.82 DPS) | yes | Darkspear Raider's Cloak (272076, -0.33 DPS) [vendor]; Arcane Cloak (8286, -0.53 DPS) [world_drop]; Wingveil Cloak (10802, -0.78 DPS, sim-verified) [dungeon] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 91.1 healing_power points (9.00 DPS) | yes | Doomsayer's Robe (4746, -0.26 DPS, sim-verified) [quest]; Feathered Breastplate (8349, -0.89 DPS) [crafted]; Stormcloth Vest (10020, -1.07 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.4 healing_power points (4.59 DPS) | yes | Nethergeld Cuffs (254061, -0.66 DPS, sim-verified) [crafted]; Enchanted Kodo Bracers (13119, -1.23 DPS) [world_drop]; Venomshroud Armguards (14439, -1.78 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.9 healing_power points (8.99 DPS) | yes | Gilded Gloves (254095, -2.41 DPS) [crafted]; Mender's Leather Gauntlets (252551, -2.56 DPS, sim-verified) [crafted]; Earthenweave Gloves (254075, -3.39 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.6 healing_power points (5.59 DPS) | yes | Gilded Cord (254037, -0.81 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -0.81 DPS) [crafted]; Mender's Leather Belt (252523, -1.16 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 85.8 healing_power points (8.48 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.21 DPS) [dungeon]; Dalewind Trousers (13008, -0.24 DPS) [world_drop]; Jinxed Hoodoo Kilt (9474, -0.89 DPS, sim-verified) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.9 healing_power points (6.71 DPS) | yes | Sergeant Major's Restored Leather Boots (220884, -0.44 DPS, sim-verified) [vendor]; Earthenweave Boots (254093, -1.12 DPS) [crafted]; Furen's Boots (13100, -1.25 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.6 healing_power points (4.21 DPS) | yes | Darkspear Signet (272069, -1.14 DPS) [vendor]; Snake Hoop (6750, -1.22 DPS) [quest]; Coldwater Ring (4550, -1.28 DPS) [quest] |
| finger2 | Choking Band (11868) | Ogre Head On A Stick = Party [quest] | 36.1 healing_power points (3.57 DPS) | yes | Snake Hoop (6750, -0.59 DPS) [quest]; Coldwater Ring (4550, -0.64 DPS) [quest]; Darkspear Signet (272069, -1.10 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (200.4 DPS) | yes | Evonice's Landin' Pilla (18951, -3.82 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.37 DPS) [quest]; Uther's Strength (11302, -4.79 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (200.4 DPS) | yes | Evonice's Landin' Pilla (18951, -0.55 DPS) [quest]; Uther's Strength (11302, -0.70 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -1.10 DPS) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (200.4 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -1.84 DPS) [dungeon]; Hand of Righteousness (7721, -2.52 DPS) [dungeon] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | 41.2 healing_power points (4.08 DPS) | yes | Thermotastic Egg Timer (9644, -1.05 DPS) [quest]; Windchaser Orb (15965, -1.05 DPS) [world_drop]; Hibernal Sphere (15937, -1.80 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Living Shoulders; back: Caretaker's Cape; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Choking Band; trinket1: Darkspear Voodoo Seal; main_hand: Charstone Dirk; off_hand: Twisting Essence Jar

No-known-source sample (15 of 580, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 341.3. Weights run: 2.2s. Verify run: 23.5s. 1222 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.045, intellect=2.354 ± 0.064, spirit=4.397 ± 0.050, mp5=7.029 ± 0.048, crit=0.099 ± 0.014 per rating point (14 rating = 1%, 1.383 per %), spell_haste=not significant (0.397 ± 0.651)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 176.9 healing_power points (17.44 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -2.59 DPS) [vendor]; Feralheart Headdress (226786, -3.17 DPS) [vendor]; Wildheart Cowl (16720, -5.69 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 94.4 healing_power points (9.30 DPS) | yes | Enchanted South Seas Kelp (19612, -3.08 DPS) [quest]; Pristine Enchanted South Seas Kelp (19613, -3.08 DPS) [quest]; Lady Alizabeth's Pendant (13002, -4.07 DPS, sim-verified) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 145.1 healing_power points (14.31 DPS) | yes | Feralheart Mantle (226785, -0.30 DPS, sim-verified) [vendor]; Lieutenant Commander's Dragonhide Pauldrons (227201, -2.17 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -3.16 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 65.5 healing_power points (6.46 DPS) | yes | Cloak of Earth and Sky (279270, -0.39 DPS) [crafted]; Caretaker's Cape (19530, -0.43 DPS) [rep]; Darkspear Raider's Cloak (272063, -2.40 DPS, sim-verified) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 156.6 healing_power points (15.44 DPS) | yes | Ironfeather Breastplate (15066, -1.03 DPS, sim-verified) [crafted]; Wildheart Vest (16706, -1.05 DPS) [dungeon]; Ghostweave Vest (14141, -1.23 DPS) [crafted] |
| wrist | Feralheart Bindings (226782) | Mokvar [vendor] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.36 DPS) [dungeon]; Bracers of Hope (22667, -5.45 DPS, sim-verified) [quest] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 129.0 healing_power points (12.72 DPS) | yes | Wildheart Gloves (16717, -1.24 DPS, sim-verified) [dungeon]; Devout Gloves (16692, -1.55 DPS) [dungeon]; Swarmtender's Gloves (275607, -2.47 DPS) [crafted] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Elderwild Waistcord (279252, -0.82 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.37 DPS) [crafted]; Caretaker's Cord (272398, -4.35 DPS, sim-verified) [vendor] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 160.4 healing_power points (15.81 DPS) | yes | Knight-Captain's Dragonhide Legguards (227200, -0.34 DPS, sim-verified) [vendor]; Feralheart Pants (226787, -1.25 DPS) [vendor]; Elderwild Pants (279254, -1.47 DPS) [crafted] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 152.6 healing_power points (15.04 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.62 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -3.79 DPS) [crafted]; Devout Sandals (16691, -4.27 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 91.3 healing_power points (9.00 DPS) | yes | Band of the Hierophant (13096, -1.92 DPS) [world_drop]; Channeler's Ring (272406, -2.29 DPS) [vendor]; Eye of Adaegus (5266, -2.41 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 73.7 healing_power points (7.26 DPS) | yes | Channeler's Ring (272406, -0.54 DPS) [vendor]; Eye of Adaegus (5266, -0.67 DPS) [world_drop]; Band of the Hierophant (13096, -1.59 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+11.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -5.19 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.06 DPS) [quest]; Shard of the Splithooves (10659, -6.24 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.86 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, -0.04 DPS) [world_drop]; Wind Spirit Staff (6689, -0.28 DPS) [dungeon] |
| off_hand | Lapidis Tankard of Tidesippe (4696) | World drop [world_drop] | 82.1 healing_power points (8.09 DPS) | yes | Lei of the Lifegiver (19312, +0.00 DPS, sim-verified) [rep]; Grand Marshal's Tome of Restoration (234590, -1.99 DPS) [pvp]; Arcane Star (15931, -2.31 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Feralheart Bindings; waist: Feralheart Cord; legs: Devout Skirt; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lapidis Tankard of Tidesippe

No-known-source sample (15 of 1222, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 640.8. Weights run: 1.8s. Verify run: 20.9s. 1222 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.324, intellect=2.826 ± 0.130, spirit=2.735 ± 0.120, mp5=3.745 ± 0.156, crit=0.317 ± 0.049 per rating point (14 rating = 1%, 4.436 per %), spell_haste=not significant (-1.810 ± 2.036)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 162.1 healing_power points (23.63 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -4.00 DPS) [vendor]; Devout Crown (16693, -4.85 DPS) [dungeon]; Feralheart Headdress (226786, -7.96 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 70.9 healing_power points (10.33 DPS) | yes | Enchanted South Seas Kelp (19612, -2.63 DPS) [quest]; Pristine Enchanted South Seas Kelp (19613, -2.63 DPS) [quest]; Drake Tooth Necklace (21531, -3.90 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 108.8 healing_power points (15.86 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, -0.41 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -0.43 DPS) [vendor]; Devout Mantle (16695, -11.59 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 70.3 healing_power points (10.24 DPS) | yes | Darkspear Raider's Cloak (272076, -2.48 DPS) [vendor]; Caretaker's Cape (19530, -3.26 DPS) [rep] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 134.5 healing_power points (19.60 DPS) | yes | Field Marshal's Dragonhide Tunic (231702, -1.56 DPS) [vendor]; Feralheart Embrace (226783, -1.64 DPS) [vendor]; Devout Robe (16690, -2.22 DPS, sim-verified) [dungeon] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 81.9 healing_power points (11.94 DPS) | yes | Marshal's Dragonhide Bracers (16445, -0.13 DPS) [pvp]; Feralheart Bindings (226782, -1.04 DPS, sim-verified) [vendor]; Devout Bracers (16697, -3.05 DPS) [dungeon] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 104.3 healing_power points (15.20 DPS) | yes | Wildheart Gloves (16717, -1.52 DPS) [dungeon]; Devout Gloves (16692, -2.12 DPS) [dungeon]; Raider Handwraps (272097, -12.88 DPS, sim-verified) [vendor] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | 97.0 healing_power points (14.14 DPS) | yes | Marshal's Dragonhide Waistguard (16447, +0.00 DPS) [pvp]; Devout Belt (16696, -0.72 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -8.55 DPS, sim-verified) [crafted] |
| legs | Feralheart Pants (226787) | Mokvar [vendor] | sim-verified (640.8 DPS) | yes | Devout Skirt (16694, +0.00 DPS) [dungeon]; Marshal's Dragonhide Legguards (231703, -1.89 DPS) [pvp]; Knight-Captain's Dragonhide Legguards (227200, -6.58 DPS, sim-verified) [vendor] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | sim-verified (640.8 DPS) | yes | Mooncloth Boots (15802, -2.89 DPS) [crafted]; Devout Sandals (16691, -4.83 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -6.71 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 79.2 healing_power points (11.54 DPS) | yes | Channeler's Ring (272406, -2.58 DPS) [vendor]; Band of the Hierophant (13096, -3.03 DPS) [world_drop]; Blessed Band of Light (272407, -3.83 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 65.2 healing_power points (9.51 DPS) | yes | Band of the Hierophant (13096, -1.01 DPS) [world_drop]; Blessed Band of Light (272407, -1.80 DPS) [vendor]; Channeler's Ring (272406, -3.48 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (640.8 DPS) | yes | Ankh of Life (1713, -3.40 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.20 DPS) [quest]; Shard of the Splithooves (10659, -4.91 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (640.8 DPS) | yes | Evonice's Landin' Pilla (18951, -2.14 DPS) [quest]; Shard of the Splithooves (10659, -2.85 DPS) [quest]; Ankh of Life (1713, -5.22 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (640.8 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, -3.46 DPS) [world_drop]; Wind Spirit Staff (6689, -3.74 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | 58.7 healing_power points (8.56 DPS) | yes | Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Arcane Star (15931, -1.25 DPS) [world_drop]; Lapidis Tankard of Tidesippe (4696, -5.71 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Bracers of Hope; waist: Feralheart Cord; legs: Feralheart Pants; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1222, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 43.9. Weights run: 1.9s. Verify run: 12.0s. 178 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.284 ± 0.006, spirit=1.552 ± 0.012, mp5=3.307 ± 0.021, crit=0.044 ± 0.004 per rating point (14 rating = 1%, 0.622 per %), spell_haste=-5.580 ± 0.149

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.86 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.23 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.7 healing_power points (0.48 DPS) | yes | Scholarly Pendant (277203, -2.00 DPS, sim-verified) [quest] |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (43.9 DPS) | yes | Forest Leather Mantle (4709, -0.38 DPS) [world_drop]; Prospector's Pads (14566, -0.38 DPS) [world_drop]; Magician's Mantle (12998, -1.17 DPS, sim-verified) [world_drop] |
| back | Regent's Cloak (5969) | Ravenclaw Regent [world] | sim-verified (43.9 DPS) | yes | Battle Healer's Cloak (20427, +0.00 DPS, sim-verified) [rep]; Traveler's Shawl (277289, +0.00 DPS) [quest]; Spirit Cloak (4792, -0.16 DPS) [vendor] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-verified (43.9 DPS) | yes | Wisdom's Leather Armor (252493, +0.00 DPS) [crafted]; Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Robe of the Moccasin (6465, -2.14 DPS, sim-verified) [dungeon] |
| wrist | Drakewing Bands (12999) | World drop [world_drop] | sim-verified (43.9 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Crystalline Cuffs (14148, -0.22 DPS) [dungeon]; Owl Bracers (4796, -0.30 DPS) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | sim-verified (43.9 DPS) | yes | Pristine Gloves (253913, -0.03 DPS) [crafted]; Magefist Gloves (12977, -0.26 DPS) [world_drop]; Blight Gloves (279877, -2.22 DPS, sim-verified) [quest] |
| waist | Novice Ardent's Sash (253887) | Tailoring [crafted] | sim-verified (43.9 DPS) | yes | Wisdom's Leather Belt (252433, -0.14 DPS) [crafted]; Keller's Girdle (2911, -0.44 DPS) [world_drop]; Pristine Sash (253925, -1.36 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.50 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (43.9 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Pristine Boots (253889, +0.00 DPS) [crafted]; Black Whelp Slippers (252424, -2.20 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.9 healing_power points (1.22 DPS) | yes | Deep Fathom Ring (6463, -0.42 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; Advisor's Ring (20426, -0.54 DPS) [rep] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; Advisor's Ring (20426, -0.28 DPS) [rep]; Deep Fathom Ring (6463, -2.00 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | sim-verified (43.9 DPS) | yes | Advisor's Gnarled Staff (20425, -0.05 DPS) [pvp]; Gnarled Necromancer's Staff (251534, -0.28 DPS) [quest]; Staff of Orgrimmar (15444, -1.94 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Slime-encrusted Pads; back: Regent's Cloak; chest: Armor of the Fang; wrist: Drakewing Bands; hands: Wisdom's Leather Gloves; waist: Novice Ardent's Sash; legs: Wisdom's Leather Pants; feet: Footpads of the Fang; finger1: Black Pearl Ring; finger2: Band of Purification; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 178, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance; 241089 Scarlet Dagger; 248005 Northshire Hammer; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263006 Scout Ranger's Tunic; 263007 Skyseer's Vest

### Band 30 (tauren, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 101.0. Weights run: 1.9s. Verify run: 6.9s. 332 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.002, intellect=1.454 ± 0.020, spirit=2.901 ± 0.026, mp5=4.443 ± 0.039, crit=0.058 ± 0.007 per rating point (14 rating = 1%, 0.818 per %), spell_haste=not significant (-0.802 ± 0.370)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (101.0 DPS) | yes | Embalmed Shroud (7691, +0.00 DPS) [dungeon]; Whisperwind Headdress (6688, -0.26 DPS) [dungeon]; Enduring Cap (3020, -0.70 DPS) [world_drop] |
| neck | Pendant of Myzrael (4614) | Razorfen Downs: Splinterbone Captain [dungeon] | sim-verified (101.0 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, +0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.29 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 35.1 healing_power points (3.58 DPS) | yes | Batwing Mantle (6697, -0.47 DPS) [dungeon]; Nightsky Mantle (4718, -0.91 DPS) [world_drop]; Desert Shoulders (15457, -1.36 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.2 healing_power points (3.08 DPS) | yes | Battle Healer's Cloak (19529, -0.57 DPS) [rep]; Darkspear Raider's Cloak (272078, -1.01 DPS) [vendor]; Construct Cloak (279848, -1.01 DPS) [quest] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.1 healing_power points (4.50 DPS) | yes | Pristine Gown (253961, -0.24 DPS) [crafted]; Pressed Felt Robe (1997, -0.50 DPS) [world]; Robe of the Moccasin (6465, -0.62 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.6 healing_power points (2.10 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Drakewing Bands (12999, -0.33 DPS) [world_drop]; Dokebi Bracers (14580, -0.47 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.1 healing_power points (4.29 DPS) | yes | Tattered Mittens (270030, -1.50 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.63 DPS) [dungeon]; Blight Gloves (279877, -1.77 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.1 healing_power points (4.60 DPS) | yes | Silver-lined Belt (13011, -2.08 DPS) [world_drop]; Dokebi Cord (14578, -2.23 DPS) [world_drop]; Lilac Sash (6780, -2.38 DPS) [quest] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.5 healing_power points (5.25 DPS) | yes | Earthen Leggings (253999, -0.47 DPS) [crafted]; Pristine Leggings (253987, -0.49 DPS) [crafted]; Stormrider's Leather Kilt (252518, -0.96 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.26 DPS) | yes | Soggy Boots (274747, -1.30 DPS) [vendor]; Vorrel's Boots (7751, -1.30 DPS) [quest]; Glinteye Slippers (273024, -1.60 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.5 healing_power points (3.11 DPS) | yes | The Queen's Jewel (13094, -0.45 DPS) [world_drop]; Black Pearl Ring (6332, -1.04 DPS) [world]; Monkey Ring (6748, -1.04 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.3 healing_power points (2.68 DPS) | yes | The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Black Pearl Ring (6332, -0.61 DPS) [world]; Monkey Ring (6748, -0.61 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Diplomatic Scepter (9457, -2.99 DPS) [dungeon]; Death Speaker Scepter (2816, -5.78 DPS) [dungeon]; Gnarled Ash Staff (791, -8.27 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; wrist: Mindthrust Bracers; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Electrocutioner Lagnut; main_hand: Wind Spirit Staff

No-known-source sample (15 of 332, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 145.5. Weights run: 2.0s. Verify run: 15.6s. 431 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.352 ± 0.022, spirit=2.509 ± 0.016, mp5=4.836 ± 0.048, crit=0.058 ± 0.008 per rating point (14 rating = 1%, 0.814 per %), spell_haste=not significant (0.504 ± 0.436)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 48.1 healing_power points (4.94 DPS) | yes | Embalmed Shroud (7691, -0.20 DPS) [dungeon]; Imperial Leather Helm (6433, -0.44 DPS) [dungeon]; Earthen Silk Hood (254015, -0.99 DPS, sim-verified) [crafted] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 27.0 healing_power points (2.78 DPS) | yes | Amberglow Talisman (10824, -0.28 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon]; Necklace of Harmony (5180, -1.19 DPS) [world] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Ghostly Mantle (3324, -0.67 DPS) [quest]; Sheepshear Mantle (13115, -0.93 DPS, sim-verified) [world_drop] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 32.6 healing_power points (3.35 DPS) | yes | Battle Healer's Cloak (19528, +0.00 DPS, sim-verified) [rep]; Glowing Thresher Cape (6901, -0.57 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.75 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.1 healing_power points (8.03 DPS) | yes | Dreamweave Vest (10021, -3.16 DPS) [crafted]; Civinad Robes (9623, -3.18 DPS) [quest]; Doomsayer's Robe (4746, -4.91 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.5 healing_power points (3.14 DPS) | yes | Mindthrust Bracers (1974, -0.95 DPS) [dungeon]; Earthen Silk Cuffs (254019, -0.95 DPS, sim-verified) [crafted]; Dryad's Wrist Bindings (19597, -1.01 DPS) [pvp] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (145.5 DPS) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Mender's Leather Gloves (252530, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, -2.19 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.9 healing_power points (4.92 DPS) | yes | Mender's Leather Belt (252523, -0.95 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (145.5 DPS) | yes | Wisdom's Leather Leggings (252519, -0.47 DPS) [crafted]; Pristine Leggings (253987, -0.89 DPS) [crafted]; Stoneweaver Leggings (9407, -3.17 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 51.0 healing_power points (5.25 DPS) | yes | Furen's Boots (13100, -0.70 DPS, sim-verified) [world_drop]; Everlast Boots (10359, -0.86 DPS) [quest]; Thoughtcast Boots (10578, -1.07 DPS) [dungeon] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.0 healing_power points (2.98 DPS) | yes | Welken Ring (5011, -0.48 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.56 DPS) [dungeon]; The Queen's Jewel (13094, -0.64 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 27.0 healing_power points (2.78 DPS) | yes | Electrocutioner Lagnut (9447, -0.36 DPS) [dungeon]; The Queen's Jewel (13094, -0.44 DPS) [world_drop]; Welken Ring (5011, -0.48 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (145.5 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (145.5 DPS) | yes | Hand of Righteousness (7721, -0.39 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -2.61 DPS) [dungeon]; Manual Crowd Pummeler (9449, -16.59 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life

No-known-source sample (15 of 431, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 198.3. Weights run: 2.1s. Verify run: 10.1s. 562 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.534 ± 0.024, spirit=2.780 ± 0.024, mp5=4.433 ± 0.022, crit=0.074 ± 0.011 per rating point (14 rating = 1%, 1.036 per %), spell_haste=1.008 ± 0.239

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 85.1 healing_power points (8.41 DPS) | yes | Blood Guard's Restored Leather Helm (220875, -0.50 DPS) [vendor]; Engineer's Guild Headpiece (9534, -1.31 DPS, sim-verified) [quest]; Cassandra's Grace (13102, -1.59 DPS) [world_drop] |
| neck | Darkmoon Necklace (19303) | Lhara [vendor] | 35.8 healing_power points (3.54 DPS) | yes | Necklace of Calisea (1714, -0.55 DPS) [dungeon]; Choker of the High Shaman (4112, -0.58 DPS) [quest]; Horizon Choker (13085, -0.84 DPS, sim-verified) [world_drop] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 67.1 healing_power points (6.64 DPS) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.20 DPS) [vendor]; Mender's Leather Shoulder (252538, -0.74 DPS, sim-verified) [crafted]; Nethergeld Shoulders (254049, -1.21 DPS) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | 38.7 healing_power points (3.82 DPS) | yes | Wingveil Cloak (10802, -0.31 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -0.33 DPS) [vendor]; Cloak of Blight (6832, -0.51 DPS, sim-verified) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 91.1 healing_power points (9.00 DPS) | yes | Doomsayer's Robe (4746, -0.20 DPS, sim-verified) [quest]; Feathered Breastplate (8349, -0.89 DPS) [crafted]; Stormcloth Vest (10020, -1.07 DPS) [crafted] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.4 healing_power points (4.59 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Enchanted Kodo Bracers (13119, -1.23 DPS) [world_drop]; Venomshroud Armguards (14439, -1.78 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.9 healing_power points (8.99 DPS) | yes | Mender's Leather Gauntlets (252551, -2.26 DPS, sim-verified) [crafted]; Gilded Gloves (254095, -2.41 DPS) [crafted]; Earthenweave Gloves (254075, -3.39 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.6 healing_power points (5.59 DPS) | yes | Gilded Cord (254037, -0.70 DPS, sim-verified) [crafted]; Mender's Leather Waistguard (252477, -0.81 DPS) [crafted]; Mender's Leather Belt (252523, -1.16 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 85.8 healing_power points (8.48 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.16 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.21 DPS) [dungeon]; Dalewind Trousers (13008, -0.24 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.9 healing_power points (6.71 DPS) | yes | First Sergeant's Restored Leather Boots (220885, +0.00 DPS, sim-verified) [vendor]; Earthenweave Boots (254093, -1.12 DPS) [crafted]; Furen's Boots (13100, -1.25 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.6 healing_power points (4.21 DPS) | yes | Snake Hoop (6750, -1.22 DPS) [quest]; Coldwater Ring (4550, -1.28 DPS) [quest]; Welken Ring (5011, -1.53 DPS) [world_drop] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | 31.0 healing_power points (3.07 DPS) | yes | Coldwater Ring (4550, -0.14 DPS) [quest]; Welken Ring (5011, -0.39 DPS) [world_drop]; Snake Hoop (6750, -0.51 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (198.3 DPS) | yes | Evonice's Landin' Pilla (18951, -3.82 DPS) [quest]; Uther's Strength (11302, -4.79 DPS) [world_drop]; Alchemist's Stone (13503, -5.47 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (198.3 DPS) | yes | Uther's Strength (11302, -0.46 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -0.55 DPS) [quest]; Alchemist's Stone (13503, -2.20 DPS) [crafted] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (198.3 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -1.84 DPS) [dungeon]; Hand of Righteousness (7721, -2.52 DPS) [dungeon] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | 41.2 healing_power points (4.08 DPS) | yes | Hibernal Sphere (15937, -0.88 DPS, sim-verified) [world_drop]; Thermotastic Egg Timer (9644, -1.05 DPS) [quest]; Windchaser Orb (15965, -1.05 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Darkmoon Necklace; shoulder: Living Shoulders; back: Battle Healer's Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; main_hand: Charstone Dirk; off_hand: Twisting Essence Jar

No-known-source sample (15 of 562, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 339.5. Weights run: 2.2s. Verify run: 20.0s. 1218 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.045, intellect=2.354 ± 0.064, spirit=4.397 ± 0.050, mp5=7.029 ± 0.048, crit=0.099 ± 0.014 per rating point (14 rating = 1%, 1.383 per %), spell_haste=not significant (0.397 ± 0.651)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 176.9 healing_power points (17.44 DPS) | yes | Champion's Dragonhide Headdress (227205, -2.59 DPS) [vendor]; Feralheart Headdress (226786, -3.17 DPS) [vendor]; Wildheart Cowl (16720, -5.55 DPS, sim-verified) [dungeon] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 94.4 healing_power points (9.30 DPS) | yes | Enchanted South Seas Kelp (19612, -3.08 DPS) [quest]; Pristine Enchanted South Seas Kelp (19613, -3.08 DPS) [quest]; Lady Alizabeth's Pendant (13002, -4.13 DPS, sim-verified) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 145.1 healing_power points (14.31 DPS) | yes | Feralheart Mantle (226785, -0.77 DPS, sim-verified) [vendor]; Champion's Dragonhide Pauldrons (227207, -2.17 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, -3.16 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 65.5 healing_power points (6.46 DPS) | yes | Cloak of Earth and Sky (279270, -0.39 DPS) [crafted]; Battle Healer's Cloak (19526, -0.43 DPS) [rep]; Darkspear Raider's Cloak (272063, -2.39 DPS, sim-verified) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 156.6 healing_power points (15.44 DPS) | yes | Ironfeather Breastplate (15066, -0.59 DPS, sim-verified) [crafted]; Wildheart Vest (16706, -1.05 DPS) [dungeon]; Ghostweave Vest (14141, -1.23 DPS) [crafted] |
| wrist | Feralheart Bindings (226782) | Mokvar [vendor] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Devout Bracers (16697, -1.36 DPS) [dungeon]; Bracers of Hope (22667, -6.24 DPS, sim-verified) [quest] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 129.0 healing_power points (12.72 DPS) | yes | Wildheart Gloves (16717, -1.17 DPS, sim-verified) [dungeon]; Devout Gloves (16692, -1.55 DPS) [dungeon]; Swarmtender's Gloves (275607, -2.47 DPS) [crafted] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Elderwild Waistcord (279252, -0.82 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.37 DPS) [crafted]; Caretaker's Cord (272398, -4.67 DPS, sim-verified) [vendor] |
| legs | Feralheart Pants (226787) | Mokvar [vendor] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; Elderwild Pants (279254, -0.22 DPS) [crafted]; Devout Skirt (16694, -6.03 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 152.6 healing_power points (15.04 DPS) | yes | Incandescent Mooncloth Boots (227862, -1.39 DPS) [vendor]; Mooncloth Boots (15802, -3.79 DPS) [crafted]; Devout Sandals (16691, -4.27 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 91.3 healing_power points (9.00 DPS) | yes | Band of the Hierophant (13096, -1.92 DPS) [world_drop]; Channeler's Ring (272406, -2.29 DPS) [vendor]; Eye of Adaegus (5266, -2.41 DPS) [world_drop] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 73.7 healing_power points (7.26 DPS) | yes | Channeler's Ring (272406, -0.54 DPS) [vendor]; Eye of Adaegus (5266, -0.67 DPS) [world_drop]; Band of the Hierophant (13096, -0.86 DPS, sim-verified) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+10.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -5.19 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.06 DPS) [quest]; Shard of the Splithooves (10659, -6.24 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Ankh of Life (1713, -0.40 DPS, sim-verified) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, -0.04 DPS) [world_drop]; Wind Spirit Staff (6689, -0.28 DPS) [dungeon] |
| off_hand | Lapidis Tankard of Tidesippe (4696) | World drop [world_drop] | 82.1 healing_power points (8.09 DPS) | yes | Lei of the Lifegiver (19312, +0.00 DPS, sim-verified) [rep]; High Warlord's Tome of Mending (234564, -1.99 DPS) [pvp]; Arcane Star (15931, -2.31 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Feralheart Bindings; waist: Feralheart Cord; legs: Feralheart Pants; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lapidis Tankard of Tidesippe

No-known-source sample (15 of 1218, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 639.2. Weights run: 1.8s. Verify run: 20.4s. 1218 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.324, intellect=2.826 ± 0.130, spirit=2.735 ± 0.120, mp5=3.745 ± 0.156, crit=0.317 ± 0.049 per rating point (14 rating = 1%, 4.436 per %), spell_haste=not significant (-1.810 ± 2.036)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 162.1 healing_power points (23.63 DPS) | yes | Champion's Dragonhide Headdress (227205, -4.00 DPS) [vendor]; Devout Crown (16693, -4.85 DPS) [dungeon]; Feralheart Headdress (226786, -8.87 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 70.9 healing_power points (10.33 DPS) | yes | Enchanted South Seas Kelp (19612, -2.63 DPS) [quest]; Pristine Enchanted South Seas Kelp (19613, -2.63 DPS) [quest]; Drake Tooth Necklace (21531, -4.52 DPS, sim-verified) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 108.8 healing_power points (15.86 DPS) | yes | Champion's Dragonhide Pauldrons (227207, -0.41 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, -0.43 DPS) [vendor]; Devout Mantle (16695, -11.26 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 70.3 healing_power points (10.24 DPS) | yes | Darkspear Raider's Cloak (272076, -2.48 DPS) [vendor]; Battle Healer's Cloak (19526, -3.26 DPS) [rep] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 134.5 healing_power points (19.60 DPS) | yes | Warlord's Dragonhide Tunic (231674, -1.56 DPS) [vendor]; Feralheart Embrace (226783, -1.64 DPS) [vendor]; Devout Robe (16690, -3.30 DPS, sim-verified) [dungeon] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 81.9 healing_power points (11.94 DPS) | yes | General's Dragonhide Bracers (16553, -0.13 DPS) [pvp]; Feralheart Bindings (226782, -1.34 DPS, sim-verified) [vendor]; Devout Bracers (16697, -3.05 DPS) [dungeon] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 104.3 healing_power points (15.20 DPS) | yes | Wildheart Gloves (16717, -1.52 DPS) [dungeon]; Devout Gloves (16692, -2.12 DPS) [dungeon]; Raider Handwraps (272097, -12.96 DPS, sim-verified) [vendor] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | 97.0 healing_power points (14.14 DPS) | yes | General's Dragonhide Belt (16556, +0.00 DPS) [pvp]; Devout Belt (16696, -0.72 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -9.35 DPS, sim-verified) [crafted] |
| legs | Feralheart Pants (226787) | Mokvar [vendor] | sim-verified (639.3 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, +0.00 DPS) [vendor]; General's Dragonhide Legguards (231673, -1.89 DPS) [vendor]; Devout Skirt (16694, -9.38 DPS, sim-verified) [dungeon] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | sim-verified (639.3 DPS) | yes | Mooncloth Boots (15802, -2.89 DPS) [crafted]; Devout Sandals (16691, -4.83 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -7.26 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | 79.2 healing_power points (11.54 DPS) | yes | Channeler's Ring (272406, -2.58 DPS) [vendor]; Band of the Hierophant (13096, -3.03 DPS) [world_drop]; Blessed Band of Light (272407, -3.83 DPS) [vendor] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | 65.2 healing_power points (9.51 DPS) | yes | Band of the Hierophant (13096, -1.01 DPS) [world_drop]; Blessed Band of Light (272407, -1.80 DPS) [vendor]; Channeler's Ring (272406, -3.72 DPS, sim-verified) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (639.3 DPS) | yes | Ankh of Life (1713, -3.40 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.20 DPS) [quest]; Shard of the Splithooves (10659, -4.91 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (639.3 DPS) | yes | Ankh of Life (1713, -1.34 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.14 DPS) [quest]; Darkmoon Card: Blue Dragon (19288, -5.22 DPS, sim-verified) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-verified (639.3 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Staff of Hale Magefire (13000, -3.46 DPS) [world_drop]; Wind Spirit Staff (6689, -3.74 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | 58.7 healing_power points (8.56 DPS) | yes | High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Arcane Star (15931, -1.25 DPS) [world_drop]; Lapidis Tankard of Tidesippe (4696, -6.26 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Bracers of Hope; waist: Feralheart Cord; legs: Feralheart Pants; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Serenity Field; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1218, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

