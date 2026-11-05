# Sim-accuracy round 23: platform lane

Site branch `wt/r23-platform` (from int/r23 25432c0f1). Fork branch `lane/r23-platform` (from int/r23 bff826207). The lane owns sim/core, the web, the addon data and shared spell math. Client 1.60.1.70790.

Method: the four seed items in the brief's order, then a fresh re-audit of round 22's shared changes. Every sim and sim-running test went through `$R/sim.sh`. Every CONFIRMED-BUG has a test that was red before its fix.

**Confirmed and fixed: 2** (E1 `R22 a0a2818ed` on an `old` regex, M2 `old`). Plus one feature built (M1, the mana-income value). **Cross-class DPS movers: none.** No spec-breakdown number moves (checked below). E1 changes what every raid-build reader gets for Assassination only (addon default build, dps-per-point, the BiS ranker's raid pass, talent-search `-raid`).

## Live logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **Still no live Forever logs we can read.** api.foreversixty.gg `/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct). foreverlogs.gg `/api/reports/public` answers 403. Warcraft Logs v1 answers 401 on www, classic and fresh (the key in GSM is no longer accepted; not printed or stored). | last looked 11:06 UTC 5 Nov | OPEN (no data) | - | - |

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | **The mana-income APL value (round 22 C1) is built.** `APLValue.mana_income_per_second = 79 { double window_seconds = 1 }` (0 means 10 s) is the mana the character gained over the trailing window divided by the window, from every `AddMana` source (regen ticks, Judgement of Wisdom, potions, runes, Innervate), net of overflow; before the window has elapsed it divides by the elapsed time. The record lives on the mana bar (`sim/core/mana_income.go`) and is written only once a rotation builds the value, so no other rotation pays for it; it is pruned to the longest window read and cleared per iteration. Site: `apl-gates` knows the node (it gates on nothing a card names), and `apl-condition` words the budget shape `currentMana + manaIncomePerSecond x remainingTime >= ...` as "your mana and what you regain before the fight ends cover what is left to cast". The addon export only reads `currentLevel` clauses, so it needs nothing. `make apl-check` passes. Handed to the paladin lane with a usage note (`$R/platform/for-paladin.md`). | fork `proto/apl.proto`, `sim/core/apl.pb.go` (protoc 36.2, protoc-gen-go v1.36.6), `mana_income.go`, `mana.go`, `apl_value.go`, `apl_values_resources.go`; tests `TestValueManaIncomePerSecondReadsTheTrailingWindow`, `TestValueManaIncomePerSecondCountsOnlyWhatTheBarTook` (red: undefined); site `sim/request/mana_income_test.go` `TestALineGatedOnManaIncomeReadsTheCharactersRegen` (red on the pinned engine: unknown field "manaIncomePerSecond"; green: a 0.1 mana/s budget casts, a 100,000 mana/s budget never does, on a real priest); vitest `reads the mana budget` (red, then green); `go test -tags with_db ./sim/core` passes | Feature (not a bug) | - | E 61dc24ce5, S 3fd4ce4fa |
| M2 | **A product of a number and a duration panicked or multiplied nanoseconds.** The pacing budget needs `income x remainingTime`. `APLValueMath` took its type from the left operand: number x duration called `GetFloat`/`GetInt` on the duration and the engine reported "Unimplemented GetFloat" (found by M1's site test); duration x number multiplied two durations in nanoseconds. Constants too: Go's `time.ParseDuration` accepts "0", so a bare "0" const is a duration. Now a product or quotient of a number and a duration is worked in floats, the duration in seconds. Every curated and fork APL multiplies float by float (`remainingTimePercent x const`, `currentMana / currentManaPercent`), so no rotation moves. | fork `sim/core/apl_values_operators.go` (`coerceMathOperands`); test `TestMathOfANumberAndADurationWorksInSeconds` (red: helper missing; the shape panicked in the site test), covering level x time, time x level, const x time, time x const | CONFIRMED-BUG (engine, latent: no shipped line uses it) | old (upstream wowsims) | E 7bb3c301a |
| E1 | **A guide build ending in an empty tree read as the leveling build.** Assassination's raid build (round 22) spends nothing in Subtlety: `raidBuild: '...:00530311251051031/30230302020122/:'`. The three FS1 regexes in `sim/leveling/talents.go` required a digit in every tree, so the raid line never matched and `GuideRaidBuildTalents` fell back to the leveling build. This is the site's own parser, not just the scratch helper (`rt` is built from it). Every raid-build reader got Assassination's leveling build: `addon-builds` (the committed `addon-builds.json` from the round-22 regen 7a138c265 has the Assassination raid order equal to the leveling order, so the addon's "Assassination raid" default is the leveling build), `dps-per-point`, `leveling-bis`'s raid-preset pass (the queued BiS re-rank would rank Assassination raid 60 on the leveling build) and `talent-search -raid`. The tree groups now allow an empty tree, as the web's FS1 parser (`planner/fs1.ts` `parseTrees`) already did. A dry run of addon-builds (reverted, not committed) changes `rogue-assassination raid` and nothing else. `TestCommittedAddonBuildsFollowTheGuides` now fails until that file is regenerated. | `sim/leveling/talents.go`; tests `TestGuideBuildLinesKeepAnEmptyTree` (red: the leveling trees for the raid line, and no match for `/302303/`), `TestParseBuildCodeKeepsAnEmptyTree` (red: "not an FS1 build code"); `go test ./sim/leveling` and `./sim/cmd/leveling-bis -run RaidBuild` pass | CONFIRMED-BUG (site tools and published addon data) | R22 a0a2818ed (the first empty-tree build; the regex itself is old) | S d8c453fdf |
| E2 | **The scratch `rt` helper.** It is `sim/cmd/zz-raidtalents` over `leveling.GuideRaidBuildTalents`, so E1 was its bug too. A lane copy built from this branch (`$R/platform/rt`) prints `00530311251051031-30230302020122-` for Assassination and the same strings as `$R/rt` for every other spec checked. The integrator should rebuild `$R/rt` from int/r23 once it takes d8c453fdf. | `$R/platform/rt` | Note (scratch tool) | - | - |
| S1 | **Sub-20 damage-spell penalty: no new evidence.** forever-bugs has no issue on it ("downrank", "spell power level": only #401, unrelated). The web still has only SpellCoeffsForever (Classic Era model) and the Classic wikis. No logs. Round 22's call stands: no cut for heals (three sources), damage OPEN-low. | https://github.com/ClassicWoWCommunity/forever-bugs (searched 11:05 UTC 5 Nov); https://www.curseforge.com/wow/addons/spellcoeffsforever ; https://warcraft.wiki.gg/wiki/Downranking | OPEN-low (settle: a level 25 to 30 shaman's or mage's tooltip for Lightning Bolt 3 or Frostbolt 3 at two spell damage values 20+ apart) | - | - |

## Fresh re-audit of round 22's shared changes

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Unknown auras in APL conditions (E 30bf8bdb0).** Unknown `auraIsActive`/`dotIsActive` read a bool `false`, remaining time a `0ms` duration. `auraNumStacks` returns the const "0", which `newValueConst` types as a **duration** (ParseDuration accepts "0"), not an int. Compared with an int bound, both sides coerce to durations (n becomes n seconds), so 0 against any bound keeps its order and every comparison reads as with an int. Harmless; noted so nobody puts unknown stacks into math. | `sim/core/apl_values_aura.go:221-229`, `apl_values_operators.go:23-52` | Aligned (cosmetic type) | - | - |
| A2 | **Talent learn levels and card row start levels (S b4a82d5ad, 6c639ea7f, regenerated 7a138c265).** The committed addon-data matches round 22's table: Survival Counterattack from 32, Enhancement Stormstrike 25, Fury Death Wish 30, Restoration Swiftmend 43, Arcane Power 40. U1 reads `addon-builds.json`'s orders, so until E1's regeneration Assassination's raid-only talents fall back to the leveling order; they are all passives (Precision, Composure, Endurance, Flawless Execution, Weapon Expertise, Aggression), so no card row moves. | `data/builds/1.60.1.70790/addon-data.json` rotations | Aligned | - | - |
| A3 | **Raid rank rows on the card (S b00a321d1).** `addonrotation.published_preset` is the raid preset from 60, the same rule as `sim/request.PublishedPreset`; a row naming no preset applies at every level. | `data/pipeline/addonrotation.py:323-346`; `sim/request/published_presets.go:30` | Aligned | - | - |
| A4 | **Kept steps (S 03cf741de).** `continuesRun` folds a step with no note or the same note and starts a new row on a different note; generated words appear only when every step of the run reads. | `web/src/lib/bis/rotation-collapse.ts:45-62` | Aligned | - | - |
| A5 | **Raid-buff scrolls (E 145b1fdd3).** +17 only with no totem of that stat, and not twice when the consume slot holds the same scroll. A scroll still stacks with a different consume in the slot (an elixir or Juju), which matches Classic's separate scroll category, single-source as in round 22. No preset sets the flags. | `sim/core/buffs.go:328-344` | Aligned (no preset uses it) | - | - |
| A6 | **Crit-aura exclusivity (E 5d23a8747).** `RaidHasDruidCritAura` gives 3% once for Moonkin Aura or Leader of the Pack. The Moonkin form adds its own 3% only when the raid buffs carry neither (`druid.go:246-253`, `forms.go:316-327`); a feral's own Leader of the Pack sets the raid flag after that check, so it is counted once. The sourcing of "exclusive" is the druid lane's seed. | engine read | Aligned (code) | - | - |

## Before/after

No published number moves. Spot check, raid 60, horde, seed 1, 3,000 iterations, frozen BiS, guide raid build from the fixed parser on both sides: paladin-retribution 779.7 / 779.7, priest-discipline 526.0 / 526.0 HPS, rogue-assassination 779.4 / 779.4 (`$R/sb-before` from `$R/site-before/sim` against `$R/platform/sb`). M1 records gains only when a rotation reads the value; M2 changes only number-times-duration math, which no curated APL has.

## Files to regenerate (integrator, at the pin)

- `python -m pipeline simproto --engine <fork at the pin>` (data/proto/apl.proto, `pipeline/simproto/apl_pb2.py`). Until then `data/tests/test_apl.py` rejects any curated APL using `manaIncomePerSecond` (paladin lane).
- Fork `ui/core/proto` TS is not regenerated (as with currentLevel in round 19).
- `go run ./sim/cmd/addon-builds -repo-root .` (rogue-assassination raid only), then the addon data step (`addon-data.json`, `Data.lua`) and `make simdb` as the usual chain. `dps-per-point` for Assassination's raid build.
- Rebuild the scratch `$R/rt` from int/r23.

## Measure at refresh

- The BiS re-rank must run on a site with S d8c453fdf, or Assassination's raid-60 picks are ranked on the leveling build.

## For other lanes

- **Paladin:** `$R/platform/for-paladin.md`. Use `currentMana + manaIncomePerSecond x remainingTime >= <cost> x <casts left>` with OpGe; income is net of overflow (0 at a full bar) and lags inside the five-second rule; never a bare "0" const in this math.
- **Rogue:** the addon's Assassination raid default build is the leveling build until addon-builds is regenerated.

## For the guide lane

- None.

## Still open

- S1 (sub-20 damage spells), L1 (logs).
