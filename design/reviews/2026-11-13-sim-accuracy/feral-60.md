# Sim accuracy, 2026-11-13: Feral cat and the bear at 50 and 60

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main 9b2facbc7 (client 1.60.1.70938, activated in a4bce48ae); fork from master b0a43f004 (the pin). Scope: the Feral cat (DPS) and the bear tank, bands 50 and 60, bare and raid 60: the engine against the 70938 tables for every cat and bear ability, talent and proc the cards and builds use; the open Rip and Ferocious Bite attack power shares; every threat factor on the bear; then the logs at the highest levels.

**Totals: 1 confirmed bug, fixed, old (from the 2026-11-12 feral lane's C1): Rake's hit carried the paw's weapon damage; the logs give a hit with no weapon roll at all, the rank's amount plus the same attack power share a tick takes. Cat bare 30 to 60 -1.7 to -2.5%, raid 60 -0.7 to -0.8%; Rake still beats the same list without it by 16 to 18% bare and 9.9% raid, so no rotation, build or prose advice moves. 0 new data: 70938 and its push change no druid number a card uses. The bear's per-ability threat factors stay unconfirmed: no client row, push row or Blizzard line states Maul's or Lacerate's threat, and no Forever threat measurement exists in public. Engine pin b0a43f004 to a0dbf6d8c.**

## Method

1. **Client.** The full wago.tools DB2 dumps for 1.60.1.70790 and 1.60.1.70938 with `useHotfixes=1` (SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellLevels, SpellCategories, SpellAuraOptions, SpellClassOptions, Spell, SpellEquippedItems, Talent, TraitDefinition and its effect points, CurvePoint, SkillLineAbility); every druid row (SpellClassSet 7, the druid talent ranks, every trait spell and two hops of triggers) diffed between the two builds (`feral-60/diff.py`, output `feral-60/client_diff_70790_70938.txt`), and every cat and bear rank's effect rows dumped (`feral-60/dump.py`), threat effect 63 and `BonusCoefficientFromAP` included.
2. **Blizzard and the community** on bear threat: the class-changes thread t/2360696 and the hotfix threads t/2403040 and t/2404092 to 13 November, Wowhead's Forever tooltips, Reddit, the Blizzard forums, guide sites, and the Classic threat libraries.
3. **Logs.** Our own seven uploaded reports (advanced logging: `base_amount` before armor, the caster's attack power, power and level on its own events), read from the cached parquet (`feral-60/scan.py`, `seq.py`, `var.py`, `rakeap.py`, `bite.py`; player names replaced by cat1 to cat9 in the committed tables); foreverlogs.gg's 26 public reports (2815 to 2854; none newer); Warcraft Logs has no Forever partition (hunters-60, tanks-60).
4. **Sims** on the Windows box (`tools/remote-test.sh --engine-sha`, `feral-60/jobs.sh`, one spec-breakdown at a time, `nice -n 15`, GOMAXPROCS=5, 10,000 iterations, seed 1, both factions' published races): the cat at bare 30 to 60 and raid 60 with the published list and the same list without its Rake line, at the pin and on the fix.

## Findings

| # | finding | source | origin | status |
|---|---|---|---|---|
| R1 | **Rake's hit is the rank's amount plus a tick's attack power share; it carries no weapon damage.** The engine (round 2026-11-12 C1) landed the paw's rolled weapon damage plus the rank's 27/39/56/76. Every non-crit hit of one cat in our advanced logs is the same number: cat1 at 29, 61 on 20 hits (`base_amount` = `amount`, Rake takes no armor) while its white swings rolled 56 to 69 before armor and its Claws 147 to 167; foreverlogs' Valdo (2847), one of the five cats C1 read, 48 to 49 on every hit while his ticks read 40 to 42. On every logged cat the hit sits the ranks' 7 (27 less 20) times the cat's multipliers above the tick: cat1 61 against 53, cat2 54/56 against 46/48, cat3 57 against 49. And the hit moves with attack power exactly as the tick does: cat2's advanced attack power went 584 to 619 between two fights and both its hit and its tick rose by 2. C1's "0.93 to 1.01 of the paw" came from dividing the means by an armor factor; the hit never rolled. | `feral-60/rake_own.txt`, `var.py`; foreverlogs 2847 | old (2026-11-12 C1) | **Fixed**: F a0dbf6d8c, `rakeHitDamage(roll, attackPower)` = roll + 8% of attack power, the tick's share; test F f01c02790 red at the pin |
| F1 | **Ferocious Bite's attack power share is consistent with 3.75% a point and not with Era's 3%.** Combo points are in our logs (SPELL_ENERGIZE, power 4, one per builder and per Blood Frenzy crit), so a solo cat's points at each bite are exact. cat1 at 29 (rank 1, 1328618), 23 bites read before crit and armor: 115 at 2 points (a 202 at a counted 2, the first bite of a pull, left out as a miscount), 157.2 at 3 (n=16, 150 to 162, the client's roll width), 204.5 at 4, 246.3 at 5: 44.6 a point. The flat step is the client's 30, so attack power adds 14.6 a point. Read against the same cat's Rake (its tick less the rank's 20, over the engine's 8%, after the 7 x M split above: about 362 attack power), that is 0.040 a point, against the engine's 0.0375 and Era's 0.03 (which would give 40.9). The attack power figure leans on the Rake share, which was itself fitted on stat-derived attack power, so this is a consistency check, not a measurement. | `feral-60/finishers_own.txt`, `bite.py` | - | Aligned (consistent), still open for a level 60 raid log |
| F2 | **Rip's attack power share: one logged Rip, and it may not snapshot.** cat3 at 27 (rank 2) ticked 58, 58 (crit 118), then 50, 50, 50 after leaving Cat Form mid-Rip (Rake had already fallen off before the second 58). Either the attack power part is read on every tick (the cat's form attack power gone) or Rend and Tear needs the form on every tick. The engine snapshots Rip at application. In a raid fight the cat never leaves form; only mid-Rip attack power procs would read differently. One Rip, its combo points unclear (a dodged Rip in between): open. | `feral-60/seq.py` on report gxrgvz4yzsi2 fight 102 | - | Open (O1) |
| T1 | **No Forever source states Maul's or Lacerate's threat.** The 70938 tables (push layered) carry a threat effect (63) on Demoralizing Roar alone (8/16/26/34/42 plus 0.8 a level, the engine's); no Maul, Lacerate, Swipe, Primal Bite or Faerie Fire rank has one, and threat multipliers are a server table the client never holds. Blizzard's only bear threat numbers are the forms (30 to 50% on 8 Oct, 50 to 115% on 3 Nov) and "Primal Bite: threat roughly doubled" (1 Oct); Faerie Fire "now generates threat" in the forms (16 Oct) with no amount. Wowhead's Forever tooltips say "high amount of threat" for Lacerate and Primal Bite and nothing for Maul. No threat-meter capture, log or addon table for Forever exists in public (Reddit, forums, guide sites; wowdee's 21 Oct bear page lists the per-ability values as untested). The engine's Maul and Swipe x1.75 and Faerie Fire (Feral) 108 are LibThreatClassic2's Classic Era values (`ClassModules/Classic/Druid.lua`: `Maul(amount) return amount * 1.75`; Faerie Fire 108 at rank 4, lower ranks scaled with a "NEED MORE INFO" comment), the community model Classic threat meters run on; cmangos' flat 207 for Maul rank 7 is the 2007 patch's form. Lacerate's and Primal Bite's x2.0 have no source beyond the text. | wago 70938h; t/2360696; LibThreatClassic2 | - | Unconfirmed (O4 stands) |

Strength of the bear threat evidence, for O4: Bear Form's +115% is a Blizzard statement and a push row (strong). Maul x1.75 is a Classic Era community model with no Forever confirmation (medium for Era, none for Forever). Lacerate x2.0 and Primal Bite x2.0 rest on "high amount of threat" and "roughly doubled" (weak). Faerie Fire (Feral)'s 108 is Era's rank 4 value used at every rank (weak; under 1% of a bear's threat, and not cast in the raid 60 rotation). The order bear over warrior (tanks-60 F2) therefore still rests on Maul and Lacerate; a `UnitDetailedThreatSituation` reading on one target is what would settle it.

