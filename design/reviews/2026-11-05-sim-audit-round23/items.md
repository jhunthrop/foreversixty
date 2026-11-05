# Round 23, items lane: on-use and proc effects and set bonuses on the published BiS

**CROSS-CLASS.** Both engine commits are shared item code (`sim/common`, `sim/core`). E 7b17db966 moves every mage set that wears Robe of the Archmage (all three mage specs, raid and bare 60). E 8376eb7ba moves every set that wears Ankh of Life and also uses a Demonic Rune, healthstone, mana gem or Thistle Tea (the raid 60 Restoration druid and Holy priest sets; the bare bands carry no rune).

Scope: every item in the frozen published BiS ($R/bis-frozen, 336 sets, 849 distinct items; 80 carry an ItemEffect row), ranked by how many sets wear it, and every set bonus those sets reach. Client: 1.60.1.70790 from wago.tools, shipped and `useHotfixes=1` (ItemEffect, ItemXItemEffect, SpellEffect, SpellAuraOptions, SpellProcsPerMinute, SpellCooldowns, SpellCategories, SpellCategory, SpellMisc, SpellDuration, Spell, ItemSet, ItemSetSpell), fetched 5 Nov ~13:00 UTC. Server text: nether.wowhead.com/forever/tooltip.

Method:
- **Diff against the last full audit.** Round 12 audited every BiS item effect on 70753. I diffed all 93 spells the BiS items reach (ItemEffect rows plus their trigger spells) between 70753 and 70790 (both with hotfixes): **no row changed**. So no new-data drift; anything wrong is an old reading.
- **New reading this round: the item's own cooldown row.** An item's ItemEffect row carries `CoolDownMSec`, `SpellCategoryID` and `CategoryCoolDownMSec`. When the row states a cooldown, it replaces the spell's own SpellCooldowns and SpellCategories rows (the server reads the item first and falls back to the spell only when the item states no cooldown). Earlier rounds read the spell rows for two items. Patch 1.4.0's note on Robe of the Archmage settles which wins (I1).
- **Handlers.** Every effect row was matched to its engine registration (`sim/common/item_effects.go`, `item_effects/client_procs.go`, `weapon_proc_rolls.go`, class relic files) and the amounts, durations and cooldowns read against the client rows.

