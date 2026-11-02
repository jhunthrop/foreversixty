# Round 10 healer sweep: Restoration shaman, level 60 (raid and bare)

Engine: r10/healers/engine at 4d3129cae (read-only). Site: wt r10-healers. Client rows: r9/casters/raw (1.60.1.70707 plus the hotfix push), read through tools/sp.py and directly from the CSVs: SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellLevels, SpellAuraOptions, SpellClassOptions, TraitDefinitionEffectPoints/CurvePoint, ItemEffect/ItemXItemEffect, ItemSetSpell and SkillLineAbility. Sim item rows come from `protoc --decode` of data/builds/1.60.1.70707/simdb.bin. Wowhead Forever tooltips (974, item 279249) were fetched 2026-11-02. No sim runs.

Published 60 build, both bands and both factions: `0000000000000000-055000000000000000-5530003315513151`, tauren (horde) and dwarf (alliance). Talents: Thundering Strikes 5, Ancestral Knowledge 5, Improved Healing Wave 5, Tidal Focus 5, Mindfulness 3, Ancestral Healing 3, Healing Focus 3, Water Shield, Tidal Mastery 5, Restorative Totems 5, Mana Tide Totem, Healing Way 3, Nature's Swiftness, Purification 5, Riptide. **Ranged slot empty in all four 60 sets.**

## Findings

| # | finding | engine | client | class | size |
|---|---|---|---|---|---|
| R1 | **The engine cannot model the Restoration relics, so the 60 BiS sets carry no totem.** Totem of Urgency (279249, epic, BoE, Enchanting-crafted, req. 60) is in the sim item table, but no effect is registered for it. It is equippable with zero value, so the gear search never picks it. Tidal Totem (272431) has the same gap, but its vendor source is `opens: later`, so it is not a Phase 1 item. | `sim/shaman/items.go:26-35`: `relicClassMasks` maps only LB (0x1), CL (0x2) and FS (1<<28). The comment says healing families "cannot be registered". `relic_mods_auto_gen.go` already carries both rows. `items.go:116-124` registers Thunder and Burning only. simdb.bin has rows for 279249 and 272431. | ItemXItemEffect 168129 → ItemEffect 235973 (TriggerType 1, equip) → spell 1306448 "Lesser Healing Wave Crit Chance": e0 aura 107, misc 7 (crit chance), bp 4, EffectSpellClassMask_0 0x80 (Lesser Healing Wave 10468 is mask 0x80). Tidal Totem: 1291076, aura 108, misc 14, -5, mask 0x40. Wowhead Forever item 279249 confirms the 4% equip text. Source: loot.json `crafted:enchanting`, no `opens`. Availability rests on our loot data only (single source). | CONFIRMED-BUG (gap). Fix: add `{0x40 → ShamanSpellMaskHealingWave}` and `{0x80 → ShamanSpellMaskLesserHealingWave}` to `relicClassMasks` and register both items with `NewEquipModItemEffect`. | Lesser Healing Wave +4% crit gives about +2% on its own healing (crit ×1.5), plus a few more Water Shield globes. Lesser Healing Wave is the tank <30% and member <25% heal, so probably 10-20% of healing. Estimate +0.2-0.4% HPS, and the ranged slot becomes non-empty. Tidal Totem is worth more to the mana-bound bare set (mana lasts 73 s), but only later. |
| R2 | Hygiene: `purificationHealingModifier()` is dead code. Its comment ("for the Healing Stream Totem's heal") contradicts `water_totems.go:78-80` and the client: Purification's masks 0x1c0 and m2 0x10 do not include Healing Stream's 0x2000. | `talents_restoration.go:106-111`, no callers (grep). | 16178 e0/e1 masks as quoted. | CONFIRMED (dead code, no effect) | none |
| R3 | Earth Shield (974) is listed in trainables/shaman.json as `source: class_spell`, level 50. That source is the heuristic for spells with no SkillLineAbility row. The rotation's note says a Forever shaman cannot learn it, and the client agrees: there is no SLA row for 974. The Wowhead Forever tooltip shows only "Requires level 50" (a SpellLevels row) and 15% base mana, 6 charges, a heal of 150. | Not registered. The APL engineering note says it is unlearnable. | SkillLineAbility: no row for 974 (UA and Hydra Shot were learnable without one in 70009, which is why the heuristic exists). | OPEN. A level-50+ beta shaman's trainer list or spellbook settles it. | If learnable: a large change, a free tank heal on every hit taken. Today it is 0. |
| R4 (carry-over) | Purification ×1.10 and Healing Way ×1.25 on Healing Wave are added (×1.35). A multiplicative server would give ×1.375. Re-checked: both are aura 108, misc 0, and they share the op. | `talents_restoration.go:77-91` | 16178 e0 and 29206 e0 | OPEN (round 9 S3). An in-game Healing Wave with one talent dropped settles it. | about +0.5% HPS if multiplicative |
| R5 (carry-over) | Chain Heal jumps to the lowest health percent; a server core may pick the largest health deficit. A party-only jump rule is not modelled. | `chain_heal.go:78-91` | ImplicitTarget 45, EffectChainTargets 3, ChainAmplitude 0.5 (1064/10622/10623) | OPEN (round 9 S4/S5) | under 0.5% |