## Checked, nothing to change (client 70938)

The druid diff 70790 to 70938 (push layered on both) moves no number a cat or bear card uses: Natural Instinct's mask gains Lacerate's hit and the Pounce Bleeds (the engine already reaches them through the crit multiplier, 2026-11-12), Pounce and Feral Charge (Cat) attribute bits, Pounce Bleed's new class mask row (which makes it show in the cat ladder's "learned but unused" lists), Wild Growth's doubled coefficient and Swiftmend's text (the healers' MR58), Rip rank 4's per-point float rounding, and Reduced Claw and Rake Cost (27851, an item effect) dropping its Mangle bit. Every rank of Rip (base and per point), Ferocious Bite (base, roll, per point, per energy), Maul, Swipe, Lacerate (bleed and hit), Primal Bite, Demoralizing Roar, Rake and Faerie Fire matches the engine as the 2026-11-12 lane recorded on 70790. No `BonusCoefficientFromAP` is set on any cat or bear rank, so the attack power shares stay server-side. Tiger's Fury, Savage Roar, Mangle and Berserk are not in the trees or the trainables (Mangle survives only in an item effect's text); Omen of Clarity, Blood Frenzy, Predatory Strikes, Heart of the Wild, Leader of the Pack, Feral Instinct, Thick Hide, Enrage, Frenzied Regeneration and Dire Bear Form are unchanged rows.

## Logs at the highest levels

