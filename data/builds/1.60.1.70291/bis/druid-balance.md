# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 30.2. Weights run: 1.5s. Verify run: 0.9s. 194 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.782 ± 0.005, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.438 per %), hit=0.102 ± 0.005 per rating point (10 rating = 1%, 1.018 per %), spell_haste=-0.851 ± 0.032, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.000, arcane_power=0.967 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Wisdom's Leather Hood (252507, -0.40 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS) [crafted]; Trapper's Leather Hood (252505, -0.84 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 spell_power points (1.60 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -1.07 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -1.09 DPS, sim-verified) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.3 spell_power points (0.58 DPS) | yes | Sanguine Cape (14376, -0.16 DPS) [world_drop]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.53 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.39 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.9 spell_power points (0.52 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.10 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.1 spell_power points (1.08 DPS) | yes | Windfelt Gloves (5630, -0.24 DPS) [quest]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.54 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.12 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.72 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.3 spell_power points (2.03 DPS) | yes | Dreamer's Leggings (270016, -0.35 DPS) [quest]; Wisdom's Leather Pants (252503, -0.47 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.65 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.1 spell_power points (1.34 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.29 DPS) [crafted]; Black Whelp Slippers (252424, -0.53 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.6 spell_power points (0.87 DPS) | yes | Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Sludge-Stained Band (286535, -0.47 DPS) [world]; Volcanic Rock Ring (12053, -0.56 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.66 DPS) | yes | Sludge-Stained Band (286535, -0.27 DPS) [world]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.06 DPS) | yes | Channeler's Staff (4437, -0.23 DPS) [world]; Lesser Staff of the Spire (1300, -0.44 DPS) [world]; Twisted Chanter's Staff (890, -1.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 194, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 65.1. Weights run: 1.6s. Verify run: 0.9s. 350 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.548 ± 0.007, crit=0.035 ± 0.002 per rating point (14 rating = 1%, 0.497 per %), hit=0.123 ± 0.005 per rating point (10 rating = 1%, 1.225 per %), spell_haste=0.320 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.09 DPS) [world_drop]; Silk Headband (7050, -0.47 DPS) [crafted]; Totemic Leather Helm (252456, -1.23 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.94 DPS) | yes | Crystal Starfire Medallion (5003, -1.52 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.52 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.07 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.9 spell_power points (2.62 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS) [dungeon]; Fairywing Mantle (9536, -0.56 DPS) [quest]; Magician's Mantle (12998, -0.75 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Hillman's Cloak (3719, -0.04 DPS) [crafted]; Cloak of Rot (4462, -0.15 DPS) [world]; Vine Pruner's Cloak (279835, -2.54 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.1 spell_power points (3.04 DPS) | yes | Guardian Armor (4256, -0.31 DPS) [crafted]; Beguiler Robes (7728, -0.52 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.63 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.69 DPS) | yes | Glowing Magical Bracelets (13106, -0.59 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Technician's Bracers (270042, -1.08 DPS) [quest] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.0 spell_power points (1.89 DPS) | yes | Stormrider's Leather Gloves (252498, -0.53 DPS) [crafted]; Shilly Mitts (9609, -0.57 DPS) [quest]; Gloves of Insight (9698, -0.57 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (2.50 DPS) | yes | Highlander's Cloth Girdle (20099, -0.12 DPS) [rep]; Moss Cinch (6911, -0.24 DPS) [dungeon]; Belt of Arugal (6392, -0.50 DPS) [dungeon] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 16.0 spell_power points (3.02 DPS) | yes | Dark Ritual Leggings (270031, +0.00 DPS, sim-verified) [quest]; Abomination Skin Leggings (23173, -0.50 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.52 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.04 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -1.99 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.32 DPS) | yes | Black Widow Band (6199, -0.60 DPS) [world]; Snake Hoop (6750, -0.60 DPS) [quest]; Minor Channeling Ring (1449, -2.74 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Widow Band (6199, -0.41 DPS) [world]; Snake Hoop (6750, -0.41 DPS) [quest]; Minor Channeling Ring (1449, -1.08 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hardened Root Staff (1317) | What Comes Around... [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wind Spirit Staff (6689, -0.80 DPS) [dungeon]; Royal Diplomatic Scepter (9457, -0.94 DPS) [dungeon]; Manual Crowd Pummeler (9449, -11.13 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Hardened Root Staff

