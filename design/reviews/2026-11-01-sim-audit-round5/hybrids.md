# Sim audit round 5: hybrids (2026-11-01)

This report covers Druid (Balance, Feral cat, Feral bear, Restoration), Paladin (Holy, Protection, Retribution) and Shaman (Elemental, Enhancement, Restoration).

**Setup:**
- Client 1.60.1.70707.
- Engine: pin f8a5f94ce, plus fork branch `lane/r5-hybrids`.
  - The branch is rebuilt on the casters lane's racial commit (E e4d9ff60c, cherry-picked as e074b0415) and its Demonic Rune fix (de2ebaf86, here 78032c455).
  - It also carries the warlock lane's healing fix (fce7dd53d, here b03f7e479).
  - This lane had duplicated the racial fixes; that commit was dropped.
- Site: branch `wt/r5-hybrids` (this report only).

**New source this round: the server's hotfix push for build 70707.**
- wago.tools/hotfixes keeps every row Blizzard's server pushed to a 70707 client. The capture is from 2026-10-31 10:59 UTC.
- Dumped with `tools/hfdump.py` (lane scratch): 586 SpellEffect rows, 68 SpellAuraOptions, 41 SpellCooldowns, 1747 SpellMisc, plus SpellPower, SpellLevels, SkillLineAbility, ChrRaceRacialAbility, Trait* and CurvePoint rows.
- It was diffed against the shipped 70707 CSVs.
- Copies for every lane: `scratchpad/r5/hotfix70707/` (`diff-vs-base.tsv`, one line per changed, new or removed row, with spell names).
- The engine and the site's spell pipeline read only the shipped tables. Every druid, paladin and shaman row in the diff (67 rows that are not the blanket `Attributes_14` flag) was checked against the engine, and so were the racials.
- wowforevertalents.com's 31 October entry (the 30 October hotfix cache over 70685) was read on the integrator's request. Its paladin and druid lines agree with the wago push.

Fix-sha column: "E" is a fork commit on `lane/r5-hybrids`.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| X1 | **Assigned single-source rows, settled by the hotfix push.** Touch of the Grave's internal cooldown is 15 s (not 25). The Human Spirit is 10% (not 5). Endurance's 2% critical bonus is removed. Orc Savage is replaced by Command (pets +5%). Berserking lasts 15 s (not 10). Wind Blessed is 2% (not 1). | **Push rows:**<br>- SpellAuraOptions 243729/243733 ProcCategoryRecovery 25000 → 15000.<br>- SpellEffect 698126 5 → 10.<br>- SpellEffect 1322984 (aura 163) REMOVED.<br>- SkillLineAbility 11710 (Savage) REMOVED; 59880-59882 (Command, Orc race mask) added; ChrRaceRacialAbility 680 "Command".<br>- SpellMisc 329971 duration index 1 → 8.<br>- SpellEffect 1268898/1354094 3 → 2.<br>**Second source:** Wowhead Forever's live tooltips agree on every one, and so do the ChrRaceRacialAbility race-screen rows.<br>**Fix:** made in this lane first; on the integrator's request the casters lane's identical commit is kept. | CONFIRMED-BUG (two sources each) | E e074b0415 (casters' e4d9ff60c) |
| X2 | **Blood Fury: 2 or 3 minutes.** | The push has no SpellCooldowns 55612 row, so the client keeps 120000 ms. Wowhead's live text says 2 min.<br>The same 16 October notes changed War Stomp, and that change is in the shipped table (90 s). Only Blizzard's notes say 3 min. | Aligned: client table (2 min) | - |
| X3 | **Kings with Heart of the Lion.** | No new source. Moot for the preset and the page (round 4). | OPEN (carried) | - |
| X4 | Moonkin Aura and Leader of the Pack become all-crit (push: SpellEffect 701335 aura 57 → 290, 702727 aura 52 → 290). | The engine already grants unified `stats.Crit` for both (`sim/core/buffs.go:250`). | Aligned | - |
| X5 | **Integrator item: Benediction (18608) has no healing in simdb.** | The client agrees with simdb:<br>- ItemSparse has Int, Sta, Spi and shadow resistance, with no stat 41/42.<br>- ItemXItemEffect gives only 23041 (the Anathema use) and 23236 (2% Holy crit).<br>- Neither the push nor the seeded hotfix ItemSparse has an 18608 row.<br>Wowhead's "195 healing and 65 damage" has nothing in the client behind it. | Aligned (pipeline follows the client) | - |
| X6 | **Ice Barrier, for the casters lane.** | wago's 70707 push sets both effects of every rank (1364979/80 and the rest) from 3 to **5**. wowforevertalents' 30 October cache reads 15 and 5. The two hotfix sources disagree on the first effect. | OPEN, for the casters lane | - |
| X7 | Items: the push has 866 ItemSparse rows at 70707. The 478 that overlap our seeded `raw/hotfixes/ItemSparse.csv` match it, apart from float formatting. 388 are not in our file; one of them, Arena Master 18706, is in a band-60 BiS set (Protection paladin, horde). | Diff by value set, `hotfix70707/itemsparse.json`. | OPEN, for the ranker and data lane | - |