## Items (ranked by sets wearing them)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| I1 | **Robe of the Archmage (12 sets: all three mage specs, raid and bare 60) was tied to the mana gems and Demonic Rune for a minute both ways.** Using the robe held the gems and rune 60 s; a gem or rune held the robe 60 s. The engine read spell 18385's category 30 (60 s). The robe's own ItemEffect row (99905) is a 5 min cooldown in category **0** with no category cooldown, and that row is what the server uses for an item. Blizzard said so for this item in patch 1.4.0: "No longer triggers the cooldowns of a Mage's mana gems. Mana gems are now on their own unique cooldown category." The gems' own rows are category 1153 for 2 min, which the engine's conjured timer already shares with the runes and healthstones. Now the robe has only its own 5 min cooldown. | ItemEffect 99905 (`300000, cat 0, 0`); mana gems 5513/5514/8007/8008 and Demonic Rune 12662: `cat 1153, 120000`. Patch 1.4.0 note: https://warcraft.wiki.gg/wiki/Robe_of_the_Archmage (read 5 Nov). Engine: `sim/common/item_effects.go` robe handler, `sim/core/character.go` RestoreCategoryRecovery (removed). | CONFIRMED (client row + Blizzard patch note) | old (Era's code held the conjured timer after the robe; eca6175f9, round 11, added the reverse hold from the spell row) | E 7b17db966 |
| I2 | **Ankh of Life (70 sets) did not share the 2 min conjured category.** Its ItemEffect row (97284) is a 5 min cooldown in category 1153 with a 2 min category cooldown: the same category and time as every healthstone, Demonic and Dark Rune, mana gem and Thistle Tea row. The engine gave it only its own timer, so a healer could use the Ankh and a Demonic Rune back to back. Now it reads and starts the conjured timer for 2 min. Only the raid 60 Restoration druid and Holy priest sets carry both (raid consumes include Demonic Rune); the other 68 sets have no rune, and DPS sets never use the Ankh (it fires on a target under 60% health). | ItemEffect 97284 (`300000, cat 1153, 120000`); healthstones 9421/19004..19013, runes 12662/20520, gems, Thistle Tea 7676 all `cat 1153, 120000`. The old comment at `item_effects.go` said 1153 "is not shared with anything this sim casts". | CONFIRMED (client rows; the same item-row rule as I1) | old (round 8 handler) | E 8376eb7ba |
| I3 | Frozen Heart (92) and Molten Heart (34) of the Mountain: 29 Frost/Shadow damage and 55 attack power and ranged attack power for 15 s, 5 min cooldown, category 1141 for 15 s (engine: shared offensive-trinket timer, 15 s). Biome doubling meets no published fight. The 13 sets wearing both pair them correctly; the Hauberks' Rage (1 set each) sits in the same 1141. | ItemEffect 213858/213860; tooltips 1249110, 1249113 | Aligned | - | - |
| I4 | Gloves of Holy Might (44): 112 Holy with variance 0.2, ranged table, 3x vs Undead, on melee hits at the unstated rate. Uther's Strength (40): 4% on hit, 216 absorb (not modelled; defensive). Rune of the Guard Captain (47), the Enriched/Darkspear Seals, Runes of Perfection and Duty: equip stats in the pipeline; the "In a Forest or Grassland area" riders are dropped by `equip.AREA_CONDITION` (checked on all ten 70790 rider rows). | client rows; `data/pipeline/simdb/equip.py` | Aligned | - | - |
| I5 | Thunderbrew's Boot Flask (33): 75 Fire a second for 5 s to all targets, 5 min, no category. Dormant Heart (13): +40 healing/+13 damage for 15 s on the four named heals. Manual Crowd Pummeler (8): 50% attack speed 30 s, 3 min, 3 charges. Teebu's, Fiery War Axe, Toxic Revenger, Electrified Dagger, Glacial Blade, Hammer of the Northern Wind, Coldrage Dagger, Black Menace, Claw of the Shadowmancer, Thornspike, Verigan's Fist, Gut Ripper, Bow of Searing Arrows, Aegis, Wall of the Dead, Cloak of Flames, Sporid Cape, Annihilator, Thrash Blade, the relics: amounts, schools, durations and cooldowns match the 70790 rows. Their chance-on-hit rates are the server's and unchanged (Arcanite Champion stays with platform). | client rows vs handlers | Aligned | - | - |
| I6 | Not modelled, no DPS: Archmagister's Faceted Pendant use (10% undead XP), Robe of the Void (heal pet), Orb of Soran'ruk (self heal over time), the PvP talismans and Arena Grand Master (absorbs), insignias, Horned Viking Helmet, Parachute Cloak, Truefaith Vestments (Fade cooldown), Lifeforce Dirk and Julie's Dagger (self heals), Staff of Westfall and Forsaken Greataxe (zone movement speed), The Unstoppable Force (stun), Satchel of Copper Bombs (thrown attacks only; no APL throws). | client rows | Aligned | - | - |

## Set bonuses

| # | finding | evidence | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | 29 set-bonus thresholds are reached by published sets (Defias Leather 2-5 on 16 low-band sets; Vestments of the Devout and Magister's Regalia 2-5; Ironfeather 2; Stormshroud 2-3; Devilsaur 2; Embrace of the Viper 2; Dreadmist 2-3; the PvP and Highlander sets; Valor 2; Deathbone 2). Every one matches the client ItemSet/ItemSetSpell rows: the engine's set table (`sim/core/client_sets_gen.go`) is generated from them, and the models read the proc rate, amount and school from the rows (Devious Strike 5% for 15/s over 5 s; Lightning 10% for 20 Nature with variance 0.5; Revitalize 4% for 30 energy; Sudden Insight 5% for 200 mana; Ironfeather +20 damage and healing; Devilsaur +2% hit). Defias 3-piece's Humanoid Slaying meets no published target. | clienthf ItemSet, ItemSetSpell; `sim/common/item_sets/*` | Aligned | - | - |

