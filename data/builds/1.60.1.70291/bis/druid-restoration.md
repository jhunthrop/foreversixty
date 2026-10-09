# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 48.2. Weights run: 3.6s. Verify run: 11.8s. 194 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.289 ± 0.004, spirit=1.552 ± 0.008, mp5=3.307 ± 0.015, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.609 per %), spell_haste=-5.588 ± 0.106

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.23 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 6.2 healing_power points (0.64 DPS) | yes | Scholarly Pendant (277203, -1.34 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.6 healing_power points (1.20 DPS) | yes | Forest Leather Mantle (4709, -0.56 DPS) [world_drop]; Prospector's Pads (14566, -0.56 DPS) [world_drop]; Slime-encrusted Pads (6461, -0.82 DPS, sim-verified) [dungeon] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | 12.1 healing_power points (1.25 DPS) | yes | Spirit Cloak (4792, -0.61 DPS) [vendor]; Sylvan Cloak (4793, -0.61 DPS) [vendor]; Regent's Cloak (5969, -2.21 DPS, sim-verified) [world] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 26.1 healing_power points (2.69 DPS) | yes | Filigreed Pristine Gown (253901, -0.21 DPS) [crafted]; Corsair's Overshirt (5202, -0.28 DPS) [dungeon]; Robe of the Moccasin (6465, -1.02 DPS, sim-verified) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 16.4 healing_power points (1.69 DPS) | yes | Owl Bracers (4796, -1.02 DPS) [vendor]; Bravo's Armbands (270015, -1.05 DPS) [quest]; Drakewing Bands (12999, -2.58 DPS, sim-verified) [world_drop] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 15.2 healing_power points (1.56 DPS) | yes | Magefist Gloves (12977, -0.26 DPS) [world_drop]; Bright Gloves (3066, -0.39 DPS) [world_drop]; Pristine Gloves (253913, -1.53 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 16.2 healing_power points (1.66 DPS) | yes | Wisdom's Leather Belt (252433, -0.31 DPS) [crafted]; Keller's Girdle (2911, -0.60 DPS) [world_drop]; Novice Ardent's Sash (253887, -1.74 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.50 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 17.4 healing_power points (1.80 DPS) | yes | Pristine Boots (253889, -0.47 DPS) [crafted]; Kimbra Boots (6191, -0.79 DPS) [quest]; Black Whelp Slippers (252424, -1.92 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.9 healing_power points (1.23 DPS) | yes | Deep Fathom Ring (6463, -0.43 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.54 DPS) [rep] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Lorekeeper's Ring (20431, -0.28 DPS) [rep]; Deep Fathom Ring (6463, -2.12 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 15.8 healing_power points (1.62 DPS) | yes | Twisted Chanter's Staff (890, -0.30 DPS) [world_drop]; Gnarled Hermit's Staff (1539, -0.50 DPS) [world]; Staff of the Blessed Seer (2271, -1.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Black Pearl Ring; finger2: Band of Purification; main_hand: Staff of Westfall

No-known-source sample (15 of 194, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 100.2. Weights run: 5.0s. Verify run: 15.6s. 350 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.454 ± 0.014, spirit=2.917 ± 0.018, mp5=4.441 ± 0.027, crit=0.052 ± 0.005 per rating point (14 rating = 1%, 0.735 per %), spell_haste=not significant (-1.023 ± 0.259)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embalmed Shroud (7691) | Scarlet Monastery: Fallen Champion [dungeon] | 51.0 healing_power points (5.19 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Whisperwind Headdress (6688, -0.30 DPS) [dungeon]; Enduring Cap (3020, -0.75 DPS) [world_drop] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 17.5 healing_power points (1.78 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.30 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 30.6 healing_power points (3.12 DPS) | yes | Batwing Mantle (6697, +0.00 DPS, sim-verified) [dungeon]; Faerie Mantle (5820, -0.45 DPS) [quest]; Nightsky Mantle (4718, -0.45 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 32.2 healing_power points (3.28 DPS) | yes | Glowing Thresher Cape (6901, -0.12 DPS, sim-verified) [dungeon]; Prelacy Cape (7004, -0.38 DPS) [quest]; Caretaker's Cape (19533, -0.77 DPS) [rep] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.2 healing_power points (4.50 DPS) | yes | Pristine Gown (253961, -0.24 DPS) [crafted]; Pressed Felt Robe (1997, -0.49 DPS) [world]; Robe of the Moccasin (6465, -0.62 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.6 healing_power points (2.10 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop]; Drakewing Bands (12999, -0.69 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.3 healing_power points (4.30 DPS) | yes | Zodiac Gloves (7106, -1.63 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.63 DPS) [dungeon]; Shilly Mitts (9609, -1.65 DPS, sim-verified) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.2 healing_power points (4.60 DPS) | yes | Dokebi Cord (14578, -2.23 DPS) [world_drop]; Resilient Cord (14406, -2.53 DPS) [world_drop]; Silver-lined Belt (13011, -2.91 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.6 healing_power points (5.25 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Earthen Leggings (253999, -0.81 DPS, sim-verified) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.26 DPS) | yes | Glinteye Slippers (273024, -1.59 DPS) [dungeon]; Boots of the Enchanter (4325, -1.88 DPS) [crafted]; Soggy Boots (274747, -2.09 DPS, sim-verified) [vendor] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.6 healing_power points (3.12 DPS) | yes | The Queen's Jewel (13094, -0.44 DPS) [world_drop]; Monkey Ring (6748, -1.04 DPS) [quest]; Ring of Calm (6790, -1.04 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.18 DPS, sim-verified) [world_drop]; Monkey Ring (6748, -0.61 DPS) [quest]; Ring of Calm (6790, -0.61 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (100.2 DPS) | yes | Royal Diplomatic Scepter (9457, -3.01 DPS) [dungeon]; Death Speaker Scepter (2816, -5.80 DPS) [dungeon]; Manual Crowd Pummeler (9449, -15.36 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Embalmed Shroud; neck: Pendant of Myzrael; shoulder: Mantle of Honor; back: Repairman's Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Electrocutioner Lagnut; main_hand: Wind Spirit Staff

No-known-source sample (15 of 350, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 149.6. Weights run: 5.5s. Verify run: 35.0s. 461 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.347 ± 0.016, spirit=2.509 ± 0.011, mp5=4.854 ± 0.034, crit=0.056 ± 0.005 per rating point (14 rating = 1%, 0.782 per %), spell_haste=not significant (0.538 ± 0.314)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 67.5 healing_power points (6.93 DPS) | yes | Holy Shroud (2721, -1.99 DPS) [world_drop]; Miner's Hat of the Deep (9429, -2.00 DPS) [dungeon]; Whitemane's Chapeau (7720, -2.84 DPS, sim-verified) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-verified (149.6 DPS) | yes | Triune Amulet (7722, +0.00 DPS) [dungeon]; Amberglow Talisman (10824, -0.20 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Crimson Silk Shoulders (7059, -1.00 DPS) [crafted]; Sheepshear Mantle (13115, -2.44 DPS, sim-verified) [world_drop] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | 30.5 healing_power points (3.14 DPS) | yes | Glowing Thresher Cape (6901, -0.36 DPS) [dungeon]; Prelacy Cape (7004, -0.46 DPS) [quest]; Repairman's Cape (9605, -1.34 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.1 healing_power points (8.02 DPS) | yes | Dreamweave Vest (10021, -3.17 DPS) [crafted]; Civinad Robes (9623, -3.18 DPS) [quest]; Doomsayer's Robe (4746, -5.74 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.5 healing_power points (3.13 DPS) | yes | Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Earthen Silk Cuffs (254019, -1.14 DPS, sim-verified) [crafted]; Silkstream Cuffs (16791, -1.41 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (149.6 DPS) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Restorer's Fine Gloves (270059, +0.00 DPS) [quest]; Earthen Silk Gloves (254017, -2.20 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.8 healing_power points (4.91 DPS) | yes | Mender's Leather Belt (252523, -0.90 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (149.6 DPS) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Wisdom's Leather Leggings (252519, -0.48 DPS) [crafted]; Warchief Kilt (7760, -4.31 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 51.0 healing_power points (5.24 DPS) | yes | Everlast Boots (10359, -0.86 DPS) [quest]; Thoughtcast Boots (10578, -1.07 DPS) [dungeon]; Furen's Boots (13100, -1.93 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.1 healing_power points (2.99 DPS) | yes | Welken Ring (5011, -0.50 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.57 DPS) [dungeon]; The Queen's Jewel (13094, -0.65 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 27.0 healing_power points (2.77 DPS) | yes | Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; Welken Ring (5011, -0.43 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.43 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (149.6 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (149.6 DPS) | yes | Hand of Righteousness (7721, -0.39 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -2.60 DPS) [dungeon]; Manual Crowd Pummeler (9449, -17.19 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Caretaker's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life

No-known-source sample (15 of 461, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 203.2. Weights run: 5.6s. Verify run: 26.0s. 604 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.528 ± 0.017, spirit=2.753 ± 0.017, mp5=4.398 ± 0.016, crit=0.072 ± 0.007 per rating point (14 rating = 1%, 1.009 per %), spell_haste=0.983 ± 0.178

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 84.5 healing_power points (8.38 DPS) | yes | Knight-Lieutenant's Restored Leather Helm (220874, -0.48 DPS) [vendor]; Engineer's Guild Headpiece (9534, -0.49 DPS) [quest]; Papal Fez (9431, -1.17 DPS) [dungeon] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | sim-verified (203.2 DPS) | yes | Darkmoon Necklace (19303, -0.57 DPS) [vendor]; Horizon Choker (13085, -0.88 DPS) [world_drop]; Gemshard Heart (17707, -0.94 DPS) [dungeon] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 66.8 healing_power points (6.62 DPS) | yes | Knight-Lieutenant's Restored Leather Spaulders (220876, -0.20 DPS) [vendor]; Mender's Leather Shoulder (252538, -0.97 DPS) [crafted]; Nethergeld Shoulders (254049, -1.19 DPS) [crafted] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Wingveil Cloak (10802, -0.32 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -0.33 DPS) [vendor] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-verified (203.2 DPS) | yes | Vestments of the Atal'ai Prophet (10806, -1.73 DPS) [dungeon]; Ghostweave Vest (14141, -1.82 DPS) [crafted]; Forest's Embrace (22272, -2.19 DPS) [quest] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.58 DPS) | yes | Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Aristocratic Cuffs (12546, -0.67 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.4 healing_power points (8.96 DPS) | yes | Mender's Leather Gauntlets (252551, -2.07 DPS) [crafted]; Gilded Gloves (254095, -2.39 DPS) [crafted]; Earthenweave Gloves (254075, -3.38 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.3 healing_power points (5.58 DPS) | yes | Gilded Cord (254037, -0.55 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.79 DPS) [crafted]; Mender's Leather Belt (252523, -1.15 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 85.2 healing_power points (8.45 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.16 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.22 DPS) [dungeon]; Dalewind Trousers (13008, -0.27 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.3 healing_power points (6.67 DPS) | yes | Sergeant Major's Restored Leather Boots (220884, -0.44 DPS) [vendor]; Mistwalker Boots (10629, -0.55 DPS) [dungeon]; Earthenweave Boots (254093, -1.09 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.2 healing_power points (4.18 DPS) | yes | Choking Band (11868, -0.64 DPS) [quest]; Darkspear Signet (272069, -1.13 DPS) [vendor]; Cyclopean Band (11824, -1.14 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (203.2 DPS) | yes | Choking Band (11868, -0.09 DPS) [quest]; Darkspear Signet (272069, -0.58 DPS) [vendor]; Cyclopean Band (11824, -0.59 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -3.81 DPS) [quest]; Thunderbrew's Boot Flask (744, -4.36 DPS) [quest]; Uther's Strength (11302, -4.76 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.55 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.09 DPS) [quest]; Uther's Strength (11302, -1.49 DPS) [world_drop] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wind Spirit Staff (6689, -1.88 DPS) [dungeon]; Hand of Righteousness (7721, -2.54 DPS) [dungeon]; Spire of Hakkar (10844, -3.53 DPS) [world] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Desertwalker Cane (12471, +0.00 DPS) [dungeon]; Cloud Stone (17737, +0.00 DPS) [dungeon]; Enthralled Sphere (11625, -0.58 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Lei of Lilies; shoulder: Living Shoulders; back: Caretaker's Cape; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Charstone Dirk; off_hand: Twisting Essence Jar

No-known-source sample (15 of 604, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 344.8. Weights run: 5.6s. Verify run: 69.1s. 1450 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 173.1 healing_power points (17.46 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -2.67 DPS) [vendor]; Wildheart Cowl (16720, -2.81 DPS) [dungeon]; Feralheart Headdress (226786, -3.26 DPS) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 92.3 healing_power points (9.31 DPS) | yes | Lady Maye's Pendant (14558, -0.72 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.94 DPS) [world_drop]; Heart of the Fiend (13960, -1.72 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 143.0 healing_power points (14.42 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, -2.32 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -3.38 DPS) [vendor]; Feralheart Mantle (226785, -4.10 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Butcher's Apron (12608, +0.00 DPS) [dungeon]; Frostweaver Cape (12968, +0.00 DPS) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Alanna's Embrace (13314, -0.17 DPS) [dungeon]; Ironfeather Breastplate (15066, -0.54 DPS) [crafted] |
| wrist | Feralheart Bindings (226782) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bleak Howler Armguards (13208, +0.00 DPS) [dungeon]; Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 125.5 healing_power points (12.66 DPS) | yes | Wildheart Gloves (16717, -0.47 DPS) [dungeon]; Hands of the Exalted Herald (12554, -1.23 DPS) [dungeon]; Devout Gloves (16692, -1.56 DPS) [dungeon] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cord (272398, +0.00 DPS) [vendor]; Elderwild Waistcord (279252, -0.72 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.34 DPS) [crafted] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | sim-verified (344.7 DPS) | yes | Devout Skirt (16694, -3.54 DPS) [dungeon]; Haunting Specter Leggings (11929, -4.48 DPS) [dungeon]; Knight-Captain's Dragonhide Legguards (227200, -4.55 DPS) [vendor] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 148.3 healing_power points (14.96 DPS) | yes | Incandescent Mooncloth Boots (227862, -1.33 DPS) [vendor]; Mooncloth Boots (15802, -3.78 DPS) [crafted]; Devout Sandals (16691, -4.26 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -1.01 DPS) [dungeon]; Emerald Flame Ring (18395, -1.35 DPS) [dungeon]; Band of Mending (22334, -1.73 DPS) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, +0.00 DPS) [dungeon]; Emerald Flame Ring (18395, +0.00 DPS) [dungeon]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -2.77 DPS) [dungeon]; Ankh of Life (1713, -5.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.07 DPS) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-verified (344.7 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, -0.12 DPS) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dancing Sliver (15854, -0.35 DPS) [quest]; Wind Spirit Staff (6689, -0.37 DPS) [dungeon]; Staff of Hale Magefire (13000, -0.44 DPS) [world_drop] |
| off_hand | Lapidis Tankard of Tidesippe (4696) | World drop [world_drop] | 79.6 healing_power points (8.03 DPS) | yes | Book of the Dead (13353, -0.33 DPS) [dungeon]; Thaurissan's Royal Scepter (11928, -0.37 DPS) [dungeon]; Lei of the Lifegiver (19312, -0.53 DPS) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Feralheart Bindings; waist: Feralheart Cord; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Royal Seal of Eldre'Thalas; off_hand: Lapidis Tankard of Tidesippe

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 643.2. Weights run: 4.2s. Verify run: 45.9s. 1450 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.204, intellect=2.451 ± 0.080, spirit=2.281 ± 0.075, mp5=3.336 ± 0.099, crit=0.327 ± 0.030 per rating point (14 rating = 1%, 4.573 per %), spell_haste=not significant (-0.743 ± 1.216)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 149.0 healing_power points (24.64 DPS) | yes | Lieutenant Commander's Dragonhide Headdress (227199, -4.67 DPS) [vendor]; Sanctified Leather Helm (22689, -5.80 DPS) [quest]; Feralheart Headdress (226786, -8.25 DPS, sim-verified) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Drake Tooth Necklace (21531, -1.65 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 101.8 healing_power points (16.83 DPS) | yes | Lieutenant Commander's Dragonhide Pauldrons (227201, -1.16 DPS) [vendor]; Field Marshal's Dragonhide Pauldrons (231705, -1.32 DPS) [vendor]; Devout Mantle (16695, -11.67 DPS, sim-verified) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 66.5 healing_power points (10.99 DPS) | yes | Cloak of the Cosmos (18389, -2.24 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -2.25 DPS) [vendor]; Frostweaver Cape (12968, -5.35 DPS, sim-verified) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Alanna's Embrace (13314, -0.61 DPS) [dungeon]; Devout Robe (16690, -0.88 DPS) [dungeon]; Tunic of Undead Slaying (23089, -20.61 DPS, sim-verified) [world] |
| wrist | Feralheart Bindings (226782) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Bracers (16445, +0.00 DPS) [pvp]; Bracers of Mending (23129, +0.00 DPS) [dungeon]; Bracers of Hope (22667, -7.51 DPS, sim-verified) [quest] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Handwraps (272097, -0.31 DPS) [vendor]; Wildheart Gloves (16717, -1.83 DPS) [dungeon]; Hands of the Exalted Herald (12554, -8.27 DPS, sim-verified) [dungeon] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | 85.9 healing_power points (14.20 DPS) | yes | Marshal's Dragonhide Waistguard (16447, -0.30 DPS) [pvp]; Devout Belt (16696, -0.89 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -8.29 DPS, sim-verified) [crafted] |
| legs | Knight-Captain's Dragonhide Legguards (227200) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of Arcana (12756, +0.00 DPS) [quest]; Devout Skirt (16694, -1.03 DPS) [dungeon]; Feralheart Pants (226787, -1.35 DPS) [vendor] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, -2.68 DPS) [crafted]; Faith Healer's Boots (22247, -3.53 DPS) [dungeon]; Incandescent Mooncloth Boots (227862, -6.34 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, -1.83 DPS) [dungeon]; Band of Mending (22334, -1.97 DPS) [dungeon]; Naglering (11669, -14.37 DPS, sim-verified) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, +0.00 DPS) [dungeon]; Band of Mending (22334, +0.00 DPS) [dungeon]; Seal of Rivendare (13345, -0.37 DPS) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-verified (+11.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -0.33 DPS) [vendor]; Mindtap Talisman (18371, -1.21 DPS) [dungeon]; Briarwood Reed (12930, -2.48 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.33 DPS) [vendor]; Mindtap Talisman (18371, -2.21 DPS) [dungeon]; Briarwood Reed (12930, -3.48 DPS) [dungeon] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, +0.00 DPS) [world_drop]; Wind Spirit Staff (6689, -4.94 DPS) [dungeon]; Hand of Righteousness (7721, -6.18 DPS) [dungeon] |
| off_hand | Lei of the Lifegiver (19312) | Stormpike Guard [rep] | sim-verified (643.3 DPS) | yes | Thaurissan's Royal Scepter (11928, +0.00 DPS) [dungeon]; Book of the Dead (13353, +0.00 DPS) [dungeon]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Feralheart Bindings; waist: Feralheart Cord; legs: Knight-Captain's Dragonhide Legguards; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-00000000000000000000-5050010000000000)

Set DPS (verified): 43.9. Weights run: 3.6s. Verify run: 22.6s. 184 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.289 ± 0.004, spirit=1.552 ± 0.008, mp5=3.307 ± 0.015, crit=0.044 ± 0.003 per rating point (14 rating = 1%, 0.609 per %), spell_haste=-5.588 ± 0.106

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (1.85 DPS) | yes | Wisdom's Leather Hood (252507, -0.13 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.23 DPS) [crafted]; Stormrider's Leather Hood (252506, -0.42 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.7 healing_power points (0.48 DPS) | yes | Scholarly Pendant (277203, -2.00 DPS, sim-verified) [quest] |
| shoulder | Slime-encrusted Pads (6461) | Wailing Caverns: Mutanus the Devourer [dungeon] | sim-verified (43.9 DPS) | yes | Forest Leather Mantle (4709, -0.38 DPS) [world_drop]; Prospector's Pads (14566, -0.38 DPS) [world_drop]; Magician's Mantle (12998, -1.17 DPS, sim-verified) [world_drop] |
| back | Regent's Cloak (5969) | Ravenclaw Regent [world] | sim-verified (43.9 DPS) | yes | Battle Healer's Cloak (20427, +0.00 DPS, sim-verified) [rep]; Traveler's Shawl (277289, +0.00 DPS) [quest]; Spirit Cloak (4792, -0.16 DPS) [vendor] |
| chest | Armor of the Fang (6473) | Wailing Caverns: Lord Pythas [dungeon] | sim-verified (43.9 DPS) | yes | Wisdom's Leather Armor (252493, +0.00 DPS) [crafted]; Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Robe of the Moccasin (6465, -2.14 DPS, sim-verified) [dungeon] |
| wrist | Drakewing Bands (12999) | World drop [world_drop] | sim-verified (43.9 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Crystalline Cuffs (14148, -0.21 DPS) [dungeon]; Owl Bracers (4796, -0.30 DPS) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | sim-verified (43.9 DPS) | yes | Pristine Gloves (253913, -0.03 DPS) [crafted]; Magefist Gloves (12977, -0.26 DPS) [world_drop]; Blight Gloves (279877, -2.22 DPS, sim-verified) [quest] |
| waist | Novice Ardent's Sash (253887) | Tailoring [crafted] | sim-verified (43.9 DPS) | yes | Wisdom's Leather Belt (252433, -0.14 DPS) [crafted]; Keller's Girdle (2911, -0.43 DPS) [world_drop]; Pristine Sash (253925, -1.36 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 33.9 healing_power points (3.50 DPS) | yes | Filigreed Pristine Leggings (253937, -0.27 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.61 DPS) [world_drop]; Scarecrow Trousers (4434, -2.03 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (43.9 DPS) | yes | Wisdom's Leather Boots (252444, +0.00 DPS) [crafted]; Pristine Boots (253889, +0.00 DPS) [crafted]; Black Whelp Slippers (252424, -2.20 DPS, sim-verified) [crafted] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 11.9 healing_power points (1.23 DPS) | yes | Deep Fathom Ring (6463, -0.43 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; Advisor's Ring (20426, -0.54 DPS) [rep] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 9.3 healing_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Advisor's Ring (20426, -0.28 DPS) [rep]; Deep Fathom Ring (6463, -2.00 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | sim-verified (43.9 DPS) | yes | Advisor's Gnarled Staff (20425, -0.05 DPS) [pvp]; Gnarled Necromancer's Staff (251534, -0.27 DPS) [quest]; Staff of Orgrimmar (15444, -1.94 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Slime-encrusted Pads; back: Regent's Cloak; chest: Armor of the Fang; wrist: Drakewing Bands; hands: Wisdom's Leather Gloves; waist: Novice Ardent's Sash; legs: Wisdom's Leather Pants; feet: Footpads of the Fang; finger1: Black Pearl Ring; finger2: Band of Purification; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 184, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-00000000000000000000-5050035110010000)

Set DPS (verified): 101.0. Weights run: 5.0s. Verify run: 16.4s. 342 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.001, intellect=1.454 ± 0.014, spirit=2.917 ± 0.018, mp5=4.441 ± 0.027, crit=0.052 ± 0.005 per rating point (14 rating = 1%, 0.735 per %), spell_haste=not significant (-1.023 ± 0.259)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (101.0 DPS) | yes | Whisperwind Headdress (6688, -0.25 DPS) [dungeon]; Enduring Cap (3020, -0.69 DPS) [world_drop]; Embalmed Shroud (7691, -1.02 DPS, sim-verified) [dungeon] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 17.5 healing_power points (1.78 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.00 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.30 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 35.3 healing_power points (3.59 DPS) | yes | Batwing Mantle (6697, -0.60 DPS, sim-verified) [dungeon]; Nightsky Mantle (4718, -0.92 DPS) [world_drop]; Desert Shoulders (15457, -1.36 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 30.3 healing_power points (3.09 DPS) | yes | Battle Healer's Cloak (19529, +0.00 DPS, sim-verified) [rep]; Amy's Blanket (13005, -1.01 DPS) [world_drop]; Construct Cloak (279848, -1.01 DPS) [quest] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 44.2 healing_power points (4.50 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Pressed Felt Robe (1997, -0.49 DPS) [world]; Robe of the Moccasin (6465, -0.62 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 20.6 healing_power points (2.10 DPS) | yes | Nightsky Wristbands (6407, -0.32 DPS) [world_drop]; Drakewing Bands (12999, -0.44 DPS, sim-verified) [world_drop]; Dokebi Bracers (14580, -0.46 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 42.3 healing_power points (4.30 DPS) | yes | Tattered Mittens (270030, -0.73 DPS, sim-verified) [quest]; Hotshot Pilot's Gloves (9491, -1.63 DPS) [dungeon]; Blight Gloves (279877, -1.78 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 45.2 healing_power points (4.60 DPS) | yes | Dokebi Cord (14578, -2.23 DPS) [world_drop]; Lilac Sash (6780, -2.38 DPS) [quest]; Silver-lined Belt (13011, -2.81 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.6 healing_power points (5.25 DPS) | yes | Pristine Leggings (253987, -0.49 DPS) [crafted]; Earthen Leggings (253999, -0.71 DPS, sim-verified) [crafted]; Stormrider's Leather Kilt (252518, -0.95 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 41.8 healing_power points (4.26 DPS) | yes | Soggy Boots (274747, -1.29 DPS) [vendor]; Glinteye Slippers (273024, -1.59 DPS) [dungeon]; Vorrel's Boots (7751, -2.53 DPS, sim-verified) [quest] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 30.6 healing_power points (3.12 DPS) | yes | The Queen's Jewel (13094, -0.44 DPS) [world_drop]; Monkey Ring (6748, -1.04 DPS) [quest]; Tiger Band (6749, -1.04 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) | Gnomeregan: Electrocutioner 6000 [dungeon] | 26.4 healing_power points (2.69 DPS) | yes | The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Monkey Ring (6748, -0.61 DPS) [quest]; Tiger Band (6749, -0.61 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Diplomatic Scepter (9457, -3.01 DPS) [dungeon]; Death Speaker Scepter (2816, -5.80 DPS) [dungeon]; Manual Crowd Pummeler (9449, -15.24 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; wrist: Mindthrust Bracers; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Electrocutioner Lagnut; main_hand: Wind Spirit Staff

No-known-source sample (15 of 342, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-00000000000000000000-5050035153112000)

Set DPS (verified): 148.8. Weights run: 5.5s. Verify run: 35.5s. 446 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.000, intellect=1.347 ± 0.016, spirit=2.509 ± 0.011, mp5=4.854 ± 0.034, crit=0.056 ± 0.005 per rating point (14 rating = 1%, 0.782 per %), spell_haste=not significant (0.538 ± 0.314)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 67.5 healing_power points (6.93 DPS) | yes | Holy Shroud (2721, -1.99 DPS) [world_drop]; Miner's Hat of the Deep (9429, -2.00 DPS) [dungeon]; Whitemane's Chapeau (7720, -2.40 DPS, sim-verified) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-verified (148.8 DPS) | yes | Triune Amulet (7722, +0.00 DPS) [dungeon]; Amberglow Talisman (10824, -0.20 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon] |
| shoulder | Earthen Silk Shoulders (254033) | Tailoring [crafted] | 38.1 healing_power points (3.91 DPS) | yes | Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Ghostly Mantle (3324, -0.67 DPS) [quest]; Sheepshear Mantle (13115, -1.71 DPS, sim-verified) [world_drop] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-verified (148.8 DPS) | yes | Glowing Thresher Cape (6901, -0.36 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.54 DPS) [quest]; Cloak of Blight (6832, -1.50 DPS, sim-verified) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 78.1 healing_power points (8.02 DPS) | yes | Dreamweave Vest (10021, -3.17 DPS) [crafted]; Civinad Robes (9623, -3.18 DPS) [quest]; Doomsayer's Robe (4746, -5.65 DPS, sim-verified) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 30.5 healing_power points (3.13 DPS) | yes | Earthen Silk Cuffs (254019, -0.23 DPS, sim-verified) [crafted]; Mindthrust Bracers (1974, -0.94 DPS) [dungeon]; Dryad's Wrist Bindings (19597, -1.01 DPS) [pvp] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Old (9395, +0.00 DPS) [world_drop]; Mender's Leather Gloves (252530, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, -1.30 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.8 healing_power points (4.91 DPS) | yes | Mender's Leather Belt (252523, -0.49 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -1.49 DPS) [world_drop]; Sutarn's Ring (13105, -1.98 DPS) [world_drop] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stoneweaver Leggings (9407, +0.00 DPS) [dungeon]; Wisdom's Leather Leggings (252519, -0.48 DPS) [crafted]; Warchief Kilt (7760, -3.48 DPS, sim-verified) [dungeon] |
| feet | Mender's Leather Shoes (252533) | Leatherworking [crafted] | 51.0 healing_power points (5.24 DPS) | yes | Everlast Boots (10359, -0.86 DPS) [quest]; Thoughtcast Boots (10578, -1.07 DPS) [dungeon]; Furen's Boots (13100, -1.86 DPS, sim-verified) [world_drop] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 29.1 healing_power points (2.99 DPS) | yes | Welken Ring (5011, -0.50 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.57 DPS) [dungeon]; The Queen's Jewel (13094, -0.65 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 27.0 healing_power points (2.77 DPS) | yes | Welken Ring (5011, -0.28 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; The Queen's Jewel (13094, -0.43 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Righteousness (7721, -0.39 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -2.60 DPS) [dungeon]; Manual Crowd Pummeler (9449, -16.95 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Earthen Silk Shoulders; back: Battle Healer's Cloak; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Stormcloth Pants; feet: Mender's Leather Shoes; finger1: Darkspear Signet; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life

No-known-source sample (15 of 446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 4300000000000000-00000000000000000000-5050035153113200)

Set DPS (verified): 202.5. Weights run: 5.6s. Verify run: 25.0s. 585 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.003, intellect=1.528 ± 0.017, spirit=2.753 ± 0.017, mp5=4.398 ± 0.016, crit=0.072 ± 0.007 per rating point (14 rating = 1%, 1.009 per %), spell_haste=0.983 ± 0.178

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 84.5 healing_power points (8.38 DPS) | yes | Blood Guard's Restored Leather Helm (220875, -0.48 DPS) [vendor]; Engineer's Guild Headpiece (9534, -0.49 DPS) [quest]; Papal Fez (9431, -1.17 DPS) [dungeon] |
| neck | Lei of Lilies (1315) | World drop [world_drop] | sim-verified (202.5 DPS) | yes | Darkmoon Necklace (19303, -0.57 DPS) [vendor]; Horizon Choker (13085, -0.88 DPS) [world_drop]; Gemshard Heart (17707, -0.94 DPS) [dungeon] |
| shoulder | Living Shoulders (15061) | Leatherworking [crafted] | 66.8 healing_power points (6.62 DPS) | yes | Blood Guard's Restored Leather Spaulders (220877, -0.20 DPS) [vendor]; Mender's Leather Shoulder (252538, -0.97 DPS) [crafted]; Nethergeld Shoulders (254049, -1.19 DPS) [crafted] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Cloak of Blight (6832, -0.27 DPS) [quest]; Wingveil Cloak (10802, -0.32 DPS) [dungeon] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | sim-verified (202.5 DPS) | yes | Vestments of the Atal'ai Prophet (10806, -1.73 DPS) [dungeon]; Ghostweave Vest (14141, -1.82 DPS) [crafted]; Forest's Embrace (22272, -2.19 DPS) [quest] |
| wrist | Mender's Leather Bracers (252543) | Leatherworking [crafted] | 46.2 healing_power points (4.58 DPS) | yes | Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Aristocratic Cuffs (12546, -0.67 DPS) [dungeon]; Enchanted Kodo Bracers (13119, -1.25 DPS) [world_drop] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 90.4 healing_power points (8.96 DPS) | yes | Mender's Leather Gauntlets (252551, -2.07 DPS) [crafted]; Gilded Gloves (254095, -2.39 DPS) [crafted]; Earthenweave Gloves (254075, -3.38 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 56.3 healing_power points (5.58 DPS) | yes | Gilded Cord (254037, -0.55 DPS) [crafted]; Mender's Leather Waistguard (252477, -0.79 DPS) [crafted]; Mender's Leather Belt (252523, -1.15 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 85.2 healing_power points (8.45 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.16 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.22 DPS) [dungeon]; Dalewind Trousers (13008, -0.27 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 67.3 healing_power points (6.67 DPS) | yes | First Sergeant's Restored Leather Boots (220885, -0.44 DPS) [vendor]; Mistwalker Boots (10629, -0.55 DPS) [dungeon]; Earthenweave Boots (254093, -1.09 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 42.2 healing_power points (4.18 DPS) | yes | Darkspear Signet (272069, -1.13 DPS) [vendor]; Cyclopean Band (11824, -1.14 DPS) [dungeon]; Snake Hoop (6750, -1.21 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (202.5 DPS) | yes | Darkspear Signet (272069, -0.58 DPS) [vendor]; Cyclopean Band (11824, -0.59 DPS) [dungeon]; Snake Hoop (6750, -0.67 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -3.81 DPS) [quest]; Uther's Strength (11302, -4.76 DPS) [world_drop]; Alchemist's Stone (13503, -5.45 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.55 DPS) [quest]; Uther's Strength (11302, -1.49 DPS) [world_drop]; Alchemist's Stone (13503, -2.18 DPS) [crafted] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wind Spirit Staff (6689, -1.88 DPS) [dungeon]; Hand of Righteousness (7721, -2.54 DPS) [dungeon]; Spire of Hakkar (10844, -3.53 DPS) [world] |
| off_hand | Twisting Essence Jar (249456) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Desertwalker Cane (12471, +0.00 DPS) [dungeon]; Cloud Stone (17737, +0.00 DPS) [dungeon]; Enthralled Sphere (11625, -0.58 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Lei of Lilies; shoulder: Living Shoulders; back: Battle Healer's Cloak; chest: Embrace of the Wind Serpent; wrist: Mender's Leather Bracers; hands: Feralheart Gauntlets; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Brainlash; trinket1: Darkspear Voodoo Seal; main_hand: Charstone Dirk; off_hand: Twisting Essence Jar

No-known-source sample (15 of 585, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 343.6. Weights run: 5.6s. Verify run: 67.0s. 1444 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.022, intellect=2.234 ± 0.044, spirit=4.276 ± 0.034, mp5=6.865 ± 0.033, crit=0.098 ± 0.010 per rating point (14 rating = 1%, 1.366 per %), spell_haste=not significant (-0.039 ± 0.480)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 173.1 healing_power points (17.46 DPS) | yes | Champion's Dragonhide Headdress (227205, -2.67 DPS) [vendor]; Wildheart Cowl (16720, -2.81 DPS) [dungeon]; Feralheart Headdress (226786, -3.26 DPS) [vendor] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | 92.3 healing_power points (9.31 DPS) | yes | Lady Maye's Pendant (14558, -0.72 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.94 DPS) [world_drop]; Heart of the Fiend (13960, -1.72 DPS) [dungeon] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 143.0 healing_power points (14.42 DPS) | yes | Champion's Dragonhide Pauldrons (227207, -2.32 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, -3.38 DPS) [vendor]; Feralheart Mantle (226785, -4.10 DPS) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Butcher's Apron (12608, +0.00 DPS) [dungeon]; Frostweaver Cape (12968, +0.00 DPS) [dungeon] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Alanna's Embrace (13314, -0.17 DPS) [dungeon]; Ironfeather Breastplate (15066, -0.54 DPS) [crafted] |
| wrist | Feralheart Bindings (226782) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bleak Howler Armguards (13208, +0.00 DPS) [dungeon]; Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | 125.5 healing_power points (12.66 DPS) | yes | Wildheart Gloves (16717, -0.47 DPS) [dungeon]; Hands of the Exalted Herald (12554, -1.23 DPS) [dungeon]; Devout Gloves (16692, -1.56 DPS) [dungeon] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cord (272398, +0.00 DPS) [vendor]; Elderwild Waistcord (279252, -0.72 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.34 DPS) [crafted] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | 191.0 healing_power points (19.26 DPS) | yes | Devout Skirt (16694, -3.54 DPS) [dungeon]; Haunting Specter Leggings (11929, -4.48 DPS) [dungeon]; Legionnaire's Dragonhide Legguards (227206, -4.55 DPS) [vendor] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | 148.3 healing_power points (14.96 DPS) | yes | Incandescent Mooncloth Boots (227862, -1.33 DPS) [vendor]; Mooncloth Boots (15802, -3.78 DPS) [crafted]; Devout Sandals (16691, -4.26 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -1.01 DPS) [dungeon]; Emerald Flame Ring (18395, -1.35 DPS) [dungeon]; Band of Mending (22334, -1.73 DPS) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-verified (343.6 DPS) | yes | The Postmaster's Seal (13392, +0.00 DPS) [dungeon]; Emerald Flame Ring (18395, +0.00 DPS) [dungeon]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindtap Talisman (18371, -2.77 DPS) [dungeon]; Ankh of Life (1713, -5.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -6.07 DPS) [quest] |
| trinket2 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-verified (343.6 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Mindtap Talisman (18371, +0.00 DPS) [dungeon]; Evonice's Landin' Pilla (18951, -0.12 DPS) [quest] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dancing Sliver (15854, -0.35 DPS) [quest]; Wind Spirit Staff (6689, -0.37 DPS) [dungeon]; Staff of Hale Magefire (13000, -0.44 DPS) [world_drop] |
| off_hand | Lapidis Tankard of Tidesippe (4696) | World drop [world_drop] | 79.6 healing_power points (8.03 DPS) | yes | Book of the Dead (13353, -0.33 DPS) [dungeon]; Thaurissan's Royal Scepter (11928, -0.37 DPS) [dungeon]; Lei of the Lifegiver (19312, -0.53 DPS) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Feralheart Bindings; waist: Feralheart Cord; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket2: Royal Seal of Eldre'Thalas; off_hand: Lapidis Tankard of Tidesippe

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 4532200100000000-00000000000000000000-5050035153113200)

Set DPS (verified): 642.0. Weights run: 4.2s. Verify run: 44.0s. 1444 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.204, intellect=2.451 ± 0.080, spirit=2.281 ± 0.075, mp5=3.336 ± 0.099, crit=0.327 ± 0.030 per rating point (14 rating = 1%, 4.573 per %), spell_haste=not significant (-0.743 ± 1.216)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 149.0 healing_power points (24.64 DPS) | yes | Champion's Dragonhide Headdress (227205, -4.67 DPS) [vendor]; Feralheart Headdress (226786, -5.51 DPS) [vendor]; Sanctified Leather Helm (22689, -5.80 DPS) [quest] |
| neck | Wavefront Necklace (20685) | Lord Skwol [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jeweled Amulet of Cainwyn (1443, +0.00 DPS) [world_drop]; Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Drake Tooth Necklace (21531, -1.65 DPS) [quest] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 101.8 healing_power points (16.83 DPS) | yes | Champion's Dragonhide Pauldrons (227207, -1.16 DPS) [vendor]; Warlord's Dragonhide Pauldrons (231672, -1.32 DPS) [vendor]; Devout Mantle (16695, -3.11 DPS) [dungeon] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 66.5 healing_power points (10.99 DPS) | yes | Frostweaver Cape (12968, -1.61 DPS) [dungeon]; Cloak of the Cosmos (18389, -2.24 DPS) [dungeon]; Darkspear Raider's Cloak (272063, -2.25 DPS) [vendor] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Alanna's Embrace (13314, -0.61 DPS) [dungeon]; Devout Robe (16690, -0.88 DPS) [dungeon]; Feralheart Embrace (226783, -1.31 DPS) [vendor] |
| wrist | Feralheart Bindings (226782) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Bracers (16553, +0.00 DPS) [pvp]; Bracers of Hope (22667, +0.00 DPS) [quest]; Bracers of Mending (23129, +0.00 DPS) [dungeon] |
| hands | Feralheart Gauntlets (226784) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Raider Handwraps (272097, -0.31 DPS) [vendor]; Wildheart Gloves (16717, -1.83 DPS) [dungeon] |
| waist | Feralheart Cord (226780) | Mokvar [vendor] | 85.9 healing_power points (14.20 DPS) | yes | Wisdom of the Timbermaw (19047, -0.01 DPS) [crafted]; General's Dragonhide Belt (16556, -0.30 DPS) [pvp]; Devout Belt (16696, -0.89 DPS) [dungeon] |
| legs | Leggings of Arcana (12756) | Leggings of Arcana [quest] | sim-verified (642.1 DPS) | yes | Legionnaire's Dragonhide Legguards (227206, -2.65 DPS) [vendor]; Devout Skirt (16694, -3.67 DPS) [dungeon]; Feralheart Pants (226787, -4.00 DPS) [vendor] |
| feet | Feralheart Sandals (226781) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Incandescent Mooncloth Boots (227862, +0.00 DPS) [vendor]; Mooncloth Boots (15802, -2.68 DPS) [crafted]; Faith Healer's Boots (22247, -3.53 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234033) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, -1.83 DPS) [dungeon]; Band of Mending (22334, -1.97 DPS) [dungeon]; Seal of Rivendare (13345, -2.66 DPS) [dungeon] |
| finger2 | Band of Piety (22681) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emerald Flame Ring (18395, +0.00 DPS) [dungeon]; Band of Mending (22334, +0.00 DPS) [dungeon]; Seal of Rivendare (13345, -0.37 DPS) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18470) | The Emerald Dream... [quest] | sim-verified (642.1 DPS) | yes | Serenity Field (272439, -0.33 DPS) [vendor]; Mindtap Talisman (18371, -1.21 DPS) [dungeon]; Briarwood Reed (12930, -2.48 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.33 DPS) [vendor]; Mindtap Talisman (18371, -2.21 DPS) [dungeon]; Briarwood Reed (12930, -3.48 DPS) [dungeon] |
| main_hand | Charstone Dirk (17710) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wind Spirit Staff (6689, -4.94 DPS) [dungeon]; Hand of Righteousness (7721, -6.18 DPS) [dungeon]; Staff of Hale Magefire (13000, -6.31 DPS) [world_drop] |
| off_hand | Lei of the Lifegiver (19312) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thaurissan's Royal Scepter (11928, +0.00 DPS) [dungeon]; Book of the Dead (13353, +0.00 DPS) [dungeon]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Wavefront Necklace; shoulder: Argent Elite Shoulders; back: Hide of the Wild; chest: Mooncloth Vest; wrist: Feralheart Bindings; waist: Feralheart Cord; legs: Leggings of Arcana; feet: Feralheart Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Piety; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; off_hand: Lei of the Lifegiver

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

