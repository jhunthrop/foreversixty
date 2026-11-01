# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 22.5. Weights run: 2.7s. Verify run: 1.6s. 258 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.577 ± 0.009, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.031 per %), hit=0.188 ± 0.007 per rating point (10 rating = 1%, 1.884 per %), spell_haste=not significant (-0.026 ± 0.112), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.872 ± 0.003, spirit=not significant (-0.023 ± 0.111), mp5=not significant (0.137 ± 0.151)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (0.75 DPS) | yes | Pristine Circlet (253949, -0.25 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.36 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -0.37 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.2 spell_power points (0.85 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.44 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.52 DPS) [crafted]; Slime-encrusted Pads (6461, -0.81 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Sanguine Cape (14376, -0.14 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.19 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 10.9 spell_power points (0.91 DPS) | yes | Filigreed Pristine Gown (253901, -0.25 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.36 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 3.3 spell_power points (0.27 DPS) | yes | Bright Bracers (3647, -0.08 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.13 DPS) [dungeon]; Owl Bracers (4796, -0.66 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.3 spell_power points (0.61 DPS) | yes | Gnoll Casting Gloves (892, -0.11 DPS) [world]; Windfelt Gloves (5630, -0.13 DPS) [quest]; Serpent Gloves (5970, -0.42 DPS, sim-verified) [dungeon] |
| waist | Stormrider's Leather Belt (252432) (or Pristine Sash (253925)) | Leatherworking [crafted] | 6.3 spell_power points (0.53 DPS) | yes | Pristine Sash (253925, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 13.5 spell_power points (1.12 DPS) | yes | Filigreed Pristine Leggings (253937, -0.33 DPS) [crafted]; Totemic Leather Pants (252446, -0.35 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.36 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.3 spell_power points (0.78 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.28 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.51 DPS) | yes | Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; Furlfeather Ring (286745, -0.43 DPS) [world]; The 1 Ring (8350, -0.46 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.5 spell_power points (0.29 DPS) | yes | Furlfeather Ring (286745, -0.21 DPS) [world]; The 1 Ring (8350, -0.24 DPS) [world]; Volcanic Rock Ring (12053, -0.29 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Twisted Chanter's Staff (890, -0.19 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.38 DPS) [world]; Staff of Westfall (2042, -0.43 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Stormrider's Leather Belt; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18845 Insignia of the Horde

### Band 30 (dwarf, 5510331300000000-000000000000000000-0000000000000000)

Set DPS (verified): 45.7. Weights run: 2.9s. Verify run: 1.7s. 458 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.630 ± 0.012, crit=0.091 ± 0.002 per rating point (14 rating = 1%, 1.270 per %), hit=0.247 ± 0.008 per rating point (10 rating = 1%, 2.467 per %), spell_haste=-2.228 ± 0.170, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.806 ± 0.002, spirit=0.435 ± 0.068, mp5=0.936 ± 0.131

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (45.7 DPS) | yes | Holy Shroud (2721, +0.00 DPS) [world_drop]; Enduring Cap (3020, +0.00 DPS) [world_drop]; Totemic Leather Helm (252456, -0.03 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-verified (45.7 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, +0.00 DPS) [world_drop]; Pendant of Myzrael (4614, -0.13 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, -0.19 DPS) [dungeon]; Magician's Mantle (12998, -0.44 DPS) [world_drop]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-verified (45.7 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Glowing Thresher Cape (6901, +0.00 DPS) [dungeon]; Repairman's Cape (9605, +0.00 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (45.7 DPS) | yes | Beguiler Robes (7728, +0.00 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.14 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.8 spell_power points (1.30 DPS) | yes | Mindthrust Bracers (1974, -0.64 DPS) [dungeon]; Nightsky Wristbands (6407, -0.74 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.74 DPS) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | sim-verified (45.7 DPS) | yes | Gloves of Meditation (4318, +0.00 DPS) [crafted]; Gloves of Old (9395, +0.00 DPS) [world_drop]; Shilly Mitts (9609, +0.00 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.5 spell_power points (1.71 DPS) | yes | Belt of Arugal (6392, -0.27 DPS) [dungeon]; Mender's Leather Belt (252523, -0.34 DPS) [crafted]; Moss Cinch (6911, -0.39 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (45.7 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS) [dungeon]; Wisdom's Leather Leggings (252519, -0.12 DPS) [crafted]; Dark Ritual Leggings (270031, -0.17 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.2 spell_power points (1.45 DPS) | yes | Acidic Walkers (9454, -0.15 DPS) [dungeon]; Glinteye Slippers (273024, -0.39 DPS) [dungeon]; Spidersilk Boots (4320, -0.40 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (1.01 DPS) | yes | Minor Channeling Ring (1449, -0.32 DPS) [quest]; Electrocutioner Lagnut (9447, -0.34 DPS) [dungeon]; Runemender's Seal (273031, -0.34 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.8 spell_power points (0.97 DPS) | yes | Minor Channeling Ring (1449, -0.28 DPS) [quest]; Electrocutioner Lagnut (9447, -0.30 DPS) [dungeon]; Runemender's Seal (273031, -0.30 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (45.7 DPS) | yes | Bloodscalp Channeling Staff (1998, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, +0.00 DPS) [dungeon]; Reef Axe (6905, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 458, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (dwarf, 5510331300001350-010000000000000000-0000000000000000)

Set DPS (verified): 82.2. Weights run: 3.2s. Verify run: 2.2s. 706 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.936 ± 0.025, crit=0.247 ± 0.008 per rating point (14 rating = 1%, 3.459 per %), hit=0.431 ± 0.019 per rating point (10 rating = 1%, 4.305 per %), spell_haste=-2.455 ± 0.273, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.692 ± 0.002, spirit=not significant (0.632 ± 0.203), mp5=1.844 ± 0.324

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 29.6 spell_power points (3.19 DPS) | yes | Big Voodoo Mask (8201, -0.81 DPS) [crafted]; Whitemane's Chapeau (7720, -0.89 DPS, sim-verified) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.91 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.5 spell_power points (1.56 DPS) | yes | Necklace of Calisea (1714, -0.38 DPS) [dungeon]; Amberglow Talisman (10824, -0.88 DPS) [quest]; Triune Amulet (7722, -1.01 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.2 spell_power points (2.07 DPS) | yes | Sheepshear Mantle (13115, -0.04 DPS) [world_drop]; Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Mantle of Doan (7712, -0.16 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 10.8 spell_power points (1.16 DPS) | yes | Long Silken Cloak (4326, -0.01 DPS) [crafted]; Guardian Cloak (5965, -0.01 DPS) [crafted]; Blackforge Cape (6424, -0.08 DPS) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 35.3 spell_power points (3.81 DPS) | yes | Robe of the Magi (1716, -0.48 DPS) [world_drop]; Robe of Power (7054, -0.54 DPS) [crafted]; Beguiler Robes (7728, -1.08 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.5 spell_power points (1.57 DPS) | yes | Green Whelp Bracers (7386, -0.09 DPS) [crafted]; Guardian Leather Bracers (4260, -0.31 DPS) [crafted]; Turtle Scale Bracers (8198, -0.42 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 30.9 spell_power points (3.34 DPS) | yes | Dreamweave Gloves (10019, -0.51 DPS) [crafted]; Red Mageweave Gloves (10018, -1.14 DPS) [crafted]; Skycaller's Leather Gloves (252529, -1.24 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.0 spell_power points (2.27 DPS) | yes | Gilded Cord (254037, -0.19 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.31 DPS) [crafted]; Star Belt (4329, -0.59 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.2 spell_power points (2.72 DPS) | yes | Warchief Kilt (7760, -0.30 DPS) [dungeon]; Stoneweaver Leggings (9407, -0.34 DPS) [dungeon]; Misplaced Pantaloons (276201, -0.52 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.4 spell_power points (3.07 DPS) | yes | Skycaller's Leather Shoes (252532, -0.62 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.62 DPS) [crafted]; Mender's Leather Shoes (252533, -1.02 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (1.58 DPS) | yes | Snake Hoop (6750, -0.43 DPS) [quest]; Welken Ring (5011, -0.60 DPS) [world_drop]; Voodoo Band (1996, -0.67 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 11.5 spell_power points (1.24 DPS) | yes | Snake Hoop (6750, -0.09 DPS) [quest]; Welken Ring (5011, -0.26 DPS) [world_drop]; Voodoo Band (1996, -0.33 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (82.2 DPS) | yes | Illusionary Rod (7713, -1.94 DPS) [dungeon]; Bloodscalp Channeling Staff (1998, -1.98 DPS) [world]; Mograine's Might (7723, -11.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 706, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (dwarf, 5510331300001350-030000000000000000-5030000000000000)

Set DPS (verified): 106.4. Weights run: 3.3s. Verify run: 2.4s. 909 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.801 ± 0.028, crit=0.298 ± 0.009 per rating point (14 rating = 1%, 4.174 per %), hit=0.491 ± 0.023 per rating point (10 rating = 1%, 4.913 per %), spell_haste=not significant (-1.235 ± 0.318), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.691 ± 0.002, spirit=not significant (0.576 ± 0.267), mp5=not significant (1.146 ± 0.335)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 39.6 spell_power points (4.57 DPS) | yes | Red Mageweave Headband (10033, -0.53 DPS) [crafted]; Charged Scorpid Helm (252581, -1.04 DPS) [crafted]; Dreamweave Circlet (10041, -1.45 DPS, sim-verified) [crafted] |
| neck | Mindburst Medallion (11196) | Tremors of the Earth [quest] | 13.7 spell_power points (1.58 DPS) | yes | Horizon Choker (13085, -0.02 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -0.03 DPS) [dungeon]; Triune Amulet (7722, -0.19 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.6 spell_power points (4.11 DPS) | yes | Rockshard Pauldrons (9411, -1.12 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -1.38 DPS) [crafted]; Rotgrip Mantle (17732, -2.37 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.3 spell_power points (2.57 DPS) | yes | Big Voodoo Cloak (8216, -1.16 DPS) [crafted]; Imperial Red Cloak (8248, -1.29 DPS) [world_drop]; Runecloth Cloak (13860, -2.42 DPS, sim-verified) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 40.8 spell_power points (4.71 DPS) | yes | Dreamweave Vest (10021, -0.87 DPS) [crafted]; Flamestrider Robes (11747, -0.97 DPS) [dungeon]; Acumen Robes (17775, -1.61 DPS, sim-verified) [quest] |
| wrist | Skycaller's Leather Bracers (252542) (or Skycaller's Mail Bracers (252571)) | Leatherworking [crafted] | 18.5 spell_power points (2.13 DPS) | yes | Skycaller's Mail Bracers (252571, +0.00 DPS) [crafted]; Runic Leather Bracers (15092, -0.06 DPS) [crafted]; Nethergeld Cuffs (254061, -0.35 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 30.9 spell_power points (3.56 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.49 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.49 DPS) [crafted]; Dreamweave Gloves (10019, -0.65 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.6 spell_power points (2.95 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Tooled Leather Belt (252544, -0.37 DPS) [crafted]; Satyrmane Sash (17755, -0.42 DPS) [dungeon] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 28.0 spell_power points (3.23 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.02 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.12 DPS) [dungeon]; Red Mageweave Pants (10009, -0.50 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.0 spell_power points (3.23 DPS) | yes | Skycaller's Leather Boots (252471, -0.49 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.49 DPS) [crafted]; Greaves of Withering Despair (22240, -1.26 DPS, sim-verified) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.61 DPS) | yes | Eye of Adaegus (5266, -0.26 DPS) [world_drop]; Snake Hoop (6750, -0.46 DPS) [quest]; Ring of Power Regulation (273030, -0.53 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.50 DPS) | yes | Eye of Adaegus (5266, -0.15 DPS) [world_drop]; Snake Hoop (6750, -0.35 DPS) [quest]; Ring of Power Regulation (273030, -0.41 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (106.4 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (106.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.13 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (106.4 DPS) | yes | Spire of Hakkar (10844, -0.30 DPS) [world]; Spellshifter Rod (9527, -1.02 DPS, sim-verified) [quest]; Zum'rah's Vexing Cane (18082, -1.36 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Mindburst Medallion; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Skycaller's Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Senior Designer's Pantaloons; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Evonice's Landin' Pilla; main_hand: Kindling Stave

No-known-source sample (15 of 909, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (dwarf, 5510331300001350-030000000000000000-5033520000000000)

Set DPS (verified): 146.0. Weights run: 3.4s. Verify run: 6.4s. 1856 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.867 ± 0.034, crit=0.361 ± 0.010 per rating point (14 rating = 1%, 5.052 per %), hit=0.635 ± 0.027 per rating point (10 rating = 1%, 6.346 per %), spell_haste=-1.861 ± 0.401, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.701 ± 0.002, spirit=not significant (1.028 ± 0.322), mp5=2.342 ± 0.404

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 55.0 spell_power points (6.57 DPS) | yes | Blue Dragonscale Helm (252604, -0.96 DPS) [crafted]; Dreadmist Mask (16698, -1.04 DPS) [dungeon]; Coif of Elements (16667, -1.28 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.7 spell_power points (2.47 DPS) | yes | Horizon Choker (13085, -0.53 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -0.63 DPS) [world_drop]; Mindburst Medallion (11196, -1.98 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 40.6 spell_power points (4.84 DPS) | yes | Mantle of the Timbermaw (19050, -0.86 DPS) [crafted]; Argent Shoulders (19059, -1.00 DPS) [crafted]; Shroud of the Nathrezim (18720, -4.67 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.4 spell_power points (3.03 DPS) | yes | Cloak of Earth and Sky (279270, -0.06 DPS) [crafted]; Runecloth Cloak (13860, -1.13 DPS) [crafted]; Sergeant Major's Cape (16337, -1.22 DPS) [pvp] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 60.2 spell_power points (7.19 DPS) | yes | Ghostweave Vest (14141, -1.42 DPS) [crafted]; Mooncloth Vest (14138, -1.71 DPS) [crafted]; Feathered Breastplate (8349, -4.42 DPS, sim-verified) [crafted] |
| wrist | Bindings of Elements (16671) | Stratholme: Crypt Beast [dungeon] | sim-verified (146.0 DPS) | yes | Charged Scorpid Bracers (252593, -0.34 DPS) [crafted]; Devout Bracers (16697, -0.35 DPS) [dungeon]; Spitfire Bracers (20481, -2.66 DPS, sim-verified) [crafted] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 36.8 spell_power points (4.40 DPS) | yes | Magister's Gloves (16684, -0.39 DPS) [dungeon]; Devout Gloves (16692, -0.68 DPS) [dungeon]; Spitfire Gauntlets (20480, -3.80 DPS, sim-verified) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 32.4 spell_power points (3.87 DPS) | yes | Devout Belt (16696, -0.21 DPS) [dungeon]; Dreadmist Belt (16702, -0.24 DPS) [dungeon]; Elderwild Waistcord (279252, -0.46 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.7 spell_power points (7.00 DPS) | yes | Luminary Kilt (11823, -1.32 DPS) [dungeon]; Kilt of Elements (16668, -1.33 DPS) [dungeon]; Red Dragonscale Leggings (252603, -1.73 DPS, sim-verified) [crafted] |
| feet | Boots of Elements (16670) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (146.0 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Earthen Silk Slippers (254013, -0.03 DPS) [crafted]; Wolfrunner Shoes (13101, -1.62 DPS, sim-verified) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 30.2 spell_power points (3.61 DPS) | yes | Eye of Adaegus (5266, -1.52 DPS) [world_drop]; Philanthropist's Ring (281635, -1.90 DPS) [quest]; Choking Band (11868, -2.01 DPS) [quest] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 20.0 spell_power points (2.39 DPS) | yes | Eye of Adaegus (5266, -0.29 DPS) [world_drop]; Philanthropist's Ring (281635, -0.67 DPS) [quest]; Choking Band (11868, -0.79 DPS) [quest] |
| trinket1 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (146.0 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.25 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (146.0 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (146.0 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -6.59 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Bindings of Elements; hands: Gauntlets of Elements; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Boots of Elements; finger1: Maiden's Circle; finger2: Band of the Hierophant; trinket1: Evonice's Landin' Pilla; trinket2: Frozen Heart of the Mountain; main_hand: Hand of Edward the Odd

No-known-source sample (15 of 1856, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (dwarf, 5510331300001351-055000000000000000-0530020000000000)

Set DPS (verified): 486.2. Weights run: 2.1s. Verify run: 4.0s. 1856 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.500 ± 0.019, crit=0.369 ± 0.008 per rating point (14 rating = 1%, 5.167 per %), hit=0.683 ± 0.038 per rating point (10 rating = 1%, 6.831 per %), spell_haste=not significant (1.762 ± 0.464), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.645 ± 0.001, spirit=not significant (-0.032 ± 0.208), mp5=not significant (0.827 ± 0.284)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 40.8 spell_power points (18.12 DPS) | yes | Soothsayer's Headdress (17740, -4.58 DPS) [dungeon]; Red Mageweave Headband (10033, -5.25 DPS) [crafted]; Living Crown (252561, -5.35 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.0 spell_power points (7.54 DPS) | yes | Glowing Eye of Mordresh (10769, -3.45 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.55 DPS) [quest]; Triune Amulet (7722, -3.99 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.7 spell_power points (12.72 DPS) | yes | Ironfeather Shoulders (15067, -1.63 DPS) [crafted]; Rotgrip Mantle (17732, -2.96 DPS) [dungeon]; Argent Shoulders (19059, -4.69 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.0 spell_power points (7.54 DPS) | yes | Cloak of Earth and Sky (279270, -3.25 DPS) [crafted]; Big Voodoo Cloak (8216, -3.33 DPS) [crafted]; Runecloth Cloak (13860, -3.83 DPS, sim-verified) [crafted] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | 33.0 spell_power points (14.64 DPS) | yes | Acumen Robes (17775, -1.77 DPS) [quest]; Fernpulse Jerkin (15786, -2.65 DPS, sim-verified) [quest]; Ironfeather Breastplate (15066, -2.66 DPS) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 16.0 spell_power points (7.10 DPS) | yes | Runic Leather Bracers (15092, -0.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.11 DPS) [crafted]; Skycaller's Mail Bracers (252571, -1.11 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 30.5 spell_power points (13.51 DPS) | yes | Storm Gauntlets (12632, -2.49 DPS) [crafted]; Earth Warder's Gloves (21318, -3.75 DPS) [quest]; Skycaller's Leather Gauntlets (252550, -4.42 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 22.0 spell_power points (9.76 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -1.33 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.77 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.2 spell_power points (23.15 DPS) | yes | Red Dragonscale Leggings (252603, -7.17 DPS) [crafted]; Luminary Kilt (11823, -8.95 DPS) [dungeon]; Pristine Scorpid Leggings (252606, -9.41 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 26.0 spell_power points (11.54 DPS) | yes | Greaves of Withering Despair (22240, -1.41 DPS) [dungeon]; Venomspew Footpads (275606, -1.77 DPS) [crafted]; Earthen Silk Slippers (254013, -2.40 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.5 spell_power points (9.54 DPS) | yes | Philanthropist's Ring (281635, -3.99 DPS) [quest]; Snake Hoop (6750, -5.77 DPS) [quest]; Ring of Power Regulation (273030, -5.78 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.77 DPS) | yes | Philanthropist's Ring (281635, -0.22 DPS) [quest]; Snake Hoop (6750, -2.00 DPS) [quest]; Ring of Power Regulation (273030, -2.00 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (486.2 DPS) | yes | Uther's Strength (11302, -2.66 DPS, sim-verified) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (486.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (486.2 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -14.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-verified (486.2 DPS) | yes | Totem of the Storm (272432, -2.35 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Vest of Elements; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; trinket2: Shard of the Splithooves; main_hand: Hand of Edward the Odd; ranged: Totem of Thunder

No-known-source sample (15 of 1856, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (orc, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 22.6. Weights run: 2.7s. Verify run: 1.6s. 247 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.577 ± 0.009, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.031 per %), hit=0.188 ± 0.007 per rating point (10 rating = 1%, 1.884 per %), spell_haste=not significant (-0.026 ± 0.112), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.872 ± 0.003, spirit=not significant (-0.023 ± 0.111), mp5=not significant (0.137 ± 0.151)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (0.75 DPS) | yes | Pristine Circlet (253949, -0.25 DPS) [crafted]; Trapper's Leather Hood (252505, -0.37 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.38 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.2 spell_power points (0.85 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.52 DPS) [crafted]; Slime-encrusted Pads (6461, -0.81 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -1.15 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Sanguine Cape (14376, -0.14 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.48 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 10.9 spell_power points (0.91 DPS) | yes | Filigreed Pristine Gown (253901, -0.25 DPS) [crafted]; Totemic Leather Armor (252435, -0.26 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.38 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 3.3 spell_power points (0.27 DPS) | yes | Owl Bracers (4796, -0.03 DPS) [vendor]; Bright Bracers (3647, -0.08 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.03 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.3 spell_power points (0.61 DPS) | yes | Serpent Gloves (5970, -0.03 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.11 DPS) [world]; Pristine Gloves (253913, -0.13 DPS) [crafted] |
| waist | Stormrider's Leather Belt (252432) (or Pristine Sash (253925)) | Leatherworking [crafted] | 6.3 spell_power points (0.53 DPS) | yes | Pristine Sash (253925, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 13.5 spell_power points (1.12 DPS) | yes | Filigreed Pristine Leggings (253937, -0.33 DPS) [crafted]; Totemic Leather Pants (252446, -0.35 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.38 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.3 spell_power points (0.78 DPS) | yes | Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.28 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.61 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.5 spell_power points (0.29 DPS) | yes | Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; Furlfeather Ring (286745, -0.21 DPS) [world]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.9 spell_power points (0.24 DPS) | yes | Furlfeather Ring (286745, -0.16 DPS) [world]; The 1 Ring (8350, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.33 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Crescent Staff (6505, -0.33 DPS) [quest]; Lesser Staff of the Spire (1300, -0.38 DPS) [world]; Twisted Chanter's Staff (890, -0.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Stormrider's Leather Belt; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 5510331300000000-000000000000000000-0000000000000000)

Set DPS (verified): 45.8. Weights run: 2.9s. Verify run: 1.8s. 451 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.630 ± 0.012, crit=0.091 ± 0.002 per rating point (14 rating = 1%, 1.270 per %), hit=0.247 ± 0.008 per rating point (10 rating = 1%, 2.467 per %), spell_haste=-2.228 ± 0.170, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.806 ± 0.002, spirit=0.435 ± 0.068, mp5=0.936 ± 0.131

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (45.8 DPS) | yes | Holy Shroud (2721, +0.00 DPS) [world_drop]; Enduring Cap (3020, +0.00 DPS) [world_drop]; Totemic Leather Helm (252456, -0.03 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-verified (45.8 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, +0.00 DPS) [world_drop]; Pendant of Myzrael (4614, -0.13 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.7 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, -0.19 DPS) [dungeon]; Magician's Mantle (12998, -0.44 DPS) [world_drop]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | sim-verified (45.8 DPS) | yes | Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Glowing Thresher Cape (6901, +0.00 DPS) [dungeon]; Windsong Drape (15468, +0.00 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | sim-verified (45.8 DPS) | yes | Beguiler Robes (7728, +0.00 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.14 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.8 spell_power points (1.30 DPS) | yes | Mindthrust Bracers (1974, -0.64 DPS) [dungeon]; Nightsky Wristbands (6407, -0.74 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.74 DPS) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.3 spell_power points (1.13 DPS) | yes | Jutebraid Gloves (10654, -0.12 DPS) [quest]; Gloves of Old (9395, -0.21 DPS) [world_drop]; Gloves of Meditation (4318, -0.24 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.5 spell_power points (1.71 DPS) | yes | Belt of Arugal (6392, -0.27 DPS) [dungeon]; Mender's Leather Belt (252523, -0.34 DPS) [crafted]; Moss Cinch (6911, -0.39 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (45.8 DPS) | yes | Gaze Dreamer Pants (6903, +0.00 DPS) [dungeon]; Wisdom's Leather Leggings (252519, -0.12 DPS) [crafted]; Dark Ritual Leggings (270031, -0.17 DPS) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.2 spell_power points (1.45 DPS) | yes | Acidic Walkers (9454, -0.15 DPS) [dungeon]; Glinteye Slippers (273024, -0.39 DPS) [dungeon]; Spidersilk Boots (4320, -0.40 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (1.01 DPS) | yes | Solid Gold Band (279845, -0.33 DPS) [quest]; Electrocutioner Lagnut (9447, -0.34 DPS) [dungeon]; Runemender's Seal (273031, -0.34 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.8 spell_power points (0.97 DPS) | yes | Solid Gold Band (279845, -0.29 DPS) [quest]; Electrocutioner Lagnut (9447, -0.30 DPS) [dungeon]; Runemender's Seal (273031, -0.30 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (45.8 DPS) | yes | Bloodscalp Channeling Staff (1998, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, +0.00 DPS) [dungeon]; Reef Axe (6905, +0.00 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 451, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (orc, 5510331300001350-010000000000000000-0000000000000000)

Set DPS (verified): 82.3. Weights run: 3.2s. Verify run: 2.1s. 688 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.936 ± 0.025, crit=0.247 ± 0.008 per rating point (14 rating = 1%, 3.459 per %), hit=0.431 ± 0.019 per rating point (10 rating = 1%, 4.305 per %), spell_haste=-2.455 ± 0.273, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.692 ± 0.002, spirit=not significant (0.632 ± 0.203), mp5=1.844 ± 0.324

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 29.6 spell_power points (3.19 DPS) | yes | Big Voodoo Mask (8201, -0.81 DPS) [crafted]; Electromagnetic Gigaflux Reactivator (9492, -0.91 DPS) [dungeon]; Whitemane's Chapeau (7720, -1.38 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.5 spell_power points (1.56 DPS) | yes | Necklace of Calisea (1714, -0.38 DPS) [dungeon]; Triune Amulet (7722, -0.74 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -0.88 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.2 spell_power points (2.07 DPS) | yes | Sheepshear Mantle (13115, -0.04 DPS) [world_drop]; Green Silken Shoulders (7057, -0.09 DPS) [crafted]; Mantle of Doan (7712, -0.16 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 10.8 spell_power points (1.16 DPS) | yes | Long Silken Cloak (4326, -0.01 DPS) [crafted]; Guardian Cloak (5965, -0.01 DPS) [crafted]; Blackforge Cape (6424, -0.08 DPS) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 35.3 spell_power points (3.81 DPS) | yes | Robe of the Magi (1716, -0.48 DPS) [world_drop]; Robe of Power (7054, -0.54 DPS) [crafted]; Beguiler Robes (7728, -1.08 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 14.5 spell_power points (1.57 DPS) | yes | Green Whelp Bracers (7386, -0.09 DPS) [crafted]; Guardian Leather Bracers (4260, -0.31 DPS) [crafted]; Radiant Silver Bracers (4545, -0.33 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 30.9 spell_power points (3.34 DPS) | yes | Dreamweave Gloves (10019, -1.10 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.14 DPS) [crafted]; Skycaller's Leather Gloves (252529, -1.24 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.0 spell_power points (2.27 DPS) | yes | Gilded Cord (254037, -0.19 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.31 DPS) [crafted]; Star Belt (4329, -0.59 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.2 spell_power points (2.72 DPS) | yes | Stoneweaver Leggings (9407, -0.34 DPS) [dungeon]; Misplaced Pantaloons (276201, -0.52 DPS) [vendor]; Warchief Kilt (7760, -1.33 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.4 spell_power points (3.07 DPS) | yes | Skycaller's Leather Shoes (252532, -0.62 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.62 DPS) [crafted]; Mender's Leather Shoes (252533, -1.02 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.7 spell_power points (1.58 DPS) | yes | Snake Hoop (6750, -0.43 DPS) [quest]; Welken Ring (5011, -0.60 DPS) [world_drop]; Voodoo Band (1996, -0.67 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 11.5 spell_power points (1.24 DPS) | yes | Snake Hoop (6750, -0.09 DPS) [quest]; Welken Ring (5011, -0.26 DPS) [world_drop]; Voodoo Band (1996, -0.33 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (82.3 DPS) | yes | Illusionary Rod (7713, -1.94 DPS) [dungeon]; Bloodscalp Channeling Staff (1998, -1.98 DPS) [world]; Mograine's Might (7723, -12.80 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Icy Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan

No-known-source sample (15 of 688, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (orc, 5510331300001350-030000000000000000-5030000000000000)

Set DPS (verified): 107.1. Weights run: 3.3s. Verify run: 2.8s. 879 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.801 ± 0.028, crit=0.298 ± 0.009 per rating point (14 rating = 1%, 4.174 per %), hit=0.491 ± 0.023 per rating point (10 rating = 1%, 4.913 per %), spell_haste=not significant (-1.235 ± 0.318), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.691 ± 0.002, spirit=not significant (0.576 ± 0.267), mp5=not significant (1.146 ± 0.335)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 39.6 spell_power points (4.57 DPS) | yes | Dreamweave Circlet (10041, -0.43 DPS) [crafted]; Red Mageweave Headband (10033, -0.53 DPS) [crafted]; Charged Scorpid Helm (252581, -1.04 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (107.1 DPS) | yes | Glowing Eye of Mordresh (10769, -0.01 DPS) [dungeon]; Triune Amulet (7722, -0.17 DPS) [dungeon]; Mindburst Medallion (11196, -1.10 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.6 spell_power points (4.11 DPS) | yes | Rockshard Pauldrons (9411, -1.12 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -1.38 DPS) [crafted]; Rotgrip Mantle (17732, -1.61 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.3 spell_power points (2.57 DPS) | yes | Deep Woodlands Cloak (19121, -0.35 DPS) [quest]; Runecloth Cloak (13860, -0.79 DPS) [crafted]; Big Voodoo Cloak (8216, -1.16 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | 40.8 spell_power points (4.71 DPS) | yes | Dreamweave Vest (10021, -0.87 DPS) [crafted]; Stone Guard's Pulsing Breastplate (220844, -0.96 DPS) [vendor]; Acumen Robes (17775, -1.46 DPS, sim-verified) [quest] |
| wrist | Skycaller's Leather Bracers (252542) (or Skycaller's Mail Bracers (252571)) | Leatherworking [crafted] | 18.5 spell_power points (2.13 DPS) | yes | Skycaller's Mail Bracers (252571, +0.00 DPS) [crafted]; Runic Leather Bracers (15092, -0.06 DPS) [crafted]; Nethergeld Cuffs (254061, -0.35 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 30.9 spell_power points (3.56 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.49 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.49 DPS) [crafted]; Dreamweave Gloves (10019, -0.65 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.6 spell_power points (2.95 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Tooled Leather Belt (252544, -0.37 DPS) [crafted]; Satyrmane Sash (17755, -0.42 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | PvP rank 10 · Stone Guard · Horde [vendor] | 32.5 spell_power points (3.75 DPS) | yes | Senior Designer's Pantaloons (11841, -0.52 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -0.55 DPS) [dungeon]; Jinxed Hoodoo Kilt (9474, -0.64 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.0 spell_power points (3.23 DPS) | yes | Greaves of Withering Despair (22240, -0.47 DPS) [dungeon]; Skycaller's Leather Boots (252471, -0.49 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.49 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.61 DPS) | yes | Eye of Adaegus (5266, -0.26 DPS) [world_drop]; Snake Hoop (6750, -0.46 DPS) [quest]; Ring of Power Regulation (273030, -0.53 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.50 DPS) | yes | Eye of Adaegus (5266, -0.15 DPS) [world_drop]; Snake Hoop (6750, -0.35 DPS) [quest]; Ring of Power Regulation (273030, -0.41 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.16 DPS) [quest]; Spire of Hakkar (10844, -0.30 DPS) [world]; Zum'rah's Vexing Cane (18082, -1.36 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Skycaller's Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (orc, 5510331300001350-030000000000000000-5033520000000000)

Set DPS (verified): 143.3. Weights run: 3.4s. Verify run: 4.9s. 1812 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.867 ± 0.034, crit=0.361 ± 0.010 per rating point (14 rating = 1%, 5.052 per %), hit=0.635 ± 0.027 per rating point (10 rating = 1%, 6.346 per %), spell_haste=-1.861 ± 0.401, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.701 ± 0.002, spirit=not significant (1.028 ± 0.322), mp5=2.342 ± 0.404

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 55.0 spell_power points (6.57 DPS) | yes | Dreadmist Mask (16698, -1.04 DPS) [dungeon]; Warlord's Mail Helm (231663, -1.15 DPS) [pvp]; Blue Dragonscale Helm (252604, -1.69 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.7 spell_power points (2.47 DPS) | yes | Horizon Choker (13085, -0.53 DPS) [world_drop]; Lady Alizabeth's Pendant (13002, -0.63 DPS) [world_drop]; Mindburst Medallion (11196, -1.51 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 40.6 spell_power points (4.84 DPS) | yes | Warlord's Mail Spaulders (231659, -0.08 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.86 DPS) [crafted]; Shroud of the Nathrezim (18720, -6.83 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.4 spell_power points (3.03 DPS) | yes | Deep Woodlands Cloak (19121, -0.66 DPS) [quest]; Runecloth Cloak (13860, -1.13 DPS) [crafted]; Cloak of Earth and Sky (279270, -1.47 DPS, sim-verified) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | 60.2 spell_power points (7.19 DPS) | yes | Ghostweave Vest (14141, -1.42 DPS) [crafted]; Mooncloth Vest (14138, -1.71 DPS) [crafted]; Feathered Breastplate (8349, -6.72 DPS, sim-verified) [crafted] |
| wrist | Spitfire Bracers (20481) | Leatherworking [crafted] | 25.1 spell_power points (2.99 DPS) | yes | Charged Scorpid Bracers (252593, -0.46 DPS) [crafted]; Devout Bracers (16697, -0.46 DPS) [dungeon]; Bindings of Elements (16671, -1.87 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 36.8 spell_power points (4.40 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Spitfire Gauntlets (20480, -0.37 DPS) [crafted]; Magister's Gloves (16684, -0.39 DPS) [dungeon] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 32.4 spell_power points (3.87 DPS) | yes | General's Mail Waistband (16575, -0.05 DPS) [pvp]; Dreadmist Belt (16702, -0.24 DPS) [dungeon]; Devout Belt (16696, -1.33 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.7 spell_power points (7.00 DPS) | yes | General's Mail Leggings (231664, -0.57 DPS) [pvp]; Luminary Kilt (11823, -1.32 DPS) [dungeon]; Red Dragonscale Leggings (252603, -2.58 DPS, sim-verified) [crafted] |
| feet | Wolfrunner Shoes (13101) | World drop [world_drop] | 33.8 spell_power points (4.04 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Boots of Elements (16670, -0.28 DPS) [dungeon]; Mooncloth Boots (15802, -1.56 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 30.2 spell_power points (3.61 DPS) | yes | Eye of Adaegus (5266, -1.52 DPS) [world_drop]; Philanthropist's Ring (281635, -1.90 DPS) [quest]; Ring of Power Regulation (273030, -2.06 DPS) [dungeon] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 20.0 spell_power points (2.39 DPS) | yes | Philanthropist's Ring (281635, -0.67 DPS) [quest]; Ring of Power Regulation (273030, -0.83 DPS) [dungeon]; Eye of Adaegus (5266, -2.28 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (143.3 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (143.3 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (143.3 DPS) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -8.26 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Spitfire Bracers; hands: Gauntlets of Elements; waist: Wisdom of the Timbermaw; legs: Ironfeather Leggings; feet: Wolfrunner Shoes; finger1: Maiden's Circle; finger2: Band of the Hierophant; trinket2: Mark of the Chosen; main_hand: Hand of Edward the Odd

No-known-source sample (15 of 1812, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (orc, 5510331300001351-055000000000000000-0530020000000000)

Set DPS (verified): 492.6. Weights run: 2.1s. Verify run: 5.1s. 1812 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.500 ± 0.019, crit=0.369 ± 0.008 per rating point (14 rating = 1%, 5.167 per %), hit=0.683 ± 0.038 per rating point (10 rating = 1%, 6.831 per %), spell_haste=not significant (1.762 ± 0.464), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.645 ± 0.001, spirit=not significant (-0.032 ± 0.208), mp5=not significant (0.827 ± 0.284)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 40.8 spell_power points (18.12 DPS) | yes | Warlord's Mail Helm (231663, -1.18 DPS) [pvp]; Champion's Mail Helm (227162, -3.40 DPS) [pvp]; Living Crown (252561, -4.06 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.0 spell_power points (7.54 DPS) | yes | Glowing Eye of Mordresh (10769, -2.65 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.55 DPS) [quest]; Triune Amulet (7722, -3.99 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (492.6 DPS) | yes | Champion's Mail Spaulders (227160, +0.00 DPS) [vendor]; Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -8.16 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.0 spell_power points (7.54 DPS) | yes | Deep Woodlands Cloak (19121, -0.22 DPS) [quest]; Runecloth Cloak (13860, -1.77 DPS) [crafted]; Cloak of Earth and Sky (279270, -3.25 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (492.6 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Vest of Elements (16666, -7.57 DPS, sim-verified) [dungeon] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 16.0 spell_power points (7.10 DPS) | yes | Runic Leather Bracers (15092, -0.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.11 DPS) [crafted]; Skycaller's Mail Bracers (252571, -1.11 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 30.5 spell_power points (13.51 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, -1.68 DPS) [vendor]; Storm Gauntlets (12632, -2.49 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 22.0 spell_power points (9.76 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -1.33 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.77 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.2 spell_power points (23.15 DPS) | yes | General's Mail Leggings (231664, -3.11 DPS) [pvp]; Red Dragonscale Leggings (252603, -7.17 DPS) [crafted]; Pristine Scorpid Leggings (252606, -7.58 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 26.0 spell_power points (11.54 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Earthen Silk Slippers (254013, -0.89 DPS) [crafted]; Greaves of Withering Despair (22240, -1.41 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.5 spell_power points (9.54 DPS) | yes | Philanthropist's Ring (281635, -3.99 DPS) [quest]; Snake Hoop (6750, -5.77 DPS) [quest]; Ring of Power Regulation (273030, -5.78 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.77 DPS) | yes | Philanthropist's Ring (281635, -0.22 DPS) [quest]; Snake Hoop (6750, -2.00 DPS) [quest]; Ring of Power Regulation (273030, -2.00 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (492.6 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (492.6 DPS) | yes | Uther's Strength (11302, -0.07 DPS) [world_drop]; Shard of the Splithooves (10659, -0.52 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (492.6 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -11.59 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Enchanting [crafted] | sim-verified (492.6 DPS) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hand of Edward the Odd; ranged: Totem of Thunder

No-known-source sample (15 of 1812, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

