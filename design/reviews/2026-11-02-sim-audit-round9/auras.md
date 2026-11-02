# Sim audit round 9: auras lane

Scope: every buff, debuff, totem, blessing, aura, consumable and racial the sim applies from outside a spec's own kit, at bands 20/30/40/50/60.

Sources:
- The 1.60.1.70707 client tables with the hotfix push layered in. SpellEffect, SpellMisc, SpellLevels, SpellAuraOptions, ItemEffect and SkillLineAbility match the sha256s in `data/builds/1.60.1.70707/hotfix-push.json` (SpellEffect `eaff7efc…`).
- Also ItemSparse, Item, ItemXItemEffect, SpellCooldowns and ChrRaceRacialAbility, and SpellLabel and SpellAuraRestrictions from wago.tools for 70707.
- Wowhead Forever's server tooltips (`nether.wowhead.com/forever/tooltip/...`), read 2026-11-02. A tooltip and the client row together count as one source.

Engine branch: `lane/r9-auras` (from 40021374a). Site branch: `wt/r9-auras`.

**Found: 11 CONFIRMED-BUG rows, all fixed.**
- At the level-60 raid preset, which moves published numbers:
  - the foods (C2)
  - the Blasted Lands buffs (Z1)
  - Survival's stones (S1)
- Raid-buffed /sim rows below 45 only: the dumplings' level gate (C1).
- Custom /sim panel only:
  - Thorns, Retribution Aura and Stoneskin ranks (R1 to R3)
  - the party Mana Tide (M1)
  - the raid Winter's Chill flag (W1, the integrator's candidate)
- Latent, no published spec: Command on guardians (K1).
- Still OPEN:
  - armor-debuff stacking (D1)
  - Forever's new elixirs and potions (C5)
- The crit lane's list of which crit table each outside source hits is under "For other lanes".

## Reactive raid buffs, Mana Tide and Winter's Chill

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | **Thorns did 18 Nature damage per hit at every level.**<br>- Ranks 1 to 6 do 3, 6, 9, 12, 15 and 18, at levels 6, 14, 24, 34, 44 and 54.<br>- The tristate's improved form added Era's Improved Thorns (+75%). Forever's druid tree has no such talent. | SpellEffect 467/782/1075/8914/9756/9910 (aura 15) and SpellLevels. Wowhead Forever 467: "causing 3 Nature damage"; 9910: "18". `talents/druid.json` has no talent naming Thorns. Engine before: `core/buffs.go:609`. Test first: `TestThornsFollowsTheClientRanks`. Site pin: `TestReactiveAuraRanksMatchTheClient`. | CONFIRMED-BUG (custom panel only) | E b7c81850b, S 3bbc74b29 |
| R2 | **Retribution Aura did 20 Holy damage per hit at every level.**<br>- Ranks 1 to 5 do 5, 8, 12, 16 and 20, at 16, 26, 36, 46 and 56.<br>- The improved form (+50%) is gone: no Improved Retribution Aura in Forever. | SpellEffect 7294/10298/10299/10300/10301 (aura 15). Wowhead Forever 10301: "Causes 20 Holy damage". Engine before: `core/buffs.go:573`. Test first: `TestRetributionAuraFollowsTheClientRanks`. | CONFIRMED-BUG (custom panel only) | E b7c81850b, S 3bbc74b29 |
| R3 | **Stoneskin Totem took 30 off every melee hit at every level, and the amount grew each iteration.**<br>- The client's Stoneskin auras reduce Physical damage taken by 4, 7, 11, 16, 22 and 30. The totem ranks are learned at 4, 14, 24, 34, 44 and 54.<br>- The aura's expire added the reduction again instead of removing it. Every aura expires at the end of an iteration, so iteration 2 took 60 a hit, iteration 3 took 90, and so on.<br>- Ranged hits count now too: aura 14 with misc 1 covers the whole Physical school.<br>- Guardian Totems (Elemental tier 1) exists in Forever and still adds 10% a point. | SpellEffect 8072/8156/8157/10403/10404/10405 (aura 14, misc 1). Wowhead Forever 10408: "reducing Physical damage taken by 30". `core/aura.go:547` (doneIteration deactivates every aura). Engine before: `core/buffs.go:566-568`. Tests first: `TestStoneskinFollowsTheClientRanks`, `TestStoneskinExpiresCleanly`. | CONFIRMED-BUG (custom panel only) | E b7c81850b, S 3bbc74b29 |
| M1 | **A party's Mana Tide Totem restored 290 mana a tick at every level.** The healers lane had not taken this row.<br>- The ticks restore 88, 133, 197 and 290 every 3 s, from the totem ranks learned at 1, 38, 48 and 58.<br>- The talent is Restoration tier 3 (level 25), so a party has no Mana Tide before 25.<br>- The shaman's own Mana Tide already took its rank by level. | SpellEffect 16191/17355/17360/1325257 (aura 24, period 3000); totems 16190/17354/17359/1325256. `talents/shaman.json`: Mana Tide Totem is tier 3. Engine before: `core/buffs.go:1242`. Test first: `TestPartyManaTideFollowsTheClientRanks`. | CONFIRMED-BUG (custom panel `mana_tide_totems` only) | E 23aee9172 |
| W1 | **The raid's Winter's Chill flag fed a Frost mage's Cone of Cold.** (Candidate from the integrator.)<br>- Forever's Winter's Chill (12579, aura 271) raises only the applying mage's own next Cone of Cold.<br>- The flag put five stacks on the boss under the same aura the mage's talent uses (`core.WintersChillAura`, label "Winter's Chill"), and Cone of Cold read and spent them.<br>- The flag now applies nothing; the mage stacks its own.<br>- No preset carries it. The published Frost rotation casts no Cone of Cold, so no published number moves: mage-frost with `-buffs winters_chill` is the same before and after. | SpellEffect 12579 eff0 aura 271, misc 16 (Frost), mask 512. `mage/talents.go:533`, `mage/cone_of_cold.go:134-153`. Engine before: `core/debuffs.go:51-65`. Test first: `TestRaidWintersChillGivesNoOneStacks`. TestP1Mage regenerated, because the engine's test raid sets the flag. | CONFIRMED-BUG (custom panel only) | E d1dc542c7 |

