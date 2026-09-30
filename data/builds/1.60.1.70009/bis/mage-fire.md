# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 32.1. Weights run: 0.8s. Verify run: 0.8s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.495 ± 0.285, crit=1.741 ± 0.088, hit=3.693 ± 0.034, spell_haste=not significant (0.159 ± 0.399), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (31.8 DPS) | yes | Flying Tiger Goggles (4368, -0.45 DPS) [crafted]; Lucky Fishing Hat (19972, -0.45 DPS) [quest]; Shadow Goggles (4373, -0.77 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 18.5 | yes | Reinforced Woolen Shoulders (4315, -0.45 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.09 DPS) [crafted]; Slime-encrusted Pads (6461, -1.39 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.5 | yes | Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Feyscale Cloak (6632, -0.26 DPS) [dungeon]; Sanguine Cape (14376, -0.33 DPS, sim-verified) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 12.5 | yes | Mystic's Wrap (14369, -0.15 DPS) [world_drop]; Mystic's Robe (14371, -0.15 DPS) [world_drop]; Gray Woolen Robe (2585, -0.50 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 9.0 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.23 DPS) [world_drop]; Mystic's Bracelets (14366, -0.45 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 10.5 | yes | Pristine Gloves (253913, -0.15 DPS) [crafted]; Tomb Robber's Gloves (280096, -0.17 DPS, sim-verified) [quest]; Adept's Gloves (4768, -0.23 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (31.5 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Tarantula Silk Sash (3229, -0.19 DPS) [world]; Keller's Girdle (2911, -0.53 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 21.0 | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.79 DPS) [world_drop]; Filigreed Silky Leggings (253939, -0.90 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 13.0 | yes | Pristine Boots (253889, -0.22 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.53 DPS) [world]; Kimbra Boots (6191, -0.53 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 9.0 | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Lorekeeper's Ring (20431, -0.30 DPS) [rep]; Volcanic Rock Ring (12053, -0.34 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 8.0 | yes | Lorekeeper's Ring (20431, -0.23 DPS) [rep]; Volcanic Rock Ring (12053, -0.26 DPS) [world_drop]; Loop of Sacrifice (281673, -0.42 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 14.9 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.23 DPS) [world]; Lesser Staff of the Spire (1300, -0.45 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 298.7 | yes | Skycaller (12984, -0.66 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Blight Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff; 20440 Protector's Sword

