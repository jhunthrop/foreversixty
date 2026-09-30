# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 32.4. Weights run: 0.9s. Verify run: 0.8s. 215 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.268 ± 0.163, crit=0.822 ± 0.037, hit=1.874 ± 0.118, spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 11.0 | yes | Totemic Leather Hood (252448, -0.19 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.38 DPS) [crafted]; Trapper's Leather Hood (252505, -0.87 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.4 | yes | Reinforced Woolen Shoulders (4315, -0.34 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.19 DPS) [crafted]; Rough Bronze Shoulders (3480, -1.57 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.8 | yes | Sanguine Cape (14376, -0.13 DPS, sim-verified) [world_drop]; Heavy Woolen Cloak (4311, -0.17 DPS) [crafted]; Feyscale Cloak (6632, -0.27 DPS) [dungeon] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 14.3 | yes | Acolyte's Chain Shirt (250491, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Armor (252493, -0.29 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.29 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (31.4 DPS) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.12 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.55 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | sim-verified (31.5 DPS) | yes | Acolyte's Gloves (250511, -0.07 DPS) [crafted]; Blight Gloves (279877, -0.11 DPS) [quest]; Fletcher's Gloves (7348, -0.64 DPS, sim-verified) [crafted] |
| waist | Stormrider's Leather Belt (252432) | Leatherworking [crafted] | sim-verified (31.3 DPS) | yes | Pristine Sash (253925, +0.00 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.39 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.1 | yes | Stormrider's Leather Pants (252502, -0.15 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.31 DPS) [crafted]; Dreamer's Leggings (270016, -0.67 DPS, sim-verified) [quest] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.3 | yes | Spidersilk Boots (4320, +0.00 DPS, sim-verified) [crafted]; Acolyte's Boots (250506, -0.19 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.19 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 7.6 | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Lorekeeper's Ring (20431, -0.25 DPS) [rep]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.5 | yes | Lorekeeper's Ring (20431, -0.24 DPS) [rep]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Loop of Sacrifice (281673, -0.57 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 12.7 | yes | Channeler's Staff (4437, -0.24 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.45 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -1.38 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Stormrider's Leather Belt; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; trinket1: Rune of Duty; trinket2: Rune of Perfection; main_hand: Twisted Chanter's Staff; ranged: Kajaric Icon

No-known-source sample (15 of 215, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7956 Bronze Warhammer; 10047 Simple Kilt; 10421 Rough Copper Vest; 14147 Cavedweller Bracers

### Band 30 (dwarf, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 62.8. Weights run: 1.1s. Verify run: 0.9s. 373 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.138 ± 0.340), crit=1.742 ± 0.105, hit=2.217 ± 0.197, spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 18.2 | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.42 DPS) [crafted]; Nightsky Cowl (4039, -0.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 | yes | Crystal Starfire Medallion (5003, -0.93 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.93 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.14 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.2 | yes | Fairywing Mantle (9536, -0.30 DPS) [quest]; Death Speaker Mantle (6685, -0.37 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.40 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.1 | yes | Darkspear Raider's Cloak (272078, -0.01 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.15 DPS) [quest]; Resilient Cape (14400, -0.23 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.8 | yes | Guardian Armor (4256, -0.06 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.43 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.70 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (60.8 DPS) | yes | Nightsky Wristbands (6407, -0.22 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.22 DPS) [quest]; Glowing Magical Bracelets (13106, -0.81 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | sim-verified (61.4 DPS) | yes | Stormrider's Leather Gloves (252498, -0.70 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.74 DPS) [dungeon]; Fletcher's Gloves (7348, -1.43 DPS, sim-verified) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | sim-verified (60.8 DPS) | yes | Highlander's Cloth Girdle (20099, -0.05 DPS) [rep]; Guardian Belt (4258, -0.10 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.83 DPS, sim-verified) [crafted] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 22.5 | yes | Abomination Skin Leggings (23173, -0.46 DPS, sim-verified) [dungeon]; Guardian Pants (5962, -0.54 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.54 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 | yes | Stormrider's Leather Boots (252443, -0.33 DPS) [crafted]; Spidersilk Boots (4320, -0.34 DPS) [crafted]; Acidic Walkers (9454, -1.44 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.0 | yes | Lorekeeper's Ring (19525, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Snake Hoop (6750, -0.51 DPS, sim-verified) [quest] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | sim-verified (60.7 DPS) | yes | Lorekeeper's Ring (19525, -0.03 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.04 DPS) [dungeon]; Snake Hoop (6750, -0.69 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (54.2 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Talisman of Arathor (21119, -0.50 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (60.2 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Glimmering Staff (249392, +0.00 DPS) [crafted]; Mechanic's Pipehammer (9604, -5.59 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Prefect's Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Black Widow Band; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 373, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (dwarf, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 78.6. Weights run: 1.2s. Verify run: 1.0s. 512 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.663 ± 0.634), crit=2.997 ± 0.195, hit=4.584 ± 0.401, spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 42.0 | yes | Spellpower Goggles Xtreme (10502, +0.00 DPS, sim-verified) [crafted]; Big Voodoo Mask (8201, -2.11 DPS) [crafted]; Augural Shroud (2620, -2.16 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Darkspear Warding Pendant (272074, -0.56 DPS) [vendor]; Necklace of Calisea (1714, -0.61 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -0.68 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.18 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.18 DPS) [vendor]; Icy Cloak (4327, -0.21 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.36 DPS) [crafted]; Elemental Raiment (9434, -0.44 DPS) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.0 | yes | Arcane Runed Bracers (4744, -0.09 DPS) [quest]; Spidertank Oilrag (9448, -0.09 DPS) [dungeon]; Turtle Scale Bracers (8198, -0.64 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Dragonscale Gauntlets (8347), Shadowskin Gloves (18238)) | World drop [world_drop] | 42.0 | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Fletcher's Gloves (7348, -0.79 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (78.6 DPS) | yes | Highlander's Cloth Girdle (20098, -0.03 DPS) [rep]; Skycaller's Leather Belt (252522, -0.26 DPS) [crafted]; Highlander's Chain Girdle (20089, -1.08 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Crimson Silk Pantaloons (7062, -0.47 DPS) [crafted]; Abomination Skin Leggings (23173, -0.68 DPS) [dungeon]; Kodohide Legguards (285338, -0.78 DPS, sim-verified) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -0.35 DPS, sim-verified) [crafted]; Skycaller's Mail Boots (252563, -0.57 DPS) [crafted]; Mender's Leather Shoes (252533, -1.01 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 | yes | Ring of Forlorn Spirits (2043, -0.53 DPS) [quest]; Reedknot Ring (9622, -0.62 DPS) [quest]; Minor Channeling Ring (1449, -0.68 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.18 DPS) [quest]; Lorekeeper's Ring (19525, -0.18 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.28 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (64.5 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Rune of Duty (21567, -0.71 DPS, sim-verified) [rep] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (77.3 DPS) | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Mechanic's Pipehammer (9604, +0.00 DPS) [quest]; Fiery War Axe (870, -7.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; ranged: Totem of Ancestral Protection

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (dwarf, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 104.1. Weights run: 1.2s. Verify run: 1.1s. 674 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (3.110 ± 0.789), crit=3.848 ± 0.256, hit=6.081 ± 0.559, spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | sim-verified (101.0 DPS) | yes | Knight-Lieutenant's Mail Helmet (223075, -0.49 DPS) [vendor]; Winged Helm (13112, -0.85 DPS) [world_drop]; Eye of Theradras (17715, -1.33 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 43.5 | yes | Darkspear Warding Pendant (272073, -0.49 DPS, sim-verified) [vendor]; Scorn's Icy Choker (23169, -1.58 DPS) [dungeon]; Mindburst Medallion (11196, -1.66 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 77.2 | yes | Knight-Lieutenant's Mail Epaulets (223073, -0.44 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -0.72 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.08 DPS) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 43.5 | yes | Imperial Red Cloak (8248, +0.00 DPS, sim-verified) [world_drop]; Darkspear Raider's Cloak (272077, -0.82 DPS) [vendor]; Runecloth Cloak (13860, -0.85 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (100.8 DPS) | yes | Wildthorn Mail (12624, -0.83 DPS) [crafted]; Knight's Mail Armor (223078, -1.13 DPS, sim-verified) [vendor]; Runecloth Robe (13858, -1.53 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 41.1 | yes | Shizzle's Nozzle Wiper (11917, -0.33 DPS) [quest]; Forgotten Wraps (9433, -0.55 DPS, sim-verified) [world_drop]; Imperial Red Bracers (8247, -0.61 DPS) [world_drop] |
| hands | Fists of The Five Thunders (227022) | Mokvar [vendor] | 94.0 | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Sergeant Major's Mail Gauntlets (223076, -2.16 DPS) [vendor]; Raider Handguards (272102, -2.48 DPS) [vendor] |
| waist | Highlander's Lizardhide Girdle (20103) (or Highlander's Mail Girdle (20118)) | The League of Arathor [rep] | 85.0 | yes | Highlander's Mail Girdle (20118, +0.00 DPS, sim-verified) [vendor]; Highlander's Cloth Girdle (20097, -0.58 DPS) [rep]; Dawnspire Cord (12466, -1.75 DPS) [dungeon] |
| legs | Knight's Mail Legplates (223074) | Captain Dirgehammer [vendor] | sim-verified (101.5 DPS) | yes | Stormshroud Pants (15057, -1.88 DPS, sim-verified) [crafted]; Kilt of the Atal'ai Prophet (10807, -2.84 DPS) [dungeon]; Red Mageweave Pants (10009, -3.51 DPS) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 102.9 | yes | Skycaller's Leather Boots (252471, +0.00 DPS, sim-verified) [crafted]; Skycaller's Mail Sabatons (252577, -4.73 DPS) [crafted]; Mender's Leather Boots (252472, -5.26 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 60.8 | yes | Ogremind Ring (1993, -3.44 DPS) [world_drop]; Voodoo Band (1996, -3.44 DPS) [world_drop]; Mindbender Loop (5009, -3.44 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 28.7 | yes | Voodoo Band (1996, -0.61 DPS) [world_drop]; Mindbender Loop (5009, -0.61 DPS) [world_drop]; Ogremind Ring (1993, -1.00 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (89.0 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS, sim-verified) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (99.9 DPS) | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Hammer of the Northern Wind (810, -8.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Fists of The Five Thunders; waist: Highlander's Lizardhide Girdle; legs: Knight's Mail Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Guardian Talisman

No-known-source sample (15 of 674, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (dwarf, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 169.1. Weights run: 1.1s. Verify run: 1.0s. 1303 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.046 ± 0.865), crit=5.588 ± 0.380, hit=6.652 ± 0.739, spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Crown (240123) | Leonid Barthalomew the Revered [vendor] | sim-verified (161.6 DPS) | yes | Mask of the Unforgiven (13404, -1.52 DPS) [dungeon]; Bloodvine Goggles (19999, -1.52 DPS) [crafted]; Soulcrusher Headpiece (240096, -2.03 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (159.5 DPS) | yes | Blazefury Medallion (17111, -0.76 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -1.02 DPS) [quest]; Amulet of the Dawn (22657, -4.30 DPS) [quest] |
| shoulder | Soulcrusher Mantle (240125) | Leonid Barthalomew the Revered [vendor] | 198.3 | yes | Rugged Mantle of the Timbermaw (227808, -2.70 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -7.76 DPS) [crafted]; Shroud of the Nathrezim (18720, -8.35 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 90.9 | yes | Earthweave Cloak (21187, -2.11 DPS) [quest]; Howler's Furs (272414, -2.11 DPS) [vendor]; Chromatic Cloak (18509, -4.20 DPS, sim-verified) [crafted] |
| chest | Soulcrusher Embrace (240109) | Leonid Barthalomew the Revered [vendor] | sim-verified (162.1 DPS) | yes | Soulcrusher Armor (240128, -2.56 DPS, sim-verified) [vendor]; Soulcrusher Chestguard (240101, -3.40 DPS) [vendor]; Bloodvine Vest (19682, -3.48 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 93.5 | yes | Soulcrusher Bracers (240108, +0.00 DPS, sim-verified) [vendor]; Soulcrusher Vambraces (240137, -0.98 DPS) [vendor]; Primal Batskin Bracers (19687, -2.34 DPS) [crafted] |
| hands | Soulcrusher Mitts (240122) | Leonid Barthalomew the Revered [vendor] | 207.0 | yes | Soulcrusher Handguards (240095, -1.86 DPS, sim-verified) [vendor]; Soulcrusher Grips (240130, -3.20 DPS) [vendor]; Stormshroud Gloves (21278, -5.40 DPS) [crafted] |
| waist | Soulcrusher Girdle (240099) | Leonid Barthalomew the Revered [vendor] | 197.3 | yes | Soulcrusher Waistguard (240107, -1.04 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -7.14 DPS) [crafted]; Knowledge of the Timbermaw (228190, -7.17 DPS) [vendor] |
| legs | Soulcrusher Legguards (240097) | Leonid Barthalomew the Revered [vendor] | sim-verified (166.3 DPS) | yes | Sentinel's Silk Leggings (237815, -1.18 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.18 DPS) [vendor]; Sentinel's Chain Leggings (237819, -6.73 DPS, sim-verified) [vendor] |
| feet | Soulcrusher Greaves (240110) | Leonid Barthalomew the Revered [vendor] | 119.0 | yes | Soulcrusher Sabatons (240102, -2.09 DPS) [vendor]; Greaves of Withering Despair (22240, -2.69 DPS) [dungeon]; Bloodvine Boots (19684, -2.96 DPS, sim-verified) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (159.5 DPS) | yes | Wrath of Cenarius (21190, -1.18 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.32 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -3.54 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (159.5 DPS) | yes | Wrath of Cenarius (21190, -0.99 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.32 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -3.54 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (161.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -0.69 DPS) [world_drop] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (159.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -1.97 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (159.5 DPS) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Fist of Cenarius (21188, +0.00 DPS, sim-verified) [quest]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (23199) (or Totem of Thunder (228176), Tidal Totem (272431), Totem of the Storm (272432), Burning Totem (272433), Totem of Urgency (279249), Totem of Ancestral Protection (249443), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | World drop [world_drop] | 0.0 | yes | Totem of Thunder (228176, +0.00 DPS, sim-verified) [vendor]; Tidal Totem (272431, +0.00 DPS) [vendor]; Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Soulcrusher Crown; neck: Medallion of the Dawn; shoulder: Soulcrusher Mantle; back: Arcanoweave Cloak; chest: Soulcrusher Embrace; wrist: Rockfury Bracers; hands: Soulcrusher Mitts; waist: Soulcrusher Girdle; legs: Soulcrusher Legguards; feet: Soulcrusher Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Darkmoon Card: Maelstrom; ranged: Totem of the Storm

No-known-source sample (15 of 1303, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (orc, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 32.2. Weights run: 0.9s. Verify run: 0.8s. 211 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.268 ± 0.163, crit=0.822 ± 0.037, hit=1.874 ± 0.118, spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 11.0 | yes | Totemic Leather Hood (252448, -0.19 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.38 DPS) [crafted]; Trapper's Leather Hood (252505, -0.67 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.4 | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.19 DPS) [crafted]; Rough Bronze Shoulders (3480, -1.57 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.8 | yes | Sanguine Cape (14376, -0.16 DPS, sim-verified) [world_drop]; Heavy Woolen Cloak (4311, -0.17 DPS) [crafted]; Feyscale Cloak (6632, -0.27 DPS) [dungeon] |
| chest | Acolyte's Chain Shirt (250491) | Blacksmithing [crafted] | sim-verified (30.6 DPS) | yes | Wisdom's Leather Armor (252493, +0.00 DPS) [crafted]; Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.39 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (30.7 DPS) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.42 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | sim-verified (31.4 DPS) | yes | Acolyte's Gloves (250511, -0.07 DPS) [crafted]; Blight Gloves (279877, -0.11 DPS) [quest]; Fletcher's Gloves (7348, -1.18 DPS, sim-verified) [crafted] |
| waist | Stormrider's Leather Belt (252432) | Leatherworking [crafted] | sim-verified (30.9 DPS) | yes | Pristine Sash (253925, +0.00 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.66 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.1 | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Acolyte's Chain Leggings (250496, -0.31 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.3 | yes | Spidersilk Boots (4320, +0.00 DPS, sim-verified) [crafted]; Acolyte's Boots (250506, -0.19 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.19 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 7.6 | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; Sludge-Stained Band (286535, -0.44 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (30.9 DPS) | yes | Volcanic Rock Ring (12053, -0.11 DPS) [world_drop]; Sludge-Stained Band (286535, -0.19 DPS) [world]; Loop of Sacrifice (281673, -0.67 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 12.7 | yes | Channeler's Staff (4437, -0.24 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.45 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -1.15 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Acolyte's Chain Shirt; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Stormrider's Leather Belt; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Kajaric Icon

No-known-source sample (15 of 211, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7956 Bronze Warhammer; 10047 Simple Kilt; 10421 Rough Copper Vest; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade

### Band 30 (orc, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 60.1. Weights run: 1.1s. Verify run: 0.8s. 370 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.138 ± 0.340), crit=1.742 ± 0.105, hit=2.217 ± 0.197, spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 18.2 | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.42 DPS) [crafted]; Nightsky Cowl (4039, -0.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 | yes | Crystal Starfire Medallion (5003, -0.93 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.93 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.47 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.2 | yes | Fairywing Mantle (9536, -0.30 DPS) [quest]; Death Speaker Mantle (6685, -0.38 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.40 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.1 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.23 DPS) [world_drop]; Soft Willow Cape (16661, -0.34 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.8 | yes | Guardian Armor (4256, -0.26 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.43 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.70 DPS) [crafted] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 9.1 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -0.23 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.23 DPS) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | sim-verified (59.0 DPS) | yes | Oilrag Handwraps (16741, -0.04 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.21 DPS) [crafted]; Fletcher's Gloves (7348, -0.94 DPS, sim-verified) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | sim-verified (58.7 DPS) | yes | Defiler's Cloth Girdle (20164, -0.05 DPS) [rep]; Guardian Belt (4258, -0.10 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.64 DPS, sim-verified) [crafted] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 22.5 | yes | Guardian Pants (5962, -0.54 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.54 DPS) [crafted]; Abomination Skin Leggings (23173, -0.65 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 | yes | Stormrider's Leather Boots (252443, -0.33 DPS) [crafted]; Spidersilk Boots (4320, -0.34 DPS) [crafted]; Acidic Walkers (9454, -1.33 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.0 | yes | Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.20 DPS) [vendor]; Snake Hoop (6750, -0.65 DPS, sim-verified) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | sim-verified (58.6 DPS) | yes | Lavishly Jeweled Ring (1156, -0.02 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor]; Snake Hoop (6750, -0.63 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (51.0 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Defiler's Talisman (21120, -0.49 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (58.1 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, +0.00 DPS) [quest]; Glimmering Staff (249392, -6.98 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Glowing Magical Bracelets; hands: Jutebraid Gloves; waist: Prefect's Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Advisor's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 370, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7956 Bronze Warhammer; 7958 Bronze Battle Axe; 9362 Brilliant Gold Ring; 10047 Simple Kilt

### Band 40 (orc, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 75.9. Weights run: 1.2s. Verify run: 1.0s. 510 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.663 ± 0.634), crit=2.997 ± 0.195, hit=4.584 ± 0.401, spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (74.4 DPS) | yes | Big Voodoo Mask (8201, -0.24 DPS) [crafted]; Augural Shroud (2620, -0.30 DPS) [world]; Raging Berserker's Helm (7719, -0.76 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Necklace of Calisea (1714, -0.24 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.56 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.68 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 | yes | Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.18 DPS) [quest]; Green Silken Shoulders (7057, -0.91 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.18 DPS) [vendor]; Icy Cloak (4327, -0.21 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.36 DPS) [crafted]; Elemental Raiment (9434, -0.44 DPS) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.0 | yes | Radiant Silver Bracers (4545, -0.06 DPS) [quest]; Spidertank Oilrag (9448, -0.09 DPS) [dungeon]; Turtle Scale Bracers (8198, -0.56 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Dragonscale Gauntlets (8347), Shadowskin Gloves (18238)) | World drop [world_drop] | 42.0 | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Fletcher's Gloves (7348, -0.70 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (75.1 DPS) | yes | Defiler's Cloth Girdle (20166, -0.03 DPS) [rep]; Skycaller's Leather Belt (252522, -0.26 DPS) [crafted]; Defiler's Chain Girdle (20153, -1.42 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Crimson Silk Pantaloons (7062, -0.47 DPS) [crafted]; Abomination Skin Leggings (23173, -0.68 DPS) [dungeon]; Kodohide Legguards (285338, -0.74 DPS, sim-verified) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -0.44 DPS, sim-verified) [crafted]; Skycaller's Mail Boots (252563, -0.57 DPS) [crafted]; Mender's Leather Shoes (252533, -1.01 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 | yes | Reedknot Ring (9622, -0.62 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.71 DPS) [vendor]; Ogremind Ring (1993, -0.83 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.18 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.27 DPS) [vendor]; Reedknot Ring (9622, -1.46 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (73.4 DPS) | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Fiery War Axe (870, -3.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; ranged: Totem of Ancestral Protection

No-known-source sample (15 of 510, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 7956 Bronze Warhammer

### Band 50 (orc, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 104.6. Weights run: 1.2s. Verify run: 1.0s. 660 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (3.110 ± 0.789), crit=3.848 ± 0.256, hit=6.081 ± 0.559, spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 113.0 | yes | Blood Guard's Pulsing Helmet (220848, -0.28 DPS, sim-verified) [vendor]; Blood Guard's Inscribed Skullcap (220842, -2.19 DPS) [vendor]; Red Mageweave Headband (10033, -2.80 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 43.5 | yes | Darkspear Warding Pendant (272073, -0.35 DPS, sim-verified) [vendor]; Scorn's Icy Choker (23169, -1.58 DPS) [dungeon]; Mindburst Medallion (11196, -1.66 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (102.9 DPS) | yes | Blood Guard's Mail Epaulets (220823, -0.69 DPS) [vendor]; Blood Guard's Inscribed Shoulder Pads (220841, -0.69 DPS) [vendor]; Blood Guard's Pulsing Shoulders (220849, -1.56 DPS, sim-verified) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | sim-verified (102.4 DPS) | yes | Imperial Red Cloak (8248, -0.51 DPS) [world_drop]; Darkspear Raider's Cloak (272077, -0.51 DPS) [vendor]; Darkspear Raider's Cloak (272076, -1.08 DPS, sim-verified) [vendor] |
| chest | Stone Guard's Pulsing Breastplate (220844) | Lady Palanseer [vendor] | 116.1 | yes | Stone Guard's Mail Armor (220826, -1.72 DPS, sim-verified) [vendor]; Stone Guard's Inscribed Chestpiece (220838, -2.20 DPS) [vendor]; Acumen Robes (17775, -3.08 DPS) [quest] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 41.1 | yes | Forgotten Wraps (9433, +0.00 DPS, sim-verified) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.33 DPS) [quest]; Imperial Red Bracers (8247, -0.61 DPS) [world_drop] |
| hands | Fists of The Five Thunders (227022) | Mokvar [vendor] | 94.0 | yes | First Sergeant's Pulsing Gauntlets (220845, -0.77 DPS, sim-verified) [vendor]; Raider Handwraps (272098, -1.63 DPS) [vendor]; First Sergeant's Mail Gauntlets (220831, -2.16 DPS) [vendor] |
| waist | Highlander's Mail Girdle (20118) (or Defiler's Lizardhide Girdle (20174), Defiler's Mail Girdle (20196)) | Samuel Hawke [vendor] | 85.0 | yes | Defiler's Lizardhide Girdle (20174, +0.00 DPS, sim-verified) [rep]; Defiler's Mail Girdle (20196, +0.00 DPS) [rep]; Defiler's Cloth Girdle (20165, -0.58 DPS) [rep] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 116.1 | yes | Stormshroud Pants (15057, -1.57 DPS, sim-verified) [crafted]; Stone Guard's Mail Legplates (220834, -2.20 DPS) [vendor]; Stone Guard's Inscribed Legplates (220839, -2.20 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 102.9 | yes | Skycaller's Leather Boots (252471, +0.00 DPS, sim-verified) [crafted]; Skycaller's Mail Sabatons (252577, -4.73 DPS) [crafted]; Mender's Leather Boots (252472, -5.26 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 60.8 | yes | Ogremind Ring (1993, -3.44 DPS) [world_drop]; Voodoo Band (1996, -3.44 DPS) [world_drop]; Mindbender Loop (5009, -3.44 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 28.7 | yes | Ogremind Ring (1993, -0.40 DPS, sim-verified) [world_drop]; Voodoo Band (1996, -0.61 DPS) [world_drop]; Mindbender Loop (5009, -0.61 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (93.6 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -1.61 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (93.6 DPS) | yes | Frozen Heart of the Mountain (249469, -1.41 DPS, sim-verified) [crafted]; Uther's Strength (11302, -3.22 DPS) [world_drop]; Tidal Charm (1404, -3.75 DPS) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (102.0 DPS) | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Fiery War Axe (870, -3.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Runic Leather Bracers; hands: Fists of The Five Thunders; waist: Highlander's Mail Girdle; legs: Stone Guard's Pulsing Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Rune of the Guard Captain

No-known-source sample (15 of 660, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (orc, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 170.1. Weights run: 1.1s. Verify run: 0.9s. 1242 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.046 ± 0.865), crit=5.588 ± 0.380, hit=6.652 ± 0.739, spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulcrusher Headpiece (240096) | Leonid Barthalomew the Revered [vendor] | 286.5 | yes | Soulcrusher Crown (240123, +0.00 DPS, sim-verified) [vendor]; Mask of the Unforgiven (13404, -6.52 DPS) [dungeon]; Bloodvine Goggles (19999, -6.52 DPS) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (161.4 DPS) | yes | Beads of Ogre Might (22150, -1.02 DPS) [quest]; Blazefury Medallion (17111, -2.02 DPS, sim-verified) [world]; Amulet of the Dawn (22657, -4.30 DPS) [quest] |
| shoulder | Soulcrusher Mantle (240125) | Leonid Barthalomew the Revered [vendor] | 198.3 | yes | Champion's Mail Pauldrons (227154, -4.19 DPS) [vendor]; Warlord's Mail Pauldrons (231654, -6.51 DPS, sim-verified) [vendor]; Rugged Mantle of the Timbermaw (227808, -6.71 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 90.9 | yes | Earthweave Cloak (21187, -2.11 DPS) [quest]; Howler's Furs (272414, -2.11 DPS) [vendor]; Chromatic Cloak (18509, -4.76 DPS, sim-verified) [crafted] |
| chest | Soulcrusher Embrace (240109) | Leonid Barthalomew the Revered [vendor] | sim-verified (164.0 DPS) | yes | Soulcrusher Chestguard (240101, -3.40 DPS) [vendor]; Soulcrusher Armor (240128, -3.45 DPS, sim-verified) [vendor]; Bloodvine Vest (19682, -3.48 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 93.5 | yes | Soulcrusher Bracers (240108, -0.12 DPS, sim-verified) [vendor]; Soulcrusher Vambraces (240137, -0.98 DPS) [vendor]; Primal Batskin Bracers (19687, -2.34 DPS) [crafted] |
| hands | Soulcrusher Mitts (240122) | Leonid Barthalomew the Revered [vendor] | 207.0 | yes | Soulcrusher Handguards (240095, -1.98 DPS, sim-verified) [vendor]; Soulcrusher Grips (240130, -3.20 DPS) [vendor]; General's Mail Vices (231655, -4.85 DPS) [vendor] |
| waist | Soulcrusher Girdle (240099) | Leonid Barthalomew the Revered [vendor] | 197.3 | yes | Soulcrusher Waistguard (240107, -4.33 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -7.14 DPS) [crafted]; Knowledge of the Timbermaw (228190, -7.17 DPS) [vendor] |
| legs | Soulcrusher Legguards (240097) | Leonid Barthalomew the Revered [vendor] | sim-verified (166.2 DPS) | yes | Sentinel's Silk Leggings (237815, -1.18 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.18 DPS) [vendor]; Sentinel's Chain Leggings (237819, -5.62 DPS, sim-verified) [vendor] |
| feet | Soulcrusher Greaves (240110) | Leonid Barthalomew the Revered [vendor] | 119.0 | yes | Soulcrusher Sabatons (240102, -2.09 DPS) [vendor]; Greaves of Withering Despair (22240, -2.69 DPS) [dungeon]; Bloodvine Boots (19684, -3.67 DPS, sim-verified) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (161.4 DPS) | yes | Wrath of Cenarius (21190, -2.75 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.32 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -3.54 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (161.4 DPS) | yes | Wrath of Cenarius (21190, -2.65 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.32 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -3.54 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (159.6 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.69 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (161.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.39 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.65 DPS, sim-verified) [quest] |
| main_hand | Fist of Cenarius (21188) | Champion's Battlegear [quest] | sim-verified (161.4 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Ironbark Staff (20220, -2.56 DPS) [rep]; Manual Crowd Pummeler (9449, -5.84 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Totem of the Storm (23199) (or Totem of Thunder (228176), Tidal Totem (272431), Totem of the Storm (272432), Burning Totem (272433), Totem of Urgency (279249), Totem of Ancestral Protection (249443), Kajaric Icon (206387), Polished Driftwood Icon (249398), Tempest Icon (206382), Dyadic Icon (206381), Galvanic Icon (206386), Sulfurous Icon (206388), Voltaic Icon (225838)) | World drop [world_drop] | 0.0 | yes | Totem of Thunder (228176, +0.00 DPS, sim-verified) [vendor]; Tidal Totem (272431, +0.00 DPS) [vendor]; Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Soulcrusher Headpiece; neck: Medallion of the Dawn; shoulder: Soulcrusher Mantle; back: Arcanoweave Cloak; chest: Soulcrusher Embrace; wrist: Rockfury Bracers; hands: Soulcrusher Mitts; waist: Soulcrusher Girdle; legs: Soulcrusher Legguards; feet: Soulcrusher Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Fist of Cenarius; ranged: Totem of the Storm

No-known-source sample (15 of 1242, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