No level 50 or 60 druid log exists. foreverlogs.gg has no report after 2854 (11 Nov; its tanks top out at 43, tanks-60), Warcraft Logs has no Forever partition, and our seven reports top out at a level 29 cat and a level 24 bear (`feral-60/scan.txt`: 30 druids casting cat or bear abilities, levels 7 to 29). What they do show, besides R1 and F1: a solo cat at 29 opens on Faerie Fire and Rake, builds with Claw, and finishes with Ferocious Bite at 3 points on most mobs (16 of 23 bites), Rip almost never (the mobs die first); energy is spent as it comes. Nothing there bears on a raid boss's finisher mix or the bear's rage at 60; the bear's Maul is the paw plus its flat (cat-level bears' Mauls roll with the weapon, 85 to 99 on one bear at 11), as recorded.

## Before and after (R1)

Box, 10,000 iterations, seed 1, the published entries and rotation; pin b0a43f004 against fork a0dbf6d8c (`feral-60/cat.txt`). Alliance / Horde:

| band | published rotation | Rake line over the list without it (fixed engine) |
|---|---|---|
| bare 30 | -2.45 / -2.44% | +16.0 / +16.2% |
| bare 40 | -2.34 / -2.40% | +17.2 / +17.4% |
| bare 50 | -1.97 / -1.88% | +18.1 / +18.5% |
| bare 60 | -1.79 / -1.74% | +16.4 / +17.1% |
| raid 60 | -0.73 / -0.80% | +9.85 / +9.88% |

The list without Rake is identical on both engines, so the whole move is Rake's hit. Bare 20 (level 20, before Rake) and the bear (no Rake) do not move; the bear suite golden is unchanged. The cat ladder golden's rungs 30 to 60 move -3.3 to -5.6% (the walk's bare gear); Ferocious Bite still lands at raid 60 (3.9 a fight) and bare 60 (0.95).

## Measure at refresh

| what | specs / bands | why |
|---|---|---|
| MR104: the cat on Rake's hit without weapon damage | druid-feral bare bands 30 to 60 and raid 60: picks, weights (weapon DPS no longer reaches Rake's hit, attack power and Strength do), the race sweep, the published set_dps | R1: bare -1.7 to -2.5%, raid 60 -0.7 to -0.8% |

## Open

- **O1** Rip's attack power share and whether Rip snapshots (F2): one Rip in our logs; a level 60 single-target raid log with advanced logging settles both, and Ferocious Bite's share with them (F1).
- **O4** (tanks-60) the bear's Maul, Lacerate and Primal Bite threat factors: unconfirmed (T1).
- **MR71** Pounce's attack power half: still no Pounce tick in any log.
- Not this lane: at the pin, `sim/conformance` fails on a stale `warlock.golden.md`, and the ladder goldens of mage-arcane, mage-fire, mage-frost, paladin-retribution, priest-shadow, the three warlocks and warrior-arms differ from a fresh run at their level 10 rungs (70938's starter gear ids), not touched here.

## Gates

- **Fork** (`feral-60/fork_tests.sh` on the box): `TestRakeHitIsTheRanksAmountPlusTheTicksAttackPowerShare` red at b0a43f004 (1076 for 156; 954 for 34), green at a0dbf6d8c; `go test -tags=with_db -p 1` over `sim/druid/...` and `sim/core` ok; no suite golden moved; `sim/conformance` fails on the warlock golden at the pin as well (not this change).
- **Site**: the cat ladder golden regenerated on the box (`FOREVER_UPDATE_GOLDEN=1 go test -run Ladder ./request/`, only `druid-feral.golden.md` committed); `go vet` and `go test` over `request`, `cmd/spec-breakdown` (`TestFeralCatKeepsRakeUp` included) and `specs`; `make apl-check` (the fork's card copy keeps no notes); the web guide and copy suites.
- **Review**: wow-player SHIP on the guide, card note and coach lines (`feral-60/wow-player.md`); its one check, that Rake still out-buys Shred for its energy, is the right column above.

## Sim changes in this landing

- **Fork `lane/sim-feral60`**: f01c02790 (test, red), a0dbf6d8c (fix); master fast-forwarded to a0dbf6d8c.
- **Site**: the pin (`sim/enginever/version.go`, `web/src/lib/sim/version.ts`, the sim fixtures); the Rake line's note in `data/curated/apl/druid-feral.json` and its `sim/request/apl` copy; three guide lines and the coach line; the cat ladder golden; this report, `feral-60/` (scripts, anonymized tables, the job list, the no-Rake list, the review), the ledger row and MR104.
- **No regen**: no BiS rank, data run, simdb rebuild or Data.lua render; `addon-data.json` keeps the old Rake note until the nightly.

## Sources

- wago.tools DB2 CSVs for 1.60.1.70790 and 1.60.1.70938, `useHotfixes=1`.
- Blizzard t/2360696 (1 Oct, 8 Oct, 16 Oct, 3 Nov), t/2403040, t/2404092; Wowhead Forever tooltips 9881, 1235827, 1238070.
- github.com/dfherr/LibThreatClassic2 `ClassModules/Classic/Druid.lua`; wowdee.com's Forever bear threat article (21 Oct).
- Our uploaded reports (ctxicu4cuyek, gxrgvz4yzsi2, 4q5el2xsgoy7 and four others); foreverlogs.gg 2815 to 2854.
