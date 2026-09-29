# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 37.5. Weights run: 0.9s. Verify run: 0.9s. 252 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.128 ± 0.138), crit=1.335 ± 0.054, hit=3.164 ± 0.119, spell_haste=1.014 ± 0.120, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -0.77 DPS) [crafted]; Lucky Fishing Hat (19972, -0.77 DPS) [quest]; Shadow Goggles (4373, -1.47 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.5 | yes | Slime-encrusted Pads (6461, -0.71 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.72 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Caretaker's Cape (20428, -0.13 DPS) [rep]; Feyscale Cloak (6632, -0.41 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 | yes | Green Woolen Vest (2582, -0.21 DPS) [crafted]; Bloody Apron (6226, -0.21 DPS) [dungeon]; Gray Woolen Robe (2585, -1.09 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.6 | yes | Bright Bracers (3647, -0.02 DPS) [dungeon]; Seer's Cuffs (3645, -0.07 DPS) [dungeon]; Tabitha's Cuffs (251486, -0.60 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.16 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.34 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.61 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Keller's Girdle (2911, -0.45 DPS) [dungeon]; Novice Arcanist's Sash (253885, -0.87 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.0 | yes | Filigreed Pristine Leggings (253937, -0.42 DPS) [crafted]; Colorful Kilt (10048, -0.64 DPS) [crafted]; Silk-threaded Trousers (1929, -1.02 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 | yes | Red Woolen Boots (4313, -0.45 DPS) [crafted]; Pristine Boots (253889, -0.53 DPS) [crafted]; Feather Padded Treads (285345, -0.88 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.3 | yes | Sludge-Stained Band (286535, -0.29 DPS) [world]; Lavishly Jeweled Ring (1156, -0.58 DPS) [dungeon]; Black Pearl Ring (6332, -0.64 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon]; Sludge-Stained Band (286535, -0.57 DPS, sim-verified) [world]; Black Pearl Ring (6332, -0.61 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 1.3 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.07 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 5.0 | yes | Cookie's Stirring Rod (5198, -0.14 DPS, sim-verified) [dungeon]; Sable Wand (7607, -0.26 DPS) [quest]; Torchlight Wand (5240, -0.38 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 252, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 61.7. Weights run: 0.9s. Verify run: 1.0s. 475 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.065 ± 0.228), crit=2.229 ± 0.108, hit=3.372 ± 0.161, spell_haste=1.051 ± 0.234, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | Razorfen Downs: Withered Battle Boar [dungeon] | 11.0 | yes | Embalmed Shroud (7691, -0.46 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.46 DPS) [crafted]; Silk Headband (7050, -0.47 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.4 | yes | Crystal Starfire Medallion (5003, -1.08 DPS) [dungeon]; Pendant of Myzrael (4614, -1.12 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.25 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.6 | yes | Moonlit Amice (11884, -0.39 DPS) [quest]; Death Speaker Mantle (6685, -0.44 DPS) [dungeon]; Invoker's Mantle (215365, -0.91 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, +0.18 DPS, sim-verified) [crafted]; Prelacy Cape (7004, -0.15 DPS) [quest]; Caretaker's Cape (19533, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 9.8 | yes | Robes of Arcana (5770, -0.28 DPS) [crafted]; Death Speaker Robes (6682, -0.32 DPS) [dungeon]; Tree Bark Jacket (1486, -1.29 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -1.31 DPS) [quest]; Nightsky Wristbands (6407, -1.31 DPS, sim-verified) [dungeon]; Mindthrust Bracers (1974, -1.32 DPS) [dungeon] |
| hands | Serpent Gloves (5970) (or Shilly Mitts (9609)) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Shilly Mitts (9609, +0.42 DPS, sim-verified) [quest]; Gnoll Casting Gloves (892, -0.15 DPS) [world]; Truefaith Gloves (7049, -0.27 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.2 | yes | Belt of Arugal (6392, -0.03 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.48 DPS) [dungeon]; Invoker's Cord (215366, -0.59 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 | yes | Pristine Leggings (253987, -0.31 DPS) [crafted]; Silk-threaded Trousers (1929, -0.38 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.86 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.5 | yes | Nimbus Boots (6998, -0.22 DPS) [quest]; Acidic Walkers (9454, -0.29 DPS) [dungeon]; Spidersilk Boots (4320, -1.63 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.28 DPS) [quest]; Lorekeeper's Ring (20431, -0.30 DPS) [rep]; Electrocutioner Lagnut (9447, -0.61 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.46 DPS) [world]; Minor Channeling Ring (1449, -0.87 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 0.7 | yes | Twisted Chanter's Staff (890, +0.07 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -0.01 DPS) [quest]; Channeler's Staff (4437, -0.03 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | Gnomeregan: Dark Iron Agent [dungeon] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Sizzle Stick (8071, -0.02 DPS, sim-verified) [quest]; Spellcrafter Wand (6677, -0.15 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 475, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (gnome, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 90.9. Weights run: 0.9s. Verify run: 1.0s. 641 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.379 ± 0.386), crit=3.424 ± 0.185, hit=5.185 ± 0.341, spell_haste=2.354 ± 0.492, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -0.18 DPS, sim-verified) [world]; Holy Shroud (2721, -1.45 DPS) [dungeon]; Enchanter's Cowl (4322, -1.62 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 | yes | Necklace of Calisea (1714, -0.61 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.07 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 | yes | Green Silken Shoulders (7057, +0.04 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.07 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.9 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.13 DPS) [crafted]; Caretaker's Cape (19532, -0.27 DPS) [rep] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 24.3 | yes | Dreamweave Vest (10021, +0.52 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.83 DPS) [crafted]; Crimson Silk Vest (7058, -1.23 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.90 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.29 DPS) [quest]; Earthen Silk Cuffs (254019, -0.72 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.5 | yes | Red Mageweave Gloves (10018, -0.68 DPS) [crafted]; Gilded Handwraps (254021, -1.28 DPS) [crafted]; Black Mageweave Gloves (10003, -2.62 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.5 | yes | Deathmage Sash (10771, -0.41 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.49 DPS) [rep]; Star Belt (4329, -1.10 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.5 | yes | Crimson Silk Pantaloons (7062, -0.86 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.94 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.95 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.98 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.24 DPS) [crafted]; Acidic Walkers (9454, -2.31 DPS) [dungeon] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.29 DPS) [quest]; Lorekeeper's Ring (19525, -0.29 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.43 DPS) [vendor] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 | yes | Reedknot Ring (9622, -0.18 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.29 DPS) [vendor]; Minor Channeling Ring (1449, -0.32 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 50.2 | yes | Windweaver Staff (7757, -0.01 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -6.65 DPS) [dungeon]; Glimmering Staff (249392, -6.65 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Burning Sliver (5249) (or Twisted Nether Wand (249144)) | Crushridge Warmongers [quest] | 6.0 | yes | Twisted Nether Wand (249144, +0.56 DPS, sim-verified) [crafted]; Fizzle's Zippy Lighter (6729, -0.12 DPS) [quest]; Wand of Eventide (5214, -0.14 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Forlorn Spirits; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Burning Sliver

No-known-source sample (15 of 641, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (gnome, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 118.5. Weights run: 0.9s. Verify run: 1.0s. 848 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.825 ± 0.694), crit=5.572 ± 0.315, hit=7.662 ± 0.520, spell_haste=not significant (0.844 ± 0.797), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 55.5 | yes | Dreamweave Circlet (10041, -2.13 DPS) [crafted]; Bad Mojo Mask (9470, -2.25 DPS) [dungeon]; Eye of Theradras (17715, -2.38 DPS, sim-verified) [dungeon] |
| neck | Horizon Choker (13085) | Azuregos [world] | 25.5 | yes | Scorn's Icy Choker (23169, -0.97 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -1.13 DPS) [quest]; Darkspear Warding Pendant (272073, -1.20 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 45.8 | yes | Inquisitor's Shawl (19507, -1.98 DPS) [dungeon]; Green Silken Shoulders (7057, -2.33 DPS) [crafted]; Red Mageweave Shoulders (10029, -2.36 DPS, sim-verified) [crafted] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 25.5 | yes | Spritecaster Cape (11623, +0.43 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.26 DPS) [crafted]; Big Voodoo Cloak (8216, -0.54 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 55.5 | yes | Runecloth Tunic (13857, -2.41 DPS) [crafted]; Hibernal Robe (8113, -2.49 DPS) [dungeon]; Runecloth Robe (13858, -2.76 DPS, sim-verified) [crafted] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | 21.9 | yes | Bloodband Bracers (11469, -0.17 DPS, sim-verified) [quest]; Imperial Red Bracers (8247, -0.24 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.28 DPS) [crafted] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 29.2 | yes | Runecloth Gloves (13863, +0.13 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.44 DPS) [crafted]; Dreamweave Gloves (10019, -0.52 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 96.1 | yes | Dawnspire Cord (12466, +0.22 DPS, sim-verified) [world]; Deathmage Sash (10771, -8.09 DPS) [dungeon]; Satyrmane Sash (17755, -8.37 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 35.9 | yes | Crimson Silk Pantaloons (7062, -0.55 DPS) [crafted]; Imperial Red Pants (8251, -1.12 DPS) [dungeon]; Kilt of the Atal'ai Prophet (10807, -1.14 DPS, sim-verified) [dungeon] |
| feet | Southsea Mojo Boots (20641) | Southsea Shakedown [quest] | 28.1 | yes | Earthen Silk Slippers (254013, +0.63 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -0.56 DPS) [crafted]; Vinerot Sandals (17748, -0.74 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 76.6 | yes | Voodoo Band (1996, -8.37 DPS) [world]; Mindbender Loop (5009, -8.37 DPS) [dungeon]; Black Widow Band (6199, -8.37 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | Boulderfist Shaman [world] | 12.8 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world]; Mindbender Loop (5009, +0.00 DPS) [dungeon]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, -3.65 DPS, sim-verified) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Uther's Strength (11302, +0.14 DPS, sim-verified) [world]; Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 99.9 | yes | Illusionary Rod (7713, -2.02 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -2.87 DPS) [dungeon]; Spellshifter Rod (9527, -7.59 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Charged Lightning Rod (11860) | Ledger from Tanaris [quest] | 12.3 | yes | Cairnstone Sliver (9654, -0.42 DPS) [quest]; Lesser Eternal Wand (249232, -0.56 DPS) [crafted]; Fizzle's Zippy Lighter (6729, -2.26 DPS, sim-verified) [quest] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Acumen Robes; wrist: Shizzle's Nozzle Wiper; hands: Red Mageweave Gloves; waist: Highlander's Cloth Girdle; feet: Southsea Mojo Boots; finger1: Blackstone Ring; finger2: Ogremind Ring; trinket1: Darkspear Voodoo Seal; main_hand: Kindling Stave; ranged: Charged Lightning Rod

No-known-source sample (15 of 848, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 225.9. Weights run: 0.9s. Verify run: 1.0s. 1297 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.124 ± 1.054), crit=8.680 ± 0.479, hit=12.852 ± 0.837, spell_haste=not significant (1.136 ± 1.257), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 409.4 | yes | Bloodvine Goggles (19999, -6.38 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -15.41 DPS) [quest]; Field Marshal's Coronet (16441, -31.57 DPS) [vendor] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | 10.6 | yes | Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Onyxia Tooth Pendant (18404, -2.45 DPS, sim-verified) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 140.1 | yes | Lieutenant Commander's Silk Mantle (23319, -0.28 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.28 DPS) [pvp]; Champion's Silk Mantle (23264, -0.42 DPS, sim-verified) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 128.5 | yes | Chromatic Cloak (18509, -0.46 DPS, sim-verified) [crafted]; Drape of Vaulted Secrets (21415, -13.67 DPS) [quest]; Hide of the Wild (18510, -14.15 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 300.4 | yes | Bloodvine Vest (19682, -7.35 DPS, sim-verified) [crafted]; Enigma Robes (21343, -17.11 DPS) [quest]; Robe of the Archmage (14152, -17.16 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 155.5 | yes | Frostfire Bindings (22503, +1.57 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -16.55 DPS) [rep]; Dryad's Wrist Bindings (19596, -16.83 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 253.3 | yes | Sorcerer's Gloves (22066, -0.12 DPS, sim-verified) [quest]; Dreadmist Wraps (16705, -15.44 DPS) [dungeon]; Frostfire Gloves (22501, -26.84 DPS) [quest] |
| waist | Frostfire Belt (22502) | Frostfire Belt [quest] | 159.1 | yes | Highlander's Cloth Girdle (20047, -2.86 DPS) [rep]; Highlander's Cloth Girdle (20097, -3.50 DPS) [rep]; Belt of the Archmage (18405, -6.06 DPS, sim-verified) [crafted] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 177.7 | yes | Enigma Leggings (21346, -2.37 DPS) [quest]; Marshal's Silk Leggings (16442, -2.97 DPS) [vendor]; Bloodvine Leggings (19683, -8.24 DPS, sim-verified) [crafted] |
| feet | Enigma Boots (21344) | Enigma Boots [quest] | 158.4 | yes | Marshal's Silk Footwraps (16437, -0.89 DPS) [vendor]; General's Silk Boots (16539, -0.89 DPS) [vendor]; Frostfire Sandals (22500, -0.92 DPS, sim-verified) [quest] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 250.0 | yes | Band of Earthen Might (21182, -4.14 DPS, sim-verified) [quest]; Ritssyn's Ring of Chaos (21836, -12.93 DPS) [world]; Mindtear Band (20632, -13.21 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 166.3 | yes | Ritssyn's Ring of Chaos (21836, -2.47 DPS) [world]; Mindtear Band (20632, -2.75 DPS) [world]; Band of Earthen Might (21182, -5.86 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, -0.34 DPS, sim-verified) [world] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 | yes | Uther's Strength (11302, +0.12 DPS, sim-verified) [world]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 326.9 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Ironbark Staff (20069, -10.32 DPS) [rep]; Blade of Vaulted Secrets (21413, -24.53 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Burning Sliver (5249, -0.25 DPS) [quest]; Twisted Nether Wand (249144, -0.25 DPS) [crafted]; Dreambough Wand (249234, -4.22 DPS, sim-verified) [crafted] |

**New at 60:** head: Frostfire Circlet; neck: Jewel of Kajaro; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Frostfire Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; trinket1: Ankh of Life; trinket2: Thunderbrew's Boot Flask; main_hand: High Warlord's War Staff; ranged: Lesser Eternal Wand

No-known-source sample (15 of 1297, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 33.3. Weights run: 0.9s. Verify run: 1.0s. 247 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.128 ± 0.138), crit=1.335 ± 0.054, hit=3.164 ± 0.119, spell_haste=1.014 ± 0.120, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -0.77 DPS) [crafted]; Lucky Fishing Hat (19972, -0.77 DPS) [quest]; Shadow Goggles (4373, -0.93 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.5 | yes | Double-Stitched Woolen Shoulders (4314, -0.56 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.71 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Battle Healer's Cloak (20427, -0.13 DPS) [rep]; Feyscale Cloak (6632, -0.29 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.6 | yes | Green Woolen Vest (2582, -0.21 DPS) [crafted]; Bloody Apron (6226, -0.21 DPS) [dungeon]; Gray Woolen Robe (2585, -0.62 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 0.8 | yes | Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest]; Owlbeard Bracers (16981, -0.40 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.34 DPS) [crafted]; Apothecary Gloves (10919, -0.38 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.5 | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Keller's Girdle (2911, -0.45 DPS) [dungeon]; Novice Arcanist's Sash (253885, -0.47 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 10.0 | yes | Filigreed Pristine Leggings (253937, -0.42 DPS) [crafted]; Colorful Kilt (10048, -0.64 DPS) [crafted]; Silk-threaded Trousers (1929, -0.97 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 | yes | Red Woolen Boots (4313, -0.45 DPS) [crafted]; Pristine Boots (253889, -0.53 DPS) [crafted]; Feather Padded Treads (285345, -0.69 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon]; Black Pearl Ring (6332, -0.61 DPS) [world]; The 1 Ring (8350, -0.63 DPS) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, +0.25 DPS, sim-verified) [dungeon]; Black Pearl Ring (6332, -0.35 DPS) [world]; The 1 Ring (8350, -0.37 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | 1.3 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.03 DPS) [world]; Lesser Staff of the Spire (1300, -0.07 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 5.0 | yes | Cookie's Stirring Rod (5198, -0.21 DPS, sim-verified) [dungeon]; Torchlight Wand (5240, -0.38 DPS) [quest]; Greater Magic Wand (11288, -0.38 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 56.9. Weights run: 0.9s. Verify run: 1.0s. 470 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.065 ± 0.228), crit=2.229 ± 0.108, hit=3.372 ± 0.161, spell_haste=1.051 ± 0.234, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | Razorfen Downs: Withered Battle Boar [dungeon] | 11.0 | yes | Embalmed Shroud (7691, -0.46 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.46 DPS) [crafted]; Silk Headband (7050, -0.56 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.4 | yes | Darkspear Warding Pendant (272075, -1.07 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.08 DPS) [dungeon]; Pendant of Myzrael (4614, -1.12 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.6 | yes | Invoker's Mantle (215365, -0.34 DPS) [crafted]; Death Speaker Mantle (6685, -0.44 DPS) [dungeon]; Chestnut Mantle (17695, -1.24 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.30 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 9.8 | yes | High Robe of the Adjudicator (3461, -0.26 DPS) [quest]; Robes of Arcana (5770, -0.28 DPS) [crafted]; Tree Bark Jacket (1486, -1.15 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes | Owlbeard Bracers (16981, -1.25 DPS, sim-verified) [quest]; Nightsky Wristbands (6407, -1.31 DPS) [dungeon]; Tabitha's Cuffs (251486, -1.31 DPS) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 6.3 | yes | Gnoll Casting Gloves (892, -0.05 DPS) [world]; Truefaith Gloves (7049, -0.17 DPS) [crafted]; Serpent Gloves (5970, -0.75 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.2 | yes | Warsong Sash (16975, -0.17 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.30 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.48 DPS) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.5 | yes | Pristine Leggings (253987, -0.31 DPS) [crafted]; Silk-threaded Trousers (1929, -0.38 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.68 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.5 | yes | Acidic Walkers (9454, -0.29 DPS) [dungeon]; Boots of the Enchanter (4325, -0.37 DPS) [crafted]; Spidersilk Boots (4320, -1.58 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.30 DPS) [rep]; Electrocutioner Lagnut (9447, -0.61 DPS) [dungeon]; Sludge-Stained Band (286535, -0.61 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.46 DPS) [world]; Sacred Band (6669, -0.61 DPS) [quest]; Electrocutioner Lagnut (9447, -0.84 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 0.7 | yes | Twisted Chanter's Staff (890, +0.10 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -0.01 DPS) [quest]; Channeler's Staff (4437, -0.03 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | Gnomeregan: Dark Iron Agent [dungeon] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Cookie's Stirring Rod (5198, -0.30 DPS) [dungeon]; Sizzle Stick (8071, -0.34 DPS, sim-verified) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 470, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (troll, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 81.7. Weights run: 0.9s. Verify run: 1.0s. 636 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.379 ± 0.386), crit=3.424 ± 0.185, hit=5.185 ± 0.341, spell_haste=2.354 ± 0.492, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -0.21 DPS, sim-verified) [world]; Holy Shroud (2721, -1.45 DPS) [dungeon]; Enchanter's Cowl (4322, -1.62 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.3 | yes | Necklace of Calisea (1714, -0.69 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -0.96 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.07 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.4 | yes | Green Silken Shoulders (7057, +0.17 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.07 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.9 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.13 DPS) [crafted]; Battle Healer's Cloak (19528, -0.27 DPS) [rep] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 24.3 | yes | Dreamweave Vest (10021, +0.66 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.83 DPS) [crafted]; Crimson Silk Vest (7058, -1.23 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan [dungeon] | 9.0 | yes | Radiant Silver Bracers (4545, -0.10 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.29 DPS) [quest]; Earthen Silk Cuffs (254019, -0.72 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.5 | yes | Red Mageweave Gloves (10018, -0.68 DPS) [crafted]; Black Mageweave Gloves (10003, -1.11 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.28 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.5 | yes | Star Belt (4329, +0.04 DPS, sim-verified) [crafted]; Deathmage Sash (10771, -0.41 DPS) [dungeon]; Defiler's Cloth Girdle (20164, -0.49 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.5 | yes | Crimson Silk Pantaloons (7062, -0.77 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.94 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.95 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.74 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.24 DPS) [crafted]; Acidic Walkers (9454, -2.31 DPS) [dungeon] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.29 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.43 DPS) [vendor]; Advisor's Ring (20426, -0.58 DPS) [rep] |
| finger2 | Reedknot Ring (9622) | Jarl Needs a Blade [quest] | 7.0 | yes | Sea Giant's Toe Ring (274746, +0.78 DPS, sim-verified) [vendor]; Electrocutioner Lagnut (9447, -0.58 DPS) [dungeon]; Sludge-Stained Band (286535, -0.58 DPS) [world] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 50.2 | yes | Windweaver Staff (7757, +0.34 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -6.65 DPS) [dungeon]; Glimmering Staff (249392, -6.65 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 6.0 | yes | Wand of Eventide (5214, -0.14 DPS) [dungeon]; Sizzle Stick (8071, -0.14 DPS) [quest]; Fizzle's Zippy Lighter (6729, -0.91 DPS, sim-verified) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; finger2: Reedknot Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 636, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (troll, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 108.0. Weights run: 0.9s. Verify run: 1.0s. 843 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.825 ± 0.694), crit=5.572 ± 0.315, hit=7.662 ± 0.520, spell_haste=not significant (0.844 ± 0.797), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 55.5 | yes | Eye of Theradras (17715, -1.63 DPS, sim-verified) [dungeon]; Dreamweave Circlet (10041, -2.13 DPS) [crafted]; Bad Mojo Mask (9470, -2.25 DPS) [dungeon] |
| neck | Horizon Choker (13085) | Azuregos [world] | 25.5 | yes | Scorn's Icy Choker (23169, -1.06 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -1.13 DPS) [quest]; Darkspear Warding Pendant (272073, -1.20 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 45.8 | yes | Red Mageweave Shoulders (10029, -1.64 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -1.98 DPS) [dungeon]; Green Silken Shoulders (7057, -2.33 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 28.4 | yes | Spritecaster Cape (11623, -0.46 DPS) [dungeon]; Runecloth Cloak (13860, -0.63 DPS) [crafted]; Darkspear Raider's Cloak (272076, -1.78 DPS, sim-verified) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 55.5 | yes | Runecloth Robe (13858, -1.99 DPS, sim-verified) [crafted]; Runecloth Tunic (13857, -2.41 DPS) [crafted]; Hibernal Robe (8113, -2.49 DPS) [dungeon] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | 21.9 | yes | Bloodband Bracers (11469, +0.25 DPS, sim-verified) [quest]; Imperial Red Bracers (8247, -0.24 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.28 DPS) [crafted] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 29.2 | yes | Runecloth Gloves (13863, +0.68 DPS, sim-verified) [crafted]; Greenleaf Handwraps (19116, -0.29 DPS) [quest]; Stormcloth Gloves (10011, -0.44 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 96.1 | yes | Dawnspire Cord (12466, -0.22 DPS, sim-verified) [world]; Deathmage Sash (10771, -8.09 DPS) [dungeon]; Satyrmane Sash (17755, -8.37 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 35.9 | yes | Kilt of the Atal'ai Prophet (10807, -0.41 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -0.55 DPS) [crafted]; Imperial Red Pants (8251, -1.12 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Black Mageweave Boots (10026, -0.03 DPS) [crafted]; Vinerot Sandals (17748, -0.21 DPS) [dungeon]; Southsea Mojo Boots (20641, -1.06 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 76.6 | yes | Voodoo Band (1996, -8.37 DPS) [world]; Mindbender Loop (5009, -8.37 DPS) [dungeon]; Black Widow Band (6199, -8.37 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | Boulderfist Shaman [world] | 12.8 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world]; Mindbender Loop (5009, +0.00 DPS) [dungeon]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Tidal Charm (1404, -0.79 DPS) [vendor]; Guardian Talisman (1490, -0.79 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 99.9 | yes | Illusionary Rod (7713, -0.97 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -2.87 DPS) [dungeon]; Spellshifter Rod (9527, -7.59 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Charged Lightning Rod (11860) | Ledger from Tanaris [quest] | 12.3 | yes | Fizzle's Zippy Lighter (6729, -0.37 DPS) [quest]; Nature's Breath (19118, -0.44 DPS, sim-verified) [quest]; Lesser Eternal Wand (249232, -0.56 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Shizzle's Nozzle Wiper; hands: Red Mageweave Gloves; waist: Defiler's Cloth Girdle; finger1: Blackstone Ring; finger2: Ogremind Ring; trinket1: Darkspear Voodoo Seal; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Charged Lightning Rod

No-known-source sample (15 of 843, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (troll, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 208.0. Weights run: 0.9s. Verify run: 1.0s. 1292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.124 ± 1.054), crit=8.680 ± 0.479, hit=12.852 ± 0.837, spell_haste=not significant (1.136 ± 1.257), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 409.4 | yes | Bloodvine Goggles (19999, -9.53 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -15.41 DPS) [quest]; Field Marshal's Coronet (16441, -31.57 DPS) [vendor] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | 10.6 | yes | Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Onyxia Tooth Pendant (18404, -2.07 DPS, sim-verified) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 140.1 | yes | Champion's Silk Mantle (23264, -0.19 DPS, sim-verified) [vendor]; Lieutenant Commander's Silk Mantle (23319, -0.28 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.28 DPS) [pvp] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 128.5 | yes | Chromatic Cloak (18509, -0.30 DPS, sim-verified) [crafted]; Drape of Vaulted Secrets (21415, -13.67 DPS) [quest]; Hide of the Wild (18510, -14.15 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 300.4 | yes | Bloodvine Vest (19682, -6.21 DPS, sim-verified) [crafted]; Enigma Robes (21343, -17.11 DPS) [quest]; Robe of the Archmage (14152, -17.16 DPS) [crafted] |
| wrist | Frostfire Bindings (22503) | Frostfire Bindings [quest] | 28.9 | yes | Dryad's Wrist Bindings (19595, -0.73 DPS) [rep]; Dryad's Wrist Bindings (19596, -1.01 DPS) [rep]; Rockfury Bracers (21186, -4.40 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gloves (22066) | Just Compensation [quest] | 142.3 | yes | Dreadmist Wraps (16705, -1.58 DPS) [dungeon]; Gloves of Spell Mastery (14146, -3.23 DPS, sim-verified) [crafted]; Frostfire Gloves (22501, -12.98 DPS) [quest] |
| waist | Frostfire Belt (22502) | Frostfire Belt [quest] | 159.1 | yes | Defiler's Cloth Girdle (20163, -2.86 DPS) [rep]; Defiler's Cloth Girdle (20165, -3.50 DPS) [rep]; Belt of the Archmage (18405, -3.69 DPS, sim-verified) [crafted] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 177.7 | yes | Enigma Leggings (21346, -2.37 DPS) [quest]; Marshal's Silk Leggings (16442, -2.97 DPS) [vendor]; Bloodvine Leggings (19683, -8.93 DPS, sim-verified) [crafted] |
| feet | Enigma Boots (21344) | Enigma Boots [quest] | 158.4 | yes | Frostfire Sandals (22500, -0.56 DPS, sim-verified) [quest]; Marshal's Silk Footwraps (16437, -0.89 DPS) [vendor]; General's Silk Boots (16539, -0.89 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 250.0 | yes | Band of Earthen Might (21182, -3.49 DPS, sim-verified) [quest]; Ritssyn's Ring of Chaos (21836, -12.93 DPS) [world]; Mindtear Band (20632, -13.21 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 166.3 | yes | Ritssyn's Ring of Chaos (21836, -2.47 DPS) [world]; Mindtear Band (20632, -2.75 DPS) [world]; Band of Earthen Might (21182, -8.40 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 90.0 | yes | Uther's Strength (11302, -0.76 DPS, sim-verified) [world]; Tidal Charm (1404, -11.24 DPS) [vendor]; Guardian Talisman (1490, -11.24 DPS) [quest] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 326.9 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Ironbark Staff (20220, -10.32 DPS) [rep]; Blade of Vaulted Secrets (21413, -24.53 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Twisted Nether Wand (249144, -0.25 DPS) [crafted]; Charged Lightning Rod (11860, -0.31 DPS) [quest]; Dreambough Wand (249234, -1.15 DPS, sim-verified) [crafted] |

**New at 60:** head: Frostfire Circlet; neck: Jewel of Kajaro; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Frostfire Robe; wrist: Frostfire Bindings; hands: Sorcerer's Gloves; waist: Frostfire Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; trinket1: Ankh of Life; trinket2: Rune of the Guard Captain; main_hand: High Warlord's War Staff; ranged: Lesser Eternal Wand

No-known-source sample (15 of 1292, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

