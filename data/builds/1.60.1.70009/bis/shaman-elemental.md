# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 25.9. Weights run: 1.6s. Verify run: 1.2s. 229 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.723 ± 0.018, crit=0.059 ± 0.003 per rating point (14 rating = 1%, 0.822 per %), hit=0.153 ± 0.002 per rating point (10 rating = 1%, 1.532 per %), spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 11.0 spell_power points (1.05 DPS) | yes | Totemic Leather Hood (252448, -0.25 DPS, sim-verified) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.38 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.48 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | sim-verified (25.9 DPS) | yes | Scholarly Pendant (277203, -0.26 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 spell_power points (1.10 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.63 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.72 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.40 DPS) | yes | Feyscale Cloak (6632, -0.11 DPS) [dungeon]; Caretaker's Cape (20428, -0.11 DPS) [rep]; Heavy Woolen Cloak (4311, -0.17 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 11.6 spell_power points (1.11 DPS) | yes | Acolyte's Chain Shirt (250491, -0.29 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.29 DPS) [crafted]; Crusader's Chain Shirt (250492, -0.66 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.3 spell_power points (0.42 DPS) | yes | Mindthrust Bracers (1974, -0.07 DPS) [dungeon]; Ratchet Wristwraps (274742, -0.14 DPS) [vendor]; Owl Bracers (4796, -0.26 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.9 spell_power points (0.76 DPS) | yes | Acolyte's Gloves (250511, -0.12 DPS) [crafted]; Windfelt Gloves (5630, -0.17 DPS) [quest]; Serpent Gloves (5970, -0.28 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.9 spell_power points (0.66 DPS) | yes | Novice Arcanist's Sash (253885, -0.07 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.45 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.8 spell_power points (1.42 DPS) | yes | Stormrider's Leather Pants (252502, -0.24 DPS, sim-verified) [crafted]; Acolyte's Chain Leggings (250496, -0.26 DPS) [crafted]; Dreamer's Leggings (270016, -0.27 DPS) [quest] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (0.95 DPS) | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Acolyte's Boots (250506, -0.22 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.22 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.4 spell_power points (0.62 DPS) | yes | Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Loop of Sacrifice (281673, -0.27 DPS) [quest]; Sludge-Stained Band (286535, -0.33 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.48 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Sludge-Stained Band (286535, -0.19 DPS) [world]; Lavishly Jeweled Ring (1156, -0.32 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | Westfall: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.77 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.07 DPS) [quest]; Channeler's Staff (4437, -0.21 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Stormrider's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 229, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 30 (dwarf, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 49.8. Weights run: 1.7s. Verify run: 1.2s. 393 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.811 ± 0.026, crit=0.124 ± 0.008 per rating point (14 rating = 1%, 1.742 per %), hit=0.230 ± 0.003 per rating point (10 rating = 1%, 2.297 per %), spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 14.1 spell_power points (1.41 DPS) | yes | Crusader's Chain Helm (250502, -0.07 DPS, sim-verified) [crafted]; Enduring Cap (3020, -0.11 DPS) [world_drop]; Totemic Leather Helm (252456, -0.21 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.9 spell_power points (1.18 DPS) | yes | Darkspear Warding Pendant (272075, -0.84 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.3 spell_power points (1.63 DPS) | yes | Fairywing Mantle (9536, -0.30 DPS) [quest]; Magician's Mantle (12998, -0.40 DPS) [world_drop]; Death Speaker Mantle (6685, -0.50 DPS, sim-verified) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.5 spell_power points (0.65 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.02 DPS) [quest]; Hillman's Cloak (3719, -0.15 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 19.5 spell_power points (1.95 DPS) | yes | Guardian Armor (4256, -0.17 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.36 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.47 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.90 DPS) | yes | Nightsky Wristbands (6407, -0.41 DPS) [world_drop]; Technician's Bracers (270042, -0.41 DPS) [quest]; Glowing Magical Bracelets (13106, -1.38 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 12.9 spell_power points (1.29 DPS) | yes | Truefaith Gloves (7049, -0.55 DPS) [crafted]; Acolyte's Gloves (250511, -0.59 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.68 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.9 spell_power points (1.48 DPS) | yes | Prefect's Belt (250559, -0.22 DPS) [crafted]; Justicar's Belt (250560, -0.29 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.43 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 18.9 spell_power points (1.89 DPS) | yes | Abomination Skin Leggings (23173, -0.32 DPS, sim-verified) [dungeon]; Stormrider's Leather Pants (252502, -0.40 DPS) [crafted]; Guardian Pants (5962, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.7 spell_power points (1.27 DPS) | yes | Spidersilk Boots (4320, -0.24 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.26 DPS) [crafted]; Acidic Walkers (9454, -1.11 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.70 DPS) | yes | Black Widow Band (6199, -0.13 DPS) [world]; Snake Hoop (6750, -0.13 DPS) [quest]; Minor Channeling Ring (1449, -1.59 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (49.8 DPS) | yes | Black Widow Band (6199, -0.03 DPS) [world]; Snake Hoop (6750, -0.03 DPS) [quest]; Minor Channeling Ring (1449, -0.94 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Talisman of Arathor (21119, -2.19 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 0.0 spell_power points (0.00 DPS) | yes | Scorn's Focal Dagger (23168, -1.40 DPS) [dungeon]; Glimmering Staff (249392, -1.41 DPS) [crafted]; Manual Crowd Pummeler (9449, -3.43 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 393, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 63.1. Weights run: 1.8s. Verify run: 1.4s. 543 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.227 ± 0.045, crit=0.214 ± 0.014 per rating point (14 rating = 1%, 2.997 per %), hit=0.380 ± 0.006 per rating point (10 rating = 1%, 3.799 per %), spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 26.2 spell_power points (2.33 DPS) | yes | Augural Shroud (2620, -0.26 DPS) [world]; Spellpower Goggles Xtreme (10502, -0.46 DPS) [crafted]; Corpseshroud (10574, -3.21 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.4 spell_power points (1.28 DPS) | yes | Necklace of Calisea (1714, -0.51 DPS) [world_drop]; Triune Amulet (7722, -0.51 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.49 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.9 spell_power points (2.04 DPS) | yes | Green Silken Shoulders (7057, -0.20 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Sheepshear Mantle (13115, -0.30 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Long Silken Cloak (4326, -0.12 DPS) [crafted]; Guardian Cloak (5965, -0.12 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -2.19 DPS, sim-verified) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.03 DPS) [crafted]; Big Voodoo Robe (8200, -0.26 DPS) [crafted]; Robe of the Magi (1716, -0.98 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 13.4 spell_power points (1.19 DPS) | yes | Windchaser Cuffs (14429, -0.21 DPS) [world_drop]; Mistscape Bracers (4045, -0.32 DPS) [world_drop]; Turtle Scale Bracers (8198, -0.49 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.14 DPS) | yes | Dreamweave Gloves (10019, -0.10 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.17 DPS) [crafted]; Red Mageweave Gloves (10018, -0.78 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.4 spell_power points (2.26 DPS) | yes | Highlander's Lizardhide Girdle (20104, -0.62 DPS) [rep]; Highlander's Mail Girdle (20119, -0.62 DPS) [vendor]; Highlander's Cloth Girdle (20098, -0.81 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.7 spell_power points (2.56 DPS) | yes | Crimson Silk Pantaloons (7062, -0.41 DPS, sim-verified) [crafted]; Kodohide Legguards (285338, -0.47 DPS) [world]; Abomination Skin Leggings (23173, -0.88 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.14 DPS) | yes | Skycaller's Leather Shoes (252532, -0.05 DPS, sim-verified) [crafted]; Skycaller's Mail Boots (252563, -0.21 DPS) [crafted]; Mender's Leather Shoes (252533, -0.66 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.4 spell_power points (1.54 DPS) | yes | Ogremind Ring (1993, -0.78 DPS) [world_drop]; Voodoo Band (1996, -0.78 DPS) [world_drop]; Mindbender Loop (5009, -0.78 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.80 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Mindbender Loop (5009, -0.04 DPS) [world_drop]; Voodoo Band (1996, -0.44 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 0.0 spell_power points (0.00 DPS) | yes | Spellforce Rod (1664, -0.46 DPS) [world_drop]; Mograine's Might (7723, -0.49 DPS) [dungeon]; Gut Ripper (2164, -2.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life

