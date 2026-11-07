# Leveling BiS: Discipline

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 46.4. Weights run: 13.2s. Verify run: 71.9s. 186 eligible items had no known source.

Race sweep: human 46.4 ± 0.0, gnome 45.9 ± 0.0, night-elf 44.8 ± 0.0, dwarf 44.5 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.027, intellect=0.700 ± 0.005, spirit=0.609 ± 0.006, mp5=2.069 ± 0.017, crit=0.056 ± 0.003 per rating point (14 rating = 1%, 0.781 per %), spell_haste=not significant (0.013 ± 0.043)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.21 DPS) | yes | Shadow Goggles (4373, -0.85 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.91 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 1.2 healing_power points (0.15 DPS) | yes | Tarnished Locket (279870, -0.32 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.3 healing_power points (1.39 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.56 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.90 DPS) [crafted] |
| back | Caretaker's Cape (20428) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.40 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.75 DPS) [crafted]; Black Whelp Cloak (7283, -0.89 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 18.3 healing_power points (2.25 DPS) | yes | Corsair's Overshirt (5202, -0.27 DPS, sim-verified) [dungeon]; Robe of the Moccasin (6465, -0.40 DPS) [dungeon]; Bloody Apron (6226, -0.65 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.7 healing_power points (1.19 DPS) | yes | Repurposed Hair Band (281256, -0.87 DPS) [quest]; Seer's Cuffs (3645, -0.96 DPS) [dungeon]; Bright Bracers (3647, -1.48 DPS, sim-verified) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.1 healing_power points (1.61 DPS) | yes | Serpent Gloves (5970, -0.80 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.87 DPS) [world]; Magefist Gloves (12977, -0.88 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 13.8 healing_power points (1.70 DPS) | yes | Novice Ardent's Sash (253887, -0.44 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.95 DPS) [crafted]; Keller's Girdle (2911, -1.01 DPS) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 24.6 healing_power points (3.03 DPS) | yes | Darkweave Breeches (12987, -1.85 DPS, sim-verified) [world_drop]; Colorful Kilt (10048, -2.04 DPS) [crafted]; Silk-threaded Trousers (1929, -2.17 DPS) [dungeon] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.1 healing_power points (1.36 DPS) | yes | Spidersilk Boots (4320, -0.14 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.50 DPS) [crafted]; Kimbra Boots (6191, -0.80 DPS) [quest] |
| finger1 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 9.1 healing_power points (1.12 DPS) | yes | Lavishly Jeweled Ring (1156, -0.61 DPS) [dungeon]; Band of Purification (12996, -0.67 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.86 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.4 healing_power points (0.79 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS, sim-verified) [dungeon]; Band of Purification (12996, -0.34 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 55.2 healing_power points (6.77 DPS) | yes | Living Root (6631, -1.06 DPS, sim-verified) [dungeon]; Staff of the Blessed Seer (2271, -2.59 DPS) [dungeon]; Evocator's Blade (2567, -2.79 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.25 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Flaring Baton (5326, -0.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Caretaker's Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Staff of Westfall; ranged: Sizzle Stick

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (human, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 93.1. Weights run: 15.1s. Verify run: 80.8s. 308 eligible items had no known source.

Race sweep: human 93.1 ± 0.0, gnome 89.7 ± 0.0, night-elf 87.9 ± 0.0, dwarf 87.2 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.036, intellect=0.715 ± 0.008, spirit=1.155 ± 0.009, mp5=2.499 ± 0.009, crit=0.087 ± 0.004 per rating point (14 rating = 1%, 1.217 per %), spell_haste=not significant (0.226 ± 0.085)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.9 healing_power points (5.42 DPS) | yes | Pristine Circlet (253949, -2.98 DPS) [crafted]; Filigreed Pristine Circlet (253975, -3.14 DPS, sim-verified) [crafted]; Nightsky Cowl (4039, -3.47 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.5 healing_power points (1.01 DPS) | yes | Glowing Green Talisman (5002, -0.07 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop]; Pendant of Myzrael (4614, -0.28 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.4 healing_power points (2.09 DPS) | yes | Death Speaker Mantle (6685, -0.13 DPS, sim-verified) [dungeon]; Mantle of Honor (3560, -0.32 DPS) [quest]; Nightsky Mantle (4718, -0.53 DPS) [world_drop] |
| back | Caretaker's Cape (19533) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.09 DPS, sim-verified) [dungeon]; Repairman's Cape (9605, -0.50 DPS) [quest]; Prelacy Cape (7004, -0.97 DPS) [quest] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 29.6 healing_power points (4.02 DPS) | yes | Death Speaker Robes (6682, -0.33 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.38 DPS) [dungeon]; Robes of Arugal (6324, -0.64 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.5 healing_power points (2.24 DPS) | yes | Mindthrust Bracers (1974, -0.63 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.19 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 17.4 healing_power points (2.37 DPS) | yes | Pristine Gloves (253913, -0.58 DPS) [crafted]; Gloves of Meditation (4318, -0.59 DPS) [crafted]; Truefaith Gloves (7049, -0.86 DPS, sim-verified) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 16.9 healing_power points (2.30 DPS) | yes | Highlander's Cloth Girdle (20099, -0.51 DPS) [rep]; Novice Ardent's Sash (253887, -0.72 DPS) [crafted]; Pristine Sash (253925, -0.80 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 32.8 healing_power points (4.45 DPS) | yes | Filigreed Pristine Leggings (253937, -0.80 DPS) [crafted]; Earthen Leggings (253999, -0.98 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.41 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 29.6 healing_power points (4.02 DPS) | yes | Boots of the Enchanter (4325, -2.09 DPS) [crafted]; Glinteye Slippers (273024, -2.13 DPS) [dungeon]; Acidic Walkers (9454, -2.32 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.1 healing_power points (2.32 DPS) | yes | Ring of Power Regulation (273030, -0.49 DPS) [dungeon]; Lorekeeper's Ring (19525, -0.69 DPS) [rep]; The Queen's Jewel (13094, -0.87 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 17.1 healing_power points (2.32 DPS) | yes | Ring of Power Regulation (273030, -0.53 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19525, -0.69 DPS) [rep]; The Queen's Jewel (13094, -0.87 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.29 DPS) [dungeon]; Staff of Westfall (2042, -1.29 DPS) [quest]; Staff of the Friar (3415, -2.95 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 19.9 healing_power points (2.70 DPS) | yes | Eye of Paleth (2943, -0.94 DPS) [quest]; Orb of Mistmantle (13031, -1.08 DPS, sim-verified) [world_drop]; Orb of Mystic Insight (249394, -1.17 DPS) [crafted] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 7.9 healing_power points (1.07 DPS) | yes | Wand of Eventide (5214, -0.39 DPS) [world_drop]; Greater Mystic Wand (217287, -0.39 DPS) [crafted]; Thunderwood (13062, -0.81 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Caretaker's Cape; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 005005031330011000-33300000000000000-00000000000000000)

Set DPS (verified): 141.9. Weights run: 25.3s. Verify run: 161.8s. 402 eligible items had no known source.

Race sweep: human 141.9 ± 0.1, gnome 138.4 ± 0.0, night-elf 136.2 ± 0.0, dwarf 135.1 ± 0.0 (heal_score); the guide race human is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.152, intellect=1.632 ± 0.029, spirit=1.991 ± 0.031, mp5=3.897 ± 0.029, crit=0.336 ± 0.020 per rating point (14 rating = 1%, 4.710 per %), spell_haste=2.787 ± 0.501

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 67.7 healing_power points (8.26 DPS) | yes | Whitemane's Chapeau (7720, -1.07 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -2.44 DPS) [dungeon]; Holy Shroud (2721, -2.77 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.1 healing_power points (4.04 DPS) | yes | Necklace of Calisea (1714, -1.27 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -1.36 DPS) [dungeon]; Amberglow Talisman (10824, -1.61 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 43.2 healing_power points (5.28 DPS) | yes | Earthen Silk Shoulders (254033, -0.75 DPS, sim-verified) [crafted]; Mantle of Doan (7712, -1.28 DPS) [dungeon]; Mistscape Mantle (4734, -1.87 DPS) [dungeon] |
| back | Caretaker's Cape (19532) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.61 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.76 DPS) [pvp]; Expeditionary's Cape (271922, -0.96 DPS, sim-verified) [rep] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 73.9 healing_power points (9.03 DPS) | yes | Doomsayer's Robe (4746, -2.70 DPS) [quest]; Civinad Robes (9623, -2.93 DPS) [quest]; Dreamweave Vest (10021, -3.01 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.7 healing_power points (2.53 DPS) | yes | Mindthrust Bracers (1974, -0.10 DPS) [dungeon]; Earthen Silk Cuffs (254019, -0.25 DPS, sim-verified) [crafted]; Silkstream Cuffs (16791, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 38.5 healing_power points (4.70 DPS) | yes | Gilded Handwraps (254021, -0.62 DPS) [crafted]; Earthen Silk Gloves (254017, -0.62 DPS, sim-verified) [crafted]; Gloves of Old (9395, -0.78 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.0 healing_power points (5.74 DPS) | yes | Windchaser Cinch (14435, -1.96 DPS) [world_drop]; Sutarn's Ring (13105, -2.29 DPS) [world_drop]; Deathmage Sash (10771, -2.48 DPS, sim-verified) [dungeon] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 56.9 healing_power points (6.95 DPS) | yes | Pristine Leggings (253987, -1.66 DPS) [crafted]; Stormcloth Pants (10010, -1.89 DPS, sim-verified) [crafted]; Aurora Pants (4044, -2.33 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Thoughtcast Boots (10578, -0.06 DPS) [dungeon]; Furen's Boots (13100, -0.33 DPS, sim-verified) [world_drop]; Earthen Silk Slippers (254013, -0.91 DPS) [crafted] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 22.9 healing_power points (2.80 DPS) | yes | Welken Ring (5011, -0.10 DPS) [world_drop]; Lorekeeper's Ring (19524, -0.27 DPS) [rep]; The Queen's Jewel (13094, -0.46 DPS) [world_drop] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 22.9 healing_power points (2.80 DPS) | yes | Lorekeeper's Ring (19524, -0.27 DPS) [rep]; Welken Ring (5011, -0.43 DPS, sim-verified) [world_drop]; The Queen's Jewel (13094, -0.46 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Friar (3415, -3.56 DPS) [dungeon]; Odo's Ley Staff (6318, -4.16 DPS) [dungeon]; Staff of Jordan (873, -5.04 DPS, sim-verified) [world_drop] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 30.0 healing_power points (3.66 DPS) | yes | Orb of Souls (249395, -0.61 DPS) [crafted]; Orb of Lorica (11262, -0.84 DPS, sim-verified) [quest]; Aurora Sphere (7610, -1.29 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.8 healing_power points (1.92 DPS) | yes | Wand of Mana Concentration (273043, -0.18 DPS) [world]; Flash Wand (5248, -0.40 DPS) [quest]; Goblin Igniter (5253, -0.40 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Caretaker's Cape; chest: Stormcloth Vest; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Stoneweaver Leggings; trinket1: Ankh of Life; trinket2: Enriched Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 198.5. Weights run: 26.5s. Verify run: 242.0s. 517 eligible items had no known source.

Race sweep: human 198.6 ± 0.1, gnome 190.6 ± 0.1, night-elf 187.4 ± 0.1, dwarf 186.1 ± 0.1 (heal_score); the guide race human is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.098, intellect=1.300 ± 0.031, spirit=2.718 ± 0.030, mp5=6.029 ± 0.034, crit=0.393 ± 0.016 per rating point (14 rating = 1%, 5.504 per %), spell_haste=not significant (0.971 ± 0.548)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Satin Cover (220896, +0.00 DPS) [vendor]; Cassandra's Grace (13102, -0.01 DPS) [world_drop]; Gemburst Circlet (10751, -0.30 DPS, sim-verified) [quest] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Gemshard Heart (17707, -0.10 DPS) [dungeon]; Horizon Choker (13085, -0.12 DPS) [world_drop]; Lei of Lilies (1315, -0.43 DPS, sim-verified) [world_drop] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 52.7 healing_power points (4.65 DPS) | yes | Knight-Lieutenant's Satin Pads (220894, +0.00 DPS) [vendor]; Kentic Amice (11624, -0.64 DPS, sim-verified) [dungeon]; Hibernal Mantle (8111, -1.11 DPS) [world_drop] |
| back | Caretaker's Cape (19531) | Silverwing Sentinels [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Spritecaster Cape (11623, -0.02 DPS) [dungeon]; Wingveil Cloak (10802, -0.30 DPS) [dungeon]; Featherskin Cape (10843, -1.52 DPS, sim-verified) [world] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Wind Serpent (12462, -0.70 DPS) [world]; Vestments of the Atal'ai Prophet (10806, -2.11 DPS) [dungeon]; Stormcloth Vest (10020, -2.81 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 42.7 healing_power points (3.77 DPS) | yes | Aristocratic Cuffs (12546, -1.12 DPS, sim-verified) [dungeon]; Spidertank Oilrag (9448, -1.38 DPS) [dungeon]; Venomshroud Armguards (14439, -1.39 DPS) [world_drop] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 63.7 healing_power points (5.62 DPS) | yes | Gloves of the Atal'ai Prophet (10808, -0.83 DPS) [dungeon]; Earthenweave Gloves (254075, -1.03 DPS, sim-verified) [crafted]; Ghostweave Gloves (14142, -1.23 DPS) [crafted] |
| waist | Earthenweave Cord (254077) | Tailoring [crafted] | 55.9 healing_power points (4.93 DPS) | yes | Gilded Cord (254037, -0.59 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -1.06 DPS) [crafted]; Gilded Waistcord (254081, -1.34 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 83.0 healing_power points (7.33 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.41 DPS) [dungeon]; Rainstrider Leggings (11123, -0.85 DPS) [quest]; Dalewind Trousers (13008, -0.95 DPS, sim-verified) [world_drop] |
| feet | Earthenweave Boots (254093) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Furen's Boots (13100, -0.29 DPS) [world_drop]; Vinerot Sandals (17748, -0.44 DPS) [dungeon]; Mistwalker Boots (10629, -1.21 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 40.4 healing_power points (3.57 DPS) | yes | Choking Band (11868, -0.45 DPS) [quest]; Brainlash (6440, -0.65 DPS) [dungeon]; Cyclopean Band (11824, -1.01 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 36.1 healing_power points (3.19 DPS) | yes | Brainlash (6440, -0.27 DPS) [dungeon]; Cyclopean Band (11824, -0.63 DPS) [dungeon]; Choking Band (11868, -1.01 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271903, -0.22 DPS) [rep]; Evonice's Landin' Pilla (18951, -0.48 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.96 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Enriched Seal (271903, -0.69 DPS, sim-verified) [rep] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, -4.90 DPS) [crafted]; Charstone Dirk (17710, -6.16 DPS) [dungeon]; Resurgence Rod (17743, -7.49 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cairnstone Sliver (9654) | The Morrow Stone [quest] | 20.1 healing_power points (1.77 DPS) | yes | Wand of Mana Concentration (273043, -0.25 DPS, sim-verified) [world]; Jaina's Firestarter (13064, -0.36 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -0.52 DPS) [quest] |

**New at 50:** shoulder: Nethergeld Shoulders; back: Caretaker's Cape; chest: Ghostweave Vest; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Earthenweave Cord; legs: Senior Designer's Pantaloons; feet: Earthenweave Boots; finger1: Eye of Adaegus; finger2: Lorekeeper's Ring; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Cairnstone Sliver

No-known-source sample (15 of 517, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 287.7. Weights run: 27.3s. Verify run: 836.3s. 1359 eligible items had no known source.

Race sweep: human 287.5 ± 0.2, gnome 286.8 ± 0.2, night-elf 274.6 ± 0.2, dwarf 270.8 ± 0.2 (heal_score); the guide race human is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.053, intellect=1.517 ± 0.066, spirit=2.679 ± 0.052, mp5=5.318 ± 0.086, crit=0.549 ± 0.033 per rating point (14 rating = 1%, 7.688 per %), spell_haste=not significant (2.247 ± 1.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 96.6 healing_power points (9.61 DPS) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -6.50 DPS, sim-verified) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 55.6 healing_power points (5.53 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.15 DPS) [world_drop]; Animated Chain Necklace (18723, -0.65 DPS) [dungeon]; Heart of the Fiend (13960, -0.78 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -6.01 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -0.48 DPS) [rep]; Featherskin Cape (10843, -0.60 DPS) [world]; Frostweaver Cape (12968, -2.03 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 115.9 healing_power points (11.53 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Embrace of the Wind Serpent (12462, -0.97 DPS) [world]; Mooncloth Vest (14138, -2.83 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmthalak's Shackles (13958, +0.00 DPS) [quest]; Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, -4.57 DPS, sim-verified) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Swarmtender's Gloves (275607, -0.61 DPS) [crafted]; Magister's Gloves (16684, -0.99 DPS) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 87.3 healing_power points (8.69 DPS) | yes | Marshal's Satin Sash (17609, -1.22 DPS) [pvp]; Devout Belt (16696, -2.17 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.81 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 108.4 healing_power points (10.78 DPS) | yes | Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Knight-Captain's Satin Legguards (227125, -0.81 DPS) [pvp]; Haunting Specter Leggings (11929, -10.47 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Satin Walkers (231627, -0.04 DPS) [vendor]; Faith Healer's Boots (22247, -4.62 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; The Postmaster's Seal (13392, -0.23 DPS) [dungeon]; Maiden's Circle (13001, -0.77 DPS) [world_drop] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; The Postmaster's Seal (13392, -0.55 DPS) [dungeon]; Maiden's Circle (13001, -1.08 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Shard of the Splithooves (10659, -0.02 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.53 DPS) [quest]; Enriched Seal (271903, -0.55 DPS) [rep] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, -1.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.51 DPS) [quest]; Second Wind (11819, -3.15 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Warhammer (23454, +0.00 DPS) [vendor]; Resurgence Rod (17743, -4.74 DPS) [quest]; Hand of Edward the Odd (2243, -5.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 46.5 healing_power points (4.62 DPS) | yes | Brilliant Wand (249385, -2.67 DPS) [crafted]; Jaina's Firestarter (13064, -2.92 DPS) [world_drop]; Cairnstone Sliver (9654, -4.71 DPS, sim-verified) [quest] |

**New at 60:** head: Devout Crown; neck: Lady Maye's Pendant; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Penitent's Cinch; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Rosewine Circle; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1359, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 600.0. Weights run: 18.6s. Verify run: 592.5s. 1359 eligible items had no known source.

Race sweep: human 599.7 ± 0.3, gnome 596.3 ± 0.3, night-elf 582.2 ± 0.3, dwarf 579.2 ± 0.3 (heal_score); the guide race human is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.306, intellect=1.044 ± 0.056, spirit=0.949 ± 0.055, mp5=2.727 ± 0.064, crit=0.858 ± 0.047 per rating point (14 rating = 1%, 12.006 per %), spell_haste=not significant (3.646 ± 1.236)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Animated Chain Necklace (18723, +0.00 DPS) [dungeon]; Jeweled Amulet of Cainwyn (1443, -0.02 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19530, -2.45 DPS) [rep]; Spritecaster Cape (11623, -4.70 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 88.2 healing_power points (25.98 DPS) | yes | Field Marshal's Satin Tunic (231624, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -1.33 DPS) [dungeon]; Field Marshal's Satin Vestments (17605, -4.44 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Nethergeld Cuffs (254061, +0.00 DPS) [crafted] |
| hands | Theramore Handwraps (271945) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; Marshal's Satin Handwraps (231623, +0.00 DPS) [vendor]; Gilded Gloves (254095, -1.52 DPS) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Penitent's Cinch (272394, +0.00 DPS) [vendor]; Gilded Waistcord (254081, -0.60 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Satin Legguards (227125, +0.00 DPS) [pvp]; Marshal's Satin Legguards (231626, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, +0.00 DPS) [vendor] |
| feet | Mooncloth Boots (15802) | Tailoring [crafted] | 51.0 healing_power points (15.01 DPS) | yes | Marshal's Satin Walkers (231627, +0.00 DPS) [vendor]; Faith Healer's Boots (22247, -0.31 DPS) [dungeon]; Swarmtender's Footpads (275609, -2.34 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, -5.89 DPS) [dungeon]; Weakness Analyzer (272438, -5.89 DPS) [vendor]; Uther's Strength (11302, -7.07 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -13.79 DPS) [quest]; Grand Marshal's Stave (234571, -17.62 DPS) [pvp]; Ironbark Staff (20069, -18.01 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.4 healing_power points (8.09 DPS) | yes | Brilliant Wand (249385, -3.28 DPS) [crafted]; Bonecreeper Stylus (13938, -3.61 DPS) [dungeon]; Oblivion's Touch (18761, -4.70 DPS) [dungeon] |

**New at 60:** head: Devout Crown; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Theramore Handwraps; waist: Devout Belt; legs: Devout Skirt; feet: Mooncloth Boots; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1359, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 005005000010000000-00000000000000000-00000000000000000)

Set DPS (verified): 45.4. Weights run: 13.2s. Verify run: 69.1s. 179 eligible items had no known source.

Race sweep: undead 45.4 ± 0.0, troll 43.4 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.027, intellect=0.700 ± 0.005, spirit=0.609 ± 0.006, mp5=2.069 ± 0.017, crit=0.056 ± 0.003 per rating point (14 rating = 1%, 0.781 per %), spell_haste=not significant (0.013 ± 0.043)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 18.0 healing_power points (2.21 DPS) | yes | Shadow Goggles (4373, -1.30 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -1.91 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.8 healing_power points (0.22 DPS) | yes | Scholarly Pendant (277203, -0.07 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.3 healing_power points (1.39 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.42 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.90 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.44 DPS, sim-verified) [crafted]; Pearl-clasped Cloak (5542, -0.75 DPS) [crafted]; Black Whelp Cloak (7283, -0.89 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 18.3 healing_power points (2.25 DPS) | yes | Corsair's Overshirt (5202, -0.30 DPS) [dungeon]; Robe of the Moccasin (6465, -0.40 DPS) [dungeon]; Bloody Apron (6226, -0.65 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 9.7 healing_power points (1.19 DPS) | yes | Bright Bracers (3647, -0.85 DPS) [world_drop]; Owlbeard Bracers (16981, -0.90 DPS) [quest]; Tabitha's Cuffs (251486, -1.27 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 13.1 healing_power points (1.61 DPS) | yes | Serpent Gloves (5970, -0.81 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.87 DPS) [world]; Magefist Gloves (12977, -0.88 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 13.8 healing_power points (1.70 DPS) | yes | Novice Ardent's Sash (253887, -0.51 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.95 DPS) [crafted]; Keller's Girdle (2911, -1.01 DPS) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 24.6 healing_power points (3.03 DPS) | yes | Darkweave Breeches (12987, -1.93 DPS, sim-verified) [world_drop]; Colorful Kilt (10048, -2.04 DPS) [crafted]; Silk-threaded Trousers (1929, -2.17 DPS) [dungeon] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 11.1 healing_power points (1.36 DPS) | yes | Spidersilk Boots (4320, -0.07 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.50 DPS) [crafted]; Woolen Boots (2583, -0.80 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 9.1 healing_power points (1.12 DPS) | yes | Loop of Sacrifice (281673, -0.69 DPS) [quest]; Volcanic Rock Ring (12053, -0.86 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.48 DPS, sim-verified) [dungeon] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | sim-verified (45.4 DPS) | yes | Loop of Sacrifice (281673, -0.02 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.21 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 48.0 healing_power points (5.90 DPS) | yes | Staff of the Blessed Seer (2271, -1.25 DPS, sim-verified) [dungeon]; Kris of Orgrimmar (15443, -1.85 DPS) [quest]; Evocator's Blade (2567, -1.92 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) (or Greater Magic Wand (11288), Torchlight Wand (5240)) | Deviate Eradication [quest] | 2.0 healing_power points (0.25 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Wand of Decay (5252, -0.07 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Living Root; ranged: Sizzle Stick

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 005005031330001000-00000000000000000-00000000000000000)

Set DPS (verified): 94.6. Weights run: 15.1s. Verify run: 77.3s. 301 eligible items had no known source.

Race sweep: undead 94.6 ± 0.0, troll 87.1 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.036, intellect=0.715 ± 0.008, spirit=1.155 ± 0.009, mp5=2.499 ± 0.009, crit=0.087 ± 0.004 per rating point (14 rating = 1%, 1.217 per %), spell_haste=not significant (0.226 ± 0.085)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 39.9 healing_power points (5.42 DPS) | yes | Pristine Circlet (253949, -2.98 DPS) [crafted]; Filigreed Pristine Circlet (253975, -3.06 DPS, sim-verified) [crafted]; Nightsky Cowl (4039, -3.47 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 7.5 healing_power points (1.01 DPS) | yes | Glowing Green Talisman (5002, -0.07 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.16 DPS) [world_drop]; Pendant of Myzrael (4614, -0.24 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.4 healing_power points (2.09 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS) [dungeon]; Ghostly Mantle (3324, -0.34 DPS) [quest]; Nightsky Mantle (4718, -0.53 DPS) [world_drop] |
| back | Battle Healer's Cloak (19529) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.19 DPS) [dungeon]; Windsong Drape (15468, -1.09 DPS) [quest]; Darkspear Raider's Cloak (272078, -1.14 DPS) [vendor] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | 29.6 healing_power points (4.02 DPS) | yes | Death Speaker Robes (6682, -0.11 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.38 DPS) [dungeon]; Robes of Arugal (6324, -0.64 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 16.5 healing_power points (2.24 DPS) | yes | Mindthrust Bracers (1974, -0.55 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.19 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.46 DPS) [world_drop] |
| hands | Gloves of Old (9395) | World drop [world_drop] | 17.4 healing_power points (2.37 DPS) | yes | Truefaith Gloves (7049, -0.55 DPS, sim-verified) [crafted]; Pristine Gloves (253913, -0.58 DPS) [crafted]; Gloves of Meditation (4318, -0.59 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 16.9 healing_power points (2.30 DPS) | yes | Pristine Sash (253925, -0.46 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -0.51 DPS) [rep]; Novice Ardent's Sash (253887, -0.72 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 32.8 healing_power points (4.45 DPS) | yes | Filigreed Pristine Leggings (253937, -0.80 DPS) [crafted]; Earthen Leggings (253999, -1.11 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.41 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 29.6 healing_power points (4.02 DPS) | yes | Boots of the Enchanter (4325, -2.09 DPS) [crafted]; Glinteye Slippers (273024, -2.13 DPS) [dungeon]; Acidic Walkers (9454, -2.26 DPS, sim-verified) [dungeon] |
| finger1 | Electrocutioner Lagnut (9447) (or Runemender's Seal (273031)) | Gnomeregan: Electrocutioner 6000 [dungeon] | 17.1 healing_power points (2.32 DPS) | yes | Solid Gold Band (279845, -0.04 DPS) [quest]; Ring of Power Regulation (273030, -0.49 DPS) [dungeon]; Advisor's Ring (19521, -0.69 DPS) [rep] |
| finger2 | Runemender's Seal (273031) | Arcane Anomaly [world] | 17.1 healing_power points (2.32 DPS) | yes | Solid Gold Band (279845, -0.08 DPS, sim-verified) [quest]; Ring of Power Regulation (273030, -0.49 DPS) [dungeon]; Advisor's Ring (19521, -0.69 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Death Speaker Scepter (2816) | Razorfen Kraul: Death Speaker Jargba [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Odo's Ley Staff (6318, -0.29 DPS) [dungeon]; Living Root (6631, -2.34 DPS) [dungeon]; Staff of the Friar (3415, -2.98 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Souls (249395) | Enchanting [crafted] | 19.9 healing_power points (2.70 DPS) | yes | Orb of Mistmantle (13031, -1.03 DPS, sim-verified) [world_drop]; Orb of Mystic Insight (249394, -1.17 DPS) [crafted]; Strength of Will (4837, -1.76 DPS) [vendor] |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 7.9 healing_power points (1.07 DPS) | yes | Wand of Eventide (5214, -0.39 DPS) [world_drop]; Greater Mystic Wand (217287, -0.39 DPS) [crafted]; Thunderwood (13062, -0.75 DPS, sim-verified) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Battle Healer's Cloak; chest: Pristine Gown; wrist: Spidertank Oilrag; hands: Gloves of Old; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Electrocutioner Lagnut; finger2: Runemender's Seal; main_hand: Death Speaker Scepter; off_hand: Orb of Souls; ranged: Wand of Mana Concentration

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 005005031330011000-33300000000000000-00000000000000000)

Set DPS (verified): 146.9. Weights run: 25.3s. Verify run: 160.2s. 395 eligible items had no known source.

Race sweep: undead 146.9 ± 0.1, troll 134.3 ± 0.0 (heal_score); the guide race undead is best.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.152, intellect=1.632 ± 0.029, spirit=1.991 ± 0.031, mp5=3.897 ± 0.029, crit=0.336 ± 0.020 per rating point (14 rating = 1%, 4.710 per %), spell_haste=2.787 ± 0.501

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 67.7 healing_power points (8.26 DPS) | yes | Whitemane's Chapeau (7720, -0.91 DPS) [dungeon]; Miner's Hat of the Deep (9429, -2.44 DPS) [dungeon]; Holy Shroud (2721, -2.77 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 33.1 healing_power points (4.04 DPS) | yes | Necklace of Calisea (1714, -0.94 DPS) [dungeon]; Triune Amulet (7722, -1.36 DPS) [dungeon]; Amberglow Talisman (10824, -1.61 DPS) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 43.2 healing_power points (5.28 DPS) | yes | Earthen Silk Shoulders (254033, -1.13 DPS) [crafted]; Mantle of Doan (7712, -1.28 DPS) [dungeon]; Mistscape Mantle (4734, -1.87 DPS) [dungeon] |
| back | Battle Healer's Cloak (19528) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272077, -0.25 DPS) [vendor]; Cloak of Blight (6832, -0.25 DPS) [quest]; Glowing Thresher Cape (6901, -0.61 DPS) [dungeon] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | 73.9 healing_power points (9.03 DPS) | yes | Dreamweave Vest (10021, -1.63 DPS) [crafted]; Doomsayer's Robe (4746, -2.70 DPS) [quest]; Civinad Robes (9623, -2.93 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.7 healing_power points (2.53 DPS) | yes | Earthen Silk Cuffs (254019, -0.09 DPS) [crafted]; Mindthrust Bracers (1974, -0.10 DPS) [dungeon]; Radiant Silver Bracers (4545, -0.44 DPS) [quest] |
| hands | Gilded Handwraps (254021) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; Earthen Silk Gloves (254017, +0.00 DPS) [crafted]; Gloves of Old (9395, -0.17 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | 47.0 healing_power points (5.74 DPS) | yes | Deathmage Sash (10771, -1.90 DPS) [dungeon]; Windchaser Cinch (14435, -1.96 DPS) [world_drop]; Sutarn's Ring (13105, -2.29 DPS) [world_drop] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 56.9 healing_power points (6.95 DPS) | yes | Stormcloth Pants (10010, -1.48 DPS) [crafted]; Pristine Leggings (253987, -1.66 DPS) [crafted]; Aurora Pants (4044, -2.33 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Furen's Boots (13100, +0.00 DPS) [world_drop]; Thoughtcast Boots (10578, -0.06 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.91 DPS) [crafted] |
| finger1 | Darkspear Signet (272070) | Creeg Bothunk [vendor] | 23.4 healing_power points (2.86 DPS) | yes | Electrocutioner Lagnut (9447, -0.05 DPS) [dungeon]; Runemender's Seal (273031, -0.05 DPS) [world]; Welken Ring (5011, -0.16 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Welken Ring (5011, +0.00 DPS) [world_drop]; Electrocutioner Lagnut (9447, +0.00 DPS) [dungeon]; Runemender's Seal (273031, +0.00 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Hand of Righteousness (7721) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Staff of the Friar (3415, -3.56 DPS) [dungeon]; Odo's Ley Staff (6318, -4.16 DPS) [dungeon] |
| off_hand | Beacon of Hope (9393) | Uldaman [dungeon] | 30.0 healing_power points (3.66 DPS) | yes | Orb of Souls (249395, -0.61 DPS) [crafted]; Omega Orb (7749, -1.13 DPS) [quest]; Prophetic Cane (6803, -1.27 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 15.8 healing_power points (1.92 DPS) | yes | Wand of Mana Concentration (273043, -0.18 DPS) [world]; Flash Wand (5248, -0.40 DPS) [quest]; Goblin Igniter (5253, -0.40 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Battle Healer's Cloak; chest: Stormcloth Vest; hands: Gilded Handwraps; waist: Gilded Cord; legs: Stoneweaver Leggings; finger1: Darkspear Signet; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Darkspear Voodoo Seal; main_hand: Hand of Righteousness; off_hand: Beacon of Hope; ranged: Jaina's Firestarter

No-known-source sample (15 of 395, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 005005031330011430-33501000000000000-00000000000000000)

Set DPS (verified): 202.5. Weights run: 26.5s. Verify run: 238.3s. 510 eligible items had no known source.

Race sweep: undead 202.5 ± 0.1, troll 184.6 ± 0.1 (heal_score); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.098, intellect=1.300 ± 0.031, spirit=2.718 ± 0.030, mp5=6.029 ± 0.034, crit=0.393 ± 0.016 per rating point (14 rating = 1%, 5.504 per %), spell_haste=not significant (0.971 ± 0.548)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Cassandra's Grace (13102) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Papal Fez (9431, +0.00 DPS) [dungeon]; Gemburst Circlet (10751, +0.00 DPS) [quest]; Blood Guard's Satin Cover (220899, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lei of Lilies (1315, +0.00 DPS) [world_drop]; Gemshard Heart (17707, -0.10 DPS) [dungeon]; Horizon Choker (13085, -0.12 DPS) [world_drop] |
| shoulder | Nethergeld Shoulders (254049) | Tailoring [crafted] | 52.7 healing_power points (4.65 DPS) | yes | Blood Guard's Satin Pads (220901, +0.00 DPS) [vendor]; Kentic Amice (11624, -0.49 DPS) [dungeon]; Hibernal Mantle (8111, -1.11 DPS) [world_drop] |
| back | Battle Healer's Cloak (19527) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Featherskin Cape (10843, +0.00 DPS) [world]; Spritecaster Cape (11623, -0.02 DPS) [dungeon]; Cloak of Blight (6832, -0.26 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Vestments of the Atal'ai Prophet (10806, +0.00 DPS) [dungeon]; Embrace of the Wind Serpent (12462, +0.00 DPS) [world]; Ghostweave Vest (14141, +0.00 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 42.7 healing_power points (3.77 DPS) | yes | Aristocratic Cuffs (12546, -0.61 DPS) [dungeon]; Spidertank Oilrag (9448, -1.38 DPS) [dungeon]; Venomshroud Armguards (14439, -1.39 DPS) [world_drop] |
| hands | Gilded Gloves (254095) | Tailoring [crafted] | 63.7 healing_power points (5.62 DPS) | yes | Earthenweave Gloves (254075, -0.69 DPS) [crafted]; Gloves of the Atal'ai Prophet (10808, -0.83 DPS) [dungeon]; Ghostweave Gloves (14142, -1.23 DPS) [crafted] |
| waist | Gilded Waistcord (254081) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gilded Cord (254037, +0.00 DPS) [crafted]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Earthenweave Cord (254077, +0.00 DPS) [crafted] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 83.0 healing_power points (7.33 DPS) | yes | Dalewind Trousers (13008, -0.30 DPS) [world_drop]; Kilt of the Atal'ai Prophet (10807, -0.41 DPS) [dungeon]; Rainstrider Leggings (11123, -0.85 DPS) [quest] |
| feet | Earthenweave Boots (254093) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mistwalker Boots (10629, +0.00 DPS) [dungeon]; Furen's Boots (13100, -0.29 DPS) [world_drop]; Vinerot Sandals (17748, -0.44 DPS) [dungeon] |
| finger1 | Advisor's Ring (19519) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Adaegus (5266, +0.00 DPS) [world_drop]; Brainlash (6440, -0.27 DPS) [dungeon]; Cyclopean Band (11824, -0.63 DPS) [dungeon] |
| finger2 | Darkspear Signet (272069) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Adaegus (5266, -0.16 DPS) [world_drop]; Brainlash (6440, -0.80 DPS) [dungeon]; Cyclopean Band (11824, -1.17 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.48 DPS) [quest]; Uther's Strength (11302, -1.29 DPS) [world_drop]; Alchemist's Stone (13503, -1.92 DPS) [crafted] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, -0.26 DPS) [quest]; Uther's Strength (11302, -1.07 DPS) [world_drop]; Alchemist's Stone (13503, -1.70 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -2.99 DPS) [quest]; Dreamstaff (249454, -4.90 DPS) [crafted]; Charstone Dirk (17710, -6.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | 17.3 healing_power points (1.52 DPS) | yes | Chillnail Splinter (10704, -0.08 DPS) [quest]; Jaina's Firestarter (13064, -0.11 DPS) [world_drop]; Captain Rackmore's Tiller (16789, -0.27 DPS) [quest] |

**New at 50:** head: Cassandra's Grace; shoulder: Nethergeld Shoulders; back: Battle Healer's Cloak; wrist: Nethergeld Cuffs; hands: Gilded Gloves; waist: Gilded Waistcord; legs: Senior Designer's Pantaloons; feet: Earthenweave Boots; finger1: Advisor's Ring; finger2: Darkspear Signet; trinket2: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff; ranged: Wand of Mana Concentration

No-known-source sample (15 of 510, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 288.4. Weights run: 27.3s. Verify run: 844.9s. 1354 eligible items had no known source.

Race sweep: undead 288.4 ± 0.2, troll 271.3 ± 0.2 (heal_score); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.053, intellect=1.517 ± 0.066, spirit=2.679 ± 0.052, mp5=5.318 ± 0.086, crit=0.549 ± 0.033 per rating point (14 rating = 1%, 7.688 per %), spell_haste=not significant (2.247 ± 1.028)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 96.6 healing_power points (9.61 DPS) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Mooncloth Circlet (14140, -5.76 DPS, sim-verified) [crafted] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | 55.6 healing_power points (5.53 DPS) | yes | Jeweled Amulet of Cainwyn (1443, -0.45 DPS, sim-verified) [world_drop]; Animated Chain Necklace (18723, -0.65 DPS) [dungeon]; Heart of the Fiend (13960, -0.78 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -7.90 DPS, sim-verified) [dungeon] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -0.48 DPS) [rep]; Featherskin Cape (10843, -0.60 DPS) [world]; Frostweaver Cape (12968, -2.67 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 115.9 healing_power points (11.53 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Embrace of the Wind Serpent (12462, -0.97 DPS) [world]; Mooncloth Vest (14138, -4.12 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Wyrmthalak's Shackles (13958, -8.37 DPS, sim-verified) [quest] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swarmtender's Gloves (275607, -0.61 DPS) [crafted]; Magister's Gloves (16684, -0.99 DPS) [dungeon]; Hands of the Exalted Herald (12554, -6.82 DPS, sim-verified) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 87.3 healing_power points (8.69 DPS) | yes | General's Satin Cinch (17621, -1.22 DPS) [pvp]; Devout Belt (16696, -2.17 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -3.92 DPS, sim-verified) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 108.4 healing_power points (10.78 DPS) | yes | General's Satin Legguards (231634, +0.00 DPS) [vendor]; Legionnaire's Satin Legguards (227123, -0.81 DPS) [pvp]; Haunting Specter Leggings (11929, -11.59 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; General's Satin Walkers (231630, -0.04 DPS) [vendor]; Mooncloth Boots (15802, -6.28 DPS, sim-verified) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (288.4 DPS) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; The Postmaster's Seal (13392, -0.23 DPS) [dungeon]; Darkspear Signet (272068, -0.45 DPS) [vendor] |
| finger2 | Rosewine Circle (13178) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Band of Mending (22334, +0.00 DPS) [dungeon]; The Postmaster's Seal (13392, -0.55 DPS) [dungeon]; Darkspear Signet (272068, -0.77 DPS) [vendor] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.51 DPS) [quest]; Darkspear Voodoo Seal (272061, -0.53 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.98 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.51 DPS) [quest]; Second Wind (11819, -3.55 DPS, sim-verified) [dungeon] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Battle Mace (23464, +0.00 DPS) [vendor]; Resurgence Rod (17743, -4.74 DPS) [quest]; Hand of Edward the Odd (2243, -5.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 46.5 healing_power points (4.62 DPS) | yes | Jaina's Firestarter (13064, -2.92 DPS) [world_drop]; Bonecreeper Stylus (13938, -2.93 DPS) [dungeon]; Brilliant Wand (249385, -6.73 DPS, sim-verified) [crafted] |

**New at 60:** head: Devout Crown; neck: Lady Maye's Pendant; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Devout Gloves; waist: Penitent's Cinch; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Rosewine Circle; trinket1: Shard of the Splithooves; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1354, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005005031330011430-33505003030000000-00000000000000000)

Set DPS (verified): 592.7. Weights run: 18.6s. Verify run: 578.5s. 1354 eligible items had no known source.

Race sweep: undead 592.5 ± 0.4, troll 578.1 ± 0.3 (heal_score); the guide race undead is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.114, intellect=0.890 ± 0.041, spirit=0.942 ± 0.041, mp5=2.250 ± 0.043, crit=0.638 ± 0.024 per rating point (14 rating = 1%, 8.933 per %), spell_haste=4.014 ± 0.914

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Animated Chain Necklace (18723, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, -0.17 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.35 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Shoulders (14139, +0.00 DPS) [crafted]; Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp] |
| back | Drape of Recovery (272413) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19526, -1.35 DPS) [rep]; Raincaster Drape (12110, -1.99 DPS) [quest] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 88.1 healing_power points (18.37 DPS) | yes | Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robes of the Exalted (13346, -1.10 DPS) [dungeon]; Warlord's Satin Robes (231612, -4.41 DPS) [pvp] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dryad's Wrist Bindings (19595, +0.00 DPS) [rep]; Earthenweave Cuffs (254131, +0.00 DPS) [crafted] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of the Exalted Herald (12554, +0.00 DPS) [dungeon]; General's Satin Handwraps (231633, +0.00 DPS) [vendor]; Gilded Gloves (254095, -0.54 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Penitent's Cinch (272394, -1.32 DPS) [vendor]; Devout Belt (16696, -3.06 DPS) [dungeon]; Gilded Waistcord (254081, -3.11 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 59.0 healing_power points (12.31 DPS) | yes | Legionnaire's Satin Legguards (227123, +0.00 DPS) [pvp]; General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Faith Healer's Boots (22247, +0.00 DPS) [dungeon]; General's Satin Walkers (231630, +0.00 DPS) [vendor] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rosewine Circle (13178, +0.00 DPS) [dungeon]; Fordring's Seal (16058, +0.00 DPS) [quest]; Band of Mending (22334, +0.00 DPS) [dungeon] |
| trinket1 | Dormant Heart of the Mountain (249473) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, -4.17 DPS) [dungeon]; Weakness Analyzer (272438, -4.17 DPS) [vendor]; Uther's Strength (11302, -5.01 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Resurgence Rod (17743, -9.91 DPS) [quest]; Dreamstaff (249454, -12.80 DPS) [crafted]; High Warlord's Battle Mace (23464, -13.49 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 27.4 healing_power points (5.71 DPS) | yes | Brilliant Wand (249385, -2.53 DPS) [crafted]; Bonecreeper Stylus (13938, -2.67 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -3.41 DPS) [dungeon] |

**New at 60:** head: Devout Crown; shoulder: Devout Mantle; back: Drape of Recovery; chest: Truefaith Vestments; wrist: Devout Bracers; hands: Raider Handwraps; waist: Wisdom of the Timbermaw; legs: Devout Skirt; feet: Devout Sandals; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Dormant Heart of the Mountain; trinket2: Serenity Field; ranged: Torch of Light

No-known-source sample (15 of 1354, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