## Druid: Balance

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | Every Balance hotfix row checks out against the engine:<br>- Naturalist's mask moves (only Healing Touch and Wild Growth cast time; the engine follows the live text).<br>- Moonkin Aura goes to aura 290 (X4).<br>- Genesis 10% at 5/5 (CurvePoint rows; the engine reads 2% a rank). | Push diff; `sim/druid/talents.go:717`. | Aligned | - |
| B2 | Moves -4.6% (alliance) / -3.9% (horde) from the warlock lane's healing fix (items' healing no longer adds a hidden third as damage). | fce7dd53d. | (outside this lane) | b03f7e479 |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Shifting Power's 8 s cooldown, a hotfix (SpellCooldowns 101953 16000 → 8000), is in the engine. | `shifting_power.go:24`; Wowhead "8 sec cooldown". | Aligned | - |
| F2 | Ferocious Bite's new low rank (1328618, level 24) leaves the level-60 rank (31018, 2.7 an energy) as it was. Predatory Discipline is modelled. | Push SkillLineAbility 12256 and 59885; `ferocious_bite.go`. | Aligned | - |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| G1 | Primal Bite and Lacerate match the push:<br>- Primal Bite levels 40/50/60 for 42/62/77; 1238073's skill line row is removed and 1238070 sits at 60.<br>- Lacerate's new talent ranks 1328617/1328616 are 4 a tick and 5 stacks.<br>The engine already reads these from Wowhead. | Push SpellEffect 1071863, 1235070; SpellLevels 148268-148270; SkillLineAbility 53141-53143, 59883/59884. | Aligned | - |