No-known-source sample (15 of 543, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (dwarf, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 77.9. Weights run: 1.9s. Verify run: 1.5s. 716 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.216 ± 0.055, crit=0.275 ± 0.018 per rating point (14 rating = 1%, 3.848 per %), hit=0.505 ± 0.007 per rating point (10 rating = 1%, 5.046 per %), spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -0.71 DPS) [crafted]; Chief Architect's Monocle (11839, -0.74 DPS) [dungeon]; Red Mageweave Headband (10033, -0.76 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.24 DPS) [dungeon]; Mindburst Medallion (11196, -0.33 DPS) [quest]; Arcane Crystal Pendant (20037, -2.57 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 39.3 spell_power points (3.46 DPS) | yes | Lead Surveyor's Mantle (11842, -0.64 DPS) [dungeon]; Kentic Amice (11624, -0.84 DPS) [dungeon]; Rotgrip Mantle (17732, -1.08 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.3 spell_power points (1.88 DPS) | yes | Runecloth Cloak (13860, -0.23 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.38 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -2.89 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 43.3 spell_power points (3.82 DPS) | yes | Feathered Breastplate (8349, -1.07 DPS) [crafted]; Robes of Insight (940, -1.14 DPS) [world_drop]; Runecloth Robe (13858, -1.64 DPS, sim-verified) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 22.2 spell_power points (1.95 DPS) | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Skycaller's Mail Bracers (252571, -0.32 DPS) [crafted]; Aristocratic Cuffs (12546, -0.35 DPS) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 33.8 spell_power points (2.97 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.60 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.60 DPS) [crafted]; Raider Handguards (272102, -0.77 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 30.6 spell_power points (2.70 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -0.13 DPS) [dungeon]; Satyrmane Sash (17755, -0.39 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Big Voodoo Pants (8202, -0.13 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -0.33 DPS) [dungeon]; Spellshock Leggings (9484, -3.41 DPS, sim-verified) [dungeon] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 28.4 spell_power points (2.50 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS, sim-verified) [crafted]; Greaves of Withering Despair (22240, -0.01 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.39 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 18.2 spell_power points (1.61 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -0.32 DPS) [dungeon]; Band of the Unicorn (7553, -0.46 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindseye Circle (10634, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Cyclopean Band (11824, -2.06 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Smoking Heart of the Mountain (11811, -0.70 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.64 DPS) [quest]; Mechanic's Pipehammer (9604, -0.90 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 716, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (dwarf, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 165.1. Weights run: 1.9s. Verify run: 1.3s. 1598 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.219 ± 0.086, crit=0.399 ± 0.027 per rating point (14 rating = 1%, 5.588 per %), hit=0.656 ± 0.011 per rating point (10 rating = 1%, 6.555 per %), spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Crown (240123) | Leonid Barthalomew the Revered [vendor] | 92.4 spell_power points (8.01 DPS) | yes | Blue Dragonscale Helm (252604, -3.37 DPS) [crafted]; Soulcrusher Headpiece (240096, -3.48 DPS, sim-verified) [vendor]; Black Dragonscale Helm (252605, -3.51 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.9 spell_power points (2.67 DPS) | yes | Beads of Ogre Mojo (22149, -0.44 DPS, sim-verified) [quest]; Archlight Talisman (15856, -0.58 DPS) [quest]; Arcane Crystal Pendant (20037, -0.65 DPS) [quest] |
| shoulder | Soulcrusher Mantle (240125) | Leonid Barthalomew the Revered [vendor] | 71.6 spell_power points (6.21 DPS) | yes | Darkspear Shoulderguards (272958, -1.97 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.31 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -3.96 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.3 spell_power points (2.80 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -0.64 DPS) [world_drop]; Spritecaster Cape (11623, -0.95 DPS) [dungeon] |
| chest | Soulcrusher Embrace (240109) | Leonid Barthalomew the Revered [vendor] | sim-verified (165.1 DPS) | yes | Soulcrusher Tunic (240092, -1.91 DPS) [vendor]; Soulcrusher Chestguard (240101, -2.78 DPS) [vendor]; Tunic of Undead Slaying (23089, -16.57 DPS, sim-verified) [world] |
| wrist | Soulcrusher Bindings (240127) | Leonid Barthalomew the Revered [vendor] | sim-verified (165.1 DPS) | yes | Soulcrusher Wristguards (240100, -1.43 DPS) [vendor]; Soulcrusher Bracers (240108, -1.91 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -9.62 DPS, sim-verified) [world] |
| hands | Soulcrusher Mitts (240122) | Leonid Barthalomew the Revered [vendor] | 67.4 spell_power points (5.84 DPS) | yes | Raider Handguards (272101, -2.24 DPS) [vendor]; Raider Handwraps (272097, -2.30 DPS) [vendor]; Soulcrusher Handguards (240095, -3.26 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.4 spell_power points (5.06 DPS) | yes | Soulcrusher Cord (240126, +0.00 DPS, sim-verified) [vendor]; Soulcrusher Girdle (240099, -0.28 DPS) [vendor]; Soulcrusher Waistguard (240107, -1.10 DPS) [vendor] |
| legs | Soulcrusher Kilt (240124) | Leonid Barthalomew the Revered [vendor] | 88.5 spell_power points (7.67 DPS) | yes | Leggings of Elemental Fury (23665, -1.98 DPS) [world_drop]; Ironfeather Leggings (252486, -1.99 DPS) [crafted]; Soulcrusher Legguards (240097, -2.85 DPS, sim-verified) [vendor] |
| feet | Soulcrusher Greaves (240110) | Leonid Barthalomew the Revered [vendor] | 64.8 spell_power points (5.62 DPS) | yes | Bloodvine Boots (19684, -1.71 DPS) [crafted]; Slippers of The Five Thunders (227007, -2.11 DPS) [vendor]; Soulcrusher Boots (240093, -3.43 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (165.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.28 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.37 DPS) [vendor]; Naglering (11669, -8.21 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (165.1 DPS) | yes | Ritssyn's Ring of Chaos (21836, -0.42 DPS) [world_drop]; Cauterizing Band (19140, -0.50 DPS) [world_drop]; Naglering (11669, -5.99 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (165.1 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (165.1 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, -0.17 DPS) [dungeon]; Weakness Analyzer (272438, -1.08 DPS, sim-verified) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (165.1 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -11.64 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (23199) (or Totem of Thunder (228176), Tidal Totem (272431), Totem of the Storm (272432), Burning Totem (272433), Totem of Urgency (279249), Totem of Ancestral Protection (249443), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Totem of Thunder (228176, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Soulcrusher Crown; neck: Amulet of the Dawn; shoulder: Soulcrusher Mantle; back: Arcanoweave Cloak; chest: Soulcrusher Embrace; wrist: Soulcrusher Bindings; hands: Soulcrusher Mitts; waist: Knowledge of the Timbermaw; legs: Soulcrusher Kilt; feet: Soulcrusher Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Crackling Staff; ranged: Totem of the Storm

No-known-source sample (15 of 1598, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (orc, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 25.8. Weights run: 1.6s. Verify run: 1.1s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.723 ± 0.018, crit=0.059 ± 0.003 per rating point (14 rating = 1%, 0.822 per %), hit=0.153 ± 0.002 per rating point (10 rating = 1%, 1.532 per %), spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 11.0 spell_power points (1.05 DPS) | yes | Totemic Leather Hood (252448, -0.25 DPS, sim-verified) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.38 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.48 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 spell_power points (1.10 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.72 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.77 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Battle Healer's Cloak (20427, -0.10 DPS) [rep]; Pearl-clasped Cloak (5542, -0.52 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 11.6 spell_power points (1.11 DPS) | yes | Acolyte's Chain Shirt (250491, -0.29 DPS) [crafted]; Wisdom's Leather Armor (252493, -0.29 DPS) [crafted]; Crusader's Chain Shirt (250492, -0.82 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.3 spell_power points (0.42 DPS) | yes | Mindthrust Bracers (1974, -0.07 DPS) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest]; Owl Bracers (4796, -0.16 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.9 spell_power points (0.76 DPS) | yes | Acolyte's Gloves (250511, -0.12 DPS) [crafted]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Serpent Gloves (5970, -0.44 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) (or Stormrider's Leather Belt (252432)) | Tailoring [crafted] | 6.9 spell_power points (0.66 DPS) | yes | Novice Arcanist's Sash (253885, -0.07 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.46 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Acolyte's Chain Leggings (250496, -0.22 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.29 DPS) [crafted]; Abomination Skin Leggings (23173, -0.35 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (0.95 DPS) | yes | Stormrider's Leather Boots (252443, -0.20 DPS, sim-verified) [crafted]; Acolyte's Boots (250506, -0.22 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.22 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.48 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Sludge-Stained Band (286535, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 4.3 spell_power points (0.42 DPS) | yes | Sludge-Stained Band (286535, -0.13 DPS) [world]; Loop of Sacrifice (281673, -0.18 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | Westfall: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.77 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.07 DPS) [quest]; Channeler's Staff (4437, -0.21 DPS) [world]; Twisted Chanter's Staff (890, -0.31 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders

### Band 30 (orc, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 47.2. Weights run: 1.7s. Verify run: 1.2s. 388 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.811 ± 0.026, crit=0.124 ± 0.008 per rating point (14 rating = 1%, 1.742 per %), hit=0.230 ± 0.003 per rating point (10 rating = 1%, 2.297 per %), spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 14.1 spell_power points (1.41 DPS) | yes | Enduring Cap (3020, -0.11 DPS) [world_drop]; Totemic Leather Helm (252456, -0.21 DPS) [crafted]; Crusader's Chain Helm (250502, -0.50 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.9 spell_power points (1.18 DPS) | yes | Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.25 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.3 spell_power points (1.63 DPS) | yes | Death Speaker Mantle (6685, -0.08 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.30 DPS) [quest]; Magician's Mantle (12998, -0.40 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.5 spell_power points (0.65 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.15 DPS) [crafted]; Windsong Drape (15468, -0.15 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 19.5 spell_power points (1.95 DPS) | yes | Death Speaker Robes (6682, -0.36 DPS) [dungeon]; Guardian Armor (4256, -0.42 DPS, sim-verified) [crafted]; Stormrider's Leather Tunic (252510, -0.47 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.90 DPS) | yes | Nightsky Wristbands (6407, -0.41 DPS) [world_drop]; Technician's Bracers (270042, -0.41 DPS) [quest]; Glowing Magical Bracelets (13106, -1.22 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.6 spell_power points (1.06 DPS) | yes | Jutebraid Gloves (10654, +0.00 DPS, sim-verified) [quest]; Stormrider's Leather Gloves (252498, -0.24 DPS) [crafted]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.9 spell_power points (1.48 DPS) | yes | Prefect's Belt (250559, -0.22 DPS) [crafted]; Justicar's Belt (250560, -0.29 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.56 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 18.9 spell_power points (1.89 DPS) | yes | Stormrider's Leather Pants (252502, -0.40 DPS) [crafted]; Guardian Pants (5962, -0.44 DPS) [crafted]; Abomination Skin Leggings (23173, -0.57 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.7 spell_power points (1.27 DPS) | yes | Spidersilk Boots (4320, -0.24 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.26 DPS) [crafted]; Acidic Walkers (9454, -1.04 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.70 DPS) | yes | Black Widow Band (6199, -0.13 DPS) [world]; Snake Hoop (6750, -0.13 DPS) [quest]; Advisor's Ring (20426, -0.20 DPS) [rep] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.60 DPS) | yes | Black Widow Band (6199, -0.03 DPS) [world]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Snake Hoop (6750, -0.45 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (47.2 DPS) | yes | Defiler's Talisman (21120, -2.13 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (47.2 DPS) | yes | Glimmering Staff (249392, -0.01 DPS) [crafted]; Twisted Chanter's Staff (890, -0.09 DPS) [world_drop]; Manual Crowd Pummeler (9449, -1.64 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.9 spell_power points (1.18 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.66 DPS) [vendor]; Seedcloud Buckler (6630, -0.69 DPS) [dungeon]; Witch's Finger (16887, -1.73 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 40 (orc, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 59.5. Weights run: 1.8s. Verify run: 1.4s. 537 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.227 ± 0.045, crit=0.214 ± 0.014 per rating point (14 rating = 1%, 2.997 per %), hit=0.380 ± 0.006 per rating point (10 rating = 1%, 3.799 per %), spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 26.2 spell_power points (2.33 DPS) | yes | Augural Shroud (2620, -0.26 DPS) [world]; Spellpower Goggles Xtreme (10502, -0.46 DPS) [crafted]; Corpseshroud (10574, -2.60 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.4 spell_power points (1.28 DPS) | yes | Necklace of Calisea (1714, -0.51 DPS) [world_drop]; Triune Amulet (7722, -0.51 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.82 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.9 spell_power points (2.04 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Sheepshear Mantle (13115, -0.30 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Long Silken Cloak (4326, -0.12 DPS) [crafted]; Guardian Cloak (5965, -0.12 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -2.05 DPS, sim-verified) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.03 DPS) [crafted]; Big Voodoo Robe (8200, -0.26 DPS) [crafted]; Robe of the Magi (1716, -0.58 DPS, sim-verified) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 13.8 spell_power points (1.23 DPS) | yes | Guardian Leather Bracers (4260, +0.00 DPS, sim-verified) [crafted]; Turtle Scale Bracers (8198, -0.15 DPS) [crafted]; Windchaser Cuffs (14429, -0.25 DPS) [world_drop] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.14 DPS) | yes | Dreamweave Gloves (10019, -0.10 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.17 DPS) [crafted]; Red Mageweave Gloves (10018, -1.50 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.4 spell_power points (2.26 DPS) | yes | Defiler's Cloth Girdle (20166, -0.38 DPS, sim-verified) [rep]; Highlander's Mail Girdle (20119, -0.62 DPS) [vendor]; Defiler's Lizardhide Girdle (20173, -0.62 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.7 spell_power points (2.56 DPS) | yes | Kodohide Legguards (285338, -0.47 DPS) [world]; Crimson Silk Pantaloons (7062, -0.82 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.88 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.14 DPS) | yes | Skycaller's Mail Boots (252563, -0.21 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.49 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -0.66 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.4 spell_power points (1.54 DPS) | yes | Ogremind Ring (1993, -0.78 DPS) [world_drop]; Voodoo Band (1996, -0.78 DPS) [world_drop]; Mindbender Loop (5009, -0.78 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.80 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Mindbender Loop (5009, -0.04 DPS) [world_drop]; Voodoo Band (1996, -0.79 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Mograine's Might (7723, -0.03 DPS) [dungeon]; Windweaver Staff (7757, -0.14 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Dreamweave Vest; wrist: Radiant Silver Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; main_hand: Spellforce Rod

No-known-source sample (15 of 537, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 78.1. Weights run: 1.9s. Verify run: 1.4s. 697 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.216 ± 0.055, crit=0.275 ± 0.018 per rating point (14 rating = 1%, 3.848 per %), hit=0.505 ± 0.007 per rating point (10 rating = 1%, 5.046 per %), spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -0.71 DPS) [crafted]; Chief Architect's Monocle (11839, -0.74 DPS) [dungeon]; Red Mageweave Headband (10033, -1.11 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.24 DPS) [dungeon]; Mindburst Medallion (11196, -0.33 DPS) [quest]; Arcane Crystal Pendant (20037, -2.41 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 39.3 spell_power points (3.46 DPS) | yes | Rotgrip Mantle (17732, -0.56 DPS, sim-verified) [dungeon]; Lead Surveyor's Mantle (11842, -0.64 DPS) [dungeon]; Kentic Amice (11624, -0.84 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mantle of Lady Falther'ess (23178, -0.12 DPS) [dungeon]; Runecloth Cloak (13860, -0.23 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.94 DPS, sim-verified) [quest] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 43.3 spell_power points (3.82 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.62 DPS, sim-verified) [vendor]; Runecloth Robe (13858, -1.03 DPS) [crafted]; Feathered Breastplate (8349, -1.07 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 22.2 spell_power points (1.95 DPS) | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Skycaller's Mail Bracers (252571, -0.32 DPS) [crafted]; Aristocratic Cuffs (12546, -0.35 DPS) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 33.8 spell_power points (2.97 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.60 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.60 DPS) [crafted]; Raider Handguards (272102, -0.82 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 30.6 spell_power points (2.70 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -0.13 DPS) [dungeon]; Satyrmane Sash (17755, -0.39 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 37.6 spell_power points (3.32 DPS) | yes | Red Mageweave Pants (10009, -0.80 DPS) [crafted]; Big Voodoo Pants (8202, -0.92 DPS) [crafted]; Spellshock Leggings (9484, -5.18 DPS, sim-verified) [dungeon] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 28.4 spell_power points (2.50 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS, sim-verified) [crafted]; Greaves of Withering Despair (22240, -0.01 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.39 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 18.2 spell_power points (1.61 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Mindseye Circle (10634, -0.32 DPS) [dungeon]; Band of the Unicorn (7553, -0.46 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindseye Circle (10634, -0.24 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Cyclopean Band (11824, -1.87 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.64 DPS) [quest]; Radiant Staff (249453, -1.07 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 697, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 156.1. Weights run: 1.9s. Verify run: 1.2s. 1531 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.219 ± 0.086, crit=0.399 ± 0.027 per rating point (14 rating = 1%, 5.588 per %), hit=0.656 ± 0.011 per rating point (10 rating = 1%, 6.555 per %), spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Crown (240123) | Leonid Barthalomew the Revered [vendor] | 92.4 spell_power points (8.01 DPS) | yes | Soulcrusher Headpiece (240096, -2.60 DPS, sim-verified) [vendor]; Coif of The Five Thunders (227002, -3.06 DPS) [quest]; Blue Dragonscale Helm (252604, -3.37 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.9 spell_power points (2.67 DPS) | yes | Archlight Talisman (15856, -0.58 DPS) [quest]; Arcane Crystal Pendant (20037, -0.65 DPS) [quest]; Beads of Ogre Mojo (22149, -0.85 DPS, sim-verified) [quest] |
| shoulder | Soulcrusher Mantle (240125) | Leonid Barthalomew the Revered [vendor] | 71.6 spell_power points (6.21 DPS) | yes | Darkspear Shoulderguards (272958, -1.97 DPS) [vendor]; Warlord's Mail Spaulders (231659, -2.21 DPS) [pvp]; Rugged Mantle of the Timbermaw (227808, -3.74 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | sim-verified (156.1 DPS) | yes | Crystalline Threaded Cape (20697, -0.11 DPS) [world_drop]; Deep Woodlands Cloak (19121, -0.28 DPS) [quest]; Arcanoweave Cloak (272411, -2.04 DPS, sim-verified) [vendor] |
| chest | Soulcrusher Embrace (240109) | Leonid Barthalomew the Revered [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Soulcrusher Tunic (240092, -1.91 DPS) [vendor]; Soulcrusher Chestguard (240101, -2.78 DPS) [vendor]; Tunic of Undead Slaying (23089, -15.26 DPS, sim-verified) [world] |
| wrist | Soulcrusher Bindings (240127) | Leonid Barthalomew the Revered [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Soulcrusher Wristguards (240100, -1.43 DPS) [vendor]; Soulcrusher Bracers (240108, -1.91 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -8.39 DPS, sim-verified) [world] |
| hands | Soulcrusher Mitts (240122) | Leonid Barthalomew the Revered [vendor] | 67.4 spell_power points (5.84 DPS) | yes | General's Mail Gauntlets (231660, -1.95 DPS) [pvp]; Raider Handguards (272101, -2.24 DPS) [vendor]; Soulcrusher Handguards (240095, -2.53 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.4 spell_power points (5.06 DPS) | yes | Soulcrusher Cord (240126, -0.24 DPS, sim-verified) [vendor]; Soulcrusher Girdle (240099, -0.28 DPS) [vendor]; Soulcrusher Waistguard (240107, -1.10 DPS) [vendor] |
| legs | Soulcrusher Kilt (240124) | Leonid Barthalomew the Revered [vendor] | 88.5 spell_power points (7.67 DPS) | yes | Leggings of Elemental Fury (23665, -1.98 DPS) [world_drop]; Ironfeather Leggings (252486, -1.99 DPS) [crafted]; Soulcrusher Legguards (240097, -2.56 DPS, sim-verified) [vendor] |
| feet | Soulcrusher Greaves (240110) | Leonid Barthalomew the Revered [vendor] | 64.8 spell_power points (5.62 DPS) | yes | Bloodvine Boots (19684, -1.71 DPS) [crafted]; General's Mail Boots (16573, -2.09 DPS) [vendor]; Soulcrusher Boots (240093, -3.24 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.28 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.37 DPS) [vendor]; Naglering (11669, -6.46 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 0.0 spell_power points (0.00 DPS) | yes | Ritssyn's Ring of Chaos (21836, -0.42 DPS) [world_drop]; Cauterizing Band (19140, -0.50 DPS) [world_drop]; Naglering (11669, -5.04 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, -1.04 DPS, sim-verified) [vendor] |
| main_hand | Hammer of Divine Might (22333) | Scholomance: Kormok [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS, sim-verified) [world_drop]; High Warlord's Destroyer (23465, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (23199) (or Totem of Thunder (228176), Tidal Totem (272431), Totem of the Storm (272432), Burning Totem (272433), Totem of Urgency (279249), Totem of Ancestral Protection (249443), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Totem of Thunder (228176, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Soulcrusher Crown; neck: Amulet of the Dawn; shoulder: Soulcrusher Mantle; back: Hide of the Wild; chest: Soulcrusher Embrace; wrist: Soulcrusher Bindings; hands: Soulcrusher Mitts; waist: Knowledge of the Timbermaw; legs: Soulcrusher Kilt; feet: Soulcrusher Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Hammer of Divine Might; ranged: Totem of the Storm

No-known-source sample (15 of 1531, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