## Round 22 open items (no DPS)

- Thrown spread (I2 of round 22): Footman's Pocketknife's Wowhead Forever tooltip still reads 9 - 10 (5 Nov 13:45 UTC); the client hotfix row says spread 0.6. Still OPEN; an in-game tooltip settles it.
- Bomb satchels (I5): Satchel of Dark Iron Bombs still reads 2 - 39 on Wowhead; unchanged, OPEN.
- Era classic-db weapons (I7): Azuresong Mageblade 17103 and Claw of Chromaggus 18832 still 404 on Wowhead Forever. OPEN until raid content ships.

## Before/after

Raid 60, frozen BiS, guide raid builds (`ALLRT=1 BISDIR=$R/bis-frozen measure.sh`), 10000 iterations. Before: $R/sb-before (site 25432c0f1, fork bff826207). After: E 8376eb7ba on site wt/r23-items. DPS, or effective HPS for healers.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 |
|---|---|---|---|---|
| mage-arcane | 711.5 -> 710.8 (-0.10%) | -0.08% | -0.19% | -0.09% |
| mage-fire | 725.0 -> 724.4 (-0.08%) | -0.03% | -0.17% | +0.01% |
| mage-frost | 619.2 -> 619.8 (+0.10%) | +0.16% | +0.11% | +0.00% |
| druid-restoration (HPS) | 927.3 -> 918.6 (-0.94%) | -0.94% | -0.95% | -0.96% |
| priest-holy (HPS) | 771.2 -> 771.2 (+0.00%) | +0.00% | 788.0 -> 766.0 (-2.79%) | -2.67% |

The mage moves are inside error at every seed: the robe was rarely held back in a 3 minute fight. The healers lose the second use of the Ankh or the rune inside a 2 min window. The alliance Holy priest set does not wear the Ankh at raid 60; the horde set does. No other published spec wears the robe, and no other raid 60 set wears the Ankh beside a rune. Bare bands: no rune in the bare preset, so the Ankh change does not reach them; mage bare 60 wears the robe and gems and moves the same way as raid (not measured separately, within error at raid).

## Files to regenerate

- Fork golden `sim/mage/TestP1Mage.results` was regenerated in the same commit (E 7b17db966); it moves only with the robe's timing.
- Site: nothing in this lane. The BiS re-rank (queued on land/bis-final) and the ladder goldens pick up the mage change at the integrator's regen.

## Measure at refresh

- Mage BiS at 60: the robe is now worth more beside the gems; it is already picked in all six mage 60 sets.
- Restoration druid and Holy priest raid 60: the Ankh now competes with the Demonic Rune for the 2 min category.

## For other lanes

- Paladin: Sanctified Orb 20512's handler (`sim/paladin/items.go`) still holds the conjured timer a minute on use, on the same spell-category reading I1 removed for the robe. The 70790 client has no ItemXItemEffect row for 20512, so the item row cannot be read; it is in no published set. Check its tooltip before changing it.
- Mage: the mana gem and Demonic Rune keep sharing the 2 min conjured timer, which matches their rows.

## For the guide lane

- Done in this lane (S, same commit as the report): arcane.md:78, fire.md:81 and frost.md:75 said "using any of the three locks the other two out for a minute". They now say the robe is on its own 5 minute cooldown and shares nothing with the gems or the rune. `_prose.test.ts` passes. The mage lane may touch the same paragraphs; the sentence is self-contained.

## Still open

- Chance-on-hit rates of the BiS weapons and Gloves of Holy Might are the server's; no client field states them (unchanged since round 8).
- Logs: looked 5 Nov 13:49 UTC. api.foreversixty.gg/v1/reports/recent lists only the 20 Oct test report. A Forever mage log with the robe and a mana gem used inside 60 s would confirm I1 in play; a healer log with the Ankh and a rune inside 2 min would refute I2.