## Druid: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Wild Growth gains a 3 s cast (SpellMisc 792796/7 cast time index 1 → 14). The engine already has 3 s, and the guide build does not cast it. | `wild_growth.go:20`. | Aligned | - |
| R2 | Bloom's rows (1328621/1328622) arrive in the push. They match the engine's 100% of Spirit and no cast procs. | Push SpellEffect 1367726-8. | Aligned | - |
| R3 | +1.0% from the casters lane's Demonic Rune health cost. | de2ebaf86. | (outside this lane) | 78032c455 |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | **Righteous Fury was +60% Holy threat only.** It is +33% all threat and +90% on Holy. | The shipped row (25780 effect 0) already reads 90. The push adds effect 4 (1367626: aura 10, +33%, every school). Wowhead and wowforevertalents read "Increases all threat generated by 33%, and all threat generated by your Holy attacks by an additional 90%".<br>The two aura 10 effects are multiplied (1.33 × 1.9), as the server's threat rule multiplies every aura 10. **Test first:** `TestRighteousFuryRaisesAllThreatAndHolyMore`. | CONFIRMED-BUG (fresh, two sources) | E f942519c3 |
| T2 | **Holy Shield's mana stayed at the old bracket** (150/195/240). | Push SpellPower 185892, 186383 and 184913 → 100/150/195. Wowhead's live tooltips and wowforevertalents agree. **Tests first:** `TestHolyShieldRanksCostTheHotfixedMana` and `TestHolyShieldCostsTheHotfixedMana`. | CONFIRMED-BUG (fresh, two sources) | E 7156faea2 |
| T3 | Swift Judgement (1 min, does not consume the Seal) and Sacred Duty (no Stamina) match the live text. Shieldsworn, Iron Creed 30%, Redoubt 30% and Shield Specialization 30%/100% were checked against wowforevertalents' rank lines. Judgement of Fury's taunt goes from 4 to 3 s, which has no threat effect. | `swift_judgement.go:18`; `tank_talents.go`. | Aligned | - |
| T4 | Undecided: the two Righteous Fury threat effects could add instead of multiply (2.23 against 2.527 on Holy, about -12% Holy threat). | Rests on the modern server's aura rule. | OPEN: a threat log | - |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| L1 | The Human Spirit 10% (X1): alliance +0.4%. | X1. | CONFIRMED-BUG (X1) | E e074b0415 |
| L2 | **Integrator item: Champion of the Light gave a healing paladin the share as healing too.** | The 70707 row 1311084 is aura 174 only (spell damage of an Intellect percent), with no 175. The push adds only the Hammer of Wrath cast-time effect (1367233). The live text names spell damage only. No published build takes it. **Test first:** `TestChampionOfTheLightGivesSpellDamageOnly`. | CONFIRMED-BUG (three sources, no published move) | E fb8acebd9 |
| L3 | Holy Power (50% critical healing bonus at 5/5), Divine Precision (8%), Boundless Light 15/30% and Reverence 30% match the push and the live text. Oath of Light's amounts (400/115/230) are new in the push. The engine still treats Oath of Light as unmodelled; Boundless Light covers allies without an Oath. | wowforevertalents 31 October; `healing_talents.go:41`. | Aligned (Oath unmodelled, as before) | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Touch of the Grave 15 s (X1): horde +0.7%. | X1. | CONFIRMED-BUG (X1) | E e074b0415 |
| P2 | **Assigned item 2, Unsealed Wrath on special attacks: new evidence, client table.** Talent 1310735's SpellAuraOptions ProcTypeMask is 4 (melee auto-attack only, no 0x10 melee ability). So white swings store and consume copies, and Holy Strike does neither, which is the engine's reading (`isSealSwing`). The push also removes the talent's -20% seal cost (1340538 → 0); the engine never modelled it. | SpellAuraOptions 247745; push SpellEffect 1340538. | Aligned (one source, client table) | - |
| P3 | Holy Strike's flat bonus under the weapon percent: no new evidence. The push touches 10333 only by float rounding. | Push diff. | OPEN | - |
| P4 | **Assigned item 4.** Seal twist: Seal of Command and Seal of Righteousness are both cast, with Twist of Light up 99.8%; the gate behaves.<br>Weapons on this engine, 10000 iterations, seeds 1 / 50001:<br>- Alliance: published Arcanite Champion 635.6 / 636.0; The Needler 566.4; Blade of Hanna 576.8.<br>- Horde: published The Needler (13060) 583.1 / 582.9; Arcanite Champion 650.2 / 651.6 (+11.5%).<br>The published horde pick is from a pin-640da215f ranking; the next bis.yml run should take Arcanite Champion. Grand Marshal's Glaive (PvP) sims 783 / 800 if the ranker counts it. | `-gear main_hand:<id>` runs. | Ranker lane (spec side sane) | - |

