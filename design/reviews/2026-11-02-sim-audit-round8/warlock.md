# Sim-accuracy round 8: warlock (2026-11-02)

Specs: Affliction, Demonology, Destruction and their demons. Client 1.60.1.70707 with the hotfix push layered in (the layered `raw/` fetched 2026-11-01 21:42 UTC; `python -m pipeline push-check` logic re-run 2026-11-02: wago's push totals equal the committed `hotfix-push.json`, so no newer push). Pin cd4dfdcc0, site origin/main 170b7b84c. Lane engine `lane/r8-warlock` (fork), site branch `wt/r8-warlock`.

Method: spec-breakdown, raid preset, band 60, 10000 iterations, seeds 1 and 50001, horde troll and alliance gnome; "before" is a detached origin/main on a detached cd4dfdcc0, each with its own `make simdb`. Client rows were read from the layered SpellEffect, SpellMisc, SpellAuraOptions, SpellClassOptions, SpellCooldowns, SpellPower, SpellLevels, SkillLineAbility, ItemEffect and ItemXItemEffect, with each talent's class masks decoded against the warlock family (SpellClassSet 5). The committed `spellconst/warlock.json` regenerates byte-for-byte from the layered tables, so the fork's conformance report (which reads that file) is current.

Totals: 2 CONFIRMED-BUGs, both fixed, neither moves DPS. 1 large OPEN (Kindling Stave's Ignition).

## Assigned item 1: round 7's warlock fixes, side effects

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Mana potion now drinks on each cooldown: Major Mana Potion 1.99 a fight and Demonic Rune 2.00 in all three specs, about 1,798 and 1,195 mana each (client 17531 bp 1800; 16666 bp 1200 mana, 800 self-damage). Out of mana 0.0 s (Affliction) and 0.5 s (Demonology, Destruction). The rune and the potion are separate categories in Classic and both are used. | spec-breakdown resource table | Aligned | - |
| R2 | Elixir of Shadow Power (11474: aura 13, +40, school 32) and Elixir of Fire Power (7844: +10, school 4) now both on every warlock. Nothing else in the consume list moved; Flask of Supreme Power (17628, +150), Greater Arcane Elixir (17539, +35) and Brilliant Wizard Oil (enchant 2628, spell 25113: +36 damage, +36 healing, +1% crit) match the client. | client SpellEffect | Aligned | - |
| R3 | Life Tap: (420 + 1 a level above 56, so 424 at 60, + Spirit) x (1 + 10% a rank of Improved Life Tap), no spell power, as 11689's text ("Spirit increases the amount converted") and effect 0 (dummy 420, 1 a level) say. Spirit in `weight_stats` is still there and its test passes. | `lifetap.go`; client 11689 | Aligned | - |
| R4 | Held adoption rows. Affliction's Improved Bane of Agony (+0.7%) and Destruction's Immolate below Conflagrate (+0.1 to +0.2%) stay held: this round's engine changes (R-rows below) touch no damage, so there is no new evidence. | - | Aligned (held) | - |

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
| C2 | **Kindling Stave's Ignition is not modeled, and its rate is not in the client.** Forever binds Kindling Stave (11750, on 15 published warlock sets, Arcane mage 60 raid and bare, Fire mage and Elemental 50 bare) to a new equip spell, 1300781 Ignition: aura 42 on "deal harmful spell" (ProcTypeMask 0x10000), ProcChance 100, ProcCategoryRecovery 10 ms, triggering 1300782 for 70 Fire at Variance 0.3 ("60 to 81", Wowhead Forever spell 1300781). The text says "have a chance". Blizzard's Forever rows use ProcChance 100 both for "chance" procs (Ignition, Necropile Drain 1299737, Galgann's Firehammer 1293980, Draconic Infused Emblem 1318931) and for always-on ones written without "chance" (Sear 1291568: "Harmful spell casts sear the target"), so 100 is not the rate: the server holds it. A probe at 100% (branch `lane/r8-warlock-ignition-probe`, 0fd4c0311) gives the upper bound: Affliction +5.0%, Demonology +4.5%, Arcane +10.8% (each missile procs) at 60; Fire +15.4%, Elemental +13.2%, warlocks +10 to +12% at band 50 bare. The commit is not on the lane branch. Settled by a level-48+ log with Kindling Stave equipped (Ignition hits per harmful cast), or a Blizzard statement. | client ItemXItemEffect 11750, SpellAuraOptions 1300781, 1291568, 1293980, 1299737; https://nether.wowhead.com/forever/tooltip/spell/1300781 (read 2026-11-02) | OPEN (single source for the rate: none) | - |
| C3 | Other effects on the published warlock gear: Frozen Heart of the Mountain's Chthonic Power (1249110: +29 Frost and Shadow, school 48, 15 s, 5 min, one a fight) matches the engine; Rune of the Guard Captain's Authority (attack power) and Robe of the Void's Heal Pet, Archmagister's Faceted Pendant (experience) and Cookie's Stirring Rod (cooking) are not damage; Uther's Strength's Uther's Light is a melee-struck shield. No published warlock set carries a set bonus. | client ItemEffect / ItemXItemEffect for every item in `bis/warlock-*.json` | Aligned | - |
| C4 | Racials, hit and crit: Troll Berserking and Gnome rows unchanged in the push; Suppression is 5% hit on every spell (18174 aura 55 and 54, mask 0) and the demons take 8/17 of it on melee, all of it on spells (#394). Shadow Bolt misses 10.3 to 10.5% against the level-63 target on all three sets. | 18174; spec-breakdown | Aligned | - |
| C5 | Rotations: Affliction (Corruption, Agony, Siphon Life, Immolate, Shadow Bolt), Demonology (Imp sacrifice, Succubus, Soul Link, Doom then Agony in the last minute, Immolate, Decimation Soul Fire, Shadow Bolt) and Destruction (Corruption, Doom/Agony, Immolate, Conflagrate, Shadowburn, Incinerate) cast what their builds take; Bane of Doom and Bane of Agony never overlap (the second only while Doom is down). | APLs; spec-breakdown | Aligned | - |

## Before/after

Raid preset, band 60, 10000 iterations. Before: origin/main 170b7b84c on cd4dfdcc0. After: this branch on lane/r8-warlock c17f0c0ae. Neither fix touches damage, so DPS is identical to the decimal.

| Spec | Horde s1 | Horde s50001 | Alliance s1 | Alliance s50001 |
|---|---|---|---|---|
| Affliction | 574.5 to 574.5 (0.00%) | 575.0 to 575.0 | 573.2 to 573.2 | 573.3 to 573.3 |
| Demonology | 642.5 to 642.5 (0.00%) | 643.3 to 643.3 | 639.2 to 639.2 | 638.9 to 638.9 |
| Destruction | 639.1 to 639.1 (0.00%) | 639.5 to 639.5 | 636.4 to 636.4 | 636.4 to 636.4 |

The Ignition probe (C2, not on the branch), same method: Affliction 574.5 to 603.6 (+5.07% / +4.94%, alliance +5.01% / +4.92%), Demonology 642.5 to 671.9 (+4.58% / +4.46%, alliance +4.47% / +4.41%), Destruction raid unchanged (it wears Elemental Mage Staff).

The fork suite (`-tags=with_db ./sim/warlock/... ./sim/core/... ./sim/common/... ./sim/conformance/...`) passes; no golden moved.

## For other lanes

- Integrator / casters / hybrids: C2 decides several published sets (warlocks, Arcane, Fire 50, Elemental 50). Today the ranker scores Kindling Stave on stats alone.
- Whoever owns item_sets: Necropile Raiment's Necropile Drain (1299737) is modeled as always proccing (`alwaysProcs`, client_damage_procs.go:62), but its text is "have a chance", the same row shape as C2. Its 10 ms recovery is also unit-wide: when a cast-time spell lands in the same instant the next instant spell starts, the second cast's proc is swallowed (Bane of Agony and Siphon Life lost theirs in the probe until the recovery was kept per spell).
- Engine (core): C1 changes the buffs-tab uptime of any prepull aura that fades or spans the pull, every class.

## Still open

- C2: Ignition's rate (and Necropile Drain's).
- Demonology at bands 30 to 50 takes Demonic Sacrifice without Demonic Pact, so the prepull's Succubus summon cancels the Imp sacrifice (Burning Shadow fades at about -38 s). Whether keeping the sacrifice beats the Succubus there is a rotation-search question; `-rotation` cannot be probed below 60 because it skips the band's rank rewrite.
- From round 7, unchanged: Improved Corruption with Shadow Mastery's periodic half (add or multiply), demon crit and the level-60 sheets (#331), shadow wands and Shadow multipliers (#419), Siphon Life tick crit (held for the owner).

Sources (read 2026-11-02): wago.tools layered 70707 tables (scratch `r7/casters/raw70707`, fetched 2026-11-01 21:42 UTC) and https://wago.tools/hotfixes totals; https://nether.wowhead.com/forever/tooltip/spell/1300781, /1300782, /1292997, /1322312, /8397; https://nether.wowhead.com/forever/tooltip/item/5198, /273049 (11750 returns no record).
