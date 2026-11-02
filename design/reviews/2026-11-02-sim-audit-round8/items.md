# Sim audit round 8: items lane

New lane. It covers item data from the client to the engine: simdb.bin and the item files (`data/pipeline/normalize`, `data/pipeline/simdb`, `wowhead_items`, `classicdb_items`), the temporary enchants the presets use, and the engine's item effects.

**Method.** This was a measured diff of every item, not a sample.
- **Picks.** All 719 distinct items any published BiS band picks (`data/builds/1.60.1.70707/bis/*.json`, bands 20 to 60, both presets).
- **Simdb.** Regenerated locally from the layered 70707 raw tables (hotfix push in), then decoded.
- **Wowhead Forever.** Each pick was compared field by field with its tooltip (662 fetched 2026-11-01; the other 57 are Wowhead-supplement ids that Wowhead Forever 404s). The fields were stats, armor, weapon min, max, speed and school, required level, classes, unique and item set. All 314 client ranged, wand and thrown weapons were compared the same way, because round 4 checked melee only.
- **Effects.** Every ItemEffect row on a pick was checked against the engine's registrations, and so was every set bonus a band reaches.

Scripts are in `scratchpad/r8/items/scripts/` (`cmpwh.py`, `effects.py`, `simdbdiff.py`).