## Raid buffs and debuffs, re-verified by the raid's level

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | **A script checked every `BuffRanks` row in `sim/core`** against the layered SpellEffect and SpellLevels: the id, the level and the amount.<br>- Arcane Intellect, Blessing of Might, Blessing of Wisdom, Mark of the Wild (armor, stats, resistances), Battle Shout (per level, capped).<br>- Trueshot Aura: 15/22/30/40/50, and ranged attack power only (the client row is aura 124 alone).<br>- Heart of the Lion (280 at 60), Devotion Aura, Power Word: Fortitude 4 to 70, Divine Spirit 17 to 40 (none before 30), Blood Pact.<br>- Strength of Earth (aura spells: 7/14/25/42), Grace of Air, Mana Spring (aura spells: 4/6/8/10 every 2 s, so 25 mp5 at 60), Flametongue Totem.<br>- Sunder 90 to 450, Faerie Fire 175 to 505, Curse of Recklessness 175 to 505, Curse of the Elements 4 to 10% and 30 to 75, Hunter's Mark 26 to 143, Judgement of the Crusader 23 to 161.<br>- Judgement of Wisdom and of Light were checked on the spells their debuffs trigger.<br>- Every row matches. | `scripts/ranks.py` (lane scratch) against SpellEffect sha `eaff7efc…` | Aligned | - |
| B2 | Curse of Recklessness is 505 armor (aura 674) plus aura 92 (no fleeing), with no attack power, at ranks 14/28/42/56. | SpellEffect 11717; Wowhead Forever 11717 | Aligned | - |
| B3 | **Leader of the Pack and Moonkin Aura.**<br>- Leader of the Pack is Feral tier 4 (level 30). Moonkin Aura is Balance tier 6 (level 40). Both are 3%, aura 290.<br>- In this client aura 290 is "Critical Strike chance": melee, ranged and spell. Era's Leader of the Pack was melee only and its Moonkin Aura spell only.<br>- Once the crit lane splits the tables, each must write both (crit list below). | SpellEffect 24932/24907. Wowhead Forever 24907: "Increases Critical Strike chance by 3%". Talent text of 17007. | Aligned now; crit-lane note | - |
| B4 | Mana Spring under the raid preset still counts once: the shaman's own totem adds only its excess over the raid's. Restorative Totems (exists, 5% a point) and the set bonus scale only the shaman's own totem. | `shaman/water_totems.go:166-172`, `shaman/shaman.go:230` | Aligned | - |
| B5 | Judgement of Wisdom and Judgement of Light amounts and levels match their triggered spells.<br>- The 50% chance is the engine's constant. The client's debuffs carry ProcChance 100 and trigger dummy "Intermediate" spells (1826, 5373), so the client does not state it. | SpellAuraOptions 20355 (mask 139944), 20346 (mask 40). Wowhead Forever 20355: "have a chance of restoring 59". | Aligned (the chance is single-source, unchanged) | - |
| B6 | **The custom panel's improved forms of raid buffs.** Most belong to talents Forever's trees do not have:<br>- Improved Mark of the Wild, Power Word: Fortitude, Blessing of Might, Blessing of Wisdom, Devotion Aura, Battle Shout and Hunter's Mark are absent.<br>- Enhancing Totems is absent, so Strength of Earth's and Grace of Air's improved forms are too.<br>- Improved Imp (Blood Pact), Restorative Totems (Mana Spring) and Guardian Totems (Stoneskin) exist.<br>- None of these improved forms is in a preset. | `data/builds/1.60.1.70707/talents/*.json` talent names | For the page lane (see below) | - |
| D1 | **Do Sunder, Faerie Fire and Curse of Recklessness stack under aura 674?** (round 8's D5)<br>- All four armor debuffs, Expose Armor included, are aura 674 with misc 1 and separate SpellLabels.<br>- None has a SpellAuraRestrictions row, and no tooltip says they do not stack.<br>- The beta logs on foreverlogs.gg are levels 13 to 30. There is no level-60 log before launch. | SpellEffect, SpellLabel, SpellAuraRestrictions (70707). Tooltips 11597, 9907, 11717, 11198. | OPEN: a level-60 log comparing physical hits with and without each debuff | - |

## Consumables

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **Smoked Desert Dumplings (level 45) rode the Raid-buffed rows from level 1** for the strength users (Arms, Fury, Retribution, Enhancement, Feral) and the three tanks.<br>- The level gate joins a preset id to its client item by name. `food_smoked_desert_dumpling` is singular and the item (20452) is plural, so the join missed and nothing gated it.<br>- So bands 20, 30 and 40 got 20 Strength their characters cannot eat. This is a side effect of round 8's C2, which made it these specs' food.<br>- An explicit name map covers the id now. A new test fails when any carried consumable names no client item. | ItemSparse 20452 RequiredLevel 45. `sim/request/item_levels.go`. Before: `raid-consumes.json` listed the dumplings from `from_level: 1`. Tests first: `TestEveryCarriedConsumableNamesAClientItem` (failed for 8 specs) and a `TestItemLevelsReadTheClientsRequiredLevel` row. | CONFIRMED-BUG (Raid-buffed /sim below 45) | S 2b11bec66 |
| C2 | **The raid preset fed casters, agility users and the paladin tank weaker foods than Forever has.**<br>- Two stronger Well Feds were unmodelled:<br>&nbsp;&nbsp;- Baked Salmon 13935 (Lobster Stew 13933 is the same), level 45: 28 spell damage. Nightfin Soup gives 22.<br>&nbsp;&nbsp;- Poached Sunscale Salmon 13932, level 35 (Mightfish Steak 13934 is the same at 45): 40 attack power and 40 ranged attack power.<br>- The engine models both now (Food 14 and 15).<br>- Measured at band 60 raid, both factions, seeds 1 and 50001, 10,000 iterations:<br>&nbsp;&nbsp;- Damage casters: Baked Salmon is +0.5 to +0.7% over Nightfin.<br>&nbsp;&nbsp;- Rogues and hunters: Sunscale is +0.2 to +0.8% over Grilled Squid.<br>&nbsp;&nbsp;- Protection paladin: Baked Salmon is +0.7% threat over the dumplings. Its threat is Holy spell damage. Round 8 gave it the dumplings, and even Nightfin already beat them by 0.4%.<br>&nbsp;&nbsp;- Strength users and the warrior and bear tanks keep the dumplings. Sunscale is 0.1% behind for them.<br>- A slot now holds the first consumable the character can use. So Nightfin stands behind Baked Salmon until 45, and Sunscale behind the dumplings until 45. | Items 13932 to 13935 and their Nutritious Food spells (aura 227 of 40 or 28). They trigger Well Fed 1249519 (aura 99 and aura 124) and Well Fed 1249520 (aura 13, misc 126). Wowhead Forever: "gain 28 Spell Damage", "gain 40 Attack Power". Tests first: `TestFoodsMatchTheClient` (engine), `TestEachConsumableSlotCarriesTheFirstUsableEntry` and `TestLayerWithNoKitIsThePresetItself` (site). | CONFIRMED-BUG | E b731861fb, S b5cc8b1f6 |
| Z1 | **The Blasted Lands buffs were one choice for everyone.**<br>- The preset gave R.O.I.D.S. (25 Strength) to every melee spec and to the hunters, and nothing to casters, healers or tanks.<br>- The five buffs share one hour-long category (103), so a raider takes the one their stats want.<br>- Measured at band 60 raid, horde, 3,000 iterations, seeds 1 and 50001; the final picks were re-measured in the Before/after below:<br>&nbsp;&nbsp;- Hunters: Ground Scorpok Assay (25 Agility) is +2.0 to +2.3% over R.O.I.D.S. for Beast Mastery and Marksmanship, and +0.7% for Survival.<br>&nbsp;&nbsp;- Rogues +0.8 to +0.9% and Feral +0.5%, also with the Assay.<br>&nbsp;&nbsp;- Damage casters: Cerebral Cortex Compound (25 Intellect) +0.4 to +1.2% over nothing.<br>&nbsp;&nbsp;- Healers: Cortex for the Holy paladin (+1.5%) and the Restoration shaman (+2.7%). Gizzard Gum (25 Spirit) for the Restoration druid (+2.3%) and both priests (+1.4%).<br>&nbsp;&nbsp;- Tanks: R.O.I.D.S. gives +1.2 to +1.8% threat.<br>&nbsp;&nbsp;- Strength users keep R.O.I.D.S. The Assay is 0.3 to 0.7% behind for them.<br>- The preset now lists the spec's consumables first, then the class's, then the group's, and a slot keeps its first entry. So a rogue's pick beats the melee group's. | Items 8410/8411/8412/8423/8424: use spells 10667/10668/10669/10692/10693, aura 29 at 25 (Strength, Stamina, Agility, Intellect, Spirit), all SpellCategory 103 with a 3,600,000 ms category cooldown. Test first: the new cases in `TestEachConsumableSlotCarriesTheFirstUsableEntry`. They failed on the previous preset. | CONFIRMED-BUG | S efe053f06 |
| S1 | **Survival fought in melee with unstoned weapons.**<br>- Survival is melee in Forever and dual wields (`leveling.DualWieldSpecs`). The hunters' "ranged" group has no stones.<br>- It now carries an Elemental Sharpening Stone on the main hand (2% crit, level 50) and a Dense Sharpening Stone on the off hand (+8, level 35).<br>- That measured +2.4% (horde, 10,000 iterations: 767.4 -> 785.8 at seed 1). | Enchants 2506 (spell 22755, aura 52) and 1643. Test first: the Survival cases in `TestEachConsumableSlotCarriesTheFirstUsableEntry`. | CONFIRMED-BUG | S be3314ab5 |
| C3 | Round 8's foods and elixirs still read the client:<br>- Grilled Squid 1% (aura 290). Tender Wolf Steak and Hot Wolf Ribs Agility. Runn Tum 15 Intellect. The sagefish spell damage. Smoked Desert Dumplings 20 Strength. Nightfin 22. Sage's Tea 44 healing.<br>- Elixir of Greater Defense 450, Elixir of Fortitude 200, Greater Stoneshield 2000, Flask of the Titans 1200, Flask of Distilled Wisdom 2000.<br>- Juju Power 30. Juju Might 40 and 40 ranged. R.O.I.D.S. 25. Rumsey 15.<br>- Fire Power 10, Frost Power 15, Shadow Power 40, Greater Arcane 35.<br>- Mongoose 25 and 2%. Brilliant Wizard Oil 36 and 1%. Elixir of Sages 25 and 2%. | `tools/foods.py` and `tools/it.py` over the layered tables | Aligned | - |
| C4 | **Juju Power, Juju Might, R.O.I.D.S. and Demonic Rune have RequiredLevel 0.**<br>- Their spells' SpellLevel (65, 50, 36, 58) gates a buff cast on a low-level target, not using the item. So a level-20 character can use them, as twinks do.<br>- The rune's life cost cannot kill a low-level character in the sim: it fires only when 1500 mana is missing from the bar. | ItemSparse. SpellLevels 16323/16329/10667/16666. `core/consumes.go:1073-1077`. | Aligned | - |
| C5 | **Forever added elixirs and 30-second potions the engine does not model.** Every item and its recipe item is in the client:<br>- Elixir of the Grizzly 250351 (25 Strength, 2% crit). Ferocity 250350 (18 Strength, 18 Agility). Cunning 250328 (25 Agility, 25 Intellect).<br>- Elixir of the Owl 250337 (25 Intellect, 2% crit). The Whale 250348. Nature Power 250343 (40 Nature spell damage).<br>- Frenzy potions 250940 to 250943 (28 to 80 attack power for 30 s). Spellblasting potions 250934 to 250937 (16 to 47 spell damage for 30 s).<br>- What stops them going in:<br>&nbsp;&nbsp;- `loot.json` names no source for the recipes (250364 to 250387, 250981 to 250987), so their launch phase is unknown.<br>&nbsp;&nbsp;- Whether the new elixirs stack with Mongoose or Greater Arcane is unknown. | `scripts/consumables.py` over the layered Item, ItemSparse, ItemEffect and SpellEffect. Wowhead Forever 250351, 250343, 250337, 250943, 250937. | OPEN: the recipes' sources and phase, and a beta character sheet with Mongoose plus the Grizzly | - |

## Racials of the published races

The published races are Human, Dwarf, Night Elf and Gnome for the Alliance, and Orc, Troll, Tauren and Undead for the Horde (`bis/<spec>.json`). Forever's new combinations among them are the Orc mage, Dwarf shaman, Undead paladin, Gnome priest and Troll warlock.

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| K1 | **Orc's Command gave every pet 5% more damage.**<br>- The client gives it only to hunter pets (20576, ClassMask 4) and warlock demons (20575, ClassMask 256). The other classes' Command is the dummy 21563.<br>- So an orc warrior or shaman with an item guardian (Dragon's Call's whelp, the Core Hound set) gave it 5% it should not have.<br>- No published spec moves: the published orcs are mages, shamans and Arms warriors, and no orc BiS carries a guardian item. | SpellEffect and SkillLineAbility 20575, 20576, 21563. Engine before: `core/racials.go:247-253`. Test first: the shaman case in `TestOrcHasCommandAndNoSavage`. | CONFIRMED-BUG (latent) | E e56f5e29a |
| K2 | **Every combat racial of the eight published races matches the layered rows.** The hotfix push now carries the values the engine took from Blizzard's notes.<br>- **Human:** Ruthless 1% (aura 290, all crit). The Human Spirit 10% (the push changed the shipped 5).<br>- **Orc:** Blood Fury 10% attack power, ranged attack power and spell power (auras 166, 167 and 317), 15 s, 2 min cooldown. Savage has no SkillLineAbility row.<br>- **Dwarf:** Artisan's Precision 1% melee and ranged hit (aura 54) and 1% spell hit (aura 55), which covers the Dwarf shaman. Big Game Hunter 5%. Stoneform -10% Physical damage taken, 8 s, 3 min.<br>- **Night Elf:** Elune's Light 10% (aura 290), 15 s, 3 min. Quickness 1% dodge.<br>- **Gnome:** Expansive Mind 5% (20591, ClassMask 400, covers the priest). Eureka! per class, 8 s, 2 min, from the class masks. The priest's periodic +10% names Mind Flay but not Shadow Word: Pain, as the client mask states.<br>- **Troll:** Berserking 10% melee, ranged and spell haste for 15 s (the push changed 10 s), 3 min. Beast Slaying 5%.<br>- **Tauren:** Endurance 5% Health. The pushed row has no aura 163.<br>- **Undead:** Touch of the Grave 20% for warrior, paladin and rogue (so the Undead paladin), 30% for priest, mage and warlock. ProcCategoryRecovery 15000. A 5% maximum-Health drain. | `tools/sp.py` over SpellEffect, SpellMisc, SpellCooldowns, SpellAuraOptions and SkillLineAbility. Wowhead Forever 1260189 and 1260201: "(Proc chance: 20%, 15s cooldown)". | Aligned (apart from K1) | - |

## Before/after

**Band 60, raid preset.** The lane engine head plus the site branch, against 40021374a with origin/main (55a1212f7). Both factions, seeds 1 and 50001, 10,000 iterations. Absolute values are horde at seed 1. The metric is DPS, effective HPS for healers and threat per second for tanks.

| spec | horde s1 / s50001 | alliance s1 / s50001 | horde s1 absolute |
|---|---|---|---|
| druid-balance | +1.19 / +1.15 | +1.21 / +1.17 | 579.5 -> 586.4 |
| druid-feral-bear | +1.86 / +1.86 | +1.88 / +1.87 | 856.3 -> 872.2 |
| druid-feral | +0.39 / +0.43 | +0.29 / +0.29 | 608.9 -> 611.3 |
| druid-restoration | +2.30 / +2.41 | +2.30 / +2.52 | 777.3 -> 795.2 |
| hunter-beast-mastery | +2.27 / +2.25 | +2.25 / +2.25 | 661.8 -> 676.8 |
| hunter-marksmanship | +2.53 / +2.52 | +2.53 / +2.54 | 653.1 -> 669.6 |
| hunter-survival | +3.46 / +3.46 | +3.45 / +3.44 | 759.5 -> 785.8 |
| mage-arcane | +2.00 / +1.94 | +2.01 / +1.99 | 601.4 -> 613.4 |
| mage-fire | +1.42 / +1.35 | +1.66 / +1.59 | 670.8 -> 680.3 |
| mage-frost | +0.90 / +0.90 | +0.89 / +0.89 | 601.5 -> 606.9 |
| paladin-holy | +1.56 / +1.56 | +1.42 / +1.31 | 648.2 -> 658.3 |
| paladin-protection | +1.99 / +1.98 | +2.02 / +2.04 | 1018.2 -> 1038.5 |
| paladin-retribution | +0.00 / +0.00 | +0.00 / +0.00 | 658.8 -> 658.8 |
| priest-discipline | +1.44 / +1.34 | +1.31 / +1.34 | 620.0 -> 628.9 |
| priest-holy | +1.35 / +1.35 | +1.31 / +1.25 | 674.1 -> 683.2 |
| priest-shadow | +0.77 / +0.79 | +0.78 / +0.79 | 724.2 -> 729.8 |
| rogue-assassination | +1.36 / +1.35 | +1.39 / +1.37 | 609.3 -> 617.6 |
| rogue-combat | +1.35 / +1.35 | +1.35 / +1.35 | 606.0 -> 614.2 |
| rogue-subtlety | +1.60 / +1.60 | +1.60 / +1.60 | 532.6 -> 541.1 |
| shaman-elemental | +1.93 / +1.89 | +2.09 / +2.03 | 476.8 -> 486.0 |
| shaman-enhancement | +0.00 / +0.00 | +0.00 / +0.00 | 594.0 -> 594.0 |
| shaman-restoration | +2.63 / +2.67 | +2.70 / +2.67 | 838.9 -> 861.0 |
| warlock-affliction | +1.16 / +1.21 | +1.18 / +1.18 | 585.2 -> 592.0 |
| warlock-demonology | +1.18 / +1.22 | +1.36 / +1.29 | 654.9 -> 662.6 |
| warlock-destruction | +1.11 / +1.20 | +1.17 / +1.10 | 649.7 -> 656.9 |
| warrior-arms | +0.00 / +0.00 | +0.00 / +0.00 | 720.8 -> 720.8 |
| warrior-fury | +0.00 / +0.00 | +0.00 / +0.00 | 757.8 -> 757.8 |
| warrior-protection | +1.47 / +1.56 | +1.48 / +1.52 | 1036.9 -> 1052.1 |

What moved each group:
- Damage casters: Baked Salmon and Cerebral Cortex Compound.
- Rogues, hunters and Feral: Ground Scorpok Assay and Poached Sunscale Salmon. Survival also gets its stones.
- Healers: Cortex or Gizzard Gum.
- Tanks: R.O.I.D.S., and the paladin tank's Baked Salmon.
- Arms, Fury, Retribution and Enhancement do not move.

**Bare bands 20 to 50 (published numbers below 60):** no change. None of these commits touches a class kit. Spot check at bands 20 and 40 (horde, seed 1, 10,000 iterations), before and after identical: Holy priest 30.5 and 119.3, Protection warrior 73.4 and 371.4, Balance 25.8 and 156.1, Restoration shaman 46.5 and 144.4, Combat rogue 40.7 and 109.8.

**Raid-buffed /sim below 60:** C1 takes the dumplings' 20 Strength off the strength users and tanks at bands 20 to 40. They eat Poached Sunscale Salmon from 35 instead. No published number uses these rows.

## For other lanes

- **Crit lane: which table each outside crit source hits.** In the client, aura 52 is weapon crit, aura 57 spell crit, and aura 290 "Critical Strike chance" (melee, ranged and spell).
  - Elemental Sharpening Stone (enchant 2506, spell 22755, aura 52, 2%): melee, on the stoned weapon. The engine (`core/consumes.go:195-201`) adds unified Crit and takes ranged back off, so today it wrongly gives spell crit.
  - Brilliant Wizard Oil (25113 eff2, aura 57, 1%): spell only. The engine (`core/consumes.go:146`) gives unified Crit, so today it wrongly gives melee crit.
  - These are all crit, so after the split each must write both tables:
    - Mongoose (17538 eff1, aura 290, 2%)
    - Grilled Squid's Well Fed (1249523, aura 290, 1%)
    - Elixir of Sages (1250941 eff1, aura 290, 2%)
    - Leader of the Pack (24932, 3%) and Moonkin Aura (24907, 3%); Era had these as one table each
    - Ruthless (20597, 1%) and Elune's Light (1259799, 10%)
  - Artisan's Precision is aura 54 plus aura 55, both hit tables.
- **Integrator.**
  - Every lane commit that moves a published number is a site preset commit (S b5cc8b1f6, efe053f06, be3314ab5). They move almost every spec at 60 (Before/after).
  - The engine commits move no published number, but S b5cc8b1f6 needs E b731861fb (the new Food enum values) pinned with it.
  - `data/proto/common.proto` needs the two new Food values at the pin (fork `proto/common.proto`).
  - `sim/request/IDS.md` is regenerated against the lane engine.
  - Fork goldens regenerated: TestP1Mage (W1), and TestFeralTank, TestProtection and TestP1TankWarrior (R1 and R2, the test raid's improved Thorns and Retribution Aura), in E 738ea7d1b.
  - No `TestRotationLadder` golden was regenerated.
- **Platform lane.**
  - Add `winters_chill` to INERT_BUFFS (W1).
  - The custom panel could hide the improved forms whose talents Forever lacks (B6).
- **Logs owner.**
  - `logs/engine/mechanics/consumables/catalogue.json` and the coach notes still name Grilled Squid, Nightfin and R.O.I.D.S. as the preset picks.
  - The new picks: Poached Sunscale Salmon, Baked Salmon, Ground Scorpok Assay, Cerebral Cortex Compound and Gizzard Gum.
- **Ranker lane / data owner.** The C5 recipes have no source in `loot.json`. Their sources decide whether Forever's new elixirs and potions belong in the preset.

## Still open

1. D1: whether Forever's armor debuffs (aura 674) stack with each other.
2. C5: Forever's new elixirs and potions. The recipes' sources and phase are unknown, and so is the stacking rule.
3. Blessing of Kings against Heart of the Lion stacking (round 7's C2), unchanged.
4. The 50% Judgement of Wisdom and Light chance is server-side and single-source.
