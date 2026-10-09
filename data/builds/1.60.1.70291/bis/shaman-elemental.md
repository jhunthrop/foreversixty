# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 25.0. Weights run: 2.4s. Verify run: 1.5s. 226 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.655 ± 0.009, crit=0.073 ± 0.002 per rating point (14 rating = 1%, 1.018 per %), hit=0.203 ± 0.006 per rating point (10 rating = 1%, 2.029 per %), spell_haste=-1.772 ± 0.076, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.06 DPS) [crafted]; Totemic Leather Hood (252448, -0.32 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (0.91 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.58 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Pearl-clasped Cloak (5542, -0.00 DPS) [crafted]; Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.03 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.24 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.05 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.05 DPS) [vendor]; Owl Bracers (4796, -1.41 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.6 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Windfelt Gloves (5630, -0.14 DPS) [quest]; Serpent Gloves (5970, -0.35 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.44 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.19 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Dreamer's Leggings (270016, -0.25 DPS) [quest]; Wisdom's Leather Pants (252503, -0.28 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (0.80 DPS) | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.30 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.3 spell_power points (0.53 DPS) | yes | Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Sludge-Stained Band (286535, -0.28 DPS) [world]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.42 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS, sim-verified) [dungeon]; Sludge-Stained Band (286535, -0.17 DPS) [world]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Twisted Chanter's Staff (890, -0.12 DPS) [world_drop]; Channeler's Staff (4437, -0.23 DPS) [world]; Lesser Staff of the Spire (1300, -0.34 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 226, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 5530311300000000-000000000000000000-0000000000000000)