### Band 30 (gnome, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 60.3. Weights run: 0.8s. Verify run: 0.8s. 216 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.956 ± 0.566), crit=3.069 ± 0.225, hit=4.265 ± 0.054, spell_haste=not significant (0.233 ± 0.614), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.6 | yes | Holy Shroud (2721, -0.44 DPS) [world_drop]; Nightsky Cowl (4039, -0.47 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 | yes | Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.32 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.6 | yes | Death Speaker Mantle (6685, -0.20 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.7 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.08 DPS) [quest]; Resilient Cape (14400, -0.18 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.4 | yes | Pristine Gown (253961, -0.74 DPS) [crafted]; Death Speaker Robes (6682, -0.75 DPS, sim-verified) [dungeon]; Tree Bark Jacket (1486, -0.81 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.31 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.31 DPS) [quest]; Glowing Magical Bracelets (13106, -1.01 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 14.5 | yes | Truefaith Gloves (7049, -0.63 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.66 DPS) [dungeon]; Serpent Gloves (5970, -0.72 DPS) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.9 | yes | Crimson Silk Belt (7055, -0.15 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.7 | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -0.45 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.47 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.7 | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -0.80 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Black Widow Band (6199, -0.03 DPS) [world]; Snake Hoop (6750, -0.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 | yes | Black Widow Band (6199, +0.00 DPS, sim-verified) [world]; Snake Hoop (6750, -0.02 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (59.2 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Talisman of Arathor (21119, -2.33 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.5 | yes | Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Channeler's Staff (4437, -0.27 DPS) [world]; Twisted Chanter's Staff (890, -0.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 349.6 | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 216, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 102.3. Weights run: 0.9s. Verify run: 0.8s. 293 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (0.249 ± 0.763), crit=2.900 ± 0.288, hit=4.986 ± 0.077, spell_haste=6.506 ± 0.958, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (101.4 DPS) | yes | Holy Shroud (2721, -0.33 DPS) [world_drop]; Silk Headband (7050, -0.59 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.10 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 | yes | Necklace of Calisea (1714, -0.64 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.88 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.95 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.2 | yes | Green Silken Shoulders (7057, -0.03 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.03 DPS) [crafted]; Caretaker's Cape (19532, -0.16 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.5 | yes | Dreamweave Vest (10021, -0.43 DPS) [crafted]; Robe of Power (7054, -0.85 DPS) [crafted]; Elemental Raiment (9434, -1.28 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (101.5 DPS) | yes | Condor Bracers (15864, -0.26 DPS) [quest]; Earthen Silk Cuffs (254019, -0.65 DPS) [crafted]; Arcane Runed Bracers (4744, -1.20 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | Red Mageweave Gloves (10018, -0.72 DPS) [crafted]; Gilded Handwraps (254021, -1.21 DPS) [crafted]; Black Mageweave Gloves (10003, -2.01 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.0 | yes | Highlander's Cloth Girdle (20099, -0.43 DPS) [rep]; Deathmage Sash (10771, -0.56 DPS) [dungeon]; Star Belt (4329, -0.64 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.0 | yes | Crimson Silk Pantaloons (7062, -0.75 DPS) [crafted]; Abomination Skin Leggings (23173, -0.78 DPS) [dungeon]; Gaze Dreamer Pants (6903, -2.93 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, +0.00 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.09 DPS) [crafted]; Acidic Walkers (9454, -2.23 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 | yes | Ring of Forlorn Spirits (2043, -0.46 DPS) [quest]; Reedknot Ring (9622, -0.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.72 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.90 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (100.2 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Rune of Duty (21567, +0.00 DPS, sim-verified) [rep] |
| main_hand | Windweaver Staff (7757) | Scarlet Monastery: Houndmaster Loksey [dungeon] | sim-verified (101.5 DPS) | yes | Staff of Jordan (873, -0.13 DPS) [world_drop]; Glimmering Staff (249392, -0.13 DPS) [crafted]; Illusionary Rod (7713, -1.21 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 302.5 | yes | Nether Force Wand (11263, -1.55 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.29 DPS) [quest]; Ragefire Wand (7513, -2.34 DPS) [quest] |

**New at 40:** head: Augural Shroud; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Windweaver Staff; ranged: Jaina's Firestarter

No-known-source sample (15 of 293, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 140.0. Weights run: 0.8s. Verify run: 0.8s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.053 ± 1.050), crit=4.846 ± 0.439, hit=7.393 ± 0.114, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 94.5 | yes | Eye of Theradras (17715, -1.16 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -6.91 DPS) [crafted]; Dreamweave Circlet (10041, -7.99 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.7 | yes | Mindburst Medallion (11196, -0.31 DPS) [quest]; Darkspear Warding Pendant (272073, -0.67 DPS) [vendor]; Scorn's Icy Choker (23169, -1.10 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) | Captain Dirgehammer [vendor] | 85.3 | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -7.93 DPS) [crafted]; Inquisitor's Shawl (19507, -8.20 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.3 | yes | Darkspear Raider's Cloak (272076, -0.71 DPS) [vendor]; Big Voodoo Cloak (8216, -0.74 DPS) [crafted]; Runecloth Cloak (13860, -1.20 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (138.3 DPS) | yes | Runecloth Robe (13858, -1.42 DPS) [crafted]; Runecloth Tunic (13857, -1.46 DPS) [crafted]; Knight's Dreadweave Vest (220886, -1.66 DPS, sim-verified) [vendor] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (138.0 DPS) | yes | Forgotten Wraps (9433, -0.22 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.22 DPS) [quest]; Bloodband Bracers (11469, -1.41 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 100.7 | yes | Raider Handwraps (272098, -1.24 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -9.92 DPS) [vendor]; Dreamweave Gloves (10019, -9.96 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 82.1 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.31 DPS) [dungeon]; Deathmage Sash (10771, -7.53 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 92.5 | yes | Red Mageweave Pants (10009, -0.52 DPS, sim-verified) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.95 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -8.98 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 91.4 | yes | Earthen Silk Slippers (254013, -0.42 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -9.00 DPS) [crafted]; Southsea Mojo Boots (20641, -9.11 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 73.9 | yes | Band of the Unicorn (7553, -7.73 DPS) [world_drop]; Lorekeeper's Ring (19523, -7.86 DPS) [rep]; Lorekeeper's Ring (19524, -8.24 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.3 | yes | Lorekeeper's Ring (19523, -0.55 DPS) [rep]; Lorekeeper's Ring (19524, -0.93 DPS) [rep]; Band of the Unicorn (7553, -1.64 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (135.7 DPS) | yes | Uther's Strength (11302, -0.46 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -8.44 DPS) [quest]; Tidal Charm (1404, -8.44 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (135.7 DPS) | yes | Illusionary Rod (7713, -0.80 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.60 DPS) [dungeon]; Shortsword of Vengeance (754, -2.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 | yes | Noxious Shooter (17745, -0.54 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.73 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Guardian Talisman; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 391.3. Weights run: 0.9s. Verify run: 0.9s. 777 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.134 ± 1.662), crit=7.715 ± 0.552, hit=13.218 ± 0.226, spell_haste=not significant (-2.222 ± 2.347), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bloodvine Goggles (19999) | Engineering [crafted] | 372.4 | yes | Fireleaf Circlet (240056, -0.51 DPS, sim-verified) [vendor]; Field Marshal's Coronet (16441, -45.80 DPS) [vendor]; Field Marshal's Coronet (231604, -45.80 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (386.8 DPS) | yes | Jewel of Kajaro (19601, -4.68 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -5.22 DPS) [quest]; Amulet of the Dawn (22657, -22.12 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 166.0 | yes | Rugged Mantle of the Timbermaw (227808, -4.76 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.67 DPS) [crafted]; Lieutenant Commander's Silk Mantle (23319, -6.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 157.3 | yes | Howler's Furs (272414, -5.41 DPS) [vendor]; Stalwart Cloak (272415, -5.41 DPS) [vendor]; Earthweave Cloak (21187, -13.84 DPS, sim-verified) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | sim-verified (385.0 DPS) | yes | Bloodvine Vest (19682, -4.53 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -22.65 DPS) [vendor]; Robe of the Archmage (14152, -30.01 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 161.2 | yes | Fireleaf Wristwraps (240044, -2.49 DPS) [vendor]; Rockfury Bracers (21186, -11.21 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -28.09 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 273.9 | yes | Gloves of Spell Mastery (14146, -16.74 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -22.00 DPS) [vendor]; Sorcerer's Gloves (22066, -24.58 DPS) [quest] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 192.9 | yes | Fireleaf Waistguard (240045, -5.41 DPS) [vendor]; Belt of the Archmage (18405, -10.10 DPS) [crafted]; Knowledge of the Timbermaw (228190, -10.18 DPS, sim-verified) [vendor] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (385.8 DPS) | yes | Bloodvine Leggings (19683, -2.93 DPS) [crafted]; Sorcerer's Leggings (226933, -4.37 DPS) [quest]; Sentinel's Silk Leggings (237815, -5.31 DPS, sim-verified) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 169.3 | yes | Marshal's Silk Footwraps (16437, -0.06 DPS) [vendor]; Marshal's Silk Footwraps (231606, -0.06 DPS) [vendor]; Fireleaf Boots (240050, -0.96 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (386.8 DPS) | yes | Wrath of Cenarius (21190, -9.37 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -15.08 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -15.75 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (386.8 DPS) | yes | Wrath of Cenarius (21190, -9.37 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -15.08 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -15.75 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (384.7 DPS) | yes | Weakness Analyzer (272438, -20.94 DPS) [vendor]; Uther's Strength (11302, -24.40 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -25.69 DPS) [quest] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (386.8 DPS) | yes | Uther's Strength (11302, -1.73 DPS) [world_drop]; Weakness Analyzer (272438, -2.04 DPS, sim-verified) [vendor]; Thunderbrew's Boot Flask (744, -3.02 DPS) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (386.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -13.84 DPS, sim-verified) [world_drop]; Sageblade (22383, -19.09 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 361.4 | yes | Wand of Biting Cold (19108, -13.55 DPS, sim-verified) [quest]; Stormrager (16997, -14.14 DPS) [quest]; Brilliant Wand (249385, -14.82 DPS) [crafted] |

**New at 60:** head: Bloodvine Goggles; neck: Beads of Ogre Might; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Frozen Heart of the Mountain; trinket2: Serenity Field; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 777, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 30.5. Weights run: 0.8s. Verify run: 0.8s. 121 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.495 ± 0.285, crit=1.741 ± 0.088, hit=3.693 ± 0.034, spell_haste=not significant (0.159 ± 0.399), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | sim-verified (28.9 DPS) | yes | Flying Tiger Goggles (4368, -0.45 DPS) [crafted]; Lucky Fishing Hat (19972, -0.45 DPS) [quest]; Shadow Goggles (4373, -0.74 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 18.5 | yes | Reinforced Woolen Shoulders (4315, -0.18 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.09 DPS) [crafted]; Slime-encrusted Pads (6461, -1.39 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.5 | yes | Sanguine Cape (14376, -0.06 DPS, sim-verified) [world_drop]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Feyscale Cloak (6632, -0.26 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 12.5 | yes | Mystic's Wrap (14369, -0.15 DPS) [world_drop]; Mystic's Robe (14371, -0.15 DPS) [world_drop]; Gray Woolen Robe (2585, -0.65 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (28.7 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.11 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.48 DPS, sim-verified) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 10.5 | yes | Tomb Robber's Gloves (280096, +0.00 DPS, sim-verified) [quest]; Pristine Gloves (253913, -0.15 DPS) [crafted]; Adept's Gloves (4768, -0.23 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (28.9 DPS) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Tarantula Silk Sash (3229, -0.19 DPS) [world]; Keller's Girdle (2911, -0.68 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (28.6 DPS) | yes | Darkweave Breeches (12987, -0.34 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.41 DPS, sim-verified) [dungeon]; Filigreed Silky Leggings (253939, -0.45 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 13.0 | yes | Pristine Boots (253889, +0.00 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.53 DPS) [world]; Sanguine Sandals (14374, -0.53 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 9.0 | yes | Loop of Sacrifice (281673, -0.19 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.34 DPS) [world_drop]; Sludge-Stained Band (286535, -0.45 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (28.8 DPS) | yes | Volcanic Rock Ring (12053, -0.04 DPS) [world_drop]; Sludge-Stained Band (286535, -0.15 DPS) [world]; Loop of Sacrifice (281673, -0.62 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 14.9 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.23 DPS) [world]; Lesser Staff of the Spire (1300, -0.45 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 298.7 | yes | Skycaller (12984, -0.59 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 121, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263432 Skulker's Shiv

### Band 30 (orc, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 56.3. Weights run: 0.8s. Verify run: 0.8s. 211 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.956 ± 0.566), crit=3.069 ± 0.225, hit=4.265 ± 0.054, spell_haste=not significant (0.233 ± 0.614), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.6 | yes | Nightsky Cowl (4039, -0.14 DPS, sim-verified) [world_drop]; Holy Shroud (2721, -0.44 DPS) [world_drop]; Shadow Hood (4323, -0.48 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.7 | yes | Darkspear Warding Pendant (272075, -0.85 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.85 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.6 | yes | Death Speaker Mantle (6685, -0.07 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.7 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.18 DPS) [world_drop]; Hillman's Cloak (3719, -0.25 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.4 | yes | Death Speaker Robes (6682, -0.34 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.74 DPS) [crafted]; Tree Bark Jacket (1486, -0.81 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.31 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.31 DPS) [quest]; Glowing Magical Bracelets (13106, -0.60 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 10.8 | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.30 DPS) [dungeon]; Serpent Gloves (5970, -0.36 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.9 | yes | Crimson Silk Belt (7055, -0.04 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -0.19 DPS) [dungeon]; Invoker's Cord (215366, -0.20 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (56.3 DPS) | yes | Gaze Dreamer Pants (6903, -0.16 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.19 DPS) [crafted]; Abomination Skin Leggings (23173, -0.96 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.7 | yes | Spidersilk Boots (4320, -0.27 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -1.08 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Snake Hoop (6750, -0.03 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 6.7 | yes | Snake Hoop (6750, +0.00 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.07 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (55.0 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Defiler's Talisman (21120, -2.28 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 10.5 | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Channeler's Staff (4437, -0.27 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 349.6 | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.02 DPS) [crafted]; Gravestone Scepter (7001, -4.50 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Black Widow Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 211, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19545 Scout's Blade

### Band 40 (orc, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 92.4. Weights run: 0.9s. Verify run: 0.8s. 288 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (0.249 ± 0.763), crit=2.900 ± 0.288, hit=4.986 ± 0.077, spell_haste=6.506 ± 0.958, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -1.31 DPS) [world_drop]; Silk Headband (7050, -1.57 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 | yes | Necklace of Calisea (1714, -0.51 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.88 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.95 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.2 | yes | Inquisitor's Shawl (19507, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest]; Green Silken Shoulders (7057, -0.38 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.03 DPS) [crafted]; Battle Healer's Cloak (19528, -0.16 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.5 | yes | Dreamweave Vest (10021, -0.43 DPS) [crafted]; Robe of Power (7054, -0.85 DPS) [crafted]; Elemental Raiment (9434, -2.12 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Radiant Silver Bracers (4545, -0.39 DPS) [quest]; Earthen Silk Cuffs (254019, -0.65 DPS) [crafted]; Condor Bracers (15864, -0.96 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | Red Mageweave Gloves (10018, -0.72 DPS) [crafted]; Gilded Handwraps (254021, -1.21 DPS) [crafted]; Black Mageweave Gloves (10003, -1.83 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.0 | yes | Defiler's Cloth Girdle (20164, -0.43 DPS) [rep]; Warsong Sash (16975, -0.52 DPS) [quest]; Star Belt (4329, -1.35 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 17.0 | yes | Crimson Silk Pantaloons (7062, -0.75 DPS) [crafted]; Abomination Skin Leggings (23173, -0.78 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.46 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.51 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.09 DPS) [crafted]; Acidic Walkers (9454, -2.23 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.5 | yes | Reedknot Ring (9622, -0.59 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.72 DPS) [vendor]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.39 DPS) [vendor]; Reedknot Ring (9622, -0.96 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 42.1 | yes | Windweaver Staff (7757, +0.00 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -5.15 DPS) [world_drop]; Glimmering Staff (249392, -5.15 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 302.5 | yes | Nether Force Wand (11263, -1.22 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.29 DPS) [quest]; Ragefire Wand (7513, -2.34 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 126.9. Weights run: 0.8s. Verify run: 0.8s. 375 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.053 ± 1.050), crit=4.846 ± 0.439, hit=7.393 ± 0.114, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 94.5 | yes | Eye of Theradras (17715, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -6.91 DPS) [crafted]; Dreamweave Circlet (10041, -7.99 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.7 | yes | Scorn's Icy Choker (23169, -0.27 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.31 DPS) [quest]; Darkspear Warding Pendant (272073, -0.67 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (125.9 DPS) | yes | Red Mageweave Shoulders (10029, -1.16 DPS) [crafted]; Inquisitor's Shawl (19507, -1.43 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.76 DPS, sim-verified) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.5 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.51 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.85 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (126.2 DPS) | yes | Runecloth Robe (13858, -1.42 DPS) [crafted]; Runecloth Tunic (13857, -1.46 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -2.08 DPS, sim-verified) [vendor] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 14.5 | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Forgotten Wraps (9433, -0.23 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.23 DPS) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 100.7 | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -9.92 DPS) [vendor]; Dreamweave Gloves (10019, -9.96 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 82.1 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.31 DPS) [dungeon]; Deathmage Sash (10771, -7.53 DPS) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 92.5 | yes | Red Mageweave Pants (10009, -0.78 DPS, sim-verified) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.95 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -8.98 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 91.4 | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -9.00 DPS) [crafted]; Southsea Mojo Boots (20641, -9.11 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 73.9 | yes | Band of the Unicorn (7553, -7.73 DPS) [world_drop]; Advisor's Ring (19519, -7.86 DPS) [rep]; Advisor's Ring (19520, -8.24 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.3 | yes | Band of the Unicorn (7553, +0.00 DPS, sim-verified) [world_drop]; Advisor's Ring (19519, -0.55 DPS) [rep]; Advisor's Ring (19520, -0.93 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (123.5 DPS) | yes | Uther's Strength (11302, -0.43 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -8.44 DPS) [vendor]; Guardian Talisman (1490, -8.44 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (123.5 DPS) | yes | Uther's Strength (11302, -0.56 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -6.57 DPS) [vendor]; Guardian Talisman (1490, -6.57 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (123.5 DPS) | yes | Illusionary Rod (7713, -0.80 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.60 DPS) [dungeon]; Shortsword of Vengeance (754, -3.45 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 | yes | Noxious Shooter (17745, -0.04 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.73 DPS) [crafted] |

**New at 50:** head: Blood Guard's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 375, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 370.5. Weights run: 0.9s. Verify run: 0.8s. 774 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.134 ± 1.662), crit=7.715 ± 0.552, hit=13.218 ± 0.226, spell_haste=not significant (-2.222 ± 2.347), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (355.8 DPS) | yes | Warlord's Silk Cowl (16533, -6.32 DPS) [vendor]; Warlord's Silk Cowl (231601, -6.32 DPS) [vendor]; Bloodvine Goggles (19999, -6.60 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (347.7 DPS) | yes | Jewel of Kajaro (19601, -2.93 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -5.22 DPS) [quest]; Amulet of the Dawn (22657, -22.12 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 166.0 | yes | Rugged Mantle of the Timbermaw (227808, -4.24 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.67 DPS) [crafted]; Champion's Silk Mantle (23264, -6.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 157.3 | yes | Howler's Furs (272414, -5.41 DPS) [vendor]; Stalwart Cloak (272415, -5.41 DPS) [vendor]; Earthweave Cloak (21187, -12.85 DPS, sim-verified) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | sim-verified (356.5 DPS) | yes | Bloodvine Vest (19682, -7.32 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -22.65 DPS) [vendor]; Robe of the Archmage (14152, -30.01 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 161.2 | yes | Rockfury Bracers (21186, -1.17 DPS, sim-verified) [quest]; Fireleaf Wristwraps (240044, -2.49 DPS) [vendor]; Dryad's Wrist Bindings (19595, -28.09 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 273.9 | yes | Gloves of Spell Mastery (14146, -15.61 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -22.00 DPS) [vendor]; Sorcerer's Gloves (22066, -24.58 DPS) [quest] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 192.9 | yes | Fireleaf Waistguard (240045, -5.41 DPS) [vendor]; Knowledge of the Timbermaw (228190, -5.90 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -10.10 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (355.9 DPS) | yes | Bloodvine Leggings (19683, -2.93 DPS) [crafted]; Sorcerer's Leggings (226933, -4.37 DPS) [quest]; Sentinel's Silk Leggings (237815, -6.68 DPS, sim-verified) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 169.3 | yes | Fireleaf Boots (240050, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (16539, -0.06 DPS) [vendor]; General's Silk Boots (231597, -0.06 DPS) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (347.7 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Channeler's Ring (272406, +0.00 DPS) [vendor]; Don Julio's Band (19325, -1.64 DPS, sim-verified) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (355.9 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Signet Ring of the Bronze Dragonflight (234028, -0.68 DPS) [vendor]; Band of Earthen Might (21182, -6.68 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (346.1 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -1.73 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.48 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (346.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted]; Weakness Analyzer (272438, -15.23 DPS) [vendor]; Uther's Strength (11302, -18.69 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (347.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -13.28 DPS, sim-verified) [world_drop]; Sageblade (22383, -19.09 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 361.4 | yes | Stormrager (16997, -14.14 DPS) [quest]; Brilliant Wand (249385, -14.82 DPS) [crafted]; Wand of Biting Cold (19108, -17.37 DPS, sim-verified) [quest] |

**New at 60:** head: Fireleaf Circlet; neck: Beads of Ogre Might; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Bloodvine Boots; finger1: Wrath of Cenarius; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Serenity Field; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 774, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

