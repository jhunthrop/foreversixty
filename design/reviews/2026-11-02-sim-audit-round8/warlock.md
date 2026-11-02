# Sim-accuracy round 8: warlock (2026-11-02)

Specs: Affliction, Demonology, Destruction and their demons. Client 1.60.1.70707 with the hotfix push layered in (the layered `raw/` fetched 2026-11-01 21:42 UTC; `python -m pipeline push-check` logic re-run 2026-11-02: wago's push totals equal the committed `hotfix-push.json`, so no newer push). Pin cd4dfdcc0, site origin/main 170b7b84c. Lane engine `lane/r8-warlock` (fork, head fa9d1ca90), site branch `wt/r8-warlock`.

Method: spec-breakdown, raid preset, band 60, 10000 iterations, seeds 1 and 50001, horde troll and alliance gnome; "before" is a detached origin/main on a detached cd4dfdcc0, each with its own `make simdb`. Client rows were read from the layered SpellEffect, SpellMisc, SpellAuraOptions, SpellClassOptions, SpellCooldowns, SpellPower, SpellLevels, SkillLineAbility, ItemEffect and ItemXItemEffect, with each talent's class masks decoded against the warlock family (SpellClassSet 5). The committed `spellconst/warlock.json` regenerates byte-for-byte from the layered tables, so the fork's conformance report (which reads that file) is current.

Totals: 5 CONFIRMED-BUGs, all fixed: two move no DPS (C1, T6); Kindling Stave's Ignition (C2) and Necropile Drain's rate and recovery (C6, C7) move DPS by +0.1 to +0.5% where the staff is worn. Ignition's rate is an estimate (C2). The spec-breakdown `-rotation` probe now works below 60 (S a0eb42beb).

## Assigned item 1: round 7's warlock fixes, side effects

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Mana potion now drinks on each cooldown: Major Mana Potion 1.99 a fight and Demonic Rune 2.00 in all three specs, about 1,798 and 1,195 mana each (client 17531 bp 1800; 16666 bp 1200 mana, 800 self-damage). Out of mana 0.0 s (Affliction) and 0.5 s (Demonology, Destruction). The rune and the potion are separate categories in Classic and both are used. | spec-breakdown resource table | Aligned | - |
| R2 | Elixir of Shadow Power (11474: aura 13, +40, school 32) and Elixir of Fire Power (7844: +10, school 4) now both on every warlock. Nothing else in the consume list moved; Flask of Supreme Power (17628, +150), Greater Arcane Elixir (17539, +35) and Brilliant Wizard Oil (enchant 2628, spell 25113: +36 damage, +36 healing, +1% crit) match the client. | client SpellEffect | Aligned | - |
| R3 | Life Tap: (420 + 1 a level above 56, so 424 at 60, + Spirit) x (1 + 10% a rank of Improved Life Tap), no spell power, as 11689's text ("Spirit increases the amount converted") and effect 0 (dummy 420, 1 a level) say. Spirit in `weight_stats` is still there and its test passes. | `lifetap.go`; client 11689 | Aligned | - |
| R4 | Held adoption rows. Affliction's Improved Bane of Agony (+0.7%) and Destruction's Immolate below Conflagrate (+0.1 to +0.2%) stay held. Ignition (C2) adds a flat Fire proc of about one a minute that no talent in either swap scales (Affliction has no Fire talent; Destruction's raid set does not wear the staff), so neither gap moves toward 1%. | - | Aligned (held) | - |

## Assigned item 2: ranks per band, coefficients, amounts

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| K1 | Every rank the engine registers at levels 10, 20, 30, 38, 40, 50 and 60 matches the client's learn level, mana cost, cast time, cooldown, GCD, duration and direct coefficient, talent-gated spells included (Conflagrate's Forever ranks 1293817 at 25 and 1293818 at 32, Incinerate 412758/1293812/1293813, Shadowburn, Siphon Life, Wrack, Bane of Havoc, Amplify Curse, Fel Domination). Only Death Coil's and Curse of Recklessness' control durations are unmodeled, which no DPS number reads. | fork `sim/core/testdata/conformance/warlock.golden.md` (passes on lane engine) against spellconst regenerated from the layered tables | Aligned | - |
| K2 | Damage tables: Shadow Bolt 1-10, Shadowburn, Conflagrate, Incinerate, Searing Pain, Soul Fire, Immolate (hit and tick), Rain of Fire, Corruption, Drain Soul, Bane of Agony, Bane of Doom, Drain Life, Siphon Life and Death Coil all match spellconst rank by rank (`TestWarlock*TablesMatchClient`). Spot checks on the periodic coefficients the golden does not print: Immolate 0.2 hit / 0.13 a tick, 5 ticks of 3 s; Corruption 0.2 a tick, 6 ticks of 3 s at ranks 3-6 (4 and 5 ticks at ranks 1 and 2, 12 and 15 s); Bane of Agony 0.133 a tick, 12 ticks of 2 s with the half/full/one-and-a-half ramp (total 12 x 46); Bane of Doom 1742 at 4.0 after 60 s; Siphon Life 41 + 0.05 a tick, 10 ticks; Wrack 36 + 0.143 a second for 6 s. | client 11668, 11672, 11713, 603, 18881, 1316697 | Aligned | - |
| K3 | Lash of Pain rank 6 (11780) 50 + 0.429, 160 mana, 12 s category cooldown, learned at 60 (11779 at 52); Improved Sayaad 10% a rank (18754 bp 30 at 3/3). | client 11780, 11779, 7814, 18754 | Aligned | - |
| K4 | Live casts at bands 20, 30, 40, 50 (bare preset) are the top trainable rank at each band for every spell all three rotations cast (for example band 40: Shadow Bolt 7641, Corruption 7648, Bane of Agony 11711, Immolate 11665, Siphon Life 18879, Conflagrate 17962, Incinerate 412758, Shadowburn 18869, Lash of Pain 7816, Life Tap 11687). Shadow Bolt 25307, Corruption 25311 and Immolate 25309 stay off as Ahn'Qiraj book ranks (book-ranks.json; ledger, blocked on logs). | spec-breakdown `-band`; `trainables/warlock.json` | Aligned | - |

## Assigned item 3: talents (client masks decoded)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Ruin (17959, crit bonus +100% at 5/5) and Cataclysm (17778, cost -10%) masks cover Shadow Bolt, Searing Pain, Soul Fire, Shadowburn, Immolate, Conflagrate, Incinerate, Rain of Fire, Hellfire (and Bane of Havoc for Cataclysm): exactly the spells the engine flags Destruction. | masks of 17959, 17778 | Aligned | - |
| T2 | Agonizing Flames (17927): +10% crit on Searing Pain; +10% direct damage on the Destruction list above, Shadow Bolt included; +10% periodic on Immolate and Rain of Fire. The engine's spell-level additive covers both halves. | 17927 effects 0-2 | Aligned | - |
| T3 | Shadow Mastery (18271): +2% a rank, direct mask Shadow Bolt, Corruption, Death Coil, Shadowburn, Drain Mana; periodic mask the banes, Corruption, Drain Life, Drain Soul, Siphon Life, Wrack. Engine list matches. Malevolence (1310949): +1% crit a rank on the warlock's Shadow list and Lash of Pain (owner spell mods reach demon spells). Pandemic (427712) on Corruption, banes, drains, Siphon Life, Wrack. | masks of 18271, 1310949 | Aligned | - |
| T4 | Improved Shadow Bolt: Shadow Vulnerability 17794, aura 270 on the caster, Shadow, 4% a rank, 12 s (DurationIndex 29), no SpellAuraOptions row so no charges. Nightfall 2% a rank on Corruption, Drain Soul, Drain Life, Wrack; Shadow Trance 17941 is -100% cast time on Shadow Bolt, 10 s, consumed by the instant bolt. Bane -0.5 s (Shadow Bolt, Immolate, Incinerate) and -2 s Soul Fire at 5/5. Improved Corruption -2 s and +10% periodic. | 17794, 18094, 17941, 17788, 17810 | Aligned | - |
| T5 | Demonology: Burning Shadow 18789 (Imp sacrifice) +15% Shadow; Touch of Fire 18791 (Succubus) +15% Fire, as Forever's text swaps them. Master Demonologist 23759/23761 +10% Fire (Imp) and Shadow (Succubus) on both units; Soul Link 25228 +3% all schools and 30% split; Demonic Knowledge +100% of level; Unholy Power and Fel Vitality 15% at max; Decimation -90% Soul Fire cooldown, -40% cast in the 10 s window, +6% Shadow Bolt and Searing Pain below 35%; Master Summoner -4 s, -40%. | 18789, 18791, 23759, 23761, 25228, 412732, 18769, 18731, 440870 | Aligned | - |
| T6 | **Demonic Sacrifice's Voidwalker and Felhunter buffs had each other's ids.** The amounts were right (Voidwalker 2% Mana, Felhunter 3% Health every 4 s, Forever's text), but the Voidwalker buff carried 18790 "Fel Stamina" and the Felhunter buff 18792 "Fel Energy". The client's 18792 Fel Energy is aura 21 (power) 2 every 4000 ms and 18790 Fel Stamina is aura 20 (health) 3 every 4000 ms. No published build sacrifices either demon; the buffs tab names follow the client now. Test `TestDemonicSacrificeBuffsCarryTheClientIDs` (failed, then passed). | client 18790, 18792; talent text 18788 | CONFIRMED-BUG | E c17f0c0ae |
| T7 | Destruction: Aftermath +10% Immolate hit a rank (18119 bp 50 at 5/5), Fire and Brimstone +25% Conflagrate crit (412751), Shadow and Flame +10% Shadow after Conflagrate and +10% Fire after Shadowburn for 20 s (Flame 426311, DurationIndex 18), Conflagrate keeps Immolate at 5/5; Incinerate +25% with Immolate up (effect 1 dummy 25). | 18119, 412751, 426316, 426311, 412758 | Aligned | - |

## Cross-cutting (fresh re-audit)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **Auras that faded before the pull showed negative uptime, and auras across the pull lost their prepull share.** The engine added "fade time minus max(start, 0)" to an aura's uptime, so an aura that faded at -45 s added -45 s. Demonology's prepull summons the Imp, sacrifices it at -45 s and summons the Succubus, so its result read Master Demonologist (Imp) at -25.0% and Master Demonologist (Succubus), up the whole fight, at 66.6%; Burning Shadow read -21.1% at bands 30 and 40. The /sim buffs tab shows these numbers (`sim/adapter` UptimeMS). Uptime now runs from the later of application and the pull to the later of fade and the pull: 0.0% and 99.9%. Core change, so every class with a prepull aura that fades reads right too. No damage changes. Test `TestAuraUptimeCountsOnlyTheFight` (failed at -45 s, then passed). | `sim/core/aura.go:674`; spec-breakdown aura table | CONFIRMED-BUG | E ffbb59f33 |
| C2 | **Kindling Stave's Ignition was not modeled.** Forever binds Kindling Stave (11750: 15 published warlock sets, Arcane mage 60 raid and bare horde, Fire mage and Elemental 50 bare) to a new equip spell, 1300781 Ignition: aura 42 on "deal harmful spell" (ProcTypeMask 0x10000), triggering 1300782 for 70 Fire at Variance 0.3 ("60 to 81", Wowhead Forever). The rows hold no rate: SpellAuraOptions 246565 is ProcChance 100, ProcCharges 0, SpellProcsPerMinuteID 0, ProcCategoryRecovery 10 ms; there is no SpellCooldowns or SpellCategories row for 1300781 (1300782 has only DefenseType 1), and the hotfix push leaves every one of these rows as shipped. The text says "have a chance". The client's 100 is not the rate. Forever writes ProcChance 100 for always-on procs worded without "chance" (Sear 1291568: "Harmful spell casts sear the target") and for "chance" ones alike. The beta logs show the "chance" ones proccing rarely: Dream Venom (1291807, set 162 5-piece; ProcChance 100, no PPM row, no recovery, "a chance") applied 10 times in Vijal's 358 auto attacks, and Rotmending (1321587, set 2133 5-piece; ProcChance 100, 2 s recovery, "a chance") 10 times in Altu's about 258 heals. Both come to about one proc a minute (2.8% a 2 to 2.7 s swing; 3.9% a cast of about 2 s). That is also the rate the engine already reads for every unstated "chance on" bonus (Bloodmail, the rage and energy bonuses). So Ignition fires at one a minute: a harmful cast rolls 1 x its cast time / 60, with an instant counted as the 1.5 s GCD. The new shared helper is `core.MakeHarmfulSpellProcAura`, with the rate as `core.UnstatedProcsPerMinute`. The estimate rests on two 10-proc samples from other items, so a level-48+ log with the staff would refine it; a 100% probe (branch `lane/r8-warlock-ignition-probe`, not merged) gave +5% at 60 and +10 to +15% at 50 as the ceiling. Test `TestKindlingStaveIgnitionFiresAtOneAMinute` (failed: no proc; then passed). | client ItemXItemEffect 11750; SpellAuraOptions 246565, 246447, 1291568, 1291807, 1321587; wago SpellCategories, SpellProcsPerMinute (70707, read 2026-11-02); foreverlogs.gg event caches (scratch r3 casters/hybrids, 2,150 files); https://nether.wowhead.com/forever/tooltip/spell/1300781 | CONFIRMED-BUG (rate estimated from two log samples) | E fa9d1ca90 |
| C3 | Other effects on the published warlock gear: Frozen Heart of the Mountain's Chthonic Power (1249110: +29 Frost and Shadow, school 48, 15 s, 5 min, one a fight) matches the engine; Rune of the Guard Captain's Authority (attack power) and Robe of the Void's Heal Pet, Archmagister's Faceted Pendant (experience) and Cookie's Stirring Rod (cooking) are not damage; Uther's Strength's Uther's Light is a melee-struck shield. No published warlock set carries a set bonus. | client ItemEffect / ItemXItemEffect for every item in `bis/warlock-*.json` | Aligned | - |
| C4 | Racials, hit and crit: Troll Berserking and Gnome rows unchanged in the push; Suppression is 5% hit on every spell (18174 aura 55 and 54, mask 0) and the demons take 8/17 of it on melee, all of it on spells (#394). Shadow Bolt misses 10.3 to 10.5% against the level-63 target on all three sets. | 18174; spec-breakdown | Aligned | - |
| C5 | Rotations: Affliction (Corruption, Agony, Siphon Life, Immolate, Shadow Bolt), Demonology (Imp sacrifice, Succubus, Soul Link, Doom then Agony in the last minute, Immolate, Decimation Soul Fire, Shadow Bolt) and Destruction (Corruption, Doom/Agony, Immolate, Conflagrate, Shadowburn, Incinerate) cast what their builds take; Bane of Doom and Bane of Agony never overlap (the second only while Doom is down). | APLs; spec-breakdown | Aligned | - |
| C6 | **Necropile Drain fired on every harmful cast.** Necropile Raiment's 3-piece (1299737) is the same row as Ignition (SpellAuraOptions 246447: ProcChance 100, no PPM row, 10 ms; "have a chance to steal"), but the engine read it as always proccing, the one "chance on" bonus that did not take the unstated rate. It now takes the same one a minute through the same helper. No published set wears Necropile Raiment, so no published number moves; any local run with it does. Test `TestNecropileDrainLeechesTheClientAmountAtOneAMinute` (failed, then passed). That test also divided by Hits plus ResistedHits, counting a partial resist twice; ResistedHits sit inside Hits. | client 1299737; same logs as C2 | CONFIRMED-BUG | E 3026483e6 |
| C7 | **The 10 ms recovery was unit-wide.** It exists so that one cast's several hit results proc once. The engine can land a cast-time spell in the same instant the next instant cast starts, so a unit-wide window swallowed the second cast's proc: with Ignition firing on every cast, Bane of Agony and Siphon Life (cast right after a Shadow Bolt lands) never procced, which is 15 of 78 casts a fight. The recovery is now kept per spell (`MakeHarmfulSpellProcAura`, used by Ignition and Necropile Drain). Test `TestHarmfulSpellProcRollsPerMinuteAndKeepsItsRecoveryPerSpell` (did not build before the helper; passes). | spec-breakdown probe: Ignition attempts 59.6 against 64.5 harmful casts before, 64.4 after | CONFIRMED-BUG | E 3026483e6 |

## Before/after

10000 iterations. Before: origin/main 170b7b84c on cd4dfdcc0. After: wt/r8-warlock on lane/r8-warlock fa9d1ca90. Only Ignition moves DPS (C1 and T6 touch no damage; no published set wears Necropile Raiment). Seed 1 / seed 50001.

| Spec, band, preset | Horde s1 | Horde % (1 / 50001) | Alliance s1 | Alliance % (1 / 50001) |
|---|---|---|---|---|
| Affliction 60 raid | 574.5 to 575.5 | +0.17 / +0.10 | 573.2 to 573.7 | +0.09 / +0.09 |
| Demonology 60 raid | 642.5 to 644.5 | +0.31 / +0.14 | 639.2 to 639.6 | +0.06 / +0.14 |
| Destruction 60 raid | 639.1 to 639.1 | 0.00 / 0.00 | 636.4 to 636.4 | 0.00 / 0.00 |
| Affliction 60 bare | 290.7 to 291.4 | +0.24 / +0.24 | 290.2 to 290.9 | +0.24 / +0.28 |
| Demonology 60 bare | 325.6 to 326.1 | +0.15 / +0.18 | 324.5 to 325.6 | +0.34 / +0.25 |
| Destruction 60 bare | 358.5 to 358.5 | 0.00 / 0.00 | 360.6 to 361.7 | +0.31 / +0.25 |
| Affliction 50 bare | 195.9 to 196.6 | +0.36 / +0.36 | 197.5 to 198.2 | +0.35 / +0.41 |
| Demonology 50 bare | 208.1 to 208.8 | +0.34 / +0.38 | 207.1 to 208.0 | +0.43 / +0.34 |
| Destruction 50 bare | 238.3 to 239.0 | +0.29 / +0.25 | 239.4 to 240.2 | +0.33 / +0.38 |
| Arcane 60 raid | 592.5 to 593.2 | +0.12 / +0.19 | 589.5 to 589.5 | 0.00 / 0.00 |
| Arcane 60 bare | 318.4 to 319.8 | +0.44 / +0.28 | 326.7 to 326.7 | 0.00 / 0.00 |
| Fire 50 bare | 181.8 to 182.2 | +0.22 / +0.38 | 186.9 to 187.6 | +0.37 / +0.53 |
| Elemental 50 bare | 107.2 to 107.7 | +0.47 / +0.37 | 106.0 to 106.4 | +0.38 / +0.47 |
| Fire, Frost, Shadow, Balance, Elemental 60 raid | unchanged | 0.00 | unchanged | 0.00 |

The rows read 0.00 where the published set does not wear the staff (Destruction raid and horde bare, alliance Arcane). The fork suite (`-tags=with_db ./sim/warlock/... ./sim/core/... ./sim/common/... ./sim/mage/... ./sim/shaman/... ./sim/conformance/...`) passes; no golden moved. sim/request TestRotationLadder goldens are not regenerated (the integrator's job at the pin); the warlock rungs at 50 and 60 move with C2.

### Demonology, bands 30 to 50

The bare builds at 30, 40 and 50 take Demonic Sacrifice without Demonic Pact. The shared prepull summons the Imp, sacrifices it and summons the Succubus, and the summon cancels the sacrifice, so Burning Shadow fades at about -38 s: confirmed. Measured with the new band-aware `-rotation` (horde, 10000 iterations, seeds 1 / 50001):

| Band | published | keep the sacrifice (no Succubus, no Soul Link) | no sacrifice (Succubus only) |
|---|---|---|---|
| 30 | 80.5 / 80.4 | 73.1 / 73.1 (-9.2%) | 80.5 / 80.4 (0.0%) |
| 40 | 142.0 / 141.9 | 108.0 / 107.9 (-24%) | 142.4 / 142.4 (+0.3%) |
| 50 | 208.8 / 208.7 | 165.6 / 165.8 (-21%) | 209.1 / 209.1 (+0.15%) |

Keeping the demon is right at every one of these bands. The wasted sacrifice costs 0.0 to 0.3% (prepull mana and the Succubus arriving later), below the 1% bar the rotation searches use, so the APL is unchanged. A clean fix needs the prepull gated on Demonic Pact, which the APL cannot test (the talent is passive and registers no spell for `spellIsKnown`). The alternative is a band build without Demonic Sacrifice at 30 to 50, which is a builds question.

## For other lanes

- Integrator: E fa9d1ca90 moves Arcane (60 horde), Fire (50 bare) and Elemental (50 bare) by +0.1 to +0.5% (table above), and the ranker will now see Ignition on Kindling Stave.
- Item procs owners: `core.MakeHarmfulSpellProcAura` and `core.UnstatedProcsPerMinute` are the shared rule for any other Forever "chance on harmful spell cast" item. Several are unmodeled: Galgann's Firehammer 1293980, Draconic Infused Emblem 1318931, Moonshadow 1300435, Blood Siphon 1299948, Soulrip 1307315, Blood Sacrifice 1294644. None is on a published set I checked.
- Engine (core): C1 changes the buffs-tab uptime of any prepull aura that fades or spans the pull, every class.
- Builds: Demonology 30 to 50 (above).

## Still open

- C2 and C6's rate is one a minute from two 10-proc log samples (Dream Venom, Rotmending); a level-48+ Kindling Stave log, or a Blizzard statement, would settle it.
- From round 7, unchanged: Improved Corruption with Shadow Mastery's periodic half (add or multiply), demon crit and the level-60 sheets (#331), shadow wands and Shadow multipliers (#419), Siphon Life tick crit (held for the owner).

Sources (read 2026-11-02): wago.tools layered 70707 tables (scratch `r7/casters/raw70707`, fetched 2026-11-01 21:42 UTC) and https://wago.tools/hotfixes totals; https://nether.wowhead.com/forever/tooltip/spell/1300781, /1300782, /1292997, /1322312, /8397; https://nether.wowhead.com/forever/tooltip/item/5198, /273049 (11750 returns no record).