## Item data (pipeline)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | **Stat 45 gave no healing.** ItemSparse stat 45 (ITEM_MOD_SPELL_POWER) mapped to spell power only. The engine heals with the healing stat alone (`Spell.HealingPower`), so 1,569 client rows were worth nothing to a healer. The worst cases are Dreamweave Gloves 10019 and the Devout set's mixed rows. Wowhead's `splpwr` on a record with no `splheal` is the same stat and had the same gap. | Wowhead Forever shows stat 45 as "Increases damage **and healing** … by up to N" and adds it into the healing total next to stats 41 and 42. Atiesh 22631 has 45 = 120 and 41 = 180 and reads "healing 300, damage 120". Field Marshal's Satin Hood 231622 has 45 = 21 and 41 = 41 and reads "healing 62, damage 21". Earthen Silk Slippers 254013 has 45 = 18 and 42 = 6 and reads "healing 18, damage 24". In the comparator, 98 of the 662 picks disagreed on healing; 0 disagree now. Test `test_spell_power_heals.py` failed 5 of 7 first. | CONFIRMED-BUG | S 57648f5d3 |
| I2 | **Wowhead percentages were rounded to integers on the sim path.** Blackcrow 12651 (12 picks: rogues and Survival, raid 60) gave 1% hit for its 0.9%. Savage Gladiator Chain gave 2% crit for 1.5%. Might of Hakkar's 0.5% hit became 0 under banker's rounding. | `raw/wowhead/items.json` `hitpct` 0.9 / `critstrkpct` 1.5 / 0.5; `wowhead_items.planner_percent_stats` used `int(round())`. Test `test_a_fractional_percent_stat_reaches_the_engine_unrounded` failed first. | CONFIRMED-BUG (single source: these ids are only in Wowhead's planner feed) | S 946dde232 |
| I3 | **Stat 50 was read as mastery.** It is ITEM_MOD_EXTRA_ARMOR, and it was dropped. It is now StatBonusArmor in simdb only; the planner states armor once, from the curve. Tank survival only; no published number moves. | Wowhead Forever: Tree Bark Jacket 1486 "+60 Armor", Sergeant's Cloak 18427 "+40 Armor", and the Knight's and Warden's leathers. Test `test_item_stat_50_is_bonus_armor_for_the_engine_only` failed first. | CONFIRMED-BUG | S 946dde232 |
| I4 | Everything else on the 662 compared picks agrees: primary stats, armor, school power, resistances, mp5, hit/crit/defense/dodge, block value, spell penetration, weapon damage, speed and school, required level, set id, unique, classes. (The relics' "all classes" allowlist is narrowed by proficiency.) | `cmp2.txt`: the remaining rows are parser gaps ("+N Ranged Attack Power", "+N Block Value" beside the shield's own block, the Rune of the Guard Captain's terrain line). | Aligned | - |
| I5 | Ranged, wand and thrown weapons: 302 of 305 match Wowhead Forever exactly, school included. The 3 misses are the engineering Satchels of Bombs (thrown, QualityModifier −5.33 to −16.33): ours is 0-11 and Wowhead's is 2-10. Their only picks are Protection warrior ranged at 20 and 50, which never throws. | Curve: `ItemDamageThrown` 18 q2 8.257 gives avg 5.85 and spread 4.95. No single rule fits all three. | OPEN (no published number) | - |
| I6 | Red Dragonscale Leggings 252603 (Enhancement bare 60): Int 32 against the tooltip's 31. 45 × 7000/10000 is exactly 31.5. A float32 percent would give 31, but Wowhead's own feed says 32, and the other 32 exact halves all round up. | Tooltip 2026-11-01; `raw/wowhead/items.json`. | OPEN (one point, one item; a hotfix is likely) | - |
| I7 | The 57 Wowhead-supplement picks (the ItemSparse row is a server hotfix the client never got) carry the feed's stats one for one, except I1 and I2. Their ItemEffect links do exist in the client (Kindling Stave, Spire of Hakkar, Mark of the Chosen, Black Menace) and are handled in the engine below. | `pick_effects.json`. | Aligned after I1/I2 | - |

## Item effects (engine)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| E1 | **Manual Crowd Pummeler kept its haste up all fight.** It had a 30 sec cooldown on a 30 sec buff. Forever's is a 3 minute cooldown with 3 charges. It is a published main hand for Feral 30/40/50, Retribution 30/40, Enhancement 30 and Arms 30. | ItemEffect 98990: CoolDownMSec 180000, Charges 3. Wowhead Forever "(3 Min Cooldown) 3 Charges". Test `TestManualCrowdPummelerIsOnAThreeMinuteCooldown` failed (30s). | CONFIRMED-BUG | E febc3c25a |
| E2 | **Fiery War Axe** (Enhancement 40) rolled Era's 155-197 plus 8 a tick. The client is 231-293 plus 20 every 2 sec. | SpellEffect 18796: 262, variance 0.2386; aura 3 at 20, period 2000. Tooltip "231 to 293 … an additional 60 over 6 sec". Test failed first. | CONFIRMED-BUG | E febc3c25a |
| E3 | **Bow of Searing Arrows** (BM, MM 40) rolled 18-26; the client is 23-33. | SpellEffect 29638: 28, variance 0.3636; tooltip "23 to 33". | CONFIRMED-BUG | E febc3c25a |
| E4 | **Gut Ripper** (rogues, Survival, Fury 40/50) lacked its "100% increased damage to targets below 25% health". The roll itself was fixed by the physical lane (cherry-picked). | SpellEffect 18107 effects 1 and 2 (dummy 100, 25); tooltip. Test `TestGutRipperDoublesInTheExecutePhase`. | CONFIRMED-BUG | E febc3c25a |
| E5 | **Not implemented.** These picks had no engine effect at all:<br>- Black Menace 6831: 77 Shadow (rogues 40)<br>- Verigan's Fist 6953: 39-59 Holyfire (Retribution and Holy 20)<br>- Toxic Revenger 9453: 14 Nature / 3 sec / 12 sec (rogues 30)<br>- Gloves of Holy Might 867: 101-123 Holy, 3x vs Undead (Survival 40/50)<br>- Spire of Hakkar 10844: Blood Siphon, 28 / 2 sec / 8 sec (Balance, Arcane, Frost 50)<br>- Thunderbrew's Boot Flask 744: 75 Fire / sec / 5 sec, 5 min (Frost 50)<br>- Fire Hardened and Brutal Hauberk: +30 rage, 1 h (Protection warrior 30)<br>- Aegis of the Scarlet Commander 7726: 13-15 Holy on a block (tanks 40)<br>- Wall of the Dead 1979: +150 armor, 3% when struck (tanks 50)<br>- Sporid Cape 6629: 1 Nature, and Cloak of Flames 3475: 7 Fire, both when struck | The client rows are named per item in `client_procs.go`, each checked against Wowhead Forever's server tooltip (2026-11-02). A "chance on hit" has no rate in the client. The procs use `core.UnstatedProcsPerMinute` (1 a minute), the warlock lane's helper, which beta-log fits back. Spire of Hakkar uses `core.MakeHarmfulSpellProcAura`. Tests `TestItemProcsRollTheClientsDamage`, `TestClientPeriodicItemEffects`, `TestRageHauberksGrantThirtyRage` and others all failed on cd4dfdcc0. | CONFIRMED-BUG (rate: single source, the shared convention) | E febc3c25a |
| E6 | **The wizard oils gave spell damage only.** The client also gives healing (aura 135). No published set heals with a wizard oil. | SpellItemEnchantment 2623/2626/2627/2628 → 25109/25110/25111/25113: aura 13 and 135 at 8/16/24/36. Test `TestWizardOilsMatchTheClient`. | CONFIRMED-BUG (no published number) | E 78262bdb1 |
| E7 | Already right, checked against the client: Teebu's 140, Arcanite Champion (120 Str, 270-450), Robe of the Archmage (375-625, 5 min), both Hearts of the Mountain (29 / 55, 5 min, 15 sec category), Mark of the Chosen, Ankh of Life, Hand of Edward's Focus, and the relics (Idol of the Moon 33, Idol of the Dream +2 sec, Totem of Thunder 1%, Libram of Invocation 5%, Talons of Wrath 50% for 35 mana). | ItemEffect, SpellEffect and SpellAuraOptions rows; tooltips. | Aligned | - |
| E8 | **Overlaps with other lanes, cherry-picked and not redone.**<br>- The physical lane's client weapon-proc table, E 3fcf9342d: Nightblade 161-301, Hammer of the Northern Wind 40-60, Coldrage 1293790 34-50, Gut Ripper's roll, Claw of the Shadowmancer, Thornspike.<br>- The warlock lane's Kindling Stave Ignition and unstated-rate helpers, E fa9d1ca90 and 3026483e6.<br>Black Menace and Verigan's Fist joined that table. `TestItemProcsRollTheClientsDamage` re-checks the overlapping rolls. | | Aligned (owned elsewhere) | cherry-picks 6b7ee8b04, 9c1064ef7, 81539d6d7 |
| E9 | **No sim effect, by design.** None of these changes a published number:<br>- Lifeforce 10750, Uther's Light 11302, Julie's Dagger: self-heals<br>- the Satchels of Bombs: thrown, and no published spec throws<br>- Archmagister's Faceted Pendant: an undead experience buff<br>- Guardian Talisman, Ring of Saviors, Arena trinkets: damage taken<br>- slow fall, run speed, fishing, cooking, mechanic duration | ItemEffect rows. | Aligned | - |

## Set bonuses reached by published bands

Every set any band wears 2 or more of was checked against the engine's conformance golden:
- Gladiator, Deathbone Guardian, Stormshroud, Devilsaur, Ironfeather, Defias Leather, Scarlet Crusade
- Magister's, Devout, Shadowcraft, Elements, Lightforge, Valor
- the PvP rank sets 1618-1665

The flat bonuses are "declared, matches". The hand models are Deathbone Surge, Lightning, Revitalize, Devious Strike, the 5% mana procs and Rogue Armor Energize. They read their amounts, chances (10, 4, 5, 5, 15%) and recovery straight from the client rows. Rogue Armor Energize states ProcChance 100, so it takes the unstated rule. Aligned.

## Temporary enchants the presets use

- **Brilliant Wizard Oil:** 36 and 1% spell crit. It now also heals (E6).
- **Brilliant Mana Oil:** 15 mp5 and 30 healing.
- **Elemental Sharpening Stone:** 2% crit.
- **Dense Sharpening Stone:** +8 damage.

All four match SpellItemEnchantment 2628, 2629, 2506 and 1643 and their spells. Permanent enchants are not used by any published sim, so `enchants.json` was spot-checked only.

## Before/after

10,000 iterations, seeds 1 and 50001. Before is the round-7 regen simdb on engine cd4dfdcc0. After is this lane's regen simdb on `lane/r8-items`. Using the round-7 regen as the base keeps out round 7's held-item regen: Holy paladin bare 60 is +13.2% from that alone, and it is not this lane's. The after includes the cherry-picked physical and warlock commits. Bare preset below 60; seed 50001 is within 0.3 points unless shown.

| spec | band | alliance / horde, seed 1 | absolute, alliance seed 1 | cause |
|---|---|---|---|---|
| Arms | 30 | −23.1 / −23.0 | 73.2 → 56.3 | E1 |
| Enhancement | 30 | −19.3 / −19.2 | 85.3 → 68.8 | E1 |
| Retribution | 30 / 40 | −18.7 / −18.4; −17.8 / −17.5 | 88.6 → 72.0 | E1 |
| Feral | 30 / 40 / 50 | −14.4 / −14.6; −12.3 / −12.6; −11.6 / −11.9 | 98.8 → 84.6 | E1 |
| Retribution | 20 | +3.0 / 0.0 (+3.5 at 50001) | 42.9 → 44.2 | Verigan's Fist |
| Protection paladin, threat | 40 / 20 | +3.8 / +3.9; +1.6 / +1.5 | 191.9 → 199.2 | Aegis, Sporid |
| Enhancement | 40 | +2.5 / +2.4 | 122.4 → 125.4 | E2 |
| Survival | 40 / 50 | +2.3 / +2.3; +1.7 / 0.0 | 171.3 → 175.2 | Gut Ripper, Gloves |
| Rogues | 30 | +1.5 to +2.3 | Assassination 57.0 → 58.2 | Toxic Revenger, physical's Claw/Thornspike |
| Rogues | 40 | +0.6 to +1.5 | Combat 109.7 → 110.9 | Black Menace, Gut Ripper, Coldrage |
| Protection warrior, threat | 30 / 40 / 20 / 50 | +1.5 / +1.6; +1.6 / +1.7; +0.6 / +0.5; +0.1 / +0.1 | 155.4 → 157.8 | hauberks, Aegis |
| Frost / Arcane / Balance | 50 | +1.4 / +0.6; +1.0 / +1.0; +0.9 / +0.8 | 205.7 → 208.5 | Boot Flask, Blood Siphon |
| Fury | 40 / 50 | +0.6 / +0.6; +1.0 / 0.0 | 127.4 → 128.1 | Gut Ripper; physical's Hammer |
| Arms | 40 | +0.6 / +0.6 | 96.0 → 96.6 | physical's Nightblade |
| BM, MM | 40 | +0.3 | 190.8 → 191.4 | E3 |
| Bear, threat | 20 | +0.8 / +1.0 | 48.6 → 49.0 | Sporid Cape |
| Healers (I1) | 20 | +1.1 to +1.5, Holy paladin horde +2.1 | Resto shaman 50.2 → 50.8 | I1 |
| Restoration shaman | 50 | +2.4 / +2.4 | 225.5 → 230.9 | I1 |
| Holy paladin | 30 / 50 | +1.3 / +1.2; 0.0 / +1.4 | 61.0 → 61.8 | I1 |
| Restoration druid | 50 | +0.6 / +0.6 | 274.3 → 275.9 | I1 |

**Raid 60** (rogues, Survival, Arms, Fury, Retribution, Enhancement, Arcane, Affliction):
- rogues, Survival, Arms and Fury: −0.03 to −0.16 (I2, Blackcrow's 0.9% hit)
- Affliction and Arcane horde: +0.1 to +0.2 (the warlock lane's Ignition)
- every other combination: 0.00

All five healers are 0.00 at raid and bare 60. Their published sets were picked by a ranker that weighed stat-45 items at zero healing, so the full effect of I1 shows only after a re-rank.

## For other lanes

- **Ranker / integrator.** I1 revalues every stat-45 item for healers (1,569 client rows) and E1 removes the Manual Crowd Pummeler's free uptime. Healer picks at every band, and the MCP bands' main hands (Feral, Retribution, Enhancement, Arms 30 to 50), should be re-ranked at the owner's refresh. Measured here on the published sets only.
- **Integrator.** This lane carries cherry-picks of E 3fcf9342d (physical) and fa9d1ca90 and 3026483e6 (warlock), as 6b7ee8b04, 9c1064ef7 and 81539d6d7. Merge order should take those lanes' originals first.
- **Integrator.** The CharacterStats goldens for mage and Elemental gain healing from E6. Those are fork goldens, already regenerated. `TestRotationLadder` was not touched.
- **Casters / hybrids (OPEN, cross-cutting).** The engine has one Crit stat. The client keeps melee crit (aura 52, e.g. Elemental Sharpening Stone 22755) apart from spell crit (aura 57, e.g. Brilliant Wizard Oil). So a Retribution or Enhancement stone also crits their spells. Not changed here.

## Still open

1. I5, the Satchels' thrown damage. I6, Red Dragonscale Leggings' half point.
2. **Proc rates.** Every new proc in E5, and the physical table, uses the 1-a-minute convention. Beta logs with Black Menace, Verigan's Fist, Toxic Revenger, Gloves of Holy Might or Spire of Hakkar would settle each rate.
3. Thornspike registers only in the site's build. The fork's `db.json` lacks it (physical lane's guard), so its fork test is skipped.

Sources: wago.tools db2 1.60.1.70707, layered hotfix push (`data/builds/1.60.1.70707/raw`); nether.wowhead.com/forever/tooltip/item and /spell (read 2026-11-01 and 2026-11-02).
