# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 21.2. Weights run: 0.8s. Verify run: 0.8s. 127 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.325 ± 0.127, crit=0.376 ± 0.034, hit=1.173 ± 0.026, spell_haste=not significant (1.258 ± 0.345), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (19.5 DPS) | yes | Shadow Goggles (4373, -1.16 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.9 spell_power points (1.50 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.01 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.15 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Sanguine Cape (14376, -0.19 DPS, sim-verified) [world_drop]; Feyscale Cloak (6632, -0.26 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 11.6 spell_power points (1.03 DPS) | yes | Mystic's Wrap (14369, -0.21 DPS) [world_drop]; Mystic's Robe (14371, -0.21 DPS) [world_drop]; Gray Woolen Robe (2585, -0.52 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 8.0 spell_power points (0.71 DPS) | yes | Mindthrust Bracers (1974, -0.01 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.24 DPS) [world_drop]; Mystic's Bracelets (14366, -0.47 DPS) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (19.1 DPS) | yes | Tomb Robber's Gloves (280096, -0.00 DPS) [quest]; Serpent Gloves (5970, -0.09 DPS) [dungeon]; Blight Gloves (279877, -0.74 DPS, sim-verified) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (19.1 DPS) | yes | Novice Arcanist's Sash (253885, -0.12 DPS) [crafted]; Tarantula Silk Sash (3229, -0.24 DPS) [world]; Keller's Girdle (2911, -0.73 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (18.6 DPS) | yes | Abomination Skin Leggings (23173, -0.28 DPS, sim-verified) [dungeon]; Darkweave Breeches (12987, -0.41 DPS) [world_drop]; Filigreed Silky Leggings (253939, -0.53 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 12.3 spell_power points (1.09 DPS) | yes | Pristine Boots (253889, -0.10 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.62 DPS) [world]; Kimbra Boots (6191, -0.62 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 8.0 spell_power points (0.71 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Lorekeeper's Ring (20431, -0.26 DPS) [rep]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.7 spell_power points (0.68 DPS) | yes | Lorekeeper's Ring (20431, -0.24 DPS) [rep]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; Loop of Sacrifice (281673, -0.48 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 13.3 spell_power points (1.18 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.24 DPS) [world]; Lesser Staff of the Spire (1300, -0.47 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 254.3 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.29 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 127, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (gnome, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 53.3. Weights run: 0.8s. Verify run: 0.6s. 219 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=1.020 ± 0.088, crit=0.303 ± 0.045, hit=2.022 ± 0.044, spell_haste=not significant (0.291 ± 0.092), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 16.2 spell_power points (1.45 DPS) | yes | Shadow Hood (4323, -0.45 DPS) [crafted]; Resilient Cap (14401, -0.45 DPS) [world_drop]; Nightsky Cowl (4039, -0.66 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.1 spell_power points (1.18 DPS) | yes | Crystal Starfire Medallion (5003, -0.81 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.81 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.27 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.2 spell_power points (1.63 DPS) | yes | Fairywing Mantle (9536, -0.27 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop]; Death Speaker Mantle (6685, -0.51 DPS, sim-verified) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 8.2 spell_power points (0.73 DPS) | yes | Repairman's Cape (9605, -0.10 DPS) [quest]; Darkspear Raider's Cloak (272078, -0.13 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.18 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 22.3 spell_power points (2.00 DPS) | yes | Mechbuilder's Overalls (9508, -0.62 DPS) [dungeon]; Pristine Gown (253961, -0.73 DPS) [crafted]; Death Speaker Robes (6682, -1.07 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Nightsky Wristbands (6407, -0.26 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.26 DPS) [quest]; Glowing Magical Bracelets (13106, -1.09 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 15.2 spell_power points (1.36 DPS) | yes | Truefaith Gloves (7049, -0.64 DPS) [crafted]; Blight Gloves (279877, -0.72 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.84 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 14.1 spell_power points (1.26 DPS) | yes | Invoker's Cord (215366, -0.18 DPS) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Crimson Silk Belt (7055, -0.95 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 17.2 spell_power points (1.54 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.45 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.46 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 14.1 spell_power points (1.27 DPS) | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.63 DPS) [crafted]; Acidic Walkers (9454, -1.79 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 7.1 spell_power points (0.64 DPS) | yes | Minor Channeling Ring (1449, -0.01 DPS) [quest]; Lorekeeper's Ring (19525, -0.01 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.1 spell_power points (0.64 DPS) | yes | Lorekeeper's Ring (19525, -0.01 DPS) [rep]; Minor Channeling Ring (1449, -0.02 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (53.3 DPS) | yes | Talisman of Arathor (21119, -2.43 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 11.2 spell_power points (1.01 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Channeler's Staff (4437, -0.27 DPS) [world]; Twisted Chanter's Staff (890, -0.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 373.4 spell_power points (33.48 DPS) | yes | Starfaller (13063, -0.76 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 79.8. Weights run: 0.9s. Verify run: 0.6s. 304 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.487 ± 0.184, crit=0.512 ± 0.078, hit=3.296 ± 0.077, spell_haste=not significant (0.544 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (79.8 DPS) | yes | Thinking Cap (2624, -0.06 DPS) [world_drop]; Miner's Hat of the Deep (9429, -0.06 DPS) [dungeon]; Corpseshroud (10574, -3.18 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 15.9 spell_power points (1.54 DPS) | yes | Necklace of Calisea (1714, -0.53 DPS) [world_drop]; Triune Amulet (7722, -0.53 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.44 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 26.3 spell_power points (2.55 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.38 DPS) [dungeon]; Death Speaker Mantle (6685, -0.38 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 16.4 spell_power points (1.58 DPS) | yes | Long Silken Cloak (4326, +0.00 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.28 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -0.29 DPS) [dungeon] |
| chest | Robe of Power (7054) | Tailoring [crafted] | 31.8 spell_power points (3.08 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of the Magi (1716, -0.09 DPS) [world_drop]; Green Silk Armor (7065, -0.34 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 13.4 spell_power points (1.29 DPS) | yes | Aurora Bracers (4043, -0.07 DPS, sim-verified) [world_drop]; Mistscape Bracers (4045, -0.14 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.14 DPS) [quest] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 25.9 spell_power points (2.50 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.39 DPS) [crafted]; Town Clerk's Mittens (270029, -0.53 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.3 spell_power points (2.83 DPS) | yes | Highlander's Cloth Girdle (20098, -0.56 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.91 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -1.11 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 31.8 spell_power points (3.08 DPS) | yes | Crimson Silk Pantaloons (7062, -0.78 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.06 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.25 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.32 DPS) | yes | Acidic Walkers (9454, -0.69 DPS) [dungeon]; Gilded Slippers (254001, -0.80 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.07 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 18.9 spell_power points (1.83 DPS) | yes | Voodoo Band (1996, -0.82 DPS) [world_drop]; Mindbender Loop (5009, -0.82 DPS) [world_drop]; Black Widow Band (6199, -0.82 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 10.4 spell_power points (1.01 DPS) | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Windweaver Staff (7757) | Scarlet Monastery: Houndmaster Loksey [dungeon] | sim-verified (76.6 DPS) | yes | Staff of Jordan (873, -0.57 DPS) [world_drop]; Glimmering Staff (249392, -0.57 DPS) [crafted]; Gut Ripper (2164, -1.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 416.3 spell_power points (40.24 DPS) | yes | Umbral Wand (5216, -0.74 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.65 DPS) [dungeon]; Twisted Nether Wand (249144, -5.66 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of Power; wrist: Windchaser Cuffs; hands: Red Mageweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ogremind Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Windweaver Staff; ranged: Jaina's Firestarter

No-known-source sample (15 of 304, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (gnome, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 119.3. Weights run: 0.5s. Verify run: 0.6s. 401 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.962 ± 0.656), crit=1.607 ± 0.120, hit=6.384 ± 0.118, spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 60.0 spell_power points (6.26 DPS) | yes | Red Mageweave Headband (10033, -0.19 DPS) [crafted]; Chief Architect's Monocle (11839, -0.74 DPS) [dungeon]; Eye of Theradras (17715, -0.99 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 27.5 spell_power points (2.86 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.82 DPS) [quest]; Scorn's Icy Choker (23169, -0.91 DPS) [dungeon]; Gemshard Heart (17707, -2.84 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 48.3 spell_power points (5.03 DPS) | yes | Red Mageweave Shoulders (10029, -1.24 DPS) [crafted]; Inquisitor's Shawl (19507, -1.65 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.86 DPS, sim-verified) [vendor] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 27.5 spell_power points (2.86 DPS) | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.29 DPS) [crafted]; Big Voodoo Cloak (8216, -0.50 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 58.2 spell_power points (6.07 DPS) | yes | Robes of Insight (940, -0.96 DPS) [world_drop]; Runecloth Robe (13858, -1.45 DPS) [crafted]; Knight's Dreadweave Vest (220886, -2.14 DPS, sim-verified) [vendor] |
| wrist | Forgotten Wraps (9433) | World drop [world_drop] | sim-verified (119.3 DPS) | yes | Shizzle's Nozzle Wiper (11917, +0.00 DPS) [quest]; Bloodband Bracers (11469, -0.09 DPS) [quest]; Aristocratic Cuffs (12546, -1.47 DPS, sim-verified) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 50.2 spell_power points (5.23 DPS) | yes | Virtuous Hands (226958, -0.55 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.03 DPS) [vendor]; Red Mageweave Gloves (10018, -2.04 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 43.3 spell_power points (4.51 DPS) | yes | Deathmage Sash (10771, -0.71 DPS) [dungeon]; Satyrmane Sash (17755, -1.01 DPS) [dungeon]; Highlander's Cloth Girdle (20097, -2.26 DPS, sim-verified) [rep] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 58.0 spell_power points (6.05 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.56 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -2.14 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.56 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 89.5 spell_power points (9.33 DPS) | yes | Southsea Mojo Boots (20641, -0.99 DPS, sim-verified) [quest]; Gilded Sandals (254107, -6.34 DPS) [crafted]; Coldstone Slippers (18697, -6.46 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 63.8 spell_power points (6.65 DPS) | yes | Mindseye Circle (10634, -4.20 DPS) [dungeon]; Philanthropist's Ring (281635, -4.38 DPS) [quest]; Woodseed Hoop (17768, -4.81 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.4 spell_power points (3.07 DPS) | yes | Mindseye Circle (10634, +0.00 DPS, sim-verified) [dungeon]; Philanthropist's Ring (281635, -0.80 DPS) [quest]; Woodseed Hoop (17768, -1.23 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (117.8 DPS) | yes | Uther's Strength (11302, -0.04 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (117.8 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -1.13 DPS) [dungeon]; Spellshifter Rod (9527, -1.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.98 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -14.40 DPS, sim-verified) [quest] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Acumen Robes; wrist: Forgotten Wraps; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Brainlash; trinket1: Guardian Talisman; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 401, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (gnome, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 242.4. Weights run: 1.0s. Verify run: 0.9s. 996 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=1.188 ± 0.238, crit=1.001 ± 0.150, hit=7.692 ± 0.171, spell_haste=not significant (0.018 ± 0.293), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Circlet of Revelation (239585) | Leonid Barthalomew the Revered [vendor] | sim-verified (197.2 DPS) | yes | Field Marshal's Headdress (17602, -2.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, -2.00 DPS) [vendor]; Bloodvine Goggles (19999, -9.42 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (187.8 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Amulet of the Dawn (22657, -5.23 DPS) [quest]; Beads of Ogre Mojo (22149, -5.59 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | sim-verified (189.8 DPS) | yes | Shoulderpads of Revelation (239586, -0.13 DPS) [vendor]; Field Marshal's Satin Mantle (17604, -1.13 DPS) [vendor]; Virtuous Epaulets (226955, -2.03 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 102.4 spell_power points (11.53 DPS) | yes | Earthweave Cloak (21187, -2.87 DPS, sim-verified) [quest]; Howler's Furs (272414, -2.87 DPS) [vendor]; Stalwart Cloak (272415, -2.87 DPS) [vendor] |
| chest | Robe of Revelation (239591) | Leonid Barthalomew the Revered [vendor] | sim-verified (199.4 DPS) | yes | Field Marshal's Satin Vestments (17605, -1.09 DPS) [vendor]; Field Marshal's Satin Robe (231618, -1.09 DPS) [vendor]; Bloodvine Vest (19682, -11.62 DPS, sim-verified) [crafted] |
| wrist | Bindings of Revelation (239588) | Leonid Barthalomew the Revered [vendor] | sim-verified (194.2 DPS) | yes | Wrists of Revelation (239583, -1.87 DPS) [vendor]; Dryad's Wrist Bindings (19595, -2.18 DPS) [rep]; Rockfury Bracers (21186, -6.37 DPS, sim-verified) [quest] |
| hands | Gloves of Revelation (239584) | Leonid Barthalomew the Revered [vendor] | sim-verified (193.7 DPS) | yes | Gloves of Spell Mastery (14146, -2.47 DPS) [crafted]; Hands of Revelation (239574, -3.46 DPS) [vendor]; Dreadmist Wraps (16705, -5.90 DPS, sim-verified) [dungeon] |
| waist | Belt of Revelation (239590) | Leonid Barthalomew the Revered [vendor] | sim-verified (192.9 DPS) | yes | Belt of the Archmage (18405, -0.70 DPS) [crafted]; Girdle of Revelation (239582, -2.15 DPS) [vendor]; Knowledge of the Timbermaw (228190, -5.11 DPS, sim-verified) [vendor] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (191.1 DPS) | yes | Leggings of Revelation (239587, -0.70 DPS) [vendor]; Sentinel's Silk Leggings (237815, -2.49 DPS) [vendor]; Bloodvine Leggings (19683, -3.30 DPS, sim-verified) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 114.9 spell_power points (12.94 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.53 DPS, sim-verified) [vendor]; Argent Elite Boots (227816, -4.28 DPS) [vendor]; Sandals of Revelation (239589, -6.45 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (187.8 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.36 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.47 DPS) [vendor]; Wrath of Cenarius (21190, -2.62 DPS, sim-verified) [quest] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (187.8 DPS) | yes | Don Julio's Band (19325, -0.30 DPS) [rep]; Band of Earthen Might (21182, -0.30 DPS) [quest]; Wrath of Cenarius (21190, -2.12 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (187.8 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.90 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -2.40 DPS, sim-verified) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (187.8 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Thunderbrew's Boot Flask (744, -1.47 DPS, sim-verified) [quest]; Uther's Strength (11302, -1.80 DPS) [world_drop] |
| main_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (187.8 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS, sim-verified) [world_drop]; Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 31.0 spell_power points (3.49 DPS) | yes | Grand Marshal's Tome of Power (23452, +0.00 DPS) [vendor]; Grand Marshal's Tome of Power (234589, +0.00 DPS) [vendor]; Spirit of Aquementas (11904, -1.53 DPS, sim-verified) [quest] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (215.5 DPS) | yes | Wand of Biting Cold (19108, -13.44 DPS) [quest]; Oblivion's Touch (18761, -13.46 DPS) [dungeon]; Cold Snap (19130, -27.70 DPS, sim-verified) [world] |

**New at 60:** head: Circlet of Revelation; neck: Beads of Ogre Might; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Revelation; wrist: Bindings of Revelation; hands: Gloves of Revelation; waist: Belt of Revelation; legs: Magister's Leggings; feet: Bloodvine Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Channeler's Ring; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Persuader; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 996, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 26.0. Weights run: 0.8s. Verify run: 0.8s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.325 ± 0.127, crit=0.376 ± 0.034, hit=1.173 ± 0.026, spell_haste=not significant (1.258 ± 0.345), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (23.6 DPS) | yes | Shadow Goggles (4373, -1.16 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.9 spell_power points (1.50 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.15 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.0 spell_power points (0.53 DPS) | yes | Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Sanguine Cape (14376, -0.19 DPS, sim-verified) [world_drop]; Feyscale Cloak (6632, -0.26 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 11.6 spell_power points (1.03 DPS) | yes | Mystic's Wrap (14369, -0.21 DPS) [world_drop]; Mystic's Robe (14371, -0.21 DPS) [world_drop]; Gray Woolen Robe (2585, -0.51 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 8.0 spell_power points (0.71 DPS) | yes | Featherbead Bracers (15452, -0.12 DPS) [quest]; Mindthrust Bracers (1974, -0.22 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.24 DPS) [world_drop] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (23.2 DPS) | yes | Tomb Robber's Gloves (280096, -0.00 DPS) [quest]; Serpent Gloves (5970, -0.09 DPS) [dungeon]; Blight Gloves (279877, -0.76 DPS, sim-verified) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (23.3 DPS) | yes | Novice Arcanist's Sash (253885, -0.12 DPS) [crafted]; Tarantula Silk Sash (3229, -0.24 DPS) [world]; Keller's Girdle (2911, -0.78 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (22.8 DPS) | yes | Abomination Skin Leggings (23173, -0.30 DPS, sim-verified) [dungeon]; Darkweave Breeches (12987, -0.41 DPS) [world_drop]; Filigreed Silky Leggings (253939, -0.53 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 12.3 spell_power points (1.09 DPS) | yes | Pristine Boots (253889, -0.09 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.62 DPS) [world]; Sanguine Sandals (14374, -0.62 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 8.0 spell_power points (0.71 DPS) | yes | Loop of Sacrifice (281673, -0.01 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop]; Sludge-Stained Band (286535, -0.44 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (23.0 DPS) | yes | Volcanic Rock Ring (12053, -0.09 DPS) [world_drop]; Sludge-Stained Band (286535, -0.18 DPS) [world]; Loop of Sacrifice (281673, -0.49 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 13.3 spell_power points (1.18 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.24 DPS) [world]; Lesser Staff of the Spire (1300, -0.47 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 254.3 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.48 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 209613 Insignia of the Alliance

### Band 30 (undead, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 48.3. Weights run: 0.8s. Verify run: 0.6s. 216 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=1.020 ± 0.088, crit=0.303 ± 0.045, hit=2.022 ± 0.044, spell_haste=not significant (0.291 ± 0.092), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 16.2 spell_power points (1.45 DPS) | yes | Nightsky Cowl (4039, -0.44 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.45 DPS) [crafted]; Resilient Cap (14401, -0.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.1 spell_power points (1.18 DPS) | yes | Darkspear Warding Pendant (272075, -0.56 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.81 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.81 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 18.2 spell_power points (1.63 DPS) | yes | Death Speaker Mantle (6685, -0.03 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.27 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 8.2 spell_power points (0.73 DPS) | yes | Darkspear Raider's Cloak (272078, -0.06 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.18 DPS) [world_drop]; Soft Willow Cape (16661, -0.27 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 22.3 spell_power points (2.00 DPS) | yes | Death Speaker Robes (6682, -0.27 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.62 DPS) [dungeon]; Pristine Gown (253961, -0.73 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Nightsky Wristbands (6407, -0.26 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.26 DPS) [quest]; Glowing Magical Bracelets (13106, -0.64 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 11.1 spell_power points (1.00 DPS) | yes | Truefaith Gloves (7049, -0.27 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.31 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.36 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 14.1 spell_power points (1.26 DPS) | yes | Crimson Silk Belt (7055, -0.02 DPS, sim-verified) [crafted]; Invoker's Cord (215366, -0.18 DPS) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (48.3 DPS) | yes | Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.19 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.67 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 14.1 spell_power points (1.27 DPS) | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.63 DPS) [crafted]; Acidic Walkers (9454, -0.77 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 7.1 spell_power points (0.64 DPS) | yes | Advisor's Ring (19521, -0.01 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.1 spell_power points (0.64 DPS) | yes | Advisor's Ring (19521, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (47.6 DPS) | yes | Defiler's Talisman (21120, -1.53 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 11.2 spell_power points (1.01 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Channeler's Staff (4437, -0.27 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 373.4 spell_power points (33.48 DPS) | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 216, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 72.7. Weights run: 0.9s. Verify run: 0.6s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.487 ± 0.184, crit=0.512 ± 0.078, hit=3.296 ± 0.077, spell_haste=not significant (0.544 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (72.7 DPS) | yes | Thinking Cap (2624, -0.06 DPS) [world_drop]; Miner's Hat of the Deep (9429, -0.06 DPS) [dungeon]; Corpseshroud (10574, -2.96 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 15.9 spell_power points (1.54 DPS) | yes | Necklace of Calisea (1714, -0.53 DPS) [world_drop]; Triune Amulet (7722, -0.53 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.01 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 26.3 spell_power points (2.55 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.38 DPS) [dungeon]; Death Speaker Mantle (6685, -0.38 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 16.4 spell_power points (1.58 DPS) | yes | Guardian Cloak (5965, -0.28 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -0.29 DPS) [dungeon]; Long Silken Cloak (4326, -0.63 DPS, sim-verified) [crafted] |
| chest | Robe of Power (7054) | Tailoring [crafted] | 31.8 spell_power points (3.08 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of the Magi (1716, -0.09 DPS) [world_drop]; Green Silk Armor (7065, -0.34 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 15.9 spell_power points (1.54 DPS) | yes | Aurora Bracers (4043, -0.39 DPS) [world_drop]; Mistscape Bracers (4045, -0.39 DPS) [world_drop]; Windchaser Cuffs (14429, -0.40 DPS, sim-verified) [world_drop] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 25.9 spell_power points (2.50 DPS) | yes | Dreamweave Gloves (10019, -0.28 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.39 DPS) [crafted]; Gilded Handwraps (254021, -0.72 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 29.3 spell_power points (2.83 DPS) | yes | Defiler's Cloth Girdle (20166, -0.91 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.91 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -1.11 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 31.8 spell_power points (3.08 DPS) | yes | Crimson Silk Pantaloons (7062, -0.63 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.06 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.25 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.32 DPS) | yes | Gilded Slippers (254001, -0.29 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Boots of the Maharishi (9658, -1.03 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 18.9 spell_power points (1.83 DPS) | yes | Voodoo Band (1996, -0.82 DPS) [world_drop]; Mindbender Loop (5009, -0.82 DPS) [world_drop]; Black Widow Band (6199, -0.82 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 10.4 spell_power points (1.01 DPS) | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Windweaver Staff (7757) | Scarlet Monastery: Houndmaster Loksey [dungeon] | sim-verified (69.7 DPS) | yes | Staff of Jordan (873, -0.57 DPS) [world_drop]; Glimmering Staff (249392, -0.57 DPS) [crafted]; Gut Ripper (2164, -2.74 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 416.3 spell_power points (40.24 DPS) | yes | Umbral Wand (5216, -1.57 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.65 DPS) [dungeon]; Twisted Nether Wand (249144, -5.66 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of Power; wrist: Radiant Silver Bracers; hands: Red Mageweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ogremind Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Windweaver Staff; ranged: Jaina's Firestarter

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 105.3. Weights run: 0.5s. Verify run: 0.6s. 398 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.962 ± 0.656), crit=1.607 ± 0.120, hit=6.384 ± 0.118, spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 60.0 spell_power points (6.26 DPS) | yes | Red Mageweave Headband (10033, -0.19 DPS) [crafted]; Chief Architect's Monocle (11839, -0.74 DPS) [dungeon]; Eye of Theradras (17715, -0.95 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 27.5 spell_power points (2.86 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.82 DPS) [quest]; Scorn's Icy Choker (23169, -0.91 DPS) [dungeon]; Gemshard Heart (17707, -2.71 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 48.3 spell_power points (5.03 DPS) | yes | Red Mageweave Shoulders (10029, -1.24 DPS) [crafted]; Inquisitor's Shawl (19507, -1.65 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.23 DPS, sim-verified) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 29.7 spell_power points (3.09 DPS) | yes | Spritecaster Cape (11623, -0.40 DPS) [dungeon]; Runecloth Cloak (13860, -0.52 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.28 DPS, sim-verified) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 58.2 spell_power points (6.07 DPS) | yes | Robes of Insight (940, -0.96 DPS) [world_drop]; Runecloth Robe (13858, -1.45 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -2.46 DPS, sim-verified) [vendor] |
| wrist | Aristocratic Cuffs (12546) | Blackrock Depths: Anvilrage Overseer [dungeon] | 29.4 spell_power points (3.07 DPS) | yes | Forgotten Wraps (9433, +0.00 DPS, sim-verified) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.61 DPS) [quest]; Bloodband Bracers (11469, -0.71 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 50.2 spell_power points (5.23 DPS) | yes | Virtuous Hands (226958, -0.72 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -2.03 DPS) [vendor]; Red Mageweave Gloves (10018, -2.04 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 43.3 spell_power points (4.51 DPS) | yes | Deathmage Sash (10771, -0.71 DPS) [dungeon]; Satyrmane Sash (17755, -1.01 DPS) [dungeon]; Defiler's Cloth Girdle (20165, -1.84 DPS, sim-verified) [rep] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 58.0 spell_power points (6.05 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.24 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -2.14 DPS) [crafted]; Crimson Silk Pantaloons (7062, -2.56 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 89.5 spell_power points (9.33 DPS) | yes | Southsea Mojo Boots (20641, -1.65 DPS, sim-verified) [quest]; Gilded Sandals (254107, -6.34 DPS) [crafted]; Coldstone Slippers (18697, -6.46 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 63.8 spell_power points (6.65 DPS) | yes | Mindseye Circle (10634, -4.20 DPS) [dungeon]; Philanthropist's Ring (281635, -4.38 DPS) [quest]; Woodseed Hoop (17768, -4.81 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.4 spell_power points (3.07 DPS) | yes | Mindseye Circle (10634, +0.00 DPS, sim-verified) [dungeon]; Philanthropist's Ring (281635, -0.80 DPS) [quest]; Woodseed Hoop (17768, -1.23 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (105.3 DPS) | yes | Uther's Strength (11302, -0.43 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (105.3 DPS) | yes | Uther's Strength (11302, -0.33 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (105.3 DPS) | yes | Hammer of the Northern Wind (810, +0.00 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -1.13 DPS) [dungeon]; Spellshifter Rod (9527, -1.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.98 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -16.76 DPS, sim-verified) [quest] |

**New at 50:** head: Blood Guard's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Aristocratic Cuffs; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blackstone Ring; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 398, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 228.0. Weights run: 1.0s. Verify run: 1.0s. 993 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=1.188 ± 0.238, crit=1.001 ± 0.150, hit=7.692 ± 0.171, spell_haste=not significant (0.018 ± 0.293), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Circlet of Revelation (239585) | Leonid Barthalomew the Revered [vendor] | sim-verified (178.2 DPS) | yes | Warlord's Satin Cowl (17623, -2.00 DPS) [vendor]; Warlord's Satin Crown (231615, -2.00 DPS) [vendor]; Bloodvine Goggles (19999, -8.71 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (169.5 DPS) | yes | Blazefury Medallion (17111, +0.00 DPS, sim-verified) [world]; Amulet of the Dawn (22657, -5.23 DPS) [quest]; Beads of Ogre Mojo (22149, -5.59 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | sim-verified (171.8 DPS) | yes | Shoulderpads of Revelation (239586, -0.13 DPS) [vendor]; Warlord's Satin Mantle (17622, -1.13 DPS) [vendor]; Virtuous Epaulets (226955, -2.30 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 102.4 spell_power points (11.53 DPS) | yes | Howler's Furs (272414, -2.87 DPS) [vendor]; Stalwart Cloak (272415, -2.87 DPS) [vendor]; Earthweave Cloak (21187, -4.28 DPS, sim-verified) [quest] |
| chest | Robe of Revelation (239591) | Leonid Barthalomew the Revered [vendor] | sim-verified (180.7 DPS) | yes | Warlord's Satin Robes (17624, -1.09 DPS) [vendor]; Warlord's Satin Robes (231612, -1.09 DPS) [vendor]; Bloodvine Vest (19682, -11.21 DPS, sim-verified) [crafted] |
| wrist | Bindings of Revelation (239588) | Leonid Barthalomew the Revered [vendor] | sim-verified (175.9 DPS) | yes | Wrists of Revelation (239583, -1.87 DPS) [vendor]; Dryad's Wrist Bindings (19595, -2.18 DPS) [rep]; Rockfury Bracers (21186, -6.41 DPS, sim-verified) [quest] |
| hands | Gloves of Revelation (239584) | Leonid Barthalomew the Revered [vendor] | sim-verified (175.3 DPS) | yes | Gloves of Spell Mastery (14146, -2.47 DPS) [crafted]; Hands of Revelation (239574, -3.46 DPS) [vendor]; Dreadmist Wraps (16705, -5.79 DPS, sim-verified) [dungeon] |
| waist | Belt of Revelation (239590) | Leonid Barthalomew the Revered [vendor] | sim-verified (172.3 DPS) | yes | Belt of the Archmage (18405, -0.70 DPS) [crafted]; Girdle of Revelation (239582, -2.15 DPS) [vendor]; Knowledge of the Timbermaw (228190, -2.82 DPS, sim-verified) [vendor] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (173.5 DPS) | yes | Leggings of Revelation (239587, -0.70 DPS) [vendor]; Sentinel's Silk Leggings (237815, -2.49 DPS) [vendor]; Bloodvine Leggings (19683, -4.05 DPS, sim-verified) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 114.9 spell_power points (12.94 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -3.40 DPS, sim-verified) [vendor]; Argent Elite Boots (227816, -4.28 DPS) [vendor]; Sandals of Revelation (239589, -6.45 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (169.5 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.36 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -0.47 DPS) [vendor]; Wrath of Cenarius (21190, -2.08 DPS, sim-verified) [quest] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-verified (169.5 DPS) | yes | Don Julio's Band (19325, -0.30 DPS) [rep]; Band of Earthen Might (21182, -0.30 DPS) [quest]; Wrath of Cenarius (21190, -1.65 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (169.5 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.90 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (169.5 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.16 DPS, sim-verified) [world_drop] |
| main_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (169.5 DPS) | yes | Hand of Edward the Odd (2243, +0.00 DPS, sim-verified) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Swiftstrike Cudgel (11964, -1.58 DPS) [quest] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 31.0 spell_power points (3.49 DPS) | yes | High Warlord's Tome of Destruction (23468, +0.00 DPS) [vendor]; High Warlord's Tome of Destruction (234563, +0.00 DPS) [vendor]; Spirit of Aquementas (11904, -1.35 DPS, sim-verified) [quest] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (198.7 DPS) | yes | Wand of Biting Cold (19108, -13.44 DPS) [quest]; Oblivion's Touch (18761, -13.46 DPS) [dungeon]; Cold Snap (19130, -29.20 DPS, sim-verified) [world] |

**New at 60:** head: Circlet of Revelation; neck: Beads of Ogre Might; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Revelation; wrist: Bindings of Revelation; hands: Gloves of Revelation; waist: Belt of Revelation; legs: Magister's Leggings; feet: Bloodvine Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Channeler's Ring; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Persuader; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 993, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

