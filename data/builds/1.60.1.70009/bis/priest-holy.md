# Leveling BiS: Holy

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 30.4. Weights run: 9.8s. Verify run: 4.7s. 150 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.688 ± 0.012, spirit=1.209 ± 0.009, mp5=3.159 ± 0.024, crit=0.145 ± 0.007 per rating point (14 rating = 1%, 2.034 per %), spell_haste=not significant (0.016 ± 0.050)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, -0.02 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.50 DPS) [crafted] |
| neck | Scholarly Pendant (277203) (or Tarnished Locket (279870)) | Friend of the Library [quest] | 4.8 healing_power points (0.22 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.2 healing_power points (0.69 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.39 DPS) [crafted] |
| back | Sanguine Cape (14376) | World drop [world_drop] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Caretaker's Cape (20428, -0.02 DPS, sim-verified) [rep]; Regent's Cloak (5969, -0.03 DPS) [world]; Seer's Cape (6378, -0.04 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 25.1 healing_power points (1.15 DPS) | yes | Robe of the Moccasin (6465, +0.00 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.20 DPS) [dungeon]; Seer's Robe (2981, -0.52 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 17.9 healing_power points (0.82 DPS) | yes | Bright Bracers (3647, +0.00 DPS, sim-verified) [world_drop]; Repurposed Hair Band (281256, -0.55 DPS) [quest]; Seer's Cuffs (3645, -0.63 DPS) [dungeon] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 16.1 healing_power points (0.73 DPS) | yes | Magefist Gloves (12977, +0.00 DPS, sim-verified) [world_drop]; Bright Gloves (3066, -0.20 DPS) [world_drop]; Tomb Robber's Gloves (280096, -0.27 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.8 healing_power points (0.81 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.19 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.32 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 33.0 healing_power points (1.51 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -0.82 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.82 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 14.1 healing_power points (0.64 DPS) | yes | Kimbra Boots (6191, +0.00 DPS, sim-verified) [quest]; Bluegill Sandals (1560, -0.27 DPS) [world]; Smoldering Boots (3076, -0.31 DPS) [world] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 10.6 healing_power points (0.49 DPS) | yes | Band of Purification (12996, -0.15 DPS) [world_drop]; Lorekeeper's Ring (20431, -0.20 DPS) [rep]; Deep Fathom Ring (6463, -0.21 DPS) [dungeon] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 10.1 healing_power points (0.46 DPS) | yes | Band of Purification (12996, +0.00 DPS, sim-verified) [world_drop]; Lorekeeper's Ring (20431, -0.17 DPS) [rep]; Deep Fathom Ring (6463, -0.19 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 16.9 healing_power points (0.77 DPS) | yes | Staff of Westfall (2042, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.15 DPS) [world]; Staff of the Blessed Seer (2271, -0.22 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Flaring Baton (5326) | The Escape [quest] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Moonstone Wand (15204, +0.00 DPS) [quest]; Sable Wand (7607, -0.01 DPS, sim-verified) [quest]; Dwarven Flamestick (5241, -0.04 DPS) [quest] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Sanguine Cape; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Flaring Baton

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 47.5. Weights run: 9.9s. Verify run: 4.9s. 244 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.030, intellect=2.337 ± 0.017, spirit=1.458 ± 0.016, mp5=3.536 ± 0.029, crit=0.206 ± 0.011 per rating point (14 rating = 1%, 2.882 per %), spell_haste=not significant (0.321 ± 0.091)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.10 DPS, sim-verified) [world_drop]; Resilient Cap (14401, -0.18 DPS) [world_drop]; Shadow Hood (4323, -0.46 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Kaleidoscope Chain (13084, -0.01 DPS, sim-verified) [world_drop]; Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.11 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 33.0 healing_power points (1.59 DPS) | yes | Mantle of Honor (3560, +0.00 DPS, sim-verified) [quest]; Nightsky Mantle (4718, -0.34 DPS) [world_drop]; Death Speaker Mantle (6685, -0.35 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 23.1 healing_power points (1.11 DPS) | yes | Prelacy Cape (7004, -0.07 DPS, sim-verified) [quest]; Repairman's Cape (9605, -0.17 DPS) [quest]; Caretaker's Cape (19533, -0.20 DPS) [rep] |
| chest | Death Speaker Robes (6682) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 45.7 healing_power points (2.20 DPS) | yes | Pristine Gown (253961, +0.00 DPS, sim-verified) [crafted]; Beguiler Robes (7728, -0.29 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.51 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.3 healing_power points (1.07 DPS) | yes | Glowing Magical Bracelets (13106, +0.00 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.19 DPS) [world_drop]; Stonecloth Bindings (14416, -0.51 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan: Caverndeep Burrower [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Town Clerk's Mittens (270029, -0.01 DPS) [quest]; Gloves of Old (9395, -0.02 DPS, sim-verified) [world_drop]; Truefaith Gloves (7049, -0.19 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 20.3 healing_power points (0.98 DPS) | yes | Resilient Cord (14406, +0.00 DPS, sim-verified) [world_drop]; Dreamer's Belt (4829, -0.05 DPS) [vendor]; Wizard's Belt (4827, -0.08 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 45.6 healing_power points (2.20 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.70 DPS) [crafted]; Necromancer Leggings (2277, -0.96 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.2 healing_power points (2.03 DPS) | yes | Acidic Walkers (9454, +0.00 DPS, sim-verified) [dungeon]; Frothing Slippers (254003, -0.96 DPS) [crafted]; Fiery Slippers (254005, -0.96 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 26.6 healing_power points (1.28 DPS) | yes | Sea Giant's Toe Ring (274746, -0.08 DPS, sim-verified) [vendor]; The Queen's Jewel (13094, -0.49 DPS) [world_drop]; Darkspear Signet (272071, -0.60 DPS) [vendor] |
| finger2 | Black Widow Band (6199) | Leech Widow [world] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | The Queen's Jewel (13094, -0.00 DPS) [world_drop]; Sea Giant's Toe Ring (274746, -0.02 DPS, sim-verified) [vendor]; Darkspear Signet (272071, -0.11 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.02 DPS, sim-verified) [dungeon]; Lorekeeper's Staff (212580, -0.26 DPS) [vendor]; Glimmering Staff (249392, -0.38 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 9.3 healing_power points (0.45 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Consecrated Wand (5244, -0.08 DPS) [quest]; Flaring Baton (5326, -0.22 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Scorn's Icy Choker; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Death Speaker Robes; hands: Hotshot Pilot's Gloves; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Black Widow Band; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 244, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 58.9. Weights run: 9.9s. Verify run: 4.9s. 328 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.077, intellect=10.374 ± 0.104, spirit=3.924 ± 0.103, mp5=4.519 ± 0.121, crit=0.250 ± 0.027 per rating point (14 rating = 1%, 3.494 per %), spell_haste=not significant (0.067 ± 0.119)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman: Shadowforge Relic Hunter [dungeon] | 233.7 healing_power points (4.67 DPS) | yes | Corpseshroud (10574, -0.26 DPS) [dungeon]; Miner's Hat of the Deep (9429, -0.36 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.62 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.10 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.29 DPS, sim-verified) [quest] |
| shoulder | Windchaser Amice (14432) (or Inquisitor's Shawl (19507)) | World drop [world_drop] | 134.9 healing_power points (2.70 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.02 DPS) [dungeon]; Batwing Mantle (6697, -0.02 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackforge Cape (6424, -0.07 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.15 DPS) [vendor] |
| chest | Silksand Tunic (14417) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Silksand Wraps (14425, +0.00 DPS) [world_drop]; Red Mageweave Vest (10007, -0.26 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.59 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 93.4 healing_power points (1.87 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -0.21 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 135.5 healing_power points (2.71 DPS) | yes | Town Clerk's Mittens (270029, -0.07 DPS, sim-verified) [quest]; Red Mageweave Gloves (10018, -0.63 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.66 DPS) [dungeon] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Sutarn's Ring (13105, -0.03 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -0.08 DPS) [quest]; Deathmage Sash (10771, -0.41 DPS, sim-verified) [dungeon] |
| legs | Crimson Silk Pantaloons (7062) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Stoneweaver Leggings (9407, -0.10 DPS) [dungeon]; Aurora Pants (4044, -0.11 DPS, sim-verified) [world_drop]; Red Mageweave Pants (10009, -0.21 DPS) [crafted] |
| feet | Thoughtcast Boots (10578) | Razorfen Downs: Withered Warrior [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gilded Slippers (254001, -0.10 DPS) [crafted]; Furen's Boots (13100, -0.13 DPS, sim-verified) [world_drop]; Kodo Rustler Boots (15697, -0.21 DPS) [quest] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 100.1 healing_power points (2.00 DPS) | yes | Ogremind Ring (1993, -0.31 DPS) [world_drop]; Mindbender Loop (5009, -0.39 DPS) [world_drop]; Welken Ring (5011, -0.41 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 84.4 healing_power points (1.69 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.08 DPS) [world_drop]; Welken Ring (5011, -0.10 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windweaver Staff (7757, -0.03 DPS) [dungeon]; Gut Ripper (2164, -0.14 DPS, sim-verified) [world_drop]; Glimmering Staff (249392, -0.86 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 74.0 healing_power points (1.48 DPS) | yes | Goblin Igniter (5253, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.41 DPS) [quest]; Starfaller (13063, -0.65 DPS) [world_drop] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Windchaser Amice; back: Mantle of Lady Falther'ess; chest: Silksand Tunic; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Crimson Silk Pantaloons; feet: Thoughtcast Boots; finger2: Voodoo Band; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 328, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 67.8. Weights run: 10.1s. Verify run: 5.0s. 423 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.044, intellect=6.209 ± 0.041, spirit=0.334 ± 0.023, mp5=1.779 ± 0.072, crit=0.299 ± 0.027 per rating point (14 rating = 1%, 4.181 per %), spell_haste=not significant (-0.173 ± 0.128)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 168.7 healing_power points (4.16 DPS) | yes | Soulcatcher Halo (10630, -0.25 DPS) [dungeon]; Papal Fez (9431, -0.94 DPS) [dungeon]; Bad Mojo Mask (9470, -0.94 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 88.3 healing_power points (2.18 DPS) | yes | Gemshard Heart (17707, -0.09 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.65 DPS) [quest]; Darkspear Warding Pendant (272073, -0.80 DPS) [vendor] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Rotgrip Mantle (17732, -0.06 DPS, sim-verified) [dungeon]; Imperial Red Mantle (8250, -0.09 DPS) [world_drop]; Red Mageweave Shoulders (10029, -0.09 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 88.6 healing_power points (2.19 DPS) | yes | Imperial Red Cloak (8248, -0.08 DPS, sim-verified) [world_drop]; Mantle of Lady Falther'ess (23178, -0.59 DPS) [dungeon]; Keeper's Cloak (14665, -0.65 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 160.3 healing_power points (3.95 DPS) | yes | Hibernal Robe (8113, -0.18 DPS, sim-verified) [world_drop]; Acumen Robes (17775, -0.89 DPS) [quest]; Embrace of the Wind Serpent (12462, -1.10 DPS) [world] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 95.1 healing_power points (2.35 DPS) | yes | Shizzle's Nozzle Wiper (11917, -0.09 DPS, sim-verified) [quest]; Forgotten Wraps (9433, -0.51 DPS) [world_drop]; Imperial Red Bracers (8247, -0.66 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 158.6 healing_power points (3.91 DPS) | yes | Virtuous Mitts (226950, -0.34 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.60 DPS) [crafted]; Silkweb Gloves (11634, -1.79 DPS) [dungeon] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 118.0 healing_power points (2.91 DPS) | yes | Serenity Belt (13144, +0.00 DPS, sim-verified) [world_drop]; Imperial Red Sash (8253, -0.61 DPS) [world_drop]; Deathmage Sash (10771, -0.61 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 117.8 healing_power points (2.91 DPS) | yes | Knight's Satin Leggings (220893, -0.16 DPS, sim-verified) [vendor]; Venomshroud Leggings (14444, -0.51 DPS) [world_drop]; Imperial Red Pants (8251, -0.61 DPS) [world_drop] |
| feet | Coldstone Slippers (18697) | Scholomance: Old Treasure Chest [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Sergeant Major's Satin Boots (220895, -0.14 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -0.18 DPS) [crafted]; Highborne Footpads (14447, -0.48 DPS) [world_drop] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 94.8 healing_power points (2.34 DPS) | yes | Woodseed Hoop (17768, -0.96 DPS) [quest]; Cyclopean Band (11824, -1.01 DPS) [dungeon]; Snake Hoop (6750, -1.21 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 74.5 healing_power points (1.84 DPS) | yes | Woodseed Hoop (17768, -0.07 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.51 DPS) [dungeon]; Snake Hoop (6750, -0.71 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.58 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.59 DPS) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.02 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.03 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -0.65 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.99 DPS) [quest]; Radiant Staff (249453, -1.61 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 38.3 healing_power points (0.94 DPS) | yes | Cairnstone Sliver (9654, -0.14 DPS) [quest]; Flash Wand (5248, -0.31 DPS) [quest]; Goblin Igniter (5253, -0.31 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Kentic Amice; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Coldstone Slippers; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 208.0. Weights run: 10.5s. Verify run: 4.5s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.108, intellect=1.584 ± 0.040, spirit=8.731 ± 0.410, mp5=3.135 ± 0.048, crit=0.267 ± 0.018 per rating point (14 rating = 1%, 3.737 per %), spell_haste=not significant (-0.439 ± 0.295)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crown of Caer Darrow (13986) | The Lich, Ras Frostwhisper [quest] | 206.3 healing_power points (19.06 DPS) | yes | Virtuous Crown (226947, -0.57 DPS, sim-verified) [quest]; Earthenweave Crown (254137, -0.84 DPS) [crafted]; Lieutenant Commander's Satin Hood (227121, -1.44 DPS) [pvp] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 144.2 healing_power points (13.32 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Heart of the Fiend (13960, -0.49 DPS) [dungeon]; Lady Alizabeth's Pendant (13002, -1.22 DPS) [world_drop] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -0.48 DPS) [dungeon]; Argent Elite Shoulders (227888, -0.50 DPS, sim-verified) [vendor] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Butcher's Apron (12608, -0.19 DPS, sim-verified) [dungeon]; Frostweaver Cape (12968, -1.25 DPS) [dungeon]; Indomitable Cloak (14683, -2.20 DPS) [world_drop] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 288.9 healing_power points (26.69 DPS) | yes | Vestments of the Atal'ai Prophet (10806, -0.90 DPS, sim-verified) [dungeon]; Alanna's Embrace (13314, -5.78 DPS) [dungeon]; Mooncloth Vest (14138, -6.03 DPS) [crafted] |
| wrist | Bracers of Mending (23129) | Dire Maul: Revanchion [dungeon] | 145.8 healing_power points (13.47 DPS) | yes | Bracers of Hope (22667, -0.51 DPS) [quest]; Wyrmthalak's Shackles (13958, -0.58 DPS, sim-verified) [quest]; Virtuous Bracers (226949, -1.61 DPS) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 198.6 healing_power points (18.35 DPS) | yes | Swarmtender's Gloves (275607, -0.12 DPS, sim-verified) [crafted]; Devout Gloves (16692, -1.78 DPS) [dungeon]; Gloves of the Atal'ai Prophet (10808, -2.21 DPS) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 232.5 healing_power points (21.48 DPS) | yes | Mana-infused Cord (279267, -1.17 DPS, sim-verified) [crafted]; Marshal's Satin Sash (17609, -6.02 DPS) [pvp]; Virtuous Belt (226948, -7.48 DPS) [quest] |
| legs | Haunting Specter Leggings (11929) | Blackrock Depths: Chest of The Seven [dungeon] | 263.5 healing_power points (24.34 DPS) | yes | Devout Skirt (16694, -0.27 DPS, sim-verified) [dungeon]; Wolfshear Leggings (13206, -1.77 DPS) [dungeon]; Rainstrider Leggings (11123, -2.56 DPS) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | 221.5 healing_power points (20.46 DPS) | yes | Incandescent Mooncloth Boots (227862, -0.52 DPS, sim-verified) [vendor]; Devout Sandals (16691, -4.27 DPS) [dungeon]; Mistwalker Boots (10629, -4.77 DPS) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, -0.71 DPS, sim-verified) [dungeon]; Choking Band (11868, -3.67 DPS) [quest]; Band of the Hierophant (13096, -3.82 DPS) [world_drop] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Choking Band (11868, -0.07 DPS) [quest]; Band of the Hierophant (13096, -0.22 DPS) [world_drop]; Naglering (11669, -1.10 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Mindtap Talisman (18371, -0.11 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18469, -0.26 DPS, sim-verified) [quest] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, -0.42 DPS, sim-verified) [world_drop]; Soulkeeper (1607, -2.57 DPS) [world_drop]; Staff of Hale Magefire (13000, -3.74 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 113.0 healing_power points (10.44 DPS) | yes | Cairnstone Sliver (9654, -0.57 DPS, sim-verified) [quest]; Jaina's Firestarter (13064, -7.15 DPS) [world_drop]; Goblin Igniter (5253, -7.44 DPS) [quest] |

**New at 60:** head: Crown of Caer Darrow; neck: The Eye of Zuldazar; shoulder: Virtuous Mantle; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Bracers of Mending; hands: Virtuous Mitts; waist: Penitent's Cinch; legs: Haunting Specter Leggings; feet: Virtuous Sandals; finger1: The Postmaster's Seal; finger2: Eye of Adaegus; trinket2: Serenity Field; main_hand: Dancing Sliver; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (gnome, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 337.0. Weights run: 8.9s. Verify run: 3.8s. 1129 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.137, intellect=1.431 ± 0.022, spirit=2.274 ± 0.032, mp5=3.081 ± 0.041, crit=0.295 ± 0.015 per rating point (14 rating = 1%, 4.124 per %), spell_haste=not significant (0.053 ± 0.215)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Hood (227121, +0.00 DPS) [pvp]; Field Marshal's Satin Hood (231622, +0.00 DPS) [vendor]; Virtuous Crown (226947, -1.26 DPS, sim-verified) [quest] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 59.0 healing_power points (9.84 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Wavefront Necklace (20685, -0.30 DPS) [world]; Lady Maye's Pendant (14558, -1.51 DPS) [world_drop] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Satin Mantle (227119, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -0.81 DPS) [dungeon]; Argent Elite Shoulders (227888, -3.50 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 56.3 healing_power points (9.39 DPS) | yes | Caretaker's Cape (19530, -2.02 DPS) [rep]; Drape of Recovery (272413, -2.16 DPS) [vendor]; Frostweaver Cape (12968, -5.33 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 109.4 healing_power points (18.23 DPS) | yes | Robes of the Exalted (13346, -1.30 DPS, sim-verified) [dungeon]; Field Marshal's Satin Tunic (231624, -1.31 DPS) [vendor]; Virtuous Robe (226945, -1.88 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 61.0 healing_power points (10.17 DPS) | yes | Bracers of Mending (23129, -0.70 DPS, sim-verified) [dungeon]; Virtuous Bracers (226949, -0.86 DPS) [quest]; Marshal's Satin Bracers (17606, -1.85 DPS) [pvp] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-verified (+5.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Desert Bloom Gloves (20717, -1.56 DPS) [quest]; Devout Gloves (16692, -1.82 DPS) [dungeon]; Virtuous Mitts (226950, -5.42 DPS, sim-verified) [vendor] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 77.6 healing_power points (12.93 DPS) | yes | Virtuous Belt (226948, -0.70 DPS, sim-verified) [quest]; Wisdom of the Timbermaw (19047, -1.01 DPS) [crafted]; Marshal's Satin Sash (17609, -1.76 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 97.8 healing_power points (16.30 DPS) | yes | Knight-Captain's Satin Legguards (227125, -0.74 DPS) [pvp]; Marshal's Satin Legguards (231626, -1.38 DPS) [vendor]; Virtuous Skirt (226946, -1.60 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mooncloth Boots (15802, -2.78 DPS) [crafted]; Incandescent Mooncloth Boots (227862, -3.25 DPS, sim-verified) [vendor]; Faith Healer's Boots (22247, -3.30 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.91 DPS) [vendor]; Band of Piety (22681, -1.01 DPS) [quest]; Naglering (11669, -10.66 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.20 DPS) [vendor]; Band of Piety (22681, -0.29 DPS) [quest]; Naglering (11669, -9.26 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (+9.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -0.55 DPS) [vendor]; Mindtap Talisman (18371, -1.91 DPS) [dungeon]; Briarwood Reed (12930, -2.72 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.88 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -2.05 DPS) [dungeon]; Briarwood Reed (12930, -2.87 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -1.21 DPS) [dungeon]; Staff of Hale Magefire (13000, -1.25 DPS) [world_drop]; Hand of Edward the Odd (2243, -9.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 42.0 healing_power points (7.00 DPS) | yes | Cairnstone Sliver (9654, -3.18 DPS, sim-verified) [quest]; Sparkling Crystal Wand (20672, -4.02 DPS) [world]; Bonecreeper Stylus (13938, -4.22 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: The Eye of Zuldazar; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Hands of the Exalted Herald; waist: Penitent's Cinch; legs: Devout Skirt; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Emerald Flame Ring; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1129, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (troll, 000000000000000000-03503000000000000-000000000000000000)

Set DPS (verified): 29.6. Weights run: 9.8s. Verify run: 5.2s. 140 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.013, intellect=1.688 ± 0.012, spirit=1.209 ± 0.009, mp5=3.159 ± 0.024, crit=0.145 ± 0.007 per rating point (14 rating = 1%, 2.034 per %), spell_haste=not significant (0.016 ± 0.050)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Goggles (4373) | Engineering [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, -0.00 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.50 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 4.8 healing_power points (0.22 DPS) | yes | Roadwatcher's Confidence (281265, -0.06 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 15.2 healing_power points (0.69 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -0.39 DPS) [crafted] |
| back | Battle Healer's Cloak (20427) | Warsong Outriders [rep] | 11.4 healing_power points (0.52 DPS) | yes | Sanguine Cape (14376, +0.00 DPS, sim-verified) [world_drop]; Regent's Cloak (5969, -0.25 DPS) [world]; Traveler's Shawl (277289, -0.25 DPS) [quest] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 25.1 healing_power points (1.15 DPS) | yes | Robe of the Moccasin (6465, +0.00 DPS, sim-verified) [dungeon]; Corsair's Overshirt (5202, -0.20 DPS) [dungeon]; Seer's Robe (2981, -0.52 DPS) [world_drop] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 17.9 healing_power points (0.82 DPS) | yes | Tabitha's Cuffs (251486, +0.00 DPS, sim-verified) [quest]; Featherbead Bracers (15452, -0.43 DPS) [quest]; Crystalline Cuffs (14148, -0.50 DPS) [dungeon] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 17.9 healing_power points (0.82 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Magefist Gloves (12977, -0.21 DPS) [world_drop]; Bright Gloves (3066, -0.29 DPS) [world_drop] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 17.8 healing_power points (0.81 DPS) | yes | Novice Ardent's Sash (253887, +0.00 DPS, sim-verified) [crafted]; Keller's Girdle (2911, -0.19 DPS) [world_drop]; Tarantula Silk Sash (3229, -0.32 DPS) [world] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 33.0 healing_power points (1.51 DPS) | yes | Darkweave Breeches (12987, +0.00 DPS, sim-verified) [world_drop]; Filigreed Silky Leggings (253939, -0.82 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.82 DPS) [crafted] |
| feet | Pristine Boots (253889) | Tailoring [crafted] | 14.1 healing_power points (0.64 DPS) | yes | Bluegill Sandals (1560, +0.00 DPS, sim-verified) [world]; Smoldering Boots (3076, -0.31 DPS) [world]; Sanguine Sandals (14374, -0.33 DPS) [world_drop] |
| finger1 | Black Pearl Ring (6332) | Lady Vespira [world] | 10.6 healing_power points (0.49 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; Band of Purification (12996, -0.15 DPS) [world_drop]; Advisor's Ring (20426, -0.20 DPS) [rep] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 10.1 healing_power points (0.46 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Band of Purification (12996, -0.13 DPS) [world_drop]; Advisor's Ring (20426, -0.17 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) | The Wrath of Rath'mael [quest] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Staff of Orgrimmar (15444, -0.01 DPS, sim-verified) [quest]; Advisor's Gnarled Staff (20425, -0.03 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Decay (5252) (or Flaring Baton (5326)) | Beren's Peril [quest] | 3.4 healing_power points (0.15 DPS) | yes | Flaring Baton (5326, +0.00 DPS) [quest]; Wisesight Wand (286750, -0.08 DPS) [world] |

**New at 20:** head: Shadow Goggles; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Battle Healer's Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Pristine Boots; finger1: Black Pearl Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff; ranged: Wand of Decay

No-known-source sample (15 of 140, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 209613 Insignia of the Alliance; 241089 Scarlet Dagger; 248007 Militia Shortblade

### Band 30 (troll, 000000000000000000-03505003030110000-000000000000000000)

Set DPS (verified): 47.4. Weights run: 9.9s. Verify run: 5.5s. 231 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.030, intellect=2.337 ± 0.017, spirit=1.458 ± 0.016, mp5=3.536 ± 0.029, crit=0.206 ± 0.011 per rating point (14 rating = 1%, 2.882 per %), spell_haste=not significant (0.321 ± 0.091)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightsky Cowl (4039) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.10 DPS, sim-verified) [world_drop]; Resilient Cap (14401, -0.18 DPS) [world_drop]; Shadow Hood (4323, -0.46 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.2 healing_power points (0.73 DPS) | yes | Scorn's Icy Choker (23169, +0.00 DPS, sim-verified) [dungeon]; Crystal Starfire Medallion (5003, -0.07 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.17 DPS) [vendor] |
| shoulder | Batwing Mantle (6697) | Razorfen Kraul: Blind Hunter [dungeon] | 33.0 healing_power points (1.59 DPS) | yes | Nightsky Mantle (4718, +0.00 DPS, sim-verified) [world_drop]; Death Speaker Mantle (6685, -0.35 DPS) [dungeon]; Mantle of Woe (7750, -0.52 DPS) [quest] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 23.1 healing_power points (1.11 DPS) | yes | Battle Healer's Cloak (19529, +0.00 DPS, sim-verified) [rep]; Cloak of Rot (4462, -0.21 DPS) [world]; Glowing Thresher Cape (6901, -0.21 DPS) [dungeon] |
| chest | Pristine Gown (253961) | Tailoring [crafted] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Robes (6682, -0.01 DPS, sim-verified) [dungeon]; Beguiler Robes (7728, -0.12 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.34 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 22.3 healing_power points (1.07 DPS) | yes | Glowing Magical Bracelets (13106, -0.17 DPS) [world_drop]; Nightsky Wristbands (6407, -0.19 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.40 DPS) [quest] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan: Caverndeep Burrower [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of Old (9395, -0.01 DPS, sim-verified) [world_drop]; Blight Gloves (279877, -0.11 DPS) [quest]; Truefaith Gloves (7049, -0.19 DPS) [crafted] |
| waist | Lilac Sash (6780) | Centaur Bounty [quest] | 25.4 healing_power points (1.22 DPS) | yes | Pristine Sash (253925, +0.00 DPS, sim-verified) [crafted]; Resilient Cord (14406, -0.27 DPS) [world_drop]; Dreamer's Belt (4829, -0.30 DPS) [vendor] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 45.6 healing_power points (2.20 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Earthen Leggings (253999, -0.70 DPS) [crafted]; Necromancer Leggings (2277, -0.96 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 42.2 healing_power points (2.03 DPS) | yes | Acidic Walkers (9454, +0.00 DPS, sim-verified) [dungeon]; Frothing Slippers (254003, -0.96 DPS) [crafted]; Fiery Slippers (254005, -0.96 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 26.6 healing_power points (1.28 DPS) | yes | Sea Giant's Toe Ring (274746, -0.06 DPS, sim-verified) [vendor]; The Queen's Jewel (13094, -0.49 DPS) [world_drop]; Darkspear Signet (272071, -0.60 DPS) [vendor] |
| finger2 | Black Widow Band (6199) | Leech Widow [world] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | The Queen's Jewel (13094, -0.00 DPS) [world_drop]; Sea Giant's Toe Ring (274746, -0.02 DPS, sim-verified) [vendor]; Darkspear Signet (272071, -0.11 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Wind Spirit Staff (6689) | Razorfen Kraul: Earthcaller Halmgar [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Death Speaker Scepter (2816, -0.02 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (19569, -0.15 DPS) [pvp]; Lorekeeper's Staff (212580, -0.26 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) (or Lesser Mystic Wand (11289)) | World drop [world_drop] | 9.3 healing_power points (0.45 DPS) | yes | Lesser Mystic Wand (11289, +0.00 DPS) [crafted]; Wand of Decay (5252, -0.22 DPS) [quest]; Flaring Baton (5326, -0.22 DPS) [quest] |

**New at 30:** head: Nightsky Cowl; neck: Kaleidoscope Chain; shoulder: Batwing Mantle; back: Darkspear Raider's Cloak; chest: Pristine Gown; hands: Hotshot Pilot's Gloves; waist: Lilac Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Black Widow Band; trinket1: Darkspear Voodoo Seal; main_hand: Wind Spirit Staff; ranged: Starfaller

No-known-source sample (15 of 231, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (troll, 000000000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 58.0. Weights run: 9.9s. Verify run: 5.4s. 311 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.077, intellect=10.374 ± 0.104, spirit=3.924 ± 0.103, mp5=4.519 ± 0.121, crit=0.250 ± 0.027 per rating point (14 rating = 1%, 3.494 per %), spell_haste=not significant (0.067 ± 0.119)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs: Withered Warrior [dungeon] | sim-verified (1.7 DPS) | yes | Papal Fez (9431, -0.04 DPS, sim-verified) [dungeon]; Miner's Hat of the Deep (9429, -0.10 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -0.36 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Necklace of Calisea (1714, +0.00 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.10 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.14 DPS, sim-verified) [quest] |
| shoulder | Windchaser Amice (14432) (or Inquisitor's Shawl (19507)) | World drop [world_drop] | 134.9 healing_power points (2.70 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS) [dungeon]; Mistscape Mantle (4734, -0.02 DPS) [dungeon]; Batwing Mantle (6697, -0.02 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackforge Cape (6424, -0.07 DPS) [dungeon]; Darkspear Raider's Cloak (272077, -0.09 DPS, sim-verified) [vendor] |
| chest | Silksand Tunic (14417) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Silksand Wraps (14425, +0.00 DPS) [world_drop]; Red Mageweave Vest (10007, -0.30 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.59 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 93.4 healing_power points (1.87 DPS) | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [dungeon]; Radiant Silver Bracers (4545, -0.21 DPS) [quest]; Enchanted Stonecloth Bracers (4979, -0.21 DPS) [quest] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | 135.5 healing_power points (2.71 DPS) | yes | Red Mageweave Gloves (10018, -0.07 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.66 DPS) [dungeon]; Gloves of Old (9395, -0.73 DPS) [world_drop] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Sutarn's Ring (13105, -0.03 DPS) [world_drop]; Razzeric's Customized Seatbelt (6726, -0.08 DPS) [quest]; Deathmage Sash (10771, -0.27 DPS, sim-verified) [dungeon] |
| legs | Crimson Silk Pantaloons (7062) | Tailoring [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Aurora Pants (4044, -0.09 DPS, sim-verified) [world_drop]; Stoneweaver Leggings (9407, -0.10 DPS) [dungeon]; Red Mageweave Pants (10009, -0.21 DPS) [crafted] |
| feet | Boots of the Maharishi (9658) | Gordunni Cobalt [quest] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Thoughtcast Boots (10578, -0.07 DPS) [dungeon]; Furen's Boots (13100, -0.09 DPS, sim-verified) [world_drop]; Gilded Slippers (254001, -0.17 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 100.1 healing_power points (2.00 DPS) | yes | Ogremind Ring (1993, -0.31 DPS) [world_drop]; Mindbender Loop (5009, -0.39 DPS) [world_drop]; Welken Ring (5011, -0.41 DPS) [world_drop] |
| finger2 | Voodoo Band (1996) (or Ogremind Ring (1993)) | Bloodscalp Witch Doctor [world] | 84.4 healing_power points (1.69 DPS) | yes | Ogremind Ring (1993, +0.00 DPS) [world_drop]; Mindbender Loop (5009, -0.08 DPS) [world_drop]; Welken Ring (5011, -0.10 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windweaver Staff (7757, -0.03 DPS) [dungeon]; Gut Ripper (2164, -0.09 DPS, sim-verified) [world_drop]; Advisor's Gnarled Staff (19568, -0.74 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 74.0 healing_power points (1.48 DPS) | yes | Goblin Igniter (5253, +0.00 DPS, sim-verified) [quest]; Flash Wand (5248, -0.41 DPS) [quest]; Starfaller (13063, -0.65 DPS) [world_drop] |

**New at 40:** head: Corpseshroud; neck: Triune Amulet; shoulder: Windchaser Amice; back: Mantle of Lady Falther'ess; chest: Silksand Tunic; wrist: Windchaser Cuffs; hands: Stormcloth Gloves; waist: Gilded Cord; legs: Crimson Silk Pantaloons; feet: Boots of the Maharishi; finger2: Voodoo Band; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 311, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (troll, 025003000000000000-03505003030121431-000000000000000000)

Set DPS (verified): 63.2. Weights run: 10.1s. Verify run: 5.6s. 403 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.044, intellect=6.209 ± 0.041, spirit=0.334 ± 0.023, mp5=1.779 ± 0.072, crit=0.299 ± 0.027 per rating point (14 rating = 1%, 4.181 per %), spell_haste=not significant (-0.173 ± 0.128)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chief Architect's Monocle (11839) | Blackrock Depths: Fineous Darkvire [dungeon] | 168.7 healing_power points (4.16 DPS) | yes | Soulcatcher Halo (10630, -0.25 DPS) [dungeon]; Papal Fez (9431, -0.94 DPS) [dungeon]; Bad Mojo Mask (9470, -0.94 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 88.3 healing_power points (2.18 DPS) | yes | Gemshard Heart (17707, -0.08 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.65 DPS) [quest]; Darkspear Warding Pendant (272073, -0.80 DPS) [vendor] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Rotgrip Mantle (17732, -0.03 DPS, sim-verified) [dungeon]; Imperial Red Mantle (8250, -0.09 DPS) [world_drop]; Red Mageweave Shoulders (10029, -0.09 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 88.6 healing_power points (2.19 DPS) | yes | Imperial Red Cloak (8248, -0.07 DPS, sim-verified) [world_drop]; Mantle of Lady Falther'ess (23178, -0.59 DPS) [dungeon]; Keeper's Cloak (14665, -0.65 DPS) [world_drop] |
| chest | Robes of Insight (940) | World drop [world_drop] | 160.3 healing_power points (3.95 DPS) | yes | Hibernal Robe (8113, -0.14 DPS, sim-verified) [world_drop]; Acumen Robes (17775, -0.89 DPS) [quest]; Embrace of the Wind Serpent (12462, -1.10 DPS) [world] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 95.1 healing_power points (2.35 DPS) | yes | Shizzle's Nozzle Wiper (11917, -0.07 DPS, sim-verified) [quest]; Forgotten Wraps (9433, -0.51 DPS) [world_drop]; Imperial Red Bracers (8247, -0.66 DPS) [world_drop] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 158.6 healing_power points (3.91 DPS) | yes | Virtuous Mitts (226950, -0.29 DPS, sim-verified) [vendor]; Gilded Gloves (254095, -1.60 DPS) [crafted]; Greenleaf Handwraps (19116, -1.68 DPS) [quest] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 118.0 healing_power points (2.91 DPS) | yes | Serenity Belt (13144, -0.05 DPS, sim-verified) [world_drop]; Imperial Red Sash (8253, -0.61 DPS) [world_drop]; Deathmage Sash (10771, -0.61 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 117.8 healing_power points (2.91 DPS) | yes | Stone Guard's Satin Leggings (220902, -0.13 DPS, sim-verified) [vendor]; Venomshroud Leggings (14444, -0.51 DPS) [world_drop]; Imperial Red Pants (8251, -0.61 DPS) [world_drop] |
| feet | Coldstone Slippers (18697) | Scholomance: Old Treasure Chest [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | First Sergeant's Satin Boots (220900, -0.12 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -0.18 DPS) [crafted]; Highborne Footpads (14447, -0.48 DPS) [world_drop] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 94.8 healing_power points (2.34 DPS) | yes | Woodseed Hoop (17768, -0.96 DPS) [quest]; Cyclopean Band (11824, -1.01 DPS) [dungeon]; Snake Hoop (6750, -1.21 DPS) [quest] |
| finger2 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 74.5 healing_power points (1.84 DPS) | yes | Woodseed Hoop (17768, -0.08 DPS, sim-verified) [quest]; Cyclopean Band (11824, -0.51 DPS) [dungeon]; Snake Hoop (6750, -0.71 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -0.21 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.58 DPS) [quest]; Alchemists' Stone (13503, -0.63 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.02 DPS) [quest]; Alchemists' Stone (13503, -0.07 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -0.55 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.99 DPS) [quest]; Radiant Staff (249453, -1.61 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 38.3 healing_power points (0.94 DPS) | yes | Nature's Breath (19118, -0.02 DPS) [quest]; Flash Wand (5248, -0.31 DPS) [quest]; Goblin Igniter (5253, -0.31 DPS) [quest] |

**New at 50:** head: Chief Architect's Monocle; neck: Horizon Choker; shoulder: Kentic Amice; back: Darkspear Raider's Cloak; chest: Robes of Insight; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Kilt of the Atal'ai Prophet; feet: Coldstone Slippers; finger1: Brainlash; finger2: Mindseye Circle; trinket1: Darkspear Voodoo Seal; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 202.5. Weights run: 10.5s. Verify run: 4.9s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.108, intellect=1.584 ± 0.040, spirit=8.731 ± 0.410, mp5=3.135 ± 0.048, crit=0.267 ± 0.018 per rating point (14 rating = 1%, 3.737 per %), spell_haste=not significant (-0.439 ± 0.295)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crown of Caer Darrow (13986) | The Lich, Ras Frostwhisper [quest] | 206.3 healing_power points (19.06 DPS) | yes | Earthenweave Crown (254137, -0.84 DPS) [crafted]; Champion's Satin Hood (227118, -1.44 DPS) [pvp]; Virtuous Crown (226947, -1.66 DPS, sim-verified) [quest] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 144.2 healing_power points (13.32 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Heart of the Fiend (13960, -0.49 DPS) [dungeon]; Lady Alizabeth's Pendant (13002, -1.22 DPS) [world_drop] |
| shoulder | Argent Elite Shoulders (227888) | Argent Quartermaster Hasana [vendor] | 165.1 healing_power points (15.26 DPS) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Virtuous Mantle (226951, -0.55 DPS, sim-verified) [quest]; Mantle of the Scarlet Crusade (22405, -2.12 DPS) [dungeon] |
| back | Featherskin Cape (10843) | Avatar of Hakkar [world] | sim-verified (15.5 DPS) | yes | Butcher's Apron (12608, -0.20 DPS, sim-verified) [dungeon]; Frostweaver Cape (12968, -1.25 DPS) [dungeon]; Cloak of Blight (6832, -2.20 DPS) [quest] |
| chest | Embrace of the Wind Serpent (12462) | Avatar of Hakkar [world] | 288.9 healing_power points (26.69 DPS) | yes | Vestments of the Atal'ai Prophet (10806, -1.98 DPS, sim-verified) [dungeon]; Alanna's Embrace (13314, -5.78 DPS) [dungeon]; Mooncloth Vest (14138, -6.03 DPS) [crafted] |
| wrist | Bracers of Mending (23129) | Dire Maul: Revanchion [dungeon] | 145.8 healing_power points (13.47 DPS) | yes | Wyrmthalak's Shackles (13958, -0.45 DPS, sim-verified) [quest]; Bracers of Hope (22667, -0.51 DPS) [quest]; Virtuous Bracers (226949, -1.61 DPS) [quest] |
| hands | Virtuous Mitts (226950) | Mokvar [vendor] | 198.6 healing_power points (18.35 DPS) | yes | Swarmtender's Gloves (275607, -1.59 DPS, sim-verified) [crafted]; Devout Gloves (16692, -1.78 DPS) [dungeon]; Gloves of the Atal'ai Prophet (10808, -2.21 DPS) [dungeon] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 232.5 healing_power points (21.48 DPS) | yes | Mana-infused Cord (279267, -2.24 DPS, sim-verified) [crafted]; General's Satin Cinch (17621, -6.02 DPS) [pvp]; Virtuous Belt (226948, -7.48 DPS) [quest] |
| legs | Haunting Specter Leggings (11929) | Blackrock Depths: Chest of The Seven [dungeon] | 263.5 healing_power points (24.34 DPS) | yes | Devout Skirt (16694, +0.00 DPS, sim-verified) [dungeon]; Wolfshear Leggings (13206, -1.77 DPS) [dungeon]; Rainstrider Leggings (11123, -2.56 DPS) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | 221.5 healing_power points (20.46 DPS) | yes | Incandescent Mooncloth Boots (227862, -1.69 DPS, sim-verified) [vendor]; Devout Sandals (16691, -4.27 DPS) [dungeon]; Mistwalker Boots (10629, -4.77 DPS) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, -2.32 DPS, sim-verified) [dungeon]; Band of the Hierophant (13096, -3.82 DPS) [world_drop]; Emerald Flame Ring (18395, -4.56 DPS) [dungeon] |
| finger2 | Eye of Adaegus (5266) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of the Hierophant (13096, -0.22 DPS) [world_drop]; Emerald Flame Ring (18395, -0.96 DPS) [dungeon]; Naglering (11669, -2.15 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Mindtap Talisman (18371, -0.58 DPS, sim-verified) [dungeon] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Briarwood Reed (12930, -0.52 DPS, sim-verified) [dungeon] |
| main_hand | Dancing Sliver (15854) | Dawn's Gambit [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Edward the Odd (2243, -2.21 DPS, sim-verified) [world_drop]; Soulkeeper (1607, -2.57 DPS) [world_drop]; Staff of Hale Magefire (13000, -3.74 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 113.0 healing_power points (10.44 DPS) | yes | Chillnail Splinter (10704, -2.06 DPS, sim-verified) [quest]; Jaina's Firestarter (13064, -7.15 DPS) [world_drop]; Eyepoker (6797, -7.22 DPS) [quest] |

**New at 60:** head: Crown of Caer Darrow; neck: The Eye of Zuldazar; shoulder: Argent Elite Shoulders; back: Featherskin Cape; chest: Embrace of the Wind Serpent; wrist: Bracers of Mending; hands: Virtuous Mitts; waist: Penitent's Cinch; legs: Haunting Specter Leggings; feet: Virtuous Sandals; finger1: The Postmaster's Seal; finger2: Eye of Adaegus; trinket1: Royal Seal of Eldre'Thalas; trinket2: Serenity Field; main_hand: Dancing Sliver; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (troll, 025003031303000000-03505003030121431-000000000000000000)

Set DPS (verified): 327.5. Weights run: 8.9s. Verify run: 4.2s. 1120 eligible items had no known source.

Stat weights (normalized to healing_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): healing_power=1.000 ± 0.137, intellect=1.431 ± 0.022, spirit=2.274 ± 0.032, mp5=3.081 ± 0.041, crit=0.295 ± 0.015 per rating point (14 rating = 1%, 4.124 per %), spell_haste=not significant (0.053 ± 0.215)

| Slot | Item | Source | Score (healing_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Hood (227118, +0.00 DPS) [pvp]; Warlord's Satin Hood (231635, +0.00 DPS) [vendor]; Virtuous Crown (226947, -1.32 DPS, sim-verified) [quest] |
| neck | The Eye of Zuldazar (19593) (or The All-Seeing Eye of Zuldazar (19594)) | The Eye of Zuldazar [quest] | 59.0 healing_power points (9.84 DPS) | yes | The All-Seeing Eye of Zuldazar (19594, +0.00 DPS) [quest]; Wavefront Necklace (20685, -0.30 DPS) [world]; Lady Maye's Pendant (14558, -1.51 DPS) [world_drop] |
| shoulder | Virtuous Mantle (226951) | Anthion's Parting Words [quest] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Satin Mantle (227120, +0.00 DPS) [pvp]; Mantle of the Scarlet Crusade (22405, -0.81 DPS) [dungeon]; Argent Elite Shoulders (227888, -1.94 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | 56.3 healing_power points (9.39 DPS) | yes | Battle Healer's Cloak (19526, -2.02 DPS) [rep]; Drape of Recovery (272413, -2.16 DPS) [vendor]; Frostweaver Cape (12968, -3.25 DPS, sim-verified) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 109.4 healing_power points (18.23 DPS) | yes | Robes of the Exalted (13346, -1.21 DPS, sim-verified) [dungeon]; Warlord's Satin Tunic (231632, -1.31 DPS) [vendor]; Virtuous Robe (226945, -1.88 DPS) [quest] |
| wrist | Bracers of Hope (22667) | Epic Armaments of Battle - Friend of the Dawn [quest] | 61.0 healing_power points (10.17 DPS) | yes | Bracers of Mending (23129, -0.10 DPS) [dungeon]; Virtuous Bracers (226949, -0.86 DPS) [quest]; General's Satin Bracers (17619, -1.85 DPS) [pvp] |
| hands | Hands of the Exalted Herald (12554) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Desert Bloom Gloves (20717, -1.56 DPS) [quest]; Devout Gloves (16692, -1.82 DPS) [dungeon]; Virtuous Mitts (226950, -4.43 DPS, sim-verified) [vendor] |
| waist | Penitent's Cinch (272394) | Pix Xizzix [vendor] | 77.6 healing_power points (12.93 DPS) | yes | Wisdom of the Timbermaw (19047, -1.01 DPS) [crafted]; Virtuous Belt (226948, -1.38 DPS, sim-verified) [quest]; General's Satin Cinch (17621, -1.76 DPS) [pvp] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | 97.8 healing_power points (16.30 DPS) | yes | Legionnaire's Satin Legguards (227123, -0.74 DPS) [pvp]; General's Satin Legguards (231634, -1.38 DPS) [vendor]; Virtuous Skirt (226946, -1.42 DPS, sim-verified) [quest] |
| feet | Virtuous Sandals (226952) | Anthion's Parting Words [quest] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Incandescent Mooncloth Boots (227862, -2.30 DPS, sim-verified) [vendor]; Mooncloth Boots (15802, -2.78 DPS) [crafted]; Faith Healer's Boots (22247, -3.30 DPS) [dungeon] |
| finger1 | Band of Mending (22334) | Stratholme: Balnazzar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.91 DPS) [vendor]; Band of Piety (22681, -1.01 DPS) [quest]; Naglering (11669, -8.85 DPS, sim-verified) [dungeon] |
| finger2 | Emerald Flame Ring (18395) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.20 DPS) [vendor]; Band of Piety (22681, -0.29 DPS) [quest]; Naglering (11669, -7.30 DPS, sim-verified) [dungeon] |
| trinket1 | Royal Seal of Eldre'Thalas (18469) | Holy Bologna: What the Light Won't Tell You [quest] | sim-verified (+6.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -0.55 DPS) [vendor]; Mindtap Talisman (18371, -1.91 DPS) [dungeon]; Briarwood Reed (12930, -2.72 DPS) [dungeon] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -0.27 DPS, sim-verified) [vendor]; Mindtap Talisman (18371, -2.05 DPS) [dungeon]; Briarwood Reed (12930, -2.87 DPS) [dungeon] |
| main_hand | Redemption (22406) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Metanoia (22394, -1.21 DPS) [dungeon]; Staff of Hale Magefire (13000, -1.25 DPS) [world_drop]; Hand of Edward the Odd (2243, -7.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 42.0 healing_power points (7.00 DPS) | yes | Sparkling Crystal Wand (20672, -1.79 DPS, sim-verified) [world]; Bonecreeper Stylus (13938, -4.22 DPS) [dungeon]; Oblivion's Touch (18761, -4.38 DPS) [dungeon] |

**New at 60:** head: Devout Crown; neck: The Eye of Zuldazar; shoulder: Virtuous Mantle; back: Hide of the Wild; chest: Truefaith Vestments; wrist: Bracers of Hope; hands: Hands of the Exalted Herald; waist: Penitent's Cinch; legs: Devout Skirt; feet: Virtuous Sandals; finger1: Band of Mending; finger2: Emerald Flame Ring; trinket1: Royal Seal of Eldre'Thalas; trinket2: Darkspear Voodoo Seal; main_hand: Redemption; ranged: Torch of Light

No-known-source sample (15 of 1120, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