## Crit split (round 9) re-check: aligned

- Heals roll `spellTableCritRating` (Crit + SpellCrit) plus spell `BonusCritRating`: direct heals (`spell_result.go:277-281`) and heal-over-time ticks (`spell_outcome.go` `outcomeHealingCritTick`). Intellect → SpellCrit at 0.0169 a point at 60 (59.2 Int per 1%). Base spell crit 2.3% (Era value, unconfirmed for 1.60, as the table comment says).
- Thundering Strikes 16255: aura 290, curve 85275 gives 1-5 → `stats.Crit` (`talents.go:64`). It reaches heals. Aligned.
- Raid buffs: Leader of the Pack 24932 and Moonkin Aura 24907 are both aura 290, bp 3 (eff 65, party area aura). The engine adds `stats.Crit` for both (`core/buffs.go:231-239`), so heals get both. Aligned with the client. Note: both are party auras, and the preset assumes the healer shares a party with a moonkin and a feral druid; that is engine-wide.
- Elixir of Sages 1250941: aura 29 Spirit 25, aura 290 at 2 → `{Spirit 25, Crit 2%}` (`consumes.go:635`). Aligned.
- Every other proc that reads crit:
  - Water Shield's crit proc reads `OutcomeCrit` on `OnHealDealt` and `OnSpellHitDealt` and is table-agnostic.
  - Ancestral Healing is unmodelled; it is armor only (16177 aura 101, misc 1, curve 8/17/25), and the profile's hits ignore armor.
  - Elemental Devastation is not taken. Elemental Focus 16164 is now a level-1 SLA passive (8836), but its Clearcasting 16246 mask (m0 0x98100003, m1 0x1000) has no heal bits.
  - Nothing in this build is limited to one table.

## Aligned, spell by spell (re-derived from client rows)

- **Healing Wave 9 (10396)**, the 60 cast:
  - bp 1344, Variance 0.13251 → 1254.95-1433.05, matching `HealingWaveBaseDamage[9]`.
  - Growth 5.5 a level, SpellLevel 56, MaxLevel 61 (+22 at 60). Coefficient 0.857.
  - Mana 560, cast 3000 ms, GCD 1500 (cat 133). Single target: ChainTargets 0, so the stray ChainAmplitude 0.2 is inert.
  - Rank 10 (25357) is the AQ book rank; the rank rewrite drops it.
- **Lesser Healing Wave 6 (10468):**
  - bp 820, Variance 0.10909, growth 4.2, SpellLevel 60, MaxLevel 65. Coefficient 0.429.
  - Mana 380, cast 1500. Matches the hand override at `lesser_healing_wave.go:15-19` (the generated 27624 row is the BC id, correctly bypassed).
- **Chain Heal 3 (10623):**
  - bp 490, Variance 0.13220 → 457.61-522.39. Growth 3.3, SpellLevel 54, MaxLevel 59. Coefficient 0.714.
  - Mana 405, cast 2500. ChainTargets 3, amplitude 0.5 → 1, 0.5, 0.25 (`chain_heal.go:15-16,47-57`).
  - Riptide's bonus (1239243 e2 aura 4, bp 25) applies to the primary only, ×1.25. Each target rolls its own crit.
