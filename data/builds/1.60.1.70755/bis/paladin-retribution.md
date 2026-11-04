# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 41.7. Weights run: 2.3s. Verify run: 3.9s. 262 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.484 ± 0.026, crit=0.841 ± 0.023 per rating point (14 rating = 1%, 11.774 per %), hit=1.314 ± 0.065 per rating point (10 rating = 1%, 13.136 per %), melee_haste=8.238 ± 0.564, intellect=2.782 ± 0.044, spell_power=1.772 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | sim-verified (41.7 DPS) | yes | Crusader's Silvered Chain Helm (250532, -0.02 DPS) [crafted]; Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Trapper's Leather Hood (252505, -0.76 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.0 attack_power points (0.13 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS) [quest]; Tarnished Locket (279870, -0.39 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 33.9 attack_power points (1.48 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.61 DPS) [crafted]; Serpent's Shoulders (5404, -1.15 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.17 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 11.9 attack_power points (0.52 DPS) | yes | Sanguine Cape (14376, -0.03 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.15 DPS) [crafted] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (41.7 DPS) | yes | Crusader's Chain Shirt (250492, +0.00 DPS) [crafted]; Stormrider's Leather Armor (252492, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -1.53 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 13.9 attack_power points (0.61 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Bright Bracers (3647, -0.12 DPS) [world_drop]; Bravo's Armbands (270015, -0.15 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.7 DPS) | yes | Acolyte's Gloves (250511, +0.00 DPS) [crafted]; Stormrider's Leather Gloves (252498, +0.00 DPS) [crafted]; Serpent Gloves (5970, -1.94 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (41.7 DPS) | yes | Keller's Girdle (2911, +0.00 DPS) [world_drop]; Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Cobrahn's Grasp (6460, -2.38 DPS, sim-verified) [dungeon] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 41.5 attack_power points (1.81 DPS) | yes | Stormrider's Leather Pants (252502, -0.31 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.42 DPS) [crafted]; Trapper's Leather Pants (252501, -0.53 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (41.7 DPS) | yes | Spidersilk Boots (4320, +0.00 DPS) [crafted]; Acolyte's Boots (250506, +0.00 DPS) [crafted]; Stormrider's Leather Boots (252443, -1.93 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 19.7 attack_power points (0.86 DPS) | yes | Volcanic Rock Ring (12053, -0.49 DPS) [world_drop]; Demon Band (12054, -0.51 DPS) [world_drop]; The 1 Ring (8350, -0.58 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 14.4 attack_power points (0.63 DPS) | yes | Demon Band (12054, -0.28 DPS) [world_drop]; The 1 Ring (8350, -0.36 DPS) [world]; Volcanic Rock Ring (12053, -0.46 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 375.4 attack_power points (16.34 DPS) | yes | Smite's Mighty Hammer (7230, -3.11 DPS) [dungeon]; Sword of the Fallen (286977, -3.14 DPS, sim-verified) [dungeon]; Duskbringer (2205, -3.53 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Blackened Defias Armor; wrist: Owl Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Blackened Defias Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; main_hand: Verigan's Fist