Set DPS (verified): 49.4. Weights run: 2.8s. Verify run: 1.6s. 395 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.639 ± 0.011, crit=0.089 ± 0.002 per rating point (14 rating = 1%, 1.253 per %), hit=0.254 ± 0.008 per rating point (10 rating = 1%, 2.537 per %), spell_haste=-2.308 ± 0.136, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.29 DPS) | yes | Holy Shroud (2721, -0.15 DPS) [world_drop]; Enduring Cap (3020, -0.23 DPS) [world_drop]; Totemic Leather Helm (252456, -0.70 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.13 DPS) | yes | Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.53 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (1.54 DPS) | yes | Fairywing Mantle (9536, -0.31 DPS) [quest]; Death Speaker Mantle (6685, -0.36 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.42 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Cloak of Rot (4462, -0.05 DPS) [world]; Darkspear Raider's Cloak (272078, -0.05 DPS) [vendor]; Vine Pruner's Cloak (279835, -0.51 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (1.81 DPS) | yes | Guardian Armor (4256, -0.20 DPS) [crafted]; Beguiler Robes (7728, -0.33 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.57 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.94 DPS) | yes | Nightsky Wristbands (6407, -0.54 DPS) [world_drop]; Technician's Bracers (270042, -0.54 DPS) [quest]; Glowing Magical Bracelets (13106, -1.23 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 11.0 spell_power points (1.15 DPS) | yes | Shilly Mitts (9609, -0.42 DPS) [quest]; Gloves of Insight (9698, -0.42 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.99 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.8 spell_power points (1.44 DPS) | yes | Moss Cinch (6911, -0.19 DPS) [dungeon]; Belt of Arugal (6392, -0.30 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.55 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 17.0 spell_power points (1.78 DPS) | yes | Dark Ritual Leggings (270031, -0.32 DPS) [quest]; Stormrider's Leather Pants (252502, -0.33 DPS) [crafted]; Abomination Skin Leggings (23173, -0.78 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.20 DPS) | yes | Spidersilk Boots (4320, -0.20 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.24 DPS) [crafted]; Acidic Walkers (9454, -1.50 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.73 DPS) | yes | Black Widow Band (6199, -0.26 DPS) [world]; Snake Hoop (6750, -0.26 DPS) [quest]; Minor Channeling Ring (1449, -1.70 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Widow Band (6199, -0.16 DPS) [world]; Snake Hoop (6750, -0.16 DPS) [quest]; Minor Channeling Ring (1449, -0.97 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wind Spirit Staff (6689, -0.40 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -0.52 DPS) [dungeon]; Manual Crowd Pummeler (9449, -6.86 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Hardened Root Staff

No-known-source sample (15 of 395, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 5530311300103050-010000000000000000-0000000000000000)

Set DPS (verified): 69.9. Weights run: 3.0s. Verify run: 2.1s. 622 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.915 ± 0.024, crit=0.233 ± 0.007 per rating point (14 rating = 1%, 3.264 per %), hit=0.393 ± 0.022 per rating point (10 rating = 1%, 3.929 per %), spell_haste=-2.566 ± 0.289, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.944 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 21.8 spell_power points (2.23 DPS) | yes | Augural Shroud (2620, -0.17 DPS) [world]; Corpseshroud (10574, -0.45 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.83 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.3 spell_power points (1.46 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.53 DPS) [quest]; Scorn's Icy Choker (23169, -0.79 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -0.81 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.9 spell_power points (1.93 DPS) | yes | Bloodmage Mantle (7684, -0.17 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest]; Green Silken Shoulders (7057, -0.74 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 17.2 spell_power points (1.76 DPS) | yes | Guardian Cloak (5965, -0.68 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.73 DPS) [vendor]; Long Silken Cloak (4326, -1.81 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (69.9 DPS) | yes | Robe of Power (7054, -0.13 DPS) [crafted]; Big Voodoo Robe (8200, -0.45 DPS) [crafted]; Robe of the Magi (1716, -0.91 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.5 spell_power points (1.17 DPS) | yes | Turtle Scale Bracers (8198, -0.09 DPS) [crafted]; Arcane Runed Bracers (4744, -0.25 DPS) [quest]; Spidertank Oilrag (9448, -0.25 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.45 DPS) | yes | Dreamweave Gloves (10019, -0.24 DPS) [crafted]; Red Mageweave Gloves (10018, -0.39 DPS) [crafted]; Prospector Gloves (4980, -1.12 DPS, sim-verified) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.7 spell_power points (2.11 DPS) | yes | Skycaller's Leather Belt (252522, -0.53 DPS) [crafted]; Gilded Cord (254037, -0.55 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.10 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.0 spell_power points (2.55 DPS) | yes | Crimson Silk Pantaloons (7062, -0.52 DPS) [crafted]; Kodohide Legguards (285338, -0.69 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -0.88 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.45 DPS) | yes | Skycaller's Leather Shoes (252532, -0.47 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.47 DPS) [crafted]; Mender's Leather Shoes (252533, -0.98 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.49 DPS) | yes | Ring of Forlorn Spirits (2043, -0.67 DPS) [quest]; Reedknot Ring (9622, -0.77 DPS) [quest]; Minor Channeling Ring (1449, -0.79 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.92 DPS) | yes | Reedknot Ring (9622, -0.20 DPS) [quest]; Minor Channeling Ring (1449, -0.22 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.09 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Hand of Righteousness (7721, -2.07 DPS) [dungeon] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 15.1 spell_power points (1.55 DPS) | yes | Thrash's Trash (276204, -0.12 DPS) [vendor]; Orb of Mystic Insight (249394, -0.27 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.14 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of Lorica

No-known-source sample (15 of 622, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 5530311300103050-030000000000000000-5300000000000000)

Set DPS (verified): 98.2. Weights run: 3.1s. Verify run: 2.5s. 796 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.859 ± 0.029, crit=0.295 ± 0.009 per rating point (14 rating = 1%, 4.128 per %), hit=0.512 ± 0.027 per rating point (10 rating = 1%, 5.121 per %), spell_haste=-2.547 ± 0.340, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.937 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.2 spell_power points (3.48 DPS) | yes | Soothsayer's Headdress (17740, -0.03 DPS) [dungeon]; Dreamweave Circlet (10041, -0.63 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.88 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.9 spell_power points (1.33 DPS) | yes | Horizon Choker (13085, -0.18 DPS) [world_drop]; Mindburst Medallion (11196, -0.26 DPS) [quest]; Scorn's Icy Choker (23169, -1.27 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 32.2 spell_power points (3.09 DPS) | yes | Lead Surveyor's Mantle (11842, -0.50 DPS) [dungeon]; Kentic Amice (11624, -0.67 DPS) [dungeon]; Rotgrip Mantle (17732, -1.15 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.2 spell_power points (1.84 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.23 DPS) [dungeon]; Runecloth Cloak (13860, -0.32 DPS) [crafted]; Big Voodoo Cloak (8216, -0.62 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.2 spell_power points (3.48 DPS) | yes | Feathered Breastplate (8349, -0.83 DPS) [crafted]; Robe of the Magi (1716, -0.87 DPS) [world_drop]; Runecloth Tunic (13857, -0.94 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.6 spell_power points (1.79 DPS) | yes | Skycaller's Leather Bracers (252542, -0.25 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.25 DPS) [crafted]; Nethergeld Cuffs (254061, -0.54 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of the Greatfather (17721, -0.18 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.21 DPS) [crafted]; Raider Handguards (272102, -1.00 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | sim-verified (98.2 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.36 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.6 spell_power points (3.04 DPS) | yes | Big Voodoo Pants (8202, -0.77 DPS) [crafted]; Turtle Scale Leggings (8185, -1.15 DPS) [crafted]; Red Mageweave Pants (10009, -2.28 DPS, sim-verified) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 24.4 spell_power points (2.35 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Earthen Silk Slippers (254013, -0.04 DPS) [crafted]; Greaves of Withering Despair (22240, -0.47 DPS) [dungeon] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.0 spell_power points (1.44 DPS) | yes | Band of the Unicorn (7553, -0.19 DPS) [world_drop]; Brainlash (6440, -0.20 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.29 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.37 DPS) | yes | Band of the Unicorn (7553, -0.12 DPS) [world_drop]; Brainlash (6440, -0.14 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.22 DPS) [rep] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.33 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.07 DPS) [dungeon]; Blade of Eternal Darkness (17780, -10.35 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spellshifter Rod

No-known-source sample (15 of 796, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5530311300103050-030000000000000000-5533020000000000)

Set DPS (verified): 176.3. Weights run: 3.0s. Verify run: 7.6s. 1759 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.914 ± 0.026, crit=0.319 ± 0.009 per rating point (14 rating = 1%, 4.466 per %), hit=0.501 ± 0.029 per rating point (10 rating = 1%, 5.014 per %), spell_haste=not significant (-0.457 ± 0.321), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (176.3 DPS) | yes | Black Dragonscale Helm (252605, -0.26 DPS) [crafted]; Magister's Crown (16686, -0.35 DPS) [dungeon]; Blue Dragonscale Helm (252604, -6.81 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 26.9 spell_power points (3.07 DPS) | yes | Beads of Ogre Mojo (22149, -0.33 DPS) [quest]; Orb of the Darkmoon (19426, -0.56 DPS) [quest]; Chains of the Lich (23125, -0.56 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (176.3 DPS) | yes | Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -3.04 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.3 spell_power points (3.24 DPS) | yes | Crystalline Threaded Cape (20697, -0.53 DPS) [world]; Hide of the Wild (18510, -0.59 DPS) [crafted]; Spritecaster Cape (11623, -1.01 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (176.3 DPS) | yes | Chestplate of Tranquility (18373, +0.00 DPS) [dungeon]; Robe of Everlasting Night (18385, +0.00 DPS) [dungeon]; Vest of Elements (16666, -6.68 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-verified (176.3 DPS) | yes | Modest Armguards (18458, -0.93 DPS) [dungeon]; Sublime Wristguards (18497, -0.93 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.27 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (176.3 DPS) | yes | Hands of Power (13253, -0.14 DPS) [dungeon]; Raider Handguards (272102, -0.47 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.1 spell_power points (5.73 DPS) | yes | Stormseeker's Girdle (272399, -1.39 DPS) [vendor]; Girdle of Insight (18504, -1.61 DPS) [crafted]; Belt of the Archmage (18405, -3.19 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.9 spell_power points (6.74 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.01 DPS) [vendor]; Red Dragonscale Leggings (252603, -1.11 DPS) [crafted]; Sentinel's Silk Leggings (237815, -1.68 DPS, sim-verified) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 35.6 spell_power points (4.07 DPS) | yes | Omnicast Boots (11822, -0.53 DPS) [dungeon]; Waterspout Boots (18322, -0.59 DPS) [dungeon]; Dragonrider Boots (18102, -1.71 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (176.3 DPS) | yes | Songstone of Ironforge (12543, -1.93 DPS) [quest]; Maiden's Circle (13001, -1.93 DPS) [world_drop]; Naglering (11669, -10.37 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (176.3 DPS) | yes | Songstone of Ironforge (12543, -0.88 DPS) [quest]; Maiden's Circle (13001, -0.88 DPS) [world_drop]; Naglering (11669, -7.59 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (176.3 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Second Wind (11819, -1.64 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (176.3 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (176.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.46 DPS) [dungeon]; Hand of Edward the Odd (2243, -17.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (176.3 DPS) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; back: Arcanoweave Cloak; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; ranged: Totem of Thunder

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 5530311300103051-020000000000000000-0533520000000000)

Set DPS (verified): 504.7. Weights run: 1.8s. Verify run: 5.5s. 1759 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.357 ± 0.017, crit=0.327 ± 0.008 per rating point (14 rating = 1%, 4.577 per %), hit=0.649 ± 0.039 per rating point (10 rating = 1%, 6.487 per %), spell_haste=2.374 ± 0.470, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.780 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 37.9 spell_power points (14.89 DPS) | yes | Crimson Felt Hat (18727, -1.99 DPS) [dungeon]; Living Crown (252561, -2.55 DPS) [crafted]; Soothsayer's Headdress (17740, -3.75 DPS) [dungeon] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.64 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.93 DPS) [quest]; Diana's Pearl Necklace (22403, -1.44 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.9 spell_power points (14.51 DPS) | yes | Burial Shawl (18681, -4.41 DPS) [dungeon]; Darkspear Shoulderguards (272958, -4.44 DPS) [vendor]; Mantle of the Timbermaw (19050, -5.02 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.3 spell_power points (9.95 DPS) | yes | Crystalline Threaded Cape (20697, -1.54 DPS) [world]; Amplifying Cloak (18350, -2.88 DPS) [dungeon]; Hide of the Wild (18510, -3.05 DPS) [crafted] |
| chest | Robe of Everlasting Night (18385) | Dire Maul: Immol'thar [dungeon] | sim-verified (504.7 DPS) | yes | Vest of Elements (16666, -0.59 DPS) [dungeon]; Chestplate of Tranquility (18373, -0.59 DPS) [dungeon]; Tunic of Undead Slaying (23089, -19.74 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-verified (504.7 DPS) | yes | Modest Armguards (18458, -3.65 DPS) [dungeon]; Sublime Wristguards (18497, -3.65 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -11.54 DPS, sim-verified) [world] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 28.1 spell_power points (11.05 DPS) | yes | Gloves of the Greatfather (17721, -1.63 DPS) [crafted]; Raider Handguards (272101, -1.82 DPS) [vendor]; Storm Gauntlets (12632, -2.43 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.3 spell_power points (15.45 DPS) | yes | Belt of the Archmage (18405, -4.63 DPS, sim-verified) [crafted]; Barrage Girdle (18721, -5.58 DPS) [dungeon]; Stormseeker's Girdle (272399, -5.80 DPS) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.0 spell_power points (19.25 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.68 DPS) [vendor]; Sentinel's Silk Leggings (237815, -1.68 DPS) [vendor]; Pristine Scorpid Leggings (252606, -4.49 DPS) [crafted] |
| feet | Waterspout Boots (18322) | Dire Maul: Hydrospawn [dungeon] | 27.1 spell_power points (10.66 DPS) | yes | Omnicast Boots (11822, -1.12 DPS) [dungeon]; Earthen Silk Slippers (254013, -1.23 DPS) [crafted]; Slippers of The Five Thunders (227007, -4.38 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (504.7 DPS) | yes | Rune Band of Wizardry (22339, -5.98 DPS) [dungeon]; Maiden's Circle (13001, -6.76 DPS) [world_drop]; Naglering (11669, -20.02 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (504.7 DPS) | yes | Rune Band of Wizardry (22339, -1.35 DPS) [dungeon]; Maiden's Circle (13001, -2.13 DPS) [world_drop]; Naglering (11669, -12.43 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (504.7 DPS) | yes | Weakness Analyzer (272438, -2.75 DPS) [vendor]; Serenity Field (272439, -5.89 DPS) [vendor]; Burst of Knowledge (11832, -6.68 DPS) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (504.7 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -4.46 DPS, sim-verified) [dungeon] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (504.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.81 DPS) [quest]; Hand of Edward the Odd (2243, -37.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (504.7 DPS) | yes | Totem of the Storm (272432, -3.14 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Everlasting Night; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Waterspout Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Kindling Stave; ranged: Totem of Thunder

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.7. Weights run: 2.4s. Verify run: 1.4s. 206 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.655 ± 0.009, crit=0.073 ± 0.002 per rating point (14 rating = 1%, 1.018 per %), hit=0.203 ± 0.006 per rating point (10 rating = 1%, 2.029 per %), spell_haste=-1.772 ± 0.076, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.06 DPS) [crafted]; Totemic Leather Hood (252448, -0.33 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (0.91 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.58 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.81 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.40 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.03 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.25 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Featherbead Bracers (15452), Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.3 spell_power points (0.27 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.02 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.6 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.14 DPS) [crafted]; Serpent Gloves (5970, -0.42 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.46 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.19 DPS) | yes | Stormrider's Leather Pants (252502, -0.19 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (0.80 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.30 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.42 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Sludge-Stained Band (286535, -0.17 DPS) [world]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.9 spell_power points (0.33 DPS) | yes | Loop of Sacrifice (281673, -0.05 DPS) [quest]; Sludge-Stained Band (286535, -0.08 DPS) [world]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -0.12 DPS) [world_drop]; Channeler's Staff (4437, -0.23 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 206, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 5530311300000000-000000000000000000-0000000000000000)

Set DPS (verified): 50.4. Weights run: 2.8s. Verify run: 1.4s. 377 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.639 ± 0.011, crit=0.089 ± 0.002 per rating point (14 rating = 1%, 1.253 per %), hit=0.254 ± 0.008 per rating point (10 rating = 1%, 2.537 per %), spell_haste=-2.308 ± 0.136, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.29 DPS) | yes | Holy Shroud (2721, -0.15 DPS) [world_drop]; Enduring Cap (3020, -0.23 DPS) [world_drop]; Totemic Leather Helm (252456, -0.56 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.13 DPS) | yes | Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.10 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (1.54 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS) [dungeon]; Mantle of Woe (7750, -0.22 DPS) [quest]; Fairywing Mantle (9536, -0.31 DPS) [quest] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 5.1 spell_power points (0.53 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Hillman's Cloak (3719, -0.01 DPS) [crafted]; Windsong Drape (15468, -0.01 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (1.81 DPS) | yes | Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Guardian Armor (4256, -0.20 DPS) [crafted]; Beguiler Robes (7728, -0.33 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.94 DPS) | yes | Nightsky Wristbands (6407, -0.54 DPS) [world_drop]; Technician's Bracers (270042, -0.54 DPS) [quest]; Glowing Magical Bracelets (13106, -0.93 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.3 spell_power points (1.07 DPS) | yes | Jutebraid Gloves (10654, -0.11 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.28 DPS) [crafted]; Serpent Gloves (5970, -0.34 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.8 spell_power points (1.44 DPS) | yes | Defiler's Cloth Girdle (20164, -0.10 DPS) [rep]; Moss Cinch (6911, -0.19 DPS) [dungeon]; Warsong Sash (16975, -0.30 DPS) [quest] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 17.0 spell_power points (1.78 DPS) | yes | Dark Ritual Leggings (270031, -0.32 DPS) [quest]; Stormrider's Leather Pants (252502, -0.33 DPS) [crafted]; Abomination Skin Leggings (23173, -0.46 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.20 DPS) | yes | Spidersilk Boots (4320, -0.20 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.24 DPS) [crafted]; Acidic Walkers (9454, -1.13 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.73 DPS) | yes | Black Widow Band (6199, -0.26 DPS) [world]; Snake Hoop (6750, -0.26 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.63 DPS) | yes | Snake Hoop (6750, -0.16 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Black Widow Band (6199, -0.59 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (50.4 DPS) | yes | Royal Diplomatic Scepter (9457, -0.12 DPS) [dungeon]; Scorn's Focal Dagger (23168, -1.80 DPS) [dungeon]; Manual Crowd Pummeler (9449, -8.77 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff

No-known-source sample (15 of 377, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 5530311300103050-010000000000000000-0000000000000000)

Set DPS (verified): 68.9. Weights run: 3.0s. Verify run: 2.1s. 584 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.915 ± 0.024, crit=0.233 ± 0.007 per rating point (14 rating = 1%, 3.264 per %), hit=0.393 ± 0.022 per rating point (10 rating = 1%, 3.929 per %), spell_haste=-2.566 ± 0.289, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.944 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 21.8 spell_power points (2.23 DPS) | yes | Augural Shroud (2620, -0.17 DPS) [world]; Corpseshroud (10574, -0.45 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.68 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.3 spell_power points (1.46 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.53 DPS) [quest]; Darkspear Warding Pendant (272074, -0.81 DPS) [vendor]; Scorn's Icy Choker (23169, -0.82 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.9 spell_power points (1.93 DPS) | yes | Bloodmage Mantle (7684, -0.17 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest]; Green Silken Shoulders (7057, -0.77 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 17.2 spell_power points (1.76 DPS) | yes | Guardian Cloak (5965, -0.68 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.73 DPS) [vendor]; Long Silken Cloak (4326, -1.37 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (68.9 DPS) | yes | Robe of Power (7054, -0.13 DPS) [crafted]; Zealot's Robe (17043, -0.21 DPS) [quest]; Robe of the Magi (1716, -0.86 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.5 spell_power points (1.17 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.02 DPS) [quest]; Turtle Scale Bracers (8198, -0.09 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.45 DPS) | yes | Red Mageweave Gloves (10018, -0.39 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.49 DPS) [crafted]; Dreamweave Gloves (10019, -0.81 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.7 spell_power points (2.11 DPS) | yes | Skycaller's Leather Belt (252522, -0.53 DPS) [crafted]; Gilded Cord (254037, -0.55 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.28 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.0 spell_power points (2.55 DPS) | yes | Crimson Silk Pantaloons (7062, -0.52 DPS) [crafted]; Abomination Skin Leggings (23173, -0.88 DPS) [dungeon]; Kodohide Legguards (285338, -0.99 DPS, sim-verified) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.45 DPS) | yes | Skycaller's Leather Shoes (252532, -0.47 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.47 DPS) [crafted]; Mender's Leather Shoes (252533, -0.98 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.6 spell_power points (1.49 DPS) | yes | Reedknot Ring (9622, -0.77 DPS) [quest]; Voodoo Band (1996, -0.83 DPS) [world]; Black Widow Band (6199, -0.83 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.92 DPS) | yes | Voodoo Band (1996, -0.26 DPS) [world]; Black Widow Band (6199, -0.26 DPS) [world]; Reedknot Ring (9622, -1.03 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Skullbreaker (17039, -0.30 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (1.43 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Mystic Insight (249394, -0.15 DPS) [crafted]; Prophetic Cane (6803, -0.31 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 5530311300103050-030000000000000000-5300000000000000)

Set DPS (verified): 98.3. Weights run: 3.1s. Verify run: 2.4s. 737 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.859 ± 0.029, crit=0.295 ± 0.009 per rating point (14 rating = 1%, 4.128 per %), hit=0.512 ± 0.027 per rating point (10 rating = 1%, 5.121 per %), spell_haste=-2.547 ± 0.340, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.937 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.2 spell_power points (3.48 DPS) | yes | Soothsayer's Headdress (17740, -0.03 DPS) [dungeon]; Dreamweave Circlet (10041, -0.63 DPS) [crafted]; Blood Guard's Pulsing Helmet (220848, -0.83 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.9 spell_power points (1.33 DPS) | yes | Scorn's Icy Choker (23169, -0.17 DPS) [dungeon]; Horizon Choker (13085, -0.18 DPS) [world_drop]; Mindburst Medallion (11196, -0.26 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 32.2 spell_power points (3.09 DPS) | yes | Lead Surveyor's Mantle (11842, -0.50 DPS) [dungeon]; Kentic Amice (11624, -0.67 DPS) [dungeon]; Rotgrip Mantle (17732, -0.99 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.7 spell_power points (1.90 DPS) | yes | Spritecaster Cape (11623, -0.06 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.29 DPS) [dungeon]; Runecloth Cloak (13860, -0.37 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.2 spell_power points (3.48 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.26 DPS) [vendor]; Feathered Breastplate (8349, -0.83 DPS) [crafted]; Robe of the Magi (1716, -0.87 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.6 spell_power points (1.79 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.25 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of the Greatfather (17721, -0.18 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.21 DPS) [crafted]; Raider Handguards (272102, -1.38 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | sim-verified (98.3 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.36 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 33.4 spell_power points (3.21 DPS) | yes | Spellshock Leggings (9484, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -0.88 DPS) [crafted]; Big Voodoo Pants (8202, -0.95 DPS) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 24.4 spell_power points (2.35 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Earthen Silk Slippers (254013, -0.04 DPS) [crafted]; Greaves of Withering Despair (22240, -0.47 DPS) [dungeon] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.0 spell_power points (1.44 DPS) | yes | Band of the Unicorn (7553, -0.19 DPS) [world_drop]; Brainlash (6440, -0.20 DPS) [dungeon]; Advisor's Ring (19519, -0.29 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.37 DPS) | yes | Band of the Unicorn (7553, -0.12 DPS) [world_drop]; Brainlash (6440, -0.14 DPS) [dungeon]; Advisor's Ring (19519, -0.22 DPS) [rep] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.13 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -0.33 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -1.07 DPS) [dungeon]; Blade of Eternal Darkness (17780, -10.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; feet: Skycaller's Leather Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Rune of the Guard Captain; main_hand: Spellshifter Rod

No-known-source sample (15 of 737, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 5530311300103050-030000000000000000-5533020000000000)

Set DPS (verified): 173.8. Weights run: 3.0s. Verify run: 7.8s. 1679 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.914 ± 0.026, crit=0.319 ± 0.009 per rating point (14 rating = 1%, 4.466 per %), hit=0.501 ± 0.029 per rating point (10 rating = 1%, 5.014 per %), spell_haste=not significant (-0.457 ± 0.321), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Coif of The Five Thunders (227002) | Saving the Best for Last [quest] | 48.3 spell_power points (5.52 DPS) | yes | Warlord's Mail Helm (231663, -0.29 DPS) [pvp]; Living Crown (252561, -0.78 DPS) [crafted]; Blue Dragonscale Helm (252604, -5.47 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 26.9 spell_power points (3.07 DPS) | yes | Beads of Ogre Mojo (22149, -0.33 DPS) [quest]; Orb of the Darkmoon (19426, -0.56 DPS) [quest]; Chains of the Lich (23125, -0.56 DPS) [dungeon] |
| shoulder | Pauldrons of The Five Thunders (227003) | Anthion's Parting Words [quest] | sim-verified (173.8 DPS) | yes | Rugged Mantle of the Timbermaw (227808, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -4.66 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.3 spell_power points (3.24 DPS) | yes | Crystalline Threaded Cape (20697, -0.53 DPS) [world]; Hide of the Wild (18510, -0.59 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.93 DPS) [quest] |
| chest | Vest of The Five Thunders (227004) | Saving the Best for Last [quest] | sim-verified (173.8 DPS) | yes | Vest of Elements (16666, +0.00 DPS) [dungeon]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -4.38 DPS, sim-verified) [crafted] |
| wrist | Bindings of The Five Thunders (227001) | An Earnest Proposition [quest] | sim-verified (173.8 DPS) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep] |
| hands | Gauntlets of The Five Thunders (227006) | Just Compensation [quest] | sim-verified (173.8 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Raider Handguards (272101, -4.90 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 50.1 spell_power points (5.73 DPS) | yes | Stormseeker's Girdle (272399, -1.39 DPS) [vendor]; Girdle of Insight (18504, -1.61 DPS) [crafted]; Belt of the Archmage (18405, -2.80 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.9 spell_power points (6.74 DPS) | yes | General's Mail Leggings (231664, -0.52 DPS) [pvp]; Sentinel's Lizardhide Pants (237817, -1.01 DPS) [vendor]; Sentinel's Silk Leggings (237815, -1.52 DPS, sim-verified) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 35.6 spell_power points (4.07 DPS) | yes | General's Mail Sabatons (231661, -0.01 DPS) [vendor]; Omnicast Boots (11822, -0.53 DPS) [dungeon]; Dragonrider Boots (18102, -6.32 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (173.8 DPS) | yes | Eye of Orgrimmar (12545, -1.93 DPS) [quest]; Maiden's Circle (13001, -1.93 DPS) [world_drop]; Naglering (11669, -9.79 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (173.8 DPS) | yes | Eye of Orgrimmar (12545, -0.88 DPS) [quest]; Maiden's Circle (13001, -0.88 DPS) [world_drop]; Naglering (11669, -7.98 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (173.8 DPS) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (173.8 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (173.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.46 DPS) [dungeon]; Hand of Edward the Odd (2243, -17.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (173.8 DPS) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Coif of The Five Thunders; neck: Amulet of the Dawn; shoulder: Pauldrons of The Five Thunders; back: Arcanoweave Cloak; chest: Vest of The Five Thunders; wrist: Bindings of The Five Thunders; hands: Gauntlets of The Five Thunders; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; ranged: Totem of Thunder

No-known-source sample (15 of 1679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (orc, 5530311300103051-020000000000000000-0533520000000000)

Set DPS (verified): 501.2. Weights run: 1.8s. Verify run: 5.6s. 1679 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.357 ± 0.017, crit=0.327 ± 0.008 per rating point (14 rating = 1%, 4.577 per %), hit=0.649 ± 0.039 per rating point (10 rating = 1%, 6.487 per %), spell_haste=2.374 ± 0.470, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.780 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 37.9 spell_power points (14.89 DPS) | yes | Warlord's Mail Helm (231663, -1.26 DPS) [pvp]; Crimson Felt Hat (18727, -1.99 DPS) [dungeon]; Coif of The Five Thunders (227002, -2.02 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.64 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.93 DPS) [quest]; Diana's Pearl Necklace (22403, -1.44 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.9 spell_power points (14.51 DPS) | yes | Warlord's Mail Spaulders (231659, -2.22 DPS) [pvp]; Pauldrons of The Five Thunders (227003, -3.21 DPS, sim-verified) [quest]; Mantle of the Timbermaw (19050, -4.21 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.3 spell_power points (9.95 DPS) | yes | Crystalline Threaded Cape (20697, -1.54 DPS) [world]; Amplifying Cloak (18350, -2.88 DPS) [dungeon]; Hide of the Wild (18510, -3.05 DPS) [crafted] |
| chest | Robe of Everlasting Night (18385) | Dire Maul: Immol'thar [dungeon] | sim-verified (501.2 DPS) | yes | Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Legionnaire's Mail Breastplate (227165, -0.25 DPS) [vendor]; Tunic of Undead Slaying (23089, -19.79 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-verified (501.2 DPS) | yes | Modest Armguards (18458, -3.65 DPS) [dungeon]; Sublime Wristguards (18497, -3.65 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -12.04 DPS, sim-verified) [world] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 28.1 spell_power points (11.05 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Gloves of the Greatfather (17721, -1.63 DPS) [crafted]; Storm Gauntlets (12632, -3.61 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.3 spell_power points (15.45 DPS) | yes | Belt of the Archmage (18405, -4.39 DPS, sim-verified) [crafted]; Barrage Girdle (18721, -5.58 DPS) [dungeon]; Stormseeker's Girdle (272399, -5.80 DPS) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.0 spell_power points (19.25 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.68 DPS) [vendor]; Sentinel's Silk Leggings (237815, -1.68 DPS) [vendor]; General's Mail Leggings (231664, -3.09 DPS) [pvp] |
| feet | Waterspout Boots (18322) | Dire Maul: Hydrospawn [dungeon] | 27.1 spell_power points (10.66 DPS) | yes | General's Mail Sabatons (231661, -0.42 DPS) [vendor]; Omnicast Boots (11822, -1.12 DPS) [dungeon]; Slippers of The Five Thunders (227007, -5.26 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (501.2 DPS) | yes | Rune Band of Wizardry (22339, -5.98 DPS) [dungeon]; Maiden's Circle (13001, -6.76 DPS) [world_drop]; Naglering (11669, -19.90 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (501.2 DPS) | yes | Rune Band of Wizardry (22339, -1.35 DPS) [dungeon]; Maiden's Circle (13001, -2.13 DPS) [world_drop]; Naglering (11669, -12.49 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (501.2 DPS) | yes | Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (501.2 DPS) | yes | Weakness Analyzer (272438, -2.75 DPS) [vendor]; Royal Seal of Eldre'Thalas (18471, -3.85 DPS, sim-verified) [quest]; Serenity Field (272439, -5.89 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (501.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.81 DPS) [quest]; Hand of Edward the Odd (2243, -38.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (501.2 DPS) | yes | Totem of the Storm (272432, -3.10 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Everlasting Night; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Waterspout Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Kindling Stave; ranged: Totem of Thunder

No-known-source sample (15 of 1679, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

