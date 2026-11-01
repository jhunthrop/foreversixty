# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5510000000000000)

Set DPS (verified): 26.7. Weights run: 6.2s. Verify run: 23.1s. 227 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=1.938 ± 0.009, spirit=1.454 ± 0.008, mp5=4.313 ± 0.013, crit=0.120 ± 0.006 per rating point (14 rating = 1%, 1.677 per %), spell_haste=0.228 ± 0.043

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (26.7 DPS) | yes | Wisdom's Leather Hood (252507, -0.04 DPS) [crafted]; Trapper's Leather Hood (252505, -0.11 DPS) [crafted]; Shadow Goggles (4373, -0.44 DPS, sim-verified) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.8 healing_power points (0.25 DPS) | yes | Scholarly Pendant (277203, -0.08 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.4 healing_power points (0.74 DPS) | yes | Slime-encrusted Pads (6461, -0.18 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.41 DPS) [crafted]; Forest Leather Mantle (4709, -0.49 DPS) [world_drop] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 7.8 healing_power points (0.33 DPS) | yes | Seer's Cape (6378, -0.04 DPS) [dungeon]; Spirit Cloak (4792, -0.08 DPS) [vendor]; Sylvan Cloak (4793, -0.08 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.1 healing_power points (1.23 DPS) | yes | Filigreed Pristine Gown (253901, -0.13 DPS, sim-verified) [crafted]; Robe of the Moccasin (6465, -0.23 DPS) [dungeon]; Corsair's Overshirt (5202, -0.27 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.6 healing_power points (0.95 DPS) | yes | Drakewing Bands (12999, -0.59 DPS) [world_drop]; Bright Bracers (3647, -0.63 DPS) [world_drop]; Owl Bracers (4796, -0.67 DPS, sim-verified) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 17.8 healing_power points (0.75 DPS) | yes | Pristine Gloves (253913, -0.04 DPS) [crafted]; Magefist Gloves (12977, -0.09 DPS) [world_drop]; Bright Gloves (3066, -0.18 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.8 healing_power points (0.79 DPS) | yes | Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.14 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.27 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 37.4 healing_power points (1.58 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.64 DPS) [world_drop]; Dreamer's Leggings (270016, -0.68 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.7 healing_power points (0.87 DPS) | yes | Black Whelp Slippers (252424, -0.19 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.25 DPS) [crafted]; Kimbra Boots (6191, -0.36 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.6 healing_power points (0.49 DPS) | yes | Deep Fathom Ring (6463, -0.18 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; Minor Channeling Ring (1449, -0.33 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 8.7 healing_power points (0.37 DPS) | yes | Deep Fathom Ring (6463, -0.05 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; Minor Channeling Ring (1449, -0.20 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 19.4 healing_power points (0.82 DPS) | yes | Staff of Westfall (2042, +0.00 DPS, sim-verified) [quest]; Staff of the Blessed Seer (2271, -0.20 DPS) [dungeon]; Taskmaster Axe (5194, -0.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Tarnished Locket; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 227, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18845 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (dwarf, 0000000000000000-000000000000000000-5530003311000000)

Set DPS (verified): 46.7. Weights run: 6.2s. Verify run: 22.8s. 427 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=2.153 ± 0.013, spirit=2.159 ± 0.012, mp5=5.680 ± 0.012, crit=0.215 ± 0.012 per rating point (14 rating = 1%, 3.014 per %), spell_haste=not significant (0.076 ± 0.029)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.6 healing_power points (2.33 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.61 DPS) [world_drop]; Dokebi Hat (14584, -0.71 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.2 healing_power points (0.81 DPS) | yes | Crystal Starfire Medallion (5003, -0.16 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.20 DPS) [dungeon]; Glowing Green Talisman (5002, -0.20 DPS) [world_drop] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 30.2 healing_power points (1.42 DPS) | yes | Nightsky Mantle (4718, -0.10 DPS) [world_drop]; Faerie Mantle (5820, -0.20 DPS) [quest]; Talbar Mantle (10657, -0.28 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 24.3 healing_power points (1.14 DPS) | yes | Prelacy Cape (7004, +0.00 DPS, sim-verified) [quest]; Repairman's Cape (9605, -0.14 DPS) [quest]; Sergeant Major's Cape (16315, -0.33 DPS) [pvp] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 43.9 healing_power points (2.06 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.01 DPS) [dungeon]; Beguiler Robes (7728, -0.04 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 27.8 healing_power points (1.31 DPS) | yes | Nightsky Wristbands (6407, -0.34 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.50 DPS) [world_drop]; Spidertank Oilrag (9448, -0.51 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 36.7 healing_power points (1.72 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.10 DPS, sim-verified) [dungeon]; Zodiac Gloves (7106, -0.61 DPS) [quest]; Silver-thread Gloves (6393, -0.71 DPS) [world_drop] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 44.9 healing_power points (2.11 DPS) | yes | Dokebi Cord (14578, -1.10 DPS) [world_drop]; Resilient Cord (14406, -1.10 DPS) [world_drop]; Silver-lined Belt (13011, -1.24 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.2 healing_power points (2.41 DPS) | yes | Pristine Leggings (253987, -0.16 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.45 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.55 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.7 healing_power points (2.06 DPS) | yes | Frothing Slippers (254003, -0.94 DPS) [crafted]; Black Slippers (254007, -0.94 DPS) [crafted]; Acidic Walkers (9454, -1.39 DPS, sim-verified) [dungeon] |
| finger1 | The Queen's Jewel (13094) | World drop [world_drop] | 21.6 healing_power points (1.01 DPS) | yes | Runemender's Seal (273031, -0.02 DPS) [world]; Ring of Power Regulation (273030, -0.21 DPS) [dungeon]; Ring of Calm (6790, -0.30 DPS) [quest] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 21.1 healing_power points (0.99 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Ring of Power Regulation (273030, -0.19 DPS) [dungeon]; Ring of Calm (6790, -0.28 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (46.7 DPS) | yes | Hardened Root Staff (1317, -1.34 DPS) [quest]; Manual Crowd Pummeler (9449, -1.40 DPS, sim-verified) [dungeon]; Spine of the Devourer (273040, -1.51 DPS) [world] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 26.0 healing_power points (1.22 DPS) | yes | Satyr's Rod (15962, -0.41 DPS) [world_drop]; Orb of Mistmantle (13031, -0.44 DPS, sim-verified) [world_drop]; Insignia Buckler (4066, -0.51 DPS) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: The Queen's Jewel; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 427, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (dwarf, 0000000000000000-000000000000000000-5530003315510000)

Set DPS (verified): 78.1. Weights run: 7.3s. Verify run: 31.0s. 676 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.037, intellect=2.478 ± 0.016, spirit=2.379 ± 0.012, mp5=6.301 ± 0.019, crit=0.334 ± 0.019 per rating point (14 rating = 1%, 4.679 per %), spell_haste=not significant (0.031 ± 0.039)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 77.5 healing_power points (3.75 DPS) | yes | Whitemane's Chapeau (7720, -0.79 DPS, sim-verified) [dungeon]; Enduring Cap (3020, -1.03 DPS) [world_drop]; Imperial Leather Helm (6433, -1.05 DPS) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 34.0 healing_power points (1.65 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.49 DPS) [quest]; Explorers' League Commendation (7746, -0.61 DPS) [quest] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 53.9 healing_power points (2.61 DPS) | yes | Inquisitor's Shawl (19507, -0.32 DPS) [dungeon]; Mistscape Mantle (4734, -0.72 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.82 DPS) [crafted] |
| back | Blackforge Cape (6424) | Uldaman: Ancient Treasure [dungeon] | 29.3 healing_power points (1.42 DPS) | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Ceremonial Centaur Blanket (6789, -0.14 DPS) [quest]; Glowing Thresher Cape (6901, -0.16 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 77.0 healing_power points (3.73 DPS) | yes | Dreamweave Vest (10021, -1.04 DPS) [crafted]; Black Mageweave Robe (10001, -1.16 DPS) [crafted]; Deathchill Armor (10764, -1.36 DPS, sim-verified) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 33.7 healing_power points (1.63 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Windchaser Cuffs (14429, -0.55 DPS) [world_drop]; Nightsky Wristbands (6407, -0.57 DPS) [world_drop] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 47.5 healing_power points (2.30 DPS) | yes | Mender's Leather Gloves (252530, +0.00 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.32 DPS) [world_drop]; Stormcloth Gloves (10011, -0.33 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 56.1 healing_power points (2.72 DPS) | yes | Mender's Leather Belt (252523, -0.52 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.72 DPS) [world_drop]; Prefect's Belt (250559, -0.81 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 64.2 healing_power points (3.11 DPS) | yes | Stoneweaver Leggings (9407, -0.14 DPS) [dungeon]; Misplaced Pantaloons (276201, -0.16 DPS, sim-verified) [vendor]; Wisdom's Leather Leggings (252519, -0.44 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 58.0 healing_power points (2.81 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Furen's Boots (13100, -0.13 DPS) [world_drop]; Gilded Slippers (254001, -0.54 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 29.0 healing_power points (1.41 DPS) | yes | Ogremind Ring (1993, -0.22 DPS) [world_drop]; The Queen's Jewel (13094, -0.24 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.31 DPS) [dungeon] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 24.5 healing_power points (1.19 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.09 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -1.82 DPS) [dungeon]; Hypnotic Blade (7714, -2.04 DPS) [dungeon] |
| off_hand | Battered Viking Shield (9403) | Uldaman: Olaf [dungeon] | sim-verified (78.1 DPS) | yes | Ravager's Shield (14777, -0.01 DPS) [world_drop]; Orb of Souls (249395, -0.22 DPS) [crafted]; Orb of Lorica (11262, -7.66 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Sheepshear Mantle; back: Blackforge Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Battered Viking Shield

No-known-source sample (15 of 676, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 0000000000000000-000000000000000000-5530003315513151)

Set DPS (verified): 147.8. Weights run: 10.9s. Verify run: 47.0s. 879 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.068, intellect=3.447 ± 0.040, spirit=2.672 ± 0.025, mp5=7.230 ± 0.036, crit=0.477 ± 0.031 per rating point (14 rating = 1%, 6.681 per %), spell_haste=not significant (0.165 ± 0.090)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 115.7 healing_power points (6.74 DPS) | yes | Winged Helm (13112, -0.56 DPS) [world_drop]; Papal Fez (9431, -1.11 DPS) [dungeon]; Serpentskin Helm (8261, -1.19 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 58.9 healing_power points (3.44 DPS) | yes | Necklace of Calisea (1714, -0.94 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.07 DPS) [dungeon]; Mindburst Medallion (11196, -1.45 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 90.3 healing_power points (5.26 DPS) | yes | Mender's Leather Shoulders (252538, -0.97 DPS) [crafted]; Mender's Mail Shoulders (252569, -0.97 DPS) [crafted]; Sheepshear Mantle (13115, -1.12 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 48.6 healing_power points (2.83 DPS) | yes | Highborne Cloak (14450, -0.40 DPS) [world_drop]; Blackforge Cape (6424, -0.60 DPS) [dungeon]; Chainlink Towel (9648, -0.67 DPS) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 111.3 healing_power points (6.49 DPS) | yes | Feathered Breastplate (8349, -0.74 DPS) [crafted]; Earth Warder's Vest (21311, -1.35 DPS) [quest]; Deathchill Armor (10764, -1.58 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 59.2 healing_power points (3.45 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.57 DPS) [quest] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 86.1 healing_power points (5.02 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.07 DPS) [crafted]; Keeper's Gloves (14666, -1.25 DPS) [world_drop] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-verified (147.8 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Bloodlust Belt (14803, +0.00 DPS) [world_drop]; Mender's Mail Belt (252591, +0.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 116.1 healing_power points (6.77 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.87 DPS) [dungeon]; Dalewind Trousers (13008, -1.07 DPS) [world_drop]; Senior Designer's Pantaloons (11841, -1.10 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 81.0 healing_power points (4.72 DPS) | yes | Vinerot Sandals (17748, -0.66 DPS) [dungeon]; Furen's Boots (13100, -0.82 DPS) [world_drop]; Mender's Leather Shoes (252533, -0.83 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 52.7 healing_power points (3.07 DPS) | yes | Choking Band (11868, -1.05 DPS) [quest]; Coldwater Ring (4550, -1.07 DPS) [quest]; Voodoo Band (1996, -1.20 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 35.9 healing_power points (2.09 DPS) | yes | Choking Band (11868, -0.07 DPS) [quest]; Coldwater Ring (4550, -0.09 DPS) [quest]; Voodoo Band (1996, -0.22 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.43 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.31 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.62 DPS) [quest]; Alchemist's Stone (13503, -1.25 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -6.16 DPS) [dungeon]; Spellshifter Rod (9527, -8.20 DPS) [quest]; Hand of Righteousness (7721, -9.25 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 221.4. Weights run: 11.0s. Verify run: 81.8s. 1826 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.103, intellect=5.230 ± 0.065, spirit=4.626 ± 0.056, mp5=11.677 ± 0.062, crit=0.688 ± 0.047 per rating point (14 rating = 1%, 9.638 per %), spell_haste=not significant (0.203 ± 0.185)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 231.9 healing_power points (10.83 DPS) | yes | Crown of the Penitent (13216, -0.24 DPS) [quest]; Devout Crown (16693, -0.79 DPS) [dungeon]; Black Dragonscale Helm (252605, -1.10 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 91.7 healing_power points (4.29 DPS) | yes | Lady Alizabeth's Pendant (13002, -1.04 DPS) [world_drop]; Necklace of Calisea (1714, -1.06 DPS) [dungeon]; Opaline Medallion (16623, -1.11 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 162.5 healing_power points (7.59 DPS) | yes | Drakesfire Epaulets (13133, -0.93 DPS) [world_drop]; Ironfeather Shoulders (15067, -0.97 DPS) [crafted]; Hero's Pauldrons (8310, -1.60 DPS) [dungeon] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 78.8 healing_power points (3.68 DPS) | yes | Imperial Red Cloak (8248, -0.13 DPS) [world_drop]; Ethereal Mist Cape (11873, -0.36 DPS) [quest]; Highborne Cloak (14450, -0.62 DPS) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, +0.00 DPS) [crafted]; Dreamwalker Armor (13123, -0.83 DPS) [world_drop]; Ironfeather Breastplate (15066, -1.06 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 99.1 healing_power points (4.63 DPS) | yes | Magister's Bindings (16683, -0.16 DPS) [dungeon]; Magiskull Cuffs (13107, -0.48 DPS) [world_drop]; Spitfire Bracers (20481, -0.49 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 145.9 healing_power points (6.82 DPS) | yes | Magister's Gloves (16684, -0.37 DPS) [dungeon]; Gauntlets of Elements (16672, -0.92 DPS) [dungeon]; Mender's Leather Gauntlets (252551, -1.23 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 157.2 healing_power points (7.35 DPS) | yes | Elunarian Belt (14465, -0.30 DPS) [world_drop]; Wisdom of the Timbermaw (19047, -0.53 DPS) [crafted]; Cord of Elements (16673, -1.68 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 208.8 healing_power points (9.76 DPS) | yes | Red Dragonscale Leggings (252603, -1.00 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.19 DPS) [dungeon]; Magister's Leggings (16687, -1.55 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Sandals of the Insurgent (13111, -0.35 DPS) [world_drop]; Furen's Boots (13100, -1.46 DPS) [world_drop] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 103.2 healing_power points (4.82 DPS) | yes | Maiden's Circle (13001, -1.81 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.89 DPS) [quest]; Choking Band (11868, -2.01 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 86.9 healing_power points (4.06 DPS) | yes | Maiden's Circle (13001, -1.05 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.13 DPS) [quest]; Choking Band (11868, -1.25 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -1.11 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.54 DPS) [quest]; Alchemist's Stone (13503, -2.41 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (221.4 DPS) | yes | Evonice's Landin' Pilla (18951, -0.43 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.86 DPS) [quest]; Alchemist's Stone (13503, -1.73 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -5.18 DPS) [world_drop]; Hand of Edward the Odd (2243, -6.50 DPS) [world_drop]; Charstone Dirk (17710, -6.84 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves

No-known-source sample (15 of 1826, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 612.4. Weights run: 7.9s. Verify run: 63.1s. 1826 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.183, intellect=1.485 ± 0.029, spirit=0.801 ± 0.022, mp5=2.411 ± 0.028, crit=0.564 ± 0.030 per rating point (14 rating = 1%, 7.894 per %), spell_haste=1.639 ± 0.403

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 110.9 healing_power points (29.46 DPS) | yes | Devout Crown (16693, -9.82 DPS, sim-verified) [dungeon]; Crown of the Penitent (13216, -12.34 DPS) [quest]; Black Dragonscale Helm (252605, -12.38 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.9 healing_power points (6.61 DPS) | yes | Horizon Choker (13085, -0.23 DPS) [world_drop]; Opaline Medallion (16623, -1.48 DPS) [quest]; Necklace of Calisea (1714, -2.36 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 49.4 healing_power points (13.12 DPS) | yes | Eternity Pauldrons (279255, -1.32 DPS) [crafted]; Mender's Leather Shoulders (252538, -1.71 DPS) [crafted]; Mooncloth Shoulders (14139, -7.29 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 19.5 healing_power points (5.19 DPS) | yes | Sergeant Major's Cape (16337, +0.00 DPS, sim-verified) [pvp]; Highborne Cloak (14450, -0.79 DPS) [world_drop]; Prelacy Cape (7004, -0.99 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (612.4 DPS) | yes | Stormcloth Vest (10020, -0.86 DPS) [crafted]; Mooncloth Robe (18486, -1.20 DPS) [crafted]; Mooncloth Vest (14138, -4.76 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (612.4 DPS) | yes | Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -5.09 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (612.4 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, +0.00 DPS) [crafted]; Mender's Leather Gauntlets (252551, -3.59 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (612.4 DPS) | yes | Mender's Leather Waistguard (252477, -0.02 DPS) [crafted]; Mender's Mail Belt (252591, -0.02 DPS) [crafted]; Wisdom of the Timbermaw (19047, -2.91 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (612.4 DPS) | yes | Elderwild Pants (279254, -1.45 DPS) [crafted]; Earthenweave Leggings (254125, -3.95 DPS) [crafted]; Red Dragonscale Leggings (252603, -4.13 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 55.2 healing_power points (14.66 DPS) | yes | Mender's Leather Boots (252472, -2.89 DPS) [crafted]; Mender's Mail Sabatons (252579, -2.89 DPS) [crafted]; Gilded Sandals (254107, -3.60 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 23.7 healing_power points (6.28 DPS) | yes | Glowing Crystal Ring (18402, -1.55 DPS) [quest]; Maiden's Circle (13001, -2.25 DPS) [world_drop]; Frigid Ring (18679, -2.34 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 18.5 healing_power points (4.92 DPS) | yes | Glowing Crystal Ring (18402, -0.19 DPS) [quest]; Maiden's Circle (13001, -0.88 DPS) [world_drop]; Frigid Ring (18679, -0.97 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (612.4 DPS) | yes | Ankh of Life (1713, -1.29 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.71 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.14 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (612.4 DPS) | yes | Ankh of Life (1713, -1.60 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, -2.65 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.08 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (612.4 DPS) | yes | Charstone Dirk (17710, -16.76 DPS) [dungeon]; Hand of Edward the Odd (2243, -20.15 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -26.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Uther's Strength

No-known-source sample (15 of 1826, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (tauren, 0000000000000000-000000000000000000-5510000000000000)

Set DPS (verified): 26.2. Weights run: 6.2s. Verify run: 20.6s. 216 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.012, intellect=1.938 ± 0.009, spirit=1.454 ± 0.008, mp5=4.313 ± 0.013, crit=0.120 ± 0.006 per rating point (14 rating = 1%, 1.677 per %), spell_haste=0.228 ± 0.043

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 18.4 healing_power points (0.78 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.06 DPS) [crafted]; Trapper's Leather Hood (252505, -0.12 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.4 healing_power points (0.18 DPS) | yes | Scholarly Pendant (277203, -0.07 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.4 healing_power points (0.74 DPS) | yes | Slime-encrusted Pads (6461, -0.22 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.41 DPS) [crafted]; Forest Leather Mantle (4709, -0.49 DPS) [world_drop] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 7.8 healing_power points (0.33 DPS) | yes | Seer's Cape (6378, -0.07 DPS, sim-verified) [dungeon]; Spirit Cloak (4792, -0.08 DPS) [vendor]; Antiquated Cloak (4799, -0.08 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.1 healing_power points (1.23 DPS) | yes | Filigreed Pristine Gown (253901, -0.14 DPS, sim-verified) [crafted]; Robe of the Moccasin (6465, -0.23 DPS) [dungeon]; Corsair's Overshirt (5202, -0.27 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 22.6 healing_power points (0.95 DPS) | yes | Featherbead Bracers (15452, -0.55 DPS) [quest]; Drakewing Bands (12999, -0.59 DPS) [world_drop]; Owl Bracers (4796, -0.63 DPS, sim-verified) [vendor] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 20.8 healing_power points (0.88 DPS) | yes | Wisdom's Leather Gloves (252499, +0.00 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Magefist Gloves (12977, -0.22 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.8 healing_power points (0.79 DPS) | yes | Wisdom's Leather Belt (252433, -0.13 DPS) [crafted]; Keller's Girdle (2911, -0.14 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.16 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 37.4 healing_power points (1.58 DPS) | yes | Filigreed Pristine Leggings (253937, -0.14 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.64 DPS) [world_drop]; Ghastly Trousers (15449, -0.81 DPS) [quest] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.7 healing_power points (0.87 DPS) | yes | Black Whelp Slippers (252424, -0.22 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.25 DPS) [crafted]; Bluegill Sandals (1560, -0.46 DPS) [world] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.6 healing_power points (0.49 DPS) | yes | Band of Purification (12996, -0.12 DPS) [world_drop]; Deep Fathom Ring (6463, -0.18 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 9.7 healing_power points (0.41 DPS) | yes | Deep Fathom Ring (6463, -0.10 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; Band of Purification (12996, -0.25 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Orgrimmar (15444) | Hidden Enemies [quest] | 23.6 healing_power points (1.00 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.06 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -0.18 DPS) [world_drop]; Staff of the Blessed Seer (2271, -0.38 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Staff of Orgrimmar

No-known-source sample (15 of 216, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection; 21568 Rune of Duty

### Band 30 (tauren, 0000000000000000-000000000000000000-5530003311000000)

Set DPS (verified): 46.3. Weights run: 6.2s. Verify run: 21.8s. 420 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.008, intellect=2.153 ± 0.013, spirit=2.159 ± 0.012, mp5=5.680 ± 0.012, crit=0.215 ± 0.012 per rating point (14 rating = 1%, 3.014 per %), spell_haste=not significant (0.076 ± 0.029)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.6 healing_power points (2.33 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.61 DPS) [world_drop]; Dokebi Hat (14584, -0.71 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.2 healing_power points (0.81 DPS) | yes | Crystal Starfire Medallion (5003, -0.16 DPS, sim-verified) [world_drop]; Pendant of Myzrael (4614, -0.20 DPS) [dungeon]; Glowing Green Talisman (5002, -0.20 DPS) [world_drop] |
| shoulder | Ghostly Mantle (3324) | Deathstalkers in Shadowfang [quest] | 28.4 healing_power points (1.34 DPS) | yes | Nightsky Mantle (4718, +0.00 DPS, sim-verified) [world_drop]; Talbar Mantle (10657, -0.20 DPS) [quest]; Death Speaker Mantle (6685, -0.22 DPS) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 24.3 healing_power points (1.14 DPS) | yes | Nightsky Cloak (4719, -0.53 DPS) [world_drop]; Insignia Cloak (4722, -0.53 DPS) [world_drop]; Amy's Blanket (13005, -0.60 DPS, sim-verified) [world_drop] |
| chest | Wisdom's Leather Tunic (252511) | Leatherworking [crafted] | 43.9 healing_power points (2.06 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.01 DPS) [dungeon]; Beguiler Robes (7728, -0.04 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 27.8 healing_power points (1.31 DPS) | yes | Nightsky Wristbands (6407, -0.40 DPS, sim-verified) [world_drop]; Glowing Magical Bracelets (13106, -0.50 DPS) [world_drop]; Spidertank Oilrag (9448, -0.51 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 36.7 healing_power points (1.72 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.11 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.51 DPS) [quest]; Tattered Mittens (270030, -0.65 DPS) [quest] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 44.9 healing_power points (2.11 DPS) | yes | Lilac Sash (6780, -0.89 DPS) [quest]; Dokebi Cord (14578, -1.10 DPS) [world_drop]; Silver-lined Belt (13011, -1.24 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.2 healing_power points (2.41 DPS) | yes | Pristine Leggings (253987, -0.09 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.45 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.55 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.7 healing_power points (2.06 DPS) | yes | Frothing Slippers (254003, -0.94 DPS) [crafted]; Black Slippers (254007, -0.94 DPS) [crafted]; Acidic Walkers (9454, -1.40 DPS, sim-verified) [dungeon] |
| finger1 | The Queen's Jewel (13094) | World drop [world_drop] | 21.6 healing_power points (1.01 DPS) | yes | Runemender's Seal (273031, -0.02 DPS) [world]; Solid Gold Band (279845, -0.18 DPS) [quest]; Ring of Power Regulation (273030, -0.21 DPS) [dungeon] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 21.1 healing_power points (0.99 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; Solid Gold Band (279845, -0.16 DPS) [quest]; Ring of Power Regulation (273030, -0.19 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-verified (46.3 DPS) | yes | Manual Crowd Pummeler (9449, -1.40 DPS, sim-verified) [dungeon]; Spine of the Devourer (273040, -1.51 DPS) [world]; Tendonscraper Dagger (273053, -1.51 DPS) [world] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 26.0 healing_power points (1.22 DPS) | yes | Satyr's Rod (15962, -0.41 DPS) [world_drop]; Witch's Finger (16887, -0.41 DPS) [quest]; Orb of Mistmantle (13031, -0.48 DPS, sim-verified) [world_drop] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Ghostly Mantle; back: Glowing Thresher Cape; chest: Wisdom's Leather Tunic; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: The Queen's Jewel; finger2: Electrocutioner Lagnut; main_hand: Death Speaker Scepter; off_hand: Orb of Souls

No-known-source sample (15 of 420, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-000000000000000000-5530003315510000)

Set DPS (verified): 77.3. Weights run: 7.3s. Verify run: 28.2s. 658 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.037, intellect=2.478 ± 0.016, spirit=2.379 ± 0.012, mp5=6.301 ± 0.019, crit=0.334 ± 0.019 per rating point (14 rating = 1%, 4.679 per %), spell_haste=not significant (0.031 ± 0.039)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 77.5 healing_power points (3.75 DPS) | yes | Enduring Cap (3020, -1.03 DPS) [world_drop]; Imperial Leather Helm (6433, -1.05 DPS) [dungeon]; Whitemane's Chapeau (7720, -1.09 DPS, sim-verified) [dungeon] |
| neck | Necklace of Calisea (1714) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 34.0 healing_power points (1.65 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.49 DPS) [quest]; Kaleidoscope Chain (13084, -0.71 DPS) [world_drop] |
| shoulder | Sheepshear Mantle (13115) | World drop [world_drop] | 53.9 healing_power points (2.61 DPS) | yes | Inquisitor's Shawl (19507, -0.27 DPS, sim-verified) [dungeon]; Mistscape Mantle (4734, -0.72 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.82 DPS) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 30.9 healing_power points (1.50 DPS) | yes | Blackforge Cape (6424, +0.00 DPS, sim-verified) [dungeon]; Ceremonial Centaur Blanket (6789, -0.22 DPS) [quest]; Glowing Thresher Cape (6901, -0.24 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 77.0 healing_power points (3.73 DPS) | yes | Dreamweave Vest (10021, -1.04 DPS) [crafted]; Black Mageweave Robe (10001, -1.16 DPS) [crafted]; Deathchill Armor (10764, -1.26 DPS, sim-verified) [dungeon] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 33.7 healing_power points (1.63 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Windchaser Cuffs (14429, -0.55 DPS) [world_drop]; Nightsky Wristbands (6407, -0.57 DPS) [world_drop] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 47.5 healing_power points (2.30 DPS) | yes | Mender's Leather Gloves (252530, -0.20 DPS) [crafted]; Gloves of Old (9395, -0.32 DPS) [world_drop]; Stormcloth Gloves (10011, -0.33 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 56.1 healing_power points (2.72 DPS) | yes | Mender's Leather Belt (252523, -0.68 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.72 DPS) [world_drop]; Prefect's Belt (250559, -0.81 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 64.2 healing_power points (3.11 DPS) | yes | Stoneweaver Leggings (9407, -0.14 DPS) [dungeon]; Misplaced Pantaloons (276201, -0.26 DPS, sim-verified) [vendor]; Wisdom's Leather Leggings (252519, -0.44 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 58.0 healing_power points (2.81 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Furen's Boots (13100, -0.13 DPS) [world_drop]; Gilded Slippers (254001, -0.54 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 29.0 healing_power points (1.41 DPS) | yes | Ogremind Ring (1993, -0.22 DPS) [world_drop]; The Queen's Jewel (13094, -0.24 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.31 DPS) [dungeon] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 24.5 healing_power points (1.19 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; The Queen's Jewel (13094, -0.02 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.09 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (77.3 DPS) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (77.3 DPS) | yes | Illusionary Rod (7713, -1.82 DPS) [dungeon]; Hypnotic Blade (7714, -2.04 DPS) [dungeon]; Gut Ripper (2164, -8.37 DPS, sim-verified) [world_drop] |
| off_hand | Battered Viking Shield (9403) | Uldaman: Olaf [dungeon] | 31.8 healing_power points (1.54 DPS) | yes | Ravager's Shield (14777, -0.01 DPS) [world_drop]; Forcestone Buckler (17508, -0.04 DPS) [quest]; Orb of Souls (249395, -0.22 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Necklace of Calisea; shoulder: Sheepshear Mantle; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; finger2: Voodoo Band; trinket1: Ankh of Life; main_hand: Hand of Righteousness; off_hand: Battered Viking Shield

No-known-source sample (15 of 658, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (tauren, 0000000000000000-000000000000000000-5530003315513151)

Set DPS (verified): 147.3. Weights run: 10.9s. Verify run: 46.3s. 849 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.068, intellect=3.447 ± 0.040, spirit=2.672 ± 0.025, mp5=7.230 ± 0.036, crit=0.477 ± 0.031 per rating point (14 rating = 1%, 6.681 per %), spell_haste=not significant (0.165 ± 0.090)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 115.7 healing_power points (6.74 DPS) | yes | Winged Helm (13112, -0.56 DPS) [world_drop]; Papal Fez (9431, -1.11 DPS) [dungeon]; Serpentskin Helm (8261, -1.19 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 58.9 healing_power points (3.44 DPS) | yes | Necklace of Calisea (1714, -0.94 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.07 DPS) [dungeon]; Choker of the High Shaman (4112, -1.19 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 90.3 healing_power points (5.26 DPS) | yes | Mender's Leather Shoulders (252538, -0.97 DPS) [crafted]; Mender's Mail Shoulders (252569, -0.97 DPS) [crafted]; Sheepshear Mantle (13115, -1.12 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 48.6 healing_power points (2.83 DPS) | yes | Highborne Cloak (14450, -0.40 DPS) [world_drop]; Blackforge Cape (6424, -0.60 DPS) [dungeon]; Chainlink Towel (9648, -0.67 DPS) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 111.3 healing_power points (6.49 DPS) | yes | Feathered Breastplate (8349, -0.74 DPS) [crafted]; Earth Warder's Vest (21311, -1.35 DPS) [quest]; Deathchill Armor (10764, -1.58 DPS) [dungeon] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 59.2 healing_power points (3.45 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.57 DPS) [quest] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 86.1 healing_power points (5.02 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.07 DPS) [crafted]; Keeper's Gloves (14666, -1.25 DPS) [world_drop] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-verified (147.3 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Bloodlust Belt (14803, +0.00 DPS) [world_drop]; Mender's Mail Belt (252591, +0.00 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 116.1 healing_power points (6.77 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.87 DPS) [dungeon]; Dalewind Trousers (13008, -1.07 DPS) [world_drop]; Senior Designer's Pantaloons (11841, -1.10 DPS) [dungeon] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 81.0 healing_power points (4.72 DPS) | yes | Vinerot Sandals (17748, -0.66 DPS) [dungeon]; Furen's Boots (13100, -0.82 DPS) [world_drop]; Mender's Leather Shoes (252533, -0.83 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 52.7 healing_power points (3.07 DPS) | yes | Coldwater Ring (4550, -1.07 DPS) [quest]; Ogremind Ring (1993, -1.20 DPS) [world_drop]; Voodoo Band (1996, -1.20 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 35.9 healing_power points (2.09 DPS) | yes | Coldwater Ring (4550, -0.09 DPS) [quest]; Ogremind Ring (1993, -0.22 DPS) [world_drop]; Voodoo Band (1996, -0.22 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (147.3 DPS) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.43 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (147.3 DPS) | yes | Evonice's Landin' Pilla (18951, -0.31 DPS) [quest]; Alchemist's Stone (13503, -1.25 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Charstone Dirk (17710, -6.16 DPS) [dungeon]; Spellshifter Rod (9527, -8.20 DPS) [quest]; Hand of Righteousness (7721, -9.25 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 849, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (tauren, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 220.2. Weights run: 11.0s. Verify run: 79.9s. 1782 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.103, intellect=5.230 ± 0.065, spirit=4.626 ± 0.056, mp5=11.677 ± 0.062, crit=0.688 ± 0.047 per rating point (14 rating = 1%, 9.638 per %), spell_haste=not significant (0.203 ± 0.185)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 231.9 healing_power points (10.83 DPS) | yes | Devout Crown (16693, -0.79 DPS) [dungeon]; Black Dragonscale Helm (252605, -1.10 DPS) [crafted]; Crown of the Penitent (13216, -3.79 DPS, sim-verified) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 91.7 healing_power points (4.29 DPS) | yes | Necklace of Calisea (1714, -1.06 DPS) [dungeon]; Choker of the High Shaman (4112, -1.33 DPS) [quest]; Lady Alizabeth's Pendant (13002, -2.71 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 162.5 healing_power points (7.59 DPS) | yes | Ironfeather Shoulders (15067, -0.97 DPS) [crafted]; Hero's Pauldrons (8310, -1.60 DPS) [dungeon]; Drakesfire Epaulets (13133, -7.94 DPS, sim-verified) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 76.0 healing_power points (3.55 DPS) | yes | Cloak of Earth and Sky (279270, -0.53 DPS) [crafted]; Chainlink Towel (9648, -0.66 DPS) [quest]; Highborne Cloak (14450, -0.85 DPS, sim-verified) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamwalker Armor (13123, -0.83 DPS) [world_drop]; Ironfeather Breastplate (15066, -1.06 DPS) [crafted]; Mooncloth Vest (14138, -5.30 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 99.1 healing_power points (4.63 DPS) | yes | General's Mail Bracers (16576, -0.03 DPS) [pvp]; Magiskull Cuffs (13107, -0.48 DPS) [world_drop]; Magister's Bindings (16683, -5.76 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 145.9 healing_power points (6.82 DPS) | yes | Gauntlets of Elements (16672, -0.92 DPS) [dungeon]; Mender's Leather Gauntlets (252551, -1.23 DPS) [crafted]; Magister's Gloves (16684, -6.43 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 157.2 healing_power points (7.35 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.53 DPS) [crafted]; Elunarian Belt (14465, -6.39 DPS, sim-verified) [world_drop] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 208.8 healing_power points (9.76 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.19 DPS) [dungeon]; Magister's Leggings (16687, -1.55 DPS) [dungeon]; Red Dragonscale Leggings (252603, -5.36 DPS, sim-verified) [crafted] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Sandals of the Insurgent (13111, -0.35 DPS) [world_drop]; Furen's Boots (13100, -1.46 DPS) [world_drop]; Mooncloth Boots (15802, -2.89 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 103.2 healing_power points (4.82 DPS) | yes | Chemist's Ring (15702, -1.51 DPS) [quest]; Maiden's Circle (13001, -1.81 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.89 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 86.9 healing_power points (4.06 DPS) | yes | Chemist's Ring (15702, -0.48 DPS, sim-verified) [quest]; Maiden's Circle (13001, -1.05 DPS) [world_drop]; Glowing Crystal Ring (18402, -1.13 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -1.11 DPS) [quest]; Alchemist's Stone (13503, -2.41 DPS) [crafted]; Uther's Strength (11302, -2.43 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.57 DPS, sim-verified) [quest]; Alchemist's Stone (13503, -1.73 DPS) [crafted]; Uther's Strength (11302, -1.75 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, -5.18 DPS) [world_drop]; Charstone Dirk (17710, -6.84 DPS) [dungeon]; Hand of Edward the Odd (2243, -13.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves

No-known-source sample (15 of 1782, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (tauren, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 611.0. Weights run: 7.9s. Verify run: 61.4s. 1782 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.183, intellect=1.485 ± 0.029, spirit=0.801 ± 0.022, mp5=2.411 ± 0.028, crit=0.564 ± 0.030 per rating point (14 rating = 1%, 7.894 per %), spell_haste=1.639 ± 0.403

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 110.9 healing_power points (29.46 DPS) | yes | Devout Crown (16693, -9.65 DPS, sim-verified) [dungeon]; Champion's Mail Skullcap (227168, -11.54 DPS) [vendor]; Crown of the Penitent (13216, -12.34 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.9 healing_power points (6.61 DPS) | yes | Horizon Choker (13085, -0.23 DPS) [world_drop]; Necklace of Calisea (1714, -2.36 DPS) [dungeon]; Archlight Talisman (15856, -2.66 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 49.4 healing_power points (13.12 DPS) | yes | Champion's Mail Epaulets (227166, +0.00 DPS) [vendor]; Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor]; Mooncloth Shoulders (14139, -7.28 DPS, sim-verified) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 19.5 healing_power points (5.19 DPS) | yes | Highborne Cloak (14450, -1.02 DPS, sim-verified) [world_drop]; Blackforge Cape (6424, -1.18 DPS) [dungeon]; Keeper's Cloak (14665, -1.25 DPS) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (611.0 DPS) | yes | Legionnaire's Mail Chestguard (227171, +0.00 DPS) [vendor]; Warlord's Mail Chestguard (231668, -0.28 DPS) [vendor]; Mooncloth Vest (14138, -5.12 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (611.0 DPS) | yes | Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -5.22 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (611.0 DPS) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Mender's Leather Gauntlets (252551, -3.55 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (611.0 DPS) | yes | Mender's Leather Waistguard (252477, -0.02 DPS) [crafted]; Mender's Mail Belt (252591, -0.02 DPS) [crafted]; Wisdom of the Timbermaw (19047, -2.92 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (611.0 DPS) | yes | Legionnaire's Mail Pants (227167, +0.00 DPS) [vendor]; General's Mail Pants (231670, +0.00 DPS) [vendor]; Red Dragonscale Leggings (252603, -4.05 DPS, sim-verified) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 55.2 healing_power points (14.66 DPS) | yes | General's Mail Boots (231667, -1.45 DPS) [pvp]; Blood Guard's Mail Boots (227170, -2.23 DPS) [vendor]; Gilded Sandals (254107, -3.36 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 23.7 healing_power points (6.28 DPS) | yes | Chemist's Ring (15702, -1.49 DPS) [quest]; Glowing Crystal Ring (18402, -1.55 DPS) [quest]; Maiden's Circle (13001, -2.25 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 18.5 healing_power points (4.92 DPS) | yes | Chemist's Ring (15702, -0.12 DPS) [quest]; Glowing Crystal Ring (18402, -0.19 DPS) [quest]; Maiden's Circle (13001, -0.88 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (611.0 DPS) | yes | Ankh of Life (1713, -1.29 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.71 DPS) [quest]; Alchemist's Stone (13503, -2.99 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (611.0 DPS) | yes | Evonice's Landin' Pilla (18951, -1.81 DPS, sim-verified) [quest]; Ankh of Life (1713, -2.23 DPS) [world_drop]; Alchemist's Stone (13503, -3.93 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (611.0 DPS) | yes | Charstone Dirk (17710, -16.76 DPS) [dungeon]; Hand of Edward the Odd (2243, -19.71 DPS, sim-verified) [world_drop]; Hand of Righteousness (7721, -26.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Uther's Strength

No-known-source sample (15 of 1782, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

