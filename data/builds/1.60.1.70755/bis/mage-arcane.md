# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 141005000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 47.8. Weights run: 3.1s. Verify run: 3.6s. 178 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.941 ± 0.013, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.122 per %), hit=0.327 ± 0.015 per rating point (10 rating = 1%, 3.273 per %), spell_haste=not significant (0.121 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.906 ± 0.004, spirit=not significant (0.196 ± 0.052), mp5=0.796 ± 0.101

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 10.0 spell_power points (1.94 DPS) | yes | Shadow Goggles (4373, -0.80 DPS) [crafted]; Pristine Circlet (253949, -0.92 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.79 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.4 spell_power points (0.08 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.5 spell_power points (2.62 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.65 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.85 DPS) [crafted]; Slime-encrusted Pads (6461, -2.16 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Pearl-clasped Cloak (5542, -0.13 DPS, sim-verified) [crafted]; Sanguine Cape (14376, -0.20 DPS) [world_drop]; Caretaker's Cape (20428, -0.27 DPS) [rep] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 13.5 spell_power points (2.62 DPS) | yes | Filigreed Pristine Gown (253901, -0.73 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.92 DPS) [crafted]; Mystic's Wrap (14369, -1.34 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.1 spell_power points (1.38 DPS) | yes | Bright Bracers (3647, -0.37 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.94 DPS) [quest]; Mystic's Bracelets (14366, -1.02 DPS) [world_drop] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.2 spell_power points (1.79 DPS) | yes | Serpent Gloves (5970, -0.42 DPS) [dungeon]; Pristine Gloves (253913, -0.46 DPS) [crafted]; Gnoll Casting Gloves (892, -0.62 DPS) [world] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 10.1 spell_power points (1.97 DPS) | yes | Keller's Girdle (2911, -0.50 DPS) [world_drop]; Pristine Sash (253925, -0.55 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.64 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 16.4 spell_power points (3.20 DPS) | yes | Filigreed Pristine Leggings (253937, -0.92 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.68 DPS) [world_drop]; Silk-threaded Trousers (1929, -1.83 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.8 spell_power points (2.10 DPS) | yes | Pearly Boots (253897, -0.39 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.96 DPS) [crafted]; Red Woolen Boots (4313, -1.13 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 spell_power points (1.34 DPS) | yes | Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.79 DPS) [world_drop]; Band of Purification (12996, -1.11 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 6.6 spell_power points (1.28 DPS) | yes | Lavishly Jeweled Ring (1156, -0.65 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop]; Band of Purification (12996, -1.05 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 30.9 spell_power points (6.01 DPS) | yes | Staff of the Blessed Seer (2271, -0.92 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.34 DPS) [dungeon]; Buzzer Blade (2169, -1.73 DPS) [dungeon] |
| off_hand | Mathystra Channeling Orb (281291) | Baron Marinous [quest] | 4.7 spell_power points (0.92 DPS) | yes | Tear of Grief (5611, -0.11 DPS) [quest]; Dwarven Tome (279898, -0.29 DPS) [quest]; Seer's Fine Stein (7608, -0.33 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Magic Wand (11288, -0.90 DPS) [crafted]; Sizzle Stick (8071, -0.90 DPS, sim-verified) [quest]; Torchlight Wand (5240, -0.99 DPS) [quest] |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Mathystra Channeling Orb; ranged: Skycaller

No-known-source sample (15 of 178, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff; 20440 Protector's Sword

### Band 30 (gnome, 141005213100300000-00000000000000000-000000000000000000)

Set DPS (verified): 91.1. Weights run: 3.2s. Verify run: 4.5s. 303 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.555 ± 0.010, crit=0.067 ± 0.002 per rating point (14 rating = 1%, 0.936 per %), hit=0.243 ± 0.012 per rating point (10 rating = 1%, 2.433 per %), spell_haste=not significant (0.430 ± 0.125), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.593 ± 0.003, spirit=not significant (0.212 ± 0.081), mp5=not significant (0.199 ± 0.115)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 12.3 spell_power points (2.88 DPS) | yes | Enchanter's Cowl (4322, -0.17 DPS) [crafted]; Silk Headband (7050, -0.77 DPS) [crafted]; Filigreed Pearly Circlet (253983, -0.93 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.3 spell_power points (0.78 DPS) | yes | Kaleidoscope Chain (13084, -0.06 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.11 DPS) [world_drop]; Seal of the Expedition (271920, -0.13 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.29 DPS) | yes | Death Speaker Mantle (6685, -0.44 DPS) [dungeon]; Magician's Mantle (12998, -0.94 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Expeditionary's Cape (271923) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.02 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.02 DPS) [quest]; Heavy Woolen Cloak (4311, -0.05 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 17.4 spell_power points (4.08 DPS) | yes | Mechbuilder's Overalls (9508, -0.63 DPS) [dungeon]; Green Silk Armor (7065, -0.69 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -1.00 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.6 spell_power points (2.25 DPS) | yes | Glowing Magical Bracelets (13106, -1.17 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -1.32 DPS) [world_drop]; Mindthrust Bracers (1974, -1.46 DPS) [dungeon] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 9.2 spell_power points (2.17 DPS) | yes | Serpent Gloves (5970, -0.52 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.55 DPS) [quest]; Truefaith Gloves (7049, -0.60 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.7 spell_power points (2.98 DPS) | yes | Belt of Arugal (6392, -0.22 DPS) [dungeon]; Crimson Silk Belt (7055, -0.65 DPS) [crafted]; Invoker's Cord (215366, -0.68 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 13.9 spell_power points (3.27 DPS) | yes | Pearly Leggings (253995, -0.16 DPS) [crafted]; Pristine Leggings (253987, -0.46 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.75 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 12.4 spell_power points (2.92 DPS) | yes | Gilded Slippers (254001, -0.17 DPS) [crafted]; Acidic Walkers (9454, -0.51 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (2.06 DPS) | yes | Ring of Power Regulation (273030, -0.51 DPS) [dungeon]; Minor Channeling Ring (1449, -0.63 DPS) [quest]; Electrocutioner Lagnut (9447, -1.01 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.4 spell_power points (1.74 DPS) | yes | Ring of Power Regulation (273030, -0.19 DPS) [dungeon]; Minor Channeling Ring (1449, -0.30 DPS) [quest]; Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (91.1 DPS) | yes | Soulstaff (249393, -1.06 DPS) [crafted]; Tendonscraper Dagger (273053, -1.14 DPS) [world]; Glimmering Staff (249392, -2.83 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (2.43 DPS) | yes | Arcane Infused Rod (279838, -0.67 DPS, sim-verified) [quest]; Orb of Souls (249395, -1.19 DPS) [crafted]; Eye of Paleth (2943, -1.49 DPS) [quest] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-verified (91.1 DPS) | yes | Thunderwood (13062, -0.11 DPS) [world_drop]; Greater Mystic Wand (217287, -0.29 DPS) [crafted]; Wand of Eventide (5214, -0.38 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Expeditionary's Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 303, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 173.8. Weights run: 3.4s. Verify run: 4.0s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.737 ± 0.022, crit=0.173 ± 0.004 per rating point (14 rating = 1%, 2.415 per %), hit=0.375 ± 0.022 per rating point (10 rating = 1%, 3.754 per %), spell_haste=not significant (0.107 ± 0.297), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.617 ± 0.003, spirit=not significant (0.223 ± 0.250), mp5=not significant (0.228 ± 0.264)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 22.5 spell_power points (5.71 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.39 DPS) [crafted]; Augural Shroud (2620, -1.06 DPS) [world]; Whitemane's Chapeau (7720, -1.39 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.9 spell_power points (3.27 DPS) | yes | Triune Amulet (7722, -0.49 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.57 DPS) [dungeon]; Seal of the Expedition (271919, -1.96 DPS) [rep] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.6 spell_power points (4.20 DPS) | yes | Green Silken Shoulders (7057, -0.12 DPS) [crafted]; Bloodmage Mantle (7684, -0.24 DPS) [dungeon]; Berylline Pads (4197, -0.56 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (173.8 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Expeditionary's Cape (271922, -0.17 DPS) [rep]; Icy Cloak (4327, -0.34 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 27.8 spell_power points (7.04 DPS) | yes | Robe of the Magi (1716, -0.05 DPS) [world_drop]; Robe of Power (7054, -0.79 DPS) [crafted]; Elemental Raiment (9434, -1.71 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.7 spell_power points (2.46 DPS) | yes | Arcane Runed Bracers (4744, -0.17 DPS) [quest]; Condor Bracers (15864, -0.68 DPS) [quest]; Windchaser Cuffs (14429, -0.77 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.5 spell_power points (5.71 DPS) | yes | Black Mageweave Gloves (10003, -1.90 DPS) [crafted]; Golden Handwraps (254029, -2.21 DPS) [crafted]; Red Mageweave Gloves (10018, -2.61 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.1 spell_power points (4.58 DPS) | yes | Highlander's Cloth Girdle (20098, -0.28 DPS) [rep]; Golden Cord (254045, -0.56 DPS) [crafted]; Gilded Cord (254037, -0.72 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.8 spell_power points (5.79 DPS) | yes | Crimson Silk Pantaloons (7062, -1.79 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.84 DPS) [dungeon]; Pearly Leggings (253995, -2.01 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.6 spell_power points (6.48 DPS) | yes | Gilded Slippers (254001, -3.17 DPS) [crafted]; Acidic Walkers (9454, -3.49 DPS) [dungeon]; Golden Slippers (254009, -4.45 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (3.47 DPS) | yes | Lorekeeper's Ring (19524, -1.01 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.44 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.7 spell_power points (2.46 DPS) | yes | Lorekeeper's Ring (19524, -0.00 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.43 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (173.8 DPS) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (173.8 DPS) | yes | Celestial Stave (9517, -7.78 DPS) [quest]; Hypnotic Blade (7714, -8.12 DPS) [dungeon]; Illusionary Rod (7713, -11.61 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-verified (173.8 DPS) | yes | Starfaller (13063, -0.39 DPS) [world_drop]; Twisted Nether Wand (249144, -0.52 DPS) [crafted]; Fizzle's Zippy Lighter (6729, -0.54 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 141005213100301531-40000000000000000-051000000000000000)

Set DPS (verified): 252.1. Weights run: 3.5s. Verify run: 5.2s. 491 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.704 ± 0.028, crit=0.358 ± 0.008 per rating point (14 rating = 1%, 5.006 per %), hit=0.863 ± 0.044 per rating point (10 rating = 1%, 8.627 per %), spell_haste=not significant (0.542 ± 0.603), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.646 ± 0.002, spirit=not significant (0.495 ± 0.248), mp5=not significant (0.524 ± 0.262)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 34.0 spell_power points (12.09 DPS) | yes | Red Mageweave Headband (10033, -1.70 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.32 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.48 DPS) [crafted] |
| neck | Mindburst Medallion (11196) | Tremors of the Earth [quest] | 12.7 spell_power points (4.52 DPS) | yes | Horizon Choker (13085, -0.31 DPS) [world_drop]; Triune Amulet (7722, -0.44 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.76 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.7 spell_power points (9.14 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -2.25 DPS) [vendor]; Nethergeld Shoulders (254049, -2.87 DPS) [crafted]; Nethershine Shoulders (254057, -3.31 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (252.1 DPS) | yes | Expeditionary's Cape (271921, -3.15 DPS) [rep]; Big Voodoo Cloak (8216, -3.51 DPS) [crafted]; Runecloth Cloak (13860, -3.85 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.1 spell_power points (11.77 DPS) | yes | Robe of the Magi (1716, -1.56 DPS) [world_drop]; Dreamweave Vest (10021, -2.31 DPS, sim-verified) [crafted]; Robe of Power (7054, -2.37 DPS) [crafted] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 15.8 spell_power points (5.62 DPS) | yes | Nethergeld Cuffs (254061, -0.50 DPS) [crafted]; Bloodband Bracers (11469, -1.06 DPS) [quest]; Spidertank Oilrag (9448, -1.86 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.3 spell_power points (8.64 DPS) | yes | Golden Gloves (254103, -0.36 DPS) [crafted]; Runecloth Gloves (13863, -0.53 DPS) [crafted]; Gilded Gloves (254095, -1.17 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 21.0 spell_power points (7.49 DPS) | yes | Dawnspire Cord (12466, -0.59 DPS) [dungeon]; Golden Waistcord (254089, -0.87 DPS) [crafted]; Ghostweave Cord (254073, -1.90 DPS, sim-verified) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 25.8 spell_power points (9.20 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.45 DPS) [dungeon]; Red Mageweave Pants (10009, -1.21 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.28 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.5 spell_power points (9.77 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.60 DPS) [vendor]; Golden Sandals (254115, -2.92 DPS) [crafted]; Gilded Sandals (254107, -3.60 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 14.1 spell_power points (5.02 DPS) | yes | Band of the Unicorn (7553, -0.39 DPS) [world_drop]; Eye of Adaegus (5266, -1.40 DPS) [world_drop]; Snake Hoop (6750, -1.63 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.81 DPS) | yes | Eye of Adaegus (5266, -1.19 DPS) [world_drop]; Snake Hoop (6750, -1.42 DPS) [quest]; Band of the Unicorn (7553, -1.91 DPS, sim-verified) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (252.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.69 DPS, sim-verified) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (252.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest]; Uther's Strength (11302, -1.88 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (252.1 DPS) | yes | Glowing Brightwood Staff (812, -5.89 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -6.59 DPS) [dungeon]; Spellshifter Rod (9527, -7.75 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-verified (252.1 DPS) | yes | Nether Force Wand (11263, -0.70 DPS) [quest]; Noxious Shooter (17745, -0.73 DPS) [dungeon]; Charged Lightning Rod (11860, -0.74 DPS) [quest] |

**New at 50:** head: Dreamweave Circlet; neck: Mindburst Medallion; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethershine Cuffs; waist: Satyrmane Sash; legs: Senior Designer's Pantaloons; finger1: Lorekeeper's Ring; finger2: Philanthropist's Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Ankh of Life; ranged: Lesser Eternal Wand

No-known-source sample (15 of 491, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 141005213100301531-40000000000000000-053500030000000000)

Set DPS (verified): 393.0. Weights run: 3.6s. Verify run: 17.7s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.820 ± 0.041, crit=0.551 ± 0.013 per rating point (14 rating = 1%, 7.711 per %), hit=1.180 ± 0.065 per rating point (10 rating = 1%, 11.803 per %), spell_haste=not significant (0.883 ± 0.937), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.603 ± 0.002, spirit=not significant (0.599 ± 0.405), mp5=not significant (1.066 ± 0.433)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreadmist Mask (16698) | Scholomance: Darkmaster Gandling [dungeon] | 40.1 spell_power points (15.48 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -1.42 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.2 spell_power points (7.81 DPS) | yes | Mindburst Medallion (11196, -2.43 DPS) [quest]; Horizon Choker (13085, -2.44 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -2.57 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 35.4 spell_power points (13.67 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Theramore Shoulderpads (271953, -0.47 DPS) [rep]; Lieutenant Commander's Silk Mantle (227102, -0.48 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (393.0 DPS) | yes | Expeditionary's Cape (271907, -2.24 DPS) [rep]; Runecloth Cloak (13860, -2.69 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (393.0 DPS) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -4.60 DPS) [pvp]; Magister's Robes (16688, -11.56 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 32.8 spell_power points (12.66 DPS) | yes | Marshal's Silk Bracers (16438, -4.32 DPS) [pvp]; Runecloth Cuffs (254123, -5.24 DPS) [crafted]; Dreadmist Bracers (16703, -6.08 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 32.6 spell_power points (12.61 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Theramore Handwraps (271945, +0.00 DPS) [rep] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | 27.9 spell_power points (10.77 DPS) | yes | Marshal's Silk Sash (16439, -0.01 DPS) [pvp]; Stormpike Cloth Girdle (19094, -0.65 DPS) [rep]; Highlander's Cloth Girdle (20047, -3.99 DPS, sim-verified) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 55.6 spell_power points (21.50 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -3.85 DPS) [pvp]; Magister's Leggings (16687, -5.07 DPS) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.8 spell_power points (11.53 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.25 DPS) [pvp]; Wolfrunner Shoes (13101, -0.47 DPS) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 27.3 spell_power points (10.56 DPS) | yes | Don Julio's Band (19325, -3.02 DPS) [rep]; Lorekeeper's Ring (19522, -3.51 DPS) [rep] |
| finger2 | Deep Rooted Ring (19109) | Ivus the Forest Lord [world] | 20.2 spell_power points (7.81 DPS) | yes | Don Julio's Band (19325, -0.27 DPS) [rep]; Lorekeeper's Ring (19522, -0.75 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (393.0 DPS) | yes | Ankh of Life (1713, -1.33 DPS) [world_drop]; Shard of the Splithooves (10659, -1.63 DPS) [quest]; Uther's Strength (11302, -1.79 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (393.0 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (393.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -4.10 DPS) [world_drop]; Ironbark Staff (20069, -8.03 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (11.21 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -0.97 DPS) [pvp]; Tome of Arcane Domination (19308, -2.87 DPS, sim-verified) [rep] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (393.0 DPS) | yes | Brilliant Wand (249385, -0.10 DPS) [crafted]; Greater Eternal Wand (249237, -2.65 DPS) [crafted]; Lesser Eternal Wand (249232, -3.45 DPS) [crafted] |

**New at 60:** head: Dreadmist Mask; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Wisdom of the Timbermaw; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Deep Rooted Ring; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 131205213100301531-30000000000000000-053500030000000000)

Set DPS (verified): 678.8. Weights run: 4.3s. Verify run: 18.8s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.657 ± 0.029, crit=0.508 ± 0.011 per rating point (14 rating = 1%, 7.112 per %), hit=1.037 ± 0.051 per rating point (10 rating = 1%, 10.366 per %), spell_haste=not significant (2.088 ± 0.570), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.683 ± 0.002, spirit=not significant (0.438 ± 0.247), mp5=not significant (0.388 ± 0.291)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreadmist Mask (16698) | Scholomance: Darkmaster Gandling [dungeon] | 34.4 spell_power points (15.83 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -0.71 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.6 spell_power points (8.55 DPS) | yes | Mindburst Medallion (11196, -2.97 DPS) [quest]; Triune Amulet (7722, -3.46 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -3.96 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 32.7 spell_power points (15.04 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -0.72 DPS) [pvp]; Theramore Shoulderpads (271953, -1.34 DPS) [rep] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.91 DPS) [crafted]; Stormpike Sage's Cloak (19086, -3.03 DPS) [rep]; Expeditionary's Cape (271907, -3.42 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Silk Vestments (231603, -0.50 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.23 DPS) [pvp]; Magister's Robes (16688, -13.49 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 30.3 spell_power points (13.97 DPS) | yes | Runecloth Cuffs (254123, -5.87 DPS) [crafted]; Marshal's Silk Bracers (16438, -6.20 DPS) [pvp]; Nethershine Cuffs (254069, -6.75 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.8 spell_power points (13.72 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Theramore Handwraps (271945, -4.41 DPS, sim-verified) [rep] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 25.1 spell_power points (11.54 DPS) | yes | Stormpike Cloth Girdle (19094, -0.22 DPS) [rep]; Wisdom of the Timbermaw (19047, -0.42 DPS) [crafted]; Dreadmist Belt (16702, -1.34 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 52.8 spell_power points (24.32 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -5.22 DPS) [pvp]; Magister's Leggings (16687, -8.02 DPS) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 27.9 spell_power points (12.84 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.14 DPS) [pvp]; Earthen Silk Slippers (254013, -0.38 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 25.2 spell_power points (11.62 DPS) | yes | Lorekeeper's Ring (19522, -4.45 DPS) [rep]; Deep Rooted Ring (19109, -13.40 DPS, sim-verified) [world] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lorekeeper's Ring (19522, -0.89 DPS) [rep]; Deep Rooted Ring (19109, -3.67 DPS, sim-verified) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -1.53 DPS) [world_drop]; Ankh of Life (1713, -1.88 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.28 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -3.97 DPS, sim-verified) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Sageclaw (20070, -6.08 DPS) [rep]; Ironbark Staff (20069, -10.98 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.36 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -2.21 DPS) [pvp]; Tome of Arcane Domination (19308, -3.07 DPS) [rep] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torch of Light (279246, -1.26 DPS) [crafted]; Greater Eternal Wand (249237, -2.12 DPS) [crafted]; Nether Force Wand (11263, -2.55 DPS) [quest] |

**New at 60:** head: Dreadmist Mask; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 141005000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 45.0. Weights run: 3.1s. Verify run: 3.4s. 172 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.941 ± 0.013, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.122 per %), hit=0.327 ± 0.015 per rating point (10 rating = 1%, 3.273 per %), spell_haste=not significant (0.121 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.906 ± 0.004, spirit=not significant (0.196 ± 0.052), mp5=0.796 ± 0.101

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 10.0 spell_power points (1.94 DPS) | yes | Shadow Goggles (4373, -0.80 DPS) [crafted]; Pristine Circlet (253949, -0.85 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.79 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.6 spell_power points (0.11 DPS) | yes | Scholarly Pendant (277203, -0.04 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.5 spell_power points (2.62 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.85 DPS) [crafted]; Slime-encrusted Pads (6461, -2.16 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.01 DPS) [crafted]; Sanguine Cape (14376, -0.21 DPS) [world_drop]; Battle Healer's Cloak (20427, -0.28 DPS) [rep] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 13.5 spell_power points (2.62 DPS) | yes | Filigreed Pristine Gown (253901, -0.68 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.92 DPS) [crafted]; Mystic's Wrap (14369, -1.34 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 7.1 spell_power points (1.38 DPS) | yes | Bright Bracers (3647, -0.65 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.80 DPS) [quest]; Owlbeard Bracers (16981, -0.82 DPS) [quest] |
| hands | Pearly Gloves (253921) | Tailoring [crafted] | 9.2 spell_power points (1.79 DPS) | yes | Pristine Gloves (253913, -0.46 DPS) [crafted]; Serpent Gloves (5970, -0.50 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.62 DPS) [world] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 10.1 spell_power points (1.97 DPS) | yes | Keller's Girdle (2911, -0.50 DPS) [world_drop]; Pristine Sash (253925, -0.51 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.64 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 16.4 spell_power points (3.20 DPS) | yes | Filigreed Pristine Leggings (253937, -0.85 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -1.68 DPS) [world_drop]; Silk-threaded Trousers (1929, -1.83 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.8 spell_power points (2.10 DPS) | yes | Pearly Boots (253897, -0.56 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.96 DPS) [crafted]; Red Woolen Boots (4313, -1.13 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 6.6 spell_power points (1.28 DPS) | yes | Loop of Sacrifice (281673, -0.37 DPS) [quest]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop]; Band of Purification (12996, -1.05 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.6 spell_power points (1.10 DPS) | yes | Loop of Sacrifice (281673, -0.18 DPS) [quest]; Volcanic Rock Ring (12053, -0.55 DPS) [world_drop]; Band of Purification (12996, -0.87 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.18 DPS) [dungeon]; Crescent Staff (6505, -0.23 DPS, sim-verified) [quest]; Gnarled Necromancer's Staff (251534, -0.56 DPS) [quest] |
| off_hand | Seer's Fine Stein (7608) | World drop [world_drop] | 3.0 spell_power points (0.59 DPS) | yes | Spellbinder Orb (15926, -0.04 DPS) [dungeon]; Tork Wrench (11855, -0.22 DPS) [quest]; Ancestral Orb (15944, -0.41 DPS) [world_drop] |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cookie's Stirring Rod (5198, -1.00 DPS) [dungeon]; Greater Magic Wand (11288, -1.01 DPS) [crafted]; Sizzle Stick (8071, -1.07 DPS, sim-verified) [quest] |

**New at 20:** head: Pearly Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Pearly Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Evocator's Blade; off_hand: Seer's Fine Stein; ranged: Skycaller

No-known-source sample (15 of 172, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade

### Band 30 (orc, 141005213100300000-00000000000000000-000000000000000000)

Set DPS (verified): 89.8. Weights run: 3.2s. Verify run: 3.8s. 295 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.555 ± 0.010, crit=0.067 ± 0.002 per rating point (14 rating = 1%, 0.936 per %), hit=0.243 ± 0.012 per rating point (10 rating = 1%, 2.433 per %), spell_haste=not significant (0.430 ± 0.125), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.593 ± 0.003, spirit=not significant (0.212 ± 0.081), mp5=not significant (0.199 ± 0.115)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 12.3 spell_power points (2.88 DPS) | yes | Enchanter's Cowl (4322, -0.17 DPS) [crafted]; Silk Headband (7050, -0.77 DPS) [crafted]; Filigreed Pearly Circlet (253983, -0.93 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.3 spell_power points (0.78 DPS) | yes | Kaleidoscope Chain (13084, -0.06 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.11 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.13 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.29 DPS) | yes | Death Speaker Mantle (6685, -0.92 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.94 DPS) [world_drop]; Invoker's Mantle (215365, -0.99 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | sim-verified (89.8 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Hillman's Cloak (3719, -0.02 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.05 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 17.4 spell_power points (4.08 DPS) | yes | Green Silk Armor (7065, -0.27 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.63 DPS) [dungeon]; Death Speaker Robes (6682, -1.00 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.6 spell_power points (2.25 DPS) | yes | Nightsky Wristbands (6407, -1.32 DPS) [world_drop]; Mindthrust Bracers (1974, -1.46 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.71 DPS, sim-verified) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 9.2 spell_power points (2.17 DPS) | yes | Serpent Gloves (5970, -0.52 DPS) [dungeon]; Truefaith Gloves (7049, -0.60 DPS) [crafted]; Jutebraid Gloves (10654, -0.66 DPS, sim-verified) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.7 spell_power points (2.98 DPS) | yes | Warsong Sash (16975, -0.39 DPS) [quest]; Belt of Arugal (6392, -0.65 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.65 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 13.9 spell_power points (3.27 DPS) | yes | Pearly Leggings (253995, -0.16 DPS) [crafted]; Pristine Leggings (253987, -0.46 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.75 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 12.4 spell_power points (2.92 DPS) | yes | Gilded Slippers (254001, -0.17 DPS) [crafted]; Acidic Walkers (9454, -0.51 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (2.06 DPS) | yes | Ring of Power Regulation (273030, -0.51 DPS) [dungeon]; Solid Gold Band (279845, -0.87 DPS) [quest]; Electrocutioner Lagnut (9447, -1.01 DPS) [dungeon] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.4 spell_power points (1.74 DPS) | yes | Ring of Power Regulation (273030, -0.19 DPS) [dungeon]; Solid Gold Band (279845, -0.55 DPS) [quest]; Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (89.8 DPS) | yes | Soulstaff (249393, -1.06 DPS) [crafted]; Tendonscraper Dagger (273053, -1.14 DPS) [world]; Glimmering Staff (249392, -3.40 DPS, sim-verified) [crafted] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (2.43 DPS) | yes | Witch's Finger (16887, -1.38 DPS) [quest]; Satyr's Rod (15962, -1.55 DPS) [world_drop]; Orb of Souls (249395, -2.02 DPS, sim-verified) [crafted] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-verified (89.8 DPS) | yes | Greater Mystic Wand (217287, -0.35 DPS) [crafted]; Wand of Eventide (5214, -0.49 DPS) [world_drop]; Thunderwood (13062, -0.78 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 295, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19545 Scout's Blade

### Band 40 (orc, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 170.5. Weights run: 3.4s. Verify run: 5.1s. 379 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.737 ± 0.022, crit=0.173 ± 0.004 per rating point (14 rating = 1%, 2.415 per %), hit=0.375 ± 0.022 per rating point (10 rating = 1%, 3.754 per %), spell_haste=not significant (0.107 ± 0.297), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.617 ± 0.003, spirit=not significant (0.223 ± 0.250), mp5=not significant (0.228 ± 0.264)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 22.5 spell_power points (5.71 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.54 DPS, sim-verified) [crafted]; Augural Shroud (2620, -1.06 DPS) [world]; Whitemane's Chapeau (7720, -1.39 DPS) [dungeon] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 12.9 spell_power points (3.27 DPS) | yes | Triune Amulet (7722, -0.85 DPS) [dungeon]; Necklace of Calisea (1714, -1.57 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.96 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.6 spell_power points (4.20 DPS) | yes | Green Silken Shoulders (7057, -0.12 DPS) [crafted]; Bloodmage Mantle (7684, -0.24 DPS) [dungeon]; Berylline Pads (4197, -0.56 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.17 DPS) [vendor]; Icy Cloak (4327, -0.34 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 27.8 spell_power points (7.04 DPS) | yes | Robe of the Magi (1716, -0.05 DPS) [world_drop]; Robe of Power (7054, -0.79 DPS) [crafted]; Elemental Raiment (9434, -1.71 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.9 spell_power points (2.51 DPS) | yes | Spidertank Oilrag (9448, -0.05 DPS) [dungeon]; Condor Bracers (15864, -0.73 DPS) [quest]; Windchaser Cuffs (14429, -0.83 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.5 spell_power points (5.71 DPS) | yes | Black Mageweave Gloves (10003, -1.90 DPS) [crafted]; Golden Handwraps (254029, -2.21 DPS) [crafted]; Red Mageweave Gloves (10018, -2.54 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.1 spell_power points (4.58 DPS) | yes | Defiler's Cloth Girdle (20166, -0.28 DPS) [rep]; Golden Cord (254045, -0.56 DPS) [crafted]; Gilded Cord (254037, -0.72 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.8 spell_power points (5.79 DPS) | yes | Crimson Silk Pantaloons (7062, -1.32 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.84 DPS) [dungeon]; Pearly Leggings (253995, -2.01 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.6 spell_power points (6.48 DPS) | yes | Gilded Slippers (254001, -3.17 DPS) [crafted]; Acidic Walkers (9454, -3.49 DPS) [dungeon]; Golden Slippers (254009, -3.65 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (3.47 DPS) | yes | Reedknot Ring (9622, -1.70 DPS) [quest]; Ring of Power Regulation (273030, -1.78 DPS) [dungeon]; Snake Hoop (6750, -4.89 DPS, sim-verified) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.45 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.68 DPS) [quest]; Ring of Power Regulation (273030, -0.76 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Celestial Stave (9517, -7.78 DPS) [quest]; Hypnotic Blade (7714, -8.12 DPS) [dungeon]; Illusionary Rod (7713, -10.66 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dancing Flame (6806, -0.37 DPS) [quest]; Starfaller (13063, -0.40 DPS) [world_drop]; Twisted Nether Wand (249144, -0.53 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 379, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 141005213100301531-40000000000000000-051000000000000000)

Set DPS (verified): 247.1. Weights run: 3.5s. Verify run: 5.2s. 483 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.704 ± 0.028, crit=0.358 ± 0.008 per rating point (14 rating = 1%, 5.006 per %), hit=0.863 ± 0.044 per rating point (10 rating = 1%, 8.627 per %), spell_haste=not significant (0.542 ± 0.603), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.646 ± 0.002, spirit=not significant (0.495 ± 0.248), mp5=not significant (0.524 ± 0.262)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 34.0 spell_power points (12.09 DPS) | yes | Red Mageweave Headband (10033, -0.32 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.32 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.48 DPS) [crafted] |
| neck | Mindburst Medallion (11196) | Broken Alliances [quest] | 12.7 spell_power points (4.52 DPS) | yes | Glowing Eye of Mordresh (10769, -0.02 DPS) [dungeon]; Horizon Choker (13085, -0.31 DPS) [world_drop]; Triune Amulet (7722, -0.44 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 25.7 spell_power points (9.14 DPS) | yes | Nethershine Shoulders (254057, -2.08 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.25 DPS) [vendor]; Nethergeld Shoulders (254049, -2.87 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (247.1 DPS) | yes | Runecloth Cloak (13860, -2.34 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.49 DPS, sim-verified) [quest]; Darkspear Raider's Cloak (272076, -3.15 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 33.1 spell_power points (11.77 DPS) | yes | Dreamweave Vest (10021, -0.65 DPS) [crafted]; Robe of the Magi (1716, -1.56 DPS) [world_drop]; Robe of Power (7054, -2.37 DPS) [crafted] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 15.8 spell_power points (5.62 DPS) | yes | Nethergeld Cuffs (254061, -0.50 DPS) [crafted]; Bloodband Bracers (11469, -1.06 DPS) [quest]; Spidertank Oilrag (9448, -1.86 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 24.3 spell_power points (8.64 DPS) | yes | Golden Gloves (254103, -0.36 DPS) [crafted]; Runecloth Gloves (13863, -0.53 DPS) [crafted]; Gilded Gloves (254095, -1.17 DPS) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 21.0 spell_power points (7.49 DPS) | yes | Ghostweave Cord (254073, -0.57 DPS) [crafted]; Dawnspire Cord (12466, -0.59 DPS) [dungeon]; Golden Waistcord (254089, -0.87 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 25.8 spell_power points (9.20 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.14 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.45 DPS) [dungeon]; Red Mageweave Pants (10009, -1.21 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.5 spell_power points (9.77 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.32 DPS, sim-verified) [vendor]; Golden Sandals (254115, -2.92 DPS) [crafted]; Gilded Sandals (254107, -3.60 DPS) [crafted] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | 14.1 spell_power points (5.02 DPS) | yes | Band of the Unicorn (7553, -0.39 DPS) [world_drop]; Eye of Adaegus (5266, -1.40 DPS) [world_drop]; Snake Hoop (6750, -1.63 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (4.81 DPS) | yes | Band of the Unicorn (7553, -0.19 DPS) [world_drop]; Eye of Adaegus (5266, -1.19 DPS) [world_drop]; Snake Hoop (6750, -1.42 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (247.1 DPS) | yes | Rune of the Guard Captain (19120, -0.61 DPS) [quest]; Uther's Strength (11302, -0.63 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.00 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (247.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (247.1 DPS) | yes | Spire of Hakkar (10844, -5.41 DPS, sim-verified) [world]; Glowing Brightwood Staff (812, -6.08 DPS) [world_drop]; Kindling Stave (11750, -6.59 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | sim-verified (247.1 DPS) | yes | Noxious Shooter (17745, +0.00 DPS) [dungeon]; Lesser Eternal Wand (249232, +0.00 DPS) [crafted]; Charged Lightning Rod (11860, -0.04 DPS) [quest] |

**New at 50:** head: Dreamweave Circlet; neck: Mindburst Medallion; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethershine Cuffs; waist: Satyrmane Sash; legs: Senior Designer's Pantaloons; finger1: Advisor's Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Ankh of Life

No-known-source sample (15 of 483, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 141005213100301531-40000000000000000-053500030000000000)

Set DPS (verified): 388.8. Weights run: 3.6s. Verify run: 22.1s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.748 ± 0.031, crit=0.413 ± 0.010 per rating point (14 rating = 1%, 5.777 per %), hit=0.845 ± 0.041 per rating point (10 rating = 1%, 8.450 per %), spell_haste=not significant (1.114 ± 0.583), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.599 ± 0.003, spirit=not significant (0.229 ± 0.350), mp5=not significant (0.040 ± 0.350)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Dreadmist Mask (16698, -9.22 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 19.5 spell_power points (5.89 DPS) | yes | Glowing Eye of Mordresh (10769, -1.96 DPS) [dungeon]; Mindburst Medallion (11196, -2.37 DPS) [quest]; Horizon Choker (13085, -2.45 DPS) [world_drop] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -6.28 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.34 DPS) [quest]; Runecloth Cloak (13860, -1.47 DPS) [crafted]; Frostwolf Advisor's Cloak (19085, -1.77 DPS) [rep] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Silk Raiment (231596, -0.57 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.27 DPS) [pvp]; Magister's Robes (16688, -8.26 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.6 spell_power points (8.94 DPS) | yes | Runecloth Cuffs (254123, -3.36 DPS) [crafted]; General's Silk Cuffs (16538, -3.96 DPS) [pvp]; Dreadmist Bracers (16703, -4.60 DPS) [dungeon] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Gloves of Spell Mastery (14146, -8.59 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Defiler's Cloth Girdle (20163, +0.00 DPS) [rep]; Frostwolf Cloth Belt (19090, -10.85 DPS, sim-verified) [rep] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -7.36 DPS, sim-verified) [rep] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.0 spell_power points (8.76 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.15 DPS) [pvp]; Earthen Silk Slippers (254013, -1.02 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.6 spell_power points (7.44 DPS) | yes | Don Julio's Band (19325, -3.14 DPS) [rep]; Advisor's Ring (19518, -3.16 DPS) [rep]; Philanthropist's Ring (281635, -3.28 DPS) [quest] |
| finger2 | Deep Rooted Ring (19109) | Ivus the Forest Lord [world] | 19.5 spell_power points (5.89 DPS) | yes | Don Julio's Band (19325, -1.59 DPS) [rep]; Advisor's Ring (19518, -1.61 DPS) [rep]; Philanthropist's Ring (281635, -1.73 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.49 DPS) [world_drop]; Ankh of Life (1713, -1.47 DPS) [world_drop]; Alchemist's Stone (13503, -2.62 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.93 DPS, sim-verified) [world_drop]; Ankh of Life (1713, -0.96 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.10 DPS) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Mindfang (20214, -3.98 DPS) [rep]; Ironbark Staff (20220, -9.04 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (8.77 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -1.07 DPS) [pvp]; Spirit of Aquementas (11904, -3.23 DPS, sim-verified) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Torch of Light (279246, -0.90 DPS, sim-verified) [crafted]; Greater Eternal Wand (249237, -2.00 DPS) [crafted]; Lesser Eternal Wand (249232, -2.80 DPS) [crafted] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Magister's Mantle; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Deep Rooted Ring; trinket2: Rune of the Guard Captain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 131205213100301531-30000000000000000-053500030000000000)

Set DPS (verified): 668.6. Weights run: 4.3s. Verify run: 17.4s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.694 ± 0.037, crit=0.671 ± 0.014 per rating point (14 rating = 1%, 9.396 per %), hit=1.434 ± 0.068 per rating point (10 rating = 1%, 14.344 per %), spell_haste=5.233 ± 0.693, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.740 ± 0.003, spirit=not significant (0.224 ± 0.282), mp5=not significant (0.009 ± 0.330)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 32.9 spell_power points (17.13 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Dreadmist Mask (16698, -0.12 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.9 spell_power points (9.87 DPS) | yes | Glowing Eye of Mordresh (10769, -3.33 DPS) [dungeon]; Mindburst Medallion (11196, -3.99 DPS) [quest]; Horizon Choker (13085, -4.34 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 35.4 spell_power points (18.45 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.30 DPS) [pvp]; Darkspear Shoulderpads (272103, -6.08 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (668.6 DPS) | yes | Deep Woodlands Cloak (19121, -0.66 DPS) [quest]; Runecloth Cloak (13860, -2.58 DPS) [crafted]; Cloak of Flames (3475, -2.69 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (668.6 DPS) | yes | Warlord's Silk Raiment (231596, -1.14 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -7.51 DPS) [pvp]; Magister's Robes (16688, -15.51 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.1 spell_power points (15.17 DPS) | yes | Runecloth Cuffs (254123, -5.82 DPS) [crafted]; Nethershine Cuffs (254069, -7.04 DPS) [crafted]; General's Silk Cuffs (16538, -7.13 DPS) [pvp] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 34.7 spell_power points (18.09 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Raider Handwraps (272097, -8.82 DPS, sim-verified) [vendor] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 27.6 spell_power points (14.36 DPS) | yes | Frostwolf Cloth Belt (19090, -1.36 DPS) [rep]; Wisdom of the Timbermaw (19047, -2.51 DPS) [crafted]; Magister's Belt (16685, -3.12 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 43.4 spell_power points (22.62 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -0.06 DPS) [pvp]; Magister's Leggings (16687, -4.36 DPS, sim-verified) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP reward · Horde [vendor] | 28.6 spell_power points (14.89 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.14 DPS) [dungeon]; Blood Guard's Silk Walkers (227110, -0.34 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.2 spell_power points (12.61 DPS) | yes | Deep Rooted Ring (19109, -2.74 DPS) [world]; Blackstone Ring (17713, -5.14 DPS) [dungeon]; Advisor's Ring (19518, -5.30 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 23.7 spell_power points (12.37 DPS) | yes | Blackstone Ring (17713, -4.89 DPS) [dungeon]; Advisor's Ring (19518, -5.05 DPS) [rep]; Deep Rooted Ring (19109, -5.54 DPS, sim-verified) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (668.6 DPS) | yes | Uther's Strength (11302, -2.10 DPS) [world_drop]; Ankh of Life (1713, -3.83 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -4.06 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (668.6 DPS) | yes | Uther's Strength (11302, -3.60 DPS) [world_drop]; Ankh of Life (1713, -5.32 DPS) [world_drop]; Darkspear Voodoo Seal (272059, -9.16 DPS, sim-verified) [vendor] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (668.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -7.74 DPS) [world_drop]; Ironbark Staff (20220, -15.72 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (15.11 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -2.23 DPS) [pvp]; Tome of Arcane Domination (19308, -5.89 DPS, sim-verified) [rep] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (668.6 DPS) | yes | Greater Eternal Wand (249237, -2.55 DPS) [crafted]; Nether Force Wand (11263, -2.73 DPS) [quest]; Torch of Light (279246, -2.76 DPS) [crafted] |

**New at 60:** head: Red Mageweave Headband; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Maiden's Circle; finger2: Don Julio's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