No-known-source sample (15 of 350, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 85.2. Weights run: 1.9s. Verify run: 1.1s. 461 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.398 ± 0.005, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.005 per %), hit=0.170 ± 0.007 per rating point (10 rating = 1%, 1.702 per %), spell_haste=not significant (-0.271 ± 0.073), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.29 DPS) | yes | Augural Shroud (2620, -0.76 DPS, sim-verified) [world]; Big Voodoo Mask (8201, -1.31 DPS) [crafted]; Living Cowl (5608, -1.64 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.2 spell_power points (2.08 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.27 DPS) [quest]; Darkspear Warding Pendant (272074, -1.51 DPS) [vendor]; Scorn's Icy Choker (23169, -1.54 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (2.57 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.6 spell_power points (2.57 DPS) | yes | Guardian Cloak (5965, -0.94 DPS) [crafted]; Icy Cloak (4327, -1.14 DPS) [crafted]; Long Silken Cloak (4326, -2.34 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.12 DPS) [world_drop]; Robe of Power (7054, -0.57 DPS) [crafted]; Robe of the Magi (1716, -1.18 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Leather Bracers (4260, -0.13 DPS) [crafted]; Condor Bracers (15864, -0.41 DPS) [quest]; Arcane Runed Bracers (4744, -1.43 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.91 DPS) | yes | Prospector Gloves (4980, -1.35 DPS) [quest]; Dreamweave Gloves (10019, -1.70 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.84 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.6 spell_power points (3.19 DPS) | yes | Star Belt (4329, -0.53 DPS) [crafted]; Deathmage Sash (10771, -0.54 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.65 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.8 spell_power points (3.84 DPS) | yes | Dark Ritual Leggings (270031, -0.98 DPS) [quest]; Crimson Silk Pantaloons (7062, -1.14 DPS) [crafted]; Kodohide Legguards (285338, -1.30 DPS, sim-verified) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.91 DPS) | yes | Skycaller's Leather Shoes (252532, -1.60 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.70 DPS) [crafted]; Gilded Slippers (254001, -2.91 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (2.45 DPS) | yes | Ring of Forlorn Spirits (2043, -0.82 DPS) [quest]; Reedknot Ring (9622, -1.02 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.22 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.84 DPS) | yes | Reedknot Ring (9622, -0.41 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.61 DPS) [vendor]; Ring of Forlorn Spirits (2043, -1.61 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.16 DPS) [dungeon]; Hand of Righteousness (7721, -3.31 DPS) [dungeon] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.86 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Lorica (11262, -0.82 DPS) [quest]; Orb of Mystic Insight (249394, -0.94 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 461, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 140.7. Weights run: 1.9s. Verify run: 1.5s. 604 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.468 ± 0.009, crit=0.116 ± 0.006 per rating point (14 rating = 1%, 1.626 per %), hit=0.277 ± 0.013 per rating point (10 rating = 1%, 2.768 per %), spell_haste=not significant (0.045 ± 0.138), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.377 ± 0.001, arcane_power=0.623 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.0 spell_power points (5.69 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -0.82 DPS) [crafted]; Red Mageweave Headband (10033, -1.25 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.7 spell_power points (2.04 DPS) | yes | Mindburst Medallion (11196, -0.37 DPS) [quest]; Horizon Choker (13085, -0.79 DPS) [world_drop]; Scorn's Icy Choker (23169, -1.67 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.4 spell_power points (4.62 DPS) | yes | Kentic Amice (11624, -0.81 DPS) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.93 DPS) [vendor]; Rotgrip Mantle (17732, -2.86 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (3.18 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.68 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted]; Big Voodoo Cloak (8216, -1.44 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.4 spell_power points (5.37 DPS) | yes | Feathered Breastplate (8349, -0.89 DPS) [crafted]; Dreamweave Vest (10021, -1.17 DPS) [crafted]; Robe of the Magi (1716, -2.87 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.7 spell_power points (2.78 DPS) | yes | Skycaller's Leather Bracers (252542, -0.27 DPS) [crafted]; Nethergeld Cuffs (254061, -0.83 DPS) [crafted]; Mender's Leather Bracers (252543, -1.02 DPS) [crafted] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 25.1 spell_power points (4.75 DPS) | yes | Bloodfire Talons (12464, -0.90 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -0.92 DPS) [crafted]; Gloves of the Greatfather (17721, -2.62 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.6 spell_power points (4.10 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -0.56 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.09 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.7 spell_power points (5.24 DPS) | yes | Big Voodoo Pants (8202, -1.52 DPS) [crafted]; Red Mageweave Pants (10009, -1.53 DPS) [crafted]; Knight's Crackling Leather Leggings (220864, -2.61 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.55 DPS) | yes | Skycaller's Leather Boots (252471, -0.73 DPS) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -1.05 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.46 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.46 DPS) | yes | Lorekeeper's Ring (19523, -0.19 DPS) [rep]; Philanthropist's Ring (281635, -3.40 DPS, sim-verified) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (140.7 DPS) | yes | Lorekeeper's Ring (19523, -0.05 DPS) [rep]; Philanthropist's Ring (281635, -1.49 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -1.77 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.16 DPS) [quest]; Spire of Hakkar (10844, -0.69 DPS) [world]; Blade of Eternal Darkness (17780, -19.96 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Cyclopean Band; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave

No-known-source sample (15 of 604, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 243.8. Weights run: 1.9s. Verify run: 4.6s. 1450 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.666 ± 0.014, crit=0.145 ± 0.008 per rating point (14 rating = 1%, 2.037 per %), hit=0.350 ± 0.017 per rating point (10 rating = 1%, 3.500 per %), spell_haste=not significant (0.325 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.375 ± 0.001, arcane_power=0.625 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cowl (226773) | Saving the Best for Last [quest] | 40.0 spell_power points (8.06 DPS) | yes | Field Marshal's Dragonhide Helm (231695, -0.34 DPS) [vendor]; Living Crown (252561, -0.61 DPS) [crafted]; Crimson Felt Hat (18727, -0.95 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 23.7 spell_power points (4.76 DPS) | yes | Orb of the Darkmoon (19426, -0.33 DPS) [quest]; Beads of Ogre Mojo (22149, -0.54 DPS) [quest]; Chains of the Lich (23125, -2.63 DPS, sim-verified) [dungeon] |
| shoulder | Feralheart Spaulders (226778) | Anthion's Parting Words [quest] | sim-verified (243.8 DPS) | yes | Burial Shawl (18681, +0.00 DPS) [dungeon]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Rugged Mantle of the Timbermaw (227808, -4.59 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.8 spell_power points (5.00 DPS) | yes | Crystalline Threaded Cape (20697, -0.44 DPS) [world]; Hide of the Wild (18510, -0.84 DPS) [crafted]; Amplifying Cloak (18350, -1.37 DPS) [dungeon] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-verified (243.8 DPS) | yes | Field Marshal's Dragonhide Armor (231696, -0.19 DPS) [vendor]; Chestplate of Tranquility (18373, -0.34 DPS) [dungeon]; Tunic of Undead Slaying (23089, -22.65 DPS, sim-verified) [world] |
| wrist | Feralheart Wraps (226775) | Mokvar [vendor] | sim-verified (243.8 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (243.8 DPS) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Raider Handwraps (272097, -0.34 DPS) [vendor]; Hands of Power (13253, -6.58 DPS, sim-verified) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.2 spell_power points (8.69 DPS) | yes | Elunite Cord (272401, -2.25 DPS) [vendor]; Girdle of Insight (18504, -2.58 DPS) [crafted]; Belt of the Archmage (18405, -4.89 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.0 spell_power points (10.48 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.87 DPS) [vendor]; Sentinel's Silk Leggings (22752, -2.29 DPS) [rep] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 32.7 spell_power points (6.58 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.47 DPS) [pvp]; Dragonrider Boots (18102, -0.81 DPS) [dungeon]; Waterspout Boots (18322, -8.44 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (243.8 DPS) | yes | Songstone of Ironforge (12543, -2.99 DPS) [quest]; Maiden's Circle (13001, -2.99 DPS) [world_drop]; Naglering (11669, -14.90 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (243.8 DPS) | yes | Songstone of Ironforge (12543, -1.34 DPS) [quest]; Maiden's Circle (13001, -1.34 DPS) [world_drop]; Naglering (11669, -11.52 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (243.8 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (243.8 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (243.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.27 DPS) [dungeon]; Hand of Edward the Odd (2243, -24.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Cowl; neck: Amulet of the Dawn; shoulder: Feralheart Spaulders; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Feralheart Wraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Second Wind; main_hand: Spellshifter Rod

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 4132220115501051-05000000000000000000-5053000000000000)

Set DPS (verified): 531.0. Weights run: 1.1s. Verify run: 4.5s. 1450 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.233 ± 0.011, crit=0.284 ± 0.011 per rating point (14 rating = 1%, 3.969 per %), hit=0.598 ± 0.022 per rating point (10 rating = 1%, 5.982 per %), spell_haste=3.830 ± 0.105, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.194 ± 0.000, arcane_power=0.806 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cowl (226773) | Saving the Best for Last [quest] | 34.2 spell_power points (15.93 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Dragonhide Helm (231695, -1.18 DPS) [vendor]; Living Crown (252561, -2.32 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.26 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.85 DPS) [quest]; Diana's Pearl Necklace (22403, -2.40 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.5 spell_power points (16.07 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, -1.97 DPS) [pvp]; Feralheart Spaulders (226778, -3.12 DPS) [quest]; Argent Shoulders (19059, -4.41 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.8 spell_power points (11.12 DPS) | yes | Crystalline Threaded Cape (20697, -1.36 DPS) [world]; Amplifying Cloak (18350, -2.72 DPS) [dungeon]; Hide of the Wild (18510, -3.50 DPS) [crafted] |
| chest | Knight-Captain's Dragonhide Armor (227195) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Robe of Everlasting Night (18385, +0.00 DPS) [dungeon]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Feralheart Vest (226776, -12.57 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -4.45 DPS) [dungeon]; Runecloth Cuffs (254123, -4.91 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -11.48 DPS, sim-verified) [world] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-verified (531.0 DPS) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Feralheart Hands (226777, -0.92 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 36.1 spell_power points (16.83 DPS) | yes | Belt of the Archmage (18405, -5.11 DPS, sim-verified) [crafted]; Elunite Cord (272401, -6.77 DPS) [vendor]; Ban'thok Sash (11662, -7.26 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 46.2 spell_power points (21.52 DPS) | yes | Sentinel's Silk Leggings (237815, -1.82 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.82 DPS) [vendor]; Skyshroud Leggings (13170, -4.80 DPS) [dungeon] |
| feet | Knight-Lieutenant's Dragonhide Boots (227194) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feralheart Galoshes (226774, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Waterspout Boots (18322, -5.44 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.57 DPS) [dungeon]; Maiden's Circle (13001, -7.67 DPS) [world_drop]; Naglering (11669, -19.84 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.20 DPS) [dungeon]; Maiden's Circle (13001, -2.30 DPS) [world_drop]; Naglering (11669, -11.84 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+24.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.26 DPS) [vendor]; Serenity Field (272439, -5.50 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -7.93 DPS) [dungeon] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.37 DPS) [world]; Hand of Edward the Odd (2243, -40.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Knight-Captain's Dragonhide Armor; wrist: Dryad's Wrist Bindings; hands: Gloves of the Greatfather; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Knight-Lieutenant's Dragonhide Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 29.0. Weights run: 1.5s. Verify run: 0.9s. 184 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.782 ± 0.005, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.438 per %), hit=0.102 ± 0.005 per rating point (10 rating = 1%, 1.018 per %), spell_haste=-0.851 ± 0.032, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.000, arcane_power=0.967 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Trapper's Leather Hood (252505, -0.38 DPS, sim-verified) [crafted]; Wisdom's Leather Hood (252507, -0.40 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 spell_power points (1.60 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.04 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.07 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.3 spell_power points (0.58 DPS) | yes | Heavy Woolen Cloak (4311, -0.05 DPS) [crafted]; Sanguine Cape (14376, -0.16 DPS) [world_drop]; Black Whelp Cloak (7283, -0.18 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.39 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Featherbead Bracers (15452), Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.9 spell_power points (0.52 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.10 DPS) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.1 spell_power points (1.08 DPS) | yes | Serpent Gloves (5970, -0.15 DPS) [dungeon]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.27 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.12 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.71 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.3 spell_power points (2.03 DPS) | yes | Stormrider's Leather Pants (252502, -0.08 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.47 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.61 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Boots (252444, -0.27 DPS) [crafted]; Spidersilk Boots (4320, -0.31 DPS, sim-verified) [crafted]; Black Whelp Slippers (252424, -0.50 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Sludge-Stained Band (286535, -0.27 DPS) [world]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.7 spell_power points (0.62 DPS) | yes | Sludge-Stained Band (286535, -0.23 DPS) [world]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; Loop of Sacrifice (281673, -0.33 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.06 DPS) | yes | Twisted Chanter's Staff (890, -0.02 DPS) [world_drop]; Channeler's Staff (4437, -0.23 DPS) [world]; Gnarled Necromancer's Staff (251534, -0.83 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 184, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 66.3. Weights run: 1.6s. Verify run: 1.0s. 342 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.548 ± 0.007, crit=0.035 ± 0.002 per rating point (14 rating = 1%, 0.497 per %), hit=0.123 ± 0.005 per rating point (10 rating = 1%, 1.225 per %), spell_haste=0.320 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.26 DPS) | yes | Holy Shroud (2721, -0.19 DPS) [world_drop]; Silk Headband (7050, -0.56 DPS) [crafted]; Enchanter's Cowl (4322, -1.28 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.94 DPS) | yes | Crystal Starfire Medallion (5003, -1.52 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.52 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.79 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.9 spell_power points (2.62 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS) [dungeon]; Mantle of Woe (7750, -0.44 DPS) [quest]; Fairywing Mantle (9536, -0.56 DPS) [quest] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (66.3 DPS) | yes | Cloak of Rot (4462, -0.12 DPS) [world]; Darkspear Raider's Cloak (272078, -0.12 DPS) [vendor]; Hillman's Cloak (3719, -0.66 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.1 spell_power points (3.04 DPS) | yes | Mechbuilder's Overalls (9508, -0.29 DPS) [dungeon]; Guardian Armor (4256, -0.31 DPS) [crafted]; Beguiler Robes (7728, -0.52 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.69 DPS) | yes | Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Technician's Bracers (270042, -1.08 DPS) [quest]; Glowing Magical Bracelets (13106, -2.14 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.1 spell_power points (1.90 DPS) | yes | Jutebraid Gloves (10654, -0.26 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.55 DPS) [crafted]; Serpent Gloves (5970, -0.58 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (2.50 DPS) | yes | Moss Cinch (6911, -0.24 DPS) [dungeon]; Warsong Sash (16975, -0.43 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.94 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 16.0 spell_power points (3.02 DPS) | yes | Abomination Skin Leggings (23173, -0.50 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.52 DPS) [crafted]; Dark Ritual Leggings (270031, -1.19 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.04 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -1.80 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.32 DPS) | yes | Black Widow Band (6199, -0.60 DPS) [world]; Snake Hoop (6750, -0.60 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.70 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.13 DPS) | yes | Snake Hoop (6750, -0.41 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; Black Widow Band (6199, -2.60 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Diplomatic Scepter (9457, -0.14 DPS) [dungeon]; Scorn's Focal Dagger (23168, -3.15 DPS) [dungeon]; Manual Crowd Pummeler (9449, -13.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Wind Spirit Staff

No-known-source sample (15 of 342, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 84.9. Weights run: 1.9s. Verify run: 1.2s. 446 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.398 ± 0.005, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.005 per %), hit=0.170 ± 0.007 per rating point (10 rating = 1%, 1.702 per %), spell_haste=not significant (-0.271 ± 0.073), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.29 DPS) | yes | Augural Shroud (2620, -1.23 DPS) [world]; Big Voodoo Mask (8201, -1.31 DPS) [crafted]; Living Cowl (5608, -1.64 DPS) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.2 spell_power points (2.08 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.27 DPS) [quest]; Darkspear Warding Pendant (272074, -1.51 DPS) [vendor]; Scorn's Icy Choker (23169, -1.55 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (2.57 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.6 spell_power points (2.57 DPS) | yes | Guardian Cloak (5965, -0.94 DPS) [crafted]; Icy Cloak (4327, -1.14 DPS) [crafted]; Long Silken Cloak (4326, -2.35 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (84.9 DPS) | yes | Elemental Raiment (9434, -0.12 DPS) [world_drop]; Zealot's Robe (17043, -0.53 DPS) [quest]; Robe of the Magi (1716, -1.54 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.84 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.37 DPS) [quest]; Guardian Leather Bracers (4260, -0.89 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.91 DPS) | yes | Dreamweave Gloves (10019, -1.39 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.84 DPS) [crafted]; Red Mageweave Gloves (10018, -1.84 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.6 spell_power points (3.19 DPS) | yes | Star Belt (4329, -0.53 DPS) [crafted]; Deathmage Sash (10771, -0.54 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.65 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.8 spell_power points (3.84 DPS) | yes | Dark Ritual Leggings (270031, -0.98 DPS) [quest]; Crimson Silk Pantaloons (7062, -1.14 DPS) [crafted]; Kodohide Legguards (285338, -1.30 DPS, sim-verified) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.91 DPS) | yes | Skycaller's Leather Shoes (252532, -1.65 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.70 DPS) [crafted]; Gilded Slippers (254001, -2.91 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (2.45 DPS) | yes | Reedknot Ring (9622, -1.02 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.22 DPS) [vendor]; Sludge-Stained Band (286535, -1.84 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.84 DPS) | yes | Sea Giant's Toe Ring (274746, -0.61 DPS) [vendor]; Sludge-Stained Band (286535, -1.23 DPS) [world]; Reedknot Ring (9622, -1.84 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.16 DPS) [dungeon]; Skullbreaker (17039, -0.49 DPS) [quest] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.86 DPS) | yes | Thrash's Trash (276204, +0.00 DPS) [vendor]; Orb of Mystic Insight (249394, -0.94 DPS) [crafted]; Omega Orb (7749, -1.02 DPS) [quest] |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 139.7. Weights run: 1.9s. Verify run: 1.4s. 585 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.468 ± 0.009, crit=0.116 ± 0.006 per rating point (14 rating = 1%, 1.626 per %), hit=0.277 ± 0.013 per rating point (10 rating = 1%, 2.768 per %), spell_haste=not significant (0.045 ± 0.138), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.377 ± 0.001, arcane_power=0.623 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.0 spell_power points (5.69 DPS) | yes | Red Mageweave Headband (10033, -0.31 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -0.82 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.7 spell_power points (2.04 DPS) | yes | Scorn's Icy Choker (23169, -0.18 DPS) [dungeon]; Mindburst Medallion (11196, -0.37 DPS) [quest]; Horizon Choker (13085, -0.79 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.4 spell_power points (4.62 DPS) | yes | Kentic Amice (11624, -0.81 DPS) [dungeon]; Blood Guard's Crackling Leather Spaulders (220871, -0.93 DPS) [vendor]; Rotgrip Mantle (17732, -2.57 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (3.18 DPS) | yes | Deep Woodlands Cloak (19121, -0.11 DPS) [quest]; Mantle of Lady Falther'ess (23178, -0.68 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.4 spell_power points (5.37 DPS) | yes | Feathered Breastplate (8349, -0.89 DPS) [crafted]; Dreamweave Vest (10021, -1.17 DPS) [crafted]; Robe of the Magi (1716, -1.94 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.7 spell_power points (2.78 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.27 DPS) [crafted] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 25.1 spell_power points (4.75 DPS) | yes | Bloodfire Talons (12464, -0.90 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -0.92 DPS) [crafted]; Gloves of the Greatfather (17721, -2.08 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.6 spell_power points (4.10 DPS) | yes | Ban'thok Sash (11662, -0.32 DPS) [dungeon]; Satyrmane Sash (17755, -0.56 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.09 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.7 spell_power points (5.24 DPS) | yes | Big Voodoo Pants (8202, -1.52 DPS) [crafted]; Red Mageweave Pants (10009, -1.53 DPS) [crafted]; Stone Guard's Crackling Leather Leggings (220865, -2.82 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.55 DPS) | yes | Skycaller's Leather Boots (252471, -0.73 DPS) [crafted]; First Sergeant's Crackling Leather Boots (220863, -1.05 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.46 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.46 DPS) | yes | Advisor's Ring (19519, -0.19 DPS) [rep]; Philanthropist's Ring (281635, -3.34 DPS, sim-verified) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (139.7 DPS) | yes | Advisor's Ring (19519, -0.05 DPS) [rep]; Philanthropist's Ring (281635, -2.35 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.66 DPS) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.16 DPS) [quest]; Spire of Hakkar (10844, -0.69 DPS) [world]; Blade of Eternal Darkness (17780, -19.10 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Cyclopean Band; trinket1: Uther's Strength; main_hand: Kindling Stave

No-known-source sample (15 of 585, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 240.4. Weights run: 1.9s. Verify run: 3.8s. 1444 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.666 ± 0.014, crit=0.145 ± 0.008 per rating point (14 rating = 1%, 2.037 per %), hit=0.350 ± 0.017 per rating point (10 rating = 1%, 3.500 per %), spell_haste=not significant (0.325 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.375 ± 0.001, arcane_power=0.625 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cowl (226773) | Saving the Best for Last [quest] | 40.0 spell_power points (8.06 DPS) | yes | Warlord's Dragonhide Helm (231678, -0.34 DPS) [vendor]; Living Crown (252561, -0.61 DPS) [crafted]; Crimson Felt Hat (18727, -0.95 DPS) [dungeon] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 23.7 spell_power points (4.76 DPS) | yes | Orb of the Darkmoon (19426, -0.33 DPS) [quest]; Beads of Ogre Mojo (22149, -0.54 DPS) [quest]; Chains of the Lich (23125, -2.47 DPS, sim-verified) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 39.0 spell_power points (7.86 DPS) | yes | Warlord's Dragonhide Spaulders (231681, -0.94 DPS) [vendor]; Burial Shawl (18681, -1.69 DPS) [dungeon]; Feralheart Spaulders (226778, -1.72 DPS) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.8 spell_power points (5.00 DPS) | yes | Crystalline Threaded Cape (20697, -0.44 DPS) [world]; Hide of the Wild (18510, -0.84 DPS) [crafted]; Amplifying Cloak (18350, -1.37 DPS) [dungeon] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Armor (231679, -0.19 DPS) [vendor]; Chestplate of Tranquility (18373, -0.34 DPS) [dungeon]; Tunic of Undead Slaying (23089, -18.15 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -1.75 DPS) [dungeon]; Runecloth Cuffs (254123, -1.95 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -11.37 DPS, sim-verified) [world] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (240.4 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Raider Handwraps (272097, -0.34 DPS) [vendor]; Hands of Power (13253, -7.35 DPS, sim-verified) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.2 spell_power points (8.69 DPS) | yes | Elunite Cord (272401, -2.25 DPS) [vendor]; Girdle of Insight (18504, -2.58 DPS) [crafted]; Belt of the Archmage (18405, -5.18 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.0 spell_power points (10.48 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.87 DPS) [vendor]; Outrider's Silk Leggings (22747, -2.29 DPS) [rep]; Sentinel's Silk Leggings (237815, -3.26 DPS, sim-verified) [vendor] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 32.7 spell_power points (6.58 DPS) | yes | General's Dragonhide Boots (231682, -0.47 DPS) [pvp]; Dragonrider Boots (18102, -0.81 DPS) [dungeon]; Waterspout Boots (18322, -5.15 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -2.99 DPS) [quest]; Maiden's Circle (13001, -2.99 DPS) [world_drop]; Naglering (11669, -14.30 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.34 DPS) [quest]; Maiden's Circle (13001, -1.34 DPS) [world_drop]; Naglering (11669, -9.97 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+13.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -0.27 DPS) [dungeon]; Hand of Edward the Odd (2243, -23.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Cowl; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Burst of Knowledge; main_hand: Spellshifter Rod

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 4132220115501051-05000000000000000000-5053000000000000)

Set DPS (verified): 518.3. Weights run: 1.1s. Verify run: 3.5s. 1444 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.233 ± 0.011, crit=0.284 ± 0.011 per rating point (14 rating = 1%, 3.969 per %), hit=0.598 ± 0.022 per rating point (10 rating = 1%, 5.982 per %), spell_haste=3.830 ± 0.105, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.194 ± 0.000, arcane_power=0.806 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cowl (226773) | Saving the Best for Last [quest] | 34.2 spell_power points (15.93 DPS) | yes | Crimson Felt Hat (18727, -1.07 DPS) [dungeon]; Warlord's Dragonhide Helm (231678, -1.18 DPS) [vendor]; Living Crown (252561, -2.32 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.26 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.85 DPS) [quest]; Diana's Pearl Necklace (22403, -2.40 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 34.5 spell_power points (16.07 DPS) | yes | Warlord's Dragonhide Spaulders (231681, -1.97 DPS) [vendor]; Feralheart Spaulders (226778, -3.12 DPS) [quest]; Argent Shoulders (19059, -4.41 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.8 spell_power points (11.12 DPS) | yes | Crystalline Threaded Cape (20697, -1.36 DPS) [world]; Amplifying Cloak (18350, -2.72 DPS) [dungeon]; Hide of the Wild (18510, -3.50 DPS) [crafted] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Robe of Everlasting Night (18385, -0.08 DPS) [dungeon]; Tunic of Undead Slaying (23089, -14.83 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -4.45 DPS) [dungeon]; Runecloth Cuffs (254123, -4.91 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -11.77 DPS, sim-verified) [world] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Feralheart Hands (226777, -0.92 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 36.1 spell_power points (16.83 DPS) | yes | Belt of the Archmage (18405, -5.46 DPS, sim-verified) [crafted]; Elunite Cord (272401, -6.77 DPS) [vendor]; Ban'thok Sash (11662, -7.26 DPS) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 46.2 spell_power points (21.52 DPS) | yes | Sentinel's Silk Leggings (237815, -1.82 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.82 DPS) [vendor]; Skyshroud Leggings (13170, -4.80 DPS) [dungeon] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | sim-verified (518.3 DPS) | yes | Waterspout Boots (18322, +0.00 DPS) [dungeon]; General's Dragonhide Boots (231682, -0.68 DPS) [pvp]; Earthen Silk Slippers (254013, -0.80 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -6.57 DPS) [dungeon]; Maiden's Circle (13001, -7.67 DPS) [world_drop]; Naglering (11669, -21.10 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.20 DPS) [dungeon]; Maiden's Circle (13001, -2.30 DPS) [world_drop]; Naglering (11669, -12.12 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+24.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.26 DPS) [vendor]; Serenity Field (272439, -5.46 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -7.93 DPS) [dungeon] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.37 DPS) [world]; The Lobotomizer (19324, -41.75 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; hands: Gloves of the Greatfather; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

