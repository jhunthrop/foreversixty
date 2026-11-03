# Leveling BiS: Restoration

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-000000000000000000-5510000000000000)

Set DPS (verified): 29.2. Weights run: 4.7s. Verify run: 18.2s. 258 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=1.987 ± 0.015, spirit=1.481 ± 0.014, mp5=4.402 ± 0.018, crit=0.118 ± 0.006 per rating point (14 rating = 1%, 1.646 per %), spell_haste=not significant (0.159 ± 0.067)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Hood (252507, -0.04 DPS) [crafted]; Trapper's Leather Hood (252505, -0.09 DPS) [crafted]; Shadow Goggles (4373, -0.37 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 3.0 healing_power points (0.12 DPS) | yes | Tarnished Locket (279870, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.9 healing_power points (0.74 DPS) | yes | Slime-encrusted Pads (6461, -0.26 DPS, sim-verified) [dungeon]; Serpent's Shoulders (5404, -0.31 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.41 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 7.9 healing_power points (0.33 DPS) | yes | Seer's Cape (6378, -0.04 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.08 DPS) [crafted]; Spirit Cloak (4792, -0.08 DPS) [vendor] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.4 healing_power points (1.21 DPS) | yes | Filigreed Pristine Gown (253901, -0.13 DPS, sim-verified) [crafted]; Robe of the Moccasin (6465, -0.23 DPS) [dungeon]; Corsair's Overshirt (5202, -0.27 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 23.1 healing_power points (0.95 DPS) | yes | Drakewing Bands (12999, -0.59 DPS) [world_drop]; Bright Bracers (3647, -0.63 DPS) [world_drop]; Owl Bracers (4796, -0.64 DPS, sim-verified) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 17.9 healing_power points (0.74 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.09 DPS) [world_drop]; Bright Gloves (3066, -0.17 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.9 healing_power points (0.78 DPS) | yes | Wisdom's Leather Belt (252433, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.13 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.25 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 37.8 healing_power points (1.56 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.62 DPS) [world_drop]; Stormrider's Leather Pants (252502, -0.82 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.9 healing_power points (0.86 DPS) | yes | Black Whelp Slippers (252424, -0.17 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.25 DPS) [crafted]; Kimbra Boots (6191, -0.35 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.9 healing_power points (0.49 DPS) | yes | Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; Minor Channeling Ring (1449, -0.33 DPS) [quest]; The 1 Ring (8350, -0.35 DPS) [world] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 8.9 healing_power points (0.37 DPS) | yes | Volcanic Rock Ring (12053, +0.00 DPS, sim-verified) [world_drop]; Minor Channeling Ring (1449, -0.20 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.17 DPS) [dungeon]; Lesser Staff of the Spire (1300, -0.28 DPS) [world]; Twisted Chanter's Staff (890, -3.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Band of Purification; main_hand: Staff of Westfall

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18845 Insignia of the Horde

### Band 30 (dwarf, 0000000000000000-000000000000000000-5530003311000000)

Set DPS (verified): 51.9. Weights run: 4.8s. Verify run: 16.4s. 458 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.032, intellect=2.172 ± 0.024, spirit=2.134 ± 0.023, mp5=5.589 ± 0.022, crit=0.227 ± 0.015 per rating point (14 rating = 1%, 3.178 per %), spell_haste=not significant (-0.050 ± 0.253)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.7 healing_power points (2.36 DPS) | yes | Holy Shroud (2721, -0.18 DPS) [world_drop]; Nightsky Cowl (4039, -0.62 DPS) [world_drop]; Dokebi Hat (14584, -0.72 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.2 healing_power points (0.82 DPS) | yes | Crystal Starfire Medallion (5003, -0.17 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.20 DPS) [world]; Pendant of Myzrael (4614, -0.21 DPS) [dungeon] |
| shoulder | Mantle of Honor (3560) | Bride of the Embalmer [quest] | 30.1 healing_power points (1.43 DPS) | yes | Nightsky Mantle (4718, -0.10 DPS) [world_drop]; Faerie Mantle (5820, -0.20 DPS) [quest]; Death Speaker Mantle (6685, -0.30 DPS) [dungeon] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 24.1 healing_power points (1.14 DPS) | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Sergeant Major's Cape (16315, -0.33 DPS) [pvp]; Amy's Blanket (13005, -0.43 DPS) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 43.9 healing_power points (2.09 DPS) | yes | Pristine Gown (253961, -0.01 DPS) [crafted]; Beguiler Robes (7728, -0.04 DPS) [dungeon]; Wisdom's Leather Tunic (252511, -0.35 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 27.6 healing_power points (1.31 DPS) | yes | Glowing Magical Bracelets (13106, -0.49 DPS) [world_drop]; Nightsky Wristbands (6407, -0.49 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.52 DPS) [dungeon] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 36.5 healing_power points (1.73 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.09 DPS, sim-verified) [dungeon]; Zodiac Gloves (7106, -0.61 DPS) [quest]; Silver-thread Gloves (6393, -0.71 DPS) [world_drop] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 44.8 healing_power points (2.13 DPS) | yes | Resilient Cord (14406, -1.11 DPS) [world_drop]; Stalker's Leather Belt (252521, -1.11 DPS) [crafted]; Silver-lined Belt (13011, -1.33 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.1 healing_power points (2.43 DPS) | yes | Pristine Leggings (253987, -0.09 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.46 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.55 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.7 healing_power points (2.08 DPS) | yes | Frothing Slippers (254003, -0.95 DPS) [crafted]; Fiery Slippers (254005, -0.95 DPS) [crafted]; Acidic Walkers (9454, -1.49 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 23.9 healing_power points (1.14 DPS) | yes | The Queen's Jewel (13094, -0.12 DPS) [world_drop]; Ring of Power Regulation (273030, -0.34 DPS) [dungeon]; Defias Mage Ring (1077, -0.42 DPS) [world] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 23.9 healing_power points (1.14 DPS) | yes | The Queen's Jewel (13094, -0.33 DPS, sim-verified) [world_drop]; Ring of Power Regulation (273030, -0.34 DPS) [dungeon]; Defias Mage Ring (1077, -0.42 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-verified (51.9 DPS) | yes | Gnarled Ash Staff (791, -0.01 DPS) [world_drop]; Mechanic's Pipehammer (9604, -0.29 DPS) [quest]; Death Speaker Scepter (2816, -0.43 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Mantle of Honor; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Staff of the Friar

No-known-source sample (15 of 458, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (dwarf, 0000000000000000-000000000000000000-5530003315510000)

Set DPS (verified): 80.2. Weights run: 5.6s. Verify run: 21.6s. 706 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.056, intellect=2.355 ± 0.042, spirit=2.365 ± 0.038, mp5=6.199 ± 0.036, crit=0.318 ± 0.024 per rating point (14 rating = 1%, 4.453 per %), spell_haste=not significant (0.126 ± 0.369)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 83.3 healing_power points (4.09 DPS) | yes | Whitemane's Chapeau (7720, -0.97 DPS, sim-verified) [dungeon]; Imperial Leather Helm (6433, -1.43 DPS) [dungeon]; Enduring Cap (3020, -1.43 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 38.8 healing_power points (1.91 DPS) | yes | Amberglow Talisman (10824, -0.75 DPS) [quest]; Explorers' League Commendation (7746, -0.86 DPS) [quest]; Necklace of Calisea (1714, -0.90 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 52.6 healing_power points (2.58 DPS) | yes | Sheepshear Mantle (13115, -0.45 DPS, sim-verified) [world_drop]; Mistscape Mantle (4734, -0.73 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.77 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 28.3 healing_power points (1.39 DPS) | yes | Blackforge Cape (6424, -0.00 DPS) [dungeon]; Ceremonial Centaur Blanket (6789, -0.11 DPS) [quest]; Glowing Thresher Cape (6901, -0.12 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 76.9 healing_power points (3.78 DPS) | yes | Deathchill Armor (10764, -0.59 DPS, sim-verified) [dungeon]; Doomsayer's Robe (4746, -0.90 DPS) [quest]; Civinad Robes (9623, -1.11 DPS) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 33.1 healing_power points (1.62 DPS) | yes | Mindthrust Bracers (1974, -0.13 DPS) [dungeon]; Earthen Silk Cuffs (254019, -0.55 DPS) [crafted]; Silkstream Cuffs (16791, -0.58 DPS) [quest] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 46.0 healing_power points (2.26 DPS) | yes | Mender's Leather Gloves (252530, +0.00 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.29 DPS) [world_drop]; Stormcloth Gloves (10011, -0.33 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 55.0 healing_power points (2.70 DPS) | yes | Mender's Leather Belt (252523, -0.53 DPS, sim-verified) [crafted]; Windchaser Cinch (14435, -0.73 DPS) [world_drop]; Prefect's Belt (250559, -0.81 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 71.2 healing_power points (3.50 DPS) | yes | Stoneweaver Leggings (9407, -0.14 DPS, sim-verified) [dungeon]; Misplaced Pantaloons (276201, -0.48 DPS) [vendor]; Wisdom's Leather Leggings (252519, -0.84 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 57.0 healing_power points (2.80 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Furen's Boots (13100, -0.13 DPS) [world_drop]; Gilded Slippers (254001, -0.54 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 28.3 healing_power points (1.39 DPS) | yes | Runemender's Seal (273031, -0.14 DPS) [world]; The Queen's Jewel (13094, -0.23 DPS) [world_drop]; Voodoo Band (1996, -0.23 DPS) [world] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 25.6 healing_power points (1.25 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.09 DPS) [world_drop]; Voodoo Band (1996, -0.10 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (80.2 DPS) | yes | Mograine's Might (7723, -1.62 DPS) [dungeon]; Silver Spade (4128, -1.73 DPS) [quest]; Death Speaker Scepter (2816, -4.36 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Sergeant Major's Cape; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 706, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (dwarf, 0000000000000000-000000000000000000-5530003315513151)

Set DPS (verified): 147.2. Weights run: 8.5s. Verify run: 38.3s. 909 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.142, intellect=3.147 ± 0.061, spirit=2.619 ± 0.046, mp5=6.942 ± 0.054, crit=0.390 ± 0.032 per rating point (14 rating = 1%, 5.462 per %), spell_haste=not significant (-0.201 ± 0.462)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 109.7 healing_power points (6.66 DPS) | yes | Papal Fez (9431, -0.64 DPS) [dungeon]; Winged Helm (13112, -0.67 DPS) [world_drop]; Eye of Theradras (17715, -1.11 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 54.5 healing_power points (3.31 DPS) | yes | Glowing Eye of Mordresh (10769, -0.57 DPS) [dungeon]; Necklace of Calisea (1714, -0.86 DPS) [dungeon]; Mindburst Medallion (11196, -1.37 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 83.9 healing_power points (5.10 DPS) | yes | Mender's Leather Shoulders (252538, -0.80 DPS) [crafted]; Mender's Mail Shoulders (252569, -0.80 DPS) [crafted]; Nethergeld Shoulders (254049, -1.03 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 45.1 healing_power points (2.74 DPS) | yes | Highborne Cloak (14450, -0.38 DPS) [world_drop]; Blackforge Cape (6424, -0.57 DPS) [dungeon]; Chainlink Towel (9648, -0.57 DPS) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 107.9 healing_power points (6.55 DPS) | yes | Feathered Breastplate (8349, -0.83 DPS) [crafted]; Deathchill Armor (10764, -1.23 DPS) [dungeon]; Earth Warder's Vest (21311, -1.54 DPS) [quest] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 56.7 healing_power points (3.45 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.68 DPS) [quest] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 82.9 healing_power points (5.04 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.09 DPS) [crafted]; Keeper's Gloves (14666, -1.34 DPS) [world_drop] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-verified (147.2 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Bloodlust Belt (14803, -0.01 DPS) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 109.8 healing_power points (6.67 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.76 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.95 DPS) [dungeon]; Dalewind Trousers (13008, -0.97 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 77.6 healing_power points (4.71 DPS) | yes | Vinerot Sandals (17748, -0.67 DPS) [dungeon]; Mender's Leather Shoes (252533, -0.80 DPS) [crafted]; Mender's Mail Boots (252565, -0.80 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 50.3 healing_power points (3.06 DPS) | yes | Choking Band (11868, -0.99 DPS) [quest]; Coldwater Ring (4550, -1.05 DPS) [quest]; Voodoo Band (1996, -1.24 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 34.1 healing_power points (2.07 DPS) | yes | Choking Band (11868, -0.00 DPS) [quest]; Coldwater Ring (4550, -0.06 DPS) [quest]; Voodoo Band (1996, -0.25 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.46 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.32 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.64 DPS) [quest]; Alchemist's Stone (13503, -1.27 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (147.2 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Charstone Dirk (17710, +0.00 DPS) [dungeon]; Spire of Hakkar (10844, -0.66 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 909, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 220.0. Weights run: 8.7s. Verify run: 65.8s. 1856 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.242, intellect=6.165 ± 0.147, spirit=4.940 ± 0.118, mp5=12.611 ± 0.131, crit=0.947 ± 0.084 per rating point (14 rating = 1%, 13.263 per %), spell_haste=not significant (-0.383 ± 0.908)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Dragonscale Helm (252605, -0.24 DPS) [crafted]; Devout Crown (16693, -0.48 DPS) [dungeon]; Crown of the Penitent (13216, -3.48 DPS, sim-verified) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 106.1 healing_power points (4.58 DPS) | yes | Necklace of Calisea (1714, -1.22 DPS) [dungeon]; Opaline Medallion (16623, -1.26 DPS, sim-verified) [quest]; Lady Alizabeth's Pendant (13002, -1.38 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 184.9 healing_power points (7.99 DPS) | yes | Drakesfire Epaulets (13133, -1.06 DPS) [world_drop]; Shroud of the Nathrezim (18720, -1.45 DPS) [dungeon]; Ironfeather Shoulders (15067, -2.06 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 88.8 healing_power points (3.84 DPS) | yes | Ethereal Mist Cape (11873, -0.48 DPS) [quest]; Highborne Cloak (14450, -0.59 DPS) [world_drop]; Imperial Red Cloak (8248, -0.70 DPS, sim-verified) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamwalker Armor (13123, -0.86 DPS) [world_drop]; Ironfeather Breastplate (15066, -1.29 DPS) [crafted]; Mooncloth Vest (14138, -4.34 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Magiskull Cuffs (13107, -0.21 DPS) [world_drop]; Spitfire Bracers (20481, -0.42 DPS) [crafted]; Magister's Bindings (16683, -5.00 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 160.6 healing_power points (6.94 DPS) | yes | Gauntlets of Elements (16672, -0.86 DPS) [dungeon]; Magister's Gloves (16684, -1.16 DPS, sim-verified) [dungeon]; Spitfire Gauntlets (20480, -1.18 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 178.8 healing_power points (7.72 DPS) | yes | Wisdom of the Timbermaw (19047, -0.73 DPS) [crafted]; Elunarian Belt (14465, -1.27 DPS, sim-verified) [world_drop]; Cord of Elements (16673, -1.70 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 230.1 healing_power points (9.94 DPS) | yes | Red Dragonscale Leggings (252603, -0.55 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.05 DPS) [dungeon]; Magister's Leggings (16687, -1.46 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Sandals of the Insurgent (13111, -0.37 DPS) [world_drop]; Furen's Boots (13100, -1.49 DPS) [world_drop]; Mooncloth Boots (15802, -2.40 DPS, sim-verified) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 116.0 healing_power points (5.01 DPS) | yes | Glowing Crystal Ring (18402, -1.81 DPS) [quest]; Maiden's Circle (13001, -1.87 DPS) [world_drop]; Welken Ring (5011, -2.18 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 96.3 healing_power points (4.16 DPS) | yes | Glowing Crystal Ring (18402, -0.75 DPS, sim-verified) [quest]; Maiden's Circle (13001, -1.01 DPS) [world_drop]; Welken Ring (5011, -1.33 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.71 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.13 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.56 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS, sim-verified) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; Charstone Dirk (17710, -0.61 DPS) [dungeon]; Hand of Edward the Odd (2243, -12.38 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; back: Sergeant Major's Cape; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Uther's Strength

No-known-source sample (15 of 1856, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (dwarf, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 642.0. Weights run: 6.4s. Verify run: 58.2s. 1856 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.377, intellect=1.236 ± 0.054, spirit=0.738 ± 0.046, mp5=2.083 ± 0.054, crit=0.499 ± 0.053 per rating point (14 rating = 1%, 6.982 per %), spell_haste=not significant (2.412 ± 1.109)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 105.6 healing_power points (32.55 DPS) | yes | Mooncloth Circlet (14140, -13.42 DPS) [crafted]; Devout Crown (16693, -13.83 DPS) [dungeon]; Crown of the Penitent (13216, -15.37 DPS) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.9 healing_power points (9.21 DPS) | yes | Horizon Choker (13085, -2.97 DPS) [world_drop]; Opaline Medallion (16623, -4.26 DPS) [quest]; Necklace of Calisea (1714, -4.95 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Eternity Pauldrons (279255, -0.10 DPS) [crafted]; Living Shoulders (15061, -0.93 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 16.5 healing_power points (5.10 DPS) | yes | Sergeant Major's Cape (16337, -0.23 DPS) [pvp]; Highborne Cloak (14450, -0.76 DPS) [world_drop]; Repairman's Cape (9605, -0.97 DPS) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormcloth Vest (10020, +0.00 DPS) [crafted]; Mooncloth Vest (14138, +0.00 DPS) [crafted]; Red Dragonscale Breastplate (15047, -1.15 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Leather Gauntlets (252551, +0.00 DPS) [crafted]; Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, +0.00 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (642.0 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Mender's Mail Belt (252591, +0.00 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Red Dragonscale Leggings (252603, +0.00 DPS) [crafted]; Elderwild Pants (279254, -0.56 DPS) [crafted]; Earthenweave Leggings (254125, -3.30 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.9 healing_power points (15.69 DPS) | yes | Swarmtender's Footpads (275609, -2.43 DPS) [crafted]; Broodwatcher's Treaders (275621, -2.43 DPS) [crafted]; Gilded Sandals (254107, -2.70 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 20.5 healing_power points (6.31 DPS) | yes | Glowing Crystal Ring (18402, -1.74 DPS) [quest]; Electrocutioner Lagnut (9447, -1.95 DPS) [dungeon]; Runemender's Seal (273031, -1.95 DPS) [world] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 16.3 healing_power points (5.02 DPS) | yes | Glowing Crystal Ring (18402, -0.44 DPS) [quest]; Electrocutioner Lagnut (9447, -0.65 DPS) [dungeon]; Runemender's Seal (273031, -0.65 DPS) [world] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -1.12 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.58 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.03 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -2.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.27 DPS) [quest]; Thunderbrew's Boot Flask (744, -3.73 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (642.0 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Uther's Strength

No-known-source sample (15 of 1856, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (tauren, 0000000000000000-000000000000000000-5510000000000000)

Set DPS (verified): 27.3. Weights run: 4.7s. Verify run: 17.1s. 247 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.021, intellect=1.987 ± 0.015, spirit=1.481 ± 0.014, mp5=4.402 ± 0.018, crit=0.118 ± 0.006 per rating point (14 rating = 1%, 1.646 per %), spell_haste=not significant (0.159 ± 0.067)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | 18.8 healing_power points (0.78 DPS) | yes | Pristine Circlet (253949, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.08 DPS) [crafted]; Trapper's Leather Hood (252505, -0.12 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 4.4 healing_power points (0.18 DPS) | yes | Scholarly Pendant (277203, -0.05 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 17.9 healing_power points (0.74 DPS) | yes | Serpent's Shoulders (5404, -0.31 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.41 DPS) [crafted]; Slime-encrusted Pads (6461, -0.45 DPS, sim-verified) [dungeon] |
| back | Sanguine Cape (14376) | World drop [world_drop] | 7.9 healing_power points (0.33 DPS) | yes | Pearl-clasped Cloak (5542, -0.08 DPS) [crafted]; Spirit Cloak (4792, -0.08 DPS) [vendor]; Seer's Cape (6378, -0.14 DPS, sim-verified) [dungeon] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | 29.4 healing_power points (1.21 DPS) | yes | Filigreed Pristine Gown (253901, -0.13 DPS, sim-verified) [crafted]; Robe of the Moccasin (6465, -0.23 DPS) [dungeon]; Corsair's Overshirt (5202, -0.27 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 23.1 healing_power points (0.95 DPS) | yes | Drakewing Bands (12999, -0.59 DPS) [world_drop]; Bright Bracers (3647, -0.63 DPS) [world_drop]; Owl Bracers (4796, -0.77 DPS, sim-verified) [vendor] |
| hands | Wisdom's Leather Gloves (252499) | Leatherworking [crafted] | 17.9 healing_power points (0.74 DPS) | yes | Pristine Gloves (253913, -0.07 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.09 DPS) [world_drop]; Bright Gloves (3066, -0.17 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 18.9 healing_power points (0.78 DPS) | yes | Wisdom's Leather Belt (252433, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.13 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.40 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | 37.8 healing_power points (1.56 DPS) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.62 DPS) [world_drop]; Stormrider's Leather Pants (252502, -0.82 DPS) [crafted] |
| feet | Wisdom's Leather Boots (252444) | Leatherworking [crafted] | 20.9 healing_power points (0.86 DPS) | yes | Pristine Boots (253889, -0.25 DPS) [crafted]; Black Whelp Slippers (252424, -0.27 DPS, sim-verified) [crafted]; Stormrider's Leather Boots (252443, -0.45 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 11.9 healing_power points (0.49 DPS) | yes | Band of Purification (12996, -0.13 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; The 1 Ring (8350, -0.35 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 9.9 healing_power points (0.41 DPS) | yes | Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; Band of Purification (12996, -0.22 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.27 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) | The Wrath of Rath'mael [quest] | sim-verified (27.3 DPS) | yes | Staff of the Blessed Seer (2271, +0.00 DPS) [dungeon]; Hammerbone (270018, +0.00 DPS) [quest]; Twisted Chanter's Staff (890, -1.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Shadow Goggles; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Wisdom's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Wisdom's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-000000000000000000-5530003311000000)

Set DPS (verified): 51.6. Weights run: 4.8s. Verify run: 15.3s. 451 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.032, intellect=2.172 ± 0.024, spirit=2.134 ± 0.023, mp5=5.589 ± 0.022, crit=0.227 ± 0.015 per rating point (14 rating = 1%, 3.178 per %), spell_haste=not significant (-0.050 ± 0.253)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 49.7 healing_power points (2.36 DPS) | yes | Holy Shroud (2721, -0.08 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.62 DPS) [world_drop]; Dokebi Hat (14584, -0.72 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.2 healing_power points (0.82 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.20 DPS) [world]; Pendant of Myzrael (4614, -0.21 DPS) [dungeon]; Crystal Starfire Medallion (5003, -0.23 DPS, sim-verified) [world_drop] |
| shoulder | Nightsky Mantle (4718) | World drop [world_drop] | 28.0 healing_power points (1.33 DPS) | yes | Death Speaker Mantle (6685, -0.07 DPS, sim-verified) [dungeon]; Emblazoned Shoulders (6399, -0.31 DPS) [world_drop]; Desert Shoulders (15457, -0.31 DPS) [quest] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 24.1 healing_power points (1.14 DPS) | yes | Resilient Cape (14400, -0.53 DPS) [world_drop]; Silver-thread Cloak (4713, -0.53 DPS) [world_drop]; Amy's Blanket (13005, -0.67 DPS, sim-verified) [world_drop] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 43.9 healing_power points (2.09 DPS) | yes | Pristine Gown (253961, -0.01 DPS) [crafted]; Beguiler Robes (7728, -0.04 DPS) [dungeon]; Wisdom's Leather Tunic (252511, -0.31 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 27.6 healing_power points (1.31 DPS) | yes | Glowing Magical Bracelets (13106, -0.49 DPS) [world_drop]; Spidertank Oilrag (9448, -0.52 DPS) [dungeon]; Nightsky Wristbands (6407, -0.53 DPS, sim-verified) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 36.5 healing_power points (1.73 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.14 DPS, sim-verified) [dungeon]; Silver-thread Gloves (6393, -0.71 DPS) [world_drop]; Truefaith Gloves (7049, -0.71 DPS) [crafted] |
| waist | Mender's Leather Belt (252523) | Leatherworking [crafted] | 44.8 healing_power points (2.13 DPS) | yes | Lilac Sash (6780, -0.90 DPS) [quest]; Resilient Cord (14406, -1.11 DPS) [world_drop]; Silver-lined Belt (13011, -1.29 DPS, sim-verified) [world_drop] |
| legs | Wisdom's Leather Leggings (252519) | Leatherworking [crafted] | 51.1 healing_power points (2.43 DPS) | yes | Pristine Leggings (253987, -0.16 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.46 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.55 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 43.7 healing_power points (2.08 DPS) | yes | Frothing Slippers (254003, -0.95 DPS) [crafted]; Fiery Slippers (254005, -0.95 DPS) [crafted]; Acidic Walkers (9454, -1.51 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 23.9 healing_power points (1.14 DPS) | yes | Solid Gold Band (279845, -0.11 DPS) [quest]; The Queen's Jewel (13094, -0.12 DPS) [world_drop]; Ring of Power Regulation (273030, -0.34 DPS) [dungeon] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 23.9 healing_power points (1.14 DPS) | yes | Solid Gold Band (279845, -0.08 DPS, sim-verified) [quest]; The Queen's Jewel (13094, -0.12 DPS) [world_drop]; Ring of Power Regulation (273030, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Friar (3415) | Blackfathom Deeps [dungeon] | sim-verified (51.6 DPS) | yes | Gnarled Ash Staff (791, -0.01 DPS) [world_drop]; Odo's Ley Staff (6318, -0.31 DPS) [dungeon]; Death Speaker Scepter (2816, -0.51 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Kaleidoscope Chain; shoulder: Nightsky Mantle; back: Glowing Thresher Cape; chest: Death Speaker Robes; hands: Gloves of Old; waist: Mender's Leather Belt; legs: Wisdom's Leather Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Staff of the Friar

No-known-source sample (15 of 451, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 0000000000000000-000000000000000000-5530003315510000)

Set DPS (verified): 79.4. Weights run: 5.6s. Verify run: 20.3s. 688 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.056, intellect=2.355 ± 0.042, spirit=2.365 ± 0.038, mp5=6.199 ± 0.036, crit=0.318 ± 0.024 per rating point (14 rating = 1%, 4.453 per %), spell_haste=not significant (0.126 ± 0.369)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 83.3 healing_power points (4.09 DPS) | yes | Whitemane's Chapeau (7720, -1.07 DPS, sim-verified) [dungeon]; Imperial Leather Helm (6433, -1.43 DPS) [dungeon]; Enduring Cap (3020, -1.43 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 38.8 healing_power points (1.91 DPS) | yes | Amberglow Talisman (10824, -0.75 DPS) [quest]; Necklace of Calisea (1714, -0.97 DPS, sim-verified) [dungeon]; Kaleidoscope Chain (13084, -0.98 DPS) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 52.6 healing_power points (2.58 DPS) | yes | Sheepshear Mantle (13115, -0.40 DPS, sim-verified) [world_drop]; Mistscape Mantle (4734, -0.73 DPS) [dungeon]; Earthen Silk Shoulders (254033, -0.77 DPS) [crafted] |
| back | Cloak of Blight (6832) | Nothing But The Truth [quest] | 30.7 healing_power points (1.51 DPS) | yes | Blackforge Cape (6424, +0.00 DPS, sim-verified) [dungeon]; Ceremonial Centaur Blanket (6789, -0.23 DPS) [quest]; Glowing Thresher Cape (6901, -0.24 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 76.9 healing_power points (3.78 DPS) | yes | Deathchill Armor (10764, -0.60 DPS, sim-verified) [dungeon]; Doomsayer's Robe (4746, -0.90 DPS) [quest]; Civinad Robes (9623, -1.11 DPS) [quest] |
| wrist | Enchanted Kodo Bracers (13119) | World drop [world_drop] | 33.1 healing_power points (1.62 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.55 DPS) [crafted]; Silkstream Cuffs (16791, -0.58 DPS) [quest] |
| hands | Bonefingers (10765) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 46.0 healing_power points (2.26 DPS) | yes | Mender's Leather Gloves (252530, -0.19 DPS) [crafted]; Gloves of Old (9395, -0.29 DPS) [world_drop]; Stormcloth Gloves (10011, -0.33 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 55.0 healing_power points (2.70 DPS) | yes | Windchaser Cinch (14435, -0.73 DPS) [world_drop]; Mender's Leather Belt (252523, -0.74 DPS, sim-verified) [crafted]; Prefect's Belt (250559, -0.81 DPS) [crafted] |
| legs | Warchief Kilt (7760) | Scarlet Monastery [dungeon] | 71.2 healing_power points (3.50 DPS) | yes | Stoneweaver Leggings (9407, -0.21 DPS, sim-verified) [dungeon]; Misplaced Pantaloons (276201, -0.48 DPS) [vendor]; Wisdom's Leather Leggings (252519, -0.84 DPS) [crafted] |
| feet | Mender's Leather Shoes (252533) (or Mender's Mail Boots (252565)) | Leatherworking [crafted] | 57.0 healing_power points (2.80 DPS) | yes | Mender's Mail Boots (252565, +0.00 DPS) [crafted]; Furen's Boots (13100, -0.13 DPS) [world_drop]; Gilded Slippers (254001, -0.54 DPS) [crafted] |
| finger1 | Welken Ring (5011) | World drop [world_drop] | 28.3 healing_power points (1.39 DPS) | yes | Runemender's Seal (273031, -0.14 DPS) [world]; The Queen's Jewel (13094, -0.23 DPS) [world_drop]; Voodoo Band (1996, -0.23 DPS) [world] |
| finger2 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 25.6 healing_power points (1.25 DPS) | yes | Runemender's Seal (273031, +0.00 DPS) [world]; The Queen's Jewel (13094, -0.09 DPS) [world_drop]; Voodoo Band (1996, -0.10 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (79.4 DPS) | yes | Mograine's Might (7723, -1.62 DPS) [dungeon]; Silver Spade (4128, -1.73 DPS) [quest]; Death Speaker Scepter (2816, -4.44 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Cloak of Blight; chest: Stormcloth Vest; wrist: Enchanted Kodo Bracers; hands: Bonefingers; waist: Gilded Cord; legs: Warchief Kilt; feet: Mender's Leather Shoes; finger1: Welken Ring; finger2: Electrocutioner Lagnut; trinket1: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 688, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (tauren, 0000000000000000-000000000000000000-5530003315513151)

Set DPS (verified): 146.7. Weights run: 8.5s. Verify run: 37.3s. 879 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.142, intellect=3.147 ± 0.061, spirit=2.619 ± 0.046, mp5=6.942 ± 0.054, crit=0.390 ± 0.032 per rating point (14 rating = 1%, 5.462 per %), spell_haste=not significant (-0.201 ± 0.462)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Gemburst Circlet (10751) | The God Hakkar [quest] | 109.7 healing_power points (6.66 DPS) | yes | Papal Fez (9431, -0.64 DPS) [dungeon]; Winged Helm (13112, -0.67 DPS) [world_drop]; Eye of Theradras (17715, -1.11 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 54.5 healing_power points (3.31 DPS) | yes | Glowing Eye of Mordresh (10769, -0.57 DPS) [dungeon]; Necklace of Calisea (1714, -0.86 DPS) [dungeon]; Choker of the High Shaman (4112, -1.08 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 83.9 healing_power points (5.10 DPS) | yes | Mender's Leather Shoulders (252538, -0.80 DPS) [crafted]; Mender's Mail Shoulders (252569, -0.80 DPS) [crafted]; Nethergeld Shoulders (254049, -1.03 DPS) [crafted] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 45.1 healing_power points (2.74 DPS) | yes | Highborne Cloak (14450, -0.38 DPS) [world_drop]; Blackforge Cape (6424, -0.57 DPS) [dungeon]; Chainlink Towel (9648, -0.57 DPS) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 107.9 healing_power points (6.55 DPS) | yes | Feathered Breastplate (8349, -0.83 DPS) [crafted]; Deathchill Armor (10764, -1.23 DPS) [dungeon]; Earth Warder's Vest (21311, -1.54 DPS) [quest] |
| wrist | Mender's Leather Bracers (252543) (or Mender's Mail Bracers (252573)) | Leatherworking [crafted] | 56.7 healing_power points (3.45 DPS) | yes | Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, -0.10 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.68 DPS) [quest] |
| hands | Mender's Leather Gauntlets (252551) (or Mender's Mail Gauntlets (252587)) | Leatherworking [crafted] | 82.9 healing_power points (5.04 DPS) | yes | Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted]; Gilded Gloves (254095, -0.09 DPS) [crafted]; Keeper's Gloves (14666, -1.34 DPS) [world_drop] |
| waist | Mender's Leather Waistguard (252477) | Leatherworking [crafted] | sim-verified (146.6 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Mender's Mail Belt (252591, +0.00 DPS) [crafted]; Bloodlust Belt (14803, -0.01 DPS) [world_drop] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 109.8 healing_power points (6.67 DPS) | yes | Jinxed Hoodoo Kilt (9474, -0.76 DPS) [dungeon]; Senior Designer's Pantaloons (11841, -0.95 DPS) [dungeon]; Dalewind Trousers (13008, -0.97 DPS) [world_drop] |
| feet | Sandals of the Insurgent (13111) | World drop [world_drop] | 77.6 healing_power points (4.71 DPS) | yes | Vinerot Sandals (17748, -0.67 DPS) [dungeon]; Mender's Leather Shoes (252533, -0.80 DPS) [crafted]; Mender's Mail Boots (252565, -0.80 DPS) [crafted] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 50.3 healing_power points (3.06 DPS) | yes | Coldwater Ring (4550, -1.05 DPS) [quest]; Ogremind Ring (1993, -1.24 DPS) [world_drop]; Voodoo Band (1996, -1.24 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 34.1 healing_power points (2.07 DPS) | yes | Coldwater Ring (4550, -0.06 DPS) [quest]; Ogremind Ring (1993, -0.25 DPS) [world_drop]; Voodoo Band (1996, -0.25 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Alchemist's Stone (13503, -0.46 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.32 DPS) [quest]; Alchemist's Stone (13503, -1.27 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (146.6 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Charstone Dirk (17710, +0.00 DPS) [dungeon]; Spire of Hakkar (10844, -0.66 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Gemburst Circlet; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; chest: Ghostweave Vest; wrist: Mender's Leather Bracers; hands: Mender's Leather Gauntlets; waist: Mender's Leather Waistguard; legs: Kilt of the Atal'ai Prophet; feet: Sandals of the Insurgent; finger1: Eye of Adaegus; finger2: Welken Ring; trinket1: Uther's Strength; trinket2: Ankh of Life; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (tauren, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 219.1. Weights run: 8.7s. Verify run: 63.7s. 1812 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.242, intellect=6.165 ± 0.147, spirit=4.940 ± 0.118, mp5=12.611 ± 0.131, crit=0.947 ± 0.084 per rating point (14 rating = 1%, 13.263 per %), spell_haste=not significant (-0.383 ± 0.908)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crown of the Penitent (13216, +0.00 DPS) [quest]; Black Dragonscale Helm (252605, -0.24 DPS) [crafted]; Devout Crown (16693, -0.48 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 106.1 healing_power points (4.58 DPS) | yes | Necklace of Calisea (1714, -1.22 DPS) [dungeon]; Lady Alizabeth's Pendant (13002, -1.38 DPS) [world_drop]; Choker of the High Shaman (4112, -1.54 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | 184.9 healing_power points (7.99 DPS) | yes | Ironfeather Shoulders (15067, -0.95 DPS) [crafted]; Drakesfire Epaulets (13133, -1.06 DPS) [world_drop]; Shroud of the Nathrezim (18720, -1.45 DPS) [dungeon] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 87.6 healing_power points (3.78 DPS) | yes | Highborne Cloak (14450, -0.53 DPS) [world_drop]; Cloak of Earth and Sky (279270, -0.80 DPS) [crafted]; Blackforge Cape (6424, -0.80 DPS) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Vest (14138, +0.00 DPS) [crafted]; Dreamwalker Armor (13123, -0.86 DPS) [world_drop]; Ironfeather Breastplate (15066, -1.29 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Bracers (16576, +0.00 DPS) [pvp]; Magister's Bindings (16683, +0.00 DPS) [dungeon]; Magiskull Cuffs (13107, -0.21 DPS) [world_drop] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | 160.6 healing_power points (6.94 DPS) | yes | Magister's Gloves (16684, -0.22 DPS) [dungeon]; Gauntlets of Elements (16672, -0.86 DPS) [dungeon]; Spitfire Gauntlets (20480, -1.18 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 178.8 healing_power points (7.72 DPS) | yes | General's Mail Waistband (16575, +0.00 DPS) [pvp]; Elunarian Belt (14465, -0.26 DPS) [world_drop]; Wisdom of the Timbermaw (19047, -0.73 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 230.1 healing_power points (9.94 DPS) | yes | Red Dragonscale Leggings (252603, -0.55 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.05 DPS) [dungeon]; Magister's Leggings (16687, -1.46 DPS) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Sandals of the Insurgent (13111, -0.37 DPS) [world_drop]; General's Mail Boots (231667, -1.38 DPS) [pvp] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 116.0 healing_power points (5.01 DPS) | yes | Chemist's Ring (15702, -1.49 DPS) [quest]; Glowing Crystal Ring (18402, -1.81 DPS) [quest]; Maiden's Circle (13001, -1.87 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 96.3 healing_power points (4.16 DPS) | yes | Chemist's Ring (15702, -0.64 DPS) [quest]; Glowing Crystal Ring (18402, -0.96 DPS) [quest]; Maiden's Circle (13001, -1.01 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -1.13 DPS) [quest]; Alchemist's Stone (13503, -2.41 DPS) [crafted]; Uther's Strength (11302, -2.49 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (219.1 DPS) | yes | Evonice's Landin' Pilla (18951, -0.43 DPS) [quest]; Alchemist's Stone (13503, -1.71 DPS) [crafted]; Uther's Strength (11302, -1.78 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; Charstone Dirk (17710, -0.61 DPS) [dungeon]; Soulkeeper (1607, -1.17 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves

No-known-source sample (15 of 1812, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (tauren, 0000000000000000-055000000000000000-5530003315513151)

Set DPS (verified): 639.8. Weights run: 6.4s. Verify run: 54.6s. 1812 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.377, intellect=1.236 ± 0.054, spirit=0.738 ± 0.046, mp5=2.083 ± 0.054, crit=0.499 ± 0.053 per rating point (14 rating = 1%, 6.982 per %), spell_haste=not significant (2.412 ± 1.109)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 105.6 healing_power points (32.55 DPS) | yes | Mooncloth Circlet (14140, -13.42 DPS) [crafted]; Champion's Mail Skullcap (227168, -13.46 DPS) [vendor]; Devout Crown (16693, -13.83 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.9 healing_power points (9.21 DPS) | yes | Horizon Choker (13085, -2.97 DPS) [world_drop]; Necklace of Calisea (1714, -4.95 DPS) [dungeon]; Archlight Talisman (15856, -5.40 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Mail Epaulets (227166, +0.00 DPS) [vendor]; Warlord's Mail Epaulets (231665, +0.00 DPS) [vendor] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 16.5 healing_power points (5.10 DPS) | yes | Highborne Cloak (14450, -0.76 DPS) [world_drop]; Wingveil Cloak (10802, -1.06 DPS) [dungeon]; Glowing Thresher Cape (6901, -1.12 DPS) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormcloth Vest (10020, +0.00 DPS) [crafted]; Mooncloth Vest (14138, +0.00 DPS) [crafted]; Legionnaire's Mail Chestguard (227171, +0.00 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mender's Leather Bracers (252543, +0.00 DPS) [crafted]; Mender's Mail Bracers (252573, +0.00 DPS) [crafted]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Mail Gloves (231666, +0.00 DPS) [vendor]; Mender's Leather Gauntlets (252551, +0.00 DPS) [crafted]; Mender's Mail Gauntlets (252587, +0.00 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (639.9 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Mender's Leather Waistguard (252477, +0.00 DPS) [crafted]; Mender's Mail Belt (252591, +0.00 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Mail Pants (227167, +0.00 DPS) [vendor]; General's Mail Pants (231670, +0.00 DPS) [vendor]; Red Dragonscale Leggings (252603, +0.00 DPS) [crafted] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 50.9 healing_power points (15.69 DPS) | yes | General's Mail Boots (231667, -1.58 DPS) [pvp]; Blood Guard's Mail Boots (227170, -2.18 DPS) [vendor]; Swarmtender's Footpads (275609, -2.43 DPS) [crafted] |
| finger1 | Band of the Hierophant (13096) | World drop [world_drop] | 20.5 healing_power points (6.31 DPS) | yes | Chemist's Ring (15702, -1.59 DPS) [quest]; Glowing Crystal Ring (18402, -1.74 DPS) [quest]; Solid Gold Band (279845, -1.78 DPS) [quest] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | 16.3 healing_power points (5.02 DPS) | yes | Chemist's Ring (15702, -0.30 DPS) [quest]; Glowing Crystal Ring (18402, -0.44 DPS) [quest]; Solid Gold Band (279845, -0.49 DPS) [quest] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -1.12 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.58 DPS) [quest]; Alchemist's Stone (13503, -2.94 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -2.82 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.27 DPS) [quest]; Alchemist's Stone (13503, -4.64 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (639.9 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Charstone Dirk (17710, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Glowing Eye of Mordresh; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Band of the Hierophant; finger2: Eye of Adaegus; trinket1: Shard of the Splithooves; trinket2: Uther's Strength

No-known-source sample (15 of 1812, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