- **Riptide 3 (1239243):**
  - Direct heal: bp 841, Variance 0.1 → 798.95-883.05, growth 4.9, coefficient 0.214.
  - Heal over time: e1 aura 8, 161 a tick, RealPointsPerLevel 3.7 (0 at 60), period 3000, duration 15000 → 5 ticks, coefficient 0.1 (`riptide.go:15-35`).
  - CategoryRecoveryTime 6000 (category 1228, shared timer), mana 385, GCD 1500. Class mask m2 0x10.
  - Attributes_8 0x1200 includes PERIODIC_CAN_CRIT (0x200), and 408521/1239242/1239243 are in `periodic_crit_auto_gen.go:324,344,345`, so the ticks crit at the healer's spell crit plus Tidal Mastery (spell `BonusCritRating`).
  - Ticks go through `DealPeriodicHealing`, so they do not reach `OnHealDealt`. Water Shield's ProcTypeMask 0x362A8 has no DONE_PERIODIC (0x40000) bit, so crit ticks rightly give no globe.
  - Purification e1 (misc 22 DOT, m2 0x10, 10%) is matched by the additive caster multiplier on the snapshot; the net 10% is the same.
- **Healing Stream Totem 5 (10463 → heal 10461):**
  - Totem: 80 mana, GCD 1000, 300 s, mask 0x80000.
  - Heal: eff 35 aura 8, 11 every 2000 ms, coefficient 0.022, mask 0x2000, Attributes_8 0x1001200 (PERIODIC_CAN_CRIT) → `water_totems.go:92-95` rolls the shaman's healing crit and deals a periodic heal (no Water Shield globe, which is correct).
  - Restorative Totems e1 (aura 108, misc 22, mask 0x2000, curve 85319 → 50%) is applied as `DamageMultiplier` 1.5 on the whole tick, as a DOT mod applies to the final amount.
  - Tidal Mastery (0x5c0: Healing Wave, Lesser Healing Wave, Chain Heal and Lightning Shield; Lightning Shield and Rolling Thunder share bit 0x400, plus Riptide) and Purification (0x1c0 plus Riptide) do not include 0x2000, and the engine applies neither. Aligned.
  - That a totem's tick uses its owner's crit and healing is single-source (server-core behaviour). OPEN-low, no evidence against it.
- **Mana Tide Totem 4 (1325256 → 1325257):**
  - 60 mana, CategoryRecoveryTime 300000 (category 591), GCD 1000, duration 13000.
  - Restore: eff 35 aura 24, 290 every 3000 → 4 ticks to party mana users (`mana_tide_totem.go:15-19,39`). Learned at 58.
  - It holds the water slot, and dropping it replaces Healing Stream.
- **Mana Spring** (not cast by the list): rank 4 (10497 → 10491 family). Restorative Totems e0 (misc 8 ALL_EFFECTS, mask 0x4000 = Mana Spring, 25%) gives `restorativeManaSpringMultiplier`. 300 s, GCD 1 s.
- **Water Shield (408510 → 408511):**
  - ProcCharges 3, ProcCategoryRecovery 3500, proc 100, duration 600000, no cost, GCD 1500.
  - 408511 eff 137 gives 2% of maximum mana.
  - ProcTypeMask 0x362A8: taken melee, ranged and spell; dealt helpful and harmful magic; not periodic. The engine matches it (`water_shield.go:36-108`).
  - e1 (aura 4, bp 1, a recent row 1360851) is a dummy with no stated effect; no gap is evident.
  - The APL note was fixed since round 9 S2.
- **Nature's Swiftness 16188:**
  - RecoveryTime 180000, COOLDOWN_ON_EVENT (Attributes_0 0x2000000). Engine starts the cooldown on consume (`talents.go:369-410`).
  - Mask m0 0x9c3 (Lightning Bolt, Chain Lightning, Healing Wave, Lesser Healing Wave, Chain Heal, Ghost Wolf) plus m2 0x100000 (Healing Rain). The engine's Nature-with-cast-time rule covers every cast-time spell of this kit.
- **Talent curves** (TraitDefinitionEffectPoints → CurvePoint):
  - Improved Healing Wave: -100 to -500 (aura 107, misc 10, 0x40).
  - Tidal Focus: -1 to -5% cost (aura 108, misc 14, 0x1c0 plus Riptide; the totems are excluded), plus aura 54 and 55 hit 1-5 → `stats.Hit`.
  - Mindfulness 1223033: 17/33/50 (aura 134).
  - Healing Focus: aura 108, misc 9, 0x1c0 plus m2 0x100000 → pushback on Healing Wave, Lesser Healing Wave and Chain Heal.
  - Tidal Mastery: 1-5 (aura 107, misc 7).
  - Restorative Totems: 5-25 and 10-50.
  - Healing Way 29206: 8/17/25, now a flat aura 108 on 0x40 (proc 101, no stacking buff). FOREVER-CHANGE, matched.
  - Purification: 2-10 on both effects.
  - Ancestral Knowledge: 2-10% Intellect (aura 137, misc_1 8 = 1<<Intellect).
  - Thundering Strikes: 1-5 (aura 290).
- **Healing Rain** (415236/415242) and **Earthliving** (415292): their SkillLineAbility rows are AcquireMethod 3 (Season of Discovery runes), so they are not learnable. Correctly absent from the engine.

## Gear, racials, consumables: aligned

- **Sim item rows** (simdb decode): healing = Wowhead `splheal`+`splpwr` (or `splheal` alone when Wowhead lists `spldmg`), spell power = `splpwr`/`spldmg`:
  - Glowing Brightwood Staff 812: healing 150, spell power 50 (Wowhead `splheal` 150, `spldmg` 50).
  - Glowing Eye of Mordresh: 20/7.
  - Mooncloth Boots: 24/8; client ItemSparse stat 41/42 editors 10000/3333.
  - Living Crown: 73/25.
  - Devout Robe: 20/7.
  - Uther's Strength: 18/6.
  - Site-display note: items/shaman.json row 812 omits its healing and spell power (`stats_source` null). It is display only; the sim's row is right.
- **Devout set 182** (ItemSetSpell): 2 = 18679 (resist 8), 3 = 9346 (aura 13 and 135, 18 each), 4 = 27778 (2% when struck, absorb 488; noSim, it is a struck utility), 5 = 450576 (5% chance, ProcTypeMask 0x14000 → 450554, 200 mana; modelled on cast complete with the spell-damage and healing masks, so Healing Stream ticks, flagged NoOnCastComplete, are excluded), 6 = 18378 (8 mp5). Seven or six pieces worn, all thresholds met.
- **Trinkets:**
  - Shard of the Splithooves: 6 mp5.
  - Uther's Strength: 8397, 2% when struck in melee (ProcTypeMask 0x28) → absorb 216; it never procs on the healer.
  - Ankh of Life (bare horde): 1317424, bp 720, Variance 0.2 → 648-792, no coefficient, off the GCD, 5 min (`item_effects.go:3474-3512`).
- **Racials:** tauren Endurance 20550 (aura 133, 5% health), War Stomp and Cultivation; dwarf as in round 9 (Artisan's Precision hit, Stoneform 20594). None affects healing.
- **Raid consumables:**
  - Flask of Distilled Wisdom 17627: aura 418, 2000 mana.
  - Greater Mageblood Elixir 1250948: 20 mp5.
  - Greater Cleric's Elixir 1250926: 40 healing.
  - Elixir of Sages 1250941: Spirit 25, 2% crit.
  - Sage's Tea 1249925: aura 227 at 44 → Well Fed 1249927, aura 135 → 44 healing.
  - Brilliant Mana Oil: enchant 2629 → 25116, aura 85 at 15 and aura 135 at 30 → `{MP5 15, Healing 30}`.
  - Major Mana Potion 13444 and Demonic Rune 12662 through the shared potion and conjured timers.
  - Each matches `consumes.go`.