## Shaman: Elemental

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| E1 | Orc loses Savage (X1): horde -0.7%. With fce7dd53d: -3.4 / -4.1% in all. | X1, B2. | CONFIRMED-BUG (X1) | E e074b0415 |
| E2 | Every talent the build takes was checked against the live text and the engine constants:<br>- Lightning Overload 10%, half damage.<br>- Elemental Alacrity 0.5 s; Elemental Fury +100%.<br>- Concussion, Call of Flame, Convection.<br>- Totem of Thunder +1% Lightning Bolt crit (client relic row).<br>The push has only float rounding on Lava Burst. | `sim/shaman/talents.go`; `relic_mods_auto_gen.go`. | Aligned | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | Orc loses Savage (X1): horde -0.8%. | X1. | CONFIRMED-BUG (X1) | E e074b0415 |
| N2 | **Assigned item 4: the weapon is sane.** Forest Defender's Axe 618.4 / 634.1 and Arcanite Champion 619.4 / 634.6 tie within 0.2%; Fiery War Axe is -11%. | `-gear` runs. | Aligned | - |
| N3 | Mental Dexterity 100%, Mental Quickness 30%, Flurry 25%, Maelstrom Weapon and Rage of the Farseer match the live text. | `sim/shaman/talents.go`. | Aligned | - |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | **Assigned single-source row: Healing Stream's tick crit.** Every rank's heal (5672 to 10461) carries PERIODIC_CAN_CRIT (Attributes_8 0x200). The push leaves it unchanged, and no second source exists, so the client's table decides. **Test first:** `TestHealingStreamTicksCanCrit`. | SpellMisc 322278 and others. | CONFIRMED-BUG (single source, client rule) | E c40020435 |
| H2 | **Side effect of H1, caught in measurement:** each tick was dealt as a direct heal, so every crit fired Water Shield (+9.6% HPS from mana). The client's tick is a periodic heal aura, and Water Shield's proc mask (223904, 0x362A8) has no periodic bit. The tick now deals through `DealPeriodicHealing`. **Test first:** `TestHealingStreamCritsDoNotTriggerWaterShield` (562 mana in 20 s, then 0). | SpellAuraOptions 223904. | CONFIRMED-BUG (in this lane's own fix, before merge) | E 8d3575c5c |

## Heal profile: error and sensitivity (assigned item 3, for the page lane and the integrator)

Effective HPS, raid preset, band 60, 10000 iterations, lane engine (with fce7dd53d and de2ebaf86). Each cell is seed 1 / seed 50001. The ×1.5 column is the same profile with the tank hit at 1725 and the pulse at 450.

| Healer | Alliance ×1.65 | Alliance ×1.5 | Horde ×1.65 | Horde ×1.5 | ×1.5 → ×1.65 |
|---|---|---|---|---|---|
| Restoration druid | 741.5 / 741.6 | 731.1 / 731.3 | 739.2 / 739.7 | 729.9 / 729.8 | +1.4% |
| Restoration shaman | 640.0 / 640.3 | 617.4 / 617.6 | 638.3 / 638.9 | 616.3 / 616.3 | +3.6% |
| Holy priest | 636.4 / 636.9 | 657.1 / 656.8 | 615.0 / 615.5 | 631.8 / 632.1 | -3.1% / -2.7% |
| Holy paladin | 551.6 / 551.3 | 540.5 / 540.2 | 541.9 / 541.5 | 531.3 / 531.3 | +2.0% |
| Discipline | 536.0 / 536.2 | 537.4 / 537.7 | 545.7 / 545.9 | 549.5 / 549.5 | -0.3% / -0.7% |

- **Seed error:** at most 0.6 HPS between seeds, 0.1% or less for every healer.
- **Tuning sensitivity:** the profile choice moves healers by -3.1% to +3.6% relative to each other, about 30 times the seed error.
- **Ties under the page rule.** These pairs sit within the tuning sensitivity, and the order flips between ×1.5 and ×1.65:
  - Restoration shaman and Holy priest: alliance 640.0 against 636.4 at ×1.65, 617.4 against 657.1 at ×1.5.
  - Holy paladin and Discipline: horde 541.9 against 545.7.
- The druid's lead over every other healer exceeds both the seed error and the profile sensitivity.

## Before/after (relative %)

**Method:**
- Raid preset, band 60, 10000 iterations, seeds 1 / 50001.
- **Before:** pin f8a5f94ce with origin/main.
- **After:** lane head 9e64eb77b with `wt/r5-hybrids`.
- Healers are scored in effective HPS, tanks in threat a second.

| Spec | Changes | Alliance | Horde | Absolute, seed 1 (alliance; horde) |
|---|---|---|---|---|
| Balance | fce7dd53d | -4.6 / -4.5 | -3.9 / -3.9 | 524.9 → 501.0; 509.4 → 489.5 |
| Feral | - | 0.0 | 0.0 | 611.5; 613.1 |
| Feral bear (threat) | - | 0.0 | 0.0 | 679.7; 672.5 |
| Restoration druid | de2ebaf86 | +1.0 / +1.0 | +1.0 / +1.0 | 734.4 → 741.5; 732.0 → 739.2 |
| Holy paladin | X1 | +0.4 / +0.3 | 0.0 | 549.5 → 551.6; 541.9 |
| Protection paladin (threat) | T1, T2, X1 | +50.0 / +50.1 | +51.1 / +51.0 | 685.4 → 1028.1; 690.2 → 1043.2 |
| Retribution | X1 | 0.0 | +0.8 / +0.6 | 635.7; 578.6 → 583.1 |
| Elemental | X1, fce7dd53d | -3.4 / -3.4 | -4.1 / -4.1 | 464.7 → 449.1; 466.3 → 447.0 |
| Enhancement | X1, fce7dd53d | -0.5 / -0.5 | -1.4 / -1.4 | 618.4 → 615.0; 639.8 → 630.6 |
| Restoration shaman | H1, H2 | -0.1 / -0.1 | -0.2 / -0.1 | 640.9 → 640.0; 639.6 → 638.3 |

The lane's own commits move nothing outside the group: Righteous Fury, Holy Shield, Healing Stream and Champion of the Light are class-only. The moves outside the group come from the merged e4d9ff60c, de2ebaf86 and fce7dd53d, and those lanes report them.

Nothing nightly-owned was committed. `TestRotationLadder` was not regenerated. Fork goldens regenerated:
- TestProtection.
- TestElemental and TestEnhancement (on the merged base).
- The paladin conformance golden (Holy Shield's hotfixed mana against the shipped table).

## Still open

1. **T4.** Righteous Fury's two threat effects, multiplied or added. A Protection paladin threat log settles it.
2. **P3.** Holy Strike's flat bonus under the weapon percent. A level-60 Retribution log settles it.
3. **X3.** Kings with Heart of the Lion (moot for defaults).
4. **X6.** Ice Barrier's first effect: 5 (wago's 70707 push) or 15 (wowforevertalents). For the casters lane.
5. **X7.** 388 ItemSparse hotfix rows at 70707 are not in our seeded hotfix file. For the data and ranker lane.
6. **P4.** Horde Retribution's weapon: the next bis.yml run should pick Arcanite Champion.
7. **Pipeline (cross-cutting).** The shipped-table generators (spellconst, conformance, relics) do not read the server's hotfix push. Hand-written code covers the rows found here, but a `hotfix_merge` for spell tables would stop the next push from drifting silently.

## Sources

- wago.tools hotfix feed, build 1.60.1.70707, push of 2026-10-31 10:59 UTC (https://wago.tools/hotfixes, read 2026-11-01). Dump and diff in `scratchpad/r5/hotfix70707/`.
- wago.tools shipped 70707 tables (read 2026-11-01): SpellEffect, SpellAuraOptions, SpellCooldowns, SpellCategories, SpellMisc, SpellPower, SpellLevels, SkillLineAbility, ChrRaceRacialAbility, ItemSparse, ItemEffect, ItemXItemEffect, TraitDefinition, TraitDefinitionEffectPoints and CurvePoint.
- Wowhead Forever tooltips (read 2026-11-01): 20550, 20598, 20572, 20574, 20575, 20554, 1259710, 1260189, 1260201, 1322605, 20925, 20927, 20928, 25780, 1311084, 1310735, 407995, 1238069, 1238070, 1238073, and item 18608.
- wowforevertalents.com patch notes, 31 October 2026 entry (https://wowforevertalents.com/patch-notes, read 2026-11-01).