No-known-source sample (15 of 262, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers; 14148 Crystalline Cuffs

### Band 30 (human, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 85.8. Weights run: 2.4s. Verify run: 2.4s. 484 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.267 ± 0.023, crit=0.964 ± 0.025 per rating point (14 rating = 1%, 13.497 per %), hit=1.585 ± 0.069 per rating point (10 rating = 1%, 15.846 per %), melee_haste=9.455 ± 0.568, intellect=0.752 ± 0.017, spell_power=2.062 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 50.9 attack_power points (2.45 DPS) | yes | Tusken Helm (6686, -1.20 DPS) [dungeon]; Totemic Leather Helm (252456, -1.26 DPS) [crafted]; Barbaric Iron Helm (7915, -1.63 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 16.5 attack_power points (0.80 DPS) | yes | Kaleidoscope Chain (13084, -0.02 DPS) [world_drop]; River Pride Choker (13087, -0.41 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.51 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 25.3 attack_power points (1.22 DPS) | yes | Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop]; Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.28 DPS) [crafted] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 16.1 attack_power points (0.78 DPS) | yes | Hawkeye's Cloak (14593, -0.06 DPS) [world_drop]; Hillman's Cloak (3719, -0.28 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.28 DPS) [quest] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 36.7 attack_power points (1.77 DPS) | yes | Hard Gold Cuirass (250533, -0.11 DPS) [crafted]; Totemic Leather Tunic (252451, -0.12 DPS) [crafted]; Avenger's Armor (1488, -0.33 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.6 attack_power points (0.90 DPS) | yes | Pugilist Bracers (4438, -0.12 DPS) [dungeon]; Yorgen Bracers (13012, -0.13 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.14 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 22.0 attack_power points (1.06 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.00 DPS) [world]; Insignia Gloves (6408, -0.02 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.05 DPS) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 36.7 attack_power points (1.77 DPS) | yes | Skirmisher's Leather Belt (252461, -0.06 DPS) [crafted]; Prowler's Leather Belt (252459, -0.39 DPS) [crafted]; Officer's Belt (250556, -0.44 DPS) [crafted] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 40.8 attack_power points (1.97 DPS) | yes | Totemic Leather Leggings (252458, -0.09 DPS) [crafted]; Totemic Leather Pants (252446, -0.26 DPS) [crafted]; Ferine Leggings (6690, -0.28 DPS) [dungeon] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 28.4 attack_power points (1.37 DPS) | yes | Disjointed Shoes (277226, -0.30 DPS) [quest]; Crusader's Boots (250507, -0.39 DPS) [crafted]; Totemic Leather Boots (252442, -0.42 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.96 DPS) | yes | Seal of Wrynn (2933, -0.34 DPS) [quest]; Ring of Power Regulation (273030, -0.36 DPS) [dungeon]; Minor Channeling Ring (1449, -0.39 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 16.1 attack_power points (0.78 DPS) | yes | Seal of Wrynn (2933, -0.16 DPS) [quest]; Ring of Power Regulation (273030, -0.18 DPS) [dungeon]; Minor Channeling Ring (1449, -0.21 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (85.8 DPS) | yes | Felbough Arm (273037, -0.59 DPS) [world]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -20.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Sergeant Major's Cape; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; hands: Serpent Gloves; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Snake Hoop; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 484, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 126.9. Weights run: 2.4s. Verify run: 2.6s. 672 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=1.525 ± 0.030, crit=1.475 ± 0.036 per rating point (14 rating = 1%, 20.647 per %), hit=1.985 ± 0.088 per rating point (10 rating = 1%, 19.849 per %), melee_haste=11.125 ± 0.653, intellect=1.748 ± 0.018, spell_power=2.065 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 51.3 attack_power points (2.67 DPS) | yes | Papal Fez (9431, -0.27 DPS) [dungeon]; Icemetal Barbute (10763, -0.28 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.69 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.4 attack_power points (1.48 DPS) | yes | Kaleidoscope Chain (13084, -0.37 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.48 DPS, sim-verified) [world]; Triune Amulet (7722, -0.62 DPS) [dungeon] |
| shoulder | Shining Mithril Pauldrons (250541) | Blacksmithing [crafted] | 43.1 attack_power points (2.25 DPS) | yes | Green Silken Shoulders (7057, -0.38 DPS) [crafted]; Bloodmage Mantle (7684, -0.46 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.54 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 31.9 attack_power points (1.66 DPS) | yes | Long Silken Cloak (4326, -0.56 DPS) [crafted]; Guardian Cloak (5965, -0.56 DPS) [crafted]; Hawkeye's Cloak (14593, -0.79 DPS) [world_drop] |
| chest | Shining Mithril Breastplate (250540) | Blacksmithing [crafted] | 57.5 attack_power points (2.99 DPS) | yes | Robe of the Magi (1716, -0.08 DPS) [world_drop]; Dreamweave Vest (10021, -0.24 DPS) [crafted]; Icemail Jerkin (1981, -2.06 DPS, sim-verified) [world_drop] |
| wrist | Hard Gold Bracers (250535) | Blacksmithing [crafted] | 24.6 attack_power points (1.28 DPS) | yes | Guardian Leather Bracers (4260, -0.09 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.18 DPS) [crafted]; Turtle Scale Bracers (8198, -0.18 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 49.5 attack_power points (2.58 DPS) | yes | Skirmisher's Leather Gloves (252528, -0.44 DPS) [crafted]; Dreamweave Gloves (10019, -0.47 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.49 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 40.7 attack_power points (2.12 DPS) | yes | Skirmisher's Leather Belt (252461, -0.05 DPS) [crafted]; Justicar's Belt (250560, -0.19 DPS) [crafted]; Prowler's Leather Belt (252459, -0.32 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 49.9 attack_power points (2.60 DPS) | yes | Ferine Leggings (6690, -0.15 DPS) [dungeon]; Totemic Leather Leggings (252458, -0.24 DPS) [crafted]; Hard Gold Leggings (250536, -0.25 DPS) [crafted] |
| feet | Skirmisher's Mail Boots (252564) | Leatherworking [crafted] | 49.8 attack_power points (2.59 DPS) | yes | Earthen Silk Slippers (254013, -0.01 DPS) [crafted]; Justicar's Boots (250550, -0.24 DPS) [crafted]; Shining Mithril Boots (250542, -0.35 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 29.4 attack_power points (1.53 DPS) | yes | Thunderbrow Ring (13097, -0.44 DPS) [world_drop]; Mark of Kern (2262, -0.49 DPS) [dungeon]; Assault Band (13095, -0.49 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 21.1 attack_power points (1.10 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Mark of Kern (2262, -0.06 DPS) [dungeon]; Assault Band (13095, -0.06 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (126.9 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -18.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** neck: Glowing Eye of Mordresh; shoulder: Shining Mithril Pauldrons; back: Sergeant Major's Cape; chest: Shining Mithril Breastplate; wrist: Hard Gold Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Skirmisher's Mail Boots; finger1: Philanthropist's Ring

No-known-source sample (15 of 672, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (human, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 172.2. Weights run: 2.6s. Verify run: 3.0s. 869 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.485 ± 0.029, crit=1.769 ± 0.038 per rating point (14 rating = 1%, 24.766 per %), hit=2.227 ± 0.098 per rating point (10 rating = 1%, 22.269 per %), melee_haste=11.213 ± 0.984, intellect=4.041 ± 0.070, spell_power=2.485 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 128.0 attack_power points (6.78 DPS) | yes | Knight-Lieutenant's Lamellar Helm (220819, -0.61 DPS) [vendor]; Eye of Theradras (17715, -1.40 DPS) [dungeon]; Soothsayer's Headdress (17740, -2.04 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 56.6 attack_power points (3.00 DPS) | yes | Mindburst Medallion (11196, -0.92 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.07 DPS, sim-verified) [dungeon]; Archmagister's Faceted Pendant (273049, -1.08 DPS) [world] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 118.1 attack_power points (6.25 DPS) | yes | Knight-Lieutenant's Lamellar Pauldrons (220818, -1.25 DPS) [vendor]; Rockshard Pauldrons (9411, -1.54 DPS) [dungeon]; Rotgrip Mantle (17732, -2.57 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 59.0 attack_power points (3.13 DPS) | yes | Big Voodoo Cloak (8216, -0.54 DPS) [crafted]; Sergeant Major's Cape (16336, -0.67 DPS) [pvp]; Runecloth Cloak (13860, -1.16 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 128.0 attack_power points (6.78 DPS) | yes | Knight's Lamellar Chestplate (220815, -0.94 DPS) [vendor]; Spiderfang Carapace (11633, -1.22 DPS, sim-verified) [dungeon]; Runecloth Robe (13858, -1.69 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 65.3 attack_power points (3.46 DPS) | yes | Skycaller's Leather Bracers (252542, -0.64 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.64 DPS) [crafted]; Skirmisher's Mail Bracers (252572, -0.93 DPS, sim-verified) [crafted] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | 94.7 attack_power points (5.01 DPS) | yes | Skirmisher's Mail Gauntlets (252586, -0.65 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.98 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.98 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 91.7 attack_power points (4.86 DPS) | yes | Skycaller's Leather Waistguard (252476, -0.18 DPS) [crafted]; Skycaller's Mail Belt (252589, -0.18 DPS) [crafted]; Deathmage Sash (10771, -0.72 DPS) [dungeon] |
| legs | Knight's Lamellar Legplates (220816) | Captain Dirgehammer [vendor] | 110.3 attack_power points (5.84 DPS) | yes | Imperial Red Pants (8251, -1.45 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -1.60 DPS) [dungeon]; Red Mageweave Pants (10009, -3.35 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 90.0 attack_power points (4.77 DPS) | yes | Skycaller's Mail Sabatons (252577, -0.44 DPS) [crafted]; Sergeant Major's Lamellar Boots (220814, -0.80 DPS) [vendor]; Skycaller's Leather Boots (252471, -2.50 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 45.1 attack_power points (2.39 DPS) | yes | Snake Hoop (6750, -0.53 DPS) [quest]; Band of the Unicorn (7553, -0.68 DPS) [world_drop]; Voodoo Band (1996, -0.89 DPS) [world] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.3 attack_power points (2.24 DPS) | yes | Band of the Unicorn (7553, -0.53 DPS) [world_drop]; Voodoo Band (1996, -0.74 DPS) [world]; Snake Hoop (6750, -2.29 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (172.2 DPS) | yes | Uther's Strength (11302, -0.81 DPS, sim-verified) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (172.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (172.2 DPS) | yes | Blight (7959, -1.72 DPS, sim-verified) [crafted]; Warmonger (13052, -3.34 DPS) [world_drop]; Stoneraven (13059, -4.80 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Stonerender Gauntlets; waist: Dawnspire Cord; legs: Knight's Lamellar Legplates; feet: Greaves of Withering Despair; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 869, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (human, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 243.6. Weights run: 2.6s. Verify run: 8.3s. 1792 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.005 ± 0.048, crit=2.828 ± 0.069 per rating point (14 rating = 1%, 39.592 per %), hit=4.042 ± 0.174 per rating point (10 rating = 1%, 40.418 per %), melee_haste=20.373 ± 1.984, intellect=6.329 ± 0.108, spell_power=2.510 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 322.2 attack_power points (17.44 DPS) | yes | Field Marshal's Lamellar Faceguard (16474, -2.55 DPS) [vendor]; Field Marshal's Lamellar Headguard (231648, -4.69 DPS) [vendor]; Magister's Crown (16686, -5.24 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 93.4 attack_power points (5.06 DPS) | yes | Opaline Medallion (16623, -0.60 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.36 DPS) [dungeon]; Horizon Choker (13085, -2.85 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 164.5 attack_power points (8.91 DPS) | yes | Field Marshal's Lamellar Shoulders (231651, +0.00 DPS) [vendor]; Magister's Mantle (16689, -0.28 DPS) [dungeon]; Ironfeather Shoulders (15067, -2.46 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 84.3 attack_power points (4.56 DPS) | yes | Spritecaster Cape (11623, -0.60 DPS) [dungeon]; Imperial Red Cloak (8248, -0.79 DPS) [world_drop]; Runecloth Cloak (13860, -2.13 DPS, sim-verified) [crafted] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (243.6 DPS) | yes | Field Marshal's Lamellar Chestplate (231641, +0.00 DPS) [pvp]; Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Magister's Robes (16688, -5.17 DPS, sim-verified) [dungeon] |
| wrist | Magiskull Cuffs (13107) | World drop [world_drop] | 107.6 attack_power points (5.82 DPS) | yes | Tranquil Wristguards (279256, -0.49 DPS) [crafted]; Spitfire Bracers (20481, -0.68 DPS) [crafted]; Runecloth Cuffs (254123, -0.90 DPS) [crafted] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | sim-verified (243.6 DPS) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Mourning Handguards (279257, -4.94 DPS, sim-verified) [crafted] |
| waist | Cord of Elements (16673) | Blackrock Spire: Smolderthorn Witch Doctor [dungeon] | sim-verified (243.6 DPS) | yes | Elunarian Belt (14465, +0.00 DPS) [world_drop]; Magister's Belt (16685, +0.00 DPS) [dungeon]; Lightforge Belt (16723, -3.55 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-verified (243.6 DPS) | yes | Marshal's Lamellar Legplates (16475, +0.00 DPS) [vendor]; Marshal's Lamellar Leggings (231647, +0.00 DPS) [vendor]; Red Dragonscale Leggings (252603, -4.21 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 131.3 attack_power points (7.11 DPS) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 89.5 attack_power points (4.84 DPS) | yes | Band of the Hierophant (13096, -1.42 DPS) [world_drop]; Frigid Ring (18679, -1.42 DPS) [world]; Blackstone Ring (17713, -1.57 DPS) [dungeon] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 75.9 attack_power points (4.11 DPS) | yes | Band of the Hierophant (13096, -0.69 DPS) [world_drop]; Blackstone Ring (17713, -0.84 DPS) [dungeon]; Frigid Ring (18679, -1.33 DPS, sim-verified) [world] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (243.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (243.6 DPS) | yes | Uther's Strength (11302, -1.55 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (243.6 DPS) | yes | Blade of Hanna (2801, +0.00 DPS, sim-verified) [world_drop]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-verified (243.6 DPS) | yes | Libram of Fervor (23203, -1.87 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Sergeant Major's Cape; chest: Vest of Elements; wrist: Magiskull Cuffs; hands: Gauntlets of Elements; waist: Cord of Elements; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Glowing Crystal Ring; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive

No-known-source sample (15 of 1792, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (human, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 661.1. Weights run: 2.7s. Verify run: 6.8s. 1792 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.006, agility=1.852 ± 0.037, crit=2.406 ± 0.050 per rating point (14 rating = 1%, 33.683 per %), hit=5.216 ± 0.321 per rating point (10 rating = 1%, 52.157 per %), melee_haste=23.914 ± 3.643, intellect=2.596 ± 0.183, spell_power=1.387 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 215.2 attack_power points (32.60 DPS) | yes | Stalwart Helm (250599, -4.00 DPS, sim-verified) [crafted]; Black Dragonscale Helm (252605, -6.67 DPS) [crafted]; Field Marshal's Lamellar Faceguard (16474, -6.69 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 42.6 attack_power points (6.45 DPS) | yes | Horizon Choker (13085, -0.95 DPS) [world_drop]; Skibi's Pendant (13089, -0.97 DPS) [world_drop]; Opaline Medallion (16623, -1.34 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 103.9 attack_power points (15.74 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS, sim-verified) [crafted]; Field Marshal's Lamellar Shoulders (231651, +0.00 DPS) [vendor]; Shroud of the Nathrezim (18720, -2.88 DPS) [dungeon] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 54.9 attack_power points (8.32 DPS) | yes | Cape of the Black Baron (13340, -1.08 DPS) [dungeon]; Spritecaster Cape (11623, -3.02 DPS) [dungeon]; Cloak of Flames (3475, -8.42 DPS, sim-verified) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+8.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Lamellar Breastplate (231649, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -2.25 DPS) [dungeon]; Dawn Armor (252483, -8.21 DPS, sim-verified) [crafted] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 64.7 attack_power points (9.79 DPS) | yes | Spitfire Bracers (20481, -2.05 DPS) [crafted]; Vambraces of the Sadist (13400, -2.13 DPS) [dungeon]; Lordly Armguards (13135, -2.51 DPS) [world_drop] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-verified (+8.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Lamellar Gloves (231643, +0.00 DPS) [pvp]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Gauntlets of Elements (16672, -8.30 DPS, sim-verified) [dungeon] |
| waist | Lightforge Belt (16723) | Stratholme: Rockwing Gargoyle [dungeon] | 98.9 attack_power points (14.98 DPS) | yes | Brigam Girdle (13142, -1.95 DPS) [dungeon]; Might of the Timbermaw (19044, -3.36 DPS) [crafted]; Marshal's Lamellar Belt (16470, -3.52 DPS) [pvp] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Lamellar Legplates (16475, +0.00 DPS) [vendor]; Marshal's Lamellar Leggings (231647, +0.00 DPS) [vendor]; Champion's Legplates (250598, -0.48 DPS) [crafted] |
| feet | Depleted Thorium Sabatons (279263) | Blacksmithing [crafted] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Lamellar Boots (16472, +0.00 DPS) [vendor]; Marshal's Lamellar Sabatons (231652, +0.00 DPS) [vendor]; Greaves of Withering Despair (22240, -7.56 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 72.2 attack_power points (10.93 DPS) | yes | Band of the Penitent (13217, -5.83 DPS) [quest]; Glowing Crystal Ring (18402, -6.21 DPS) [quest]; Philanthropist's Ring (281635, -6.86 DPS) [quest] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 43.1 attack_power points (6.53 DPS) | yes | Band of the Penitent (13217, -1.43 DPS) [quest]; Glowing Crystal Ring (18402, -1.82 DPS) [quest]; Philanthropist's Ring (281635, -2.47 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -5.85 DPS) [world_drop] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Legionite Glaive (250619, -33.08 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Archlight Talisman; shoulder: Lightforge Spaulders; back: Sergeant Major's Cape; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; waist: Lightforge Belt; legs: Ironfeather Leggings; feet: Depleted Thorium Sabatons; finger1: Blackstone Ring; finger2: Maiden's Circle; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Arcanite Champion

No-known-source sample (15 of 1792, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (undead, 00000000000000000-00000000000000000-05231000000000000)

Set DPS (verified): 39.8. Weights run: 2.3s. Verify run: 3.8s. 252 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.484 ± 0.026, crit=0.841 ± 0.023 per rating point (14 rating = 1%, 11.774 per %), hit=1.314 ± 0.065 per rating point (10 rating = 1%, 13.136 per %), melee_haste=8.238 ± 0.564, intellect=2.782 ± 0.044, spell_power=1.772 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | sim-verified (39.8 DPS) | yes | Crusader's Silvered Chain Helm (250532, -0.02 DPS) [crafted]; Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Trapper's Leather Hood (252505, -0.39 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.0 attack_power points (0.13 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 33.9 attack_power points (1.48 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.61 DPS) [crafted]; Serpent's Shoulders (5404, -1.15 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.17 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 11.9 attack_power points (0.52 DPS) | yes | Sanguine Cape (14376, -0.03 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.15 DPS) [crafted] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (39.8 DPS) | yes | Crusader's Chain Shirt (250492, +0.00 DPS) [crafted]; Stormrider's Leather Armor (252492, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -1.41 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 13.9 attack_power points (0.61 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Bright Bracers (3647, -0.12 DPS) [world_drop]; Patterned Bronze Bracers (2868, -0.17 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.8 DPS) | yes | Acolyte's Gloves (250511, +0.00 DPS) [crafted]; Stormrider's Leather Gloves (252498, +0.00 DPS) [crafted]; Serpent Gloves (5970, -1.94 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (39.8 DPS) | yes | Keller's Girdle (2911, +0.00 DPS) [world_drop]; Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Cobrahn's Grasp (6460, -2.32 DPS, sim-verified) [dungeon] |
| legs | Totemic Leather Pants (252446) | Leatherworking [crafted] | 41.5 attack_power points (1.81 DPS) | yes | Stormrider's Leather Pants (252502, -0.31 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.42 DPS) [crafted]; Trapper's Leather Pants (252501, -0.62 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (39.8 DPS) | yes | Spidersilk Boots (4320, +0.00 DPS) [crafted]; Acolyte's Boots (250506, +0.00 DPS) [crafted]; Stormrider's Leather Boots (252443, -1.55 DPS, sim-verified) [crafted] |
| finger1 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 19.9 attack_power points (0.87 DPS) | yes | Volcanic Rock Ring (12053, -0.50 DPS) [world_drop]; Demon Band (12054, -0.52 DPS) [world_drop]; The 1 Ring (8350, -0.59 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 19.7 attack_power points (0.86 DPS) | yes | Volcanic Rock Ring (12053, -0.49 DPS) [world_drop]; Demon Band (12054, -0.51 DPS) [world_drop]; The 1 Ring (8350, -0.58 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | sim-verified (39.8 DPS) | yes | Sword of the Fallen (286977, +0.00 DPS) [dungeon]; Hammerbone (270018, -0.16 DPS) [quest]; Duskbringer (2205, -0.41 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Blackened Defias Armor; wrist: Owl Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Totemic Leather Pants; feet: Blackened Defias Boots; finger1: Loop of Sacrifice; finger2: Lavishly Jeweled Ring; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 252, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band

### Band 30 (undead, 00000000000000000-00000000000000000-05233301001300000)

Set DPS (verified): 87.2. Weights run: 2.4s. Verify run: 2.5s. 473 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.267 ± 0.023, crit=0.964 ± 0.025 per rating point (14 rating = 1%, 13.497 per %), hit=1.585 ± 0.069 per rating point (10 rating = 1%, 15.846 per %), melee_haste=9.455 ± 0.568, intellect=0.752 ± 0.017, spell_power=2.062 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 50.9 attack_power points (2.45 DPS) | yes | Tusken Helm (6686, -1.20 DPS) [dungeon]; Totemic Leather Helm (252456, -1.26 DPS) [crafted]; Barbaric Iron Helm (7915, -1.63 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 16.5 attack_power points (0.80 DPS) | yes | Kaleidoscope Chain (13084, -0.02 DPS) [world_drop]; River Pride Choker (13087, -0.41 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.65 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 25.3 attack_power points (1.22 DPS) | yes | Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop]; Death Speaker Mantle (6685, -0.23 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.28 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.72 DPS) | yes | Hillman's Cloak (3719, -0.22 DPS) [crafted]; Windsong Drape (15468, -0.22 DPS) [quest]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop] |
| chest | Crusader's Silvered Chain Shirt (250522) | Blacksmithing [crafted] | 36.7 attack_power points (1.77 DPS) | yes | Totemic Leather Tunic (252451, -0.12 DPS) [crafted]; Avenger's Armor (1488, -0.33 DPS) [dungeon]; Hard Gold Cuirass (250533, -0.39 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 18.6 attack_power points (0.90 DPS) | yes | Pugilist Bracers (4438, -0.12 DPS) [dungeon]; Yorgen Bracers (13012, -0.13 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.14 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 22.0 attack_power points (1.06 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.00 DPS) [world]; Insignia Gloves (6408, -0.02 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.05 DPS) [world_drop] |
| waist | Justicar's Belt (250560) | Blacksmithing [crafted] | 36.7 attack_power points (1.77 DPS) | yes | Skirmisher's Leather Belt (252461, -0.37 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -0.39 DPS) [crafted]; Officer's Belt (250556, -0.44 DPS) [crafted] |
| legs | Crusader's Silvered Chain Leggings (250527) | Blacksmithing [crafted] | 40.8 attack_power points (1.97 DPS) | yes | Totemic Leather Pants (252446, -0.26 DPS) [crafted]; Ferine Leggings (6690, -0.28 DPS) [dungeon]; Totemic Leather Leggings (252458, -0.46 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | 28.4 attack_power points (1.37 DPS) | yes | Totemic Leather Boots (252442, -0.30 DPS, sim-verified) [crafted]; Crusader's Boots (250507, -0.39 DPS) [crafted]; Gilded Slippers (254001, -0.42 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.96 DPS) | yes | Band of the Fist (17694, -0.32 DPS) [quest]; Ring of Power Regulation (273030, -0.36 DPS) [dungeon]; Tiger Band (6749, -0.47 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 16.1 attack_power points (0.78 DPS) | yes | Band of the Fist (17694, -0.15 DPS) [quest]; Ring of Power Regulation (273030, -0.18 DPS) [dungeon]; Tiger Band (6749, -0.30 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (87.2 DPS) | yes | Felbough Arm (273037, -0.59 DPS) [world]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -20.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Crusader's Chain Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hawkeye's Cloak; chest: Crusader's Silvered Chain Shirt; wrist: Spidertank Oilrag; hands: Serpent Gloves; waist: Justicar's Belt; legs: Crusader's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Snake Hoop; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 473, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (undead, 01000000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 127.7. Weights run: 2.4s. Verify run: 2.7s. 658 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.040 ± 0.004, agility=1.525 ± 0.030, crit=1.475 ± 0.036 per rating point (14 rating = 1%, 20.647 per %), hit=1.985 ± 0.088 per rating point (10 rating = 1%, 19.849 per %), melee_haste=11.125 ± 0.653, intellect=1.748 ± 0.018, spell_power=2.065 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Chain Helm (250502) | Blacksmithing [crafted] | 51.3 attack_power points (2.67 DPS) | yes | Papal Fez (9431, -0.27 DPS) [dungeon]; Icemetal Barbute (10763, -0.28 DPS) [dungeon]; Raging Berserker's Helm (7719, -0.52 DPS, sim-verified) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 28.4 attack_power points (1.48 DPS) | yes | Kaleidoscope Chain (13084, -0.37 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.50 DPS, sim-verified) [world]; Triune Amulet (7722, -0.62 DPS) [dungeon] |
| shoulder | Shining Mithril Pauldrons (250541) | Blacksmithing [crafted] | 43.1 attack_power points (2.25 DPS) | yes | Green Silken Shoulders (7057, -0.38 DPS) [crafted]; Bloodmage Mantle (7684, -0.46 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.53 DPS, sim-verified) [dungeon] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 21.1 attack_power points (1.10 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Icy Cloak (4327, -0.35 DPS) [crafted] |
| chest | Shining Mithril Breastplate (250540) | Blacksmithing [crafted] | 57.5 attack_power points (2.99 DPS) | yes | Robe of the Magi (1716, -0.08 DPS) [world_drop]; Dreamweave Vest (10021, -0.24 DPS) [crafted]; Icemail Jerkin (1981, -1.54 DPS, sim-verified) [world_drop] |
| wrist | Hard Gold Bracers (250535) | Blacksmithing [crafted] | 24.6 attack_power points (1.28 DPS) | yes | Guardian Leather Bracers (4260, -0.09 DPS) [crafted]; Radiant Silver Bracers (4545, -0.12 DPS) [quest]; Tough Scorpid Bracers (8205, -0.18 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 49.5 attack_power points (2.58 DPS) | yes | Skirmisher's Leather Gloves (252528, -0.44 DPS) [crafted]; Dreamweave Gloves (10019, -0.45 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.49 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 40.7 attack_power points (2.12 DPS) | yes | Skirmisher's Leather Belt (252461, -0.05 DPS) [crafted]; Justicar's Belt (250560, -0.19 DPS) [crafted]; Prowler's Leather Belt (252459, -0.32 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 49.9 attack_power points (2.60 DPS) | yes | Ferine Leggings (6690, -0.15 DPS) [dungeon]; Totemic Leather Leggings (252458, -0.24 DPS) [crafted]; Hard Gold Leggings (250536, -0.25 DPS) [crafted] |
| feet | Skirmisher's Mail Boots (252564) | Leatherworking [crafted] | 49.8 attack_power points (2.59 DPS) | yes | Earthen Silk Slippers (254013, -0.01 DPS) [crafted]; Justicar's Boots (250550, -0.24 DPS) [crafted]; Shining Mithril Boots (250542, -0.35 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 29.4 attack_power points (1.53 DPS) | yes | Thunderbrow Ring (13097, -0.44 DPS) [world_drop]; Mark of Kern (2262, -0.49 DPS) [dungeon]; Assault Band (13095, -0.49 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 21.1 attack_power points (1.10 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Mark of Kern (2262, -0.06 DPS) [dungeon]; Assault Band (13095, -0.06 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (127.7 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -18.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** neck: Glowing Eye of Mordresh; shoulder: Shining Mithril Pauldrons; back: Long Silken Cloak; chest: Shining Mithril Breastplate; wrist: Hard Gold Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Skirmisher's Mail Boots; finger1: Philanthropist's Ring

No-known-source sample (15 of 658, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler

### Band 50 (undead, 05024000000000000-00000000000000000-05233301001301350)

Set DPS (verified): 174.2. Weights run: 2.6s. Verify run: 2.9s. 868 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.485 ± 0.029, crit=1.769 ± 0.038 per rating point (14 rating = 1%, 24.766 per %), hit=2.227 ± 0.098 per rating point (10 rating = 1%, 22.269 per %), melee_haste=11.213 ± 0.984, intellect=4.041 ± 0.070, spell_power=2.485 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 128.0 attack_power points (6.78 DPS) | yes | Soothsayer's Headdress (17740, -0.54 DPS) [dungeon]; Eye of Theradras (17715, -1.40 DPS) [dungeon]; Charged Scorpid Helm (252581, -1.78 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 56.6 attack_power points (3.00 DPS) | yes | Glowing Eye of Mordresh (10769, -0.36 DPS) [dungeon]; Mindburst Medallion (11196, -0.92 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.08 DPS) [world] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 118.1 attack_power points (6.25 DPS) | yes | Rotgrip Mantle (17732, -0.69 DPS) [dungeon]; Rockshard Pauldrons (9411, -1.54 DPS) [dungeon]; Earthslag Shoulders (11632, -1.58 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 66.2 attack_power points (3.51 DPS) | yes | Spritecaster Cape (11623, -0.38 DPS) [dungeon]; Runecloth Cloak (13860, -0.61 DPS) [crafted]; Big Voodoo Cloak (8216, -0.92 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 128.0 attack_power points (6.78 DPS) | yes | Spiderfang Carapace (11633, -0.66 DPS) [dungeon]; Runecloth Robe (13858, -1.69 DPS) [crafted]; Feathered Breastplate (8349, -2.14 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 65.3 attack_power points (3.46 DPS) | yes | Skirmisher's Mail Bracers (252572, -0.28 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.64 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.64 DPS) [crafted] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | 94.7 attack_power points (5.01 DPS) | yes | Skirmisher's Mail Gauntlets (252586, -0.76 DPS, sim-verified) [crafted]; Skycaller's Leather Gauntlets (252550, -0.98 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.98 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 91.7 attack_power points (4.86 DPS) | yes | Skycaller's Leather Waistguard (252476, +0.00 DPS, sim-verified) [crafted]; Skycaller's Mail Belt (252589, -0.18 DPS) [crafted]; Deathmage Sash (10771, -0.72 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 83.3 attack_power points (4.41 DPS) | yes | Imperial Red Pants (8251, +0.00 DPS, sim-verified) [world_drop]; Kilt of the Atal'ai Prophet (10807, -0.16 DPS) [dungeon]; Big Voodoo Pants (8202, -0.30 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 90.0 attack_power points (4.77 DPS) | yes | Skycaller's Leather Boots (252471, -0.44 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.44 DPS) [crafted]; Mender's Leather Boots (252472, -1.23 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 45.1 attack_power points (2.39 DPS) | yes | Snake Hoop (6750, -0.53 DPS) [quest]; Band of the Unicorn (7553, -0.68 DPS) [world_drop]; Voodoo Band (1996, -0.89 DPS) [world] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.3 attack_power points (2.24 DPS) | yes | Snake Hoop (6750, -0.38 DPS) [quest]; Band of the Unicorn (7553, -0.53 DPS) [world_drop]; Voodoo Band (1996, -0.74 DPS) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (174.2 DPS) | yes | Uther's Strength (11302, -2.26 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (174.2 DPS) | yes | Uther's Strength (11302, -0.27 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (174.2 DPS) | yes | Blight (7959, +0.00 DPS) [crafted]; Warmonger (13052, -3.34 DPS) [world_drop]; Stoneraven (13059, -4.80 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Stonerender Gauntlets; waist: Dawnspire Cord; feet: Greaves of Withering Despair; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 868, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60 (undead, 05025000200000000-50200000000000000-05233301001301350)

Set DPS (verified): 249.3. Weights run: 2.6s. Verify run: 6.4s. 1811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.005 ± 0.048, crit=2.828 ± 0.069 per rating point (14 rating = 1%, 39.592 per %), hit=4.042 ± 0.174 per rating point (10 rating = 1%, 40.418 per %), melee_haste=20.373 ± 1.984, intellect=6.329 ± 0.108, spell_power=2.510 ± 0.004

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 322.2 attack_power points (17.44 DPS) | yes | Magister's Crown (16686, -4.39 DPS, sim-verified) [dungeon]; Blue Dragonscale Helm (252604, -5.69 DPS) [crafted]; Stalwart Helm (250599, -6.47 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 93.4 attack_power points (5.06 DPS) | yes | Glowing Eye of Mordresh (10769, -1.36 DPS) [dungeon]; Mindburst Medallion (11196, -2.19 DPS) [quest]; Horizon Choker (13085, -2.30 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 164.5 attack_power points (8.91 DPS) | yes | Magister's Mantle (16689, -0.28 DPS) [dungeon]; Shroud of the Nathrezim (18720, -0.33 DPS) [dungeon]; Ironfeather Shoulders (15067, -2.97 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 87.1 attack_power points (4.71 DPS) | yes | Spritecaster Cape (11623, -0.76 DPS) [dungeon]; Runecloth Cloak (13860, -0.85 DPS, sim-verified) [crafted]; Imperial Red Cloak (8248, -0.95 DPS) [world_drop] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 207.3 attack_power points (11.22 DPS) | yes | Vest of Elements (16666, -1.25 DPS) [dungeon]; Mooncloth Robe (18486, -1.44 DPS) [crafted]; Blue Dragonscale Breastplate (15048, -2.20 DPS, sim-verified) [crafted] |
| wrist | Magiskull Cuffs (13107) | World drop [world_drop] | 107.6 attack_power points (5.82 DPS) | yes | Tranquil Wristguards (279256, -0.49 DPS) [crafted]; Spitfire Bracers (20481, -0.68 DPS) [crafted]; Runecloth Cuffs (254123, -0.90 DPS) [crafted] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 130.7 attack_power points (7.07 DPS) | yes | Spitfire Gauntlets (20480, +0.00 DPS, sim-verified) [crafted]; Gauntlets of Elements (16672, -0.87 DPS) [dungeon]; Stonerender Gauntlets (17007, -1.07 DPS) [world_drop] |
| waist | Lightforge Belt (16723) | Stratholme: Rockwing Gargoyle [dungeon] | 154.7 attack_power points (8.37 DPS) | yes | Elunarian Belt (14465, -0.81 DPS) [world_drop]; Devout Belt (16696, -0.98 DPS) [dungeon]; Magister's Belt (16685, -2.99 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-verified (249.3 DPS) | yes | Luminary Kilt (11823, -3.63 DPS) [dungeon]; Champion's Legplates (250598, -4.02 DPS) [crafted]; Red Dragonscale Leggings (252603, -4.71 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 131.3 attack_power points (7.11 DPS) | yes | Omnicast Boots (11822, -0.28 DPS) [dungeon]; Mooncloth Boots (15802, -1.23 DPS) [crafted]; Skycaller's Leather Boots (252471, -1.30 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 89.5 attack_power points (4.84 DPS) | yes | Band of the Hierophant (13096, -1.42 DPS) [world_drop]; Chemist's Ring (15702, -1.42 DPS) [quest]; Frigid Ring (18679, -1.42 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 75.9 attack_power points (4.11 DPS) | yes | Band of the Hierophant (13096, -0.69 DPS) [world_drop]; Chemist's Ring (15702, -0.69 DPS) [quest]; Frigid Ring (18679, -0.69 DPS) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (249.3 DPS) | yes | Uther's Strength (11302, -2.99 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (249.3 DPS) | yes | Uther's Strength (11302, -1.15 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (249.3 DPS) | yes | Blade of Hanna (2801, +0.00 DPS, sim-verified) [world_drop]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-verified (249.3 DPS) | yes | Libram of Fervor (23203, -1.27 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Magister's Robes; wrist: Magiskull Cuffs; hands: Mourning Handguards; waist: Lightforge Belt; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Glowing Crystal Ring; main_hand: Legionite Glaive

No-known-source sample (15 of 1811, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

### Band 60, raid preset (undead, 05005000200000000-50300000000000000-05233301001301351)

Set DPS (verified): 660.0. Weights run: 2.7s. Verify run: 6.7s. 1811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.006, agility=1.852 ± 0.037, crit=2.406 ± 0.050 per rating point (14 rating = 1%, 33.683 per %), hit=5.216 ± 0.321 per rating point (10 rating = 1%, 52.157 per %), melee_haste=23.914 ± 3.643, intellect=2.596 ± 0.183, spell_power=1.387 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 215.2 attack_power points (32.60 DPS) | yes | Stalwart Helm (250599, -3.79 DPS, sim-verified) [crafted]; Black Dragonscale Helm (252605, -6.67 DPS) [crafted]; Mask of the Unforgiven (13404, -11.70 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 42.6 attack_power points (6.45 DPS) | yes | Horizon Choker (13085, -0.95 DPS) [world_drop]; Skibi's Pendant (13089, -0.97 DPS) [world_drop]; Woven Ivy Necklace (19159, -1.73 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 103.9 attack_power points (15.74 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS, sim-verified) [crafted]; Shroud of the Nathrezim (18720, -2.88 DPS) [dungeon]; Ironfeather Shoulders (15067, -4.73 DPS) [crafted] |
| back | Cloak of Flames (3475) | World drop [world_drop] | 52.2 attack_power points (7.90 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deep Woodlands Cloak (19121, -1.84 DPS) [quest]; Spritecaster Cape (11623, -2.60 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Savage Gladiator Chain (11726, -2.25 DPS) [dungeon]; Dawn Armor (252483, -6.59 DPS, sim-verified) [crafted]; Icemail Jerkin (1981, -7.61 DPS) [world_drop] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 64.7 attack_power points (9.79 DPS) | yes | Spitfire Bracers (20481, -2.05 DPS) [crafted]; Vambraces of the Sadist (13400, -2.13 DPS) [dungeon]; Lordly Armguards (13135, -2.51 DPS) [world_drop] |
| hands | Stonerender Gauntlets (17007) | World drop [world_drop] | sim-verified (+7.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Storm Gauntlets (12632, -0.29 DPS) [crafted]; Timbermaw Brawlers (19049, -0.35 DPS) [crafted]; Gauntlets of Elements (16672, -7.92 DPS, sim-verified) [dungeon] |
| waist | Lightforge Belt (16723) | Stratholme: Rockwing Gargoyle [dungeon] | 98.9 attack_power points (14.98 DPS) | yes | Brigam Girdle (13142, -1.95 DPS) [dungeon]; Might of the Timbermaw (19044, -3.36 DPS) [crafted]; Verek's Leash (22242, -4.63 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Legplates (250598, -0.48 DPS) [crafted]; Pristine Scorpid Leggings (252606, -1.53 DPS) [crafted]; Red Dragonscale Leggings (252603, -3.38 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 93.4 attack_power points (14.14 DPS) | yes | Depleted Thorium Sabatons (279263, -4.60 DPS) [crafted]; Omnicast Boots (11822, -5.22 DPS) [dungeon]; Timmy's Galoshes (13402, -5.44 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 72.2 attack_power points (10.93 DPS) | yes | Band of the Penitent (13217, -5.83 DPS) [quest]; Glowing Crystal Ring (18402, -6.21 DPS) [quest]; Philanthropist's Ring (281635, -6.86 DPS) [quest] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 43.1 attack_power points (6.53 DPS) | yes | Band of the Penitent (13217, -1.43 DPS) [quest]; Glowing Crystal Ring (18402, -1.82 DPS) [quest]; Philanthropist's Ring (281635, -2.47 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+7.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -4.78 DPS) [crafted]; Uther's Strength (11302, -10.63 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Legionite Glaive (250619, -29.70 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Archlight Talisman; shoulder: Lightforge Spaulders; back: Cloak of Flames; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; waist: Lightforge Belt; legs: Ironfeather Leggings; finger1: Blackstone Ring; finger2: Maiden's Circle; trinket2: Molten Heart of the Mountain; main_hand: Arcanite Champion

No-known-source sample (15 of 1811, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders

